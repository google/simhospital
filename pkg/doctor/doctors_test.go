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

package doctor

import (
	"context"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/test/testwrite"
)

var (
	singleDoctor = `
- id: "id-1"
  surname: "surname-1"
  firstname: "firstname-1"
  prefix: "prefix-1"
  specialty: "specialty-1"`

	twoDoctors = `
- id: "id-1"
  surname: "surname-1"
  firstname: "firstname-1"
  prefix: "prefix-1"
  specialty: "specialty-1"
- id: "id-2"
  surname: "surname-2"
  firstname: "firstname-2"
  prefix: "prefix-2"
  specialty: "specialty-2"`
)

func TestLoadDoctors(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "success - single doctor",
			input:   singleDoctor,
			wantErr: false,
		},
		{
			name:    "success - two doctors",
			input:   twoDoctors,
			wantErr: false,
		},
		{
			name: "inexistent field",
			input: `
- id: "id-1"
  inexistent-field: "surname-1"`,
			wantErr: true,
		},
		{
			name: "empty field",
			input: `
- id: "id-1"
  surname: ""
  firstname: "firstname-1"
  prefix: "prefix-1"
  specialty: "specialty-1"`,
			wantErr: true,
		},
		{
			name: "absent fields",
			input: `
- id: "id-1"
  surname: "surname-2"`,
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fName := testwrite.BytesToFile(t, []byte(tc.input))
			if _, err := LoadDoctors(ctx, fName); (err == nil) == tc.wantErr {
				t.Errorf("LoadDoctors(%s) got err %v; want err: %t", tc.input, err, tc.wantErr)
			}
		})
	}
}

func TestDoctorsGet(t *testing.T) {
	ctx := context.Background()
	fName := testwrite.BytesToFile(t, []byte(singleDoctor))
	d, err := LoadDoctors(ctx, fName)
	if err != nil {
		t.Fatalf("LoadDoctors() failed with %v", err)
	}
	if got, want := len(d.m), 1; got != want {
		t.Fatalf("len(d.m)=%v, want %v", got, want)
	}

	testDoctor := &ir.Doctor{
		ID:        "id-1",
		Surname:   "surname-1",
		FirstName: "firstname-1",
		Prefix:    "prefix-1",
		Specialty: "specialty-1",
	}
	var nilDoctor *ir.Doctor

	tests := []struct {
		name       string
		f          func() *ir.Doctor
		wantDoctor *ir.Doctor
	}{{
		name:       "by id",
		f:          func() *ir.Doctor { return d.GetByID("id-1") },
		wantDoctor: testDoctor,
	}, {
		name:       "by id but not in map",
		f:          func() *ir.Doctor { return d.GetByID("NotInMap") },
		wantDoctor: nilDoctor,
	}, {
		name:       "by name",
		f:          func() *ir.Doctor { return d.GetByName("firstname-1", "surname-1") },
		wantDoctor: testDoctor,
	}, {
		name:       "by name but not in map",
		f:          func() *ir.Doctor { return d.GetByName("NotInMap", "surname-1") },
		wantDoctor: nilDoctor,
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			doctor := tc.f()
			if diff := cmp.Diff(tc.wantDoctor, doctor); diff != "" {
				t.Errorf("Doctor got diff (-want, +got):\n%s", diff)
			}
		})
	}
}

func TestDoctorsAdd(t *testing.T) {
	ctx := context.Background()
	fName := testwrite.BytesToFile(t, []byte(singleDoctor))
	d, err := LoadDoctors(ctx, fName)
	if err != nil {
		t.Fatalf("LoadDoctors(%s) failed with %v", singleDoctor, err)
	}

	doctorToAdd := &ir.Doctor{
		ID:        "id-2",
		Surname:   "surname-2",
		FirstName: "firstname-2",
		Prefix:    "prefix-2",
		Specialty: "specialty-2",
	}
	if err := d.Add(doctorToAdd); err != nil {
		t.Fatalf("Add(%+v) failed with %v", doctorToAdd, err)
	}
	if err := d.Add(doctorToAdd); err == nil {
		t.Errorf("Add(%+v) got nil err; want non-nil err", doctorToAdd)
	}

	gotDoctor := d.GetByID("id-2")
	if diff := cmp.Diff(doctorToAdd, gotDoctor); diff != "" {
		t.Errorf(`GetByID("id-2") got diff (-want, +got):\n%s`, diff)
	}
}

func TestDoctorsGetRandomDoctor(t *testing.T) {
	ctx := context.Background()
	// Somewhat arbitrary number of doctors and runs, but chosen in a way that the probability that
	// all doctors are picked is large: small number of doctors, large number of runs.
	nDoctors := 2
	runs := 1000

	fName := testwrite.BytesToFile(t, []byte(twoDoctors))
	d, err := LoadDoctors(ctx, fName)
	if err != nil {
		t.Fatalf("LoadDoctors(%s) failed with %v", twoDoctors, err)
	}

	pickedIDs := map[string]int{}
	for i := 0; i < runs; i++ {
		randomDoctor := d.GetRandomDoctor()
		if randomDoctor == nil {
			t.Error("GetRandomDoctor() got <nil>; want not nil")
			continue
		}
		pickedIDs[randomDoctor.ID]++
	}

	if got := len(pickedIDs); got != nDoctors {
		t.Errorf("len(pickedIDs) = %d, want %d items", got, nDoctors)
	}
	delta := float64(runs / 2)
	for k, got := range pickedIDs {
		if want := runs / len(pickedIDs); math.Abs(float64(got)-float64(want)) >= delta {
			t.Errorf("pickedIDs[%q] = %d, want within %.1f of %d", k, got, delta, want)
		}
	}
}
