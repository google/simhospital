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

// Package hardcoded contains features to manage hardcoded messages in Simulated
// Hospital. Messages are filtered according to a user-supplied regular expression.
// A single message is selected from the filtered messages, according to a uniform
// distribution. Its missing fields are filled in, according to user-supplied
// values and a unique MessageControlID is added, so that the message is considered
// unique by the downstream processes.
package hardcoded

import (
	"context"
	"fmt"
	"math/rand"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"github.com/google/simhospital/pkg/files"
	"github.com/google/simhospital/pkg/generator/header"
	"github.com/google/simhospital/pkg/hl7"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/logging"
	"github.com/google/simhospital/pkg/message"
)

// pidSegmentPlaceholder defines a placeholder for the PID segments, which
// should be filled in when the message is returned by this class.
const pidSegmentPlaceholder = "PID_SEGMENT_PLACEHOLDER"

// validExtensions defines which file extensions are valid for hardcoded messages.
var validExtensions = []string{".yml", ".yaml"}

var log = logging.ForCallerPackage()

// Manager contains all hardcoded messages.
type Manager struct {
	// messages is a map of all hardcoded messages indexed by their name.
	messages map[string]string
	// generator produces unique message IDs.
	generator *header.MessageControlGenerator
}

type hardcodedMessage struct {
	Segments []string
}

// NewManager returns a Manager for the messages contained in the given directory.
func NewManager(ctx context.Context, messageDir string, headerGenerator *header.MessageControlGenerator) (*Manager, error) {
	files, err := files.List(ctx, messageDir)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read hardcoded messages directory: %s", messageDir)
	}

	messages := map[string]string{}
	for _, file := range files {
		if !fileExtensionIsValid(file.Name()) {
			log.Warnf("File name has invalid extension %s, expected one of %+v. Skipping...", file.Name(), validExtensions)
			continue
		}

		messagesInFile, err := parseFile(ctx, file)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot parse hardcoded messages file %s", file.FullPath())
		}

		for name, msg := range messagesInFile {
			if _, ok := messages[name]; ok {
				return nil, fmt.Errorf("hardcoded message with name %q re-declared", name)
			}
			messages[name] = msg
		}
	}

	if len(messages) == 0 {
		return nil, errors.New("no hardcoded messages were loaded")
	}

	log.Infof("Loaded a total of %d hardcoded messages from dir %s", len(messages), messageDir)
	for name := range messages {
		log.Infof(" - %s", name)
	}

	return &Manager{
		messages:  messages,
		generator: headerGenerator,
	}, nil
}

func parseFile(ctx context.Context, file files.File) (map[string]string, error) {
	input := map[string]hardcodedMessage{}
	data, err := file.Read(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read file: %s", file.FullPath())
	}
	err = yaml.UnmarshalStrict(data, &input)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot unmarshal messages from file: %s", file.FullPath())
	}
	messages := map[string]string{}
	for k, v := range input {
		messages[k] = strings.Join(v.Segments, message.SegmentTerminator)
	}
	return messages, nil
}

// Message returns a uniformly random hardcoded message from the set of messages
// with names that match the provided regular expression. If the chosen hardcoded
// message has a PID segment placeholder, this will be filled in with a PID
// segment based on the provided Person argument. Otherwise, the PID segment
// will be left as it is. If there aren't any messages that match the provided
// regular expression or if the regular expression is malformed, this method
// returns an error.
func (m Manager) Message(toIncludeRegex string, p *ir.Person, t time.Time) (*message.HL7Message, error) {
	filtered := m.filterMessages(toIncludeRegex)
	if len(filtered) == 0 {
		return nil, fmt.Errorf("cannot get hardcoded message: no messages match the regular expression %q", toIncludeRegex)
	}
	log.Debugf("Selected %d hardcoded messages based on the regex %q: %v", len(filtered), toIncludeRegex, filtered)

	name := filtered[rand.Intn(len(filtered))]
	log.Infof("Hardcoded message with name %s chosen at random", name)

	msg := m.messages[name]
	return m.buildMessage(msg, p, t)
}

func (m Manager) filterMessages(toIncludeRegex string) []string {
	if toIncludeRegex == "" {
		log.Warning("Ignoring empty regexp while filtering hardcoded messages")
		return nil
	}

	var regexps []*regexp.Regexp
	for _, expStr := range strings.Split(toIncludeRegex, ",") {
		if expStr == "" {
			log.Warning("Ignoring empty regexp while filtering hardcoded messages")
			continue
		}
		exp, err := regexp.Compile(expStr)
		if err != nil {
			log.WithError(err).Warningf("Ignoring invalid regexp %q while filtering hardcoded messages", expStr)
			continue
		}
		regexps = append(regexps, exp)
	}

	if len(regexps) == 0 {
		return nil
	}

	var filtered []string
	for name := range m.messages {
		for _, exp := range regexps {
			if exp.MatchString(name) {
				filtered = append(filtered, name)
				break
			}
		}
	}

	return filtered
}

func (m Manager) buildMessage(msg string, p *ir.Person, t time.Time) (*message.HL7Message, error) {
	pid, err := message.BuildPID(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	msg = strings.Replace(msg, pidSegmentPlaceholder, pid, 1)
	d, err := message.ToHL7Date(t)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot convert %t to HL7 date", t)
	}
	msg = fmt.Sprintf(msg, d, m.generator.NewMessageControlID())

	msgType, err := messageType(msg)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get message type")
	}
	return &message.HL7Message{
		Type:    msgType,
		Message: msg,
	}, nil
}

func messageType(msg string) (*message.Type, error) {
	mo := hl7.NewParseMessageOptions()
	mo.TimezoneLoc = time.UTC
	parsedM, err := hl7.ParseMessageWithOptions([]byte(msg), mo)
	if err != nil {
		return nil, errors.Wrap(err, "message is not parsable")
	}
	msh, err := parsedM.MSH()
	if err != nil {
		return nil, errors.Wrap(err, "cannot get MSH segment from parsed message")
	}
	if msh.MessageType == nil {
		return nil, errors.New("MSH.MessageType is not present in parsed message")
	}
	return &message.Type{
		MessageType:  msh.MessageType.MessageCode.String(),
		TriggerEvent: msh.MessageType.TriggerEvent.String(),
	}, nil
}

func fileExtensionIsValid(fileName string) bool {
	for _, ext := range validExtensions {
		if path.Ext(fileName) == ext {
			return true
		}
	}
	return false
}
