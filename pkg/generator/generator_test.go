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

package generator

import (
	"context"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/constants"
	"github.com/google/simhospital/pkg/doctor"
	"github.com/google/simhospital/pkg/generator/codedelement"
	"github.com/google/simhospital/pkg/generator/header"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/orderprofile"
	"github.com/google/simhospital/pkg/pathway"
	"github.com/google/simhospital/pkg/state"
	"github.com/google/simhospital/pkg/test"
	"github.com/google/simhospital/pkg/test/testaddress"
	"github.com/google/simhospital/pkg/test/testclock"
	"github.com/google/simhospital/pkg/test/testdate"
	"github.com/google/simhospital/pkg/test/testid"
	"github.com/google/simhospital/pkg/test/testperson"
	"github.com/google/simhospital/pkg/test/testwrite"
)

var (
	urineElectrolytesCE = &ir.CodedElement{ID: "lpdc-3969", Text: "UREA AND ELECTROLYTES", CodingSystem: "WinPath"}

	oneDayAgoDuration  = -24 * time.Hour
	twoDaysAgoDuration = -48 * time.Hour

	defaultDate = time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)

	defaultCities = []string{"London", "Cambridge"}
)

func TestNewGeneratorPublicMessageConfiguration(t *testing.T) {
	ctx := context.Background()
	// If some files are not parsable or accessible, things will crash.
	c, err := config.LoadHL7Config(ctx, test.MessageConfigProd)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigProd, err)
	}
	h, err := config.LoadHeaderConfig(ctx, test.HeaderConfigProd)
	if err != nil {
		t.Fatalf("LoadHeaderConfig(%s) failed with %v", test.HeaderConfigProd, err)
	}
	f := test.DataFiles[test.Prod]
	dc, err := config.LoadData(ctx, f, c)
	if err != nil {
		t.Fatalf("LoadData(%+v, %+v) failed with %v", f, c, err)
	}
	d, err := doctor.LoadDoctors(ctx, test.DoctorsConfigProd)
	if err != nil {
		t.Fatalf("LoadDoctors(%s) failed with %v", test.DoctorsConfigProd, err)
	}
	op, err := orderprofile.Load(ctx, test.OrderProfilesConfigProd, c)
	if err != nil {
		t.Fatalf("orderprofile.Load(%s, %+v) failed with %v", test.OrderProfilesConfigProd, c, err)
	}
	NewGenerator(
		Config{
			Clock:         testclock.New(defaultDate),
			HL7Config:     c,
			Header:        h,
			Data:          dc,
			Doctors:       d,
			OrderProfiles: op,
		},
	)
}

func TestNewPerson(t *testing.T) {
	ctx := context.Background()
	g := testGenerator(ctx, t, Config{})
	pathway := &pathway.Person{}
	got := g.NewPerson(pathway)
	if got == nil {
		t.Errorf("g.NewPerson(%v)=%v, want non-nil value", pathway, got)
	}
}

func originalPatientInfo() *ir.PatientInfo {
	return &ir.PatientInfo{
		Person:     testperson.New(),
		Diagnoses:  []*ir.DiagnosisOrProcedure{},
		Procedures: []*ir.DiagnosisOrProcedure{},
		Allergies:  []*ir.Allergy{},
	}
}

func TestUpdateFromPathway(t *testing.T) {
	ctx := context.Background()
	g := testGenerator(ctx, t, Config{})

	cases := []struct {
		name    string
		patient *ir.PatientInfo
		pathway *pathway.UpdatePerson
		want    func() *ir.PatientInfo
	}{
		{
			name:    "No changes",
			patient: originalPatientInfo(),
			pathway: &pathway.UpdatePerson{},
			want:    originalPatientInfo,
		}, {
			name:    "Update Person: MRN and NHS",
			patient: originalPatientInfo(),
			pathway: &pathway.UpdatePerson{
				Person: &pathway.Person{MRN: "888888", NHS: "9377865972"},
			},
			want: func() *ir.PatientInfo {
				pi := originalPatientInfo()
				pi.Person.MRN = "888888"
				pi.Person.NHS = "9377865972"
				return pi
			},
		}, {
			name: "Update Diagnosis and Procedure: init to empty slices",
			patient: &ir.PatientInfo{
				Person:    testperson.New(),
				Allergies: []*ir.Allergy{},
			},
			pathway: &pathway.UpdatePerson{},
			want:    originalPatientInfo,
		}, {
			name:    "Update Allergies",
			patient: originalPatientInfo(),
			pathway: &pathway.UpdatePerson{
				Allergies: []pathway.Allergy{
					{
						Type:                   "FOOD",
						Code:                   "E",
						Description:            "egg-containing compound",
						Severity:               "SEVERE",
						Reaction:               "Rash",
						CodingSystem:           "AL",
						IdentificationDateTime: &pathway.DateTime{Time: &defaultDate},
					},
				},
			},
			want: func() *ir.PatientInfo {
				pi := originalPatientInfo()
				pi.Allergies = []*ir.Allergy{
					{
						Type:                   "FOOD",
						Description:            ir.CodedElement{ID: "E", Text: "egg-containing compound", CodingSystem: "AL"},
						Severity:               "SEVERE",
						Reaction:               "Rash",
						IdentificationDateTime: ir.NewValidTime(defaultDate),
					},
				}
				return pi
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := new(ir.PatientInfo)
			*got = *tc.patient

			g.UpdateFromPathway(got, tc.pathway)
			if diff := cmp.Diff(tc.want(), got); diff != "" {
				t.Errorf("g.UpdateFromPathway(%+v, %+v) diff: (-want, +got):\n%s", tc.patient, tc.pathway, diff)
			}
		})
	}
}

