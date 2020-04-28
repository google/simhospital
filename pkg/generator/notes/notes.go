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

// Package notes contains functions needed to generate a message.ClinicalNote object given the pathway.ClinicalNote object.
package notes

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/generator/text"
	"github.com/google/simhospital/pkg/message"
	"github.com/google/simhospital/pkg/pathway"
)

const (
	// defaultNumSentences is a number of random sentences to generate.
	defaultNumSentences = 20

	// File types.
	txt = "txt"
	pdf = "pdf"
	jpg = "jpg"
	png = "png"
)

// Generator generates random clinical notes.
type Generator struct {
	config        map[string][]config.ClinicalNote
	types         []string
	textGenerator *text.Generator
	numSentences  int
}

// NewGenerator returns a new Generator struct.
func NewGenerator(d *config.Data, t *text.Generator) *Generator {
	return &Generator{
		config:        d.NotesConfig,
		types:         d.ClinicalNoteTypes,
		textGenerator: t,
		numSentences:  defaultNumSentences,
	}
}

// RandomDocumentForClinicalNote generates a ClinicalNote from the given pathway event.
// If the content type is a txt, random text is generated as the content.
// Otherwise a random file matching the content type is read from the list of sample files as the content.
func (g *Generator) RandomDocumentForClinicalNote(np *pathway.ClinicalNote) (*message.ClinicalNote, error) {
	note := &message.ClinicalNote{
		DocumentType: g.documentType(np.DocumentType),
		ContentType:  np.ContentType,
		DocumentID:   g.id(np.DocumentID),
	}

	if np.DocumentContent != "" {
		if np.ContentType != txt {
			return nil, fmt.Errorf("cannot give explicit content for ContentType %q, only ContentType %q is supported",
				np.ContentType, txt)
		}
		note.DocumentContent = np.DocumentContent
		return note, nil
	}

	switch np.ContentType {
	case txt:
		sentences := g.textGenerator.RandomSentences(g.numSentences)
		note.DocumentContent = strings.Join(sentences, ". ")
	// pdf, jpg and png document contents can contain delimiters used in HL7 messages eg, pipes.
	// So these files are encoded in base64 to make sure the HL7 messages are parsable.
	case pdf, jpg, png:
		content, err := g.note(np.ContentType)
		if err != nil {
			return nil, errors.Wrapf(err, "generate note for ContentType %q", np.ContentType)
		}
		note.DocumentContent = base64.StdEncoding.EncodeToString(content)
		note.DocumentEncoding = "base64"
	// rtf, xhtml etc
	default:
		content, err := g.note(np.ContentType)
		if err != nil {
			return nil, errors.Wrapf(err, "generate note for ContentType %q", np.ContentType)
		}
		note.DocumentContent = string(content)
	}
	return note, nil
}

// RandomNotesForResult generates between 0 to 2 notes, with the following probabilities:
// 0.4 - no notes
// 0.5 - 1 note
// 0.1 - 2 notes
// Each note has between 1 - 10 random words.
func (g *Generator) RandomNotesForResult() []string {
	switch r := rand.Intn(10); {
	case r < 4:
		return nil
	case r < 9:
		return g.textGenerator.RandomSentences(1)
	default:
		return g.textGenerator.RandomSentences(2)
	}
}

func (g *Generator) note(contentType string) ([]byte, error) {
	notes, ok := g.config[contentType]
	if !ok || len(notes) == 0 {
		return nil, fmt.Errorf("no sample Notes found for %s ContentType: ContentType not supported", contentType)
	}
	clinicalNote := notes[rand.Intn(len(notes))]
	return ioutil.ReadFile(clinicalNote.Path)
}

// id returns the given ID if not empty, otherwise generates a random string with "random-" prefix.
func (g *Generator) id(currID string) string {
	if currID != "" {
		return currID
	}
	var buffer bytes.Buffer
	for i := 0; i < 10; i++ {
		buffer.WriteString(strconv.Itoa(rand.Intn(10)))
	}
	return fmt.Sprintf("random-%v", buffer.String())
}

// documentType returns the given type if not empty,
// otherwise returns a random Clinical Note type from the list of types.
func (g *Generator) documentType(currType string) string {
	if currType != "" {
		return currType
	}
	return g.types[rand.Intn(len(g.types))]
}
