// Copyright 2023 Google LLC
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

// This file contains the schemas for HL7 messages, segments and values for HL7v2 version 2.2.0.
// It has been auto-generated from the HL7v2 specification.

package hl7

import "reflect"

// AD represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type AD struct {
	StreetAddress              *ST `hl7:"false,Street Address"`
	OtherDesignation           *ST `hl7:"false,Other Designation"`
	City                       *ST `hl7:"false,City"`
	StateOrProvince            *ST `hl7:"false,State Or Province"`
	ZipOrPostalCode            *ID `hl7:"false,Zip Or Postal Code"`
	Country                    *ID `hl7:"false,Country"`
	Type                       *ID `hl7:"false,Type"`
	OtherGeographicDesignation *ST `hl7:"false,Other Geographic Designation"`
}

// CE represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CE struct {
	Identifier                  *ID `hl7:"false,Identifier"`
	Text                        *ST `hl7:"false,Text"`
	NameOfCodingSystem          *ST `hl7:"false,Name Of Coding System"`
	AlternateIdentifier         *ST `hl7:"false,Alternate Identifier"`
	AlternateText               *ST `hl7:"false,Alternate Text"`
	NameOfAlternateCodingSystem *ST `hl7:"false,Name Of Alternate Coding System"`
}

// CK_ACCOUNT_NO represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CK_ACCOUNT_NO struct {
	AccountNumber    *NM `hl7:"false,Account Number"`
	CheckDigit       *NM `hl7:"false,Check Digit"`
	CheckDigitScheme *ID `hl7:"false,Check Digit Scheme"`
	FacilityID       *ID `hl7:"false,Facility ID"`
}

// CK_PAT_ID represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CK_PAT_ID struct {
	PatientID        *ST `hl7:"false,Patient ID"`
	CheckDigit       *NM `hl7:"false,Check Digit"`
	CheckDigitScheme *ID `hl7:"false,Check Digit Scheme"`
	FacilityID       *ID `hl7:"false,Facility ID"`
}

// CM_ABS_RANGE represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_ABS_RANGE struct {
	Range            *CM `hl7:"false,Range"`
	NumericChange    *NM `hl7:"false,Numeric Change"`
	PercentPerChange *NM `hl7:"false,Percent Per Change"`
	Days             *NM `hl7:"false,Days"`
}

// CM_AUI represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_AUI struct {
	AuthorizationNumber *ST `hl7:"false,Authorization Number"`
	Date                *DT `hl7:"false,Date"`
	Source              *ST `hl7:"false,Source"`
}

// CM_BATCH_TOTAL represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_BATCH_TOTAL struct {
	BatchTotal1 *NM `hl7:"false,Batch Total 1"`
	BatchTotal2 *NM `hl7:"false,Batch Total 2"`
}

// CM_CCD represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_CCD struct {
	WhenToCharge *ID `hl7:"false,When To Charge"`
	DateTime     *TS `hl7:"false,Date/Time"`
}

// CM_DDI represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_DDI struct {
	DelayDays    *ST `hl7:"false,Delay Days"`
	Amount       *NM `hl7:"false,Amount"`
	NumberOfDays *NM `hl7:"false,Number Of Days"`
}

// CM_DIN represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_DIN struct {
	Date            *TS `hl7:"false,Date"`
	InstitutionName *CE `hl7:"false,Institution Name"`
}

// CM_DLD represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_DLD struct {
	DischargeLocation *ID `hl7:"false,Discharge Location"`
	EffectiveDate     *TS `hl7:"false,Effective Date"`
}

// CM_DLT represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_DLT struct {
	Range            *CM `hl7:"false,Range"`
	NumericThreshold *NM `hl7:"false,Numeric Threshold"`
	Change           *ST `hl7:"false,Change"`
	LengthOfTimeDays *NM `hl7:"false,Length Of Time-Days"`
}

// CM_DTN represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_DTN struct {
	DayType      *ID `hl7:"false,Day Type"`
	NumberOfDays *NM `hl7:"false,Number Of Days"`
}

// CM_EIP represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_EIP struct {
	ParentSPlacerOrderNumber *ST `hl7:"false,Parent´s Placer Order Number"`
	ParentSFillerOrderNumber *ST `hl7:"false,Parent´s Filler Order Number"`
}

// CM_ELD represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_ELD struct {
	SegmentID            *ST `hl7:"false,Segment-ID"`
	Sequence             *NM `hl7:"false,Sequence"`
	FieldPosition        *NM `hl7:"false,Field-Position"`
	CodeIdentifyingError *CE `hl7:"false,Code Identifying Error"`
}

// CM_FILLER represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_FILLER struct {
	UniqueFillerId      *ID `hl7:"false,Unique Filler Id"`
	FillerApplicationID *ID `hl7:"false,Filler Application ID"`
}

// CM_FINANCE represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_FINANCE struct {
	FinancialClassID *ID `hl7:"false,Financial Class ID"`
	EffectiveDate    *TS `hl7:"false,Effective Date"`
}

// CM_GROUP_ID represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_GROUP_ID struct {
	UniqueGroupId       *ID `hl7:"false,Unique Group Id"`
	PlacerApplicationId *ID `hl7:"false,Placer Application Id"`
}

// CM_INTERNAL_LOCATION represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_INTERNAL_LOCATION struct {
	NurseUnitStation *ID `hl7:"false,Nurse Unit (Station)"`
	Room             *ID `hl7:"false,Room"`
	Bed              *ID `hl7:"false,Bed"`
	FacilityID       *ID `hl7:"false,Facility ID"`
	BedStatus        *ID `hl7:"false,Bed Status"`
	Etage            *ID `hl7:"false,Etage"`
	Klinik           *ID `hl7:"false,Klinik"`
	Zentrum          *ID `hl7:"false,Zentrum"`
}

// CM_JOB_CODE represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_JOB_CODE struct {
	JobCode                *ID `hl7:"false,Job Code"`
	EmployeeClassification *ID `hl7:"false,Employee Classification"`
}

// CM_LA1 represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_LA1 struct {
	DispenseDeliverToLocation *CM `hl7:"false,Dispense / Deliver To Location"`
	Location                  *AD `hl7:"false,Location"`
}

// CM_LICENSE_NO represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_LICENSE_NO struct {
	LicenseNumber               *ST `hl7:"false,License Number"`
	IssuingStateProvinceCountry *ST `hl7:"false,Issuing State,Province,Country"`
}

// CM_MOC represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_MOC struct {
	DollarAmount *ST `hl7:"false,Dollar Amount"`
	ChargeCode   *ST `hl7:"false,Charge Code"`
}

// CM_MSG represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_MSG struct {
	MessageType  *ID `hl7:"false,Message Type"`
	TriggerEvent *ID `hl7:"false,Trigger Event"`
}

// CM_NDL represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_NDL struct {
	InterpreterTechnician *CN `hl7:"false,Interpreter / Technician"`
	StartDateTime         *TS `hl7:"false,Start Date/Time"`
	EndDateTime           *TS `hl7:"false,End Date/Time"`
	Location              *CM `hl7:"false,Location"`
}

// CM_OCD represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_OCD struct {
	OccurrenceCode *ID `hl7:"false,Occurrence Code"`
	OccurrenceDate *DT `hl7:"false,Occurrence Date"`
}

// CM_OSP represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_OSP struct {
	OccurrenceSpanCode      *ID `hl7:"false,Occurrence Span Code"`
	OccurrenceSpanStartDate *DT `hl7:"false,Occurrence Span Start Date"`
	OccurrenceSpanStopDate  *DT `hl7:"false,Occurrence Span Stop Date"`
}

// CM_PAT_ID represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_PAT_ID struct {
	PatientID        *ST `hl7:"false,Patient ID"`
	CheckDigit       *NM `hl7:"false,Check Digit"`
	CheckDigitScheme *ID `hl7:"false,Check Digit Scheme"`
	FacilityID       *ID `hl7:"false,Facility ID"`
	Type             *ID `hl7:"false,Type"`
}

// CM_PAT_ID_0192 represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_PAT_ID_0192 struct {
	PatientID        *ST `hl7:"false,Patient ID"`
	CheckDigit       *NM `hl7:"false,Check Digit"`
	CheckDigitScheme *ID `hl7:"false,Check Digit Scheme"`
	FacilityID       *ID `hl7:"false,Facility ID"`
	Type             *ID `hl7:"false,Type"`
}

// CM_PCF represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_PCF struct {
	PreCertificationPatientType *ID `hl7:"false,Pre-Certification Patient Type"`
	PreCerticationRequired      *ID `hl7:"false,Pre-Certication Required"`
	PreCertificationWindow      *TS `hl7:"false,Pre-Certification Window"`
}

// CM_PEN represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_PEN struct {
	PenaltyID     *ID `hl7:"false,Penalty ID"`
	PenaltyAmount *NM `hl7:"false,Penalty Amount"`
}

// CM_PIP represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_PIP struct {
	Privilege      *CE `hl7:"false,Privilege"`
	PrivilegeClass *CE `hl7:"false,Privilege Class"`
	ExpirationDate *DT `hl7:"false,Expiration Date"`
	ActivationDate *DT `hl7:"false,Activation Date"`
}

// CM_PLACER represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_PLACER struct {
	UniquePlacerId    *ID `hl7:"false,Unique Placer Id"`
	PlacerApplication *ID `hl7:"false,Placer Application"`
}

// CM_PLN represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_PLN struct {
	IDNumber                  *ST `hl7:"false,ID Number"`
	TypeOfIDNumber            *ID `hl7:"false,Type Of ID Number"`
	StateOtherQualifiyingInfo *ST `hl7:"false,State/Other Qualifiying Info"`
}

// CM_POSITION represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_POSITION struct {
	Saal  *ST `hl7:"false,Saal"`
	Tisch *ST `hl7:"false,Tisch"`
	Stuhl *ST `hl7:"false,Stuhl"`
}

// CM_PRACTITIONER represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_PRACTITIONER struct {
	ProcedurePractitionerID   *CN `hl7:"false,Procedure Practitioner  ID"`
	ProcedurePractitionerType *ID `hl7:"false,Procedure Practitioner Type"`
}

// CM_PTA represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_PTA struct {
	PolicyType  *ID `hl7:"false,Policy Type"`
	AmountClass *ID `hl7:"false,Amount Class"`
	Amount      *NM `hl7:"false,Amount"`
}

// CM_RANGE represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_RANGE struct {
	LowValue  *CE `hl7:"false,Low Value"`
	HighValue *CE `hl7:"false,High Value"`
}

// CM_RFR represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_RFR struct {
	ReferenceRange      *CE `hl7:"false,Reference Range"`
	Sex                 *ID `hl7:"false,Sex"`
	AgeRange            *CE `hl7:"false,Age Range"`
	GestationalAgeRange *CE `hl7:"false,Gestational Age Range"`
	Species             *ST `hl7:"false,Species"`
	RaceSubspecies      *ID `hl7:"false,Race / Subspecies"`
	TextCondition       *ST `hl7:"false,Text Condition"`
}

// CM_RI represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_RI struct {
	RepeatPattern        *ST `hl7:"false,Repeat Pattern"`
	ExplicitTimeIntevall *ST `hl7:"false,Explicit Time Intevall"`
}

// CM_RMC represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_RMC struct {
	RoomType       *ID `hl7:"false,Room Type"`
	AmountType     *ID `hl7:"false,Amount Type"`
	CoverageAmount *NM `hl7:"false,Coverage Amount"`
}

// CM_SPD represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_SPD struct {
	SpecialtyName       *ST `hl7:"false,Specialty Name"`
	GoverningBoard      *ST `hl7:"false,Governing Board"`
	EligibleOrCertified *ID `hl7:"false,Eligible Or Certified"`
	DateOfCertification *DT `hl7:"false,Date Of Certification"`
}

// CM_SPS represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_SPS struct {
	SpecimenSourceNameOrCode *CE `hl7:"false,Specimen Source Name Or Code"`
	Additives                *TX `hl7:"false,Additives"`
	Freetext                 *TX `hl7:"false,Freetext"`
	BodySite                 *CE `hl7:"false,Body Site"`
	SiteModifier             *CE `hl7:"false,Site Modifier"`
}

// CM_UVC represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_UVC struct {
	ValueCode   *ID `hl7:"false,Value Code"`
	ValueAmount *NM `hl7:"false,Value Amount"`
}

// CM_VR represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_VR struct {
	FirstDataCodeValue *ST `hl7:"false,First Data Code Value"`
	LastDataCodeCalue  *ST `hl7:"false,Last Data Code Calue"`
}

// CN_PERSON represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CN_PERSON struct {
	IDNumber            *ID `hl7:"false,ID Number"`
	FamiliyName         *ST `hl7:"false,Familiy Name"`
	GivenName           *ST `hl7:"false,Given Name"`
	MiddleInitialOrName *ST `hl7:"false,Middle Initial Or Name"`
	SuffixEGJROrIII     *ST `hl7:"false,Suffix (E.G. JR Or III)"`
	PrefixEGDR          *ST `hl7:"false,Prefix (E.G. DR)"`
	DegreeEGMD          *ST `hl7:"false,Degree (E.G. MD)"`
	SourceTableId       *ID `hl7:"false,Source Table Id"`
}

// CN_PHYSICIAN represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CN_PHYSICIAN struct {
	PhysicianID         *ID `hl7:"false,Physician ID"`
	FamiliyName         *ST `hl7:"false,Familiy Name"`
	GivenName           *ST `hl7:"false,Given Name"`
	MiddleInitialOrName *ST `hl7:"false,Middle Initial Or Name"`
	SuffixEGJROrIII     *ST `hl7:"false,Suffix (E.G. JR Or III)"`
	PrefixEGDR          *ST `hl7:"false,Prefix (E.G. DR)"`
	DegreeEGMD          *ST `hl7:"false,Degree (E.G. MD)"`
	SourceTableId       *ID `hl7:"false,Source Table Id"`
	Adresse             *AD `hl7:"false,Adresse"`
	Telefon             *TN `hl7:"false,Telefon"`
	Faxnummer           *TN `hl7:"false,Faxnummer"`
	OnlineNummer        *TN `hl7:"false,Online-Nummer"`
	EMail               *ST `hl7:"false,E-Mail"`
}

// CQ_QUANTITY represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CQ_QUANTITY struct {
	Quantity *ST `hl7:"false,Quantity"`
	Units    *ST `hl7:"false,Units"`
}

// PN represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type PN struct {
	FamiliyName         *ST `hl7:"false,Familiy Name"`
	GivenName           *ST `hl7:"false,Given Name"`
	MiddleInitialOrName *ST `hl7:"false,Middle Initial Or Name"`
	SuffixEGJROrIII     *ST `hl7:"false,Suffix (E.G. JR Or III)"`
	PrefixEGDR          *ST `hl7:"false,Prefix (E.G. DR)"`
	DegreeEGMD          *ST `hl7:"false,Degree (E.G. MD)"`
}

// TQ represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type TQ struct {
	Quantity        *CQ `hl7:"false,Quantity"`
	Interval        *CM `hl7:"false,Interval"`
	Duration        *ST `hl7:"false,Duration"`
	StartDateTime   *TS `hl7:"false,Start Date/Time"`
	EndDateTime     *TS `hl7:"false,End Date/Time"`
	Priority        *ID `hl7:"false,Priority"`
	Condition       *ST `hl7:"false,Condition"`
	Text            *TX `hl7:"false,Text"`
	Conjunction     *ID `hl7:"false,Conjunction"`
	OrderSequencing *ST `hl7:"false,Order Sequencing"`
}

// ACC represents the corresponding HL7 segment.
// Definition from HL7 2.2
type ACC struct {
	AccidentDateTime *TS `hl7:"false,Accident Date / Time"` // ACC-1
	AccidentCode     *ID `hl7:"false,Accident Code"`        // ACC-2
	AccidentLocation *ST `hl7:"false,Accident Location"`    // ACC-3
}

func (s *ACC) SegmentName() string {
	return "ACC"
}

// ADD represents the corresponding HL7 segment.
// Definition from HL7 2.2
type ADD struct {
	AddendumContinuationPointer *ST `hl7:"false,Addendum Continuation Pointer"` // ADD-1
}

func (s *ADD) SegmentName() string {
	return "ADD"
}

// AL1 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type AL1 struct {
	SetIDAllergy                   *SI `hl7:"true,Set ID - Allergy"`                      // AL1-1
	AllergyType                    *ID `hl7:"false,Allergy Type"`                         // AL1-2
	AllergyCodeMnemonicDescription *CE `hl7:"true,Allergy Code / Mnemonic / Description"` // AL1-3
	AllergySeverity                *ID `hl7:"false,Allergy Severity"`                     // AL1-4
	AllergyReaction                *ST `hl7:"false,Allergy Reaction"`                     // AL1-5
	IdentificationDate             *DT `hl7:"false,Identification Date"`                  // AL1-6
}

func (s *AL1) SegmentName() string {
	return "AL1"
}

// BHS represents the corresponding HL7 segment.
// Definition from HL7 2.2
type BHS struct {
	BatchFieldSeparator       *ST `hl7:"true,Batch Field Separator"`        // BHS-1
	BatchEncodingCharacters   *ST `hl7:"true,Batch Encoding Characters"`    // BHS-2
	BatchSendingApplication   *ST `hl7:"false,Batch Sending Application"`   // BHS-3
	BatchSendingFacility      *ST `hl7:"false,Batch Sending Facility"`      // BHS-4
	BatchReceivingApplication *ST `hl7:"false,Batch Receiving Application"` // BHS-5
	BatchReceivingFacility    *ST `hl7:"false,Batch Receiving Facility"`    // BHS-6
	BatchCreationDateTime     *TS `hl7:"false,Batch Creation Date / Time"`  // BHS-7
	BatchSecurity             *ST `hl7:"false,Batch Security"`              // BHS-8
	BatchNameIDType           *ST `hl7:"false,Batch Name / ID / Type"`      // BHS-9
	BatchComment              *ST `hl7:"false,Batch Comment"`               // BHS-10
	BatchControlID            *ST `hl7:"false,Batch Control ID"`            // BHS-11
	ReferenceBatchControlID   *ST `hl7:"false,Reference Batch Control ID"`  // BHS-12
}

func (s *BHS) SegmentName() string {
	return "BHS"
}

// BLG represents the corresponding HL7 segment.
// Definition from HL7 2.2
type BLG struct {
	WhenToCharge *CM `hl7:"false,When To Charge"` // BLG-1
	ChargeType   *ID `hl7:"false,Charge Type"`    // BLG-2
	AccountID    *CK `hl7:"false,Account ID"`     // BLG-3
}

func (s *BLG) SegmentName() string {
	return "BLG"
}

// BTS represents the corresponding HL7 segment.
// Definition from HL7 2.2
type BTS struct {
	BatchMessageCount *ST  `hl7:"false,Batch Message Count"` // BTS-1
	BatchComment      *ST  `hl7:"false,Batch Comment"`       // BTS-2
	BatchTotals       []CM `hl7:"false,Batch Totals"`        // BTS-3
}

func (s *BTS) SegmentName() string {
	return "BTS"
}

// DG1 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type DG1 struct {
	SetIDDiagnosis          *SI `hl7:"true,Set ID - Diagnosis"`         // DG1-1
	DiagnosisCodingMethod   *ID `hl7:"true,Diagnosis Coding Method"`    // DG1-2
	DiagnosisCode           *ID `hl7:"false,Diagnosis Code"`            // DG1-3
	DiagnosisDescription    *ST `hl7:"false,Diagnosis Description"`     // DG1-4
	DiagnosisDateTime       *TS `hl7:"false,Diagnosis Date / Time"`     // DG1-5
	DiagnosisDRGType        *ID `hl7:"true,Diagnosis / DRG Type"`       // DG1-6
	MajorDiagnosticCategory *CE `hl7:"false,Major Diagnostic Category"` // DG1-7
	DiagnosticRelatedGroup  *ID `hl7:"false,Diagnostic Related Group"`  // DG1-8
	DRGApprovalIndicator    *ID `hl7:"false,DRG Approval Indicator"`    // DG1-9
	DRGGrouperReviewCode    *ID `hl7:"false,DRG Grouper Review Code"`   // DG1-10
	OutlierType             *ID `hl7:"false,Outlier Type"`              // DG1-11
	OutlierDays             *NM `hl7:"false,Outlier Days"`              // DG1-12
	OutlierCost             *NM `hl7:"false,Outlier Cost"`              // DG1-13
	GrouperVersionAndType   *ST `hl7:"false,Grouper Version And Type"`  // DG1-14
	DiagnosisDRGPriority    *NM `hl7:"false,Diagnosis / DRG Priority"`  // DG1-15
	DiagnosingClinician     *CN `hl7:"false,Diagnosing Clinician"`      // DG1-16
}

func (s *DG1) SegmentName() string {
	return "DG1"
}

// DSC represents the corresponding HL7 segment.
// Definition from HL7 2.2
type DSC struct {
	ContinuationPointer *ST `hl7:"false,Continuation Pointer"` // DSC-1
}

func (s *DSC) SegmentName() string {
	return "DSC"
}

// DSP represents the corresponding HL7 segment.
// Definition from HL7 2.2
type DSP struct {
	SetIDDisplayData  *SI `hl7:"false,Set ID - Display Data"` // DSP-1
	DisplayLevel      *SI `hl7:"false,Display Level"`         // DSP-2
	DataLine          *TX `hl7:"true,Data Line"`              // DSP-3
	LogicalBreakPoint *ST `hl7:"false,Logical Break Point"`   // DSP-4
	ResultID          *TX `hl7:"false,Result ID"`             // DSP-5
}

func (s *DSP) SegmentName() string {
	return "DSP"
}

// ERR represents the corresponding HL7 segment.
// Definition from HL7 2.2
type ERR struct {
	ErrorCodeAndLocation []CM `hl7:"true,Error Code And Location"` // ERR-1
}

func (s *ERR) SegmentName() string {
	return "ERR"
}

// EVN represents the corresponding HL7 segment.
// Definition from HL7 2.2
type EVN struct {
	EventTypeCode        *ID `hl7:"true,Event Type Code"`            // EVN-1
	DateTimeOfEvent      *TS `hl7:"true,Date / Time Of Event"`       // EVN-2
	DateTimePlannedEvent *TS `hl7:"false,Date / Time Planned Event"` // EVN-3
	EventReasonCode      *ID `hl7:"false,Event Reason Code"`         // EVN-4
	OperatorID           *ID `hl7:"false,Operator ID"`               // EVN-5
}

func (s *EVN) SegmentName() string {
	return "EVN"
}

// FHS represents the corresponding HL7 segment.
// Definition from HL7 2.2
type FHS struct {
	FileFieldSeparator       *ST `hl7:"true,File Field Separator"`        // FHS-1
	FileEncodingCharacters   *ST `hl7:"true,File Encoding Characters"`    // FHS-2
	FileSendingApplication   *ST `hl7:"false,File Sending Application"`   // FHS-3
	FileSendingFacility      *ST `hl7:"false,File Sending Facility"`      // FHS-4
	FileReceivingApplication *ST `hl7:"false,File Receiving Application"` // FHS-5
	FileReceivingFacility    *ST `hl7:"false,File Receiving Facility"`    // FHS-6
	FileCreationDateTime     *TS `hl7:"false,File Creation Date / Time"`  // FHS-7
	FileSecurity             *ST `hl7:"false,File Security"`              // FHS-8
	FileNameID               *ST `hl7:"false,File Name / ID"`             // FHS-9
	FileHeaderComment        *ST `hl7:"false,File Header Comment"`        // FHS-10
	FileControlID            *ST `hl7:"false,File Control ID"`            // FHS-11
	ReferenceFileControlID   *ST `hl7:"false,Reference File Control ID"`  // FHS-12
}

func (s *FHS) SegmentName() string {
	return "FHS"
}

// FT1 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type FT1 struct {
	SetIDFinancialTransaction       *SI  `hl7:"false,Set ID - Financial Transaction"`      // FT1-1
	TransactionID                   *ST  `hl7:"false,Transaction ID"`                      // FT1-2
	TransactionBatchID              *ST  `hl7:"false,Transaction Batch ID"`                // FT1-3
	TransactionDate                 *DT  `hl7:"true,Transaction Date"`                     // FT1-4
	TransactionPostingDate          *DT  `hl7:"false,Transaction Posting Date"`            // FT1-5
	TransactionType                 *ID  `hl7:"true,Transaction Type"`                     // FT1-6
	TransactionCode                 *CE  `hl7:"true,Transaction Code"`                     // FT1-7
	TransactionDescription          *ST  `hl7:"false,Transaction Description"`             // FT1-8
	TransactionDescriptionAlternate *ST  `hl7:"false,Transaction Description - Alternate"` // FT1-9
	TransactionQuantity             *NM  `hl7:"false,Transaction Quantity"`                // FT1-10
	TransactionAmountExtended       *NM  `hl7:"false,Transaction Amount - Extended"`       // FT1-11
	TransactionAmountUnit           *NM  `hl7:"false,Transaction Amount - Unit"`           // FT1-12
	DepartmentCode                  *CE  `hl7:"false,Department Code"`                     // FT1-13
	InsurancePlanID                 *ID  `hl7:"true,Insurance Plan ID"`                    // FT1-14
	InsuranceAmount                 *NM  `hl7:"false,Insurance Amount"`                    // FT1-15
	AssignedPatientLocation         *CM  `hl7:"false,Assigned Patient Location"`           // FT1-16
	FeeSchedule                     *ID  `hl7:"false,Fee Schedule"`                        // FT1-17
	PatientType                     *ID  `hl7:"false,Patient Type"`                        // FT1-18
	DiagnosisCode                   []CE `hl7:"false,Diagnosis Code"`                      // FT1-19
	PerformedByCode                 *CN  `hl7:"false,Performed By Code"`                   // FT1-20
	OrderedByCode                   *CN  `hl7:"false,Ordered By Code"`                     // FT1-21
	UnitCost                        *NM  `hl7:"false,Unit Cost"`                           // FT1-22
	FillerOrderNumber               *CM  `hl7:"false,Filler Order Number"`                 // FT1-23
}

func (s *FT1) SegmentName() string {
	return "FT1"
}

// FTS represents the corresponding HL7 segment.
// Definition from HL7 2.2
type FTS struct {
	FileBatchCount     *NM `hl7:"false,File Batch Count"`     // FTS-1
	FileTrailerComment *ST `hl7:"false,File Trailer Comment"` // FTS-2
}

func (s *FTS) SegmentName() string {
	return "FTS"
}

// GT1 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type GT1 struct {
	SetIDGuarantor                *SI  `hl7:"true,Set ID - Guarantor"`                 // GT1-1
	GuarantorNumber               *CK  `hl7:"false,Guarantor Number"`                  // GT1-2
	GuarantorName                 *PN  `hl7:"true,Guarantor Name"`                     // GT1-3
	GuarantorSpouseName           *PN  `hl7:"false,Guarantor Spouse Name"`             // GT1-4
	GuarantorAddress              *AD  `hl7:"false,Guarantor Address"`                 // GT1-5
	GuarantorPhoneNumberHome      []TN `hl7:"false,Guarantor Phone Number - Home"`     // GT1-6
	GuarantorPhoneNumberBusiness  []TN `hl7:"false,Guarantor Phone Number - Business"` // GT1-7
	GuarantorDateOfBirth          *DT  `hl7:"false,Guarantor Date Of Birth"`           // GT1-8
	GuarantorSex                  *ID  `hl7:"false,Guarantor Sex"`                     // GT1-9
	GuarantorType                 *ID  `hl7:"false,Guarantor Type"`                    // GT1-10
	GuarantorRelationship         *ID  `hl7:"false,Guarantor Relationship"`            // GT1-11
	GuarantorSocialSecurityNumber *ST  `hl7:"false,Guarantor Social Security Number"`  // GT1-12
	GuarantorDateBegin            *DT  `hl7:"false,Guarantor Date - Begin"`            // GT1-13
	GuarantorDateEnd              *DT  `hl7:"false,Guarantor Date - End"`              // GT1-14
	GuarantorPriority             *NM  `hl7:"false,Guarantor Priority"`                // GT1-15
	GuarantorEmployerName         *ST  `hl7:"false,Guarantor Employer Name"`           // GT1-16
	GuarantorEmployerAddress      *AD  `hl7:"false,Guarantor Employer Address"`        // GT1-17
	GuarantorEmployPhoneNumber    []TN `hl7:"false,Guarantor Employ Phone Number"`     // GT1-18
	GuarantorEmployeeIDNumber     *ST  `hl7:"false,Guarantor Employee ID Number"`      // GT1-19
	GuarantorEmploymentStatus     *ID  `hl7:"false,Guarantor Employment Status"`       // GT1-20
	GuarantorOrganization         *ST  `hl7:"false,Guarantor Organization"`            // GT1-21
}

func (s *GT1) SegmentName() string {
	return "GT1"
}

// IN1 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type IN1 struct {
	SetIDInsurance                 *SI  `hl7:"true,Set ID - Insurance"`                   // IN1-1
	InsurancePlanID                *ID  `hl7:"true,Insurance Plan ID"`                    // IN1-2
	InsuranceCompanyID             *ST  `hl7:"true,Insurance Company ID"`                 // IN1-3
	InsuranceCompanyName           *ST  `hl7:"false,Insurance Company Name"`              // IN1-4
	InsuranceCompanyAddress        *AD  `hl7:"false,Insurance Company Address"`           // IN1-5
	InsuranceCompanyContactPers    *PN  `hl7:"false,Insurance Company Contact Pers"`      // IN1-6
	InsuranceCompanyPhoneNumber    []TN `hl7:"false,Insurance Company Phone Number"`      // IN1-7
	GroupNumber                    *ST  `hl7:"false,Group Number"`                        // IN1-8
	GroupName                      *ST  `hl7:"false,Group Name"`                          // IN1-9
	InsuredSGroupEmployerID        *ST  `hl7:"false,Insured'S Group Employer ID"`         // IN1-10
	InsuredSGroupEmployerName      *ST  `hl7:"false,Insured'S Group Employer Name"`       // IN1-11
	PlanEffectiveDate              *DT  `hl7:"false,Plan Effective Date"`                 // IN1-12
	PlanExpirationDate             *DT  `hl7:"false,Plan Expiration Date"`                // IN1-13
	AuthorizationInformation       *CM  `hl7:"false,Authorization Information"`           // IN1-14
	PlanType                       *ID  `hl7:"false,Plan Type"`                           // IN1-15
	NameOfInsured                  *PN  `hl7:"false,Name Of Insured"`                     // IN1-16
	InsuredSRelationshipToPatient  *ID  `hl7:"false,Insured'S Relationship To Patient"`   // IN1-17
	InsuredSDateOfBirth            *DT  `hl7:"false,Insured'S Date Of Birth"`             // IN1-18
	InsuredSAddress                *AD  `hl7:"false,Insured'S Address"`                   // IN1-19
	AssignmentOfBenefits           *ID  `hl7:"false,Assignment Of Benefits"`              // IN1-20
	CoordinationOfBenefits         *ID  `hl7:"false,Coordination Of Benefits"`            // IN1-21
	CoordinationOfBenefitsPriority *ST  `hl7:"false,Coordination Of Benefits - Priority"` // IN1-22
	NoticeOfAdmissionCode          *ID  `hl7:"false,Notice Of Admission Code"`            // IN1-23
	NoticeOfAdmissionDate          *DT  `hl7:"false,Notice Of Admission Date"`            // IN1-24
	ReportOfEligibilityCode        *ID  `hl7:"false,Report Of Eligibility Code"`          // IN1-25
	ReportOfEligibilityDate        *DT  `hl7:"false,Report Of Eligibility Date"`          // IN1-26
	ReleaseInformationCode         *ID  `hl7:"false,Release Information Code"`            // IN1-27
	PreAdmitCertification          *ST  `hl7:"false,Pre-Admit Certification"`             // IN1-28
	VerificationDateTime           *TS  `hl7:"false,Verification Date / Time"`            // IN1-29
	VerificationBy                 *CN  `hl7:"false,Verification By"`                     // IN1-30
	TypeOfAgreementCode            *ID  `hl7:"false,Type Of Agreement Code"`              // IN1-31
	BillingStatus                  *ID  `hl7:"false,Billing Status"`                      // IN1-32
	LifetimeReserveDays            *NM  `hl7:"false,Lifetime Reserve Days"`               // IN1-33
	DelayBeforeLifetimeReserveDays *NM  `hl7:"false,Delay Before Lifetime Reserve Days"`  // IN1-34
	CompanyPlanCode                *ID  `hl7:"false,Company Plan Code"`                   // IN1-35
	PolicyNumber                   *ST  `hl7:"false,Policy Number"`                       // IN1-36
	PolicyDeductible               *NM  `hl7:"false,Policy Deductible"`                   // IN1-37
	PolicyLimitAmount              *NM  `hl7:"false,Policy Limit - Amount"`               // IN1-38
	PolicyLimitDays                *NM  `hl7:"false,Policy Limit - Days"`                 // IN1-39
	RoomRateSemiPrivate            *NM  `hl7:"false,Room Rate - Semi-Private"`            // IN1-40
	RoomRatePrivate                *NM  `hl7:"false,Room Rate - Private"`                 // IN1-41
	InsuredSEmploymentStatus       *CE  `hl7:"false,Insured'S Employment Status"`         // IN1-42
	InsuredSSex                    *ID  `hl7:"false,Insured'S Sex"`                       // IN1-43
	InsuredSEmployerAddress        *AD  `hl7:"false,Insured'S Employer Address"`          // IN1-44
	VerificationStatus             *ST  `hl7:"false,Verification Status"`                 // IN1-45
	PriorInsurancePlanID           *ID  `hl7:"false,Prior Insurance Plan ID"`             // IN1-46
}

func (s *IN1) SegmentName() string {
	return "IN1"
}

// IN2 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type IN2 struct {
	InsuredSEmployeeID                        *ST  `hl7:"false,Insured'S Employee ID"`                          // IN2-1
	InsuredSSocialSecurityNumber              *ST  `hl7:"false,Insured'S Social Security Number"`               // IN2-2
	InsuredSEmployerName                      *CN  `hl7:"false,Insured'S Employer Name"`                        // IN2-3
	EmployerInformationData                   *ID  `hl7:"false,Employer Information Data"`                      // IN2-4
	MailClaimParty                            *ID  `hl7:"false,Mail Claim Party"`                               // IN2-5
	MedicareHealthInsuranceCardNumber         *NM  `hl7:"false,Medicare Health Insurance Card Number"`          // IN2-6
	MedicaidCaseName                          *PN  `hl7:"false,Medicaid Case Name"`                             // IN2-7
	MedicaidCaseNumber                        *NM  `hl7:"false,Medicaid Case Number"`                           // IN2-8
	ChampusSponsorName                        *PN  `hl7:"false,Champus Sponsor Name"`                           // IN2-9
	ChampusIDNumber                           *NM  `hl7:"false,Champus ID Number"`                              // IN2-10
	DependentOfChampusRecipient               *ID  `hl7:"false,Dependent Of Champus Recipient"`                 // IN2-11
	ChampusOrganization                       *ST  `hl7:"false,Champus Organization"`                           // IN2-12
	ChampusStation                            *ST  `hl7:"false,Champus Station"`                                // IN2-13
	ChampusService                            *ID  `hl7:"false,Champus Service"`                                // IN2-14
	ChampusRankGrade                          *ID  `hl7:"false,Champus Rank / Grade"`                           // IN2-15
	ChampusStatus                             *ID  `hl7:"false,Champus Status"`                                 // IN2-16
	ChampusRetireDate                         *DT  `hl7:"false,Champus Retire Date"`                            // IN2-17
	ChampusNonAvailabilityCertificationOnFile *ID  `hl7:"false,Champus Non-Availability Certification On File"` // IN2-18
	BabyCoverage                              *ID  `hl7:"false,Baby Coverage"`                                  // IN2-19
	CombineBabyBill                           *ID  `hl7:"false,Combine Baby Bill"`                              // IN2-20
	BloodDeductible                           *NM  `hl7:"false,Blood Deductible"`                               // IN2-21
	SpecialCoverageApprovalName               *PN  `hl7:"false,Special Coverage Approval Name"`                 // IN2-22
	SpecialCoverageApprovalTitle              *ST  `hl7:"false,Special Coverage Approval Title"`                // IN2-23
	NonCoveredInsuranceCode                   []ID `hl7:"false,Non-Covered Insurance Code"`                     // IN2-24
	PayorID                                   *ST  `hl7:"false,Payor ID"`                                       // IN2-25
	PayorSubscriberID                         *ST  `hl7:"false,Payor Subscriber ID"`                            // IN2-26
	EligibilitySource                         *ID  `hl7:"false,Eligibility Source"`                             // IN2-27
	RoomCoverageTypeAmount                    []CM `hl7:"false,Room Coverage Type / Amount"`                    // IN2-28
	PolicyTypeAmount                          []CM `hl7:"false,Policy Type / Amount"`                           // IN2-29
	DailyDeductible                           *CM  `hl7:"false,Daily Deductible"`                               // IN2-30
}

