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
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/google/simhospital/pkg/constants"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/location"
	"github.com/google/simhospital/pkg/logging"
	"github.com/google/simhospital/pkg/message"
	"github.com/google/simhospital/pkg/pathway"
	"github.com/google/simhospital/pkg/state"
)

const locationError = "patient location error"

func (h *Hospital) occupyBed(loc, bed string) (*ir.PatientLocation, error) {
	if bed != "" {
		return h.locationManager.OccupySpecificBed(loc, bed)
	}
	return h.locationManager.OccupyAvailableBed(loc)
}

func (h *Hospital) processAdmission(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo

	*logLocal = *logLocal.WithField(keyLocation, e.Step.Admission.Loc)
	if patientInfo.ExpectedAdmitDateTime.Valid {
		patientInfo.AdmissionDate = patientInfo.ExpectedAdmitDateTime
		patientInfo.Location = patientInfo.PendingLocation
		logLocal.Debugf("Entered reserved bed %s", patientInfo.PendingLocation)
	} else {
		patientInfo.AdmissionDate = ir.NewValidTime(e.EventTime)
		loc, err := h.occupyBed(e.Step.Admission.Loc, e.Step.Admission.Bed)
		if err != nil {
			return errors.Wrap(err, locationError)
		}
		patientInfo.Location = loc
	}

	if ec := patientInfo.LatestEncounter(); ec != nil && ec.IsPending {
		ec.UpdateStatus(patientInfo.AdmissionDate, constants.EncounterStatusArrived)
		ec.UpdateLocation(patientInfo.AdmissionDate, patientInfo.Location)
		ec.IsPending = false
	} else {
		patientInfo.AddEncounter(patientInfo.AdmissionDate, constants.EncounterStatusArrived, patientInfo.Location)
	}

	patientInfo.AdmitReason = e.Step.Admission.AdmitReason
	patientInfo.PendingLocation = nil
	patientInfo.ExpectedAdmitDateTime = ir.NewInvalidTime()
	patientInfo.Class = h.messageConfig.PatientClass.Inpatient
	patientInfo.VisitID = h.generator.NewVisitID()
	patientInfo.AccountStatus = h.messageConfig.PatientAccountStatus.Arrived
	h.generator.AddAllergies(patientInfo, e.Step.Admission.Allergies)
	h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)

	msg, err := message.BuildAdmissionADTA01(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A01 message")
	}
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) processOrder(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	patient := h.patients.Get(e.PatientMRN)
	patientInfo := patient.PatientInfo

	h.setAdmissionDetailsIfMissing(patientInfo, e.EventTime)

	o := patient.GetOrder(e.Step.Order.OrderID)
	if o == nil {
		o = h.generator.NewOrder(e.Step.Order, e.EventTime)
		o.MessageControlIDOriginalOrder = msgHeader.MessageControlID
		patient.AddOrder(e.Step.Order.OrderID, o)
		h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)
	}
	if orderStatus := e.Step.Order.OrderStatus; orderStatus != "" {
		o.OrderStatus = orderStatus
	}

	msg, err := message.BuildOrderORMO01(msgHeader, patientInfo, o, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ORM^O01 message")
	}
	if err := h.queueMessage(logLocal, msg, e); err != nil {
		return err
	}
	if e.Step.Order.NoAcknowledgementMessage {
		return nil
	}
	msgHeader = h.generator.NewHeader(&e.Step)
	o.OrderControl = h.messageConfig.OrderControl.OK
	delay := h.orderAckDelay.Random()
	orderAckMessageTime := e.MessageTime.Add(delay)
	msg, err = message.BuildPathologyORRO02(msgHeader, patientInfo, o, orderAckMessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ORR^O02 message")
	}
	e.MessageTime = orderAckMessageTime
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) processResults(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	patient := h.patients.Get(e.PatientMRN)
	patientInfo := patient.PatientInfo

	h.setAdmissionDetailsIfMissing(patientInfo, e.EventTime)
	o, err := h.generator.SetResults(patient.GetOrder(e.Step.Result.OrderID), e.Step.Result, e.EventTime)
	if err != nil {
		return errors.Wrap(err, "cannot set results in Results event")
	}
	o.OrderControl = h.messageConfig.OrderControl.WithObservations
	patient.AddOrder(e.Step.Result.OrderID, o)
	h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)

	te := strings.ToUpper(e.Step.Result.TriggerEvent)
	var msg *message.HL7Message
	switch te {
	case constants.R03:
		msg, err = message.BuildResultORUR03(msgHeader, patientInfo, o, e.MessageTime)
	case constants.R32:
		msg, err = message.BuildResultORUR32(msgHeader, patientInfo, o, e.MessageTime)
	default:
		msg, err = message.BuildResultORUR01(msgHeader, patientInfo, o, e.MessageTime)
	}
	if err != nil {
		return errors.Wrapf(err, "cannot build ORU message; trigger event is %s", te)
	}
	if !e.Step.Result.ExpectCorrection {
		o.NumberOfPreviousResults += len(o.Results)
	}
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) processClinicalNote(ctx context.Context, e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	patient := h.patients.Get(e.PatientMRN)
	patientInfo := patient.PatientInfo
	h.setAdmissionDetailsIfMissing(patientInfo, e.EventTime)

	o := patient.GetOrder(e.Step.ClinicalNote.DocumentID)
	o, err := h.generator.OrderWithClinicalNote(ctx, o, e.Step.ClinicalNote, e.EventTime)
	if err != nil {
		return errors.Wrap(err, "cannot generate a Clinical Note")
	}
	patient.AddOrder(e.Step.ClinicalNote.DocumentID, o)
	msg, err := message.BuildResultORUR01(msgHeader, patientInfo, o, e.MessageTime)
	if err != nil {
		return errors.Wrapf(err, "cannot build ORU^R01 message")
	}
	o.NumberOfPreviousResults += len(o.Results)
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) processDocument(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	patient := h.patients.Get(e.PatientMRN)
	patientInfo := patient.PatientInfo
	updateType := e.Step.Document.UpdateType
	var d *ir.Document

	if updateType != "" {
		d = patient.GetDocument(e.Step.Document.ID)
		if d == nil {
			log.WithError(fmt.Errorf("update type, %q, was requested for document ID, %q, but document does not exist", updateType, e.Step.Document.ID))
			return fmt.Errorf("Document.ID does not exist")
		}
		if err := h.generator.UpdateDocumentContent(d, e.Step.Document); err != nil {
			return errors.Wrap(err, "cannot update document")
		}
	} else {
		if _, ok := patient.Documents[e.Step.Document.ID]; ok {
			log.WithError(fmt.Errorf("ID was set to %q, but document ID already exists", e.Step.Document.ID))
			return fmt.Errorf("Document.ID already exists")
		}
		d = h.generator.NewDocument(e.EventTime, e.Step.Document)
		patient.AddDocument(e.Step.Document.ID, d)
	}

	msg, err := message.BuildDocumentNotificationMDMT02(msgHeader, patientInfo, d, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build MDM^T02 message")
	}
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) processDischarge(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	mrn := e.PatientMRN
	patient := h.patients.Get(mrn)
	patientInfo := patient.PatientInfo
	pathwayName := e.PathwayName

	dischargeTime := e.EventTime
	if e.Step.Discharge.DischargeTime != nil {
		dischargeTime = *e.Step.Discharge.DischargeTime
	}
	setDischargeDate(patientInfo, dischargeTime)
	patientInfo.AccountStatus = h.messageConfig.PatientAccountStatus.Finished
	h.generator.AddAllergies(patientInfo, e.Step.Discharge.Allergies)
	h.updateDeathInfo(logLocal, now, pathwayName, patientInfo, e.Step.Parameters)

	ec := patientInfo.LatestEncounter()
	if ec == nil {
		// No Encounters exist, so we just create and end one here.
		ec = patientInfo.AddEncounter(patientInfo.DischargeDate, constants.EncounterStatusInProgress, patientInfo.Location)
	}
	ec.EndEncounter(patientInfo.DischargeDate, constants.EncounterStatusFinished)
	ec.IsPending = false
	msg, err := message.BuildDischargeADTA03(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A03 message")
	}
	if patientInfo.AdmissionDate.Valid {
		counters.SimulatedHospital.AdmissionDurationMinutes.With(prometheus.Labels{
			"pathway_name": pathwayName,
		}).Observe(e.EventTime.Sub(patientInfo.AdmissionDate.Time).Minutes())
	}
	if patientInfo.Location != nil {
		*logLocal = *logLocal.WithField(keyLocation, patientInfo.Location.Poc)
	}
	patient.PushPastVisit(patientInfo.VisitID)
	patient = h.resetPatient(logLocal, pathwayName, patient, mrn)
	h.patients.Put(patient)
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) processDischargeInError(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo

	dischargeTime := e.EventTime
	if e.Step.DischargeInError.DischargeTime != nil {
		dischargeTime = *e.Step.DischargeInError.DischargeTime
	}
	setDischargeDate(patientInfo, dischargeTime)
	patientInfo.AccountStatus = h.messageConfig.PatientAccountStatus.Finished
	h.generator.AddAllergies(patientInfo, e.Step.DischargeInError.Allergies)
	h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)

	msg, err := message.BuildDischargeADTA03(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A03 message")
	}
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) processTransferOrTransferInError(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo
	pathwayName := e.PathwayName
	var loc string
	var bed string
	eventTime := e.EventTime
	if e.Step.StepType() == pathway.StepTransfer {
		loc = e.Step.Transfer.Loc
		bed = e.Step.Transfer.Bed
		patientInfo.PriorLocation = h.freeLocation(logLocal, patientInfo, pathwayName)
	} else if e.Step.StepType() == pathway.StepTransferInError {
		loc = e.Step.TransferInError.Loc
		bed = e.Step.TransferInError.Bed
		patientInfo.PriorLocation = patientInfo.Location
		patientInfo.PriorLocationForCancelTransfer = patientInfo.PriorLocation
	}
	*logLocal = *logLocal.WithField(keyLocation, loc)

	if patientInfo.ExpectedTransferDateTime.Valid {
		patientInfo.TransferDate = patientInfo.ExpectedTransferDateTime
		eventTime = patientInfo.TransferDate.Time
		patientInfo.Location = patientInfo.PendingLocation
		logLocal.Debugf("Entered reserved bed %s", patientInfo.PendingLocation)
	} else {
		patientInfo.TransferDate = ir.NewValidTime(e.EventTime)
		// Even if this transfer is in error, we simulate the new bed being allocated to the new
		// patient, as that's what the system will think has happened.
		// The old bed is not released, as the patient is still physically there, thus can't be
		// physically given to a new patient.
		loc, err := h.occupyBed(loc, bed)
		if err != nil {
			// It's safe to call freeSpecificLocation even if the prior location was already freed.
			h.freeSpecificLocation(logLocal, patientInfo.PriorLocation, pathwayName)
			return errors.Wrap(err, locationError)
		}
		patientInfo.Location = loc
	}
	patientInfo.PendingLocation = nil
	patientInfo.ExpectedTransferDateTime = ir.NewInvalidTime()
	h.updateDeathInfo(logLocal, now, pathwayName, patientInfo, e.Step.Parameters)

	if ec := patientInfo.LatestEncounter(); ec != nil {
		ec.UpdateStatus(patientInfo.TransferDate, constants.EncounterStatusArrived)
		ec.UpdateLocation(patientInfo.TransferDate, patientInfo.Location)
		ec.IsPending = false
	} else {
		patientInfo.AddEncounter(patientInfo.TransferDate, constants.EncounterStatusArrived, patientInfo.Location)
	}

	msg, err := message.BuildTransferADTA02(msgHeader, patientInfo, eventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A02 message")
	}
	patientInfo.PriorLocation = nil
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) cancelVisit(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	mrn := e.PatientMRN
	patient := h.patients.Get(mrn)
	patientInfo := patient.PatientInfo
	pathwayName := e.PathwayName
	h.updateDeathInfo(logLocal, now, pathwayName, patientInfo, e.Step.Parameters)
	patientInfo.AccountStatus = h.messageConfig.PatientAccountStatus.Cancelled
	msg, err := message.BuildCancelVisitADTA11(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A11 message")
	}
	if err := h.queueMessage(logLocal, msg, e); err != nil {
		return err
	}
	patient.PushPastVisit(patientInfo.VisitID)
	patient = h.resetPatient(logLocal, pathwayName, patient, mrn)
	h.patients.Put(patient)
	return nil
}

