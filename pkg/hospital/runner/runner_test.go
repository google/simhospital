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
	"net/http"
	"testing"

	"github.com/google/simhospital/pkg/hospital"
	. "github.com/google/simhospital/pkg/hospital/runner"
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
