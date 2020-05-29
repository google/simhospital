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

// Package testaddress contains functionality to generate deterministic addresses for testing.
package testaddress

import (
	"fmt"

	"github.com/google/simhospital/pkg/ir"
)

// ArbitraryGenerator is a generator of addresses with predefined country and cities,
// to be used when the exact values do not matter.
var ArbitraryGenerator = Generator{Country: "GBR", Cities: []string{"Cambridge", "London"}}

// Generator generates addresses for testing.
type Generator struct {
	n       int
	Country string
	Cities  []string
}

// Random returns a deterministically different British address each time. The address will be in the following formats with equal probabilities:
// 		n Test House
//		n Test Street
// 		AB# #CD
//		City
//		Country
//
// Where:
// n is a integer number that starts from 1 and increases by 1 for every address generated.
//
// AB# #CD is a UK post code, where # is a digit that ranges between 0-9, increasing by 1 for every
// address generated and wrapping to 0. It is equivalent to the ones digit of n.
//
// City is the nth city in the list of cities configured for the generator. After it reaches the end
// of the list, it will wrap back around to the beginning of the list.
func (g *Generator) Random() *ir.Address {
	g.n++

	return &ir.Address{
		FirstLine:  fmt.Sprintf("%d Test House", g.n),
		SecondLine: fmt.Sprintf("%d Test Street", g.n),
		City:       g.city(),
		PostalCode: g.postCode(),
		Country:    g.Country,
		Type:       "HOME",
	}
}

func (g Generator) postCode() string {
	return fmt.Sprintf("AB%d %dCD", g.n%10, g.n%10)
}

func (g Generator) city() string {
	return g.Cities[g.n%len(g.Cities)]
}
