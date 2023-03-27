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

// Package fhircore contains utility functions to deal with core FHIR types.
package fhircore

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
)

// UnixMicro returns the number of microseconds elapsed since January 1, 1970 UTC.
func UnixMicro(t time.Time) int64 {
	return t.UnixNano() / int64(time.Microsecond/time.Nanosecond)
}

func Id(i string) *pb.Id {
	return &pb.Id{
		Value: i,
	}
}

func String(s string) *pb.String {
	return &pb.String{Value: s}
}

func Integer(i int32) *pb.Integer {
	return &pb.Integer{Value: i}
}

func Boolean(b bool) *pb.Boolean {
	return &pb.Boolean{Value: b}
}

func Uri(u string) *pb.Uri {
	return &pb.Uri{Value: u}
}

func Code(c string) *pb.Code {
	return &pb.Code{Value: c}
}

func Date(t time.Time, p pb.Date_Precision) *pb.Date {
	return &pb.Date{
		Precision: p,
		ValueUs:   UnixMicro(t),
	}
}

func Instant(t time.Time, z string) *pb.Instant {
	return &pb.Instant{
		Timezone: z,
		ValueUs:  UnixMicro(t),
	}
}

// DateTime creates a FHIR DateTime from the provided time, using z as the timezone.
// If precision p is set to seconds, the time is truncated.
func DateTime(t time.Time, z string, p pb.DateTime_Precision) *pb.DateTime {
	if p == pb.DateTime_SECOND {
		t = t.Truncate(time.Second)
	}
	return &pb.DateTime{
		Timezone:  z,
		ValueUs:   UnixMicro(t),
		Precision: p,
	}
}

func Time(t time.Time, p pb.Time_Precision) *pb.Time {
	return &pb.Time{
		Precision: p,
		ValueUs:   UnixMicro(t),
	}
}

// fromUnixMicro creates a time.Time object from a Unix timestamp in microseconds. The returned time will be in UTC.
func fromUnixMicro(usec int64) time.Time {
	return fromUnixMicroDate(usec).UTC()
}

// DateTimeToTime returns time.Time from the FHIR proto, assuming that the time unit is microseconds,
// which is a standard for streams-internal interfaces.
func DateTimeToTime(t *pb.DateTime) time.Time {
	return fromUnixMicro(t.GetValueUs())
}

// fromUnixMicroDate creates a time.Time object from a Unix timestamp in microseconds. This function should be used for
// Date instances where we have no Time component.
func fromUnixMicroDate(usec int64) time.Time {
	// Logic adapted from time.Unix but for microseconds.
	sec := usec / 1e6
	usec -= sec * 1e6
	if usec < 0 { // If usec was negative to begin with
		usec += 1e6
		sec--
	}
	nsec := usec * 1e3

	return time.Unix(sec, nsec)
}

func Coding(code string, system string, display string) *pb.Coding {
	return &pb.Coding{
		Code:    Code(code),
		System:  Uri(system),
		Display: String(display),
	}
}

func Decimal(s string) (*pb.Decimal, error) {
	if strings.ContainsAny(s, "eE") {
		return nil, fmt.Errorf("Can't convert %v to Decimal. Exponents are not supported", s)
	}
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, err
	}
	return &pb.Decimal{
		Value: s,
	}, nil
}
