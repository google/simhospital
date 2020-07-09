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

// Package resource contains functionality for generating resources from PatientInfo.
package resource

import (
	"errors"
	"io"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/constants"
	"github.com/google/simhospital/pkg/gender"
	"github.com/google/simhospital/pkg/generator/codedelement"
	"github.com/google/simhospital/pkg/generator/id"
	"github.com/google/simhospital/pkg/generator/order"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/logging"

	cpb "google/fhir/proto/r4/core/codes_go_proto"
	dpb "google/fhir/proto/r4/core/datatypes_go_proto"
	aipb "google/fhir/proto/r4/core/resources/allergy_intolerance_go_proto"
	r4pb "google/fhir/proto/r4/core/resources/bundle_and_contained_resource_go_proto"
	encounterpb "google/fhir/proto/r4/core/resources/encounter_go_proto"
	locationpb "google/fhir/proto/r4/core/resources/location_go_proto"
	observationpb "google/fhir/proto/r4/core/resources/observation_go_proto"
	patientpb "google/fhir/proto/r4/core/resources/patient_go_proto"
)

const microPerNano = int64(time.Microsecond / time.Nanosecond)

var log = logging.ForCallerPackage()

// Default value for cpb.AddressUseCode_Value is AddressUseCode_INVALID_UNINITIALIZED.
var internalToFHIRAddressType = map[string]cpb.AddressUseCode_Value{
	"HOME": cpb.AddressUseCode_HOME,
	"WORK": cpb.AddressUseCode_WORK,
}

// Default value for cpb.EncounterStatusCode is EncounterStatusCode_INVALID_UNINITIALIZED.
var internalToFHIREncounterStatus = map[string]cpb.EncounterStatusCode_Value{
	constants.EncounterStatusPlanned:    cpb.EncounterStatusCode_PLANNED,
	constants.EncounterStatusInProgress: cpb.EncounterStatusCode_IN_PROGRESS,
	constants.EncounterStatusArrived:    cpb.EncounterStatusCode_ARRIVED,
	constants.EncounterStatusFinished:   cpb.EncounterStatusCode_FINISHED,
	constants.EncounterStatusCancelled:  cpb.EncounterStatusCode_CANCELLED,
	constants.EncounterStatusUnknown:    cpb.EncounterStatusCode_UNKNOWN,
}

// Marshaller defines an object that can marshal a protocol buffer message.
type Marshaller interface {
	Marshal(proto.Message) ([]byte, error)
}

// GeneratorConfig is the configuration for resource generators.
type GeneratorConfig struct {
	Writer      io.Writer
	HL7Config   *config.HL7Config
	IDGenerator id.Generator
	Output      Output
	Marshaller  Marshaller
}

// NewFHIRWriter constructs and returns a new FHIRWriter.
func NewFHIRWriter(cfg GeneratorConfig) (*FHIRWriter, error) {
	ac, err := codedelement.NewAllergyConvertor(cfg.HL7Config)
	if err != nil {
		return nil, err
	}
	return &FHIRWriter{
		gc:          gender.NewConvertor(cfg.HL7Config),
		oc:          order.NewConvertor(cfg.HL7Config),
		ac:          ac,
		cc:          codedelement.NewCodingSystemConvertor(cfg.HL7Config),
		idGenerator: cfg.IDGenerator,
		output:      cfg.Output,
		marshaller:  cfg.Marshaller,
	}, nil
}

// FHIRWriter generates FHIR resources as protocol buffers, and writes them to writer.
type FHIRWriter struct {
	gc          gender.Convertor
	oc          order.Convertor
	ac          codedelement.AllergyConvertor
	cc          codedelement.CodingSystemConvertor
	idGenerator id.Generator
	count       int
	output      Output
	marshaller  Marshaller
}

// Generate generates FHIR resources from PatientInfo.
func (w *FHIRWriter) Generate(p *ir.PatientInfo) error {
	if p == nil {
		return errors.New("cannot generate resources from nil PatientInfo")
	}
	b := w.bundle(p)
	bytes, err := w.marshaller.Marshal(b)
	if err != nil {
		return err
	}
	f, err := w.output.New(p)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.Write(bytes); err != nil {
		return err
	}
	w.count = w.count + len(b.GetEntry()) + 1 // Include the Bundle resource itself.
	return nil
}

