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

package document

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/pathway"
	"github.com/google/simhospital/pkg/test/testtext"
)

var (
	date        = time.Date(2018, 2, 12, 1, 25, 0, 0, time.UTC)
	HL7Document = config.HL7Document{
		Types: []string{"AR", "CD", "CN", "DI", "DS"},
	}
	textGenerator = &testtext.Generator{Text: []string{"sample-text-1", "sample-text-2"}}
)

func TestDocument(t *testing.T) {
	obsID := "id"
	obsText := "text"
	cs := "cs"
	tests := []struct {
		name  string
		input *pathway.Document
		want  *ir.Document
	}{{
		name: "Fixed values",
		input: &pathway.Document{
			DocumentType:      "DS",
			CompletionStatus:  "IP",
			ObsIdentifierID:   &obsID,
			ObsIdentifierText: &obsText,
			ObsIdentifierCS:   &cs,
		},
		want: &ir.Document{
			ActivityDateTime:         ir.NewValidTime(date),
			EditDateTime:             ir.NewValidTime(date),
			DocumentType:             "DS",
			DocumentCompletionStatus: "IP",
			ObservationIdentifier: &ir.CodedElement{
				ID:           obsID,
				Text:         obsText,
				CodingSystem: cs,
			},
		},
	}, {
		name: "Default values",
		input: &pathway.Document{
			// We fix the document type because, if not specified, it is picked at random and we wouldn't be able
			// to do assertions on the entire field.
			// TestDocument_RandomDocumentType tests unspecified document types.
			DocumentType: "DS",
		},
		want: &ir.Document{
			ActivityDateTime:         ir.NewValidTime(date),
			EditDateTime:             ir.NewValidTime(date),
			DocumentType:             "DS",
			DocumentCompletionStatus: "DO",
			ObservationIdentifier: &ir.CodedElement{
				ID:           "Established Patient 15",
				Text:         "Established Patient 15",
				CodingSystem: "Simulated Hospital",
			},
		},
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := Generator{DocumentConfig: &HL7Document, TextGenerator: textGenerator}
			got := g.Document(date, tc.input)

			// UniqueDocumentNumber is randomly generated with each document generation.
			// ContentLine contains randomly generated sentences with each document generation.
			ignore := cmpopts.IgnoreFields(ir.Document{}, "UniqueDocumentNumber", "ContentLine")
			if diff := cmp.Diff(tc.want, got, ignore); diff != "" {
				t.Errorf("g.Document() got diff (-want, +got):\n%s", diff)
			}

			if got, want := len(got.UniqueDocumentNumber), 13; got != want {
				t.Errorf("len(g.Document().UniqueDocumentNumber) got %v, want %v", got, want)
			}

			gotContent := len(got.ContentLine)
			if gotContent < 10 || gotContent > 50 {
				t.Errorf("len(g.Document().ContentLine) got %v, want in the range [10,50]", gotContent)
			}
		})
	}
}

func TestDocument_UpdateDocumentContent(t *testing.T) {
	tests := []struct {
		name                 string
		input                *pathway.Document
		originalContentLines []string
		wantContentLine      []string
		wantErr              bool
	}{{
		name: "Append",
		input: &pathway.Document{
			UpdateType: "append",
			NumRandomContentLines: &pathway.Interval{
				From: 2,
				To:   2,
			},
		},
		originalContentLines: []string{"content-1"},
		wantContentLine:      []string{"content-1", "sample-text-1", "sample-text-2"},
	}, {
		name: "Overwrite",
		input: &pathway.Document{
			UpdateType: "overwrite",
			NumRandomContentLines: &pathway.Interval{
				From: 2,
				To:   2,
			},
		},
		originalContentLines: []string{"content-1"},
		wantContentLine:      []string{"sample-text-1", "sample-text-2"},
	}, {
		name: "Append with fixed Ending",
		input: &pathway.Document{
			UpdateType: "append",
			NumRandomContentLines: &pathway.Interval{
				From: 2,
				To:   2,
			},
			EndingContentLines: []string{"ending-text"},
		},
		originalContentLines: []string{"content-1"},
		wantContentLine:      []string{"content-1", "sample-text-1", "sample-text-2", "ending-text"},
	}, {
		name: "Append with fixed Header",
		input: &pathway.Document{
			UpdateType: "append",
			NumRandomContentLines: &pathway.Interval{
				From: 2,
				To:   2,
			},
			HeaderContentLines: []string{"header-text"},
		},
		originalContentLines: []string{"content-1"},
		wantContentLine:      []string{"content-1", "header-text", "sample-text-1", "sample-text-2"},
	}, {
		name: "Empty update",
		input: &pathway.Document{
			NumRandomContentLines: &pathway.Interval{
				From: 2,
				To:   2,
			},
			HeaderContentLines: []string{"header-text"},
		},
		wantErr: true,
	}, {
		name: "Invalid update type",
		input: &pathway.Document{
			UpdateType: "appendd",
			NumRandomContentLines: &pathway.Interval{
				From: 2,
				To:   2,
			},
			HeaderContentLines: []string{"header-text"},
		},
		wantErr: true,
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := Generator{DocumentConfig: &HL7Document, TextGenerator: textGenerator}
			d := &ir.Document{
				ActivityDateTime:         ir.NewValidTime(date),
				EditDateTime:             ir.NewValidTime(date),
				DocumentType:             "DS",
				DocumentCompletionStatus: "DO",
				ObservationIdentifier: &ir.CodedElement{
					ID:           "Established Patient 15",
					Text:         "Established Patient 15",
					CodingSystem: "Simulated Hospital",
				},
				ContentLine: tc.originalContentLines,
			}
			err := g.UpdateDocumentContent(d, tc.input)
			if gotErr := err != nil; gotErr != tc.wantErr {
				t.Errorf("g.UpdateDocumentContent(%v, %v) got error %v; want err? %t", d, tc.input, err, tc.wantErr)
			}
			if diff := cmp.Diff(tc.wantContentLine, d.ContentLine); diff != "" {
				t.Errorf("g.UpdateDocument(%v, %v) got diff: \n%s", d, tc.input, diff)
			}
		})
	}
}

