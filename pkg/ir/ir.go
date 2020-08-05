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

// Package ir contains data structures related to internal representations of entities within Simulated Hospital.
package ir

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"
	"time"

	"github.com/google/simhospital/pkg/constants"
)

// Person represents a person.
type Person struct {
	Prefix         string
	FirstName      string
	MiddleName     string
	Surname        string
	Suffix         string
	Degree         string
	Gender         string
	Ethnicity      *Ethnicity
	Birth          NullTime
	DateOfDeath    NullTime
	Address        *Address
	PhoneNumber    string
	MRN            string
	NHS            string
	DeathIndicator string
}

// Text returns a human-readable representation of a person i.e. their full name.
func (p *Person) Text() string {
	return joinNonEmpty([]string{p.Prefix, p.FirstName, p.MiddleName, p.Surname, p.Suffix}, " ")
}

// AlternateText returns a more concise representation of a person.
func (p *Person) AlternateText() string {
	return joinNonEmpty([]string{p.FirstName, p.Surname}, " ")
}

// CodedElement represents a HL7v2 Coded Element: https://hl7-definition.caristix.com/v2/HL7v2.2/DataTypes/CE.
type CodedElement struct {
	ID            string
	Text          string
	CodingSystem  string
	AlternateText string
}

// Order represents a clinical order.
type Order struct {
	// OrderProfile is the order profile for the order.
	OrderProfile *CodedElement
	// Placer is the PlacerOrderNumber to be set in the ORC and OBR segments.
	Placer string
	// Filler is the FillerOrderNumber to be set in the ORC and OBR segments.
	Filler string
	// OrderDateTime is the ORC -> Date/Time of Transaction.
	OrderDateTime NullTime
	// CollectedDateTime is the
	// OBR / OBX -> Observation Date Time (the same for all observations for one report).
	CollectedDateTime NullTime
	// ReceivedInLabDateTime is the OBR -> Specimen Received in Lab.
	ReceivedInLabDateTime NullTime
	// ReportedDateTime is the OBR -> Results Rpt/Status Change.
	ReportedDateTime NullTime
	// OrderControl is the ORC -> Order Control
	// (https://www.hl7.org/fhir/v2/0119/index.html).
	OrderControl string
	// MessageControlIDOriginalOrder is the MSH / MSA -> Message Control ID corresponding to the original Order message.
	MessageControlIDOriginalOrder string
	// OrderStatus is the ORC -> Order Status
	// (http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/Default.aspx?version=HL7%20v2.5.1&table=0038)
	OrderStatus string
	// ResultsStatus is the OBR -> Result Status
	// (http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/Default.aspx?version=HL7%20v2.5.1&table=0123)
	ResultsStatus string
	// Results (observations) relevant to ORU messages. They translate into OBX segments and contain
	// clinically relevant information, for instance:
	// "OBX|1|NM|lpdc-3384^Urea^WinPath||5.0|MMOLL|2.1 - 7.1||||F|||||".
	Results []*Result
	// ResultsForORM are the results to be included in ORM messages. They translate into OBX segments, as for Results.
	// However these are usually not clinically relevant and contain less information than the results
	// in the Results field that contain proper observations.
	// For instance, "OBX|3|CD|PERSONUKRES||Yes".
	ResultsForORM []*Result
	// NotesForORM are the notes for ORM messages. These still generate NTE segments, but such segments are located before
	// the OBX segments and refer to the order in general instead of the results as the Notes field in
	// the Result struct.
	NotesForORM      []string
	OrderingProvider *Doctor
	SpecimenSource   string
	// DiagnosticServID is the value to be set in the Diagnostic Serv Sect ID (OBR.24) field.
	// If the value matches DiagnosticServIDMDOC, the order is for a document/clinical note.
	DiagnosticServID string
	// NumberOfPreviousResults is used to keep track of how many results were already sent for this order.
	// This allows for starting with the correct OBX SetID when sending new results linked to that order.
	NumberOfPreviousResults int
}

