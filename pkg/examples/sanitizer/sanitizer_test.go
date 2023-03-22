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

package sanitizer

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/google/simhospital/pkg/hl7"
	"github.com/google/simhospital/pkg/test/testmetrics"
)

const (
	testMSH         = "MSH|^~\\&|CERNER|SF|CARELINK|AKI2|20141128001635||ORU^R01|2014112800163507740000|T|2.3|||AL||44|ASCII"
	testMSHNoRFH    = "MSH|^~\\&|CERNER|HNAM|CARELINK|AKI2|20141128001635||ORU^R01|2014112800163507740000|T|2.3|||AL||44|ASCII"
	testORURO3MSH   = "MSH|^~\\&|HNAM|SF|SHIP|TQUEST|20160717150510||ORU^R03|Q500000486T598370000||2.3"
	testSFORUR01MSH = "MSH|^~\\&|SOLITON|SF|FHIRAPI|OS|20170601171317||ORU^R01|1990855733|P|2.3"
	testPID         = "PID|1|||||"
	//                               ┌ Address 0: empty. This makes sure the indexes are incremented correctly.
	//                               |┌ Address 1: zip code and no address type
	//                               ||      ┌ Address 2: zip code with address type
	//                               ||      |              ┌ Address 3: address type field, to become zip code
	//                               ||      |              |              ┌ Address 4: no zip code, no address type
	testPIDAddress = "PID|1||||||||||~^^^^29~^^^^39^^HOME^^~^^^^^^N1C 4AB^~^^^^^^||"
	//                               ┌ Responsible Observers
	testOBX1 = "OBX|1|FT||||||||||||||C1111111^Doe^John^Joe^^^^^HOSPITAL^CD:222222^^^CD:1111^~ZZZZ^Doe^Jane^Dr^^^"
	testOBX2 = "OBX|1|FT||||||||||||||C3333333^Doe^Jane^Dr^^^^^HOSPITAL^CD:444444^^^CD:2222^~ZZZZ^Doe^Jane^Dr^^^"
	testOBX3 = "OBX|1|FT||||||||||||||C3333333^Doe^Jane^Dr^^^^^HOSPITAL^CD:444444"
	//                     ┌ Units
	testOBX4 = "OBX|1|FT||||a^b^c||||||||||C3333333^Doe^Jane^Dr^^^^^HOSPITAL^CD:444444"
	testOBX5 = "OBX|1|FT||||a||||||||||C3333333^Doe^Jane^Dr^^^^^HOSPITAL^CD:444444"
	testOBX6 = "OBX|1|FT||||a^^c||||||||||C3333333^Doe^Jane^Dr^^^^^HOSPITAL^CD:444444"
	//                ┌ Universal Service ID
	//                |                       ┌ Quantity/Timing
	testOBR = "OBR|1|||1|||||||||||||||||||||||1^^20130815114600CD:61640612~^^^^R|||||||||"

	rewriteWarningMetricName = "message_sanitizer_rewrite_warning_total"
)

func TestMain(m *testing.M) {
	hl7.TimezoneAndLocation("Europe/London")
	retCode := m.Run()
	os.Exit(retCode)
}

func TestSanitizeMessage_PatientAddress_MultipleAddresses(t *testing.T) {
	var s MessageSanitizer

	msg := &Message{
		Message: strings.Join([]string{testMSH, "PID|1||||||||||~^^^^29~^^^^39^^HOME^^~^^^^^^N1C 4AB^~^^^^^^||"}, "\r"),
	}

	m, err := hl7.ParseMessage([]byte(msg.Message))
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid.PatientAddress[3].ZipOrPostalCode != nil {
		t.Errorf("pid.PatientAddress[3].ZipOrPostalCode is %q, want <nil>.", pid.PatientAddress[3].ZipOrPostalCode)
	}
	if got, want := pid.PatientAddress[3].AddressType.String(), "N1C 4AB"; got != want {
		t.Errorf("pid.PatientAddress[3].AddressType got %v, want %v.", got, want)
	}

	m, err = s.SanitizeMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("SanitizeMessage() failed with %v", err)
	}
	// If we parse the PID more than once we should get the same results.
	for i := 0; i < 2; i++ {
		pid, err = m.PID()
		if err != nil {
			t.Fatalf("PID() failed with %v", err)
		}
		if got, want := pid.PatientAddress[3].ZipOrPostalCode.String(), "N1C 4AB"; got != want {
			t.Errorf("pid.PatientAddress[3].ZipOrPostalCode got %v, want %v.", got, want)
		}
		if pid.PatientAddress[3].AddressType == nil {
			t.Error("pid.PatientAddress[3].AddressType is <nil>, want non nil.")
		}
		// Make sure other addresses did not change.
		if got, want := len(pid.PatientAddress), 5; got != want {
			t.Errorf("len(pid.PatientAddress) got %v, want %v.", got, want)
		}
		if got, want := pid.PatientAddress[1].ZipOrPostalCode.String(), "29"; got != want {
			t.Errorf("pid.PatientAddress[1].ZipOrPostalCode got %v, want %v.", got, want)
		}
		if got, want := pid.PatientAddress[2].ZipOrPostalCode.String(), "39"; got != want {
			t.Errorf("pid.PatientAddress[2].ZipOrPostalCode got %v, want %v.", got, want)
		}
		if got, want := pid.PatientAddress[2].AddressType.String(), "HOME"; got != want {
			t.Errorf("pid.PatientAddress[2].AddressType got %v, want %v.", got, want)
		}
		if got := pid.PatientAddress[4].ZipOrPostalCode; got != nil {
			t.Errorf("pid.PatientAddress[4].ZipOrPostalCode is %q, want <nil>.", got)
		}
		if got := pid.PatientAddress[4].AddressType; got != nil {
			t.Errorf("pid.PatientAddress[4].AddressType is %q, want <nil>.", got)
		}
	}
}

