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

package codedelement

import (
	"testing"
	"time"
)

func TestSimpleDateGenerator(t *testing.T) {
	now := time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)
	oneYearAgo := time.Date(2018, 1, 20, 0, 0, 0, 0, time.UTC)

	sdg := SimpleDateGenerator{}

	for i := 0; i < 100; i++ {
		got := sdg.Random(now)
		if got.After(now) || got.Before(oneYearAgo) {
			t.Errorf("sdg.Random(%v)=%v, want between (%v, %v)", now, oneYearAgo, now)
		}
	}
}
