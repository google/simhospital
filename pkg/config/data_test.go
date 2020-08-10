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
	"context"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/sample"
	"github.com/google/simhospital/pkg/test/testwrite"
)

var (
	defaultData = []byte(`
allergy:
  reactions:
    - "reaction 1"
    - "reaction 2"
patient_name:
  degrees:
    - "B.A."
    - "Ph.D."
address:
  country: "UK"
  cities:
    - "London"
    - "Croydon"
  `)

	invalidData = []byte(`
inexistent_field: arbitrary-string`)

	defaultNouns = []byte(`aardvark
abacus
abbey`)

	defaultNotetypes = []byte(`ED Depart Summary
Gynaecology & Maternity
Surgery Inpatients`)

	defaultAllergies = []byte(`
J30.1,Allergy1,59
J45.0,Allergy2,2556
`)

	defaultDiagnoses = []byte(`
A01.1,Diagnosis1,1
A02.1,Diagnosis2,1
`)

	defaultProcedures = []byte(`
P24.9,Procedure1,1
P25.8,Procedure2,1
`)

	defaultGirls = []byte(`
RANK,1904,1914,1924,1934,1944,1954,1964,1974,1984,1994
,,,,,,,,,,
1,MARY,MARY,MARGARET,MARGARET,MARGARET,SUSAN,SUSAN,SARAH,SARAH,REBECCA
2,FLORENCE,MARGARET,MARY,JEAN,PATRICIA,LINDA,JULIE,CLAIRE,LAURA,LAUREN
3,DORIS,DORIS,JOAN,MARY,CHRISTINE,CHRISTINE,KAREN,NICOLA,GEMMA,JESSICA`)

	invalidGirlsDuplicatedHeader = []byte(`
RANK,1904,1914,1924,1934,1944,1954,1964,1974,1984,1994
RANK,1904,1914,1924,1934,1944,1954,1964,1974,1984,1994
,,,,,,,,,,
1,MARY,MARY,MARGARET,MARGARET,MARGARET,SUSAN,SUSAN,SARAH,SARAH,REBECCA
2,FLORENCE,MARGARET,MARY,JEAN,PATRICIA,LINDA,JULIE,CLAIRE,LAURA,LAUREN
3,DORIS,DORIS,JOAN,MARY,CHRISTINE,CHRISTINE,KAREN,NICOLA,GEMMA,JESSICA`)

	invalidGirlsHeaderAfterData = []byte(`
,,,,,,,,,,
1,MARY,MARY,MARGARET,MARGARET,MARGARET,SUSAN,SUSAN,SARAH,SARAH,REBECCA
RANK,1904,1914,1924,1934,1944,1954,1964,1974,1984,1994
2,FLORENCE,MARGARET,MARY,JEAN,PATRICIA,LINDA,JULIE,CLAIRE,LAURA,LAUREN
3,DORIS,DORIS,JOAN,MARY,CHRISTINE,CHRISTINE,KAREN,NICOLA,GEMMA,JESSICA`)

	invalidGirlsHeaderOutOfOrder = []byte(`
RANK,1904,1914,1924,1934,1944,1954,1964,1974,1994,1984
,,,,,,,,,,
1,MARY,MARY,MARGARET,MARGARET,MARGARET,SUSAN,SUSAN,SARAH,SARAH,REBECCA
2,FLORENCE,MARGARET,MARY,JEAN,PATRICIA,LINDA,JULIE,CLAIRE,LAURA,LAUREN
3,DORIS,DORIS,JOAN,MARY,CHRISTINE,CHRISTINE,KAREN,NICOLA,GEMMA,JESSICA`)

	defaultBoys = []byte(`
RANK,1904,1914,1924,1934,1944,1954,1964,1974,1984,1994
,,,,,,,,,,
1,WILLIAM,JOHN,JOHN,JOHN,JOHN,DAVID,DAVID,PAUL,CHRISTOPHER,THOMAS
2,JOHN,WILLIAM,WILLIAM,PETER,DAVID,JOHN,PAUL,MARK,JAMES,JAMES
3,GEORGE,GEORGE,GEORGE,WILLIAM,MICHAEL,STEPHEN,ANDREW,DAVID,DAVID,JACK`)

	defaultSurnames = []byte(`# Comment
Forster

Nathenson
Croock
`)

	defaultEthnicities = []byte(`
White,White,1
Asian,Asian,2
Black,Black,2
Other,Other,2
nil,nil,3`)

	defaultPatients = []byte(`
EMERGENCY,EMERGENCY,10
OUTPATIENT,OUTPATIENT,40
`)

	sampleNotes = []struct {
		content  string
		filename string
	}{{
		content:  "meme content",
		filename: "test1.png",
	}, {
		content:  "more meme content",
		filename: "test2.jpg",
	}}
)

