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

package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testPath   = "simulated-hospital-patients/api/syncPathwayStarter"
	testAPIKey = "2389r-13ywfeu"
)

// getTestHandler is needed to start the httptest server.
// The returned handlerFunc fails the test if invoked.
func getTestHandler(t *testing.T) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		t.Fatal("test entered test handler, this should not happen")
	}
}

func TestAuthenticationMiddleware(t *testing.T) {
	tests := []struct {
		description     string
		wantBody        string
		authHeaderValue string
		wantCode        int
		handlerFunc     func(rw http.ResponseWriter, req *http.Request)
	}{
		{
			description:     "Request with invalid Authorization header",
			wantBody:        "Invalid Authorization header",
			wantCode:        http.StatusUnauthorized,
			authHeaderValue: "qwerty",
			handlerFunc:     getTestHandler(t),
		},
		{
			description:     "Request with valid Authorization header",
			wantBody:        "",
			wantCode:        http.StatusOK,
			authHeaderValue: testAPIKey,
			handlerFunc: func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
			},
		},
		{
			description:     "Request with empty Authorization header",
			wantBody:        "An Authorization header is required",
			wantCode:        http.StatusBadRequest,
			authHeaderValue: "",
			handlerFunc:     getTestHandler(t),
		},
	}

	client := http.Client{}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			ts := httptest.NewServer(Middleware(tc.handlerFunc, testAPIKey))
			defer ts.Close()

			requestPath := fmt.Sprintf("%s/%s", ts.URL, testPath)
			req, err := http.NewRequest("POST", requestPath, nil)
			if err != nil {
				t.Fatalf("http.NewRequest(%v, %v) failed with %v", "POST", requestPath, err)
			}

			req.Header.Add("Authorization", tc.authHeaderValue)
			res, err := client.Do(req)
			if err != nil {
				t.Fatalf("client.Do(%v) failed with %v", req, err)
			}
			defer res.Body.Close()

			if res.StatusCode != tc.wantCode {
				t.Fatalf("client.Do(%v).StatusCode got %d, want %d", req, res.StatusCode, tc.wantCode)
			}

			if tc.wantBody == "" {
				return
			}
			var actualBody interface{}
			if err = json.NewDecoder(res.Body).Decode(&actualBody); err != nil {
				t.Fatalf("json.NewDecoder(res.Body) failed with %v", err)
			}
			if actualBody != tc.wantBody {
				t.Fatalf("client.Do(%v).Body got %q, want %q", req, actualBody, tc.wantBody)
			}
		})
	}
}
