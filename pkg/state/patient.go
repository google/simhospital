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

package state

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/google/simhospital/pkg/message"
)

const generatedOrderIDPattern = "generated-%d"

// Patient represents a patient in Simulated Hospital.
type Patient struct {
	PatientInfo *message.PatientInfo
	Orders      map[string]*message.Order
	PastVisits  []uint64
}

// GetOrder retrieves an order by its identifier.
func (p *Patient) GetOrder(orderID string) *message.Order {
	return p.Orders[orderID]
}

// AddOrder adds an order to the map against the specified order ID, so that it can be looked up later on.
// If the orderID is not specified (ie is an empty string), a unique ID is generated.
func (p *Patient) AddOrder(orderID string, order *message.Order) {
	if orderID == "" {
		p.Orders[fmt.Sprintf(generatedOrderIDPattern, len(p.Orders))] = order
	} else {
		p.Orders[orderID] = order
	}
}

// PushPastVisit appends a visit number to the patients PastVisits slice.
func (p *Patient) PushPastVisit(visit uint64) {
	p.PastVisits = append(p.PastVisits, visit)
}

// PopPastVisit gets and deletes the most recent past visit.
// PopPastVisit returns an error if there are no past visits.
func (p *Patient) PopPastVisit() (uint64, error) {
	if len(p.PastVisits) == 0 {
		return 0, errors.New("No past visits")
	}
	idx := len(p.PastVisits) - 1
	visit := p.PastVisits[idx]
	p.PastVisits = p.PastVisits[:idx]
	return visit, nil
}

// Marshal marshals a Patient into JSON.
func (p Patient) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

// ID returns the patient's ID.
func (p Patient) ID() (string, error) {
	if p.PatientInfo == nil || p.PatientInfo.Person == nil {
		return "", errors.New("cannot get ID: No PatientInfo or PatientInfo.Person")
	}
	return p.PatientInfo.Person.MRN, nil
}
