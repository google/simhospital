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
	"math"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/simhospital/pkg/test/testclock"
	"github.com/google/simhospital/pkg/test/testlocation"
	"github.com/google/simhospital/pkg/test/testwrite"
)

const defaultPathwayName = "random_pathway"

var (
	tenSecondsAgo = -10 * time.Second
	july182006    = time.Date(2006, 7, 18, 0, 0, 0, 0, time.UTC)
)

func newDefaultParser(ctx context.Context, t *testing.T, now time.Time) *Parser {
	t.Helper()
	return &Parser{
		Clock:           testclock.New(now),
		OrderProfiles:   emptyOrderProfiles,
		Doctors:         emptyDoctors,
		LocationManager: testlocation.NewLocationManager(ctx, t, "ED", "Renal"),
	}
}

func TestFileExtensionValidation(t *testing.T) {
	ctx := context.Background()
	p := []byte(`
pathway1:
  persons:
    main_patient:
      gender: F
  pathway:
    - discharge: {}
`)

	cases := []struct {
		name    string
		wantErr bool
	}{{
		name: "valid*.yml",
	}, {
		name: "valid*.yaml",
	}, {
		name: "valid*.json",
	}, {
		name:    "no_extension",
		wantErr: true,
	}, {
		name:    "invalid_extension*.jpg",
		wantErr: true,
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			dir := testwrite.BytesToDir(t, []byte(p), tc.name)

			p := newDefaultParser(ctx, t, time.Now())

			pathways, err := p.ParsePathways(ctx, dir)

			// If a filename is valid, we expect its pathways to be parsed.
			// Otherwise, we expect an error because no files were parsed.
			if (err != nil) != tc.wantErr {
				t.Errorf("ParsePathways(%s) got err=%v; want error? %t", dir, err, tc.wantErr)
			}

			if err != nil || tc.wantErr {
				return
			}

			want := 1
			if got := len(pathways); got != want {
				t.Errorf("len(pathways) got %d pathways, %d 1", got, want)
			}
		})
	}
}

func TestParsePathways_Valid(t *testing.T) {
	ctx := context.Background()
	p1 := []byte(`
pathway1:
  persons:
    main_patient:
      gender: F
  pathway:
    - discharge: {}
`)

	cases := []struct {
		name         string
		p2           []byte
		wantErr      bool
		wantPathways int
	}{{
		name: "valid pathways",
		p2: []byte(`
pathway2:
 persons:
   main_patient:
     gender: F
 pathway:
   - discharge: {}
`),
		wantPathways: 2,
	}, {
		name: "multiple yml pathways in file",
		p2: []byte(`
pathway2:
 persons:
   main_patient:
     gender: F
 pathway:
   - discharge: {}

pathway3:
 persons:
   main_patient:
     gender: F
 pathway:
   - discharge: {}
`),
		wantPathways: 3,
	}, {
		name: "multiple json pathways in file",
		p2: []byte(`
{
 "pathway2": {
   "persons": {
     "main_patient": {
       "gender": "F"
     }
   },
   "pathway": [
     {
       "discharge": {}
     }
   ]
 },
 "pathway3": {
   "persons": {
     "main_patient": {
       "gender": "F"
     }
   },
   "pathway": [
     {
       "discharge": {}
     }
   ]
 }
}
`),
		wantPathways: 3,
	}, {
		name: "first yml then json in the same file",
		p2: []byte(`
pathway3:
  persons:
    main_patient:
      gender: F
  pathway:
    - discharge: {}

{
  "pathway2": {
    "persons": {
      "main_patient": {
        "gender": "F"
      }
    },
    "pathway": [
      {
        "discharge": {}
      }
    ]
  },
}
`),
		wantErr: true,
	}, {
		name: "first json then json again in the same file",
		p2: []byte(`
{
  "pathway2": {
    "persons": {
      "main_patient": {
        "gender": "F"
      }
    },
    "pathway": [
      {
        "discharge": {}
      }
    ]
  },
}

{
  "pathway_ignored": {
    "persons": {
      "main_patient": {
        "gender": "F"
      }
    },
    "pathway": [
      {
        "discharge": {}
      }
    ]
  },
}
`),
		// Only the first JSON file is loaded.
		wantPathways: 2,
	}, {
		name: "first json then yml in the same file",
		p2: []byte(`
{
  "pathway2": {
    "persons": {
      "main_patient": {
        "gender": "F"
      }
    },
    "pathway": [
      {
        "discharge": {}
      }
    ]
  },
}

pathway_ignored_1:
  persons:
    main_patient:
      gender: F
  pathway:
    - discharge: {}

{
  "pathway_ignored_2": {
    "persons": {
      "main_patient": {
        "gender": "F"
      }
    },
    "pathway": [
      {
        "discharge": {}
      }
    ]
  },
}
`),
		// Only the first JSON file is loaded.
		wantPathways: 2,
	}, {
		name: "pathway re-declared",
		p2: []byte(`
pathway1:
 persons:
   main_patient:
     gender: F
 pathway:
   - discharge: {}
`),
		wantErr: true,
	}, {
		name: "invalid pathway",
		p2: []byte(`
pathway2:
 persons:
   main_patient:
     gender: invalid
 pathway:
   - discharge: {}
`),
		wantErr: true,
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mainDir := testwrite.TempDir(t)
			testwrite.BytesToFileInExistingDir(t, p1, mainDir, "pathway1.yml")
			testwrite.BytesToFileInExistingDir(t, tc.p2, mainDir, "pathway2.yml")

			p := newDefaultParser(ctx, t, time.Now())

			pathways, err := p.ParsePathways(ctx, mainDir)
			if (err != nil) != tc.wantErr {
				t.Errorf("ParsePathways(%s) got err=%v; want error? %t", mainDir, err, tc.wantErr)
			}

			if got := len(pathways); got != tc.wantPathways {
				t.Errorf("ParsePathways(%s) got %d pathways, want %d", mainDir, got, tc.wantPathways)
			}
		})
	}
}

