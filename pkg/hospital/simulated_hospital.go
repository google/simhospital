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

// Package hospital implements the main functionality of Simulated Hospital.
package hospital

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/encoding/prototext"
	"github.com/google/simhospital/pkg/clock"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/doctor"
	"github.com/google/simhospital/pkg/generator"
	"github.com/google/simhospital/pkg/generator/header"
	"github.com/google/simhospital/pkg/generator/id"
	"github.com/google/simhospital/pkg/generator/person"
	"github.com/google/simhospital/pkg/hardcoded"
	"github.com/google/simhospital/pkg/hl7"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/location"
	"github.com/google/simhospital/pkg/logging"
	"github.com/google/simhospital/pkg/message"
	"github.com/google/simhospital/pkg/monitoring"
	"github.com/google/simhospital/pkg/orderprofile"
	"github.com/google/simhospital/pkg/pathway"
	"github.com/google/simhospital/pkg/processor"
	"github.com/google/simhospital/pkg/resource/cloud"
	"github.com/google/simhospital/pkg/resource"
	"github.com/google/simhospital/pkg/state/persist"
	"github.com/google/simhospital/pkg/state"
)

const (
	datetimeLayout = "2006-01-02 15:04:05"

	keyPathwayName           = "pathway_name"
	keyMessageName           = "message_name"
	keyEventType             = "event_type"
	keyEvent                 = "event"
	keyMessage               = "message"
	keyMessageType           = "message_type"
	keyLocation              = "location"
	keyNextEvent             = "next_event"
	keyNextEventType         = "next_event_type"
	keyIndex                 = "step_index"
	keyExpectedMessageTime   = "expected_message_time"
	keyExpectedEventTime     = "expected_event_time"
	keyExpectedNextEventTime = "expected_next_event_time"
	keyIsHistorical          = "is_historical"
	keyPatientID             = "mrn"

	inconsistentQueueError = "inconsistent event queue"

	unknown = "unknown"
)

var (
	log = logging.ForCallerPackage()

	defaultOrderAckDelay = &pathway.Delay{
		From: time.Second,
		To:   11 * time.Second,
	}

	counters struct {
		SimulatedHospital struct {
			PathwaysTotal            *prometheus.CounterVec   `help:"Number of pathways that were successfully started" labels:"pathway_name"`
			MessagesTotal            *prometheus.CounterVec   `help:"Number of messages sent" labels:"pathway_name,message_type,trigger_event"`
			ErrorsTotal              *prometheus.CounterVec   `help:"Number of errors" labels:"pathway_name,reason"`
			PathwayDurationMinutes   *prometheus.HistogramVec `help:"Duration (minutes) of the generated pathway, by pathway name" labels:"pathway_name" buckets:"1,5,10,30,60,180,720,1440,2880"`
			AdmissionDurationMinutes *prometheus.HistogramVec `help:"Duration (minutes) of the admissions in the generated pathways, by pathway name" labels:"pathway_name" buckets:"1,5,10,30,60,180,720,1440,2880"`
			MessageDelaySeconds      prometheus.Histogram     `help:"Difference, in seconds, between the time a message was expected to be sent, and the time when it was really sent" buckets:"1,5,10,30,60,180"`
		}
	}
)

// EventProcessor defines a processor of events that is run before, instead of or after the standard event processing logic.
type EventProcessor interface {
	// Process processes the given event and returns any HL7 messages that must be sent as the result of the processing, if any.
	Process(*state.Event, *ir.PatientInfo, *processor.Config) ([]*message.HL7Message, error)
	// Matches returns whether the given event can be processed by the processor.
	Matches(*state.Event) bool
}

// MessageProcessor defines a processor of messages that is run before, instead of or after the standard message processing logic.
type MessageProcessor interface {
	// Process processes the given message.
	Process(*state.HL7Message) error
	// Matches returns whether the given message can be processed by the processor.
	Matches(*state.HL7Message) bool
}

