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

package orderprofile

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/google/simhospital/pkg/constants"
)

type valRange struct {
	from float64
	to   float64
}

func TestValueGeneratorFromRange(t *testing.T) {
	cases := []struct {
		name       string
		inRange    string
		wantNormal valRange
		wantHigh   valRange
		wantLow    valRange
	}{
		{name: "simple", inRange: "12 - 24", wantNormal: valRange{from: 12, to: 24}, wantHigh: valRange{from: 24, to: 240}, wantLow: valRange{from: 0, to: 12}},
		{name: "simple", inRange: "0.5 - 1.6", wantNormal: valRange{from: 0.5, to: 1.6}, wantHigh: valRange{from: 1.6, to: 16}, wantLow: valRange{from: 0, to: 0.5}},
		{name: "simple", inRange: "-2 - 3", wantNormal: valRange{from: -2, to: 3}, wantHigh: valRange{from: 3, to: 30}, wantLow: valRange{from: -20, to: -2}},
		{name: "simple", inRange: "-2 - -0.5", wantNormal: valRange{from: -2, to: -0.5}, wantHigh: valRange{from: -0.5, to: 0}, wantLow: valRange{from: -20, to: -2}},
		{name: "no spaces", inRange: "12-24", wantNormal: valRange{from: 12, to: 24}, wantHigh: valRange{from: 24, to: 240}, wantLow: valRange{from: 0, to: 12}},
		{name: "no spaces", inRange: "0.5-1.6", wantNormal: valRange{from: 0.5, to: 1.6}, wantHigh: valRange{from: 1.6, to: 16}, wantLow: valRange{from: 0, to: 0.5}},
		{name: "no spaces", inRange: "-2-3", wantNormal: valRange{from: -2, to: 3}, wantHigh: valRange{from: 3, to: 30}, wantLow: valRange{from: -20, to: -2}},
		{name: "no spaces", inRange: "-2--0.5", wantNormal: valRange{from: -2, to: -0.5}, wantHigh: valRange{from: -0.5, to: 0}, wantLow: valRange{from: -20, to: -2}},
		{name: "with brackets", inRange: "[ 12 - 24 ]", wantNormal: valRange{from: 12, to: 24}, wantHigh: valRange{from: 24, to: 240}, wantLow: valRange{from: 0, to: 12}},
		{name: "with brackets", inRange: "[ 0.5 - 1.6 ]", wantNormal: valRange{from: 0.5, to: 1.6}, wantHigh: valRange{from: 1.6, to: 16}, wantLow: valRange{from: 0, to: 0.5}},
		{name: "with brackets", inRange: "[ -2 - 3 ]", wantNormal: valRange{from: -2, to: 3}, wantHigh: valRange{from: 3, to: 30}, wantLow: valRange{from: -20, to: -2}},
		{name: "with brackets no spaces", inRange: "[12-24]", wantNormal: valRange{from: 12, to: 24}, wantHigh: valRange{from: 24, to: 240}, wantLow: valRange{from: 0, to: 12}},
		{name: "with brackets", inRange: "[ -2 - -0.5 ]", wantNormal: valRange{from: -2, to: -0.5}, wantHigh: valRange{from: -0.5, to: 0}, wantLow: valRange{from: -20, to: -2}},
		{name: "with brackets no spaces", inRange: "[0.5-1.6]", wantNormal: valRange{from: 0.5, to: 1.6}, wantHigh: valRange{from: 1.6, to: 16}, wantLow: valRange{from: 0, to: 0.5}},
		{name: "with brackets no spaces", inRange: "[-2-3]", wantNormal: valRange{from: -2, to: 3}, wantHigh: valRange{from: 3, to: 30}, wantLow: valRange{from: -20, to: -2}},
		{name: "with brackets no spaces", inRange: "[-2--0.5]", wantNormal: valRange{from: -2, to: -0.5}, wantHigh: valRange{from: -0.5, to: 0}, wantLow: valRange{from: -20, to: -2}},
		{name: "duplicated range", inRange: "12-24^12^24", wantNormal: valRange{from: 12, to: 24}, wantHigh: valRange{from: 24, to: 240}, wantLow: valRange{from: 0, to: 12}},
		{name: "duplicated range", inRange: "0.5-1.6^0.5^1.6", wantNormal: valRange{from: 0.5, to: 1.6}, wantHigh: valRange{from: 1.6, to: 16}, wantLow: valRange{from: 0, to: 0.5}},
		{name: "duplicated range", inRange: "-2-3^-2^3", wantNormal: valRange{from: -2, to: 3}, wantHigh: valRange{from: 3, to: 30}, wantLow: valRange{from: -20, to: -2}},
		{name: "duplicated range", inRange: "-2--0.5^-2^-0.5", wantNormal: valRange{from: -2, to: -0.5}, wantHigh: valRange{from: -0.5, to: 0}, wantLow: valRange{from: -20, to: -2}},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s-%s", tc.name, tc.inRange), func(t *testing.T) {
			vg, err := ValueGeneratorFromRange(tc.inRange)
			if err != nil {
				t.Fatalf("ValueGeneratorFromRange(%s) failed with err %v", tc.inRange, err)
			}

			// Generate Normal, AbnormalHigh and AbnormalLow values multiple times.
			for i := 0; i < 1; i++ {
				gotNormal, err := vg.Normal()
				if err != nil {
					t.Fatalf("Normal() failed with err %v", err)
				}
				gotf, err := strconv.ParseFloat(gotNormal, 64)
				if err != nil {
					t.Fatalf("ParseFloat(%q, 64) failed with err %v", gotNormal, err)
				}
				if gotf <= tc.wantNormal.from || gotf >= tc.wantNormal.to {
					t.Errorf("Normal() = %q, want in range (%f, %f)", gotNormal, tc.wantNormal.from, tc.wantNormal.to)
				}

				gotHigh, err := vg.AbnormalHigh()
				if err != nil {
					t.Fatalf("AbnormalHigh() failed with err %v", err)
				}
				gotf, err = strconv.ParseFloat(gotHigh, 64)
				if err != nil {
					t.Fatalf("ParseFloat(%q, 64) failed with err %v", gotHigh, err)
				}
				if gotf <= tc.wantHigh.from || gotf >= tc.wantHigh.to {
					t.Errorf("AbnormalHigh() = %q, want in range (%f, %f)", gotHigh, tc.wantHigh.from, tc.wantHigh.to)
				}

				gotLow, err := vg.AbnormalLow()
				if err != nil {
					t.Fatalf("AbnormalLow() failed with err %v", err)
				}
				gotf, err = strconv.ParseFloat(gotLow, 64)
				if err != nil {
					t.Fatalf("ParseFloat(%q, 64) failed with err %v", gotLow, err)
				}
				if gotf <= tc.wantLow.from || gotf >= tc.wantLow.to {
					t.Errorf("AbnormalLow() = %q, want in range (%f, %f)", gotLow, tc.wantLow.from, tc.wantLow.to)
				}
			}
		})
	}
}

