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

package orderprofile

import (
	"fmt"
	"math/rand"
	"regexp"

	"github.com/pkg/errors"
	"github.com/google/simhospital/pkg/constants"
)

const valueFormat = "%.2f"

var (
	fromToRangeRegExp = []*regexp.Regexp{
		// simple ranges,ie:
		// from-to
		// from - to
		// [ from - to ]
		// [from-to]
		regexp.MustCompile("^\\[? ?(-?[0-9|/.]+) ?- ?(-?[0-9|/.]+) ?\\]?$"),

		// duplicated value ranges, ie:
		// from-to^from^to
		regexp.MustCompile("^(-?[0-9|/.]+)-(-?[0-9|/.]+)(\\^-?[0-9|/.]+)+$"),
	}

	lessRangeRegExp = []*regexp.Regexp{
		// less than ranges, ie:
		// <to^^<to
		// <to^<to
		// <=to^^<=to
		// <=to^<=to
		regexp.MustCompile("^<=?(-?[0-9|/.]+)\\^+<=?-?[0-9|/.]+$"),

		// less than ranges with brackets, ie:
		// [ < to ]
		// [ <= to ]
		// [<to]
		// [<=to]
		regexp.MustCompile("^\\[? ?<=? ?(-?[0-9|/.]+) ?\\]?$"),
	}
	greaterRangeRegExp = []*regexp.Regexp{
		// greater than ranges, ie:
		// >from^^>from
		// >from^>from
		// >=from^^>=from
		// >=from^>=from
		regexp.MustCompile("^>=?(-?[0-9|/.]+)\\^+>=?-?[0-9|/.]+$"),

		// greater than ranges with brackets, ie:
		// [ > from ]
		// [ >= from ]
		// [>from]
		// [>=from]
		regexp.MustCompile("^\\[? ?>=? ?(-?[0-9|/.]+) ?\\]?$"),
	}
)

// ValueGenerator generates the value within the range (exclusive) given the normal value range.
//
// The ranges specified for the order profiles may sometimes be:
// inclusive for one border, eg.: >=5.5,
// exclusive, eg.: >0.25
// unknown inclusivity, eg.: 0.5 - 1.6 - probably exclusive, 70-120^70^120 - probably inclusive
// It is though safer (and easier) to always treat the range as exclusive.
type ValueGenerator struct {
	from validFloat
	to   validFloat
}

type validFloat struct {
	value float64
	valid bool
}

func newValidFloat(f float64) validFloat {
	return validFloat{value: f, valid: true}
}

func newInvalidFloat() validFloat {
	return validFloat{valid: false}
}

// IsHigh returns whether the value v is above the range represented by ValueGenerator.
func (g *ValueGenerator) IsHigh(v float64) bool {
	return g.to.valid && v > g.to.value
}

// IsLow returns whether the value v is below the range represented by ValueGenerator.
func (g *ValueGenerator) IsLow(v float64) bool {
	return g.from.valid && v < g.from.value
}

// IsNormal returns whether the value v is within range represented by ValueGenerator.
func (g *ValueGenerator) IsNormal(v float64) bool {
	if g.from.valid && g.to.valid {
		return v >= g.from.value && v <= g.to.value
	}
	return (g.from.valid && v >= g.from.value) || (g.to.valid && v <= g.to.value)
}

// Random returns the random value based on the randomType, which is either within normal ranges,
// or outside the normal ranges (ie: higher or lower).
// Returns error if the random value cannot be generated.
func (g *ValueGenerator) Random(randomType string) (string, error) {
	switch randomType {
	case constants.AbnormalLow:
		return g.AbnormalLow()
	case constants.AbnormalHigh:
		return g.AbnormalHigh()
	case constants.NormalValue:
		return g.Normal()
	default:
		log.WithField("random_type", randomType).Error("Unknown random type")
		return "", errors.New("unknown random type")
	}
}

