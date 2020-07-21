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

// Package location provides functionality to manage locations.
package location

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/yaml.v2"
	"github.com/google/simhospital/pkg/files"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/logging"
	"github.com/google/simhospital/pkg/monitoring"
)

const aAndEID = "ED"

var (
	unknownLocation = "unknown location"
	counters        struct {
		SimulatedHospital struct {
			OccupiedBeds *prometheus.GaugeVec `help:"Number of occupied beds" labels:"poc"`
		}
	}
	log = logging.ForCallerPackage()
)

// Manager is a manager of locations that contains multiple room managers.
type Manager struct {
	RoomManagers map[string]*RoomManager
}

// RoomManager is a manager of rooms.
type RoomManager struct {
	Poc      string
	Facility string
	Building string
	Floor    string
	Room     string
	Type     string
	// occupiedBeds is a counter of occupied beds for each point of care.
	// This counter is modified through OccupyAvailableBed and FreeBed.
	occupiedBeds int
	// isBedOccupied is the set of bed names that are currently occupied.
	// If true, bed is occupied; if false, bed is not occupied.
	isBedOccupied map[string]bool
}

func init() {
	if err := monitoring.CreateAndRegisterMetricsFromStruct(&counters); err != nil {
		log.WithError(err).Fatal("Cannot register metrics from the 'location' package")
	}
}

// NewManager returns a location Manager.
func NewManager(ctx context.Context, fileName string) (*Manager, error) {
	roomManagers := map[string]*RoomManager{}

	data, err := files.Read(ctx, fileName)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot parse locations file %s", fileName)
	}

	if err = yaml.Unmarshal(data, &roomManagers); err != nil {
		return nil, errors.Wrapf(err, "cannot unmarshal locations from file %s", fileName)
	}

	log.WithField("file", fileName).Infof("Found %d locations", len(roomManagers))
	for n, rm := range roomManagers {
		t := rm.Type
		if t == "" {
			t = "BED"
		}
		roomManagers[n] = &RoomManager{
			Poc:           rm.Poc,
			Facility:      rm.Facility,
			Building:      rm.Building,
			Floor:         rm.Floor,
			Room:          rm.Room,
			Type:          t,
			occupiedBeds:  0,
			isBedOccupied: make(map[string]bool),
		}
		log.Infof(" - id: %s, poc: %s", n, rm.Poc)
	}
	if _, ok := roomManagers[aAndEID]; !ok {
		return nil, fmt.Errorf("no ED Location found, this is a required Location. File: %s", fileName)
	}
	return &Manager{RoomManagers: roomManagers}, nil
}

// GetAAndELocation returns the ED location.
func (m *Manager) GetAAndELocation() *ir.PatientLocation {
	roomManager := m.RoomManagers[aAndEID]
	if roomManager == nil {
		return &ir.PatientLocation{LocationType: "ED"}
	}
	return &ir.PatientLocation{
		Poc:          roomManager.Poc,
		Facility:     roomManager.Facility,
		LocationType: roomManager.Type,
		Building:     roomManager.Building,
	}
}

// OccupyAvailableBed picks an available bed in the given location and occupies it.
// This would be equivalent to assigning a patient to a bed.
// Returns an error if the location doesn't exist.
func (m *Manager) OccupyAvailableBed(locationName string) (*ir.PatientLocation, error) {
	if _, ok := m.RoomManagers[locationName]; !ok {
		return nil, fmt.Errorf("%s: %s", unknownLocation, locationName)
	}

	// Bed names start in Bed 1 as that's how most humans count.
	for i := 1; ; i++ {
		bedName := fmt.Sprintf("Bed %d", i)
		if location, err := m.OccupySpecificBed(locationName, bedName); err == nil {
			return location, nil
		}
	}
}

// OccupySpecificBed occupies the given bed in the given location.
// Returns an error if the location doesn't exist, or the bed is already occupied.
func (m *Manager) OccupySpecificBed(locationName string, bedName string) (*ir.PatientLocation, error) {
	roomManager, ok := m.RoomManagers[locationName]
	if !ok {
		return nil, fmt.Errorf("%s: %s", unknownLocation, locationName)
	}
	if roomManager.isBedOccupied[bedName] {
		return nil, fmt.Errorf("bed %q in location %q already occupied", bedName, locationName)
	}
	roomManager.isBedOccupied[bedName] = true
	roomManager.occupiedBeds++
	counters.SimulatedHospital.OccupiedBeds.With(prometheus.Labels{
		"poc": locationName,
	}).Set(float64(roomManager.occupiedBeds))
	log.Debugf("Occupied beds in %s: %d", locationName, roomManager.occupiedBeds)
	return &ir.PatientLocation{
		Poc:          roomManager.Poc,
		Room:         roomManager.Room,
		Bed:          bedName,
		Facility:     roomManager.Facility,
		LocationType: roomManager.Type,
		Building:     roomManager.Building,
		Floor:        roomManager.Floor,
	}, nil
}

// FreeBed marks the bed given in the patient location as available.
// It returns an error if the patient location is not a bed.
func (m *Manager) FreeBed(pl *ir.PatientLocation) error {
	if !IsBed(pl) {
		return fmt.Errorf("patient location %+v is not a bed", pl)
	}
	for key, roomManager := range m.RoomManagers {
		matches, err := m.Matches(key, pl)
		if err != nil {
			return errors.Wrapf(err, "finding matching location for %+v failed", pl)
		}
		if !matches {
			continue
		}
		if !roomManager.isBedOccupied[pl.Bed] {
			// This is not a huge deal, but it shouldn't happen.
			// We return it as an error so that we can monitor it.
			return fmt.Errorf("patient location %+v already free", pl)
		}
		roomManager.isBedOccupied[pl.Bed] = false
		roomManager.occupiedBeds--
		counters.SimulatedHospital.OccupiedBeds.With(prometheus.Labels{
			"poc": pl.Poc,
		}).Set(float64(roomManager.occupiedBeds))
		log.Debugf("Occupied beds in %s: %d", pl.Poc, roomManager.occupiedBeds)
		return nil
	}
	return fmt.Errorf("%s: %+v", unknownLocation, pl)
}

// Matches returns whether the location name matches the patient location.
// If pl is nil, this method returns an error.
func (m *Manager) Matches(locationName string, pl *ir.PatientLocation) (bool, error) {
	if pl == nil {
		return false, errors.New("nil patient location")
	}
	roomManager, ok := m.RoomManagers[locationName]
	if !ok {
		return false, fmt.Errorf("%s: %s", unknownLocation, locationName)
	}
	return roomManager.equalToPatientLocation(pl), nil
}

// OccupiedBeds returns the number of beds that are currently occupied.
func (r *RoomManager) OccupiedBeds() int {
	return r.occupiedBeds
}

func (r *RoomManager) equalToPatientLocation(pl *ir.PatientLocation) bool {
	return r.Poc == pl.Poc && r.Facility == pl.Facility &&
		r.Building == pl.Building && r.Floor == pl.Floor && r.Room == pl.Room
}

// IsBed returns whether the given location is a bed.
func IsBed(pl *ir.PatientLocation) bool {
	return pl.Bed != ""
}
