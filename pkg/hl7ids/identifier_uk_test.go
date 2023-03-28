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

package hl7ids

import (
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/hl7"
	"github.com/google/simhospital/pkg/test/testhl7"
)

func TestMain(m *testing.M) {
	hl7.TimezoneAndLocation("Europe/London")
	retCode := m.Run()
	os.Exit(retCode)
}

func TestGetMRNNumber(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{{
		name: "Find valid MRN in PID-3-1",
		in:   "PID|1||2222^^^^MRN~5576027981^^^NHSNBR||",
		want: "2222",
	}, {
		name: "MRN keyword included in longer string is found",
		in:   "PID|1||2222^^^^SFA MRN~5576027981^^^NHSNBR||",
		want: "2222",
	}, {
		name: "Find MRN in PID-3-1 after NHS",
		in:   "PID|1||5576027981^^^NHSNBR~2222^^^^MRN||",
		want: "2222",
	}, {
		name: "PID-3-1 is preferred over PID-4",
		in:   "PID|1||2222^^^^MRN~5576027981^^^NHSNBR|1111^^^^MRN^|",
		want: "2222",
	}, {
		name: "Empty MRN in PID-3-1 finds PID-4",
		in:   "PID|1||^^^^MRN~5576027981^^^NHSNBR|1111^^^^MRN^||",
		want: "1111",
	}, {
		name: "PID-2 is ignored",
		in:   "PID|1|1111^^^^MRN^|^^^^MRN~5576027981^^^NHSNBR|",
		want: "",
	}, {
		name: "An id without MRN keyword is not found",
		in:   "PID|1||2222^^^SFA^~5576027981^^^NHSNBR||",
		want: "",
	}, {
		name: "Multiple MRNs in PID-3-1 finds the first one",
		in:   "PID|1||2222^^^^MRN~1111^^^^MRN^~5576027981^^^NHSNBR||",
		want: "2222",
	}, {
		name: "All fields empty finds nothing",
		in:   "PID|1||||",
		want: "",
	}, {
		name: "MRN keyword cannot be in CX-2",
		in:   "PID|1||2222^MRN||",
		want: "",
	}, {
		name: "MRN keyword can be in CX-3",
		in:   "PID|1||2222^^MRN||",
		want: "2222",
	}, {
		name: "MRN keyword can be in CX-4",
		in:   "PID|1||2222^^^MRN||",
		want: "2222",
	}, {
		name: "MRN keyword can be in CX-5",
		in:   "PID|1||2222^^^^MRN||",
		want: "2222",
	}, {
		name: "MRN keyword can be in CX-6",
		in:   "PID|1||2222^^^^^MRN||",
		want: "2222",
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := testhl7.Parse(t, strings.Join([]string{testhl7.SegmentMSH, test.in}, "\r"))
			pid, err := m.PID()
			if err != nil {
				t.Fatalf("PID() failed with %v", err)
			}
			got := GetMRNNumber(pid)
			if got != test.want {
				t.Errorf("GetMRNNumber(%s) got %q, want %q", test.in, got, test.want)
			}
		})
	}
}

func TestGetMRNNumberWithOptions(t *testing.T) {
	tests := []struct {
		name string
		in   string
		opts *Options
		want string
	}{{
		name: "PID-4 is preferred over PID-3 if specified in the options",
		opts: &Options{
			IsValidMRN: DefaultOptions().IsValidMRN,
			MRNFrom:    []int{4, 3},
		},
		in:   "PID|1||2222^^^^MRN~5576027981^^^NHSNBR|1111^^^^MRN^|",
		want: "1111",
	}, {
		name: "PID-2 is preferred if specified in the options",
		opts: &Options{
			IsValidMRN: DefaultOptions().IsValidMRN,
			MRNFrom:    []int{2, 3, 4},
		},
		in:   "PID|1|1111^^^^MRN^|^^^^MRN~5576027981^^^NHSNBR|",
		want: "1111",
	}, {
		name: "We can specify a different MRN keyword",
		opts: &Options{
			IsValidMRN: func(c hl7.CX) bool { return HasKeywordInAnyField(c, "SFA") },
			MRNFrom:    DefaultOptions().MRNFrom,
		},
		in:   "PID|1||2222^^^SFA^~5576027981^^^NHSNBR||",
		want: "2222",
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := testhl7.Parse(t, strings.Join([]string{testhl7.SegmentMSH, test.in}, "\r"))
			pid, err := m.PID()
			if err != nil {
				t.Fatalf("PID() failed with %v", err)
			}
			got := GetMRNNumberWithOptions(pid, test.opts)
			if got != test.want {
				t.Errorf("GetMRNNumberWithOptions(%s, %+v) got %q, want %q", test.in, test.opts, got, test.want)
			}
		})
	}
}