// Normal returns random number formatted as string within the normal range,
// ie between (g.from, g.to) exclusive.
//
// If ValueGenerator represents a right open range, ie: >g.from (g.to is invalid):
// - if "from" is positive, the normal value is generated from (g.from, 10 x g.from).
// - if "from" is negative, the normal value is generated from (g.from, 0) to only allow negative numbers.
// - if "from" is 0, the normal value is generated from (g.from, 10); this is an arbitrary choice,
//   as we don't really know what order of magnitude the values should be in.
//
// If ValueGenerator represents a left open range, ie: <g.to (g.from is invalid):
// - if "to" is positive, the normal value is generated from (0, g.to) to only allow positive values.
// - if "to" is negative, the normal value is generated from (10 x g.to, g.to).
// - if "to" is 0, the normal value is generated from (-10, g.to); this is an arbitrary choice,
//   as we don't really know what order of magnitude the values should be in.
//
// If g == nil, returns 0.
// Returns error if both: start and end of the range are open.
func (g *ValueGenerator) Normal() (string, error) {
	if g == nil {
		return fmt.Sprintf(valueFormat, 0.0), nil
	}
	if !g.from.valid && !g.to.valid {
		log.WithField("value_generator", g).Error("Cannot generate normal value if both: start and end of the range are open")
		return "", errors.New("cannot generate normal value if both: start and end of the range are open")
	}

	var from, to float64
	if g.from.valid {
		from = g.from.value
	} else if g.to.value == 0 {
		// if end of the range is equal to 0, start of the range is set to -10; this is an arbitrary choice,
		// as we don't really know what order of magnitude the values should be in
		from = -10
	} else if g.to.value > 0 {
		// if end of range is positive, set "from" to 0 to only allow positive numbers
		from = 0
	} else {
		// if end of range is negative, set "from" to 10 x start of the range,
		// to avoid absurd small numbers being generated
		from = g.to.value * 10
	}
	if g.to.valid {
		to = g.to.value
	} else if from == 0 {
		// if start of the range is equal to 0, end of the range is set to 10; this is an arbitrary choice,
		// as we don't really know what order of magnitude the values should be in
		to = 10
	} else if from > 0 {
		// if the start of the range is positive, set end of the range to 10 x start of the range,
		// to avoid absurd huge numbers being generated
		to = from * 10
	} else {
		// if end of the range is negative, set end of the range to 0 to only allow negative numbers
		to = 0
	}

	return randomFromRange(from, to)
}

// AbnormalLow returns a random number formatted as string, which is lower than the normal range.
// If the start of normal range is positive, it will return value between (0, g.from) exclusive.
// If the start of normal range is negative, it will return value between (10 x g.from, g.from) exclusive.
// Returns an error in any of the following situations:
// - the receiver is nil
// - the start of the normal range is open
// - the start of the normal range is 0 -> the assumption is that if start of the normal range is positive,
//   the negative numbers are invalid, thus it is impossible to generate the abnormal low value if range starts at 0
func (g *ValueGenerator) AbnormalLow() (string, error) {
	if g == nil {
		return "", errors.New("cannot generate abnormal low value for nil ValueGenerator")
	}
	if !g.from.valid || g.from.value == 0 {
		log.WithField("value_generator", g).Error("Cannot generate abnormal low value for open or zero start range")
		return "", errors.New("cannot generate abnormal low value for open or zero start range")
	}
	var from, to float64
	if g.from.value > 0 {
		from, to = 0, g.from.value
	} else {
		from, to = 10*g.from.value, g.from.value
	}

	return randomFromRange(from, to)
}

// AbnormalHigh returns a random number formatted as string, which is higher than the normal range.
// If the end of normal range is positive, it will return value between (g.to, 10 x g.to) exclusive.
// If the end of normal range is negative, it will return value between (g.to, 0) exclusive.
// Returns an error in any of the following situations:
// - the receiver is nil
// - the end of the normal range is open
// - the end of the normal range is 0 -> the assumption is that if the end of the normal range is negative,
//   the positive numbers are invalid, thus it is impossible to generate the abnormal high value if range ends at 0
func (g *ValueGenerator) AbnormalHigh() (string, error) {
	if g == nil {
		return "", errors.New("cannot generate abnormal high value for nil ValueGenerator")
	}
	if !g.to.valid || g.to.value == 0 {
		log.WithField("value_generator", g).Error("Cannot generate abnormal high value for open or zero end range")
		return "", errors.New("cannot generate abnormal high value for open or zero end range")
	}
	var from, to float64
	if g.to.value > 0 {
		from, to = g.to.value, 10*g.to.value
	} else {
		from, to = g.to.value, 0
	}

	return randomFromRange(from, to)
}

