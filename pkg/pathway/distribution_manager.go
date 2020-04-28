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

package pathway

import (
	"fmt"
	"math"
	"sort"

	"github.com/pkg/errors"

	"github.com/google/simhospital/pkg/sample"
)

const (
	// maxSignificantDigits is the maximum number of decimal digits allowed in the
	// 'percentage_of_patients' field in pathways.
	maxSignificantDigits = 3
	// defaultPercentage is the default 'percentage_of_patients' for pathways that don't specify it.
	defaultPercentage = 1
)

// DistributionManager manages a given distribution of pathways.
type DistributionManager struct {
	// pathways is a map of pathways indexed by their name.
	pathways map[string]Pathway
	// distribution is the distribution of pathways.
	distribution *sample.DiscreteDistribution
	// basic is a list of the names of pathways with at least one step and exactly one person.
	basic []string
	// pathwayNames is a list of the names of all pathways this manager contains.
	pathwayNames []string
}

// getPathway returns the pathway for the specified name.
// If the pathway does not exist, then it returns an error.
func (m DistributionManager) getPathway(pathwayName string) (Pathway, error) {
	pathway, ok := m.pathways[pathwayName]
	if !ok {
		return pathway, fmt.Errorf("pathwayName %s does not exist within manager.pathways", pathwayName)
	}
	return pathway.Runnable()
}

// GetPathway gets the pathway with the given name.
// If the name provided is not valid, it returns an error.
func (m DistributionManager) GetPathway(pathwayName string) (*Pathway, error) {
	pathway, err := m.getPathway(pathwayName)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get pathway with name %s", pathwayName)
	}
	return &pathway, nil
}

// NextPathway returns the next pathway to run.
// This is chosen based on the expected frequency for each pathway in the original pathway list.
func (m DistributionManager) NextPathway() (*Pathway, error) {
	nextPathwayName := m.distribution.Random().(string)
	return m.GetPathway(nextPathwayName)
}

// AllPathwayNames returns all the pathway names in this manager.
func (m DistributionManager) AllPathwayNames() []string {
	return m.pathwayNames
}

// NewDistributionManager creates a new DistributionManager with the given pathway map.
// All pathways are initialised.
func NewDistributionManager(pathways map[string]Pathway) (DistributionManager, error) {
	m := DistributionManager{
		pathways: pathways,
		basic:    []string{},
	}
	for k, v := range m.pathways {
		v.Init(k)
		if len(v.Pathway) > 0 && v.Persons.HasOnePerson() {
			m.basic = append(m.basic, k)
		}
		m.pathways[k] = v
		m.pathwayNames = append(m.pathwayNames, k)
	}

	m.distribution = &sample.DiscreteDistribution{WeightedValues: calculateDistribution(m.pathways)}
	printPathwayStats(m.pathways)
	return m, nil
}

func calculateDistribution(pathways map[string]Pathway) []sample.WeightedValue {
	var weighted []sample.WeightedValue
	accPercentage := 0.0
	// We'll later share the remaining percentage budget among the pathways without an explicit one.
	var noPercentage []string
	for k, v := range pathways {
		if v.Percentage == nil {
			noPercentage = append(noPercentage, k)
		} else if *v.Percentage > 0 {
			accPercentage, weighted = addPercentage(v.Percentage.Float(), accPercentage, k, weighted)
		}
	}

	// Share the remaining percentage_of_patients budget among the pathways without an explicit one.
	optimal := 100.0
	if len(noPercentage) > 0 {
		remaining := optimal - accPercentage
		perPathway := calculateBudgetPerPathway(remaining, len(noPercentage))
		log.Infof("Setting pathway frequency %v%% for %d pathways without explicit percentage_of_patients: %v", perPathway, len(noPercentage), noPercentage)
		for _, k := range noPercentage {
			accPercentage, weighted = addPercentage(perPathway, accPercentage, k, weighted)
		}
	}
	log.Infof("Accumulated percentage_of_patients is %.3f. The closer to 100, "+
		"the closer the actual distribution will be to the pathways' percentage_of_patients", accPercentage)
	return weighted
}

func addPercentage(f, acc float64, name string, weighted []sample.WeightedValue) (float64, []sample.WeightedValue) {
	// Our distribution expects an integer.
	frequency := f * math.Pow(10.0, float64(maxSignificantDigits))
	weighted = append(weighted, sample.WeightedValue{Value: name, Frequency: uint(frequency)})
	return acc + f, weighted
}

// calculateBudgetPerPathway splits the remaining percentage budget among n pathways and returns
// the percentage that should be set per pathway. If there's not enough budget to set relevant
// percentages, it returns a default percentage.
func calculateBudgetPerPathway(remaining float64, n int) float64 {
	perPathway := remaining / float64(n)
	if perPathway <= 0 {
		log.Warningf("Cannot split remaining percentage budget %.3f among %d pathways, invalid percentage per pathway: %v. Setting default percentage: %v%%",
			remaining, n, perPathway, defaultPercentage)
		perPathway = defaultPercentage
	}
	return round(perPathway)
}

func printPathwayStats(pathways map[string]Pathway) {
	var names []string
	disabled := map[string]bool{}
	for name, p := range pathways {
		names = append(names, name)
		if p.Percentage != nil && *p.Percentage == 0 {
			disabled[name] = true
		}
	}

	sort.Strings(names)

	log.Infof("Loaded %d pathways:", len(pathways))
	for _, name := range names {
		suffix := ""
		if disabled[name] {
			suffix = " (percentage set to 0; this pathway will not be run)"
		}
		log.Infof(" - %s%s", name, suffix)
	}
}
