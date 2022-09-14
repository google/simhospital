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

package starter_test

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/doctor"
	"github.com/google/simhospital/pkg/hl7"
	"github.com/google/simhospital/pkg/hospital"
	"github.com/google/simhospital/pkg/location"
	"github.com/google/simhospital/pkg/logging"
	"github.com/google/simhospital/pkg/orderprofile"
	"github.com/google/simhospital/pkg/pathway"
	. "github.com/google/simhospital/pkg/starter"
	"github.com/google/simhospital/pkg/test"
	"github.com/google/simhospital/pkg/test/testclock"
	"github.com/google/simhospital/pkg/test/testhl7"
	"github.com/google/simhospital/pkg/test/testhospital"
	"github.com/google/simhospital/pkg/test/testlocation"
	"github.com/google/simhospital/pkg/test/testwrite"
)

const (
	// validHL7 is an arbitrary valid HL7 message.
	validHL7 = `MSH|^~\\&|CERNER|RAL|CARELINK|AKI2|20141128001635||ADT^A01|2014112800163507740000|T|2.3|||AL||44|ASCII
EVN|R01|20170329021843|||216865551019^Osman^Arthur^^^Dr^^^DRNBR^PRSNL^^^ORGDR
PID|1|843124^^^RAL MRN^MRN^|843124^^^RAL MRN^MRN^CD:5294405~1231231235^^^NHSNBR^NHSNMBR||ZZZTEST^PAUL^^^MR^^CURRENT||19000524|1|ZZZTEST^PAUL^^^MR^^PREVIOUS||100 The Street^Any Square^LONDON^^ZZ99 1BA^GBR^HOME^^||0205551234^HOME^CD:4072430~0205551234^CD:4580206||||CATHOLIC|3393034^^^RAL Encounter Num^FINNBR^||||C|||0|||||`
	// invalidHL7 is an invalid HL7 message with a problem in the MSH.6 field.
	invalidHL7    = `MSH|^~\\&|CERNER|RAL|CARELINK|AKI2|not-a-date||ADT^A01|2014112800163507740000|T|2.3|||AL||44|ASCII`
	invalidHL7Err = "errors (1): error in MSH-6-Date/Time Of Message: bad TS value: invalid length"
	testLoc       = "Ward 1"
	testLocAE     = "ED"
)

var (
	now = time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC)

	delay = &pathway.Delay{
		From: time.Second,
		To:   5 * time.Second,
	}
	arbitrarySteps = []pathway.Step{
		{Admission: &pathway.Admission{Loc: testLoc}},
		{Delay: delay},
		{Discharge: &pathway.Discharge{}},
	}
	pathway1 = pathway.Pathway{
		Persons: &pathway.Persons{
			"main-patient": pathway.Person{
				FirstName: "John",
				Surname:   "Doe",
			},
		},
		Pathway: arbitrarySteps,
	}
	// wantFailStartWithParameters contains the pathways from the Prod config that are allowed to fail to be started with a patient name or MRN.
	wantFailStartWithParameters = map[string]bool{
		"aki_scenario_with_merge":                   true,
		"aki_scenario_with_merge_aki_3_then_aki_2":  true,
		"aki_scenario_with_merge_and_complications": true,
		"explicit_merge_after_aki_result":           true,
		"explicit_merge_before_aki_result":          true,
		"merge":                                     true,
		"merge_json":                                true,
	}
	// overridesMRN contains the pathways from the Prod config that override the MRN.
	overridesMRN = map[string]bool{
		"admit_patient_with_mrn_and_all_demographics":      true,
		"admit_patient_with_mrn_and_all_demographics_json": true,
	}
)

func TestMain(m *testing.M) {
	logging.SetLogLevel(logrus.DebugLevel)
	hl7.TimezoneAndLocation("Europe/London")

	retCode := m.Run()

	os.Exit(retCode)
}

func TestNewPathwayStarter(t *testing.T) {
	ctx := context.Background()
	ps := newTestPathwayStarter(ctx, t, map[string]pathway.Pathway{}, hospital.Config{})
	defer ps.Hospital.Close()
	ts := httptest.NewServer(ps)
	defer ts.Close()

	response := getResponse(t, ts.URL)
	if got, want := response, UnavailableResponse; got != want {
		t.Errorf("GET request got %s, want %s", got, want)
	}
}

