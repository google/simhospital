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
	"github.com/google/simhospital/pkg/generator/text"
	"github.com/google/simhospital/pkg/message"
	"github.com/google/simhospital/pkg/pathway"
)

var (
	date        = time.Date(2018, 2, 12, 1, 25, 0, 0, time.UTC)
	nouns       = []string{"one", "two", "three", "four", "five"}
	HL7Document = config.HL7Document{
		Types: []string{"AR", "CD", "CN", "DI", "DS"},
	}
	sampleText = []string{"sample-text-1", "sample-text-2"}
)

func TestDocument(t *testing.T) {
	obsID := "id"
	obsText := "text"
	cs := "cs"
	tests := []struct {
		name  string
		input *pathway.Document
		want  *message.Document
	}{{
		name: "Fixed values",
		input: &pathway.Document{
			DocumentType:      "DS",
			CompletionStatus:  "IP",
			ObsIdentifierID:   &obsID,
			ObsIdentifierText: &obsText,
			ObsIdentifierCS:   &cs,
		},
		want: &message.Document{
			ActivityDateTime:         message.NewValidTime(date),
			EditDateTime:             message.NewValidTime(date),
			DocumentType:             "DS",
			DocumentCompletionStatus: "IP",
			ObservationIdentifier: &message.CodedElement{
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
		want: &message.Document{
			ActivityDateTime:         message.NewValidTime(date),
			EditDateTime:             message.NewValidTime(date),
			DocumentType:             "DS",
			DocumentCompletionStatus: "DO",
			ObservationIdentifier: &message.CodedElement{
				ID:           "Established Patient 15",
				Text:         "Established Patient 15",
				CodingSystem: "Simulated Hospital",
			},
		},
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := Generator{DocumentConfig: &HL7Document, TextGenerator: &text.NounGenerator{Nouns: nouns}}
			got := g.Document(date, tc.input)

			// UniqueDocumentNumber is randomly generated with each document generation.
			// ContentLine contains randomly generated sentences with each document generation.
			ignore := cmpopts.IgnoreFields(message.Document{}, "UniqueDocumentNumber", "ContentLine")
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

func TestDocument_RandomDocumentTypeDoesntModifyInput(t *testing.T) {
	dt := &pathway.Document{}
	g := Generator{DocumentConfig: &HL7Document, TextGenerator: &text.NounGenerator{Nouns: nouns}}
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
	g := Generator{DocumentConfig: &HL7Document, TextGenerator: &text.NounGenerator{Nouns: nouns}}
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

func TestDocument_NumContentLines(t *testing.T) {
	d := &pathway.Document{DocumentType: "DS", NumContentLines: &pathway.Interval{From: 30, To: 30}}
	g := Generator{DocumentConfig: &HL7Document, TextGenerator: &text.NounGenerator{Nouns: nouns}}
	gotDoc := g.Document(date, d)
	if got := len(gotDoc.ContentLine); got != 30 {
		t.Errorf("len(g.Document().ContentLine) got %v, want 30", got)
	}
}

func TestDocument_SetEndingContentLines(t *testing.T) {
	d := &pathway.Document{EndingContentLines: sampleText}
	g := Generator{DocumentConfig: &HL7Document, TextGenerator: &text.NounGenerator{Nouns: nouns}}
	gotDoc := g.Document(date, d)
	l := len(gotDoc.ContentLine)
	last := []string{gotDoc.ContentLine[l-2], gotDoc.ContentLine[l-1]}
	want := []string{"sample-text-1", "sample-text-2"}
	if diff := cmp.Diff(want, last); diff != "" {
		t.Errorf("Last two lines of the content returned diff:\n%s", diff)
	}
}
