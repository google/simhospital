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

package message

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/pkg/errors"
	"github.com/google/simhospital/pkg/constants"
	"github.com/google/simhospital/pkg/hl7"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/logging"
)

// The fields in this block are HL7 message types Simulated Hospital supports.
// https://corepointhealth.com/resource-center/hl7-resources/hl7-messages/
const (
	// ADT represents an ADT HL7v2 message.
	ADT = "ADT"
	// ORM represents an ORM HL7v2 message.
	ORM = "ORM"
	// ORR represents an ORR HL7v2 message.
	ORR = "ORR"
	// ORU represents an ORU HL7v2 message.
	ORU = "ORU"
	// MDM represents an MDM HL7v2 message.
	MDM = "MDM"
)

// DiagnosticServIDMDOC is the value of the Diagnostic Serv ID field (OBR_24) for clinical documents.
const DiagnosticServIDMDOC = "MDOC"

// SegmentTerminator is the string used to terminate segments in HL7v2 messages.
const SegmentTerminator = constants.SegmentTerminatorStr
const (
	listItemsSeparator           = "~"
	componentSeparator           = "^"
	escapedComponentSeparator    = "\\S\\"
	subComponentSeparator        = "&"
	escapedSubComponentSeparator = "\\T\\"
	lineBreak                    = "\n"
	escapedLineBreak             = "\\.br\\"
	backwardSlash                = "\\"
	escapedBackwardSlash         = "\\E\\"
)

// HL7Message represents a HL7 Message.
type HL7Message struct {
	Type    *Type
	Message string
}

// Type represents the message type for a HL7 Message.
type Type struct {
	MessageType  string
	TriggerEvent string
}

// HeaderInfo contains information relevant to a header of a HL7 Message.
type HeaderInfo struct {
	SendingApplication   string
	SendingFacility      string
	ReceivingApplication string
	ReceivingFacility    string
	// MessageControlID is the MSH -> Message Control ID.
	MessageControlID string
}

var (
	log = logging.ForCallerPackage()

	funcMap = template.FuncMap{
		"HL7_date":     ToHL7Date,
		"HL7_repeated": toHL7RepeatedField,
		"expand_mrns":  expandMRNs,
		"HL7_unit":     toHL7Unit,
		"escape_HL7":   escapeHL7,
	}
)

// ToHL7Date converts a date into a string with HL7 date format.
func ToHL7Date(t ir.Formattable) (string, error) {
	nt, ok := t.(ir.NullTime)
	if ok && !nt.Valid {
		return "", nil
	}
	if nt.Location() != time.UTC {
		// To avoid mistakes, make sure all times are provided in UTC.
		return "", fmt.Errorf("found time with non-UTC location: %v", nt)
	}
	// Get the midnight time at the HL7 Location.
	if nt.Midnight {
		localT := t.In(hl7.Location)
		// Set the location so that the conversion to hl7.Location afterwards is a no-op.
		t = time.Date(localT.Year(), localT.Month(), localT.Day(), 0, 0, 0, 0, hl7.Location)
	}
	return t.In(hl7.Location).Format("20060102150405"), nil
}

// toHL7RepeatedField transforms the given string, where multiple values are separated with \n,
// to multiple HL7v2 values separated by the default multiple item separator.
func toHL7RepeatedField(s string) string {
	return strings.Replace(s, "\n", listItemsSeparator, -1)
}

func expandMRNs(mrns []string) (string, error) {
	fields := make([]string, len(mrns))
	for i, m := range mrns {
		f, err := executeTemplate(parsedCXMRNTemplate, struct {
			MRN string
		}{m})
		if err != nil {
			return "", errors.Wrap(err, "cannot expand MRNs")
		}
		fields[i] = f
	}
	return strings.Join(fields, listItemsSeparator), nil
}

func toHL7Unit(s string) string {
	return strings.Replace(s, componentSeparator, escapedComponentSeparator, -1)
}

func escapeHL7(s string) string {
	r := strings.NewReplacer(
		componentSeparator, escapedComponentSeparator,
		subComponentSeparator, escapedSubComponentSeparator,
		lineBreak, escapedLineBreak,
		backwardSlash, escapedBackwardSlash,
	)
	return r.Replace(s)
}

// Constants for segments and templates.
const (
	MSH             = "MSH"
	MSA             = "MSA"
	EVN             = "EVN"
	PID             = "PID"
	ORC             = "ORC"
	OBR             = "OBR"
	OBRClinicalNote = "OBRClinicalNote"
	OBX             = "OBX"
	OBXClinicalNote = "OBXClinicalNote"
	OBXForMDM       = "OBXForMDM"
	PV1             = "PV1"
	PV2             = "PV2"
	NK1             = "NK1"
	AL1             = "AL1"
	NTE             = "NTE"
	MRG             = "MRG"
	DG1             = "DG1"
	PD1             = "PD1"
	PR1             = "PR1"
	TXA             = "TXA"
)

