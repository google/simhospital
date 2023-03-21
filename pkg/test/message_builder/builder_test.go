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

package message_builder

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/simhospital/pkg/hl7"
)

func TestMain(m *testing.M) {
	hl7.TimezoneAndLocation("UTC")
	retCode := m.Run()
	os.Exit(retCode)
}

func TestBuildMessage_NoMessageType(t *testing.T) {
	builder := NewBuilderForTests()
	builder.MessageType.TriggerEvent = "O01"

	if _, err := builder.buildMessage(); err == nil {
		t.Error("buildMessage() got nil error, want error because the MessageType hasn't been set")
	}
}

func TestBuildMessage_NoTrigger(t *testing.T) {
	builder := NewBuilderForTests()
	builder.MessageType.MessageType = "ADT"

	if _, err := builder.buildMessage(); err == nil {
		t.Error("buildMessage() got nil error, want error because the TriggerEvent hasn't been set")
	}
}

func TestBuildMessage_UnknownMessageType(t *testing.T) {
	builder := NewBuilderForTests()
	builder.MessageType.MessageType = "ADT"
	builder.MessageType.MessageType = "O01"

	if _, err := builder.buildMessage(); err == nil {
		t.Error("buildMessage() got nil error, want error because the combination of TriggerEvent and MessageType is not correct")
	}
}

func TestBuildMessage_IncrementCurrentDate(t *testing.T) {
	builder := NewBuilderForTests()
	builder.MessageType = AdtA01 // Arbitrary message type (irrelevant for this test).

	first := builder.currentDate

	// Invoking BuildMessage moves the time forward, a time period equal to *timeIncrements.
	builder.BuildMessage(t)
	second := builder.currentDate

	if second.Sub(first) != *timeIncrements {
		t.Fatalf("BuildMessage.currentDate.Sub(%v) got %v, want %v.", first, second.Sub(first), *timeIncrements)
	}
}

func TestBuildMessage_ValidMessage(t *testing.T) {
	for _, messageType := range supportedMessages {
		t.Run(fmt.Sprintf("%s-%s", messageType.MessageType, messageType.TriggerEvent), func(t *testing.T) {
			builder := NewBuilderForTests()
			builder.MessageType.MessageType = messageType.MessageType
			builder.MessageType.TriggerEvent = messageType.TriggerEvent

			message := builder.BuildMessage(t)
			if len(message) == 0 {
				t.Fatal("BuildMessage() message is empty, want non empty")
			}

			// Assert that the correct message was created.
			mo := hl7.NewParseMessageOptions()
			mo.TimezoneLoc = time.UTC
			m, err := hl7.ParseMessageWithOptions([]byte(message), mo)
			if err != nil {
				t.Fatalf("ParseMessageWithOptions failed with %v", err)
			}
			if m == nil {
				t.Fatal("ParseMessageWithOptions got message=<nil>, want not nil.")
			}
			msh, err := m.MSH()
			if err != nil {
				t.Fatalf("MSH() failed with %v", err)
			}
			if msh == nil {
				t.Fatal("MSN() got msh=<nil>, want not nil.")
			}
			if got, want := msh.MessageType.MessageCode.String(), messageType.MessageType; got != want {
				t.Errorf("msh.MessageType.MessageCode got %v, want %v.", got, want)
			}
			if got, want := msh.MessageType.TriggerEvent.String(), messageType.TriggerEvent; got != want {
				t.Errorf("msh.MessageType.TriggerEvent got %v, want %v.", got, want)
			}
		})
	}
}

func TestCurrentTime(t *testing.T) {
	builder := NewBuilderForTests()
	builder.MessageType = AdtA01 // Arbitrary message type (irrelevant for this test).

	first := builder.currentDate
	firstT := builder.CurrentTime()
	if got, want := firstT.Time, first; !firstT.Time.Equal(first) {
		t.Errorf("builder.CurrentTime() got %v, want %v.", got, want)
	}

	// If the current time changes, the time returned by CurrentTime must still be in sync with
	// "currentDate".
	builder.BuildMessage(t)
	second := builder.currentDate
	secondT := builder.CurrentTime()
	if got, want := secondT.Time, second; !secondT.Time.Equal(second) {
		t.Errorf("builder.CurrentTime() got %v, want %v.", got, want)
	}
}