func randomFromRange(from float64, to float64) (string, error) {
	for i := 0; i < 100; i++ {
		f := rand.Float64()*(to-from) + from

		// The rand.Float64() returns value between [0.0, 1.0), ie the start of the range is inclusive, while
		// the number generated by the ValueGenerator needs to be exclusive.
		// Furthermore, the value is formatted as a string, when it loses some precisions, ie. it is formatted
		// to 2 decimal places only.
		// We need to ensure, that the generated value is still within the range (exclusive for both: start and end
		// of the range) after the value is formatted.
		s := fmt.Sprintf(valueFormat, f)
		if afterFormatting, _ := floatFromString(s); afterFormatting > from && afterFormatting < to {
			return s, nil
		}
	}

	log.WithField("from", from).WithField("to", to).Error("Failed to generate the value from range after 100 attempts")
	return "", errors.New("failed to generate the value from range after 100 attempts")
}

// ValueGeneratorFromRange returns ValueGenerator created from string
// containing the normal value range.
// String may be in one of the following formats:
//
// from-to
// from - to
// [ from - to ]
// [from-to]
// from-to^from^to
// where both: "from" and "to" are either positive or negative floating point numbers
//
// or:
//
// <to^^<to
// <to^<to
// <=to^^<=to
// <=to^<=to
// [ < to ]
// [ <= to ]
// [<to]
// [<=to]
// where "to" is either positive or negative floating point number; "from" is set to invalid float (open start range)
//
// or:
//
// >from^^>from
// >from^>from
// >=from^^>=from
// >=from^>=from
// [ > from ]
// [ >= from ]
// [>from]
// [>=from]
// where "from" is either positive or negative floating point number; "to" is set to invalid float (open end of range)
//
// Returns error if the string cannot be parsed.
func ValueGeneratorFromRange(s string) (*ValueGenerator, error) {
	var matchGroups []string
	for _, r := range fromToRangeRegExp {
		if g := r.FindStringSubmatch(s); len(g) >= 3 {
			matchGroups = g
			break
		}
	}

	if len(matchGroups) >= 3 {
		from, err := floatFromString(matchGroups[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse from value: %s", matchGroups[1])
		}
		to, err := floatFromString(matchGroups[2])
		if err != nil {
			return nil, fmt.Errorf("failed to parse to value: %s", matchGroups[2])
		}
		if from >= to {
			return nil, fmt.Errorf("Cannot create ValueGenerator: start of the range [%f] "+
				"is greater than end of the range [%f]", from, to)
		}
		return &ValueGenerator{
			from: newValidFloat(from),
			to:   newValidFloat(to),
		}, nil
	}

	return valueGeneratorFromLessGreaterRange(s)
}

func valueGeneratorFromLessGreaterRange(s string) (*ValueGenerator, error) {
	for _, r := range lessRangeRegExp {
		matchGroups := r.FindStringSubmatch(s)
		if len(matchGroups) != 2 {
			continue
		}

		to, err := floatFromString(matchGroups[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse to value of less range: %s", matchGroups[1])
		}
		return &ValueGenerator{
			from: newInvalidFloat(),
			to:   newValidFloat(to),
		}, nil
	}

	for _, r := range greaterRangeRegExp {
		matchGroups := r.FindStringSubmatch(s)
		if len(matchGroups) != 2 {
			continue
		}

		from, err := floatFromString(matchGroups[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse from value of greater range: %s", matchGroups[1])
		}
		return &ValueGenerator{
			from: newValidFloat(from),
			to:   newInvalidFloat(),
		}, nil
	}

	return nil, fmt.Errorf("failed to parse the range: %s", s)
}
