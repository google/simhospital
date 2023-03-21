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

// Package keyed contains functions to handle HL7 messages in realistic settings, for instance:
// - They are loaded from a database and have a key
// - They contain custom Z-Segments
// - They have been somewhat validated - and thus we can ignore errors in parsing.
package keyed

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/google/simhospital/pkg/hl7"
)

var errSegmentNotFound = errors.New("Segment not found, multiple entries for this segment might exist with" +
	" a different name")

const (
	MessageTypeADT         = "adt"
	MessageTypeORU         = "oru"
	MessageTypeORR         = "orr"
	MessageTypeORM         = "orm"
	MessageTriggerEventA03 = "A03"
	MessageTriggerEventA11 = "A11"
	MessageTriggerEventA13 = "A13"
	MessageTriggerEventA17 = "A17"
	MessageTriggerEventA23 = "A23"
	MessageTriggerEventA34 = "A34"
	MessageTriggerEventA40 = "A40"
)

// Message is a wrapper of an hl7.Message that contains the Message key from the database,
// mostly for logging.
// After parsing, this is the main type of message that could be passed around.
type Message struct {
	// Sanitized hl7.Message, see MessageSanitizer.
	// Sanitized messages are parsable, so errors to retrieve segments can be ignored.
	*hl7.Message
	// The message's key from the database.
	Key []byte

	ParsedMessage any

	// Trigger event for this message, A01, A02, R01 etc
	TriggerEvent string

	// True if the message is reliable.
	Reliable bool

	// The SendingFacility (MSH.SendingFacility.NamespaceID) so that it can be accessed easily.
	SendingFacility string

	// The message type for this message: adt, oru, etc.
	// This is lower case for compatibility with current metrics.
	MessageType string

	MessageDateTime time.Time

	SendingApplication string
}

// GetPIDByName gets the PID segment from this message with the provided name. This function should be used when
// a PID segment does not have the default name of "PID".
func (k Message) GetPIDByName(s string) (*hl7.PID, error) {
	v := reflect.ValueOf(k.ParsedMessage)
	f := v.Elem().FieldByName(s)
	if !f.IsValid() || f.IsNil() {
		return nil, errSegmentNotFound
	}
	return f.Interface().(*hl7.PID), nil
}

// GetPV1ByName gets the PV1 segment from this message with the provided name. This function should be used when
// a PV1 segment does not have the default name of "PV1".
func (k Message) GetPV1ByName(s string) (*hl7.PV1, error) {
	v := reflect.ValueOf(k.ParsedMessage)
	f := v.Elem().FieldByName(s)
	if !f.IsValid() || f.IsNil() {
		return nil, errSegmentNotFound
	}
	return f.Interface().(*hl7.PV1), nil
}

func (k Message) GetSerialisedKeyedMessage() string {
	v := reflect.ValueOf(k.ParsedMessage)
	return fmt.Sprintf("%+v", v)
}

// SafeGetMSH returns the MSH segment. Keyed Messages have a present and parsable MSH segment.
func (k Message) SafeGetMSH() *hl7.MSH {
	msh, _ := k.MSH()
	return msh
}

// SafeGetPID returns the PID segment. Keyed Messages have a present and parsable PID segment.
func (k Message) SafeGetPID() *hl7.PID {
	pid, _ := k.PID()
	return pid
}

// SafeGetZCM returns the ZCM segment. Keyed Messages have a present and parsable ZCM segment.
func (k Message) SafeGetZCM() *hl7.ZCM {
	zdm, _ := k.ZCM()
	return zdm
}

// SafeGetPV1 returns the PV1 segment. Keyed Messages have a parsable PV1 segment, so errors
// can be ignored. The returned segment might be nil.
func (k Message) SafeGetPV1() *hl7.PV1 {
	pv1, _ := k.PV1()
	return pv1
}

// SafeGetOBR returns the OBR segment. Keyed Messages have a present and parsable OBR segment.
// The returned segment might be nil.
func (k Message) SafeGetOBR() *hl7.OBR {
	obr, _ := k.OBR()
	return obr
}

// SafeGetORC returns the ORC segment. Keyed Messages have a present and parsable ORC segment.
// The returned segment might be nil.
func (k Message) SafeGetORC() *hl7.ORC {
	orc, _ := k.ORC()
	return orc
}

// SafeGetOBX returns the OBX segment. Keyed Messages have a present and parsable OBX segment.
// The returned segment might be nil.
func (k Message) SafeGetOBX() *hl7.OBX {
	obx, _ := k.OBX()
	return obx
}