func TestNewPerson_MRNsAreIncremental(t *testing.T) {
	ctx := context.Background()
	g := testGenerator(ctx, t, Config{})

	person := g.NewPerson(&pathway.Person{})

	patientInfo := &ir.PatientInfo{Person: person}
	g.UpdateFromPathway(patientInfo, &pathway.UpdatePerson{Person: &pathway.Person{}})

	g.UpdateFromPathway(patientInfo, &pathway.UpdatePerson{Person: &pathway.Person{MRN: "888888"}})

	// Even if we've updated multiple patients, only NewPerson calls should generate new MRNs.
	person = g.NewPerson(&pathway.Person{})
	if got, want := person.MRN, "2"; got != want {
		t.Errorf("UpdateFromPathway() MRN got %v, want %v", got, want)
	}
}

func TestUpdateFromPathwaySetProceduresAndDiagnoses(t *testing.T) {
	ctx := context.Background()
	// Allow only one doctor to be "randomly" chosen for testing purposes.
	fName := testwrite.BytesToFile(t, []byte(`
- id: "id-1"
  surname: "surname-1"
  firstname: "firstname-1"
  prefix: "prefix-1"
  specialty: "specialty-1"`))

	d, err := doctor.LoadDoctors(ctx, fName)
	if err != nil {
		t.Fatalf("LoadDoctors() failed with %v", err)
	}

	wantDoctor := &ir.Doctor{
		ID:        "id-1",
		Surname:   "surname-1",
		FirstName: "firstname-1",
		Prefix:    "prefix-1",
		Specialty: "specialty-1",
	}

	tmpConfig := testwrite.BytesToFile(t, []byte(`
allergy:
  types:
    - "FA"
  severities:
    - "SV"
  coding_system: "ZAL"
procedure:
  types:
    - "P"
  coding_system: "PCS"
diagnosis:
  types:
    - "D"
  coding_system: "DCS"
`))

	// Allow only one single procedure to be "randomly" chosen for testing purposes.
	tmpProcedures := testwrite.BytesToFile(t, []byte(`
P24.9,Procedure1,1
`))

	// Allow only one single diagnosis to be "randomly" chosen for testing purposes.
	tmpDiagnoses := testwrite.BytesToFile(t, []byte(`
A01.1,Diagnosis1,1
`))

	c, err := config.LoadHL7Config(ctx, tmpConfig)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", tmpConfig, err)
	}
	f := test.DataFiles[test.Test]
	f.Procedures = tmpProcedures
	f.Diagnoses = tmpDiagnoses
	dc, err := config.LoadData(ctx, f, c)
	if err != nil {
		t.Fatalf("LoadData(%+v, %+v) failed with %v", f, c, err)
	}
	now := time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC)
	oneDayAgo := time.Date(2018, 2, 11, 0, 0, 0, 0, time.UTC)
	twoDaysAgo := now.Add(twoDaysAgoDuration)

	tests := []struct {
		name           string
		procedures     []*pathway.DiagnosisOrProcedure
		diagnoses      []*pathway.DiagnosisOrProcedure
		wantProcedures []*ir.DiagnosisOrProcedure
		wantDiagnoses  []*ir.DiagnosisOrProcedure
	}{{
		name: "Random, no datetime recorded",
		procedures: []*pathway.DiagnosisOrProcedure{
			{Description: constants.RandomString, DateTime: &pathway.DateTime{NoDateTimeRecorded: true}},
		},
		diagnoses: []*pathway.DiagnosisOrProcedure{
			{Description: constants.RandomString, DateTime: &pathway.DateTime{NoDateTimeRecorded: true}},
		},
		wantProcedures: []*ir.DiagnosisOrProcedure{{
			Description: &ir.CodedElement{
				ID:           "P24.9",
				Text:         "Procedure1",
				CodingSystem: "PCS",
			},
			Type:      "P",
			Clinician: wantDoctor,
			DateTime:  ir.NewInvalidTime(),
		}},
		wantDiagnoses: []*ir.DiagnosisOrProcedure{{
			Description: &ir.CodedElement{
				ID:           "A01.1",
				Text:         "Diagnosis1",
				CodingSystem: "DCS",
			},
			Type:      "D",
			Clinician: wantDoctor,
			DateTime:  ir.NewInvalidTime(),
		}},
	}, {
		name: "all fields present, Time DateTime",
		procedures: []*pathway.DiagnosisOrProcedure{{
			Type:        "A",
			Code:        "A01.1",
			Description: "proc1",
			DateTime:    &pathway.DateTime{Time: &defaultDate},
		}},
		diagnoses: []*pathway.DiagnosisOrProcedure{{
			Type:        "B",
			Code:        "B01.1",
			Description: "diag1",
			DateTime:    &pathway.DateTime{Time: &defaultDate},
		}},
		wantProcedures: []*ir.DiagnosisOrProcedure{{
			Type: "A",
			Description: &ir.CodedElement{
				ID:   "A01.1",
				Text: "proc1",
			},
			Clinician: wantDoctor,
			DateTime:  ir.NewValidTime(defaultDate),
		}},
		wantDiagnoses: []*ir.DiagnosisOrProcedure{{
			Type: "B",
			Description: &ir.CodedElement{
				ID:   "B01.1",
				Text: "diag1",
			},
			Clinician: wantDoctor,
			DateTime:  ir.NewValidTime(defaultDate),
		}},
	}, {
		name: "all fields present, TimeFromNow DateTime",
		procedures: []*pathway.DiagnosisOrProcedure{{
			Type:        "B",
			Code:        "B01.1",
			Description: "proc1",
			DateTime:    &pathway.DateTime{TimeFromNow: &twoDaysAgoDuration},
		}},
		diagnoses: []*pathway.DiagnosisOrProcedure{{
			Type:        "A",
			Code:        "A01.1",
			Description: "diag1",
			DateTime:    &pathway.DateTime{TimeFromNow: &twoDaysAgoDuration},
		}},
		wantProcedures: []*ir.DiagnosisOrProcedure{{
			Type: "B",
			Description: &ir.CodedElement{
				ID:   "B01.1",
				Text: "proc1",
			},
			Clinician: wantDoctor,
			DateTime:  ir.NewValidTime(twoDaysAgo),
		}},
		wantDiagnoses: []*ir.DiagnosisOrProcedure{{
			Type: "A",
			Description: &ir.CodedElement{
				ID:   "A01.1",
				Text: "diag1",
			},
			Clinician: wantDoctor,
			DateTime:  ir.NewValidTime(twoDaysAgo),
		}},
	}, {
		name: "all fields present, NoDateTimeRecorded DateTime",
		procedures: []*pathway.DiagnosisOrProcedure{{
			Type:        "C",
			Code:        "C01.1",
			Description: "proc1",
			DateTime:    &pathway.DateTime{NoDateTimeRecorded: true},
		}},
		diagnoses: []*pathway.DiagnosisOrProcedure{{
			Type:        "D",
			Code:        "D01.1",
			Description: "diag1",
			DateTime:    &pathway.DateTime{NoDateTimeRecorded: true},
		}},
		wantProcedures: []*ir.DiagnosisOrProcedure{{
			Type: "C",
			Description: &ir.CodedElement{
				ID:   "C01.1",
				Text: "proc1",
			},
			Clinician: wantDoctor,
			DateTime:  ir.NewInvalidTime(),
		}},
		wantDiagnoses: []*ir.DiagnosisOrProcedure{{
			Type: "D",
			Description: &ir.CodedElement{
				ID:   "D01.1",
				Text: "diag1",
			},
			Clinician: wantDoctor,
			DateTime:  ir.NewInvalidTime(),
		}},
	}, {
		name: "mapped description",
		procedures: []*pathway.DiagnosisOrProcedure{{
			Type:     "A",
			Code:     "P24.9",
			DateTime: &pathway.DateTime{Time: &defaultDate},
		}},
		diagnoses: []*pathway.DiagnosisOrProcedure{{
			Type:     "B",
			Code:     "A01.1",
			DateTime: &pathway.DateTime{Time: &defaultDate},
		}},
		wantProcedures: []*ir.DiagnosisOrProcedure{{
			Type: "A",
			Description: &ir.CodedElement{
				ID:   "P24.9",
				Text: "Procedure1",
			},
			Clinician: wantDoctor,
			DateTime:  ir.NewValidTime(defaultDate),
		}},
		wantDiagnoses: []*ir.DiagnosisOrProcedure{{
			Type: "B",
			Description: &ir.CodedElement{
				ID:   "A01.1",
				Text: "Diagnosis1",
			},
			Clinician: wantDoctor,
			DateTime:  ir.NewValidTime(defaultDate),
		}},
	}, {
		name: "no description",
		procedures: []*pathway.DiagnosisOrProcedure{{
			Type:     "A",
			Code:     "XXX.1",
			DateTime: &pathway.DateTime{Time: &defaultDate},
		}},
		diagnoses: []*pathway.DiagnosisOrProcedure{{
			Type:     "B",
			Code:     "YYY.2",
			DateTime: &pathway.DateTime{Time: &defaultDate},
		}},
		wantProcedures: []*ir.DiagnosisOrProcedure{{
			Type: "A",
			Description: &ir.CodedElement{
				ID: "XXX.1",
			},
			Clinician: wantDoctor,
			DateTime:  ir.NewValidTime(defaultDate),
		}},
		wantDiagnoses: []*ir.DiagnosisOrProcedure{{
			Type: "B",
			Description: &ir.CodedElement{
				ID: "YYY.2",
			},
			Clinician: wantDoctor,
			DateTime:  ir.NewValidTime(defaultDate),
		}},
	}, {
		name: "mapped code",
		procedures: []*pathway.DiagnosisOrProcedure{{
			Type:        "A",
			Description: "Procedure1",
			DateTime:    &pathway.DateTime{Time: &defaultDate},
		}},
		diagnoses: []*pathway.DiagnosisOrProcedure{{
			Type:        "B",
			Description: "Diagnosis1",
			DateTime:    &pathway.DateTime{Time: &defaultDate},
		}},
		wantProcedures: []*ir.DiagnosisOrProcedure{{
			Type: "A",
			Description: &ir.CodedElement{
				ID:   "P24.9",
				Text: "Procedure1",
			},
			Clinician: wantDoctor,
			DateTime:  ir.NewValidTime(defaultDate),
		}},
		wantDiagnoses: []*ir.DiagnosisOrProcedure{{
			Type: "B",
			Description: &ir.CodedElement{
				ID:   "A01.1",
				Text: "Diagnosis1",
			},
			Clinician: wantDoctor,
			DateTime:  ir.NewValidTime(defaultDate),
		}},
	}, {
		name: "no code",
		procedures: []*pathway.DiagnosisOrProcedure{{
			Type:        "A",
			Description: "ProcedureX",
			DateTime:    &pathway.DateTime{Time: &defaultDate},
		}},
		diagnoses: []*pathway.DiagnosisOrProcedure{{
			Type:        "B",
			Description: "DiagnosisX",
			DateTime:    &pathway.DateTime{Time: &defaultDate},
		}},
		wantProcedures: []*ir.DiagnosisOrProcedure{{
			Type: "A",
			Description: &ir.CodedElement{
				Text: "ProcedureX",
			},
			Clinician: wantDoctor,
			DateTime:  ir.NewValidTime(defaultDate),
		}},
		wantDiagnoses: []*ir.DiagnosisOrProcedure{{
			Type: "B",
			Description: &ir.CodedElement{
				Text: "DiagnosisX",
			},
			Clinician: wantDoctor,
			DateTime:  ir.NewValidTime(defaultDate),
		}},
	}, {
		name: "From Pathway, multiple",
		procedures: []*pathway.DiagnosisOrProcedure{
			{
				Type:        "A",
				Code:        "A01.1",
				Description: "proc1",
				DateTime:    &pathway.DateTime{TimeFromNow: &oneDayAgoDuration},
			},
			{
				Type:        "B",
				Code:        "B01.1",
				Description: "proc2",
				DateTime:    &pathway.DateTime{Time: &twoDaysAgo},
			},
		},
		diagnoses: []*pathway.DiagnosisOrProcedure{
			{
				Type:        "C",
				Code:        "C01.1",
				Description: "diag1",
				DateTime:    &pathway.DateTime{TimeFromNow: &oneDayAgoDuration},
			},
			{
				Type:        "D",
				Code:        "D01.1",
				Description: "diag2",
				DateTime:    &pathway.DateTime{Time: &twoDaysAgo},
			},
		},
		wantProcedures: []*ir.DiagnosisOrProcedure{{
			Description: &ir.CodedElement{
				ID:   "A01.1",
				Text: "proc1",
			},
			Type:      "A",
			Clinician: wantDoctor,
			DateTime:  ir.NewValidTime(oneDayAgo),
		}, {
			Description: &ir.CodedElement{
				ID:   "B01.1",
				Text: "proc2",
			},
			Type:      "B",
			Clinician: wantDoctor,
			DateTime:  ir.NewValidTime(twoDaysAgo),
		}},
		wantDiagnoses: []*ir.DiagnosisOrProcedure{{
			Description: &ir.CodedElement{
				ID:   "C01.1",
				Text: "diag1",
			},
			Type:      "C",
			Clinician: wantDoctor,
			DateTime:  ir.NewValidTime(oneDayAgo),
		}, {
			Description: &ir.CodedElement{
				ID:   "D01.1",
				Text: "diag2",
			},
			Type:      "D",
			Clinician: wantDoctor,
			DateTime:  ir.NewValidTime(twoDaysAgo),
		}},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := originalPatientInfo()
			want := originalPatientInfo()
			want.Diagnoses = tc.wantDiagnoses
			want.Procedures = tc.wantProcedures

			pathway := &pathway.UpdatePerson{Procedures: tc.procedures, Diagnoses: tc.diagnoses}

			dg := &testdate.Generator{}
			g := testGeneratorWithDate(ctx, t, now, Config{HL7Config: c, Data: dc, Doctors: d, DateGenerator: dg})
			g.UpdateFromPathway(got, pathway)

			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("g.UpdateFromPathway(%+v, %+v) diff: (-want, +got):\n%s", originalPatientInfo(), pathway, diff)
			}
		})
	}
}