func (h *Hospital) cancelTransfer(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	mrn := e.PatientMRN
	patient := h.patients.Get(mrn)
	patientInfo := patient.PatientInfo
	pathwayName := e.PathwayName
	patientInfo.Location, patientInfo.PriorLocation = patientInfo.PriorLocationForCancelTransfer, h.freeLocation(logLocal, patientInfo, pathwayName)
	h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)

	msg, err := message.BuildCancelTransferADTA12(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A12 message")
	}
	patientInfo.TransferDate = ir.NewInvalidTime()
	patientInfo.PriorLocation = nil
	patientInfo.PriorLocationForCancelTransfer = nil
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) cancelDischarge(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo
	h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)
	patientInfo.AccountStatus = h.messageConfig.PatientAccountStatus.Arrived
	msg, err := message.BuildCancelDischargeADTA13(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A13 message")
	}
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) pendingAdmission(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo
	pendingLocation, err := h.occupyBed(e.Step.PendingAdmission.Loc, e.Step.PendingAdmission.Bed)
	if err != nil {
		return errors.Wrap(err, locationError)
	}
	patientInfo.AccountStatus = h.messageConfig.PatientAccountStatus.Planned
	patientInfo.PendingLocation = pendingLocation
	patientInfo.ExpectedAdmitDateTime = ir.NewValidTime(e.EventTime.Add(*e.Step.PendingAdmission.ExpectedAdmissionTimeFromNow))
	h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)

	// A planned Encounter has its start time set to the time that it is expected to begin.
	// In the case that there are two consecutive PendingEncounter steps, the first Encounter
	// will never be finished, since only the latest Encounter is checked.
	ec := patientInfo.AddEncounter(patientInfo.ExpectedAdmitDateTime, constants.EncounterStatusPlanned, nil)
	ec.IsPending = true

	msg, err := message.BuildPendingAdmissionADTA14(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A14 message")
	}
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) pendingDischarge(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo
	patientInfo.ExpectedDischargeDateTime = ir.NewValidTime(e.EventTime.Add(*e.Step.PendingDischarge.ExpectedDischargeTimeFromNow))
	h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)

	ec := patientInfo.LatestEncounter()
	if ec != nil {
		ec.UpdateStatus(patientInfo.ExpectedDischargeDateTime, constants.EncounterStatusPlanned)
	} else {
		ec = patientInfo.AddEncounter(patientInfo.ExpectedDischargeDateTime, constants.EncounterStatusPlanned, patientInfo.Location)
	}
	ec.IsPending = true

	msg, err := message.BuildPendingDischargeADTA16(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A16 message")
	}
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) pendingTransfer(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo
	pendingLocation, err := h.occupyBed(e.Step.PendingTransfer.Loc, e.Step.PendingTransfer.Bed)
	if err != nil {
		return errors.Wrap(err, locationError)
	}
	patientInfo.PendingLocation = pendingLocation
	patientInfo.ExpectedTransferDateTime = ir.NewValidTime(e.EventTime.Add(*e.Step.PendingTransfer.ExpectedTransferTimeFromNow))
	h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)

	ec := patientInfo.AddEncounter(patientInfo.ExpectedTransferDateTime, constants.EncounterStatusPlanned, nil)
	ec.IsPending = true

	msg, err := message.BuildPendingTransferADTA15(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A15 message")
	}
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) registration(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo
	patientInfo.AdmissionDate = ir.NewValidTime(e.EventTime)
	patientInfo.VisitID = h.generator.NewVisitID()
	h.generator.AddAllergies(patientInfo, e.Step.Registration.Allergies)
	h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)
	patientInfo.AccountStatus = h.messageConfig.PatientAccountStatus.Planned

	if e.Step.Registration.PatientClass != "" {
		patientInfo.Class = e.Step.Registration.PatientClass
		patientInfo.Type = e.Step.Registration.PatientClass
	} else {
		generated := h.generator.NewRegistrationPatientClassAndType()
		patientInfo.Class = generated.Class
		patientInfo.Type = generated.Type
	}

	msg, err := message.BuildRegistrationADTA04(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A04 message")
	}
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) preadmission(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo
	pendingLocation, err := h.occupyBed(e.Step.PreAdmission.Loc, e.Step.PreAdmission.Bed)
	if err != nil {
		return errors.Wrap(err, locationError)
	}
	patientInfo.PendingLocation = pendingLocation
	patientInfo.ExpectedAdmitDateTime = ir.NewValidTime(e.EventTime.Add(*e.Step.PreAdmission.ExpectedAdmissionTimeFromNow))
	patientInfo.AccountStatus = h.messageConfig.PatientAccountStatus.Planned
	h.generator.AddAllergies(patientInfo, e.Step.PreAdmission.Allergies)
	h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)

	msg, err := message.BuildPreAdmitADTA05(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A05 message")
	}
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) usePatient(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	mrn := e.PatientMRN
	useMRN := e.ResolveMRN(e.Step.UsePatient.Patient)
	if h.patients.Get(useMRN) == nil {
		logLocal.Errorf("Unknown MRN in use_patient: %v", useMRN)
		return errors.New("unknown MRN in use_patient")
	}
	if useMRN != mrn && useMRN != pathway.Current {
		e.PatientMRN = useMRN
		*logLocal = *logLocal.WithField("mrn", e.PatientMRN).WithField("previous_mrn", mrn)
		logLocal.Infof("Switching patient used in pathway, useMRN = %v", useMRN)
	}
	return nil
}