func TestDocument_RandomDocumentTypeDoesntModifyInput(t *testing.T) {
	dt := &pathway.Document{}
	g := Generator{DocumentConfig: &HL7Document, TextGenerator: textGenerator}
	gotDocument := g.Document(date, dt)

	if got := gotDocument.DocumentType; got == "" {
		t.Errorf(`g.Document().DocumentType got %v, want != ""`, got)
	}

	if got, want := dt.DocumentType, ""; got != want {
		t.Errorf("&pathway.Document.DocumentType got %v, want %v", got, want)
	}
}

func TestDocument_RandomDocumentType(t *testing.T) {
	rand.Seed(1)
	g := Generator{DocumentConfig: &HL7Document, TextGenerator: textGenerator}
	runs := float64(1000)
	docTypeDistr := map[string]int{}
	for i := 0; i < int(runs); i++ {
		doc := g.Document(date, &pathway.Document{})
		docType := doc.DocumentType
		docTypeDistr[docType]++
	}

	// Check that each document type is chosen with equal probability.
	wantNum := int(runs / 5)
	wantDocTypeDistr := map[string]int{
		"AR": wantNum, "CD": wantNum, "CN": wantNum, "DI": wantNum, "DS": wantNum,
	}
	// Allow an error of half of the expected probability (arbitrary).
	delta := float64(wantNum) * 0.5

	for k, v := range docTypeDistr {
		want, ok := wantDocTypeDistr[k]
		if !ok {
			t.Errorf("DocumentType in generated Document got %v, want DocumentType to be in %v", k, wantDocTypeDistr)
		}
		if math.Abs(float64(v)-float64(want)) >= delta {
			t.Errorf("docTypeDistr[%v] = %d, want within %v of %d", k, v, delta, want)
		}
	}
}

func TestDocument_ContentLine(t *testing.T) {
	textGenerator := &testtext.Generator{Text: []string{"sentence 1", "sentence 2", "sentence 3"}}
	tests := []struct {
		name string
		d    *pathway.Document
		want []string
	}{{
		name: "Interval with 0",
		d: &pathway.Document{
			DocumentType:          "DS",
			NumRandomContentLines: &pathway.Interval{From: 0, To: 0},
		},
		want: nil,
	}, {
		name: "Random 5 sentences",
		d: &pathway.Document{
			DocumentType:          "DS",
			NumRandomContentLines: &pathway.Interval{From: 4, To: 4},
		},
		want: []string{"sentence 1", "sentence 2", "sentence 3", "sentence 1"},
	}, {
		name: "Empty interval",
		d: &pathway.Document{
			DocumentType:          "DS",
			NumRandomContentLines: &pathway.Interval{},
		},
		want: nil,
	}, {
		name: "Ending only",
		d: &pathway.Document{
			DocumentType:          "DS",
			NumRandomContentLines: &pathway.Interval{},
			EndingContentLines:    []string{"ending 1", "ending 2"},
		},
		want: []string{"ending 1", "ending 2"},
	}, {
		name: "Random and ending",
		d: &pathway.Document{
			DocumentType:          "DS",
			NumRandomContentLines: &pathway.Interval{From: 4, To: 4},
			EndingContentLines:    []string{"ending 1", "ending 2"},
		},
		want: []string{"sentence 1", "sentence 2", "sentence 3", "sentence 1", "ending 1", "ending 2"},
	}, {
		name: "Header and Random",
		d: &pathway.Document{
			DocumentType:          "DS",
			NumRandomContentLines: &pathway.Interval{From: 4, To: 4},
			HeaderContentLines:    []string{"header 1", "header 2"},
		},
		want: []string{"header 1", "header 2", "sentence 1", "sentence 2", "sentence 3", "sentence 1"},
	}, {
		name: "Header only",
		d: &pathway.Document{
			DocumentType:          "DS",
			NumRandomContentLines: &pathway.Interval{},
			HeaderContentLines:    []string{"header 1", "header 2"},
		},
		want: []string{"header 1", "header 2"},
	}, {
		name: "Header and Random and Ending",
		d: &pathway.Document{
			DocumentType:          "DS",
			NumRandomContentLines: &pathway.Interval{From: 4, To: 4},
			HeaderContentLines:    []string{"header 1", "header 2"},
			EndingContentLines:    []string{"ending 1", "ending 2"},
		},
		want: []string{"header 1", "header 2", "sentence 1", "sentence 2", "sentence 3", "sentence 1", "ending 1", "ending 2"},
	}, {
		name: "Header and Ending",
		d: &pathway.Document{
			DocumentType:          "DS",
			NumRandomContentLines: &pathway.Interval{},
			HeaderContentLines:    []string{"header 1", "header 2"},
			EndingContentLines:    []string{"ending 1", "ending 2"},
		},
		want: []string{"header 1", "header 2", "ending 1", "ending 2"},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := Generator{DocumentConfig: &HL7Document, TextGenerator: textGenerator}
			gotDoc := g.Document(date, tc.d)
			if diff := cmp.Diff(tc.want, gotDoc.ContentLine); diff != "" {
				t.Errorf("g.Document(%v, %+v).ContentLine got diff:\n%s", date, tc.d, diff)
			}
		})
	}
}