func TestServeHTTP_AllProdPathwaysCanBeStartedByName(t *testing.T) {
	ctx := context.Background()
	hl7Config, err := config.LoadHL7Config(ctx, test.MessageConfigProd)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigProd, err)
	}
	d, err := doctor.LoadDoctors(ctx, test.DoctorsConfigProd)
	if err != nil {
		t.Fatalf("LoadDoctors(%s) failed with %v", test.DoctorsConfigProd, err)
	}
	op, err := orderprofile.Load(ctx, test.OrderProfilesConfigProd, hl7Config)
	if err != nil {
		t.Fatalf("ParseOrderProfiles(%s, %+v) failed with %v", test.OrderProfilesConfigProd, hl7Config, err)
	}
	lm, err := location.NewManager(ctx, test.LocationsConfigProd)
	if err != nil {
		t.Fatalf("location.NewManager(%s) failed with %v", test.LocationsConfigProd, err)
	}
	// Get all of the pathways from the config files.
	p := &pathway.Parser{Clock: testclock.New(time.Now()), OrderProfiles: op, Doctors: d, LocationManager: lm}
	pathways, err := p.ParsePathways(ctx, test.PathwaysDirProd)
	if err != nil {
		t.Fatalf("ParsePathways(%s) failed with %v", test.PathwaysDirProd, err)
	}

	// Start each one by name.
	for k := range pathways {
		t.Run(k, func(t *testing.T) {
			ps := newTestPathwayStarter(ctx, t, pathways, hospital.Config{}).PathwayStarter
			defer ps.Hospital.Close()
			response := serveHTTP(t, ps, k)
			if want := responseStarted(k); !strings.Contains(response, want) {
				t.Errorf("ServeHTTP(%s) got response=%v, want response to contain %q", k, response, want)
			}
		})
	}
}

func TestServeHTTP_AllProdPathwaysCanBeStartedWithMRNAndPatientName(t *testing.T) {
	ctx := context.Background()
	hl7Config, err := config.LoadHL7Config(ctx, test.MessageConfigProd)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigProd, err)
	}
	d, err := doctor.LoadDoctors(ctx, test.DoctorsConfigProd)
	if err != nil {
		t.Fatalf("LoadDoctors(%s) failed with %v", test.DoctorsConfigProd, err)
	}
	op, err := orderprofile.Load(ctx, test.OrderProfilesConfigProd, hl7Config)
	if err != nil {
		t.Fatalf("ParseOrderProfiles(%s, %+v) failed with %v", test.OrderProfilesConfigProd, hl7Config, err)
	}
	lm, err := location.NewManager(ctx, test.LocationsConfigProd)
	if err != nil {
		t.Fatalf("location.NewManager(%s) failed with %v", test.LocationsConfigProd, err)
	}
	// Get all of the pathways from the config files.
	p := &pathway.Parser{Clock: testclock.New(time.Now()), OrderProfiles: op, Doctors: d, LocationManager: lm}
	pathways, err := p.ParsePathways(ctx, test.PathwaysDirProd)
	if err != nil {
		t.Fatalf("ParsePathways(%s) failed with %v", test.PathwaysDirProd, err)
	}

	tests := []struct {
		name         string
		reqParams    string
		wantResponse string
		assertsMRN   bool
	}{{
		name:         "By name",
		reqParams:    "Michael Jackson",
		wantResponse: "Michael Jackson",
	}, {
		name:         "By name and assert MRN",
		reqParams:    "Michael Jackson",
		wantResponse: "Michael Jackson (MRN 1)",
		assertsMRN:   true,
	}, {
		name:         "MRN",
		reqParams:    "1234567890",
		wantResponse: "(MRN 1234567890)",
	}}

	// Start each one by name and patient.
	for k := range pathways {
		for _, tc := range tests {
			req := fmt.Sprintf("%s: %s", k, tc.reqParams)
			t.Run(fmt.Sprintf("%s_%s", k, tc.name), func(t *testing.T) {
				ps := newTestPathwayStarter(ctx, t, pathways, hospital.Config{}).PathwayStarter
				defer ps.Hospital.Close()
				response := serveHTTP(t, ps, req)
				gotStarted := strings.Contains(response, responseStarted(k))
				wantStarted := !wantFailStartWithParameters[k]
				if gotStarted != wantStarted {
					t.Errorf("ServeHTTP(%s) got started=%t (response=%v); want started? %v", req, gotStarted, response, wantStarted)
				}
				if !gotStarted || !wantStarted {
					return
				}
				if tc.assertsMRN && overridesMRN[k] {
					// If we're here, the pathway sets the MRN, and the MRN is not explicitly set in the request.
					// We cannot assert on the entire response since we don't know the MRN that was set in the pathway.
					return
				}
				if want := tc.wantResponse; !strings.Contains(response, want) {
					t.Errorf("ServeHTTP(%s) got response=%v, want response to contain %q", req, response, want)
				}
			})
		}
	}
}

