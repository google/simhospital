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

package orderprofile

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/constants"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/test"
	"github.com/google/simhospital/pkg/test/testwrite"
)

var (
	hl7ConfigFile = []byte(`
coding_system: "WinPath"
`)

	ureaOP = []byte(`
UREA AND ELECTROLYTES:
  universal_service_id: lpdc-3969
  test_types:
    Creatinine:
      id: lpdc-2012
      ref_range: 49 - 92
      unit: UMOLL
      value: '70'
      value_type: NM`)

	ureaOPValPrefix = []byte(`
UREA AND ELECTROLYTES:
  universal_service_id: lpdc-3969
  test_types:
    Creatinine:
      id: lpdc-2012
      ref_range: 49 - 92
      unit: UMOLL
      value: '>70'
      value_type: NM`)

	ohOPTX = []byte(`
17-OH Prog:
  universal_service_id: OHPROG
  test_types:
    17-Hydroxy Progesterone:
      id: OHPROG
      value_type: TX
      value: See note
      unit: NMOLL
      ref_range: <=9.6^^<=9.6`)

	ohOPCE = []byte(`
17-OH Prog:
  universal_service_id: OHPROG
  test_types:
    17-Hydroxy Progesterone:
      id: OHPROG
      value_type: CE
      value: See note`)

	ohOPNM = []byte(`
17-OH Prog:
  universal_service_id: OHPROG
  test_types:
    17-Hydroxy Progesterone:
      id: OHPROG
      value_type: NM
      value: See note
      unit: NMOLL
      ref_range: 0 - 0.45`)

	ohOPNMInvalidRefRange = []byte(`
17-OH Prog:
  universal_service_id: OHPROG
  test_types:
    17-Hydroxy Progesterone:
      id: OHPROG
      value_type: NM
      value: 0.20
      unit: NMOLL
      ref_range: "-"`)

	ohOPNMNoRefRange = []byte(`
17-OH Prog:
  universal_service_id: OHPROG
  test_types:
    17-Hydroxy Progesterone:
      id: OHPROG
      value_type: NM
      value: 2.50
      unit: NMOLL`)

	vitalSignsTT = []byte(`
Vital Signs:
  universal_service_id: us-0005
  test_types:
    MDC_TEMP:
      id: tt-0005-09
      coding_system: MDC
      ref_range: 36-38
      unit: MDC_DIM_DEGC
      value: 37.5
      value_type: NM`)

	vitalSignsOP = []byte(`
Vital Signs:
  universal_service_id: us-0005
  coding_system: MDC
  test_types:
    MDC_TEMP:
      id: tt-0005-09
      ref_range: 36-38
      unit: MDC_DIM_DEGC
      value: 37.5
      value_type: NM`)

	invalidOP = []byte(`
17-OH Prog:
  service_id: OHPROG
  test_types:
    17-Hydroxy Progesterone:
      tt_id: OHPROG
      value: 2.50
      unit: NMOLL`)
)

func TestLoad(t *testing.T) {
	ctx := context.Background()
	hl7Config := loadHL7Config(ctx, t)
	cases := []struct {
		name          string
		opFileContent []byte
		wantLoadErr   bool
		opName        string
		ttName        string
		wantUS        ir.CodedElement
		wantTT        *TestType
	}{
		{
			name:          "Default Coding System",
			opFileContent: ureaOP,
			opName:        "UREA AND ELECTROLYTES",
			ttName:        "Creatinine",
			wantUS:        ir.CodedElement{ID: "lpdc-3969", Text: "UREA AND ELECTROLYTES", CodingSystem: "WinPath"},
			wantTT: &TestType{
				Name:      ir.CodedElement{ID: "lpdc-2012", Text: "Creatinine", CodingSystem: "WinPath"},
				Unit:      "UMOLL",
				ValueType: "NM",
				RefRange:  "49 - 92",
			},
		}, {
			name:          "Custom Coding System for test type",
			opFileContent: vitalSignsTT,
			opName:        "Vital Signs",
			ttName:        "MDC_TEMP",
			wantUS:        ir.CodedElement{ID: "us-0005", Text: "Vital Signs", CodingSystem: "WinPath"},
			wantTT: &TestType{
				Name:      ir.CodedElement{ID: "tt-0005-09", Text: "MDC_TEMP", CodingSystem: "MDC"},
				Unit:      "MDC_DIM_DEGC",
				ValueType: "NM",
				RefRange:  "36-38",
			},
		}, {
			name:          "Custom Coding System for order profile",
			opFileContent: vitalSignsOP,
			opName:        "Vital Signs",
			ttName:        "MDC_TEMP",
			wantUS:        ir.CodedElement{ID: "us-0005", Text: "Vital Signs", CodingSystem: "MDC"},
			wantTT: &TestType{
				Name:      ir.CodedElement{ID: "tt-0005-09", Text: "MDC_TEMP", CodingSystem: "WinPath"},
				Unit:      "MDC_DIM_DEGC",
				ValueType: "NM",
				RefRange:  "36-38",
			},
		}, {
			name:          "Invalid Order Profile",
			opFileContent: invalidOP,
			wantLoadErr:   true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fName := testwrite.BytesToFile(t, tc.opFileContent)

			ops, err := Load(ctx, fName, hl7Config)
			if gotErr := err != nil; gotErr != tc.wantLoadErr {
				t.Fatalf("Load(%s, %+v) got err %v, want err? %t", fName, hl7Config, err, tc.wantLoadErr)
			}
			if tc.wantLoadErr || err != nil {
				return
			}

			op, ok := ops.Get(tc.opName)
			if !ok {
				t.Fatalf("Get(%s) got ok = %t, want true", tc.opName, ok)
			}
			if diff := cmp.Diff(tc.wantUS, op.UniversalService); diff != "" {
				t.Errorf("[%+v].UniversalService -want, +got:\n%s", op, diff)
			}

			tt, ok := op.TestTypes[tc.ttName]
			if !ok {
				t.Fatalf("[%+v][%s] got ok = %t, want true", op, tc.ttName, ok)
			}
			if diff := cmp.Diff(tc.wantTT, tt, cmpopts.IgnoreUnexported(TestType{}), cmpopts.IgnoreFields(TestType{}, "ValueGenerator")); diff != "" {
				t.Errorf("[%+v][%s] got -want, +got:\n%s", op, tc.ttName, diff)
			}
		})
	}
}

