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

package message

import (
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/hl7"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/test/testhl7"
)

const (
	defaultDoctorID        = "216865551019"
	defaultDoctorSurname   = "Osman"
	defaultDoctorFirstName = "Arthur"
	defaultDoctorPrefix    = "Dr"

	defaultProcedureCodingMethod    = "SNMCT"
	defaultProcedureCodeID          = "A01.1"
	defaultProcedureCodeDescription = "Hemispherectomy"
	defaultProcedureFunctionalType  = "A"
	defaultProcedurePriority        = "0"

	rtfContent = `{\rtf1\ansi{\fonttbl\f0\fswiss Helvetica;}\f0\pard
This is some {\b bold} text.\par
}`
)

var (
	defaultAdmissionDate         = time.Date(2017, 1, 26, 15, 24, 21, 0, time.UTC)
	defaultDischargeDate         = time.Date(2018, 2, 26, 15, 24, 21, 0, time.UTC)
	defaultTransferDate          = time.Date(2018, 4, 28, 22, 38, 13, 0, time.UTC)
	defaultExpectedAdmissionDate = time.Date(2017, 1, 26, 15, 24, 22, 0, time.UTC)
	defaultExpectedDischargeDate = time.Date(2017, 1, 26, 15, 24, 23, 0, time.UTC)
	defaultExpectedTransferDate  = time.Date(2017, 1, 26, 15, 24, 24, 0, time.UTC)
	defaultDiagnoseDate          = time.Date(2017, 1, 28, 15, 24, 24, 0, time.UTC)
	defaultProcedureDate         = time.Date(2017, 1, 29, 15, 24, 24, 0, time.UTC)
	defaultDateOfDeath           = time.Date(2020, 5, 26, 19, 28, 28, 0, time.UTC)
	defaultProcedureDateTime     = time.Date(2017, 1, 29, 15, 24, 24, 0, time.UTC)
)

func TestMain(m *testing.M) {
	hl7.TimezoneAndLocation("Europe/London")
	retCode := m.Run()
	os.Exit(retCode)
}

func TestBuildPID(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *ir.Person
		want  string
	}{{
		name: "Female",
		setup: func() *ir.Person {
			return testPersonFemale()
		},
		want: "PID|1|12529150521124992^^^SIMULATOR MRN^MRN|12529150521124992^^^SIMULATOR MRN^MRN~3333381389^^^NHSNBR^NHSNMBR||Smiths^Helen^Matilda^Junior^Miss^Dr^CURRENT||19940704133518|F|||1 Goodwill Hunting Road^Kings Cross^London^^N1C 4AG^GBR^HOME||020 7031 3000^HOME|||||||||A^White British^^^|||||||20200526202828|DECEASED",
	}, {
		name: "Missing Data",
		setup: func() *ir.Person {
			return &ir.Person{
				Prefix:    "Miss",
				FirstName: "Helen",
				Surname:   "Smiths",
				Gender:    "F",
				MRN:       "12529150521124992",
				NHS:       "3333381389",
			}
		},
		want: "PID|1|12529150521124992^^^SIMULATOR MRN^MRN|12529150521124992^^^SIMULATOR MRN^MRN~3333381389^^^NHSNBR^NHSNMBR||Smiths^Helen^^^Miss^^CURRENT|||F||||||||||||||||||||||",
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := tc.setup()
			got, err := BuildPID(p)
			if err != nil {
				t.Fatalf("BuildPID(%v) failed with %v", p, err)
			}
			if got != tc.want {
				t.Errorf("BuildPID(%v)=%v, want %v", p, got, tc.want)
			}
		})
	}
}

func TestBuildMSH(t *testing.T) {
	now := time.Date(2018, 1, 26, 15, 24, 21, 0, time.UTC)
	header := testHeader()
	mt := &Type{"ORU", "R01"}

	want := "MSH|^~\\&|CERNER|RAL1|STREAMS|RAL|20180126152421||ORU^R01|1|T|2.3|||AL||44|ASCII"
	got, err := BuildMSH(now, mt, header)
	if err != nil {
		t.Fatalf("BuildMSH(%v, %v, %v) failed with %v", now, mt, header, err)
	}
	if got != want {
		t.Errorf("BuildMSH(%v, %v, %v)=%v, want %v", now, mt, header, got, want)
	}
}

func TestBuildMSA(t *testing.T) {
	want := "MSA|AA|1"
	got, err := BuildMSA("1")
	if err != nil {
		t.Fatalf("BuildMSA(%v) failed with %v", "1", err)
	}
	if got != want {
		t.Errorf("BuildMSA(%v)=%v, want %v", "1", got, want)
	}
}

func TestBuildEVN(t *testing.T) {
	now := time.Date(2018, 1, 26, 15, 24, 21, 0, time.UTC)
	occurred := ir.NewValidTime(time.Date(2018, 1, 26, 15, 24, 23, 0, time.UTC))
	planned := ir.NewValidTime(time.Date(2018, 1, 26, 15, 24, 22, 0, time.UTC))
	operator := testDoctor()
	mt := &Type{"ORU", "R01"}

	want := "EVN|R01|20180126152421|20180126152422||216865551019^Osman^Arthur^^^Dr^^^DRNBR^PRSNL^^^ORGDR|20180126152423"
	got, err := BuildEVN(now, mt, planned, operator, occurred)
	if err != nil {
		t.Fatalf("BuildEVN(%v, %v, %v, %v, %v) failed with %v", now, mt, planned, operator, occurred, err)
	}
	if got != want {
		t.Errorf("BuildEVN(%v, %v, %v, %v, %v)=%v, want %v", now, mt, planned, operator, occurred, got, want)
	}
}

func TestBuildEVN_NoOccurredOrPlannedTime(t *testing.T) {
	now := time.Date(2018, 1, 26, 15, 24, 21, 0, time.UTC)
	operator := testDoctor()
	invalidTime := ir.NewInvalidTime()
	mt := &Type{"ORU", "R01"}

	want := "EVN|R01|20180126152421|||216865551019^Osman^Arthur^^^Dr^^^DRNBR^PRSNL^^^ORGDR|"
	got, err := BuildEVN(now, mt, invalidTime, operator, invalidTime)
	if err != nil {
		t.Fatalf("BuildEVN(%v, %v, %v, %v, %v) failed with %v", now, mt, invalidTime, operator, invalidTime, err)
	}
	if got != want {
		t.Errorf("BuildEVN(%v, %v, %v, %v, %v)=%v, want %v", now, mt, invalidTime, operator, invalidTime, got, want)
	}
}

func TestBuildORC(t *testing.T) {
	now := time.Date(2018, 1, 26, 15, 24, 21, 0, time.UTC)
	o := testOrder(now)

	want := "ORC|RE|9984058|1902082||IP||||20180126152421"
	got, err := BuildORC(o)
	if err != nil {
		t.Fatalf("BuildORC(%v) failed with %v", o, err)
	}
	if got != want {
		t.Errorf("BuildORC(%v)=%v, want %v", o, got, want)
	}
}

func TestBuildORC_NoOrderDateTime(t *testing.T) {
	o := &ir.Order{
		Placer:       "9984058",
		Filler:       "1902082",
		OrderControl: "RE",
		OrderStatus:  "IP",
	}
	want := "ORC|RE|9984058|1902082||IP||||"
	got, err := BuildORC(o)
	if err != nil {
		t.Fatalf("BuildORC(%v) failed with %v", o, err)
	}
	if got != want {
		t.Errorf("BuildORC(%v)=%v, want %v", o, got, want)
	}
}

func TestBuildOBR(t *testing.T) {
	now := time.Date(2018, 1, 26, 15, 24, 21, 0, time.UTC)

	tests := []struct {
		name    string
		setup   func() *ir.Order
		want    string
		wantErr bool
	}{{
		name: "Regular",
		setup: func() *ir.Order {
			return testOrder(now)
		},
		want: "OBR|1|9984058|1902082|lpdc-3969^UREA AND ELECTROLYTES^WinPath^^||20180126152421|||||||||||||||||||C||1",
	}, {
		name: "AllDates",
		setup: func() *ir.Order {
			o := testOrder(now)
			o.CollectedDateTime = ir.NewValidTime(time.Date(2018, 1, 26, 15, 45, 23, 0, time.UTC))
			o.ReceivedInLabDateTime = ir.NewValidTime(time.Date(2018, 1, 26, 16, 32, 55, 0, time.UTC))
			o.ReportedDateTime = ir.NewValidTime(time.Date(2018, 1, 26, 16, 51, 21, 0, time.UTC))
			return o
		},
		want: "OBR|1|9984058|1902082|lpdc-3969^UREA AND ELECTROLYTES^WinPath^^||20180126152421|20180126154523|||||||20180126163255||||||||20180126165121|||C||1",
	}, {
		name: "NoTimezone",
		setup: func() *ir.Order {
			o := testOrder(now)
			o.CollectedDateTime = ir.NewValidTime(time.Now())
			return o
		},
		wantErr: true,
	}, {
		name: "Midnight Dates",
		setup: func() *ir.Order {
			o := testOrder(now)
			o.CollectedDateTime = ir.NewMidnightTime(time.Date(2018, 1, 26, 23, 45, 23, 0, time.UTC))
			o.ReceivedInLabDateTime = ir.NewValidTime(time.Date(2018, 1, 26, 16, 32, 55, 0, time.UTC))
			o.ReportedDateTime = ir.NewValidTime(time.Date(2018, 1, 26, 16, 51, 21, 0, time.UTC))
			return o
		},
		want: "OBR|1|9984058|1902082|lpdc-3969^UREA AND ELECTROLYTES^WinPath^^||20180126152421|20180126000000|||||||20180126163255||||||||20180126165121|||C||1",
	}, {
		name: "WithOrderingProvider",
		setup: func() *ir.Order {
			o := testOrder(now)
			o.OrderingProvider = testDoctor()
			return o
		},
		want: "OBR|1|9984058|1902082|lpdc-3969^UREA AND ELECTROLYTES^WinPath^^||20180126152421||||||||||216865551019^Osman^Arthur^^^Dr^^^DRNBR^PRSNL^^^ORGDR|||||||||C||1",
	}, {
		name: "WithSpecimenSource",
		setup: func() *ir.Order {
			o := testOrder(now)
			o.SpecimenSource = "source"
			return o
		},
		want: "OBR|1|9984058|1902082|lpdc-3969^UREA AND ELECTROLYTES^WinPath^^||20180126152421|||||||||source||||||||||C||1",
	}, {
		name: "ClinicalNote",
		setup: func() *ir.Order {
			return orderWithClinicalNote(now, "content")
		},
		want: "OBR|1||document_id|document-type^document-type^^^document-title||||||||||||216865551019^Osman^Arthur^^^Dr^^^DRNBR^PRSNL^^^ORGDR||||||||MDOC|||1",
	}, {
		// We use BuildOBR, but we could use any method that uses the ToHL7Date method.
		name: "No UTC",
		setup: func() *ir.Order {
			now := time.Date(2018, 1, 26, 15, 24, 21, 0, time.Local)
			return testOrder(now)
		},
		wantErr: true,
	}, {
		name: "Escape OrderProfile",
		setup: func() *ir.Order {
			o := testOrder(now)
			o.OrderProfile = &ir.CodedElement{
				ID:            "Urea & Electrolytes",
				Text:          "Creatinine & Glucose",
				CodingSystem:  "& not escaped",
				AlternateText: "Some text with &",
			}
			return o
		},
		want: "OBR|1|9984058|1902082|Urea \\T\\ Electrolytes^Creatinine \\T\\ Glucose^& not escaped^^Some text with \\T\\||20180126152421|||||||||||||||||||C||1",
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			o := tc.setup()
			got, err := BuildOBR(o)
			if gotErr := err != nil; gotErr != tc.wantErr {
				t.Fatalf("BuildOBR(%v) failed with error %v, want error? %t", o, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("BuildOBR(%v)=%v, want %v", o, got, tc.want)
			}
		})
	}
}

