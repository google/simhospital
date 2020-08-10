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
	"context"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/logging"
	"github.com/google/simhospital/pkg/pathway"
	"github.com/google/simhospital/pkg/processor"
	"github.com/google/simhospital/pkg/state"
)

// HasEvents returns whether there are events in the Event queue, independently of when they are due.
func (h *Hospital) HasEvents() bool {
	return !h.eventQ.Empty()
}

// runNextEvent consumes the next event from the Events queue and runs it.
// runNextEvent returns an error if the queue is empty or there was any problem running the event.
func (h *Hospital) runNextEvent(ctx context.Context) error {
	if !h.HasEvents() {
		return errors.New("runNextEvent() was invoked on an empty Event queue")
	}
	// The item retrieved by pq.Get() might not be exactly the same one as the pq.Peek() above (as the items are
	// submitted to the queue concurrently by another thread). This is OK, because we want to process the item
	// which due date is the earliest first.
	consistentBefore := h.eventQ.IsConsistent()
	i, err := h.eventQ.Get()
	if err != nil {
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": unknown,
			"reason":       "event_queue_get",
		}).Inc()
		return errors.Wrap(err, "failed to get event from queue")
	}
	item := *i
	event, ok := item.(state.Event)
	if !ok {
		log.WithError(err).Errorf("Unknown item consumed from the event queue: %v", item)
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": unknown,
			"reason":       "event_cast",
		}).Inc()
		// An item of the wrong type will never be able to be processed. The item has been
		// removed from the queue already, so there's no point in signaling this further.
		return nil
	}

	h.runEvent(ctx, event)

	if consistentAfter := h.eventQ.IsConsistent(); !consistentAfter && consistentBefore {
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": event.PathwayName,
			"reason":       inconsistentQueueError,
		}).Inc()
	}
	return nil
}

func (h *Hospital) queueFirstEvent(p pathway.Pathway, patientIDs map[pathway.PatientID]string, patients ...*state.Patient) error {
	first, history, steps := getNextEvents(p.History, p.Pathway)

	firstPatient := patients[0]
	logLocal := log.WithField(keyPathwayName, p.Name()).
		WithField(keyEventType, first.StepType()).
		WithField(keyPatientID, firstPatient.PatientInfo.Person.MRN)
	logLocal.Info("Queuing first pathway event")
	logLocal.WithField(keyEvent, first).Debug("Queuing first pathway event")

	mrn := firstPatient.PatientInfo.Person.MRN
	for _, patient := range patients {
		h.patients.Put(patient)
	}

	now := h.clock.Now()

	eventTime, msgTime := calculateTimes(now, first.Parameters)

	consistentBefore := h.eventQ.IsConsistent()
	event := state.Event{
		EventTime:      eventTime,
		MessageTime:    msgTime,
		PathwayName:    p.Name(),
		PatientMRN:     mrn,
		Step:           *first,
		Pathway:        steps,
		History:        history,
		PathwayStarted: now,
		IsHistorical:   len(p.History) > 0,
		Index:          0,
		PatientIDs:     patientIDs,
	}
	if err := h.eventQ.Put(event); err != nil {
		return err
	}
	if consistentAfter := h.eventQ.IsConsistent(); !consistentAfter && consistentBefore {
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": event.PathwayName,
			"reason":       inconsistentQueueError,
		}).Inc()
	}
	return nil
}

