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
	"fmt"
	"sort"

	"github.com/pkg/errors"
)

// Collection represents a collection of pathways.
type Collection struct {
	// pathways is a map of pathways indexed by their name.
	pathways map[string]Pathway
	// pathwayNames is a sorted list of the names of all pathways.
	pathwayNames []string
}

// GetPathway gets the pathway with the given name.
// If the name provided is not valid or does not exist, it returns an error.
func (c Collection) GetPathway(pathwayName string) (*Pathway, error) {
	pathway, ok := c.pathways[pathwayName]
	if !ok {
		return nil, fmt.Errorf("pathwayName %s does not exist within Collection", pathwayName)
	}
	runnable, err := pathway.Runnable()
	if err != nil {
		return nil, errors.Wrapf(err, "pathway with name %s is not runnable", pathwayName)
	}
	return &runnable, nil
}

// PathwayNames returns the names of all the pathways in this Collection, alphabetically sorted.
func (c Collection) PathwayNames() []string {
	return c.pathwayNames
}

// Pathways returns the pathways in this Collection.
func (c Collection) Pathways() map[string]Pathway {
	return c.pathways
}

// NewCollection creates a new Collection with the given pathway map.
// All pathways are initialised.
func NewCollection(pathways map[string]Pathway) (Collection, error) {
	p := Collection{
		pathways: map[string]Pathway{},
	}
	for k, v := range pathways {
		v.Init(k)
		p.pathways[k] = v
		p.pathwayNames = append(p.pathwayNames, k)
	}
	sort.Strings(p.pathwayNames)
	return p, nil
}

// Print prints the names of the pathways in the collection, with optional suffixes.
func (c Collection) Print(suffixes map[string]string) {
	log.Infof("Loaded %d pathways:", len(c.pathwayNames))
	for _, name := range c.pathwayNames {
		log.Infof(" - %s%s", name, suffixes[name])
	}
}
