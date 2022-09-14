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

package header

import (
	"context"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/logging"
	"github.com/google/simhospital/pkg/message"
	"github.com/google/simhospital/pkg/pathway"
	"github.com/google/simhospital/pkg/test"
	"github.com/google/simhospital/pkg/test/testwrite"
)

var arbitraryStep = &pathway.Step{Discharge: &pathway.Discharge{}}

func TestMain(m *testing.M) {
	logging.SetLogLevel(logrus.DebugLevel)

	retCode := m.Run()

	os.Exit(retCode)
}

func TestMessageControlIDsAreIncremental(t *testing.T) {
	ctx := context.Background()
	headerCFG, err := config.LoadHeaderConfig(ctx, test.HeaderConfigTest)
	if err != nil {
		t.Fatalf("LoadHeaderConfig(%s) failed with %v", test.HeaderConfigTest, err)
	}
	g := &Generator{Header: headerCFG, MsgCtrlGen: &MessageControlGenerator{}}

	// We have different ways to get the Message Control ID. Make sure that all of them
	// move the ID forward.
	wantIDs := []string{"1", "2", "3", "4"}
	h := g.NewHeader(arbitraryStep)
	var gotIDs []string
	gotIDs = append(gotIDs, h.MessageControlID)
	gotIDs = append(gotIDs, g.MsgCtrlGen.NewMessageControlID())
	gotIDs = append(gotIDs, g.NewHeader(arbitraryStep).MessageControlID)
	gotIDs = append(gotIDs, g.MsgCtrlGen.NewMessageControlID())

	if diff := cmp.Diff(wantIDs, gotIDs); diff != "" {
		t.Errorf("MessageControlID got diff (-want, +got):\n%s ", diff)
	}
}