func TestParsePathwaysUnmarshalStrict(t *testing.T) {
	ctx := context.Background()
	p := newDefaultParser(ctx, t, time.Now())
	cases := []struct {
		name              string
		pathwayDefinition []byte
		wantErr           bool
		wantErrContains   string
	}{{
		name: "valid pathway",
		pathwayDefinition: []byte(`
test_pathway:
  persons:
    main_patient:
      gender: F
  pathway:
    - discharge: {}
`),
		wantErr: false,
	}, {
		name: "valid pathway as json",
		pathwayDefinition: []byte(`
{
  "test_pathway": {
    "persons": {
      "main_patient": {
        "gender": "F"
      }
    },
    "pathway": [
      {
        "discharge": {}
      }
    ]
  }
}`),
		wantErr: false,
	}, {
		name: "invalid yml pathway: unknown field",
		pathwayDefinition: []byte(`
test_pathway:
  unknown_field: arbitrary_string
  persons:
    main_patient:
      gender: F
  pathway:
    - discharge: {}
`),
		wantErr:         true,
		wantErrContains: "unknown_field",
	}, {
		name: "invalid json pathway: unknown field",
		pathwayDefinition: []byte(`
{
  "test_pathway": {
    "unknown_field": "arbitrary_string",
    "persons": {
      "main_patient": {
        "gender": "F"
      }
    },
    "pathway": [
      {
        "discharge": {}
      }
    ]
  }
}`),
		wantErr:         true,
		wantErrContains: "unknown_field",
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mainDir := writePathwayToDir(t, tc.pathwayDefinition)

			_, err := p.ParsePathways(ctx, mainDir)
			gotErr := err != nil
			if gotErr != tc.wantErr {
				t.Errorf("ParsePathways(%s) got err %v, want err? %t", string(tc.pathwayDefinition), err, tc.wantErr)
			}
			if !gotErr || !tc.wantErr {
				return
			}

			if !strings.Contains(err.Error(), tc.wantErrContains) {
				t.Errorf("ParsePathways(%s) got err %v, want err to contain %q", string(tc.pathwayDefinition), err, tc.wantErrContains)
			}
		})
	}
}

func TestParseSinglePathway(t *testing.T) {
	ctx := context.Background()
	delay := &Delay{
		From: time.Second,
		To:   5 * time.Second,
	}
	pathway := Pathway{
		Persons: &Persons{defaultPatientID: {}},
		Pathway: []Step{
			{Admission: &Admission{Loc: "Renal"}},
			{Delay: delay},
			{Discharge: &Discharge{}},
		},
	}
	p := newDefaultParser(ctx, t, time.Now())

	cases := []struct {
		name              string
		pathwayDefinition []byte
		wantName          string
		want              Pathway
		wantErr           bool
	}{{
		name: "valid yml pathway without name",
		pathwayDefinition: []byte(`
pathway:
  - admission:
      loc: Renal
  - delay:
      from: 1s
      to: 5s
  - discharge: {}
`),
		want:     pathway,
		wantName: UnknownPathwayName,
	}, {
		name: "valid json pathway without name",
		pathwayDefinition: []byte(`
{
  "pathway": [
    {
      "admission": {
        "loc": "Renal"
      }
    },
    {
      "delay": {
        "from": "1s",
        "to": "5s"
      }
    },
    {
      "discharge": {}
    }
  ]
}
`),
		want:     pathway,
		wantName: UnknownPathwayName,
	}, {
		name: "valid json pathway without name, no brackets",
		pathwayDefinition: []byte(`
"pathway": [
	{
		"admission": {
			"loc": "Renal"
		}
	},
	{
		"delay": {
			"from": "1s",
			"to": "5s"
		}
	},
	{
		"discharge": {}
	}
]
`),
		want:     pathway,
		wantName: UnknownPathwayName,
	}, {
		name: "valid yml pathway with name",
		pathwayDefinition: []byte(`
random_pathway:
  pathway:
    - admission:
        loc: Renal
    - delay:
        from: 1s
        to: 5s
    - discharge: {}
`),
		wantName: "random_pathway",
		want:     pathway,
	}, {
		name: "valid json pathway with name",
		pathwayDefinition: []byte(`
{
  "random_pathway": {
    "pathway": [
      {
        "admission": {
          "loc": "Renal"
        }
      },
      {
        "delay": {
          "from": "1s",
          "to": "5s"
        }
      },
      {
        "discharge": {}
      }
    ]
  }
}`),
		wantName: "random_pathway",
		want:     pathway,
	}, {
		name: "valid json pathway with name, no brackets",
		pathwayDefinition: []byte(`
"random_pathway": {
	"pathway": [
		{
			"admission": {
				"loc": "Renal"
			}
		},
		{
			"delay": {
				"from": "1s",
				"to": "5s"
			}
		},
		{
			"discharge": {}
		}
	]
}`),
		wantName: "random_pathway",
		want:     pathway,
	}, {
		name:              "invalid yml format",
		pathwayDefinition: []byte("pathway with invalid yml format"),
		want:              Pathway{},
		wantErr:           true,
	}, {
		name:              "invalid pathway: both: history and pathway empty",
		pathwayDefinition: []byte("pathway:\n"),
		want:              Pathway{},
		wantErr:           true,
	}, {
		name: "invalid yml pathway: too many pathways",
		pathwayDefinition: []byte(`
random_pathway:
  pathway:
    - admission:
        loc: Renal
    - delay:
        from: 1s
        to: 5s
    - discharge: {}
another_random_pathway:
  pathway:
    - admission:
        loc: Renal
`),
		want:    Pathway{},
		wantErr: true,
	}, {
		name: "invalid json pathway: too many pathways",
		pathwayDefinition: []byte(`
{
  "random_pathway": {
    "pathway": [
      {
        "admission": {
          "loc": "Renal"
        }
      },
      {
        "delay": {
          "from": "1s",
          "to": "5s"
        }
      },
      {
        "discharge": {}
      }
    ]
  },
  "another_random_pathway": {
    "pathway": [
      {
        "admission": {
          "loc": "Renal"
        }
      }
    ]
  }
}
`),
		want:    Pathway{},
		wantErr: true,
	}, {
		name: "invalid yml pathway: unknown field",
		pathwayDefinition: []byte(`
random_pathway:
  unknown_field: arbitrary_string
  pathway:
    - admission:
        loc: Renal
    - discharge: {}
`),
		want:    Pathway{},
		wantErr: true,
	}, {
		name: "invalid json pathway: unknown field",
		pathwayDefinition: []byte(`
{
  "random_pathway": {
    "unknown_field": "arbitrary_string",
    "pathway": [
      {
        "admission": {
          "loc": "Renal"
        }
      },
      {
        "discharge": {}
      }
    ]
  }
}
`),
		want:    Pathway{},
		wantErr: true,
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := p.ParseSinglePathway(tc.pathwayDefinition)

			gotErr := err != nil
			if gotErr != tc.wantErr {
				t.Fatalf("ParseSinglePathway(%s) got err %v, want err? %t", string(tc.pathwayDefinition), err, tc.wantErr)
			}
			if diff := cmp.Diff(tc.want, got, cmpopts.IgnoreUnexported(Pathway{}, Step{})); diff != "" {
				t.Errorf("ParseSinglePathway(%s) got diff (-want, +got):\n%s", string(tc.pathwayDefinition), diff)
			}
			if gotErr || tc.wantErr {
				return
			}

			if gotName := got.Name(); gotName != tc.wantName {
				t.Errorf("ParseSinglePathway(%v).Name()=%q, want %q.", string(tc.pathwayDefinition), gotName, tc.wantName)
			}
		})
	}
}

