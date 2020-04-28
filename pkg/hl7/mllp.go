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
	"io"
	"net"

	"github.com/pkg/errors"
	"github.com/google/simhospital/pkg/logging"
)

const (
	mllpStartBlock     = 0x0b
	mllpEndBlock       = 0x1c
	mllpCarriageReturn = 0x0d
)

var log = logging.ForCallerPackage()

// MLLPClient reads and writes MLLP.
type MLLPClient struct {
	r *bufio.Reader
	w io.Writer
}

// NewMLLPClient returns a new MLLPClient.
func NewMLLPClient(c net.Conn) *MLLPClient {
	return &MLLPClient{bufio.NewReader(c), c}
}

// Write writes a message to MLLP client.
// Returns an error if writing a message fails.
func (c *MLLPClient) Write(message []byte) error {
	// Don't wrap the errors with errors.Wrap(err, "..."). Using errors.Cause(err) on an error wrapped like that
	// does not retrieve the original cause, and this is needed by downstreams services.
	// We log it instead, as warnings to avoid duplicate error logging.
	log.Debugf("MLLP: Writing %q", message)
	if _, err := c.w.Write([]byte{mllpStartBlock}); err != nil {
		log.WithError(err).Warning("Cannot write mllp start block")
		return err
	}
	if _, err := c.w.Write(message); err != nil {
		log.WithError(err).Warning("Cannot write mllp message")
		return err
	}
	if _, err := c.w.Write([]byte{mllpEndBlock, mllpCarriageReturn}); err != nil {
		log.WithError(err).Warning("Cannot write mllp end block")
		return err
	}
	return nil
}

func (c *MLLPClient) Read() ([]byte, error) {
	log.Debug("MLLP: reading")
	b, err := c.r.ReadByte()
	if err != nil {
		return nil, errors.Wrap(err, "cannot read the first byte of mllp message")
	}
	if b != mllpStartBlock {
		return nil, errors.New("mllp: protocol error, missing Start Block")
	}
	payload, err := c.r.ReadBytes(mllpEndBlock)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read mllp message")
	}
	// Remove the mllpEndBlock
	payload = payload[:len(payload)-1]
	b, err = c.r.ReadByte()
	if err != nil {
		return nil, errors.Wrap(err, "cannot read the last byte of mllp message")
	}
	if b != mllpCarriageReturn {
		return nil, errors.New("mllp: protocol error, missing End Carriage Return")
	}
	return payload, nil
}
