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

// Package clock provides convenient functionality to manage the time.
package clock

import "time"

// Clock provides functionality to manage the time.
type Clock interface {
	// Now is the current time as seen by the clock, in UTC.
	Now() time.Time
}

// RealTimeClock is a real time Clock.
type RealTimeClock struct{}

// Now is the current time in UTC.
func (c *RealTimeClock) Now() time.Time {
	return time.Now().UTC()
}