func TestGetNHSNumber(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{{
		name: "Valid NHS in PID-3-1",
		in:   "PID|1||2222^^^^MRN~5576027981^^^NHSNBR|1111^^^^MRN^|",
		want: "5576027981",
	}, {
		name: "An invalid NHS is not found",
		in:   "PID|1||2222^^^^MRN~1122231236^^^NHSNBR|1111^^^^MRN^|",
		want: "",
	}, {
		name: "NHS in PID-3-1 is preferred over PID-4",
		in:   "PID|1||2222^^^^MRN~4010232137^^^NHSNBR|5576027981^^^NHSNBR|",
		want: "4010232137",
	}, {
		name: "Empty NHS in PID-3-1 finds PID-4",
		in:   "PID|1||^^^^MRN~5576027981^^^NHSNBR|5576027981^^^NHSNBR||",
		want: "5576027981",
	}, {
		name: "PID-2 is ignored",
		in:   "PID|1|5576027981^^^NHSNBR||",
		want: "",
	}, {
		name: "An id without NHS keyword is not found",
		in:   "PID|1||2222^^^SFA^~5576027981^^^NBR||",
		want: "",
	}, {
		name: "Multiple NHSs in PID-3-1 finds the first one",
		in:   "PID|1||2222^^^^MRN~4010232137^^^NHSNBR~5576027981^^^NHSNBR||",
		want: "4010232137",
	}, {
		name: "All fields empty finds nothing",
		in:   "PID|1||||",
		want: "",
	}, {
		name: "NHS keyword cannot be in CX-2",
		in:   "PID|1||5576027981^NHS||",
		want: "",
	}, {
		name: "NHS keyword can be in CX-3",
		in:   "PID|1||5576027981^^NHS||",
		want: "5576027981",
	}, {
		name: "NHS keyword can be in CX-4",
		in:   "PID|1||5576027981^^^NHS||",
		want: "5576027981",
	}, {
		name: "NHS keyword can be in CX-5",
		in:   "PID|1||5576027981^^^^NHS||",
		want: "5576027981",
	}, {
		name: "NHS keyword can be in CX-6",
		in:   "PID|1||5576027981^^^^^NHS||",
		want: "5576027981",
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := testhl7.Parse(t, strings.Join([]string{testhl7.SegmentMSH, test.in}, "\r"))
			pid, err := m.PID()
			if err != nil {
				t.Fatalf("PID() failed with %v", err)
			}
			got := GetNHSNumber(pid)
			if got != test.want {
				t.Errorf("GetNHSNumber(%s) got %q, want %q", test.in, got, test.want)
			}
		})
	}
}

