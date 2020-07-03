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

package config

// CodeMapping contains mappings from FHIR to those defined in the HL7 configuration.
type CodeMapping struct {
	FHIR FHIRMapping
}

// FHIRMapping contains the mappings from FHIR value sets to those defined in the HL7 configuration.
type FHIRMapping struct {
	CodingSystems map[string]string `yaml:"coding_systems"`
	// AllergySeverities contains the mapping from FHIR allergy severities to those defined in the HL7 configuration.
	// Reference: https://www.hl7.org/fhir/valueset-reaction-event-severity.html
	AllergySeverities map[string][]string `yaml:"allergy_severities"`
	// AllergyTypes contains the mapping from FHIR allergy categories to those defined in the HL7 configuration.
	// Reference: https://www.hl7.org/fhir/valueset-allergy-intolerance-category.html
	AllergyTypes map[string][]string `yaml:"allergy_types"`
}