func TestServeHTTP_PathwayByName(t *testing.T) {
	ctx := context.Background()
	// Inject just one name to be chosen "randomly" so that we can make assertions later.
	names := `
RANK,1904
1,Michael
`
	surnames := `Jackson`
	fNames := testwrite.BytesToFile(t, []byte(names))
	fSurnames := testwrite.BytesToFile(t, []byte(surnames))
	df := test.DataFiles[test.Test]
	df.Boys = fNames
	df.Girls = fNames
	df.Surnames = fSurnames
	cfg := hospital.Config{DataFiles: df}

	tests := []struct {
		name         string
		pathways     map[string]pathway.Pathway
		req          string
		wantResponse string
		wantEvents   int
	}{{
		name: "No person section",
		pathways: map[string]pathway.Pathway{
			"pathway1": {
				Pathway: arbitrarySteps,
			}},
		req:          "pathway1",
		wantResponse: responseStartedWithParams("pathway1", "Michael Jackson (MRN 1)"),
		wantEvents:   3,
	}, {
		name: "Person section",
		pathways: map[string]pathway.Pathway{
			"pathway1": {
				Persons: &pathway.Persons{
					"main-patient": pathway.Person{
						FirstName: "John",
						Surname:   "Doe",
					},
				},
				Pathway: arbitrarySteps,
			}},
		req:          "pathway1",
		wantResponse: responseStartedWithParams("pathway1", "John Doe (MRN 1)"),
		wantEvents:   3,
	}, {
		name: "Patient Name with Person section",
		pathways: map[string]pathway.Pathway{
			"pathway1": {
				Persons: &pathway.Persons{
					"main-patient": pathway.Person{
						FirstName: "John",
						Surname:   "Doe",
					},
				},
				Pathway: arbitrarySteps,
			}},
		req:          "pathway1: Jane Taylor",
		wantResponse: responseStartedWithParams("pathway1", "Jane Taylor (MRN 1)"),
		wantEvents:   3,
	}, {
		name: "Patient Name with no Person section",
		pathways: map[string]pathway.Pathway{
			"pathway1": {
				Pathway: arbitrarySteps,
			}},
		req:          "pathway1: Jane Taylor",
		wantResponse: responseStartedWithParams("pathway1", "Jane Taylor (MRN 1)"),
		wantEvents:   3,
	}, {
		name: "MRN with Person section",
		pathways: map[string]pathway.Pathway{
			"pathway1": {
				Persons: &pathway.Persons{
					"main-patient": pathway.Person{
						FirstName: "John",
						Surname:   "Doe",
					},
				},
				Pathway: arbitrarySteps,
			}},
		req:          "pathway1: 1234567890",
		wantResponse: responseStartedWithParams("pathway1", "John Doe (MRN 1234567890)"),
		wantEvents:   3,
	}, {
		name: "MRN with no Person section",
		pathways: map[string]pathway.Pathway{
			"pathway1": {
				Pathway: arbitrarySteps,
			}},
		req:          "pathway1: 1234567890",
		wantResponse: responseStartedWithParams("pathway1", "Michael Jackson (MRN 1234567890)"),
		wantEvents:   3,
	}, {
		name: "Unknown pathway",
		pathways: map[string]pathway.Pathway{
			"pathway1": pathway1,
		},
		req:          "pathway2",
		wantResponse: StartErrPrefix,
	}, {
		name: "Unknown pathway and patient",
		pathways: map[string]pathway.Pathway{
			"pathway1": pathway1,
		},
		req:          "pathway2: Jane Taylor",
		wantResponse: StartErrPrefix,
	}, {
		name: "Multiple persons",
		req:  "pathway1: Jane Taylor",
		pathways: map[string]pathway.Pathway{
			"pathway1": {
				Persons: &pathway.Persons{
					// There are two persons in this section; only one is allowed.
					"main-patient": pathway.Person{
						FirstName: "John",
						Surname:   "Doe",
					},
					"second-patient": pathway.Person{},
				},
				Pathway: arbitrarySteps,
			},
		},
		wantResponse: StartErrPrefix,
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ps := newTestPathwayStarter(ctx, t, tc.pathways, cfg)
			defer ps.Hospital.Close()
			response := serveHTTP(t, ps.PathwayStarter, tc.req)

			if tc.wantResponse == "" {
				t.Error("tc.wantResponse is empty, want non empty")
			}
			if want := tc.wantResponse; !strings.Contains(response, want) {
				t.Errorf("ServeHTTP(%s) got response=%v, want response to contain %q", tc.req, response, want)
			}

			got, _ := ps.Hospital.ConsumeQueues(ctx, t)
			if want := tc.wantEvents; got != want {
				t.Errorf("ServeHTTP(%s) numEvents = %d, want %d", tc.req, got, want)
			}
		})
	}
}

