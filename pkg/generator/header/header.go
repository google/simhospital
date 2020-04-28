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

// Package header contains functionality to generate headers.
package header

import (
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/message"
	"github.com/google/simhospital/pkg/pathway"
)

// Generator is a generator of headers.
type Generator struct {
	Header     *config.Header
	MsgCtrlGen *MessageControlGenerator
}

// NewHeader returns a HeaderInfo for the given step with a unique Message Control ID
// and values from the header configuration file and the step.
// The step must be non-nil.
func (g *Generator) NewHeader(step *pathway.Step) *message.HeaderInfo {
	header := g.Header.Default
	if step.StepType() == pathway.StepResults && g.Header.ORU != nil {
		header = *g.Header.ORU
	}
	h := &message.HeaderInfo{
		ReceivingApplication: header.ReceivingApplication,
		ReceivingFacility:    header.ReceivingFacility,
		SendingFacility:      header.SendingFacility,
		SendingApplication:   header.SendingApplication,
		MessageControlID:     g.MsgCtrlGen.NewMessageControlID(),
	}
	params := step.Parameters
	if params == nil {
		return h
	}
	if params.ReceivingApplication != "" {
		h.ReceivingApplication = params.ReceivingApplication
	}
	if params.ReceivingFacility != "" {
		h.ReceivingFacility = params.ReceivingFacility
	}
	if params.SendingFacility != "" {
		h.SendingFacility = params.SendingFacility
	}
	if params.SendingApplication != "" {
		h.SendingApplication = params.SendingApplication
	}
	return h
}
