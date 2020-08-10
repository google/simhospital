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

package hospital_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/simhospital/pkg/constants"
	. "github.com/google/simhospital/pkg/hospital"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/pathway"
	"github.com/google/simhospital/pkg/test/testresource"
)

// TestGenerateResources verifies the interaction between the `GenerateResources`
// step in the pathway and the `ResourceWriter`.
// It only tests that the contents of the `Person` field has been populated correctly,
// as this is all that is necessary to ensure that writing is happening as intended.
// Population of the remaining fields is covered by other tests.
func TestGenerateResources(t *testing.T) {
	ctx := context.Background()
	testPerson := pathway.Person{
		FirstName: "FIRST_NAME_1",
		Surname:   "SURNAME",
		MRN:       "MRN-NUMBER",
		NHS:       "NHS-NUMBER",
		Address: &pathway.Address{
			FirstLine:  "FIRST_LINE",
			SecondLine: "SECOND_LINE",
			City:       "CITY",
			Postcode:   "POSTCODE",
			Country:    "COUNTRY",
			Type:       "TYPE",
		},
		DateOfBirth: &now,
		Gender:      "M",
	}

	testPerson2 := testPerson
	testPerson2.FirstName = "FIRST_NAME_2"

	testPerson3 := testPerson
	testPerson3.FirstName = "FIRST_NAME_3"

	tests := []struct {
		name    string
		pathway pathway.Pathway
		want    []*ir.Person
	}{{
		name: "Admission",
		pathway: pathway.Pathway{
			Persons: &pathway.Persons{pathway.PatientID(testPerson.MRN): testPerson},
			Pathway: []pathway.Step{
				{Admission: &pathway.Admission{Loc: testLocAE}},
				{GenerateResources: &pathway.GenerateResources{}},
			},
		},
		want: []*ir.Person{pathwayPersonToIRPerson(testPerson)},
	}, {
		name: "Update person before and after",
		pathway: pathway.Pathway{
			Persons: &pathway.Persons{pathway.PatientID(testPerson.MRN): testPerson},
			Pathway: []pathway.Step{
				{Admission: &pathway.Admission{Loc: testLocAE}},
				{UpdatePerson: &pathway.UpdatePerson{
					Person: &pathway.Person{FirstName: "FIRST_NAME_2"},
				}},
				{GenerateResources: &pathway.GenerateResources{}},
				{UpdatePerson: &pathway.UpdatePerson{
					Person: &pathway.Person{FirstName: "FIRST_NAME_3"},
				}},
			},
		},
		want: []*ir.Person{pathwayPersonToIRPerson(testPerson2)},
	}, {
		name: "Two GenerateResources",
		pathway: pathway.Pathway{
			Persons: &pathway.Persons{pathway.PatientID(testPerson.MRN): testPerson},
			Pathway: []pathway.Step{
				{Admission: &pathway.Admission{Loc: testLocAE}},
				{UpdatePerson: &pathway.UpdatePerson{
					Person: &pathway.Person{FirstName: "FIRST_NAME_2"},
				}},
				{GenerateResources: &pathway.GenerateResources{}},
				{UpdatePerson: &pathway.UpdatePerson{
					Person: &pathway.Person{FirstName: "FIRST_NAME_3"},
				}},
				{GenerateResources: &pathway.GenerateResources{}},
			},
		},
		want: []*ir.Person{pathwayPersonToIRPerson(testPerson2), pathwayPersonToIRPerson(testPerson3)},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rw := testresource.NewWriter()

			pathways := map[string]pathway.Pathway{
				testPathwayName: tc.pathway,
			}

			hospital := newHospital(ctx, t, Config{ResourceWriter: rw}, pathways)
			defer hospital.Close()

			if err := hospital.StartNextPathway(); err != nil {
				t.Fatalf("StartNextPathway() failed with %v", err)
			}

			hospital.ConsumeQueues(ctx, t)

			if got, want := len(rw.Resources), len(tc.want); got != want {
				t.Fatalf("len(rw.Resources) = %d, want %v", got, want)
			}

			// We specify options to equate nil with empty slices and ignore unstable fields.
			// These fields are populated randomly by a generator, and it is not possible to
			// (easily) specify what data to be generated from a test.
			opts := []cmp.Option{
				cmpopts.EquateEmpty(),
				cmpopts.IgnoreFields(ir.Person{}, "Prefix", "Suffix", "PhoneNumber",
					"DeathIndicator", "Degree", "Ethnicity", "DateOfDeath", "MiddleName"),
			}

			var got []*ir.Person
			for _, r := range rw.Resources {
				got = append(got, r.Person)
			}
			if diff := cmp.Diff(tc.want, got, opts...); diff != "" {
				t.Errorf("StartNextPathway() resources returned diff (-want +got):\n%s", diff)
			}
		})
	}
}