func TestBuildOBR_MidnightDatesAndDifferentTimezone(t *testing.T) {
	originalTz := hl7.Timezone
	if err := hl7.TimezoneAndLocation("Europe/Madrid"); err != nil {
		t.Fatalf("hl7.TimezoneAndLocation(%q) failed with %v", "Europe/Madrid", err)
	}
	defer hl7.TimezoneAndLocation(originalTz)

	now := time.Date(2018, 1, 26, 15, 24, 21, 0, time.UTC)
	o := testOrder(now)
	o.CollectedDateTime = ir.NewMidnightTime(time.Date(2018, 1, 26, 23, 45, 23, 0, time.UTC))
	o.ReceivedInLabDateTime = ir.NewValidTime(time.Date(2018, 1, 26, 23, 32, 55, 0, time.UTC))
	o.ReportedDateTime = ir.NewValidTime(time.Date(2018, 1, 26, 23, 51, 21, 0, time.UTC))

	// Some of the dates should be set the day after (27th).
	want := "OBR|1|9984058|1902082|lpdc-3969^UREA AND ELECTROLYTES^WinPath^^||20180126162421|20180127000000|||||||20180127003255||||||||20180127005121|||C||1"
	got, err := BuildOBR(o)
	if err != nil {
		t.Fatalf("BuildOBR(%v) failed with %v", o, err)
	}
	if got != want {
		t.Errorf("BuildOBR(%v)=%v, want %v", o, got, want)
	}
}

func TestBuildOBX(t *testing.T) {
	now := time.Date(2018, 1, 26, 15, 24, 21, 0, time.UTC)

	tests := []struct {
		name  string
		setup func() *ir.Order
		want  string
	}{{
		name: "Regular",
		setup: func() *ir.Order {
			o := testOrderWithResult(now)
			o.Results[0].ObservationDateTime = ir.NewValidTime(time.Date(2018, 1, 26, 15, 45, 23, 0, time.UTC))
			return o
		},
		want: "OBX|1|NM|lpdc-2011^Creatinine^WinPath^^||700|UML|39.00 - 308.00|HIGH|||F|||20180126154523||",
	}, {
		name: "Escape Unit",
		setup: func() *ir.Order {
			o := testOrderWithResult(now)
			o.Results[0].Unit = "10^9 g/L"
			o.Results[0].ObservationDateTime = ir.NewValidTime(time.Date(2018, 1, 26, 15, 45, 23, 0, time.UTC))
			return o
		},
		want: "OBX|1|NM|lpdc-2011^Creatinine^WinPath^^||700|10\\S\\9 g/L|39.00 - 308.00|HIGH|||F|||20180126154523||",
	}, {
		name: "Escape Reference Range",
		setup: func() *ir.Order {
			o := testOrderWithResult(now)
			o.Results[0].Range = "39.00 ^ 308.00"
			o.Results[0].ObservationDateTime = ir.NewValidTime(time.Date(2018, 1, 26, 15, 45, 23, 0, time.UTC))
			return o
		},
		want: "OBX|1|NM|lpdc-2011^Creatinine^WinPath^^||700|UML|39.00 \\S\\ 308.00|HIGH|||F|||20180126154523||",
	}, {
		name: "Escape TestName",
		setup: func() *ir.Order {
			o := testOrderWithResult(now)
			o.Results[0].TestName = &ir.CodedElement{
				ID:            "Urea & Electrolytes",
				Text:          "Creatinine & Glucose",
				CodingSystem:  "& not escaped",
				AlternateText: "Some text with &",
			}
			return o
		},
		want: "OBX|1|NM|Urea \\T\\ Electrolytes^Creatinine \\T\\ Glucose^& not escaped^^Some text with \\T\\||700|UML|39.00 - 308.00|HIGH|||F|||||",
	}, {
		name: "Replace New Line",
		setup: func() *ir.Order {
			o := testOrder(now)
			o.Results = []*ir.Result{{
				TestName: &ir.CodedElement{
					ID:           "lpdc-2011",
					Text:         "Creatinine",
					CodingSystem: "WinPath",
				},
				Value:               "This is the result.\nAnd this is second line.",
				ValueType:           "TX",
				Status:              "F",
				ObservationDateTime: ir.NewValidTime(time.Date(2018, 1, 26, 15, 45, 23, 0, time.UTC)),
			}}
			return o
		},
		want: "OBX|1|TX|lpdc-2011^Creatinine^WinPath^^||This is the result.~And this is second line.||||||F|||20180126154523||",
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			o := tc.setup()
			got, err := BuildOBX(1, o.Results[0], o)
			if err != nil {
				t.Fatalf("BuildOBX(%v,%v,%v) failed with %v", 1, o.Results[0], o, err)
			}
			if got != tc.want {
				t.Errorf("BuildOBX(%v,%v,%v)=%v, want %v", 1, o.Results[0], o, got, tc.want)
			}
		})
	}
}

func TestBuildOBXForClinicalNote(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *ir.Order
		want  string
	}{{
		name: "Clinical Note",
		setup: func() *ir.Order {
			orderTime := time.Date(2018, 1, 26, 15, 24, 21, 0, time.UTC)
			return orderWithClinicalNote(orderTime, "some-content")
		},
		want: "OBX|1||ECG^ECG||^^PNG^BASE64^some-content|||||||||20180126152421||216865551019^Osman^Arthur^^^Dr^^^DRNBR^PRSNL^^^ORGDR",
	}, {
		name: "clinical note with an rtf file",
		setup: func() *ir.Order {
			orderTime := time.Date(2018, 1, 26, 15, 24, 21, 0, time.UTC)
			return orderWithClinicalNote(orderTime, rtfContent)
		},
		want: `OBX|1||ECG^ECG||^^PNG^BASE64^{\E\rtf1\E\ansi{\E\fonttbl\E\f0\E\fswiss Helvetica;}\E\f0\E\pard\.br\This is some {\E\b bold} text.\E\par\.br\}|||||||||20180126152421||216865551019^Osman^Arthur^^^Dr^^^DRNBR^PRSNL^^^ORGDR`,
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			o := tc.setup()
			got, err := BuildOBXForClinicalNote(1, 0, o.Results[0], o)
			if err != nil {
				t.Fatalf("BuildOBXForClinicalNote(%d, %d, %v, %v) failed with %v", 1, 0, o.Results[0], o, err)
			}
			if got != tc.want {
				t.Errorf("BuildOBXForClinicalNote(%d, %d, %v, %v)=%v, want %v", 1, 0, o.Results[0], o, got, tc.want)
			}
		})
	}
}

