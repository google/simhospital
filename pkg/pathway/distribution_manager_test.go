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
	"math"
	"math/rand"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestDistributionManager_GetPathway(t *testing.T) {
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
		include     []string
		exclude     []string
		want        *Pathway
		wantError   bool
	}{{
		name:        "valid pathway name",
		pathwayName: "pathway1",
		want:        wantPathway1,
	}, {
		name:        "inexistent pathway name",
		pathwayName: "inexistent-pathway",
		wantError:   true,
	}, {
		// include and exclude should not have any effect on GetPathway, they affect NextPathway only.
		name:        "excluded pathway name",
		exclude:     []string{"pathway2"},
		pathwayName: "pathway2",
		want:        wantPathway2,
	}, {
		name:        "not included pathway name",
		include:     []string{"pathway1"},
		pathwayName: "pathway2",
		want:        wantPathway2,
	}, {
		name:        "all excluded",
		exclude:     []string{".*"},
		pathwayName: "pathway2",
		want:        wantPathway2,
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m, err := NewDistributionManager(pathways, tc.include, tc.exclude)
			if err != nil {
				t.Fatalf("NewDistributionManager(%v, %v, %v) failed with %v", pathways, tc.include, tc.exclude, err)
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

func TestDistributionManager_PathwayNames(t *testing.T) {
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
		include  []string
		exclude  []string
		want     []string
	}{{
		name:     "one pathway",
		pathways: map[string]Pathway{"pathway1": pathway},
		want:     []string{"pathway1"},
	}, {
		name:     "two pathways",
		pathways: map[string]Pathway{"pathway1": pathway, "pathway2": pathway},
		want:     []string{"pathway1", "pathway2"},
	}, {
		// include and exclude should not have any effect on GetPathway, they affect NextPathway only.
		name:     "excluded pathway name",
		pathways: map[string]Pathway{"pathway1": pathway, "pathway2": pathway},
		exclude:  []string{"pathway2"},
		want:     []string{"pathway1", "pathway2"},
	}, {
		name:     "not included pathway name",
		pathways: map[string]Pathway{"pathway1": pathway, "pathway2": pathway},
		include:  []string{"pathway1"},
		want:     []string{"pathway1", "pathway2"},
	}, {
		name:     "all excluded",
		pathways: map[string]Pathway{"pathway1": pathway, "pathway2": pathway},
		exclude:  []string{".*"},
		want:     []string{"pathway1", "pathway2"},
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m, err := NewDistributionManager(tc.pathways, tc.include, tc.exclude)
			if err != nil {
				t.Fatalf("NewDistributionManager(%v, %v, %v) failed with %v", tc.pathways, tc.include, tc.exclude, err)
			}
			got := m.PathwayNames()
			if diff := cmp.Diff(tc.want, got, cmpopts.SortSlices(func(x, y string) bool { return strings.Compare(x, y) > 0 })); diff != "" {
				t.Errorf("manager.PathwayNames() -want, +got:\n%s", diff)
			}
		})
	}
}

func TestDistributionManager_NextPathway(t *testing.T) {
	rand.Seed(1)

	pathway1, pathway2, pathway3 := "pathway1", "pathway2", "pathway3"
	cases := []struct {
		name        string
		include     []string
		exclude     []string
		wantErrNew  bool
		wantErrNext bool
		percentages map[string]*Percentage
		wantFreq    map[string]float64
	}{{
		name:        "100% across 2 pathways",
		percentages: map[string]*Percentage{pathway1: NewPercentage(80), pathway2: NewPercentage(20)},
		wantFreq:    map[string]float64{pathway1: 80, pathway2: 20},
	}, {
		name:        "100% across 2 pathways, include explicit",
		percentages: map[string]*Percentage{pathway1: NewPercentage(80), pathway2: NewPercentage(20)},
		include:     []string{"pathway1", "pathway2"},
		wantFreq:    map[string]float64{pathway1: 80, pathway2: 20},
	}, {
		name:        "100% across 2 pathways, include regexp",
		percentages: map[string]*Percentage{pathway1: NewPercentage(80), pathway2: NewPercentage(20)},
		include:     []string{"pathway.*"},
		wantFreq:    map[string]float64{pathway1: 80, pathway2: 20},
	}, {
		name:        "100% across 2 pathways, include match-all regexp",
		percentages: map[string]*Percentage{pathway1: NewPercentage(80), pathway2: NewPercentage(20)},
		include:     []string{".*"},
		wantFreq:    map[string]float64{pathway1: 80, pathway2: 20},
	}, {
		name:        "100% across 2 pathways, include match-all full regexp",
		percentages: map[string]*Percentage{pathway1: NewPercentage(80), pathway2: NewPercentage(20)},
		include:     []string{"^.*$"},
		wantFreq:    map[string]float64{pathway1: 80, pathway2: 20},
	}, {
		name:        "100% across 2 pathways, include one",
		percentages: map[string]*Percentage{pathway1: NewPercentage(80), pathway2: NewPercentage(20)},
		include:     []string{"pathway1"},
		wantFreq:    map[string]float64{pathway1: 100},
	}, {
		name:        "100% across 2 pathways, include does not match any",
		percentages: map[string]*Percentage{pathway1: NewPercentage(80), pathway2: NewPercentage(20)},
		include:     []string{"does not match any"},
		wantErrNext: true,
	}, {
		name:        "100% across 2 pathways, exclude one",
		percentages: map[string]*Percentage{pathway1: NewPercentage(80), pathway2: NewPercentage(20)},
		exclude:     []string{"pathway1"},
		wantFreq:    map[string]float64{pathway2: 100},
	}, {
		name:        "100% across 2 pathways, exclude all",
		percentages: map[string]*Percentage{pathway1: NewPercentage(80), pathway2: NewPercentage(20)},
		exclude:     []string{"pathway1", "pathway2"},
		wantErrNext: true,
	}, {
		name:        "100% across 2 pathways, exclude regexp",
		percentages: map[string]*Percentage{pathway1: NewPercentage(80), pathway2: NewPercentage(20)},
		exclude:     []string{"pathway.*"},
		wantErrNext: true,
	}, {
		name:        "100% across 2 pathways, exclude match-all regexp",
		percentages: map[string]*Percentage{pathway1: NewPercentage(80), pathway2: NewPercentage(20)},
		exclude:     []string{".*"},
		wantErrNext: true,
	}, {
		name:        "100% across 2 pathways, exclude match-all full regexp",
		percentages: map[string]*Percentage{pathway1: NewPercentage(80), pathway2: NewPercentage(20)},
		exclude:     []string{"^.*$"},
		wantErrNext: true,
	}, {
		name:        "100% across 2 pathways, exclude does not match any",
		percentages: map[string]*Percentage{pathway1: NewPercentage(80), pathway2: NewPercentage(20)},
		exclude:     []string{"does not match anything"},
		wantFreq:    map[string]float64{pathway1: 80, pathway2: 20},
	}, {
		name:        "100% across 2 pathways, include and exclude excludes",
		percentages: map[string]*Percentage{pathway1: NewPercentage(80), pathway2: NewPercentage(20)},
		include:     []string{"pathway1", "pathway2"},
		exclude:     []string{"pathway1"},
		wantFreq:    map[string]float64{pathway2: 100},
	}, {
		name:        "invalid include",
		percentages: map[string]*Percentage{pathway1: NewPercentage(80), pathway2: NewPercentage(20)},
		include:     []string{"[0-9]++"},
		wantErrNew:  true,
	}, {
		name:        "invalid exclude",
		percentages: map[string]*Percentage{pathway1: NewPercentage(80), pathway2: NewPercentage(20)},
		exclude:     []string{"[0-9]++"},
		wantErrNew:  true,
	}, {
		name:        "percentage doesn't sum up to 100",
		percentages: map[string]*Percentage{pathway1: NewPercentage(60), pathway2: NewPercentage(20)},
		wantFreq:    map[string]float64{pathway1: 60, pathway2: 20},
	}, {
		name:        "remaining percentage split between pathways",
		percentages: map[string]*Percentage{pathway1: NewPercentage(50), pathway2: nil, pathway3: nil},
		wantFreq:    map[string]float64{pathway1: 50, pathway2: 25, pathway3: 25},
	}, {
		name:        "default percentage is 1 if no budget remaining",
		percentages: map[string]*Percentage{pathway1: NewPercentage(100), pathway2: nil, pathway3: nil},
		wantFreq:    map[string]float64{pathway1: 100, pathway2: 1, pathway3: 1},
	}, {
		name:        "pathway with negative or zero percentage is not run",
		percentages: map[string]*Percentage{pathway1: NewPercentage(100), pathway2: NewPercentage(0), pathway3: NewPercentage(-1)},
		wantFreq:    map[string]float64{pathway1: 100},
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			pathways := map[string]Pathway{}
			for pname, perc := range tc.percentages {
				pathways[pname] = Pathway{
					Percentage: perc,
					Pathway: []Step{
						{Admission: &Admission{}},
						{Discharge: &Discharge{}},
					},
				}
			}
			manager, err := NewDistributionManager(pathways, tc.include, tc.exclude)
			if gotErr := err != nil; gotErr != tc.wantErrNew {
				t.Fatalf("NewDistributionManager(%+v, %v, %v) got err %v, want err? %t", pathways, tc.include, tc.exclude, err, tc.wantErrNew)
			}
			if err != nil || tc.wantErrNew {
				return
			}

			gotFreq := make(map[string]int)
			runs := float64(10000)
			for i := 1; i <= int(runs); i++ {
				pathway, err := manager.NextPathway()
				if gotErr := err != nil; gotErr != tc.wantErrNext {
					t.Fatalf("manager.NextPathway() got err %v, want err? %t", err, tc.wantErrNext)
				}
				if err != nil || tc.wantErrNext {
					return
				}
				gotFreq[pathway.Name()]++
				if _, ok := tc.wantFreq[pathway.Name()]; !ok {
					t.Errorf("manager.NextPathway() pathway name %q, want one of %v", pathway.Name(), tc.wantFreq)
				}
			}
			// Allow an error of 1% of the number of runs.
			delta := runs / 100
			allFreq := 0.0
			for _, v := range tc.wantFreq {
				allFreq += v
			}
			for k, freq := range tc.wantFreq {
				if want := (freq / allFreq) * runs; math.Abs(float64(gotFreq[k])-want) >= delta {
					t.Errorf("gotFreq[%q] = %d, want within %v of %v", k, gotFreq[k], delta, want)
				}
			}
		})
	}
}
