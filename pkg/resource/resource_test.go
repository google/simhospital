// Copyright 2020 Google LLC
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

package resource

import (
	"bytes"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/testing/protocmp"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/constants"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/test/testid"
	"github.com/google/simhospital/pkg/test/testresource"

	cpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/codes_go_proto"
	dpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
	aipb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/allergy_intolerance_go_proto"
	r4pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/bundle_and_contained_resource_go_proto"
	conditionpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/condition_go_proto"
	encounterpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/encounter_go_proto"
	locationpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/location_go_proto"
	observationpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/observation_go_proto"
	patientpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/patient_go_proto"
	practitionerpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/practitioner_go_proto"
	procedurepb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/procedure_go_proto"
)

var (
	delay           = time.Hour * 5
	now             = ir.NewValidTime(time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC))
	later           = ir.NewValidTime(now.Add(delay))
	evenLater       = ir.NewValidTime(later.Add(delay))
	nowMicros       = now.UnixNano() / 1000
	laterMicros     = later.UnixNano() / 1000
	evenLaterMicros = evenLater.UnixNano() / 1000
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name        string
		patientInfo *ir.PatientInfo
		bundleType  string
		want        *r4pb.Bundle
	}{{
		name:       "Patient with all fields",
		bundleType: Batch,
		patientInfo: &ir.PatientInfo{
			Person: &ir.Person{
				MRN:         "1234",
				Prefix:      "Dr",
				FirstName:   "William",
				MiddleName:  "George",
				Surname:     "Burr",
				Suffix:      "MD",
				PhoneNumber: "01234567890",
				Gender:      "M",
				DateOfDeath: now,
				Address: &ir.Address{
					FirstLine:  "FIRST_LINE",
					SecondLine: "SECOND_LINE",
					City:       "CITY",
					PostalCode: "ABC DEF",
					Country:    "COUNTRY",
					Type:       "HOME",
				},
			},
			Class: "IMP",
			Allergies: []*ir.Allergy{{
				Description: ir.CodedElement{
					ID:            "ID",
					Text:          "TEXT",
					CodingSystem:  "SYSTEM",
					AlternateText: "ALTERNATE_TEXT",
				},
				Severity:               "SEVERE",
				Reaction:               "REACTION",
				Type:                   "FOOD",
				IdentificationDateTime: now,
			}, {
				Description: ir.CodedElement{
					ID:            "ID",
					Text:          "TEXT",
					CodingSystem:  "SYSTEM",
					AlternateText: "ALTERNATE_TEXT",
				},
				Severity:               "MODERATE",
				Reaction:               "REACTION",
				Type:                   "MEDICATION",
				IdentificationDateTime: now,
			}},
			Encounters: []*ir.Encounter{{
				Status:      constants.EncounterStatusFinished,
				StatusStart: evenLater,
				Start:       now,
				End:         evenLater,
				StatusHistory: []*ir.StatusHistory{{
					Status: constants.EncounterStatusPlanned,
					Start:  now,
					End:    now,
				}, {
					Status: constants.EncounterStatusArrived,
					Start:  now,
					End:    later,
				}, {
					Status: constants.EncounterStatusInProgress,
					Start:  later,
					End:    evenLater,
				}},
				Procedures: []*ir.DiagnosisOrProcedure{{
					Description: &ir.CodedElement{
						ID:           "ID_1",
						Text:         "PROCEDURE",
						CodingSystem: "SYSTEM",
					},
					Type: "TYPE",
					Clinician: &ir.Doctor{
						ID:        "ID",
						Prefix:    "Dr",
						FirstName: "Doctor",
						Surname:   "Doctorson",
						Specialty: "Doctoring",
					},
					DateTime: now,
				}, {
					Description: &ir.CodedElement{
						ID:           "ID_2",
						Text:         "PROCEDURE",
						CodingSystem: "SYSTEM",
					},
					Type: "TYPE",
					Clinician: &ir.Doctor{
						ID:        "ID",
						Prefix:    "Dr",
						FirstName: "Doctor",
						Surname:   "Doctorson",
						Specialty: "Doctoring",
					},
					DateTime: later,
				}},
				Diagnoses: []*ir.DiagnosisOrProcedure{{
					Description: &ir.CodedElement{
						ID:           "ID",
						Text:         "DIAGNOSIS",
						CodingSystem: "SYSTEM",
					},
					Type: "TYPE",
					Clinician: &ir.Doctor{
						ID:        "ID",
						Prefix:    "Dr",
						FirstName: "Doctor",
						Surname:   "Doctorson",
						Specialty: "Doctoring",
					},
					DateTime: later,
				}},
				Orders: []*ir.Order{{
					OrderDateTime: later,
					Results: []*ir.Result{{
						TestName: &ir.CodedElement{
							ID:           "TEST_ID_1",
							Text:         "TEST_NAME_1",
							CodingSystem: "SYSTEM",
						},
						Value:        "VALUE",
						Unit:         "UNIT",
						AbnormalFlag: "H",
						Notes:        []string{"NOTE_1", "NOTE_2"},
						Status:       "C",
					}, {
						TestName: &ir.CodedElement{
							ID:           "TEST_ID_2",
							Text:         "TEST_NAME_2",
							CodingSystem: "SYSTEM",
						},
						Value:  "VALUE",
						Unit:   "UNIT",
						Status: "F",
					}},
				}},
				LocationHistory: []*ir.LocationHistory{{
					Location: &ir.PatientLocation{
						Poc:      "POC",
						Room:     "ROOM",
						Bed:      "BED",
						Floor:    "FLOOR",
						Building: "BUILDING",
						Facility: "FACILITY",
					},
					Start: now,
					End:   later,
				}, {
					Location: &ir.PatientLocation{
						Poc:      "POC",
						Room:     "ROOM",
						Bed:      "BED",
						Floor:    "FLOOR",
						Building: "BUILDING",
						Facility: "FACILITY",
					},
					Start: later,
					End:   later,
				}, {
					Location: &ir.PatientLocation{
						Building: "BUILDING",
					},
					Start: later,
					End:   evenLater,
				}},
			}, {
				Status:      constants.EncounterStatusInProgress,
				StatusStart: evenLater,
				Start:       evenLater,
			}},
		},
		want: &r4pb.Bundle{
			Type: &r4pb.Bundle_TypeCode{Value: cpb.BundleTypeCode_BATCH},
			Entry: []*r4pb.Bundle_Entry{{
				FullUrl: &dpb.Uri{Value: "Patient/1"},
				Request: &r4pb.Bundle_Entry_Request{
					Method: &r4pb.Bundle_Entry_Request_MethodCode{Value: cpb.HTTPVerbCode_POST},
					Url:    &dpb.Uri{Value: "Patient"},
				},
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Patient{
						&patientpb.Patient{
							Id:         &dpb.Id{Value: "1"},
							Identifier: []*dpb.Identifier{{Value: &dpb.String{Value: "1234"}}},
							Text: &dpb.Narrative{
								Div:    &dpb.Xhtml{Value: "<div><p>Dr William George Burr MD</p></div>"},
								Status: &dpb.Narrative_StatusCode{Value: cpb.NarrativeStatusCode_GENERATED},
							},
							Name: []*dpb.HumanName{{
								Prefix: []*dpb.String{{Value: "Dr"}},
								Family: &dpb.String{Value: "Burr"},
								Given:  []*dpb.String{{Value: "William"}, {Value: "George"}},
								Suffix: []*dpb.String{{Value: "MD"}},
							}},
							Gender: &patientpb.Patient_GenderCode{Value: cpb.AdministrativeGenderCode_MALE},
							Telecom: []*dpb.ContactPoint{{
								Value:  &dpb.String{Value: "01234567890"},
								System: &dpb.ContactPoint_SystemCode{Value: cpb.ContactPointSystemCode_PHONE},
								Use:    &dpb.ContactPoint_UseCode{Value: cpb.ContactPointUseCode_HOME},
							}},
							Deceased: &patientpb.Patient_DeceasedX{
								Choice: &patientpb.Patient_DeceasedX_DateTime{
									DateTime: &dpb.DateTime{ValueUs: nowMicros, Precision: dpb.DateTime_SECOND},
								},
							},
							Address: []*dpb.Address{{
								Use:  &dpb.Address_UseCode{Value: cpb.AddressUseCode_HOME},
								Type: &dpb.Address_TypeCode{Value: cpb.AddressTypeCode_BOTH},
								Line: []*dpb.String{{
									Value: "FIRST_LINE",
								}, {
									Value: "SECOND_LINE",
								}},
								City:       &dpb.String{Value: "CITY"},
								PostalCode: &dpb.String{Value: "ABC DEF"},
								Country:    &dpb.String{Value: "COUNTRY"},
							}},
						},
					},
				},
			}, {
				FullUrl: &dpb.Uri{Value: "AllergyIntolerance/2"},
				Request: &r4pb.Bundle_Entry_Request{
					Method: &r4pb.Bundle_Entry_Request_MethodCode{Value: cpb.HTTPVerbCode_POST},
					Url:    &dpb.Uri{Value: "AllergyIntolerance"},
				},
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_AllergyIntolerance{
						&aipb.AllergyIntolerance{
							Id:   &dpb.Id{Value: "2"},
							Type: &aipb.AllergyIntolerance_TypeCode{Value: cpb.AllergyIntoleranceTypeCode_ALLERGY},
							ClinicalStatus: &dpb.CodeableConcept{
								Coding: []*dpb.Coding{{
									Code: &dpb.Code{
										Value: "active",
									},
									System: &dpb.Uri{
										Value: "http://terminology.hl7.org/CodeSystem/allergyintolerance-clinical",
									},
									Display: &dpb.String{Value: "Active"},
								}},
							},
							RecordedDate: &dpb.DateTime{ValueUs: nowMicros, Precision: dpb.DateTime_SECOND},
							Category: []*aipb.AllergyIntolerance_CategoryCode{{
								Value: cpb.AllergyIntoleranceCategoryCode_FOOD,
							}},
							Patient: &dpb.Reference{
								Reference: &dpb.Reference_PatientId{
									&dpb.ReferenceId{Value: "1"},
								},
								Display: &dpb.String{Value: "William Burr"},
							},
							Code: &dpb.CodeableConcept{
								Coding: []*dpb.Coding{{
									System:  &dpb.Uri{Value: "SYSTEM_URI"},
									Code:    &dpb.Code{Value: "ID"},
									Display: &dpb.String{Value: "TEXT"},
								}},
							},
							Reaction: []*aipb.AllergyIntolerance_Reaction{{
								Manifestation: []*dpb.CodeableConcept{{
									Text: &dpb.String{Value: "REACTION"},
								}},
								Severity: &aipb.AllergyIntolerance_Reaction_SeverityCode{
									Value: cpb.AllergyIntoleranceSeverityCode_SEVERE,
								},
							}},
						},
					},
				},
			}, {
				FullUrl: &dpb.Uri{Value: "AllergyIntolerance/3"},
				Request: &r4pb.Bundle_Entry_Request{
					Method: &r4pb.Bundle_Entry_Request_MethodCode{Value: cpb.HTTPVerbCode_POST},
					Url:    &dpb.Uri{Value: "AllergyIntolerance"},
				},
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_AllergyIntolerance{
						&aipb.AllergyIntolerance{
							Id: &dpb.Id{Value: "3"},
							ClinicalStatus: &dpb.CodeableConcept{
								Coding: []*dpb.Coding{{
									Code: &dpb.Code{
										Value: "active",
									},
									System:  &dpb.Uri{Value: "http://terminology.hl7.org/CodeSystem/allergyintolerance-clinical"},
									Display: &dpb.String{Value: "Active"},
								}},
							},
							Type:         &aipb.AllergyIntolerance_TypeCode{Value: cpb.AllergyIntoleranceTypeCode_ALLERGY},
							RecordedDate: &dpb.DateTime{ValueUs: nowMicros, Precision: dpb.DateTime_SECOND},
							Category: []*aipb.AllergyIntolerance_CategoryCode{{
								Value: cpb.AllergyIntoleranceCategoryCode_MEDICATION,
							}},
							Code: &dpb.CodeableConcept{
								Coding: []*dpb.Coding{
									{
										System:  &dpb.Uri{Value: "SYSTEM_URI"},
										Code:    &dpb.Code{Value: "ID"},
										Display: &dpb.String{Value: "TEXT"},
									},
								},
							},
							Patient: &dpb.Reference{
								Reference: &dpb.Reference_PatientId{
									&dpb.ReferenceId{Value: "1"},
								},
								Display: &dpb.String{Value: "William Burr"},
							},
							Reaction: []*aipb.AllergyIntolerance_Reaction{{
								Manifestation: []*dpb.CodeableConcept{{
									Text: &dpb.String{Value: "REACTION"},
								}},
								Severity: &aipb.AllergyIntolerance_Reaction_SeverityCode{
									Value: cpb.AllergyIntoleranceSeverityCode_MODERATE,
								},
							}},
						},
					},
				},
			}, {
				FullUrl: &dpb.Uri{Value: "Location/5"},
				Request: &r4pb.Bundle_Entry_Request{
					Method: &r4pb.Bundle_Entry_Request_MethodCode{Value: cpb.HTTPVerbCode_POST},
					Url:    &dpb.Uri{Value: "Location"},
				},
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Location{
						&locationpb.Location{
							Id:   &dpb.Id{Value: "5"},
							Name: &dpb.String{Value: "BED, POC, ROOM, FLOOR, BUILDING, FACILITY"},
							Text: &dpb.Narrative{
								Div: &dpb.Xhtml{
									Value: "<div><p>BED, POC, ROOM, FLOOR, BUILDING, FACILITY</p></div>",
								},
								Status: &dpb.Narrative_StatusCode{Value: cpb.NarrativeStatusCode_GENERATED},
							},
						},
					},
				},
			}, {
				FullUrl: &dpb.Uri{Value: "Location/6"},
				Request: &r4pb.Bundle_Entry_Request{
					Method: &r4pb.Bundle_Entry_Request_MethodCode{Value: cpb.HTTPVerbCode_POST},
					Url:    &dpb.Uri{Value: "Location"},
				},
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Location{
						&locationpb.Location{
							Id:   &dpb.Id{Value: "6"},
							Name: &dpb.String{Value: "BUILDING"},
							Text: &dpb.Narrative{
								Div:    &dpb.Xhtml{Value: "<div><p>BUILDING</p></div>"},
								Status: &dpb.Narrative_StatusCode{Value: cpb.NarrativeStatusCode_GENERATED},
							},
						},
					},
				},
			}, {
				FullUrl: &dpb.Uri{Value: "Practitioner/7"},
				Request: &r4pb.Bundle_Entry_Request{
					Method: &r4pb.Bundle_Entry_Request_MethodCode{Value: cpb.HTTPVerbCode_POST},
					Url:    &dpb.Uri{Value: "Practitioner"},
				},
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Practitioner{
						&practitionerpb.Practitioner{
							Id: &dpb.Id{Value: "7"},
							Identifier: []*dpb.Identifier{{
								Value: &dpb.String{Value: "ID"},
							}},
							Text: &dpb.Narrative{
								Div:    &dpb.Xhtml{Value: "<div><p>Dr Doctor Doctorson</p></div>"},
								Status: &dpb.Narrative_StatusCode{Value: cpb.NarrativeStatusCode_GENERATED},
							},
							Name: []*dpb.HumanName{{
								Family: &dpb.String{Value: "Doctorson"},
								Given:  []*dpb.String{{Value: "Doctor"}},
								Prefix: []*dpb.String{{Value: "Dr"}},
							}},
						},
					},
				},
			}, {
				FullUrl: &dpb.Uri{Value: "Procedure/8"},
				Request: &r4pb.Bundle_Entry_Request{
					Method: &r4pb.Bundle_Entry_Request_MethodCode{Value: cpb.HTTPVerbCode_POST},
					Url:    &dpb.Uri{Value: "Procedure"},
				},
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Procedure{
						&procedurepb.Procedure{
							Id: &dpb.Id{Value: "8"},
							Text: &dpb.Narrative{
								Div:    &dpb.Xhtml{Value: "<div><p>PROCEDURE by Dr Doctor Doctorson</p></div>"},
								Status: &dpb.Narrative_StatusCode{Value: cpb.NarrativeStatusCode_GENERATED},
							},
							Status: &procedurepb.Procedure_StatusCode{
								Value: cpb.EventStatusCode_COMPLETED,
							},
							Category: &dpb.CodeableConcept{
								Text: &dpb.String{Value: "TYPE"},
							},
							Code: &dpb.CodeableConcept{
								Coding: []*dpb.Coding{{
									Code:    &dpb.Code{Value: "ID_1"},
									System:  &dpb.Uri{Value: "SYSTEM_URI"},
									Display: &dpb.String{Value: "PROCEDURE"},
								}},
							},
							Performer: []*procedurepb.Procedure_Performer{{
								Actor: &dpb.Reference{
									Reference: &dpb.Reference_PractitionerId{
										&dpb.ReferenceId{Value: "7"},
									},
									Display: &dpb.String{Value: "Doctor Doctorson"},
								},
							}},
							Performed: &procedurepb.Procedure_PerformedX{
								Choice: &procedurepb.Procedure_PerformedX_DateTime{
									&dpb.DateTime{ValueUs: nowMicros, Precision: dpb.DateTime_SECOND},
								},
							},
							Encounter: &dpb.Reference{
								Reference: &dpb.Reference_EncounterId{
									&dpb.ReferenceId{Value: "4"},
								},
							},
							Subject: &dpb.Reference{
								Reference: &dpb.Reference_PatientId{
									PatientId: &dpb.ReferenceId{Value: "1"},
								},
								Display: &dpb.String{Value: "William Burr"},
							},
						},
					},
				},
			}, {
				FullUrl: &dpb.Uri{Value: "Procedure/9"},
				Request: &r4pb.Bundle_Entry_Request{
					Method: &r4pb.Bundle_Entry_Request_MethodCode{Value: cpb.HTTPVerbCode_POST},
					Url:    &dpb.Uri{Value: "Procedure"},
				},
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Procedure{
						&procedurepb.Procedure{
							Id: &dpb.Id{Value: "9"},
							Text: &dpb.Narrative{
								Div:    &dpb.Xhtml{Value: "<div><p>PROCEDURE by Dr Doctor Doctorson</p></div>"},
								Status: &dpb.Narrative_StatusCode{Value: cpb.NarrativeStatusCode_GENERATED},
							},
							Status: &procedurepb.Procedure_StatusCode{
								Value: cpb.EventStatusCode_COMPLETED,
							},
							Category: &dpb.CodeableConcept{
								Text: &dpb.String{Value: "TYPE"},
							},
							Code: &dpb.CodeableConcept{
								Coding: []*dpb.Coding{{
									Code:    &dpb.Code{Value: "ID_2"},
									System:  &dpb.Uri{Value: "SYSTEM_URI"},
									Display: &dpb.String{Value: "PROCEDURE"},
								}},
							},
							Performer: []*procedurepb.Procedure_Performer{{
								Actor: &dpb.Reference{
									Reference: &dpb.Reference_PractitionerId{
										&dpb.ReferenceId{Value: "7"},
									},
									Display: &dpb.String{Value: "Doctor Doctorson"},
								},
							}},
							Performed: &procedurepb.Procedure_PerformedX{
								Choice: &procedurepb.Procedure_PerformedX_DateTime{
									&dpb.DateTime{ValueUs: laterMicros, Precision: dpb.DateTime_SECOND},
								},
							},
							Encounter: &dpb.Reference{
								Reference: &dpb.Reference_EncounterId{
									&dpb.ReferenceId{Value: "4"},
								},
							},
							Subject: &dpb.Reference{
								Reference: &dpb.Reference_PatientId{
									PatientId: &dpb.ReferenceId{Value: "1"},
								},
								Display: &dpb.String{Value: "William Burr"},
							},
						},
					},
				},
			}, {
				FullUrl: &dpb.Uri{Value: "Condition/10"},
				Request: &r4pb.Bundle_Entry_Request{
					Method: &r4pb.Bundle_Entry_Request_MethodCode{Value: cpb.HTTPVerbCode_POST},
					Url:    &dpb.Uri{Value: "Condition"},
				},
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Condition{
						&conditionpb.Condition{
							Id: &dpb.Id{Value: "10"},
							Text: &dpb.Narrative{
								Div:    &dpb.Xhtml{Value: "<div><p>DIAGNOSIS by Dr Doctor Doctorson</p></div>"},
								Status: &dpb.Narrative_StatusCode{Value: cpb.NarrativeStatusCode_GENERATED},
							},
							Code: &dpb.CodeableConcept{
								Coding: []*dpb.Coding{{
									Code:    &dpb.Code{Value: "ID"},
									System:  &dpb.Uri{Value: "SYSTEM_URI"},
									Display: &dpb.String{Value: "DIAGNOSIS"},
								}},
							},
							Recorder: &dpb.Reference{
								Reference: &dpb.Reference_PractitionerId{
									&dpb.ReferenceId{Value: "7"},
								},
								Display: &dpb.String{Value: "Doctor Doctorson"},
							},
							RecordedDate: &dpb.DateTime{ValueUs: laterMicros, Precision: dpb.DateTime_SECOND},
							Encounter: &dpb.Reference{
								Reference: &dpb.Reference_EncounterId{
									&dpb.ReferenceId{Value: "4"},
								},
							},
							Subject: &dpb.Reference{
								Reference: &dpb.Reference_PatientId{
									PatientId: &dpb.ReferenceId{Value: "1"},
								},
								Display: &dpb.String{Value: "William Burr"},
							},
						},
					},
				},
			}, {
				FullUrl: &dpb.Uri{Value: "Encounter/4"},
				Request: &r4pb.Bundle_Entry_Request{
					Method: &r4pb.Bundle_Entry_Request_MethodCode{Value: cpb.HTTPVerbCode_POST},
					Url:    &dpb.Uri{Value: "Encounter"},
				},
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Encounter{
						&encounterpb.Encounter{
							Id: &dpb.Id{Value: "4"},
							ClassValue: &dpb.Coding{
								Code: &dpb.Code{Value: "IMP"},
							},
							Status: &encounterpb.Encounter_StatusCode{Value: cpb.EncounterStatusCode_FINISHED},
							Text: &dpb.Narrative{
								Div: &dpb.Xhtml{
									Value: "<div><p>Status: finished</p><p>Active from Mon Feb 12 00:00:00 2018 until Mon Feb 12 10:00:00 2018</p></div>",
								},
								Status: &dpb.Narrative_StatusCode{Value: cpb.NarrativeStatusCode_GENERATED},
							},
							Period: &dpb.Period{
								Start: &dpb.DateTime{ValueUs: nowMicros, Precision: dpb.DateTime_SECOND},
								End:   &dpb.DateTime{ValueUs: evenLaterMicros, Precision: dpb.DateTime_SECOND},
							},
							StatusHistory: []*encounterpb.Encounter_StatusHistory{{
								Status: &encounterpb.Encounter_StatusHistory_StatusCode{Value: cpb.EncounterStatusCode_PLANNED},
								Period: &dpb.Period{
									Start: &dpb.DateTime{ValueUs: nowMicros, Precision: dpb.DateTime_SECOND},
									End:   &dpb.DateTime{ValueUs: nowMicros, Precision: dpb.DateTime_SECOND},
								},
							}, {
								Status: &encounterpb.Encounter_StatusHistory_StatusCode{Value: cpb.EncounterStatusCode_ARRIVED},
								Period: &dpb.Period{
									Start: &dpb.DateTime{ValueUs: nowMicros, Precision: dpb.DateTime_SECOND},
									End:   &dpb.DateTime{ValueUs: laterMicros, Precision: dpb.DateTime_SECOND},
								},
							}, {
								Status: &encounterpb.Encounter_StatusHistory_StatusCode{Value: cpb.EncounterStatusCode_IN_PROGRESS},
								Period: &dpb.Period{
									Start: &dpb.DateTime{ValueUs: laterMicros, Precision: dpb.DateTime_SECOND},
									End:   &dpb.DateTime{ValueUs: evenLaterMicros, Precision: dpb.DateTime_SECOND},
								},
							}},
							Location: []*encounterpb.Encounter_Location{{
								Location: &dpb.Reference{
									Reference: &dpb.Reference_LocationId{
										&dpb.ReferenceId{Value: "5"},
									},
									Display: &dpb.String{Value: "BED, POC, ROOM, FLOOR, BUILDING, FACILITY"},
								},
								Period: &dpb.Period{
									Start: &dpb.DateTime{ValueUs: nowMicros, Precision: dpb.DateTime_SECOND},
									End:   &dpb.DateTime{ValueUs: laterMicros, Precision: dpb.DateTime_SECOND},
								},
							}, {
								Location: &dpb.Reference{
									Reference: &dpb.Reference_LocationId{
										&dpb.ReferenceId{Value: "5"},
									},
									Display: &dpb.String{Value: "BED, POC, ROOM, FLOOR, BUILDING, FACILITY"},
								},
								Period: &dpb.Period{
									Start: &dpb.DateTime{ValueUs: laterMicros, Precision: dpb.DateTime_SECOND},
									End:   &dpb.DateTime{ValueUs: laterMicros, Precision: dpb.DateTime_SECOND},
								},
							}, {
								Location: &dpb.Reference{
									Reference: &dpb.Reference_LocationId{
										&dpb.ReferenceId{Value: "6"},
									},
									Display: &dpb.String{Value: "BUILDING"},
								},
								Period: &dpb.Period{
									Start: &dpb.DateTime{ValueUs: laterMicros, Precision: dpb.DateTime_SECOND},
									End:   &dpb.DateTime{ValueUs: evenLaterMicros, Precision: dpb.DateTime_SECOND},
								},
							}},
							Diagnosis: []*encounterpb.Encounter_Diagnosis{{
								Condition: &dpb.Reference{
									Reference: &dpb.Reference_ProcedureId{
										&dpb.ReferenceId{Value: "8"},
									},
									Display: &dpb.String{Value: "PROCEDURE by Dr Doctor Doctorson"},
								},
							}, {
								Condition: &dpb.Reference{
									Reference: &dpb.Reference_ProcedureId{
										&dpb.ReferenceId{Value: "9"},
									},
									Display: &dpb.String{Value: "PROCEDURE by Dr Doctor Doctorson"},
								},
							}, {
								Condition: &dpb.Reference{
									Reference: &dpb.Reference_ConditionId{
										&dpb.ReferenceId{Value: "10"},
									},
									Display: &dpb.String{Value: "DIAGNOSIS by Dr Doctor Doctorson"},
								},
							}},
						},
					},
				},
			}, {
				FullUrl: &dpb.Uri{Value: "Observation/11"},
				Request: &r4pb.Bundle_Entry_Request{
					Method: &r4pb.Bundle_Entry_Request_MethodCode{Value: cpb.HTTPVerbCode_POST},
					Url:    &dpb.Uri{Value: "Observation"},
				},
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Observation{
						&observationpb.Observation{
							Id: &dpb.Id{Value: "11"},
							Code: &dpb.CodeableConcept{
								Coding: []*dpb.Coding{{
									Code:    &dpb.Code{Value: "TEST_ID_1"},
									System:  &dpb.Uri{Value: "SYSTEM_URI"},
									Display: &dpb.String{Value: "TEST_NAME_1"},
								}},
							},
							Encounter: &dpb.Reference{
								Reference: &dpb.Reference_EncounterId{&dpb.ReferenceId{Value: "4"}},
							},
							Text: &dpb.Narrative{
								Div: &dpb.Xhtml{
									Value: "<div><p>TEST_NAME_1: VALUE UNIT (H)</p><p>NOTE_1; NOTE_2</p></div>",
								},
								Status: &dpb.Narrative_StatusCode{Value: cpb.NarrativeStatusCode_GENERATED},
							},
							Status: &observationpb.Observation_StatusCode{Value: cpb.ObservationStatusCode_AMENDED},
							Subject: &dpb.Reference{
								Reference: &dpb.Reference_PatientId{
									&dpb.ReferenceId{Value: "1"},
								},
								Display: &dpb.String{Value: "William Burr"},
							},
							Value: &observationpb.Observation_ValueX{
								Choice: &observationpb.Observation_ValueX_Quantity{
									&dpb.Quantity{Value: &dpb.Decimal{Value: "VALUE"}, Unit: &dpb.String{Value: "UNIT"}},
								},
							},
							Effective: &observationpb.Observation_EffectiveX{
								Choice: &observationpb.Observation_EffectiveX_DateTime{
									&dpb.DateTime{ValueUs: laterMicros, Precision: dpb.DateTime_SECOND},
								},
							},
							Note: []*dpb.Annotation{{
								Text: &dpb.Markdown{Value: "NOTE_1"},
							}, {
								Text: &dpb.Markdown{Value: "NOTE_2"},
							}},
						},
					},
				},
			}, {
				FullUrl: &dpb.Uri{Value: "Observation/12"},
				Request: &r4pb.Bundle_Entry_Request{
					Method: &r4pb.Bundle_Entry_Request_MethodCode{Value: cpb.HTTPVerbCode_POST},
					Url:    &dpb.Uri{Value: "Observation"},
				},
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Observation{
						&observationpb.Observation{
							Id: &dpb.Id{Value: "12"},
							Code: &dpb.CodeableConcept{
								Coding: []*dpb.Coding{{
									Code:    &dpb.Code{Value: "TEST_ID_2"},
									System:  &dpb.Uri{Value: "SYSTEM_URI"},
									Display: &dpb.String{Value: "TEST_NAME_2"},
								}},
							},
							Status: &observationpb.Observation_StatusCode{Value: cpb.ObservationStatusCode_FINAL},
							Encounter: &dpb.Reference{
								Reference: &dpb.Reference_EncounterId{&dpb.ReferenceId{Value: "4"}},
							},
							Text: &dpb.Narrative{
								Div: &dpb.Xhtml{
									Value: "<div><p>TEST_NAME_2: VALUE UNIT</p></div>",
								},
								Status: &dpb.Narrative_StatusCode{Value: cpb.NarrativeStatusCode_GENERATED},
							},
							Subject: &dpb.Reference{
								Reference: &dpb.Reference_PatientId{
									&dpb.ReferenceId{Value: "1"},
								},
								Display: &dpb.String{Value: "William Burr"},
							},
							Value: &observationpb.Observation_ValueX{
								Choice: &observationpb.Observation_ValueX_Quantity{
									&dpb.Quantity{Value: &dpb.Decimal{Value: "VALUE"}, Unit: &dpb.String{Value: "UNIT"}},
								},
							},
							Effective: &observationpb.Observation_EffectiveX{
								Choice: &observationpb.Observation_EffectiveX_DateTime{
									&dpb.DateTime{ValueUs: laterMicros, Precision: dpb.DateTime_SECOND},
								},
							},
						},
					},
				},
			}, {
				FullUrl: &dpb.Uri{Value: "Encounter/13"},
				Request: &r4pb.Bundle_Entry_Request{
					Method: &r4pb.Bundle_Entry_Request_MethodCode{Value: cpb.HTTPVerbCode_POST},
					Url:    &dpb.Uri{Value: "Encounter"},
				},
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Encounter{
						&encounterpb.Encounter{
							Id: &dpb.Id{Value: "13"},
							ClassValue: &dpb.Coding{
								Code: &dpb.Code{Value: "IMP"},
							},
							Status: &encounterpb.Encounter_StatusCode{Value: cpb.EncounterStatusCode_IN_PROGRESS},
							Period: &dpb.Period{
								Start: &dpb.DateTime{ValueUs: evenLaterMicros, Precision: dpb.DateTime_SECOND},
							},
							Text: &dpb.Narrative{
								Div: &dpb.Xhtml{
									Value: "<div><p>Status: in-progress</p><p>Active from Mon Feb 12 10:00:00 2018</p></div>",
								},
								Status: &dpb.Narrative_StatusCode{Value: cpb.NarrativeStatusCode_GENERATED},
							},
						},
					},
				},
			}},
		},
	}, {
		name:       "Patient with missing fields",
		bundleType: Collection,
		patientInfo: &ir.PatientInfo{
			Person: &ir.Person{
				MRN:       "8888",
				FirstName: "Elisa",
				Surname:   "Mogollon",
				Address: &ir.Address{
					FirstLine:  "FIRST_LINE",
					City:       "CITY",
					Country:    "COUNTRY",
					PostalCode: "ABC DEF",
					Type:       "UNKNOWN",
				},
			}},
		want: &r4pb.Bundle{
			Type: &r4pb.Bundle_TypeCode{Value: cpb.BundleTypeCode_COLLECTION},
			Entry: []*r4pb.Bundle_Entry{{
				FullUrl: &dpb.Uri{Value: "Patient/1"},
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Patient{
						&patientpb.Patient{
							Id:         &dpb.Id{Value: "1"},
							Identifier: []*dpb.Identifier{{Value: &dpb.String{Value: "8888"}}},
							Text: &dpb.Narrative{
								Div:    &dpb.Xhtml{Value: "<div><p>Elisa Mogollon</p></div>"},
								Status: &dpb.Narrative_StatusCode{Value: cpb.NarrativeStatusCode_GENERATED},
							},
							Name: []*dpb.HumanName{{
								Family: &dpb.String{Value: "Mogollon"},
								Given:  []*dpb.String{{Value: "Elisa"}},
							}},
							Gender: &patientpb.Patient_GenderCode{Value: cpb.AdministrativeGenderCode_UNKNOWN},
							Address: []*dpb.Address{{
								Line:       []*dpb.String{{Value: "FIRST_LINE"}},
								City:       &dpb.String{Value: "CITY"},
								Country:    &dpb.String{Value: "COUNTRY"},
								PostalCode: &dpb.String{Value: "ABC DEF"},
								Type:       &dpb.Address_TypeCode{Value: cpb.AddressTypeCode_BOTH},
								Use:        &dpb.Address_UseCode{Value: cpb.AddressUseCode_INVALID_UNINITIALIZED},
							}},
							Deceased: &patientpb.Patient_DeceasedX{
								Choice: &patientpb.Patient_DeceasedX_Boolean{
									Boolean: &dpb.Boolean{
										Value: false,
									},
								},
							},
						},
					},
				},
			}},
		},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var b bytes.Buffer
			cfg := GeneratorConfig{
				HL7Config: &config.HL7Config{
					Gender: config.Gender{Male: "M", Female: "F"},
					ResultStatus: config.ResultStatus{
						Final:     "F",
						Corrected: "C",
					},
					Allergy: config.HL7Allergy{
						Types:      []string{"FOOD", "MEDICATION"},
						Severities: []string{"MILD", "MODERATE", "SEVERE"},
					},
					Mapping: config.CodeMapping{
						FHIR: config.FHIRMapping{
							CodingSystems: map[string]string{"SYSTEM": "SYSTEM_URI"},
							AllergySeverities: map[string][]string{
								"Severe":   {"SEVERE"},
								"Moderate": {"MODERATE"},
								"Mild":     {"MILD"},
							},
							AllergyTypes: map[string][]string{
								"Food":       {"FOOD"},
								"Medication": {"MEDICATION"},
							},
						},
					},
				},
				IDGenerator: &testid.Generator{},
				Output:      &testresource.ByteOutput{Bytes: &b},
				Marshaller:  prototext.MarshalOptions{},
				BundleType:  tc.bundleType,
			}

			w, err := NewFHIRWriter(cfg)
			if err != nil {
				t.Fatalf("NewFHIRWriter(%v) failed with: %v", cfg, err)
			}
			if err := w.Generate(tc.patientInfo); err != nil {
				t.Fatalf("w.Generate(%v) failed with: %v", tc.patientInfo, err)
			}
			if err := w.Close(); err != nil {
				t.Errorf("w.Close() failed with: %v", err)
			}

			got := &r4pb.Bundle{}
			if err := prototext.Unmarshal(b.Bytes(), got); err != nil {
				t.Fatalf("prototext.Unmarshal(%v, %v) failed with: %v", b.String(), got, err)
			}

			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("w.Generate(%v) returned diff (-want +got):\n%s", tc.patientInfo, diff)
			}
		})
	}
}
