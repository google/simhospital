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

// Package testwrite contains functionality to write files for testing.
package testwrite

import (
	"io/ioutil"
	"os"
	"testing"
)

const defaultTestTmpFileName = "arbitrary-file"

// BytesToFile writes the given bytes to an arbitrary file and returns the file name.
func BytesToFile(t *testing.T, b []byte) string {
	t.Helper()
	name := defaultTestTmpFileName
	tmp, err := ioutil.TempFile("", name)
	if err != nil {
		t.Fatalf("TempFile(%s) failed with %v", name, err)
	}
	if _, err = tmp.Write(b); err != nil {
		os.Remove(tmp.Name())
		t.Fatalf("Write(%s) failed with %v", b, err)
	}
	return tmp.Name()
}

// BytesToFileWithName writes the given bytes to a file with the given name.
func BytesToFileWithName(t *testing.T, b []byte, name string) {
	t.Helper()
	// Only read and write permission for the owner is required as the file is only
	// created, read and then deleted by the same user.
	if err := ioutil.WriteFile(name, b, 0600); err != nil {
		t.Fatalf("WriteFile(%s) failed with %v", name, err)
	}
}
