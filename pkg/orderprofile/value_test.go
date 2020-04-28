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

package orderprofile

import (
	"testing"
)

func TestParseValueFloat(t *testing.T) {
	tcs := []struct {
		strVal   string
		wantPref string
		wantVal  float64
	}{
		{strVal: "70", wantPref: "", wantVal: 70},
		{strVal: "70.5", wantPref: "", wantVal: 70.5},
		{strVal: "-70.5", wantPref: "", wantVal: -70.5},
		{strVal: ">70", wantPref: ">", wantVal: 70},
		{strVal: ">70.5", wantPref: ">", wantVal: 70.5},
		{strVal: ">=70.5", wantPref: ">=", wantVal: 70.5},
		{strVal: ">=-70.5", wantPref: ">=", wantVal: -70.5},
		{strVal: ">= -70.5", wantPref: ">=", wantVal: -70.5},
		{strVal: "<70", wantPref: "<", wantVal: 70},
		{strVal: "<70.5", wantPref: "<", wantVal: 70.5},
		{strVal: "<=70.5", wantPref: "<=", wantVal: 70.5},
		{strVal: "<=-70.5", wantPref: "<=", wantVal: -70.5},
		{strVal: "<= -70.5", wantPref: "<=", wantVal: -70.5},
	}
	for _, tc := range tcs {
		t.Run(tc.strVal, func(t *testing.T) {
			gotPref, gotVal, err := ValueFromString(tc.strVal)
			if err != nil {
				t.Fatalf("ValueFromString(%q) failed with %v", tc.strVal, err)
			}
			if gotPref != tc.wantPref {
				t.Errorf("ValueFromString(%q) got prefix %q, want %q", tc.strVal, gotPref, tc.wantPref)
			}
			if gotVal != tc.wantVal {
				t.Errorf("ValueFromString(%q) got value %v, want %v", tc.strVal, gotVal, tc.wantVal)
			}
		})
	}
}

func TestParseValueCannotParse(t *testing.T) {
	badValues := []string{"not float", "<1.2.3"}
	for _, badVal := range badValues {
		t.Run(badVal, func(t *testing.T) {
			gotPref, gotVal, err := ValueFromString(badVal)
			if err == nil {
				t.Errorf("ValueFromString(%q) got err=<nil>, want non-nil error", badVal)
			}
			if gotPref != "" {
				t.Errorf("ValueFromString(%q) got prefix %q, want empty string", badVal, gotPref)
			}
			if gotVal != 0 {
				t.Errorf("ValueFromString(%q) got value %f, want 0", badVal, gotVal)
			}
		})
	}
}
