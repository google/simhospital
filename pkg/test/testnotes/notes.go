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

// Package testnotes contains functionality to generate deterministic notes for testing.
package testnotes

import (
	"errors"

	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/pathway"
)

// Generator generates notes for testing.
type Generator struct {
	n     int
	Notes [][]string
}

// RandomNotesForResult returns notes deterministically, iterating through the notes the generator has been initialized with.
func (g Generator) RandomNotesForResult() []string {
	result := g.Notes[g.n]
	g.n = (g.n + 1) % len(g.Notes)
	return result
}

// RandomDocumentForClinicalNote is not supported in Generator and it returns an error.
func (g Generator) RandomDocumentForClinicalNote(_ *pathway.ClinicalNote) (*ir.ClinicalNote, error) {
	return nil, errors.New("method RandomDocumentForClinicalNote() not supported in testnotes.Generator")
}
