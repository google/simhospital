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
	"testing"

	cpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/codes_go_proto"
	pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
	observationpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/observation_go_proto"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	fhir "github.com/google/simhospital/pkg/fhircore"
)

func TestToQuantity(t *testing.T) {
	testCases := []struct {
		in   string
		want *pb.Quantity
	}{{
		in: "<15",
		want: &pb.Quantity{
			Value: &pb.Decimal{Value: "15"},
			Comparator: &pb.Quantity_ComparatorCode{
				Value:     cpb.QuantityComparatorCode_LESS_THAN,
				Extension: comparatorExtension("<"),
			},
		},
	}, {
		in: "[<15]",
		want: &pb.Quantity{
			Value: &pb.Decimal{Value: "15"},
			Comparator: &pb.Quantity_ComparatorCode{
				Value:     cpb.QuantityComparatorCode_LESS_THAN,
				Extension: comparatorExtension("<"),
			},
		},
	}, {
		in: ">1.50",
		want: &pb.Quantity{
			Value: &pb.Decimal{Value: "1.50"},
			Comparator: &pb.Quantity_ComparatorCode{
				Value:     cpb.QuantityComparatorCode_GREATER_THAN,
				Extension: comparatorExtension(">"),
			},
		},
	}, {
		in: ">=1",
		want: &pb.Quantity{
			Value: &pb.Decimal{Value: "1"},
			Comparator: &pb.Quantity_ComparatorCode{
				Value:     cpb.QuantityComparatorCode_GREATER_THAN_OR_EQUAL_TO,
				Extension: comparatorExtension(">="),
			},
		},
	}, {
		in: "<=-100.0",
		want: &pb.Quantity{
			Value: &pb.Decimal{Value: "-100.0"},
			Comparator: &pb.Quantity_ComparatorCode{
				Value:     cpb.QuantityComparatorCode_LESS_THAN_OR_EQUAL_TO,
				Extension: comparatorExtension("<="),
			},
		},
	}, {
		in: ">  1",
		want: &pb.Quantity{
			Value: &pb.Decimal{Value: "1"},
			Comparator: &pb.Quantity_ComparatorCode{
				Value:     cpb.QuantityComparatorCode_GREATER_THAN,
				Extension: comparatorExtension(">"),
			},
		},
	}, {
		in: "<= -100.0",
		want: &pb.Quantity{
			Value: &pb.Decimal{Value: "-100.0"},
			Comparator: &pb.Quantity_ComparatorCode{
				Value:     cpb.QuantityComparatorCode_LESS_THAN_OR_EQUAL_TO,
				Extension: comparatorExtension("<="),
			},
		},
	}, {
		in: "-100.0",
		want: &pb.Quantity{
			Value: &pb.Decimal{Value: "-100.0"},
		},
	}, {
		in: "2000.10",
		want: &pb.Quantity{
			Value: &pb.Decimal{Value: "2000.10"},
		},
	}}
	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			got, err := ToQuantity(context.Background(), tc.in)
			if err != nil {
				t.Fatalf("ToQuantity(%q) failed with %v", tc.in, err)
			}

			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("ToQuantity(%q) returned diff (-want +got):\n%s", tc.in, diff)
			}
		})
	}
}

