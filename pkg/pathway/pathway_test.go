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

package pathway

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/constants"
	"github.com/google/simhospital/pkg/orderprofile"
	"github.com/google/simhospital/pkg/test/testclock"
)

func TestGender(t *testing.T) {
	cases := []struct {
		g    Gender
		want bool
	}{
		{g: Male, want: true},
		{g: Female, want: true},
		{g: "invalid", want: true},
		{g: "", want: false},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.g), func(t *testing.T) {
			if got := tc.g.IsSet(); got != tc.want {
				t.Errorf("[%v].IsSet()=%t, want %t", tc.g, got, tc.want)
			}
		})
	}
}

func TestOptionalRandomString(t *testing.T) {
	cases := []struct {
		s           OptionalRandomString
		wantIsSet   bool
		wantIsFixed bool
		wantString  string
	}{
		{s: "", wantIsSet: false, wantIsFixed: false, wantString: ""},
		{s: constants.RandomString, wantIsSet: true, wantIsFixed: false, wantString: constants.RandomString},
		{s: "value", wantIsSet: true, wantIsFixed: true, wantString: "value"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.s), func(t *testing.T) {
			if got := tc.s.IsSet(); got != tc.wantIsSet {
				t.Errorf("[%v].IsSet()=%t, want %t", tc.s, got, tc.wantIsSet)
			}
			if got := tc.s.IsFixedValue(); got != tc.wantIsFixed {
				t.Errorf("[%v].IsFixedValue()=%t, want %t", tc.s, got, tc.wantIsFixed)
			}
			if got := tc.s.String(); got != tc.wantString {
				t.Errorf("[%v].String()=%s, want %s", tc.s, got, tc.wantString)
			}
		})
	}
}

func TestAgeBirthdate(t *testing.T) {
	now := time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC)
	clock := testclock.New(now)
	cases := []struct {
		name         string
		age          *Age
		wantYearFrom int
		wantYearTo   int
	}{
		{
			name:         "20-30",
			age:          &Age{From: 20, To: 30},
			wantYearFrom: 1988,
			wantYearTo:   1998,
		},
		{
			name:         "30",
			age:          &Age{From: 30, To: 30},
			wantYearFrom: 1988,
			wantYearTo:   1988,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.age.Birthdate(clock)
			if got.Year() < tc.wantYearFrom || got.Year() > tc.wantYearTo {
				t.Errorf("(%+v).Birthdate(%+v)=%v, want year between (%d, %d)", tc.age, clock, got, tc.wantYearFrom, tc.wantYearTo)
			}
		})
	}
}

func TestBirthdateDayOfYear(t *testing.T) {
	now := time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC)
	clock := testclock.New(now)

	cases := []struct {
		name      string
		dayOfYear int
		wantMonth int
		wantDay   int
	}{
		// 23rd day of year should be January 23.
		{
			name:      "Jan 23",
			dayOfYear: 23,
			wantMonth: 1,
			wantDay:   23,
		},
		// 42nd day of year should be February 11.
		{
			name:      "Feb 11",
			dayOfYear: 42,
			wantMonth: 2,
			wantDay:   11,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			age := &Age{DayOfYear: tc.dayOfYear}
			dob := age.Birthdate(clock)
			if got, want := int(dob.Month()), tc.wantMonth; got != want {
				t.Errorf("(%+v).Birthdate(_).Month()=%v, want %v", age, got, want)
			}
			if got, want := dob.Day(), tc.wantDay; got != want {
				t.Errorf("(%+v).Birthdate(_).Day()=%v, want %v", age, got, want)
			}
		})
	}
}

func TestRandomBirthdate(t *testing.T) {
	now := time.Date(2018, 2, 12, 0, 0, 0, 0, time.UTC)
	clock := testclock.New(now)

	wantYearFrom := now.Year() - 100
	wantYearTo := now.Year() - 1
	got := RandomBirthdate(clock)
	if got.Year() < wantYearFrom || got.Year() > wantYearTo {
		t.Errorf("RandomBirthdate(%+v)=%v, want year between (%d, %d)", clock, got, wantYearFrom, wantYearTo)
	}
}