func TestBuildPV1(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() *ir.PatientInfo
		want    string
		wantErr bool
	}{{
		name: "Regular",
		setup: func() *ir.PatientInfo {
			return &ir.PatientInfo{
				Class:           "INPATIENT",
				Type:            "EMERGENCY",
				VisitID:         12341234,
				HospitalService: "180",
				Location: &ir.PatientLocation{
					Poc:          "RAL 12 West",
					Room:         "Bay01",
					Bed:          "Bed10",
					Facility:     "RAL RF",
					LocationType: "BED",
					Building:     "RFH",
					Floor:        "Floor1",
				},
				PriorLocation: &ir.PatientLocation{
					Poc:          "RAL 11 East",
					Room:         "Bay02",
					Bed:          "Bed11",
					Facility:     "RAL RF",
					LocationType: "BED",
					Building:     "RFH",
					Floor:        "Floor1",
				},
				PendingLocation: &ir.PatientLocation{
					Poc:          "RAL 12 West",
					Room:         "Bay01",
					Bed:          "Bed10",
					Facility:     "RAL RF",
					LocationType: "BED",
					Building:     "RFH",
					Floor:        "Floor1",
				},
				TemporaryLocation: &ir.PatientLocation{
					Poc:      "X-RAY",
					Facility: "RAL RF",
					Building: "RFH",
					Floor:    "Floor1",
				},
				PriorTemporaryLocation: &ir.PatientLocation{
					Poc:      "Hallway",
					Facility: "RAL RF",
					Building: "RFH",
					Floor:    "Floor1",
				},
				AttendingDoctor: testDoctor(),
				AdmissionDate:   ir.NewValidTime(time.Date(2018, 4, 28, 22, 38, 44, 0, time.UTC)),
				DischargeDate:   ir.NewValidTime(time.Date(2018, 4, 29, 21, 45, 30, 0, time.UTC)),
			}
		},
		want: "PV1|1|INPATIENT|RAL 12 West^Bay01^Bed10^RAL RF^^BED^RFH^Floor1|28b||RAL 11 East^Bay02^Bed11^RAL RF^^BED^RFH^Floor1|216865551019^Osman^Arthur^^^Dr^^^DRNBR^PRSNL^^^ORGDR|||180|X-RAY^^^RAL RF^^^RFH^Floor1|||||||EMERGENCY|12341234^^^^visitid|||||||||||||||||||||||RAL 12 West^Bay01^Bed10^RAL RF^^BED^RFH^Floor1|Hallway^^^RAL RF^^^RFH^Floor1|20180428233844|20180429224530|",
	}, {
		name: "Missing Data",
		setup: func() *ir.PatientInfo {
			return &ir.PatientInfo{
				Class:           "OUTPATIENT",
				HospitalService: "180",
				AdmissionDate:   ir.NewInvalidTime(),
			}
		},
		want: "PV1|1|OUTPATIENT||28b||||||180||||||||||||||||||||||||||||||||||||",
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			patientInfo := tc.setup()
			got, err := BuildPV1(patientInfo)
			if err != nil {
				t.Fatalf("BuildPV1(%v) failed with %v", patientInfo, err)
			}
			if got != tc.want {
				t.Errorf("BuildPV1(%v)=%v, want %v", patientInfo, got, tc.want)
			}
		})
	}
}

func TestBuildPseudoPV1(t *testing.T) {
	want := "PV1|1|N|"
	if got := BuildPseudoPV1(); got != want {
		t.Errorf("BuildPseudoPV1()=%v, want %v", got, want)
	}
}

func TestBuildPV2(t *testing.T) {
	patientInfo := &ir.PatientInfo{
		PriorPendingLocation: &ir.PatientLocation{
			Poc:          "RAL 12 West",
			Room:         "Bay01",
			Bed:          "Bed10",
			Facility:     "RAL RF",
			LocationType: "BED",
			Building:     "RFH",
			Floor:        "Floor1",
		},
		AdmitReason:               "Eye Problems",
		ExpectedAdmitDateTime:     ir.NewValidTime(time.Date(2018, 4, 28, 22, 38, 44, 0, time.UTC)),
		ExpectedDischargeDateTime: ir.NewValidTime(time.Date(2018, 4, 29, 21, 45, 30, 0, time.UTC)),
	}

	want := "PV2|RAL 12 West^Bay01^Bed10^RAL RF^^BED^RFH^Floor1||^Eye Problems|||||20180428233844|20180429224530"
	got, err := BuildPV2(patientInfo)
	if err != nil {
		t.Fatalf("BuildPV2(%v) failed with %v", patientInfo, err)
	}
	if got != want {
		t.Errorf("BuildPV2(%v)=%v, want %v", patientInfo, got, want)
	}
}

func TestBuildNK1(t *testing.T) {
	p := &ir.AssociatedParty{
		Person: &ir.Person{
			Prefix:     "Mr",
			FirstName:  "John",
			MiddleName: "George",
			Surname:    "Smiths",
			Suffix:     "Senior",
			Gender:     "M",
			Address: &ir.Address{
				FirstLine:  "5 Goodwill Hunting Road",
				City:       "London",
				PostalCode: "N1D 4AG",
				Country:    "GBR",
				Type:       "HOME",
			},
			PhoneNumber: "020 7031 4000",
			MRN:         "21124992125291505",
			NHS:         "3338933381",
		},
		Relationship: &ir.CodedElement{ID: "S", Text: "SPOUSE"},
		ContactRole:  &ir.CodedElement{ID: "F", Text: "FAMILYMEM"},
	}

	want := "NK1|3|Smiths^John^George^Senior^Mr^^CURRENT|S^SPOUSE^^^|5 Goodwill Hunting Road^^London^^N1D 4AG^GBR^HOME|020 7031 4000^HOME||F^FAMILYMEM^^^||||||||M|"
	got, err := BuildNK1(3, p)
	if err != nil {
		t.Fatalf("BuildNK1(%v, %v) failed with %v", 3, p, err)
	}
	if got != want {
		t.Errorf("BuildNK1(%v, %v)=%v, want %v", 3, p, got, want)
	}
}

func TestBuildNK1_missingData(t *testing.T) {
	p := &ir.AssociatedParty{
		Person: &ir.Person{
			FirstName:   "John",
			Surname:     "Smiths",
			Gender:      "M",
			PhoneNumber: "020 7031 4000",
			MRN:         "21124992125291505",
			NHS:         "3338933381",
		},
	}

	want := "NK1|3|Smiths^John^^^^^CURRENT|||020 7031 4000^HOME||||||||||M|"
	got, err := BuildNK1(3, p)
	if err != nil {
		t.Fatalf("BuildNK1(%v, %v) failed with %v", 3, p, err)
	}
	if got != want {
		t.Errorf("BuildNK1(%v, %v)=%v, want %v", 3, p, got, want)
	}
}

func TestBuildAL1(t *testing.T) {
	tests := []struct {
		name        string
		allergy     *ir.Allergy
		expectedMsg string
	}{
		{
			"Moderate food allergy, valid IdentificationDate",
			&ir.Allergy{
				Type:                   "FA",
				Description:            ir.CodedElement{ID: "E", Text: "egg-containing compound", CodingSystem: "ZAL"},
				Severity:               "MO",
				Reaction:               "Skin rash",
				IdentificationDateTime: ir.NewValidTime(time.Date(2018, 4, 28, 22, 38, 44, 0, time.UTC)),
			},
			"AL1|2|FA|E^egg-containing compound^ZAL^^|MO|Skin rash|20180428233844",
		},
		{
			"Severe drug allergy, missing IdentificationDate",
			&ir.Allergy{
				Type:        "DA",
				Description: ir.CodedElement{ID: "E", Text: "egg-containing drug", CodingSystem: "ZAL"},
				Severity:    "SV",
				Reaction:    "Rash",
			},
			"AL1|2|DA|E^egg-containing drug^ZAL^^|SV|Rash|",
		},
		{
			"Mild miscellaneous allergy, invalid IdentificationDate",
			&ir.Allergy{
				Type:                   "MA",
				Description:            ir.CodedElement{ID: "E", Text: "eggshell-containing liquid", CodingSystem: "ZAL"},
				Severity:               "MI",
				Reaction:               "Skin boils",
				IdentificationDateTime: ir.NewInvalidTime(),
			},
			"AL1|2|MA|E^eggshell-containing liquid^ZAL^^|MI|Skin boils|",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildAL1(2, tt.allergy)
			if err != nil {
				t.Fatalf("BuildAL1(%v, %v) failed with %v", 2, tt.allergy, err)
			}
			if want := tt.expectedMsg; got != want {
				t.Errorf("BuildAL1(%v, %v) = %v, want %v", 2, tt.allergy, got, want)
			}
		})
	}
}

func TestBuildNTE(t *testing.T) {
	want := "NTE|2||Test note|"
	got, err := BuildNTE(2, "Test note")
	if err != nil {
		t.Fatalf("BuildNTE(%v, %v) failed with %v", 2, "Test note", err)
	}
	if got != want {
		t.Errorf("BuildNTE(%v, %v)=%v, want %v", 2, "Test note", got, want)
	}
}

func TestBuildMRG_OneMRN(t *testing.T) {
	mrns := []string{"123"}

	want := "MRG|123^^^SIMULATOR MRN^MRN|"
	got, err := BuildMRG(mrns)
	if err != nil {
		t.Fatalf("BuildMRG(%v) failed with %v", mrns, err)
	}
	if got != want {
		t.Errorf("BuildMRG(%v)=%v, want %v", mrns, got, want)
	}
}

func TestBuildMRG_MultipleMRNs(t *testing.T) {
	mrns := []string{"123", "456"}

	want := "MRG|123^^^SIMULATOR MRN^MRN~456^^^SIMULATOR MRN^MRN|"
	got, err := BuildMRG(mrns)
	if err != nil {
		t.Fatalf("BuildMRG(%v) failed with %v", mrns, err)
	}
	if got != want {
		t.Errorf("BuildMRG(%v)=%v, want %v", mrns, got, want)
	}
}

func TestBuildDG1(t *testing.T) {
	diagnose := testDiagnosis()
	want := "DG1|2|SNMCT|A01.0^Typhoid fever^^^|Typhoid fever|20170128152424|Admitting|||||||||0|216865551019^Osman^Arthur^^^Dr^^^DRNBR^PRSNL^^^ORGDR"
	got, err := BuildDG1(2, diagnose)
	if err != nil {
		t.Fatalf("BuildDG1(%v, %v) failed with %v", 2, diagnose, err)
	}
	if got != want {
		t.Errorf("BuildDG1(%v, %v)=%v, want %v", 2, diagnose, got, want)
	}
}

