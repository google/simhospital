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

package hl7

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var (
	msh  = []byte("MSH|^~\\&|CERNER|RAL|CARELINK|AKI2|20141128001635||ADT^A01|2014112800163507740000|T|2.3|||AL||44|ASCII")
	evn  = []byte("EVN|R01|20170329021843|||216865551019^Osman^Arthur^^^Dr^^^DRNBR^PRSNL^^^ORGDR")
	pid  = []byte("PID|1|843124^^^RAL MRN^MRN^|843124^^^RAL MRN^MRN^CD:5294405~1231231235^^^NHSNBR^NHSNMBR||ZZZTEST^PAUL^^^MR^^CURRENT||19000524|1|ZZZTEST^PAUL^^^MR^^PREVIOUS||100 The Street^Any Square^LONDON^^ZZ99 1BA^GBR^HOME^^||0205551234^HOME^CD:4072430~0205551234^CD:4580206||||CATHOLIC|3393034^^^RAL Encounter Num^FINNBR^||||C|||0|||||")
	pv1  = []byte("PV1|1|INPATIENT|RAL 10 EAST^Bay01^Bed10^RAL RF^^BED^RFH|28b||^^^^^^|C3335493^Harber^Mark^^^^^^NHSCSLTNBR^PRSNL^^^NONGP^~895000428014^Harber^Mark^^^^^^DRNBR^PRSNL^^^ORGDR^|C3335493^Harber^Mark^^^^^^NHSCSLTNBR^PRSNL^^^NONGP^~895000428014^Harber^Mark^^^^^^DRNBR^PRSNL^^^ORGDR^||361||||19||||INPATIENT|6153479^^^^VISITID||||||||||||||||||||RAL RF||ACTIVE|||20141127180000")
	nk1  = []byte("NK1|1|ZZZTESTWIFE^JANE^^^^^CURRENT|SPOUSE|96 The Street^^LONDON^^ZZ99 1AA^^^^|02011115555||FAMILYMEM|||||||||||||")
	obr  = []byte("OBR|1|5081597|3847829|lpdc-3969^UREA AND ELECTROLYTES||20170329021843|20170329021843")
	obx1 = []byte("OBX|1|CD|PASOVERSEAS||8")
	obx2 = []byte("OBX|2|CD|PASSITECODE||thisite")
	obx3 = []byte("OBX|3|CD|PERSONUKRES||YES")

	// wantPID is the representation of the "pid" variable as a *PID struct.
	wantPID = &PID{
		SetIDPID:  NewSI(1),
		PatientID: &CX{IDNumber: NewST("843124"), AssigningAuthority: &HD{NamespaceID: NewIS("RAL MRN")}, IdentifierTypeCode: NewID("MRN")},
		PatientIdentifierList: []CX{
			{IDNumber: NewST("843124"), AssigningAuthority: &HD{NamespaceID: NewIS("RAL MRN")}, IdentifierTypeCode: NewID("MRN"), AssigningFacility: &HD{NamespaceID: NewIS("CD:5294405")}},
			{IDNumber: NewST("1231231235"), AssigningAuthority: &HD{NamespaceID: NewIS("NHSNBR")}, IdentifierTypeCode: NewID("NHSNMBR")},
		},
		PatientName: []XPN{
			{FamilyName: &FN{Surname: NewST("ZZZTEST")}, GivenName: NewST("PAUL"), PrefixEGDR: NewST("MR"), NameTypeCode: NewID("CURRENT")},
		},
		DateTimeOfBirth:   &TS{Time: time.Date(1900, 5, 24, 0, 0, 0, 0, time.UTC), Precision: TSPrecision(2)},
		AdministrativeSex: NewIS("1"),
		PatientAlias: []XPN{
			{FamilyName: &FN{Surname: NewST("ZZZTEST")}, GivenName: NewST("PAUL"), PrefixEGDR: NewST("MR"), NameTypeCode: NewID("PREVIOUS")},
		},
		Religion: &CE{Identifier: NewST("CATHOLIC")},
		PatientAddress: []XAD{
			{StreetAddress: &SAD{StreetOrMailingAddress: NewST("100 The Street")}, OtherDesignation: NewST("Any Square"), City: NewST("LONDON"), ZipOrPostalCode: NewST("ZZ99 1BA"), Country: NewID("GBR"), AddressType: NewID("HOME")},
		},
		PhoneNumberHome: []XTN{
			{Number: NewST("0205551234"), TelecommunicationUseCode: NewID("HOME"), TelecommunicationEquipmentType: NewID("CD:4072430")},
			{Number: NewST("0205551234"), TelecommunicationUseCode: NewID("CD:4580206")},
		},
		PatientAccountNumber: &CX{IDNumber: NewST("3393034"), AssigningAuthority: &HD{NamespaceID: NewIS("RAL Encounter Num")}, IdentifierTypeCode: NewID("FINNBR")},
		EthnicGroup:          []CWE{{Identifier: NewST("C")}},
		BirthOrder:           NewNM(0),
	}

	segmentTerminatorBytes = []byte(SegmentTerminatorStr)
)