func TestServeHTTP_PathwayByName_MRNRetrievesPreviousPatient(t *testing.T) {
	ctx := context.Background()
	// Inject just one name to be chosen "randomly" so that we can make assertions later.
	names := `
RANK,1904
1,Michael
`
	surnames := `Jackson`
	fNames := testwrite.BytesToFile(t, []byte(names))
	fSurnames := testwrite.BytesToFile(t, []byte(surnames))
	df := test.DataFiles[test.Test]
	df.Boys = fNames
	df.Girls = fNames
	df.Surnames = fSurnames
	cfg := hospital.Config{DataFiles: df}

	pathways := map[string]pathway.Pathway{
		"pathway1": {
			Pathway: []pathway.Step{
				{Admission: &pathway.Admission{Loc: testLoc}},
				{Delay: delay},
				{Discharge: &pathway.Discharge{}},
			},
		},
		"pathway_with_dave_grohl": {
			Persons: &pathway.Persons{
				"my-patient": {
					FirstName: "Dave",
					Surname:   "Grohl",
				},
			},
			Pathway: []pathway.Step{
				{Admission: &pathway.Admission{Loc: testLoc}},
				{Delay: delay},
				{UsePatient: &pathway.UsePatient{Patient: pathway.PatientID("my-patient")}},
				{Discharge: &pathway.Discharge{}},
			},
		},
	}
	ps := newTestPathwayStarter(ctx, t, pathways, cfg)
	defer ps.Hospital.Close()
	req1 := "pathway1: Jane Taylor-Swift"
	response := serveHTTP(t, ps.PathwayStarter, req1)

	if want := responseStartedWithParams("pathway1", "Jane", "Taylor-Swift", "(MRN 1)"); !strings.Contains(response, want) {
		t.Errorf("ServeHTTP(%s) got response=%v, want response to contain %q", req1, response, want)
	}

	// MRNs are consecutive numbers starting in 1. Starting the previous pathway should have created
	// a patient with MRN 1 and name Jane Taylor-Swift. We should reuse the same patient in a new pathway
	// by entering the MRN only.
	req2 := "pathway1: 1"
	response = serveHTTP(t, ps.PathwayStarter, req2)

	if want := responseStartedWithParams("pathway1", "Jane", "Taylor-Swift", "(MRN 1)"); !strings.Contains(response, want) {
		t.Errorf("ServeHTTP(%s) got response=%v, want response to contain %q", req2, response, want)
	}

	// If we now start the pathway without specifying a patient name it should generate a different patient.
	req3 := "pathway1"
	response = serveHTTP(t, ps.PathwayStarter, req3)

	if want := responseStartedWithParams("pathway1", "Michael", "Jackson", "(MRN 2)"); !strings.Contains(response, want) {
		t.Errorf("ServeHTTP(%s) got response=%v, want response to contain %q", req3, response, want)
	}

	// Start a different pathway that has an assigned Persons section with an existing MRN.
	// We load the patient with such MRN and ignore the Persons section.
	req4 := "pathway_with_dave_grohl: 1"
	response = serveHTTP(t, ps.PathwayStarter, req4)

	if want := responseStartedWithParams("pathway_with_dave_grohl", "Jane", "Taylor-Swift", "(MRN 1)"); !strings.Contains(response, want) {
		t.Errorf("ServeHTTP(%s) got response=%v, want response to contain %q", req4, response, want)
	}
	got, _ := ps.Hospital.ConsumeQueues(ctx, t)
	if want := 13; got != want {
		t.Errorf("ServeHTTP(%s) numEvents = %d, want %d", req4, got, want)
	}
}

