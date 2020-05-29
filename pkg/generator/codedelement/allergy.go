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
	"math/rand"

	"github.com/google/simhospital/pkg/clock"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/pathway"
)

// AllergyGenerator provides functionality to generate an allergy.
type AllergyGenerator struct {
	*Generator
	severities   []string
	reactions    []string
	percentage   int
	maxAllergies int
}

// DeriveIdentificationDateTime returns an absolute or invalid NullTime value for an
// allergy's IdentificationDateTime.
func (g *AllergyGenerator) DeriveIdentificationDateTime(a pathway.Allergy) ir.NullTime {
	return g.nullTimeOrRandom(a.IdentificationDateTime)
}

// DeriveCodingSystem uses the coding system configured for the allergy, or defaults to the
// allergy coding system in the default message configuration.
func (g *AllergyGenerator) DeriveCodingSystem(config config.HL7Allergy, a pathway.Allergy) string {
	if a.CodingSystem != "" {
		return a.CodingSystem
	}
	return config.CodingSystem
}

// randomSeverity returns a random severity value, where each value has an equal probability to be selected.
func (g *AllergyGenerator) randomSeverity() string {
	return g.severities[rand.Intn(len(g.severities))]
}

// randomReaction returns a random reaction value, where each value has an equal probability to be selected.
func (g *AllergyGenerator) randomReaction() string {
	return g.reactions[rand.Intn(len(g.reactions))]
}

// randomIdentificationDateTime returns a random identification datetime.
func (g *AllergyGenerator) randomIdentificationDateTime() ir.NullTime {
	return g.nullTimeOrRandom(nil)
}

// GenerateRandomDistinctAllergies generates a list of allergies.
// The list will have at least one item with probability percentage.
// After that, the final number of items is picked randomly between 1 to maxAllergies (both inclusive).
func (g *AllergyGenerator) GenerateRandomDistinctAllergies() []*ir.Allergy {
	var generatedAllergies []*ir.Allergy
	ra := rand.Intn(100)
	if ra >= g.percentage {
		return generatedAllergies
	}
	allergyCount := rand.Intn(g.maxAllergies) + 1
	selectedCodes := map[string]bool{}
	for len(generatedAllergies) < allergyCount {
		a := g.Random()
		if !selectedCodes[a.ID] {
			selectedCodes[a.ID] = true

			generatedAllergies = append(generatedAllergies, &ir.Allergy{
				Type:                   g.RandomType(),
				Description:            *a,
				Severity:               g.randomSeverity(),
				Reaction:               g.randomReaction(),
				IdentificationDateTime: g.randomIdentificationDateTime(),
			})
		}
	}
	return generatedAllergies
}

// NewAllergyGenerator creates a new Generator with the allergies from the given configurations.
func NewAllergyGenerator(hc *config.HL7Config, d *config.Data, c clock.Clock, dg DateGenerator) *AllergyGenerator {
	return &AllergyGenerator{
		Generator:    newGenerator(d.Allergies, hc.Allergy.Types, c, dg),
		severities:   hc.Allergy.Severities,
		reactions:    d.Allergy.Reactions,
		percentage:   d.Allergy.Percentage,
		maxAllergies: d.Allergy.MaximumAllergies,
	}
}
