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
	"context"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/test"
	"github.com/google/simhospital/pkg/test/testclock"
	"github.com/google/simhospital/pkg/test/testdate"
	"github.com/google/simhospital/pkg/test/testwrite"

	cpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/codes_go_proto"
)

var (
	defaultDate = time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)
	reactions   = []string{"Sneezing", "Itchy nose", "Runny nose"}
)

func TestAllergyGenerator_RandomNoItems(t *testing.T) {
	ctx := context.Background()
	fName := testwrite.BytesToFile(t, []byte(``))
	configHL7, err := config.LoadHL7Config(ctx, test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	f := test.DataFiles[test.Test]
	f.Allergies = fName
	data, err := config.LoadData(ctx, f, configHL7)
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

	ctx := context.Background()
	configHL7, err := config.LoadHL7Config(ctx, test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	f := test.DataFiles[test.Test]
	f.Allergies = fName
	data, err := config.LoadData(ctx, f, configHL7)
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
		field func(*ir.Allergy) string
		want  []string
		count map[string]int
	}{
		"Type":     {func(a *ir.Allergy) string { return a.Type }, configHL7.Allergy.Types, typesCount},
		"Severity": {func(a *ir.Allergy) string { return a.Severity }, configHL7.Allergy.Severities, severitiesCount},
		"Reaction": {func(a *ir.Allergy) string { return a.Reaction }, reactions, reactionsCount},
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
			if got, want := a.IdentificationDateTime, ir.NewValidTime(defaultDate); got != want {
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

func TestNewAllergyConvertor(t *testing.T) {
	tests := []struct {
		name      string
		hl7Config *config.HL7Config
		wantErr   bool
	}{{
		name: "Valid mapping",
		hl7Config: &config.HL7Config{
			Mapping: config.CodeMapping{
				FHIR: config.FHIRMapping{
					AllergySeverities: map[string][]string{
						"SEVERE": []string{"SV", "SEVERE"},
						"MILD":   []string{"MI", "MILD"},
					},
					AllergyTypes: map[string][]string{
						"FOOD": []string{"V1"},
					},
				},
			},
		},
	}, {
		name: "Invalid severity mapping",
		hl7Config: &config.HL7Config{
			Mapping: config.CodeMapping{
				FHIR: config.FHIRMapping{
					AllergySeverities: map[string][]string{
						"SEVERE":  []string{"SV", "SEVERE"},
						"INVALID": []string{"INVALID"},
					},
					AllergyTypes: map[string][]string{
						"FOOD": []string{"V1"},
					},
				},
			},
		},
		wantErr: true,
	}, {
		name: "Invalid type mapping",
		hl7Config: &config.HL7Config{
			Mapping: config.CodeMapping{
				FHIR: config.FHIRMapping{
					AllergySeverities: map[string][]string{
						"SEVERE": []string{"SV", "SEVERE"},
						"MILD":   []string{"MI", "MILD"},
					},
					AllergyTypes: map[string][]string{
						"INVALID": []string{"INVALID"},
					},
				},
			},
		},
		wantErr: true,
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// AllergyConvertor has no exported fields, so we skip the comparison of the convertor itself.
			_, err := NewAllergyConvertor(tc.hl7Config)
			if gotErr := err != nil; gotErr != tc.wantErr {
				t.Errorf("NewAllergyConvertor(%v) got err %v, want error=%t", tc.hl7Config, err, tc.wantErr)
			}
		})
	}
}

func TestAllergyConvertorSeverityHL7ToFHIR(t *testing.T) {
	hl7Config := &config.HL7Config{
		Mapping: config.CodeMapping{
			FHIR: config.FHIRMapping{
				AllergySeverities: map[string][]string{
					"severe":   []string{"SV", "SEVERE"},
					"moderate": []string{"MO", "MODERATE"},
					"mild":     []string{"MI", "MILD"},
				},
			},
		},
	}

	wantMapping := map[string]cpb.AllergyIntoleranceSeverityCode_Value{
		"":         cpb.AllergyIntoleranceSeverityCode_INVALID_UNINITIALIZED,
		"UNKNOWN":  cpb.AllergyIntoleranceSeverityCode_INVALID_UNINITIALIZED,
		"MILD":     cpb.AllergyIntoleranceSeverityCode_MILD,
		"MI":       cpb.AllergyIntoleranceSeverityCode_MILD,
		"MODERATE": cpb.AllergyIntoleranceSeverityCode_MODERATE,
		"MO":       cpb.AllergyIntoleranceSeverityCode_MODERATE,
		"SEVERE":   cpb.AllergyIntoleranceSeverityCode_SEVERE,
		"SV":       cpb.AllergyIntoleranceSeverityCode_SEVERE,
	}
	c, err := NewAllergyConvertor(hl7Config)
	if err != nil {
		t.Fatalf("NewAllergyConvertor(%v) failed with %v", hl7Config, err)
	}

	for k, v := range wantMapping {
		t.Run(fmt.Sprintf("%v-%v", k, v), func(t *testing.T) {
			if got, want := c.SeverityHL7ToFHIR(k), v; got != want {
				t.Errorf("c.SeverityHL7ToFHIR(%v)=%v, want %v", k, got, want)
			}
		})
	}
}

func TestAllergyConvertorTypeHL7ToFHIR(t *testing.T) {
	hl7Config := &config.HL7Config{
		Mapping: config.CodeMapping{
			FHIR: config.FHIRMapping{
				AllergyTypes: map[string][]string{
					"food":        []string{"FA", "FOOD"},
					"medication":  []string{"MA", "MEDICATION"},
					"environment": []string{"EA", "ENVIRONMENT"},
					"biologic":    []string{"BIOLOGIC"},
				},
			},
		},
	}

	wantMapping := map[string]cpb.AllergyIntoleranceCategoryCode_Value{
		"":            cpb.AllergyIntoleranceCategoryCode_INVALID_UNINITIALIZED,
		"UNKNOWN":     cpb.AllergyIntoleranceCategoryCode_INVALID_UNINITIALIZED,
		"FOOD":        cpb.AllergyIntoleranceCategoryCode_FOOD,
		"FA":          cpb.AllergyIntoleranceCategoryCode_FOOD,
		"MEDICATION":  cpb.AllergyIntoleranceCategoryCode_MEDICATION,
		"MA":          cpb.AllergyIntoleranceCategoryCode_MEDICATION,
		"ENVIRONMENT": cpb.AllergyIntoleranceCategoryCode_ENVIRONMENT,
		"EA":          cpb.AllergyIntoleranceCategoryCode_ENVIRONMENT,
		"BIOLOGIC":    cpb.AllergyIntoleranceCategoryCode_BIOLOGIC,
	}
	c, err := NewAllergyConvertor(hl7Config)
	if err != nil {
		t.Fatalf("NewAllergyConvertor(%v) failed with %v", hl7Config, err)
	}

	for k, v := range wantMapping {
		t.Run(fmt.Sprintf("%v-%v", k, v), func(t *testing.T) {
			if got, want := c.TypeHL7ToFHIR(k), v; got != want {
				t.Errorf("c.TypeHL7ToFHIR(%v)=%v, want %v", k, got, want)
			}
		})
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
