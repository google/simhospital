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

package runner_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/simhospital/pkg/hl7"
	"github.com/google/simhospital/pkg/hospital"
	. "github.com/google/simhospital/pkg/hospital/runner"
	"github.com/google/simhospital/pkg/test/testclock"
	"github.com/google/simhospital/pkg/test/testhospital"
	"github.com/google/simhospital/pkg/test/testwrite"
)

const (
	nonEmptyString        = "nonempty"
	validDashboardAddress = ":8000"
)

var (
	testAPIHandler            = func(_ http.ResponseWriter, _ *http.Request) {}
	testEndpointAndHandler    = EndpointAndHandler{Endpoint: "testEndpoint", Handler: testAPIHandler}
	testAPIEndpointAndHandler = APIEndpointAndHandler{EndpointAndHandler: testEndpointAndHandler, HTTPMethod: "POST"}
)

func TestNewRunner(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{{
		name: "valid api endpoint and nonempty config",
		config: Config{
			DashboardURI:           nonEmptyString,
			DashboardAddress:       validDashboardAddress,
			DashboardStaticDir:     nonEmptyString,
			AuthenticatedEndpoints: []APIEndpointAndHandler{testAPIEndpointAndHandler},
			AuthenticatedAPIConfig: APIConfig{APIPort: nonEmptyString, APIKey: nonEmptyString},
		},
		wantErr: false,
	}, {
		name: "no api endpoints and nonempty config",
		config: Config{
			DashboardURI:           nonEmptyString,
			DashboardAddress:       validDashboardAddress,
			DashboardStaticDir:     nonEmptyString,
			AuthenticatedEndpoints: []APIEndpointAndHandler{},
			AuthenticatedAPIConfig: APIConfig{APIPort: nonEmptyString, APIKey: nonEmptyString},
		},
		wantErr: false,
	}, {
		name: "valid api endpoint and empty API config",
		config: Config{
			DashboardURI:           nonEmptyString,
			DashboardAddress:       validDashboardAddress,
			DashboardStaticDir:     nonEmptyString,
			AuthenticatedEndpoints: []APIEndpointAndHandler{testAPIEndpointAndHandler},
			AuthenticatedAPIConfig: APIConfig{},
		},
		wantErr: true,
	}, {
		name: "valid api endpoint and missing APIPort",
		config: Config{
			DashboardURI:           nonEmptyString,
			DashboardAddress:       validDashboardAddress,
			DashboardStaticDir:     nonEmptyString,
			AuthenticatedEndpoints: []APIEndpointAndHandler{testAPIEndpointAndHandler},
			AuthenticatedAPIConfig: APIConfig{APIKey: nonEmptyString},
		},
		wantErr: true,
	}, {
		name: "valid api endpoint and missing APIKey",
		config: Config{
			DashboardURI:           nonEmptyString,
			DashboardAddress:       validDashboardAddress,
			DashboardStaticDir:     nonEmptyString,
			AuthenticatedEndpoints: []APIEndpointAndHandler{testAPIEndpointAndHandler},
			AuthenticatedAPIConfig: APIConfig{APIPort: nonEmptyString},
		},
		wantErr: true,
	}, {
		name: "no api endpoints and valid dashboard config",
		config: Config{
			DashboardURI:       nonEmptyString,
			DashboardAddress:   validDashboardAddress,
			DashboardStaticDir: nonEmptyString,
		},
		wantErr: false,
	}, {
		name: "missing DashboardURI",
		config: Config{
			DashboardAddress:   validDashboardAddress,
			DashboardStaticDir: nonEmptyString,
		},
		wantErr: true,
	}, {
		name: "missing DashboardAddress",
		config: Config{
			DashboardURI:       nonEmptyString,
			DashboardStaticDir: nonEmptyString,
		},
		wantErr: true,
	}, {
		name: "invalid DashboardAddress",
		config: Config{
			DashboardURI:       nonEmptyString,
			DashboardAddress:   nonEmptyString,
			DashboardStaticDir: nonEmptyString,
		},
		wantErr: true,
	}, {
		name: "missing DashboardStaticDir",
		config: Config{
			DashboardURI:     nonEmptyString,
			DashboardAddress: validDashboardAddress,
		},
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(&hospital.Hospital{}, tt.config)
			gotErr := err != nil
			if gotErr != tt.wantErr {
				t.Errorf("New(%+v) got error: %v, did want error? %t", tt.config, err, tt.wantErr)
			}
		})
	}
}

// This test only covers the case when Run() stops.
func TestRunner_Run(t *testing.T) {
	ctx := context.Background()
	// test_pathway is an arbitrary pathway that sends 4 (arbitrary number) messages.
	// This pathway refers to values in the files in the pkg/test/data package,
	// more specifically the order profiles and the locations.
	// This test will fail if those values change.
	b := []byte(`
test_pathway:
  historical_data:
    - result:
        order_profile: UREA AND ELECTROLYTES
        results:
          - test_name: Creatinine
            value: 126.00
            unit: UMOLL
            abnormal_flag: HIGH
      parameters:
        time_from_now: -48h
  pathway:
    - admission:
        loc: Renal
    - result:
        order_profile: UREA AND ELECTROLYTES
        results:
          - test_name: Creatinine
            value: 153.00
            unit: UMOLL
            abnormal_flag: HIGH
    - discharge: {}`)
	mainDir := testwrite.BytesToDir(t, b, "pathway.yml")

	hl7.TimezoneAndLocation("Europe/London")
	// now is an arbitrary date in the past.
	now := time.Date(2020, 2, 12, 0, 0, 0, 0, time.UTC)

	args := testhospital.Arguments
	args.PathwayArguments.Dir = mainDir
	args.PathwayArguments.Names = []string{"test_pathway"}

	tests := []struct {
		maxPathways  int
		wantMessages int
	}{
		{maxPathways: 0, wantMessages: 0},
		{maxPathways: 1, wantMessages: 4},
		{maxPathways: 2, wantMessages: 8},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d", tc.maxPathways), func(t *testing.T) {
			clock := testclock.WithTick(now, time.Second)

			h := testhospital.New(ctx, t, testhospital.Config{
				Config:    hospital.Config{Clock: clock},
				Arguments: args,
			})
			defer h.Close()

			config := Config{
				DashboardURI:       nonEmptyString,
				DashboardAddress:   ":0000",
				DashboardStaticDir: nonEmptyString,
				MaxPathways:        tc.maxPathways,
				PathwaysPerHour:    3600, // Create the pathways quickly.
				Clock:              clock,
			}

			runner, err := New(h.Hospital, config)
			if err != nil {
				t.Fatalf("New(%+v) failed with %v", config, err)
			}
			runner.Run(context.Background())
			messages := h.Sender.GetSentMessages()
			if got, want := len(messages), tc.wantMessages; got != want {
				t.Errorf("h.Sender.GetSentMessages() got %d messages, want %v", got, want)
			}
		})
	}
}