func TestNewHeader(t *testing.T) {
	headerFile := testwrite.BytesToFile(t, []byte(`
default:
  sending_application: default_sa
  sending_facility: default_sf
  receiving_application: default_ra
  receiving_facility: default_rf
oru:
  sending_application: oru_sa
  sending_facility: oru_sf
  receiving_application: oru_ra
  receiving_facility: oru_rf
`))

	ctx := context.Background()
	headerCFG, err := config.LoadHeaderConfig(ctx, headerFile)
	if err != nil {
		t.Fatalf("LoadHeaderConfig(%s) failed with %v", test.HeaderConfigTest, err)
	}
	tests := []struct {
		name string
		step *pathway.Step
		want *message.HeaderInfo
	}{{
		name: "Non ORU with no overrides",
		step: &pathway.Step{Admission: &pathway.Admission{}},
		want: &message.HeaderInfo{
			SendingApplication:   "default_sa",
			SendingFacility:      "default_sf",
			ReceivingApplication: "default_ra",
			ReceivingFacility:    "default_rf",
			MessageControlID:     "1",
		},
	}, {
		name: "Order with no overrides",
		step: &pathway.Step{Order: &pathway.Order{}},
		want: &message.HeaderInfo{
			SendingApplication:   "default_sa",
			SendingFacility:      "default_sf",
			ReceivingApplication: "default_ra",
			ReceivingFacility:    "default_rf",
			MessageControlID:     "1",
		},
	}, {
		name: "ORU with no overrides",
		step: &pathway.Step{Result: &pathway.Results{}},
		want: &message.HeaderInfo{
			SendingApplication:   "oru_sa",
			SendingFacility:      "oru_sf",
			ReceivingApplication: "oru_ra",
			ReceivingFacility:    "oru_rf",
			MessageControlID:     "1",
		},
	}, {
		name: "Non ORU overrides Sending Facility",
		step: &pathway.Step{
			Admission: &pathway.Admission{},
			Parameters: &pathway.Parameters{
				SendingFacility: "override",
			},
		},
		want: &message.HeaderInfo{
			SendingApplication:   "default_sa",
			SendingFacility:      "override",
			ReceivingApplication: "default_ra",
			ReceivingFacility:    "default_rf",
			MessageControlID:     "1",
		},
	}, {
		name: "Non ORU overrides Sending Application",
		step: &pathway.Step{
			Admission: &pathway.Admission{},
			Parameters: &pathway.Parameters{
				SendingApplication: "override",
			},
		},
		want: &message.HeaderInfo{
			SendingApplication:   "override",
			SendingFacility:      "default_sf",
			ReceivingApplication: "default_ra",
			ReceivingFacility:    "default_rf",
			MessageControlID:     "1",
		},
	}, {
		name: "Non ORU overrides Receiving Facility",
		step: &pathway.Step{
			Admission: &pathway.Admission{},
			Parameters: &pathway.Parameters{
				ReceivingFacility: "override",
			},
		},
		want: &message.HeaderInfo{
			SendingApplication:   "default_sa",
			SendingFacility:      "default_sf",
			ReceivingApplication: "default_ra",
			ReceivingFacility:    "override",
			MessageControlID:     "1",
		},
	}, {
		name: "Non ORU overrides Receiving Application",
		step: &pathway.Step{
			Admission: &pathway.Admission{},
			Parameters: &pathway.Parameters{
				ReceivingApplication: "override",
			},
		},
		want: &message.HeaderInfo{
			SendingApplication:   "default_sa",
			SendingFacility:      "default_sf",
			ReceivingApplication: "override",
			ReceivingFacility:    "default_rf",
			MessageControlID:     "1",
		},
	}, {
		name: "ORU overrides Sending Facility",
		step: &pathway.Step{
			Result: &pathway.Results{},
			Parameters: &pathway.Parameters{
				SendingFacility: "override",
			},
		},
		want: &message.HeaderInfo{
			SendingApplication:   "oru_sa",
			SendingFacility:      "override",
			ReceivingApplication: "oru_ra",
			ReceivingFacility:    "oru_rf",
			MessageControlID:     "1",
		},
	}, {
		name: "ORU overrides Sending Application",
		step: &pathway.Step{
			Result: &pathway.Results{},
			Parameters: &pathway.Parameters{
				SendingApplication: "override",
			},
		},
		want: &message.HeaderInfo{
			SendingApplication:   "override",
			SendingFacility:      "oru_sf",
			ReceivingApplication: "oru_ra",
			ReceivingFacility:    "oru_rf",
			MessageControlID:     "1",
		},
	}, {
		name: "ORU overrides Receiving Facility",
		step: &pathway.Step{
			Result: &pathway.Results{},
			Parameters: &pathway.Parameters{
				ReceivingFacility: "override",
			},
		},
		want: &message.HeaderInfo{
			SendingApplication:   "oru_sa",
			SendingFacility:      "oru_sf",
			ReceivingApplication: "oru_ra",
			ReceivingFacility:    "override",
			MessageControlID:     "1",
		},
	}, {
		name: "ORU overrides Receiving Application",
		step: &pathway.Step{
			Result: &pathway.Results{},
			Parameters: &pathway.Parameters{
				ReceivingApplication: "override",
			},
		},
		want: &message.HeaderInfo{
			SendingApplication:   "oru_sa",
			SendingFacility:      "oru_sf",
			ReceivingApplication: "override",
			ReceivingFacility:    "oru_rf",
			MessageControlID:     "1",
		},
	}, {
		name: "ORU overrides all",
		step: &pathway.Step{
			Result: &pathway.Results{},
			Parameters: &pathway.Parameters{
				SendingApplication:   "override_sa",
				SendingFacility:      "override_sf",
				ReceivingApplication: "override_ra",
				ReceivingFacility:    "override_rf",
			},
		},
		want: &message.HeaderInfo{
			SendingApplication:   "override_sa",
			SendingFacility:      "override_sf",
			ReceivingApplication: "override_ra",
			ReceivingFacility:    "override_rf",
			MessageControlID:     "1",
		},
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := &Generator{Header: headerCFG, MsgCtrlGen: &MessageControlGenerator{}}

			if diff := cmp.Diff(tc.want, g.NewHeader(tc.step)); diff != "" {
				t.Errorf("NewHeader(%v) got diff (-want, +got):\n%s ", tc.step, diff)
			}
		})
	}
}
