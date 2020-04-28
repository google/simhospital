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

// Package teststate provides test-only implementation for state-specific interfaces.
package teststate

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/golang-collections/go-datastructures/queue"
	"github.com/pkg/errors"
	"github.com/google/simhospital/pkg/state/persist"
)

const (
	// Type is the type of an Item.
	Type = "test-type"
)

var (
	// Item1 is an Item with ID 1.
	Item1 = NewItem("1")
	// Item2 is an Item with ID 2.
	Item2 = NewItem("2")
)

// Item implements persist.MarshallableItem and MarshallableQueueItem interfaces for testing.
type Item struct {
	// I is the item's ID.
	I string
	// S represents whether the Marshal operation on this object will be successful.
	S bool
}

// ID returns the item's ID.
func (i Item) ID() (string, error) {
	return i.I, nil
}

// Marshal marshals item for persisting.
func (i Item) Marshal() ([]byte, error) {
	if i.S {
		return json.Marshal(i)
	}
	return nil, errors.New("unsuccessful marshal")
}

// Compare compares two items in a priority queue.
func (i Item) Compare(other queue.Item) int {
	o, ok := other.(Item)
	if !ok {
		panic(fmt.Sprintf("Incompatible type: trying to compare Item with %v", other))
	}
	if i.I > o.I {
		return 1
	}
	return -1
}

// NewItem creates a Item with success status.
func NewItem(id string) Item {
	return Item{
		I: id,
		S: true,
	}
}

// ItemSyncer implements the persist.ItemSyncer interface using a map.
// It tracks the LoadByID requests made to the syncer in the form of an
// internal `reqs` map for testing purposes.
type ItemSyncer struct {
	m    map[string]persist.MarshallableItem
	reqs map[string]int
	// delete indicates whether to delete items or not.
	delete   bool
	nDeletes int
}

// NewItemSyncer initializes the ItemSyncer.
func NewItemSyncer() *ItemSyncer {
	return &ItemSyncer{m: map[string]persist.MarshallableItem{}, reqs: map[string]int{}}
}

// NewItemSyncerWithDelete initializes the ItemSyncer with a value for whether to delete items or not.
func NewItemSyncerWithDelete(delete bool) *ItemSyncer {
	return &ItemSyncer{m: map[string]persist.MarshallableItem{}, reqs: map[string]int{}, delete: delete}
}

// Write writes an item to the map.
func (s *ItemSyncer) Write(item persist.MarshallableItem) error {
	id, err := item.ID()
	if err != nil {
		return errors.Wrap(err, "cannot get ID")
	}
	s.m[id] = item
	return nil
}

// Delete deletes an item from the map.
func (s *ItemSyncer) Delete(item persist.MarshallableItem) error {
	if !s.delete {
		return nil
	}
	s.nDeletes++
	id, err := item.ID()
	if err != nil {
		return errors.Wrap(err, "cannot get ID")
	}
	delete(s.m, id)
	return nil
}

// LoadAll returns a slice of all the items in the map, sorted by id.
func (s *ItemSyncer) LoadAll() ([]persist.MarshallableItem, error) {
	keys := make([]string, 0)
	for id := range s.m {
		keys = append(keys, id)
	}
	sort.Strings(keys)

	sorted := make([]persist.MarshallableItem, len(keys))
	for i, k := range keys {
		sorted[i] = s.m[k]
	}
	return sorted, nil
}

// LoadByID returns an item in the map with the provided id, if it exists.
func (s *ItemSyncer) LoadByID(id string) (persist.MarshallableItem, error) {
	s.reqs[id]++ // track LoadByID calls made to ItemSyncer for testing.
	return s.m[id], nil
}

// Count returns number of elements in the syncer for testing.
func (s *ItemSyncer) Count() int {
	return len(s.m)
}

// CountDeletes returns the number of deletions requested.
func (s *ItemSyncer) CountDeletes() int {
	return s.nDeletes
}

// WasRequested returns whether LoadByID was ever called with a specific id.
func (s *ItemSyncer) WasRequested(id string) bool {
	return s.reqs[id] > 0
}
