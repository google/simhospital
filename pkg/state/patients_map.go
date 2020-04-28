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
	"sync"

	"github.com/pkg/errors"
	"github.com/google/simhospital/pkg/state/persist"
)

// PatientsMap contains the map of patients, indexed by their ID.
type PatientsMap struct {
	m     map[string]*Patient
	mutex sync.Mutex
	// syncer receives all operations (Put, Get, Delete, etc.) performed on the PatientsMap.
	// Optional.
	syncer      persist.ItemSyncer
	deleteFromM bool
}

// NewPatientsMap returns a map-based store for tracking patients inside Simulated Hospital.
// NewPatientsMap takes an optional ItemSyncer which can be used to persist items into other storage (e.g., a database)
// so that if the execution stops for any reason, the items can be loaded from such storage.
// If the syncer is set, all operations (Put, Get, Delete, etc.) performed on PatientsMap are performed on the syncer too.
// Pass in a nil ItemSyncer if you only wish to use the in-memory map for storing state.
// deleteFromMap indicates whether patients are deleted from the internal map to save memory.
func NewPatientsMap(syncer persist.ItemSyncer, deleteFromMap bool) *PatientsMap {
	return &PatientsMap{
		m:           map[string]*Patient{},
		mutex:       sync.Mutex{},
		syncer:      syncer,
		deleteFromM: deleteFromMap,
	}
}

// Put adds a patient to the internal patients map and the syncer.
func (m *PatientsMap) Put(p *Patient) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.put(p)
	if m.syncer != nil {
		m.syncer.Write(*p)
	}
}

func (m *PatientsMap) put(p *Patient) error {
	id, err := p.ID()
	if err != nil {
		return errors.Wrap(err, "cannot get ID")
	}
	m.m[id] = p
	return nil
}

// Get returns a patient if its identifier is present within the internal patients map or the syncer, in this order.
func (m *PatientsMap) Get(id string) *Patient {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	logLocal := log.WithField("patient_id", id)
	if p, ok := m.m[id]; ok {
		logLocal.Debug("Patient found in the map")
		return p
	}
	if m.syncer == nil {
		logLocal.Debug("Patient not found in the map; syncer not set")
		return nil
	}
	return m.getFromSyncer(id)
}

func (m *PatientsMap) getFromSyncer(id string) *Patient {
	logLocal := log.WithField("patient_id", id)
	item, err := m.syncer.LoadByID(id)
	if err != nil {
		logLocal.WithError(err).Error("Could not get patient from synced storage")
		return nil
	}
	if item == nil {
		logLocal.Debug("Patient not found in map or syncer")
		return nil
	}
	logLocal.Debug("Patient found via syncer")
	p := item.(Patient)
	id, err = p.ID()
	if err != nil {
		logLocal.WithError(err).Error("Found the patient in the synced storage, but could not get the patient ID")
		return nil
	}
	m.m[id] = &p
	return m.m[id]
}

// Delete deletes a patient from the internal patients map and the syncer, by its identifier.
func (m *PatientsMap) Delete(id string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if item, ok := m.m[id]; ok && m.syncer != nil {
		if err := m.syncer.Delete(item); err != nil {
			log.WithField("patient_id", id).WithError(err).Error("Cannot delete patient from the syncer")
		}
	}
	if m.deleteFromM {
		delete(m.m, id)
	}
}

// Len returns the length of the patients map.
func (m *PatientsMap) Len() int {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return len(m.m)
}

// PatientUnmarshaller is an unmarshaller of patients.
type PatientUnmarshaller struct{}

// Unmarshal unmarshals JSON input into a Patient.
func (u *PatientUnmarshaller) Unmarshal(b []byte) (persist.MarshallableItem, error) {
	var p Patient
	err := json.Unmarshal(b, &p)
	return p, err
}
