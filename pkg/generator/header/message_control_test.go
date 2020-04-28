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
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewMessageControlIDMessageControlIDsAreIncremental(t *testing.T) {
	g := &MessageControlGenerator{}

	want := []string{"1", "2", "3"}
	var got []string
	for range want {
		got = append(got, g.NewMessageControlID())
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("NewMessageControlID() got diff (-want, +got):\n%s", diff)
	}
}

func TestNewMessageControlIDIsThreadSafe(t *testing.T) {
	g := &MessageControlGenerator{}

	want := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	gotChan := make(chan string)
	defer close(gotChan)
	for range want {
		go func() {
			got := g.NewMessageControlID()
			gotChan <- got
		}()
	}

	got := make([]string, len(want))
	for i := range want {
		g := <-gotChan
		got[i] = g
	}
	sort.Strings(got)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("NewMessageControlID() got diff (-want, +got):\n%s", diff)
	}
}