// Close closes the FHIRWriter.
func (w *FHIRWriter) Close() error {
	log.Infof("Resources successfully written by the FHIRWriter: %d", w.count)
	return nil
}

// bundle converts PatientInfo into FHIR and returns an R4 Bundle. Bundle is the top-level
// record encapsulating a patient's medical history.
func (w *FHIRWriter) bundle(p *ir.PatientInfo) *r4pb.Bundle {
	bundle := &r4pb.Bundle{
		Type: &r4pb.Bundle_TypeCode{
			Value: cpb.BundleTypeCode_COLLECTION,
		},
	}

	patient, patientRef := w.patient(p)
	bundle.Entry = append(bundle.Entry, patient)

	allergies := w.allergies(p.Allergies, patientRef)
	for _, a := range allergies {
		bundle.Entry = append(bundle.Entry, a)
	}

	for _, ec := range p.Encounters {
		encounter, encounterRef := w.encounter(ec)
		bundle.Entry = append(bundle.Entry, encounter)

		e := encounter.GetResource().GetEncounter()
		for _, lh := range ec.LocationHistory {
			location, locationRef := w.location(lh.Location)
			bundle.Entry = append(bundle.Entry, location)
			e.Location = append(e.Location, w.encounterLocation(locationRef, lh.Start, lh.End))
		}

		for _, o := range ec.Orders {
			observations := w.observations(encounterRef, patientRef, o)
			bundle.Entry = append(bundle.Entry, observations...)
		}
	}

	return bundle
}

func (w *FHIRWriter) patient(patient *ir.PatientInfo) (*r4pb.Bundle_Entry, *dpb.Reference) {
	id := w.idGenerator.NewID()
	displayName := patient.Person.FirstName + " " + patient.Person.Surname

	entry := &r4pb.Bundle_Entry{
		Resource: &r4pb.ContainedResource{
			OneofResource: &r4pb.ContainedResource_Patient{
				&patientpb.Patient{
					Id:         &dpb.Id{Value: id},
					Identifier: w.personIdentifier(patient.Person),
					Name:       w.humanName(patient.Person),
					Address:    w.address(patient.Person.Address),
					Deceased:   w.deceased(patient.Person),
					Telecom:    w.telecom(patient.Person.PhoneNumber),
					Gender: &patientpb.Patient_GenderCode{
						Value: w.gc.HL7ToFHIR(patient.Person.Gender),
					},
				},
			},
		},
	}

	ref := &dpb.Reference{
		Reference: &dpb.Reference_PatientId{&dpb.ReferenceId{Value: id}},
		Display:   &dpb.String{Value: displayName},
	}

	return entry, ref
}

func (w *FHIRWriter) allergies(allergies []*ir.Allergy, patientRef *dpb.Reference) []*r4pb.Bundle_Entry {
	var entries []*r4pb.Bundle_Entry
	for _, a := range allergies {
		id := w.idGenerator.NewID()

		entry := &r4pb.Bundle_Entry{
			Resource: &r4pb.ContainedResource{
				OneofResource: &r4pb.ContainedResource_AllergyIntolerance{
					&aipb.AllergyIntolerance{
						Id: &dpb.Id{Value: id},
						// Simulated Hospital does not yet distinguish between allergies and intolerances.
						Type: &aipb.AllergyIntolerance_TypeCode{Value: cpb.AllergyIntoleranceTypeCode_ALLERGY},
						Category: []*aipb.AllergyIntolerance_CategoryCode{{
							Value: w.ac.TypeHL7ToFHIR(a.Type),
						}},
						Reaction: []*aipb.AllergyIntolerance_Reaction{{
							Description: &dpb.String{Value: a.Reaction},
							Severity: &aipb.AllergyIntolerance_Reaction_SeverityCode{
								Value: w.ac.SeverityHL7ToFHIR(a.Severity),
							},
						}},
						Code:         w.codeableConcept(a.Description),
						RecordedDate: w.dateTime(a.IdentificationDateTime),
						Patient:      patientRef,
					},
				},
			},
		}
		entries = append(entries, entry)
	}
	return entries
}

