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
	"fmt"
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

	cpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/codes_go_proto"
	dpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
	aipb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/allergy_intolerance_go_proto"
	r4pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/bundle_and_contained_resource_go_proto"
	conditionpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/condition_go_proto"
	encounterpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/encounter_go_proto"
	locationpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/location_go_proto"
	observationpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/observation_go_proto"
	patientpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/patient_go_proto"
	practitionerpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/practitioner_go_proto"
	procedurepb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/procedure_go_proto"
)

const (
	// Batch denotes the transaction bundle type: intended to be processed by a server as a group of
	// independent actions.
	// Reference: http://hl7.org/fhir/valueset-bundle-type.html
	Batch = "BATCH"
	// Collection denotes the collection bundle type: a set of resources collected into a single
	// document for ease of distribution.
	// Reference: http://hl7.org/fhir/valueset-bundle-type.html
	Collection   = "COLLECTION"
	microPerNano = int64(time.Microsecond / time.Nanosecond)
)

var (
	log = logging.ForCallerPackage()

	bundleTypes = map[string]cpb.BundleTypeCode_Value{
		Batch:      cpb.BundleTypeCode_BATCH,
		Collection: cpb.BundleTypeCode_COLLECTION,
		"":         cpb.BundleTypeCode_BATCH,
	}

	// Default value for cpb.AddressUseCode_Value is AddressUseCode_INVALID_UNINITIALIZED.
	internalToFHIRAddressType = map[string]cpb.AddressUseCode_Value{
		"HOME": cpb.AddressUseCode_HOME,
		"WORK": cpb.AddressUseCode_WORK,
	}

	// Default value for cpb.EncounterStatusCode is EncounterStatusCode_INVALID_UNINITIALIZED.
	internalToFHIREncounterStatus = map[string]cpb.EncounterStatusCode_Value{
		constants.EncounterStatusPlanned:    cpb.EncounterStatusCode_PLANNED,
		constants.EncounterStatusInProgress: cpb.EncounterStatusCode_IN_PROGRESS,
		constants.EncounterStatusArrived:    cpb.EncounterStatusCode_ARRIVED,
		constants.EncounterStatusFinished:   cpb.EncounterStatusCode_FINISHED,
		constants.EncounterStatusCancelled:  cpb.EncounterStatusCode_CANCELLED,
		constants.EncounterStatusUnknown:    cpb.EncounterStatusCode_UNKNOWN,
	}
)

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
	// BundleType is the type of bundle to generate, and defaults to Batch if unspecified.
	BundleType string
}

// NewFHIRWriter constructs and returns a new FHIRWriter.
func NewFHIRWriter(cfg GeneratorConfig) (*FHIRWriter, error) {
	ac, err := codedelement.NewAllergyConvertor(cfg.HL7Config)
	if err != nil {
		return nil, err
	}

	bundleTypeCode, err := bundleType(cfg.BundleType)
	if err != nil {
		return nil, err
	}

	return &FHIRWriter{
		gc:             gender.NewConvertor(cfg.HL7Config),
		oc:             order.NewConvertor(cfg.HL7Config),
		ac:             ac,
		cc:             codedelement.NewCodingSystemConvertor(cfg.HL7Config),
		idGenerator:    cfg.IDGenerator,
		output:         cfg.Output,
		marshaller:     cfg.Marshaller,
		locations:      make(map[ir.PatientLocation]*dpb.Reference),
		doctors:        make(map[ir.Doctor]*dpb.Reference),
		bundleTypeCode: bundleTypeCode,
	}, nil
}