func TestBuildPR1(t *testing.T) {
	procedure := testProcedure()
	want := "PR1|2|SNMCT|A01.1^Hemispherectomy^^^|Hemispherectomy|20170129152424|A||||||216865551019^Osman^Arthur^^^Dr^^^DRNBR^PRSNL^^^ORGDR||0||"
	got, err := BuildPR1(2, procedure)
	if err != nil {
		t.Fatalf("BuildPR1(%v, %v) failed with %v", 2, procedure, err)
	}
	if got != want {
		t.Errorf("BuildPR1(%v, %v)=%v, want %v", 2, procedure, got, want)
	}
}

func TestBuildTXA(t *testing.T) {
	d := document()
	p := &ir.PatientInfo{
		AttendingDoctor: &ir.Doctor{
			ID:        "743857BT34",
			Surname:   "Davis",
			FirstName: "Olive",
		},
	}
	want := "TXA|1|DS||20190615091340|743857BT34^Davis^Olive^^^^^^DRNBR^PRSNL^^^ORGDR|||20191104081340||||9298345CE5003|||||DO||||||"
	got, err := BuildTXA(p, d)
	if err != nil {
		t.Fatalf("BuildTXA(%v, %v) failed with %v", p, d, err)
	}
	if got != want {
		t.Errorf("BuildTXA(%v, %v) = %v, want %v", p, d, got, want)
	}
}

func TestBuildOBXForMDM(t *testing.T) {
	observationIdentifier := &ir.CodedElement{
		ID:           "Established Patient 15",
		Text:         "Established Patient 15",
		CodingSystem: "Simulation",
	}
	contentLine := "Name : SULLY, J K (65yo, F) ID# 47Q66Q585"
	want := "OBX|2|TX|Established Patient 15^Established Patient 15^Simulation^^|1|Name : SULLY, J K (65yo, F) ID# 47Q66Q585||||||F||||||"
	got, err := BuildOBXForMDM(2, observationIdentifier, contentLine)
	if err != nil {
		t.Fatalf("BuildOBXForMDM(%v, %v, %v) failed with %v", 2, observationIdentifier, contentLine, err)
	}
	if got != want {
		t.Errorf("BuildOBXForMDM(%v, %v, %v) = %v, want %v", 2, observationIdentifier, contentLine, got, want)
	}
}

func TestPD1(t *testing.T) {
	tests := []struct {
		name            string
		primaryFacility *ir.PrimaryFacility
		expected        string
	}{
		{
			"Nil primary facility",
			nil,
			"PD1||||",
		},
		{
			"Populated Primary Facility",
			&ir.PrimaryFacility{
				Organization: "ORG",
				ID:           "12345",
			},
			"PD1|||ORG^^12345|",
		},
		{
			"All empty",
			&ir.PrimaryFacility{
				Organization: "",
				ID:           "",
			},
			"PD1|||^^|",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			patientInfo := testPatientInfo()
			patientInfo.PrimaryFacility = test.primaryFacility

			pd1, err := BuildPD1(patientInfo)
			if err != nil {
				t.Fatalf("BuildPD1(%v) failed with %v", patientInfo, err)
			}
			if got, want := pd1, test.expected; got != want {
				t.Errorf("BuildPD1(%v)=%v, want %v", patientInfo, got, want)
			}
		})
	}
}

func TestBuildPathologyORRO02(t *testing.T) {
	now := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	order := testOrder(now)
	header := testHeader()

	orrO02, err := BuildPathologyORRO02(header, patientInfo, order, msgTime)
	if err != nil {
		t.Fatalf("BuildPathologyORRO02(%v, %v, %v, %v) failed with %v", header, patientInfo, order, msgTime, err)
	}
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(orrO02.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", orrO02.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ORR"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "O02"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	msa, err := m.MSA()
	if err != nil {
		t.Fatalf("MSA() failed with %v", err)
	}
	if msa == nil {
		t.Fatal("MSA() got nil MSA segment, want non nil")
	}
	if got, want := msa.AcknowledgmentCode.String(), "AA"; got != want {
		t.Errorf("msa.AcknowledgementCode.String()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	orc, err := m.ORC()
	if err != nil {
		t.Fatalf("ORC() failed with %v", err)
	}
	if orc == nil {
		t.Error("ORC() got nil ORC segment, want non nil")
	}
}

func TestBuildAdmissionADTA01(t *testing.T) {
	admissionTime := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adt, err := BuildAdmissionADTA01(header, patientInfo, admissionTime, msgTime)
	if err != nil {
		t.Fatalf("BuildAdmissionADTA01(%v, %v, %v, %v) failed with %v", header, patientInfo, admissionTime, msgTime, err)
	}

	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A01"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A01"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), admissionTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}

	al1, err := m.AL1()
	if err != nil {
		t.Fatalf("AL1() failed with %v", err)
	}
	if al1 == nil {
		t.Error("AL1() got nil AL1 segment, want non nil")
	}

	nk1, err := m.NK1()
	if err != nil {
		t.Fatalf("NK1() failed with %v", err)
	}
	if nk1 == nil {
		t.Error("NK1() got nil NK1 segment, want non nil")
	}

	pv2, err := m.PV2()
	if err != nil {
		t.Errorf("PV1() failed with %v", err)
	}
	if pv2 == nil {
		t.Errorf("PV2() got <nil> PV2 segment, want non nil")
	}
	want := &hl7.CE{Text: hl7.NewST("Eye problems")}
	if diff := cmp.Diff(want, pv2.AdmitReason); diff != "" {
		t.Errorf("pv2.AdmitReason mismatch (-want, +got)=\n%s", diff)
	}
}

func TestBuildTransferADTA02(t *testing.T) {
	transferTime := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adt, err := BuildTransferADTA02(header, patientInfo, transferTime, msgTime)
	if err != nil {
		t.Fatalf("BuildTransferADTA02(%v, %v, %v, %v) failed with %v", header, patientInfo, transferTime, msgTime, err)
	}
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A02"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A02"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), transferTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Fatal("PV1() got nil PV1 segment, want non nil")
	}
	if pv1.PriorPatientLocation == nil {
		t.Fatal("pv1.PriorPatientLocation is <nil>, want non nil")
	}
	if pv1.AssignedPatientLocation == nil {
		t.Fatal("pv1.AssignedPatientLocation is <nil>, want non nil")
	}
	if got, want := pv1.AssignedPatientLocation.PointOfCare.String(), "RAL 12 West"; got != want {
		t.Errorf("pv1.AssignedPatientLocation.PointOfCare.String()=%v, want %v", got, want)
	}
	if got, want := pv1.PriorPatientLocation.PointOfCare.String(), "RAL 12 East"; got != want {
		t.Errorf("pv1.PriorPatientLocation.PointOfCare.String()=%v, want %v", got, want)
	}
}

func TestBuildDischargeADTA03(t *testing.T) {
	dischargeTime := time.Date(2018, 4, 28, 22, 38, 44, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 44, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adtA03, err := BuildDischargeADTA03(header, patientInfo, dischargeTime, msgTime)
	if err != nil {
		t.Fatalf("BuildDischargeADTA03(%v, %v, %v, %v) failed with %v", header, patientInfo, dischargeTime, msgTime, err)
	}
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adtA03.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adtA03.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A03"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A03"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), dischargeTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}
}

func TestBuildRegistrationADTA04(t *testing.T) {
	registrationTime := time.Date(2018, 4, 28, 22, 38, 44, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 44, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adtA04, err := BuildRegistrationADTA04(header, patientInfo, registrationTime, msgTime)
	if err != nil {
		t.Fatalf("BuildRegistrationADTA04(%v, %v, %v, %v) failed with %v", header, patientInfo, registrationTime, msgTime, err)
	}
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adtA04.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adtA04.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A04"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A04"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), registrationTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}

	al1, err := m.AL1()
	if err != nil {
		t.Fatalf("AL1() failed with %v", err)
	}
	if al1 == nil {
		t.Error("AL1() got nil AL1 segment, want non nil")
	}
}

func TestBuildTrackDepartureADTA09(t *testing.T) {
	departureTime := time.Date(2018, 7, 2, 15, 36, 30, 0, time.UTC)
	msgTime := time.Date(2018, 7, 2, 15, 37, 30, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adtA09, err := BuildTrackDepartureADTA09(header, patientInfo, departureTime, msgTime)
	if err != nil {
		t.Fatalf("BuildTrackDepartureADTA09(%v, %v, %v, %v) failed with %v", header, patientInfo, departureTime, msgTime, err)
	}
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adtA09.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adtA09.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A09"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A09"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), departureTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}
}

func TestBuildTrackArrivalADTA10(t *testing.T) {
	arrivalTime := time.Date(2018, 7, 3, 11, 26, 30, 0, time.UTC)
	msgTime := time.Date(2018, 7, 3, 12, 27, 30, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adtA10, err := BuildTrackArrivalADTA10(header, patientInfo, arrivalTime, msgTime)
	if err != nil {
		t.Fatalf("BuildTrackArrivalADTA10(%v, %v, %v, %v) failed with %v", header, patientInfo, arrivalTime, msgTime, err)
	}
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adtA10.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adtA10.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A10"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A10"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), arrivalTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}
}

func TestBuildCancelVisitADTA11(t *testing.T) {
	cancelVisitTime := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adt, err := BuildCancelVisitADTA11(header, patientInfo, cancelVisitTime, msgTime)
	if err != nil {
		t.Fatalf("BuildCancelVisitADTA11(%v, %v, %v, %v) failed with %v", header, patientInfo, cancelVisitTime, msgTime, err)
	}
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A11"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A11"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), cancelVisitTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}
	if got, want := evn.EventOccurred.Time.Second(), defaultAdmissionDate.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.EventOccurred.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Fatal("PV1() got nil PV1 segment, want non nil")
	}
	if pv1.AssignedPatientLocation == nil {
		t.Fatal("pv1.AssignedPatientLocation is <nil>, want non nil")
	}
	if got, want := pv1.PriorPatientLocation.PointOfCare.String(), "RAL 12 East"; got != want {
		t.Errorf("pv1.PriorPatientLocation.PointOfCare.String()=%v, want %v", got, want)
	}
}

