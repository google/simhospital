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

package hardcoded

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/generator/header"
	"github.com/google/simhospital/pkg/hl7"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/message"
	"github.com/google/simhospital/pkg/test/testwrite"
)

const (
	nhsName = "InvalidNHSNum"
	// nhsYML is an ADT A01 admission message inline with HL7 standard:
	// http://hl7-definition.caristix.com:9010/HL7%20v2.3/triggerEvent/ADT_A01
	// The NHS number in PID.3 segment is invalid.
	nhsYML = nhsName + `:
  segments:
    - "MSH|^~\\&|SIMHOSP|SFAC|RAPP|RFAC|%s||ADT^A01|%s|T|2.3|||AL||44|ASCII"
    - "EVN|A01|20180709162444|||216865551019^ZZZDOCTOR^Foo^Bar Baz^^^^^DRNBR^PRSNL^^^ORGDR^|"
    - "PID||20180709162444249^^^^MRN|20180709162444249^^^^MRN~2278322867^^^NHSNMBR||FN20180709162444^Invalid NHS^^^MR^^CURRENT||19610306|M|||1L20180709162444^2L20180709162444^3L20180709162444^^09162444^C20180709162444^HOME^^|||||||||||P||||||||"
    - "PD1|||^^E83039|G9601039^Subel^Barry||||0||||"
    - "PV1|1|INPATIENT|ED^ED^Chair 06^Simulated Hospital^^BED^^1^||||123^dr. J Smith|||||||||||INPATIENT|2018070916244445^^^^visitid|||||||||||||||||||||||||20180403000000||||||||"`
	nhsFieldsSet = "MSH|^~\\&|SIMHOSP|SFAC|RAPP|RFAC|20180212000000||ADT^A01|1|T|2.3|||AL||44|ASCII\rEVN|A01|20180709162444|||216865551019^ZZZDOCTOR^Foo^Bar Baz^^^^^DRNBR^PRSNL^^^ORGDR^|\rPID||20180709162444249^^^^MRN|20180709162444249^^^^MRN~2278322867^^^NHSNMBR||FN20180709162444^Invalid NHS^^^MR^^CURRENT||19610306|M|||1L20180709162444^2L20180709162444^3L20180709162444^^09162444^C20180709162444^HOME^^|||||||||||P||||||||\rPD1|||^^E83039|G9601039^Subel^Barry||||0||||\rPV1|1|INPATIENT|ED^ED^Chair 06^Simulated Hospital^^BED^^1^||||123^dr. J Smith|||||||||||INPATIENT|2018070916244445^^^^visitid|||||||||||||||||||||||||20180403000000||||||||"
	oruName      = "InvalidORU_MissingPlacerAndFiller"
	// oruYML is an ORU R01 observation message inline with HL7 standard:
	// http://hl7-definition.caristix.com:9010/Default.aspx?version=HL7%20v2.3&triggerEvent=ORU_R01
	// The Placer and Filler numbers in ORB.2 and ORB.3 segments are not provided.
	oruYML = oruName + `:
  segments:
    - "MSH|^~\\&|SIMHOSP|SFAC|RAPP|RFAC|%s||ORU^R01|%s|T|2.3|||AL||44|ASCII"
    - "PID||20180709163256171^^^^MRN|20180709163256171^^^^MRN~2278322869^^^^NHSNMBR||FN20180709163256^Invalid ORU^^^MR^^CURRENT||19990804|M|||1L20180709163256^2L20180709163256^3L20180709163256^^09163256^C20180709163256^HOME^^|||||||^^^SIMHOSP^||||P||||||||"
    - "PV1|1|INPATIENT|ED^ED^Chair 06^Simulated Hospital^^BED^^2^||||123^dr. J Smith|||161||||19||||INPATIENT|2018070916325657^^^^visitid|||||||||||||||||||||||||20180403000000||||||||"
    - "ORC|RE||||||||20180709163256||||||20180709163256||||"
    - "OBR|1|||us-0003^UREA AND ELECTROLYTES||20180709163256|20180709163256|||||||||456^DR. J.G. BLACK|||||306360202|20180709163256||1|f||1^^^20180709163256^^r|^||"
    - "OBX|1|NM|tt-0003-01^Creatinine^||126|UMOLL|49-92|HH|||F|||||"
    - "OBX|2|NM|tt-0003-03^Sodium^||10|MMOLL|135 - 145|LL|||F|||||"
    - "OBX|3|NM|tt-0003-02^Potassium^||3.75|MMOLL|3.5 - 5.3||||F|||||"
    - "OBX|4|NM|tt-0003-05^eGFR (MDRD)^||None|MLMIN|[ ]|None|||F|||||"
    - "OBX|5|NM|tt-0003-04^Urea^||7.97|MMOLL|2.5 - 7.8||||F|||||"`
	oruFieldsSet   = "MSH|^~\\&|SIMHOSP|SFAC|RAPP|RFAC|20180212000000||ORU^R01|1|T|2.3|||AL||44|ASCII\rPID||20180709163256171^^^^MRN|20180709163256171^^^^MRN~2278322869^^^^NHSNMBR||FN20180709163256^Invalid ORU^^^MR^^CURRENT||19990804|M|||1L20180709163256^2L20180709163256^3L20180709163256^^09163256^C20180709163256^HOME^^|||||||^^^SIMHOSP^||||P||||||||\rPV1|1|INPATIENT|ED^ED^Chair 06^Simulated Hospital^^BED^^2^||||123^dr. J Smith|||161||||19||||INPATIENT|2018070916325657^^^^visitid|||||||||||||||||||||||||20180403000000||||||||\rORC|RE||||||||20180709163256||||||20180709163256||||\rOBR|1|||us-0003^UREA AND ELECTROLYTES||20180709163256|20180709163256|||||||||456^DR. J.G. BLACK|||||306360202|20180709163256||1|f||1^^^20180709163256^^r|^||\rOBX|1|NM|tt-0003-01^Creatinine^||126|UMOLL|49-92|HH|||F|||||\rOBX|2|NM|tt-0003-03^Sodium^||10|MMOLL|135 - 145|LL|||F|||||\rOBX|3|NM|tt-0003-02^Potassium^||3.75|MMOLL|3.5 - 5.3||||F|||||\rOBX|4|NM|tt-0003-05^eGFR (MDRD)^||None|MLMIN|[ ]|None|||F|||||\rOBX|5|NM|tt-0003-04^Urea^||7.97|MMOLL|2.5 - 7.8||||F|||||"
	missingPIDName = "MissingPIDMessageName"
	missingPIDYml  = missingPIDName + `:
  segments:
    - "MSH|^~\\&|sending_application_reliable|sending_facility|receiving_application|receiving_facility|%s||ADT^A03|%s|T|2.3|||AL||44|ASCII"
    - "EVN|A03|20180212000000|||id-1^surname-1^firstname-1^^^prefix-1^^^DRNBR^PRSNL^^^ORGDR|"
    - "PID_SEGMENT_PLACEHOLDER"
    - "PD1|||Test Primary Facility^^123|"
    - "PV1|1|PREADMIT||28b|||id-1^surname-1^firstname-1^^^prefix-1^^^DRNBR^PRSNL^^^ORGDR|||specialty-1||||||||PREADMIT|3448412528941593955^^^^visitid||||||||||||||||||||||FINISHED|||20180212000000|20180212000000|"`
	missingPIDFieldsSet = "MSH|^~\\&|sending_application_reliable|sending_facility|receiving_application|receiving_facility|20180212000000||ADT^A03|1|T|2.3|||AL||44|ASCII\rEVN|A03|20180212000000|||id-1^surname-1^firstname-1^^^prefix-1^^^DRNBR^PRSNL^^^ORGDR|\rPID|1|98186149572547538^^^SIMULATOR MRN^MRN|98186149572547538^^^SIMULATOR MRN^MRN~8676784620^^^NHSNBR^NHSNMBR||Johnson^John^^^^^CURRENT|||1|||Carbeth Road^20 Bull Lane^London^^SL9 8RA^GBR^HOME|||||||||||||||||||\rPD1|||Test Primary Facility^^123|\rPV1|1|PREADMIT||28b|||id-1^surname-1^firstname-1^^^prefix-1^^^DRNBR^PRSNL^^^ORGDR|||specialty-1||||||||PREADMIT|3448412528941593955^^^^visitid||||||||||||||||||||||FINISHED|||20180212000000|20180212000000|"
	// Unparsable message because it has receiving_facility (string) in the MSH.6 field (timestamp).
	unparsableYML = `Unparsable:
  segments:
    - "MSH|^~\\&||sending_application_reliable|sending_facility|receiving_application|receiving_facility|%s||ADT^A03|%s|T|2.3|||AL||44|ASCII"
    - "EVN|A03|20180212000000|||id-1^surname-1^firstname-1^^^prefix-1^^^DRNBR^PRSNL^^^ORGDR|"
    - "PID_SEGMENT_PLACEHOLDER"
    - "PD1|||Test Primary Facility^^123|"
    - "PV1|1|PREADMIT||28b|||id-1^surname-1^firstname-1^^^prefix-1^^^DRNBR^PRSNL^^^ORGDR|||specialty-1||||||||PREADMIT|3448412528941593955^^^^visitid||||||||||||||||||||||FINISHED|||20180212000000|20180212000000|"`
	noMessageTypeYML = `NoMessageType:
  segments:
    - "MSH|^~\\&|sending_application_reliable|sending_facility|receiving_application|receiving_facility|%s|||%s|T|2.3|||AL||44|ASCII"
    - "EVN|A03|20180212000000|||id-1^surname-1^firstname-1^^^prefix-1^^^DRNBR^PRSNL^^^ORGDR|"
    - "PID_SEGMENT_PLACEHOLDER"
    - "PD1|||Test Primary Facility^^123|"
    - "PV1|1|PREADMIT||28b|||id-1^surname-1^firstname-1^^^prefix-1^^^DRNBR^PRSNL^^^ORGDR|||specialty-1||||||||PREADMIT|3448412528941593955^^^^visitid||||||||||||||||||||||FINISHED|||20180212000000|20180212000000|"`
)

