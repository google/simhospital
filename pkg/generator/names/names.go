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

// Package names provides functionality to generate names for humans.
package names

import (
	"math/rand"

	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/gender"
)

// Generator is a generator of names.
type Generator struct {
	Data *config.Data
}

// Prefix returns a random prefix based on the given gender.
func (g Generator) Prefix(gen gender.Internal) string {
	switch gen {
	case gender.Male:
		return random(g.Data.PatientName.MalePrefixes)
	case gender.Female:
		return random(g.Data.PatientName.FemalePrefixes)
	default:
		return ""
	}
}

// FirstName returns a random first name based on the given gender and the year the person was born.
func (g Generator) FirstName(gen gender.Internal, year int) string {
	switch gen {
	case gender.Male:
		return randomByYear(g.Data.FirstNames.Boys, year)
	case gender.Female:
		return randomByYear(g.Data.FirstNames.Girls, year)
	default:
		return ""
	}
}

// MiddleName returns a random middle name based on the given gender.
func (g Generator) MiddleName(gen gender.Internal) string {
	if rand.Intn(100) < g.Data.PatientName.MiddlenamePercentage {
		switch gen {
		case gender.Male:
			return randomName(g.Data.FirstNames.Boys)
		case gender.Female:
			return randomName(g.Data.FirstNames.Girls)
		}
	}
	return ""
}

// Suffix returns a random suffix.
func (g Generator) Suffix() string {
	return randomWithProb(g.Data.PatientName.Suffixes, g.Data.PatientName.SuffixPercentage)
}

// Degree returns a random degree.
func (g Generator) Degree() string {
	return randomWithProb(g.Data.PatientName.Degrees, g.Data.PatientName.DegreePercentage)
}

// Surname returns a random surname.
func (g Generator) Surname() string {
	return random(g.Data.Surnames)
}

// randomWithProb returns a random item from the slice with the probability p/100, where p is an int between [0, 100),
// or an empty string otherwise.
func randomWithProb(s []string, p int) string {
	if rand.Intn(100) < p {
		return random(s)
	}
	return ""
}

// random returns a random item from the slice.
func random(s []string) string {
	return s[rand.Intn(len(s))]
}

// randomByYear returns a random name from the set of Names which were popular among people born in a given year.
// Every name from the given by-year set is equally probable.
func randomByYear(n *config.Names, year int) string {
	if year > n.MaxYear {
		year = n.MaxYear
	}
	if year < n.MinYear {
		year = n.MinYear
	}

	// Find the first year present in the census that is greater or equal to the given year.
	censusYear := year
	for ; censusYear < n.MaxYear; censusYear++ {
		if _, ok := n.ByYear[censusYear]; ok {
			break
		}
	}
	return n.ByYear[censusYear][rand.Intn(len(n.ByYear[censusYear]))]
}

// randomName returns a random name. Each name has the same probability to be returned.
func randomName(n *config.Names) string {
	return n.All[rand.Intn(len(n.All))]
}
