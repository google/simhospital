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

package names

import (
	"math"
	"testing"

	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/gender"
)

func TestPrefixMiddleNameSuffixDegreeSurname(t *testing.T) {
	malePrefix := []string{"Mr", "Master"}
	femalePrefix := []string{"Ms", "Mrs", "Miss"}
	suffixes := []string{"dr", "prof", "sir"}
	degrees := []string{"B.A.", "Ph.D."}
	surnames := []string{"Smith", "Bond", "Forster"}
	firstNames := &config.FirstNamesByCensus{
		Girls: &config.Names{
			All: []string{"Mary", "Margaret", "Susan"},
		},
		Boys: &config.Names{
			All: []string{"William", "John", "David"},
		},
	}

	cases := []struct {
		name                 string
		gender               gender.Internal
		middlenamePercentage int
		suffixPercentage     int
		degreePercantage     int
		funcName             string
		funcCall             func(g Generator, gen gender.Internal) string
		wantOneOf            []string
		wantPercentage       int
	}{
		{
			name:           "Male prefix",
			gender:         gender.Male,
			funcName:       "Prefix",
			funcCall:       func(g Generator, gen gender.Internal) string { return g.Prefix(gen) },
			wantOneOf:      []string{"Mr", "Master"},
			wantPercentage: 100,
		},
		{
			name:           "Female prefix",
			gender:         gender.Female,
			funcName:       "Prefix",
			funcCall:       func(g Generator, gen gender.Internal) string { return g.Prefix(gen) },
			wantOneOf:      []string{"Ms", "Mrs", "Miss"},
			wantPercentage: 100,
		},
		{
			name:           "Other prefix",
			gender:         gender.Unknown,
			funcName:       "Prefix",
			funcCall:       func(g Generator, gen gender.Internal) string { return g.Prefix(gen) },
			wantOneOf:      []string{},
			wantPercentage: 0,
		},
		{
			name:                 "Male middlename",
			gender:               gender.Male,
			middlenamePercentage: 50,
			funcName:             "MiddleName",
			funcCall:             func(g Generator, gen gender.Internal) string { return g.MiddleName(gen) },
			wantOneOf:            []string{"William", "John", "David"},
			wantPercentage:       50,
		},
		{
			name:                 "Male middlename always populated",
			gender:               gender.Male,
			middlenamePercentage: 100,
			funcName:             "MiddleName",
			funcCall:             func(g Generator, gen gender.Internal) string { return g.MiddleName(gen) },
			wantOneOf:            []string{"William", "John", "David"},
			wantPercentage:       100,
		},
		{
			name:                 "Male middlename never populated",
			gender:               gender.Male,
			middlenamePercentage: 0,
			funcName:             "MiddleName",
			funcCall:             func(g Generator, gen gender.Internal) string { return g.MiddleName(gen) },
			wantOneOf:            []string{},
			wantPercentage:       0,
		},
		{
			name:                 "Female middlename",
			gender:               gender.Female,
			middlenamePercentage: 50,
			funcName:             "MiddleName",
			funcCall:             func(g Generator, gen gender.Internal) string { return g.MiddleName(gen) },
			wantOneOf:            []string{"Mary", "Margaret", "Susan"},
			wantPercentage:       50,
		},
		{
			name:                 "Other middlename",
			gender:               gender.Unknown,
			middlenamePercentage: 50,
			funcName:             "MiddleName",
			funcCall:             func(g Generator, gen gender.Internal) string { return g.MiddleName(gen) },
			wantOneOf:            []string{},
			wantPercentage:       0,
		},
		{
			name:             "Suffixes 80% probability",
			suffixPercentage: 80,
			funcName:         "Suffix",
			funcCall:         func(g Generator, gen gender.Internal) string { return g.Suffix() },
			wantOneOf:        suffixes,
			wantPercentage:   80,
		},
		{
			name:             "Degrees 20% probability",
			degreePercantage: 20,
			funcName:         "Degree",
			funcCall:         func(g Generator, gen gender.Internal) string { return g.Degree() },
			wantOneOf:        degrees,
			wantPercentage:   20,
		},
		{
			name:           "Surname",
			funcName:       "Surname",
			funcCall:       func(g Generator, gen gender.Internal) string { return g.Surname() },
			wantOneOf:      surnames,
			wantPercentage: 100,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data := &config.Data{
				PatientName: config.PatientName{
					MalePrefixes:         malePrefix,
					FemalePrefixes:       femalePrefix,
					MiddlenamePercentage: tc.middlenamePercentage,
					Suffixes:             suffixes,
					SuffixPercentage:     tc.suffixPercentage,
					Degrees:              degrees,
					DegreePercentage:     tc.degreePercantage,
				},
				FirstNames: firstNames,
				Surnames:   surnames,
			}
			g := Generator{Data: data}

			runs := 1000
			gotAll := map[string]int{}
			for i := 0; i < runs; i++ {
				got := tc.funcCall(g, tc.gender)
				if got != "" && !contains(got, tc.wantOneOf) {
					t.Errorf("%s(%v)=%v, want one of %v or empty string", tc.funcName, tc.gender, got, tc.wantOneOf)
				}
				gotAll[got]++
			}

			wantNonEmpty := float64(runs) * (float64(tc.wantPercentage) / 100.0)
			delta := float64(wantNonEmpty) / 5.0
			wantFreq := wantNonEmpty / float64(len(tc.wantOneOf))
			for _, item := range tc.wantOneOf {
				if gotFreq := gotAll[item]; math.Abs(float64(gotFreq)-wantFreq) > delta {
					t.Errorf("gotAll[%s] = %d, want within %v of %f", item, gotFreq, delta, wantFreq)
				}
			}
			wantEmpty := float64(runs) - wantNonEmpty
			deltaEmpty := float64(wantEmpty) / 5.0
			if gotEmptyFreq := gotAll[""]; math.Abs(float64(gotEmptyFreq)-wantEmpty) > deltaEmpty {
				t.Errorf("gotAll[%q] = %d, want within %v of %f", "", gotEmptyFreq, deltaEmpty, wantEmpty)
			}
		})
	}
}