func TestSanitizeMessage_FixPostcodes(t *testing.T) {
	var s MessageSanitizer

	testCases := []struct {
		in   string
		want string
	}{
		// Rewritten.
		{"^^^^N1C 4ABN1C 4AB^^^^", "N1C 4AB"},
		{"^^^^n1C 4aBn1C 4aB^^^^", "n1C 4aB"},
		{"^^^^NE1 4ABNE1 4AB^^^^", "NE1 4AB"},
		{"^^^^N12 4ABN12 4AB^^^^", "N12 4AB"},
		{"^^^^N1C 4AB N1C 4AB^^^^", "N1C 4AB"},
		{"^^^^N1C 4ABN1C 4AB^^^^", "N1C 4AB"},
		{"^^^^N1C4ABN1C4AB^^^^", "N1C4AB"},
		{"^^^^N1 8TEN1 8TE^^^^", "N1 8TE"},
		{"^^^^NW1C 4ABNW1C 4AB^^^^", "NW1C 4AB"},
		{"^^^^N1C N1C^^^^", "N1C"},

		// Not rewritten.
		{"^^^^^^^^", ""},
		{"^^^^ ^^^^", " "},
		{"^^^^B^^^^", "B"},
		{"^^^^AB^^^^", "AB"},
		{"^^^^N1C 4AB^^^^", "N1C 4AB"},
		{"^^^^N1C 4AB.N1C 4AB^^^^", "N1C 4AB.N1C 4AB"},
		{"^^^^N12 ABCN12 ABC^^^^", "N12 ABCN12 ABC"},
		{"^^^^N1C 4ABN1C 4ABN1C 4AB^^^^", "N1C 4ABN1C 4ABN1C 4AB"},

		// Rewrite the postcode from other fields.
		{"LONDON N1C 4AB^^^^^^", "N1C 4AB"},
		{"^LONDON N1C 4AB^^^^^", "N1C 4AB"},
		{"^^LONDON N1C 4AB^^^^", "N1C 4AB"},
		{"^^^LONDON N1C 4AB^^^", "N1C 4AB"},
		{"^^^^^LONDON N1C 4AB", "N1C 4AB"},
		{"^^^^^^LONDON N1C 4AB", "N1C 4AB"},
		// Postcodes with no spaces.
		{"LONDON N1C4AB^^^^^^", "N1C4AB"},
		{"LONDON N1C4ABN1C4AB^^^^^^", "N1C4AB"},
		// Two postcodes return the first one.
		{"LONDON N1C 4ABLONDON S2X 8BE^^^^^^", "N1C 4AB"},
		// Postcodes with lowercase letters.
		{"LONDON n1C 4aB^^^^^^", "n1C 4aB"},

		// Priorities: Address type is the preferred field. After that, priorities go in order.
		{"LONDON S2X 8BE^^^^^^LONDON N1C 4AB", "N1C 4AB"},
		{"LONDON N1C 4AB^LONDON S2X 8BE^^^^^", "N1C 4AB"},
		{"^LONDON N1C 4AB^LONDON S2X 8BE^^^^", "N1C 4AB"},
		{"^^LONDON N1C 4AB^LONDON S2X 8BE^^^", "N1C 4AB"},
		{"^^^LONDON N1C 4AB^^LONDON S2X 8BE^^", "N1C 4AB"},

		// If the postcode field is present, it's not rewritten.
		{"LONDON N1C 4AB^^^^LONDON S2X 8BE^^", "LONDON S2X 8BE"},
		// Invalid postcodes are not rewritten.
		{"N1B 432^^^^^^", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			pidPatientAddressPlaceholder := "PID|||||||||||%s|||"
			msg := &Message{
				Message: strings.Join([]string{testMSH, fmt.Sprintf(pidPatientAddressPlaceholder, tc.in)}, "\r"),
			}
			m, err := hl7.ParseMessage([]byte(msg.Message))
			if err != nil {
				t.Fatalf("ParseMessage(%v) failed with %v", msg.Message, err)
			}
			_, err = m.PID()
			if err != nil {
				t.Fatalf("PID() failed with %v", err)
			}

			m, err = s.SanitizeMessage(context.Background(), msg)
			if err != nil {
				t.Fatalf("SanitizeMessage() failed with %v", err)
			}
			pid, err := m.PID()
			if err != nil {
				t.Fatalf("PID() (after sanitization) failed with %v", err)
			}
			got := ""
			if len(pid.PatientAddress) > 0 {
				got = pid.PatientAddress[0].ZipOrPostalCode.String()
			}
			if got != tc.want {
				t.Errorf("pid.PatientAddress[0].ZipOrPostalCode got %v, want %v.", got, tc.want)
			}
		})
	}
}

