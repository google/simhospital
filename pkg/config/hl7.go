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

package config

import (
	"context"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"github.com/google/simhospital/pkg/files"
)

// HL7Config is the configuration for HL7 messages where the values are important for HL7 compliance or specific uses of the HL7 standard.
type HL7Config struct {
	Allergy HL7Allergy

	Diagnosis HL7Diagnosis

	Document HL7Document

	Procedure HL7Procedure

	OrderControl OrderControl `yaml:"order_control"`

	ResultStatus ResultStatus `yaml:"result_status"`

	DocumentStatus DocumentStatus `yaml:"document_status"`

	OrderStatus OrderStatus `yaml:"order_status"`

	PatientClass PatientClass `yaml:"patient_class"`

	PatientAccountStatus PatientAccountStatus `yaml:"patient_account_status"`

	Gender Gender

	AbnormalFlags AbnormalFlags `yaml:"abnormal_flags"`

	// PrimaryFacility is the patient's primary facility.
	// If none of the PrimaryFacility fields are set, we want the entire PD1.3 Patient Primary Facility field to be
	// empty in the resulting HL7 message instead of a having field with separators but empty values.
	// Making PrimaryFacility a pointer is an easy way to detect when all fields are unset.
	PrimaryFacility *PrimaryFacility `yaml:"primary_facility"`

	// HospitalService is the default value to be set in the PV1.10-Hospital Service field.
	// This is overridden per pathway by the pathway's Consultant.
	HospitalService string `yaml:"hospital_service"`

	// CodingSystem is the default coding system of Order Profiles and their Test Types.
	// It is used to construct the Coded Element.
	CodingSystem string `yaml:"coding_system"`

	Mapping CodeMapping
}

// Header contains the configuration of the Message Header (MSH segment).
type Header struct {
	// Default is the default configuration for all messages.
	// Required.
	Default HeaderForType `yaml:"default"`
	// ORU is the configuration for ORU messages.
	// Optional. If not present, ORU messages will use the Default.
	ORU *HeaderForType
}

// HeaderForType contains the fields in the Message Header (MSH segment).
// All fields must be present.
type HeaderForType struct {
	// SendingApplication is the value to set in MSH-3 Sending Application.
	SendingApplication string `yaml:"sending_application"`
	// SendingFacility is the value to set in MSH-4 Sending Facility.
	SendingFacility string `yaml:"sending_facility"`
	// ReceivingApplication is the value to set in MSH-5 Receiving Application.
	ReceivingApplication string `yaml:"receiving_application"`
	// ReceivingFacility is the value to set in MSH-6 Receiving Facility.
	ReceivingFacility string `yaml:"receiving_facility"`
}

// HL7Allergy contains the configuration for AL1 segment (allergies).
type HL7Allergy struct {
	// Types is a list of the possible types of allergy types to be set in the AL1.2.AllergyTypes field.
	Types []string
	// Severities is a list of the possible types of allergy severities to be set in the AL1.4.AllergySeverity field.
	Severities []string
	// CodingSystem is the allergy coding system to be set in the CE.3.NameOfCodingSystem field in the
	// AL1.3.AllergyCode/Mnemonic/Description.
	CodingSystem string `yaml:"coding_system"`
}

// HL7Diagnosis is the configuration for DG1 segment (diagnosis).
type HL7Diagnosis struct {
	// Types is a list of the possible types of diagnosis to be set in the DG1.6.DiagnosisTypes field.
	Types []string
	// CodingSystem is the diagnosis coding system to be set in the CE.3.NameOfCodingSystem field in the
	// DG1.3.Diagnosis Code - DG1.
	CodingSystem string `yaml:"coding_system"`
}

// HL7Document contains configuration for a TXA segment (document).
type HL7Document struct {
	// Types is the document types to be set in the TXA.2.DocumentType field.
	Types []string
}

// HL7Procedure is the configuration for PR1 segment (procedure).
type HL7Procedure struct {
	// Types is the possible types of procedure to be set in the PR1.6.ProcedureTypes field.
	Types []string
	// CodingSystem is the procedure coding system to be set in the CE.3.NameOfCodingSystem field in the
	// PR1.3.Procedure Code.
	CodingSystem string `yaml:"coding_system"`
}

// OrderControl contains the values for the ORC.1 Order Control field.
// Values: http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/Default.aspx?version=HL7+v2.3.1&table=0119
type OrderControl struct {
	// New represents a New Order.
	New string
	// OK means that the Order/service is accepted & OK.
	OK string
	// WithObservations is the order control value for a status of "Observations/Performed Service to follow" (the results
	// have arrived).
	WithObservations string `yaml:"with_observations"`
}

