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

package logging

import (
	"testing"
)

func TestPackageName(t *testing.T) {
	tests := []struct {
		filename string
		want     string
	}{{
		filename: "/test/xxx/simulated_hospital/file.go",
		want:     "",
	}, {
		filename: "/test/xxx/simulated_hospital/x/file.go",
		want:     "x",
	}, {
		filename: "/test/xxx/simulated_hospital/x/y/file.go",
		want:     "x/y",
	}, {
		filename: "/test/xxx/simulated_hospital/x/y/z/file.go",
		want:     "x/y/z",
	}, {
		filename: "/test/xxx/file.go",
		want:     "/test/xxx",
	}}
	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			if got := packageName(tt.filename); got != tt.want {
				t.Errorf("packageName(%v) = %q, want %q", tt.filename, got, tt.want)
			}
		})
	}
}
