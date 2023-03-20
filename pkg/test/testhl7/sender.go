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

// Package testhl7 contains utility functions and helpers for testing with HL7 messages.
package testhl7

// Sender is an hl7.Sender that saves all sent messages in a list.
type Sender struct {
	messages []string
	err      error
}

// SenderWithError returns a Sender that returns the provided error when Send() is invoked.
func SenderWithError(err error) *Sender {
	return &Sender{err: err}
}

// Send saves all sent messages to a list if s.err is nil.
// Otherwise, it returns s.err.
func (s *Sender) Send(message []byte) error {
	if s.err != nil {
		return s.err
	}
	s.messages = append(s.messages, string(message))
	return nil
}

// Close is no-op.
func (s *Sender) Close() error {
	return nil
}

// GetSentMessages returns the messages sent by this sender.
func (s *Sender) GetSentMessages() []string {
	return s.messages
}

// EraseSentHistory erases the history of sent messages.
func (s *Sender) EraseSentHistory() {
	s.messages = []string{}
}
