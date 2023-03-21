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

package keyed

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/google/simhospital/pkg/hl7"
)

const (
	mshTemplate = "MSH|^~\\&|CERNER|RAL|CARELINK|AKI2|20141128001635||%s|2014112800163507740000|T|2.3|||AL||44|ASCII"
	pidTemplate = "PID|1|%s^^^RAL MRN^MRN^|843124^^^RAL MRN^MRN^CD:5294405~1231231235^^^NHSNBR^NHSNMBR||ZZZTEST^%s^^^MR^^CURRENT||19000524|1|ZZZTEST^PAUL^^^MR^^PREVIOUS||100 The Street^Any Square^LONDON^^ZZ99 1BA^GBR^HOME^^||0205551234^HOME^CD:4072430~0205551234^CD:4580206||||CATHOLIC|3393034^^^RAL Encounter Num^FINNBR^||||C|||0|||||"
	pv1Template = "PV1|1|INPATIENT|RAL 10 EAST^Bay01^%s^RAL RF^^BED^RFH|28b||^^^^^^|C3335493^Harber^Mark^^^^^^NHSCSLTNBR^PRSNL^^^NONGP^~895000428014^Harber^Mark^^^^^^DRNBR^PRSNL^^^ORGDR^|C3335493^Harber^Mark^^^^^^NHSCSLTNBR^PRSNL^^^NONGP^~895000428014^Harber^Mark^^^^^^DRNBR^PRSNL^^^ORGDR^||361||||19||||INPATIENT|6153479^^^^VISITID||||||||||||||||||||RAL RF||ACTIVE|||20141127180000"
)

var (
	msh = []byte("MSH|^~\\&|CERNER|RAL|CARELINK|AKI2|20141128001635||ADT^A01|2014112800163507740000|T|2.3|||AL||44|ASCII")
	pid = []byte("PID|1|843124^^^RAL MRN^MRN^|843124^^^RAL MRN^MRN^CD:5294405~1231231235^^^NHSNBR^NHSNMBR||ZZZTEST^PAUL^^^MR^^CURRENT||19000524|1|ZZZTEST^PAUL^^^MR^^PREVIOUS||100 The Street^Any Square^LONDON^^ZZ99 1BA^GBR^HOME^^||0205551234^HOME^CD:4072430~0205551234^CD:4580206||||CATHOLIC|3393034^^^RAL Encounter Num^FINNBR^||||C|||0|||||")
	pv1 = []byte("PV1|1|INPATIENT|RAL 10 EAST^Bay01^Bed10^RAL RF^^BED^RFH|28b||^^^^^^|C3335493^Harber^Mark^^^^^^NHSCSLTNBR^PRSNL^^^NONGP^~895000428014^Harber^Mark^^^^^^DRNBR^PRSNL^^^ORGDR^|C3335493^Harber^Mark^^^^^^NHSCSLTNBR^PRSNL^^^NONGP^~895000428014^Harber^Mark^^^^^^DRNBR^PRSNL^^^ORGDR^||361||||19||||INPATIENT|6153479^^^^VISITID||||||||||||||||||||RAL RF||ACTIVE|||20141127180000")
)

func TestMain(m *testing.M) {
	hl7.TimezoneAndLocation("Europe/London")
	retCode := m.Run()
	os.Exit(retCode)
}

func TestSafeGetPID(t *testing.T) {
	mt, err := hl7.ParseMessage(bytes.Join([][]byte{msh, pid}, []byte("\r")))
	if err != nil {
		t.Errorf("ParseMessage() failed with %v", err)
	}

	m := Message{Message: mt}

	pid := m.SafeGetPID()
	if pid == nil {
		t.Error("PID() = <nil>, want not <nil>")
	}
	if pid.PatientName[0].GivenName.String() != "PAUL" {
		t.Errorf("pid.PatientName[0].GivenName got %v, want %v", pid.PatientName[0].GivenName.String(), "PAUL")
	}
	if pid.PatientName[0].FamilyName.Surname.String() != "ZZZTEST" {
		t.Errorf("pid.PatientName[0].FamilyName.Surname got %v, want %v", pid.PatientName[0].FamilyName.Surname.String(), "ZZZTEST")
	}
}