func TestParseSinglePathwayWithPatients(t *testing.T) {
	ctx := context.Background()
	p := newDefaultParser(ctx, t, time.Now())
	cases := []struct {
		name              string
		pathwayDefinition []byte
		want              Pathway
	}{
		{
			name: "use_patient use MRN",
			pathwayDefinition: []byte(`
pathway:
  - admission:
      loc: Renal
  - use_patient:
      patient: 1234
  - discharge: {}
`),
			want: Pathway{
				Persons: &Persons{defaultPatientID: {}},
				Pathway: []Step{
					Step{Admission: &Admission{Loc: "Renal"}},
					Step{UsePatient: &UsePatient{Patient: "1234"}},
					Step{Discharge: &Discharge{}},
				},
			},
		}, {
			name: "use_patient referes to persons section",
			pathwayDefinition: []byte(`
persons:
  first_patient: {}
  second_patient: {}
pathway:
  - use_patient:
      patient: first_patient
  - admission:
      loc: Renal
  - use_patient:
      patient: second_patient
  - discharge: {}
`),
			want: Pathway{
				Persons: &Persons{"first_patient": {}, "second_patient": {}},
				Pathway: []Step{
					Step{UsePatient: &UsePatient{Patient: "first_patient"}},
					Step{Admission: &Admission{Loc: "Renal"}},
					Step{UsePatient: &UsePatient{Patient: "second_patient"}},
					Step{Discharge: &Discharge{}},
				},
			},
		}, {
			name: "use_patient refers to persons section and MRN",
			pathwayDefinition: []byte(`
persons:
  first_patient: {}
  second_patient: {}
pathway:
  - use_patient:
      patient: first_patient
  - admission:
      loc: Renal
  - use_patient:
      patient: second_patient
  - use_patient:
      patient: 1234
  - discharge: {}
`),
			want: Pathway{
				Persons: &Persons{"first_patient": {}, "second_patient": {}},
				Pathway: []Step{
					Step{UsePatient: &UsePatient{Patient: "first_patient"}},
					Step{Admission: &Admission{Loc: "Renal"}},
					Step{UsePatient: &UsePatient{Patient: "second_patient"}},
					Step{UsePatient: &UsePatient{Patient: "1234"}},
					Step{Discharge: &Discharge{}},
				},
			},
		}, {
			name: "bed_swap use MRN",
			pathwayDefinition: []byte(`
pathway:
  - admission:
      loc: Renal
  - bed_swap:
      patient_1: 1234
      patient_2: 5678
  - discharge: {}
`),
			want: Pathway{
				Persons: &Persons{defaultPatientID: {}},
				Pathway: []Step{
					Step{Admission: &Admission{Loc: "Renal"}},
					Step{BedSwap: &BedSwap{Patient1: "1234", Patient2: "5678"}},
					Step{Discharge: &Discharge{}},
				},
			},
		}, {
			name: "bed_swap use CURRENT and MRN",
			pathwayDefinition: []byte(`
pathway:
  - admission:
      loc: Renal
  - bed_swap:
      patient_1: CURRENT
      patient_2: 5678
  - discharge: {}
`),
			want: Pathway{
				Persons: &Persons{defaultPatientID: {}},
				Pathway: []Step{
					Step{Admission: &Admission{Loc: "Renal"}},
					Step{BedSwap: &BedSwap{Patient1: "CURRENT", Patient2: "5678"}},
					Step{Discharge: &Discharge{}},
				},
			},
		}, {
			name: "bed_swap refers to persons",
			pathwayDefinition: []byte(`
persons:
  patient_1: {}
  patient_2: {}
pathway:
  - use_patient:
      patient: patient_1
  - admission:
      loc: Renal
  - use_patient:
      patient: patient_2
  - admission:
      loc: Renal
  - bed_swap:
      patient_1: patient_1
      patient_2: patient_2
  - discharge: {}
`),
			want: Pathway{
				Persons: &Persons{"patient_1": {}, "patient_2": {}},
				Pathway: []Step{
					Step{UsePatient: &UsePatient{Patient: "patient_1"}},
					Step{Admission: &Admission{Loc: "Renal"}},
					Step{UsePatient: &UsePatient{Patient: "patient_2"}},
					Step{Admission: &Admission{Loc: "Renal"}},
					Step{BedSwap: &BedSwap{Patient1: "patient_1", Patient2: "patient_2"}},
					Step{Discharge: &Discharge{}},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := p.ParseSinglePathway(tc.pathwayDefinition)
			if err != nil {
				t.Fatalf("ParseSinglePathway(%s) failed with %v", string(tc.pathwayDefinition), err)
			}
			if diff := cmp.Diff(tc.want, got, cmpopts.IgnoreUnexported(Pathway{}, Step{})); diff != "" {
				t.Errorf("ParseSinglePathway(%s) got diff (-want, +got):\n%s", string(tc.pathwayDefinition), diff)
			}
		})
	}
}

func TestParseSinglePathwayDateTimeAllergies(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name              string
		pathwayDefinition []byte
		wantAllergy       Allergy
	}{
		{
			name: "allergy IdentificationDateTime Time",
			pathwayDefinition: []byte(`
            pathway:
              - admission:
                  loc: Renal
                  allergies:
                  - code: J301
                    identification_datetime:
                        time: 2006-07-18
            `),
			wantAllergy: Allergy{
				Code: "J301",
				IdentificationDateTime: &DateTime{
					Time: &july182006,
				},
			},
		},
		{
			name: "allergy IdentificationDateTime TimeFromNow",
			pathwayDefinition: []byte(`
            pathway:
              - admission:
                  loc: Renal
                  allergies:
                  - code: J301
                    identification_datetime:
                        time_from_now: -10s
            `),
			wantAllergy: Allergy{
				Code: "J301",
				IdentificationDateTime: &DateTime{
					TimeFromNow: &tenSecondsAgo,
				},
			},
		},
		{
			name: "allergy IdentificationDateTime NoDateTimeRecorded",
			pathwayDefinition: []byte(`
            pathway:
              - admission:
                  loc: Renal
                  allergies:
                  - code: J301
                    identification_datetime:
                        no_datetime_recorded: true
            `),
			wantAllergy: Allergy{
				Code: "J301",
				IdentificationDateTime: &DateTime{
					NoDateTimeRecorded: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := newDefaultParser(ctx, t, time.Now())
			got, err := p.ParseSinglePathway(tt.pathwayDefinition)
			if err != nil {
				t.Fatalf("ParseSinglePathway(%s) failed with %v", string(tt.pathwayDefinition), err)
			}
			if diff := cmp.Diff([]Allergy{tt.wantAllergy}, got.Pathway[0].Admission.Allergies); diff != "" {
				t.Errorf("ParseSinglePathway(%s).Pathway[0].Admission.Allergies got diff (-want +got):\n%s", string(tt.pathwayDefinition), diff)
			}
		})
	}
}

func TestParseSinglePathwayDateTimeDiagnoses(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name              string
		pathwayDefinition []byte
		wantDiagnosis     *DiagnosisOrProcedure
	}{
		{
			name: "diagnosis DateTime Time",
			pathwayDefinition: []byte(`
            pathway:
              - admission:
                  loc: Renal
              - update_person:
                  diagnoses:
                  - type: Working
                    code: A101.1
                    description: diagnosis1
                    datetime:
                        time: 2006-07-18
            `),
			wantDiagnosis: &DiagnosisOrProcedure{
				Type:        "Working",
				Code:        "A101.1",
				Description: "diagnosis1",
				DateTime: &DateTime{
					Time: &july182006,
				},
			},
		},
		{
			name: "diagnosis DateTime TimeFromNow",
			pathwayDefinition: []byte(`
            pathway:
              - admission:
                  loc: Renal
              - update_person:
                  diagnoses:
                  - type: Working
                    code: A101.1
                    description: diagnosis1
                    datetime:
                        time_from_now: -10s
            `),
			wantDiagnosis: &DiagnosisOrProcedure{
				Type:        "Working",
				Code:        "A101.1",
				Description: "diagnosis1",
				DateTime: &DateTime{
					TimeFromNow: &tenSecondsAgo,
				},
			},
		},
		{
			name: "diagnosis DateTime NoDateTimeRecorded",
			pathwayDefinition: []byte(`
            pathway:
              - admission:
                  loc: Renal
              - update_person:
                  diagnoses:
                  - type: Working
                    code: A101.1
                    description: diagnosis1
                    datetime:
                        no_datetime_recorded: true
                    `),
			wantDiagnosis: &DiagnosisOrProcedure{
				Type:        "Working",
				Code:        "A101.1",
				Description: "diagnosis1",
				DateTime: &DateTime{
					NoDateTimeRecorded: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := newDefaultParser(ctx, t, time.Now())
			got, err := p.ParseSinglePathway(tt.pathwayDefinition)
			if err != nil {
				t.Fatalf("ParseSinglePathway(%s) failed with %v", string(tt.pathwayDefinition), err)
			}
			if diff := cmp.Diff([]*DiagnosisOrProcedure{tt.wantDiagnosis}, got.Pathway[1].UpdatePerson.Diagnoses); diff != "" {
				t.Errorf("ParseSinglePathway(%s).Pathway[1].UpdatePerson.Diagnoses got diff (-want, +got):\n%s", string(tt.pathwayDefinition), diff)
			}
		})
	}
}

func TestParseSinglePathwayDateTimeProcedures(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name              string
		pathwayDefinition []byte
		wantProcedure     *DiagnosisOrProcedure
	}{
		{
			name: "procedure DateTime Time",
			pathwayDefinition: []byte(`
            pathway:
              - admission:
                  loc: Renal
              - update_person:
                  procedures:
                  - type: A
                    code: A102.2
                    description: procedure1
                    datetime:
                        time: 2006-07-18
            `),
			wantProcedure: &DiagnosisOrProcedure{
				Type:        "A",
				Code:        "A102.2",
				Description: "procedure1",
				DateTime: &DateTime{
					Time: &july182006,
				},
			},
		},
		{
			name: "procedure DateTime TimeFromNow",
			pathwayDefinition: []byte(`
            pathway:
              - admission:
                  loc: Renal
              - update_person:
                  procedures:
                  - type: A
                    code: A102.2
                    description: procedure1
                    datetime:
                        time_from_now: -10s
                    `),
			wantProcedure: &DiagnosisOrProcedure{
				Type:        "A",
				Code:        "A102.2",
				Description: "procedure1",
				DateTime: &DateTime{
					TimeFromNow: &tenSecondsAgo,
				},
			},
		},
		{
			name: "procedure DateTime NoDateTimeRecorded",
			pathwayDefinition: []byte(`
            pathway:
              - admission:
                  loc: Renal
              - update_person:
                  procedures:
                  - type: A
                    code: A102.2
                    description: procedure1
                    datetime:
                        no_datetime_recorded: true
            `),
			wantProcedure: &DiagnosisOrProcedure{
				Type:        "A",
				Code:        "A102.2",
				Description: "procedure1",
				DateTime: &DateTime{
					NoDateTimeRecorded: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := newDefaultParser(ctx, t, time.Now())
			got, err := p.ParseSinglePathway(tt.pathwayDefinition)
			if err != nil {
				t.Fatalf("ParseSinglePathway(%s) failed with %v", string(tt.pathwayDefinition), err)
			}
			if diff := cmp.Diff([]*DiagnosisOrProcedure{tt.wantProcedure}, got.Pathway[1].UpdatePerson.Procedures); diff != "" {
				t.Errorf("ParseSinglePathway(%s).Pathway[1].UpdatePerson.Procedures got diff (-want, +got):\n%s", string(tt.pathwayDefinition), diff)
			}
		})
	}
}

func TestParseSinglePathwayInvalidDateTime(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name              string
		pathwayDefinition []byte
	}{
		{
			name: "allergy IdentificationDateTime Time & TimeFromNow",
			pathwayDefinition: []byte(`
            pathway:
              - admission:
                  loc: Renal
                  allergies:
                  - code: J301
                    identification_datetime:
                        time: 2006-07-18
                        time_from_now: -10s
            `),
		},
		{
			name: "allergy IdentificationDateTime Time & NoRecordedDateTime",
			pathwayDefinition: []byte(`
            pathway:
              - admission:
                  loc: Renal
                  allergies:
                  - code: J301
                    identification_datetime:
                        time_from_now: -10s
                        no_datetime_recorded: true
            `),
		},
		{
			name: "allergy IdentificationDateTime TimeFromNow & NoDateTimeRecorded",
			pathwayDefinition: []byte(`
            pathway:
              - admission:
                  loc: Renal
                  allergies:
                  - code: J301
                    identification_datetime:
                        time: 2006-07-18
                        no_datetime_recorded: true
            `),
		},
		{
			name: "diagnosis DateTime Time & TimeFromNow",
			pathwayDefinition: []byte(`
            pathway:
              - admission:
                  loc: Renal
              - update_person:
                  diagnoses:
                  - type: Working
                    code: A101.1
                    description: diagnosis1
                    datetime:
                        time: 2006-07-18
                        time_from_now: -10s
            `),
		},
		{
			name: "diagnosis DateTime Time & NoRecordedDateTime",
			pathwayDefinition: []byte(`
            pathway:
              - admission:
                  loc: Renal
              - update_person:
                  diagnoses:
                  - type: Working
                    code: A101.1
                    description: diagnosis1
                    datetime:
                        time: 2006-07-18
                        no_datetime_recorded: true
            `),
		},
		{
			name: "procedure DateTime TimeFromNow & NoDateTimeRecorded",
			pathwayDefinition: []byte(`
            pathway:
              - admission:
                  loc: Renal
              - update_person:
                  procedures:
                  - type: A
                    code: A102.2
                    description: procedure1
                    datetime:
                        time_from_now: -10s
                        no_datetime_recorded: true
            `),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := newDefaultParser(ctx, t, time.Now())
			if _, err := p.ParseSinglePathway(tt.pathwayDefinition); err == nil {
				t.Fatalf("ParseSinglePathway(%s) got err=<nil>, want non-nil err", string(tt.pathwayDefinition))
			}
		})
	}
}