func TestMain(m *testing.M) {
	TimezoneAndLocation("Europe/London")
	retCode := m.Run()
	os.Exit(retCode)
}

func TestParsePIDWithinMessage(t *testing.T) {
	m, err := ParseMessage(bytes.Join([][]byte{msh, pid, nk1, obx1, obx2, obx3}, segmentTerminatorBytes))
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	got, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}

	if diff := cmp.Diff(wantPID, got); diff != "" {
		t.Errorf("PID() diff: (-want, +got):\n%s", diff)
	}
}

func TestParsePD1WithinMessageWhenNonPresent(t *testing.T) {
	m, err := ParseMessage(bytes.Join([][]byte{msh, pid}, segmentTerminatorBytes))
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 != nil {
		t.Errorf("PD1() got %+v, want <nil>", pd1)
	}
}

func TestParseOBXWithinMessageReturnsFirstOfMany(t *testing.T) {
	m, err := ParseMessage(bytes.Join([][]byte{msh, pid, nk1, obx1, obx2, obx3}, segmentTerminatorBytes))
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	obx, err := m.OBX()
	if err != nil {
		t.Fatalf("OBX() failed with %v", err)
	}
	if got, want := *obx.ObservationIdentifier.Identifier, ST("PASOVERSEAS"); got != want {
		t.Errorf("OBX().ObservationIdentifier got %q, want %q", got, want)
	}
}

func TestParseAllOBXWithinMessage(t *testing.T) {
	m, err := ParseMessage(bytes.Join([][]byte{msh, pid, nk1, obx1, obx2, obx3}, segmentTerminatorBytes))
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	obx, err := m.AllOBX()
	if err != nil {
		t.Fatalf("AllOBX() failed with %v", err)
	}
	if got, want := len(obx), 3; got != want {
		t.Fatalf("len(AllOBX) = %d, want %d", got, want)
	}
	tests := []struct {
		i    int
		want string
	}{
		{0, "PASOVERSEAS"},
		{1, "PASSITECODE"},
		{2, "PERSONUKRES"},
	}
	for _, tc := range tests {
		if got, want := *obx[tc.i].ObservationIdentifier.Identifier, ST(tc.want); got != want {
			t.Errorf("OBX()[%d].ObservationIdentifie.Identifier got %q, want %q", tc.i, got, want)
		}
	}
}

func TestParseAllPD1WithinMessageWhenNotPresent(t *testing.T) {
	m, err := ParseMessage(bytes.Join([][]byte{msh, pid, nk1, obx1, obx2, obx3}, segmentTerminatorBytes))
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	pd1, err := m.AllPD1()
	if err != nil {
		t.Fatalf("AllPD1() failed with %v", err)
	}
	if got, want := len(pd1), 0; got != want {
		t.Errorf("len(AllPD1) = %d, want %d", got, want)
	}
}

func TestAll(t *testing.T) {
	m, err := ParseMessage(bytes.Join([][]byte{msh, pid, nk1, obx1, obx2, obx3}, segmentTerminatorBytes))
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	s, err := m.All()
	if err != nil {
		t.Fatalf("All() failed with %v", err)
	}
	if got, want := len(s), 6; got != want {
		t.Fatalf("len(All()) = %d, want %d", got, want)
	}
	if _, ok := s[0].(*MSH); !ok {
		t.Errorf("s[0] is of type %T, want *MSH", s[0])
	}
	if _, ok := s[1].(*PID); !ok {
		t.Errorf("s[1] is of type %T, want *PID", s[1])
	}
	if _, ok := s[2].(*NK1); !ok {
		t.Errorf("s[2] is of type %T, want *NK1", s[2])
	}
	if _, ok := s[3].(*OBX); !ok {
		t.Errorf("s[3] is of type %T, want *OBX", s[3])
	}
	if _, ok := s[4].(*OBX); !ok {
		t.Errorf("s[4] is of type %T, want *OBX", s[4])
	}
	if _, ok := s[5].(*OBX); !ok {
		t.Errorf("s[5] is of type %T, want *OBX", s[5])
	}
}

