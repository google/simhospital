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

package cloud

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/api/option"
)

const (
	project  = "test-id"
	location = "test-location"
	dataset  = "test-dataset"
	store    = "test-store"
)

func TestOutputWrite(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	bytesToWrite := []byte("test bytes")

	tests := []struct {
		name       string
		statusCode int
		wantErr    bool
	}{{
		name:       "Status OK",
		statusCode: 200,
	}, {
		name:       "Status error",
		statusCode: 403,
		wantErr:    true,
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var got []byte
			var err error
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if got, err = ioutil.ReadAll(r.Body); err != nil {
					t.Errorf("ioutil.ReadAll(%v) failed with %v", r.Body, err)
					w.WriteHeader(http.StatusBadRequest)
				}
				w.WriteHeader(tc.statusCode)
				w.Write([]byte("response body"))
			}))
			defer ts.Close()

			opts := []option.ClientOption{option.WithHTTPClient(ts.Client()), option.WithEndpoint(ts.URL)}
			output, err := NewOutput(ctx, project, location, dataset, store, opts...)
			if err != nil {
				t.Fatalf("NewOutput(%v, %s, %s, %s, %s, %v) failed with %v", ctx, project, location, dataset, store, opts, err)
			}

			writer, err := output.New(nil)
			if err != nil {
				t.Fatalf("%T.New(%v) failed with %v", output, nil, err)
			}
			defer writer.Close()

			n, err := writer.Write(bytesToWrite)
			if gotErr := (err != nil); gotErr != tc.wantErr {
				t.Errorf("%T.Write(%s) got err %v, want error? %t", writer, bytesToWrite, err, tc.wantErr)
			}

			if gotN, wantN := n, len(bytesToWrite); !tc.wantErr && gotN != wantN {
				t.Errorf("%T.Write(%s) = %d, want %d", writer, bytesToWrite, gotN, wantN)
			}

			if diff := cmp.Diff(bytesToWrite, got); diff != "" {
				t.Errorf("%T.Write(%s) returned diff (-want +got):\n%s", writer, bytesToWrite, diff)
			}
		})
	}
}
