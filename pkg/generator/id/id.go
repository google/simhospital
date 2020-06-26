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

// Package id provides the functionality to generate identifiers.
package id

import "github.com/google/uuid"

// Generator is an interface to generate identifiers.
type Generator interface {
	NewID() string
}

// UUIDGenerator is a wrapper for github.com/google/uuid that implements the Generator interface.
type UUIDGenerator struct{}

// NewID returns a new random UUID.
func (*UUIDGenerator) NewID() string {
	return uuid.New().String()
}
