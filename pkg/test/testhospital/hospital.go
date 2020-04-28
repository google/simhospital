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

// Package testhospital contains functionality to test hospitals.
// It allows to create hospitals easily and run the events and messages.
package testhospital

import (
	"testing"
	"time"

	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/doctor"
	"github.com/google/simhospital/pkg/generator/header"
	"github.com/google/simhospital/pkg/hardcoded"
	"github.com/google/simhospital/pkg/hospital"
	"github.com/google/simhospital/pkg/location"
	"github.com/google/simhospital/pkg/orderprofile"
	"github.com/google/simhospital/pkg/pathway"
	"github.com/google/simhospital/pkg/test"
	"github.com/google/simhospital/pkg/test/testaddress"
	"github.com/google/simhospital/pkg/test/testclock"
	"github.com/google/simhospital/pkg/test/testhl7"
	"github.com/google/simhospital/pkg/test/testid"
)

const defaultClockTick = time.Second

// now is an arbitrary time to be used by default.
var now = time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC)

// Hospital represents a hospital that can be used for testing.
// Some of the fields this hospital is created with are exposed so that they can be accessed in the tests.
type Hospital struct {
	*hospital.Hospital
	clock           *testclock.Clock
	Sender          *testhl7.Sender
	Parser          *pathway.Parser
	PathwayManager  pathway.Manager
	LocationManager *location.Manager
	MessageConfig   *config.HL7Config
}

// Config is the configuration of a test hospital.
type Config struct {
	hospital.Config
	test.ConfigFiles
}

// New creates a new Hospital for test.
func New(t *testing.T, cfg Config) *Hospital {
	t.Helper()
	return WithTime(t, cfg, now)
}

// WithTime creates a new Hospital for test initialised with the given time.
func WithTime(t *testing.T, cfg Config, now time.Time) *Hospital {
	t.Helper()

	if cfg.HL7Config == nil {
		c, err := config.LoadHL7Config(cfg.MessageConfigFile)
		if err != nil {
			t.Fatalf("config.LoadHL7Config(%s) failed with %v", cfg.MessageConfigFile, err)
		}
		cfg.HL7Config = c
	}
	if cfg.LocationManager == nil {
		locationManager, err := location.NewManager(cfg.LocationsFile)
		if err != nil {
			t.Fatalf("location.NewManager(%s) failed with %v", cfg.LocationsFile, err)
		}
		cfg.LocationManager = locationManager
	}
	if cfg.Header == nil {
		headerCFG, err := config.LoadHeaderConfig(cfg.HeaderConfigFile)
		if err != nil {
			t.Fatalf("config.LoadHeaderConfig(%s) failed with %v", cfg.HeaderConfigFile, err)
		}
		cfg.Header = headerCFG
	}
	if cfg.DataFiles == (config.DataFiles{}) {
		cfg.DataFiles = test.DataFiles[test.Test]
	}
	if cfg.Doctors == nil {
		d, err := doctor.LoadDoctors(cfg.DoctorFile)
		if err != nil {
			t.Fatalf("doctor.LoadDoctors(%s) failed with %v", cfg.DoctorFile, err)
		}
		cfg.Doctors = d
	}
	if cfg.OrderProfiles == nil {
		op, err := orderprofile.Load(cfg.OrderProfilesFile, cfg.HL7Config)
		if err != nil {
			t.Fatalf("orderprofile.Load(%s, %+v) failed with %v", cfg.OrderProfilesFile, cfg.HL7Config, err)
		}
		cfg.OrderProfiles = op
	}
	if cfg.MessageControlGenerator == nil {
		cfg.MessageControlGenerator = &header.MessageControlGenerator{}
	}
	if cfg.MessagesManager == nil {
		hm, err := hardcoded.NewManager(cfg.HardcodedMessagesConfigDir, cfg.MessageControlGenerator)
		if err != nil {
			t.Fatalf("hardcoded.NewManager(%s) failed with %v", cfg.HardcodedMessagesConfigDir, err)
		}
		cfg.MessagesManager = hm
	}
	cfg.Clock = testclock.New(now)
	pathwayP := &pathway.Parser{Clock: cfg.Clock, OrderProfiles: cfg.OrderProfiles, Doctors: cfg.Doctors, LocationManager: cfg.LocationManager}
	var pathways map[string]pathway.Pathway
	if cfg.PathwayManager == nil {
		var err error
		pathways, err = pathwayP.ParsePathways(cfg.PathwaysDir, nil, nil)
		if err != nil {
			t.Fatalf("ParsePathways(%s, %v, %v) failed with %v", cfg.PathwaysDir, nil, nil, err)
		}
		pathwayManager, err := pathway.NewDistributionManager(pathways)
		if err != nil {
			t.Fatalf("pathway.NewDistributionManager() failed with %v", err)
		}
		cfg.PathwayManager = pathwayManager
	}
	if cfg.Sender == nil {
		cfg.Sender = &testhl7.Sender{}
	}

	cfg.AdditionalConfig.AddressGenerator = &testaddress.ArbitraryGenerator
	cfg.AdditionalConfig.MRNGenerator = &testid.Generator{}
	cfg.AdditionalConfig.PlacerGenerator = &testid.Generator{}
	cfg.AdditionalConfig.FillerGenerator = &testid.Generator{}

	h, err := hospital.NewHospital(cfg.Config)
	if err != nil {
		t.Fatalf("NewHospital() failed with %v", err)
	}
	return &Hospital{
		Hospital:        h,
		clock:           cfg.Clock.(*testclock.Clock),
		Sender:          cfg.Sender.(*testhl7.Sender),
		Parser:          pathwayP,
		PathwayManager:  cfg.PathwayManager,
		LocationManager: cfg.LocationManager,
		MessageConfig:   cfg.HL7Config,
	}
}

