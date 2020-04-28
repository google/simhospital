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
	"time"

	"github.com/golang-collections/go-datastructures/queue"
	"github.com/google/simhospital/pkg/pathway"
	"github.com/google/simhospital/pkg/state/persist"
)

// Event is a stateful object representing Simulated Hospital events currently in progress.
type Event struct {
	EventTime      time.Time
	MessageTime    time.Time
	PathwayName    string
	PatientMRN     string
	Step           pathway.Step
	History        []pathway.Step
	Pathway        []pathway.Step
	PathwayStarted time.Time
	IsHistorical   bool
	Index          int
	// PatientIDs is a map from PatientID to MRN; only set if the pathway this event belongs to had a Persons section.
	PatientIDs map[pathway.PatientID]string
}

func (e Event) String() string {
	return fmt.Sprintf("time:%v, messageTime:%v pathwayName:%v, index:%v, mrn:%v", e.EventTime, e.MessageTime, e.PathwayName, e.Index, e.PatientMRN)
}

// ResolveMRN transforms the given PatientID into an MRN.
// If the given PatientID is not in this event's patient map, it is assumed that it is an MRN already.
func (e Event) ResolveMRN(patientID pathway.PatientID) string {
	if mrn, ok := e.PatientIDs[patientID]; ok {
		return mrn
	}
	return string(patientID)
}

// Compare compares the current event with the given one.
// This method is part of the queue.Item interface and allows the events to be added to the priority queue.
// The items on the priority queue need to be sorted by their eventTime date. If two items have exactly
// the same eventTime date, they will be ordered by insertion time, ie: the item inserted earlier
// will be first.
func (e Event) Compare(other queue.Item) int {
	o, ok := other.(Event)
	if !ok {
		log.Fatalf("Incompatible type: trying to compare Event with %v", other)
	}
	if e.EventTime.Unix() > o.EventTime.Unix() {
		return 1
	}
	return -1
}

// ID returns this event's identifier.
func (e Event) ID() (string, error) {
	return id(fmt.Sprintf("%v", e)), nil

}

// Marshal marshals the event for persistence.
func (e Event) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

// EventUnmarshaller can be used to unmarshal events.
type EventUnmarshaller struct{}

// Unmarshal unmarshals events.
func (u EventUnmarshaller) Unmarshal(b []byte) (persist.MarshallableItem, error) {
	var e Event
	err := json.Unmarshal(b, &e)
	return e, err
}