const (
	locationTemplate      = "LocationTmpl"
	doctorTemplate        = "DoctorTmpl"
	personNameTemplate    = "PersonNameTmpl"
	addressTemplate       = "AddressTmpl"
	homeNumberTemplate    = "HomeNumberTmpl"
	ceTemplate            = "CETmpl"
	ceNoteTemplate        = "CENoteTmpl"
	ceAdmitReasonTemplate = "CEAdmitReasonTmpl"
	cxVisitTemplate       = "CXVisitTmpl"
	cxMRNTemplate         = "CXMRNTmpl"
	primFacTemplate       = "PrimFacTmpl"
	noteTemplate          = "NoteTmpl"
)

var (
	// locationTmpl represents the data type PL: Person Location
	// http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/segment/PV1?version=HL7%20v2.3.1&dataType=PL
	locationTmpl = "{{.Poc}}^{{.Room}}^{{.Bed}}^{{.Facility}}^^{{.LocationType}}^{{.Building}}^{{.Floor}}"

	// doctorTmpl represents the data type XCN: Extended Composite ID Number And Name For Persons
	// http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/segment/PV1?version=HL7%20v2.3.1&dataType=XCN
	doctorTmpl = "{{.ID}}^{{.Surname}}^{{.FirstName}}^^^{{.Prefix}}^^^DRNBR^PRSNL^^^ORGDR"

	// personNameTmpl represents the data type XPN: Extended Person Name
	// http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/segment/PID?version=HL7%20v2.3.1&dataType=XPN
	personNameTmpl = "{{.Surname}}^{{.FirstName}}^{{.MiddleName}}^{{.Suffix}}^{{.Prefix}}^{{.Degree}}^CURRENT"

	// addressTmpl represents the data type XAD: Extended Address
	// http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/segment/PID?version=HL7%20v2.3.1&dataType=XAD
	addressTmpl = "{{.FirstLine}}^{{.SecondLine}}^{{.City}}^^{{.PostalCode}}^{{.Country}}^{{.Type}}"

	// homeNumberTmpl represents the data type XTN: Extended Telecommunication Number
	// http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/segment/PID?version=HL7%20v2.3.1&dataType=XTN
	homeNumberTmpl = "{{.}}^HOME"

	// ceTmpl represents the data type CE: Coded Element
	// http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/segment/PID?version=HL7%20v2.3.1&dataType=CE
	ceTmpl = "{{escape_HL7 .ID}}^{{escape_HL7 .Text}}^{{.CodingSystem}}^^{{escape_HL7 .AlternateText}}"
	// ceNoteTmpl is the CE template for notes.
	// When the OBX.Observation Identifier field is used to send Notes, this is the Document Type; e.g. ECG/Discharge Summary.
	ceNoteTmpl = "{{.DocumentType}}^{{.DocumentType}}"
	// ceAdmitReasonTmpl is CE template for Admit Reason in PV2.3 field.
	ceAdmitReasonTmpl = "^{{.}}"

	// primFacTmpl represents the data type XON: Extended Composite Name And Identification Number For Organizations
	// http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/segment/PD1?version=HL7%20v2.3.1&dataType=XON
	primFacTmpl = "{{.Organization}}^^{{.ID}}"

	// cxVisitTmpl represents the data type CX: Extended Composite ID with Check Digit
	// http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/Default.aspx?version=HL7%20v2.5.1&dataType=CX
	cxVisitTmpl = "{{.}}^^^^visitid"
	// cxMRNTmpl is the template for MRNs.
	cxMRNTmpl = "{{.MRN}}^^^SIMULATOR MRN^MRN"
	// stOBXNoteVal is the template for the OBX.Observation Value for documents.
	stOBXNoteVal = "^^{{.ContentType}}^{{.DocumentEncoding}}^{{escape_HL7 .DocumentContent}}"

	parsedCXMRNTemplate = mustParseTemplateWithoutFuncs(cxMRNTemplate, cxMRNTmpl)
)

