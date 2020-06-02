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

// Package document contains functions needed to generate a ir.Document object.
package document

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/generator/text"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/pathway"
)

const (
	// minLines is the minimum number of lines in a generated document.
	minLines = 10
	// maxLines is the maximum number of lines in a generated document.
	maxLines = 50
	// udnLength is the set number of characters in a UniqueDocumentNumber.
	udnLength = 13
	// completionStatusDocumented corresponds to a completion status of Documented.
	// This is the default completion status.
	completionStatusDocumented = "DO"
	// obsID is an arbitrary string to be used as the default ID and Text of the Observation Identifier.
	obsID = "Established Patient 15"
	// obsCS is an arbitrary string to be used as the default Coding System of the Observation Identifier.
	obsCS = "Simulated Hospital"
)

var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// Generator generates a document.
type Generator struct {
	DocumentConfig *config.HL7Document
	TextGenerator  text.Generator
}

// Document returns a Document from the given configuration.
// See pathway.Document for information on how every field is populated.
// A UniqueDocumentNumber is randomly generated.
// ActivityDateTime and EditDateTime are set to eventTime.
func (g *Generator) Document(eventTime time.Time, d *pathway.Document) *ir.Document {
	e := ir.NewValidTime(eventTime)
	docType := d.DocumentType
	status := d.CompletionStatus
	id := obsID
	text := obsID
	cs := obsCS

	if docType == "" {
		docType = g.DocumentConfig.Types[rand.Intn(len(g.DocumentConfig.Types))]
	}
	if status == "" {
		status = completionStatusDocumented
	}
	if d.ObsIdentifierID != nil {
		id = *d.ObsIdentifierID
	}
	if d.ObsIdentifierText != nil {
		text = *d.ObsIdentifierText
	}
	if d.ObsIdentifierCS != nil {
		cs = *d.ObsIdentifierCS
	}
	return &ir.Document{
		ActivityDateTime:         e,
		EditDateTime:             e,
		DocumentCompletionStatus: status,
		DocumentType:             docType,
		ObservationIdentifier: &ir.CodedElement{
			ID:           id,
			Text:         text,
			CodingSystem: cs,
		},
		UniqueDocumentNumber: randomUniqueDocumentNumber(),
		ContentLine:          g.content(d),
	}
}

// UpdateDocumentContent updates an existing document, 'dm', content in place based on a pathway.Document configuration, 'dp'.
// Update type is determined in the pathway can be set to "append" or "overwrite".
// An error is returned if the set update type is not one of the above mentioned.
func (g *Generator) UpdateDocumentContent(dm *ir.Document, dp *pathway.Document) error {
	newContent := g.content(dp)
	switch dp.UpdateType {
	case pathway.Append:
		dm.ContentLine = append(dm.ContentLine, newContent...)
	case pathway.Overwrite:
		dm.ContentLine = newContent
	default:
		return fmt.Errorf("update type was %q, expected %q or %q", dp.UpdateType, pathway.Append, pathway.Overwrite)
	}
	return nil
}

func (g *Generator) content(d *pathway.Document) []string {
	randLen := 0
	if d.NumRandomContentLines == nil {
		randLen = g.defaultRandomNumContentLines()
	} else {
		randLen = d.NumRandomContentLines.Random()
	}
	var contentLine []string
	if len(d.HeaderContentLines) != 0 {
		contentLine = append(contentLine, d.HeaderContentLines...)
	}
	if randLen != 0 {
		contentLine = append(contentLine, g.TextGenerator.Sentences(randLen)...)
	}
	if len(d.EndingContentLines) != 0 {
		contentLine = append(contentLine, d.EndingContentLines...)
	}
	return contentLine
}

func randomUniqueDocumentNumber() string {
	udn := make([]rune, udnLength)
	for i := range udn {
		udn[i] = chars[rand.Intn(len(chars))]
	}
	return string(udn)
}

func (g Generator) defaultRandomNumContentLines() int {
	i := &pathway.Interval{To: maxLines, From: minLines}
	return i.Random()
}
