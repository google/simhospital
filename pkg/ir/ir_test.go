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

	want := []*Encounter{{
		Status:      constants.EncounterStatusCancelled,
		StatusStart: later,
		Start:       now,
		End:         later,
		StatusHistory: []*StatusHistory{
			{Status: constants.EncounterStatusPlanned, Start: now, End: later},
		},
	}, {
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

func TestPatientInfo_AddOrderToEncounter(t *testing.T) {
	tests := []struct {
		name   string
		p      *PatientInfo
		orders []*Order
		want   []*Encounter
	}{{
		name: "Add Order to existing Encounter",
		p: &PatientInfo{
			Encounters: []*Encounter{{
				Status:      constants.EncounterStatusArrived,
				StatusStart: now,
				Start:       now,
				End:         NewInvalidTime(),
			}},
		},
		orders: []*Order{testOrder(), testOrder()},
		want: []*Encounter{{
			Status:      constants.EncounterStatusInProgress,
			StatusStart: now,
			Start:       now,
			End:         NewInvalidTime(),
			Orders:      []*Order{testOrder(), testOrder()},
			StatusHistory: []*StatusHistory{{
				Status: constants.EncounterStatusArrived,
				Start:  now,
				End:    now,
			}},
		}},
	}, {
		name:   "Multiple new Orders",
		p:      &PatientInfo{},
		orders: []*Order{testOrder(), testOrder()},
		want: []*Encounter{{
			Status:      constants.EncounterStatusFinished,
			StatusStart: later,
			Start:       now,
			End:         later,
			StatusHistory: []*StatusHistory{{
				Status: constants.EncounterStatusInProgress,
				Start:  now,
				End:    later,
			}},
			Orders: []*Order{testOrder()},
		}, {
			Status:      constants.EncounterStatusFinished,
			StatusStart: later,
			Start:       now,
			End:         later,
			StatusHistory: []*StatusHistory{{
				Status: constants.EncounterStatusInProgress,
				Start:  now,
				End:    later,
			}},
			Orders: []*Order{testOrder()},
		}},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, o := range tc.orders {
				tc.p.AddOrderToEncounter(o)
			}

			got := tc.p.Encounters
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("p.Encounters returned encounters diff (-want +got):\n%s", diff)
			}
		})
	}
}

func testOrder() *Order {
	return &Order{
		OrderProfile:                  &CodedElement{ID: "ORDER_PROFILE", Text: "ORDER_PROFILE"},
		Placer:                        "PLACER",
		Filler:                        "FILLER",
		OrderDateTime:                 now,
		CollectedDateTime:             now,
		ReceivedInLabDateTime:         later,
		ReportedDateTime:              later,
		MessageControlIDOriginalOrder: "6",
		NumberOfPreviousResults:       5,
		OrderControl:                  "ORDER_CONTROL",
		OrderStatus:                   "ORDER_STATIS",
		DiagnosticServID:              "SERV_ID",
		ResultsStatus:                 "RESULTS_STATUS",
		SpecimenSource:                "SPECIMEN_SOURCE",
		OrderingProvider: &Doctor{
			ID:        "ID",
			Surname:   "SURNAME",
			FirstName: "FIRST_NAME",
			Prefix:    "PREFIX",
			Specialty: "SPECIALTY",
		},
		NotesForORM: []string{"NOTES_1", "NOTES_2", "NOTES_3"},
		Results: []*Result{{
			TestName:            &CodedElement{ID: "TEST_NAME", Text: "TEST_NAME"},
			Value:               "VALUE",
			Unit:                "UNIT",
			ValueType:           "VALUE_TYPE",
			ObservationDateTime: later,
			Notes:               []string{"NOTES_1", "NOTES_2", "NOTES_3"},
			Range:               "RANGE",
			AbnormalFlag:        "ABNORMAL_FLAG",
			Status:              "STATUS",
			ClinicalNote: &ClinicalNote{
				DateTime:      now,
				DocumentTitle: "DOCUMENT_TITLE",
				DocumentType:  "DOCUMENT_TYPE",
				DocumentID:    "DOCUMENT_ID",
				Contents: []*ClinicalNoteContent{{
					ObservationDateTime: now,
					ContentType:         "CONTENT_TYPE",
					DocumentEncoding:    "DOCUMENT_ENCODING",
					DocumentContent:     "DOCUMENT_CONTENT",
				}},
			},
		}},
	}
}
