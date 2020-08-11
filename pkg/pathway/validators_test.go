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
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pkg/errors"
	"github.com/google/simhospital/pkg/constants"
	"github.com/google/simhospital/pkg/doctor"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/orderprofile"
	"github.com/google/simhospital/pkg/test"
	"github.com/google/simhospital/pkg/test/testclock"
	"github.com/google/simhospital/pkg/test/testlocation"
	"github.com/google/simhospital/pkg/test/testwrite"
)

var (
	urea = &orderprofile.OrderProfile{
		UniversalService: ir.CodedElement{ID: "lpdc-3969", Text: "UREA AND ELECTROLYTES", CodingSystem: "WinPath"},
		TestTypes: map[string]*orderprofile.TestType{
			"Creatinine": {
				Name:      ir.CodedElement{ID: "lpdc-2012", Text: "Creatinine", CodingSystem: "WinPath"},
				Unit:      "UMOLL",
				ValueType: "NM",
				RefRange:  "49 - 92",
			},
		},
	}
	ureaOP = orderprofile.New(map[string]*orderprofile.OrderProfile{
		"UREA AND ELECTROLYTES": urea,
	})
	emptyOP                = orderprofile.New(map[string]*orderprofile.OrderProfile{})
	defaultClock           = testclock.New(time.Now())
	defaultValid           func(*Pathway) error
	defaultLocationManager = testlocation.ManagerWithAAndE
	emptyOrderProfiles     = orderprofile.New(make(map[string]*orderprofile.OrderProfile, 0))
	emptyDoctors           = &doctor.Doctors{}
	pathwayName            = "pathway_name"
)

func TestResultValid(t *testing.T) {
	openEnd := &orderprofile.OrderProfile{
		UniversalService: ir.CodedElement{ID: "lpdc-3969", Text: "Open End", CodingSystem: "WinPath"},
		TestTypes: map[string]*orderprofile.TestType{
			"Open End": {
				Name:      ir.CodedElement{ID: "lpdc-2012", Text: "Open End", CodingSystem: "WinPath"},
				Unit:      "UMOLL",
				ValueType: "NM",
				RefRange:  ">=5.5",
			},
		},
	}
	openEndOP := orderprofile.New(map[string]*orderprofile.OrderProfile{
		"UREA AND ELECTROLYTES": openEnd,
	})
	openStart := &orderprofile.OrderProfile{
		UniversalService: ir.CodedElement{ID: "lpdc-3969", Text: "Open Start", CodingSystem: "WinPath"},
		TestTypes: map[string]*orderprofile.TestType{
			"Open Start": {
				Name:      ir.CodedElement{ID: "lpdc-2012", Text: "Open Start", CodingSystem: "WinPath"},
				Unit:      "UMOLL",
				ValueType: "NM",
				RefRange:  "<5.5",
			},
		},
	}
	openStartOP := orderprofile.New(map[string]*orderprofile.OrderProfile{
		"UREA AND ELECTROLYTES": openStart,
	})

	cases := []struct {
		r       *Result
		op      *orderprofile.OrderProfiles
		wantErr bool
	}{
		// valid nil
		{r: nil, op: emptyOP, wantErr: false},
		{r: nil, op: ureaOP, wantErr: false},
		// valid numerical value and unit
		{r: &Result{TestName: "Creatinine", Value: "12", Unit: "UMOLL"}, op: emptyOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "<12", Unit: "UMOLL"}, op: emptyOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "0.25", Unit: "UMOLL"}, op: emptyOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "52", Unit: "UMOLL"}, op: ureaOP, wantErr: false},
		// Empty value can be set for known or unknown order profile
		{r: &Result{TestName: "Creatinine", Value: constants.EmptyString, Unit: "UMOLL"}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: constants.EmptyString, Unit: constants.EmptyString}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: constants.EmptyString, Unit: "UMOLL"}, op: emptyOP, wantErr: false},
		// Invalid: numerical value, but unit missing
		{r: &Result{TestName: "Creatinine", Value: "12"}, op: emptyOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "12"}, op: ureaOP, wantErr: true},
		// Invalid: value missing
		{r: &Result{TestName: "Creatinine", Unit: "UMOLL"}, op: emptyOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Unit: "UMOLL"}, op: ureaOP, wantErr: true},
		// valid textual value
		{r: &Result{TestName: "Creatinine", Value: "Normal result"}, op: emptyOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "Normal result below 5.5"}, op: emptyOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "Normal result below 5.5", Unit: constants.EmptyString}, op: emptyOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "Normal result"}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: constants.EmptyString}, op: ureaOP, wantErr: false},
		// invalid: unit not allowed for textual value
		{r: &Result{TestName: "Creatinine", Value: "Normal result", Unit: "UMOLL"}, op: emptyOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "Normal result", Unit: "UMOLL"}, op: ureaOP, wantErr: true},
		// invalid: random value and unit, but using different random value
		{r: &Result{TestName: "Creatinine", Value: constants.NormalValue, Unit: constants.AbnormalHigh}, op: emptyOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: constants.NormalValue, Unit: constants.AbnormalHigh}, op: ureaOP, wantErr: true},
		// invalid: cannot generate HIGH abnormal value for open-ended range
		{r: &Result{TestName: "Open End", Value: constants.AbnormalLow}, op: openEndOP, wantErr: false},
		{r: &Result{TestName: "Open End", Value: constants.NormalValue}, op: openEndOP, wantErr: false},
		{r: &Result{TestName: "Open End", Value: constants.AbnormalHigh}, op: openEndOP, wantErr: true},
		{r: &Result{TestName: "Open End", Value: constants.AbnormalLow, Unit: "UMOLL", ReferenceRange: ">=5.5"}, op: emptyOP, wantErr: false},
		{r: &Result{TestName: "Open End", Value: constants.NormalValue, Unit: "UMOLL", ReferenceRange: ">=5.5"}, op: emptyOP, wantErr: false},
		{r: &Result{TestName: "Open End", Value: constants.AbnormalHigh, Unit: "UMOLL", ReferenceRange: ">=5.5"}, op: emptyOP, wantErr: true},
		// invalid: cannot generate LOW abnormal value for open-start range
		{r: &Result{TestName: "Open Start", Value: constants.AbnormalHigh}, op: openStartOP, wantErr: false},
		{r: &Result{TestName: "Open Start", Value: constants.NormalValue}, op: openStartOP, wantErr: false},
		{r: &Result{TestName: "Open Start", Value: constants.AbnormalLow}, op: openStartOP, wantErr: true},
		{r: &Result{TestName: "Open Start", Value: constants.AbnormalHigh, Unit: "UMOLL", ReferenceRange: "<5"}, op: emptyOP, wantErr: false},
		{r: &Result{TestName: "Open Start", Value: constants.NormalValue, Unit: "UMOLL", ReferenceRange: "<5"}, op: emptyOP, wantErr: false},
		{r: &Result{TestName: "Open Start", Value: constants.AbnormalLow, Unit: "UMOLL", ReferenceRange: "<5"}, op: emptyOP, wantErr: true},
		// invalid: non-existing test type in pre-defined order profile
		{r: &Result{TestName: "Sodium", Value: "52", Unit: "UMOLL"}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Sodium", Value: "52", Unit: "UMOLL"}, op: emptyOP, wantErr: false},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("id:%d-result:%+v-op:%v-valid:%t", i, tc.r, tc.op, !tc.wantErr), func(t *testing.T) {
			p := Pathway{
				Pathway: []Step{
					{Result: &Results{OrderProfile: "UREA AND ELECTROLYTES", Results: []*Result{tc.r}}},
				},
			}
			p.Init(pathwayName)

			err := p.Valid(defaultClock, tc.op, emptyDoctors, defaultLocationManager, defaultValid)
			if gotErr := err != nil; gotErr != tc.wantErr {
				t.Errorf("[%+v].Valid(_, %+v, _, _) got err %v; want err? %t", p, tc.op, err, tc.wantErr)
			}
		})
	}
}

func TestResultValidRandomValue(t *testing.T) {
	randomValues := []string{constants.NormalValue, constants.AbnormalHigh, constants.AbnormalLow}

	for _, v := range randomValues {
		t.Run(v, func(t *testing.T) {

			cases := []struct {
				r       *Result
				op      *orderprofile.OrderProfiles
				wantErr bool
			}{
				// random value - matching order profile
				{r: &Result{TestName: "Creatinine", Value: v}, op: ureaOP, wantErr: false},
				{r: &Result{TestName: "Creatinine", Value: v, Unit: v}, op: ureaOP, wantErr: false},
				{r: &Result{TestName: "Creatinine", Value: v, Unit: "UMOLL"}, op: ureaOP, wantErr: true},
				// random value - overridden ref range: Unit must be specified to non-random
				{r: &Result{TestName: "Creatinine", Value: v, ReferenceRange: "500 - 700", Unit: "UMOLL"}, op: ureaOP, wantErr: false},
				{r: &Result{TestName: "Creatinine", Value: v, ReferenceRange: "500 - 700", Unit: "UMOLL"}, op: emptyOP, wantErr: false},
				{r: &Result{TestName: "Creatinine", Value: v, ReferenceRange: "500 - 700"}, op: ureaOP, wantErr: true},
				{r: &Result{TestName: "Creatinine", Value: v, ReferenceRange: "500 - 700"}, op: emptyOP, wantErr: true},
				{r: &Result{TestName: "Creatinine", Value: v, ReferenceRange: "500 - 700", Unit: v}, op: ureaOP, wantErr: true},
				{r: &Result{TestName: "Creatinine", Value: v, ReferenceRange: "500 - 700", Unit: v}, op: emptyOP, wantErr: true},
				// random value - no matching order profile and no ref ranges
				{r: &Result{TestName: "Creatinine", Value: v}, op: emptyOP, wantErr: true},
				{r: &Result{TestName: "Creatinine", Value: v, Unit: v}, op: emptyOP, wantErr: true},
				{r: &Result{TestName: "Creatinine", Value: v, Unit: "UMOLL"}, op: emptyOP, wantErr: true},
			}

			for i, tc := range cases {
				t.Run(fmt.Sprintf("id:%d-result:%+v-op:%v-valid:%t", i, tc.r, tc.op, !tc.wantErr), func(t *testing.T) {
					p := Pathway{
						Pathway: []Step{
							{Result: &Results{OrderProfile: "UREA AND ELECTROLYTES", Results: []*Result{tc.r}}},
						},
					}
					p.Init(pathwayName)

					err := p.Valid(defaultClock, tc.op, emptyDoctors, defaultLocationManager, defaultValid)
					if gotErr := err != nil; gotErr != tc.wantErr {
						t.Errorf("[%+v].Valid(_, %+v, _, _) got err %v; want err? %t", p, tc.op, err, tc.wantErr)
					}
				})
			}
		})
	}
}

