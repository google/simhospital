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
	"github.com/google/simhospital/pkg/message"
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
	urineElectrolytesCE = &message.CodedElement{ID: "lpdc-3969", Text: "UREA AND ELECTROLYTES", CodingSystem: "WinPath"}

	oneDayAgoDuration  = -24 * time.Hour
	twoDaysAgoDuration = -48 * time.Hour

	defaultDate = time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)

	defaultCities = []string{"London", "Cambridge"}
)

func TestNewGeneratorPublicMessageConfiguration(t *testing.T) {
	// If some files are not parsable or accessible, things will crash.
	c, err := config.LoadHL7Config(test.MessageConfigProd)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigProd, err)
	}
	h, err := config.LoadHeaderConfig(test.HeaderConfigProd)
	if err != nil {
		t.Fatalf("LoadHeaderConfig(%s) failed with %v", test.HeaderConfigProd, err)
	}
	f := test.DataFiles[test.Prod]
	dc, err := config.LoadData(f, c)
	if err != nil {
		t.Fatalf("LoadData(%+v, %+v) failed with %v", f, c, err)
	}
	d, err := doctor.LoadDoctors(test.DoctorsConfigProd)
	if err != nil {
		t.Fatalf("LoadDoctors(%s) failed with %v", test.DoctorsConfigProd, err)
	}
	op, err := orderprofile.Load(test.OrderProfilesConfigProd, c)
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
	g := testGenerator(t, Config{})
	pathway := &pathway.Person{}
	got := g.NewPerson(pathway)
	if got == nil {
		t.Errorf("g.NewPerson(%v)=%v, want non-nil value", pathway, got)
	}
}

func originalPatientInfo() *message.PatientInfo {
	return &message.PatientInfo{
		Person:     testperson.New(),
		Diagnoses:  []*message.DiagnosisOrProcedure{},
		Procedures: []*message.DiagnosisOrProcedure{},
		Allergies:  []*message.Allergy{},
	}
}

func TestUpdateFromPathway(t *testing.T) {
	g := testGenerator(t, Config{})

	cases := []struct {
		name    string
		patient *message.PatientInfo
		pathway *pathway.UpdatePerson
		want    func() *message.PatientInfo
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
			want: func() *message.PatientInfo {
				pi := originalPatientInfo()
				pi.Person.MRN = "888888"
				pi.Person.NHS = "9377865972"
				return pi
			},
		}, {
			name: "Update Diagnosis and Procedure: init to empty slices",
			patient: &message.PatientInfo{
				Person:    testperson.New(),
				Allergies: []*message.Allergy{},
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
			want: func() *message.PatientInfo {
				pi := originalPatientInfo()
				pi.Allergies = []*message.Allergy{
					{
						Type:                   "FOOD",
						Description:            message.CodedElement{ID: "E", Text: "egg-containing compound", CodingSystem: "AL"},
						Severity:               "SEVERE",
						Reaction:               "Rash",
						IdentificationDateTime: message.NewValidTime(defaultDate),
					},
				}
				return pi
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := new(message.PatientInfo)
			*got = *tc.patient

			g.UpdateFromPathway(got, tc.pathway)
			if diff := cmp.Diff(tc.want(), got); diff != "" {
				t.Errorf("g.UpdateFromPathway(%+v, %+v) diff: (-want, +got):\n%s", tc.patient, tc.pathway, diff)
			}
		})
	}
}

func TestNewPerson_MRNsAreIncremental(t *testing.T) {
	g := testGenerator(t, Config{})

	person := g.NewPerson(&pathway.Person{})

	patientInfo := &message.PatientInfo{Person: person}
	g.UpdateFromPathway(patientInfo, &pathway.UpdatePerson{Person: &pathway.Person{}})

	g.UpdateFromPathway(patientInfo, &pathway.UpdatePerson{Person: &pathway.Person{MRN: "888888"}})

	// Even if we've updated multiple patients, only NewPerson calls should generate new MRNs.
	person = g.NewPerson(&pathway.Person{})
	if got, want := person.MRN, "2"; got != want {
		t.Errorf("UpdateFromPathway() MRN got %v, want %v", got, want)
	}
}