func (s *IN2) SegmentName() string {
	return "IN2"
}

// IN3 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type IN3 struct {
	SetIDInsuranceCertification        *SI  `hl7:"true,Set ID - Insurance Certification"`       // IN3-1
	CertificationNumber                *ST  `hl7:"false,Certification Number"`                  // IN3-2
	CertifiedBy                        *CN  `hl7:"false,Certified By"`                          // IN3-3
	CertificationRequired              *ID  `hl7:"false,Certification Required"`                // IN3-4
	Penalty                            *CM  `hl7:"false,Penalty"`                               // IN3-5
	CertificationDateTime              *TS  `hl7:"false,Certification Date / Time"`             // IN3-6
	CertificationModifyDateTime        *TS  `hl7:"false,Certification Modify Date / Time"`      // IN3-7
	Operator                           *CN  `hl7:"false,Operator"`                              // IN3-8
	CertificationBeginDate             *DT  `hl7:"false,Certification Begin Date"`              // IN3-9
	CertificationEndDate               *DT  `hl7:"false,Certification End Date"`                // IN3-10
	Days                               *CM  `hl7:"false,Days"`                                  // IN3-11
	NonConcurCodeDescription           *CE  `hl7:"false,Non-Concur Code / Description"`         // IN3-12
	NonConcurEffectiveDateTime         *TS  `hl7:"false,Non-Concur Effective Date / Time"`      // IN3-13
	PhysicianReviewer                  *CN  `hl7:"false,Physician Reviewer"`                    // IN3-14
	CertificationContact               *ST  `hl7:"false,Certification Contact"`                 // IN3-15
	CertificationContactPhoneNumber    []TN `hl7:"false,Certification Contact Phone Number"`    // IN3-16
	AppealReason                       *CE  `hl7:"false,Appeal Reason"`                         // IN3-17
	CertificationAgency                *CE  `hl7:"false,Certification Agency"`                  // IN3-18
	CertificationAgencyPhoneNumber     []TN `hl7:"false,Certification Agency Phone Number"`     // IN3-19
	PreCertificationRequiredWindow     []CM `hl7:"false,Pre-Certification Required / Window"`   // IN3-20
	CaseManager                        *ST  `hl7:"false,Case Manager"`                          // IN3-21
	SecondOpinionDate                  *DT  `hl7:"false,Second Opinion Date"`                   // IN3-22
	SecondOpinionStatus                *ID  `hl7:"false,Second Opinion Status"`                 // IN3-23
	SecondOpinionDocumentationReceived *ID  `hl7:"false,Second Opinion Documentation Received"` // IN3-24
	SecondOpinionPractitioner          *CN  `hl7:"false,Second Opinion Practitioner"`           // IN3-25
}

func (s *IN3) SegmentName() string {
	return "IN3"
}

// MFA represents the corresponding HL7 segment.
// Definition from HL7 2.2
type MFA struct {
	RecordLevelEventCode     *ID  `hl7:"true,Record-Level Event Code"`       // MFA-1
	MFNControlID             *ST  `hl7:"false,MFN Control ID"`               // MFA-2
	EventCompletionDateTime  *TS  `hl7:"false,Event Completion Date / Time"` // MFA-3
	ErrorReturnCodeAndOrText *CE  `hl7:"true,Error Return Code And/Or Text"` // MFA-4
	PrimaryKeyValue          []CE `hl7:"true,Primary Key Value"`             // MFA-5
}

func (s *MFA) SegmentName() string {
	return "MFA"
}

// MFE represents the corresponding HL7 segment.
// Definition from HL7 2.2
type MFE struct {
	RecordLevelEventCode *ID  `hl7:"true,Record-Level Event Code"` // MFE-1
	MFNControlID         *ST  `hl7:"false,MFN Control ID"`         // MFE-2
	EffectiveDateTime    *TS  `hl7:"false,Effective Date / Time"`  // MFE-3
	PrimaryKeyValue      []CE `hl7:"true,Primary Key Value"`       // MFE-4
}

func (s *MFE) SegmentName() string {
	return "MFE"
}

// MFI represents the corresponding HL7 segment.
// Definition from HL7 2.2
type MFI struct {
	MasterFileIdentifier            *CE `hl7:"true,Master File Identifier"`              // MFI-1
	MasterFileApplicationIdentifier *ID `hl7:"false,Master File Application Identifier"` // MFI-2
	FileLevelEventCode              *ID `hl7:"true,File-Level Event Code"`               // MFI-3
	EnteredDateTime                 *TS `hl7:"false,Entered Date / Time"`                // MFI-4
	EffectiveDateTime               *TS `hl7:"false,Effective Date / Time"`              // MFI-5
	ResponseLevelCode               *ID `hl7:"true,Response Level Code"`                 // MFI-6
}

func (s *MFI) SegmentName() string {
	return "MFI"
}

// MRG represents the corresponding HL7 segment.
// Definition from HL7 2.2
type MRG struct {
	PriorPatientIDInternal    *CM `hl7:"true,Prior Patient ID - Internal"`   // MRG-1
	PriorAlternatePatientID   *CM `hl7:"false,Prior Alternate Patient ID"`   // MRG-2
	PriorPatientAccountNumber *CK `hl7:"false,Prior Patient Account Number"` // MRG-3
	PriorPatientIDExternal    *CK `hl7:"false,Prior Patient ID - External"`  // MRG-4
}

func (s *MRG) SegmentName() string {
	return "MRG"
}

// MSA represents the corresponding HL7 segment.
// Definition from HL7 2.2
type MSA struct {
	AcknowledgementCode        *ID `hl7:"true,Acknowledgement Code"`          // MSA-1
	MessageControlID           *ST `hl7:"true,Message Control ID"`            // MSA-2
	TextMessage                *ST `hl7:"false,Text Message"`                 // MSA-3
	ExpectedSequenceNumber     *NM `hl7:"false,Expected Sequence Number"`     // MSA-4
	DelayedAcknowledgementType *ID `hl7:"false,Delayed Acknowledgement Type"` // MSA-5
	ErrorCondition             *CE `hl7:"false,Error Condition"`              // MSA-6
}

func (s *MSA) SegmentName() string {
	return "MSA"
}

// MSH represents the corresponding HL7 segment.
// Definition from HL7 2.2
type MSH struct {
	// Missing: MSH.1
	EncodingCharacters             *Delimiters `hl7:"true,Encoding Characters"`               // MSH-2
	SendingApplication             *ST         `hl7:"false,Sending Application"`              // MSH-3
	SendingFacility                *ST         `hl7:"false,Sending Facility"`                 // MSH-4
	ReceivingApplication           *ST         `hl7:"false,Receiving Application"`            // MSH-5
	ReceivingFacility              *ST         `hl7:"false,Receiving Facility"`               // MSH-6
	DateTimeOfMessage              *TS         `hl7:"false,Date / Time Of Message"`           // MSH-7
	Security                       *ST         `hl7:"false,Security"`                         // MSH-8
	MessageType                    *CM         `hl7:"true,Message Type"`                      // MSH-9
	MessageControlID               *ST         `hl7:"true,Message Control ID"`                // MSH-10
	ProcessingID                   *ID         `hl7:"true,Processing ID"`                     // MSH-11
	VersionID                      *ID         `hl7:"true,Version ID"`                        // MSH-12
	SequenceNumber                 *NM         `hl7:"false,Sequence Number"`                  // MSH-13
	ContinuationPointer            *ST         `hl7:"false,Continuation Pointer"`             // MSH-14
	AcceptAcknowledgementType      *ID         `hl7:"false,Accept Acknowledgement Type"`      // MSH-15
	ApplicationAcknowledgementType *ID         `hl7:"false,Application Acknowledgement Type"` // MSH-16
	CountryCode                    *ID         `hl7:"false,Country Code"`                     // MSH-17
}

func (s *MSH) SegmentName() string {
	return "MSH"
}

// NCK represents the corresponding HL7 segment.
// Definition from HL7 2.2
type NCK struct {
	SystemDateTime *TS `hl7:"true,System Date/Time"` // NCK-1
}

func (s *NCK) SegmentName() string {
	return "NCK"
}

// NK1 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type NK1 struct {
	SetIDNextOfKin          *SI  `hl7:"true,Set ID - Next Of Kin"`          // NK1-1
	Name                    *PN  `hl7:"false,Name"`                         // NK1-2
	Relationship            *CE  `hl7:"false,Relationship"`                 // NK1-3
	Address                 *AD  `hl7:"false,Address"`                      // NK1-4
	PhoneNumber             []TN `hl7:"false,Phone Number"`                 // NK1-5
	BusinessPhoneNumber     *TN  `hl7:"false,Business Phone Number"`        // NK1-6
	ContactRole             *CE  `hl7:"false,Contact Role"`                 // NK1-7
	StartDate               *DT  `hl7:"false,Start Date"`                   // NK1-8
	EndDate                 *DT  `hl7:"false,End Date"`                     // NK1-9
	NextOfKin               *ST  `hl7:"false,Next Of Kin"`                  // NK1-10
	NextOfKinJobCodeClass   *CM  `hl7:"false,Next Of Kin Job Code / Class"` // NK1-11
	NextOfKinEmployeeNumber *ST  `hl7:"false,Next Of Kin Employee Number"`  // NK1-12
	OrganizationName        *ST  `hl7:"false,Organization Name"`            // NK1-13
}

func (s *NK1) SegmentName() string {
	return "NK1"
}

// NPU represents the corresponding HL7 segment.
// Definition from HL7 2.2
type NPU struct {
	BedLocation *CM `hl7:"true,Bed Location"` // NPU-1
	BedStatus   *ID `hl7:"false,Bed Status"`  // NPU-2
}

func (s *NPU) SegmentName() string {
	return "NPU"
}

// NSC represents the corresponding HL7 segment.
// Definition from HL7 2.2
type NSC struct {
	NetworkChangeType  *ID `hl7:"true,Network Change Type"`  // NSC-1
	CurrentCPU         *ST `hl7:"false,Current CPU"`         // NSC-2
	CurrentFileserver  *ST `hl7:"false,Current Fileserver"`  // NSC-3
	CurrentApplication *ST `hl7:"false,Current Application"` // NSC-4
	CurrentFacility    *ST `hl7:"false,Current Facility"`    // NSC-5
	NewCPU             *ST `hl7:"false,New CPU"`             // NSC-6
	NewFileserver      *ST `hl7:"false,New Fileserver"`      // NSC-7
	NewApplication     *ST `hl7:"false,New Application"`     // NSC-8
	NewFacility        *ST `hl7:"false,New Facility"`        // NSC-9
}

func (s *NSC) SegmentName() string {
	return "NSC"
}

// NST represents the corresponding HL7 segment.
// Definition from HL7 2.2
type NST struct {
	StatisticsAvailable    *ID `hl7:"true,Statistics Available"`      // NST-1
	SourceIdentifier       *ST `hl7:"false,Source Identifier"`        // NST-2
	SourceType             *ID `hl7:"false,Source Type"`              // NST-3
	StatisticsStart        *TS `hl7:"false,Statistics Start"`         // NST-4
	StatisticsEnd          *TS `hl7:"false,Statistics End"`           // NST-5
	ReceiveCharacterCount  *NM `hl7:"false,Receive Character Count"`  // NST-6
	SendCharacterCount     *NM `hl7:"false,Send Character Count"`     // NST-7
	MessageReceived        *NM `hl7:"false,Message Received"`         // NST-8
	MessageSent            *NM `hl7:"false,Message Sent"`             // NST-9
	ChecksumErrorsReceived *NM `hl7:"false,Checksum Errors Received"` // NST-10
	LengthErrorsReceived   *NM `hl7:"false,Length Errors Received"`   // NST-11
	OtherErrorsReceived    *NM `hl7:"false,Other Errors Received"`    // NST-12
	ConnectTimeouts        *NM `hl7:"false,Connect Timeouts"`         // NST-13
	ReceiveTimeouts        *NM `hl7:"false,Receive Timeouts"`         // NST-14
	NetworkErrors          *NM `hl7:"false,Network Errors"`           // NST-15
}

func (s *NST) SegmentName() string {
	return "NST"
}

// NTE represents the corresponding HL7 segment.
// Definition from HL7 2.2
type NTE struct {
	SetIDNotesAndComments *SI  `hl7:"false,Set ID - Notes And Comments"` // NTE-1
	SourceOfComment       *ID  `hl7:"false,Source Of Comment"`           // NTE-2
	Comment               []FT `hl7:"false,Comment"`                     // NTE-3
}

func (s *NTE) SegmentName() string {
	return "NTE"
}

// OBR represents the corresponding HL7 segment.
// Definition from HL7 SYNTHETIC
type OBR struct {
	SetIDOBR                                   *SI   `hl7:"false,Set ID - OBR"`                                    // OBR-1
	PlacerOrderNumber                          *EI   `hl7:"false,Placer Order Number"`                             // OBR-2
	FillerOrderNumber                          *EI   `hl7:"false,Filler Order Number"`                             // OBR-3
	UniversalServiceIdentifier                 *CWE  `hl7:"true,Universal Service Identifier"`                     // OBR-4
	Priority                                   *ID   `hl7:"false,Priority"`                                        // OBR-5
	RequestedDateTime                          *TS   `hl7:"false,Requested Date/Time"`                             // OBR-6
	ObservationDateTime                        *TS   `hl7:"false,Observation Date/Time"`                           // OBR-7
	ObservationEndDateTime                     *TS   `hl7:"false,Observation End Date/Time"`                       // OBR-8
	CollectionVolume                           *CQ   `hl7:"false,Collection Volume"`                               // OBR-9
	CollectorIdentifier                        []XCN `hl7:"false,Collector Identifier"`                            // OBR-10
	SpecimenActionCode                         *ID   `hl7:"false,Specimen Action Code"`                            // OBR-11
	DangerCode                                 *CE   `hl7:"false,Danger Code"`                                     // OBR-12
	RelevantClinicalInformation                *ST   `hl7:"false,Relevant Clinical Information"`                   // OBR-13
	SpecimenReceivedDateTime                   *TS   `hl7:"false,Specimen Received Date/Time"`                     // OBR-14
	SpecimenSource                             *SPS  `hl7:"false,Specimen Source"`                                 // OBR-15
	OrderingProvider                           []XCN `hl7:"false,Ordering Provider"`                               // OBR-16
	OrderCallbackPhoneNumber                   []XTN `hl7:"false,Order Callback Phone Number"`                     // OBR-17
	PlacerField1                               *ST   `hl7:"false,Placer Field 1"`                                  // OBR-18
	PlacerField2                               *ST   `hl7:"false,Placer Field 2"`                                  // OBR-19
	FillerField1                               *ST   `hl7:"false,Filler Field 1"`                                  // OBR-20
	FillerField2                               *ST   `hl7:"false,Filler Field 2"`                                  // OBR-21
	ResultsRptStatusChngDateTime               *TS   `hl7:"false,Results Rpt/Status Chng - Date/Time"`             // OBR-22
	ChargeToPractice                           *MOC  `hl7:"false,Charge To Practice"`                              // OBR-23
	DiagnosticServSectID                       *ID   `hl7:"false,Diagnostic Serv Sect ID"`                         // OBR-24
	ResultStatus                               *ID   `hl7:"false,Result Status"`                                   // OBR-25
	ParentResult                               *PRL  `hl7:"false,Parent Result"`                                   // OBR-26
	QuantityTiming                             []TQ  `hl7:"false,Quantity/Timing"`                                 // OBR-27
	ResultCopiesTo                             []XCN `hl7:"false,Result Copies To"`                                // OBR-28
	Parent                                     *EIP  `hl7:"false,Parent"`                                          // OBR-29
	TransportationMode                         *ID   `hl7:"false,Transportation Mode"`                             // OBR-30
	ReasonForStudy                             []CE  `hl7:"false,Reason For Study"`                                // OBR-31
	PrincipalResultInterpreter                 *NDL  `hl7:"false,Principal Result Interpreter"`                    // OBR-32
	AssistantResultInterpreter                 []NDL `hl7:"false,Assistant Result Interpreter"`                    // OBR-33
	Technician                                 []NDL `hl7:"false,Technician"`                                      // OBR-34
	Transcriptionist                           []NDL `hl7:"false,Transcriptionist"`                                // OBR-35
	ScheduledDateTime                          *TS   `hl7:"false,Scheduled Date/Time"`                             // OBR-36
	NumberOfSampleContainers                   *NM   `hl7:"false,Number Of Sample Containers *"`                   // OBR-37
	TransportLogisticsOfCollectedSample        []CE  `hl7:"false,Transport Logistics Of Collected Sample"`         // OBR-38
	CollectorSComment                          []CE  `hl7:"false,Collector'S Comment *"`                           // OBR-39
	TransportArrangementResponsibility         *CE   `hl7:"false,Transport Arrangement Responsibility"`            // OBR-40
	TransportArranged                          *ID   `hl7:"false,Transport Arranged"`                              // OBR-41
	EscortRequired                             *ID   `hl7:"false,Escort Required"`                                 // OBR-42
	PlannedPatientTransportComment             []CE  `hl7:"false,Planned Patient Transport Comment"`               // OBR-43
	ProcedureCode                              *CE   `hl7:"false,Procedure Code"`                                  // OBR-44
	ProcedureCodeModifier                      []CE  `hl7:"false,Procedure Code Modifier"`                         // OBR-45
	PlacerSupplementalServiceInformation       []CE  `hl7:"false,Placer Supplemental Service Information"`         // OBR-46
	FillerSupplementalServiceInformation       []CE  `hl7:"false,Filler Supplemental Service Information"`         // OBR-47
	MedicallyNecessaryDuplicateProcedureReason *CWE  `hl7:"false,Medically Necessary Duplicate Procedure Reason."` // OBR-48
	ResultHandling                             *IS   `hl7:"false,Result Handling"`                                 // OBR-49
	ParentUniversalServiceIdentifier           *CWE  `hl7:"false,Parent Universal Service Identifier"`             // OBR-50
}

func (s *OBR) SegmentName() string {
	return "OBR"
}

// OBX represents the corresponding HL7 segment.
// Definition from HL7 SYNTHETIC
type OBX struct {
	SetIDOBX                               *SI   `hl7:"false,Set ID - OBX"`                                  // OBX-1
	ValueType                              *ID   `hl7:"false,Value Type"`                                    // OBX-2
	ObservationIdentifier                  *CWE  `hl7:"true,Observation Identifier"`                         // OBX-3
	ObservationSubID                       *ST   `hl7:"false,Observation Sub-ID"`                            // OBX-4
	ObservationValue                       []Any `hl7:"false,Observation Value"`                             // OBX-5
	Units                                  *CE   `hl7:"false,Units"`                                         // OBX-6
	ReferencesRange                        *ST   `hl7:"false,References Range"`                              // OBX-7
	AbnormalFlags                          []IS  `hl7:"false,Abnormal Flags"`                                // OBX-8
	Probability                            *NM   `hl7:"false,Probability"`                                   // OBX-9
	NatureOfAbnormalTest                   []ID  `hl7:"false,Nature Of Abnormal Test"`                       // OBX-10
	ObservationResultStatus                *ID   `hl7:"true,Observation Result Status"`                      // OBX-11
	EffectiveDateOfReferenceRangeValues    *TS   `hl7:"false,Effective Date Of Reference Range Values"`      // OBX-12
	UserDefinedAccessChecks                *ST   `hl7:"false,User Defined Access Checks"`                    // OBX-13
	DateTimeOfTheObservation               *TS   `hl7:"false,Date/Time Of The Observation"`                  // OBX-14
	ProducerSReference                     *CE   `hl7:"false,Producer'S Reference"`                          // OBX-15
	ResponsibleObserver                    []XCN `hl7:"false,Responsible Observer"`                          // OBX-16
	ObservationMethod                      []CE  `hl7:"false,Observation Method"`                            // OBX-17
	EquipmentInstanceIdentifier            []EI  `hl7:"false,Equipment Instance Identifier"`                 // OBX-18
	DateTimeOfTheAnalysis                  *TS   `hl7:"false,Date/Time Of The Analysis"`                     // OBX-19
	ReservedForHarmonizationWithVersion26A []XON `hl7:"false,Reserved For Harmonization With Version 2.6 A"` // OBX-20
	ReservedForHarmonizationWithVersion26B []XON `hl7:"false,Reserved For Harmonization With Version 2.6 B"` // OBX-21
	ReservedForHarmonizationWithVersion26C []XON `hl7:"false,Reserved For Harmonization With Version 2.6 C"` // OBX-22
	PerformingOrganizationName             *XON  `hl7:"false,Performing Organization Name"`                  // OBX-23
	PerformingOrganizationAddress          *XAD  `hl7:"false,Performing Organization Address"`               // OBX-24
	PerformingOrganizationMedicalDirector  *XCN  `hl7:"false,Performing Organization Medical Director"`      // OBX-25
}

func (s *OBX) SegmentName() string {
	return "OBX"
}

// ODS represents the corresponding HL7 segment.
// Definition from HL7 2.2
type ODS struct {
	Type                           *ID  `hl7:"true,Type"`                                 // ODS-1
	ServicePeriod                  []CE `hl7:"false,Service Period"`                      // ODS-2
	DietSupplementOrPreferenceCode []CE `hl7:"true,Diet, Supplement, Or Preference Code"` // ODS-3
	TextInstruction                []ST `hl7:"false,Text Instruction"`                    // ODS-4
}

func (s *ODS) SegmentName() string {
	return "ODS"
}

// ODT represents the corresponding HL7 segment.
// Definition from HL7 2.2
type ODT struct {
	TrayType        *CE  `hl7:"true,Tray Type"`         // ODT-1
	ServicePeriod   []CE `hl7:"false,Service Period"`   // ODT-2
	TextInstruction []ST `hl7:"false,Text Instruction"` // ODT-3
}

func (s *ODT) SegmentName() string {
	return "ODT"
}

// OM1 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type OM1 struct {
	SegmentTypeID                                 *ST  `hl7:"false,Segment Type ID"`                                       // OM1-1
	SequenceNumberTestObservationMasterFile       *NM  `hl7:"false,Sequence Number - Test/ Observation Master File"`       // OM1-2
	ProducerSTestObservationID                    *CE  `hl7:"true,Producer'S Test / Observation ID"`                       // OM1-3
	PermittedDataTypes                            []ID `hl7:"false,Permitted Data Types"`                                  // OM1-4
	SpecimenRequired                              *ID  `hl7:"true,Specimen Required"`                                      // OM1-5
	ProducerID                                    *CE  `hl7:"true,Producer ID"`                                            // OM1-6
	ObservationDescription                        *TX  `hl7:"false,Observation Description"`                               // OM1-7
	OtherTestObservationIDsForTheObservation      *CE  `hl7:"false,Other Test / Observation IDs For The Observation"`      // OM1-8
	OtherNames                                    []ST `hl7:"true,Other Names"`                                            // OM1-9
	PreferredReportNameForTheObservation          *ST  `hl7:"false,Preferred Report Name For The Observation"`             // OM1-10
	PreferredShortNameOrMnemonicForObservation    *ST  `hl7:"false,Preferred Short Name Or Mnemonic For Observation"`      // OM1-11
	PreferredLongNameForTheObservation            *ST  `hl7:"false,Preferred Long Name For The Observation"`               // OM1-12
	Orderability                                  *ID  `hl7:"false,Orderability"`                                          // OM1-13
	IdentityOfInstrumentUsedToPerformThisStudy    []CE `hl7:"false,Identity Of Instrument Used To Perform This Study"`     // OM1-14
	CodedRepresentationOfMethod                   []CE `hl7:"false,Coded Representation Of Method"`                        // OM1-15
	Portable                                      *ID  `hl7:"false,Portable"`                                              // OM1-16
	ObservationProducingDepartmentSection         []ID `hl7:"false,Observation Producing Department / Section"`            // OM1-17
	TelephoneNumberOfSection                      *TN  `hl7:"false,Telephone Number Of Section"`                           // OM1-18
	NatureOfTestObservation                       *ID  `hl7:"true,Nature Of Test / Observation"`                           // OM1-19
	ReportSubheader                               *CE  `hl7:"false,Report Subheader"`                                      // OM1-20
	ReportDisplayOrder                            *ST  `hl7:"false,Report Display Order"`                                  // OM1-21
	DateTimeStampForAnyChangeInDefinitionForObs   *TS  `hl7:"true,Date / Time Stamp For Any Change In Definition For Obs"` // OM1-22
	EffectiveDateTimeOfChange                     *TS  `hl7:"false,Effective Date / Time Of Change"`                       // OM1-23
	TypicalTurnAroundTime                         *NM  `hl7:"false,Typical Turn-Around Time"`                              // OM1-24
	ProcessingTime                                *NM  `hl7:"false,Processing Time"`                                       // OM1-25
	ProcessingPriority                            []ID `hl7:"false,Processing Priority"`                                   // OM1-26
	ReportingPriority                             *ID  `hl7:"false,Reporting Priority"`                                    // OM1-27
	OutsideSiteSWhereObservationMayBePerformed    []CE `hl7:"false,Outside Site(S) Where Observation May Be Performed"`    // OM1-28
	AddressOfOutsideSiteS                         []AD `hl7:"false,Address Of Outside Site(S)"`                            // OM1-29
	PhoneNumberOfOutsideSite                      []TN `hl7:"false,Phone Number Of Outside Site"`                          // OM1-30
	ConfidentialityCode                           *ID  `hl7:"false,Confidentiality Code"`                                  // OM1-31
	ObservationsRequiredToInterpretTheObservation []CE `hl7:"false,Observations Required To Interpret The Observation"`    // OM1-32
	InterpretationOfObservations                  *TX  `hl7:"false,Interpretation Of Observations"`                        // OM1-33
	ContraindicationsToObservations               []CE `hl7:"false,Contraindications To Observations"`                     // OM1-34
	ReflexTestsObservations                       []CE `hl7:"false,Reflex Tests / Observations"`                           // OM1-35
	RulesThatTriggerReflexTesting                 *ST  `hl7:"false,Rules That Trigger Reflex Testing"`                     // OM1-36
	FixedCannedMessage                            []CE `hl7:"false,Fixed Canned Message"`                                  // OM1-37
	PatientPreparation                            *TX  `hl7:"false,Patient Preparation"`                                   // OM1-38
	ProcedureMedication                           *CE  `hl7:"false,Procedure Medication"`                                  // OM1-39
	FactorsThatMayAffectTheObservation            *TX  `hl7:"false,Factors That May Affect The Observation"`               // OM1-40
	TestObservationPerformanceSchedule            []ST `hl7:"false,Test / Observation Performance Schedule"`               // OM1-41
	DescriptionOfTestMethods                      *TX  `hl7:"false,Description Of Test Methods"`                           // OM1-42
}

func (s *OM1) SegmentName() string {
	return "OM1"
}

// OM2 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type OM2 struct {
	SegmentTypeID                                     *ST  `hl7:"false,Segment Type ID"`                                              // OM2-1
	SequenceNumberTestObservationMasterFile           *NM  `hl7:"false,Sequence Number - Test/ Observation Master File"`              // OM2-2
	UnitsOfMeasure                                    *CE  `hl7:"false,Units Of Measure"`                                             // OM2-3
	RangeOfDecimalPrecision                           *NM  `hl7:"false,Range Of Decimal Precision"`                                   // OM2-4
	CorrespondingSIUnitsOfMeasure                     *CE  `hl7:"false,Corresponding SI Units Of Measure"`                            // OM2-5
	SIConversionFactor                                []TX `hl7:"true,SI Conversion Factor"`                                          // OM2-6
	ReferenceNormalRangeOrdinalContinuousObservations []CM `hl7:"false,Reference (Normal) Range - Ordinal & Continuous Observations"` // OM2-7
	CriticalRangeForOrdinalAndContinuousObservations  *CM  `hl7:"false,Critical Range For Ordinal And Continuous Observations"`       // OM2-8
	AbsoluteRangeForOrdinalAndContinuousObservations  *CM  `hl7:"false,Absolute Range For Ordinal And Continuous Observations"`       // OM2-9
	DeltaCheckCriteria                                []CM `hl7:"false,Delta Check Criteria"`                                         // OM2-10
	MinimumMeaningfulIncrements                       *NM  `hl7:"false,Minimum Meaningful Increments"`                                // OM2-11
}

func (s *OM2) SegmentName() string {
	return "OM2"
}

// OM3 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type OM3 struct {
	SegmentTypeID                               *ST  `hl7:"false,Segment Type ID"`                                  // OM3-1
	SequenceNumberTestObservationMasterFile     *NM  `hl7:"false,Sequence Number - Test/ Observation Master File"`  // OM3-2
	PreferredCodingSystem                       *ID  `hl7:"false,Preferred Coding System"`                          // OM3-3
	ValidCodedAnswers                           []CE `hl7:"false,Valid Coded Answers"`                              // OM3-4
	NormalTestCodesForCategoricalObservations   []CE `hl7:"false,Normal Test Codes For Categorical Observations"`   // OM3-5
	AbnormalTestCodesForCategoricalObservations *CE  `hl7:"false,Abnormal Test Codes For Categorical Observations"` // OM3-6
	CriticalTestCodesForCategoricalObservations *CE  `hl7:"false,Critical Test Codes For Categorical Observations"` // OM3-7
	DataType                                    *ID  `hl7:"false,Data Type"`                                        // OM3-8
}

func (s *OM3) SegmentName() string {
	return "OM3"
}

// OM4 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type OM4 struct {
	SegmentTypeID                           *ST  `hl7:"false,Segment Type ID"`                                 // OM4-1
	SequenceNumberTestObservationMasterFile *NM  `hl7:"false,Sequence Number - Test/ Observation Master File"` // OM4-2
	DerivedSpecimen                         *ID  `hl7:"false,Derived Specimen"`                                // OM4-3
	ContainerDescription                    *TX  `hl7:"false,Container Description"`                           // OM4-4
	ContainerVolume                         *NM  `hl7:"false,Container Volume"`                                // OM4-5
	ContainerUnits                          *CE  `hl7:"false,Container Units"`                                 // OM4-6
	Specimen                                *CE  `hl7:"false,Specimen"`                                        // OM4-7
	Additive                                *CE  `hl7:"false,Additive"`                                        // OM4-8
	Preparation                             *TX  `hl7:"false,Preparation"`                                     // OM4-9
	SpecialHandlingRequirements             *TX  `hl7:"false,Special Handling Requirements"`                   // OM4-10
	NormalCollectionVolume                  *CQ  `hl7:"false,Normal Collection Volume"`                        // OM4-11
	MinimumCollectionVolume                 *CQ  `hl7:"false,Minimum Collection Volume"`                       // OM4-12
	SpecimenRequirements                    *TX  `hl7:"false,Specimen Requirements"`                           // OM4-13
	SpecimenPriorities                      []ID `hl7:"false,Specimen Priorities"`                             // OM4-14
	SpecimenRetentionTime                   *CQ  `hl7:"false,Specimen Retention Time"`                         // OM4-15
}

func (s *OM4) SegmentName() string {
	return "OM4"
}

// OM5 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type OM5 struct {
	SegmentTypeID                                       *ST  `hl7:"false,Segment Type ID"`                                              // OM5-1
	SequenceNumberTestObservationMasterFile             *NM  `hl7:"false,Sequence Number - Test/ Observation Master File"`              // OM5-2
	TestsObservationsIncludedWithinAnOrderedTestBattery []CE `hl7:"false,Tests / Observations Included Within An Ordered Test Battery"` // OM5-3
	ObservationIDSuffixes                               *ST  `hl7:"false,Observation ID Suffixes"`                                      // OM5-4
}

func (s *OM5) SegmentName() string {
	return "OM5"
}

// OM6 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type OM6 struct {
	SegmentTypeID                           *ST `hl7:"false,Segment Type ID"`                                 // OM6-1
	SequenceNumberTestObservationMasterFile *NM `hl7:"false,Sequence Number - Test/ Observation Master File"` // OM6-2
	DerivationRule                          *TX `hl7:"false,Derivation Rule"`                                 // OM6-3
}

func (s *OM6) SegmentName() string {
	return "OM6"
}

// ORC represents the corresponding HL7 segment.
// Definition from HL7 2.2
type ORC struct {
	OrderControl           *ID  `hl7:"true,Order Control"`                // ORC-1
	PlacerOrderNumber      *CM  `hl7:"false,Placer Order Number"`         // ORC-2
	FillerOrderNumber      *CM  `hl7:"false,Filler Order Number"`         // ORC-3
	PlacerGroupNumber      *CM  `hl7:"false,Placer Group Number"`         // ORC-4
	OrderStatus            *ID  `hl7:"false,Order Status"`                // ORC-5
	ResponseFlag           *ID  `hl7:"false,Response Flag"`               // ORC-6
	QuantityTiming         []TQ `hl7:"false,Quantity / Timing"`           // ORC-7
	Parent                 *CM  `hl7:"false,Parent"`                      // ORC-8
	DateTimeOfTransaction  *TS  `hl7:"false,Date / Time Of Transaction"`  // ORC-9
	EnteredBy              *CN  `hl7:"false,Entered By"`                  // ORC-10
	VerifiedBy             *CN  `hl7:"false,Verified By"`                 // ORC-11
	OrderingProvider       *CN  `hl7:"false,Ordering Provider"`           // ORC-12
	EntererSLocation       *CM  `hl7:"false,Enterer'S Location"`          // ORC-13
	CallBackPhoneNumber    []TN `hl7:"false,Call Back Phone Number"`      // ORC-14
	OrderEffectiveDateTime *TS  `hl7:"false,Order Effective Date / Time"` // ORC-15
	OrderControlCodeReason *CE  `hl7:"false,Order Control Code Reason"`   // ORC-16
	EnteringOrganization   *CE  `hl7:"false,Entering Organization"`       // ORC-17
	EnteringDevice         *CE  `hl7:"false,Entering Device"`             // ORC-18
	ActionBy               *CN  `hl7:"false,Action By"`                   // ORC-19
}

func (s *ORC) SegmentName() string {
	return "ORC"
}

// ORO represents the corresponding HL7 segment.
// Definition from HL7 2.1
type ORO struct {
	ORDERITEMID       *CE  `hl7:"false,ORDER ITEM ID"`      // ORO-1
	SUBSTITUTEALLOWED *ID  `hl7:"false,SUBSTITUTE ALLOWED"` // ORO-2
	RESULTSCOPIESTO   []CN `hl7:"false,RESULTS COPIES TO"`  // ORO-3
	STOCKLOCATION     *ID  `hl7:"false,STOCK LOCATION"`     // ORO-4
}

func (s *ORO) SegmentName() string {
	return "ORO"
}

// PID represents the corresponding HL7 segment.
// Definition from HL7 2.2
type PID struct {
	SetIDPatientID              *SI  `hl7:"false,Set ID - Patient ID"`               // PID-1
	PatientIDExternalID         *CK  `hl7:"false,Patient ID (External ID)"`          // PID-2
	PatientIDInternalID         []CM `hl7:"true,Patient ID (Internal ID)"`           // PID-3
	AlternatePatientID          *ST  `hl7:"false,Alternate Patient ID"`              // PID-4
	PatientName                 *PN  `hl7:"true,Patient Name"`                       // PID-5
	MotherSMaidenName           *ST  `hl7:"false,Mother'S Maiden Name"`              // PID-6
	DateOfBirth                 *TS  `hl7:"false,Date Of Birth"`                     // PID-7
	Sex                         *ID  `hl7:"false,Sex"`                               // PID-8
	PatientAlias                []PN `hl7:"false,Patient Alias"`                     // PID-9
	Race                        *CWE `hl7:"false,Race"`                              // PID-10
	PatientAddress              []AD `hl7:"false,Patient Address"`                   // PID-11
	CountyCode                  *ID  `hl7:"false,County Code"`                       // PID-12
	PhoneNumberHome             []TN `hl7:"false,Phone Number - Home"`               // PID-13
	PhoneNumberBusiness         []TN `hl7:"false,Phone Number - Business"`           // PID-14
	LanguagePatient             *ST  `hl7:"false,Language - Patient"`                // PID-15
	MaritalStatus               *ID  `hl7:"false,Marital Status"`                    // PID-16
	Religion                    *ID  `hl7:"false,Religion"`                          // PID-17
	PatientAccountNumber        *CK  `hl7:"false,Patient Account Number"`            // PID-18
	SocialSecurityNumberPatient *ST  `hl7:"false,Social Security Number - Patient"`  // PID-19
	DriverSLicenseNumberPatient *CM  `hl7:"false,Driver'S License Number - Patient"` // PID-20
	MotherSIdentifier           *CK  `hl7:"false,Mother'S Identifier"`               // PID-21
	EthnicGroup                 *CWE `hl7:"false,Ethnic Group"`                      // PID-22
	BirthPlace                  *ST  `hl7:"false,Birth Place"`                       // PID-23
	MultipleBirthIndicator      *ID  `hl7:"false,Multiple Birth Indicator"`          // PID-24
	BirthOrder                  *NM  `hl7:"false,Birth Order"`                       // PID-25
	Citizenship                 []ID `hl7:"false,Citizenship"`                       // PID-26
	VeteransMilitaryStatus      *ST  `hl7:"false,Veterans Military Status"`          // PID-27
}

