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

// Package text contains functions needed to generate random sentences.
package text

import (
	"math/rand"
	"strings"
)

// Generator is a generator of text.
type Generator interface {
	// Sentences generates n sentences.
	Sentences(int) []string
}

// maxWords is the maximum number of words generated in a sentence.
const maxWords = 10

// NounGenerator generates text by concatenating nouns.
type NounGenerator struct {
	Nouns []string
}

// randomSentence returns a random sentence consisting of between [1, max] random nouns,
// separated by an empty space.
// The first word starts with a capital letter.
func (g *NounGenerator) randomSentence(max int) string {
	r := rand.Intn(max)
	w := make([]string, 0)
	for i := 0; i <= r; i++ {
		w = append(w, g.randomNoun())
	}
	w[0] = strings.Title(w[0])
	return strings.Join(w, " ")
}

func (g *NounGenerator) randomNoun() string {
	return g.Nouns[rand.Intn(len(g.Nouns))]
}

// Sentences returns an array of n random sentences.
func (g *NounGenerator) Sentences(n int) []string {
	contentLine := make([]string, 0, n)
	for i := 0; i < n; i++ {
		contentLine = append(contentLine, g.randomSentence(maxWords))
	}
	return contentLine
}
