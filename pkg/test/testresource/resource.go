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

// Package testresource contains utilities for testing with resources.
package testresource

import (
	"bytes"
	"encoding/gob"
	"errors"

	"github.com/google/simhospital/pkg/ir"
)

// Writer is a resource writer that exposes the resources it has written to be retrieved by tests.
type Writer struct {
	enc       *gob.Encoder
	dec       *gob.Decoder
	Resources []*ir.PatientInfo
}

// Generate appends a copy of `p` to `Resources`.
func (w *Writer) Generate(p *ir.PatientInfo) error {
	if p == nil {
		return errors.New("PatientInfo is nil")
	}

	if err := w.enc.Encode(p); err != nil {
		return err
	}

	var pCopy ir.PatientInfo

	if err := w.dec.Decode(&pCopy); err != nil {
		return err
	}

	w.Resources = append(w.Resources, &pCopy)
	return nil
}

// Close exists to implement the hospital.ResourceWriter interface and is a no-op.
func (w *Writer) Close() error {
	return nil
}

// NewWriter initialises gob encoders and returns a new TestWriter.
func NewWriter() *Writer {
	var b bytes.Buffer
	return &Writer{enc: gob.NewEncoder(&b), dec: gob.NewDecoder(&b)}
}
