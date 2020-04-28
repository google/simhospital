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
	"sync"

	"github.com/golang-collections/go-datastructures/queue"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/google/simhospital/pkg/state/persist"
)

// ErrSyncerNotSet is returned if no ItemSyncer is set, but the load operation was called.
var ErrSyncerNotSet = errors.New("ItemSyncer not set; cannot load from the syncer")

// WrappedQueue is a queue that contains both a PriorityQueue and an internal struct with objects that are marshallable
// (for persisting).
type WrappedQueue struct {
	q        *queue.PriorityQueue
	m        map[string]MarshallableQueueItem
	mutex    sync.Mutex
	itemType string
	// syncer receives all operations performed on the WrappedQueue.
	// If set, an ItemSyncer persists the items in the WrappedQueue in some kind of storage
	// (e.g., a DB), so that the persisted items are the same as the items inside the queue.
	// That way, if the execution stops for any reason, the items can be loaded from such storage.
	// The WrappedQueue forwards item additions and deletions to the syncer, so that when an item is
	// added or removed from the WrappedQueue, the syncer also adds or deletes the item in the storage.
	// Before forwarding Write and Delete requests, it checks whether the syncer has been set.
	syncer persist.ItemSyncer
	// consistent keeps track of whether the queue is, or has been, in an inconsistent state.
	// If the queue is not consistent, we print warnings but we don't stop the execution.
	// Once the queue is an inconsistent state, it is never marked as consistent again. We use this
	// attribute to decide which warnings to print as there's no point in printing unhelpful warnings.
	// Callers that want to detect changes in the inconsistency can use the IsConsistent method.
	consistent bool
}

// NewWrappedQueue returns a priority-queue based store for tracking the state of various objects inside Simulated
// Hospital. It accepts an itemType specifying what type of stateful object will be stored in the queue and an
// ItemSyncer for persisting state. Pass in a nil ItemSyncer if you only wish to use the in-memory priority queue
// for storing  state.
// If ItemSyncer is specified, it also loads all items from the syncer.
// If loading all items from syncer fails, it returns an error, but also returns a valid WrappedQueue.
// The user can decide whether loading the items on the initialization is critical for their use case.
func NewWrappedQueue(itemType string, syncer persist.ItemSyncer) (*WrappedQueue, error) {
	w := &WrappedQueue{
		q:          queue.NewPriorityQueue(100),
		m:          map[string]MarshallableQueueItem{},
		mutex:      sync.Mutex{},
		consistent: true,
		itemType:   itemType,
		syncer:     syncer,
	}

	if w.syncer != nil {
		if err := w.LoadFromSyncer(); err != nil {
			return w, errors.Wrap(err, "cannot load from syncer")
		}
	}

	return w, nil
}

// MarshallableQueueItem is the interface for items that can be stored in a MarshallableQueue.
type MarshallableQueueItem interface {
	persist.MarshallableItem
	Compare(queue.Item) int
}

// LoadFromSyncer loads all items from the syncer into the queue.
func (q *WrappedQueue) LoadFromSyncer() error {
	if q.syncer == nil {
		return ErrSyncerNotSet
	}
	items, err := q.syncer.LoadAll()
	if err != nil {
		return err
	}
	queueItems := make([]MarshallableQueueItem, len(items))
	for i, item := range items {
		queueItems[i] = item.(MarshallableQueueItem)
	}
	// Send a "nil" syncer so that the items aren't persisted: they have been persisted already.
	return q.put(nil, queueItems...)
}

// Len returns the number of items in the queue.
func (q *WrappedQueue) Len() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return len(q.m)
}

// Put inserts all arguments into the queue.
func (q *WrappedQueue) Put(items ...MarshallableQueueItem) error {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.put(q.syncer, items...)
}

func (q *WrappedQueue) put(syncer persist.ItemSyncer, items ...MarshallableQueueItem) error {
	for _, item := range items {
		if syncer != nil {
			syncer.Write(item)
		}
		err := q.q.Put(item)
		if err != nil {
			return err
		}
		q.addToInternalMap(item)
		if q.q.Len() != len(q.m) && q.consistent {
			log.Warningf("Queues out of sync after Put method: #priority queue: %d, #wrapped map: %d", q.q.Len(), len(q.m))
			q.consistent = false
		} else {
			counters.SimulatedHospital.PendingItem.With(prometheus.Labels{
				"item_type": q.itemType,
			}).Inc()
		}
	}
	return nil
}

func (q *WrappedQueue) addToInternalMap(i MarshallableQueueItem) error {
	s, err := i.ID()
	if err != nil {
		return errors.Wrap(err, "cannot get item ID")
	}
	logLocal := log.WithField("item_id", s)
	logLocal.Debugf("Adding item: %v", i)
	_, ok := q.m[s]
	if ok {
		logLocal.Warning("Key collision, elements can be lost")
	}
	q.m[s] = i
	return nil
}

// Get retrieves the next item from the queue, removing it from all internal data structures and the syncer.
func (q *WrappedQueue) Get() (*MarshallableQueueItem, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	i, err := q.q.Get(1)
	if err != nil {
		return nil, err
	}
	if len(i) != 1 {
		return nil, errors.Wrapf(err, "Failed to consume item in queue, expected 1, got %d", len(i))
	}
	item := i[0].(MarshallableQueueItem)
	if q.syncer != nil {
		q.syncer.Delete(item)
	}
	id, err := item.ID()
	if err != nil {
		return nil, errors.Wrap(err, "cannot get item ID")
	}
	_, ok := q.m[id]
	if !ok {
		log.WithField("item_id", id).Warning("Elements out of sync: asked to remove an item that wasn't present")
	} else {
		delete(q.m, id)
		counters.SimulatedHospital.PendingItem.With(prometheus.Labels{
			"item_type": q.itemType,
		}).Dec()
	}
	if q.q.Len() != len(q.m) && q.consistent {
		log.Warningf("Elements out of sync after GET method: #priority queue: %d, #wrapped map: %d", q.q.Len(), len(q.m))
		q.consistent = false
	}
	return &item, nil
}

// Peek returns the next item in the queue without removing it from the queue.
func (q *WrappedQueue) Peek() queue.Item {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.q.Peek()
}

// Empty returns whether the queue is empty.
func (q *WrappedQueue) Empty() bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.q.Empty()
}

// IsConsistent returns whether the queue and mapping of MarshallableQueueItems have ever fallen out of sync.
func (q *WrappedQueue) IsConsistent() bool {
	return q.consistent
}
