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

package config

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/test/testwrite"
)

type note struct {
	content  string
	filename string
}

func TestLoadNotesConfigFromDir(t *testing.T) {
	ctx := context.Background()
	mainDir := testwrite.TempDir(t)

	testCases := []struct {
		name  string
		notes []note
		want  map[string][]ClinicalNote
	}{{
		name: "single png file",
		notes: []note{{
			content:  "last png content",
			filename: "test1.png",
		}},
		want: map[string][]ClinicalNote{
			"png": {{
				DocumentFileName: "test1.png",
				Path:             filepath.Join(mainDir, "test1.png"),
				ContentType:      "png",
			}},
		},
	}, {
		name: "two jpg files",
		notes: []note{{
			content:  "meme content",
			filename: "test1.jpg",
		}, {
			content:  "more meme content",
			filename: "test2.jpg",
		}},
		want: map[string][]ClinicalNote{
			"jpg": {{
				DocumentFileName: "test1.jpg",
				Path:             filepath.Join(mainDir, "test1.jpg"),
				ContentType:      "jpg",
			}, {
				DocumentFileName: "test2.jpg",
				Path:             filepath.Join(mainDir, "test2.jpg"),
				ContentType:      "jpg",
			}},
		},
	}, {
		name: "2 of jpg, rtf and pdf each",
		notes: []note{{
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
		}},
		want: map[string][]ClinicalNote{
			"pdf": {{
				DocumentFileName: "test1.pdf",
				Path:             filepath.Join(mainDir, "test1.pdf"),
				ContentType:      "pdf",
			}, {
				DocumentFileName: "test2.pdf",
				Path:             filepath.Join(mainDir, "test2.pdf"),
				ContentType:      "pdf",
			}},
			"rtf": {{
				DocumentFileName: "test1.rtf",
				Path:             filepath.Join(mainDir, "test1.rtf"),
				ContentType:      "rtf",
			}, {
				DocumentFileName: "test2.rtf",
				Path:             filepath.Join(mainDir, "test2.rtf"),
				ContentType:      "rtf",
			}},
			"jpg": {{
				DocumentFileName: "test1.jpg",
				Path:             filepath.Join(mainDir, "test1.jpg"),
				ContentType:      "jpg",
			}, {
				DocumentFileName: "test2.jpg",
				Path:             filepath.Join(mainDir, "test2.jpg"),
				ContentType:      "jpg",
			}},
		},
	}, {
		name: "attempting to read a non existing file",
		notes: []note{{
			content:  "new content",
			filename: "testfilenamenotrightformat",
		}},
		want: map[string][]ClinicalNote{},
	}}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			for _, n := range tt.notes {
				content := []byte(n.content)
				testwrite.BytesToFileInExistingDir(t, content, mainDir, n.filename)
			}
			got, err := LoadNotesConfig(ctx, mainDir)
			if err != nil {
				t.Fatalf("LoadNotesConfig(%s) failed with error %v", mainDir, err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("LoadNotesConfig(%s) got diff (-want, +got):\n%s", mainDir, diff)
			}
		})
	}
}
