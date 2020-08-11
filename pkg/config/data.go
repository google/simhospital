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

import (
	"bufio"
	"bytes"
	"context"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"github.com/google/simhospital/pkg/files"
	"github.com/google/simhospital/pkg/logging"
	"github.com/google/simhospital/pkg/sample"
)

var log = logging.ForCallerPackage()

// MappableWeightedValue is used where a mapping needs to be kept alongside the WeightedValue when loading data.
type MappableWeightedValue struct {
	WeightedVal sample.WeightedValue
	Mapping     Mapping
}

// Mapping represents a key value pair.
type Mapping struct {
	Key   string
	Value string
}

// Data contains data for generating HL7 messages.
type Data struct {
	Allergy      DataAllergy
	PatientName  PatientName
	Address      Address
	Nouns        []string
	Surnames     []string
	FirstNames   *FirstNamesByCensus
	Diagnoses    []MappableWeightedValue
	Procedures   []MappableWeightedValue
	Allergies    []MappableWeightedValue
	Ethnicities  []sample.WeightedValue
	PatientClass []sample.WeightedValue
	// NotesConfig maps file extensions with available list of sample notes.
	NotesConfig       map[string][]ClinicalNote
	ClinicalNoteTypes []string
}

type simpleConfig struct {
	Allergy     DataAllergy
	PatientName PatientName `yaml:"patient_name"`
	Address     Address
}

// DataAllergy contains data for generating allergies.
type DataAllergy struct {
	// Reactions contains the possible types of allergy reactions to be set in
	// the AL1.5.AllergyReaction field.
	Reactions []string
	// Percentage is a percentage of people with allergies, from 0 to 100.
	Percentage int
	// MaximumAllergies is the maximum number of allergies generated per person.
	MaximumAllergies int `yaml:"maximum_allergies"`
}

// PatientName contains data for generating patient name.
type PatientName struct {
	// Degrees contains the list of degrees for patient names.
	Degrees []string
	// DegreePercentage  is a percentage of people with degrees, from 0 to 100.
	DegreePercentage int `yaml:"degree_percentage"`
	// Suffixes contains a list of suffixes for patient names.
	Suffixes []string
	// SuffixPercentage is a percentage of people with suffixes, from 0 to 100.
	SuffixPercentage int `yaml:"suffix_percentage"`
	// FemalePrefixes contains a list of female name prefixes.
	FemalePrefixes []string `yaml:"female_prefixes"`
	// MalePrefixes contains a list of male name prefixes.
	MalePrefixes []string `yaml:"male_prefixes"`
	// MiddlenamePercentage is a percentage of people with middlenames,
	// from 0 to 100.
	MiddlenamePercentage int `yaml:"middlename_percentage"`
}

// Address contains data for generating the address.
type Address struct {
	// Cities contains a list of cities for address.
	Cities []string
	// Streets contains a list of street suffixes for addresses.
	Streets []string
	// Country is the Country to set in the XAD.6 Country field.
	Country string
	// Types contains a list of types of addresses to be set in the
	// XAD.7 Address Type field.
	Types []string
}

// DataFiles are the files to load data configuration from.
// All fields are required.
type DataFiles struct {
	DataConfig        string
	Nouns             string
	Surnames          string
	Girls             string
	Boys              string
	Procedures        string
	Diagnoses         string
	Allergies         string
	Ethnicities       string
	PatientClass      string
	SampleNotesDir    string
	ClinicalNoteTypes string
}

// LoadData loads the data configuration from the given data files.
func LoadData(ctx context.Context, f DataFiles, hc *HL7Config) (*Data, error) {
	data, err := files.Read(ctx, f.DataConfig)
	if err != nil {
		return nil, err
	}

	var c simpleConfig
	err = yaml.UnmarshalStrict(data, &c)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot unmarshal configuration file %q", f.DataConfig)
	}
	nouns, err := textFileToList(ctx, f.Nouns)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load nouns from file %q", f.Nouns)
	}
	firstNames, err := census(ctx, f.Girls, f.Boys)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load first names from files girls=%s, boys=%s", f.Girls, f.Boys)
	}
	surnames, err := textFileToList(ctx, f.Surnames)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load surnames from file %q", f.Surnames)
	}
	allergies, err := loadCodedElements(ctx, f.Allergies, hc.Allergy.CodingSystem, false)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load allergies from file %q", f.Allergies)
	}
	log.WithField("file", f.Allergies).Infof("Loaded %d allergies", len(allergies))

	diagnoses, err := loadCodedElements(ctx, f.Diagnoses, hc.Diagnosis.CodingSystem, true)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load diagnoses from file %q", f.Diagnoses)
	}
	log.WithField("file", f.Diagnoses).Infof("Loaded %d diagnoses", len(diagnoses))

	procedures, err := loadCodedElements(ctx, f.Procedures, hc.Procedure.CodingSystem, true)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load procedures from file %q", f.Procedures)
	}
	log.WithField("file", f.Procedures).Infof("Loaded %d procedures", len(procedures))

	ethnicities, err := ethnicities(ctx, f.Ethnicities)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load ethnicities from file %q", f.Ethnicities)
	}
	log.WithField("file", f.Ethnicities).Infof("Loaded %d ethnicities", len(ethnicities))

	patientClass, err := patientClass(ctx, f.PatientClass)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load patient classes from file %q", f.PatientClass)
	}
	log.WithField("file", f.PatientClass).Infof("Loaded %d patient classes", len(patientClass))

	noteTypes, err := textFileToList(ctx, f.ClinicalNoteTypes)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load Clinical Note types from file %q", f.ClinicalNoteTypes)
	}
	log.WithField("file", f.ClinicalNoteTypes).Infof("Loaded %d Clinical Note types", len(noteTypes))

	notesConfig, err := LoadNotesConfig(ctx, f.SampleNotesDir)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot load NotesConfig from directory %q", f.SampleNotesDir)
	}
	log.WithField("directory", f.SampleNotesDir).Infof("Loaded %d different types of sample clinical notes", len(notesConfig))

	return &Data{
		Allergy:           c.Allergy,
		PatientName:       c.PatientName,
		Address:           c.Address,
		Nouns:             nouns,
		Surnames:          surnames,
		FirstNames:        firstNames,
		Diagnoses:         diagnoses,
		Procedures:        procedures,
		Allergies:         allergies,
		Ethnicities:       ethnicities,
		PatientClass:      patientClass,
		NotesConfig:       notesConfig,
		ClinicalNoteTypes: noteTypes,
	}, nil
}

// textFileToList takes a filename as input and returns a list of strings where each
// string is a line in the input file.
// Lines that start with # are ignored.
func textFileToList(ctx context.Context, fileName string) ([]string, error) {
	b, err := files.Read(ctx, fileName)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot open file: %q", fileName)
	}

	var lines []string
	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "#") && line != "" {
			lines = append(lines, line)
		}
	}

	return lines, nil
}
