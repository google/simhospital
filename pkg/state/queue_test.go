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

package state

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/state/persist"
	"github.com/google/simhospital/pkg/test/teststate"
)

func TestWrappedQueue_Put(t *testing.T) {
	tests := []struct {
		name        string
		syncer      persist.ItemSyncer
		insertOrder []MarshallableQueueItem
	}{
		{
			name:        "no syncer, insert in order",
			syncer:      nil,
			insertOrder: []MarshallableQueueItem{teststate.Item1, teststate.Item2},
		},
		{
			name:        "no syncer, insert out of order",
			syncer:      nil,
			insertOrder: []MarshallableQueueItem{teststate.Item2, teststate.Item1},
		},
		{
			name:        "with syncer, insert in order",
			syncer:      teststate.NewItemSyncer(),
			insertOrder: []MarshallableQueueItem{teststate.Item2, teststate.Item1},
		},
		{
			name:        "with syncer, insert out of order",
			syncer:      teststate.NewItemSyncer(),
			insertOrder: []MarshallableQueueItem{teststate.Item2, teststate.Item1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wq, err := NewWrappedQueue(teststate.Type, tt.syncer)
			if err != nil {
				t.Fatalf("NewWrappedQueue(%s, %v) failed with %v", teststate.Type, tt.syncer, err)
			}
			wq.Put(tt.insertOrder...)

			for _, want := range []MarshallableQueueItem{teststate.Item1, teststate.Item2} {
				// Confirm item added in order to priority queue.
				if got, _ := wq.q.Get(1); !cmp.Equal(want, got[0]) {
					t.Errorf("wq.q.Get(1) = %v, want: %v", got, want)
				}
				// Confirm item added to internal map.
				wantID, err := want.ID()
				if err != nil {
					t.Fatalf("want.ID() failed with %v", err)
				}
				if got := wq.m[wantID]; !cmp.Equal(want, got) {
					t.Errorf("wq.m[%q] = %v, want: %v", wantID, got, want)
				}
				// Confirm item added to synced store.
				if wq.syncer != nil {
					if got, _ := wq.syncer.LoadByID(wantID); !cmp.Equal(want, got) {
						t.Errorf("syncer LoadByID(%q) = %v, want: %v", wantID, got, want)
					}
				}
			}
		})
	}
}

func TestWrappedQueue_Get(t *testing.T) {
	tests := []struct {
		name   string
		syncer persist.ItemSyncer
	}{
		{name: "no syncer", syncer: nil},
		{name: "with syncer", syncer: teststate.NewItemSyncerWithDelete(true)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wq, err := NewWrappedQueue(teststate.Type, tt.syncer)
			if err != nil {
				t.Fatalf("NewWrappedQueue(%s, %v) failed with %v", teststate.Type, tt.syncer, err)
			}
			wq.Put(teststate.Item1, teststate.Item2)

			l := wq.Len()
			for _, wantItem := range []MarshallableQueueItem{teststate.Item1, teststate.Item2} {
				if got, _ := wq.Get(); !cmp.Equal(wantItem, *got) {
					t.Errorf("wq.Get() = %v, want: %v", *got, wantItem)
				}
				l--

				// Confirm item is removed from priority queue.
				if got, want := wq.Len(), l; got != want {
					t.Errorf("wq.Len() = %d, want: %d", got, want)
				}
				// Confirm item is removed from internal map.
				wantItemID, err := wantItem.ID()
				if err != nil {
					t.Fatalf("want.ID() failed with %v", err)
				}
				if got := wq.m[wantItemID]; got != nil {
					t.Errorf("wq.m[%q] = %v, want: nil", wantItemID, got)
				}
				// Confirm item removed from synced store.
				if wq.syncer != nil {
					if got, _ := wq.syncer.LoadByID(wantItemID); got != nil {
						t.Errorf("syncer LoadByID(%q) = %v, want: nil", wantItemID, got)
					}
				}
			}
		})
	}
}

func TestWrappedQueue_LoadFromSyncer(t *testing.T) {
	empty := teststate.NewItemSyncer()
	nonEmpty := teststate.NewItemSyncer()
	nonEmpty.Write(teststate.Item1)
	nonEmpty.Write(teststate.Item2)

	tests := []struct {
		name   string
		syncer persist.ItemSyncer
		want   []MarshallableQueueItem
	}{
		{name: "empty synced store", syncer: empty, want: []MarshallableQueueItem{}},
		{name: "non-empty synced store", syncer: nonEmpty, want: []MarshallableQueueItem{teststate.Item1, teststate.Item2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wq, err := NewWrappedQueue(teststate.Type, tt.syncer)
			if err != nil {
				t.Fatalf("NewWrappedQueue(%s, %v) failed with %v", teststate.Type, tt.syncer, err)
			}
			if got, want := wq.Len(), len(tt.want); got != want {
				t.Fatalf("wq.Len() = %d, want: %d", got, want)
			}
			for _, want := range tt.want {
				// Confirm item added in order to priority queue.
				if got, _ := wq.q.Get(1); !cmp.Equal(want, got[0]) {
					t.Errorf("wq.q.Get(1) = %v, want: %v", got, want)
				}
				// Confirm item added to internal map.
				wantID, err := want.ID()
				if err != nil {
					t.Fatalf("want.ID() failed with %v", err)
				}
				if got := wq.m[wantID]; !cmp.Equal(want, got) {
					t.Errorf("wq.m[%q] = %v, want: %v", wantID, got, want)
				}
			}
		})
	}
}

