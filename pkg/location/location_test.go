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

package location_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/simhospital/pkg/ir"
	. "github.com/google/simhospital/pkg/location"
	"github.com/google/simhospital/pkg/test/testlocation"
	"github.com/google/simhospital/pkg/test/testwrite"
)

const aAndEID = "ED"

var (
	aAndEBed1 = &ir.PatientLocation{
		Poc:          "ED",
		Facility:     "Simulated Hospital",
		Building:     "Building-1",
		Floor:        "7",
		Room:         "Room-1",
		LocationType: "BED",
		// Beds are filled in order starting from 1.
		Bed: "Bed 1",
	}

	bed2      = "Bed 2"
	aAndEBed2 = &ir.PatientLocation{
		Poc:          "ED",
		Facility:     "Simulated Hospital",
		Building:     "Building-1",
		Floor:        "7",
		LocationType: "BED",
		Room:         "Room-1",
		Bed:          bed2,
	}
)

func TestNewManager(t *testing.T) {
	ctx := context.Background()
	loc := []byte(`
Ward 1:
  poc: Ward 1
  facility: Simulated Hospital
  building: Building-1
  floor: 7
  room: Room-1

ED:
  poc: ED
  facility: Simulated Hospital
  building: Building-1
  floor: 7
  room: Room-2
  type: ED

Ward 1 Room 2:
  poc: Ward 1
  facility: Simulated Hospital
  building: Building-1
  floor: 7
  room: Room-3`)

	noAandE := []byte(`
Ward 1:
  poc: Ward 1
  facility: Simulated Hospital
  building: Building-1
  floor: 7
  room: Room-1`)

	invalid := []byte(`
poc: Ward 1
facility: Simulated Hospital
building: Building-1
floor: 7
room: Room-1`)

	cases := []struct {
		name       string
		locContent []byte
		want       *Manager
		wantErr    bool
	}{
		{
			name:       "valid",
			locContent: loc,
			want: &Manager{
				RoomManagers: map[string]*RoomManager{
					"Ward 1": {
						Poc:      "Ward 1",
						Facility: "Simulated Hospital",
						Building: "Building-1",
						Floor:    "7",
						Room:     "Room-1",
						Type:     "BED",
					},
					"ED": {
						Poc:      "ED",
						Facility: "Simulated Hospital",
						Building: "Building-1",
						Floor:    "7",
						Room:     "Room-2",
						Type:     "ED",
					},
					"Ward 1 Room 2": {
						Poc:      "Ward 1",
						Facility: "Simulated Hospital",
						Building: "Building-1",
						Floor:    "7",
						Room:     "Room-3",
						Type:     "BED",
					},
				},
			},
		}, {
			name:       "invalid yml",
			locContent: invalid,
			wantErr:    true,
		}, {
			name:       "ED not defined",
			locContent: noAandE,
			wantErr:    true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fName := testwrite.BytesToFile(t, tc.locContent)

			got, err := NewManager(ctx, fName)
			gotErr := err != nil
			if gotErr != tc.wantErr {
				t.Errorf("NewManager(%s) got err %v, want err? %t", string(tc.locContent), err, tc.wantErr)
			}
			if gotErr || tc.wantErr {
				return
			}

			if diff := cmp.Diff(tc.want, got, cmpopts.IgnoreUnexported(RoomManager{})); diff != "" {
				t.Errorf("NewManager(%s) got diff (-want, +got):\n%s", string(tc.locContent), diff)
			}
		})
	}
}

func TestManagerOccupyAvailableBed(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name    string
		poc     string
		want    *ir.PatientLocation
		wantErr bool
	}{
		{
			name: "existing poc",
			poc:  aAndEID,
			want: aAndEBed1,
		}, {
			name:    "unknown poc",
			poc:     "unknown-ward",
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			locationManager := testlocation.NewLocationManager(ctx, t, aAndEID)

			got, err := locationManager.OccupyAvailableBed(tc.poc)
			gotErr := err != nil
			if gotErr != tc.wantErr {
				t.Fatalf("OccupyAvailableBed(%s) got err %v; want err? %t", aAndEID, err, tc.wantErr)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("OccupyAvailableBed(%s) got diff (-want, +got):\n%s", aAndEID, diff)
			}
		})
	}
}

