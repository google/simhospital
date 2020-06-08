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
	"strings"
	"testing"
)

// The tests in this file use methods in parser.go; the tests have been separated into a different file
// to differentiate functionality.

func TestParseMessageWithOptions_BrokenFieldsCanBeRewritten(t *testing.T) {
	badPid := []byte("PID|1||||||||||||^^^^NOT A NUMBER|")
	mshBadPid := bytes.Join([][]byte{msh, badPid}, []byte("\r"))
	m, err := ParseMessage(mshBadPid)
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	_, err = m.PID()
	if err == nil {
		t.Fatal("PID returned err=<nil>, want error")
	}
	rewrite := func(t Token) *RewriteResult {
		if t.Location == "PID-13-Phone Number - Home/XTN-5-Country Code" {
			return RewriteResultReplaceValue([]byte("1"))
		}
		return RewriteResultNoop()
	}
	mo := NewParseMessageOptions()
	mo.Rewrites = &[]Rewrite{rewrite}
	m, err = ParseMessageWithOptions(mshBadPid, mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions() failed with %v", err)
	}
	_, err = m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
}

func TestParseMessageWithOptions_Rewrite(t *testing.T) {
	pidWithoutGivenName := []byte("PID|1|||||Johnson")
	emptyPID := []byte("PID|1|")

	tests := []struct {
		name               string
		pid                []byte
		rewrite1           func(token Token) *RewriteResult
		rewrite2           func(token Token) *RewriteResult
		wantMothersSurname string
		wantMothersGiven   string
		checkPatientName   bool
		wantSurname        string
		wantGiven          string
		skipNameCheck      bool
	}{{
		name: "Single rewrite",
		pid:  pidWithoutGivenName,
		rewrite1: func(t Token) *RewriteResult {
			if t.Location == "PID-6-Mother'S Maiden Name" {
				return RewriteResultReplaceValue([]byte("Doe^Jane"))
			}
			return RewriteResultNoop()
		},
		wantMothersGiven:   "Jane",
		wantMothersSurname: "Doe",
	}, {
		name: "Multiple rewrites in the same field",
		pid:  pidWithoutGivenName,
		rewrite1: func(t Token) *RewriteResult {
			if t.Location == "PID-6-Mother'S Maiden Name" {
				return RewriteResultReplaceValue([]byte("Surname^Name"))
			}
			return RewriteResultNoop()
		},
		rewrite2: func(token Token) *RewriteResult {
			if token.Location == "PID-6-Mother'S Maiden Name" {
				if got, want := string(token.Value), "Surname^Name"; got != want {
					t.Errorf("token.Value got %q, want %q", got, want)
				}
				return RewriteResultReplaceValue([]byte("Doe^Jane"))
			}
			return RewriteResultNoop()
		},
		wantMothersGiven:   "Jane",
		wantMothersSurname: "Doe",
		wantSurname:        "Johnson",
		wantGiven:          "",
	}, {
		name: "Multiple rewrites in different fields",
		pid:  pidWithoutGivenName,
		rewrite1: func(t Token) *RewriteResult {
			if t.Location == "PID-6-Mother'S Maiden Name" {
				return RewriteResultReplaceValue([]byte("Doe^Jane"))
			}
			return RewriteResultNoop()
		},
		rewrite2: func(t Token) *RewriteResult {
			if t.Location == "PID-5-Patient Name" {
				return RewriteResultReplaceValue([]byte("Doe^Jane"))
			}
			return RewriteResultNoop()
		},
		wantMothersSurname: "Doe",
		wantMothersGiven:   "Jane",
		checkPatientName:   true,
		wantSurname:        "Doe",
		wantGiven:          "Jane",
	}, {
		name: "Entire segment",
		pid:  emptyPID,
		rewrite1: func(t Token) *RewriteResult {
			isPIDSegment := strings.Index(string(t.Value), "PID") == 0
			if isPIDSegment {
				return RewriteResultReplaceValue([]byte("PID|1|||||Doe^Jane"))
			}
			return RewriteResultNoop()
		},
		wantMothersGiven:   "Jane",
		wantMothersSurname: "Doe",
		skipNameCheck:      true,
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			marshalled := bytes.Join([][]byte{msh, tc.pid}, []byte("\r"))
			m, err := ParseMessage(marshalled)
			if err != nil {
				t.Fatalf("ParseMessage() failed with %v", err)
			}
			pid, err := m.PID()
			if err != nil {
				t.Fatalf("PID() failed with %v", err)
			}
			if !tc.skipNameCheck {
				if got := pid.MotherSMaidenName[0].GivenName; got != nil {
					t.Errorf("GivenName is %q, want <nil>", got)
				}
				if got, want := string(*pid.MotherSMaidenName[0].FamilyName.Surname), "Johnson"; got != want {
					t.Errorf("Surname got %q, want %q", got, want)
				}
			}

			mo := NewParseMessageOptions()
			if tc.rewrite2 != nil {
				mo.Rewrites = &[]Rewrite{tc.rewrite1, tc.rewrite2}
			} else {
				mo.Rewrites = &[]Rewrite{tc.rewrite1}
			}
			m, err = ParseMessageWithOptions(marshalled, mo)
			if err != nil {
				t.Fatalf("ParseMessageWithOptions() failed with %v", err)
			}
			pid, err = m.PID()
			if err != nil {
				t.Fatalf("PID() failed with %v", err)
			}
			if got, want := string(*pid.MotherSMaidenName[0].FamilyName.Surname), tc.wantMothersSurname; got != want {
				t.Errorf("Mother's surname got %q, want %q", got, want)
			}
			if got, want := string(*pid.MotherSMaidenName[0].GivenName), tc.wantMothersGiven; got != want {
				t.Errorf("Mother's GivenName got %q, want %q", got, want)
			}
			if tc.checkPatientName {
				if got, want := string(*pid.PatientName[0].FamilyName.Surname), tc.wantSurname; got != want {
					t.Errorf("Surname got %q, want %q", got, want)
				}
				if got, want := string(*pid.PatientName[0].GivenName), tc.wantGiven; got != want {
					t.Errorf("GivenName got %q, want %q", got, want)
				}
			}
		})
	}
}

