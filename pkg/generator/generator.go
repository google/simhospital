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

// Package generator implements functionality to generate various patient related information, including, but not limited to:
// - person information, ie: name, surname, ethnicity, address, etc.,
// - patient type and class,
// - orders and test results,
// - allergies,
// - diagnosis,
// - procedures.
//
// The data is generated based on information provided in the pathway.
package generator

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/simhospital/pkg/clock"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/doctor"
	"github.com/google/simhospital/pkg/gender"
	"github.com/google/simhospital/pkg/generator/address"
	"github.com/google/simhospital/pkg/generator/codedelement"
	"github.com/google/simhospital/pkg/generator/document"
	"github.com/google/simhospital/pkg/generator/header"
	"github.com/google/simhospital/pkg/generator/id"
	"github.com/google/simhospital/pkg/generator/names"
	"github.com/google/simhospital/pkg/generator/notes"
	"github.com/google/simhospital/pkg/generator/order"
	"github.com/google/simhospital/pkg/generator/person"
	"github.com/google/simhospital/pkg/generator/text"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/logging"
	"github.com/google/simhospital/pkg/message"
	"github.com/google/simhospital/pkg/orderprofile"
	"github.com/google/simhospital/pkg/pathway"
	"github.com/google/simhospital/pkg/state"
)

var log = logging.ForCallerPackage()

type randomIDGenerator struct{}

func (g *randomIDGenerator) NewID() string {
	return fmt.Sprintf("%d", rand.Uint32())
}

// Generator implements functionality to generate various patient related information based on the information provided
// in the pathway.
type Generator struct {
	personGenerator       *person.Generator
	patientClassGenerator patientClassGenerator
	allergyGenerator      *codedelement.AllergyGenerator
	diagnosisGenerator    diagnosisOrProcedureGenerator
	procedureGenerator    diagnosisOrProcedureGenerator
	messageConfig         *config.HL7Config
	doctors               *doctor.Doctors
	headerGenerator       *header.Generator
	orderGenerator        *order.Generator
	documentGenerator     *document.Generator
}

type diagnosisOrProcedureGenerator interface {
	RandomOrFromPathway(*pathway.DateTime, *pathway.DiagnosisOrProcedure) *ir.DiagnosisOrProcedure
}

// NewPerson returns a new person based on pathway.Person.
func (g Generator) NewPerson(pathwayPerson *pathway.Person) *ir.Person {
	return g.personGenerator.NewPerson(pathwayPerson)
}

// UpdateFromPathway updates PatientInfo with information from pathway.
// It Updates:
// - person information
// - diagnoses
// - procedures
// - allergies
func (g Generator) UpdateFromPathway(patientInfo *ir.PatientInfo, updatePerson *pathway.UpdatePerson) {
	if updatePerson.Person != nil {
		g.personGenerator.UpdatePersonFromPathway(patientInfo.Person, updatePerson.Person)
	}
	g.setDiagnoses(patientInfo, updatePerson.Diagnoses)
	g.setProcedures(patientInfo, updatePerson.Procedures)
	g.AddAllergies(patientInfo, updatePerson.Allergies)
}

// NewPatient returns a new patient based on Person information and a doctor provided.
func (g Generator) NewPatient(person *ir.Person, doctor *ir.Doctor) *state.Patient {
	p := &state.Patient{
		PatientInfo: &ir.PatientInfo{
			Class:  g.messageConfig.PatientClass.Outpatient,
			Person: person,
			// The Hospital Service might be overridden later with the doctor's specialty.
			HospitalService: g.messageConfig.HospitalService,
			AttendingDoctor: doctor,
		},
		// The code downstream assumes that Orders exists.
		Orders:    make(map[string]*ir.Order),
		Documents: make(map[string]*ir.Document),
	}
	// If none of the g.messageConfig.PrimaryFacility fields is set, we want the resulting HL7 message to have the entire
	// PD1.3 Patient Primary Facility field empty. This is achieved by leaving p.PatientInfo.PrimaryFacility nil.
	if g.messageConfig.PrimaryFacility != nil {
		p.PatientInfo.PrimaryFacility = &ir.PrimaryFacility{
			Organization: g.messageConfig.PrimaryFacility.OrganizationName,
			ID:           g.messageConfig.PrimaryFacility.IDNumber,
		}
	}
	if doctor != nil {
		docWithSpecialty := g.doctors.GetByID(doctor.ID)
		if docWithSpecialty != nil && docWithSpecialty.Specialty != "" {
			p.PatientInfo.HospitalService = docWithSpecialty.Specialty
		}
	}
	return p
}