func TestResultValidReferenceRange(t *testing.T) {
	cases := []struct {
		name    string
		r       *Result
		wantErr bool
	}{
		{
			name: "valid: reference range when numerical value specified",
			r:    &Result{TestName: "Creatinine", Value: "150", Unit: "UMOLL", ReferenceRange: "145 - 550"}, wantErr: false,
		}, {
			name: "valid: reference range when textual value specified",
			r:    &Result{TestName: "Creatinine", Value: "Normal result", ReferenceRange: "145 - 550"}, wantErr: false,
		}, {
			name: "invalid: testName is required",
			r:    &Result{Value: "150", Unit: "UMOLL", ReferenceRange: "145 - 550"}, wantErr: true,
		}, {
			name: "valid: reference range can be set for random value",
			r:    &Result{TestName: "Creatinine", Value: constants.NormalValue, Unit: "UMOLL", ReferenceRange: "145 - 550"}, wantErr: false,
		}, {
			name: "invalid: reference range cannot be set when value missing",
			r:    &Result{TestName: "Creatinine", ReferenceRange: "145 - 550"}, wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := Pathway{
				Pathway: []Step{
					{Result: &Results{OrderProfile: "UREA AND ELECTROLYTES", Results: []*Result{tc.r}}},
				},
			}
			p.Init(pathwayName)

			err := p.Valid(defaultClock, emptyOP, emptyDoctors, defaultLocationManager, defaultValid)
			if gotErr := err != nil; gotErr != tc.wantErr {
				t.Errorf("[%+v].Valid(_, %+v, _, _) got err %v; want err? %t", p, emptyOP, err, tc.wantErr)
			}
		})
	}
}

func TestResultValidAbnormalFlag(t *testing.T) {
	noRefRange := &orderprofile.OrderProfile{
		UniversalService: ir.CodedElement{ID: "lpdc-3969", Text: "UREA AND ELECTROLYTES", CodingSystem: "WinPath"},
		TestTypes: map[string]*orderprofile.TestType{
			"Creatinine": {
				Name:      ir.CodedElement{ID: "lpdc-2012", Text: "Creatinine", CodingSystem: "WinPath"},
				Unit:      "UMOLL",
				ValueType: "NM",
			},
		},
	}
	noRefRangeOP := orderprofile.New(map[string]*orderprofile.OrderProfile{
		"UREA AND ELECTROLYTES": noRefRange,
	})
	nonParsableRefRange := &orderprofile.OrderProfile{
		UniversalService: ir.CodedElement{ID: "lpdc-3969", Text: "UREA AND ELECTROLYTES", CodingSystem: "WinPath"},
		TestTypes: map[string]*orderprofile.TestType{
			"Creatinine": {
				Name:      ir.CodedElement{ID: "lpdc-2012", Text: "Creatinine", CodingSystem: "WinPath"},
				Unit:      "UMOLL",
				ValueType: "NM",
				RefRange:  "-",
			},
		},
	}
	nonParsableRefRangeOP := orderprofile.New(map[string]*orderprofile.OrderProfile{
		"UREA AND ELECTROLYTES": nonParsableRefRange,
	})

	cases := []struct {
		r       *Result
		op      *orderprofile.OrderProfiles
		wantErr bool
	}{
		// Fails if invalid AbnormalFlag value
		{r: &Result{TestName: "Creatinine", Value: "12", Unit: "UMOLL", ReferenceRange: "10 - 20", AbnormalFlag: "L"}, op: ureaOP, wantErr: true},

		// Fail if value and abnormal flag doesn't match - use ref ranges specified in the pathway
		{r: &Result{TestName: "Creatinine", Value: "12", Unit: "UMOLL", ReferenceRange: "10 - 20", AbnormalFlag: constants.AbnormalFlagEmpty}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "12", Unit: "UMOLL", ReferenceRange: "10 - 20", AbnormalFlag: constants.AbnormalFlagNormal}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "5", Unit: "UMOLL", ReferenceRange: "10 - 20", AbnormalFlag: constants.AbnormalFlagLow}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "25", Unit: "UMOLL", ReferenceRange: "10 - 20", AbnormalFlag: constants.AbnormalFlagHigh}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "25", Unit: "UMOLL", ReferenceRange: "10 - 20", AbnormalFlag: constants.AbnormalFlagEmpty}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "25", Unit: "UMOLL", ReferenceRange: "10 - 20", AbnormalFlag: constants.AbnormalFlagNormal}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "25", Unit: "UMOLL", ReferenceRange: "10 - 20", AbnormalFlag: constants.AbnormalFlagLow}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "15", Unit: "UMOLL", ReferenceRange: "10 - 20", AbnormalFlag: constants.AbnormalFlagLow}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "5", Unit: "UMOLL", ReferenceRange: "10 - 20", AbnormalFlag: constants.AbnormalFlagHigh}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "5", Unit: "UMOLL", ReferenceRange: "10 - 20", AbnormalFlag: constants.AbnormalFlagEmpty}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "5", Unit: "UMOLL", ReferenceRange: "10 - 20", AbnormalFlag: constants.AbnormalFlagNormal}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "5", Unit: "UMOLL", ReferenceRange: "10 - 20", AbnormalFlag: constants.AbnormalFlagDefault}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "5", Unit: "UMOLL", ReferenceRange: "-", AbnormalFlag: constants.AbnormalFlagEmpty}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "5", Unit: "UMOLL", ReferenceRange: "-", AbnormalFlag: constants.AbnormalFlagNormal}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "5", Unit: "UMOLL", ReferenceRange: "-", AbnormalFlag: constants.AbnormalFlagHigh}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "5", Unit: "UMOLL", ReferenceRange: "-", AbnormalFlag: constants.AbnormalFlagDefault}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "Normal Value", ReferenceRange: "-", AbnormalFlag: constants.AbnormalFlagEmpty}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "Normal Value", ReferenceRange: "-", AbnormalFlag: constants.AbnormalFlagNormal}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "High Value", ReferenceRange: "-", AbnormalFlag: constants.AbnormalFlagHigh}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "Low Value", ReferenceRange: "-", AbnormalFlag: constants.AbnormalFlagLow}, op: ureaOP, wantErr: false},

		// Fail if value and abnormal flag doesn't match - use ref ranges specified in the pathway, nil order profile
		{r: &Result{TestName: "Creatinine", Value: "12", Unit: "UMOLL", ReferenceRange: "10 - 20", AbnormalFlag: constants.AbnormalFlagEmpty}, op: emptyOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "25", Unit: "UMOLL", ReferenceRange: "10 - 20", AbnormalFlag: constants.AbnormalFlagEmpty}, op: emptyOP, wantErr: true},

		// Fail if value and abnormal flag doesn't match - use ref ranges from order profile
		{r: &Result{TestName: "Creatinine", Value: "52", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagEmpty}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "52", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagNormal}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "40", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagLow}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "95", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagHigh}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "95", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagEmpty}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "95", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagNormal}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "95", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagLow}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "52", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagLow}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "40", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagHigh}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "40", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagEmpty}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "40", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagNormal}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "40", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagDefault}, op: ureaOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "40", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagEmpty}, op: nonParsableRefRangeOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "40", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagNormal}, op: nonParsableRefRangeOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "40", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagHigh}, op: nonParsableRefRangeOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "40", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagDefault}, op: nonParsableRefRangeOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "40", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagDefault}, op: nonParsableRefRangeOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "Normal Value", AbnormalFlag: constants.AbnormalFlagEmpty}, op: nonParsableRefRangeOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "Normal Value", AbnormalFlag: constants.AbnormalFlagNormal}, op: nonParsableRefRangeOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "High Value", AbnormalFlag: constants.AbnormalFlagHigh}, op: nonParsableRefRangeOP, wantErr: false},
		{r: &Result{TestName: "Creatinine", Value: "Low Value", AbnormalFlag: constants.AbnormalFlagLow}, op: nonParsableRefRangeOP, wantErr: false},

		// Fail if AbnormalFlagDefault, but no ref range neither in the pathway nor in order profile
		{r: &Result{TestName: "Creatinine", Value: "5", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagDefault}, op: emptyOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "5", Unit: "UMOLL", AbnormalFlag: constants.AbnormalFlagDefault}, op: noRefRangeOP, wantErr: true},

		// Cannot derive abnormal flag if the value is empty or textual.
		{r: &Result{TestName: "Creatinine", Value: constants.EmptyString, AbnormalFlag: constants.AbnormalFlagDefault}, op: emptyOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: constants.EmptyString, AbnormalFlag: constants.AbnormalFlagDefault}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "Normal Value", AbnormalFlag: constants.AbnormalFlagDefault}, op: emptyOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: "Normal Value", AbnormalFlag: constants.AbnormalFlagDefault}, op: ureaOP, wantErr: true},

		// If value is set to NORMAL / ABNORMAL_HIGH / ABNORMAL_LOW, the abnormal flag will be derived automatically,
		// and so cannot be specified in the pathway
		{r: &Result{TestName: "Creatinine", Value: constants.AbnormalHigh, AbnormalFlag: constants.AbnormalFlagLow}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: constants.AbnormalLow, AbnormalFlag: constants.AbnormalFlagLow}, op: ureaOP, wantErr: true},
		{r: &Result{TestName: "Creatinine", Value: constants.NormalValue, AbnormalFlag: constants.AbnormalFlagHigh}, op: ureaOP, wantErr: true},

		// Fail if random value on absent (nil) order profile and invalid ReferenceRange on the pathway.
		{r: &Result{TestName: "Test1", Value: constants.NormalValue, Unit: "UMOLL", ReferenceRange: "10 - 20", AbnormalFlag: constants.AbnormalFlagEmpty}, op: emptyOP, wantErr: false},
		{r: &Result{TestName: "Test1", Value: constants.AbnormalHigh, Unit: "UMOLL", ReferenceRange: "10 - 20", AbnormalFlag: constants.AbnormalFlagEmpty}, op: emptyOP, wantErr: false},
		{r: &Result{TestName: "Test1", Value: constants.NormalValue, Unit: "UMOLL", ReferenceRange: "", AbnormalFlag: constants.AbnormalFlagEmpty}, op: emptyOP, wantErr: true},
		{r: &Result{TestName: "Test1", Value: constants.NormalValue, Unit: "UMOLL", ReferenceRange: "-", AbnormalFlag: constants.AbnormalFlagEmpty}, op: emptyOP, wantErr: true},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("id:%d-result:%+v-op:%v-valid:%t", i, tc.r, tc.op, !tc.wantErr), func(t *testing.T) {
			p := Pathway{
				Pathway: []Step{
					{Result: &Results{OrderProfile: "UREA AND ELECTROLYTES", Results: []*Result{tc.r}}},
				},
			}
			p.Init(pathwayName)

			err := p.Valid(defaultClock, tc.op, emptyDoctors, defaultLocationManager, defaultValid)
			if gotErr := err != nil; gotErr != tc.wantErr {
				t.Errorf("[%+v].Valid(_, %+v, _, _) got err %v; want err? %t", p, tc.op, err, tc.wantErr)
			}
		})
	}
}

