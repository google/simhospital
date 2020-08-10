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

package notes

import (
	"context"
	"encoding/base64"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/generator/text"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/pathway"
	"github.com/google/simhospital/pkg/test/testwrite"
)

var (
	noteTypes   = []string{"ED Depart Summary", "Gynaecology & Maternity", "Surgery Inpatients"}
	nouns       = []string{"one", "two", "three", "four", "five"}
	defaultDate = time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)
)

const testNumSentences = 15

type wantContent struct {
	wantEncoding            string
	wantExplicitContent     bool
	wantObservationDateTime ir.NullTime
}

type tuple struct {
	pathway          *pathway.ClinicalNote
	contents         []*wantContent
	wantDocumentType string
	eventDateTime    time.Time
}

func TestRandomDocumentForClinicalNote(t *testing.T) {
	rand.Seed(1)
	ctx := context.Background()
	nc, fc := testSetup(ctx, t)
	g := &Generator{
		textGenerator: &text.NounGenerator{Nouns: nouns},
		config:        nc,
		types:         noteTypes,
		numSentences:  testNumSentences,
	}
	testCases := []struct {
		name         string
		wantRandomID bool
		wantDateTime ir.NullTime
		tuples       []*tuple
	}{{
		name: "pdf, id is non-explicit, base64 encoding",
		tuples: []*tuple{{
			pathway: &pathway.ClinicalNote{
				ContentType: "pdf",
			},
			contents: []*wantContent{{
				wantEncoding:            "base64",
				wantObservationDateTime: ir.NewValidTime(defaultDate),
			}},
			eventDateTime: defaultDate,
		}},
		wantRandomID: true,
		wantDateTime: ir.NewValidTime(defaultDate),
	}, {
		name: "rtf, id is non-explicit, no encoding",
		tuples: []*tuple{{
			pathway: &pathway.ClinicalNote{
				ContentType: "rtf",
			},
			contents: []*wantContent{{
				wantObservationDateTime: ir.NewValidTime(defaultDate),
			}},
			eventDateTime: defaultDate,
		}},
		wantRandomID: true,
		wantDateTime: ir.NewValidTime(defaultDate),
	}, {
		name: "jpg, id is explicit, base64 encoding",
		tuples: []*tuple{{
			pathway: &pathway.ClinicalNote{
				ContentType: "jpg",
				DocumentID:  "some-id",
			},
			contents: []*wantContent{{
				wantEncoding:            "base64",
				wantObservationDateTime: ir.NewValidTime(defaultDate),
			}},
			eventDateTime: defaultDate,
		}},
		wantDateTime: ir.NewValidTime(defaultDate),
	}, {
		name: "pdf, id is explicit, base64 encoding",
		tuples: []*tuple{{
			pathway: &pathway.ClinicalNote{
				ContentType: "pdf",
				DocumentID:  "some-id",
			},
			contents: []*wantContent{{
				wantEncoding:            "base64",
				wantObservationDateTime: ir.NewValidTime(defaultDate),
			}},
			eventDateTime: defaultDate,
		}},
		wantDateTime: ir.NewValidTime(defaultDate),
	}, {
		name: "txt, document type explicitly given",
		tuples: []*tuple{{
			pathway: &pathway.ClinicalNote{
				ContentType:  "txt",
				DocumentType: "ED",
				DocumentID:   "some-id",
			},
			contents: []*wantContent{{
				wantObservationDateTime: ir.NewValidTime(defaultDate),
			}},
			wantDocumentType: "ED",
			eventDateTime:    defaultDate,
		}},
		wantDateTime: ir.NewValidTime(defaultDate),
	}, {
		name: "txt, explicit content given",
		tuples: []*tuple{{
			pathway: &pathway.ClinicalNote{
				ContentType:     "txt",
				DocumentContent: "short content",
			},
			contents: []*wantContent{{
				wantExplicitContent:     true,
				wantObservationDateTime: ir.NewValidTime(defaultDate),
			}},
			eventDateTime: defaultDate,
		}},
		wantRandomID: true,
		wantDateTime: ir.NewValidTime(defaultDate),
	}, {
		name: "two pathways with same id, request a pdf, and then a text addendum with given content",
		tuples: []*tuple{{
			pathway: &pathway.ClinicalNote{
				ContentType: "pdf",
				DocumentID:  "some-id",
			},
			contents: []*wantContent{{
				wantEncoding:            "base64",
				wantObservationDateTime: ir.NewValidTime(defaultDate),
			}},
			eventDateTime: defaultDate,
		}, {
			pathway: &pathway.ClinicalNote{
				ContentType:     "txt",
				DocumentContent: "new content on addendum",
				DocumentID:      "some-id",
			},
			contents: []*wantContent{{
				wantEncoding:            "base64",
				wantObservationDateTime: ir.NewValidTime(defaultDate),
			}, {
				wantExplicitContent:     true,
				wantObservationDateTime: ir.NewValidTime(time.Date(2019, 1, 21, 0, 0, 0, 0, time.UTC)),
			}},
			eventDateTime: time.Date(2019, 1, 21, 0, 0, 0, 0, time.UTC),
		}},
		wantDateTime: ir.NewValidTime(defaultDate),
	}, {
		name: "two pathways with same id, request a txt, and then a txt addendum with new title and type",
		tuples: []*tuple{{
			pathway: &pathway.ClinicalNote{
				ContentType:   "txt",
				DocumentID:    "some-id",
				DocumentType:  "type",
				DocumentTitle: "title",
			},
			contents: []*wantContent{{
				wantObservationDateTime: ir.NewValidTime(defaultDate),
			}},
			wantDocumentType: "type",
			eventDateTime:    defaultDate,
		}, {
			pathway: &pathway.ClinicalNote{
				ContentType:   "txt",
				DocumentID:    "some-id",
				DocumentType:  "new-type",
				DocumentTitle: "new-title",
			},
			contents: []*wantContent{{
				wantObservationDateTime: ir.NewValidTime(defaultDate),
			}, {
				wantObservationDateTime: ir.NewValidTime(time.Date(2019, 1, 21, 0, 0, 0, 0, time.UTC)),
			}},
			wantDocumentType: "new-type",
			eventDateTime:    time.Date(2019, 1, 21, 0, 0, 0, 0, time.UTC),
		}},
		wantDateTime: ir.NewValidTime(defaultDate),
	}, {
		name: "two pathways with same id, request a txt with a type and title, and then a txt addendum with empty title and type",
		tuples: []*tuple{{
			pathway: &pathway.ClinicalNote{
				ContentType:   "txt",
				DocumentID:    "some-id",
				DocumentTitle: "title",
			},
			contents: []*wantContent{{
				wantObservationDateTime: ir.NewValidTime(defaultDate),
			}},
			eventDateTime: defaultDate,
		}, {
			pathway: &pathway.ClinicalNote{
				ContentType: "txt",
				DocumentID:  "some-id",
			},
			contents: []*wantContent{{
				wantObservationDateTime: ir.NewValidTime(defaultDate),
			}, {
				wantObservationDateTime: ir.NewValidTime(time.Date(2019, 1, 21, 0, 0, 0, 0, time.UTC)),
			}},
			eventDateTime: time.Date(2019, 1, 21, 0, 0, 0, 0, time.UTC),
		}},
		wantDateTime: ir.NewValidTime(defaultDate),
	}}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var (
				got *ir.ClinicalNote
				err error
			)
			for index, tuple := range tt.tuples {
				t.Run(fmt.Sprintf("pair index: %d", index), func(t *testing.T) {
					got, err = g.RandomDocumentForClinicalNote(ctx, tuple.pathway, got, tuple.eventDateTime)
					if err != nil {
						t.Fatalf("g.RandomDocumentForClinicalNote(%v) failed with error %v", tuple.pathway, err)
					}
					if got == nil {
						t.Fatalf("g.RandomDocumentForClinicalNote(%v)=<nil>, want *ir.ClinicalNote{...} instance", tuple.pathway)
					}
					if gotRandomID := strings.HasPrefix(got.DocumentID, "random-"); tt.wantRandomID != gotRandomID {
						t.Errorf("g.RandomDocumentForClinicalNote(%v) got DocumentID=%s, is random? %t, want random? %t",
							tuple.pathway, got.DocumentID, gotRandomID, tt.wantRandomID)
					}
					if tuple.wantDocumentType == "" {
						if !contains(got.DocumentType, g.types) {
							t.Errorf("g.RandomDocumentForClinicalNote(%v) got DocumentType=%s, want one of %v",
								tuple.pathway, got.DocumentType, g.types)
						}
					} else if got, want := got.DocumentType, tuple.wantDocumentType; got != want {
						t.Errorf("g.RandomDocumentForClinicalNote(%v) got DocumentType=%s, want %s",
							tuple.pathway, got, want)
					}
					if diff := cmp.Diff(tt.wantDateTime, got.DateTime); diff != "" {
						t.Errorf("g.RandomDocumentForClinicalNote(%v).DateTime mismatch (-want, +got)=\n%s", tuple.pathway, diff)
					}
					for contentIndex, content := range tuple.contents {
						if got, want := content.wantEncoding, got.Contents[contentIndex].DocumentEncoding; got != want {
							t.Errorf("content index: %d; g.RandomDocumentForClinicalNote(%v) got DocumentEncoding=%s, want %s",
								contentIndex, tuple.pathway, got, want)
						}
						if tuple.pathway.ContentType != "txt" && !contains(got.Contents[contentIndex].DocumentContent, fc[tuple.pathway.ContentType]) {
							t.Errorf("content index: %d; g.RandomDocumentForClinicalNote(%v) got DocumentContent=%s, want one of %v",
								contentIndex, tuple.pathway, got.Contents[index].DocumentContent, fc[tuple.pathway.ContentType])
						}
						if got, want := got.Contents[contentIndex].DocumentContent, tuple.pathway.DocumentContent; content.wantExplicitContent && got != want {
							t.Errorf("content index: %d; g.RandomDocumentForClinicalNote(%v) got DocumentContent=%s, want %s",
								contentIndex, tuple.pathway, got, want)
						}
						if diff := cmp.Diff(content.wantObservationDateTime, got.Contents[contentIndex].ObservationDateTime); diff != "" {
							t.Errorf("content index: %d; g.RandomDocumentForClinicalNote(%v) got ObservationDateTime mismatch (-want, +got)=\n%s",
								contentIndex, tuple.pathway, diff)
						}
						if got.Contents[0].ContentType != "txt" || content.wantExplicitContent {
							return
						}
						sentences := strings.Split(got.Contents[index].DocumentContent, ".")
						if got, want := len(sentences), testNumSentences; got != want {
							t.Errorf("content index: %d; len(g.RandomDocumentForClinicalNote(%v)) got len(DocumentContent)=%v, want %v",
								contentIndex, tuple.pathway, got, want)
						}
					}
				})
			}
		})
	}
}

