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

package fhircore

import (
	"testing"
	"time"

	pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
)

func TestUnixMicro(t *testing.T) {
	want := int64(12345000678)
	tt := time.Unix(12345, 678000)

	if got, want := UnixMicro(tt), want; got != want {
		t.Errorf("UnixMicro(%v) got %v, want %v", tt, got, want)
	}
}

func TestDateTime_TruncateMs(t *testing.T) {
	input := time.Date(2018, 10, 23, 01, 19, 35, 800000000, time.UTC)
	want := time.Date(2018, 10, 23, 01, 19, 35, 0, time.UTC)
	dt := DateTime(input, "", pb.DateTime_SECOND)
	if got, want := dt.GetValueUs(), UnixMicro(want); got != want {
		t.Errorf("DateTime.GetvalueUs() got %v, want %v", got, want)
	}
}

func TestDateTime_DoNotTruncateMs(t *testing.T) {
	input := time.Date(2018, 10, 23, 01, 19, 40, 800000000, time.UTC)
	dt := DateTime(input, "", pb.DateTime_MILLISECOND)
	if got, want := dt.GetValueUs(), UnixMicro(input); got != want {
		t.Errorf("DateTime.GetvalueUs() got %v, want %v", got, want)
	}
}
