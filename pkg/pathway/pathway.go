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

// Package pathway contains definition of pathway in Simulated Hospital,
// as well as functionality to parse pathways and manage them.
package pathway

import (
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/google/simhospital/pkg/clock"
	"github.com/google/simhospital/pkg/constants"
	"github.com/google/simhospital/pkg/logging"
	"github.com/google/simhospital/pkg/orderprofile"
)

// The following constants represent step types.
const (
	StepDelay                  = "Delay"
	StepAdmission              = "Admission"
	StepOrder                  = "Order"
	StepResults                = "Result"
	StepDischarge              = "Discharge"
	StepRegistration           = "Registration"
	StepPreAdmission           = "PreAdmission"
	StepTransfer               = "Transfer"
	StepAddPerson              = "AddPerson"
	StepUpdatePerson           = "UpdatePerson"
	StepMerge                  = "Merge"
	StepBedSwap                = "BedSwap"
	StepTransferInError        = "TransferInError"
	StepDischargeInError       = "DischargeInError"
	StepCancelVisit            = "CancelVisit"
	StepCancelTransfer         = "CancelTransfer"
	StepCancelDischarge        = "CancelDischarge"
	StepPendingAdmission       = "PendingAdmission"
	StepPendingDischarge       = "PendingDischarge"
	StepPendingTransfer        = "PendingTransfer"
	StepCancelPendingAdmission = "CancelPendingAdmission"
	StepCancelPendingDischarge = "CancelPendingDischarge"
	StepCancelPendingTransfer  = "CancelPendingTransfer"
	StepDeleteVisit            = "DeleteVisit"
	StepTrackDeparture         = "TrackDeparture"
	StepTrackArrival           = "TrackArrival"
	StepUsePatient             = "UsePatient"
	StepAutoGenerate           = "AutoGenerate"
	StepClinicalNote           = "ClinicalNote"
	StepHardcodedMessage       = "HardcodedMessage"
	StepDocument               = "Document"
	StepGeneric                = "Generic"
	StepGenerateResources      = "GenerateResources"
)

const (
	// Current is the keyword used to mark the patient currently used in the pathway.
	Current = "CURRENT"

	stepInvalid      = "Invalid"
	defaultPatientID = "main-patient"
)

// Constants for the three situations of patient departing Tracking and Arrival.
// See http://www.hl7.eu/refactored/msgADT_A09.html and http://www.hl7.eu/refactored/msgADT_A10.html
const (
	// TrackMode represents the patient track mode.
	TrackMode = "track"
	// TransitMode represents the patient in transit mode.
	TransitMode = "transit"
	// TemporaryMode represents a temporary location change.
	TemporaryMode = "temporary"
)

// Constants for the possible update types in a Document step.
const (
	Append    = "append"
	Overwrite = "overwrite"
)

var (
	log = logging.ForCallerPackage()

	randomValues = map[string]bool{constants.NormalValue: true, constants.AbnormalHigh: true, constants.AbnormalLow: true}

	extractDecimalsRegexp = regexp.MustCompile(`[0-9]+.([0-9]+)`)
)

// Person represents a person in the pathway.
type Person struct {
	// Age is the age of the person. Only one of Age or DateOfBirth may be used at a time.
	Age *Age
	// DateOfBirth is the date of birth of the person.
	// Only one of Age or DateOfBirth may be used at a time.
	DateOfBirth *time.Time `yaml:"date_of_birth"`
	Gender      Gender
	FirstName   string `yaml:"first_name"`
	Surname     OptionalRandomString
	Address     *Address
	NHS         string
	MRN         string
}

// Gender represents a gender of the person.
type Gender OptionalRandomString

// IsSet returns true if Gender is not empty.
func (g Gender) IsSet() bool {
	return g != ""
}

const (
	// Male is a male Gender.
	Male = Gender("M")
	// Female is a female Gender.
	Female = Gender("F")
)

// UpdatePerson is a step to update an existing person.
// It produces an ADT^A08 message (Update patient information) if the person is an inpatient,
// or an ADT^A31 (Update person information) if the person is not an inpatient.
type UpdatePerson struct {
	Person     *Person
	Diagnoses  []*DiagnosisOrProcedure
	Procedures []*DiagnosisOrProcedure
	Allergies  []Allergy
}

// OptionalRandomString is a string that can be set to a normal string, or RANDOM,
// or omitted.
type OptionalRandomString string

// IsSet returns true if OptionalRandomString is not empty.
func (s OptionalRandomString) IsSet() bool {
	return s != ""
}

// IsFixedValue returns true if a given string is a fixed value that can be used "as is".
func (s OptionalRandomString) IsFixedValue() bool {
	return IsFixedValue(s.String())
}

// IsFixedValue returns true if "s" is a fixed value that can be used "as is",
// that is not a RANDOM and not empty (which indicates randomly generated values).
func IsFixedValue(s string) bool {
	return s != "" && s != constants.RandomString
}

func (s OptionalRandomString) String() string {
	return string(s)
}

// Address represents a physical address, e.g., a patient's home.
// It is used to populate the PID.11 segment (XAD - Extended Address).
// Eg.: 1 Goodwill Hunting Road^^London^^N1C 4AG^GBR^HOME
type Address struct {
	// AllRandom indicates that all address fields should be generated randomly.
	// If it is set, none of the other fields can be set.
	AllRandom  bool                 `yaml:"all_random"`
	FirstLine  OptionalRandomString `yaml:"first_line"`
	SecondLine OptionalRandomString `yaml:"second_line"`
	City       OptionalRandomString
	Postcode   OptionalRandomString
	Country    OptionalRandomString
	// Type is a type of an address, eg.: HOME.
	Type OptionalRandomString
}

// Parameters contain additional step parameters.
type Parameters struct {
	// DelayMessage is the delay between when the event happened, and when the HL7 message should be sent.
	// Both ends of the interval in the Delay must be positive.
	DelayMessage *Delay `yaml:"delay_message,omitempty"`
	// TimeFromNow is the time offset between now and when the event happened.
	// This is only allowed in historical steps and must contain a negative value,
	// otherwise the pathway is considered invalid.
	// To specify positive offsets in a pathway, use Delay steps.
	TimeFromNow *time.Duration `yaml:"time_from_now,omitempty"`
	// Status is an indicator for whether the patient is dead or not, to be sent in the PID segment.
	Status *DeathStatus
	// SendingApplication to use for this message, if different from the default.
	SendingApplication string `yaml:"sending_application,omitempty"`
	// ReceivingApplication to use for this message, if different from the default.
	ReceivingApplication string `yaml:"receiving_application,omitempty"`
	// SendingFacility to use for this message, if different from the default.
	SendingFacility string `yaml:"sending_facility,omitempty"`
	// ReceivingFacility to use for this message, if different from the default.
	ReceivingFacility string `yaml:"receiving_facility,omitempty"`
	// Custom are other parameters that can be used for custom processing.
	Custom map[string]string `yaml:"custom,omitempty"`
}