// runEvent: the default processing functionality builds the related HL7 msg and adds it
// to the Message queue. This can be overridden with custom processing, and additionally, custom pre/post
// processing logic can run before/after the override or the default processing. Finally, this method queues
// the next event to be run as part of the pathway.
// If there's an error (e.g. event pre/override/default/post processing logic fails, the message cannot be added
// to the queue, etc.), the next event isn't added to the queue, and thus the entire pathway is stopped. In that
// case the patient is deleted from the internal map.
func (h *Hospital) runEvent(ctx context.Context, e state.Event) {
	pathwayName := e.PathwayName
	mrn := e.PatientMRN
	logLocal := log.WithField(keyPathwayName, pathwayName).
		WithField(keyPatientID, mrn)
	patient := h.patients.Get(mrn)
	if patient == nil {
		logLocal.Error("unknown MRN in event")
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": e.PathwayName,
			"reason":       "unknown_mrn",
		}).Inc()
		return
	}
	patientInfo := patient.PatientInfo
	// Advance the current time in the pathway, if needed.
	now := e.EventTime
	if e.IsHistorical {
		// We are loading historical data: the pathway has not started yet and time has not moved.
		// We want the next event time to be set relative to when the pathway starts, not relative to
		// when the current historical event is taking place.
		now = e.PathwayStarted
	}

	if e.Step.StepType() == pathway.StepDelay {
		now = now.Add(e.Step.Delay.Random())
	}

	if _, err := h.runEventProcessors(logLocal, &e, patientInfo, h.processors.EventPre); err != nil {
		logLocal.WithError(err).Error("event pre processing failed")
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": e.PathwayName,
			"reason":       "event_pre_processor",
		}).Inc()
		return
	}

	processed, err := h.runEventProcessors(logLocal, &e, patientInfo, h.processors.EventOverride)
	if err != nil {
		logLocal.WithError(err).Error("event override processing failed")
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": e.PathwayName,
			"reason":       "event_override_processor",
		}).Inc()
		return
	}

	if !processed {
		if err := h.processEventType(ctx, &e, logLocal, now); err != nil {
			logLocal.WithError(err).Errorf("cannot process event type %v, deleting patient", e.Step.StepType())
			counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
				"pathway_name": pathwayName,
				"reason":       err.Error(),
			}).Inc()
			h.patients.Delete(e.PatientMRN)
			return
		}
		// We make sure to persist the data in the internal map into the internal database before deleting it.
		// If we don't do this, then the internal map and internal database get out of sync.
		h.patients.Put(h.patients.Get(e.PatientMRN))
	}

	if _, err := h.runEventProcessors(logLocal, &e, patientInfo, h.processors.EventPost); err != nil {
		logLocal.WithError(err).Error("event post processing failed")
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": e.PathwayName,
			"reason":       "event_post_processor",
		}).Inc()
		return
	}

	// Event processing might have changed the patient's MRN.
	mrn = e.PatientMRN

	// Queue the next event, if any.
	first, history, pathwaySteps := getNextEvents(e.History, e.Pathway)
	if first != nil {
		eventTime, msgTime := calculateTimes(now, first.Parameters)

		logLocal = logLocal.
			WithField(keyNextEventType, first.StepType()).
			WithField(keyExpectedNextEventTime, eventTime.UTC().Format(datetimeLayout))

		logLocal.Info("Queuing next event")
		logLocal.WithField(keyNextEvent, first).Debug("Queuing next event")

		consistentBefore := h.eventQ.IsConsistent()
		event := state.Event{
			EventTime:      eventTime,
			MessageTime:    msgTime,
			PathwayName:    pathwayName,
			PatientMRN:     mrn,
			Step:           *first,
			Pathway:        pathwaySteps,
			History:        history,
			PathwayStarted: e.PathwayStarted,
			IsHistorical:   len(e.History) > 0,
			Index:          e.Index + 1,
			PatientIDs:     e.PatientIDs,
		}
		if err := h.eventQ.Put(event); err != nil {
			logLocal.WithError(err).Error("Failed to put the next event on the priority queue")
			counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
				"pathway_name": pathwayName,
				"reason":       "Failed to put the next event on the priority queue",
			}).Inc()
		}
		if consistentAfter := h.eventQ.IsConsistent(); !consistentAfter && consistentBefore {
			counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
				"pathway_name": pathwayName,
				"reason":       inconsistentQueueError,
			}).Inc()
		}
	} else {
		// The pathway has finished!
		// We assume the pathway steps are sorted in chronological order, so the time of the last event is
		// the time the pathway finishes.
		logLocal.Info("Pathway finished!")
		h.patients.Delete(mrn)
		if len(e.Pathway) == 0 && !e.IsHistorical {
			// The last step is a Pathway step, as opposed to a historical step.
			// Note we don't export the metric if the pathway has historical steps only, as there's no
			// pathway duration in that case.
			// We assume the pathway steps are sorted in chronological order, so the time of the last event is
			// the time the pathway finishes.
			counters.SimulatedHospital.PathwayDurationMinutes.With(prometheus.Labels{
				"pathway_name": pathwayName,
			}).Observe(now.Sub(e.PathwayStarted).Minutes())
		}
	}
}

// getNextEvents gets the first event to be run, either from the historical steps or the pathway (if
// there are no historical steps), and returns the updated lists of historical and pathway steps.
func getNextEvents(historicalSteps []pathway.Step, pathwaySteps []pathway.Step) (first *pathway.Step, history []pathway.Step, steps []pathway.Step) {
	if len(historicalSteps) > 0 {
		first = &historicalSteps[0]
		history = historicalSteps[1:(len(historicalSteps))]
		steps = pathwaySteps
	} else if len(pathwaySteps) > 0 {
		first = &pathwaySteps[0]
		steps = pathwaySteps[1:(len(pathwaySteps))]
	}
	return
}

func (h *Hospital) runEventProcessors(logLocal *logging.SimulatedHospitalLogger, e *state.Event, patientInfo *ir.PatientInfo, ps []EventProcessor) (bool, error) {
	processed := false
	for _, p := range ps {
		if !p.Matches(e) {
			continue
		}
		logLocal.Debugf("Running custom processor for event %s", e.Step.StepType())
		processed = true

		msgs, err := p.Process(e, patientInfo, &processor.Config{
			Generator:       h.generator,
			LocationManager: h.locationManager,
			HL7Config:       h.messageConfig,
		})
		if err != nil {
			return false, err
		}
		for _, m := range msgs {
			if err := h.queueMessage(logLocal, m, e); err != nil {
				return false, err
			}
		}
	}
	return processed, nil
}
