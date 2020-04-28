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

// Package testclock provides a convenient Clock implementation for testing.
package testclock

import (
	"time"
)

// Clock is a clock used for testing.
type Clock struct {
	now  time.Time
	tick time.Duration
}

// New creates a new Clock for testing with a tick of 0.
// For this clock, Now() always returns the value the clock was created with.
func New(now time.Time) *Clock {
	return WithTick(now, time.Duration(0))
}

// WithTick creates a new Clock for testing with the time and tick specified.
func WithTick(now time.Time, tick time.Duration) *Clock {
	return &Clock{now: now, tick: tick}
}

// Now returns the current time as seen by the Clock and advances the time
// the duration of the tick.
func (c *Clock) Now() time.Time {
	nowToReturn := c.now
	c.Advance(c.tick)
	return nowToReturn
}

// Advance advances the clock the specified duration.
func (c *Clock) Advance(d time.Duration) {
	c.now = c.now.Add(d)
}