// ResourceWriter defines an object which can produce resources from a patient record.
type ResourceWriter interface {
	// Generate generates resources from the given PatientInfo.
	Generate(*ir.PatientInfo) error
	Close() error
}

// Arguments contains the arguments used to create a default Simulated Hospital Config.
type Arguments struct {
	// LocationsFile to create the Config.LocationManager.
	// Also required to create Config.PathwayParser and Config.PathwayManager.
	LocationsFile *string

	// HardcodedMessagesDir to create Config.MessagesManager.
	HardcodedMessagesDir *string

	// Hl7ConfigFile to create Config.HL7Config.
	// Also required to create Config.OrderProfiles.
	Hl7ConfigFile *string

	// HeaderConfigFile to create Config.Header.
	HeaderConfigFile *string

	// DoctorsFile to create Config.Doctors.
	// Also required to create Config.PathwayParser and Config.PathwayManager.
	DoctorsFile *string

	// OrderProfilesFile to create Config.OrderProfiles.
	// Also required to create Config.PathwayParser and Config.PathwayManager.
	OrderProfilesFile *string

	// ResourceArguments to create ResourceWriter.
	ResourceArguments *ResourceArguments

	// DeletePatientsFromMemory to set as Config.DeletePatientsFromMemory.
	DeletePatientsFromMemory bool

	// PathwayArguments to create Config.PathwayManager.
	PathwayArguments *PathwayArguments

	// SenderArguments to create Config.Sender.
	SenderArguments *SenderArguments

	// DataFiles to set as Config.DataFiles.
	DataFiles *config.DataFiles

	// MessageControlGenerator for the hospital. If not set, a default generator is created.
	MessageControlGenerator *header.MessageControlGenerator

	// Clock for this hospital. If not set, a default clock is created.
	Clock clock.Clock
}

// PathwayArguments contains arguments to create a Pathway Manager.
type PathwayArguments struct {
	// Dir contains all pathways to be used to create a Pathway Manager.
	Dir string

	// Names contains all pathway names to be included when creating a Pathway Manager.
	Names []string

	// ExcludeNames contains pathways names to be excluded when creating a Pathway Manager.
	ExcludeNames []string

	// Type is the way pathways are picked to be run - either "distribution" or "deterministic".
	Type string
}

// SenderArguments contains arguments to create a Sender.
type SenderArguments struct {
	// Output specified where the generated HL7 messages will be sent.
	Output string

	// OutputFile is a file path to write messages if Output=file.
	OutputFile string

	// MllpDestination is Host:Port to which MLLP messages will be sent if Output=mllp.
	MllpDestination string

	// MllpKeepAlive is whether to send keep-alive messages on the MLLP connection.
	// Only relevant if Output=mllp.
	MllpKeepAlive bool

	// MllpKeepAliveInterval is an interval between keep-alive messages.
	// Only relevant if Output=mllp and MllpKeepAlive=true.
	MllpKeepAliveInterval *time.Duration
}

// ResourceArguments contains arguments to create a ResourceWriter.
type ResourceArguments struct {
	Output    string
	OutputDir string
	Format    string

	// Arguments to connect to a Cloud FHIR store.
	// Only relevant if Output=cloud.
	CloudProjectID string
	CloudLocation  string
	CloudDataset   string
	CloudDatastore string
}

// Config contains the configuration for Simulated Hospital.
type Config struct {
	// The location manager.
	// Required.
	LocationManager *location.Manager

	// The generator of message control IDs.
	// Required.
	MessageControlGenerator *header.MessageControlGenerator

	// The manager of hardcoded messages.
	// Required.
	MessagesManager *hardcoded.Manager

	// The configuration for HL7 messages.
	// Required.
	HL7Config *config.HL7Config

	// The configuration for the headers of the HL7 messages.
	// Required.
	Header *config.Header

	// The files to load the data from.
	DataFiles config.DataFiles

	// Doctors is the set of doctors to be used in pathways.
	Doctors *doctor.Doctors

	// OrderProfiles are the order profiles to be used in pathways.
	OrderProfiles *orderprofile.OrderProfiles

	// PathwayParser is used to parse pathways.
	PathwayParser *pathway.Parser

	// PathwayManager is the pathway manager to use for the hospital.
	PathwayManager pathway.Manager

	// Clock is the clock for the hospital.
	Clock clock.Clock

	// Sender contains the sender of HL7 messages.
	Sender hl7.Sender

	// Whether patients are deleted from the in-memory map after their pathways finish.
	// Deleting patients saves memory, but patients cannot be reused for other pathways.
	DeletePatientsFromMemory bool

	// ResourceWriter is used to write resources.
	ResourceWriter ResourceWriter

	// Additional configuration.
	// Optional.
	AdditionalConfig AdditionalConfig
}

