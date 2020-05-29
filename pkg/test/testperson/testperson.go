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

// Package testperson contains test Person data.
package testperson

import (
	"time"

	"github.com/google/simhospital/pkg/ir"
)

// New returns a Person with all fields populated.
func New() *ir.Person {
	return &ir.Person{
		Prefix:     "Miss",
		FirstName:  "Alice",
		MiddleName: "Mary",
		Surname:    "Johnes",
		Suffix:     "Sr",
		Degree:     "Ph.D.",
		Gender:     "F",
		Ethnicity:  &ir.Ethnicity{ID: "Asian", Text: "Asian"},
		Address: &ir.Address{
			FirstLine:  "999 Small House",
			SecondLine: "Short Street",
			City:       "Croydon",
			PostalCode: "XY9 8ZZ",
			Country:    "UK",
			Type:       "HOME",
		},
		MRN:   "123456",
		NHS:   "0714630667",
		Birth: ir.NewValidTime(time.Date(1998, 1, 20, 0, 0, 0, 0, time.UTC)),
	}
}
