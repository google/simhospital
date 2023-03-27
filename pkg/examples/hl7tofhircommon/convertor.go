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

// Package hl7tofhircommon contains utilities to convert HL7v2 values into FHIR.
package hl7tofhircommon

import (
	cpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/codes_go_proto"
	"github.com/google/simhospital/pkg/hl7tofhirmap"
)

// Convertor knows how to map values found in HL7v2 messages to their FHIR equivalent.
type Convertor struct {
	*hl7tofhirmap.Convertor
	DeceasedMap     map[string]bool
	CodingSystemMap map[string]string
}

// NewConvertor with a few additions to the default mappings.
func NewConvertor() *Convertor {

	// http://hl7-definition.caristix.com:9010/Default.aspx?version=HL7%20v2.5&table=0085
	observationStatusCodeMap := map[string]cpb.ObservationStatusCode_Value{}
	for k, v := range hl7tofhirmap.DefaultObservationStatusCodeMap {
		observationStatusCodeMap[k] = v
	}
	// Record coming over is a correction and thus replaces a final result.
	observationStatusCodeMap["C"] = cpb.ObservationStatusCode_CORRECTED
	// Deletes the OBX record.
	observationStatusCodeMap["D"] = cpb.ObservationStatusCode_ENTERED_IN_ERROR
	// Alternative code for deletion of OBX record.
	observationStatusCodeMap["INERR2"] = cpb.ObservationStatusCode_ENTERED_IN_ERROR
	// Final results; Can only be changed with a corrected result.
	observationStatusCodeMap["F"] = cpb.ObservationStatusCode_FINAL
	// Specimen in lab; results pending
	observationStatusCodeMap["I"] = cpb.ObservationStatusCode_REGISTERED
	// HL7OBXResultPendingNonStandard is the same as Pending, but is not included in HL7 standard
	observationStatusCodeMap["IP"] = cpb.ObservationStatusCode_REGISTERED
	// Not asked; used to affirmatively document that the observation identified in the OBX was not sought when
	// the universal service ID in OBR-4 implies that it would be sought.
	observationStatusCodeMap["N"] = cpb.ObservationStatusCode_UNKNOWN
	// Order detail description only (no result)
	observationStatusCodeMap["O"] = cpb.ObservationStatusCode_REGISTERED
	// Preliminary results
	observationStatusCodeMap["P"] = cpb.ObservationStatusCode_PRELIMINARY
	// Results entered -- not verified
	observationStatusCodeMap["R"] = cpb.ObservationStatusCode_PRELIMINARY
	// Partial results
	observationStatusCodeMap["S"] = cpb.ObservationStatusCode_PRELIMINARY
	// Results cannot be obtained for this observation
	observationStatusCodeMap["X"] = cpb.ObservationStatusCode_PRELIMINARY
	// Results status change to final without retransmitting results already sent as ‘preliminary.’
	// E.g., radiology changes status from preliminary to final
	observationStatusCodeMap["U"] = cpb.ObservationStatusCode_FINAL
	// Post original as wrong, e.g., transmitted for wrong patient
	observationStatusCodeMap["W"] = cpb.ObservationStatusCode_CANCELLED

	drStatusCodeMap := map[string]cpb.DiagnosticReportStatusCode_Value{}
	for k, v := range hl7tofhirmap.DefaultDiagnosticReportStatusCodeMap {
		drStatusCodeMap[k] = v
	}
	drStatusCodeMap["A"] = cpb.DiagnosticReportStatusCode_PARTIAL
	drStatusCodeMap["R"] = cpb.DiagnosticReportStatusCode_PARTIAL
	drStatusCodeMap["C"] = cpb.DiagnosticReportStatusCode_CORRECTED
	drStatusCodeMap["F"] = cpb.DiagnosticReportStatusCode_FINAL
	drStatusCodeMap["I"] = cpb.DiagnosticReportStatusCode_REGISTERED
	drStatusCodeMap["O"] = cpb.DiagnosticReportStatusCode_REGISTERED
	drStatusCodeMap["P"] = cpb.DiagnosticReportStatusCode_PRELIMINARY
	drStatusCodeMap["S"] = cpb.DiagnosticReportStatusCode_REGISTERED
	drStatusCodeMap["X"] = cpb.DiagnosticReportStatusCode_CANCELLED
	drStatusCodeMap["Y"] = cpb.DiagnosticReportStatusCode_UNKNOWN
	drStatusCodeMap["Z"] = cpb.DiagnosticReportStatusCode_UNKNOWN

	reqStatusCodeMap := map[string]cpb.RequestStatusCode_Value{}
	for k, v := range hl7tofhirmap.DefaultRequestStatusCodeMap {
		reqStatusCodeMap[k] = v
	}
	reqStatusCodeMap["CM"] = cpb.RequestStatusCode_COMPLETED
	reqStatusCodeMap["DC"] = cpb.RequestStatusCode_UNKNOWN
	reqStatusCodeMap["A"] = cpb.RequestStatusCode_UNKNOWN
	reqStatusCodeMap["RP"] = cpb.RequestStatusCode_UNKNOWN

	// We map most of these to UNKNOWN, because the statuses that we'd like to set
	// (see the comments) aren't part of RequestStatusCode_Value.
	reqStatusCodeMap["NW"] = cpb.RequestStatusCode_UNKNOWN       // Should be REQUESTED
	reqStatusCodeMap["ORDERED"] = cpb.RequestStatusCode_UNKNOWN  // Should be REQUESTED
	reqStatusCodeMap["IL"] = cpb.RequestStatusCode_UNKNOWN       // Should be IN_PROGRESS
	reqStatusCodeMap["INLAB"] = cpb.RequestStatusCode_UNKNOWN    // Should be IN_PROGRESS
	reqStatusCodeMap["DP"] = cpb.RequestStatusCode_UNKNOWN       // Should be IN_PROGRESS
	reqStatusCodeMap["DISPATCH"] = cpb.RequestStatusCode_UNKNOWN // Should be IN_PROGRESS
	reqStatusCodeMap["SCHED"] = cpb.RequestStatusCode_UNKNOWN    // Should be PLANNED
	reqStatusCodeMap["SC"] = cpb.RequestStatusCode_UNKNOWN       // Should be PLANNED
	reqStatusCodeMap["SUSPEND"] = cpb.RequestStatusCode_UNKNOWN  // Should be SUSPENDED
	reqStatusCodeMap["HD"] = cpb.RequestStatusCode_UNKNOWN       // Should be SUSPENDED
	reqStatusCodeMap["TRANCANC"] = cpb.RequestStatusCode_UNKNOWN // Should be CANCELLED

	nameUseCodeMap := map[string]cpb.NameUseCode_Value{}
	for k, v := range hl7tofhirmap.DefaultNameUseCodeMap {
		nameUseCodeMap[k] = v
	}
	nameUseCodeMap["CURRENT"] = cpb.NameUseCode_USUAL
	nameUseCodeMap["HISTORIC"] = cpb.NameUseCode_OLD

	contactPointUseCodeMap := map[string]cpb.ContactPointUseCode_Value{}
	for k, v := range hl7tofhirmap.DefaultContactPointUseCodeMap {
		contactPointUseCodeMap[k] = v
	}
	contactPointUseCodeMap["ORN"] = cpb.ContactPointUseCode_HOME
	contactPointUseCodeMap["PRN"] = cpb.ContactPointUseCode_HOME
	contactPointUseCodeMap["VHN"] = cpb.ContactPointUseCode_HOME
	contactPointUseCodeMap["WPN"] = cpb.ContactPointUseCode_WORK
	contactPointUseCodeMap["BUSINESS"] = cpb.ContactPointUseCode_WORK

	genderMap := map[string]cpb.AdministrativeGenderCode_Value{}
	for k, v := range hl7tofhirmap.DefaultAdministrativeGenderCodeMap {
		genderMap[k] = v
	}
	genderMap["1"] = cpb.AdministrativeGenderCode_MALE
	genderMap["2"] = cpb.AdministrativeGenderCode_FEMALE
	genderMap["0"] = cpb.AdministrativeGenderCode_UNKNOWN
	genderMap["9"] = cpb.AdministrativeGenderCode_OTHER
	genderMap["M"] = cpb.AdministrativeGenderCode_MALE
	genderMap["F"] = cpb.AdministrativeGenderCode_FEMALE
	genderMap["U"] = cpb.AdministrativeGenderCode_UNKNOWN
	genderMap["m"] = cpb.AdministrativeGenderCode_MALE
	genderMap["f"] = cpb.AdministrativeGenderCode_FEMALE
	genderMap["u"] = cpb.AdministrativeGenderCode_UNKNOWN

	reqPriorityCode := map[string]cpb.RequestPriorityCode_Value{}
	for k, v := range hl7tofhirmap.DefaultRequestPriorityCodeMap {
		reqPriorityCode[k] = v
	}
	reqPriorityCode["HI"] = cpb.RequestPriorityCode_URGENT

	aiCategory := map[string]cpb.AllergyIntoleranceCategoryCode_Value{}
	for k, v := range hl7tofhirmap.DefaultAllergyIntoleranceCategoryCodeMap {
		aiCategory[k] = v
	}
	aiCategory["DRUG"] = cpb.AllergyIntoleranceCategoryCode_MEDICATION

	aiSeverity := map[string]cpb.AllergyIntoleranceSeverityCode_Value{}
	for k, v := range hl7tofhirmap.DefaultAllergyIntoleranceSeverityCodeMap {
		aiSeverity[k] = v
	}
	aiSeverity["LOW"] = cpb.AllergyIntoleranceSeverityCode_MILD
	aiSeverity["MEDIUM"] = cpb.AllergyIntoleranceSeverityCode_MODERATE
	aiSeverity["HIGH"] = cpb.AllergyIntoleranceSeverityCode_SEVERE

	return &Convertor{
		Convertor: &hl7tofhirmap.Convertor{
			NameUseCodeMap:                    nameUseCodeMap,
			ContactPointUseCodeMap:            contactPointUseCodeMap,
			ObservationStatusCodeMap:          observationStatusCodeMap,
			AdministrativeGenderCodeMap:       genderMap,
			DiagnosticReportStatusCodeMap:     drStatusCodeMap,
			RequestStatusCodeMap:              reqStatusCodeMap,
			RequestPriorityCodeMap:            reqPriorityCode,
			AllergyIntoleranceSeverityCodeMap: aiSeverity,
			AllergyIntoleranceCategoryCodeMap: aiCategory,
		},
		DeceasedMap: map[string]bool{
			"YES":      true,
			"DECEASED": true,
			"Y":        true,
			"NO":       false,
			"N":        false,
		},
		CodingSystemMap: map[string]string{
			"SNM3": "http://snomed.info/sct",
			"ACME": "https://acme.lab/resultcodes",
		},
	}
}
