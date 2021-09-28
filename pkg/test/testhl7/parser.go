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

package testhl7

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/google/simhospital/pkg/hl7"
)

// Parse parses the given message.
func Parse(t *testing.T, message string) *hl7.Message {
	t.Helper()
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%s, %v) failed with %v", message, mo, err)
	}
	if m == nil {
		t.Fatalf("ParseMessageWithOptions(%s, %v) got nil message, want non nil", message, mo)
	}
	return m
}

// Fields computes the given function that returns strings for each message.
func Fields(t *testing.T, messages []string, f func(*testing.T, string) string) []string {
	t.Helper()
	var result []string
	for _, m := range messages {
		result = append(result, f(t, m))
	}
	return result
}

// TimeFields computes the given function that returns times for each message.
func TimeFields(t *testing.T, messages []string, f func(*testing.T, string) *time.Time) []*time.Time {
	t.Helper()
	var result []*time.Time
	for _, m := range messages {
		result = append(result, f(t, m))
	}
	return result
}

// MSH returns the message's MSH segment.
func MSH(t *testing.T, message string) *hl7.MSH {
	t.Helper()
	m := Parse(t, message)

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	return msh
}

// OBX returns the message's OBX segment.
func OBX(t *testing.T, message string) *hl7.OBX {
	t.Helper()
	m := Parse(t, message)

	obx, err := m.OBX()
	if err != nil {
		t.Fatalf("OBX() failed with %v", err)
	}
	if obx == nil {
		t.Fatal("OBX() got nil OBX segment, want non nil")
	}
	return obx
}

// OBR returns the message's OBR segment.
func OBR(t *testing.T, message string) *hl7.OBR {
	t.Helper()
	m := Parse(t, message)

	obx, err := m.OBR()
	if err != nil {
		t.Fatalf("OBR() failed with %v", err)
	}
	if obx == nil {
		t.Fatal("OBR() got nil OBR segment, want non nil")
	}
	return obx
}

// MSA returns the message's MSA segment.
func MSA(t *testing.T, message string) *hl7.MSA {
	t.Helper()
	m := Parse(t, message)

	msa, err := m.MSA()
	if err != nil {
		t.Fatalf("MSA() failed with %v", err)
	}
	if msa == nil {
		t.Fatal("MSA() got nil MSA segment, want non nil")
	}
	return msa
}

// EVN returns the message's EVN segment.
func EVN(t *testing.T, message string) *hl7.EVN {
	t.Helper()
	m := Parse(t, message)

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	return evn
}

// PV1 returns the message's PV1 segment.
func PV1(t *testing.T, message string) *hl7.PV1 {
	t.Helper()
	m := Parse(t, message)

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Fatal("PV1() got nil PV1 segment, want non nil")
	}
	return pv1
}

// PID returns the message's PID segment.
func PID(t *testing.T, message string) *hl7.PID {
	t.Helper()
	m := Parse(t, message)

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Fatal("PID() got nil PID segment, want non nil")
	}
	return pid
}

// PV2 returns the message's PV2 segment.
func PV2(t *testing.T, message string) *hl7.PV2 {
	t.Helper()
	m := Parse(t, message)

	pv2, err := m.PV2()
	if err != nil {
		t.Fatalf("PV2() failed with %v", err)
	}
	if pv2 == nil {
		t.Fatal("PV2() got nil PV2 segment, want non nil")
	}
	return pv2
}

// PV2OrNil returns the message's PV2 segment, or nil if there's no PV2 segment.
func PV2OrNil(t *testing.T, message string) *hl7.PV2 {
	t.Helper()
	m := Parse(t, message)

	pv2, err := m.PV2()
	if err != nil {
		t.Fatalf("PV2() failed with %v", err)
	}
	return pv2
}

// TXA returns the message's TXA segment.
func TXA(t *testing.T, message string) *hl7.TXA {
	t.Helper()
	m := Parse(t, message)

	txa, err := m.TXA()
	if err != nil {
		t.Fatalf("TXA() failed with %v", err)
	}
	return txa
}

