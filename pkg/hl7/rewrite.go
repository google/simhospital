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

package hl7

import "fmt"

// Rewrite represents functions that take a Token, and return the underlying byte sequence value,
// having potentially rewritten that sequence to correct known message errors.
// A RewriteResult is also returned and it determines how the byte sequence is used:
// ReplaceValue: the token Value is replaced with the byte sequence returned.
// Noop: the token is not modified. The byte sequence is ignored.
// DeleteToken: the token is removed entirely. The byte sequence is ignored.
//
// This function is invoked for each segment, field and value within the message, and thus allows
// rewriting all those elements. Use the token's Location to obtain the current position within the
// message.
// Empty fields or values can be rewritten; however, the placeholders for such fields or values need
// to be present. If you need to add fields or values in order to rewrite them, rewrite the entire
// segment first and add the placeholders.
// Fields or values can be made nil by rewriting them with []byte{} or returning DeleteToken.
// Segments can be removed by returning DeleteToken.
// It is not possible to add segments with this function.
//
// Examples:
//
// 1. Rewrite a value within a field.
//
//   // Rewrite the Zip code of a PID.
//                                      ┌ Placeholder for the zip code
//   // Example PID: "PID|1||||||||||^^^^^|"
//   rewrite := func(t Token) *RewriteResult {
//    if t.Location == "PID-11-Patient Address/XAD-5-Zip Or Postal Code" {
//      return RewriteResultReplaceValue([]byte("49999"))
//    }
//    return RewriteResultNoop()
//   }
//
// 2. Rewrite an entire field.
//
//   // Rewrite the entire Mother's Maiden Name field or a PID. The fields needs to be present.
//                             ┌ Placeholder for the Mother's Maiden Name.
//   // Example PID: "PID|1||||||"
//   rewrite := func(t Token) *RewriteResult {
//     if t.Location == "PID-6-Mother'S Maiden Name" {
//       return RewriteResultReplaceValue([]byte("Doe^Jane"))
//     }
//     return RewriteResultNoop()
//   }
//
// 3. Rewrite an entire segment.
//
//   // Rewrite the entire PID to add more placeholders.
//   // Example PID: "PID|1||"
//   rewrite := func(t Token) *RewriteResult {
//     if t.Location == "PID" {
//       return RewriteResultReplaceValue([]byte("PID|1|||||Doe^Jane"))
//     }
//     return RewriteResultNoop()
//   }
//
// 4. Delete an entire segment.
//
//   // Delete all OBX segments.
//   // Example OBX: "OBX|1||"
//   rewrite := func(t Token) *RewriteResult {
//     if t.Location == "OBX" {
//       return RewriteResultDeleteToken()
//     }
//     return RewriteResultNoop()
//   }
//
type Rewrite func(t Token) *RewriteResult

// RewriteResult determines the action to take for a rewrite, optionally including a new value.
// Use the RewriteResult* functions to create instances.
type RewriteResult struct {
	value  []byte
	action rewriteAction
}

type rewriteAction int

const (
	// The int default (0) is not a valid action.
	noop         rewriteAction = 1
	replaceValue rewriteAction = 2
	deleteToken  rewriteAction = 3
)

// RewriteResultNoop returns a RewriteResult that is a no-op.
func RewriteResultNoop() *RewriteResult {
	return &RewriteResult{action: noop}
}

// RewriteResultReplaceValue returns a RewriteResult that replaces the current value with newValue.
func RewriteResultReplaceValue(newValue []byte) *RewriteResult {
	return &RewriteResult{value: newValue, action: replaceValue}
}

// RewriteResultDeleteToken returns a RewriteResult that deletes the current token.
func RewriteResultDeleteToken() *RewriteResult {
	return &RewriteResult{action: deleteToken}
}

// NopRewrite returns the underlying byte sequence value for a Token without
// any rewriting.
func NopRewrite(_ Token) *RewriteResult {
	return RewriteResultNoop()
}

// rewrite applies the Rewrite functions to the token in order.
func rewrite(c *Context, input Token) (*RewriteResult, error) {
	// Need to keep track of the latest rewrite that had some effect,
	// e.g. if there's a ReplaceValue and then a Noop, we should return the ReplaceValue.
	finalResult := RewriteResultNoop()
	for _, rewrite := range c.Rewrite {
		result := rewrite(input)
		switch result.action {
		case noop:
			// Do nothing.
		case replaceValue:
			input.Value = result.value
			finalResult = result
		case deleteToken:
			// No point doing more rewrites after a deletion.
			return result, nil
		default:
			return nil, fmt.Errorf("Unknown rewriteAction value: %v", result.action)
		}
	}
	return finalResult, nil
}