var templates = map[string]*template.Template{
	MSH: mustParseTemplate(MSH, "MSH|^~\\&|{{.Header.SendingApplication}}|{{.Header.SendingFacility}}|{{.Header.ReceivingApplication}}|{{.Header.ReceivingFacility}}|{{HL7_date .T}}||{{.MsgType.MessageType}}^{{.MsgType.TriggerEvent}}|{{.Header.MessageControlID}}|T|2.3|||AL||44|ASCII"),
	MSA: mustParseTemplate(MSA, "MSA|AA|{{.OrderMessageControlID}}"),
	EVN: mustParseTemplates(EVN, map[string]string{
		doctorTemplate: doctorTmpl,
		EVN:            `EVN|{{.MsgType.TriggerEvent}}|{{HL7_date .T}}|{{HL7_date .DateTimePlannedEvent}}||{{template "DoctorTmpl" .Operator}}|{{HL7_date .EventOccurredDateTime}}`,
	}),
	PID: mustParseTemplates(PID, map[string]string{
		personNameTemplate: personNameTmpl,
		addressTemplate:    addressTmpl,
		homeNumberTemplate: homeNumberTmpl,
		ceTemplate:         ceTmpl,
		cxMRNTemplate:      cxMRNTmpl,
		PID:                `PID|1|{{template "CXMRNTmpl" .}}|{{template "CXMRNTmpl" .}}~{{.NHS}}^^^NHSNBR^NHSNMBR||{{template "PersonNameTmpl" .}}||{{HL7_date .Birth}}|{{.Gender}}|||{{template "AddressTmpl" .Address}}||{{template "HomeNumberTmpl" .PhoneNumber}}|||||||||{{template "CETmpl" .Ethnicity}}|||||||{{HL7_date .DateOfDeath}}|{{.DeathIndicator}}`,
	}),
	MRG: mustParseTemplate(MRG, "MRG|{{expand_mrns .MRNs}}|"),
	ORC: mustParseTemplate(ORC, "ORC|{{.OrderControl}}|{{.Placer}}|{{.Filler}}||{{.OrderStatus}}||||{{HL7_date .OrderDateTime}}"),
	OBR: mustParseTemplates(OBR, map[string]string{
		ceTemplate:     ceTmpl,
		doctorTemplate: doctorTmpl,
		OBR:            `OBR|1|{{.Placer}}|{{.Filler}}|{{template "CETmpl" .OrderProfile}}||{{HL7_date .OrderDateTime}}|{{HL7_date .CollectedDateTime}}|||||||{{HL7_date .ReceivedInLabDateTime}}|{{.SpecimenSource}}|{{template "DoctorTmpl" .OrderingProvider}}||||||{{HL7_date .ReportedDateTime}}||{{.DiagnosticServID}}|{{.ResultsStatus}}||1`,
	}),
	OBRClinicalNote: mustParseTemplates(OBR, map[string]string{
		ceTemplate:     ceTmpl,
		doctorTemplate: doctorTmpl,
		OBR:            `OBR|1|{{.Placer}}|{{.DocumentID}}|{{template "CETmpl" .OrderProfile}}||{{HL7_date .OrderDateTime}}|{{HL7_date .CollectedDateTime}}|||||||{{HL7_date .ReceivedInLabDateTime}}|{{.SpecimenSource}}|{{template "DoctorTmpl" .OrderingProvider}}||||||{{HL7_date .ReportedDateTime}}||{{.DiagnosticServID}}|{{.ResultsStatus}}||1`,
	}),
	OBX: mustParseTemplates(OBX, map[string]string{
		ceTemplate: ceTmpl,
		OBX:        `OBX|{{.ID}}|{{.ValueType}}|{{template "CETmpl" .TestName}}||{{HL7_repeated .Value}}|{{HL7_unit .Unit}}|{{escape_HL7 .Range}}|{{.AbnormalFlag}}|||{{.Status}}|||{{HL7_date .ObservationDateTime}}||`,
	}),
	OBXClinicalNote: mustParseTemplates(OBX, map[string]string{
		ceNoteTemplate: ceNoteTmpl,
		noteTemplate:   stOBXNoteVal,
		doctorTemplate: doctorTmpl,
		OBX:            `OBX|{{.ID}}|{{.ValueType}}|{{template "CENoteTmpl" .ClinicalNote}}||{{template "NoteTmpl" .Content}}|||||||||{{HL7_date .ObservationDateTime}}||{{template "DoctorTmpl" .OrderingProvider}}`,
	}),
	OBXForMDM: mustParseTemplates(OBX, map[string]string{
		ceTemplate: ceTmpl,
		OBX:        `OBX|{{.ID}}|TX|{{template "CETmpl" .ObservationIdentifier}}|1|{{.Content}}||||||F||||||`,
	}),
	PV1: mustParseTemplates(PV1, map[string]string{
		locationTemplate: locationTmpl,
		doctorTemplate:   doctorTmpl,
		cxVisitTemplate:  cxVisitTmpl,
		PV1:              `PV1|1|{{.Class}}|{{template "LocationTmpl" .Location}}|28b||{{template "LocationTmpl" .PriorLocation}}|{{template "DoctorTmpl" .AttendingDoctor}}|||{{.HospitalService}}|{{template "LocationTmpl" .TemporaryLocation}}|||||||{{.Type}}|{{template "CXVisitTmpl" .VisitID}}||||||||||||||||||||||{{.AccountStatus}}|{{template "LocationTmpl" .PendingLocation}}|{{template "LocationTmpl" .PriorTemporaryLocation}}|{{HL7_date .AdmissionDate}}|{{HL7_date .DischargeDate}}|`,
	}),
	PV2: mustParseTemplates(PV2, map[string]string{
		locationTemplate:      locationTmpl,
		ceAdmitReasonTemplate: ceAdmitReasonTmpl,
		PV2:                   `PV2|{{template "LocationTmpl" .PriorPendingLocation}}||{{template "CEAdmitReasonTmpl" .AdmitReason}}|||||{{HL7_date .ExpectedAdmitDateTime}}|{{HL7_date .ExpectedDischargeDateTime}}`,
	}),
	NK1: mustParseTemplates(NK1, map[string]string{
		personNameTemplate: personNameTmpl,
		addressTemplate:    addressTmpl,
		homeNumberTemplate: homeNumberTmpl,
		ceTemplate:         ceTmpl,
		NK1:                `NK1|{{.ID}}|{{template "PersonNameTmpl" .}}|{{template "CETmpl" .Relationship}}|{{template "AddressTmpl" .Address}}|{{template "HomeNumberTmpl" .PhoneNumber}}||{{template "CETmpl" .ContactRole}}||||||||{{.Gender}}|`,
	}),
	AL1: mustParseTemplates(AL1, map[string]string{
		ceTemplate: ceTmpl,
		AL1:        `AL1|{{.ID}}|{{.Type}}|{{template "CETmpl" .Description}}|{{.Severity}}|{{.Reaction}}|{{HL7_date .IdentificationDateTime}}`,
	}),
	NTE: mustParseTemplate(NTE, `NTE|{{.ID}}||{{.Note}}|`),
	DG1: mustParseTemplates(DG1, map[string]string{
		ceTemplate:     ceTmpl,
		doctorTemplate: doctorTmpl,
		DG1:            `DG1|{{.ID}}|SNMCT|{{template "CETmpl" .Description}}|{{.Description.Text}}|{{HL7_date .DateTime}}|{{.Type}}|||||||||0|{{template "DoctorTmpl" .Clinician}}`,
	}),
	PD1: mustParseTemplates(PD1, map[string]string{
		primFacTemplate: primFacTmpl,
		PD1:             `PD1|||{{template "PrimFacTmpl" .PrimaryFacility}}|`,
	}),
	PR1: mustParseTemplates(PR1, map[string]string{
		ceTemplate:     ceTmpl,
		doctorTemplate: doctorTmpl,
		PR1:            `PR1|{{.ID}}|SNMCT|{{template "CETmpl" .Description}}|{{.Description.Text}}|{{HL7_date .DateTime}}|{{.Type}}||||||{{template "DoctorTmpl" .Clinician}}||0||`,
	}),
	TXA: mustParseTemplates(TXA, map[string]string{
		doctorTemplate: doctorTmpl,
		TXA:            `TXA|1|{{.DocumentType}}||{{HL7_date .ActivityDateTime}}|{{template "DoctorTmpl" .AttendingDoctor}}|||{{HL7_date .EditDateTime}}||||{{.UniqueDocumentNumber}}|||||{{.DocumentCompletionStatus}}||||||`,
	}),
}