func (s *PID) SegmentName() string {
	return "PID"
}

// PR1 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type PR1 struct {
	SetIDProcedure        *SI  `hl7:"true,Set ID - Procedure"`      // PR1-1
	ProcedureCodingMethod []ID `hl7:"true,Procedure Coding Method"` // PR1-2
	ProcedureCode         []ID `hl7:"true,Procedure Code"`          // PR1-3
	ProcedureDescription  []ST `hl7:"false,Procedure Description"`  // PR1-4
	ProcedureDateTime     *TS  `hl7:"true,Procedure Date / Time"`   // PR1-5
	ProcedureType         *ID  `hl7:"true,Procedure Type"`          // PR1-6
	ProcedureMinutes      *NM  `hl7:"false,Procedure Minutes"`      // PR1-7
	Anesthesiologist      *CN  `hl7:"false,Anesthesiologist"`       // PR1-8
	AnesthesiaCode        *ID  `hl7:"false,Anesthesia Code"`        // PR1-9
	AnesthesiaMinutes     *NM  `hl7:"false,Anesthesia Minutes"`     // PR1-10
	Surgeon               *CN  `hl7:"false,Surgeon"`                // PR1-11
	ProcedurePractitioner []CM `hl7:"false,Procedure Practitioner"` // PR1-12
	ConsentCode           *ID  `hl7:"false,Consent Code"`           // PR1-13
	ProcedurePriority     *NM  `hl7:"false,Procedure Priority"`     // PR1-14
}

func (s *PR1) SegmentName() string {
	return "PR1"
}

// PRA represents the corresponding HL7 segment.
// Definition from HL7 2.2
type PRA struct {
	PRAPrimaryKeyValue    *ST  `hl7:"true,PRA - Primary Key Value"`  // PRA-1
	PractitionerGroup     []CE `hl7:"false,Practitioner Group"`      // PRA-2
	PractitionerCategory  []ID `hl7:"false,Practitioner Category"`   // PRA-3
	ProviderBilling       *ID  `hl7:"false,Provider Billing"`        // PRA-4
	Specialty             []CM `hl7:"false,Specialty"`               // PRA-5
	PractitionerIDNumbers []CM `hl7:"false,Practitioner ID Numbers"` // PRA-6
	Privileges            []CM `hl7:"false,Privileges"`              // PRA-7
}

func (s *PRA) SegmentName() string {
	return "PRA"
}

// PV1 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type PV1 struct {
	SetIDPatientVisit       *SI  `hl7:"false,Set ID - Patient Visit"`      // PV1-1
	PatientClass            *ID  `hl7:"true,Patient Class"`                // PV1-2
	AssignedPatientLocation *CM  `hl7:"false,Assigned Patient Location"`   // PV1-3
	AdmissionType           *ID  `hl7:"false,Admission Type"`              // PV1-4
	PreadmitNumber          *ST  `hl7:"false,Preadmit Number"`             // PV1-5
	PriorPatientLocation    *CM  `hl7:"false,Prior Patient Location"`      // PV1-6
	AttendingDoctor         *CN  `hl7:"false,Attending Doctor"`            // PV1-7
	ReferringDoctor         *CN  `hl7:"false,Referring Doctor"`            // PV1-8
	ConsultingDoctor        []CN `hl7:"false,Consulting Doctor"`           // PV1-9
	HospitalService         *ID  `hl7:"false,Hospital Service"`            // PV1-10
	TemporaryLocation       *CM  `hl7:"false,Temporary Location"`          // PV1-11
	PreadmitTestIndicator   *ID  `hl7:"false,Preadmit Test Indicator"`     // PV1-12
	ReadmissionIndicator    *ID  `hl7:"false,Readmission Indicator"`       // PV1-13
	AdmitSource             *ID  `hl7:"false,Admit Source"`                // PV1-14
	AmbulatoryStatus        []ID `hl7:"false,Ambulatory Status"`           // PV1-15
	VIPIndicator            *ID  `hl7:"false,VIP Indicator"`               // PV1-16
	AdmittingDoctor         *CN  `hl7:"false,Admitting Doctor"`            // PV1-17
	PatientType             *ID  `hl7:"false,Patient Type"`                // PV1-18
	VisitNumber             *CM  `hl7:"false,Visit Number"`                // PV1-19
	FinancialClass          []CM `hl7:"false,Financial Class"`             // PV1-20
	ChargePriceIndicator    *ID  `hl7:"false,Charge Price Indicator"`      // PV1-21
	CourtesyCode            *ID  `hl7:"false,Courtesy Code"`               // PV1-22
	CreditRating            *ID  `hl7:"false,Credit Rating"`               // PV1-23
	ContractCode            []ID `hl7:"false,Contract Code"`               // PV1-24
	ContractEffectiveDate   []DT `hl7:"false,Contract Effective Date"`     // PV1-25
	ContractAmount          []NM `hl7:"false,Contract Amount"`             // PV1-26
	ContractPeriod          []NM `hl7:"false,Contract Period"`             // PV1-27
	InterestCode            *ID  `hl7:"false,Interest Code"`               // PV1-28
	TransferToBadDebtCode   *ID  `hl7:"false,Transfer To Bad Debt - Code"` // PV1-29
	TransferToBadDebtDate   *DT  `hl7:"false,Transfer To Bad Debt - Date"` // PV1-30
	BadDebtAgencyCode       *ID  `hl7:"false,Bad Debt Agency Code"`        // PV1-31
	BadDebtTransferAmount   *NM  `hl7:"false,Bad Debt Transfer Amount"`    // PV1-32
	BadDebtRecoveryAmount   *NM  `hl7:"false,Bad Debt Recovery Amount"`    // PV1-33
	DeleteAccountIndicator  *ID  `hl7:"false,Delete Account Indicator"`    // PV1-34
	DeleteAccountDate       *DT  `hl7:"false,Delete Account Date"`         // PV1-35
	DischargeDisposition    *ID  `hl7:"false,Discharge Disposition"`       // PV1-36
	DischargedToLocation    *CM  `hl7:"false,Discharged To Location"`      // PV1-37
	DietType                *ID  `hl7:"false,Diet Type"`                   // PV1-38
	ServicingFacility       *ID  `hl7:"false,Servicing Facility"`          // PV1-39
	BedStatus               *ID  `hl7:"false,Bed Status"`                  // PV1-40
	AccountStatus           *ID  `hl7:"false,Account Status"`              // PV1-41
	PendingLocation         *CM  `hl7:"false,Pending Location"`            // PV1-42
	PriorTemporaryLocation  *CM  `hl7:"false,Prior Temporary Location"`    // PV1-43
	AdmitDateTime           *TS  `hl7:"false,Admit Date / Time"`           // PV1-44
	DischargeDateTime       *TS  `hl7:"false,Discharge Date / Time"`       // PV1-45
	CurrentPatientBalance   *NM  `hl7:"false,Current Patient Balance"`     // PV1-46
	TotalCharges            *NM  `hl7:"false,Total Charges"`               // PV1-47
	TotalAdjustments        *NM  `hl7:"false,Total Adjustments"`           // PV1-48
	TotalPayments           *NM  `hl7:"false,Total Payments"`              // PV1-49
	AlternateVisitID        *CM  `hl7:"false,Alternate Visit ID"`          // PV1-50
}

func (s *PV1) SegmentName() string {
	return "PV1"
}

// PV2 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type PV2 struct {
	PriorPendingLocation     *CM  `hl7:"false,Prior Pending Location"`     // PV2-1
	AccommodationCode        *CE  `hl7:"false,Accommodation Code"`         // PV2-2
	AdmitReason              *CE  `hl7:"false,Admit Reason"`               // PV2-3
	TransferReason           *CE  `hl7:"false,Transfer Reason"`            // PV2-4
	PatientValuables         []ST `hl7:"false,Patient Valuables"`          // PV2-5
	PatientValuablesLocation *ST  `hl7:"false,Patient Valuables Location"` // PV2-6
	VisitUserCode            *ID  `hl7:"false,Visit User Code"`            // PV2-7
	ExpectedAdmitDate        *DT  `hl7:"false,Expected Admit Date"`        // PV2-8
	ExpectedDischargeDate    *DT  `hl7:"false,Expected Discharge Date"`    // PV2-9
}

func (s *PV2) SegmentName() string {
	return "PV2"
}

// QRD represents the corresponding HL7 segment.
// Definition from HL7 2.2
type QRD struct {
	QueryDateTime              *TS  `hl7:"true,Query Date / Time"`               // QRD-1
	QueryFormatCode            *ID  `hl7:"true,Query Format Code"`               // QRD-2
	QueryPriority              *ID  `hl7:"true,Query Priority"`                  // QRD-3
	QueryID                    *ST  `hl7:"true,Query ID"`                        // QRD-4
	DeferredResponseType       *ID  `hl7:"false,Deferred Response Type"`         // QRD-5
	DeferredResponseDateTime   *TS  `hl7:"false,Deferred Response Date / Time"`  // QRD-6
	QuantityLimitedRequest     *CQ  `hl7:"true,Quantity Limited Request"`        // QRD-7
	WhoSubjectFilter           []ST `hl7:"true,Who Subject Filter"`              // QRD-8
	WhatSubjectFilter          []ID `hl7:"true,What Subject Filter"`             // QRD-9
	WhatDepartmentDataCode     []ST `hl7:"true,What Department Data Code"`       // QRD-10
	WhatDataCodeValueQualifier []CM `hl7:"false,What Data Code Value Qualifier"` // QRD-11
	QueryResultsLevel          *ID  `hl7:"false,Query Results Level"`            // QRD-12
}

func (s *QRD) SegmentName() string {
	return "QRD"
}

// QRF represents the corresponding HL7 segment.
// Definition from HL7 2.2
type QRF struct {
	WhereSubjectFilter           []ST `hl7:"true,Where Subject Filter"`                // QRF-1
	WhenDataStartDateTime        *TS  `hl7:"false,When Data Start Date / Time"`        // QRF-2
	WhenDataEndDateTime          *TS  `hl7:"false,When Data End Date / Time"`          // QRF-3
	WhatUserQualifier            []ST `hl7:"false,What User Qualifier"`                // QRF-4
	OtherQRYSubjectFilter        []ST `hl7:"false,Other QRY Subject Filter"`           // QRF-5
	WhichDateTimeQualifier       []ID `hl7:"false,Which Date / Time Qualifier"`        // QRF-6
	WhichDateTimeStatusQualifier []ID `hl7:"false,Which Date / Time Status Qualifier"` // QRF-7
	DateTimeSelectionQualifier   []ID `hl7:"false,Date / Time Selection Qualifier"`    // QRF-8
}

func (s *QRF) SegmentName() string {
	return "QRF"
}

// RQ1 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type RQ1 struct {
	AnticipatedPrice     *SI `hl7:"false,Anticipated Price"`      // RQ1-1
	ManufacturerID       *CE `hl7:"false,Manufacturer ID"`        // RQ1-2
	ManufacturerSCatalog *ST `hl7:"false,Manufacturer'S Catalog"` // RQ1-3
	VendorID             *CE `hl7:"false,Vendor ID"`              // RQ1-4
	VendorCatalog        *ST `hl7:"false,Vendor Catalog"`         // RQ1-5
	Taxable              *ID `hl7:"false,Taxable"`                // RQ1-6
	SubstituteAllowed    *ID `hl7:"false,Substitute Allowed"`     // RQ1-7
}

func (s *RQ1) SegmentName() string {
	return "RQ1"
}

// RQD represents the corresponding HL7 segment.
// Definition from HL7 2.2
type RQD struct {
	RequisitionLineNumber    *SI `hl7:"false,Requisition Line Number"`     // RQD-1
	ItemCodeInternal         *CE `hl7:"false,Item Code - Internal"`        // RQD-2
	ItemCodeExternal         *CE `hl7:"false,Item Code - External"`        // RQD-3
	HospitalItemCode         *CE `hl7:"false,Hospital Item Code"`          // RQD-4
	RequisitionQuantity      *NM `hl7:"false,Requisition Quantity"`        // RQD-5
	RequisitionUnitOfMeasure *CE `hl7:"false,Requisition Unit Of Measure"` // RQD-6
	DepartmentCostCenter     *ID `hl7:"false,Department Cost Center"`      // RQD-7
	ItemNaturalAccountCode   *ID `hl7:"false,Item Natural Account Code"`   // RQD-8
	DeliverToID              *CE `hl7:"false,Deliver-To ID"`               // RQD-9
	DateNeeded               *DT `hl7:"false,Date Needed"`                 // RQD-10
}

func (s *RQD) SegmentName() string {
	return "RQD"
}

// RX1 represents the corresponding HL7 segment.
// Definition from HL7 2.1
type RX1 struct {
	UNUSED1              *ST  `hl7:"false,UNUSED"`                   // RX1-1
	UNUSED2              *ST  `hl7:"false,UNUSED"`                   // RX1-2
	ROUTE                *ST  `hl7:"false,ROUTE"`                    // RX1-3
	SITEADMINISTERED     *ST  `hl7:"false,SITE ADMINISTERED"`        // RX1-4
	IVSOLUTIONRATE       *CQ  `hl7:"false,IV SOLUTION RATE"`         // RX1-5
	DRUGSTRENGTH         *CQ  `hl7:"false,DRUG STRENGTH"`            // RX1-6
	FINALCONCENTRATION   *NM  `hl7:"false,FINAL CONCENTRATION"`      // RX1-7
	FINALVOLUMEINML      *NM  `hl7:"false,FINAL VOLUME IN ML."`      // RX1-8
	DRUGDOSE             *CM  `hl7:"false,DRUG DOSE"`                // RX1-9
	DRUGROLE             *ID  `hl7:"false,DRUG ROLE"`                // RX1-10
	PRESCRIPTIONSEQUENCE *NM  `hl7:"false,PRESCRIPTION SEQUENCE #"`  // RX1-11
	QUANTITYDISPENSED    *CQ  `hl7:"false,QUANTITY DISPENSED"`       // RX1-12
	UNUSED3              *ST  `hl7:"false,UNUSED"`                   // RX1-13
	DRUGID               *CE  `hl7:"false,DRUG ID"`                  // RX1-14
	COMPONENTDRUGIDS     []ID `hl7:"false,COMPONENT DRUG IDS"`       // RX1-15
	PRESCRIPTIONTYPE     *ID  `hl7:"false,PRESCRIPTION TYPE"`        // RX1-16
	SUBSTITUTIONSTATUS   *ID  `hl7:"false,SUBSTITUTION STATUS"`      // RX1-17
	RXORDERSTATUS        *ID  `hl7:"false,RX ORDER STATUS"`          // RX1-18
	NUMBEROFREFILLS      *NM  `hl7:"false,NUMBER OF REFILLS"`        // RX1-19
	UNUSED4              *ST  `hl7:"false,UNUSED"`                   // RX1-20
	REFILLSREMAINING     *NM  `hl7:"false,REFILLS REMAINING"`        // RX1-21
	DEACLASS             *ID  `hl7:"false,DEA CLASS"`                // RX1-22
	ORDERINGMDSDEANUMBER *NM  `hl7:"false,ORDERING MD'S DEA NUMBER"` // RX1-23
	UNUSED5              *ST  `hl7:"false,UNUSED"`                   // RX1-24
	LASTREFILLDATETIME   *TS  `hl7:"false,LAST REFILL DATE/TIME"`    // RX1-25
	RXNUMBER             *ST  `hl7:"false,RX NUMBER"`                // RX1-26
	PRNSTATUS            *ID  `hl7:"false,PRN STATUS"`               // RX1-27
	PHARMACYINSTRUCTIONS []TX `hl7:"false,PHARMACY INSTRUCTIONS"`    // RX1-28
	PATIENTINSTRUCTIONS  []TX `hl7:"false,PATIENT INSTRUCTIONS"`     // RX1-29
	INSTRUCTIONS         []TX `hl7:"false,INSTRUCTIONS"`             // RX1-30
}

func (s *RX1) SegmentName() string {
	return "RX1"
}

// RXA represents the corresponding HL7 segment.
// Definition from HL7 2.2
type RXA struct {
	GiveSubIDCounter              *NM `hl7:"true,Give Sub-ID Counter"`                 // RXA-1
	AdministrationSubIDCounter    *NM `hl7:"true,Administration Sub-ID Counter"`       // RXA-2
	DateTimeStartOfAdministration *TS `hl7:"true,Date / Time Start Of Administration"` // RXA-3
	DateTimeEndOfAdministration   *TS `hl7:"true,Date / Time End Of Administration"`   // RXA-4
	AdministeredCode              *CE `hl7:"true,Administered Code"`                   // RXA-5
	AdministeredAmount            *NM `hl7:"true,Administered Amount"`                 // RXA-6
	AdministeredUnits             *CE `hl7:"false,Administered Units"`                 // RXA-7
	AdministeredDosageForm        *CE `hl7:"false,Administered Dosage Form"`           // RXA-8
	AdministrationNotes           *ST `hl7:"false,Administration Notes"`               // RXA-9
	AdministeringProvider         *CN `hl7:"false,Administering Provider"`             // RXA-10
	AdministeredAtLocation        *CM `hl7:"false,Administered-At Location"`           // RXA-11
	AdministeredPerTimeUnit       *ST `hl7:"false,Administered Per (Time Unit)"`       // RXA-12
}

func (s *RXA) SegmentName() string {
	return "RXA"
}

// RXC represents the corresponding HL7 segment.
// Definition from HL7 2.2
type RXC struct {
	RXComponentType *ID `hl7:"true,RX Component Type"` // RXC-1
	ComponentCode   *CE `hl7:"true,Component Code"`    // RXC-2
	ComponentAmount *NM `hl7:"true,Component Amount"`  // RXC-3
	ComponentUnits  *CE `hl7:"true,Component Units"`   // RXC-4
}

func (s *RXC) SegmentName() string {
	return "RXC"
}

// RXD represents the corresponding HL7 segment.
// Definition from HL7 2.2
type RXD struct {
	DispenseSubIDCounter                  *NM  `hl7:"false,Dispense Sub-ID Counter"`                  // RXD-1
	DispenseGiveCode                      *CE  `hl7:"true,Dispense / Give Code"`                      // RXD-2
	DateTimeDispensed                     *TS  `hl7:"false,Date / Time Dispensed"`                    // RXD-3
	ActualDispenseAmount                  *NM  `hl7:"true,Actual Dispense Amount"`                    // RXD-4
	ActualDispenseUnits                   *CE  `hl7:"false,Actual Dispense Units"`                    // RXD-5
	ActualDosageForm                      *CE  `hl7:"false,Actual Dosage Form"`                       // RXD-6
	PrescriptionNumber                    *ST  `hl7:"true,Prescription Number"`                       // RXD-7
	NumberOfRefillsRemaining              *NM  `hl7:"false,Number Of Refills Remaining"`              // RXD-8
	DispenseNotes                         []ST `hl7:"false,Dispense Notes"`                           // RXD-9
	DispensingProvider                    *CN  `hl7:"false,Dispensing Provider"`                      // RXD-10
	SubstitutionStatus                    *ID  `hl7:"false,Substitution Status"`                      // RXD-11
	TotalDailyDose                        *CQ  `hl7:"false,Total Daily Dose"`                         // RXD-12
	DeliverToLocation                     *CM  `hl7:"false,Deliver-To Location"`                      // RXD-13
	NeedsHumanReview                      *ID  `hl7:"false,Needs Human Review"`                       // RXD-14
	PharmacySpecialDispensingInstructions *CE  `hl7:"false,Pharmacy Special Dispensing Instructions"` // RXD-15
}

func (s *RXD) SegmentName() string {
	return "RXD"
}

// RXE represents the corresponding HL7 segment.
// Definition from HL7 2.2
type RXE struct {
	QuantityTiming                            []TQ `hl7:"false,Quantity / Timing"`                                   // RXE-1
	GiveCode                                  *CE  `hl7:"true,Give Code"`                                            // RXE-2
	GiveAmountMinimum                         *NM  `hl7:"true,Give Amount - Minimum"`                                // RXE-3
	GiveAmountMaximum                         *NM  `hl7:"false,Give Amount - Maximum"`                               // RXE-4
	GiveUnits                                 *CE  `hl7:"true,Give Units"`                                           // RXE-5
	GiveDosageForm                            *CE  `hl7:"false,Give Dosage Form"`                                    // RXE-6
	ProviderSAdministrationInstructions       []CE `hl7:"false,Provider'S Administration Instructions"`              // RXE-7
	DeliverToLocation                         *CM  `hl7:"false,Deliver-To Location"`                                 // RXE-8
	SubstitutionStatus                        *ID  `hl7:"false,Substitution Status"`                                 // RXE-9
	DispenseAmount                            *NM  `hl7:"false,Dispense Amount"`                                     // RXE-10
	DispenseUnits                             *CE  `hl7:"false,Dispense Units"`                                      // RXE-11
	NumberOfRefills                           *NM  `hl7:"false,Number Of Refills"`                                   // RXE-12
	OrderingProviderSDEANumber                *CN  `hl7:"false,Ordering Provider'S DEA Number"`                      // RXE-13
	PharmacistVerifierID                      *CN  `hl7:"false,Pharmacist Verifier ID"`                              // RXE-14
	PrescriptionNumber                        *ST  `hl7:"true,Prescription Number"`                                  // RXE-15
	NumberOfRefillsRemaining                  *NM  `hl7:"false,Number Of Refills Remaining"`                         // RXE-16
	NumberOfRefillsDosesDispensed             *NM  `hl7:"false,Number Of Refills / Doses Dispensed"`                 // RXE-17
	DateTimeOfMostRecentRefillOrDoseDispensed *TS  `hl7:"false,Date / Time Of Most Recent Refill Or Dose Dispensed"` // RXE-18
	TotalDailyDose                            *CQ  `hl7:"false,Total Daily Dose"`                                    // RXE-19
	NeedsHumanReview                          *ID  `hl7:"false,Needs Human Review"`                                  // RXE-20
	PharmacySpecialDispensingInstructions     *CE  `hl7:"false,Pharmacy Special Dispensing Instructions"`            // RXE-21
	GivePerTimeUnit                           *ST  `hl7:"false,Give Per (Time Unit)"`                                // RXE-22
	GiveRateAmount                            *CE  `hl7:"false,Give Rate Amount"`                                    // RXE-23
	GiveRateUnits                             *CE  `hl7:"false,Give Rate Units"`                                     // RXE-24
}

func (s *RXE) SegmentName() string {
	return "RXE"
}

// RXG represents the corresponding HL7 segment.
// Definition from HL7 2.2
type RXG struct {
	GiveSubIDCounter                          *NM  `hl7:"true,Give Sub-ID Counter"`                           // RXG-1
	DispenseSubIDCounter                      *NM  `hl7:"false,Dispense Sub-ID Counter"`                      // RXG-2
	QuantityTiming                            []TQ `hl7:"false,Quantity / Timing"`                            // RXG-3
	GiveCode                                  *CE  `hl7:"true,Give Code"`                                     // RXG-4
	GiveAmountMinimum                         *NM  `hl7:"true,Give Amount - Minimum"`                         // RXG-5
	GiveAmountMaximum                         *NM  `hl7:"false,Give Amount - Maximum"`                        // RXG-6
	GiveUnits                                 *CE  `hl7:"true,Give Units"`                                    // RXG-7
	GiveDosageForm                            *CE  `hl7:"false,Give Dosage Form"`                             // RXG-8
	AdministrationNotes                       *ST  `hl7:"false,Administration Notes"`                         // RXG-9
	SubstitutionStatus                        *ID  `hl7:"false,Substitution Status"`                          // RXG-10
	DeliverToLocation                         *CM  `hl7:"false,Deliver-To Location"`                          // RXG-11
	NeedsHumanReview                          *ID  `hl7:"false,Needs Human Review"`                           // RXG-12
	PharmacySpecialAdministrationInstructions []CE `hl7:"false,Pharmacy Special Administration Instructions"` // RXG-13
	GivePerTimeUnit                           *ST  `hl7:"false,Give Per (Time Unit)"`                         // RXG-14
	GiveRateAmount                            *CE  `hl7:"false,Give Rate Amount"`                             // RXG-15
	GiveRateUnits                             *CE  `hl7:"false,Give Rate Units"`                              // RXG-16
}

func (s *RXG) SegmentName() string {
	return "RXG"
}

// RXO represents the corresponding HL7 segment.
// Definition from HL7 2.2
type RXO struct {
	RequestedGiveCode                   *CE  `hl7:"true,Requested Give Code"`                     // RXO-1
	RequestedGiveAmountMinimum          *NM  `hl7:"true,Requested Give Amount - Minimum"`         // RXO-2
	RequestedGiveAmountMaximum          *NM  `hl7:"false,Requested Give Amount - Maximum"`        // RXO-3
	RequestedGiveUnits                  *CE  `hl7:"true,Requested Give Units"`                    // RXO-4
	RequestedDosageForm                 *CE  `hl7:"false,Requested Dosage Form"`                  // RXO-5
	ProviderSPharmacyInstructions       []CE `hl7:"false,Provider'S Pharmacy Instructions"`       // RXO-6
	ProviderSAdministrationInstructions []CE `hl7:"false,Provider'S Administration Instructions"` // RXO-7
	DeliverToLocation                   *CM  `hl7:"false,Deliver-To Location"`                    // RXO-8
	AllowSubstitutions                  *ID  `hl7:"false,Allow Substitutions"`                    // RXO-9
	RequestedDispenseCode               *CE  `hl7:"false,Requested Dispense Code"`                // RXO-10
	RequestedDispenseAmount             *NM  `hl7:"false,Requested Dispense Amount"`              // RXO-11
	RequestedDispenseUnits              *CE  `hl7:"false,Requested Dispense Units"`               // RXO-12
	NumberOfRefills                     *NM  `hl7:"false,Number Of Refills"`                      // RXO-13
	OrderingProviderSDEANumber          *CN  `hl7:"false,Ordering Provider'S DEA Number"`         // RXO-14
	PharmacistVerifierID                *CN  `hl7:"false,Pharmacist Verifier ID"`                 // RXO-15
	NeedsHumanReview                    *ID  `hl7:"false,Needs Human Review"`                     // RXO-16
	RequestedGivePerTimeUnit            *ST  `hl7:"false,Requested Give Per (Time Unit)"`         // RXO-17
}

func (s *RXO) SegmentName() string {
	return "RXO"
}

// RXR represents the corresponding HL7 segment.
// Definition from HL7 2.2
type RXR struct {
	Route                *CE `hl7:"true,Route"`                  // RXR-1
	Site                 *CE `hl7:"false,Site"`                  // RXR-2
	AdministrationDevice *CE `hl7:"false,Administration Device"` // RXR-3
	AdministrationMethod *CE `hl7:"false,Administration Method"` // RXR-4
}

func (s *RXR) SegmentName() string {
	return "RXR"
}

// STF represents the corresponding HL7 segment.
// Definition from HL7 2.2
type STF struct {
	STFPrimaryKeyValue       *CE  `hl7:"true,STF - Primary Key Value"`      // STF-1
	StaffIDCode              []CE `hl7:"false,Staff ID Code"`               // STF-2
	StaffName                *PN  `hl7:"false,Staff Name"`                  // STF-3
	StaffType                []ID `hl7:"false,Staff Type"`                  // STF-4
	Sex                      *ID  `hl7:"false,Sex"`                         // STF-5
	DateOfBirth              *TS  `hl7:"false,Date Of Birth"`               // STF-6
	ActiveInactive           *ID  `hl7:"false,Active / Inactive"`           // STF-7
	Department               []CE `hl7:"false,Department"`                  // STF-8
	Service                  []CE `hl7:"false,Service"`                     // STF-9
	Phone                    []TN `hl7:"false,Phone"`                       // STF-10
	OfficeHomeAddress        []AD `hl7:"false,Office / Home Address"`       // STF-11
	ActivationDate           []CM `hl7:"false,Activation Date"`             // STF-12
	InactivationDate         []CM `hl7:"false,Inactivation Date"`           // STF-13
	BackupPersonID           []CE `hl7:"false,Backup Person ID"`            // STF-14
	EMailAddress             []ST `hl7:"false,E-Mail Address"`              // STF-15
	PreferredMethodOfContact *ID  `hl7:"false,Preferred Method Of Contact"` // STF-16
}

func (s *STF) SegmentName() string {
	return "STF"
}

// UB1 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type UB1 struct {
	SetIDUB82                 *SI  `hl7:"false,Set ID - UB82"`                       // UB1-1
	BloodDeductible43         *NM  `hl7:"false,Blood Deductible (43)"`               // UB1-2
	BloodFurnishedPintsOf40   *NM  `hl7:"false,Blood Furnished Pints Of (40)"`       // UB1-3
	BloodReplacedPints41      *NM  `hl7:"false,Blood Replaced Pints (41)"`           // UB1-4
	BloodNotReplacedPints42   *NM  `hl7:"false,Blood Not Replaced Pints (42)"`       // UB1-5
	CoInsuranceDays25         *NM  `hl7:"false,Co-Insurance Days (25)"`              // UB1-6
	ConditionCode3539         []ID `hl7:"false,Condition Code (35-39)"`              // UB1-7
	CoveredDays23             *NM  `hl7:"false,Covered Days (23)"`                   // UB1-8
	NonCoveredDays24          *NM  `hl7:"false,Non-Covered Days (24)"`               // UB1-9
	ValueAmountAndCode4649    []CM `hl7:"false,Value Amount And Code (46-49)"`       // UB1-10
	NumberOfGraceDays90       *NM  `hl7:"false,Number Of Grace Days (90)"`           // UB1-11
	SpecialProgramIndicator44 *ID  `hl7:"false,Special Program Indicator (44)"`      // UB1-12
	PSROURApprovalIndicator87 *ID  `hl7:"false,PSRO / UR Approval Indicator (87)"`   // UB1-13
	PSROURApprovedStayFrom88  *DT  `hl7:"false,PSRO / UR Approved Stay - From (88)"` // UB1-14
	PSROURApprovedStayTo89    *DT  `hl7:"false,PSRO / UR Approved Stay - To (89)"`   // UB1-15
	Occurrence2832            []CM `hl7:"false,Occurrence (28-32)"`                  // UB1-16
	OccurrenceSpan33          *ID  `hl7:"false,Occurrence Span (33)"`                // UB1-17
	OccurrenceSpanStartDate33 *DT  `hl7:"false,Occurrence Span Start Date (33)"`     // UB1-18
	OccurrenceSpanEndDate33   *DT  `hl7:"false,Occurrence Span End Date (33)"`       // UB1-19
	UB82Locator2              *ST  `hl7:"false,UB-82 Locator 2"`                     // UB1-20
	UB82Locator9              *ST  `hl7:"false,UB-82 Locator 9"`                     // UB1-21
	UB82Locator27             *ST  `hl7:"false,UB-82 Locator 27"`                    // UB1-22
	UB82Locator45             *ST  `hl7:"false,UB-82 Locator 45"`                    // UB1-23
}

func (s *UB1) SegmentName() string {
	return "UB1"
}

// UB2 represents the corresponding HL7 segment.
// Definition from HL7 2.2
type UB2 struct {
	SetIDUB92                 *SI  `hl7:"false,Set ID - UB92"`                     // UB2-1
	CoInsuranceDays9          *ST  `hl7:"false,Co-Insurance Days (9)"`             // UB2-2
	ConditionCode2430         []ID `hl7:"false,Condition Code (24-30)"`            // UB2-3
	CoveredDays7              *ST  `hl7:"false,Covered Days (7)"`                  // UB2-4
	NonCoveredDays8           *ST  `hl7:"false,Non-Covered Days (8)"`              // UB2-5
	ValueAmountAndCode3941    []CM `hl7:"false,Value Amount And Code (39-41)"`     // UB2-6
	OccurrenceCodeAndDate3235 []CM `hl7:"false,Occurrence Code And Date (32-35)"`  // UB2-7
	OccurrenceSpanCodeDates36 []CM `hl7:"false,Occurrence Span Code / Dates (36)"` // UB2-8
	UB92Locator2State         []ST `hl7:"false,UB92 Locator 2 (State)"`            // UB2-9
	UB92Locator11State        []ST `hl7:"false,UB92 Locator 11 (State)"`           // UB2-10
	UB92Locator31National     *ST  `hl7:"false,UB92 Locator 31 (National)"`        // UB2-11
	DocumentControlNumber37   []ST `hl7:"false,Document Control Number (37)"`      // UB2-12
	UB92Locator49National     []ST `hl7:"false,UB92 Locator 49 (National)"`        // UB2-13
	UB92Locator56State        []ST `hl7:"false,UB92 Locator 56 (State)"`           // UB2-14
	UB92Locator57National     *ST  `hl7:"false,UB92 Locator 57 (National)"`        // UB2-15
	UB92Locator78State        []ST `hl7:"false,UB92 Locator 78 (State)"`           // UB2-16
}

func (s *UB2) SegmentName() string {
	return "UB2"
}

// URD represents the corresponding HL7 segment.
// Definition from HL7 2.2
type URD struct {
	RUDateTime              *TS  `hl7:"false,R/U Date / Time"`               // URD-1
	ReportPriority          *ID  `hl7:"false,Report Priority"`               // URD-2
	RUWhoSubjectDefinition  []ST `hl7:"true,R/U Who Subject Definition"`     // URD-3
	RUWhatSubjectDefinition []ID `hl7:"false,R/U What Subject Definition"`   // URD-4
	RUWhatDepartmentCode    []ST `hl7:"false,R/U What Department Code"`      // URD-5
	RUDisplayPrintLocations []ST `hl7:"false,R/U Display / Print Locations"` // URD-6
	RUResultsLevel          *ID  `hl7:"false,R/U Results Level"`             // URD-7
}

func (s *URD) SegmentName() string {
	return "URD"
}

// URS represents the corresponding HL7 segment.
// Definition from HL7 2.2
type URS struct {
	RUWhereSubjectDefinition        []ST `hl7:"true,R/U Where Subject Definition"`            // URS-1
	RUWhenDataStartDateTime         *TS  `hl7:"false,R/U When Data Start Date / Time"`        // URS-2
	RUWhenDataEndDateTime           *TS  `hl7:"false,R/U When Data End Date / Time"`          // URS-3
	RUWhatUserQualifier             []ST `hl7:"false,R/U What User Qualifier"`                // URS-4
	RUOtherResultsSubjectDefinition []ST `hl7:"false,R/U Other Results Subject Definition"`   // URS-5
	RUWhichDateTimeQualifier        []ID `hl7:"false,R/U Which Date / Time Qualifier"`        // URS-6
	RUWhichDateTimeStatusQualifier  []ID `hl7:"false,R/U Which Date / Time Status Qualifier"` // URS-7
	RUDateTimeSelectionQualifier    []ID `hl7:"false,R/U Date / Time Selection Qualifier"`    // URS-8
}

func (s *URS) SegmentName() string {
	return "URS"
}

// ACK represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ACK struct {
	MSH   *MSH `hl7:"true,MSH"`
	MSA   *MSA `hl7:"true,MSA"`
	ERR   *ERR `hl7:"false,ERR"`
	Other []interface{}
}

func (s *ACK) MessageTypeName() string {
	return "ACK"
}

// ADR_A19 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADR_A19 struct {
	MSH            *MSH                     `hl7:"true,MSH"`
	MSA            *MSA                     `hl7:"true,MSA"`
	ERR            *ERR                     `hl7:"false,ERR"`
	QRD            *QRD                     `hl7:"true,QRD"`
	QUERY_RESPONSE []ADR_A19_QUERY_RESPONSE `hl7:"true,QUERY_RESPONSE"`
	DSC            *DSC                     `hl7:"false,DSC"`
	Other          []interface{}
}

func (s *ADR_A19) MessageTypeName() string {
	return "ADR_A19"
}

// ADR_A19_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADR_A19_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADR_A19_INSURANCE) MessageTypeSubStructName() string {
	return "ADR_A19_INSURANCE"
}