// AdditionalConfig contains optional configuration options for Simulated Hospital
// used to extend the main functionality.
// All fields are optional.
type AdditionalConfig struct {
	// ValidFn contains any additional validation for all parsed pathways.
	ValidFn func(*pathway.Pathway) error

	// Processors for events and messages.
	Processors Processors

	// ItemSyncers contains a map of item type to the ItemSyncer that will be used to sync those items
	// to more permanent storage. Currently, the following items are supported to be synced:
	//   - event
	//   - message
	//   - patient
	ItemSyncers map[string]persist.ItemSyncer

	// OrderAckDelay is the delay in sending Order Acknowledgement (ORR^O02) messages
	// after the corresponding Order message.
	OrderAckDelay *pathway.Delay

	// AddressGenerator generates random addresses.
	AddressGenerator person.AddressGenerator

	// MRNGenerator generates new MRNs.
	MRNGenerator id.Generator

	// PlacerGenerator generates Placer Order Numbers.
	PlacerGenerator id.Generator

	// FillerGenerator generates Filler Order Numbers.
	FillerGenerator id.Generator
}

// DefaultConfig returns a default Config from Arguments.
// Config may be only partially populated if some Arguments are not specified.
// It is the responsibility of the caller to initialize missing components of the Config.
func DefaultConfig(ctx context.Context, arguments Arguments) (Config, error) {
	c := Config{
		MessageControlGenerator:  &header.MessageControlGenerator{},
		Clock:                    &clock.RealTimeClock{},
		DeletePatientsFromMemory: arguments.DeletePatientsFromMemory,
	}

	if arguments.MessageControlGenerator != nil {
		c.MessageControlGenerator = arguments.MessageControlGenerator
	}

	if arguments.Clock != nil {
		c.Clock = arguments.Clock
	}

	var err error
	if arguments.LocationsFile != nil {
		if c.LocationManager, err = location.NewManager(ctx, *arguments.LocationsFile); err != nil {
			return Config{}, errors.Wrap(err, "cannot create Location Manager")
		}
	}

	if arguments.HardcodedMessagesDir != nil {
		if c.MessagesManager, err = hardcoded.NewManager(ctx, *arguments.HardcodedMessagesDir, c.MessageControlGenerator); err != nil {
			return Config{}, errors.Wrap(err, "cannot create Hardcoded Messages Manager")
		}
	}

	if arguments.Hl7ConfigFile != nil {
		if c.HL7Config, err = config.LoadHL7Config(ctx, *arguments.Hl7ConfigFile); err != nil {
			return Config{}, errors.Wrap(err, "cannot load the message HL7 configuration")
		}
	}

	if arguments.HeaderConfigFile != nil {
		if c.Header, err = config.LoadHeaderConfig(ctx, *arguments.HeaderConfigFile); err != nil {
			return Config{}, errors.Wrap(err, "cannot load the header configuration")
		}
	}

	if arguments.DoctorsFile != nil {
		if c.Doctors, err = doctor.LoadDoctors(ctx, *arguments.DoctorsFile); err != nil {
			return Config{}, errors.Wrap(err, "cannot load the doctors configuration")
		}
	}

	if arguments.OrderProfilesFile != nil && c.HL7Config != nil {
		if c.OrderProfiles, err = orderprofile.Load(ctx, *arguments.OrderProfilesFile, c.HL7Config); err != nil {
			return Config{}, errors.Wrap(err, "cannot load the order profiles")
		}
	}

	if arguments.SenderArguments != nil {
		if c.Sender, err = hl7Sender(*arguments.SenderArguments); err != nil {
			return Config{}, errors.Wrap(err, "cannot create the sender")
		}
	}

	if arguments.ResourceArguments != nil && c.HL7Config != nil {
		if c.ResourceWriter, err = resourceWriter(ctx, *arguments.ResourceArguments, c.HL7Config); err != nil {
			return Config{}, errors.Wrap(err, "cannot create the resource writer")
		}
	}

	if c.OrderProfiles != nil && c.Doctors != nil && c.LocationManager != nil {
		c.PathwayParser = &pathway.Parser{Clock: c.Clock, OrderProfiles: c.OrderProfiles, Doctors: c.Doctors, LocationManager: c.LocationManager}

		if arguments.PathwayArguments != nil {
			if c.PathwayManager, err = pathwayManager(ctx, c.PathwayParser, *arguments.PathwayArguments); err != nil {
				return Config{}, errors.Wrap(err, "cannot create pathway manager")
			}
		}
	}

	if arguments.DataFiles != nil {
		c.DataFiles = *arguments.DataFiles
	}

	return c, nil
}