// NewDoctor returns a new doctor based on the Consultant information from the pathway.
// If consultant is not specified, it returns a random doctor.
// Otherwise, it attempts to lookup an existic doctor basd on consultant ID. If any doctor is found, it returns it.
// Othwerise creates a new doctor from Consultant information, with the default speciality defined in
// messageConfig.HospitalService.
func (g Generator) NewDoctor(c *pathway.Consultant) *ir.Doctor {
	if c == nil {
		return g.doctors.GetRandomDoctor()
	}
	if doctor := g.doctors.GetByID(*c.ID); doctor != nil {
		return doctor
	}
	newDoctor := &ir.Doctor{
		// A valid pathway.Consultant has all the fields set, so we can just dereference.
		ID:        *c.ID,
		Surname:   *c.Surname,
		Prefix:    *c.Prefix,
		FirstName: *c.FirstName,
		Specialty: g.messageConfig.HospitalService,
	}
	g.doctors.Add(newDoctor)
	return newDoctor
}

// ResetPatient returns a Patient based on the given Patient.
// Medical History (Orders, Encounters) and general information is kept, but other
// information is cleared as if the patient was a new patient.
func (g Generator) ResetPatient(p *state.Patient) *state.Patient {
	newP := g.NewPatient(p.PatientInfo.Person, p.PatientInfo.AttendingDoctor)
	newP.Orders = p.Orders
	newP.PatientInfo.HospitalService = p.PatientInfo.HospitalService
	newP.PatientInfo.Encounters = p.PatientInfo.Encounters
	newP.PastVisits = p.PastVisits
	newP.PatientInfo.PrimaryFacility = p.PatientInfo.PrimaryFacility
	newP.PatientInfo.Allergies = p.PatientInfo.Allergies
	return newP
}

// AddAllergies adds allergies specified in the pathway to the patientInfo:
// * If there are any allergies specified in the pathways, they are always added to existing allergies on the patientInfo.
// * If the allergies were not specified in the pathway (ie. allergies is nil) and the allergies on the patientInfo
//   have not been initialised yet (ie are also nil), initialise them to an empty slice (to make sure we'll not make an
//   attempt to generate them on the next ADT-like event, as that would increase the likelihood of the patient having
//   allergies), and then generate them.
// * If the allergies from the pathway are explicitly set to empty slice, the allergies on the patient info are also set
//   to empty slice.
func (g Generator) AddAllergies(patientInfo *ir.PatientInfo, allergies []pathway.Allergy) {
	switch {
	case len(allergies) > 0:
		// If the pathway allergies are set, add them to the existing ones.
		if patientInfo.Allergies == nil {
			patientInfo.Allergies = []*ir.Allergy{}
		}
		patientInfo.Allergies = append(patientInfo.Allergies, g.getDedupedAllergiesFromPathway(patientInfo, allergies)...)
	case allergies == nil && patientInfo.Allergies == nil:
		// Initialise the allergies to an empty slice so that they're not nil anymore.
		patientInfo.Allergies = []*ir.Allergy{}
		patientInfo.Allergies = append(patientInfo.Allergies, g.allergyGenerator.GenerateRandomDistinctAllergies()...)
	case allergies != nil && len(allergies) == 0:
		// Allergies were set explicitly as an empty slice in the pathway.
		patientInfo.Allergies = []*ir.Allergy{}
	}
}

