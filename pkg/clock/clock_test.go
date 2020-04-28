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

package clock

import (
	"testing"
	"time"
)

func TestRealTimeClock(t *testing.T) {
	before := time.Now().UTC()

	rtc := &RealTimeClock{}
	first := rtc.Now()
	if first.Before(before) {
		t.Errorf("rtc.Now() got %v; want after or equal to %v", first, before)
	}

	second := rtc.Now()
	if second.Before(first) {
		t.Errorf("Second invocation of rtc.Now() got %v; want after or equal to %v", second, first)
	}

	after := time.Now().UTC()
	if after.Before(second) {
		t.Errorf("rtc.Now() got %v; want before or equal to %v", second, after)
	}
}
