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

package state

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/state/persist"
	"github.com/google/simhospital/pkg/test/teststate"
)

const (
	testID  = "1"
	testID2 = "2"
)

var (
	testInfo1 = &Patient{PatientInfo: &ir.PatientInfo{Person: &ir.Person{MRN: testID}}}
	testInfo2 = &Patient{PatientInfo: &ir.PatientInfo{Person: &ir.Person{MRN: testID2}}}
)

func TestPatientsMap_Put_DifferentIDs(t *testing.T) {
	tests := []struct {
		name string
		pm   *PatientsMap
	}{
		{name: "no syncer", pm: NewPatientsMap(nil, true)},
		{name: "with syncer", pm: NewPatientsMap(teststate.NewItemSyncer(), true)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pm.Put(testInfo1)
			tt.pm.Put(testInfo2)

			for _, want := range []*Patient{testInfo1, testInfo2} {
				// Confirm patient was added to internal map.
				wantID, err := want.ID()
				if err != nil {
					t.Fatalf("want.ID() failed with %v", err)
				}
				if got := tt.pm.m[wantID]; !cmp.Equal(want, got) {
					t.Errorf("pm.m[%q] = %v, want: %v", wantID, got, want)
				}
				// Confirm patient was added to synced store.
				if tt.pm.syncer != nil {
					if got, _ := tt.pm.syncer.LoadByID(wantID); !cmp.Equal(*want, got) {
						t.Errorf("syncer LoadByID(%q) = %v, want: %v", wantID, got, want)
					}
				}
			}
		})
	}
}

func TestPatientsMap_Put_DuplicateIDs(t *testing.T) {
	tests := []struct {
		name string
		pm   *PatientsMap
	}{
		{name: "no syncer", pm: NewPatientsMap(nil, true)},
		{name: "with syncer", pm: NewPatientsMap(teststate.NewItemSyncer(), true)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pm.Put(testInfo1)
			testInfo1ID, err := testInfo1.ID()
			if err != nil {
				t.Fatalf("testInfo1.ID() failed with %v", err)
			}
			second := &Patient{PatientInfo: &ir.PatientInfo{Person: &ir.Person{MRN: testInfo1ID}}}
			tt.pm.Put(second)

			for _, want := range []*Patient{testInfo1, second} {
				// Confirm patient was added to internal map.
				wantID, err := want.ID()
				if err != nil {
					t.Fatalf("want.ID() failed with %v", err)
				}
				if got := tt.pm.m[wantID]; !cmp.Equal(second, got) {
					t.Errorf("pm.m[%q] = %v, want: %v", wantID, got, second)
				}
				// Confirm patient was added to synced store.
				if tt.pm.syncer != nil {
					if got, _ := tt.pm.syncer.LoadByID(wantID); !cmp.Equal(*second, got) {
						t.Errorf("syncer LoadByID(%q) = %v, want: %v", wantID, got, second)
					}
				}
			}
		})
	}
}

func TestPatientsMap_Get_PatientInMap(t *testing.T) {
	pm := NewPatientsMap(teststate.NewItemSyncer(), true)
	pm.Put(testInfo1)

	testInfo1ID, err := testInfo1.ID()
	if err != nil {
		t.Fatalf("testInfo1.ID() failed with %v", err)
	}
	if got := pm.Get(testInfo1ID); !cmp.Equal(testInfo1, got) {
		t.Errorf("pm.Get(%q) = %v, want: %v", testInfo1ID, got, testInfo1)
	}
	if pm.syncer.(*teststate.ItemSyncer).WasRequested(testInfo1ID) {
		t.Errorf("syncer LoadByID(%q) was called, want: not called", testInfo1ID)
	}
}

