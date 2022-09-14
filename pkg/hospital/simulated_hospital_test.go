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
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/testing/protocmp"
	"github.com/google/simhospital/pkg/generator/header"
	"github.com/google/simhospital/pkg/hardcoded"
	"github.com/google/simhospital/pkg/hl7"
	. "github.com/google/simhospital/pkg/hospital"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/logging"
	"github.com/google/simhospital/pkg/message"
	"github.com/google/simhospital/pkg/pathway"
	"github.com/google/simhospital/pkg/processor"
	"github.com/google/simhospital/pkg/state/persist"
	"github.com/google/simhospital/pkg/state"
	"github.com/google/simhospital/pkg/test"
	"github.com/google/simhospital/pkg/test/testhl7"
	"github.com/google/simhospital/pkg/test/testhospital"
	"github.com/google/simhospital/pkg/test/testlocation"
	"github.com/google/simhospital/pkg/test/testmetrics"
	"github.com/google/simhospital/pkg/test/teststate"
	"github.com/google/simhospital/pkg/test/testwrite"
)

var (
	// now is an arbitrary time.
	now       = time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC)
	delay     = time.Hour * 5
	later     = now.Add(delay)
	evenLater = later.Add(delay)
	oneDayAgo = -24 * time.Hour
)

const (
	testPathwayName     = "test_pathway"
	testLoc             = "Ward 1"
	testLocAE           = "ED"
	hardcodedMessageYml = `DischargeHardcodedMessage:
  segments:
    - "MSH|^~\\&|sending_application_reliable|sending_facility|receiving_application|receiving_facility|%s||ADT^A03|%s|T|2.3|||AL||44|ASCII"
    - "EVN|A03|20180212000000|||id-1^surname-1^firstname-1^^^prefix-1^^^DRNBR^PRSNL^^^ORGDR|"
    - "PID_SEGMENT_PLACEHOLDER"
    - "PD1|||Test Primary Facility^^123|"
    - "PV1|1|PREADMIT||28b|||id-1^surname-1^firstname-1^^^prefix-1^^^DRNBR^PRSNL^^^ORGDR|||specialty-1||||||||PREADMIT|3448412528941593955^^^^visitid||||||||||||||||||||||FINISHED|||20180212000000|20180212000000|"`
)

func TestMain(m *testing.M) {
	logging.SetLogLevel(logrus.DebugLevel)
	hl7.TimezoneAndLocation("Europe/London")

	retCode := m.Run()

	os.Exit(retCode)
}

func TestStartPathwayWithDelays(t *testing.T) {
	ctx := context.Background()
	type duration struct {
		from time.Duration
		to   time.Duration
	}
	oneMinute := time.Minute
	fiveMinutes := 5 * time.Minute
	oneDayAgo := -24 * time.Hour
	oneHourAgo := -time.Hour
	twoHours := 2 * time.Hour
	fiveHours := 5 * time.Hour

	tests := []struct {
		name              string
		pathway           pathway.Pathway
		wantMessageDelays map[string]duration
		wantEventDelays   map[string]duration
		wantMessageTypes  map[string]bool
	}{{
		name: "No Delay",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{Discharge: &pathway.Discharge{}},
		}},
		wantMessageDelays: map[string]duration{"ADT^A01": {from: time.Duration(0), to: time.Duration(0)}},
		wantMessageTypes:  map[string]bool{"ADT^A01": true, "ADT^A03": true},
	}, {
		name: "With Delay step",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{Delay: &pathway.Delay{From: oneMinute, To: fiveMinutes}},
			{Discharge: &pathway.Discharge{}},
		}},
		wantMessageDelays: map[string]duration{"ADT^A03": {from: oneMinute, to: fiveMinutes}},
		wantEventDelays:   map[string]duration{"ADT^A03": {from: oneMinute, to: fiveMinutes}},
		wantMessageTypes:  map[string]bool{"ADT^A01": true, "ADT^A03": true},
	}, {
		name: "With message delay",
		pathway: pathway.Pathway{Pathway: []pathway.Step{{
			Admission: &pathway.Admission{Loc: testLoc},
			Parameters: &pathway.Parameters{
				DelayMessage: &pathway.Delay{From: oneMinute, To: fiveMinutes},
			},
		}, {
			Discharge: &pathway.Discharge{},
		}}},
		wantMessageDelays: map[string]duration{"ADT^A01": {from: oneMinute, to: fiveMinutes}},
		wantEventDelays:   map[string]duration{"ADT^A01": {from: time.Duration(0), to: time.Duration(0)}},
		wantMessageTypes:  map[string]bool{"ADT^A01": true, "ADT^A03": true},
	}, {
		name: "Event in the past",
		pathway: pathway.Pathway{Pathway: []pathway.Step{{
			Discharge: &pathway.Discharge{},
		}, {
			Admission: &pathway.Admission{Loc: testLoc},
			Parameters: &pathway.Parameters{
				TimeFromNow: &oneDayAgo,
			},
		}}},
		wantMessageDelays: map[string]duration{"ADT^A01": {from: oneDayAgo, to: oneDayAgo}},
		wantMessageTypes:  map[string]bool{"ADT^A01": true, "ADT^A03": true},
	}, {
		name: "Historical step",
		pathway: pathway.Pathway{
			History: []pathway.Step{{
				Discharge: &pathway.Discharge{},
				Parameters: &pathway.Parameters{
					TimeFromNow: &oneDayAgo,
				}},
			},
			Pathway: []pathway.Step{
				{Admission: &pathway.Admission{Loc: testLoc}},
			},
		},
		wantMessageDelays: map[string]duration{
			"ADT^A03": {from: oneDayAgo, to: oneDayAgo},
			"ADT^A01": {from: time.Duration(0), to: time.Duration(0)},
		},
		wantMessageTypes: map[string]bool{"ADT^A01": true, "ADT^A03": true},
	}, {
		name: "Event in the past with short message delay",
		pathway: pathway.Pathway{
			Pathway: []pathway.Step{{Discharge: &pathway.Discharge{}}, {
				Admission: &pathway.Admission{Loc: testLoc},
				Parameters: &pathway.Parameters{
					TimeFromNow: &oneDayAgo,
					DelayMessage: &pathway.Delay{
						From: oneMinute,
						To:   fiveMinutes,
					},
				},
			}},
		},
		wantMessageDelays: map[string]duration{
			"ADT^A03": {from: time.Duration(0), to: time.Duration(0)},
			"ADT^A01": {from: oneDayAgo + oneMinute, to: oneDayAgo + fiveMinutes},
		},
		wantEventDelays: map[string]duration{
			"ADT^A03": {from: time.Duration(0), to: time.Duration(0)},
			"ADT^A01": {from: oneDayAgo, to: oneDayAgo},
		},
		wantMessageTypes: map[string]bool{"ADT^A01": true, "ADT^A03": true},
	}, {
		name: "Event in the past with long message delay",
		pathway: pathway.Pathway{
			Pathway: []pathway.Step{{Discharge: &pathway.Discharge{}}, {
				Admission: &pathway.Admission{Loc: testLoc},
				Parameters: &pathway.Parameters{
					TimeFromNow: &oneHourAgo,
					DelayMessage: &pathway.Delay{
						From: twoHours,
						To:   fiveHours,
					},
				}},
			},
		},
		wantMessageTypes: map[string]bool{"ADT^A01": true, "ADT^A03": true},
		wantMessageDelays: map[string]duration{
			"ADT^A03": {from: time.Duration(0), to: time.Duration(0)},
			"ADT^A01": {from: time.Hour, to: 4 * time.Hour},
		},
		wantEventDelays: map[string]duration{
			"ADT^A03": {from: time.Duration(0), to: time.Duration(0)},
			"ADT^A01": {from: oneHourAgo, to: oneHourAgo},
		},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			now := time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC)

			pathways := map[string]pathway.Pathway{
				testPathwayName: tc.pathway,
			}
			hospital := hospitalWithTime(ctx, t, Config{}, pathways, now)
			defer hospital.Close()

			startPathway(t, hospital, testPathwayName)
			_, messages := hospital.ConsumeQueues(ctx, t)
			if got, want := len(messages), len(tc.wantMessageTypes); got != want {
				t.Errorf("StartPathway(%v) generated %v messages, want %v", testPathwayName, got, want)
			}

			gotMessageTypes := map[string]bool{}
			gotMessages := map[string]string{}
			for _, m := range messages {
				mt := testhl7.MessageType(t, m)
				gotMessageTypes[mt] = true
				gotMessages[mt] = m
			}

			if diff := cmp.Diff(tc.wantMessageTypes, gotMessageTypes); diff != "" {
				t.Errorf("StartPathway(%v) generated message types with diff (-want, +got):\n%s", testPathwayName, diff)
			}

			for mt, wantDelay := range tc.wantMessageDelays {
				delay := time.Duration(testhl7.MessageDateTime(t, gotMessages[mt]).UnixNano() - now.UnixNano())
				if delay < wantDelay.from || delay > wantDelay.to {
					t.Errorf("Messages[%s] delay got %v, want it in the range (%v, %v)", mt, delay, wantDelay.from, wantDelay.to)
				}
			}
			for mt, wantDelay := range tc.wantEventDelays {
				delay := time.Duration(testhl7.EventDateTime(t, gotMessages[mt]).UnixNano() - now.UnixNano())
				if delay < wantDelay.from || delay > wantDelay.to {
					t.Errorf("Events[%s] delay got %v, want it in the range (%v, %v)", mt, delay, wantDelay.from, wantDelay.to)
				}
			}
		})
	}
}

func TestStartPathway_OrderAckAndResult(t *testing.T) {
	ctx := context.Background()
	results := []*pathway.Result{{
		TestName: "Creatinine",
		Value:    "52",
		Unit:     "UMOLL",
	}}
	tests := []struct {
		name              string
		steps             []pathway.Step
		orderAckDelay     time.Duration
		wantMessageTypes  []string
		wantPlacerNumbers []string
		wantOrderStatuses []string
	}{{
		name: "Same Test, Ack delay",
		steps: []pathway.Step{{
			Order: &pathway.Order{
				OrderID:      "order1",
				OrderProfile: "UREA AND ELECTROLYTES",
			},
		}, {
			Result: &pathway.Results{
				OrderID:      "order1",
				OrderProfile: "UREA AND ELECTROLYTES",
				Results:      results,
			}},
		},
		orderAckDelay:     time.Second,
		wantMessageTypes:  []string{"ORM^O01", "ORU^R01", "ORR^O02"},
		wantPlacerNumbers: []string{"1", "1", "1"},
		wantOrderStatuses: []string{"IP", "CM", "IP"},
	}, {
		name: "Same Test, no Ack delay",
		steps: []pathway.Step{{
			Order: &pathway.Order{
				OrderID:      "order1",
				OrderProfile: "UREA AND ELECTROLYTES",
			},
		}, {
			Result: &pathway.Results{
				OrderID:      "order1",
				OrderProfile: "UREA AND ELECTROLYTES",
				Results:      results,
			}},
		},
		wantMessageTypes:  []string{"ORM^O01", "ORR^O02", "ORU^R01"},
		wantPlacerNumbers: []string{"1", "1", "1"},
		wantOrderStatuses: []string{"IP", "IP", "CM"},
	}, {
		name: "Different Test, Ack delay",
		steps: []pathway.Step{{
			Order: &pathway.Order{
				OrderProfile: "X-Ray",
			},
		}, {
			Result: &pathway.Results{
				OrderProfile: "UREA AND ELECTROLYTES",
				Results:      results,
			}},
		},
		orderAckDelay:     time.Second,
		wantMessageTypes:  []string{"ORM^O01", "ORU^R01", "ORR^O02"},
		wantPlacerNumbers: []string{"1", "2", "1"},
		wantOrderStatuses: []string{"IP", "CM", "IP"},
	}, {
		name: "Different Test, no Ack delay",
		steps: []pathway.Step{{
			Order: &pathway.Order{
				OrderProfile: "X-Ray",
			},
		}, {
			Result: &pathway.Results{
				OrderProfile: "UREA AND ELECTROLYTES",
				Results:      results,
			}},
		},
		wantMessageTypes:  []string{"ORM^O01", "ORR^O02", "ORU^R01"},
		wantPlacerNumbers: []string{"1", "1", "2"},
		wantOrderStatuses: []string{"IP", "IP", "CM"},
	}, {
		name: "Same ID, no Order ACK",
		steps: []pathway.Step{{
			Order: &pathway.Order{
				OrderID:                  "order1",
				OrderProfile:             "UREA AND ELECTROLYTES",
				NoAcknowledgementMessage: true,
			},
		}, {
			Result: &pathway.Results{
				OrderID:      "order1",
				OrderProfile: "UREA AND ELECTROLYTES",
				Results: []*pathway.Result{{
					TestName: "Creatinine",
					Value:    "500",
					Unit:     "UMOLL",
				}},
			}},
		},
		wantMessageTypes:  []string{"ORM^O01", "ORU^R01"},
		wantPlacerNumbers: []string{"1", "1"},
		wantOrderStatuses: []string{"IP", "CM"},
	}, {
		name: "Order without ack, then another status update",
		steps: []pathway.Step{{
			Order: &pathway.Order{
				OrderID:                  "order1",
				OrderProfile:             "UREA AND ELECTROLYTE",
				NoAcknowledgementMessage: true,
			},
		}, {
			Order: &pathway.Order{
				OrderID:                  "order1",
				OrderProfile:             "UREA AND ELECTROLYTE",
				OrderStatus:              "IL",
				NoAcknowledgementMessage: true,
			},
		}},
		wantMessageTypes:  []string{"ORM^O01", "ORM^O01"},
		wantPlacerNumbers: []string{"1", "1"},
		wantOrderStatuses: []string{"IP", "IL"},
	}, {
		name: "Order with ack, then another status update, then a result",
		steps: []pathway.Step{{
			Order: &pathway.Order{
				OrderID:      "order1",
				OrderProfile: "UREA AND ELECTROLYTE",
			},
		}, {
			Order: &pathway.Order{
				OrderID:                  "order1",
				OrderProfile:             "UREA AND ELECTROLYTE",
				OrderStatus:              "IL",
				NoAcknowledgementMessage: true,
			},
		}, {
			Result: &pathway.Results{
				OrderID:      "order1",
				OrderProfile: "UREA AND ELECTROLYTES",
				Results: []*pathway.Result{{
					TestName: "Creatinine",
					Value:    "500",
					Unit:     "UMOLL",
				}},
			}},
		},
		wantMessageTypes:  []string{"ORM^O01", "ORR^O02", "ORM^O01", "ORU^R01"},
		wantPlacerNumbers: []string{"1", "1", "1", "1"},
		wantOrderStatuses: []string{"IP", "IP", "IL", "CM"},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pathways := map[string]pathway.Pathway{
				testPathwayName: {Pathway: tc.steps},
			}
			hospital := newHospital(ctx, t, Config{AdditionalConfig: AdditionalConfig{OrderAckDelay: &pathway.Delay{From: tc.orderAckDelay, To: tc.orderAckDelay}}}, pathways)

			startPathway(t, hospital, testPathwayName)
			_, messages := hospital.ConsumeQueues(ctx, t)
			if got, want := len(messages), len(tc.wantMessageTypes); got != want {
				t.Errorf("StartPathway(%v) generated %v messages, want %v", testPathwayName, got, want)
			}

			gotMessageTypes := testhl7.Fields(t, messages, testhl7.MessageType)
			if diff := cmp.Diff(tc.wantMessageTypes, gotMessageTypes); diff != "" {
				t.Errorf("StartPathway(%v) generated message types with diff (-want, +got):\n%s", testPathwayName, diff)
			}

			gotPlacerNumbers := testhl7.Fields(t, messages, testhl7.PlacerNumber)
			if diff := cmp.Diff(tc.wantPlacerNumbers, gotPlacerNumbers); diff != "" {
				t.Errorf("StartPathway(%v) generated message Placer numbers with diff (-want, +got):\n%s", testPathwayName, diff)
			}

			gotOrderStatuses := testhl7.Fields(t, messages, testhl7.OrderStatus)
			if diff := cmp.Diff(tc.wantOrderStatuses, gotOrderStatuses); diff != "" {
				t.Errorf("StartPathway(%v) generated ir.Order statuses with diff (-want, +got):\n%s", testPathwayName, diff)
			}
		})
	}
}

func TestStartPathway_OrderAckDelayIsRandom(t *testing.T) {
	ctx := context.Background()
	rand.Seed(1)
	pathways := map[string]pathway.Pathway{
		testPathwayName: {Pathway: []pathway.Step{{
			Order: &pathway.Order{
				OrderID:      "order1",
				OrderProfile: "UREA AND ELECTROLYTES",
			},
		}}},
	}

	// delays tracks how many times each delay, in seconds, happens.
	delays := map[float64]int{}
	minOrderDelay := time.Second
	maxOrderDelay := 10 * time.Second
	for i := minOrderDelay.Seconds(); i <= maxOrderDelay.Seconds(); i++ {
		delays[i] = 0
	}
	hospital := newHospital(ctx, t, Config{AdditionalConfig: AdditionalConfig{OrderAckDelay: &pathway.Delay{From: minOrderDelay, To: maxOrderDelay + time.Second}}}, pathways)
	defer hospital.Close()
	count := 1000
	wantMessageTypes := []string{"ORM^O01", "ORR^O02"}
	for i := 0; i < count; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			startPathway(t, hospital, testPathwayName)
			_, messages := hospital.ConsumeQueues(ctx, t)
			gotMessageTypes := testhl7.Fields(t, messages, testhl7.MessageType)
			if diff := cmp.Diff(wantMessageTypes, gotMessageTypes); diff != "" {
				t.Fatalf("StartPathway(%v) generated message types with diff (-want, +got):\n%s", testPathwayName, diff)
			}

			orderAck := messages[1]
			gotMessageControlIDs := testhl7.Fields(t, messages, testhl7.MessageControlIDFromMSH)
			wantOrderMsgCtrlID := strconv.Itoa(i*2 + 1)
			wantOrderAckMsgCtrlID := strconv.Itoa(i*2 + 2)
			if diff := cmp.Diff([]string{wantOrderMsgCtrlID, wantOrderAckMsgCtrlID}, gotMessageControlIDs); diff != "" {
				t.Errorf("StartPathway(%v) generated message control IDs with diff (-want, +got):\n%s", testPathwayName, diff)
			}
			referredMessageControlID := testhl7.MessageControlIDFromMSA(t, orderAck)
			if got, want := referredMessageControlID, wantOrderMsgCtrlID; got != want {
				t.Errorf("referredMessageControlID=%v, want %v", got, want)
			}

			orderDT := testhl7.MessageDateTime(t, messages[0])
			orderAckDT := testhl7.MessageDateTime(t, messages[1])
			delay := orderAckDT.Sub(orderDT).Seconds()
			if _, ok := delays[delay]; !ok {
				t.Errorf("OrderAcknowledgement delay got %v, want between %v and %v inclusive", delay, minOrderDelay, maxOrderDelay)
			}
			delays[delay]++
		})
	}
	want := count / len(delays)
	delta := float64(want / 2) // Allow some error.

	for k, v := range delays {
		if math.Abs(float64(v)-float64(want)) >= delta {
			t.Errorf("delays[%v] = %d, want within %.1f of %d", k, v, delta, want)
		}
	}
}

func TestRunPathwayWithMultipleHeaderFields(t *testing.T) {
	ctx := context.Background()
	sa1, sa2 := "sa-1", "sa-2"
	sf1, sf2 := "sf-1", "sf-2"
	pathways := map[string]pathway.Pathway{
		testPathwayName: {Pathway: []pathway.Step{
			{
				Admission: &pathway.Admission{Loc: testLoc},
				Parameters: &pathway.Parameters{
					SendingApplication: sa1,
					SendingFacility:    sf1,
				},
			},
			{
				Result: &pathway.Results{},
				Parameters: &pathway.Parameters{
					SendingApplication: sa2,
					SendingFacility:    sf2,
				},
			},
			// The Sending Facility isn't specified so it is the default one.
			{Result: &pathway.Results{}},
		}},
	}

	hospital := newHospital(ctx, t, Config{}, pathways)
	defer hospital.Close()
	startPathway(t, hospital, testPathwayName)
	_, messages := hospital.ConsumeQueues(ctx, t)
	if got, want := len(messages), 3; got != want {
		t.Fatalf("len(messages) = %d, want %d", got, want)
	}

	var gotSA []string
	var gotSF []string
	for _, m := range messages {
		gotSA = append(gotSA, testhl7.SendingApplication(t, m))
		gotSF = append(gotSF, testhl7.SendingFacility(t, m))
	}

	wantSA := []string{sa1, sa2, "sending_application"}
	if diff := cmp.Diff(wantSA, gotSA); diff != "" {
		t.Errorf("SendingApplication got diff (-want, +got)\n%s", diff)
	}

	wantSF := []string{sf1, sf2, "sending_facility"}
	if diff := cmp.Diff(wantSF, gotSF); diff != "" {
		t.Errorf("SendingFacility got diff (-want, +got)\n%s", diff)
	}
}