var arbitraryTime = time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC)

func TestMain(m *testing.M) {
	hl7.TimezoneAndLocation("Europe/London")
	retCode := m.Run()
	os.Exit(retCode)
}

func TestNewManager(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		description string
		yml         string
		wantErr     bool
	}{{
		description: "success",
		yml:         missingPIDYml,
	}, {
		description: "no hardcoded messages",
		yml:         "",
		wantErr:     true,
	}, {
		description: "re-declared message name",
		yml:         fmt.Sprintf("%s\n\n%s", missingPIDYml, missingPIDYml),
		wantErr:     true,
	}}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			dir := writeYmlToFile(t, tc.yml)
			mcg := &header.MessageControlGenerator{}
			if _, err := NewManager(ctx, dir, mcg); (err != nil) != tc.wantErr {
				t.Errorf("NewManager(%q, %v) got err %v, want err? %t", tc.yml, mcg, err, tc.wantErr)
			}
		})
	}
}

func TestFileExtensionValidation(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		wantErr bool
	}{{
		name: "valid*.yml",
	}, {
		name: "valid*.yaml",
	}, {
		name:    "no_extension",
		wantErr: true,
	}, {
		name:    "invalid_extension*.jpg",
		wantErr: true,
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dir := testwrite.BytesToDir(t, []byte(nhsYML), tc.name)

			mcg := &header.MessageControlGenerator{}
			mgr, err := NewManager(ctx, dir, mcg)

			// If a filename is valid, we expect it to be parsed into messages.
			// Otherwise, we expect an error because no files were parsed.
			if (err != nil) != tc.wantErr {
				t.Errorf("NewManager(%s, %v) got err %v, want err? %t", dir, mcg, err, tc.wantErr)
			}

			if err != nil || tc.wantErr {
				return
			}

			want := 1
			if got := len(mgr.messages); got != want {
				t.Errorf("len(mgr.messages) got %d messages, want %d", got, want)
			}
		})
	}
}

