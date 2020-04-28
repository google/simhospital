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

// ConfigFiles contains file paths to configure a hospital.
type ConfigFiles struct {
	DoctorFile                 string
	OrderProfilesFile          string
	PathwaysDir                string
	MessageConfigFile          string
	HeaderConfigFile           string
	LocationsFile              string
	HardcodedMessagesConfigDir string
}

var (
	base = path.Join(currentDir())

	// testConfigDir is the directory for test files for testing.
	testConfigDir = path.Join(base, "pkg", "test", "data")
	// prodConfigDir is the directory for prod files used during tests.
	// Some tests load the prod values to make sure the files are correct.
	prodConfigDir = path.Join(base, "configs")

	// These file paths are used for testing only.
	allergiesConfig            = path.Join(testConfigDir, "sh_allergies_test.csv")
	proceduresConfig           = path.Join(testConfigDir, "sh_procedures_test.csv")
	diagnosesConfig            = path.Join(testConfigDir, "sh_diagnoses_test.csv")
	ComplexOrderProfilesConfig = path.Join(testConfigDir, "sh_complex_order_profiles_test.yml")
	DoctorsConfig              = path.Join(testConfigDir, "sh_doctors_test.yml")
	ethnicityConfig            = path.Join(testConfigDir, "sh_ethnicity_test.csv")
	MessageConfig              = path.Join(testConfigDir, "sh_message_config_test.yml")
	HeaderConfig               = path.Join(testConfigDir, "sh_header_config_test.yml")
	dataMessageConfig          = path.Join(testConfigDir, "sh_data_message_config_test.yml")
	OrderProfilesConfig        = path.Join(testConfigDir, "sh_order_profiles_test.yml")
	patientClassConfig         = path.Join(testConfigDir, "sh_patient_class_test.csv")
	surnamesConfig             = path.Join(testConfigDir, "surnames_test.txt")
	boysConfig                 = path.Join(testConfigDir, "historicname_boys_test.csv")
	girlsConfig                = path.Join(testConfigDir, "historicname_girls_test.csv")
	hardcodedMessagesDir       = path.Join(testConfigDir, "hardcoded")

	publicAllergiesConfig         = path.Join(prodConfigDir, "hl7_messages", "allergies.csv")
	PublicProceduresConfig        = path.Join(prodConfigDir, "hl7_messages", "procedures.csv")
	PublicDiagnosesConfig         = path.Join(prodConfigDir, "hl7_messages", "diagnoses.csv")
	PublicDoctorsConfig           = path.Join(prodConfigDir, "hl7_messages", "doctors.yml")
	publicEthnicityConfig         = path.Join(prodConfigDir, "hl7_messages", "ethnicity.csv")
	PublicOrderProfilesConfig     = path.Join(prodConfigDir, "hl7_messages", "order_profiles.yml")
	publicPatientClassConfig      = path.Join(prodConfigDir, "hl7_messages", "patient_class.csv")
	PublicMessageConfig           = path.Join(prodConfigDir, "hl7_messages", "hl7.yml")
	PublicHeaderConfig            = path.Join(prodConfigDir, "hl7_messages", "header.yml")
	PublicDataConfig              = path.Join(prodConfigDir, "hl7_messages", "data.yml")
	PublicNounsConfig             = path.Join(prodConfigDir, "hl7_messages", "third_party", "nouns.txt")
	PublicSurnamesConfig          = path.Join(prodConfigDir, "hl7_messages", "third_party", "surnames.txt")
	PublicBoysConfig              = path.Join(prodConfigDir, "hl7_messages", "third_party", "historicname_tcm77-254032-boys.csv")
	PublicGirlsConfig             = path.Join(prodConfigDir, "hl7_messages", "third_party", "historicname_tcm77-254032-girls.csv")
	PublicClinicalNotes           = path.Join(prodConfigDir, "hl7_messages", "third_party", "notes")
	PublicClinicalNoteTypesConfig = path.Join(prodConfigDir, "hl7_messages", "third_party", "note_types.txt")

	PublicLocationsConfig          = path.Join(prodConfigDir, "hl7_messages", "locations.yml")
	PathwaysConfigDir              = path.Join(testConfigDir, "sh_pathways")
	ProdPathwaysDir                = path.Join(prodConfigDir, "pathways")
	ProdHardcodedMessagesConfigDir = path.Join(prodConfigDir, "hardcoded_messages")

	DataFiles = map[ConfigType]config.DataFiles{
		Test: {
			DataConfig:        dataMessageConfig,
			Nouns:             PublicNounsConfig,
			Allergies:         allergiesConfig,
			Procedures:        proceduresConfig,
			Diagnoses:         diagnosesConfig,
			Surnames:          surnamesConfig,
			Girls:             girlsConfig,
			Boys:              boysConfig,
			Ethnicities:       ethnicityConfig,
			PatientClass:      patientClassConfig,
			ClinicalNoteTypes: PublicClinicalNoteTypesConfig,
			SampleNotesDir:    PublicClinicalNotes,
		},
		Prod: {
			DataConfig:        PublicDataConfig,
			Nouns:             PublicNounsConfig,
			Allergies:         publicAllergiesConfig,
			Procedures:        PublicProceduresConfig,
			Diagnoses:         PublicDiagnosesConfig,
			Surnames:          PublicSurnamesConfig,
			Girls:             PublicGirlsConfig,
			Boys:              PublicBoysConfig,
			Ethnicities:       publicEthnicityConfig,
			PatientClass:      publicPatientClassConfig,
			ClinicalNoteTypes: PublicClinicalNoteTypesConfig,
			SampleNotesDir:    PublicClinicalNotes,
		},
	}
	// HospitalFiles contains paths to files for testing.
	HospitalFiles = map[ConfigType]ConfigFiles{
		Test: {
			DoctorFile:                 DoctorsConfig,
			OrderProfilesFile:          OrderProfilesConfig,
			PathwaysDir:                PathwaysConfigDir,
			MessageConfigFile:          MessageConfig,
			HeaderConfigFile:           HeaderConfig,
			HardcodedMessagesConfigDir: hardcodedMessagesDir,
		},
	}
)

func currentDir() string {
	dir, _ := os.Getwd()
	return dir
}