func resourceWriter(ctx context.Context, arguments ResourceArguments, hl7Config *config.HL7Config) (ResourceWriter, error) {
	output, err := resourceOutput(ctx, arguments)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create resource output")
	}
	marshaller, err := resourceMarshaller(arguments)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create resource marshaller")
	}

	cfg := resource.GeneratorConfig{
		HL7Config:   hl7Config,
		IDGenerator: &id.UUIDGenerator{},
		Output:      output,
		Marshaller:  marshaller,
	}

	return resource.NewFHIRWriter(cfg)
}

func resourceOutput(ctx context.Context, arguments ResourceArguments) (resource.Output, error) {
	switch arguments.Output {
	case "stdout":
		return &resource.StdOutput{}, nil
	case "file":
		return resource.NewDirectoryOutput(arguments.OutputDir)
	case "cloud":
		return cloud.NewOutput(ctx, arguments.CloudProjectID, arguments.CloudLocation, arguments.CloudDataset, arguments.CloudDatastore)
	default:
		return nil, errors.Errorf("unsupported output type %q", arguments.Output)
	}
}

func resourceMarshaller(arguments ResourceArguments) (resource.Marshaller, error) {
	switch arguments.Format {
	case "json":
		return resource.NewJSONMarshaller()
	case "proto":
		return prototext.MarshalOptions{Multiline: true, Indent: "  "}, nil
	default:
		return nil, errors.Errorf("unsupported output format %q", arguments.Format)
	}
}

func hl7Sender(arguments SenderArguments) (hl7.Sender, error) {
	switch arguments.Output {
	case "stdout":
		return hl7.NewStdoutSender(), nil
	case "mllp":
		return hl7.NewMLLPSender(arguments.MllpDestination, arguments.MllpKeepAlive, *arguments.MllpKeepAliveInterval)
	case "file":
		return hl7.NewFileSender(arguments.OutputFile)
	default:
		return nil, errors.Errorf("unsupported output type %q", arguments.Output)
	}
}

func pathwayManager(ctx context.Context, p *pathway.Parser, arguments PathwayArguments) (pathway.Manager, error) {
	pathways, err := p.ParsePathways(ctx, arguments.Dir)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse pathways for Pathway Manager")
	}
	switch arguments.Type {
	case "distribution":
		return pathway.NewDistributionManager(pathways, arguments.Names, arguments.ExcludeNames)
	case "deterministic":
		return pathway.NewDeterministicManager(pathways, arguments.Names)
	default:
		return nil, errors.Errorf("unsupported pathway manager type %q", arguments.Type)
	}
}