func TestResultGetValueTypeAndUnit(t *testing.T) {
	cases := []struct {
		name      string
		r         *Result
		wantValue string
		wantType  string
		wantUnit  string
	}{
		{
			name:      "numerical value",
			r:         &Result{TestName: "Creatinine", Value: "12", Unit: "UMOLL"},
			wantValue: "12",
			wantType:  "NM",
			wantUnit:  "UMOLL",
		}, {
			name:      "textual value",
			r:         &Result{TestName: "Creatinine", Value: "the value"},
			wantValue: "the value",
			wantType:  "TX",
			wantUnit:  "",
		}, {
			name:      "empty value",
			r:         &Result{TestName: "Creatinine", Value: constants.EmptyString, Unit: constants.EmptyString},
			wantValue: "",
			wantType:  "",
			wantUnit:  "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.r.GetValue(); got != tc.wantValue {
				t.Errorf("[%+v].GetValue()=%v, want %v", tc.r, got, tc.wantValue)
			}
			if got := tc.r.GetValueType(); got != tc.wantType {
				t.Errorf("[%+v].GetValueType()=%v, want %v", tc.r, got, tc.wantType)
			}
			if got := tc.r.GetUnit(); got != tc.wantUnit {
				t.Errorf("[%+v].GetUnit()=%v, want %v", tc.r, got, tc.wantUnit)
			}
		})
	}
}

func TestResultGetAbnormalFlag(t *testing.T) {
	secondaryValueGenerator, err := orderprofile.ValueGeneratorFromRange("60-80")
	if err != nil {
		t.Fatalf("ValueGeneratorFromRange() failed with %v", err)
	}
	cases := []struct {
		name        string
		r           *Result
		secondaryVG *orderprofile.ValueGenerator
		want        constants.AbnormalFlag
	}{
		{
			name:        "explicit empty flag",
			r:           &Result{TestName: "Creatinine", AbnormalFlag: constants.AbnormalFlagEmpty},
			secondaryVG: secondaryValueGenerator,
			want:        constants.AbnormalFlagEmpty,
		}, {
			name:        "explicit normal flag",
			r:           &Result{TestName: "Creatinine", AbnormalFlag: constants.AbnormalFlagNormal},
			secondaryVG: secondaryValueGenerator,
			want:        constants.AbnormalFlagEmpty,
		}, {
			name:        "explicit high flag",
			r:           &Result{TestName: "Creatinine", AbnormalFlag: constants.AbnormalFlagHigh},
			secondaryVG: secondaryValueGenerator,
			want:        constants.AbnormalFlagHigh,
		}, {
			name:        "explicit low flag",
			r:           &Result{TestName: "Creatinine", AbnormalFlag: constants.AbnormalFlagLow},
			secondaryVG: secondaryValueGenerator,
			want:        constants.AbnormalFlagLow,
		}, {
			name:        "low flag from secondary generator",
			r:           &Result{TestName: "Creatinine", Value: "12", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagDefault},
			secondaryVG: secondaryValueGenerator,
			want:        constants.AbnormalFlagLow,
		}, {
			name:        "normal flag from secondary generator",
			r:           &Result{TestName: "Creatinine", Value: "70", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagDefault},
			secondaryVG: secondaryValueGenerator,
			want:        constants.AbnormalFlagEmpty,
		}, {
			name:        "high flag from secondary generator",
			r:           &Result{TestName: "Creatinine", Value: "90", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagDefault},
			secondaryVG: secondaryValueGenerator,
			want:        constants.AbnormalFlagHigh,
		}, {
			name:        "low flag from reference range",
			r:           &Result{TestName: "Creatinine", Value: "12", Unit: "UMOLL", ReferenceRange: "20 - 40", AbnormalFlag: constants.AbnormalFlagDefault},
			secondaryVG: nil,
			want:        constants.AbnormalFlagLow,
		}, {
			name:        "normal flag from reference range",
			r:           &Result{TestName: "Creatinine", Value: "25", Unit: "UMOLL", ReferenceRange: "20 - 40", AbnormalFlag: constants.AbnormalFlagDefault},
			secondaryVG: nil,
			want:        constants.AbnormalFlagEmpty,
		}, {
			name:        "high flag from reference range",
			r:           &Result{TestName: "Creatinine", Value: "45", Unit: "UMOLL", ReferenceRange: "20 - 40", AbnormalFlag: constants.AbnormalFlagDefault},
			secondaryVG: nil,
			want:        constants.AbnormalFlagHigh,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.r.GetAbnormalFlag(tc.secondaryVG)
			if err != nil {
				t.Fatalf("[%+v].GetAbnormalFlag(%v) failed with %v", tc.r, tc.secondaryVG, err)
			}
			if got != tc.want {
				t.Errorf("[%+v].GetAbnormalFlag(%v)=%v, want %v", tc.r, tc.secondaryVG, got, tc.want)
			}
		})
	}
}