func TestValueGeneratorFromRange_LeftOpenRange(t *testing.T) {
	cases := []struct {
		name       string
		inRange    string
		wantNormal valRange
		wantHigh   valRange
	}{
		{name: "strictly less", inRange: "<55^^<55", wantNormal: valRange{from: 0, to: 55}, wantHigh: valRange{from: 55, to: 550}},
		{name: "strictly less", inRange: "<5.5^^<5.5", wantNormal: valRange{from: 0, to: 5.5}, wantHigh: valRange{from: 5.5, to: 55}},
		{name: "strictly less", inRange: "<-5.5^^<-5.5", wantNormal: valRange{from: -55, to: -5.5}, wantHigh: valRange{from: -5.5, to: 0}},
		{name: "less or equal", inRange: "<=55^^<=55", wantNormal: valRange{from: 0, to: 55}, wantHigh: valRange{from: 55, to: 550}},
		{name: "less or equal", inRange: "<=5.5^^<=5.5", wantNormal: valRange{from: 0, to: 5.5}, wantHigh: valRange{from: 5.5, to: 55}},
		{name: "less or equal", inRange: "<=-5.5^<=-5.5", wantNormal: valRange{from: -55, to: -5.5}, wantHigh: valRange{from: -5.5, to: 0}},
		{name: "strictly less with brackets", inRange: "[ < 55 ]", wantNormal: valRange{from: 0, to: 55}, wantHigh: valRange{from: 55, to: 550}},
		{name: "strictly less with brackets", inRange: "[ < 5.5 ]", wantNormal: valRange{from: 0, to: 5.5}, wantHigh: valRange{from: 5.5, to: 55}},
		{name: "strictly less with brackets", inRange: "[ < -5.5 ]", wantNormal: valRange{from: -55, to: -5.5}, wantHigh: valRange{from: -5.5, to: 0}},
		{name: "less or equal with brackets", inRange: "[ <= 55 ]", wantNormal: valRange{from: 0, to: 55}, wantHigh: valRange{from: 55, to: 550}},
		{name: "less or equal with brackets", inRange: "[ <= 5.5 ]", wantNormal: valRange{from: 0, to: 5.5}, wantHigh: valRange{from: 5.5, to: 55}},
		{name: "less or equal with brackets", inRange: "[ <= -5.5 ]", wantNormal: valRange{from: -55, to: -5.5}, wantHigh: valRange{from: -5.5, to: 0}},
		{name: "brackets no spaces", inRange: "[<=55]", wantNormal: valRange{from: 0, to: 55}, wantHigh: valRange{from: 55, to: 550}},
		{name: "brackets no spaces", inRange: "[<5.5]", wantNormal: valRange{from: 0, to: 5.5}, wantHigh: valRange{from: 5.5, to: 55}},
		{name: "brackets no spaces", inRange: "[<=-5.5]", wantNormal: valRange{from: -55, to: -5.5}, wantHigh: valRange{from: -5.5, to: 0}},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s-%s", tc.name, tc.inRange), func(t *testing.T) {
			vg, err := ValueGeneratorFromRange(tc.inRange)
			if err != nil {
				t.Fatalf("ValueGeneratorFromRange(%s) failed with err %v", tc.inRange, err)
			}

			// Generate Normal, AbnormalHigh and AbnormalLow values multiple times.
			for i := 0; i < 1; i++ {
				gotNormal, err := vg.Normal()
				if err != nil {
					t.Fatalf("Normal() failed with err %v", err)
				}
				gotf, err := strconv.ParseFloat(gotNormal, 64)
				if err != nil {
					t.Fatalf("ParseFloat(%q, 64) failed with err %v", gotNormal, err)
				}
				if gotf <= tc.wantNormal.from || gotf >= tc.wantNormal.to {
					t.Errorf("Normal() = %q, want in range (%f, %f)", gotNormal, tc.wantNormal.from, tc.wantNormal.to)
				}

				gotHigh, err := vg.AbnormalHigh()
				if err != nil {
					t.Fatalf("AbnormalHigh() failed with err %v", err)
				}
				gotf, err = strconv.ParseFloat(gotHigh, 64)
				if err != nil {
					t.Fatalf("ParseFloat(%q, 64) failed with err %v", gotHigh, err)
				}
				if gotf <= tc.wantHigh.from || gotf >= tc.wantHigh.to {
					t.Errorf("AbnormalHigh() = %q, want in range (%f, %f)", gotHigh, tc.wantHigh.from, tc.wantHigh.to)
				}

				if _, err = vg.AbnormalLow(); err == nil {
					t.Error("AbnormalLow() got nil err, want non-nill err")
				}
			}
		})
	}
}