func bundleType(bundleType string) (cpb.BundleTypeCode_Value, error) {
	if bundleTypeCode, ok := bundleTypes[bundleType]; ok {
		return bundleTypeCode, nil
	}
	return cpb.BundleTypeCode_INVALID_UNINITIALIZED,
		fmt.Errorf("invalid bundle type, expected one of %+v", keys(bundleTypes))
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
	// locationMap and doctorMap ensure that equivalent locations and doctors are only generated
	// once, preventing duplicates.
	locations      map[ir.PatientLocation]*dpb.Reference
	doctors        map[ir.Doctor]*dpb.Reference
	bundleTypeCode cpb.BundleTypeCode_Value
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
			Value: cpb.BundleTypeCode_BATCH,
		},
	}

	bundle.Type = &r4pb.Bundle_TypeCode{Value: w.bundleTypeCode}

	patient, patientRef := w.patient(p.Person)
	addEntry(bundle, patient)

	allergies := w.allergies(p.Allergies, patientRef)
	addEntry(bundle, allergies...)

	for _, ec := range p.Encounters {
		encounter, encounterRef := w.encounter(ec, p.Class)

		e := encounter.GetResource().GetEncounter()
		for _, lh := range ec.LocationHistory {
			location, locationRef := w.location(lh.Location)
			addEntry(bundle, location)
			e.Location = append(e.Location, w.encounterLocation(locationRef, lh.Start, lh.End))
		}

		for _, pr := range ec.Procedures {
			practitioner, practitionerRef := w.practitioner(pr.Clinician)
			addEntry(bundle, practitioner)

			procedure, procedureRef := w.procedure(pr, patientRef, practitionerRef, encounterRef)
			addEntry(bundle, procedure)
			e.Diagnosis = append(e.Diagnosis, w.encounterDiagnosis(procedureRef))
		}

		for _, d := range ec.Diagnoses {
			practitioner, practitionerRef := w.practitioner(d.Clinician)
			addEntry(bundle, practitioner)

			condition, conditionRef := w.condition(d, patientRef, practitionerRef, encounterRef)
			addEntry(bundle, condition)
			e.Diagnosis = append(e.Diagnosis, w.encounterDiagnosis(conditionRef))
		}
		addEntry(bundle, encounter)

		for _, o := range ec.Orders {
			observations := w.observations(encounterRef, patientRef, o)
			addEntry(bundle, observations...)
		}
	}
	return bundle
}

func addEntry(bundle *r4pb.Bundle, entries ...*r4pb.Bundle_Entry) {
	for _, entry := range entries {
		if entry == nil {
			continue
		}
		bundle.Entry = append(bundle.Entry, entry)
	}
}

