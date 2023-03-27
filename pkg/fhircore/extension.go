// Copyright 2023 Google LLC
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

package fhircore

import (
	"errors"
	"fmt"
	"sort"
	"time"

	pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
)

var (
	ErrMissingExtension = errors.New("extension not present")
	ErrMissingValue     = errors.New("malformed extension: missing value")
	ErrValue            = errors.New("malformed extension: cannot get value of the expected type")
)

// AddNewExtensionsOnly adds the "target" extensions to "source", as long as the extensions aren't present in "source".
func AddNewExtensionsOnly(source []*pb.Extension, target []*pb.Extension) []*pb.Extension {
	for _, te := range target {
		if existing := FindFirstExtension(source, te.GetUrl().GetValue()); existing == nil {
			source = append(source, te)
		}
	}
	return source
}

// AddOrUpdateExtension either updates the provided extension, if there is already one with the same
// URL, or inserts it at the end.
// It returns the updated list of extensions.
// If there are multiple extensions with the same URL, it returns an error.
func AddOrUpdateExtension(extensions []*pb.Extension, newExt *pb.Extension) ([]*pb.Extension, error) {
	url := newExt.GetUrl().GetValue()
	e, i := findAllExtensions(extensions, url)
	if len(e) > 1 {
		return nil, fmt.Errorf("Found %d %q extensions, expected 1", len(e), url)
	}
	if len(e) == 0 { // No existing extension.
		return append(extensions, newExt), nil
	}
	extensions[i[0]] = newExt
	return extensions, nil
}

func findAllExtensions(extensions []*pb.Extension, url string) ([]*pb.Extension, []int) {
	exts := []*pb.Extension{}
	var indexes []int
	for i, e := range extensions {
		if url == e.GetUrl().GetValue() {
			exts = append(exts, e)
			indexes = append(indexes, i)
		}
	}
	return exts, indexes
}

// FindAllExtensions returns all the extensions with the provided URL.
func FindAllExtensions(extensions []*pb.Extension, url string) []*pb.Extension {
	exts, _ := findAllExtensions(extensions, url)
	return exts
}

// FindFirstExtension returns the first Extension in extensions with the provided URL, or nil if not found.
func FindFirstExtension(extensions []*pb.Extension, url string) *pb.Extension {
	exts := FindAllExtensions(extensions, url)
	if len(exts) > 0 {
		return exts[0]
	}
	return nil
}

// FindLastExtension returns the last Extension in extensions with the provided URL, or nil if not found.
func FindLastExtension(extensions []*pb.Extension, url string) *pb.Extension {
	exts := FindAllExtensions(extensions, url)
	reverseExtensions(exts)
	if len(exts) > 0 {
		return exts[0]
	}
	return nil
}

func reverseExtensions(exts []*pb.Extension) {
	sort.SliceStable(exts, func(i, j int) bool {
		return i > j
	})
}

// getOnlyExtension gets the extension with the provided URL.
// It returns an ErrMissingExtension error if the extension is not present, and it also returns an
// error if there are multiple extensions with the provided URL.
// If possible, avoid using this method directly and use the methods that check the type of the
// extension (GetStringExtension, GetIntegerExtension...).
func getOnlyExtension(extensions []*pb.Extension, url string) (*pb.Extension, error) {
	exts := FindAllExtensions(extensions, url)
	if len(exts) == 0 {
		return nil, ErrMissingExtension
	}
	if len(exts) > 1 {
		return nil, fmt.Errorf("Found %d %q extensions, expected 1", len(exts), url)
	}
	// Callers of this method generally want to get the contained Value. Returning an error here
	// means they don't need to do this check themselves.
	if exts[0].GetValue() == nil {
		return nil, ErrMissingValue
	}
	return exts[0], nil
}

// GetDateTimeExtension gets the DateTime value from the extension with the provided name.
// If the extension is not found, it returns ErrMissingExtension.
// If the extension is found but it's not of DateTime type, it returns an error.
// Use DateTimeToTime on the returned value to convert it to time.Time.
func GetDateTimeExtension(extensions []*pb.Extension, name string) (*pb.DateTime, error) {
	e, err := getOnlyExtension(extensions, name)
	if err != nil {
		return nil, err
	}
	if _, ok := e.GetValue().GetChoice().(*pb.Extension_ValueX_DateTime); ok {
		return e.GetValue().GetDateTime(), nil
	}
	return nil, fmt.Errorf("Unexpected extension type for %q: expected DateTime", name)
}