// Processors for events and messages that enable configurable event/message processing logic to be run before,
// instead of or after the standard event/message processing logic.
type Processors struct {
	// EventOverride is a slice of event processors that can be used to process an event
	// instead of the standard processing logic.
	EventOverride []EventProcessor

	// EventPre is a slice of event processors that can be used to process an event
	// before the standard event processing logic or its overrides have started.
	EventPre []EventProcessor

	// EventPost is a slice of event processors that can be used to process an event
	// after the standard event processing logic or its overrides have finished.
	EventPost []EventProcessor

	// MessageOverride is a slice of message processors that can be used to process a message
	// instead of the standard message processing logic.
	MessageOverride []MessageProcessor

	// MessagePre is a slice of message processors that can be used to process a message
	// before the standard message processing logic or its overrides have started.
	MessagePre []MessageProcessor

	// MessagePost is a slice of message processors that can be used to process a message
	// after the standard message processing logic or its overrides have finished.
	MessagePost []MessageProcessor
}

// Hospital contains the main functionality of Simulated Hospital.
type Hospital struct {
	clock                   clock.Clock
	sender                  hl7.Sender
	generator               *generator.Generator
	locationManager         *location.Manager
	messageQ                *state.WrappedQueue
	eventQ                  *state.WrappedQueue
	pathwayManager          pathway.Manager
	hardcodedMessageManager *hardcoded.Manager
	patients                *state.PatientsMap
	processors              Processors
	resourceWriter          ResourceWriter
	messageConfig           *config.HL7Config
	orderAckDelay           *pathway.Delay
}

func init() {
	if err := monitoring.CreateAndRegisterMetricsFromStruct(&counters); err != nil {
		log.WithError(err).Fatal("Cannot register metrics from the 'hospital' package")
	}
}

// RunNextEventIfDue checks if there is an event available on the event queue and if it is due,
// and if so, it runs the next event.
// Returns true it there was an event for processing and the event ran successfully, false otherwise.
func (h *Hospital) RunNextEventIfDue(ctx context.Context) (bool, error) {
	if !h.hasDueEvent() {
		return false, nil
	}
	err := h.runNextEvent(ctx)
	return err == nil, err
}

// ProcessNextMessageIfDue checks if there is a message available on the message queue and if it is due,
// and if so, it processes the next message.
// Returns true if there was a message for processing and the processing was successful, false otherwise.
func (h *Hospital) ProcessNextMessageIfDue() (bool, error) {
	if !h.hasDueMessage() {
		return false, nil
	}
	err := h.processNextMessage()
	return err == nil, err
}

func (h Hospital) hasDueEvent() bool {
	i := h.eventQ.Peek()
	if i == nil {
		return false
	}
	e, ok := i.(state.Event)
	if !ok {
		log.Fatalf("Unknown item type %v, want state.Event", i)
	}
	return h.isTimeDue(e.EventTime)
}

func (h Hospital) hasDueMessage() bool {
	i := h.messageQ.Peek()
	if i == nil {
		return false
	}
	m, ok := i.(state.HL7Message)
	if !ok {
		log.Fatalf("Unknown item type %v, want state.HL7Message", i)
	}
	return h.isTimeDue(m.MessageTime)
}

func (h Hospital) isTimeDue(t time.Time) bool {
	return t.Unix() <= h.clock.Now().Unix()
}

// GetPathway gets the pathway with the given name configured in the hospital's pathway manager.
func (h *Hospital) GetPathway(pathwayName string) (*pathway.Pathway, error) {
	return h.pathwayManager.GetPathway(pathwayName)
}

// StartNextPathway starts the next pathway.
func (h *Hospital) StartNextPathway() error {
	p, err := h.pathwayManager.NextPathway()
	if err != nil {
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": "unknown",
			"reason":       "get_pathway_failure",
		}).Inc()
		return errors.Wrap(err, "cannot get next pathway")
	}
	if _, err := h.StartPathway(p); err != nil {
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": p.Name(),
			"reason":       "pathway_start_failure",
		}).Inc()
		return errors.Wrap(err, "cannot start pathway")
	}
	return nil
}

