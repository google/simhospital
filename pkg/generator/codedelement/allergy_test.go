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

package codedelement

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/message"
	"github.com/google/simhospital/pkg/test"
	"github.com/google/simhospital/pkg/test/testclock"
	"github.com/google/simhospital/pkg/test/testdate"
	"github.com/google/simhospital/pkg/test/testwrite"
)

var (
	defaultDate = time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)
	reactions   = []string{"Sneezing", "Itchy nose", "Runny nose"}
)

func TestAllergyGenerator_RandomNoItems(t *testing.T) {
	fName := testwrite.BytesToFile(t, []byte(``))

	configHL7, err := config.LoadHL7Config(test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	f := test.DataFiles[test.Test]
	f.Allergies = fName
	data, err := config.LoadData(f, configHL7)
	if err != nil {
		t.Fatalf("LoadData(%+v, %+v) failed with %v", f, configHL7, err)
	}
	g := NewAllergyGenerator(configHL7, data, testclock.New(defaultDate), &testdate.Generator{})
	if r := g.Random(); r != nil {
		t.Errorf("NewAllergyGenerator().Random() = %v, want <nil>", r)
	}
}

func TestAllergyGenerator_Random(t *testing.T) {
	rand.Seed(1)
	fName := testwrite.BytesToFile(t, []byte(`
J30.1,Allergy1,59
J45.0,Allergy2,2556
J23.1,Allergy3,200
T78.1,Allergy4,23
`))

	codeToDesc := map[string]string{"J30.1": "Allergy1", "J45.0": "Allergy2", "J23.1": "Allergy3", "T78.1": "Allergy4"}

	configHL7, err := config.LoadHL7Config(test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	f := test.DataFiles[test.Test]
	f.Allergies = fName
	data, err := config.LoadData(f, configHL7)
	if err != nil {
		t.Fatalf("LoadData(%+v, %+v) failed with %v", f, configHL7, err)
	}
	g := NewAllergyGenerator(configHL7, data, testclock.New(defaultDate), &testdate.Generator{})
	if len(g.WeightedValues) != 4 {
		t.Fatalf("len(allergies.WeightedValues) = %d, want %d", len(g.WeightedValues), 4)
	}

	runs := 1000
	generated := float64(0)
	typesCount := make(map[string]int)
	severitiesCount := make(map[string]int)
	reactionsCount := make(map[string]int)
	fieldExpectations := map[string]struct {
		field func(*message.Allergy) string
		want  []string
		count map[string]int
	}{
		"Type":     {func(a *message.Allergy) string { return a.Type }, configHL7.Allergy.Types, typesCount},
		"Severity": {func(a *message.Allergy) string { return a.Severity }, configHL7.Allergy.Severities, severitiesCount},
		"Reaction": {func(a *message.Allergy) string { return a.Reaction }, reactions, reactionsCount},
	}
	for i := 0; i < runs; i++ {
		allergies := g.GenerateRandomDistinctAllergies()
		for _, a := range allergies {
			generated++
			if _, ok := codeToDesc[a.Description.ID]; !ok {
				t.Errorf("GenerateRandomDistinctAllergies().Description.ID = %s, want one of %v",
					a.Description.ID, codeToDesc)
			}
			if got, want := a.Description.Text, codeToDesc[a.Description.ID]; got != want {
				t.Errorf("GenerateRandomDistinctAllergies().Description.Text got %v, want %v", got, want)
			}
			if got, want := a.IdentificationDateTime, message.NewValidTime(defaultDate); got != want {
				t.Errorf("GenerateRandomDistinctAllergies().IdentificationDateTime got %v, want %v", got, want)
			}
			for k, v := range fieldExpectations {
				got := v.field(a)
				if !contains(v.want, got) {
					t.Errorf("%s=%q, want one of %v", k, got, v.want)
				}
				v.count[got]++
			}
		}
	}

	// Allow an error of 20% of the number of generated.
	delta := generated / 5
	for k, got := range typesCount {
		if want := generated / float64(len(configHL7.Allergy.Types)); math.Abs(float64(got)-float64(want)) >= delta {
			t.Errorf("typesCount[%q] = %d, want within %.1f of %d", k, got, delta, want)
		}
	}
	for k, got := range severitiesCount {
		if want := generated / float64(len(severitiesCount)); math.Abs(float64(got)-float64(want)) >= delta {
			t.Errorf("severitiesCount[%q] = %d, want within %.1f of %d", k, got, delta, want)
		}
	}
	for k, got := range reactionsCount {
		if want := generated / float64(len(reactions)); math.Abs(float64(got)-float64(want)) >= delta {
			t.Errorf("reactionsCount[%q] = %d, want within %.1f of %d", k, got, delta, want)
		}
	}
}

func contains(strings []string, target string) bool {
	for _, s := range strings {
		if s == target {
			return true
		}
	}
	return false
}
