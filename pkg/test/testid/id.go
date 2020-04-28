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

// Package testid contains functionality to generate identifiers in a deterministic way.
package testid

import "strconv"

// Generator is a generator of identifiers.
type Generator struct {
	nID int
}

// NewID returns a new identifier that increments with every invocation, starting with "1".
func (g *Generator) NewID() string {
	g.nID++
	return strconv.Itoa(g.nID)
}