func TestRunPathway_StepTypes(t *testing.T) {
	ctx := context.Background()
	type metric struct {
		name     string
		labels   map[string]string
		wantDiff float64
	}

	yesterday := now.Add(oneDayAgo)
	oneDay := 24 * time.Hour
	twoHours := 2 * time.Hour
	// originalNumContentLines, appendNumContentLines and overwriteNumContentLines are arbitrarily chosen.
	originalNumContentLines := 11
	appendNumContentLines := 4
	overwriteNumContentLines := 26
	tests := []struct {
		name             string
		pathway          pathway.Pathway
		wantMessageTypes []string
		want             func(t *testing.T, messages []string, hospital *testhospital.Hospital)
		wantMetrics      []metric
	}{{
		name: "Admission with Discharge",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc, AdmitReason: "Eye problems"}},
			{Discharge: &pathway.Discharge{DischargeTime: &yesterday}},
		}},
		wantMessageTypes: []string{"ADT^A01", "ADT^A03"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			admitPV2 := testhl7.PV2(t, messages[0])
			if admitPV2.AdmitReason == nil {
				t.Fatal("admitPV2.AdmitReason=<nil>; want non nil")
			}
			if got, want := admitPV2.AdmitReason.Text.SanitizedString(), "Eye problems"; got != want {
				t.Errorf("admitPV2.AdmitReason.Text.SanitizedString()=%q, want %q", got, want)
			}
			dischargePV1 := testhl7.PV1(t, messages[1])
			if got, want := len(dischargePV1.DischargeDateTime), 1; got != want {
				t.Errorf("len(dischargePV1.DischargeDateTime)=%v, want %v", got, want)
			}
			if got, want := dischargePV1.DischargeDateTime[0].Time, yesterday; got != want {
				t.Errorf("dischargePV1.DischargeDateTime[0].Time=%v, want %v", got, want)
			}
			wantAccountStatus := []string{hospital.MessageConfig.PatientAccountStatus.Arrived, hospital.MessageConfig.PatientAccountStatus.Finished}
			gotAccountStatus := testhl7.Fields(t, messages, testhl7.AccountStatus)
			if diff := cmp.Diff(wantAccountStatus, gotAccountStatus); diff != "" {
				t.Errorf("StartPathway(%v) generated AccountStatus with diff (-want, +got):\n%s", testPathwayName, diff)
			}
			// The patient is discharged and thus the bed is free.
			if got, want := hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds(), 0; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds()=%v, want %v", got, want)
			}
		},
	}, {
		name: "Registration with Discharge",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Registration: &pathway.Registration{PatientClass: "PREADMIT"}},
			{Discharge: &pathway.Discharge{DischargeTime: &yesterday}},
		}},
		wantMessageTypes: []string{"ADT^A04", "ADT^A03"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			wantAccountStatus := []string{hospital.MessageConfig.PatientAccountStatus.Planned, hospital.MessageConfig.PatientAccountStatus.Finished}
			gotAccountStatus := testhl7.Fields(t, messages, testhl7.AccountStatus)
			if diff := cmp.Diff(wantAccountStatus, gotAccountStatus); diff != "" {
				t.Errorf("StartPathway(%v) generated AccountStatus with diff (-want, +got):\n%s", testPathwayName, diff)
			}
			wantPatientClass := []string{"PREADMIT", "PREADMIT"}
			gotPatientClass := testhl7.Fields(t, messages, testhl7.PatientClass)
			if diff := cmp.Diff(wantPatientClass, gotPatientClass); diff != "" {
				t.Errorf("StartPathway(%v) generated PatientClass with diff (-want, +got):\n%s", testPathwayName, diff)
			}
		},
	}, {
		name: "Admission with Document",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{Document: &pathway.Document{DocumentType: "DS"}},
		}},
		wantMessageTypes: []string{"ADT^A01", "MDM^T02"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			documentTXA := testhl7.TXA(t, messages[1])
			if got, want := documentTXA.DocumentType.String(), "DS"; got != want {
				t.Errorf("documentTXA.DocumentType.String() got %v, want %v", got, want)
			}
			if got, want := documentTXA.DocumentCompletionStatus.String(), "DO"; got != want {
				t.Errorf("documentTXA.DocumentCompletionStatus.String() got %v, want %v", got, want)
			}
			if got, want := len(documentTXA.UniqueDocumentNumber.EntityIdentifier.String()), 13; got != want {
				t.Errorf("len(documentTXA.UniqueDocumentNumber) got %v, want %v", got, want)
			}
			documentOBXs := testhl7.AllOBX(t, messages[1])
			if got, want := len(documentOBXs), 1; got < want {
				t.Errorf("len(documentOBXs) got %v, want greater than %v", got, want)
			}
		},
	}, {
		name: "Document with Document append update",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Document: &pathway.Document{ID: "id1", NumRandomContentLines: &pathway.Interval{From: originalNumContentLines, To: originalNumContentLines}}},
			{Document: &pathway.Document{ID: "id1", UpdateType: "append", NumRandomContentLines: &pathway.Interval{From: appendNumContentLines, To: appendNumContentLines}}},
		}},
		wantMessageTypes: []string{"MDM^T02", "MDM^T02"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			documentTXA := testhl7.TXA(t, messages[0])
			updateDocumentTXA := testhl7.TXA(t, messages[1])
			if diff := cmp.Diff(documentTXA, updateDocumentTXA); diff != "" {
				t.Errorf("Updated documentTXA got diff:\n%s", diff)
			}
			documentOBXs := testhl7.AllOBX(t, messages[0])
			updateDocumentOBXs := testhl7.AllOBX(t, messages[1])
			sumOfContentLines := originalNumContentLines + appendNumContentLines
			if got, want := len(updateDocumentOBXs), sumOfContentLines; got != want {
				t.Errorf("len(documentOBXs) got %v, want %v (%d old lines + %d new lines)", got, want, originalNumContentLines, appendNumContentLines)
			}
			if diff := cmp.Diff(documentOBXs[:4], updateDocumentOBXs[:4]); diff != "" {
				t.Errorf("Updated Document content got diff \n%s", diff)
			}
			if got, want := updateDocumentTXA.UniqueDocumentNumber.EntityIdentifier.String(), documentTXA.UniqueDocumentNumber.EntityIdentifier.String(); got != want {
				t.Errorf("updateDocumentTXA.UniqueDocumentNumber got %v, want %v", got, want)
			}
		},
	}, {
		name: "Document with Document overwrite update",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Document: &pathway.Document{ID: "id1", NumRandomContentLines: &pathway.Interval{From: originalNumContentLines, To: originalNumContentLines}}},
			{Document: &pathway.Document{ID: "id1", UpdateType: "overwrite", NumRandomContentLines: &pathway.Interval{From: overwriteNumContentLines, To: overwriteNumContentLines}}},
		}},
		wantMessageTypes: []string{"MDM^T02", "MDM^T02"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			documentTXA := testhl7.TXA(t, messages[0])
			updateDocumentTXA := testhl7.TXA(t, messages[1])

			if diff := cmp.Diff(documentTXA, updateDocumentTXA); diff != "" {
				t.Errorf("Updated documentTXA got diff:\n%s", diff)
			}
			updateDocumentOBXs := testhl7.AllOBX(t, messages[1])
			if got, want := len(updateDocumentOBXs), overwriteNumContentLines; got != want {
				t.Errorf("Updated len(documentOBXs) got %v, want %v", got, want)
			}
		},
	}, {
		name: "Document with multiple updates",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Document: &pathway.Document{ID: "id1", NumRandomContentLines: &pathway.Interval{From: originalNumContentLines, To: originalNumContentLines}}},
			{Document: &pathway.Document{ID: "id1", UpdateType: "overwrite", NumRandomContentLines: &pathway.Interval{From: overwriteNumContentLines, To: overwriteNumContentLines}}},
			{Document: &pathway.Document{ID: "id1", UpdateType: "append", NumRandomContentLines: &pathway.Interval{From: appendNumContentLines, To: appendNumContentLines}}},
		}},
		wantMessageTypes: []string{"MDM^T02", "MDM^T02", "MDM^T02"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			documentTXA := testhl7.TXA(t, messages[0])
			finalUpdatedDocumentTXA := testhl7.TXA(t, messages[2])

			if diff := cmp.Diff(documentTXA, finalUpdatedDocumentTXA); diff != "" {
				t.Errorf("Updated documentTXA got diff:\n%s", diff)
			}
			finalUpdatedDocumentOBXs := testhl7.AllOBX(t, messages[2])
			finalSumOfContentLines := overwriteNumContentLines + appendNumContentLines
			if got, want := len(finalUpdatedDocumentOBXs), finalSumOfContentLines; got != want {
				t.Errorf("Updated len(documentOBXs) got %v, want %v", got, want)
			}
		},
	}, {
		name: "Document with existing Document ID",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Document: &pathway.Document{ID: "id1"}},
			{Document: &pathway.Document{ID: "id1"}},
		}},
		wantMessageTypes: []string{"MDM^T02"},
		wantMetrics: []metric{{
			name: "simulated_hospital_errors_total",
			labels: map[string]string{
				"pathway_name": testPathwayName,
				"reason":       "Document.ID already exists",
			},
			wantDiff: 1,
		}},
	}, {
		name: "Document update with non-existing ID",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Document: &pathway.Document{ID: "id1", UpdateType: "overwrite"}},
		}},
		wantMessageTypes: nil,
		wantMetrics: []metric{{
			name: "simulated_hospital_errors_total",
			labels: map[string]string{
				"pathway_name": testPathwayName,
				"reason":       "Document.ID does not exist",
			},
			wantDiff: 1,
		}},
	}, {
		name: "Document update with unsupported update type",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Document: &pathway.Document{ID: "id1"}},
			{Document: &pathway.Document{ID: "id1", UpdateType: "delete"}},
		}},
		wantMessageTypes: []string{"MDM^T02"},
		wantMetrics: []metric{{
			name: "simulated_hospital_errors_total",
			labels: map[string]string{
				"pathway_name": testPathwayName,
				"reason":       `cannot update document: update type was "delete", expected "append" or "overwrite"`,
			},
			wantDiff: 1,
		}},
	}, {
		name: "PreAdmission",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{PreAdmission: &pathway.PreAdmission{Loc: testLoc, ExpectedAdmissionTimeFromNow: &oneDay}},
			{Admission: &pathway.Admission{Loc: testLoc}},
		}},
		wantMessageTypes: []string{"ADT^A05", "ADT^A01"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			wantAccountStatus := []string{hospital.MessageConfig.PatientAccountStatus.Planned, hospital.MessageConfig.PatientAccountStatus.Arrived}
			gotAccountStatus := testhl7.Fields(t, messages, testhl7.AccountStatus)
			if diff := cmp.Diff(wantAccountStatus, gotAccountStatus); diff != "" {
				t.Errorf("StartPathway(%v) generated AccountStatus with diff (-want, +got):\n%s", testPathwayName, diff)
			}
			preAdmission, admission := messages[0], messages[1]
			preAdmissionPV1 := testhl7.PV1(t, messages[0])
			preAdmissionPV2 := testhl7.PV2(t, preAdmission)
			preAdmissionEVN := testhl7.EVN(t, preAdmission)
			admissionPV1 := testhl7.PV1(t, admission)

			if got, want := preAdmissionPV1.PendingLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLoc].Poc; got != want {
				t.Errorf("preAdmissionPV1.PendingLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if admissionPV1.PendingLocation != nil {
				t.Errorf("admission.PendingLocation got %+v, want <nil>", admissionPV1.PendingLocation)
			}
			if got, want := *admissionPV1.AdmitDateTime, *preAdmissionPV2.ExpectedAdmitDateTime; got != want {
				t.Errorf("*admission.AdmitDateTime=%v, want %v", got, want)
			}
			if got, want := *preAdmissionEVN.DateTimePlannedEvent, *preAdmissionPV2.ExpectedAdmitDateTime; got != want {
				t.Errorf("*preAdmissionEVN.DateTimePlannedEvent=%v, want %v", got, want)
			}
		},
	}, {
		name: "Transfer and Discharge",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{Transfer: &pathway.Transfer{Loc: testLoc}},
			{Result: &pathway.Results{}},
			{Discharge: &pathway.Discharge{}},
		}},
		wantMessageTypes: []string{"ADT^A01", "ADT^A02", "ORU^R01", "ADT^A03"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			transfer, result := messages[1], messages[2]

			transferPV1 := testhl7.PV1(t, transfer)
			if got, want := transferPV1.PriorPatientLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLocAE].Poc; got != want {
				t.Errorf("transferPV1.PriorPatientLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if got, want := transferPV1.AssignedPatientLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLoc].Poc; got != want {
				t.Errorf("transferPV1.AssignedPatientLocation.PointOfCare.String()=%v, want %v", got, want)
			}

			// The PriorLocation should only be set in the Transfer message.
			resultPV1 := testhl7.PV1(t, result)
			if resultPV1.PriorPatientLocation != nil {
				t.Errorf("resultPV1.PriorPatientLocation got %+v, want <nil>", resultPV1.PriorPatientLocation)
			}

			// The patient is discharged and thus the bed is free.
			if got, want := hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds(), 0; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds()=%v, want %v", got, want)
			}
		},
	}, {
		name: "Cancel Visit after Admission",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{CancelVisit: &pathway.CancelVisit{}},
		}},
		wantMessageTypes: []string{"ADT^A01", "ADT^A11"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			wantAccountStatus := []string{hospital.MessageConfig.PatientAccountStatus.Arrived, hospital.MessageConfig.PatientAccountStatus.Cancelled}
			gotAccountStatus := testhl7.Fields(t, messages, testhl7.AccountStatus)
			if diff := cmp.Diff(wantAccountStatus, gotAccountStatus); diff != "" {
				t.Errorf("StartPathway(%v) generated AccountStatus with diff (-want, +got):\n%s", testPathwayName, diff)
			}
			if got, want := hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds(), 0; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds()=%v, want %v", got, want)
			}
			admission, cancelVisit := messages[0], messages[1]
			admissionVisitNumber := testhl7.VisitNumber(t, admission)
			wantVisitNumber := []string{admissionVisitNumber, admissionVisitNumber}
			gotVisitNumber := testhl7.Fields(t, messages, testhl7.VisitNumber)
			if diff := cmp.Diff(wantVisitNumber, gotVisitNumber); diff != "" {
				t.Errorf("StartPathway(%v) generated VisitNumber with diff (-want, +got):\n%s", testPathwayName, diff)
			}

			admissionEVN := testhl7.EVN(t, admission)
			cancelVisitEVN := testhl7.EVN(t, cancelVisit)
			if diff := cmp.Diff(admissionEVN.RecordedDateTime, cancelVisitEVN.EventOccurred); diff != "" {
				t.Errorf("cancelVisitEVN.EventOccurred got diff (-want, +got):\n%s", diff)
			}
		},
	}, {
		name: "Cancel Visit after Registration",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Registration: &pathway.Registration{}},
			{CancelVisit: &pathway.CancelVisit{}},
		}},
		wantMessageTypes: []string{"ADT^A04", "ADT^A11"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			wantAccountStatus := []string{hospital.MessageConfig.PatientAccountStatus.Planned, hospital.MessageConfig.PatientAccountStatus.Cancelled}
			gotAccountStatus := testhl7.Fields(t, messages, testhl7.AccountStatus)
			if diff := cmp.Diff(wantAccountStatus, gotAccountStatus); diff != "" {
				t.Errorf("StartPathway(%v) generated AccountStatus with diff (-want, +got):\n%s", testPathwayName, diff)
			}
			if got, want := hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds(), 0; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds()=%v, want %v", got, want)
			}
			registration, cancelVisit := messages[0], messages[1]
			registrationVisitNumber := testhl7.VisitNumber(t, registration)
			wantVisitNumber := []string{registrationVisitNumber, registrationVisitNumber}
			gotVisitNumber := testhl7.Fields(t, messages, testhl7.VisitNumber)
			if diff := cmp.Diff(wantVisitNumber, gotVisitNumber); diff != "" {
				t.Errorf("StartPathway(%v) generated VisitNumber with diff (-want, +got):\n%s", testPathwayName, diff)
			}

			registrationEVN := testhl7.EVN(t, registration)
			cancelVisitEVN := testhl7.EVN(t, cancelVisit)
			if diff := cmp.Diff(registrationEVN.RecordedDateTime, cancelVisitEVN.EventOccurred); diff != "" {
				t.Errorf("cancelVisitEVN.EventOccurred got diff (-want, +got):\n%s", diff)
			}
		},
	}, {
		name: "Result after CancelVisit",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Result: &pathway.Results{}},
			{Admission: &pathway.Admission{Loc: testLoc}},
			{CancelVisit: &pathway.CancelVisit{}},
			{Result: &pathway.Results{}},
		}},
		wantMessageTypes: []string{"ORU^R01", "ADT^A01", "ADT^A11", "ORU^R01"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			wantPatientClass := []string{
				hospital.MessageConfig.PatientClass.Outpatient,
				hospital.MessageConfig.PatientClass.Inpatient,
				hospital.MessageConfig.PatientClass.Inpatient,
				hospital.MessageConfig.PatientClass.Outpatient,
			}
			gotPatientClass := testhl7.Fields(t, messages, testhl7.PatientClass)
			if diff := cmp.Diff(wantPatientClass, gotPatientClass); diff != "" {
				t.Errorf("StartPathway(%v) generated PatientClass with diff (-want, +got):\n%s", testPathwayName, diff)
			}
			result := messages[3]
			if got := testhl7.PV1(t, result).VisitNumber; got != nil {
				t.Errorf("pv1.VisitNumber for Result after CancelVisit got %+v, want <nil>", got)
			}
		},
	}, {
		name: "CancelTransfer",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Result: &pathway.Results{}},
			{Admission: &pathway.Admission{Loc: testLoc}},
			{TransferInError: &pathway.TransferInError{Loc: testLocAE}},
			{Result: &pathway.Results{}},
			{CancelTransfer: &pathway.CancelTransfer{}},
		}},
		wantMessageTypes: []string{"ORU^R01", "ADT^A01", "ADT^A02", "ORU^R01", "ADT^A12"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			result1, transfer, result2, cancelTransfer := messages[0], messages[2], messages[3], messages[4]

			if got, want := hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds(), 1; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds()=%v, want %v", got, want)
			}
			if got, want := hospital.LocationManager.RoomManagers[testLocAE].OccupiedBeds(), 0; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLocAE].OccupiedBeds()=%v, want %v", got, want)
			}

			transferEVN := testhl7.EVN(t, transfer)
			cancelTransferEVN := testhl7.EVN(t, cancelTransfer)
			if diff := cmp.Diff(transferEVN.RecordedDateTime, cancelTransferEVN.EventOccurred); diff != "" {
				t.Errorf("transferEVN.EventOccurred got diff (-want, +got):\n%s", diff)
			}

			cancelTransferPV1 := testhl7.PV1(t, cancelTransfer)
			if got, want := cancelTransferPV1.PriorPatientLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLocAE].Poc; got != want {
				t.Errorf("cancelTransferPV1.PriorPatientLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if got, want := cancelTransferPV1.AssignedPatientLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLoc].Poc; got != want {
				t.Errorf("cancelTransferPV1.AssignedPatientLocation.PointOfCare.String()=%v, want %v", got, want)
			}

			transferPV1 := testhl7.PV1(t, transfer)
			if diff := cmp.Diff(transferPV1.AssignedPatientLocation, cancelTransferPV1.PriorPatientLocation); diff != "" {
				t.Errorf("cancelTransferPV1.PriorPatientLocation got diff (-want, +got):\n%s", diff)
			}
			if diff := cmp.Diff(transferPV1.PriorPatientLocation, cancelTransferPV1.AssignedPatientLocation); diff != "" {
				t.Errorf("cancelTransferPV1.AssignedPatientLocation diff (-want, +got):\n%s", diff)
			}

			// The PriorLocation should only be set in the TransferInError and CancelTransfer messages.
			result1PV1 := testhl7.PV1(t, result1)
			if result1PV1.PriorPatientLocation != nil {
				t.Errorf("result1PV1.PriorPatientLocation = %+v, want <nil>", result1PV1.PriorPatientLocation)
			}
			result2PV1 := testhl7.PV1(t, result2)
			if result2PV1.PriorPatientLocation != nil {
				t.Errorf("result2PV1.PriorPatientLocation = %+v, want <nil>", result2PV1.PriorPatientLocation)
			}
		},
	}, {
		name: "CancelDischarge",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Result: &pathway.Results{}},
			{Admission: &pathway.Admission{Loc: testLoc}},
			{DischargeInError: &pathway.DischargeInError{}},
			{CancelDischarge: &pathway.CancelDischarge{}},
		}},
		wantMessageTypes: []string{"ORU^R01", "ADT^A01", "ADT^A03", "ADT^A13"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			wantAccountStatus := []string{
				"",
				hospital.MessageConfig.PatientAccountStatus.Arrived,
				hospital.MessageConfig.PatientAccountStatus.Finished,
				hospital.MessageConfig.PatientAccountStatus.Arrived,
			}
			gotAccountStatus := testhl7.Fields(t, messages, testhl7.AccountStatus)
			if diff := cmp.Diff(wantAccountStatus, gotAccountStatus); diff != "" {
				t.Errorf("StartPathway(%v) generated AccountStatus with diff (-want, +got):\n%s", testPathwayName, diff)
			}

			if got, want := hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds(), 1; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds()=%v, want %v", got, want)
			}

			discharge, cancelDischarge := messages[2], messages[3]
			dischargeEVN := testhl7.EVN(t, discharge)
			cancelDischargeEVN := testhl7.EVN(t, cancelDischarge)
			if diff := cmp.Diff(dischargeEVN.RecordedDateTime, cancelDischargeEVN.EventOccurred); diff != "" {
				t.Errorf("cancelDischargeEVN.EventOccurred diff (-want, +got):\n%s", diff)
			}

			dischargePV1 := testhl7.PV1(t, discharge)
			if dischargePV1.PriorPatientLocation != nil {
				t.Errorf("dischargePV1.PriorPatientLocation = %+v, want <nil>", dischargePV1.PriorPatientLocation)
			}

			cancelDischargePV1 := testhl7.PV1(t, cancelDischarge)
			if got, want := cancelDischargePV1.AssignedPatientLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLoc].Poc; got != want {
				t.Errorf("cancelDischargePV1.AssignedPatientLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if cancelDischargePV1.PriorPatientLocation != nil {
				t.Errorf("cancelDischargePV1.PriorPatientLocation = %+v, want <nil>", cancelDischargePV1.PriorPatientLocation)
			}
		},
	}, {
		name: "CancelDischarge with DischargeTime",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Result: &pathway.Results{}},
			{Admission: &pathway.Admission{Loc: testLoc}},
			{DischargeInError: &pathway.DischargeInError{DischargeTime: &yesterday}},
			{CancelDischarge: &pathway.CancelDischarge{}},
		}},
		wantMessageTypes: []string{"ORU^R01", "ADT^A01", "ADT^A03", "ADT^A13"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			if got, want := hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds(), 1; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds()=%v, want %v", got, want)
			}
			discharge := messages[2]
			dischargePV1 := testhl7.PV1(t, discharge)
			if got, want := len(dischargePV1.DischargeDateTime), 1; got != want {
				t.Errorf("len(dischargePV1.DischargeDateTime)=%v, want %v", got, want)
			}
			if got, want := dischargePV1.DischargeDateTime[0].Time, yesterday; got != want {
				t.Errorf("dischargePV1.DischargeDateTime[0].Time=%v, want %v", got, want)
			}
		},
	}, {
		name: "CancelPendingTransfer",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{PendingTransfer: &pathway.PendingTransfer{Loc: testLoc, ExpectedTransferTimeFromNow: &twoHours}},
			{CancelPendingTransfer: &pathway.CancelPendingTransfer{}},
			{Result: &pathway.Results{}},
		}},
		wantMessageTypes: []string{"ADT^A01", "ADT^A15", "ADT^A26", "ORU^R01"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			pendingTransfer, cancelPendingTransfer, result := messages[1], messages[2], messages[3]
			cancelPendingTransferPV2 := testhl7.PV2(t, cancelPendingTransfer)
			if got, want := cancelPendingTransferPV2.PriorPendingLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLoc].Poc; got != want {
				t.Errorf("cancelPendingTransferPV2.PriorPendingLocation.PointOfCare.String()=%v, want %v", got, want)
			}

			pendingTransferPV1 := testhl7.PV1(t, pendingTransfer)
			pendingTransferEVN := testhl7.EVN(t, pendingTransfer)
			cancelPendingTransferPV1 := testhl7.PV1(t, cancelPendingTransfer)
			cancelPendingTransferEVN := testhl7.EVN(t, cancelPendingTransfer)
			if cancelPendingTransferPV1.PendingLocation != nil {
				t.Error("cancelPendingTransferPV1.PendingLocation is something, want <nil>.")
			}
			if diff := cmp.Diff(pendingTransferEVN.DateTimePlannedEvent, cancelPendingTransferEVN.EventOccurred); diff != "" {
				t.Errorf("cancelPendingTransferEVN.EventOccurred got diff (-want, +got):\n%s", diff)
			}
			if cancelPendingTransferPV1.VisitNumber == nil {
				t.Errorf("cancelPendingTransferPV1.VisitNumber = %+v, want <nil>", cancelPendingTransferPV1.VisitNumber)
			}
			if diff := cmp.Diff(pendingTransferPV1.VisitNumber, cancelPendingTransferPV1.VisitNumber); diff != "" {
				t.Errorf("cancelPendingTransferPV1.VisitNumber got diff (-want, +got):\n%s", diff)
			}

			// The pending location should have been cleared.
			admitPV1 := testhl7.PV1(t, result)
			resultPV1 := testhl7.PV1(t, result)
			if resultPV1.PendingLocation != nil {
				t.Errorf("resultPV1.PendingLocation = %+v want <nil>", resultPV1.PendingLocation)
			}
			if diff := cmp.Diff(admitPV1.AssignedPatientLocation, resultPV1.AssignedPatientLocation); diff != "" {
				t.Errorf("resultPV1.AssignedPatientLocation (-want, +got):\n%s", diff)
			}

			// The transfer didn't happen.
			if got, want := hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds(), 0; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds()=%v, want %v", got, want)
			}
			if got, want := hospital.LocationManager.RoomManagers[testLocAE].OccupiedBeds(), 1; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLocAE].OccupiedBeds()=%v, want %v", got, want)
			}
		},
	}, {
		name: "CancelTransfer",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{PendingTransfer: &pathway.PendingTransfer{Loc: testLoc, ExpectedTransferTimeFromNow: &twoHours}},
			{TransferInError: &pathway.TransferInError{Loc: testLoc}},
			{CancelTransfer: &pathway.CancelTransfer{}},
		}},
		wantMessageTypes: []string{"ADT^A01", "ADT^A15", "ADT^A02", "ADT^A12"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			pendingTransfer, cancel := messages[1], messages[3]

			pendingTransferEVN := testhl7.EVN(t, pendingTransfer)
			cancelEVN := testhl7.EVN(t, cancel)

			if diff := cmp.Diff(pendingTransferEVN.DateTimePlannedEvent, cancelEVN.EventOccurred); diff != "" {
				t.Errorf("cancelEVN.EventOccurred diff (-want, +got):\n%s", diff)
			}

			if got, want := hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds(), 0; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds()=%v, want %v", got, want)
			}
			if got, want := hospital.LocationManager.RoomManagers[testLocAE].OccupiedBeds(), 1; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLocAE].OccupiedBeds()=%v, want %v", got, want)
			}
		},
	}, {
		name: "Transfer after PendingTransfer",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{PendingTransfer: &pathway.PendingTransfer{Loc: testLoc, ExpectedTransferTimeFromNow: &twoHours}},
			{Transfer: &pathway.Transfer{Loc: testLoc}},
		}},
		wantMessageTypes: []string{"ADT^A01", "ADT^A15", "ADT^A02"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			pendingTransfer, transfer := messages[1], messages[2]
			pendingTransferPV1 := testhl7.PV1(t, pendingTransfer)
			transferPV1 := testhl7.PV1(t, transfer)

			if got, want := pendingTransferPV1.PendingLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLoc].Poc; got != want {
				t.Errorf("pendingTransferPV1.PendingLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if transferPV1.PendingLocation != nil {
				t.Error("transferPV1.PendingLocation is something, want <nil>.")
			}
			pendingTransferEVN := testhl7.EVN(t, pendingTransfer)
			transferEVN := testhl7.EVN(t, transfer)
			if diff := cmp.Diff(pendingTransferEVN.DateTimePlannedEvent, transferEVN.RecordedDateTime); diff != "" {
				t.Errorf("transferEVN.RecordedDateTime diff (-want, +got):\n%s", diff)
			}

			if got, want := hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds(), 1; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds()=%v, want %v", got, want)
			}
			if got, want := hospital.LocationManager.RoomManagers[testLocAE].OccupiedBeds(), 0; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLocAE].OccupiedBeds()=%v, want %v", got, want)
			}
		},
	}, {
		name: "Allergies in Admission DischargeInError Discharge",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{AddPerson: &pathway.AddPerson{Allergies: []pathway.Allergy{{Description: "Allergy1"}}}},
			{Registration: &pathway.Registration{Allergies: []pathway.Allergy{{Description: "Allergy2"}}}},
			{Admission: &pathway.Admission{Loc: testLoc, Allergies: []pathway.Allergy{{Description: "Allergy3"}}}},
			{DischargeInError: &pathway.DischargeInError{Allergies: []pathway.Allergy{{Description: "Allergy4"}}}},
			{CancelDischarge: &pathway.CancelDischarge{}},
			{Discharge: &pathway.Discharge{Allergies: []pathway.Allergy{{Description: "Allergy5"}}}},
		}},
		wantMessageTypes: []string{"ADT^A28", "ADT^A04", "ADT^A01", "ADT^A03", "ADT^A13", "ADT^A03"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			// Allergies accumulate between steps, and cancelDischarge steps don't contain allergies.
			want := []int{1, 2, 3, 4, 0, 5}
			var got []int
			for _, m := range messages {
				got = append(got, len(testhl7.AllAL1(t, m)))
			}
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("Number of allergies mismatch (-want, +got):\n%s", diff)
			}
		},
	}, {
		name: "Pending Admission",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{PendingAdmission: &pathway.PendingAdmission{Loc: testLoc, ExpectedAdmissionTimeFromNow: &twoHours}},
			{Admission: &pathway.Admission{Loc: testLoc}},
		}},
		wantMessageTypes: []string{"ADT^A14", "ADT^A01"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			wantAccountStatus := []string{
				hospital.MessageConfig.PatientAccountStatus.Planned,
				hospital.MessageConfig.PatientAccountStatus.Arrived,
			}
			gotAccountStatus := testhl7.Fields(t, messages, testhl7.AccountStatus)
			if diff := cmp.Diff(wantAccountStatus, gotAccountStatus); diff != "" {
				t.Errorf("StartPathway(%v) generated AccountStatus with diff (-want, +got):\n%s", testPathwayName, diff)
			}

			pendingAdmission, admission := messages[0], messages[1]
			pendingAdmissionPV1 := testhl7.PV1(t, pendingAdmission)
			pendingAdmissionPV2 := testhl7.PV2(t, pendingAdmission)
			pendingAdmissionEVN := testhl7.EVN(t, pendingAdmission)
			admissionPV1 := testhl7.PV1(t, admission)

			if got, want := pendingAdmissionPV1.PendingLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLoc].Poc; got != want {
				t.Errorf("pendingAdmissionPV1.PendingLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if admissionPV1.PendingLocation != nil {
				t.Errorf("admissionPV1.PendingLocation is %+v, want <nil>", admissionPV1.PendingLocation)
			}
			if got, want := *admissionPV1.AdmitDateTime, *pendingAdmissionPV2.ExpectedAdmitDateTime; got != want {
				t.Errorf("*admitPV1.AdmitDateTime=%v, want %v", got, want)
			}
			if got, want := *pendingAdmissionEVN.DateTimePlannedEvent, *pendingAdmissionPV2.ExpectedAdmitDateTime; got != want {
				t.Errorf("*pendingAdmissionEVN.DateTimePlannedEvent=%v, want %v", got, want)
			}
		},
	}, {
		name: "Cancel Pending Admission",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{PendingAdmission: &pathway.PendingAdmission{Loc: testLoc, ExpectedAdmissionTimeFromNow: &twoHours}},
			{CancelPendingAdmission: &pathway.CancelPendingAdmission{}},
			{Result: &pathway.Results{}},
		}},
		wantMessageTypes: []string{"ADT^A14", "ADT^A27", "ORU^R01"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			wantAccountStatus := []string{
				hospital.MessageConfig.PatientAccountStatus.Planned,
				hospital.MessageConfig.PatientAccountStatus.Cancelled,
				hospital.MessageConfig.PatientAccountStatus.Cancelled,
			}
			gotAccountStatus := testhl7.Fields(t, messages, testhl7.AccountStatus)
			if diff := cmp.Diff(wantAccountStatus, gotAccountStatus); diff != "" {
				t.Errorf("StartPathway(%v) generated AccountStatus with diff (-want, +got):\n%s", testPathwayName, diff)
			}

			pendingAdmission, cancelPendingAdmission, result := messages[0], messages[1], messages[2]
			cancelPV2 := testhl7.PV2(t, cancelPendingAdmission)
			if got, want := cancelPV2.PriorPendingLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLoc].Poc; got != want {
				t.Errorf("cancelPV2.PriorPendingLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			// The Cancellation event contains the date of the planned admission.
			cancelPendingAdmissionEVN := testhl7.EVN(t, cancelPendingAdmission)
			pendingAdmissionEVN := testhl7.EVN(t, pendingAdmission)
			if got, want := *cancelPendingAdmissionEVN.EventOccurred, *pendingAdmissionEVN.DateTimePlannedEvent; got != want {
				t.Errorf("*cancelPendingAdmissionEVN.EventOccurred=%v, want %v", got, want)
			}

			// The admission didn't happen: the patient should be an outpatient.
			wantPatientClass := []string{"OUTPATIENT", "OUTPATIENT", "OUTPATIENT"}
			gotPatientClass := testhl7.Fields(t, messages, testhl7.PatientClass)
			if diff := cmp.Diff(wantPatientClass, gotPatientClass); diff != "" {
				t.Errorf("StartPathway(%v) generated PatientClass with diff (-want, +got):\n%s", testPathwayName, diff)
			}
			resultPV1 := testhl7.PV1(t, result)
			if resultPV1.PendingLocation != nil {
				t.Errorf("resultPV1.PendingLocation is %+v, want <nil>", resultPV1.PendingLocation)
			}
			if got, want := hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds(), 0; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds()=%v, want %v", got, want)
			}
		},
	}, {
		name: "Pending Discharge",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{PendingDischarge: &pathway.PendingDischarge{ExpectedDischargeTimeFromNow: &twoHours}},
			{Discharge: &pathway.Discharge{}},
		}},
		wantMessageTypes: []string{"ADT^A16", "ADT^A03"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			pendingDischarge, discharge := messages[0], messages[1]
			pendingDischargePV2 := testhl7.PV2(t, pendingDischarge)
			pendingDischargeEVN := testhl7.EVN(t, pendingDischarge)
			dischargePV1 := testhl7.PV1(t, discharge)

			if got, want := len(dischargePV1.DischargeDateTime), 1; got != want {
				t.Errorf("len(dischargePV1.DischargeDateTime)=%v, want %v", got, want)
			}
			if got, want := dischargePV1.DischargeDateTime[0], *pendingDischargePV2.ExpectedDischargeDateTime; got != want {
				t.Errorf("dischargePV1.DischargeDateTime[0]=%v, want %v", got, want)
			}
			if diff := cmp.Diff(pendingDischargePV2.ExpectedDischargeDateTime, pendingDischargeEVN.DateTimePlannedEvent); diff != "" {
				t.Errorf("pendingDischargeEVN.DateTimePlannedEvent diff (-want, +got):\n%s", diff)
			}
		},
	}, {
		name: "Cancel Pending Discharge",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{PendingDischarge: &pathway.PendingDischarge{ExpectedDischargeTimeFromNow: &twoHours}},
			{CancelPendingDischarge: &pathway.CancelPendingDischarge{}},
			{Result: &pathway.Results{}},
		}},
		wantMessageTypes: []string{"ADT^A01", "ADT^A16", "ADT^A25", "ORU^R01"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			pendingDischarge, cancelPendingDischarge, result := messages[1], messages[2], messages[3]
			// The Cancellation event contains the date of the planned discharge.
			cancelPendingDischargeEVN := testhl7.EVN(t, cancelPendingDischarge)
			pendingDischargeEVN := testhl7.EVN(t, pendingDischarge)
			if diff := cmp.Diff(pendingDischargeEVN.DateTimePlannedEvent, cancelPendingDischargeEVN.EventOccurred); diff != "" {
				t.Errorf("cancelPendingDischargeEVN.EventOccurred diff (-want, +got):\n%s", diff)
			}

			// The discharge didn't happen: The location should have been the same as the original location,
			// and the patient should still be an inpatient.
			wantPatientClass := []string{"INPATIENT", "INPATIENT", "INPATIENT", "INPATIENT"}
			gotPatientClass := testhl7.Fields(t, messages, testhl7.PatientClass)
			if diff := cmp.Diff(wantPatientClass, gotPatientClass); diff != "" {
				t.Errorf("StartPathway(%v) generated PatientClass with diff (-want, +got):\n%s", testPathwayName, diff)
			}
			admitPV1 := testhl7.PV1(t, result)
			resultPV1 := testhl7.PV1(t, result)
			if diff := cmp.Diff(admitPV1.AssignedPatientLocation, resultPV1.AssignedPatientLocation); diff != "" {
				t.Errorf("resultPV1.AssignedPatientLocation diff (-want, +got):\n%s", diff)
			}
			if got, want := resultPV1.PatientClass.String(), hospital.MessageConfig.PatientClass.Inpatient; got != want {
				t.Errorf("resultPV1.PatientClass.String()=%v, want %v", got, want)
			}
			if got, want := hospital.LocationManager.RoomManagers[testLocAE].OccupiedBeds(), 1; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLocAE].OccupiedBeds()=%v, want %v", got, want)
			}
		},
	}, {
		name: "TrackDeparture_Track",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{TrackDeparture: &pathway.TrackDeparture{DestinationLoc: testLoc, Mode: pathway.TrackMode}},
		}},
		wantMessageTypes: []string{"ADT^A01", "ADT^A09"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			trackDeparture := messages[1]
			trackDeparturePV1 := testhl7.PV1(t, trackDeparture)

			if got, want := trackDeparturePV1.PriorPatientLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLocAE].Poc; got != want {
				t.Errorf("trackDeparturePV1.PriorPatientLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if got, want := trackDeparturePV1.AssignedPatientLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLoc].Poc; got != want {
				t.Errorf("trackDeparturePV1.AssignedPatientLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if got, want := hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds(), 1; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds()=%v, want %v", got, want)
			}
			if got, want := hospital.LocationManager.RoomManagers[testLocAE].OccupiedBeds(), 0; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLocAE].OccupiedBeds()=%v, want %v", got, want)
			}
		},
	}, {
		name: "TrackArrival_Track",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{TrackArrival: &pathway.TrackArrival{Loc: testLoc, Mode: pathway.TrackMode}},
		}},
		wantMessageTypes: []string{"ADT^A01", "ADT^A10"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			trackArrival := messages[1]
			trackArrivalPV1 := testhl7.PV1(t, trackArrival)

			if got, want := trackArrivalPV1.PriorPatientLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLocAE].Poc; got != want {
				t.Errorf("trackArrivalPV1.PriorPatientLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if got, want := trackArrivalPV1.AssignedPatientLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLoc].Poc; got != want {
				t.Errorf("trackArrivalPV1.AssignedPatientLocation.PointOfCare.String()=%v, want %v", got, want)
			}

			if got, want := hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds(), 1; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds()=%v, want %v", got, want)
			}
			if got, want := hospital.LocationManager.RoomManagers[testLocAE].OccupiedBeds(), 0; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLocAE].OccupiedBeds()=%v, want %v", got, want)
			}
		},
	}, {
		name: "TrackDeparture_Transit and TrackArrival_Transit",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{TrackDeparture: &pathway.TrackDeparture{DestinationLoc: testLoc, Mode: pathway.TransitMode}},
			{TrackArrival: &pathway.TrackArrival{Loc: testLoc, Mode: pathway.TransitMode}},
		}},
		wantMessageTypes: []string{"ADT^A01", "ADT^A09", "ADT^A10"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			transitDeparture, transitArrival := messages[1], messages[2]

			transitDeparturePV1 := testhl7.PV1(t, transitDeparture)
			transitArrivalPV1 := testhl7.PV1(t, transitArrival)

			if got, want := transitDeparturePV1.PriorPatientLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLocAE].Poc; got != want {
				t.Errorf("transitDeparturePV1.PriorPatientLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if got, want := transitDeparturePV1.PendingLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLoc].Poc; got != want {
				t.Errorf("transitDeparturePV1.PendingLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if transitDeparturePV1.AssignedPatientLocation != nil {
				t.Errorf("transitDeparturePV1.AssignedPatientLocation is %+v, want <nil>", transitDeparturePV1.AssignedPatientLocation)
			}

			if got, want := transitArrivalPV1.AssignedPatientLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLoc].Poc; got != want {
				t.Errorf("transitArrivalPV1.AssignedPatientLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if transitArrivalPV1.PendingLocation != nil {
				t.Errorf("transitArrivalPV1.PendingLocation is %+v, want <nil>", transitArrivalPV1.PendingLocation)
			}

			if got, want := hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds(), 1; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds()=%v, want %v", got, want)
			}
			if got, want := hospital.LocationManager.RoomManagers[testLocAE].OccupiedBeds(), 0; got != want {
				t.Errorf("hospital.LocationManager.RoomManagers[testLocAE].OccupiedBeds()=%v, want %v", got, want)
			}
		},
	}, {
		name: "TrackDeparture_Temporary and TrackArrival_Temporary",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{TrackDeparture: &pathway.TrackDeparture{DestinationLoc: "Hallway", Mode: pathway.TemporaryMode}},
			{TrackDeparture: &pathway.TrackDeparture{DestinationLoc: "X-RAY", Mode: pathway.TemporaryMode}},
			{TrackArrival: &pathway.TrackArrival{Loc: "X-RAY", Mode: pathway.TemporaryMode, IsTemporary: true}},
			{TrackArrival: &pathway.TrackArrival{Loc: testLocAE, Mode: pathway.TemporaryMode}},
		}},
		wantMessageTypes: []string{"ADT^A01", "ADT^A09", "ADT^A09", "ADT^A10", "ADT^A10"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			firstTempDeparture, secondTempDeparture, firstTempArrival, secondTempArrival := messages[1], messages[2], messages[3], messages[4]

			firstTempDeparturePV1 := testhl7.PV1(t, firstTempDeparture)
			secondTempDeparturePV1 := testhl7.PV1(t, secondTempDeparture)
			firstTempArrivalPV1 := testhl7.PV1(t, firstTempArrival)
			secondTempArrivalPV1 := testhl7.PV1(t, secondTempArrival)

			if got, want := firstTempDeparturePV1.PriorPatientLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLocAE].Poc; got != want {
				t.Errorf("firstTempDeparturePV1.PriorPatientLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if got, want := firstTempDeparturePV1.TemporaryLocation.PointOfCare.String(), "Hallway"; got != want {
				t.Errorf("firstTempDeparturePV1.TempLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if firstTempDeparturePV1.AssignedPatientLocation != nil {
				t.Errorf("firstTempDeparturePV1.AssignedPatientLocation is %+v, want <nil>", firstTempDeparturePV1.AssignedPatientLocation)
			}
			if got, want := secondTempDeparturePV1.PriorPatientLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLocAE].Poc; got != want {
				t.Errorf("secondTempDeparturePV1.PriorPatientLocation.PointOfCare.String()=%v, want %v", got, want)
			}

			if got, want := firstTempArrivalPV1.PriorTemporaryLocation.PointOfCare.String(), "Hallway"; got != want {
				t.Errorf("firstTempArrivalPV1.PriorTempLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if got, want := firstTempArrivalPV1.TemporaryLocation.PointOfCare.String(), "X-RAY"; got != want {
				t.Errorf("firstTempArrivalPV1.TempLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if firstTempArrivalPV1.AssignedPatientLocation != nil {
				t.Errorf("firstTempArrivalPV1.AssignedPatientLocation is %+v, want <nil>", firstTempArrivalPV1.AssignedPatientLocation)
			}

			if got, want := secondTempArrivalPV1.PriorTemporaryLocation.PointOfCare.String(), "X-RAY"; got != want {
				t.Errorf("secondTemporaryArrivalPV1.PriorTemporaryLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if got, want := secondTempArrivalPV1.AssignedPatientLocation.PointOfCare.String(), hospital.LocationManager.RoomManagers[testLocAE].Poc; got != want {
				t.Errorf("secondTemporaryArrivalPV1.AssignedPatientLocation.PointOfCare.String()=%v, want %v", got, want)
			}
			if secondTempArrivalPV1.TemporaryLocation != nil {
				t.Errorf("secondTemporaryArrivalPV1.TemporaryLocation is %+v, want <nil>", secondTempArrivalPV1.TemporaryLocation)
			}
		},
	}, {
		name: "DeleteVisit",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{Discharge: &pathway.Discharge{}},
			{DeleteVisit: &pathway.DeleteVisit{}},
		}},
		wantMessageTypes: []string{"ADT^A01", "ADT^A03", "ADT^A23"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			admit, deleteVisit := messages[0], messages[2]
			visitID := testhl7.PV1(t, admit).VisitNumber.IDNumber.SanitizedString()
			deleteVisitID := testhl7.PV1(t, deleteVisit).VisitNumber.IDNumber.SanitizedString()
			if got, want := deleteVisitID, visitID; got != want {
				t.Errorf("deleteVisitID=%v, want %v", got, want)
			}
		},
	}, {
		name: "DeleteVisit_NoPastVisit",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{DeleteVisit: &pathway.DeleteVisit{}},
		}},
		wantMessageTypes: nil, // The event fails and no messages are generated.
		wantMetrics: []metric{{
			name: "simulated_hospital_errors_total",
			labels: map[string]string{
				"pathway_name": testPathwayName,
				"reason":       "Patient PastVisits empty",
			},
			wantDiff: 1,
		}},
	}, {
		name: "DeleteVisit_PastVisitAlreadyDeleted",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{Discharge: &pathway.Discharge{}},
			{DeleteVisit: &pathway.DeleteVisit{}},
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{DeleteVisit: &pathway.DeleteVisit{}},
		}},
		// The second delete fails because there all past visits have been deleted.
		wantMessageTypes: []string{"ADT^A01", "ADT^A03", "ADT^A23", "ADT^A01"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			admit, deleteVisit := messages[0], messages[2]

			visitID := testhl7.PV1(t, admit).VisitNumber.IDNumber.SanitizedString()
			deleteVisitID := testhl7.PV1(t, deleteVisit).VisitNumber.IDNumber.SanitizedString()
			if got, want := deleteVisitID, visitID; got != want {
				t.Errorf("deleteVisitID=%v, want %v", got, want)
			}
		},
		wantMetrics: []metric{{
			name: "simulated_hospital_errors_total",
			labels: map[string]string{
				"pathway_name": testPathwayName,
				"reason":       "Patient PastVisits empty",
			},
			wantDiff: 1,
		}},
	}, {
		name: "DeleteVisit_WithOngoingVisit",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{Discharge: &pathway.Discharge{}},
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{DeleteVisit: &pathway.DeleteVisit{}},
			{Discharge: &pathway.Discharge{}},
		}},
		wantMessageTypes: []string{"ADT^A01", "ADT^A03", "ADT^A01", "ADT^A23", "ADT^A03"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			// The first visit has been deleted, and the second is correctly discharged.
			admit, admit2, deleteVisit, discharge2 := messages[0], messages[2], messages[3], messages[4]

			visitID := testhl7.PV1(t, admit).VisitNumber.IDNumber.SanitizedString()
			deleteVisitID := testhl7.PV1(t, deleteVisit).VisitNumber.IDNumber.SanitizedString()
			if got, want := deleteVisitID, visitID; got != want {
				t.Errorf("deleteVisitID=%v, want %v", got, want)
			}

			admit2VisitID := testhl7.PV1(t, admit2).VisitNumber.IDNumber.SanitizedString()
			discharge2VisitID := testhl7.PV1(t, discharge2).VisitNumber.IDNumber.SanitizedString()
			if got, want := discharge2VisitID, admit2VisitID; got != want {
				t.Errorf("discharge2VisitID=%v, want %v", got, want)
			}
		},
	}, {
		name: "Result",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Result: &pathway.Results{}},
		}},
		wantMessageTypes: []string{"ORU^R01"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			resultORC := testhl7.ORC(t, messages[0])
			if got, want := resultORC.OrderControl.String(), hospital.MessageConfig.OrderControl.WithObservations; got != want {
				t.Errorf("resultORC.OrderControl.String() = %q, want %q", got, want)
			}
		},
	}, {
		name: "Result_TriggerEvent_Empty",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Result: &pathway.Results{}},
		}},
		wantMessageTypes: []string{"ORU^R01"},
	}, {
		name: "Result_TriggerEvent_R01",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Result: &pathway.Results{TriggerEvent: "R01"}},
		}},
		wantMessageTypes: []string{"ORU^R01"},
	}, {
		name: "Result_TriggerEvent_r01",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Result: &pathway.Results{TriggerEvent: "r01"}},
		}},
		wantMessageTypes: []string{"ORU^R01"},
	}, {
		name: "Result_TriggerEvent_r03",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Result: &pathway.Results{TriggerEvent: "r03"}},
		}},
		wantMessageTypes: []string{"ORU^R03"},
	}, {
		name: "Result_TriggerEvent_R03",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Result: &pathway.Results{TriggerEvent: "R03"}},
		}},
		wantMessageTypes: []string{"ORU^R03"},
	}, {
		name: "Result_TriggerEvent_R32",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Result: &pathway.Results{TriggerEvent: "R32"}},
		}},
		wantMessageTypes: []string{"ORU^R32"},
	}, {
		name: "Result_TriggerEvent_other_trigger_event",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Result: &pathway.Results{TriggerEvent: "something-else"}},
		}},
		wantMessageTypes: []string{"ORU^R01"},
	}, {
		name: "UsePatient",
		pathway: pathway.Pathway{
			Persons: &pathway.Persons{
				"main-patient": {
					FirstName: "First first name",
				},
			},
			Pathway: []pathway.Step{
				{AddPerson: &pathway.AddPerson{}},
				{Result: &pathway.Results{}},
				{
					UpdatePerson: &pathway.UpdatePerson{
						Person: &pathway.Person{
							FirstName: "Second first name",
						},
						Diagnoses: []*pathway.DiagnosisOrProcedure{
							{Code: "A01.1", DateTime: &pathway.DateTime{TimeFromNow: &oneDayAgo}},
						},
						Procedures: []*pathway.DiagnosisOrProcedure{
							{Code: "P01.1", DateTime: &pathway.DateTime{TimeFromNow: &oneDayAgo}},
						},
						Allergies: []pathway.Allergy{
							{Description: "Allergy1", IdentificationDateTime: &pathway.DateTime{Time: &now}},
						},
					},
				},
				{Admission: &pathway.Admission{Loc: testLoc}},
				{
					UpdatePerson: &pathway.UpdatePerson{
						Person: &pathway.Person{
							FirstName: "Third first name",
						},
						Diagnoses: []*pathway.DiagnosisOrProcedure{
							{Code: "A02.2", DateTime: &pathway.DateTime{TimeFromNow: &oneDayAgo}},
						},
						Allergies: []pathway.Allergy{
							{Description: "Allergy2", IdentificationDateTime: &pathway.DateTime{Time: &now}},
						},
					},
				},
			},
		},
		wantMessageTypes: []string{"ADT^A28", "ORU^R01", "ADT^A31", "ADT^A01", "ADT^A08"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			addPerson, updatePerson, updatePatient := messages[0], messages[2], messages[4]
			wantFirstNames := []string{
				"First first name", "First first name", "Second first name", "Second first name", "Third first name",
			}
			gotFirstNames := testhl7.Fields(t, messages, testhl7.FirstName)
			if diff := cmp.Diff(wantFirstNames, gotFirstNames); diff != "" {
				t.Errorf("StartPathway(%v) generated message with patient first names with diff (-want, +got):\n%s", testPathwayName, diff)
			}
			mrn1 := testhl7.MRN(t, addPerson)
			// The MRN shouldn't change, only the first name.
			wantMRNs := []string{mrn1, mrn1, mrn1, mrn1, mrn1}
			gotMRNs := testhl7.Fields(t, messages, testhl7.MRN)
			if diff := cmp.Diff(wantMRNs, gotMRNs); diff != "" {
				t.Errorf("StartPathway(%v) generated message with patient MRNs with diff (-want, +got):\n%s", testPathwayName, diff)
			}
			// Some allergies might have been added in this step already.
			originalAllergies := len(testhl7.AllAL1(t, addPerson))

			// Assert procedure and diagnosis populated.
			if got, want := len(testhl7.AllDG1(t, updatePerson)), 1; got != want {
				t.Errorf("len(getAllDG1(t, updatePerson))=%v, want %v", got, want)
			}
			if got, want := len(testhl7.AllPR1(t, updatePerson)), 1; got != want {
				t.Errorf("len(getAllPR1(t, updatePerson))=%v, want %v", got, want)
			}
			if got, want := len(testhl7.AllAL1(t, updatePerson)), originalAllergies+1; got != want {
				t.Errorf("len(getAllAL1(t, updatePerson))=%v, want %v", got, want)
			}

			// Procedures and diagnoses are cleaned up between steps; allergies accumulated.
			if got, want := len(testhl7.AllDG1(t, updatePatient)), 1; got != want {
				t.Errorf("len(getAllDG1(t, updatePatient))=%v, want %v", got, want)
			}
			if got, want := len(testhl7.AllPR1(t, updatePatient)), 0; got != want {
				t.Errorf("len(getAllPR1(t, updatePatient))=%v, want %v", got, want)
			}
			if got, want := len(testhl7.AllAL1(t, updatePatient)), originalAllergies+2; got != want {
				t.Errorf("len(getAllAL1(t, updatePatient))=%v, want %v", got, want)
			}
		},
	}, {
		name: "UsePatient_FromPersonSection",
		pathway: pathway.Pathway{
			Persons: &pathway.Persons{
				"patient-1": {FirstName: "Patient 1"},
				"patient-2": {FirstName: "Patient 2"},
			},
			Pathway: []pathway.Step{
				{UsePatient: &pathway.UsePatient{Patient: "patient-1"}},
				{Admission: &pathway.Admission{Loc: testLoc}},
				{UsePatient: &pathway.UsePatient{Patient: "patient-2"}},
				{Admission: &pathway.Admission{Loc: testLoc}},
				{UsePatient: &pathway.UsePatient{Patient: "patient-1"}},
				{Discharge: &pathway.Discharge{}},
				{UsePatient: &pathway.UsePatient{Patient: "patient-2"}},
				{Discharge: &pathway.Discharge{}},
			},
		},
		wantMessageTypes: []string{"ADT^A01", "ADT^A01", "ADT^A03", "ADT^A03"},
		want: func(t *testing.T, messages []string, hospital *testhospital.Hospital) {
			admit1, admit2, discharge1, discharge2 := messages[0], messages[1], messages[2], messages[3]
			// Note we cannot make the assertions using the MRN, as we don't know in what order the two items
			// in Persons have been processed.
			if got, want := testhl7.FirstName(t, admit1), "Patient 1"; got != want {
				t.Errorf("Got %v, want %v.", got, want)
			}
			if got, want := testhl7.FirstName(t, admit2), "Patient 2"; got != want {
				t.Errorf("Got %v, want %v.", got, want)
			}
			if got, want := testhl7.FirstName(t, discharge1), "Patient 1"; got != want {
				t.Errorf("Got %v, want %v.", got, want)
			}
			if got, want := testhl7.FirstName(t, discharge2), "Patient 2"; got != want {
				t.Errorf("Got %v, want %v.", got, want)
			}
		},
	}, {
		name: "UsePatient_Itself",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
			// This is the current patient's MRN.
			{UsePatient: &pathway.UsePatient{Patient: "1"}},
			{Result: &pathway.Results{}},
			{Discharge: &pathway.Discharge{}},
		}},
		wantMessageTypes: []string{"ADT^A01", "ORU^R01", "ADT^A03"},
	}, {
		name: "UsePatient with unknown MRN",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{UsePatient: &pathway.UsePatient{Patient: "non-existent"}},
			{Result: &pathway.Results{}},
			{Discharge: &pathway.Discharge{}},
		}},
		wantMessageTypes: []string{"ADT^A01"},
		wantMetrics: []metric{{
			name: "simulated_hospital_errors_total",
			labels: map[string]string{
				"pathway_name": testPathwayName,
				"reason":       "unknown MRN in use_patient",
			},
			wantDiff: 1,
		}},
	}, {
		name: "TrackArrival_Track as first event",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{TrackArrival: &pathway.TrackArrival{Loc: testLoc, Mode: pathway.TransitMode}},
		}},
		wantMetrics: []metric{{
			name: "simulated_hospital_errors_total",
			labels: map[string]string{
				"pathway_name": testPathwayName,
				"reason":       "patient location error: nil patient location",
			},
			wantDiff: 1,
		}},
	}, {
		name: "TrackArrival_Track after discharge",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{Discharge: &pathway.Discharge{}},
			{TrackArrival: &pathway.TrackArrival{Loc: testLoc, Mode: pathway.TransitMode}},
		}},
		wantMessageTypes: []string{"ADT^A01", "ADT^A03"},
		wantMetrics: []metric{{
			name: "simulated_hospital_errors_total",
			labels: map[string]string{
				"pathway_name": testPathwayName,
				"reason":       "patient location error: nil patient location",
			},
			wantDiff: 1,
		}},
	}, {
		name: "Admission Discharge with GenerateResources",
		pathway: pathway.Pathway{Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{GenerateResources: &pathway.GenerateResources{}}, // Should not generate any messages.
			{Discharge: &pathway.Discharge{}},
		}},
		wantMessageTypes: []string{"ADT^A01", "ADT^A03"},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mr := testmetrics.NewRetrieverFromGatherer(t)

			hospital := newHospital(ctx, t, Config{}, map[string]pathway.Pathway{testPathwayName: tc.pathway})
			defer hospital.Close()
			startPathway(t, hospital, testPathwayName)
			_, messages := hospital.ConsumeQueues(ctx, t)
			if got, want := len(messages), len(tc.wantMessageTypes); got != want {
				t.Errorf("StartPathway(%v) generated %v messages, want %v", testPathwayName, got, want)
			}

			gotMessageTypes := testhl7.Fields(t, messages, testhl7.MessageType)
			if diff := cmp.Diff(tc.wantMessageTypes, gotMessageTypes); diff != "" {
				t.Errorf("StartPathway(%v) generated message types with diff (-want, +got):\n%s", testPathwayName, diff)
			}
			if tc.want != nil {
				tc.want(t, messages, hospital)
			}

			for _, m := range tc.wantMetrics {
				initial, final := mr.GetCounterValues(t, m.name, m.labels)
				if got, want := final-initial, m.wantDiff; got != want {
					t.Errorf("Metric %s[%v] is incremented by %f (initial=%f, final=%f); want %f", m.name, m.labels, got, initial, final, want)
				}
			}
		})
	}
}