func TestNewPatient(t *testing.T) {
	ctx := context.Background()
	dName := testwrite.BytesToFile(t, []byte(`
- id: "id-1"
  surname: "surname-1"
  firstname: "firstname-1"
  prefix: "prefix-1"
  specialty: "specialty-1"`))

	d, err := doctor.LoadDoctors(ctx, dName)
	if err != nil {
		t.Fatalf("LoadDoctors(%s) failed with %v", dName, err)
	}

	existingDoctor := &ir.Doctor{
		ID:        "id-1",
		Surname:   "surname-1",
		FirstName: "firstname-1",
		Prefix:    "prefix-1",
		Specialty: "specialty-1",
	}

	// This empty configuration file simulates a nil primary facility and nil patient class.
	nilHL7Name := testwrite.BytesToFile(t, []byte(``))

	nilHL7Config, err := config.LoadHL7Config(ctx, nilHL7Name)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", nilHL7Name, err)
	}

	hl7Name := testwrite.BytesToFile(t, []byte(`
patient_class:
  outpatient: "OUTPATIENT"
  inpatient: "INPATIENT"
primary_facility:
  organization_name: "Test Primary Facility"
  id_number: "123"
hospital_service: "180"
`))

	hl7Config, err := config.LoadHL7Config(ctx, hl7Name)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", hl7Name, err)
	}

	newDoctor := &ir.Doctor{
		ID:        "123",
		Surname:   "Osman",
		FirstName: "Arthur",
		Prefix:    "Dr",
		Specialty: "Nephrologist",
	}

	person := testperson.New()
	cases := []struct {
		name   string
		conf   Config
		doctor *ir.Doctor
		want   *state.Patient
	}{
		{
			name:   "Nil doctor, primary facility, hospital service and patient class not specified in config",
			conf:   Config{HL7Config: nilHL7Config, Doctors: d},
			doctor: nil,
			want: &state.Patient{
				PatientInfo: &ir.PatientInfo{
					Class:           "",
					Person:          person,
					HospitalService: "",
				},
				Orders:    make(map[string]*ir.Order),
				Documents: make(map[string]*ir.Document),
			},
		}, {
			name:   "Existing doctor, override hospital service",
			conf:   Config{HL7Config: nilHL7Config, Doctors: d},
			doctor: existingDoctor,
			want: &state.Patient{
				PatientInfo: &ir.PatientInfo{
					Class:           "",
					Person:          person,
					HospitalService: existingDoctor.Specialty,
					AttendingDoctor: existingDoctor,
				},
				Orders:    make(map[string]*ir.Order),
				Documents: make(map[string]*ir.Document),
			},
		}, {
			name:   "New doctor, don't override hospital service",
			conf:   Config{HL7Config: nilHL7Config, Doctors: d},
			doctor: newDoctor,
			want: &state.Patient{
				PatientInfo: &ir.PatientInfo{
					Class:           "",
					Person:          person,
					HospitalService: "",
					AttendingDoctor: newDoctor,
				},
				Orders:    make(map[string]*ir.Order),
				Documents: make(map[string]*ir.Document),
			},
		}, {
			name:   "Nil doctor, primary facility, hospital service and patient class from config",
			conf:   Config{HL7Config: hl7Config, Doctors: d},
			doctor: newDoctor,
			want: &state.Patient{
				PatientInfo: &ir.PatientInfo{
					Class:           "OUTPATIENT",
					Person:          person,
					HospitalService: "180",
					AttendingDoctor: newDoctor,
					PrimaryFacility: &ir.PrimaryFacility{
						Organization: "Test Primary Facility",
						ID:           "123",
					},
				},
				Orders:    make(map[string]*ir.Order),
				Documents: make(map[string]*ir.Document),
			},
		}, {
			name:   "Existing doctor, defined config, override hospital service",
			conf:   Config{HL7Config: hl7Config, Doctors: d},
			doctor: existingDoctor,
			want: &state.Patient{
				PatientInfo: &ir.PatientInfo{
					Class:           "OUTPATIENT",
					Person:          person,
					HospitalService: existingDoctor.Specialty,
					AttendingDoctor: existingDoctor,
					PrimaryFacility: &ir.PrimaryFacility{
						Organization: "Test Primary Facility",
						ID:           "123",
					},
				},
				Orders:    make(map[string]*ir.Order),
				Documents: make(map[string]*ir.Document),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := testGenerator(ctx, t, tc.conf)

			got := g.NewPatient(person, tc.doctor)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("g.NewPatient(%+v, %+v) diff: (-want, +got):\n%s", person, tc.doctor, diff)
			}
		})
	}
}

