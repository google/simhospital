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

package rate

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestHeartbeat(t *testing.T) {
	cases := []struct {
		rate        float64
		per         time.Duration
		want        time.Duration
		wantElapsed time.Duration
	}{
		{
			rate:        1,
			per:         time.Hour,
			want:        time.Hour,
			wantElapsed: time.Hour,
		}, {
			rate:        2,
			per:         time.Hour,
			want:        30 * time.Minute,
			wantElapsed: 30 * time.Minute,
		}, {
			rate:        4,
			per:         time.Minute,
			want:        15 * time.Second,
			wantElapsed: 15 * time.Second,
		}, {
			rate:        4,
			per:         0,
			want:        0,
			wantElapsed: 0,
		}, {
			rate:        0,
			per:         time.Hour,
			want:        time.Duration(1<<63 - 1),
			wantElapsed: 0,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.want), func(t *testing.T) {
			c := NewController(tc.rate, tc.per)
			if got := c.Heartbeat(); got != tc.want {
				t.Errorf("(%+v).Heartbeat() got %v; want %v", c, got, tc.want)
			}
			if got := c.InitialElapsed(); got != tc.wantElapsed {
				t.Errorf("(%+v).InitialElapsed() got %v; want %v", c, got, tc.wantElapsed)
			}
		})
	}
}

func TestMessageRateHandlerGET(t *testing.T) {
	want := float64(1)
	c := NewController(want, time.Hour)
	ts := httptest.NewServer(c)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatalf(`http.NewRequest("GET", %v, nil) failed with %v`, ts.URL, err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("http.DefaultClient.Do(%v) failed with %v", req, err)
	}
	defer resp.Body.Close()

	gotBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll(%v) failed with %v", resp.Body, err)
	}

	got, err := strconv.ParseFloat(string(gotBody), 64)
	if err != nil {
		t.Fatalf("strconv.ParseFloat(string(%v), 64) failed with %v", gotBody, err)
	}

	if got != want {
		t.Errorf("GET request got %v; want %v", got, want)
	}
}

func TestMessageRateHandlerPOST(t *testing.T) {
	c := NewController(1, time.Hour)

	postDone := make(chan bool, 1)
	defer close(postDone)
	go func() {
		select {
		case r := <-c.RateChanged():
			if !r {
				t.Errorf("c.RateChanged() got %v; want true", r)
			}
		case <-postDone:
			t.Error("c.RateChanged() returned no value")
		}
	}()

	ts := httptest.NewServer(c)
	defer ts.Close()

	newValue := "value=2.5"
	postReq, err := http.NewRequest("POST", ts.URL, strings.NewReader(newValue))
	if err != nil {
		t.Fatalf(`http.NewRequest("POST", %v, %v) failed with %v`, ts.URL, newValue, err)
	}

	_, err = http.DefaultClient.Do(postReq)
	if err != nil {
		t.Fatalf("http.DefaultClient.Do(%v) failed with %v", postReq, err)
	}
	postDone <- true

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatalf(`http.NewRequest("GET", %v, nil) failed with %v`, ts.URL, err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("http.DefaultClient.Do(%v) failed with %v", req, err)
	}
	defer resp.Body.Close()

	gotBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll(%v) failed with %v", resp.Body, err)
	}

	got, err := strconv.ParseFloat(string(gotBody), 64)
	if err != nil {
		t.Fatalf("strconv.ParseFloat(string(%v), 64) failed with %v", gotBody, err)
	}

	if want := 2.5; got != want {
		t.Errorf("GET request got %v; want %v", got, want)
	}
}

func TestMessageRateHandlerPOSTError(t *testing.T) {
	c := NewController(1, time.Hour)

	// RateChanged channel is a blocking channel.
	// Continuously consume values from it to unblock successful POST requests.
	// The actual values are ignored.
	done := make(chan bool)
	defer close(done)
	go func() {
		for {
			select {
			case <-c.RateChanged():
				continue
			case <-done:
				return
			}
		}
	}()

	ts := httptest.NewServer(c)
	defer ts.Close()

	cases := []struct {
		value    string
		wantCode int
	}{
		{value: "value=2.5", wantCode: http.StatusOK},
		{value: "invalid=2.5", wantCode: http.StatusInternalServerError},
		{value: "=2.5", wantCode: http.StatusInternalServerError},
		{value: "2.5", wantCode: http.StatusInternalServerError},
		{value: "value=2.5.3", wantCode: http.StatusInternalServerError},
		{value: "value=2,5", wantCode: http.StatusInternalServerError},
		{value: "value=newvalue", wantCode: http.StatusInternalServerError},
		{value: "value=", wantCode: http.StatusInternalServerError},
	}

	for _, tc := range cases {
		t.Run(tc.value, func(t *testing.T) {
			postReq, err := http.NewRequest("POST", ts.URL, strings.NewReader(tc.value))
			if err != nil {
				t.Fatalf(`http.NewRequest("POST", %v, %v) failed with %v`, ts.URL, tc.value, err)
			}

			resp, err := http.DefaultClient.Do(postReq)
			if err != nil {
				t.Fatalf("http.DefaultClient.Do(%v) failed with %v", postReq, err)
			}

			if gotCode := resp.StatusCode; gotCode != tc.wantCode {
				t.Errorf("POST response status code got %v; want %v", gotCode, tc.wantCode)
			}
		})
	}
}