func TestPatientsMap_Get_PatientInSyncer(t *testing.T) {
	testInfo1ID, err := testInfo1.ID()
	if err != nil {
		t.Fatalf("testInfo1.ID() failed with %v", err)
	}
	pm := NewPatientsMap(teststate.NewItemSyncer(), true)
	pm.Put(testInfo1)
	pm.Delete(testInfo1ID) // Delete from map only.

	if got := pm.Get(testInfo1ID); !cmp.Equal(testInfo1, got) {
		t.Errorf("pm.Get(%q) = %v, want: %v", testInfo1ID, got, testInfo1)
	}
	if !pm.syncer.(*teststate.ItemSyncer).WasRequested(testInfo1ID) {
		t.Errorf("syncer LoadByID(%q) was not called, want: called", testInfo1ID)
	}
	// Confirm that internal map is updated with patient returned from syncer.
	if got := pm.m[testInfo1ID]; !cmp.Equal(testInfo1, got) {
		t.Errorf("pm.m[%q] = %v, want: %v", testInfo1ID, got, testInfo1)
	}
}

func TestPatientsMap_Get_PatientNotFound(t *testing.T) {
	pm := NewPatientsMap(teststate.NewItemSyncer(), true)
	testInfo1ID, err := testInfo1.ID()
	if err != nil {
		t.Fatalf("testInfo1.ID() failed with %v", err)
	}
	if got := pm.Get(testInfo1ID); got != nil {
		t.Errorf("pm.Get(%q) = %v, want: nil", testInfo1ID, got)
	}
}

func TestPatientsMap_Delete(t *testing.T) {
	tests := []struct {
		deleteInSyncer bool
		deleteInMap    bool
		wantInSyncer   persist.MarshallableItem
		wantInMap      *Patient
	}{{
		deleteInSyncer: true,
		deleteInMap:    true,
		wantInSyncer:   nil,
		wantInMap:      nil,
	}, {
		deleteInSyncer: false,
		deleteInMap:    true,
		wantInSyncer:   *testInfo1,
		wantInMap:      nil,
	}, {
		deleteInSyncer: true,
		deleteInMap:    false,
		wantInSyncer:   nil,
		wantInMap:      testInfo1,
	}, {
		deleteInSyncer: false,
		deleteInMap:    false,
		wantInSyncer:   *testInfo1,
		wantInMap:      testInfo1,
	}}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("deleteInSyncer=%t_deleteInMap=%t", tc.deleteInSyncer, tc.deleteInMap), func(t *testing.T) {
			pm := NewPatientsMap(teststate.NewItemSyncerWithDelete(tc.deleteInSyncer), tc.deleteInMap)
			pm.Put(testInfo1)
			testInfo1ID, err := testInfo1.ID()
			if err != nil {
				t.Fatalf("testInfo1.ID() failed with %v", err)
			}
			if got := pm.Get(testInfo1ID); !cmp.Equal(testInfo1, got) {
				t.Errorf("pm.Get(%q) = %v, want: %v", testInfo1ID, got, testInfo1)
			}

			pm.Delete(testInfo1ID)

			if gotInMap := pm.m[testInfo1ID]; !cmp.Equal(tc.wantInMap, gotInMap) {
				t.Errorf("pm.m[%q] = %v, want: %v", testInfo1ID, gotInMap, tc.wantInMap)
			}
			gotInSyncer, err := pm.syncer.LoadByID(testInfo1ID)
			if err != nil {
				t.Fatalf("syncer LoadByID(%q) failed with %v", testInfo1ID, err)
			}
			if !cmp.Equal(tc.wantInSyncer, gotInSyncer) {
				t.Errorf("syncer LoadByID(%q) = %+v, want: %+v", testInfo1ID, gotInSyncer, tc.wantInSyncer)
			}
		})
	}
}

func TestPatientsMap_Len(t *testing.T) {
	pm := NewPatientsMap(teststate.NewItemSyncer(), true)
	testInfo1ID, err := testInfo1.ID()
	if err != nil {
		t.Fatalf("testInfo1.ID() failed with %v", err)
	}

	pm.Put(testInfo1)
	if got, want := pm.Len(), 1; got != want {
		t.Errorf("pm.Len() = %d, want: %d", got, want)
	}

	pm.Put(testInfo2)
	if got, want := pm.Len(), 2; got != want {
		t.Errorf("pm.Len() = %d, want: %d", got, want)
	}

	pm.Delete(testInfo1ID) // Delete from map only.
	if got, want := pm.Len(), 1; got != want {
		t.Errorf("pm.Len() = %d, want: %d", got, want)
	}
}