func TestParseAutogenerate(t *testing.T) {
	ctx := context.Background()
	fourHoursAgo := -4 * time.Hour
	threeHoursAgo := -3 * time.Hour
	twoHoursAgo := -2 * time.Hour
	oneHourAgo := -1 * time.Hour

	customParams := &Parameters{
		SendingApplication:   "sa",
		SendingFacility:      "sf",
		ReceivingApplication: "ra",
		ReceivingFacility:    "rf",
		Custom:               map[string]string{"key": "value"},
	}

	cases := []struct {
		name              string
		pathwayDefinition []byte
		want              Pathway
	}{
		{
			name: "negative time no previous history",
			pathwayDefinition: []byte(`
random_pathway:
  pathway:
    - admission:
        loc: Renal
    - autogenerate:
        from: -10s
        to: -10s
    - discharge: {}
`),
			want: Pathway{
				Persons: &Persons{defaultPatientID: {}},
				History: []Step{
					Step{Result: &Results{OrderProfile: "RANDOM"}, Parameters: &Parameters{TimeFromNow: &tenSecondsAgo}},
				},
				Pathway: []Step{
					Step{Admission: &Admission{Loc: "Renal"}},
					Step{Discharge: &Discharge{}},
				},
			},
		}, {
			name: "negative time before first step in the history",
			pathwayDefinition: []byte(`
random_pathway:
  historical_data:
    - admission:
        loc: Renal
      parameters:
        time_from_now: -3h
    - discharge: {}
      parameters:
        time_from_now: -1h
  pathway:
    - autogenerate:
        from: -4h
        to: -4h
`),
			want: Pathway{
				Persons: &Persons{defaultPatientID: {}},
				History: []Step{
					Step{Result: &Results{OrderProfile: "RANDOM"}, Parameters: &Parameters{TimeFromNow: &fourHoursAgo}},
					Step{Admission: &Admission{Loc: "Renal"}, Parameters: &Parameters{TimeFromNow: &threeHoursAgo}},
					Step{Discharge: &Discharge{}, Parameters: &Parameters{TimeFromNow: &oneHourAgo}},
				},
				Pathway: []Step{},
			},
		}, {
			name: "negative time same as the first step in the history - insert befre first step",
			pathwayDefinition: []byte(`
random_pathway:
  historical_data:
    - admission:
        loc: Renal
      parameters:
        time_from_now: -3h
    - discharge: {}
      parameters:
        time_from_now: -1h
  pathway:
    - autogenerate:
        from: -3h
        to: -3h
`),
			want: Pathway{
				Persons: &Persons{defaultPatientID: {}},
				History: []Step{
					Step{Result: &Results{OrderProfile: "RANDOM"}, Parameters: &Parameters{TimeFromNow: &threeHoursAgo}},
					Step{Admission: &Admission{Loc: "Renal"}, Parameters: &Parameters{TimeFromNow: &threeHoursAgo}},
					Step{Discharge: &Discharge{}, Parameters: &Parameters{TimeFromNow: &oneHourAgo}},
				},
				Pathway: []Step{},
			},
		}, {
			name: "negative time between two steps in the history",
			pathwayDefinition: []byte(`
random_pathway:
  historical_data:
    - admission:
        loc: Renal
      parameters:
        time_from_now: -3h
    - discharge: {}
      parameters:
        time_from_now: -1h
  pathway:
    - autogenerate:
        from: -2h
        to: -2h
`),
			want: Pathway{
				Persons: &Persons{defaultPatientID: {}},
				History: []Step{
					Step{Admission: &Admission{Loc: "Renal"}, Parameters: &Parameters{TimeFromNow: &threeHoursAgo}},
					Step{Result: &Results{OrderProfile: "RANDOM"}, Parameters: &Parameters{TimeFromNow: &twoHoursAgo}},
					Step{Discharge: &Discharge{}, Parameters: &Parameters{TimeFromNow: &oneHourAgo}},
				},
				Pathway: []Step{},
			},
		}, {
			name: "negative time last step in the history",
			pathwayDefinition: []byte(`
random_pathway:
  historical_data:
    - admission:
        loc: Renal
      parameters:
        time_from_now: -3h
    - discharge: {}
      parameters:
        time_from_now: -1h
  pathway:
    - autogenerate:
        from: -10s
        to: -10s
`),
			want: Pathway{
				Persons: &Persons{defaultPatientID: {}},
				History: []Step{
					Step{Admission: &Admission{Loc: "Renal"}, Parameters: &Parameters{TimeFromNow: &threeHoursAgo}},
					Step{Discharge: &Discharge{}, Parameters: &Parameters{TimeFromNow: &oneHourAgo}},
					Step{Result: &Results{OrderProfile: "RANDOM"}, Parameters: &Parameters{TimeFromNow: &tenSecondsAgo}},
				},
				Pathway: []Step{},
			},
		}, {
			name: "positive time break delay",
			pathwayDefinition: []byte(`
random_pathway:
  pathway:
    - admission:
        loc: Renal
    - delay:
        from: 5s
        to: 5s
    - autogenerate:
        from: 2s
        to: 2s
    - discharge: {}
`),
			want: Pathway{
				Persons: &Persons{defaultPatientID: {}},
				Pathway: []Step{
					Step{Admission: &Admission{Loc: "Renal"}},
					Step{Delay: &Delay{From: 2 * time.Second, To: 2 * time.Second}},
					Step{Result: &Results{OrderProfile: "RANDOM"}},
					Step{Delay: &Delay{From: 3 * time.Second, To: 3 * time.Second}},
					Step{Discharge: &Discharge{}},
				},
			},
		}, {
			name: "positive time add delay",
			pathwayDefinition: []byte(`
random_pathway:
  pathway:
    - admission:
        loc: Renal
    - autogenerate:
        from: 2s
        to: 2s
    - discharge: {}
`),
			want: Pathway{
				Persons: &Persons{defaultPatientID: {}},
				Pathway: []Step{
					Step{Admission: &Admission{Loc: "Renal"}},
					Step{Discharge: &Discharge{}},
					Step{Delay: &Delay{From: 2 * time.Second, To: 2 * time.Second}},
					Step{Result: &Results{OrderProfile: "RANDOM"}},
				},
			},
		}, {
			name: "positive time at the end of an existing delay",
			pathwayDefinition: []byte(`
random_pathway:
  pathway:
    - admission:
        loc: Renal
    - delay:
        from: 1s
        to: 1s
    - autogenerate:
        from: 1s
        to: 1s
    - discharge: {}
`),
			want: Pathway{
				Persons: &Persons{defaultPatientID: {}},
				Pathway: []Step{
					Step{Admission: &Admission{Loc: "Renal"}},
					Step{Delay: &Delay{From: 1 * time.Second, To: 1 * time.Second}},
					Step{Result: &Results{OrderProfile: "RANDOM"}},
					Step{Discharge: &Discharge{}},
				},
			},
		}, {
			name: "multiple results negative and positive",
			pathwayDefinition: []byte(`
random_pathway:
  pathway:
    - autogenerate:
        result:
          order_profile: Vital Signs
        from: -3h
        to: 1h
        every: 1h
    - admission:
        loc: Renal
    - delay:
        from: 2h
        to: 2h
    - discharge: {}
    - autogenerate:
        result:
          order_profile: Creatinine
        from: 0h
        to: 0h
`),
			want: Pathway{
				Persons: &Persons{defaultPatientID: {}},
				History: []Step{
					Step{Result: &Results{OrderProfile: "Vital Signs"}, Parameters: &Parameters{TimeFromNow: &threeHoursAgo}},
					Step{Result: &Results{OrderProfile: "Vital Signs"}, Parameters: &Parameters{TimeFromNow: &twoHoursAgo}},
					Step{Result: &Results{OrderProfile: "Vital Signs"}, Parameters: &Parameters{TimeFromNow: &oneHourAgo}},
				},
				Pathway: []Step{
					Step{Admission: &Admission{Loc: "Renal"}},
					Step{Result: &Results{OrderProfile: "Vital Signs"}},
					Step{Result: &Results{OrderProfile: "Creatinine"}},
					Step{Delay: &Delay{From: 1 * time.Hour, To: 1 * time.Hour}},
					Step{Result: &Results{OrderProfile: "Vital Signs"}},
					Step{Delay: &Delay{From: 1 * time.Hour, To: 1 * time.Hour}},
					Step{Discharge: &Discharge{}},
				},
			},
		}, {
			name: "with parameters",
			pathwayDefinition: []byte(`
random_pathway:
  pathway:
    - autogenerate:
        result:
          order_profile: Vital Signs
        from: -1h
        to: 1h
        every: 1h
      parameters:
        sending_application: sa
        sending_facility: sf
        receiving_application: ra
        receiving_facility: rf
        custom:
          key: value
`),
			want: Pathway{
				Persons: &Persons{defaultPatientID: {}},
				History: []Step{
					Step{
						Result: &Results{OrderProfile: "Vital Signs"},
						Parameters: &Parameters{
							TimeFromNow:          &oneHourAgo,
							SendingApplication:   "sa",
							SendingFacility:      "sf",
							ReceivingApplication: "ra",
							ReceivingFacility:    "rf",
							Custom:               map[string]string{"key": "value"},
						},
					}},
				Pathway: []Step{
					Step{Result: &Results{OrderProfile: "Vital Signs"}, Parameters: customParams},
					Step{Delay: &Delay{From: time.Hour, To: time.Hour}},
					Step{Result: &Results{OrderProfile: "Vital Signs"}, Parameters: customParams},
				},
			},
		}, {
			name: "with use_patient",
			pathwayDefinition: []byte(`
random_pathway:
  historical_data:
    - use_patient:
        patient: 1234
    - result:
        order_profile: Creatinine
      parameters:
        time_from_now: -1h
  pathway:
    - autogenerate:
        result:
          order_profile: Vital Signs
        from: -1h
        to: 0h
        every: 1h
    - use_patient:
        patient: 1234
    - admission:
        loc: Renal
    - discharge: {}
`),
			want: Pathway{
				Persons: &Persons{defaultPatientID: {}},
				History: []Step{
					Step{UsePatient: &UsePatient{Patient: "1234"}},
					Step{Result: &Results{OrderProfile: "Vital Signs"}, Parameters: &Parameters{TimeFromNow: &oneHourAgo}},
					Step{Result: &Results{OrderProfile: "Creatinine"}, Parameters: &Parameters{TimeFromNow: &oneHourAgo}},
				},
				Pathway: []Step{
					Step{UsePatient: &UsePatient{Patient: "1234"}},
					Step{Admission: &Admission{Loc: "Renal"}},
					Step{Discharge: &Discharge{}},
					Step{Result: &Results{OrderProfile: "Vital Signs"}},
				},
			},
		},
	}

	for _, tc := range cases {
		for parseFunc, got := range parsePathwayDefinition(ctx, t, tc.pathwayDefinition) {
			t.Run(fmt.Sprintf("%s-%s", tc.name, parseFunc.name), func(t *testing.T) {
				if diff := cmp.Diff(tc.want, got, cmpopts.IgnoreUnexported(Pathway{}, Step{})); diff != "" {
					t.Errorf("%s got diff (-want, +got):\n%s", parseFunc.errMsg, diff)
				}
			})
		}
	}
}

