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

package codedelement

import (
	"context"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/pathway"
	"github.com/google/simhospital/pkg/test"
	"github.com/google/simhospital/pkg/test/testclock"
	"github.com/google/simhospital/pkg/test/testdate"
	"github.com/google/simhospital/pkg/test/testwrite"
)

func TestDiagOrProcGenerator_Random(t *testing.T) {
	ctx := context.Background()
	rand.Seed(1)
	diagnosisFilename := testwrite.BytesToFile(t, []byte(`
A01.1,Diagnosis1,1
A02.1,Diagnosis2,1
A03.1,Diagnosis3,1
A04.1,Diagnosis4,1
`))
	procedureFilename := testwrite.BytesToFile(t, []byte(`
P01.1,Procedure1,1
P02.1,Procedure2,1
P03.1,Procedure3,1
P04.1,Procedure4,1
`))

	f := test.DataFiles[test.Test]
	f.Diagnoses = diagnosisFilename
	f.Procedures = procedureFilename

	c, err := config.LoadHL7Config(ctx, test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	data, err := config.LoadData(ctx, f, c)
	if err != nil {
		t.Fatalf("LoadData(%+v, %+v) failed with %v", f, c, err)
	}

	tclock := testclock.New(defaultDate)
	pathwayDate := time.Date(2016, 1, 2, 0, 0, 0, 0, time.UTC)
	dg := &testdate.Generator{}

	tests := []struct {
		name             string
		g                *DiagOrProcGenerator
		wantCodingSystem string
		wantTypes        []string
		wantRandom       bool
		wantRandomCode   map[string]string
		input            *pathway.DiagnosisOrProcedure
		want             *ir.DiagnosisOrProcedure
	}{{
		name: "Diagnosis from pathway",
		g:    NewDiagnosisGenerator(c, data, tclock, dg),
		input: &pathway.DiagnosisOrProcedure{
			Type:        "some-type",
			Description: "description",
			Code:        "code",
		},
		want: &ir.DiagnosisOrProcedure{
			Description: &ir.CodedElement{ID: "code", Text: "description"},
			Type:        "some-type",
			DateTime:    ir.NullTime{Valid: true, Time: pathwayDate},
		},
		wantTypes: c.Diagnosis.Types,
	}, {
		name: "Random Diagnosis",
		g:    NewDiagnosisGenerator(c, data, tclock, dg),
		input: &pathway.DiagnosisOrProcedure{
			Description: "RANDOM",
		},
		wantRandomCode:   map[string]string{"A01.1": "Diagnosis1", "A02.1": "Diagnosis2", "A03.1": "Diagnosis3", "A04.1": "Diagnosis4"},
		wantTypes:        c.Diagnosis.Types,
		wantCodingSystem: c.Diagnosis.CodingSystem,
	}, {
		name: "Procedure from pathway",
		g:    NewProcedureGenerator(c, data, tclock, dg),
		input: &pathway.DiagnosisOrProcedure{
			Type:        "some-type",
			Description: "description",
			Code:        "code",
		},
		want: &ir.DiagnosisOrProcedure{
			Description: &ir.CodedElement{ID: "code", Text: "description"},
			Type:        "some-type",
			DateTime:    ir.NullTime{Valid: true, Time: pathwayDate},
		},
		wantTypes: c.Procedure.Types,
	}, {
		name: "Random Procedure",
		g:    NewProcedureGenerator(c, data, tclock, dg),
		input: &pathway.DiagnosisOrProcedure{
			Description: "RANDOM",
		},
		wantRandomCode:   map[string]string{"P01.1": "Procedure1", "P02.1": "Procedure2", "P03.1": "Procedure3", "P04.1": "Procedure4"},
		wantTypes:        c.Procedure.Types,
		wantCodingSystem: c.Procedure.CodingSystem,
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := tc.g
			dt := &pathway.DateTime{Time: &pathwayDate}
			if tc.want != nil {
				got := g.RandomOrFromPathway(dt, tc.input)
				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Errorf("RandomOrFromPathway(%v, %+v) got diff (-want, +got):\n%s", dt, tc.input, diff)
				}
				return
			}

			// Randomly generated items.
			runs := float64(1000)
			typesCount := make(map[string]int)
			for i := 0; i < int(runs); i++ {
				got := g.RandomOrFromPathway(dt, tc.input)
				if got == nil {
					t.Fatalf("RandomOrFromPathway(%v, %+v) got nil, want not nil", dt, tc.input)
				}

				desc := got.Description
				id := desc.ID
				if _, ok := tc.wantRandomCode[id]; !ok {
					t.Errorf("RandomOrFromPathway(%v, %+v).ID = %s, want one of %v", dt, tc.input, id, tc.wantRandomCode)
				}
				if got, want := desc.Text, tc.wantRandomCode[id]; got != want {
					t.Errorf("RandomOrFromPathway(%v, %+v).Text got %v, want %v", dt, tc.input, got, want)
				}
				if got, want := desc.CodingSystem, tc.wantCodingSystem; got != want {
					t.Errorf("RandomOrFromPathway(%v, %+v).CodingSystem got %v, want %v", dt, tc.input, got, want)
				}

				if got, want := got.DateTime, ir.NewValidTime(pathwayDate); got != want {
					t.Errorf("RandomOrFromPathway(%v, %+v).DateTime got %v, want %v", dt, tc.input, got, want)
				}

				gotRType := got.Type
				if !contains(tc.wantTypes, gotRType) {
					t.Errorf("RandomOrFromPathway(%v, %+v).Type=%q, want one of %v", dt, tc.input, gotRType, tc.wantTypes)
				}
				typesCount[gotRType]++
			}

			// Allow an error of 20% of the number of runs.
			delta := runs / 5
			for k, got := range typesCount {
				if want := runs / float64(len(tc.wantTypes)); math.Abs(float64(got)-float64(want)) >= delta {
					t.Errorf("typesCount[%q] = %d, want within %.1f of %d", k, got, delta, want)
				}
			}
		})
	}
}

