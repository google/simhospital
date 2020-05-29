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

// Package starter contains functionality to start pathways by an endpoint.
package starter

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/google/simhospital/pkg/hl7"
	"github.com/google/simhospital/pkg/hospital"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/logging"
	"github.com/google/simhospital/pkg/pathway"
)

const (
	// UnavailableResponse is the default response for the PathwayStarter.
	UnavailableResponse = ""
	// DefaultPathwayNameFromUI is the default pathway name to be used for pathways that are started without a name.
	DefaultPathwayNameFromUI = "pathway_started_from_ui"
	// StartErrPrefix is the prefix for the error message in the response when something went wrong.
	StartErrPrefix = "Cannot start pathway or send messages"
)

var (
	pathwayNameRegex = regexp.MustCompile("^[a-zA-Z0-9_]+$")
	// pathwayAndMRNRegex matches a pathway definition of the form: pathway_name: MRN
	pathwayAndMRNRegex = regexp.MustCompile("^([a-zA-Z0-9_]+): ([a-zA-Z0-9]+)$")
	// pathwayAndPatientRegex matches a pathway definition of the form: pathway_name: FirstName LastName
	pathwayAndPatientRegex = regexp.MustCompile("^([a-zA-Z0-9_]+): ([a-zA-Z-]+) ([a-zA-Z-]+)$")
	log                    = logging.ForCallerPackage()
)

// PathwayStarter provides a handler for starting pathways on demand.
type PathwayStarter struct {
	Hospital       *hospital.Hospital
	Parser         *pathway.Parser
	PathwayManager pathway.Manager
	Sender         hl7.Sender
	response       string
}

// ServeHTTP handles the requests to start specific pathway.
// Use a POST request to start a pathway or send HL7v2 messages.
// The request can have the following formats:
// * A pathway name, e.g. "aki_scenario_1"
// * A pathway name plus the MRN of a patient, e.g. "aki_scenario_1: 1234567890"
// * A pathway name plus the first and last name of a patient, e.g. "aki_scenario_1: Tina Turner"
// * A pathway in YML format as in the pathway config files
// * One HL7v2 message where every segment is in a different line (separated by \n)
// * Multiple HL7v2 messages where messages are separated by blank lines (separated by \n\n)
// Use a GET request to get the latest response that mentions whether the request was successful.
func (ps *PathwayStarter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Return the response if available.
		if ps.response != UnavailableResponse {
			w.Write([]byte(ps.response))
			ps.response = UnavailableResponse
		}
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			ps.response = "The pathway definition was not read successfully."
			log.WithError(err).Warning("The pathway definition was not read successfully")
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		resp, err := ps.requestPathwayOrSendMessage(string(body))
		if err != nil {
			log.WithError(err).Error("Cannot process PathwayStarter request")
			ps.response = fmt.Sprintf("%s:\n%v", StartErrPrefix, err.Error())
			return
		}
		log.Infof("ServeHTTP response: %v", resp)
		// The UI expects lines separated by \n.
		ps.response = strings.Join(resp, "\n")
	case "PUT":
		http.Error(w, "Method not implemented",
			http.StatusInternalServerError)
	case "DELETE":
		http.Error(w, "Method not implemented",
			http.StatusInternalServerError)
	default:
		http.Error(w, "Error",
			http.StatusInternalServerError)
	}
}

// parseBytesAsHL7 parses the given bytes as a HL7 message.
// parseBytesAsHL7 returns an error if the message is not parsable.
func parseBytesAsHL7(b []byte) error {
	_, err := hl7.ParseMessage(b)
	return errors.Wrap(err, "cannot parse message")
}

// sendMessages sends the HL7 messages contained in the given text.
// The format of the text needs to be as follows: (1) the text needs to start with "MSH", and
// (2) every message is separated by the previous one by one or more empty lines.
// Before sending a message, sendMessages checks that the message is a valid HL7 message.
// sendMessages returns the list of strings that describe the outcome to be set as the response, or an error.
// Invalid messages and problems sending messages return errors.
func (ps *PathwayStarter) sendMessages(text string) ([]string, error) {
	// The error is capitalized because it is displayed in the UI.
	errStr := "Error sending messages: %v. Number of messages successfully sent: %d"
	count := 0
	for _, section := range strings.Split(text, "\n\n") {
		if section == "" {
			continue
		}
		m := strings.Replace(section, "\n", hl7.SegmentTerminatorStr, -1)
		b := []byte(m)
		if err := parseBytesAsHL7(b); err != nil {
			return nil, fmt.Errorf(errStr, err, count)
		}
		if err := ps.Sender.Send(b); err != nil {
			return nil, fmt.Errorf(errStr, err, count)
		}
		count++
	}
	return []string{fmt.Sprintf("Number of messages sent: %d", count)}, nil
}