func TestGetPID_MultiplePID(t *testing.T) {
	m, err := hl7.ParseMessage(bytes.Join([][]byte{
		[]byte(fmt.Sprintf(mshTemplate, "ADT^A17")),
		[]byte(fmt.Sprintf(pidTemplate, "123", "name1")),
		[]byte(fmt.Sprintf(pidTemplate, "456", "name2"))}, []byte("\r")))
	if err != nil {
		t.Errorf("ParseMessage() failed with %v", err)
	}
	mt, err := m.ParseMessageType()
	if err != nil {
		t.Fatalf("ParseMessageType() failed with %v", err)
	}

	km := Message{
		ParsedMessage: mt,
	}

	pid, err := km.GetPIDByName("PID1")
	if err != nil {
		t.Fatalf("GetPIDByName() failed with %v", err)
	}
	if pid.PatientName[0].GivenName.String() != "name1" {
		t.Errorf("pid.PatientName[0].GivenName got %v, want %v", pid.PatientName[0].GivenName.String(), "name1")
	}
	if pid.PatientID.IDNumber.String() != "123" {
		t.Errorf("pid.PatientID.IDNumber got %v, want %v", pid.PatientID.IDNumber.String(), "123")
	}
	pid, err = km.GetPIDByName("PID2")
	if err != nil {
		t.Fatalf("GetPIDByName() failed with %v", err)
	}
	if got, want := pid.PatientName[0].GivenName.String(), "name2"; got != want {
		t.Errorf("pid.PatientName[0].GivenName got %v, want %v", got, want)
	}
	if pid.PatientID == nil {
		t.Fatal("pid.PatientID = <nil>, want not <nil>")
	}
	if got, want := pid.PatientID.IDNumber.String(), "456"; got != want {
		t.Errorf("pid.PatientID.IDNumber got %v, want %v", got, want)
	}
}

func TestSafeGetPV1(t *testing.T) {
	mt, err := hl7.ParseMessage(bytes.Join([][]byte{msh, pid, pv1}, []byte("\r")))
	if err != nil {
		t.Errorf("ParseMessage() failed with %v", err)
	}

	m := Message{Message: mt}

	pv1 := m.SafeGetPV1()
	if pv1 == nil {
		t.Fatal("PV1() = <nil>, want not <nil>")
	}
	if pv1.VisitNumber == nil {
		t.Fatal("pv1.VisitNumber = <nil>, want not <nil>")
	}
	if got, want := pv1.VisitNumber.IDNumber.String(), "6153479"; got != want {
		t.Errorf("pv1.VisitNumber.IDNumber got %v, want %v", got, want)
	}
}

func TestGetPV1_MultiplePV1(t *testing.T) {
	m, err := hl7.ParseMessage(bytes.Join([][]byte{
		[]byte(fmt.Sprintf(mshTemplate, "ADT^A17")),
		[]byte(fmt.Sprintf(pv1Template, "bed1")),
		[]byte(fmt.Sprintf(pv1Template, "bed2"))}, []byte("\r")))
	if err != nil {
		t.Errorf("ParseMessage() failed with %v", err)
	}
	mt, err := m.ParseMessageType()
	if err != nil {
		t.Fatalf("ParseMessageType() failed with %v", err)
	}

	km := Message{
		ParsedMessage: mt,
	}

	pv1, err := km.GetPV1ByName("PV11")
	if err != nil {
		t.Errorf("GetPV1ByName(%q) failed with %v", "PV11", err)
	}
	if pv1.AssignedPatientLocation.Bed.String() != "bed1" {
		t.Errorf("pv1.AssignedPatientLocation.Bed got %v, want %v", pv1.AssignedPatientLocation.Bed.String(), "bed1")
	}
	pv1, err = km.GetPV1ByName("PV12")
	if err != nil {
		t.Errorf("GetPV1ByName(%q) failed with %v", "PV12", err)
	}
	if pv1.AssignedPatientLocation == nil {
		t.Fatal("pv1.AssignedPatientLocation = <nil>, want not <nil>")
	}
	if got, want := pv1.AssignedPatientLocation.Bed.String(), "bed2"; got != want {
		t.Errorf("pv1.AssignedPatientLocation.Bed got %v, want %v", got, want)
	}
}

func TestSafeGetMSH(t *testing.T) {
	mt, err := hl7.ParseMessage(bytes.Join([][]byte{msh}, []byte("\r")))
	if err != nil {
		t.Errorf("ParseMessage() failed with %v", err)
	}

	m := Message{Message: mt}

	msh := m.SafeGetMSH()
	if msh == nil {
		t.Fatal("MSH() = <nil>, want not <nil>")
	}
	if msh.MessageControlID == nil {
		t.Fatal("msh.MessageControlID = <nil>, want not <nil>")
	}
	if got, want := msh.MessageControlID.String(), "2014112800163507740000"; got != want {
		t.Errorf("msh.MessageControlID got %v, want %v", got, want)
	}
}