func TestBuildCancelTransferADTA12(t *testing.T) {
	cancelTransferTime := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adt, err := BuildCancelTransferADTA12(header, patientInfo, cancelTransferTime, msgTime)
	if err != nil {
		t.Fatalf("BuildCancelTransferADTA12(%v, %v, %v, %v) failed with %v", header, patientInfo, cancelTransferTime, msgTime, err)
	}
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A12"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A12"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), cancelTransferTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}
	if got, want := evn.EventOccurred.Time.Second(), defaultTransferDate.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.EventOccurred.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Fatal("PV1() got nil PV1 segment, want non nil")
	}
	if pv1.AssignedPatientLocation == nil {
		t.Fatal("pv1.AssignedPatientLocation is <nil>, want non nil")
	}
}

func TestBuildCancelDischargeADTA13(t *testing.T) {
	cancelDischargeTime := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adt, err := BuildCancelDischargeADTA13(header, patientInfo, cancelDischargeTime, msgTime)
	if err != nil {
		t.Fatalf("BuildCancelDischargeADTA13(%v, %v, %v, %v) failed with %v", header, patientInfo, cancelDischargeTime, msgTime, err)
	}
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A13"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A13"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), cancelDischargeTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}
	if got, want := evn.EventOccurred.Time.Second(), defaultDischargeDate.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.REventOccurred.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Fatal("PV1() got nil PV1 segment, want non nil")
	}
	if pv1.AssignedPatientLocation == nil {
		t.Fatal("pv1.AssignedPatientLocation is <nil>, want non nil")
	}
	if pv1.PriorPatientLocation == nil {
		t.Fatal("pv1.PriorPatientLocation is <nil>, want non nil")
	}
	if got, want := pv1.PriorPatientLocation.PointOfCare.String(), "RAL 12 East"; got != want {
		t.Errorf("pv1.PriorPatientLocation.PointOfCare.String()=%v, want %v", got, want)
	}
}

func TestBuildPendingAdmissionA14(t *testing.T) {
	pendingAdmissionTime := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adt, err := BuildPendingAdmissionADTA14(header, patientInfo, pendingAdmissionTime, msgTime)
	if err != nil {
		t.Fatalf("BuildPendingAdmissionADTA14(%v, %v, %v, %v) failed with %v", header, patientInfo, pendingAdmissionTime, msgTime, err)
	}
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A14"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A14"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), pendingAdmissionTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}
	if got, want := evn.DateTimePlannedEvent.Time.Second(), defaultExpectedAdmissionDate.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.DateTimePlannedEvent.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}

	pv2, err := m.PV2()
	if err != nil {
		t.Fatalf("PV2() failed with %v", err)
	}
	if pv2 == nil {
		t.Error("PV2() got nil PV2 segment, want non nil")
	}
}

func TestBuildPreAdmitA05(t *testing.T) {
	preAdmitTime := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adt, err := BuildPreAdmitADTA05(header, patientInfo, preAdmitTime, msgTime)
	if err != nil {
		t.Fatalf("BuildPreAdmitADTA05(%v, %v, %v, %v) failed with %v", header, patientInfo, preAdmitTime, msgTime, err)
	}
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A05"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A05"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), preAdmitTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}
	if got, want := evn.DateTimePlannedEvent.Time.Second(), defaultExpectedAdmissionDate.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.DateTimePlannedEvent.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}

	pv2, err := m.PV2()
	if err != nil {
		t.Fatalf("PV2() failed with %v", err)
	}
	if pv2 == nil {
		t.Error("PV2() got nil PV2 segment, want non nil")
	}
}

func TestBuildPendingDischargeA16(t *testing.T) {
	pendingDischargeTime := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adt, err := BuildPendingDischargeADTA16(header, patientInfo, pendingDischargeTime, msgTime)
	if err != nil {
		t.Fatalf("BuildPendingDischargeADTA16(%v, %v, %v, %v) failed with %v", header, patientInfo, pendingDischargeTime, msgTime, err)
	}

	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A16"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A16"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), pendingDischargeTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}
	if got, want := evn.DateTimePlannedEvent.Time.Second(), defaultExpectedDischargeDate.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.DateTimePlannedEvent.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}

	pv2, err := m.PV2()
	if err != nil {
		t.Fatalf("PV2() failed with %v", err)
	}
	if pv2 == nil {
		t.Error("PV2() got nil PV2 segment, want non nil")
	}
}

func TestBuildPendingTransferADTA15(t *testing.T) {
	pendingTransferTime := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adt, err := BuildPendingTransferADTA15(header, patientInfo, pendingTransferTime, msgTime)
	if err != nil {
		t.Fatalf("BuildPendingTransferADTA15(%v, %v, %v, %v) failed with %v", header, patientInfo, pendingTransferTime, msgTime, err)
	}
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}

	if got, want := msh.MessageType.TriggerEvent.String(), "A15"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A15"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), pendingTransferTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}
	if got, want := evn.DateTimePlannedEvent.Time.Second(), defaultExpectedTransferDate.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.EDateTimePlannedEvent.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}
}

func TestBuildDeleteVisitA23(t *testing.T) {
	deleteTime := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adt, err := BuildDeleteVisitADTA23(header, patientInfo, deleteTime, msgTime)
	if err != nil {
		t.Fatalf("BuildDeleteVisitADTA23(%v, %v, %v, %v) failed with %v", header, patientInfo, deleteTime, msgTime, err)
	}
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A23"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A23"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), deleteTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}
}

func TestBuildCancelPendingDischargeA25(t *testing.T) {
	cancelPendingDischargeTime := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adt, err := BuildCancelPendingDischargeADTA25(header, patientInfo, cancelPendingDischargeTime, msgTime)
	if err != nil {
		t.Fatalf("BuildCancelPendingDischargeADTA25(%v, %v, %v, %v) failed with %v", header, patientInfo, cancelPendingDischargeTime, msgTime, err)
	}

	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A25"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A25"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), cancelPendingDischargeTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}
	if got, want := evn.EventOccurred.Time.Second(), defaultExpectedDischargeDate.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.EventOccurred.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}

	pv2, err := m.PV2()
	if err != nil {
		t.Fatalf("PV2() failed with %v", err)
	}
	if pv2 == nil {
		t.Error("PV2() got nil PV2 segment, want non nil")
	}
}

func TestBuildCancelPendingTransferADTA26(t *testing.T) {
	cancelPendingTransferTime := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adt, err := BuildCancelPendingTransferADTA26(header, patientInfo, cancelPendingTransferTime, msgTime)
	if err != nil {
		t.Fatalf("BuildCancelPendingTransferADTA26(%v, %v, %v, %v) failed with %v", header, patientInfo, cancelPendingTransferTime, msgTime, err)
	}
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A26"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A26"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), cancelPendingTransferTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}
	if got, want := evn.EventOccurred.Time.Second(), defaultExpectedTransferDate.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.EventOccurred.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}

	pv2, err := m.PV2()
	if err != nil {
		t.Fatalf("PV2() failed with %v", err)
	}
	if pv2 == nil {
		t.Error("PV2() got nil PV2 segment, want non nil")
	}
}

func TestBuildCancelPendingAdmitADTA27(t *testing.T) {
	cancelPendingAdmitTime := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adt, err := BuildCancelPendingAdmitADTA27(header, patientInfo, cancelPendingAdmitTime, msgTime)
	if err != nil {
		t.Errorf("BuildCancelPendingAdmitADTA27(%v, %v, %v, %v) failed with %v", header, patientInfo, cancelPendingAdmitTime, msgTime, err)
	}

	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A27"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A27"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), cancelPendingAdmitTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}
	if got, want := evn.EventOccurred.Time.Second(), defaultExpectedAdmissionDate.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.EventOccurred.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}

	pv2, err := m.PV2()
	if err != nil {
		t.Fatalf("PV2() failed with %v", err)
	}
	if pv2 == nil {
		t.Error("PV2() got nil PV2 segment, want non nil")
	}
}

func TestBuildUpdatePatientA08(t *testing.T) {
	now := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adt, err := BuildUpdatePatientADTA08(header, patientInfo, now, msgTime)
	if err != nil {
		t.Errorf("BuildUpdatePatientADTA08(%v, %v, %v, %v) failed with %v", header, patientInfo, now, msgTime, err)
	}

	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A08"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A08"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), now.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}
	if got, want := pv1.PatientClass.String(), "N"; got != want {
		t.Errorf("pv1.PatientClass.String()=%v, want %v", got, want)
	}

	dg1, err := m.DG1()
	if err != nil {
		t.Fatalf("DG1() failed with %v", err)
	}
	if dg1 == nil {
		t.Error("DG1() got nil DG1 segment, want non nil")
	}

	pr1, err := m.PR1()
	if err != nil {
		t.Fatalf("PR1() failed with %v", err)
	}
	wantPR1 := &hl7.PR1{
		SetIDPR1:              &hl7.SI{Valid: true},
		ProcedureCodingMethod: hl7.NewIS(defaultProcedureCodingMethod),
		ProcedureCode: &hl7.CE{
			Identifier: hl7.NewST(defaultProcedureCodeID),
			Text:       hl7.NewST(defaultProcedureCodeDescription),
		},
		ProcedureDescription:    hl7.NewST(defaultProcedureCodeDescription),
		ProcedureDateTime:       &hl7.TS{Time: defaultProcedureDateTime, Precision: hl7.SecondPrecision},
		ProcedureFunctionalType: hl7.NewIS(defaultProcedureFunctionalType),
		ProcedurePriority:       hl7.NewID(defaultProcedurePriority),
		ProcedurePractitioner: []hl7.XCN{{
			IDNumber:           hl7.NewST(defaultDoctorID),
			GivenName:          hl7.NewST(defaultDoctorFirstName),
			PrefixEGDR:         hl7.NewST(defaultDoctorPrefix),
			FamilyName:         &hl7.FN{Surname: hl7.NewST(defaultDoctorSurname)},
			AssigningAuthority: &hl7.HD{NamespaceID: hl7.NewIS("DRNBR")},
			NameTypeCode:       hl7.NewID("PRSNL"),
			IdentifierTypeCode: hl7.NewID("ORGDR"),
		}},
	}
	if diff := cmp.Diff(wantPR1, pr1); diff != "" {
		t.Errorf("PR1() got diff (-want, +got):\n%v", diff)
	}

	al1, err := m.AL1()
	if err != nil {
		t.Fatalf("AL1() failed with %v", err)
	}
	if al1 == nil {
		t.Error("AL1() got nil AL1 segment, want non nil")
	}
}

