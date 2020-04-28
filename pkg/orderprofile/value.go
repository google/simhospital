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
	"regexp"
	"strconv"
	"strings"

	"github.com/google/simhospital/pkg/logging"
)

var (
	log = logging.ForCallerPackage()

	valueRexEx = regexp.MustCompile("^([>|<]?=?) *(-?[0-9|/.]+)$")
)

// ValueFromString parses string in one of the following formats:
// value
// >value
// >=value
// <value
// <=value
// where value is either positive of negative floating point number.
// Returns the prefix ( > / >= / < / <= ) and the parsed number.
// Note: prefix is optional and defaults to an empty string.
// Returns an error if value can not be parsed.
func ValueFromString(s string) (string, float64, error) {
	groups := valueRexEx.FindStringSubmatch(s)
	if len(groups) == 3 {
		v, err := floatFromString(groups[2])
		if err != nil {
			return "", 0, fmt.Errorf("Cannot parse %s to float", s)
		}

		return groups[1], v, nil
	}
	return "", 0, fmt.Errorf("Cannot parse %s", s)
}

// floatFromString parses the string to a float64.
// Trimps spaces before trying to parse it.
// Returns an error if the string cannot be parsed.
func floatFromString(s string) (float64, error) {
	return strconv.ParseFloat(strings.TrimSpace(s), 64)
}