func TestGetNHSNumberWithOptions(t *testing.T) {
	tests := []struct {
		name string
		in   string
		opts *Options
		want string
	}{{
		name: "PID-4 is preferred if specified in the options",
		in:   "PID|1||2222^^^^MRN~4010232137^^^NHSNBR|5576027981^^^NHSNBR|",
		opts: &Options{
			NHSFrom:    []int{4},
			IsValidNHS: DefaultOptions().IsValidNHS,
		},
		want: "5576027981",
	}, {
		name: "PID-2 is taken into account if specified in the options",
		in:   "PID|1|5576027981^^^NHSNBR||",
		opts: &Options{
			NHSFrom:    []int{2},
			IsValidNHS: DefaultOptions().IsValidNHS,
		},
		want: "5576027981",
	}, {
		name: "We can specify a different NHS keyword",
		opts: &Options{
			IsValidNHS: func(c hl7.CX) bool {
				return HasKeywordInAnyField(c, "OTHER_KEY") && NHSNumberIsValid(c.IDNumber.SanitizedString())
			},
			NHSFrom: DefaultOptions().NHSFrom,
		},
		in:   "PID|1||2222^^^SFA^~5576027981^^^OTHER_KEY||",
		want: "5576027981",
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := testhl7.Parse(t, strings.Join([]string{testhl7.SegmentMSH, test.in}, "\r"))
			pid, err := m.PID()
			if err != nil {
				t.Fatalf("PID() failed with %v", err)
			}
			got := GetNHSNumberWithOptions(pid, test.opts)
			if got != test.want {
				t.Errorf("GetNHSNumberWithOptions(%s, %v) got %q, want %q", test.in, test.opts, got, test.want)
			}
		})
	}
}

func TestGetAllMRNs(t *testing.T) {
	st1111 := hl7.ST("1111")
	st2222 := hl7.ST("2222")
	idMRN := hl7.ID("MRN")
	idMRNOther := hl7.ID("OTHER_MRN")
	want2222 := hl7.CX{IDNumber: &st2222, IdentifierTypeCode: &idMRN}
	want1111 := hl7.CX{IDNumber: &st1111, IdentifierTypeCode: &idMRN}
	want222Other := hl7.CX{IDNumber: &st2222, IdentifierTypeCode: &idMRNOther}

	tests := []struct {
		name string
		in   string
		want []hl7.CX
	}{{
		name: "Find valid MRN in PID-3-1",
		in:   "PID|1||2222^^^^MRN~5576027981^^^NHSNBR||",
		want: []hl7.CX{want2222},
	}, {
		name: "Concatenates multiple MRNs from the same field",
		in:   "PID|1||2222^^^^MRN~5576027981^^^NHSNBR~1111^^^^MRN^||",
		want: []hl7.CX{want2222, want1111},
	}, {
		name: "Concatenates multiple MRNs from different fields",
		in:   "PID|1||2222^^^^MRN~5576027981^^^NHSNBR|1111^^^^MRN^|",
		want: []hl7.CX{want2222, want1111},
	}, {
		name: "Dedupes from the same field",
		in:   "PID|1||2222^^^^MRN~5576027981^^^NHSNBR~2222^^^^MRN^||",
		want: []hl7.CX{want2222},
	}, {
		name: "Dedupes from different fields",
		in:   "PID|1||2222^^^^MRN~5576027981^^^NHSNBR|2222^^^^MRN^|",
		want: []hl7.CX{want2222},
	}, {
		name: "MRNs with different sub-fields are considered different",
		in:   "PID|1||2222^^^^MRN~5576027981^^^NHSNBR|2222^^^^OTHER_MRN^|",
		want: []hl7.CX{want2222, want222Other},
	}, {
		name: "PID-2 is ignored",
		in:   "PID|1|333^^MRN|2222^^^^MRN~5576027981^^^NHSNBR|1111^^^^MRN^|",
		want: []hl7.CX{want2222, want1111},
	}, {
		name: "All fields empty finds nothing",
		in:   "PID|1||||",
		want: []hl7.CX{},
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := testhl7.Parse(t, strings.Join([]string{testhl7.SegmentMSH, test.in}, "\r"))
			pid, err := m.PID()
			if err != nil {
				t.Fatalf("PID() failed with %v", err)
			}
			got := GetAllMRNs(pid)
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("GetAllMRNs(%s) returned diff (-want +got):\n%s", test.in, diff)
			}
		})
	}
}

func TestNHSNumberIsValid(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{{
		name: "valid",
		in:   "5576027981",
		want: true,
	}, {
		name: "invalid",
		in:   "1122231236",
		want: false,
	}, {
		name: "empty",
		in:   "",
		want: false,
	}, {
		name: "not a number",
		in:   "not a number",
		want: false,
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := NHSNumberIsValid(test.in); got != test.want {
				t.Errorf("NHSNumberIsValid(%s) got %t, want %t", test.in, got, test.want)
			}
		})
	}
}
