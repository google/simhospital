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

package generator

import (
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/sample"
)

type patientClassGenerator struct {
	*sample.DiscreteDistribution
}

func newPatientClassAndTypeGenerator(d *config.Data) patientClassGenerator {
	return patientClassGenerator{DiscreteDistribution: &sample.DiscreteDistribution{WeightedValues: d.PatientClass}}
}

func (eg patientClassGenerator) Random() *config.PatientClassAndType {
	e := eg.DiscreteDistribution.Random()
	if e == nil {
		return nil
	}
	return e.(*config.PatientClassAndType)
}