func TestRandomisedValueWithFlag_NumericalValue(t *testing.T) {
	ctx := context.Background()
	hl7Config := loadHL7Config(ctx, t)
	ureaKey := "UREA AND ELECTROLYTES"
	creaKey := "Creatinine"

	cases := []struct {
		name          string
		opFileContent []byte
		wantPrefix    string
	}{
		{
			name:          "Numerical value",
			opFileContent: ureaOP,
			wantPrefix:    "",
		}, {
			name:          "Numerical value with prefix",
			opFileContent: ureaOPValPrefix,
			wantPrefix:    ">",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fName := testwrite.BytesToFile(t, tc.opFileContent)

			ops, err := Load(ctx, fName, hl7Config)
			if err != nil {
				t.Fatalf("Load(%s, %+v) failed with %v", fName, hl7Config, err)
			}

			op, ok := ops.Get(ureaKey)
			if !ok {
				t.Fatalf("Get(%s) got ok = %t, want true", ureaKey, ok)
			}
			tt, ok := op.TestTypes[creaKey]
			if !ok {
				t.Fatalf("[%+v][%s] got ok = %t, want true", op, creaKey, ok)
			}

			randomTypes := []struct {
				randomType string
				wantFrom   float64
				wantTo     float64
				wantFlag   constants.AbnormalFlag
			}{
				{randomType: constants.NormalValue, wantFrom: 49, wantTo: 92, wantFlag: constants.AbnormalFlagEmpty},
				{randomType: constants.AbnormalHigh, wantFrom: 92, wantTo: 920, wantFlag: constants.AbnormalFlagHigh},
				{randomType: constants.AbnormalLow, wantFrom: 0, wantTo: 49, wantFlag: constants.AbnormalFlagLow},
			}
			for _, rType := range randomTypes {
				t.Run(rType.randomType, func(t *testing.T) {
					gotValue, gotFlag, err := tt.RandomisedValueWithFlag(rType.randomType)
					if err != nil {
						t.Fatalf("[%+v].RandomisedValueWithFlag(%v) failed with %v", tt, rType.randomType, err)
					}
					if !strings.HasPrefix(gotValue, tc.wantPrefix) {
						t.Errorf("[%+v].RandomisedValueWithFlag(%q) got value=%q, want value with prefix %q", tt, rType.randomType, gotValue, tc.wantPrefix)
					}
					gotValue = strings.TrimPrefix(gotValue, tc.wantPrefix)
					gotf, err := strconv.ParseFloat(gotValue, 64)
					if err != nil {
						t.Fatalf("ParseFloat(%q, 64) failed with %v", gotValue, err)
					}
					if gotf <= rType.wantFrom || gotf >= rType.wantTo {
						t.Errorf("[%+v].RandomisedValueWithFlag(%q) got value=%q, want in range (%f, %f)", tt, rType.randomType, gotValue, rType.wantFrom, rType.wantTo)
					}
					if gotFlag != rType.wantFlag {
						t.Errorf("[%+v].RandomisedValueWithFlag(%q) got flag=%q, want %q", tt, rType.randomType, gotFlag, rType.wantFlag)
					}
				})
			}
		})
	}
}

