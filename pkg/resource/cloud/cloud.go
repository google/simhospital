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

// Package cloud contains functionality to write to a Cloud FHIR store.
package cloud

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
	"google.golang.org/api/healthcare/v1"
	"google.golang.org/api/option"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/logging"
)

var log = logging.ForCallerPackage()

const parentFormat = "projects/%s/locations/%s/datasets/%s/fhirStores/%s"

// Output is a resource.Output that returns writers to a Cloud FHIR store.
type Output struct {
	svc    *healthcare.ProjectsLocationsDatasetsFhirStoresFhirService
	parent string
}

// New creates a new cloud WriterCloser.
func (o *Output) New(*ir.PatientInfo) (io.WriteCloser, error) {
	return &WriterCloser{svc: o.svc, parent: o.parent}, nil
}

// NewOutput creates a new healthcare service and returns an Ouput. If invalid details are given,
// no error will be returned until an actual attempt to write is made (in WriteCloser.Write()).
func NewOutput(ctx context.Context, projectID, location, dataset, datastore string, opts ...option.ClientOption) (*Output, error) {
	if err := checkEmptyArg("projectID", projectID); err != nil {
		return nil, err
	}
	if err := checkEmptyArg("location", location); err != nil {
		return nil, err
	}
	if err := checkEmptyArg("dataset", dataset); err != nil {
		return nil, err
	}
	if err := checkEmptyArg("datastore", datastore); err != nil {
		return nil, err
	}
	healthcareService, err := healthcare.NewService(ctx, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "could not create healthcare service")
	}
	return &Output{
		svc:    healthcareService.Projects.Locations.Datasets.FhirStores.Fhir,
		parent: fmt.Sprintf(parentFormat, projectID, location, dataset, datastore),
	}, nil
}

// WriterCloser is an io.WriterCloser that writes to a Cloud FHIR store.
type WriterCloser struct {
	svc    *healthcare.ProjectsLocationsDatasetsFhirStoresFhirService
	parent string
}

// Write writes to the specified Cloud FHIR store. All consecutive writes will persist, regardless
// of whether Close is called or not.
func (c *WriterCloser) Write(b []byte) (int, error) {
	call := c.svc.ExecuteBundle(c.parent, bytes.NewReader(b))
	call.Header().Set("Content-Type", "application/fhir+json;charset=utf-8")

	resp, err := call.Do()
	if err != nil {
		return 0, errors.Wrap(err, "could not make call")
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return 0, errors.Wrap(err, "could not read response")
	}

	// The Cloud FHIR store usually returns 201 for creation of resources and possibly others for
	// different actions, so we accept all 2xx codes.
	if resp.StatusCode > 299 {
		return 0, errors.Errorf("response returned status %s: %s", resp.Status, respBytes)
	}

	log.Infof("Response <%d>: %s", resp.StatusCode, respBytes)
	return int(resp.Request.ContentLength), nil
}

// Close is a no-op.
func (c *WriterCloser) Close() error {
	return nil
}

func checkEmptyArg(arg string, s string) error {
	if s == "" {
		return errors.Errorf("%s unspecified, this is required", arg)
	}
	return nil
}