func TestResultGetRandomType(t *testing.T) {
	cases := []struct {
		name    string
		r       *Result
		want    string
		wantErr bool
	}{
		{name: "nil Result - normal value", r: nil, want: constants.NormalValue},
		{name: "normal value", r: &Result{Value: constants.NormalValue}, want: constants.NormalValue},
		{name: "high value", r: &Result{Value: constants.AbnormalHigh}, want: constants.AbnormalHigh},
		{name: "low value", r: &Result{Value: constants.AbnormalLow}, want: constants.AbnormalLow},
		{name: "invalid", r: &Result{Value: "invalid"}, wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.r.GetRandomType()
			if (err != nil) != tc.wantErr {
				t.Fatalf("[%+v].GetRandomType got err %v, want err? %t", tc.r, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("[%+v].GetRandomType() got %s, want %s", tc.r, got, tc.want)
			}
		})
	}
}

func TestResultIsValueRandom(t *testing.T) {
	cases := []struct {
		value string
		want  bool
	}{
		{value: constants.NormalValue, want: true},
		{value: constants.AbnormalHigh, want: true},
		{value: constants.AbnormalLow, want: true},
		{value: "12.5", want: false},
		{value: "Normal value", want: false},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s-random:%t", tc.value, tc.want), func(t *testing.T) {
			r := Result{Value: tc.value}
			if got := r.IsValueRandom(); got != tc.want {
				t.Errorf("[%+v].IsValueRandom()=%v, want %v", r, got, tc.want)
			}
		})
	}
}

func TestStepStepType(t *testing.T) {
	cases := []struct {
		step Step
		want string
	}{
		{step: Step{Delay: &Delay{}}, want: StepDelay},
		{step: Step{Delay: &Delay{}, Parameters: &Parameters{}}, want: StepDelay},
		{step: Step{Admission: &Admission{}}, want: StepAdmission},
		{step: Step{Order: &Order{}}, want: StepOrder},
		{step: Step{Result: &Results{}}, want: StepResults},
		{step: Step{Transfer: &Transfer{}}, want: StepTransfer},
		{step: Step{Discharge: &Discharge{}}, want: StepDischarge},
		{step: Step{Registration: &Registration{}}, want: StepRegistration},
		{step: Step{AddPerson: &AddPerson{}}, want: StepAddPerson},
		{step: Step{UpdatePerson: &UpdatePerson{}}, want: StepUpdatePerson},
		{step: Step{Merge: &Merge{}}, want: StepMerge},
		{step: Step{BedSwap: &BedSwap{}}, want: StepBedSwap},
		{step: Step{TransferInError: &TransferInError{}}, want: StepTransferInError},
		{step: Step{DischargeInError: &DischargeInError{}}, want: StepDischargeInError},
		{step: Step{CancelVisit: &CancelVisit{}}, want: StepCancelVisit},
		{step: Step{CancelTransfer: &CancelTransfer{}}, want: StepCancelTransfer},
		{step: Step{CancelDischarge: &CancelDischarge{}}, want: StepCancelDischarge},
		{step: Step{PendingAdmission: &PendingAdmission{}}, want: StepPendingAdmission},
		{step: Step{PendingDischarge: &PendingDischarge{}}, want: StepPendingDischarge},
		{step: Step{PendingTransfer: &PendingTransfer{}}, want: StepPendingTransfer},
		{step: Step{CancelPendingTransfer: &CancelPendingTransfer{}}, want: StepCancelPendingTransfer},
		{step: Step{CancelPendingAdmission: &CancelPendingAdmission{}}, want: StepCancelPendingAdmission},
		{step: Step{CancelPendingDischarge: &CancelPendingDischarge{}}, want: StepCancelPendingDischarge},
		{step: Step{TrackDeparture: &TrackDeparture{}}, want: StepTrackDeparture},
		{step: Step{TrackArrival: &TrackArrival{}}, want: StepTrackArrival},
		{step: Step{UsePatient: &UsePatient{}}, want: StepUsePatient},
		{step: Step{Delay: &Delay{}, Admission: &Admission{}}, want: stepInvalid},
		{step: Step{Parameters: &Parameters{}}, want: stepInvalid},
		{step: Step{AutoGenerate: &AutoGenerate{}}, want: StepAutoGenerate},
		{step: Step{ClinicalNote: &ClinicalNote{}}, want: StepClinicalNote},
		{step: Step{HardcodedMessage: &HardcodedMessage{}}, want: StepHardcodedMessage},
		{step: Step{Document: &Document{}}, want: StepDocument},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.want), func(t *testing.T) {
			if got := tc.step.StepType(); got != tc.want {
				t.Errorf("[%+v].StepType()=%v, want %v", tc.step, got, tc.want)
			}
		})
	}
}

