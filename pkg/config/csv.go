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
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"

	"github.com/pkg/errors"
	"github.com/google/simhospital/pkg/files"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/sample"
)

// nilKey is a keyword to use in CSV files that are loaded with loadCSVWithFrequency.
// Rows for which all items (except the frequency) match this keyword are represented with a nil RecordWithFreq.Value.
const nilKey = "nil"

// RecordWithFreq stores a record as a list of strings (e.g. from CSV) and its associated frequency value.
type RecordWithFreq struct {
	Value  map[string]string
	Weight uint
}

// loadCSVWithFrequency loads a CSV file where each row is a list of strings and the last
// column is a frequency represented as int, and returns a slice where each element corresponds
// to one row in the file. Rows that start with # are ignored.
// The columnKeys parameter are the keys to be set in the RecordWithFreq.Value map. The keys are
// mapped to the items in the rows in order: the first key will be used for the first element in
// each row, and successively.
// All rows are expected to have the same number of columns.
// Rows for which all items (except the frequency) are the keyword "nil" are represented as a
// nil RecordWithFreq.Value. Callers need to check for the presence of a nil Value.
// Example format for the file:
//   # Distribution of patient classes.
//   OUTPATIENT,EMERGENCY,10
//   nil,nil,20
// Output for columnKeys ("class", "type") :
// []RecordWithFreq {
//	{
//		Value: map[string]string{
//			"class": "OUTPATIENT",
//			"type": "EMERGENCY",
//		},
//		Weight: 10,
//	},
//	{
//		Value: nil,
//		Weight: 20,
//	},
//  (etc).
//}
func loadCSVWithFrequency(ctx context.Context, fName string, columnKeys []string) ([]RecordWithFreq, error) {
	b, err := files.Read(ctx, fName)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot open file %s", fName)
	}
	reader := csv.NewReader(bytes.NewReader(b))
	reader.Comment = '#'
	var records []RecordWithFreq
	nColumns := len(columnKeys) + 1
	for record, err := reader.Read(); err != io.EOF; record, err = reader.Read() {
		if err != nil {
			return nil, errors.Wrapf(err, "cannot read file %s", fName)
		}

		if len(record) != nColumns {
			return nil, fmt.Errorf("cannot load frequencies from file %s: got %d elements in one line, want %d", fName, len(record), nColumns)
		}

		m := map[string]string{}
		countNil := 0
		for i := 0; i < nColumns-1; i++ {
			m[columnKeys[i]] = record[i]
			if record[i] == nilKey {
				countNil++
			}
		}
		if countNil == nColumns-1 {
			m = nil
		}
		frequencyField := record[nColumns-1]
		weight, err := strconv.ParseInt(frequencyField, 10, 32)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot load frequencies from file %s: got frequency %s, want int", fName, frequencyField)
		}
		records = append(records, RecordWithFreq{
			Value:  m,
			Weight: uint(weight),
		})
	}
	return records, nil
}

// loadCodedElements loads a CSV file where each row contains one coded element and its frequency.
// The first element of each row is the code of the coded element, the second is its description,
// and the last one is the frequency.
// allowNil specifies whether nil rows are allowed. A nil row is of the form "nil,nil,<frequency>".
func loadCodedElements(ctx context.Context, fileName string, codingSystem string, allowNil bool) ([]MappableWeightedValue, error) {
	idKey := "id"
	textKey := "text"

	recordsWithFrequency, err := loadCSVWithFrequency(ctx, fileName, []string{idKey, textKey})
	if err != nil {
		return nil, err
	}

	var values []MappableWeightedValue
	for _, record := range recordsWithFrequency {
		if !allowNil && record.Value == nil {
			return nil, fmt.Errorf("found nil value in file %s; nil values not supported", fileName)
		}
		id := record.Value[idKey]
		key := record.Value[textKey]
		values = append(values, MappableWeightedValue{
			WeightedVal: sample.WeightedValue{
				Value: &ir.CodedElement{
					ID:           id,
					Text:         key,
					CodingSystem: codingSystem,
				},
				Frequency: record.Weight,
			},
			Mapping: Mapping{Key: id, Value: key},
		})
	}
	return values, nil
}

func ethnicities(ctx context.Context, fileName string) ([]sample.WeightedValue, error) {
	idKey := "id"
	textKey := "text"

	recordsWithFrequency, err := loadCSVWithFrequency(ctx, fileName, []string{idKey, textKey})
	if err != nil {
		return nil, err
	}

	var distr []sample.WeightedValue
	for _, record := range recordsWithFrequency {
		var e *ir.Ethnicity
		if record.Value != nil {
			e = &ir.Ethnicity{
				ID:   record.Value[idKey],
				Text: record.Value[textKey],
			}
		}
		distr = append(distr, sample.WeightedValue{
			Value:     e,
			Frequency: record.Weight,
		})
	}
	return distr, nil
}

// PatientClassAndType represents a class and type pair.
type PatientClassAndType struct {
	Class string
	Type  string
}

func patientClass(ctx context.Context, fileName string) ([]sample.WeightedValue, error) {
	classKey := "class"
	typeKey := "type"

	recordsWithFrequency, err := loadCSVWithFrequency(ctx, fileName, []string{classKey, typeKey})
	if err != nil {
		return nil, err
	}

	var distr []sample.WeightedValue
	for _, record := range recordsWithFrequency {
		if record.Value == nil {
			return nil, fmt.Errorf("cannot load patient classes from file %s: nil records are not supported", fileName)
		}
		patientC := record.Value[classKey]
		patientT := record.Value[typeKey]
		distr = append(distr, sample.WeightedValue{
			Value:     &PatientClassAndType{Class: patientC, Type: patientT},
			Frequency: record.Weight,
		})
	}
	return distr, nil
}
