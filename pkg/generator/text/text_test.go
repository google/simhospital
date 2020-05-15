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

package text

import (
	"math"
	"math/rand"
	"strings"
	"testing"
)

var nouns = []string{"one", "two", "three", "four", "five"}

func TestNounGenerator_Sentences_Length(t *testing.T) {
	g := NounGenerator{Nouns: nouns}
	if got, want := len(g.Sentences(10)), 10; got != want {
		t.Errorf("len(g.Sentences(%v) got %v, want %v", 10, got, want)
	}
}

func TestNounGenerator_Sentences_TextIsRandom(t *testing.T) {
	rand.Seed(1)
	g := &NounGenerator{Nouns: nouns}
	runs := float64(100)
	sentencesPerRun := 10
	sentenceLenDistr := map[int]int{}
	wordDistr := map[string]int{}
	wordCount := 0
	for i := 0; i < int(runs); i++ {
		sentences := g.Sentences(sentencesPerRun)
		if ls := len(sentences); ls != 10 {
			t.Errorf("len(sentences)=%d, want exactly 10. Sentences: %v", ls, sentences)
		}
		for _, s := range sentences {
			words := strings.Split(s, " ")
			lw := len(words)
			sentenceLenDistr[lw]++
			if lw > 10 {
				// This check is repeated later, but here we print the words which is useful to know what went wrong.
				t.Errorf("len(words)=%d, want no more than 10. Sentences: %v", lw, words)
			}
			for _, w := range words {
				wordDistr[strings.ToLower(w)]++
				wordCount++
			}
		}
	}

	// Check that every sentence has from 1 to 10 words with equal probability.

	wantNum := int(runs * float64(sentencesPerRun) / 10)
	wantSentenceLenDistr := map[int]int{
		1: wantNum, 2: wantNum, 3: wantNum, 4: wantNum, 5: wantNum, 6: wantNum, 7: wantNum, 8: wantNum, 9: wantNum, 10: wantNum,
	}
	// Allow an error of half of the expected probability (arbitrary).
	delta := float64(wantNum) * 0.5

	for k, v := range sentenceLenDistr {
		want, ok := wantSentenceLenDistr[k]
		if !ok {
			t.Errorf("len(sentence) got %v, want length to be in %v", k, wantSentenceLenDistr)
		}
		if math.Abs(float64(v)-float64(want)) >= delta {
			t.Errorf("sentenceLenDistr[%v] = %d, want within %v of %d", k, v, delta, want)
		}
	}

	// Check that all words are picked with equal probability.
	wantP := int(wordCount / 5)
	wantWordDistr := map[string]int{
		"one": wantP, "two": wantP, "three": wantP, "four": wantP, "five": wantP,
	}
	// Allow an error of half of the expected probability (arbitrary).
	delta = float64(wantP) * 0.5
	for k, v := range wordDistr {
		want, ok := wantWordDistr[k]
		if !ok {
			t.Errorf("word in generated note is %q, want word to be in %v", k, wantWordDistr)
		}
		if math.Abs(float64(v)-float64(want)) >= delta {
			t.Errorf("wordDistr[%v] = %d, want within %v of %d", k, v, delta, want)
		}
	}
}