// Result represents a clinical result.
type Result struct {
	TestName            *CodedElement
	Value               string
	Unit                string
	ValueType           string
	Range               string
	AbnormalFlag        string
	ObservationDateTime NullTime
	// Status is the OBX -> Observation Result Status
	// (http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/Default.aspx?version=HL7%20v2.5.1&table=0085).
	Status       string
	Notes        []string
	ClinicalNote *ClinicalNote
}

// Text returns a human-readable representation of the result.
func (r *Result) Text() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%s: %s %s", r.TestName.Text, r.Value, r.Unit)
	if strings.TrimSpace(r.AbnormalFlag) != "" {
		fmt.Fprintf(&sb, " (%s)", r.AbnormalFlag)
	}
	return sb.String()
}

// ClinicalNoteContent contains data used to generate an OBX segment in a ClinicalNote HL7 message.
type ClinicalNoteContent struct {
	// ObservationDateTime can be different from the DateTime field in ClinicalNote struct.
	// ObservationDateTime is set when the corresponding content is generated whereas ClinicalNote.DateTime is set when the ClinicalNote is generated.
	ObservationDateTime NullTime
	ContentType         string
	DocumentEncoding    string
	DocumentContent     string
}

// ClinicalNote represents a Clinical Note.
// A clinical note is a document with information about a patient. Even if "document" could be more accurate,
// we prefer to keep the term that clinicians use.
type ClinicalNote struct {
	DateTime      NullTime
	DocumentTitle string
	DocumentType  string
	DocumentID    string
	Contents      []*ClinicalNoteContent
}

// Document represents a generic document.
// It is used to populate the TXA and OBX segments of an MDM message.
type Document struct {
	// Fields used in TXA segment.
	ActivityDateTime         NullTime
	EditDateTime             NullTime
	DocumentType             string
	DocumentCompletionStatus string
	UniqueDocumentNumber     string

	// Fields used in OBX segments.
	// ObservationIdentifier populates the OBX.3 (Observation Identifier) field in each OBX segment.
	ObservationIdentifier *CodedElement
	// ContentLine contains values to be set in the OBX.5 (Observation Value) field.
	// Each line generates a different OBX segment.
	ContentLine []string
}

// Ethnicity is a HL7v2 coded element to represent ethnicities.
type Ethnicity CodedElement

// Address represents a physical address, e.g., a patient's home.
// Example: 1 Goodwill Hunting Road^^London^^N1C 4AG^GBR^HOME
type Address struct {
	FirstLine  string
	SecondLine string
	City       string
	PostalCode string
	Country    string
	// Type is the type of the address, eg. HOME or WORK.
	Type string
}

// PatientLocation represents a patient location within a clinical facility.
// Example: RAL 12 West^Bay01^Bed10^RAL RF^^BED^RFH^Floor 1.
type PatientLocation struct {
	Poc          string
	Room         string
	Bed          string
	Facility     string
	LocationType string
	Building     string
	Floor        string
}

// Name returns the name for this location by concatenating non-empty fields.
func (p *PatientLocation) Name() string {
	return joinNonEmpty([]string{p.Bed, p.Poc, p.Room, p.Floor, p.Building, p.Facility}, ", ")
}

// LocationHistory represents a patient location along with the period for which the patient was
// at that location.
type LocationHistory struct {
	Location *PatientLocation
	Start    NullTime
	End      NullTime
}

// Doctor represents a doctor.
// Example: 216865551019^Osman^Arthur^^^Dr^^^DRNBR^PRSNL^^^ORGDR.
type Doctor struct {
	ID        string
	Surname   string
	FirstName string
	Prefix    string
	Specialty string // This field is not used in message building.
}

// AssociatedParty represents a person associated to another person.
type AssociatedParty struct {
	*Person
	Relationship *CodedElement
	ContactRole  *CodedElement
}

// Allergy represents an allergy.
type Allergy struct {
	Type                   string
	Description            CodedElement
	Severity               string
	Reaction               string
	IdentificationDateTime NullTime
}