func TestStartPathway_OccupiedBed(t *testing.T) {
	ctx := context.Background()
	type preoccupiedBed struct {
		loc   string
		bedID string
	}
	type wantOccupiedBed struct {
		loc  string
		want int
	}

	tests := []struct {
		name             string
		preoccupiedBeds  []preoccupiedBed
		steps            []pathway.Step
		wantMessageTypes []string
		wantOccupiedBeds []wantOccupiedBed
	}{{
		name: "Transfer and Discharge frees the bed",
		steps: []pathway.Step{
			{Transfer: &pathway.Transfer{Loc: testLoc}},
			{Discharge: &pathway.Discharge{}},
		},
		wantMessageTypes: []string{"ADT^A02", "ADT^A03"},
		wantOccupiedBeds: []wantOccupiedBed{{loc: testLoc, want: 0}},
	}, {
		name: "Transfer keeps the bed",
		steps: []pathway.Step{
			{Transfer: &pathway.Transfer{Loc: testLoc}},
		},
		wantMessageTypes: []string{"ADT^A02"},
		wantOccupiedBeds: []wantOccupiedBed{{loc: testLoc, want: 1}},
	}, {
		name: "Specific bed in Admission",
		steps: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE, Bed: "specific bed"}},
		},
		wantMessageTypes: []string{"ADT^A01"},
		wantOccupiedBeds: []wantOccupiedBed{{loc: testLocAE, want: 1}},
	}, {
		name: "Specific bed in Transfer",
		steps: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE, Bed: "specific bed"}},
			{Transfer: &pathway.Transfer{Loc: testLocAE, Bed: "a different bed"}},
		},
		wantMessageTypes: []string{"ADT^A01", "ADT^A02"},
		wantOccupiedBeds: []wantOccupiedBed{{loc: testLocAE, want: 1}},
	}, {
		name: "Admit to occupied bed",
		steps: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE, Bed: "bed 1"}},
			{Admission: &pathway.Admission{Loc: testLocAE, Bed: "bed 1"}},
		},
		// The transfer event fails because we can't occupy an already occupied bed.
		wantMessageTypes: []string{"ADT^A01"},
		wantOccupiedBeds: []wantOccupiedBed{{loc: testLocAE, want: 1}},
	}, {
		name: "Transfer to occupied bed",
		preoccupiedBeds: []preoccupiedBed{{
			loc: testLocAE, bedID: "arbitrary-preoccupied-bed",
		}},
		steps: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{Transfer: &pathway.Transfer{Loc: testLocAE, Bed: "arbitrary-preoccupied-bed"}},
		},
		// The transfer event fails because we can't occupy an already occupied bed.
		wantMessageTypes: []string{"ADT^A01"},
		wantOccupiedBeds: []wantOccupiedBed{
			{loc: testLocAE, want: 1},
			{loc: testLoc, want: 0},
		},
	}, {
		name: "Cancel Visit after Admission",
		steps: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{CancelVisit: &pathway.CancelVisit{}},
		},
		wantMessageTypes: []string{"ADT^A01", "ADT^A11"},
		wantOccupiedBeds: []wantOccupiedBed{{loc: testLoc, want: 0}},
	}, {
		name: "Cancel Visit after Registration",
		steps: []pathway.Step{
			{Registration: &pathway.Registration{}},
			{CancelVisit: &pathway.CancelVisit{}},
		},
		wantMessageTypes: []string{"ADT^A04", "ADT^A11"},
		wantOccupiedBeds: []wantOccupiedBed{{loc: testLoc, want: 0}},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pathways := map[string]pathway.Pathway{
				testPathwayName: {Pathway: tc.steps},
			}
			cfg := Config{
				LocationManager: testlocation.NewLocationManager(ctx, t, testLoc, testLocAE),
			}
			hospital := newHospital(ctx, t, cfg, pathways)
			defer hospital.Close()
			for _, preoccupiedBed := range tc.preoccupiedBeds {
				if _, err := hospital.LocationManager.OccupySpecificBed(preoccupiedBed.loc, preoccupiedBed.bedID); err != nil {
					t.Fatalf("OccupySpecificBed(%s, %s) failed with %v", preoccupiedBed.loc, preoccupiedBed.bedID, err)
				}
			}

			startPathway(t, hospital, testPathwayName)
			_, messages := hospital.ConsumeQueues(ctx, t)
			if got, want := len(messages), len(tc.wantMessageTypes); got != want {
				t.Errorf("StartPathway(%v) generated %v messages, want %v", testPathwayName, got, want)
			}

			gotMessageTypes := testhl7.Fields(t, messages, testhl7.MessageType)
			if diff := cmp.Diff(tc.wantMessageTypes, gotMessageTypes); diff != "" {
				t.Errorf("StartPathway(%v) generated message types with diff (-want, +got):\n%s", testPathwayName, diff)
			}

			for _, wantOccupiedBeds := range tc.wantOccupiedBeds {
				if got, want := hospital.LocationManager.RoomManagers[wantOccupiedBeds.loc].OccupiedBeds(), wantOccupiedBeds.want; got != want {
					t.Errorf("hospital.LocationManager.RoomManagers[%v].OccupiedBeds()=%v, want %v", wantOccupiedBeds.loc, got, want)
				}
			}
		})
	}
}