// BuildDocumentNotificationMDMT02 builds and returns a HL7 MDM^T02 message.
func BuildDocumentNotificationMDMT02(h *HeaderInfo, p *ir.PatientInfo, d *ir.Document, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  MDM,
		TriggerEvent: "T02",
	}

	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)
	txa, err := BuildTXA(p, d)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build TXA segment")
	}
	segments = append(segments, txa)
	for id, note := range d.ContentLine {
		obx, err := BuildOBXForMDM(id+1, d.ObservationIdentifier, note)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build OBX segment")
		}
		segments = append(segments, obx)
	}

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildResultORUR01 builds and returns a HL7 ORU^R01 message.
func BuildResultORUR01(h *HeaderInfo, p *ir.PatientInfo, o *ir.Order, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ORU,
		TriggerEvent: "R01",
	}

	segments, err := segmentsORU(h, p, o, msgTime, msgType)
	if err != nil {
		return nil, err
	}

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildResultORUR03 builds and returns a HL7 ORU^R03 message.
func BuildResultORUR03(h *HeaderInfo, p *ir.PatientInfo, o *ir.Order, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ORU,
		TriggerEvent: "R03",
	}

	segments, err := segmentsORU(h, p, o, msgTime, msgType)
	if err != nil {
		return nil, err
	}

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildResultORUR32 builds and returns a HL7 ORU^R32 message.
func BuildResultORUR32(h *HeaderInfo, p *ir.PatientInfo, o *ir.Order, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ORU,
		TriggerEvent: "R32",
	}

	segments, err := segmentsORU(h, p, o, msgTime, msgType)
	if err != nil {
		return nil, err
	}

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

func segmentsORU(h *HeaderInfo, p *ir.PatientInfo, o *ir.Order, msgTime time.Time, msgType *Type) ([]string, error) {
	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)
	orc, err := BuildORC(o)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build ORC segment")
	}
	segments = append(segments, orc)
	obr, err := BuildOBR(o)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build OBR segment")
	}
	segments = append(segments, obr)

	if o.DiagnosticServID == DiagnosticServIDMDOC {
		return clinicalNotesOBX(o, segments)
	}
	return resultsOBX(o, segments)
}

