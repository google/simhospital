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

	cpb "google/fhir/proto/r4/core/codes_go_proto"
	dpb "google/fhir/proto/r4/core/datatypes_go_proto"
	r4pb "google/fhir/proto/r4/core/resources/bundle_and_contained_resource_go_proto"
	encounterpb "google/fhir/proto/r4/core/resources/encounter_go_proto"
	observationpb "google/fhir/proto/r4/core/resources/observation_go_proto"
	patientpb "google/fhir/proto/r4/core/resources/patient_go_proto"
)

var (
	delay     = time.Hour * 5
	now       = ir.NewValidTime(time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC))
	later     = ir.NewValidTime(now.Add(delay))
	evenLater = ir.NewValidTime(later.Add(delay))
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name        string
		patientInfo *ir.PatientInfo
		want        *r4pb.Bundle
	}{{
		name: "Patient with all fields",
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
				Orders: []*ir.Order{{
					OrderDateTime: later,
					Results: []*ir.Result{{
						TestName:     &ir.CodedElement{ID: "TEST_ID_1", Text: "TEST_NAME_1"},
						Value:        "VALUE",
						Unit:         "UNIT",
						AbnormalFlag: "H",
						Notes:        []string{"NOTE_1", "NOTE_2"},
						Status:       "C",
					}, {
						TestName: &ir.CodedElement{ID: "TEST_ID_2", Text: "TEST_NAME_2"},
						Value:    "VALUE",
						Unit:     "UNIT",
					}},
				}},
			}, {
				Status:      constants.EncounterStatusInProgress,
				StatusStart: evenLater,
				Start:       evenLater,
			}},
		},
		want: &r4pb.Bundle{
			Entry: []*r4pb.Bundle_Entry{{
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Patient{
						&patientpb.Patient{
							Id:         &dpb.Id{Value: "1"},
							Identifier: []*dpb.Identifier{{Value: &dpb.String{Value: "1234"}}},
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
									DateTime: &dpb.DateTime{
										ValueUs:   now.Unix(),
										Precision: dpb.DateTime_SECOND,
									},
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
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Encounter{
						&encounterpb.Encounter{
							Id:     &dpb.Id{Value: "2"},
							Status: &encounterpb.Encounter_StatusCode{Value: cpb.EncounterStatusCode_FINISHED},
							Period: &dpb.Period{
								Start: &dpb.DateTime{ValueUs: now.Unix(), Precision: dpb.DateTime_SECOND},
								End:   &dpb.DateTime{ValueUs: evenLater.Unix(), Precision: dpb.DateTime_SECOND},
							},
							StatusHistory: []*encounterpb.Encounter_StatusHistory{{
								Status: &encounterpb.Encounter_StatusHistory_StatusCode{Value: cpb.EncounterStatusCode_PLANNED},
								Period: &dpb.Period{
									Start: &dpb.DateTime{ValueUs: now.Unix(), Precision: dpb.DateTime_SECOND},
									End:   &dpb.DateTime{ValueUs: now.Unix(), Precision: dpb.DateTime_SECOND},
								},
							}, {
								Status: &encounterpb.Encounter_StatusHistory_StatusCode{Value: cpb.EncounterStatusCode_ARRIVED},
								Period: &dpb.Period{
									Start: &dpb.DateTime{ValueUs: now.Unix(), Precision: dpb.DateTime_SECOND},
									End:   &dpb.DateTime{ValueUs: later.Unix(), Precision: dpb.DateTime_SECOND},
								},
							}, {
								Status: &encounterpb.Encounter_StatusHistory_StatusCode{Value: cpb.EncounterStatusCode_IN_PROGRESS},
								Period: &dpb.Period{
									Start: &dpb.DateTime{ValueUs: later.Unix(), Precision: dpb.DateTime_SECOND},
									End:   &dpb.DateTime{ValueUs: evenLater.Unix(), Precision: dpb.DateTime_SECOND},
								},
							}},
						},
					},
				},
			}, {
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Observation{
						&observationpb.Observation{
							Id: &dpb.Id{Value: "3"},
							Encounter: &dpb.Reference{
								Reference: &dpb.Reference_EncounterId{&dpb.ReferenceId{Value: "2"}},
							},
							Text: &dpb.Narrative{
								Div: &dpb.Xhtml{
									Value: "<div><p>TEST_NAME_1: VALUE UNIT (H)</p><p>NOTE_1; NOTE_2</p></div>",
								},
							},
							Status: &observationpb.Observation_StatusCode{Value: cpb.ObservationStatusCode_AMENDED},
							Subject: &dpb.Reference{
								Reference: &dpb.Reference_PatientId{&dpb.ReferenceId{Value: "1"}},
								Display:   &dpb.String{Value: "William Burr"},
							},
							Value: &observationpb.Observation_ValueX{
								Choice: &observationpb.Observation_ValueX_Quantity{
									&dpb.Quantity{Value: &dpb.Decimal{Value: "VALUE"}, Unit: &dpb.String{Value: "UNIT"}},
								},
							},
							Effective: &observationpb.Observation_EffectiveX{
								Choice: &observationpb.Observation_EffectiveX_DateTime{
									&dpb.DateTime{
										ValueUs:   later.Unix(),
										Precision: dpb.DateTime_SECOND,
									},
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
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Observation{
						&observationpb.Observation{
							Id: &dpb.Id{Value: "4"},
							Encounter: &dpb.Reference{
								Reference: &dpb.Reference_EncounterId{&dpb.ReferenceId{Value: "2"}},
							},
							Text: &dpb.Narrative{
								Div: &dpb.Xhtml{
									Value: "<div><p>TEST_NAME_2: VALUE UNIT</p></div>",
								},
							},
							Status: &observationpb.Observation_StatusCode{Value: cpb.ObservationStatusCode_INVALID_UNINITIALIZED},
							Subject: &dpb.Reference{
								Reference: &dpb.Reference_PatientId{&dpb.ReferenceId{Value: "1"}},
								Display:   &dpb.String{Value: "William Burr"},
							},
							Value: &observationpb.Observation_ValueX{
								Choice: &observationpb.Observation_ValueX_Quantity{
									&dpb.Quantity{Value: &dpb.Decimal{Value: "VALUE"}, Unit: &dpb.String{Value: "UNIT"}},
								},
							},
							Effective: &observationpb.Observation_EffectiveX{
								Choice: &observationpb.Observation_EffectiveX_DateTime{
									&dpb.DateTime{
										ValueUs:   later.Unix(),
										Precision: dpb.DateTime_SECOND,
									},
								},
							},
						},
					},
				},
			}, {
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Encounter{
						&encounterpb.Encounter{
							Id:     &dpb.Id{Value: "5"},
							Status: &encounterpb.Encounter_StatusCode{Value: cpb.EncounterStatusCode_IN_PROGRESS},
							Period: &dpb.Period{
								Start: &dpb.DateTime{ValueUs: evenLater.Unix(), Precision: dpb.DateTime_SECOND},
							},
						},
					},
				},
			}},
		},
	}, {
		name: "Patient with missing fields",
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
			Entry: []*r4pb.Bundle_Entry{{
				Resource: &r4pb.ContainedResource{
					OneofResource: &r4pb.ContainedResource_Patient{
						&patientpb.Patient{
							Id:         &dpb.Id{Value: "1"},
							Identifier: []*dpb.Identifier{{Value: &dpb.String{Value: "8888"}}},
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
								Use:        &dpb.Address_UseCode{Value: cpb.AddressUseCode_INVALID_UNINITIALIZED},
								Type:       &dpb.Address_TypeCode{Value: cpb.AddressTypeCode_BOTH},
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
				},
				Writer:      &b,
				IDGenerator: &testid.Generator{},
			}

			w := NewFHIRWriter(cfg)
			if err := w.Generate(tc.patientInfo); err != nil {
				t.Fatalf("w.Generate(%v) failed: %v", tc.patientInfo, err)
			}
			if err := w.Close(); err != nil {
				t.Errorf("w.Close() failed: %v", err)
			}

			got := &r4pb.Bundle{}
			if err := prototext.Unmarshal(b.Bytes(), got); err != nil {
				t.Fatalf("prototext.Unmarshal(%v, %v) failed: %v", b.String(), got, err)
			}

			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("w.Generate(%v) returned diff (-want +got):\n%s", tc.patientInfo, diff)
			}
		})
	}
}