// TestRunPathway_UpdateDeathInformation tests the ability to set death information in individual steps.
// TrackArrival in 'transit' mode, DeleteVisit, Merge and BedSwap require a complex setup and they cannot be run individually.
// Given that there's a low chance that those steps will be used to update death information, they aren't tested.
func TestRunPathway_UpdateDeathInformation(t *testing.T) {
	ctx := context.Background()
	now := time.Date(2018, 2, 12, 0, 30, 0, 0, time.UTC)
	testTime := 3 * time.Minute
	threeMinutes := 3 * time.Minute
	threeMinutesAgo := now.Add(-threeMinutes)

	tests := []struct {
		name               string
		params             *pathway.Parameters
		wantDeathIndicator string
		wantDeathTime      *time.Time
		wantLocationFree   string
	}{{
		name: "Death Indicator Y, TimeSinceDeath",
		params: &pathway.Parameters{
			Status: &pathway.DeathStatus{DeathIndicator: "Y", TimeSinceDeath: &threeMinutes},
		},
		wantDeathIndicator: "Y",
		wantDeathTime:      &threeMinutesAgo,
		wantLocationFree:   testLoc,
	}, {
		name: "Death Indicator N",
		params: &pathway.Parameters{
			Status: &pathway.DeathStatus{DeathIndicator: "N"},
		},
		wantDeathIndicator: "N",
		wantDeathTime:      nil,
		wantLocationFree:   "",
	}, {
		name: "Death Indicator Y, TimeOfDeath",
		params: &pathway.Parameters{
			Status: &pathway.DeathStatus{DeathIndicator: "Y", TimeOfDeath: &now},
		},
		wantDeathIndicator: "Y",
		wantDeathTime:      &now,
		wantLocationFree:   testLoc,
	}, {
		name: "Death Indicator Y, TimeSinceDeath and TimeOfDeath",
		params: &pathway.Parameters{
			// Note that this pathway would be invalid because both TimeSinceDeath and TimeOfDeath are set at the same time,
			// but we keep the test because the Hospital handles this situation.
			Status: &pathway.DeathStatus{DeathIndicator: "Y", TimeSinceDeath: &threeMinutes, TimeOfDeath: &now},
		},
		wantDeathIndicator: "Y",
		wantDeathTime:      &now,
		wantLocationFree:   testLoc,
	}}

	steps := []pathway.Step{
		{Admission: &pathway.Admission{Loc: testLoc}},
		{AddPerson: &pathway.AddPerson{}},
		{Result: &pathway.Results{}},
		{Discharge: &pathway.Discharge{}},
		{Registration: &pathway.Registration{}},
		{UpdatePerson: &pathway.UpdatePerson{}},
		{PendingAdmission: &pathway.PendingAdmission{Loc: testLoc, ExpectedAdmissionTimeFromNow: &testTime}},
		{TrackDeparture: &pathway.TrackDeparture{DestinationLoc: testLoc, Mode: pathway.TrackMode}},
		{TrackDeparture: &pathway.TrackDeparture{DestinationLoc: testLoc, Mode: pathway.TemporaryMode}},
		{PreAdmission: &pathway.PreAdmission{Loc: testLoc, ExpectedAdmissionTimeFromNow: &testTime}},
		{TransferInError: &pathway.TransferInError{Loc: testLoc}},
		{Order: &pathway.Order{OrderProfile: "RANDOM", NoAcknowledgementMessage: true}},
		{Transfer: &pathway.Transfer{Loc: testLoc}},
		{DischargeInError: &pathway.DischargeInError{}},
		{CancelTransfer: &pathway.CancelTransfer{}},
		{CancelDischarge: &pathway.CancelDischarge{}},
		{CancelVisit: &pathway.CancelVisit{}},
		{PendingDischarge: &pathway.PendingDischarge{ExpectedDischargeTimeFromNow: &testTime}},
		{PendingTransfer: &pathway.PendingTransfer{Loc: testLoc, ExpectedTransferTimeFromNow: &testTime}},
		{CancelPendingDischarge: &pathway.CancelPendingDischarge{}},
		{CancelPendingAdmission: &pathway.CancelPendingAdmission{}},
		{CancelPendingTransfer: &pathway.CancelPendingTransfer{}},
		{TrackDeparture: &pathway.TrackDeparture{DestinationLoc: testLoc, Mode: pathway.TransitMode}},
		{TrackArrival: &pathway.TrackArrival{Loc: testLoc, Mode: pathway.TrackMode}},
		{TrackArrival: &pathway.TrackArrival{Loc: testLoc, Mode: pathway.TemporaryMode}},
	}

	for _, tc := range tests {
		for _, step := range steps {
			t.Run(fmt.Sprintf("DeathIndicator:%v-Step:%v", tc.params.Status.DeathIndicator, step.StepType()), func(t *testing.T) {
				step.Parameters = tc.params
				pathways := map[string]pathway.Pathway{
					testPathwayName: {Pathway: []pathway.Step{step, {Result: &pathway.Results{}}}},
				}

				h := hospitalWithTime(ctx, t, Config{}, pathways, now)
				defer h.Close()
				startPathway(t, h, testPathwayName)

				_, messages := h.ConsumeQueues(ctx, t)
				if got, want := len(messages), 2; got != want {
					t.Errorf("StartPathway(%v) generated %v messages, want %v", testPathwayName, got, want)
				}

				// Test both messages: the Death Indicator should remain for the next message.
				wantDeathIndicators := []string{tc.wantDeathIndicator, tc.wantDeathIndicator}
				gotDeathIndicators := testhl7.Fields(t, messages, testhl7.DeathIndicator)
				if diff := cmp.Diff(wantDeathIndicators, gotDeathIndicators); diff != "" {
					t.Errorf("StartPathway(%v) generated death indicator with diff (-want, +got):\n%s", testPathwayName, diff)
				}

				wantDeathTime := []*time.Time{tc.wantDeathTime, tc.wantDeathTime}
				gotDeathTime := testhl7.TimeFields(t, messages, testhl7.DeathTime)
				if diff := cmp.Diff(wantDeathTime, gotDeathTime); diff != "" {
					t.Errorf("StartPathway(%v) generated death time indicator with diff (-want, +got):\n%s", testPathwayName, diff)
				}

				if tc.wantLocationFree == "" {
					return
				}
				pv1 := testhl7.PV1(t, messages[0])
				if pv1.AssignedPatientLocation != nil {
					t.Errorf("PV1.AssignedPatientLocation = %+v, want <nil>", pv1.AssignedPatientLocation)
				}
				if pv1.PendingLocation != nil {
					t.Errorf("PV1.PendingLocation = %+v, want <nil>", pv1.PendingLocation)
				}
				if pv1.TemporaryLocation != nil {
					t.Errorf("PV1.TemporaryLocation = %+v, want <nil>", pv1.TemporaryLocation)
				}
				if got, want := h.LocationManager.RoomManagers[tc.wantLocationFree].OccupiedBeds(), 0; got != want {
					t.Errorf("h.locationManager.RoomManagers[%s].OccupiedBeds()=%v, want %v", tc.wantLocationFree, got, want)
				}

				// Not all messages have PV2 segments.
				if pv2 := testhl7.PV2OrNil(t, messages[0]); pv2 != nil {
					if pv2.ExpectedAdmitDateTime != nil {
						t.Errorf("PV2.ExpectedAdmitDateTime = %+v, want <nil>", pv2.ExpectedAdmitDateTime)
					}
					if pv2.ExpectedDischargeDateTime != nil {
						t.Errorf("PV2.ExpectedDischargeDateTime = %+v, want <nil>", pv2.ExpectedDischargeDateTime)
					}
				}
			})
		}
	}
}