// StartPathway starts the given pathway.
// StartsPathway initialises the patient or patients this pathway refers to, and queues the first event for processing.
// StartPathway returns:
//   - the list of persons that were generated as a result of running this pathway.
//   - an error if something unexpected happened.
func (h *Hospital) StartPathway(p *pathway.Pathway) ([]*ir.Person, error) {
	logLocal := log.WithField(keyPathwayName, p.Name())

	if p.Persons == nil || len(*p.Persons) == 0 {
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": p.Name(),
			"reason":       "invalid_persons_section",
		}).Inc()
		return nil, errors.New("invalid Persons section")
	}

	idsToMRN := map[pathway.PatientID]string{}
	var mbPersons []*ir.Person
	var patients []*state.Patient

	persons := *p.Persons
	i := 1
	for id, person := range persons {
		newPerson, p := h.newOrExistingPatient(&person, p.Consultant)
		patients = append(patients, p)
		mbPersons = append(mbPersons, newPerson)
		logLocal.Infof("Starting pathway, person %v: %s %s %s",
			id, newPerson.FirstName, newPerson.Surname, newPerson.MRN)
		idsToMRN[id] = newPerson.MRN
		i++
	}
	if err := h.queueFirstEvent(*p, idsToMRN, patients...); err != nil {
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": p.Name(),
			"reason":       "queue_first_event",
		}).Inc()
		return nil, errors.Wrap(err, "Failed to queue first event")
	}

	counters.SimulatedHospital.PathwaysTotal.With(
		prometheus.Labels{"pathway_name": p.Name()},
	).Inc()

	return mbPersons, nil
}

// newOrExistingPatient returns a new or existing patient.
// If the MRN of the given person is set, this method searches for existing patients with such MRN.
// If it is not set, or if an existing patient isn't found, it generates a new person and patient.
// If the patient already exists, this method updates the patient with information contained in the
// pathway.
func (h Hospital) newOrExistingPatient(person *pathway.Person, consultant *pathway.Consultant) (*ir.Person, *state.Patient) {
	if person.MRN != "" {
		if p := h.patients.Get(person.MRN); p != nil {
			// After we load the patient, update the information with the information in the pathway.
			h.generator.UpdateFromPathway(p.PatientInfo, &pathway.UpdatePerson{Person: person})
			return p.PatientInfo.Person, p
		}
	}
	return h.newPatient(person, consultant)
}

// newPatient returns a new person and patient.
func (h Hospital) newPatient(person *pathway.Person, consultant *pathway.Consultant) (*ir.Person, *state.Patient) {
	newPerson := h.generator.NewPerson(person)
	newConsultant := h.generator.NewDoctor(consultant)
	return newPerson, h.generator.NewPatient(newPerson, newConsultant)
}

// calculateTimes calculates the time in which the event should take place, and the message should
// be sent, based on the current time and the specified delays (if any).
func calculateTimes(now time.Time, params *pathway.Parameters) (eventTime time.Time, msgTime time.Time) {
	eventTime = now
	msgTime = now
	if params != nil {
		if params.TimeFromNow != nil {
			eventTime = eventTime.Add(*params.TimeFromNow)
		}
		msgTime = eventTime.Add(params.DelayMessage.Random())
	}
	return
}