// DeleteAllExtensions deletes all extensions with the matching URL from a list of Extension and returns the pruned list.
func DeleteAllExtensions(extensions []*pb.Extension, url string) []*pb.Extension {
	_, indexesToDelete := findAllExtensions(extensions, url)
	// Reverse indexesToDelete to start deleting from the end. Otherwise the indexes would change.
	for i, j := 0, len(indexesToDelete)-1; i < j; i, j = i+1, j-1 {
		indexesToDelete[i], indexesToDelete[j] = indexesToDelete[j], indexesToDelete[i]
	}

	for _, i := range indexesToDelete {
		extensions = append(extensions[:i], extensions[i+1:]...)
	}
	return extensions
}

// GetStringExtension gets the string contained in the extension with the provided name.
// If the extension is not found, it returns ErrMissingExtension.
func GetStringExtension(extensions []*pb.Extension, name string) (string, error) {
	e, err := getOnlyExtension(extensions, name)
	if err != nil {
		return "", err
	}
	if _, ok := e.GetValue().GetChoice().(*pb.Extension_ValueX_StringValue); ok {
		if e.GetValue().GetStringValue() == nil {
			return "", ErrValue
		}
		return e.GetValue().GetStringValue().GetValue(), nil
	}
	return "", errors.New("Unexpected extension type: expected StringValue")
}

// GetIntegerExtension gets the integer contained in the extension with the provided name.
// If the extension is not found, it returns ErrMissingExtension.
func GetIntegerExtension(extensions []*pb.Extension, name string) (int32, error) {
	e, err := getOnlyExtension(extensions, name)
	if err != nil {
		return 0, err
	}
	if _, ok := e.GetValue().GetChoice().(*pb.Extension_ValueX_Integer); ok {
		if e.GetValue().GetInteger() == nil {
			return 0, ErrValue
		}
		return e.GetValue().GetInteger().GetValue(), nil
	}
	return 0, errors.New("Unexpected extension type: expected Integer")
}

// GetBooleanExtension gets the boolean contained in the extension with the provided name.
// If the extension is not found, it returns ErrMissingExtension.
func GetBooleanExtension(extensions []*pb.Extension, name string) (bool, error) {
	e, err := getOnlyExtension(extensions, name)
	if err != nil {
		return false, err
	}
	if _, ok := e.GetValue().GetChoice().(*pb.Extension_ValueX_Boolean); ok {
		if e.GetValue().GetBoolean() == nil {
			return false, ErrValue
		}
		return e.GetValue().GetBoolean().GetValue(), nil
	}
	return false, errors.New("Unexpected extension type: expected Boolean")
}

// DateTimeExtension creates a DateTime extension.
func DateTimeExtension(url string, t time.Time, precision pb.DateTime_Precision) *pb.Extension {
	return &pb.Extension{
		Url: &pb.Uri{Value: url},
		Value: &pb.Extension_ValueX{
			Choice: &pb.Extension_ValueX_DateTime{
				DateTime: &pb.DateTime{
					ValueUs:   UnixMicro(t.UTC()),
					Timezone:  "UTC",
					Precision: precision,
				},
			},
		},
	}
}

// IntegerExtension creates an extension with an integer value.
func IntegerExtension(url string, value int32) *pb.Extension {
	return &pb.Extension{
		Url: &pb.Uri{Value: url},
		Value: &pb.Extension_ValueX{
			Choice: &pb.Extension_ValueX_Integer{
				Integer: Integer(value),
			},
		},
	}
}

// StringExtension creates an extension with a string value.
func StringExtension(url string, value string) *pb.Extension {
	return &pb.Extension{
		Url: &pb.Uri{Value: url},
		Value: &pb.Extension_ValueX{
			Choice: &pb.Extension_ValueX_StringValue{
				StringValue: String(value),
			},
		},
	}
}

// BooleanExtension creates an extension with a boolean value.
func BooleanExtension(url string, value bool) *pb.Extension {
	return &pb.Extension{
		Url: &pb.Uri{Value: url},
		Value: &pb.Extension_ValueX{
			Choice: &pb.Extension_ValueX_Boolean{
				Boolean: Boolean(value),
			},
		},
	}
}
