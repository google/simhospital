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

// Package sample includes the functionality for sampling from a discrete distribution.
package sample

import (
	"math/rand"
)

// WeightedValue represents the value and its frequency.
type WeightedValue struct {
	Value     interface{}
	Frequency uint
}

// DiscreteDistribution represents a collection of weighted values that form a distribution.
type DiscreteDistribution struct {
	WeightedValues []WeightedValue
}

func (d DiscreteDistribution) total() uint {
	total := uint(0)
	for _, result := range d.WeightedValues {
		total += result.Frequency
	}
	return total
}

// randUint returns, as an uint, a pseudo-random number in [0,n) from the default Source.
func randUint(n uint) uint {
	return uint(rand.Intn(int(n)))
}

// Random samples the DiscreteDistribution and returns the resulting value.
func (d DiscreteDistribution) Random() interface{} {
	if d.total() == 0 {
		return nil
	}
	r := randUint(d.total())
	current := uint(0)
	for _, result := range d.WeightedValues {
		current += result.Frequency
		if r < current {
			return result.Value
		}
	}

	return nil
}
