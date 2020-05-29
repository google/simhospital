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

// Package person provides functionality to generate a person.
// This includes:
// - demographic information, ie: name, surname, gender, date of birth, ethnicity, etc.
// - address, telephone number,
// - NHS and MRN numbers.
package person

import (
	"fmt"
	"math/rand"

	"github.com/google/simhospital/pkg/clock"
	"github.com/google/simhospital/pkg/gender"
	"github.com/google/simhospital/pkg/generator/id"
	"github.com/google/simhospital/pkg/generator/names"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/pathway"
)

// AddressGenerator is an interface to generate addresses.
type AddressGenerator interface {
	Random() *ir.Address
}

// Generator is a generator of person.
type Generator struct {
	Clock              clock.Clock
	NameGenerator      *names.Generator
	GenderConvertor    gender.Convertor
	EthnicityGenerator EthnicityGenerator
	AddressGenerator   AddressGenerator
	MRNGenerator       id.Generator
}

// NewPerson returns a new person based on pathway.Person.
func (g Generator) NewPerson(pathwayPerson *pathway.Person) *ir.Person {

	// We first set the fields that can't be set from the pathway.
	person := &ir.Person{
		Suffix:      g.NameGenerator.Suffix(),
		Degree:      g.NameGenerator.Degree(),
		Ethnicity:   g.EthnicityGenerator.Random(),
		PhoneNumber: g.phoneNumber(),
	}
	g.UpdatePersonFromPathway(person, pathwayPerson)
	return person
}

// UpdatePersonFromPathway updates a person with information from a pathway. Calling this method
// populates all fields of a Person, if they were not already set.
// Fields that are set in the pathway's person always override the original person's.
// Fields that are set in the original person, but not in the pathway, are kept as they are.
// Fields that are not set in any are generated randomly.
func (g Generator) UpdatePersonFromPathway(person *ir.Person, pathwayPerson *pathway.Person) {
	if pathwayPerson == nil {
		pathwayPerson = &pathway.Person{}
	}

	// Determine values for dob and gender first, because those values are needed
	// in order to generate other variables like FirstName, and Birth.

	// For dob, the pathway person may have either Age or DateOfBirth, or none of them set.
	// If person.Birth is already set to a valid date and neither DateOfBirth
	// nor Age are explicitly set in the pathway, don't change the person.Birth.
	switch {
	case pathwayPerson.DateOfBirth != nil:
		person.Birth = ir.NewValidTime(*pathwayPerson.DateOfBirth)
	case pathwayPerson.Age != nil:
		person.Birth = ir.NewValidTime(pathwayPerson.Age.Birthdate(g.Clock))
	case !person.Birth.Valid:
		person.Birth = ir.NewValidTime(pathway.RandomBirthdate(g.Clock))
	}

	// For gender, if there is no gender set in the pathway then we need to randomly generate one
	// for use in generating future variables.
	// If it is set in pathway, then we need to use to pathway gender to generate the other variables
	// or else there may be discrepancies (e.g. Male FirstNames for Female Person, etc.)
	originalGender := person.Gender
	person.Gender = g.chooseGenderValue(pathwayPerson.Gender, person.Gender)
	internalGender := g.GenderConvertor.HL7ToInternal(person.Gender)

	if person.Gender != originalGender {
		// The Middlename and Prefix depend on the gender, so if the gender changes we need to change
		// those too, otherwise we might have inconsistent names such as a person with gender Female
		// and a prefix "Mr.".
		person.MiddleName = g.NameGenerator.MiddleName(internalGender)
		person.Prefix = g.NameGenerator.Prefix(internalGender)
	}

	// Fields that are set in the pathway's person always override the original person's.
	person.FirstName = chooseValue(pathwayPerson.FirstName, g.NameGenerator.FirstName(internalGender,
		person.Birth.Year()), person.FirstName)
	person.Surname = chooseOptionalValue(pathwayPerson.Surname, g.NameGenerator.Surname(),
		person.Surname)
	person.Address = g.mergeAddressFromPathway(pathwayPerson.Address, person.Address)
	person.NHS = chooseValue(pathwayPerson.NHS, newNHSNumber(), person.NHS)
	person.MRN = chooseValueLazy(pathwayPerson.MRN, func() string { return g.MRNGenerator.NewID() }, person.MRN)
}