func clinicalNotesOBX(o *ir.Order, segments []string) ([]string, error) {
	for _, result := range o.Results {
		for id := range result.ClinicalNote.Contents {
			obx, err := BuildOBXForClinicalNote(id+1, id, result, o)
			if err != nil {
				return nil, errors.Wrap(err, "cannot build OBX segment")
			}
			segments = append(segments, obx)
		}
	}
	return segments, nil
}

func resultsOBX(o *ir.Order, segments []string) ([]string, error) {
	for id, result := range o.Results {
		// We increment by 1 so that the first OBX has a SetID of 1 - that's how segment numbers starts.
		// We use the number of previous result for the same order so that the SetIDs of OBX segments
		// of different messages related to the same order (i.e. amendments) don't clash with the previous messages.
		obx, err := BuildOBX(o.NumberOfPreviousResults+id+1, result, o)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build OBX segment")
		}
		segments = append(segments, obx)
		for noteID, note := range result.Notes {
			nte, err := BuildNTE(noteID, note)
			if err != nil {
				return nil, errors.Wrap(err, "cannot build NTE segment")
			}
			segments = append(segments, nte)
		}
	}
	return segments, nil
}

// BuildOrderORMO01 builds and returns a HL7 ORM^O01 message.
func BuildOrderORMO01(h *HeaderInfo, p *ir.PatientInfo, o *ir.Order, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ORM,
		TriggerEvent: "O01",
	}

	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)
	orc, err := BuildORC(o)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build ORC segment")
	}
	segments = append(segments, orc)
	obr, err := BuildOBR(o)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build OBR segment")
	}
	segments = append(segments, obr)
	for noteID, note := range o.NotesForORM {
		nte, err := BuildNTE(noteID, note)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build NTE segment")
		}
		segments = append(segments, nte)
	}

	for id, result := range o.ResultsForORM {
		obx, err := BuildOBX(id+1, result, o)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build OBX segment")
		}
		segments = append(segments, obx)
		for noteID, note := range result.Notes {
			nte, err := BuildNTE(noteID, note)
			if err != nil {
				return nil, errors.Wrap(err, "cannot build NTE segment")
			}
			segments = append(segments, nte)
		}
	}
	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildPathologyORRO02 builds and returns a HL7 ORR^O02 message.
func BuildPathologyORRO02(h *HeaderInfo, p *ir.PatientInfo, o *ir.Order, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ORR,
		TriggerEvent: "O02",
	}
	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	msa, err := BuildMSA(o.MessageControlIDOriginalOrder)
	if err != nil {
		return nil, errors.Wrap(err, "MSA build MSH segment")
	}
	segments = append(segments, msa)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	orc, err := BuildORC(o)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build ORC segment")
	}
	segments = append(segments, orc)

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildAdmissionADTA01 builds and returns a HL7 ADT^A01 message.
func BuildAdmissionADTA01(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A01",
	}

	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)
	pv2, err := BuildPV2(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV2 segment")
	}
	segments = append(segments, pv2)
	for id, ap := range p.AssociatedParties {
		nk1, err := BuildNK1(id, ap)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build NK1 segment")
		}
		segments = append(segments, nk1)
	}
	for id, al := range p.Allergies {
		al1, err := BuildAL1(id, al)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build AL1 segment")
		}
		segments = append(segments, al1)
	}

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildTransferADTA02 builds and returns a HL7 ADT^A02 message.
func BuildTransferADTA02(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A02",
	}

	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildDischargeADTA03 builds and returns a HL7 ADT^A03 message.
func BuildDischargeADTA03(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A03",
	}

	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)
	for id, al := range p.Allergies {
		al1, err := BuildAL1(id, al)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build AL1 segment")
		}
		segments = append(segments, al1)
	}

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildRegistrationADTA04 builds and returns a HL7 ADT^A04 message.
func BuildRegistrationADTA04(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A04",
	}

	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)
	for id, ap := range p.AssociatedParties {
		nk1, err := BuildNK1(id, ap)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build NK1 segment")
		}
		segments = append(segments, nk1)
	}
	for id, al := range p.Allergies {
		al1, err := BuildAL1(id, al)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build AL1 segment")
		}
		segments = append(segments, al1)
	}

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildPreAdmitADTA05 builds and returns a HL7 ADT^A05 message.
func BuildPreAdmitADTA05(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A05",
	}
	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, p.ExpectedAdmitDateTime, p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)
	pv2, err := BuildPV2(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV2 segment")
	}
	segments = append(segments, pv2)
	for id, al := range p.Allergies {
		al1, err := BuildAL1(id, al)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build AL1 segment")
		}
		segments = append(segments, al1)
	}
	for id, ap := range p.AssociatedParties {
		nk1, err := BuildNK1(id, ap)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build NK1 segment")
		}
		segments = append(segments, nk1)
	}
	for id, d := range p.Diagnoses {
		dg1, err := BuildDG1(id, d)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build DG1 segment")
		}
		segments = append(segments, dg1)
	}
	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildUpdatePatientADTA08 builds and returns a HL7 ADT^A08 message.