func TestParseAutoGenerateGetPathwayDifferentEveryTime(t *testing.T) {
	ctx := context.Background()
	rand.Seed(1)

	pathwayDefinition := []byte(`
pathway_autogenerate:
  pathway:
    - admission:
        loc: Renal
    - delay:
        from: 0s
        to: 2s
    - discharge: {}
    - autogenerate:
        result:
          order_profile: Vital Signs
        from: -1s
        to: 1s
        every: 1s
`)

	mainDir := writePathwayToDir(t, pathwayDefinition)

	p := newDefaultParser(ctx, t, time.Now())
	pathways, err := p.ParsePathways(ctx, mainDir)
	if err != nil {
		t.Fatalf("ParsePathways(%s failed with %v", string(pathwayDefinition), err)
	}

	pathwayName := "pathway_autogenerate"
	pathway, ok := pathways[pathwayName]
	if !ok {
		t.Fatalf("ParsePathways(%s)=%+v, want %q to exist", string(pathwayDefinition), pathways, pathwayName)
	}

	oneSecondAgo := -1 * time.Second
	delay0To1 := "delay0To1"
	delay1 := "delay1"
	delay1To2 := "delay1To2"

	// Every time Runnable() is called on the pathway_autogenerate pathway,
	// a slightly different pathway is generated.
	// This is because the delay defined in the pathway can be between 0s-2s.
	// When inserting autogenerate steps, the actual value of the delay needs to
	// be known, so in practice all delays in such pathway will be a concrete
	// values, ie Delay.From == Delay.To.
	// Depending on the actual value of that delay step, there can be 3 different
	// pathways generated:
	wantPathways := map[string]Pathway{
		// 0s < delay < 1s
		// In this case, Discharge step follows the delay defined in the pathway.
		// Vital Signs result is inserted at the end of the pathway after an additional
		// delay, so that the sum of both delays is equal 1s.
		delay0To1: Pathway{
			Persons: &Persons{defaultPatientID: {}},
			History: []Step{
				Step{Result: &Results{OrderProfile: "Vital Signs"}, Parameters: &Parameters{TimeFromNow: &oneSecondAgo}},
			},
			Pathway: []Step{
				Step{Admission: &Admission{Loc: "Renal"}},
				Step{Result: &Results{OrderProfile: "Vital Signs"}}, // 0s
				Step{Delay: &Delay{From: 0 * time.Second, To: 1 * time.Second}},
				Step{Discharge: &Discharge{}},
				Step{Delay: &Delay{From: 0 * time.Hour, To: 1 * time.Hour}},
				Step{Result: &Results{OrderProfile: "Vital Signs"}}, // 1s
			},
		},
		// delay == 1s
		// In this case, Vital Signs result is inserted after 1s delay,
		// immediatelly followed by the Discharge.
		// Note that this pathway is very unlikely.
		delay1: Pathway{
			Persons: &Persons{defaultPatientID: {}},
			History: []Step{
				Step{Result: &Results{OrderProfile: "Vital Signs"}, Parameters: &Parameters{TimeFromNow: &oneSecondAgo}},
			},
			Pathway: []Step{
				Step{Admission: &Admission{Loc: "Renal"}},
				Step{Result: &Results{OrderProfile: "Vital Signs"}}, // 0s
				Step{Delay: &Delay{From: 1 * time.Second, To: 1 * time.Second}},
				Step{Result: &Results{OrderProfile: "Vital Signs"}}, // 1s
				Step{Discharge: &Discharge{}},
			},
		},
		// 1s < delay < 2s
		// In this case, the original delay is split into two:
		// the first delay is exactly 1s, after which the Vital Signs result is inserted.
		// It is followed by the remaining delay, followed by the Discharge step.
		delay1To2: Pathway{
			Persons: &Persons{defaultPatientID: {}},
			History: []Step{
				Step{Result: &Results{OrderProfile: "Vital Signs"}, Parameters: &Parameters{TimeFromNow: &oneSecondAgo}},
			},
			Pathway: []Step{
				Step{Admission: &Admission{Loc: "Renal"}},
				Step{Result: &Results{OrderProfile: "Vital Signs"}}, // 0s
				Step{Delay: &Delay{From: 1 * time.Second, To: 1 * time.Second}},
				Step{Result: &Results{OrderProfile: "Vital Signs"}}, // 1s
				Step{Delay: &Delay{From: 0 * time.Hour, To: 1 * time.Hour}},
				Step{Discharge: &Discharge{}},
			},
		},
	}

	runs := 1000
	gotFreq := make(map[string]int)
	for i := 0; i < runs; i++ {
		got, err := pathway.Runnable()
		if err != nil {
			t.Fatalf("[%v].Runnable() failed with %v", pathway, err)
		}

		matchFound := false
		for k, want := range wantPathways {
			if diff := cmp.Diff(want, got, cmpopts.IgnoreUnexported(Pathway{}, Step{}), cmpopts.IgnoreFields(Delay{}, "From", "To")); diff == "" {
				gotFreq[k]++
				matchFound = true

				switch k {
				case delay0To1:
					firstDelay := got.Pathway[2].Delay
					if firstDelay.From != firstDelay.To {
						t.Errorf("%s pathway: first delay want From == To, got %+v", delay0To1, firstDelay)
					}
					secondDelay := got.Pathway[4].Delay
					if secondDelay.From != secondDelay.To {
						t.Errorf("%s pathway: second delay want From == To, got %+v", delay0To1, secondDelay)
					}
					if got, want := firstDelay.From+secondDelay.From, time.Second; got != want {
						t.Errorf("%s pathway: first delay (%+v) + second delay (%+v) = %v, want %v", delay0To1, firstDelay, secondDelay, got, want)
					}

				case delay1:
					delay := got.Pathway[2].Delay
					if delay.From != delay.To {
						t.Errorf("%s pathway: delay want From == To, got %+v", delay1, delay)
					}

				case delay1To2:
					firstDelay := got.Pathway[2].Delay
					if want := time.Second; firstDelay.From != firstDelay.To || firstDelay.From != want {
						t.Errorf("%s pathway: first delay want From == To == %v, got %+v", delay1To2, want, firstDelay)
					}
					secondDelay := got.Pathway[4].Delay
					if secondDelay.From != secondDelay.To {
						t.Errorf("%s pathway: second delay want From == To, got %+v", delay1To2, secondDelay)
					}
					if want := time.Second; secondDelay.From > want {
						t.Errorf("%s pathway: second delay = %v, want less or equal than %v", delay1To2, secondDelay, want)
					}
				}
				break
			}
		}

		if !matchFound {
			t.Errorf("[%v].Runnable()=%+v, want one of %+v", pathway, got, wantPathways)
		}
	}

	// Pathways delay0To1 and delay1To2 are equally probable.
	// Pathway delay1 is very unlikely, so we don't count occurrences of it.
	delta := float64(runs) / 20
	for _, pathwayName := range []string{delay0To1, delay1To2} {
		if got, want := gotFreq[pathwayName], runs/2; math.Abs(float64(got)-float64(want)) >= delta {
			t.Errorf("gotFreq[%q] = %d, want within %v of %v", pathwayName, got, delta, want)
		}
	}
}