func TestRandomNotesForResult(t *testing.T) {
	rand.Seed(1)
	g := Generator{textGenerator: &text.NounGenerator{Nouns: nouns}}
	runs := float64(1000)
	runsWithNotes := float64(0)
	noteLenDistr := map[int]int{}
	for i := 0; i < int(runs); i++ {
		notes := g.RandomNotesForResult()
		l := len(notes)
		if l > 2 {
			// This check is repeated later, but here we print the notes which is useful to know what went wrong.
			t.Errorf("len(notes)=%d, want no more than 2. Notes: %v", l, notes)
		}
		noteLenDistr[l]++
		if len(notes) > 0 {
			runsWithNotes++
		}
	}
	// Check that all Notes have 0, 1 or 2 sentences with 0.4, 0.5 and 0.1 probability respectively.
	wantNoteLenDistr := map[int]int{
		0: int(runs * 0.4),
		1: int(runs * 0.5),
		2: int(runs * 0.1),
	}
	// Allow an error of 5% of the number of runs.
	delta := runs / 20
	for k, v := range noteLenDistr {
		want, ok := wantNoteLenDistr[k]
		if !ok {
			t.Errorf("len(notes) got %v, want length to be in %v", k, wantNoteLenDistr)
		}
		if math.Abs(float64(v)-float64(want)) >= delta {
			t.Errorf("noteLenDistr[%v] = %v, want within %v of %v", k, v, delta, want)
		}
	}
}