func BuildUpdatePatientADTA08(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A08",
	}

	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	segments = append(segments, BuildPseudoPV1())
	for id, al := range p.Allergies {
		al1, err := BuildAL1(id, al)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build AL1 segment")
		}
		segments = append(segments, al1)
	}
	for id, d := range p.Diagnoses {
		dg1, err := BuildDG1(id, d)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build DG1 segment")
		}
		segments = append(segments, dg1)
	}
	for id, p := range p.Procedures {
		pr1, err := BuildPR1(id, p)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build PR1 segment")
		}
		segments = append(segments, pr1)
	}

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildTrackDepartureADTA09 builds and returns a HL7 ADT^A09 message.
func BuildTrackDepartureADTA09(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A09",
	}

	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildTrackArrivalADTA10 builds and returns a HL7 ADT^A10 message.
func BuildTrackArrivalADTA10(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A10",
	}

	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildCancelVisitADTA11 builds and returns a HL7 ADT^A11 message.
func BuildCancelVisitADTA11(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A11",
	}

	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, p.AdmissionDate)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildBedSwapADTA17 builds and returns a HL7 ADT^A17 message.
func BuildBedSwapADTA17(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time, otherP *ir.PatientInfo) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A17",
	}

	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)
	otherPID, err := BuildPID(otherP.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, otherPID)
	segments = append(segments, pd1)
	otherPV1, err := BuildPV1(otherP)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, otherPV1)

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildAddPersonADTA28 builds and returns a HL7 ADT^A28 message.
func BuildAddPersonADTA28(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A28",
	}

	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	segments = append(segments, BuildPseudoPV1())
	for id, al := range p.Allergies {
		al1, err := BuildAL1(id, al)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build AL1 segment")
		}
		segments = append(segments, al1)
	}

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildUpdatePersonADTA31 builds and returns a HL7 ADT^A31 message.
func BuildUpdatePersonADTA31(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A31",
	}
	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	segments = append(segments, BuildPseudoPV1())
	for id, al := range p.Allergies {
		al1, err := BuildAL1(id, al)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build AL1 segment")
		}
		segments = append(segments, al1)
	}
	for id, d := range p.Diagnoses {
		dg1, err := BuildDG1(id, d)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build DG1 segment")
		}
		segments = append(segments, dg1)
	}
	for id, p := range p.Procedures {
		pr1, err := BuildPR1(id, p)
		if err != nil {
			return nil, errors.Wrap(err, "cannot build PR1 segment")
		}
		segments = append(segments, pr1)
	}

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildCancelTransferADTA12 builds and returns a HL7 ADT^A12 message.
func BuildCancelTransferADTA12(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A12",
	}

	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, p.TransferDate)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildCancelDischargeADTA13 builds and returns a HL7 ADT^A13 message.
func BuildCancelDischargeADTA13(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A13",
	}

	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, p.DischargeDate)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildPendingAdmissionADTA14 builds and returns a HL7 ADT^A14 message.
func BuildPendingAdmissionADTA14(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A14",
	}
	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	// The PV2 segment contains ExpectedAdmitDateTime as well, which is the recommendation.
	// http://www.hl7.eu/refactored/segEVN.html
	// We add it in the EVN as well for consistency with the PendingTransfer message that doesn't have
	// an equivalent in PV2.
	evn, err := BuildEVN(eventTime, msgType, p.ExpectedAdmitDateTime, p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)
	pv2, err := BuildPV2(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV2 segment")
	}
	segments = append(segments, pv2)

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildPendingTransferADTA15 builds and returns a HL7 ADT^A15 message.
func BuildPendingTransferADTA15(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A15",
	}
	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, p.ExpectedTransferDateTime, p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildPendingDischargeADTA16 builds and returns a HL7 ADT^A16 message.
func BuildPendingDischargeADTA16(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A16",
	}
	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	// See BuildPendingAdmissionADTA14 for why we send ExpectedDischargeDateTime here.
	evn, err := BuildEVN(eventTime, msgType, p.ExpectedDischargeDateTime, p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)
	pv2, err := BuildPV2(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV2 segment")
	}
	segments = append(segments, pv2)

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildDeleteVisitADTA23 builds and returns a HL7 ADT^A23 message.
func BuildDeleteVisitADTA23(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A23",
	}
	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)
	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildCancelPendingDischargeADTA25 builds and returns a HL7 ADT^A25 message.
