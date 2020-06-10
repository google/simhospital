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

package ir

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/constants"
)

var (
	delay = time.Hour * 5
	now   = NewValidTime(time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC))
	later = NewValidTime(now.Add(delay))
)

func TestPatientInfo_AddEncounter(t *testing.T) {
	p := &PatientInfo{}

	var want []*Encounter
	var got []*Encounter
	for i := 0; i < 5; i++ {
		want = append(want, &Encounter{
			Start:       now,
			Status:      constants.EncounterStatusPlanned,
			StatusStart: now,
		})
		got = append(got, p.AddEncounter(now, constants.EncounterStatusPlanned))
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("AddEncounter(%v, %v) returned diff (-want +got):\n%s", now, constants.EncounterStatusPlanned, diff)
	}
}

func TestEncounter_EndEncounter(t *testing.T) {
	p := &PatientInfo{}

	ec1 := p.AddEncounter(now, constants.EncounterStatusPlanned)
	ec2 := p.AddEncounter(later, constants.EncounterStatusInProgress)
	ec1.EndEncounter(later, constants.EncounterStatusCancelled)
	ec2.EndEncounter(later, constants.EncounterStatusFinished)

	want := []*Encounter{
		{
			Status:      constants.EncounterStatusCancelled,
			StatusStart: later,
			Start:       now,
			End:         later,
			StatusHistory: []*StatusHistory{
				{Status: constants.EncounterStatusPlanned, Start: now, End: later},
			},
		},
		{
			Status:      constants.EncounterStatusFinished,
			StatusStart: later,
			Start:       later,
			End:         later,
			StatusHistory: []*StatusHistory{
				{Status: constants.EncounterStatusInProgress, Start: later, End: later},
			},
		},
	}

	got := p.Encounters
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("p.Encounters returned encounters diff (-want +got):\n%s", diff)
	}
}

func TestPatientInfo_LatestEncounter(t *testing.T) {
	p := &PatientInfo{}

	if got := p.LatestEncounter(); got != nil {
		t.Error("p.LatestEncounter() is something, want <nil>")
	}

	// Populate Encounters slice with two encounters.
	p.AddEncounter(now, constants.EncounterStatusInProgress)
	want := p.AddEncounter(later, constants.EncounterStatusFinished)

	// We assert that both got and want point to the same struct.
	if got := p.LatestEncounter(); got != want {
		t.Errorf("p.LatestEncounter() = %p, want %p", got, want)
	}
}