func BenchmarkAll(b *testing.B) {
	m, _ := ParseMessage(bytes.Join([][]byte{msh, pid, nk1, obx1, obx2, obx3}, segmentTerminatorBytes))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.All()
	}
}

func TestParseMSHHandlesFieldSeparator(t *testing.T) {
	m, err := ParseMessage(msh)
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if got, want := msh.EncodingCharacters.Field, byte('|'); got != want {
		t.Errorf("MSH().EncodingCharacters.Field got %v, want %v", got, want)
	}
	if got, want := msh.EncodingCharacters.Component, byte('^'); got != want {
		t.Errorf("MSH().EncodingCharacters.Component got %v, want %v", got, want)
	}
	if got, want := msh.EncodingCharacters.Subcomponent, byte('&'); got != want {
		t.Errorf("MSH().EncodingCharacters.Subcomponent got %v, want %v", got, want)
	}
}

func TestErrorsContainByteOffsetWithinMessage(t *testing.T) {
	badPid := []byte("PID|1|843124^^^RAL MRN^MRN^|843124^^^RAL MRN^MRN^CD:5294405~1231231235^^^NHSNBR^NHSNMBR||ZZZTEST^PAUL^^^MR^^CURRENT||19000524|1|ZZZTEST^PAUL^^^MR^^PREVIOUS||100 The Street^Any Square^LONDON^^ZZ99 1BA^GBR^HOME^^||0205551234^HOME^CD:4072430~0205551234^CD:4580206||||CATHOLIC|3393034^^^RAL Encounter Num^FINNBR^||||C|||BROKEN|||||")
	m, err := ParseMessage(bytes.Join([][]byte{msh, badPid}, segmentTerminatorBytes))
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	pid, err := m.PID()
	// The valid parts of the message are still returned on error.
	if pid == nil {
		t.Fatal("PID() returned <nil>, want non-nil *PID")
	}
	if got, want := *pid.Religion.Identifier, ST("CATHOLIC"); got != want {
		t.Errorf("PID().Religion got %q, want %q", got, want)
	}

	if err == nil {
		t.Fatal("PID() returned err=<nil>, want error")
	}
	if _, ok := err.(ParseErrors); !ok {
		t.Fatalf("PID() returned error of type %T, want ParseError{}", err)
	}
	pe := err.(ParseErrors)
	if got, want := len(pe), 1; got != want {
		t.Fatalf("len(ParseErrors) = %d, want %d", got, want)
	}
	if got, want := pe[0].Offset, 418; got != want {
		t.Errorf("[%v].Offset got %v, want %v", pe[0], got, want)
	}
}

func TestMultipleParseErrorsReportedInASingleError(t *testing.T) {
	badPid := []byte("PID|1|843124^^^RAL MRN^MRN^|||||||||||||||||||||||BROKEN||||||||BROKEN")
	m, err := ParseMessage(bytes.Join([][]byte{msh, badPid}, segmentTerminatorBytes))
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	pid, err := m.PID()
	if pid == nil {
		t.Fatal("PID() returned <nil>, want non-nil *PID")
	}
	if err == nil {
		t.Fatal("PID() returned err=<nil>, want non-nil error")
	}
	if _, ok := err.(ParseErrors); !ok {
		t.Fatalf("PID() returned error of type %T, want ParseError{}", err)
	}
	pe := err.(ParseErrors)
	if got, want := len(pe), 2; got != want {
		t.Fatalf("len(ParseErrors) = %d, want %d", got, want)
	}
	if got, want := pe[0].Offset, 152; got != want {
		t.Errorf("[%v].Offset got %v, want %v", pe[0], got, want)
	}
	if got, want := pe[1].Offset, 166; got != want {
		t.Errorf("[%v].Offset got %v, want %v", pe[1], got, want)
	}
}