func TestManagerOccupyAvailableBedReoccupyFreedBed(t *testing.T) {
	ctx := context.Background()
	locationManager := testlocation.NewLocationManager(ctx, t, aAndEID)
	want := aAndEBed1

	got, err := locationManager.OccupyAvailableBed(aAndEID)
	if err != nil {
		t.Fatalf("OccupyAvailableBed(%s) failed with %v", aAndEID, err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("OccupyAvailableBed(%s) got diff (-want, +got):\n%s", aAndEID, diff)
	}
	if gotCount, wantCount := locationManager.RoomManagers[aAndEID].OccupiedBeds(), 1; gotCount != wantCount {
		t.Errorf("RoomManagers[%s].OccupiedBeds()=%d, want %d", aAndEID, gotCount, wantCount)
	}

	if err := locationManager.FreeBed(got); err != nil {
		t.Fatalf("FreeBed(%s) failed with %v", got, err)
	}
	if gotCount, wantCount := locationManager.RoomManagers[aAndEID].OccupiedBeds(), 0; gotCount != wantCount {
		t.Errorf("RoomManagers[%s].OccupiedBeds()=%d, want %d", aAndEID, gotCount, wantCount)
	}

	// Since we just freed the bed, we should be able to occupy it again.
	got, err = locationManager.OccupyAvailableBed(aAndEID)
	if err != nil {
		t.Fatalf("OccupyAvailableBed(%s) failed with %v", aAndEID, err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("OccupyAvailableBed(%s) got diff (-want, +got):\n%s", aAndEID, diff)
	}
	if gotCount, wantCount := locationManager.RoomManagers[aAndEID].OccupiedBeds(), 1; gotCount != wantCount {
		t.Errorf("RoomManagers[%s].OccupiedBeds()=%d, want %d", aAndEID, gotCount, wantCount)
	}
}

func TestManagerOccupySpecificBed(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name    string
		poc     string
		bed     string
		want    *ir.PatientLocation
		wantErr bool
	}{
		{
			name: "existing poc",
			poc:  aAndEID,
			bed:  bed2,
			want: aAndEBed2,
		}, {
			name:    "unknown poc",
			poc:     "unknown-ward",
			bed:     bed2,
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			locationManager := testlocation.NewLocationManager(ctx, t, aAndEID)

			got, err := locationManager.OccupySpecificBed(tc.poc, tc.bed)
			gotErr := err != nil
			if gotErr != tc.wantErr {
				t.Fatalf("OccupySpecificBed(%s, %s) got err %v, want err? %t", tc.poc, tc.bed, err, tc.wantErr)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("OccupySpecificBed(%s, %s) got diff (-want, +got):\n%s", tc.poc, tc.bed, diff)
			}
		})
	}
}

func TestManagerOccupySpecificBedNoAvailableBed(t *testing.T) {
	ctx := context.Background()
	locationManager := testlocation.NewLocationManager(ctx, t, aAndEID)

	// Mark bed as occupied.
	if _, err := locationManager.OccupySpecificBed(aAndEID, bed2); err != nil {
		t.Fatalf("OccupySpecificBed(%s, %s) failed with %v", aAndEID, bed2, err)
	}

	if _, err := locationManager.OccupySpecificBed(aAndEID, bed2); err == nil {
		t.Errorf("OccupySpecificBed(%s, %s) got nil error, want non nil", aAndEID, bed2)
	}
}