func TestSanitizeMessage_OBXUnescaping(t *testing.T) {
	for _, tt := range []struct {
		name string
		in   []string
		want [][]string
	}{
		{
			name: "obx with one escape sequence",
			in:   []string{`One \F\ Escape`},
			want: [][]string{
				{"One | Escape"},
			},
		}, {
			name: "obx with two escape sequence",
			in:   []string{`Two \F\ Escapes \S\`},
			want: [][]string{
				{"Two | Escapes ^"},
			},
		}, {
			name: "obx with one escape sequence but no spaces",
			in:   []string{`No spaces\F\between escapes`},
			want: [][]string{
				{"No spaces|between escapes"},
			},
		}, {
			name: "obx with escape sequence in the beginning of the field",
			in:   []string{`\F\Escape at index zero`},
			want: [][]string{
				{"|Escape at index zero"},
			},
		}, {
			name: "obx with new line escape",
			in:   []string{`result\.br\result`},
			want: [][]string{
				{"result\nresult"},
			},
		}, {
			name: "obx with repeated field and escape sequences",
			in:   []string{`one\.br\two~three\.br\four`},
			want: [][]string{
				{"one\ntwo", "three\nfour"},
			},
		}, {
			name: "obx with repeated field and escape sequences at the end",
			in:   []string{`foo\E\bar~bar\.br\`},
			want: [][]string{
				{"foo\\bar", "bar\n"},
			},
		}, {
			name: "empty value",
			in:   []string{""},
			want: [][]string{{}},
		}, {
			name: "multiple obxs",
			in: []string{
				`foo\E\bar~bar\.br\`,
				`one\.br\two~three\.br\four~\F\six`,
			},
			want: [][]string{
				{"foo\\bar", "bar\n"},
				{"one\ntwo", "three\nfour", "|six"},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			sanitizer := MessageSanitizer{UnescapeOBXValues: true}

			segments := []string{testMSH, testPIDAddress, testOBR}
			for _, v := range tt.in {
				segments = append(segments, fmt.Sprintf("OBX|1|TX|LABR5288^^Winpath||%s|||||||||", v))
			}
			msg := &Message{
				Message: strings.Join(segments, "\r"),
			}

			m, err := sanitizer.SanitizeMessage(context.Background(), msg)
			if err != nil {
				t.Fatalf("sanitizer.SanitizeMessage() failed with %v", err)
			}

			// Do this whole test multiple times to ensure that the unescaping is consistent across
			// multiple segment retrieval.
			for i := 0; i < 2; i++ {
				obxs, err := m.AllOBX()
				if err != nil {
					t.Fatalf("m.AllOBX() failed with %v", err)
				}

				if got, want := len(obxs), len(tt.in); got != want {
					t.Fatalf("len(obxs)=%d, want %d", got, want)
				}

				for obxIndex, obx := range obxs {
					if tt.in[obxIndex] == "" {
						if got, want := len(obx.ObservationValue), 0; got != want {
							t.Errorf("len(obx.ObservationValue)=%d, want %d", got, want)
						}
						return
					}

					if got, want := len(obx.ObservationValue), len(tt.want[obxIndex]); got != want {
						t.Fatalf("len(obxs[%d].ObservationValue)=%d, want %d", obxIndex, got, want)
					}

					for index, want := range tt.want[obxIndex] {
						if got := string(obx.ObservationValue[index]); got != want {
							t.Errorf("string(obx.ObservationValue[%d])=%q, want %q", index, got, want)
						}
					}
				}
			}
		})
	}
}

func TestSanitizeMessage(t *testing.T) {
	tcs := []struct {
		name          string
		msg           *Message
		wantErrName   string
		wantErrString string
	}{{
		name:          "No Sending Application",
		msg:           &Message{Message: strings.Join([]string{"MSH|^~\\&||SF|CARELINK|AKI2|20141128001635||ORU^R01|2014112800163507740000|T|2.3|||AL||44|ASCII"}, "\r")},
		wantErrName:   "MISSING_HL7_FIELD",
		wantErrString: "could not get sending application",
	}, {
		name:          "No DateTime",
		msg:           &Message{Message: strings.Join([]string{"MSH|^~\\&|CERNER|RF|CARELINK|AKI2|||ORU^R01|2014112800163507740000|T|2.3|||AL||44|ASCII"}, "\r")},
		wantErrName:   "MISSING_HL7_FIELD",
		wantErrString: "could not get message time stamp",
	}, {
		name:          "No MSH",
		msg:           &Message{Message: strings.Join([]string{testPIDAddress}, "\r")},
		wantErrName:   "HL7_PARSE_FAILURE",
		wantErrString: "MSH header",
	}, {
		name:          "No PID",
		msg:           &Message{Message: strings.Join([]string{"MSH|^~\\&|CERNER|SF|CARELINK|AKI2|20141128001635||ORU^R01|2014112800163507740000|T|2.3|||AL||44|ASCII"}, "\r")},
		wantErrName:   "NO_PID",
		wantErrString: "No PID segment",
	}, {
		name:          "No Sending Facility",
		msg:           &Message{Message: strings.Join([]string{"MSH|^~\\&|SA||CARELINK|AKI2|20141128001635||ORU^R01|2014112800163507740000|T|2.3|||AL||44|ASCII"}, "\r")},
		wantErrName:   "MISSING_HL7_FIELD",
		wantErrString: "could not get sending facility",
	}, {
		name:          "No Message Type",
		msg:           &Message{Message: strings.Join([]string{"MSH|^~\\&|SA|SF|CARELINK|AKI2|20141128001635|||2014112800163507740000|T|2.3|||AL||44|ASCII"}, "\r")},
		wantErrName:   "MISSING_HL7_FIELD",
		wantErrString: "could not get message type",
	}, {
		name:          "Too many ZCM",
		msg:           &Message{Message: strings.Join([]string{testMSH, testPID, "ZCM|1|"}, "\r")},
		wantErrName:   "MESSAGE_SANITIZATION_FAILURE",
		wantErrString: "expected 1 ZCM segment after sanitization, got 2",
	}, {
		name: "Unfixable PID",
		// PID with a very long and wrong Date of Death. These dates are not fixed during sanitization.
		msg:           &Message{Message: strings.Join([]string{testMSH, "PID|1||30711793^^^SF MRN^MRN||ZZZTEST^INFADMREGTWOUPAA^^^^^CURRENT||19491231|1||\"\"|^^^\"\"^^GBR^HOME^^\"\"~Chapel of St George^Central Terminal Area^HOUNSLOW^\"\"^TW6 1BP^GBR^HOME^London Heathrow Airport^\"\"||02033132587^HOME||\"\"|\"\"|\"\"|||||A||||\"\"|\"\"|\"\"|190005242515141234143141|"}, "\r")},
		wantErrName:   "PID_RETRIEVAL_FAILURE",
		wantErrString: "bad TS value: invalid length",
	}}

	var s MessageSanitizer

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_, err := s.SanitizeMessage(context.Background(), tc.msg)
			if err == nil {
				t.Fatal("SanitizeMessage, got err=<nil>, want error")
			}
			sErr, isSanitizerErr := err.(*Error)
			if !isSanitizerErr {
				t.Errorf("isSanitizerErr is false: %v", err)
			}
			if got, want := sErr.Name, tc.wantErrName; got != want {
				t.Errorf("sErr.Name got %q, want %q", got, want)
			}
			if !strings.Contains(err.Error(), tc.wantErrString) {
				t.Errorf("err got %q, want error that contains %q", err, tc.wantErrString)
			}
		})
	}
}

func TestSanitizeMessage_ZCMIsAppended(t *testing.T) {
	var s MessageSanitizer

	// Send a message without ZCM segment. The Sanitization will append the ZCM segment.
	msg := &Message{Message: strings.Join([]string{testMSH, testPID}, "\r")}
	m, err := s.SanitizeMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("SanitizeMessage() failed with %v", err)
	}
	zdm, err := m.ZCM()
	if err != nil {
		t.Fatalf("ZCM() failed with %v", err)
	}
	if zdm == nil {
		t.Error("zdm is <nil>, want non nil")
	}
}

func TestSanitizeMessage_SquashOBXes(t *testing.T) {
	var s MessageSanitizer

	testCases := []struct {
		name       string
		obxes      []string
		inSquash   bool
		wantSquash bool
		wantValue  string
	}{
		{
			name: "Squash",
			obxes: []string{
				"OBX|1||||Line 1|",
				"OBX|2||||Line 2|",
				"OBX|3||||Line 3|",
			},
			inSquash:   true,
			wantSquash: true,
			wantValue:  "Line 1\nLine 2\nLine 3",
		}, {
			name: "Squash with blank line",
			obxes: []string{
				"OBX|1||||Line 1|",
				"OBX|2|||||",
				"OBX|3||||Line 3|",
			},
			inSquash:   true,
			wantSquash: true,
			wantValue:  "Line 1\n\nLine 3",
		}, {
			name:     "SquashFullReport",
			inSquash: true,
			obxes: []string{
				"OBX|1|TX|CANKR^CT Ankle Rt||CANKR||||||F|||20160516145631",
				"OBX|2|TX|CANKR^CT Ankle Rt||Procedure Description: CT Ankle Rt||||||F|||20160516145631",
				"OBX|3|TX|CANKR^CT Ankle Rt||Patient Last Name:     YYYTEST||||||F|||20160516145631",
				"OBX|4|TX|CANKR^CT Ankle Rt||Patient First Name:    TESTFIVESOLITON||||||F|||20160516145631",
				"OBX|5|TX|CANKR^CT Ankle Rt||HIS Patient ID:        30370308||||||F|||20160516145631",
				"OBX|6|TX|CANKR^CT Ankle Rt||NHS Patient ID:        9464610409||||||F|||20160516145631",
				"OBX|7|TX|CANKR^CT Ankle Rt||Date of Birth:         17/07/1986||||||F|||20160516145631",
				"OBX|8|TX|CANKR^CT Ankle Rt||||||||F|||20160516145631",
				"OBX|9|TX|CANKR^CT Ankle Rt||Requested Date:        16/05/2016 14:53||||||F|||20160516145631",
				"OBX|10|TX|CANKR^CT Ankle Rt||Examination Date:      16/05/2016 14:54||||||F|||20160516145631",
				"OBX|11|TX|CANKR^CT Ankle Rt||||||||F|||20160516145631",
				"OBX|12|TX|CANKR^CT Ankle Rt||========== REPORT TEXT START ==========||||||F|||20160516145631",
				"OBX|13|TX|CANKR^CT Ankle Rt||||||||F|||20160516145631",
				"OBX|14|TX|CANKR^CT Ankle Rt||8500016 16/05/2016 CT Ankle Rt:||||||F|||20160516145631",
				"OBX|15|TX|CANKR^CT Ankle Rt||||||||F|||20160516145631",
				"OBX|16|TX|CANKR^CT Ankle Rt||Clinical History||||||F|||20160516145631",
				"OBX|17|TX|CANKR^CT Ankle Rt||||||||F|||20160516145631",
				"OBX|18|TX|CANKR^CT Ankle Rt||||||||F|||20160516145631",
				"OBX|19|TX|CANKR^CT Ankle Rt||Report||||||F|||20160516145631",
				"OBX|20|TX|CANKR^CT Ankle Rt||||||||F|||20160516145631",
				"OBX|21|TX|CANKR^CT Ankle Rt||RIS only test result||||||F|||20160516145631",
				"OBX|22|TX|CANKR^CT Ankle Rt||on NHS number patient||||||F|||20160516145631",
				"OBX|23|TX|CANKR^CT Ankle Rt||for interface testing||||||F|||20160516145631",
				"OBX|24|TX|CANKR^CT Ankle Rt||please ignore this result||||||F|||20160516145631",
				"OBX|25|TX|CANKR^CT Ankle Rt||||||||F|||20160516145631",
				"OBX|26|TX|CANKR^CT Ankle Rt||Dictated by Dr Soliton User||||||F|||20160516145631",
				"OBX|27|TX|CANKR^CT Ankle Rt||Authorised by Dr Soliton User  on 16/05/2016 at 14:56||||||F|||20160516145631",
				"OBX|28|TX|CANKR^CT Ankle Rt||||||||F|||20160516145631",
				"OBX|29|TX|CANKR^CT Ankle Rt||========== REPORT TEXT END ==========||||||F|||20160516145631",
			},
			wantSquash: true,
			wantValue: "CANKR\nProcedure Description: CT Ankle Rt\nPatient Last Name:     YYYTEST\nPatient First Name:    TESTFIVESOLITON\n" +
				"HIS Patient ID:        30370308\nNHS Patient ID:        9464610409\nDate of Birth:         17/07/1986\n\n" +
				"Requested Date:        16/05/2016 14:53\nExamination Date:      16/05/2016 14:54\n" +
				"\n========== REPORT TEXT START ==========\n\n8500016 16/05/2016 CT Ankle Rt:\n\nClinical History\n\n\n" +
				"Report\n\nRIS only test result\non NHS number patient\nfor interface testing\nplease ignore this result\n\n" +
				"Dictated by Dr Soliton User\nAuthorised by Dr Soliton User  on 16/05/2016 at 14:56\n\n" +
				"========== REPORT TEXT END ==========",
		}, {
			name: "SingleOBX",
			obxes: []string{
				"OBX|1||||Line 1|",
			},
			inSquash:   true,
			wantSquash: true,
			wantValue:  "Line 1",
		}, {
			name:       "NoOBXs",
			obxes:      []string{},
			inSquash:   true,
			wantSquash: false,
			wantValue:  "",
		}, {
			name: "DifferentUnits",
			obxes: []string{
				"OBX|1||||Line 1|Unit^Unit",
				"OBX|2||||Line 2|Unit^Unit",
				"OBX|3||||Line 3|Different unit^Different unit",
			},
			inSquash:   true,
			wantSquash: false,
			wantValue:  "",
		}, {
			name: "Don't squash",
			obxes: []string{
				"OBX|1||||Line 1|",
				"OBX|2||||Line 2|",
				"OBX|3||||Line 3|",
			},
			inSquash:   false,
			wantSquash: false,
			wantValue:  "",
		}}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s.ShouldSquashOBXs = func(_ *hl7.Message) bool { return tc.inSquash }

			lines := []string{testMSH, testPIDAddress}
			lines = append(lines, tc.obxes...)
			msg := &Message{
				Message: strings.Join(lines, "\r"),
			}

			_, err := hl7.ParseMessage([]byte(msg.Message))
			if err != nil {
				t.Fatalf("Unexpected error, err=%q.", err)
			}
			m, err := s.SanitizeMessage(context.Background(), msg)
			if err != nil {
				t.Fatalf("SanitizeMessage() failed with %v", err)
			}

			gotOBXes, err := m.AllOBX()
			if err != nil {
				t.Fatalf("AllOBX() failed with %v", err)
			}
			if tc.wantSquash {
				// If we're squashing there should be a single OBX, with the expectedValue.
				if len(gotOBXes) != 1 {
					t.Fatalf("len(actualOBXes) = %d, want %d", len(gotOBXes), 1)
				}
				if len(gotOBXes[0].ObservationValue) != 1 {
					t.Fatalf("len(actualOBXes[0].ObservationValue) = %d, want %d", len(gotOBXes[0].ObservationValue), 1)
				}
				if string(gotOBXes[0].ObservationValue[0]) != tc.wantValue {
					t.Errorf("Got %v, want %v.", string(gotOBXes[0].ObservationValue[0]), tc.wantValue)
				}
			} else {
				// Check that nothing changed.
				if len(gotOBXes) != len(tc.obxes) {
					t.Fatalf("len(actualOBXes) = %d, want %d", len(gotOBXes), len(tc.obxes))
				}
				for i, obx := range gotOBXes {
					segment, err := hl7.MarshalSegment(obx, m.Context)
					if err != nil {
						t.Errorf("OBX %v err=%q", i, err)
					}
					if string(segment) != tc.obxes[i] {
						t.Errorf("Got %v, want %v.", string(segment), tc.obxes[i])
					}
				}
			}
		})
	}
}

func TestSanitizeMessage_OldSendingFacility(t *testing.T) {
	var s MessageSanitizer
	s.OldSendingFacilities = map[string]string{
		"OLD": "NEW",
	}
	mshWithOldSendingFacility := "MSH|^~\\&|CERNER|OLD|CARELINK|AKI2|20141128001635||ORU^R01|2014112800163507740000|T|2.3|||AL||44|ASCII"
	msg := &Message{
		Message: strings.Join([]string{mshWithOldSendingFacility, testPIDAddress}, "\r"),
	}
	m, err := s.SanitizeMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("SanitizeMessage() failed with %v", err)
	}
	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if got, want := msh.SendingFacility.NamespaceID.String(), "NEW"; got != want {
		t.Errorf("msh.SendingFacility.NamespaceID got %v, want %v.", got, want)
	}
}

func TestSanitizeMessage_SendingFacilityNotRewritten(t *testing.T) {
	var s MessageSanitizer
	mshIn := "MSH|^~\\&|CERNER|SF|CARELINK|AKI2|20141128001635||ORU^R01|2014112800163507740000|T|2.3|||AL||44|ASCII"
	msg := &Message{
		Message: strings.Join([]string{mshIn, testPIDAddress}, "\r"),
	}
	m, err := s.SanitizeMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("SanitizeMessage() failed with %v", err)
	}
	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh.SendingFacility.NamespaceID.String() != "SF" {
		t.Errorf("Got %v, want %v.", msh.SendingFacility.NamespaceID.String(), "SF")
	}
}

func TestSanitizeMessage_MalformedGP(t *testing.T) {
	var s MessageSanitizer

	//                      ┌ A number is expected here.
	pd1 := "PD1|\"\"|\"\"|^^V81999|G9999998^NotKnown^GP code not known|\"\"||\"\"|0"
	msg := &Message{
		Message: strings.Join([]string{testMSH, testPIDAddress, pd1}, "\r"),
	}
	m, err := hl7.ParseMessage([]byte(msg.Message))
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	_, err = m.PD1()
	if err == nil {
		t.Errorf("Expected error, got err=<nil>")
	}
	m, err = s.SanitizeMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("SanitizeMessage() failed with %v", err)
	}
	parsedPD1, err := m.PD1()
	if err != nil {
		t.Errorf("Unexpected error, err=%q.", err)
	}
	zdm, err := m.ZCM()
	if err != nil {
		t.Fatalf("ZCM() failed with %v", err)
	}

	if zdm.PatientPrimaryFacilityIDNumber.String() != "V81999" {
		t.Errorf("Got %v, want %v.", zdm.PatientPrimaryFacilityIDNumber.String(), "V81999")
	}
	if parsedPD1.PatientPrimaryFacility[0].IDNumber != nil {
		t.Errorf("parsedPD1.PatientPrimaryFacility[0].IDNumber is %v, want <nil>.", parsedPD1.PatientPrimaryFacility[0].IDNumber)
	}

	// Send multiple Patient Primary Facilities: the first ID is used in all of them.
	pd1 = "PD1|\"\"|\"\"|^^V81999~^^V33444|G9999998^NotKnown^GP code not known|\"\"||\"\"|0"
	msg = &Message{
		Message: strings.Join([]string{testMSH, testPIDAddress, pd1}, "\r"),
	}
	m, err = hl7.ParseMessage([]byte(msg.Message))
	if err != nil {
		t.Errorf("Unexpected error, err=%q.", err)
	}
	_, err = m.PD1()
	if err == nil {
		t.Errorf("Expected error, got err=<nil>")
	}
	m, err = s.SanitizeMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("SanitizeMessage() failed with %v", err)
	}
	parsedPD1, err = m.PD1()
	if err != nil {
		t.Errorf("Unexpected error, err=%q.", err)
	}
	zdm, err = m.ZCM()
	if err != nil {
		t.Fatalf("ZCM() failed with %v", err)
	}

	if zdm.PatientPrimaryFacilityIDNumber.String() != "V81999" {
		t.Errorf("Got %v, want %v.", zdm.PatientPrimaryFacilityIDNumber.String(), "V81999")
	}
	if parsedPD1.PatientPrimaryFacility[0].IDNumber != nil {
		t.Errorf("parsedPD1.PatientPrimaryFacility[0].IDNumber is %v, want <nil>.", parsedPD1.PatientPrimaryFacility[0].IDNumber)
	}
	if parsedPD1.PatientPrimaryFacility[1].IDNumber != nil {
		t.Errorf("parsedPD1.PatientPrimaryFacility[1].IDNumber is %v, want <nil>.", parsedPD1.PatientPrimaryFacility[1].IDNumber)
	}
}

func TestSanitizeMessage_DateOfBirthPrecisionTruncated(t *testing.T) {
	var s MessageSanitizer
	highPrecisionDateOfBirthPID := "PID|||||||19880705000000|"
	msh := "MSH|^~\\&|CERNER|SF|CARELINK|AKI2|20141128001635||ORU^R01|2014112800163507740000|T|2.3|||AL||44|ASCII"

	msg := &Message{
		Message: strings.Join([]string{msh, highPrecisionDateOfBirthPID}, "\r"),
	}
	m, err := s.SanitizeMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("SanitizeMessage() failed with %v", err)
	}
	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if got, want := pid.DateTimeOfBirth.Precision, hl7.DayPrecision; got != want {
		t.Errorf("pid.DateTimeOfBirth.Precision=%v, want %v", got, want)
	}
	if pid.DateTimeOfBirth.Time.Format("20060102") != "19880705" {
		t.Errorf("Got %v, want %v.", pid.DateTimeOfBirth.Time.Format("20060102"), "19880705")
	}
}

func TestSanitizeMessage_DeathIndicator(t *testing.T) {
	tests := []struct {
		in                  string
		m                   map[string]string
		want                string
		wantMetricIncrement float64
	}{
		{in: "yes", want: "Y", m: CommonNormalizedDeathIndicator, wantMetricIncrement: 0},
		{in: "no", want: "N", m: CommonNormalizedDeathIndicator, wantMetricIncrement: 0},
		{in: "yes", want: "Y", m: map[string]string{"yes": "Y"}, wantMetricIncrement: 0},
		// A present map and an unrecognised input should increment the counter.
		{in: "not-included", want: "", m: map[string]string{"yes": "Y"}, wantMetricIncrement: 1},
		{in: "passthrough", want: "passthrough", m: map[string]string{}, wantMetricIncrement: 0},
		{in: "passthrough", want: "passthrough", m: nil, wantMetricIncrement: 0},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			mr := testmetrics.NewRetrieverFromGatherer(t)
			s := MessageSanitizer{DeathIndicatorNormalizer: tc.m}
			msh := "MSH|^~\\&|CERNER|SF|CARELINK|AKI2|20141128001635||ORU^R01|2014112800163507740000|T|2.3|||AL||44|ASCII"
			nullDeathIndicatorPID := "PID||||||||||||||||||||||||||||||" + tc.in
			msg := &Message{
				Message: strings.Join([]string{msh, nullDeathIndicatorPID}, "\r"),
			}
			m, err := s.SanitizeMessage(context.Background(), msg)
			if err != nil {
				t.Fatalf("SanitizeMessage() failed with %v", err)
			}
			pid, err := m.PID()
			if err != nil {
				t.Fatalf("PID() failed with %v", err)
			}
			if got, want := pid.PatientDeathIndicator.String(), tc.want; got != want {
				t.Errorf("pid.PatientDeathIndicator got %v, want %v.", got, want)
			}
			// Check that the expected "bad value" counter was incremented the expected number of times,
			// independently of how many times PID() is invoked.
			_, err = m.PID()
			if err != nil {
				t.Fatalf("PID() failed with %v", err)
			}
			metric := "hl7_unrecognized_death_indicator_value"
			initial, final := mr.GetCounterValues(t, metric, nil)
			if got, want := final-initial, tc.wantMetricIncrement; got != want {
				t.Errorf("Metric %s is incremented by %f (initial=%f, final=%f); want %f", metric, got, initial, final, want)
			}
		})
	}
}

func TestSanitizeMessage_OBXUnitsField(t *testing.T) {
	var s MessageSanitizer

	msg := &Message{
		Message: strings.Join([]string{testMSH, testPIDAddress, testOBX4, testOBX5, testOBX6}, "\r"),
	}
	_, err := hl7.ParseMessage([]byte(msg.Message))
	if err != nil {
		t.Fatalf("ParseMessage() failed with %v", err)
	}
	m, err := s.SanitizeMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("SanitizeMessage() failed with %v", err)
	}
	obx, err := m.AllOBX()
	if err != nil {
		t.Fatalf("AllOBX() failed with %v", err)
	}

	// Units: a^b^c. Nothing should be changed.
	if obx[0].Units.Identifier.String() != "a" {
		t.Errorf("Got %v, want %v.", obx[0].Units.Identifier.String(), "a")
	}
	if obx[0].Units.Text.String() != "b" {
		t.Errorf("Got %v, want %v.", obx[0].Units.Text.String(), "b")
	}
	if obx[0].Units.NameOfCodingSystem.String() != "c" {
		t.Errorf("Got %v, want %v.", obx[0].Units.NameOfCodingSystem.String(), "c")
	}

	// Units: a. Should become a^a
	if obx[1].Units.Identifier.String() != "a" {
		t.Errorf("Got %v, want %v.", obx[1].Units.Identifier.String(), "a")
	}
	if obx[1].Units.Text.String() != "a" {
		t.Errorf("Got %v, want %v.", obx[1].Units.Text.String(), "a")
	}

	// Units: a^^c. Should become a^a^c
	if obx[2].Units.Identifier.String() != "a" {
		t.Errorf("Got %v, want %v.", obx[2].Units.Identifier.String(), "a")
	}
	if obx[2].Units.Text.String() != "a" {
		t.Errorf("Got %v, want %v.", obx[2].Units.Text.String(), "a")
	}
	if obx[2].Units.NameOfCodingSystem.String() != "c" {
		t.Errorf("Got %v, want %v.", obx[2].Units.NameOfCodingSystem.String(), "c")
	}
}

func TestSanitizeMessage_Identifiers(t *testing.T) {
	type idAndType struct {
		id  string
		typ string
	}

	tests := []struct {
		name            string
		s               *MessageSanitizer
		pid             string
		wantPatientID   *idAndType
		wantPatientList []idAndType
	}{{
		name:          "MRN in PID-2, add MRN keyword",
		s:             &MessageSanitizer{RewritePatientIdentifiers: true, AddMRN: true, RewriteIdentifierLocations: []string{"PID-2-Patient ID", "PID-3-Patient Identifier List"}},
		pid:           "PID||77652365||",
		wantPatientID: &idAndType{id: "77652365", typ: "MRN"},
	}, {
		name:          "MRN in PID-2, not rewritten if not in locations",
		s:             &MessageSanitizer{RewritePatientIdentifiers: true, AddMRN: true, RewriteIdentifierLocations: []string{"PID-3-Patient Identifier List"}},
		pid:           "PID||77652365||",
		wantPatientID: &idAndType{id: "77652365", typ: ""},
	}, {
		name:            "MRN in PID-3, add MRN keyword",
		s:               &MessageSanitizer{RewritePatientIdentifiers: true, AddMRN: true, RewriteIdentifierLocations: []string{"PID-2-Patient ID", "PID-3-Patient Identifier List"}},
		pid:             "PID|||77652365|",
		wantPatientList: []idAndType{{id: "77652365", typ: "MRN"}},
	}, {
		name:            "NHS in PID-3, NHS keyword is always added",
		s:               &MessageSanitizer{RewritePatientIdentifiers: true, AddMRN: false, RewriteIdentifierLocations: []string{"PID-2-Patient ID", "PID-3-Patient Identifier List"}},
		pid:             "PID|||4644143847|",
		wantPatientList: []idAndType{{id: "4644143847", typ: "NHSNMBR"}},
	}, {
		name:          "NHS in PID-2, add NHS keyword",
		s:             &MessageSanitizer{RewritePatientIdentifiers: true, AddMRN: false, RewriteIdentifierLocations: []string{"PID-2-Patient ID", "PID-3-Patient Identifier List"}},
		pid:           "PID||4644143847||",
		wantPatientID: &idAndType{id: "4644143847", typ: "NHSNMBR"},
	}, {
		name:            "NHS in PID-2, MRN in PID-3, Add all keywords",
		s:               &MessageSanitizer{RewritePatientIdentifiers: true, AddMRN: true, RewriteIdentifierLocations: []string{"PID-2-Patient ID", "PID-3-Patient Identifier List"}},
		pid:             "PID||4644143847|77652365|",
		wantPatientID:   &idAndType{id: "4644143847", typ: "NHSNMBR"},
		wantPatientList: []idAndType{{id: "77652365", typ: "MRN"}},
	}, {
		name:            "More than 1 component shouldn't rewrite",
		s:               &MessageSanitizer{RewritePatientIdentifiers: true, AddMRN: true, RewriteIdentifierLocations: []string{"PID-2-Patient ID", "PID-3-Patient Identifier List"}},
		pid:             "PID||4644143847^test|4644143847^test|",
		wantPatientID:   &idAndType{id: "4644143847", typ: ""},
		wantPatientList: []idAndType{{id: "4644143847", typ: ""}},
	}, {
		name:          "Multiple items in list",
		s:             &MessageSanitizer{RewritePatientIdentifiers: true, AddMRN: true, RewriteIdentifierLocations: []string{"PID-2-Patient ID", "PID-3-Patient Identifier List"}},
		pid:           "PID||4644143847|77652365~12345|",
		wantPatientID: &idAndType{id: "4644143847", typ: "NHSNMBR"},
		wantPatientList: []idAndType{
			{id: "77652365", typ: "MRN"},
			{id: "12345", typ: "MRN"},
		},
	}, {
		name:          "NHS in PID-2, do not rewrite",
		s:             &MessageSanitizer{RewritePatientIdentifiers: false},
		pid:           "PID||4644143847||",
		wantPatientID: &idAndType{id: "4644143847", typ: ""},
	}, {
		name:          "MRN in PID-2, do not rewrite",
		s:             &MessageSanitizer{RewritePatientIdentifiers: false},
		pid:           "PID||77652365||",
		wantPatientID: &idAndType{id: "77652365", typ: ""},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := &Message{
				Message: strings.Join([]string{testMSH, tc.pid}, "\r"),
			}
			m, err := tc.s.SanitizeMessage(context.Background(), msg)
			if err != nil {
				t.Fatalf("SanitizeMessage() failed with %v", err)
			}

			pid, err := m.PID()
			if err != nil {
				t.Fatalf("PID() failed with %v", err)
			}
			if tc.wantPatientID != nil {
				if pid.PatientID == nil {
					t.Fatal("PatientID is nil, want something")
				}
				if got, want := pid.PatientID.IDNumber.String(), tc.wantPatientID.id; got != want {
					t.Errorf("PatientID.ID got %q, want %q", got, want)
				}
				if got, want := pid.PatientID.IdentifierTypeCode.String(), tc.wantPatientID.typ; got != want {
					t.Errorf("PatientID.IdentifierTypeCode got %q, want %q", got, want)
				}
			}
			for i, wantIDAndType := range tc.wantPatientList {
				if i >= len(pid.PatientIdentifierList) {
					t.Fatalf("PatientIdentifierList got %d items, cannot get item %d", len(pid.PatientIdentifierList), i)
				}
				item := &pid.PatientIdentifierList[i]
				if got, want := item.IDNumber.String(), wantIDAndType.id; got != want {
					t.Errorf("PatientIdentifierList(%d).IDNumber got %q, want %q", i, got, want)
				}
				if got, want := item.IdentifierTypeCode.String(), wantIDAndType.typ; got != want {
					t.Errorf("PatientIdentifierList(%v).IdentifierTypeCode got %q, want %q", i, got, want)
				}
			}
		})
	}
}

func TestSanitizeMessage_RemoveLeadingZerosFromMRN(t *testing.T) {
	var s MessageSanitizer
	s.RemoveLeadingZerosFromMRN = true

	// No changes to the identifiers.
	msg := &Message{
		Message: strings.Join([]string{
			"MSH|^~\\&|CERNER|SF|CARELINK|AKI2|20141128001635||ORU^R01|2014112800163507740000|T|2.3|||AL||44|ASCII",
			"PID||1^^^^MRN^|2^^^^SF MRN^~3^^^MRN|4^^^^SF MRN~5^^^MRN|SMITH^JOHN^JAMES||19890901|M||||||||||5425719|"}, "\r"),
	}
	m, err := s.SanitizeMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("SanitizeMessage() failed with %v", err)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid.PatientID.IDNumber.String() != "1" {
		t.Errorf("Got %v, want %v.", pid.PatientID.IDNumber.String(), "1")
	}
	if pid.PatientID.IdentifierTypeCode.String() != "MRN" {
		t.Errorf("Got %v, want %v.", pid.PatientID.IdentifierTypeCode.String(), "MRN")
	}
	if pid.PatientID.AssigningFacility.String() != "" {
		t.Errorf("Got %v, want %v.", pid.PatientID.AssigningFacility.String(), "")
	}
	if pid.PatientIdentifierList[0].IDNumber.String() != "2" {
		t.Errorf("Got %v, want %v.", pid.PatientIdentifierList[0].IDNumber.String(), "2")
	}
	if pid.PatientIdentifierList[0].IdentifierTypeCode.String() != "SF MRN" {
		t.Errorf("Got %v, want %v.", pid.PatientIdentifierList[0].IdentifierTypeCode.String(), "SF MRN")
	}
	if pid.PatientIdentifierList[1].IDNumber.String() != "3" {
		t.Errorf("Got %v, want %v.", pid.PatientIdentifierList[1].IDNumber.String(), "3")
	}
	if pid.PatientIdentifierList[1].AssigningAuthority.String() != "MRN" {
		t.Errorf("Got %v, want %v.", pid.PatientIdentifierList[1].AssigningAuthority.String(), "MRN")
	}
	if pid.AlternatePatientIDPID[0].IDNumber.String() != "4" {
		t.Errorf("Got %v, want %v.", pid.AlternatePatientIDPID[0].IDNumber.String(), "4")
	}
	if pid.AlternatePatientIDPID[0].IdentifierTypeCode.String() != "SF MRN" {
		t.Errorf("Got %v, want %v.", pid.AlternatePatientIDPID[0].IdentifierTypeCode.String(), "SF MRN")
	}
	if pid.AlternatePatientIDPID[1].IDNumber.String() != "5" {
		t.Errorf("Got %v, want %v.", pid.AlternatePatientIDPID[1].IDNumber.String(), "5")
	}
	if pid.AlternatePatientIDPID[1].AssigningAuthority.String() != "MRN" {
		t.Errorf("Got %v, want %v.", pid.AlternatePatientIDPID[1].AssigningAuthority.String(), "MRN")
	}

	// 0's get stripped from MRNs in PID-2 (External), PID-3 (Internal) and PID-4 (Alternate).
	msg = &Message{
		Message: strings.Join([]string{
			"MSH|^~\\&|CERNER|SF|CARELINK|AKI2|20141128001635||ORU^R01|2014112800163507740000|T|2.3|||AL||44|ASCII",
			"PID||01^^^^MRN^|002^^^^MRN^~0003^^^MRN~0004^^^NHSNMBR|005^^^^SF MRN~00006^^^MRN|SMITH^JOHN^JAMES||19890901|M||||||||||5425719|"}, "\r"),
	}
	m, err = s.SanitizeMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("SanitizeMessage() failed with %v", err)
	}

	pid, err = m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid.PatientID.IDNumber.String() != "1" {
		t.Errorf("Got %v, want %v.", pid.PatientID.IDNumber.String(), "1")
	}
	if pid.PatientID.IdentifierTypeCode.String() != "MRN" {
		t.Errorf("Got %v, want %v.", pid.PatientID.IdentifierTypeCode.String(), "MRN")
	}
	if pid.PatientIdentifierList[0].IDNumber.String() != "2" {
		t.Errorf("Got %v, want %v.", pid.PatientIdentifierList[0].IDNumber.String(), "2")
	}
	if pid.PatientIdentifierList[1].IDNumber.String() != "3" {
		t.Errorf("Got %v, want %v.", pid.PatientIdentifierList[1].IDNumber.String(), "3")
	}
	// 0's from NHS numbers aren't stripped.
	if pid.PatientIdentifierList[2].IDNumber.String() != "0004" {
		t.Errorf("Got %v, want %v.", pid.PatientIdentifierList[2].IDNumber.String(), "0004")
	}
	if pid.AlternatePatientIDPID[0].IDNumber.String() != "5" {
		t.Errorf("Got %v, want %v.", pid.AlternatePatientIDPID[0].IDNumber.String(), "5")
	}
	if pid.AlternatePatientIDPID[1].IDNumber.String() != "6" {
		t.Errorf("Got %v, want %v.", pid.AlternatePatientIDPID[1].IDNumber.String(), "6")
	}

	// 0's do not get stripped if the flag is off.
	s.RemoveLeadingZerosFromMRN = false
	msg = &Message{
		Message: strings.Join([]string{
			"MSH|^~\\&|CERNER|SF|CARELINK|AKI2|20141128001635||ORU^R01|2014112800163507740000|T|2.3|||AL||44|ASCII",
			"PID||01^^^^MRN^|02^^^^MRN^~03^^^MRN~0004^^^NHS||SMITH^JOHN^JAMES||19890901|M||||||||||5425719|"}, "\r"),
	}
	m, err = s.SanitizeMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("SanitizeMessage() failed with %v", err)
	}

	pid, err = m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid.PatientID.IDNumber.String() != "01" {
		t.Errorf("Got %v, want %v.", pid.PatientID.IDNumber.String(), "01")
	}
	if pid.PatientID.IdentifierTypeCode.String() != "MRN" {
		t.Errorf("Got %v, want %v.", pid.PatientID.IdentifierTypeCode.String(), "MRN")
	}
	if pid.PatientIdentifierList[0].IDNumber.String() != "02" {
		t.Errorf("Got %v, want %v.", pid.PatientIdentifierList[0].IDNumber.String(), "02")
	}
	if pid.PatientIdentifierList[1].IDNumber.String() != "03" {
		t.Errorf("Got %v, want %v.", pid.PatientIdentifierList[1].IDNumber.String(), "03")
	}
	if pid.PatientIdentifierList[2].IDNumber.String() != "0004" {
		t.Errorf("Got %v, want %v.", pid.PatientIdentifierList[2].IDNumber.String(), "0004")
	}
}

func TestSanitizeMessage_RemoveLeadingZerosFromMRNWithoutKeyword(t *testing.T) {
	var s MessageSanitizer
	s.RemoveLeadingZerosFromMRN = true
	s.RewritePatientIdentifiers = true
	s.AddMRN = true
	s.RewriteIdentifierLocations = []string{"PID-2-Patient ID", "PID-3-Patient Identifier List"}

	// Send identifiers with leading zeros and with no "MRN" keyword.
	// The zeros need to be stripped, and the MRN keyword added.
	msg := &Message{
		Message: strings.Join([]string{
			"MSH|^~\\&|Rapidcomm|SF|OS|SF|20160927000000||ORU^R32|20160927000000|T|2.3|||AL||44|ASCII",
			"PID||0001|0002~0003|0004~0005|SMITH^JOHN^JAMES||19890901|M||||||||||5425719|"}, "\r"),
	}
	m, err := s.SanitizeMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("SanitizeMessage() failed with %v", err)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid.PatientID.IDNumber.String() != "1" {
		t.Errorf("Got %v, want %v.", pid.PatientID.IDNumber.String(), "1")
	}
	if pid.PatientID.IdentifierTypeCode.String() != "MRN" {
		t.Errorf("Got %v, want %v.", pid.PatientID.IdentifierTypeCode.String(), "MRN")
	}
	if pid.PatientIdentifierList[0].IDNumber.String() != "2" {
		t.Errorf("Got %v, want %v.", pid.PatientIdentifierList[0].IDNumber.String(), "2")
	}
	if pid.PatientIdentifierList[0].IdentifierTypeCode.String() != "MRN" {
		t.Errorf("Got %v, want %v.", pid.PatientIdentifierList[0].IdentifierTypeCode.String(), "MRN")
	}
	if pid.PatientIdentifierList[1].IDNumber.String() != "3" {
		t.Errorf("Got %v, want %v.", pid.PatientIdentifierList[1].IDNumber.String(), "3")
	}
	if pid.PatientIdentifierList[1].IdentifierTypeCode.String() != "MRN" {
		t.Errorf("Got %v, want %v.", pid.PatientIdentifierList[1].IdentifierTypeCode.String(), "MRN")
	}
	// The MRN keyword is not added for identifiers in the PID-4 field (see rewritePatientIdentifiers),
	// so the zeros aren't stripped.
	if pid.AlternatePatientIDPID[0].IDNumber.String() != "0004" {
		t.Errorf("Got %v, want %v.", pid.AlternatePatientIDPID[0].IDNumber.String(), "0004")
	}
	if pid.AlternatePatientIDPID[0].IdentifierTypeCode.String() != "" {
		t.Errorf("Got %v, want %v.", pid.AlternatePatientIDPID[0].IdentifierTypeCode.String(), "")
	}
	if pid.AlternatePatientIDPID[1].IDNumber.String() != "0005" {
		t.Errorf("Got %v, want %v.", pid.AlternatePatientIDPID[1].IDNumber.String(), "0005")
	}
	if pid.AlternatePatientIDPID[1].IdentifierTypeCode.String() != "" {
		t.Errorf("Got %v, want %v.", pid.AlternatePatientIDPID[1].IdentifierTypeCode.String(), "")
	}
}
