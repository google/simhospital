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

package config

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/google/simhospital/pkg/files"
)

var digitsPattern = regexp.MustCompile("^\\d+$")

// names loads person names from the given file.
// The files contain the top baby names by year.
// It expects the following format of the file:
//
//   RANK,1904,1914,1924
//   ,,,,,,,,,,
//   1,MARY,MARY,MARGARET
//   2,FLORENCE,MARGARET,MARY
//
// The first line (the header) represents the years in chronological order.
// The rest of lines contain the rank number, followed by the Names with
// that rank per given year.
// Eg. in the example above, in 1904, the 1st most popular name was MARY,
// and the 2nd most popular name was FLORENCE.
// This allows generating the random name which was popular among people born
// in a given year.
func names(ctx context.Context, filename string) (*Names, error) {
	namesByYear := map[int][]string{}
	allNames := map[string]bool{}
	var years []int
	b, err := files.Read(ctx, filename)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(bytes.NewReader(b))
	r.Comment = '#'
	headerParsed := false
	for {
		row, err := r.Read()
		if err != nil {
			if err != io.EOF {
				log.WithError(err).Warning("Error reading file")
			}
			break
		}
		if len(row) > 0 && row[0] == "RANK" {
			// The first row contains headers - ie the years.
			if headerParsed {
				return nil, fmt.Errorf("duplicated header line: %v", row)
			}
			if years, err = parseHeaders(row); err != nil {
				return nil, errors.Wrap(err, "cannot parse headers")
			}
			headerParsed = true
		}
		if len(row) > 0 && digitsPattern.MatchString(row[0]) {
			if !headerParsed {
				return nil, fmt.Errorf("the line containing ranked names found: %v before headers are parsed", row)
			}
			for i, year := range years {
				n := strings.Title(strings.ToLower(row[i+1]))
				namesByYear[year] = append(namesByYear[year], n)
				allNames[n] = true
			}
		}
	}
	all := make([]string, 0, len(allNames))
	for k := range allNames {
		all = append(all, k)
	}

	return &Names{ByYear: namesByYear, All: all, MinYear: years[0], MaxYear: years[len(years)-1]}, nil
}

func parseHeaders(row []string) ([]int, error) {
	var years []int
	for i := 1; i < len(row); i++ {
		year, err := strconv.Atoi(row[i])
		if err != nil {
			return nil, errors.Wrapf(err, "cannot parse the year: %q", row[i])
		}
		if len(years) > 0 && year <= years[len(years)-1] {
			return nil, fmt.Errorf("years are expected in chronological order; got year %d after %v", year, years)
		}
		years = append(years, year)
	}
	return years, nil
}

// Names contains the division of names per year, and a list of all names.
type Names struct {
	// ByYear contains a list of names that were popular at the time period
	// up to a given year. Eg. if the map contains the following data:
	// 1904: [Mary], 1914: [Dory], 1924: [Florence]
	// it means, that:
	// - Mary was the most popular name in the time period up to 1904, and since
	//   there is no more historical data, it would be also used as the most
	//   popular name for the whole time period preceding 1904,
	// - Dory was the most popular name in the time period between 1905 - 1914,
	// - Florence was the most popular name in the time period between 1915 - 1924
	//   and since there is no newer data, it would be also used as the most
	//   popular name for all the years succeeding 1924.
	// Note that the time periods don't have to be equal in length.
	ByYear map[int][]string
	// All contains the list of all unique names present in ByYear map,
	// regardless of the time period when they were popular.
	All []string
	// MinYear is the minimum year present in the ByYear map.
	MinYear int
	// MaxYear is the maximum year present in the ByYear map.
	MaxYear int
}

// FirstNamesByCensus represents Girls and Boys name grouped by their
// popularity by year.
type FirstNamesByCensus struct {
	Girls *Names
	Boys  *Names
}

// census loads Girls and Boys first Names from bytes files specified by
// the respective parameters.
func census(ctx context.Context, fileNameGirls string, fileNameBoys string) (*FirstNamesByCensus, error) {
	girls, err := names(ctx, fileNameGirls)
	if err != nil {
		return nil, err
	}
	boys, err := names(ctx, fileNameBoys)
	if err != nil {
		return nil, err
	}
	return &FirstNamesByCensus{Girls: girls, Boys: boys}, nil
}
