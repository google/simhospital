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

package hospital

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/google/simhospital/pkg/logging"
	"github.com/google/simhospital/pkg/message"
	"github.com/google/simhospital/pkg/state"
)

// HasMessages returns whether there are messages in the Message queue, independently of when they are due.
func (h *Hospital) HasMessages() bool {
	return !h.messageQ.Empty()
}

// processNextMessage consumes the next message from the Message queue and processes it.
// processNextMessage returns an error if the queue is empty or there was any problem processing the message.
func (h *Hospital) processNextMessage() error {
	if !h.HasMessages() {
		return errors.New("processNextMessage() was invoked on an empty Message queue")
	}
	// The item retrieved by pq.Get() might not be exactly the same one as the pq.Peek() above (as the items are
	// submitted to the queue concurrently by another thread). This is OK, because we want to process the item
	// which due date is the earliest first.
	i, err := h.messageQ.Get()
	if err != nil {
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": unknown,
			"reason":       "message_queue_get",
		}).Inc()
		return errors.Wrap(err, "failed to get message from queue")
	}
	item := *i
	m, ok := item.(state.HL7Message)
	if !ok {
		log.WithError(err).Errorf("Unknown item consumed from the message queue: %v", item)
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": unknown,
			"reason":       "message_cast",
		}).Inc()
		// An item of the wrong type will never be able to be processed. The item has been
		// removed from the queue already, so there's no point in signaling this further.
		return nil
	}
	if err := h.processMessage(m); err != nil {
		return errors.Wrap(err, "failed to process message")
	}
	return nil
}

// processMessage processes the message.
// The default processing logic sends the message using the configured sender. This functionality, as well as what
// happens before and after, can be overridden by custom message processors.
// An error at any processing stage results in the rest of the processing stages not being carried out.
// If an error occurs, the ErrorsTotal metric is incremented.
func (h *Hospital) processMessage(m state.HL7Message) error {
	logLocal := log.WithField(keyPathwayName, m.PathwayName).
		WithField(keyMessageName, m.Name).
		WithField(keyIsHistorical, m.IsHistorical).
		WithField(keyExpectedMessageTime, m.MessageTime.UTC().Format(datetimeLayout))

	logLocal = logLocal.WithField(keyMessageType, m.Message.Type)

	if _, err := runMessageProcessors(logLocal, &m, h.processors.MessagePre); err != nil {
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": m.PathwayName,
			"reason":       "message_pre_processor",
		}).Inc()
		return errors.Wrap(err, "message pre processing failed")
	}

	processed, err := runMessageProcessors(logLocal, &m, h.processors.MessageOverride)
	if err != nil {
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": m.PathwayName,
			"reason":       "message_override_processor",
		}).Inc()
		return errors.Wrap(err, "message override processing failed")
	}

	if !processed {
		logLocal.Info("Sending message")
		logLocal.WithField(keyMessage, m).Debug("Sending message")
		if err := h.sender.Send([]byte(m.Message.Message)); err != nil {
			counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
				"pathway_name": m.PathwayName,
				"reason":       "send_message",
			}).Inc()
			return errors.Wrap(err, "cannot send message")
		}
		counters.SimulatedHospital.MessagesTotal.With(prometheus.Labels{
			"pathway_name":  m.PathwayName,
			"message_type":  strings.ToLower(m.Message.Type.MessageType),
			"trigger_event": m.Message.Type.TriggerEvent,
		}).Inc()
		counters.SimulatedHospital.MessageDelaySeconds.Observe(h.clock.Now().UTC().Sub(m.MessageTime.UTC()).Seconds())
	}

	if _, err := runMessageProcessors(logLocal, &m, h.processors.MessagePost); err != nil {
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": m.PathwayName,
			"reason":       "message_post_processor",
		}).Inc()
		return errors.Wrap(err, "message post processing failed")
	}
	return nil
}

func runMessageProcessors(logLocal *logging.SimulatedHospitalLogger, m *state.HL7Message, ps []MessageProcessor) (bool, error) {
	processed := false
	for _, p := range ps {
		if !p.Matches(m) {
			continue
		}
		logLocal.Debugf("Running custom processor for message: %s", m)
		processed = true
		if err := p.Process(m); err != nil {
			return false, err
		}
	}
	return processed, nil
}

func (h *Hospital) queueMessage(logLocal *logging.SimulatedHospitalLogger, msg *message.HL7Message, e *state.Event) error {
	name := fmt.Sprintf("%s^%s-%s", msg.Type.MessageType, msg.Type.TriggerEvent, e.PatientMRN)
	*logLocal = *logLocal.
		WithField(keyMessageName, name).
		WithField(keyMessageType, msg.Type).
		WithField(keyIsHistorical, e.IsHistorical).
		WithField(keyExpectedMessageTime, e.MessageTime.UTC().Format(datetimeLayout))
	logLocal.Info("Queuing message")
	logLocal.WithField(keyMessage, msg).Debug("Queuing message")
	err := h.messageQ.Put(state.HL7Message{
		Name:         name,
		PathwayName:  e.PathwayName,
		Message:      msg,
		MessageTime:  e.MessageTime,
		IsHistorical: e.IsHistorical,
	})
	if err != nil {
		logLocal.WithError(err).Error("Failed to put the message on the priority queue")
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": e.PathwayName,
			"reason":       "Failed to put the message on the priority queue",
		}).Inc()
		return errors.New("failed to put the message on the priority queue")
	}
	return nil
}
