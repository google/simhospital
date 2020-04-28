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

package address

import (
	"math"
	"math/rand"
	"strings"
	"testing"

	"github.com/google/simhospital/pkg/config"
)

func TestRandom(t *testing.T) {
	rand.Seed(1)

	cities := []string{"London", "Cambridge"}
	types := []string{"HOME"}
	streets := []string{"Road", "Street"}
	country := "GBR"
	nouns := []string{"aardvark", "abacus", "abbey"}
	wantNouns := []string{"Aardvark", "Abacus", "Abbey"}

	g := Generator{
		Address:           config.Address{Cities: cities, Country: country, Streets: streets, Types: types},
		Nouns:             nouns,
		PostcodeGenerator: &UKPostcode{},
	}

	gotSecondLine := 0
	runs := 100

	for i := 0; i < runs; i++ {
		got := g.Random()
		if got == nil {
			t.Fatal("g.Random() = <nil>, want non-nil value")
		}
		if got.FirstLine == "" {
			t.Error(`g.Random().FirstLine = "", want nonempty string`)
		}
		if got.SecondLine != "" {
			gotSecondLine++
		}
		if got.PostalCode == "" {
			t.Error("g.Random().PostalCode is empty, want non empty")
		}
		if !contains(cities, got.City) {
			t.Errorf("g.Random().City = %q, want one of data config cities %v", got.City, cities)
		}
		if got, want := got.Country, country; got != want {
			t.Errorf("g.Random().Country = %q, want %q", got, want)
		}
		if !contains(types, got.Type) {
			t.Errorf("g.Random().Type = %q, want one of data config address types %v", got.Type, types)
		}
		if !containsSubstring(streets, got.FirstLine) && !containsSubstring(streets, got.SecondLine) {
			t.Errorf("g.Random().FirstLine = %q, g.Random().SecondLine = %q, want one of them to contain %v", got.FirstLine, got.SecondLine, streets)
		}
		if !containsSubstring(wantNouns, got.FirstLine) && !containsSubstring(wantNouns, got.SecondLine) {
			t.Errorf("g.Random().FirstLine = %q, g.Random().SecondLine = %q, want one of them to contain %v", got.FirstLine, got.SecondLine, wantNouns)
		}
	}

	// Second Line should be populated 50% of the time.
	delta := float64(runs) / 5.0
	if want := runs / 2.0; math.Abs(float64(gotSecondLine)-float64(want)) >= delta {
		t.Errorf("gotSecondLine=%d, want within %.1f of %d", gotSecondLine, delta, want)
	}
}

func contains(set []string, target string) bool {
	for _, s := range set {
		if s == target {
			return true
		}
	}
	return false
}

func containsSubstring(set []string, target string) bool {
	for _, s := range set {
		if strings.Contains(target, s) {
			return true
		}
	}
	return false
}
