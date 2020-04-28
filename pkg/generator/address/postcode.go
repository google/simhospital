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

package address

import (
	"fmt"
	"math/rand"
)

// UKPostcode is a generator of UK postcodes.
type UKPostcode struct{}

// Random returns a random string that matches the format of a UK post code:
// XX1 1XX or XX11 1XX
//
// Where:
// X is a random letter
// 1 is a random number between [1, 9]
//
// The returned postcode might exist or not.
func (g *UKPostcode) Random() string {
	return fmt.Sprintf("%s%s%d %d%s%s", randomLetter(), randomLetter(), rand.Intn(99)+1, rand.Intn(9)+1, randomLetter(), randomLetter())
}

func randomLetter() string {
	return string(rand.Intn(int('Z')-int('A')) + int('A'))
}

// USPostcode is a generator of US zipcodes.
type USPostcode struct{}

// Random returns a random string that matches the format of a US zipcode:
// 11111
//
// Where:
// 1 is a random number between [1, 9]
//
// The returned zipcode might exist or not.
func (g *USPostcode) Random() string {
	chars := []rune("1234567890")
	udn := make([]rune, 5)
	for i := range udn {
		udn[i] = chars[rand.Intn(len(chars))]
	}
	return string(udn)
}