func (h *Hospital) merge(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	mrn := e.PatientMRN
	patientInfo := h.patients.Get(mrn).PatientInfo
	h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)
	parentMRN := e.ResolveMRN(e.Step.Merge.Parent)
	if parentMRN != pathway.Current && parentMRN != mrn {
		logLocal.Warningf("e.Step.Merge.Parent %s is different from the current MRN %s; the data from the message might not be populated correctly", parentMRN, mrn)
		return errors.New("invalid merge state")
	}
	mrns := e.Step.Merge.Children
	var msg *message.HL7Message
	var err error
	if len(mrns) == 1 && !e.Step.Merge.ForceA40 {
		childMRN := e.ResolveMRN(mrns[0])
		msg, err = message.BuildMergeADTA34(msgHeader, patientInfo, e.EventTime, e.MessageTime, childMRN)
	} else {
		childMRNs := make([]string, len(mrns))
		for i, m := range mrns {
			childMRNs[i] = e.ResolveMRN(m)
		}
		msg, err = message.BuildMergeADTA40(msgHeader, patientInfo, e.EventTime, e.MessageTime, childMRNs)
	}
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A34 or ADT^A40 message")
	}
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) bedSwap(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	mrn := e.PatientMRN
	patientInfo := h.patients.Get(mrn).PatientInfo
	mainMRN := e.ResolveMRN(e.Step.BedSwap.Patient1)
	if mainMRN != pathway.Current && mainMRN != mrn {
		logLocal.Warningf("e.Step.BedSwap.Patient1 %s is different from the current MRN %s; the data from the message might not be populated correctly", mainMRN, mrn)
		return errors.New("invalid bed swap state")
	}
	otherMRN := e.ResolveMRN(e.Step.BedSwap.Patient2)
	*logLocal = *logLocal.WithField("secondary_mrn", otherMRN)
	otherPatient := h.patients.Get(otherMRN)
	if otherPatient == nil {
		return errors.New("unknown MRN in bed swap")
	}
	otherPatientInfo := otherPatient.PatientInfo
	if patientInfo.Location == nil || otherPatientInfo.Location == nil {
		logLocal.Errorf("Cannot do bed swap: one or both of the patient locations for bed swaps is nil; first loc nil? %t; second loc nil? %t", patientInfo.Location == nil, otherPatientInfo.Location == nil)
		return errors.New("nil location in bed swap")
	}
	patientInfo.Location, otherPatientInfo.Location = otherPatientInfo.Location, patientInfo.Location
	h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)

	msg, err := message.BuildBedSwapADTA17(msgHeader, patientInfo, e.EventTime, e.MessageTime, otherPatientInfo)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A17 message")
	}
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) addPerson(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo
	h.generator.AddAllergies(patientInfo, e.Step.AddPerson.Allergies)
	h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)
	msg, err := message.BuildAddPersonADTA28(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A28 message")
	}
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) updatePerson(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo
	h.generator.UpdateFromPathway(patientInfo, e.Step.UpdatePerson)
	h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)
	var msg *message.HL7Message
	var err error
	if patientInfo.Class == h.messageConfig.PatientClass.Inpatient {
		msg, err = message.BuildUpdatePatientADTA08(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	} else {
		msg, err = message.BuildUpdatePersonADTA31(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	}
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A08 or ADT^A31 message")
	}

	if err := h.queueMessage(logLocal, msg, e); err != nil {
		return err
	}

	patientInfo.AddDiagnosesOrProceduresToEncounter(e.EventTime, patientInfo.Diagnoses, patientInfo.Procedures)

	patientInfo.Diagnoses = nil
	patientInfo.Procedures = nil
	return nil
}