func (w *FHIRWriter) codeableConcept(c ir.CodedElement) *dpb.CodeableConcept {
	return &dpb.CodeableConcept{
		Text: &dpb.String{Value: c.Text},
		Coding: []*dpb.Coding{{
			System:  &dpb.Uri{Value: w.cc.HL7ToFHIR(c.CodingSystem)},
			Code:    &dpb.Code{Value: c.ID},
			Display: &dpb.String{Value: c.Text},
		}},
	}
}

func (w *FHIRWriter) personIdentifier(pe *ir.Person) []*dpb.Identifier {
	return []*dpb.Identifier{{Value: &dpb.String{Value: pe.MRN}}}
}

func (w *FHIRWriter) humanName(pe *ir.Person) []*dpb.HumanName {
	n := &dpb.HumanName{
		Family: &dpb.String{Value: pe.Surname},
		Given:  []*dpb.String{{Value: pe.FirstName}},
	}
	if pe.MiddleName != "" {
		n.Given = append(n.Given, &dpb.String{Value: pe.MiddleName})
	}
	if pe.Prefix != "" {
		n.Prefix = []*dpb.String{{Value: pe.Prefix}}
	}
	if pe.Suffix != "" {
		n.Suffix = []*dpb.String{{Value: pe.Suffix}}
	}
	return []*dpb.HumanName{n}
}

func (w *FHIRWriter) telecom(phone string) []*dpb.ContactPoint {
	if phone == "" {
		return nil
	}
	return []*dpb.ContactPoint{{
		Value:  &dpb.String{Value: phone},
		System: &dpb.ContactPoint_SystemCode{Value: cpb.ContactPointSystemCode_PHONE},
		Use:    &dpb.ContactPoint_UseCode{Value: cpb.ContactPointUseCode_HOME},
	}}
}

func (w *FHIRWriter) deceased(pe *ir.Person) *patientpb.Patient_DeceasedX {
	if pe.DateOfDeath.Valid {
		return &patientpb.Patient_DeceasedX{
			Choice: &patientpb.Patient_DeceasedX_DateTime{
				DateTime: w.dateTime(pe.DateOfDeath),
			},
		}
	}
	return &patientpb.Patient_DeceasedX{
		Choice: &patientpb.Patient_DeceasedX_Boolean{
			Boolean: &dpb.Boolean{Value: pe.DeathIndicator != ""},
		},
	}
}

func (w *FHIRWriter) address(address *ir.Address) []*dpb.Address {
	a := &dpb.Address{
		// Confusingly, Simulated Hospital's concept of "Type" maps to FHIR's concept of "Use", *not* "Type".
		Use: &dpb.Address_UseCode{Value: internalToFHIRAddressType[address.Type]},
		// Simulated Hospital does not support this concept, so we default to "BOTH".
		Type:       &dpb.Address_TypeCode{Value: cpb.AddressTypeCode_BOTH},
		Line:       []*dpb.String{{Value: address.FirstLine}},
		City:       &dpb.String{Value: address.City},
		PostalCode: &dpb.String{Value: address.PostalCode},
		Country:    &dpb.String{Value: address.Country},
	}
	if address.SecondLine != "" {
		a.Line = append(a.GetLine(), &dpb.String{Value: address.SecondLine})
	}
	return []*dpb.Address{a}
}

func (w *FHIRWriter) encounter(ec *ir.Encounter) (*r4pb.Bundle_Entry, *dpb.Reference) {
	id := w.idGenerator.NewID()

	entry := &r4pb.Bundle_Entry{
		Resource: &r4pb.ContainedResource{
			OneofResource: &r4pb.ContainedResource_Encounter{
				&encounterpb.Encounter{
					Id: &dpb.Id{Value: id},
					Status: &encounterpb.Encounter_StatusCode{
						Value: internalToFHIREncounterStatus[ec.Status],
					},
					Period: &dpb.Period{
						Start: w.dateTime(ec.Start),
						End:   w.dateTime(ec.End),
					},
					StatusHistory: w.statusHistory(ec.StatusHistory),
				},
			},
		},
	}

	ref := &dpb.Reference{
		Reference: &dpb.Reference_EncounterId{&dpb.ReferenceId{Value: id}},
	}

	return entry, ref
}