func TestUpdateFromPathwaySetProceduresAndDiagnoses(t *testing.T) {
	// Allow only one doctor to be "randomly" chosen for testing purposes.
	fName := testwrite.BytesToFile(t, []byte(`
- id: "id-1"
  surname: "surname-1"
  firstname: "firstname-1"
  prefix: "prefix-1"
  specialty: "specialty-1"`))

	d, err := doctor.LoadDoctors(fName)
	if err != nil {
		t.Fatalf("LoadDoctors() failed with %v", err)
	}

	wantDoctor := &message.Doctor{
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

	c, err := config.LoadHL7Config(tmpConfig)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", tmpConfig, err)
	}
	f := test.DataFiles[test.Test]
	f.Procedures = tmpProcedures
	f.Diagnoses = tmpDiagnoses
	dc, err := config.LoadData(f, c)
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
		wantProcedures []*message.DiagnosisOrProcedure
		wantDiagnoses  []*message.DiagnosisOrProcedure
	}{{
		name: "Random, no datetime recorded",
		procedures: []*pathway.DiagnosisOrProcedure{
			{Description: constants.RandomString, DateTime: &pathway.DateTime{NoDateTimeRecorded: true}},
		},
		diagnoses: []*pathway.DiagnosisOrProcedure{
			{Description: constants.RandomString, DateTime: &pathway.DateTime{NoDateTimeRecorded: true}},
		},
		wantProcedures: []*message.DiagnosisOrProcedure{{
			Description: &message.CodedElement{
				ID:           "P24.9",
				Text:         "Procedure1",
				CodingSystem: "PCS",
			},
			Type:      "P",
			Clinician: wantDoctor,
			DateTime:  message.NewInvalidTime(),
		}},
		wantDiagnoses: []*message.DiagnosisOrProcedure{{
			Description: &message.CodedElement{
				ID:           "A01.1",
				Text:         "Diagnosis1",
				CodingSystem: "DCS",
			},
			Type:      "D",
			Clinician: wantDoctor,
			DateTime:  message.NewInvalidTime(),
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
		wantProcedures: []*message.DiagnosisOrProcedure{{
			Type: "A",
			Description: &message.CodedElement{
				ID:   "A01.1",
				Text: "proc1",
			},
			Clinician: wantDoctor,
			DateTime:  message.NewValidTime(defaultDate),
		}},
		wantDiagnoses: []*message.DiagnosisOrProcedure{{
			Type: "B",
			Description: &message.CodedElement{
				ID:   "B01.1",
				Text: "diag1",
			},
			Clinician: wantDoctor,
			DateTime:  message.NewValidTime(defaultDate),
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
		wantProcedures: []*message.DiagnosisOrProcedure{{
			Type: "B",
			Description: &message.CodedElement{
				ID:   "B01.1",
				Text: "proc1",
			},
			Clinician: wantDoctor,
			DateTime:  message.NewValidTime(twoDaysAgo),
		}},
		wantDiagnoses: []*message.DiagnosisOrProcedure{{
			Type: "A",
			Description: &message.CodedElement{
				ID:   "A01.1",
				Text: "diag1",
			},
			Clinician: wantDoctor,
			DateTime:  message.NewValidTime(twoDaysAgo),
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
		wantProcedures: []*message.DiagnosisOrProcedure{{
			Type: "C",
			Description: &message.CodedElement{
				ID:   "C01.1",
				Text: "proc1",
			},
			Clinician: wantDoctor,
			DateTime:  message.NewInvalidTime(),
		}},
		wantDiagnoses: []*message.DiagnosisOrProcedure{{
			Type: "D",
			Description: &message.CodedElement{
				ID:   "D01.1",
				Text: "diag1",
			},
			Clinician: wantDoctor,
			DateTime:  message.NewInvalidTime(),
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
		wantProcedures: []*message.DiagnosisOrProcedure{{
			Type: "A",
			Description: &message.CodedElement{
				ID:   "P24.9",
				Text: "Procedure1",
			},
			Clinician: wantDoctor,
			DateTime:  message.NewValidTime(defaultDate),
		}},
		wantDiagnoses: []*message.DiagnosisOrProcedure{{
			Type: "B",
			Description: &message.CodedElement{
				ID:   "A01.1",
				Text: "Diagnosis1",
			},
			Clinician: wantDoctor,
			DateTime:  message.NewValidTime(defaultDate),
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
		wantProcedures: []*message.DiagnosisOrProcedure{{
			Type: "A",
			Description: &message.CodedElement{
				ID: "XXX.1",
			},
			Clinician: wantDoctor,
			DateTime:  message.NewValidTime(defaultDate),
		}},
		wantDiagnoses: []*message.DiagnosisOrProcedure{{
			Type: "B",
			Description: &message.CodedElement{
				ID: "YYY.2",
			},
			Clinician: wantDoctor,
			DateTime:  message.NewValidTime(defaultDate),
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
		wantProcedures: []*message.DiagnosisOrProcedure{{
			Type: "A",
			Description: &message.CodedElement{
				ID:   "P24.9",
				Text: "Procedure1",
			},
			Clinician: wantDoctor,
			DateTime:  message.NewValidTime(defaultDate),
		}},
		wantDiagnoses: []*message.DiagnosisOrProcedure{{
			Type: "B",
			Description: &message.CodedElement{
				ID:   "A01.1",
				Text: "Diagnosis1",
			},
			Clinician: wantDoctor,
			DateTime:  message.NewValidTime(defaultDate),
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
		wantProcedures: []*message.DiagnosisOrProcedure{{
			Type: "A",
			Description: &message.CodedElement{
				Text: "ProcedureX",
			},
			Clinician: wantDoctor,
			DateTime:  message.NewValidTime(defaultDate),
		}},
		wantDiagnoses: []*message.DiagnosisOrProcedure{{
			Type: "B",
			Description: &message.CodedElement{
				Text: "DiagnosisX",
			},
			Clinician: wantDoctor,
			DateTime:  message.NewValidTime(defaultDate),
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
		wantProcedures: []*message.DiagnosisOrProcedure{{
			Description: &message.CodedElement{
				ID:   "A01.1",
				Text: "proc1",
			},
			Type:      "A",
			Clinician: wantDoctor,
			DateTime:  message.NewValidTime(oneDayAgo),
		}, {
			Description: &message.CodedElement{
				ID:   "B01.1",
				Text: "proc2",
			},
			Type:      "B",
			Clinician: wantDoctor,
			DateTime:  message.NewValidTime(twoDaysAgo),
		}},
		wantDiagnoses: []*message.DiagnosisOrProcedure{{
			Description: &message.CodedElement{
				ID:   "C01.1",
				Text: "diag1",
			},
			Type:      "C",
			Clinician: wantDoctor,
			DateTime:  message.NewValidTime(oneDayAgo),
		}, {
			Description: &message.CodedElement{
				ID:   "D01.1",
				Text: "diag2",
			},
			Type:      "D",
			Clinician: wantDoctor,
			DateTime:  message.NewValidTime(twoDaysAgo),
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
			g := testGeneratorWithDate(t, now, Config{HL7Config: c, Data: dc, Doctors: d, DateGenerator: dg})
			g.UpdateFromPathway(got, pathway)

			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("g.UpdateFromPathway(%+v, %+v) diff: (-want, +got):\n%s", originalPatientInfo(), pathway, diff)
			}
		})
	}
}

func TestNewPatient(t *testing.T) {
	dName := testwrite.BytesToFile(t, []byte(`
- id: "id-1"
  surname: "surname-1"
  firstname: "firstname-1"
  prefix: "prefix-1"
  specialty: "specialty-1"`))

	d, err := doctor.LoadDoctors(dName)
	if err != nil {
		t.Fatalf("LoadDoctors(%s) failed with %v", dName, err)
	}

	existingDoctor := &message.Doctor{
		ID:        "id-1",
		Surname:   "surname-1",
		FirstName: "firstname-1",
		Prefix:    "prefix-1",
		Specialty: "specialty-1",
	}

	// This empty configuration file simulates a nil primary facility and nil patient class.
	nilHL7Name := testwrite.BytesToFile(t, []byte(``))

	nilHL7Config, err := config.LoadHL7Config(nilHL7Name)
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

	hl7Config, err := config.LoadHL7Config(hl7Name)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", hl7Name, err)
	}

	newDoctor := &message.Doctor{
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
		doctor *message.Doctor
		want   *state.Patient
	}{
		{
			name:   "Nil doctor, primary facility, hospital service and patient class not specified in config",
			conf:   Config{HL7Config: nilHL7Config, Doctors: d},
			doctor: nil,
			want: &state.Patient{
				PatientInfo: &message.PatientInfo{
					Class:           "",
					Person:          person,
					HospitalService: "",
				},
				Orders: make(map[string]*message.Order),
			},
		}, {
			name:   "Existing doctor, override hospital service",
			conf:   Config{HL7Config: nilHL7Config, Doctors: d},
			doctor: existingDoctor,
			want: &state.Patient{
				PatientInfo: &message.PatientInfo{
					Class:           "",
					Person:          person,
					HospitalService: existingDoctor.Specialty,
					AttendingDoctor: existingDoctor,
				},
				Orders: make(map[string]*message.Order),
			},
		}, {
			name:   "New doctor, don't override hospital service",
			conf:   Config{HL7Config: nilHL7Config, Doctors: d},
			doctor: newDoctor,
			want: &state.Patient{
				PatientInfo: &message.PatientInfo{
					Class:           "",
					Person:          person,
					HospitalService: "",
					AttendingDoctor: newDoctor,
				},
				Orders: make(map[string]*message.Order),
			},
		}, {
			name:   "Nil doctor, primary facility, hospital service and patient class from config",
			conf:   Config{HL7Config: hl7Config, Doctors: d},
			doctor: newDoctor,
			want: &state.Patient{
				PatientInfo: &message.PatientInfo{
					Class:           "OUTPATIENT",
					Person:          person,
					HospitalService: "180",
					AttendingDoctor: newDoctor,
					PrimaryFacility: &message.PrimaryFacility{
						Organization: "Test Primary Facility",
						ID:           "123",
					},
				},
				Orders: make(map[string]*message.Order),
			},
		}, {
			name:   "Existing doctor, defined config, override hospital service",
			conf:   Config{HL7Config: hl7Config, Doctors: d},
			doctor: existingDoctor,
			want: &state.Patient{
				PatientInfo: &message.PatientInfo{
					Class:           "OUTPATIENT",
					Person:          person,
					HospitalService: existingDoctor.Specialty,
					AttendingDoctor: existingDoctor,
					PrimaryFacility: &message.PrimaryFacility{
						Organization: "Test Primary Facility",
						ID:           "123",
					},
				},
				Orders: make(map[string]*message.Order),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := testGenerator(t, tc.conf)

			got := g.NewPatient(person, tc.doctor)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("g.NewPatient(%+v, %+v) diff: (-want, +got):\n%s", person, tc.doctor, diff)
			}
		})
	}
}

func TestNewDoctor(t *testing.T) {
	hl7Name := testwrite.BytesToFile(t, []byte(`
hospital_service: "180"
`))

	hl7Config, err := config.LoadHL7Config(hl7Name)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", hl7Name, err)
	}

	dName := testwrite.BytesToFile(t, []byte(`
- id: "id-1"
  surname: "surname-1"
  firstname: "firstname-1"
  prefix: "prefix-1"
  specialty: "specialty-1"`))

	existingDoctor := &message.Doctor{
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
		want       *message.Doctor
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
			want: &message.Doctor{
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
			d, err := doctor.LoadDoctors(dName)
			if err != nil {
				t.Fatalf("LoadDoctors(%s) failed with %v", dName, err)
			}
			g := testGenerator(t, Config{HL7Config: hl7Config, Doctors: d})
			got := g.NewDoctor(tc.consultant)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("g.NewDoctor(%+v) diff: (-want, +got):\n%s", tc.consultant, diff)
			}
		})
	}
}

func TestNewDoctorAddDoctorsForFutureUse(t *testing.T) {
	// No doctors defined.
	dName := testwrite.BytesToFile(t, []byte(``))
	d, err := doctor.LoadDoctors(dName)
	if err != nil {
		t.Fatalf("LoadDoctors(%s) failed with %v", dName, err)
	}
	hl7Name := testwrite.BytesToFile(t, []byte(`
hospital_service: "180"
`))

	hl7Config, err := config.LoadHL7Config(hl7Name)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", hl7Name, err)
	}

	newID := "123"
	newSurname := "Osman"
	newName := "Arthur"
	newPrefix := "Dr"

	want := &message.Doctor{
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

	g := testGenerator(t, Config{HL7Config: hl7Config, Doctors: d})
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

	hl7Config, err := config.LoadHL7Config(fHL7Config)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", fHL7Config, err)
	}
	f := test.DataFiles[test.Test]
	f.Allergies = fAllergies
	f.DataConfig = fData
	dataConfig, err := config.LoadData(f, hl7Config)
	if err != nil {
		t.Fatalf("LoadData(%+v, %+v) failed with %v", f, hl7Config, err)
	}

	dg := &testdate.Generator{}

	g := testGeneratorWithDate(t, now, Config{HL7Config: hl7Config, Data: dataConfig, DateGenerator: dg, Clock: testclock.New(now)})

	pathwayA1 := pathway.Allergy{Type: "MEDICATION", Code: "J301", Description: "Allergic rhinitis due to pollen", Severity: "MODERATE", Reaction: "Skin rash"}
	pathwayA2 := pathway.Allergy{Type: "FOOD", Code: "E", Description: "egg-containing compound", Severity: "SEVERE", Reaction: "Rash"}

	a1 := &message.Allergy{
		Type:                   "MEDICATION",
		Description:            message.CodedElement{ID: "J301", Text: "Allergic rhinitis due to pollen", CodingSystem: "ZAL"},
		Severity:               "MODERATE",
		Reaction:               "Skin rash",
		IdentificationDateTime: message.NewValidTime(now),
	}
	a2 := &message.Allergy{
		Type:                   "FOOD",
		Description:            message.CodedElement{ID: "E", Text: "egg-containing compound", CodingSystem: "ZAL"},
		Severity:               "SEVERE",
		Reaction:               "Rash",
		IdentificationDateTime: message.NewValidTime(now),
	}

	cases := []struct {
		name    string
		patient *message.PatientInfo
		pathway []pathway.Allergy
		want    *message.PatientInfo
	}{
		{
			name:    "Nil allergy - generate random",
			patient: &message.PatientInfo{},
			pathway: nil,
			want: &message.PatientInfo{
				Allergies: []*message.Allergy{
					{
						Type:                   "FA",
						Description:            message.CodedElement{ID: "J30.1", Text: "Allergy1", CodingSystem: "ZAL"},
						Severity:               "SV",
						Reaction:               "Sneezing",
						IdentificationDateTime: message.NewValidTime(now),
					},
				},
			},
		}, {
			name:    "Nil allegy - don't generate allergies in already set to empty slice",
			patient: &message.PatientInfo{Allergies: []*message.Allergy{}},
			pathway: nil,
			want: &message.PatientInfo{
				Allergies: []*message.Allergy{},
			},
		}, {
			name:    "Explicit empty allergy slice in the pathway",
			patient: &message.PatientInfo{},
			pathway: []pathway.Allergy{},
			want: &message.PatientInfo{
				Allergies: []*message.Allergy{},
			},
		}, {
			name: "Explicit empty allergy slice in the pathway - override allergies in the PatientInfo",
			patient: &message.PatientInfo{
				Allergies: []*message.Allergy{
					{
						Type:                   "FA",
						Description:            message.CodedElement{ID: "J30.1", Text: "Allergy1", CodingSystem: "ZAL"},
						Severity:               "SV",
						Reaction:               "Sneezing",
						IdentificationDateTime: message.NewValidTime(now),
					},
				},
			},
			pathway: []pathway.Allergy{},
			want: &message.PatientInfo{
				Allergies: []*message.Allergy{},
			},
		}, {
			name:    "Explicit allergies from the pathway",
			patient: &message.PatientInfo{},
			pathway: []pathway.Allergy{
				pathwayA1,
				pathwayA2,
			},
			want: &message.PatientInfo{
				Allergies: []*message.Allergy{a1, a2},
			},
		}, {
			name:    "Explicit allergies from the pathway - override empty slice",
			patient: &message.PatientInfo{Allergies: []*message.Allergy{}},
			pathway: []pathway.Allergy{
				pathwayA1,
			},
			want: &message.PatientInfo{
				Allergies: []*message.Allergy{a1},
			},
		}, {
			name:    "Explicit allergies from the pathway with date and coding system",
			patient: &message.PatientInfo{},
			pathway: []pathway.Allergy{
				{Type: "MEDICATION", Code: "J301", Description: "Allergic rhinitis due to pollen", Severity: "MODERATE", Reaction: "Skin rash", CodingSystem: "all-code", IdentificationDateTime: &pathway.DateTime{TimeFromNow: &twoDaysAgoDuration}},
			},
			want: &message.PatientInfo{
				Allergies: []*message.Allergy{
					{
						Type:                   "MEDICATION",
						Description:            message.CodedElement{ID: "J301", Text: "Allergic rhinitis due to pollen", CodingSystem: "all-code"},
						Severity:               "MODERATE",
						Reaction:               "Skin rash",
						IdentificationDateTime: message.NewValidTime(twoDaysAgo),
					},
				},
			},
		}, {
			name:    "Explicit allergies from the pathway, no date recorded",
			patient: &message.PatientInfo{},
			pathway: []pathway.Allergy{
				{Type: "MEDICATION", Code: "J301", Description: "Allergic rhinitis due to pollen", Severity: "MODERATE", Reaction: "Skin rash", IdentificationDateTime: &pathway.DateTime{NoDateTimeRecorded: true}},
			},
			want: &message.PatientInfo{
				Allergies: []*message.Allergy{
					{
						Type:                   "MEDICATION",
						Description:            message.CodedElement{ID: "J301", Text: "Allergic rhinitis due to pollen", CodingSystem: "ZAL"},
						Severity:               "MODERATE",
						Reaction:               "Skin rash",
						IdentificationDateTime: message.NewInvalidTime(),
					},
				},
			},
		}, {
			name:    "Duplicated allergies in the pathway - deduplicate",
			patient: &message.PatientInfo{},
			pathway: []pathway.Allergy{
				pathwayA1,
				pathwayA1,
			},
			want: &message.PatientInfo{
				Allergies: []*message.Allergy{a1},
			},
		}, {
			name: "Add allergies to existing, deduplicate",
			patient: &message.PatientInfo{
				Allergies: []*message.Allergy{a1},
			},
			pathway: []pathway.Allergy{
				pathwayA1,
				pathwayA2,
			},
			want: &message.PatientInfo{
				Allergies: []*message.Allergy{a1, a2},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := new(message.PatientInfo)
			*got = *tc.patient

			g.AddAllergies(got, tc.pathway)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("g.AddAllergies(%+v, %+v) diff: (-want, +got):\n%s", tc.patient, tc.pathway, diff)
			}
		})
	}
}

func TestNewOrder(t *testing.T) {
	eventTime := time.Date(2018, 2, 12, 1, 25, 0, 0, time.UTC)
	g := testGenerator(t, Config{})

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

	g := testGenerator(t, Config{})
	eventTime := time.Date(2018, 2, 12, 1, 25, 0, 0, time.UTC)
	pw := &pathway.ClinicalNote{
		ContentType:   wantContentType,
		DocumentID:    wantID,
		DocumentTitle: wantTitle,
	}

	got, err := g.OrderWithClinicalNote(nil, pw, eventTime)
	if err != nil {
		t.Fatalf("g.OrderWithClinicalNote(nil, %+v, %v) failed with %v", pw, eventTime, err)
	}
	if got == nil {
		t.Errorf("g.OrderWithClinicalNote(nil, %+v, %v) = %v, want not nil value", pw, eventTime, got)
	}
}

func TestSetResults(t *testing.T) {
	eventTime := time.Date(2018, 2, 12, 1, 25, 0, 0, time.UTC)
	hl7Config, err := config.LoadHL7Config(test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	g := testGenerator(t, Config{HL7Config: hl7Config})

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
	hl7Config, err := config.LoadHL7Config(test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	g := testGenerator(t, Config{HL7Config: hl7Config})
	if got := g.NewDocument(defaultDate, &pathway.Document{}); got == nil {
		t.Errorf("g.NewDocument() = %v, want not nil value", got)
	}
}

func TestNewRegistrationPatientClassAndTypeRandom(t *testing.T) {
	rand.Seed(1)

	fPatientClass := testwrite.BytesToFile(t, []byte(`
EMERGENCY,EMERGENCY,1
OUTPATIENT,OUTPATIENT,4
RECURRING,RECURRING,5
`))

	hl7Config, err := config.LoadHL7Config(test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	f := test.DataFiles[test.Test]
	f.PatientClass = fPatientClass
	dataConfig, err := config.LoadData(f, hl7Config)
	if err != nil {
		t.Fatalf("LoadData(%+v, %+v) failed with %v", f, hl7Config, err)
	}

	g := testGenerator(t, Config{Data: dataConfig})

	wantFreq := map[config.PatientClassAndType]int{
		config.PatientClassAndType{Class: "EMERGENCY", Type: "EMERGENCY"}:   1,
		config.PatientClassAndType{Class: "OUTPATIENT", Type: "OUTPATIENT"}: 4,
		config.PatientClassAndType{Class: "RECURRING", Type: "RECURRING"}:   5,
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

	g := testGenerator(t, Config{})

	gotFirst := g.NewVisitID()
	gotSecond := g.NewVisitID()

	if gotFirst == gotSecond {
		t.Errorf("g.NewVisitID() consecutive invocations returned identical value=%d", gotFirst)
	}
}

func TestNewHeader(t *testing.T) {
	rand.Seed(1)

	g := testGenerator(t, Config{})

	pathway := &pathway.Step{Result: &pathway.Results{}}
	got := g.NewHeader(pathway)
	if got == nil {
		t.Errorf("g.NewHeader(%v)=%v, want non-nil value", pathway, got)
	}
}

func TestGeneratorResetPatientInfo(t *testing.T) {
	hl7Config, err := config.LoadHL7Config(test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	g := testGenerator(t, Config{HL7Config: hl7Config})

	doctor := &message.Doctor{
		ID:        "id-1",
		Surname:   "surname-1",
		FirstName: "firstname-1",
		Prefix:    "prefix-1",
		Specialty: "specialty-1",
	}
	patient := &state.Patient{
		PatientInfo: &message.PatientInfo{
			Class:           "INPATIENT",
			Person:          testperson.New(),
			HospitalService: doctor.Specialty,
			AttendingDoctor: doctor,
			PrimaryFacility: &message.PrimaryFacility{
				Organization: "Test Primary Facility",
				ID:           "123",
			},
			VisitID:   2,
			Location:  &message.PatientLocation{Poc: "Poc-1", Room: "room-1", Bed: "bed-1"},
			Allergies: []*message.Allergy{{Type: "food"}},
		},
		Orders: map[string]*message.Order{
			"order-id": urineOrder(defaultDate, hl7Config),
		},
		PastVisits: []uint64{1, 2},
	}

	want := &state.Patient{
		PatientInfo: &message.PatientInfo{
			Class:           "OUTPATIENT",
			Person:          testperson.New(),
			HospitalService: doctor.Specialty,
			AttendingDoctor: doctor,
			PrimaryFacility: &message.PrimaryFacility{
				Organization: "Test Primary Facility",
				ID:           "123",
			},
		},
		Orders: map[string]*message.Order{
			"order-id": urineOrder(defaultDate, hl7Config),
		},
		PastVisits: []uint64{1, 2},
	}

	got := g.ResetPatient(patient)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf(`g.ResetPatient(%v) diff (-want, +got):\n%s`, patient, diff)
	}
}

func urineOrder(eventTime time.Time, c *config.HL7Config) *message.Order {
	return &message.Order{
		OrderProfile:                  urineElectrolytesCE,
		Placer:                        "12345",
		OrderDateTime:                 message.NewValidTime(eventTime),
		OrderStatus:                   c.OrderStatus.InProcess,
		MessageControlIDOriginalOrder: "7777",
	}
}

func testGenerator(t *testing.T, cfg Config) *Generator {
	t.Helper()
	return testGeneratorWithDate(t, defaultDate, cfg)
}

func populateConfig(t *testing.T, now time.Time, cfg Config) Config {
	t.Helper()
	if cfg.HL7Config == nil {
		c, err := config.LoadHL7Config(test.MessageConfigTest)
		if err != nil {
			t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
		}
		cfg.HL7Config = c
	}
	if cfg.Header == nil {
		h, err := config.LoadHeaderConfig(test.HeaderConfigTest)
		if err != nil {
			t.Fatalf("LoadHeaderConfig(%s) failed with %v", test.HeaderConfigTest, err)
		}
		cfg.Header = h
	}
	if cfg.Data == nil {
		f := test.DataFiles[test.Test]
		dataCFG, err := config.LoadData(f, cfg.HL7Config)
		if err != nil {
			t.Fatalf("LoadData(%+v, %+v) failed with %v", f, cfg.HL7Config, err)
		}
		cfg.Data = dataCFG
	}
	if cfg.Doctors == nil {
		d, err := doctor.LoadDoctors(test.DoctorsConfigTest)
		if err != nil {
			t.Fatalf("LoadDoctors(%s) failed with %v", test.DoctorsConfigTest, err)
		}
		cfg.Doctors = d
	}
	if cfg.OrderProfiles == nil {
		op, err := orderprofile.Load(test.OrderProfilesConfigTest, cfg.HL7Config)
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

func testGeneratorWithDate(t *testing.T, now time.Time, cfg Config) *Generator {
	t.Helper()
	cfg = populateConfig(t, now, cfg)
	return NewGenerator(cfg)
}
