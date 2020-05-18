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

// Package testtext contains functionality to generate deterministic text for testing.
package testtext

// Generator generates text for testing.
type Generator struct {
	Text []string
}

// Sentences returns n sentences.
// The sentences are picked from g.Text, in order.
func (g Generator) Sentences(n int) []string {
	var result []string
	l := len(g.Text)
	for i := 0; i < n; i++ {
		result = append(result, g.Text[i%l])
	}
	return result
}
