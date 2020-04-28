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

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRead(t *testing.T) {
	tests := []struct {
		name        string
		reader      io.Reader
		wantErr     bool
		wantPayload []byte
	}{{
		name:        "Correct",
		reader:      bytes.NewBufferString(fmt.Sprintf("\x0bHL7 Payload\x1c\x0d")),
		wantPayload: []byte("HL7 Payload"),
	}, {
		name:    "Bad byte",
		reader:  bytes.NewBufferString(fmt.Sprintf("\xffHL7 Payload\x1c\x0d")), // \xff is invalid
		wantErr: true,
	}, {
		name:    "insufficient data",
		reader:  bytes.NewBufferString(fmt.Sprintf("\xffHL7 Pay")), // too short
		wantErr: true,
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client := &MLLPClient{r: bufio.NewReader(tc.reader), w: nil}
			payload, err := client.Read()
			gotErr := err != nil
			if gotErr != tc.wantErr {
				t.Errorf("client.Read() got err=%v, want error? %v", err, tc.wantErr)
			}
			if diff := cmp.Diff(tc.wantPayload, payload); diff != "" {
				t.Errorf("client.Read() got diff (-want, +got):\n%s", diff)
			}
		})
	}
}

func TestWrite(t *testing.T) {
	tests := []struct {
		name        string
		writer      io.Writer
		wantErr     bool
		wantPayload string
	}{{
		name:   "Correct",
		writer: &bytes.Buffer{},
	}, {
		name:    "Bad writer",
		writer:  &badWriter{5}, // 5 bytes is not enough to hold the payload
		wantErr: true,
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client := &MLLPClient{nil, tc.writer}
			err := client.Write([]byte("HL7 Payload"))
			gotErr := err != nil
			if gotErr != tc.wantErr {
				t.Errorf("client.Write() got err=%v, want error? %v", err, tc.wantErr)
			}
			if gotErr || tc.wantErr {
				return
			}
			if got, want := tc.writer.(*bytes.Buffer).String(), "\x0bHL7 Payload\x1c\x0d"; got != want {
				t.Errorf("client.writer after client.Write() got %v, want %v", got, want)
			}
		})
	}
}

type badWriter struct {
	Capacity int
}

func (w *badWriter) Write(p []byte) (n int, err error) {
	if len(p) > w.Capacity {
		return 0, io.ErrShortWrite
	}
	w.Capacity -= len(p)
	return len(p), nil
}