// ResultStatus for the OBR.25 Result Status field.
// Values: http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/Default.aspx?version=HL7%20v2.5.1&table=0123
type ResultStatus struct {
	// Final means that the results are stored and verified. Can only be changed with a corrected result.
	Final string
	// Corrected means that the record coming over is a correction and thus replaces a final result.
	Corrected string
}

// DocumentStatus is set in the OBR.25 Result Status field for a Clinical Note.
type DocumentStatus struct {
	// Authenticated means that the document is authenticated.
	Authenticated string `yaml:"authenticated"`
}

// OrderStatus for the ORC.5 Order Status field.
// Values: http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/Default.aspx?version=HL7%20v2.5.1&table=0038
type OrderStatus struct {
	// Completed means that the order is completed.
	Completed string
	// InProcess means that the status is in process, unspecified.
	InProcess string `yaml:"in_process"`
}

// PatientClass are the patient class values to set in the PV1.2.PatientClass field.
// Values: http://hl7-definition.caristix.com:9010/Default.aspx?version=HL7%20v2.5.1&table=0004
type PatientClass struct {
	// Outpatient is the default patient class for newly created patients and outpatients.
	Outpatient string
	// Inpatient is the patient class for inpatients (set after an ADT^A01 Admission message).
	Inpatient string
}

// PatientAccountStatus are the patient account status values to set in the PV1.41.AccountStatus field.
type PatientAccountStatus struct {
	// Arrived means that the patient has been admitted.
	Arrived string
	// Cancelled means that the admission has been cancelled, e.g. the visit or a pending admission
	// has been cancelled.
	Cancelled string
	// Finished means that the patient encounter has finished and the patient is no longer admitted.
	Finished string
	// Planned means that the patient is going to be admitted, but it not admitted yet.
	Planned string
}

// Gender are the values to set in the PID.8 Sex field.
// Values: http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/segment/PID?version=HL7%20v2.3.1&table=0001
type Gender struct {
	Male    string
	Female  string
	Unknown string
}

// AbnormalFlags are the abnormal flag values to set in the OBX.8 Abnormal Flags field.
// Values: http://hl7-definition.caristix.com:9010/HL7%20v2.2/table/Default.aspx?version=HL7+v2.2&table=0078
type AbnormalFlags struct {
	AboveHighNormal string `yaml:"above_high_normal"`
	BelowLowNormal  string `yaml:"below_low_normal"`
}

// PrimaryFacility is the Primary Facility to set in the PD1.3 Patient Primary Facility field. Type XON.
// http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/segment/PD1
type PrimaryFacility struct {
	// OrganizationName is the XON.1 (Organization name).
	OrganizationName string `yaml:"organization_name"`
	// IDNumber is the XON.3 (ID Number).
	// The ID is of type string in the HL7 XON data type:
	// http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/segment/PD1?version=HL7%20v2.3.1&dataType=XON.
	// We use a string because it allows for more generic values.
	IDNumber string `yaml:"id_number"`
}

// LoadHL7Config loads the HL7 configuration from the given file.
func LoadHL7Config(ctx context.Context, fileName string) (*HL7Config, error) {
	data, err := files.Read(ctx, fileName)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot parse HL7 configuration file %s", fileName)
	}

	c := new(HL7Config)
	if err := yaml.UnmarshalStrict(data, c); err != nil {
		return nil, errors.Wrapf(err, "cannot unmarshal HL7 configuration file %s", fileName)
	}

	return c, nil
}

// LoadHeaderConfig loads the Header configuration from the given file.
func LoadHeaderConfig(ctx context.Context, fileName string) (*Header, error) {
	data, err := files.Read(ctx, fileName)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot parse header configuration file %s", fileName)
	}

	h := new(Header)
	if err := yaml.UnmarshalStrict(data, h); err != nil {
		return nil, errors.Wrapf(err, "cannot unmarshal header configuration file %s", fileName)
	}

	if err := valid(h.Default); err != nil {
		return nil, errors.Wrapf(err, "invalid header configuration %s: invalid default", fileName)
	}
	if h.ORU != nil {
		if err := valid(*h.ORU); err != nil {
			return nil, errors.Wrapf(err, "invalid header configuration %s: invalid oru", fileName)
		}
	}

	return h, nil
}

func valid(h HeaderForType) error {
	if h.SendingFacility == "" {
		return errors.New("SendingFacility not set; this is required")
	}
	if h.SendingApplication == "" {
		return errors.New("SendingApplication not set; this is required")
	}
	if h.ReceivingFacility == "" {
		return errors.New("ReceivingFacility not set; this is required")
	}
	if h.ReceivingApplication == "" {
		return errors.New("ReceivingApplication not set; this is required")
	}
	return nil
}
