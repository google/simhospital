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
	"path/filepath"
	"testing"
)

const (
	defaultTmpFileName = "arbitrary-file"
	defaultTmpDirName  = "arbitrary-dir"
)

// BytesToFile writes the given bytes to an arbitrary file and returns the file name.
// Different invocations return different file names.
// The file is removed when the test finishes.
func BytesToFile(t *testing.T, b []byte) string {
	t.Helper()
	dir := TempDir(t)
	return BytesToFileInExistingDir(t, b, dir, defaultTmpFileName)
}

// BytesToFileInExistingDir writes the given bytes to a file called name in the given directory
// and returns the full file name.
// dir must be an existing directory.
// The file is removed when the test finishes.
func BytesToFileInExistingDir(t *testing.T, b []byte, dir string, name string) string {
	t.Helper()
	full := filepath.Join(dir, name)
	// Only read and write permission for the owner is required as the file is only
	// created, read and then deleted by the same user.
	if err := ioutil.WriteFile(full, b, 0600); err != nil {
		t.Fatalf("WriteFile(%s) failed with %v", full, err)
	}

	t.Cleanup(func() { os.Remove(full) })
	return full
}

// BytesToDir writes the given bytes to a file named fileName inside an arbitrary directory
// and returns the name of the directory.
// The file and the directory are removed when the test finishes.
func BytesToDir(t *testing.T, b []byte, fileName string) string {
	t.Helper()
	dir := TempDir(t)
	BytesToFileInExistingDir(t, b, dir, fileName)
	return dir
}

// TempDir creates a new directory.
// Invoking this method multiple times returns different results.
// The directory and all its contents are removed when the test finishes.
func TempDir(t *testing.T) string {
	t.Helper()
	dir, err := ioutil.TempDir("", defaultTmpDirName)
	if err != nil {
		t.Fatalf("ioutil.TempDir(%s, %s) failed with %v", "", defaultTmpDirName, err)
	}
	t.Cleanup(func() { os.RemoveAll(dir) })
	return dir
}
