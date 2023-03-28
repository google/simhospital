// Copyright 2023 Google LLC
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

// Package hl7tofhirutils contains standalone functions for converting HL7v2
// fields to FHIR.
package hl7tofhirutils

import (
	"time"

	pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
	fhir "github.com/google/simhospital/pkg/fhircore"
)

// IsNewerMessage checks if the provided extensions contains a DateTime extension with the provided
// URL, and if it does, whether the provided hl7MsgTime is more recent.
// Sometimes HL7 messages should only be taken into account if they are newer than the information
// that we already have.
func IsNewerMessage(extensions []*pb.Extension, dateExtURL string, hl7MsgTime time.Time) bool {
	ext := fhir.FindFirstExtension(extensions, dateExtURL)
	if ext.GetValue().GetDateTime() == nil {
		return true
	}
	return !hl7MsgTime.Before(fhir.DateTimeToTime(ext.GetValue().GetDateTime()))
}
