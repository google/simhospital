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

// Package constants contains constant variables used in pathway definition.
package constants

const (
	// RandomString is a keyword used to indicate, that the given value
	// should be generated randomly.
	RandomString = "RANDOM"
	// NormalValue is a keyword used to indicate, that the given value
	// should be randomly chosen from a normal range.
	NormalValue = "NORMAL"
	// AbnormalHigh is a keyword used to indicate, that the given value
	// should be set to a random abnormal high value, ie: it should be above
	// the high end of the normal range.
	AbnormalHigh = "ABNORMAL_HIGH"
	// AbnormalLow is a keyword used to indicate, that the given value
	// should be set to a random abnormal low value, ie: it should be below
	// the low end of the normal range.
	AbnormalLow = "ABNORMAL_LOW"
	// EmptyString is a keyword used to indicate, that the given value
	// should be left empty.
	EmptyString = "EMPTY"

	// MidnightDate is a keyword used to indicate, that the given date
	// should be set to midnight time.
	MidnightDate = "MIDNIGHT"

	// NumericalValueType indicates that the value is numerical.
	NumericalValueType = "NM"
	// TextualValueType indicates that the value is textual.
	TextualValueType = "TX"

	// R01 is the trigger event R01.
	R01 = "R01"
	// R03 is the trigger event R03.
	R03 = "R03"
	// R32 is the trigger event R32.
	R32 = "R32"
)
