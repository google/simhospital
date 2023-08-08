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
	"math/rand"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewDeterministicManager(t *testing.T) {
	steps := []Step{
		{Admission: &Admission{}},
		{Discharge: &Discharge{}},
	}
	pathways := map[string]Pathway{
		"pathway1": {Pathway: steps},
		"pathway2": {Pathway: steps},
	}
	cases := []struct {
		name    string
		order   []string
		wantErr bool
	}{{
		name:  "one pathway",
		order: []string{"pathway1"},
	}, {
		name:  "two pathways",
		order: []string{"pathway1", "pathway2"},
	}, {
		name:    "non-existent pathway",
		order:   []string{"pathway1", "non-existent"},
		wantErr: true,
	}, {
		name:    "nil",
		order:   nil,
		wantErr: true,
	}, {
		name:    "empty",
		order:   []string{},
		wantErr: true,
	}, {
		name:    "one empty item",
		order:   []string{""},
		wantErr: true,
	}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewDeterministicManager(pathways, tc.order)
			if gotErr := err != nil; gotErr != tc.wantErr {
				t.Fatalf("NewDeterministicManager(%v, %v) got err %v, want err? %t", pathways, tc.order, err, tc.wantErr)
			}
		})
	}
}

func TestDeterministicManager_GetPathway(t *testing.T) {
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

	cases := []struct {
		name        string
		pathwayName string
		order       []string
		want        *Pathway
		wantError   bool
	}{{
		name:        "valid pathway name",
		pathwayName: "pathway1",
		order:       []string{"pathway1", "pathway2"},
		want:        wantPathway1,
	}, {
		name:        "inexistent pathway name",
		pathwayName: "inexistent-pathway",
		order:       []string{"pathway1", "pathway2"},
		wantError:   true,
	}, {
		// "order" should not have any effect on GetPathway, it affects NextPathway only.
		name:        "valid pathway name not in order",
		order:       []string{"pathway2"},
		pathwayName: "pathway1",
		want:        wantPathway1,
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m, err := NewDeterministicManager(pathways, tc.order)
			if err != nil {
				t.Fatalf("NewDeterministicManager(%v, %v) failed with %v", pathways, tc.order, err)
			}
			got, err := m.GetPathway(tc.pathwayName)

			gotErr := err != nil
			if gotErr != tc.wantError {
				t.Errorf("manager.GetPathway(%s) got err %v, want err? %t", tc.pathwayName, err, tc.wantError)
			}
			if gotErr || tc.wantError {
				return
			}
			if diff := cmp.Diff(tc.want, got, cmpopts.IgnoreUnexported(Pathway{}, Step{})); diff != "" {
				t.Errorf("manager.GetPathway(%s) -want, +got:\n%s", tc.pathwayName, diff)
			}
		})
	}
}

func TestDeterministicManager_PathwayNames(t *testing.T) {
	pathway := Pathway{
		Pathway: []Step{
			{Admission: &Admission{}},
			{Result: &Results{}},
			{Discharge: &Discharge{}},
		},
	}

	cases := []struct {
		name     string
		pathways map[string]Pathway
		order    []string
		want     []string
	}{{
		name:     "one pathway included in order",
		pathways: map[string]Pathway{"pathway1": pathway, "pathway2": pathway},
		order:    []string{"pathway1"},
		want:     []string{"pathway1", "pathway2"},
	}, {
		name:     "all included in order",
		pathways: map[string]Pathway{"pathway1": pathway, "pathway2": pathway},
		order:    []string{"pathway1", "pathway2"},
		want:     []string{"pathway1", "pathway2"},
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m, err := NewDeterministicManager(tc.pathways, tc.order)
			if err != nil {
				t.Fatalf("NewDeterministicManager(%v, %v) failed with %v", tc.pathways, tc.order, err)
			}
			got := m.PathwayNames()
			if diff := cmp.Diff(tc.want, got, cmpopts.SortSlices(func(x, y string) bool { return strings.Compare(x, y) > 0 })); diff != "" {
				t.Errorf("manager.PathwayNames() -want, +got:\n%s", diff)
			}
		})
	}
}

func TestDeterministicManager_NextPathway(t *testing.T) {
	rand.Seed(1)

	steps := []Step{
		{Admission: &Admission{}},
		{Discharge: &Discharge{}},
	}
	pathways := map[string]Pathway{
		"pathway1": {Pathway: steps},
		"pathway2": {Pathway: steps},
		"pathway3": {Pathway: steps},
	}

	// nRuns is an arbitrary number small enough so that we can specify all test outputs, and large enough
	// so that we can test wrapping.
	nRuns := 7
	cases := []struct {
		name      string
		order     []string
		wantOrder []string
	}{{
		name:      "one pathway",
		order:     []string{"pathway1"},
		wantOrder: []string{"pathway1", "pathway1", "pathway1", "pathway1", "pathway1", "pathway1", "pathway1"},
	}, {
		name:      "two pathways",
		order:     []string{"pathway1", "pathway2"},
		wantOrder: []string{"pathway1", "pathway2", "pathway1", "pathway2", "pathway1", "pathway2", "pathway1"},
	}, {
		name:      "three pathways",
		order:     []string{"pathway1", "pathway2", "pathway3"},
		wantOrder: []string{"pathway1", "pathway2", "pathway3", "pathway1", "pathway2", "pathway3", "pathway1"},
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			manager, err := NewDeterministicManager(pathways, tc.order)
			if err != nil {
				t.Fatalf("NewDeterministicManager(%+v,%v) failed with %v", pathways, tc.order, err)
			}

			var got []string
			for i := 0; i < nRuns; i++ {
				pathway, err := manager.NextPathway()
				if err != nil {
					t.Errorf("manager.NextPathway() failed with %v", err)
				}
				got = append(got, pathway.Name())
			}
			if diff := cmp.Diff(tc.wantOrder, got); diff != "" {
				t.Errorf("manager.NextPathway() got diff:\n%s", diff)
			}
		})
	}
}
