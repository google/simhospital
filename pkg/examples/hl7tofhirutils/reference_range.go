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

package hl7tofhirutils

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
	observationpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/observation_go_proto"
	"github.com/pkg/errors"
	"github.com/google/simhospital/pkg/hl7tofhirmap"

	fhir "github.com/google/simhospital/pkg/fhircore"
	"github.com/google/simhospital/pkg/logging"
)

const ComparatorExtensionURI = "comparator"

// Characters that are not allowed for an observation value. These characters will be stripped before
// attempting to convert the value to its relevant type. Don't add any chars in here that are significant,
// like comparators (<, >, =).
var (
	blockedCharacters = []string{"[", "]"}

	numberRegexp = regexp.MustCompile(`[+-]?\d+(?:\.\d+)?(?:[eE][+-]?\d+)?`)

	comparatorsRegexp = regexp.MustCompile(fmt.Sprintf("(<|<=|>|>=) *(%v)", numberRegexp))

	rangePattern = regexp.MustCompile(fmt.Sprintf("^(%v)-(%v)(\\^(%v)\\^(%v))?$",
		numberRegexp, numberRegexp, numberRegexp, numberRegexp))

	rangePatternWithComparators = regexp.MustCompile(fmt.Sprintf("^%v([\\^]{1,2}%v)?$", comparatorsRegexp, comparatorsRegexp))

	log = logging.ForCallerPackage()
)

// ToReferenceRange returns the ReferenceRange representation of the given string.
func ToReferenceRange(ctx context.Context, str string) *observationpb.Observation_ReferenceRange {
	obsRefRange := &observationpb.Observation_ReferenceRange{
		Text: fhir.String(str),
	}

	l, h, err := parseReferenceRange(ctx, str)
	if err != nil {
		// Just ignore these fields if we cannot parse them. We send the full text in the Text
		// field, which can be inspected to find unsupported formats.
		log.WithContext(ctx).WithError(err).Debug("could not detect low and high bounds in ReferenceRange")
	}
	if l != nil {
		obsRefRange.Low = l
	}
	if h != nil {
		obsRefRange.High = h
	}
	return obsRefRange
}

// parseReferenceRange parses a string that represents a numerical range (such as "10-20") and
// returns the lower and upper limits of the range in this order. It could happen that only one
// limit makes sense, e.g., <=1 returns an upper limit but not a lower limit.
// parseReferenceRange supports ranges with format A-B or A-B^A^B, and also supports ranges enclosed
// in square brackets such as [A-B].
func parseReferenceRange(ctx context.Context, rangeStr string) (*pb.SimpleQuantity, *pb.SimpleQuantity, error) {
	// Trimming all the spaces allows us to have simpler regular expressions.
	rangeStr = strings.Replace(rangeStr, " ", "", -1)

	// Sometimes we get ranges in between square brackets. Remove the brackets.
	bracketsPattern := regexp.MustCompile(`^\[(.+)\]$`)
	if bracketsPattern.MatchString(rangeStr) {
		r := bracketsPattern.FindStringSubmatch(rangeStr)
		if len(r) == 2 {
			rangeStr = r[1]
		}
	}

	if rangeStr == "" || rangeStr == "-" || rangeStr == "[]" {
		// We get these harmless strings sometimes. Return early so that we don't try to
		// parse them and treat them as more problematic cases.
		return nil, nil, errors.New("empty ReferenceRange")
	}

	if rangePattern.MatchString(rangeStr) {
		return parseRange(ctx, rangeStr)
	}
	if rangePatternWithComparators.MatchString(rangeStr) {
		return parseRangeWithComparators(ctx, rangeStr)
	}
	return nil, nil, errors.New("unrecognized ReferenceRange format")
}