// DiagnosisOrProcedure represents a clinical diagnosis or procedure.
type DiagnosisOrProcedure struct {
	Description *CodedElement
	Type        string
	Clinician   *Doctor
	DateTime    NullTime
}

// Text returns a human-readable representation of a DiagnosisOrProcedure.
func (d *DiagnosisOrProcedure) Text() string {
	c := d.Clinician
	return fmt.Sprintf("%s by %s", d.Description.Text, joinNonEmpty([]string{c.Prefix, c.FirstName, c.Surname}, " "))
}

// PrimaryFacility represents a patient's primary clinical facility (e.g. a GP practice).
type PrimaryFacility struct {
	Organization string
	// ID is the "XON.3-Id Number" for this primary facility.
	// Id Number is numeric (NM) in HL7:
	// http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/segment/PD1?version=HL7%20v2.3.1&dataType=XON.
	// We make it a string instead because it's more generic.
	// Also, if it's not present, it appears in the HL7 message as an empty field as opposed to a 0.
	ID string
}

// PatientInfo represents a patient and related information.
type PatientInfo struct {
	Person          *Person
	Class           string // EMERGENCY / INPATIENT / OUTPATIENT / PREADMIT / RECURRING PATIENT / OBSTETRICS
	Type            string // values are defined per-trust if this field is used
	VisitID         uint64
	HospitalService string
	AdmitReason     string
	// Location is the current patient location.
	Location      *PatientLocation
	PriorLocation *PatientLocation
	// PriorLocationForCancelTransfer is the patient's PriorLocation after a CancelTransfer message.
	// After a transfer message we clear the patient's PriorLocation so that it's not included in
	// future messages. However in a CancelTransfer we need to know it so that we can re-instate it.
	PriorLocationForCancelTransfer *PatientLocation
	PendingLocation                *PatientLocation
	PriorPendingLocation           *PatientLocation
	TemporaryLocation              *PatientLocation
	PriorTemporaryLocation         *PatientLocation
	AttendingDoctor                *Doctor
	AccountStatus                  string
	AdmissionDate                  NullTime
	DischargeDate                  NullTime
	TransferDate                   NullTime
	ExpectedAdmitDateTime          NullTime
	ExpectedDischargeDateTime      NullTime
	ExpectedTransferDateTime       NullTime
	AssociatedParties              []*AssociatedParty
	Allergies                      []*Allergy
	// Diagnoses and Procedures are used in UpdatePerson to build ADT^A31 messages and are cleared
	// at the end of the event. Not to be confused with Encounter.Diagnoses and Encounter.Procedures
	// which persist diagnoses and procedures in the medical record.
	Diagnoses       []*DiagnosisOrProcedure
	Procedures      []*DiagnosisOrProcedure
	Encounters      []*Encounter
	PrimaryFacility *PrimaryFacility
	// AdditionalData allows users to enter arbitrary information about a patient's medical record.
	// It is up to the user to decide what data is stored here.
	AdditionalData interface{}
}

// LatestEncounter retrieves the latest encounter from PatientInfo.
func (p *PatientInfo) LatestEncounter() *Encounter {
	if len(p.Encounters) == 0 {
		return nil
	}
	return p.Encounters[len(p.Encounters)-1]
}

// AddEncounter creates a new Encounter, adds it to the list of Encounters, and sets its status
// and location.
func (p *PatientInfo) AddEncounter(startTime NullTime, status string, loc *PatientLocation) *Encounter {
	ec := &Encounter{Status: status, StatusStart: startTime, Start: startTime}
	p.Encounters = append(p.Encounters, ec)
	ec.UpdateLocation(startTime, loc)
	return ec
}

// EndEncounter finishes the specified Encounter and sets its status and final location end time.
// Note that this takes an Encounter as a receiver, unlike AddEncounter.
func (ec *Encounter) EndEncounter(endTime NullTime, newStatus string) {
	ec.UpdateStatus(endTime, newStatus)
	ec.End = endTime
	if l := len(ec.LocationHistory); l > 0 {
		ec.LocationHistory[l-1].End = endTime
	}
}

