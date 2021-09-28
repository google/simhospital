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
