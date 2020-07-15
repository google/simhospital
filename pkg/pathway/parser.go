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

package pathway

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"github.com/google/simhospital/pkg/clock"
	"github.com/google/simhospital/pkg/doctor"
	"github.com/google/simhospital/pkg/files"
	"github.com/google/simhospital/pkg/location"
	"github.com/google/simhospital/pkg/orderprofile"
)

// UnknownPathwayName is the default pathway name, if it is not explicitly specified.
const UnknownPathwayName = "unknown_pathway"

// validExtensions defines which file extensions are valid for pathways.
var validExtensions = []string{".json", ".yml", ".yaml"}

// Parser provides the functionality to parse the pathways.
type Parser struct {
	Clock clock.Clock
	// OrderProfiles is used to validate the results specified in the pathway.
	OrderProfiles *orderprofile.OrderProfiles
	// Doctors is used to validate doctors specified in the pathway.
	Doctors *doctor.Doctors
	// Valid function performs additional pathway validation.
	Valid func(*Pathway) error
	// LocationManager contains the patient locations.
	LocationManager *location.Manager
}

// ParsePathways parses all pathways defined in the pathwaysDir.
// Returns a map of pathway name to pathway structure.
// ParsePathways expects all pathways in the directory to be well formed and valid, and it will return an error
// if that is not the case. An error is also returned if the given regular expressions are invalid, or if there
// are no pathways defined in the directory.
// All pathways are initialised, but are not necessarily runnable yet. Ensure that Runnable() is called
// before the pathway is ran.
// Pathways can be specified in YAML or JSON.
func (p *Parser) ParsePathways(ctx context.Context, pathwaysDir string) (map[string]Pathway, error) {
	logLocal := log.WithField("pathway_dir", pathwaysDir)
	logLocal.Info("Parsing pathways from directory")
	files, err := files.List(ctx, pathwaysDir)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to read pathways files from %s", pathwaysDir)
	}

	validPathways := map[string]Pathway{}
	redeclaredPathways := make(map[string]bool, 0)
	for _, file := range files {
		if !fileExtensionIsValid(file.Name()) {
			log.Warnf("File name has invalid extension %s, expected one of %+v. Skipping...", file.Name(), validExtensions)
			continue
		}
		logLocal := logLocal.WithField("pathway_file", file.Name())
		logLocal.Info("Parsing pathways from file")
		p, err := p.parse(ctx, file)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to parse pathway file %s", file.Name())
		}

		for pathwayName, pathway := range p {
			logLocal := logLocal.WithField("pathway_name", pathwayName)
			if _, ok := validPathways[pathwayName]; ok {
				logLocal.Error("Pathway re-declared")
				redeclaredPathways[pathwayName] = true
				continue
			}

			logLocal.Debug("Adding pathway")
			validPathways[pathwayName] = pathway
		}
	}
	if len(validPathways) == 0 {
		return nil, fmt.Errorf("cannot load pathways from %s: no valid pathways", pathwaysDir)
	}

	if len(redeclaredPathways) > 0 {
		return nil, fmt.Errorf("cannot load pathways from %s: found re-declared pathways: %v", pathwaysDir, redeclaredPathways)
	}

	return validPathways, nil
}

// ParseSinglePathway parses the given pathway definition as a YAML or JSON format definition for a Pathway,
// or as a YAML or JSON format definition for a map with a single Pathway.
// In that second case, the returned pathway will have a name.
// The returned pathway is initialised and runnable.
func (p *Parser) ParseSinglePathway(pathwayDefinition []byte) (Pathway, error) {
	pathway := Pathway{}
	pathwayName := UnknownPathwayName
	// The yaml library parses JSON too, we don't need anything extra to support JSON.
	err := yaml.UnmarshalStrict(pathwayDefinition, &pathway)
	if err != nil {
		var m map[string]Pathway
		err := yaml.UnmarshalStrict(pathwayDefinition, &m)
		if err != nil {
			return Pathway{}, errors.Wrap(err, "cannot unmarshal pathway")
		}
		if len(m) != 1 {
			return Pathway{}, fmt.Errorf("too many pathways: found %d, expected 1", len(m))
		}
		for k, v := range m { // Access the only item in the map.
			pathway = v
			pathwayName = k
		}
	}
	pathway.Init(pathwayName)
	if err := pathway.Valid(p.Clock, p.OrderProfiles, p.Doctors, p.LocationManager, p.Valid); err != nil {
		return Pathway{}, errors.Wrap(err, "invalid pathway")
	}

	pathway, err = pathway.Runnable()
	if err != nil {
		return Pathway{}, errors.Wrap(err, "cannot run Runnable on pathway")
	}

	return pathway, nil
}

func (p *Parser) parse(ctx context.Context, file files.File) (map[string]Pathway, error) {
	pathways := map[string]Pathway{}

	data, err := file.Read(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "cannot parse pathways file")
	}

	// The yaml library parses JSON too, we don't need anything extra to support JSON.
	err = yaml.UnmarshalStrict(data, &pathways)
	if err != nil {
		return nil, errors.Wrap(err, "cannot unmarshal pathways")
	}

	invalidPathways := make([]string, 0)
	var allErrors []error
	for name, pathway := range pathways {
		pathway.Init(name)
		pathways[name] = pathway
		if err := pathway.Valid(p.Clock, p.OrderProfiles, p.Doctors, p.LocationManager, p.Valid); err != nil {
			log.WithField("pathway_file", file.FullPath()).WithField("pathway_name", name).
				WithError(err).Error("Invalid pathway")
			invalidPathways = append(invalidPathways, name)
			allErrors = append(allErrors, err)
		}
	}
	if len(invalidPathways) > 0 {
		return nil, fmt.Errorf("pathways %v are invalid: %v", invalidPathways, allErrors)
	}

	return pathways, nil
}

func fileExtensionIsValid(fileName string) bool {
	for _, ext := range validExtensions {
		if filepath.Ext(fileName) == ext {
			return true
		}
	}
	return false
}
