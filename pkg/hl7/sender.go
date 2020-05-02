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
	"bytes"
	"fmt"
	"net"
	"os"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

// Sender is an interface for sending HL7 messages.
type Sender interface {
	Send([]byte) error
	Close() error
}

// stdoutSender sends HL7 messages to the standard output.
type stdoutSender struct {
	count int
}

// NewStdoutSender returns a sender that sends HL7 messages to the standard output.
func NewStdoutSender() Sender {
	return &stdoutSender{}
}

// Send sends a message to stdout.
func (s *stdoutSender) Send(message []byte) error {
	fmt.Print(string(bytes.Replace(message, []byte(SegmentTerminatorStr), []byte("\n"), -1)))
	fmt.Print("\n")
	s.count++
	return nil
}

// Close prints the number of messages that have been sent.
func (s *stdoutSender) Close() error {
	log.Infof("Messages successfully sent by the stdoutSender: %d", s.count)
	return nil
}

var recoverableErrs = map[syscall.Errno]bool{syscall.EPIPE: true, syscall.ECONNRESET: true}

// mllpSender sends HL7 messages via the MLLP protocol.
type mllpSender struct {
	client              *MLLPClient
	conn                net.Conn
	address             string
	mllpKeepAlive       bool
	mllpKeepAlivePeriod time.Duration
	count               int
}

// NewMLLPSender returns a sender that sends HL7 messages via the MLLP protocol.
func NewMLLPSender(address string, mllpKeepAlive bool, mllpKeepAlivePeriod time.Duration) (Sender, error) {
	sender := &mllpSender{
		address:             address,
		mllpKeepAlive:       mllpKeepAlive,
		mllpKeepAlivePeriod: mllpKeepAlivePeriod,
	}
	if err := sender.establishConnection(); err != nil {
		return nil, errors.Wrapf(err, "cannot establish mllp connection on sender %+v", sender)
	}
	return sender, nil
}

func isRecoverable(err error) bool {
	if opError, ok := err.(*net.OpError); ok {
		if syscallErr, ok := opError.Err.(*os.SyscallError); ok {
			if errno, ok := syscallErr.Err.(syscall.Errno); ok && recoverableErrs[errno] {
				return true
			}
		}
	}

	return false
}

func (s *mllpSender) establishConnection() error {
	conn, err := net.Dial("tcp", s.address)
	if err != nil {
		return errors.Wrapf(err, "cannot connect to tcp address %s", s.address)
	}

	if s.mllpKeepAlive {
		if err := conn.(*net.TCPConn).SetKeepAlive(true); err != nil {
			return errors.Wrapf(err, "cannot set keep alive on connection %v", conn)
		}
		if err := conn.(*net.TCPConn).SetKeepAlivePeriod(s.mllpKeepAlivePeriod); err != nil {
			return errors.Wrapf(err, "cannot set keep alive period on connection %v", conn)
		}
	}

	s.conn = conn
	s.client = NewMLLPClient(conn)
	return nil
}

// Send sends a messages via the MLLP protocol.
// It returns an error if the message cannot be sent or was not acknowledged.
func (s *mllpSender) Send(message []byte) error {
	if err := s.client.Write(message); err != nil {
		if !isRecoverable(err) {
			return errors.Wrap(err, "cannot send message")
		}
		// If the socket was closed by the peer, handle it by trying to
		// write once again on a new connection.
		if err = s.establishConnection(); err != nil {
			return errors.Wrap(err, "cannot send message: error when re-establishing connection")
		}
		if err = s.client.Write(message); err != nil {
			return errors.Wrap(err, "cannot send message after re-establishing connection")
		}
	}

	ack, err := s.client.Read()
	if err != nil {
		return errors.Wrap(err, "cannot read an ack after sending message")
	}

	if _, err = ParseMessage(ack); err != nil {
		return errors.Wrap(err, "ack message cannot be parsed")
	}
	s.count++
	return nil
}

// Close closes the underlying TCP connection.
// It should be called, when the mllpSender is not needed anymore or at the program exit.
// Close prints the number of messages that have been sent.
func (s *mllpSender) Close() error {
	log.Infof("Messages successfully sent by the mllpSender: %d", s.count)
	if err := s.conn.Close(); err != nil {
		return errors.Wrap(err, "closing mllp sender connection")
	}
	return nil
}

// fileSender sends HL7 messages to a file.
type fileSender struct {
	file  *os.File
	count int
}

// NewFileSender creates a sender that sends HL7 messages to a file.
func NewFileSender(destFilename string) (Sender, error) {
	if destFilename == "" {
		return nil, errors.New("output filename must be nonempty if outputting to a file")
	}
	file, err := os.Create(destFilename)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot create output file %s", destFilename)
	}
	return &fileSender{file: file}, nil
}

// Send sends a message to the file.
func (s *fileSender) Send(message []byte) error {
	if _, err := s.file.Write(append(message, []byte("\n\n")...)); err != nil {
		return errors.Wrap(err, "cannot write a message")
	}
	s.count++
	return nil
}

// Close closes the underlying file.
// It should be called, when the mllpSender is not needed anymore or at the program exit.
// Close prints the number of messages that have been sent.
func (s *fileSender) Close() error {
	log.Infof("Messages successfully sent by the fileSender: %d", s.count)
	if err := s.file.Close(); err != nil {
		return errors.Wrap(err, "closing file sender")
	}
	return nil
}