// NewHospital creates a new Hospital.
func NewHospital(ctx context.Context, c Config) (*Hospital, error) {
	if c.MessagesManager == nil {
		return nil, errors.New("Config.MessagesManager not provided; this is required")
	}
	if c.LocationManager == nil {
		return nil, errors.New("Config.LocationManager not provided; this is required")
	}
	if c.MessageControlGenerator == nil {
		return nil, errors.New("Config.MessageControlGenerator not provided; this is required")
	}
	if c.HL7Config == nil {
		return nil, errors.New("Config.HL7Config not provided; this is required")
	}
	if c.Header == nil {
		return nil, errors.New("Config.Header not provided; this is required")
	}
	if c.Doctors == nil {
		return nil, errors.New("Config.Doctors not provided; this is required")
	}
	if c.OrderProfiles == nil {
		return nil, errors.New("Config.OrderProfiles not provided; this is required")
	}
	if c.Sender == nil {
		return nil, errors.New("Config.Sender not provided; this is required")
	}
	if c.ResourceWriter == nil {
		return nil, errors.New("Config.ResourceWriter not provided; this is required")
	}
	if c.PathwayManager == nil {
		return nil, errors.New("Config.PathwayManager not provided; this is required")
	}
	if c.Clock == nil {
		return nil, errors.New("Config.Clock not provided; this is required")
	}
	ac := c.AdditionalConfig

	dataConfig, err := config.LoadData(ctx, c.DataFiles, c.HL7Config)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load the message data configuration")
	}

	genConfig := generator.Config{
		Clock:            c.Clock,
		HL7Config:        c.HL7Config,
		Header:           c.Header,
		Data:             dataConfig,
		Doctors:          c.Doctors,
		MsgCtrlGenerator: c.MessageControlGenerator,
		OrderProfiles:    c.OrderProfiles,
		AddressGenerator: ac.AddressGenerator,
		MRNGenerator:     ac.MRNGenerator,
		PlacerGenerator:  ac.PlacerGenerator,
		FillerGenerator:  ac.FillerGenerator,
	}

	messageQ := newMessageQueue(ac.ItemSyncers[state.MessageItemType])
	eventQ := newEventQueue(ac.ItemSyncers[state.EventItemType])
	patientsMap := state.NewPatientsMap(ac.ItemSyncers[state.PatientItemType], c.DeletePatientsFromMemory)

	if ac.OrderAckDelay == nil {
		ac.OrderAckDelay = defaultOrderAckDelay
	}
	return &Hospital{
		clock:                   c.Clock,
		sender:                  c.Sender,
		generator:               generator.NewGenerator(genConfig),
		locationManager:         c.LocationManager,
		messageQ:                messageQ,
		eventQ:                  eventQ,
		pathwayManager:          c.PathwayManager,
		hardcodedMessageManager: c.MessagesManager,
		patients:                patientsMap,
		processors:              c.AdditionalConfig.Processors,
		resourceWriter:          c.ResourceWriter,
		messageConfig:           c.HL7Config,
		orderAckDelay:           ac.OrderAckDelay,
	}, nil
}

// Close closes resources held by the Hospital.
// Should be called if the Hospital is no longer needed or at the program exit.
func (h *Hospital) Close() error {
	if err := h.sender.Close(); err != nil {
		return errors.Wrap(err, "error closing sender")
	}
	if err := h.resourceWriter.Close(); err != nil {
		return errors.Wrap(err, "error closing resource writer")
	}
	return nil
}

func newMessageQueue(syncer persist.ItemSyncer) *state.WrappedQueue {
	messageQ, err := state.NewWrappedQueue(state.MessageItemType, syncer)
	if err != nil {
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": unknown,
			"reason":       "cannot load messages",
		}).Inc()
		log.WithError(err).Error("Cannot load Messages from syncer")
	}
	return messageQ
}

func newEventQueue(syncer persist.ItemSyncer) *state.WrappedQueue {
	eventQ, err := state.NewWrappedQueue(state.EventItemType, syncer)
	if err != nil {
		counters.SimulatedHospital.ErrorsTotal.With(prometheus.Labels{
			"pathway_name": unknown,
			"reason":       "cannot load messages",
		}).Inc()
		log.WithError(err).Error("Cannot load Events from syncer")
	}
	return eventQ
}

// PatientsLen returns the number of patients in the internal patients map.
func (h *Hospital) PatientsLen() int {
	return h.patients.Len()
}

// EventsLen returns the number of events waiting to be processed.
func (h *Hospital) EventsLen() int {
	return h.eventQ.Len()
}

// MessagesLen returns the number of messages waiting to be processed.
func (h *Hospital) MessagesLen() int {
	return h.messageQ.Len()
}

// PatientExists returns whether the patient with the given id is a known patient.
func (h *Hospital) PatientExists(id string) bool {
	return h.patients.Get(id) != nil
}
