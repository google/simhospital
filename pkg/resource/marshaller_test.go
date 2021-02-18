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
	"encoding/json"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	cpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/codes_go_proto"
	dpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
	r4pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/bundle_and_contained_resource_go_proto"
	patientpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/patient_go_proto"
)

type caseInsensitiveString string

// bundle defines the JSON schema to unmarshal into. We use this instead of the original
// r4pb.Bundle because json.Unmarshal does not recognise the datatypes in the proto.
type bundle struct {
	Entry        []entry
	ResourceType caseInsensitiveString
	Type         caseInsensitiveString
}

type entry struct {
	Resource resource
}

type resource struct {
	ID               string
	Identifier       []identifier
	Name             []name
	DeceasedDateTime string
	Gender           caseInsensitiveString
	ResourceType     caseInsensitiveString
	Telecom          []telecom
}

type identifier struct {
	Value string
}

type name struct {
	Family string
	Given  []string
	Suffix []string
}

type telecom struct {
	System caseInsensitiveString
	Use    caseInsensitiveString
	Value  string
}

func TestJSONMarshaller(t *testing.T) {
	b := testBundle()
	m, err := NewJSONMarshaller()
	if err != nil {
		t.Fatalf("NewJSONMarshaller() failed with %v", err)
	}

	want := bundle{
		Entry: []entry{{
			Resource: resource{
				ID: "1",
				Identifier: []identifier{{
					Value: "1234",
				}},
				Name: []name{{
					Family: "Burr",
					Given:  []string{"William", "George"},
					Suffix: []string{"MD"},
				}},
				DeceasedDateTime: "2018-02-12T00:00:00+00:00",
				Gender:           "male",
				Telecom: []telecom{{
					System: "phone",
					Use:    "home",
					Value:  "01234567890",
				}},
				ResourceType: "patient",
			},
		}},
		ResourceType: "Bundle",
		Type:         "collection",
	}

	bytes, err := m.Marshal(b)
	if err != nil {
		t.Fatalf("%T.Marshal(%s) failed with %v", m, b, err)
	}

	// We parse the JSON and ensure that it is semantically equivalent to what we expect. This
	// is important since properties can be in any order, and parsing creates a stable result.
	var got bundle
	if err := json.Unmarshal(bytes, &got); err != nil {
		t.Fatalf("json.Unmarshal(%v, %v) failed with %v", bytes, &got, err)
	}

	// We use case insensitive comparison for fields which contain case insensitive codes.
	trans := cmp.Transformer("", func(in caseInsensitiveString) string {
		return strings.ToLower(string(in))
	})

	if diff := cmp.Diff(want, got, trans); diff != "" {
		t.Errorf("json.Unmarshal(%v, %v) returned diff (-want +got):\n%s", bytes, &got, diff)
	}
}

func testBundle() *r4pb.Bundle {
	return &r4pb.Bundle{
		Type: &r4pb.Bundle_TypeCode{Value: cpb.BundleTypeCode_COLLECTION},
		Entry: []*r4pb.Bundle_Entry{{
			Resource: &r4pb.ContainedResource{
				OneofResource: &r4pb.ContainedResource_Patient{
					&patientpb.Patient{
						Id:         &dpb.Id{Value: "1"},
						Identifier: []*dpb.Identifier{{Value: &dpb.String{Value: "1234"}}},
						Name: []*dpb.HumanName{{
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
					},
				},
			},
		}},
	}
}