func TestServeHTTP_PathwayWithPatient_PatientNotPersisted(t *testing.T) {
	ctx := context.Background()
	pathways := map[string]pathway.Pathway{
		"pathway1": pathway1,
	}
	ps := newTestPathwayStarter(ctx, t, pathways, hospital.Config{})
	defer ps.Hospital.Close()
	req := "pathway1: Jane Taylor"
	response := serveHTTP(t, ps.PathwayStarter, req)

	got, _ := ps.Hospital.ConsumeQueues(ctx, t)
	if want := 3; got != want {
		t.Errorf("ServeHTTP(%s) numEvents = %d, want %d", req, got, want)
	}
	if want := responseStartedWithParams("pathway1", "Jane", "Taylor"); !strings.Contains(response, want) {
		t.Errorf("ServeHTTP(%s) got response=%v, want response to contain %q", req, response, want)
	}

	// If we start the pathway again without specifying the patient, the patient should be the
	// original one.
	req = "pathway1"
	response = serveHTTP(t, ps.PathwayStarter, req)

	got, _ = ps.Hospital.ConsumeQueues(ctx, t)
	if want := 3; got != want {
		t.Errorf("ServeHTTP(%s) numEvents = %d, want %d", req, got, want)
	}
	if want := responseStartedWithDefaultPatient(req); !strings.Contains(response, want) {
		t.Errorf("ServeHTTP(%s) got response=%v, want response to contain %q", req, response, want)
	}
}

func TestServeHTTP_PathwayFromYml(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name         string
		yml          string
		req          string
		wantEvents   int
		wantResponse string
	}{{
		name: "Valid",
		req: fmt.Sprintf(`
persons:
 main_patient:
   first_name: John
   surname: Doe

pathway:
   - admission:
       loc: %s
   - delay:
       from: 1s
       to: 5s
   - discharge: {}
`, testLoc),
		wantEvents:   3,
		wantResponse: responseStartedWithDefaultPatient(DefaultPathwayNameFromUI),
	}, {
		name: "Invalid",
		req: fmt.Sprintf(`
persons:
 main_patient:
   first_name: John
   surname: Doe

pathway:
   - admission:
       loc: %s
   - delay:
       from: 1ss
       to: 5s
   - discharge: {}
`, testLoc),
		wantEvents:   0,
		wantResponse: StartErrPrefix,
	}, {
		name: "Pathway with name",
		req: fmt.Sprintf(`
random_pathway:
 persons:
   main_patient:
     first_name: John
     surname: Doe

 pathway:
     - admission:
         loc: %s
     - delay:
         from: 1s
         to: 5s
     - discharge: {}
`, testLoc),
		wantEvents:   3,
		wantResponse: responseStartedWithDefaultPatient("random_pathway"),
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			ps := newTestPathwayStarter(ctx, t, map[string]pathway.Pathway{}, hospital.Config{})
			defer ps.Hospital.Close()
			response := serveHTTP(t, ps.PathwayStarter, tc.req)

			got, _ := ps.Hospital.ConsumeQueues(ctx, t)
			if want := tc.wantEvents; got != want {
				t.Errorf("ServeHTTP(%s) numEvents = %d, want %d", tc.req, got, want)
			}
			if !strings.Contains(response, tc.wantResponse) {
				t.Errorf("ServeHTTP(%s) got response=%v, want response to contain %q", tc.req, response, tc.wantResponse)
			}
		})
	}
}