func TestWrappedQueue_LoadFromSyncer_NoSyncer(t *testing.T) {
	wq, err := NewWrappedQueue(teststate.Type, nil)
	if err != nil {
		t.Fatalf("NewWrappedQueue(%s, nil) failed with %v", teststate.Type, err)
	}
	err = wq.LoadFromSyncer()
	if err == nil || !strings.Contains(err.Error(), "ItemSyncer not set") {
		t.Errorf("wq.LoadFromSyncer() = %v, want: error complaining about ItemSyncer not set", err)
	}
}

func TestWrappedQueue_Peek(t *testing.T) {
	tests := []struct {
		name        string
		insertOrder []MarshallableQueueItem
	}{
		{
			name:        "insert in order",
			insertOrder: []MarshallableQueueItem{teststate.Item1, teststate.Item2},
		},
		{
			name:        "insert out of order",
			insertOrder: []MarshallableQueueItem{teststate.Item2, teststate.Item1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wq, err := NewWrappedQueue(teststate.Type, nil)
			if err != nil {
				t.Fatalf("NewWrappedQueue(%s, nil) failed with %v", teststate.Type, err)
			}
			wq.Put(tt.insertOrder...)

			if got := wq.Peek(); !cmp.Equal(got, teststate.Item1) {
				t.Errorf("wq.Peek() = %v, want: %v", got, teststate.Item1)
			}
		})
	}
}

func TestWrappedQueue_Len(t *testing.T) {
	wq, err := NewWrappedQueue(teststate.Type, nil)
	if err != nil {
		t.Fatalf("NewWrappedQueue(%s, nil) failed with %v", teststate.Type, err)
	}

	wq.Put(teststate.Item1)
	if got, want := wq.Len(), 1; got != want {
		t.Errorf("wq.Len() = %d, want: %d", got, want)
	}
	wq.Put(teststate.Item2)
	if got, want := wq.Len(), 2; got != want {
		t.Errorf("wq.Len() = %d, want: %d", got, want)
	}
	wq.Peek()
	if got, want := wq.Len(), 2; got != want {
		t.Errorf("wq.Len() = %d, want: %d", got, want)
	}
	wq.Get()
	if got, want := wq.Len(), 1; got != want {
		t.Errorf("wq.Len() = %d, want: %d", got, want)
	}
	wq.Get()
	if got, want := wq.Empty(), true; got != want {
		t.Errorf("wq.Empty() = %t, want: %t", got, want)
	}
}

func TestWrappedQueue_Empty(t *testing.T) {
	wq, err := NewWrappedQueue(teststate.Type, nil)
	if err != nil {
		t.Fatalf("NewWrappedQueue(%s, nil) failed with %v", teststate.Type, err)
	}

	wq.Put(teststate.Item1)
	if got, want := wq.Empty(), false; got != want {
		t.Errorf("wq.Empty() = %t, want: %t", got, want)
	}
	wq.Get()
	if got, want := wq.Empty(), true; got != want {
		t.Errorf("wq.Empty() = %t, want: %t", got, want)
	}
}

func TestWrappedQueue_IsConsistent_DuplicateIDs(t *testing.T) {
	wq, err := NewWrappedQueue(teststate.Type, nil)
	if err != nil {
		t.Fatalf("NewWrappedQueue(%s, nil) failed with %v", teststate.Type, err)
	}
	wq.Put(teststate.Item1)
	if got, want := wq.IsConsistent(), true; got != want {
		t.Errorf("wq.IsConsistent() = %t, want: %t", got, want)
	}

	// Add item with same id as before.
	testItemID1, err := teststate.Item1.ID()
	if err != nil {
		t.Fatalf("teststate.Item1.ID() failed with %v", err)
	}
	wq.Put(teststate.NewItem(testItemID1))
	if got, want := wq.IsConsistent(), false; got != want {
		t.Errorf("wq.IsConsistent() = %t, want: %t", got, want)
	}
}

func TestWrappedQueue_IsConsistent_NonexistentItem(t *testing.T) {
	wq, err := NewWrappedQueue(teststate.Type, nil)
	if err != nil {
		t.Fatalf("NewWrappedQueue(%s, nil) failed with %v", teststate.Type, err)
	}
	wq.Put(teststate.Item1)
	if got, want := wq.IsConsistent(), true; got != want {
		t.Errorf("wq.IsConsistent() = %t, want: %t", got, want)
	}

	testItemID1, err := teststate.Item1.ID()
	if err != nil {
		t.Fatalf("teststate.Item1.ID() failed with %v", err)
	}
	delete(wq.m, testItemID1) // Delete item from internal map only.
	wq.Put(teststate.Item2)   // Trigger consistency check.
	if got, want := wq.IsConsistent(), false; got != want {
		t.Errorf("wq.IsConsistent() = %t, want: %t", got, want)
	}
}
