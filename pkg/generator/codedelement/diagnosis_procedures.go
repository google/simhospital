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
	"github.com/google/simhospital/pkg/clock"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/constants"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/pathway"
)

// DiagOrProcGenerator is a generator of Diagnoses or Procedures.
type DiagOrProcGenerator struct {
	*Generator
}

// NewDiagnosisGenerator creates a new generator of Diagnoses.
func NewDiagnosisGenerator(hc *config.HL7Config, d *config.Data, c clock.Clock, dg DateGenerator) *DiagOrProcGenerator {
	return &DiagOrProcGenerator{Generator: newGenerator(d.Diagnoses, hc.Diagnosis.Types, c, dg)}
}

// NewProcedureGenerator creates a new generator of Procedures.
func NewProcedureGenerator(hc *config.HL7Config, d *config.Data, c clock.Clock, dg DateGenerator) *DiagOrProcGenerator {
	return &DiagOrProcGenerator{Generator: newGenerator(d.Procedures, hc.Procedure.Types, c, dg)}
}

// RandomOrFromPathway returns a random ir.DiagnosisOrProcedure or one based on the pathway
// depending on the value of the pathway's Code or Description.
func (g *DiagOrProcGenerator) RandomOrFromPathway(dt *pathway.DateTime, p *pathway.DiagnosisOrProcedure) *ir.DiagnosisOrProcedure {
	t := g.nullTimeOrRandom(dt)
	if p.Code == constants.RandomString || p.Description == constants.RandomString {
		return g.random(t)
	}
	return g.fromPathway(t, p)
}

func (g *DiagOrProcGenerator) random(t ir.NullTime) *ir.DiagnosisOrProcedure {
	return &ir.DiagnosisOrProcedure{
		Description: g.Random(),
		Type:        g.RandomType(),
		DateTime:    t,
	}
}

func (g *DiagOrProcGenerator) fromPathway(t ir.NullTime, p *pathway.DiagnosisOrProcedure) *ir.DiagnosisOrProcedure {
	code, description := g.DeriveCodeAndDescription(p.Code, p.Description)
	return &ir.DiagnosisOrProcedure{
		Description: &ir.CodedElement{ID: code, Text: description},
		Type:        p.Type,
		DateTime:    t,
	}
}