func TestNewDoctor(t *testing.T) {
	ctx := context.Background()
	hl7Name := testwrite.BytesToFile(t, []byte(`
hospital_service: "180"
`))

	hl7Config, err := config.LoadHL7Config(ctx, hl7Name)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", hl7Name, err)
	}

	dName := testwrite.BytesToFile(t, []byte(`
- id: "id-1"
  surname: "surname-1"
  firstname: "firstname-1"
  prefix: "prefix-1"
  specialty: "specialty-1"`))

	existingDoctor := &ir.Doctor{
		ID:        "id-1",
		Surname:   "surname-1",
		FirstName: "firstname-1",
		Prefix:    "prefix-1",
		Specialty: "specialty-1",
	}

	dID := "id-1"
	newID := "123"
	newSurname := "Osman"
	newName := "Arthur"
	newPrefix := "Dr"

	cases := []struct {
		name       string
		consultant *pathway.Consultant
		want       *ir.Doctor
	}{
		{
			name: "Use existing doctor",
			consultant: &pathway.Consultant{
				ID: &dID,
			},
			want: existingDoctor,
		}, {
			name: "New doctor",
			consultant: &pathway.Consultant{
				ID:        &newID,
				Surname:   &newSurname,
				Prefix:    &newPrefix,
				FirstName: &newName,
			},
			want: &ir.Doctor{
				ID:        "123",
				Surname:   "Osman",
				FirstName: "Arthur",
				Prefix:    "Dr",
				Specialty: "180",
			},
		}, {
			name:       "Nil doctor - return random existing doctor",
			consultant: nil,
			want:       existingDoctor,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d, err := doctor.LoadDoctors(ctx, dName)
			if err != nil {
				t.Fatalf("LoadDoctors(%s) failed with %v", dName, err)
			}
			g := testGenerator(ctx, t, Config{HL7Config: hl7Config, Doctors: d})
			got := g.NewDoctor(tc.consultant)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("g.NewDoctor(%+v) diff: (-want, +got):\n%s", tc.consultant, diff)
			}
		})
	}
}