func TestRunPathwayMergeStep(t *testing.T) {
	ctx := context.Background()
	// Both "CURRENT" and the explicit MRN should work.
	mergeSteps := []*pathway.Merge{
		{Parent: "2", Children: toPatientIDs("1")},
		{Parent: pathway.Current, Children: toPatientIDs("1")},
	}

	for i, mergeStep := range mergeSteps {
		t.Run(fmt.Sprintf("%d-%v", i, mergeStep.Parent), func(t *testing.T) {
			delay := &pathway.Delay{From: time.Second, To: 5 * time.Second}
			longDelay := &pathway.Delay{From: time.Minute, To: 5 * time.Minute}
			veryLongDelay := &pathway.Delay{From: time.Hour, To: 5 * time.Hour}

			pathways := map[string]pathway.Pathway{
				"pathway1": {Pathway: []pathway.Step{
					{Admission: &pathway.Admission{Loc: testLocAE}},
					{Delay: longDelay},
				}},
				"pathway2": {Pathway: []pathway.Step{
					{Admission: &pathway.Admission{Loc: testLoc}},
					{Delay: delay},
					{Merge: mergeStep},
					{Result: &pathway.Results{}},
					{Delay: veryLongDelay},
					{Discharge: &pathway.Discharge{}},
				}},
			}

			hospital := newHospital(ctx, t, Config{}, pathways)
			defer hospital.Close()
			startPathway(t, hospital, "pathway1", "pathway2")
			events, messages := hospital.ConsumeQueues(ctx, t)
			if got, want := events, 8; got != want {
				t.Errorf("StartPathway(%v) and StartPathway(%v) generated %v events, want %v", "pathway1", "pathway2", got, want)
			}
			if got, want := len(messages), 5; got != want {
				t.Errorf("StartPathway(%v) and StartPathway(%v) generated %v messages, want %v", "pathway1", "pathway2", got, want)
			}

			merge := messages[2]
			wantMainMRN := "2"
			wantMergeWithMRN := []string{"1"}
			if got, want := testhl7.MRN(t, merge), wantMainMRN; got != want {
				t.Errorf("StartPathway(%v) and StartPathway(%v) got merge message with MRN = %v, want %v", "pathway1", "pathway2", got, want)
			}
			gotMergeWithMRN := testhl7.PriorPatientIdentifierList(t, merge)
			if diff := cmp.Diff(wantMergeWithMRN, gotMergeWithMRN); diff != "" {
				t.Errorf("StartPathway(%v) and StartPathway(%v) got PriorPatientIdentifierList with diff (-want, +got):\n%s", "pathway1", "pathway2", diff)
			}
		})
	}
}

func TestRunPathwayMergeStepWrongMRNSyncerInUse(t *testing.T) {
	ctx := context.Background()
	mr := testmetrics.NewRetrieverFromGatherer(t)

	pathways := map[string]pathway.Pathway{
		"pathway1": {Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{Result: &pathway.Results{}},
			{Discharge: &pathway.Discharge{}},
		}},
		// When we get to the merge, the state is invalid as it's expectedSurname that the Parent is the
		// current patient.
		"pathway2": {Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{Merge: &pathway.Merge{Parent: "1", Children: toPatientIDs("2")}},
			{Result: &pathway.Results{}},
			{Discharge: &pathway.Discharge{}},
		}},
		// Same as pathway2 but with a UsePatient step before the Merge.
		"pathway2-use-patient": {Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{UsePatient: &pathway.UsePatient{Patient: "1"}},
			{Merge: &pathway.Merge{Parent: "1", Children: toPatientIDs("2")}},
			{Result: &pathway.Results{}},
			{Discharge: &pathway.Discharge{}},
		}},
	}

	hospital := hospitalWithPatientSyncer(ctx, t, Config{}, pathways, teststate.NewItemSyncer())
	defer hospital.Close()
	startPathway(t, hospital, "pathway1", "pathway2")
	events, messages := hospital.ConsumeQueues(ctx, t)
	// The last two events aren't run due to the failure in the Merge step.
	if got, want := events, 5; got != want {
		t.Errorf("StartPathway(%v) and StartPathway(%v) generated %v events, want %v", "pathway1", "pathway2", got, want)
	}
	if got, want := len(messages), 4; got != want {
		t.Errorf("StartPathway(%v) and StartPathway(%v) generated %v messages, want %v", "pathway1", "pathway2", got, want)
	}

	startPathway(t, hospital, "pathway2-use-patient")
	events, messages = hospital.ConsumeQueues(ctx, t)
	if got, want := events, 5; got != want {
		t.Errorf("StartPathway(%v) generated %v events, want %v", "pathway2-use-patient", got, want)
	}
	if got, want := len(messages), 4; got != want {
		t.Errorf("StartPathway(%v) generated %v messages, want %v", "pathway2-use-patient", got, want)
	}

	metric := "simulated_hospital_errors_total"
	labels := map[string]string{
		"pathway_name": "pathway2",
		"reason":       "invalid merge state",
	}
	initial, final := mr.GetCounterValues(t, metric, labels)
	if got, want := final-initial, 1.0; got != want {
		t.Errorf("Metric %s[%v] is incremented by %f (initial=%f, final=%f); want %f", metric, labels, got, initial, final, want)
	}

	metric = "simulated_hospital_pathway_duration_minutes"
	labels = map[string]string{
		"pathway_name": "pathway2-use-patient",
	}
	beforeH, afterH := mr.GetHistogramBuckets(t, metric, labels)
	wantAfter := mr.IncrementBucketsCumulatively(beforeH, 1, 1)
	if diff := cmp.Diff(wantAfter, afterH, protocmp.Transform()); diff != "" {
		t.Errorf("Metric %s[%v] histogram buckets not incremented as expected (-want/+got):\n%s", metric, labels, diff)
	}
}

func TestRunPathwayMergeStepPersonsSection(t *testing.T) {
	ctx := context.Background()
	p1ID := pathway.PatientID("patient-1")
	p2ID := pathway.PatientID("patient-2")

	pathways := map[string]pathway.Pathway{
		testPathwayName: {
			Persons: &pathway.Persons{
				p1ID: {FirstName: "Patient 1"},
				p2ID: {FirstName: "Patient 2"},
			},
			Pathway: []pathway.Step{
				{UsePatient: &pathway.UsePatient{Patient: p1ID}},
				{Admission: &pathway.Admission{Loc: testLocAE}},
				{UsePatient: &pathway.UsePatient{Patient: p2ID}},
				{Admission: &pathway.Admission{Loc: testLoc}},
				{Merge: &pathway.Merge{Parent: p2ID, Children: []pathway.PatientID{p1ID}}},
			},
		},
	}

	hospital := newHospital(ctx, t, Config{}, pathways)
	defer hospital.Close()
	startPathway(t, hospital, testPathwayName)
	events, messages := hospital.ConsumeQueues(ctx, t)
	if got, want := events, 5; got != want {
		t.Errorf("StartPathway(%v) generated %v events, want %v", testPathwayName, got, want)
	}
	if got, want := len(messages), 3; got != want {
		t.Errorf("StartPathway(%v) generated %v messages, want %v", testPathwayName, got, want)
	}

	merge := messages[2]
	admit1 := messages[0]
	admit2 := messages[1]
	p1MRN := testhl7.MRN(t, admit1)
	p2MRN := testhl7.MRN(t, admit2)

	wantMainMRN := p2MRN
	wantMergeWithMRN := []string{p1MRN}
	if got, want := testhl7.MRN(t, merge), wantMainMRN; got != want {
		t.Errorf("StartPathway(%v) got merge message with MRN = %v, want %v", testPathwayName, got, want)
	}
	gotMergeWithMRN := testhl7.PriorPatientIdentifierList(t, merge)
	if diff := cmp.Diff(wantMergeWithMRN, gotMergeWithMRN); diff != "" {
		t.Errorf("StartPathway(%v) got PriorPatientIdentifierList with diff (-want, +got):\n%s", testPathwayName, diff)
	}
}

func TestRunPathwayUsePatientWithMRNs(t *testing.T) {
	ctx := context.Background()
	// An arbitrary pathway, it will create a new patient when it is run.
	pathway1 := pathway.Pathway{
		Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{Result: &pathway.Results{}},
		},
	}

	// Steps for a second pathway that starts using the first patient (UsePatient) at some point.
	steps := []pathway.Step{
		{Admission: &pathway.Admission{Loc: testLocAE}},
		{UsePatient: &pathway.UsePatient{Patient: "1"}},
		{Result: &pathway.Results{}},
		{Discharge: &pathway.Discharge{}},
	}

	// Test the UsePatient event in both History and Pathway.
	usePatientInPathway := pathway.Pathway{Pathway: steps}
	usePatientInHistory := pathway.Pathway{History: steps}

	for i, pathway2 := range []pathway.Pathway{usePatientInPathway, usePatientInHistory} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			pathways := map[string]pathway.Pathway{
				"pathway1": pathway1,
				"pathway2": pathway2,
			}

			// In order to use a patient from a previous pathway we must allow for that patient to be saved in the syncer.
			hospital := hospitalWithPatientSyncer(ctx, t, Config{}, pathways, teststate.NewItemSyncer())
			defer hospital.Close()
			startPathway(t, hospital, "pathway1", "pathway2")
			_, messages := hospital.ConsumeQueues(ctx, t)

			discharge := messages[len(messages)-1]
			if got, want := testhl7.MessageType(t, discharge), "ADT^A03"; got != want {
				t.Errorf("Last message got MessageType = %q, want %q", got, want)
			}
			// Since the second pathway switched to using the first patient, the discharge should be for
			// the first patient.
			if got, want := testhl7.MRN(t, discharge), "1"; got != want {
				t.Errorf("Last message got MRN = %v, want %v", got, want)
			}
		})
	}
}

func TestLoadPersonFromSyncer(t *testing.T) {
	ctx := context.Background()
	existingPatientMRN := "MRN-OF-EXISTING-PATIENT"
	expectedFirstName := "Expected first name"
	// Patient to be returned by syncer.
	existingPatient := state.Patient{
		PatientInfo: &ir.PatientInfo{
			Person: &ir.Person{
				FirstName: expectedFirstName,
				MRN:       existingPatientMRN,
			},
		},
		Orders: make(map[string]*ir.Order),
	}

	testSteps := []struct {
		n string
		p pathway.Pathway
	}{
		{"MRN in Persons",
			pathway.Pathway{
				Persons: &pathway.Persons{
					"main-patient": {
						MRN: existingPatientMRN,
					},
				},
				Pathway: []pathway.Step{
					{Result: &pathway.Results{}},
				},
			}},
		{"MRN in UsePatient",
			pathway.Pathway{
				Pathway: []pathway.Step{
					{UsePatient: &pathway.UsePatient{Patient: pathway.PatientID(existingPatientMRN)}},
					{Result: &pathway.Results{}},
				},
			},
		}}

	for i, testStep := range testSteps {
		t.Run(fmt.Sprintf("%d-%v", i, testStep.n), func(t *testing.T) {
			pathways := map[string]pathway.Pathway{testPathwayName: testStep.p}

			ps := teststate.NewItemSyncer()
			if err := ps.Write(existingPatient); err != nil {
				t.Fatalf("PatientSyncer.Write(%+v) failed with %v", existingPatient, err)
			}
			hospital := hospitalWithPatientSyncer(ctx, t, Config{}, pathways, ps)
			defer hospital.Close()
			startPathway(t, hospital, testPathwayName)
			_, messages := hospital.ConsumeQueues(ctx, t)
			if got, want := len(messages), 1; got != want {
				t.Fatalf("StartPathway(%v) generated %v messages, want %v", testPathwayName, got, want)
			}

			pid := testhl7.PID(t, messages[0])
			// Make sure the message had the first name returned by Syncer. Otherwise, it would have been
			// generated randomly.
			if got, want := pid.PatientName[0].GivenName.String(), expectedFirstName; got != want {
				t.Errorf("pid.PatientName[0].GivenName.String()=%v, want %v", got, want)
			}
		})
	}
}

func TestLoadExistingPatientUpdatePerson(t *testing.T) {
	ctx := context.Background()
	existingPatientMRN := "MRN-OF-EXISTING-PATIENT"
	existingFirstName := "Existing first name"
	existingSurname := "Existing surname"
	existingPatient := state.Patient{
		PatientInfo: &ir.PatientInfo{
			Person: &ir.Person{
				FirstName: existingFirstName,
				Surname:   existingSurname,
				MRN:       existingPatientMRN,
			},
		},
		Orders: make(map[string]*ir.Order),
	}

	// Create a pathway that uses the existing patient (via MRN), and make it override the surname.
	newSurname := "New surname"
	p := pathway.Pathway{
		Persons: &pathway.Persons{
			"main-patient": {
				MRN:     existingPatientMRN,
				Surname: pathway.OptionalRandomString(newSurname),
			},
		},
		Pathway: []pathway.Step{
			{Result: &pathway.Results{}},
		},
	}

	pathways := map[string]pathway.Pathway{testPathwayName: p}

	ps := teststate.NewItemSyncer()
	if err := ps.Write(existingPatient); err != nil {
		t.Fatalf("PatientSyncer.Write(%+v) failed with %v", existingPatient, err)
	}

	hospital := hospitalWithPatientSyncer(ctx, t, Config{}, pathways, ps)
	defer hospital.Close()
	startPathway(t, hospital, testPathwayName)
	_, messages := hospital.ConsumeQueues(ctx, t)
	if got, want := len(messages), 1; got != want {
		t.Fatalf("StartPathway(%v) generated %v messages, want %v", testPathwayName, got, want)
	}

	pid := testhl7.PID(t, messages[0])
	// The first name is the original one as this field was not set in the pathway.
	if got, want := pid.PatientName[0].GivenName.String(), existingFirstName; got != want {
		t.Errorf("pid.PatientName[0].GivenName.String()=%v, want %v", got, want)
	}
	if got, want := pid.PatientName[0].FamilyName.Surname.String(), newSurname; got != want {
		t.Errorf("pid.PatientName[0].FamilyName.Surname.String()=%v, want %v", got, want)
	}
}

func TestRunPathwayUsePatientWithPersonsAndMRNs(t *testing.T) {
	ctx := context.Background()
	// An arbitrary pathway, it will create a new patient when it is run.
	pathway1 := pathway.Pathway{
		Persons: &pathway.Persons{
			"main-patient": pathway.Person{
				FirstName: "Patient 1",
			},
		},
		Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
		},
	}

	// A pathway that refers to a patient created previously, via MRN.
	pathway2 := pathway.Pathway{
		Persons: &pathway.Persons{
			"patient-2": {FirstName: "Patient 2"},
		},
		Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{UsePatient: &pathway.UsePatient{Patient: "1"}},
			{Discharge: &pathway.Discharge{}},
		},
	}

	pathways := map[string]pathway.Pathway{
		"pathway1": pathway1,
		"pathway2": pathway2,
	}
	hospital := hospitalWithPatientSyncer(ctx, t, Config{}, pathways, teststate.NewItemSyncer())
	defer hospital.Close()
	startPathway(t, hospital, "pathway1", "pathway2")
	_, messages := hospital.ConsumeQueues(ctx, t)

	wantMessageTypes := []string{"ADT^A01", "ADT^A01", "ADT^A03"}
	gotMessageTypes := testhl7.Fields(t, messages, testhl7.MessageType)
	if diff := cmp.Diff(wantMessageTypes, gotMessageTypes); diff != "" {
		t.Errorf("StartPathway(%v) and StartPathway(%v) generated message types with diff (-want, +got):\n%s", "pathway1", "pathway2", diff)
	}
	wantFirstNames := []string{"Patient 1", "Patient 2", "Patient 1"}
	gotFirstNames := testhl7.Fields(t, messages, testhl7.FirstName)
	if diff := cmp.Diff(wantFirstNames, gotFirstNames); diff != "" {
		t.Errorf("StartPathway(%v) and StartPathway(%v) generated message with patient first names with diff (-want, +got):\n%s", "pathway1", "pathway2", diff)
	}
}

func TestStartPathway_InvalidPersonsSection(t *testing.T) {
	ctx := context.Background()
	mr := testmetrics.NewRetrieverFromGatherer(t)

	pathways := map[string]pathway.Pathway{
		testPathwayName: {
			// Note this pathway is Invalid as the Persons section always contains one item after parsing.
			Persons: &pathway.Persons{},
			Pathway: []pathway.Step{
				{Admission: &pathway.Admission{Loc: testLoc}},
				{Discharge: &pathway.Discharge{}},
			},
		},
	}

	hospital := newHospital(ctx, t, Config{}, pathways)
	defer hospital.Close()
	p, err := hospital.PathwayManager.GetPathway(testPathwayName)
	if err != nil {
		t.Fatalf("pathwayManager.GetPathway(%s) failed with %v", testPathwayName, err)
	}
	if _, err := hospital.StartPathway(p); err == nil {
		t.Errorf("StartPathway(%v) got nil error, want non nil error", testPathwayName)
	}
	_, messages := hospital.ConsumeQueues(ctx, t)
	if got, want := len(messages), 0; got != want {
		t.Errorf("StartPathway(%v) generated %v messages, want %v", testPathwayName, got, want)
	}

	metric := "simulated_hospital_errors_total"
	labels := map[string]string{
		"pathway_name": testPathwayName,
		"reason":       "invalid_persons_section",
	}
	initial, final := mr.GetCounterValues(t, metric, labels)
	if got, want := final-initial, 1.0; got != want {
		t.Errorf("Metric %s[%v] is incremented by %f (initial=%f, final=%f); want %f", metric, labels, got, initial, final, want)
	}
}

func TestRunPathwayDeleteVisitNoPastVisit(t *testing.T) {
	ctx := context.Background()
	pathways := map[string]pathway.Pathway{
		testPathwayName: {Pathway: []pathway.Step{
			{DeleteVisit: &pathway.DeleteVisit{}},
		}},
	}
	hospital := newHospital(ctx, t, Config{DeletePatientsFromMemory: true}, pathways)
	defer hospital.Close()
	startPathway(t, hospital, testPathwayName)
	events, messages := hospital.ConsumeQueues(ctx, t)

	// The second delete should fail because there are no visits to delete.
	if got, want := events, 1; got != want {
		t.Errorf("StartPathway(%v) generated %v events, want %v", testPathwayName, got, want)
	}
	if got, want := len(messages), 0; got != want {
		t.Errorf("StartPathway(%v) generated %v messages, want %v", testPathwayName, got, want)
	}

	if got, want := hospital.PatientsLen(), 0; got != want {
		t.Errorf("hospital.PatientsLen()=%v, want %v", got, want)
	}
}

func TestRunPathwayBedSwapStep(t *testing.T) {
	ctx := context.Background()
	// Both "CURRENT" and the explicit MRN should work.
	swapSteps := []*pathway.BedSwap{
		{Patient1: pathway.Current, Patient2: "1"},
		{Patient1: "2", Patient2: "1"},
	}

	for i, swapStep := range swapSteps {
		t.Run(fmt.Sprintf("%d-%v", i, swapStep.Patient1), func(t *testing.T) {
			delay := &pathway.Delay{From: time.Second, To: 5 * time.Second}
			longDelay := &pathway.Delay{From: time.Minute, To: 5 * time.Minute}
			veryLongDelay := &pathway.Delay{From: time.Hour, To: 5 * time.Hour}

			pathways := map[string]pathway.Pathway{
				"pathway1": {Pathway: []pathway.Step{
					{Admission: &pathway.Admission{Loc: testLocAE}},
					{Delay: delay},
					{Result: &pathway.Results{}},
					{Delay: longDelay},
					{Discharge: &pathway.Discharge{}},
				}},
				"pathway2": {Pathway: []pathway.Step{
					{Admission: &pathway.Admission{Loc: testLoc}},
					{Delay: delay},
					{BedSwap: swapStep},
					{Result: &pathway.Results{}},
					{Delay: veryLongDelay},
					{Discharge: &pathway.Discharge{}},
				}},
			}

			hospital := newHospital(ctx, t, Config{}, pathways)
			defer hospital.Close()
			startPathway(t, hospital, "pathway1", "pathway2")
			events, messages := hospital.ConsumeQueues(ctx, t)
			if got, want := events, 11; got != want {
				t.Errorf("StartPathway(%v) and StartPathway(%v) generated %v events, want %v", "pathway1", "pathway2", got, want)
			}
			if got, want := len(messages), 7; got != want {
				t.Errorf("StartPathway(%v) and StartPathway(%v) generated %v messages, want %v", "pathway1", "pathway2", got, want)
			}

			// Make sure the last messages have swapped the patient location.
			admit1 := messages[0]
			admit1POC := testhl7.PointOfCare(t, admit1)
			if diff := cmp.Diff(hospital.LocationManager.RoomManagers[testLocAE].Poc, admit1POC); diff != "" {
				t.Errorf("admit1.POC got diff (-want, +got):\n%s", diff)
			}
			admit2 := messages[1]
			admit2POC := testhl7.PointOfCare(t, admit2)
			if diff := cmp.Diff(hospital.LocationManager.RoomManagers[testLoc].Poc, admit2POC); diff != "" {
				t.Errorf("admit2.POC got diff (-want, +got):\n%s", diff)
			}
			// We've added delays to make sure the Discharges are the last two messages to be sent.
			discharge1 := messages[len(messages)-2]
			discharge1POC := testhl7.PointOfCare(t, discharge1)
			if diff := cmp.Diff(hospital.LocationManager.RoomManagers[testLoc].Poc, discharge1POC); diff != "" {
				t.Errorf("discharge1.POC got diff (-want, +got):\n%s", diff)
			}
			discharge2 := messages[len(messages)-1]
			discharge2POC := testhl7.PointOfCare(t, discharge2)
			if diff := cmp.Diff(hospital.LocationManager.RoomManagers[testLocAE].Poc, discharge2POC); diff != "" {
				t.Errorf("discharge2.POC got diff (-want, +got):\n%s", diff)
			}
		})
	}
}

func TestRunPathwayBedSwapStepPersonsSection(t *testing.T) {
	ctx := context.Background()
	p1ID := pathway.PatientID("patient-1")
	p2ID := pathway.PatientID("patient-2")

	pathways := map[string]pathway.Pathway{
		testPathwayName: {
			Persons: &pathway.Persons{
				p1ID: {FirstName: "Patient 1"},
				p2ID: {FirstName: "Patient 2"},
			},
			Pathway: []pathway.Step{
				{UsePatient: &pathway.UsePatient{Patient: p1ID}},
				{Admission: &pathway.Admission{Loc: testLocAE}},
				{UsePatient: &pathway.UsePatient{Patient: p2ID}},
				{Admission: &pathway.Admission{Loc: testLoc}},
				{BedSwap: &pathway.BedSwap{Patient1: p2ID, Patient2: p1ID}},
				{UsePatient: &pathway.UsePatient{Patient: p1ID}},
				{Discharge: &pathway.Discharge{}},
				{UsePatient: &pathway.UsePatient{Patient: p2ID}},
				{Discharge: &pathway.Discharge{}},
			},
		},
	}

	hospital := newHospital(ctx, t, Config{}, pathways)
	defer hospital.Close()
	startPathway(t, hospital, testPathwayName)
	events, messages := hospital.ConsumeQueues(ctx, t)
	if got, want := events, 9; got != want {
		t.Errorf("StartPathway(%v) generated %v events, want %v", testPathwayName, got, want)
	}
	if got, want := len(messages), 5; got != want {
		t.Errorf("StartPathway(%v) generated %v messages, want %v", testPathwayName, got, want)
	}

	// Make sure the last messages have swapped the patient location.
	admit1 := messages[0]
	admit1POC := testhl7.PointOfCare(t, admit1)
	if diff := cmp.Diff(hospital.LocationManager.RoomManagers[testLocAE].Poc, admit1POC); diff != "" {
		t.Errorf("admit1.POC got diff (-want, +got):\n%s", diff)
	}
	if got, want := testhl7.FirstName(t, admit1), "Patient 1"; got != want {
		t.Errorf("Got %v, want %v.", got, want)
	}
	admit2 := messages[1]
	admit2POC := testhl7.PointOfCare(t, admit2)
	if diff := cmp.Diff(hospital.LocationManager.RoomManagers[testLoc].Poc, admit2POC); diff != "" {
		t.Errorf("admit2.POC got diff (-want, +got):\n%s", diff)
	}
	if got, want := testhl7.FirstName(t, admit2), "Patient 2"; got != want {
		t.Errorf("Got %v, want %v.", got, want)
	}
	discharge1 := messages[len(messages)-2]
	discharge1POC := testhl7.PointOfCare(t, discharge1)
	if diff := cmp.Diff(hospital.LocationManager.RoomManagers[testLoc].Poc, discharge1POC); diff != "" {
		t.Errorf("discharge1.POC got diff (-want, +got):\n%s", diff)
	}
	if got, want := testhl7.FirstName(t, discharge1), "Patient 1"; got != want {
		t.Errorf("Got %v, want %v.", got, want)
	}
	discharge2 := messages[len(messages)-1]
	discharge2POC := testhl7.PointOfCare(t, discharge2)
	if diff := cmp.Diff(hospital.LocationManager.RoomManagers[testLocAE].Poc, discharge2POC); diff != "" {
		t.Errorf("discharge2.POC got diff (-want, +got):\n%s", diff)
	}
	if got, want := testhl7.FirstName(t, discharge2), "Patient 2"; got != want {
		t.Errorf("Got %v, want %v.", got, want)
	}
}