// The TestStepTypeIsMemoizedValid tests make sure the StepType is calculated only once. This test is
// not realistic as fields in the Steps should not be set manually but it's the only way to test
// the type is saved on the first invocation and not recalculated in subsequent invocations.
func TestStepTypeIsMemoizedValid(t *testing.T) {
	// A step with both Delay and TrackArrival set is invalid.
	invalid := &Step{Delay: &Delay{}, TrackArrival: &TrackArrival{}}
	if got, want := invalid.StepType(), stepInvalid; got != want {
		t.Errorf("invalid.StepType()=%v, want %v", got, want)
	}

	// A delay only is valid.
	delay := &Step{Delay: &Delay{}}
	if got, want := delay.StepType(), StepDelay; got != want {
		t.Errorf("delay.StepType()=%v, want %v", got, want)
	}
	// If we now add TrackArrival, the type shouldn't change, so it should stay as Delay.
	delay.TrackArrival = &TrackArrival{}
	if got, want := delay.StepType(), StepDelay; got != want {
		t.Errorf("delay.StepType()=%v, want %v", got, want)
	}
}

func TestStepTypeIsMemoizedInvalid(t *testing.T) {
	// A step with Delay and Parameters is valid.
	validDelay := &Step{Delay: &Delay{}, Parameters: &Parameters{}}
	if got, want := validDelay.StepType(), StepDelay; got != want {
		t.Errorf("validDelay.StepType()=%v, want %v", got, want)
	}

	// A step with parameters only is invalid.
	invalid := &Step{Parameters: &Parameters{}}
	if got, want := invalid.StepType(), stepInvalid; got != want {
		t.Errorf("invalid.StepType()=%v, want %v", got, want)
	}

	// If we now add a Delay it would have been valid, but the type is kept as Invalid.
	invalid.Delay = &Delay{}
	if got, want := invalid.StepType(), stepInvalid; got != want {
		t.Errorf("invalid.StepType()=%v, want %v", got, want)
	}
}

func TestOnlyPerson(t *testing.T) {
	arbitraryPerson := Person{FirstName: "first-name"}
	cases := []struct {
		name    string
		p       *Persons
		want    *Person
		wantID  PatientID
		wantErr bool
	}{
		{name: "nil Persons", p: nil, wantID: "", want: nil, wantErr: true},
		{name: "too many Persons", p: &Persons{"1": {}, "2": {}}, wantID: "", want: nil, wantErr: true},
		{name: "exactly one Person", p: &Persons{"1": arbitraryPerson}, want: &arbitraryPerson, wantID: "1", wantErr: false},
		{name: "exactly one empty Person", p: &Persons{"1": {}}, want: &Person{}, wantID: "1", wantErr: false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotID, got, err := tc.p.OnlyPerson()
			gotErr := err != nil
			if gotErr != tc.wantErr {
				t.Errorf("[%+v].OnlyPerson() got err=%v, want error? %t", tc.p, err, tc.wantErr)
			}
			if gotErr || tc.wantErr {
				return
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("[%+v].OnlyPerson() got diff (-want, +got):\n%s", tc.p, diff)
			}
			if diff := cmp.Diff(tc.wantID, gotID); diff != "" {
				t.Errorf("[%+v].OnlyPerson() got ID diff (-want, +got):\n%s", tc.p, diff)
			}
		})
	}
}

func TestHasPersonsDefined(t *testing.T) {
	arbitraryPerson := Person{FirstName: "first-name"}
	cases := []struct {
		name string
		p    *Persons
		want bool
	}{
		{name: "no Persons", p: nil, want: false},
		{name: "default Persons", p: &Persons{"main-patient": Person{}}, want: false},
		{name: "custom Persons", p: &Persons{"main-patient": arbitraryPerson}, want: true},
		{name: "multiple Persons", p: &Persons{"1": {}, "2": {}}, want: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			pathway := Pathway{Persons: tc.p}
			if got := pathway.HasPersonsDefined(); got != tc.want {
				t.Errorf("[%+v].HasPersonsDefined()=%t, want %t", pathway, got, tc.want)
			}
		})
	}
}