func TestMessage(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		description string
		yml         string
		toInclude   string
		want        *message.HL7Message
	}{{
		description: missingPIDName,
		yml:         missingPIDYml,
		toInclude:   ".*",
		want: &message.HL7Message{
			Type:    &message.Type{MessageType: "ADT", TriggerEvent: "A03"},
			Message: missingPIDFieldsSet,
		},
	}, {
		description: oruName,
		yml:         oruYML,
		toInclude:   ".*ORU.*",
		want: &message.HL7Message{
			Type:    &message.Type{MessageType: "ORU", TriggerEvent: "R01"},
			Message: oruFieldsSet,
		},
	}, {
		description: nhsName,
		yml:         nhsYML,
		toInclude:   ".*NHS.*,.*ORU.*",
		want: &message.HL7Message{
			Type:    &message.Type{MessageType: "ADT", TriggerEvent: "A01"},
			Message: nhsFieldsSet,
		},
	}, {
		description: oruName,
		yml:         oruYML,
		toInclude:   ".*NHS.*,.*ORU.*",
		want: &message.HL7Message{
			Type:    &message.Type{MessageType: "ORU", TriggerEvent: "R01"},
			Message: oruFieldsSet,
		},
	}, {
		description: fmt.Sprintf("%s,%s", oruName, nhsName),
		yml:         fmt.Sprintf("%s\n%s", oruYML, nhsYML),
		toInclude:   ".*ORU.*",
		want: &message.HL7Message{
			Type:    &message.Type{MessageType: "ORU", TriggerEvent: "R01"},
			Message: oruFieldsSet,
		},
	}}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			dir := writeYmlToFile(t, tc.yml)
			mcg := &header.MessageControlGenerator{}
			mgr, err := NewManager(ctx, dir, mcg)
			if err != nil {
				t.Fatalf("NewManager(%s, %v) failed with %v", tc.yml, mcg, err)
			}
			p := testPerson()
			got, err := mgr.Message(tc.toInclude, p, arbitraryTime)
			if err != nil {
				t.Fatalf("Message(%q, %v, %v) failed with %v", tc.toInclude, p, arbitraryTime, err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Message(%q, %v, %v) returned diff (-want +got):\n%s", tc.toInclude, p, arbitraryTime, diff)
			}
		})
	}
}