func TestResultsValid(t *testing.T) {
	cases := []struct {
		name    string
		r       *Results
		wantErr bool
	}{
		{
			name: "valid: specified order profile",
			r: &Results{
				OrderProfile: "UREA AND ELECTROLYTES",
				Results: []*Result{
					{
						TestName: "Creatinine",
						Value:    "200",
						Unit:     "UMOLL",
					},
				},
			},
			wantErr: false,
		}, {
			name:    "valid: random order profile",
			r:       &Results{OrderProfile: constants.RandomString},
			wantErr: false,
		}, {
			name: "invalid: random order profile but results specified",
			r: &Results{
				OrderProfile: constants.RandomString,
				Results: []*Result{
					{
						TestName: "Creatinine",
						Value:    "200",
						Unit:     "UMOLL",
					},
				},
			},
			wantErr: true,
		}, {
			name: "valid: override OrderStatus, ResultStatus and ResultStatus",
			r: &Results{
				OrderProfile: "UREA AND ELECTROLYTES",
				OrderStatus:  "A",
				ResultStatus: "P",
				Results: []*Result{
					{
						TestName:     "Creatinine",
						ResultStatus: "R",
						Value:        "200",
						Unit:         "UMOLL",
					},
				},
			},
			wantErr: false,
		}, {
			name: "valid: override OrderStatus and ResultStatus",
			r: &Results{
				OrderProfile: "UREA AND ELECTROLYTES",
				OrderStatus:  "A",
				ResultStatus: "P",
				Results: []*Result{
					{
						TestName: "Creatinine",
						Value:    "200",
						Unit:     "UMOLL",
					},
				},
			},
			wantErr: false,
		}, {
			name: "invalid: only OrderStatus overridden",
			r: &Results{
				OrderProfile: "UREA AND ELECTROLYTES",
				OrderStatus:  "A",
				Results: []*Result{
					{
						TestName: "Creatinine",
						Value:    "200",
						Unit:     "UMOLL",
					},
				},
			},
			wantErr: true,
		}, {
			name: "invalid: only ResultStatus overridden",
			r: &Results{
				OrderProfile: "UREA AND ELECTROLYTES",
				ResultStatus: "P",
				Results: []*Result{
					{
						TestName: "Creatinine",
						Value:    "200",
						Unit:     "UMOLL",
					},
				},
			},
			wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := Pathway{
				Pathway: []Step{
					{Result: tc.r},
				},
			}
			p.Init(pathwayName)

			err := p.Valid(defaultClock, emptyOP, emptyDoctors, defaultLocationManager, defaultValid)
			if gotErr := err != nil; gotErr != tc.wantErr {
				t.Errorf("[%+v].Valid(_, _, _, _) got err %v; want err? %t", p, err, tc.wantErr)
			}
		})
	}
}

func TestResultsValidOverrideDates(t *testing.T) {
	dates := []struct {
		value   string
		wantErr bool
	}{
		{value: constants.MidnightDate, wantErr: false},
		{value: constants.EmptyString, wantErr: false},
		{value: "20160321165530", wantErr: true},
		{value: "anything", wantErr: true},
	}

	for _, d := range dates {
		t.Run(fmt.Sprintf("collected-%s", d.value), func(t *testing.T) {
			r := &Results{
				OrderProfile:      "UREA AND ELECTROLYTES",
				CollectedDateTime: d.value,
				Results: []*Result{
					{
						TestName:     "Creatinine",
						ResultStatus: "R",
						Value:        "200",
						Unit:         "UMOLL",
					},
				},
			}
			p := Pathway{
				Pathway: []Step{
					{Result: r},
				},
			}
			p.Init(pathwayName)

			err := p.Valid(defaultClock, emptyOP, emptyDoctors, defaultLocationManager, defaultValid)
			if gotErr := err != nil; gotErr != d.wantErr {
				t.Errorf("[%+v].Valid(_, _, _, _) got err %v; want err? %t", p, err, d.wantErr)
			}
		})
	}

	for _, d := range dates {
		t.Run(fmt.Sprintf("received-in-lab-%s", d.value), func(t *testing.T) {
			r := &Results{
				OrderProfile:          "UREA AND ELECTROLYTES",
				ReceivedInLabDateTime: d.value,
				Results: []*Result{
					{
						TestName:     "Creatinine",
						ResultStatus: "R",
						Value:        "200",
						Unit:         "UMOLL",
					},
				},
			}
			p := Pathway{
				Pathway: []Step{
					{Result: r},
				},
			}
			p.Init(pathwayName)

			err := p.Valid(defaultClock, emptyOP, emptyDoctors, defaultLocationManager, defaultValid)
			if gotErr := err != nil; gotErr != d.wantErr {
				t.Errorf("[%+v].Valid(_, _, _, _) got err %v; want err? %t", p, err, d.wantErr)
			}
		})
	}
}