// ADR_A19_QUERY_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADR_A19_QUERY_RESPONSE struct {
	EVN       *EVN                `hl7:"false,EVN"`
	PID       *PID                `hl7:"true,PID"`
	NK1       []NK1               `hl7:"false,NK1"`
	PV1       *PV1                `hl7:"true,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	PR1       []PR1               `hl7:"false,PR1"`
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []ADR_A19_INSURANCE `hl7:"false,INSURANCE"`
	ACC       *ACC                `hl7:"false,ACC"`
	UB1       *UB1                `hl7:"false,UB1"`
	UB2       *UB2                `hl7:"false,UB2"`
	Other     []interface{}
}

func (s *ADR_A19_QUERY_RESPONSE) MessageTypeSubStructName() string {
	return "ADR_A19_QUERY_RESPONSE"
}

// ADT_A01 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A01 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	EVN       *EVN                `hl7:"true,EVN"`
	PID       *PID                `hl7:"true,PID"`
	NK1       []NK1               `hl7:"false,NK1"`
	PV1       *PV1                `hl7:"true,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	PR1       []PR1               `hl7:"false,PR1"`
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []ADT_A01_INSURANCE `hl7:"false,INSURANCE"`
	ACC       *ACC                `hl7:"false,ACC"`
	UB1       *UB1                `hl7:"false,UB1"`
	UB2       *UB2                `hl7:"false,UB2"`
	Other     []interface{}
}

func (s *ADT_A01) MessageTypeName() string {
	return "ADT_A01"
}

// ADT_A01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A01_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A01_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A01_INSURANCE"
}

// ADT_A02 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A02 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A02) MessageTypeName() string {
	return "ADT_A02"
}

// ADT_A03 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A03 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A03) MessageTypeName() string {
	return "ADT_A03"
}

// ADT_A04 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A04 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	EVN       *EVN                `hl7:"true,EVN"`
	PID       *PID                `hl7:"true,PID"`
	NK1       []NK1               `hl7:"false,NK1"`
	PV1       *PV1                `hl7:"true,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	PR1       []PR1               `hl7:"false,PR1"`
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []ADT_A04_INSURANCE `hl7:"false,INSURANCE"`
	ACC       *ACC                `hl7:"false,ACC"`
	UB1       *UB1                `hl7:"false,UB1"`
	UB2       *UB2                `hl7:"false,UB2"`
	Other     []interface{}
}

func (s *ADT_A04) MessageTypeName() string {
	return "ADT_A04"
}

// ADT_A04_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A04_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A04_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A04_INSURANCE"
}

// ADT_A05 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A05 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	EVN       *EVN                `hl7:"true,EVN"`
	PID       *PID                `hl7:"true,PID"`
	NK1       []NK1               `hl7:"false,NK1"`
	PV1       *PV1                `hl7:"true,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	PR1       []PR1               `hl7:"false,PR1"`
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []ADT_A05_INSURANCE `hl7:"false,INSURANCE"`
	ACC       *ACC                `hl7:"false,ACC"`
	UB1       *UB1                `hl7:"false,UB1"`
	UB2       *UB2                `hl7:"false,UB2"`
	Other     []interface{}
}

func (s *ADT_A05) MessageTypeName() string {
	return "ADT_A05"
}

// ADT_A05_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A05_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A05_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A05_INSURANCE"
}

// ADT_A06 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A06 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	EVN       *EVN                `hl7:"true,EVN"`
	PID       *PID                `hl7:"true,PID"`
	MRG       *MRG                `hl7:"false,MRG"`
	NK1       []NK1               `hl7:"false,NK1"`
	PV1       *PV1                `hl7:"true,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	PR1       []PR1               `hl7:"false,PR1"`
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []ADT_A06_INSURANCE `hl7:"false,INSURANCE"`
	ACC       *ACC                `hl7:"false,ACC"`
	UB1       *UB1                `hl7:"false,UB1"`
	UB2       *UB2                `hl7:"false,UB2"`
	Other     []interface{}
}

func (s *ADT_A06) MessageTypeName() string {
	return "ADT_A06"
}

// ADT_A06_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A06_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A06_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A06_INSURANCE"
}

// ADT_A07 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A07 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	EVN       *EVN                `hl7:"true,EVN"`
	PID       *PID                `hl7:"true,PID"`
	MRG       *MRG                `hl7:"false,MRG"`
	NK1       []NK1               `hl7:"false,NK1"`
	PV1       *PV1                `hl7:"true,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	PR1       []PR1               `hl7:"false,PR1"`
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []ADT_A07_INSURANCE `hl7:"false,INSURANCE"`
	ACC       *ACC                `hl7:"false,ACC"`
	UB1       *UB1                `hl7:"false,UB1"`
	UB2       *UB2                `hl7:"false,UB2"`
	Other     []interface{}
}

func (s *ADT_A07) MessageTypeName() string {
	return "ADT_A07"
}

// ADT_A07_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A07_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A07_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A07_INSURANCE"
}

// ADT_A08 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A08 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	EVN       *EVN                `hl7:"true,EVN"`
	PID       *PID                `hl7:"true,PID"`
	NK1       []NK1               `hl7:"false,NK1"`
	PV1       *PV1                `hl7:"true,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	PR1       []PR1               `hl7:"false,PR1"`
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []ADT_A08_INSURANCE `hl7:"false,INSURANCE"`
	ACC       *ACC                `hl7:"false,ACC"`
	UB1       *UB1                `hl7:"false,UB1"`
	UB2       *UB2                `hl7:"false,UB2"`
	Other     []interface{}
}

func (s *ADT_A08) MessageTypeName() string {
	return "ADT_A08"
}

// ADT_A08_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A08_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A08_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A08_INSURANCE"
}

// ADT_A09 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A09 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	DG1   []DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A09) MessageTypeName() string {
	return "ADT_A09"
}

// ADT_A10 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A10 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	DG1   []DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A10) MessageTypeName() string {
	return "ADT_A10"
}

// ADT_A11 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A11 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	DG1   []DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A11) MessageTypeName() string {
	return "ADT_A11"
}

// ADT_A12 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A12 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	DG1   []DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A12) MessageTypeName() string {
	return "ADT_A12"
}

// ADT_A13 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A13 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	EVN       *EVN                `hl7:"true,EVN"`
	PID       *PID                `hl7:"true,PID"`
	NK1       []NK1               `hl7:"false,NK1"`
	PV1       *PV1                `hl7:"true,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	PR1       []PR1               `hl7:"false,PR1"`
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []ADT_A13_INSURANCE `hl7:"false,INSURANCE"`
	ACC       *ACC                `hl7:"false,ACC"`
	UB1       *UB1                `hl7:"false,UB1"`
	UB2       *UB2                `hl7:"false,UB2"`
	Other     []interface{}
}

func (s *ADT_A13) MessageTypeName() string {
	return "ADT_A13"
}

// ADT_A13_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A13_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A13_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A13_INSURANCE"
}

// ADT_A14 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A14 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	EVN       *EVN                `hl7:"true,EVN"`
	PID       *PID                `hl7:"true,PID"`
	NK1       []NK1               `hl7:"false,NK1"`
	PV1       *PV1                `hl7:"true,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	PR1       []PR1               `hl7:"false,PR1"`
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []ADT_A14_INSURANCE `hl7:"false,INSURANCE"`
	ACC       *ACC                `hl7:"false,ACC"`
	UB1       *UB1                `hl7:"false,UB1"`
	UB2       *UB2                `hl7:"false,UB2"`
	Other     []interface{}
}

func (s *ADT_A14) MessageTypeName() string {
	return "ADT_A14"
}

// ADT_A14_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A14_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A14_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A14_INSURANCE"
}

// ADT_A15 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A15 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	DG1   []DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A15) MessageTypeName() string {
	return "ADT_A15"
}

// ADT_A16 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A16 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	DG1   []DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A16) MessageTypeName() string {
	return "ADT_A16"
}

// ADT_A17 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A17 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID1  *PID  `hl7:"true,PID1"`
	PV11  *PV1  `hl7:"true,PV11"`
	PV21  *PV2  `hl7:"false,PV21"`
	OBX1  []OBX `hl7:"false,OBX1"`
	PID2  *PID  `hl7:"true,PID2"`
	PV12  *PV1  `hl7:"true,PV12"`
	PV22  *PV2  `hl7:"false,PV22"`
	OBX2  []OBX `hl7:"false,OBX2"`
	Other []interface{}
}

func (s *ADT_A17) MessageTypeName() string {
	return "ADT_A17"
}

// ADT_A17_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A17_PATIENT struct {
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"true,PV1"`
	Other []interface{}
}

func (s *ADT_A17_PATIENT) MessageTypeSubStructName() string {
	return "ADT_A17_PATIENT"
}

// ADT_A18 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A18 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	MRG   *MRG `hl7:"false,MRG"`
	PV1   *PV1 `hl7:"true,PV1"`
	Other []interface{}
}

func (s *ADT_A18) MessageTypeName() string {
	return "ADT_A18"
}

// ADT_A20 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A20 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	NPU   *NPU `hl7:"true,NPU"`
	Other []interface{}
}

func (s *ADT_A20) MessageTypeName() string {
	return "ADT_A20"
}

// ADT_A21 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A21 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A21) MessageTypeName() string {
	return "ADT_A21"
}

// ADT_A22 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A22 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A22) MessageTypeName() string {
	return "ADT_A22"
}

// ADT_A23 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A23 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A23) MessageTypeName() string {
	return "ADT_A23"
}

// ADT_A24 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A24 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID1  *PID `hl7:"true,PID1"`
	PV11  *PV1 `hl7:"false,PV11"`
	PID2  *PID `hl7:"true,PID2"`
	PV12  *PV1 `hl7:"false,PV12"`
	Other []interface{}
}

func (s *ADT_A24) MessageTypeName() string {
	return "ADT_A24"
}

// ADT_A25 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A25 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A25) MessageTypeName() string {
	return "ADT_A25"
}

// ADT_A26 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A26 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A26) MessageTypeName() string {
	return "ADT_A26"
}

// ADT_A27 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A27 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A27) MessageTypeName() string {
	return "ADT_A27"
}

// ADT_A28 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A28 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	EVN       *EVN                `hl7:"true,EVN"`
	PID       *PID                `hl7:"true,PID"`
	NK1       []NK1               `hl7:"false,NK1"`
	PV1       *PV1                `hl7:"true,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	PR1       []PR1               `hl7:"false,PR1"`
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []ADT_A28_INSURANCE `hl7:"false,INSURANCE"`
	ACC       *ACC                `hl7:"false,ACC"`
	UB1       *UB1                `hl7:"false,UB1"`
	UB2       *UB2                `hl7:"false,UB2"`
	Other     []interface{}
}

func (s *ADT_A28) MessageTypeName() string {
	return "ADT_A28"
}

// ADT_A28_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A28_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A28_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A28_INSURANCE"
}

// ADT_A29 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A29 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A29) MessageTypeName() string {
	return "ADT_A29"
}

// ADT_A30 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A30 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	MRG   *MRG `hl7:"true,MRG"`
	Other []interface{}
}

func (s *ADT_A30) MessageTypeName() string {
	return "ADT_A30"
}

// ADT_A31 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A31 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	EVN       *EVN                `hl7:"true,EVN"`
	PID       *PID                `hl7:"true,PID"`
	NK1       []NK1               `hl7:"false,NK1"`
	PV1       *PV1                `hl7:"true,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	PR1       []PR1               `hl7:"false,PR1"`
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []ADT_A31_INSURANCE `hl7:"false,INSURANCE"`
	ACC       *ACC                `hl7:"false,ACC"`
	UB1       *UB1                `hl7:"false,UB1"`
	UB2       *UB2                `hl7:"false,UB2"`
	Other     []interface{}
}

func (s *ADT_A31) MessageTypeName() string {
	return "ADT_A31"
}

// ADT_A31_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A31_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ADT_A31_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A31_INSURANCE"
}

// ADT_A32 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A32 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A32) MessageTypeName() string {
	return "ADT_A32"
}

// ADT_A33 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A33 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ADT_A33) MessageTypeName() string {
	return "ADT_A33"
}

// ADT_A34 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A34 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	MRG   *MRG `hl7:"true,MRG"`
	Other []interface{}
}

func (s *ADT_A34) MessageTypeName() string {
	return "ADT_A34"
}

// ADT_A35 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A35 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	MRG   *MRG `hl7:"true,MRG"`
	Other []interface{}
}

func (s *ADT_A35) MessageTypeName() string {
	return "ADT_A35"
}

// ADT_A36 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A36 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	MRG   *MRG `hl7:"true,MRG"`
	Other []interface{}
}

func (s *ADT_A36) MessageTypeName() string {
	return "ADT_A36"
}

// ADT_A37 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ADT_A37 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID1  *PID `hl7:"true,PID1"`
	PV11  *PV1 `hl7:"false,PV11"`
	PID2  *PID `hl7:"true,PID2"`
	PV12  *PV1 `hl7:"false,PV12"`
	Other []interface{}
}

func (s *ADT_A37) MessageTypeName() string {
	return "ADT_A37"
}

// ADT_A40 represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ADT_A40 struct {
	MSH     *MSH              `hl7:"true,MSH"`
	EVN     *EVN              `hl7:"true,EVN"`
	PATIENT []ADT_A40_PATIENT `hl7:"true,PATIENT"`
	Other   []interface{}
}

func (s *ADT_A40) MessageTypeName() string {
	return "ADT_A40"
}

// ADT_A40_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ADT_A40_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	// Missing: PD1
	MRG   *MRG `hl7:"true,MRG"`
	PV1   *PV1 `hl7:"false,PV1"`
	Other []interface{}
}

func (s *ADT_A40_PATIENT) MessageTypeSubStructName() string {
	return "ADT_A40_PATIENT"
}

// ADT_A44 represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ADT_A44 struct {
	MSH     *MSH              `hl7:"true,MSH"`
	EVN     *EVN              `hl7:"true,EVN"`
	PATIENT []ADT_A44_PATIENT `hl7:"true,PATIENT"`
	Other   []interface{}
}

func (s *ADT_A44) MessageTypeName() string {
	return "ADT_A44"
}

// ADT_A44_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ADT_A44_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	// Missing: PD1
	MRG   *MRG `hl7:"true,MRG"`
	Other []interface{}
}

func (s *ADT_A44_PATIENT) MessageTypeSubStructName() string {
	return "ADT_A44_PATIENT"
}

// BAR_P01 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type BAR_P01 struct {
	MSH   *MSH            `hl7:"true,MSH"`
	EVN   *EVN            `hl7:"true,EVN"`
	PID   *PID            `hl7:"true,PID"`
	VISIT []BAR_P01_VISIT `hl7:"true,VISIT"`
	Other []interface{}
}

func (s *BAR_P01) MessageTypeName() string {
	return "BAR_P01"
}

// BAR_P01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type BAR_P01_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *BAR_P01_INSURANCE) MessageTypeSubStructName() string {
	return "BAR_P01_INSURANCE"
}

// BAR_P01_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.2
type BAR_P01_VISIT struct {
	PV1       *PV1                `hl7:"false,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	PR1       []PR1               `hl7:"false,PR1"`
	GT1       []GT1               `hl7:"false,GT1"`
	NK1       []NK1               `hl7:"false,NK1"`
	INSURANCE []BAR_P01_INSURANCE `hl7:"false,INSURANCE"`
	ACC       *ACC                `hl7:"false,ACC"`
	UB1       *UB1                `hl7:"false,UB1"`
	UB2       *UB2                `hl7:"false,UB2"`
	Other     []interface{}
}

func (s *BAR_P01_VISIT) MessageTypeSubStructName() string {
	return "BAR_P01_VISIT"
}

// BAR_P02 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type BAR_P02 struct {
	MSH     *MSH              `hl7:"true,MSH"`
	EVN     *EVN              `hl7:"true,EVN"`
	PATIENT []BAR_P02_PATIENT `hl7:"true,PATIENT"`
	Other   []interface{}
}

func (s *BAR_P02) MessageTypeName() string {
	return "BAR_P02"
}

// BAR_P02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.2
type BAR_P02_PATIENT struct {
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"false,PV1"`
	Other []interface{}
}

func (s *BAR_P02_PATIENT) MessageTypeSubStructName() string {
	return "BAR_P02_PATIENT"
}

// DFT_P03 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type DFT_P03 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"false,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	FT1   []FT1 `hl7:"true,FT1"`
	Other []interface{}
}

func (s *DFT_P03) MessageTypeName() string {
	return "DFT_P03"
}

// DSR_P04 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type DSR_P04 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MSA   *MSA  `hl7:"true,MSA"`
	ERR   *ERR  `hl7:"false,ERR"`
	QRD   *QRD  `hl7:"true,QRD"`
	QRF   *QRF  `hl7:"false,QRF"`
	DSP   []DSP `hl7:"true,DSP"`
	DSC   *DSC  `hl7:"false,DSC"`
	Other []interface{}
}

func (s *DSR_P04) MessageTypeName() string {
	return "DSR_P04"
}

// DSR_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type DSR_Q01 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MSA   *MSA  `hl7:"true,MSA"`
	ERR   *ERR  `hl7:"false,ERR"`
	QRD   *QRD  `hl7:"true,QRD"`
	QRF   *QRF  `hl7:"false,QRF"`
	DSP   []DSP `hl7:"true,DSP"`
	DSC   *DSC  `hl7:"false,DSC"`
	Other []interface{}
}

func (s *DSR_Q01) MessageTypeName() string {
	return "DSR_Q01"
}

// DSR_Q03 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type DSR_Q03 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MSA   *MSA  `hl7:"false,MSA"`
	QRD   *QRD  `hl7:"true,QRD"`
	QRF   *QRF  `hl7:"false,QRF"`
	DSP   []DSP `hl7:"true,DSP"`
	DSC   *DSC  `hl7:"false,DSC"`
	Other []interface{}
}

func (s *DSR_Q03) MessageTypeName() string {
	return "DSR_Q03"
}

// DSR_R03 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type DSR_R03 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MSA   *MSA  `hl7:"false,MSA"`
	QRD   *QRD  `hl7:"true,QRD"`
	QRF   *QRF  `hl7:"false,QRF"`
	DSP   []DSP `hl7:"true,DSP"`
	DSC   *DSC  `hl7:"false,DSC"`
	Other []interface{}
}

func (s *DSR_R03) MessageTypeName() string {
	return "DSR_R03"
}

// MCF_Q02 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type MCF_Q02 struct {
	MSH   *MSH `hl7:"true,MSH"`
	MSA   *MSA `hl7:"true,MSA"`
	Other []interface{}
}

func (s *MCF_Q02) MessageTypeName() string {
	return "MCF_Q02"
}

// MFD_M01 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFD_M01 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MFI   *MFI  `hl7:"true,MFI"`
	MFA   []MFA `hl7:"false,MFA"`
	Other []interface{}
}

func (s *MFD_M01) MessageTypeName() string {
	return "MFD_M01"
}

// MFD_M02 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFD_M02 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MFI   *MFI  `hl7:"true,MFI"`
	MFA   []MFA `hl7:"false,MFA"`
	Other []interface{}
}

func (s *MFD_M02) MessageTypeName() string {
	return "MFD_M02"
}

// MFD_M03 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFD_M03 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MFI   *MFI  `hl7:"true,MFI"`
	MFA   []MFA `hl7:"false,MFA"`
	Other []interface{}
}

func (s *MFD_M03) MessageTypeName() string {
	return "MFD_M03"
}

// MFK_M01 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFK_M01 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MSA   *MSA  `hl7:"true,MSA"`
	ERR   *ERR  `hl7:"false,ERR"`
	MFI   *MFI  `hl7:"true,MFI"`
	MFA   []MFA `hl7:"false,MFA"`
	Other []interface{}
}

func (s *MFK_M01) MessageTypeName() string {
	return "MFK_M01"
}

// MFK_M02 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFK_M02 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MSA   *MSA  `hl7:"true,MSA"`
	ERR   *ERR  `hl7:"false,ERR"`
	MFI   *MFI  `hl7:"true,MFI"`
	MFA   []MFA `hl7:"false,MFA"`
	Other []interface{}
}

func (s *MFK_M02) MessageTypeName() string {
	return "MFK_M02"
}

// MFK_M03 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFK_M03 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MSA   *MSA  `hl7:"true,MSA"`
	ERR   *ERR  `hl7:"false,ERR"`
	MFI   *MFI  `hl7:"true,MFI"`
	MFA   []MFA `hl7:"false,MFA"`
	Other []interface{}
}

func (s *MFK_M03) MessageTypeName() string {
	return "MFK_M03"
}

// MFN_M01 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFN_M01 struct {
	MSH   *MSH         `hl7:"true,MSH"`
	MFI   *MFI         `hl7:"true,MFI"`
	MF    []MFN_M01_MF `hl7:"true,MF"`
	Other []interface{}
}

func (s *MFN_M01) MessageTypeName() string {
	return "MFN_M01"
}

// MFN_M01_MF represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFN_M01_MF struct {
	MFE *MFE `hl7:"true,MFE"`
	// Missing: anyZSegment
	Other []interface{}
}

func (s *MFN_M01_MF) MessageTypeSubStructName() string {
	return "MFN_M01_MF"
}

// MFN_M02 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFN_M02 struct {
	MSH      *MSH               `hl7:"true,MSH"`
	MFI      *MFI               `hl7:"true,MFI"`
	MF_STAFF []MFN_M02_MF_STAFF `hl7:"true,MF_STAFF"`
	Other    []interface{}
}

func (s *MFN_M02) MessageTypeName() string {
	return "MFN_M02"
}

// MFN_M02_MF_STAFF represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFN_M02_MF_STAFF struct {
	MFE *MFE `hl7:"true,MFE"`
	// Missing: anyZSegment
	Other []interface{}
}

func (s *MFN_M02_MF_STAFF) MessageTypeSubStructName() string {
	return "MFN_M02_MF_STAFF"
}

// MFN_M03 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFN_M03 struct {
	MSH     *MSH              `hl7:"true,MSH"`
	MFI     *MFI              `hl7:"true,MFI"`
	MF_TEST []MFN_M03_MF_TEST `hl7:"true,MF_TEST"`
	Other   []interface{}
}

func (s *MFN_M03) MessageTypeName() string {
	return "MFN_M03"
}

// MFN_M03_MF_TEST represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFN_M03_MF_TEST struct {
	MFE *MFE `hl7:"true,MFE"`
	// Missing: anyZSegment
	Other []interface{}
}

func (s *MFN_M03_MF_TEST) MessageTypeSubStructName() string {
	return "MFN_M03_MF_TEST"
}

// MFQ_M01 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFQ_M01 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QRD   *QRD `hl7:"true,QRD"`
	QRF   *QRF `hl7:"false,QRF"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *MFQ_M01) MessageTypeName() string {
	return "MFQ_M01"
}

// MFQ_M02 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFQ_M02 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QRD   *QRD `hl7:"true,QRD"`
	QRF   *QRF `hl7:"false,QRF"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *MFQ_M02) MessageTypeName() string {
	return "MFQ_M02"
}

// MFQ_M03 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFQ_M03 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QRD   *QRD `hl7:"true,QRD"`
	QRF   *QRF `hl7:"false,QRF"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *MFQ_M03) MessageTypeName() string {
	return "MFQ_M03"
}

// MFR_M01 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFR_M01 struct {
	MSH   *MSH         `hl7:"true,MSH"`
	MSA   *MSA         `hl7:"true,MSA"`
	ERR   *ERR         `hl7:"false,ERR"`
	QRD   *QRD         `hl7:"true,QRD"`
	QRF   *QRF         `hl7:"false,QRF"`
	MFI   *MFI         `hl7:"true,MFI"`
	MF    []MFR_M01_MF `hl7:"true,MF"`
	DSC   *DSC         `hl7:"false,DSC"`
	Other []interface{}
}

func (s *MFR_M01) MessageTypeName() string {
	return "MFR_M01"
}

// MFR_M01_MF represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFR_M01_MF struct {
	MFE *MFE `hl7:"true,MFE"`
	// Missing: anyZSegment
	Other []interface{}
}

func (s *MFR_M01_MF) MessageTypeSubStructName() string {
	return "MFR_M01_MF"
}

// MFR_M02 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFR_M02 struct {
	MSH      *MSH               `hl7:"true,MSH"`
	MSA      *MSA               `hl7:"true,MSA"`
	ERR      *ERR               `hl7:"false,ERR"`
	QRD      *QRD               `hl7:"true,QRD"`
	QRF      *QRF               `hl7:"false,QRF"`
	MFI      *MFI               `hl7:"true,MFI"`
	MF_STAFF []MFR_M02_MF_STAFF `hl7:"true,MF_STAFF"`
	DSC      *DSC               `hl7:"false,DSC"`
	Other    []interface{}
}

func (s *MFR_M02) MessageTypeName() string {
	return "MFR_M02"
}

// MFR_M02_MF_STAFF represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFR_M02_MF_STAFF struct {
	MFE *MFE `hl7:"true,MFE"`
	// Missing: anyZSegment
	Other []interface{}
}

func (s *MFR_M02_MF_STAFF) MessageTypeSubStructName() string {
	return "MFR_M02_MF_STAFF"
}

// MFR_M03 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFR_M03 struct {
	MSH     *MSH              `hl7:"true,MSH"`
	MSA     *MSA              `hl7:"true,MSA"`
	ERR     *ERR              `hl7:"false,ERR"`
	QRD     *QRD              `hl7:"true,QRD"`
	QRF     *QRF              `hl7:"false,QRF"`
	MFI     *MFI              `hl7:"true,MFI"`
	MF_TEST []MFR_M03_MF_TEST `hl7:"true,MF_TEST"`
	DSC     *DSC              `hl7:"false,DSC"`
	Other   []interface{}
}

func (s *MFR_M03) MessageTypeName() string {
	return "MFR_M03"
}

// MFR_M03_MF_TEST represents the corresponding HL7 message type.
// Definition from HL7 2.2
type MFR_M03_MF_TEST struct {
	MFE *MFE `hl7:"true,MFE"`
	// Missing: anyZSegment
	Other []interface{}
}

func (s *MFR_M03_MF_TEST) MessageTypeSubStructName() string {
	return "MFR_M03_MF_TEST"
}

// NMD_N01_APP_STATS represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMD_N01_APP_STATS struct {
	NST   *NST  `hl7:"true,NST"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *NMD_N01_APP_STATS) MessageTypeSubStructName() string {
	return "NMD_N01_APP_STATS"
}

// NMD_N01_APP_STATUS represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMD_N01_APP_STATUS struct {
	NSC   *NSC  `hl7:"true,NSC"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *NMD_N01_APP_STATUS) MessageTypeSubStructName() string {
	return "NMD_N01_APP_STATUS"
}

// NMD_N01_CLOCK represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMD_N01_CLOCK struct {
	NCK   *NCK  `hl7:"true,NCK"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *NMD_N01_CLOCK) MessageTypeSubStructName() string {
	return "NMD_N01_CLOCK"
}

// NMD_N01_CLOCK_AND_STATS_WITH_NOTES represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMD_N01_CLOCK_AND_STATS_WITH_NOTES struct {
	CLOCK      *NMD_N01_CLOCK      `hl7:"false,CLOCK"`
	APP_STATS  *NMD_N01_APP_STATS  `hl7:"false,APP_STATS"`
	APP_STATUS *NMD_N01_APP_STATUS `hl7:"false,APP_STATUS"`
	Other      []interface{}
}

func (s *NMD_N01_CLOCK_AND_STATS_WITH_NOTES) MessageTypeSubStructName() string {
	return "NMD_N01_CLOCK_AND_STATS_WITH_NOTES"
}

// NMD_N01 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMD_N01 struct {
	MSH                        *MSH                                 `hl7:"true,MSH"`
	CLOCK_AND_STATS_WITH_NOTES []NMD_N01_CLOCK_AND_STATS_WITH_NOTES `hl7:"true,CLOCK_AND_STATS_WITH_NOTES"`
	Other                      []interface{}
}

func (s *NMD_N01) MessageTypeName() string {
	return "NMD_N01"
}

// NMQ_N02_CLOCK_AND_STATISTICS represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMQ_N02_CLOCK_AND_STATISTICS struct {
	NCK   *NCK `hl7:"false,NCK"`
	NST   *NST `hl7:"false,NST"`
	NSC   *NSC `hl7:"false,NSC"`
	Other []interface{}
}

func (s *NMQ_N02_CLOCK_AND_STATISTICS) MessageTypeSubStructName() string {
	return "NMQ_N02_CLOCK_AND_STATISTICS"
}

// NMQ_N02 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMQ_N02 struct {
	MSH                  *MSH                           `hl7:"true,MSH"`
	QRY_WITH_DETAIL      *NMQ_N02_QRY_WITH_DETAIL       `hl7:"false,QRY_WITH_DETAIL"`
	CLOCK_AND_STATISTICS []NMQ_N02_CLOCK_AND_STATISTICS `hl7:"true,CLOCK_AND_STATISTICS"`
	Other                []interface{}
}

func (s *NMQ_N02) MessageTypeName() string {
	return "NMQ_N02"
}

// NMQ_N02_QRY_WITH_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMQ_N02_QRY_WITH_DETAIL struct {
	QRD   *QRD `hl7:"true,QRD"`
	QRF   *QRF `hl7:"false,QRF"`
	Other []interface{}
}

func (s *NMQ_N02_QRY_WITH_DETAIL) MessageTypeSubStructName() string {
	return "NMQ_N02_QRY_WITH_DETAIL"
}

// NMR_N02_CLOCK_AND_STATS_WITH_NOTES_ALT represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMR_N02_CLOCK_AND_STATS_WITH_NOTES_ALT struct {
	NCK   *NCK  `hl7:"false,NCK"`
	NTE1  []NTE `hl7:"false,NTE1"`
	NST   *NST  `hl7:"false,NST"`
	NTE2  []NTE `hl7:"false,NTE2"`
	NSC   *NSC  `hl7:"false,NSC"`
	NTE3  []NTE `hl7:"false,NTE3"`
	Other []interface{}
}

func (s *NMR_N02_CLOCK_AND_STATS_WITH_NOTES_ALT) MessageTypeSubStructName() string {
	return "NMR_N02_CLOCK_AND_STATS_WITH_NOTES_ALT"
}

// NMR_N02 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type NMR_N02 struct {
	MSH                            *MSH                                     `hl7:"true,MSH"`
	MSA                            *MSA                                     `hl7:"true,MSA"`
	ERR                            *ERR                                     `hl7:"false,ERR"`
	QRD                            *QRD                                     `hl7:"false,QRD"`
	CLOCK_AND_STATS_WITH_NOTES_ALT []NMR_N02_CLOCK_AND_STATS_WITH_NOTES_ALT `hl7:"true,CLOCK_AND_STATS_WITH_NOTES_ALT"`
	Other                          []interface{}
}

func (s *NMR_N02) MessageTypeName() string {
	return "NMR_N02"
}

// ORF_R04 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORF_R04 struct {
	MSH            *MSH                     `hl7:"true,MSH"`
	MSA            *MSA                     `hl7:"true,MSA"`
	QUERY_RESPONSE []ORF_R04_QUERY_RESPONSE `hl7:"true,QUERY_RESPONSE"`
	ORDER          []ORF_R04_ORDER          `hl7:"true,ORDER"`
	DSC            *DSC                     `hl7:"false,DSC"`
	Other          []interface{}
}

func (s *ORF_R04) MessageTypeName() string {
	return "ORF_R04"
}

// ORF_R04_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORF_R04_OBSERVATION struct {
	OBX   *OBX  `hl7:"false,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORF_R04_OBSERVATION) MessageTypeSubStructName() string {
	return "ORF_R04_OBSERVATION"
}

// ORF_R04_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORF_R04_ORDER struct {
	ORC         *ORC                  `hl7:"false,ORC"`
	OBR         *OBR                  `hl7:"true,OBR"`
	NTE         []NTE                 `hl7:"false,NTE"`
	OBSERVATION []ORF_R04_OBSERVATION `hl7:"true,OBSERVATION"`
	Other       []interface{}
}

func (s *ORF_R04_ORDER) MessageTypeSubStructName() string {
	return "ORF_R04_ORDER"
}

// ORF_R04_QUERY_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORF_R04_QUERY_RESPONSE struct {
	QRD   *QRD  `hl7:"true,QRD"`
	QRF   *QRF  `hl7:"false,QRF"`
	PID   *PID  `hl7:"false,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORF_R04_QUERY_RESPONSE) MessageTypeSubStructName() string {
	return "ORF_R04_QUERY_RESPONSE"
}

// ORM_O01_CHOICE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORM_O01_CHOICE struct {
	// Only one of the following fields will be set.
	OBR   *OBR `hl7:"true,OBR"`
	RQD   *RQD `hl7:"true,RQD"`
	RQ1   *RQ1 `hl7:"true,RQ1"`
	RXO   *RXO `hl7:"true,RXO"`
	ODS   *ODS `hl7:"true,ODS"`
	ODT   *ODT `hl7:"true,ODT"`
	Other []interface{}
}

func (s *ORM_O01_CHOICE) MessageTypeSubStructName() string {
	return "ORM_O01_CHOICE"
}

// ORM_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORM_O01 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	NTE     []NTE            `hl7:"false,NTE"`
	PATIENT *ORM_O01_PATIENT `hl7:"false,PATIENT"`
	ORDER   []ORM_O01_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *ORM_O01) MessageTypeName() string {
	return "ORM_O01"
}

// ORM_O01_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORM_O01_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *ORM_O01_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	BLG          *BLG                  `hl7:"false,BLG"`
	Other        []interface{}
}

func (s *ORM_O01_ORDER) MessageTypeSubStructName() string {
	return "ORM_O01_ORDER"
}

// ORM_O01_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORM_O01_ORDER_DETAIL struct {
	CHOICE *ORM_O01_CHOICE `hl7:"true,CHOICE"`
	NTE1   []NTE           `hl7:"false,NTE1"`
	OBX    []OBX           `hl7:"false,OBX"`
	NTE2   []NTE           `hl7:"true,NTE2"`
	Other  []interface{}
}

func (s *ORM_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "ORM_O01_ORDER_DETAIL"
}

// ORM_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORM_O01_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	PV1   *PV1  `hl7:"false,PV1"`
	Other []interface{}
}

func (s *ORM_O01_PATIENT) MessageTypeSubStructName() string {
	return "ORM_O01_PATIENT"
}

// ORR_O02_CHOICE represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORR_O02_CHOICE struct {
	// Only one of the following fields will be set.
	OBR   *OBR `hl7:"true,OBR"`
	RQD   *RQD `hl7:"true,RQD"`
	RQ1   *RQ1 `hl7:"true,RQ1"`
	RXO   *RXO `hl7:"true,RXO"`
	ODS   *ODS `hl7:"true,ODS"`
	ODT   *ODT `hl7:"true,ODT"`
	Other []interface{}
}

func (s *ORR_O02_CHOICE) MessageTypeSubStructName() string {
	return "ORR_O02_CHOICE"
}

// ORR_O02 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORR_O02 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	MSA     *MSA             `hl7:"true,MSA"`
	NTE     []NTE            `hl7:"false,NTE"`
	PATIENT *ORR_O02_PATIENT `hl7:"false,PATIENT"`
	Other   []interface{}
}

func (s *ORR_O02) MessageTypeName() string {
	return "ORR_O02"
}

// ORR_O02_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORR_O02_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *ORR_O02_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	NTE          []NTE                 `hl7:"false,NTE"`
	Other        []interface{}
}

func (s *ORR_O02_ORDER) MessageTypeSubStructName() string {
	return "ORR_O02_ORDER"
}

// ORR_O02_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORR_O02_ORDER_DETAIL struct {
	CHOICE *ORR_O02_CHOICE `hl7:"true,CHOICE"`
	Other  []interface{}
}

func (s *ORR_O02_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "ORR_O02_ORDER_DETAIL"
}

// ORR_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORR_O02_PATIENT struct {
	PID   *PID            `hl7:"false,PID"`
	NTE   []NTE           `hl7:"false,NTE"`
	ORDER []ORR_O02_ORDER `hl7:"true,ORDER"`
	Other []interface{}
}

func (s *ORR_O02_PATIENT) MessageTypeSubStructName() string {
	return "ORR_O02_PATIENT"
}

// ORU_R01 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORU_R01 struct {
	MSH            *MSH                     `hl7:"true,MSH"`
	PATIENT_RESULT []ORU_R01_PATIENT_RESULT `hl7:"true,PATIENT_RESULT"`
	DSC            *DSC                     `hl7:"false,DSC"`
	Other          []interface{}
}

func (s *ORU_R01) MessageTypeName() string {
	return "ORU_R01"
}