func TestValueGeneratorFromRange_RightOpenRange(t *testing.T) {
	cases := []struct {
		name       string
		inRange    string
		wantNormal valRange
		wantLow    valRange
	}{
		{name: "strictly greater", inRange: ">55^^>55", wantNormal: valRange{from: 55, to: 550}, wantLow: valRange{from: 0, to: 55}},
		{name: "strictly greater", inRange: ">5.5^^>5.5", wantNormal: valRange{from: 5.5, to: 55}, wantLow: valRange{from: 0, to: 5.5}},
		{name: "strictly greater", inRange: ">-5.5^^>-5.5", wantNormal: valRange{from: -5.5, to: 0}, wantLow: valRange{from: -55, to: -5.5}},
		{name: "greater or equal", inRange: ">=55^^>=55", wantNormal: valRange{from: 55, to: 550}, wantLow: valRange{from: 0, to: 55}},
		{name: "greater or equal", inRange: ">=5.5^^>=5.5", wantNormal: valRange{from: 5.5, to: 55}, wantLow: valRange{from: 0, to: 5.5}},
		{name: "greater or equal", inRange: ">=-5.5^>=-5.5", wantNormal: valRange{from: -5.5, to: 0}, wantLow: valRange{from: -55, to: -5.5}},
		{name: "strictly greater with brackets", inRange: "[ > 55 ]", wantNormal: valRange{from: 55, to: 550}, wantLow: valRange{from: 0, to: 55}},
		{name: "strictly greater with brackets", inRange: "[ > 5.5 ]", wantNormal: valRange{from: 5.5, to: 55}, wantLow: valRange{from: 0, to: 5.5}},
		{name: "strictly greater with brackets", inRange: "[ > -5.5 ]", wantNormal: valRange{from: -5.5, to: 0}, wantLow: valRange{from: -55, to: -5.5}},
		{name: "greater or equal with brackets", inRange: "[ >= 55 ]", wantNormal: valRange{from: 55, to: 550}, wantLow: valRange{from: 0, to: 55}},
		{name: "greater or equal with brackets", inRange: "[ >= 5.5 ]", wantNormal: valRange{from: 5.5, to: 55}, wantLow: valRange{from: 0, to: 5.5}},
		{name: "greater or equal with brackets", inRange: "[ >= -5.5 ]", wantNormal: valRange{from: -5.5, to: 0}, wantLow: valRange{from: -55, to: -5.5}},
		{name: "brackets no spaces", inRange: "[>=55]", wantNormal: valRange{from: 55, to: 550}, wantLow: valRange{from: 0, to: 55}},
		{name: "brackets no spaces", inRange: "[>5.5]", wantNormal: valRange{from: 5.5, to: 55}, wantLow: valRange{from: 0, to: 5.5}},
		{name: "brackets no spaces", inRange: "[>=-5.5]", wantNormal: valRange{from: -5.5, to: 0}, wantLow: valRange{from: -55, to: -5.5}},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s-%s", tc.name, tc.inRange), func(t *testing.T) {
			vg, err := ValueGeneratorFromRange(tc.inRange)
			if err != nil {
				t.Fatalf("ValueGeneratorFromRange(%s) failed with err %v", tc.inRange, err)
			}

			// Generate Normal, AbnormalHigh and AbnormalLow values multiple times.
			for i := 0; i < 1; i++ {
				gotNormal, err := vg.Normal()
				if err != nil {
					t.Fatalf("Normal() failed with err %v", err)
				}
				gotf, err := strconv.ParseFloat(gotNormal, 64)
				if err != nil {
					t.Fatalf("ParseFloat(%q, 64) failed with err %v", gotNormal, err)
				}
				if gotf <= tc.wantNormal.from || gotf >= tc.wantNormal.to {
					t.Errorf("Normal() = %q, want in range (%f, %f)", gotNormal, tc.wantNormal.from, tc.wantNormal.to)
				}

				gotLow, err := vg.AbnormalLow()
				if err != nil {
					t.Fatalf("AbnormalLow() failed with err %v", err)
				}
				gotf, err = strconv.ParseFloat(gotLow, 64)
				if err != nil {
					t.Fatalf("ParseFloat(%q, 64) failed with err %v", gotLow, err)
				}
				if gotf <= tc.wantLow.from || gotf >= tc.wantLow.to {
					t.Errorf("AbnormalLow() = %q, want in range (%f, %f)", gotLow, tc.wantLow.from, tc.wantLow.to)
				}

				if _, err := vg.AbnormalHigh(); err == nil {
					t.Error("AbnormalHigh() got nil err, want non-nill err")
				}
			}
		})
	}
}

