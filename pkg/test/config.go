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

// Package test contains functionality to define test configs.
package test

import (
	"os"
	"path"

	"github.com/google/simhospital/pkg/config"
)

// ConfigType represents the type of configuration, eg: test, prod, etc.
type ConfigType string

const (
	// Test is the test configuration.
	Test = ConfigType("TEST")
	// Prod is the prod configuration.
	Prod = ConfigType("PROD")
)

var (
	base = path.Join(currentDir())

	// testConfigDir is the directory for test files for testing.
	testConfigDir = path.Join(base, "pkg", "test", "data")
	// prodConfigDir is the directory for prod files used during tests.
	// Some tests load the prod values to make sure the files are correct.
	prodConfigDir = path.Join(base, "configs")

	// AllergiesConfigTest is the path to the allergies config file for testing.
	AllergiesConfigTest = path.Join(testConfigDir, "sh_allergies_test.csv")
	// ProceduresConfigTest is the path to the procedures config file for testing.
	ProceduresConfigTest = path.Join(testConfigDir, "sh_procedures_test.csv")
	// DiagnosesConfigTest is the path to the diagnoses config file for testing.
	DiagnosesConfigTest = path.Join(testConfigDir, "sh_diagnoses_test.csv")
	// ComplexOrderProfilesConfigTest is the path to the config file with complex order profiles for testing.
	ComplexOrderProfilesConfigTest = path.Join(testConfigDir, "sh_complex_order_profiles_test.yml")
	// DoctorsConfigTest is the path to the doctors config file for testing.
	DoctorsConfigTest = path.Join(testConfigDir, "sh_doctors_test.yml")
	// EthnicityConfigTest is the path to the ethnicities config file for testing.
	EthnicityConfigTest = path.Join(testConfigDir, "sh_ethnicity_test.csv")
	// MessageConfigTest is the path to the message config file for testing.
	MessageConfigTest = path.Join(testConfigDir, "sh_message_config_test.yml")
	// HeaderConfigTest is the path to the header config file for testing.
	HeaderConfigTest = path.Join(testConfigDir, "sh_header_config_test.yml")
	// DataMessageConfigTest is the path to the data message config file for testing.
	DataMessageConfigTest = path.Join(testConfigDir, "sh_data_message_config_test.yml")
	// OrderProfilesConfigTest is the path to the config file with order profiles for testing.
	OrderProfilesConfigTest = path.Join(testConfigDir, "sh_order_profiles_test.yml")
	// PatientClassConfigTest is the path to the patient class config file for testing.
	PatientClassConfigTest = path.Join(testConfigDir, "sh_patient_class_test.csv")
	// SurnamesConfigTest is the path to the surnames config file for testing.
	SurnamesConfigTest = path.Join(testConfigDir, "surnames_test.txt")
	// BoysConfigTest is the path to the boys' names config file for testing.
	BoysConfigTest = path.Join(testConfigDir, "historicname_boys_test.csv")
	// GirlsConfigTest is the path to the girls' names config file for testing.
	GirlsConfigTest = path.Join(testConfigDir, "historicname_girls_test.csv")
	// LocationsConfigTest is the path to the locations config file for testing.
	LocationsConfigTest = path.Join(testConfigDir, "sh_locations_test.yml")
	// PathwaysDirTest is the path to the directory with pathways for testing.
	PathwaysDirTest = path.Join(testConfigDir, "sh_pathways")
	// HardcodedMessagesDirTest is the path to the directory with hardcoded messages for testing.
	HardcodedMessagesDirTest = path.Join(testConfigDir, "hardcoded")

	// AllergiesConfigProd is the path to the prod allergies config file.
	AllergiesConfigProd = path.Join(prodConfigDir, "hl7_messages", "allergies.csv")
	// ProceduresConfigProd is the path to the prod procedures config file.
	ProceduresConfigProd = path.Join(prodConfigDir, "hl7_messages", "procedures.csv")
	// DiagnosesConfigProd is the path to the prod diagnoses config file.
	DiagnosesConfigProd = path.Join(prodConfigDir, "hl7_messages", "diagnoses.csv")
	// DoctorsConfigProd is the path to the prod doctors config file.
	DoctorsConfigProd = path.Join(prodConfigDir, "hl7_messages", "doctors.yml")
	// EthnicityConfigProd is the path to the prod ethnicities config file.
	EthnicityConfigProd = path.Join(prodConfigDir, "hl7_messages", "ethnicity.csv")
	// OrderProfilesConfigProd is the path to the prod config file with order profiles.
	OrderProfilesConfigProd = path.Join(prodConfigDir, "hl7_messages", "order_profiles.yml")
	// PatientClassConfigProd is the path to the prod patient class config file.
	PatientClassConfigProd = path.Join(prodConfigDir, "hl7_messages", "patient_class.csv")
	// MessageConfigProd is the path to the prod message config file.
	MessageConfigProd = path.Join(prodConfigDir, "hl7_messages", "hl7.yml")
	// HeaderConfigProd is the path to the prod header config file.
	HeaderConfigProd = path.Join(prodConfigDir, "hl7_messages", "header.yml")
	// DataConfigProd is the path to the prod data config file.
	DataConfigProd = path.Join(prodConfigDir, "hl7_messages", "data.yml")
	// NounsConfigProd is the path to the prod names config file.
	NounsConfigProd = path.Join(prodConfigDir, "hl7_messages", "third_party", "nouns.txt")
	// SurnamesConfigProd is the path to the prod surnames config file.
	SurnamesConfigProd = path.Join(prodConfigDir, "hl7_messages", "third_party", "surnames.txt")
	// BoysConfigProd is the path to the prod boys' names config file.
	BoysConfigProd = path.Join(prodConfigDir, "hl7_messages", "third_party", "historicname_tcm77-254032-boys.csv")
	// GirlsConfigProd is the path to the prod girls' names config file.
	GirlsConfigProd = path.Join(prodConfigDir, "hl7_messages", "third_party", "historicname_tcm77-254032-girls.csv")
	// ClinicalNotesConfigProd is the path to the prod clinical notes config file.
	ClinicalNotesConfigProd = path.Join(prodConfigDir, "hl7_messages", "third_party", "notes")
	// ClinicalNoteTypesConfigProd is the path to the prod clinical note types config file.
	ClinicalNoteTypesConfigProd = path.Join(prodConfigDir, "hl7_messages", "third_party", "note_types.txt")
	// LocationsConfigProd is the path to the prod locations config file.
	LocationsConfigProd = path.Join(prodConfigDir, "hl7_messages", "locations.yml")
	// PathwaysDirProd is the path to the directory with prod pathways.
	PathwaysDirProd = path.Join(prodConfigDir, "pathways")
	// HardcodedMessagesDirProd is the path to the prod directory with hardcoded messages.
	HardcodedMessagesDirProd = path.Join(prodConfigDir, "hardcoded_messages")

	// DataFiles contains sets of data files for testing.
	DataFiles = map[ConfigType]config.DataFiles{
		Test: {
			DataConfig:        DataMessageConfigTest,
			Nouns:             NounsConfigProd,
			Allergies:         AllergiesConfigTest,
			Procedures:        ProceduresConfigTest,
			Diagnoses:         DiagnosesConfigTest,
			Surnames:          SurnamesConfigTest,
			Girls:             GirlsConfigTest,
			Boys:              BoysConfigTest,
			Ethnicities:       EthnicityConfigTest,
			PatientClass:      PatientClassConfigTest,
			ClinicalNoteTypes: ClinicalNoteTypesConfigProd,
			SampleNotesDir:    ClinicalNotesConfigProd,
		},
		Prod: {
			DataConfig:        DataConfigProd,
			Nouns:             NounsConfigProd,
			Allergies:         AllergiesConfigProd,
			Procedures:        ProceduresConfigProd,
			Diagnoses:         DiagnosesConfigProd,
			Surnames:          SurnamesConfigProd,
			Girls:             GirlsConfigProd,
			Boys:              BoysConfigProd,
			Ethnicities:       EthnicityConfigProd,
			PatientClass:      PatientClassConfigProd,
			ClinicalNoteTypes: ClinicalNoteTypesConfigProd,
			SampleNotesDir:    ClinicalNotesConfigProd,
		},
	}
)

func currentDir() string {
	dir, _ := os.Getwd()
	return dir
}
