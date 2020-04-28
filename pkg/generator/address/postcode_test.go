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

package address

import (
	"math/rand"
	"regexp"
	"testing"
)

func TestRandomPostcode(t *testing.T) {
	rand.Seed(1)

	tests := []struct {
		name string
		g    PostcodeGenerator
		r    *regexp.Regexp
	}{{
		name: "UK",
		g:    &UKPostcode{},
		r:    regexp.MustCompile(`^[A-Z]{2}[0-9][0-9]?\s?[0-9][A-Z]{2}$`),
	}, {
		name: "US",
		g:    &USPostcode{},
		r:    regexp.MustCompile(`^[0-9]{5}$`),
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i < 100; i++ {
				got := tc.g.Random()
				if !tc.r.MatchString(got) {
					t.Errorf(`g.Random() = %q, want matching regex %v`, got, tc.r)
				}
			}
		})
	}
}
