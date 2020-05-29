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

// Package testdate includes functionality to generate deterministic dates for testing.
package testdate

import (
	"time"

	"github.com/google/simhospital/pkg/ir"
)

// Generator is a generator of deterministic random dates.
type Generator struct{}

// Random returns the NullTime that corresponds to now.
func (g *Generator) Random(now time.Time) ir.NullTime {
	return ir.NewValidTime(now)
}