func (h *Hospital) cancelPendingAdmission(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	pathwayName := e.PathwayName
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo
	patientInfo.PriorPendingLocation = patientInfo.PendingLocation
	patientInfo.PendingLocation = nil
	patientInfo.AccountStatus = h.messageConfig.PatientAccountStatus.Cancelled
	h.updateDeathInfo(logLocal, now, pathwayName, patientInfo, e.Step.Parameters)
	msg, err := message.BuildCancelPendingAdmitADTA27(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A27 message")
	}
	h.freeSpecificLocation(logLocal, patientInfo.PriorPendingLocation, pathwayName)
	patientInfo.PriorPendingLocation = nil
	patientInfo.ExpectedAdmitDateTime = ir.NewInvalidTime()
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) cancelPendingTransfer(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	pathwayName := e.PathwayName
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo
	patientInfo.PriorPendingLocation = patientInfo.PendingLocation
	patientInfo.PendingLocation = nil
	h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)
	msg, err := message.BuildCancelPendingTransferADTA26(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A26 message")
	}
	h.freeSpecificLocation(logLocal, patientInfo.PriorPendingLocation, pathwayName)
	patientInfo.PriorPendingLocation = nil
	patientInfo.ExpectedTransferDateTime = ir.NewInvalidTime()
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) cancelPendingDischarge(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo
	h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)
	msg, err := message.BuildCancelPendingDischargeADTA25(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A25 message")
	}
	patientInfo.ExpectedDischargeDateTime = ir.NewInvalidTime()
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) deleteVisit(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	patient := h.patients.Get(e.PatientMRN)
	patientInfo := patient.PatientInfo
	pastVisitID, err := patient.PopPastVisit()
	if err != nil {
		return errors.New("Patient PastVisits empty")
	}
	// Since we cannot specify a visitID mid-pathway, default to deleting the most recent visit.
	thisVisitID := patientInfo.VisitID
	patientInfo.VisitID = pastVisitID
	h.updateDeathInfo(logLocal, now, e.PathwayName, patientInfo, e.Step.Parameters)
	msg, err := message.BuildDeleteVisitADTA23(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A23 message")
	}
	if err := h.queueMessage(logLocal, msg, e); err != nil {
		return err
	}
	patientInfo.VisitID = thisVisitID
	return nil
}