// AddOrderToEncounter either adds the specified order to the current on-going Encounter, or
// creates a new Encounter for the Order if one does not exist. In the latter case, the new
// Encounter will contain only that Order, and we set the start/end time of the Encounter to the
// order/reported time of the Order.
func (p *PatientInfo) AddOrderToEncounter(o *Order) {
	ec := p.LatestEncounter()
	if ec != nil && !ec.hasEnded() {
		ec.UpdateStatus(o.OrderDateTime, constants.EncounterStatusInProgress)
		ec.UpdateLocation(o.OrderDateTime, p.Location)
	} else {
		ec = p.AddEncounter(o.OrderDateTime, constants.EncounterStatusInProgress, p.Location)
		// If the Order does not have any Results associated with it ReportedDateTime will be nil.
		// In this case we set the end time to the OrderDateTime instead. If a Result is added in the
		// future, the Order's ReportedDateTime will be updated but the Encounter's end time will not.
		endTime := o.ReportedDateTime
		if !endTime.Valid {
			endTime = o.OrderDateTime
		}
		ec.EndEncounter(endTime, constants.EncounterStatusFinished)
	}
	ec.Orders = append(ec.Orders, o)
}

// AddDiagnosesOrProceduresToEncounter either adds the specified DiagnosisOrProcedures to the current on-going
// Encounter, or creates a new Encounter for *each* DiagnosisOrProcedure, if one does not exist.
func (p *PatientInfo) AddDiagnosesOrProceduresToEncounter(startTime time.Time, diagnoses []*DiagnosisOrProcedure, procedures []*DiagnosisOrProcedure) {
	t := NewValidTime(startTime)
	ec := p.LatestEncounter()
	if ec != nil && !ec.hasEnded() {
		// We use the event time because DiagnosisOrProcedure.DateTime is likely to be outside the
		// period of the Encounter itself.
		ec.UpdateStatus(t, constants.EncounterStatusInProgress)
		ec.Diagnoses = append(ec.Diagnoses, diagnoses...)
		ec.Procedures = append(ec.Procedures, procedures...)
	} else {
		for _, d := range diagnoses {
			ec = p.AddEncounter(t, constants.EncounterStatusInProgress, p.Location)
			ec.Diagnoses = append(ec.Diagnoses, d)
			ec.EndEncounter(t, constants.EncounterStatusFinished)
		}
		for _, pr := range procedures {
			ec = p.AddEncounter(t, constants.EncounterStatusInProgress, p.Location)
			ec.Procedures = append(ec.Procedures, pr)
			ec.EndEncounter(t, constants.EncounterStatusFinished)
		}
	}
}

// Encounter represents an interaction between a patient and healthcare provider.
type Encounter struct {
	Status string
	// StatusStart tracks the start time of the current status.
	StatusStart NullTime
	// StatusHistory tracks the lifecycle of this Encounter.
	// e.g. [planned -> arrived -> finished] or [planned -> in-progress -> cancelled]
	StatusHistory []*StatusHistory
	IsPending     bool
	// Start to End encompasses the entire period that this Encounter is active for.
	Start NullTime
	End   NullTime
	// LocationHistory tracks where the patient has been during this Encounter.
	LocationHistory []*LocationHistory
	// Orders tracks the Orders for this Encounter. Each entry in Patient.Orders is associated with
	// exactly one Encounter.
	Orders []*Order
	// Diagnoses and Procedures track the diagnoses and procedures for each Encounter. This is
	// different from PatientInfo.Procedures and PatientInfo.Diagnoses, which are used for building
	// ADT^A31 messages and are cleared after each UpdatePerson step.
	Diagnoses  []*DiagnosisOrProcedure
	Procedures []*DiagnosisOrProcedure
}

