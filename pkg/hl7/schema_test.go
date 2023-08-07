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
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts/cmpopts"
)

func TestNestingDepthNeverExceedsTwo(t *testing.T) {
	// HL7 only supports three levels of nesting (| - field, ^ - component,
	// & - subcomponent). Some of the types are broken in the specification, as
	// their nesting level makes them unencodable. As handling this situation
	// needs special casing, this test ensures the set of broken types doesn't
	// change.
	segment := reflect.ValueOf((*Segment)(nil)).Type().Elem()
	brokenTypes := map[string]bool{}
	for _, t := range Types {
		if !reflect.PtrTo(t).Implements(segment) {
			continue
		}
		for _, description := range findTypesWithMoreThanTwoLevelsOfNesting(t) {
			brokenTypes[description] = true
		}
	}
	// The TQ type in HL7 2.4 is technically unencodable, as the nesting
	// level is too deep. This is corrected in HL7 2.5, where the last
	// CE is redefined as ST. Skip this known error.
	// The XCN,DR,TS stack (and a few others ending in TS) are also
	// unencodable, though as the second component of TS has been
	// deprecated, it's not an issue in practice. This issue isn't seen
	// here as TS (incorrectly) implements Primitive currently.
	want := map[string]bool{
		"ECD,TQ,CQ,CE": true,
		"GOL,TQ,CQ,CE": true,
		"INV,TQ,CQ,CE": true,
		"OBR,TQ,CQ,CE": true,
		"ORC,TQ,CQ,CE": true,
		"QRF,TQ,CQ,CE": true,
		"RXE,TQ,CQ,CE": true,
		"RXG,TQ,CQ,CE": true,
		"SCH,TQ,CQ,CE": true,
		"URS,TQ,CQ,CE": true,
	}
	if diff := cmp.Diff(want, brokenTypes, cmpopts.SortSlices(func(a, b string) bool { return a < b })); diff != "" {
		t.Errorf("broken types comparison returned diff (-want +got):\n%s", diff)
	}
}

func TestTypesGeneratedForAllHL7Versions(t *testing.T) {
	for typ, def := range map[string]string{
		"RX1":     "a segment last defined in HL7 2.1",
		"ADT_A26": "a message type last defined in HL7 2.2",
		"ORN_O02": "a message type last defined in HL7 2.3.1",
		"ACK_N02": "a message type last defined in HL7 2.4",
	} {
		if _, ok := Types[typ]; !ok {
			t.Errorf("Missing %q in Types (%v)", typ, def)
		}
	}
}

func isCompositeType(t reflect.Type) bool {
	primitive := reflect.ValueOf((*Primitive)(nil)).Type().Elem()
	pointerToComposite := t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct && !t.Implements(primitive)
	sliceOfComposites := t.Kind() == reflect.Slice && t.Elem().Kind() == reflect.Struct && !reflect.PtrTo(t.Elem()).Implements(primitive)
	return pointerToComposite || sliceOfComposites
}

func findTypesWithMoreThanTwoLevelsOfNesting(segment reflect.Type) []string {
	type entry struct {
		t reflect.Type
		i int
	}
	// An array of descriptions of types that are nested too deeply, with entries
	// of the form "INV,TQ,CQ,CE" (describing the segment INC with a CE nested
	// within CQ nested within TQ).
	var result []string
	stack := []entry{{segment, 0}}
	for {
		if len(stack) == 0 {
			break
		}
		e := &stack[len(stack)-1]
		if e.i >= e.t.NumField() {
			stack = stack[0 : len(stack)-1]
			continue
		}
		ft := e.t.Field(e.i).Type
		e.i++
		if isCompositeType(ft) {
			stack = append(stack, entry{ft.Elem(), 0})
			if len(stack) > 3 { // 3 levels max: Field, Component, Subcomponent.
				var names []string
				for _, e := range stack {
					names = append(names, e.t.Name())
				}
				result = append(result, strings.Join(names, ","))
			}
		}
	}
	return result
}