func TestLoadDataConfig(t *testing.T) {
	ctx := context.Background()
	tmpData := testwrite.BytesToFile(t, defaultData)
	tmpInvalidData := testwrite.BytesToFile(t, invalidData)
	tmpNouns := testwrite.BytesToFile(t, defaultNouns)
	tmpNoteTypes := testwrite.BytesToFile(t, defaultNotetypes)
	tmpSurnames := testwrite.BytesToFile(t, defaultSurnames)
	tmpGirls := testwrite.BytesToFile(t, defaultGirls)
	tmpGirlsDuplicatedHeader := testwrite.BytesToFile(t, invalidGirlsDuplicatedHeader)
	tmpGirlsHeaderAfterData := testwrite.BytesToFile(t, invalidGirlsHeaderAfterData)
	tmpGirlsHeaderOutOfOrder := testwrite.BytesToFile(t, invalidGirlsHeaderOutOfOrder)
	tmpBoys := testwrite.BytesToFile(t, defaultBoys)
	tmpAllergies := testwrite.BytesToFile(t, defaultAllergies)
	tmpDiagnoses := testwrite.BytesToFile(t, defaultDiagnoses)
	tmpProcedures := testwrite.BytesToFile(t, defaultProcedures)
	tmpEthnicities := testwrite.BytesToFile(t, defaultEthnicities)
	tmpPatientClass := testwrite.BytesToFile(t, defaultPatients)

	tmpNotesDir := testwrite.TempDir(t)
	for _, n := range sampleNotes {
		content := []byte(n.content)
		testwrite.BytesToFileInExistingDir(t, content, tmpNotesDir, n.filename)
	}

	tests := []struct {
		name              string
		overrideDataFiles func(df DataFiles) DataFiles
		wantErr           bool
	}{{
		name:              "Success",
		overrideDataFiles: func(df DataFiles) DataFiles { return df },
	}, {
		name: "Unknown fields in data config",
		overrideDataFiles: func(df DataFiles) DataFiles {
			df.DataConfig = tmpInvalidData
			return df
		},
		wantErr: true,
	}, {
		name: "Duplicated header in girls config",
		overrideDataFiles: func(df DataFiles) DataFiles {
			df.Girls = tmpGirlsDuplicatedHeader
			return df
		},
		wantErr: true,
	}, {
		name: "Header after data in girls config",
		overrideDataFiles: func(df DataFiles) DataFiles {
			df.Girls = tmpGirlsHeaderAfterData
			return df
		},
		wantErr: true,
	}, {
		name: "Header years out of order in girls config",
		overrideDataFiles: func(df DataFiles) DataFiles {
			df.Girls = tmpGirlsHeaderOutOfOrder
			return df
		},
		wantErr: true,
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := DataFiles{
				DataConfig:        tmpData,
				Nouns:             tmpNouns,
				Surnames:          tmpSurnames,
				Boys:              tmpBoys,
				Girls:             tmpGirls,
				Allergies:         tmpAllergies,
				Diagnoses:         tmpDiagnoses,
				Procedures:        tmpProcedures,
				Ethnicities:       tmpEthnicities,
				PatientClass:      tmpPatientClass,
				ClinicalNoteTypes: tmpNoteTypes,
				SampleNotesDir:    tmpNotesDir,
			}
			f = tc.overrideDataFiles(f)

			hl7Config := HL7Config{
				Allergy:   HL7Allergy{CodingSystem: "allergy-cs"},
				Diagnosis: HL7Diagnosis{CodingSystem: "diagnosis-cs"},
				Procedure: HL7Procedure{CodingSystem: "procedure-cs"},
			}
			c, err := LoadData(ctx, f, &hl7Config)
			gotErr := err != nil
			if gotErr != tc.wantErr {
				t.Fatalf("LoadData(%+v, %+v) got err=%v; want err? %t", f, hl7Config, err, tc.wantErr)
			}
			if tc.wantErr {
				return
			}

			wantReactions := []string{"reaction 1", "reaction 2"}
			if diff := cmp.Diff(wantReactions, c.Allergy.Reactions); diff != "" {
				t.Errorf("LoadData(%+v, %+v) Allergy.Reactions mismatch (-want +got):\n%s", f, hl7Config, diff)
			}
			wantDegrees := []string{"B.A.", "Ph.D."}
			if diff := cmp.Diff(wantDegrees, c.PatientName.Degrees); diff != "" {
				t.Errorf("LoadData(%+v, %+v) PatientName.Degrees mismatch (-want +got):\n%s", f, hl7Config, diff)
			}
			wantAddress := Address{
				Country: "UK",
				Cities:  []string{"London", "Croydon"},
			}
			if diff := cmp.Diff(wantAddress, c.Address); diff != "" {
				t.Errorf("LoadData(%+v, %+v) Address mismatch (-want +got):\n%s", f, hl7Config, diff)
			}

			if diff := cmp.Diff([]string{"aardvark", "abacus", "abbey"}, c.Nouns); diff != "" {
				t.Errorf("LoadData(%+v, %+v) Nouns mismatch (-want, +got):\n%s", f, hl7Config, diff)
			}
			names := c.FirstNames
			if diff := cmp.Diff([]string{"Mary", "Florence", "Doris"}, names.Girls.ByYear[1904]); diff != "" {
				t.Errorf("Names.Girls.ByYear[1904] mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff([]string{"Susan", "Julie", "Karen"}, names.Girls.ByYear[1964]); diff != "" {
				t.Errorf("Names.Girls.ByYear[1964] mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff([]string{"Rebecca", "Lauren", "Jessica"}, names.Girls.ByYear[1994]); diff != "" {
				t.Errorf("Names.Girls.ByYear[1994] mismatch (-want +got):\n%s", diff)
			}
			if got, want := len(names.Girls.All), 20; got != want {
				t.Errorf("len(Names.Girls.All) = %d, want %d", got, want)
			}

			if diff := cmp.Diff([]string{"William", "John", "George"}, names.Boys.ByYear[1904]); diff != "" {
				t.Errorf("Names.Boys.ByYear[1904] mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff([]string{"David", "Paul", "Andrew"}, names.Boys.ByYear[1964]); diff != "" {
				t.Errorf("Names.Boys.ByYear[1964] mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff([]string{"Thomas", "James", "Jack"}, names.Boys.ByYear[1994]); diff != "" {
				t.Errorf("Names.Boys.ByYear[1994] mismatch (-want +got):\n%s", diff)
			}
			if got, want := len(names.Boys.All), 14; got != want {
				t.Errorf("len(Names.Boys.All) = %d, want %d", got, want)
			}

			if diff := cmp.Diff([]string{"Forster", "Nathenson", "Croock"}, c.Surnames, cmpopts.SortSlices(func(a, b string) bool { return a < b })); diff != "" {
				t.Errorf("Surnames mismatch (-want, +got):\n%s", diff)
			}

			wantAllergies := []MappableWeightedValue{{
				WeightedVal: sample.WeightedValue{
					Value:     &ir.CodedElement{ID: "J30.1", Text: "Allergy1", CodingSystem: "allergy-cs"},
					Frequency: uint(59),
				},
				Mapping: Mapping{Key: "J30.1", Value: "Allergy1"},
			}, {
				WeightedVal: sample.WeightedValue{
					Value:     &ir.CodedElement{ID: "J45.0", Text: "Allergy2", CodingSystem: "allergy-cs"},
					Frequency: uint(2556),
				},
				Mapping: Mapping{Key: "J45.0", Value: "Allergy2"},
			}}
			if diff := cmp.Diff(wantAllergies, c.Allergies); diff != "" {
				t.Errorf("LoadData(%+v, %+v) Allergies mismatch (-want, +got):\n%s", f, hl7Config, diff)
			}

			wantDiagnoses := []MappableWeightedValue{{
				WeightedVal: sample.WeightedValue{
					Value:     &ir.CodedElement{ID: "A01.1", Text: "Diagnosis1", CodingSystem: "diagnosis-cs"},
					Frequency: uint(1),
				},
				Mapping: Mapping{Key: "A01.1", Value: "Diagnosis1"},
			}, {
				WeightedVal: sample.WeightedValue{
					Value:     &ir.CodedElement{ID: "A02.1", Text: "Diagnosis2", CodingSystem: "diagnosis-cs"},
					Frequency: uint(1),
				}, Mapping: Mapping{Key: "A02.1", Value: "Diagnosis2"},
			}}
			if diff := cmp.Diff(wantDiagnoses, c.Diagnoses); diff != "" {
				t.Errorf("LoadData(%+v, %+v) Diagnoses mismatch (-want, +got):\n%s", f, hl7Config, diff)
			}

			wantProcedures := []MappableWeightedValue{{
				WeightedVal: sample.WeightedValue{
					Value:     &ir.CodedElement{ID: "P24.9", Text: "Procedure1", CodingSystem: "procedure-cs"},
					Frequency: uint(1),
				},
				Mapping: Mapping{Key: "P24.9", Value: "Procedure1"},
			}, {
				WeightedVal: sample.WeightedValue{
					Value:     &ir.CodedElement{ID: "P25.8", Text: "Procedure2", CodingSystem: "procedure-cs"},
					Frequency: uint(1),
				},
				Mapping: Mapping{Key: "P25.8", Value: "Procedure2"},
			}}
			if diff := cmp.Diff(wantProcedures, c.Procedures); diff != "" {
				t.Errorf("LoadData(%+v, %+v) Procedures mismatch (-want, +got):\n%s", f, hl7Config, diff)
			}

			wantEthnicities := []sample.WeightedValue{
				{Value: &ir.Ethnicity{ID: "White", Text: "White"}, Frequency: uint(1)},
				{Value: &ir.Ethnicity{ID: "Asian", Text: "Asian"}, Frequency: uint(2)},
				{Value: &ir.Ethnicity{ID: "Black", Text: "Black"}, Frequency: uint(2)},
				{Value: &ir.Ethnicity{ID: "Other", Text: "Other"}, Frequency: uint(2)},
				{Value: (*ir.Ethnicity)(nil), Frequency: uint(3)},
			}
			if diff := cmp.Diff(wantEthnicities, c.Ethnicities); diff != "" {
				t.Errorf("LoadData(%+v, %+v) Ethnicities mismatch (-want, +got):\n%s", f, hl7Config, diff)
			}

			wantPatientClass := []sample.WeightedValue{
				{
					Value:     &PatientClassAndType{Class: "EMERGENCY", Type: "EMERGENCY"},
					Frequency: uint(10),
				},
				{
					Value:     &PatientClassAndType{Class: "OUTPATIENT", Type: "OUTPATIENT"},
					Frequency: uint(40),
				},
			}
			if diff := cmp.Diff(wantPatientClass, c.PatientClass); diff != "" {
				t.Errorf("LoadData(%+v, %+v) PatientClass mismatch (-want, +got):\n%s", f, hl7Config, diff)
			}
			if diff := cmp.Diff([]string{"ED Depart Summary", "Gynaecology & Maternity", "Surgery Inpatients"}, c.ClinicalNoteTypes); diff != "" {
				t.Errorf("LoadData(%+v, %+v) ClinicalNoteTypes mismatch (-want, +got):\n%s", f, hl7Config, diff)
			}
			wantNoteConfig := map[string][]ClinicalNote{
				"png": {{
					DocumentFileName: "test1.png",
					Path:             filepath.Join(tmpNotesDir, "test1.png"),
					ContentType:      "png",
				}},
				"jpg": {{
					DocumentFileName: "test2.jpg",
					Path:             filepath.Join(tmpNotesDir, "test2.jpg"),
					ContentType:      "jpg",
				}},
			}
			if diff := cmp.Diff(wantNoteConfig, c.NotesConfig); diff != "" {
				t.Errorf("LoadData(%+v, %+v) NotesConfig mismatch (-want, +got):\n%s", f, hl7Config, diff)
			}
		})
	}
}

// TestLoadData_CodedElements tests how coded elements are loaded depending on the content of the file.
// Coded elements are Diagnoses, Procedures and Allergies. We load other files too in order to use the
// exposed method LoadData.
func TestLoadData_CodedElements(t *testing.T) {
	ctx := context.Background()
	tmpData := testwrite.BytesToFile(t, defaultData)
	tmpNouns := testwrite.BytesToFile(t, defaultNouns)
	tmpNoteTypes := testwrite.BytesToFile(t, defaultNotetypes)
	tmpSurnames := testwrite.BytesToFile(t, defaultSurnames)
	tmpGirls := testwrite.BytesToFile(t, defaultGirls)
	tmpBoys := testwrite.BytesToFile(t, defaultBoys)
	tmpEthnicities := testwrite.BytesToFile(t, defaultEthnicities)
	tmpPatientClass := testwrite.BytesToFile(t, defaultPatients)
	tmpNotesDir := testwrite.TempDir(t)

	csv := []byte(`
J30.1,Value1,59
J45.0,Value2,2556`)
	tmpCSV := testwrite.BytesToFile(t, csv)

	csvWithComment := []byte(`
# Lines starting with "#" are ignored.
J30.1,Value1,59
J45.0,Value2,2556`)
	tmpCSVWithComment := testwrite.BytesToFile(t, csvWithComment)

	csvWithNil := []byte(`
J30.1,Value1,59
J45.0,Value2,2556
nil,nil,300`)
	tmpCSVWithNil := testwrite.BytesToFile(t, csvWithNil)

	csvNoInt := []byte(`
J30.1,Value1,59.5
J45.0,Value2,2556.6`)
	tmpCSVNoInt := testwrite.BytesToFile(t, csvNoInt)

	// Use the same coding system in all cases to make the assertions easier.
	defaultCS := "coding-system"
	wantCSVWithNil := []MappableWeightedValue{{
		WeightedVal: sample.WeightedValue{
			Value:     &ir.CodedElement{ID: "J30.1", Text: "Value1", CodingSystem: defaultCS},
			Frequency: uint(59),
		},
		Mapping: Mapping{Key: "J30.1", Value: "Value1"},
	}, {
		WeightedVal: sample.WeightedValue{
			Value:     &ir.CodedElement{ID: "J45.0", Text: "Value2", CodingSystem: defaultCS},
			Frequency: uint(2556),
		},
		Mapping: Mapping{Key: "J45.0", Value: "Value2"},
	}, {
		WeightedVal: sample.WeightedValue{
			Value:     &ir.CodedElement{CodingSystem: defaultCS},
			Frequency: uint(300),
		},
	}}

	wantCSVNoNil := []MappableWeightedValue{{
		WeightedVal: sample.WeightedValue{
			Value:     &ir.CodedElement{ID: "J30.1", Text: "Value1", CodingSystem: defaultCS},
			Frequency: uint(59),
		},
		Mapping: Mapping{Key: "J30.1", Value: "Value1"},
	}, {
		WeightedVal: sample.WeightedValue{
			Value:     &ir.CodedElement{ID: "J45.0", Text: "Value2", CodingSystem: defaultCS},
			Frequency: uint(2556),
		},
		Mapping: Mapping{Key: "J45.0", Value: "Value2"},
	}}

	tests := []struct {
		name              string
		overrideDataFiles func(df DataFiles) DataFiles
		wantErr           bool
		wantDiagnoses     []MappableWeightedValue
		wantProcedures    []MappableWeightedValue
		wantAllergies     []MappableWeightedValue
	}{{
		name: "Regular",
		overrideDataFiles: func(df DataFiles) DataFiles {
			df.Diagnoses = tmpCSV
			df.Procedures = tmpCSV
			df.Allergies = tmpCSV
			return df
		},
		wantDiagnoses:  wantCSVNoNil,
		wantProcedures: wantCSVNoNil,
		wantAllergies:  wantCSVNoNil,
	}, {
		name: `Lines starting with "#" are ignored`,
		overrideDataFiles: func(df DataFiles) DataFiles {
			df.Diagnoses = tmpCSVWithComment
			df.Procedures = tmpCSVWithComment
			df.Allergies = tmpCSVWithComment
			return df
		},
		wantDiagnoses:  wantCSVNoNil,
		wantProcedures: wantCSVNoNil,
		wantAllergies:  wantCSVNoNil,
	}, {
		name: "Nils in Diagnoses and Procedures",
		overrideDataFiles: func(df DataFiles) DataFiles {
			df.Diagnoses = tmpCSVWithNil
			df.Procedures = tmpCSVWithNil
			df.Allergies = tmpCSV
			return df
		},
		wantDiagnoses:  wantCSVWithNil,
		wantProcedures: wantCSVWithNil,
		wantAllergies:  wantCSVNoNil,
	}, {
		name: "Allergies don't support nils",
		overrideDataFiles: func(df DataFiles) DataFiles {
			df.Diagnoses = tmpCSVWithNil
			df.Procedures = tmpCSVWithNil
			df.Allergies = tmpCSVWithNil
			return df
		},
		wantErr: true,
	}, {
		name: "No integers",
		overrideDataFiles: func(df DataFiles) DataFiles {
			df.Diagnoses = tmpCSVNoInt
			df.Procedures = tmpCSV
			df.Allergies = tmpCSV
			return df
		},
		wantErr: true,
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := DataFiles{
				DataConfig:        tmpData,
				Nouns:             tmpNouns,
				Surnames:          tmpSurnames,
				Boys:              tmpBoys,
				Girls:             tmpGirls,
				Ethnicities:       tmpEthnicities,
				PatientClass:      tmpPatientClass,
				ClinicalNoteTypes: tmpNoteTypes,
				SampleNotesDir:    tmpNotesDir,
			}
			f = tc.overrideDataFiles(f)

			hl7Config := HL7Config{
				Allergy:   HL7Allergy{CodingSystem: defaultCS},
				Diagnosis: HL7Diagnosis{CodingSystem: defaultCS},
				Procedure: HL7Procedure{CodingSystem: defaultCS},
			}
			c, err := LoadData(ctx, f, &hl7Config)
			gotErr := err != nil
			if gotErr != tc.wantErr {
				t.Fatalf("LoadData(%+v, %+v) got err=%v; want err? %t", f, hl7Config, err, tc.wantErr)
			}
			if tc.wantErr || gotErr {
				return
			}

			if diff := cmp.Diff(tc.wantAllergies, c.Allergies); diff != "" {
				t.Errorf("LoadData(%+v, %+v) Allergies mismatch (-want, +got):\n%s", f, hl7Config, diff)
			}

			if diff := cmp.Diff(tc.wantDiagnoses, c.Diagnoses); diff != "" {
				t.Errorf("LoadData(%+v, %+v) Diagnoses mismatch (-want, +got):\n%s", f, hl7Config, diff)
			}

			if diff := cmp.Diff(tc.wantProcedures, c.Procedures); diff != "" {
				t.Errorf("LoadData(%+v, %+v) Procedures mismatch (-want, +got):\n%s", f, hl7Config, diff)
			}
		})
	}
}
