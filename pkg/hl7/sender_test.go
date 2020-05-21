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
	"io/ioutil"
	"net"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/test/testwrite"
)

func TestNewFileSender(t *testing.T) {
	dir := testwrite.TempDir(t)
	filename := path.Join(dir, "output")

	fileSender, err := NewFileSender(filename)
	if err != nil {
		t.Fatalf("NewFileSender(%s) failed with %v", filename, err)
	}
	defer fileSender.Close()

	outputString := []byte("hello world")
	if err := fileSender.Send(outputString); err != nil {
		t.Errorf("Send(%v) failed with %v", outputString, err)
	}

	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("ReadFile() failed with %v", err)
	}

	if diff := cmp.Diff(append(outputString, []byte("\n\n")...), fileContents); diff != "" {
		t.Errorf("ioutil.ReadFile(%s) got diff (-want +got):\n%s", filename, diff)
	}
}

func TestNewFileSender_Error(t *testing.T) {
	dir := testwrite.TempDir(t)
	filename := path.Join(dir, "output")

	tests := []struct {
		name       string
		outputFile string
		wantErr    bool
	}{
		{"Success", filename, false},
		{"Empty filename", "", true},
		{"Nonexistent filename", "/foo/bar/baz", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewFileSender(tt.outputFile); (err != nil) != tt.wantErr {
				t.Errorf("NewFileSender(%s) got err=%v, want err? %t", tt.outputFile, err, tt.wantErr)
			}
		})
	}
}

func TestMllpSender(t *testing.T) {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf(`net.Listen("tcp", ":0") failed with %v`, err)
	}
	defer ln.Close()

	mllpSender, err := NewMLLPSender(ln.Addr().String(), true, time.Second)
	if err != nil {
		t.Fatalf("NewMLLPSender(%s, %t, %v) failed with %v", ln.Addr().String(), true, time.Second, err)
	}
	defer mllpSender.Close()

	done := make(chan bool, 1)
	defer close(done)
	want := "hl7_message"
	go func() {
		if err := mllpSender.Send([]byte(want)); err != nil {
			t.Errorf("mllpSender.Send(%s) failed with %v", want, err)
		}
		done <- true
	}()

	conn, err := ln.Accept()
	if err != nil {
		t.Fatalf("ln.Accept() failed with %v", err)
	}
	defer conn.Close()
	mllpClient := NewMLLPClient(conn)

	// Read the message sent using mllpSender and ACK it.
	// ACK needs to be a valid HL7 message.
	gotB, err := mllpClient.Read()
	if err != nil {
		t.Fatalf("mllpClient.Read() failed with %v", err)
	}
	if got := string(gotB); got != want {
		t.Errorf("mllpClient.Read() got %q, want %q", got, want)
	}

	mllpClient.Write([]byte("MSH|^~\\&|"))
	<-done
}

func TestMllpSender_InvalidAck(t *testing.T) {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf(`net.Listen("tcp", ":0") failed with %v`, err)
	}
	defer ln.Close()

	mllpSender, err := NewMLLPSender(ln.Addr().String(), true, time.Second)
	if err != nil {
		t.Fatalf("NewMLLPSender(%s, %t, %v) failed with %v", ln.Addr().String(), true, time.Second, err)
	}
	defer mllpSender.Close()

	done := make(chan bool, 1)
	defer close(done)
	msg := "hl7_message"
	go func() {
		if err := mllpSender.Send([]byte(msg)); err == nil {
			t.Errorf("mllpSender.Send(%s) got nil err, want non-nil err", msg)
		}
		done <- true
	}()

	conn, err := ln.Accept()
	if err != nil {
		t.Fatalf("ln.Accept() failed with %v", err)
	}
	defer conn.Close()
	mllpClient := NewMLLPClient(conn)
	mllpClient.Write([]byte("Ack"))
	<-done
}

func TestMllpSender_ReEstablishConn(t *testing.T) {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf(`net.Listen("tcp", ":0") failed with %v`, err)
	}
	defer ln.Close()

	addr := ln.Addr().String()

	mllpSender, err := NewMLLPSender(addr, true, time.Second)
	if err != nil {
		t.Fatalf("NewMLLPSender(%s, %t, %v) failed with %v", addr, true, time.Second, err)
	}
	defer mllpSender.Close()

	// Close and re-create the listener on the same address.
	ln.Close()
	ln, err = net.Listen("tcp", addr)
	if err != nil {
		t.Fatalf(`net.Listen("tcp", %s) failed with %v`, addr, err)
	}

	// mllpSender.Send() should re-establish the connection and send the message successfully.
	done := make(chan bool, 1)
	defer close(done)
	want := "hl7_message"
	go func() {
		if err := mllpSender.Send([]byte(want)); err != nil {
			t.Errorf("mllpSender.Send(%s) failed with %v", want, err)
		}
		done <- true
	}()

	conn, err := ln.Accept()
	if err != nil {
		t.Fatalf("ln.Accept() failed with %v", err)
	}
	defer conn.Close()
	mllpClient := NewMLLPClient(conn)

	// Read the message sent using mllpSender and ACK it.
	// ACK needs to be a valid HL7 message.
	gotB, err := mllpClient.Read()
	if err != nil {
		t.Fatalf("mllpClient.Read() failed with %v", err)
	}
	if got := string(gotB); got != want {
		t.Errorf("mllpClient.Read() got %q, want %q", got, want)
	}

	mllpClient.Write([]byte("MSH|^~\\&|"))
	<-done
}

func TestStdoutSender(t *testing.T) {
	// Capture stdout.
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() {
		os.Stdout = oldStdout
	}()

	stdoutSender := NewStdoutSender()
	defer stdoutSender.Close()

	msg := "hl7_message\rsegment_1\rsegment_2\r"
	want := "hl7_message\nsegment_1\nsegment_2\n"
	if err := stdoutSender.Send([]byte(msg)); err != nil {
		os.Stdout = oldStdout
		t.Fatalf("stdoutSender.Send(%s) failed with %v", msg, err)
	}

	w.Close()
	gotB, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout

	if got := string(gotB); !strings.Contains(got, want) {
		t.Errorf("stdoutSender.Send(%s) got %q, want containing %q", msg, got, want)
	}
}
