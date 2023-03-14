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
	"bytes"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	mshTemplate = "MSH|^~\\&|CERNER|RAL|CARELINK|AKI2|20141128001635||%s|2014112800163507740000|T|2.3|||AL||44|ASCII"
	mshOru      = []byte("MSH|^~\\&|CERNER|RAL|CARELINK|AKI|20160316190000||ORU^R01|201603161900001|T|2.3|||AL||44|ASCII")
	zal         = []byte("ZAL|SNAPSHOT|20091123104044|4219059|4219059|ALLEGY|ACTIVE|||||20081223104044|913445555041^ZZZDOCTOR^Dorian^John^^^^^DRNBR^PRSNL^^^ORGDR^|0")
)

func TestParseMessageV2(t *testing.T) {
	m, err := ParseMessageV2(bytes.Join([][]byte{msh, pid, nk1, obx1, obx2, obx3, zal}, []byte("\r")))
	if err != nil {
		t.Fatalf("ParseMessageV2() failed with %v", err)
	}
	if _, ok := m.(*ADT_A01v2); !ok {
		t.Errorf("ParseMessageV2 returned value of type %T, want *ADT_A01v2", m)
	}
	adt := m.(*ADT_A01v2)
	if adt.MSH() == nil {
		t.Error("ParseMessageType.MSH is <nil>, want not nil")
	}
	if adt.PID() == nil {
		t.Error("ParseMessageType.PID is <nil>, want not nil")
	}
	if got, want := *adt.PID().Religion.Identifier, ST("CATHOLIC"); got != want {
		t.Errorf("PID.Religion got %q, want %q", got, want)
	}
	if adt.AllNK1() == nil {
		t.Error("ParseMessageType.AllNK1 is <nil>, want not nil")
	}
	if got, want := len(adt.AllNK1()), 1; got != want {
		t.Fatalf("len(adt.AllNK1) = %d, want %d", got, want)
	}
	obxs := adt.AllOBX()
	if obxs == nil {
		t.Error("ParseMessageType.AllOBX is <nil>, want not nil")
	}
	if got, want := len(obxs), 3; got != want {
		t.Errorf("len(AllOBX) = %d, want %d", got, want)
	}
	var gotOBXIdentifier []string
	for _, obx := range obxs {
		gotOBXIdentifier = append(gotOBXIdentifier, obx.ObservationIdentifier.Identifier.String())
	}
	want := []string{"PASOVERSEAS", "PASSITECODE", "PERSONUKRES"}
	if diff := cmp.Diff(want, gotOBXIdentifier); diff != "" {
		t.Errorf("AllOBX ObservationIdentifier comparison returned diff (-want +got):\n%s", diff)
	}
	// TODO: Test that ZAL is parsed and accessible somewhere.
}

func TestParseMessageV2_WithSubGroups(t *testing.T) {
	m, err := ParseMessageV2(bytes.Join([][]byte{mshOru, pid, nk1, obx1, obx2, obx3}, []byte("\r")))
	if err != nil {
		t.Fatalf("ParseMessageV2() failed with %v", err)
	}
	if _, ok := m.(*ORU_R01v2); !ok {
		t.Errorf("ParseMessageV2 returned value of type %T, want *ORU_R01v2", m)
	}
	oru := m.(*ORU_R01v2)
	if oru.MSH() == nil {
		t.Error("ParseMessageType.oru.MSH is <nil>, want not nil")
	}
	pids := oru.GroupByPID()
	if pids == nil {
		t.Fatal("ParseMessageType.oru.GroupByPID is <nil>, want not nil")
	}
	if got, want := len(pids), 1; got != want {
		t.Fatalf("len(oru.GroupByPID) = %d, want %d", got, want)
	}
	if got, want := *pids[0].PID().Religion.Identifier, ST("CATHOLIC"); got != want {
		t.Errorf("PID.Religion got %q, want %q", got, want)
	}
	if pids[0].AllNK1() == nil {
		t.Error("pid[0].AllNK1 is <nil>, want not nil")
	}
	obrs := pids[0].GroupByOBR()
	if got, want := len(obrs), 1; got != want {
		t.Errorf("len(GroupByOBR) = %d, want %d", got, want)
	}
	obxs := obrs[0].GroupByOBX()
	if got, want := len(obxs), 3; got != want {
		t.Errorf("len(AllOBX) = %d, want %d", got, want)
	}
	var gotOBXIdentifier []string
	for _, obx := range obxs {
		gotOBXIdentifier = append(gotOBXIdentifier, obx.OBX().ObservationIdentifier.Identifier.String())
	}
	want := []string{"PASOVERSEAS", "PASSITECODE", "PERSONUKRES"}
	if diff := cmp.Diff(want, gotOBXIdentifier); diff != "" {
		t.Errorf("GroupByOBX ObservationIdentifier comparison returned diff (-want +got):\n%s", diff)
	}
}

func TestParseMessageV2_WithChoiceInXMLSchema(t *testing.T) {
	msh := []byte(fmt.Sprintf(mshTemplate, "ORM^O01"))
	m, err := ParseMessageV2(bytes.Join([][]byte{msh, evn, pid, obr}, []byte("\r")))
	if err != nil {
		t.Fatalf("ParseMessageV2() failed with %v", err)
	}
	if _, ok := m.(*ORM_O01v2); !ok {
		t.Errorf("ParseMessageV2 returned value of type %T, want *ORM_O01v2", m)
	}
	orm := m.(*ORM_O01v2)
	if got, want := orm.GroupByORC()[0].OBR().UniversalServiceIdentifier.Text.String(), "UREA AND ELECTROLYTES"; got != want {
		t.Errorf("UniversalServiceIdentifier got %q, want %q", got, want)
	}
}

func TestParseMessageV2_WithUnexpectedSegment(t *testing.T) {
	msh := []byte(fmt.Sprintf(mshTemplate, "ORU^R01"))
	// The EVN segment is unexpected here, but shouldn't cause an issue
	m, err := ParseMessageV2(bytes.Join([][]byte{msh, evn, pid, obr, obx1, obx2}, []byte("\r")))
	if err != nil {
		t.Fatalf("ParseMessageV2() failed with %v", err)
	}
	if _, ok := m.(*ORU_R01v2); !ok {
		t.Errorf("ParseMessageV2 returned value of type %T, want *ORU_R01v2", m)
	}
	if m.(*ORU_R01v2).GroupByPID() == nil {
		t.Error("GroupByPID() is <nil>, want not nil")
	}
}

func TestParseMessageV2_WithError(t *testing.T) {
	tests := []string{
		"ABC^DEF",
		"",
		"^A01",
		"ADT", // ADTs must have trigger events.
		"NST", // TODO: Fix parsing of NST messages: NST messages don't need trigger events.
	}

	for _, tc := range tests {
		t.Run(tc, func(t *testing.T) {
			badMsh := []byte(fmt.Sprintf("MSH|^~\\&|CERNER|RAL|CARELINK|AKI|20160316190000||%s|201603161900001|T|2.3|||AL||44|ASCII", tc))
			m, err := ParseMessageV2(bytes.Join([][]byte{badMsh, pid, nk1, obx1, obx2, obx3}, []byte("\r")))
			if err == nil {
				t.Error("ParseMessageV2 got err=<nil>, want error")
			}
			if m != nil {
				t.Errorf("ParseMessageV2 returned %q, want <nil>", m)
			}
		})
	}
}
