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

package hl7

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/text/encoding/unicode"
)

func TestPrimitive(t *testing.T) {
	TimezoneAndLocation("UTC")
	c := &Context{
		Decoder:     unicode.UTF8.NewDecoder(),
		Delimiters:  DefaultDelimiters,
		Nesting:     0,
		TimezoneLoc: Location,
	}

	cases := []struct {
		name string
		p    Primitive
		want Primitive
		// got is an empty placeholder of the given type,
		// where the value will be unmarshalled to.
		got Primitive
	}{{
		name: "ST",
		p:    NewST("value"),
		want: NewST("value"),
		got:  NewST(""),
	}, {
		name: "ID",
		p:    NewID("value"),
		want: NewID("value"),
		got:  NewID(""),
	}, {
		name: "SI",
		p:    NewSI(44),
		want: NewSI(44),
		got:  &SI{Valid: false},
	}, {
		name: "NM",
		p:    NewNM(44),
		want: NewNM(44),
		got:  &NM{Valid: false},
	}, {
		name: "IS",
		p:    NewIS("value"),
		want: NewIS("value"),
		got:  NewIS(""),
	}, {
		name: "DT",
		p:    NewDT("value"),
		want: NewDT("value"),
		got:  NewDT(""),
	}, {
		name: "TM",
		p:    NewTM("value"),
		want: NewTM("value"),
		got:  NewTM(""),
	}, {
		name: "TS_YearPrecision",
		p:    &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 0, time.UTC), Precision: YearPrecision},
		want: &TS{IsHL7Null: false, Time: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC), Precision: YearPrecision},
		got:  &TS{IsHL7Null: true},
	}, {
		name: "TS_MonthPrecision",
		p:    &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 0, time.UTC), Precision: MonthPrecision},
		want: &TS{IsHL7Null: false, Time: time.Date(2020, 02, 01, 0, 0, 0, 0, time.UTC), Precision: MonthPrecision},
		got:  &TS{IsHL7Null: true},
	}, {
		name: "TS_DayPrecision",
		p:    &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 0, time.UTC), Precision: DayPrecision},
		want: &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 0, 0, 0, 0, time.UTC), Precision: DayPrecision},
		got:  &TS{IsHL7Null: true},
	}, {
		name: "TS_HourPrecision",
		p:    &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 0, time.UTC), Precision: HourPrecision},
		want: &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 0, 0, 0, time.UTC), Precision: HourPrecision},
		got:  &TS{IsHL7Null: true},
	}, {
		name: "TS_MinutePrecision",
		p:    &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 0, time.UTC), Precision: MinutePrecision},
		want: &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 0, 0, time.UTC), Precision: MinutePrecision},
		got:  &TS{IsHL7Null: true},
	}, {
		name: "TS_SecondPrecision",
		p:    &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 0, time.UTC), Precision: SecondPrecision},
		want: &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 0, time.UTC), Precision: SecondPrecision},
		got:  &TS{IsHL7Null: true},
	}, {
		name: "TS_SecondPrecision_WithNanoseconds",
		p:    &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 5123456789, time.UTC), Precision: SecondPrecision},
		want: &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 35, 0, time.UTC), Precision: SecondPrecision},
		got:  &TS{IsHL7Null: true},
	}, {
		name: "TS_TenthSecondPrecision",
		p:    &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 5123456789, time.UTC), Precision: TenthSecondPrecision},
		want: &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 35, 100000000, time.UTC), Precision: TenthSecondPrecision},
		got:  &TS{IsHL7Null: true},
	}, {
		name: "TS_HundredthSecondPrecision",
		p:    &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 5123456789, time.UTC), Precision: HundredthSecondPrecision},
		want: &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 35, 120000000, time.UTC), Precision: HundredthSecondPrecision},
		got:  &TS{IsHL7Null: true},
	}, {
		name: "TS_ThousandthSecondPrecision",
		p:    &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 5123456789, time.UTC), Precision: ThousandthSecondPrecision},
		want: &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 35, 123000000, time.UTC), Precision: ThousandthSecondPrecision},
		got:  &TS{IsHL7Null: true},
	}, {
		name: "TS_TenThousandthSecondPrecision",
		p:    &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 5123456789, time.UTC), Precision: TenThousandthSecondPrecision},
		want: &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 35, 123400000, time.UTC), Precision: TenThousandthSecondPrecision},
		got:  &TS{IsHL7Null: true},
	}, {
		name: "TN",
		p:    NewTN("value"),
		want: NewTN("value"),
		got:  NewTN(""),
	}, {
		name: "FT",
		p:    NewFT("value"),
		want: NewFT("value"),
		got:  NewFT(""),
	}, {
		name: "TX",
		p:    NewTX("value"),
		want: NewTX("value"),
		got:  NewTX(""),
	}, {
		name: "CM",
		p:    NewCM([]byte("value")),
		want: NewCM([]byte("value")),
		got:  NewCM([]byte{}),
	}, {
		name: "Any",
		p:    NewAny([]byte("value")),
		want: NewAny([]byte("value")),
		got:  NewAny([]byte{}),
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b, err := tc.p.Marshal(c)
			if err != nil {
				t.Fatalf("[%+v].Marshal(%v) failed with %v", tc.p, c, err)
			}

			if err := tc.got.Unmarshal(b, c); err != nil {
				t.Fatalf("[%+v].Unmarshal(%s, %v) failed with %v", tc.got, string(b), c, err)
			}
			if diff := cmp.Diff(tc.want, tc.got); diff != "" {
				t.Errorf("[%+v].Unmarshal(%s, %v) got diff (-want, +got):\n%s", tc.got, string(b), c, diff)
			}
		})
	}
}