// ConsumeQueues consumes all events and messages and returns the number of events that were run
// and the messages that were sent.
// ConsumeQueues fails if processing a message or an event fails.
// See ConsumeQueuesWithLimit for the order in which events and messages are processed.
func (h *Hospital) ConsumeQueues(t *testing.T) (events int, messages []string) {
	t.Helper()
	return h.ConsumeQueuesWithLimit(t, -1, true)
}

// ConsumeQueuesWithLimit consumes limit events and their corresponding messages, and returns the number of events that were run
// and the messages that were sent.
// Parameters:
// * limit is the maximum number of events to run. All messages relevant to those events will be sent.
//   If there are less than limit events in the queue, this method consumes the events in the queue and returns. Set <1 for no limit.
// * 'failIfError' specifies whether errors running events or sending messages should make the method fail.
// The order in which events and messages are processed is the following:
// 1. Run all events that are due at the current time.
// 2. Send all messages that are due at the current time.
// 3. If there were no events due, advance the clock, and process any messages for that time so that we give priority
//    to the existing messages before running new events.
// 4. If there are still events in the queue, go back to step 1.
// 5. After all events are consumed, process all of the remaining messages advancing the clock if needed.
func (h *Hospital) ConsumeQueuesWithLimit(t *testing.T, limit int, failIfError bool) (events int, messages []string) {
	t.Helper()
	for h.HasEvents() {
		ran, err := h.RunNextEventIfDue()
		if err != nil && failIfError {
			t.Fatalf("hospital.RunNextEventIfDue() failed with %v", err)
		}
		messages = append(messages, h.consumeDueMessages(t, failIfError)...)
		if ran {
			events++
		} else {
			h.clock.Advance(defaultClockTick)
			messages = append(messages, h.consumeDueMessages(t, failIfError)...)
		}
		if limit > 0 && limit >= events {
			return
		}
	}
	messages = append(messages, h.consumeAllMessages(t, failIfError)...)
	return
}

func (h *Hospital) consumeDueMessages(t *testing.T, failIfError bool) []string {
	t.Helper()
	var messages []string
	for h.HasMessages() {
		processed, err := h.ProcessNextMessageIfDue()
		if err != nil && failIfError {
			t.Fatalf("hospital.ProcessNextMessageIfDue() failed with %v", err)
		}
		if !processed {
			// There are more messages, but they're not due yet.
			return messages
		}
		messages = append(messages, h.Sender.GetSentMessages()...)
		h.Sender.EraseSentHistory()
	}
	return messages
}

func (h *Hospital) consumeAllMessages(t *testing.T, failIfError bool) []string {
	t.Helper()
	var messages []string
	for h.HasMessages() {
		messages = append(messages, h.consumeDueMessages(t, failIfError)...)
		h.clock.Advance(defaultClockTick)
	}
	return messages
}