func TestErrorsContainLocations(t *testing.T) {
	badPid := []byte("PID|1|843124^^^RAL MRN^MRN^|||||||||||||||||||||||BROKEN||||||||BROKEN")
	m, err := ParseMessage(bytes.Join([][]byte{msh, badPid}, segmentTerminatorBytes))
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	_, err = m.PID()
	if err == nil {
		t.Fatal("PID() returned err=<nil>, want non-nil error")
	}
	if _, ok := err.(ParseErrors); !ok {
		t.Fatalf("PID() returned error of type %T, want ParseError{}", err)
	}
	pe := err.(ParseErrors)
	if got, want := len(pe), 2; got != want {
		t.Fatalf("len(ParseErrors) = %d, want %d", got, want)
	}
	tests := []struct {
		i    int
		want string
	}{
		{0, "PID-25-Birth Order"},
		{1, "PID-33-Last Update Date/Time"},
	}
	for _, tc := range tests {
		if got, want := pe[tc.i].Location, tc.want; got != want {
			t.Errorf("[%v].Location got %v, want %v", pe[tc.i], got, want)
		}
		if !strings.Contains(err.Error(), tc.want) {
			t.Errorf("err=%q, want to contain %q", err, tc.want)
		}
	}
}

func TestErrorsContainNumberOfErrors(t *testing.T) {
	badPid := []byte("PID|1|843124^^^RAL MRN^MRN^|||||||||||||||||||||||BROKEN||")
	m, err := ParseMessage(bytes.Join([][]byte{msh, badPid}, segmentTerminatorBytes))
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	_, err = m.PID()
	if err == nil {
		t.Fatal("PID() returned err=<nil>, want non-nil error")
	}
	wantPrefix := "errors (1): "
	if !strings.HasPrefix(err.Error(), wantPrefix) {
		t.Errorf("err=%q, want to have prefix %q", err, wantPrefix)
	}
}

func TestAllReportsErrorsForBadSegments(t *testing.T) {
	badSegment := []byte("XXX|1|||")
	m, err := ParseMessage(bytes.Join([][]byte{msh, pid, badSegment}, segmentTerminatorBytes))
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	s, err := m.All()
	if err == nil {
		t.Fatal("All() returned err=<nil>, want non-nil error")
	}
	if got, want := len(s), 2; got != want {
		t.Fatalf("len(All()) = %d, want %d", got, want)
	}
	if _, ok := s[0].(*MSH); !ok {
		t.Errorf("s[0] is of type %T, want *MSH", s[0])
	}
	if _, ok := s[1].(*PID); !ok {
		t.Errorf("s[1] is of type %T, want *PID", s[0])
	}
	if _, ok := err.(ParseErrors); !ok {
		t.Fatalf("All() returned error of type %T, want ParseError{}", err)
	}
	pe := err.(ParseErrors)
	if got, want := len(pe), 1; got != want {
		t.Fatalf("len(ParseErrors) = %d, want %d", got, want)
	}
	if got, want := pe[0].Offset, 425; got != want {
		t.Errorf("[%v].Offset got %v, want %v", pe[0], got, want)
	}
	cause, ok := pe[0].Cause.(*BadSegmentError)
	if !ok {
		t.Fatalf("[%v].Cause is of type %T, want &BadSegmentError{}", pe[0], pe[0].Cause)
	}
	if got, want := cause.Name, "XXX"; got != want {
		t.Errorf("[%v].Name got %q, want %q", cause, got, want)
	}
}

func TestAllHandlesBlankAndShortSegments(t *testing.T) {
	m, err := ParseMessage(bytes.Join([][]byte{msh, segmentTerminatorBytes, []byte("XX|0|")}, segmentTerminatorBytes))
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	s, err := m.All()
	if got, want := len(s), 1; got != want {
		t.Fatalf("len(All()) = %d, want %d", got, want)
	}
	if _, ok := s[0].(*MSH); !ok {
		t.Errorf("s[0] is of type %T, want *MSH", s[0])
	}
	if err == nil {
		t.Fatal("All() returned err=<nil>, want non-nil error")
	}
	if _, ok := err.(ParseErrors); !ok {
		t.Fatalf("All() returned error of type %T, want ParseError{}", err)
	}
	pe := err.(ParseErrors)
	if got, want := len(pe), 1; got != want {
		t.Fatalf("len(ParseErrors) = %d, want %d", got, want)
	}
	cause, ok := pe[0].Cause.(*BadSegmentError)
	if !ok {
		t.Fatalf("[%v].Cause is of type %T, want &BadSegmentError{}", pe[0], pe[0].Cause)
	}
	if got, want := cause.Name, "XX|"; got != want {
		t.Errorf("[%v].Name got %q, want %q", cause, got, want)
	}
}

