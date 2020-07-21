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

package files

import (
	"context"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/test/testwrite"
)

func TestListLocal(t *testing.T) {
	ctx := context.Background()
	dir := testwrite.TempDir(t)
	fileData := map[string]string{
		"file1": "file1.txt",
		"file2": "file2.txt",
		"file3": "file3.txt",
	}
	for name, data := range fileData {
		testwrite.BytesToFileInExistingDir(t, []byte(data), dir, name)
	}

	files, err := List(ctx, dir)
	if err != nil {
		t.Fatalf("List(%s) failed with %v", dir, err)
	}
	gotFileData := make(map[string]string)
	for _, f := range files {
		data, err := f.Read(ctx)
		if err != nil {
			t.Errorf("could not read %s: %v", f.Name(), err)
		}
		gotFileData[f.Name()] = string(data)
	}
	if diff := cmp.Diff(fileData, gotFileData); diff != "" {
		t.Errorf("List(%s) returned unexpected file data, diff (-want,+got):\n%s", dir, diff)
	}
}

func TestReadLocal(t *testing.T) {
	ctx := context.Background()
	data := []byte("this is a test")
	f := testwrite.BytesToFile(t, data)
	v, err := Read(ctx, f)
	if err != nil {
		t.Fatalf("Read(%s) failed with %v", f, err)
	}
	if got, want := string(v), string(data); got != want {
		t.Errorf("Read(%s) got: %v, want: %v", f, got, want)
	}
}

func TestReadLocalErrors(t *testing.T) {
	ctx := context.Background()
	dir := testwrite.TempDir(t)
	tests := []struct {
		name   string
		path   string
		errStr string
	}{
		{
			name:   "path is directory",
			path:   dir,
			errStr: "is a directory",
		}, {
			name:   "path doesn't exist",
			path:   "path/to/anything.txt",
			errStr: "no such file",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Read(ctx, tt.path)
			if err == nil {
				t.Fatalf("Read(%s) = %v, want error", tt.path, data)
			}
			if !strings.Contains(err.Error(), tt.errStr) {
				t.Errorf("Read(%s) = (%v, %v), want: (nil, ...%s...)", tt.path, data, err, tt.errStr)
			}
		})
	}
}

// TODO. No tests for reading from GCS as the current testing infrastructure setup doesn't allow for it.