func TestPathwayValidStep(t *testing.T) {
	negative := -15 * time.Hour
	negativeOneHour := -1 * time.Hour
	negativeTwoHours := -2 * time.Hour
	oneHour := time.Hour
	twoHours := 2 * time.Hour
	fifteenHoursAgo := time.Now().UTC().Add(negative)
	fiveMinutes := 5 * time.Minute

	cases := []struct {
		step    Step
		wantErr bool
	}{
		{step: Step{AddPerson: &AddPerson{}}},
		{step: Step{UpdatePerson: &UpdatePerson{Person: &Person{}}}},
		{step: Step{PendingAdmission: &PendingAdmission{Loc: "ED", ExpectedAdmissionTimeFromNow: &fiveMinutes}}},
		{step: Step{PendingDischarge: &PendingDischarge{ExpectedDischargeTimeFromNow: &fiveMinutes}}},
		{step: Step{PendingTransfer: &PendingTransfer{Loc: "ED", ExpectedTransferTimeFromNow: &fiveMinutes}}},
		{step: Step{TransferInError: &TransferInError{Loc: "ED"}}},
		{step: Step{DischargeInError: &DischargeInError{}}},
		{step: Step{Discharge: &Discharge{}}},
		{step: Step{Transfer: &Transfer{Loc: "ED"}}},
		{step: Step{Admission: &Admission{Loc: "ED"}}},
		{step: Step{CancelTransfer: &CancelTransfer{}}},
		{step: Step{CancelDischarge: &CancelDischarge{}}},
		{step: Step{CancelVisit: &CancelVisit{}}},
		{step: Step{Order: &Order{OrderProfile: "profile"}}},
		{step: Step{Registration: &Registration{}}},
		{step: Step{CancelPendingAdmission: &CancelPendingAdmission{}}},
		{step: Step{CancelPendingDischarge: &CancelPendingDischarge{}}},
		{step: Step{CancelPendingTransfer: &CancelPendingTransfer{}}},
		{step: Step{DeleteVisit: &DeleteVisit{}}},
		{step: Step{TrackDeparture: &TrackDeparture{DestinationLoc: "ED", Mode: "track"}}},
		{step: Step{TrackDeparture: &TrackDeparture{DestinationLoc: "ED", Mode: "transit"}}},
		{step: Step{TrackDeparture: &TrackDeparture{DestinationLoc: "ED", Mode: "temporary"}}},
		{step: Step{TrackArrival: &TrackArrival{Loc: "ED", Mode: "track"}}},
		{step: Step{TrackArrival: &TrackArrival{Loc: "ED", Mode: "transit"}}},
		{step: Step{TrackArrival: &TrackArrival{Loc: "ED", Mode: "temporary"}}},
		{step: Step{Result: &Results{OrderProfile: "profile"}}},
		{step: Step{Result: &Results{OrderProfile: "profile", TriggerEvent: ""}}},
		{step: Step{Result: &Results{OrderProfile: "profile", TriggerEvent: "R01"}}},
		{step: Step{Result: &Results{OrderProfile: "profile", TriggerEvent: "r01"}}},
		{step: Step{Result: &Results{OrderProfile: "profile", TriggerEvent: "R03"}}},
		{step: Step{Result: &Results{OrderProfile: "profile", TriggerEvent: "R32"}}},
		{step: Step{Result: &Results{OrderProfile: "profile", TriggerEvent: "something-else"}}, wantErr: true},
		{step: Step{Delay: &Delay{}}},
		{step: Step{Delay: &Delay{}, Parameters: &Parameters{}}},
		// Steps that require a location must have a valid one.
		// Steps that require other fields in addition to the location are tested later.
		{step: Step{Admission: &Admission{Loc: "ED"}}},
		{step: Step{Admission: &Admission{Loc: "nonexistent-location"}}, wantErr: true},
		{step: Step{Admission: &Admission{}}, wantErr: true},
		{step: Step{Transfer: &Transfer{Loc: "ED"}}},
		{step: Step{Transfer: &Transfer{Loc: "nonexistent-location"}}, wantErr: true},
		{step: Step{Transfer: &Transfer{}}, wantErr: true},
		{step: Step{TransferInError: &TransferInError{Loc: "ED"}}},
		{step: Step{TransferInError: &TransferInError{Loc: "nonexistent-location"}}, wantErr: true},
		{step: Step{TransferInError: &TransferInError{}}, wantErr: true},
		// Allergies have to have either Code or Description (or both).
		{step: Step{Admission: &Admission{Loc: "ED", Allergies: []Allergy{{Code: "A11.0"}}}}},
		{step: Step{Admission: &Admission{Loc: "ED", Allergies: []Allergy{{Description: "Allergy"}}}}},
		{step: Step{Admission: &Admission{Loc: "ED", Allergies: []Allergy{{Code: "A11.0", Description: "Allergy"}}}}},
		{step: Step{Admission: &Admission{Loc: "ED", Allergies: []Allergy{{}}}}, wantErr: true},
		{step: Step{Discharge: &Discharge{Allergies: []Allergy{{}}}}, wantErr: true},
		{step: Step{DischargeInError: &DischargeInError{Allergies: []Allergy{{}}}}, wantErr: true},
		{step: Step{UpdatePerson: &UpdatePerson{Allergies: []Allergy{{}}}}, wantErr: true},
		{step: Step{Registration: &Registration{Allergies: []Allergy{{}}}}, wantErr: true},
		{step: Step{AddPerson: &AddPerson{Allergies: []Allergy{{}}}}, wantErr: true},
		// Discharges don't have any required fields.
		{step: Step{Discharge: &Discharge{}}},
		// Merges must have Parent and at least one child.
		{step: Step{Merge: &Merge{}}, wantErr: true},
		{step: Step{Merge: &Merge{Parent: "1", Children: []PatientID{PatientID("9999")}}}},
		{step: Step{Merge: &Merge{Parent: "1"}}, wantErr: true},
		{step: Step{Merge: &Merge{Children: []PatientID{PatientID("9999")}}}, wantErr: true},
		// An A40 message can only be forced if Children has one single item.
		{step: Step{Merge: &Merge{Parent: "1", Children: []PatientID{PatientID("9999")}, ForceA40: true}}},
		{step: Step{Merge: &Merge{Parent: "1", Children: []PatientID{PatientID("9999"), PatientID("8888")}, ForceA40: true}}, wantErr: true},
		// Bed Swaps must have Patient1 and Patient2.
		{step: Step{BedSwap: &BedSwap{}}, wantErr: true},
		{step: Step{BedSwap: &BedSwap{Patient1: "9999", Patient2: "8888"}}},
		{step: Step{BedSwap: &BedSwap{Patient2: "8888"}}, wantErr: true},
		{step: Step{BedSwap: &BedSwap{Patient1: "9999"}}, wantErr: true},
		// Only one field must be set in a step.
		{step: Step{Delay: &Delay{}, Admission: &Admission{Loc: "ED"}}, wantErr: true},
		{step: Step{Delay: &Delay{}, Order: &Order{}}, wantErr: true},
		{step: Step{Order: &Order{OrderProfile: "profile"}, Discharge: &Discharge{}}, wantErr: true},
		// A Step must have one field different from Parameters or Metadata.
		{step: Step{Parameters: &Parameters{}}, wantErr: true},
		{step: Step{}, wantErr: true},
		// Delays must have meaningful intervals, and can't have DelayMessage parameters.
		{step: Step{Delay: &Delay{From: oneHour, To: twoHours}}},
		{step: Step{Delay: &Delay{From: twoHours, To: oneHour}}, wantErr: true},
		{step: Step{Delay: &Delay{From: negative, To: oneHour}}, wantErr: true},
		{step: Step{Delay: &Delay{}, Parameters: &Parameters{DelayMessage: &Delay{From: oneHour, To: twoHours}}}, wantErr: true},
		// Orders and Results require an OrderProfile and/or OrderID.
		{step: Step{Order: &Order{OrderProfile: "profile"}}},
		{step: Step{Result: &Results{OrderProfile: "profile"}}},
		{step: Step{Result: &Results{OrderProfile: "profile"}, Parameters: &Parameters{DelayMessage: &Delay{From: oneHour, To: twoHours}}}},
		{step: Step{Result: &Results{OrderProfile: "profile"}, Parameters: &Parameters{DelayMessage: &Delay{From: twoHours, To: oneHour}}}, wantErr: true},
		{step: Step{Result: &Results{OrderProfile: "profile"}, Parameters: &Parameters{DelayMessage: &Delay{From: negative, To: oneHour}}}, wantErr: true},
		{step: Step{Order: &Order{OrderProfile: "profile"}}},
		{step: Step{Order: &Order{OrderID: "order1", OrderProfile: "profile"}}},
		{step: Step{Order: &Order{}}, wantErr: true},
		{step: Step{Result: &Results{OrderProfile: "profile"}}},
		{step: Step{Result: &Results{OrderID: "order1", OrderProfile: "profile"}}},
		{step: Step{Result: &Results{}}, wantErr: true},
		// A PreAdmission must have both a valid Loc and a positive ExpectedAdmissionTimeFromNow. Optionally it can have allergies.
		{step: Step{PreAdmission: &PreAdmission{Loc: "ED", ExpectedAdmissionTimeFromNow: &oneHour}}},
		{step: Step{PreAdmission: &PreAdmission{Loc: "dummy", ExpectedAdmissionTimeFromNow: &oneHour}}, wantErr: true},
		{step: Step{PreAdmission: &PreAdmission{Loc: "", ExpectedAdmissionTimeFromNow: &oneHour}}, wantErr: true},
		{step: Step{PreAdmission: &PreAdmission{Loc: "ED"}}, wantErr: true},
		{step: Step{PreAdmission: &PreAdmission{ExpectedAdmissionTimeFromNow: &oneHour}}, wantErr: true},
		{step: Step{PreAdmission: &PreAdmission{Loc: "ED", ExpectedAdmissionTimeFromNow: &negative}}, wantErr: true},
		{step: Step{PreAdmission: &PreAdmission{Loc: "ED", ExpectedAdmissionTimeFromNow: &oneHour, Allergies: []Allergy{{Code: "A11.0"}}}}},
		// A PendingAdmission must have both Loc and a positive ExpectedAdmissionTimeFromNow.
		{step: Step{PendingAdmission: &PendingAdmission{Loc: "ED", ExpectedAdmissionTimeFromNow: &oneHour}}},
		{step: Step{PendingAdmission: &PendingAdmission{Loc: "nonexistent-location", ExpectedAdmissionTimeFromNow: &oneHour}}, wantErr: true},
		{step: Step{PendingAdmission: &PendingAdmission{Loc: "", ExpectedAdmissionTimeFromNow: &oneHour}}, wantErr: true},
		{step: Step{PendingAdmission: &PendingAdmission{Loc: "ED"}}, wantErr: true},
		{step: Step{PendingAdmission: &PendingAdmission{ExpectedAdmissionTimeFromNow: &oneHour}}, wantErr: true},
		{step: Step{PendingAdmission: &PendingAdmission{Loc: "ED", ExpectedAdmissionTimeFromNow: &negative}}, wantErr: true},
		// A PendingDischarge must have a positive ExpectedDischargeTimeFromNow.
		{step: Step{PendingDischarge: &PendingDischarge{ExpectedDischargeTimeFromNow: &oneHour}}},
		{step: Step{PendingDischarge: &PendingDischarge{}}, wantErr: true},
		{step: Step{PendingDischarge: &PendingDischarge{ExpectedDischargeTimeFromNow: &negative}}, wantErr: true},
		// A PendingTransfer must have both Loc and a positive ExpectedAdmissionTimeFromNow.
		{step: Step{PendingTransfer: &PendingTransfer{Loc: "ED", ExpectedTransferTimeFromNow: &oneHour}}},
		{step: Step{PendingTransfer: &PendingTransfer{Loc: "nonexistent-location", ExpectedTransferTimeFromNow: &oneHour}}, wantErr: true},
		{step: Step{PendingTransfer: &PendingTransfer{Loc: "", ExpectedTransferTimeFromNow: &oneHour}}, wantErr: true},
		{step: Step{PendingTransfer: &PendingTransfer{Loc: "ED"}}, wantErr: true},
		{step: Step{PendingTransfer: &PendingTransfer{ExpectedTransferTimeFromNow: &oneHour}}, wantErr: true},
		{step: Step{PendingTransfer: &PendingTransfer{Loc: "ED", ExpectedTransferTimeFromNow: &negative}}, wantErr: true},
		// TrackDeparture and TrackArrival must have a DestinationLoc or Loc, respectively, and a supported Mode.
		// The location must be a known location unless the mode is Temporary.
		{step: Step{TrackDeparture: &TrackDeparture{DestinationLoc: "ED"}}, wantErr: true},
		{step: Step{TrackDeparture: &TrackDeparture{DestinationLoc: "ED", Mode: "unsupported"}}, wantErr: true},
		{step: Step{TrackDeparture: &TrackDeparture{DestinationLoc: "ED", Mode: "track"}}},
		{step: Step{TrackDeparture: &TrackDeparture{DestinationLoc: "nonexistent-location", Mode: "track"}}, wantErr: true},
		{step: Step{TrackDeparture: &TrackDeparture{DestinationLoc: "", Mode: "track"}}, wantErr: true},
		{step: Step{TrackDeparture: &TrackDeparture{DestinationLoc: "ED", Mode: "transit"}}},
		{step: Step{TrackDeparture: &TrackDeparture{DestinationLoc: "nonexistent-location", Mode: "transit"}}, wantErr: true},
		{step: Step{TrackDeparture: &TrackDeparture{DestinationLoc: "", Mode: "transit"}}, wantErr: true},
		{step: Step{TrackDeparture: &TrackDeparture{DestinationLoc: "ED", Mode: "temporary"}}},
		{step: Step{TrackDeparture: &TrackDeparture{DestinationLoc: "nonexistent-location", Mode: "temporary"}}},
		{step: Step{TrackDeparture: &TrackDeparture{DestinationLoc: "", Mode: "temporary"}}, wantErr: true},
		{step: Step{TrackArrival: &TrackArrival{Loc: "ED"}}, wantErr: true},
		{step: Step{TrackArrival: &TrackArrival{Loc: "ED", Mode: "unsupported"}}, wantErr: true},
		{step: Step{TrackArrival: &TrackArrival{Loc: "ED", Mode: "track"}}},
		{step: Step{TrackArrival: &TrackArrival{Loc: "nonexistent-location", Mode: "track"}}, wantErr: true},
		{step: Step{TrackArrival: &TrackArrival{Loc: "", Mode: "track"}}, wantErr: true},
		{step: Step{TrackArrival: &TrackArrival{Loc: "ED", Mode: "transit"}}},
		{step: Step{TrackArrival: &TrackArrival{Loc: "nonexistent-location", Mode: "transit"}}, wantErr: true},
		{step: Step{TrackArrival: &TrackArrival{Loc: "", Mode: "transit"}}, wantErr: true},
		{step: Step{TrackArrival: &TrackArrival{Loc: "ED", Mode: "temporary"}}},
		{step: Step{TrackArrival: &TrackArrival{Loc: "nonexistent-location", Mode: "temporary", IsTemporary: true}}},
		{step: Step{TrackArrival: &TrackArrival{Loc: "nonexistent-location", Mode: "temporary", IsTemporary: false}}, wantErr: true},
		{step: Step{TrackArrival: &TrackArrival{Loc: "", Mode: "temporary"}}, wantErr: true},
		// TrackDeparture can have a DestinationBed if the mode is not 'temporary'.
		{step: Step{TrackDeparture: &TrackDeparture{DestinationLoc: "ED", DestinationBed: "dummy", Mode: "temporary"}}, wantErr: true},
		{step: Step{TrackDeparture: &TrackDeparture{DestinationLoc: "ED", DestinationBed: "dummy", Mode: "track"}}},
		{step: Step{TrackDeparture: &TrackDeparture{DestinationLoc: "ED", DestinationBed: "dummy", Mode: "transit"}}},
		// TrackArrival can have a Bed only if the mode is not 'transit'.
		{step: Step{TrackArrival: &TrackArrival{Loc: "ED", Bed: "dummy", Mode: "temporary"}}},
		{step: Step{TrackArrival: &TrackArrival{Loc: "ED", Bed: "dummy", Mode: "track"}}},
		{step: Step{TrackArrival: &TrackArrival{Loc: "ED", Bed: "dummy", Mode: "transit"}}, wantErr: true},
		// IsTemporary can only be set in 'temporary' mode.
		{step: Step{TrackArrival: &TrackArrival{Loc: "ED", IsTemporary: true, Mode: "temporary"}}},
		{step: Step{TrackArrival: &TrackArrival{Loc: "ED", IsTemporary: true, Mode: "track"}}, wantErr: true},
		{step: Step{TrackArrival: &TrackArrival{Loc: "ED", IsTemporary: true, Mode: "transit"}}, wantErr: true},
		// An UpdatePerson must have a valid Diagnoses: optional DateTime in the past; at least one from: Description or Code.
		{step: Step{UpdatePerson: &UpdatePerson{Diagnoses: []*DiagnosisOrProcedure{{Description: "Diagnosis", DateTime: &DateTime{TimeFromNow: &negative}}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Diagnoses: []*DiagnosisOrProcedure{{Description: "Diagnosis", DateTime: &DateTime{Time: &fifteenHoursAgo}}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Diagnoses: []*DiagnosisOrProcedure{{Description: "Diagnosis", DateTime: &DateTime{NoDateTimeRecorded: true}}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Diagnoses: []*DiagnosisOrProcedure{{Description: "Diagnosis", DateTime: &DateTime{TimeFromNow: &twoHours}}}}}, wantErr: true},
		{step: Step{UpdatePerson: &UpdatePerson{Diagnoses: []*DiagnosisOrProcedure{{Description: "Diagnosis"}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Diagnoses: []*DiagnosisOrProcedure{{Code: "A10.1", DateTime: &DateTime{TimeFromNow: &negative}}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Diagnoses: []*DiagnosisOrProcedure{{DateTime: &DateTime{TimeFromNow: &negative}}}}}, wantErr: true},
		// RANDOM diagnosis: either Code or Description (or both) set to random; TimeFromNow not required; Type cannot be set.
		{step: Step{UpdatePerson: &UpdatePerson{Diagnoses: []*DiagnosisOrProcedure{{Code: "RANDOM"}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Diagnoses: []*DiagnosisOrProcedure{{Description: "RANDOM"}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Diagnoses: []*DiagnosisOrProcedure{{Code: "RANDOM", Description: "RANDOM"}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Diagnoses: []*DiagnosisOrProcedure{{Code: "RANDOM", Description: "non random"}}}}, wantErr: true},
		{step: Step{UpdatePerson: &UpdatePerson{Diagnoses: []*DiagnosisOrProcedure{{Code: "RANDOM", DateTime: &DateTime{TimeFromNow: &negative}}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Diagnoses: []*DiagnosisOrProcedure{{Code: "RANDOM", DateTime: &DateTime{TimeFromNow: &twoHours}}}}}, wantErr: true},
		{step: Step{UpdatePerson: &UpdatePerson{Diagnoses: []*DiagnosisOrProcedure{{Description: "RANDOM", Type: "some type"}}}}, wantErr: true},
		// An UpdatePerson must have a valid Procedures: optional DateTime in the past; at least one from: Description or Code.
		{step: Step{UpdatePerson: &UpdatePerson{Procedures: []*DiagnosisOrProcedure{{Description: "Procedure", DateTime: &DateTime{TimeFromNow: &negative}}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Procedures: []*DiagnosisOrProcedure{{Description: "Procedure", DateTime: &DateTime{Time: &fifteenHoursAgo}}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Procedures: []*DiagnosisOrProcedure{{Description: "Procedure", DateTime: &DateTime{NoDateTimeRecorded: true}}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Procedures: []*DiagnosisOrProcedure{{Description: "Procedure", DateTime: &DateTime{TimeFromNow: &twoHours}}}}}, wantErr: true},
		{step: Step{UpdatePerson: &UpdatePerson{Procedures: []*DiagnosisOrProcedure{{Description: "Procedure"}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Procedures: []*DiagnosisOrProcedure{{Code: "A10.1", DateTime: &DateTime{TimeFromNow: &negative}}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Procedures: []*DiagnosisOrProcedure{{DateTime: &DateTime{TimeFromNow: &negative}}}}}, wantErr: true},
		// RANDOM procedure: either Code or Description (or both) set to random; TimeFromNow not required; Type cannot be set;
		{step: Step{UpdatePerson: &UpdatePerson{Procedures: []*DiagnosisOrProcedure{{Code: "RANDOM"}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Procedures: []*DiagnosisOrProcedure{{Description: "RANDOM"}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Procedures: []*DiagnosisOrProcedure{{Code: "RANDOM", Description: "RANDOM"}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Procedures: []*DiagnosisOrProcedure{{Code: "RANDOM", Description: "non random"}}}}, wantErr: true},
		{step: Step{UpdatePerson: &UpdatePerson{Procedures: []*DiagnosisOrProcedure{{Code: "RANDOM", DateTime: &DateTime{TimeFromNow: &negative}}}}}},
		{step: Step{UpdatePerson: &UpdatePerson{Procedures: []*DiagnosisOrProcedure{{Code: "RANDOM", DateTime: &DateTime{TimeFromNow: &twoHours}}}}}, wantErr: true},
		{step: Step{UpdatePerson: &UpdatePerson{Procedures: []*DiagnosisOrProcedure{{Description: "RANDOM", Type: "some type"}}}}, wantErr: true},
		// AutoGenerate requires a positive time interval From -> To and a positive Every if To != From;
		// Results Step is optional but if set must be valid.
		{step: Step{AutoGenerate: &AutoGenerate{From: &oneHour, To: &twoHours, Every: &oneHour}}},
		{step: Step{AutoGenerate: &AutoGenerate{From: &oneHour, To: &oneHour}}},
		{step: Step{AutoGenerate: &AutoGenerate{From: &oneHour, To: &oneHour, Every: &oneHour}}, wantErr: true},
		{step: Step{AutoGenerate: &AutoGenerate{To: &twoHours, Every: &oneHour}}, wantErr: true},
		{step: Step{AutoGenerate: &AutoGenerate{From: &oneHour, Every: &oneHour}}, wantErr: true},
		{step: Step{AutoGenerate: &AutoGenerate{From: &oneHour, To: &twoHours}}, wantErr: true},
		{step: Step{AutoGenerate: &AutoGenerate{From: &negativeTwoHours, To: &negativeOneHour, Every: &oneHour}}},
		{step: Step{AutoGenerate: &AutoGenerate{From: &negativeTwoHours, To: &oneHour, Every: &oneHour}}},
		{step: Step{AutoGenerate: &AutoGenerate{From: &negativeOneHour, To: &negativeTwoHours, Every: &oneHour}}, wantErr: true},
		{step: Step{AutoGenerate: &AutoGenerate{From: &twoHours, To: &oneHour, Every: &oneHour}}, wantErr: true},
		{step: Step{AutoGenerate: &AutoGenerate{From: &oneHour, To: &twoHours}}, wantErr: true},
		{step: Step{AutoGenerate: &AutoGenerate{From: &oneHour, To: &twoHours, Every: &negativeOneHour}}, wantErr: true},
		{step: Step{AutoGenerate: &AutoGenerate{Result: &Results{OrderProfile: "profile"}, From: &oneHour, To: &twoHours, Every: &oneHour}}},
		{step: Step{AutoGenerate: &AutoGenerate{Result: &Results{}, From: &oneHour, To: &twoHours, Every: &oneHour}}, wantErr: true},
		{step: Step{ClinicalNote: &ClinicalNote{}}, wantErr: true},
		{step: Step{ClinicalNote: &ClinicalNote{ContentType: "some_content_type"}}},
		{step: Step{ClinicalNote: &ClinicalNote{DocumentType: "Some type", ContentType: "some_content_type", DocumentID: "document_id"}}},
		{step: Step{ClinicalNote: &ClinicalNote{ContentType: "txt", DocumentContent: "some-content"}}},
		{step: Step{ClinicalNote: &ClinicalNote{ContentType: "pdf", DocumentContent: "some-content"}}, wantErr: true},
		{step: Step{HardcodedMessage: &HardcodedMessage{Regex: ".*"}}},
		{step: Step{HardcodedMessage: &HardcodedMessage{Regex: ""}}, wantErr: true},
		{step: Step{HardcodedMessage: &HardcodedMessage{}}, wantErr: true},
		{step: Step{HardcodedMessage: &HardcodedMessage{}}, wantErr: true},
		// DeathStatus cannot have both TimeSinceDeath and TimeOfDeath set at the same time.
		{step: Step{Admission: &Admission{Loc: "ED"}, Parameters: &Parameters{Status: &DeathStatus{TimeSinceDeath: &oneHour}}}},
		{step: Step{Admission: &Admission{Loc: "ED"}, Parameters: &Parameters{Status: &DeathStatus{TimeSinceDeath: &oneHour, TimeOfDeath: &fifteenHoursAgo}}}, wantErr: true},
		{step: Step{Admission: &Admission{Loc: "ED"}, Parameters: &Parameters{Status: &DeathStatus{TimeOfDeath: &fifteenHoursAgo}}}},
		{step: Step{Document: &Document{ID: "docid1"}}},
		{step: Step{Document: &Document{}}},
		// All NumRandomContentLines in Documents must be valid Intervals.
		{step: Step{Document: &Document{NumRandomContentLines: &Interval{To: -1}}}, wantErr: true},
		{step: Step{Document: &Document{NumRandomContentLines: &Interval{From: -1}}}, wantErr: true},
		{step: Step{Document: &Document{NumRandomContentLines: &Interval{To: 20}}}},
		{step: Step{Document: &Document{NumRandomContentLines: &Interval{From: 10}}}, wantErr: true},
		{step: Step{Document: &Document{NumRandomContentLines: &Interval{To: 0}}}},
		{step: Step{Document: &Document{NumRandomContentLines: &Interval{From: 30, To: 10}}}, wantErr: true},
		{step: Step{Document: &Document{NumRandomContentLines: &Interval{From: 10, To: 30}}}},
		{step: Step{Document: &Document{NumRandomContentLines: &Interval{From: 10, To: 10}}}},
		{step: Step{Document: &Document{NumRandomContentLines: &Interval{From: 0, To: 0}}}},
		{step: Step{Document: &Document{NumRandomContentLines: &Interval{}}}},
		// UpdateType must be "append" or "overwrite".
		{step: Step{Document: &Document{ID: "id1", UpdateType: "append"}}},
		{step: Step{Document: &Document{ID: "id1", UpdateType: "overwrite"}}},
		{step: Step{Document: &Document{ID: "id1", UpdateType: "delete"}}, wantErr: true},
		// Append type document updates must add at least 1 line.
		{step: Step{Document: &Document{ID: "docid1", UpdateType: "append", NumRandomContentLines: &Interval{From: 0, To: 0}}}, wantErr: true},
		{step: Step{Document: &Document{ID: "docid1", UpdateType: "append", NumRandomContentLines: &Interval{}}}, wantErr: true},
		{step: Step{Document: &Document{ID: "docid1", UpdateType: "append", HeaderContentLines: []string{"header"}, NumRandomContentLines: &Interval{}}}, wantErr: false},
		{step: Step{Document: &Document{ID: "docid1", UpdateType: "append", EndingContentLines: []string{"ending"}, NumRandomContentLines: &Interval{}}}, wantErr: false},
		{step: Step{Document: &Document{ID: "docid1", UpdateType: "append", EndingContentLines: []string{"ending"}}}, wantErr: false},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("id:%d-step:%+v-valid:%t", i, tc.step, !tc.wantErr), func(t *testing.T) {
			p := Pathway{Pathway: []Step{tc.step}}
			p.Init(pathwayName)

			err := p.Valid(defaultClock, emptyOP, emptyDoctors, defaultLocationManager, defaultValid)
			if gotErr := err != nil; gotErr != tc.wantErr {
				t.Errorf("[%+v].Valid(%v, %v, %v, _) got err %v; want err? %t", p, defaultClock, emptyOP, emptyDoctors, err, tc.wantErr)
			}
		})
	}
}

func TestPathwayValidPathway(t *testing.T) {
	twoHoursAgo := -2 * time.Hour
	oneHourAgo := -time.Hour
	oneHour := time.Hour
	twoHours := 2 * time.Hour
	second := time.Second

	stepOneHourAgo := Step{Order: &Order{OrderID: "order1", OrderProfile: "profile"}, Parameters: &Parameters{TimeFromNow: &oneHourAgo}}
	stepTwoHoursAgo := Step{Order: &Order{OrderID: "order1", OrderProfile: "profile"}, Parameters: &Parameters{TimeFromNow: &twoHoursAgo}}

	addPerson := Step{AddPerson: &AddPerson{}}
	if err := addPerson.valid(defaultClock.Now(), defaultLocationManager); err != nil {
		t.Fatalf("addPerson.valid(%v) got error %v, want nil", defaultClock.Now(), err)
	}
	discharge := Step{Discharge: &Discharge{}}
	if err := discharge.valid(defaultClock.Now(), defaultLocationManager); err != nil {
		t.Fatalf("discharge.valid(%v) got error %v, want nil", defaultClock.Now(), err)
	}
	admit := Step{Admission: &Admission{Loc: "ED"}}
	if err := admit.valid(defaultClock.Now(), defaultLocationManager); err != nil {
		t.Fatalf("admit.valid(%v) got error %v, want nil", defaultClock.Now(), err)
	}
	result := Step{Result: &Results{OrderProfile: "profile"}}
	if err := result.valid(defaultClock.Now(), defaultLocationManager); err != nil {
		t.Fatalf("result.valid() got error %v, want nil", err)
	}
	resultNegativeTimeFromNow := Step{Result: &Results{OrderProfile: "profile"}, Parameters: &Parameters{TimeFromNow: &twoHoursAgo}}
	if err := resultNegativeTimeFromNow.valid(defaultClock.Now(), defaultLocationManager); err != nil {
		t.Fatalf("resultNegativeTimeFromNow.valid(%v) got error %v, want nil", defaultClock.Now(), err)
	}

	validNote := Step{ClinicalNote: &ClinicalNote{ContentType: "some_content_type"}}
	invalidNote := Step{ClinicalNote: &ClinicalNote{}}

	twoPersons := &Persons{"first": {}, "second": {}}
	usePatientFirst := Step{UsePatient: &UsePatient{Patient: PatientID("first")}}
	usePatientSecond := Step{UsePatient: &UsePatient{Patient: PatientID("second")}}
	twoPersonsOneInvalid := &Persons{"first": {Gender: "invalid"}, "second": {}}

	ctx := context.Background()
	doctors, err := doctor.LoadDoctors(ctx, test.DoctorsConfigTest)
	if err != nil {
		t.Fatalf("LoadDoctors(%s) failed with %v", test.DoctorsConfigTest, err)
	}

	id1 := "id-1"
	first1 := "firstname-1"
	last1 := "surname-1"
	wantConsultant := &Consultant{ID: &id1, FirstName: &first1, Surname: &last1}
	consultantMissingFields := &Consultant{FirstName: &first1, Surname: &last1}
	invalidConsultant := &Consultant{ID: &first1}

	existingDoctor := doctors.GetByID(id1)
	if existingDoctor == nil {
		t.Fatal("existingDoctor is <nil>.")
	}
	if got, want := existingDoctor.FirstName, first1; got != want {
		t.Errorf("existingDoctor.FirstName=%v, want %v", got, want)
	}
	if got, want := existingDoctor.Surname, last1; got != want {
		t.Errorf("existingDoctor.Surname=%v, want %v", got, want)
	}
	if doctors.GetByID(first1) != nil {
		t.Fatal("doctors.GetByID(first1) is something, want <nil>.")
	}

	cases := []struct {
		pathway *Pathway
		wantErr bool
	}{
		// TimeFromNow must be set and negative in History.
		{pathway: &Pathway{History: []Step{resultNegativeTimeFromNow}}, wantErr: false},
		{pathway: &Pathway{History: []Step{{Result: &Results{OrderProfile: "profile"}, Parameters: &Parameters{TimeFromNow: &oneHour}}}}, wantErr: true},
		{pathway: &Pathway{History: []Step{{Result: &Results{OrderProfile: "profile"}, Parameters: &Parameters{}}}}, wantErr: true},
		{pathway: &Pathway{History: []Step{{Result: &Results{OrderProfile: "profile"}}}}, wantErr: true},
		// ...unless the step is a UsePatient step.
		{pathway: &Pathway{History: []Step{usePatientFirst}}, wantErr: false},
		// Autogenerated steps cannot be in the History.
		{pathway: &Pathway{History: []Step{{AutoGenerate: &AutoGenerate{From: &twoHoursAgo, To: &oneHourAgo, Every: &second}}}}, wantErr: true},
		// TimeFromNow is not allowed in Pathway.
		{pathway: &Pathway{Pathway: []Step{{Result: &Results{OrderProfile: "profile"}, Parameters: &Parameters{TimeFromNow: &twoHoursAgo}}}}, wantErr: true},
		{pathway: &Pathway{Pathway: []Step{{Result: &Results{OrderProfile: "profile"}, Parameters: &Parameters{TimeFromNow: &oneHour}}}}, wantErr: true},
		// Delays are not allowed in History.
		{pathway: &Pathway{History: []Step{{Delay: &Delay{From: oneHour, To: twoHours}}}}, wantErr: true},
		{pathway: &Pathway{Pathway: []Step{{Delay: &Delay{From: oneHour, To: twoHours}}}}, wantErr: false},

		// OrderID only allowed if there is something to link to
		// The step in the pathway that uses OrderID first time, has to specify order profile.
		// Every subsequent step that relates to the same  OrderID may - but doesn't have to - have the order profile
		// specified, too. If specified, it needs to match the original order profile.
		{pathway: &Pathway{Pathway: []Step{{Order: &Order{OrderID: "order1", OrderProfile: "profile"}}, {Result: &Results{OrderID: "order1"}}}}, wantErr: false},
		{pathway: &Pathway{Pathway: []Step{{Order: &Order{OrderID: "order1", OrderProfile: "profile"}}, {Result: &Results{OrderID: "order1", OrderProfile: "profile"}}}}, wantErr: false},
		{pathway: &Pathway{Pathway: []Step{{Order: &Order{OrderID: "order1", OrderProfile: "profile"}}, {Result: &Results{OrderID: "order1", OrderProfile: "profile2"}}}}, wantErr: true},
		{pathway: &Pathway{Pathway: []Step{{Order: &Order{OrderID: "order1"}}, {Result: &Results{OrderID: "order1", OrderProfile: "profile"}}}}, wantErr: true},
		{pathway: &Pathway{Pathway: []Step{{Result: &Results{OrderID: "order1", OrderProfile: "profile"}}, {Result: &Results{OrderID: "order1"}}}}, wantErr: false},
		{pathway: &Pathway{Pathway: []Step{{Result: &Results{OrderID: "order1", OrderProfile: "profile"}}, {Result: &Results{OrderID: "order1", OrderProfile: "profile"}}}}, wantErr: false},
		{pathway: &Pathway{Pathway: []Step{{Result: &Results{OrderID: "order1", OrderProfile: "profile"}}, {Result: &Results{OrderID: "order1", OrderProfile: "profile2"}}}}, wantErr: true},
		{pathway: &Pathway{Pathway: []Step{{Result: &Results{OrderID: "order1"}}, {Result: &Results{OrderID: "order1", OrderProfile: "profile"}}}}, wantErr: true},
		{pathway: &Pathway{History: []Step{{Order: &Order{OrderID: "order1", OrderProfile: "profile"}, Parameters: &Parameters{TimeFromNow: &twoHoursAgo}}}, Pathway: []Step{{Result: &Results{OrderID: "order1"}}}}, wantErr: false},
		{pathway: &Pathway{History: []Step{{Order: &Order{OrderID: "order1", OrderProfile: "profile"}, Parameters: &Parameters{TimeFromNow: &twoHoursAgo}}}, Pathway: []Step{{Result: &Results{OrderID: "order1", OrderProfile: "profile"}}}}, wantErr: false},
		{pathway: &Pathway{History: []Step{{Order: &Order{OrderID: "order1", OrderProfile: "profile"}, Parameters: &Parameters{TimeFromNow: &twoHoursAgo}}}, Pathway: []Step{{Result: &Results{OrderID: "order1", OrderProfile: "profile2"}}}}, wantErr: true},
		{pathway: &Pathway{History: []Step{{Order: &Order{OrderID: "order1"}, Parameters: &Parameters{TimeFromNow: &twoHoursAgo}}}, Pathway: []Step{{Result: &Results{OrderID: "order1", OrderProfile: "profile"}}}}, wantErr: true},
		// It's ok to have an order ID without results now.
		{pathway: &Pathway{Pathway: []Step{{Order: &Order{OrderID: "order1", OrderProfile: "profile"}}, {Result: &Results{OrderID: "order2", OrderProfile: "profile"}}}}, wantErr: false},
		{pathway: &Pathway{Pathway: []Step{{Order: &Order{OrderID: "order1", OrderProfile: "profile"}}}}, wantErr: false},
		{pathway: &Pathway{Pathway: []Step{{Order: &Order{OrderProfile: "profile"}}}}, wantErr: false},
		// AddPerson needs to be the first step.
		{pathway: &Pathway{Pathway: []Step{addPerson}}, wantErr: false},
		{pathway: &Pathway{Pathway: []Step{admit, addPerson}}, wantErr: true},
		// UsePatient requires a patient.
		{pathway: &Pathway{Pathway: []Step{{UsePatient: &UsePatient{Patient: "123"}}}}, wantErr: false},
		{pathway: &Pathway{Pathway: []Step{{UsePatient: &UsePatient{}}}}, wantErr: true},
		// UsePatient cannot have DelayMessage.
		{pathway: &Pathway{Pathway: []Step{{UsePatient: &UsePatient{Patient: "123"}, Parameters: &Parameters{DelayMessage: &Delay{}}}}}, wantErr: true},
		// History steps do not have to be ordered by TimeFromNow.
		{pathway: &Pathway{History: []Step{stepOneHourAgo, stepTwoHoursAgo}}, wantErr: false},
		{pathway: &Pathway{Pathway: []Step{{UsePatient: &UsePatient{Patient: Current}}}}, wantErr: true},
		// All persons in Persons need to be used in UsePatient steps. Steps in History count too.
		{pathway: &Pathway{Persons: twoPersons, Pathway: []Step{usePatientFirst, usePatientSecond}}, wantErr: false},
		{pathway: &Pathway{Persons: twoPersons, Pathway: []Step{usePatientFirst, {UsePatient: &UsePatient{Patient: PatientID("something-else")}}}}, wantErr: true},
		{pathway: &Pathway{Persons: twoPersons, Pathway: []Step{usePatientFirst}}, wantErr: true},
		{pathway: &Pathway{Persons: twoPersons, History: []Step{usePatientSecond}, Pathway: []Step{usePatientFirst}}, wantErr: false},
		// If there's a Persons section, the UsePatient step must be the first one...
		{pathway: &Pathway{Persons: twoPersons, Pathway: []Step{usePatientFirst, usePatientSecond, result}}, wantErr: false},
		{pathway: &Pathway{Persons: twoPersons, Pathway: []Step{result, usePatientFirst, usePatientSecond}}, wantErr: true},
		{pathway: &Pathway{Persons: twoPersons, History: []Step{usePatientFirst, usePatientSecond, resultNegativeTimeFromNow}}, wantErr: false},
		{pathway: &Pathway{Persons: twoPersons, History: []Step{resultNegativeTimeFromNow, usePatientFirst, usePatientSecond}}, wantErr: true},
		// ... unless there's just one Person.
		{pathway: &Pathway{Persons: twoPersons, Pathway: []Step{usePatientFirst, usePatientSecond}}, wantErr: false},
		// All persons in Persons need to be valid.
		{pathway: &Pathway{Persons: twoPersonsOneInvalid, Pathway: []Step{usePatientFirst, usePatientSecond}}, wantErr: true},
		{pathway: &Pathway{Persons: nil, Pathway: []Step{admit, discharge}}, wantErr: false},
		{pathway: &Pathway{Persons: nil, Pathway: []Step{admit, discharge}, Consultant: wantConsultant}, wantErr: false},
		{pathway: &Pathway{Persons: nil, Pathway: []Step{admit, discharge}, Consultant: consultantMissingFields}, wantErr: false},
		{pathway: &Pathway{Persons: nil, Pathway: []Step{admit, discharge}, Consultant: invalidConsultant}, wantErr: true},
		{pathway: &Pathway{Pathway: []Step{{Order: &Order{OrderID: "order1", OrderProfile: "profile"}}, {Result: &Results{OrderID: "order2", OrderProfile: "profile"}}, validNote}}, wantErr: false},
		{pathway: &Pathway{Pathway: []Step{{Order: &Order{OrderID: "order1", OrderProfile: "profile"}}, {Result: &Results{OrderID: "order2", OrderProfile: "profile"}}, invalidNote}}, wantErr: true},
		{pathway: &Pathway{Pathway: []Step{{Document: &Document{}}}}, wantErr: false},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("id:%d-pathway:%+v-valid:%t", i, tc.pathway, !tc.wantErr), func(t *testing.T) {
			tc.pathway.Init(fmt.Sprintf("test-pathway-%d", i))
			err := tc.pathway.Valid(defaultClock, emptyOrderProfiles, doctors, defaultLocationManager, defaultValid)
			if gotErr := err != nil; gotErr != tc.wantErr {
				t.Errorf("[%+v].Valid(%v, %v, %v, _) got err %v; want err? %t", tc.pathway, defaultClock, emptyOrderProfiles, doctors, err, tc.wantErr)
			}
			// When we validate the pathway, we update the consultant in the pathway. This is important
			// to check that all missing fields of the consultant have been added if the pathway is valid.
			if tc.pathway.Consultant != nil && !tc.wantErr {
				if diff := cmp.Diff(wantConsultant, tc.pathway.Consultant); diff != "" {
					t.Errorf("pathway.Consultant got diff (-want, +got):\n%s", diff)
				}
			}
		})
	}
}

func TestPathwayValidFn(t *testing.T) {
	pathway := &Pathway{Pathway: []Step{{Result: &Results{OrderProfile: "profile"}}}}
	pathway.Init("test-pathway")
	cases := []struct {
		desc    string
		fn      func(*Pathway) error
		wantErr bool
	}{
		{
			desc: "valid",
			fn: func(pathway *Pathway) error {
				return nil
			},
			wantErr: false,
		},
		{
			desc: "invalid",
			fn: func(pathway *Pathway) error {
				return errors.New("invalid")
			},
			wantErr: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			err := pathway.Valid(defaultClock, emptyOrderProfiles, emptyDoctors, defaultLocationManager, tt.fn)
			if gotErr := err != nil; gotErr != tt.wantErr {
				t.Errorf("[%+v].Valid(_, _, _, _) got err %v; want err? %t", pathway, err, tt.wantErr)
			}
		})
	}
}

func TestPathwayValidPercentage(t *testing.T) {
	cases := []struct {
		percentage *Percentage
		wantErr    bool
	}{
		// Percentage must be >0, and with no more than maxSignificantDigits decimal digits.
		{percentage: nil, wantErr: false},
		{percentage: NewPercentage(1), wantErr: false},
		{percentage: NewPercentage(0), wantErr: false},
		{percentage: NewPercentage(-1), wantErr: true},
		{percentage: NewPercentage(1.1), wantErr: false},
		{percentage: NewPercentage(1.111), wantErr: false},
		{percentage: NewPercentage(1.1111), wantErr: true},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("id:%d-percentage:%v", i, tc.percentage), func(t *testing.T) {
			pathway := &Pathway{
				Percentage: tc.percentage,
				Pathway:    []Step{{Result: &Results{OrderProfile: "profile"}}},
			}
			pathway.Init(pathwayName)

			err := pathway.Valid(defaultClock, emptyOrderProfiles, emptyDoctors, defaultLocationManager, defaultValid)
			if gotErr := err != nil; gotErr != tc.wantErr {
				t.Errorf("[%+v].Valid(_, _, _, _) got err %v; want err? %t", pathway, err, tc.wantErr)
			}
		})
	}
}

func TestPathwayValidPerson(t *testing.T) {
	cases := []struct {
		name    string
		persons *Persons
		wantErr bool
	}{
		// Person only allows genders of F, M, or empty.
		{name: "no person", persons: nil, wantErr: false},
		{name: "Gender F", persons: &Persons{"main_patient": Person{Gender: "F"}}, wantErr: false},
		{name: "Gender M", persons: &Persons{"main_patient": Person{Gender: "M"}}, wantErr: false},
		{name: "Gender Empty", persons: &Persons{"main_patient": Person{Gender: ""}}, wantErr: false},
		{name: "Gender Random", persons: &Persons{"main_patient": Person{Gender: "RANDOM"}}, wantErr: false},
		{name: "Gender Invalid", persons: &Persons{"main_patient": Person{Gender: "invalid"}}, wantErr: true},

		// Addresses must provide FirstLine if SecondLine is provided.
		{name: "Address first line set", persons: &Persons{"main_patient": Person{Address: &Address{FirstLine: "6 Pancras Square"}}}, wantErr: false},
		{name: "Address second line set", persons: &Persons{"main_patient": Person{Address: &Address{SecondLine: "King's Cross"}}}, wantErr: true},

		// If AllRandom is set, no other field can be set.
		{name: "Address AllRandom Postcode set", persons: &Persons{"main_patient": Person{Address: &Address{AllRandom: true, Postcode: "NW1 4AB"}}}, wantErr: true},
		{name: "Address AllRandom=false Postcode set", persons: &Persons{"main_patient": Person{Address: &Address{AllRandom: false, Postcode: "NW1 4AB"}}}, wantErr: false},
		// DayOfYear for Age must be in [0, 365]
		{name: "Age DayOfYear 0", persons: &Persons{"main_patient": Person{Age: &Age{DayOfYear: 0}}}, wantErr: false},
		{name: "Age DayOfYear 365", persons: &Persons{"main_patient": Person{Age: &Age{DayOfYear: 365}}}, wantErr: false},
		{name: "Age DayOfYear -1", persons: &Persons{"main_patient": Person{Age: &Age{DayOfYear: -1}}}, wantErr: true},
		{name: "Age DayOfYear 366", persons: &Persons{"main_patient": Person{Age: &Age{DayOfYear: 366}}}, wantErr: true},
		{name: "Age 20-30", persons: &Persons{"main_patient": Person{Age: &Age{From: 20, To: 30}}}, wantErr: false},
		{name: "Age 30-30", persons: &Persons{"main_patient": Person{Age: &Age{From: 30, To: 30}}}, wantErr: false},
		{name: "Age 30-20", persons: &Persons{"main_patient": Person{Age: &Age{From: 30, To: 20}}}, wantErr: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			pathway := &Pathway{
				Persons: tc.persons,
				Pathway: []Step{{Order: &Order{OrderProfile: "profile"}}},
			}
			pathway.Init(pathwayName)

			err := pathway.Valid(defaultClock, emptyOrderProfiles, emptyDoctors, defaultLocationManager, defaultValid)
			if gotErr := err != nil; gotErr != tc.wantErr {
				t.Errorf("[%+v].Valid(_, _, _, _) got err %v; want err? %t", pathway, err, tc.wantErr)
			}
		})
	}
}

func TestPathwayValidConsultant(t *testing.T) {
	ctx := context.Background()
	doctorsFile := []byte(`
- id: "id-1"
  surname: "surname-1"
  firstname: "firstname-1"
  prefix: "prefix-1"
  specialty: "specialty-1"
- id: "id-2"
  surname: "surname-2"
  firstname: "firstname-2"
  prefix: "prefix-2"
  specialty: "specialty-2"`)

	fName := testwrite.BytesToFile(t, doctorsFile)

	doctors, err := doctor.LoadDoctors(ctx, fName)
	if err != nil {
		t.Fatalf("LoadDoctors(%s) failed with %v", string(doctorsFile), err)
	}
	randomID := "RANDOM"

	id1 := "id-1"
	first1 := "firstname-1"
	last1 := "surname-1"
	prefix1 := "prefix-1"

	id2 := "id-2"
	last2 := "surname-2"

	customID := "custom-id"
	customPrefix := "custom"
	emptyPrefix := ""

	c1 := &Consultant{
		ID:        &id1,
		Surname:   &last1,
		FirstName: &first1,
		Prefix:    &prefix1,
	}

	invalidID := "invalid-id"

	cases := []struct {
		name    string
		c       *Consultant
		want    *Consultant
		wantID  *string
		wantErr bool
	}{
		{
			name:    "Nil consultant",
			c:       nil,
			want:    nil,
			wantErr: false,
		}, {
			name:    "Consultant ID AND name matched",
			c:       &Consultant{ID: &id1, FirstName: &first1, Surname: &last1},
			want:    c1,
			wantID:  c1.ID,
			wantErr: false,
		}, {
			name:    "Consultant ID AND name matched; prefix not changed",
			c:       &Consultant{ID: &id1, FirstName: &first1, Surname: &last1, Prefix: &customPrefix},
			want:    &Consultant{ID: &id1, FirstName: &first1, Surname: &last1, Prefix: &customPrefix},
			wantID:  c1.ID,
			wantErr: false,
		}, {
			name:    "Consultant ID AND name matched; empty prefix not changed",
			c:       &Consultant{ID: &id1, FirstName: &first1, Surname: &last1, Prefix: &emptyPrefix},
			want:    &Consultant{ID: &id1, FirstName: &first1, Surname: &last1, Prefix: &emptyPrefix},
			wantID:  c1.ID,
			wantErr: false,
		}, {
			name:    "Consultant validation failed. Either ID or First Name and Surname must be provided for a consultant to be valid.",
			c:       &Consultant{},
			want:    nil,
			wantErr: true,
		}, {
			name:    "Consultant ID matched but name did not",
			c:       &Consultant{ID: &id2, FirstName: &first1, Surname: &last1},
			want:    nil,
			wantErr: true,
		}, {
			name:    "Consultant name matched but ID did not",
			c:       &Consultant{ID: &invalidID, FirstName: &first1, Surname: &last1},
			want:    nil,
			wantErr: true,
		}, {
			name:    "No ID; Consultant name found (already had ID)",
			c:       &Consultant{FirstName: &first1, Surname: &last1},
			want:    c1,
			wantID:  c1.ID,
			wantErr: false,
		}, {
			name:    "No ID; Consultant name found (already had ID); prefix not changes",
			c:       &Consultant{FirstName: &first1, Surname: &last1, Prefix: &customPrefix},
			want:    &Consultant{ID: &id1, FirstName: &first1, Surname: &last1, Prefix: &customPrefix},
			wantID:  c1.ID,
			wantErr: false,
		}, {
			name:    "No ID; Consultant name not found (generating a new ID)",
			c:       &Consultant{FirstName: &first1, Surname: &last2},
			want:    &Consultant{FirstName: &first1, Surname: &last2},
			wantID:  &randomID,
			wantErr: false,
		}, {
			name:    "No ID; Consultant name not found (generating a new ID); custom prefix",
			c:       &Consultant{FirstName: &first1, Surname: &last2, Prefix: &customPrefix},
			want:    &Consultant{FirstName: &first1, Surname: &last2, Prefix: &customPrefix},
			wantID:  &randomID,
			wantErr: false,
		}, {
			name:    "Custom ID; Consultant name or ID not found",
			c:       &Consultant{ID: &customID, FirstName: &first1, Surname: &last2, Prefix: &customPrefix},
			want:    &Consultant{ID: &customID, FirstName: &first1, Surname: &last2, Prefix: &customPrefix},
			wantID:  &customID,
			wantErr: false,
		}, {
			name:    "Consultant ID not found. If the consultant name is not specified, the ID must be an existing ID",
			c:       &Consultant{ID: &invalidID},
			want:    nil,
			wantErr: true,
		}, {
			name:    "No name; Consultant ID found",
			c:       &Consultant{ID: &id1},
			want:    c1,
			wantID:  c1.ID,
			wantErr: false,
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d-%s", i, tc.name), func(t *testing.T) {
			pathway := &Pathway{
				Consultant: tc.c,
				Pathway:    []Step{{Order: &Order{OrderProfile: "profile"}}},
			}
			pathway.Init(pathwayName)

			err := pathway.Valid(defaultClock, emptyOrderProfiles, doctors, defaultLocationManager, defaultValid)
			if gotErr := err != nil; gotErr != tc.wantErr {
				t.Errorf("[%+v].Valid(_, _, _, _) got err %v; want err? %t", pathway, err, tc.wantErr)
			}

			got := pathway.Consultant
			if diff := cmp.Diff(tc.want, got, cmpopts.IgnoreFields(Consultant{}, "ID")); diff != "" {
				t.Errorf("pathway.Consultant got diff (-want, +got):\n%s", diff)
			}
			if got == nil {
				return
			}
			if tc.wantID == &randomID {
				if got.ID == nil || *got.ID == "" {
					t.Errorf("pathway.Consultant.ID=%v, want not-nil (autogenerated)", got.ID)
				}
			} else {
				if diff := cmp.Diff(tc.wantID, got.ID); diff != "" {
					t.Errorf("pathway.Consultant.ID got diff (-want, +got):\n%s", diff)
				}
			}
		})
	}
}