func TestSanitizedString(t *testing.T) {
	type sanitizable interface {
		SanitizedString() string
	}

	cases := []struct {
		name string
		s    sanitizable
		want string
	}{
		{name: "ST", s: NewST("value"), want: "value"},
		{name: "ST null", s: NewST(`""`), want: ""},
		{name: "ID", s: NewID("value"), want: "value"},
		{name: "ID null", s: NewID(`""`), want: ""},
		{name: "IS", s: NewIS("value"), want: "value"},
		{name: "IS null", s: NewIS(`""`), want: ""},
		{
			name: "HD",
			s: &HD{
				NamespaceID:     NewIS("NamespaceID"),
				UniversalID:     NewST("UniversalID"),
				UniversalIDType: NewID("UniversalIDType"),
			},
			want: "NamespaceID^UniversalID^UniversalIDType",
		}, {
			name: "HD every field null",
			s: &HD{
				NamespaceID:     NewIS(`""`),
				UniversalID:     NewST(`""`),
				UniversalIDType: NewID(`""`),
			},
			want: `""^""^""`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.s.SanitizedString(); got != tc.want {
				t.Errorf("[%+v].SanitizedString()=%q, want %q", tc.s, got, tc.want)
			}
		})
	}
}

type empty interface {
	Empty() bool
}

func TestEmpty(t *testing.T) {
	cases := []struct {
		name string
		f    func(s string) empty
	}{
		{name: "ST", f: func(s string) empty { return NewST(ST(s)) }},
		{name: "ID", f: func(s string) empty { return NewID(ID(s)) }},
		{name: "IS", f: func(s string) empty { return NewIS(IS(s)) }},
	}

	for _, tc := range cases {
		for k, want := range map[empty]bool{tc.f("value"): false, tc.f(""): true} {
			t.Run(fmt.Sprintf("%s-%s", tc.name, k), func(t *testing.T) {
				if got := k.Empty(); got != want {
					t.Errorf("[%+v].Empty()=%t, want %t", k, got, want)
				}
			})
		}
	}
}

func TestEmptyNilEmpty(t *testing.T) {
	cases := []struct {
		name string
		e    empty
	}{
		{name: "ST", e: new(ST)},
		{name: "ID", e: new(ID)},
		{name: "IS", e: new(IS)},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.e.Empty(); !got {
				t.Errorf("[%+v].Empty()=%t, want true", tc.e, got)
			}
		})
	}
}

func TestParseTS(t *testing.T) {
	tests := []struct {
		in           string
		outTime      string
		outPrecision TSPrecision
	}{
		{"20141128001635", "2014-11-28T00:16:35Z", SecondPrecision},
		{"20141128001635^M", "2014-11-28T00:16:35Z", MinutePrecision},
		{"20141128001635.1", "2014-11-28T00:16:35.1Z", TenthSecondPrecision},
		{"20141128001635.12", "2014-11-28T00:16:35.12Z", HundredthSecondPrecision},
		{"20141128001635.123", "2014-11-28T00:16:35.123Z", ThousandthSecondPrecision},
		{"20141128001635.1234", "2014-11-28T00:16:35.1234Z", TenThousandthSecondPrecision},
		// The following examples are from the HL7 specification
		{"19760704010159-0600", "1976-07-04T01:01:59-06:00", SecondPrecision},
		{"19760704010159-0500", "1976-07-04T01:01:59-05:00", SecondPrecision},
		{"198807050000", "1988-07-04T23:00:00Z", MinutePrecision},
		{"19880705", "1988-07-05T00:00:00Z", DayPrecision},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			var ts TS
			err := ts.Unmarshal([]byte(tt.in), testContext)
			if err != nil {
				t.Errorf("ParseTS(%q) got error %v, want err=<nil>", tt.in, err)
			}
			want, _ := time.Parse(time.RFC3339Nano, tt.outTime)
			if !want.Equal(ts.Time) {
				t.Errorf("ParseTS(%q).Time got %v, want %v", tt.in, ts.Time, want)
			}
			if diff := cmp.Diff(tt.outPrecision, ts.Precision); diff != "" {
				t.Errorf("ParseTS(%q).Precision mismatch (-want, +got)=\n%s", tt.in, diff)
			}
		})
	}
}

