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

package fhircore

import (
	"testing"

	pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

const targetCode = "target_url"

var (
	id1            = Identifier("1", targetCode)
	id2            = Identifier("2", targetCode)
	id3            = Identifier("3", targetCode)
	idNotMatching  = Identifier("4", "different_code")
	idNotMatching2 = Identifier("5", "different_code")
)

func TestAddOrUpdateIdentifier(t *testing.T) {
	tcs := []struct {
		name       string
		inExisting []*pb.Identifier
		inNew      *pb.Identifier
		want       []*pb.Identifier
	}{{
		name:       "add new",
		inExisting: []*pb.Identifier{},
		inNew:      id1,
		want:       []*pb.Identifier{id1},
	}, {
		name:       "update existing",
		inExisting: []*pb.Identifier{id1},
		inNew:      id2,
		want:       []*pb.Identifier{id2},
	}, {
		name:       "add even if there are other identifiers",
		inExisting: []*pb.Identifier{idNotMatching, id1},
		inNew:      id2,
		want:       []*pb.Identifier{idNotMatching, id2},
	}, {
		name:       "multiple matching identifiers updates the first one",
		inExisting: []*pb.Identifier{id1, id2},
		inNew:      id3,
		want:       []*pb.Identifier{id3, id2},
	}, {
		name:       "multiple matching identifiers updates the first one",
		inExisting: []*pb.Identifier{id1, id2},
		inNew:      id3,
		want:       []*pb.Identifier{id3, id2},
	}, {
		name:       "empty code does not match with other codes",
		inExisting: []*pb.Identifier{id1, id2},
		inNew:      Identifier("value", ""),
		want:       []*pb.Identifier{id1, id2, Identifier("value", "")},
	}, {
		name:       "can add first with empty code",
		inExisting: []*pb.Identifier{},
		inNew:      Identifier("value", ""),
		want:       []*pb.Identifier{Identifier("value", "")},
	}, {
		name:       "empty code always adds",
		inExisting: []*pb.Identifier{Identifier("value1", "")},
		inNew:      Identifier("value2", ""),
		want:       []*pb.Identifier{Identifier("value1", ""), Identifier("value2", "")},
	}, {
		name:       "secondary MRNs adds if different tc",
		inExisting: []*pb.Identifier{IdentifierSecondaryMRN("value1", "tc", "au", "fac")},
		inNew:      IdentifierSecondaryMRN("value2", "other-tc", "au", "fac"),
		want:       []*pb.Identifier{IdentifierSecondaryMRN("value1", "tc", "au", "fac"), IdentifierSecondaryMRN("value2", "other-tc", "au", "fac")},
	}, {
		name:       "secondary MRNs adds if different au",
		inExisting: []*pb.Identifier{IdentifierSecondaryMRN("value1", "tc", "au", "fac")},
		inNew:      IdentifierSecondaryMRN("value2", "tc", "other-au", "fac"),
		want:       []*pb.Identifier{IdentifierSecondaryMRN("value1", "tc", "au", "fac"), IdentifierSecondaryMRN("value2", "tc", "other-au", "fac")},
	}, {
		name:       "secondary MRNs adds if different fac",
		inExisting: []*pb.Identifier{IdentifierSecondaryMRN("value1", "tc", "au", "fac")},
		inNew:      IdentifierSecondaryMRN("value2", "tc", "au", "other-fac"),
		want:       []*pb.Identifier{IdentifierSecondaryMRN("value1", "tc", "au", "fac"), IdentifierSecondaryMRN("value2", "tc", "au", "other-fac")},
	}, {
		name:       "secondary MRNs updates if everything matches",
		inExisting: []*pb.Identifier{IdentifierSecondaryMRN("value1", "tc", "au", "fac")},
		inNew:      IdentifierSecondaryMRN("value1", "tc", "au", "fac"),
		want:       []*pb.Identifier{IdentifierSecondaryMRN("value1", "tc", "au", "fac")},
	}, {
		name:       "secondary MRNs adds if all fields match but value does not",
		inExisting: []*pb.Identifier{IdentifierSecondaryMRN("value1", "tc", "au", "fac")},
		inNew:      IdentifierSecondaryMRN("value2", "tc", "au", "fac"),
		want:       []*pb.Identifier{IdentifierSecondaryMRN("value1", "tc", "au", "fac"), IdentifierSecondaryMRN("value2", "tc", "au", "fac")},
	}}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := AddOrUpdateIdentifier(tc.inExisting, tc.inNew)

			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("AddOrUpdateIdentifier(%v, %v) mismatch (-want, +got)=\n%s", tc.inExisting, tc.inNew, diff)
			}
		})
	}
}

