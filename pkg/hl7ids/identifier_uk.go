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

// Package hl7ids contains utilities to get identifiers from HL7v2 messages.
package hl7ids

import (
	"strconv"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/hl7"
)

// Options contains the configuration for getting identifiers.
type Options struct {
	// The PID fields to look for, in this order, to find the MRNs. Values other than 2, 3 or 4 are ignored.
	MRNFrom []int
	// The function to apply to hl7.CX elements to check whether the item is a valid MRN.
	IsValidMRN func(hl7.CX) bool
	// The PID fields to look for, in this order, to find the NHS number. Values other than 2, 3 or 4 are ignored.
	NHSFrom []int
	// The function to apply to hl7.CX elements to check whether the item is a valid MRN.
	IsValidNHS func(hl7.CX) bool
}

// DefaultOptions returns the default options.
func DefaultOptions() *Options {
	return &Options{
		// PID-2 is deprecated after 3.1.
		MRNFrom: []int{3, 4},
		IsValidMRN: func(cx hl7.CX) bool {
			return HasKeywordInAnyField(cx, "MRN")
		},
		NHSFrom: []int{3},
		IsValidNHS: func(cx hl7.CX) bool {
			return HasKeywordInAnyField(cx, "NHS") && NHSNumberIsValid(cx.IDNumber.SanitizedString())
		},
	}
}

// GetMRNNumber retrieves the primary MRN number from the provided PID segment using the default options.
func GetMRNNumber(pid *hl7.PID) string {
	return GetMRNNumberWithOptions(pid, DefaultOptions())
}

// GetMRNNumberWithOptions retrieves the primary MRN number from the provided PID segment using the
// provided options.
func GetMRNNumberWithOptions(pid *hl7.PID, opts *Options) string {
	return getFirstIdentifier(opts.IsValidMRN, opts.MRNFrom, pid)
}

// GetNHSNumber retrieves the primary MRN number from the provided PID segment using the default options.
func GetNHSNumber(pid *hl7.PID) string {
	return GetNHSNumberWithOptions(pid, DefaultOptions())
}

// GetNHSNumberWithOptions retrieves the primary MRN number from the provided PID segment using the
// provided options.
func GetNHSNumberWithOptions(pid *hl7.PID, opts *Options) string {
	return getFirstIdentifier(opts.IsValidNHS, opts.NHSFrom, pid)
}

// GetAllMRNs returns all MRNs from the PID, using the default options.
func GetAllMRNs(pid *hl7.PID) []hl7.CX {
	return GetAllMRNsWithOptions(pid, DefaultOptions())
}

// GetAllMRNsWithOptions returns all MRNs from the PID, using the provided options.
func GetAllMRNsWithOptions(pid *hl7.PID, opts *Options) []hl7.CX {
	return getAllIdentifiers(opts.IsValidMRN, opts.MRNFrom, pid)
}

func getFirstIdentifier(isValid func(hl7.CX) bool, fromFields []int, pid *hl7.PID) string {
	if ids := getAllIdentifiers(isValid, fromFields, pid); len(ids) >= 1 {
		return ids[0].IDNumber.SanitizedString()
	}
	return ""
}

func getIdentifiersFrom(isValid func(hl7.CX) bool, list []hl7.CX) []hl7.CX {
	var mrns []hl7.CX
	for _, v := range list {
		if v.IDNumber == nil || !isValid(v) {
			continue
		}
		mrns = append(mrns, v)
	}
	return unique(mrns)
}

func exists(list []hl7.CX, item hl7.CX) bool {
	for _, l := range list {
		if cmp.Equal(item, l) {
			return true
		}
	}
	return false
}

func unique(list []hl7.CX) []hl7.CX {
	deduped := []hl7.CX{}
	for _, item := range list {
		if !exists(deduped, item) {
			deduped = append(deduped, item)
		}
	}
	return deduped
}

func getAllIdentifiers(isValid func(hl7.CX) bool, fromFields []int, pid *hl7.PID) []hl7.CX {
	var ids []hl7.CX

	for _, f := range fromFields {
		switch f {
		case 2:
			if pid.PatientID != nil && isValid(*pid.PatientID) {
				ids = append(ids, *pid.PatientID)
			}
		case 3:
			ids = append(ids, getIdentifiersFrom(isValid, pid.PatientIdentifierList)...)
		case 4:
			ids = append(ids, getIdentifiersFrom(isValid, pid.AlternatePatientIDPID)...)
		default:
			// Skip numbers that don't refer to PID fields with identifiers.
		}
	}
	return unique(ids)
}

// NHSNumberIsValid returns whether the given NHS Number is valid.
// An explanation of this algorithm can be found here: http://www.datadictionary.nhs.uk/version2/data_dictionary/data_field_notes/n/nhs_number_de.asp?shownav=0
func NHSNumberIsValid(nhs string) bool {
	if len(nhs) != 10 {
		return false
	}
	nhsNumber, err := strconv.Atoi(nhs)
	if err != nil {
		return false
	}
	sum := 0
	leadingDigits := nhsNumber / 10
	for i := 2; i <= 10; i++ {
		sum += (leadingDigits % 10) * i
		leadingDigits /= 10
	}
	remainder := sum % 11
	expectedCheckDigit := (11 - remainder) % 11
	return nhsNumber%10 == expectedCheckDigit
}

// HasKeywordInAnyField returns whether an identifier of type CX contains "keyword" in its
// Identifier Type Code, Assigning Facility, Assigning Authority or Code Identifying The Check
// Digit Scheme Employedfields.
func HasKeywordInAnyField(cx hl7.CX, keyword string) bool {
	return strings.Contains(cx.IdentifierTypeCode.String(), keyword) ||
		strings.Contains(cx.AssigningFacility.String(), keyword) ||
		strings.Contains(cx.AssigningAuthority.String(), keyword) ||
		strings.Contains(cx.CheckDigitScheme.String(), keyword)
}