func TestMarshalTS(t *testing.T) {
	for _, tt := range []struct {
		name string
		in   *TS
		want string
	}{{
		name: "TS_YearPrecision",
		in:   &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 0, testLocation), Precision: YearPrecision},
		want: "2020",
	}, {
		name: "TS_MonthPrecision",
		in:   &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 0, testLocation), Precision: MonthPrecision},
		want: "202002",
	}, {
		name: "TS_DayPrecision",
		in:   &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 0, testLocation), Precision: DayPrecision},
		want: "20200224",
	}, {
		name: "TS_HourPrecision",
		in:   &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 0, testLocation), Precision: HourPrecision},
		want: "2020022412",
	}, {
		name: "TS_MinutePrecision",
		in:   &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 0, testLocation), Precision: MinutePrecision},
		want: "202002241255",
	}, {
		name: "TS_SecondPrecision",
		in:   &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 0, testLocation), Precision: SecondPrecision},
		want: "20200224125530",
	}, {
		name: "TS_SecondPrecision_WithNanoseconds",
		in:   &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 5123456789, testLocation), Precision: SecondPrecision},
		want: "20200224125535",
	}, {
		name: "TS_TenthSecondPrecision",
		in:   &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 5123456789, testLocation), Precision: TenthSecondPrecision},
		want: "20200224125535.1",
	}, {
		name: "TS_HundredthSecondPrecision",
		in:   &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 5123456789, testLocation), Precision: HundredthSecondPrecision},
		want: "20200224125535.12",
	}, {
		name: "TS_ThousandthSecondPrecision",
		in:   &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 5123456789, testLocation), Precision: ThousandthSecondPrecision},
		want: "20200224125535.123",
	}, {
		name: "TS_TenThousandthSecondPrecision",
		in:   &TS{IsHL7Null: false, Time: time.Date(2020, 02, 24, 12, 55, 30, 5123456789, testLocation), Precision: TenThousandthSecondPrecision},
		want: "20200224125535.1234",
	}} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.in.Marshal(testContext)
			if err != nil {
				t.Fatalf("[%+v].Marshal(%+v) failed with %+v", tt.in, testContext, err)
			}
			if got, want := string(got), tt.want; got != want {
				t.Errorf("[%+v].Marshal(%+v)=%q, want %q", tt.in, testContext, got, want)
			}
		})
	}
}

func TestParseTS_ClearField(t *testing.T) {
	var ts TS
	if err := ts.Unmarshal([]byte("\"\""), testContext); err != nil {
		t.Fatalf("Unmarshal('') failed with %v", err)
	}
	if !ts.Time.IsZero() {
		t.Error("ts.Time.IsZero() is false, want true")
	}
}

func TestParseTS_Error(t *testing.T) {
	tests := []string{
		// The empty string
		"",
		// A two digit year
		"20",
		// Fractions of a second with seconds
		"201411280016.12",
		// More precision than thousandths of a second
		"20141128001635.12345",
		// A unknown (legacy) precision value
		"20141128001635^T",
		// A timezone without the correct number of digits
		"201411280016+010",
	}
	for _, tt := range tests {
		var ts TS
		if err := ts.Unmarshal([]byte(tt), testContext); err == nil {
			t.Errorf("Unmarshal(%q) got err=<nil>, want error", tt)
		}
	}
}