func TestParseMessageWithOptions_ValuesCanBeNulled(t *testing.T) {
	pidWithAddressType := []byte("PID|1||||||||||^^^^^^AddressType")
	marshalled := bytes.Join([][]byte{msh, pidWithAddressType}, []byte("\r"))
	m, err := ParseMessage(marshalled)
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if got, want := string(*pid.PatientAddress[0].AddressType), "AddressType"; got != want {
		t.Errorf("AddressType got %q, want %q", got, want)
	}
	rewrite := func(t Token) *RewriteResult {
		if t.Location == "PID-11-Patient Address/XAD-7-Address Type" {
			return RewriteResultDeleteToken()
		}
		return RewriteResultNoop()
	}
	mo := NewParseMessageOptions()
	mo.Rewrites = &[]Rewrite{rewrite}
	m, err = ParseMessageWithOptions(marshalled, mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions() failed with %v", err)
	}
	pid, err = m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if got := pid.PatientAddress[0].AddressType; got != nil {
		t.Errorf("pid.PatientAddress[0].AddressType is %q, want <nil>", got)
	}
}

func TestParseMessageWithOptions_AbsentValuesCanBeMadePresent(t *testing.T) {
	pidWithoutZip := []byte("PID|1||||||||||^^^^^|")
	marshalled := bytes.Join([][]byte{msh, pidWithoutZip}, []byte("\r"))
	m, err := ParseMessage(marshalled)
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if got := pid.PatientAddress[0].ZipOrPostalCode; got != nil {
		t.Errorf("pid.PatientAddress[0].ZipOrPostalCode is %q, want <nil>", got)
	}
	rewrite := func(t Token) *RewriteResult {
		if t.Location == "PID-11-Patient Address/XAD-5-Zip Or Postal Code" {
			return RewriteResultReplaceValue([]byte("49999"))
		}
		return RewriteResultNoop()
	}
	mo := NewParseMessageOptions()
	mo.Rewrites = &[]Rewrite{rewrite}
	m, err = ParseMessageWithOptions(marshalled, mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions() failed with %v", err)
	}
	pid, err = m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if got, want := string(*pid.PatientAddress[0].ZipOrPostalCode), "49999"; got != want {
		t.Errorf("pid.PatientAddress[0].ZipOrPostalCode got %q, want %q", got, want)
	}
}

func TestParseMessageWithOptions_EntireSegmentsCanBeDeleted(t *testing.T) {
	pid1 := []byte("PID|1|")
	marshalled := bytes.Join([][]byte{msh, pid1}, []byte("\r"))
	m, err := ParseMessage(marshalled)
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	// Try all the ways a segment can be retrieved.
	pid, err := m.PID()
	if pid == nil {
		t.Error("PID returned <nil>, want not nil")
	}
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	pids, err := m.AllPID()
	if got, want := len(pids), 1; got != want {
		t.Errorf("AllPID returned length %d, want length %d", got, want)
	}
	if err != nil {
		t.Fatalf("AllPID() failed with %v", err)
	}
	all, err := m.All()
	if got, want := len(all), 2; got != want {
		t.Errorf("All() returned length %d, want length %d", got, want)
	}
	if err != nil {
		t.Fatalf("All() failed with %v", err)
	}

	rewrite := func(t Token) *RewriteResult {
		isPIDSegment := strings.Index(string(t.Value), "PID") == 0
		if isPIDSegment {
			return RewriteResultDeleteToken()
		}
		return RewriteResultNoop()
	}
	mo := NewParseMessageOptions()
	mo.Rewrites = &[]Rewrite{rewrite}
	m, err = ParseMessageWithOptions(marshalled, mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions() failed with %v", err)
	}
	// Try all the ways a segment can be retrieved.
	pid, err = m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid != nil {
		t.Errorf("PID returned %v, want <nil>", pid)
	}
	pids, err = m.AllPID()
	if got, want := len(pids), 0; got != want {
		t.Errorf("AllPID returned length %d, want empty", got)
	}
	if err != nil {
		t.Fatalf("AllPID() failed with %v", err)
	}
	all, err = m.All()
	if err != nil {
		t.Fatalf("All() failed with %v", err)
	}
	if got, want := len(all), 1; got != want {
		t.Errorf("All() returned length %d, want length %d", got, want)
	}
}