// mergeAddressFromPathway merges the given HL7 address with a pathway address, overriding only the
// fields that are set in the pathway.
// FirstLine and SecondLine are treated as a unit; if FirstLine is
// defined in the pathway, the SecondLine field is overridden along with it,
// no matter if it is empty.
// Fields set to RANDOM are overridden with a new random value.
// If AllRandom is enabled in the pathway, all fields are overridden with new random values.
func (g Generator) mergeAddressFromPathway(a *pathway.Address, p *ir.Address) *ir.Address {
	random := g.AddressGenerator.Random()
	if p == nil {
		p = random
	}
	if a == nil {
		return p
	}
	if a.AllRandom {
		return random
	}

	// The treatment of FirstLine and SecondLine is special because it would otherwise be impossible
	// to turn a 2-line address into a 1-line address. If first line is not set, second line is never
	// modified. If first line is set, so must be the second.
	p.FirstLine = chooseOptionalValue(a.FirstLine, random.FirstLine, p.FirstLine)
	if a.FirstLine.IsSet() {
		if a.SecondLine == "" {
			// We are turning into a 1-line address, don't randomly generate second line.
			p.SecondLine = ""
		} else {
			// Deal with random variables in the second line as necessary.
			p.SecondLine = chooseOptionalValue(a.SecondLine, random.SecondLine, p.SecondLine)
		}
	}

	p.City = chooseOptionalValue(a.City, random.City, p.City)
	p.PostalCode = chooseOptionalValue(a.Postcode, random.PostalCode, p.PostalCode)
	p.Country = chooseOptionalValue(a.Country, random.Country, p.Country)
	p.Type = chooseOptionalValue(a.Type, random.Type, p.Type)
	return p
}

// chooseOptionalValue returns an appropriate value based on whether "pathwayString", "original" and "random" are set.
func chooseOptionalValue(pathwayString pathway.OptionalRandomString, random string, original string) string {
	switch {
	case pathwayString.IsFixedValue():
		return pathwayString.String()
	case !pathwayString.IsSet() && original != "":
		return original
	default:
		return random
	}
}

// chooseValue returns an appropriate value based on whether "pathwayString" and "original" are set.
func chooseValue(pathwayString string, random string, original string) string {
	return chooseValueLazy(pathwayString, func() string { return random }, original)
}

// chooseValueLazy returns an appropriate value based on whether "pathwayString" and "original" are set.
// The value from valueContainer is only retrieved if needed.
func chooseValueLazy(pathwayString string, valueFunc func() string, original string) string {
	switch {
	case pathway.IsFixedValue(pathwayString):
		return pathwayString
	case original != "":
		return original
	default:
		return valueFunc()
	}
}

// chooseGenderValue returns an appropriate HL7 gender value based on whether "pathwayGender" and "originalHL7" are set.
// If neither of them are provided, returns a random HL7 gender.
func (g Generator) chooseGenderValue(pathwayGender pathway.Gender, originalHL7 string) string {
	switch {
	case pathway.IsFixedValue(string(pathwayGender)):
		return g.GenderConvertor.PathwayToHL7(pathwayGender)
	case originalHL7 != "":
		return originalHL7
	default:
		return g.GenderConvertor.InternalToHL7(gender.Random())
	}
}

func (g Generator) phoneNumber() string {
	if rand.Intn(2) == 0 {
		// London home phone number
		return fmt.Sprintf("020 %04d %04d", rand.Intn(10000), rand.Intn(10000))
	}
	// UK mobile number
	return fmt.Sprintf("07%d %04d %04d", rand.Intn(10), rand.Intn(10000), rand.Intn(10000))
}

// Return a newly minted NHS number that will pass validation rules. See:
// http://www.datadictionary.nhs.uk/version2/data_dictionary/data_field_notes/n/nhs_number_de.asp?shownav=0
func newNHSNumber() string {
	for {
		n := rand.Intn(1000000000) * 10
		a := n / 10
		check := 0
		for i := 0; i < 9; i++ {
			check += (a % 10) * (i + 2)
			a /= 10
		}
		check = 11 - (check % 11)
		if check == 11 {
			check = 0
		}
		if check != 10 {
			return fmt.Sprintf("%010d", n+check)
		}
	}
}
