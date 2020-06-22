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

	cpb "google/fhir/proto/r4/core/codes_go_proto"
	dpb "google/fhir/proto/r4/core/datatypes_go_proto"
	r4pb "google/fhir/proto/r4/core/resources/bundle_and_contained_resource_go_proto"
	patientpb "google/fhir/proto/r4/core/resources/patient_go_proto"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/testing/protocmp"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/ir"
)

var now = ir.NewValidTime(time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC))

func TestGeneratePatient(t *testing.T) {
	tests := []struct {
		name        string
		patientInfo *ir.PatientInfo
		want        *patientpb.Patient
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
			}},
		want: &patientpb.Patient{
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
		want: &patientpb.Patient{
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
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var b bytes.Buffer
			cfg := GeneratorConfig{
				HL7Config: &config.HL7Config{Gender: config.Gender{Male: "M", Female: "F"}},
				Writer:    &b,
			}

			w := NewFHIRWriter(cfg)
			w.Generate(tc.patientInfo)
			if err := w.Close(); err != nil {
				t.Errorf("w.Close() failed: %v", err)
			}

			bundle := &r4pb.Bundle{}
			if err := prototext.Unmarshal(b.Bytes(), bundle); err != nil {
				t.Errorf("prototext.Unmarshal(%v, %v) failed: %v", b.String(), bundle, err)
			}

			if got, want := len(bundle.GetEntry()), 1; got != want {
				t.Fatalf("len(bundle.GetEntry()) = %d, want %d", got, want)
			}

			got := bundle.GetEntry()[0].GetResource().GetPatient()
			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("w.Generate(%v) returned diff (-want +got):\n%s", tc.patientInfo, diff)
			}
		})
	}
}
