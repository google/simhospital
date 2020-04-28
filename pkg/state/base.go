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

// Package state contains definitions of stateful objects in Simulated Hospital.
package state

import (
	"crypto/sha256"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/google/simhospital/pkg/logging"
	"github.com/google/simhospital/pkg/monitoring"
)

const (
	// EventItemType used for Event items.
	EventItemType = "event"

	// MessageItemType used for Message items.
	MessageItemType = "message"

	// PatientItemType used for Patient items.
	PatientItemType = "patient"
)

var (
	log = logging.ForCallerPackage()

	counters struct {
		SimulatedHospital struct {
			PendingItem *prometheus.GaugeVec `help:"Number of pending items in the queues" labels:"item_type"`
		}
	}
)

func init() {
	if err := monitoring.CreateAndRegisterMetricsFromStruct(&counters); err != nil {
		log.WithError(err).Fatal("Cannot register metrics from the 'state' package")
	}
}

func id(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return fmt.Sprintf("%X", bs)
}