func TestRunPathwayBedSwapStepWrongMRN(t *testing.T) {
	ctx := context.Background()
	mr := testmetrics.NewRetrieverFromGatherer(t)

	pathways := map[string]pathway.Pathway{
		"pathway1": {Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{Result: &pathway.Results{}},
			{Discharge: &pathway.Discharge{}},
		}},
		"pathway2": {Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{BedSwap: &pathway.BedSwap{Patient1: "wrong-mrn", Patient2: "1"}},
			{Result: &pathway.Results{}},
			{Discharge: &pathway.Discharge{}},
		}},
	}

	hospital := newHospital(ctx, t, Config{}, pathways)
	defer hospital.Close()
	startPathway(t, hospital, "pathway1", "pathway2")
	events, messages := hospital.ConsumeQueues(ctx, t)
	// The last two events aren't run due to the failure in the Bed Swap event.
	if got, want := events, 5; got != want {
		t.Errorf("StartPathway(%v) and StartPathway(%v) generated %v events, want %v", "pathway1", "pathway2", got, want)
	}
	if got, want := len(messages), 4; got != want {
		t.Errorf("StartPathway(%v) and StartPathway(%v) generated %v messages, want %v", "pathway1", "pathway2", got, want)
	}

	metric := "simulated_hospital_errors_total"
	labels := map[string]string{
		"pathway_name": "pathway2",
		"reason":       "invalid bed swap state",
	}
	initial, final := mr.GetCounterValues(t, metric, labels)
	if got, want := final-initial, 1.0; got != want {
		t.Errorf("Metric %s[%v] is incremented by %f (initial=%f, final=%f); want %f", metric, labels, got, initial, final, want)
	}
}

func TestRunPathwayBedSwapPathwayBedSwapAfterPathwayFinished(t *testing.T) {
	ctx := context.Background()
	mr := testmetrics.NewRetrieverFromGatherer(t)

	delay := &pathway.Delay{From: 5 * time.Second, To: 5 * time.Second}
	longDelay := &pathway.Delay{From: 5 * time.Minute, To: 5 * time.Minute}
	veryLongDelay := &pathway.Delay{From: 5 * time.Hour, To: 5 * time.Hour}

	pathways := map[string]pathway.Pathway{
		"pathway1": {Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
			{Delay: delay},
			{Result: &pathway.Results{}},
			{Delay: longDelay},
			{Discharge: &pathway.Discharge{}},
		}},
		"pathway2": {Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{Delay: delay},
			{Result: &pathway.Results{}},
			{Delay: veryLongDelay},
			// After the long delay, pathway1 should have finished already.
			{BedSwap: &pathway.BedSwap{Patient1: pathway.Current, Patient2: "1"}},
			{Discharge: &pathway.Discharge{}},
		}},
	}

	hospital := hospitalWithPatientSyncer(ctx, t, Config{}, pathways, teststate.NewItemSyncer())
	defer hospital.Close()
	startPathway(t, hospital, "pathway1", "pathway2")
	events, messages := hospital.ConsumeQueues(ctx, t)

	// The BedSwap event fails to complete as one location is already nil, so the last event of
	// pathway2 doesn't run, and the last two messages aren't sent.
	if got, want := events, 10; got != want {
		t.Errorf("StartPathway(%v) and StartPathway(%v) generated %v events, want %v", "pathway1", "pathway2", got, want)
	}
	if got, want := len(messages), 5; got != want {
		t.Errorf("StartPathway(%v) and StartPathway(%v) generated %v messages, want %v", "pathway1", "pathway2", got, want)
	}

	metric := "simulated_hospital_pathway_duration_minutes"
	labels := map[string]string{"pathway_name": "pathway1"}
	beforeH, afterH := mr.GetHistogramBuckets(t, metric, labels)
	wantFinal := mr.IncrementBucketsCumulatively(beforeH, 10, 1)
	if diff := cmp.Diff(wantFinal, afterH, protocmp.Transform()); diff != "" {
		t.Errorf("Metric %s[%v] histogram buckets not incremented as expected (-want/+got):\n%s", metric, labels, diff)
	}

	metric = "simulated_hospital_errors_total"
	labels = map[string]string{
		"pathway_name": "pathway2",
		"reason":       "nil location in bed swap",
	}
	initial, final := mr.GetCounterValues(t, metric, labels)
	if got, want := final-initial, 1.0; got != want {
		t.Errorf("Metric %s[%v] is incremented by %f (initial=%f, final=%f); want %f", metric, labels, got, initial, final, want)
	}
}

func TestRunPathwayBedSwapPathwayBedSwapWithNilLocation(t *testing.T) {
	ctx := context.Background()
	mr := testmetrics.NewRetrieverFromGatherer(t)

	delay := &pathway.Delay{From: 5 * time.Second, To: 5 * time.Second}
	veryLongDelay := &pathway.Delay{From: 5 * time.Hour, To: 5 * time.Hour}

	// Have a long delay, so that we make sure the patient location is nil when the bed swap event
	// arrives. This tests a code path, but it is not a valid pathway.
	pathways := map[string]pathway.Pathway{
		"pathway1": {
			Pathway: []pathway.Step{
				{Delay: veryLongDelay},
				{Discharge: &pathway.Discharge{}},
			},
		},
		"pathway2": {
			Pathway: []pathway.Step{
				{Admission: &pathway.Admission{Loc: testLoc}},
				{Delay: delay},
				// At this point, patient 1 doesn't have a location yet.
				{BedSwap: &pathway.BedSwap{Patient1: pathway.Current, Patient2: "1"}},
			},
		},
	}

	hospital := newHospital(ctx, t, Config{}, pathways)
	defer hospital.Close()
	startPathway(t, hospital, "pathway1", "pathway2")
	events, messages := hospital.ConsumeQueues(ctx, t)
	// The BedSwap event fails to complete, so such message isn't sent.
	if got, want := events, 5; got != want {
		t.Errorf("StartPathway(%v) and StartPathway(%v) generated %v events, want %v", "pathway1", "pathway2", got, want)
	}
	if got, want := len(messages), 2; got != want {
		t.Errorf("StartPathway(%v) and StartPathway(%v) generated %v messages, want %v", "pathway1", "pathway2", got, want)
	}

	metric := "simulated_hospital_pathway_duration_minutes"
	labels := map[string]string{"pathway_name": "pathway1"}
	beforeH, afterH := mr.GetHistogramBuckets(t, metric, labels)
	wantAfter := mr.IncrementBucketsCumulatively(beforeH, 720, 1)
	if diff := cmp.Diff(wantAfter, afterH, protocmp.Transform()); diff != "" {
		t.Errorf("Metric %s[%v] histogram buckets not incremented as expected (-want/+got):\n%s", metric, labels, diff)
	}
	metric = "simulated_hospital_errors_total"
	labels = map[string]string{
		"pathway_name": "pathway2",
		"reason":       "nil location in bed swap",
	}
	initial, final := mr.GetCounterValues(t, metric, labels)
	if got, want := final-initial, 1.0; got != want {
		t.Errorf("Metric %s[%v] is incremented by %f (initial=%f, final=%f); want %f", metric, labels, got, initial, final, want)
	}
}

func TestRunPathwayPathwayFinishesSinglePatient(t *testing.T) {
	ctx := context.Background()
	pathways := map[string]pathway.Pathway{
		testPathwayName: {Pathway: []pathway.Step{
			{Result: &pathway.Results{}},
			{Result: &pathway.Results{}},
		}},
	}

	hospital := newHospital(ctx, t, Config{DeletePatientsFromMemory: true}, pathways)
	defer hospital.Close()
	startPathway(t, hospital, testPathwayName)
	// The patient is added to the list when the pathway starts,
	if got, want := hospital.PatientsLen(), 1; got != want {
		t.Errorf("hospital.PatientsLen()=%v, want %v", got, want)
	}

	// The patient should not be left in the map after all events run.
	hospital.ConsumeQueues(ctx, t)
	if got, want := hospital.PatientsLen(), 0; got != want {
		t.Errorf("hospital.PatientsLen()=%v, want %v", got, want)
	}
}

func TestRunPathwayPathwayFinishesNoPatients(t *testing.T) {
	ctx := context.Background()
	shortDelay := &pathway.Delay{From: time.Second, To: time.Second}
	longDelay := &pathway.Delay{From: 10 * time.Second, To: 10 * time.Second}

	pathways := map[string]pathway.Pathway{
		"pathway1": {Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{Delay: shortDelay},
			{Result: &pathway.Results{}},
			{Discharge: &pathway.Discharge{}},
		}},
		"pathway2": {Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{Delay: longDelay},
			{Discharge: &pathway.Discharge{}},
		}},
	}

	hospital := newHospital(ctx, t, Config{DeletePatientsFromMemory: true}, pathways)
	defer hospital.Close()
	startPathway(t, hospital, "pathway1", "pathway2")
	// The patients are added to the list when the pathways start.
	if got, want := hospital.PatientsLen(), 2; got != want {
		t.Errorf("hospital.PatientsLen()=%v, want %v", got, want)
	}

	for i := 0; i < 5; i++ {
		hospital.ConsumeQueuesWithLimit(ctx, t, 5, true)
		if got, want := hospital.PatientsLen(), 2; got != want {
			t.Errorf("hospital.PatientsLen()=%v, want %v", got, want)
		}
	}

	// All patients are deleted after the pathways finish.
	hospital.ConsumeQueues(ctx, t)
	if got, want := hospital.PatientsLen(), 0; got != want {
		t.Errorf("hospital.PatientsLen()=%v, want %v", got, want)
	}
}

func TestRunPathwaySameVisitDifferentPathways(t *testing.T) {
	ctx := context.Background()
	pathways := map[string]pathway.Pathway{
		testPathwayName: {Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
		}},
		"dischargePathway": {Pathway: []pathway.Step{
			{UsePatient: &pathway.UsePatient{Patient: "1"}},
			{Discharge: &pathway.Discharge{}},
		}},
	}

	hospital := hospitalWithPatientSyncer(ctx, t, Config{}, pathways, teststate.NewItemSyncer())
	defer hospital.Close()
	// Admission in one pathway.
	startPathway(t, hospital, testPathwayName)
	events, messages := hospital.ConsumeQueues(ctx, t)
	if got, want := events, 1; got != want {
		t.Errorf("StartPathway(%v) generated %v events, want %v", testPathwayName, got, want)
	}
	wantMessageTypes := []string{"ADT^A01"}
	gotMessageTypes := testhl7.Fields(t, messages, testhl7.MessageType)
	if diff := cmp.Diff(wantMessageTypes, gotMessageTypes); diff != "" {
		t.Errorf("StartPathway(%v) generated message types with diff (-want, +got):\n%s", testPathwayName, diff)
	}
	first := messages[0]
	firstMrn := testhl7.MRN(t, first)
	firstPV1 := testhl7.PV1(t, first)
	if got, want := len(firstPV1.DischargeDateTime), 0; got != want {
		t.Errorf("len(firstPV1.DischargeDateTime)=%v, want %v", got, want)
	}
	if got, want := hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds(), 1; got != want {
		t.Errorf("hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds()=%v, want %v", got, want)
	}

	// Discharge in another pathway.
	startPathway(t, hospital, "dischargePathway")
	secondEvents, secondMessages := hospital.ConsumeQueues(ctx, t)
	if got, want := secondEvents, 2; got != want {
		t.Errorf("secondEvents=%v, want %v", got, want)
	}
	allMessages := append(messages, secondMessages...)
	wantMessageTypes = []string{"ADT^A01", "ADT^A03"}
	gotMessageTypes = testhl7.Fields(t, allMessages, testhl7.MessageType)
	if diff := cmp.Diff(wantMessageTypes, gotMessageTypes); diff != "" {
		t.Errorf("StartPathway(%v) and StartPathway(%v) generated message types with diff (-want, +got):\n%s", testPathwayName, "dischargePathway", diff)
	}
	secondMessage := secondMessages[0]
	secondMrn := testhl7.MRN(t, secondMessage)
	if got, want := secondMrn, firstMrn; got != want {
		t.Errorf("secondMrn=%v, want %v", got, want)
	}
	wantMRNs := []string{firstMrn, firstMrn}
	gotMRNs := testhl7.Fields(t, allMessages, testhl7.MRN)
	if diff := cmp.Diff(wantMRNs, gotMRNs); diff != "" {
		t.Errorf("StartPathway(%v) and StartPathway(%v) generated patient MRNs with diff (-want, +got):\n%s", testPathwayName, "dischargePathway", diff)
	}
	secondPV1 := testhl7.PV1(t, secondMessage)
	if got, want := len(secondPV1.DischargeDateTime), 1; got != want {
		t.Errorf("len(secondPV1.DischargeDateTime)=%v, want %v", got, want)
	}
	if diff := cmp.Diff(firstPV1.VisitNumber, secondPV1.VisitNumber); diff != "" {
		t.Errorf("secondPV1.VisitNumber got diff (-want, +got):\n%s", diff)
	}
	if got, want := hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds(), 0; got != want {
		t.Errorf("hospital.LocationManager.RoomManagers[testLoc].OccupiedBeds()=%v, want %v", got, want)
	}
}

func TestPersistence(t *testing.T) {
	ctx := context.Background()
	pathways := map[string]pathway.Pathway{
		testPathwayName: {Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{Result: &pathway.Results{}},
			{Discharge: &pathway.Discharge{}},
		}},
	}

	ms := teststate.NewItemSyncerWithDelete(true)
	es := teststate.NewItemSyncerWithDelete(true)
	ps := teststate.NewItemSyncer()
	syncers := map[string]persist.ItemSyncer{
		state.MessageItemType: ms,
		state.EventItemType:   es,
		state.PatientItemType: ps,
	}
	hospital := newHospital(ctx, t, Config{AdditionalConfig: AdditionalConfig{ItemSyncers: syncers}}, pathways)
	defer hospital.Close()
	// When a pathway starts, an event and a patient are persisted.
	startPathway(t, hospital, testPathwayName)
	if got, want := ps.Count(), 1; got != want {
		t.Errorf("Patient Syncer Count()=%v, want %v", got, want)
	}
	if got, want := es.Count(), 1; got != want {
		t.Errorf("Event Syncer Count()=%v, want %v", got, want)
	}
	if got, want := es.CountDeletes(), 0; got != want {
		t.Errorf("Event Syncer CountDeletes()=%v, want %v", got, want)
	}
	if got, want := ms.Count(), 0; got != want {
		t.Errorf("Message Syncer Count()=%v, want %v", got, want)
	}
	if got, want := ms.CountDeletes(), 0; got != want {
		t.Errorf("Message Syncer CountDeletes()=%v, want %v", got, want)
	}

	// The event runs and it's removed from the syncer, but it generates a message and queues the next event.
	// The message is sent so it's removed from the syncer as well.
	hospital.ConsumeQueuesWithLimit(ctx, t, 1, true)
	if got, want := ps.Count(), 1; got != want {
		t.Errorf("Patient Syncer Count()=%v, want %v", got, want)
	}
	if got, want := es.Count(), 1; got != want {
		t.Errorf("Event Syncer Count()=%v, want %v", got, want)
	}
	if got, want := es.CountDeletes(), 1; got != want {
		t.Errorf("Event Syncer CountDeletes()=%v, want %v", got, want)
	}
	if got, want := ms.Count(), 0; got != want {
		t.Errorf("Message Syncer Count()=%v, want %v", got, want)
	}
	if got, want := ms.CountDeletes(), 1; got != want {
		t.Errorf("Message Syncer CountDeletes()=%v, want %v", got, want)
	}

	// Another event runs, it generates another message.
	hospital.ConsumeQueuesWithLimit(ctx, t, 1, true)
	if got, want := ps.Count(), 1; got != want {
		t.Errorf("Patient Syncer Count()=%v, want %v", got, want)
	}
	if got, want := es.Count(), 1; got != want {
		t.Errorf("Event Syncer Count()=%v, want %v", got, want)
	}
	if got, want := es.CountDeletes(), 2; got != want {
		t.Errorf("Event Syncer CountDeletes()=%v, want %v", got, want)
	}
	if got, want := ms.Count(), 0; got != want {
		t.Errorf("Message Syncer Count()=%v, want %v", got, want)
	}
	if got, want := ms.CountDeletes(), 2; got != want {
		t.Errorf("Message Syncer CountDeletes()=%v, want %v", got, want)
	}

	// When the last event runs, there are no more events pending.
	hospital.ConsumeQueuesWithLimit(ctx, t, -1, true)
	if got, want := ps.Count(), 1; got != want {
		t.Errorf("Patient Syncer Count()=%v, want %v", got, want)
	}
	if got, want := es.Count(), 0; got != want {
		t.Errorf("Event Syncer Count()=%v, want %v", got, want)
	}
	if got, want := es.CountDeletes(), 3; got != want {
		t.Errorf("Event Syncer CountDeletes()=%v, want %v", got, want)
	}
	if got, want := ms.Count(), 0; got != want {
		t.Errorf("Message Syncer Count()=%v, want %v", got, want)
	}
	if got, want := ms.CountDeletes(), 3; got != want {
		t.Errorf("Message Syncer CountDeletes()=%v, want %v", got, want)
	}
}

func TestPersistenceLoad(t *testing.T) {
	ctx := context.Background()
	// We test that we can load by first saving a few items.
	// We use a Hospital for convenience.
	// If there's nothing in the syncer, we don't load anything.
	pathways := map[string]pathway.Pathway{
		testPathwayName: {Pathway: []pathway.Step{
			{UpdatePerson: &pathway.UpdatePerson{}},
			{Admission: &pathway.Admission{Loc: testLoc}},
			{Result: &pathway.Results{}},
			{Discharge: &pathway.Discharge{}},
		}},
	}

	ms := teststate.NewItemSyncerWithDelete(true)
	es := teststate.NewItemSyncerWithDelete(true)
	ps := teststate.NewItemSyncerWithDelete(true)
	syncers := map[string]persist.ItemSyncer{
		state.MessageItemType: ms,
		state.EventItemType:   es,
		state.PatientItemType: ps,
	}
	hospital := newHospital(ctx, t, Config{AdditionalConfig: AdditionalConfig{ItemSyncers: syncers}}, pathways)
	defer hospital.Close()
	startPathway(t, hospital, testPathwayName)
	// Run the first event, so that the first message is sent.
	hospital.ConsumeQueuesWithLimit(ctx, t, 1, true)
	if got, want := ps.Count(), 1; got != want {
		t.Errorf("Patient Syncer Count()=%v, want %v", got, want)
	}
	if got, want := es.Count(), 1; got != want {
		t.Errorf("Event Syncer Count()=%v, want %v", got, want)
	}
	if got, want := ms.Count(), 0; got != want {
		t.Errorf("Message Syncer Count()=%v, want %v", got, want)
	}
	if got, want := ms.CountDeletes(), 1; got != want {
		t.Errorf("Message Syncer CountDeletes()=%v, want %v", got, want)
	}

	// Create a new instance of Simulated Hospital with the same syncers.
	hospital2 := newHospital(ctx, t, Config{AdditionalConfig: AdditionalConfig{ItemSyncers: syncers}}, pathways)
	defer hospital2.Close()
	// We expect hospital.patients to be empty since we do not load patients regardless of load flag.
	if got, want := hospital2.PatientsLen(), 0; got != want {
		t.Fatalf("hospital2.PatientsLen()=%v, want %v", got, want)
	}
	if got, want := hospital2.EventsLen(), 1; got != want {
		t.Fatalf("hospital2.EventsLen()=%v, want %v", got, want)
	}
	if got, want := hospital2.MessagesLen(), 0; got != want {
		t.Fatalf("hospital2.MessagesLen()=%v, want %v", got, want)
	}

	events, messages := hospital2.ConsumeQueues(ctx, t)
	// The first event was already run in the first Hospital.
	if got, want := events, 3; got != want {
		t.Fatalf("StartPathway(%v) generated %v events, want %v", testPathwayName, got, want)
	}
	if got, want := len(messages), 3; got != want {
		t.Fatalf("StartPathway(%v) generated %v messages, want %v", testPathwayName, got, want)
	}
}

type eventProcessor struct {
	match bool
	msgs  []*message.HL7Message
	err   error
}

func (p *eventProcessor) Matches(e *state.Event) bool {
	return p.match
}

func (p *eventProcessor) Process(*state.Event, *ir.PatientInfo, *processor.Config) ([]*message.HL7Message, error) {
	return p.msgs, p.err
}

type messageProcessor struct {
	match           bool
	err             error
	invocationOrder *[]int
	id              int
}

func (p *messageProcessor) Matches(*state.HL7Message) bool {
	return p.match
}

func (p *messageProcessor) Process(m *state.HL7Message) error {
	*p.invocationOrder = append(*(p.invocationOrder), p.id)
	return p.err
}