type errOrRange struct {
	err      bool
	valRange valRange
}

func TestValueGeneratorFromRange_ZeroOrSmallRange(t *testing.T) {
	cases := []struct {
		name       string
		inRange    string
		wantNormal errOrRange
		wantHigh   errOrRange
		wantLow    errOrRange
	}{
		{name: "from zero", inRange: "0 - 24", wantNormal: errOrRange{valRange: valRange{from: 0, to: 24}}, wantHigh: errOrRange{valRange: valRange{from: 24, to: 240}}, wantLow: errOrRange{err: true}},
		{name: "to zero", inRange: "-2 - 0", wantNormal: errOrRange{valRange: valRange{from: -2, to: 0}}, wantHigh: errOrRange{err: true}, wantLow: errOrRange{valRange: valRange{from: -20, to: -2}}},
		{name: "grater than zero", inRange: "[>0]", wantNormal: errOrRange{valRange: valRange{from: 0, to: 10}}, wantHigh: errOrRange{err: true}, wantLow: errOrRange{err: true}},
		{name: "grater or equal to zero", inRange: "[>=0]", wantNormal: errOrRange{valRange: valRange{from: 0, to: 10}}, wantHigh: errOrRange{err: true}, wantLow: errOrRange{err: true}},
		{name: "less than zero", inRange: "[<0]", wantNormal: errOrRange{valRange: valRange{from: -10, to: 0}}, wantHigh: errOrRange{err: true}, wantLow: errOrRange{err: true}},
		{name: "less or equal to zero", inRange: "[<=0]", wantNormal: errOrRange{valRange: valRange{from: -10, to: 0}}, wantHigh: errOrRange{err: true}, wantLow: errOrRange{err: true}},
		{name: "small range", inRange: "0.1 - 0.3", wantNormal: errOrRange{valRange: valRange{from: 0.1, to: 0.3}}, wantHigh: errOrRange{valRange: valRange{from: 0.3, to: 3}}, wantLow: errOrRange{valRange: valRange{from: 0, to: 0.1}}},
		{name: "small range - impossible to generate normal value", inRange: "0.11 - 0.12", wantNormal: errOrRange{err: true}, wantHigh: errOrRange{valRange: valRange{from: 0.12, to: 1.2}}, wantLow: errOrRange{valRange: valRange{from: 0, to: 0.11}}},
		{name: "small range - impossible to generate abnormal low", inRange: "0.01 - 0.12", wantNormal: errOrRange{valRange: valRange{from: 0.01, to: 0.12}}, wantHigh: errOrRange{valRange: valRange{from: 0.12, to: 1.2}}, wantLow: errOrRange{err: true}},
		{name: "small range - impossible to generate abnormal high", inRange: "-0.12 - -0.01", wantNormal: errOrRange{valRange: valRange{from: -0.12, to: -0.01}}, wantHigh: errOrRange{err: true}, wantLow: errOrRange{valRange: valRange{from: -1.2, to: -0.12}}},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s-%s", tc.name, tc.inRange), func(t *testing.T) {
			vg, err := ValueGeneratorFromRange(tc.inRange)
			if err != nil {
				t.Fatalf("ValueGeneratorFromRange(%s) failed with err %v", tc.inRange, err)
			}

			// Generate Normal, AbnormalHigh and AbnormalLow values multiple times.
			for i := 0; i < 1; i++ {
				gotNormal, err := vg.Normal()
				if gotErr := err != nil; gotErr != tc.wantNormal.err {
					t.Fatalf("Normal() got err %v, want err? %t", err, tc.wantNormal.err)
				}
				if !tc.wantNormal.err {
					gotf, err := strconv.ParseFloat(gotNormal, 64)
					if err != nil {
						t.Fatalf("ParseFloat(%q, 64) failed with err %v", gotNormal, err)
					}
					if gotf <= tc.wantNormal.valRange.from || gotf >= tc.wantNormal.valRange.to {
						t.Errorf("Normal() = %q, want in range (%f, %f)", gotNormal, tc.wantNormal.valRange.from, tc.wantNormal.valRange.to)
					}
				}

				gotHigh, err := vg.AbnormalHigh()
				if gotErr := err != nil; gotErr != tc.wantHigh.err {
					t.Fatalf("AbnormalHigh() got err %v, want err? %t", err, tc.wantHigh.err)
				}
				if !tc.wantHigh.err {
					gotf, err := strconv.ParseFloat(gotHigh, 64)
					if err != nil {
						t.Fatalf("ParseFloat(%q, 64) failed with err %v", gotHigh, err)
					}
					if gotf <= tc.wantHigh.valRange.from || gotf >= tc.wantHigh.valRange.to {
						t.Errorf("AbnormalHigh() = %q, want in range (%f, %f)", gotHigh, tc.wantHigh.valRange.from, tc.wantHigh.valRange.to)
					}
				}

				gotLow, err := vg.AbnormalLow()
				if gotErr := err != nil; gotErr != tc.wantLow.err {
					t.Fatalf("AbnormalLow() got err %v, want err? %t", err, tc.wantLow.err)
				}

				if !tc.wantLow.err {
					gotf, err := strconv.ParseFloat(gotLow, 64)
					if err != nil {
						t.Fatalf("ParseFloat(%q, 64) failed with err %v", gotLow, err)
					}
					if gotf <= tc.wantLow.valRange.from || gotf >= tc.wantLow.valRange.to {
						t.Errorf("AbnormalLow() = %q, want in range (%f, %f)", gotLow, tc.wantLow.valRange.from, tc.wantLow.valRange.to)
					}
				}
			}
		})
	}
}