func TestFirstName(t *testing.T) {
	girlsUntil1974 := []string{"Sarah", "Clair", "Nicola"}
	girls1975To1984 := []string{"Mary", "Laura", "Gemma"}
	girlsFrom1985 := []string{"Rebeca", "Lauren", "Jessica"}

	boysUntil1974 := []string{"Paul", "Marc", "David"}
	boys1975To1984 := []string{"Christopher", "James", "David"}
	boysFrom1985 := []string{"Thomas", "James", "Jack"}

	data := &config.Data{
		FirstNames: &config.FirstNamesByCensus{
			Girls: &config.Names{
				ByYear: map[int][]string{
					1974: girlsUntil1974,
					1984: girls1975To1984,
					1994: girlsFrom1985,
				},
				MinYear: 1974,
				MaxYear: 1994,
			},
			Boys: &config.Names{
				ByYear: map[int][]string{
					1974: boysUntil1974,
					1984: boys1975To1984,
					1994: boysFrom1985,
				},
				MinYear: 1974,
				MaxYear: 1994,
			},
		},
	}
	g := Generator{Data: data}

	cases := []struct {
		name      string
		gender    gender.Internal
		year      int
		wantOneOf []string
	}{
		{
			name:      "Male name 1984",
			gender:    gender.Male,
			year:      1984,
			wantOneOf: boys1975To1984,
		},
		{
			name:      "Male name 1985",
			gender:    gender.Male,
			year:      1985,
			wantOneOf: boysFrom1985,
		},
		{
			name:      "Male name 1974 - min year specified",
			gender:    gender.Male,
			year:      1974,
			wantOneOf: boysUntil1974,
		},
		{
			name:      "Male name 1994 - max year specified",
			gender:    gender.Male,
			year:      1994,
			wantOneOf: boysFrom1985,
		},
		{
			name:      "Male name 1966 - below min year",
			gender:    gender.Male,
			year:      1966,
			wantOneOf: boysUntil1974,
		},
		{
			name:      "Male name 2012 - above max year",
			gender:    gender.Male,
			year:      2012,
			wantOneOf: boysFrom1985,
		},
		{
			name:      "Female name 1983",
			gender:    gender.Female,
			year:      1983,
			wantOneOf: girls1975To1984,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			runs := 100
			gotAll := map[string]int{}
			for i := 0; i < runs; i++ {
				got := g.FirstName(tc.gender, tc.year)
				if !contains(got, tc.wantOneOf) {
					t.Errorf("FirstName(%v, %d)=%v, want one of %v", tc.gender, tc.year, got, tc.wantOneOf)
				}
				gotAll[got]++
			}

			delta := float64(runs) / 5.0
			wantFreq := float64(runs) / float64(len(tc.wantOneOf))
			for _, item := range tc.wantOneOf {
				if gotFreq := gotAll[item]; math.Abs(float64(gotFreq)-wantFreq) > delta {
					t.Errorf("gotAll[%s] = %d, want within %v of %f", item, gotFreq, delta, wantFreq)
				}
			}
		})
	}
}

func contains(str string, strs []string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}