func TestAllIncludesZSegments(t *testing.T) {
	m, err := ParseMessage(bytes.Join([][]byte{msh, []byte("ZAL|0|")}, segmentTerminatorBytes))
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	s, err := m.All()
	if err != nil {
		t.Fatalf("All() failed with %v", err)
	}
	if got, want := len(s), 2; got != want {
		t.Errorf("len(All()) = %d, want %d", got, want)
	}
}

func TestConversionToUTF8(t *testing.T) {
	mshCases := []struct {
		name string
		msh  string
	}{
		{"Correctly formatted", "MSH|^~\\&|||||||ADT^A01||T|2.3||||||8859/1"},
		{"Trailing spaces", "MSH|^~\\&|||||||ADT^A01||T|2.3||||||8859/1  "},
		{"Leading spaces", "MSH|^~\\&|||||||ADT^A01||T|2.3||||||   8859/1"},
	}
	for _, c := range mshCases {
		t.Run(c.name, func(t *testing.T) {
			// \xe9 is e acute in the iso88591 character set.
			pidISO88591 := []byte("PID|1||||Smith^John^Jos\xe9^^MR^^CURRENT|")
			m, err := ParseMessage(bytes.Join([][]byte{[]byte(c.msh), pidISO88591}, segmentTerminatorBytes))
			if err != nil {
				t.Fatalf("ParseMessage() failed with %v", err)
			}
			pid, err := m.PID()
			if err != nil {
				t.Fatalf("PID() failed with %v", err)
			}
			if len(pid.PatientName) == 0 {
				t.Fatal("PID().PatientName is empty, want non empty")
			}
			if got, want := *pid.PatientName[0].SecondAndFurtherGivenNamesOrInitialsThereof, ST("José"); got != want {
				t.Errorf("PID().PatientName[0].SecondAndFurtherGivenNamesOrInitialsThereof got %q, want %q", got, want)
			}
		})
	}
}

func TestTimezoneAndLocation(t *testing.T) {
	cases := []struct {
		timezone string
		wantErr  bool
	}{
		{timezone: "Europe/Madrid", wantErr: false},
		{timezone: "invalid", wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.timezone, func(t *testing.T) {
			if err := TimezoneAndLocation(tc.timezone); (err != nil) != tc.wantErr {
				t.Fatalf("TimezoneAndLocation(%q) got err=%v; want err? %t", tc.timezone, err, tc.wantErr)
			}
			if !tc.wantErr && Timezone != tc.timezone {
				t.Errorf("Timezone got %q; want %q", Timezone, tc.timezone)
			}
		})
	}
}

func TestParseMessageWithOptions_AllowNullHeader(t *testing.T) {
	allowNullHeaderDisabled := NewParseMessageOptions()
	allowNullHeaderEnabled := NewParseMessageOptions()
	allowNullHeaderEnabled.AllowNullHeader = true

	tests := []struct {
		name     string
		segments [][]byte
		options  *ParseMessageOptions
		wantErr  bool
	}{{
		name:     "AllowNullHeader disabled",
		segments: [][]byte{pid, nk1, obx1, obx2, obx3},
		options:  allowNullHeaderDisabled,
		wantErr:  true,
	}, {
		name:     "AllowNullHeader enabled",
		segments: [][]byte{pid, nk1, obx1, obx2, obx3},
		options:  allowNullHeaderEnabled,
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m, err := ParseMessageWithOptions(bytes.Join(tc.segments, segmentTerminatorBytes), tc.options)
			gotErr := err != nil
			if gotErr != tc.wantErr {
				t.Fatalf("ParseMessageWithOptions() got error %v, want error? %t", err, tc.wantErr)
			}
			if err != nil || tc.wantErr {
				return
			}

			// We use the PID segment to make sure the message is still parsable, but we could have used any other segment.
			got, err := m.PID()
			if err != nil {
				t.Fatalf("PID() failed with %v", err)
			}

			if diff := cmp.Diff(wantPID, got); diff != "" {
				t.Errorf("PID() diff: (-want, +got):\n%s", diff)
			}
		})
	}
}