// requestPathwayOrSendMessage attempts to run the pathway or send a message as follows:
// 1. If the text starts with MSH, send as messages.
// 2. Attempt to run the pathway using the string as a name. If it fails,
// 3. Attempt to run the pathway using the string as: "pathway_name: MRN". If it fails,
// 4. Attempt to run the pathway using the string as: "pathway_name: FirstName Surname". If it fails,
// 5. Attempt to run the pathway using the string as a yml format definition. If it fails,
// 6. Attempt to run the pathway using the string as a yml format definition that includes the
//    pathway name, i.e.,
//    pathway_name:
//      person:
//        age: ...
// It returns the list of strings that describe the outcome to be set as the response, or an error.
func (ps *PathwayStarter) requestPathwayOrSendMessage(text string) ([]string, error) {
	if strings.HasPrefix(text, "MSH") {
		return ps.sendMessages(text)
	}
	if pathwayNameRegex.MatchString(text) {
		return ps.startPathwayWithPatient(text, nil)
	}
	if groups := pathwayAndMRNRegex.FindStringSubmatch(text); len(groups) == 3 {
		pathwayName, mrn := groups[1], groups[2]
		return ps.startPathwayWithMRN(pathwayName, mrn)
	}
	if groups := pathwayAndPatientRegex.FindStringSubmatch(text); len(groups) == 4 {
		pathwayName, firstName, lastName := groups[1], groups[2], groups[3]
		return ps.startPathwayWithPatient(pathwayName, &[2]string{firstName, lastName})
	}
	p, err := ps.Parser.ParseSinglePathway([]byte(text))
	if err != nil {
		return nil, fmt.Errorf("cannot parse pathway: %v", err)
	}
	if p.Name() == pathway.UnknownPathwayName {
		p.UpdateName(DefaultPathwayNameFromUI)
	}

	return ps.startPathwayInHospital(&p)
}

// startPathwayWithMRN starts the given pathway name for a patient with the given MRN.
// If an MRN is not provided, this method is equivalent to startPathwayWithPatient with a nil patient.
// startPathwayWithMRN returns the list of strings that describe the outcome to be set as the response, or an error.
func (ps *PathwayStarter) startPathwayWithMRN(pathwayName string, mrn string) ([]string, error) {
	if mrn == "" {
		return ps.startPathwayWithPatient(pathwayName, nil)
	}
	logLocal := log.WithField("pathway_name", pathwayName).WithField("mrn", mrn)

	logLocal.Info("Attempting to run pathway by name")
	p, err := ps.PathwayManager.GetPathway(pathwayName)
	if err != nil {
		return nil, fmt.Errorf("pathway %q is not defined", pathwayName)
	}

	if !p.Persons.HasOnePerson() {
		return nil, fmt.Errorf("pathway %q cannot be started with a single patient as it refers to more than one patient", pathwayName)
	}
	// If the MRN belongs to an existing patient, we set the MRN only. This is to avoid updating an existing
	// patient with predefined fields from the Person section of the pathway (e.g., the patient name).
	if !ps.Hospital.PatientExists(mrn) {
		updateMRN(p, mrn)
	} else if err := setMRNOnly(p, mrn); err != nil {
		return nil, errors.Wrapf(err, "cannot set MRN to the person from the pathway %q", pathwayName)
	}
	return ps.startPathwayInHospital(p)
}

func setMRNOnly(p *pathway.Pathway, mrn string) error {
	id, _, err := p.Persons.OnlyPerson()
	if err != nil {
		return errors.Wrap(err, "cannot get the only person from the pathway")
	}
	p.Persons = &pathway.Persons{
		id: pathway.Person{
			MRN: mrn,
		},
	}
	return nil
}

func updateMRN(p *pathway.Pathway, mrn string) {
	persons := *p.Persons
	for k, person := range persons {
		person.MRN = mrn
		persons[k] = person
	}
}

// startPathwayWithPatient starts the given pathway name using the given patient (first and last names) if present.
// startPathwayWithPatient returns the list of strings that describe the outcome to be set as the response, or an error.
func (ps *PathwayStarter) startPathwayWithPatient(pathwayName string, patient *[2]string) ([]string, error) {
	logLocal := log.WithField("pathway_name", pathwayName)
	if patient != nil {
		logLocal = logLocal.WithField("first_name", patient[0]).
			WithField("last_name", patient[1])
	}

	logLocal.Info("Attempting to run pathway by name")
	p, err := ps.PathwayManager.GetPathway(pathwayName)
	if err != nil {
		return nil, fmt.Errorf("pathway %q is not defined", pathwayName)
	}
	if patient == nil {
		return ps.startPathwayInHospital(p)
	}

	if !p.Persons.HasOnePerson() {
		return nil, fmt.Errorf("pathway %q cannot be started with a single patient as it refers to more than one patient", pathwayName)
	}
	persons := *p.Persons
	for k, person := range persons {
		person.FirstName = patient[0]
		person.Surname = pathway.OptionalRandomString(patient[1])
		persons[k] = person
	}
	return ps.startPathwayInHospital(p)
}

// startPathwayInHospital attempts to start the given pathway and returns the list of strings that describe the outcome
// to be set as the response, or an error.
func (ps *PathwayStarter) startPathwayInHospital(p *pathway.Pathway) ([]string, error) {
	persons, err := ps.Hospital.StartPathway(p)
	if err != nil {
		return nil, fmt.Errorf("cannot start pathway %v: %v", p.Name(), err)
	}
	return ps.pathwayStartedResponse(p.Name(), persons), nil
}

// pathwayStartedResponse returns a pathway started response from the given persons' names and MRNs.
func (ps *PathwayStarter) pathwayStartedResponse(pathwayName string, persons []*ir.Person) []string {
	allLines := []string{fmt.Sprintf("Pathway %s started with patient(s):", pathwayName)}
	for _, person := range persons {
		allLines = append(allLines, fmt.Sprintf("%s %s (MRN %s)", person.FirstName, person.Surname, person.MRN))
	}
	return allLines
}