func TestBuildAddPersonADTA28(t *testing.T) {
	now := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adt, err := BuildAddPersonADTA28(header, patientInfo, now, msgTime)
	if err != nil {
		t.Errorf("BuildAddPersonADTA28(%v, %v, %v, %v) failed with %v", header, patientInfo, now, msgTime, err)
	}

	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A28"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A28"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), now.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}
	if got, want := pv1.PatientClass.String(), "N"; got != want {
		t.Errorf("pv1.PatientClass.String()=%v, want %v", got, want)
	}

	al1, err := m.AL1()
	if err != nil {
		t.Fatalf("AL1() failed with %v", err)
	}
	if al1 == nil {
		t.Error("AL1() got nil AL1 segment, want non nil")
	}
}

func TestBuildUpdatePersonADTA31(t *testing.T) {
	now := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	adt, err := BuildUpdatePersonADTA31(header, patientInfo, now, msgTime)
	if err != nil {
		t.Fatalf("BuildUpdatePersonADTA31(%v, %v, %v, %v) failed with %v", header, patientInfo, now, msgTime, err)
	}
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A31"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A31"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), now.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}
	if got, want := pv1.PatientClass.String(), "N"; got != want {
		t.Errorf("pv1.PatientClass.String()=%v, want %v", got, want)
	}

	dg1, err := m.DG1()
	if err != nil {
		t.Fatalf("DG1() failed with %v", err)
	}
	if dg1 == nil {
		t.Error("DG1() got nil DG1 segment, want non nil")
	}

	pr1, err := m.PR1()
	if err != nil {
		t.Fatalf("PR1() failed with %v", err)
	}
	wantPR1 := &hl7.PR1{
		SetIDPR1:              &hl7.SI{Valid: true},
		ProcedureCodingMethod: hl7.NewIS(defaultProcedureCodingMethod),
		ProcedureCode: &hl7.CE{
			Identifier: hl7.NewST(defaultProcedureCodeID),
			Text:       hl7.NewST(defaultProcedureCodeDescription),
		},
		ProcedureDescription:    hl7.NewST(defaultProcedureCodeDescription),
		ProcedureDateTime:       &hl7.TS{Time: defaultProcedureDateTime, Precision: hl7.SecondPrecision},
		ProcedureFunctionalType: hl7.NewIS(defaultProcedureFunctionalType),
		ProcedurePriority:       hl7.NewID(defaultProcedurePriority),
		ProcedurePractitioner: []hl7.XCN{{
			IDNumber:           hl7.NewST(defaultDoctorID),
			GivenName:          hl7.NewST(defaultDoctorFirstName),
			PrefixEGDR:         hl7.NewST(defaultDoctorPrefix),
			FamilyName:         &hl7.FN{Surname: hl7.NewST(defaultDoctorSurname)},
			AssigningAuthority: &hl7.HD{NamespaceID: hl7.NewIS("DRNBR")},
			NameTypeCode:       hl7.NewID("PRSNL"),
			IdentifierTypeCode: hl7.NewID("ORGDR"),
		}},
	}
	if diff := cmp.Diff(wantPR1, pr1); diff != "" {
		t.Errorf("PR1() got diff (-want, +got):\n %v", diff)
	}

	al1, err := m.AL1()
	if err != nil {
		t.Fatalf("AL1() failed with %v", err)
	}
	if al1 == nil {
		t.Error("AL1() got nil AL1 segment, want non nil")
	}
}

func TestBuildResultORU(t *testing.T) {
	eventTime := time.Date(2018, 4, 28, 22, 38, 44, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 44, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	tests := []struct {
		name             string
		order            *ir.Order
		f                func(*HeaderInfo, *ir.PatientInfo, *ir.Order, time.Time) (*HL7Message, error)
		numNotes         int
		wantTriggerEvent string
		wantOBXSetIDs    []string
	}{{
		name:             "ORU^R01 for test results",
		order:            testOrderWithResult(eventTime),
		f:                BuildResultORUR01,
		numNotes:         2,
		wantTriggerEvent: "R01",
		wantOBXSetIDs:    []string{"1", "2"},
	}, {
		name:             "ORU^R01 for Doctor Notes",
		order:            orderWithClinicalNote(eventTime, "some-content"),
		f:                BuildResultORUR01,
		numNotes:         0,
		wantTriggerEvent: "R01",
		wantOBXSetIDs:    []string{"1", "2"},
	}, {
		name:             "ORU^R32",
		order:            testOrderWithResult(eventTime),
		f:                BuildResultORUR32,
		numNotes:         2,
		wantTriggerEvent: "R32",
		wantOBXSetIDs:    []string{"1", "2"},
	}, {
		name:             "ORU^R03",
		order:            testOrderWithResult(eventTime),
		f:                BuildResultORUR03,
		numNotes:         2,
		wantTriggerEvent: "R03",
		wantOBXSetIDs:    []string{"1", "2"},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oruR01, err := tt.f(header, patientInfo, tt.order, msgTime)
			if err != nil {
				t.Fatalf("BuildResultORU(%v, %v, %v, %v) failed with %v", header, patientInfo, tt.order, msgTime, err)
			}
			mo := hl7.NewParseMessageOptions()
			mo.TimezoneLoc = time.UTC
			m, err := hl7.ParseMessageWithOptions([]byte(oruR01.Message), mo)
			if err != nil {
				t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", oruR01.Message, mo, err)
			}

			msh, err := m.MSH()
			if err != nil {
				t.Fatalf("MSH() failed with %v", err)
			}
			if msh == nil {
				t.Fatal("MSH() got nil MSH segment, want non nil")
			}
			if got, want := msh.MessageType.MessageCode.String(), "ORU"; got != want {
				t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
			}
			if got, want := msh.MessageType.TriggerEvent.String(), tt.wantTriggerEvent; got != want {
				t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
			}
			if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
				t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
			}

			evn, err := m.EVN()
			if err != nil {
				t.Fatalf("EVN() failed with %v", err)
			}
			if evn != nil {
				t.Fatal("evn is not <nil>.")
			}

			pid, err := m.PID()
			if err != nil {
				t.Fatalf("PID() failed with %v", err)
			}
			if pid == nil {
				t.Error("PID() got nil PID segment, want non nil")
			}

			pv1, err := m.PV1()
			if err != nil {
				t.Fatalf("PV1() failed with %v", err)
			}
			if pv1 == nil {
				t.Error("PV1() got nil PV1 segment, want non nil")
			}

			orc, err := m.ORC()
			if err != nil {
				t.Fatalf("ORC() failed with %v", err)
			}
			if orc == nil {
				t.Error("ORC() got nil ORC segment, want non nil")
			}

			obr, err := m.OBR()
			if err != nil {
				t.Fatalf("OBR() failed with %v", err)
			}
			if obr == nil {
				t.Error("OBR() got nil OBR segment, want non nil")
			}

			obxs, err := m.AllOBX()
			if err != nil {
				t.Fatalf("AllOBX() failed with %v", err)
			}
			gotSetIDs := testhl7.OBXFieldsFromOBXs(t, obxs, testhl7.OBXSetID)
			if diff := cmp.Diff(tt.wantOBXSetIDs, gotSetIDs); diff != "" {
				t.Errorf("OBX.SetIDs got diff %v", diff)
			}

			nte, err := m.AllNTE()
			if err != nil {
				t.Errorf("AllNTE() failed with %v", err)
			}
			if got, want := len(nte), tt.numNotes; got != want {
				t.Errorf("len(nte)=%v, want %v", got, want)
			}
		})
	}
}

func TestBuildResultORU_StartCountingAtInitialSetID(t *testing.T) {
	eventTime := time.Date(2018, 4, 28, 22, 38, 44, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 44, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()

	orderWithResults := testOrderWithResult(eventTime)
	orderWithResults.NumberOfPreviousResults = 7
	orderWithResults.Results = append(orderWithResults.Results, &ir.Result{
		TestName: &ir.CodedElement{
			ID:           "lpdc-2011",
			Text:         "Creatinine",
			CodingSystem: "WinPath",
		},
		Value:        "710",
		Unit:         "UML",
		ValueType:    "NM",
		Range:        "39.00 - 308.00",
		Status:       "F",
		AbnormalFlag: "HIGH",
		Notes:        []string{"Note3", "Note4"},
	})

	tests := []struct {
		name string
		f    func(*HeaderInfo, *ir.PatientInfo, *ir.Order, time.Time) (*HL7Message, error)
	}{
		{name: "R01", f: BuildResultORUR01},
		{name: "R03", f: BuildResultORUR03},
		{name: "R32", f: BuildResultORUR32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			oru, err := tt.f(header, patientInfo, orderWithResults, msgTime)
			if err != nil {
				t.Fatalf("BuildResultORU%s(%v, %v, %v, %v) failed with %v", tt.name, header, patientInfo, orderWithResults, msgTime, err)
			}
			mo := hl7.NewParseMessageOptions()
			mo.TimezoneLoc = time.UTC
			m, err := hl7.ParseMessageWithOptions([]byte(oru.Message), mo)
			if err != nil {
				t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", oru.Message, mo, err)
			}

			ps, err := m.ParseAll("OBX")
			if err != nil {
				t.Fatalf("ParseAll() failed with %v", err)
			}
			obx, ok := ps.([]*hl7.OBX)
			if !ok {
				t.Fatalf("m.ParseAll(OBX) got %v, want slice of *hl7.OBX", ps)
			}

			if got, want := len(obx), 3; got != want {
				t.Fatalf("len(obx)=%v, want %v", got, want)
			}
			gotSetIDs := testhl7.OBXFieldsFromOBXs(t, obx, testhl7.OBXSetID)
			if diff := cmp.Diff([]string{"8", "9", "10"}, gotSetIDs); diff != "" {
				t.Errorf("OBX.SetIDs got diff %v", diff)
			}
		})
	}
}