func TestRandomInvalidContentTypeError(t *testing.T) {
	ctx := context.Background()
	nc, _ := testSetup(ctx, t)
	g := &Generator{
		textGenerator: &text.NounGenerator{Nouns: nouns},
		config:        nc,
		types:         noteTypes,
		numSentences:  testNumSentences,
	}
	for _, tt := range []struct {
		name    string
		pathway *pathway.ClinicalNote
	}{{
		name:    "ContentType not supported",
		pathway: &pathway.ClinicalNote{ContentType: "doc"},
	}, {
		name:    "Empty ContentType",
		pathway: &pathway.ClinicalNote{},
	}, {
		name: "Explicit content for pdf",
		pathway: &pathway.ClinicalNote{
			ContentType:     "pdf",
			DocumentContent: "some content",
		},
	}} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := g.RandomDocumentForClinicalNote(ctx, tt.pathway, nil, time.Time{})
			if err == nil {
				t.Errorf("g.RandomDocumentForClinicalNote(%q) = %v, %v; want <nil>, error", tt.pathway, got, err)
			}
		})
	}
}

func TestRandomUniformDistribution(t *testing.T) {
	rand.Seed(1)
	ctx := context.Background()
	nc, fc := testSetup(ctx, t)
	g := &Generator{
		textGenerator: &text.NounGenerator{Nouns: nouns},
		config:        nc,
		types:         noteTypes,
	}
	arbitraryClinicalNotes := []*pathway.ClinicalNote{{
		ContentType: "jpg",
	}, {
		ContentType: "rtf",
	}, {
		ContentType: "pdf",
	}}
	runs := float64(1000)
	contentDistr := map[string]int{}
	typesDistr := map[string]int{}
	for _, acn := range arbitraryClinicalNotes {
		for i := 0; i < int(runs); i++ {
			n, err := g.RandomDocumentForClinicalNote(ctx, acn, nil, time.Time{})
			if err != nil {
				t.Fatalf("g.RandomDocumentForClinicalNote(%v) failed with %v", acn, err)
			}
			contentDistr[n.Contents[0].DocumentContent]++
			typesDistr[n.DocumentType]++
		}
	}
	// Each ContentType in arbitraryClinicalNotes is run 1000 times.
	// For uniform distribution we expect each available note to have selected
	// equally, eg: If there are 5 sample notes for an extension, each of them
	// to have selected roughly 1000 / 5 = 200 times.
	wantContentDistr := map[string]int{}
	for _, v := range fc {
		numOfNotes := float64(len(v))
		for _, content := range v {
			wantContentDistr[content] = int(runs / numOfNotes)
		}
	}
	// Allow an error of 5% of the number of runs.
	delta := 0.05 * runs
	for k, v := range contentDistr {
		want, ok := wantContentDistr[k]
		if !ok {
			t.Errorf("content in generated note is %s, want content to be in %v", k, wantContentDistr)
		}
		if math.Abs(float64(v)-float64(want)) >= delta {
			t.Errorf("contentDistr[%s]=%d want within %v of %d", k, v, delta, want)
		}
	}
	// Check all DocumentTypes are selected uniformly.
	// Expected probability is total number of runs divided by number of noteTypes.
	wantP := int(runs * float64(len(arbitraryClinicalNotes)) / float64(len(noteTypes)))
	wantTypesDistr := map[string]int{}
	for _, noteType := range noteTypes {
		wantTypesDistr[noteType] = wantP
	}
	// Allow an error of 5% of total number of runs.
	delta = 0.05 * runs * float64(len(arbitraryClinicalNotes))
	for k, v := range typesDistr {
		want, ok := wantTypesDistr[k]
		if !ok {
			t.Errorf("DocumentType in generated note is %s, want DocumentType to be in %v", k, wantTypesDistr)
		}
		if math.Abs(float64(v)-float64(want)) >= delta {
			t.Errorf("typesDistr[%v] = %d, want within %v of %d", k, v, delta, want)
		}
	}
}