// Order is a step to place an order. It produces an ORM message followed by
// Order acknowledgement message (ORR^O02).
type Order struct {
	// OrderID links the Order with corresponding Result.
	// It doesn't need to be specified if there is only one Order-Result pair in the pathway.
	OrderID string `yaml:"order_id"`
	// OrderProfile is an order profile of the Order.
	// If set to RANDOM, the random OrderProfile will be selected.
	OrderProfile string `yaml:"order_profile"`
	// Status of the order. If order status is not provided, hl7.OrderStatus.InProcess will be used.
	OrderStatus string `yaml:"order_status"`
	// NoAcknowledgementMessage indicates, that an Order acknowledgement message (ORR^O02)
	// should not be sent following the Order message.
	// The default behaviour is that this message is always sent.
	NoAcknowledgementMessage bool `yaml:"no_acknowledgement_message"`
}

// Results is a step to generate a set of results.
// It produces an ORU message with one OBX segment per result.
type Results struct {
	// OrderID links the Results with corresponding Order.
	// It doesn't need to be specified if there is only one Order-Result pair in the pathway.
	OrderID string `yaml:"order_id"`
	// OrderProfile is an order profile of the Results.
	// If set to RANDOM, the random OrderProfile will be selected. In this case Results can not be specified,
	// as they will be generated randomly from the normal range for each test type of the selected order profile.
	OrderProfile string `yaml:"order_profile"`
	// OrderStatus is used to set the status of an Order.
	// Optional.
	// If specified, both: OrderStatus and ResultStatus need to be set.
	// If not specified, it will default to HL7Config.OrderStatus.Completed.
	OrderStatus string `yaml:"order_status"`
	// ResultStatus is used to set the status of results.
	// Optional.
	// If specified, both: OrderStatus and ResultStatus need to be set.
	// If not specified, it will default to HL7Config.ResultStatus.Final
	// (Final results; Can only be changed with a corrected result)
	// for the first result in the pathway, or to HL7Config.ResultStatus.Corrected
	// (Record coming over is a correction and thus replaces a final result)
	// for any subsequent results.
	ResultStatus string `yaml:"results_status"`
	// CollectedDateTime is the time when the test was collected.
	// Optional.
	// The valid values are:
	// - EMPTY - would set CollectedDateTime to an empty date.
	// - MIDNIGHT - would set CollectedDateTime's time to midnight.
	CollectedDateTime string `yaml:"collected_datetime"`
	// ReceivedInLabDateTime is the time when the test was received in lab.
	// Optional.
	// The valid values are:
	// - EMPTY - would set ReceivedInLabDateTime to an empty date.
	// - MIDNIGHT - would set ReceivedInLabDateTime's time to midnight.
	ReceivedInLabDateTime string `yaml:"received_in_lab_datetime"`
	// Results contain a slice of results.
	Results []*Result
	// TriggerEvent is the HL7 trigger event for the ORU message.
	// Optional.
	// Supported: R01 (default), R03 and R32, case insensitive.
	TriggerEvent string `yaml:"trigger_event"`
	// ExpectCorrection indicates that we expect a correction or amendment for the same order.
	// The next message for the same order will be treated as the amendment.
	// The NumberOfPreviousResults for this order won't advance until the amendment message is
	// received. This means that the SetIDs for the OBXs in both the initial message and the
	// amendment will be the same.
	// If ExpectCorrection is set, you can use OrderStatus and ResultStatus to set a value that
	// indicates to downstream processing systems that the order/results will be corrected later.
	ExpectCorrection bool `yaml:"expect_correction"`
}

// Result represents a single test result.
type Result struct {
	// TestName is the name of the test.
	// Required.
	TestName string `yaml:"test_name"`
	// ID is the ID of the test type, e.g. LABR or lpdc code. It overrides the
	// ir.Result.TestName.Id field, i.e. OBX.3.1 in the resulting HL7 message.
	// Optional.
	ID string
	// ResultStatus is the status of this result.
	// Optional.
	// If not specified, the value of Results.ResultStatus will be used.
	ResultStatus string `yaml:"result_status"`
	// Value is the value of this test result.
	// It may either be numerical, eg.: 0.25 / 70 / <0.5 / >= 5.1 etc,
	// or textual, eg.: "Sample haemolysised" etc.
	// If the value is numerical, Unit also needs to be specified.
	// If the value is textual, Unit cannot be specified, or this would cause the validation error.
	Value string `yaml:"value"`
	// Unit is a unit of the Value.
	// Requires if Value is numerical.
	Unit string
	// ObservationDateTimeOffset is the duration e.g. "+1h" which will set the time
	// relative to the CollectedDateTime within the enclosing Results.
	// Optional.
	// If not specified, the CollectedDateTime of the enclosing Results will be used.
	ObservationDateTimeOffset time.Duration `yaml:"observation_datetime_offset"`
	// ReferenceRange is a custom reference range for this test result.
	// Optional.
	// If not specified, the default reference range for the order profile will be used.
	ReferenceRange string `yaml:"reference_range"`
	// AbnormalFlag is an abnormal flag for this test result.
	// Optional.
	// Must be set to either HIGH or LOW if the value should be marked as abnormal.
	// NORMAL or empty string are both mapped to the normal flag (ie an empty string in HL7 message).
	// If AbnormalFlag is set to DEFAULT, it will be derived from reference ranges
	// (either custom, or from order profile) and the value.
	// The AbnormalFlag cannot be set to DEFAULT for the textual or empty value.
	AbnormalFlag constants.AbnormalFlag `yaml:"abnormal_flag"`
	// Notes are the notes that will be used to populate the NTE segments associated with this result.
	// Optional.
	Notes []string
}

// Admission is a step to admit the patient to the hospital. It produces an ADT^A01 message.
type Admission struct {
	// Loc is a location (point of care) the patient is admitted to.
	// Required.
	Loc         string
	Bed         string    `yaml:",omitempty"`
	Allergies   []Allergy `yaml:",omitempty"`
	AdmitReason string    `yaml:"admit_reason,omitempty"`
}

// Transfer is a step to transfer the patient to a different location.
// It produces an ADT^A02 message.
type Transfer struct {
	// Loc is a location (point of care) the patient is transferred to.
	// Required.
	Loc string
	Bed string
}

// Discharge is a step to discharge the patient. It produces an ADT^A03 message.
type Discharge struct {
	Note          string
	Allergies     []Allergy  `yaml:",omitempty"`
	DischargeTime *time.Time `yaml:"discharge_time"`
}

// Generic is a step for situations that do not fit the existing steps.
// Generic events have no default logic. Use Generic events together with an EventProcessor to override the custom logic.
// Use the step's Parameters.Custom to send parameters.
type Generic struct {
	// Name is an identifying name for this generic step. This allows to distinguish between generic steps if
	// for instance there are multiple types of generic steps in a pathway.
	Name string
}