func TestNewDoctorAddDoctorsForFutureUse(t *testing.T) {
	ctx := context.Background()
	// No doctors defined.
	dName := testwrite.BytesToFile(t, []byte(``))
	d, err := doctor.LoadDoctors(ctx, dName)
	if err != nil {
		t.Fatalf("LoadDoctors(%s) failed with %v", dName, err)
	}
	hl7Name := testwrite.BytesToFile(t, []byte(`
hospital_service: "180"
`))

	hl7Config, err := config.LoadHL7Config(ctx, hl7Name)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", hl7Name, err)
	}

	newID := "123"
	newSurname := "Osman"
	newName := "Arthur"
	newPrefix := "Dr"

	want := &ir.Doctor{
		ID:        "123",
		Surname:   "Osman",
		FirstName: "Arthur",
		Prefix:    "Dr",
		Specialty: "180",
	}
	consultant := &pathway.Consultant{
		ID:        &newID,
		Surname:   &newSurname,
		Prefix:    &newPrefix,
		FirstName: &newName,
	}

	g := testGenerator(ctx, t, Config{HL7Config: hl7Config, Doctors: d})
	got := g.NewDoctor(consultant)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("g.NewDoctor(%+v) diff: (-want, +got):\n%s", consultant, diff)
	}

	// Dr Arthur Osman was added to the pool of doctors, so can be randomly selected.
	got = g.NewDoctor(nil)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("g.NewDoctor(%+v) diff: (-want, +got):\n%s", nil, diff)
	}
}

