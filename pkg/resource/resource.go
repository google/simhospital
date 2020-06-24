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

	"google.golang.org/protobuf/encoding/prototext"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/constants"
	"github.com/google/simhospital/pkg/gender"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/logging"

	cpb "google/fhir/proto/r4/core/codes_go_proto"
	dpb "google/fhir/proto/r4/core/datatypes_go_proto"
	r4pb "google/fhir/proto/r4/core/resources/bundle_and_contained_resource_go_proto"
	encounterpb "google/fhir/proto/r4/core/resources/encounter_go_proto"
	patientpb "google/fhir/proto/r4/core/resources/patient_go_proto"
)

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

// GeneratorConfig is the configuration for resource generators.
type GeneratorConfig struct {
	Writer    io.Writer
	HL7Config *config.HL7Config
}

// NewFHIRWriter constructs and returns a new FHIRWriter.
func NewFHIRWriter(cfg GeneratorConfig) *FHIRWriter {
	return &FHIRWriter{
		gc:     gender.NewConvertor(cfg.HL7Config),
		writer: cfg.Writer,
	}
}

// FHIRWriter generates FHIR resources as protocol buffers, and writes them to writer.
type FHIRWriter struct {
	gc     gender.Convertor
	writer io.Writer
	count  int
}

// Generate generates FHIR resources from PatientInfo.
func (w *FHIRWriter) Generate(p *ir.PatientInfo) error {
	if p == nil {
		return errors.New("cannot generate resources from nil PatientInfo")
	}
	b := w.bundle(p)
	// TODO: Use jsonformat for output when available.
	m := prototext.MarshalOptions{Multiline: true}
	bytes, err := m.Marshal(b)
	if err != nil {
		return err
	}
	if _, err = w.writer.Write(bytes); err != nil {
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
		Entry: []*r4pb.Bundle_Entry{{
			Resource: &r4pb.ContainedResource{
				OneofResource: &r4pb.ContainedResource_Patient{
					w.patient(p),
				}},
		}},
	}

	for _, ec := range p.Encounters {
		encounter := &r4pb.Bundle_Entry{
			Resource: &r4pb.ContainedResource{
				OneofResource: &r4pb.ContainedResource_Encounter{
					w.encounter(ec),
				},
			},
		}
		bundle.Entry = append(bundle.Entry, encounter)
	}

	return bundle
}

func (w *FHIRWriter) patient(p *ir.PatientInfo) *patientpb.Patient {
	return &patientpb.Patient{
		Identifier: w.identifier(p.Person),
		Name:       w.humanName(p.Person),
		Gender:     w.gc.HL7ToFHIR(p.Person.Gender),
		Address:    w.address(p.Person.Address),
		Deceased:   w.deceased(p.Person),
		Telecom:    w.telecom(p.Person.PhoneNumber),
	}
}

func (w *FHIRWriter) identifier(pe *ir.Person) []*dpb.Identifier {
	return []*dpb.Identifier{{
		Value: &dpb.String{
			Value: pe.MRN,
		}},
	}
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

func (w *FHIRWriter) encounter(ec *ir.Encounter) *encounterpb.Encounter {
	return &encounterpb.Encounter{
		Status: &encounterpb.Encounter_StatusCode{
			Value: internalToFHIREncounterStatus[ec.Status],
		},
		Period: &dpb.Period{
			Start: w.dateTime(ec.Start),
			End:   w.dateTime(ec.End),
		},
		StatusHistory: w.statusHistory(ec.StatusHistory),
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

func (w *FHIRWriter) dateTime(time ir.NullTime) *dpb.DateTime {
	if !time.Valid {
		return nil
	}
	return &dpb.DateTime{ValueUs: time.Unix(), Precision: dpb.DateTime_SECOND}
}
