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

package hl7

import (
	"reflect"
)

// ZCM represents a generic CustoM HL7v2 segment.
// The fields in this segment are examples of what this segment can contain.
// These fields could have different goals, e.g. (a) contain certain fields that
// are located in different places in different clients (and thus allowing downstreams services to
// process them the same way), and (b) represent concepts that are useful for processing in the
// system, but are not part of the HL7v2 standard.
type ZCM struct {
	SetIDZCM *SI `hl7:"true,Set ID - ZCM"` // ZCM-1
	// For order-related messages, the discipline (e.g., RADIOLOGY) of the order, if known.
	OrderDiscipline *ST `hl7:"false,Order Discipline"` // ZCM-2
	// The original value of PD1-3-Patient Primary Facility/XON-3-ID Number.
	// The ID Number is defined as an NM field in HL7 but in reality it's not always a number, which
	// can make our parsing fail. We put it here instead.
	PatientPrimaryFacilityIDNumber *ST `hl7:"false,Patient Primary Facility ID Number"` // ZCM-3
	// A keyword with the custom message type for this message, if any. Custom message types
	// trigger special processing logic downstream.
	CustomMessageType *ST `hl7:"false,Custom Message Type"` // ZCM-4
}

// SegmentName returns the ZCM segment name.
func (s *ZCM) SegmentName() string {
	return "ZCM"
}

// AllZCM returns a slice containing all ZCM segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllZCM() ([]*ZCM, error) {
	pss, err := m.ParseAll("ZCM")
	return pss.([]*ZCM), err
}

// ZCM returns the first ZCM segment within the message, or nil if there isn't one.
func (m *Message) ZCM() (*ZCM, error) {
	ps, err := m.Parse("ZCM")
	pst, ok := ps.(*ZCM)
	if ok {
		return pst, err
	}
	return nil, err
}

func init() {
	Types["ZCM"] = reflect.TypeOf(ZCM{})
}