// Document is a step which creates a new document.
type Document struct {
	// DocumentType populates the required TXA.2-Document Type field.
	// This field is required in HL7. Simulated Hospital generates a value if this isn't set.
	DocumentType string `yaml:"document_type"`
	// CompletionStatus populates the required TXA.17-Document Completion Status field.
	// This field is required in HL7. Simulated Hospital generates a value if this isn't set.
	CompletionStatus string `yaml:"completion_status"`
	// ObsIdentifierID populates the ID of the OBX.3-Observation Identifier field.
	// Simulated Hospital generates a value if this is null, but preserves an explicit empty string.
	ObsIdentifierID *string `yaml:"observation_identifier_id"`
	// ObsIdentifierText populates the Text of the OBX.3-Observation Identifier field.
	// Simulated Hospital generates a value if this is null, but preserves an explicit empty string.
	ObsIdentifierText *string `yaml:"observation_identifier_text"`
	// ObsIdentifierCS populates the Coding System of the OBX.3-Observation Identifier field.
	// Simulated Hospital generates a value if this is null, but preserves an explicit empty string.
	ObsIdentifierCS *string `yaml:"observation_identifier_coding_system"`
	// ID is the pathway document ID that links to a document and is unrelated to the HL7 message Document ID field.
	// It is a required field if a document is being updated.
	ID string
	// UpdateType is an optional parameter that specifies the type of update to perform.
	// The supported update types are: append, overwrite.
	UpdateType string `yaml:"update_type"`
	// EndingContentLines is an optional parameter that sets the last lines in the document content.
	EndingContentLines []string `yaml:"ending_content_lines"`
	// HeaderContentLines is an optional parameter that sets the first lines in the document content.
	HeaderContentLines []string `yaml:"header_content_lines"`
	// NumRandomContentLines is an optional parameter to control the number of lines with random content.
	// Simulated Hospital chooses this number randomly between 10 and 50 if this isn't set.
	// If you want to random content, set an empty Interval.
	NumRandomContentLines *Interval `yaml:"num_random_content_lines"`
}

// Registration is a step to register the patient. It produces an ADT^A04 message.
type Registration struct {
	PatientClass string `yaml:"patient_class"`
	Allergies    []Allergy
}

// PreAdmission is a step to pre-admit the patient. It produces an ADT^A05 message.
type PreAdmission struct {
	// Loc is a location (point of care) the patient is pre-admitted to.
	// Required.
	Loc string
	Bed string
	// ExpectedAdmissionTimeFromNow is the time offset between now and when the patient
	// is expected to be admitted.
	// Required.
	ExpectedAdmissionTimeFromNow *time.Duration `yaml:"expected_admission_time_from_now"`
	Allergies                    []Allergy
}

// TransferInError step is the same as a Transfer step, but it must be used when the Transfer step
// is followed by a CancelTransfer step.
// A TransferInError step generates the same HL7 message as Transfer step (ADT^A02),
// but we use this artificial step in order to signal a different internal state in the hospital due to
// an event happening by mistake.
// Ie, after a normal Transfer step, the patient location is typically freed and thus
// made available to others, while the new location is occupied.
// If a CancelTransfer arrives afterwards because the patient was mistakenly discharged,
// the patient's previous bed might have been occupied by another patient.
// In a real hospital this cannot happen, as the bed is still physically occupied by the first patient.
// We simulate this inconsistency with a TransferInError step that sends a Transfer message,
// but keeps the bed physically occupied in our internal state as it would be in a real hospital.
type TransferInError struct {
	// Loc is a location (point of care) the patient is supposed to be transferred to.
	// Required.
	Loc string
	Bed string
}

// DischargeInError step is the same as a Discharge step, but it must be used when the Discharge step
// is followed by a CancelDischarge step.
// A DischargeInError step generates the same HL7 message as Discharge step (ADT^A03),
// but we use this artificial step in order to signal a different internal state in the hospital due to
// an event happening by mistake.
// Ie, after a normal Discharge step, the patient location is typically freed and thus
// made available to others. If a CancelDischarge arrives afterwards because the patient was
// mistakenly discharged, the patient's bed might have been occupied by another patient.
// In a real hospital this cannot happen, as the bed is still physically occupied by the first patient.
// We simulate this inconsistency with a DischargeInError step that sends a Discharge message,
// but keeps the bed physically occupied in our internal state as it would be in a real hospital.
type DischargeInError struct {
	Note          string
	Allergies     []Allergy
	DischargeTime *time.Time `yaml:"discharge_time"`
}

// CancelVisit is a step to cancel the latest admission or visit.
// It produces an ADT^A11 message.
type CancelVisit struct{}

// AddPerson is a step to create a new person. It produces an ADT^A28 message.
// It is only allowed as the first step of the pathway.
type AddPerson struct {
	Allergies []Allergy
}

// CancelTransfer is a step to cancel a Transfer. It produces an ADT^A12 message.
type CancelTransfer struct{}

// CancelDischarge is a step to cancel a Discharge. It produces an ADT^A13 message.
type CancelDischarge struct{}

// PendingAdmission step is the step preceding the Admission step.
// It marks an event as pending to happen in the future so that some things start to
// happen already now, e.g., it reserves a bed in the given point of care in our simulator.
// It produces an ADT^A14 message to communicate this.
// PendingAdmission step requires an Admission or CancelPendingAdmission step to happen at some point later.
// Note that an Admission step after a PendingAdmission will typically need less work,
// as many things (such as reserving beds for admissions) might have already been done
// in its corresponding PendingAdmission step.
// The ExpectedAdmissionTimeFromNow parameter specifies when in the future the Admission event is expected.
// This parameter is only used to build the messages. Explicit Delay steps must be used to simulate
// the actual delays between a PendingAdmission event and its corresponding Admission event.
type PendingAdmission struct {
	// Loc is a location (point of care) the patient will be admitted to.
	// Required.
	Loc string
	Bed string
	// ExpectedAdmissionTimeFromNow specifies when in the future the Admission event is expected.
	// It is only used to build the HL7 messages.
	// Required.
	ExpectedAdmissionTimeFromNow *time.Duration `yaml:"expected_admission_time_from_now"`
}

// PendingDischarge step is the step preceding Discharge step.
// It marks an event as pending to happen in the future so that some things start to
// happen already.
// It produces an ADT^A16 message.
// PendingDischarge step requires a Discharge or CancelPendingDischarge step to happen at some point later.
// The ExpectedDischargeTimeFromNow parameter specifies when in the future the Discharge event is expected.
// This parameter is only used to build the messages. Explicit Delay steps must be used to simulate
// the actual delays between a PendingDischarge event and its corresponding Discharge event.
type PendingDischarge struct {
	// ExpectedDischargeTimeFromNow specifies when in the future the Discharge event is expected.
	// It is only used to build the HL7 messages.
	// Required.
	ExpectedDischargeTimeFromNow *time.Duration `yaml:"expected_discharge_time_from_now"`
}

