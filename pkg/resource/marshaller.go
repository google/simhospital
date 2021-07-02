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

package resource

import (
	"github.com/google/fhir/go/fhirversion"
	"github.com/google/fhir/go/jsonformat"
	"google.golang.org/protobuf/proto"
)

// JSONMarshaller is a resource.Marshaller that wraps jsonformat.Marshaller.
type JSONMarshaller struct {
	marshaller *jsonformat.Marshaller
}

// Marshal marshalls the given FHIR protocol buffer message as JSON.
func (m *JSONMarshaller) Marshal(message proto.Message) ([]byte, error) {
	return m.marshaller.MarshalResource(message)
}

// NewJSONMarshaller creates and returns a new JSONMarshaller.
func NewJSONMarshaller() (*JSONMarshaller, error) {
	m, err := jsonformat.NewPrettyMarshaller(fhirversion.R4)
	if err != nil {
		return nil, err
	}
	return &JSONMarshaller{marshaller: m}, nil
}
