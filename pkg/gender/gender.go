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

// Package gender contains the functionality to convert gender values.
package gender

import (
	"math/rand"

	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/pathway"

	cpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/codes_go_proto"
)

// Internal is an internal representation of the gender.
type Internal int

const (
	// Unknown is an unknown gender.
	Unknown Internal = iota
	// Male is a male gender.
	Male
	// Female is a female gender.
	Female
)

// Random generates a random gender from the options of Male or Female
// with equal probability.
func Random() Internal {
	switch rand.Intn(2) {
	case 0:
		return Male
	default:
		return Female
	}
}

// Convertor converts gender between pathway, HL7 and internal representations.
type Convertor struct {
	internalToHL7Mapping map[Internal]string
	hl7ToInternalMapping map[string]Internal
	hl7ToFHIRMapping     map[string]cpb.AdministrativeGenderCode_Value
	hl7Gender            config.Gender
}

// NewConvertor returns a new gender Convertor based on the HL7 config.
func NewConvertor(c *config.HL7Config) Convertor {
	return Convertor{
		internalToHL7Mapping: map[Internal]string{
			Male:    c.Gender.Male,
			Female:  c.Gender.Female,
			Unknown: c.Gender.Unknown,
		},
		hl7ToInternalMapping: map[string]Internal{
			c.Gender.Male:    Male,
			c.Gender.Female:  Female,
			c.Gender.Unknown: Unknown,
		},
		// Mapping taken from https://www.hl7.org/fhir/cm-administrative-gender-v2.html
		hl7ToFHIRMapping: map[string]cpb.AdministrativeGenderCode_Value{
			c.Gender.Male:    cpb.AdministrativeGenderCode_MALE,
			c.Gender.Female:  cpb.AdministrativeGenderCode_FEMALE,
			c.Gender.Unknown: cpb.AdministrativeGenderCode_UNKNOWN,
		},
		hl7Gender: c.Gender,
	}
}

// PathwayToHL7 converts pathway.Gender to HL7 gender representation.
// If pathway.Gender is not specified as Male or Female, return Unknown.
func (gc Convertor) PathwayToHL7(g pathway.Gender) string {
	switch g {
	case pathway.Male:
		return gc.hl7Gender.Male
	case pathway.Female:
		return gc.hl7Gender.Female
	default:
		return gc.hl7Gender.Unknown
	}
}

// InternalToHL7 returns the HL7 representation for the given Internal gender.
func (gc Convertor) InternalToHL7(gender Internal) string {
	return gc.internalToHL7Mapping[gender]
}

// HL7ToInternal returns the Internal representation for the given HL7 gender.
func (gc Convertor) HL7ToInternal(gender string) Internal {
	return gc.hl7ToInternalMapping[gender]
}

// HL7ToFHIR returns the FHIR representation for the given HL7 gender.
func (gc Convertor) HL7ToFHIR(gender string) cpb.AdministrativeGenderCode_Value {
	return gc.hl7ToFHIRMapping[gender]
}