func (h *Hospital) trackDeparture(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	pathwayName := e.PathwayName
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo
	h.freeSpecificLocation(logLocal, patientInfo.Location, pathwayName)
	switch e.Step.TrackDeparture.Mode {
	case pathway.TrackMode:
		loc, err := h.occupyBed(e.Step.TrackDeparture.DestinationLoc, e.Step.TrackDeparture.DestinationBed)
		if err != nil {
			return errors.Wrap(err, locationError)
		}
		patientInfo.PriorLocation = patientInfo.Location
		patientInfo.Location = loc
	case pathway.TransitMode:
		pendingLocation, err := h.occupyBed(e.Step.TrackDeparture.DestinationLoc, e.Step.TrackDeparture.DestinationBed)
		if err != nil {
			return errors.Wrap(err, locationError)
		}
		patientInfo.PriorLocation = patientInfo.Location
		patientInfo.PendingLocation = pendingLocation
		patientInfo.Location = nil
	case pathway.TemporaryMode:
		loc := e.Step.TrackDeparture.DestinationLoc
		if patientInfo.TemporaryLocation != nil { // The patient was already in a temporary location.
			patientInfo.PriorTemporaryLocation = patientInfo.TemporaryLocation
		} else {
			patientInfo.PriorLocation = patientInfo.Location
		}
		patientInfo.TemporaryLocation = &ir.PatientLocation{Poc: loc}
		patientInfo.Location = nil
	default:
		return fmt.Errorf("unsupported mode in TrackDeparture: %q", e.Step.TrackDeparture.Mode)
	}

	h.updateDeathInfo(logLocal, now, pathwayName, patientInfo, e.Step.Parameters)
	msg, err := message.BuildTrackDepartureADTA09(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A09 message")
	}
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) trackArrival(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	msgHeader := h.generator.NewHeader(&e.Step)
	pathwayName := e.PathwayName
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo

	switch e.Step.TrackArrival.Mode {
	case pathway.TrackMode:
		patientInfo.PriorLocation = patientInfo.Location
		h.freeSpecificLocation(logLocal, patientInfo.Location, pathwayName)
		loc, err := h.occupyBed(e.Step.TrackArrival.Loc, e.Step.TrackArrival.Bed)
		if err != nil {
			return errors.Wrap(err, locationError)
		}
		patientInfo.Location = loc
	case pathway.TransitMode:
		matches, err := h.locationManager.Matches(e.Step.TrackArrival.Loc, patientInfo.PendingLocation)
		if err != nil {
			log.WithError(err).Errorf("Error matching location in TrackArrival (transit) step %q and patient location %q", e.Step.TrackArrival.Loc, patientInfo.PendingLocation)
			return errors.Wrap(err, locationError)
		}
		if !matches {
			log.WithError(err).Errorf("Location mismatch in TrackArrival (transit): step %q; patient location: %q", e.Step.TrackArrival.Loc, patientInfo.PendingLocation)
			return errors.New("transit location mismatch")
		}
		patientInfo.Location = patientInfo.PendingLocation
		patientInfo.PendingLocation = nil
	case pathway.TemporaryMode:
		// In pathway.TemporaryMode mode, the patient comes from a temporary location and arrives to a temporary or permanent location.
		switch {
		case !e.Step.TrackArrival.IsTemporary:
			// Temporary-to-permanent.
			patientInfo.PriorTemporaryLocation = patientInfo.TemporaryLocation
			patientInfo.TemporaryLocation = nil

			loc, err := h.occupyBed(e.Step.TrackArrival.Loc, e.Step.TrackArrival.Bed)
			if err != nil {
				return errors.Wrap(err, locationError)
			}
			patientInfo.Location = loc
		case patientInfo.TemporaryLocation == nil, e.Step.TrackArrival.Loc != patientInfo.TemporaryLocation.Poc:
			// Temporary-to-temporary.
			// We assume the previous location was temporary if there's none.
			// We only need to update the location when the new one is different from the current one.
			patientInfo.PriorTemporaryLocation = patientInfo.TemporaryLocation
			patientInfo.TemporaryLocation = &ir.PatientLocation{Poc: e.Step.TrackArrival.Loc}
		}
	default:
		return fmt.Errorf("unsupported mode in TrackArrival: %q", e.Step.TrackArrival.Mode)
	}
	h.updateDeathInfo(logLocal, now, pathwayName, patientInfo, e.Step.Parameters)

	msg, err := message.BuildTrackArrivalADTA10(msgHeader, patientInfo, e.EventTime, e.MessageTime)
	if err != nil {
		return errors.Wrap(err, "cannot build ADT^A10 message")
	}
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) hardcodedMessage(e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo
	toIncludeRegex := e.Step.HardcodedMessage.Regex
	msg, err := h.hardcodedMessageManager.Message(toIncludeRegex, patientInfo.Person, now)
	if err != nil {
		return errors.Wrap(err, "cannot process HardcodedMessage event")
	}
	return h.queueMessage(logLocal, msg, e)
}