func TestServeHTTP_SendMessage(t *testing.T) {
	ctx := context.Background()
	// In reality, the message definition comes from a text box in the UI which has literal new lines,
	// so we use them in the tests too (as opposed to "\n" characters).
	tests := []struct {
		name          string
		msgDefinition string
		senderErr     error
		wantMsgs      int
		wantResponse  string
	}{{
		name: "One valid message",
		msgDefinition: fmt.Sprintf(`%s
`, validHL7),
		wantMsgs:     1,
		wantResponse: "Number of messages sent: 1",
	}, {
		name: "One unparsable message",
		msgDefinition: fmt.Sprintf(`%s
`, invalidHL7),
		wantMsgs:     0,
		wantResponse: fmt.Sprintf("Error sending messages: cannot parse message: %s. Number of messages successfully sent: 0", invalidHL7Err),
	}, {
		name: "Two valid messages",
		msgDefinition: fmt.Sprintf(`%s

%s`, validHL7, validHL7),
		wantMsgs:     2,
		wantResponse: "Number of messages sent: 2",
	}, {
		name: "One valid and one invalid",
		msgDefinition: fmt.Sprintf(`%s

%s`, validHL7, invalidHL7),
		wantMsgs:     1,
		wantResponse: fmt.Sprintf("Error sending messages: cannot parse message: %s. Number of messages successfully sent: 1", invalidHL7Err),
	}, {
		name: "One invalid followed by one valid",
		msgDefinition: fmt.Sprintf(`%s

%s`, invalidHL7, validHL7),
		wantMsgs:     0,
		wantResponse: fmt.Sprintf("Error sending messages: cannot parse message: %s. Number of messages successfully sent: 0", invalidHL7Err),
	}, {
		name: "Empty line at the end is ignored",
		msgDefinition: fmt.Sprintf(`%s

`, validHL7),
		wantMsgs:     1,
		wantResponse: "Number of messages sent: 1",
	}, {
		name: "Two messages with multiple lines between them",
		msgDefinition: fmt.Sprintf(`%s



%s`, validHL7, validHL7),
		wantMsgs:     2,
		wantResponse: "Number of messages sent: 2",
	}, {
		name: "Error sending the message",
		msgDefinition: fmt.Sprintf(`%s
`, validHL7),
		senderErr:    errors.New("something went wrong"),
		wantMsgs:     0,
		wantResponse: "Error sending messages: something went wrong. Number of messages successfully sent: 0",
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pathways := map[string]pathway.Pathway{}
			sender := testhl7.SenderWithError(tc.senderErr)
			ps := newTestPathwayStarter(ctx, t, pathways, hospital.Config{Sender: sender})
			defer ps.Hospital.Close()

			response := serveHTTP(t, ps.PathwayStarter, tc.msgDefinition)

			if got, want := len(sender.GetSentMessages()), tc.wantMsgs; got != want {
				t.Errorf("len(sender.messages) = %d, want %d", got, want)
			}
			if got, want := response, tc.wantResponse; !strings.Contains(got, want) {
				t.Errorf("response = %q, want response to contain %q", got, want)
			}
		})
	}
}

// serveHTTP creates a http server for the PathwayStarter and does a POST request with the given reqStr.
func serveHTTP(t *testing.T, ps *PathwayStarter, reqStr string) string {
	ts := httptest.NewServer(ps)
	defer ts.Close()
	req, err := http.NewRequest("POST", "/pathwayStarter", strings.NewReader(reqStr))
	if err != nil {
		t.Fatalf("http.NewRequest(%v, %v, %v) failed with %v", "POST", "/pathwayStarter", reqStr, err)
	}

	ps.ServeHTTP(nil, req)

	return getResponse(t, ts.URL)
}

// getResponse does a GET request and returns the response.
func getResponse(t *testing.T, url string) string {
	getReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf(`http.NewRequest("GET", %v, nil) failed with %v`, "/pathwayStarter", err)
	}
	resp, err := http.DefaultClient.Do(getReq)
	if err != nil {
		t.Fatalf("http.DefaultClient.Do(%v) failed with %v", getReq, err)
	}
	defer resp.Body.Close()

	gotBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll(%v) failed with %v", resp.Body, err)
	}
	return string(gotBody)
}

type pathwayStarter struct {
	*PathwayStarter
	Hospital *testhospital.Hospital
}

func newTestPathwayStarter(ctx context.Context, t *testing.T, pathways map[string]pathway.Pathway, cfg hospital.Config) *pathwayStarter {
	t.Helper()
	pathwayManager, err := pathway.NewDistributionManager(pathways, nil, nil)
	if err != nil {
		t.Fatalf("pathway.NewDistributionManager(%v,%v,%v) failed with %v", pathways, nil, nil, err)
	}

	cfg.PathwayManager = pathwayManager
	cfg.LocationManager = testlocation.NewLocationManager(ctx, t, testLoc, testLocAE)
	h := testhospital.WithTime(ctx, t, testhospital.Config{Config: cfg, Arguments: testhospital.Arguments}, now)
	p := &pathway.Parser{
		Clock:           testclock.New(time.Now()),
		OrderProfiles:   cfg.OrderProfiles,
		Doctors:         cfg.Doctors,
		LocationManager: cfg.LocationManager,
	}
	ps := &PathwayStarter{Hospital: h.Hospital, Parser: p, Sender: h.Sender, PathwayManager: h.PathwayManager}
	return &pathwayStarter{Hospital: h, PathwayStarter: ps}
}

func responseStarted(pathway string) string {
	return fmt.Sprintf("Pathway %s started", pathway)
}

func responseStartedWithDefaultPatient(pathway string) string {
	return responseStartedWithParams(pathway, "John", "Doe")
}

func responseStartedWithParams(pathway string, params ...string) string {
	return fmt.Sprintf("Pathway %s started with patient(s):\n%s", pathway, strings.Join(params, " "))
}
