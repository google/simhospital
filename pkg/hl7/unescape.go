// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hl7

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	// The next variables represent valid escape sequences as defined on section 2.9.1 of the
	// HL7 2.3 specification.
	hexEscapeSeq   = regexp.MustCompile(`X[A-Fa-f0-9]+`)
	localEscapeSeq = regexp.MustCompile(`Z.+`)
	spEscapeSeq    = regexp.MustCompile(`^\.sp\+?([0-9]*)$`)

	// ErrUnrecognizedEscapeSequence symbolizes an unknown or invalid HL7 escape sequence.
	ErrUnrecognizedEscapeSequence = errors.New("Unrecognized HL7 escape sequence")
)

// unescapeSP returns the expansion of an sp escape sequence, e.g. \.sp+2\
func unescapeSP(escape string) ([]byte, error) {
	n := 0
	m := spEscapeSeq.FindStringSubmatch(escape)
	if len(m) > 1 {
		if len(m[1]) == 0 {
			n = 1 // No explicit count, specification says default to 1
		} else {
			n, _ = strconv.Atoi(m[1])
		}
	}
	if n <= 0 {
		return nil, errors.New("bad .sp escape sequence: invalid number")
	}
	r := make([]byte, n, n)
	for i := 0; i < n; i++ {
		r[i] = 0xa // ASCII newline
	}
	return r, nil
}

// UnescapeText unescapes the text field src using the rules from section 2.9.1 of the
// specification, eg \F\ for the field separator (usually |).
// Unknown escape sequences cause an ErrUnrecognizedEscapeSequence.
// TX, FT and CF fields can include any of the escape sequences defined in section 2.9.1.
// ST fields can include only a subset. If the parameter `isST` is set, only such subset is
// considered valid.
func UnescapeText(src []byte, d *Delimiters, isST bool) ([]byte, error) {
	dst := make([]byte, 0, len(src))
	// Index of the start of the current escape sequence, -1 if not currently
	// within an escape sequence.
	escapeStart := -1
	// Separator for sequences that are technically invalid, but we choose to be more
	// permissive to avoid failing to process important messages.
	// We replace those sequences with invalidSequenceSeparator.
	var invalidSequenceSeparator byte = 0x20 // ASCII space.
	for i, b := range src {
		if escapeStart < 0 {
			if b != d.Escape {
				dst = append(dst, b)
			} else {
				escapeStart = i + 1
			}
		} else if escapeStart > 0 && b == d.Escape {
			v := string(src[escapeStart:i])
			switch {
			case v == "F":
				dst = append(dst, d.Field)
			case v == "S":
				dst = append(dst, d.Component)
			case v == "T":
				dst = append(dst, d.Subcomponent)
			case v == "R":
				dst = append(dst, d.Repetition)
			case v == "E":
				dst = append(dst, d.Escape)
			case v == "":
				// Empty sequences are technically not allowed.
				dst = append(dst, invalidSequenceSeparator)
			case !isST && (v == "H" || v == "N"):
			// Ignore formatting sequences.
			case !isST && v == `X000a`:
				dst = append(dst, 0xa) // ASCII new line
			case !isST && v == `X000d`:
				dst = append(dst, 0xd) // Carriage return
			case !isST && hexEscapeSeq.MatchString(v):
			case !isST && localEscapeSeq.MatchString(v):
				// Ignore locally defined escape sequences.
			case !isST && v == ".br":
				dst = append(dst, 0xa) // ASCII new line
			case isST && v == ".br":
				// This sequence is technically not allowed in ST fields.
				dst = append(dst, invalidSequenceSeparator)
			case !isST && strings.HasPrefix(v, ".sp"):
				sp, err := unescapeSP(v)
				if err != nil {
					return dst, err
				}
				dst = append(dst, sp...)
			default:
				return dst, ErrUnrecognizedEscapeSequence
			}
			escapeStart = -1
		}
	}
	if escapeStart > 0 {
		// Unterminated sequences are technically not allowed.
		dst = append(dst, invalidSequenceSeparator)
		dst = append(dst, src[escapeStart:]...)
	}
	return dst, nil
}