// getDedupedAllergiesFromPathway returns the list of allergies from the pathway after de-duplication:
// if the allergy is set twice in the pathway, it's added to the list only once. If the allergy
// already exists for the patient, it's not added to the list.
// Note: if the same Allergy is specified with eg. different severity or reaction, it'll be added to
// the list, as there is no way of deleting / amending existing pathwayAllergies.
func (g Generator) getDedupedAllergiesFromPathway(patientInfo *ir.PatientInfo, pathwayAllergies []pathway.Allergy) []*ir.Allergy {
	var dedupedAllergies []*ir.Allergy
	existing := make(map[ir.Allergy]bool)
	for _, a := range patientInfo.Allergies {
		existing[*a] = true
	}

	for _, a := range pathwayAllergies {
		code, description := g.allergyGenerator.DeriveCodeAndDescription(a.Code, a.Description)
		idt := g.allergyGenerator.DeriveIdentificationDateTime(a)
		cs := g.allergyGenerator.DeriveCodingSystem(g.messageConfig.Allergy, a)
		allergy := &ir.Allergy{
			Type: a.Type,
			Description: ir.CodedElement{
				ID:           code,
				Text:         description,
				CodingSystem: cs,
			},
			Severity:               a.Severity,
			Reaction:               a.Reaction,
			IdentificationDateTime: idt,
		}
		if !existing[*allergy] {
			existing[*allergy] = true
			dedupedAllergies = append(dedupedAllergies, allergy)
		}
	}
	return dedupedAllergies
}

func (g Generator) setDiagnoses(patientInfo *ir.PatientInfo, diagnoses []*pathway.DiagnosisOrProcedure) {
	patientInfo.Diagnoses = make([]*ir.DiagnosisOrProcedure, len(diagnoses))
	g.setDiagnosesOrProcedures(patientInfo.Diagnoses, diagnoses, g.diagnosisGenerator)
}

func (g Generator) setProcedures(patientInfo *ir.PatientInfo, procedures []*pathway.DiagnosisOrProcedure) {
	patientInfo.Procedures = make([]*ir.DiagnosisOrProcedure, len(procedures))
	g.setDiagnosesOrProcedures(patientInfo.Procedures, procedures, g.procedureGenerator)
}

func (g Generator) setDiagnosesOrProcedures(diagnosisOrProcedure []*ir.DiagnosisOrProcedure, fromPathway []*pathway.DiagnosisOrProcedure, dpg diagnosisOrProcedureGenerator) {
	for i, p := range fromPathway {
		diagnosisOrProcedure[i] = dpg.RandomOrFromPathway(p.DateTime, p)
		// By design, diagnoses and procedures don't reuse the clinician from the pathway.
		// Past diagnoses and procedures could have been done by other clinicians, not the current one,
		// so we do not want to use the pathway's clinician.
		diagnosisOrProcedure[i].Clinician = g.NewDoctor(nil)
	}
}

// NewRegistrationPatientClassAndType returns a PatientClassAndType for a patient newly registered.
func (g Generator) NewRegistrationPatientClassAndType() *config.PatientClassAndType {
	return g.patientClassGenerator.Random()
}

// NewOrder returns a new order based on order information from the pathway and eventTime.
func (g Generator) NewOrder(o *pathway.Order, eventTime time.Time) *ir.Order {
	return g.orderGenerator.NewOrder(o, eventTime)
}

// OrderWithClinicalNote creates an order with a Clinical Note based on the pathway.
func (g Generator) OrderWithClinicalNote(ctx context.Context, o *ir.Order, n *pathway.ClinicalNote, eventTime time.Time) (*ir.Order, error) {
	return g.orderGenerator.OrderWithClinicalNote(ctx, o, n, eventTime)
}

// SetResults sets results on an existing Order based on the results information from the pathway.
// If order is nil, this also creates an Order using details in pathway.Result.
// Returns an error of the retults cannot be created.
func (g Generator) SetResults(o *ir.Order, r *pathway.Results, eventTime time.Time) (*ir.Order, error) {
	return g.orderGenerator.SetResults(o, r, eventTime)
}

// NewVisitID generates a new visit identifier.
func (g Generator) NewVisitID() uint64 {
	return rand.Uint64()
}

// NewHeader returns a new header for the given step.
func (g *Generator) NewHeader(step *pathway.Step) *message.HeaderInfo {
	return g.headerGenerator.NewHeader(step)
}