func TestValueGeneratorFromRange_InvalidRange(t *testing.T) {
	cases := []struct {
		name    string
		inRange string
		wantErr bool
	}{
		{name: "valid", inRange: "12 - 24", wantErr: false},
		{name: "invalid range", inRange: "15 - 0", wantErr: true},
		{name: "start range invalid", inRange: "0.5.6 - 1.6", wantErr: true},
		{name: "end range invalid", inRange: "0.5 - 1.6.3", wantErr: true},
		{name: "start range invalid - repeated range", inRange: "71.5.6-121.5^71.5.6^121.5", wantErr: true},
		{name: "range invalid - greater than", inRange: "[>17.3.3]", wantErr: true},
		{name: "range invalid - less than than", inRange: "[<=17.3.3]", wantErr: true},
		{name: "range not a number", inRange: "from - to", wantErr: true},
		{name: "empty range", inRange: "-", wantErr: true},
		{name: "any other string", inRange: "any other string", wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ValueGeneratorFromRange(tc.inRange)
			if gotErr := err != nil; gotErr != tc.wantErr {
				t.Fatalf("ValueGeneratorFromRange(%s) got err %v, want err? %t", tc.inRange, err, tc.wantErr)
			}
		})
	}
}

func TestValueGenerator_NotInitialised(t *testing.T) {
	cases := []struct {
		name          string
		vg            *ValueGenerator
		wantNormal    float64
		wantNormalErr bool
		wantHighErr   bool
		wantLowErr    bool
	}{
		{name: "nil ValueGenerator", vg: nil, wantNormal: 0.0, wantNormalErr: false, wantHighErr: true, wantLowErr: true},
		{name: "empty ValueGenerator", vg: new(ValueGenerator), wantNormalErr: true, wantHighErr: true, wantLowErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.vg.Normal()
			if gotErr := err != nil; gotErr != tc.wantNormalErr {
				t.Errorf("[%v].Normal() got err %v, want err? %t", tc.vg, err, tc.wantNormalErr)
			}

			if !tc.wantNormalErr {
				gotf, err := strconv.ParseFloat(got, 64)
				if err != nil {
					t.Fatalf("ParseFloat(%q, 64) failed with err %v", got, err)
				}
				if gotf != tc.wantNormal {
					t.Errorf("[%v].Normal() = %q, want %f", tc.vg, got, tc.wantNormal)
				}
			}

			_, err = tc.vg.AbnormalHigh()
			if gotErr := err != nil; gotErr != tc.wantHighErr {
				t.Errorf("[%v].AbnormalHigh() got err %v, want err? %t", tc.vg, err, tc.wantHighErr)
			}
			_, err = tc.vg.AbnormalLow()
			if gotErr := err != nil; gotErr != tc.wantLowErr {
				t.Errorf("[%v].AbnormalLow() got err %v, want err? %t", tc.vg, err, tc.wantLowErr)
			}
		})
	}
}