func TestDiagOrProcGenerator_Random_EmptyFile_NoDate(t *testing.T) {
	ctx := context.Background()
	rand.Seed(1)

	emptyFilename := testwrite.BytesToFile(t, []byte(``))

	f := test.DataFiles[test.Test]
	f.Diagnoses = emptyFilename
	f.Procedures = emptyFilename

	c, err := config.LoadHL7Config(ctx, test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	data, err := config.LoadData(ctx, f, c)
	if err != nil {
		t.Fatalf("LoadData(%+v, %+v) failed with %v", f, c, err)
	}

	tclock := testclock.New(defaultDate)
	dg := &testdate.Generator{}

	tests := []struct {
		name      string
		g         *DiagOrProcGenerator
		input     *pathway.DiagnosisOrProcedure
		wantTypes []string
	}{{
		name: "Random Diagnosis but empty file",
		g:    NewDiagnosisGenerator(c, data, tclock, dg),
		input: &pathway.DiagnosisOrProcedure{
			Description: "RANDOM",
		},
		wantTypes: c.Diagnosis.Types,
	}, {
		name: "Random Procedure but empty file",
		g:    NewProcedureGenerator(c, data, tclock, dg),
		input: &pathway.DiagnosisOrProcedure{
			Description: "RANDOM",
		},
		wantTypes: c.Procedure.Types,
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := tc.g

			runs := float64(1000)
			typesCount := make(map[string]int)
			for i := 0; i < int(runs); i++ {
				got := g.RandomOrFromPathway(nil, tc.input)
				if got == nil {
					t.Fatalf("RandomOrFromPathway(nil, %+v) got nil, want not nil", tc.input)
				}
				if got.Description != nil {
					t.Errorf("RandomOrFromPathway(nil, %+v).Description got %+v, want nil", tc.input, got.Description)
				}

				if got, want := got.DateTime, ir.NewValidTime(defaultDate); got != want {
					t.Errorf("RandomOrFromPathway(nil, %+v).DateTime got %v, want %v", tc.input, got, want)
				}

				gotType := got.Type
				if !contains(tc.wantTypes, gotType) {
					t.Errorf("RandomOrFromPathway(nil, %+v).Type=%q, want one of %v", tc.input, gotType, tc.wantTypes)
				}
				typesCount[gotType]++
			}

			// Allow an error of 20% of the number of runs.
			delta := runs / 5
			for k, got := range typesCount {
				if want := runs / float64(len(tc.wantTypes)); math.Abs(float64(got)-float64(want)) >= delta {
					t.Errorf("typesCount[%q] = %d, want within %.1f of %d", k, got, delta, want)
				}
			}
		})
	}
}
