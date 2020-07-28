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

package config

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/test/testwrite"
)

func TestLoadHL7Config(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		config  []byte
		wantErr bool
		wantOC  string
	}{{
		name: "good config",
		config: []byte(`
order_control:
  new: want-order-control-new`),
		wantErr: false,
		wantOC:  "want-order-control-new",
	}, {
		name:    "unknown fields",
		config:  []byte(`arbitrary_field: want-sending-application`),
		wantErr: true,
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tmp := testwrite.BytesToFile(t, []byte(tc.config))
			c, err := LoadHL7Config(ctx, tmp)
			gotErr := err != nil
			if gotErr != tc.wantErr {
				t.Errorf("LoadHL7Config(%s) got err %v; want error? %t", tmp, err, tc.wantErr)
			}
			if gotErr || tc.wantErr {
				return
			}
			if got, want := c.OrderControl.New, tc.wantOC; got != want {
				t.Errorf("OrderControl.New got %q, want %q", got, want)
			}
		})
	}
}

func TestLoadHeaderConfig(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name        string
		header      []byte
		wantErr     bool
		wantDefault *HeaderForType
		wantORU     *HeaderForType
	}{{
		name: "Default only",
		header: []byte(`
default:
  sending_application: want-sa
  sending_facility: want-sf
  receiving_application: want-ra
  receiving_facility: want-rf
`),
		wantDefault: &HeaderForType{
			SendingFacility:      "want-sf",
			SendingApplication:   "want-sa",
			ReceivingApplication: "want-ra",
			ReceivingFacility:    "want-rf",
		},
	}, {
		name: "Default and Override ORU",
		header: []byte(`
default:
  sending_application: want-sa
  sending_facility: want-sf
  receiving_application: want-ra
  receiving_facility: want-rf
oru:
  sending_application: want-sa-oru
  sending_facility: want-sf-oru
  receiving_application: want-ra-oru
  receiving_facility: want-rf-oru
`),
		wantDefault: &HeaderForType{
			SendingFacility:      "want-sf",
			SendingApplication:   "want-sa",
			ReceivingApplication: "want-ra",
			ReceivingFacility:    "want-rf",
		},
		wantORU: &HeaderForType{
			SendingFacility:      "want-sf-oru",
			SendingApplication:   "want-sa-oru",
			ReceivingApplication: "want-ra-oru",
			ReceivingFacility:    "want-rf-oru",
		},
	}, {
		name: "Default missing",
		header: []byte(`
oru:
  sending_application: want-sa-oru
  sending_facility: want-sf-oru
  receiving_application: want-ra-oru
  receiving_facility: want-rf-oru
`),
		wantErr: true,
	}, {
		name: "Missing Default.ReceivingFacility",
		header: []byte(`
default:
  sending_application: want-sa
  sending_facility: want-sf
  receiving_application: want-ra
`),
		wantErr: true,
	}, {
		name: "Missing Default.ReceivingApplication",
		header: []byte(`
default:
  sending_application: want-sa
  sending_facility: want-sf
  receiving_facility: want-rf
`),
		wantErr: true,
	}, {
		name: "Missing Default.SendingApplication",
		header: []byte(`
default:
  sending_facility: want-sf
  receiving_application: want-ra
  receiving_facility: want-rf
`),
		wantErr: true,
	}, {
		name: "Missing Default.SendingFacility",
		header: []byte(`
default:
  sending_application: want-sa
  receiving_application: want-ra
  receiving_facility: want-rf
`),
		wantErr: true,
	}, {
		name: "Missing OverrideORU.SendingApplication",
		header: []byte(`
default:
  sending_application: want-sa
  sending_facility: want-sf
  receiving_application: want-ra
  receiving_facility: want-rf
oru:
  sending_facility: want-sf-oru
  receiving_application: want-ra-oru
  receiving_facility: want-rf-oru
`),
		wantErr: true,
	}, {
		name: "Missing OverrideORU.SendingFacility",
		header: []byte(`
default:
  sending_application: want-sa
  sending_facility: want-sf
  receiving_application: want-ra
  receiving_facility: want-rf
oru:
  sending_application: want-sa-oru
  receiving_application: want-ra-oru
  receiving_facility: want-rf-oru
`),
		wantErr: true,
	}, {
		name: "Missing OverrideORU.ReceivingApplication",
		header: []byte(`
default:
  sending_application: want-sa
  sending_facility: want-sf
  receiving_application: want-ra
  receiving_facility: want-rf
oru:
  sending_application: want-sa-oru
  sending_facility: want-sf-oru
  receiving_facility: want-rf-oru
`),
		wantErr: true,
	}, {
		name: "Missing OverrideORU.ReceivingFacility",
		header: []byte(`
default:
  sending_application: want-sa
  sending_facility: want-sf
  receiving_application: want-ra
  receiving_facility: want-rf
oru:
  sending_application: want-sa-oru
  sending_facility: want-sf-oru
  receiving_application: want-ra-oru
`),
		wantErr: true,
	}, {
		name: "Unknown fields",
		header: []byte(`
default:
  sending_application: want-sa
  sending_facility: want-sf
  receiving_application: want-ra
  receiving_facility: want-rf
	unknown_field: unknown-field
`),
		wantErr: true,
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			headerFile := testwrite.BytesToFile(t, tc.header)
			h, err := LoadHeaderConfig(ctx, headerFile)
			gotErr := err != nil
			if gotErr != tc.wantErr {
				t.Errorf("LoadHeaderConfig(%s) got err %v; want error? %t", headerFile, err, tc.wantErr)
			}
			if gotErr || tc.wantErr {
				return
			}
			if diff := cmp.Diff(tc.wantDefault, &h.Default); diff != "" {
				t.Errorf("Header.Default got mismatch (-want, +got):\n%s", diff)
			}
			if diff := cmp.Diff(tc.wantORU, h.ORU); diff != "" {
				t.Errorf("Header.ORU got mismatch (-want, +got):\n%s", diff)
			}
		})
	}
}