func TestAddAllergies(t *testing.T) {
	now := time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC)
	twoDaysAgo := now.Add(twoDaysAgoDuration)

	fHL7Config := testwrite.BytesToFile(t, []byte(`
allergy:
  types:
    - "FA"
  severities:
    - "SV"
  coding_system: "ZAL"
`))

	// Allow only one single allergy to be "randomly" chosen for testing purposes.
	fAllergies := testwrite.BytesToFile(t, []byte(`
J30.1,Allergy1,1
`))

	// Simplified data config for testing purposes.
	fData := testwrite.BytesToFile(t, []byte(`
allergy:
  reactions:
    - "Sneezing"
  percentage: 100
  maximum_allergies: 1
`))

	ctx := context.Background()
	hl7Config, err := config.LoadHL7Config(ctx, fHL7Config)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", fHL7Config, err)
	}
	f := test.DataFiles[test.Test]
	f.Allergies = fAllergies
	f.DataConfig = fData
	dataConfig, err := config.LoadData(ctx, f, hl7Config)
	if err != nil {
		t.Fatalf("LoadData(%+v, %+v) failed with %v", f, hl7Config, err)
	}

	dg := &testdate.Generator{}

	g := testGeneratorWithDate(ctx, t, now, Config{HL7Config: hl7Config, Data: dataConfig, DateGenerator: dg, Clock: testclock.New(now)})

	pathwayA1 := pathway.Allergy{Type: "MEDICATION", Code: "J301", Description: "Allergic rhinitis due to pollen", Severity: "MODERATE", Reaction: "Skin rash"}
	pathwayA2 := pathway.Allergy{Type: "FOOD", Code: "E", Description: "egg-containing compound", Severity: "SEVERE", Reaction: "Rash"}

	a1 := &ir.Allergy{
		Type:                   "MEDICATION",
		Description:            ir.CodedElement{ID: "J301", Text: "Allergic rhinitis due to pollen", CodingSystem: "ZAL"},
		Severity:               "MODERATE",
		Reaction:               "Skin rash",
		IdentificationDateTime: ir.NewValidTime(now),
	}
	a2 := &ir.Allergy{
		Type:                   "FOOD",
		Description:            ir.CodedElement{ID: "E", Text: "egg-containing compound", CodingSystem: "ZAL"},
		Severity:               "SEVERE",
		Reaction:               "Rash",
		IdentificationDateTime: ir.NewValidTime(now),
	}

	cases := []struct {
		name    string
		patient *ir.PatientInfo
		pathway []pathway.Allergy
		want    *ir.PatientInfo
	}{
		{
			name:    "Nil allergy - generate random",
			patient: &ir.PatientInfo{},
			pathway: nil,
			want: &ir.PatientInfo{
				Allergies: []*ir.Allergy{
					{
						Type:                   "FA",
						Description:            ir.CodedElement{ID: "J30.1", Text: "Allergy1", CodingSystem: "ZAL"},
						Severity:               "SV",
						Reaction:               "Sneezing",
						IdentificationDateTime: ir.NewValidTime(now),
					},
				},
			},
		}, {
			name:    "Nil allegy - don't generate allergies in already set to empty slice",
			patient: &ir.PatientInfo{Allergies: []*ir.Allergy{}},
			pathway: nil,
			want: &ir.PatientInfo{
				Allergies: []*ir.Allergy{},
			},
		}, {
			name:    "Explicit empty allergy slice in the pathway",
			patient: &ir.PatientInfo{},
			pathway: []pathway.Allergy{},
			want: &ir.PatientInfo{
				Allergies: []*ir.Allergy{},
			},
		}, {
			name: "Explicit empty allergy slice in the pathway - override allergies in the PatientInfo",
			patient: &ir.PatientInfo{
				Allergies: []*ir.Allergy{
					{
						Type:                   "FA",
						Description:            ir.CodedElement{ID: "J30.1", Text: "Allergy1", CodingSystem: "ZAL"},
						Severity:               "SV",
						Reaction:               "Sneezing",
						IdentificationDateTime: ir.NewValidTime(now),
					},
				},
			},
			pathway: []pathway.Allergy{},
			want: &ir.PatientInfo{
				Allergies: []*ir.Allergy{},
			},
		}, {
			name:    "Explicit allergies from the pathway",
			patient: &ir.PatientInfo{},
			pathway: []pathway.Allergy{
				pathwayA1,
				pathwayA2,
			},
			want: &ir.PatientInfo{
				Allergies: []*ir.Allergy{a1, a2},
			},
		}, {
			name:    "Explicit allergies from the pathway - override empty slice",
			patient: &ir.PatientInfo{Allergies: []*ir.Allergy{}},
			pathway: []pathway.Allergy{
				pathwayA1,
			},
			want: &ir.PatientInfo{
				Allergies: []*ir.Allergy{a1},
			},
		}, {
			name:    "Explicit allergies from the pathway with date and coding system",
			patient: &ir.PatientInfo{},
			pathway: []pathway.Allergy{
				{Type: "MEDICATION", Code: "J301", Description: "Allergic rhinitis due to pollen", Severity: "MODERATE", Reaction: "Skin rash", CodingSystem: "all-code", IdentificationDateTime: &pathway.DateTime{TimeFromNow: &twoDaysAgoDuration}},
			},
			want: &ir.PatientInfo{
				Allergies: []*ir.Allergy{
					{
						Type:                   "MEDICATION",
						Description:            ir.CodedElement{ID: "J301", Text: "Allergic rhinitis due to pollen", CodingSystem: "all-code"},
						Severity:               "MODERATE",
						Reaction:               "Skin rash",
						IdentificationDateTime: ir.NewValidTime(twoDaysAgo),
					},
				},
			},
		}, {
			name:    "Explicit allergies from the pathway, no date recorded",
			patient: &ir.PatientInfo{},
			pathway: []pathway.Allergy{
				{Type: "MEDICATION", Code: "J301", Description: "Allergic rhinitis due to pollen", Severity: "MODERATE", Reaction: "Skin rash", IdentificationDateTime: &pathway.DateTime{NoDateTimeRecorded: true}},
			},
			want: &ir.PatientInfo{
				Allergies: []*ir.Allergy{
					{
						Type:                   "MEDICATION",
						Description:            ir.CodedElement{ID: "J301", Text: "Allergic rhinitis due to pollen", CodingSystem: "ZAL"},
						Severity:               "MODERATE",
						Reaction:               "Skin rash",
						IdentificationDateTime: ir.NewInvalidTime(),
					},
				},
			},
		}, {
			name:    "Duplicated allergies in the pathway - deduplicate",
			patient: &ir.PatientInfo{},
			pathway: []pathway.Allergy{
				pathwayA1,
				pathwayA1,
			},
			want: &ir.PatientInfo{
				Allergies: []*ir.Allergy{a1},
			},
		}, {
			name: "Add allergies to existing, deduplicate",
			patient: &ir.PatientInfo{
				Allergies: []*ir.Allergy{a1},
			},
			pathway: []pathway.Allergy{
				pathwayA1,
				pathwayA2,
			},
			want: &ir.PatientInfo{
				Allergies: []*ir.Allergy{a1, a2},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := new(ir.PatientInfo)
			*got = *tc.patient

			g.AddAllergies(got, tc.pathway)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("g.AddAllergies(%+v, %+v) diff: (-want, +got):\n%s", tc.patient, tc.pathway, diff)
			}
		})
	}
}