func (h *Hospital) generateResources(e *state.Event, logLocal *logging.SimulatedHospitalLogger) error {
	patientInfo := h.patients.Get(e.PatientMRN).PatientInfo
	logLocal.Info("Generating resources")
	return h.resourceWriter.Generate(patientInfo)
}

// processEventType processes the given event type.
// Most events create HL7 messages that are added to the message queue.
func (h *Hospital) processEventType(ctx context.Context, e *state.Event, logLocal *logging.SimulatedHospitalLogger, now time.Time) error {
	*logLocal = *logLocal.WithField(keyEventType, e.Step.StepType()).
		WithField(keyIsHistorical, e.IsHistorical).
		WithField(keyIndex, e.Index).
		WithField(keyExpectedEventTime, e.EventTime.UTC().Format(datetimeLayout))
	logLocal.Info("Running event")
	logLocal.WithField(keyEvent, e).Debug("Running event")
	// Most event processing events require that the MRN is for a known patient. Do the check here to avoid having to
	// do it in the individual methods.
	if h.patients.Get(e.PatientMRN) == nil {
		return errors.New("unknown MRN in event")
	}

	switch e.Step.StepType() {
	case pathway.StepDelay:
		// no-op as Delay events don't have messages.
	case pathway.StepAdmission:
		return h.processAdmission(e, logLocal, now)
	case pathway.StepOrder:
		return h.processOrder(e, logLocal, now)
	case pathway.StepResults:
		return h.processResults(e, logLocal, now)
	case pathway.StepClinicalNote:
		return h.processClinicalNote(ctx, e, logLocal, now)
	case pathway.StepDocument:
		return h.processDocument(e, logLocal, now)
	case pathway.StepDischarge:
		return h.processDischarge(e, logLocal, now)
	case pathway.StepDischargeInError:
		return h.processDischargeInError(e, logLocal, now)
	case pathway.StepTransferInError, pathway.StepTransfer:
		return h.processTransferOrTransferInError(e, logLocal, now)
	case pathway.StepCancelVisit:
		return h.cancelVisit(e, logLocal, now)
	case pathway.StepCancelTransfer:
		return h.cancelTransfer(e, logLocal, now)
	case pathway.StepCancelDischarge:
		return h.cancelDischarge(e, logLocal, now)
	case pathway.StepPendingAdmission:
		return h.pendingAdmission(e, logLocal, now)
	case pathway.StepPendingDischarge:
		return h.pendingDischarge(e, logLocal, now)
	case pathway.StepPendingTransfer:
		return h.pendingTransfer(e, logLocal, now)
	case pathway.StepRegistration:
		return h.registration(e, logLocal, now)
	case pathway.StepPreAdmission:
		return h.preadmission(e, logLocal, now)
	case pathway.StepUsePatient:
		return h.usePatient(e, logLocal, now)
	case pathway.StepMerge:
		return h.merge(e, logLocal, now)
	case pathway.StepBedSwap:
		return h.bedSwap(e, logLocal, now)
	case pathway.StepAddPerson:
		return h.addPerson(e, logLocal, now)
	case pathway.StepUpdatePerson:
		return h.updatePerson(e, logLocal, now)
	case pathway.StepCancelPendingAdmission:
		return h.cancelPendingAdmission(e, logLocal, now)
	case pathway.StepCancelPendingTransfer:
		return h.cancelPendingTransfer(e, logLocal, now)
	case pathway.StepCancelPendingDischarge:
		return h.cancelPendingDischarge(e, logLocal, now)
	case pathway.StepDeleteVisit:
		return h.deleteVisit(e, logLocal, now)
	case pathway.StepTrackDeparture:
		return h.trackDeparture(e, logLocal, now)
	case pathway.StepTrackArrival:
		return h.trackArrival(e, logLocal, now)
	case pathway.StepHardcodedMessage:
		return h.hardcodedMessage(e, logLocal, now)
	case pathway.StepAutoGenerate:
		return fmt.Errorf("unsupported event type: %s; make sure Runnable() is called on the pathway before it is ran", pathway.StepAutoGenerate)
	case pathway.StepGeneric:
		// Generic events do not have a default logic by design.
		// Add an event processor in AdditionalConfig.Processors.EventOverride to specify the behaviour for these events.
		return errors.New("missing_processor_of_generic_event")
	case pathway.StepGenerateResources:
		return h.generateResources(e, logLocal)
	default:
		return fmt.Errorf("unknown_event_type_%s", e.Step.StepType())
	}
	return nil
}