func TestBuildOrderORMO01(t *testing.T) {
	eventTime := time.Date(2018, 4, 28, 22, 38, 44, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 44, 0, time.UTC)
	patientInfo := testPatientInfo()
	order := testOrder(eventTime)
	order.ResultsForORM = []*ir.Result{
		{
			TestName: &ir.CodedElement{
				ID: "random-test-id",
			},
		},
	}
	order.NotesForORM = []string{"Random order Note 1", "Random order Note 2"}

	header := testHeader()

	ormR01, err := BuildOrderORMO01(header, patientInfo, order, msgTime)
	if err != nil {
		t.Fatalf("BuildOrderORMO01(%v, %v, %v, %v) failed with %v", header, patientInfo, order, msgTime, err)
	}
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC

	m, err := hl7.ParseMessageWithOptions([]byte(ormR01.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", ormR01.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ORM"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "O01"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn != nil {
		t.Error("EVN() got not nil EVN segment, want <nil>")
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}

	orc, err := m.ORC()
	if err != nil {
		t.Fatalf("ORC() failed with %v", err)
	}
	if orc == nil {
		t.Error("ORC() got nil ORC segment, want non nil")
	}

	obr, err := m.OBR()
	if err != nil {
		t.Fatalf("OBR() failed with %v", err)
	}
	if obr == nil {
		t.Error("OBR() got nil OBR segment, want non nil")
	}

	obx, err := m.AllOBX()
	if err != nil {
		t.Fatalf("AllOBX() failed with %v", err)
	}
	if got, want := len(obx), 1; got != want {
		t.Fatalf("len(obx)=%v, want %v", got, want)
	}
	if got, want := obx[0].ObservationIdentifier.Identifier.String(), "random-test-id"; got != want {
		t.Errorf("obx[0].ObservationIdentifier.Identifier.String()=%v, want %v", got, want)
	}
	gotSetIDs := testhl7.OBXFieldsFromOBXs(t, obx, testhl7.OBXSetID)
	if diff := cmp.Diff([]string{"1"}, gotSetIDs); diff != "" {
		t.Errorf("OBX.SetIDs got diff %v", diff)
	}

	nte, err := m.AllNTE()
	if err != nil {
		t.Fatalf("AllNTE() failed with %v", err)
	}
	if got, want := len(nte), 2; got != want {
		t.Fatalf("len(nte)=%v, want %v", got, want)
	}
	if got, want := string(nte[0].Comment[0]), "Random order Note 1"; got != want {
		t.Errorf("string(nte[0].Comment[0])=%v, want %v", got, want)
	}
	if got, want := string(nte[1].Comment[0]), "Random order Note 2"; got != want {
		t.Errorf("string(nte[1].Comment[0])=%v, want %v", got, want)
	}
}

func TestBuildBedSwapADTA17(t *testing.T) {
	mergeTime := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	patientInfo.Location.Poc = "onc-poc"
	otherPatientInfo := testPatientInfo()
	otherPatientInfo.Location.Poc = "another-poc"
	header := testHeader()

	adt, err := BuildBedSwapADTA17(header, patientInfo, mergeTime, msgTime, otherPatientInfo)
	if err != nil {
		t.Fatalf("BuildBedSwapADTA17(%v, %v, %v, %v, %v) failed with %v", header, patientInfo, mergeTime, msgTime, otherPatientInfo, err)
	}

	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A17"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A17"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), mergeTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}

	pids, err := m.AllPID()
	if err != nil {
		t.Fatalf("AllPID() failed with %v", err)
	}
	if got, want := len(pids), 2; got != want {
		t.Fatalf("len(pids)=%v, want %v", got, want)
	}

	pv1s, err := m.AllPV1()
	if err != nil {
		t.Fatalf("AllPV1() failed with %v", err)
	}
	if got, want := len(pv1s), 2; got != want {
		t.Fatalf("len(pv1s)=%v, want %v", got, want)
	}
	if diff := cmp.Diff(pv1s[0].AssignedPatientLocation, pv1s[1].AssignedPatientLocation); diff == "" {
		t.Error("PV1.AssignedPatientLocation returned no diff between the two PV1 segments; want diff")
	}
}

func TestBuildMergeADTA34(t *testing.T) {
	mergeTime := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()
	mrn := "123"

	adt, err := BuildMergeADTA34(header, patientInfo, mergeTime, msgTime, mrn)
	if err != nil {
		t.Fatalf("BuildMergeADTA34(%v, %v, %v, %v, %v) failed with %v", header, patientInfo, mergeTime, msgTime, mrn, err)
	}

	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A34"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A34"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), mergeTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	mrg, err := m.MRG()
	if err != nil {
		t.Fatalf("MRG() failed with %v", err)
	}
	if mrg == nil {
		t.Fatal("MRG() got nil MRG segment, want non nil")
	}
	if got, want := len(mrg.PriorPatientIdentifierList), 1; got != want {
		t.Errorf("len(mrg.PriorPatientIdentifierList)=%v, want %v", got, want)
	}
	if got, want := mrg.PriorPatientIdentifierList[0].IDNumber.String(), mrn; got != want {
		t.Errorf("mrg.PriorPatientIdentifierList[0].ID.String()=%v, want %v", got, want)
	}
}