func TestNewOrder(t *testing.T) {
	ctx := context.Background()
	eventTime := time.Date(2018, 2, 12, 1, 25, 0, 0, time.UTC)
	g := testGenerator(ctx, t, Config{})

	wantOp := "UREA AND ELECTROLYTES"
	p := &pathway.Order{OrderProfile: wantOp}
	got := g.NewOrder(p, eventTime)

	if got == nil {
		t.Errorf("g.NewOrder(%v, %v) = %v, want not nil value", p, eventTime, got)
	}
	if got.OrderProfile.Text != wantOp {
		t.Errorf("g.NewOrder(%v, %v) got OrderProfile Text %q, want %q", p, eventTime, got.OrderProfile.Text, wantOp)
	}
}

func TestOrderWithClinicalNote(t *testing.T) {
	wantContentType := "rtf"
	wantID := "note1.rtf"
	wantTitle := "document-title"

	ctx := context.Background()
	g := testGenerator(ctx, t, Config{})
	eventTime := time.Date(2018, 2, 12, 1, 25, 0, 0, time.UTC)
	pw := &pathway.ClinicalNote{
		ContentType:   wantContentType,
		DocumentID:    wantID,
		DocumentTitle: wantTitle,
	}

	got, err := g.OrderWithClinicalNote(ctx, nil, pw, eventTime)
	if err != nil {
		t.Fatalf("g.OrderWithClinicalNote(nil, %+v, %v) failed with %v", pw, eventTime, err)
	}
	if got == nil {
		t.Errorf("g.OrderWithClinicalNote(nil, %+v, %v) = %v, want not nil value", pw, eventTime, got)
	}
}