func setDischargeDate(patientInfo *ir.PatientInfo, dischargeTime time.Time) {
	if patientInfo.ExpectedDischargeDateTime.Valid {
		patientInfo.DischargeDate = patientInfo.ExpectedDischargeDateTime
	} else {
		patientInfo.DischargeDate = ir.NewValidTime(dischargeTime)
	}
}

// freeLocation frees the location occupied by the patient, and returns the (now free) patient location.
func (h *Hospital) freeLocation(logLocal *logging.SimulatedHospitalLogger, patientInfo *ir.PatientInfo, pathwayName string) *ir.PatientLocation {
	h.freeSpecificLocation(logLocal, patientInfo.Location, pathwayName)
	return patientInfo.Location
}

// freeSpecificLocation frees the given location.
func (h *Hospital) freeSpecificLocation(logLocal *logging.SimulatedHospitalLogger, l *ir.PatientLocation, pathwayName string) {
	if l == nil {
		logLocal.Debug("Tried to free location but no patient location specified")
		return
	}
	if !location.IsBed(l) {
		return
	}
	if err := h.locationManager.FreeBed(l); err != nil {
		// The fact that we cannot free the patient bed might cause problems for future patients, but
		// an error here is very rare, so we just monitor it but we don't stop processing the event.
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": pathwayName,
			"reason":       err.Error(),
		}).Inc()
		logLocal.WithError(err).Warning("cannot free patient bed")
	}
}

