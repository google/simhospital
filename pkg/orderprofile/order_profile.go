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

// Package orderprofile is responsible for parsing and generating Order Profiles.
package orderprofile

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/constants"
	"github.com/google/simhospital/pkg/files"
	"github.com/google/simhospital/pkg/ir"
)

// OrderProfiles contains Order Profile information.
type OrderProfiles struct {
	// op is a map of Order Profiles keyed by their names.
	op map[string]*OrderProfile
	// names is a slice of all Order Profile names.
	names []string
}

// OrderProfile contains details of an Order Profile.
type OrderProfile struct {
	// UniversalService encapsulated UniversalServiceID, Order Profile name (represented as Text)
	// and CodingSystem.
	UniversalService ir.CodedElement
	// TestTypes is a map of all Test Types for the Order Profile, keys by their names.
	TestTypes map[string]*TestType
}

// TestType represents the Test Type of the Order Profile.
type TestType struct {
	Name         ir.CodedElement
	defaultValue validString
	Unit         string
	ValueType    string
	// isNumerical indicates whether the value is numerical, or textual.
	// Note: this may not be the same as ValueType above. ValueType is used
	// to populate the HL7 message and is set to whatever is defined in the order profile,
	// while isNumerical indicates whether we treat this value as numerical or not.
	// Eg '>12.4' could have ValueType = 'TX', but isNumerical = true,
	// as we can interpret this as a number.
	isNumerical bool
	// valuePrefix is the prefix of the value.
	// Some numerical values are represented as: <0.2 / >=5.5 etc.
	// valuePrefix stores < / >= etc so that we can then construct valid random value.
	valuePrefix    string
	RefRange       string
	ValueGenerator *ValueGenerator
}

type validString struct {
	value string
	valid bool
}

func newValidString(s string) validString {
	return validString{value: s, valid: true}
}

func newInvalidString() validString {
	return validString{valid: false}
}

// New returns a new OrderProfiles from a order profiles map.
func New(m map[string]*OrderProfile) *OrderProfiles {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return &OrderProfiles{
		op:    m,
		names: keys,
	}
}

// Get returns an OrderProfile for the given Order Profile name.
func (op *OrderProfiles) Get(name string) (*OrderProfile, bool) {
	v, ok := op.op[name]
	return v, ok
}

// Generate returns a CodedElement for the given name.
// If the name is constants.RandomString, it returns a CodedElement for a random Order Profile.
// If the name is a name of any existing Order Profile, the CodedElement for that Order Profile
// is returned.
// Otherwise, returns CodedElement with ID and Text equal to given name.
func (op *OrderProfiles) Generate(name string) *ir.CodedElement {
	if name == constants.RandomString {
		name = op.names[rand.Intn(len(op.names))]
	}
	if v, ok := op.op[name]; ok {
		return &v.UniversalService
	}

	log.Warningf("No order profile defined for: %s", name)
	return &ir.CodedElement{ID: name, Text: name}
}

// RandomisedValueWithFlag generates a random value for the given type.
// Returns:
// - the value, which is either defaultValue if set, or the random value generated using the valueGenerator.
// - abnormal flag, HIGH if randomType is ABNORMAL_HIGH, LOW if randomType is ABNORMAL_LOW, or else an empty string.
// - an error if something went wrong.
func (tt *TestType) RandomisedValueWithFlag(randomType string) (string, constants.AbnormalFlag, error) {
	abnormalFlag := constants.FromRandomType(randomType)

	if tt.defaultValue.valid {
		if abnormalFlag == constants.AbnormalFlagEmpty || !tt.isNumerical {
			return tt.defaultValue.value, abnormalFlag, nil
		}
	}
	v, err := tt.ValueGenerator.Random(randomType)
	if err != nil {
		return "", "", errors.Wrap(err, "cannot generate random value with flag")
	}
	return fmt.Sprintf("%s%s", tt.valuePrefix, v), abnormalFlag, nil
}

type tt struct {
	ID           string
	CodingSystem string `yaml:"coding_system"`
	ValueType    string `yaml:"value_type"`
	Value        string
	Unit         string
	RefRange     string `yaml:"ref_range"`
}

type op struct {
	UniversalServiceID string        `yaml:"universal_service_id"`
	CodingSystem       string        `yaml:"coding_system"`
	TestTypes          map[string]tt `yaml:"test_types"`
}

func testType(ttName string, ttValue tt, codingSystem string) *TestType {
	if ttValue.CodingSystem != "" {
		codingSystem = ttValue.CodingSystem
	}
	testType := &TestType{
		Name:      ir.CodedElement{ID: ttValue.ID, Text: ttName, CodingSystem: codingSystem},
		Unit:      ttValue.Unit,
		ValueType: ttValue.ValueType,
		RefRange:  ttValue.RefRange,
	}

	if ttValue.ValueType == constants.NumericalValueType {
		prefix, _, err := ValueFromString(ttValue.Value)
		if err != nil {
			// Even though the ValueType == "NM", the example value given in the order profile yml file is
			// not actually numerical, eg: <40 copies/ml.
			// Always use the value from yml file.
			testType.defaultValue = newValidString(ttValue.Value)
			testType.isNumerical = false
			return testType
		}

		vg, err := ValueGeneratorFromRange(ttValue.RefRange)
		if err != nil {
			// Cannot parse the range, eg: [9.00am <46 ] or [ Random ].
			// Always use the value from yml file.
			testType.defaultValue = newValidString(ttValue.Value)
			testType.isNumerical = true
			return testType
		}
		// We can parse the range, do not set the testType.defaultValue but testType.ValueGenerator,
		// so that the value can be generated randomly.
		// Also set the testType.valuePrefix, so that the generated values can be prefixed properly.
		testType.defaultValue = newInvalidString()
		testType.ValueGenerator = vg
		testType.valuePrefix = prefix
		testType.isNumerical = true
		return testType
	}

	testType.defaultValue = newValidString(ttValue.Value)
	testType.isNumerical = false
	return testType
}

// Load parses the order profiles from the given file.
func Load(ctx context.Context, filename string, hl7Config *config.HL7Config) (*OrderProfiles, error) {
	data, err := files.Read(ctx, filename)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot parse order profiles file %s", filename)
	}

	parsed := map[string]op{}
	err = yaml.UnmarshalStrict(data, &parsed)
	if err != nil {
		return nil, errors.Wrapf(err, "Cannot unmarshal order profiles from %v", filename)
	}

	orderProfiles := map[string]*OrderProfile{}
	log.Info("Loading order profiles")
	for k, v := range parsed {
		testTypes := map[string]*TestType{}
		for ttName, ttValue := range v.TestTypes {
			testTypes[ttName] = testType(ttName, ttValue, hl7Config.CodingSystem)
		}

		codingSystem := hl7Config.CodingSystem
		if v.CodingSystem != "" {
			codingSystem = v.CodingSystem
		}
		orderProfiles[k] = &OrderProfile{
			UniversalService: ir.CodedElement{ID: v.UniversalServiceID, Text: k, CodingSystem: codingSystem},
			TestTypes:        testTypes,
		}
		log.Infof(" - %s", k)
	}

	return New(orderProfiles), nil
}