// NewDocument returns a NewDocument for MDM^T02 messages.
func (g Generator) NewDocument(eventTime time.Time, d *pathway.Document) *ir.Document {
	return g.documentGenerator.Document(eventTime, d)
}

// UpdateDocumentContent updates the given document for MDM^T02 messages.
func (g Generator) UpdateDocumentContent(dm *ir.Document, dp *pathway.Document) error {
	return g.documentGenerator.UpdateDocumentContent(dm, dp)
}

// Config contains the configuration for Generator.
type Config struct {
	Clock            clock.Clock
	HL7Config        *config.HL7Config
	Header           *config.Header
	AddressGenerator person.AddressGenerator
	MRNGenerator     id.Generator
	PlacerGenerator  id.Generator
	FillerGenerator  id.Generator
	textGenerator    text.Generator
	NotesGenerator   order.NotesGenerator
	DateGenerator    codedelement.DateGenerator
	Data             *config.Data
	Doctors          *doctor.Doctors
	MsgCtrlGenerator *header.MessageControlGenerator
	OrderProfiles    *orderprofile.OrderProfiles
}

// NewGenerator creates a new Generator.
func NewGenerator(cfg Config) *Generator {
	ag := cfg.AddressGenerator
	if ag == nil {
		ag = &address.Generator{Nouns: cfg.Data.Nouns, Address: cfg.Data.Address, PostcodeGenerator: &address.UKPostcode{}}
	}

	mrnGenerator := cfg.MRNGenerator
	if mrnGenerator == nil {
		mrnGenerator = &randomIDGenerator{}
	}

	placerGenerator := cfg.PlacerGenerator
	if placerGenerator == nil {
		placerGenerator = &randomIDGenerator{}
	}

	fillerGenerator := cfg.FillerGenerator
	if fillerGenerator == nil {
		fillerGenerator = &randomIDGenerator{}
	}

	tg := cfg.textGenerator
	if tg == nil {
		tg = &text.NounGenerator{Nouns: cfg.Data.Nouns}
	}

	ng := cfg.NotesGenerator
	if ng == nil {
		ng = notes.NewGenerator(cfg.Data, tg)
	}

	dg := cfg.DateGenerator
	if dg == nil {
		dg = &codedelement.SimpleDateGenerator{}
	}

	personGenerator := &person.Generator{
		Clock:              cfg.Clock,
		NameGenerator:      &names.Generator{Data: cfg.Data},
		GenderConvertor:    gender.NewConvertor(cfg.HL7Config),
		EthnicityGenerator: person.NewEthnicityGenerator(cfg.Data),
		AddressGenerator:   ag,
		MRNGenerator:       mrnGenerator,
	}

	orderGenerator := &order.Generator{
		MessageConfig:         cfg.HL7Config,
		OrderProfiles:         cfg.OrderProfiles,
		NoteGenerator:         ng,
		PlacerGenerator:       placerGenerator,
		FillerGenerator:       fillerGenerator,
		AbnormalFlagConvertor: order.NewAbnormalFlagConvertor(cfg.HL7Config),
		Doctors:               cfg.Doctors,
	}

	return &Generator{
		personGenerator:       personGenerator,
		patientClassGenerator: newPatientClassAndTypeGenerator(cfg.Data),
		messageConfig:         cfg.HL7Config,
		doctors:               cfg.Doctors,
		allergyGenerator:      codedelement.NewAllergyGenerator(cfg.HL7Config, cfg.Data, cfg.Clock, dg),
		diagnosisGenerator:    codedelement.NewDiagnosisGenerator(cfg.HL7Config, cfg.Data, cfg.Clock, dg),
		procedureGenerator:    codedelement.NewProcedureGenerator(cfg.HL7Config, cfg.Data, cfg.Clock, dg),
		headerGenerator:       &header.Generator{Header: cfg.Header, MsgCtrlGen: cfg.MsgCtrlGenerator},
		orderGenerator:        orderGenerator,
		documentGenerator:     &document.Generator{DocumentConfig: &cfg.HL7Config.Document, TextGenerator: tg},
	}
}
