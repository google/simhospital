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

	dpb "google/fhir/proto/r4/core/datatypes_go_proto"
	r4pb "google/fhir/proto/r4/core/resources/bundle_and_contained_resource_go_proto"
	patientpb "google/fhir/proto/r4/core/resources/patient_go_proto"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/testing/protocmp"
	"github.com/google/simhospital/pkg/ir"
)

func TestGeneratePatient(t *testing.T) {
	tests := []struct {
		name        string
		patientInfo *ir.PatientInfo
		want        *patientpb.Patient
	}{{
		name: "Patient with prefix and suffix",
		patientInfo: &ir.PatientInfo{
			Person: &ir.Person{
				MRN:        "1234",
				Prefix:     "Dr",
				FirstName:  "William",
				MiddleName: "George",
				Surname:    "Burr",
				Suffix:     "MD",
			}},
		want: &patientpb.Patient{
			Identifier: []*dpb.Identifier{{Value: &dpb.String{Value: "1234"}}},
			Name: []*dpb.HumanName{{
				Prefix: []*dpb.String{{Value: "Dr"}},
				Family: &dpb.String{Value: "Burr"},
				Given:  []*dpb.String{{Value: "William"}, {Value: "George"}},
				Suffix: []*dpb.String{{Value: "MD"}},
			}},
		},
	}, {
		name: "Patient with firstname and surname",
		patientInfo: &ir.PatientInfo{
			Person: &ir.Person{
				MRN:       "8888",
				FirstName: "Elisa",
				Surname:   "Mogollon",
			}},
		want: &patientpb.Patient{
			Identifier: []*dpb.Identifier{{Value: &dpb.String{Value: "8888"}}},
			Name: []*dpb.HumanName{{
				Family: &dpb.String{Value: "Mogollon"},
				Given:  []*dpb.String{{Value: "Elisa"}},
			}},
		},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var b bytes.Buffer
			cfg := GeneratorConfig{
				Writer: &b,
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