func TestRunnableOriginalNotModified(t *testing.T) {
	ctx := context.Background()
	twoSeconds := 2 * time.Second
	pathwayDefinition := []byte(`
random_pathway:
  pathway:
    - admission:
        loc: Renal
    - autogenerate:
        from: 2s
        to: 2s
    - discharge: {}
`)
	mainDir := writePathwayToDir(t, pathwayDefinition)

	p := newDefaultParser(ctx, t, time.Now())
	pathways, err := p.ParsePathways(ctx, mainDir)
	if err != nil {
		t.Fatalf("ParsePathways(%s) failed with %v", string(pathwayDefinition), err)
	}

	wantOriginal := Pathway{
		Persons: &Persons{defaultPatientID: {}},
		Pathway: []Step{
			Step{Admission: &Admission{Loc: "Renal"}},
			Step{AutoGenerate: &AutoGenerate{From: &twoSeconds, To: &twoSeconds, Result: &Results{OrderProfile: "RANDOM"}}},
			Step{Discharge: &Discharge{}},
		},
	}
	want := Pathway{
		Persons: &Persons{defaultPatientID: {}},
		Pathway: []Step{
			Step{Admission: &Admission{Loc: "Renal"}},
			Step{Discharge: &Discharge{}},
			Step{Delay: &Delay{From: 2 * time.Second, To: 2 * time.Second}},
			Step{Result: &Results{OrderProfile: "RANDOM"}},
		},
	}

	gotOriginal, ok := pathways[defaultPathwayName]
	if !ok {
		t.Fatalf("ParsePathways(%s)=%+v, want %q to exist", string(pathwayDefinition), pathways, defaultPathwayName)
	}
	if diff := cmp.Diff(wantOriginal, gotOriginal, cmpopts.IgnoreUnexported(Pathway{}, Step{})); diff != "" {
		t.Errorf("ParsePathways(%s)[%s] got diff (-want, +got):\n%s", string(pathwayDefinition), defaultPathwayName, diff)
	}

	got, err := gotOriginal.Runnable()
	if err != nil {
		t.Fatalf("[%+v].Runnable() failed with %v", gotOriginal, err)
	}
	if diff := cmp.Diff(want, got, cmpopts.IgnoreUnexported(Pathway{}, Step{})); diff != "" {
		t.Errorf("[%+v].Runnable() got diff (-want, +got):\n%s", gotOriginal, diff)
	}

	if diff := cmp.Diff(wantOriginal, gotOriginal, cmpopts.IgnoreUnexported(Pathway{}, Step{})); diff != "" {
		t.Errorf("ParsePathways(%s)[%s] got diff (-want, +got):\n%s", string(pathwayDefinition), defaultPathwayName, diff)
	}
}