// PendingTransfer step is the step preceding the Transfer step.
// It marks an event as pending to happen in the future so that some things start to
// happen already now, e.g., it reserves a bed in the given point of care in our simulator.
// It produces an ADT^A12 message to communicate this.
// PendingTransfer step requires a Transfer or CancelPendingTransfer step to happen at some point later.
// Note that an Transfer step after a PendingTransfer will typically need less work,
// as many things (such as reserving beds for transfer) might have already been done
// in its corresponding PendingTransfer step.
// The ExpectedTransferTimeFromNow parameter specifies when in the future the Transfer event is expected.
// This parameter is only used to build the messages. Explicit Delay steps must be used to simulate
// the actual delays between a PendingTransfer event and its corresponding Transfer event.
type PendingTransfer struct {
	// Loc is a location (point of care) the patient will be transferred to.
	// Required.
	Loc string
	Bed string
	// ExpectedTransferTimeFromNow specifies when in the future the Transfer event is expected.
	// It is only used to build the HL7 messages.
	// Required.
	ExpectedTransferTimeFromNow *time.Duration `yaml:"expected_transfer_time_from_now"`
}

// CancelPendingAdmission step cancels pending admission.
// It must always be preceded by the PendingAdmission step.
// It produces an ADT^A27 message.
type CancelPendingAdmission struct{}

// CancelPendingDischarge step cancels pending discharge.
// It must always be preceded by the PendingDischarge step.
// It produces an ADT^A25 message.
type CancelPendingDischarge struct{}

// CancelPendingTransfer step cancel pending transfer.
// It must always be preceded by the PendingTransfer step.
// It produces an ADT^A26 message.
type CancelPendingTransfer struct{}

// TrackDeparture step tracks the departure of a patient.
// It produces an ADT^A09 message.
// This step means that there will be a change in the patient's location,
// but an official ADT^A02 transfer hasn't been issued.
// Patient could be leaving the floor or the building, but must stay within the same healthcare institution.
// There are three modes in which this event can occur,
// see http://www.hl7.eu/refactored/msgADT_A09.html.
type TrackDeparture struct {
	// Mode is the type of departure. Supported modes are: transit, temporary or track.
	Mode string
	// DestinationLoc is the destination location (point of care) the patient is departing to.
	// Required.
	DestinationLoc string `yaml:"destination_loc"`
	// DestinationBed is the specific bed the patient is departing to.
	// Can only be set if the mode is not 'temporary'.
	// Optional.
	DestinationBed string `yaml:"destination_bed"`
}

// TrackArrival step tracks the arrival of the patient and it relates to a TrackDeparture event.
// It produces an ADT^A10 message.
// As for TrackDeparture events, there are three modes in which this event can occur,
// see http://www.hl7.eu/refactored/msgADT_A10.html.
type TrackArrival struct {
	// Mode is the type of arrival. Supported modes are: transit, temporary or track.
	Mode string
	// Loc is the destination location (point of care) the patient is arriving at.
	// Required.
	Loc string
	// Bed is the bed the patient is arriving at.
	// Optional.
	// Cannot be set in Mode 'transit'.
	Bed string
	// IsTemporary indicates whether Loc is a temporary location (e.g. X-RAY, Hallway etc.).
	// Can only be set if Mode is 'temporary'.
	IsTemporary bool `yaml:"is_temporary"`
}

// Persons maps Patient IDs to Persons used in this pathway.
type Persons map[PatientID]Person

// HasOnePerson returns whether the underlying map has exactly one person.
func (p *Persons) HasOnePerson() bool {
	return !p.isEmpty() && len(*p) == 1
}

// isEmpty returns true if the underlying map has no persons.
func (p *Persons) isEmpty() bool {
	return p == nil || len(*p) == 0
}

// isDefault returns whether the underlying map is the default one, ie:
// if it contains only one empty Person.
// Pathways without a persons section defined are initialized with default Persons
// in the Init(pathwayName) function.
func (p *Persons) isDefault() bool {
	if !p.HasOnePerson() {
		return false
	}
	patientIDs := make([]PatientID, 1)
	for pID := range *p {
		patientIDs[0] = pID
	}
	return patientIDs[0] == defaultPatientID && (*p)[defaultPatientID] == Person{}
}

// OnlyPerson returns the key and the value of the only item in the Persons map.
// OnlyPerson returns an error if the pathway does not refer to exactly one person.
func (p *Persons) OnlyPerson() (PatientID, *Person, error) {
	if !p.HasOnePerson() {
		return "", nil, errors.New("section Persons does not have one person only")
	}
	for k, v := range *p {
		return k, &v, nil
	}
	return "", nil, errors.New("no person found in the Persons section")
}

// UsePatient step defines which patient should be used from now on in the pathway.
// Cannot be set to the keyword CURRENT.
type UsePatient struct {
	// Patient is the patient ID of the patient to use in the pathway from now on.
	Patient PatientID
}

// PatientID represents a Patient ID. It can be the ID of the patient from the Persons map related to the current pathway,
// if such a section is present, or an MRN.
// Use the keyword CURRENT for the current patient in the pathway.
// See PatientsMap.Get for more details on how Simulated Hospital retrieves patients by MRN.
// Using an unknown MRN in a step will very likely result in an error.
type PatientID string

// Merge step merges two or more patients.
// This step requires one (or more) valid MRNs to exist to merge the current patient with,
// or multiple Persons being defined in the pathway, so that their identifiers can be used.
// It produces an ADT^A34 message if there is only one MRN / identifier in the Children field,
// or ADT^A40 if there are more or if the (optional) field ForceA40 set to true.
type Merge struct {
	// ForceA40 indicates to always produce ADT^A40 message, even if only two patients are merged.
	ForceA40 bool `yaml:"force_a40"`
	// Children contain the slice of patients to be merged into the Parent patient.
	// Required.
	Children []PatientID
	// Parent is the patient that the Children patients are merged to.
	// Required.
	Parent PatientID
}

// BedSwap step performs a bed swap between two patients.
// It produces an ADT^A17 message.
type BedSwap struct {
	// Patient1 is the first patient to be swapped.
	// Required.
	Patient1 PatientID `yaml:"patient_1"`
	// Patient2 is the second patient to be swapped.
	// Required.
	Patient2 PatientID `yaml:"patient_2"`
}

// Delay step is a delay between two steps in the pathway.
// It is defined as a random duration between From and To.
// Delays are not supported in Historical steps. For historical steps, use Parameters.TimeFromNow
// to specify how long in the past the event took place.
// Both From and To need to be positive, and To must be greater or equal than From.
// Otherwise the pathway is considered invalid.
type Delay struct {
	From time.Duration
	To   time.Duration
}

