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
	"context"
	"testing"
	"time"

	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/hospital"
	"github.com/google/simhospital/pkg/location"
	"github.com/google/simhospital/pkg/pathway"
	"github.com/google/simhospital/pkg/test"
	"github.com/google/simhospital/pkg/test/testaddress"
	"github.com/google/simhospital/pkg/test/testclock"
	"github.com/google/simhospital/pkg/test/testhl7"
	"github.com/google/simhospital/pkg/test/testid"
	"github.com/google/simhospital/pkg/test/testresource"
)

const defaultClockTick = time.Second

var (
	// now is an arbitrary time to be used by default.
	now           = time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC)
	dataFilesTest = test.DataFiles[test.Test]

	// Arguments contains default arguments for testing.
	Arguments = hospital.Arguments{
		DoctorsFile:          &test.DoctorsConfigTest,
		OrderProfilesFile:    &test.OrderProfilesConfigTest,
		PathwayArguments:     &hospital.PathwayArguments{Dir: test.PathwaysDirTest, Type: "distribution"},
		Hl7ConfigFile:        &test.MessageConfigTest,
		HeaderConfigFile:     &test.HeaderConfigTest,
		HardcodedMessagesDir: &test.HardcodedMessagesDirTest,
		LocationsFile:        &test.LocationsConfigTest,
		DataFiles:            &dataFilesTest,
	}
)

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
// Fields common to both Config and Arguments are taken from hospital.Config.
type Config struct {
	hospital.Config
	hospital.Arguments
}

// New creates a new Hospital for test.
func New(ctx context.Context, t *testing.T, cfg Config) *Hospital {
	t.Helper()
	return WithTime(ctx, t, cfg, now)
}

// WithTime creates a new Hospital for test initialised with the given time.
func WithTime(ctx context.Context, t *testing.T, cfg Config, now time.Time) *Hospital {
	t.Helper()

	clock := testclock.New(now)

	// Create a default config using Arguments, and then override with the fields set in Config.
	// Fields common to both Config and Arguments are taken from hospital.Config; copy them.
	cfg.Arguments.MessageControlGenerator = cfg.Config.MessageControlGenerator
	cfg.Arguments.Clock = clock
	cfg.Arguments.DeletePatientsFromMemory = cfg.Config.DeletePatientsFromMemory
	if cfg.Config.DataFiles != (config.DataFiles{}) {
		cfg.Arguments.DataFiles = &cfg.Config.DataFiles
	}
	c, err := hospital.DefaultConfig(ctx, cfg.Arguments)
	if err != nil {
		t.Fatalf("hospital.DefaultConfig(%+v) failed with %v", cfg.Arguments, err)
	}

	if cfg.LocationManager != nil {
		c.LocationManager = cfg.LocationManager
	}
	if cfg.MessagesManager != nil {
		c.MessagesManager = cfg.MessagesManager
	}
	if cfg.HL7Config != nil {
		c.HL7Config = cfg.HL7Config
	}
	if cfg.Header != nil {
		c.Header = cfg.Header
	}
	if cfg.Doctors != nil {
		c.Doctors = cfg.Doctors
	}
	if cfg.OrderProfiles != nil {
		c.OrderProfiles = cfg.OrderProfiles
	}
	if cfg.PathwayManager != nil {
		c.PathwayManager = cfg.PathwayManager
	}
	if cfg.Sender != nil {
		c.Sender = cfg.Sender
	} else {
		c.Sender = &testhl7.Sender{}
	}
	if cfg.ResourceWriter != nil {
		c.ResourceWriter = cfg.ResourceWriter
	} else {
		c.ResourceWriter = testresource.NewWriter()
	}

	c.AdditionalConfig = cfg.AdditionalConfig
	c.AdditionalConfig.AddressGenerator = &testaddress.ArbitraryGenerator
	c.AdditionalConfig.MRNGenerator = &testid.Generator{}
	c.AdditionalConfig.PlacerGenerator = &testid.Generator{}
	c.AdditionalConfig.FillerGenerator = &testid.Generator{}

	h, err := hospital.NewHospital(ctx, c)
	if err != nil {
		t.Fatalf("NewHospital(%+v) failed with %v", c, err)
	}
	return &Hospital{
		Hospital:        h,
		clock:           c.Clock.(*testclock.Clock),
		Sender:          c.Sender.(*testhl7.Sender),
		Parser:          c.PathwayParser,
		PathwayManager:  c.PathwayManager,
		LocationManager: c.LocationManager,
		MessageConfig:   c.HL7Config,
	}
}

// ConsumeQueues consumes all events and messages and returns the number of events that were run
// and the messages that were sent.
// ConsumeQueues fails if processing a message or an event fails.
// See ConsumeQueuesWithLimit for the order in which events and messages are processed.
func (h *Hospital) ConsumeQueues(ctx context.Context, t *testing.T) (events int, messages []string) {
	t.Helper()
	return h.ConsumeQueuesWithLimit(ctx, t, -1, true)
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
func (h *Hospital) ConsumeQueuesWithLimit(ctx context.Context, t *testing.T, limit int, failIfError bool) (events int, messages []string) {
	t.Helper()
	for h.HasEvents() {
		ran, err := h.RunNextEventIfDue(ctx)
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