func TestParseOverrideResultID(t *testing.T) {
	ctx := context.Background()
	pathwayDefinition := []byte(`
random_pathway:
  pathway:
    - result:
        order_profile: UREA AND ELECTROLYTES
        results:
          - test_name: Creatinine
            value: 126.00
            unit: UMOLL
            abnormal_flag: HIGH
            id: override id
`)

	for _, p := range parsePathwayDefinition(ctx, t, pathwayDefinition) {
		if got, want := len(p.Pathway), 1; got != want {
			t.Fatalf("len(p.Pathway) = %d, want %d", got, want)
		}

		pathway := p.Pathway[0]
		if pathway.Result == nil {
			t.Fatal("pathway.Result is <nil>.")
		}

		if got, want := len(pathway.Result.Results), 1; got != want {
			t.Fatalf("len(pathway.Result.Results) = %d, want %d", got, want)
		}

		if got, want := pathway.Result.Results[0].ID, "override id"; got != want {
			t.Errorf("pathway.Result.Results.Id = %s, want %s", got, want)
		}
	}
}

type parseFunc struct {
	name   string
	errMsg string
}

// parsePathwayDefinition returns the pathways that result from parsing the pathwayDefinition with
// the two methods that we have to parse pathways: ParsePathways and ParseSinglePathway.
// The pathway definition should be for a pathway with name defaultPathwayName.
// Returns a map, where the key encapsulates the method used to parse the pathway and
// an error message that can be used in the error message in the test to give more information
// on how the pathway was parsed. The value in a map is a parsed pathway.
func parsePathwayDefinition(ctx context.Context, t *testing.T, pathwayDefinition []byte) map[parseFunc]Pathway {
	t.Helper()
	mainDir := writePathwayToDir(t, pathwayDefinition)

	p := newDefaultParser(ctx, t, time.Now())
	pathways, err := p.ParsePathways(ctx, mainDir)
	if err != nil {
		t.Fatalf("ParsePathways(%s) failed with %v", string(pathwayDefinition), err)
	}

	pathway := pathways[defaultPathwayName]
	runnable, err := pathway.Runnable()
	if err != nil {
		t.Fatalf("[%v].Runnable() failed with %v", pathway, err)
	}

	fromParseSingle, err := p.ParseSinglePathway(pathwayDefinition)
	if err != nil {
		t.Fatal("err is something, want <nil>.")
	}
	if got, want := fromParseSingle.Name(), defaultPathwayName; got != want {
		t.Fatalf("fromParseSingle.Name() got %v, want %v", got, want)
	}
	if fromParseSingle.metadata == nil {
		t.Fatal("fromParseSingle.metadata is <nil>.")
	}

	return map[parseFunc]Pathway{
		parseFunc{name: "ParsePathways", errMsg: fmt.Sprintf("ParsePathways(%s, nil, nil)[%s].Runnable()", string(pathwayDefinition), defaultPathwayName)}: runnable,
		parseFunc{name: "ParseSinglePathway", errMsg: fmt.Sprintf("ParseSinglePathway(%s)", string(pathwayDefinition))}:                                    fromParseSingle,
	}
}

func writePathwayToDir(t *testing.T, pathwayDefinition []byte) string {
	t.Helper()
	return testwrite.BytesToDir(t, pathwayDefinition, "pathway1.yml")
}