func TestRunPathwayWithEventProcessor(t *testing.T) {
	ctx := context.Background()
	pathways := map[string]pathway.Pathway{
		testPathwayName: {Pathway: []pathway.Step{
			{Result: &pathway.Results{}},
			{Discharge: &pathway.Discharge{}},
		}},
	}

	err := errors.New("this processor fails")
	resultID := "ORU^R01"
	dischargeID := "ADT^A03"
	proc1 := "processor_msg_one"
	proc2 := "processor_msg_two"
	proc3 := "processor_msg_three"
	proc1HL7 := []*message.HL7Message{{Message: proc1, Type: &message.Type{MessageType: "ADT", TriggerEvent: "A01"}}}
	proc2HL7 := []*message.HL7Message{{Message: proc2, Type: &message.Type{MessageType: "ORU", TriggerEvent: "R01"}}}
	proc3HL7 := []*message.HL7Message{{Message: proc3, Type: &message.Type{MessageType: "ORR", TriggerEvent: "O02"}}}
	twoHL7 := append(proc1HL7, proc2HL7...)

	tests := []struct {
		name               string
		preProcessors      []EventProcessor
		overrideProcessors []EventProcessor
		postProcessors     []EventProcessor
		wantEvents         int
		wantErrorStr       string
		wantMsgIDs         []string
	}{{
		name:       "No processors",
		wantMsgIDs: []string{resultID, dischargeID},
		wantEvents: 2,
	}, {
		name:               "No matching processors",
		preProcessors:      []EventProcessor{&eventProcessor{err: err}},
		overrideProcessors: []EventProcessor{&eventProcessor{err: err}},
		postProcessors:     []EventProcessor{&eventProcessor{err: err}},
		wantEvents:         2,
		wantMsgIDs:         []string{resultID, dischargeID},
	}, {
		name: "First pre processor matches, second doesn't",
		preProcessors: []EventProcessor{
			&eventProcessor{match: true, msgs: proc1HL7},
			&eventProcessor{err: err},
		},
		wantEvents: 2,
		wantMsgIDs: []string{proc1, resultID, proc1, dischargeID},
	}, {
		name: "First override processor matches, second doesn't",
		overrideProcessors: []EventProcessor{
			&eventProcessor{match: true, msgs: proc2HL7},
			&eventProcessor{err: err},
		},
		wantEvents: 2,
		wantMsgIDs: []string{proc2, proc2},
	}, {
		name: "First post processor matches, second doesn't",
		postProcessors: []EventProcessor{
			&eventProcessor{match: true, msgs: proc3HL7},
			&eventProcessor{err: err},
		},
		wantEvents: 2,
		wantMsgIDs: []string{resultID, proc3, dischargeID, proc3},
	}, {
		name: "First pre processor doesn't match, second does",
		preProcessors: []EventProcessor{
			&eventProcessor{err: err},
			&eventProcessor{match: true, msgs: proc1HL7},
		},
		wantEvents: 2,
		wantMsgIDs: []string{proc1, resultID, proc1, dischargeID},
	}, {
		name: "First override processor doesn't match, second does",
		overrideProcessors: []EventProcessor{
			&eventProcessor{err: err},
			&eventProcessor{match: true, msgs: proc2HL7},
		},
		wantEvents: 2,
		wantMsgIDs: []string{proc2, proc2},
	}, {
		name: "First post processor doesn't match, second does",
		postProcessors: []EventProcessor{
			&eventProcessor{err: err},
			&eventProcessor{match: true, msgs: proc3HL7},
		},
		wantEvents: 2,
		wantMsgIDs: []string{resultID, proc3, dischargeID, proc3},
	}, {
		name:           "Matching pre and post processors with standard processing logic",
		preProcessors:  []EventProcessor{&eventProcessor{match: true, msgs: proc1HL7}},
		postProcessors: []EventProcessor{&eventProcessor{match: true, msgs: proc3HL7}},
		wantEvents:     2,
		wantMsgIDs:     []string{proc1, resultID, proc3, proc1, dischargeID, proc3},
	}, {
		name:               "Matching pre, override and post processors",
		preProcessors:      []EventProcessor{&eventProcessor{match: true, msgs: proc1HL7}},
		overrideProcessors: []EventProcessor{&eventProcessor{match: true, msgs: proc2HL7}},
		postProcessors:     []EventProcessor{&eventProcessor{match: true, msgs: proc3HL7}},
		wantEvents:         2,
		wantMsgIDs:         []string{proc1, proc2, proc3, proc1, proc2, proc3},
	}, {
		name: "First pre processor succeeds, second fails",
		preProcessors: []EventProcessor{
			&eventProcessor{match: true, msgs: proc1HL7},
			&eventProcessor{match: true, err: err},
		},
		overrideProcessors: []EventProcessor{&eventProcessor{match: true, msgs: proc2HL7}},
		postProcessors:     []EventProcessor{&eventProcessor{match: true, msgs: proc3HL7}},
		wantEvents:         1,
		wantErrorStr:       "pre",
		wantMsgIDs:         []string{proc1},
	}, {
		name:          "First override processor succeeds, second fails",
		preProcessors: []EventProcessor{&eventProcessor{match: true, msgs: proc1HL7}},
		overrideProcessors: []EventProcessor{
			&eventProcessor{match: true, msgs: proc2HL7},
			&eventProcessor{match: true, err: err},
		},
		postProcessors: []EventProcessor{&eventProcessor{match: true, msgs: proc3HL7}},
		wantEvents:     1,
		wantErrorStr:   "override",
		wantMsgIDs:     []string{proc1, proc2},
	}, {
		name:               "First post processor succeeds, second fails",
		preProcessors:      []EventProcessor{&eventProcessor{match: true, msgs: proc1HL7}},
		overrideProcessors: []EventProcessor{&eventProcessor{match: true, msgs: proc2HL7}},
		postProcessors: []EventProcessor{
			&eventProcessor{match: true, msgs: proc3HL7},
			&eventProcessor{match: true, err: err},
		},
		wantEvents:   1,
		wantErrorStr: "post",
		wantMsgIDs:   []string{proc1, proc2, proc3},
	}, {
		name: "First pre processor fails, second succeeds",
		preProcessors: []EventProcessor{
			&eventProcessor{match: true, err: err},
			&eventProcessor{match: true, msgs: proc1HL7},
		},
		overrideProcessors: []EventProcessor{&eventProcessor{match: true, msgs: proc2HL7}},
		postProcessors:     []EventProcessor{&eventProcessor{match: true, msgs: proc3HL7}},
		wantEvents:         1,
		wantErrorStr:       "pre",
	}, {
		name:          "First override processor fails, second succeeds",
		preProcessors: []EventProcessor{&eventProcessor{match: true, msgs: proc1HL7}},
		overrideProcessors: []EventProcessor{
			&eventProcessor{match: true, err: err},
			&eventProcessor{match: true, msgs: proc2HL7},
		},
		postProcessors: []EventProcessor{&eventProcessor{match: true, msgs: proc3HL7}},
		wantEvents:     1,
		wantErrorStr:   "override",
		wantMsgIDs:     []string{proc1},
	}, {
		name:               "First post processor fails, second succeeds",
		preProcessors:      []EventProcessor{&eventProcessor{match: true, msgs: proc1HL7}},
		overrideProcessors: []EventProcessor{&eventProcessor{match: true, msgs: proc2HL7}},
		postProcessors: []EventProcessor{
			&eventProcessor{match: true, err: err},
			&eventProcessor{match: true, msgs: proc3HL7},
		},
		wantEvents:   1,
		wantErrorStr: "post",
		wantMsgIDs:   []string{proc1, proc2},
	}, {
		name: "Both pre processors succeed",
		preProcessors: []EventProcessor{
			&eventProcessor{match: true, msgs: proc1HL7},
			&eventProcessor{match: true, msgs: proc2HL7},
		},
		wantEvents: 2,
		wantMsgIDs: []string{proc1, proc2, resultID, proc1, proc2, dischargeID},
	}, {
		name: "Both override processors succeed",
		overrideProcessors: []EventProcessor{
			&eventProcessor{match: true, msgs: proc1HL7},
			&eventProcessor{match: true, msgs: proc2HL7},
		},
		wantEvents: 2,
		wantMsgIDs: []string{proc1, proc2, proc1, proc2},
	}, {
		name: "Both post processors succeed",
		postProcessors: []EventProcessor{
			&eventProcessor{match: true, msgs: proc1HL7},
			&eventProcessor{match: true, msgs: proc2HL7},
		},
		wantEvents: 2,
		wantMsgIDs: []string{resultID, proc1, proc2, dischargeID, proc1, proc2},
	}, {
		name: "Pre processor without messages",
		preProcessors: []EventProcessor{
			&eventProcessor{match: true, msgs: []*message.HL7Message{}},
		},
		wantEvents: 2,
		wantMsgIDs: []string{resultID, dischargeID},
	}, {
		name: "Override processor without messages",
		overrideProcessors: []EventProcessor{
			&eventProcessor{match: true, msgs: []*message.HL7Message{}},
		},
		wantEvents: 2,
	}, {
		name: "Post processor without messages",
		postProcessors: []EventProcessor{
			&eventProcessor{match: true, msgs: []*message.HL7Message{}},
		},
		wantEvents: 2,
		wantMsgIDs: []string{resultID, dischargeID},
	}, {
		name: "Multiple messages per pre processor",
		preProcessors: []EventProcessor{
			&eventProcessor{match: true, msgs: twoHL7},
		},
		wantEvents: 2,
		wantMsgIDs: []string{proc1, proc2, resultID, proc1, proc2, dischargeID},
	}, {
		name: "Multiple messages per override processor",
		overrideProcessors: []EventProcessor{
			&eventProcessor{match: true, msgs: twoHL7},
		},
		wantEvents: 2,
		wantMsgIDs: []string{proc1, proc2, proc1, proc2},
	}, {
		name: "Multiple messages per post processor",
		postProcessors: []EventProcessor{
			&eventProcessor{match: true, msgs: twoHL7},
		},
		wantEvents: 2,
		wantMsgIDs: []string{resultID, proc1, proc2, dischargeID, proc1, proc2},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mr := testmetrics.NewRetrieverFromGatherer(t)
			cfg := Config{
				AdditionalConfig: AdditionalConfig{
					Processors: Processors{EventPost: tc.postProcessors, EventOverride: tc.overrideProcessors, EventPre: tc.preProcessors},
				},
			}
			hospital := newHospital(ctx, t, cfg, pathways)
			defer hospital.Close()
			startPathway(t, hospital, testPathwayName)
			eventCount, gotMsgs := hospital.ConsumeQueues(ctx, t)

			if got, want := len(gotMsgs), len(tc.wantMsgIDs); got != want {
				t.Fatalf("StartPathway(%v) got %d messages; want %d", testPathwayName, got, want)
			}

			for i, msgID := range tc.wantMsgIDs {
				msg := gotMsgs[i]
				if !strings.Contains(msg, msgID) {
					t.Errorf("consumeQueues(t)'s message number %d was %q, want it to contain %q", i, msg, msgID)
				}
			}

			if got, want := eventCount, tc.wantEvents; got != want {
				t.Errorf("StartPathway(%v) got %d events; want %d", testPathwayName, got, want)
			}

			if tc.wantErrorStr == "" {
				return
			}
			metric := "simulated_hospital_errors_total"
			labels := map[string]string{
				"pathway_name": testPathwayName,
				"reason":       fmt.Sprintf("event_%s_processor", tc.wantErrorStr),
			}
			initial, final := mr.GetCounterValues(t, metric, labels)
			if got, want := final-initial, 1.0; got != want {
				t.Errorf("Metric %s[%v] is incremented by %f (initial=%f, final=%f); want %f", metric, labels, got, initial, final, want)
			}
		})
	}
}

func TestRunPathwayWithMessageProcessor(t *testing.T) {
	ctx := context.Background()
	pathways := map[string]pathway.Pathway{
		testPathwayName: {Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
		}},
	}

	errString := "this processor always fails"
	err := errors.New(errString)
	tests := []struct {
		name                string
		preProcessors       []MessageProcessor
		overrideProcessors  []MessageProcessor
		postProcessors      []MessageProcessor
		wantInvocationOrder []int
		wantMsgs            int
		wantErrorStr        string
	}{{
		name:     "No processors",
		wantMsgs: 1,
	}, {
		name:               "No matching processors",
		preProcessors:      []MessageProcessor{&messageProcessor{err: err, id: 1}},
		overrideProcessors: []MessageProcessor{&messageProcessor{err: err, id: 2}},
		postProcessors:     []MessageProcessor{&messageProcessor{err: err, id: 3}},
		wantMsgs:           1,
	}, {
		name: "First pre processor matches, second doesn't",
		preProcessors: []MessageProcessor{
			&messageProcessor{match: true, id: 1},
			&messageProcessor{err: err, id: 2},
		},
		wantMsgs:            1,
		wantInvocationOrder: []int{1},
	}, {
		name: "First override processor matches, second doesn't",
		overrideProcessors: []MessageProcessor{
			&messageProcessor{match: true, id: 1},
			&messageProcessor{err: err, id: 2},
		},
		wantInvocationOrder: []int{1},
	}, {
		name: "First post processor matches, second doesn't",
		postProcessors: []MessageProcessor{
			&messageProcessor{match: true, id: 1},
			&messageProcessor{err: err, id: 2},
		},
		wantMsgs:            1,
		wantInvocationOrder: []int{1},
	}, {
		name: "First pre processor doesn't match, second does",
		preProcessors: []MessageProcessor{
			&messageProcessor{err: err, id: 1},
			&messageProcessor{match: true, id: 2},
		},
		wantMsgs:            1,
		wantInvocationOrder: []int{2},
	}, {
		name: "First override processor doesn't match, second does",
		overrideProcessors: []MessageProcessor{
			&messageProcessor{err: err, id: 1},
			&messageProcessor{match: true, id: 2},
		},
		wantInvocationOrder: []int{2},
	}, {
		name: "First post processor doesn't match, second does",
		postProcessors: []MessageProcessor{
			&messageProcessor{err: err, id: 1},
			&messageProcessor{match: true, id: 2},
		},
		wantMsgs:            1,
		wantInvocationOrder: []int{2},
	}, {
		name:                "Matching pre and post processors with standard processing logic",
		preProcessors:       []MessageProcessor{&messageProcessor{match: true, id: 1}},
		postProcessors:      []MessageProcessor{&messageProcessor{match: true, id: 2}},
		wantInvocationOrder: []int{1, 2},
		wantMsgs:            1,
	}, {
		name:                "Matching pre, override and post processors",
		preProcessors:       []MessageProcessor{&messageProcessor{match: true, id: 1}},
		overrideProcessors:  []MessageProcessor{&messageProcessor{match: true, id: 2}},
		postProcessors:      []MessageProcessor{&messageProcessor{match: true, id: 3}},
		wantInvocationOrder: []int{1, 2, 3},
	}, {
		name: "First pre processor succeeds, second fails",
		preProcessors: []MessageProcessor{
			&messageProcessor{match: true, id: 1},
			&messageProcessor{match: true, err: err, id: 2},
		},
		overrideProcessors:  []MessageProcessor{&messageProcessor{match: true, id: 3}},
		postProcessors:      []MessageProcessor{&messageProcessor{match: true, id: 4}},
		wantInvocationOrder: []int{1, 2},
		wantErrorStr:        "pre",
	}, {
		name:          "First override processor succeeds, second fails",
		preProcessors: []MessageProcessor{&messageProcessor{match: true, id: 1}},
		overrideProcessors: []MessageProcessor{
			&messageProcessor{match: true, id: 2},
			&messageProcessor{match: true, err: err, id: 3},
		},
		postProcessors:      []MessageProcessor{&messageProcessor{match: true, id: 4}},
		wantInvocationOrder: []int{1, 2, 3},
		wantErrorStr:        "override",
	}, {
		name:               "First post processor succeeds, second fails",
		preProcessors:      []MessageProcessor{&messageProcessor{match: true, id: 1}},
		overrideProcessors: []MessageProcessor{&messageProcessor{match: true, id: 2}},
		postProcessors: []MessageProcessor{
			&messageProcessor{match: true, id: 3},
			&messageProcessor{match: true, err: err, id: 4},
		},
		wantErrorStr:        "post",
		wantInvocationOrder: []int{1, 2, 3, 4},
	}, {
		name: "First pre processor fails, second succeeds",
		preProcessors: []MessageProcessor{
			&messageProcessor{match: true, err: err, id: 1},
			&messageProcessor{match: true, id: 2},
		},
		overrideProcessors:  []MessageProcessor{&messageProcessor{match: true, id: 3}},
		postProcessors:      []MessageProcessor{&messageProcessor{match: true, id: 4}},
		wantInvocationOrder: []int{1},
		wantErrorStr:        "pre",
	}, {
		name:          "First override processor fails, second succeeds",
		preProcessors: []MessageProcessor{&messageProcessor{match: true, id: 1}},
		overrideProcessors: []MessageProcessor{
			&messageProcessor{match: true, err: err, id: 2},
			&messageProcessor{match: true, id: 3},
		},
		postProcessors:      []MessageProcessor{&messageProcessor{match: true, id: 4}},
		wantInvocationOrder: []int{1, 2},
		wantErrorStr:        "override",
	}, {
		name:               "First post processor fails, second succeeds",
		preProcessors:      []MessageProcessor{&messageProcessor{match: true, id: 1}},
		overrideProcessors: []MessageProcessor{&messageProcessor{match: true, id: 2}},
		postProcessors: []MessageProcessor{
			&messageProcessor{match: true, err: err, id: 3},
			&messageProcessor{match: true, id: 4},
		},
		wantInvocationOrder: []int{1, 2, 3},
		wantErrorStr:        "post",
	}, {
		name: "Both pre processors succeed",
		preProcessors: []MessageProcessor{
			&messageProcessor{match: true, id: 1},
			&messageProcessor{match: true, id: 2},
		},
		wantMsgs:            1,
		wantInvocationOrder: []int{1, 2},
	}, {
		name: "Both override processors succeed",
		overrideProcessors: []MessageProcessor{
			&messageProcessor{match: true, id: 1},
			&messageProcessor{match: true, id: 2},
		},
		wantInvocationOrder: []int{1, 2},
	}, {
		name: "Both post processors succeed",
		postProcessors: []MessageProcessor{
			&messageProcessor{match: true, id: 1},
			&messageProcessor{match: true, id: 2},
		},
		wantInvocationOrder: []int{1, 2},
		wantMsgs:            1,
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mr := testmetrics.NewRetrieverFromGatherer(t)

			var gotInvocationOrder []int
			for _, p := range tc.preProcessors {
				mp := p.(*messageProcessor)
				mp.invocationOrder = &gotInvocationOrder
			}
			for _, p := range tc.overrideProcessors {
				mp := p.(*messageProcessor)
				mp.invocationOrder = &gotInvocationOrder
			}
			for _, p := range tc.postProcessors {
				mp := p.(*messageProcessor)
				mp.invocationOrder = &gotInvocationOrder
			}

			cfg := Config{
				AdditionalConfig: AdditionalConfig{
					Processors: Processors{MessagePost: tc.postProcessors, MessageOverride: tc.overrideProcessors, MessagePre: tc.preProcessors},
				},
			}
			hospital := newHospital(ctx, t, cfg, pathways)
			defer hospital.Close()

			p, _ := hospital.PathwayManager.GetPathway(testPathwayName)
			if _, err := hospital.StartPathway(p); err != nil {
				t.Fatalf("startPathway(%v) failed with %v", testPathwayName, err)
			}

			// Allow failures - that's exactly what we want to test.
			_, gotMsgs := hospital.ConsumeQueuesWithLimit(ctx, t, -1, false)

			if got, want := len(gotMsgs), tc.wantMsgs; got != want {
				t.Errorf("StartPathway(%v) got %d messages; want %d", testPathwayName, got, want)
			}

			if diff := cmp.Diff(tc.wantInvocationOrder, gotInvocationOrder); diff != "" {
				t.Errorf("Processor invocation order got diff (-want, +got)\n%s", diff)
			}

			if tc.wantErrorStr == "" {
				return
			}
			metric := "simulated_hospital_errors_total"
			labels := map[string]string{
				"pathway_name": testPathwayName,
				"reason":       fmt.Sprintf("message_%s_processor", tc.wantErrorStr),
			}
			initial, final := mr.GetCounterValues(t, metric, labels)
			if got, want := final-initial, 1.0; got != want {
				t.Errorf("Metric %s[%v] is incremented by %f (initial=%f, final=%f); want %f", metric, labels, got, initial, final, want)
			}
		})
	}
}

func TestGenericEventAndAdditionalData(t *testing.T) {
	ctx := context.Background()
	pathways := map[string]pathway.Pathway{
		testPathwayName: {Pathway: []pathway.Step{
			{
				Generic:    &pathway.Generic{Name: "add_medication"},
				Parameters: &pathway.Parameters{Custom: map[string]string{"medication_name": "paracetamol"}},
			},
			{
				Generic:    &pathway.Generic{Name: "add_medication"},
				Parameters: &pathway.Parameters{Custom: map[string]string{"medication_name": "ibuprofen"}},
			},
			{
				Generic:    &pathway.Generic{Name: "assert_medications"},
				Parameters: &pathway.Parameters{Custom: map[string]string{"medications_number": "2"}},
			},
		}},
	}

	assertMedicationsProc := &assertMedicationsProc{}

	tests := []struct {
		name      string
		ac        AdditionalConfig
		wantError float64
	}{{
		name:      "no_override",
		wantError: 1,
	}, {
		name: "override",
		ac:   AdditionalConfig{Processors: Processors{EventOverride: []EventProcessor{&addMedicationProc{}, assertMedicationsProc}}},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assertMedicationsProc.t = t
			mr := testmetrics.NewRetrieverFromGatherer(t)
			cfg := Config{AdditionalConfig: tc.ac}
			hospital := newHospital(ctx, t, cfg, pathways)
			defer hospital.Close()

			p, _ := hospital.PathwayManager.GetPathway(testPathwayName)
			if _, err := hospital.StartPathway(p); err != nil {
				t.Fatalf("startPathway(%v) failed with %v", testPathwayName, err)
			}

			hospital.ConsumeQueues(ctx, t)

			metric := "simulated_hospital_errors_total"
			labels := map[string]string{
				"pathway_name": testPathwayName,
				"reason":       "missing_processor_of_generic_event",
			}
			initial, final := mr.GetCounterValues(t, metric, labels)
			if got, want := final-initial, tc.wantError; got != want {
				t.Errorf("Metric %s[%v] is incremented by %f (initial=%f, final=%f); want %f", metric, labels, got, initial, final, want)
			}
		})
	}
}

// addMedicationProc is a custom event processor that adds medications to a patient's medical record.
type addMedicationProc struct{}

// Matches returns whether the given event can be processed by our custom processor.
func (p *addMedicationProc) Matches(e *state.Event) bool {
	return e.Step.Generic != nil && e.Step.Generic.Name == "add_medication"
}

// AdditionalData is the type of *ir.PatientInfo.AdditionalData.
// Medications aren't part of the regular fields of patientInfo so we use the AdditionalData field.
type AdditionalData struct {
	Medications []string
}

// Process adds a medication to the patient's medical record.
func (p *addMedicationProc) Process(e *state.Event, patientInfo *ir.PatientInfo, _ *processor.Config) ([]*message.HL7Message, error) {
	newMedication := e.Step.Parameters.Custom["medication_name"]
	var ad AdditionalData
	if patientInfo.AdditionalData != nil {
		ad = patientInfo.AdditionalData.(AdditionalData)
	}

	ad.Medications = append(ad.Medications, newMedication)
	patientInfo.AdditionalData = ad
	return nil, nil
}

// assertMedicationsProc is a custom event processor that asserts that there's a certain number of medications in the patient's medical record.
type assertMedicationsProc struct {
	t *testing.T
}

// Matches returns whether the given event can be processed by our custom processor.
func (p *assertMedicationsProc) Matches(e *state.Event) bool {
	return e.Step.Generic != nil && e.Step.Generic.Name == "assert_medications"
}

// Process makes sure that the number of medications in the patient's medical record is expected.
func (p *assertMedicationsProc) Process(e *state.Event, patientInfo *ir.PatientInfo, _ *processor.Config) ([]*message.HL7Message, error) {
	ad := patientInfo.AdditionalData.(AdditionalData)
	want, err := strconv.Atoi(e.Step.Parameters.Custom["medications_number"])
	if err != nil {
		p.t.Fatalf("strconv.Atoi(%s) failed with %v", e.Step.Parameters.Custom["medications_number"], err)
	}
	if got := len(ad.Medications); got != want {
		p.t.Errorf("printMedicationsProc.Process() found %d medications, want %d", got, want)
	}
	return nil, nil
}

