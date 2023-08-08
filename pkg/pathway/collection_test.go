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

package pathway

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCollection_GetPathway(t *testing.T) {
	steps1 := []Step{
		{Admission: &Admission{}},
		{Result: &Results{}},
		{Discharge: &Discharge{}},
	}
	steps2 := []Step{
		{Admission: &Admission{}},
		{Discharge: &Discharge{}},
	}
	pathways := map[string]Pathway{
		"pathway1": {Pathway: steps1},
		"pathway2": {Pathway: steps2},
	}
	wantPathway1 := &Pathway{
		Persons: &Persons{"main-patient": {}},
		Pathway: steps1,
	}
	wantPathway2 := &Pathway{
		Persons: &Persons{"main-patient": {}},
		Pathway: steps2,
	}

	cases := []struct {
		name        string
		pathwayName string
		want        *Pathway
		wantError   bool
	}{{
		name:        "valid pathway name pathway1",
		pathwayName: "pathway1",
		want:        wantPathway1,
	}, {
		name:        "valid pathway name pathway2",
		pathwayName: "pathway2",
		want:        wantPathway2,
	}, {
		name:        "inexistent pathway name",
		pathwayName: "inexistent-pathway",
		wantError:   true,
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c, err := NewCollection(pathways)
			if err != nil {
				t.Fatalf("NewCollection(%v) failed with %v", pathways, err)
			}
			got, err := c.GetPathway(tc.pathwayName)

			gotErr := err != nil
			if gotErr != tc.wantError {
				t.Errorf("collection.GetPathway(%s) got err %v, want err? %t", tc.pathwayName, err, tc.wantError)
			}
			if gotErr || tc.wantError {
				return
			}
			if diff := cmp.Diff(tc.want, got, cmpopts.IgnoreUnexported(Pathway{}, Step{})); diff != "" {
				t.Errorf("collection.GetPathway(%s) -want, +got:\n%s", tc.pathwayName, diff)
			}
		})
	}
}

func TestCollection_PathwayNames(t *testing.T) {
	pathway := Pathway{
		Pathway: []Step{
			{Admission: &Admission{}},
			{Discharge: &Discharge{}},
		},
	}
	pathways := map[string]Pathway{"pathway1": pathway, "pathway2": pathway, "pathway0": pathway}

	c, err := NewCollection(pathways)
	if err != nil {
		t.Fatalf("NewCollection(%v) failed with %v", pathways, err)
	}
	got := c.PathwayNames()
	want := []string{"pathway0", "pathway1", "pathway2"}
	if diff := cmp.Diff(want, got, cmpopts.SortSlices(func(x, y string) bool { return strings.Compare(x, y) > 0 })); diff != "" {
		t.Errorf("collection.PathwayNames() -want, +got:\n%s", diff)
	}
}

func TestCollection_Pathways(t *testing.T) {
	steps1 := []Step{
		{Admission: &Admission{}},
		{Result: &Results{}},
		{Discharge: &Discharge{}},
	}
	steps2 := []Step{
		{Admission: &Admission{}},
		{Discharge: &Discharge{}},
	}
	pathways := map[string]Pathway{
		"pathway1": {Pathway: steps1},
		"pathway2": {Pathway: steps2},
	}

	c, err := NewCollection(pathways)
	if err != nil {
		t.Fatalf("NewCollection(%v) failed with %v", pathways, err)
	}
	got := c.Pathways()
	want := map[string]Pathway{
		"pathway1": Pathway{
			Persons: &Persons{"main-patient": {}},
			Pathway: steps1,
		},
		"pathway2": Pathway{
			Persons: &Persons{"main-patient": {}},
			Pathway: steps2,
		},
	}

	if diff := cmp.Diff(want, got, cmpopts.IgnoreUnexported(Pathway{}, Step{})); diff != "" {
		t.Errorf("collection.Pathways() -want, +got:\n%s", diff)
	}
}