func TestGetMRN(t *testing.T) {
	tests := []struct {
		name string
		in   []*pb.Identifier
		want []string
	}{{
		name: "No MRN numbers",
		in:   []*pb.Identifier{Identifier("value", "some-code")},
		want: []string{},
	}, {
		name: "One MRN number",
		in:   []*pb.Identifier{IdentifierMRN("mrn1")},
		want: []string{"mrn1"},
	}, {
		name: "Multiple MRN numbers",
		in:   []*pb.Identifier{IdentifierMRN("mrn1"), IdentifierMRN("mrn2")},
		want: []string{"mrn1", "mrn2"},
	}, {
		name: "Multiple MRN numbers with other identifiers",
		in:   []*pb.Identifier{IdentifierMRN("mrn1"), Identifier("value", "code"), IdentifierMRN("mrn2")},
		want: []string{"mrn1", "mrn2"},
	}, {
		name: "Multiple MRN numbers with NHS only find MRN",
		in:   []*pb.Identifier{IdentifierMRN("mrn1"), IdentifierNHS("nhs1"), IdentifierMRN("mrn2")},
		want: []string{"mrn1", "mrn2"},
	}, {
		name: "MRN created with Identifier method",
		in:   []*pb.Identifier{Identifier("mrn1", MRNIdentifierCode)},
		want: []string{"mrn1"},
	}, {
		name: "Multiple identical MRN numbers",
		in:   []*pb.Identifier{IdentifierMRN("mrn1"), IdentifierMRN("mrn1")},
		want: []string{"mrn1", "mrn1"},
	}, {
		name: "Empty identifiers are ignored",
		in:   []*pb.Identifier{{}},
		want: []string{},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetMRN(tc.in)
			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("GetMRN(%v) mismatch (-want, +got)=\n%s", tc.in, diff)
			}
		})
	}
}

func TestGetNHS(t *testing.T) {
	tests := []struct {
		name string
		in   []*pb.Identifier
		want []string
	}{{
		name: "No NHS numbers",
		in:   []*pb.Identifier{Identifier("value", "some-code")},
		want: []string{},
	}, {
		name: "One NHS number",
		in:   []*pb.Identifier{IdentifierNHS("nhs1")},
		want: []string{"nhs1"},
	}, {
		name: "Multiple NHS numbers",
		in:   []*pb.Identifier{IdentifierNHS("nhs1"), IdentifierNHS("nhs2")},
		want: []string{"nhs1", "nhs2"},
	}, {
		name: "Multiple NHS numbers with other identifiers",
		in:   []*pb.Identifier{IdentifierNHS("nhs1"), Identifier("value", "code"), IdentifierNHS("nhs2")},
		want: []string{"nhs1", "nhs2"},
	}, {
		name: "Multiple NHS numbers with NHS only find NHS",
		in:   []*pb.Identifier{IdentifierNHS("nhs1"), IdentifierMRN("mrn1"), IdentifierNHS("nhs2")},
		want: []string{"nhs1", "nhs2"},
	}, {
		name: "NHS created with Identifier method",
		in:   []*pb.Identifier{Identifier("nhs1", NHSIdentifierCode)},
		want: []string{"nhs1"},
	}, {
		name: "Multiple identical NHS numbers",
		in:   []*pb.Identifier{IdentifierNHS("nhs1"), IdentifierNHS("nhs1")},
		want: []string{"nhs1", "nhs1"},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetNHS(tc.in)
			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("GetNHS(%v) mismatch (-want, +got)=\n%s", tc.in, diff)
			}
		})
	}
}

func TestFindIdentifiers(t *testing.T) {
	tests := []struct {
		name   string
		inIds  []*pb.Identifier
		inCode string
		want   []*pb.Identifier
	}{{
		name:   "Match code",
		inIds:  []*pb.Identifier{Identifier("value", "some-code")},
		inCode: "some-code",
		want:   []*pb.Identifier{Identifier("value", "some-code")},
	}, {
		name:   "Doesnt match code",
		inIds:  []*pb.Identifier{Identifier("value", "some-code")},
		inCode: "some-other-code",
		want:   []*pb.Identifier{},
	}, {
		name:   "Empty initial identifiers",
		inIds:  []*pb.Identifier{},
		inCode: "irrelevant",
		want:   []*pb.Identifier{},
	}, {
		name:   "Nil initial identifiers",
		inIds:  nil,
		inCode: "irrelevant",
		want:   []*pb.Identifier{},
	}, {
		name:   "Empty code can match",
		inIds:  []*pb.Identifier{Identifier("value1", ""), Identifier("value2", "")},
		inCode: "",
		want:   []*pb.Identifier{Identifier("value1", ""), Identifier("value2", "")},
	}, {
		name:   "Returns multiple matches",
		inIds:  []*pb.Identifier{Identifier("value1", "code"), Identifier("value2", "code"), IdentifierNHS("nhs1")},
		inCode: "code",
		want:   []*pb.Identifier{Identifier("value1", "code"), Identifier("value2", "code")},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FindIdentifiers(tc.inIds, tc.inCode)
			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("FindIdentifiers(%v, %v) mismatch (-want, +got)=\n%s", tc.inIds, tc.inCode, diff)
			}
		})
	}
}