// DeathStatus represents a patient's death status.
type DeathStatus struct {
	// DeathIndicator is the dead indicator to set in the PID.30 - Dead Indicator field.
	// Optional.
	DeathIndicator string `yaml:"death_indicator,omitempty"`
	// TimeOfDeath is the time of death to be set in the PID.29 - Patient Death Date and Time.
	// Optional.
	// Only one of TimeOfDeath or TimeSinceDeath can be set in a step.
	TimeOfDeath *time.Time `yaml:"time_of_death,omitempty"`
	// TimeSinceDeath is how long ago the patient died, used to populate the PID.29 - Patient Death Date and Time.
	// It must be a positive duration.
	// Optional.
	// Only one of TimeOfDeath or TimeSinceDeath can be set in a step.
	TimeSinceDeath *time.Duration `yaml:"time_since_death,omitempty"`
}

// DateTime is a convenience struct that should be used to configure time fields.
// A datetime value can be specified to be absolute, relative (to now), or null.
type DateTime struct {
	Time               *time.Time
	TimeFromNow        *time.Duration `yaml:"time_from_now"`
	NoDateTimeRecorded bool           `yaml:"no_datetime_recorded"`
}

// Allergy represents an allergy.
type Allergy struct {
	// Type is a type of the allergy:
	// http://hl7-definition.caristix.com:9010/HL7%20v2.5.1/segment/Default.aspx?version=HL7%20v2.5.1&table=0127
	Type string
	// Code is the code of an allergy.
	// Either Code or Description (or both) is required.
	// If Description is missing and the Code specified is on the list of allergies
	// loaded from the allergies config file, the Description will be derived from the Code.
	Code string
	// Description is a description of an allergy.
	// Either Code or Description (or both) is required.
	// If Code is missing and the Description specified is on the list of allergies
	// loaded from the allergies config file, the Code will be derived from the Description.
	Description            string
	Severity               string
	Reaction               string
	CodingSystem           string    `yaml:"coding_system,omitempty"`
	IdentificationDateTime *DateTime `yaml:"identification_datetime,omitempty"`
}

// DiagnosisOrProcedure represents a Diagnosis or Procedure.
type DiagnosisOrProcedure struct {
	// Type is a type of diagnosis or procedure.
	// It can only be set if Code / Description is not set to RANDOM.
	Type string
	// Code is a code of a Diagnosis or Procedure.
	// Either Code or Description (or both) is required.
	// If Description is missing and the Code specified is on the list of Diagnoses / Procedures
	// loaded from the config file, the Description will be derived from the Code.
	// If Code or Description is set to RANDOM, the random Diagnosis or Procedure will be generated.
	// In this case either both: Code and Description must be set to RANDOM,
	// or one of them must be set to RANDOM and another omitted.
	// Also, Type cannot be set for RANDOM Diagnosis/Procedure.
	Code string
	// Description is a description of a Diagnosis or Procedure.
	// Either Code or Description (or both) is required.
	// If Code is missing and the Description specified is on the list of Diagnoses / Procedures
	// loaded from the config file, the Code will be derived from the Description.
	// If Code or Description is set to RANDOM, the random Diagnosis or Procedure will be generated.
	// In this case either both: Code and Description must be set to RANDOM,
	// or one of them must be set to RANDOM and another omitted.
	// Also, Type cannot be set for RANDOM Diagnosis/Procedure.
	Description string
	// DateTime is a DiagnosisOrProcedure datetime - must be in the past.
	// Required for non-random Diagnosis / Procedure.
	DateTime *DateTime `yaml:"datetime"`
}

// DeleteVisit deletes the most recently discharged or cancelled visit.
// If there are no such past visits, an error will be returned by a simulated hospital.
// It ignores the active ongoing visit, if there is one. To delete an active visit use CancelVisit.
// It produces an ADT^A23 message.
type DeleteVisit struct{}

// AutoGenerate step inserts n Results steps into the pathway,
// where n is determined by time interval From -> To and period Every.
// From and To are absolute time differences from time = 0.
// If From < 0 and To > 0, Results steps will start in History at time = From,
// with the last one generated in Pathway at time = To.
// If Results are not specified, they will be generated for a random OrderProfile.
// AutoGenerate step cannot be defined in the History.
// If From and To are the same, Every cannot be set, and there will be a single
// autogenerated Results step.
type AutoGenerate struct {
	// Result represents the results that should be generated.
	Result *Results
	From   *time.Duration
	To     *time.Duration
	Every  *time.Duration
}

// ClinicalNote is a step to send a Clinical Note document.
// It generated ORU^R01 message with a single result, with the content of the document
// in the OBX-5-ObservationValue field with the appropriate encoding.
// A clinical note is a document with information about a patient. Even if "document" could be more accurate,
// we prefer to keep the term that clinicians use.
// Some examples of clinical notes include discharge notes, images, or other documents.
type ClinicalNote struct {
	DateTime        *DateTime `yaml:"datetime"`
	DocumentType    string    `yaml:"document_type"`
	ContentType     string    `yaml:"content_type"`
	DocumentID      string    `yaml:"document_id"`
	DocumentContent string    `yaml:"document_content"`
	DocumentTitle   string    `yaml:"document_title"`
}

// HardcodedMessage is a step that sends a hardcoded message with a name
// matching the provided regular expression. Messages with matching names
// should be defined in the hardcoded messages directory. If there are multiple
// matching messages, the message to be sent will be chosen at random.
type HardcodedMessage struct {
	// Regex is the regular expression that the hardcoded message name should match.
	// Required.
	Regex string
}

// Age is randomly chosen as a number of years between From and To.
// If DayOfYear is provided as a nonzero value, it is used as a 1-indexed
// value to indicate which day of a year that person was born.
// This is useful in order to generate patients with similar demographics.
type Age struct {
	From      int
	To        int
	DayOfYear int `yaml:"day_of_year"`
}

// Interval is randomly chosen as a positive number between From and To.
type Interval struct {
	From int
	To   int
}

// GetValue returns the Value of the Result.
// If Value is set to EMPTY, the empty string is returned.
func (r *Result) GetValue() string {
	return valueOrEmptyString(r.Value)
}

// GetUnit returns the Unit of the Result.
// If Unit is set to EMPTY, the empty string is returned.
func (r *Result) GetUnit() string {
	return valueOrEmptyString(r.Unit)
}