// ORU_R01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORU_R01_OBSERVATION struct {
	OBX   *OBX  `hl7:"false,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORU_R01_OBSERVATION) MessageTypeSubStructName() string {
	return "ORU_R01_OBSERVATION"
}

// ORU_R01_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORU_R01_ORDER_OBSERVATION struct {
	ORC         *ORC                  `hl7:"false,ORC"`
	OBR         *OBR                  `hl7:"true,OBR"`
	NTE         []NTE                 `hl7:"false,NTE"`
	OBSERVATION []ORU_R01_OBSERVATION `hl7:"true,OBSERVATION"`
	Other       []interface{}
}

func (s *ORU_R01_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "ORU_R01_ORDER_OBSERVATION"
}

// ORU_R01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORU_R01_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	PV1   *PV1  `hl7:"false,PV1"`
	Other []interface{}
}

func (s *ORU_R01_PATIENT) MessageTypeSubStructName() string {
	return "ORU_R01_PATIENT"
}

// ORU_R01_PATIENT_RESULT represents the corresponding HL7 message type.
// Definition from HL7 2.2
type ORU_R01_PATIENT_RESULT struct {
	PATIENT           *ORU_R01_PATIENT            `hl7:"false,PATIENT"`
	ORDER_OBSERVATION []ORU_R01_ORDER_OBSERVATION `hl7:"true,ORDER_OBSERVATION"`
	Other             []interface{}
}

func (s *ORU_R01_PATIENT_RESULT) MessageTypeSubStructName() string {
	return "ORU_R01_PATIENT_RESULT"
}

// ORU_R03 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ORU_R03 struct {
	MSH            *MSH                     `hl7:"true,MSH"`
	PATIENT_RESULT []ORU_R03_PATIENT_RESULT `hl7:"true,PATIENT_RESULT"`
	DSC            *DSC                     `hl7:"false,DSC"`
	Other          []interface{}
}

func (s *ORU_R03) MessageTypeName() string {
	return "ORU_R03"
}

// ORU_R03_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ORU_R03_OBSERVATION struct {
	OBX   *OBX  `hl7:"false,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORU_R03_OBSERVATION) MessageTypeSubStructName() string {
	return "ORU_R03_OBSERVATION"
}

// ORU_R03_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ORU_R03_ORDER_OBSERVATION struct {
	ORC         *ORC                  `hl7:"false,ORC"`
	OBR         *OBR                  `hl7:"true,OBR"`
	NTE         []NTE                 `hl7:"false,NTE"`
	OBSERVATION []ORU_R03_OBSERVATION `hl7:"true,OBSERVATION"`
	Other       []interface{}
}

func (s *ORU_R03_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "ORU_R03_ORDER_OBSERVATION"
}

// ORU_R03_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ORU_R03_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	PV1   *PV1  `hl7:"false,PV1"`
	Other []interface{}
}

func (s *ORU_R03_PATIENT) MessageTypeSubStructName() string {
	return "ORU_R03_PATIENT"
}

// ORU_R03_PATIENT_RESULT represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ORU_R03_PATIENT_RESULT struct {
	PATIENT           *ORU_R03_PATIENT            `hl7:"false,PATIENT"`
	ORDER_OBSERVATION []ORU_R03_ORDER_OBSERVATION `hl7:"true,ORDER_OBSERVATION"`
	Other             []interface{}
}

func (s *ORU_R03_PATIENT_RESULT) MessageTypeSubStructName() string {
	return "ORU_R03_PATIENT_RESULT"
}

// ORU_R32 represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ORU_R32 struct {
	MSH            *MSH                     `hl7:"true,MSH"`
	PATIENT_RESULT []ORU_R32_PATIENT_RESULT `hl7:"true,PATIENT_RESULT"`
	DSC            *DSC                     `hl7:"false,DSC"`
	Other          []interface{}
}

func (s *ORU_R32) MessageTypeName() string {
	return "ORU_R32"
}

// ORU_R32_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ORU_R32_OBSERVATION struct {
	OBX   *OBX  `hl7:"false,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORU_R32_OBSERVATION) MessageTypeSubStructName() string {
	return "ORU_R32_OBSERVATION"
}

// ORU_R32_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ORU_R32_ORDER_OBSERVATION struct {
	ORC         *ORC                  `hl7:"false,ORC"`
	OBR         *OBR                  `hl7:"true,OBR"`
	NTE         []NTE                 `hl7:"false,NTE"`
	OBSERVATION []ORU_R32_OBSERVATION `hl7:"true,OBSERVATION"`
	// Missing: CTI
	Other []interface{}
}

func (s *ORU_R32_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "ORU_R32_ORDER_OBSERVATION"
}

// ORU_R32_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ORU_R32_PATIENT struct {
	PID *PID `hl7:"true,PID"`
	// Missing: PD1
	NK1   []NK1          `hl7:"false,NK1"`
	NTE   []NTE          `hl7:"false,NTE"`
	VISIT *ORU_R32_VISIT `hl7:"false,VISIT"`
	Other []interface{}
}

func (s *ORU_R32_PATIENT) MessageTypeSubStructName() string {
	return "ORU_R32_PATIENT"
}

// ORU_R32_PATIENT_RESULT represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ORU_R32_PATIENT_RESULT struct {
	PATIENT           *ORU_R32_PATIENT            `hl7:"false,PATIENT"`
	ORDER_OBSERVATION []ORU_R32_ORDER_OBSERVATION `hl7:"true,ORDER_OBSERVATION"`
	Other             []interface{}
}

func (s *ORU_R32_PATIENT_RESULT) MessageTypeSubStructName() string {
	return "ORU_R32_PATIENT_RESULT"
}

// ORU_R32_VISIT represents the corresponding HL7 message type.
// Definition from HL7 SYNTHETIC
type ORU_R32_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *ORU_R32_VISIT) MessageTypeSubStructName() string {
	return "ORU_R32_VISIT"
}

// QRY_A19 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type QRY_A19 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QRD   *QRD `hl7:"true,QRD"`
	QRF   *QRF `hl7:"false,QRF"`
	Other []interface{}
}

func (s *QRY_A19) MessageTypeName() string {
	return "QRY_A19"
}

// QRY_P04 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type QRY_P04 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QRD   *QRD `hl7:"true,QRD"`
	QRF   *QRF `hl7:"false,QRF"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *QRY_P04) MessageTypeName() string {
	return "QRY_P04"
}

// QRY_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type QRY_Q01 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QRD   *QRD `hl7:"true,QRD"`
	QRF   *QRF `hl7:"false,QRF"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *QRY_Q01) MessageTypeName() string {
	return "QRY_Q01"
}

// QRY_Q02 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type QRY_Q02 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QRD   *QRD `hl7:"true,QRD"`
	QRF   *QRF `hl7:"false,QRF"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *QRY_Q02) MessageTypeName() string {
	return "QRY_Q02"
}

