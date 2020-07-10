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
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/constants"
)

var (
	delay     = time.Hour * 5
	now       = NewValidTime(time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC))
	later     = NewValidTime(now.Add(delay))
	evenLater = NewValidTime(later.Add(delay))
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
			LocationHistory: []*LocationHistory{{
				Location: wardBed(1),
				Start:    now,
			}},
		})
		got = append(got, p.AddEncounter(now, constants.EncounterStatusPlanned, wardBed(1)))
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("AddEncounter(%v, %v) returned diff (-want +got):\n%s", now, constants.EncounterStatusPlanned, diff)
	}
}

func TestEncounter_EndEncounter(t *testing.T) {
	p := &PatientInfo{}

	ec1 := p.AddEncounter(now, constants.EncounterStatusPlanned, wardBed(1))
	ec2 := p.AddEncounter(later, constants.EncounterStatusInProgress, wardBed(2))
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
		LocationHistory: []*LocationHistory{{
			Location: wardBed(1),
			Start:    now,
			End:      later,
		}},
	}, {
		Status:      constants.EncounterStatusFinished,
		StatusStart: later,
		Start:       later,
		End:         later,
		StatusHistory: []*StatusHistory{
			{Status: constants.EncounterStatusInProgress, Start: later, End: later},
		},
		LocationHistory: []*LocationHistory{{
			Location: wardBed(2),
			Start:    later,
			End:      later,
		}},
	}}

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
	p.AddEncounter(now, constants.EncounterStatusInProgress, wardBed(1))
	want := p.AddEncounter(later, constants.EncounterStatusFinished, wardBed(2))

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

func TestEncounter_UpdateLocation(t *testing.T) {
	tests := []struct {
		name      string
		locations []*LocationHistory
		want      []*LocationHistory
	}{{
		name: "Add three locations",
		locations: []*LocationHistory{
			{
				Location: wardBed(3),
				Start:    now,
			}, {
				Location: wardBed(1),
				Start:    later,
			}, {
				Location: wardBed(2),
				Start:    evenLater,
			},
		},
		want: []*LocationHistory{
			{
				Location: wardBed(3),
				Start:    now,
				End:      later,
			}, {
				Location: wardBed(1),
				Start:    later,
				End:      evenLater,
			}, {
				Location: wardBed(2),
				Start:    evenLater,
			},
		},
	}, {
		name: "Update twice with same location",
		locations: []*LocationHistory{
			{
				Location: wardBed(1),
				Start:    now,
			}, {
				Location: wardBed(1), // Should not be added to location history.
				Start:    later,
			},
		},
		want: []*LocationHistory{
			{
				Location: wardBed(1),
				Start:    now,
			},
		},
	}, {
		name: "Update with nil location",
		locations: []*LocationHistory{
			{
				Location: wardBed(1),
				Start:    now,
			}, {
				// Nil location.
			},
		},
		want: []*LocationHistory{
			{
				Location: wardBed(1),
				Start:    now,
			},
		},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ec := &Encounter{}

			for _, l := range tc.locations {
				ec.UpdateLocation(l.Start, l.Location)
			}

			got := ec.LocationHistory
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("ec.LocationHistory returned locations diff (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEncounter_Text(t *testing.T) {
	tests := []struct {
		name      string
		encounter *Encounter
		want      string
	}{{
		name: "Finished",
		encounter: &Encounter{
			Status: constants.EncounterStatusFinished,
			Start:  now,
			End:    later,
		},
		want: "Status: finished\nActive from Mon Feb 12 00:00:00 2018 until Mon Feb 12 05:00:00 2018",
	}, {
		name: "In progress",
		encounter: &Encounter{
			Status: constants.EncounterStatusInProgress,
			Start:  now,
		},
		want: "Status: in-progress\nActive from Mon Feb 12 00:00:00 2018",
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.encounter.Text()
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("encounter.Text() returned diff (-want +got):\n%s", diff)
			}
		})
	}
}

func TestResult_Text(t *testing.T) {
	tests := []struct {
		name   string
		result *Result
		want   string
	}{{
		name: "All fields",
		result: &Result{
			TestName:     &CodedElement{Text: "TEST NAME"},
			Value:        "VALUE",
			Unit:         "UNIT",
			AbnormalFlag: "HIGH",
		},
		want: "TEST NAME: VALUE UNIT (HIGH)",
	}, {
		name: "Missing abnormal flag",
		result: &Result{
			TestName: &CodedElement{Text: "TEST NAME"},
			Value:    "VALUE",
			Unit:     "UNIT",
		},
		want: "TEST NAME: VALUE UNIT",
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.result.Text()
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("result.Text() returned diff (-want +got):\n%s", diff)
			}
		})
	}
}

func TestPatientLocation_Name(t *testing.T) {
	location := &PatientLocation{
		Bed:      "BED 1",
		Room:     "ROOM 2",
		Floor:    "5",
		Poc:      " ", // Should be omitted from the output.
		Facility: "FACILITY",
		Building: "BUILDING",
	}

	want := "BED 1, ROOM 2, 5, BUILDING, FACILITY"

	if got := location.Name(); got != want {
		t.Errorf("%v.Name() = %q, want %q", location, got, want)
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

func wardBed(i int) *PatientLocation {
	return &PatientLocation{
		Poc:          "Ward 1",
		Facility:     "Simulated Hospital",
		Building:     "Building-1",
		Floor:        "7",
		LocationType: "BED",
		Room:         "Room-1",
		Bed:          fmt.Sprintf("Bed %d", i),
	}
}