func TestParseSI(t *testing.T) {
	tests := []struct {
		in  string
		out SI
	}{
		{"0", SI{Value: 0, Valid: true}},
		{"1", SI{Value: 1, Valid: true}},
		{"2", SI{Value: 2, Valid: true}},
		{"112233445566", SI{Value: 112233445566, Valid: true}},
		{`""`, SI{Valid: false}},
	}
	for _, tt := range tests {
		var si SI
		if err := si.Unmarshal([]byte(tt.in), testContext); err != nil {
			t.Fatalf("ParseSI(%q) failed with %v", tt.in, err)
		}
		if diff := cmp.Diff(tt.out, si); diff != "" {
			t.Errorf("ParseSI(%q) mismatch (-want, +got)=\n%s", tt.in, diff)
		}
	}
}

func TestParseSI_Error(t *testing.T) {
	tests := []string{
		"",
		"-",
		" ",
		"-1",  // Only non-negative numbers allowed.
		"1.2", // Only integer numbers allowed.
		"2-1",
	}
	for _, tt := range tests {
		var si SI
		if err := si.Unmarshal([]byte(tt), testContext); err == nil {
			t.Errorf("ParseSI(%q) got err=<nil>, want error", tt)
		}
	}
}

func TestParseNM(t *testing.T) {
	tests := []struct {
		in  string
		out NM
	}{
		{"0", NM{Value: 0.0, Valid: true}},
		{"-0", NM{Value: 0.0, Valid: true}},
		{"0.0", NM{Value: 0.0, Valid: true}},
		{"-0.0", NM{Value: 0.0, Valid: true}},
		{"0011.2200", NM{Value: 11.22, Valid: true}},   // Leading/trailing zeroes.
		{"-0011.2200", NM{Value: -11.22, Valid: true}}, // Leading/trailing zeroes.
		{"112233445566", NM{Value: 112233445566.0, Valid: true}},
		{"-112233445566", NM{Value: -112233445566.0, Valid: true}},
		{"112233445566.77", NM{Value: 112233445566.77, Valid: true}},
		{"-112233445566.77", NM{Value: -112233445566.77, Valid: true}},
		{`""`, NM{Valid: false}},
	}
	for _, tt := range tests {
		var nm NM
		if err := nm.Unmarshal([]byte(tt.in), testContext); err != nil {
			t.Errorf("ParseNM(%q) failed with %v", tt.in, err)
		}
		if diff := cmp.Diff(NM(tt.out), nm); diff != "" {
			t.Errorf("ParseNM(%q) mismatch (-want, +got)=\n%s", tt.in, diff)
		}
	}
}

func TestParseNM_Error(t *testing.T) {
	tests := []string{
		"",
		"-",
		" ",
		"2-1",
	}
	for _, tt := range tests {
		var nm NM
		if err := nm.Unmarshal([]byte(tt), testContext); err == nil {
			t.Errorf("ParseNM(%q) got err=<nil>, want error", tt)
		}
	}
}

func TestParseFT_EscapesText(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"One|Field", `One\F\Field`},
		{"Many|Fields|a|b", `Many\F\Fields\F\a\F\b`},
		{"Component1^Component2^Component3", `Component1\S\Component2\S\Component3`},
		{"Subcomponent1&Subcomponent2&", `Subcomponent1\T\Subcomponent2\T\`},
		{"Reptition1~Repetition2~Repetition3~", `Reptition1\R\Repetition2\R\Repetition3\R\`},
		{"line break 1\nline break 2\n", `line break 1\.br\line break 2\.br\`},
		{"two new lines\n\ntwo new lines\n\n", `two new lines\.br\\.br\two new lines\.br\\.br\`},
	}
	for _, test := range tests {
		ft := FT(test.in)
		out, err := ft.Marshal(testContext)
		if err != nil {
			t.Fatalf("Marshal(%q) failed with %v", test.in, err)
		}
		if got, want := string(out), test.want; got != want {
			t.Errorf("Marshal(%q) got %v, want %v", test.in, got, want)
		}
	}
}

func TestHDString(t *testing.T) {
	tests := []struct {
		in      HD
		wantOut string
	}{
		{HD{NamespaceID: NewIS("namespace"),
			UniversalID:     NewST("ID"),
			UniversalIDType: NewID("IDType")}, "namespace^ID^IDType"},
		{HD{NamespaceID: NewIS("namespace")}, "namespace"},
		{HD{UniversalID: NewST("UID")}, "^UID"},
		{HD{NamespaceID: NewIS("namespace"),
			UniversalIDType: NewID("IDType")}, "namespace^^IDType"},
		{HD{}, ""},
	}
	for _, tt := range tests {
		if got, want := tt.in.String(), tt.wantOut; got != want {
			t.Errorf("[%v].String() got %v, want %v", tt.in, got, want)
		}
	}
}