func TestMessageReturnsUniqueMessages(t *testing.T) {
	ctx := context.Background()
	dir := writeYmlToFile(t, missingPIDYml)
	mcg := &header.MessageControlGenerator{}
	mgr, err := NewManager(ctx, dir, mcg)
	if err != nil {
		t.Fatalf("NewManager(%s, %v) failed with %v", missingPIDYml, mcg, err)
	}

	p := testPerson()
	msg1, err := mgr.Message(missingPIDName, p, arbitraryTime)
	if err != nil {
		t.Fatalf("Message(%q, %v, %v) failed with %v", missingPIDName, p, arbitraryTime, err)
	}
	msg2, err := mgr.Message(missingPIDName, p, arbitraryTime)
	if err != nil {
		t.Fatalf("Message(%q, %v, %v) failed with %v", missingPIDName, p, arbitraryTime, err)
	}
	if msg1.Message == msg2.Message {
		t.Errorf("Message(%q, %v, %v) got the same message twice: %s, want messages with different headers", missingPIDName, p, arbitraryTime, msg1.Message)
	}
}

func TestMessageErrors(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		description string
		yml         string
		toInclude   string
	}{{
		description: fmt.Sprintf("no matching messages: %s", nhsName),
		yml:         nhsYML,
		toInclude:   ".*ORU.*",
	}, {
		description: fmt.Sprintf("no matching messages: %s", missingPIDName),
		yml:         missingPIDYml,
		toInclude:   ".*NHS.*,.*ORU.*",
	}, {
		description: "empty regex",
		yml:         nhsYML,
		toInclude:   "",
	}, {
		description: "list of empty or invalid regex",
		yml:         nhsYML,
		toInclude:   ",[",
	}, {
		description: "unparsable message",
		yml:         unparsableYML,
		toInclude:   ".*",
	}, {
		description: "no message type",
		yml:         noMessageTypeYML,
		toInclude:   ".*",
	}}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("NewManager(%q)", tc.description), func(t *testing.T) {
			dir := writeYmlToFile(t, tc.yml)
			mcg := &header.MessageControlGenerator{}
			mgr, err := NewManager(ctx, dir, mcg)
			if err != nil {
				t.Fatalf("NewManager(%s, %v) failed with %v", tc.yml, mcg, err)
			}

			p := testPerson()
			if _, err := mgr.Message(tc.toInclude, p, arbitraryTime); err == nil {
				t.Errorf("Message(%q, %v, %v) got nil error, want non-nil", tc.toInclude, p, arbitraryTime)
			}
		})
	}
}

func writeYmlToFile(t *testing.T, ymls ...string) string {
	t.Helper()
	yml := strings.Join(ymls, "\r")
	return testwrite.BytesToDir(t, []byte(yml), "hardcoded_messages.yml")
}

func testPerson() *ir.Person {
	return &ir.Person{
		MRN:       "98186149572547538",
		NHS:       "8676784620",
		Surname:   "Johnson",
		FirstName: "John",
		Address: &ir.Address{
			FirstLine:  "Carbeth Road",
			SecondLine: "20 Bull Lane",
			City:       "London",
			PostalCode: "SL9 8RA",
			Country:    "GBR",
			Type:       "HOME",
		},
		Gender: "1",
	}
}