// AllDG1 returns all DG1 segments.
func AllDG1(t *testing.T, message string) []*hl7.DG1 {
	t.Helper()
	m := Parse(t, message)

	dg1, err := m.AllDG1()
	if err != nil {
		t.Fatalf("AllDG1() failed with %v", err)
	}
	return dg1
}

// AllPR1 returns all PR1 segments.
func AllPR1(t *testing.T, message string) []*hl7.PR1 {
	t.Helper()
	m := Parse(t, message)

	pr1, err := m.AllPR1()
	if err != nil {
		t.Fatalf("AllPR1() failed with %v", err)
	}
	return pr1
}

// AllAL1 returns all AL1 segments.
func AllAL1(t *testing.T, message string) []*hl7.AL1 {
	t.Helper()
	m := Parse(t, message)

	al1, err := m.AllAL1()
	if err != nil {
		t.Fatalf("AllAL1() failed with %v", err)
	}
	return al1
}

// AllOBX returns all OBX segments.
func AllOBX(t *testing.T, message string) []*hl7.OBX {
	t.Helper()
	m := Parse(t, message)

	obx, err := m.AllOBX()
	if err != nil {
		t.Fatalf("AllOBX() failed with %v", err)
	}
	return obx
}

// ORC returns the message's ORC segment.
func ORC(t *testing.T, message string) *hl7.ORC {
	t.Helper()
	m := Parse(t, message)

	orc, err := m.ORC()
	if err != nil {
		t.Fatalf("ORC() failed with %v", err)
	}
	if orc == nil {
		t.Fatal("ORC() got nil ORC segment, want non nil")
	}
	return orc
}

// MRG returns the message's MRG segment.
func MRG(t *testing.T, message string) *hl7.MRG {
	t.Helper()
	m := Parse(t, message)

	mrg, err := m.MRG()
	if err != nil {
		t.Fatalf("MRG() failed with %v", err)
	}
	if mrg == nil {
		t.Fatal("MRG() got nil ORC segment, want non nil")
	}
	return mrg
}

// MessageDateTime returns the message's date time.
func MessageDateTime(t *testing.T, message string) time.Time {
	t.Helper()
	msh := MSH(t, message)
	return msh.DateTimeOfMessage.Time.UTC()
}

// PlacerNumber returns the PlacerOrderNumber.
func PlacerNumber(t *testing.T, message string) string {
	t.Helper()
	orc := ORC(t, message)
	return orc.PlacerOrderNumber.EntityIdentifier.String()
}

// DeathIndicator returns the PatientDeathIndicator.
func DeathIndicator(t *testing.T, message string) string {
	t.Helper()
	pid := PID(t, message)
	return pid.PatientDeathIndicator.String()
}

// DeathTime returns the PatientDeathIndicator.
func DeathTime(t *testing.T, message string) *time.Time {
	t.Helper()
	pid := PID(t, message)
	if pid.PatientDeathDateAndTime == nil {
		return nil
	}
	return &pid.PatientDeathDateAndTime.Time
}

// OrderStatus returns OrderStatus from the ORC segment.
func OrderStatus(t *testing.T, message string) string {
	t.Helper()
	orc := ORC(t, message)
	return orc.OrderStatus.String()
}

// OBXSetID returns the OBX's SetID.
func OBXSetID(t *testing.T, obx *hl7.OBX) string {
	t.Helper()
	return strconv.Itoa(int(obx.SetIDOBX.Value))
}

// ValueType returns the OBX's observation value type.
func ValueType(t *testing.T, obx *hl7.OBX) string {
	t.Helper()
	return obx.ValueType.String()
}

// OBXResultStatus returns the observation result status.
func OBXResultStatus(t *testing.T, obx *hl7.OBX) string {
	t.Helper()
	return obx.ObservationResultStatus.String()
}