func TestManagerOccupySpecificBedReoccupyFreedBed(t *testing.T) {
	ctx := context.Background()
	locationManager := testlocation.NewLocationManager(ctx, t, aAndEID)

	want := aAndEBed2

	// We first occupy a bed.
	got, err := locationManager.OccupySpecificBed(aAndEID, bed2)
	if err != nil {
		t.Fatalf("OccupySpecificBed(%s, %s) failed with %v", aAndEID, bed2, err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("OccupySpecificBed(%s, %s) got diff (-want, +got):\n%s", aAndEID, bed2, diff)
	}
	if gotCount, wantCount := locationManager.RoomManagers[aAndEID].OccupiedBeds(), 1; gotCount != wantCount {
		t.Errorf("RoomManagers[%s].OccupiedBeds()=%d, want %d", aAndEID, gotCount, wantCount)
	}

	if err := locationManager.FreeBed(got); err != nil {
		t.Errorf("FreeBed(%s) failed with %v", got, err)
	}
	if gotCount, wantCount := locationManager.RoomManagers[aAndEID].OccupiedBeds(), 0; gotCount != wantCount {
		t.Errorf("RoomManagers[%s].OccupiedBeds()=%d, want %d", aAndEID, gotCount, wantCount)
	}

	// Since we just freed it, we should be able to occupy it again.
	got, err = locationManager.OccupySpecificBed(aAndEID, bed2)
	if err != nil {
		t.Fatalf("OccupySpecificBed(%s, %s) failed with %v", aAndEID, bed2, err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("OccupySpecificBed(%s, %s) got diff (-want, +got):\n%s", aAndEID, bed2, diff)
	}
	if gotCount, wantCount := locationManager.RoomManagers[aAndEID].OccupiedBeds(), 1; gotCount != wantCount {
		t.Errorf("RoomManagers[%s].OccupiedBeds()=%d, want %d", aAndEID, gotCount, wantCount)
	}
}

func TestManagerFreeBed(t *testing.T) {
	ctx := context.Background()
	locationManager := testlocation.NewLocationManager(ctx, t, aAndEID)

	// We first occupy a bed.
	got, err := locationManager.OccupyAvailableBed(aAndEID)
	if err != nil {
		t.Fatalf("OccupyAvailableBed(%s) failed with %v", aAndEID, err)
	}
	if diff := cmp.Diff(aAndEBed1, got); diff != "" {
		t.Errorf("OccupyAvailableBed(%s) got diff (-want, +got):\n%s", aAndEID, diff)
	}

	// The first time we free it, everything is good.
	if err := locationManager.FreeBed(got); err != nil {
		t.Errorf("FreeBed(%s) failed with %v", got, err)
	}

	// The second time it's already free so we expect an error.
	if err := locationManager.FreeBed(got); err == nil {
		t.Errorf("FreeBed(%s) got nil err, want not nil error", got)
	}
}

func TestManagerFreeBedError(t *testing.T) {
	ctx := context.Background()
	locationManager := testlocation.NewLocationManager(ctx, t, aAndEID)

	// Get a location that is not a bed.
	aAndELoc := locationManager.GetAAndELocation()
	if aAndELoc == nil {
		t.Fatal("GetAAndELocation() got <nil>, want not nil")
	}
	if IsBed(aAndELoc) {
		t.Errorf("IsBed(%+v) is true, want false", aAndELoc)
	}

	cases := []struct {
		name string
		pl   *ir.PatientLocation
	}{
		{
			name: "not a bed",
			pl:   aAndELoc,
		}, {
			name: "location doesn't match",
			pl: &ir.PatientLocation{
				Poc:          "ED",
				Facility:     "Different Facility",
				Building:     "Building-1",
				Floor:        "7",
				LocationType: "BED",
				Room:         "Room-1",
				Bed:          bed2,
			},
		}, {
			name: "unknown location",
			pl: &ir.PatientLocation{
				Poc:          "Renal",
				Facility:     "Simulated Hospital",
				Building:     "Building-1",
				Floor:        "7",
				Room:         "Room-1",
				LocationType: "BED",
				Bed:          "Bed 1",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := locationManager.FreeBed(tc.pl); err == nil {
				t.Errorf("FreeBed(%v) got nil err, want not nil error", tc.pl)
			}
		})
	}
}

func TestManagerGetAAndELocation(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name    string
		manager *Manager
		want    *ir.PatientLocation
	}{
		{
			name:    "existing ED",
			manager: testlocation.NewLocationManager(ctx, t, aAndEID),
			want: &ir.PatientLocation{
				Poc:          aAndEID,
				Facility:     "Simulated Hospital",
				LocationType: "BED",
				Building:     "Building-1",
			},
		}, {
			name: "default ED",
			// When creating Manager using NewManager function, the validation would fail
			// if ED is not defined. It is however possible to create a Manager just by
			// defining structs, in which case ED may be missing.
			// GetAAndELocation() handles this case gracefully.
			manager: &Manager{
				RoomManagers: map[string]*RoomManager{
					"Renal": {
						Poc:      "Renal",
						Facility: "Simulated Hospital",
						Building: "Building-2",
						Floor:    "2",
						Room:     "Room-2",
					},
				},
			},
			want: &ir.PatientLocation{
				LocationType: "ED",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.manager.GetAAndELocation()
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("GetAAndELocation() got diff (-want, +got):\n%s", diff)
			}
		})
	}
}

func TestManagerMatches(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name     string
		location string
		pl       *ir.PatientLocation
		want     bool
		wantErr  bool
	}{
		{
			name:     "matching location",
			location: aAndEID,
			pl:       aAndEBed1,
			want:     true,
		}, {
			name:     "not matching location",
			location: aAndEID,
			pl: &ir.PatientLocation{
				Poc:          "Renal",
				Facility:     "Simulated Hospital",
				Building:     "Building-1",
				Floor:        "7",
				Room:         "Room-1",
				LocationType: "BED",
				Bed:          "Bed 1",
			},
			want: false,
		}, {
			name:     "nil PatientLocation",
			location: aAndEID,
			want:     false,
			wantErr:  true,
		}, {
			name:     "location does not exist",
			location: "Renal",
			pl:       aAndEBed1,
			want:     false,
			wantErr:  true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			manager := testlocation.NewLocationManager(ctx, t, aAndEID)
			got, err := manager.Matches(tc.location, tc.pl)
			gotErr := err != nil
			if gotErr != tc.wantErr {
				t.Fatalf("Matches(%s, %v) got err %v, want err? %t", tc.location, tc.pl, err, tc.wantErr)
			}

			if got != tc.want {
				t.Errorf("Matches(%s, %v) got %t, want %t", tc.location, tc.pl, got, tc.want)
			}
		})
	}
}

func TestIsBed(t *testing.T) {
	cases := []struct {
		name string
		pl   *ir.PatientLocation
		want bool
	}{
		{
			name: "is bed",
			pl: &ir.PatientLocation{
				Poc: "Ward",
				Bed: "Bed 1",
			},
			want: true,
		}, {
			name: "is not a bed",
			pl: &ir.PatientLocation{
				Poc: "Ward",
			},
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsBed(tc.pl)
			if got != tc.want {
				t.Errorf("IsBed(%v)=%t, want %t", tc.pl, got, tc.want)
			}
		})
	}
}
