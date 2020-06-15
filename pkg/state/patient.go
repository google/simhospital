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
	"github.com/google/simhospital/pkg/ir"
)

const generatedIDPattern = "generated-%d"

// Patient represents a patient in Simulated Hospital.
type Patient struct {
	PatientInfo *ir.PatientInfo
	// Orders maps from orderIDs to Orders. This is used to provide an index to all of a Patient's
	// orders so they can be looked up later on. In particular, this is used to link Results events
	// to their corresponding Order events.
	Orders     map[string]*ir.Order
	PastVisits []uint64
	Documents  map[string]*ir.Document
}

// GetOrder retrieves an order by its identifier.
func (p *Patient) GetOrder(orderID string) *ir.Order {
	return p.Orders[orderID]
}

// AddOrder adds an order to the map against the specified order ID if it does not exist, and adds
// it to the current Encounter. If the orderID is not specified (ie is an empty string), a unique
// ID is generated.
func (p *Patient) AddOrder(orderID string, order *ir.Order) {
	if orderID == "" {
		orderID = fmt.Sprintf(generatedIDPattern, len(p.Orders))
	}
	if _, ok := p.Orders[orderID]; !ok {
		p.PatientInfo.AddOrderToEncounter(order)
		p.Orders[orderID] = order
	}
}

// GetDocument retrieves an order by the pathway Document ID.
func (p *Patient) GetDocument(pathwayDocumentID string) *ir.Document {
	return p.Documents[pathwayDocumentID]
}

// AddDocument adds a document to the map against the specified pathway Document ID, so that it can be looked up and updated.
// If the pathwayDocumentID is not specified, a unique ID is generated.
func (p *Patient) AddDocument(pathwayDocumentID string, document *ir.Document) {
	if pathwayDocumentID == "" {
		p.Documents[fmt.Sprintf(generatedIDPattern, len(p.Documents))] = document
	} else {
		p.Documents[pathwayDocumentID] = document
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
