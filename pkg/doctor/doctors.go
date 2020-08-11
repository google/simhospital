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

// Package doctor provides functionality to manage doctors.
package doctor

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"github.com/google/simhospital/pkg/files"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/logging"
)

var log = logging.ForCallerPackage()

// Doctors contains and manages a set of doctors.
type Doctors struct {
	m map[string]*ir.Doctor
	k []string
}

// LoadDoctors loads the doctors from the given file.
// Returns an error if any entry is invalid.
func LoadDoctors(ctx context.Context, filename string) (*Doctors, error) {
	logLocal := log.WithField("file", filename)
	data, err := files.Read(ctx, filename)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot parse doctors from file %q", filename)
	}

	var list []*ir.Doctor
	if err = yaml.UnmarshalStrict(data, &list); err != nil {
		return nil, errors.Wrapf(err, "cannot unmarshal doctors from file %q", filename)
	}

	doctors := &Doctors{
		m: make(map[string]*ir.Doctor),
	}

	logLocal.Info("Loading doctors")
	for _, d := range list {
		if err := validate(d); err != nil {
			return nil, errors.Wrapf(err, "invalid doctor entry: %+v", d)
		}
		doctors.m[d.ID] = d
		doctors.k = append(doctors.k, d.ID)
		logLocal.Infof(" - %+v", doctors.m[d.ID])
	}

	return doctors, nil
}

func validate(d *ir.Doctor) error {
	v := reflect.ValueOf(*d)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := v.Field(i)
		if f.String() == "" {
			return fmt.Errorf("%s was empty; all fields of a Doctor must be present", t.Field(i).Name)
		}
	}
	return nil
}

// Add adds a doctor to the set of available doctors.
// Returns an error if the doctor with the same ID already exists.
func (d *Doctors) Add(doctor *ir.Doctor) error {
	if _, ok := d.m[doctor.ID]; ok {
		return fmt.Errorf("a consultant with this consultant ID %q already exists in the doctors map", doctor.ID)
	}
	d.m[doctor.ID] = doctor
	d.k = append(d.k, doctor.ID)
	return nil
}

// GetByID returns a doctor by ID or nil if no doctor is mapped to a given ID.
func (d *Doctors) GetByID(id string) *ir.Doctor {
	return d.m[id]
}

// GetByName returns a doctor by firstName and surname.
// Returns nil if no matching doctor found.
func (d *Doctors) GetByName(firstName string, surname string) *ir.Doctor {
	for _, doctor := range d.m {
		if doctor.FirstName == firstName && doctor.Surname == surname {
			return doctor
		}
	}
	return nil
}

// GetRandomDoctor returns a random doctor.
// Returns nil if no doctors are specified.
func (d *Doctors) GetRandomDoctor() *ir.Doctor {
	if len(d.k) == 0 {
		return nil
	}

	id := rand.Intn(len(d.k))
	return d.m[d.k[id]]
}