// GetAbnormalFlag returns the AbnormalFlag.
// If AbnormalFlag is set to NORMAL, returns empty string.
// If AbnormalFlag is set to DEFAULT, returns flag derived from the reference range and the value.
// Otherwise returns the AbnormalFlag value from the Result.
func (r *Result) GetAbnormalFlag(secondaryValueGenerator *orderprofile.ValueGenerator) (constants.AbnormalFlag, error) {
	if r.AbnormalFlag == constants.AbnormalFlagNormal {
		return constants.AbnormalFlagEmpty, nil
	}
	if r.AbnormalFlag == constants.AbnormalFlagDefault {
		_, v, err := orderprofile.ValueFromString(r.GetValue())
		if err != nil {
			// This should never happen, as the pathways are validated in the startup.
			return constants.AbnormalFlagEmpty, errors.Wrapf(err, "cannot derive value from string %v", r.Value)
		}
		g := secondaryValueGenerator
		if r.ReferenceRange != "" {
			g, err = orderprofile.ValueGeneratorFromRange(r.ReferenceRange)
			if err != nil {
				// This should never happen, as the pathways are validated in the startup.
				return constants.AbnormalFlagEmpty, errors.Wrapf(err, "cannot get value generator from reference range %v", r.ReferenceRange)
			}
		}
		if g != nil {
			switch {
			case g.IsLow(v):
				return constants.AbnormalFlagLow, nil
			case g.IsHigh(v):
				return constants.AbnormalFlagHigh, nil
			default:
				return constants.AbnormalFlagEmpty, nil
			}
		}
	}
	return r.AbnormalFlag, nil
}

// GenerateResources step triggers the generation of resources (i.e. FHIR) from a
// patient's health record at that point in time.
type GenerateResources struct{}

func valueOrEmptyString(s string) string {
	if s == constants.EmptyString {
		return ""
	}
	return s
}

// GetRandomType returns the Random Type of the Random Result.
// Returns one of the result's random type: [NORMAL, ABNORMAL_HIGH, ABNORMAL_LOW],
// If Result in nil, returns NORMAL - as not specifying results in pathway means: random results for each test type
// for specified order profile.
// It returns an error if the value is not one of the random types.
func (r *Result) GetRandomType() (string, error) {
	switch {
	case r == nil:
		return constants.NormalValue, nil
	case randomValues[r.Value]:
		return r.Value, nil
	default:
		return "", fmt.Errorf("invalid random type: %s", r.Value)
	}
}

// GetValueType returns the type of the Value of the Result,
// ie: either constants.NumericalValueType or constants.TextualValueType.
func (r *Result) GetValueType() string {
	if r.Value == constants.EmptyString {
		return ""
	}
	if _, _, err := orderprofile.ValueFromString(r.Value); err == nil {
		return constants.NumericalValueType
	}
	return constants.TextualValueType
}

// IsValueRandom returns whether the Value is random,
// ie: normal, abnormal high or abnormal low.
func (r *Result) IsValueRandom() bool {
	return randomValues[r.Value]
}

// Random returns random duration between [d.From, d.To).
// If d == nil, returns 0.
// If d.From == d.To, returns d.From.
func (d *Delay) Random() time.Duration {
	if d == nil {
		return 0
	}
	if d.From == d.To {
		return d.From
	}
	return time.Duration(rand.Int63n(int64(d.To)-int64(d.From)) + int64(d.From))
}

// Random returns random int between [i.From, i.To).
// If i.From == a.To, returns i.From.
func (i *Interval) Random() int {
	if i.From == i.To {
		return i.From
	}
	return rand.Intn(i.To-i.From) + i.From
}

// random returns random int between [a.From, a.To).
// If a.From == a.To, returns a.From.
func (a *Age) random() int {
	if a.From == a.To {
		return a.From
	}
	return rand.Intn(a.To-a.From) + a.From
}

// getDayOfYear returns the 0-indexed day of the year the person was born;
// this is either given in the pathway or randomized.
func (a *Age) getDayOfYear() int {
	if a.DayOfYear > 0 {
		return a.DayOfYear - 1
	}
	return rand.Intn(365)
}

// Birthdate returns the date of birth for the givem age given a clock.
func (a *Age) Birthdate(clock clock.Clock) time.Time {
	year := clock.Now().Year() - a.random()
	dayOfYear := a.getDayOfYear()
	return time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, dayOfYear)
}

// RandomBirthdate returns the date of birth, so that the age is between 1 and 100.
func RandomBirthdate(clock clock.Clock) time.Time {
	a := &Age{From: 1, To: 100}
	return a.Birthdate(clock)
}

// Step represents an event in a patient pathway. Exactly one field (Delay, Admission, etc.) should
// be set. "Parameters" can always be set in addition to that field.
type Step struct {
	Delay                  *Delay                  `yaml:",omitempty"`
	Admission              *Admission              `yaml:",omitempty"`
	Order                  *Order                  `yaml:",omitempty"`
	Result                 *Results                `yaml:",omitempty"`
	Discharge              *Discharge              `yaml:",omitempty"`
	Registration           *Registration           `yaml:",omitempty"`
	PreAdmission           *PreAdmission           `yaml:"pre_admission,omitempty"`
	Transfer               *Transfer               `yaml:",omitempty"`
	Merge                  *Merge                  `yaml:",omitempty"`
	BedSwap                *BedSwap                `yaml:"bed_swap,omitempty"`
	TransferInError        *TransferInError        `yaml:"transfer_in_error,omitempty"`
	DischargeInError       *DischargeInError       `yaml:"discharge_in_error,omitempty"`
	CancelVisit            *CancelVisit            `yaml:"cancel_visit,omitempty"`
	CancelTransfer         *CancelTransfer         `yaml:"cancel_transfer,omitempty"`
	CancelDischarge        *CancelDischarge        `yaml:"cancel_discharge,omitempty"`
	AddPerson              *AddPerson              `yaml:"add_person,omitempty"`
	UpdatePerson           *UpdatePerson           `yaml:"update_person,omitempty"`
	PendingAdmission       *PendingAdmission       `yaml:"pending_admission,omitempty"`
	PendingDischarge       *PendingDischarge       `yaml:"pending_discharge,omitempty"`
	PendingTransfer        *PendingTransfer        `yaml:"pending_transfer,omitempty"`
	CancelPendingAdmission *CancelPendingAdmission `yaml:"cancel_pending_admission,omitempty"`
	CancelPendingDischarge *CancelPendingDischarge `yaml:"cancel_pending_discharge,omitempty"`
	CancelPendingTransfer  *CancelPendingTransfer  `yaml:"cancel_pending_transfer,omitempty"`
	DeleteVisit            *DeleteVisit            `yaml:"delete_visit,omitempty"`
	TrackDeparture         *TrackDeparture         `yaml:"track_departure,omitempty"`
	TrackArrival           *TrackArrival           `yaml:"track_arrival,omitempty"`
	UsePatient             *UsePatient             `yaml:"use_patient,omitempty"`
	AutoGenerate           *AutoGenerate           `yaml:"autogenerate,omitempty"`
	ClinicalNote           *ClinicalNote           `yaml:"clinical_note,omitempty"`
	HardcodedMessage       *HardcodedMessage       `yaml:"hardcoded_message,omitempty"`
	Document               *Document               `yaml:",omitempty"`
	Generic                *Generic                `yaml:",omitempty"`
	GenerateResources      *GenerateResources      `yaml:"generate_resources,omitempty"`
	// Up to this point, only one of the fields can be set. The pathway will be considered invalid if
	// more than one of the above fields is set.

	// Parameters contain additional parameters of this step and can be set
	// is addition to the actual Step field.
	Parameters *Parameters `yaml:",omitempty"`
	// stepType is a Step Type. It is derived based on which Step field is set and is caches,
	// so that we don't need to re-calculate it. Note this means that if fields are
	// set or unset manually, the type and the set fields could be inconsistent. However in practice
	// these fields are never manually set. Note that there could be other ways to make steps invalid
	// (since they're only validated at loading time) that don't have to do with the step type.
	stepType *stepType
}