// QRY_R02 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type QRY_R02 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QRD   *QRD `hl7:"true,QRD"`
	QRF   *QRF `hl7:"true,QRF"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *QRY_R02) MessageTypeName() string {
	return "QRY_R02"
}

// UDM_Q05 represents the corresponding HL7 message type.
// Definition from HL7 2.2
type UDM_Q05 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	URD   *URD  `hl7:"true,URD"`
	URS   *URS  `hl7:"false,URS"`
	DSP   []DSP `hl7:"true,DSP"`
	DSC   *DSC  `hl7:"true,DSC"`
	Other []interface{}
}

func (s *UDM_Q05) MessageTypeName() string {
	return "UDM_Q05"
}

// GenericHL7Segment represents the corresponding HL7 segment type.
type GenericHL7Segment struct {
	segment []byte
}

func (s *GenericHL7Segment) SegmentName() string {
	return "GenericHL7Segment"
}

// ACC returns the first ACC segment within the message, or nil if there isn't one.
func (m *Message) ACC() (*ACC, error) {
	ps, err := m.Parse("ACC")
	pst, ok := ps.(*ACC)
	if ok {
		return pst, err
	}
	return nil, err
}

// ADD returns the first ADD segment within the message, or nil if there isn't one.
func (m *Message) ADD() (*ADD, error) {
	ps, err := m.Parse("ADD")
	pst, ok := ps.(*ADD)
	if ok {
		return pst, err
	}
	return nil, err
}

// AL1 returns the first AL1 segment within the message, or nil if there isn't one.
func (m *Message) AL1() (*AL1, error) {
	ps, err := m.Parse("AL1")
	pst, ok := ps.(*AL1)
	if ok {
		return pst, err
	}
	return nil, err
}

// BHS returns the first BHS segment within the message, or nil if there isn't one.
func (m *Message) BHS() (*BHS, error) {
	ps, err := m.Parse("BHS")
	pst, ok := ps.(*BHS)
	if ok {
		return pst, err
	}
	return nil, err
}

// BLG returns the first BLG segment within the message, or nil if there isn't one.
func (m *Message) BLG() (*BLG, error) {
	ps, err := m.Parse("BLG")
	pst, ok := ps.(*BLG)
	if ok {
		return pst, err
	}
	return nil, err
}

// BTS returns the first BTS segment within the message, or nil if there isn't one.
func (m *Message) BTS() (*BTS, error) {
	ps, err := m.Parse("BTS")
	pst, ok := ps.(*BTS)
	if ok {
		return pst, err
	}
	return nil, err
}

// DG1 returns the first DG1 segment within the message, or nil if there isn't one.
func (m *Message) DG1() (*DG1, error) {
	ps, err := m.Parse("DG1")
	pst, ok := ps.(*DG1)
	if ok {
		return pst, err
	}
	return nil, err
}

// DSC returns the first DSC segment within the message, or nil if there isn't one.
func (m *Message) DSC() (*DSC, error) {
	ps, err := m.Parse("DSC")
	pst, ok := ps.(*DSC)
	if ok {
		return pst, err
	}
	return nil, err
}

// DSP returns the first DSP segment within the message, or nil if there isn't one.
func (m *Message) DSP() (*DSP, error) {
	ps, err := m.Parse("DSP")
	pst, ok := ps.(*DSP)
	if ok {
		return pst, err
	}
	return nil, err
}

// ERR returns the first ERR segment within the message, or nil if there isn't one.
func (m *Message) ERR() (*ERR, error) {
	ps, err := m.Parse("ERR")
	pst, ok := ps.(*ERR)
	if ok {
		return pst, err
	}
	return nil, err
}

// EVN returns the first EVN segment within the message, or nil if there isn't one.
func (m *Message) EVN() (*EVN, error) {
	ps, err := m.Parse("EVN")
	pst, ok := ps.(*EVN)
	if ok {
		return pst, err
	}
	return nil, err
}

// FHS returns the first FHS segment within the message, or nil if there isn't one.
func (m *Message) FHS() (*FHS, error) {
	ps, err := m.Parse("FHS")
	pst, ok := ps.(*FHS)
	if ok {
		return pst, err
	}
	return nil, err
}

// FT1 returns the first FT1 segment within the message, or nil if there isn't one.
func (m *Message) FT1() (*FT1, error) {
	ps, err := m.Parse("FT1")
	pst, ok := ps.(*FT1)
	if ok {
		return pst, err
	}
	return nil, err
}

// FTS returns the first FTS segment within the message, or nil if there isn't one.
func (m *Message) FTS() (*FTS, error) {
	ps, err := m.Parse("FTS")
	pst, ok := ps.(*FTS)
	if ok {
		return pst, err
	}
	return nil, err
}

// GT1 returns the first GT1 segment within the message, or nil if there isn't one.
func (m *Message) GT1() (*GT1, error) {
	ps, err := m.Parse("GT1")
	pst, ok := ps.(*GT1)
	if ok {
		return pst, err
	}
	return nil, err
}

// IN1 returns the first IN1 segment within the message, or nil if there isn't one.
func (m *Message) IN1() (*IN1, error) {
	ps, err := m.Parse("IN1")
	pst, ok := ps.(*IN1)
	if ok {
		return pst, err
	}
	return nil, err
}

// IN2 returns the first IN2 segment within the message, or nil if there isn't one.
func (m *Message) IN2() (*IN2, error) {
	ps, err := m.Parse("IN2")
	pst, ok := ps.(*IN2)
	if ok {
		return pst, err
	}
	return nil, err
}

// IN3 returns the first IN3 segment within the message, or nil if there isn't one.
func (m *Message) IN3() (*IN3, error) {
	ps, err := m.Parse("IN3")
	pst, ok := ps.(*IN3)
	if ok {
		return pst, err
	}
	return nil, err
}

// MFA returns the first MFA segment within the message, or nil if there isn't one.
func (m *Message) MFA() (*MFA, error) {
	ps, err := m.Parse("MFA")
	pst, ok := ps.(*MFA)
	if ok {
		return pst, err
	}
	return nil, err
}

// MFE returns the first MFE segment within the message, or nil if there isn't one.
func (m *Message) MFE() (*MFE, error) {
	ps, err := m.Parse("MFE")
	pst, ok := ps.(*MFE)
	if ok {
		return pst, err
	}
	return nil, err
}

// MFI returns the first MFI segment within the message, or nil if there isn't one.
func (m *Message) MFI() (*MFI, error) {
	ps, err := m.Parse("MFI")
	pst, ok := ps.(*MFI)
	if ok {
		return pst, err
	}
	return nil, err
}

// MRG returns the first MRG segment within the message, or nil if there isn't one.
func (m *Message) MRG() (*MRG, error) {
	ps, err := m.Parse("MRG")
	pst, ok := ps.(*MRG)
	if ok {
		return pst, err
	}
	return nil, err
}

// MSA returns the first MSA segment within the message, or nil if there isn't one.
func (m *Message) MSA() (*MSA, error) {
	ps, err := m.Parse("MSA")
	pst, ok := ps.(*MSA)
	if ok {
		return pst, err
	}
	return nil, err
}

// MSH returns the first MSH segment within the message, or nil if there isn't one.
func (m *Message) MSH() (*MSH, error) {
	ps, err := m.Parse("MSH")
	pst, ok := ps.(*MSH)
	if ok {
		return pst, err
	}
	return nil, err
}

// NCK returns the first NCK segment within the message, or nil if there isn't one.
func (m *Message) NCK() (*NCK, error) {
	ps, err := m.Parse("NCK")
	pst, ok := ps.(*NCK)
	if ok {
		return pst, err
	}
	return nil, err
}

// NK1 returns the first NK1 segment within the message, or nil if there isn't one.
func (m *Message) NK1() (*NK1, error) {
	ps, err := m.Parse("NK1")
	pst, ok := ps.(*NK1)
	if ok {
		return pst, err
	}
	return nil, err
}

// NPU returns the first NPU segment within the message, or nil if there isn't one.
func (m *Message) NPU() (*NPU, error) {
	ps, err := m.Parse("NPU")
	pst, ok := ps.(*NPU)
	if ok {
		return pst, err
	}
	return nil, err
}

// NSC returns the first NSC segment within the message, or nil if there isn't one.
func (m *Message) NSC() (*NSC, error) {
	ps, err := m.Parse("NSC")
	pst, ok := ps.(*NSC)
	if ok {
		return pst, err
	}
	return nil, err
}

// NST returns the first NST segment within the message, or nil if there isn't one.
func (m *Message) NST() (*NST, error) {
	ps, err := m.Parse("NST")
	pst, ok := ps.(*NST)
	if ok {
		return pst, err
	}
	return nil, err
}

// NTE returns the first NTE segment within the message, or nil if there isn't one.
func (m *Message) NTE() (*NTE, error) {
	ps, err := m.Parse("NTE")
	pst, ok := ps.(*NTE)
	if ok {
		return pst, err
	}
	return nil, err
}

// OBR returns the first OBR segment within the message, or nil if there isn't one.
func (m *Message) OBR() (*OBR, error) {
	ps, err := m.Parse("OBR")
	pst, ok := ps.(*OBR)
	if ok {
		return pst, err
	}
	return nil, err
}

// OBX returns the first OBX segment within the message, or nil if there isn't one.
func (m *Message) OBX() (*OBX, error) {
	ps, err := m.Parse("OBX")
	pst, ok := ps.(*OBX)
	if ok {
		return pst, err
	}
	return nil, err
}

// ODS returns the first ODS segment within the message, or nil if there isn't one.
func (m *Message) ODS() (*ODS, error) {
	ps, err := m.Parse("ODS")
	pst, ok := ps.(*ODS)
	if ok {
		return pst, err
	}
	return nil, err
}

// ODT returns the first ODT segment within the message, or nil if there isn't one.
func (m *Message) ODT() (*ODT, error) {
	ps, err := m.Parse("ODT")
	pst, ok := ps.(*ODT)
	if ok {
		return pst, err
	}
	return nil, err
}

// OM1 returns the first OM1 segment within the message, or nil if there isn't one.
func (m *Message) OM1() (*OM1, error) {
	ps, err := m.Parse("OM1")
	pst, ok := ps.(*OM1)
	if ok {
		return pst, err
	}
	return nil, err
}

// OM2 returns the first OM2 segment within the message, or nil if there isn't one.
func (m *Message) OM2() (*OM2, error) {
	ps, err := m.Parse("OM2")
	pst, ok := ps.(*OM2)
	if ok {
		return pst, err
	}
	return nil, err
}

// OM3 returns the first OM3 segment within the message, or nil if there isn't one.
func (m *Message) OM3() (*OM3, error) {
	ps, err := m.Parse("OM3")
	pst, ok := ps.(*OM3)
	if ok {
		return pst, err
	}
	return nil, err
}

// OM4 returns the first OM4 segment within the message, or nil if there isn't one.
func (m *Message) OM4() (*OM4, error) {
	ps, err := m.Parse("OM4")
	pst, ok := ps.(*OM4)
	if ok {
		return pst, err
	}
	return nil, err
}

// OM5 returns the first OM5 segment within the message, or nil if there isn't one.
func (m *Message) OM5() (*OM5, error) {
	ps, err := m.Parse("OM5")
	pst, ok := ps.(*OM5)
	if ok {
		return pst, err
	}
	return nil, err
}

// OM6 returns the first OM6 segment within the message, or nil if there isn't one.
func (m *Message) OM6() (*OM6, error) {
	ps, err := m.Parse("OM6")
	pst, ok := ps.(*OM6)
	if ok {
		return pst, err
	}
	return nil, err
}

// ORC returns the first ORC segment within the message, or nil if there isn't one.
func (m *Message) ORC() (*ORC, error) {
	ps, err := m.Parse("ORC")
	pst, ok := ps.(*ORC)
	if ok {
		return pst, err
	}
	return nil, err
}

// ORO returns the first ORO segment within the message, or nil if there isn't one.
func (m *Message) ORO() (*ORO, error) {
	ps, err := m.Parse("ORO")
	pst, ok := ps.(*ORO)
	if ok {
		return pst, err
	}
	return nil, err
}

// PID returns the first PID segment within the message, or nil if there isn't one.
func (m *Message) PID() (*PID, error) {
	ps, err := m.Parse("PID")
	pst, ok := ps.(*PID)
	if ok {
		return pst, err
	}
	return nil, err
}

// PR1 returns the first PR1 segment within the message, or nil if there isn't one.
func (m *Message) PR1() (*PR1, error) {
	ps, err := m.Parse("PR1")
	pst, ok := ps.(*PR1)
	if ok {
		return pst, err
	}
	return nil, err
}

// PRA returns the first PRA segment within the message, or nil if there isn't one.
func (m *Message) PRA() (*PRA, error) {
	ps, err := m.Parse("PRA")
	pst, ok := ps.(*PRA)
	if ok {
		return pst, err
	}
	return nil, err
}

// PV1 returns the first PV1 segment within the message, or nil if there isn't one.
func (m *Message) PV1() (*PV1, error) {
	ps, err := m.Parse("PV1")
	pst, ok := ps.(*PV1)
	if ok {
		return pst, err
	}
	return nil, err
}

// PV2 returns the first PV2 segment within the message, or nil if there isn't one.
func (m *Message) PV2() (*PV2, error) {
	ps, err := m.Parse("PV2")
	pst, ok := ps.(*PV2)
	if ok {
		return pst, err
	}
	return nil, err
}

// QRD returns the first QRD segment within the message, or nil if there isn't one.
func (m *Message) QRD() (*QRD, error) {
	ps, err := m.Parse("QRD")
	pst, ok := ps.(*QRD)
	if ok {
		return pst, err
	}
	return nil, err
}

// QRF returns the first QRF segment within the message, or nil if there isn't one.
func (m *Message) QRF() (*QRF, error) {
	ps, err := m.Parse("QRF")
	pst, ok := ps.(*QRF)
	if ok {
		return pst, err
	}
	return nil, err
}

// RQ1 returns the first RQ1 segment within the message, or nil if there isn't one.
func (m *Message) RQ1() (*RQ1, error) {
	ps, err := m.Parse("RQ1")
	pst, ok := ps.(*RQ1)
	if ok {
		return pst, err
	}
	return nil, err
}

// RQD returns the first RQD segment within the message, or nil if there isn't one.
func (m *Message) RQD() (*RQD, error) {
	ps, err := m.Parse("RQD")
	pst, ok := ps.(*RQD)
	if ok {
		return pst, err
	}
	return nil, err
}

// RX1 returns the first RX1 segment within the message, or nil if there isn't one.
func (m *Message) RX1() (*RX1, error) {
	ps, err := m.Parse("RX1")
	pst, ok := ps.(*RX1)
	if ok {
		return pst, err
	}
	return nil, err
}

// RXA returns the first RXA segment within the message, or nil if there isn't one.
func (m *Message) RXA() (*RXA, error) {
	ps, err := m.Parse("RXA")
	pst, ok := ps.(*RXA)
	if ok {
		return pst, err
	}
	return nil, err
}

// RXC returns the first RXC segment within the message, or nil if there isn't one.
func (m *Message) RXC() (*RXC, error) {
	ps, err := m.Parse("RXC")
	pst, ok := ps.(*RXC)
	if ok {
		return pst, err
	}
	return nil, err
}

// RXD returns the first RXD segment within the message, or nil if there isn't one.
func (m *Message) RXD() (*RXD, error) {
	ps, err := m.Parse("RXD")
	pst, ok := ps.(*RXD)
	if ok {
		return pst, err
	}
	return nil, err
}

// RXE returns the first RXE segment within the message, or nil if there isn't one.
func (m *Message) RXE() (*RXE, error) {
	ps, err := m.Parse("RXE")
	pst, ok := ps.(*RXE)
	if ok {
		return pst, err
	}
	return nil, err
}

// RXG returns the first RXG segment within the message, or nil if there isn't one.
func (m *Message) RXG() (*RXG, error) {
	ps, err := m.Parse("RXG")
	pst, ok := ps.(*RXG)
	if ok {
		return pst, err
	}
	return nil, err
}

// RXO returns the first RXO segment within the message, or nil if there isn't one.
func (m *Message) RXO() (*RXO, error) {
	ps, err := m.Parse("RXO")
	pst, ok := ps.(*RXO)
	if ok {
		return pst, err
	}
	return nil, err
}

// RXR returns the first RXR segment within the message, or nil if there isn't one.
func (m *Message) RXR() (*RXR, error) {
	ps, err := m.Parse("RXR")
	pst, ok := ps.(*RXR)
	if ok {
		return pst, err
	}
	return nil, err
}

// STF returns the first STF segment within the message, or nil if there isn't one.
func (m *Message) STF() (*STF, error) {
	ps, err := m.Parse("STF")
	pst, ok := ps.(*STF)
	if ok {
		return pst, err
	}
	return nil, err
}

// UB1 returns the first UB1 segment within the message, or nil if there isn't one.
func (m *Message) UB1() (*UB1, error) {
	ps, err := m.Parse("UB1")
	pst, ok := ps.(*UB1)
	if ok {
		return pst, err
	}
	return nil, err
}

// UB2 returns the first UB2 segment within the message, or nil if there isn't one.
func (m *Message) UB2() (*UB2, error) {
	ps, err := m.Parse("UB2")
	pst, ok := ps.(*UB2)
	if ok {
		return pst, err
	}
	return nil, err
}

// URD returns the first URD segment within the message, or nil if there isn't one.
func (m *Message) URD() (*URD, error) {
	ps, err := m.Parse("URD")
	pst, ok := ps.(*URD)
	if ok {
		return pst, err
	}
	return nil, err
}

// URS returns the first URS segment within the message, or nil if there isn't one.
func (m *Message) URS() (*URS, error) {
	ps, err := m.Parse("URS")
	pst, ok := ps.(*URS)
	if ok {
		return pst, err
	}
	return nil, err
}

// AllACC returns a slice containing all ACC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllACC() ([]*ACC, error) {
	pss, err := m.ParseAll("ACC")
	return pss.([]*ACC), err
}

// AllADD returns a slice containing all ADD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllADD() ([]*ADD, error) {
	pss, err := m.ParseAll("ADD")
	return pss.([]*ADD), err
}

// AllAL1 returns a slice containing all AL1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllAL1() ([]*AL1, error) {
	pss, err := m.ParseAll("AL1")
	return pss.([]*AL1), err
}

// AllBHS returns a slice containing all BHS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllBHS() ([]*BHS, error) {
	pss, err := m.ParseAll("BHS")
	return pss.([]*BHS), err
}

// AllBLG returns a slice containing all BLG segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllBLG() ([]*BLG, error) {
	pss, err := m.ParseAll("BLG")
	return pss.([]*BLG), err
}

// AllBTS returns a slice containing all BTS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllBTS() ([]*BTS, error) {
	pss, err := m.ParseAll("BTS")
	return pss.([]*BTS), err
}

// AllDG1 returns a slice containing all DG1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllDG1() ([]*DG1, error) {
	pss, err := m.ParseAll("DG1")
	return pss.([]*DG1), err
}

// AllDSC returns a slice containing all DSC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllDSC() ([]*DSC, error) {
	pss, err := m.ParseAll("DSC")
	return pss.([]*DSC), err
}

// AllDSP returns a slice containing all DSP segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllDSP() ([]*DSP, error) {
	pss, err := m.ParseAll("DSP")
	return pss.([]*DSP), err
}

// AllERR returns a slice containing all ERR segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllERR() ([]*ERR, error) {
	pss, err := m.ParseAll("ERR")
	return pss.([]*ERR), err
}

// AllEVN returns a slice containing all EVN segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllEVN() ([]*EVN, error) {
	pss, err := m.ParseAll("EVN")
	return pss.([]*EVN), err
}

// AllFHS returns a slice containing all FHS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllFHS() ([]*FHS, error) {
	pss, err := m.ParseAll("FHS")
	return pss.([]*FHS), err
}

// AllFT1 returns a slice containing all FT1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllFT1() ([]*FT1, error) {
	pss, err := m.ParseAll("FT1")
	return pss.([]*FT1), err
}

// AllFTS returns a slice containing all FTS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllFTS() ([]*FTS, error) {
	pss, err := m.ParseAll("FTS")
	return pss.([]*FTS), err
}

// AllGT1 returns a slice containing all GT1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllGT1() ([]*GT1, error) {
	pss, err := m.ParseAll("GT1")
	return pss.([]*GT1), err
}

// AllIN1 returns a slice containing all IN1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllIN1() ([]*IN1, error) {
	pss, err := m.ParseAll("IN1")
	return pss.([]*IN1), err
}

// AllIN2 returns a slice containing all IN2 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllIN2() ([]*IN2, error) {
	pss, err := m.ParseAll("IN2")
	return pss.([]*IN2), err
}

// AllIN3 returns a slice containing all IN3 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllIN3() ([]*IN3, error) {
	pss, err := m.ParseAll("IN3")
	return pss.([]*IN3), err
}

// AllMFA returns a slice containing all MFA segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllMFA() ([]*MFA, error) {
	pss, err := m.ParseAll("MFA")
	return pss.([]*MFA), err
}

// AllMFE returns a slice containing all MFE segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllMFE() ([]*MFE, error) {
	pss, err := m.ParseAll("MFE")
	return pss.([]*MFE), err
}

// AllMFI returns a slice containing all MFI segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllMFI() ([]*MFI, error) {
	pss, err := m.ParseAll("MFI")
	return pss.([]*MFI), err
}

// AllMRG returns a slice containing all MRG segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllMRG() ([]*MRG, error) {
	pss, err := m.ParseAll("MRG")
	return pss.([]*MRG), err
}

// AllMSA returns a slice containing all MSA segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllMSA() ([]*MSA, error) {
	pss, err := m.ParseAll("MSA")
	return pss.([]*MSA), err
}

// AllMSH returns a slice containing all MSH segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllMSH() ([]*MSH, error) {
	pss, err := m.ParseAll("MSH")
	return pss.([]*MSH), err
}

// AllNCK returns a slice containing all NCK segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllNCK() ([]*NCK, error) {
	pss, err := m.ParseAll("NCK")
	return pss.([]*NCK), err
}

// AllNK1 returns a slice containing all NK1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllNK1() ([]*NK1, error) {
	pss, err := m.ParseAll("NK1")
	return pss.([]*NK1), err
}

// AllNPU returns a slice containing all NPU segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllNPU() ([]*NPU, error) {
	pss, err := m.ParseAll("NPU")
	return pss.([]*NPU), err
}

// AllNSC returns a slice containing all NSC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllNSC() ([]*NSC, error) {
	pss, err := m.ParseAll("NSC")
	return pss.([]*NSC), err
}

// AllNST returns a slice containing all NST segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllNST() ([]*NST, error) {
	pss, err := m.ParseAll("NST")
	return pss.([]*NST), err
}

// AllNTE returns a slice containing all NTE segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllNTE() ([]*NTE, error) {
	pss, err := m.ParseAll("NTE")
	return pss.([]*NTE), err
}

// AllOBR returns a slice containing all OBR segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllOBR() ([]*OBR, error) {
	pss, err := m.ParseAll("OBR")
	return pss.([]*OBR), err
}

// AllOBX returns a slice containing all OBX segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllOBX() ([]*OBX, error) {
	pss, err := m.ParseAll("OBX")
	return pss.([]*OBX), err
}

// AllODS returns a slice containing all ODS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllODS() ([]*ODS, error) {
	pss, err := m.ParseAll("ODS")
	return pss.([]*ODS), err
}

// AllODT returns a slice containing all ODT segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllODT() ([]*ODT, error) {
	pss, err := m.ParseAll("ODT")
	return pss.([]*ODT), err
}

// AllOM1 returns a slice containing all OM1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllOM1() ([]*OM1, error) {
	pss, err := m.ParseAll("OM1")
	return pss.([]*OM1), err
}

// AllOM2 returns a slice containing all OM2 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllOM2() ([]*OM2, error) {
	pss, err := m.ParseAll("OM2")
	return pss.([]*OM2), err
}

// AllOM3 returns a slice containing all OM3 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllOM3() ([]*OM3, error) {
	pss, err := m.ParseAll("OM3")
	return pss.([]*OM3), err
}

// AllOM4 returns a slice containing all OM4 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllOM4() ([]*OM4, error) {
	pss, err := m.ParseAll("OM4")
	return pss.([]*OM4), err
}

// AllOM5 returns a slice containing all OM5 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllOM5() ([]*OM5, error) {
	pss, err := m.ParseAll("OM5")
	return pss.([]*OM5), err
}

// AllOM6 returns a slice containing all OM6 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllOM6() ([]*OM6, error) {
	pss, err := m.ParseAll("OM6")
	return pss.([]*OM6), err
}

// AllORC returns a slice containing all ORC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllORC() ([]*ORC, error) {
	pss, err := m.ParseAll("ORC")
	return pss.([]*ORC), err
}

// AllORO returns a slice containing all ORO segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllORO() ([]*ORO, error) {
	pss, err := m.ParseAll("ORO")
	return pss.([]*ORO), err
}

// AllPID returns a slice containing all PID segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPID() ([]*PID, error) {
	pss, err := m.ParseAll("PID")
	return pss.([]*PID), err
}

// AllPR1 returns a slice containing all PR1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPR1() ([]*PR1, error) {
	pss, err := m.ParseAll("PR1")
	return pss.([]*PR1), err
}

// AllPRA returns a slice containing all PRA segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPRA() ([]*PRA, error) {
	pss, err := m.ParseAll("PRA")
	return pss.([]*PRA), err
}

// AllPV1 returns a slice containing all PV1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPV1() ([]*PV1, error) {
	pss, err := m.ParseAll("PV1")
	return pss.([]*PV1), err
}

// AllPV2 returns a slice containing all PV2 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPV2() ([]*PV2, error) {
	pss, err := m.ParseAll("PV2")
	return pss.([]*PV2), err
}

// AllQRD returns a slice containing all QRD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllQRD() ([]*QRD, error) {
	pss, err := m.ParseAll("QRD")
	return pss.([]*QRD), err
}

// AllQRF returns a slice containing all QRF segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllQRF() ([]*QRF, error) {
	pss, err := m.ParseAll("QRF")
	return pss.([]*QRF), err
}

// AllRQ1 returns a slice containing all RQ1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRQ1() ([]*RQ1, error) {
	pss, err := m.ParseAll("RQ1")
	return pss.([]*RQ1), err
}

// AllRQD returns a slice containing all RQD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRQD() ([]*RQD, error) {
	pss, err := m.ParseAll("RQD")
	return pss.([]*RQD), err
}

// AllRX1 returns a slice containing all RX1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRX1() ([]*RX1, error) {
	pss, err := m.ParseAll("RX1")
	return pss.([]*RX1), err
}

// AllRXA returns a slice containing all RXA segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRXA() ([]*RXA, error) {
	pss, err := m.ParseAll("RXA")
	return pss.([]*RXA), err
}

// AllRXC returns a slice containing all RXC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRXC() ([]*RXC, error) {
	pss, err := m.ParseAll("RXC")
	return pss.([]*RXC), err
}

// AllRXD returns a slice containing all RXD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRXD() ([]*RXD, error) {
	pss, err := m.ParseAll("RXD")
	return pss.([]*RXD), err
}

// AllRXE returns a slice containing all RXE segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRXE() ([]*RXE, error) {
	pss, err := m.ParseAll("RXE")
	return pss.([]*RXE), err
}

// AllRXG returns a slice containing all RXG segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRXG() ([]*RXG, error) {
	pss, err := m.ParseAll("RXG")
	return pss.([]*RXG), err
}

// AllRXO returns a slice containing all RXO segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRXO() ([]*RXO, error) {
	pss, err := m.ParseAll("RXO")
	return pss.([]*RXO), err
}

// AllRXR returns a slice containing all RXR segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRXR() ([]*RXR, error) {
	pss, err := m.ParseAll("RXR")
	return pss.([]*RXR), err
}

// AllSTF returns a slice containing all STF segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllSTF() ([]*STF, error) {
	pss, err := m.ParseAll("STF")
	return pss.([]*STF), err
}

// AllUB1 returns a slice containing all UB1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllUB1() ([]*UB1, error) {
	pss, err := m.ParseAll("UB1")
	return pss.([]*UB1), err
}

// AllUB2 returns a slice containing all UB2 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllUB2() ([]*UB2, error) {
	pss, err := m.ParseAll("UB2")
	return pss.([]*UB2), err
}

// AllURD returns a slice containing all URD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllURD() ([]*URD, error) {
	pss, err := m.ParseAll("URD")
	return pss.([]*URD), err
}

// AllURS returns a slice containing all URS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllURS() ([]*URS, error) {
	pss, err := m.ParseAll("URS")
	return pss.([]*URS), err
}

// v2 API
type ACKv2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
}

func (m *ACKv2) MSH() *MSH {
	return m.msh
}

func (m *ACKv2) MSA() *MSA {
	return m.msa
}

func (m *ACKv2) ERR() *ERR {
	return m.err
}

func (m ACKv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
	}, nil
}

type ADR_A19v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD             // Required
	pid []*ADR_A19_PIDv2 // Required
	dsc *DSC
}

func (m *ADR_A19v2) MSH() *MSH {
	return m.msh
}

func (m *ADR_A19v2) MSA() *MSA {
	return m.msa
}

func (m *ADR_A19v2) ERR() *ERR {
	return m.err
}

func (m *ADR_A19v2) QRD() *QRD {
	return m.qrd
}

func (m *ADR_A19v2) GroupByPID() []*ADR_A19_PIDv2 {
	return m.pid
}

func (m *ADR_A19v2) DSC() *DSC {
	return m.dsc
}

func (m ADR_A19v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"pid": m.pid,
		"dsc": m.dsc,
	}, nil
}

type ADR_A19_PIDv2 struct {
	evn *EVN
	pid *PID // Required
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADR_A19_PID_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADR_A19_PIDv2) EVN() *EVN {
	return m.evn
}

func (m *ADR_A19_PIDv2) PID() *PID {
	return m.pid
}

func (m *ADR_A19_PIDv2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADR_A19_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *ADR_A19_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *ADR_A19_PIDv2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADR_A19_PIDv2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADR_A19_PIDv2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADR_A19_PIDv2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADR_A19_PIDv2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADR_A19_PIDv2) GroupByIN1() []*ADR_A19_PID_IN1v2 {
	return m.in1
}

func (m *ADR_A19_PIDv2) ACC() *ACC {
	return m.acc
}

func (m *ADR_A19_PIDv2) UB1() *UB1 {
	return m.ub1
}

func (m *ADR_A19_PIDv2) UB2() *UB2 {
	return m.ub2
}

func (m ADR_A19_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}

type ADR_A19_PID_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADR_A19_PID_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADR_A19_PID_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADR_A19_PID_IN1v2) IN3() *IN3 {
	return m.in3
}

func (m ADR_A19_PID_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}

type ADT_A01v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A01_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A01v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A01v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A01v2) PID() *PID {
	return m.pid
}

func (m *ADT_A01v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A01v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A01v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A01v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A01v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A01v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A01v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A01v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A01v2) GroupByIN1() []*ADT_A01_IN1v2 {
	return m.in1
}

func (m *ADT_A01v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A01v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A01v2) UB2() *UB2 {
	return m.ub2
}

func (m ADT_A01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}

type ADT_A01_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A01_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A01_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A01_IN1v2) IN3() *IN3 {
	return m.in3
}

func (m ADT_A01_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}

type ADT_A02v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A02v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A02v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A02v2) PID() *PID {
	return m.pid
}

func (m *ADT_A02v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A02v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A02v2) AllOBX() []*OBX {
	return m.obx
}

func (m ADT_A02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}

type ADT_A03v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A03v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A03v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A03v2) PID() *PID {
	return m.pid
}

func (m *ADT_A03v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A03v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A03v2) AllOBX() []*OBX {
	return m.obx
}

func (m ADT_A03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}

type ADT_A04v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A04_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A04v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A04v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A04v2) PID() *PID {
	return m.pid
}

func (m *ADT_A04v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A04v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A04v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A04v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A04v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A04v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A04v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A04v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A04v2) GroupByIN1() []*ADT_A04_IN1v2 {
	return m.in1
}

func (m *ADT_A04v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A04v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A04v2) UB2() *UB2 {
	return m.ub2
}

func (m ADT_A04v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}

type ADT_A04_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A04_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A04_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A04_IN1v2) IN3() *IN3 {
	return m.in3
}

func (m ADT_A04_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}

type ADT_A05v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A05_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A05v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A05v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A05v2) PID() *PID {
	return m.pid
}

func (m *ADT_A05v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A05v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A05v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A05v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A05v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A05v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A05v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A05v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A05v2) GroupByIN1() []*ADT_A05_IN1v2 {
	return m.in1
}

func (m *ADT_A05v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A05v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A05v2) UB2() *UB2 {
	return m.ub2
}

func (m ADT_A05v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}

type ADT_A05_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A05_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A05_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A05_IN1v2) IN3() *IN3 {
	return m.in3
}

func (m ADT_A05_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}

type ADT_A06v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	mrg *MRG
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A06_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A06v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A06v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A06v2) PID() *PID {
	return m.pid
}

func (m *ADT_A06v2) MRG() *MRG {
	return m.mrg
}

func (m *ADT_A06v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A06v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A06v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A06v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A06v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A06v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A06v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A06v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A06v2) GroupByIN1() []*ADT_A06_IN1v2 {
	return m.in1
}

func (m *ADT_A06v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A06v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A06v2) UB2() *UB2 {
	return m.ub2
}

func (m ADT_A06v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"mrg": m.mrg,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}

type ADT_A06_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A06_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A06_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A06_IN1v2) IN3() *IN3 {
	return m.in3
}

func (m ADT_A06_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}

type ADT_A07v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	mrg *MRG
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A07_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A07v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A07v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A07v2) PID() *PID {
	return m.pid
}

func (m *ADT_A07v2) MRG() *MRG {
	return m.mrg
}

func (m *ADT_A07v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A07v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A07v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A07v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A07v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A07v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A07v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A07v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A07v2) GroupByIN1() []*ADT_A07_IN1v2 {
	return m.in1
}

func (m *ADT_A07v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A07v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A07v2) UB2() *UB2 {
	return m.ub2
}

func (m ADT_A07v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"mrg": m.mrg,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}

type ADT_A07_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A07_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A07_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A07_IN1v2) IN3() *IN3 {
	return m.in3
}

func (m ADT_A07_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}

type ADT_A08v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A08_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A08v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A08v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A08v2) PID() *PID {
	return m.pid
}

func (m *ADT_A08v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A08v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A08v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A08v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A08v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A08v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A08v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A08v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A08v2) GroupByIN1() []*ADT_A08_IN1v2 {
	return m.in1
}

func (m *ADT_A08v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A08v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A08v2) UB2() *UB2 {
	return m.ub2
}

func (m ADT_A08v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}

type ADT_A08_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A08_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A08_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A08_IN1v2) IN3() *IN3 {
	return m.in3
}

func (m ADT_A08_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}

type ADT_A09v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	dg1 []*DG1
}

func (m *ADT_A09v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A09v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A09v2) PID() *PID {
	return m.pid
}

func (m *ADT_A09v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A09v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A09v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A09v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m ADT_A09v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"dg1": m.dg1,
	}, nil
}

type ADT_A10v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	dg1 []*DG1
}

func (m *ADT_A10v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A10v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A10v2) PID() *PID {
	return m.pid
}

func (m *ADT_A10v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A10v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A10v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A10v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m ADT_A10v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"dg1": m.dg1,
	}, nil
}

type ADT_A11v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	dg1 []*DG1
}

func (m *ADT_A11v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A11v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A11v2) PID() *PID {
	return m.pid
}

func (m *ADT_A11v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A11v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A11v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A11v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m ADT_A11v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"dg1": m.dg1,
	}, nil
}

type ADT_A12v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	dg1 []*DG1
}

func (m *ADT_A12v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A12v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A12v2) PID() *PID {
	return m.pid
}

func (m *ADT_A12v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A12v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A12v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A12v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m ADT_A12v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"dg1": m.dg1,
	}, nil
}

type ADT_A13v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A13_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A13v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A13v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A13v2) PID() *PID {
	return m.pid
}

func (m *ADT_A13v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A13v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A13v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A13v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A13v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A13v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A13v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A13v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A13v2) GroupByIN1() []*ADT_A13_IN1v2 {
	return m.in1
}

func (m *ADT_A13v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A13v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A13v2) UB2() *UB2 {
	return m.ub2
}

func (m ADT_A13v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}

type ADT_A13_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A13_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A13_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A13_IN1v2) IN3() *IN3 {
	return m.in3
}

func (m ADT_A13_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}

type ADT_A14v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A14_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A14v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A14v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A14v2) PID() *PID {
	return m.pid
}

func (m *ADT_A14v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A14v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A14v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A14v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A14v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A14v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A14v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A14v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A14v2) GroupByIN1() []*ADT_A14_IN1v2 {
	return m.in1
}

func (m *ADT_A14v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A14v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A14v2) UB2() *UB2 {
	return m.ub2
}

func (m ADT_A14v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}

type ADT_A14_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A14_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A14_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A14_IN1v2) IN3() *IN3 {
	return m.in3
}

func (m ADT_A14_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}

type ADT_A15v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	dg1 []*DG1
}

func (m *ADT_A15v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A15v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A15v2) PID() *PID {
	return m.pid
}

func (m *ADT_A15v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A15v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A15v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A15v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m ADT_A15v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"dg1": m.dg1,
	}, nil
}

type ADT_A16v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	dg1 []*DG1
}

func (m *ADT_A16v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A16v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A16v2) PID() *PID {
	return m.pid
}

func (m *ADT_A16v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A16v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A16v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A16v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m ADT_A16v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"dg1": m.dg1,
	}, nil
}

type ADT_A17v2 struct {
	msh  *MSH // Required
	evn  *EVN // Required
	pid1 *PID // Required
	pv11 *PV1 // Required
	pv21 *PV2
	obx1 []*OBX
	pid2 *PID // Required
	pv12 *PV1 // Required
	pv22 *PV2
	obx2 []*OBX
}

func (m *ADT_A17v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A17v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A17v2) PID1() *PID {
	return m.pid1
}

func (m *ADT_A17v2) PV11() *PV1 {
	return m.pv11
}

func (m *ADT_A17v2) PV21() *PV2 {
	return m.pv21
}

func (m *ADT_A17v2) AllOBX1() []*OBX {
	return m.obx1
}

func (m *ADT_A17v2) PID2() *PID {
	return m.pid2
}

func (m *ADT_A17v2) PV12() *PV1 {
	return m.pv12
}

func (m *ADT_A17v2) PV22() *PV2 {
	return m.pv22
}

func (m *ADT_A17v2) AllOBX2() []*OBX {
	return m.obx2
}

func (m ADT_A17v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh":  m.msh,
		"evn":  m.evn,
		"pid1": m.pid1,
		"pv11": m.pv11,
		"pv21": m.pv21,
		"obx1": m.obx1,
		"pid2": m.pid2,
		"pv12": m.pv12,
		"pv22": m.pv22,
		"obx2": m.obx2,
	}, nil
}

type ADT_A18v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	mrg *MRG
	pv1 *PV1 // Required
}

func (m *ADT_A18v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A18v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A18v2) PID() *PID {
	return m.pid
}

func (m *ADT_A18v2) MRG() *MRG {
	return m.mrg
}

func (m *ADT_A18v2) PV1() *PV1 {
	return m.pv1
}

func (m ADT_A18v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"mrg": m.mrg,
		"pv1": m.pv1,
	}, nil
}

type ADT_A20v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	npu *NPU // Required
}

func (m *ADT_A20v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A20v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A20v2) NPU() *NPU {
	return m.npu
}

func (m ADT_A20v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"npu": m.npu,
	}, nil
}

type ADT_A21v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A21v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A21v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A21v2) PID() *PID {
	return m.pid
}

func (m *ADT_A21v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A21v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A21v2) AllOBX() []*OBX {
	return m.obx
}

func (m ADT_A21v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}

type ADT_A22v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A22v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A22v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A22v2) PID() *PID {
	return m.pid
}

func (m *ADT_A22v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A22v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A22v2) AllOBX() []*OBX {
	return m.obx
}

func (m ADT_A22v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}

type ADT_A23v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A23v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A23v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A23v2) PID() *PID {
	return m.pid
}

func (m *ADT_A23v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A23v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A23v2) AllOBX() []*OBX {
	return m.obx
}

func (m ADT_A23v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}

type ADT_A24v2 struct {
	msh  *MSH // Required
	evn  *EVN // Required
	pid1 *PID // Required
	pv11 *PV1
	pid2 *PID // Required
	pv12 *PV1
}

func (m *ADT_A24v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A24v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A24v2) PID1() *PID {
	return m.pid1
}

func (m *ADT_A24v2) PV11() *PV1 {
	return m.pv11
}

func (m *ADT_A24v2) PID2() *PID {
	return m.pid2
}

func (m *ADT_A24v2) PV12() *PV1 {
	return m.pv12
}

func (m ADT_A24v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh":  m.msh,
		"evn":  m.evn,
		"pid1": m.pid1,
		"pv11": m.pv11,
		"pid2": m.pid2,
		"pv12": m.pv12,
	}, nil
}

type ADT_A25v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A25v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A25v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A25v2) PID() *PID {
	return m.pid
}

func (m *ADT_A25v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A25v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A25v2) AllOBX() []*OBX {
	return m.obx
}

func (m ADT_A25v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}

type ADT_A26v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A26v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A26v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A26v2) PID() *PID {
	return m.pid
}

func (m *ADT_A26v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A26v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A26v2) AllOBX() []*OBX {
	return m.obx
}

func (m ADT_A26v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}

type ADT_A27v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A27v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A27v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A27v2) PID() *PID {
	return m.pid
}

func (m *ADT_A27v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A27v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A27v2) AllOBX() []*OBX {
	return m.obx
}

func (m ADT_A27v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}

type ADT_A28v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A28_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A28v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A28v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A28v2) PID() *PID {
	return m.pid
}

func (m *ADT_A28v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A28v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A28v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A28v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A28v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A28v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A28v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A28v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A28v2) GroupByIN1() []*ADT_A28_IN1v2 {
	return m.in1
}

func (m *ADT_A28v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A28v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A28v2) UB2() *UB2 {
	return m.ub2
}

func (m ADT_A28v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}

type ADT_A28_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A28_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A28_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A28_IN1v2) IN3() *IN3 {
	return m.in3
}

func (m ADT_A28_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}

type ADT_A29v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A29v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A29v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A29v2) PID() *PID {
	return m.pid
}

func (m *ADT_A29v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A29v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A29v2) AllOBX() []*OBX {
	return m.obx
}

func (m ADT_A29v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}

type ADT_A30v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	mrg *MRG // Required
}

func (m *ADT_A30v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A30v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A30v2) PID() *PID {
	return m.pid
}

func (m *ADT_A30v2) MRG() *MRG {
	return m.mrg
}

func (m ADT_A30v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"mrg": m.mrg,
	}, nil
}

type ADT_A31v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 []*NK1
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	in1 []*ADT_A31_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *ADT_A31v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A31v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A31v2) PID() *PID {
	return m.pid
}

func (m *ADT_A31v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ADT_A31v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A31v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A31v2) AllOBX() []*OBX {
	return m.obx
}

func (m *ADT_A31v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *ADT_A31v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *ADT_A31v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *ADT_A31v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *ADT_A31v2) GroupByIN1() []*ADT_A31_IN1v2 {
	return m.in1
}

func (m *ADT_A31v2) ACC() *ACC {
	return m.acc
}

func (m *ADT_A31v2) UB1() *UB1 {
	return m.ub1
}

func (m *ADT_A31v2) UB2() *UB2 {
	return m.ub2
}

func (m ADT_A31v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}

type ADT_A31_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *ADT_A31_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *ADT_A31_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *ADT_A31_IN1v2) IN3() *IN3 {
	return m.in3
}

func (m ADT_A31_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}

type ADT_A32v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A32v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A32v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A32v2) PID() *PID {
	return m.pid
}

func (m *ADT_A32v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A32v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A32v2) AllOBX() []*OBX {
	return m.obx
}

func (m ADT_A32v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}

type ADT_A33v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	pv2 *PV2
	obx []*OBX
}

func (m *ADT_A33v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A33v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A33v2) PID() *PID {
	return m.pid
}

func (m *ADT_A33v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A33v2) PV2() *PV2 {
	return m.pv2
}

func (m *ADT_A33v2) AllOBX() []*OBX {
	return m.obx
}

func (m ADT_A33v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
	}, nil
}

type ADT_A34v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	mrg *MRG // Required
}

func (m *ADT_A34v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A34v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A34v2) PID() *PID {
	return m.pid
}

func (m *ADT_A34v2) MRG() *MRG {
	return m.mrg
}

func (m ADT_A34v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"mrg": m.mrg,
	}, nil
}

type ADT_A35v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	mrg *MRG // Required
}

func (m *ADT_A35v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A35v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A35v2) PID() *PID {
	return m.pid
}

func (m *ADT_A35v2) MRG() *MRG {
	return m.mrg
}

func (m ADT_A35v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"mrg": m.mrg,
	}, nil
}

type ADT_A36v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	mrg *MRG // Required
}

func (m *ADT_A36v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A36v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A36v2) PID() *PID {
	return m.pid
}

func (m *ADT_A36v2) MRG() *MRG {
	return m.mrg
}

func (m ADT_A36v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"mrg": m.mrg,
	}, nil
}

type ADT_A37v2 struct {
	msh  *MSH // Required
	evn  *EVN // Required
	pid1 *PID // Required
	pv11 *PV1
	pid2 *PID // Required
	pv12 *PV1
}

func (m *ADT_A37v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A37v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A37v2) PID1() *PID {
	return m.pid1
}

func (m *ADT_A37v2) PV11() *PV1 {
	return m.pv11
}

func (m *ADT_A37v2) PID2() *PID {
	return m.pid2
}

func (m *ADT_A37v2) PV12() *PV1 {
	return m.pv12
}

func (m ADT_A37v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh":  m.msh,
		"evn":  m.evn,
		"pid1": m.pid1,
		"pv11": m.pv11,
		"pid2": m.pid2,
		"pv12": m.pv12,
	}, nil
}

type ADT_A40v2 struct {
	msh *MSH             // Required
	evn *EVN             // Required
	pid []*ADT_A40_PIDv2 // Required
}

func (m *ADT_A40v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A40v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A40v2) GroupByPID() []*ADT_A40_PIDv2 {
	return m.pid
}

func (m ADT_A40v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
	}, nil
}

type ADT_A40_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	mrg *MRG // Required
	pv1 *PV1
}

func (m *ADT_A40_PIDv2) PID() *PID {
	return m.pid
}

func (m *ADT_A40_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A40_PIDv2) MRG() *MRG {
	return m.mrg
}

func (m *ADT_A40_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m ADT_A40_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"mrg": m.mrg,
		"pv1": m.pv1,
	}, nil
}

type ADT_A44v2 struct {
	msh *MSH             // Required
	evn *EVN             // Required
	pid []*ADT_A44_PIDv2 // Required
}

func (m *ADT_A44v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A44v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A44v2) GroupByPID() []*ADT_A44_PIDv2 {
	return m.pid
}

func (m ADT_A44v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
	}, nil
}

type ADT_A44_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	mrg *MRG // Required
}

func (m *ADT_A44_PIDv2) PID() *PID {
	return m.pid
}

func (m *ADT_A44_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A44_PIDv2) MRG() *MRG {
	return m.mrg
}

func (m ADT_A44_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"mrg": m.mrg,
	}, nil
}

type BAR_P01v2 struct {
	msh *MSH             // Required
	evn *EVN             // Required
	pid *PID             // Required
	pv1 []*BAR_P01_PV1v2 // Required
}

func (m *BAR_P01v2) MSH() *MSH {
	return m.msh
}

func (m *BAR_P01v2) EVN() *EVN {
	return m.evn
}

func (m *BAR_P01v2) PID() *PID {
	return m.pid
}

func (m *BAR_P01v2) GroupByPV1() []*BAR_P01_PV1v2 {
	return m.pv1
}

func (m BAR_P01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
	}, nil
}

type BAR_P01_PV1v2 struct {
	pv1 *PV1
	pv2 *PV2
	obx []*OBX
	al1 []*AL1
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	nk1 []*NK1
	in1 []*BAR_P01_PV1_IN1v2
	acc *ACC
	ub1 *UB1
	ub2 *UB2
}

func (m *BAR_P01_PV1v2) PV1() *PV1 {
	return m.pv1
}

func (m *BAR_P01_PV1v2) PV2() *PV2 {
	return m.pv2
}

func (m *BAR_P01_PV1v2) AllOBX() []*OBX {
	return m.obx
}

func (m *BAR_P01_PV1v2) AllAL1() []*AL1 {
	return m.al1
}

func (m *BAR_P01_PV1v2) AllDG1() []*DG1 {
	return m.dg1
}

func (m *BAR_P01_PV1v2) AllPR1() []*PR1 {
	return m.pr1
}

func (m *BAR_P01_PV1v2) AllGT1() []*GT1 {
	return m.gt1
}

func (m *BAR_P01_PV1v2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *BAR_P01_PV1v2) GroupByIN1() []*BAR_P01_PV1_IN1v2 {
	return m.in1
}

func (m *BAR_P01_PV1v2) ACC() *ACC {
	return m.acc
}

func (m *BAR_P01_PV1v2) UB1() *UB1 {
	return m.ub1
}

func (m *BAR_P01_PV1v2) UB2() *UB2 {
	return m.ub2
}

func (m BAR_P01_PV1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"al1": m.al1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"nk1": m.nk1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
		"ub2": m.ub2,
	}, nil
}

type BAR_P01_PV1_IN1v2 struct {
	in1 *IN1 // Required
	in2 *IN2
	in3 *IN3
}

func (m *BAR_P01_PV1_IN1v2) IN1() *IN1 {
	return m.in1
}

func (m *BAR_P01_PV1_IN1v2) IN2() *IN2 {
	return m.in2
}

func (m *BAR_P01_PV1_IN1v2) IN3() *IN3 {
	return m.in3
}

func (m BAR_P01_PV1_IN1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"in1": m.in1,
		"in2": m.in2,
		"in3": m.in3,
	}, nil
}

type BAR_P02v2 struct {
	msh *MSH             // Required
	evn *EVN             // Required
	pid []*BAR_P02_PIDv2 // Required
}

func (m *BAR_P02v2) MSH() *MSH {
	return m.msh
}

func (m *BAR_P02v2) EVN() *EVN {
	return m.evn
}

func (m *BAR_P02v2) GroupByPID() []*BAR_P02_PIDv2 {
	return m.pid
}

func (m BAR_P02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
	}, nil
}

type BAR_P02_PIDv2 struct {
	pid *PID // Required
	pv1 *PV1
}

func (m *BAR_P02_PIDv2) PID() *PID {
	return m.pid
}

func (m *BAR_P02_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m BAR_P02_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pv1": m.pv1,
	}, nil
}

type DFT_P03v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1
	pv2 *PV2
	obx []*OBX
	ft1 []*FT1 // Required
}

func (m *DFT_P03v2) MSH() *MSH {
	return m.msh
}

func (m *DFT_P03v2) EVN() *EVN {
	return m.evn
}

func (m *DFT_P03v2) PID() *PID {
	return m.pid
}

func (m *DFT_P03v2) PV1() *PV1 {
	return m.pv1
}

func (m *DFT_P03v2) PV2() *PV2 {
	return m.pv2
}

func (m *DFT_P03v2) AllOBX() []*OBX {
	return m.obx
}

func (m *DFT_P03v2) AllFT1() []*FT1 {
	return m.ft1
}

func (m DFT_P03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obx": m.obx,
		"ft1": m.ft1,
	}, nil
}

type DSR_P04v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD // Required
	qrf *QRF
	dsp []*DSP // Required
	dsc *DSC
}

func (m *DSR_P04v2) MSH() *MSH {
	return m.msh
}

func (m *DSR_P04v2) MSA() *MSA {
	return m.msa
}

func (m *DSR_P04v2) ERR() *ERR {
	return m.err
}

func (m *DSR_P04v2) QRD() *QRD {
	return m.qrd
}

func (m *DSR_P04v2) QRF() *QRF {
	return m.qrf
}

func (m *DSR_P04v2) AllDSP() []*DSP {
	return m.dsp
}

func (m *DSR_P04v2) DSC() *DSC {
	return m.dsc
}

func (m DSR_P04v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsp": m.dsp,
		"dsc": m.dsc,
	}, nil
}

type DSR_Q01v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD // Required
	qrf *QRF
	dsp []*DSP // Required
	dsc *DSC
}

func (m *DSR_Q01v2) MSH() *MSH {
	return m.msh
}

func (m *DSR_Q01v2) MSA() *MSA {
	return m.msa
}

func (m *DSR_Q01v2) ERR() *ERR {
	return m.err
}

func (m *DSR_Q01v2) QRD() *QRD {
	return m.qrd
}

func (m *DSR_Q01v2) QRF() *QRF {
	return m.qrf
}

func (m *DSR_Q01v2) AllDSP() []*DSP {
	return m.dsp
}

func (m *DSR_Q01v2) DSC() *DSC {
	return m.dsc
}

func (m DSR_Q01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsp": m.dsp,
		"dsc": m.dsc,
	}, nil
}

type DSR_Q03v2 struct {
	msh *MSH // Required
	msa *MSA
	qrd *QRD // Required
	qrf *QRF
	dsp []*DSP // Required
	dsc *DSC
}

func (m *DSR_Q03v2) MSH() *MSH {
	return m.msh
}

func (m *DSR_Q03v2) MSA() *MSA {
	return m.msa
}

func (m *DSR_Q03v2) QRD() *QRD {
	return m.qrd
}

func (m *DSR_Q03v2) QRF() *QRF {
	return m.qrf
}

func (m *DSR_Q03v2) AllDSP() []*DSP {
	return m.dsp
}

func (m *DSR_Q03v2) DSC() *DSC {
	return m.dsc
}

func (m DSR_Q03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsp": m.dsp,
		"dsc": m.dsc,
	}, nil
}

type DSR_R03v2 struct {
	msh *MSH // Required
	msa *MSA
	qrd *QRD // Required
	qrf *QRF
	dsp []*DSP // Required
	dsc *DSC
}

func (m *DSR_R03v2) MSH() *MSH {
	return m.msh
}

func (m *DSR_R03v2) MSA() *MSA {
	return m.msa
}

func (m *DSR_R03v2) QRD() *QRD {
	return m.qrd
}

func (m *DSR_R03v2) QRF() *QRF {
	return m.qrf
}

func (m *DSR_R03v2) AllDSP() []*DSP {
	return m.dsp
}

func (m *DSR_R03v2) DSC() *DSC {
	return m.dsc
}

func (m DSR_R03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsp": m.dsp,
		"dsc": m.dsc,
	}, nil
}

type MCF_Q02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
}

func (m *MCF_Q02v2) MSH() *MSH {
	return m.msh
}

func (m *MCF_Q02v2) MSA() *MSA {
	return m.msa
}

func (m MCF_Q02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
	}, nil
}

type MFD_M01v2 struct {
	msh *MSH // Required
	mfi *MFI // Required
	mfa []*MFA
}

func (m *MFD_M01v2) MSH() *MSH {
	return m.msh
}

func (m *MFD_M01v2) MFI() *MFI {
	return m.mfi
}

func (m *MFD_M01v2) AllMFA() []*MFA {
	return m.mfa
}

func (m MFD_M01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfa": m.mfa,
	}, nil
}

type MFD_M02v2 struct {
	msh *MSH // Required
	mfi *MFI // Required
	mfa []*MFA
}

func (m *MFD_M02v2) MSH() *MSH {
	return m.msh
}

func (m *MFD_M02v2) MFI() *MFI {
	return m.mfi
}

func (m *MFD_M02v2) AllMFA() []*MFA {
	return m.mfa
}

func (m MFD_M02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfa": m.mfa,
	}, nil
}

type MFD_M03v2 struct {
	msh *MSH // Required
	mfi *MFI // Required
	mfa []*MFA
}

func (m *MFD_M03v2) MSH() *MSH {
	return m.msh
}

func (m *MFD_M03v2) MFI() *MFI {
	return m.mfi
}

func (m *MFD_M03v2) AllMFA() []*MFA {
	return m.mfa
}

func (m MFD_M03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfa": m.mfa,
	}, nil
}

type MFK_M01v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	mfi *MFI // Required
	mfa []*MFA
}

func (m *MFK_M01v2) MSH() *MSH {
	return m.msh
}

func (m *MFK_M01v2) MSA() *MSA {
	return m.msa
}

func (m *MFK_M01v2) ERR() *ERR {
	return m.err
}

func (m *MFK_M01v2) MFI() *MFI {
	return m.mfi
}

func (m *MFK_M01v2) AllMFA() []*MFA {
	return m.mfa
}

func (m MFK_M01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"mfi": m.mfi,
		"mfa": m.mfa,
	}, nil
}

type MFK_M02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	mfi *MFI // Required
	mfa []*MFA
}

func (m *MFK_M02v2) MSH() *MSH {
	return m.msh
}

func (m *MFK_M02v2) MSA() *MSA {
	return m.msa
}

func (m *MFK_M02v2) ERR() *ERR {
	return m.err
}

func (m *MFK_M02v2) MFI() *MFI {
	return m.mfi
}

func (m *MFK_M02v2) AllMFA() []*MFA {
	return m.mfa
}

func (m MFK_M02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"mfi": m.mfi,
		"mfa": m.mfa,
	}, nil
}

type MFK_M03v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	mfi *MFI // Required
	mfa []*MFA
}

func (m *MFK_M03v2) MSH() *MSH {
	return m.msh
}

func (m *MFK_M03v2) MSA() *MSA {
	return m.msa
}

func (m *MFK_M03v2) ERR() *ERR {
	return m.err
}

func (m *MFK_M03v2) MFI() *MFI {
	return m.mfi
}

func (m *MFK_M03v2) AllMFA() []*MFA {
	return m.mfa
}

func (m MFK_M03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"mfi": m.mfi,
		"mfa": m.mfa,
	}, nil
}

type MFN_M01v2 struct {
	msh *MSH             // Required
	mfi *MFI             // Required
	mfe []*MFN_M01_MFEv2 // Required
}

func (m *MFN_M01v2) MSH() *MSH {
	return m.msh
}

func (m *MFN_M01v2) MFI() *MFI {
	return m.mfi
}

func (m *MFN_M01v2) GroupByMFE() []*MFN_M01_MFEv2 {
	return m.mfe
}

func (m MFN_M01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfe": m.mfe,
	}, nil
}

type MFN_M01_MFEv2 struct {
	mfe *MFE // Required
}

func (m *MFN_M01_MFEv2) MFE() *MFE {
	return m.mfe
}

func (m MFN_M01_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
	}, nil
}

type MFN_M02v2 struct {
	msh *MSH             // Required
	mfi *MFI             // Required
	mfe []*MFN_M02_MFEv2 // Required
}

func (m *MFN_M02v2) MSH() *MSH {
	return m.msh
}

func (m *MFN_M02v2) MFI() *MFI {
	return m.mfi
}

func (m *MFN_M02v2) GroupByMFE() []*MFN_M02_MFEv2 {
	return m.mfe
}

func (m MFN_M02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfe": m.mfe,
	}, nil
}

type MFN_M02_MFEv2 struct {
	mfe *MFE // Required
}

func (m *MFN_M02_MFEv2) MFE() *MFE {
	return m.mfe
}

func (m MFN_M02_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
	}, nil
}

type MFN_M03v2 struct {
	msh *MSH             // Required
	mfi *MFI             // Required
	mfe []*MFN_M03_MFEv2 // Required
}

func (m *MFN_M03v2) MSH() *MSH {
	return m.msh
}

func (m *MFN_M03v2) MFI() *MFI {
	return m.mfi
}

func (m *MFN_M03v2) GroupByMFE() []*MFN_M03_MFEv2 {
	return m.mfe
}

func (m MFN_M03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"mfi": m.mfi,
		"mfe": m.mfe,
	}, nil
}

type MFN_M03_MFEv2 struct {
	mfe *MFE // Required
}

func (m *MFN_M03_MFEv2) MFE() *MFE {
	return m.mfe
}

func (m MFN_M03_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
	}, nil
}

type MFQ_M01v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	dsc *DSC
}

func (m *MFQ_M01v2) MSH() *MSH {
	return m.msh
}

func (m *MFQ_M01v2) QRD() *QRD {
	return m.qrd
}

func (m *MFQ_M01v2) QRF() *QRF {
	return m.qrf
}

func (m *MFQ_M01v2) DSC() *DSC {
	return m.dsc
}

func (m MFQ_M01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsc": m.dsc,
	}, nil
}

type MFQ_M02v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	dsc *DSC
}

func (m *MFQ_M02v2) MSH() *MSH {
	return m.msh
}

func (m *MFQ_M02v2) QRD() *QRD {
	return m.qrd
}

func (m *MFQ_M02v2) QRF() *QRF {
	return m.qrf
}

func (m *MFQ_M02v2) DSC() *DSC {
	return m.dsc
}

func (m MFQ_M02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsc": m.dsc,
	}, nil
}

type MFQ_M03v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	dsc *DSC
}

func (m *MFQ_M03v2) MSH() *MSH {
	return m.msh
}

func (m *MFQ_M03v2) QRD() *QRD {
	return m.qrd
}

func (m *MFQ_M03v2) QRF() *QRF {
	return m.qrf
}

func (m *MFQ_M03v2) DSC() *DSC {
	return m.dsc
}

func (m MFQ_M03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsc": m.dsc,
	}, nil
}

type MFR_M01v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD // Required
	qrf *QRF
	mfi *MFI             // Required
	mfe []*MFR_M01_MFEv2 // Required
	dsc *DSC
}

func (m *MFR_M01v2) MSH() *MSH {
	return m.msh
}

func (m *MFR_M01v2) MSA() *MSA {
	return m.msa
}

func (m *MFR_M01v2) ERR() *ERR {
	return m.err
}

func (m *MFR_M01v2) QRD() *QRD {
	return m.qrd
}

func (m *MFR_M01v2) QRF() *QRF {
	return m.qrf
}

func (m *MFR_M01v2) MFI() *MFI {
	return m.mfi
}

func (m *MFR_M01v2) GroupByMFE() []*MFR_M01_MFEv2 {
	return m.mfe
}

func (m *MFR_M01v2) DSC() *DSC {
	return m.dsc
}

func (m MFR_M01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"mfi": m.mfi,
		"mfe": m.mfe,
		"dsc": m.dsc,
	}, nil
}

type MFR_M01_MFEv2 struct {
	mfe *MFE // Required
}

func (m *MFR_M01_MFEv2) MFE() *MFE {
	return m.mfe
}

func (m MFR_M01_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
	}, nil
}

type MFR_M02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD // Required
	qrf *QRF
	mfi *MFI             // Required
	mfe []*MFR_M02_MFEv2 // Required
	dsc *DSC
}

func (m *MFR_M02v2) MSH() *MSH {
	return m.msh
}

func (m *MFR_M02v2) MSA() *MSA {
	return m.msa
}

func (m *MFR_M02v2) ERR() *ERR {
	return m.err
}

func (m *MFR_M02v2) QRD() *QRD {
	return m.qrd
}

func (m *MFR_M02v2) QRF() *QRF {
	return m.qrf
}

func (m *MFR_M02v2) MFI() *MFI {
	return m.mfi
}

func (m *MFR_M02v2) GroupByMFE() []*MFR_M02_MFEv2 {
	return m.mfe
}

func (m *MFR_M02v2) DSC() *DSC {
	return m.dsc
}

func (m MFR_M02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"mfi": m.mfi,
		"mfe": m.mfe,
		"dsc": m.dsc,
	}, nil
}

type MFR_M02_MFEv2 struct {
	mfe *MFE // Required
}

func (m *MFR_M02_MFEv2) MFE() *MFE {
	return m.mfe
}

func (m MFR_M02_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
	}, nil
}

type MFR_M03v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD // Required
	qrf *QRF
	mfi *MFI             // Required
	mfe []*MFR_M03_MFEv2 // Required
	dsc *DSC
}

func (m *MFR_M03v2) MSH() *MSH {
	return m.msh
}

func (m *MFR_M03v2) MSA() *MSA {
	return m.msa
}

func (m *MFR_M03v2) ERR() *ERR {
	return m.err
}

func (m *MFR_M03v2) QRD() *QRD {
	return m.qrd
}

func (m *MFR_M03v2) QRF() *QRF {
	return m.qrf
}

func (m *MFR_M03v2) MFI() *MFI {
	return m.mfi
}

func (m *MFR_M03v2) GroupByMFE() []*MFR_M03_MFEv2 {
	return m.mfe
}

func (m *MFR_M03v2) DSC() *DSC {
	return m.dsc
}

func (m MFR_M03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"mfi": m.mfi,
		"mfe": m.mfe,
		"dsc": m.dsc,
	}, nil
}

type MFR_M03_MFEv2 struct {
	mfe *MFE // Required
}

func (m *MFR_M03_MFEv2) MFE() *MFE {
	return m.mfe
}

func (m MFR_M03_MFEv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"mfe": m.mfe,
	}, nil
}

type NMD_N01v2 struct {
	msh *MSH             // Required
	nck []*NMD_N01_NCKv2 // Required
}

func (m *NMD_N01v2) MSH() *MSH {
	return m.msh
}

func (m *NMD_N01v2) GroupByNCK() []*NMD_N01_NCKv2 {
	return m.nck
}

func (m NMD_N01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"nck": m.nck,
	}, nil
}

type NMD_N01_NCKv2 struct {
	nck *NCK // Required
	nte []*NTE
	nst *NMD_N01_NCK_NSTv2
	nsc *NMD_N01_NCK_NSCv2
}

func (m *NMD_N01_NCKv2) NCK() *NCK {
	return m.nck
}

func (m *NMD_N01_NCKv2) AllNTE() []*NTE {
	return m.nte
}

func (m *NMD_N01_NCKv2) GroupByNST() *NMD_N01_NCK_NSTv2 {
	return m.nst
}

func (m *NMD_N01_NCKv2) GroupByNSC() *NMD_N01_NCK_NSCv2 {
	return m.nsc
}

func (m NMD_N01_NCKv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"nck": m.nck,
		"nte": m.nte,
		"nst": m.nst,
		"nsc": m.nsc,
	}, nil
}

type NMD_N01_NCK_NSTv2 struct {
	nst *NST // Required
	nte []*NTE
}

func (m *NMD_N01_NCK_NSTv2) NST() *NST {
	return m.nst
}

func (m *NMD_N01_NCK_NSTv2) AllNTE() []*NTE {
	return m.nte
}

func (m NMD_N01_NCK_NSTv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"nst": m.nst,
		"nte": m.nte,
	}, nil
}

type NMD_N01_NCK_NSCv2 struct {
	nsc *NSC // Required
	nte []*NTE
}

func (m *NMD_N01_NCK_NSCv2) NSC() *NSC {
	return m.nsc
}

func (m *NMD_N01_NCK_NSCv2) AllNTE() []*NTE {
	return m.nte
}

func (m NMD_N01_NCK_NSCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"nsc": m.nsc,
		"nte": m.nte,
	}, nil
}

type NMQ_N02v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	nck []*NMQ_N02_NCKv2 // Required
}

func (m *NMQ_N02v2) MSH() *MSH {
	return m.msh
}

func (m *NMQ_N02v2) QRD() *QRD {
	return m.qrd
}

func (m *NMQ_N02v2) QRF() *QRF {
	return m.qrf
}

func (m *NMQ_N02v2) GroupByNCK() []*NMQ_N02_NCKv2 {
	return m.nck
}

func (m NMQ_N02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"nck": m.nck,
	}, nil
}

type NMQ_N02_NCKv2 struct {
	nck *NCK
	nst *NST
	nsc *NSC
}

func (m *NMQ_N02_NCKv2) NCK() *NCK {
	return m.nck
}

func (m *NMQ_N02_NCKv2) NST() *NST {
	return m.nst
}

func (m *NMQ_N02_NCKv2) NSC() *NSC {
	return m.nsc
}

func (m NMQ_N02_NCKv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"nck": m.nck,
		"nst": m.nst,
		"nsc": m.nsc,
	}, nil
}

type NMR_N02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	err *ERR
	qrd *QRD
	nck []*NMR_N02_NCKv2 // Required
}

func (m *NMR_N02v2) MSH() *MSH {
	return m.msh
}

func (m *NMR_N02v2) MSA() *MSA {
	return m.msa
}

func (m *NMR_N02v2) ERR() *ERR {
	return m.err
}

func (m *NMR_N02v2) QRD() *QRD {
	return m.qrd
}

func (m *NMR_N02v2) GroupByNCK() []*NMR_N02_NCKv2 {
	return m.nck
}

func (m NMR_N02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"err": m.err,
		"qrd": m.qrd,
		"nck": m.nck,
	}, nil
}

type NMR_N02_NCKv2 struct {
	nck  *NCK
	nte1 []*NTE
	nst  *NST
	nte2 []*NTE
	nsc  *NSC
	nte3 []*NTE
}

func (m *NMR_N02_NCKv2) NCK() *NCK {
	return m.nck
}

func (m *NMR_N02_NCKv2) AllNTE1() []*NTE {
	return m.nte1
}

func (m *NMR_N02_NCKv2) NST() *NST {
	return m.nst
}

func (m *NMR_N02_NCKv2) AllNTE2() []*NTE {
	return m.nte2
}

func (m *NMR_N02_NCKv2) NSC() *NSC {
	return m.nsc
}

func (m *NMR_N02_NCKv2) AllNTE3() []*NTE {
	return m.nte3
}

func (m NMR_N02_NCKv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"nck":  m.nck,
		"nte1": m.nte1,
		"nst":  m.nst,
		"nte2": m.nte2,
		"nsc":  m.nsc,
		"nte3": m.nte3,
	}, nil
}

type ORF_R04v2 struct {
	msh *MSH             // Required
	msa *MSA             // Required
	qrd []*ORF_R04_QRDv2 // Required
	obr []*ORF_R04_OBRv2 // Required
	dsc *DSC
}

func (m *ORF_R04v2) MSH() *MSH {
	return m.msh
}

func (m *ORF_R04v2) MSA() *MSA {
	return m.msa
}

func (m *ORF_R04v2) GroupByQRD() []*ORF_R04_QRDv2 {
	return m.qrd
}

func (m *ORF_R04v2) GroupByOBR() []*ORF_R04_OBRv2 {
	return m.obr
}

func (m *ORF_R04v2) DSC() *DSC {
	return m.dsc
}

func (m ORF_R04v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"qrd": m.qrd,
		"obr": m.obr,
		"dsc": m.dsc,
	}, nil
}

type ORF_R04_QRDv2 struct {
	qrd *QRD // Required
	qrf *QRF
	pid *PID
	nte []*NTE
}

func (m *ORF_R04_QRDv2) QRD() *QRD {
	return m.qrd
}

func (m *ORF_R04_QRDv2) QRF() *QRF {
	return m.qrf
}

func (m *ORF_R04_QRDv2) PID() *PID {
	return m.pid
}

func (m *ORF_R04_QRDv2) AllNTE() []*NTE {
	return m.nte
}

func (m ORF_R04_QRDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"qrd": m.qrd,
		"qrf": m.qrf,
		"pid": m.pid,
		"nte": m.nte,
	}, nil
}

type ORF_R04_OBRv2 struct {
	orc *ORC
	obr *OBR // Required
	nte []*NTE
	obx []*ORF_R04_OBR_OBXv2 // Required
}

func (m *ORF_R04_OBRv2) ORC() *ORC {
	return m.orc
}

func (m *ORF_R04_OBRv2) OBR() *OBR {
	return m.obr
}

func (m *ORF_R04_OBRv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORF_R04_OBRv2) GroupByOBX() []*ORF_R04_OBR_OBXv2 {
	return m.obx
}

func (m ORF_R04_OBRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"nte": m.nte,
		"obx": m.obx,
	}, nil
}

type ORF_R04_OBR_OBXv2 struct {
	obx *OBX
	nte []*NTE
}

func (m *ORF_R04_OBR_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *ORF_R04_OBR_OBXv2) AllNTE() []*NTE {
	return m.nte
}

func (m ORF_R04_OBR_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}

type ORM_O01v2 struct {
	msh *MSH // Required
	nte []*NTE
	pid *ORM_O01_PIDv2
	orc []*ORM_O01_ORCv2 // Required
}

func (m *ORM_O01v2) MSH() *MSH {
	return m.msh
}

func (m *ORM_O01v2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORM_O01v2) GroupByPID() *ORM_O01_PIDv2 {
	return m.pid
}

func (m *ORM_O01v2) GroupByORC() []*ORM_O01_ORCv2 {
	return m.orc
}

func (m ORM_O01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"nte": m.nte,
		"pid": m.pid,
		"orc": m.orc,
	}, nil
}

type ORM_O01_PIDv2 struct {
	pid *PID // Required
	nte []*NTE
	pv1 *PV1
}

func (m *ORM_O01_PIDv2) PID() *PID {
	return m.pid
}

func (m *ORM_O01_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORM_O01_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m ORM_O01_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"nte": m.nte,
		"pv1": m.pv1,
	}, nil
}

type ORM_O01_ORCv2 struct {
	orc  *ORC // Required
	obr  *OBR // Required
	rqd  *RQD // Required
	rq1  *RQ1 // Required
	rxo  *RXO // Required
	ods  *ODS // Required
	odt  *ODT // Required
	nte1 []*NTE
	obx  []*OBX
	nte2 []*NTE // Required
	blg  *BLG
}

func (m *ORM_O01_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *ORM_O01_ORCv2) OBR() *OBR {
	return m.obr
}

func (m *ORM_O01_ORCv2) RQD() *RQD {
	return m.rqd
}

func (m *ORM_O01_ORCv2) RQ1() *RQ1 {
	return m.rq1
}

func (m *ORM_O01_ORCv2) RXO() *RXO {
	return m.rxo
}

func (m *ORM_O01_ORCv2) ODS() *ODS {
	return m.ods
}

func (m *ORM_O01_ORCv2) ODT() *ODT {
	return m.odt
}

func (m *ORM_O01_ORCv2) AllNTE1() []*NTE {
	return m.nte1
}

func (m *ORM_O01_ORCv2) AllOBX() []*OBX {
	return m.obx
}

func (m *ORM_O01_ORCv2) AllNTE2() []*NTE {
	return m.nte2
}

func (m *ORM_O01_ORCv2) BLG() *BLG {
	return m.blg
}

func (m ORM_O01_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc":  m.orc,
		"obr":  m.obr,
		"rqd":  m.rqd,
		"rq1":  m.rq1,
		"rxo":  m.rxo,
		"ods":  m.ods,
		"odt":  m.odt,
		"nte1": m.nte1,
		"obx":  m.obx,
		"nte2": m.nte2,
		"blg":  m.blg,
	}, nil
}

type ORR_O02v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	nte []*NTE
	pid *ORR_O02_PIDv2
}

func (m *ORR_O02v2) MSH() *MSH {
	return m.msh
}

func (m *ORR_O02v2) MSA() *MSA {
	return m.msa
}

func (m *ORR_O02v2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORR_O02v2) GroupByPID() *ORR_O02_PIDv2 {
	return m.pid
}

func (m ORR_O02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"msa": m.msa,
		"nte": m.nte,
		"pid": m.pid,
	}, nil
}

type ORR_O02_PIDv2 struct {
	pid *PID
	nte []*NTE
	orc []*ORR_O02_PID_ORCv2 // Required
}

func (m *ORR_O02_PIDv2) PID() *PID {
	return m.pid
}

func (m *ORR_O02_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORR_O02_PIDv2) GroupByORC() []*ORR_O02_PID_ORCv2 {
	return m.orc
}

func (m ORR_O02_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"nte": m.nte,
		"orc": m.orc,
	}, nil
}

type ORR_O02_PID_ORCv2 struct {
	orc *ORC // Required
	obr *OBR // Required
	rqd *RQD // Required
	rq1 *RQ1 // Required
	rxo *RXO // Required
	ods *ODS // Required
	odt *ODT // Required
	nte []*NTE
}

func (m *ORR_O02_PID_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *ORR_O02_PID_ORCv2) OBR() *OBR {
	return m.obr
}

func (m *ORR_O02_PID_ORCv2) RQD() *RQD {
	return m.rqd
}

func (m *ORR_O02_PID_ORCv2) RQ1() *RQ1 {
	return m.rq1
}

func (m *ORR_O02_PID_ORCv2) RXO() *RXO {
	return m.rxo
}

func (m *ORR_O02_PID_ORCv2) ODS() *ODS {
	return m.ods
}

func (m *ORR_O02_PID_ORCv2) ODT() *ODT {
	return m.odt
}

func (m *ORR_O02_PID_ORCv2) AllNTE() []*NTE {
	return m.nte
}

func (m ORR_O02_PID_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"rqd": m.rqd,
		"rq1": m.rq1,
		"rxo": m.rxo,
		"ods": m.ods,
		"odt": m.odt,
		"nte": m.nte,
	}, nil
}

type ORU_R01v2 struct {
	msh *MSH             // Required
	pid []*ORU_R01_PIDv2 // Required
	dsc *DSC
}

func (m *ORU_R01v2) MSH() *MSH {
	return m.msh
}

func (m *ORU_R01v2) GroupByPID() []*ORU_R01_PIDv2 {
	return m.pid
}

func (m *ORU_R01v2) DSC() *DSC {
	return m.dsc
}

func (m ORU_R01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"pid": m.pid,
		"dsc": m.dsc,
	}, nil
}

type ORU_R01_PIDv2 struct {
	pid *PID // Required
	nte []*NTE
	pv1 *PV1
	obr []*ORU_R01_PID_OBRv2 // Required
}

func (m *ORU_R01_PIDv2) PID() *PID {
	return m.pid
}

func (m *ORU_R01_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORU_R01_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *ORU_R01_PIDv2) GroupByOBR() []*ORU_R01_PID_OBRv2 {
	return m.obr
}

func (m ORU_R01_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"nte": m.nte,
		"pv1": m.pv1,
		"obr": m.obr,
	}, nil
}

type ORU_R01_PID_OBRv2 struct {
	orc *ORC
	obr *OBR // Required
	nte []*NTE
	obx []*ORU_R01_PID_OBR_OBXv2 // Required
}

func (m *ORU_R01_PID_OBRv2) ORC() *ORC {
	return m.orc
}

func (m *ORU_R01_PID_OBRv2) OBR() *OBR {
	return m.obr
}

func (m *ORU_R01_PID_OBRv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORU_R01_PID_OBRv2) GroupByOBX() []*ORU_R01_PID_OBR_OBXv2 {
	return m.obx
}

func (m ORU_R01_PID_OBRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"nte": m.nte,
		"obx": m.obx,
	}, nil
}

type ORU_R01_PID_OBR_OBXv2 struct {
	obx *OBX
	nte []*NTE
}

func (m *ORU_R01_PID_OBR_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *ORU_R01_PID_OBR_OBXv2) AllNTE() []*NTE {
	return m.nte
}

func (m ORU_R01_PID_OBR_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}

type ORU_R03v2 struct {
	msh *MSH             // Required
	pid []*ORU_R03_PIDv2 // Required
	dsc *DSC
}

func (m *ORU_R03v2) MSH() *MSH {
	return m.msh
}

func (m *ORU_R03v2) GroupByPID() []*ORU_R03_PIDv2 {
	return m.pid
}

func (m *ORU_R03v2) DSC() *DSC {
	return m.dsc
}

func (m ORU_R03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"pid": m.pid,
		"dsc": m.dsc,
	}, nil
}

type ORU_R03_PIDv2 struct {
	pid *PID // Required
	nte []*NTE
	pv1 *PV1
	obr []*ORU_R03_PID_OBRv2 // Required
}

func (m *ORU_R03_PIDv2) PID() *PID {
	return m.pid
}

func (m *ORU_R03_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORU_R03_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *ORU_R03_PIDv2) GroupByOBR() []*ORU_R03_PID_OBRv2 {
	return m.obr
}

func (m ORU_R03_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"nte": m.nte,
		"pv1": m.pv1,
		"obr": m.obr,
	}, nil
}

type ORU_R03_PID_OBRv2 struct {
	orc *ORC
	obr *OBR // Required
	nte []*NTE
	obx []*ORU_R03_PID_OBR_OBXv2 // Required
}

func (m *ORU_R03_PID_OBRv2) ORC() *ORC {
	return m.orc
}

func (m *ORU_R03_PID_OBRv2) OBR() *OBR {
	return m.obr
}

func (m *ORU_R03_PID_OBRv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORU_R03_PID_OBRv2) GroupByOBX() []*ORU_R03_PID_OBR_OBXv2 {
	return m.obx
}

func (m ORU_R03_PID_OBRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"nte": m.nte,
		"obx": m.obx,
	}, nil
}

type ORU_R03_PID_OBR_OBXv2 struct {
	obx *OBX
	nte []*NTE
}

func (m *ORU_R03_PID_OBR_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *ORU_R03_PID_OBR_OBXv2) AllNTE() []*NTE {
	return m.nte
}

func (m ORU_R03_PID_OBR_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}

type ORU_R32v2 struct {
	msh *MSH             // Required
	pid []*ORU_R32_PIDv2 // Required
	dsc *DSC
}

func (m *ORU_R32v2) MSH() *MSH {
	return m.msh
}

func (m *ORU_R32v2) GroupByPID() []*ORU_R32_PIDv2 {
	return m.pid
}

func (m *ORU_R32v2) DSC() *DSC {
	return m.dsc
}

func (m ORU_R32v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"pid": m.pid,
		"dsc": m.dsc,
	}, nil
}

type ORU_R32_PIDv2 struct {
	pid *PID // Required
	pd1 *PD1
	nk1 []*NK1
	nte []*NTE
	pv1 *PV1 // Required
	pv2 *PV2
	obr []*ORU_R32_PID_OBRv2 // Required
}

func (m *ORU_R32_PIDv2) PID() *PID {
	return m.pid
}

func (m *ORU_R32_PIDv2) PD1() *PD1 {
	return m.pd1
}

func (m *ORU_R32_PIDv2) AllNK1() []*NK1 {
	return m.nk1
}

func (m *ORU_R32_PIDv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORU_R32_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m *ORU_R32_PIDv2) PV2() *PV2 {
	return m.pv2
}

func (m *ORU_R32_PIDv2) GroupByOBR() []*ORU_R32_PID_OBRv2 {
	return m.obr
}

func (m ORU_R32_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pd1": m.pd1,
		"nk1": m.nk1,
		"nte": m.nte,
		"pv1": m.pv1,
		"pv2": m.pv2,
		"obr": m.obr,
	}, nil
}

type ORU_R32_PID_OBRv2 struct {
	orc *ORC
	obr *OBR // Required
	nte []*NTE
	obx []*ORU_R32_PID_OBR_OBXv2 // Required
	cti []*CTI
}

func (m *ORU_R32_PID_OBRv2) ORC() *ORC {
	return m.orc
}

func (m *ORU_R32_PID_OBRv2) OBR() *OBR {
	return m.obr
}

func (m *ORU_R32_PID_OBRv2) AllNTE() []*NTE {
	return m.nte
}

func (m *ORU_R32_PID_OBRv2) GroupByOBX() []*ORU_R32_PID_OBR_OBXv2 {
	return m.obx
}

func (m *ORU_R32_PID_OBRv2) AllCTI() []*CTI {
	return m.cti
}

func (m ORU_R32_PID_OBRv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"nte": m.nte,
		"obx": m.obx,
		"cti": m.cti,
	}, nil
}

type ORU_R32_PID_OBR_OBXv2 struct {
	obx *OBX
	nte []*NTE
}

func (m *ORU_R32_PID_OBR_OBXv2) OBX() *OBX {
	return m.obx
}

func (m *ORU_R32_PID_OBR_OBXv2) AllNTE() []*NTE {
	return m.nte
}

func (m ORU_R32_PID_OBR_OBXv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"obx": m.obx,
		"nte": m.nte,
	}, nil
}

type QRY_A19v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
}

func (m *QRY_A19v2) MSH() *MSH {
	return m.msh
}

func (m *QRY_A19v2) QRD() *QRD {
	return m.qrd
}

func (m *QRY_A19v2) QRF() *QRF {
	return m.qrf
}

func (m QRY_A19v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
	}, nil
}

type QRY_P04v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	dsc *DSC
}

func (m *QRY_P04v2) MSH() *MSH {
	return m.msh
}

func (m *QRY_P04v2) QRD() *QRD {
	return m.qrd
}

func (m *QRY_P04v2) QRF() *QRF {
	return m.qrf
}

func (m *QRY_P04v2) DSC() *DSC {
	return m.dsc
}

func (m QRY_P04v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsc": m.dsc,
	}, nil
}

type QRY_Q01v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	dsc *DSC
}

func (m *QRY_Q01v2) MSH() *MSH {
	return m.msh
}

func (m *QRY_Q01v2) QRD() *QRD {
	return m.qrd
}

func (m *QRY_Q01v2) QRF() *QRF {
	return m.qrf
}

func (m *QRY_Q01v2) DSC() *DSC {
	return m.dsc
}

func (m QRY_Q01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsc": m.dsc,
	}, nil
}

type QRY_Q02v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	dsc *DSC
}

func (m *QRY_Q02v2) MSH() *MSH {
	return m.msh
}

func (m *QRY_Q02v2) QRD() *QRD {
	return m.qrd
}

func (m *QRY_Q02v2) QRF() *QRF {
	return m.qrf
}

func (m *QRY_Q02v2) DSC() *DSC {
	return m.dsc
}

func (m QRY_Q02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsc": m.dsc,
	}, nil
}

type QRY_R02v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF // Required
	dsc *DSC
}

func (m *QRY_R02v2) MSH() *MSH {
	return m.msh
}

func (m *QRY_R02v2) QRD() *QRD {
	return m.qrd
}

func (m *QRY_R02v2) QRF() *QRF {
	return m.qrf
}

func (m *QRY_R02v2) DSC() *DSC {
	return m.dsc
}

func (m QRY_R02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsc": m.dsc,
	}, nil
}

type UDM_Q05v2 struct {
	msh *MSH // Required
	urd *URD // Required
	urs *URS
	dsp []*DSP // Required
	dsc *DSC   // Required
}

func (m *UDM_Q05v2) MSH() *MSH {
	return m.msh
}

func (m *UDM_Q05v2) URD() *URD {
	return m.urd
}

func (m *UDM_Q05v2) URS() *URS {
	return m.urs
}

func (m *UDM_Q05v2) AllDSP() []*DSP {
	return m.dsp
}

func (m *UDM_Q05v2) DSC() *DSC {
	return m.dsc
}

func (m UDM_Q05v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"urd": m.urd,
		"urs": m.urs,
		"dsp": m.dsp,
		"dsc": m.dsc,
	}, nil
}

// Types maps the name of an HL7 segment or message type to the type of the struct that
// represents that segment or message type.
var Types = map[string]reflect.Type{}

func init() {
	Types["ACC"] = reflect.TypeOf(ACC{})
	Types["ADD"] = reflect.TypeOf(ADD{})
	Types["AL1"] = reflect.TypeOf(AL1{})
	Types["BHS"] = reflect.TypeOf(BHS{})
	Types["BLG"] = reflect.TypeOf(BLG{})
	Types["BTS"] = reflect.TypeOf(BTS{})
	Types["DG1"] = reflect.TypeOf(DG1{})
	Types["DSC"] = reflect.TypeOf(DSC{})
	Types["DSP"] = reflect.TypeOf(DSP{})
	Types["ERR"] = reflect.TypeOf(ERR{})
	Types["EVN"] = reflect.TypeOf(EVN{})
	Types["FHS"] = reflect.TypeOf(FHS{})
	Types["FT1"] = reflect.TypeOf(FT1{})
	Types["FTS"] = reflect.TypeOf(FTS{})
	Types["GT1"] = reflect.TypeOf(GT1{})
	Types["IN1"] = reflect.TypeOf(IN1{})
	Types["IN2"] = reflect.TypeOf(IN2{})
	Types["IN3"] = reflect.TypeOf(IN3{})
	Types["MFA"] = reflect.TypeOf(MFA{})
	Types["MFE"] = reflect.TypeOf(MFE{})
	Types["MFI"] = reflect.TypeOf(MFI{})
	Types["MRG"] = reflect.TypeOf(MRG{})
	Types["MSA"] = reflect.TypeOf(MSA{})
	Types["MSH"] = reflect.TypeOf(MSH{})
	Types["NCK"] = reflect.TypeOf(NCK{})
	Types["NK1"] = reflect.TypeOf(NK1{})
	Types["NPU"] = reflect.TypeOf(NPU{})
	Types["NSC"] = reflect.TypeOf(NSC{})
	Types["NST"] = reflect.TypeOf(NST{})
	Types["NTE"] = reflect.TypeOf(NTE{})
	Types["OBR"] = reflect.TypeOf(OBR{})
	Types["OBX"] = reflect.TypeOf(OBX{})
	Types["ODS"] = reflect.TypeOf(ODS{})
	Types["ODT"] = reflect.TypeOf(ODT{})
	Types["OM1"] = reflect.TypeOf(OM1{})
	Types["OM2"] = reflect.TypeOf(OM2{})
	Types["OM3"] = reflect.TypeOf(OM3{})
	Types["OM4"] = reflect.TypeOf(OM4{})
	Types["OM5"] = reflect.TypeOf(OM5{})
	Types["OM6"] = reflect.TypeOf(OM6{})
	Types["ORC"] = reflect.TypeOf(ORC{})
	Types["ORO"] = reflect.TypeOf(ORO{})
	Types["PID"] = reflect.TypeOf(PID{})
	Types["PR1"] = reflect.TypeOf(PR1{})
	Types["PRA"] = reflect.TypeOf(PRA{})
	Types["PV1"] = reflect.TypeOf(PV1{})
	Types["PV2"] = reflect.TypeOf(PV2{})
	Types["QRD"] = reflect.TypeOf(QRD{})
	Types["QRF"] = reflect.TypeOf(QRF{})
	Types["RQ1"] = reflect.TypeOf(RQ1{})
	Types["RQD"] = reflect.TypeOf(RQD{})
	Types["RX1"] = reflect.TypeOf(RX1{})
	Types["RXA"] = reflect.TypeOf(RXA{})
	Types["RXC"] = reflect.TypeOf(RXC{})
	Types["RXD"] = reflect.TypeOf(RXD{})
	Types["RXE"] = reflect.TypeOf(RXE{})
	Types["RXG"] = reflect.TypeOf(RXG{})
	Types["RXO"] = reflect.TypeOf(RXO{})
	Types["RXR"] = reflect.TypeOf(RXR{})
	Types["STF"] = reflect.TypeOf(STF{})
	Types["UB1"] = reflect.TypeOf(UB1{})
	Types["UB2"] = reflect.TypeOf(UB2{})
	Types["URD"] = reflect.TypeOf(URD{})
	Types["URS"] = reflect.TypeOf(URS{})
	Types["ACK"] = reflect.TypeOf(ACK{})
	Types["ADR_A19"] = reflect.TypeOf(ADR_A19{})
	Types["ADR_A19_INSURANCE"] = reflect.TypeOf(ADR_A19_INSURANCE{})
	Types["ADR_A19_QUERY_RESPONSE"] = reflect.TypeOf(ADR_A19_QUERY_RESPONSE{})
	Types["ADT_A01"] = reflect.TypeOf(ADT_A01{})
	Types["ADT_A01_INSURANCE"] = reflect.TypeOf(ADT_A01_INSURANCE{})
	Types["ADT_A02"] = reflect.TypeOf(ADT_A02{})
	Types["ADT_A03"] = reflect.TypeOf(ADT_A03{})
	Types["ADT_A04"] = reflect.TypeOf(ADT_A04{})
	Types["ADT_A04_INSURANCE"] = reflect.TypeOf(ADT_A04_INSURANCE{})
	Types["ADT_A05"] = reflect.TypeOf(ADT_A05{})
	Types["ADT_A05_INSURANCE"] = reflect.TypeOf(ADT_A05_INSURANCE{})
	Types["ADT_A06"] = reflect.TypeOf(ADT_A06{})
	Types["ADT_A06_INSURANCE"] = reflect.TypeOf(ADT_A06_INSURANCE{})
	Types["ADT_A07"] = reflect.TypeOf(ADT_A07{})
	Types["ADT_A07_INSURANCE"] = reflect.TypeOf(ADT_A07_INSURANCE{})
	Types["ADT_A08"] = reflect.TypeOf(ADT_A08{})
	Types["ADT_A08_INSURANCE"] = reflect.TypeOf(ADT_A08_INSURANCE{})
	Types["ADT_A09"] = reflect.TypeOf(ADT_A09{})
	Types["ADT_A10"] = reflect.TypeOf(ADT_A10{})
	Types["ADT_A11"] = reflect.TypeOf(ADT_A11{})
	Types["ADT_A12"] = reflect.TypeOf(ADT_A12{})
	Types["ADT_A13"] = reflect.TypeOf(ADT_A13{})
	Types["ADT_A13_INSURANCE"] = reflect.TypeOf(ADT_A13_INSURANCE{})
	Types["ADT_A14"] = reflect.TypeOf(ADT_A14{})
	Types["ADT_A14_INSURANCE"] = reflect.TypeOf(ADT_A14_INSURANCE{})
	Types["ADT_A15"] = reflect.TypeOf(ADT_A15{})
	Types["ADT_A16"] = reflect.TypeOf(ADT_A16{})
	Types["ADT_A17"] = reflect.TypeOf(ADT_A17{})
	Types["ADT_A17_PATIENT"] = reflect.TypeOf(ADT_A17_PATIENT{})
	Types["ADT_A18"] = reflect.TypeOf(ADT_A18{})
	Types["ADT_A20"] = reflect.TypeOf(ADT_A20{})
	Types["ADT_A21"] = reflect.TypeOf(ADT_A21{})
	Types["ADT_A22"] = reflect.TypeOf(ADT_A22{})
	Types["ADT_A23"] = reflect.TypeOf(ADT_A23{})
	Types["ADT_A24"] = reflect.TypeOf(ADT_A24{})
	Types["ADT_A25"] = reflect.TypeOf(ADT_A25{})
	Types["ADT_A26"] = reflect.TypeOf(ADT_A26{})
	Types["ADT_A27"] = reflect.TypeOf(ADT_A27{})
	Types["ADT_A28"] = reflect.TypeOf(ADT_A28{})
	Types["ADT_A28_INSURANCE"] = reflect.TypeOf(ADT_A28_INSURANCE{})
	Types["ADT_A29"] = reflect.TypeOf(ADT_A29{})
	Types["ADT_A30"] = reflect.TypeOf(ADT_A30{})
	Types["ADT_A31"] = reflect.TypeOf(ADT_A31{})
	Types["ADT_A31_INSURANCE"] = reflect.TypeOf(ADT_A31_INSURANCE{})
	Types["ADT_A32"] = reflect.TypeOf(ADT_A32{})
	Types["ADT_A33"] = reflect.TypeOf(ADT_A33{})
	Types["ADT_A34"] = reflect.TypeOf(ADT_A34{})
	Types["ADT_A35"] = reflect.TypeOf(ADT_A35{})
	Types["ADT_A36"] = reflect.TypeOf(ADT_A36{})
	Types["ADT_A37"] = reflect.TypeOf(ADT_A37{})
	Types["ADT_A40"] = reflect.TypeOf(ADT_A40{})
	Types["ADT_A40_PATIENT"] = reflect.TypeOf(ADT_A40_PATIENT{})
	Types["ADT_A44"] = reflect.TypeOf(ADT_A44{})
	Types["ADT_A44_PATIENT"] = reflect.TypeOf(ADT_A44_PATIENT{})
	Types["BAR_P01"] = reflect.TypeOf(BAR_P01{})
	Types["BAR_P01_INSURANCE"] = reflect.TypeOf(BAR_P01_INSURANCE{})
	Types["BAR_P01_VISIT"] = reflect.TypeOf(BAR_P01_VISIT{})
	Types["BAR_P02"] = reflect.TypeOf(BAR_P02{})
	Types["BAR_P02_PATIENT"] = reflect.TypeOf(BAR_P02_PATIENT{})
	Types["DFT_P03"] = reflect.TypeOf(DFT_P03{})
	Types["DSR_P04"] = reflect.TypeOf(DSR_P04{})
	Types["DSR_Q01"] = reflect.TypeOf(DSR_Q01{})
	Types["DSR_Q03"] = reflect.TypeOf(DSR_Q03{})
	Types["DSR_R03"] = reflect.TypeOf(DSR_R03{})
	Types["MCF_Q02"] = reflect.TypeOf(MCF_Q02{})
	Types["MFD_M01"] = reflect.TypeOf(MFD_M01{})
	Types["MFD_M02"] = reflect.TypeOf(MFD_M02{})
	Types["MFD_M03"] = reflect.TypeOf(MFD_M03{})
	Types["MFK_M01"] = reflect.TypeOf(MFK_M01{})
	Types["MFK_M02"] = reflect.TypeOf(MFK_M02{})
	Types["MFK_M03"] = reflect.TypeOf(MFK_M03{})
	Types["MFN_M01"] = reflect.TypeOf(MFN_M01{})
	Types["MFN_M01_MF"] = reflect.TypeOf(MFN_M01_MF{})
	Types["MFN_M02"] = reflect.TypeOf(MFN_M02{})
	Types["MFN_M02_MF_STAFF"] = reflect.TypeOf(MFN_M02_MF_STAFF{})
	Types["MFN_M03"] = reflect.TypeOf(MFN_M03{})
	Types["MFN_M03_MF_TEST"] = reflect.TypeOf(MFN_M03_MF_TEST{})
	Types["MFQ_M01"] = reflect.TypeOf(MFQ_M01{})
	Types["MFQ_M02"] = reflect.TypeOf(MFQ_M02{})
	Types["MFQ_M03"] = reflect.TypeOf(MFQ_M03{})
	Types["MFR_M01"] = reflect.TypeOf(MFR_M01{})
	Types["MFR_M01_MF"] = reflect.TypeOf(MFR_M01_MF{})
	Types["MFR_M02"] = reflect.TypeOf(MFR_M02{})
	Types["MFR_M02_MF_STAFF"] = reflect.TypeOf(MFR_M02_MF_STAFF{})
	Types["MFR_M03"] = reflect.TypeOf(MFR_M03{})
	Types["MFR_M03_MF_TEST"] = reflect.TypeOf(MFR_M03_MF_TEST{})
	Types["NMD_N01_APP_STATS"] = reflect.TypeOf(NMD_N01_APP_STATS{})
	Types["NMD_N01_APP_STATUS"] = reflect.TypeOf(NMD_N01_APP_STATUS{})
	Types["NMD_N01_CLOCK"] = reflect.TypeOf(NMD_N01_CLOCK{})
	Types["NMD_N01_CLOCK_AND_STATS_WITH_NOTES"] = reflect.TypeOf(NMD_N01_CLOCK_AND_STATS_WITH_NOTES{})
	Types["NMD_N01"] = reflect.TypeOf(NMD_N01{})
	Types["NMQ_N02_CLOCK_AND_STATISTICS"] = reflect.TypeOf(NMQ_N02_CLOCK_AND_STATISTICS{})
	Types["NMQ_N02"] = reflect.TypeOf(NMQ_N02{})
	Types["NMQ_N02_QRY_WITH_DETAIL"] = reflect.TypeOf(NMQ_N02_QRY_WITH_DETAIL{})
	Types["NMR_N02_CLOCK_AND_STATS_WITH_NOTES_ALT"] = reflect.TypeOf(NMR_N02_CLOCK_AND_STATS_WITH_NOTES_ALT{})
	Types["NMR_N02"] = reflect.TypeOf(NMR_N02{})
	Types["ORF_R04"] = reflect.TypeOf(ORF_R04{})
	Types["ORF_R04_OBSERVATION"] = reflect.TypeOf(ORF_R04_OBSERVATION{})
	Types["ORF_R04_ORDER"] = reflect.TypeOf(ORF_R04_ORDER{})
	Types["ORF_R04_QUERY_RESPONSE"] = reflect.TypeOf(ORF_R04_QUERY_RESPONSE{})
	Types["ORM_O01_CHOICE"] = reflect.TypeOf(ORM_O01_CHOICE{})
	Types["ORM_O01"] = reflect.TypeOf(ORM_O01{})
	Types["ORM_O01_ORDER"] = reflect.TypeOf(ORM_O01_ORDER{})
	Types["ORM_O01_ORDER_DETAIL"] = reflect.TypeOf(ORM_O01_ORDER_DETAIL{})
	Types["ORM_O01_PATIENT"] = reflect.TypeOf(ORM_O01_PATIENT{})
	Types["ORR_O02_CHOICE"] = reflect.TypeOf(ORR_O02_CHOICE{})
	Types["ORR_O02"] = reflect.TypeOf(ORR_O02{})
	Types["ORR_O02_ORDER"] = reflect.TypeOf(ORR_O02_ORDER{})
	Types["ORR_O02_ORDER_DETAIL"] = reflect.TypeOf(ORR_O02_ORDER_DETAIL{})
	Types["ORR_O02_PATIENT"] = reflect.TypeOf(ORR_O02_PATIENT{})
	Types["ORU_R01"] = reflect.TypeOf(ORU_R01{})
	Types["ORU_R01_OBSERVATION"] = reflect.TypeOf(ORU_R01_OBSERVATION{})
	Types["ORU_R01_ORDER_OBSERVATION"] = reflect.TypeOf(ORU_R01_ORDER_OBSERVATION{})
	Types["ORU_R01_PATIENT"] = reflect.TypeOf(ORU_R01_PATIENT{})
	Types["ORU_R01_PATIENT_RESULT"] = reflect.TypeOf(ORU_R01_PATIENT_RESULT{})
	Types["ORU_R03"] = reflect.TypeOf(ORU_R03{})
	Types["ORU_R03_OBSERVATION"] = reflect.TypeOf(ORU_R03_OBSERVATION{})
	Types["ORU_R03_ORDER_OBSERVATION"] = reflect.TypeOf(ORU_R03_ORDER_OBSERVATION{})
	Types["ORU_R03_PATIENT"] = reflect.TypeOf(ORU_R03_PATIENT{})
	Types["ORU_R03_PATIENT_RESULT"] = reflect.TypeOf(ORU_R03_PATIENT_RESULT{})
	Types["ORU_R32"] = reflect.TypeOf(ORU_R32{})
	Types["ORU_R32_OBSERVATION"] = reflect.TypeOf(ORU_R32_OBSERVATION{})
	Types["ORU_R32_ORDER_OBSERVATION"] = reflect.TypeOf(ORU_R32_ORDER_OBSERVATION{})
	Types["ORU_R32_PATIENT"] = reflect.TypeOf(ORU_R32_PATIENT{})
	Types["ORU_R32_PATIENT_RESULT"] = reflect.TypeOf(ORU_R32_PATIENT_RESULT{})
	Types["ORU_R32_VISIT"] = reflect.TypeOf(ORU_R32_VISIT{})
	Types["QRY_A19"] = reflect.TypeOf(QRY_A19{})
	Types["QRY_P04"] = reflect.TypeOf(QRY_P04{})
	Types["QRY_Q01"] = reflect.TypeOf(QRY_Q01{})
	Types["QRY_Q02"] = reflect.TypeOf(QRY_Q02{})
	Types["QRY_R02"] = reflect.TypeOf(QRY_R02{})
	Types["UDM_Q05"] = reflect.TypeOf(UDM_Q05{})
	Types["GenericHL7Segment"] = reflect.TypeOf(GenericHL7Segment{})
	Types["ACKv2"] = reflect.TypeOf(ACKv2{})
	Types["ADR_A19v2"] = reflect.TypeOf(ADR_A19v2{})
	Types["ADR_A19_PIDv2"] = reflect.TypeOf(ADR_A19_PIDv2{})
	Types["ADR_A19_PID_IN1v2"] = reflect.TypeOf(ADR_A19_PID_IN1v2{})
	Types["ADT_A01v2"] = reflect.TypeOf(ADT_A01v2{})
	Types["ADT_A01_IN1v2"] = reflect.TypeOf(ADT_A01_IN1v2{})
	Types["ADT_A02v2"] = reflect.TypeOf(ADT_A02v2{})
	Types["ADT_A03v2"] = reflect.TypeOf(ADT_A03v2{})
	Types["ADT_A04v2"] = reflect.TypeOf(ADT_A04v2{})
	Types["ADT_A04_IN1v2"] = reflect.TypeOf(ADT_A04_IN1v2{})
	Types["ADT_A05v2"] = reflect.TypeOf(ADT_A05v2{})
	Types["ADT_A05_IN1v2"] = reflect.TypeOf(ADT_A05_IN1v2{})
	Types["ADT_A06v2"] = reflect.TypeOf(ADT_A06v2{})
	Types["ADT_A06_IN1v2"] = reflect.TypeOf(ADT_A06_IN1v2{})
	Types["ADT_A07v2"] = reflect.TypeOf(ADT_A07v2{})
	Types["ADT_A07_IN1v2"] = reflect.TypeOf(ADT_A07_IN1v2{})
	Types["ADT_A08v2"] = reflect.TypeOf(ADT_A08v2{})
	Types["ADT_A08_IN1v2"] = reflect.TypeOf(ADT_A08_IN1v2{})
	Types["ADT_A09v2"] = reflect.TypeOf(ADT_A09v2{})
	Types["ADT_A10v2"] = reflect.TypeOf(ADT_A10v2{})
	Types["ADT_A11v2"] = reflect.TypeOf(ADT_A11v2{})
	Types["ADT_A12v2"] = reflect.TypeOf(ADT_A12v2{})
	Types["ADT_A13v2"] = reflect.TypeOf(ADT_A13v2{})
	Types["ADT_A13_IN1v2"] = reflect.TypeOf(ADT_A13_IN1v2{})
	Types["ADT_A14v2"] = reflect.TypeOf(ADT_A14v2{})
	Types["ADT_A14_IN1v2"] = reflect.TypeOf(ADT_A14_IN1v2{})
	Types["ADT_A15v2"] = reflect.TypeOf(ADT_A15v2{})
	Types["ADT_A16v2"] = reflect.TypeOf(ADT_A16v2{})
	Types["ADT_A17v2"] = reflect.TypeOf(ADT_A17v2{})
	Types["ADT_A18v2"] = reflect.TypeOf(ADT_A18v2{})
	Types["ADT_A20v2"] = reflect.TypeOf(ADT_A20v2{})
	Types["ADT_A21v2"] = reflect.TypeOf(ADT_A21v2{})
	Types["ADT_A22v2"] = reflect.TypeOf(ADT_A22v2{})
	Types["ADT_A23v2"] = reflect.TypeOf(ADT_A23v2{})
	Types["ADT_A24v2"] = reflect.TypeOf(ADT_A24v2{})
	Types["ADT_A25v2"] = reflect.TypeOf(ADT_A25v2{})
	Types["ADT_A26v2"] = reflect.TypeOf(ADT_A26v2{})
	Types["ADT_A27v2"] = reflect.TypeOf(ADT_A27v2{})
	Types["ADT_A28v2"] = reflect.TypeOf(ADT_A28v2{})
	Types["ADT_A28_IN1v2"] = reflect.TypeOf(ADT_A28_IN1v2{})
	Types["ADT_A29v2"] = reflect.TypeOf(ADT_A29v2{})
	Types["ADT_A30v2"] = reflect.TypeOf(ADT_A30v2{})
	Types["ADT_A31v2"] = reflect.TypeOf(ADT_A31v2{})
	Types["ADT_A31_IN1v2"] = reflect.TypeOf(ADT_A31_IN1v2{})
	Types["ADT_A32v2"] = reflect.TypeOf(ADT_A32v2{})
	Types["ADT_A33v2"] = reflect.TypeOf(ADT_A33v2{})
	Types["ADT_A34v2"] = reflect.TypeOf(ADT_A34v2{})
	Types["ADT_A35v2"] = reflect.TypeOf(ADT_A35v2{})
	Types["ADT_A36v2"] = reflect.TypeOf(ADT_A36v2{})
	Types["ADT_A37v2"] = reflect.TypeOf(ADT_A37v2{})
	Types["ADT_A40v2"] = reflect.TypeOf(ADT_A40v2{})
	Types["ADT_A40_PIDv2"] = reflect.TypeOf(ADT_A40_PIDv2{})
	Types["ADT_A44v2"] = reflect.TypeOf(ADT_A44v2{})
	Types["ADT_A44_PIDv2"] = reflect.TypeOf(ADT_A44_PIDv2{})
	Types["BAR_P01v2"] = reflect.TypeOf(BAR_P01v2{})
	Types["BAR_P01_PV1v2"] = reflect.TypeOf(BAR_P01_PV1v2{})
	Types["BAR_P01_PV1_IN1v2"] = reflect.TypeOf(BAR_P01_PV1_IN1v2{})
	Types["BAR_P02v2"] = reflect.TypeOf(BAR_P02v2{})
	Types["BAR_P02_PIDv2"] = reflect.TypeOf(BAR_P02_PIDv2{})
	Types["DFT_P03v2"] = reflect.TypeOf(DFT_P03v2{})
	Types["DSR_P04v2"] = reflect.TypeOf(DSR_P04v2{})
	Types["DSR_Q01v2"] = reflect.TypeOf(DSR_Q01v2{})
	Types["DSR_Q03v2"] = reflect.TypeOf(DSR_Q03v2{})
	Types["DSR_R03v2"] = reflect.TypeOf(DSR_R03v2{})
	Types["MCF_Q02v2"] = reflect.TypeOf(MCF_Q02v2{})
	Types["MFD_M01v2"] = reflect.TypeOf(MFD_M01v2{})
	Types["MFD_M02v2"] = reflect.TypeOf(MFD_M02v2{})
	Types["MFD_M03v2"] = reflect.TypeOf(MFD_M03v2{})
	Types["MFK_M01v2"] = reflect.TypeOf(MFK_M01v2{})
	Types["MFK_M02v2"] = reflect.TypeOf(MFK_M02v2{})
	Types["MFK_M03v2"] = reflect.TypeOf(MFK_M03v2{})
	Types["MFN_M01v2"] = reflect.TypeOf(MFN_M01v2{})
	Types["MFN_M01_MFEv2"] = reflect.TypeOf(MFN_M01_MFEv2{})
	Types["MFN_M02v2"] = reflect.TypeOf(MFN_M02v2{})
	Types["MFN_M02_MFEv2"] = reflect.TypeOf(MFN_M02_MFEv2{})
	Types["MFN_M03v2"] = reflect.TypeOf(MFN_M03v2{})
	Types["MFN_M03_MFEv2"] = reflect.TypeOf(MFN_M03_MFEv2{})
	Types["MFQ_M01v2"] = reflect.TypeOf(MFQ_M01v2{})
	Types["MFQ_M02v2"] = reflect.TypeOf(MFQ_M02v2{})
	Types["MFQ_M03v2"] = reflect.TypeOf(MFQ_M03v2{})
	Types["MFR_M01v2"] = reflect.TypeOf(MFR_M01v2{})
	Types["MFR_M01_MFEv2"] = reflect.TypeOf(MFR_M01_MFEv2{})
	Types["MFR_M02v2"] = reflect.TypeOf(MFR_M02v2{})
	Types["MFR_M02_MFEv2"] = reflect.TypeOf(MFR_M02_MFEv2{})
	Types["MFR_M03v2"] = reflect.TypeOf(MFR_M03v2{})
	Types["MFR_M03_MFEv2"] = reflect.TypeOf(MFR_M03_MFEv2{})
	Types["NMD_N01v2"] = reflect.TypeOf(NMD_N01v2{})
	Types["NMD_N01_NCKv2"] = reflect.TypeOf(NMD_N01_NCKv2{})
	Types["NMD_N01_NCK_NSTv2"] = reflect.TypeOf(NMD_N01_NCK_NSTv2{})
	Types["NMD_N01_NCK_NSCv2"] = reflect.TypeOf(NMD_N01_NCK_NSCv2{})
	Types["NMQ_N02v2"] = reflect.TypeOf(NMQ_N02v2{})
	Types["NMQ_N02_NCKv2"] = reflect.TypeOf(NMQ_N02_NCKv2{})
	Types["NMR_N02v2"] = reflect.TypeOf(NMR_N02v2{})
	Types["NMR_N02_NCKv2"] = reflect.TypeOf(NMR_N02_NCKv2{})
	Types["ORF_R04v2"] = reflect.TypeOf(ORF_R04v2{})
	Types["ORF_R04_QRDv2"] = reflect.TypeOf(ORF_R04_QRDv2{})
	Types["ORF_R04_OBRv2"] = reflect.TypeOf(ORF_R04_OBRv2{})
	Types["ORF_R04_OBR_OBXv2"] = reflect.TypeOf(ORF_R04_OBR_OBXv2{})
	Types["ORM_O01v2"] = reflect.TypeOf(ORM_O01v2{})
	Types["ORM_O01_PIDv2"] = reflect.TypeOf(ORM_O01_PIDv2{})
	Types["ORM_O01_ORCv2"] = reflect.TypeOf(ORM_O01_ORCv2{})
	Types["ORR_O02v2"] = reflect.TypeOf(ORR_O02v2{})
	Types["ORR_O02_PIDv2"] = reflect.TypeOf(ORR_O02_PIDv2{})
	Types["ORR_O02_PID_ORCv2"] = reflect.TypeOf(ORR_O02_PID_ORCv2{})
	Types["ORU_R01v2"] = reflect.TypeOf(ORU_R01v2{})
	Types["ORU_R01_PIDv2"] = reflect.TypeOf(ORU_R01_PIDv2{})
	Types["ORU_R01_PID_OBRv2"] = reflect.TypeOf(ORU_R01_PID_OBRv2{})
	Types["ORU_R01_PID_OBR_OBXv2"] = reflect.TypeOf(ORU_R01_PID_OBR_OBXv2{})
	Types["ORU_R03v2"] = reflect.TypeOf(ORU_R03v2{})
	Types["ORU_R03_PIDv2"] = reflect.TypeOf(ORU_R03_PIDv2{})
	Types["ORU_R03_PID_OBRv2"] = reflect.TypeOf(ORU_R03_PID_OBRv2{})
	Types["ORU_R03_PID_OBR_OBXv2"] = reflect.TypeOf(ORU_R03_PID_OBR_OBXv2{})
	Types["ORU_R32v2"] = reflect.TypeOf(ORU_R32v2{})
	Types["ORU_R32_PIDv2"] = reflect.TypeOf(ORU_R32_PIDv2{})
	Types["ORU_R32_PID_OBRv2"] = reflect.TypeOf(ORU_R32_PID_OBRv2{})
	Types["ORU_R32_PID_OBR_OBXv2"] = reflect.TypeOf(ORU_R32_PID_OBR_OBXv2{})
	Types["QRY_A19v2"] = reflect.TypeOf(QRY_A19v2{})
	Types["QRY_P04v2"] = reflect.TypeOf(QRY_P04v2{})
	Types["QRY_Q01v2"] = reflect.TypeOf(QRY_Q01v2{})
	Types["QRY_Q02v2"] = reflect.TypeOf(QRY_Q02v2{})
	Types["QRY_R02v2"] = reflect.TypeOf(QRY_R02v2{})
	Types["UDM_Q05v2"] = reflect.TypeOf(UDM_Q05v2{})
}

var FollowSets = map[string]StringSet{
	"ACKv2.err": StringSet{},
	"ACKv2.msa": StringSet{
		"ERR": true,
	},
	"ACKv2.msh": StringSet{
		"ERR": true,
		"MSA": true,
	},
	"ADR_A19_PID_IN1v2.in1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PID_IN1v2.in2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PID_IN1v2.in3": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.acc": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.dg1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.gt1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.pr1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.ub1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19_PIDv2.ub2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19v2.dsc": StringSet{},
	"ADR_A19v2.err": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19v2.msa": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"ERR": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"ERR": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MSA": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"QRD": true,
		"UB1": true,
		"UB2": true,
	},
	"ADR_A19v2.qrd": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"DSC": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A01v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A01v2.ub2": StringSet{},
	"ADT_A02v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A02v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A02v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A02v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A02v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A02v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A03v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A03v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A03v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A03v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A03v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A03v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A04_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A04v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A04v2.ub2": StringSet{},
	"ADT_A05_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A05v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A05v2.ub2": StringSet{},
	"ADT_A06_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MRG": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.mrg": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MRG": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MRG": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A06v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A06v2.ub2": StringSet{},
	"ADT_A07_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MRG": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.mrg": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MRG": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"MRG": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A07v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A07v2.ub2": StringSet{},
	"ADT_A08_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A08v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A08v2.ub2": StringSet{},
	"ADT_A09v2.dg1": StringSet{
		"DG1": true,
	},
	"ADT_A09v2.evn": StringSet{
		"DG1": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A09v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A09v2.obx": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A09v2.pid": StringSet{
		"DG1": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A09v2.pv1": StringSet{
		"DG1": true,
		"OBX": true,
		"PV2": true,
	},
	"ADT_A09v2.pv2": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A10v2.dg1": StringSet{
		"DG1": true,
	},
	"ADT_A10v2.evn": StringSet{
		"DG1": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A10v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A10v2.obx": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A10v2.pid": StringSet{
		"DG1": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A10v2.pv1": StringSet{
		"DG1": true,
		"OBX": true,
		"PV2": true,
	},
	"ADT_A10v2.pv2": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A11v2.dg1": StringSet{
		"DG1": true,
	},
	"ADT_A11v2.evn": StringSet{
		"DG1": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A11v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A11v2.obx": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A11v2.pid": StringSet{
		"DG1": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A11v2.pv1": StringSet{
		"DG1": true,
		"OBX": true,
		"PV2": true,
	},
	"ADT_A11v2.pv2": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A12v2.dg1": StringSet{
		"DG1": true,
	},
	"ADT_A12v2.evn": StringSet{
		"DG1": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A12v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A12v2.obx": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A12v2.pid": StringSet{
		"DG1": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A12v2.pv1": StringSet{
		"DG1": true,
		"OBX": true,
		"PV2": true,
	},
	"ADT_A12v2.pv2": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A13_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A13v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A13v2.ub2": StringSet{},
	"ADT_A14_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A14v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A14v2.ub2": StringSet{},
	"ADT_A15v2.dg1": StringSet{
		"DG1": true,
	},
	"ADT_A15v2.evn": StringSet{
		"DG1": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A15v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A15v2.obx": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A15v2.pid": StringSet{
		"DG1": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A15v2.pv1": StringSet{
		"DG1": true,
		"OBX": true,
		"PV2": true,
	},
	"ADT_A15v2.pv2": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A16v2.dg1": StringSet{
		"DG1": true,
	},
	"ADT_A16v2.evn": StringSet{
		"DG1": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A16v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A16v2.obx": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A16v2.pid": StringSet{
		"DG1": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A16v2.pv1": StringSet{
		"DG1": true,
		"OBX": true,
		"PV2": true,
	},
	"ADT_A16v2.pv2": StringSet{
		"DG1": true,
		"OBX": true,
	},
	"ADT_A17v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A17v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A17v2.obx1": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A17v2.obx2": StringSet{
		"OBX": true,
	},
	"ADT_A17v2.pid1": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A17v2.pid2": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A17v2.pv11": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A17v2.pv12": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A17v2.pv21": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A17v2.pv22": StringSet{
		"OBX": true,
	},
	"ADT_A18v2.evn": StringSet{
		"MRG": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A18v2.mrg": StringSet{
		"PV1": true,
	},
	"ADT_A18v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A18v2.pid": StringSet{
		"MRG": true,
		"PV1": true,
	},
	"ADT_A18v2.pv1": StringSet{},
	"ADT_A20v2.evn": StringSet{
		"NPU": true,
	},
	"ADT_A20v2.msh": StringSet{
		"EVN": true,
		"NPU": true,
	},
	"ADT_A20v2.npu": StringSet{},
	"ADT_A21v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A21v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A21v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A21v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A21v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A21v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A22v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A22v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A22v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A22v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A22v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A22v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A23v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A23v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A23v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A23v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A23v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A23v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A24v2.evn": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A24v2.msh": StringSet{
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A24v2.pid1": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A24v2.pid2": StringSet{
		"PV1": true,
	},
	"ADT_A24v2.pv11": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A24v2.pv12": StringSet{},
	"ADT_A25v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A25v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A25v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A25v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A25v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A25v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A26v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A26v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A26v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A26v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A26v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A26v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A27v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A27v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A27v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A27v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A27v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A27v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A28_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A28v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A28v2.ub2": StringSet{},
	"ADT_A29v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A29v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A29v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A29v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A29v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A29v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A30v2.evn": StringSet{
		"MRG": true,
		"PID": true,
	},
	"ADT_A30v2.mrg": StringSet{},
	"ADT_A30v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PID": true,
	},
	"ADT_A30v2.pid": StringSet{
		"MRG": true,
	},
	"ADT_A31_IN1v2.in1": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31_IN1v2.in2": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31_IN1v2.in3": StringSet{
		"ACC": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.acc": StringSet{
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.gt1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.pr1": StringSet{
		"ACC": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"OBX": true,
		"PR1": true,
		"UB1": true,
		"UB2": true,
	},
	"ADT_A31v2.ub1": StringSet{
		"UB2": true,
	},
	"ADT_A31v2.ub2": StringSet{},
	"ADT_A32v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A32v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A32v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A32v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A32v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A32v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A33v2.evn": StringSet{
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A33v2.msh": StringSet{
		"EVN": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A33v2.obx": StringSet{
		"OBX": true,
	},
	"ADT_A33v2.pid": StringSet{
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"ADT_A33v2.pv1": StringSet{
		"OBX": true,
		"PV2": true,
	},
	"ADT_A33v2.pv2": StringSet{
		"OBX": true,
	},
	"ADT_A34v2.evn": StringSet{
		"MRG": true,
		"PID": true,
	},
	"ADT_A34v2.mrg": StringSet{},
	"ADT_A34v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PID": true,
	},
	"ADT_A34v2.pid": StringSet{
		"MRG": true,
	},
	"ADT_A35v2.evn": StringSet{
		"MRG": true,
		"PID": true,
	},
	"ADT_A35v2.mrg": StringSet{},
	"ADT_A35v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PID": true,
	},
	"ADT_A35v2.pid": StringSet{
		"MRG": true,
	},
	"ADT_A36v2.evn": StringSet{
		"MRG": true,
		"PID": true,
	},
	"ADT_A36v2.mrg": StringSet{},
	"ADT_A36v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PID": true,
	},
	"ADT_A36v2.pid": StringSet{
		"MRG": true,
	},
	"ADT_A37v2.evn": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A37v2.msh": StringSet{
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A37v2.pid1": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A37v2.pid2": StringSet{
		"PV1": true,
	},
	"ADT_A37v2.pv11": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A37v2.pv12": StringSet{},
	"ADT_A40_PIDv2.mrg": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A40_PIDv2.pd1": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A40_PIDv2.pid": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A40_PIDv2.pv1": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A40v2.evn": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A40v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A44_PIDv2.mrg": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
	},
	"ADT_A44_PIDv2.pd1": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
	},
	"ADT_A44_PIDv2.pid": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
	},
	"ADT_A44v2.evn": StringSet{
		"MRG": true,
		"PD1": true,
		"PID": true,
	},
	"ADT_A44v2.msh": StringSet{
		"EVN": true,
		"MRG": true,
		"PD1": true,
		"PID": true,
	},
	"BAR_P01_PV1_IN1v2.in1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1_IN1v2.in2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1_IN1v2.in3": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.acc": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.al1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.dg1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.gt1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.nk1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.obx": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.pr1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.pv1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.pv2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.ub1": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01_PV1v2.ub2": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01v2.evn": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01v2.msh": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P01v2.pid": StringSet{
		"ACC": true,
		"AL1": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"IN2": true,
		"IN3": true,
		"NK1": true,
		"OBX": true,
		"PR1": true,
		"PV1": true,
		"PV2": true,
		"UB1": true,
		"UB2": true,
	},
	"BAR_P02_PIDv2.pid": StringSet{
		"PID": true,
		"PV1": true,
	},
	"BAR_P02_PIDv2.pv1": StringSet{
		"PID": true,
		"PV1": true,
	},
	"BAR_P02v2.evn": StringSet{
		"PID": true,
		"PV1": true,
	},
	"BAR_P02v2.msh": StringSet{
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"DFT_P03v2.evn": StringSet{
		"FT1": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"DFT_P03v2.ft1": StringSet{
		"FT1": true,
	},
	"DFT_P03v2.msh": StringSet{
		"EVN": true,
		"FT1": true,
		"OBX": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"DFT_P03v2.obx": StringSet{
		"FT1": true,
		"OBX": true,
	},
	"DFT_P03v2.pid": StringSet{
		"FT1": true,
		"OBX": true,
		"PV1": true,
		"PV2": true,
	},
	"DFT_P03v2.pv1": StringSet{
		"FT1": true,
		"OBX": true,
		"PV2": true,
	},
	"DFT_P03v2.pv2": StringSet{
		"FT1": true,
		"OBX": true,
	},
	"DSR_P04v2.dsc": StringSet{},
	"DSR_P04v2.dsp": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"DSR_P04v2.err": StringSet{
		"DSC": true,
		"DSP": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_P04v2.msa": StringSet{
		"DSC": true,
		"DSP": true,
		"ERR": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_P04v2.msh": StringSet{
		"DSC": true,
		"DSP": true,
		"ERR": true,
		"MSA": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_P04v2.qrd": StringSet{
		"DSC": true,
		"DSP": true,
		"QRF": true,
	},
	"DSR_P04v2.qrf": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"DSR_Q01v2.dsc": StringSet{},
	"DSR_Q01v2.dsp": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"DSR_Q01v2.err": StringSet{
		"DSC": true,
		"DSP": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_Q01v2.msa": StringSet{
		"DSC": true,
		"DSP": true,
		"ERR": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_Q01v2.msh": StringSet{
		"DSC": true,
		"DSP": true,
		"ERR": true,
		"MSA": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_Q01v2.qrd": StringSet{
		"DSC": true,
		"DSP": true,
		"QRF": true,
	},
	"DSR_Q01v2.qrf": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"DSR_Q03v2.dsc": StringSet{},
	"DSR_Q03v2.dsp": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"DSR_Q03v2.msa": StringSet{
		"DSC": true,
		"DSP": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_Q03v2.msh": StringSet{
		"DSC": true,
		"DSP": true,
		"MSA": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_Q03v2.qrd": StringSet{
		"DSC": true,
		"DSP": true,
		"QRF": true,
	},
	"DSR_Q03v2.qrf": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"DSR_R03v2.dsc": StringSet{},
	"DSR_R03v2.dsp": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"DSR_R03v2.msa": StringSet{
		"DSC": true,
		"DSP": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_R03v2.msh": StringSet{
		"DSC": true,
		"DSP": true,
		"MSA": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_R03v2.qrd": StringSet{
		"DSC": true,
		"DSP": true,
		"QRF": true,
	},
	"DSR_R03v2.qrf": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"MCF_Q02v2.msa": StringSet{},
	"MCF_Q02v2.msh": StringSet{
		"MSA": true,
	},
	"MFD_M01v2.mfa": StringSet{
		"MFA": true,
	},
	"MFD_M01v2.mfi": StringSet{
		"MFA": true,
	},
	"MFD_M01v2.msh": StringSet{
		"MFA": true,
		"MFI": true,
	},
	"MFD_M02v2.mfa": StringSet{
		"MFA": true,
	},
	"MFD_M02v2.mfi": StringSet{
		"MFA": true,
	},
	"MFD_M02v2.msh": StringSet{
		"MFA": true,
		"MFI": true,
	},
	"MFD_M03v2.mfa": StringSet{
		"MFA": true,
	},
	"MFD_M03v2.mfi": StringSet{
		"MFA": true,
	},
	"MFD_M03v2.msh": StringSet{
		"MFA": true,
		"MFI": true,
	},
	"MFK_M01v2.err": StringSet{
		"MFA": true,
		"MFI": true,
	},
	"MFK_M01v2.mfa": StringSet{
		"MFA": true,
	},
	"MFK_M01v2.mfi": StringSet{
		"MFA": true,
	},
	"MFK_M01v2.msa": StringSet{
		"ERR": true,
		"MFA": true,
		"MFI": true,
	},
	"MFK_M01v2.msh": StringSet{
		"ERR": true,
		"MFA": true,
		"MFI": true,
		"MSA": true,
	},
	"MFK_M02v2.err": StringSet{
		"MFA": true,
		"MFI": true,
	},
	"MFK_M02v2.mfa": StringSet{
		"MFA": true,
	},
	"MFK_M02v2.mfi": StringSet{
		"MFA": true,
	},
	"MFK_M02v2.msa": StringSet{
		"ERR": true,
		"MFA": true,
		"MFI": true,
	},
	"MFK_M02v2.msh": StringSet{
		"ERR": true,
		"MFA": true,
		"MFI": true,
		"MSA": true,
	},
	"MFK_M03v2.err": StringSet{
		"MFA": true,
		"MFI": true,
	},
	"MFK_M03v2.mfa": StringSet{
		"MFA": true,
	},
	"MFK_M03v2.mfi": StringSet{
		"MFA": true,
	},
	"MFK_M03v2.msa": StringSet{
		"ERR": true,
		"MFA": true,
		"MFI": true,
	},
	"MFK_M03v2.msh": StringSet{
		"ERR": true,
		"MFA": true,
		"MFI": true,
		"MSA": true,
	},
	"MFN_M01_MFEv2.mfe": StringSet{
		"MFE": true,
	},
	"MFN_M01v2.mfi": StringSet{
		"MFE": true,
	},
	"MFN_M01v2.msh": StringSet{
		"MFE": true,
		"MFI": true,
	},
	"MFN_M02_MFEv2.mfe": StringSet{
		"MFE": true,
	},
	"MFN_M02v2.mfi": StringSet{
		"MFE": true,
	},
	"MFN_M02v2.msh": StringSet{
		"MFE": true,
		"MFI": true,
	},
	"MFN_M03_MFEv2.mfe": StringSet{
		"MFE": true,
	},
	"MFN_M03v2.mfi": StringSet{
		"MFE": true,
	},
	"MFN_M03v2.msh": StringSet{
		"MFE": true,
		"MFI": true,
	},
	"MFQ_M01v2.dsc": StringSet{},
	"MFQ_M01v2.msh": StringSet{
		"DSC": true,
		"QRD": true,
		"QRF": true,
	},
	"MFQ_M01v2.qrd": StringSet{
		"DSC": true,
		"QRF": true,
	},
	"MFQ_M01v2.qrf": StringSet{
		"DSC": true,
	},
	"MFQ_M02v2.dsc": StringSet{},
	"MFQ_M02v2.msh": StringSet{
		"DSC": true,
		"QRD": true,
		"QRF": true,
	},
	"MFQ_M02v2.qrd": StringSet{
		"DSC": true,
		"QRF": true,
	},
	"MFQ_M02v2.qrf": StringSet{
		"DSC": true,
	},
	"MFQ_M03v2.dsc": StringSet{},
	"MFQ_M03v2.msh": StringSet{
		"DSC": true,
		"QRD": true,
		"QRF": true,
	},
	"MFQ_M03v2.qrd": StringSet{
		"DSC": true,
		"QRF": true,
	},
	"MFQ_M03v2.qrf": StringSet{
		"DSC": true,
	},
	"MFR_M01_MFEv2.mfe": StringSet{
		"DSC": true,
		"MFE": true,
	},
	"MFR_M01v2.dsc": StringSet{},
	"MFR_M01v2.err": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M01v2.mfi": StringSet{
		"DSC": true,
		"MFE": true,
	},
	"MFR_M01v2.msa": StringSet{
		"DSC": true,
		"ERR": true,
		"MFE": true,
		"MFI": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M01v2.msh": StringSet{
		"DSC": true,
		"ERR": true,
		"MFE": true,
		"MFI": true,
		"MSA": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M01v2.qrd": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
		"QRF": true,
	},
	"MFR_M01v2.qrf": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
	},
	"MFR_M02_MFEv2.mfe": StringSet{
		"DSC": true,
		"MFE": true,
	},
	"MFR_M02v2.dsc": StringSet{},
	"MFR_M02v2.err": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M02v2.mfi": StringSet{
		"DSC": true,
		"MFE": true,
	},
	"MFR_M02v2.msa": StringSet{
		"DSC": true,
		"ERR": true,
		"MFE": true,
		"MFI": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M02v2.msh": StringSet{
		"DSC": true,
		"ERR": true,
		"MFE": true,
		"MFI": true,
		"MSA": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M02v2.qrd": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
		"QRF": true,
	},
	"MFR_M02v2.qrf": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
	},
	"MFR_M03_MFEv2.mfe": StringSet{
		"DSC": true,
		"MFE": true,
	},
	"MFR_M03v2.dsc": StringSet{},
	"MFR_M03v2.err": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M03v2.mfi": StringSet{
		"DSC": true,
		"MFE": true,
	},
	"MFR_M03v2.msa": StringSet{
		"DSC": true,
		"ERR": true,
		"MFE": true,
		"MFI": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M03v2.msh": StringSet{
		"DSC": true,
		"ERR": true,
		"MFE": true,
		"MFI": true,
		"MSA": true,
		"QRD": true,
		"QRF": true,
	},
	"MFR_M03v2.qrd": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
		"QRF": true,
	},
	"MFR_M03v2.qrf": StringSet{
		"DSC": true,
		"MFE": true,
		"MFI": true,
	},
	"NMD_N01_NCK_NSCv2.nsc": StringSet{
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMD_N01_NCK_NSCv2.nte": StringSet{
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMD_N01_NCK_NSTv2.nst": StringSet{
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMD_N01_NCK_NSTv2.nte": StringSet{
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMD_N01_NCKv2.nck": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMD_N01_NCKv2.nte": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMD_N01v2.msh": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMQ_N02_NCKv2.nck": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
	},
	"NMQ_N02_NCKv2.nsc": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
	},
	"NMQ_N02_NCKv2.nst": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
	},
	"NMQ_N02v2.msh": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"QRD": true,
		"QRF": true,
	},
	"NMQ_N02v2.qrd": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"QRF": true,
	},
	"NMQ_N02v2.qrf": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
	},
	"NMR_N02_NCKv2.nck": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMR_N02_NCKv2.nsc": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMR_N02_NCKv2.nst": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMR_N02_NCKv2.nte1": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMR_N02_NCKv2.nte2": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMR_N02_NCKv2.nte3": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"NMR_N02v2.err": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
		"QRD": true,
	},
	"NMR_N02v2.msa": StringSet{
		"ERR": true,
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
		"QRD": true,
	},
	"NMR_N02v2.msh": StringSet{
		"ERR": true,
		"MSA": true,
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
		"QRD": true,
	},
	"NMR_N02v2.qrd": StringSet{
		"NCK": true,
		"NSC": true,
		"NST": true,
		"NTE": true,
	},
	"ORF_R04_OBR_OBXv2.nte": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
	},
	"ORF_R04_OBR_OBXv2.obx": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
	},
	"ORF_R04_OBRv2.nte": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
	},
	"ORF_R04_OBRv2.obr": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
	},
	"ORF_R04_OBRv2.orc": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
	},
	"ORF_R04_QRDv2.nte": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
	},
	"ORF_R04_QRDv2.pid": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
	},
	"ORF_R04_QRDv2.qrd": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
	},
	"ORF_R04_QRDv2.qrf": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
	},
	"ORF_R04v2.dsc": StringSet{},
	"ORF_R04v2.msa": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
	},
	"ORF_R04v2.msh": StringSet{
		"DSC": true,
		"MSA": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"QRD": true,
		"QRF": true,
	},
	"ORM_O01_ORCv2.blg": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.nte1": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.nte2": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.obr": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.obx": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.ods": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.odt": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.orc": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.rq1": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.rqd": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_ORCv2.rxo": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_PIDv2.nte": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PV1": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_PIDv2.pid": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PV1": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01_PIDv2.pv1": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01v2.msh": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PID": true,
		"PV1": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORM_O01v2.nte": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PID": true,
		"PV1": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PID_ORCv2.nte": StringSet{
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PID_ORCv2.obr": StringSet{
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PID_ORCv2.ods": StringSet{
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PID_ORCv2.odt": StringSet{
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PID_ORCv2.orc": StringSet{
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PID_ORCv2.rq1": StringSet{
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PID_ORCv2.rqd": StringSet{
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PID_ORCv2.rxo": StringSet{
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PIDv2.nte": StringSet{
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02_PIDv2.pid": StringSet{
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02v2.msa": StringSet{
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PID": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02v2.msh": StringSet{
		"MSA": true,
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PID": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORR_O02v2.nte": StringSet{
		"NTE": true,
		"OBR": true,
		"ODS": true,
		"ODT": true,
		"ORC": true,
		"PID": true,
		"RQ1": true,
		"RQD": true,
		"RXO": true,
	},
	"ORU_R01_PID_OBR_OBXv2.nte": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
	},
	"ORU_R01_PID_OBR_OBXv2.obx": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
	},
	"ORU_R01_PID_OBRv2.nte": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
	},
	"ORU_R01_PID_OBRv2.obr": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
	},
	"ORU_R01_PID_OBRv2.orc": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
	},
	"ORU_R01_PIDv2.nte": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PV1": true,
	},
	"ORU_R01_PIDv2.pid": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PV1": true,
	},
	"ORU_R01_PIDv2.pv1": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PV1": true,
	},
	"ORU_R01v2.dsc": StringSet{},
	"ORU_R01v2.msh": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PV1": true,
	},
	"ORU_R03_PID_OBR_OBXv2.nte": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
	},
	"ORU_R03_PID_OBR_OBXv2.obx": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
	},
	"ORU_R03_PID_OBRv2.nte": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
	},
	"ORU_R03_PID_OBRv2.obr": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
	},
	"ORU_R03_PID_OBRv2.orc": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PV1": true,
	},
	"ORU_R03_PIDv2.nte": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PV1": true,
	},
	"ORU_R03_PIDv2.pid": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PV1": true,
	},
	"ORU_R03_PIDv2.pv1": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PV1": true,
	},
	"ORU_R03v2.dsc": StringSet{},
	"ORU_R03v2.msh": StringSet{
		"DSC": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PID": true,
		"PV1": true,
	},
	"ORU_R32_PID_OBR_OBXv2.nte": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PID_OBR_OBXv2.obx": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PID_OBRv2.cti": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PID_OBRv2.nte": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PID_OBRv2.obr": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PID_OBRv2.orc": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PIDv2.nk1": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PIDv2.nte": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PIDv2.pd1": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PIDv2.pid": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PIDv2.pv1": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32_PIDv2.pv2": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"ORU_R32v2.dsc": StringSet{},
	"ORU_R32v2.msh": StringSet{
		"CTI": true,
		"DSC": true,
		"NK1": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
		"PV2": true,
	},
	"QRY_A19v2.msh": StringSet{
		"QRD": true,
		"QRF": true,
	},
	"QRY_A19v2.qrd": StringSet{
		"QRF": true,
	},
	"QRY_A19v2.qrf": StringSet{},
	"QRY_P04v2.dsc": StringSet{},
	"QRY_P04v2.msh": StringSet{
		"DSC": true,
		"QRD": true,
		"QRF": true,
	},
	"QRY_P04v2.qrd": StringSet{
		"DSC": true,
		"QRF": true,
	},
	"QRY_P04v2.qrf": StringSet{
		"DSC": true,
	},
	"QRY_Q01v2.dsc": StringSet{},
	"QRY_Q01v2.msh": StringSet{
		"DSC": true,
		"QRD": true,
		"QRF": true,
	},
	"QRY_Q01v2.qrd": StringSet{
		"DSC": true,
		"QRF": true,
	},
	"QRY_Q01v2.qrf": StringSet{
		"DSC": true,
	},
	"QRY_Q02v2.dsc": StringSet{},
	"QRY_Q02v2.msh": StringSet{
		"DSC": true,
		"QRD": true,
		"QRF": true,
	},
	"QRY_Q02v2.qrd": StringSet{
		"DSC": true,
		"QRF": true,
	},
	"QRY_Q02v2.qrf": StringSet{
		"DSC": true,
	},
	"QRY_R02v2.dsc": StringSet{},
	"QRY_R02v2.msh": StringSet{
		"DSC": true,
		"QRD": true,
		"QRF": true,
	},
	"QRY_R02v2.qrd": StringSet{
		"DSC": true,
		"QRF": true,
	},
	"QRY_R02v2.qrf": StringSet{
		"DSC": true,
	},
	"UDM_Q05v2.dsc": StringSet{},
	"UDM_Q05v2.dsp": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"UDM_Q05v2.msh": StringSet{
		"DSC": true,
		"DSP": true,
		"URD": true,
		"URS": true,
	},
	"UDM_Q05v2.urd": StringSet{
		"DSC": true,
		"DSP": true,
		"URS": true,
	},
	"UDM_Q05v2.urs": StringSet{
		"DSC": true,
		"DSP": true,
	},
}
