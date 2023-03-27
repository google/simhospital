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

// Package fhir contains functionality for generating FHIR resources from PatientInfo.
package fhir

import (
	"io"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/gender"
	"github.com/google/simhospital/pkg/generator/codedelement"
	"github.com/google/simhospital/pkg/generator/id"
	"github.com/google/simhospital/pkg/generator/order"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/logging"

	cpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/codes_go_proto"
	dpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
	r4pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/resources/bundle_and_contained_resource_go_proto"
)

const microPerNano = int64(time.Microsecond / time.Nanosecond)

var log = logging.ForCallerPackage()

// Marshaller defines an object that can marshal a protocol buffer message.
type Marshaller interface {
	Marshal(proto.Message) ([]byte, error)
}

// Output defines an object which returns a writer that resources are written to.
type Output interface {
	New(string) (io.WriteCloser, error)
}

// BundlerConfig is the configuration for resource generators.
type BundlerConfig struct {
	HL7Config   *config.HL7Config
	IDGenerator id.Generator
	// BundleType is the type of bundle to generate, and defaults to Batch if unspecified.
	BundleType string
}

// NewBundler constructs and returns a new Bundler.
func NewBundler(cfg BundlerConfig) (*Bundler, error) {
	ac, err := codedelement.NewAllergyConvertor(cfg.HL7Config)
	if err != nil {
		return nil, err
	}

	bundleTypeCode, err := bundleType(cfg.BundleType)
	if err != nil {
		return nil, err
	}

	return &Bundler{
		gc:             gender.NewConvertor(cfg.HL7Config),
		oc:             order.NewConvertor(cfg.HL7Config),
		ac:             ac,
		cc:             codedelement.NewCodingSystemConvertor(cfg.HL7Config),
		idGenerator:    cfg.IDGenerator,
		locations:      make(map[ir.PatientLocation]*dpb.Reference),
		doctors:        make(map[ir.Doctor]*dpb.Reference),
		bundleTypeCode: bundleTypeCode,
	}, nil
}

// Bundler generates FHIR resources as protocol buffers.
type Bundler struct {
	gc          gender.Convertor
	oc          order.Convertor
	ac          codedelement.AllergyConvertor
	cc          codedelement.CodingSystemConvertor
	idGenerator id.Generator
	// locationMap and doctorMap ensure that equivalent locations and doctors are only generated
	// once, preventing duplicates.
	locations      map[ir.PatientLocation]*dpb.Reference
	doctors        map[ir.Doctor]*dpb.Reference
	bundleTypeCode cpb.BundleTypeCode_Value
}

// Writer writes FHIR resources protocol buffers.
type Writer struct {
	Bundler    *Bundler
	count      int
	Output     Output
	Marshaller Marshaller
}

// Generate generates FHIR resources from PatientInfo.
func (w *Writer) Generate(p *ir.PatientInfo) error {
	b, err := w.Bundler.Generate(p)
	if err != nil {
		return err
	}

	pe := p.Person
	filename := strings.Join([]string{pe.FirstName, pe.MiddleName, pe.Surname, pe.MRN}, "_")
	return w.writeBundle(filename, b)
}

func (w *Writer) writeBundle(filename string, b *r4pb.Bundle) error {
	f, err := w.Output.New(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	bytes, err := w.Marshaller.Marshal(b)
	if err != nil {
		return err
	}
	if _, err = f.Write(bytes); err != nil {
		return err
	}
	w.count = w.count + len(b.GetEntry()) + 1 // Include the Bundle resource itself.
	return nil
}

// Close closes the Writer.
func (w *Writer) Close() error {
	log.Infof("FHIR Resources successfully written by the FHIR Writer: %d", w.count)
	return nil
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