func TestToReferenceRange(t *testing.T) {
	tests := []struct {
		in   string
		want *observationpb.Observation_ReferenceRange
	}{{
		"     1    -    2    ",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("     1    -    2    "),
			Low:  &pb.SimpleQuantity{Value: &pb.Decimal{Value: "1"}},
			High: &pb.SimpleQuantity{Value: &pb.Decimal{Value: "2"}},
		},
	}, {
		"-1    -    2",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("-1    -    2"),
			Low:  &pb.SimpleQuantity{Value: &pb.Decimal{Value: "-1"}},
			High: &pb.SimpleQuantity{Value: &pb.Decimal{Value: "2"}},
		},
	}, {
		"-1 - -2",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("-1 - -2"),
			Low:  &pb.SimpleQuantity{Value: &pb.Decimal{Value: "-1"}},
			High: &pb.SimpleQuantity{Value: &pb.Decimal{Value: "-2"}},
		},
	}, {
		"-1--2",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("-1--2"),
			Low:  &pb.SimpleQuantity{Value: &pb.Decimal{Value: "-1"}},
			High: &pb.SimpleQuantity{Value: &pb.Decimal{Value: "-2"}},
		},
	}, {
		"1--2",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("1--2"),
			Low:  &pb.SimpleQuantity{Value: &pb.Decimal{Value: "1"}},
			High: &pb.SimpleQuantity{Value: &pb.Decimal{Value: "-2"}},
		},
	}, {
		"1-2",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("1-2"),
			Low:  &pb.SimpleQuantity{Value: &pb.Decimal{Value: "1"}},
			High: &pb.SimpleQuantity{Value: &pb.Decimal{Value: "2"}},
		},
	}, {
		"+1-+2",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("+1-+2"),
			Low:  &pb.SimpleQuantity{Value: &pb.Decimal{Value: "+1"}},
			High: &pb.SimpleQuantity{Value: &pb.Decimal{Value: "+2"}},
		},
	}, {
		"1-2^1^2",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("1-2^1^2"),
			Low:  &pb.SimpleQuantity{Value: &pb.Decimal{Value: "1"}},
			High: &pb.SimpleQuantity{Value: &pb.Decimal{Value: "2"}},
		},
	}, {
		"1--2^1^-2",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("1--2^1^-2"),
			Low:  &pb.SimpleQuantity{Value: &pb.Decimal{Value: "1"}},
			High: &pb.SimpleQuantity{Value: &pb.Decimal{Value: "-2"}},
		},
	}, {
		"1.1-2.2",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("1.1-2.2"),
			Low:  &pb.SimpleQuantity{Value: &pb.Decimal{Value: "1.1"}},
			High: &pb.SimpleQuantity{Value: &pb.Decimal{Value: "2.2"}},
		},
	}, {
		"+1-+2^+1^+2",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("+1-+2^+1^+2"),
			Low:  &pb.SimpleQuantity{Value: &pb.Decimal{Value: "+1"}},
			High: &pb.SimpleQuantity{Value: &pb.Decimal{Value: "+2"}},
		},
	}, {
		"100-1000^100^1000",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("100-1000^100^1000"),
			Low:  &pb.SimpleQuantity{Value: &pb.Decimal{Value: "100"}},
			High: &pb.SimpleQuantity{Value: &pb.Decimal{Value: "1000"}},
		},
	}, {
		"[1-2]",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("[1-2]"),
			Low:  &pb.SimpleQuantity{Value: &pb.Decimal{Value: "1"}},
			High: &pb.SimpleQuantity{Value: &pb.Decimal{Value: "2"}},
		},
	}, {
		" [ 1 - 2 ] ",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String(" [ 1 - 2 ] "),
			Low:  &pb.SimpleQuantity{Value: &pb.Decimal{Value: "1"}},
			High: &pb.SimpleQuantity{Value: &pb.Decimal{Value: "2"}},
		},
	}, {
		"<1",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("<1"),
			High: &pb.SimpleQuantity{
				Value:     &pb.Decimal{Value: "1"},
				Extension: comparatorExtension("<"),
			},
		},
	}, {
		" <=1 ",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String(" <=1 "),
			High: &pb.SimpleQuantity{
				Value:     &pb.Decimal{Value: "1"},
				Extension: comparatorExtension("<="),
			},
		},
	}, {
		" >= 2 ",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String(" >= 2 "),
			Low: &pb.SimpleQuantity{
				Value:     &pb.Decimal{Value: "2"},
				Extension: comparatorExtension(">="),
			},
		},
	}, {
		" >= 2.5 ^>= 2.5 ",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String(" >= 2.5 ^>= 2.5 "),
			Low: &pb.SimpleQuantity{
				Value:     &pb.Decimal{Value: "2.5"},
				Extension: comparatorExtension(">="),
			},
		},
	}, {
		" >= 2.5 ^^>= 2.5 ",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String(" >= 2.5 ^^>= 2.5 "),
			Low: &pb.SimpleQuantity{
				Value:     &pb.Decimal{Value: "2.5"},
				Extension: comparatorExtension(">="),
			},
		},
	}, {
		"-",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("-"),
		},
	}, {
		"",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String(""),
		},
	}, {
		"   ",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("   "),
		},
	}, {
		"[   ]",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("[   ]"),
		},
	}, {
		"[]",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("[]"),
		},
	}, {
		">=2.5^^<=2.5",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String(">=2.5^^<=2.5"),
		},
	}, {
		"1-2^10^2",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("1-2^10^2"),
		},
	}, {
		"1-2^1^3",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("1-2^1^3"),
		},
	}, {
		"NonNumber-String",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("NonNumber-String"),
		},
	}, {
		"NonNumber-String^NonNumber^String",
		&observationpb.Observation_ReferenceRange{
			Text: fhir.String("NonNumber-String^NonNumber^String"),
		},
	}}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			got := ToReferenceRange(context.Background(), tc.in)

			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("ToReferenceRange(%q) returned diff (-want +got):\n%s", tc.in, diff)
			}
		})
	}
}

func comparatorExtension(in string) []*pb.Extension {
	return []*pb.Extension{fhir.StringExtension(ComparatorExtensionURI, in)}
}
