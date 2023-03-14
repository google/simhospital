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

package hl7

import (
	"reflect"
	"unsafe"
)

// ParseMessageV2 parses a message, and returns a V2 type.
// TODO: Support Rewriting.
func ParseMessageV2(input []byte) (interface{}, error) {
	m, err := ParseMessage(input)
	if err != nil {
		return nil, err
	}
	name, err := m.messageTypeName()
	if err != nil {
		return nil, err
	}
	t, ok := Types[name+"v2"]
	if !ok {
		return nil, &BadMessageTypeError{Name: name}
	}
	segments, err := m.All()
	if err != nil {
		return nil, err
	}

	result := reflect.New(t)
	fillGroup(result, segments)
	return result.Interface(), nil
}

// fillGroup takes an empty struct representing a group of HL7 segments, g, and
// a list of segments, and attempts to match those segments to fields within the
// group. Nested subgroups are created when necessary. Segments are assigned
// to fields by looking at the type of the field. If the current segment
// doesn't match the current field, we consult the follow set for that field
// to decide whether that segment can be assigned to a following field - in
// which case, the current field is skipped, as the segment eventually assigned.
// If the segment can't be assigned to any following field, it's discarded.
// TODO: Return an error if required segments are missing
// TODO: Communicate skipped segments to the caller
func fillGroup(g reflect.Value, segments []interface{}) []interface{} {
	g = g.Elem()
	for i := 0; i < g.NumField() && len(segments) > 0; {
		f := g.Field(i)
		if isSegmentType(f.Type()) {
			st := reflect.TypeOf(segments[0])
			if f.Type().Kind() == reflect.Slice && f.Type().Elem() == st {
				f = reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
				existing := f
				if existing.IsNil() {
					existing = reflect.MakeSlice(f.Type(), 0, 0)
				}
				f.Set(reflect.Append(existing, reflect.ValueOf(segments[0])))
				segments = segments[1:]
			} else if f.Type() == st {
				f = reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
				f.Set(reflect.ValueOf(segments[0]))
				i++
				segments = segments[1:]
			} else {
				fs, ok := FollowSets[followSetKey(g, i)]
				if !ok || !fs[st.Elem().Name()] {
					segments = segments[1:]
				} else {
					i++
				}
			}
		} else if f.Type().Kind() == reflect.Slice {
			new := reflect.New(f.Type().Elem().Elem()) // Elem.Elem -> []*Group to Group
			remaining := fillGroup(new, segments)
			if len(remaining) < len(segments) {
				f = reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
				existing := f
				if existing.IsNil() {
					existing = reflect.MakeSlice(f.Type(), 0, 0)
				}
				f.Set(reflect.Append(existing, new))
				segments = remaining
			} else {
				i++
			}
		} else {
			new := reflect.New(f.Type().Elem()) // Elem -> *Group to Group
			remaining := fillGroup(new, segments)
			if len(remaining) < len(segments) {
				f = reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
				f.Set(new)
				segments = remaining
			}
			i++
		}
	}
	return segments
}

func followSetKey(g reflect.Value, i int) string {
	return g.Type().Name() + "." + g.Type().Field(i).Name
}