// parseRange parses a range of the form A-B^A^B or A-B and returns its lower and upper limits.
func parseRange(ctx context.Context, rangeStr string) (*pb.SimpleQuantity, *pb.SimpleQuantity, error) {
	r := rangePattern.FindStringSubmatch(rangeStr)
	// r contains something of the form [ 1-2^1^3 1 2 ^1^3 1 3 ], which has 6 elements.
	// If the range is of the form A-B, the last three groups will be the empty string.
	if len(r) != 6 {
		return nil, nil, errors.New("could not parse ReferenceRange of format A-B^A^B or A-B")
	}

	l1, err := fhir.Decimal(r[1])
	if err != nil {
		return nil, nil, errors.New("could not parse ReferenceRange of format A-B^A^B or A-B")
	}
	h1, err := fhir.Decimal(r[2])
	if err != nil {
		return nil, nil, errors.New("could not parse ReferenceRange of format A-B^A^B or A-B")
	}

	// Check if this number is of the form A-B^A^B. If so, parse the numbers after the "^" and
	// make sure they match the previous ones.
	if r[3] != "" {
		l2, err := fhir.Decimal(r[4])
		if err != nil {
			return nil, nil, errors.New("could not parse ReferenceRange of format A-B^A^B or A-B")
		}
		h2, err := fhir.Decimal(r[5])
		if err != nil {
			return nil, nil, errors.New("could not parse ReferenceRange of format A-B^A^B or A-B")
		}
		if !decimalEquals(l1, l2) || !decimalEquals(h1, h2) {
			// If the numbers do not match, it is not possible for us to know
			// which ones are the correct ones, so we consider it an error.
			return nil, nil, errors.New("could not parse ReferenceRange of format A-B^A^B or A-B")
		}
	}
	return &pb.SimpleQuantity{Value: l1}, &pb.SimpleQuantity{Value: h1}, nil
}

func decimalEquals(d1, d2 *pb.Decimal) bool {
	return d1.GetValue() == d2.GetValue()
}

// parseRangeWithComparators parses a range of the form <B, <=B, >A, >=A, <=B^^<=B, >=A^>=A or
// <B^<B and returns its lower and upper limits.
func parseRangeWithComparators(ctx context.Context, rangeStr string) (*pb.SimpleQuantity, *pb.SimpleQuantity, error) {
	r := rangePatternWithComparators.FindStringSubmatch(rangeStr)
	// r contains something of the form [>=2.5^>=2.5 >= 2.5 ^>=2.5 >= 2.5], which has 6
	// elements. If the range doesn't have a ^ character, the last three groups will be the
	// empty string.
	if len(r) != 6 {
		return nil, nil, errors.New("could not parse ReferenceRange with comparators, e.g., <=B^^<=B")
	}

	// Parse the number (and its comparator) before the "^".
	c := r[1]
	n, err := fhir.Decimal(r[2])
	if err != nil {
		return nil, nil, errors.New("could not parse ReferenceRange with comparators, e.g., <=B^^<=B")
	}
	// Check if this number is of the form <=B^^<=B, >=A^>=A or <B^<B. If so, parse the number
	// (and its comparator) after the "^" and make sure it matches the previous one.
	if r[3] != "" {
		c2 := r[4]
		n2, err := fhir.Decimal(r[5])
		if err != nil {
			return nil, nil, errors.New("could not parse ReferenceRange with comparators, e.g., <=B^^<=B")
		}
		if !decimalEquals(n, n2) || c != c2 {
			return nil, nil, errors.New("could not parse ReferenceRange with comparators, e.g., <=B^^<=B")
		}
	}
	// SimpleQuantity doesn't have a native place for comparators.
	s := &pb.SimpleQuantity{
		Value: n,
		Extension: []*pb.Extension{
			fhir.StringExtension(ComparatorExtensionURI, c),
		},
	}
	if c == ">" || c == ">=" {
		return s, nil, nil
	}
	return nil, s, nil
}

// ToQuantity returns the *pb.Quantity representation of the given string.
// This function returns an error if it fails to parse the provided string as a Decimal.
func ToQuantity(ctx context.Context, str string) (*pb.Quantity, error) {
	// Check for blocked characters
	for _, v := range blockedCharacters {
		str = strings.Replace(str, v, "", -1)
	}
	// Check for comparators.
	var comparator string
	if comparatorsRegexp.MatchString(str) {
		m := comparatorsRegexp.FindStringSubmatch(str)
		if len(m) == 3 {
			comparator = m[1]
			str = m[2]
		} else {
			log.WithContext(ctx).
				WithField("num_matches", len(m)).
				Error("Found comparators but length of matches is not 3.")
		}
	}

	decimal, err := fhir.Decimal(str)
	if err != nil {
		return nil, err
	}

	valueQuantity := &observationpb.Observation_ValueX{
		Choice: &observationpb.Observation_ValueX_Quantity{
			Quantity: &pb.Quantity{Value: decimal},
		},
	}
	if comparator != "" {
		valueQuantity.GetQuantity().Comparator = &pb.Quantity_ComparatorCode{Value: hl7tofhirmap.DefaultQuantityComparatorCodeMap[comparator]}
		valueQuantity.GetQuantity().Comparator.Extension = []*pb.Extension{
			fhir.StringExtension(ComparatorExtensionURI, comparator),
		}
	}
	return valueQuantity.GetQuantity(), nil
}