func TestValueGenerator_Random(t *testing.T) {
	cases := []struct {
		name       string
		inRange    string
		randomType string
		wantErr    bool
		wantRange  valRange
	}{
		{name: "normal from range", inRange: "12 - 24", randomType: constants.NormalValue, wantRange: valRange{from: 12, to: 24}},
		{name: "high from range", inRange: "12 - 24", randomType: constants.AbnormalHigh, wantRange: valRange{from: 24, to: 240}},
		{name: "low from range", inRange: "12 - 24", randomType: constants.AbnormalLow, wantRange: valRange{from: 0, to: 12}},
		{name: "normal from open range", inRange: "[>12]", randomType: constants.NormalValue, wantRange: valRange{from: 12, to: 120}},
		{name: "low from open range", inRange: "[>12]", randomType: constants.AbnormalLow, wantRange: valRange{from: 0, to: 12}},
		{name: "invalid random type", inRange: "12 - 24", randomType: "any", wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			vg, err := ValueGeneratorFromRange(tc.inRange)
			if err != nil {
				t.Fatalf("ValueGeneratorFromRange(%s) failed with err %v", tc.inRange, err)
			}

			got, err := vg.Random(tc.randomType)
			if gotErr := err != nil; gotErr != tc.wantErr {
				t.Fatalf("Random(%s) got err %v, want err? %t", tc.randomType, err, tc.wantErr)
			}
			if tc.wantErr || err != nil {
				return
			}

			gotf, err := strconv.ParseFloat(got, 64)
			if err != nil {
				t.Fatalf("ParseFloat(%q, 64) failed with err %v", got, err)
			}
			if gotf <= tc.wantRange.from || gotf >= tc.wantRange.to {
				t.Errorf("Random(%s) = %q, want in range (%f, %f)", tc.randomType, got, tc.wantRange.from, tc.wantRange.to)
			}
		})
	}
}

