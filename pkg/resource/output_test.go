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

package resource

import (
	"os"
	"path"
	"testing"

	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/test/testwrite"
)

func TestDirectoryOutput(t *testing.T) {
	tmpDir := testwrite.TempDir(t)
	wantPath1 := path.Join(tmpDir, "FIRSTNAME_MIDDLENAME_SURNAME_MRN")
	wantPath2 := path.Join(tmpDir, "FIRSTNAME_MIDDLENAME_SURNAME_MRN_1")

	o, err := NewDirectoryOutput(tmpDir)
	if err != nil {
		t.Fatalf("NewDirectoryOutput(%s) failed with: %v", tmpDir, err)
	}

	if fileExists(t, wantPath1) {
		t.Errorf("file %s already exists", wantPath1)
	}
	if fileExists(t, wantPath2) {
		t.Errorf("file %s already exists", wantPath2)
	}

	p := &ir.PatientInfo{
		Person: &ir.Person{
			FirstName:  "FIRSTNAME",
			MiddleName: "MIDDLENAME",
			Surname:    "SURNAME",
			MRN:        "MRN",
		},
	}

	w1, err := o.New(p)
	if err != nil {
		t.Fatalf("o.New(%v) failed with: %v", p, err)
	}
	w1.Close()

	if !fileExists(t, wantPath1) {
		t.Errorf("o.New(%v) did not create the wanted file %s", p, wantPath1)
	}

	w2, err := o.New(p)
	if err != nil {
		t.Fatalf("o.New(%v) (second call) failed with: %v", p, err)
	}
	w2.Close()

	if !fileExists(t, wantPath2) {
		t.Errorf("o.New(%v) (second call) did not created the wanted file %s", p, wantPath2)
	}
}

func fileExists(t *testing.T, path string) bool {
	t.Helper()
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	// If we cannot determine whether or not a file exists, something is wrong with our testing
	// environment and we cannot continue.
	t.Fatalf("os.Stat(%s) failed with: %s", path, err)
	return false
}
