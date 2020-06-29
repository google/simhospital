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

package testresource

import (
	"bytes"
	"errors"
	"io"

	"github.com/google/simhospital/pkg/ir"
)

// ByteOutput is a resource output that writes to a byte buffer.
type ByteOutput struct {
	Bytes *bytes.Buffer
}

// New returns the ByteWriteCloser.
func (o *ByteOutput) New(p *ir.PatientInfo) (io.WriteCloser, error) {
	if o.Bytes == nil {
		return nil, errors.New("nil byte buffer")
	}
	return &byteWriteCloser{Bytes: o.Bytes}, nil
}

// byteWriteCloser is a wrapper around bytes.Buffer which implements the WriteCloser interface.
type byteWriteCloser struct {
	Bytes *bytes.Buffer
}

// Write writes bytes to the byte buffer.
func (w *byteWriteCloser) Write(bytes []byte) (int, error) {
	return w.Bytes.Write(bytes)
}

// Close is a no-op.
func (*byteWriteCloser) Close() error {
	return nil
}