func (w *FHIRWriter) patient(person *ir.Person) (*r4pb.Bundle_Entry, *dpb.Reference) {
	id := w.idGenerator.NewID()

	entry := &r4pb.Bundle_Entry{
		Resource: &r4pb.ContainedResource{
			OneofResource: &r4pb.ContainedResource_Patient{
				&patientpb.Patient{
					Id:         &dpb.Id{Value: id},
					Identifier: w.identifier(person.MRN),
					Name:       w.humanName(person),
					Address:    w.address(person.Address),
					Deceased:   w.deceased(person),
					Telecom:    w.telecom(person.PhoneNumber),
					Gender: &patientpb.Patient_GenderCode{
						Value: w.gc.HL7ToFHIR(person.Gender),
					},
					Text: w.narrative(person.Text()),
				},
			},
		},
	}

	ref := &dpb.Reference{
		Reference: &dpb.Reference_PatientId{&dpb.ReferenceId{Value: id}},
		Display:   &dpb.String{Value: person.AlternateText()},
	}

	return w.addURL(entry, id, "Patient"), ref
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
						// Simulated Hospital does not support the concept of ClinicalStatus, so we default to
						// a hardcoded "active" value.
						ClinicalStatus: &dpb.CodeableConcept{
							Coding: []*dpb.Coding{{
								Code: &dpb.Code{Value: "active"},
								System: &dpb.Uri{
									Value: "http://terminology.hl7.org/CodeSystem/allergyintolerance-clinical",
								},
								Display: &dpb.String{Value: "Active"},
							}},
						},
						// Simulated Hospital does not yet distinguish between allergies and intolerances.
						Type: &aipb.AllergyIntolerance_TypeCode{Value: cpb.AllergyIntoleranceTypeCode_ALLERGY},
						Category: []*aipb.AllergyIntolerance_CategoryCode{{
							Value: w.ac.TypeHL7ToFHIR(a.Type),
						}},
						Reaction: []*aipb.AllergyIntolerance_Reaction{{
							Manifestation: []*dpb.CodeableConcept{{
								Text: &dpb.String{Value: a.Reaction},
							}},
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
		entries = append(entries, w.addURL(entry, id, "AllergyIntolerance"))
	}
	return entries
}

func (w *FHIRWriter) codeableConcept(c ir.CodedElement) *dpb.CodeableConcept {
	return &dpb.CodeableConcept{
		// The Text field should only be used if the code and coding system are unknown.
		Coding: []*dpb.Coding{{
			System:  &dpb.Uri{Value: w.cc.HL7ToFHIR(c.CodingSystem)},
			Code:    &dpb.Code{Value: c.ID},
			Display: &dpb.String{Value: c.Text},
		}},
	}
}

func (w *FHIRWriter) identifier(identifier string) []*dpb.Identifier {
	return []*dpb.Identifier{{Value: &dpb.String{Value: identifier}}}
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

func (w *FHIRWriter) encounter(encounter *ir.Encounter, class string) (*r4pb.Bundle_Entry, *dpb.Reference) {
	id := w.idGenerator.NewID()

	entry := &r4pb.Bundle_Entry{
		Resource: &r4pb.ContainedResource{
			OneofResource: &r4pb.ContainedResource_Encounter{
				&encounterpb.Encounter{
					Id:   &dpb.Id{Value: id},
					Text: w.narrative(encounter.Text()),
					ClassValue: &dpb.Coding{
						Code: &dpb.Code{Value: class},
					},
					Status: &encounterpb.Encounter_StatusCode{
						Value: internalToFHIREncounterStatus[encounter.Status],
					},
					Period: &dpb.Period{
						Start: w.dateTime(encounter.Start),
						End:   w.dateTime(encounter.End),
					},
					StatusHistory: w.statusHistory(encounter.StatusHistory),
				},
			},
		},
	}

	ref := &dpb.Reference{
		Reference: &dpb.Reference_EncounterId{&dpb.ReferenceId{Value: id}},
	}

	return w.addURL(entry, id, "Encounter"), ref
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

// encounterDiagnosis constructs the `Encounter.Diagnosis` field. `ref` must be a reference to
// either a condition or procedure.
// Reference: https://www.hl7.org/fhir/encounter-definitions.html#Encounter.diagnosis.condition
func (w *FHIRWriter) encounterDiagnosis(ref *dpb.Reference) *encounterpb.Encounter_Diagnosis {
	return &encounterpb.Encounter_Diagnosis{
		Condition: ref,
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
		o := &observationpb.Observation{
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
		}

		if r.TestName != nil {
			o.Code = w.codeableConcept(*r.TestName)
		}

		entry := &r4pb.Bundle_Entry{
			Resource: &r4pb.ContainedResource{
				OneofResource: &r4pb.ContainedResource_Observation{o},
			},
		}

		observations = append(observations, w.addURL(entry, id, "Observation"))
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
		for _, s := range strings.Split(p, "\n") {
			fmt.Fprintf(&sb, "<p>%s</p>", s)
		}
	}
	sb.WriteString("</div>")
	return &dpb.Narrative{
		Div:    &dpb.Xhtml{Value: sb.String()},
		Status: &dpb.Narrative_StatusCode{Value: cpb.NarrativeStatusCode_GENERATED},
	}
}

func (w *FHIRWriter) location(location *ir.PatientLocation) (*r4pb.Bundle_Entry, *dpb.Reference) {
	if location == nil {
		return nil, nil
	}
	if ref, ok := w.locations[*location]; ok {
		return nil, ref
	}

	id := w.idGenerator.NewID()
	name := location.Name()

	entry := &r4pb.Bundle_Entry{
		Resource: &r4pb.ContainedResource{
			OneofResource: &r4pb.ContainedResource_Location{
				&locationpb.Location{
					Id:   &dpb.Id{Value: id},
					Name: &dpb.String{Value: name},
					Text: w.narrative(name),
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

	w.locations[*location] = ref

	return w.addURL(entry, id, "Location"), ref
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

func (w *FHIRWriter) procedure(procedure *ir.DiagnosisOrProcedure, patientRef *dpb.Reference, practitionerRef *dpb.Reference, encounterRef *dpb.Reference) (*r4pb.Bundle_Entry, *dpb.Reference) {
	id := w.idGenerator.NewID()
	p := &procedurepb.Procedure{
		Id: &dpb.Id{Value: id},
		Performed: &procedurepb.Procedure_PerformedX{
			Choice: &procedurepb.Procedure_PerformedX_DateTime{
				w.dateTime(procedure.DateTime),
			},
		},
		Status: &procedurepb.Procedure_StatusCode{Value: cpb.EventStatusCode_COMPLETED},
		Category: &dpb.CodeableConcept{
			Text: &dpb.String{Value: procedure.Type},
		},
		Encounter: encounterRef,
		Performer: []*procedurepb.Procedure_Performer{{
			Actor: practitionerRef,
		}},
		Text:    w.narrative(procedure.Text()),
		Subject: patientRef,
	}

	if procedure.Description != nil {
		p.Code = w.codeableConcept(*procedure.Description)
	}

	entry := &r4pb.Bundle_Entry{
		Resource: &r4pb.ContainedResource{
			OneofResource: &r4pb.ContainedResource_Procedure{p},
		},
	}

	ref := &dpb.Reference{
		Reference: &dpb.Reference_ProcedureId{
			&dpb.ReferenceId{Value: id},
		},
		Display: &dpb.String{Value: procedure.Text()},
	}

	return w.addURL(entry, id, "Procedure"), ref
}

func (w *FHIRWriter) condition(diagnosis *ir.DiagnosisOrProcedure, patientRef *dpb.Reference, practitionerRef *dpb.Reference, encounterRef *dpb.Reference) (*r4pb.Bundle_Entry, *dpb.Reference) {
	id := w.idGenerator.NewID()

	d := &conditionpb.Condition{
		Id:           &dpb.Id{Value: id},
		RecordedDate: w.dateTime(diagnosis.DateTime),
		Recorder:     practitionerRef,
		Encounter:    encounterRef,
		Text:         w.narrative(diagnosis.Text()),
		Subject:      patientRef,
	}

	if diagnosis.Description != nil {
		d.Code = w.codeableConcept(*diagnosis.Description)
	}

	entry := &r4pb.Bundle_Entry{
		Resource: &r4pb.ContainedResource{
			OneofResource: &r4pb.ContainedResource_Condition{d},
		},
	}

	ref := &dpb.Reference{
		Reference: &dpb.Reference_ConditionId{
			&dpb.ReferenceId{Value: id},
		},
		Display: &dpb.String{Value: diagnosis.Text()},
	}

	return w.addURL(entry, id, "Condition"), ref
}

func (w *FHIRWriter) practitioner(doctor *ir.Doctor) (*r4pb.Bundle_Entry, *dpb.Reference) {
	if doctor == nil {
		return nil, nil
	}
	if ref, ok := w.doctors[*doctor]; ok {
		return nil, ref
	}

	id := w.idGenerator.NewID()
	person := &ir.Person{
		Prefix:    doctor.Prefix,
		FirstName: doctor.FirstName,
		Surname:   doctor.Surname,
	}

	entry := &r4pb.Bundle_Entry{
		Resource: &r4pb.ContainedResource{
			OneofResource: &r4pb.ContainedResource_Practitioner{
				&practitionerpb.Practitioner{
					Id:         &dpb.Id{Value: id},
					Identifier: w.identifier(doctor.ID),
					Name:       w.humanName(person),
					Text:       w.narrative(person.Text()),
				},
			},
		},
	}

	ref := &dpb.Reference{
		Reference: &dpb.Reference_PractitionerId{
			&dpb.ReferenceId{Value: id},
		},
		Display: &dpb.String{Value: person.AlternateText()},
	}

	w.doctors[*doctor] = ref

	return w.addURL(entry, id, "Practitioner"), ref
}

// addURL adds the FullURL field to the resource, and if the bundle type is set to Batch the
// Request field is also set to provide execution information for the server. `url` is the HTTP URL
// for the resource, and is usually the resource type. addURL should only be called from internal
// methods where `entry` has already been constructed via a struct literal.
func (w *FHIRWriter) addURL(entry *r4pb.Bundle_Entry, id, url string) *r4pb.Bundle_Entry {
	if w.bundleTypeCode == cpb.BundleTypeCode_BATCH {
		entry.Request = w.request(url)
	}
	entry.FullUrl = &dpb.Uri{Value: fmt.Sprintf("%s/%s", url, id)}
	return entry
}

func (w *FHIRWriter) request(url string) *r4pb.Bundle_Entry_Request {
	return &r4pb.Bundle_Entry_Request{
		Url: &dpb.Uri{Value: url},
		// Currently, we only support the creation of resources (POST).
		Method: &r4pb.Bundle_Entry_Request_MethodCode{
			Value: cpb.HTTPVerbCode_POST,
		},
	}
}

func unixMicro(t time.Time) int64 {
	return t.UnixNano() / microPerNano
}

func keys(m map[string]cpb.BundleTypeCode_Value) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