func TestSetResults(t *testing.T) {
	ctx := context.Background()
	eventTime := time.Date(2018, 2, 12, 1, 25, 0, 0, time.UTC)
	hl7Config, err := config.LoadHL7Config(ctx, test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	g := testGenerator(ctx, t, Config{HL7Config: hl7Config})

	pathwayR := &pathway.Results{
		OrderProfile: "UREA AND ELECTROLYTES",
		Results: []*pathway.Result{
			{
				TestName: "Creatinine",
				Value:    "52",
				Unit:     "UMOLL",
				Notes:    []string{"Note1", "Note2"},
			},
		},
	}

	order := urineOrder(eventTime, hl7Config)

	got, err := g.SetResults(order, pathwayR, eventTime)
	if err != nil {
		t.Fatalf("SetResults(%v, %v, %v) failed with %v", order, pathwayR, eventTime, err)
	}
	if got == nil {
		t.Errorf("SetResults(%v, %v, %v) = %v, want non-nil value", order, pathwayR, eventTime, got)
	}
}

func TestNewDocument(t *testing.T) {
	ctx := context.Background()
	hl7Config, err := config.LoadHL7Config(ctx, test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	g := testGenerator(ctx, t, Config{HL7Config: hl7Config})
	if got := g.NewDocument(defaultDate, &pathway.Document{}); got == nil {
		t.Errorf("g.NewDocument() = %v, want not nil value", got)
	}
}

func TestNewRegistrationPatientClassAndTypeRandom(t *testing.T) {
	ctx := context.Background()
	rand.Seed(1)

	fPatientClass := testwrite.BytesToFile(t, []byte(`
EMERGENCY,EMERGENCY,1
OUTPATIENT,OUTPATIENT,4
RECURRING,RECURRING,5
`))

	hl7Config, err := config.LoadHL7Config(ctx, test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	f := test.DataFiles[test.Test]
	f.PatientClass = fPatientClass
	dataConfig, err := config.LoadData(ctx, f, hl7Config)
	if err != nil {
		t.Fatalf("LoadData(%+v, %+v) failed with %v", f, hl7Config, err)
	}

	g := testGenerator(ctx, t, Config{Data: dataConfig})

	wantFreq := map[config.PatientClassAndType]int{
		{Class: "EMERGENCY", Type: "EMERGENCY"}:   1,
		{Class: "OUTPATIENT", Type: "OUTPATIENT"}: 4,
		{Class: "RECURRING", Type: "RECURRING"}:   5,
	}

	runs := 1000
	gotFreq := map[config.PatientClassAndType]int{}
	for i := 0; i < runs; i++ {
		got := g.NewRegistrationPatientClassAndType()
		if _, ok := wantFreq[*got]; !ok {
			t.Errorf("g.NewRegistrationPatientClassAndType()=%v, want one of %v", got, wantFreq)
		}
		gotFreq[*got]++
	}

	// Allow an error of 10% of the number of runs.
	delta := float64(runs / 10)
	for k, v := range gotFreq {
		if want := runs * wantFreq[k] / 10; math.Abs(float64(v)-float64(want)) >= delta {
			t.Errorf("wantFreq[%q]=%d, want within %v of %d", k, v, delta, want)
		}
	}
}

func TestNewVisitID(t *testing.T) {
	rand.Seed(1)

	ctx := context.Background()
	g := testGenerator(ctx, t, Config{})

	gotFirst := g.NewVisitID()
	gotSecond := g.NewVisitID()

	if gotFirst == gotSecond {
		t.Errorf("g.NewVisitID() consecutive invocations returned identical value=%d", gotFirst)
	}
}

func TestNewHeader(t *testing.T) {
	rand.Seed(1)

	ctx := context.Background()
	g := testGenerator(ctx, t, Config{})

	pathway := &pathway.Step{Result: &pathway.Results{}}
	got := g.NewHeader(pathway)
	if got == nil {
		t.Errorf("g.NewHeader(%v)=%v, want non-nil value", pathway, got)
	}
}

func TestGeneratorResetPatientInfo(t *testing.T) {
	ctx := context.Background()
	hl7Config, err := config.LoadHL7Config(ctx, test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	g := testGenerator(ctx, t, Config{HL7Config: hl7Config})
	now := ir.NewValidTime(time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC))

	doctor := &ir.Doctor{
		ID:        "id-1",
		Surname:   "surname-1",
		FirstName: "firstname-1",
		Prefix:    "prefix-1",
		Specialty: "specialty-1",
	}
	patient := &state.Patient{
		PatientInfo: &ir.PatientInfo{
			Class:           "INPATIENT",
			Person:          testperson.New(),
			HospitalService: doctor.Specialty,
			AttendingDoctor: doctor,
			PrimaryFacility: &ir.PrimaryFacility{
				Organization: "Test Primary Facility",
				ID:           "123",
			},
			VisitID:   2,
			Location:  &ir.PatientLocation{Poc: "Poc-1", Room: "room-1", Bed: "bed-1"},
			Allergies: []*ir.Allergy{{Type: "food"}},
			Encounters: []*ir.Encounter{
				{
					Status:      constants.EncounterStatusArrived,
					StatusStart: now,
					Start:       now,
					End:         now,
					StatusHistory: []*ir.StatusHistory{
						{
							Status: constants.EncounterStatusPlanned,
							Start:  now,
							End:    now,
						},
					},
				},
			},
		},
		Orders: map[string]*ir.Order{
			"order-id": urineOrder(defaultDate, hl7Config),
		},
		Documents:  map[string]*ir.Document{},
		PastVisits: []uint64{1, 2},
	}

	want := &state.Patient{
		PatientInfo: &ir.PatientInfo{
			Class:           "OUTPATIENT",
			Person:          testperson.New(),
			HospitalService: doctor.Specialty,
			AttendingDoctor: doctor,
			Allergies:       []*ir.Allergy{{Type: "food"}},
			PrimaryFacility: &ir.PrimaryFacility{
				Organization: "Test Primary Facility",
				ID:           "123",
			},
			Encounters: []*ir.Encounter{
				{
					Status:      constants.EncounterStatusArrived,
					StatusStart: now,
					Start:       now,
					End:         now,
					StatusHistory: []*ir.StatusHistory{
						{
							Status: constants.EncounterStatusPlanned,
							Start:  now,
							End:    now,
						},
					},
				},
			},
		},
		Orders: map[string]*ir.Order{
			"order-id": urineOrder(defaultDate, hl7Config),
		},
		Documents:  map[string]*ir.Document{},
		PastVisits: []uint64{1, 2},
	}

	got := g.ResetPatient(patient)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf(`g.ResetPatient(%v) diff (-want, +got):\n%s`, patient, diff)
	}
}

func urineOrder(eventTime time.Time, c *config.HL7Config) *ir.Order {
	return &ir.Order{
		OrderProfile:                  urineElectrolytesCE,
		Placer:                        "12345",
		OrderDateTime:                 ir.NewValidTime(eventTime),
		OrderStatus:                   c.OrderStatus.InProcess,
		MessageControlIDOriginalOrder: "7777",
	}
}

func testGenerator(ctx context.Context, t *testing.T, cfg Config) *Generator {
	t.Helper()
	return testGeneratorWithDate(ctx, t, defaultDate, cfg)
}

func populateConfig(ctx context.Context, t *testing.T, now time.Time, cfg Config) Config {
	t.Helper()
	if cfg.HL7Config == nil {
		c, err := config.LoadHL7Config(ctx, test.MessageConfigTest)
		if err != nil {
			t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
		}
		cfg.HL7Config = c
	}
	if cfg.Header == nil {
		h, err := config.LoadHeaderConfig(ctx, test.HeaderConfigTest)
		if err != nil {
			t.Fatalf("LoadHeaderConfig(%s) failed with %v", test.HeaderConfigTest, err)
		}
		cfg.Header = h
	}
	if cfg.Data == nil {
		f := test.DataFiles[test.Test]
		dataCFG, err := config.LoadData(ctx, f, cfg.HL7Config)
		if err != nil {
			t.Fatalf("LoadData(%+v, %+v) failed with %v", f, cfg.HL7Config, err)
		}
		cfg.Data = dataCFG
	}
	if cfg.Doctors == nil {
		d, err := doctor.LoadDoctors(ctx, test.DoctorsConfigTest)
		if err != nil {
			t.Fatalf("LoadDoctors(%s) failed with %v", test.DoctorsConfigTest, err)
		}
		cfg.Doctors = d
	}
	if cfg.OrderProfiles == nil {
		op, err := orderprofile.Load(ctx, test.OrderProfilesConfigTest, cfg.HL7Config)
		if err != nil {
			t.Fatalf("orderprofile.Load(%s, %+v) failed with %v", test.OrderProfilesConfigTest, cfg.HL7Config, err)
		}
		cfg.OrderProfiles = op
	}
	if cfg.DateGenerator == nil {
		cfg.DateGenerator = &codedelement.SimpleDateGenerator{}
	}
	cfg.AddressGenerator = &testaddress.Generator{Country: "GBR", Cities: defaultCities}
	cfg.MRNGenerator = &testid.Generator{}
	cfg.PlacerGenerator = &testid.Generator{}
	cfg.FillerGenerator = &testid.Generator{}
	cfg.MsgCtrlGenerator = &header.MessageControlGenerator{}
	cfg.Clock = testclock.New(now)
	return cfg
}

func testGeneratorWithDate(ctx context.Context, t *testing.T, now time.Time, cfg Config) *Generator {
	t.Helper()
	cfg = populateConfig(ctx, t, now, cfg)
	return NewGenerator(cfg)
}