func TestRandomisedValueWithFlag_NonNumericalValue(t *testing.T) {
	ctx := context.Background()
	hl7Config := loadHL7Config(ctx, t)
	ohKey := "17-OH Prog"
	hydKey := "17-Hydroxy Progesterone"
	wantUS := ir.CodedElement{ID: "OHPROG", Text: "17-OH Prog", CodingSystem: "WinPath"}
	cases := []struct {
		name          string
		opFileContent []byte
		wantTT        *TestType
		wantValue     string
	}{
		{
			name:          "TX value type",
			opFileContent: ohOPTX,
			wantTT: &TestType{
				Name:      ir.CodedElement{ID: "OHPROG", Text: "17-Hydroxy Progesterone", CodingSystem: "WinPath"},
				Unit:      "NMOLL",
				ValueType: "TX",
				RefRange:  "<=9.6^^<=9.6",
			},
			wantValue: "See note",
		}, {
			name:          "CE value type",
			opFileContent: ohOPCE,
			wantTT: &TestType{
				Name:      ir.CodedElement{ID: "OHPROG", Text: "17-Hydroxy Progesterone", CodingSystem: "WinPath"},
				ValueType: "CE",
			},
			wantValue: "See note",
		}, {
			name:          "NM value type but value not numerical - default to value",
			opFileContent: ohOPNM,
			wantTT: &TestType{
				Name:      ir.CodedElement{ID: "OHPROG", Text: "17-Hydroxy Progesterone", CodingSystem: "WinPath"},
				Unit:      "NMOLL",
				ValueType: "NM",
				RefRange:  "0 - 0.45",
			},
			wantValue: "See note",
		},
	}

	for _, tc := range cases {
		fName := testwrite.BytesToFile(t, tc.opFileContent)

		ops, err := Load(ctx, fName, hl7Config)
		if err != nil {
			t.Fatalf("Load(%s, %+v) failed with %v", fName, hl7Config, err)
		}

		op, ok := ops.Get(ohKey)
		if !ok {
			t.Fatalf("Get(%s) got ok = %t, want true", ohKey, ok)
		}
		if diff := cmp.Diff(wantUS, op.UniversalService); diff != "" {
			t.Errorf("[%+v].UniversalService -want, +got:\n%s", op, diff)
		}

		tt, ok := op.TestTypes[hydKey]
		if !ok {
			t.Fatalf("[%+v][%s] got ok = %t, want true", op, hydKey, ok)
		}
		if diff := cmp.Diff(tc.wantTT, tt, cmpopts.IgnoreUnexported(TestType{}), cmpopts.IgnoreFields(TestType{}, "ValueGenerator")); diff != "" {
			t.Errorf("[%+v][%s] got -want, +got:\n%s", op, hydKey, diff)
		}

		randomTypes := []struct {
			randomType string
			wantFlag   constants.AbnormalFlag
		}{
			{randomType: constants.NormalValue, wantFlag: constants.AbnormalFlagEmpty},
			{randomType: constants.AbnormalHigh, wantFlag: constants.AbnormalFlagHigh},
			{randomType: constants.AbnormalLow, wantFlag: constants.AbnormalFlagLow},
		}

		for _, rType := range randomTypes {
			t.Run(fmt.Sprintf("%s-%s", tc.name, rType.randomType), func(t *testing.T) {
				gotValue, gotFlag, err := tt.RandomisedValueWithFlag(rType.randomType)
				if err != nil {
					t.Fatalf("[%+v].RandomisedValueWithFlag(%v) failed with %v", tt, rType.randomType, err)
				}
				if gotValue != tc.wantValue {
					t.Errorf("[%+v].RandomisedValueWithFlag(%q) got value=%q, want %q", tt, rType.randomType, gotValue, tc.wantValue)
				}
				if gotFlag != rType.wantFlag {
					t.Errorf("[%+v].RandomisedValueWithFlag(%q) got flag=%q, want %q", tt, rType.randomType, gotFlag, rType.wantFlag)
				}
			})
		}
	}
}