func pathwayPersonToIRPerson(p pathway.Person) *ir.Person {
	return &ir.Person{
		FirstName: p.FirstName,
		Surname:   string(p.Surname),
		MRN:       p.MRN,
		NHS:       p.NHS,
		Address: &ir.Address{
			FirstLine:  string(p.Address.FirstLine),
			SecondLine: string(p.Address.SecondLine),
			City:       string(p.Address.City),
			Country:    string(p.Address.Country),
			PostalCode: string(p.Address.Postcode),
			Type:       string(p.Address.Type),
		},
		Gender: string(p.Gender),
		Birth:  ir.NewValidTime(*p.DateOfBirth),
	}
}

func TestEncounters(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		pathway pathway.Pathway
		want    []*ir.Encounter
	}{{
		name: "PendingAdmission Admission PendingDischarge Discharge",
		pathway: pathway.Pathway{
			Pathway: []pathway.Step{
				{PendingAdmission: &pathway.PendingAdmission{Loc: testLocAE, ExpectedAdmissionTimeFromNow: &delay}},
				{Delay: &pathway.Delay{From: delay, To: delay}},
				{Admission: &pathway.Admission{Loc: testLocAE}},
				{PendingDischarge: &pathway.PendingDischarge{ExpectedDischargeTimeFromNow: &delay}},
				{Delay: &pathway.Delay{From: delay, To: delay}},
				{Discharge: &pathway.Discharge{}},
				{GenerateResources: &pathway.GenerateResources{}},
			}},
		want: []*ir.Encounter{{
			Start:       ir.NewValidTime(later),
			End:         ir.NewValidTime(evenLater),
			Status:      constants.EncounterStatusFinished,
			StatusStart: ir.NewValidTime(evenLater),
			LocationHistory: []*ir.LocationHistory{{
				Location: aAndEBed(1),
				Start:    ir.NewValidTime(later),
				End:      ir.NewValidTime(evenLater),
			}},
			StatusHistory: []*ir.StatusHistory{{
				Start:  ir.NewValidTime(later),
				End:    ir.NewValidTime(later),
				Status: constants.EncounterStatusPlanned,
			}, {
				Start:  ir.NewValidTime(later),
				End:    ir.NewValidTime(evenLater),
				Status: constants.EncounterStatusArrived,
			}, {
				Start:  ir.NewValidTime(evenLater),
				End:    ir.NewValidTime(evenLater),
				Status: constants.EncounterStatusPlanned,
			}},
		}},
	}, {
		name: "Admission and Discharge twice",
		pathway: pathway.Pathway{
			Pathway: []pathway.Step{
				{Admission: &pathway.Admission{Loc: testLocAE}},
				{Delay: &pathway.Delay{From: delay, To: delay}},
				{Discharge: &pathway.Discharge{}},
				{Admission: &pathway.Admission{Loc: testLoc}},
				{Delay: &pathway.Delay{From: delay, To: delay}},
				{Discharge: &pathway.Discharge{DischargeTime: &evenLater}},
				{GenerateResources: &pathway.GenerateResources{}},
			}},
		want: []*ir.Encounter{{
			Start:       ir.NewValidTime(now),
			End:         ir.NewValidTime(later),
			Status:      constants.EncounterStatusFinished,
			StatusStart: ir.NewValidTime(later),
			LocationHistory: []*ir.LocationHistory{{
				Location: aAndEBed(1),
				Start:    ir.NewValidTime(now),
				End:      ir.NewValidTime(later),
			}},
			StatusHistory: []*ir.StatusHistory{{
				Start:  ir.NewValidTime(now),
				End:    ir.NewValidTime(later),
				Status: constants.EncounterStatusArrived,
			}},
		}, {
			Start:       ir.NewValidTime(later),
			End:         ir.NewValidTime(evenLater),
			Status:      constants.EncounterStatusFinished,
			StatusStart: ir.NewValidTime(evenLater),
			LocationHistory: []*ir.LocationHistory{{
				Location: wardBed(1),
				Start:    ir.NewValidTime(later),
				End:      ir.NewValidTime(evenLater),
			}},
			StatusHistory: []*ir.StatusHistory{{
				Start:  ir.NewValidTime(later),
				End:    ir.NewValidTime(evenLater),
				Status: constants.EncounterStatusArrived,
			}},
		}},
	}, {
		name: "PendingAdmission twice",
		pathway: pathway.Pathway{
			Pathway: []pathway.Step{
				{PendingAdmission: &pathway.PendingAdmission{Loc: testLoc, ExpectedAdmissionTimeFromNow: &delay}},
				{PendingAdmission: &pathway.PendingAdmission{Loc: testLoc, ExpectedAdmissionTimeFromNow: &delay}},
				{Delay: &pathway.Delay{From: delay, To: delay}},
				{Admission: &pathway.Admission{Loc: testLoc}},
				{GenerateResources: &pathway.GenerateResources{}},
			}},
		want: []*ir.Encounter{{
			Start:       ir.NewValidTime(later),
			End:         ir.NewInvalidTime(),
			Status:      constants.EncounterStatusPlanned,
			StatusStart: ir.NewValidTime(later),
			IsPending:   true,
		}, {
			Start:       ir.NewValidTime(later),
			End:         ir.NewInvalidTime(),
			Status:      constants.EncounterStatusArrived,
			StatusStart: ir.NewValidTime(later),
			LocationHistory: []*ir.LocationHistory{{
				Location: wardBed(2),
				Start:    ir.NewValidTime(later),
				End:      ir.NewInvalidTime(),
			}},
			StatusHistory: []*ir.StatusHistory{{
				Start:  ir.NewValidTime(later),
				End:    ir.NewValidTime(later),
				Status: constants.EncounterStatusPlanned,
			}},
		}},
	}, {
		name: "Multiple Orders",
		pathway: pathway.Pathway{
			Pathway: []pathway.Step{
				{Admission: &pathway.Admission{Loc: testLoc}},
				{Order: &pathway.Order{OrderID: "ORDER_ID_1"}},
				{Order: &pathway.Order{OrderID: "ORDER_ID_2"}},
				{Delay: &pathway.Delay{From: delay, To: delay}},
				{Result: &pathway.Results{
					OrderID: "ORDER_ID_2",
					Results: []*pathway.Result{
						{TestName: "TEST_NAME_2"},
					},
				}},
				{Delay: &pathway.Delay{From: delay, To: delay}},
				{Result: &pathway.Results{ // Will be overwritten.
					OrderID: "ORDER_ID_1",
					Results: []*pathway.Result{
						{TestName: "TEST_NAME_1_1"},
					},
				}},
				{Result: &pathway.Results{
					OrderID: "ORDER_ID_1",
					Results: []*pathway.Result{
						{TestName: "TEST_NAME_1_2"},
						{TestName: "TEST_NAME_1_3"},
					},
				}},
				{Discharge: &pathway.Discharge{}},
				{GenerateResources: &pathway.GenerateResources{}},
			},
		},
		want: []*ir.Encounter{{
			Status:      constants.EncounterStatusFinished,
			Start:       ir.NewValidTime(now),
			End:         ir.NewValidTime(evenLater),
			StatusStart: ir.NewValidTime(evenLater),
			LocationHistory: []*ir.LocationHistory{{
				Location: wardBed(1),
				Start:    ir.NewValidTime(now),
				End:      ir.NewValidTime(evenLater),
			}},
			StatusHistory: []*ir.StatusHistory{{
				Status: constants.EncounterStatusArrived,
				Start:  ir.NewValidTime(now),
				End:    ir.NewValidTime(now),
			}, {
				Status: constants.EncounterStatusInProgress,
				Start:  ir.NewValidTime(now),
				End:    ir.NewValidTime(evenLater),
			}},
			Orders: []*ir.Order{{
				OrderDateTime:    ir.NewValidTime(now),
				ReportedDateTime: ir.NewValidTime(evenLater),
				Results: []*ir.Result{
					{TestName: &ir.CodedElement{ID: "TEST_NAME_1_2", Text: "TEST_NAME_1_2"}},
					{TestName: &ir.CodedElement{ID: "TEST_NAME_1_3", Text: "TEST_NAME_1_3"}},
				},
			}, {
				OrderDateTime:    ir.NewValidTime(now),
				ReportedDateTime: ir.NewValidTime(later),
				Results: []*ir.Result{
					{TestName: &ir.CodedElement{ID: "TEST_NAME_2", Text: "TEST_NAME_2"}},
				},
			}},
		}},
	}, {
		name: "Update Orders after ending Encounter",
		pathway: pathway.Pathway{
			Pathway: []pathway.Step{
				{Order: &pathway.Order{OrderID: "ORDER_ID"}}, // Will create and finish a new Encounter.
				{Delay: &pathway.Delay{From: delay, To: delay}},
				{Result: &pathway.Results{
					OrderID: "ORDER_ID",
					Results: []*pathway.Result{{TestName: "TEST_NAME"}},
				}},
				{GenerateResources: &pathway.GenerateResources{}},
			}},
		want: []*ir.Encounter{{
			Status:      constants.EncounterStatusFinished,
			StatusStart: ir.NewValidTime(now),
			Start:       ir.NewValidTime(now),
			End:         ir.NewValidTime(now),
			StatusHistory: []*ir.StatusHistory{{
				Status: constants.EncounterStatusInProgress,
				Start:  ir.NewValidTime(now),
				End:    ir.NewValidTime(now),
			}},
			LocationHistory: []*ir.LocationHistory{{
				Location: hospitalLoc(),
				Start:    ir.NewValidTime(now),
				End:      ir.NewValidTime(now),
			}},
			Orders: []*ir.Order{{
				OrderDateTime:    ir.NewValidTime(now),
				ReportedDateTime: ir.NewValidTime(later),
				Results: []*ir.Result{
					{TestName: &ir.CodedElement{ID: "TEST_NAME", Text: "TEST_NAME"}},
				},
			}},
		}},
	}, {
		name: "Transfer twice",
		pathway: pathway.Pathway{
			Pathway: []pathway.Step{
				{Admission: &pathway.Admission{Loc: testLoc}},
				{Transfer: &pathway.Transfer{Loc: testLocAE}},
				{Delay: &pathway.Delay{From: delay, To: delay}},
				{Transfer: &pathway.Transfer{Loc: testLoc}},
				{Delay: &pathway.Delay{From: delay, To: delay}},
				{Discharge: &pathway.Discharge{}},
				{GenerateResources: &pathway.GenerateResources{}},
			},
		},
		want: []*ir.Encounter{{
			Status:      constants.EncounterStatusFinished,
			StatusStart: ir.NewValidTime(evenLater),
			Start:       ir.NewValidTime(now),
			End:         ir.NewValidTime(evenLater),
			StatusHistory: []*ir.StatusHistory{{
				Status: constants.EncounterStatusArrived,
				Start:  ir.NewValidTime(now),
				End:    ir.NewValidTime(evenLater),
			}},
			LocationHistory: []*ir.LocationHistory{{
				Location: wardBed(1),
				Start:    ir.NewValidTime(now),
				End:      ir.NewValidTime(now),
			}, {
				Location: aAndEBed(1),
				Start:    ir.NewValidTime(now),
				End:      ir.NewValidTime(later),
			}, {
				Location: wardBed(1),
				Start:    ir.NewValidTime(later),
				End:      ir.NewValidTime(evenLater),
			}},
		}},
	}, {
		name: "Pending transfer",
		pathway: pathway.Pathway{
			Pathway: []pathway.Step{
				{PendingTransfer: &pathway.PendingTransfer{Loc: testLoc, ExpectedTransferTimeFromNow: &delay}},
				{Transfer: &pathway.Transfer{Loc: testLoc}},
				{GenerateResources: &pathway.GenerateResources{}},
			},
		},
		want: []*ir.Encounter{{
			Status:      constants.EncounterStatusArrived,
			StatusStart: ir.NewValidTime(later),
			Start:       ir.NewValidTime(later),
			End:         ir.NewInvalidTime(),
			StatusHistory: []*ir.StatusHistory{{
				Status: constants.EncounterStatusPlanned,
				Start:  ir.NewValidTime(later),
				End:    ir.NewValidTime(later),
			}},
			LocationHistory: []*ir.LocationHistory{{
				Location: wardBed(1),
				Start:    ir.NewValidTime(later),
				End:      ir.NewInvalidTime(),
			}},
		}},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rw := testresource.NewWriter()

			pathways := map[string]pathway.Pathway{
				testPathwayName: tc.pathway,
			}

			hospital := hospitalWithTime(ctx, t, Config{ResourceWriter: rw}, pathways, now)
			defer hospital.Close()

			if err := hospital.StartNextPathway(); err != nil {
				t.Fatalf("StartNextPathway() failed with %v", err)
			}

			hospital.ConsumeQueues(ctx, t)

			if got, want := len(rw.Resources), 1; got != want {
				t.Fatalf("len(rw.Resources) = %d, want %d", got, want)
			}

			p := rw.Resources[0]

			if got, want := len(p.Encounters), len(tc.want); got != want {
				t.Fatalf("len(p.Encounters) = %d, want %d", got, want)
			}

			opts := []cmp.Option{
				cmpopts.EquateEmpty(),
				// We ignore any fields that are randomly generated.
				cmpopts.IgnoreFields(ir.Result{}, "Status", "ObservationDateTime", "Notes", "ValueType"),
				cmpopts.IgnoreFields(ir.Order{}, "NumberOfPreviousResults", "MessageControlIDOriginalOrder",
					"OrderProfile", "Placer", "Filler", "OrderControl", "OrderStatus", "ResultsStatus",
					"ReceivedInLabDateTime", "CollectedDateTime"),
			}

			var got []*ir.Encounter
			for _, e := range p.Encounters {
				got = append(got, e)
			}
			if diff := cmp.Diff(tc.want, got, opts...); diff != "" {
				t.Errorf("StartNextPathway() encounters returned diff (-want +got):\n%s", diff)
			}
		})
	}
}

func wardBed(i int) *ir.PatientLocation {
	l := aAndEBed(i)
	l.Poc = "Ward 1"
	return l
}

func aAndEBed(i int) *ir.PatientLocation {
	l := hospitalLoc()
	l.Room = "Room-1"
	l.Floor = "7"
	l.Bed = fmt.Sprintf("Bed %d", i)
	return l
}

func hospitalLoc() *ir.PatientLocation {
	return &ir.PatientLocation{
		Poc:          "ED",
		Facility:     "Simulated Hospital",
		LocationType: "BED",
		Building:     "Building-1",
	}
}
