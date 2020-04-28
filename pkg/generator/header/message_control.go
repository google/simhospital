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

package header

import (
	"strconv"
	"sync"
)

// MessageControlGenerator is a generator of  Message Control IDs.
type MessageControlGenerator struct {
	mu     sync.Mutex
	nextID uint64
}

// NewMessageControlID returns a unique Message Control ID to use in the MSH.10 Message Control ID field (type ST).
// The IDs are incremental integer numbers starting with 1 - this is an easy way of generating unique values.
// NewMessageControlID is safe for concurrent use.
func (g *MessageControlGenerator) NewMessageControlID() string {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.nextID++
	return strconv.FormatUint(g.nextID, 10)
}
