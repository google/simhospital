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

package hl7tofhirutils

import (
	"testing"
	"time"

	pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
	fhir "github.com/google/simhospital/pkg/fhircore"
)

const dateExt = "DateExtURL"

func TestIsNewerMessage(t *testing.T) {
	tcs := []struct {
		name     string
		inExt    []*pb.Extension
		inHL7Msg time.Time
		want     bool
	}{{
		name:     "Old message",
		inHL7Msg: time.Date(2017, 5, 25, 16, 18, 59, 0, time.UTC),
		inExt:    dtExtension(dateExt, time.Date(2017, 5, 25, 16, 19, 59, 0, time.UTC)),
		want:     false,
	}, {
		name:     "Newer message",
		inHL7Msg: time.Date(2018, 5, 25, 16, 18, 59, 0, time.UTC),
		inExt:    dtExtension(dateExt, time.Date(2017, 5, 25, 16, 19, 59, 0, time.UTC)),
		want:     true,
	}, {
		name:     "Equal time is treated as newer",
		inHL7Msg: time.Date(2017, 5, 25, 16, 19, 59, 0, time.UTC),
		inExt:    dtExtension(dateExt, time.Date(2017, 5, 25, 16, 19, 59, 0, time.UTC)),
		want:     true,
	}, {
		name:     "No matching extension is treated as newer",
		inHL7Msg: time.Date(2017, 5, 25, 16, 19, 59, 0, time.UTC),
		inExt:    dtExtension("some other ext", time.Date(2017, 5, 25, 16, 19, 59, 0, time.UTC)),
		want:     true,
	}, {
		name:     "No date time extension is treated as newer",
		inHL7Msg: time.Date(2017, 5, 25, 16, 19, 59, 0, time.UTC),
		inExt:    []*pb.Extension{fhir.StringExtension(dateExt, "a string")},
		want:     true,
	}}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := IsNewerMessage(tc.inExt, dateExt, tc.inHL7Msg)
			if got != tc.want {
				t.Errorf("IsNewerMessage() got %v, want %v", got, tc.want)
			}
		})
	}
}

func dtExtension(extURL string, t time.Time) []*pb.Extension {
	return []*pb.Extension{
		fhir.DateTimeExtension(extURL, t, pb.DateTime_SECOND),
	}
}
