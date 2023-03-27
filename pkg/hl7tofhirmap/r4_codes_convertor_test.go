// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package hl7tofhirmap contains utility functions to convert HL7v2 to FHIR.
package hl7tofhirmap

import (
	"testing"

	cpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/codes_go_proto"
)

func TestConvertor_AbstractTypeCode_Default(t *testing.T) {
	c := &Convertor{}
	if got, want := c.AbstractTypeCode("TYPE"), cpb.AbstractTypeCode_TYPE; got != want {
		t.Errorf(`AbstractTypeCode("TYPE") got %s, want %s`, got, want)
	}
}

func TestConvertor_AbstractTypeCode_Default_NotExists(t *testing.T) {
	c := &Convertor{}
	if got, want := c.AbstractTypeCode("non-existent"), cpb.AbstractTypeCode_INVALID_UNINITIALIZED; got != want {
		t.Errorf(`AbstractTypeCode("non-existent") got %s, want %s`, got, want)
	}
}

func TestConvertor_AccountStatusCode_Default_NotExistsButHasUnknown(t *testing.T) {
	c := &Convertor{}
	if got, want := c.AccountStatusCode("non-existent"), cpb.AccountStatusCode_INVALID_UNINITIALIZED; got != want {
		t.Errorf(`AccountStatusCode("non-existent") got %s, want %s`, got, want)
	}
}

func TestConvertor_AbstractTypeCode_WithMapOverride(t *testing.T) {
	c := &Convertor{AbstractTypeCodeMap: map[string]cpb.AbstractTypeCode_Value{
		"ANOTHER_TYPE": cpb.AbstractTypeCode_TYPE,
	}}
	if got, want := c.AbstractTypeCode("TYPE"), cpb.AbstractTypeCode_INVALID_UNINITIALIZED; got != want {
		t.Errorf(`AbstractTypeCode("TYPE") got %s, want %s`, got, want)
	}
	if got, want := c.AbstractTypeCode("ANOTHER_TYPE"), cpb.AbstractTypeCode_TYPE; got != want {
		t.Errorf(`AbstractTypeCode("ANOTHER_TYPE") got %s, want %s`, got, want)
	}
}