func testSetup(ctx context.Context, t *testing.T) (map[string][]config.ClinicalNote, map[string][]string) {
	t.Helper()
	mainDir := testwrite.TempDir(t)
	fileContents := make(map[string][]string)
	for _, n := range []struct {
		content  string
		filename string
	}{{
		content:  "new content",
		filename: "test2.rtf",
	}, {
		content:  "some rtf content",
		filename: "test1.rtf",
	}, {
		content:  "some jpg content",
		filename: "test1.jpg",
	}, {
		content:  "some more jpg content",
		filename: "test2.jpg",
	}, {
		content:  "long pdf content",
		filename: "test1.pdf",
	}, {
		content:  "short pdf content",
		filename: "test2.pdf",
	}} {
		contentBytes := []byte(n.content)
		split := strings.Split(n.filename, ".")
		if ext := split[len(split)-1]; ext == "jpg" || ext == "pdf" || ext == "png" {
			fileContents[ext] = append(fileContents[ext], base64.StdEncoding.EncodeToString(contentBytes))
		} else {
			fileContents[ext] = append(fileContents[ext], n.content)
		}
		testwrite.BytesToFileInExistingDir(t, contentBytes, mainDir, n.filename)
	}
	nc, err := config.LoadNotesConfig(ctx, mainDir)
	if err != nil {
		t.Fatalf("config.LoadNotesConfig(%s) failed with error %v", mainDir, err)
	}
	return nc, fileContents
}

func contains(str string, strs []string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}