// OBXFieldsFromOBXs applies the function f to all of the given OBX segments.
func OBXFieldsFromOBXs(t *testing.T, obxs []*hl7.OBX, f func(*testing.T, *hl7.OBX) string) []string {
	var fromOBX []string
	for _, o := range obxs {
		fromOBX = append(fromOBX, f(t, o))
	}
	return fromOBX
}

// OBXFields applies the function f to all of the OBX segments in the messages.
func OBXFields(t *testing.T, messages []string, f func(*testing.T, *hl7.OBX) string) [][]string {
	t.Helper()
	var result [][]string
	for _, m := range messages {
		obxs := AllOBX(t, m)
		fromOBX := OBXFieldsFromOBXs(t, obxs, f)
		result = append(result, fromOBX)
	}
	return result
}

// MessageControlIDFromMSH returns the message control ID from the MSH segment.
func MessageControlIDFromMSH(t *testing.T, message string) string {
	t.Helper()
	msh := MSH(t, message)
	return msh.MessageControlID.String()
}

// MessageControlIDFromMSA returns the message control ID from the MSA segment.
func MessageControlIDFromMSA(t *testing.T, message string) string {
	t.Helper()
	msa := MSA(t, message)
	return msa.MessageControlID.String()
}

// PointOfCare returns the point of care from the PV1.AssignedPatientLocation.
func PointOfCare(t *testing.T, message string) string {
	t.Helper()
	pv1 := PV1(t, message)
	return pv1.AssignedPatientLocation.PointOfCare.String()
}

// AccountStatus returns the PV1.AccountStatus.
func AccountStatus(t *testing.T, message string) string {
	t.Helper()
	pv1 := PV1(t, message)
	return pv1.AccountStatus.String()
}

// PatientClass returns the PV1.PatientClass.
func PatientClass(t *testing.T, message string) string {
	t.Helper()
	pv1 := PV1(t, message)
	return pv1.PatientClass.String()
}

// VisitNumber returns the VisitNumber ID.
func VisitNumber(t *testing.T, message string) string {
	t.Helper()
	pv1 := PV1(t, message)
	return pv1.VisitNumber.IDNumber.String()
}

// EventDateTime returns the event's date time.
func EventDateTime(t *testing.T, message string) time.Time {
	t.Helper()
	evn := EVN(t, message)
	return evn.RecordedDateTime.Time.UTC()
}

// FirstName returns the patient first name from the message.
func FirstName(t *testing.T, message string) string {
	t.Helper()
	pid := PID(t, message)
	return pid.PatientName[0].GivenName.String()
}

// MRN returns the patient identifier.
func MRN(t *testing.T, message string) string {
	t.Helper()
	pid := PID(t, message)
	return pid.PatientID.IDNumber.String()
}

// PriorPatientIdentifierList returns the PriorPatientIdentifierList from the MRG segment.
func PriorPatientIdentifierList(t *testing.T, message string) []string {
	t.Helper()
	mrg := MRG(t, message)
	var ids []string
	for _, id := range mrg.PriorPatientIdentifierList {
		ids = append(ids, id.IDNumber.String())
	}
	return ids
}

// SendingApplication returns the message's sending application.
func SendingApplication(t *testing.T, message string) string {
	t.Helper()
	return MSH(t, message).SendingApplication.String()
}

// SendingFacility returns the message's sending facility.
func SendingFacility(t *testing.T, message string) string {
	t.Helper()
	return MSH(t, message).SendingFacility.String()
}

// MessageType returns this message's message type.
func MessageType(t *testing.T, message string) string {
	t.Helper()
	msh := MSH(t, message)
	return fmt.Sprintf("%s^%s", msh.MessageType.MessageCode.String(), msh.MessageType.TriggerEvent.String())
}

// ObservationValue returns this message's OBX observation value.
func ObservationValue(t *testing.T, message string) string {
	t.Helper()
	obx := OBX(t, message)
	return string(obx.ObservationValue[0])
}