// Text returns a human-readable representation of an Encounter.
func (ec *Encounter) Text() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Status: %s\nActive from %s", ec.Status, ec.Start.Format(time.ANSIC))
	if ec.hasEnded() {
		fmt.Fprintf(&sb, " until %s", ec.End.Format(time.ANSIC))
	}
	return sb.String()
}

func (ec *Encounter) hasEnded() bool {
	return ec.Status == constants.EncounterStatusFinished || ec.Status == constants.EncounterStatusCancelled
}

// StatusHistory represents an Encounter's status and its period.
type StatusHistory struct {
	// Status must be a value from the value set EncounterStatus:
	// https://www.hl7.org/fhir/codesystem-encounter-status.html
	Status string
	Start  NullTime
	End    NullTime
}

// UpdateStatus ends the current status and appends a new entry to an Encounter's status history.
// If the new status is the same as the current one, or the start time is before the Encounter's
// start time, this is a no-op.
func (ec *Encounter) UpdateStatus(startTime NullTime, newStatus string) {
	if ec.Status == newStatus || startTime.Before(ec.Start.Time) {
		return
	}
	oldStatus := &StatusHistory{Status: ec.Status, Start: ec.StatusStart, End: startTime}
	ec.StatusHistory = append(ec.StatusHistory, oldStatus)
	ec.Status = newStatus
	ec.StatusStart = startTime
}

// UpdateLocation sets the end time of the current location and appends a new entry to an
// Encounter's location history. If the new location is the same as the current one this is a no-op.
func (ec *Encounter) UpdateLocation(startTime NullTime, newLocation *PatientLocation) {
	if newLocation == nil {
		return
	}
	if l := len(ec.LocationHistory); l > 0 {
		oldLocation := ec.LocationHistory[l-1]
		if *newLocation == *oldLocation.Location {
			return
		}
		oldLocation.End = startTime
	}
	newLocationHistory := &LocationHistory{
		Location: newLocation,
		Start:    startTime,
	}
	ec.LocationHistory = append(ec.LocationHistory, newLocationHistory)
}

// NullTime represents a time that can be null.
type NullTime struct {
	time.Time
	Valid    bool
	Midnight bool
}

// GobEncode returns the gob encoding of NullTime.
// This is necessary to prevent `time.Time.GobEncode()` being called instead,
// which will discard the `Valid` and `Midnight` fields.
func (t NullTime) GobEncode() ([]byte, error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	if err := enc.Encode(t.Time); err != nil {
		return nil, err
	}
	if err := enc.Encode(t.Valid); err != nil {
		return nil, err
	}
	if err := enc.Encode(t.Midnight); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// GobDecode performs the inverse of GobEncode.
// It modifies the receiver, so it must take a pointer receiver.
func (t *NullTime) GobDecode(data []byte) error {
	b := bytes.NewBuffer(data)
	dec := gob.NewDecoder(b)
	if err := dec.Decode(&t.Time); err != nil {
		return err
	}
	if err := dec.Decode(&t.Valid); err != nil {
		return err
	}
	return dec.Decode(&t.Midnight)
}

// NewMidnightTime returns a NullTime from the given time with Midnight and Valid set.
func NewMidnightTime(t time.Time) NullTime {
	return NullTime{
		Time:     t,
		Valid:    true,
		Midnight: true,
	}
}

// NewValidTime returns a NullTime from the given time with Valid set.
func NewValidTime(t time.Time) NullTime {
	return NullTime{
		Time:  t,
		Valid: true,
	}
}

// NewInvalidTime returns an invalid NullTime.
func NewInvalidTime() NullTime {
	return NullTime{
		Valid: false,
	}
}

// Formattable is an interface for formatting dates in different locations.
type Formattable interface {
	In(loc *time.Location) time.Time
}

func joinNonEmpty(parts []string, separator string) string {
	var toJoin []string
	for _, s := range parts {
		if strings.TrimSpace(s) != "" {
			toJoin = append(toJoin, s)
		}
	}
	return strings.Join(toJoin, separator)
}