func TestBuildMergeADTA40(t *testing.T) {
	mergeTime := time.Date(2018, 4, 28, 22, 38, 14, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 14, 0, time.UTC)
	patientInfo := testPatientInfo()
	header := testHeader()
	mrns := []string{"123", "456"}

	adt, err := BuildMergeADTA40(header, patientInfo, mergeTime, msgTime, mrns)
	if err != nil {
		t.Fatalf("BuildMergeADTA40(%v, %v, %v, %v, %v) failed with %v", header, patientInfo, mergeTime, msgTime, mrns, err)
	}

	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(adt.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", adt.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "ADT"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "A40"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "A40"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), mergeTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pd1, err := m.PD1()
	if err != nil {
		t.Fatalf("PD1() failed with %v", err)
	}
	if pd1 == nil {
		t.Error("PD1() got nil PD1 segment, want non nil")
	}

	mrg, err := m.MRG()
	if err != nil {
		t.Fatalf("MRG() failed with %v", err)
	}
	if mrg == nil {
		t.Fatal("MRG() got nil MRG segment, want non nil")
	}
	if got, want := len(mrg.PriorPatientIdentifierList), 2; got != want {
		t.Errorf("len(mrg.PriorPatientIdentifierList)=%v, want %v", got, want)
	}
	if got, want := mrg.PriorPatientIdentifierList[0].IDNumber.String(), "123"; got != want {
		t.Errorf("mrg.PriorPatientIdentifierList[0].ID.String()=%v, want %v", got, want)
	}
	if got, want := mrg.PriorPatientIdentifierList[1].IDNumber.String(), "456"; got != want {
		t.Errorf("mrg.PriorPatientIdentifierList[1].ID.String()=%v, want %v", got, want)
	}
}

func TestBuildDocumentNotificationMDMT02(t *testing.T) {
	eventTime := time.Date(2018, 4, 28, 22, 38, 44, 0, time.UTC)
	msgTime := time.Date(2018, 4, 28, 22, 39, 44, 0, time.UTC)
	document := document()
	patientInfo := testPatientInfo()
	header := testHeader()

	mdm, err := BuildDocumentNotificationMDMT02(header, patientInfo, document, eventTime, msgTime)
	if err != nil {
		t.Fatalf("BuildDocumentWithContent(%v, %v, %v, %v, %v) failed with %v", header, patientInfo, document, eventTime, msgTime, err)
	}

	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	m, err := hl7.ParseMessageWithOptions([]byte(mdm.Message), mo)
	if err != nil {
		t.Fatalf("ParseMessageWithOptions(%v, %v) failed with %v", mdm.Message, mo, err)
	}

	msh, err := m.MSH()
	if err != nil {
		t.Fatalf("MSH() failed with %v", err)
	}
	if msh == nil {
		t.Fatal("MSH() got nil MSH segment, want non nil")
	}
	if got, want := msh.MessageType.MessageCode.String(), "MDM"; got != want {
		t.Errorf("msh.MessageType.MessageType.String()=%v, want %v", got, want)
	}
	if got, want := msh.MessageType.TriggerEvent.String(), "T02"; got != want {
		t.Errorf("msh.MessageType.TriggerEvent.String()=%v, want %v", got, want)
	}
	if got, want := msh.DateTimeOfMessage.Time.Second(), msgTime.In(hl7.Location).Second(); got != want {
		t.Errorf("msh.DateTimeOfMessage.Time.Second()=%v, want %v", got, want)
	}

	evn, err := m.EVN()
	if err != nil {
		t.Fatalf("EVN() failed with %v", err)
	}
	if evn == nil {
		t.Fatal("EVN() got nil EVN segment, want non nil")
	}
	if got, want := evn.EventTypeCode.String(), "T02"; got != want {
		t.Errorf("evn.EventTypeCode.String()=%v, want %v", got, want)
	}
	if got, want := evn.RecordedDateTime.Time.Second(), eventTime.In(hl7.Location).Second(); got != want {
		t.Errorf("evn.RecordedDateTime.Time.Second()=%v, want %v", got, want)
	}

	pid, err := m.PID()
	if err != nil {
		t.Fatalf("PID() failed with %v", err)
	}
	if pid == nil {
		t.Error("PID() got nil PID segment, want non nil")
	}

	pv1, err := m.PV1()
	if err != nil {
		t.Fatalf("PV1() failed with %v", err)
	}
	if pv1 == nil {
		t.Error("PV1() got nil PV1 segment, want non nil")
	}

	txa, err := m.TXA()
	if err != nil {
		t.Fatalf("TXA() failed with %v", err)
	}
	if txa == nil {
		t.Error("TXA() got nil TXA segment, want non nil")
	}

	obx, err := m.AllOBX()
	if err != nil {
		t.Fatalf("AllOBX() failed with %v", err)
	}
	if obx == nil {
		t.Fatal("OBX() got nil OBX segment, want non nil")
	}
	if got, want := len(obx), 2; got != want {
		t.Fatalf("len(obx)=%v, want %v", got, want)
	}

	gotSetIDs := testhl7.OBXFieldsFromOBXs(t, obx, testhl7.OBXSetID)
	if diff := cmp.Diff([]string{"1", "2"}, gotSetIDs); diff != "" {
		t.Errorf("OBX.SetIDs got diff %v", diff)
	}
	gotValueType := testhl7.OBXFieldsFromOBXs(t, obx, testhl7.ValueType)
	if diff := cmp.Diff([]string{"TX", "TX"}, gotValueType); diff != "" {
		t.Errorf("OBX.ValueType got diff %v", diff)
	}
}

func testOrderWithResult(now time.Time) *ir.Order {
	order := testOrder(now)
	order.Results = []*ir.Result{{
		TestName: &ir.CodedElement{
			ID:           "lpdc-2011",
			Text:         "Creatinine",
			CodingSystem: "WinPath",
		},
		Value:               "700",
		Unit:                "UML",
		ValueType:           "NM",
		Range:               "39.00 - 308.00",
		Status:              "F",
		AbnormalFlag:        "HIGH",
		Notes:               []string{"Note1", "Note2"},
		ObservationDateTime: order.CollectedDateTime,
	}, {
		TestName: &ir.CodedElement{
			ID:           "lpdc-2804",
			Text:         "Potassium",
			CodingSystem: "WinPath",
		},
		Value: "600",
	}}
	return order
}

func testOrder(now time.Time) *ir.Order {
	return &ir.Order{
		OrderProfile: &ir.CodedElement{
			ID:           "lpdc-3969",
			Text:         "UREA AND ELECTROLYTES",
			CodingSystem: "WinPath",
		},
		Placer:        "9984058",
		Filler:        "1902082",
		OrderDateTime: ir.NewValidTime(now),
		OrderControl:  "RE",
		OrderStatus:   "IP",
		ResultsStatus: "C",
	}
}

func orderWithClinicalNote(now time.Time, content string) *ir.Order {
	return &ir.Order{
		OrderProfile: &ir.CodedElement{
			ID:            "document-type",
			Text:          "document-type",
			AlternateText: "document-title",
		},
		DiagnosticServID: "MDOC",
		OrderingProvider: testDoctor(),
		Results: []*ir.Result{{
			ObservationDateTime: ir.NewValidTime(now),
			ClinicalNote: &ir.ClinicalNote{
				DateTime:     ir.NewValidTime(now),
				DocumentType: "ECG",
				DocumentID:   "document_id",
				Contents: []*ir.ClinicalNoteContent{{
					ContentType:      "PNG",
					DocumentContent:  content,
					DocumentEncoding: "BASE64",
				}, {
					ContentType:     "rtf",
					DocumentContent: content,
				}},
			},
		}},
	}
}

func document() *ir.Document {
	return &ir.Document{
		ActivityDateTime:         ir.NewValidTime(time.Date(2019, 6, 15, 8, 13, 40, 0, time.UTC)),
		EditDateTime:             ir.NewValidTime(time.Date(2019, 11, 4, 8, 13, 40, 0, time.UTC)),
		DocumentCompletionStatus: "DO",
		DocumentType:             "DS",
		ObservationIdentifier: &ir.CodedElement{
			ID:           "Established Patient 15",
			Text:         "Established Patient 15",
			CodingSystem: "Simulation",
		},
		UniqueDocumentNumber: "9298345CE5003",
		ContentLine: []string{
			"Patient",
			"Name : SULLY, J K (65yo, F) ID# 47Q66Q585",
		},
	}
}

func testPatientInfo() *ir.PatientInfo {
	ap := &ir.AssociatedParty{
		Person: &ir.Person{
			Prefix:     "Mr",
			FirstName:  "John",
			MiddleName: "George",
			Surname:    "Smiths",
			Suffix:     "Senior",
			Gender:     "M",
			Address: &ir.Address{
				FirstLine:  "5 Goodwill Hunting Road",
				City:       "London",
				PostalCode: "N1D 4AG",
				Country:    "GBR",
				Type:       "HOME",
			},
			PhoneNumber: "020 7031 4000",
			MRN:         "21124992125291505",
			NHS:         "3338933381",
		},
		Relationship: &ir.CodedElement{ID: "S", Text: "SPOUSE"},
		ContactRole:  &ir.CodedElement{ID: "F", Text: "FAMILYMEM"},
	}
	al := &ir.Allergy{
		Type:        "FA", // Food allergy.
		Description: ir.CodedElement{ID: "E", Text: "egg-containing compound", CodingSystem: "ZAL"},
		Severity:    "MO", // Moderate.
		Reaction:    "Skin rash",
	}
	p := testPersonFemale()
	patientInfo := &ir.PatientInfo{
		Person:          p,
		Class:           "INPATIENT",
		Type:            "EMERGENCY",
		VisitID:         12341234,
		HospitalService: "180",
		Location: &ir.PatientLocation{
			Poc:          "RAL 12 West",
			Room:         "Bay01",
			Bed:          "Bed10",
			Facility:     "RAL RF",
			LocationType: "BED",
			Building:     "RFH",
		},
		PriorLocation: &ir.PatientLocation{
			Poc:          "RAL 12 East",
			Room:         "Bay02",
			Bed:          "Bed11",
			Facility:     "RAL RF",
			LocationType: "BED",
			Building:     "RFH",
		},
		AttendingDoctor:           testDoctor(),
		AdmissionDate:             ir.NewValidTime(defaultAdmissionDate),
		DischargeDate:             ir.NewValidTime(defaultDischargeDate),
		TransferDate:              ir.NewValidTime(defaultTransferDate),
		ExpectedAdmitDateTime:     ir.NewValidTime(defaultExpectedAdmissionDate),
		ExpectedDischargeDateTime: ir.NewValidTime(defaultExpectedDischargeDate),
		ExpectedTransferDateTime:  ir.NewValidTime(defaultExpectedTransferDate),
		AssociatedParties:         []*ir.AssociatedParty{ap},
		Allergies:                 []*ir.Allergy{al},
		Diagnoses:                 []*ir.DiagnosisOrProcedure{testDiagnosis()},
		Procedures:                []*ir.DiagnosisOrProcedure{testProcedure()},
		AdmitReason:               "Eye problems",
	}
	return patientInfo
}

func testDoctor() *ir.Doctor {
	return &ir.Doctor{
		ID:        defaultDoctorID,
		Surname:   defaultDoctorSurname,
		FirstName: defaultDoctorFirstName,
		Prefix:    defaultDoctorPrefix,
	}
}

func testHeader() *HeaderInfo {
	return &HeaderInfo{
		SendingApplication:   "CERNER",
		SendingFacility:      "RAL1",
		ReceivingApplication: "STREAMS",
		ReceivingFacility:    "RAL",
		MessageControlID:     "1",
	}
}

func testPersonFemale() *ir.Person {
	return &ir.Person{
		Prefix:      "Miss",
		FirstName:   "Helen",
		MiddleName:  "Matilda",
		Surname:     "Smiths",
		Suffix:      "Junior",
		Degree:      "Dr",
		Gender:      "F",
		Ethnicity:   &ir.Ethnicity{ID: "A", Text: "White British"},
		Birth:       ir.NewValidTime(time.Date(1994, 7, 4, 12, 35, 18, 0, time.UTC)),
		DateOfDeath: ir.NewValidTime(defaultDateOfDeath),
		Address: &ir.Address{
			FirstLine:  "1 Goodwill Hunting Road",
			SecondLine: "Kings Cross",
			City:       "London",
			PostalCode: "N1C 4AG",
			Country:    "GBR",
			Type:       "HOME",
		},
		PhoneNumber:    "020 7031 3000",
		MRN:            "12529150521124992",
		NHS:            "3333381389",
		DeathIndicator: "DECEASED",
	}
}

func testDiagnosis() *ir.DiagnosisOrProcedure {
	return &ir.DiagnosisOrProcedure{
		Description: &ir.CodedElement{ID: "A01.0", Text: "Typhoid fever"},
		Type:        "Admitting",
		Clinician:   testDoctor(),
		DateTime:    ir.NewValidTime(defaultDiagnoseDate),
	}
}

func testProcedure() *ir.DiagnosisOrProcedure {
	return &ir.DiagnosisOrProcedure{
		Description: &ir.CodedElement{ID: defaultProcedureCodeID, Text: defaultProcedureCodeDescription},
		Type:        defaultProcedureFunctionalType,
		Clinician:   testDoctor(),
		DateTime:    ir.NewValidTime(defaultProcedureDate),
	}
}