type stepType struct {
	stepType string
}

// Consultant is the consultant to use whenever a consultant is needed in the pathway,
// except for when inserting diagnoses and procedures, where an arbitrary consultant is used.
type Consultant struct {
	ID        *string
	FirstName *string `yaml:"first_name"`
	Surname   *string
	Prefix    *string
}

// Pathway represents a pathway.
type Pathway struct {
	Percentage *Percentage `yaml:"percentage_of_patients,omitempty"`
	// Persons contain persons this pathway relates to.
	// The code that uses pathways can assume that Persons will always be present after the pathway
	// is parsed through ParsePathways or ParseSinglePathway.
	// This simplifies the checks that the downstream services need to make.
	// If Persons is not already set in the pathway, a Persons section with one entry will be added.
	Persons    *Persons
	Consultant *Consultant
	Pathway    []Step
	History    []Step `yaml:"historical_data,omitempty"`
	// metadata contains pathway's metadata and is set when the pathway is initialised
	// through Init(pathwayName).
	metadata *pathwayMetadata
}

type pathwayMetadata struct {
	// messageCount is the number of messages this pathway generates.
	messageCount int
	// mane is the name of the pathway.
	name string
}

// Percentage represents a percentage.
type Percentage float64

// Float returns Percentage as rounded float64.
func (p Percentage) Float() float64 {
	return round(float64(p))
}

// round rounds a float number to maxSignificantDigits decimal digits.
func round(f float64) float64 {
	s := strconv.FormatFloat(f, 'f', maxSignificantDigits, 64)
	rounded, _ := strconv.ParseFloat(s, 1)
	return rounded
}

func (p Percentage) significantDigits() int {
	r := extractDecimalsRegexp.FindStringSubmatch(fmt.Sprintf("%v", p))
	if len(r) >= 2 {
		return len(r[1])
	}
	return 0
}

// NewPercentage returns a new Percentage from the float64.
func NewPercentage(f float64) *Percentage {
	n := Percentage(f)
	return &n
}

// Name returns a pathway name.
func (p *Pathway) Name() string {
	return p.metadata.name
}

// UpdateName updates a pathway name.
func (p *Pathway) UpdateName(pathwayName string) {
	p.metadata.name = pathwayName
}

// MessageCount returns the number of messages that the pathway generates.
func (p *Pathway) MessageCount() (int, error) {
	if p.metadata == nil {
		log.Errorf("Pathway %v hasn't been initialised", p)
		return 0, errors.Errorf("Pathway %v hasn't been initialised", p)
	}
	return p.metadata.messageCount, nil
}

// HasPersonsDefined returns whether the pathway has persons explicitly defined,
// ie: the Persons map is not empty, and is not default.
func (p *Pathway) HasPersonsDefined() bool {
	return !p.Persons.isEmpty() && !p.Persons.isDefault()
}

// StepType returns the step type for this step, based on the field of the Step struct that is set.
func (s *Step) StepType() string {
	if s.stepType != nil {
		return s.stepType.stepType
	}
	name := stepInvalid
	v := reflect.ValueOf(s).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if !f.IsNil() && isStepTypeField(v, f) {
			if name != stepInvalid {
				// Only one field of the fields that define the step type can be set at once.
				s.stepType = &stepType{stepInvalid}
				return stepInvalid
			}
			name = t.Field(i).Name
		}
	}
	s.stepType = &stepType{name}
	return name
}

// isStepTypeField returns whether the field "f" is a field that defines the step type of a Step,
// e.g., "Admission" or "Results".
func isStepTypeField(v reflect.Value, f reflect.Value) bool {
	return f != v.FieldByName("Parameters") && f != v.FieldByName("stepType")
}

// GetDateTime evaluates the configured DateTime value and returns an absolute time or nil.
func (d *DateTime) GetDateTime(now time.Time) *time.Time {
	switch {
	case d.Time != nil:
		return d.Time
	case d.TimeFromNow != nil:
		time := now.Add(*d.TimeFromNow)
		return &time
	case d.NoDateTimeRecorded:
		return nil
	default:
		return nil
	}
}

// IsBefore evaluates the configured DateTime value and returns whether it is before a specified datetime.
// If the configured datetime value is null, it returns true.
func (d *DateTime) IsBefore(now time.Time) bool {
	dt := d.GetDateTime(now)
	return dt == nil || now.After(*dt)
}

// Init initialises the metadata for the pathway and each step.
func (p *Pathway) Init(pathwayName string) {
	if p.metadata == nil {
		p.metadata = &pathwayMetadata{}
	}
	// The code downstream assumes, for simplicity, that Persons will have at least one person populated.
	if p.Persons == nil {
		p.Persons = &Persons{defaultPatientID: {}}
	}

	// Make sure messageCount does not duplicate the number of messages in case Init(pathwayName) is called
	// multiple times.
	p.metadata.messageCount = 0
	for _, s := range append(p.History, p.Pathway...) {
		p.metadata.messageCount += s.numberOfMessages()
	}

	p.metadata.name = pathwayName
}

// numberOfMessages returns the number of messages the step generates.
func (s *Step) numberOfMessages() int {
	switch {
	case s.UsePatient != nil || s.Delay != nil:
		return 0
	case s.Order != nil && s.Order.NoAcknowledgementMessage:
		return 1
	case s.Order != nil:
		return 2
	default:
		return 1
	}
}

// Runnable returns the pathway that is ready to be ran.
// It never modifies the original pathway, but rather creates a copy.
// If the pathway has AutoGenerate steps, it parses them and generates relevant steps.
// Returns an error if AutoGenerate steps cannot be parsed.
func (p *Pathway) Runnable() (Pathway, error) {
	pathway := p.getCopy()
	if pathway.hasAutoGenerateStep() {
		if err := pathway.parseAutoGenerate(); err != nil {
			return Pathway{}, errors.Wrap(err, "cannot parse AutoGenerate step")
		}
	}
	return pathway, nil
}

// insertAtTime modifies the pathway by inserting s at time t.
// If time is negative, the step is inserted in History; if positive, in Pathway.
// (1) It will add a delay step at pathway end if pathway time < t, or
// (2) Break up the delay step into two smaller ones if t is in the middle of it,
// 	   and insert the step in between those two delays.
func (p *Pathway) insertAtTime(s Step, t time.Duration) error {
	if t < time.Duration(0) {
		return p.insertInHistory(s, t)
	}
	return p.insertInPathway(s, t)
}

