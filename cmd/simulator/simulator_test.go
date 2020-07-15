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

package main

import (
	"context"
	"flag"
	"os"
	"path"
	"testing"

	"github.com/google/simhospital/pkg/test"
)

var (
	base = path.Join(currentDir())
	// These tests use the default value of the flags, which point to configuration files.
	// testDep is a dummy variable that points to the test library that loads those files for testing.
	testDep = test.Prod
)

func TestRunner(t *testing.T) {
	ctx := context.Background()
	// Prevent the connection to an MLLP server.
	if err := flag.Set("output", "stdout"); err != nil {
		t.Errorf("flag.Set(%v, %v) failed with %v", "output", "stdout", err)
	}

	if _, err := createRunner(ctx); err == nil {
		t.Error("createRunner() got nil error, want non nil error")
	}

	if err := flag.Set("local_path", base); err != nil {
		t.Errorf("flag.Set(%v, %v) failed with %v", "local_path", base, err)
	}
	if _, err := createRunner(ctx); err != nil {
		t.Errorf("createRunner(local_path=%v) failed with %v", base, err)
	}
}

func currentDir() string {
	dir, _ := os.Getwd()
	return dir
}
