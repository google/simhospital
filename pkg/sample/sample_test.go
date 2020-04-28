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

package sample

import (
	"math/rand"
	"testing"
)

func TestDiscreteDistribution_Random_IsFromDistribution(t *testing.T) {
	testResults := []WeightedValue{
		{Value: "Garnet", Frequency: 10},
		{Value: "Amethyst", Frequency: 5},
		{Value: "Pearl", Frequency: 17},
		{Value: "Steven", Frequency: 8},
	}
	d := DiscreteDistribution{WeightedValues: testResults}
	want := map[string]bool{
		"Garnet":   true,
		"Amethyst": true,
		"Pearl":    true,
		"Steven":   true,
	}
	for i := 1; i <= 10; i++ {
		got := d.Random()
		if !want[got.(string)] {
			t.Errorf("d.Random() got %q; want a member of: %v", got, want)
		}
	}
}

func TestDiscreteDistribution_Random_IsProportionateToDistribution(t *testing.T) {
	rand.Seed(1)
	testResults := []WeightedValue{
		{Value: "Red", Frequency: 1},
		{Value: "Orange", Frequency: 1},
		{Value: "Yellow", Frequency: 1},
		{Value: "Green", Frequency: 1},
		{Value: "Blue", Frequency: 1},
		{Value: "Violet", Frequency: 15},
	}
	d := DiscreteDistribution{WeightedValues: testResults}
	want := map[string]bool{
		"Red":    true,
		"Orange": true,
		"Yellow": true,
		"Green":  true,
		"Blue":   true,
		"Violet": true,
	}

	gotCount := make(map[string]float64)
	runs := 10000
	for i := 1; i <= runs; i++ {
		got := d.Random()
		gotCount[got.(string)]++
		if !want[got.(string)] {
			t.Errorf("d.Random() got %q; want a member of: %v", got, want)
		}
	}
	// Allow an error of 1% of the number of runs.
	runsFl := float64(runs)
	delta := runsFl / 100

	for w := range want {
		if w == "Violet" {
			// We expect "Violet" to be chosen 15 times every 20 runs, so about 75% of the time.
			if min, max, got := 0.75*runsFl-delta, 0.75*runsFl+delta, gotCount["Violet"]; got < min || got > max {
				t.Errorf("count of element Violet: got %f; want between [%f, %f]", got, min, max)
			}
		} else {
			// The rest of them are chosen once every 20 runs, so 5% of the time.
			if min, max, got := 0.05*runsFl-delta, 0.05*runsFl+delta, gotCount[w]; got < min || got > max {
				t.Errorf("count of element %q: got %f; want between [%f, %f]", w, got, min, max)
			}
		}
	}
}

func TestDiscreteDistribution_Random_NoElements(t *testing.T) {
	var testResults []WeightedValue
	d := DiscreteDistribution{WeightedValues: testResults}
	for i := 1; i <= 10; i++ {
		if got := d.Random(); got != nil {
			t.Errorf("d.Random() got %v, want <nil>.", got)
		}
	}
}