func BuildCancelPendingDischargeADTA25(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A25",
	}
	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, p.ExpectedDischargeDateTime)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)
	pv2, err := BuildPV2(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV2 segment")
	}
	segments = append(segments, pv2)

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildCancelPendingTransferADTA26 builds and returns a HL7 ADT^A26 message.
func BuildCancelPendingTransferADTA26(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A26",
	}
	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, p.ExpectedTransferDateTime)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)
	pv2, err := BuildPV2(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV2 segment")
	}
	segments = append(segments, pv2)

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildCancelPendingAdmitADTA27 builds and returns a HL7 ADT^A27 message.
func BuildCancelPendingAdmitADTA27(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A27",
	}
	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, p.ExpectedAdmitDateTime)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)
	pv2, err := BuildPV2(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV2 segment")
	}
	segments = append(segments, pv2)

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildMergeADTA34 builds and returns a HL7 ADT^A34 message.
func BuildMergeADTA34(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time, withMRN string) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A34",
	}

	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	mrg, err := BuildMRG([]string{withMRN})
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MRG segment")
	}
	segments = append(segments, mrg)

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildMergeADTA40 builds and returns a HL7 ADT^A40 message.
func BuildMergeADTA40(h *HeaderInfo, p *ir.PatientInfo, eventTime time.Time, msgTime time.Time, withMRN []string) (*HL7Message, error) {
	msgType := &Type{
		MessageType:  ADT,
		TriggerEvent: "A40",
	}

	var segments []string
	msh, err := BuildMSH(msgTime, msgType, h)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MSH segment")
	}
	segments = append(segments, msh)
	evn, err := BuildEVN(eventTime, msgType, ir.NewInvalidTime(), p.AttendingDoctor, ir.NewInvalidTime())
	if err != nil {
		return nil, errors.Wrap(err, "cannot build EVN segment")
	}
	segments = append(segments, evn)
	pid, err := BuildPID(p.Person)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PID segment")
	}
	segments = append(segments, pid)
	pd1, err := BuildPD1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PD1 segment")
	}
	segments = append(segments, pd1)
	mrg, err := BuildMRG(withMRN)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build MRG segment")
	}
	segments = append(segments, mrg)
	pv1, err := BuildPV1(p)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build PV1 segment")
	}
	segments = append(segments, pv1)

	return &HL7Message{
		Type:    msgType,
		Message: strings.Join(segments, SegmentTerminator),
	}, nil
}

// BuildMSH builds and returns a HL7 MSH segment.
func BuildMSH(t time.Time, messageType *Type, header *HeaderInfo) (string, error) {
	return executeTemplate(templates[MSH], struct {
		T       *time.Time
		MsgType *Type
		Header  *HeaderInfo
	}{&t, messageType, header})
}

// BuildMSA builds and returns a HL7 MSA segment.
func BuildMSA(orderMessageControlID string) (string, error) {
	return executeTemplate(templates[MSA], struct {
		OrderMessageControlID string
	}{OrderMessageControlID: orderMessageControlID})
}

// BuildEVN builds and returns a HL7 EVN segment.
func BuildEVN(t time.Time, messageType *Type, planned ir.NullTime, operator *ir.Doctor, occurred ir.NullTime) (string, error) {
	return executeTemplate(templates[EVN], struct {
		T                     *time.Time
		MsgType               *Type
		DateTimePlannedEvent  ir.NullTime
		Operator              *ir.Doctor
		EventOccurredDateTime ir.NullTime
	}{&t, messageType, planned, operator, occurred})
}

// BuildPID builds and returns a HL7 PID segment.
func BuildPID(p *ir.Person) (string, error) {
	return executeTemplate(templates[PID], p)
}

// BuildPV1 builds and returns a HL7 PV1 segment.
func BuildPV1(p *ir.PatientInfo) (string, error) {
	return executeTemplate(templates[PV1], p)
}

// BuildPseudoPV1 builds and returns a HL7 PV1 segment without any patient information.
// A PV1 that some messages need to send for backwards compatibility but where the visit is not
// relevant to the message, e.g. ADT^08. The PatientClass is set to N - Not applicable.
func BuildPseudoPV1() string {
	return `PV1|1|N|`
}

// BuildPV2 builds and returns a HL7 PV2 segment.
func BuildPV2(p *ir.PatientInfo) (string, error) {
	return executeTemplate(templates[PV2], p)
}

// BuildNK1 builds and returns a HL7 NK1 segment.
func BuildNK1(id int, p *ir.AssociatedParty) (string, error) {
	return executeTemplate(templates[NK1], struct {
		*ir.AssociatedParty
		ID int
	}{p, id})
}

// BuildAL1 builds and returns a HL7 AL1 segment.
func BuildAL1(id int, a *ir.Allergy) (string, error) {
	return executeTemplate(templates[AL1], struct {
		*ir.Allergy
		ID int
	}{a, id})
}

// BuildORC builds and returns a HL7 ORC segment.
func BuildORC(o *ir.Order) (string, error) {
	return executeTemplate(templates[ORC], &o)
}

// BuildOBR builds and returns a HL7 OBR segment.
func BuildOBR(o *ir.Order) (string, error) {
	// If this order is sending a ClinicalNote, use the appropriate OBR template.
	var key, documentID string
	if o.DiagnosticServID == DiagnosticServIDMDOC {
		key = OBRClinicalNote
		documentID = o.Results[0].ClinicalNote.DocumentID
	} else {
		key = OBR
	}
	return executeTemplate(templates[key], struct {
		*ir.Order
		DocumentID string
	}{o, documentID})
}