func TestValueGenerator_IsNormal_IsHigh_IsLow(t *testing.T) {
	cases := []struct {
		name     string
		inRange  string
		val      float64
		isHigh   bool
		isLow    bool
		isNormal bool
	}{
		{name: "normal", inRange: "1.2 - 5.5", val: 2, isNormal: true, isHigh: false, isLow: false},
		{name: "normal - equal to left", inRange: "1.2 - 5.5", val: 1.2, isNormal: true, isHigh: false, isLow: false},
		{name: "normal - equal to right", inRange: "1.2 - 5.5", val: 5.5, isNormal: true, isHigh: false, isLow: false},
		{name: "high", inRange: "1.2 - 5.5", val: 5.6, isNormal: false, isHigh: true, isLow: false},
		{name: "low", inRange: "1.2 - 5.5", val: 1.1, isNormal: false, isHigh: false, isLow: true},
		{name: "normal", inRange: "-10.2 - -5.5", val: -6, isNormal: true, isHigh: false, isLow: false},
		{name: "high", inRange: "-10.2 - -5.5", val: -5.4, isNormal: false, isHigh: true, isLow: false},
		{name: "low", inRange: "-10.2 - -5.5", val: -10.5, isNormal: false, isHigh: false, isLow: true},
		{name: "open left range - normal", inRange: ">1.2", val: 5.6, isNormal: true, isHigh: false, isLow: false},
		{name: "open left range - low", inRange: ">1.2", val: 1.1, isNormal: false, isHigh: false, isLow: true},
		{name: "open right range - normal", inRange: "<4.5", val: 0.2, isNormal: true, isHigh: false, isLow: false},
		{name: "open right range - high", inRange: "<4.5", val: 5.6, isNormal: false, isHigh: true, isLow: false},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s-%s", tc.name, tc.inRange), func(t *testing.T) {
			vg, err := ValueGeneratorFromRange(tc.inRange)
			if err != nil {
				t.Fatalf("ValueGeneratorFromRange(%s) failed with err %v", tc.inRange, err)
			}

			if got, want := vg.IsNormal(tc.val), tc.isNormal; got != want {
				t.Errorf("IsNormal(%v)=%v, want %v", tc.val, got, want)
			}
			if got, want := vg.IsHigh(tc.val), tc.isHigh; got != want {
				t.Errorf("IsHigh(%v)=%v, want %v", tc.val, got, want)
			}
			if got, want := vg.IsLow(tc.val), tc.isLow; got != want {
				t.Errorf("IsLow(%v)=%v, want %v", tc.val, got, want)
			}
		})
	}
}
