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
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/google/simhospital/pkg/ir"
)

// Output defines an object which returns a writer that resource are written to.
type Output interface {
	New(*ir.PatientInfo) (io.WriteCloser, error)
}

// StdOutput is a resource output that wraps stdout.
type StdOutput struct{}

// New just returns os.Stdout.
func (o *StdOutput) New(*ir.PatientInfo) (io.WriteCloser, error) {
	return os.Stdout, nil
}

// DirectoryOutput is a resource output that stores information in multiple files in the same
// directory. Each file stores resources for a single patient at a specific point in time. Files
// are never updated.
type DirectoryOutput struct {
	path  string
	count map[string]int
}

// New returns a new file as a writer for the given patient. It generates filenames as the
// concatenation of the person's full name and MRN (to avoid the slim possibility of collisions).
// If resources are generated more than once for the same Patient, a suffix is added. Consecutive
// writes will append to the file.
func (o *DirectoryOutput) New(p *ir.PatientInfo) (io.WriteCloser, error) {
	pe := p.Person
	filename := strings.Join([]string{pe.FirstName, pe.MiddleName, pe.Surname, pe.MRN}, "_")
	path := path.Join(o.path, filename)

	// Avoid name collisions when generating resources multiple times for the same person.

	if c := o.count[filename]; c > 0 {
		path = fmt.Sprintf("%s_%d", path, c)
	}
	o.count[filename]++

	return os.Create(path)
}

// NewDirectoryOutput returns a new DirectoryOutput based on the given path.
func NewDirectoryOutput(path string) (*DirectoryOutput, error) {
	if path == "" {
		return nil, errors.New("path is empty")
	}

	// Create the directory if it does not already exist.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.Mkdir(path, os.ModeDir); err != nil {
			return nil, errors.Wrapf(err, "cannot create directory %q", path)
		}
	} else if err != nil {
		return nil, errors.Wrapf(err, "cannot create Directory Output using %q", path)
	}

	return &DirectoryOutput{path: path, count: make(map[string]int)}, nil
}