// testPathwayManager implements pathway.Manager interface.
// testPathwayManager always returns the same pathway consisting of one Admission and one Discharge.
type testPathwayManager struct{}

func (tpm *testPathwayManager) GetPathway(pathwayName string) (*pathway.Pathway, error) {
	return tpm.NextPathway()
}

func (tpm *testPathwayManager) NextPathway() (*pathway.Pathway, error) {
	p := &pathway.Pathway{Pathway: []pathway.Step{
		{Admission: &pathway.Admission{Loc: testLocAE}},
		{Discharge: &pathway.Discharge{}},
	}}
	p.Init(testPathwayName)
	return p, nil
}

func TestRunPathwayCustomPathwayManager(t *testing.T) {
	ctx := context.Background()
	pm := &testPathwayManager{}
	cfg := Config{
		PathwayManager:  pm,
		LocationManager: testlocation.NewLocationManager(ctx, t, testLoc, testLocAE),
		DataFiles:       test.DataFiles[test.Test],
	}
	hospital := testhospital.WithTime(ctx, t, testhospital.Config{Config: cfg, Arguments: testhospital.Arguments}, now)
	defer hospital.Close()
	p, err := pm.NextPathway()
	if err != nil {
		t.Fatalf("pathwayManager.NextPathway() failed with %v", err)
	}
	startPathway(t, hospital, p.Name())
	_, messages := hospital.ConsumeQueues(ctx, t)

	wantMessageTypes := []string{"ADT^A01", "ADT^A03"}
	gotMessageTypes := testhl7.Fields(t, messages, testhl7.MessageType)
	if diff := cmp.Diff(wantMessageTypes, gotMessageTypes); diff != "" {
		t.Errorf("StartPathway(%v) generated message types with diff (-want, +got):\n%s", p.Name(), diff)
	}
}

func TestRunPathwayWithHardcodedMessage(t *testing.T) {
	ctx := context.Background()
	wantMessageTypes := []string{"ADT^A01", "ADT^A03"}
	pathways := map[string]pathway.Pathway{
		testPathwayName: {Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{HardcodedMessage: &pathway.HardcodedMessage{Regex: "DischargeHardcodedMessage"}},
		}},
	}

	seed := "hardcoded_messages.yml"
	dir := testwrite.BytesToDir(t, []byte(hardcodedMessageYml), seed)
	msgControlGen := &header.MessageControlGenerator{}
	hardcodedMessagesManager, err := hardcoded.NewManager(ctx, dir, msgControlGen)
	if err != nil {
		t.Fatalf("NewManager(%s) failed with %v", hardcodedMessageYml, err)
	}

	hospital := newHospital(ctx, t, Config{MessagesManager: hardcodedMessagesManager, MessageControlGenerator: msgControlGen}, pathways)
	defer hospital.Close()
	startPathway(t, hospital, testPathwayName)
	_, messages := hospital.ConsumeQueues(ctx, t)
	if got, want := len(messages), len(wantMessageTypes); got != want {
		t.Fatalf("len(messages) = %d, want %d", got, want)
	}

	gotMessageTypes := testhl7.Fields(t, messages, testhl7.MessageType)
	if diff := cmp.Diff(wantMessageTypes, gotMessageTypes); diff != "" {
		t.Errorf("StartPathway(%v) generated message types with diff (-want, +got):\n%s", testPathwayName, diff)
	}

	pid0 := testhl7.PID(t, messages[0])
	pid1 := testhl7.PID(t, messages[1])

	if diff := cmp.Diff(pid0, pid1); diff != "" {
		t.Errorf("PID comparison got diff (-want, +got):\n%s", diff)
	}
}

func newHospital(ctx context.Context, t *testing.T, cfg Config, pathways map[string]pathway.Pathway) *testhospital.Hospital {
	t.Helper()
	return hospitalWithTime(ctx, t, cfg, pathways, now)
}

func hospitalWithPatientSyncer(ctx context.Context, t *testing.T, cfg Config, pathways map[string]pathway.Pathway, ps persist.ItemSyncer) *testhospital.Hospital {
	t.Helper()
	cfg.AdditionalConfig.ItemSyncers = map[string]persist.ItemSyncer{state.PatientItemType: ps}
	return newHospital(ctx, t, cfg, pathways)
}

func hospitalWithTime(ctx context.Context, t *testing.T, cfg Config, pathways map[string]pathway.Pathway, now time.Time) *testhospital.Hospital {
	t.Helper()
	pm, err := pathway.NewDistributionManager(pathways, nil, nil)
	if err != nil {
		t.Fatalf("pathway.NewDistributionManager(%v,%v,%v) failed with %v", pathways, nil, nil, err)
	}
	cfg.PathwayManager = pm
	cfg.LocationManager = testlocation.NewLocationManager(ctx, t, testLoc, testLocAE)
	return testhospital.WithTime(ctx, t, testhospital.Config{Config: cfg, Arguments: testhospital.Arguments}, now)
}

// startPathway starts pathways in the hospital by queuing their first events.
func startPathway(t *testing.T, h *testhospital.Hospital, pathwayNames ...string) {
	t.Helper()
	for _, pathwayName := range pathwayNames {
		p, err := h.PathwayManager.GetPathway(pathwayName)
		if err != nil {
			t.Fatalf("pathwayManager.GetPathway(%s) failed with %v", pathwayName, err)
		}
		if _, err := h.StartPathway(p); err != nil {
			t.Fatalf("StartPathway(%v) failed with %v", pathwayName, err)
		}
	}
}

func toPatientIDs(strings ...string) []pathway.PatientID {
	ids := make([]pathway.PatientID, len(strings))
	for i, s := range strings {
		ids[i] = pathway.PatientID(s)
	}
	return ids
}

func TestRunPathwayResultsOnSameOrDifferentOrders(t *testing.T) {
	ctx := context.Background()
	now := time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC)
	timeFromNow := -30 * time.Minute
	orderProfile := "US Pelvis TA and transvaginal"
	result := &pathway.Result{
		TestName: "UPELD",
		Value:    "The left ovary is normal. The right ovary is enlarged.",
	}

	for _, tt := range []struct {
		name             string
		pathway          pathway.Pathway
		wantSetID        [][]string
		wantResultStatus [][]string
	}{{
		name: "same order ID",
		pathway: pathway.Pathway{
			History: []pathway.Step{{
				Result: &pathway.Results{
					OrderID:      "uspelvis_transa_transv1",
					OrderProfile: orderProfile,
					Results:      []*pathway.Result{result, result},
				},
				Parameters: &pathway.Parameters{TimeFromNow: &timeFromNow},
			}},
			Pathway: []pathway.Step{{
				Result: &pathway.Results{
					OrderID:      "uspelvis_transa_transv1",
					OrderProfile: orderProfile,
					Results:      []*pathway.Result{result, result},
				},
			}}},
		wantResultStatus: [][]string{{"F", "F"}, {"C", "C"}},
		wantSetID:        [][]string{{"1", "2"}, {"3", "4"}},
	}, {
		name: "same order ID and override status in first message",
		pathway: pathway.Pathway{
			History: []pathway.Step{{
				Result: &pathway.Results{
					OrderID:      "uspelvis_transa_transv1",
					OrderProfile: orderProfile,
					Results:      []*pathway.Result{result, result},
					OrderStatus:  "P",
					ResultStatus: "P",
				},
				Parameters: &pathway.Parameters{TimeFromNow: &timeFromNow},
			}},
			Pathway: []pathway.Step{{
				Result: &pathway.Results{
					OrderID:      "uspelvis_transa_transv1",
					OrderProfile: orderProfile,
					Results:      []*pathway.Result{result, result},
				},
			}}},
		wantResultStatus: [][]string{{"P", "P"}, {"F", "F"}},
		wantSetID:        [][]string{{"1", "2"}, {"3", "4"}},
	}, {
		name: "same order ID and override status in both messages",
		pathway: pathway.Pathway{
			History: []pathway.Step{{
				Result: &pathway.Results{
					OrderID:      "uspelvis_transa_transv1",
					OrderProfile: orderProfile,
					Results:      []*pathway.Result{result, result},
					OrderStatus:  "P",
					ResultStatus: "P",
				},
				Parameters: &pathway.Parameters{TimeFromNow: &timeFromNow},
			}},
			Pathway: []pathway.Step{{
				Result: &pathway.Results{
					OrderID:      "uspelvis_transa_transv1",
					OrderProfile: orderProfile,
					Results:      []*pathway.Result{result, result},
					OrderStatus:  "C",
					ResultStatus: "C",
				},
			}}},
		wantResultStatus: [][]string{{"P", "P"}, {"C", "C"}},
		wantSetID:        [][]string{{"1", "2"}, {"3", "4"}},
	}, {
		name: "different order ID and override status in first message",
		pathway: pathway.Pathway{
			History: []pathway.Step{{
				Result: &pathway.Results{
					OrderID:      "uspelvis_transa_transv1",
					OrderProfile: orderProfile,
					Results:      []*pathway.Result{result, result},
				},
				Parameters: &pathway.Parameters{TimeFromNow: &timeFromNow},
			}},
			Pathway: []pathway.Step{{
				Result: &pathway.Results{
					OrderID:      "another_uspelvis_transa_transv",
					OrderProfile: orderProfile,
					Results:      []*pathway.Result{result, result},
				},
			}}},
		wantResultStatus: [][]string{{"F", "F"}, {"F", "F"}},
		wantSetID:        [][]string{{"1", "2"}, {"1", "2"}},
	}, {
		name: "expect correction",
		pathway: pathway.Pathway{
			History: []pathway.Step{{
				Result: &pathway.Results{
					OrderID:          "uspelvis_transa_transv1",
					OrderProfile:     orderProfile,
					Results:          []*pathway.Result{result, result},
					ExpectCorrection: true,
					OrderStatus:      "P",
					ResultStatus:     "P",
				},
				Parameters: &pathway.Parameters{TimeFromNow: &timeFromNow},
			}},
			Pathway: []pathway.Step{{
				Result: &pathway.Results{
					OrderID:      "uspelvis_transa_transv1",
					OrderProfile: orderProfile,
					Results:      []*pathway.Result{result, result},
					OrderStatus:  "C",
					ResultStatus: "C",
				},
			}, {
				Result: &pathway.Results{
					OrderID:      "uspelvis_transa_transv1",
					OrderProfile: orderProfile,
					Results:      []*pathway.Result{result, result},
					OrderStatus:  "F",
					ResultStatus: "F",
				},
			}}},
		wantResultStatus: [][]string{{"P", "P"}, {"C", "C"}, {"F", "F"}},
		wantSetID:        [][]string{{"1", "2"}, {"1", "2"}, {"3", "4"}},
	}, {
		name: "different order ID in between",
		pathway: pathway.Pathway{
			History: []pathway.Step{{
				Result: &pathway.Results{
					OrderID:      "uspelvis_transa_transv1",
					OrderProfile: orderProfile,
					Results:      []*pathway.Result{result, result},
					OrderStatus:  "P",
					ResultStatus: "P",
				},
				Parameters: &pathway.Parameters{TimeFromNow: &timeFromNow},
			}},
			Pathway: []pathway.Step{{
				Result: &pathway.Results{
					OrderID:      "another_uspelvis_transa_transv",
					OrderProfile: orderProfile,
					Results:      []*pathway.Result{result, result},
				},
			}, {
				Result: &pathway.Results{
					OrderID:      "uspelvis_transa_transv1",
					OrderProfile: orderProfile,
					Results:      []*pathway.Result{result, result},
					OrderStatus:  "C",
					ResultStatus: "C",
				},
			}}},
		wantResultStatus: [][]string{{"P", "P"}, {"F", "F"}, {"C", "C"}},
		wantSetID:        [][]string{{"1", "2"}, {"1", "2"}, {"3", "4"}},
	}, {
		name: "expect correction with different order ID in between",
		pathway: pathway.Pathway{
			History: []pathway.Step{{
				Result: &pathway.Results{
					OrderID:          "uspelvis_transa_transv1",
					OrderProfile:     orderProfile,
					Results:          []*pathway.Result{result, result},
					ExpectCorrection: true,
					OrderStatus:      "P",
					ResultStatus:     "P",
				},
				Parameters: &pathway.Parameters{TimeFromNow: &timeFromNow},
			}},
			Pathway: []pathway.Step{{
				Result: &pathway.Results{
					OrderID:      "another_uspelvis_transa_transv",
					OrderProfile: orderProfile,
					Results:      []*pathway.Result{result, result},
				},
			}, {
				Result: &pathway.Results{
					OrderID:      "uspelvis_transa_transv",
					OrderProfile: orderProfile,
					Results:      []*pathway.Result{result, result},
					OrderStatus:  "C",
					ResultStatus: "C",
				},
			}}},
		wantResultStatus: [][]string{{"P", "P"}, {"F", "F"}, {"C", "C"}},
		wantSetID:        [][]string{{"1", "2"}, {"1", "2"}, {"1", "2"}},
	}} {
		t.Run(tt.name, func(t *testing.T) {
			pathways := map[string]pathway.Pathway{
				testPathwayName: tt.pathway,
			}

			hospital := hospitalWithTime(ctx, t, Config{}, pathways, now)
			defer hospital.Close()
			startPathway(t, hospital, testPathwayName)
			_, messages := hospital.ConsumeQueues(ctx, t)
			if got, want := len(messages), len(tt.wantResultStatus); got != want {
				t.Errorf("StartPathway(%v) generated %v messages, want %v", testPathwayName, got, want)
			}

			gotSetID := testhl7.OBXFields(t, messages, testhl7.OBXSetID)
			if diff := cmp.Diff(tt.wantSetID, gotSetID); diff != "" {
				t.Errorf("StartPathway(%v) got OBX SetID diff (-want, +got):\n%s", testPathwayName, diff)
			}

			gotResultStatus := testhl7.OBXFields(t, messages, testhl7.OBXResultStatus)
			if diff := cmp.Diff(tt.wantResultStatus, gotResultStatus); diff != "" {
				t.Errorf("StartPathway(%v) got OBX ResultStatus diff (-want, +got):\n%s", testPathwayName, diff)
			}
		})
	}
}

func TestPathwayClinicalNote(t *testing.T) {
	ctx := context.Background()
	for _, tt := range []struct {
		name                  string
		steps                 []*pathway.ClinicalNote
		wantFillerOrderNumber []*hl7.CWE
		wantObservationValues [][]string
	}{{
		name: "single ClinicalNote step",
		steps: []*pathway.ClinicalNote{{
			ContentType:     "txt",
			DocumentType:    "type",
			DocumentID:      "id",
			DocumentContent: "content",
		}},
		wantFillerOrderNumber: []*hl7.CWE{{Identifier: hl7.NewST("type"), Text: hl7.NewST("type")}},
		wantObservationValues: [][]string{{"^^txt^^content"}},
	}, {
		name: "two ClinicalNote step",
		steps: []*pathway.ClinicalNote{{
			ContentType:     "txt",
			DocumentTitle:   "title",
			DocumentID:      "new-id",
			DocumentType:    "new-type",
			DocumentContent: "new-content",
		}, {
			ContentType:     "txt",
			DocumentTitle:   "new-title",
			DocumentID:      "second-id",
			DocumentType:    "second-type",
			DocumentContent: "second-content",
		}},
		wantFillerOrderNumber: []*hl7.CWE{
			{Identifier: hl7.NewST("new-type"), Text: hl7.NewST("new-type"), AlternateText: hl7.NewST("title")},
			{Identifier: hl7.NewST("second-type"), Text: hl7.NewST("second-type"), AlternateText: hl7.NewST("new-title")},
		},
		wantObservationValues: [][]string{{"^^txt^^new-content"}, {"^^txt^^second-content"}},
	}, {
		name: "one clinical note followed by an addendum",
		steps: []*pathway.ClinicalNote{{
			ContentType:     "txt",
			DocumentTitle:   "title",
			DocumentID:      "id",
			DocumentType:    "type",
			DocumentContent: "content",
		}, {
			ContentType:     "txt",
			DocumentTitle:   "new-title",
			DocumentID:      "id",
			DocumentType:    "new-type",
			DocumentContent: "new-content",
		}},
		wantFillerOrderNumber: []*hl7.CWE{
			{Identifier: hl7.NewST("type"), Text: hl7.NewST("type"), AlternateText: hl7.NewST("title")},
			{Identifier: hl7.NewST("new-type"), Text: hl7.NewST("new-type"), AlternateText: hl7.NewST("new-title")},
		},
		wantObservationValues: [][]string{{"^^txt^^content"}, {"^^txt^^content", "^^txt^^new-content"}},
	}} {
		t.Run(tt.name, func(t *testing.T) {
			now := time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC)
			var steps []pathway.Step
			for _, step := range tt.steps {
				steps = append(steps, pathway.Step{ClinicalNote: step})
			}
			pathways := map[string]pathway.Pathway{
				testPathwayName: {Pathway: steps},
			}

			hospital := hospitalWithTime(ctx, t, Config{}, pathways, now)
			defer hospital.Close()
			startPathway(t, hospital, testPathwayName)

			_, msgs := hospital.ConsumeQueues(ctx, t)
			if got, want := len(msgs), len(tt.steps); got != want {
				t.Fatalf("len(msgs)=%d, want %d", got, want)
			}
			for index, m := range msgs {
				t.Run(fmt.Sprintf("step index: %d", index), func(t *testing.T) {
					if got, want := testhl7.MessageType(t, m), "ORU^R01"; got != want {
						t.Errorf(" testhl7.MessageType(t, %v)=%q want %q", m, got, want)
					}
					obr := testhl7.OBR(t, m)
					if got, want := obr.DiagnosticServSectID.String(), message.DiagnosticServIDMDOC; got != want {
						t.Errorf("obr.DiagnosticServSectID.String()=%s, want %s", got, want)
					}
					if obr.FillerOrderNumber == nil {
						t.Fatal("obr.FillerOrderNumber=<nil> want *hl7.EI{...} object")
					}
					step := tt.steps[index]
					if gotDocID, wantDocID := obr.FillerOrderNumber.EntityIdentifier.String(), step.DocumentID; gotDocID != wantDocID {
						t.Errorf("FillerOrderNumber.NamespaceID got %q, want %q", gotDocID, wantDocID)
					}
					if diff := cmp.Diff(tt.wantFillerOrderNumber[index], obr.UniversalServiceIdentifier); diff != "" {
						t.Errorf("obr.UniversalServiceIdentifier mismatch (-want, +got):\n%s", diff)
					}
					obxs := testhl7.AllOBX(t, m)
					if got, want := len(obxs), len(tt.wantObservationValues[index]); got != want {
						t.Fatalf("len(obxs)=%d want %d", got, want)
					}
					for obxIndex, obx := range obxs {
						if obx.ResponsibleObserver == nil {
							t.Errorf("obx index: %d; obx.ResponsibleObserver=<nil>, want something", obxIndex)
						}
						if got, want := string(obx.ObservationValue[0]), tt.wantObservationValues[index][obxIndex]; got != want {
							t.Errorf("obx index: %d; obx.ObservationValue=%q want %q", obxIndex, got, want)
						}
					}
				})
			}
		})
	}
}

func TestHasEvents_HasMessages_RunNextEventIfDue_ProcessNextMessageIfDue(t *testing.T) {
	ctx := context.Background()
	pathways := map[string]pathway.Pathway{
		testPathwayName: {Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
		}},
	}
	hospital := newHospital(ctx, t, Config{}, pathways)
	defer hospital.Close()

	// At the beginning, there are no events and no messages.
	if got, want := hospital.HasEvents(), false; got != want {
		t.Errorf("hospital.HasEvents() got %t, want %t", got, want)
	}
	if ran, err := hospital.RunNextEventIfDue(ctx); ran || err != nil {
		t.Errorf("hospital.RunNextEventIfDue() got (%v, %v), want (false, nil)", ran, err)
	}
	if got, want := hospital.HasMessages(), false; got != want {
		t.Errorf("hospital.HasMessages() got %t, want %t", got, want)
	}
	if ran, err := hospital.ProcessNextMessageIfDue(); ran || err != nil {
		t.Errorf("hospital.ProcessNextMessageIfDue() got (%v, %v), want (false, nil)", ran, err)
	}

	// When the pathway starts, there are only events.
	startPathway(t, hospital, testPathwayName)
	if got, want := hospital.HasEvents(), true; got != want {
		t.Errorf("hospital.HasEvents() got %t, want %t", got, want)
	}
	if got, want := hospital.HasMessages(), false; got != want {
		t.Errorf("hospital.HasMessages() got %t, want %t", got, want)
	}
	if ran, err := hospital.ProcessNextMessageIfDue(); ran || err != nil {
		t.Errorf("hospital.ProcessNextMessageIfDue() got (%v, %v), want (false, nil)", ran, err)
	}

	// When the event runs, there are messages.
	if ran, err := hospital.RunNextEventIfDue(ctx); !ran || err != nil {
		t.Errorf("hospital.RunNextEventIfDue() got (%v, %v), want (true, nil)", ran, err)
	}
	if got, want := hospital.HasEvents(), false; got != want {
		t.Errorf("hospital.HasEvents() got %t, want %t", got, want)
	}
	if got, want := hospital.HasMessages(), true; got != want {
		t.Errorf("hospital.HasMessages() got %t, want %t", got, want)
	}
	if ran, err := hospital.ProcessNextMessageIfDue(); !ran || err != nil {
		t.Errorf("hospital.ProcessNextMessageIfDue() got (%v, %v), want (true, nil)", ran, err)
	}

	// After everything is processed, there are no messages or events.
	if got, want := hospital.HasEvents(), false; got != want {
		t.Errorf("hospital.HasEvents() got %t, want %t", got, want)
	}
	if got, want := hospital.HasMessages(), false; got != want {
		t.Errorf("hospital.HasMessages() got %t, want %t", got, want)
	}
}

func TestStartNextPathway(t *testing.T) {
	ctx := context.Background()
	rand.Seed(1)
	pathways := map[string]pathway.Pathway{
		"pathway1": {Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLocAE}},
		}},
		"pathway2": {Pathway: []pathway.Step{
			{Admission: &pathway.Admission{Loc: testLoc}},
			{Result: &pathway.Results{}},
			{Discharge: &pathway.Discharge{}},
		}},
	}
	hospital := newHospital(ctx, t, Config{}, pathways)
	defer hospital.Close()
	runs := 100
	// We expect to have either 1 or 3 messages per pathway, depending on which one gets picked every time.
	countMessageLen := map[int]int{1: 0, 3: 0}
	for i := 0; i < runs; i++ {
		if err := hospital.StartNextPathway(); err != nil {
			t.Fatalf("StartNextPathway() failed with %v", err)
		}
		_, messages := hospital.ConsumeQueues(ctx, t)
		got := len(messages)
		if _, ok := countMessageLen[got]; !ok {
			t.Errorf("StartPathway(%v) generated %v messages, want 1 or 3", testPathwayName, got)
		}
		countMessageLen[got]++
	}
	want := runs / len(countMessageLen)
	delta := float64(want / 5) // Allow some error.

	for k, v := range countMessageLen {
		if math.Abs(float64(v)-float64(want)) >= delta {
			t.Errorf("countMessageLen[%v] = %d, want within %.1f of %d", k, v, delta, want)
		}
	}
}