// BuildOBX builds and returns a HL7 OBX segment.
func BuildOBX(id int, r *ir.Result, o *ir.Order) (string, error) {
	return executeTemplate(templates[OBX], struct {
		*ir.Result
		ID                  int
		ObservationDateTime ir.NullTime
		OrderingProvider    *ir.Doctor
	}{r, id, r.ObservationDateTime, o.OrderingProvider})
}

// BuildOBXForClinicalNote build and returns a HL7 OBX segment for a Clinical Note.
func BuildOBXForClinicalNote(id, contentIndex int, r *ir.Result, o *ir.Order) (string, error) {
	return executeTemplate(templates[OBXClinicalNote], struct {
		*ir.Result
		ID                  int
		Content             *ir.ClinicalNoteContent
		ObservationDateTime ir.NullTime
		DiagnosticServID    string
		OrderingProvider    *ir.Doctor
	}{r, id, r.ClinicalNote.Contents[contentIndex], r.ObservationDateTime, o.DiagnosticServID, o.OrderingProvider})
}

// BuildOBXForMDM builds and returns a HL7 OBX segment for MDMT02 type for an MDM message.
func BuildOBXForMDM(id int, o *ir.CodedElement, line string) (string, error) {
	return executeTemplate(templates[OBXForMDM], struct {
		ID                    int
		ObservationIdentifier *ir.CodedElement
		Content               string
	}{id, o, line})
}

// BuildNTE builds and returns a HL7 NTE segment.
func BuildNTE(id int, note string) (string, error) {
	return executeTemplate(templates[NTE], struct {
		Note string
		ID   int
	}{note, id})
}

// BuildPD1 builds and returns a HL7 PD1 segment.
func BuildPD1(p *ir.PatientInfo) (string, error) {
	return executeTemplate(templates[PD1], struct {
		*ir.PrimaryFacility
	}{p.PrimaryFacility})
}

// BuildMRG builds and returns a HL7 MRG segment.
func BuildMRG(mrns []string) (string, error) {
	return executeTemplate(templates[MRG], struct {
		MRNs []string
	}{mrns})
}

// BuildDG1 builds and returns a HL7 DG1 segment.
func BuildDG1(id int, diagnose *ir.DiagnosisOrProcedure) (string, error) {
	return executeTemplate(templates[DG1], struct {
		*ir.DiagnosisOrProcedure
		ID int
	}{DiagnosisOrProcedure: diagnose, ID: id})
}

// BuildPR1 builds and returns a HL7 PR1 segment.
func BuildPR1(id int, procedure *ir.DiagnosisOrProcedure) (string, error) {
	return executeTemplate(templates[PR1], struct {
		*ir.DiagnosisOrProcedure
		ID int
	}{DiagnosisOrProcedure: procedure, ID: id})
}

// BuildTXA builds and returns a HL7 TXA segment.
func BuildTXA(p *ir.PatientInfo, d *ir.Document) (string, error) {
	return executeTemplate(templates[TXA], struct {
		*ir.Document
		AttendingDoctor *ir.Doctor
	}{d, p.AttendingDoctor})
}

func mustParseTemplate(name string, t string) *template.Template {
	tmpl, err := template.New(name).Funcs(funcMap).Parse(t)
	if err != nil {
		log.WithError(err).Fatalf("Cannot parse template: %s", name)
	}
	return tmpl
}

func mustParseTemplateWithoutFuncs(name string, t string) *template.Template {
	tmpl, err := template.New(name).Parse(t)
	if err != nil {
		log.WithError(err).Fatalf("Cannot parse template: %s", name)
	}
	return tmpl
}

func mustParseTemplates(name string, templates map[string]string) *template.Template {
	tmpl := template.New(name).Funcs(funcMap)
	var err error

	for name, t := range templates {
		// Define sub-templates that can be referenced by their names.
		// Only execute sub-templates if the element passed to it is neither nil nor empty
		// (0 / false / slice, map or string of length 0).
		tmpl, err = tmpl.Parse(fmt.Sprintf(`{{define "%s"}}{{if .}}%s{{end}}{{end}}`, name, t))
		if err != nil {
			log.WithError(err).Fatalf("Cannot parse template: %s", name)
		}
	}
	return tmpl
}

func executeTemplate(tmpl *template.Template, data interface{}) (string, error) {
	var buffer bytes.Buffer
	err := tmpl.Execute(&buffer, data)
	if err != nil {
		return "", errors.Wrapf(err, "cannot execute the template: %s", tmpl.Name())
	}
	return buffer.String(), nil
}

func (m Type) String() string {
	return fmt.Sprintf("%s^%s", m.MessageType, m.TriggerEvent)
}

func (m HL7Message) String() string {
	return fmt.Sprintf("message:[type:%v msg:%v]",
		m.Type, strings.Replace(m.Message, SegmentTerminator, " ", -1))
}
