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
	"io"

	dpb "google/fhir/proto/r4/core/datatypes_go_proto"
	r4pb "google/fhir/proto/r4/core/resources/bundle_and_contained_resource_go_proto"
	patientpb "google/fhir/proto/r4/core/resources/patient_go_proto"
	"google.golang.org/protobuf/encoding/prototext"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/logging"
)

var log = logging.ForCallerPackage()

// GeneratorConfig is the configuration for resource generators.
// This currently only contains Writer, but more fields will be added as necessary e.g. HL7Config
// for convertors and generators.
type GeneratorConfig struct {
	Writer io.Writer
}

// NewFHIRWriter constructs and returns a new FHIRWriter.
// This currently only sets the fields of FHIRWriter, but setup code will be added as necessary
// e.g. constructing a convertor.
func NewFHIRWriter(cfg GeneratorConfig) *FHIRWriter {
	return &FHIRWriter{
		writer: cfg.Writer,
	}
}

// FHIRWriter generates FHIR resources as protocol buffers, and writes them to writer.
type FHIRWriter struct {
	writer io.Writer
	count  int
}

// Generate generates FHIR resources from PatientInfo.
func (w *FHIRWriter) Generate(p *ir.PatientInfo) error {
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
	return &r4pb.Bundle{
		Entry: []*r4pb.Bundle_Entry{{
			Resource: &r4pb.ContainedResource{
				OneofResource: &r4pb.ContainedResource_Patient{
					w.patient(p),
				}},
		}},
	}
}

func (w *FHIRWriter) patient(p *ir.PatientInfo) *patientpb.Patient {
	return &patientpb.Patient{
		Identifier: w.identifier(p.Person),
		Name:       w.humanName(p.Person),
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
