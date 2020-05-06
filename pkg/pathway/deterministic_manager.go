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

package pathway

import (
	"fmt"

	"github.com/pkg/errors"
)

// DeterministicManager is a pathway manager that returns pathways in a fixed order, specified in the order field.
// After the last pathway runs, it starts from the beginning.
type DeterministicManager struct {
	// Collection contains the available pathways and methods to access them.
	Collection
	// order contains pathway names in the order in which they are run.
	order []string
	// nextIdx is the index of the next pathway to be run from the order field.
	nextIdx int
}

// NextPathway returns the next pathway to run based on the order specified in the manager.
func (m *DeterministicManager) NextPathway() (*Pathway, error) {
	name := m.order[m.nextIdx]
	m.nextIdx = (m.nextIdx + 1) % len(m.order)
	return m.GetPathway(name)
}

// NewDeterministicManager creates a new DeterministicManager with the given pathway map and the order in which pathways will be run.
// All pathways are initialised.
// The order slice must have at least one element, and all elements must correspond to existing pathways.
// Otherwise NewDeterministicManager returns an error.
func NewDeterministicManager(pathways map[string]Pathway, order []string) (*DeterministicManager, error) {
	collection, err := NewCollection(pathways)
	if err != nil {
		return nil, err
	}
	if len(order) == 0 {
		return nil, errors.New("the order slice has no elements")
	}
	for _, p := range order {
		if _, err := collection.GetPathway(p); err != nil {
			return nil, fmt.Errorf("pathway %q is unknown", p)
		}
	}
	m := DeterministicManager{
		Collection: collection,
		order:      order,
	}
	m.print()
	return &m, nil
}

func (m DeterministicManager) print() {
	m.Collection.Print(nil)
	log.Infof("Pathways will be run in the following order: %v", m.order)
}
