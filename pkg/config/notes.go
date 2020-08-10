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

// Package config provides functionality for the configuration of Simulated Hospital, especially around loading such configuration from files.
package config

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/google/simhospital/pkg/files"
)

// ClinicalNote represents a document that contains a clinical note.
type ClinicalNote struct {
	DocumentFileName string
	Path             string
	ContentType      string
}

// LoadNotesConfig loads all the files in the sample notes directory. Each sample note in the directory have to be named
// in the <filename>.<extension> format. We rely on this formatting to extract the ContentType of the
// document. Files not named in the right format will be skipped.
func LoadNotesConfig(ctx context.Context, directory string) (map[string][]ClinicalNote, error) {
	files, err := files.List(ctx, directory)
	if err != nil {
		return nil, errors.Wrapf(err, "Error reading Sample Notes directory %s", directory)
	}

	typeToNotes := make(map[string][]ClinicalNote)
	for _, f := range files {
		split := strings.Split(f.Name(), ".")
		if len(split) <= 1 {
			log.WithField("sample_notes_directory", directory).
				Warningf("Sample file name %s is in the wrong format, want <filename>.<extension> format. Skipping...", f.Name())
			continue
		}
		ext := split[len(split)-1]

		typeToNotes[ext] = append(typeToNotes[ext], ClinicalNote{
			DocumentFileName: f.Name(),
			Path:             f.FullPath(),
			ContentType:      ext,
		})
	}
	return typeToNotes, nil
}