func (w *FHIRWriter) encounterLocation(locationRef *dpb.Reference, start ir.NullTime, end ir.NullTime) *encounterpb.Encounter_Location {
	return &encounterpb.Encounter_Location{
		Location: locationRef,
		Period: &dpb.Period{
			Start: w.dateTime(start),
			End:   w.dateTime(end),
		},
	}
}

func (w *FHIRWriter) statusHistory(statusHistory []*ir.StatusHistory) []*encounterpb.Encounter_StatusHistory {
	var sh []*encounterpb.Encounter_StatusHistory
	for _, s := range statusHistory {
		h := &encounterpb.Encounter_StatusHistory{
			Status: &encounterpb.Encounter_StatusHistory_StatusCode{
				Value: internalToFHIREncounterStatus[s.Status],
			},
			Period: &dpb.Period{
				Start: w.dateTime(s.Start),
				End:   w.dateTime(s.End),
			},
		}
		sh = append(sh, h)
	}
	return sh
}

func (w *FHIRWriter) observations(encounterRef *dpb.Reference, patientRef *dpb.Reference, order *ir.Order) []*r4pb.Bundle_Entry {
	var observations []*r4pb.Bundle_Entry
	for _, r := range order.Results {
		id := w.idGenerator.NewID()
		entry := &r4pb.Bundle_Entry{
			Resource: &r4pb.ContainedResource{
				OneofResource: &r4pb.ContainedResource_Observation{
					&observationpb.Observation{
						Encounter: encounterRef,
						Subject:   patientRef,
						Id:        &dpb.Id{Value: id},
						Note:      w.notes(r.Notes),
						Status: &observationpb.Observation_StatusCode{
							Value: w.oc.HL7ToFHIR(r.Status),
						},
						Text: w.narrative(r.Text(), strings.Join(r.Notes, "; ")),
						Effective: &observationpb.Observation_EffectiveX{
							Choice: &observationpb.Observation_EffectiveX_DateTime{
								DateTime: w.dateTime(order.OrderDateTime),
							},
						},
						Value: &observationpb.Observation_ValueX{
							Choice: &observationpb.Observation_ValueX_Quantity{
								Quantity: &dpb.Quantity{
									Value: &dpb.Decimal{Value: r.Value},
									Unit:  &dpb.String{Value: r.Unit},
								},
							},
						},
					},
				},
			},
		}
		observations = append(observations, entry)
	}
	return observations
}

func (w *FHIRWriter) narrative(paragraphs ...string) *dpb.Narrative {
	var sb strings.Builder
	sb.WriteString("<div>")
	for _, p := range paragraphs {
		if p == "" {
			continue
		}
		sb.WriteString("<p>")
		sb.WriteString(p)
		sb.WriteString("</p>")
	}
	sb.WriteString("</div>")
	return &dpb.Narrative{Div: &dpb.Xhtml{Value: sb.String()}}
}

func (w *FHIRWriter) location(l *ir.PatientLocation) (*r4pb.Bundle_Entry, *dpb.Reference) {
	id := w.idGenerator.NewID()
	name := l.Name()

	entry := &r4pb.Bundle_Entry{
		Resource: &r4pb.ContainedResource{
			OneofResource: &r4pb.ContainedResource_Location{
				&locationpb.Location{
					Id:   &dpb.Id{Value: id},
					Name: &dpb.String{Value: name},
				},
			},
		},
	}

	ref := &dpb.Reference{
		Reference: &dpb.Reference_LocationId{
			&dpb.ReferenceId{Value: id},
		},
		Display: &dpb.String{Value: name},
	}

	return entry, ref
}

func (w *FHIRWriter) notes(notes []string) []*dpb.Annotation {
	var annotations []*dpb.Annotation
	for _, n := range notes {
		a := &dpb.Annotation{Text: &dpb.Markdown{Value: n}}
		annotations = append(annotations, a)
	}
	return annotations
}

func (w *FHIRWriter) dateTime(t ir.NullTime) *dpb.DateTime {
	if !t.Valid {
		return nil
	}
	return &dpb.DateTime{ValueUs: unixMicro(t.Time), Precision: dpb.DateTime_SECOND}
}

func unixMicro(t time.Time) int64 {
	return t.UnixNano() / microPerNano
}