// resetPatient clears a patient's state. This is usually needed after a discharge or a cancel
// admission events, so that potential future events (e.g., a test result without an admit) don't
// carry the inpatient information.
func (h *Hospital) resetPatient(logLocal *logging.SimulatedHospitalLogger, pathwayName string, patient *state.Patient, mrn string) *state.Patient {
	h.freeLocation(logLocal, patient.PatientInfo, pathwayName)
	return h.generator.ResetPatient(patient)
}

func updatePersonDeath(parameters *pathway.Parameters, now time.Time, person *ir.Person) {
	switch {
	case parameters.Status.TimeOfDeath != nil:
		person.DateOfDeath = ir.NewValidTime(*parameters.Status.TimeOfDeath)
	case parameters.Status.TimeSinceDeath != nil:
		person.DateOfDeath = ir.NewValidTime(now.Add(-*parameters.Status.TimeSinceDeath))
	default:
		person.DateOfDeath = ir.NewInvalidTime()
	}
	person.DeathIndicator = parameters.Status.DeathIndicator
}

// updateDeathInfo updates the patient's death status and frees patient locations if the patient is declared dead.
func (h *Hospital) updateDeathInfo(logLocal *logging.SimulatedHospitalLogger, now time.Time, pathwayName string, patientInfo *ir.PatientInfo, parameters *pathway.Parameters) {
	// Update death status if new info provided in current step.
	person := patientInfo.Person
	if parameters != nil && parameters.Status != nil {
		updatePersonDeath(parameters, now, person)
	}
	// If patient declared undead, the step processing can continue as is.
	if !person.DateOfDeath.Valid {
		return
	}
	// Otherwise, the patient was declared dead in this or a previous step.
	// Free the occupied bed and mark the different locations as the previous locations.
	// Temporary locations aren't 'freed' since they don't have a limited number of spots.
	if patientInfo.TemporaryLocation != nil {
		patientInfo.PriorTemporaryLocation = patientInfo.TemporaryLocation
		patientInfo.TemporaryLocation = nil
	}
	if patientInfo.Location != nil {
		h.freeSpecificLocation(logLocal, patientInfo.Location, pathwayName)
		patientInfo.PriorLocation = patientInfo.Location
		patientInfo.Location = nil
	}
	if patientInfo.PendingLocation != nil {
		patientInfo.ExpectedAdmitDateTime = ir.NewInvalidTime()
		patientInfo.ExpectedTransferDateTime = ir.NewInvalidTime()
		h.freeSpecificLocation(logLocal, patientInfo.PendingLocation, pathwayName)
		patientInfo.PriorPendingLocation = patientInfo.PendingLocation
		patientInfo.PendingLocation = nil
	}
	// Pending Discharge Step could have any of (pending/temporary/assigned) location occupied when called.
	patientInfo.ExpectedDischargeDateTime = ir.NewInvalidTime()
}

func (h *Hospital) setAdmissionDetailsIfMissing(patientInfo *ir.PatientInfo, eventTime time.Time) {
	// If the admission date is not set it means that the patient was not admitted.
	if patientInfo.AdmissionDate.Valid {
		return
	}
	patientInfo.Location = h.locationManager.GetAAndELocation()
	patientInfo.AdmissionDate = ir.NewValidTime(eventTime)
}