func TestMessageCount(t *testing.T) {
	delay := &Delay{
		From: time.Second,
		To:   5 * time.Second,
	}

	cases := []struct {
		name    string
		steps   []Step
		history []Step
		want    int
	}{
		{
			name: "admit and discharge",
			steps: []Step{
				{Admission: &Admission{}},
				{Discharge: &Discharge{}},
			},
			history: []Step{},
			want:    2,
		}, {
			name: "with history",
			steps: []Step{
				{Admission: &Admission{}},
				{Discharge: &Discharge{}},
			},
			history: []Step{
				{Result: &Results{}}},
			want: 3,
		}, {
			name: "order step also generates acknowledgement",
			steps: []Step{
				{Admission: &Admission{}},
				{Order: &Order{OrderProfile: "random-profile"}},
				{Discharge: &Discharge{}},
			},
			history: []Step{},
			want:    4,
		}, {
			name: "order without acknowledgement",
			steps: []Step{
				{Admission: &Admission{}},
				{Order: &Order{OrderProfile: "random-profile", NoAcknowledgementMessage: true}},
				{Discharge: &Discharge{}},
			},
			history: []Step{},
			want:    3,
		}, {
			name: "use patient step doesn't generate message",
			steps: []Step{
				{Admission: &Admission{}},
				{UsePatient: &UsePatient{Patient: "main-patient"}},
				{Discharge: &Discharge{}},
			},
			history: []Step{},
			want:    2,
		}, {
			name: "delay step doesn't genetare message",
			steps: []Step{
				{Admission: &Admission{}},
				{Delay: delay},
				{Discharge: &Discharge{}},
			},
			history: []Step{},
			want:    2,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			pathway := Pathway{
				Pathway: tc.steps,
				History: tc.history,
			}
			pathwayName := "pathway1"

			pathway.Init(pathwayName)
			got, err := pathway.MessageCount()
			if err != nil {
				t.Fatalf("[%+v].MessageCount() failed with %v", pathway, err)
			}
			if got != tc.want {
				t.Errorf("[%+v].MessageCount()=%d, want %d", pathway, got, tc.want)
			}
		})
	}
}

func TestMessageCount_MetadataNotInitialized(t *testing.T) {
	pathway := Pathway{
		Pathway: []Step{
			{Admission: &Admission{}},
			{Discharge: &Discharge{}},
		},
	}
	if _, err := pathway.MessageCount(); err == nil {
		t.Errorf("[%+v].MessageCount() got nil err, want not-nil err", pathway)
	}
}

func TestInitCalledMultipleTimes(t *testing.T) {
	pathway := Pathway{
		Pathway: []Step{
			{Admission: &Admission{}},
			{Discharge: &Discharge{}},
		},
	}
	name1 := "name-1"
	name2 := "name-2"
	name3 := "name-3"

	pathway.Init(name1)
	if got, want := pathway.Name(), name1; got != want {
		t.Errorf("[%+v].Name()=%s, want %s", pathway, got, want)
	}
	mc, err := pathway.MessageCount()
	if err != nil {
		t.Fatalf("[%+v].MessageCount() failed with %v", pathway, err)
	}
	if got, want := mc, 2; got != want {
		t.Errorf("[%+v].MessageCount()=%d, want %d", pathway, got, want)
	}

	// Init() called second time - pathway name is updated,
	// messages only count once.
	pathway.Init(name2)
	if got, want := pathway.Name(), name2; got != want {
		t.Errorf("[%+v].Name()=%s, want %s", pathway, got, want)
	}
	mc, err = pathway.MessageCount()
	if err != nil {
		t.Fatalf("[%+v].MessageCount() failed with %v", pathway, err)
	}
	if got, want := mc, 2; got != want {
		t.Errorf("[%+v].MessageCount()=%d, want %d", pathway, got, want)
	}

	// UpdateName to update the pathway name. Message count unchanged.
	pathway.UpdateName(name3)
	if got, want := pathway.Name(), name3; got != want {
		t.Errorf("[%+v].Name()=%s, want %s", pathway, got, want)
	}
	mc, err = pathway.MessageCount()
	if err != nil {
		t.Fatalf("[%+v].MessageCount() failed with %v", pathway, err)
	}
	if got, want := mc, 2; got != want {
		t.Errorf("[%+v].MessageCount()=%d, want %d", pathway, got, want)
	}
}