func (p *Pathway) insertInHistory(s Step, t time.Duration) error {
	var clonedParams Parameters
	if s.Parameters != nil {
		clonedParams = *s.Parameters
	}
	s.Parameters = &clonedParams
	s.Parameters.TimeFromNow = &t
	i := historyIndexAtTime(p.History, t)
	newP, err := insertAtIndex(p.History, s, i)
	if err != nil {
		return errors.Wrapf(err, "cannot insert historical step %v at index %v", s, i)
	}
	p.History = newP
	return nil
}

func (p *Pathway) insertInPathway(s Step, t time.Duration) error {
	i, pathwayTime, err := pathwayIndexAtTime(p.Pathway, t)
	if err != nil {
		return errors.Wrapf(err, "cannot calculate index at time %v", t)
	}
	timeDiff := t - pathwayTime

	if timeDiff > 0 {
		newP, err := insertAtIndex(p.Pathway, Step{Delay: &Delay{From: timeDiff, To: timeDiff}}, i)
		if err != nil {
			return errors.Wrapf(err, "cannot insert step %v at index %v", &Delay{From: timeDiff, To: timeDiff}, i)
		}
		p.Pathway = newP
		i++
	} else if timeDiff < 0 {
		delay := p.Pathway[i].Delay.From
		if timeDiff != -delay {
			// We need to break up the delay into two.
			first := Step{Delay: &Delay{From: delay + timeDiff, To: delay + timeDiff}}
			second := Step{Delay: &Delay{From: (-1) * timeDiff, To: (-1) * timeDiff}}

			p.Pathway = append(p.Pathway[:i], p.Pathway[i+1:]...)
			newP, err := insertAtIndex(p.Pathway, first, i)
			if err != nil {
				return errors.Wrapf(err, "cannot insert step %v at index %v", first, i)
			}
			p.Pathway = newP
			i++
			newP, err = insertAtIndex(p.Pathway, second, i)
			if err != nil {
				return errors.Wrapf(err, "cannot insert step %v at index %v", second, i)
			}
			p.Pathway = newP
		}
	}

	newP, err := insertAtIndex(p.Pathway, s, i)
	if err != nil {
		return errors.Wrapf(err, "cannot insert step %v at index %v", s, i)
	}
	p.Pathway = newP
	return nil
}

// pathwayIndexAtTime determines pathway time and returns the index when pathway time = t.
// Pathway time is determined by setting .From = .To of a Delay to a value returned by Random(),
// which results in all subsequent calls to Random() returning that value.
// This is done for all Delay steps up to time t; if there are none, pathway time is zero.
//	(1) If there is a delay that ends at time t, we return index after that delay and t.
//	(2) If the pathway has a delay step during and lasting over t,
//		we return index where delay is and pathway time after that step.
//	(3) If we traverse the whole pathway and pathway time is less than t,
//		we return index after the last element, and current pathway time.
func pathwayIndexAtTime(p []Step, t time.Duration) (int, time.Duration, error) {
	if t < time.Duration(0) {
		return 0, time.Duration(0), fmt.Errorf("time is negative: %v", t)
	}

	pathwayTime := time.Duration(0)
	for i, x := range p {
		if x.Delay != nil {
			delay := x.Delay.Random()
			p[i] = Step{Delay: &Delay{From: delay, To: delay}}

			pathwayTime += delay
			if pathwayTime == t {
				return i + 1, pathwayTime, nil
			}
			if pathwayTime > t {
				return i, pathwayTime, nil
			}
		}
	}

	// We have reached the end of the pathway.
	return len(p), pathwayTime, nil
}

// historyIndexAtTime returns the index of the first history step that is scheduled to happen
// exactly at t or immediately after.
// If all steps happen before t, it returns the total length.
// Note that History steps might not be necessarily sorted, so it could happen that a later step
// has a time earlier than the index returned by this method.
func historyIndexAtTime(h []Step, t time.Duration) int {
	for i, s := range h {
		if s.Parameters != nil && s.Parameters.TimeFromNow != nil && *s.Parameters.TimeFromNow >= t {
			return i
		}
	}
	return len(h)
}

// insertAtIndex inserts step s into the pathway p at the index i.
// The caller is responsible for ensuring inserting s at i is safe and insertion doesn't invalidate the pathway.
// e.g. If s = CancelAdmission, which requires an Admission step before it,
// caller must insert the Admission before inserting CancelAdmission.
func insertAtIndex(p []Step, s Step, i int) ([]Step, error) {
	if i < 0 {
		return nil, fmt.Errorf("index is negative: %v", i)
	}

	p = append(p, s)
	if len(p) > 1 && i != len(p) {
		copy(p[i+1:], p[i:])
		p[i] = s
	}

	return p, nil
}

func (p *Pathway) parseAutoGenerate() error {
	var agSteps []Step

	// Removing AutoGenerate steps from the pathway in reverse order.
	r := p.reverseAutoGenerateIndices()
	for _, i := range r {
		agSteps = append([]Step{p.Pathway[i]}, agSteps...)

		if i == len(p.Pathway)-1 {
			p.Pathway = p.Pathway[:i]
		} else {
			p.Pathway = append(p.Pathway[:i], p.Pathway[i+1:]...)
		}
	}

	// Add Results steps to the pathway.
	for _, s := range agSteps {
		resultTime := *s.AutoGenerate.From
		for resultTime <= *s.AutoGenerate.To {
			if err := p.insertAtTime(Step{Result: s.AutoGenerate.Result, Parameters: s.Parameters}, resultTime); err != nil {
				return errors.Wrapf(err, "cannot insert at time %v", resultTime)
			}

			if s.AutoGenerate.Every == nil {
				// Generate a single result.
				break
			}
			resultTime += *s.AutoGenerate.Every
		}
	}
	return nil
}

func (p *Pathway) hasAutoGenerateStep() bool {
	for _, step := range p.Pathway {
		if step.AutoGenerate != nil {
			return true
		}
	}
	return false
}

// reverseAutoGenerateIndices() returns a slice of indices sorted from highest to lowest,
// where each index represents the index in Pathway where AutoGenerate step appears.
func (p *Pathway) reverseAutoGenerateIndices() []int {
	var indices []int
	for i, s := range p.Pathway {
		if s.AutoGenerate != nil {
			indices = append(indices, i)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(indices)))
	return indices
}

func (p *Pathway) getCopy() Pathway {
	copied := *p
	if p.Persons != nil {
		m := make(Persons)
		for k, v := range *p.Persons {
			m[k] = v
		}
		copied.Persons = &m
	}

	if p.History != nil {
		copied.History = make([]Step, len(p.History))
		copy(copied.History, p.History)
	}
	if p.Pathway != nil {
		copied.Pathway = make([]Step, len(p.Pathway))
		copy(copied.Pathway, p.Pathway)
	}

	if p.metadata != nil {
		metadata := *(p.metadata)
		copied.metadata = &metadata
	}

	return copied
}