func TestRandomisedValueWithFlagRange_Error(t *testing.T) {
	ctx := context.Background()
	hl7Config := loadHL7Config(ctx, t)
	ohKey := "17-OH Prog"
	hydKey := "17-Hydroxy Progesterone"

	tests := []struct {
		name          string
		opFileContent []byte
		wantValue     string
	}{{
		name:          "Invalid range",
		opFileContent: ohOPNMInvalidRefRange,
		wantValue:     "0.20",
	}, {
		name:          "No range",
		opFileContent: ohOPNMNoRefRange,
		wantValue:     "2.50",
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			fName := testwrite.BytesToFile(t, tc.opFileContent)

			ops, err := Load(ctx, fName, hl7Config)
			if err != nil {
				t.Fatalf("Load(%s, %+v) failed with %v", fName, hl7Config, err)
			}
			op, ok := ops.Get(ohKey)
			if !ok {
				t.Fatalf("Get(%s) got ok = %t, want true", ohKey, ok)
			}
			tt, ok := op.TestTypes[hydKey]
			if !ok {
				t.Fatalf("[%+v][%s] got ok = %t, want true", op, hydKey, ok)
			}

			// The value_type is NM, but we cannot parse the range, so always use default value.
			gotValue, gotFlag, err := tt.RandomisedValueWithFlag(constants.NormalValue)
			if err != nil {
				t.Fatalf("RandomisedValueWithFlag(%v) failed with %v", constants.NormalValue, err)
			}
			if gotValue != tc.wantValue {
				t.Errorf("RandomisedValueWithFlag(%v) value=%v, want %v", constants.NormalValue, gotValue, tc.wantValue)
			}
			if wantFlag := constants.AbnormalFlagEmpty; gotFlag != wantFlag {
				t.Errorf("RandomisedValueWithFlag(%v) abnormalFlag=%v, want %v", constants.NormalValue, gotFlag, wantFlag)
			}

			// Since we cannot parse the ref range, it is not possible to generate high / low values in this case.
			randomisedValues := []string{constants.AbnormalHigh, constants.AbnormalLow}

			for _, val := range randomisedValues {
				t.Run(fmt.Sprintf("%s-%s", tc.name, val), func(t *testing.T) {
					if _, _, err := tt.RandomisedValueWithFlag(val); err == nil {
						t.Errorf("RandomisedValueWithFlag(%v) returned nil error, want non-nil error", val)
					}
				})
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	ctx := context.Background()
	hl7Config := loadHL7Config(ctx, t)
	fName := testwrite.BytesToFile(t, ureaOP)

	orderProfiles, err := Load(ctx, fName, hl7Config)
	if err != nil {
		t.Fatalf("Load(%s, %+v) failed with %v", fName, hl7Config, err)
	}

	cases := []struct {
		name  string
		input string
		want  *ir.CodedElement
	}{
		{
			name:  "Existing order profile",
			input: "UREA AND ELECTROLYTES",
			want:  &ir.CodedElement{ID: "lpdc-3969", Text: "UREA AND ELECTROLYTES", CodingSystem: "WinPath"},
		}, {
			name:  "Non-existing order profile",
			input: "other",
			want:  &ir.CodedElement{ID: "other", Text: "other"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := orderProfiles.Generate(tc.input)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Generate(%v) -want, +got:\n%s", tc.input, diff)
			}
		})
	}
}

func TestGenerateRandom(t *testing.T) {
	ctx := context.Background()
	hl7Config := loadHL7Config(ctx, t)
	orderProfiles, err := Load(ctx, test.ComplexOrderProfilesConfigTest, hl7Config)
	if err != nil {
		t.Fatalf("Load(%s, %+v) failed with %v", test.ComplexOrderProfilesConfigTest, hl7Config, err)
	}
	input := constants.RandomString
	allOrderProfiles := []string{
		"UREA AND ELECTROLYTES",
		"17-OH Prog",
		"17-OH Prog CE",
		"ACETYCHOLINE RECEPTOR AB",
		"ACh Receptor Ab",
		"17-OH PROGESTERONE GC-MS",
		"1HR DDAVP-HAEM A",
	}
	gotAll := make(map[string]int)
	runs := 1000
	for i := 0; i < runs; i++ {
		got := orderProfiles.Generate(input)

		if !contains(allOrderProfiles, got.Text) {
			t.Errorf("Generate(%v) got Text=%q, want one of %v", input, got.Text, allOrderProfiles)
		}
		gotAll[got.Text]++
	}
	delta := float64(runs) / 5.0
	wantFreq := float64(runs) / float64(len(allOrderProfiles))
	for _, item := range allOrderProfiles {
		if gotFreq := gotAll[item]; math.Abs(float64(gotFreq)-wantFreq) > delta {
			t.Errorf("gotAll[%s] = %d, want within %v of %f", item, gotFreq, delta, wantFreq)
		}
	}
}

func loadHL7Config(ctx context.Context, t *testing.T) *config.HL7Config {
	t.Helper()
	fConfig := testwrite.BytesToFile(t, hl7ConfigFile)
	hl7Config, err := config.LoadHL7Config(ctx, fConfig)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", fConfig, err)
	}
	return hl7Config
}

func contains(strings []string, target string) bool {
	for _, s := range strings {
		if s == target {
			return true
		}
	}
	return false
}
