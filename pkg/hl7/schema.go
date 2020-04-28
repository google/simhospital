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

// This file contains the schemas for HL7 messages, segments and values.
// It has been auto-generated from the HL7v2 schemas.

package hl7

import "reflect"

// AD represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type AD struct {
	StreetAddress              *ST `hl7:"false,Street Address"`
	OtherDesignation           *ST `hl7:"false,Other Designation"`
	City                       *ST `hl7:"false,City"`
	StateOrProvince            *ST `hl7:"false,State Or Province"`
	ZipOrPostalCode            *ST `hl7:"false,Zip Or Postal Code"`
	Country                    *ID `hl7:"false,Country"`
	AddressType                *ID `hl7:"false,Address Type"`
	OtherGeographicDesignation *ST `hl7:"false,Other Geographic Designation"`
}

// AUI represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type AUI struct {
	AuthorizationNumber *ST `hl7:"false,Authorization Number"`
	Date                *DT `hl7:"false,Date"`
	Source              *ST `hl7:"false,Source"`
}

// CCD represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type CCD struct {
	WhenToChargeCode *ID `hl7:"false,When To Charge Code"`
	DateTime         *TS `hl7:"false,Date/Time"`
}

// CCP represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type CCP struct {
	ChannelCalibrationSensitivityCorrectionFactor *NM `hl7:"false,Channel Calibration Sensitivity Correction Factor"`
	ChannelCalibrationBaseline                    *NM `hl7:"false,Channel Calibration Baseline"`
	ChannelCalibrationTimeSkew                    *NM `hl7:"false,Channel Calibration Time Skew"`
}

// CD represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type CD struct {
	ChannelIdentifier            *WVI `hl7:"false,Channel Identifier"`
	WaveformSource               *WVS `hl7:"false,Waveform Source"`
	ChannelSensitivityUnits      *CSU `hl7:"false,Channel Sensitivity/Units"`
	ChannelCalibrationParameters *CCP `hl7:"false,Channel Calibration Parameters"`
	ChannelSamplingFrequency     *NM  `hl7:"false,Channel Sampling Frequency"`
	MinimumMaximumDataValues     *NR  `hl7:"false,Minimum/Maximum Data Values"`
}

// CE represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type CE struct {
	Identifier                  *ST `hl7:"false,Identifier"`
	Text                        *ST `hl7:"false,Text"`
	NameOfCodingSystem          *IS `hl7:"false,Name Of Coding System"`
	AlternateIdentifier         *ST `hl7:"false,Alternate Identifier"`
	AlternateText               *ST `hl7:"false,Alternate Text"`
	NameOfAlternateCodingSystem *IS `hl7:"false,Name Of Alternate Coding System"`
}

// CF represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type CF struct {
	Identifier                  *ID `hl7:"false,Identifier"`
	FormattedText               *FT `hl7:"false,Formatted Text"`
	NameOfCodingSystem          *IS `hl7:"false,Name Of Coding System"`
	AlternateIdentifier         *ID `hl7:"false,Alternate Identifier"`
	AlternateFormattedText      *FT `hl7:"false,Alternate Formatted Text"`
	NameOfAlternateCodingSystem *IS `hl7:"false,Name Of Alternate Coding System"`
}

// CK represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type CK struct {
	IDNumber                                   *NM `hl7:"false,ID Number"`
	CheckDigit                                 *NM `hl7:"false,Check Digit"`
	CodeIdentifyingTheCheckDigitSchemeEmployed *ID `hl7:"false,Code Identifying The Check Digit Scheme Employed"`
	AssigningAuthority                         *HD `hl7:"false,Assigning Authority"`
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
// Definition from HL7 2.3
type CM_ABS_RANGE struct {
	Range            *CM `hl7:"false,Range"`
	NumericChange    *NM `hl7:"false,Numeric Change"`
	PercentPerChange *NM `hl7:"false,Percent Per Change"`
	Days             *NM `hl7:"false,Days"`
}

// CM_AUI represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_AUI struct {
	AuthorizationNumber *ST `hl7:"false,Authorization Number"`
	Date                *TS `hl7:"false,Date"`
	Source              *ST `hl7:"false,Source"`
}

// CM_BATCH_TOTAL represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_BATCH_TOTAL struct {
	BatchTotal1 *NM `hl7:"false,Batch Total 1"`
	BatchTotal2 *NM `hl7:"false,Batch Total 2"`
}

// CM_CCD represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_CCD struct {
	WhenToChargeCode *ID `hl7:"false,When To Charge Code"`
	DateTime         *TS `hl7:"false,Date/Time"`
}

// CM_DDI represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_DDI struct {
	DelayDays    *NM `hl7:"false,Delay Days"`
	Amount       *NM `hl7:"false,Amount"`
	NumberOfDays *NM `hl7:"false,Number Of Days"`
}

// CM_DIN represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_DIN struct {
	Date            *TS `hl7:"false,Date"`
	InstitutionName *CE `hl7:"false,Institution Name"`
}

// CM_DLD represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_DLD struct {
	DischargeLocation *ID `hl7:"false,Discharge Location"`
	EffectiveDate     *TS `hl7:"false,Effective Date"`
}

// CM_DLT represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_DLT struct {
	Range            *CM `hl7:"false,Range"`
	NumericThreshold *NM `hl7:"false,Numeric Threshold"`
	Change           *ST `hl7:"false,Change"`
	LengthOfTimeDays *NM `hl7:"false,Length Of Time-Days"`
}

// CM_DTN represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_DTN struct {
	DayType      *IS `hl7:"false,Day Type"`
	NumberOfDays *NM `hl7:"false,Number Of Days"`
}

// CM_EIP represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_EIP struct {
	ParentSPlacerOrderNumber *EI `hl7:"false,Parent´s Placer Order Number"`
	ParentSFillerOrderNumber *EI `hl7:"false,Parent´s Filler Order Number"`
}

// CM_ELD represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_ELD struct {
	SegmentID            *ST `hl7:"false,Segment ID"`
	Sequence             *NM `hl7:"false,Sequence"`
	FieldPosition        *NM `hl7:"false,Field Position"`
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
// Definition from HL7 2.3
type CM_LA1 struct {
	PointOfCare                *ST `hl7:"false,Point Of Care"`
	Room                       *IS `hl7:"false,Room"`
	Bed                        *IS `hl7:"false,Bed"`
	Facility                   *HD `hl7:"false,Facility"`
	LocationStatus             *IS `hl7:"false,Location Status"`
	PersonLocationType         *IS `hl7:"false,Person Location Type"`
	Building                   *IS `hl7:"false,Building"`
	Floor                      *ST `hl7:"false,Floor"`
	StreetAddress              *ST `hl7:"false,Street Address"`
	OtherDesignation           *ST `hl7:"false,Other Designation"`
	City                       *ST `hl7:"false,City"`
	StateOrProvince            *ST `hl7:"false,State Or Province"`
	ZipOrPostalCode            *ST `hl7:"false,Zip Or Postal Code"`
	Country                    *ID `hl7:"false,Country"`
	AddressType                *ID `hl7:"false,Address Type"`
	OtherGeographicDesignation *ST `hl7:"false,Other Geographic Designation"`
}

// CM_LICENSE_NO represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CM_LICENSE_NO struct {
	LicenseNumber               *ST `hl7:"false,License Number"`
	IssuingStateProvinceCountry *ST `hl7:"false,Issuing State,Province,Country"`
}

// CM_MOC represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_MOC struct {
	DollarAmount *MO `hl7:"false,Dollar Amount"`
	ChargeCode   *CE `hl7:"false,Charge Code"`
}

// CM_MSG represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_MSG struct {
	MessageType  *ID `hl7:"false,Message Type"`
	TriggerEvent *ID `hl7:"false,Trigger Event"`
}

// CM_NDL represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_NDL struct {
	Name               *CN `hl7:"false,Name"`
	StartDateTime      *TS `hl7:"false,Start Date/Time"`
	EndDateTime        *TS `hl7:"false,End Date/Time"`
	PointOfCare        *IS `hl7:"false,Point Of Care"`
	Room               *IS `hl7:"false,Room"`
	Bed                *IS `hl7:"false,Bed"`
	Facility           *HD `hl7:"false,Facility"`
	LocationStatus     *IS `hl7:"false,Location Status"`
	PersonLocationType *IS `hl7:"false,Person Location Type"`
	Building           *IS `hl7:"false,Building"`
	Floor              *ST `hl7:"false,Floor"`
}

// CM_OCD represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_OCD struct {
	OccurrenceCode *CE `hl7:"false,Occurrence Code"`
	OccurrenceDate *DT `hl7:"false,Occurrence Date"`
}

// CM_OSP represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_OSP struct {
	OccurrenceSpanCode      *CE `hl7:"false,Occurrence Span Code"`
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
// Definition from HL7 2.3
type CM_PCF struct {
	PreCertificationPatientType *IS `hl7:"false,Pre-Certification Patient Type"`
	PreCertificationRequired    *ID `hl7:"false,Pre-Certification Required"`
	PreCertificationWindwow     *TS `hl7:"false,Pre-Certification Windwow"`
}

// CM_PEN represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_PEN struct {
	PenaltyType   *IS `hl7:"false,Penalty Type"`
	PenaltyAmount *NM `hl7:"false,Penalty Amount"`
}

// CM_PI represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_PI struct {
	IDNumber            *ST `hl7:"false,ID Number"`
	TypeOfIDNumber      *IS `hl7:"false,Type Of ID Number"`
	OtherQualifyingInfo *ST `hl7:"false,Other Qualifying Info"`
}

// CM_PIP represents the corresponding HL7 datatype.
// Definition from HL7 2.3
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
// Definition from HL7 2.3
type CM_PLN struct {
	IDNumber                 *ST `hl7:"false,ID Number"`
	TypeOfIDNumber           *IS `hl7:"false,Type Of ID Number"`
	StateOtherQualifyingInfo *ST `hl7:"false,State/Other Qualifying Info"`
	ExpirationDate           *DT `hl7:"false,Expiration Date"`
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

// CM_PRL represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_PRL struct {
	OBX3ObservationIdentifierOfParentResult *CE `hl7:"false,OBX-3 Observation Identifier Of Parent Result"`
	OBX4SubIDOfParentResult                 *ST `hl7:"false,OBX-4 Sub-ID Of Parent Result"`
	PartOfOBX5ObservationResultFromParent   *TX `hl7:"false,Part Of OBX-5 Observation Result From Parent"`
}

// CM_PTA represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_PTA struct {
	PolicyType  *IS `hl7:"false,Policy Type"`
	AmountClass *IS `hl7:"false,Amount Class"`
	Amount      *NM `hl7:"false,Amount"`
}

// CM_RANGE represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_RANGE struct {
	LowValue  *CE `hl7:"false,Low Value"`
	HighValue *CE `hl7:"false,High Value"`
}

// CM_RFR represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_RFR struct {
	ReferenceRange *CM `hl7:"false,Reference Range"`
	Sex            *IS `hl7:"false,Sex"`
	AgeRange       *CM `hl7:"false,Age Range"`
	AgeGestation   *CM `hl7:"false,Age Gestation"`
	Species        *TX `hl7:"false,Species"`
	RaceSubspecies *ST `hl7:"false,Race/Subspecies"`
	Conditions     *TX `hl7:"false,Conditions"`
}

// CM_RI represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_RI struct {
	RepeatPattern        *IS `hl7:"false,Repeat Pattern"`
	ExplicitTimeInterval *ST `hl7:"false,Explicit Time Interval"`
}

// CM_RMC represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_RMC struct {
	RoomType       *IS `hl7:"false,Room Type"`
	AmountType     *IS `hl7:"false,Amount Type"`
	CoverageAmount *NM `hl7:"false,Coverage Amount"`
}

// CM_SPD represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_SPD struct {
	SpecialtyName       *ST `hl7:"false,Specialty Name"`
	GoverningBoard      *ST `hl7:"false,Governing Board"`
	EligibleOrCertified *ID `hl7:"false,Eligible Or Certified"`
	DateOfCertification *DT `hl7:"false,Date Of Certification"`
}

// CM_SPS represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_SPS struct {
	SpecimenSourceNameOrCode     *CE `hl7:"false,Specimen Source Name Or Code"`
	Additives                    *TX `hl7:"false,Additives"`
	Freetext                     *TX `hl7:"false,Freetext"`
	BodySite                     *CE `hl7:"false,Body Site"`
	SiteModifier                 *CE `hl7:"false,Site Modifier"`
	CollectionModifierMethodCode *CE `hl7:"false,Collection Modifier Method Code"`
}

// CM_UVC represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_UVC struct {
	ValueCode   *IS `hl7:"false,Value Code"`
	ValueAmount *NM `hl7:"false,Value Amount"`
}

// CM_VR represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_VR struct {
	FirstDataCodeValue *ST `hl7:"false,First Data Code Value"`
	LastDataCodeCalue  *ST `hl7:"false,Last Data Code Calue"`
}

// CM_WVI represents the corresponding HL7 datatype.
// Definition from HL7 2.3
type CM_WVI struct {
	ChannelNumber *NM `hl7:"false,Channel Number"`
	ChannelName   *ST `hl7:"false,Channel Name"`
}

// CN represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type CN struct {
	IDNumber                                    *ST `hl7:"false,ID Number"`
	FamilyName                                  *FN `hl7:"false,Family Name"`
	GivenName                                   *ST `hl7:"false,Given Name"`
	SecondAndFurtherGivenNamesOrInitialsThereof *ST `hl7:"false,Second And Further Given Names Or Initials Thereof"`
	SuffixEGJROrIII                             *ST `hl7:"false,Suffix (E.G., JR Or III)"`
	PrefixEGDR                                  *ST `hl7:"false,Prefix (E.G., DR)"`
	DegreeEGMD                                  *IS `hl7:"false,Degree (E.G., MD)"`
	SourceTable                                 *IS `hl7:"false,Source Table"`
	AssigningAuthority                          *HD `hl7:"false,Assigning Authority"`
}

// CNE represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type CNE struct {
	Identifier                     *ST `hl7:"false,Identifier"`
	Text                           *ST `hl7:"false,Text"`
	NameOfCodingSystem             *IS `hl7:"false,Name Of Coding System"`
	AlternateIdentifier            *ST `hl7:"false,Alternate Identifier"`
	AlternateText                  *ST `hl7:"false,Alternate Text"`
	NameOfAlternateCodingSystem    *IS `hl7:"false,Name Of Alternate Coding System"`
	CodingSystemVersionID          *ST `hl7:"false,Coding System Version ID"`
	AlternateCodingSystemVersionID *ST `hl7:"false,Alternate Coding System Version ID"`
	OriginalText                   *ST `hl7:"false,Original Text"`
}

// CNN represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type CNN struct {
	IDNumber                                    *ST `hl7:"false,ID Number"`
	FamilyName                                  *ST `hl7:"false,Family Name"`
	GivenName                                   *ST `hl7:"false,Given Name"`
	SecondAndFurtherGivenNamesOrInitialsThereof *ST `hl7:"false,Second And Further Given Names Or Initials Thereof"`
	SuffixEGJROrIII                             *ST `hl7:"false,Suffix (E.G., JR Or III)"`
	PrefixEGDR                                  *ST `hl7:"false,Prefix (E.G., DR)"`
	DegreeEGMD                                  *IS `hl7:"false,Degree (E.G., MD)"`
	SourceTable                                 *IS `hl7:"false,Source Table"`
	AssigningAuthorityNamespaceID               *IS `hl7:"false,Assigning Authority Namespace ID"`
	AssigningAuthorityUniversalID               *ST `hl7:"false,Assigning Authority Universal ID"`
	AssigningAuthorityUniversalIDType           *ID `hl7:"false,Assigning Authority Universal ID Type"`
}

// CNS231 represents the corresponding HL7 datatype.
// Definition from HL7 2.3.1
type CNS231 struct {
	StartingNotificationReferenceNumber *NM `hl7:"false,Starting Notification Reference Number"`
	EndingNotificationReferenceNumber   *NM `hl7:"false,Ending Notification Reference Number"`
	StartingNotificationDateTime        *TS `hl7:"false,Starting Notification Date/Time"`
	EndingNotificationDateTime          *TS `hl7:"false,Ending Notification Date/Time"`
	StartingNotificationCode            *CE `hl7:"false,Starting Notification Code"`
	EndingNotificationCode              *CE `hl7:"false,Ending Notification Code"`
	DegreeEGMD                          *IS `hl7:"false,Degree (E.G., MD)"`
	SourceTable                         *IS `hl7:"false,Source Table"`
	AssigningAuthorityNamespaceID       *IS `hl7:"false,Assigning Authority Namespace ID"`
	AssigningAuthorityUniversalID       *ST `hl7:"false,Assigning Authority Universal ID"`
	AssigningAuthorityUniversalIDType   *ID `hl7:"false,Assigning Authority Universal ID Type"`
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

// CP represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type CP struct {
	Price      *MO `hl7:"false,Price"`
	PriceType  *ID `hl7:"false,Price Type"`
	FromValue  *NM `hl7:"false,From Value"`
	ToValue    *NM `hl7:"false,To Value"`
	RangeUnits *CE `hl7:"false,Range Units"`
	RangeType  *ID `hl7:"false,Range Type"`
}

// CQ represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type CQ struct {
	Quantity *NM `hl7:"false,Quantity"`
	Units    *CE `hl7:"false,Units"`
}

// CQ_QUANTITY represents the corresponding HL7 datatype.
// Definition from HL7 2.2
type CQ_QUANTITY struct {
	Quantity *ST `hl7:"false,Quantity"`
	Units    *ST `hl7:"false,Units"`
}

// CSU represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type CSU struct {
	ChannelSensitivity                 *NM `hl7:"false,Channel Sensitivity"`
	UnitOfMeasureIdentifier            *ST `hl7:"false,Unit Of Measure Identifier"`
	UnitOfMeasureDescription           *ST `hl7:"false,Unit Of Measure Description"`
	UnitOfMeasureCodingSystem          *IS `hl7:"false,Unit Of Measure Coding System"`
	AlternateUnitOfMeasureIdentifier   *ST `hl7:"false,Alternate Unit Of Measure Identifier"`
	AlternateUnitOfMeasureDescription  *ST `hl7:"false,Alternate Unit Of Measure Description"`
	AlternateUnitOfMeasureCodingSystem *IS `hl7:"false,Alternate Unit Of Measure Coding System"`
}

// CWE represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type CWE struct {
	Identifier                     *ST `hl7:"false,Identifier"`
	Text                           *ST `hl7:"false,Text"`
	NameOfCodingSystem             *IS `hl7:"false,Name Of Coding System"`
	AlternateIdentifier            *ST `hl7:"false,Alternate Identifier"`
	AlternateText                  *ST `hl7:"false,Alternate Text"`
	NameOfAlternateCodingSystem    *IS `hl7:"false,Name Of Alternate Coding System"`
	CodingSystemVersionID          *ST `hl7:"false,Coding System Version ID"`
	AlternateCodingSystemVersionID *ST `hl7:"false,Alternate Coding System Version ID"`
	OriginalText                   *ST `hl7:"false,Original Text"`
}

// CX represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type CX struct {
	ID                                         *ST `hl7:"false,ID"`
	CheckDigit                                 *ST `hl7:"false,Check Digit"`
	CodeIdentifyingTheCheckDigitSchemeEmployed *ID `hl7:"false,Code Identifying The Check Digit Scheme Employed"`
	AssigningAuthority                         *HD `hl7:"false,Assigning Authority"`
	IdentifierTypeCode                         *ID `hl7:"false,Identifier Type Code"`
	AssigningFacility                          *HD `hl7:"false,Assigning Facility"`
	EffectiveDate                              *DT `hl7:"false,Effective Date"`
	ExpirationDate                             *DT `hl7:"false,Expiration Date"`
}

// DDI represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type DDI struct {
	DelayDays    *NM `hl7:"false,Delay Days"`
	Amount       *NM `hl7:"false,Amount"`
	NumberOfDays *NM `hl7:"false,Number Of Days"`
}

// DIN represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type DIN struct {
	Date            *TS `hl7:"false,Date"`
	InstitutionName *CE `hl7:"false,Institution Name"`
}

// DLD represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type DLD struct {
	DischargeLocation *ID `hl7:"false,Discharge Location"`
	EffectiveDate     *TS `hl7:"false,Effective Date"`
}

// DLN represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type DLN struct {
	DriverSLicenseNumber        *ST `hl7:"false,Driver´s License Number"`
	IssuingStateProvinceCountry *IS `hl7:"false,Issuing State, Province, Country"`
	ExpirationDate              *DT `hl7:"false,Expiration Date"`
}

// DLT represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type DLT struct {
	Range             *NR `hl7:"false,Range"`
	NumericThreshold  *NM `hl7:"false,Numeric Threshold"`
	ChangeComputation *ST `hl7:"false,Change Computation"`
	LengthOfTimeDays  *NM `hl7:"false,Length Of Time-Days"`
}

// DR represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type DR struct {
	RangeStartDateTime *TS `hl7:"false,Range Start Date/Time"`
	RangeEndDateTime   *TS `hl7:"false,Range End Date/Time"`
}

// DTN represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type DTN struct {
	DayType      *IS `hl7:"false,Day Type"`
	NumberOfDays *NM `hl7:"false,Number Of Days"`
}

// ED represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type ED struct {
	SourceApplication *HD `hl7:"false,Source Application"`
	TypeOfData        *ID `hl7:"false,Type Of Data"`
	DataSubtype       *ID `hl7:"false,Data Subtype"`
	Encoding          *ID `hl7:"false,Encoding"`
	Data              *ST `hl7:"false,Data"`
}

// EI represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type EI struct {
	EntityIdentifier *ST `hl7:"false,Entity Identifier"`
	NamespaceID      *IS `hl7:"false,Namespace ID"`
	UniversalID      *ST `hl7:"false,Universal ID"`
	UniversalIDType  *ID `hl7:"false,Universal ID Type"`
}

// EIP represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type EIP struct {
	ParentSPlacerOrderNumber *EI `hl7:"false,Parent´s Placer Order Number"`
	ParentSFillerOrderNumber *EI `hl7:"false,Parent´s Filler Order Number"`
}

// ELD represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type ELD struct {
	SegmentID            *ST `hl7:"false,Segment ID"`
	Sequence             *NM `hl7:"false,Sequence"`
	FieldPosition        *NM `hl7:"false,Field Position"`
	CodeIdentifyingError *CE `hl7:"false,Code Identifying Error"`
}

// FC represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type FC struct {
	FinancialClass *IS `hl7:"false,Financial Class"`
	EffectiveDate  *TS `hl7:"false,Effective Date"`
}

// FN represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type FN struct {
	Surname                        *ST `hl7:"false,Surname"`
	OwnSurnamePrefix               *ST `hl7:"false,Own Surname Prefix"`
	OwnSurname                     *ST `hl7:"false,Own Surname"`
	SurnamePrefixFromPartnerSpouse *ST `hl7:"false,Surname Prefix From Partner/Spouse"`
	SurnameFromPartnerSpouse       *ST `hl7:"false,Surname From Partner/Spouse"`
}

// HD represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type HD struct {
	NamespaceID     *IS `hl7:"false,Namespace ID"`
	UniversalID     *ST `hl7:"false,Universal ID"`
	UniversalIDType *ID `hl7:"false,Universal ID Type"`
}

// JCC represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type JCC struct {
	JobCode  *IS `hl7:"false,Job Code"`
	JobClass *IS `hl7:"false,Job Class"`
}

// LA1 represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type LA1 struct {
	PointOfCare        *IS `hl7:"false,Point Of Care"`
	Room               *IS `hl7:"false,Room"`
	Bed                *IS `hl7:"false,Bed"`
	Facility           *HD `hl7:"false,Facility"`
	LocationStatus     *IS `hl7:"false,Location Status"`
	PersonLocationType *IS `hl7:"false,Person Location Type"`
	Building           *IS `hl7:"false,Building"`
	Floor              *IS `hl7:"false,Floor"`
	Address            *AD `hl7:"false,Address"`
}

// LA2 represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type LA2 struct {
	PointOfCare                *IS `hl7:"false,Point Of Care"`
	Room                       *IS `hl7:"false,Room"`
	Bed                        *IS `hl7:"false,Bed"`
	Facility                   *HD `hl7:"false,Facility"`
	LocationStatus             *IS `hl7:"false,Location Status"`
	PersonLocationType         *IS `hl7:"false,Person Location Type"`
	Building                   *IS `hl7:"false,Building"`
	Floor                      *IS `hl7:"false,Floor"`
	StreetAddress              *ST `hl7:"false,Street Address"`
	OtherDesignation           *ST `hl7:"false,Other Designation"`
	City                       *ST `hl7:"false,City"`
	StateOrProvince            *ST `hl7:"false,State Or Province"`
	ZipOrPostalCode            *ST `hl7:"false,Zip Or Postal Code"`
	Country                    *ID `hl7:"false,Country"`
	AddressType                *ID `hl7:"false,Address Type"`
	OtherGeographicDesignation *ST `hl7:"false,Other Geographic Designation"`
}

// MA represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type MA struct {
	Sample1FromChannel1 *NM `hl7:"false,Sample 1 From Channel 1"`
	Sample1FromChannel2 *NM `hl7:"false,Sample 1 From Channel 2"`
	Sample1FromChannel3 *NM `hl7:"false,Sample 1 From Channel 3"`
	Sample1FromChannel4 *NM `hl7:"false,Sample 1 From Channel 4"`
	Sample1FromChannel5 *NM `hl7:"false,Sample 1 From Channel 5"`
	Sample1FromChannel6 *NM `hl7:"false,Sample 1 From Channel 6"`
}

// MO represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type MO struct {
	Quantity     *NM `hl7:"false,Quantity"`
	Denomination *ID `hl7:"false,Denomination"`
}

// MOC represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type MOC struct {
	DollarAmount *MO `hl7:"false,Dollar Amount"`
	ChargeCode   *CE `hl7:"false,Charge Code"`
}

// MOP represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type MOP struct {
	MoneyOrPercentageIndicator *IS `hl7:"false,Money Or Percentage Indicator"`
	MoneyOrPercentageQuantity  *NM `hl7:"false,Money Or Percentage Quantity"`
}

// MSG represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type MSG struct {
	MessageType      *ID `hl7:"false,Message Type"`
	TriggerEvent     *ID `hl7:"false,Trigger Event"`
	MessageStructure *ID `hl7:"false,Message Structure"`
}

// NA represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type NA struct {
	Value1 *NM `hl7:"false,Value1"`
	Value2 *NM `hl7:"false,Value2"`
	Value3 *NM `hl7:"false,Value3"`
	Value4 *NM `hl7:"false,Value4"`
}

// NDL represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type NDL struct {
	Name               *CNN `hl7:"false,Name"`
	StartDateTime      *TS  `hl7:"false,Start Date/Time"`
	EndDateTime        *TS  `hl7:"false,End Date/Time"`
	PointOfCare        *IS  `hl7:"false,Point Of Care"`
	Room               *IS  `hl7:"false,Room"`
	Bed                *IS  `hl7:"false,Bed"`
	Facility           *HD  `hl7:"false,Facility"`
	LocationStatus     *IS  `hl7:"false,Location Status"`
	PersonLocationType *IS  `hl7:"false,Person Location Type"`
	Building           *IS  `hl7:"false,Building"`
	Floor              *IS  `hl7:"false,Floor"`
}

// NR represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type NR struct {
	LowValue  *NM `hl7:"false,Low Value"`
	HighValue *NM `hl7:"false,High Value"`
}

// OCD represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type OCD struct {
	OccurrenceCode *IS `hl7:"false,Occurrence Code"`
	OccurrenceDate *DT `hl7:"false,Occurrence Date"`
}

// OSD represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type OSD struct {
	SequenceResultsFlag               *ID `hl7:"false,Sequence/Results Flag"`
	PlacerOrderNumberEntityIdentifier *ST `hl7:"false,Placer Order Number: Entity Identifier"`
	PlacerOrderNumberNamespaceID      *IS `hl7:"false,Placer Order Number: Namespace ID"`
	FillerOrderNumberEntityIdentifier *ST `hl7:"false,Filler Order Number: Entity Identifier"`
	FillerOrderNumberNamespaceID      *IS `hl7:"false,Filler Order Number: Namespace ID"`
	SequenceConditionValue            *ST `hl7:"false,Sequence Condition Value"`
	MaximumNumberOfRepeats            *NM `hl7:"false,Maximum Number Of Repeats"`
	PlacerOrderNumberUniversalID      *ST `hl7:"false,Placer Order Number: Universal ID"`
	PlacerOrderNumberUniversalIDType  *ID `hl7:"false,Placer Order Number; Universal ID Type"`
	FillerOrderNumberUniversalID      *ST `hl7:"false,Filler Order Number: Universal ID"`
	FillerOrderNumberUniversalIDType  *ID `hl7:"false,Filler Order Number: Universal ID Type"`
}

// OSP represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type OSP struct {
	OccurrenceSpanCode      *CE `hl7:"false,Occurrence Span Code"`
	OccurrenceSpanStartDate *DT `hl7:"false,Occurrence Span Start Date"`
	OccurrenceSpanStopDate  *DT `hl7:"false,Occurrence Span Stop Date"`
}

// PCF represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type PCF struct {
	PreCertificationPatientType *IS `hl7:"false,Pre-Certification Patient Type"`
	PreCertificationRequired    *ID `hl7:"false,Pre-Certification Required"`
	PreCertificationWindow      *TS `hl7:"false,Pre-Certification Window"`
}

// PI represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type PI struct {
	IDNumber            *ST `hl7:"false,ID Number"`
	TypeOfIDNumber      *IS `hl7:"false,Type Of ID Number"`
	OtherQualifyingInfo *ST `hl7:"false,Other Qualifying Info"`
}

// PIP represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type PIP struct {
	Privilege      *CE `hl7:"false,Privilege"`
	PrivilegeClass *CE `hl7:"false,Privilege Class"`
	ExpirationDate *DT `hl7:"false,Expiration Date"`
	ActivationDate *DT `hl7:"false,Activation Date"`
	Facility       *EI `hl7:"false,Facility"`
}

// PL represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type PL struct {
	PointOfCare         *IS `hl7:"false,Point Of Care"`
	Room                *IS `hl7:"false,Room"`
	Bed                 *IS `hl7:"false,Bed"`
	Facility            *HD `hl7:"false,Facility"`
	LocationStatus      *IS `hl7:"false,Location Status"`
	PersonLocationType  *IS `hl7:"false,Person Location Type"`
	Building            *IS `hl7:"false,Building"`
	Floor               *IS `hl7:"false,Floor"`
	LocationDescription *ST `hl7:"false,Location Description"`
}

// PLN represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type PLN struct {
	IDNumber                 *ST `hl7:"false,ID Number"`
	TypeOfIDNumber           *IS `hl7:"false,Type Of ID Number"`
	StateOtherQualifyingInfo *ST `hl7:"false,State/Other Qualifying Info"`
	ExpirationDate           *DT `hl7:"false,Expiration Date"`
}

// PN represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type PN struct {
	FamilyName                                  *FN `hl7:"false,Family Name"`
	GivenName                                   *ST `hl7:"false,Given Name"`
	SecondAndFurtherGivenNamesOrInitialsThereof *ST `hl7:"false,Second And Further Given Names Or Initials Thereof"`
	SuffixEGJROrIII                             *ST `hl7:"false,Suffix (E.G., JR Or III)"`
	PrefixEGDR                                  *ST `hl7:"false,Prefix (E.G., DR)"`
	DegreeEGMD                                  *IS `hl7:"false,Degree (E.G., MD)"`
}

// PPN represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type PPN struct {
	IDNumber                                    *ST `hl7:"false,ID Number"`
	FamilyName                                  *FN `hl7:"false,Family Name"`
	GivenName                                   *ST `hl7:"false,Given Name"`
	SecondAndFurtherGivenNamesOrInitialsThereof *ST `hl7:"false,Second And Further Given Names Or Initials Thereof"`
	SuffixEGJROrIII                             *ST `hl7:"false,Suffix (E.G., JR Or III)"`
	PrefixEGDR                                  *ST `hl7:"false,Prefix (E.G., DR)"`
	DegreeEGMD                                  *IS `hl7:"false,Degree (E.G., MD)"`
	SourceTable                                 *IS `hl7:"false,Source Table"`
	AssigningAuthority                          *HD `hl7:"false,Assigning Authority"`
	NameTypeCode                                *ID `hl7:"false,Name Type Code"`
	IdentifierCheckDigit                        *ST `hl7:"false,Identifier Check Digit"`
	CodeIdentifyingTheCheckDigitSchemeEmployed  *ID `hl7:"false,Code Identifying The Check Digit Scheme Employed"`
	IdentifierTypeCode                          *IS `hl7:"false,Identifier Type Code"`
	AssigningFacility                           *HD `hl7:"false,Assigning Facility"`
	DateTimeActionPerformed                     *TS `hl7:"false,Date/Time Action Performed"`
	NameRepresentationCode                      *ID `hl7:"false,Name Representation Code"`
	NameContext                                 *CE `hl7:"false,Name Context"`
	NameValidityRange                           *DR `hl7:"false,Name Validity Range"`
	NameAssemblyOrder                           *ID `hl7:"false,Name Assembly Order"`
}

// PRL represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type PRL struct {
	OBX3ObservationIdentifierOfParentResult *CE `hl7:"false,OBX-3 Observation Identifier Of Parent Result"`
	OBX4SubIDOfParentResult                 *ST `hl7:"false,OBX-4 Sub-ID Of Parent Result"`
	PartOfOBX5ObservationResultFromParent   *TX `hl7:"false,Part Of OBX-5 Observation Result From Parent"`
}

// PT represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type PT struct {
	ProcessingID   *ID `hl7:"false,Processing ID"`
	ProcessingMode *ID `hl7:"false,Processing Mode"`
}

// PTA represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type PTA struct {
	PolicyType  *IS `hl7:"false,Policy Type"`
	AmountClass *IS `hl7:"false,Amount Class"`
	Amount      *NM `hl7:"false,Amount"`
}

// QIP represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type QIP struct {
	SegmentFieldName   *ST `hl7:"false,Segment Field Name"`
	Value1Value2Value3 *ST `hl7:"false,Value1&Value2&Value3"`
}

// QSC represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type QSC struct {
	SegmentFieldName      *ST `hl7:"false,Segment Field Name"`
	RelationalOperator    *ID `hl7:"false,Relational Operator"`
	Value                 *ST `hl7:"false,Value"`
	RelationalConjunction *ID `hl7:"false,Relational Conjunction"`
}

// RCD represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type RCD struct {
	SegmentFieldName   *ST `hl7:"false,Segment Field Name"`
	HL7DateType        *ST `hl7:"false,HL7 Date Type"`
	MaximumColumnWidth *NM `hl7:"false,Maximum Column Width"`
}

// RFR represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type RFR struct {
	NumericRange      *NR `hl7:"false,Numeric Range"`
	AdministrativeSex *IS `hl7:"false,Administrative Sex"`
	AgeRange          *NR `hl7:"false,Age Range"`
	GestationalRange  *NR `hl7:"false,Gestational Range"`
	Species           *TX `hl7:"false,Species"`
	RaceSubspecies    *ST `hl7:"false,Race/Subspecies"`
	Conditions        *TX `hl7:"false,Conditions"`
}

// RI represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type RI struct {
	RepeatPattern        *IS `hl7:"false,Repeat Pattern"`
	ExplicitTimeInterval *ST `hl7:"false,Explicit Time Interval"`
}

// RMC represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type RMC struct {
	RoomType       *IS `hl7:"false,Room Type"`
	AmountType     *IS `hl7:"false,Amount Type"`
	CoverageAmount *NM `hl7:"false,Coverage Amount"`
}

// RP represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type RP struct {
	Pointer       *ST `hl7:"false,Pointer"`
	ApplicationID *HD `hl7:"false,Application ID"`
	TypeOfData    *ID `hl7:"false,Type Of Data"`
	Subtype       *ID `hl7:"false,Subtype"`
}

// SAD represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type SAD struct {
	StreetOrMailingAddress *ST `hl7:"false,Street Or Mailing Address"`
	StreetName             *ST `hl7:"false,Street Name"`
	DwellingNumber         *ST `hl7:"false,Dwelling Number"`
}

// SCV represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type SCV struct {
	ParameterClass *IS `hl7:"false,Parameter Class"`
	ParameterValue *ST `hl7:"false,Parameter Value"`
}

// SN represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type SN struct {
	Comparator      *ST `hl7:"false,Comparator"`
	Num1            *NM `hl7:"false,Num1"`
	SeparatorSuffix *ST `hl7:"false,Separator/Suffix"`
	Num2            *NM `hl7:"false,Num2"`
}

// SPD represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type SPD struct {
	SpecialtyName       *ST `hl7:"false,Specialty Name"`
	GoverningBoard      *ST `hl7:"false,Governing Board"`
	EligibleOrCertified *ID `hl7:"false,Eligible Or Certified"`
	DateOfCertification *DT `hl7:"false,Date Of Certification"`
}

// SPS represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type SPS struct {
	SpecimenSourceNameOrCode     *CE `hl7:"false,Specimen Source Name Or Code"`
	Additives                    *TX `hl7:"false,Additives"`
	Freetext                     *TX `hl7:"false,Freetext"`
	BodySite                     *CE `hl7:"false,Body Site"`
	SiteModifier                 *CE `hl7:"false,Site Modifier"`
	CollectionModifierMethodCode *CE `hl7:"false,Collection Modifier Method Code"`
	SpecimenRole                 *CE `hl7:"false,Specimen Role"`
}

// SRT represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type SRT struct {
	SortByField *ST `hl7:"false,Sort-By Field"`
	Sequencing  *ID `hl7:"false,Sequencing"`
}

// TQ represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type TQ struct {
	Quantity             *CQ  `hl7:"false,Quantity"`
	Interval             *RI  `hl7:"false,Interval"`
	Duration             *ST  `hl7:"false,Duration"`
	StartDateTime        *TS  `hl7:"false,Start Date/Time"`
	EndDateTime          *TS  `hl7:"false,End Date/Time"`
	Priority             *ST  `hl7:"false,Priority"`
	Condition            *ST  `hl7:"false,Condition"`
	Text                 *TX  `hl7:"false,Text"`
	ConjunctionComponent *ID  `hl7:"false,Conjunction Component"`
	OrderSequencing      *OSD `hl7:"false,Order Sequencing"`
	OccurrenceDuration   *CE  `hl7:"false,Occurrence Duration"`
	TotalOccurrences     *NM  `hl7:"false,Total Occurrences"`
}

// UVC represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type UVC struct {
	ValueCode   *IS `hl7:"false,Value Code"`
	ValueAmount *NM `hl7:"false,Value Amount"`
}

// VH represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type VH struct {
	StartDayRange  *ID `hl7:"false,Start Day Range"`
	EndDayRange    *ID `hl7:"false,End Day Range"`
	StartHourRange *TM `hl7:"false,Start Hour Range"`
	EndHourRange   *TM `hl7:"false,End Hour Range"`
}

// VID represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type VID struct {
	VersionID                *ID `hl7:"false,Version ID"`
	InternationalizationCode *CE `hl7:"false,Internationalization Code"`
	InternationalVersionID   *CE `hl7:"false,International Version ID"`
}

// VR represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type VR struct {
	FirstDataCodeValue *ST `hl7:"false,First Data Code Value"`
	LastDataCodeCalue  *ST `hl7:"false,Last Data Code Calue"`
}

// WVI represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type WVI struct {
	ChannelNumber *NM `hl7:"false,Channel Number"`
	ChannelName   *ST `hl7:"false,Channel Name"`
}

// WVS represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type WVS struct {
	SourceName1 *ST `hl7:"false,Source Name 1"`
	SourceName2 *ST `hl7:"false,Source Name 2"`
}

// XAD represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type XAD struct {
	StreetAddress              *SAD `hl7:"false,Street Address"`
	OtherDesignation           *ST  `hl7:"false,Other Designation"`
	City                       *ST  `hl7:"false,City"`
	StateOrProvince            *ST  `hl7:"false,State Or Province"`
	ZipOrPostalCode            *ST  `hl7:"false,Zip Or Postal Code"`
	Country                    *ID  `hl7:"false,Country"`
	AddressType                *ID  `hl7:"false,Address Type"`
	OtherGeographicDesignation *ST  `hl7:"false,Other Geographic Designation"`
	CountyParishCode           *IS  `hl7:"false,County/Parish Code"`
	CensusTract                *IS  `hl7:"false,Census Tract"`
	AddressRepresentationCode  *ID  `hl7:"false,Address Representation Code"`
	AddressValidityRange       *DR  `hl7:"false,Address Validity Range"`
}

// XCN represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type XCN struct {
	IDNumber                                    *ST `hl7:"false,ID Number"`
	FamilyName                                  *FN `hl7:"false,Family Name"`
	GivenName                                   *ST `hl7:"false,Given Name"`
	SecondAndFurtherGivenNamesOrInitialsThereof *ST `hl7:"false,Second And Further Given Names Or Initials Thereof"`
	SuffixEGJROrIII                             *ST `hl7:"false,Suffix (E.G., JR Or III)"`
	PrefixEGDR                                  *ST `hl7:"false,Prefix (E.G., DR)"`
	DegreeEGMD                                  *IS `hl7:"false,Degree (E.G., MD)"`
	SourceTable                                 *IS `hl7:"false,Source Table"`
	AssigningAuthority                          *HD `hl7:"false,Assigning Authority"`
	NameTypeCode                                *ID `hl7:"false,Name Type Code"`
	IdentifierCheckDigit                        *ST `hl7:"false,Identifier Check Digit"`
	CodeIdentifyingTheCheckDigitSchemeEmployed  *ID `hl7:"false,Code Identifying The Check Digit Scheme Employed"`
	IdentifierTypeCode                          *IS `hl7:"false,Identifier Type Code"`
	AssigningFacility                           *HD `hl7:"false,Assigning Facility"`
	NameRepresentationCode                      *ID `hl7:"false,Name Representation Code"`
	NameContext                                 *CE `hl7:"false,Name Context"`
	NameValidityRange                           *DR `hl7:"false,Name Validity Range"`
	NameAssemblyOrder                           *ID `hl7:"false,Name Assembly Order"`
}

// XON represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type XON struct {
	OrganizationName                           *ST `hl7:"false,Organization Name"`
	OrganizationNameTypeCode                   *IS `hl7:"false,Organization Name Type Code"`
	IDNumber                                   *NM `hl7:"false,ID Number"`
	CheckDigit                                 *NM `hl7:"false,Check Digit"`
	CodeIdentifyingTheCheckDigitSchemeEmployed *ID `hl7:"false,Code Identifying The Check Digit Scheme Employed"`
	AssigningAuthority                         *HD `hl7:"false,Assigning Authority"`
	IdentifierTypeCode                         *IS `hl7:"false,Identifier Type Code"`
	AssigningFacilityID                        *HD `hl7:"false,Assigning Facility ID"`
	NameRepresentationCode                     *ID `hl7:"false,Name Representation Code"`
}

// XPN represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type XPN struct {
	FamilyName                                  *FN `hl7:"false,Family Name"`
	GivenName                                   *ST `hl7:"false,Given Name"`
	SecondAndFurtherGivenNamesOrInitialsThereof *ST `hl7:"false,Second And Further Given Names Or Initials Thereof"`
	SuffixEGJROrIII                             *ST `hl7:"false,Suffix (E.G., JR Or III)"`
	PrefixEGDR                                  *ST `hl7:"false,Prefix (E.G., DR)"`
	DegreeEGMD                                  *IS `hl7:"false,Degree (E.G., MD)"`
	NameTypeCode                                *ID `hl7:"false,Name Type Code"`
	NameRepresentationCode                      *ID `hl7:"false,Name Representation Code"`
	NameContext                                 *CE `hl7:"false,Name Context"`
	NameValidityRange                           *DR `hl7:"false,Name Validity Range"`
	NameAssemblyOrder                           *ID `hl7:"false,Name Assembly Order"`
}

// XTN represents the corresponding HL7 datatype.
// Definition from HL7 2.4
type XTN struct {
	Number                         *TN `hl7:"false,Number"`
	TelecommunicationUseCode       *ID `hl7:"false,Telecommunication Use Code"`
	TelecommunicationEquipmentType *ID `hl7:"false,Telecommunication Equipment Type"`
	EmailAddress                   *ST `hl7:"false,Email Address"`
	CountryCode                    *NM `hl7:"false,Country Code"`
	AreaCityCode                   *NM `hl7:"false,Area/City Code"`
	PhoneNumber                    *NM `hl7:"false,Phone Number"`
	Extension                      *NM `hl7:"false,Extension"`
	AnyText                        *ST `hl7:"false,Any Text"`
}

// ABS represents the corresponding HL7 segment.
// Definition from HL7 2.4
type ABS struct {
	DischargeCareProvider      *XCN `hl7:"false,Discharge Care Provider"`       // ABS-1
	TransferMedicalServiceCode *CE  `hl7:"false,Transfer Medical Service Code"` // ABS-2
	SeverityOfIllnessCode      *CE  `hl7:"false,Severity Of Illness Code"`      // ABS-3
	DateTimeOfAttestation      *TS  `hl7:"false,Date/Time Of Attestation"`      // ABS-4
	AttestedBy                 *XCN `hl7:"false,Attested By"`                   // ABS-5
	TriageCode                 *CE  `hl7:"false,Triage Code"`                   // ABS-6
	AbstractCompletionDateTime *TS  `hl7:"false,Abstract Completion Date/Time"` // ABS-7
	AbstractedBy               *XCN `hl7:"false,Abstracted By"`                 // ABS-8
	CaseCategoryCode           *CE  `hl7:"false,Case Category Code"`            // ABS-9
	CaesarianSectionIndicator  *ID  `hl7:"false,Caesarian Section Indicator"`   // ABS-10
	GestationCategoryCode      *CE  `hl7:"false,Gestation Category Code"`       // ABS-11
	GestationPeriodWeeks       *NM  `hl7:"false,Gestation Period - Weeks"`      // ABS-12
	NewbornCode                *CE  `hl7:"false,Newborn Code"`                  // ABS-13
	StillbornIndicator         *ID  `hl7:"false,Stillborn Indicator"`           // ABS-14
}

func (s *ABS) SegmentName() string {
	return "ABS"
}

// ACC represents the corresponding HL7 segment.
// Definition from HL7 2.4
type ACC struct {
	AccidentDateTime            *TS  `hl7:"false,Accident Date/Time"`             // ACC-1
	AccidentCode                *CE  `hl7:"false,Accident Code"`                  // ACC-2
	AccidentLocation            *ST  `hl7:"false,Accident Location"`              // ACC-3
	AutoAccidentState           *CE  `hl7:"false,Auto Accident State"`            // ACC-4
	AccidentJobRelatedIndicator *ID  `hl7:"false,Accident Job Related Indicator"` // ACC-5
	AccidentDeathIndicator      *ID  `hl7:"false,Accident Death Indicator"`       // ACC-6
	EnteredBy                   *XCN `hl7:"false,Entered By"`                     // ACC-7
	AccidentDescription         *ST  `hl7:"false,Accident Description"`           // ACC-8
	BroughtInBy                 *ST  `hl7:"false,Brought In By"`                  // ACC-9
	PoliceNotifiedIndicator     *ID  `hl7:"false,Police Notified Indicator"`      // ACC-10
}

func (s *ACC) SegmentName() string {
	return "ACC"
}

// ADD represents the corresponding HL7 segment.
// Definition from HL7 2.4
type ADD struct {
	AddendumContinuationPointer *ST `hl7:"false,Addendum Continuation Pointer"` // ADD-1
}

func (s *ADD) SegmentName() string {
	return "ADD"
}

// AFF represents the corresponding HL7 segment.
// Definition from HL7 2.4
type AFF struct {
	SetIDAFF                                     *SI  `hl7:"true,Set ID - AFF"`                                      // AFF-1
	ProfessionalOrganization                     *XON `hl7:"true,Professional Organization"`                         // AFF-2
	ProfessionalOrganizationAddress              *XAD `hl7:"false,Professional Organization Address"`                // AFF-3
	ProfessionalOrganizationAffiliationDateRange []DR `hl7:"false,Professional Organization Affiliation Date Range"` // AFF-4
	ProfessionalAffiliationAdditionalInformation *ST  `hl7:"false,Professional Affiliation Additional Information"`  // AFF-5
}

func (s *AFF) SegmentName() string {
	return "AFF"
}

// AIG represents the corresponding HL7 segment.
// Definition from HL7 2.4
type AIG struct {
	SetIDAIG                 *SI  `hl7:"true,Set ID - AIG"`                  // AIG-1
	SegmentActionCode        *ID  `hl7:"false,Segment Action Code"`          // AIG-2
	ResourceID               *CE  `hl7:"false,Resource ID"`                  // AIG-3
	ResourceType             *CE  `hl7:"true,Resource Type"`                 // AIG-4
	ResourceGroup            []CE `hl7:"false,Resource Group"`               // AIG-5
	ResourceQuantity         *NM  `hl7:"false,Resource Quantity"`            // AIG-6
	ResourceQuantityUnits    *CE  `hl7:"false,Resource Quantity Units"`      // AIG-7
	StartDateTime            *TS  `hl7:"false,Start Date/Time"`              // AIG-8
	StartDateTimeOffset      *NM  `hl7:"false,Start Date/Time Offset"`       // AIG-9
	StartDateTimeOffsetUnits *CE  `hl7:"false,Start Date/Time Offset Units"` // AIG-10
	Duration                 *NM  `hl7:"false,Duration"`                     // AIG-11
	DurationUnits            *CE  `hl7:"false,Duration Units"`               // AIG-12
	AllowSubstitutionCode    *IS  `hl7:"false,Allow Substitution Code"`      // AIG-13
	FillerStatusCode         *CE  `hl7:"false,Filler Status Code"`           // AIG-14
}

func (s *AIG) SegmentName() string {
	return "AIG"
}

// AIL represents the corresponding HL7 segment.
// Definition from HL7 2.4
type AIL struct {
	SetIDAIL                 *SI `hl7:"true,Set ID - AIL"`                  // AIL-1
	SegmentActionCode        *ID `hl7:"false,Segment Action Code"`          // AIL-2
	LocationResourceID       *PL `hl7:"false,Location Resource ID"`         // AIL-3
	LocationTypeAIL          *CE `hl7:"true,Location Type-AIL"`             // AIL-4
	LocationGroup            *CE `hl7:"false,Location Group"`               // AIL-5
	StartDateTime            *TS `hl7:"false,Start Date/Time"`              // AIL-6
	StartDateTimeOffset      *NM `hl7:"false,Start Date/Time Offset"`       // AIL-7
	StartDateTimeOffsetUnits *CE `hl7:"false,Start Date/Time Offset Units"` // AIL-8
	Duration                 *NM `hl7:"false,Duration"`                     // AIL-9
	DurationUnits            *CE `hl7:"false,Duration Units"`               // AIL-10
	AllowSubstitutionCode    *IS `hl7:"false,Allow Substitution Code"`      // AIL-11
	FillerStatusCode         *CE `hl7:"false,Filler Status Code"`           // AIL-12
}

func (s *AIL) SegmentName() string {
	return "AIL"
}

// AIP represents the corresponding HL7 segment.
// Definition from HL7 2.4
type AIP struct {
	SetIDAIP                 *SI   `hl7:"true,Set ID - AIP"`                  // AIP-1
	SegmentActionCode        *ID   `hl7:"false,Segment Action Code"`          // AIP-2
	PersonnelResourceID      []XCN `hl7:"false,Personnel Resource ID"`        // AIP-3
	ResourceRole             *CE   `hl7:"true,Resource Role"`                 // AIP-4
	ResourceGroup            *CE   `hl7:"false,Resource Group"`               // AIP-5
	StartDateTime            *TS   `hl7:"false,Start Date/Time"`              // AIP-6
	StartDateTimeOffset      *NM   `hl7:"false,Start Date/Time Offset"`       // AIP-7
	StartDateTimeOffsetUnits *CE   `hl7:"false,Start Date/Time Offset Units"` // AIP-8
	Duration                 *NM   `hl7:"false,Duration"`                     // AIP-9
	DurationUnits            *CE   `hl7:"false,Duration Units"`               // AIP-10
	AllowSubstitutionCode    *IS   `hl7:"false,Allow Substitution Code"`      // AIP-11
	FillerStatusCode         *CE   `hl7:"false,Filler Status Code"`           // AIP-12
}

func (s *AIP) SegmentName() string {
	return "AIP"
}

// AIS represents the corresponding HL7 segment.
// Definition from HL7 2.4
type AIS struct {
	SetIDAIS                             *SI  `hl7:"true,Set ID - AIS"`                             // AIS-1
	SegmentActionCode                    *ID  `hl7:"false,Segment Action Code"`                     // AIS-2
	UniversalServiceIdentifier           *CE  `hl7:"true,Universal Service Identifier"`             // AIS-3
	StartDateTime                        *TS  `hl7:"false,Start Date/Time"`                         // AIS-4
	StartDateTimeOffset                  *NM  `hl7:"false,Start Date/Time Offset"`                  // AIS-5
	StartDateTimeOffsetUnits             *CE  `hl7:"false,Start Date/Time Offset Units"`            // AIS-6
	Duration                             *NM  `hl7:"false,Duration"`                                // AIS-7
	DurationUnits                        *CE  `hl7:"false,Duration Units"`                          // AIS-8
	AllowSubstitutionCode                *IS  `hl7:"false,Allow Substitution Code"`                 // AIS-9
	FillerStatusCode                     *CE  `hl7:"false,Filler Status Code"`                      // AIS-10
	PlacerSupplementalServiceInformation []CE `hl7:"false,Placer Supplemental Service Information"` // AIS-11
	FillerSupplementalServiceInformation []CE `hl7:"false,Filler Supplemental Service Information"` // AIS-12
}

func (s *AIS) SegmentName() string {
	return "AIS"
}

// AL1 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type AL1 struct {
	SetIDAL1                        *CE  `hl7:"true,Set ID - AL1"`                       // AL1-1
	AllergenTypeCode                *CE  `hl7:"false,Allergen Type Code"`                // AL1-2
	AllergenCodeMnemonicDescription *CE  `hl7:"true,Allergen Code/Mnemonic/Description"` // AL1-3
	AllergySeverityCode             *CE  `hl7:"false,Allergy Severity Code"`             // AL1-4
	AllergyReactionCode             []ST `hl7:"false,Allergy Reaction Code"`             // AL1-5
	IdentificationDate              *DT  `hl7:"false,Identification Date"`               // AL1-6
}

func (s *AL1) SegmentName() string {
	return "AL1"
}

// APR represents the corresponding HL7 segment.
// Definition from HL7 2.4
type APR struct {
	TimeSelectionCriteria     []SCV `hl7:"false,Time Selection Criteria"`     // APR-1
	ResourceSelectionCriteria []SCV `hl7:"false,Resource Selection Criteria"` // APR-2
	LocationSelectionCriteria []SCV `hl7:"false,Location Selection Criteria"` // APR-3
	SlotSpacingCriteria       *NM   `hl7:"false,Slot Spacing Criteria"`       // APR-4
	FillerOverrideCriteria    []SCV `hl7:"false,Filler Override Criteria"`    // APR-5
}

func (s *APR) SegmentName() string {
	return "APR"
}

// ARQ represents the corresponding HL7 segment.
// Definition from HL7 2.4
type ARQ struct {
	PlacerAppointmentID         *EI   `hl7:"true,Placer Appointment ID"`            // ARQ-1
	FillerAppointmentID         *EI   `hl7:"false,Filler Appointment ID"`           // ARQ-2
	OccurrenceNumber            *NM   `hl7:"false,Occurrence Number"`               // ARQ-3
	PlacerGroupNumber           *EI   `hl7:"false,Placer Group Number"`             // ARQ-4
	ScheduleID                  *CE   `hl7:"false,Schedule ID"`                     // ARQ-5
	RequestEventReason          *CE   `hl7:"false,Request Event Reason"`            // ARQ-6
	AppointmentReason           *CE   `hl7:"false,Appointment Reason"`              // ARQ-7
	AppointmentType             *CE   `hl7:"false,Appointment Type"`                // ARQ-8
	AppointmentDuration         *NM   `hl7:"false,Appointment Duration"`            // ARQ-9
	AppointmentDurationUnits    *CE   `hl7:"false,Appointment Duration Units"`      // ARQ-10
	RequestedStartDateTimeRange []DR  `hl7:"false,Requested Start Date/Time Range"` // ARQ-11
	PriorityARQ                 *ST   `hl7:"false,Priority-ARQ"`                    // ARQ-12
	RepeatingInterval           *RI   `hl7:"false,Repeating Interval"`              // ARQ-13
	RepeatingIntervalDuration   *ST   `hl7:"false,Repeating Interval Duration"`     // ARQ-14
	PlacerContactPerson         []XCN `hl7:"true,Placer Contact Person"`            // ARQ-15
	PlacerContactPhoneNumber    []XTN `hl7:"false,Placer Contact Phone Number"`     // ARQ-16
	PlacerContactAddress        []XAD `hl7:"false,Placer Contact Address"`          // ARQ-17
	PlacerContactLocation       *PL   `hl7:"false,Placer Contact Location"`         // ARQ-18
	EnteredByPerson             []XCN `hl7:"true,Entered By Person"`                // ARQ-19
	EnteredByPhoneNumber        []XTN `hl7:"false,Entered By Phone Number"`         // ARQ-20
	EnteredByLocation           *PL   `hl7:"false,Entered By Location"`             // ARQ-21
	ParentPlacerAppointmentID   *EI   `hl7:"false,Parent Placer Appointment ID"`    // ARQ-22
	ParentFillerAppointmentID   *EI   `hl7:"false,Parent Filler Appointment ID"`    // ARQ-23
	PlacerOrderNumber           []EI  `hl7:"false,Placer Order Number"`             // ARQ-24
	FillerOrderNumber           []EI  `hl7:"false,Filler Order Number"`             // ARQ-25
}

func (s *ARQ) SegmentName() string {
	return "ARQ"
}

// AUT represents the corresponding HL7 segment.
// Definition from HL7 2.4
type AUT struct {
	AuthorizingPayorPlanID       *CE `hl7:"false,Authorizing Payor, Plan ID"`      // AUT-1
	AuthorizingPayorCompanyID    *CE `hl7:"true,Authorizing Payor, Company ID"`    // AUT-2
	AuthorizingPayorCompanyName  *ST `hl7:"false,Authorizing Payor, Company Name"` // AUT-3
	AuthorizationEffectiveDate   *TS `hl7:"false,Authorization Effective Date"`    // AUT-4
	AuthorizationExpirationDate  *TS `hl7:"false,Authorization Expiration Date"`   // AUT-5
	AuthorizationIdentifier      *EI `hl7:"false,Authorization Identifier"`        // AUT-6
	ReimbursementLimit           *CP `hl7:"false,Reimbursement Limit"`             // AUT-7
	RequestedNumberOfTreatments  *NM `hl7:"false,Requested Number Of Treatments"`  // AUT-8
	AuthorizedNumberOfTreatments *NM `hl7:"false,Authorized Number Of Treatments"` // AUT-9
	ProcessDate                  *TS `hl7:"false,Process Date"`                    // AUT-10
}

func (s *AUT) SegmentName() string {
	return "AUT"
}

// BHS represents the corresponding HL7 segment.
// Definition from HL7 2.4
type BHS struct {
	BatchFieldSeparator       *ST `hl7:"true,Batch Field Separator"`        // BHS-1
	BatchEncodingCharacters   *ST `hl7:"true,Batch Encoding Characters"`    // BHS-2
	BatchSendingApplication   *ST `hl7:"false,Batch Sending Application"`   // BHS-3
	BatchSendingFacility      *ST `hl7:"false,Batch Sending Facility"`      // BHS-4
	BatchReceivingApplication *ST `hl7:"false,Batch Receiving Application"` // BHS-5
	BatchReceivingFacility    *ST `hl7:"false,Batch Receiving Facility"`    // BHS-6
	BatchCreationDateTime     *TS `hl7:"false,Batch Creation Date/Time"`    // BHS-7
	BatchSecurity             *ST `hl7:"false,Batch Security"`              // BHS-8
	BatchNameIDType           *ST `hl7:"false,Batch Name/ID/Type"`          // BHS-9
	BatchComment              *ST `hl7:"false,Batch Comment"`               // BHS-10
	BatchControlID            *ST `hl7:"false,Batch Control ID"`            // BHS-11
	ReferenceBatchControlID   *ST `hl7:"false,Reference Batch Control ID"`  // BHS-12
}

func (s *BHS) SegmentName() string {
	return "BHS"
}

// BLC represents the corresponding HL7 segment.
// Definition from HL7 2.4
type BLC struct {
	BloodProductCode *CE `hl7:"false,Blood Product Code"` // BLC-1
	BloodAmount      *CQ `hl7:"false,Blood Amount"`       // BLC-2
}

func (s *BLC) SegmentName() string {
	return "BLC"
}

// BLG represents the corresponding HL7 segment.
// Definition from HL7 2.4
type BLG struct {
	WhenToCharge *CCD `hl7:"false,When To Charge"` // BLG-1
	ChargeType   *ID  `hl7:"false,Charge Type"`    // BLG-2
	AccountID    *CX  `hl7:"false,Account ID"`     // BLG-3
}

func (s *BLG) SegmentName() string {
	return "BLG"
}

// BTS represents the corresponding HL7 segment.
// Definition from HL7 2.4
type BTS struct {
	BatchMessageCount *ST  `hl7:"false,Batch Message Count"` // BTS-1
	BatchComment      *ST  `hl7:"false,Batch Comment"`       // BTS-2
	BatchTotals       []NM `hl7:"false,Batch Totals"`        // BTS-3
}

func (s *BTS) SegmentName() string {
	return "BTS"
}

// CDM represents the corresponding HL7 segment.
// Definition from HL7 2.4
type CDM struct {
	PrimaryKeyValueCDM           *CE   `hl7:"true,Primary Key Value - CDM"`         // CDM-1
	ChargeCodeAlias              []CE  `hl7:"false,Charge Code Alias"`              // CDM-2
	ChargeDescriptionShort       *ST   `hl7:"true,Charge Description Short"`        // CDM-3
	ChargeDescriptionLong        *ST   `hl7:"false,Charge Description Long"`        // CDM-4
	DescriptionOverrideIndicator *IS   `hl7:"false,Description Override Indicator"` // CDM-5
	ExplodingCharges             []CE  `hl7:"false,Exploding Charges"`              // CDM-6
	ProcedureCode                []CE  `hl7:"false,Procedure Code"`                 // CDM-7
	ActiveInactiveFlag           *ID   `hl7:"false,Active/Inactive Flag"`           // CDM-8
	InventoryNumber              []CE  `hl7:"false,Inventory Number"`               // CDM-9
	ResourceLoad                 *NM   `hl7:"false,Resource Load"`                  // CDM-10
	ContractNumber               []CK  `hl7:"false,Contract Number"`                // CDM-11
	ContractOrganization         []XON `hl7:"false,Contract Organization"`          // CDM-12
	RoomFeeIndicator             *ID   `hl7:"false,Room Fee Indicator"`             // CDM-13
}

func (s *CDM) SegmentName() string {
	return "CDM"
}

// CM0 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type CM0 struct {
	SetIDCM0                *SI   `hl7:"false,Set ID - CM0"`               // CM0-1
	SponsorStudyID          *EI   `hl7:"true,Sponsor Study ID"`            // CM0-2
	AlternateStudyID        []EI  `hl7:"false,Alternate Study ID"`         // CM0-3
	TitleOfStudy            *ST   `hl7:"true,Title Of Study"`              // CM0-4
	ChairmanOfStudy         []XCN `hl7:"false,Chairman Of Study"`          // CM0-5
	LastIRBApprovalDate     *DT   `hl7:"false,Last IRB Approval Date"`     // CM0-6
	TotalAccrualToDate      *NM   `hl7:"false,Total Accrual To Date"`      // CM0-7
	LastAccrualDate         *DT   `hl7:"false,Last Accrual Date"`          // CM0-8
	ContactForStudy         []XCN `hl7:"false,Contact For Study"`          // CM0-9
	ContactSTelephoneNumber *XTN  `hl7:"false,Contact'S Telephone Number"` // CM0-10
	ContactSAddress         []XAD `hl7:"false,Contact'S Address"`          // CM0-11
}

func (s *CM0) SegmentName() string {
	return "CM0"
}

// CM1 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type CM1 struct {
	SetIDCM1                *SI `hl7:"true,Set ID - CM1"`               // CM1-1
	StudyPhaseIdentifier    *CE `hl7:"true,Study Phase Identifier"`     // CM1-2
	DescriptionOfStudyPhase *ST `hl7:"true,Description Of Study Phase"` // CM1-3
}

func (s *CM1) SegmentName() string {
	return "CM1"
}

// CM2 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type CM2 struct {
	SetIDCM2                     *SI  `hl7:"false,Set ID- CM2"`                     // CM2-1
	ScheduledTimePoint           *CE  `hl7:"true,Scheduled Time Point"`             // CM2-2
	DescriptionOfTimePoint       *ST  `hl7:"false,Description Of Time Point"`       // CM2-3
	EventsScheduledThisTimePoint []CE `hl7:"true,Events Scheduled This Time Point"` // CM2-4
}

func (s *CM2) SegmentName() string {
	return "CM2"
}

// CNS represents the corresponding HL7 segment.
// Definition from HL7 2.4
type CNS struct {
	StartingNotificationReferenceNumber *NM `hl7:"false,Starting Notification Reference Number"` // CNS-1
	EndingNotificationReferenceNumber   *NM `hl7:"false,Ending Notification Reference Number"`   // CNS-2
	StartingNotificationDateTime        *TS `hl7:"false,Starting Notification Date/Time"`        // CNS-3
	EndingNotificationDateTime          *TS `hl7:"false,Ending Notification Date/Time"`          // CNS-4
	StartingNotificationCode            *CE `hl7:"false,Starting Notification Code"`             // CNS-5
	EndingNotificationCode              *CE `hl7:"false,Ending Notification Code"`               // CNS-6
}

func (s *CNS) SegmentName() string {
	return "CNS"
}

// CSP represents the corresponding HL7 segment.
// Definition from HL7 2.4
type CSP struct {
	StudyPhaseIdentifier    *CE `hl7:"true,Study Phase Identifier"`       // CSP-1
	DateTimeStudyPhaseBegan *TS `hl7:"true,Date/Time Study Phase Began"`  // CSP-2
	DateTimeStudyPhaseEnded *TS `hl7:"false,Date/Time Study Phase Ended"` // CSP-3
	StudyPhaseEvaluability  *CE `hl7:"false,Study Phase Evaluability"`    // CSP-4
}

func (s *CSP) SegmentName() string {
	return "CSP"
}

// CSR represents the corresponding HL7 segment.
// Definition from HL7 2.4
type CSR struct {
	SponsorStudyID                     *EI   `hl7:"true,Sponsor Study ID"`                        // CSR-1
	AlternateStudyID                   *EI   `hl7:"false,Alternate Study ID"`                     // CSR-2
	InstitutionRegisteringThePatient   *CE   `hl7:"false,Institution Registering The Patient"`    // CSR-3
	SponsorPatientID                   *CX   `hl7:"true,Sponsor Patient ID"`                      // CSR-4
	AlternatePatientIDCSR              *CX   `hl7:"false,Alternate Patient ID - CSR"`             // CSR-5
	DateTimeOfPatientStudyRegistration *TS   `hl7:"true,Date/Time Of Patient Study Registration"` // CSR-6
	PersonPerformingStudyRegistration  []XCN `hl7:"false,Person Performing Study Registration"`   // CSR-7
	StudyAuthorizingProvider           []XCN `hl7:"true,Study Authorizing Provider"`              // CSR-8
	DateTimePatientStudyConsentSigned  *TS   `hl7:"false,Date/Time Patient Study Consent Signed"` // CSR-9
	PatientStudyEligibilityStatus      *CE   `hl7:"false,Patient Study Eligibility Status"`       // CSR-10
	StudyRandomizationDateTime         []TS  `hl7:"false,Study Randomization Date/Time"`          // CSR-11
	RandomizedStudyArm                 []CE  `hl7:"false,Randomized Study Arm"`                   // CSR-12
	StratumForStudyRandomization       []CE  `hl7:"false,Stratum For Study Randomization"`        // CSR-13
	PatientEvaluabilityStatus          *CE   `hl7:"false,Patient Evaluability Status"`            // CSR-14
	DateTimeEndedStudy                 *TS   `hl7:"false,Date/Time Ended Study"`                  // CSR-15
	ReasonEndedStudy                   *CE   `hl7:"false,Reason Ended Study"`                     // CSR-16
}

func (s *CSR) SegmentName() string {
	return "CSR"
}

// CSS represents the corresponding HL7 segment.
// Definition from HL7 2.4
type CSS struct {
	StudyScheduledTimePoint        *CE  `hl7:"true,Study Scheduled Time Point"`          // CSS-1
	StudyScheduledPatientTimePoint *TS  `hl7:"false,Study Scheduled Patient Time Point"` // CSS-2
	StudyQualityControlCodes       []CE `hl7:"false,Study Quality Control Codes"`        // CSS-3
}

func (s *CSS) SegmentName() string {
	return "CSS"
}

// CTD represents the corresponding HL7 segment.
// Definition from HL7 2.4
type CTD struct {
	ContactRole                     []CE  `hl7:"true,Contact Role"`                       // CTD-1
	ContactName                     []XPN `hl7:"false,Contact Name"`                      // CTD-2
	ContactAddress                  []XAD `hl7:"false,Contact Address"`                   // CTD-3
	ContactLocation                 *PL   `hl7:"false,Contact Location"`                  // CTD-4
	ContactCommunicationInformation []XTN `hl7:"false,Contact Communication Information"` // CTD-5
	PreferredMethodOfContact        *CE   `hl7:"false,Preferred Method Of Contact"`       // CTD-6
	ContactIdentifiers              []PI  `hl7:"false,Contact Identifiers"`               // CTD-7
}

func (s *CTD) SegmentName() string {
	return "CTD"
}

// CTI represents the corresponding HL7 segment.
// Definition from HL7 2.4
type CTI struct {
	SponsorStudyID          *EI `hl7:"true,Sponsor Study ID"`            // CTI-1
	StudyPhaseIdentifier    *CE `hl7:"false,Study Phase Identifier"`     // CTI-2
	StudyScheduledTimePoint *CE `hl7:"false,Study Scheduled Time Point"` // CTI-3
}

func (s *CTI) SegmentName() string {
	return "CTI"
}

// DB1 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type DB1 struct {
	SetIDDB1                   *SI  `hl7:"true,Set ID - DB1"`                    // DB1-1
	DisabledPersonCode         *IS  `hl7:"false,Disabled Person Code"`           // DB1-2
	DisabledPersonIdentifier   []CX `hl7:"false,Disabled Person Identifier"`     // DB1-3
	DisabilityIndicator        *ID  `hl7:"false,Disability Indicator"`           // DB1-4
	DisabilityStartDate        *DT  `hl7:"false,Disability Start Date"`          // DB1-5
	DisabilityEndDate          *DT  `hl7:"false,Disability End Date"`            // DB1-6
	DisabilityReturnToWorkDate *DT  `hl7:"false,Disability Return To Work Date"` // DB1-7
	DisabilityUnableToWorkDate *DT  `hl7:"false,Disability Unable To Work Date"` // DB1-8
}

func (s *DB1) SegmentName() string {
	return "DB1"
}

// DG1 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type DG1 struct {
	SetIDDG1                *SI   `hl7:"true,Set ID - DG1"`               // DG1-1
	DiagnosisCodingMethod   *ID   `hl7:"false,Diagnosis Coding Method"`   // DG1-2
	DiagnosisCodeDG1        *CE   `hl7:"false,Diagnosis Code - DG1"`      // DG1-3
	DiagnosisDescription    *ST   `hl7:"false,Diagnosis Description"`     // DG1-4
	DiagnosisDateTime       *TS   `hl7:"false,Diagnosis Date/Time"`       // DG1-5
	DiagnosisType           *IS   `hl7:"true,Diagnosis Type"`             // DG1-6
	MajorDiagnosticCategory *CE   `hl7:"false,Major Diagnostic Category"` // DG1-7
	DiagnosticRelatedGroup  *CE   `hl7:"false,Diagnostic Related Group"`  // DG1-8
	DRGApprovalIndicator    *ID   `hl7:"false,DRG Approval Indicator"`    // DG1-9
	DRGGrouperReviewCode    *IS   `hl7:"false,DRG Grouper Review Code"`   // DG1-10
	OutlierType             *CE   `hl7:"false,Outlier Type"`              // DG1-11
	OutlierDays             *NM   `hl7:"false,Outlier Days"`              // DG1-12
	OutlierCost             *CP   `hl7:"false,Outlier Cost"`              // DG1-13
	GrouperVersionAndType   *ST   `hl7:"false,Grouper Version And Type"`  // DG1-14
	DiagnosisPriority       *ID   `hl7:"false,Diagnosis Priority"`        // DG1-15
	DiagnosingClinician     []XCN `hl7:"false,Diagnosing Clinician"`      // DG1-16
	DiagnosisClassification *IS   `hl7:"false,Diagnosis Classification"`  // DG1-17
	ConfidentialIndicator   *ID   `hl7:"false,Confidential Indicator"`    // DG1-18
	AttestationDateTime     *TS   `hl7:"false,Attestation Date/Time"`     // DG1-19
}

func (s *DG1) SegmentName() string {
	return "DG1"
}

// DRG represents the corresponding HL7 segment.
// Definition from HL7 2.4
type DRG struct {
	DiagnosticRelatedGroup *CE `hl7:"false,Diagnostic Related Group"` // DRG-1
	DRGAssignedDateTime    *TS `hl7:"false,DRG Assigned Date/Time"`   // DRG-2
	DRGApprovalIndicator   *ID `hl7:"false,DRG Approval Indicator"`   // DRG-3
	DRGGrouperReviewCode   *IS `hl7:"false,DRG Grouper Review Code"`  // DRG-4
	OutlierType            *CE `hl7:"false,Outlier Type"`             // DRG-5
	OutlierDays            *NM `hl7:"false,Outlier Days"`             // DRG-6
	OutlierCost            *CP `hl7:"false,Outlier Cost"`             // DRG-7
	DRGPayor               *IS `hl7:"false,DRG Payor"`                // DRG-8
	OutlierReimbursement   *CP `hl7:"false,Outlier Reimbursement"`    // DRG-9
	ConfidentialIndicator  *ID `hl7:"false,Confidential Indicator"`   // DRG-10
	DRGTransferType        *IS `hl7:"false,DRG Transfer Type"`        // DRG-11
}

func (s *DRG) SegmentName() string {
	return "DRG"
}

// DSC represents the corresponding HL7 segment.
// Definition from HL7 2.4
type DSC struct {
	ContinuationPointer *ST `hl7:"false,Continuation Pointer"` // DSC-1
	ContinuationStyle   *ID `hl7:"false,Continuation Style"`   // DSC-2
}

func (s *DSC) SegmentName() string {
	return "DSC"
}

// DSP represents the corresponding HL7 segment.
// Definition from HL7 2.4
type DSP struct {
	SetIDDSP          *SI `hl7:"false,Set ID - DSP"`        // DSP-1
	DisplayLevel      *SI `hl7:"false,Display Level"`       // DSP-2
	DataLine          *TX `hl7:"true,Data Line"`            // DSP-3
	LogicalBreakPoint *ST `hl7:"false,Logical Break Point"` // DSP-4
	ResultID          *TX `hl7:"false,Result ID"`           // DSP-5
}

func (s *DSP) SegmentName() string {
	return "DSP"
}

// ECD represents the corresponding HL7 segment.
// Definition from HL7 2.4
type ECD struct {
	ReferenceCommandNumber  *NM  `hl7:"true,Reference Command Number"`   // ECD-1
	RemoteControlCommand    *CE  `hl7:"true,Remote Control Command"`     // ECD-2
	ResponseRequired        *ID  `hl7:"false,Response Required"`         // ECD-3
	RequestedCompletionTime *TQ  `hl7:"false,Requested Completion Time"` // ECD-4
	Parameters              []ST `hl7:"false,Parameters"`                // ECD-5
}

func (s *ECD) SegmentName() string {
	return "ECD"
}

// ECR represents the corresponding HL7 segment.
// Definition from HL7 2.4
type ECR struct {
	CommandResponse           *CE  `hl7:"true,Command Response"`             // ECR-1
	DateTimeCompleted         *TS  `hl7:"true,Date/Time Completed"`          // ECR-2
	CommandResponseParameters []ST `hl7:"false,Command Response Parameters"` // ECR-3
}

func (s *ECR) SegmentName() string {
	return "ECR"
}

// EDU represents the corresponding HL7 segment.
// Definition from HL7 2.4
type EDU struct {
	SetIDEDU                                    *SI  `hl7:"true,Set ID - EDU"`                                      // EDU-1
	AcademicDegree                              *IS  `hl7:"false,Academic Degree"`                                  // EDU-2
	AcademicDegreeProgramDateRange              *DR  `hl7:"false,Academic Degree Program Date Range"`               // EDU-3
	AcademicDegreeProgramParticipationDateRange *DR  `hl7:"false,Academic Degree Program Participation Date Range"` // EDU-4
	AcademicDegreeGrantedDate                   *DT  `hl7:"false,Academic Degree Granted Date"`                     // EDU-5
	School                                      *XON `hl7:"false,School"`                                           // EDU-6
	SchoolTypeCode                              *CE  `hl7:"false,School Type Code"`                                 // EDU-7
	SchoolAddress                               *XAD `hl7:"false,School Address"`                                   // EDU-8
}

func (s *EDU) SegmentName() string {
	return "EDU"
}

// EQL represents the corresponding HL7 segment.
// Definition from HL7 2.4
type EQL struct {
	QueryTag                *ST `hl7:"false,Query Tag"`                 // EQL-1
	QueryResponseFormatCode *ID `hl7:"true,Query/Response Format Code"` // EQL-2
	EQLQueryName            *CE `hl7:"true,EQL Query Name"`             // EQL-3
	EQLQueryStatement       *ST `hl7:"true,EQL Query Statement"`        // EQL-4
}

func (s *EQL) SegmentName() string {
	return "EQL"
}

// EQP represents the corresponding HL7 segment.
// Definition from HL7 2.4
type EQP struct {
	EventType       *CE `hl7:"true,Event Type"`       // EQP-1
	FileName        *ST `hl7:"false,File Name"`       // EQP-2
	StartDateTime   *TS `hl7:"true,Start Date/Time"`  // EQP-3
	EndDateTime     *TS `hl7:"false,End Date/Time"`   // EQP-4
	TransactionData *FT `hl7:"true,Transaction Data"` // EQP-5
}

func (s *EQP) SegmentName() string {
	return "EQP"
}

// EQU represents the corresponding HL7 segment.
// Definition from HL7 2.4
type EQU struct {
	EquipmentInstanceIdentifier *EI `hl7:"true,Equipment Instance Identifier"` // EQU-1
	EventDateTime               *TS `hl7:"true,Event Date/Time"`               // EQU-2
	EquipmentState              *CE `hl7:"false,Equipment State"`              // EQU-3
	LocalRemoteControlState     *CE `hl7:"false,Local/Remote Control State"`   // EQU-4
	AlertLevel                  *CE `hl7:"false,Alert Level"`                  // EQU-5
}

func (s *EQU) SegmentName() string {
	return "EQU"
}

// ERQ represents the corresponding HL7 segment.
// Definition from HL7 2.4
type ERQ struct {
	QueryTag           *ST   `hl7:"false,Query Tag"`            // ERQ-1
	EventIdentifier    *CE   `hl7:"true,Event Identifier"`      // ERQ-2
	InputParameterList []QIP `hl7:"false,Input Parameter List"` // ERQ-3
}

func (s *ERQ) SegmentName() string {
	return "ERQ"
}

// ERR represents the corresponding HL7 segment.
// Definition from HL7 2.4
type ERR struct {
	ErrorCodeAndLocation []ELD `hl7:"true,Error Code And Location"` // ERR-1
}

func (s *ERR) SegmentName() string {
	return "ERR"
}

// EVN represents the corresponding HL7 segment.
// Definition from HL7 2.4
type EVN struct {
	EventTypeCode        *ID   `hl7:"false,Event Type Code"`         // EVN-1
	RecordedDateTime     *TS   `hl7:"true,Recorded Date/Time"`       // EVN-2
	DateTimePlannedEvent *TS   `hl7:"false,Date/Time Planned Event"` // EVN-3
	EventReasonCode      *IS   `hl7:"false,Event Reason Code"`       // EVN-4
	OperatorID           []XCN `hl7:"false,Operator ID"`             // EVN-5
	EventOccurred        *TS   `hl7:"false,Event Occurred"`          // EVN-6
	EventFacility        *HD   `hl7:"false,Event Facility"`          // EVN-7
}

func (s *EVN) SegmentName() string {
	return "EVN"
}

// FAC represents the corresponding HL7 segment.
// Definition from HL7 2.4
type FAC struct {
	FacilityIDFAC                       *EI   `hl7:"true,Facility ID-FAC"`                        // FAC-1
	FacilityType                        *ID   `hl7:"false,Facility Type"`                         // FAC-2
	FacilityAddress                     []XAD `hl7:"true,Facility Address"`                       // FAC-3
	FacilityTelecommunication           *XTN  `hl7:"true,Facility Telecommunication"`             // FAC-4
	ContactPerson                       []XCN `hl7:"false,Contact Person"`                        // FAC-5
	ContactTitle                        []ST  `hl7:"false,Contact Title"`                         // FAC-6
	ContactAddress                      []XAD `hl7:"false,Contact Address"`                       // FAC-7
	ContactTelecommunication            []XTN `hl7:"false,Contact Telecommunication"`             // FAC-8
	SignatureAuthority                  []XCN `hl7:"true,Signature Authority"`                    // FAC-9
	SignatureAuthorityTitle             *ST   `hl7:"false,Signature Authority Title"`             // FAC-10
	SignatureAuthorityAddress           []XAD `hl7:"false,Signature Authority Address"`           // FAC-11
	SignatureAuthorityTelecommunication *XTN  `hl7:"false,Signature Authority Telecommunication"` // FAC-12
}

func (s *FAC) SegmentName() string {
	return "FAC"
}

// FHS represents the corresponding HL7 segment.
// Definition from HL7 2.4
type FHS struct {
	FileFieldSeparator       *ST `hl7:"true,File Field Separator"`        // FHS-1
	FileEncodingCharacters   *ST `hl7:"true,File Encoding Characters"`    // FHS-2
	FileSendingApplication   *ST `hl7:"false,File Sending Application"`   // FHS-3
	FileSendingFacility      *ST `hl7:"false,File Sending Facility"`      // FHS-4
	FileReceivingApplication *ST `hl7:"false,File Receiving Application"` // FHS-5
	FileReceivingFacility    *ST `hl7:"false,File Receiving Facility"`    // FHS-6
	FileCreationDateTime     *TS `hl7:"false,File Creation Date/Time"`    // FHS-7
	FileSecurity             *ST `hl7:"false,File Security"`              // FHS-8
	FileNameID               *ST `hl7:"false,File Name/ID"`               // FHS-9
	FileHeaderComment        *ST `hl7:"false,File Header Comment"`        // FHS-10
	FileControlID            *ST `hl7:"false,File Control ID"`            // FHS-11
	ReferenceFileControlID   *ST `hl7:"false,Reference File Control ID"`  // FHS-12
}

func (s *FHS) SegmentName() string {
	return "FHS"
}

// FT1 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type FT1 struct {
	SetIDFT1                  *SI   `hl7:"false,Set ID - FT1"`                  // FT1-1
	TransactionID             *ST   `hl7:"false,Transaction ID"`                // FT1-2
	TransactionBatchID        *ST   `hl7:"false,Transaction Batch ID"`          // FT1-3
	TransactionDate           *TS   `hl7:"true,Transaction Date"`               // FT1-4
	TransactionPostingDate    *TS   `hl7:"false,Transaction Posting Date"`      // FT1-5
	TransactionType           *IS   `hl7:"true,Transaction Type"`               // FT1-6
	TransactionCode           *CE   `hl7:"true,Transaction Code"`               // FT1-7
	TransactionDescription    *ST   `hl7:"false,Transaction Description"`       // FT1-8
	TransactionDescriptionAlt *ST   `hl7:"false,Transaction Description - Alt"` // FT1-9
	TransactionQuantity       *NM   `hl7:"false,Transaction Quantity"`          // FT1-10
	TransactionAmountExtended *CP   `hl7:"false,Transaction Amount - Extended"` // FT1-11
	TransactionAmountUnit     *CP   `hl7:"false,Transaction Amount - Unit"`     // FT1-12
	DepartmentCode            *CE   `hl7:"false,Department Code"`               // FT1-13
	InsurancePlanID           *CE   `hl7:"false,Insurance Plan ID"`             // FT1-14
	InsuranceAmount           *CP   `hl7:"false,Insurance Amount"`              // FT1-15
	AssignedPatientLocation   *PL   `hl7:"false,Assigned Patient Location"`     // FT1-16
	FeeSchedule               *IS   `hl7:"false,Fee Schedule"`                  // FT1-17
	PatientType               *IS   `hl7:"false,Patient Type"`                  // FT1-18
	DiagnosisCodeFT1          []CE  `hl7:"false,Diagnosis Code - FT1"`          // FT1-19
	PerformedByCode           []XCN `hl7:"false,Performed By Code"`             // FT1-20
	OrderedByCode             []XCN `hl7:"false,Ordered By Code"`               // FT1-21
	UnitCost                  *CP   `hl7:"false,Unit Cost"`                     // FT1-22
	FillerOrderNumber         *EI   `hl7:"false,Filler Order Number"`           // FT1-23
	EnteredByCode             []XCN `hl7:"false,Entered By Code"`               // FT1-24
	ProcedureCode             *CE   `hl7:"false,Procedure Code"`                // FT1-25
	ProcedureCodeModifier     []CE  `hl7:"false,Procedure Code Modifier"`       // FT1-26
}

func (s *FT1) SegmentName() string {
	return "FT1"
}

// FTS represents the corresponding HL7 segment.
// Definition from HL7 2.4
type FTS struct {
	FileBatchCount     *NM `hl7:"false,File Batch Count"`     // FTS-1
	FileTrailerComment *ST `hl7:"false,File Trailer Comment"` // FTS-2
}

func (s *FTS) SegmentName() string {
	return "FTS"
}

// GOL represents the corresponding HL7 segment.
// Definition from HL7 2.4
type GOL struct {
	ActionCode                  *ID   `hl7:"true,Action Code"`                       // GOL-1
	ActionDateTime              *TS   `hl7:"true,Action Date/Time"`                  // GOL-2
	GoalID                      *CE   `hl7:"true,Goal ID"`                           // GOL-3
	GoalInstanceID              *EI   `hl7:"true,Goal Instance ID"`                  // GOL-4
	EpisodeOfCareID             *EI   `hl7:"false,Episode Of Care ID"`               // GOL-5
	GoalListPriority            *NM   `hl7:"false,Goal List Priority"`               // GOL-6
	GoalEstablishedDateTime     *TS   `hl7:"false,Goal Established Date/Time"`       // GOL-7
	ExpectedGoalAchieveDateTime *TS   `hl7:"false,Expected Goal Achieve Date/Time"`  // GOL-8
	GoalClassification          *CE   `hl7:"false,Goal Classification"`              // GOL-9
	GoalManagementDiscipline    *CE   `hl7:"false,Goal Management Discipline"`       // GOL-10
	CurrentGoalReviewStatus     *CE   `hl7:"false,Current Goal Review Status"`       // GOL-11
	CurrentGoalReviewDateTime   *TS   `hl7:"false,Current Goal Review Date/Time"`    // GOL-12
	NextGoalReviewDateTime      *TS   `hl7:"false,Next Goal Review Date/Time"`       // GOL-13
	PreviousGoalReviewDateTime  *TS   `hl7:"false,Previous Goal Review Date/Time"`   // GOL-14
	GoalReviewInterval          *TQ   `hl7:"false,Goal Review Interval"`             // GOL-15
	GoalEvaluation              *CE   `hl7:"false,Goal Evaluation"`                  // GOL-16
	GoalEvaluationComment       []ST  `hl7:"false,Goal Evaluation Comment"`          // GOL-17
	GoalLifeCycleStatus         *CE   `hl7:"false,Goal Life Cycle Status"`           // GOL-18
	GoalLifeCycleStatusDateTime *TS   `hl7:"false,Goal Life Cycle Status Date/Time"` // GOL-19
	GoalTargetType              []CE  `hl7:"false,Goal Target Type"`                 // GOL-20
	GoalTargetName              []XPN `hl7:"false,Goal Target Name"`                 // GOL-21
}

func (s *GOL) SegmentName() string {
	return "GOL"
}

// GP1 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type GP1 struct {
	TypeOfBillCode              *IS  `hl7:"true,Type Of Bill Code"`               // GP1-1
	RevenueCode                 []IS `hl7:"false,Revenue Code"`                   // GP1-2
	OverallClaimDispositionCode *IS  `hl7:"false,Overall Claim Disposition Code"` // GP1-3
	OCEEditsPerVisitCode        []IS `hl7:"false,OCE Edits Per Visit Code"`       // GP1-4
	OutlierCost                 *CP  `hl7:"false,Outlier Cost"`                   // GP1-5
}

func (s *GP1) SegmentName() string {
	return "GP1"
}

// GP2 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type GP2 struct {
	RevenueCode                         *IS  `hl7:"false,Revenue Code"`                           // GP2-1
	NumberOfServiceUnits                *NM  `hl7:"false,Number Of Service Units"`                // GP2-2
	Charge                              *CP  `hl7:"false,Charge"`                                 // GP2-3
	ReimbursementActionCode             *IS  `hl7:"false,Reimbursement Action Code"`              // GP2-4
	DenialOrRejectionCode               *IS  `hl7:"false,Denial Or Rejection Code"`               // GP2-5
	OCEEditCode                         []IS `hl7:"false,OCE Edit Code"`                          // GP2-6
	AmbulatoryPaymentClassificationCode *CE  `hl7:"false,Ambulatory Payment Classification Code"` // GP2-7
	ModifierEditCode                    []IS `hl7:"false,Modifier Edit Code"`                     // GP2-8
	PaymentAdjustmentCode               *IS  `hl7:"false,Payment Adjustment Code"`                // GP2-9
	PackagingStatusCode                 *IS  `hl7:"false,Packaging Status Code"`                  // GP2-10
	ExpectedHCFAPaymentAmount           *CP  `hl7:"false,Expected HCFA Payment Amount"`           // GP2-11
	ReimbursementTypeCode               *IS  `hl7:"false,Reimbursement Type Code"`                // GP2-12
	CoPayAmount                         *CP  `hl7:"false,Co-Pay Amount"`                          // GP2-13
	PayRatePerUnit                      *NM  `hl7:"false,Pay Rate Per Unit"`                      // GP2-14
}

func (s *GP2) SegmentName() string {
	return "GP2"
}

// GT1 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type GT1 struct {
	SetIDGT1                           *SI   `hl7:"true,Set ID - GT1"`                            // GT1-1
	GuarantorNumber                    []CX  `hl7:"false,Guarantor Number"`                       // GT1-2
	GuarantorName                      []XPN `hl7:"true,Guarantor Name"`                          // GT1-3
	GuarantorSpouseName                []XPN `hl7:"false,Guarantor Spouse Name"`                  // GT1-4
	GuarantorAddress                   []XAD `hl7:"false,Guarantor Address"`                      // GT1-5
	GuarantorPhNumHome                 []XTN `hl7:"false,Guarantor Ph Num - Home"`                // GT1-6
	GuarantorPhNumBusiness             []XTN `hl7:"false,Guarantor Ph Num - Business"`            // GT1-7
	GuarantorDateTimeOfBirth           *TS   `hl7:"false,Guarantor Date/Time Of Birth"`           // GT1-8
	GuarantorAdministrativeSex         *IS   `hl7:"false,Guarantor Administrative Sex"`           // GT1-9
	GuarantorType                      *IS   `hl7:"false,Guarantor Type"`                         // GT1-10
	GuarantorRelationship              *CE   `hl7:"false,Guarantor Relationship"`                 // GT1-11
	GuarantorSSN                       *ST   `hl7:"false,Guarantor SSN"`                          // GT1-12
	GuarantorDateBegin                 *DT   `hl7:"false,Guarantor Date - Begin"`                 // GT1-13
	GuarantorDateEnd                   *DT   `hl7:"false,Guarantor Date - End"`                   // GT1-14
	GuarantorPriority                  *NM   `hl7:"false,Guarantor Priority"`                     // GT1-15
	GuarantorEmployerName              []XPN `hl7:"false,Guarantor Employer Name"`                // GT1-16
	GuarantorEmployerAddress           []XAD `hl7:"false,Guarantor Employer Address"`             // GT1-17
	GuarantorEmployerPhoneNumber       []XTN `hl7:"false,Guarantor Employer Phone Number"`        // GT1-18
	GuarantorEmployeeIDNumber          []CX  `hl7:"false,Guarantor Employee ID Number"`           // GT1-19
	GuarantorEmploymentStatus          *IS   `hl7:"false,Guarantor Employment Status"`            // GT1-20
	GuarantorOrganizationName          []XON `hl7:"false,Guarantor Organization Name"`            // GT1-21
	GuarantorBillingHoldFlag           *ID   `hl7:"false,Guarantor Billing Hold Flag"`            // GT1-22
	GuarantorCreditRatingCode          *CE   `hl7:"false,Guarantor Credit Rating Code"`           // GT1-23
	GuarantorDeathDateAndTime          *TS   `hl7:"false,Guarantor Death Date And Time"`          // GT1-24
	GuarantorDeathFlag                 *ID   `hl7:"false,Guarantor Death Flag"`                   // GT1-25
	GuarantorChargeAdjustmentCode      *CE   `hl7:"false,Guarantor Charge Adjustment Code"`       // GT1-26
	GuarantorHouseholdAnnualIncome     *CP   `hl7:"false,Guarantor Household Annual Income"`      // GT1-27
	GuarantorHouseholdSize             *NM   `hl7:"false,Guarantor Household Size"`               // GT1-28
	GuarantorEmployerIDNumber          []CX  `hl7:"false,Guarantor Employer ID Number"`           // GT1-29
	GuarantorMaritalStatusCode         *CE   `hl7:"false,Guarantor Marital Status Code"`          // GT1-30
	GuarantorHireEffectiveDate         *DT   `hl7:"false,Guarantor Hire Effective Date"`          // GT1-31
	EmploymentStopDate                 *DT   `hl7:"false,Employment Stop Date"`                   // GT1-32
	LivingDependency                   *IS   `hl7:"false,Living Dependency"`                      // GT1-33
	AmbulatoryStatus                   []IS  `hl7:"false,Ambulatory Status"`                      // GT1-34
	Citizenship                        []CE  `hl7:"false,Citizenship"`                            // GT1-35
	PrimaryLanguage                    *CE   `hl7:"false,Primary Language"`                       // GT1-36
	LivingArrangement                  *IS   `hl7:"false,Living Arrangement"`                     // GT1-37
	PublicityCode                      *CE   `hl7:"false,Publicity Code"`                         // GT1-38
	ProtectionIndicator                *ID   `hl7:"false,Protection Indicator"`                   // GT1-39
	StudentIndicator                   *IS   `hl7:"false,Student Indicator"`                      // GT1-40
	Religion                           *CE   `hl7:"false,Religion"`                               // GT1-41
	MotherSMaidenName                  []XPN `hl7:"false,Mother'S Maiden Name"`                   // GT1-42
	Nationality                        *CE   `hl7:"false,Nationality"`                            // GT1-43
	EthnicGroup                        []CE  `hl7:"false,Ethnic Group"`                           // GT1-44
	ContactPersonSName                 []XPN `hl7:"false,Contact Person'S Name"`                  // GT1-45
	ContactPersonSTelephoneNumber      []XTN `hl7:"false,Contact Person'S Telephone Number"`      // GT1-46
	ContactReason                      *CE   `hl7:"false,Contact Reason"`                         // GT1-47
	ContactRelationship                *IS   `hl7:"false,Contact Relationship"`                   // GT1-48
	JobTitle                           *ST   `hl7:"false,Job Title"`                              // GT1-49
	JobCodeClass                       *JCC  `hl7:"false,Job Code/Class"`                         // GT1-50
	GuarantorEmployerSOrganizationName []XON `hl7:"false,Guarantor Employer'S Organization Name"` // GT1-51
	Handicap                           *IS   `hl7:"false,Handicap"`                               // GT1-52
	JobStatus                          *IS   `hl7:"false,Job Status"`                             // GT1-53
	GuarantorFinancialClass            *FC   `hl7:"false,Guarantor Financial Class"`              // GT1-54
	GuarantorRace                      []CE  `hl7:"false,Guarantor Race"`                         // GT1-55
}

func (s *GT1) SegmentName() string {
	return "GT1"
}

// IAM represents the corresponding HL7 segment.
// Definition from HL7 2.4
type IAM struct {
	SetIDIAM                             *SI  `hl7:"true,Set ID - IAM"`                              // IAM-1
	AllergenTypeCode                     *CE  `hl7:"false,Allergen Type Code"`                       // IAM-2
	AllergenCodeMnemonicDescription      *CE  `hl7:"true,Allergen Code/Mnemonic/Description"`        // IAM-3
	AllergySeverityCode                  *CE  `hl7:"false,Allergy Severity Code"`                    // IAM-4
	AllergyReactionCode                  []ST `hl7:"false,Allergy Reaction Code"`                    // IAM-5
	AllergyActionCode                    *CNE `hl7:"true,Allergy Action Code"`                       // IAM-6
	AllergyUniqueIdentifier              *EI  `hl7:"true,Allergy Unique Identifier"`                 // IAM-7
	ActionReason                         *ST  `hl7:"false,Action Reason"`                            // IAM-8
	SensitivityToCausativeAgentCode      *CE  `hl7:"false,Sensitivity To Causative Agent Code"`      // IAM-9
	AllergenGroupCodeMnemonicDescription *CE  `hl7:"false,Allergen Group Code/Mnemonic/Description"` // IAM-10
	OnsetDate                            *DT  `hl7:"false,Onset Date"`                               // IAM-11
	OnsetDateText                        *ST  `hl7:"false,Onset Date Text"`                          // IAM-12
	ReportedDateTime                     *TS  `hl7:"false,Reported Date/Time"`                       // IAM-13
	ReportedBy                           *XPN `hl7:"false,Reported By"`                              // IAM-14
	RelationshipToPatientCode            *CE  `hl7:"false,Relationship To Patient Code"`             // IAM-15
	AlertDeviceCode                      *CE  `hl7:"false,Alert Device Code"`                        // IAM-16
	AllergyClinicalStatusCode            *CE  `hl7:"false,Allergy Clinical Status Code"`             // IAM-17
	StatusedByPerson                     *XCN `hl7:"false,Statused By Person"`                       // IAM-18
	StatusedByOrganization               *XON `hl7:"false,Statused By Organization"`                 // IAM-19
	StatusedAtDateTime                   *TS  `hl7:"false,Statused At Date/Time"`                    // IAM-20
}

func (s *IAM) SegmentName() string {
	return "IAM"
}

// IN1 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type IN1 struct {
	SetIDIN1                      *SI   `hl7:"true,Set ID - IN1"`                       // IN1-1
	InsurancePlanID               *CE   `hl7:"true,Insurance Plan ID"`                  // IN1-2
	InsuranceCompanyID            []CX  `hl7:"true,Insurance Company ID"`               // IN1-3
	InsuranceCompanyName          []XON `hl7:"false,Insurance Company Name"`            // IN1-4
	InsuranceCompanyAddress       []XAD `hl7:"false,Insurance Company Address"`         // IN1-5
	InsuranceCoContactPerson      []XPN `hl7:"false,Insurance Co Contact Person"`       // IN1-6
	InsuranceCoPhoneNumber        []XTN `hl7:"false,Insurance Co Phone Number"`         // IN1-7
	GroupNumber                   *ST   `hl7:"false,Group Number"`                      // IN1-8
	GroupName                     []XON `hl7:"false,Group Name"`                        // IN1-9
	InsuredSGroupEmpID            []CX  `hl7:"false,Insured'S Group Emp ID"`            // IN1-10
	InsuredSGroupEmpName          []XON `hl7:"false,Insured'S Group Emp Name"`          // IN1-11
	PlanEffectiveDate             *DT   `hl7:"false,Plan Effective Date"`               // IN1-12
	PlanExpirationDate            *DT   `hl7:"false,Plan Expiration Date"`              // IN1-13
	AuthorizationInformation      *AUI  `hl7:"false,Authorization Information"`         // IN1-14
	PlanType                      *IS   `hl7:"false,Plan Type"`                         // IN1-15
	NameOfInsured                 []XPN `hl7:"false,Name Of Insured"`                   // IN1-16
	InsuredSRelationshipToPatient *CE   `hl7:"false,Insured'S Relationship To Patient"` // IN1-17
	InsuredSDateOfBirth           *TS   `hl7:"false,Insured'S Date Of Birth"`           // IN1-18
	InsuredSAddress               []XAD `hl7:"false,Insured'S Address"`                 // IN1-19
	AssignmentOfBenefits          *IS   `hl7:"false,Assignment Of Benefits"`            // IN1-20
	CoordinationOfBenefits        *IS   `hl7:"false,Coordination Of Benefits"`          // IN1-21
	CoordOfBenPriority            *ST   `hl7:"false,Coord Of Ben. Priority"`            // IN1-22
	NoticeOfAdmissionFlag         *ID   `hl7:"false,Notice Of Admission Flag"`          // IN1-23
	NoticeOfAdmissionDate         *DT   `hl7:"false,Notice Of Admission Date"`          // IN1-24
	ReportOfEligibilityFlag       *ID   `hl7:"false,Report Of Eligibility Flag"`        // IN1-25
	ReportOfEligibilityDate       *DT   `hl7:"false,Report Of Eligibility Date"`        // IN1-26
	ReleaseInformationCode        *IS   `hl7:"false,Release Information Code"`          // IN1-27
	PreAdmitCert                  *ST   `hl7:"false,Pre-Admit Cert"`                    // IN1-28
	VerificationDateTime          *TS   `hl7:"false,Verification Date/Time"`            // IN1-29
	VerificationBy                []XCN `hl7:"false,Verification By"`                   // IN1-30
	TypeOfAgreementCode           *IS   `hl7:"false,Type Of Agreement Code"`            // IN1-31
	BillingStatus                 *IS   `hl7:"false,Billing Status"`                    // IN1-32
	LifetimeReserveDays           *NM   `hl7:"false,Lifetime Reserve Days"`             // IN1-33
	DelayBeforeLRDay              *NM   `hl7:"false,Delay Before L.R. Day"`             // IN1-34
	CompanyPlanCode               *IS   `hl7:"false,Company Plan Code"`                 // IN1-35
	PolicyNumber                  *ST   `hl7:"false,Policy Number"`                     // IN1-36
	PolicyDeductible              *CP   `hl7:"false,Policy Deductible"`                 // IN1-37
	PolicyLimitAmount             *CP   `hl7:"false,Policy Limit - Amount"`             // IN1-38
	PolicyLimitDays               *NM   `hl7:"false,Policy Limit - Days"`               // IN1-39
	RoomRateSemiPrivate           *CP   `hl7:"false,Room Rate - Semi-Private"`          // IN1-40
	RoomRatePrivate               *CP   `hl7:"false,Room Rate - Private"`               // IN1-41
	InsuredSEmploymentStatus      *CE   `hl7:"false,Insured'S Employment Status"`       // IN1-42
	InsuredSAdministrativeSex     *IS   `hl7:"false,Insured'S Administrative Sex"`      // IN1-43
	InsuredSEmployerSAddress      []XAD `hl7:"false,Insured'S Employer'S Address"`      // IN1-44
	VerificationStatus            *ST   `hl7:"false,Verification Status"`               // IN1-45
	PriorInsurancePlanID          *IS   `hl7:"false,Prior Insurance Plan ID"`           // IN1-46
	CoverageType                  *IS   `hl7:"false,Coverage Type"`                     // IN1-47
	Handicap                      *IS   `hl7:"false,Handicap"`                          // IN1-48
	InsuredSIDNumber              []CX  `hl7:"false,Insured'S ID Number"`               // IN1-49
}

func (s *IN1) SegmentName() string {
	return "IN1"
}

// IN2 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type IN2 struct {
	InsuredSEmployeeID                   []CX  `hl7:"false,Insured'S Employee ID"`                     // IN2-1
	InsuredSSocialSecurityNumber         *ST   `hl7:"false,Insured'S Social Security Number"`          // IN2-2
	InsuredSEmployerSNameAndID           []XCN `hl7:"false,Insured'S Employer'S Name And ID"`          // IN2-3
	EmployerInformationData              *IS   `hl7:"false,Employer Information Data"`                 // IN2-4
	MailClaimParty                       []IS  `hl7:"false,Mail Claim Party"`                          // IN2-5
	MedicareHealthInsCardNumber          *ST   `hl7:"false,Medicare Health Ins Card Number"`           // IN2-6
	MedicaidCaseName                     []XPN `hl7:"false,Medicaid Case Name"`                        // IN2-7
	MedicaidCaseNumber                   *ST   `hl7:"false,Medicaid Case Number"`                      // IN2-8
	MilitarySponsorName                  []XPN `hl7:"false,Military Sponsor Name"`                     // IN2-9
	MilitaryIDNumber                     *ST   `hl7:"false,Military ID Number"`                        // IN2-10
	DependentOfMilitaryRecipient         *CE   `hl7:"false,Dependent Of Military Recipient"`           // IN2-11
	MilitaryOrganization                 *ST   `hl7:"false,Military Organization"`                     // IN2-12
	MilitaryStation                      *ST   `hl7:"false,Military Station"`                          // IN2-13
	MilitaryService                      *IS   `hl7:"false,Military Service"`                          // IN2-14
	MilitaryRankGrade                    *IS   `hl7:"false,Military Rank/Grade"`                       // IN2-15
	MilitaryStatus                       *IS   `hl7:"false,Military Status"`                           // IN2-16
	MilitaryRetireDate                   *DT   `hl7:"false,Military Retire Date"`                      // IN2-17
	MilitaryNonAvailCertOnFile           *ID   `hl7:"false,Military Non-Avail Cert On File"`           // IN2-18
	BabyCoverage                         *ID   `hl7:"false,Baby Coverage"`                             // IN2-19
	CombineBabyBill                      *ID   `hl7:"false,Combine Baby Bill"`                         // IN2-20
	BloodDeductible                      *ST   `hl7:"false,Blood Deductible"`                          // IN2-21
	SpecialCoverageApprovalName          []XPN `hl7:"false,Special Coverage Approval Name"`            // IN2-22
	SpecialCoverageApprovalTitle         *ST   `hl7:"false,Special Coverage Approval Title"`           // IN2-23
	NonCoveredInsuranceCode              []IS  `hl7:"false,Non-Covered Insurance Code"`                // IN2-24
	PayorID                              []CX  `hl7:"false,Payor ID"`                                  // IN2-25
	PayorSubscriberID                    []CX  `hl7:"false,Payor Subscriber ID"`                       // IN2-26
	EligibilitySource                    *IS   `hl7:"false,Eligibility Source"`                        // IN2-27
	RoomCoverageTypeAmount               []RMC `hl7:"false,Room Coverage Type/Amount"`                 // IN2-28
	PolicyTypeAmount                     []PTA `hl7:"false,Policy Type/Amount"`                        // IN2-29
	DailyDeductible                      *DDI  `hl7:"false,Daily Deductible"`                          // IN2-30
	LivingDependency                     *IS   `hl7:"false,Living Dependency"`                         // IN2-31
	AmbulatoryStatus                     []IS  `hl7:"false,Ambulatory Status"`                         // IN2-32
	Citizenship                          []CE  `hl7:"false,Citizenship"`                               // IN2-33
	PrimaryLanguage                      *CE   `hl7:"false,Primary Language"`                          // IN2-34
	LivingArrangement                    *IS   `hl7:"false,Living Arrangement"`                        // IN2-35
	PublicityCode                        *CE   `hl7:"false,Publicity Code"`                            // IN2-36
	ProtectionIndicator                  *ID   `hl7:"false,Protection Indicator"`                      // IN2-37
	StudentIndicator                     *IS   `hl7:"false,Student Indicator"`                         // IN2-38
	Religion                             *CE   `hl7:"false,Religion"`                                  // IN2-39
	MotherSMaidenName                    []XPN `hl7:"false,Mother'S Maiden Name"`                      // IN2-40
	Nationality                          *CE   `hl7:"false,Nationality"`                               // IN2-41
	EthnicGroup                          []CE  `hl7:"false,Ethnic Group"`                              // IN2-42
	MaritalStatus                        []CE  `hl7:"false,Marital Status"`                            // IN2-43
	InsuredSEmploymentStartDate          *DT   `hl7:"false,Insured'S Employment Start Date"`           // IN2-44
	EmploymentStopDate                   *DT   `hl7:"false,Employment Stop Date"`                      // IN2-45
	JobTitle                             *ST   `hl7:"false,Job Title"`                                 // IN2-46
	JobCodeClass                         *JCC  `hl7:"false,Job Code/Class"`                            // IN2-47
	JobStatus                            *IS   `hl7:"false,Job Status"`                                // IN2-48
	EmployerContactPersonName            []XPN `hl7:"false,Employer Contact Person Name"`              // IN2-49
	EmployerContactPersonPhoneNumber     []XTN `hl7:"false,Employer Contact Person Phone Number"`      // IN2-50
	EmployerContactReason                *IS   `hl7:"false,Employer Contact Reason"`                   // IN2-51
	InsuredSContactPersonSName           []XPN `hl7:"false,Insured'S Contact Person'S Name"`           // IN2-52
	InsuredSContactPersonPhoneNumber     []XTN `hl7:"false,Insured'S Contact Person Phone Number"`     // IN2-53
	InsuredSContactPersonReason          []IS  `hl7:"false,Insured'S Contact Person Reason"`           // IN2-54
	RelationshipToThePatientStartDate    *DT   `hl7:"false,Relationship To The Patient Start Date"`    // IN2-55
	RelationshipToThePatientStopDate     []DT  `hl7:"false,Relationship To The Patient Stop Date"`     // IN2-56
	InsuranceCoContactReason             *IS   `hl7:"false,Insurance Co. Contact Reason"`              // IN2-57
	InsuranceCoContactPhoneNumber        *XTN  `hl7:"false,Insurance Co Contact Phone Number"`         // IN2-58
	PolicyScope                          *IS   `hl7:"false,Policy Scope"`                              // IN2-59
	PolicySource                         *IS   `hl7:"false,Policy Source"`                             // IN2-60
	PatientMemberNumber                  *CX   `hl7:"false,Patient Member Number"`                     // IN2-61
	GuarantorSRelationshipToInsured      *CE   `hl7:"false,Guarantor'S Relationship To Insured"`       // IN2-62
	InsuredSPhoneNumberHome              []XTN `hl7:"false,Insured'S Phone Number - Home"`             // IN2-63
	InsuredSEmployerPhoneNumber          []XTN `hl7:"false,Insured'S Employer Phone Number"`           // IN2-64
	MilitaryHandicappedProgram           *CE   `hl7:"false,Military Handicapped Program"`              // IN2-65
	SuspendFlag                          *ID   `hl7:"false,Suspend Flag"`                              // IN2-66
	CopayLimitFlag                       *ID   `hl7:"false,Copay Limit Flag"`                          // IN2-67
	StoplossLimitFlag                    *ID   `hl7:"false,Stoploss Limit Flag"`                       // IN2-68
	InsuredOrganizationNameAndID         []XON `hl7:"false,Insured Organization Name And ID"`          // IN2-69
	InsuredEmployerOrganizationNameAndID []XON `hl7:"false,Insured Employer Organization Name And ID"` // IN2-70
	Race                                 []CE  `hl7:"false,Race"`                                      // IN2-71
	HCFAPatientSRelationshipToInsured    *CE   `hl7:"false,HCFA Patient'S Relationship To Insured"`    // IN2-72
}

func (s *IN2) SegmentName() string {
	return "IN2"
}

// IN3 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type IN3 struct {
	SetIDIN3                           *SI   `hl7:"true,Set ID - IN3"`                           // IN3-1
	CertificationNumber                *CX   `hl7:"false,Certification Number"`                  // IN3-2
	CertifiedBy                        []XCN `hl7:"false,Certified By"`                          // IN3-3
	CertificationRequired              *ID   `hl7:"false,Certification Required"`                // IN3-4
	Penalty                            *MOP  `hl7:"false,Penalty"`                               // IN3-5
	CertificationDateTime              *TS   `hl7:"false,Certification Date/Time"`               // IN3-6
	CertificationModifyDateTime        *TS   `hl7:"false,Certification Modify Date/Time"`        // IN3-7
	Operator                           []XCN `hl7:"false,Operator"`                              // IN3-8
	CertificationBeginDate             *DT   `hl7:"false,Certification Begin Date"`              // IN3-9
	CertificationEndDate               *DT   `hl7:"false,Certification End Date"`                // IN3-10
	Days                               *DTN  `hl7:"false,Days"`                                  // IN3-11
	NonConcurCodeDescription           *CE   `hl7:"false,Non-Concur Code/Description"`           // IN3-12
	NonConcurEffectiveDateTime         *TS   `hl7:"false,Non-Concur Effective Date/Time"`        // IN3-13
	PhysicianReviewer                  []XCN `hl7:"false,Physician Reviewer"`                    // IN3-14
	CertificationContact               *ST   `hl7:"false,Certification Contact"`                 // IN3-15
	CertificationContactPhoneNumber    []XTN `hl7:"false,Certification Contact Phone Number"`    // IN3-16
	AppealReason                       *CE   `hl7:"false,Appeal Reason"`                         // IN3-17
	CertificationAgency                *CE   `hl7:"false,Certification Agency"`                  // IN3-18
	CertificationAgencyPhoneNumber     []XTN `hl7:"false,Certification Agency Phone Number"`     // IN3-19
	PreCertificationReqWindow          []PCF `hl7:"false,Pre-Certification Req/Window"`          // IN3-20
	CaseManager                        *ST   `hl7:"false,Case Manager"`                          // IN3-21
	SecondOpinionDate                  *DT   `hl7:"false,Second Opinion Date"`                   // IN3-22
	SecondOpinionStatus                *IS   `hl7:"false,Second Opinion Status"`                 // IN3-23
	SecondOpinionDocumentationReceived []IS  `hl7:"false,Second Opinion Documentation Received"` // IN3-24
	SecondOpinionPhysician             []XCN `hl7:"false,Second Opinion Physician"`              // IN3-25
}

func (s *IN3) SegmentName() string {
	return "IN3"
}

// INV represents the corresponding HL7 segment.
// Definition from HL7 2.4
type INV struct {
	SubstanceIdentifier          *CE  `hl7:"true,Substance Identifier"`            // INV-1
	SubstanceStatus              []CE `hl7:"true,Substance Status"`                // INV-2
	SubstanceType                *CE  `hl7:"false,Substance Type"`                 // INV-3
	InventoryContainerIdentifier *CE  `hl7:"false,Inventory Container Identifier"` // INV-4
	ContainerCarrierIdentifier   *CE  `hl7:"false,Container Carrier Identifier"`   // INV-5
	PositionOnCarrier            *CE  `hl7:"false,Position On Carrier"`            // INV-6
	InitialQuantity              *NM  `hl7:"false,Initial Quantity"`               // INV-7
	CurrentQuantity              *NM  `hl7:"false,Current Quantity"`               // INV-8
	AvailableQuantity            *NM  `hl7:"false,Available Quantity"`             // INV-9
	ConsumptionQuantity          *NM  `hl7:"false,Consumption Quantity"`           // INV-10
	QuantityUnits                *CE  `hl7:"false,Quantity Units"`                 // INV-11
	ExpirationDateTime           *TS  `hl7:"false,Expiration Date/Time"`           // INV-12
	FirstUsedDateTime            *TS  `hl7:"false,First Used Date/Time"`           // INV-13
	OnBoardStabilityDuration     *TQ  `hl7:"false,On Board Stability Duration"`    // INV-14
	TestFluidIdentifierS         []CE `hl7:"false,Test/Fluid Identifier(S)"`       // INV-15
	ManufacturerLotNumber        *ST  `hl7:"false,Manufacturer Lot Number"`        // INV-16
	ManufacturerIdentifier       *CE  `hl7:"false,Manufacturer Identifier"`        // INV-17
	SupplierIdentifier           *CE  `hl7:"false,Supplier Identifier"`            // INV-18
}

func (s *INV) SegmentName() string {
	return "INV"
}

// ISD represents the corresponding HL7 segment.
// Definition from HL7 2.4
type ISD struct {
	ReferenceInteractionNumberUniqueIdentifier *NM `hl7:"true,Reference Interaction Number (Unique Identifier)"` // ISD-1
	InteractionTypeIdentifier                  *CE `hl7:"false,Interaction Type Identifier"`                     // ISD-2
	InteractionActiveState                     *CE `hl7:"true,Interaction Active State"`                         // ISD-3
}

func (s *ISD) SegmentName() string {
	return "ISD"
}

// LAN represents the corresponding HL7 segment.
// Definition from HL7 2.4
type LAN struct {
	SetIDLAN                *SI  `hl7:"true,Set ID - LAN"`               // LAN-1
	LanguageCode            *CE  `hl7:"true,Language Code"`              // LAN-2
	LanguageAbilityCode     []CE `hl7:"false,Language Ability Code"`     // LAN-3
	LanguageProficiencyCode *CE  `hl7:"false,Language Proficiency Code"` // LAN-4
}

func (s *LAN) SegmentName() string {
	return "LAN"
}

// LCC represents the corresponding HL7 segment.
// Definition from HL7 2.4
type LCC struct {
	PrimaryKeyValueLCC *PL  `hl7:"true,Primary Key Value - LCC"` // LCC-1
	LocationDepartment *CE  `hl7:"true,Location Department"`     // LCC-2
	AccommodationType  []CE `hl7:"false,Accommodation Type"`     // LCC-3
	ChargeCode         []CE `hl7:"true,Charge Code"`             // LCC-4
}

func (s *LCC) SegmentName() string {
	return "LCC"
}

// LCH represents the corresponding HL7 segment.
// Definition from HL7 2.4
type LCH struct {
	PrimaryKeyValueLCH             *PL `hl7:"true,Primary Key Value - LCH"`           // LCH-1
	SegmentActionCode              *ID `hl7:"false,Segment Action Code"`              // LCH-2
	SegmentUniqueKey               *EI `hl7:"false,Segment Unique Key"`               // LCH-3
	LocationCharacteristicID       *CE `hl7:"true,Location Characteristic ID"`        // LCH-4
	LocationCharacteristicValueLCH *CE `hl7:"true,Location Characteristic Value-LCH"` // LCH-5
}

func (s *LCH) SegmentName() string {
	return "LCH"
}

// LDP represents the corresponding HL7 segment.
// Definition from HL7 2.4
type LDP struct {
	PrimaryKeyValueLDP  *PL  `hl7:"true,Primary Key Value - LDP"`  // LDP-1
	LocationDepartment  *CE  `hl7:"true,Location Department"`      // LDP-2
	LocationService     []IS `hl7:"false,Location Service"`        // LDP-3
	SpecialtyType       []CE `hl7:"false,Specialty Type"`          // LDP-4
	ValidPatientClasses []IS `hl7:"false,Valid Patient Classes"`   // LDP-5
	ActiveInactiveFlag  *ID  `hl7:"false,Active/Inactive Flag"`    // LDP-6
	ActivationDateLDP   *TS  `hl7:"false,Activation Date  LDP"`    // LDP-7
	InactivationDateLDP *TS  `hl7:"false,Inactivation Date - LDP"` // LDP-8
	InactivatedReason   *ST  `hl7:"false,Inactivated Reason"`      // LDP-9
	VisitingHours       []VH `hl7:"false,Visiting Hours"`          // LDP-10
	ContactPhone        *XTN `hl7:"false,Contact Phone"`           // LDP-11
	LocationCostCenter  *CE  `hl7:"false,Location Cost Center"`    // LDP-12
}

func (s *LDP) SegmentName() string {
	return "LDP"
}

// LOC represents the corresponding HL7 segment.
// Definition from HL7 2.4
type LOC struct {
	PrimaryKeyValueLOC  *PL   `hl7:"true,Primary Key Value - LOC"`  // LOC-1
	LocationDescription *ST   `hl7:"false,Location Description"`    // LOC-2
	LocationTypeLOC     []IS  `hl7:"true,Location Type - LOC"`      // LOC-3
	OrganizationNameLOC []XON `hl7:"false,Organization Name - LOC"` // LOC-4
	LocationAddress     []XAD `hl7:"false,Location Address"`        // LOC-5
	LocationPhone       []XTN `hl7:"false,Location Phone"`          // LOC-6
	LicenseNumber       []CE  `hl7:"false,License Number"`          // LOC-7
	LocationEquipment   []IS  `hl7:"false,Location Equipment"`      // LOC-8
	LocationServiceCode *IS   `hl7:"false,Location Service Code"`   // LOC-9
}

func (s *LOC) SegmentName() string {
	return "LOC"
}

// LRL represents the corresponding HL7 segment.
// Definition from HL7 2.4
type LRL struct {
	PrimaryKeyValueLRL                      *PL   `hl7:"true,Primary Key Value - LRL"`                     // LRL-1
	SegmentActionCode                       *ID   `hl7:"false,Segment Action Code"`                        // LRL-2
	SegmentUniqueKey                        *EI   `hl7:"false,Segment Unique Key"`                         // LRL-3
	LocationRelationshipID                  *CE   `hl7:"true,Location Relationship ID"`                    // LRL-4
	OrganizationalLocationRelationshipValue []XON `hl7:"false,Organizational Location Relationship Value"` // LRL-5
	PatientLocationRelationshipValue        *PL   `hl7:"false,Patient Location Relationship Value"`        // LRL-6
}

func (s *LRL) SegmentName() string {
	return "LRL"
}

// MFA represents the corresponding HL7 segment.
// Definition from HL7 2.4
type MFA struct {
	RecordLevelEventCode      *ID  `hl7:"true,Record-Level Event Code"`       // MFA-1
	MFNControlID              *ST  `hl7:"false,MFN Control ID"`               // MFA-2
	EventCompletionDateTime   *TS  `hl7:"false,Event Completion Date/Time"`   // MFA-3
	MFNRecordLevelErrorReturn *CE  `hl7:"true,MFN Record Level Error Return"` // MFA-4
	PrimaryKeyValueMFA        []CE `hl7:"true,Primary Key Value - MFA"`       // MFA-5
	PrimaryKeyValueTypeMFA    []ID `hl7:"true,Primary Key Value Type - MFA"`  // MFA-6
}

func (s *MFA) SegmentName() string {
	return "MFA"
}

// MFE represents the corresponding HL7 segment.
// Definition from HL7 2.4
type MFE struct {
	RecordLevelEventCode *ID   `hl7:"true,Record-Level Event Code"` // MFE-1
	MFNControlID         *ST   `hl7:"false,MFN Control ID"`         // MFE-2
	EffectiveDateTime    *TS   `hl7:"false,Effective Date/Time"`    // MFE-3
	PrimaryKeyValueMFE   []Any `hl7:"true,Primary Key Value - MFE"` // MFE-4
	PrimaryKeyValueType  []ID  `hl7:"true,Primary Key Value Type"`  // MFE-5
}

func (s *MFE) SegmentName() string {
	return "MFE"
}

// MFI represents the corresponding HL7 segment.
// Definition from HL7 2.4
type MFI struct {
	MasterFileIdentifier            *CE `hl7:"true,Master File Identifier"`              // MFI-1
	MasterFileApplicationIdentifier *HD `hl7:"false,Master File Application Identifier"` // MFI-2
	FileLevelEventCode              *ID `hl7:"true,File-Level Event Code"`               // MFI-3
	EnteredDateTime                 *TS `hl7:"false,Entered Date/Time"`                  // MFI-4
	EffectiveDateTime               *TS `hl7:"false,Effective Date/Time"`                // MFI-5
	ResponseLevelCode               *ID `hl7:"true,Response Level Code"`                 // MFI-6
}

func (s *MFI) SegmentName() string {
	return "MFI"
}

// MRG represents the corresponding HL7 segment.
// Definition from HL7 2.4
type MRG struct {
	PriorPatientIdentifierList []CX  `hl7:"true,Prior Patient Identifier List"` // MRG-1
	PriorAlternatePatientID    []CX  `hl7:"false,Prior Alternate Patient ID"`   // MRG-2
	PriorPatientAccountNumber  *CX   `hl7:"false,Prior Patient Account Number"` // MRG-3
	PriorPatientID             *CX   `hl7:"false,Prior Patient ID"`             // MRG-4
	PriorVisitNumber           *CX   `hl7:"false,Prior Visit Number"`           // MRG-5
	PriorAlternateVisitID      *CX   `hl7:"false,Prior Alternate Visit ID"`     // MRG-6
	PriorPatientName           []XPN `hl7:"false,Prior Patient Name"`           // MRG-7
}

func (s *MRG) SegmentName() string {
	return "MRG"
}

// MSA represents the corresponding HL7 segment.
// Definition from HL7 2.4
type MSA struct {
	AcknowledgementCode       *ID `hl7:"true,Acknowledgement Code"`         // MSA-1
	MessageControlID          *ST `hl7:"true,Message Control ID"`           // MSA-2
	TextMessage               *ST `hl7:"false,Text Message"`                // MSA-3
	ExpectedSequenceNumber    *NM `hl7:"false,Expected Sequence Number"`    // MSA-4
	DelayedAcknowledgmentType *ID `hl7:"false,Delayed Acknowledgment Type"` // MSA-5
	ErrorCondition            *CE `hl7:"false,Error Condition"`             // MSA-6
}

func (s *MSA) SegmentName() string {
	return "MSA"
}

// MSH represents the corresponding HL7 segment.
// Definition from HL7 2.4
type MSH struct {
	// Missing: MSH.1
	EncodingCharacters                  *Delimiters `hl7:"true,Encoding Characters"`                      // MSH-2
	SendingApplication                  *HD         `hl7:"false,Sending Application"`                     // MSH-3
	SendingFacility                     *HD         `hl7:"false,Sending Facility"`                        // MSH-4
	ReceivingApplication                *HD         `hl7:"false,Receiving Application"`                   // MSH-5
	ReceivingFacility                   *HD         `hl7:"false,Receiving Facility"`                      // MSH-6
	DateTimeOfMessage                   *TS         `hl7:"true,Date/Time Of Message"`                     // MSH-7
	Security                            *ST         `hl7:"false,Security"`                                // MSH-8
	MessageType                         *MSG        `hl7:"true,Message Type"`                             // MSH-9
	MessageControlID                    *ST         `hl7:"true,Message Control ID"`                       // MSH-10
	ProcessingID                        *PT         `hl7:"true,Processing ID"`                            // MSH-11
	VersionID                           *VID        `hl7:"true,Version ID"`                               // MSH-12
	SequenceNumber                      *NM         `hl7:"false,Sequence Number"`                         // MSH-13
	ContinuationPointer                 *ST         `hl7:"false,Continuation Pointer"`                    // MSH-14
	AcceptAcknowledgmentType            *ID         `hl7:"false,Accept Acknowledgment Type"`              // MSH-15
	ApplicationAcknowledgmentType       *ID         `hl7:"false,Application Acknowledgment Type"`         // MSH-16
	CountryCode                         *ID         `hl7:"false,Country Code"`                            // MSH-17
	CharacterSet                        []ID        `hl7:"false,Character Set"`                           // MSH-18
	PrincipalLanguageOfMessage          *CE         `hl7:"false,Principal Language Of Message"`           // MSH-19
	AlternateCharacterSetHandlingScheme *ID         `hl7:"false,Alternate Character Set Handling Scheme"` // MSH-20
	ConformanceStatementID              []ID        `hl7:"false,Conformance Statement ID"`                // MSH-21
}

func (s *MSH) SegmentName() string {
	return "MSH"
}

// NCK represents the corresponding HL7 segment.
// Definition from HL7 2.4
type NCK struct {
	SystemDateTime *TS `hl7:"true,System Date/Time"` // NCK-1
}

func (s *NCK) SegmentName() string {
	return "NCK"
}

// NDS represents the corresponding HL7 segment.
// Definition from HL7 2.4
type NDS struct {
	NotificationReferenceNumber *NM `hl7:"true,Notification Reference Number"` // NDS-1
	NotificationDateTime        *TS `hl7:"true,Notification Date/Time"`        // NDS-2
	NotificationAlertSeverity   *CE `hl7:"true,Notification Alert Severity"`   // NDS-3
	NotificationCode            *CE `hl7:"true,Notification Code"`             // NDS-4
}

func (s *NDS) SegmentName() string {
	return "NDS"
}

// NK1 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type NK1 struct {
	SetIDNK1                                 *SI   `hl7:"true,Set ID - NK1"`                                      // NK1-1
	Name                                     []XPN `hl7:"false,Name"`                                             // NK1-2
	Relationship                             *CE   `hl7:"false,Relationship"`                                     // NK1-3
	Address                                  []XAD `hl7:"false,Address"`                                          // NK1-4
	PhoneNumber                              []XTN `hl7:"false,Phone Number"`                                     // NK1-5
	BusinessPhoneNumber                      []XTN `hl7:"false,Business Phone Number"`                            // NK1-6
	ContactRole                              *CE   `hl7:"false,Contact Role"`                                     // NK1-7
	StartDate                                *DT   `hl7:"false,Start Date"`                                       // NK1-8
	EndDate                                  *DT   `hl7:"false,End Date"`                                         // NK1-9
	NextOfKinAssociatedPartiesJobTitle       *ST   `hl7:"false,Next Of Kin / Associated Parties Job Title"`       // NK1-10
	NextOfKinAssociatedPartiesJobCodeClass   *JCC  `hl7:"false,Next Of Kin / Associated Parties Job Code/Class"`  // NK1-11
	NextOfKinAssociatedPartiesEmployeeNumber *CX   `hl7:"false,Next Of Kin / Associated Parties Employee Number"` // NK1-12
	OrganizationNameNK1                      []XON `hl7:"false,Organization Name - NK1"`                          // NK1-13
	MaritalStatus                            *CE   `hl7:"false,Marital Status"`                                   // NK1-14
	AdministrativeSex                        *IS   `hl7:"false,Administrative Sex"`                               // NK1-15
	DateTimeOfBirth                          *TS   `hl7:"false,Date/Time Of Birth"`                               // NK1-16
	LivingDependency                         []IS  `hl7:"false,Living Dependency"`                                // NK1-17
	AmbulatoryStatus                         []IS  `hl7:"false,Ambulatory Status"`                                // NK1-18
	Citizenship                              []CE  `hl7:"false,Citizenship"`                                      // NK1-19
	PrimaryLanguage                          *CE   `hl7:"false,Primary Language"`                                 // NK1-20
	LivingArrangement                        *IS   `hl7:"false,Living Arrangement"`                               // NK1-21
	PublicityCode                            *CE   `hl7:"false,Publicity Code"`                                   // NK1-22
	ProtectionIndicator                      *ID   `hl7:"false,Protection Indicator"`                             // NK1-23
	StudentIndicator                         *IS   `hl7:"false,Student Indicator"`                                // NK1-24
	Religion                                 *CE   `hl7:"false,Religion"`                                         // NK1-25
	MotherSMaidenName                        []XPN `hl7:"false,Mother'S Maiden Name"`                             // NK1-26
	Nationality                              *CE   `hl7:"false,Nationality"`                                      // NK1-27
	EthnicGroup                              []CE  `hl7:"false,Ethnic Group"`                                     // NK1-28
	ContactReason                            []CE  `hl7:"false,Contact Reason"`                                   // NK1-29
	ContactPersonSName                       []XPN `hl7:"false,Contact Person'S Name"`                            // NK1-30
	ContactPersonSTelephoneNumber            []XTN `hl7:"false,Contact Person'S Telephone Number"`                // NK1-31
	ContactPersonSAddress                    []XAD `hl7:"false,Contact Person'S Address"`                         // NK1-32
	NextOfKinAssociatedPartySIdentifiers     []CX  `hl7:"false,Next Of Kin/Associated Party'S Identifiers"`       // NK1-33
	JobStatus                                *IS   `hl7:"false,Job Status"`                                       // NK1-34
	Race                                     []CE  `hl7:"false,Race"`                                             // NK1-35
	Handicap                                 *IS   `hl7:"false,Handicap"`                                         // NK1-36
	ContactPersonSocialSecurityNumber        *ST   `hl7:"false,Contact Person Social Security Number"`            // NK1-37
}

func (s *NK1) SegmentName() string {
	return "NK1"
}

// NPU represents the corresponding HL7 segment.
// Definition from HL7 2.4
type NPU struct {
	BedLocation *PL `hl7:"true,Bed Location"` // NPU-1
	BedStatus   *IS `hl7:"false,Bed Status"`  // NPU-2
}

func (s *NPU) SegmentName() string {
	return "NPU"
}

// NSC represents the corresponding HL7 segment.
// Definition from HL7 2.4
type NSC struct {
	ApplicationChangeType *IS `hl7:"true,Application Change Type"` // NSC-1
	CurrentCPU            *ST `hl7:"false,Current CPU"`            // NSC-2
	CurrentFileserver     *ST `hl7:"false,Current Fileserver"`     // NSC-3
	CurrentApplication    *HD `hl7:"false,Current Application"`    // NSC-4
	CurrentFacility       *HD `hl7:"false,Current Facility"`       // NSC-5
	NewCPU                *ST `hl7:"false,New CPU"`                // NSC-6
	NewFileserver         *ST `hl7:"false,New Fileserver"`         // NSC-7
	NewApplication        *HD `hl7:"false,New Application"`        // NSC-8
	NewFacility           *HD `hl7:"false,New Facility"`           // NSC-9
}

func (s *NSC) SegmentName() string {
	return "NSC"
}

// NST represents the corresponding HL7 segment.
// Definition from HL7 2.4
type NST struct {
	StatisticsAvailable           *ID `hl7:"true,Statistics Available"`              // NST-1
	SourceIdentifier              *ST `hl7:"false,Source Identifier"`                // NST-2
	SourceType                    *ID `hl7:"false,Source Type"`                      // NST-3
	StatisticsStart               *TS `hl7:"false,Statistics Start"`                 // NST-4
	StatisticsEnd                 *TS `hl7:"false,Statistics End"`                   // NST-5
	ReceiveCharacterCount         *NM `hl7:"false,Receive Character Count"`          // NST-6
	SendCharacterCount            *NM `hl7:"false,Send Character Count"`             // NST-7
	MessagesReceived              *NM `hl7:"false,Messages Received"`                // NST-8
	MessagesSent                  *NM `hl7:"false,Messages Sent"`                    // NST-9
	ChecksumErrorsReceived        *NM `hl7:"false,Checksum Errors Received"`         // NST-10
	LengthErrorsReceived          *NM `hl7:"false,Length Errors Received"`           // NST-11
	OtherErrorsReceived           *NM `hl7:"false,Other Errors Received"`            // NST-12
	ConnectTimeouts               *NM `hl7:"false,Connect Timeouts"`                 // NST-13
	ReceiveTimeouts               *NM `hl7:"false,Receive Timeouts"`                 // NST-14
	ApplicationControlLevelErrors *NM `hl7:"false,Application Control-Level Errors"` // NST-15
}

func (s *NST) SegmentName() string {
	return "NST"
}

// NTE represents the corresponding HL7 segment.
// Definition from HL7 2.4
type NTE struct {
	SetIDNTE        *SI  `hl7:"false,Set ID - NTE"`      // NTE-1
	SourceOfComment *ID  `hl7:"false,Source Of Comment"` // NTE-2
	Comment         []FT `hl7:"false,Comment"`           // NTE-3
	CommentType     *CE  `hl7:"false,Comment Type"`      // NTE-4
}

func (s *NTE) SegmentName() string {
	return "NTE"
}

// OBR represents the corresponding HL7 segment.
// Definition from HL7 2.4
type OBR struct {
	SetIDOBR                             *SI   `hl7:"false,Set ID - OBR"`                              // OBR-1
	PlacerOrderNumber                    *EI   `hl7:"false,Placer Order Number"`                       // OBR-2
	FillerOrderNumber                    *EI   `hl7:"false,Filler Order Number"`                       // OBR-3
	UniversalServiceIdentifier           *CE   `hl7:"true,Universal Service Identifier"`               // OBR-4
	Priority                             *ID   `hl7:"false,Priority"`                                  // OBR-5
	RequestedDateTime                    *TS   `hl7:"false,Requested Date/Time"`                       // OBR-6
	ObservationDateTime                  *TS   `hl7:"false,Observation Date/Time #"`                   // OBR-7
	ObservationEndDateTime               *TS   `hl7:"false,Observation End Date/Time #"`               // OBR-8
	CollectionVolume                     *CQ   `hl7:"false,Collection Volume *"`                       // OBR-9
	CollectorIdentifier                  []XCN `hl7:"false,Collector Identifier *"`                    // OBR-10
	SpecimenActionCode                   *ID   `hl7:"false,Specimen Action Code *"`                    // OBR-11
	DangerCode                           *CE   `hl7:"false,Danger Code"`                               // OBR-12
	RelevantClinicalInfo                 *ST   `hl7:"false,Relevant Clinical Info."`                   // OBR-13
	SpecimenReceivedDateTime             *TS   `hl7:"false,Specimen Received Date/Time *"`             // OBR-14
	SpecimenSource                       *SPS  `hl7:"false,Specimen Source"`                           // OBR-15
	OrderingProvider                     []XCN `hl7:"false,Ordering Provider"`                         // OBR-16
	OrderCallbackPhoneNumber             []XTN `hl7:"false,Order Callback Phone Number"`               // OBR-17
	PlacerField1                         *ST   `hl7:"false,Placer Field 1"`                            // OBR-18
	PlacerField2                         *ST   `hl7:"false,Placer Field 2"`                            // OBR-19
	FillerField1                         *ST   `hl7:"false,Filler Field 1 +"`                          // OBR-20
	FillerField2                         *ST   `hl7:"false,Filler Field 2 +"`                          // OBR-21
	ResultsRptStatusChngDateTime         *TS   `hl7:"false,Results Rpt/Status Chng - Date/Time +"`     // OBR-22
	ChargeToPractice                     *MOC  `hl7:"false,Charge To Practice +"`                      // OBR-23
	DiagnosticServSectID                 *ID   `hl7:"false,Diagnostic Serv Sect ID"`                   // OBR-24
	ResultStatus                         *ID   `hl7:"false,Result Status +"`                           // OBR-25
	ParentResult                         *PRL  `hl7:"false,Parent Result +"`                           // OBR-26
	QuantityTiming                       []TQ  `hl7:"false,Quantity/Timing"`                           // OBR-27
	ResultCopiesTo                       []XCN `hl7:"false,Result Copies To"`                          // OBR-28
	Parent                               *EIP  `hl7:"false,Parent"`                                    // OBR-29
	TransportationMode                   *ID   `hl7:"false,Transportation Mode"`                       // OBR-30
	ReasonForStudy                       []CE  `hl7:"false,Reason For Study"`                          // OBR-31
	PrincipalResultInterpreter           *NDL  `hl7:"false,Principal Result Interpreter +"`            // OBR-32
	AssistantResultInterpreter           []NDL `hl7:"false,Assistant Result Interpreter +"`            // OBR-33
	Technician                           []NDL `hl7:"false,Technician +"`                              // OBR-34
	Transcriptionist                     []NDL `hl7:"false,Transcriptionist +"`                        // OBR-35
	ScheduledDateTime                    *TS   `hl7:"false,Scheduled Date/Time +"`                     // OBR-36
	NumberOfSampleContainers             *NM   `hl7:"false,Number Of Sample Containers *"`             // OBR-37
	TransportLogisticsOfCollectedSample  []CE  `hl7:"false,Transport Logistics Of Collected Sample *"` // OBR-38
	CollectorSComment                    []CE  `hl7:"false,Collector'S Comment *"`                     // OBR-39
	TransportArrangementResponsibility   *CE   `hl7:"false,Transport Arrangement Responsibility"`      // OBR-40
	TransportArranged                    *ID   `hl7:"false,Transport Arranged"`                        // OBR-41
	EscortRequired                       *ID   `hl7:"false,Escort Required"`                           // OBR-42
	PlannedPatientTransportComment       []CE  `hl7:"false,Planned Patient Transport Comment"`         // OBR-43
	ProcedureCode                        *CE   `hl7:"false,Procedure Code"`                            // OBR-44
	ProcedureCodeModifier                []CE  `hl7:"false,Procedure Code Modifier"`                   // OBR-45
	PlacerSupplementalServiceInformation []CE  `hl7:"false,Placer Supplemental Service Information"`   // OBR-46
	FillerSupplementalServiceInformation []CE  `hl7:"false,Filler Supplemental Service Information"`   // OBR-47
}

func (s *OBR) SegmentName() string {
	return "OBR"
}

// OBX represents the corresponding HL7 segment.
// Definition from HL7 2.4
type OBX struct {
	SetIDOBX                       *SI   `hl7:"false,Set ID - OBX"`                       // OBX-1
	ValueType                      *ID   `hl7:"false,Value Type"`                         // OBX-2
	ObservationIdentifier          *CE   `hl7:"true,Observation Identifier"`              // OBX-3
	ObservationSubId               *ST   `hl7:"false,Observation Sub-Id"`                 // OBX-4
	ObservationValue               []Any `hl7:"false,Observation Value"`                  // OBX-5
	Units                          *CE   `hl7:"false,Units"`                              // OBX-6
	ReferencesRange                *ST   `hl7:"false,References Range"`                   // OBX-7
	AbnormalFlags                  *IS   `hl7:"false,Abnormal Flags"`                     // OBX-8
	Probability                    []NM  `hl7:"false,Probability"`                        // OBX-9
	NatureOfAbnormalTest           *ID   `hl7:"false,Nature Of Abnormal Test"`            // OBX-10
	ObservationResultStatus        *ID   `hl7:"true,Observation Result Status"`           // OBX-11
	DateLastObservationNormalValue *TS   `hl7:"false,Date Last Observation Normal Value"` // OBX-12
	UserDefinedAccessChecks        *ST   `hl7:"false,User Defined Access Checks"`         // OBX-13
	DateTimeOfTheObservation       *TS   `hl7:"false,Date/Time Of The Observation"`       // OBX-14
	ProducerSID                    *CE   `hl7:"false,Producer'S ID"`                      // OBX-15
	ResponsibleObserver            *XCN  `hl7:"false,Responsible Observer"`               // OBX-16
	ObservationMethod              []CE  `hl7:"false,Observation Method"`                 // OBX-17
	EquipmentInstanceIdentifier    []EI  `hl7:"false,Equipment Instance Identifier"`      // OBX-18
	DateTimeOfTheAnalysis          *TS   `hl7:"false,Date/Time Of The Analysis"`          // OBX-19
}

func (s *OBX) SegmentName() string {
	return "OBX"
}

// ODS represents the corresponding HL7 segment.
// Definition from HL7 2.4
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
// Definition from HL7 2.4
type ODT struct {
	TrayType        *CE  `hl7:"true,Tray Type"`         // ODT-1
	ServicePeriod   []CE `hl7:"false,Service Period"`   // ODT-2
	TextInstruction *ST  `hl7:"false,Text Instruction"` // ODT-3
}

func (s *ODT) SegmentName() string {
	return "ODT"
}

// OM1 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type OM1 struct {
	SequenceNumberTestObservationMasterFile                *NM   `hl7:"true,Sequence Number - Test/ Observation Master File"`                   // OM1-1
	ProducerSServiceTestObservationID                      *CE   `hl7:"true,Producer'S Service/Test/Observation ID"`                            // OM1-2
	PermittedDataTypes                                     []ID  `hl7:"false,Permitted Data Types"`                                             // OM1-3
	SpecimenRequired                                       *ID   `hl7:"true,Specimen Required"`                                                 // OM1-4
	ProducerID                                             *CE   `hl7:"true,Producer ID"`                                                       // OM1-5
	ObservationDescription                                 *TX   `hl7:"false,Observation Description"`                                          // OM1-6
	OtherServiceTestObservationIDsForTheObservation        *CE   `hl7:"false,Other Service/Test/Observation IDs For The Observation"`           // OM1-7
	OtherNames                                             []ST  `hl7:"true,Other Names"`                                                       // OM1-8
	PreferredReportNameForTheObservation                   *ST   `hl7:"false,Preferred Report Name For The Observation"`                        // OM1-9
	PreferredShortNameOrMnemonicForObservation             *ST   `hl7:"false,Preferred Short Name Or Mnemonic For Observation"`                 // OM1-10
	PreferredLongNameForTheObservation                     *ST   `hl7:"false,Preferred Long Name For The Observation"`                          // OM1-11
	Orderability                                           *ID   `hl7:"false,Orderability"`                                                     // OM1-12
	IdentityOfInstrumentUsedToPerformThisStudy             []CE  `hl7:"false,Identity Of Instrument Used To Perform This Study"`                // OM1-13
	CodedRepresentationOfMethod                            []CE  `hl7:"false,Coded Representation Of Method"`                                   // OM1-14
	PortableDeviceIndicator                                *ID   `hl7:"false,Portable Device Indicator"`                                        // OM1-15
	ObservationProducingDepartmentSection                  []CE  `hl7:"false,Observation Producing Department/Section"`                         // OM1-16
	TelephoneNumberOfSection                               *XTN  `hl7:"false,Telephone Number Of Section"`                                      // OM1-17
	NatureOfServiceTestObservation                         *IS   `hl7:"true,Nature Of Service/Test/Observation"`                                // OM1-18
	ReportSubheader                                        *CE   `hl7:"false,Report Subheader"`                                                 // OM1-19
	ReportDisplayOrder                                     *ST   `hl7:"false,Report Display Order"`                                             // OM1-20
	DateTimeStampForAnyChangeInDefinitionForTheObservation *TS   `hl7:"false,Date/Time Stamp For Any Change In Definition For The Observation"` // OM1-21
	EffectiveDateTimeOfChange                              *TS   `hl7:"false,Effective Date/Time Of Change"`                                    // OM1-22
	TypicalTurnAroundTime                                  *NM   `hl7:"false,Typical Turn-Around Time"`                                         // OM1-23
	ProcessingTime                                         *NM   `hl7:"false,Processing Time"`                                                  // OM1-24
	ProcessingPriority                                     []ID  `hl7:"false,Processing Priority"`                                              // OM1-25
	ReportingPriority                                      *ID   `hl7:"false,Reporting Priority"`                                               // OM1-26
	OutsideSiteSWhereObservationMayBePerformed             []CE  `hl7:"false,Outside Site(S) Where Observation May Be Performed"`               // OM1-27
	AddressOfOutsideSiteS                                  []XAD `hl7:"false,Address Of Outside Site(S)"`                                       // OM1-28
	PhoneNumberOfOutsideSite                               *XTN  `hl7:"false,Phone Number Of Outside Site"`                                     // OM1-29
	ConfidentialityCode                                    *IS   `hl7:"false,Confidentiality Code"`                                             // OM1-30
	ObservationsRequiredToInterpretTheObservation          *CE   `hl7:"false,Observations Required To Interpret The Observation"`               // OM1-31
	InterpretationOfObservations                           *TX   `hl7:"false,Interpretation Of Observations"`                                   // OM1-32
	ContraindicationsToObservations                        *CE   `hl7:"false,Contraindications To Observations"`                                // OM1-33
	ReflexTestsObservations                                []CE  `hl7:"false,Reflex Tests/Observations"`                                        // OM1-34
	RulesThatTriggerReflexTesting                          *TX   `hl7:"false,Rules That Trigger Reflex Testing"`                                // OM1-35
	FixedCannedMessage                                     *CE   `hl7:"false,Fixed Canned Message"`                                             // OM1-36
	PatientPreparation                                     *TX   `hl7:"false,Patient Preparation"`                                              // OM1-37
	ProcedureMedication                                    *CE   `hl7:"false,Procedure Medication"`                                             // OM1-38
	FactorsThatMayAffectAffectTheObservation               *TX   `hl7:"false,Factors That May Affect Affect The Observation"`                   // OM1-39
	ServiceTestObservationPerformanceSchedule              []ST  `hl7:"false,Service/Test/Observation Performance Schedule"`                    // OM1-40
	DescriptionOfTestMethods                               *TX   `hl7:"false,Description Of Test Methods"`                                      // OM1-41
	KindOfQuantityObserved                                 *CE   `hl7:"false,Kind Of Quantity Observed"`                                        // OM1-42
	PointVersusInterval                                    *CE   `hl7:"false,Point Versus Interval"`                                            // OM1-43
	ChallengeInformation                                   *TX   `hl7:"false,Challenge Information"`                                            // OM1-44
	RelationshipModifier                                   *CE   `hl7:"false,Relationship Modifier"`                                            // OM1-45
	TargetAnatomicSiteOfTest                               *CE   `hl7:"false,Target Anatomic Site Of Test"`                                     // OM1-46
	ModalityOfImagingMeasurement                           *CE   `hl7:"false,Modality Of Imaging Measurement"`                                  // OM1-47
}

func (s *OM1) SegmentName() string {
	return "OM1"
}

// OM2 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type OM2 struct {
	SequenceNumberTestObservationMasterFile              *NM   `hl7:"false,Sequence Number - Test/ Observation Master File"`                // OM2-1
	UnitsOfMeasure                                       *CE   `hl7:"false,Units Of Measure"`                                               // OM2-2
	RangeOfDecimalPrecision                              []NM  `hl7:"false,Range Of Decimal Precision"`                                     // OM2-3
	CorrespondingSIUnitsOfMeasure                        *CE   `hl7:"false,Corresponding SI Units Of Measure"`                              // OM2-4
	SIConversionFactor                                   *TX   `hl7:"false,SI Conversion Factor"`                                           // OM2-5
	ReferenceNormalRangeOrdinalAndContinuousObservations *RFR  `hl7:"false,Reference (Normal) Range - Ordinal And Continuous Observations"` // OM2-6
	CriticalRangeForOrdinalAndContinuousObservations     *NR   `hl7:"false,Critical Range For Ordinal And Continuous Observations"`         // OM2-7
	AbsoluteRangeForOrdinalAndContinuousObservations     *RFR  `hl7:"false,Absolute Range For Ordinal And Continuous Observations"`         // OM2-8
	DeltaCheckCriteria                                   []DLT `hl7:"false,Delta Check Criteria"`                                           // OM2-9
	MinimumMeaningfulIncrements                          *NM   `hl7:"false,Minimum Meaningful Increments"`                                  // OM2-10
}

func (s *OM2) SegmentName() string {
	return "OM2"
}

// OM3 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type OM3 struct {
	SequenceNumberTestObservationMasterFile     *NM  `hl7:"false,Sequence Number - Test/ Observation Master File"`  // OM3-1
	PreferredCodingSystem                       *CE  `hl7:"false,Preferred Coding System"`                          // OM3-2
	ValidCodedAnswers                           *CE  `hl7:"false,Valid Coded Answers"`                              // OM3-3
	NormalTextCodesForCategoricalObservations   []CE `hl7:"false,Normal Text/Codes For Categorical Observations"`   // OM3-4
	AbnormalTextCodesForCategoricalObservations *CE  `hl7:"false,Abnormal Text/Codes For Categorical Observations"` // OM3-5
	CriticalTextCodesForCategoricalObservations *CE  `hl7:"false,Critical Text/Codes For Categorical Observations"` // OM3-6
	ValueType                                   *ID  `hl7:"false,Value Type"`                                       // OM3-7
}

func (s *OM3) SegmentName() string {
	return "OM3"
}

// OM4 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type OM4 struct {
	SequenceNumberTestObservationMasterFile *NM  `hl7:"false,Sequence Number - Test/ Observation Master File"` // OM4-1
	DerivedSpecimen                         *ID  `hl7:"false,Derived Specimen"`                                // OM4-2
	ContainerDescription                    *TX  `hl7:"false,Container Description"`                           // OM4-3
	ContainerVolume                         *NM  `hl7:"false,Container Volume"`                                // OM4-4
	ContainerUnits                          *CE  `hl7:"false,Container Units"`                                 // OM4-5
	Specimen                                *CE  `hl7:"false,Specimen"`                                        // OM4-6
	Additive                                *CE  `hl7:"false,Additive"`                                        // OM4-7
	Preparation                             *TX  `hl7:"false,Preparation"`                                     // OM4-8
	SpecialHandlingRequirements             *TX  `hl7:"false,Special Handling Requirements"`                   // OM4-9
	NormalCollectionVolume                  *CQ  `hl7:"false,Normal Collection Volume"`                        // OM4-10
	MinimumCollectionVolume                 *CQ  `hl7:"false,Minimum Collection Volume"`                       // OM4-11
	SpecimenRequirements                    *TX  `hl7:"false,Specimen Requirements"`                           // OM4-12
	SpecimenPriorities                      []ID `hl7:"false,Specimen Priorities"`                             // OM4-13
	SpecimenRetentionTime                   *CQ  `hl7:"false,Specimen Retention Time"`                         // OM4-14
}

func (s *OM4) SegmentName() string {
	return "OM4"
}

// OM5 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type OM5 struct {
	SequenceNumberTestObservationMasterFile            *NM  `hl7:"false,Sequence Number - Test/ Observation Master File"`           // OM5-1
	TestObservationsIncludedWithinAnOrderedTestBattery []CE `hl7:"false,Test/Observations Included Within An Ordered Test Battery"` // OM5-2
	ObservationIDSuffixes                              *ST  `hl7:"false,Observation ID Suffixes"`                                   // OM5-3
}

func (s *OM5) SegmentName() string {
	return "OM5"
}

// OM6 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type OM6 struct {
	SequenceNumberTestObservationMasterFile *NM `hl7:"false,Sequence Number - Test/ Observation Master File"` // OM6-1
	DerivationRule                          *TX `hl7:"false,Derivation Rule"`                                 // OM6-2
}

func (s *OM6) SegmentName() string {
	return "OM6"
}

// OM7 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type OM7 struct {
	SequenceNumberTestObservationMasterFile *NM  `hl7:"true,Sequence Number - Test/ Observation Master File"` // OM7-1
	UniversalServiceIdentifier              *CE  `hl7:"true,Universal Service Identifier"`                    // OM7-2
	CategoryIdentifier                      []CE `hl7:"false,Category Identifier"`                            // OM7-3
	CategoryDescription                     *TX  `hl7:"false,Category Description"`                           // OM7-4
	CategorySynonym                         []ST `hl7:"false,Category Synonym"`                               // OM7-5
	EffectiveTestServiceStartDateTime       *TS  `hl7:"false,Effective Test/Service Start Date/Time"`         // OM7-6
	EffectiveTestServiceEndDateTime         *TS  `hl7:"false,Effective Test/Service End Date/Time"`           // OM7-7
	TestServiceDefaultDurationQuantity      *NM  `hl7:"false,Test/Service Default Duration Quantity"`         // OM7-8
	TestServiceDefaultDurationUnits         *CE  `hl7:"false,Test/Service Default Duration Units"`            // OM7-9
	TestServiceDefaultFrequency             *IS  `hl7:"false,Test/Service Default Frequency"`                 // OM7-10
	ConsentIndicator                        *ID  `hl7:"false,Consent Indicator"`                              // OM7-11
	ConsentIdentifier                       *CE  `hl7:"false,Consent Identifier"`                             // OM7-12
	ConsentEffectiveStartDateTime           *TS  `hl7:"false,Consent Effective Start Date/Time"`              // OM7-13
	ConsentEffectiveEndDateTime             *TS  `hl7:"false,Consent Effective End Date/Time"`                // OM7-14
	ConsentIntervalQuantity                 *NM  `hl7:"false,Consent Interval Quantity"`                      // OM7-15
	ConsentIntervalUnits                    *CE  `hl7:"false,Consent Interval Units"`                         // OM7-16
	ConsentWaitingPeriodQuantity            *NM  `hl7:"false,Consent Waiting Period Quantity"`                // OM7-17
	ConsentWaitingPeriodUnits               *CE  `hl7:"false,Consent Waiting Period Units"`                   // OM7-18
	EffectiveDateTimeOfChange               *TS  `hl7:"false,Effective Date/Time Of Change"`                  // OM7-19
	EnteredBy                               *XCN `hl7:"false,Entered By"`                                     // OM7-20
	OrderableAtLocation                     []PL `hl7:"false,Orderable-At Location"`                          // OM7-21
	FormularyStatus                         *IS  `hl7:"false,Formulary Status"`                               // OM7-22
	SpecialOrderIndicator                   *ID  `hl7:"false,Special Order Indicator"`                        // OM7-23
	PrimaryKeyValueCDM                      []CE `hl7:"false,Primary Key Value - CDM"`                        // OM7-24
}

func (s *OM7) SegmentName() string {
	return "OM7"
}

// ORC represents the corresponding HL7 segment.
// Definition from HL7 2.4
type ORC struct {
	OrderControl                  *ID   `hl7:"true,Order Control"`                     // ORC-1
	PlacerOrderNumber             *EI   `hl7:"false,Placer Order Number"`              // ORC-2
	FillerOrderNumber             *EI   `hl7:"false,Filler Order Number"`              // ORC-3
	PlacerGroupNumber             *EI   `hl7:"false,Placer Group Number"`              // ORC-4
	OrderStatus                   *ID   `hl7:"false,Order Status"`                     // ORC-5
	ResponseFlag                  *ID   `hl7:"false,Response Flag"`                    // ORC-6
	QuantityTiming                []TQ  `hl7:"false,Quantity/Timing"`                  // ORC-7
	Parent                        *EIP  `hl7:"false,Parent"`                           // ORC-8
	DateTimeOfTransaction         *TS   `hl7:"false,Date/Time Of Transaction"`         // ORC-9
	EnteredBy                     []XCN `hl7:"false,Entered By"`                       // ORC-10
	VerifiedBy                    []XCN `hl7:"false,Verified By"`                      // ORC-11
	OrderingProvider              []XCN `hl7:"false,Ordering Provider"`                // ORC-12
	EntererSLocation              *PL   `hl7:"false,Enterer'S Location"`               // ORC-13
	CallBackPhoneNumber           []XTN `hl7:"false,Call Back Phone Number"`           // ORC-14
	OrderEffectiveDateTime        *TS   `hl7:"false,Order Effective Date/Time"`        // ORC-15
	OrderControlCodeReason        *CE   `hl7:"false,Order Control Code Reason"`        // ORC-16
	EnteringOrganization          *CE   `hl7:"false,Entering Organization"`            // ORC-17
	EnteringDevice                *CE   `hl7:"false,Entering Device"`                  // ORC-18
	ActionBy                      []XCN `hl7:"false,Action By"`                        // ORC-19
	AdvancedBeneficiaryNoticeCode *CE   `hl7:"false,Advanced Beneficiary Notice Code"` // ORC-20
	OrderingFacilityName          []XON `hl7:"false,Ordering Facility Name"`           // ORC-21
	OrderingFacilityAddress       []XAD `hl7:"false,Ordering Facility Address"`        // ORC-22
	OrderingFacilityPhoneNumber   []XTN `hl7:"false,Ordering Facility Phone Number"`   // ORC-23
	OrderingProviderAddress       []XAD `hl7:"false,Ordering Provider Address"`        // ORC-24
	OrderStatusModifier           *CWE  `hl7:"false,Order Status Modifier"`            // ORC-25
}

func (s *ORC) SegmentName() string {
	return "ORC"
}

// ORG represents the corresponding HL7 segment.
// Definition from HL7 2.4
type ORG struct {
	SetIDORG                                   *SI `hl7:"true,Set ID - ORG"`                                      // ORG-1
	OrganizationUnitCode                       *CE `hl7:"false,Organization Unit Code"`                           // ORG-2
	OrganizationUnitTypeCodeORG                *CE `hl7:"false,Organization Unit Type Code - ORG"`                // ORG-3
	PrimaryOrgUnitIndicator                    *ID `hl7:"false,Primary Org Unit Indicator"`                       // ORG-4
	PractitionerOrgUnitIdentifier              *CX `hl7:"false,Practitioner Org Unit Identifier"`                 // ORG-5
	HealthCareProviderTypeCode                 *CE `hl7:"false,Health Care Provider Type Code"`                   // ORG-6
	HealthCareProviderClassificationCode       *CE `hl7:"false,Health Care Provider Classification Code"`         // ORG-7
	HealthCareProviderAreaOfSpecializationCode *CE `hl7:"false,Health Care Provider Area Of Specialization Code"` // ORG-8
	EffectiveDateRange                         *DR `hl7:"false,Effective Date Range"`                             // ORG-9
	EmploymentStatusCode                       *CE `hl7:"false,Employment Status Code"`                           // ORG-10
	BoardApprovalIndicator                     *ID `hl7:"false,Board Approval Indicator"`                         // ORG-11
	PrimaryCarePhysicianIndicator              *ID `hl7:"false,Primary Care Physician Indicator"`                 // ORG-12
}

func (s *ORG) SegmentName() string {
	return "ORG"
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

// PCR represents the corresponding HL7 segment.
// Definition from HL7 2.4
type PCR struct {
	ImplicatedProduct                 *CE  `hl7:"true,Implicated Product"`                     // PCR-1
	GenericProduct                    *IS  `hl7:"false,Generic Product"`                       // PCR-2
	ProductClass                      *CE  `hl7:"false,Product Class"`                         // PCR-3
	TotalDurationOfTherapy            *CQ  `hl7:"false,Total Duration Of Therapy"`             // PCR-4
	ProductManufactureDate            *TS  `hl7:"false,Product Manufacture Date"`              // PCR-5
	ProductExpirationDate             *TS  `hl7:"false,Product Expiration Date"`               // PCR-6
	ProductImplantationDate           *TS  `hl7:"false,Product Implantation Date"`             // PCR-7
	ProductExplantationDate           *TS  `hl7:"false,Product Explantation Date"`             // PCR-8
	SingleUseDevice                   *IS  `hl7:"false,Single Use Device"`                     // PCR-9
	IndicationForProductUse           *CE  `hl7:"false,Indication For Product Use"`            // PCR-10
	ProductProblem                    *IS  `hl7:"false,Product Problem"`                       // PCR-11
	ProductSerialLotNumber            []ST `hl7:"false,Product Serial/Lot Number"`             // PCR-12
	ProductAvailableForInspection     *IS  `hl7:"false,Product Available For Inspection"`      // PCR-13
	ProductEvaluationPerformed        *CE  `hl7:"false,Product Evaluation Performed"`          // PCR-14
	ProductEvaluationStatus           *CE  `hl7:"false,Product Evaluation Status"`             // PCR-15
	ProductEvaluationResults          *CE  `hl7:"false,Product Evaluation Results"`            // PCR-16
	EvaluatedProductSource            *ID  `hl7:"false,Evaluated Product Source"`              // PCR-17
	DateProductReturnedToManufacturer *TS  `hl7:"false,Date Product Returned To Manufacturer"` // PCR-18
	DeviceOperatorQualifications      *ID  `hl7:"false,Device Operator Qualifications"`        // PCR-19
	RelatednessAssessment             *ID  `hl7:"false,Relatedness Assessment"`                // PCR-20
	ActionTakenInResponseToTheEvent   []ID `hl7:"false,Action Taken In Response To The Event"` // PCR-21
	EventCausalityObservations        []ID `hl7:"false,Event Causality Observations"`          // PCR-22
	IndirectExposureMechanism         []ID `hl7:"false,Indirect Exposure Mechanism"`           // PCR-23
}

func (s *PCR) SegmentName() string {
	return "PCR"
}

// PD1 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type PD1 struct {
	LivingDependency                        []IS  `hl7:"false,Living Dependency"`                           // PD1-1
	LivingArrangement                       *IS   `hl7:"false,Living Arrangement"`                          // PD1-2
	PatientPrimaryFacility                  []XON `hl7:"false,Patient Primary Facility"`                    // PD1-3
	PatientPrimaryCareProviderNameIDNo      []XCN `hl7:"false,Patient Primary Care Provider Name & ID No."` // PD1-4
	StudentIndicator                        *IS   `hl7:"false,Student Indicator"`                           // PD1-5
	Handicap                                *IS   `hl7:"false,Handicap"`                                    // PD1-6
	LivingWillCode                          *IS   `hl7:"false,Living Will Code"`                            // PD1-7
	OrganDonorCode                          *IS   `hl7:"false,Organ Donor Code"`                            // PD1-8
	SeparateBill                            *ID   `hl7:"false,Separate Bill"`                               // PD1-9
	DuplicatePatient                        []CX  `hl7:"false,Duplicate Patient"`                           // PD1-10
	PublicityCode                           *CE   `hl7:"false,Publicity Code"`                              // PD1-11
	ProtectionIndicator                     *ID   `hl7:"false,Protection Indicator"`                        // PD1-12
	ProtectionIndicatorEffectiveDate        *DT   `hl7:"false,Protection Indicator Effective Date"`         // PD1-13
	PlaceOfWorship                          []XON `hl7:"false,Place Of Worship"`                            // PD1-14
	AdvanceDirectiveCode                    []CE  `hl7:"false,Advance Directive Code"`                      // PD1-15
	ImmunizationRegistryStatus              *IS   `hl7:"false,Immunization Registry Status"`                // PD1-16
	ImmunizationRegistryStatusEffectiveDate *DT   `hl7:"false,Immunization Registry Status Effective Date"` // PD1-17
	PublicityCodeEffectiveDate              *DT   `hl7:"false,Publicity Code Effective Date"`               // PD1-18
	MilitaryBranch                          *IS   `hl7:"false,Military Branch"`                             // PD1-19
	MilitaryRankGrade                       *IS   `hl7:"false,Military Rank/Grade"`                         // PD1-20
	MilitaryStatus                          *IS   `hl7:"false,Military Status"`                             // PD1-21
}

func (s *PD1) SegmentName() string {
	return "PD1"
}

// PDA represents the corresponding HL7 segment.
// Definition from HL7 2.4
type PDA struct {
	DeathCauseCode                 []CE `hl7:"false,Death Cause Code"`                   // PDA-1
	DeathLocation                  *PL  `hl7:"false,Death Location"`                     // PDA-2
	DeathCertifiedIndicator        *ID  `hl7:"false,Death Certified Indicator"`          // PDA-3
	DeathCertificateSignedDateTime *TS  `hl7:"false,Death Certificate Signed Date/Time"` // PDA-4
	DeathCertifiedBy               *XCN `hl7:"false,Death Certified By"`                 // PDA-5
	AutopsyIndicator               *ID  `hl7:"false,Autopsy Indicator"`                  // PDA-6
	AutopsyStartAndEndDateTime     *DR  `hl7:"false,Autopsy Start And End Date/Time"`    // PDA-7
	AutopsyPerformedBy             *XCN `hl7:"false,Autopsy Performed By"`               // PDA-8
	CoronerIndicator               *ID  `hl7:"false,Coroner Indicator"`                  // PDA-9
}

func (s *PDA) SegmentName() string {
	return "PDA"
}

// PDC represents the corresponding HL7 segment.
// Definition from HL7 2.4
type PDC struct {
	ManufacturerDistributor []XON `hl7:"true,Manufacturer/Distributor"` // PDC-1
	Country                 *CE   `hl7:"true,Country"`                  // PDC-2
	BrandName               *ST   `hl7:"true,Brand Name"`               // PDC-3
	DeviceFamilyName        *ST   `hl7:"false,Device Family Name"`      // PDC-4
	GenericName             *CE   `hl7:"false,Generic Name"`            // PDC-5
	ModelIdentifier         []ST  `hl7:"false,Model Identifier"`        // PDC-6
	CatalogueIdentifier     *ST   `hl7:"false,Catalogue Identifier"`    // PDC-7
	OtherIdentifier         []ST  `hl7:"false,Other Identifier"`        // PDC-8
	ProductCode             *CE   `hl7:"false,Product Code"`            // PDC-9
	MarketingBasis          *ID   `hl7:"false,Marketing Basis"`         // PDC-10
	MarketingApprovalID     *ST   `hl7:"false,Marketing Approval ID"`   // PDC-11
	LabeledShelfLife        *CQ   `hl7:"false,Labeled Shelf Life"`      // PDC-12
	ExpectedShelfLife       *CQ   `hl7:"false,Expected Shelf Life"`     // PDC-13
	DateFirstMarketed       *TS   `hl7:"false,Date First Marketed"`     // PDC-14
	DateLastMarketed        *TS   `hl7:"false,Date Last Marketed"`      // PDC-15
}

func (s *PDC) SegmentName() string {
	return "PDC"
}

// PEO represents the corresponding HL7 segment.
// Definition from HL7 2.4
type PEO struct {
	EventIdentifiersUsed                  []CE  `hl7:"false,Event Identifiers Used"`                      // PEO-1
	EventSymptomDiagnosisCode             []CE  `hl7:"false,Event Symptom/Diagnosis Code"`                // PEO-2
	EventOnsetDateTime                    *TS   `hl7:"true,Event Onset Date/Time"`                        // PEO-3
	EventExacerbationDateTime             *TS   `hl7:"false,Event Exacerbation Date/Time"`                // PEO-4
	EventImprovedDateTime                 *TS   `hl7:"false,Event Improved Date/Time"`                    // PEO-5
	EventEndedDataTime                    *TS   `hl7:"false,Event Ended Data/Time"`                       // PEO-6
	EventLocationOccurredAddress          []XAD `hl7:"false,Event Location Occurred Address"`             // PEO-7
	EventQualification                    []ID  `hl7:"false,Event Qualification"`                         // PEO-8
	EventSerious                          *ID   `hl7:"false,Event Serious"`                               // PEO-9
	EventExpected                         *ID   `hl7:"false,Event Expected"`                              // PEO-10
	EventOutcome                          []ID  `hl7:"false,Event Outcome"`                               // PEO-11
	PatientOutcome                        *ID   `hl7:"false,Patient Outcome"`                             // PEO-12
	EventDescriptionFromOthers            []FT  `hl7:"false,Event Description From Others"`               // PEO-13
	EventFromOriginalReporter             []FT  `hl7:"false,Event From Original Reporter"`                // PEO-14
	EventDescriptionFromPatient           []FT  `hl7:"false,Event Description From Patient"`              // PEO-15
	EventDescriptionFromPractitioner      []FT  `hl7:"false,Event Description From Practitioner"`         // PEO-16
	EventDescriptionFromAutopsy           []FT  `hl7:"false,Event Description From Autopsy"`              // PEO-17
	CauseOfDeath                          []CE  `hl7:"false,Cause Of Death"`                              // PEO-18
	PrimaryObserverName                   []XPN `hl7:"false,Primary Observer Name"`                       // PEO-19
	PrimaryObserverAddress                []XAD `hl7:"false,Primary Observer Address"`                    // PEO-20
	PrimaryObserverTelephone              []XTN `hl7:"false,Primary Observer Telephone"`                  // PEO-21
	PrimaryObserverSQualification         *ID   `hl7:"false,Primary Observer'S Qualification"`            // PEO-22
	ConfirmationProvidedBy                *ID   `hl7:"false,Confirmation Provided By"`                    // PEO-23
	PrimaryObserverAwareDateTime          *TS   `hl7:"false,Primary Observer Aware Date/Time"`            // PEO-24
	PrimaryObserverSIdentityMayBeDivulged *ID   `hl7:"false,Primary Observer'S Identity May Be Divulged"` // PEO-25
}

func (s *PEO) SegmentName() string {
	return "PEO"
}

// PES represents the corresponding HL7 segment.
// Definition from HL7 2.4
type PES struct {
	SenderOrganizationName []XON `hl7:"false,Sender Organization Name"` // PES-1
	SenderIndividualName   []XCN `hl7:"false,Sender Individual Name"`   // PES-2
	SenderAddress          []XAD `hl7:"false,Sender Address"`           // PES-3
	SenderTelephone        []XTN `hl7:"false,Sender Telephone"`         // PES-4
	SenderEventIdentifier  *EI   `hl7:"false,Sender Event Identifier"`  // PES-5
	SenderSequenceNumber   *NM   `hl7:"false,Sender Sequence Number"`   // PES-6
	SenderEventDescription []FT  `hl7:"false,Sender Event Description"` // PES-7
	SenderComment          *FT   `hl7:"false,Sender Comment"`           // PES-8
	SenderAwareDateTime    *TS   `hl7:"false,Sender Aware Date/Time"`   // PES-9
	EventReportDate        *TS   `hl7:"true,Event Report Date"`         // PES-10
	EventReportTimingType  []ID  `hl7:"false,Event Report Timing/Type"` // PES-11
	EventReportSource      *ID   `hl7:"false,Event Report Source"`      // PES-12
	EventReportedTo        []ID  `hl7:"false,Event Reported To"`        // PES-13
}

func (s *PES) SegmentName() string {
	return "PES"
}

// PID represents the corresponding HL7 segment.
// Definition from HL7 2.4
type PID struct {
	SetIDPID                    *SI   `hl7:"false,Set ID - PID"`                      // PID-1
	PatientID                   *CX   `hl7:"false,Patient ID"`                        // PID-2
	PatientIdentifierList       []CX  `hl7:"true,Patient Identifier List"`            // PID-3
	AlternatePatientIDPID       []CX  `hl7:"false,Alternate Patient ID - PID"`        // PID-4
	PatientName                 []XPN `hl7:"true,Patient Name"`                       // PID-5
	MotherSMaidenName           []XPN `hl7:"false,Mother'S Maiden Name"`              // PID-6
	DateTimeOfBirth             *TS   `hl7:"false,Date/Time Of Birth"`                // PID-7
	AdministrativeSex           *IS   `hl7:"false,Administrative Sex"`                // PID-8
	PatientAlias                []XPN `hl7:"false,Patient Alias"`                     // PID-9
	Race                        []CE  `hl7:"false,Race"`                              // PID-10
	PatientAddress              []XAD `hl7:"false,Patient Address"`                   // PID-11
	CountyCode                  *IS   `hl7:"false,County Code"`                       // PID-12
	PhoneNumberHome             []XTN `hl7:"false,Phone Number - Home"`               // PID-13
	PhoneNumberBusiness         []XTN `hl7:"false,Phone Number - Business"`           // PID-14
	PrimaryLanguage             *CE   `hl7:"false,Primary Language"`                  // PID-15
	MaritalStatus               *CE   `hl7:"false,Marital Status"`                    // PID-16
	Religion                    *CE   `hl7:"false,Religion"`                          // PID-17
	PatientAccountNumber        *CX   `hl7:"false,Patient Account Number"`            // PID-18
	SSNNumberPatient            *ST   `hl7:"false,SSN Number - Patient"`              // PID-19
	DriverSLicenseNumberPatient *DLN  `hl7:"false,Driver'S License Number - Patient"` // PID-20
	MotherSIdentifier           []CX  `hl7:"false,Mother'S Identifier"`               // PID-21
	EthnicGroup                 []CE  `hl7:"false,Ethnic Group"`                      // PID-22
	BirthPlace                  *ST   `hl7:"false,Birth Place"`                       // PID-23
	MultipleBirthIndicator      *ID   `hl7:"false,Multiple Birth Indicator"`          // PID-24
	BirthOrder                  *NM   `hl7:"false,Birth Order"`                       // PID-25
	Citizenship                 []CE  `hl7:"false,Citizenship"`                       // PID-26
	VeteransMilitaryStatus      *CE   `hl7:"false,Veterans Military Status"`          // PID-27
	Nationality                 *CE   `hl7:"false,Nationality"`                       // PID-28
	PatientDeathDateAndTime     *TS   `hl7:"false,Patient Death Date And Time"`       // PID-29
	PatientDeathIndicator       *ID   `hl7:"false,Patient Death Indicator"`           // PID-30
	IdentityUnknownIndicator    *ID   `hl7:"false,Identity Unknown Indicator"`        // PID-31
	IdentityReliabilityCode     []IS  `hl7:"false,Identity Reliability Code"`         // PID-32
	LastUpdateDateTime          *TS   `hl7:"false,Last Update Date/Time"`             // PID-33
	LastUpdateFacility          *HD   `hl7:"false,Last Update Facility"`              // PID-34
	SpeciesCode                 *CE   `hl7:"false,Species Code"`                      // PID-35
	BreedCode                   *CE   `hl7:"false,Breed Code"`                        // PID-36
	Strain                      *ST   `hl7:"false,Strain"`                            // PID-37
	ProductionClassCode         *CE   `hl7:"false,Production Class Code"`             // PID-38
}

func (s *PID) SegmentName() string {
	return "PID"
}

// PR1 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type PR1 struct {
	SetIDPR1                *SI   `hl7:"true,Set ID - PR1"`               // PR1-1
	ProcedureCodingMethod   *IS   `hl7:"false,Procedure Coding Method"`   // PR1-2
	ProcedureCode           *CE   `hl7:"true,Procedure Code"`             // PR1-3
	ProcedureDescription    *ST   `hl7:"false,Procedure Description"`     // PR1-4
	ProcedureDateTime       *TS   `hl7:"true,Procedure Date/Time"`        // PR1-5
	ProcedureFunctionalType *IS   `hl7:"false,Procedure Functional Type"` // PR1-6
	ProcedureMinutes        *NM   `hl7:"false,Procedure Minutes"`         // PR1-7
	Anesthesiologist        []XCN `hl7:"false,Anesthesiologist"`          // PR1-8
	AnesthesiaCode          *IS   `hl7:"false,Anesthesia Code"`           // PR1-9
	AnesthesiaMinutes       *NM   `hl7:"false,Anesthesia Minutes"`        // PR1-10
	Surgeon                 []XCN `hl7:"false,Surgeon"`                   // PR1-11
	ProcedurePractitioner   []XCN `hl7:"false,Procedure Practitioner"`    // PR1-12
	ConsentCode             *CE   `hl7:"false,Consent Code"`              // PR1-13
	ProcedurePriority       *ID   `hl7:"false,Procedure Priority"`        // PR1-14
	AssociatedDiagnosisCode *CE   `hl7:"false,Associated Diagnosis Code"` // PR1-15
	ProcedureCodeModifier   []CE  `hl7:"false,Procedure Code Modifier"`   // PR1-16
	ProcedureDRGType        *IS   `hl7:"false,Procedure DRG Type"`        // PR1-17
	TissueTypeCode          []CE  `hl7:"false,Tissue Type Code"`          // PR1-18
}

func (s *PR1) SegmentName() string {
	return "PR1"
}

// PRA represents the corresponding HL7 segment.
// Definition from HL7 2.4
type PRA struct {
	PrimaryKeyValuePRA                        *CE   `hl7:"false,Primary Key Value - PRA"`                      // PRA-1
	PractitionerGroup                         []CE  `hl7:"false,Practitioner Group"`                           // PRA-2
	PractitionerCategory                      []IS  `hl7:"false,Practitioner Category"`                        // PRA-3
	ProviderBilling                           *ID   `hl7:"false,Provider Billing"`                             // PRA-4
	Specialty                                 []SPD `hl7:"false,Specialty"`                                    // PRA-5
	PractitionerIDNumbers                     []PLN `hl7:"false,Practitioner ID Numbers"`                      // PRA-6
	Privileges                                []PIP `hl7:"false,Privileges"`                                   // PRA-7
	DateEnteredPractice                       *DT   `hl7:"false,Date Entered Practice"`                        // PRA-8
	Institution                               *CE   `hl7:"false,Institution"`                                  // PRA-9
	DateLeftPractice                          *DT   `hl7:"false,Date Left Practice"`                           // PRA-10
	GovernmentReimbursementBillingEligibility []CE  `hl7:"false,Government Reimbursement Billing Eligibility"` // PRA-11
	SetIDPRA                                  *SI   `hl7:"false,Set ID - PRA"`                                 // PRA-12
}

func (s *PRA) SegmentName() string {
	return "PRA"
}

// PRB represents the corresponding HL7 segment.
// Definition from HL7 2.4
type PRB struct {
	ActionCode                                        *ID  `hl7:"true,Action Code"`                                              // PRB-1
	ActionDateTime                                    *TS  `hl7:"true,Action Date/Time"`                                         // PRB-2
	ProblemID                                         *CE  `hl7:"true,Problem ID"`                                               // PRB-3
	ProblemInstanceID                                 *EI  `hl7:"true,Problem Instance ID"`                                      // PRB-4
	EpisodeOfCareID                                   *EI  `hl7:"false,Episode Of Care ID"`                                      // PRB-5
	ProblemListPriority                               *NM  `hl7:"false,Problem List Priority"`                                   // PRB-6
	ProblemEstablishedDateTime                        *TS  `hl7:"false,Problem Established Date/Time"`                           // PRB-7
	AnticipatedProblemResolutionDateTime              *TS  `hl7:"false,Anticipated Problem Resolution Date/Time"`                // PRB-8
	ActualProblemResolutionDateTime                   *TS  `hl7:"false,Actual Problem Resolution Date/Time"`                     // PRB-9
	ProblemClassification                             *CE  `hl7:"false,Problem Classification"`                                  // PRB-10
	ProblemManagementDiscipline                       []CE `hl7:"false,Problem Management Discipline"`                           // PRB-11
	ProblemPersistence                                *CE  `hl7:"false,Problem Persistence"`                                     // PRB-12
	ProblemConfirmationStatus                         *CE  `hl7:"false,Problem Confirmation Status"`                             // PRB-13
	ProblemLifeCycleStatus                            *CE  `hl7:"false,Problem Life Cycle Status"`                               // PRB-14
	ProblemLifeCycleStatusDateTime                    *TS  `hl7:"false,Problem Life Cycle Status Date/Time"`                     // PRB-15
	ProblemDateOfOnset                                *TS  `hl7:"false,Problem Date Of Onset"`                                   // PRB-16
	ProblemOnsetText                                  *ST  `hl7:"false,Problem Onset Text"`                                      // PRB-17
	ProblemRanking                                    *CE  `hl7:"false,Problem Ranking"`                                         // PRB-18
	CertaintyOfProblem                                *CE  `hl7:"false,Certainty Of Problem"`                                    // PRB-19
	ProbabilityOfProblem01                            *NM  `hl7:"false,Probability Of Problem (0-1)"`                            // PRB-20
	IndividualAwarenessOfProblem                      *CE  `hl7:"false,Individual Awareness Of Problem"`                         // PRB-21
	ProblemPrognosis                                  *CE  `hl7:"false,Problem Prognosis"`                                       // PRB-22
	IndividualAwarenessOfPrognosis                    *CE  `hl7:"false,Individual Awareness Of Prognosis"`                       // PRB-23
	FamilySignificantOtherAwarenessOfProblemPrognosis *ST  `hl7:"false,Family/Significant Other Awareness Of Problem/Prognosis"` // PRB-24
	SecuritySensitivity                               *CE  `hl7:"false,Security/Sensitivity"`                                    // PRB-25
}

func (s *PRB) SegmentName() string {
	return "PRB"
}

// PRC represents the corresponding HL7 segment.
// Definition from HL7 2.4
type PRC struct {
	PrimaryKeyValuePRC  *CE  `hl7:"true,Primary Key Value - PRC"` // PRC-1
	FacilityIDPRC       []CE `hl7:"false,Facility ID - PRC"`      // PRC-2
	Department          []CE `hl7:"false,Department"`             // PRC-3
	ValidPatientClasses []IS `hl7:"false,Valid Patient Classes"`  // PRC-4
	Price               []CP `hl7:"false,Price"`                  // PRC-5
	Formula             []ST `hl7:"false,Formula"`                // PRC-6
	MinimumQuantity     *NM  `hl7:"false,Minimum Quantity"`       // PRC-7
	MaximumQuantity     *NM  `hl7:"false,Maximum Quantity"`       // PRC-8
	MinimumPrice        *MO  `hl7:"false,Minimum Price"`          // PRC-9
	MaximumPrice        *MO  `hl7:"false,Maximum Price"`          // PRC-10
	EffectiveStartDate  *TS  `hl7:"false,Effective Start Date"`   // PRC-11
	EffectiveEndDate    *TS  `hl7:"false,Effective End Date"`     // PRC-12
	PriceOverrideFlag   *IS  `hl7:"false,Price Override Flag"`    // PRC-13
	BillingCategory     []CE `hl7:"false,Billing Category"`       // PRC-14
	ChargeableFlag      *ID  `hl7:"false,Chargeable Flag"`        // PRC-15
	ActiveInactiveFlag  *ID  `hl7:"false,Active/Inactive Flag"`   // PRC-16
	Cost                *MO  `hl7:"false,Cost"`                   // PRC-17
	ChargeOnIndicator   *IS  `hl7:"false,Charge On Indicator"`    // PRC-18
}

func (s *PRC) SegmentName() string {
	return "PRC"
}

// PRD represents the corresponding HL7 segment.
// Definition from HL7 2.4
type PRD struct {
	ProviderRole                     []CE  `hl7:"true,Provider Role"`                          // PRD-1
	ProviderName                     []XPN `hl7:"false,Provider Name"`                         // PRD-2
	ProviderAddress                  []XAD `hl7:"false,Provider Address"`                      // PRD-3
	ProviderLocation                 *PL   `hl7:"false,Provider Location"`                     // PRD-4
	ProviderCommunicationInformation []XTN `hl7:"false,Provider Communication Information"`    // PRD-5
	PreferredMethodOfContact         *CE   `hl7:"false,Preferred Method Of Contact"`           // PRD-6
	ProviderIdentifiers              []PI  `hl7:"false,Provider Identifiers"`                  // PRD-7
	EffectiveStartDateOfProviderRole *TS   `hl7:"false,Effective Start Date Of Provider Role"` // PRD-8
	EffectiveEndDateOfProviderRole   *TS   `hl7:"false,Effective End Date Of Provider Role"`   // PRD-9
}

func (s *PRD) SegmentName() string {
	return "PRD"
}

// PSH represents the corresponding HL7 segment.
// Definition from HL7 2.4
type PSH struct {
	ReportType                                         *ST  `hl7:"true,Report Type"`                                                // PSH-1
	ReportFormIdentifier                               *ST  `hl7:"false,Report Form Identifier"`                                    // PSH-2
	ReportDate                                         *TS  `hl7:"true,Report Date"`                                                // PSH-3
	ReportIntervalStartDate                            *TS  `hl7:"false,Report Interval Start Date"`                                // PSH-4
	ReportIntervalEndDate                              *TS  `hl7:"false,Report Interval End Date"`                                  // PSH-5
	QuantityManufactured                               *CQ  `hl7:"false,Quantity Manufactured"`                                     // PSH-6
	QuantityDistributed                                *CQ  `hl7:"false,Quantity Distributed"`                                      // PSH-7
	QuantityDistributedMethod                          *ID  `hl7:"false,Quantity Distributed Method"`                               // PSH-8
	QuantityDistributedComment                         *FT  `hl7:"false,Quantity Distributed Comment"`                              // PSH-9
	QuantityInUse                                      *CQ  `hl7:"false,Quantity In Use"`                                           // PSH-10
	QuantityInUseMethod                                *ID  `hl7:"false,Quantity In Use Method"`                                    // PSH-11
	QuantityInUseComment                               *FT  `hl7:"false,Quantity In Use Comment"`                                   // PSH-12
	NumberOfProductExperienceReportsFiledByFacility    []NM `hl7:"false,Number Of Product Experience Reports Filed By Facility"`    // PSH-13
	NumberOfProductExperienceReportsFiledByDistributor []NM `hl7:"false,Number Of Product Experience Reports Filed By Distributor"` // PSH-14
}

func (s *PSH) SegmentName() string {
	return "PSH"
}

// PTH represents the corresponding HL7 segment.
// Definition from HL7 2.4
type PTH struct {
	ActionCode                           *ID `hl7:"true,Action Code"`                                 // PTH-1
	PathwayID                            *CE `hl7:"true,Pathway ID"`                                  // PTH-2
	PathwayInstanceID                    *EI `hl7:"true,Pathway Instance ID"`                         // PTH-3
	PathwayEstablishedDateTime           *TS `hl7:"true,Pathway Established Date/Time"`               // PTH-4
	PathwayLifeCycleStatus               *CE `hl7:"false,Pathway Life Cycle Status"`                  // PTH-5
	ChangePathwayLifeCycleStatusDateTime *TS `hl7:"false,Change Pathway Life Cycle Status Date/Time"` // PTH-6
}

func (s *PTH) SegmentName() string {
	return "PTH"
}

// PV1 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type PV1 struct {
	SetIDPV1                *SI   `hl7:"false,Set ID - PV1"`              // PV1-1
	PatientClass            *IS   `hl7:"true,Patient Class"`              // PV1-2
	AssignedPatientLocation *PL   `hl7:"false,Assigned Patient Location"` // PV1-3
	AdmissionType           *IS   `hl7:"false,Admission Type"`            // PV1-4
	PreadmitNumber          *CX   `hl7:"false,Preadmit Number"`           // PV1-5
	PriorPatientLocation    *PL   `hl7:"false,Prior Patient Location"`    // PV1-6
	AttendingDoctor         []XCN `hl7:"false,Attending Doctor"`          // PV1-7
	ReferringDoctor         []XCN `hl7:"false,Referring Doctor"`          // PV1-8
	ConsultingDoctor        []XCN `hl7:"false,Consulting Doctor"`         // PV1-9
	HospitalService         *IS   `hl7:"false,Hospital Service"`          // PV1-10
	TemporaryLocation       *PL   `hl7:"false,Temporary Location"`        // PV1-11
	PreadmitTestIndicator   *IS   `hl7:"false,Preadmit Test Indicator"`   // PV1-12
	ReAdmissionIndicator    *IS   `hl7:"false,Re-Admission Indicator"`    // PV1-13
	AdmitSource             *IS   `hl7:"false,Admit Source"`              // PV1-14
	AmbulatoryStatus        []IS  `hl7:"false,Ambulatory Status"`         // PV1-15
	VIPIndicator            *IS   `hl7:"false,VIP Indicator"`             // PV1-16
	AdmittingDoctor         []XCN `hl7:"false,Admitting Doctor"`          // PV1-17
	PatientType             *IS   `hl7:"false,Patient Type"`              // PV1-18
	VisitNumber             *CX   `hl7:"false,Visit Number"`              // PV1-19
	FinancialClass          []FC  `hl7:"false,Financial Class"`           // PV1-20
	ChargePriceIndicator    *IS   `hl7:"false,Charge Price Indicator"`    // PV1-21
	CourtesyCode            *IS   `hl7:"false,Courtesy Code"`             // PV1-22
	CreditRating            *IS   `hl7:"false,Credit Rating"`             // PV1-23
	ContractCode            []IS  `hl7:"false,Contract Code"`             // PV1-24
	ContractEffectiveDate   []DT  `hl7:"false,Contract Effective Date"`   // PV1-25
	ContractAmount          []NM  `hl7:"false,Contract Amount"`           // PV1-26
	ContractPeriod          []NM  `hl7:"false,Contract Period"`           // PV1-27
	InterestCode            *IS   `hl7:"false,Interest Code"`             // PV1-28
	TransferToBadDebtCode   *IS   `hl7:"false,Transfer To Bad Debt Code"` // PV1-29
	TransferToBadDebtDate   *DT   `hl7:"false,Transfer To Bad Debt Date"` // PV1-30
	BadDebtAgencyCode       *IS   `hl7:"false,Bad Debt Agency Code"`      // PV1-31
	BadDebtTransferAmount   *NM   `hl7:"false,Bad Debt Transfer Amount"`  // PV1-32
	BadDebtRecoveryAmount   *NM   `hl7:"false,Bad Debt Recovery Amount"`  // PV1-33
	DeleteAccountIndicator  *IS   `hl7:"false,Delete Account Indicator"`  // PV1-34
	DeleteAccountDate       *DT   `hl7:"false,Delete Account Date"`       // PV1-35
	DischargeDisposition    *IS   `hl7:"false,Discharge Disposition"`     // PV1-36
	DischargedToLocation    *DLD  `hl7:"false,Discharged To Location"`    // PV1-37
	DietType                *CE   `hl7:"false,Diet Type"`                 // PV1-38
	ServicingFacility       *IS   `hl7:"false,Servicing Facility"`        // PV1-39
	BedStatus               *IS   `hl7:"false,Bed Status"`                // PV1-40
	AccountStatus           *IS   `hl7:"false,Account Status"`            // PV1-41
	PendingLocation         *PL   `hl7:"false,Pending Location"`          // PV1-42
	PriorTemporaryLocation  *PL   `hl7:"false,Prior Temporary Location"`  // PV1-43
	AdmitDateTime           *TS   `hl7:"false,Admit Date/Time"`           // PV1-44
	DischargeDateTime       []TS  `hl7:"false,Discharge Date/Time"`       // PV1-45
	CurrentPatientBalance   *NM   `hl7:"false,Current Patient Balance"`   // PV1-46
	TotalCharges            *NM   `hl7:"false,Total Charges"`             // PV1-47
	TotalAdjustments        *NM   `hl7:"false,Total Adjustments"`         // PV1-48
	TotalPayments           *NM   `hl7:"false,Total Payments"`            // PV1-49
	AlternateVisitID        *CX   `hl7:"false,Alternate Visit ID"`        // PV1-50
	VisitIndicator          *IS   `hl7:"false,Visit Indicator"`           // PV1-51
	OtherHealthcareProvider []XCN `hl7:"false,Other Healthcare Provider"` // PV1-52
}

func (s *PV1) SegmentName() string {
	return "PV1"
}

// PV2 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type PV2 struct {
	PriorPendingLocation              *PL   `hl7:"false,Prior Pending Location"`               // PV2-1
	AccommodationCode                 *CE   `hl7:"false,Accommodation Code"`                   // PV2-2
	AdmitReason                       *CE   `hl7:"false,Admit Reason"`                         // PV2-3
	TransferReason                    *CE   `hl7:"false,Transfer Reason"`                      // PV2-4
	PatientValuables                  []ST  `hl7:"false,Patient Valuables"`                    // PV2-5
	PatientValuablesLocation          *ST   `hl7:"false,Patient Valuables Location"`           // PV2-6
	VisitUserCode                     []IS  `hl7:"false,Visit User Code"`                      // PV2-7
	ExpectedAdmitDateTime             *TS   `hl7:"false,Expected Admit Date/Time"`             // PV2-8
	ExpectedDischargeDateTime         *TS   `hl7:"false,Expected Discharge Date/Time"`         // PV2-9
	EstimatedLengthOfInpatientStay    *NM   `hl7:"false,Estimated Length Of Inpatient Stay"`   // PV2-10
	ActualLengthOfInpatientStay       *NM   `hl7:"false,Actual Length Of Inpatient Stay"`      // PV2-11
	VisitDescription                  *ST   `hl7:"false,Visit Description"`                    // PV2-12
	ReferralSourceCode                []XCN `hl7:"false,Referral Source Code"`                 // PV2-13
	PreviousServiceDate               *DT   `hl7:"false,Previous Service Date"`                // PV2-14
	EmploymentIllnessRelatedIndicator *ID   `hl7:"false,Employment Illness Related Indicator"` // PV2-15
	PurgeStatusCode                   *IS   `hl7:"false,Purge Status Code"`                    // PV2-16
	PurgeStatusDate                   *DT   `hl7:"false,Purge Status Date"`                    // PV2-17
	SpecialProgramCode                *IS   `hl7:"false,Special Program Code"`                 // PV2-18
	RetentionIndicator                *ID   `hl7:"false,Retention Indicator"`                  // PV2-19
	ExpectedNumberOfInsurancePlans    *NM   `hl7:"false,Expected Number Of Insurance Plans"`   // PV2-20
	VisitPublicityCode                *IS   `hl7:"false,Visit Publicity Code"`                 // PV2-21
	VisitProtectionIndicator          *ID   `hl7:"false,Visit Protection Indicator"`           // PV2-22
	ClinicOrganizationName            []XON `hl7:"false,Clinic Organization Name"`             // PV2-23
	PatientStatusCode                 *IS   `hl7:"false,Patient Status Code"`                  // PV2-24
	VisitPriorityCode                 *IS   `hl7:"false,Visit Priority Code"`                  // PV2-25
	PreviousTreatmentDate             *DT   `hl7:"false,Previous Treatment Date"`              // PV2-26
	ExpectedDischargeDisposition      *IS   `hl7:"false,Expected Discharge Disposition"`       // PV2-27
	SignatureOnFileDate               *DT   `hl7:"false,Signature On File Date"`               // PV2-28
	FirstSimilarIllnessDate           *DT   `hl7:"false,First Similar Illness Date"`           // PV2-29
	PatientChargeAdjustmentCode       *CE   `hl7:"false,Patient Charge Adjustment Code"`       // PV2-30
	RecurringServiceCode              *IS   `hl7:"false,Recurring Service Code"`               // PV2-31
	BillingMediaCode                  *ID   `hl7:"false,Billing Media Code"`                   // PV2-32
	ExpectedSurgeryDateAndTime        *TS   `hl7:"false,Expected Surgery Date And Time"`       // PV2-33
	MilitaryPartnershipCode           *ID   `hl7:"false,Military Partnership Code"`            // PV2-34
	MilitaryNonAvailabilityCode       *ID   `hl7:"false,Military Non-Availability Code"`       // PV2-35
	NewbornBabyIndicator              *ID   `hl7:"false,Newborn Baby Indicator"`               // PV2-36
	BabyDetainedIndicator             *ID   `hl7:"false,Baby Detained Indicator"`              // PV2-37
	ModeOfArrivalCode                 *CE   `hl7:"false,Mode Of Arrival Code"`                 // PV2-38
	RecreationalDrugUseCode           []CE  `hl7:"false,Recreational Drug Use Code"`           // PV2-39
	AdmissionLevelOfCareCode          *CE   `hl7:"false,Admission Level Of Care Code"`         // PV2-40
	PrecautionCode                    []CE  `hl7:"false,Precaution Code"`                      // PV2-41
	PatientConditionCode              *CE   `hl7:"false,Patient Condition Code"`               // PV2-42
	LivingWillCode                    *IS   `hl7:"false,Living Will Code"`                     // PV2-43
	OrganDonorCode                    *IS   `hl7:"false,Organ Donor Code"`                     // PV2-44
	AdvanceDirectiveCode              []CE  `hl7:"false,Advance Directive Code"`               // PV2-45
	PatientStatusEffectiveDate        *DT   `hl7:"false,Patient Status Effective Date"`        // PV2-46
	ExpectedLOAReturnDateTime         *TS   `hl7:"false,Expected LOA Return Date/Time"`        // PV2-47
}

func (s *PV2) SegmentName() string {
	return "PV2"
}

// QAK represents the corresponding HL7 segment.
// Definition from HL7 2.4
type QAK struct {
	QueryTag            *ST `hl7:"false,Query Tag"`             // QAK-1
	QueryResponseStatus *ID `hl7:"false,Query Response Status"` // QAK-2
	MessageQueryName    *CE `hl7:"false,Message Query Name"`    // QAK-3
	HitCountTotal       *NM `hl7:"false,Hit Count Total"`       // QAK-4
	ThisPayload         *NM `hl7:"false,This Payload"`          // QAK-5
	HitsRemaining       *NM `hl7:"false,Hits Remaining"`        // QAK-6
}

func (s *QAK) SegmentName() string {
	return "QAK"
}

// QID represents the corresponding HL7 segment.
// Definition from HL7 2.4
type QID struct {
	QueryTag         *ST `hl7:"true,Query Tag"`          // QID-1
	MessageQueryName *CE `hl7:"true,Message Query Name"` // QID-2
}

func (s *QID) SegmentName() string {
	return "QID"
}

// QPD represents the corresponding HL7 segment.
// Definition from HL7 2.4
type QPD struct {
	MessageQueryName                 *CE  `hl7:"true,Message Query Name"`                      // QPD-1
	QueryTag                         *ST  `hl7:"false,Query Tag"`                              // QPD-2
	UserParametersInSuccessiveFields *Any `hl7:"false,User Parameters (In Successive Fields)"` // QPD-3
}

func (s *QPD) SegmentName() string {
	return "QPD"
}

// QRD represents the corresponding HL7 segment.
// Definition from HL7 2.4
type QRD struct {
	QueryDateTime            *TS   `hl7:"true,Query Date/Time"`              // QRD-1
	QueryFormatCode          *ID   `hl7:"true,Query Format Code"`            // QRD-2
	QueryPriority            *ID   `hl7:"true,Query Priority"`               // QRD-3
	QueryID                  *ST   `hl7:"true,Query ID"`                     // QRD-4
	DeferredResponseType     *ID   `hl7:"false,Deferred Response Type"`      // QRD-5
	DeferredResponseDateTime *TS   `hl7:"false,Deferred Response Date/Time"` // QRD-6
	QuantityLimitedRequest   *CQ   `hl7:"true,Quantity Limited Request"`     // QRD-7
	WhoSubjectFilter         []XCN `hl7:"true,Who Subject Filter"`           // QRD-8
	WhatSubjectFilter        []CE  `hl7:"true,What Subject Filter"`          // QRD-9
	WhatDepartmentDataCode   []CE  `hl7:"true,What Department Data Code"`    // QRD-10
	WhatDataCodeValueQual    []VR  `hl7:"false,What Data Code Value Qual."`  // QRD-11
	QueryResultsLevel        *ID   `hl7:"false,Query Results Level"`         // QRD-12
}

func (s *QRD) SegmentName() string {
	return "QRD"
}

// QRF represents the corresponding HL7 segment.
// Definition from HL7 2.4
type QRF struct {
	WhereSubjectFilter           []ST `hl7:"true,Where Subject Filter"`              // QRF-1
	WhenDataStartDateTime        *TS  `hl7:"false,When Data Start Date/Time"`        // QRF-2
	WhenDataEndDateTime          *TS  `hl7:"false,When Data End Date/Time"`          // QRF-3
	WhatUserQualifier            []ST `hl7:"false,What User Qualifier"`              // QRF-4
	OtherQRYSubjectFilter        []ST `hl7:"false,Other QRY Subject Filter"`         // QRF-5
	WhichDateTimeQualifier       []ID `hl7:"false,Which Date/Time Qualifier"`        // QRF-6
	WhichDateTimeStatusQualifier []ID `hl7:"false,Which Date/Time Status Qualifier"` // QRF-7
	DateTimeSelectionQualifier   []ID `hl7:"false,Date/Time Selection Qualifier"`    // QRF-8
	WhenQuantityTimingQualifier  *TQ  `hl7:"false,When Quantity/Timing Qualifier"`   // QRF-9
	SearchConfidenceThreshold    *NM  `hl7:"false,Search Confidence Threshold"`      // QRF-10
}

func (s *QRF) SegmentName() string {
	return "QRF"
}

// QRI represents the corresponding HL7 segment.
// Definition from HL7 2.4
type QRI struct {
	CandidateConfidence *NM  `hl7:"false,Candidate Confidence"` // QRI-1
	MatchReasonCode     []IS `hl7:"false,Match Reason Code"`    // QRI-2
	AlgorithmDescriptor *CE  `hl7:"false,Algorithm Descriptor"` // QRI-3
}

func (s *QRI) SegmentName() string {
	return "QRI"
}

// RCP represents the corresponding HL7 segment.
// Definition from HL7 2.4
type RCP struct {
	QueryPriority            *ID   `hl7:"false,Query Priority"`              // RCP-1
	QuantityLimitedRequest   *CQ   `hl7:"false,Quantity Limited Request"`    // RCP-2
	ResponseModality         *CE   `hl7:"false,Response Modality"`           // RCP-3
	ExecutionAndDeliveryTime *TS   `hl7:"false,Execution And Delivery Time"` // RCP-4
	ModifyIndicator          *ID   `hl7:"false,Modify Indicator"`            // RCP-5
	SortByField              []SRT `hl7:"false,Sort-By Field"`               // RCP-6
	SegmentGroupInclusion    []ID  `hl7:"false,Segment Group Inclusion"`     // RCP-7
}

func (s *RCP) SegmentName() string {
	return "RCP"
}

// RDF represents the corresponding HL7 segment.
// Definition from HL7 2.4
type RDF struct {
	NumberOfColumnsPerRow *NM   `hl7:"true,Number Of Columns Per Row"` // RDF-1
	ColumnDescription     []RCD `hl7:"true,Column Description"`        // RDF-2
}

func (s *RDF) SegmentName() string {
	return "RDF"
}

// RDT represents the corresponding HL7 segment.
// Definition from HL7 2.4
type RDT struct {
	ColumnValue *Any `hl7:"true,Column Value"` // RDT-1
}

func (s *RDT) SegmentName() string {
	return "RDT"
}

// RF1 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type RF1 struct {
	ReferralStatus                *CE  `hl7:"false,Referral Status"`                // RF1-1
	ReferralPriority              *CE  `hl7:"false,Referral Priority"`              // RF1-2
	ReferralType                  *CE  `hl7:"false,Referral Type"`                  // RF1-3
	ReferralDisposition           []CE `hl7:"false,Referral Disposition"`           // RF1-4
	ReferralCategory              *CE  `hl7:"false,Referral Category"`              // RF1-5
	OriginatingReferralIdentifier *EI  `hl7:"true,Originating Referral Identifier"` // RF1-6
	EffectiveDate                 *TS  `hl7:"false,Effective Date"`                 // RF1-7
	ExpirationDate                *TS  `hl7:"false,Expiration Date"`                // RF1-8
	ProcessDate                   *TS  `hl7:"false,Process Date"`                   // RF1-9
	ReferralReason                []CE `hl7:"false,Referral Reason"`                // RF1-10
	ExternalReferralIdentifier    []EI `hl7:"false,External Referral Identifier"`   // RF1-11
}

func (s *RF1) SegmentName() string {
	return "RF1"
}

// RGS represents the corresponding HL7 segment.
// Definition from HL7 2.4
type RGS struct {
	SetIDRGS          *SI `hl7:"true,Set ID - RGS"`         // RGS-1
	SegmentActionCode *ID `hl7:"false,Segment Action Code"` // RGS-2
	ResourceGroupID   *CE `hl7:"false,Resource Group ID"`   // RGS-3
}

func (s *RGS) SegmentName() string {
	return "RGS"
}

// RMI represents the corresponding HL7 segment.
// Definition from HL7 2.4
type RMI struct {
	RiskManagementIncidentCode *CE `hl7:"false,Risk Management Incident Code"` // RMI-1
	DateTimeIncident           *TS `hl7:"false,Date/Time Incident"`            // RMI-2
	IncidentTypeCode           *CE `hl7:"false,Incident Type Code"`            // RMI-3
}

func (s *RMI) SegmentName() string {
	return "RMI"
}

// ROL represents the corresponding HL7 segment.
// Definition from HL7 2.4
type ROL struct {
	RoleInstanceID          *EI   `hl7:"false,Role Instance ID"`             // ROL-1
	ActionCode              *ID   `hl7:"true,Action Code"`                   // ROL-2
	RoleROL                 *CE   `hl7:"true,Role-ROL"`                      // ROL-3
	RolePerson              []XCN `hl7:"true,Role Person"`                   // ROL-4
	RoleBeginDateTime       *TS   `hl7:"false,Role Begin Date/Time"`         // ROL-5
	RoleEndDateTime         *TS   `hl7:"false,Role End Date/Time"`           // ROL-6
	RoleDuration            *CE   `hl7:"false,Role Duration"`                // ROL-7
	RoleActionReason        *CE   `hl7:"false,Role Action Reason"`           // ROL-8
	ProviderType            []CE  `hl7:"false,Provider Type"`                // ROL-9
	OrganizationUnitTypeROL *CE   `hl7:"false,Organization Unit Type - ROL"` // ROL-10
	OfficeHomeAddress       []XAD `hl7:"false,Office/Home Address"`          // ROL-11
	Phone                   []XTN `hl7:"false,Phone"`                        // ROL-12
}

func (s *ROL) SegmentName() string {
	return "ROL"
}

// RQ1 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type RQ1 struct {
	AnticipatedPrice       *ST `hl7:"false,Anticipated Price"`       // RQ1-1
	ManufacturerIdentifier *CE `hl7:"false,Manufacturer Identifier"` // RQ1-2
	ManufacturerSCatalog   *ST `hl7:"false,Manufacturer'S Catalog"`  // RQ1-3
	VendorID               *CE `hl7:"false,Vendor ID"`               // RQ1-4
	VendorCatalog          *ST `hl7:"false,Vendor Catalog"`          // RQ1-5
	Taxable                *ID `hl7:"false,Taxable"`                 // RQ1-6
	SubstituteAllowed      *ID `hl7:"false,Substitute Allowed"`      // RQ1-7
}

func (s *RQ1) SegmentName() string {
	return "RQ1"
}

// RQD represents the corresponding HL7 segment.
// Definition from HL7 2.4
type RQD struct {
	RequisitionLineNumber    *SI `hl7:"false,Requisition Line Number"`     // RQD-1
	ItemCodeInternal         *CE `hl7:"false,Item Code - Internal"`        // RQD-2
	ItemCodeExternal         *CE `hl7:"false,Item Code - External"`        // RQD-3
	HospitalItemCode         *CE `hl7:"false,Hospital Item Code"`          // RQD-4
	RequisitionQuantity      *NM `hl7:"false,Requisition Quantity"`        // RQD-5
	RequisitionUnitOfMeasure *CE `hl7:"false,Requisition Unit Of Measure"` // RQD-6
	DeptCostCenter           *IS `hl7:"false,Dept. Cost Center"`           // RQD-7
	ItemNaturalAccountCode   *IS `hl7:"false,Item Natural Account Code"`   // RQD-8
	DeliverToID              *CE `hl7:"false,Deliver To ID"`               // RQD-9
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
// Definition from HL7 2.4
type RXA struct {
	GiveSubIDCounter                *NM   `hl7:"true,Give Sub-ID Counter"`                 // RXA-1
	AdministrationSubIDCounter      *NM   `hl7:"true,Administration Sub-ID Counter"`       // RXA-2
	DateTimeStartOfAdministration   *TS   `hl7:"true,Date/Time Start Of Administration"`   // RXA-3
	DateTimeEndOfAdministration     *TS   `hl7:"true,Date/Time End Of Administration"`     // RXA-4
	AdministeredCode                *CE   `hl7:"true,Administered Code"`                   // RXA-5
	AdministeredAmount              *NM   `hl7:"true,Administered Amount"`                 // RXA-6
	AdministeredUnits               *CE   `hl7:"false,Administered Units"`                 // RXA-7
	AdministeredDosageForm          *CE   `hl7:"false,Administered Dosage Form"`           // RXA-8
	AdministrationNotes             []CE  `hl7:"false,Administration Notes"`               // RXA-9
	AdministeringProvider           []XCN `hl7:"false,Administering Provider"`             // RXA-10
	AdministeredAtLocation          *LA2  `hl7:"false,Administered-At Location"`           // RXA-11
	AdministeredPerTimeUnit         *ST   `hl7:"false,Administered Per (Time Unit)"`       // RXA-12
	AdministeredStrength            *NM   `hl7:"false,Administered Strength"`              // RXA-13
	AdministeredStrengthUnits       *CE   `hl7:"false,Administered Strength Units"`        // RXA-14
	SubstanceLotNumber              []ST  `hl7:"false,Substance Lot Number"`               // RXA-15
	SubstanceExpirationDate         []TS  `hl7:"false,Substance Expiration Date"`          // RXA-16
	SubstanceManufacturerName       []CE  `hl7:"false,Substance Manufacturer Name"`        // RXA-17
	SubstanceTreatmentRefusalReason []CE  `hl7:"false,Substance/Treatment Refusal Reason"` // RXA-18
	Indication                      []CE  `hl7:"false,Indication"`                         // RXA-19
	CompletionStatus                *ID   `hl7:"false,Completion Status"`                  // RXA-20
	ActionCodeRXA                   *ID   `hl7:"false,Action Code-RXA"`                    // RXA-21
	SystemEntryDateTime             *TS   `hl7:"false,System Entry Date/Time"`             // RXA-22
}

func (s *RXA) SegmentName() string {
	return "RXA"
}

// RXC represents the corresponding HL7 segment.
// Definition from HL7 2.4
type RXC struct {
	RXComponentType        *ID  `hl7:"true,RX Component Type"`         // RXC-1
	ComponentCode          *CE  `hl7:"true,Component Code"`            // RXC-2
	ComponentAmount        *NM  `hl7:"true,Component Amount"`          // RXC-3
	ComponentUnits         *CE  `hl7:"true,Component Units"`           // RXC-4
	ComponentStrength      *NM  `hl7:"false,Component Strength"`       // RXC-5
	ComponentStrengthUnits *CE  `hl7:"false,Component Strength Units"` // RXC-6
	SupplementaryCode      []CE `hl7:"false,Supplementary Code"`       // RXC-7
}

func (s *RXC) SegmentName() string {
	return "RXC"
}

// RXD represents the corresponding HL7 segment.
// Definition from HL7 2.4
type RXD struct {
	DispenseSubIDCounter                                    *NM   `hl7:"true,Dispense Sub-ID Counter"`                                        // RXD-1
	DispenseGiveCode                                        *CE   `hl7:"true,Dispense/Give Code"`                                             // RXD-2
	DateTimeDispensed                                       *TS   `hl7:"true,Date/Time Dispensed"`                                            // RXD-3
	ActualDispenseAmount                                    *NM   `hl7:"true,Actual Dispense Amount"`                                         // RXD-4
	ActualDispenseUnits                                     *CE   `hl7:"false,Actual Dispense Units"`                                         // RXD-5
	ActualDosageForm                                        *CE   `hl7:"false,Actual Dosage Form"`                                            // RXD-6
	PrescriptionNumber                                      *ST   `hl7:"true,Prescription Number"`                                            // RXD-7
	NumberOfRefillsRemaining                                *NM   `hl7:"false,Number Of Refills Remaining"`                                   // RXD-8
	DispenseNotes                                           []ST  `hl7:"false,Dispense Notes"`                                                // RXD-9
	DispensingProvider                                      []XCN `hl7:"false,Dispensing Provider"`                                           // RXD-10
	SubstitutionStatus                                      *ID   `hl7:"false,Substitution Status"`                                           // RXD-11
	TotalDailyDose                                          *CQ   `hl7:"false,Total Daily Dose"`                                              // RXD-12
	DispenseToLocation                                      *LA2  `hl7:"false,Dispense-To Location"`                                          // RXD-13
	NeedsHumanReview                                        *ID   `hl7:"false,Needs Human Review"`                                            // RXD-14
	PharmacyTreatmentSupplierSSpecialDispensingInstructions []CE  `hl7:"false,Pharmacy/Treatment Supplier'S Special Dispensing Instructions"` // RXD-15
	ActualStrength                                          *NM   `hl7:"false,Actual Strength"`                                               // RXD-16
	ActualStrengthUnit                                      *CE   `hl7:"false,Actual Strength Unit"`                                          // RXD-17
	SubstanceLotNumber                                      []ST  `hl7:"false,Substance Lot Number"`                                          // RXD-18
	SubstanceExpirationDate                                 []TS  `hl7:"false,Substance Expiration Date"`                                     // RXD-19
	SubstanceManufacturerName                               []CE  `hl7:"false,Substance Manufacturer Name"`                                   // RXD-20
	Indication                                              []CE  `hl7:"false,Indication"`                                                    // RXD-21
	DispensePackageSize                                     *NM   `hl7:"false,Dispense Package Size"`                                         // RXD-22
	DispensePackageSizeUnit                                 *CE   `hl7:"false,Dispense Package Size Unit"`                                    // RXD-23
	DispensePackageMethod                                   *ID   `hl7:"false,Dispense Package Method"`                                       // RXD-24
	SupplementaryCode                                       []CE  `hl7:"false,Supplementary Code"`                                            // RXD-25
	InitiatingLocation                                      *CE   `hl7:"false,Initiating Location"`                                           // RXD-26
	PackagingAssemblyLocation                               *CE   `hl7:"false,Packaging/Assembly Location"`                                   // RXD-27
}

func (s *RXD) SegmentName() string {
	return "RXD"
}

// RXE represents the corresponding HL7 segment.
// Definition from HL7 2.4
type RXE struct {
	QuantityTiming                                          *TQ   `hl7:"true,Quantity/Timing"`                                                // RXE-1
	GiveCode                                                *CE   `hl7:"true,Give Code"`                                                      // RXE-2
	GiveAmountMinimum                                       *NM   `hl7:"true,Give Amount - Minimum"`                                          // RXE-3
	GiveAmountMaximum                                       *NM   `hl7:"false,Give Amount - Maximum"`                                         // RXE-4
	GiveUnits                                               *CE   `hl7:"true,Give Units"`                                                     // RXE-5
	GiveDosageForm                                          *CE   `hl7:"false,Give Dosage Form"`                                              // RXE-6
	ProviderSAdministrationInstructions                     []CE  `hl7:"false,Provider'S Administration Instructions"`                        // RXE-7
	DeliverToLocation                                       *LA1  `hl7:"false,Deliver-To Location"`                                           // RXE-8
	SubstitutionStatus                                      *ID   `hl7:"false,Substitution Status"`                                           // RXE-9
	DispenseAmount                                          *NM   `hl7:"false,Dispense Amount"`                                               // RXE-10
	DispenseUnits                                           *CE   `hl7:"false,Dispense Units"`                                                // RXE-11
	NumberOfRefills                                         *NM   `hl7:"false,Number Of Refills"`                                             // RXE-12
	OrderingProviderSDEANumber                              []XCN `hl7:"false,Ordering Provider'S DEA Number"`                                // RXE-13
	PharmacistTreatmentSupplierSVerifierID                  []XCN `hl7:"false,Pharmacist/Treatment Supplier'S Verifier ID"`                   // RXE-14
	PrescriptionNumber                                      *ST   `hl7:"false,Prescription Number"`                                           // RXE-15
	NumberOfRefillsRemaining                                *NM   `hl7:"false,Number Of Refills Remaining"`                                   // RXE-16
	NumberOfRefillsDosesDispensed                           *NM   `hl7:"false,Number Of Refills/Doses Dispensed"`                             // RXE-17
	DTOfMostRecentRefillOrDoseDispensed                     *TS   `hl7:"false,D/T Of Most Recent Refill Or Dose Dispensed"`                   // RXE-18
	TotalDailyDose                                          *CQ   `hl7:"false,Total Daily Dose"`                                              // RXE-19
	NeedsHumanReview                                        *ID   `hl7:"false,Needs Human Review"`                                            // RXE-20
	PharmacyTreatmentSupplierSSpecialDispensingInstructions []CE  `hl7:"false,Pharmacy/Treatment Supplier'S Special Dispensing Instructions"` // RXE-21
	GivePerTimeUnit                                         *ST   `hl7:"false,Give Per (Time Unit)"`                                          // RXE-22
	GiveRateAmount                                          *ST   `hl7:"false,Give Rate Amount"`                                              // RXE-23
	GiveRateUnits                                           *CE   `hl7:"false,Give Rate Units"`                                               // RXE-24
	GiveStrength                                            *NM   `hl7:"false,Give Strength"`                                                 // RXE-25
	GiveStrengthUnits                                       *CE   `hl7:"false,Give Strength Units"`                                           // RXE-26
	GiveIndication                                          []CE  `hl7:"false,Give Indication"`                                               // RXE-27
	DispensePackageSize                                     *NM   `hl7:"false,Dispense Package Size"`                                         // RXE-28
	DispensePackageSizeUnit                                 *CE   `hl7:"false,Dispense Package Size Unit"`                                    // RXE-29
	DispensePackageMethod                                   *ID   `hl7:"false,Dispense Package Method"`                                       // RXE-30
	SupplementaryCode                                       []CE  `hl7:"false,Supplementary Code"`                                            // RXE-31
}

func (s *RXE) SegmentName() string {
	return "RXE"
}

// RXG represents the corresponding HL7 segment.
// Definition from HL7 2.4
type RXG struct {
	GiveSubIDCounter                                            *NM  `hl7:"true,Give Sub-ID Counter"`                                                // RXG-1
	DispenseSubIDCounter                                        *NM  `hl7:"false,Dispense Sub-ID Counter"`                                           // RXG-2
	QuantityTiming                                              *TQ  `hl7:"true,Quantity/Timing"`                                                    // RXG-3
	GiveCode                                                    *CE  `hl7:"true,Give Code"`                                                          // RXG-4
	GiveAmountMinimum                                           *NM  `hl7:"true,Give Amount - Minimum"`                                              // RXG-5
	GiveAmountMaximum                                           *NM  `hl7:"false,Give Amount - Maximum"`                                             // RXG-6
	GiveUnits                                                   *CE  `hl7:"true,Give Units"`                                                         // RXG-7
	GiveDosageForm                                              *CE  `hl7:"false,Give Dosage Form"`                                                  // RXG-8
	AdministrationNotes                                         []CE `hl7:"false,Administration Notes"`                                              // RXG-9
	SubstitutionStatus                                          *ID  `hl7:"false,Substitution Status"`                                               // RXG-10
	DispenseToLocation                                          *LA2 `hl7:"false,Dispense-To Location"`                                              // RXG-11
	NeedsHumanReview                                            *ID  `hl7:"false,Needs Human Review"`                                                // RXG-12
	PharmacyTreatmentSupplierSSpecialAdministrationInstructions []CE `hl7:"false,Pharmacy/Treatment Supplier'S Special Administration Instructions"` // RXG-13
	GivePerTimeUnit                                             *ST  `hl7:"false,Give Per (Time Unit)"`                                              // RXG-14
	GiveRateAmount                                              *ST  `hl7:"false,Give Rate Amount"`                                                  // RXG-15
	GiveRateUnits                                               *CE  `hl7:"false,Give Rate Units"`                                                   // RXG-16
	GiveStrength                                                *NM  `hl7:"false,Give Strength"`                                                     // RXG-17
	GiveStrengthUnits                                           *CE  `hl7:"false,Give Strength Units"`                                               // RXG-18
	SubstanceLotNumber                                          []ST `hl7:"false,Substance Lot Number"`                                              // RXG-19
	SubstanceExpirationDate                                     []TS `hl7:"false,Substance Expiration Date"`                                         // RXG-20
	SubstanceManufacturerName                                   []CE `hl7:"false,Substance Manufacturer Name"`                                       // RXG-21
	Indication                                                  []CE `hl7:"false,Indication"`                                                        // RXG-22
}

func (s *RXG) SegmentName() string {
	return "RXG"
}

// RXO represents the corresponding HL7 segment.
// Definition from HL7 2.4
type RXO struct {
	RequestedGiveCode                      *CE   `hl7:"false,Requested Give Code"`                         // RXO-1
	RequestedGiveAmountMinimum             *NM   `hl7:"false,Requested Give Amount - Minimum"`             // RXO-2
	RequestedGiveAmountMaximum             *NM   `hl7:"false,Requested Give Amount - Maximum"`             // RXO-3
	RequestedGiveUnits                     *CE   `hl7:"false,Requested Give Units"`                        // RXO-4
	RequestedDosageForm                    *CE   `hl7:"false,Requested Dosage Form"`                       // RXO-5
	ProviderSPharmacyTreatmentInstructions []CE  `hl7:"false,Provider'S Pharmacy/Treatment Instructions"`  // RXO-6
	ProviderSAdministrationInstructions    []CE  `hl7:"false,Provider'S Administration Instructions"`      // RXO-7
	DeliverToLocation                      *LA1  `hl7:"false,Deliver-To Location"`                         // RXO-8
	AllowSubstitutions                     *ID   `hl7:"false,Allow Substitutions"`                         // RXO-9
	RequestedDispenseCode                  *CE   `hl7:"false,Requested Dispense Code"`                     // RXO-10
	RequestedDispenseAmount                *NM   `hl7:"false,Requested Dispense Amount"`                   // RXO-11
	RequestedDispenseUnits                 *CE   `hl7:"false,Requested Dispense Units"`                    // RXO-12
	NumberOfRefills                        *NM   `hl7:"false,Number Of Refills"`                           // RXO-13
	OrderingProviderSDEANumber             []XCN `hl7:"false,Ordering Provider'S DEA Number"`              // RXO-14
	PharmacistTreatmentSupplierSVerifierID []XCN `hl7:"false,Pharmacist/Treatment Supplier'S Verifier ID"` // RXO-15
	NeedsHumanReview                       *ID   `hl7:"false,Needs Human Review"`                          // RXO-16
	RequestedGivePerTimeUnit               *ST   `hl7:"false,Requested Give Per (Time Unit)"`              // RXO-17
	RequestedGiveStrength                  *NM   `hl7:"false,Requested Give Strength"`                     // RXO-18
	RequestedGiveStrengthUnits             *CE   `hl7:"false,Requested Give Strength Units"`               // RXO-19
	Indication                             []CE  `hl7:"false,Indication"`                                  // RXO-20
	RequestedGiveRateAmount                *ST   `hl7:"false,Requested Give Rate Amount"`                  // RXO-21
	RequestedGiveRateUnits                 *CE   `hl7:"false,Requested Give Rate Units"`                   // RXO-22
	TotalDailyDose                         *CQ   `hl7:"false,Total Daily Dose"`                            // RXO-23
	SupplementaryCode                      []CE  `hl7:"false,Supplementary Code"`                          // RXO-24
}

func (s *RXO) SegmentName() string {
	return "RXO"
}

// RXR represents the corresponding HL7 segment.
// Definition from HL7 2.4
type RXR struct {
	Route                *CE `hl7:"true,Route"`                  // RXR-1
	AdministrationSite   *CE `hl7:"false,Administration Site"`   // RXR-2
	AdministrationDevice *CE `hl7:"false,Administration Device"` // RXR-3
	AdministrationMethod *CE `hl7:"false,Administration Method"` // RXR-4
	RoutingInstruction   *CE `hl7:"false,Routing Instruction"`   // RXR-5
}

func (s *RXR) SegmentName() string {
	return "RXR"
}

// SAC represents the corresponding HL7 segment.
// Definition from HL7 2.4
type SAC struct {
	ExternalAccessionIdentifier       *EI  `hl7:"false,External Accession Identifier"`         // SAC-1
	AccessionIdentifier               *EI  `hl7:"false,Accession Identifier"`                  // SAC-2
	ContainerIdentifier               *EI  `hl7:"false,Container Identifier"`                  // SAC-3
	PrimaryParentContainerIdentifier  *EI  `hl7:"false,Primary (Parent) Container Identifier"` // SAC-4
	EquipmentContainerIdentifier      *EI  `hl7:"false,Equipment Container Identifier"`        // SAC-5
	SpecimenSource                    *SPS `hl7:"false,Specimen Source"`                       // SAC-6
	RegistrationDateTime              *TS  `hl7:"false,Registration Date/Time"`                // SAC-7
	ContainerStatus                   *CE  `hl7:"false,Container Status"`                      // SAC-8
	CarrierType                       *CE  `hl7:"false,Carrier Type"`                          // SAC-9
	CarrierIdentifier                 *EI  `hl7:"false,Carrier Identifier"`                    // SAC-10
	PositionInCarrier                 *NA  `hl7:"false,Position In Carrier"`                   // SAC-11
	TrayTypeSAC                       *CE  `hl7:"false,Tray Type - SAC"`                       // SAC-12
	TrayIdentifier                    *EI  `hl7:"false,Tray Identifier"`                       // SAC-13
	PositionInTray                    *NA  `hl7:"false,Position In Tray"`                      // SAC-14
	Location                          []CE `hl7:"false,Location"`                              // SAC-15
	ContainerHeight                   *NM  `hl7:"false,Container Height"`                      // SAC-16
	ContainerDiameter                 *NM  `hl7:"false,Container Diameter"`                    // SAC-17
	BarrierDelta                      *NM  `hl7:"false,Barrier Delta"`                         // SAC-18
	BottomDelta                       *NM  `hl7:"false,Bottom Delta"`                          // SAC-19
	ContainerHeightDiameterDeltaUnits *CE  `hl7:"false,Container Height/Diameter/Delta Units"` // SAC-20
	ContainerVolume                   *NM  `hl7:"false,Container Volume"`                      // SAC-21
	AvailableVolume                   *NM  `hl7:"false,Available Volume"`                      // SAC-22
	InitialSpecimenVolume             *NM  `hl7:"false,Initial Specimen Volume"`               // SAC-23
	VolumeUnits                       *CE  `hl7:"false,Volume  Units"`                         // SAC-24
	SeparatorType                     *CE  `hl7:"false,Separator Type"`                        // SAC-25
	CapType                           *CE  `hl7:"false,Cap Type"`                              // SAC-26
	Additive                          []CE `hl7:"false,Additive"`                              // SAC-27
	SpecimenComponent                 *CE  `hl7:"false,Specimen Component"`                    // SAC-28
	DilutionFactor                    *SN  `hl7:"false,Dilution Factor"`                       // SAC-29
	Treatment                         *CE  `hl7:"false,Treatment"`                             // SAC-30
	Temperature                       *SN  `hl7:"false,Temperature"`                           // SAC-31
	HemolysisIndex                    *NM  `hl7:"false,Hemolysis Index"`                       // SAC-32
	HemolysisIndexUnits               *CE  `hl7:"false,Hemolysis Index Units"`                 // SAC-33
	LipemiaIndex                      *NM  `hl7:"false,Lipemia Index"`                         // SAC-34
	LipemiaIndexUnits                 *CE  `hl7:"false,Lipemia Index Units"`                   // SAC-35
	IcterusIndex                      *NM  `hl7:"false,Icterus Index"`                         // SAC-36
	IcterusIndexUnits                 *CE  `hl7:"false,Icterus Index Units"`                   // SAC-37
	FibrinIndex                       *NM  `hl7:"false,Fibrin Index"`                          // SAC-38
	FibrinIndexUnits                  *CE  `hl7:"false,Fibrin Index Units"`                    // SAC-39
	SystemInducedContaminants         []CE `hl7:"false,System Induced Contaminants"`           // SAC-40
	DrugInterference                  []CE `hl7:"false,Drug Interference"`                     // SAC-41
	ArtificialBlood                   *CE  `hl7:"false,Artificial Blood"`                      // SAC-42
	SpecialHandlingConsiderations     []CE `hl7:"false,Special Handling Considerations"`       // SAC-43
	OtherEnvironmentalFactors         []CE `hl7:"false,Other Environmental Factors"`           // SAC-44
}

func (s *SAC) SegmentName() string {
	return "SAC"
}

// SCH represents the corresponding HL7 segment.
// Definition from HL7 2.4
type SCH struct {
	PlacerAppointmentID       *EI   `hl7:"false,Placer Appointment ID"`        // SCH-1
	FillerAppointmentID       *EI   `hl7:"false,Filler Appointment ID"`        // SCH-2
	OccurrenceNumber          *NM   `hl7:"false,Occurrence Number"`            // SCH-3
	PlacerGroupNumber         *EI   `hl7:"false,Placer Group Number"`          // SCH-4
	ScheduleID                *CE   `hl7:"false,Schedule ID"`                  // SCH-5
	EventReason               *CE   `hl7:"true,Event Reason"`                  // SCH-6
	AppointmentReason         *CE   `hl7:"false,Appointment Reason"`           // SCH-7
	AppointmentType           *CE   `hl7:"false,Appointment Type"`             // SCH-8
	AppointmentDuration       *NM   `hl7:"false,Appointment Duration"`         // SCH-9
	AppointmentDurationUnits  *CE   `hl7:"false,Appointment Duration Units"`   // SCH-10
	AppointmentTimingQuantity []TQ  `hl7:"true,Appointment Timing Quantity"`   // SCH-11
	PlacerContactPerson       []XCN `hl7:"false,Placer Contact Person"`        // SCH-12
	PlacerContactPhoneNumber  *XTN  `hl7:"false,Placer Contact Phone Number"`  // SCH-13
	PlacerContactAddress      []XAD `hl7:"false,Placer Contact Address"`       // SCH-14
	PlacerContactLocation     *PL   `hl7:"false,Placer Contact Location"`      // SCH-15
	FillerContactPerson       []XCN `hl7:"true,Filler Contact Person"`         // SCH-16
	FillerContactPhoneNumber  *XTN  `hl7:"false,Filler Contact Phone Number"`  // SCH-17
	FillerContactAddress      []XAD `hl7:"false,Filler Contact Address"`       // SCH-18
	FillerContactLocation     *PL   `hl7:"false,Filler Contact Location"`      // SCH-19
	EnteredByPerson           []XCN `hl7:"true,Entered By Person"`             // SCH-20
	EnteredByPhoneNumber      []XTN `hl7:"false,Entered By Phone Number"`      // SCH-21
	EnteredByLocation         *PL   `hl7:"false,Entered By Location"`          // SCH-22
	ParentPlacerAppointmentID *EI   `hl7:"false,Parent Placer Appointment ID"` // SCH-23
	ParentFillerAppointmentID *EI   `hl7:"false,Parent Filler Appointment ID"` // SCH-24
	FillerStatusCode          *CE   `hl7:"false,Filler Status Code"`           // SCH-25
	PlacerOrderNumber         []EI  `hl7:"false,Placer Order Number"`          // SCH-26
	FillerOrderNumber         []EI  `hl7:"false,Filler Order Number"`          // SCH-27
}

func (s *SCH) SegmentName() string {
	return "SCH"
}

// SID represents the corresponding HL7 segment.
// Definition from HL7 2.4
type SID struct {
	ApplicationMethodIdentifier     *CE `hl7:"false,Application / Method Identifier"`   // SID-1
	SubstanceLotNumber              *ST `hl7:"false,Substance Lot Number"`              // SID-2
	SubstanceContainerIdentifier    *ST `hl7:"false,Substance Container Identifier"`    // SID-3
	SubstanceManufacturerIdentifier *CE `hl7:"false,Substance Manufacturer Identifier"` // SID-4
}

func (s *SID) SegmentName() string {
	return "SID"
}

// SPR represents the corresponding HL7 segment.
// Definition from HL7 2.4
type SPR struct {
	QueryTag                *ST   `hl7:"false,Query Tag"`                 // SPR-1
	QueryResponseFormatCode *ID   `hl7:"true,Query/Response Format Code"` // SPR-2
	StoredProcedureName     *CE   `hl7:"true,Stored Procedure Name"`      // SPR-3
	InputParameterList      []QIP `hl7:"false,Input Parameter List"`      // SPR-4
}

func (s *SPR) SegmentName() string {
	return "SPR"
}

// STF represents the corresponding HL7 segment.
// Definition from HL7 2.4
type STF struct {
	PrimaryKeyValueSTF            *CE   `hl7:"false,Primary Key Value - STF"`          // STF-1
	StaffIDCode                   []CX  `hl7:"false,Staff ID Code"`                    // STF-2
	StaffName                     []XPN `hl7:"false,Staff Name"`                       // STF-3
	StaffType                     []IS  `hl7:"false,Staff Type"`                       // STF-4
	AdministrativeSex             *IS   `hl7:"false,Administrative Sex"`               // STF-5
	DateTimeOfBirth               *TS   `hl7:"false,Date/Time Of Birth"`               // STF-6
	ActiveInactiveFlag            *ID   `hl7:"false,Active/Inactive Flag"`             // STF-7
	Department                    []CE  `hl7:"false,Department"`                       // STF-8
	HospitalService               []CE  `hl7:"false,Hospital Service"`                 // STF-9
	Phone                         []XTN `hl7:"false,Phone"`                            // STF-10
	OfficeHomeAddress             []XAD `hl7:"false,Office/Home Address"`              // STF-11
	InstitutionActivationDate     []DIN `hl7:"false,Institution Activation Date"`      // STF-12
	InstitutionInactivationDate   []DIN `hl7:"false,Institution Inactivation Date"`    // STF-13
	BackupPersonID                []CE  `hl7:"false,Backup Person ID"`                 // STF-14
	EMailAddress                  []ST  `hl7:"false,E-Mail Address"`                   // STF-15
	PreferredMethodOfContact      *CE   `hl7:"false,Preferred Method Of Contact"`      // STF-16
	MaritalStatus                 *CE   `hl7:"false,Marital Status"`                   // STF-17
	JobTitle                      *ST   `hl7:"false,Job Title"`                        // STF-18
	JobCodeClass                  *JCC  `hl7:"false,Job Code/Class"`                   // STF-19
	EmploymentStatusCode          *CE   `hl7:"false,Employment Status Code"`           // STF-20
	AdditionalInsuredOnAuto       *ID   `hl7:"false,Additional Insured On  Auto"`      // STF-21
	DriverSLicenseNumberStaff     *DLN  `hl7:"false,Driver'S License Number - Staff"`  // STF-22
	CopyAutoIns                   *ID   `hl7:"false,Copy  Auto Ins"`                   // STF-23
	AutoInsExpires                *DT   `hl7:"false,Auto Ins. Expires"`                // STF-24
	DateLastDMVReview             *DT   `hl7:"false,Date Last DMV Review"`             // STF-25
	DateNextDMVReview             *DT   `hl7:"false,Date Next DMV Review"`             // STF-26
	Race                          *CE   `hl7:"false,Race"`                             // STF-27
	EthnicGroup                   *CE   `hl7:"false,Ethnic Group"`                     // STF-28
	ReActivationApprovalIndicator *ID   `hl7:"false,Re-Activation Approval Indicator"` // STF-29
}

func (s *STF) SegmentName() string {
	return "STF"
}

// TCC represents the corresponding HL7 segment.
// Definition from HL7 2.4
type TCC struct {
	UniversalServiceIdentifier            *CE  `hl7:"true,Universal Service Identifier"`                // TCC-1
	TestApplicationIdentifier             *EI  `hl7:"true,Test Application Identifier"`                 // TCC-2
	SpecimenSource                        *SPS `hl7:"false,Specimen Source"`                            // TCC-3
	AutoDilutionFactorDefault             *SN  `hl7:"false,Auto-Dilution Factor Default"`               // TCC-4
	RerunDilutionFactorDefault            *SN  `hl7:"false,Rerun Dilution Factor Default"`              // TCC-5
	PreDilutionFactorDefault              *SN  `hl7:"false,Pre-Dilution Factor Default"`                // TCC-6
	EndogenousContentOfPreDilutionDiluent *SN  `hl7:"false,Endogenous Content Of Pre-Dilution Diluent"` // TCC-7
	InventoryLimitsWarningLevel           *NM  `hl7:"false,Inventory Limits Warning Level"`             // TCC-8
	AutomaticRerunAllowed                 *ID  `hl7:"false,Automatic Rerun Allowed"`                    // TCC-9
	AutomaticRepeatAllowed                *ID  `hl7:"false,Automatic Repeat Allowed"`                   // TCC-10
	AutomaticReflexAllowed                *ID  `hl7:"false,Automatic Reflex Allowed"`                   // TCC-11
	EquipmentDynamicRange                 *SN  `hl7:"false,Equipment Dynamic Range"`                    // TCC-12
	Units                                 *CE  `hl7:"false,Units"`                                      // TCC-13
	ProcessingType                        *CE  `hl7:"false,Processing Type"`                            // TCC-14
}

func (s *TCC) SegmentName() string {
	return "TCC"
}

// TCD represents the corresponding HL7 segment.
// Definition from HL7 2.4
type TCD struct {
	UniversalServiceIdentifier            *CE `hl7:"true,Universal Service Identifier"`                // TCD-1
	AutoDilutionFactor                    *SN `hl7:"false,Auto-Dilution Factor"`                       // TCD-2
	RerunDilutionFactor                   *SN `hl7:"false,Rerun Dilution Factor"`                      // TCD-3
	PreDilutionFactor                     *SN `hl7:"false,Pre-Dilution Factor"`                        // TCD-4
	EndogenousContentOfPreDilutionDiluent *SN `hl7:"false,Endogenous Content Of Pre-Dilution Diluent"` // TCD-5
	AutomaticRepeatAllowed                *ID `hl7:"false,Automatic Repeat Allowed"`                   // TCD-6
	ReflexAllowed                         *ID `hl7:"false,Reflex Allowed"`                             // TCD-7
	AnalyteRepeatStatus                   *CE `hl7:"false,Analyte Repeat Status"`                      // TCD-8
}

func (s *TCD) SegmentName() string {
	return "TCD"
}

// TXA represents the corresponding HL7 segment.
// Definition from HL7 2.4
type TXA struct {
	SetIDTXA                                 *SI   `hl7:"true,Set ID- TXA"`                                       // TXA-1
	DocumentType                             *IS   `hl7:"true,Document Type"`                                     // TXA-2
	DocumentContentPresentation              *ID   `hl7:"false,Document Content Presentation"`                    // TXA-3
	ActivityDateTime                         *TS   `hl7:"false,Activity Date/Time"`                               // TXA-4
	PrimaryActivityProviderCodeName          []XCN `hl7:"false,Primary Activity Provider Code/Name"`              // TXA-5
	OriginationDateTime                      *TS   `hl7:"false,Origination Date/Time"`                            // TXA-6
	TranscriptionDateTime                    *TS   `hl7:"false,Transcription Date/Time"`                          // TXA-7
	EditDateTime                             []TS  `hl7:"false,Edit Date/Time"`                                   // TXA-8
	OriginatorCodeName                       []XCN `hl7:"false,Originator Code/Name"`                             // TXA-9
	AssignedDocumentAuthenticator            []XCN `hl7:"false,Assigned Document Authenticator"`                  // TXA-10
	TranscriptionistCodeName                 []XCN `hl7:"false,Transcriptionist Code/Name"`                       // TXA-11
	UniqueDocumentNumber                     *EI   `hl7:"true,Unique Document Number"`                            // TXA-12
	ParentDocumentNumber                     *EI   `hl7:"false,Parent Document Number"`                           // TXA-13
	PlacerOrderNumber                        []EI  `hl7:"false,Placer Order Number"`                              // TXA-14
	FillerOrderNumber                        *EI   `hl7:"false,Filler Order Number"`                              // TXA-15
	UniqueDocumentFileName                   *ST   `hl7:"false,Unique Document File Name"`                        // TXA-16
	DocumentCompletionStatus                 *ID   `hl7:"true,Document Completion Status"`                        // TXA-17
	DocumentConfidentialityStatus            *ID   `hl7:"false,Document Confidentiality Status"`                  // TXA-18
	DocumentAvailabilityStatus               *ID   `hl7:"false,Document Availability Status"`                     // TXA-19
	DocumentStorageStatus                    *ID   `hl7:"false,Document Storage Status"`                          // TXA-20
	DocumentChangeReason                     *ST   `hl7:"false,Document Change Reason"`                           // TXA-21
	AuthenticationPersonTimeStamp            []PPN `hl7:"false,Authentication Person, Time Stamp"`                // TXA-22
	DistributedCopiesCodeAndNameOfRecipients []XCN `hl7:"false,Distributed Copies (Code And Name Of Recipients)"` // TXA-23
}

func (s *TXA) SegmentName() string {
	return "TXA"
}

// UB1 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type UB1 struct {
	SetIDUB1                  *SI   `hl7:"false,Set ID - UB1"`                    // UB1-1
	BloodDeductible43         *NM   `hl7:"false,Blood Deductible  (43)"`          // UB1-2
	BloodFurnishedPintsOf40   *NM   `hl7:"false,Blood Furnished-Pints Of (40)"`   // UB1-3
	BloodReplacedPints41      *NM   `hl7:"false,Blood Replaced-Pints (41)"`       // UB1-4
	BloodNotReplacedPints42   *NM   `hl7:"false,Blood Not Replaced-Pints(42)"`    // UB1-5
	CoInsuranceDays25         *NM   `hl7:"false,Co-Insurance Days (25)"`          // UB1-6
	ConditionCode3539         []IS  `hl7:"false,Condition Code (35-39)"`          // UB1-7
	CoveredDays23             *NM   `hl7:"false,Covered Days - (23)"`             // UB1-8
	NonCoveredDays24          *NM   `hl7:"false,Non Covered Days - (24)"`         // UB1-9
	ValueAmountCode4649       []UVC `hl7:"false,Value Amount & Code (46-49)"`     // UB1-10
	NumberOfGraceDays90       *NM   `hl7:"false,Number Of Grace Days (90)"`       // UB1-11
	SpecialProgramIndicator44 *CE   `hl7:"false,Special Program Indicator (44)"`  // UB1-12
	PSROURApprovalIndicator87 *CE   `hl7:"false,PSRO/UR Approval Indicator (87)"` // UB1-13
	PSROURApprovedStayFm88    *DT   `hl7:"false,PSRO/UR Approved Stay-Fm (88)"`   // UB1-14
	PSROURApprovedStayTo89    *DT   `hl7:"false,PSRO/UR Approved Stay-To (89)"`   // UB1-15
	Occurrence2832            []OCD `hl7:"false,Occurrence (28-32)"`              // UB1-16
	OccurrenceSpan33          *CE   `hl7:"false,Occurrence Span (33)"`            // UB1-17
	OccurSpanStartDate33      *DT   `hl7:"false,Occur Span Start Date(33)"`       // UB1-18
	OccurSpanEndDate33        *DT   `hl7:"false,Occur Span End Date (33)"`        // UB1-19
	UB82Locator2              *ST   `hl7:"false,UB-82 Locator 2"`                 // UB1-20
	UB82Locator9              *ST   `hl7:"false,UB-82 Locator 9"`                 // UB1-21
	UB82Locator27             *ST   `hl7:"false,UB-82 Locator 27"`                // UB1-22
	UB82Locator45             *ST   `hl7:"false,UB-82 Locator 45"`                // UB1-23
}

func (s *UB1) SegmentName() string {
	return "UB1"
}

// UB2 represents the corresponding HL7 segment.
// Definition from HL7 2.4
type UB2 struct {
	SetIDUB2                  *SI   `hl7:"false,Set ID - UB2"`                    // UB2-1
	CoInsuranceDays9          *ST   `hl7:"false,Co-Insurance Days (9)"`           // UB2-2
	ConditionCode2430         []IS  `hl7:"false,Condition Code (24-30)"`          // UB2-3
	CoveredDays7              *ST   `hl7:"false,Covered Days (7)"`                // UB2-4
	NonCoveredDays8           *ST   `hl7:"false,Non-Covered Days (8)"`            // UB2-5
	ValueAmountCode           []UVC `hl7:"false,Value Amount & Code"`             // UB2-6
	OccurrenceCodeDate3235    []OCD `hl7:"false,Occurrence Code & Date (32-35)"`  // UB2-7
	OccurrenceSpanCodeDates36 []OSP `hl7:"false,Occurrence Span Code/Dates (36)"` // UB2-8
	UB92Locator2State         []ST  `hl7:"false,UB92 Locator 2 (State)"`          // UB2-9
	UB92Locator11State        []ST  `hl7:"false,UB92 Locator 11 (State)"`         // UB2-10
	UB92Locator31National     *ST   `hl7:"false,UB92 Locator 31 (National)"`      // UB2-11
	DocumentControlNumber     []ST  `hl7:"false,Document Control Number"`         // UB2-12
	UB92Locator49National     []ST  `hl7:"false,UB92 Locator 49 (National)"`      // UB2-13
	UB92Locator56State        []ST  `hl7:"false,UB92 Locator 56 (State)"`         // UB2-14
	UB92Locator57National     *ST   `hl7:"false,UB92 Locator 57 (National)"`      // UB2-15
	UB92Locator78State        []ST  `hl7:"false,UB92 Locator 78 (State)"`         // UB2-16
	SpecialVisitCount         *NM   `hl7:"false,Special Visit Count"`             // UB2-17
}

func (s *UB2) SegmentName() string {
	return "UB2"
}

// URD represents the corresponding HL7 segment.
// Definition from HL7 2.4
type URD struct {
	RUDateTime              *TS   `hl7:"false,R/U Date/Time"`               // URD-1
	ReportPriority          *ID   `hl7:"false,Report Priority"`             // URD-2
	RUWhoSubjectDefinition  []XCN `hl7:"true,R/U Who Subject Definition"`   // URD-3
	RUWhatSubjectDefinition []CE  `hl7:"false,R/U What Subject Definition"` // URD-4
	RUWhatDepartmentCode    []CE  `hl7:"false,R/U What Department Code"`    // URD-5
	RUDisplayPrintLocations []ST  `hl7:"false,R/U Display/Print Locations"` // URD-6
	RUResultsLevel          *ID   `hl7:"false,R/U Results Level"`           // URD-7
}

func (s *URD) SegmentName() string {
	return "URD"
}

// URS represents the corresponding HL7 segment.
// Definition from HL7 2.4
type URS struct {
	RUWhereSubjectDefinition        []ST `hl7:"true,R/U Where Subject Definition"`          // URS-1
	RUWhenDataStartDateTime         *TS  `hl7:"false,R/U When Data Start Date/Time"`        // URS-2
	RUWhenDataEndDateTime           *TS  `hl7:"false,R/U When Data End Date/Time"`          // URS-3
	RUWhatUserQualifier             []ST `hl7:"false,R/U What User Qualifier"`              // URS-4
	RUOtherResultsSubjectDefinition []ST `hl7:"false,R/U Other Results Subject Definition"` // URS-5
	RUWhichDateTimeQualifier        []ID `hl7:"false,R/U Which Date/Time Qualifier"`        // URS-6
	RUWhichDateTimeStatusQualifier  []ID `hl7:"false,R/U Which Date/Time Status Qualifier"` // URS-7
	RUDateTimeSelectionQualifier    []ID `hl7:"false,R/U Date/Time Selection Qualifier"`    // URS-8
	RUQuantityTimingQualifier       *TQ  `hl7:"false,R/U Quantity/Timing Qualifier"`        // URS-9
}

func (s *URS) SegmentName() string {
	return "URS"
}

// VAR represents the corresponding HL7 segment.
// Definition from HL7 2.4
type VAR struct {
	VarianceInstanceID     *EI   `hl7:"true,Variance Instance ID"`       // VAR-1
	DocumentedDateTime     *TS   `hl7:"true,Documented Date/Time"`       // VAR-2
	StatedVarianceDateTime *TS   `hl7:"false,Stated Variance Date/Time"` // VAR-3
	VarianceOriginator     []XCN `hl7:"false,Variance Originator"`       // VAR-4
	VarianceClassification *CE   `hl7:"false,Variance Classification"`   // VAR-5
	VarianceDescription    []ST  `hl7:"false,Variance Description"`      // VAR-6
}

func (s *VAR) SegmentName() string {
	return "VAR"
}

// VTQ represents the corresponding HL7 segment.
// Definition from HL7 2.4
type VTQ struct {
	QueryTag                *ST   `hl7:"false,Query Tag"`                 // VTQ-1
	QueryResponseFormatCode *ID   `hl7:"true,Query/Response Format Code"` // VTQ-2
	VTQueryName             *CE   `hl7:"true,VT Query Name"`              // VTQ-3
	VirtualTableName        *CE   `hl7:"true,Virtual Table Name"`         // VTQ-4
	SelectionCriteria       []QSC `hl7:"false,Selection Criteria"`        // VTQ-5
}

func (s *VTQ) SegmentName() string {
	return "VTQ"
}

// ACK represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ACK struct {
	MSH   *MSH `hl7:"true,MSH"`
	MSA   *MSA `hl7:"true,MSA"`
	ERR   *ERR `hl7:"false,ERR"`
	Other []interface{}
}

func (s *ACK) MessageTypeName() string {
	return "ACK"
}

// ACK_N02 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ACK_N02 struct {
	MSH   *MSH `hl7:"true,MSH"`
	MSA   *MSA `hl7:"true,MSA"`
	Other []interface{}
}

func (s *ACK_N02) MessageTypeName() string {
	return "ACK_N02"
}

// ADR_A19 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADR_A19 struct {
	MSH            *MSH                     `hl7:"true,MSH"`
	MSA            *MSA                     `hl7:"true,MSA"`
	ERR            *ERR                     `hl7:"false,ERR"`
	QAK            *QAK                     `hl7:"false,QAK"`
	QRD            *QRD                     `hl7:"true,QRD"`
	QRF            *QRF                     `hl7:"false,QRF"`
	QUERY_RESPONSE []ADR_A19_QUERY_RESPONSE `hl7:"true,QUERY_RESPONSE"`
	DSC            *DSC                     `hl7:"false,DSC"`
	Other          []interface{}
}

func (s *ADR_A19) MessageTypeName() string {
	return "ADR_A19"
}

// ADR_A19_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADR_A19_INSURANCE struct {
	IN1   *IN1  `hl7:"true,IN1"`
	IN2   *IN2  `hl7:"false,IN2"`
	IN3   []IN3 `hl7:"false,IN3"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *ADR_A19_INSURANCE) MessageTypeSubStructName() string {
	return "ADR_A19_INSURANCE"
}

// ADR_A19_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADR_A19_PROCEDURE struct {
	PR1   *PR1  `hl7:"true,PR1"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *ADR_A19_PROCEDURE) MessageTypeSubStructName() string {
	return "ADR_A19_PROCEDURE"
}

// ADR_A19_QUERY_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADR_A19_QUERY_RESPONSE struct {
	EVN       *EVN                `hl7:"false,EVN"`
	PID       *PID                `hl7:"true,PID"`
	PD1       *PD1                `hl7:"false,PD1"`
	ROL1      []ROL               `hl7:"false,ROL1"`
	NK1       []NK1               `hl7:"false,NK1"`
	PV1       *PV1                `hl7:"true,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	ROL2      []ROL               `hl7:"false,ROL2"`
	DB1       []DB1               `hl7:"false,DB1"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	DRG       *DRG                `hl7:"false,DRG"`
	PROCEDURE []ADR_A19_PROCEDURE `hl7:"false,PROCEDURE"`
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
// Definition from HL7 2.4
type ADT_A01 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	EVN       *EVN                `hl7:"true,EVN"`
	PID       *PID                `hl7:"true,PID"`
	PD1       *PD1                `hl7:"false,PD1"`
	ROL1      []ROL               `hl7:"false,ROL1"`
	NK1       []NK1               `hl7:"false,NK1"`
	PV1       *PV1                `hl7:"true,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	ROL2      []ROL               `hl7:"false,ROL2"`
	DB1       []DB1               `hl7:"false,DB1"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	DRG       *DRG                `hl7:"false,DRG"`
	PROCEDURE []ADT_A01_PROCEDURE `hl7:"false,PROCEDURE"`
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []ADT_A01_INSURANCE `hl7:"false,INSURANCE"`
	ACC       *ACC                `hl7:"false,ACC"`
	UB1       *UB1                `hl7:"false,UB1"`
	UB2       *UB2                `hl7:"false,UB2"`
	PDA       *PDA                `hl7:"false,PDA"`
	Other     []interface{}
}

func (s *ADT_A01) MessageTypeName() string {
	return "ADT_A01"
}

// ADT_A01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A01_INSURANCE struct {
	IN1   *IN1  `hl7:"true,IN1"`
	IN2   *IN2  `hl7:"false,IN2"`
	IN3   []IN3 `hl7:"false,IN3"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *ADT_A01_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A01_INSURANCE"
}

// ADT_A01_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A01_PROCEDURE struct {
	PR1   *PR1  `hl7:"true,PR1"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *ADT_A01_PROCEDURE) MessageTypeSubStructName() string {
	return "ADT_A01_PROCEDURE"
}

// ADT_A02 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A02 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PD1   *PD1  `hl7:"false,PD1"`
	ROL1  []ROL `hl7:"false,ROL1"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	ROL2  []ROL `hl7:"false,ROL2"`
	DB1   []DB1 `hl7:"false,DB1"`
	OBX   []OBX `hl7:"false,OBX"`
	PDA   *PDA  `hl7:"false,PDA"`
	Other []interface{}
}

func (s *ADT_A02) MessageTypeName() string {
	return "ADT_A02"
}

// ADT_A03 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A03 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	EVN       *EVN                `hl7:"true,EVN"`
	PID       *PID                `hl7:"true,PID"`
	PD1       *PD1                `hl7:"false,PD1"`
	ROL1      []ROL               `hl7:"false,ROL1"`
	PV1       *PV1                `hl7:"true,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	ROL2      []ROL               `hl7:"false,ROL2"`
	DB1       []DB1               `hl7:"false,DB1"`
	DG1       []DG1               `hl7:"false,DG1"`
	DRG       *DRG                `hl7:"false,DRG"`
	PROCEDURE []ADT_A03_PROCEDURE `hl7:"false,PROCEDURE"`
	OBX       []OBX               `hl7:"false,OBX"`
	PDA       *PDA                `hl7:"false,PDA"`
	Other     []interface{}
}

func (s *ADT_A03) MessageTypeName() string {
	return "ADT_A03"
}

// ADT_A03_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A03_PROCEDURE struct {
	PR1   *PR1  `hl7:"true,PR1"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *ADT_A03_PROCEDURE) MessageTypeSubStructName() string {
	return "ADT_A03_PROCEDURE"
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
// Definition from HL7 2.4
type ADT_A05 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	EVN       *EVN                `hl7:"true,EVN"`
	PID       *PID                `hl7:"true,PID"`
	PD1       *PD1                `hl7:"false,PD1"`
	ROL1      []ROL               `hl7:"false,ROL1"`
	NK1       []NK1               `hl7:"false,NK1"`
	PV1       *PV1                `hl7:"true,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	ROL2      []ROL               `hl7:"false,ROL2"`
	DB1       []DB1               `hl7:"false,DB1"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	DRG       *DRG                `hl7:"false,DRG"`
	PROCEDURE []ADT_A05_PROCEDURE `hl7:"false,PROCEDURE"`
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
// Definition from HL7 2.4
type ADT_A05_INSURANCE struct {
	IN1   *IN1  `hl7:"true,IN1"`
	IN2   *IN2  `hl7:"false,IN2"`
	IN3   []IN3 `hl7:"false,IN3"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *ADT_A05_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A05_INSURANCE"
}

// ADT_A05_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A05_PROCEDURE struct {
	PR1   *PR1  `hl7:"true,PR1"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *ADT_A05_PROCEDURE) MessageTypeSubStructName() string {
	return "ADT_A05_PROCEDURE"
}

// ADT_A06 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A06 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	EVN       *EVN                `hl7:"true,EVN"`
	PID       *PID                `hl7:"true,PID"`
	PD1       *PD1                `hl7:"false,PD1"`
	ROL1      []ROL               `hl7:"false,ROL1"`
	MRG       *MRG                `hl7:"false,MRG"`
	NK1       []NK1               `hl7:"false,NK1"`
	PV1       *PV1                `hl7:"true,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	ROL2      []ROL               `hl7:"false,ROL2"`
	DB1       []DB1               `hl7:"false,DB1"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	DRG       *DRG                `hl7:"false,DRG"`
	PROCEDURE []ADT_A06_PROCEDURE `hl7:"false,PROCEDURE"`
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
// Definition from HL7 2.4
type ADT_A06_INSURANCE struct {
	IN1   *IN1  `hl7:"true,IN1"`
	IN2   *IN2  `hl7:"false,IN2"`
	IN3   []IN3 `hl7:"false,IN3"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *ADT_A06_INSURANCE) MessageTypeSubStructName() string {
	return "ADT_A06_INSURANCE"
}

// ADT_A06_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A06_PROCEDURE struct {
	PR1   *PR1  `hl7:"true,PR1"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *ADT_A06_PROCEDURE) MessageTypeSubStructName() string {
	return "ADT_A06_PROCEDURE"
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
// Definition from HL7 2.4
type ADT_A09 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PD1   *PD1  `hl7:"false,PD1"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	DB1   []DB1 `hl7:"false,DB1"`
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
// Definition from HL7 2.3.1
type ADT_A12 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PD1   *PD1  `hl7:"false,PD1"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	DB1   []DB1 `hl7:"false,DB1"`
	OBX   []OBX `hl7:"false,OBX"`
	DG1   *DG1  `hl7:"false,DG1"`
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
// Definition from HL7 2.4
type ADT_A15 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PD1   *PD1  `hl7:"false,PD1"`
	ROL1  []ROL `hl7:"false,ROL1"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	ROL2  []ROL `hl7:"false,ROL2"`
	DB1   []DB1 `hl7:"false,DB1"`
	OBX   []OBX `hl7:"false,OBX"`
	DG1   []DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A15) MessageTypeName() string {
	return "ADT_A15"
}

// ADT_A16 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A16 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PD1   *PD1  `hl7:"false,PD1"`
	ROL1  []ROL `hl7:"false,ROL1"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	ROL2  []ROL `hl7:"false,ROL2"`
	DB1   []DB1 `hl7:"false,DB1"`
	OBX   []OBX `hl7:"false,OBX"`
	DG1   []DG1 `hl7:"false,DG1"`
	DRG   *DRG  `hl7:"false,DRG"`
	Other []interface{}
}

func (s *ADT_A16) MessageTypeName() string {
	return "ADT_A16"
}

// ADT_A17 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A17 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID1  *PID  `hl7:"true,PID1"`
	PD11  *PD1  `hl7:"false,PD11"`
	PV11  *PV1  `hl7:"true,PV11"`
	PV21  *PV2  `hl7:"false,PV21"`
	DB11  []DB1 `hl7:"false,DB11"`
	OBX1  []OBX `hl7:"false,OBX1"`
	PID2  *PID  `hl7:"true,PID2"`
	PD12  *PD1  `hl7:"false,PD12"`
	PV12  *PV1  `hl7:"true,PV12"`
	PV22  *PV2  `hl7:"false,PV22"`
	DB12  []DB1 `hl7:"false,DB12"`
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
// Definition from HL7 2.4
type ADT_A18 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PD1   *PD1 `hl7:"false,PD1"`
	MRG   *MRG `hl7:"true,MRG"`
	PV1   *PV1 `hl7:"true,PV1"`
	Other []interface{}
}

func (s *ADT_A18) MessageTypeName() string {
	return "ADT_A18"
}

// ADT_A20 represents the corresponding HL7 message type.
// Definition from HL7 2.4
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
// Definition from HL7 2.4
type ADT_A21 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PD1   *PD1  `hl7:"false,PD1"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	DB1   []DB1 `hl7:"false,DB1"`
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
// Definition from HL7 2.4
type ADT_A24 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID1  *PID  `hl7:"true,PID1"`
	PD11  *PD1  `hl7:"false,PD11"`
	PV11  *PV1  `hl7:"false,PV11"`
	DB11  []DB1 `hl7:"false,DB11"`
	PID2  *PID  `hl7:"true,PID2"`
	PD12  *PD1  `hl7:"false,PD12"`
	PV12  *PV1  `hl7:"false,PV12"`
	DB12  []DB1 `hl7:"false,DB12"`
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
// Definition from HL7 2.4
type ADT_A30 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PD1   *PD1 `hl7:"false,PD1"`
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
// Definition from HL7 2.4
type ADT_A37 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID1  *PID  `hl7:"true,PID1"`
	PD11  *PD1  `hl7:"false,PD11"`
	PV11  *PV1  `hl7:"false,PV11"`
	DB11  []DB1 `hl7:"false,DB11"`
	PID2  *PID  `hl7:"true,PID2"`
	PD12  *PD1  `hl7:"false,PD12"`
	PV12  *PV1  `hl7:"false,PV12"`
	DB12  []DB1 `hl7:"false,DB12"`
	Other []interface{}
}

func (s *ADT_A37) MessageTypeName() string {
	return "ADT_A37"
}

// ADT_A38 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A38 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PD1   *PD1  `hl7:"false,PD1"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	DB1   []DB1 `hl7:"false,DB1"`
	OBX   []OBX `hl7:"false,OBX"`
	DG1   []DG1 `hl7:"false,DG1"`
	DRG   *DRG  `hl7:"false,DRG"`
	Other []interface{}
}

func (s *ADT_A38) MessageTypeName() string {
	return "ADT_A38"
}

// ADT_A39 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A39 struct {
	MSH     *MSH              `hl7:"true,MSH"`
	EVN     *EVN              `hl7:"true,EVN"`
	PATIENT []ADT_A39_PATIENT `hl7:"true,PATIENT"`
	Other   []interface{}
}

func (s *ADT_A39) MessageTypeName() string {
	return "ADT_A39"
}

// ADT_A39_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A39_PATIENT struct {
	PID   *PID `hl7:"true,PID"`
	PD1   *PD1 `hl7:"false,PD1"`
	MRG   *MRG `hl7:"true,MRG"`
	PV1   *PV1 `hl7:"false,PV1"`
	Other []interface{}
}

func (s *ADT_A39_PATIENT) MessageTypeSubStructName() string {
	return "ADT_A39_PATIENT"
}

// ADT_A43 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A43 struct {
	MSH     *MSH              `hl7:"true,MSH"`
	EVN     *EVN              `hl7:"true,EVN"`
	PATIENT []ADT_A43_PATIENT `hl7:"true,PATIENT"`
	Other   []interface{}
}

func (s *ADT_A43) MessageTypeName() string {
	return "ADT_A43"
}

// ADT_A43_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A43_PATIENT struct {
	PID   *PID `hl7:"true,PID"`
	PD1   *PD1 `hl7:"false,PD1"`
	MRG   *MRG `hl7:"true,MRG"`
	Other []interface{}
}

func (s *ADT_A43_PATIENT) MessageTypeSubStructName() string {
	return "ADT_A43_PATIENT"
}

// ADT_A44 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
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
// Definition from HL7 2.3.1
type ADT_A44_PATIENT struct {
	PID   *PID `hl7:"true,PID"`
	PD1   *PD1 `hl7:"false,PD1"`
	MRG   *MRG `hl7:"true,MRG"`
	Other []interface{}
}

func (s *ADT_A44_PATIENT) MessageTypeSubStructName() string {
	return "ADT_A44_PATIENT"
}

// ADT_A45 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A45 struct {
	MSH        *MSH                 `hl7:"true,MSH"`
	EVN        *EVN                 `hl7:"true,EVN"`
	PID        *PID                 `hl7:"true,PID"`
	PD1        *PD1                 `hl7:"false,PD1"`
	MERGE_INFO []ADT_A45_MERGE_INFO `hl7:"true,MERGE_INFO"`
	Other      []interface{}
}

func (s *ADT_A45) MessageTypeName() string {
	return "ADT_A45"
}

// ADT_A45_MERGE_INFO represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A45_MERGE_INFO struct {
	MRG   *MRG `hl7:"true,MRG"`
	PV1   *PV1 `hl7:"true,PV1"`
	Other []interface{}
}

func (s *ADT_A45_MERGE_INFO) MessageTypeSubStructName() string {
	return "ADT_A45_MERGE_INFO"
}

// ADT_A50 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A50 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PD1   *PD1 `hl7:"false,PD1"`
	MRG   *MRG `hl7:"true,MRG"`
	PV1   *PV1 `hl7:"true,PV1"`
	Other []interface{}
}

func (s *ADT_A50) MessageTypeName() string {
	return "ADT_A50"
}

// ADT_A52 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A52 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PD1   *PD1 `hl7:"false,PD1"`
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *ADT_A52) MessageTypeName() string {
	return "ADT_A52"
}

// ADT_A54 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A54 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PD1   *PD1  `hl7:"false,PD1"`
	ROL1  []ROL `hl7:"false,ROL1"`
	PV1   *PV1  `hl7:"true,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	ROL2  []ROL `hl7:"false,ROL2"`
	Other []interface{}
}

func (s *ADT_A54) MessageTypeName() string {
	return "ADT_A54"
}

// ADT_A60 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A60 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"false,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	IAM   []IAM `hl7:"false,IAM"`
	Other []interface{}
}

func (s *ADT_A60) MessageTypeName() string {
	return "ADT_A60"
}

// ADT_A61 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ADT_A61 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PD1   *PD1  `hl7:"false,PD1"`
	PV1   *PV1  `hl7:"true,PV1"`
	ROL   []ROL `hl7:"false,ROL"`
	PV2   *PV2  `hl7:"false,PV2"`
	Other []interface{}
}

func (s *ADT_A61) MessageTypeName() string {
	return "ADT_A61"
}

// ARD_A19 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ARD_A19 struct {
	MSH            *MSH                     `hl7:"true,MSH"`
	MSA            *MSA                     `hl7:"true,MSA"`
	ERR            *ERR                     `hl7:"false,ERR"`
	QRD            *QRD                     `hl7:"true,QRD"`
	QRF            *QRF                     `hl7:"false,QRF"`
	QUERY_RESPONSE []ARD_A19_QUERY_RESPONSE `hl7:"true,QUERY_RESPONSE"`
	DSC            *DSC                     `hl7:"false,DSC"`
	Other          []interface{}
}

func (s *ARD_A19) MessageTypeName() string {
	return "ARD_A19"
}

// ARD_A19_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ARD_A19_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ARD_A19_INSURANCE) MessageTypeSubStructName() string {
	return "ARD_A19_INSURANCE"
}

// ARD_A19_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ARD_A19_PROCEDURE struct {
	PR1   *PR1  `hl7:"true,PR1"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *ARD_A19_PROCEDURE) MessageTypeSubStructName() string {
	return "ARD_A19_PROCEDURE"
}

// ARD_A19_QUERY_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ARD_A19_QUERY_RESPONSE struct {
	EVN       *EVN                `hl7:"false,EVN"`
	PID       *PID                `hl7:"true,PID"`
	PD1       *PD1                `hl7:"false,PD1"`
	NK1       []NK1               `hl7:"false,NK1"`
	PV1       *PV1                `hl7:"true,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	DB1       []DB1               `hl7:"false,DB1"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	DRG       *DRG                `hl7:"false,DRG"`
	PROCEDURE []ARD_A19_PROCEDURE `hl7:"false,PROCEDURE"`
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []ARD_A19_INSURANCE `hl7:"false,INSURANCE"`
	ACC       *ACC                `hl7:"false,ACC"`
	UB1       *UB1                `hl7:"false,UB1"`
	UB2       *UB2                `hl7:"false,UB2"`
	Other     []interface{}
}

func (s *ARD_A19_QUERY_RESPONSE) MessageTypeSubStructName() string {
	return "ARD_A19_QUERY_RESPONSE"
}

// BAR_P01 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type BAR_P01 struct {
	MSH   *MSH            `hl7:"true,MSH"`
	EVN   *EVN            `hl7:"true,EVN"`
	PID   *PID            `hl7:"true,PID"`
	PD1   *PD1            `hl7:"false,PD1"`
	ROL   []ROL           `hl7:"false,ROL"`
	VISIT []BAR_P01_VISIT `hl7:"true,VISIT"`
	Other []interface{}
}

func (s *BAR_P01) MessageTypeName() string {
	return "BAR_P01"
}

// BAR_P01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type BAR_P01_INSURANCE struct {
	IN1   *IN1  `hl7:"true,IN1"`
	IN2   *IN2  `hl7:"false,IN2"`
	IN3   []IN3 `hl7:"false,IN3"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *BAR_P01_INSURANCE) MessageTypeSubStructName() string {
	return "BAR_P01_INSURANCE"
}

// BAR_P01_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type BAR_P01_PROCEDURE struct {
	PR1   *PR1  `hl7:"true,PR1"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *BAR_P01_PROCEDURE) MessageTypeSubStructName() string {
	return "BAR_P01_PROCEDURE"
}

// BAR_P01_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type BAR_P01_VISIT struct {
	PV1       *PV1                `hl7:"false,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	ROL       []ROL               `hl7:"false,ROL"`
	DB1       []DB1               `hl7:"false,DB1"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	DRG       *DRG                `hl7:"false,DRG"`
	PROCEDURE []BAR_P01_PROCEDURE `hl7:"false,PROCEDURE"`
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
// Definition from HL7 2.4
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
// Definition from HL7 2.4
type BAR_P02_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	PD1   *PD1  `hl7:"false,PD1"`
	PV1   *PV1  `hl7:"false,PV1"`
	DB1   []DB1 `hl7:"false,DB1"`
	Other []interface{}
}

func (s *BAR_P02_PATIENT) MessageTypeSubStructName() string {
	return "BAR_P02_PATIENT"
}

// BAR_P05 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type BAR_P05 struct {
	MSH   *MSH            `hl7:"true,MSH"`
	EVN   *EVN            `hl7:"true,EVN"`
	PID   *PID            `hl7:"true,PID"`
	PD1   *PD1            `hl7:"false,PD1"`
	ROL   []ROL           `hl7:"false,ROL"`
	VISIT []BAR_P05_VISIT `hl7:"true,VISIT"`
	Other []interface{}
}

func (s *BAR_P05) MessageTypeName() string {
	return "BAR_P05"
}

// BAR_P05_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type BAR_P05_INSURANCE struct {
	IN1   *IN1  `hl7:"true,IN1"`
	IN2   *IN2  `hl7:"false,IN2"`
	IN3   []IN3 `hl7:"false,IN3"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *BAR_P05_INSURANCE) MessageTypeSubStructName() string {
	return "BAR_P05_INSURANCE"
}

// BAR_P05_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type BAR_P05_PROCEDURE struct {
	PR1   *PR1  `hl7:"true,PR1"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *BAR_P05_PROCEDURE) MessageTypeSubStructName() string {
	return "BAR_P05_PROCEDURE"
}

// BAR_P05_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type BAR_P05_VISIT struct {
	PV1       *PV1                `hl7:"false,PV1"`
	PV2       *PV2                `hl7:"false,PV2"`
	ROL       []ROL               `hl7:"false,ROL"`
	DB1       []DB1               `hl7:"false,DB1"`
	OBX       []OBX               `hl7:"false,OBX"`
	AL1       []AL1               `hl7:"false,AL1"`
	DG1       []DG1               `hl7:"false,DG1"`
	DRG       *DRG                `hl7:"false,DRG"`
	PROCEDURE []BAR_P05_PROCEDURE `hl7:"false,PROCEDURE"`
	GT1       []GT1               `hl7:"false,GT1"`
	NK1       []NK1               `hl7:"false,NK1"`
	INSURANCE []BAR_P05_INSURANCE `hl7:"false,INSURANCE"`
	ACC       *ACC                `hl7:"false,ACC"`
	UB1       *UB1                `hl7:"false,UB1"`
	UB2       *UB2                `hl7:"false,UB2"`
	ABS       *ABS                `hl7:"false,ABS"`
	BLC       []BLC               `hl7:"false,BLC"`
	RMI       *RMI                `hl7:"false,RMI"`
	Other     []interface{}
}

func (s *BAR_P05_VISIT) MessageTypeSubStructName() string {
	return "BAR_P05_VISIT"
}

// BAR_P06 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type BAR_P06 struct {
	MSH     *MSH              `hl7:"true,MSH"`
	EVN     *EVN              `hl7:"true,EVN"`
	PATIENT []BAR_P06_PATIENT `hl7:"true,PATIENT"`
	Other   []interface{}
}

func (s *BAR_P06) MessageTypeName() string {
	return "BAR_P06"
}

// BAR_P06_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type BAR_P06_PATIENT struct {
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"false,PV1"`
	Other []interface{}
}

func (s *BAR_P06_PATIENT) MessageTypeSubStructName() string {
	return "BAR_P06_PATIENT"
}

// BAR_P10 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type BAR_P10 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	EVN       *EVN                `hl7:"true,EVN"`
	PID       *PID                `hl7:"true,PID"`
	PV1       *PV1                `hl7:"true,PV1"`
	DG1       []DG1               `hl7:"false,DG1"`
	GP1       *GP1                `hl7:"true,GP1"`
	PROCEDURE []BAR_P10_PROCEDURE `hl7:"false,PROCEDURE"`
	Other     []interface{}
}

func (s *BAR_P10) MessageTypeName() string {
	return "BAR_P10"
}

// BAR_P10_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type BAR_P10_PROCEDURE struct {
	PR1   *PR1 `hl7:"true,PR1"`
	GP2   *GP2 `hl7:"false,GP2"`
	Other []interface{}
}

func (s *BAR_P10_PROCEDURE) MessageTypeSubStructName() string {
	return "BAR_P10_PROCEDURE"
}

// CRM_C01 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type CRM_C01 struct {
	MSH     *MSH              `hl7:"true,MSH"`
	PATIENT []CRM_C01_PATIENT `hl7:"true,PATIENT"`
	Other   []interface{}
}

func (s *CRM_C01) MessageTypeName() string {
	return "CRM_C01"
}

// CRM_C01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type CRM_C01_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"false,PV1"`
	CSR   *CSR  `hl7:"true,CSR"`
	CSP   []CSP `hl7:"false,CSP"`
	Other []interface{}
}

func (s *CRM_C01_PATIENT) MessageTypeSubStructName() string {
	return "CRM_C01_PATIENT"
}

// CSU_C09 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type CSU_C09 struct {
	MSH     *MSH              `hl7:"true,MSH"`
	PATIENT []CSU_C09_PATIENT `hl7:"true,PATIENT"`
	Other   []interface{}
}

func (s *CSU_C09) MessageTypeName() string {
	return "CSU_C09"
}

// CSU_C09_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type CSU_C09_PATIENT struct {
	PID         *PID                  `hl7:"true,PID"`
	PD1         *PD1                  `hl7:"false,PD1"`
	NTE         []NTE                 `hl7:"false,NTE"`
	VISIT       *CSU_C09_VISIT        `hl7:"false,VISIT"`
	CSR         *CSR                  `hl7:"true,CSR"`
	STUDY_PHASE []CSU_C09_STUDY_PHASE `hl7:"true,STUDY_PHASE"`
	Other       []interface{}
}

func (s *CSU_C09_PATIENT) MessageTypeSubStructName() string {
	return "CSU_C09_PATIENT"
}

// CSU_C09_RX_ADMIN represents the corresponding HL7 message type.
// Definition from HL7 2.4
type CSU_C09_RX_ADMIN struct {
	RXA   *RXA `hl7:"true,RXA"`
	RXR   *RXR `hl7:"true,RXR"`
	Other []interface{}
}

func (s *CSU_C09_RX_ADMIN) MessageTypeSubStructName() string {
	return "CSU_C09_RX_ADMIN"
}

// CSU_C09_STUDY_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type CSU_C09_STUDY_OBSERVATION struct {
	ORC   *ORC  `hl7:"false,ORC"`
	OBR   *OBR  `hl7:"true,OBR"`
	OBX   []OBX `hl7:"true,OBX"`
	Other []interface{}
}

func (s *CSU_C09_STUDY_OBSERVATION) MessageTypeSubStructName() string {
	return "CSU_C09_STUDY_OBSERVATION"
}

// CSU_C09_STUDY_PHARM represents the corresponding HL7 message type.
// Definition from HL7 2.4
type CSU_C09_STUDY_PHARM struct {
	ORC      *ORC               `hl7:"false,ORC"`
	RX_ADMIN []CSU_C09_RX_ADMIN `hl7:"true,RX_ADMIN"`
	Other    []interface{}
}

func (s *CSU_C09_STUDY_PHARM) MessageTypeSubStructName() string {
	return "CSU_C09_STUDY_PHARM"
}

// CSU_C09_STUDY_PHASE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type CSU_C09_STUDY_PHASE struct {
	CSP            []CSP                    `hl7:"false,CSP"`
	STUDY_SCHEDULE []CSU_C09_STUDY_SCHEDULE `hl7:"true,STUDY_SCHEDULE"`
	Other          []interface{}
}

func (s *CSU_C09_STUDY_PHASE) MessageTypeSubStructName() string {
	return "CSU_C09_STUDY_PHASE"
}

// CSU_C09_STUDY_SCHEDULE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type CSU_C09_STUDY_SCHEDULE struct {
	CSS               *CSS                        `hl7:"false,CSS"`
	STUDY_OBSERVATION []CSU_C09_STUDY_OBSERVATION `hl7:"true,STUDY_OBSERVATION"`
	STUDY_PHARM       []CSU_C09_STUDY_PHARM       `hl7:"true,STUDY_PHARM"`
	Other             []interface{}
}

func (s *CSU_C09_STUDY_SCHEDULE) MessageTypeSubStructName() string {
	return "CSU_C09_STUDY_SCHEDULE"
}

// CSU_C09_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type CSU_C09_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *CSU_C09_VISIT) MessageTypeSubStructName() string {
	return "CSU_C09_VISIT"
}

// DFT_P03_COMMON_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P03_COMMON_ORDER struct {
	ORC         *ORC                  `hl7:"false,ORC"`
	ORDER       *DFT_P03_ORDER        `hl7:"false,ORDER"`
	OBSERVATION []DFT_P03_OBSERVATION `hl7:"false,OBSERVATION"`
	Other       []interface{}
}

func (s *DFT_P03_COMMON_ORDER) MessageTypeSubStructName() string {
	return "DFT_P03_COMMON_ORDER"
}

// DFT_P03 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P03 struct {
	MSH          *MSH                   `hl7:"true,MSH"`
	EVN          *EVN                   `hl7:"true,EVN"`
	PID          *PID                   `hl7:"true,PID"`
	PD1          *PD1                   `hl7:"false,PD1"`
	ROL1         []ROL                  `hl7:"false,ROL1"`
	PV1          *PV1                   `hl7:"false,PV1"`
	PV2          *PV2                   `hl7:"false,PV2"`
	ROL2         []ROL                  `hl7:"false,ROL2"`
	DB1          []DB1                  `hl7:"false,DB1"`
	COMMON_ORDER []DFT_P03_COMMON_ORDER `hl7:"false,COMMON_ORDER"`
	FINANCIAL    []DFT_P03_FINANCIAL    `hl7:"true,FINANCIAL"`
	DG1          []DG1                  `hl7:"false,DG1"`
	DRG          *DRG                   `hl7:"false,DRG"`
	GT1          []GT1                  `hl7:"false,GT1"`
	INSURANCE    []DFT_P03_INSURANCE    `hl7:"false,INSURANCE"`
	ACC          *ACC                   `hl7:"false,ACC"`
	Other        []interface{}
}

func (s *DFT_P03) MessageTypeName() string {
	return "DFT_P03"
}

// DFT_P03_FINANCIAL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P03_FINANCIAL struct {
	FT1                    *FT1                             `hl7:"true,FT1"`
	FINANCIAL_PROCEDURE    []DFT_P03_FINANCIAL_PROCEDURE    `hl7:"false,FINANCIAL_PROCEDURE"`
	FINANCIAL_COMMON_ORDER []DFT_P03_FINANCIAL_COMMON_ORDER `hl7:"false,FINANCIAL_COMMON_ORDER"`
	Other                  []interface{}
}

func (s *DFT_P03_FINANCIAL) MessageTypeSubStructName() string {
	return "DFT_P03_FINANCIAL"
}

// DFT_P03_FINANCIAL_COMMON_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P03_FINANCIAL_COMMON_ORDER struct {
	ORC                   *ORC                            `hl7:"false,ORC"`
	FINANCIAL_ORDER       *DFT_P03_FINANCIAL_ORDER        `hl7:"false,FINANCIAL_ORDER"`
	FINANCIAL_OBSERVATION []DFT_P03_FINANCIAL_OBSERVATION `hl7:"false,FINANCIAL_OBSERVATION"`
	Other                 []interface{}
}

func (s *DFT_P03_FINANCIAL_COMMON_ORDER) MessageTypeSubStructName() string {
	return "DFT_P03_FINANCIAL_COMMON_ORDER"
}

// DFT_P03_FINANCIAL_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P03_FINANCIAL_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *DFT_P03_FINANCIAL_OBSERVATION) MessageTypeSubStructName() string {
	return "DFT_P03_FINANCIAL_OBSERVATION"
}

// DFT_P03_FINANCIAL_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P03_FINANCIAL_ORDER struct {
	OBR   *OBR  `hl7:"true,OBR"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *DFT_P03_FINANCIAL_ORDER) MessageTypeSubStructName() string {
	return "DFT_P03_FINANCIAL_ORDER"
}

// DFT_P03_FINANCIAL_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P03_FINANCIAL_PROCEDURE struct {
	PR1   *PR1  `hl7:"true,PR1"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *DFT_P03_FINANCIAL_PROCEDURE) MessageTypeSubStructName() string {
	return "DFT_P03_FINANCIAL_PROCEDURE"
}

// DFT_P03_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P03_INSURANCE struct {
	IN1   *IN1  `hl7:"true,IN1"`
	IN2   *IN2  `hl7:"false,IN2"`
	IN3   []IN3 `hl7:"false,IN3"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *DFT_P03_INSURANCE) MessageTypeSubStructName() string {
	return "DFT_P03_INSURANCE"
}

// DFT_P03_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P03_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *DFT_P03_OBSERVATION) MessageTypeSubStructName() string {
	return "DFT_P03_OBSERVATION"
}

// DFT_P03_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P03_ORDER struct {
	OBR   *OBR  `hl7:"true,OBR"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *DFT_P03_ORDER) MessageTypeSubStructName() string {
	return "DFT_P03_ORDER"
}

// DFT_P11_COMMON_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P11_COMMON_ORDER struct {
	ORC         *ORC                  `hl7:"false,ORC"`
	ORDER       *DFT_P11_ORDER        `hl7:"false,ORDER"`
	OBSERVATION []DFT_P11_OBSERVATION `hl7:"false,OBSERVATION"`
	Other       []interface{}
}

func (s *DFT_P11_COMMON_ORDER) MessageTypeSubStructName() string {
	return "DFT_P11_COMMON_ORDER"
}

// DFT_P11 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P11 struct {
	MSH          *MSH                   `hl7:"true,MSH"`
	EVN          *EVN                   `hl7:"true,EVN"`
	PID          *PID                   `hl7:"true,PID"`
	PD1          *PD1                   `hl7:"false,PD1"`
	ROL1         []ROL                  `hl7:"false,ROL1"`
	PV1          *PV1                   `hl7:"false,PV1"`
	PV2          *PV2                   `hl7:"false,PV2"`
	ROL2         []ROL                  `hl7:"false,ROL2"`
	DB1          []DB1                  `hl7:"false,DB1"`
	COMMON_ORDER []DFT_P11_COMMON_ORDER `hl7:"false,COMMON_ORDER"`
	DG1          []DG1                  `hl7:"false,DG1"`
	DRG          *DRG                   `hl7:"false,DRG"`
	GT1          []GT1                  `hl7:"false,GT1"`
	INSURANCE    []DFT_P11_INSURANCE    `hl7:"false,INSURANCE"`
	ACC          *ACC                   `hl7:"false,ACC"`
	FINANCIAL    []DFT_P11_FINANCIAL    `hl7:"true,FINANCIAL"`
	Other        []interface{}
}

func (s *DFT_P11) MessageTypeName() string {
	return "DFT_P11"
}

// DFT_P11_FINANCIAL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P11_FINANCIAL struct {
	FT1                    *FT1                             `hl7:"true,FT1"`
	FINANCIAL_PROCEDURE    []DFT_P11_FINANCIAL_PROCEDURE    `hl7:"false,FINANCIAL_PROCEDURE"`
	FINANCIAL_COMMON_ORDER []DFT_P11_FINANCIAL_COMMON_ORDER `hl7:"false,FINANCIAL_COMMON_ORDER"`
	DG1                    []DG1                            `hl7:"false,DG1"`
	DRG                    *DRG                             `hl7:"false,DRG"`
	GT1                    []GT1                            `hl7:"false,GT1"`
	FINANCIAL_INSURANCE    []DFT_P11_FINANCIAL_INSURANCE    `hl7:"false,FINANCIAL_INSURANCE"`
	Other                  []interface{}
}

func (s *DFT_P11_FINANCIAL) MessageTypeSubStructName() string {
	return "DFT_P11_FINANCIAL"
}

// DFT_P11_FINANCIAL_COMMON_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P11_FINANCIAL_COMMON_ORDER struct {
	ORC                   *ORC                            `hl7:"false,ORC"`
	FINANCIAL_ORDER       *DFT_P11_FINANCIAL_ORDER        `hl7:"false,FINANCIAL_ORDER"`
	FINANCIAL_OBSERVATION []DFT_P11_FINANCIAL_OBSERVATION `hl7:"false,FINANCIAL_OBSERVATION"`
	Other                 []interface{}
}

func (s *DFT_P11_FINANCIAL_COMMON_ORDER) MessageTypeSubStructName() string {
	return "DFT_P11_FINANCIAL_COMMON_ORDER"
}

// DFT_P11_FINANCIAL_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P11_FINANCIAL_INSURANCE struct {
	IN1   *IN1  `hl7:"true,IN1"`
	IN2   *IN2  `hl7:"false,IN2"`
	IN3   []IN3 `hl7:"false,IN3"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *DFT_P11_FINANCIAL_INSURANCE) MessageTypeSubStructName() string {
	return "DFT_P11_FINANCIAL_INSURANCE"
}

// DFT_P11_FINANCIAL_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P11_FINANCIAL_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *DFT_P11_FINANCIAL_OBSERVATION) MessageTypeSubStructName() string {
	return "DFT_P11_FINANCIAL_OBSERVATION"
}

// DFT_P11_FINANCIAL_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P11_FINANCIAL_ORDER struct {
	OBR   *OBR  `hl7:"true,OBR"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *DFT_P11_FINANCIAL_ORDER) MessageTypeSubStructName() string {
	return "DFT_P11_FINANCIAL_ORDER"
}

// DFT_P11_FINANCIAL_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P11_FINANCIAL_PROCEDURE struct {
	PR1   *PR1  `hl7:"true,PR1"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *DFT_P11_FINANCIAL_PROCEDURE) MessageTypeSubStructName() string {
	return "DFT_P11_FINANCIAL_PROCEDURE"
}

// DFT_P11_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P11_INSURANCE struct {
	IN1   *IN1  `hl7:"true,IN1"`
	IN2   *IN2  `hl7:"false,IN2"`
	IN3   []IN3 `hl7:"false,IN3"`
	ROL   []ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *DFT_P11_INSURANCE) MessageTypeSubStructName() string {
	return "DFT_P11_INSURANCE"
}

// DFT_P11_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P11_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *DFT_P11_OBSERVATION) MessageTypeSubStructName() string {
	return "DFT_P11_OBSERVATION"
}

// DFT_P11_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DFT_P11_ORDER struct {
	OBR   *OBR  `hl7:"true,OBR"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *DFT_P11_ORDER) MessageTypeSubStructName() string {
	return "DFT_P11_ORDER"
}

// DOC_T12 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DOC_T12 struct {
	MSH    *MSH             `hl7:"true,MSH"`
	MSA    *MSA             `hl7:"true,MSA"`
	ERR    *ERR             `hl7:"false,ERR"`
	QAK    *QAK             `hl7:"false,QAK"`
	QRD    *QRD             `hl7:"true,QRD"`
	RESULT []DOC_T12_RESULT `hl7:"true,RESULT"`
	DSC    *DSC             `hl7:"false,DSC"`
	Other  []interface{}
}

func (s *DOC_T12) MessageTypeName() string {
	return "DOC_T12"
}

// DOC_T12_RESULT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type DOC_T12_RESULT struct {
	EVN   *EVN  `hl7:"false,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	TXA   *TXA  `hl7:"true,TXA"`
	OBX   []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *DOC_T12_RESULT) MessageTypeSubStructName() string {
	return "DOC_T12_RESULT"
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
// Definition from HL7 2.4
type DSR_Q01 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MSA   *MSA  `hl7:"true,MSA"`
	ERR   *ERR  `hl7:"false,ERR"`
	QAK   *QAK  `hl7:"false,QAK"`
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
// Definition from HL7 2.4
type DSR_Q03 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MSA   *MSA  `hl7:"false,MSA"`
	ERR   *ERR  `hl7:"false,ERR"`
	QAK   *QAK  `hl7:"false,QAK"`
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

// EAC_U07 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type EAC_U07 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EQU   *EQU  `hl7:"true,EQU"`
	ECD   []ECD `hl7:"true,ECD"`
	SAC   *SAC  `hl7:"false,SAC"`
	CNS   *CNS  `hl7:"false,CNS"`
	ROL   *ROL  `hl7:"false,ROL"`
	Other []interface{}
}

func (s *EAC_U07) MessageTypeName() string {
	return "EAC_U07"
}

// EAN_U09 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type EAN_U09 struct {
	MSH          *MSH                   `hl7:"true,MSH"`
	EQU          *EQU                   `hl7:"true,EQU"`
	NOTIFICATION []EAN_U09_NOTIFICATION `hl7:"true,NOTIFICATION"`
	ROL          *ROL                   `hl7:"false,ROL"`
	Other        []interface{}
}

func (s *EAN_U09) MessageTypeName() string {
	return "EAN_U09"
}

// EAN_U09_NOTIFICATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type EAN_U09_NOTIFICATION struct {
	NDS   *NDS `hl7:"true,NDS"`
	NTE   *NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *EAN_U09_NOTIFICATION) MessageTypeSubStructName() string {
	return "EAN_U09_NOTIFICATION"
}

// EAR_U08_COMMAND_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type EAR_U08_COMMAND_RESPONSE struct {
	ECD   *ECD `hl7:"true,ECD"`
	SAC   *SAC `hl7:"false,SAC"`
	ECR   *ECR `hl7:"true,ECR"`
	Other []interface{}
}

func (s *EAR_U08_COMMAND_RESPONSE) MessageTypeSubStructName() string {
	return "EAR_U08_COMMAND_RESPONSE"
}

// EAR_U08 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type EAR_U08 struct {
	MSH              *MSH                       `hl7:"true,MSH"`
	EQU              *EQU                       `hl7:"true,EQU"`
	COMMAND_RESPONSE []EAR_U08_COMMAND_RESPONSE `hl7:"true,COMMAND_RESPONSE"`
	ROL              *ROL                       `hl7:"false,ROL"`
	Other            []interface{}
}

func (s *EAR_U08) MessageTypeName() string {
	return "EAR_U08"
}

// EDR_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type EDR_Q01 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MSA   *MSA  `hl7:"true,MSA"`
	ERR   *ERR  `hl7:"false,ERR"`
	QAK   *QAK  `hl7:"true,QAK"`
	DSP   []DSP `hl7:"true,DSP"`
	DSC   *DSC  `hl7:"false,DSC"`
	Other []interface{}
}

func (s *EDR_Q01) MessageTypeName() string {
	return "EDR_Q01"
}

// EDR_R07 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type EDR_R07 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MSA   *MSA  `hl7:"true,MSA"`
	ERR   *ERR  `hl7:"false,ERR"`
	QAK   *QAK  `hl7:"true,QAK"`
	DSP   []DSP `hl7:"true,DSP"`
	DSC   *DSC  `hl7:"false,DSC"`
	Other []interface{}
}

func (s *EDR_R07) MessageTypeName() string {
	return "EDR_R07"
}

// EQQ_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type EQQ_Q01 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EQL   *EQL `hl7:"true,EQL"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *EQQ_Q01) MessageTypeName() string {
	return "EQQ_Q01"
}

// EQQ_Q04 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type EQQ_Q04 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EQL   *EQL `hl7:"true,EQL"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *EQQ_Q04) MessageTypeName() string {
	return "EQQ_Q04"
}

// ERP_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ERP_Q01 struct {
	MSH   *MSH `hl7:"true,MSH"`
	MSA   *MSA `hl7:"true,MSA"`
	ERR   *ERR `hl7:"false,ERR"`
	QAK   *QAK `hl7:"true,QAK"`
	ERQ   *ERQ `hl7:"true,ERQ"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *ERP_Q01) MessageTypeName() string {
	return "ERP_Q01"
}

// ERP_R09 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ERP_R09 struct {
	MSH   *MSH `hl7:"true,MSH"`
	MSA   *MSA `hl7:"true,MSA"`
	ERR   *ERR `hl7:"false,ERR"`
	QAK   *QAK `hl7:"true,QAK"`
	ERQ   *ERQ `hl7:"true,ERQ"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *ERP_R09) MessageTypeName() string {
	return "ERP_R09"
}

// ESR_U02 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ESR_U02 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EQU   *EQU `hl7:"true,EQU"`
	ROL   *ROL `hl7:"false,ROL"`
	Other []interface{}
}

func (s *ESR_U02) MessageTypeName() string {
	return "ESR_U02"
}

// ESU_U01 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ESU_U01 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EQU   *EQU  `hl7:"true,EQU"`
	ISD   []ISD `hl7:"false,ISD"`
	ROL   *ROL  `hl7:"false,ROL"`
	Other []interface{}
}

func (s *ESU_U01) MessageTypeName() string {
	return "ESU_U01"
}

// INR_U06 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type INR_U06 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EQU   *EQU  `hl7:"true,EQU"`
	INV   []INV `hl7:"true,INV"`
	ROL   *ROL  `hl7:"false,ROL"`
	Other []interface{}
}

func (s *INR_U06) MessageTypeName() string {
	return "INR_U06"
}

// INU_U05 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type INU_U05 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EQU   *EQU  `hl7:"true,EQU"`
	INV   []INV `hl7:"true,INV"`
	ROL   *ROL  `hl7:"false,ROL"`
	Other []interface{}
}

func (s *INU_U05) MessageTypeName() string {
	return "INU_U05"
}

// LSU_U12 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type LSU_U12 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EQU   *EQU  `hl7:"true,EQU"`
	EQP   []EQP `hl7:"true,EQP"`
	ROL   *ROL  `hl7:"false,ROL"`
	Other []interface{}
}

func (s *LSU_U12) MessageTypeName() string {
	return "LSU_U12"
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

// MDM_T01 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MDM_T01 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"true,PV1"`
	TXA   *TXA `hl7:"true,TXA"`
	Other []interface{}
}

func (s *MDM_T01) MessageTypeName() string {
	return "MDM_T01"
}

// MDM_T02 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MDM_T02 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"true,PV1"`
	TXA   *TXA  `hl7:"true,TXA"`
	OBX   []OBX `hl7:"true,OBX"`
	Other []interface{}
}

func (s *MDM_T02) MessageTypeName() string {
	return "MDM_T02"
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
// Definition from HL7 2.4
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
// Definition from HL7 2.3
type MFK_M02 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MSA   *MSA  `hl7:"true,MSA"`
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
// Definition from HL7 2.4
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
// Definition from HL7 2.4
type MFN_M01_MF struct {
	MFE *MFE `hl7:"true,MFE"`
	// Missing: anyZSegment
	Other []interface{}
}

func (s *MFN_M01_MF) MessageTypeSubStructName() string {
	return "MFN_M01_MF"
}

// MFN_M02 represents the corresponding HL7 message type.
// Definition from HL7 2.4
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
// Definition from HL7 2.4
type MFN_M02_MF_STAFF struct {
	MFE   *MFE `hl7:"true,MFE"`
	STF   *STF `hl7:"true,STF"`
	PRA   *PRA `hl7:"false,PRA"`
	ORG   *ORG `hl7:"false,ORG"`
	Other []interface{}
}

func (s *MFN_M02_MF_STAFF) MessageTypeSubStructName() string {
	return "MFN_M02_MF_STAFF"
}

// MFN_M03 represents the corresponding HL7 message type.
// Definition from HL7 2.4
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
// Definition from HL7 2.4
type MFN_M03_MF_TEST struct {
	MFE *MFE `hl7:"true,MFE"`
	OM1 *OM1 `hl7:"true,OM1"`
	// Missing: anyHL7Segment
	Other []interface{}
}

func (s *MFN_M03_MF_TEST) MessageTypeSubStructName() string {
	return "MFN_M03_MF_TEST"
}

// MFN_M04 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M04 struct {
	MSH    *MSH             `hl7:"true,MSH"`
	MFI    *MFI             `hl7:"true,MFI"`
	MF_CDM []MFN_M04_MF_CDM `hl7:"true,MF_CDM"`
	Other  []interface{}
}

func (s *MFN_M04) MessageTypeName() string {
	return "MFN_M04"
}

// MFN_M04_MF_CDM represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M04_MF_CDM struct {
	MFE   *MFE  `hl7:"true,MFE"`
	CDM   *CDM  `hl7:"true,CDM"`
	PRC   []PRC `hl7:"false,PRC"`
	Other []interface{}
}

func (s *MFN_M04_MF_CDM) MessageTypeSubStructName() string {
	return "MFN_M04_MF_CDM"
}

// MFN_M05 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M05 struct {
	MSH         *MSH                  `hl7:"true,MSH"`
	MFI         *MFI                  `hl7:"true,MFI"`
	MF_LOCATION []MFN_M05_MF_LOCATION `hl7:"true,MF_LOCATION"`
	Other       []interface{}
}

func (s *MFN_M05) MessageTypeName() string {
	return "MFN_M05"
}

// MFN_M05_MF_LOCATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M05_MF_LOCATION struct {
	MFE         *MFE                  `hl7:"true,MFE"`
	LOC         *LOC                  `hl7:"true,LOC"`
	LCH         []LCH                 `hl7:"false,LCH"`
	LRL         []LRL                 `hl7:"false,LRL"`
	MF_LOC_DEPT []MFN_M05_MF_LOC_DEPT `hl7:"true,MF_LOC_DEPT"`
	Other       []interface{}
}

func (s *MFN_M05_MF_LOCATION) MessageTypeSubStructName() string {
	return "MFN_M05_MF_LOCATION"
}

// MFN_M05_MF_LOC_DEPT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M05_MF_LOC_DEPT struct {
	LDP   *LDP  `hl7:"true,LDP"`
	LCH   []LCH `hl7:"false,LCH"`
	LCC   []LCC `hl7:"false,LCC"`
	Other []interface{}
}

func (s *MFN_M05_MF_LOC_DEPT) MessageTypeSubStructName() string {
	return "MFN_M05_MF_LOC_DEPT"
}

// MFN_M06 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M06 struct {
	MSH           *MSH                    `hl7:"true,MSH"`
	MFI           *MFI                    `hl7:"true,MFI"`
	MF_CLIN_STUDY []MFN_M06_MF_CLIN_STUDY `hl7:"true,MF_CLIN_STUDY"`
	Other         []interface{}
}

func (s *MFN_M06) MessageTypeName() string {
	return "MFN_M06"
}

// MFN_M06_MF_CDM represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M06_MF_CDM struct {
	MFE   *MFE  `hl7:"true,MFE"`
	CDM   *CDM  `hl7:"true,CDM"`
	PRC   []PRC `hl7:"false,PRC"`
	Other []interface{}
}

func (s *MFN_M06_MF_CDM) MessageTypeSubStructName() string {
	return "MFN_M06_MF_CDM"
}

// MFN_M06_MF_CLIN_STUDY represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M06_MF_CLIN_STUDY struct {
	MFE                   *MFE                            `hl7:"true,MFE"`
	CM0                   *CM0                            `hl7:"true,CM0"`
	MF_PHASE_SCHED_DETAIL []MFN_M06_MF_PHASE_SCHED_DETAIL `hl7:"false,MF_PHASE_SCHED_DETAIL"`
	Other                 []interface{}
}

func (s *MFN_M06_MF_CLIN_STUDY) MessageTypeSubStructName() string {
	return "MFN_M06_MF_CLIN_STUDY"
}

// MFN_M06_MF_PHASE_SCHED_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M06_MF_PHASE_SCHED_DETAIL struct {
	CM1   *CM1  `hl7:"true,CM1"`
	CM2   []CM2 `hl7:"false,CM2"`
	Other []interface{}
}

func (s *MFN_M06_MF_PHASE_SCHED_DETAIL) MessageTypeSubStructName() string {
	return "MFN_M06_MF_PHASE_SCHED_DETAIL"
}

// MFN_M07 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M07 struct {
	MSH                 *MSH                          `hl7:"true,MSH"`
	MFI                 *MFI                          `hl7:"true,MFI"`
	MF_CLIN_STUDY_SCHED []MFN_M07_MF_CLIN_STUDY_SCHED `hl7:"true,MF_CLIN_STUDY_SCHED"`
	Other               []interface{}
}

func (s *MFN_M07) MessageTypeName() string {
	return "MFN_M07"
}

// MFN_M07_MF_CLIN_STUDY represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M07_MF_CLIN_STUDY struct {
	MFE                   *MFE                            `hl7:"true,MFE"`
	CM0                   *CM0                            `hl7:"true,CM0"`
	MF_PHASE_SCHED_DETAIL []MFN_M07_MF_PHASE_SCHED_DETAIL `hl7:"false,MF_PHASE_SCHED_DETAIL"`
	Other                 []interface{}
}

func (s *MFN_M07_MF_CLIN_STUDY) MessageTypeSubStructName() string {
	return "MFN_M07_MF_CLIN_STUDY"
}

// MFN_M07_MF_CLIN_STUDY_SCHED represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M07_MF_CLIN_STUDY_SCHED struct {
	MFE   *MFE  `hl7:"true,MFE"`
	CM0   *CM0  `hl7:"true,CM0"`
	CM2   []CM2 `hl7:"false,CM2"`
	Other []interface{}
}

func (s *MFN_M07_MF_CLIN_STUDY_SCHED) MessageTypeSubStructName() string {
	return "MFN_M07_MF_CLIN_STUDY_SCHED"
}

// MFN_M07_MF_PHASE_SCHED_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3
type MFN_M07_MF_PHASE_SCHED_DETAIL struct {
	CM1   *CM1  `hl7:"true,CM1"`
	CM2   []CM2 `hl7:"false,CM2"`
	Other []interface{}
}

func (s *MFN_M07_MF_PHASE_SCHED_DETAIL) MessageTypeSubStructName() string {
	return "MFN_M07_MF_PHASE_SCHED_DETAIL"
}

// MFN_M08 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M08 struct {
	MSH             *MSH                      `hl7:"true,MSH"`
	MFI             *MFI                      `hl7:"true,MFI"`
	MF_TEST_NUMERIC []MFN_M08_MF_TEST_NUMERIC `hl7:"true,MF_TEST_NUMERIC"`
	Other           []interface{}
}

func (s *MFN_M08) MessageTypeName() string {
	return "MFN_M08"
}

// MFN_M08_MF_NUMERIC_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type MFN_M08_MF_NUMERIC_OBSERVATION struct {
	OM2   *OM2 `hl7:"false,OM2"`
	OM3   *OM3 `hl7:"false,OM3"`
	OM4   *OM4 `hl7:"false,OM4"`
	Other []interface{}
}

func (s *MFN_M08_MF_NUMERIC_OBSERVATION) MessageTypeSubStructName() string {
	return "MFN_M08_MF_NUMERIC_OBSERVATION"
}

// MFN_M08_MF_TEST_NUMERIC represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M08_MF_TEST_NUMERIC struct {
	MFE   *MFE `hl7:"true,MFE"`
	OM1   *OM1 `hl7:"true,OM1"`
	OM2   *OM2 `hl7:"false,OM2"`
	OM3   *OM3 `hl7:"false,OM3"`
	OM4   *OM4 `hl7:"false,OM4"`
	Other []interface{}
}

func (s *MFN_M08_MF_TEST_NUMERIC) MessageTypeSubStructName() string {
	return "MFN_M08_MF_TEST_NUMERIC"
}

// MFN_M09 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M09 struct {
	MSH                 *MSH                          `hl7:"true,MSH"`
	MFI                 *MFI                          `hl7:"true,MFI"`
	MF_TEST_CATEGORICAL []MFN_M09_MF_TEST_CATEGORICAL `hl7:"true,MF_TEST_CATEGORICAL"`
	Other               []interface{}
}

func (s *MFN_M09) MessageTypeName() string {
	return "MFN_M09"
}

// MFN_M09_MF_TEST_CATEGORICAL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M09_MF_TEST_CATEGORICAL struct {
	MFE                *MFE                        `hl7:"true,MFE"`
	OM1                *OM1                        `hl7:"true,OM1"`
	MF_TEST_CAT_DETAIL *MFN_M09_MF_TEST_CAT_DETAIL `hl7:"false,MF_TEST_CAT_DETAIL"`
	Other              []interface{}
}

func (s *MFN_M09_MF_TEST_CATEGORICAL) MessageTypeSubStructName() string {
	return "MFN_M09_MF_TEST_CATEGORICAL"
}

// MFN_M09_MF_TEST_CAT_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M09_MF_TEST_CAT_DETAIL struct {
	OM3   *OM3  `hl7:"true,OM3"`
	OM4   []OM4 `hl7:"false,OM4"`
	Other []interface{}
}

func (s *MFN_M09_MF_TEST_CAT_DETAIL) MessageTypeSubStructName() string {
	return "MFN_M09_MF_TEST_CAT_DETAIL"
}

// MFN_M10 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M10 struct {
	MSH               *MSH                        `hl7:"true,MSH"`
	MFI               *MFI                        `hl7:"true,MFI"`
	MF_TEST_BATTERIES []MFN_M10_MF_TEST_BATTERIES `hl7:"true,MF_TEST_BATTERIES"`
	Other             []interface{}
}

func (s *MFN_M10) MessageTypeName() string {
	return "MFN_M10"
}

// MFN_M10_MF_TEST_BATTERIES represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M10_MF_TEST_BATTERIES struct {
	MFE                 *MFE                         `hl7:"true,MFE"`
	OM1                 *OM1                         `hl7:"true,OM1"`
	MF_TEST_BATT_DETAIL *MFN_M10_MF_TEST_BATT_DETAIL `hl7:"false,MF_TEST_BATT_DETAIL"`
	Other               []interface{}
}

func (s *MFN_M10_MF_TEST_BATTERIES) MessageTypeSubStructName() string {
	return "MFN_M10_MF_TEST_BATTERIES"
}

// MFN_M10_MF_TEST_BATT_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M10_MF_TEST_BATT_DETAIL struct {
	OM5   *OM5  `hl7:"true,OM5"`
	OM4   []OM4 `hl7:"false,OM4"`
	Other []interface{}
}

func (s *MFN_M10_MF_TEST_BATT_DETAIL) MessageTypeSubStructName() string {
	return "MFN_M10_MF_TEST_BATT_DETAIL"
}

// MFN_M11 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M11 struct {
	MSH                *MSH                         `hl7:"true,MSH"`
	MFI                *MFI                         `hl7:"true,MFI"`
	MF_TEST_CALCULATED []MFN_M11_MF_TEST_CALCULATED `hl7:"true,MF_TEST_CALCULATED"`
	Other              []interface{}
}

func (s *MFN_M11) MessageTypeName() string {
	return "MFN_M11"
}

// MFN_M11_MF_TEST_CALCULATED represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M11_MF_TEST_CALCULATED struct {
	MFE                 *MFE                         `hl7:"true,MFE"`
	OM1                 *OM1                         `hl7:"true,OM1"`
	MF_TEST_CALC_DETAIL *MFN_M11_MF_TEST_CALC_DETAIL `hl7:"false,MF_TEST_CALC_DETAIL"`
	Other               []interface{}
}

func (s *MFN_M11_MF_TEST_CALCULATED) MessageTypeSubStructName() string {
	return "MFN_M11_MF_TEST_CALCULATED"
}

// MFN_M11_MF_TEST_CALC_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M11_MF_TEST_CALC_DETAIL struct {
	OM6   *OM6 `hl7:"true,OM6"`
	OM2   *OM2 `hl7:"true,OM2"`
	Other []interface{}
}

func (s *MFN_M11_MF_TEST_CALC_DETAIL) MessageTypeSubStructName() string {
	return "MFN_M11_MF_TEST_CALC_DETAIL"
}

// MFN_M12 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M12 struct {
	MSH               *MSH                        `hl7:"true,MSH"`
	MFI               *MFI                        `hl7:"true,MFI"`
	MF_OBS_ATTRIBUTES []MFN_M12_MF_OBS_ATTRIBUTES `hl7:"true,MF_OBS_ATTRIBUTES"`
	Other             []interface{}
}

func (s *MFN_M12) MessageTypeName() string {
	return "MFN_M12"
}

// MFN_M12_MF_OBS_ATTRIBUTES represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFN_M12_MF_OBS_ATTRIBUTES struct {
	MFE   *MFE `hl7:"true,MFE"`
	OM1   *OM1 `hl7:"true,OM1"`
	OM7   *OM7 `hl7:"false,OM7"`
	Other []interface{}
}

func (s *MFN_M12_MF_OBS_ATTRIBUTES) MessageTypeSubStructName() string {
	return "MFN_M12_MF_OBS_ATTRIBUTES"
}

// MFQ_M01 represents the corresponding HL7 message type.
// Definition from HL7 2.4
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
// Definition from HL7 2.4
type MFR_M01 struct {
	MSH      *MSH               `hl7:"true,MSH"`
	MSA      *MSA               `hl7:"true,MSA"`
	ERR      *ERR               `hl7:"false,ERR"`
	QAK      *QAK               `hl7:"false,QAK"`
	QRD      *QRD               `hl7:"true,QRD"`
	QRF      *QRF               `hl7:"false,QRF"`
	MFI      *MFI               `hl7:"true,MFI"`
	MF_QUERY []MFR_M01_MF_QUERY `hl7:"true,MF_QUERY"`
	DSC      *DSC               `hl7:"false,DSC"`
	Other    []interface{}
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

// MFR_M01_MF_QUERY represents the corresponding HL7 message type.
// Definition from HL7 2.4
type MFR_M01_MF_QUERY struct {
	MFE *MFE `hl7:"true,MFE"`
	// Missing: anyZSegment
	Other []interface{}
}

func (s *MFR_M01_MF_QUERY) MessageTypeSubStructName() string {
	return "MFR_M01_MF_QUERY"
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

// NMD_N02_APP_STATS represents the corresponding HL7 message type.
// Definition from HL7 2.4
type NMD_N02_APP_STATS struct {
	NST   *NST  `hl7:"true,NST"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *NMD_N02_APP_STATS) MessageTypeSubStructName() string {
	return "NMD_N02_APP_STATS"
}

// NMD_N02_APP_STATUS represents the corresponding HL7 message type.
// Definition from HL7 2.4
type NMD_N02_APP_STATUS struct {
	NSC   *NSC  `hl7:"true,NSC"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *NMD_N02_APP_STATUS) MessageTypeSubStructName() string {
	return "NMD_N02_APP_STATUS"
}

// NMD_N02_CLOCK represents the corresponding HL7 message type.
// Definition from HL7 2.4
type NMD_N02_CLOCK struct {
	NCK   *NCK  `hl7:"true,NCK"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *NMD_N02_CLOCK) MessageTypeSubStructName() string {
	return "NMD_N02_CLOCK"
}

// NMD_N02_CLOCK_AND_STATS_WITH_NOTES represents the corresponding HL7 message type.
// Definition from HL7 2.4
type NMD_N02_CLOCK_AND_STATS_WITH_NOTES struct {
	CLOCK      *NMD_N02_CLOCK      `hl7:"false,CLOCK"`
	APP_STATS  *NMD_N02_APP_STATS  `hl7:"false,APP_STATS"`
	APP_STATUS *NMD_N02_APP_STATUS `hl7:"false,APP_STATUS"`
	Other      []interface{}
}

func (s *NMD_N02_CLOCK_AND_STATS_WITH_NOTES) MessageTypeSubStructName() string {
	return "NMD_N02_CLOCK_AND_STATS_WITH_NOTES"
}

// NMD_N02 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type NMD_N02 struct {
	MSH                        *MSH                                 `hl7:"true,MSH"`
	CLOCK_AND_STATS_WITH_NOTES []NMD_N02_CLOCK_AND_STATS_WITH_NOTES `hl7:"true,CLOCK_AND_STATS_WITH_NOTES"`
	Other                      []interface{}
}

func (s *NMD_N02) MessageTypeName() string {
	return "NMD_N02"
}

// NMQ_N01_CLOCK_AND_STATISTICS represents the corresponding HL7 message type.
// Definition from HL7 2.4
type NMQ_N01_CLOCK_AND_STATISTICS struct {
	NCK   *NCK `hl7:"false,NCK"`
	NST   *NST `hl7:"false,NST"`
	NSC   *NSC `hl7:"false,NSC"`
	Other []interface{}
}

func (s *NMQ_N01_CLOCK_AND_STATISTICS) MessageTypeSubStructName() string {
	return "NMQ_N01_CLOCK_AND_STATISTICS"
}

// NMQ_N01 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type NMQ_N01 struct {
	MSH                  *MSH                           `hl7:"true,MSH"`
	QRY_WITH_DETAIL      *NMQ_N01_QRY_WITH_DETAIL       `hl7:"false,QRY_WITH_DETAIL"`
	CLOCK_AND_STATISTICS []NMQ_N01_CLOCK_AND_STATISTICS `hl7:"true,CLOCK_AND_STATISTICS"`
	Other                []interface{}
}

func (s *NMQ_N01) MessageTypeName() string {
	return "NMQ_N01"
}

// NMQ_N01_QRY_WITH_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type NMQ_N01_QRY_WITH_DETAIL struct {
	QRD   *QRD `hl7:"true,QRD"`
	QRF   *QRF `hl7:"false,QRF"`
	Other []interface{}
}

func (s *NMQ_N01_QRY_WITH_DETAIL) MessageTypeSubStructName() string {
	return "NMQ_N01_QRY_WITH_DETAIL"
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

// NMR_N01_CLOCK_AND_STATS_WITH_NOTES_ALT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type NMR_N01_CLOCK_AND_STATS_WITH_NOTES_ALT struct {
	NCK   *NCK  `hl7:"false,NCK"`
	NTE1  []NTE `hl7:"false,NTE1"`
	NST   *NST  `hl7:"false,NST"`
	NTE2  []NTE `hl7:"false,NTE2"`
	NSC   *NSC  `hl7:"false,NSC"`
	NTE3  []NTE `hl7:"false,NTE3"`
	Other []interface{}
}

func (s *NMR_N01_CLOCK_AND_STATS_WITH_NOTES_ALT) MessageTypeSubStructName() string {
	return "NMR_N01_CLOCK_AND_STATS_WITH_NOTES_ALT"
}

// NMR_N01 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type NMR_N01 struct {
	MSH                            *MSH                                     `hl7:"true,MSH"`
	MSA                            *MSA                                     `hl7:"true,MSA"`
	ERR                            *ERR                                     `hl7:"false,ERR"`
	QRD                            *QRD                                     `hl7:"false,QRD"`
	CLOCK_AND_STATS_WITH_NOTES_ALT []NMR_N01_CLOCK_AND_STATS_WITH_NOTES_ALT `hl7:"true,CLOCK_AND_STATS_WITH_NOTES_ALT"`
	Other                          []interface{}
}

func (s *NMR_N01) MessageTypeName() string {
	return "NMR_N01"
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

// OMD_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMD_O01 struct {
	MSH        *MSH                 `hl7:"true,MSH"`
	NTE        []NTE                `hl7:"false,NTE"`
	PATIENT    *OMD_O01_PATIENT     `hl7:"false,PATIENT"`
	ORDER_DIET []OMD_O01_ORDER_DIET `hl7:"true,ORDER_DIET"`
	ORDER_TRAY []OMD_O01_ORDER_TRAY `hl7:"false,ORDER_TRAY"`
	Other      []interface{}
}

func (s *OMD_O01) MessageTypeName() string {
	return "OMD_O01"
}

// OMD_O01_DIET represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMD_O01_DIET struct {
	ODS         []ODS                 `hl7:"true,ODS"`
	NTE         []NTE                 `hl7:"false,NTE"`
	OBSERVATION []OMD_O01_OBSERVATION `hl7:"true,OBSERVATION"`
	Other       []interface{}
}

func (s *OMD_O01_DIET) MessageTypeSubStructName() string {
	return "OMD_O01_DIET"
}

// OMD_O01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMD_O01_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *OMD_O01_INSURANCE) MessageTypeSubStructName() string {
	return "OMD_O01_INSURANCE"
}

// OMD_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMD_O01_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OMD_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "OMD_O01_OBSERVATION"
}

// OMD_O01_ORDER_DIET represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMD_O01_ORDER_DIET struct {
	ORC   *ORC          `hl7:"true,ORC"`
	DIET  *OMD_O01_DIET `hl7:"false,DIET"`
	Other []interface{}
}

func (s *OMD_O01_ORDER_DIET) MessageTypeSubStructName() string {
	return "OMD_O01_ORDER_DIET"
}

// OMD_O01_ORDER_TRAY represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMD_O01_ORDER_TRAY struct {
	ORC   *ORC  `hl7:"true,ORC"`
	ODT   []ODT `hl7:"true,ODT"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OMD_O01_ORDER_TRAY) MessageTypeSubStructName() string {
	return "OMD_O01_ORDER_TRAY"
}

// OMD_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMD_O01_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NTE           []NTE                  `hl7:"false,NTE"`
	PATIENT_VISIT *OMD_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE     []OMD_O01_INSURANCE    `hl7:"false,INSURANCE"`
	GT1           *GT1                   `hl7:"false,GT1"`
	AL1           []AL1                  `hl7:"false,AL1"`
	Other         []interface{}
}

func (s *OMD_O01_PATIENT) MessageTypeSubStructName() string {
	return "OMD_O01_PATIENT"
}

// OMD_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMD_O01_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *OMD_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "OMD_O01_PATIENT_VISIT"
}

// OMD_O03 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMD_O03 struct {
	MSH        *MSH                 `hl7:"true,MSH"`
	NTE        []NTE                `hl7:"false,NTE"`
	PATIENT    *OMD_O03_PATIENT     `hl7:"false,PATIENT"`
	ORDER_DIET []OMD_O03_ORDER_DIET `hl7:"true,ORDER_DIET"`
	ORDER_TRAY []OMD_O03_ORDER_TRAY `hl7:"false,ORDER_TRAY"`
	Other      []interface{}
}

func (s *OMD_O03) MessageTypeName() string {
	return "OMD_O03"
}

// OMD_O03_DIET represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMD_O03_DIET struct {
	ODS         []ODS                 `hl7:"true,ODS"`
	NTE         []NTE                 `hl7:"false,NTE"`
	OBSERVATION []OMD_O03_OBSERVATION `hl7:"false,OBSERVATION"`
	Other       []interface{}
}

func (s *OMD_O03_DIET) MessageTypeSubStructName() string {
	return "OMD_O03_DIET"
}

// OMD_O03_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMD_O03_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *OMD_O03_INSURANCE) MessageTypeSubStructName() string {
	return "OMD_O03_INSURANCE"
}

// OMD_O03_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMD_O03_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OMD_O03_OBSERVATION) MessageTypeSubStructName() string {
	return "OMD_O03_OBSERVATION"
}

// OMD_O03_ORDER_DIET represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMD_O03_ORDER_DIET struct {
	ORC   *ORC          `hl7:"true,ORC"`
	DIET  *OMD_O03_DIET `hl7:"false,DIET"`
	Other []interface{}
}

func (s *OMD_O03_ORDER_DIET) MessageTypeSubStructName() string {
	return "OMD_O03_ORDER_DIET"
}

// OMD_O03_ORDER_TRAY represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMD_O03_ORDER_TRAY struct {
	ORC   *ORC  `hl7:"true,ORC"`
	ODT   []ODT `hl7:"true,ODT"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OMD_O03_ORDER_TRAY) MessageTypeSubStructName() string {
	return "OMD_O03_ORDER_TRAY"
}

// OMD_O03_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMD_O03_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NTE           []NTE                  `hl7:"false,NTE"`
	PATIENT_VISIT *OMD_O03_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE     []OMD_O03_INSURANCE    `hl7:"false,INSURANCE"`
	GT1           *GT1                   `hl7:"false,GT1"`
	AL1           []AL1                  `hl7:"false,AL1"`
	Other         []interface{}
}

func (s *OMD_O03_PATIENT) MessageTypeSubStructName() string {
	return "OMD_O03_PATIENT"
}

// OMD_O03_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMD_O03_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *OMD_O03_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "OMD_O03_PATIENT_VISIT"
}

// OMG_O19 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMG_O19 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	NTE     []NTE            `hl7:"false,NTE"`
	PATIENT *OMG_O19_PATIENT `hl7:"false,PATIENT"`
	ORDER   []OMG_O19_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *OMG_O19) MessageTypeName() string {
	return "OMG_O19"
}

// OMG_O19_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMG_O19_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *OMG_O19_INSURANCE) MessageTypeSubStructName() string {
	return "OMG_O19_INSURANCE"
}

// OMG_O19_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMG_O19_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OMG_O19_OBSERVATION) MessageTypeSubStructName() string {
	return "OMG_O19_OBSERVATION"
}

// OMG_O19_OBSERVATION_PRIOR represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMG_O19_OBSERVATION_PRIOR struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OMG_O19_OBSERVATION_PRIOR) MessageTypeSubStructName() string {
	return "OMG_O19_OBSERVATION_PRIOR"
}

// OMG_O19_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMG_O19_ORDER struct {
	ORC          *ORC                   `hl7:"true,ORC"`
	OBR          *OBR                   `hl7:"true,OBR"`
	NTE          []NTE                  `hl7:"false,NTE"`
	CTD          *CTD                   `hl7:"false,CTD"`
	DG1          []DG1                  `hl7:"false,DG1"`
	OBSERVATION  []OMG_O19_OBSERVATION  `hl7:"false,OBSERVATION"`
	PRIOR_RESULT []OMG_O19_PRIOR_RESULT `hl7:"false,PRIOR_RESULT"`
	FT1          []FT1                  `hl7:"false,FT1"`
	CTI          []CTI                  `hl7:"false,CTI"`
	BLG          *BLG                   `hl7:"false,BLG"`
	Other        []interface{}
}

func (s *OMG_O19_ORDER) MessageTypeSubStructName() string {
	return "OMG_O19_ORDER"
}

// OMG_O19_ORDER_PRIOR represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMG_O19_ORDER_PRIOR struct {
	ORC               *ORC                        `hl7:"false,ORC"`
	OBR               *OBR                        `hl7:"true,OBR"`
	NTE               []NTE                       `hl7:"false,NTE"`
	CTD               *CTD                        `hl7:"false,CTD"`
	OBSERVATION_PRIOR []OMG_O19_OBSERVATION_PRIOR `hl7:"true,OBSERVATION_PRIOR"`
	Other             []interface{}
}

func (s *OMG_O19_ORDER_PRIOR) MessageTypeSubStructName() string {
	return "OMG_O19_ORDER_PRIOR"
}

// OMG_O19_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMG_O19_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NTE           []NTE                  `hl7:"false,NTE"`
	PATIENT_VISIT *OMG_O19_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE     []OMG_O19_INSURANCE    `hl7:"false,INSURANCE"`
	GT1           *GT1                   `hl7:"false,GT1"`
	AL1           []AL1                  `hl7:"false,AL1"`
	Other         []interface{}
}

func (s *OMG_O19_PATIENT) MessageTypeSubStructName() string {
	return "OMG_O19_PATIENT"
}

// OMG_O19_PATIENT_PRIOR represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMG_O19_PATIENT_PRIOR struct {
	PID   *PID `hl7:"true,PID"`
	PD1   *PD1 `hl7:"false,PD1"`
	Other []interface{}
}

func (s *OMG_O19_PATIENT_PRIOR) MessageTypeSubStructName() string {
	return "OMG_O19_PATIENT_PRIOR"
}

// OMG_O19_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMG_O19_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *OMG_O19_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "OMG_O19_PATIENT_VISIT"
}

// OMG_O19_PATIENT_VISIT_PRIOR represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMG_O19_PATIENT_VISIT_PRIOR struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *OMG_O19_PATIENT_VISIT_PRIOR) MessageTypeSubStructName() string {
	return "OMG_O19_PATIENT_VISIT_PRIOR"
}

// OMG_O19_PRIOR_RESULT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMG_O19_PRIOR_RESULT struct {
	PATIENT_PRIOR       *OMG_O19_PATIENT_PRIOR       `hl7:"false,PATIENT_PRIOR"`
	PATIENT_VISIT_PRIOR *OMG_O19_PATIENT_VISIT_PRIOR `hl7:"false,PATIENT_VISIT_PRIOR"`
	AL1                 []AL1                        `hl7:"false,AL1"`
	ORDER_PRIOR         []OMG_O19_ORDER_PRIOR        `hl7:"true,ORDER_PRIOR"`
	Other               []interface{}
}

func (s *OMG_O19_PRIOR_RESULT) MessageTypeSubStructName() string {
	return "OMG_O19_PRIOR_RESULT"
}

// OML_O21_CONTAINER_1 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OML_O21_CONTAINER_1 struct {
	SAC   *SAC  `hl7:"true,SAC"`
	OBX   []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *OML_O21_CONTAINER_1) MessageTypeSubStructName() string {
	return "OML_O21_CONTAINER_1"
}

// OML_O21_CONTAINER_2 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OML_O21_CONTAINER_2 struct {
	SAC   *SAC  `hl7:"true,SAC"`
	OBX   []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *OML_O21_CONTAINER_2) MessageTypeSubStructName() string {
	return "OML_O21_CONTAINER_2"
}

// OML_O21 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OML_O21 struct {
	MSH           *MSH                    `hl7:"true,MSH"`
	NTE           []NTE                   `hl7:"false,NTE"`
	PATIENT       *OML_O21_PATIENT        `hl7:"false,PATIENT"`
	ORDER_GENERAL []OML_O21_ORDER_GENERAL `hl7:"true,ORDER_GENERAL"`
	Other         []interface{}
}

func (s *OML_O21) MessageTypeName() string {
	return "OML_O21"
}

// OML_O21_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OML_O21_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *OML_O21_INSURANCE) MessageTypeSubStructName() string {
	return "OML_O21_INSURANCE"
}

// OML_O21_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OML_O21_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	TCD   *TCD  `hl7:"false,TCD"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OML_O21_OBSERVATION) MessageTypeSubStructName() string {
	return "OML_O21_OBSERVATION"
}

// OML_O21_OBSERVATION_PRIOR represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OML_O21_OBSERVATION_PRIOR struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OML_O21_OBSERVATION_PRIOR) MessageTypeSubStructName() string {
	return "OML_O21_OBSERVATION_PRIOR"
}

// OML_O21_OBSERVATION_REQUEST represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OML_O21_OBSERVATION_REQUEST struct {
	OBR          *OBR                   `hl7:"true,OBR"`
	CONTAINER_2  []OML_O21_CONTAINER_2  `hl7:"false,CONTAINER_2"`
	TCD          *TCD                   `hl7:"false,TCD"`
	NTE          []NTE                  `hl7:"false,NTE"`
	DG1          []DG1                  `hl7:"false,DG1"`
	OBSERVATION  []OML_O21_OBSERVATION  `hl7:"false,OBSERVATION"`
	PRIOR_RESULT []OML_O21_PRIOR_RESULT `hl7:"false,PRIOR_RESULT"`
	Other        []interface{}
}

func (s *OML_O21_OBSERVATION_REQUEST) MessageTypeSubStructName() string {
	return "OML_O21_OBSERVATION_REQUEST"
}

// OML_O21_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OML_O21_ORDER struct {
	ORC                 *ORC                         `hl7:"true,ORC"`
	OBSERVATION_REQUEST *OML_O21_OBSERVATION_REQUEST `hl7:"false,OBSERVATION_REQUEST"`
	FT1                 []FT1                        `hl7:"false,FT1"`
	CTI                 []CTI                        `hl7:"false,CTI"`
	BLG                 *BLG                         `hl7:"false,BLG"`
	Other               []interface{}
}

func (s *OML_O21_ORDER) MessageTypeSubStructName() string {
	return "OML_O21_ORDER"
}

// OML_O21_ORDER_GENERAL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OML_O21_ORDER_GENERAL struct {
	CONTAINER_1 *OML_O21_CONTAINER_1 `hl7:"false,CONTAINER_1"`
	ORDER       []OML_O21_ORDER      `hl7:"true,ORDER"`
	Other       []interface{}
}

func (s *OML_O21_ORDER_GENERAL) MessageTypeSubStructName() string {
	return "OML_O21_ORDER_GENERAL"
}

// OML_O21_ORDER_PRIOR represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OML_O21_ORDER_PRIOR struct {
	ORC               *ORC                        `hl7:"false,ORC"`
	OBR               *OBR                        `hl7:"true,OBR"`
	NTE               []NTE                       `hl7:"false,NTE"`
	OBSERVATION_PRIOR []OML_O21_OBSERVATION_PRIOR `hl7:"true,OBSERVATION_PRIOR"`
	Other             []interface{}
}

func (s *OML_O21_ORDER_PRIOR) MessageTypeSubStructName() string {
	return "OML_O21_ORDER_PRIOR"
}

// OML_O21_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OML_O21_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NTE           []NTE                  `hl7:"false,NTE"`
	PATIENT_VISIT *OML_O21_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE     []OML_O21_INSURANCE    `hl7:"false,INSURANCE"`
	GT1           *GT1                   `hl7:"false,GT1"`
	AL1           []AL1                  `hl7:"false,AL1"`
	Other         []interface{}
}

func (s *OML_O21_PATIENT) MessageTypeSubStructName() string {
	return "OML_O21_PATIENT"
}

// OML_O21_PATIENT_PRIOR represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OML_O21_PATIENT_PRIOR struct {
	PID   *PID `hl7:"true,PID"`
	PD1   *PD1 `hl7:"false,PD1"`
	Other []interface{}
}

func (s *OML_O21_PATIENT_PRIOR) MessageTypeSubStructName() string {
	return "OML_O21_PATIENT_PRIOR"
}

// OML_O21_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OML_O21_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *OML_O21_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "OML_O21_PATIENT_VISIT"
}

// OML_O21_PATIENT_VISIT_PRIOR represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OML_O21_PATIENT_VISIT_PRIOR struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *OML_O21_PATIENT_VISIT_PRIOR) MessageTypeSubStructName() string {
	return "OML_O21_PATIENT_VISIT_PRIOR"
}

// OML_O21_PRIOR_RESULT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OML_O21_PRIOR_RESULT struct {
	PATIENT_PRIOR       *OML_O21_PATIENT_PRIOR       `hl7:"false,PATIENT_PRIOR"`
	PATIENT_VISIT_PRIOR *OML_O21_PATIENT_VISIT_PRIOR `hl7:"false,PATIENT_VISIT_PRIOR"`
	AL1                 []AL1                        `hl7:"false,AL1"`
	ORDER_PRIOR         []OML_O21_ORDER_PRIOR        `hl7:"true,ORDER_PRIOR"`
	Other               []interface{}
}

func (s *OML_O21_PRIOR_RESULT) MessageTypeSubStructName() string {
	return "OML_O21_PRIOR_RESULT"
}

// OMN_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMN_O01 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	NTE     []NTE            `hl7:"false,NTE"`
	PATIENT *OMN_O01_PATIENT `hl7:"false,PATIENT"`
	ORDER   []OMN_O01_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *OMN_O01) MessageTypeName() string {
	return "OMN_O01"
}

// OMN_O01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMN_O01_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *OMN_O01_INSURANCE) MessageTypeSubStructName() string {
	return "OMN_O01_INSURANCE"
}

// OMN_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMN_O01_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OMN_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "OMN_O01_OBSERVATION"
}

// OMN_O01_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMN_O01_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *OMN_O01_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	BLG          *BLG                  `hl7:"false,BLG"`
	Other        []interface{}
}

func (s *OMN_O01_ORDER) MessageTypeSubStructName() string {
	return "OMN_O01_ORDER"
}

// OMN_O01_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMN_O01_ORDER_DETAIL struct {
	RQD         *RQD                  `hl7:"true,RQD"`
	RQ1         *RQ1                  `hl7:"false,RQ1"`
	NTE         []NTE                 `hl7:"false,NTE"`
	OBSERVATION []OMN_O01_OBSERVATION `hl7:"false,OBSERVATION"`
	Other       []interface{}
}

func (s *OMN_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "OMN_O01_ORDER_DETAIL"
}

// OMN_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMN_O01_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NTE           []NTE                  `hl7:"false,NTE"`
	PATIENT_VISIT *OMN_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE     []OMN_O01_INSURANCE    `hl7:"false,INSURANCE"`
	GT1           *GT1                   `hl7:"false,GT1"`
	AL1           []AL1                  `hl7:"false,AL1"`
	Other         []interface{}
}

func (s *OMN_O01_PATIENT) MessageTypeSubStructName() string {
	return "OMN_O01_PATIENT"
}

// OMN_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMN_O01_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *OMN_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "OMN_O01_PATIENT_VISIT"
}

// OMN_O07 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMN_O07 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	NTE     []NTE            `hl7:"false,NTE"`
	PATIENT *OMN_O07_PATIENT `hl7:"false,PATIENT"`
	ORDER   []OMN_O07_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *OMN_O07) MessageTypeName() string {
	return "OMN_O07"
}

// OMN_O07_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMN_O07_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *OMN_O07_INSURANCE) MessageTypeSubStructName() string {
	return "OMN_O07_INSURANCE"
}

// OMN_O07_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMN_O07_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OMN_O07_OBSERVATION) MessageTypeSubStructName() string {
	return "OMN_O07_OBSERVATION"
}

// OMN_O07_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMN_O07_ORDER struct {
	ORC         *ORC                  `hl7:"true,ORC"`
	RQD         *RQD                  `hl7:"true,RQD"`
	RQ1         *RQ1                  `hl7:"false,RQ1"`
	NTE         []NTE                 `hl7:"false,NTE"`
	OBSERVATION []OMN_O07_OBSERVATION `hl7:"false,OBSERVATION"`
	BLG         *BLG                  `hl7:"false,BLG"`
	Other       []interface{}
}

func (s *OMN_O07_ORDER) MessageTypeSubStructName() string {
	return "OMN_O07_ORDER"
}

// OMN_O07_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMN_O07_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NTE           []NTE                  `hl7:"false,NTE"`
	PATIENT_VISIT *OMN_O07_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE     []OMN_O07_INSURANCE    `hl7:"false,INSURANCE"`
	GT1           *GT1                   `hl7:"false,GT1"`
	AL1           []AL1                  `hl7:"false,AL1"`
	Other         []interface{}
}

func (s *OMN_O07_PATIENT) MessageTypeSubStructName() string {
	return "OMN_O07_PATIENT"
}

// OMN_O07_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMN_O07_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *OMN_O07_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "OMN_O07_PATIENT_VISIT"
}

// OMP_O09_COMPONENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMP_O09_COMPONENT struct {
	RXC   []RXC `hl7:"true,RXC"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OMP_O09_COMPONENT) MessageTypeSubStructName() string {
	return "OMP_O09_COMPONENT"
}

// OMP_O09 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMP_O09 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	NTE     []NTE            `hl7:"false,NTE"`
	PATIENT *OMP_O09_PATIENT `hl7:"false,PATIENT"`
	ORDER   []OMP_O09_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *OMP_O09) MessageTypeName() string {
	return "OMP_O09"
}

// OMP_O09_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMP_O09_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *OMP_O09_INSURANCE) MessageTypeSubStructName() string {
	return "OMP_O09_INSURANCE"
}

// OMP_O09_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMP_O09_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OMP_O09_OBSERVATION) MessageTypeSubStructName() string {
	return "OMP_O09_OBSERVATION"
}

// OMP_O09_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMP_O09_ORDER struct {
	ORC         *ORC                  `hl7:"true,ORC"`
	RXO         *RXO                  `hl7:"true,RXO"`
	NTE         []NTE                 `hl7:"false,NTE"`
	RXR         []RXR                 `hl7:"true,RXR"`
	COMPONENT   *OMP_O09_COMPONENT    `hl7:"false,COMPONENT"`
	OBSERVATION []OMP_O09_OBSERVATION `hl7:"false,OBSERVATION"`
	FT1         []FT1                 `hl7:"false,FT1"`
	BLG         *BLG                  `hl7:"false,BLG"`
	Other       []interface{}
}

func (s *OMP_O09_ORDER) MessageTypeSubStructName() string {
	return "OMP_O09_ORDER"
}

// OMP_O09_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMP_O09_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NTE           []NTE                  `hl7:"false,NTE"`
	PATIENT_VISIT *OMP_O09_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE     []OMP_O09_INSURANCE    `hl7:"false,INSURANCE"`
	GT1           *GT1                   `hl7:"false,GT1"`
	AL1           []AL1                  `hl7:"false,AL1"`
	Other         []interface{}
}

func (s *OMP_O09_PATIENT) MessageTypeSubStructName() string {
	return "OMP_O09_PATIENT"
}

// OMP_O09_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMP_O09_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *OMP_O09_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "OMP_O09_PATIENT_VISIT"
}

// OMS_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMS_O01 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	NTE     []NTE            `hl7:"false,NTE"`
	PATIENT *OMS_O01_PATIENT `hl7:"false,PATIENT"`
	ORDER   []OMS_O01_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *OMS_O01) MessageTypeName() string {
	return "OMS_O01"
}

// OMS_O01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMS_O01_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *OMS_O01_INSURANCE) MessageTypeSubStructName() string {
	return "OMS_O01_INSURANCE"
}

// OMS_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMS_O01_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OMS_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "OMS_O01_OBSERVATION"
}

// OMS_O01_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMS_O01_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *OMS_O01_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	BLG          *BLG                  `hl7:"false,BLG"`
	Other        []interface{}
}

func (s *OMS_O01_ORDER) MessageTypeSubStructName() string {
	return "OMS_O01_ORDER"
}

// OMS_O01_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMS_O01_ORDER_DETAIL struct {
	RQD         *RQD                  `hl7:"true,RQD"`
	NTE         []NTE                 `hl7:"false,NTE"`
	OBSERVATION []OMS_O01_OBSERVATION `hl7:"false,OBSERVATION"`
	Other       []interface{}
}

func (s *OMS_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "OMS_O01_ORDER_DETAIL"
}

// OMS_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMS_O01_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NTE           []NTE                  `hl7:"false,NTE"`
	PATIENT_VISIT *OMS_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE     []OMS_O01_INSURANCE    `hl7:"false,INSURANCE"`
	GT1           *GT1                   `hl7:"false,GT1"`
	AL1           []AL1                  `hl7:"false,AL1"`
	Other         []interface{}
}

func (s *OMS_O01_PATIENT) MessageTypeSubStructName() string {
	return "OMS_O01_PATIENT"
}

// OMS_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OMS_O01_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *OMS_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "OMS_O01_PATIENT_VISIT"
}

// OMS_O05 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMS_O05 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	NTE     []NTE            `hl7:"false,NTE"`
	PATIENT *OMS_O05_PATIENT `hl7:"false,PATIENT"`
	ORDER   []OMS_O05_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *OMS_O05) MessageTypeName() string {
	return "OMS_O05"
}

// OMS_O05_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMS_O05_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *OMS_O05_INSURANCE) MessageTypeSubStructName() string {
	return "OMS_O05_INSURANCE"
}

// OMS_O05_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMS_O05_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OMS_O05_OBSERVATION) MessageTypeSubStructName() string {
	return "OMS_O05_OBSERVATION"
}

// OMS_O05_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMS_O05_ORDER struct {
	ORC         *ORC                  `hl7:"true,ORC"`
	RQD         *RQD                  `hl7:"true,RQD"`
	RQ1         *RQ1                  `hl7:"false,RQ1"`
	NTE         []NTE                 `hl7:"false,NTE"`
	OBSERVATION []OMS_O05_OBSERVATION `hl7:"false,OBSERVATION"`
	BLG         *BLG                  `hl7:"false,BLG"`
	Other       []interface{}
}

func (s *OMS_O05_ORDER) MessageTypeSubStructName() string {
	return "OMS_O05_ORDER"
}

// OMS_O05_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMS_O05_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NTE           []NTE                  `hl7:"false,NTE"`
	PATIENT_VISIT *OMS_O05_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE     []OMS_O05_INSURANCE    `hl7:"false,INSURANCE"`
	GT1           *GT1                   `hl7:"false,GT1"`
	AL1           []AL1                  `hl7:"false,AL1"`
	Other         []interface{}
}

func (s *OMS_O05_PATIENT) MessageTypeSubStructName() string {
	return "OMS_O05_PATIENT"
}

// OMS_O05_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OMS_O05_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *OMS_O05_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "OMS_O05_PATIENT_VISIT"
}

// ORD_O02 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORD_O02 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *ORD_O02_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *ORD_O02) MessageTypeName() string {
	return "ORD_O02"
}

// ORD_O02_ORDER_DIET represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORD_O02_ORDER_DIET struct {
	ORC   *ORC  `hl7:"true,ORC"`
	ODS   []ODS `hl7:"false,ODS"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORD_O02_ORDER_DIET) MessageTypeSubStructName() string {
	return "ORD_O02_ORDER_DIET"
}

// ORD_O02_ORDER_TRAY represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORD_O02_ORDER_TRAY struct {
	ORC   *ORC  `hl7:"true,ORC"`
	ODT   []ODT `hl7:"false,ODT"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORD_O02_ORDER_TRAY) MessageTypeSubStructName() string {
	return "ORD_O02_ORDER_TRAY"
}

// ORD_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORD_O02_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORD_O02_PATIENT) MessageTypeSubStructName() string {
	return "ORD_O02_PATIENT"
}

// ORD_O02_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORD_O02_RESPONSE struct {
	PATIENT    *ORD_O02_PATIENT     `hl7:"false,PATIENT"`
	ORDER_DIET []ORD_O02_ORDER_DIET `hl7:"true,ORDER_DIET"`
	ORDER_TRAY []ORD_O02_ORDER_TRAY `hl7:"false,ORDER_TRAY"`
	Other      []interface{}
}

func (s *ORD_O02_RESPONSE) MessageTypeSubStructName() string {
	return "ORD_O02_RESPONSE"
}

// ORD_O04 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORD_O04 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *ORD_O04_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *ORD_O04) MessageTypeName() string {
	return "ORD_O04"
}

// ORD_O04_ORDER_DIET represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORD_O04_ORDER_DIET struct {
	ORC   *ORC  `hl7:"true,ORC"`
	ODS   []ODS `hl7:"false,ODS"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORD_O04_ORDER_DIET) MessageTypeSubStructName() string {
	return "ORD_O04_ORDER_DIET"
}

// ORD_O04_ORDER_TRAY represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORD_O04_ORDER_TRAY struct {
	ORC   *ORC  `hl7:"true,ORC"`
	ODT   []ODT `hl7:"false,ODT"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORD_O04_ORDER_TRAY) MessageTypeSubStructName() string {
	return "ORD_O04_ORDER_TRAY"
}

// ORD_O04_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORD_O04_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORD_O04_PATIENT) MessageTypeSubStructName() string {
	return "ORD_O04_PATIENT"
}

// ORD_O04_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORD_O04_RESPONSE struct {
	PATIENT    *ORD_O04_PATIENT     `hl7:"false,PATIENT"`
	ORDER_DIET []ORD_O04_ORDER_DIET `hl7:"true,ORDER_DIET"`
	ORDER_TRAY []ORD_O04_ORDER_TRAY `hl7:"false,ORDER_TRAY"`
	Other      []interface{}
}

func (s *ORD_O04_RESPONSE) MessageTypeSubStructName() string {
	return "ORD_O04_RESPONSE"
}

// ORF_R04 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORF_R04 struct {
	MSH      *MSH               `hl7:"true,MSH"`
	MSA      *MSA               `hl7:"true,MSA"`
	QRD      *QRD               `hl7:"true,QRD"`
	QRF      *QRF               `hl7:"false,QRF"`
	RESPONSE []ORF_R04_RESPONSE `hl7:"true,RESPONSE"`
	ERR      *ERR               `hl7:"false,ERR"`
	QAK      *QAK               `hl7:"false,QAK"`
	DSC      *DSC               `hl7:"false,DSC"`
	Other    []interface{}
}

func (s *ORF_R04) MessageTypeName() string {
	return "ORF_R04"
}

// ORF_R04_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORF_R04_OBSERVATION struct {
	OBX   *OBX  `hl7:"false,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORF_R04_OBSERVATION) MessageTypeSubStructName() string {
	return "ORF_R04_OBSERVATION"
}

// ORF_R04_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORF_R04_ORDER struct {
	ORC         *ORC                  `hl7:"false,ORC"`
	OBR         *OBR                  `hl7:"true,OBR"`
	NTE         []NTE                 `hl7:"false,NTE"`
	CTD         *CTD                  `hl7:"false,CTD"`
	OBSERVATION []ORF_R04_OBSERVATION `hl7:"true,OBSERVATION"`
	CTI         []CTI                 `hl7:"false,CTI"`
	Other       []interface{}
}

func (s *ORF_R04_ORDER) MessageTypeSubStructName() string {
	return "ORF_R04_ORDER"
}

// ORF_R04_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORF_R04_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORF_R04_PATIENT) MessageTypeSubStructName() string {
	return "ORF_R04_PATIENT"
}

// ORF_R04_QUERY_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORF_R04_QUERY_RESPONSE struct {
	PATIENT *ORF_R04_PATIENT `hl7:"false,PATIENT"`
	ORDER   []ORF_R04_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *ORF_R04_QUERY_RESPONSE) MessageTypeSubStructName() string {
	return "ORF_R04_QUERY_RESPONSE"
}

// ORF_R04_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORF_R04_RESPONSE struct {
	PATIENT *ORF_R04_PATIENT `hl7:"false,PATIENT"`
	ORDER   []ORF_R04_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *ORF_R04_RESPONSE) MessageTypeSubStructName() string {
	return "ORF_R04_RESPONSE"
}

// ORG_O20 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORG_O20 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *ORG_O20_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *ORG_O20) MessageTypeName() string {
	return "ORG_O20"
}

// ORG_O20_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORG_O20_ORDER struct {
	ORC   *ORC  `hl7:"true,ORC"`
	OBR   *OBR  `hl7:"false,OBR"`
	NTE   []NTE `hl7:"false,NTE"`
	CTI   []CTI `hl7:"false,CTI"`
	Other []interface{}
}

func (s *ORG_O20_ORDER) MessageTypeSubStructName() string {
	return "ORG_O20_ORDER"
}

// ORG_O20_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORG_O20_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORG_O20_PATIENT) MessageTypeSubStructName() string {
	return "ORG_O20_PATIENT"
}

// ORG_O20_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORG_O20_RESPONSE struct {
	PATIENT *ORG_O20_PATIENT `hl7:"false,PATIENT"`
	ORDER   []ORG_O20_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *ORG_O20_RESPONSE) MessageTypeSubStructName() string {
	return "ORG_O20_RESPONSE"
}

// ORL_O22_CONTAINER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORL_O22_CONTAINER struct {
	SAC   *SAC  `hl7:"true,SAC"`
	OBX   []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *ORL_O22_CONTAINER) MessageTypeSubStructName() string {
	return "ORL_O22_CONTAINER"
}

// ORL_O22 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORL_O22 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *ORL_O22_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *ORL_O22) MessageTypeName() string {
	return "ORL_O22"
}

// ORL_O22_GENERAL_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORL_O22_GENERAL_ORDER struct {
	CONTAINER *ORL_O22_CONTAINER `hl7:"false,CONTAINER"`
	ORDER     []ORL_O22_ORDER    `hl7:"false,ORDER"`
	Other     []interface{}
}

func (s *ORL_O22_GENERAL_ORDER) MessageTypeSubStructName() string {
	return "ORL_O22_GENERAL_ORDER"
}

// ORL_O22_OBSERVATION_REQUEST represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORL_O22_OBSERVATION_REQUEST struct {
	OBR   *OBR  `hl7:"true,OBR"`
	SAC   []SAC `hl7:"false,SAC"`
	Other []interface{}
}

func (s *ORL_O22_OBSERVATION_REQUEST) MessageTypeSubStructName() string {
	return "ORL_O22_OBSERVATION_REQUEST"
}

// ORL_O22_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORL_O22_ORDER struct {
	ORC                 *ORC                         `hl7:"true,ORC"`
	OBSERVATION_REQUEST *ORL_O22_OBSERVATION_REQUEST `hl7:"false,OBSERVATION_REQUEST"`
	Other               []interface{}
}

func (s *ORL_O22_ORDER) MessageTypeSubStructName() string {
	return "ORL_O22_ORDER"
}

// ORL_O22_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORL_O22_PATIENT struct {
	PID           *PID                    `hl7:"true,PID"`
	GENERAL_ORDER []ORL_O22_GENERAL_ORDER `hl7:"true,GENERAL_ORDER"`
	Other         []interface{}
}

func (s *ORL_O22_PATIENT) MessageTypeSubStructName() string {
	return "ORL_O22_PATIENT"
}

// ORL_O22_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORL_O22_RESPONSE struct {
	PATIENT *ORL_O22_PATIENT `hl7:"false,PATIENT"`
	Other   []interface{}
}

func (s *ORL_O22_RESPONSE) MessageTypeSubStructName() string {
	return "ORL_O22_RESPONSE"
}

// ORM_O01_CHOICE represents the corresponding HL7 message type.
// Definition from HL7 2.4
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
// Definition from HL7 2.4
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

// ORM_O01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORM_O01_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *ORM_O01_INSURANCE) MessageTypeSubStructName() string {
	return "ORM_O01_INSURANCE"
}

// ORM_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORM_O01_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORM_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "ORM_O01_OBSERVATION"
}

// ORM_O01_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORM_O01_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *ORM_O01_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	FT1          []FT1                 `hl7:"false,FT1"`
	CTI          []CTI                 `hl7:"false,CTI"`
	BLG          *BLG                  `hl7:"false,BLG"`
	Other        []interface{}
}

func (s *ORM_O01_ORDER) MessageTypeSubStructName() string {
	return "ORM_O01_ORDER"
}

// ORM_O01_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORM_O01_ORDER_DETAIL struct {
	CHOICE      *ORM_O01_CHOICE       `hl7:"true,CHOICE"`
	NTE         []NTE                 `hl7:"false,NTE"`
	CTD         *CTD                  `hl7:"false,CTD"`
	DG1         []DG1                 `hl7:"false,DG1"`
	OBSERVATION []ORM_O01_OBSERVATION `hl7:"false,OBSERVATION"`
	Other       []interface{}
}

func (s *ORM_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "ORM_O01_ORDER_DETAIL"
}

// ORM_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORM_O01_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NTE           []NTE                  `hl7:"false,NTE"`
	PATIENT_VISIT *ORM_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE     []ORM_O01_INSURANCE    `hl7:"false,INSURANCE"`
	GT1           *GT1                   `hl7:"false,GT1"`
	AL1           []AL1                  `hl7:"false,AL1"`
	Other         []interface{}
}

func (s *ORM_O01_PATIENT) MessageTypeSubStructName() string {
	return "ORM_O01_PATIENT"
}

// ORM_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORM_O01_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *ORM_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "ORM_O01_PATIENT_VISIT"
}

// ORN_O02 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORN_O02 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *ORN_O02_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *ORN_O02) MessageTypeName() string {
	return "ORN_O02"
}

// ORN_O02_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORN_O02_ORDER struct {
	ORC   *ORC  `hl7:"true,ORC"`
	RQD   *RQD  `hl7:"true,RQD"`
	RQ1   *RQ1  `hl7:"false,RQ1"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORN_O02_ORDER) MessageTypeSubStructName() string {
	return "ORN_O02_ORDER"
}

// ORN_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORN_O02_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORN_O02_PATIENT) MessageTypeSubStructName() string {
	return "ORN_O02_PATIENT"
}

// ORN_O02_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORN_O02_RESPONSE struct {
	PATIENT *ORN_O02_PATIENT `hl7:"false,PATIENT"`
	ORDER   []ORN_O02_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *ORN_O02_RESPONSE) MessageTypeSubStructName() string {
	return "ORN_O02_RESPONSE"
}

// ORN_O08 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORN_O08 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *ORN_O08_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *ORN_O08) MessageTypeName() string {
	return "ORN_O08"
}

// ORN_O08_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORN_O08_ORDER struct {
	ORC   *ORC  `hl7:"true,ORC"`
	RQD   *RQD  `hl7:"true,RQD"`
	RQ1   *RQ1  `hl7:"false,RQ1"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORN_O08_ORDER) MessageTypeSubStructName() string {
	return "ORN_O08_ORDER"
}

// ORN_O08_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORN_O08_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORN_O08_PATIENT) MessageTypeSubStructName() string {
	return "ORN_O08_PATIENT"
}

// ORN_O08_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORN_O08_RESPONSE struct {
	PATIENT *ORN_O08_PATIENT `hl7:"false,PATIENT"`
	ORDER   []ORN_O08_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *ORN_O08_RESPONSE) MessageTypeSubStructName() string {
	return "ORN_O08_RESPONSE"
}

// ORP_O10 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORP_O10 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *ORP_O10_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *ORP_O10) MessageTypeName() string {
	return "ORP_O10"
}

// ORP_O10_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORP_O10_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *ORP_O10_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other        []interface{}
}

func (s *ORP_O10_ORDER) MessageTypeSubStructName() string {
	return "ORP_O10_ORDER"
}

// ORP_O10_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORP_O10_ORDER_DETAIL struct {
	RXO   *RXO  `hl7:"true,RXO"`
	NTE1  []NTE `hl7:"false,NTE1"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	NTE2  []NTE `hl7:"false,NTE2"`
	Other []interface{}
}

func (s *ORP_O10_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "ORP_O10_ORDER_DETAIL"
}

// ORP_O10_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORP_O10_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORP_O10_PATIENT) MessageTypeSubStructName() string {
	return "ORP_O10_PATIENT"
}

// ORP_O10_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORP_O10_RESPONSE struct {
	PATIENT *ORP_O10_PATIENT `hl7:"false,PATIENT"`
	ORDER   []ORP_O10_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *ORP_O10_RESPONSE) MessageTypeSubStructName() string {
	return "ORP_O10_RESPONSE"
}

// ORR_O02_CHOICE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
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
// Definition from HL7 2.4
type ORR_O02 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *ORR_O02_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *ORR_O02) MessageTypeName() string {
	return "ORR_O02"
}

// ORR_O02_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORR_O02_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *ORR_O02_ORDER_DETAIL `hl7:"true,ORDER_DETAIL"`
	NTE          []NTE                 `hl7:"false,NTE"`
	CTI          []CTI                 `hl7:"false,CTI"`
	Other        []interface{}
}

func (s *ORR_O02_ORDER) MessageTypeSubStructName() string {
	return "ORR_O02_ORDER"
}

// ORR_O02_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORR_O02_ORDER_DETAIL struct {
	// Only one of the following fields will be set.
	OBR   *OBR `hl7:"true,OBR"`
	RQD   *RQD `hl7:"true,RQD"`
	RQ1   *RQ1 `hl7:"true,RQ1"`
	RXO   *RXO `hl7:"true,RXO"`
	ODS   *ODS `hl7:"true,ODS"`
	ODT   *ODT `hl7:"true,ODT"`
	Other []interface{}
}

func (s *ORR_O02_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "ORR_O02_ORDER_DETAIL"
}

// ORR_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORR_O02_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORR_O02_PATIENT) MessageTypeSubStructName() string {
	return "ORR_O02_PATIENT"
}

// ORR_O02_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORR_O02_RESPONSE struct {
	PATIENT *ORR_O02_PATIENT `hl7:"false,PATIENT"`
	ORDER   []ORR_O02_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *ORR_O02_RESPONSE) MessageTypeSubStructName() string {
	return "ORR_O02_RESPONSE"
}

// ORS_O02 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORS_O02 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *ORS_O02_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *ORS_O02) MessageTypeName() string {
	return "ORS_O02"
}

// ORS_O02_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORS_O02_ORDER struct {
	ORC   *ORC  `hl7:"true,ORC"`
	RQD   *RQD  `hl7:"true,RQD"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORS_O02_ORDER) MessageTypeSubStructName() string {
	return "ORS_O02_ORDER"
}

// ORS_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORS_O02_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORS_O02_PATIENT) MessageTypeSubStructName() string {
	return "ORS_O02_PATIENT"
}

// ORS_O02_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORS_O02_RESPONSE struct {
	PATIENT *ORS_O02_PATIENT `hl7:"false,PATIENT"`
	ORDER   []ORS_O02_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *ORS_O02_RESPONSE) MessageTypeSubStructName() string {
	return "ORS_O02_RESPONSE"
}

// ORS_O06 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORS_O06 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	MSA     *MSA             `hl7:"true,MSA"`
	ERR     *ERR             `hl7:"false,ERR"`
	NTE     []NTE            `hl7:"false,NTE"`
	RSPONSE *ORS_O06_RSPONSE `hl7:"false,RSPONSE"`
	Other   []interface{}
}

func (s *ORS_O06) MessageTypeName() string {
	return "ORS_O06"
}

// ORS_O06_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORS_O06_ORDER struct {
	ORC   *ORC  `hl7:"true,ORC"`
	RQD   *RQD  `hl7:"true,RQD"`
	RQ1   *RQ1  `hl7:"false,RQ1"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORS_O06_ORDER) MessageTypeSubStructName() string {
	return "ORS_O06_ORDER"
}

// ORS_O06_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORS_O06_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORS_O06_PATIENT) MessageTypeSubStructName() string {
	return "ORS_O06_PATIENT"
}

// ORS_O06_RSPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORS_O06_RSPONSE struct {
	PATIENT *ORS_O06_PATIENT `hl7:"false,PATIENT"`
	ORDER   []ORS_O06_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *ORS_O06_RSPONSE) MessageTypeSubStructName() string {
	return "ORS_O06_RSPONSE"
}

// ORU_R01 represents the corresponding HL7 message type.
// Definition from HL7 2.4
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
// Definition from HL7 2.4
type ORU_R01_OBSERVATION struct {
	OBX   *OBX  `hl7:"false,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORU_R01_OBSERVATION) MessageTypeSubStructName() string {
	return "ORU_R01_OBSERVATION"
}

// ORU_R01_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORU_R01_ORDER_OBSERVATION struct {
	ORC         *ORC                  `hl7:"false,ORC"`
	OBR         *OBR                  `hl7:"true,OBR"`
	NTE         []NTE                 `hl7:"false,NTE"`
	CTD         *CTD                  `hl7:"false,CTD"`
	OBSERVATION []ORU_R01_OBSERVATION `hl7:"true,OBSERVATION"`
	FT1         []FT1                 `hl7:"false,FT1"`
	CTI         []CTI                 `hl7:"false,CTI"`
	Other       []interface{}
}

func (s *ORU_R01_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "ORU_R01_ORDER_OBSERVATION"
}

// ORU_R01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORU_R01_PATIENT struct {
	PID   *PID           `hl7:"true,PID"`
	PD1   *PD1           `hl7:"false,PD1"`
	NK1   []NK1          `hl7:"false,NK1"`
	NTE   []NTE          `hl7:"false,NTE"`
	VISIT *ORU_R01_VISIT `hl7:"false,VISIT"`
	Other []interface{}
}

func (s *ORU_R01_PATIENT) MessageTypeSubStructName() string {
	return "ORU_R01_PATIENT"
}

// ORU_R01_PATIENT_RESULT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORU_R01_PATIENT_RESULT struct {
	PATIENT           *ORU_R01_PATIENT            `hl7:"false,PATIENT"`
	ORDER_OBSERVATION []ORU_R01_ORDER_OBSERVATION `hl7:"true,ORDER_OBSERVATION"`
	Other             []interface{}
}

func (s *ORU_R01_PATIENT_RESULT) MessageTypeSubStructName() string {
	return "ORU_R01_PATIENT_RESULT"
}

// ORU_R01_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3
type ORU_R01_RESPONSE struct {
	PATIENT           *ORU_R01_PATIENT            `hl7:"false,PATIENT"`
	ORDER_OBSERVATION []ORU_R01_ORDER_OBSERVATION `hl7:"true,ORDER_OBSERVATION"`
	Other             []interface{}
}

func (s *ORU_R01_RESPONSE) MessageTypeSubStructName() string {
	return "ORU_R01_RESPONSE"
}

// ORU_R01_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ORU_R01_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *ORU_R01_VISIT) MessageTypeSubStructName() string {
	return "ORU_R01_VISIT"
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
// Definition from HL7 2.3.1
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
// Definition from HL7 2.3.1
type ORU_R32_OBSERVATION struct {
	OBX   *OBX  `hl7:"false,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORU_R32_OBSERVATION) MessageTypeSubStructName() string {
	return "ORU_R32_OBSERVATION"
}

// ORU_R32_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORU_R32_ORDER_OBSERVATION struct {
	ORC         *ORC                  `hl7:"false,ORC"`
	OBR         *OBR                  `hl7:"true,OBR"`
	NTE         []NTE                 `hl7:"false,NTE"`
	OBSERVATION []ORU_R32_OBSERVATION `hl7:"true,OBSERVATION"`
	CTI         []CTI                 `hl7:"false,CTI"`
	Other       []interface{}
}

func (s *ORU_R32_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "ORU_R32_ORDER_OBSERVATION"
}

// ORU_R32_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORU_R32_PATIENT struct {
	PID   *PID           `hl7:"true,PID"`
	PD1   *PD1           `hl7:"false,PD1"`
	NK1   []NK1          `hl7:"false,NK1"`
	NTE   []NTE          `hl7:"false,NTE"`
	VISIT *ORU_R32_VISIT `hl7:"false,VISIT"`
	Other []interface{}
}

func (s *ORU_R32_PATIENT) MessageTypeSubStructName() string {
	return "ORU_R32_PATIENT"
}

// ORU_R32_PATIENT_RESULT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORU_R32_PATIENT_RESULT struct {
	PATIENT           *ORU_R32_PATIENT            `hl7:"false,PATIENT"`
	ORDER_OBSERVATION []ORU_R32_ORDER_OBSERVATION `hl7:"true,ORDER_OBSERVATION"`
	Other             []interface{}
}

func (s *ORU_R32_PATIENT_RESULT) MessageTypeSubStructName() string {
	return "ORU_R32_PATIENT_RESULT"
}

// ORU_R32_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type ORU_R32_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *ORU_R32_VISIT) MessageTypeSubStructName() string {
	return "ORU_R32_VISIT"
}

// OSQ_Q06 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OSQ_Q06 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QRD   *QRD `hl7:"true,QRD"`
	QRF   *QRF `hl7:"false,QRF"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *OSQ_Q06) MessageTypeName() string {
	return "OSQ_Q06"
}

// OSR_Q06_CHOICE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OSR_Q06_CHOICE struct {
	// Only one of the following fields will be set.
	OBR   *OBR `hl7:"true,OBR"`
	RQD   *RQD `hl7:"true,RQD"`
	RQ1   *RQ1 `hl7:"true,RQ1"`
	RXO   *RXO `hl7:"true,RXO"`
	ODS   *ODS `hl7:"true,ODS"`
	ODT   *ODT `hl7:"true,ODT"`
	Other []interface{}
}

func (s *OSR_Q06_CHOICE) MessageTypeSubStructName() string {
	return "OSR_Q06_CHOICE"
}

// OSR_Q06 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OSR_Q06 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	QRD      *QRD              `hl7:"true,QRD"`
	QRF      *QRF              `hl7:"false,QRF"`
	RESPONSE *OSR_Q06_RESPONSE `hl7:"false,RESPONSE"`
	DSC      *DSC              `hl7:"false,DSC"`
	Other    []interface{}
}

func (s *OSR_Q06) MessageTypeName() string {
	return "OSR_Q06"
}

// OSR_Q06_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type OSR_Q06_OBSERVATION struct {
	ORC    *ORC            `hl7:"true,ORC"`
	CHOICE *OSR_Q06_CHOICE `hl7:"true,CHOICE"`
	NTE    []NTE           `hl7:"false,NTE"`
	CTI    []CTI           `hl7:"false,CTI"`
	Other  []interface{}
}

func (s *OSR_Q06_OBSERVATION) MessageTypeSubStructName() string {
	return "OSR_Q06_OBSERVATION"
}

// OSR_Q06_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OSR_Q06_ORDER struct {
	ORC    *ORC            `hl7:"true,ORC"`
	CHOICE *OSR_Q06_CHOICE `hl7:"true,CHOICE"`
	NTE    []NTE           `hl7:"false,NTE"`
	CTI    []CTI           `hl7:"false,CTI"`
	Other  []interface{}
}

func (s *OSR_Q06_ORDER) MessageTypeSubStructName() string {
	return "OSR_Q06_ORDER"
}

// OSR_Q06_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OSR_Q06_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OSR_Q06_PATIENT) MessageTypeSubStructName() string {
	return "OSR_Q06_PATIENT"
}

// OSR_Q06_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OSR_Q06_RESPONSE struct {
	PATIENT *OSR_Q06_PATIENT `hl7:"false,PATIENT"`
	ORDER   []OSR_Q06_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *OSR_Q06_RESPONSE) MessageTypeSubStructName() string {
	return "OSR_Q06_RESPONSE"
}

// OUL_R21_CONTAINER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OUL_R21_CONTAINER struct {
	SAC   *SAC  `hl7:"true,SAC"`
	SID   *SID  `hl7:"false,SID"`
	OBX   []OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *OUL_R21_CONTAINER) MessageTypeSubStructName() string {
	return "OUL_R21_CONTAINER"
}

// OUL_R21 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OUL_R21 struct {
	MSH               *MSH                        `hl7:"true,MSH"`
	NTE               *NTE                        `hl7:"false,NTE"`
	PATIENT           *OUL_R21_PATIENT            `hl7:"false,PATIENT"`
	VISIT             *OUL_R21_VISIT              `hl7:"false,VISIT"`
	ORDER_OBSERVATION []OUL_R21_ORDER_OBSERVATION `hl7:"true,ORDER_OBSERVATION"`
	DSC               *DSC                        `hl7:"false,DSC"`
	Other             []interface{}
}

func (s *OUL_R21) MessageTypeName() string {
	return "OUL_R21"
}

// OUL_R21_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OUL_R21_OBSERVATION struct {
	OBX   *OBX  `hl7:"false,OBX"`
	TCD   *TCD  `hl7:"false,TCD"`
	SID   []SID `hl7:"false,SID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OUL_R21_OBSERVATION) MessageTypeSubStructName() string {
	return "OUL_R21_OBSERVATION"
}

// OUL_R21_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OUL_R21_ORDER_OBSERVATION struct {
	CONTAINER   *OUL_R21_CONTAINER    `hl7:"false,CONTAINER"`
	ORC         *ORC                  `hl7:"false,ORC"`
	OBR         *OBR                  `hl7:"true,OBR"`
	NTE         []NTE                 `hl7:"false,NTE"`
	OBSERVATION []OUL_R21_OBSERVATION `hl7:"true,OBSERVATION"`
	CTI         []CTI                 `hl7:"false,CTI"`
	Other       []interface{}
}

func (s *OUL_R21_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "OUL_R21_ORDER_OBSERVATION"
}

// OUL_R21_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OUL_R21_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	PD1   *PD1  `hl7:"false,PD1"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *OUL_R21_PATIENT) MessageTypeSubStructName() string {
	return "OUL_R21_PATIENT"
}

// OUL_R21_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type OUL_R21_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *OUL_R21_VISIT) MessageTypeSubStructName() string {
	return "OUL_R21_VISIT"
}

// PEX_P07_ASSOCIATED_PERSON represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PEX_P07_ASSOCIATED_PERSON struct {
	NK1                 *NK1                          `hl7:"true,NK1"`
	ASSOCIATED_RX_ORDER *PEX_P07_ASSOCIATED_RX_ORDER  `hl7:"false,ASSOCIATED_RX_ORDER"`
	ASSOCIATED_RX_ADMIN []PEX_P07_ASSOCIATED_RX_ADMIN `hl7:"false,ASSOCIATED_RX_ADMIN"`
	PRB                 []PRB                         `hl7:"false,PRB"`
	OBX                 []OBX                         `hl7:"false,OBX"`
	Other               []interface{}
}

func (s *PEX_P07_ASSOCIATED_PERSON) MessageTypeSubStructName() string {
	return "PEX_P07_ASSOCIATED_PERSON"
}

// PEX_P07_ASSOCIATED_RX_ADMIN represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PEX_P07_ASSOCIATED_RX_ADMIN struct {
	RXA   *RXA `hl7:"true,RXA"`
	RXR   *RXR `hl7:"false,RXR"`
	Other []interface{}
}

func (s *PEX_P07_ASSOCIATED_RX_ADMIN) MessageTypeSubStructName() string {
	return "PEX_P07_ASSOCIATED_RX_ADMIN"
}

// PEX_P07_ASSOCIATED_RX_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PEX_P07_ASSOCIATED_RX_ORDER struct {
	RXE   *RXE  `hl7:"true,RXE"`
	RXR   []RXR `hl7:"false,RXR"`
	Other []interface{}
}

func (s *PEX_P07_ASSOCIATED_RX_ORDER) MessageTypeSubStructName() string {
	return "PEX_P07_ASSOCIATED_RX_ORDER"
}

// PEX_P07 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PEX_P07 struct {
	MSH        *MSH                 `hl7:"true,MSH"`
	EVN        *EVN                 `hl7:"true,EVN"`
	PID        *PID                 `hl7:"true,PID"`
	PD1        *PD1                 `hl7:"false,PD1"`
	NTE        []NTE                `hl7:"false,NTE"`
	VISIT      *PEX_P07_VISIT       `hl7:"false,VISIT"`
	EXPERIENCE []PEX_P07_EXPERIENCE `hl7:"true,EXPERIENCE"`
	Other      []interface{}
}

func (s *PEX_P07) MessageTypeName() string {
	return "PEX_P07"
}

// PEX_P07_EXPERIENCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PEX_P07_EXPERIENCE struct {
	PES             *PES                      `hl7:"true,PES"`
	PEX_OBSERVATION []PEX_P07_PEX_OBSERVATION `hl7:"true,PEX_OBSERVATION"`
	Other           []interface{}
}

func (s *PEX_P07_EXPERIENCE) MessageTypeSubStructName() string {
	return "PEX_P07_EXPERIENCE"
}

// PEX_P07_PEX_CAUSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PEX_P07_PEX_CAUSE struct {
	PCR               *PCR                        `hl7:"true,PCR"`
	RX_ORDER          *PEX_P07_RX_ORDER           `hl7:"false,RX_ORDER"`
	RX_ADMINISTRATION []PEX_P07_RX_ADMINISTRATION `hl7:"false,RX_ADMINISTRATION"`
	PRB               []PRB                       `hl7:"false,PRB"`
	OBX               []OBX                       `hl7:"false,OBX"`
	NTE               []NTE                       `hl7:"false,NTE"`
	ASSOCIATED_PERSON *PEX_P07_ASSOCIATED_PERSON  `hl7:"false,ASSOCIATED_PERSON"`
	STUDY             []PEX_P07_STUDY             `hl7:"false,STUDY"`
	Other             []interface{}
}

func (s *PEX_P07_PEX_CAUSE) MessageTypeSubStructName() string {
	return "PEX_P07_PEX_CAUSE"
}

// PEX_P07_PEX_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PEX_P07_PEX_OBSERVATION struct {
	PEO       *PEO                `hl7:"true,PEO"`
	PEX_CAUSE []PEX_P07_PEX_CAUSE `hl7:"true,PEX_CAUSE"`
	Other     []interface{}
}

func (s *PEX_P07_PEX_OBSERVATION) MessageTypeSubStructName() string {
	return "PEX_P07_PEX_OBSERVATION"
}

// PEX_P07_RX_ADMINISTRATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PEX_P07_RX_ADMINISTRATION struct {
	RXA   *RXA `hl7:"true,RXA"`
	RXR   *RXR `hl7:"false,RXR"`
	Other []interface{}
}

func (s *PEX_P07_RX_ADMINISTRATION) MessageTypeSubStructName() string {
	return "PEX_P07_RX_ADMINISTRATION"
}

// PEX_P07_RX_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PEX_P07_RX_ORDER struct {
	RXE   *RXE  `hl7:"true,RXE"`
	RXR   []RXR `hl7:"false,RXR"`
	Other []interface{}
}

func (s *PEX_P07_RX_ORDER) MessageTypeSubStructName() string {
	return "PEX_P07_RX_ORDER"
}

// PEX_P07_STUDY represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PEX_P07_STUDY struct {
	CSR   *CSR  `hl7:"true,CSR"`
	CSP   []CSP `hl7:"false,CSP"`
	Other []interface{}
}

func (s *PEX_P07_STUDY) MessageTypeSubStructName() string {
	return "PEX_P07_STUDY"
}

// PEX_P07_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PEX_P07_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PEX_P07_VISIT) MessageTypeSubStructName() string {
	return "PEX_P07_VISIT"
}

// PGL_PC6_CHOICE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type PGL_PC6_CHOICE struct {
	// Only one of the following fields will be set.
	OBR   *OBR `hl7:"true,OBR"`
	RXO   *RXO `hl7:"true,RXO"`
	Other []interface{}
}

func (s *PGL_PC6_CHOICE) MessageTypeSubStructName() string {
	return "PGL_PC6_CHOICE"
}

// PGL_PC6 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PGL_PC6 struct {
	MSH           *MSH                   `hl7:"true,MSH"`
	PID           *PID                   `hl7:"true,PID"`
	PATIENT_VISIT *PGL_PC6_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	GOAL          []PGL_PC6_GOAL         `hl7:"true,GOAL"`
	Other         []interface{}
}

func (s *PGL_PC6) MessageTypeName() string {
	return "PGL_PC6"
}

// PGL_PC6_GOAL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PGL_PC6_GOAL struct {
	GOL         *GOL                  `hl7:"true,GOL"`
	NTE         []NTE                 `hl7:"false,NTE"`
	VAR         []VAR                 `hl7:"false,VAR"`
	GOAL_ROLE   []PGL_PC6_GOAL_ROLE   `hl7:"false,GOAL_ROLE"`
	PATHWAY     []PGL_PC6_PATHWAY     `hl7:"false,PATHWAY"`
	OBSERVATION []PGL_PC6_OBSERVATION `hl7:"false,OBSERVATION"`
	PROBLEM     []PGL_PC6_PROBLEM     `hl7:"false,PROBLEM"`
	ORDER       []PGL_PC6_ORDER       `hl7:"false,ORDER"`
	Other       []interface{}
}

func (s *PGL_PC6_GOAL) MessageTypeSubStructName() string {
	return "PGL_PC6_GOAL"
}

// PGL_PC6_GOAL_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PGL_PC6_GOAL_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PGL_PC6_GOAL_ROLE) MessageTypeSubStructName() string {
	return "PGL_PC6_GOAL_ROLE"
}

// PGL_PC6_OBRRXO_SUPPGRP represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PGL_PC6_OBRRXO_SUPPGRP struct {
	// Only one of the following fields will be set.
	OBR   *OBR `hl7:"true,OBR"`
	RXO   *RXO `hl7:"true,RXO"`
	Other []interface{}
}

func (s *PGL_PC6_OBRRXO_SUPPGRP) MessageTypeSubStructName() string {
	return "PGL_PC6_OBRRXO_SUPPGRP"
}

// PGL_PC6_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PGL_PC6_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PGL_PC6_OBSERVATION) MessageTypeSubStructName() string {
	return "PGL_PC6_OBSERVATION"
}

// PGL_PC6_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PGL_PC6_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *PGL_PC6_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other        []interface{}
}

func (s *PGL_PC6_ORDER) MessageTypeSubStructName() string {
	return "PGL_PC6_ORDER"
}

// PGL_PC6_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PGL_PC6_ORDER_DETAIL struct {
	OBRRXO_SUPPGRP    *PGL_PC6_OBRRXO_SUPPGRP     `hl7:"true,OBRRXO_SUPPGRP"`
	NTE               []NTE                       `hl7:"false,NTE"`
	VAR               []VAR                       `hl7:"false,VAR"`
	ORDER_OBSERVATION []PGL_PC6_ORDER_OBSERVATION `hl7:"false,ORDER_OBSERVATION"`
	Other             []interface{}
}

func (s *PGL_PC6_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "PGL_PC6_ORDER_DETAIL"
}

// PGL_PC6_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PGL_PC6_ORDER_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PGL_PC6_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "PGL_PC6_ORDER_OBSERVATION"
}

// PGL_PC6_PATHWAY represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PGL_PC6_PATHWAY struct {
	PTH   *PTH  `hl7:"true,PTH"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PGL_PC6_PATHWAY) MessageTypeSubStructName() string {
	return "PGL_PC6_PATHWAY"
}

// PGL_PC6_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PGL_PC6_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PGL_PC6_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "PGL_PC6_PATIENT_VISIT"
}

// PGL_PC6_PROBLEM represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PGL_PC6_PROBLEM struct {
	PRB                 *PRB                          `hl7:"true,PRB"`
	NTE                 []NTE                         `hl7:"false,NTE"`
	VAR                 []VAR                         `hl7:"false,VAR"`
	PROBLEM_ROLE        []PGL_PC6_PROBLEM_ROLE        `hl7:"false,PROBLEM_ROLE"`
	PROBLEM_OBSERVATION []PGL_PC6_PROBLEM_OBSERVATION `hl7:"false,PROBLEM_OBSERVATION"`
	Other               []interface{}
}

func (s *PGL_PC6_PROBLEM) MessageTypeSubStructName() string {
	return "PGL_PC6_PROBLEM"
}

// PGL_PC6_PROBLEM_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PGL_PC6_PROBLEM_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PGL_PC6_PROBLEM_OBSERVATION) MessageTypeSubStructName() string {
	return "PGL_PC6_PROBLEM_OBSERVATION"
}

// PGL_PC6_PROBLEM_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PGL_PC6_PROBLEM_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PGL_PC6_PROBLEM_ROLE) MessageTypeSubStructName() string {
	return "PGL_PC6_PROBLEM_ROLE"
}

// PIN_I07 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type PIN_I07 struct {
	MSH                 *MSH                         `hl7:"true,MSH"`
	PROVIDER            []PIN_I07_PROVIDER           `hl7:"true,PROVIDER"`
	PID                 *PID                         `hl7:"true,PID"`
	NK1                 []NK1                        `hl7:"false,NK1"`
	GUARANTOR_INSURANCE *PIN_I07_GUARANTOR_INSURANCE `hl7:"false,GUARANTOR_INSURANCE"`
	NTE                 []NTE                        `hl7:"false,NTE"`
	Other               []interface{}
}

func (s *PIN_I07) MessageTypeName() string {
	return "PIN_I07"
}

// PIN_I07_GUARANTOR_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type PIN_I07_GUARANTOR_INSURANCE struct {
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []PIN_I07_INSURANCE `hl7:"true,INSURANCE"`
	Other     []interface{}
}

func (s *PIN_I07_GUARANTOR_INSURANCE) MessageTypeSubStructName() string {
	return "PIN_I07_GUARANTOR_INSURANCE"
}

// PIN_I07_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type PIN_I07_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *PIN_I07_INSURANCE) MessageTypeSubStructName() string {
	return "PIN_I07_INSURANCE"
}

// PIN_I07_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type PIN_I07_PROVIDER struct {
	PRD   *PRD  `hl7:"true,PRD"`
	CTD   []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *PIN_I07_PROVIDER) MessageTypeSubStructName() string {
	return "PIN_I07_PROVIDER"
}

// PMU_B01 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PMU_B01 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	STF   *STF  `hl7:"true,STF"`
	PRA   []PRA `hl7:"false,PRA"`
	ORG   []ORG `hl7:"false,ORG"`
	AFF   []AFF `hl7:"false,AFF"`
	LAN   []LAN `hl7:"false,LAN"`
	EDU   []EDU `hl7:"false,EDU"`
	Other []interface{}
}

func (s *PMU_B01) MessageTypeName() string {
	return "PMU_B01"
}

// PMU_B03 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PMU_B03 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	STF   *STF `hl7:"true,STF"`
	Other []interface{}
}

func (s *PMU_B03) MessageTypeName() string {
	return "PMU_B03"
}

// PMU_B04 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PMU_B04 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	STF   *STF  `hl7:"true,STF"`
	PRA   []PRA `hl7:"false,PRA"`
	ORG   *ORG  `hl7:"false,ORG"`
	Other []interface{}
}

func (s *PMU_B04) MessageTypeName() string {
	return "PMU_B04"
}

// PPG_PCG_CHOICE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type PPG_PCG_CHOICE struct {
	// Only one of the following fields will be set.
	OBR   *OBR `hl7:"true,OBR"`
	RXO   *RXO `hl7:"true,RXO"`
	Other []interface{}
}

func (s *PPG_PCG_CHOICE) MessageTypeSubStructName() string {
	return "PPG_PCG_CHOICE"
}

// PPG_PCG represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPG_PCG struct {
	MSH           *MSH                   `hl7:"true,MSH"`
	PID           *PID                   `hl7:"true,PID"`
	PATIENT_VISIT *PPG_PCG_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	PATHWAY       []PPG_PCG_PATHWAY      `hl7:"true,PATHWAY"`
	Other         []interface{}
}

func (s *PPG_PCG) MessageTypeName() string {
	return "PPG_PCG"
}

// PPG_PCG_GOAL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPG_PCG_GOAL struct {
	GOL              *GOL                       `hl7:"true,GOL"`
	NTE              []NTE                      `hl7:"false,NTE"`
	VAR              []VAR                      `hl7:"false,VAR"`
	GOAL_ROLE        []PPG_PCG_GOAL_ROLE        `hl7:"false,GOAL_ROLE"`
	GOAL_OBSERVATION []PPG_PCG_GOAL_OBSERVATION `hl7:"false,GOAL_OBSERVATION"`
	PROBLEM          []PPG_PCG_PROBLEM          `hl7:"false,PROBLEM"`
	ORDER            []PPG_PCG_ORDER            `hl7:"false,ORDER"`
	Other            []interface{}
}

func (s *PPG_PCG_GOAL) MessageTypeSubStructName() string {
	return "PPG_PCG_GOAL"
}

// PPG_PCG_GOAL_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPG_PCG_GOAL_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPG_PCG_GOAL_OBSERVATION) MessageTypeSubStructName() string {
	return "PPG_PCG_GOAL_OBSERVATION"
}

// PPG_PCG_GOAL_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPG_PCG_GOAL_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPG_PCG_GOAL_ROLE) MessageTypeSubStructName() string {
	return "PPG_PCG_GOAL_ROLE"
}

// PPG_PCG_OBRRXO_SUPPGRP represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPG_PCG_OBRRXO_SUPPGRP struct {
	// Only one of the following fields will be set.
	OBR   *OBR `hl7:"true,OBR"`
	RXO   *RXO `hl7:"true,RXO"`
	Other []interface{}
}

func (s *PPG_PCG_OBRRXO_SUPPGRP) MessageTypeSubStructName() string {
	return "PPG_PCG_OBRRXO_SUPPGRP"
}

// PPG_PCG_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPG_PCG_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *PPG_PCG_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other        []interface{}
}

func (s *PPG_PCG_ORDER) MessageTypeSubStructName() string {
	return "PPG_PCG_ORDER"
}

// PPG_PCG_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPG_PCG_ORDER_DETAIL struct {
	OBRRXO_SUPPGRP    *PPG_PCG_OBRRXO_SUPPGRP     `hl7:"true,OBRRXO_SUPPGRP"`
	NTE               []NTE                       `hl7:"false,NTE"`
	VAR               []VAR                       `hl7:"false,VAR"`
	ORDER_OBSERVATION []PPG_PCG_ORDER_OBSERVATION `hl7:"false,ORDER_OBSERVATION"`
	Other             []interface{}
}

func (s *PPG_PCG_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "PPG_PCG_ORDER_DETAIL"
}

// PPG_PCG_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPG_PCG_ORDER_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPG_PCG_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "PPG_PCG_ORDER_OBSERVATION"
}

// PPG_PCG_PATHWAY represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPG_PCG_PATHWAY struct {
	PTH          *PTH                   `hl7:"true,PTH"`
	NTE          []NTE                  `hl7:"false,NTE"`
	VAR          []VAR                  `hl7:"false,VAR"`
	PATHWAY_ROLE []PPG_PCG_PATHWAY_ROLE `hl7:"false,PATHWAY_ROLE"`
	GOAL         []PPG_PCG_GOAL         `hl7:"false,GOAL"`
	Other        []interface{}
}

func (s *PPG_PCG_PATHWAY) MessageTypeSubStructName() string {
	return "PPG_PCG_PATHWAY"
}

// PPG_PCG_PATHWAY_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPG_PCG_PATHWAY_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPG_PCG_PATHWAY_ROLE) MessageTypeSubStructName() string {
	return "PPG_PCG_PATHWAY_ROLE"
}

// PPG_PCG_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPG_PCG_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PPG_PCG_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "PPG_PCG_PATIENT_VISIT"
}

// PPG_PCG_PROBLEM represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPG_PCG_PROBLEM struct {
	PRB                 *PRB                          `hl7:"true,PRB"`
	NTE                 []NTE                         `hl7:"false,NTE"`
	VAR                 []VAR                         `hl7:"false,VAR"`
	PROBLEM_ROLE        []PPG_PCG_PROBLEM_ROLE        `hl7:"false,PROBLEM_ROLE"`
	PROBLEM_OBSERVATION []PPG_PCG_PROBLEM_OBSERVATION `hl7:"false,PROBLEM_OBSERVATION"`
	Other               []interface{}
}

func (s *PPG_PCG_PROBLEM) MessageTypeSubStructName() string {
	return "PPG_PCG_PROBLEM"
}

// PPG_PCG_PROBLEM_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPG_PCG_PROBLEM_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPG_PCG_PROBLEM_OBSERVATION) MessageTypeSubStructName() string {
	return "PPG_PCG_PROBLEM_OBSERVATION"
}

// PPG_PCG_PROBLEM_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPG_PCG_PROBLEM_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPG_PCG_PROBLEM_ROLE) MessageTypeSubStructName() string {
	return "PPG_PCG_PROBLEM_ROLE"
}

// PPP_PCB_CHOICE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPP_PCB_CHOICE struct {
	// Only one of the following fields will be set.
	OBR   *OBR `hl7:"true,OBR"`
	RXO   *RXO `hl7:"true,RXO"`
	Other []interface{}
}

func (s *PPP_PCB_CHOICE) MessageTypeSubStructName() string {
	return "PPP_PCB_CHOICE"
}

// PPP_PCB represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPP_PCB struct {
	MSH           *MSH                   `hl7:"true,MSH"`
	PID           *PID                   `hl7:"true,PID"`
	PATIENT_VISIT *PPP_PCB_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	PATHWAY       []PPP_PCB_PATHWAY      `hl7:"true,PATHWAY"`
	Other         []interface{}
}

func (s *PPP_PCB) MessageTypeName() string {
	return "PPP_PCB"
}

// PPP_PCB_GOAL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPP_PCB_GOAL struct {
	GOL              *GOL                       `hl7:"true,GOL"`
	NTE              []NTE                      `hl7:"false,NTE"`
	VAR              []VAR                      `hl7:"false,VAR"`
	GOAL_ROLE        []PPP_PCB_GOAL_ROLE        `hl7:"false,GOAL_ROLE"`
	GOAL_OBSERVATION []PPP_PCB_GOAL_OBSERVATION `hl7:"false,GOAL_OBSERVATION"`
	Other            []interface{}
}

func (s *PPP_PCB_GOAL) MessageTypeSubStructName() string {
	return "PPP_PCB_GOAL"
}

// PPP_PCB_GOAL_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPP_PCB_GOAL_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPP_PCB_GOAL_OBSERVATION) MessageTypeSubStructName() string {
	return "PPP_PCB_GOAL_OBSERVATION"
}

// PPP_PCB_GOAL_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPP_PCB_GOAL_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPP_PCB_GOAL_ROLE) MessageTypeSubStructName() string {
	return "PPP_PCB_GOAL_ROLE"
}

// PPP_PCB_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPP_PCB_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *PPP_PCB_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other        []interface{}
}

func (s *PPP_PCB_ORDER) MessageTypeSubStructName() string {
	return "PPP_PCB_ORDER"
}

// PPP_PCB_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPP_PCB_ORDER_DETAIL struct {
	CHOICE            *PPP_PCB_CHOICE             `hl7:"true,CHOICE"`
	NTE               []NTE                       `hl7:"false,NTE"`
	VAR               []VAR                       `hl7:"false,VAR"`
	ORDER_OBSERVATION []PPP_PCB_ORDER_OBSERVATION `hl7:"false,ORDER_OBSERVATION"`
	Other             []interface{}
}

func (s *PPP_PCB_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "PPP_PCB_ORDER_DETAIL"
}

// PPP_PCB_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPP_PCB_ORDER_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPP_PCB_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "PPP_PCB_ORDER_OBSERVATION"
}

// PPP_PCB_PATHWAY represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPP_PCB_PATHWAY struct {
	PTH          *PTH                   `hl7:"true,PTH"`
	NTE          []NTE                  `hl7:"false,NTE"`
	VAR          []VAR                  `hl7:"false,VAR"`
	PATHWAY_ROLE []PPP_PCB_PATHWAY_ROLE `hl7:"false,PATHWAY_ROLE"`
	PROBLEM      []PPP_PCB_PROBLEM      `hl7:"false,PROBLEM"`
	Other        []interface{}
}

func (s *PPP_PCB_PATHWAY) MessageTypeSubStructName() string {
	return "PPP_PCB_PATHWAY"
}

// PPP_PCB_PATHWAY_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPP_PCB_PATHWAY_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPP_PCB_PATHWAY_ROLE) MessageTypeSubStructName() string {
	return "PPP_PCB_PATHWAY_ROLE"
}

// PPP_PCB_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPP_PCB_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PPP_PCB_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "PPP_PCB_PATIENT_VISIT"
}

// PPP_PCB_PROBLEM represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPP_PCB_PROBLEM struct {
	PRB                 *PRB                          `hl7:"true,PRB"`
	NTE                 []NTE                         `hl7:"false,NTE"`
	VAR                 []VAR                         `hl7:"false,VAR"`
	PROBLEM_ROLE        []PPP_PCB_PROBLEM_ROLE        `hl7:"false,PROBLEM_ROLE"`
	PROBLEM_OBSERVATION []PPP_PCB_PROBLEM_OBSERVATION `hl7:"false,PROBLEM_OBSERVATION"`
	GOAL                []PPP_PCB_GOAL                `hl7:"false,GOAL"`
	ORDER               []PPP_PCB_ORDER               `hl7:"false,ORDER"`
	Other               []interface{}
}

func (s *PPP_PCB_PROBLEM) MessageTypeSubStructName() string {
	return "PPP_PCB_PROBLEM"
}

// PPP_PCB_PROBLEM_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPP_PCB_PROBLEM_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPP_PCB_PROBLEM_OBSERVATION) MessageTypeSubStructName() string {
	return "PPP_PCB_PROBLEM_OBSERVATION"
}

// PPP_PCB_PROBLEM_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPP_PCB_PROBLEM_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPP_PCB_PROBLEM_ROLE) MessageTypeSubStructName() string {
	return "PPP_PCB_PROBLEM_ROLE"
}

// PPR_PC1_CHOICE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPR_PC1_CHOICE struct {
	// Only one of the following fields will be set.
	OBR   *OBR `hl7:"true,OBR"`
	RXO   *RXO `hl7:"true,RXO"`
	Other []interface{}
}

func (s *PPR_PC1_CHOICE) MessageTypeSubStructName() string {
	return "PPR_PC1_CHOICE"
}

// PPR_PC1 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPR_PC1 struct {
	MSH           *MSH                   `hl7:"true,MSH"`
	PID           *PID                   `hl7:"true,PID"`
	PATIENT_VISIT *PPR_PC1_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	PROBLEM       []PPR_PC1_PROBLEM      `hl7:"true,PROBLEM"`
	Other         []interface{}
}

func (s *PPR_PC1) MessageTypeName() string {
	return "PPR_PC1"
}

// PPR_PC1_GOAL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPR_PC1_GOAL struct {
	GOL              *GOL                       `hl7:"true,GOL"`
	NTE              []NTE                      `hl7:"false,NTE"`
	VAR              []VAR                      `hl7:"false,VAR"`
	GOAL_ROLE        []PPR_PC1_GOAL_ROLE        `hl7:"false,GOAL_ROLE"`
	GOAL_OBSERVATION []PPR_PC1_GOAL_OBSERVATION `hl7:"false,GOAL_OBSERVATION"`
	Other            []interface{}
}

func (s *PPR_PC1_GOAL) MessageTypeSubStructName() string {
	return "PPR_PC1_GOAL"
}

// PPR_PC1_GOAL_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPR_PC1_GOAL_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPR_PC1_GOAL_OBSERVATION) MessageTypeSubStructName() string {
	return "PPR_PC1_GOAL_OBSERVATION"
}

// PPR_PC1_GOAL_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPR_PC1_GOAL_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPR_PC1_GOAL_ROLE) MessageTypeSubStructName() string {
	return "PPR_PC1_GOAL_ROLE"
}

// PPR_PC1_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPR_PC1_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *PPR_PC1_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other        []interface{}
}

func (s *PPR_PC1_ORDER) MessageTypeSubStructName() string {
	return "PPR_PC1_ORDER"
}

// PPR_PC1_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPR_PC1_ORDER_DETAIL struct {
	CHOICE            *PPR_PC1_CHOICE             `hl7:"true,CHOICE"`
	NTE               []NTE                       `hl7:"false,NTE"`
	VAR               []VAR                       `hl7:"false,VAR"`
	ORDER_OBSERVATION []PPR_PC1_ORDER_OBSERVATION `hl7:"false,ORDER_OBSERVATION"`
	Other             []interface{}
}

func (s *PPR_PC1_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "PPR_PC1_ORDER_DETAIL"
}

// PPR_PC1_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPR_PC1_ORDER_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPR_PC1_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "PPR_PC1_ORDER_OBSERVATION"
}

// PPR_PC1_PATHWAY represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPR_PC1_PATHWAY struct {
	PTH   *PTH  `hl7:"true,PTH"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPR_PC1_PATHWAY) MessageTypeSubStructName() string {
	return "PPR_PC1_PATHWAY"
}

// PPR_PC1_PATHWAY_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type PPR_PC1_PATHWAY_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPR_PC1_PATHWAY_OBSERVATION) MessageTypeSubStructName() string {
	return "PPR_PC1_PATHWAY_OBSERVATION"
}

// PPR_PC1_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPR_PC1_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PPR_PC1_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "PPR_PC1_PATIENT_VISIT"
}

// PPR_PC1_PROBLEM represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPR_PC1_PROBLEM struct {
	PRB                 *PRB                          `hl7:"true,PRB"`
	NTE                 []NTE                         `hl7:"false,NTE"`
	VAR                 []VAR                         `hl7:"false,VAR"`
	PROBLEM_ROLE        []PPR_PC1_PROBLEM_ROLE        `hl7:"false,PROBLEM_ROLE"`
	PATHWAY             []PPR_PC1_PATHWAY             `hl7:"false,PATHWAY"`
	PROBLEM_OBSERVATION []PPR_PC1_PROBLEM_OBSERVATION `hl7:"false,PROBLEM_OBSERVATION"`
	GOAL                []PPR_PC1_GOAL                `hl7:"false,GOAL"`
	ORDER               []PPR_PC1_ORDER               `hl7:"false,ORDER"`
	Other               []interface{}
}

func (s *PPR_PC1_PROBLEM) MessageTypeSubStructName() string {
	return "PPR_PC1_PROBLEM"
}

// PPR_PC1_PROBLEM_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPR_PC1_PROBLEM_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPR_PC1_PROBLEM_OBSERVATION) MessageTypeSubStructName() string {
	return "PPR_PC1_PROBLEM_OBSERVATION"
}

// PPR_PC1_PROBLEM_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPR_PC1_PROBLEM_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPR_PC1_PROBLEM_ROLE) MessageTypeSubStructName() string {
	return "PPR_PC1_PROBLEM_ROLE"
}

// PPT_PCL_CHOICE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPT_PCL_CHOICE struct {
	// Only one of the following fields will be set.
	OBR   *OBR `hl7:"true,OBR"`
	RXO   *RXO `hl7:"true,RXO"`
	Other []interface{}
}

func (s *PPT_PCL_CHOICE) MessageTypeSubStructName() string {
	return "PPT_PCL_CHOICE"
}

// PPT_PCL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPT_PCL struct {
	MSH     *MSH              `hl7:"true,MSH"`
	MSA     *MSA              `hl7:"true,MSA"`
	ERR     *ERR              `hl7:"false,ERR"`
	QAK     *QAK              `hl7:"false,QAK"`
	QRD     *QRD              `hl7:"true,QRD"`
	PATIENT []PPT_PCL_PATIENT `hl7:"true,PATIENT"`
	Other   []interface{}
}

func (s *PPT_PCL) MessageTypeName() string {
	return "PPT_PCL"
}

// PPT_PCL_GOAL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPT_PCL_GOAL struct {
	GOL              *GOL                       `hl7:"true,GOL"`
	NTE              []NTE                      `hl7:"false,NTE"`
	VAR              []VAR                      `hl7:"false,VAR"`
	GOAL_ROLE        []PPT_PCL_GOAL_ROLE        `hl7:"false,GOAL_ROLE"`
	GOAL_OBSERVATION []PPT_PCL_GOAL_OBSERVATION `hl7:"false,GOAL_OBSERVATION"`
	PROBLEM          []PPT_PCL_PROBLEM          `hl7:"false,PROBLEM"`
	ORDER            []PPT_PCL_ORDER            `hl7:"false,ORDER"`
	Other            []interface{}
}

func (s *PPT_PCL_GOAL) MessageTypeSubStructName() string {
	return "PPT_PCL_GOAL"
}

// PPT_PCL_GOAL_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPT_PCL_GOAL_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPT_PCL_GOAL_OBSERVATION) MessageTypeSubStructName() string {
	return "PPT_PCL_GOAL_OBSERVATION"
}

// PPT_PCL_GOAL_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPT_PCL_GOAL_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPT_PCL_GOAL_ROLE) MessageTypeSubStructName() string {
	return "PPT_PCL_GOAL_ROLE"
}

// PPT_PCL_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPT_PCL_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *PPT_PCL_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other        []interface{}
}

func (s *PPT_PCL_ORDER) MessageTypeSubStructName() string {
	return "PPT_PCL_ORDER"
}

// PPT_PCL_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPT_PCL_ORDER_DETAIL struct {
	CHOICE            *PPT_PCL_CHOICE             `hl7:"true,CHOICE"`
	NTE               []NTE                       `hl7:"false,NTE"`
	VAR               []VAR                       `hl7:"false,VAR"`
	ORDER_OBSERVATION []PPT_PCL_ORDER_OBSERVATION `hl7:"false,ORDER_OBSERVATION"`
	Other             []interface{}
}

func (s *PPT_PCL_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "PPT_PCL_ORDER_DETAIL"
}

// PPT_PCL_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPT_PCL_ORDER_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPT_PCL_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "PPT_PCL_ORDER_OBSERVATION"
}

// PPT_PCL_PATHWAY represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPT_PCL_PATHWAY struct {
	PTH          *PTH                   `hl7:"true,PTH"`
	NTE          []NTE                  `hl7:"false,NTE"`
	VAR          []VAR                  `hl7:"false,VAR"`
	PATHWAY_ROLE []PPT_PCL_PATHWAY_ROLE `hl7:"false,PATHWAY_ROLE"`
	GOAL         []PPT_PCL_GOAL         `hl7:"false,GOAL"`
	Other        []interface{}
}

func (s *PPT_PCL_PATHWAY) MessageTypeSubStructName() string {
	return "PPT_PCL_PATHWAY"
}

// PPT_PCL_PATHWAY_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPT_PCL_PATHWAY_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPT_PCL_PATHWAY_ROLE) MessageTypeSubStructName() string {
	return "PPT_PCL_PATHWAY_ROLE"
}

// PPT_PCL_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPT_PCL_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PATIENT_VISIT *PPT_PCL_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	PATHWAY       []PPT_PCL_PATHWAY      `hl7:"true,PATHWAY"`
	Other         []interface{}
}

func (s *PPT_PCL_PATIENT) MessageTypeSubStructName() string {
	return "PPT_PCL_PATIENT"
}

// PPT_PCL_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPT_PCL_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PPT_PCL_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "PPT_PCL_PATIENT_VISIT"
}

// PPT_PCL_PROBLEM represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPT_PCL_PROBLEM struct {
	PRB                 *PRB                          `hl7:"true,PRB"`
	NTE                 []NTE                         `hl7:"false,NTE"`
	VAR                 []VAR                         `hl7:"false,VAR"`
	PROBLEM_ROLE        []PPT_PCL_PROBLEM_ROLE        `hl7:"false,PROBLEM_ROLE"`
	PROBLEM_OBSERVATION []PPT_PCL_PROBLEM_OBSERVATION `hl7:"false,PROBLEM_OBSERVATION"`
	Other               []interface{}
}

func (s *PPT_PCL_PROBLEM) MessageTypeSubStructName() string {
	return "PPT_PCL_PROBLEM"
}

// PPT_PCL_PROBLEM_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPT_PCL_PROBLEM_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPT_PCL_PROBLEM_OBSERVATION) MessageTypeSubStructName() string {
	return "PPT_PCL_PROBLEM_OBSERVATION"
}

// PPT_PCL_PROBLEM_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPT_PCL_PROBLEM_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPT_PCL_PROBLEM_ROLE) MessageTypeSubStructName() string {
	return "PPT_PCL_PROBLEM_ROLE"
}

// PPV_PCA_CHOICE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPV_PCA_CHOICE struct {
	// Only one of the following fields will be set.
	OBR   *OBR `hl7:"true,OBR"`
	RXO   *RXO `hl7:"true,RXO"`
	Other []interface{}
}

func (s *PPV_PCA_CHOICE) MessageTypeSubStructName() string {
	return "PPV_PCA_CHOICE"
}

// PPV_PCA represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPV_PCA struct {
	MSH     *MSH              `hl7:"true,MSH"`
	MSA     *MSA              `hl7:"true,MSA"`
	ERR     *ERR              `hl7:"false,ERR"`
	QAK     *QAK              `hl7:"false,QAK"`
	QRD     *QRD              `hl7:"true,QRD"`
	PATIENT []PPV_PCA_PATIENT `hl7:"true,PATIENT"`
	Other   []interface{}
}

func (s *PPV_PCA) MessageTypeName() string {
	return "PPV_PCA"
}

// PPV_PCA_GOAL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPV_PCA_GOAL struct {
	GOL              *GOL                       `hl7:"true,GOL"`
	NTE              []NTE                      `hl7:"false,NTE"`
	VAR              []VAR                      `hl7:"false,VAR"`
	GOAL_ROLE        []PPV_PCA_GOAL_ROLE        `hl7:"false,GOAL_ROLE"`
	GOAL_PATHWAY     []PPV_PCA_GOAL_PATHWAY     `hl7:"false,GOAL_PATHWAY"`
	GOAL_OBSERVATION []PPV_PCA_GOAL_OBSERVATION `hl7:"false,GOAL_OBSERVATION"`
	PROBLEM          []PPV_PCA_PROBLEM          `hl7:"false,PROBLEM"`
	ORDER            []PPV_PCA_ORDER            `hl7:"false,ORDER"`
	Other            []interface{}
}

func (s *PPV_PCA_GOAL) MessageTypeSubStructName() string {
	return "PPV_PCA_GOAL"
}

// PPV_PCA_GOAL_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPV_PCA_GOAL_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPV_PCA_GOAL_OBSERVATION) MessageTypeSubStructName() string {
	return "PPV_PCA_GOAL_OBSERVATION"
}

// PPV_PCA_GOAL_PATHWAY represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPV_PCA_GOAL_PATHWAY struct {
	PTH   *PTH  `hl7:"true,PTH"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPV_PCA_GOAL_PATHWAY) MessageTypeSubStructName() string {
	return "PPV_PCA_GOAL_PATHWAY"
}

// PPV_PCA_GOAL_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPV_PCA_GOAL_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPV_PCA_GOAL_ROLE) MessageTypeSubStructName() string {
	return "PPV_PCA_GOAL_ROLE"
}

// PPV_PCA_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPV_PCA_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *PPV_PCA_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other        []interface{}
}

func (s *PPV_PCA_ORDER) MessageTypeSubStructName() string {
	return "PPV_PCA_ORDER"
}

// PPV_PCA_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPV_PCA_ORDER_DETAIL struct {
	CHOICE            *PPV_PCA_CHOICE             `hl7:"true,CHOICE"`
	NTE               []NTE                       `hl7:"false,NTE"`
	VAR               []VAR                       `hl7:"false,VAR"`
	ORDER_OBSERVATION []PPV_PCA_ORDER_OBSERVATION `hl7:"false,ORDER_OBSERVATION"`
	Other             []interface{}
}

func (s *PPV_PCA_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "PPV_PCA_ORDER_DETAIL"
}

// PPV_PCA_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPV_PCA_ORDER_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPV_PCA_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "PPV_PCA_ORDER_OBSERVATION"
}

// PPV_PCA_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPV_PCA_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PATIENT_VISIT *PPV_PCA_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	GOAL          []PPV_PCA_GOAL         `hl7:"true,GOAL"`
	Other         []interface{}
}

func (s *PPV_PCA_PATIENT) MessageTypeSubStructName() string {
	return "PPV_PCA_PATIENT"
}

// PPV_PCA_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPV_PCA_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PPV_PCA_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "PPV_PCA_PATIENT_VISIT"
}

// PPV_PCA_PROBLEM represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPV_PCA_PROBLEM struct {
	PRB                 *PRB                          `hl7:"true,PRB"`
	NTE                 []NTE                         `hl7:"false,NTE"`
	VAR                 []VAR                         `hl7:"false,VAR"`
	PROBLEM_ROLE        []PPV_PCA_PROBLEM_ROLE        `hl7:"false,PROBLEM_ROLE"`
	PROBLEM_OBSERVATION []PPV_PCA_PROBLEM_OBSERVATION `hl7:"false,PROBLEM_OBSERVATION"`
	Other               []interface{}
}

func (s *PPV_PCA_PROBLEM) MessageTypeSubStructName() string {
	return "PPV_PCA_PROBLEM"
}

// PPV_PCA_PROBLEM_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPV_PCA_PROBLEM_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PPV_PCA_PROBLEM_OBSERVATION) MessageTypeSubStructName() string {
	return "PPV_PCA_PROBLEM_OBSERVATION"
}

// PPV_PCA_PROBLEM_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PPV_PCA_PROBLEM_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PPV_PCA_PROBLEM_ROLE) MessageTypeSubStructName() string {
	return "PPV_PCA_PROBLEM_ROLE"
}

// PRR_PC5_CHOICE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PRR_PC5_CHOICE struct {
	// Only one of the following fields will be set.
	OBR   *OBR `hl7:"true,OBR"`
	RXO   *RXO `hl7:"true,RXO"`
	Other []interface{}
}

func (s *PRR_PC5_CHOICE) MessageTypeSubStructName() string {
	return "PRR_PC5_CHOICE"
}

// PRR_PC5 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PRR_PC5 struct {
	MSH     *MSH              `hl7:"true,MSH"`
	MSA     *MSA              `hl7:"true,MSA"`
	ERR     *ERR              `hl7:"false,ERR"`
	QAK     *QAK              `hl7:"false,QAK"`
	QRD     *QRD              `hl7:"true,QRD"`
	PATIENT []PRR_PC5_PATIENT `hl7:"true,PATIENT"`
	Other   []interface{}
}

func (s *PRR_PC5) MessageTypeName() string {
	return "PRR_PC5"
}

// PRR_PC5_GOAL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PRR_PC5_GOAL struct {
	GOL              *GOL                       `hl7:"true,GOL"`
	NTE              []NTE                      `hl7:"false,NTE"`
	VAR              []VAR                      `hl7:"false,VAR"`
	GOAL_ROLE        []PRR_PC5_GOAL_ROLE        `hl7:"false,GOAL_ROLE"`
	GOAL_OBSERVATION []PRR_PC5_GOAL_OBSERVATION `hl7:"false,GOAL_OBSERVATION"`
	Other            []interface{}
}

func (s *PRR_PC5_GOAL) MessageTypeSubStructName() string {
	return "PRR_PC5_GOAL"
}

// PRR_PC5_GOAL_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PRR_PC5_GOAL_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PRR_PC5_GOAL_OBSERVATION) MessageTypeSubStructName() string {
	return "PRR_PC5_GOAL_OBSERVATION"
}

// PRR_PC5_GOAL_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PRR_PC5_GOAL_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PRR_PC5_GOAL_ROLE) MessageTypeSubStructName() string {
	return "PRR_PC5_GOAL_ROLE"
}

// PRR_PC5_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PRR_PC5_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *PRR_PC5_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other        []interface{}
}

func (s *PRR_PC5_ORDER) MessageTypeSubStructName() string {
	return "PRR_PC5_ORDER"
}

// PRR_PC5_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PRR_PC5_ORDER_DETAIL struct {
	CHOICE            *PRR_PC5_CHOICE             `hl7:"true,CHOICE"`
	NTE               []NTE                       `hl7:"false,NTE"`
	VAR               []VAR                       `hl7:"false,VAR"`
	ORDER_OBSERVATION []PRR_PC5_ORDER_OBSERVATION `hl7:"false,ORDER_OBSERVATION"`
	Other             []interface{}
}

func (s *PRR_PC5_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "PRR_PC5_ORDER_DETAIL"
}

// PRR_PC5_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PRR_PC5_ORDER_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PRR_PC5_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "PRR_PC5_ORDER_OBSERVATION"
}

// PRR_PC5_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PRR_PC5_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PATIENT_VISIT *PRR_PC5_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	PROBLEM       []PRR_PC5_PROBLEM      `hl7:"true,PROBLEM"`
	Other         []interface{}
}

func (s *PRR_PC5_PATIENT) MessageTypeSubStructName() string {
	return "PRR_PC5_PATIENT"
}

// PRR_PC5_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PRR_PC5_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PRR_PC5_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "PRR_PC5_PATIENT_VISIT"
}

// PRR_PC5_PROBLEM represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PRR_PC5_PROBLEM struct {
	PRB                 *PRB                          `hl7:"true,PRB"`
	NTE                 []NTE                         `hl7:"false,NTE"`
	VAR                 []VAR                         `hl7:"false,VAR"`
	PROBLEM_ROLE        []PRR_PC5_PROBLEM_ROLE        `hl7:"false,PROBLEM_ROLE"`
	PROBLEM_PATHWAY     []PRR_PC5_PROBLEM_PATHWAY     `hl7:"false,PROBLEM_PATHWAY"`
	PROBLEM_OBSERVATION []PRR_PC5_PROBLEM_OBSERVATION `hl7:"false,PROBLEM_OBSERVATION"`
	GOAL                []PRR_PC5_GOAL                `hl7:"false,GOAL"`
	ORDER               []PRR_PC5_ORDER               `hl7:"false,ORDER"`
	Other               []interface{}
}

func (s *PRR_PC5_PROBLEM) MessageTypeSubStructName() string {
	return "PRR_PC5_PROBLEM"
}

// PRR_PC5_PROBLEM_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PRR_PC5_PROBLEM_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PRR_PC5_PROBLEM_OBSERVATION) MessageTypeSubStructName() string {
	return "PRR_PC5_PROBLEM_OBSERVATION"
}

// PRR_PC5_PROBLEM_PATHWAY represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PRR_PC5_PROBLEM_PATHWAY struct {
	PTH   *PTH  `hl7:"true,PTH"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PRR_PC5_PROBLEM_PATHWAY) MessageTypeSubStructName() string {
	return "PRR_PC5_PROBLEM_PATHWAY"
}

// PRR_PC5_PROBLEM_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PRR_PC5_PROBLEM_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PRR_PC5_PROBLEM_ROLE) MessageTypeSubStructName() string {
	return "PRR_PC5_PROBLEM_ROLE"
}

// PTR_PCF_CHOICE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PTR_PCF_CHOICE struct {
	// Only one of the following fields will be set.
	OBR   *OBR `hl7:"true,OBR"`
	RXO   *RXO `hl7:"true,RXO"`
	Other []interface{}
}

func (s *PTR_PCF_CHOICE) MessageTypeSubStructName() string {
	return "PTR_PCF_CHOICE"
}

// PTR_PCF represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PTR_PCF struct {
	MSH     *MSH              `hl7:"true,MSH"`
	MSA     *MSA              `hl7:"true,MSA"`
	ERR     *ERR              `hl7:"false,ERR"`
	QAK     *QAK              `hl7:"false,QAK"`
	QRD     *QRD              `hl7:"true,QRD"`
	PATIENT []PTR_PCF_PATIENT `hl7:"true,PATIENT"`
	Other   []interface{}
}

func (s *PTR_PCF) MessageTypeName() string {
	return "PTR_PCF"
}

// PTR_PCF_GOAL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PTR_PCF_GOAL struct {
	GOL              *GOL                       `hl7:"true,GOL"`
	NTE              []NTE                      `hl7:"false,NTE"`
	VAR              []VAR                      `hl7:"false,VAR"`
	GOAL_ROLE        []PTR_PCF_GOAL_ROLE        `hl7:"false,GOAL_ROLE"`
	GOAL_OBSERVATION []PTR_PCF_GOAL_OBSERVATION `hl7:"false,GOAL_OBSERVATION"`
	Other            []interface{}
}

func (s *PTR_PCF_GOAL) MessageTypeSubStructName() string {
	return "PTR_PCF_GOAL"
}

// PTR_PCF_GOAL_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PTR_PCF_GOAL_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PTR_PCF_GOAL_OBSERVATION) MessageTypeSubStructName() string {
	return "PTR_PCF_GOAL_OBSERVATION"
}

// PTR_PCF_GOAL_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PTR_PCF_GOAL_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PTR_PCF_GOAL_ROLE) MessageTypeSubStructName() string {
	return "PTR_PCF_GOAL_ROLE"
}

// PTR_PCF_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PTR_PCF_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *PTR_PCF_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other        []interface{}
}

func (s *PTR_PCF_ORDER) MessageTypeSubStructName() string {
	return "PTR_PCF_ORDER"
}

// PTR_PCF_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PTR_PCF_ORDER_DETAIL struct {
	CHOICE            *PTR_PCF_CHOICE             `hl7:"true,CHOICE"`
	NTE               []NTE                       `hl7:"false,NTE"`
	VAR               []VAR                       `hl7:"false,VAR"`
	ORDER_OBSERVATION []PTR_PCF_ORDER_OBSERVATION `hl7:"false,ORDER_OBSERVATION"`
	Other             []interface{}
}

func (s *PTR_PCF_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "PTR_PCF_ORDER_DETAIL"
}

// PTR_PCF_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PTR_PCF_ORDER_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PTR_PCF_ORDER_OBSERVATION) MessageTypeSubStructName() string {
	return "PTR_PCF_ORDER_OBSERVATION"
}

// PTR_PCF_PATHWAY represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PTR_PCF_PATHWAY struct {
	PTH          *PTH                   `hl7:"true,PTH"`
	NTE          []NTE                  `hl7:"false,NTE"`
	VAR          []VAR                  `hl7:"false,VAR"`
	PATHWAY_ROLE []PTR_PCF_PATHWAY_ROLE `hl7:"false,PATHWAY_ROLE"`
	PROBLEM      []PTR_PCF_PROBLEM      `hl7:"false,PROBLEM"`
	Other        []interface{}
}

func (s *PTR_PCF_PATHWAY) MessageTypeSubStructName() string {
	return "PTR_PCF_PATHWAY"
}

// PTR_PCF_PATHWAY_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PTR_PCF_PATHWAY_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PTR_PCF_PATHWAY_ROLE) MessageTypeSubStructName() string {
	return "PTR_PCF_PATHWAY_ROLE"
}

// PTR_PCF_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PTR_PCF_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PATIENT_VISIT *PTR_PCF_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	PATHWAY       []PTR_PCF_PATHWAY      `hl7:"true,PATHWAY"`
	Other         []interface{}
}

func (s *PTR_PCF_PATIENT) MessageTypeSubStructName() string {
	return "PTR_PCF_PATIENT"
}

// PTR_PCF_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PTR_PCF_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *PTR_PCF_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "PTR_PCF_PATIENT_VISIT"
}

// PTR_PCF_PROBLEM represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PTR_PCF_PROBLEM struct {
	PRB                 *PRB                          `hl7:"true,PRB"`
	NTE                 []NTE                         `hl7:"false,NTE"`
	VAR                 []VAR                         `hl7:"false,VAR"`
	PROBLEM_ROLE        []PTR_PCF_PROBLEM_ROLE        `hl7:"false,PROBLEM_ROLE"`
	PROBLEM_OBSERVATION []PTR_PCF_PROBLEM_OBSERVATION `hl7:"false,PROBLEM_OBSERVATION"`
	GOAL                []PTR_PCF_GOAL                `hl7:"false,GOAL"`
	ORDER               []PTR_PCF_ORDER               `hl7:"false,ORDER"`
	Other               []interface{}
}

func (s *PTR_PCF_PROBLEM) MessageTypeSubStructName() string {
	return "PTR_PCF_PROBLEM"
}

// PTR_PCF_PROBLEM_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PTR_PCF_PROBLEM_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *PTR_PCF_PROBLEM_OBSERVATION) MessageTypeSubStructName() string {
	return "PTR_PCF_PROBLEM_OBSERVATION"
}

// PTR_PCF_PROBLEM_ROLE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type PTR_PCF_PROBLEM_ROLE struct {
	ROL   *ROL  `hl7:"true,ROL"`
	VAR   []VAR `hl7:"false,VAR"`
	Other []interface{}
}

func (s *PTR_PCF_PROBLEM_ROLE) MessageTypeSubStructName() string {
	return "PTR_PCF_PROBLEM_ROLE"
}

// QBP_K13 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type QBP_K13 struct {
	MSH            *MSH                    `hl7:"true,MSH"`
	MSA            *MSA                    `hl7:"true,MSA"`
	ERR            *ERR                    `hl7:"false,ERR"`
	QAK            *QAK                    `hl7:"true,QAK"`
	QPD            *QPD                    `hl7:"true,QPD"`
	ROW_DEFINITION *QBP_K13_ROW_DEFINITION `hl7:"false,ROW_DEFINITION"`
	DSC            *DSC                    `hl7:"false,DSC"`
	Other          []interface{}
}

func (s *QBP_K13) MessageTypeName() string {
	return "QBP_K13"
}

// QBP_K13_ROW_DEFINITION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type QBP_K13_ROW_DEFINITION struct {
	RDF   *RDF  `hl7:"true,RDF"`
	RDT   []RDT `hl7:"false,RDT"`
	Other []interface{}
}

func (s *QBP_K13_ROW_DEFINITION) MessageTypeSubStructName() string {
	return "QBP_K13_ROW_DEFINITION"
}

// QBP_Q11 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type QBP_Q11 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QPD   *QPD `hl7:"true,QPD"`
	RCP   *RCP `hl7:"true,RCP"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *QBP_Q11) MessageTypeName() string {
	return "QBP_Q11"
}

// QBP_Q13 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type QBP_Q13 struct {
	MSH   *MSH         `hl7:"true,MSH"`
	QPD   *QPD         `hl7:"true,QPD"`
	QBP   *QBP_Q13_QBP `hl7:"false,QBP"`
	RDF   *RDF         `hl7:"false,RDF"`
	RCP   *RCP         `hl7:"true,RCP"`
	DSC   *DSC         `hl7:"false,DSC"`
	Other []interface{}
}

func (s *QBP_Q13) MessageTypeName() string {
	return "QBP_Q13"
}

// QBP_Q13_QBP represents the corresponding HL7 message type.
// Definition from HL7 2.4
type QBP_Q13_QBP struct {
	// Missing: anyZSegment
	Other []interface{}
}

func (s *QBP_Q13_QBP) MessageTypeSubStructName() string {
	return "QBP_Q13_QBP"
}

// QBP_Q15 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type QBP_Q15 struct {
	MSH *MSH `hl7:"true,MSH"`
	QPD *QPD `hl7:"true,QPD"`
	// Missing: anyZSegment
	RCP   *RCP `hl7:"true,RCP"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *QBP_Q15) MessageTypeName() string {
	return "QBP_Q15"
}

// QBP_Q21 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type QBP_Q21 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QPD   *QPD `hl7:"true,QPD"`
	RCP   *RCP `hl7:"true,RCP"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *QBP_Q21) MessageTypeName() string {
	return "QBP_Q21"
}

// QBP_Qnn represents the corresponding HL7 message type.
// Definition from HL7 2.4
type QBP_Qnn struct {
	MSH   *MSH `hl7:"true,MSH"`
	QPD   *QPD `hl7:"true,QPD"`
	RDF   *RDF `hl7:"false,RDF"`
	RCP   *RCP `hl7:"true,RCP"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *QBP_Qnn) MessageTypeName() string {
	return "QBP_Qnn"
}

// QBP_Z73 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type QBP_Z73 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QPD   *QPD `hl7:"true,QPD"`
	RCP   *RCP `hl7:"true,RCP"`
	Other []interface{}
}

func (s *QBP_Z73) MessageTypeName() string {
	return "QBP_Z73"
}

// QCK_Q02 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type QCK_Q02 struct {
	MSH   *MSH `hl7:"true,MSH"`
	MSA   *MSA `hl7:"true,MSA"`
	ERR   *ERR `hl7:"false,ERR"`
	QAK   *QAK `hl7:"false,QAK"`
	Other []interface{}
}

func (s *QCK_Q02) MessageTypeName() string {
	return "QCK_Q02"
}

// QCN_J01 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type QCN_J01 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QID   *QID `hl7:"true,QID"`
	Other []interface{}
}

func (s *QCN_J01) MessageTypeName() string {
	return "QCN_J01"
}

// QRY_A19 represents the corresponding HL7 message type.
// Definition from HL7 2.4
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

// QRY_PC4 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type QRY_PC4 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QRD   *QRD `hl7:"true,QRD"`
	QRF   *QRF `hl7:"false,QRF"`
	Other []interface{}
}

func (s *QRY_PC4) MessageTypeName() string {
	return "QRY_PC4"
}

// QRY_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.4
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
// Definition from HL7 2.4
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
// Definition from HL7 2.4
type QRY_R02 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QRD   *QRD `hl7:"true,QRD"`
	QRF   *QRF `hl7:"true,QRF"`
	Other []interface{}
}

func (s *QRY_R02) MessageTypeName() string {
	return "QRY_R02"
}

// QRY_T12 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type QRY_T12 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QRD   *QRD `hl7:"true,QRD"`
	QRF   *QRF `hl7:"false,QRF"`
	Other []interface{}
}

func (s *QRY_T12) MessageTypeName() string {
	return "QRY_T12"
}

// QSB_Q16 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type QSB_Q16 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QPD   *QPD `hl7:"true,QPD"`
	RCP   *RCP `hl7:"true,RCP"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *QSB_Q16) MessageTypeName() string {
	return "QSB_Q16"
}

// QVR_Q17 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type QVR_Q17 struct {
	MSH *MSH `hl7:"true,MSH"`
	QPD *QPD `hl7:"true,QPD"`
	// Missing: anyZSegment
	RCP   *RCP `hl7:"true,RCP"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *QVR_Q17) MessageTypeName() string {
	return "QVR_Q17"
}

// RAR_RAR represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RAR_RAR struct {
	MSH        *MSH                 `hl7:"true,MSH"`
	MSA        *MSA                 `hl7:"true,MSA"`
	ERR        *ERR                 `hl7:"false,ERR"`
	DEFINITION []RAR_RAR_DEFINITION `hl7:"true,DEFINITION"`
	DSC        *DSC                 `hl7:"false,DSC"`
	Other      []interface{}
}

func (s *RAR_RAR) MessageTypeName() string {
	return "RAR_RAR"
}

// RAR_RAR_DEFINITION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RAR_RAR_DEFINITION struct {
	QRD     *QRD             `hl7:"true,QRD"`
	QRF     *QRF             `hl7:"false,QRF"`
	PATIENT *RAR_RAR_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RAR_RAR_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RAR_RAR_DEFINITION) MessageTypeSubStructName() string {
	return "RAR_RAR_DEFINITION"
}

// RAR_RAR_ENCODING represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RAR_RAR_ENCODING struct {
	RXE   *RXE  `hl7:"true,RXE"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RAR_RAR_ENCODING) MessageTypeSubStructName() string {
	return "RAR_RAR_ENCODING"
}

// RAR_RAR_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RAR_RAR_ORDER struct {
	ORC      *ORC              `hl7:"true,ORC"`
	ENCODING *RAR_RAR_ENCODING `hl7:"false,ENCODING"`
	RXA      []RXA             `hl7:"true,RXA"`
	RXR      *RXR              `hl7:"true,RXR"`
	Other    []interface{}
}

func (s *RAR_RAR_ORDER) MessageTypeSubStructName() string {
	return "RAR_RAR_ORDER"
}

// RAR_RAR_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RAR_RAR_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RAR_RAR_PATIENT) MessageTypeSubStructName() string {
	return "RAR_RAR_PATIENT"
}

// RAS_O01_COMPONENTS represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RAS_O01_COMPONENTS struct {
	RXC   []RXC `hl7:"true,RXC"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RAS_O01_COMPONENTS) MessageTypeSubStructName() string {
	return "RAS_O01_COMPONENTS"
}

// RAS_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RAS_O01 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	NTE     []NTE            `hl7:"false,NTE"`
	PATIENT *RAS_O01_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RAS_O01_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RAS_O01) MessageTypeName() string {
	return "RAS_O01"
}

// RAS_O01_ENCODING represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RAS_O01_ENCODING struct {
	RXE   *RXE  `hl7:"true,RXE"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RAS_O01_ENCODING) MessageTypeSubStructName() string {
	return "RAS_O01_ENCODING"
}

// RAS_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RAS_O01_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RAS_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "RAS_O01_OBSERVATION"
}

// RAS_O01_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RAS_O01_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *RAS_O01_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	ENCODING     *RAS_O01_ENCODING     `hl7:"false,ENCODING"`
	RXA          []RXA                 `hl7:"true,RXA"`
	RXR          *RXR                  `hl7:"true,RXR"`
	OBSERVATION  []RAS_O01_OBSERVATION `hl7:"false,OBSERVATION"`
	CTI          []CTI                 `hl7:"false,CTI"`
	Other        []interface{}
}

func (s *RAS_O01_ORDER) MessageTypeSubStructName() string {
	return "RAS_O01_ORDER"
}

// RAS_O01_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RAS_O01_ORDER_DETAIL struct {
	RXO                     *RXO                             `hl7:"true,RXO"`
	ORDER_DETAIL_SUPPLEMENT *RAS_O01_ORDER_DETAIL_SUPPLEMENT `hl7:"false,ORDER_DETAIL_SUPPLEMENT"`
	Other                   []interface{}
}

func (s *RAS_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RAS_O01_ORDER_DETAIL"
}

// RAS_O01_ORDER_DETAIL_SUPPLEMENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RAS_O01_ORDER_DETAIL_SUPPLEMENT struct {
	NTE        []NTE               `hl7:"true,NTE"`
	RXR        []RXR               `hl7:"true,RXR"`
	COMPONENTS *RAS_O01_COMPONENTS `hl7:"false,COMPONENTS"`
	Other      []interface{}
}

func (s *RAS_O01_ORDER_DETAIL_SUPPLEMENT) MessageTypeSubStructName() string {
	return "RAS_O01_ORDER_DETAIL_SUPPLEMENT"
}

// RAS_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RAS_O01_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NTE           []NTE                  `hl7:"false,NTE"`
	AL1           []AL1                  `hl7:"false,AL1"`
	PATIENT_VISIT *RAS_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	Other         []interface{}
}

func (s *RAS_O01_PATIENT) MessageTypeSubStructName() string {
	return "RAS_O01_PATIENT"
}

// RAS_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RAS_O01_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RAS_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "RAS_O01_PATIENT_VISIT"
}

// RAS_O17_COMPONENTS represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RAS_O17_COMPONENTS struct {
	RXC   []RXC `hl7:"true,RXC"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RAS_O17_COMPONENTS) MessageTypeSubStructName() string {
	return "RAS_O17_COMPONENTS"
}

// RAS_O17 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RAS_O17 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	NTE     []NTE            `hl7:"false,NTE"`
	PATIENT *RAS_O17_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RAS_O17_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RAS_O17) MessageTypeName() string {
	return "RAS_O17"
}

// RAS_O17_ENCODING represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RAS_O17_ENCODING struct {
	RXE   *RXE  `hl7:"true,RXE"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RAS_O17_ENCODING) MessageTypeSubStructName() string {
	return "RAS_O17_ENCODING"
}

// RAS_O17_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RAS_O17_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RAS_O17_OBSERVATION) MessageTypeSubStructName() string {
	return "RAS_O17_OBSERVATION"
}

// RAS_O17_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RAS_O17_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *RAS_O17_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	ENCODING     *RAS_O17_ENCODING     `hl7:"false,ENCODING"`
	RXA          []RXA                 `hl7:"true,RXA"`
	RXR          *RXR                  `hl7:"true,RXR"`
	OBSERVATION  []RAS_O17_OBSERVATION `hl7:"false,OBSERVATION"`
	CTI          []CTI                 `hl7:"false,CTI"`
	Other        []interface{}
}

func (s *RAS_O17_ORDER) MessageTypeSubStructName() string {
	return "RAS_O17_ORDER"
}

// RAS_O17_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RAS_O17_ORDER_DETAIL struct {
	RXO                     *RXO                             `hl7:"true,RXO"`
	ORDER_DETAIL_SUPPLEMENT *RAS_O17_ORDER_DETAIL_SUPPLEMENT `hl7:"false,ORDER_DETAIL_SUPPLEMENT"`
	Other                   []interface{}
}

func (s *RAS_O17_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RAS_O17_ORDER_DETAIL"
}

// RAS_O17_ORDER_DETAIL_SUPPLEMENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RAS_O17_ORDER_DETAIL_SUPPLEMENT struct {
	NTE        []NTE               `hl7:"true,NTE"`
	RXR        []RXR               `hl7:"true,RXR"`
	COMPONENTS *RAS_O17_COMPONENTS `hl7:"false,COMPONENTS"`
	Other      []interface{}
}

func (s *RAS_O17_ORDER_DETAIL_SUPPLEMENT) MessageTypeSubStructName() string {
	return "RAS_O17_ORDER_DETAIL_SUPPLEMENT"
}

// RAS_O17_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RAS_O17_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NTE           []NTE                  `hl7:"false,NTE"`
	AL1           []AL1                  `hl7:"false,AL1"`
	PATIENT_VISIT *RAS_O17_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	Other         []interface{}
}

func (s *RAS_O17_PATIENT) MessageTypeSubStructName() string {
	return "RAS_O17_PATIENT"
}

// RAS_O17_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RAS_O17_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RAS_O17_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "RAS_O17_PATIENT_VISIT"
}

// RCI_I05 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RCI_I05 struct {
	MSH         *MSH                  `hl7:"true,MSH"`
	MSA         *MSA                  `hl7:"true,MSA"`
	QRD         *QRD                  `hl7:"true,QRD"`
	QRF         *QRF                  `hl7:"false,QRF"`
	PROVIDER    []RCI_I05_PROVIDER    `hl7:"true,PROVIDER"`
	PID         *PID                  `hl7:"true,PID"`
	DG1         []DG1                 `hl7:"false,DG1"`
	DRG         []DRG                 `hl7:"false,DRG"`
	AL1         []AL1                 `hl7:"false,AL1"`
	OBSERVATION []RCI_I05_OBSERVATION `hl7:"false,OBSERVATION"`
	NTE         []NTE                 `hl7:"false,NTE"`
	Other       []interface{}
}

func (s *RCI_I05) MessageTypeName() string {
	return "RCI_I05"
}

// RCI_I05_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RCI_I05_OBSERVATION struct {
	OBR     *OBR              `hl7:"true,OBR"`
	NTE     []NTE             `hl7:"false,NTE"`
	RESULTS []RCI_I05_RESULTS `hl7:"false,RESULTS"`
	Other   []interface{}
}

func (s *RCI_I05_OBSERVATION) MessageTypeSubStructName() string {
	return "RCI_I05_OBSERVATION"
}

// RCI_I05_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RCI_I05_PROVIDER struct {
	PRD   *PRD  `hl7:"true,PRD"`
	CTD   []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RCI_I05_PROVIDER) MessageTypeSubStructName() string {
	return "RCI_I05_PROVIDER"
}

// RCI_I05_RESULTS represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RCI_I05_RESULTS struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RCI_I05_RESULTS) MessageTypeSubStructName() string {
	return "RCI_I05_RESULTS"
}

// RCL_I06 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RCL_I06 struct {
	MSH      *MSH               `hl7:"true,MSH"`
	MSA      *MSA               `hl7:"true,MSA"`
	QRD      *QRD               `hl7:"true,QRD"`
	QRF      *QRF               `hl7:"false,QRF"`
	PROVIDER []RCL_I06_PROVIDER `hl7:"true,PROVIDER"`
	PID      *PID               `hl7:"true,PID"`
	DG1      []DG1              `hl7:"false,DG1"`
	DRG      []DRG              `hl7:"false,DRG"`
	AL1      []AL1              `hl7:"false,AL1"`
	NTE      []NTE              `hl7:"false,NTE"`
	DSP      []DSP              `hl7:"false,DSP"`
	DSC      *DSC               `hl7:"false,DSC"`
	Other    []interface{}
}

func (s *RCL_I06) MessageTypeName() string {
	return "RCL_I06"
}

// RCL_I06_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RCL_I06_PROVIDER struct {
	PRD   *PRD  `hl7:"true,PRD"`
	CTD   []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RCL_I06_PROVIDER) MessageTypeSubStructName() string {
	return "RCL_I06_PROVIDER"
}

// RDE_O01_COMPONENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDE_O01_COMPONENT struct {
	RXC   []RXC `hl7:"true,RXC"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDE_O01_COMPONENT) MessageTypeSubStructName() string {
	return "RDE_O01_COMPONENT"
}

// RDE_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDE_O01 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	NTE     []NTE            `hl7:"false,NTE"`
	PATIENT *RDE_O01_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RDE_O01_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RDE_O01) MessageTypeName() string {
	return "RDE_O01"
}

// RDE_O01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDE_O01_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *RDE_O01_INSURANCE) MessageTypeSubStructName() string {
	return "RDE_O01_INSURANCE"
}

// RDE_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDE_O01_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDE_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "RDE_O01_OBSERVATION"
}

// RDE_O01_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDE_O01_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *RDE_O01_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	RXE          *RXE                  `hl7:"true,RXE"`
	RXR          []RXR                 `hl7:"true,RXR"`
	RXC          []RXC                 `hl7:"false,RXC"`
	OBSERVATION  []RDE_O01_OBSERVATION `hl7:"false,OBSERVATION"`
	CTI          []CTI                 `hl7:"false,CTI"`
	Other        []interface{}
}

func (s *RDE_O01_ORDER) MessageTypeSubStructName() string {
	return "RDE_O01_ORDER"
}

// RDE_O01_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDE_O01_ORDER_DETAIL struct {
	RXO       *RXO               `hl7:"true,RXO"`
	NTE       []NTE              `hl7:"false,NTE"`
	RXR       []RXR              `hl7:"true,RXR"`
	COMPONENT *RDE_O01_COMPONENT `hl7:"false,COMPONENT"`
	Other     []interface{}
}

func (s *RDE_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RDE_O01_ORDER_DETAIL"
}

// RDE_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDE_O01_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NTE           []NTE                  `hl7:"false,NTE"`
	PATIENT_VISIT *RDE_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE     []RDE_O01_INSURANCE    `hl7:"false,INSURANCE"`
	GT1           *GT1                   `hl7:"false,GT1"`
	AL1           []AL1                  `hl7:"false,AL1"`
	Other         []interface{}
}

func (s *RDE_O01_PATIENT) MessageTypeSubStructName() string {
	return "RDE_O01_PATIENT"
}

// RDE_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDE_O01_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RDE_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "RDE_O01_PATIENT_VISIT"
}

// RDE_O11_COMPONENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDE_O11_COMPONENT struct {
	RXC   []RXC `hl7:"true,RXC"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDE_O11_COMPONENT) MessageTypeSubStructName() string {
	return "RDE_O11_COMPONENT"
}

// RDE_O11 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDE_O11 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	NTE     []NTE            `hl7:"false,NTE"`
	PATIENT *RDE_O11_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RDE_O11_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RDE_O11) MessageTypeName() string {
	return "RDE_O11"
}

// RDE_O11_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDE_O11_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *RDE_O11_INSURANCE) MessageTypeSubStructName() string {
	return "RDE_O11_INSURANCE"
}

// RDE_O11_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDE_O11_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDE_O11_OBSERVATION) MessageTypeSubStructName() string {
	return "RDE_O11_OBSERVATION"
}

// RDE_O11_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDE_O11_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *RDE_O11_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	RXE          *RXE                  `hl7:"true,RXE"`
	RXR          []RXR                 `hl7:"true,RXR"`
	RXC          []RXC                 `hl7:"false,RXC"`
	OBSERVATION  []RDE_O11_OBSERVATION `hl7:"false,OBSERVATION"`
	CTI          []CTI                 `hl7:"false,CTI"`
	Other        []interface{}
}

func (s *RDE_O11_ORDER) MessageTypeSubStructName() string {
	return "RDE_O11_ORDER"
}

// RDE_O11_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDE_O11_ORDER_DETAIL struct {
	RXO       *RXO               `hl7:"true,RXO"`
	NTE       []NTE              `hl7:"false,NTE"`
	RXR       []RXR              `hl7:"true,RXR"`
	COMPONENT *RDE_O11_COMPONENT `hl7:"false,COMPONENT"`
	Other     []interface{}
}

func (s *RDE_O11_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RDE_O11_ORDER_DETAIL"
}

// RDE_O11_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDE_O11_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NTE           []NTE                  `hl7:"false,NTE"`
	PATIENT_VISIT *RDE_O11_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE     []RDE_O11_INSURANCE    `hl7:"false,INSURANCE"`
	GT1           *GT1                   `hl7:"false,GT1"`
	AL1           []AL1                  `hl7:"false,AL1"`
	Other         []interface{}
}

func (s *RDE_O11_PATIENT) MessageTypeSubStructName() string {
	return "RDE_O11_PATIENT"
}

// RDE_O11_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDE_O11_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RDE_O11_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "RDE_O11_PATIENT_VISIT"
}

// RDO_O01_COMPONENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDO_O01_COMPONENT struct {
	RXC   []RXC `hl7:"true,RXC"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDO_O01_COMPONENT) MessageTypeSubStructName() string {
	return "RDO_O01_COMPONENT"
}

// RDO_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDO_O01 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	NTE     []NTE            `hl7:"false,NTE"`
	PATIENT *RDO_O01_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RDO_O01_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RDO_O01) MessageTypeName() string {
	return "RDO_O01"
}

// RDO_O01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDO_O01_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *RDO_O01_INSURANCE) MessageTypeSubStructName() string {
	return "RDO_O01_INSURANCE"
}

// RDO_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDO_O01_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDO_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "RDO_O01_OBSERVATION"
}

// RDO_O01_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDO_O01_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *RDO_O01_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	BLG          *BLG                  `hl7:"false,BLG"`
	Other        []interface{}
}

func (s *RDO_O01_ORDER) MessageTypeSubStructName() string {
	return "RDO_O01_ORDER"
}

// RDO_O01_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDO_O01_ORDER_DETAIL struct {
	RXO         *RXO                  `hl7:"true,RXO"`
	NTE         []NTE                 `hl7:"false,NTE"`
	RXR         []RXR                 `hl7:"true,RXR"`
	COMPONENT   *RDO_O01_COMPONENT    `hl7:"false,COMPONENT"`
	OBSERVATION []RDO_O01_OBSERVATION `hl7:"false,OBSERVATION"`
	Other       []interface{}
}

func (s *RDO_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RDO_O01_ORDER_DETAIL"
}

// RDO_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDO_O01_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NTE           []NTE                  `hl7:"false,NTE"`
	PATIENT_VISIT *RDO_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	INSURANCE     []RDO_O01_INSURANCE    `hl7:"false,INSURANCE"`
	GT1           *GT1                   `hl7:"false,GT1"`
	AL1           []AL1                  `hl7:"false,AL1"`
	Other         []interface{}
}

func (s *RDO_O01_PATIENT) MessageTypeSubStructName() string {
	return "RDO_O01_PATIENT"
}

// RDO_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDO_O01_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RDO_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "RDO_O01_PATIENT_VISIT"
}

// RDR_RDR represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDR_RDR struct {
	MSH        *MSH                 `hl7:"true,MSH"`
	MSA        *MSA                 `hl7:"true,MSA"`
	ERR        *ERR                 `hl7:"false,ERR"`
	DEFINITION []RDR_RDR_DEFINITION `hl7:"true,DEFINITION"`
	DSC        *DSC                 `hl7:"false,DSC"`
	Other      []interface{}
}

func (s *RDR_RDR) MessageTypeName() string {
	return "RDR_RDR"
}

// RDR_RDR_DEFINITION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDR_RDR_DEFINITION struct {
	QRD     *QRD             `hl7:"true,QRD"`
	QRF     *QRF             `hl7:"false,QRF"`
	PATIENT *RDR_RDR_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RDR_RDR_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RDR_RDR_DEFINITION) MessageTypeSubStructName() string {
	return "RDR_RDR_DEFINITION"
}

// RDR_RDR_DISPENSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDR_RDR_DISPENSE struct {
	RXD   *RXD  `hl7:"true,RXD"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RDR_RDR_DISPENSE) MessageTypeSubStructName() string {
	return "RDR_RDR_DISPENSE"
}

// RDR_RDR_ENCODING represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDR_RDR_ENCODING struct {
	RXE   *RXE  `hl7:"true,RXE"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RDR_RDR_ENCODING) MessageTypeSubStructName() string {
	return "RDR_RDR_ENCODING"
}

// RDR_RDR_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDR_RDR_ORDER struct {
	ORC      *ORC               `hl7:"true,ORC"`
	ENCODING *RDR_RDR_ENCODING  `hl7:"false,ENCODING"`
	DISPENSE []RDR_RDR_DISPENSE `hl7:"true,DISPENSE"`
	Other    []interface{}
}

func (s *RDR_RDR_ORDER) MessageTypeSubStructName() string {
	return "RDR_RDR_ORDER"
}

// RDR_RDR_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDR_RDR_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDR_RDR_PATIENT) MessageTypeSubStructName() string {
	return "RDR_RDR_PATIENT"
}

// RDS_O01_COMPONENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDS_O01_COMPONENT struct {
	RXC   []RXC `hl7:"true,RXC"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDS_O01_COMPONENT) MessageTypeSubStructName() string {
	return "RDS_O01_COMPONENT"
}

// RDS_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDS_O01 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	NTE     []NTE            `hl7:"false,NTE"`
	PATIENT *RDS_O01_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RDS_O01_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RDS_O01) MessageTypeName() string {
	return "RDS_O01"
}

// RDS_O01_ENCODING represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDS_O01_ENCODING struct {
	RXE   *RXE  `hl7:"true,RXE"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RDS_O01_ENCODING) MessageTypeSubStructName() string {
	return "RDS_O01_ENCODING"
}

// RDS_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDS_O01_OBSERVATION struct {
	OBX   *OBX  `hl7:"false,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDS_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "RDS_O01_OBSERVATION"
}

// RDS_O01_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDS_O01_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *RDS_O01_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	ENCODING     *RDS_O01_ENCODING     `hl7:"false,ENCODING"`
	RXD          *RXD                  `hl7:"true,RXD"`
	RXR          []RXR                 `hl7:"true,RXR"`
	RXC          []RXC                 `hl7:"false,RXC"`
	OBSERVATION  []RDS_O01_OBSERVATION `hl7:"false,OBSERVATION"`
	Other        []interface{}
}

func (s *RDS_O01_ORDER) MessageTypeSubStructName() string {
	return "RDS_O01_ORDER"
}

// RDS_O01_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDS_O01_ORDER_DETAIL struct {
	RXO                     *RXO                             `hl7:"true,RXO"`
	ORDER_DETAIL_SUPPLEMENT *RDS_O01_ORDER_DETAIL_SUPPLEMENT `hl7:"false,ORDER_DETAIL_SUPPLEMENT"`
	Other                   []interface{}
}

func (s *RDS_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RDS_O01_ORDER_DETAIL"
}

// RDS_O01_ORDER_DETAIL_SUPPLEMENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDS_O01_ORDER_DETAIL_SUPPLEMENT struct {
	NTE       []NTE              `hl7:"true,NTE"`
	RXR       []RXR              `hl7:"true,RXR"`
	COMPONENT *RDS_O01_COMPONENT `hl7:"false,COMPONENT"`
	Other     []interface{}
}

func (s *RDS_O01_ORDER_DETAIL_SUPPLEMENT) MessageTypeSubStructName() string {
	return "RDS_O01_ORDER_DETAIL_SUPPLEMENT"
}

// RDS_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDS_O01_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NTE           []NTE                  `hl7:"false,NTE"`
	AL1           []AL1                  `hl7:"false,AL1"`
	PATIENT_VISIT *RDS_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	Other         []interface{}
}

func (s *RDS_O01_PATIENT) MessageTypeSubStructName() string {
	return "RDS_O01_PATIENT"
}

// RDS_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RDS_O01_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RDS_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "RDS_O01_PATIENT_VISIT"
}

// RDS_O13_COMPONENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDS_O13_COMPONENT struct {
	RXC   []RXC `hl7:"true,RXC"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDS_O13_COMPONENT) MessageTypeSubStructName() string {
	return "RDS_O13_COMPONENT"
}

// RDS_O13 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDS_O13 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	NTE     []NTE            `hl7:"false,NTE"`
	PATIENT *RDS_O13_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RDS_O13_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RDS_O13) MessageTypeName() string {
	return "RDS_O13"
}

// RDS_O13_ENCODING represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDS_O13_ENCODING struct {
	RXE   *RXE  `hl7:"true,RXE"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RDS_O13_ENCODING) MessageTypeSubStructName() string {
	return "RDS_O13_ENCODING"
}

// RDS_O13_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDS_O13_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RDS_O13_OBSERVATION) MessageTypeSubStructName() string {
	return "RDS_O13_OBSERVATION"
}

// RDS_O13_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDS_O13_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *RDS_O13_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	ENCODING     *RDS_O13_ENCODING     `hl7:"false,ENCODING"`
	RXD          *RXD                  `hl7:"true,RXD"`
	RXR          []RXR                 `hl7:"true,RXR"`
	RXC          []RXC                 `hl7:"false,RXC"`
	OBSERVATION  []RDS_O13_OBSERVATION `hl7:"false,OBSERVATION"`
	FT1          []FT1                 `hl7:"false,FT1"`
	Other        []interface{}
}

func (s *RDS_O13_ORDER) MessageTypeSubStructName() string {
	return "RDS_O13_ORDER"
}

// RDS_O13_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDS_O13_ORDER_DETAIL struct {
	RXO                     *RXO                             `hl7:"true,RXO"`
	ORDER_DETAIL_SUPPLEMENT *RDS_O13_ORDER_DETAIL_SUPPLEMENT `hl7:"false,ORDER_DETAIL_SUPPLEMENT"`
	Other                   []interface{}
}

func (s *RDS_O13_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RDS_O13_ORDER_DETAIL"
}

// RDS_O13_ORDER_DETAIL_SUPPLEMENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDS_O13_ORDER_DETAIL_SUPPLEMENT struct {
	NTE       []NTE              `hl7:"true,NTE"`
	RXR       []RXR              `hl7:"true,RXR"`
	COMPONENT *RDS_O13_COMPONENT `hl7:"false,COMPONENT"`
	Other     []interface{}
}

func (s *RDS_O13_ORDER_DETAIL_SUPPLEMENT) MessageTypeSubStructName() string {
	return "RDS_O13_ORDER_DETAIL_SUPPLEMENT"
}

// RDS_O13_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDS_O13_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NTE           []NTE                  `hl7:"false,NTE"`
	AL1           []AL1                  `hl7:"false,AL1"`
	PATIENT_VISIT *RDS_O13_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	Other         []interface{}
}

func (s *RDS_O13_PATIENT) MessageTypeSubStructName() string {
	return "RDS_O13_PATIENT"
}

// RDS_O13_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDS_O13_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RDS_O13_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "RDS_O13_PATIENT_VISIT"
}

// RDY_K15 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RDY_K15 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MSA   *MSA  `hl7:"true,MSA"`
	ERR   *ERR  `hl7:"false,ERR"`
	QAK   *QAK  `hl7:"true,QAK"`
	QPD   *QPD  `hl7:"true,QPD"`
	DSP   []DSP `hl7:"false,DSP"`
	DSC   *DSC  `hl7:"false,DSC"`
	Other []interface{}
}

func (s *RDY_K15) MessageTypeName() string {
	return "RDY_K15"
}

// REF_I12_AUTCTD_SUPPGRP2 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type REF_I12_AUTCTD_SUPPGRP2 struct {
	AUT   *AUT `hl7:"true,AUT"`
	CTD   *CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *REF_I12_AUTCTD_SUPPGRP2) MessageTypeSubStructName() string {
	return "REF_I12_AUTCTD_SUPPGRP2"
}

// REF_I12_AUTHORIZATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type REF_I12_AUTHORIZATION struct {
	AUT   *AUT `hl7:"true,AUT"`
	CTD   *CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *REF_I12_AUTHORIZATION) MessageTypeSubStructName() string {
	return "REF_I12_AUTHORIZATION"
}

// REF_I12_AUTHORIZATION_CONTACT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type REF_I12_AUTHORIZATION_CONTACT struct {
	AUT   *AUT `hl7:"true,AUT"`
	CTD   *CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *REF_I12_AUTHORIZATION_CONTACT) MessageTypeSubStructName() string {
	return "REF_I12_AUTHORIZATION_CONTACT"
}

// REF_I12 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type REF_I12 struct {
	MSH                   *MSH                           `hl7:"true,MSH"`
	RF1                   *RF1                           `hl7:"false,RF1"`
	AUTHORIZATION_CONTACT *REF_I12_AUTHORIZATION_CONTACT `hl7:"false,AUTHORIZATION_CONTACT"`
	PROVIDER_CONTACT      []REF_I12_PROVIDER_CONTACT     `hl7:"true,PROVIDER_CONTACT"`
	PID                   *PID                           `hl7:"true,PID"`
	NK1                   []NK1                          `hl7:"false,NK1"`
	GT1                   []GT1                          `hl7:"false,GT1"`
	INSURANCE             []REF_I12_INSURANCE            `hl7:"false,INSURANCE"`
	ACC                   *ACC                           `hl7:"false,ACC"`
	DG1                   []DG1                          `hl7:"false,DG1"`
	DRG                   []DRG                          `hl7:"false,DRG"`
	AL1                   []AL1                          `hl7:"false,AL1"`
	PROCEDURE             []REF_I12_PROCEDURE            `hl7:"false,PROCEDURE"`
	OBSERVATION           []REF_I12_OBSERVATION          `hl7:"false,OBSERVATION"`
	PATIENT_VISIT         *REF_I12_PATIENT_VISIT         `hl7:"false,PATIENT_VISIT"`
	NTE                   []NTE                          `hl7:"false,NTE"`
	Other                 []interface{}
}

func (s *REF_I12) MessageTypeName() string {
	return "REF_I12"
}

// REF_I12_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type REF_I12_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *REF_I12_INSURANCE) MessageTypeSubStructName() string {
	return "REF_I12_INSURANCE"
}

// REF_I12_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type REF_I12_OBSERVATION struct {
	OBR           *OBR                    `hl7:"true,OBR"`
	NTE           []NTE                   `hl7:"false,NTE"`
	RESULTS_NOTES []REF_I12_RESULTS_NOTES `hl7:"false,RESULTS_NOTES"`
	Other         []interface{}
}

func (s *REF_I12_OBSERVATION) MessageTypeSubStructName() string {
	return "REF_I12_OBSERVATION"
}

// REF_I12_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type REF_I12_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *REF_I12_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "REF_I12_PATIENT_VISIT"
}

// REF_I12_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type REF_I12_PROCEDURE struct {
	PR1             *PR1                     `hl7:"true,PR1"`
	AUTCTD_SUPPGRP2 *REF_I12_AUTCTD_SUPPGRP2 `hl7:"false,AUTCTD_SUPPGRP2"`
	Other           []interface{}
}

func (s *REF_I12_PROCEDURE) MessageTypeSubStructName() string {
	return "REF_I12_PROCEDURE"
}

// REF_I12_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type REF_I12_PROVIDER struct {
	PRD   *PRD  `hl7:"true,PRD"`
	CTD   []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *REF_I12_PROVIDER) MessageTypeSubStructName() string {
	return "REF_I12_PROVIDER"
}

// REF_I12_PROVIDER_CONTACT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type REF_I12_PROVIDER_CONTACT struct {
	PRD   *PRD  `hl7:"true,PRD"`
	CTD   []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *REF_I12_PROVIDER_CONTACT) MessageTypeSubStructName() string {
	return "REF_I12_PROVIDER_CONTACT"
}

// REF_I12_RESULTS represents the corresponding HL7 message type.
// Definition from HL7 2.3
type REF_I12_RESULTS struct {
	OBR         *OBR                  `hl7:"true,OBR"`
	NTE         []NTE                 `hl7:"false,NTE"`
	OBSERVATION []REF_I12_OBSERVATION `hl7:"false,OBSERVATION"`
	Other       []interface{}
}

func (s *REF_I12_RESULTS) MessageTypeSubStructName() string {
	return "REF_I12_RESULTS"
}

// REF_I12_RESULTS_NOTES represents the corresponding HL7 message type.
// Definition from HL7 2.4
type REF_I12_RESULTS_NOTES struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *REF_I12_RESULTS_NOTES) MessageTypeSubStructName() string {
	return "REF_I12_RESULTS_NOTES"
}

// REF_I12_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type REF_I12_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *REF_I12_VISIT) MessageTypeSubStructName() string {
	return "REF_I12_VISIT"
}

// RER_RER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RER_RER struct {
	MSH        *MSH                 `hl7:"true,MSH"`
	MSA        *MSA                 `hl7:"true,MSA"`
	ERR        *ERR                 `hl7:"false,ERR"`
	DEFINITION []RER_RER_DEFINITION `hl7:"true,DEFINITION"`
	DSC        *DSC                 `hl7:"false,DSC"`
	Other      []interface{}
}

func (s *RER_RER) MessageTypeName() string {
	return "RER_RER"
}

// RER_RER_DEFINITION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RER_RER_DEFINITION struct {
	QRD     *QRD             `hl7:"true,QRD"`
	QRF     *QRF             `hl7:"false,QRF"`
	PATIENT *RER_RER_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RER_RER_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RER_RER_DEFINITION) MessageTypeSubStructName() string {
	return "RER_RER_DEFINITION"
}

// RER_RER_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RER_RER_ORDER struct {
	ORC   *ORC  `hl7:"true,ORC"`
	RXE   *RXE  `hl7:"true,RXE"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RER_RER_ORDER) MessageTypeSubStructName() string {
	return "RER_RER_ORDER"
}

// RER_RER_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RER_RER_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RER_RER_PATIENT) MessageTypeSubStructName() string {
	return "RER_RER_PATIENT"
}

// RGR_RGR represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RGR_RGR struct {
	MSH        *MSH                 `hl7:"true,MSH"`
	MSA        *MSA                 `hl7:"true,MSA"`
	ERR        *ERR                 `hl7:"false,ERR"`
	DEFINITION []RGR_RGR_DEFINITION `hl7:"true,DEFINITION"`
	DSC        *DSC                 `hl7:"false,DSC"`
	Other      []interface{}
}

func (s *RGR_RGR) MessageTypeName() string {
	return "RGR_RGR"
}

// RGR_RGR_DEFINITION represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RGR_RGR_DEFINITION struct {
	QRD     *QRD             `hl7:"true,QRD"`
	QRF     *QRF             `hl7:"false,QRF"`
	PATIENT *RGR_RGR_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RGR_RGR_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RGR_RGR_DEFINITION) MessageTypeSubStructName() string {
	return "RGR_RGR_DEFINITION"
}

// RGR_RGR_ENCODING represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RGR_RGR_ENCODING struct {
	RXE   *RXE  `hl7:"true,RXE"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RGR_RGR_ENCODING) MessageTypeSubStructName() string {
	return "RGR_RGR_ENCODING"
}

// RGR_RGR_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RGR_RGR_ORDER struct {
	ORC      *ORC              `hl7:"true,ORC"`
	ENCODING *RGR_RGR_ENCODING `hl7:"false,ENCODING"`
	RXG      []RXG             `hl7:"true,RXG"`
	RXR      []RXR             `hl7:"true,RXR"`
	RXC      []RXC             `hl7:"false,RXC"`
	Other    []interface{}
}

func (s *RGR_RGR_ORDER) MessageTypeSubStructName() string {
	return "RGR_RGR_ORDER"
}

// RGR_RGR_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RGR_RGR_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RGR_RGR_PATIENT) MessageTypeSubStructName() string {
	return "RGR_RGR_PATIENT"
}

// RGV_O01_COMPONENTS represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RGV_O01_COMPONENTS struct {
	RXC   []RXC `hl7:"true,RXC"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RGV_O01_COMPONENTS) MessageTypeSubStructName() string {
	return "RGV_O01_COMPONENTS"
}

// RGV_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RGV_O01 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	NTE     []NTE            `hl7:"false,NTE"`
	PATIENT *RGV_O01_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RGV_O01_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RGV_O01) MessageTypeName() string {
	return "RGV_O01"
}

// RGV_O01_ENCODING represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RGV_O01_ENCODING struct {
	RXE   *RXE  `hl7:"true,RXE"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RGV_O01_ENCODING) MessageTypeSubStructName() string {
	return "RGV_O01_ENCODING"
}

// RGV_O01_GIVE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RGV_O01_GIVE struct {
	RXG         *RXG                  `hl7:"true,RXG"`
	RXR         []RXR                 `hl7:"true,RXR"`
	RXC         []RXC                 `hl7:"false,RXC"`
	OBSERVATION []RGV_O01_OBSERVATION `hl7:"true,OBSERVATION"`
	Other       []interface{}
}

func (s *RGV_O01_GIVE) MessageTypeSubStructName() string {
	return "RGV_O01_GIVE"
}

// RGV_O01_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RGV_O01_OBSERVATION struct {
	OBX   *OBX  `hl7:"false,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RGV_O01_OBSERVATION) MessageTypeSubStructName() string {
	return "RGV_O01_OBSERVATION"
}

// RGV_O01_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RGV_O01_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *RGV_O01_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	ENCODING     *RGV_O01_ENCODING     `hl7:"false,ENCODING"`
	GIVE         []RGV_O01_GIVE        `hl7:"true,GIVE"`
	Other        []interface{}
}

func (s *RGV_O01_ORDER) MessageTypeSubStructName() string {
	return "RGV_O01_ORDER"
}

// RGV_O01_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RGV_O01_ORDER_DETAIL struct {
	RXO                     *RXO                             `hl7:"true,RXO"`
	ORDER_DETAIL_SUPPLEMENT *RGV_O01_ORDER_DETAIL_SUPPLEMENT `hl7:"false,ORDER_DETAIL_SUPPLEMENT"`
	Other                   []interface{}
}

func (s *RGV_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RGV_O01_ORDER_DETAIL"
}

// RGV_O01_ORDER_DETAIL_SUPPLEMENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RGV_O01_ORDER_DETAIL_SUPPLEMENT struct {
	NTE        []NTE               `hl7:"true,NTE"`
	RXR        []RXR               `hl7:"true,RXR"`
	COMPONENTS *RGV_O01_COMPONENTS `hl7:"false,COMPONENTS"`
	Other      []interface{}
}

func (s *RGV_O01_ORDER_DETAIL_SUPPLEMENT) MessageTypeSubStructName() string {
	return "RGV_O01_ORDER_DETAIL_SUPPLEMENT"
}

// RGV_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RGV_O01_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	NTE           []NTE                  `hl7:"false,NTE"`
	AL1           []AL1                  `hl7:"false,AL1"`
	PATIENT_VISIT *RGV_O01_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	Other         []interface{}
}

func (s *RGV_O01_PATIENT) MessageTypeSubStructName() string {
	return "RGV_O01_PATIENT"
}

// RGV_O01_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RGV_O01_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RGV_O01_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "RGV_O01_PATIENT_VISIT"
}

// RGV_O15_COMPONENTS represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RGV_O15_COMPONENTS struct {
	RXC   []RXC `hl7:"true,RXC"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RGV_O15_COMPONENTS) MessageTypeSubStructName() string {
	return "RGV_O15_COMPONENTS"
}

// RGV_O15 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RGV_O15 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	NTE     []NTE            `hl7:"false,NTE"`
	PATIENT *RGV_O15_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RGV_O15_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RGV_O15) MessageTypeName() string {
	return "RGV_O15"
}

// RGV_O15_ENCODING represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RGV_O15_ENCODING struct {
	RXE   *RXE  `hl7:"true,RXE"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RGV_O15_ENCODING) MessageTypeSubStructName() string {
	return "RGV_O15_ENCODING"
}

// RGV_O15_GIVE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RGV_O15_GIVE struct {
	RXG         *RXG                  `hl7:"true,RXG"`
	RXR         []RXR                 `hl7:"true,RXR"`
	RXC         []RXC                 `hl7:"false,RXC"`
	OBSERVATION []RGV_O15_OBSERVATION `hl7:"true,OBSERVATION"`
	Other       []interface{}
}

func (s *RGV_O15_GIVE) MessageTypeSubStructName() string {
	return "RGV_O15_GIVE"
}

// RGV_O15_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RGV_O15_OBSERVATION struct {
	OBX   *OBX  `hl7:"false,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RGV_O15_OBSERVATION) MessageTypeSubStructName() string {
	return "RGV_O15_OBSERVATION"
}

// RGV_O15_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RGV_O15_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *RGV_O15_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	ENCODING     *RGV_O15_ENCODING     `hl7:"false,ENCODING"`
	GIVE         []RGV_O15_GIVE        `hl7:"true,GIVE"`
	Other        []interface{}
}

func (s *RGV_O15_ORDER) MessageTypeSubStructName() string {
	return "RGV_O15_ORDER"
}

// RGV_O15_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RGV_O15_ORDER_DETAIL struct {
	RXO                     *RXO                             `hl7:"true,RXO"`
	ORDER_DETAIL_SUPPLEMENT *RGV_O15_ORDER_DETAIL_SUPPLEMENT `hl7:"false,ORDER_DETAIL_SUPPLEMENT"`
	Other                   []interface{}
}

func (s *RGV_O15_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RGV_O15_ORDER_DETAIL"
}

// RGV_O15_ORDER_DETAIL_SUPPLEMENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RGV_O15_ORDER_DETAIL_SUPPLEMENT struct {
	NTE        []NTE               `hl7:"true,NTE"`
	RXR        []RXR               `hl7:"true,RXR"`
	COMPONENTS *RGV_O15_COMPONENTS `hl7:"false,COMPONENTS"`
	Other      []interface{}
}

func (s *RGV_O15_ORDER_DETAIL_SUPPLEMENT) MessageTypeSubStructName() string {
	return "RGV_O15_ORDER_DETAIL_SUPPLEMENT"
}

// RGV_O15_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RGV_O15_PATIENT struct {
	PID           *PID                   `hl7:"true,PID"`
	NTE           []NTE                  `hl7:"false,NTE"`
	AL1           []AL1                  `hl7:"false,AL1"`
	PATIENT_VISIT *RGV_O15_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	Other         []interface{}
}

func (s *RGV_O15_PATIENT) MessageTypeSubStructName() string {
	return "RGV_O15_PATIENT"
}

// RGV_O15_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RGV_O15_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RGV_O15_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "RGV_O15_PATIENT_VISIT"
}

// ROR_ROR represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ROR_ROR struct {
	MSH        *MSH                 `hl7:"true,MSH"`
	MSA        *MSA                 `hl7:"true,MSA"`
	ERR        *ERR                 `hl7:"false,ERR"`
	DEFINITION []ROR_ROR_DEFINITION `hl7:"true,DEFINITION"`
	DSC        *DSC                 `hl7:"false,DSC"`
	Other      []interface{}
}

func (s *ROR_ROR) MessageTypeName() string {
	return "ROR_ROR"
}

// ROR_ROR_DEFINITION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ROR_ROR_DEFINITION struct {
	QRD     *QRD             `hl7:"true,QRD"`
	QRF     *QRF             `hl7:"false,QRF"`
	PATIENT *ROR_ROR_PATIENT `hl7:"false,PATIENT"`
	ORDER   []ROR_ROR_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *ROR_ROR_DEFINITION) MessageTypeSubStructName() string {
	return "ROR_ROR_DEFINITION"
}

// ROR_ROR_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ROR_ROR_ORDER struct {
	ORC   *ORC  `hl7:"true,ORC"`
	RXO   *RXO  `hl7:"true,RXO"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *ROR_ROR_ORDER) MessageTypeSubStructName() string {
	return "ROR_ROR_ORDER"
}

// ROR_ROR_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type ROR_ROR_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ROR_ROR_PATIENT) MessageTypeSubStructName() string {
	return "ROR_ROR_PATIENT"
}

// RPA_I08_AUTCTD_SUPPGRP2 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPA_I08_AUTCTD_SUPPGRP2 struct {
	AUT   *AUT `hl7:"true,AUT"`
	CTD   *CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RPA_I08_AUTCTD_SUPPGRP2) MessageTypeSubStructName() string {
	return "RPA_I08_AUTCTD_SUPPGRP2"
}

// RPA_I08_AUTHORIZATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPA_I08_AUTHORIZATION struct {
	AUT   *AUT `hl7:"true,AUT"`
	CTD   *CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RPA_I08_AUTHORIZATION) MessageTypeSubStructName() string {
	return "RPA_I08_AUTHORIZATION"
}

// RPA_I08 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPA_I08 struct {
	MSH           *MSH                   `hl7:"true,MSH"`
	MSA           *MSA                   `hl7:"true,MSA"`
	RF1           *RF1                   `hl7:"false,RF1"`
	AUTHORIZATION *RPA_I08_AUTHORIZATION `hl7:"false,AUTHORIZATION"`
	PROVIDER      []RPA_I08_PROVIDER     `hl7:"true,PROVIDER"`
	PID           *PID                   `hl7:"true,PID"`
	NK1           []NK1                  `hl7:"false,NK1"`
	GT1           []GT1                  `hl7:"false,GT1"`
	INSURANCE     []RPA_I08_INSURANCE    `hl7:"false,INSURANCE"`
	ACC           *ACC                   `hl7:"false,ACC"`
	DG1           []DG1                  `hl7:"false,DG1"`
	DRG           []DRG                  `hl7:"false,DRG"`
	AL1           []AL1                  `hl7:"false,AL1"`
	PROCEDURE     []RPA_I08_PROCEDURE    `hl7:"true,PROCEDURE"`
	OBSERVATION   []RPA_I08_OBSERVATION  `hl7:"false,OBSERVATION"`
	VISIT         *RPA_I08_VISIT         `hl7:"false,VISIT"`
	NTE           []NTE                  `hl7:"false,NTE"`
	Other         []interface{}
}

func (s *RPA_I08) MessageTypeName() string {
	return "RPA_I08"
}

// RPA_I08_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPA_I08_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *RPA_I08_INSURANCE) MessageTypeSubStructName() string {
	return "RPA_I08_INSURANCE"
}

// RPA_I08_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPA_I08_OBSERVATION struct {
	OBR     *OBR              `hl7:"true,OBR"`
	NTE     []NTE             `hl7:"false,NTE"`
	RESULTS []RPA_I08_RESULTS `hl7:"false,RESULTS"`
	Other   []interface{}
}

func (s *RPA_I08_OBSERVATION) MessageTypeSubStructName() string {
	return "RPA_I08_OBSERVATION"
}

// RPA_I08_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPA_I08_PROCEDURE struct {
	PR1             *PR1                     `hl7:"true,PR1"`
	AUTCTD_SUPPGRP2 *RPA_I08_AUTCTD_SUPPGRP2 `hl7:"false,AUTCTD_SUPPGRP2"`
	Other           []interface{}
}

func (s *RPA_I08_PROCEDURE) MessageTypeSubStructName() string {
	return "RPA_I08_PROCEDURE"
}

// RPA_I08_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPA_I08_PROVIDER struct {
	PRD   *PRD  `hl7:"true,PRD"`
	CTD   []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RPA_I08_PROVIDER) MessageTypeSubStructName() string {
	return "RPA_I08_PROVIDER"
}

// RPA_I08_RESULTS represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPA_I08_RESULTS struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RPA_I08_RESULTS) MessageTypeSubStructName() string {
	return "RPA_I08_RESULTS"
}

// RPA_I08_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPA_I08_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RPA_I08_VISIT) MessageTypeSubStructName() string {
	return "RPA_I08_VISIT"
}

// RPI_I01 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPI_I01 struct {
	MSH                 *MSH                         `hl7:"true,MSH"`
	MSA                 *MSA                         `hl7:"true,MSA"`
	PROVIDER            []RPI_I01_PROVIDER           `hl7:"true,PROVIDER"`
	PID                 *PID                         `hl7:"true,PID"`
	NK1                 []NK1                        `hl7:"false,NK1"`
	GUARANTOR_INSURANCE *RPI_I01_GUARANTOR_INSURANCE `hl7:"false,GUARANTOR_INSURANCE"`
	NTE                 []NTE                        `hl7:"false,NTE"`
	Other               []interface{}
}

func (s *RPI_I01) MessageTypeName() string {
	return "RPI_I01"
}

// RPI_I01_GUARANTOR_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPI_I01_GUARANTOR_INSURANCE struct {
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []RPI_I01_INSURANCE `hl7:"true,INSURANCE"`
	Other     []interface{}
}

func (s *RPI_I01_GUARANTOR_INSURANCE) MessageTypeSubStructName() string {
	return "RPI_I01_GUARANTOR_INSURANCE"
}

// RPI_I01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPI_I01_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *RPI_I01_INSURANCE) MessageTypeSubStructName() string {
	return "RPI_I01_INSURANCE"
}

// RPI_I01_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPI_I01_PROVIDER struct {
	PRD   *PRD  `hl7:"true,PRD"`
	CTD   []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RPI_I01_PROVIDER) MessageTypeSubStructName() string {
	return "RPI_I01_PROVIDER"
}

// RPI_I04 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPI_I04 struct {
	MSH                 *MSH                         `hl7:"true,MSH"`
	MSA                 *MSA                         `hl7:"true,MSA"`
	PROVIDER            []RPI_I04_PROVIDER           `hl7:"true,PROVIDER"`
	PID                 *PID                         `hl7:"true,PID"`
	NK1                 []NK1                        `hl7:"false,NK1"`
	GUARANTOR_INSURANCE *RPI_I04_GUARANTOR_INSURANCE `hl7:"false,GUARANTOR_INSURANCE"`
	NTE                 []NTE                        `hl7:"false,NTE"`
	Other               []interface{}
}

func (s *RPI_I04) MessageTypeName() string {
	return "RPI_I04"
}

// RPI_I04_GUARANTOR_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPI_I04_GUARANTOR_INSURANCE struct {
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []RPI_I04_INSURANCE `hl7:"true,INSURANCE"`
	Other     []interface{}
}

func (s *RPI_I04_GUARANTOR_INSURANCE) MessageTypeSubStructName() string {
	return "RPI_I04_GUARANTOR_INSURANCE"
}

// RPI_I04_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPI_I04_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *RPI_I04_INSURANCE) MessageTypeSubStructName() string {
	return "RPI_I04_INSURANCE"
}

// RPI_I04_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPI_I04_PROVIDER struct {
	PRD   *PRD  `hl7:"true,PRD"`
	CTD   []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RPI_I04_PROVIDER) MessageTypeSubStructName() string {
	return "RPI_I04_PROVIDER"
}

// RPL_I02 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPL_I02 struct {
	MSH      *MSH               `hl7:"true,MSH"`
	MSA      *MSA               `hl7:"true,MSA"`
	PROVIDER []RPL_I02_PROVIDER `hl7:"true,PROVIDER"`
	NTE      []NTE              `hl7:"false,NTE"`
	DSP      []DSP              `hl7:"false,DSP"`
	DSC      *DSC               `hl7:"false,DSC"`
	Other    []interface{}
}

func (s *RPL_I02) MessageTypeName() string {
	return "RPL_I02"
}

// RPL_I02_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPL_I02_PROVIDER struct {
	PRD   *PRD  `hl7:"true,PRD"`
	CTD   []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RPL_I02_PROVIDER) MessageTypeSubStructName() string {
	return "RPL_I02_PROVIDER"
}

// RPR_I03 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPR_I03 struct {
	MSH      *MSH               `hl7:"true,MSH"`
	MSA      *MSA               `hl7:"true,MSA"`
	PROVIDER []RPR_I03_PROVIDER `hl7:"true,PROVIDER"`
	PID      []PID              `hl7:"false,PID"`
	NTE      []NTE              `hl7:"false,NTE"`
	Other    []interface{}
}

func (s *RPR_I03) MessageTypeName() string {
	return "RPR_I03"
}

// RPR_I03_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RPR_I03_PROVIDER struct {
	PRD   *PRD  `hl7:"true,PRD"`
	CTD   []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RPR_I03_PROVIDER) MessageTypeSubStructName() string {
	return "RPR_I03_PROVIDER"
}

// RQA_I08_AUTCTD_SUPPGRP2 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQA_I08_AUTCTD_SUPPGRP2 struct {
	AUT   *AUT `hl7:"true,AUT"`
	CTD   *CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RQA_I08_AUTCTD_SUPPGRP2) MessageTypeSubStructName() string {
	return "RQA_I08_AUTCTD_SUPPGRP2"
}

// RQA_I08_AUTHORIZATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQA_I08_AUTHORIZATION struct {
	AUT   *AUT `hl7:"true,AUT"`
	CTD   *CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RQA_I08_AUTHORIZATION) MessageTypeSubStructName() string {
	return "RQA_I08_AUTHORIZATION"
}

// RQA_I08 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQA_I08 struct {
	MSH                 *MSH                         `hl7:"true,MSH"`
	RF1                 *RF1                         `hl7:"false,RF1"`
	AUTHORIZATION       *RQA_I08_AUTHORIZATION       `hl7:"false,AUTHORIZATION"`
	PROVIDER            []RQA_I08_PROVIDER           `hl7:"true,PROVIDER"`
	PID                 *PID                         `hl7:"true,PID"`
	NK1                 []NK1                        `hl7:"false,NK1"`
	GUARANTOR_INSURANCE *RQA_I08_GUARANTOR_INSURANCE `hl7:"false,GUARANTOR_INSURANCE"`
	ACC                 *ACC                         `hl7:"false,ACC"`
	DG1                 []DG1                        `hl7:"false,DG1"`
	DRG                 []DRG                        `hl7:"false,DRG"`
	AL1                 []AL1                        `hl7:"false,AL1"`
	PROCEDURE           []RQA_I08_PROCEDURE          `hl7:"false,PROCEDURE"`
	OBSERVATION         []RQA_I08_OBSERVATION        `hl7:"false,OBSERVATION"`
	VISIT               *RQA_I08_VISIT               `hl7:"false,VISIT"`
	NTE                 []NTE                        `hl7:"false,NTE"`
	Other               []interface{}
}

func (s *RQA_I08) MessageTypeName() string {
	return "RQA_I08"
}

// RQA_I08_GUARANTOR_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQA_I08_GUARANTOR_INSURANCE struct {
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []RQA_I08_INSURANCE `hl7:"true,INSURANCE"`
	Other     []interface{}
}

func (s *RQA_I08_GUARANTOR_INSURANCE) MessageTypeSubStructName() string {
	return "RQA_I08_GUARANTOR_INSURANCE"
}

// RQA_I08_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQA_I08_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *RQA_I08_INSURANCE) MessageTypeSubStructName() string {
	return "RQA_I08_INSURANCE"
}

// RQA_I08_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQA_I08_OBSERVATION struct {
	OBR     *OBR              `hl7:"true,OBR"`
	NTE     []NTE             `hl7:"false,NTE"`
	RESULTS []RQA_I08_RESULTS `hl7:"false,RESULTS"`
	Other   []interface{}
}

func (s *RQA_I08_OBSERVATION) MessageTypeSubStructName() string {
	return "RQA_I08_OBSERVATION"
}

// RQA_I08_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQA_I08_PROCEDURE struct {
	PR1             *PR1                     `hl7:"true,PR1"`
	AUTCTD_SUPPGRP2 *RQA_I08_AUTCTD_SUPPGRP2 `hl7:"false,AUTCTD_SUPPGRP2"`
	Other           []interface{}
}

func (s *RQA_I08_PROCEDURE) MessageTypeSubStructName() string {
	return "RQA_I08_PROCEDURE"
}

// RQA_I08_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQA_I08_PROVIDER struct {
	PRD   *PRD  `hl7:"true,PRD"`
	CTD   []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RQA_I08_PROVIDER) MessageTypeSubStructName() string {
	return "RQA_I08_PROVIDER"
}

// RQA_I08_RESULTS represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQA_I08_RESULTS struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RQA_I08_RESULTS) MessageTypeSubStructName() string {
	return "RQA_I08_RESULTS"
}

// RQA_I08_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQA_I08_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RQA_I08_VISIT) MessageTypeSubStructName() string {
	return "RQA_I08_VISIT"
}

// RQC_I05 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQC_I05 struct {
	MSH      *MSH               `hl7:"true,MSH"`
	QRD      *QRD               `hl7:"true,QRD"`
	QRF      *QRF               `hl7:"false,QRF"`
	PROVIDER []RQC_I05_PROVIDER `hl7:"true,PROVIDER"`
	PID      *PID               `hl7:"true,PID"`
	NK1      []NK1              `hl7:"false,NK1"`
	GT1      []GT1              `hl7:"false,GT1"`
	NTE      []NTE              `hl7:"false,NTE"`
	Other    []interface{}
}

func (s *RQC_I05) MessageTypeName() string {
	return "RQC_I05"
}

// RQC_I05_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQC_I05_PROVIDER struct {
	PRD   *PRD  `hl7:"true,PRD"`
	CTD   []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RQC_I05_PROVIDER) MessageTypeSubStructName() string {
	return "RQC_I05_PROVIDER"
}

// RQC_I06 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RQC_I06 struct {
	MSH      *MSH               `hl7:"true,MSH"`
	QRD      *QRD               `hl7:"true,QRD"`
	QRF      *QRF               `hl7:"false,QRF"`
	PROVIDER []RQC_I06_PROVIDER `hl7:"true,PROVIDER"`
	PID      *PID               `hl7:"true,PID"`
	NK1      []NK1              `hl7:"false,NK1"`
	GT1      *GT1               `hl7:"false,GT1"`
	NTE      []NTE              `hl7:"false,NTE"`
	Other    []interface{}
}

func (s *RQC_I06) MessageTypeName() string {
	return "RQC_I06"
}

// RQC_I06_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RQC_I06_PROVIDER struct {
	PRD   *PRD  `hl7:"true,PRD"`
	CTD   []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RQC_I06_PROVIDER) MessageTypeSubStructName() string {
	return "RQC_I06_PROVIDER"
}

// RQI_I01 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQI_I01 struct {
	MSH                 *MSH                         `hl7:"true,MSH"`
	PROVIDER            []RQI_I01_PROVIDER           `hl7:"true,PROVIDER"`
	PID                 *PID                         `hl7:"true,PID"`
	NK1                 []NK1                        `hl7:"false,NK1"`
	GUARANTOR_INSURANCE *RQI_I01_GUARANTOR_INSURANCE `hl7:"false,GUARANTOR_INSURANCE"`
	NTE                 []NTE                        `hl7:"false,NTE"`
	Other               []interface{}
}

func (s *RQI_I01) MessageTypeName() string {
	return "RQI_I01"
}

// RQI_I01_GUARANTOR_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQI_I01_GUARANTOR_INSURANCE struct {
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []RQI_I01_INSURANCE `hl7:"true,INSURANCE"`
	Other     []interface{}
}

func (s *RQI_I01_GUARANTOR_INSURANCE) MessageTypeSubStructName() string {
	return "RQI_I01_GUARANTOR_INSURANCE"
}

// RQI_I01_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQI_I01_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *RQI_I01_INSURANCE) MessageTypeSubStructName() string {
	return "RQI_I01_INSURANCE"
}

// RQI_I01_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQI_I01_PROVIDER struct {
	PRD   *PRD  `hl7:"true,PRD"`
	CTD   []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RQI_I01_PROVIDER) MessageTypeSubStructName() string {
	return "RQI_I01_PROVIDER"
}

// RQP_I04 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQP_I04 struct {
	MSH      *MSH               `hl7:"true,MSH"`
	PROVIDER []RQP_I04_PROVIDER `hl7:"true,PROVIDER"`
	PID      *PID               `hl7:"true,PID"`
	NK1      []NK1              `hl7:"false,NK1"`
	GT1      []GT1              `hl7:"false,GT1"`
	NTE      []NTE              `hl7:"false,NTE"`
	Other    []interface{}
}

func (s *RQP_I04) MessageTypeName() string {
	return "RQP_I04"
}

// RQP_I04_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQP_I04_PROVIDER struct {
	PRD   *PRD  `hl7:"true,PRD"`
	CTD   []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RQP_I04_PROVIDER) MessageTypeSubStructName() string {
	return "RQP_I04_PROVIDER"
}

// RQQ_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RQQ_Q01 struct {
	MSH   *MSH `hl7:"true,MSH"`
	ERQ   *ERQ `hl7:"true,ERQ"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *RQQ_Q01) MessageTypeName() string {
	return "RQQ_Q01"
}

// RQQ_Q09 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RQQ_Q09 struct {
	MSH   *MSH `hl7:"true,MSH"`
	ERQ   *ERQ `hl7:"true,ERQ"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *RQQ_Q09) MessageTypeName() string {
	return "RQQ_Q09"
}

// RRA_O02_ADMINISTRATION represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRA_O02_ADMINISTRATION struct {
	RXA   []RXA `hl7:"true,RXA"`
	RXR   *RXR  `hl7:"true,RXR"`
	Other []interface{}
}

func (s *RRA_O02_ADMINISTRATION) MessageTypeSubStructName() string {
	return "RRA_O02_ADMINISTRATION"
}

// RRA_O02 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRA_O02 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *RRA_O02_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *RRA_O02) MessageTypeName() string {
	return "RRA_O02"
}

// RRA_O02_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRA_O02_ORDER struct {
	ORC            *ORC                    `hl7:"true,ORC"`
	ADMINISTRATION *RRA_O02_ADMINISTRATION `hl7:"false,ADMINISTRATION"`
	Other          []interface{}
}

func (s *RRA_O02_ORDER) MessageTypeSubStructName() string {
	return "RRA_O02_ORDER"
}

// RRA_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRA_O02_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RRA_O02_PATIENT) MessageTypeSubStructName() string {
	return "RRA_O02_PATIENT"
}

// RRA_O02_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRA_O02_RESPONSE struct {
	PATIENT *RRA_O02_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RRA_O02_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RRA_O02_RESPONSE) MessageTypeSubStructName() string {
	return "RRA_O02_RESPONSE"
}

// RRA_O18_ADMINISTRATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRA_O18_ADMINISTRATION struct {
	RXA   []RXA `hl7:"true,RXA"`
	RXR   *RXR  `hl7:"true,RXR"`
	Other []interface{}
}

func (s *RRA_O18_ADMINISTRATION) MessageTypeSubStructName() string {
	return "RRA_O18_ADMINISTRATION"
}

// RRA_O18 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRA_O18 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *RRA_O18_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *RRA_O18) MessageTypeName() string {
	return "RRA_O18"
}

// RRA_O18_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRA_O18_ORDER struct {
	ORC            *ORC                    `hl7:"true,ORC"`
	ADMINISTRATION *RRA_O18_ADMINISTRATION `hl7:"false,ADMINISTRATION"`
	Other          []interface{}
}

func (s *RRA_O18_ORDER) MessageTypeSubStructName() string {
	return "RRA_O18_ORDER"
}

// RRA_O18_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRA_O18_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RRA_O18_PATIENT) MessageTypeSubStructName() string {
	return "RRA_O18_PATIENT"
}

// RRA_O18_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRA_O18_RESPONSE struct {
	PATIENT *RRA_O18_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RRA_O18_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RRA_O18_RESPONSE) MessageTypeSubStructName() string {
	return "RRA_O18_RESPONSE"
}

// RRD_O02 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRD_O02 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *RRD_O02_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *RRD_O02) MessageTypeName() string {
	return "RRD_O02"
}

// RRD_O02_DISPENSE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRD_O02_DISPENSE struct {
	RXD   *RXD  `hl7:"true,RXD"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RRD_O02_DISPENSE) MessageTypeSubStructName() string {
	return "RRD_O02_DISPENSE"
}

// RRD_O02_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRD_O02_ORDER struct {
	ORC      *ORC              `hl7:"true,ORC"`
	DISPENSE *RRD_O02_DISPENSE `hl7:"false,DISPENSE"`
	Other    []interface{}
}

func (s *RRD_O02_ORDER) MessageTypeSubStructName() string {
	return "RRD_O02_ORDER"
}

// RRD_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRD_O02_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RRD_O02_PATIENT) MessageTypeSubStructName() string {
	return "RRD_O02_PATIENT"
}

// RRD_O02_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRD_O02_RESPONSE struct {
	PATIENT *RRD_O02_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RRD_O02_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RRD_O02_RESPONSE) MessageTypeSubStructName() string {
	return "RRD_O02_RESPONSE"
}

// RRD_O14 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRD_O14 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *RRD_O14_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *RRD_O14) MessageTypeName() string {
	return "RRD_O14"
}

// RRD_O14_DISPENSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRD_O14_DISPENSE struct {
	RXD   *RXD  `hl7:"true,RXD"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RRD_O14_DISPENSE) MessageTypeSubStructName() string {
	return "RRD_O14_DISPENSE"
}

// RRD_O14_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRD_O14_ORDER struct {
	ORC      *ORC              `hl7:"true,ORC"`
	DISPENSE *RRD_O14_DISPENSE `hl7:"false,DISPENSE"`
	Other    []interface{}
}

func (s *RRD_O14_ORDER) MessageTypeSubStructName() string {
	return "RRD_O14_ORDER"
}

// RRD_O14_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRD_O14_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RRD_O14_PATIENT) MessageTypeSubStructName() string {
	return "RRD_O14_PATIENT"
}

// RRD_O14_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRD_O14_RESPONSE struct {
	PATIENT *RRD_O14_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RRD_O14_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RRD_O14_RESPONSE) MessageTypeSubStructName() string {
	return "RRD_O14_RESPONSE"
}

// RRE_O02 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRE_O02 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *RRE_O02_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *RRE_O02) MessageTypeName() string {
	return "RRE_O02"
}

// RRE_O02_ENCODING represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRE_O02_ENCODING struct {
	RXE   *RXE  `hl7:"true,RXE"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RRE_O02_ENCODING) MessageTypeSubStructName() string {
	return "RRE_O02_ENCODING"
}

// RRE_O02_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRE_O02_ORDER struct {
	ORC      *ORC              `hl7:"true,ORC"`
	ENCODING *RRE_O02_ENCODING `hl7:"false,ENCODING"`
	Other    []interface{}
}

func (s *RRE_O02_ORDER) MessageTypeSubStructName() string {
	return "RRE_O02_ORDER"
}

// RRE_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRE_O02_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RRE_O02_PATIENT) MessageTypeSubStructName() string {
	return "RRE_O02_PATIENT"
}

// RRE_O02_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRE_O02_RESPONSE struct {
	PATIENT *RRE_O02_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RRE_O02_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RRE_O02_RESPONSE) MessageTypeSubStructName() string {
	return "RRE_O02_RESPONSE"
}

// RRE_O12 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRE_O12 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *RRE_O12_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *RRE_O12) MessageTypeName() string {
	return "RRE_O12"
}

// RRE_O12_ENCODING represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRE_O12_ENCODING struct {
	RXE   *RXE  `hl7:"true,RXE"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RRE_O12_ENCODING) MessageTypeSubStructName() string {
	return "RRE_O12_ENCODING"
}

// RRE_O12_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRE_O12_ORDER struct {
	ORC      *ORC              `hl7:"true,ORC"`
	ENCODING *RRE_O12_ENCODING `hl7:"false,ENCODING"`
	Other    []interface{}
}

func (s *RRE_O12_ORDER) MessageTypeSubStructName() string {
	return "RRE_O12_ORDER"
}

// RRE_O12_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRE_O12_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RRE_O12_PATIENT) MessageTypeSubStructName() string {
	return "RRE_O12_PATIENT"
}

// RRE_O12_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRE_O12_RESPONSE struct {
	PATIENT *RRE_O12_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RRE_O12_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RRE_O12_RESPONSE) MessageTypeSubStructName() string {
	return "RRE_O12_RESPONSE"
}

// RRG_O02 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRG_O02 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *RRG_O02_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *RRG_O02) MessageTypeName() string {
	return "RRG_O02"
}

// RRG_O02_GIVE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRG_O02_GIVE struct {
	RXG   *RXG  `hl7:"true,RXG"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RRG_O02_GIVE) MessageTypeSubStructName() string {
	return "RRG_O02_GIVE"
}

// RRG_O02_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRG_O02_ORDER struct {
	ORC   *ORC          `hl7:"true,ORC"`
	GIVE  *RRG_O02_GIVE `hl7:"false,GIVE"`
	Other []interface{}
}

func (s *RRG_O02_ORDER) MessageTypeSubStructName() string {
	return "RRG_O02_ORDER"
}

// RRG_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRG_O02_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RRG_O02_PATIENT) MessageTypeSubStructName() string {
	return "RRG_O02_PATIENT"
}

// RRG_O02_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRG_O02_RESPONSE struct {
	PATIENT *RRG_O02_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RRG_O02_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RRG_O02_RESPONSE) MessageTypeSubStructName() string {
	return "RRG_O02_RESPONSE"
}

// RRG_O16 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRG_O16 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *RRG_O16_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *RRG_O16) MessageTypeName() string {
	return "RRG_O16"
}

// RRG_O16_GIVE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRG_O16_GIVE struct {
	RXG   *RXG  `hl7:"true,RXG"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RRG_O16_GIVE) MessageTypeSubStructName() string {
	return "RRG_O16_GIVE"
}

// RRG_O16_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRG_O16_ORDER struct {
	ORC   *ORC          `hl7:"true,ORC"`
	GIVE  *RRG_O16_GIVE `hl7:"false,GIVE"`
	Other []interface{}
}

func (s *RRG_O16_ORDER) MessageTypeSubStructName() string {
	return "RRG_O16_ORDER"
}

// RRG_O16_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRG_O16_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RRG_O16_PATIENT) MessageTypeSubStructName() string {
	return "RRG_O16_PATIENT"
}

// RRG_O16_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRG_O16_RESPONSE struct {
	PATIENT *RRG_O16_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RRG_O16_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RRG_O16_RESPONSE) MessageTypeSubStructName() string {
	return "RRG_O16_RESPONSE"
}

// RRI_I12_AUTCTD_SUPPGRP2 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRI_I12_AUTCTD_SUPPGRP2 struct {
	AUT   *AUT `hl7:"true,AUT"`
	CTD   *CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RRI_I12_AUTCTD_SUPPGRP2) MessageTypeSubStructName() string {
	return "RRI_I12_AUTCTD_SUPPGRP2"
}

// RRI_I12_AUTHORIZATION represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRI_I12_AUTHORIZATION struct {
	AUT   *AUT `hl7:"true,AUT"`
	CTD   *CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RRI_I12_AUTHORIZATION) MessageTypeSubStructName() string {
	return "RRI_I12_AUTHORIZATION"
}

// RRI_I12_AUTHORIZATION_CONTACT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRI_I12_AUTHORIZATION_CONTACT struct {
	AUT   *AUT `hl7:"true,AUT"`
	CTD   *CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RRI_I12_AUTHORIZATION_CONTACT) MessageTypeSubStructName() string {
	return "RRI_I12_AUTHORIZATION_CONTACT"
}

// RRI_I12 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRI_I12 struct {
	MSH                   *MSH                           `hl7:"true,MSH"`
	MSA                   *MSA                           `hl7:"false,MSA"`
	RF1                   *RF1                           `hl7:"false,RF1"`
	AUTHORIZATION_CONTACT *RRI_I12_AUTHORIZATION_CONTACT `hl7:"false,AUTHORIZATION_CONTACT"`
	PROVIDER_CONTACT      []RRI_I12_PROVIDER_CONTACT     `hl7:"true,PROVIDER_CONTACT"`
	PID                   *PID                           `hl7:"true,PID"`
	ACC                   *ACC                           `hl7:"false,ACC"`
	DG1                   []DG1                          `hl7:"false,DG1"`
	DRG                   []DRG                          `hl7:"false,DRG"`
	AL1                   []AL1                          `hl7:"false,AL1"`
	PROCEDURE             []RRI_I12_PROCEDURE            `hl7:"false,PROCEDURE"`
	OBSERVATION           []RRI_I12_OBSERVATION          `hl7:"false,OBSERVATION"`
	PATIENT_VISIT         *RRI_I12_PATIENT_VISIT         `hl7:"false,PATIENT_VISIT"`
	NTE                   []NTE                          `hl7:"false,NTE"`
	Other                 []interface{}
}

func (s *RRI_I12) MessageTypeName() string {
	return "RRI_I12"
}

// RRI_I12_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRI_I12_OBSERVATION struct {
	OBR           *OBR                    `hl7:"true,OBR"`
	NTE           []NTE                   `hl7:"false,NTE"`
	RESULTS_NOTES []RRI_I12_RESULTS_NOTES `hl7:"false,RESULTS_NOTES"`
	Other         []interface{}
}

func (s *RRI_I12_OBSERVATION) MessageTypeSubStructName() string {
	return "RRI_I12_OBSERVATION"
}

// RRI_I12_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRI_I12_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RRI_I12_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "RRI_I12_PATIENT_VISIT"
}

// RRI_I12_PROCEDURE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRI_I12_PROCEDURE struct {
	PR1             *PR1                     `hl7:"true,PR1"`
	AUTCTD_SUPPGRP2 *RRI_I12_AUTCTD_SUPPGRP2 `hl7:"false,AUTCTD_SUPPGRP2"`
	Other           []interface{}
}

func (s *RRI_I12_PROCEDURE) MessageTypeSubStructName() string {
	return "RRI_I12_PROCEDURE"
}

// RRI_I12_PROVIDER represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRI_I12_PROVIDER struct {
	PRD   *PRD  `hl7:"true,PRD"`
	CTD   []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RRI_I12_PROVIDER) MessageTypeSubStructName() string {
	return "RRI_I12_PROVIDER"
}

// RRI_I12_PROVIDER_CONTACT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRI_I12_PROVIDER_CONTACT struct {
	PRD   *PRD  `hl7:"true,PRD"`
	CTD   []CTD `hl7:"false,CTD"`
	Other []interface{}
}

func (s *RRI_I12_PROVIDER_CONTACT) MessageTypeSubStructName() string {
	return "RRI_I12_PROVIDER_CONTACT"
}

// RRI_I12_RESULTS represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRI_I12_RESULTS struct {
	OBR         *OBR                  `hl7:"true,OBR"`
	NTE         []NTE                 `hl7:"false,NTE"`
	OBSERVATION []RRI_I12_OBSERVATION `hl7:"false,OBSERVATION"`
	Other       []interface{}
}

func (s *RRI_I12_RESULTS) MessageTypeSubStructName() string {
	return "RRI_I12_RESULTS"
}

// RRI_I12_RESULTS_NOTES represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RRI_I12_RESULTS_NOTES struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RRI_I12_RESULTS_NOTES) MessageTypeSubStructName() string {
	return "RRI_I12_RESULTS_NOTES"
}

// RRI_I12_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.3
type RRI_I12_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RRI_I12_VISIT) MessageTypeSubStructName() string {
	return "RRI_I12_VISIT"
}

// RRO_O02 represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRO_O02 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	NTE      []NTE             `hl7:"false,NTE"`
	RESPONSE *RRO_O02_RESPONSE `hl7:"false,RESPONSE"`
	Other    []interface{}
}

func (s *RRO_O02) MessageTypeName() string {
	return "RRO_O02"
}

// RRO_O02_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRO_O02_ORDER struct {
	ORC          *ORC                  `hl7:"true,ORC"`
	ORDER_DETAIL *RRO_O02_ORDER_DETAIL `hl7:"false,ORDER_DETAIL"`
	Other        []interface{}
}

func (s *RRO_O02_ORDER) MessageTypeSubStructName() string {
	return "RRO_O02_ORDER"
}

// RRO_O02_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRO_O02_ORDER_DETAIL struct {
	RXO   *RXO  `hl7:"true,RXO"`
	NTE1  []NTE `hl7:"false,NTE1"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	NTE2  []NTE `hl7:"false,NTE2"`
	Other []interface{}
}

func (s *RRO_O02_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RRO_O02_ORDER_DETAIL"
}

// RRO_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRO_O02_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RRO_O02_PATIENT) MessageTypeSubStructName() string {
	return "RRO_O02_PATIENT"
}

// RRO_O02_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.3.1
type RRO_O02_RESPONSE struct {
	PATIENT *RRO_O02_PATIENT `hl7:"false,PATIENT"`
	ORDER   []RRO_O02_ORDER  `hl7:"true,ORDER"`
	Other   []interface{}
}

func (s *RRO_O02_RESPONSE) MessageTypeSubStructName() string {
	return "RRO_O02_RESPONSE"
}

// RSP_K11 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_K11 struct {
	MSH *MSH `hl7:"true,MSH"`
	MSA *MSA `hl7:"true,MSA"`
	ERR *ERR `hl7:"false,ERR"`
	QAK *QAK `hl7:"true,QAK"`
	QPD *QPD `hl7:"true,QPD"`
	// Missing: anyZSegment
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *RSP_K11) MessageTypeName() string {
	return "RSP_K11"
}

// RSP_K13 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_K13 struct {
	MSH            *MSH                    `hl7:"true,MSH"`
	MSA            *MSA                    `hl7:"true,MSA"`
	ERR            *ERR                    `hl7:"false,ERR"`
	QAK            *QAK                    `hl7:"true,QAK"`
	QPD            *QPD                    `hl7:"true,QPD"`
	ROW_DEFINITION *RSP_K13_ROW_DEFINITION `hl7:"false,ROW_DEFINITION"`
	DSC            *DSC                    `hl7:"false,DSC"`
	Other          []interface{}
}

func (s *RSP_K13) MessageTypeName() string {
	return "RSP_K13"
}

// RSP_K13_ROW_DEFINITION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_K13_ROW_DEFINITION struct {
	RDF   *RDF  `hl7:"true,RDF"`
	RDT   []RDT `hl7:"false,RDT"`
	Other []interface{}
}

func (s *RSP_K13_ROW_DEFINITION) MessageTypeSubStructName() string {
	return "RSP_K13_ROW_DEFINITION"
}

// RSP_K15 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_K15 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MSA   *MSA  `hl7:"true,MSA"`
	ERR   *ERR  `hl7:"false,ERR"`
	QAK   *QAK  `hl7:"true,QAK"`
	QPD   *QPD  `hl7:"true,QPD"`
	DSP   []DSP `hl7:"false,DSP"`
	DSC   *DSC  `hl7:"false,DSC"`
	Other []interface{}
}

func (s *RSP_K15) MessageTypeName() string {
	return "RSP_K15"
}

// RSP_K21 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_K21 struct {
	MSH            *MSH                    `hl7:"true,MSH"`
	MSA            *MSA                    `hl7:"true,MSA"`
	ERR            *ERR                    `hl7:"false,ERR"`
	QAK            *QAK                    `hl7:"true,QAK"`
	QPD            *QPD                    `hl7:"true,QPD"`
	QUERY_RESPONSE *RSP_K21_QUERY_RESPONSE `hl7:"false,QUERY_RESPONSE"`
	DSC            *DSC                    `hl7:"false,DSC"`
	Other          []interface{}
}

func (s *RSP_K21) MessageTypeName() string {
	return "RSP_K21"
}

// RSP_K21_QUERY_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_K21_QUERY_RESPONSE struct {
	PID   *PID `hl7:"true,PID"`
	PD1   *PD1 `hl7:"false,PD1"`
	Other []interface{}
}

func (s *RSP_K21_QUERY_RESPONSE) MessageTypeSubStructName() string {
	return "RSP_K21_QUERY_RESPONSE"
}

// RSP_K22 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_K22 struct {
	MSH            *MSH                     `hl7:"true,MSH"`
	MSA            *MSA                     `hl7:"true,MSA"`
	ERR            *ERR                     `hl7:"false,ERR"`
	QAK            *QAK                     `hl7:"true,QAK"`
	QPD            *QPD                     `hl7:"true,QPD"`
	QUERY_RESPONSE []RSP_K22_QUERY_RESPONSE `hl7:"false,QUERY_RESPONSE"`
	DSC            *DSC                     `hl7:"false,DSC"`
	Other          []interface{}
}

func (s *RSP_K22) MessageTypeName() string {
	return "RSP_K22"
}

// RSP_K22_QUERY_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_K22_QUERY_RESPONSE struct {
	PID   *PID `hl7:"true,PID"`
	PD1   *PD1 `hl7:"false,PD1"`
	QRI   *QRI `hl7:"false,QRI"`
	Other []interface{}
}

func (s *RSP_K22_QUERY_RESPONSE) MessageTypeSubStructName() string {
	return "RSP_K22_QUERY_RESPONSE"
}

// RSP_K23 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_K23 struct {
	MSH   *MSH `hl7:"true,MSH"`
	MSA   *MSA `hl7:"true,MSA"`
	ERR   *ERR `hl7:"false,ERR"`
	QAK   *QAK `hl7:"true,QAK"`
	QPD   *QPD `hl7:"true,QPD"`
	PID   *PID `hl7:"false,PID"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *RSP_K23) MessageTypeName() string {
	return "RSP_K23"
}

// RSP_K24 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_K24 struct {
	MSH   *MSH `hl7:"true,MSH"`
	MSA   *MSA `hl7:"true,MSA"`
	ERR   *ERR `hl7:"false,ERR"`
	QAK   *QAK `hl7:"true,QAK"`
	QPD   *QPD `hl7:"true,QPD"`
	PID   *PID `hl7:"false,PID"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *RSP_K24) MessageTypeName() string {
	return "RSP_K24"
}

// RSP_K25 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_K25 struct {
	MSH   *MSH            `hl7:"true,MSH"`
	MSA   *MSA            `hl7:"true,MSA"`
	ERR   *ERR            `hl7:"false,ERR"`
	QAK   *QAK            `hl7:"true,QAK"`
	QPD   *QPD            `hl7:"true,QPD"`
	RCP   *RCP            `hl7:"true,RCP"`
	STAFF []RSP_K25_STAFF `hl7:"true,STAFF"`
	DSC   *DSC            `hl7:"false,DSC"`
	Other []interface{}
}

func (s *RSP_K25) MessageTypeName() string {
	return "RSP_K25"
}

// RSP_K25_STAFF represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_K25_STAFF struct {
	STF   *STF  `hl7:"true,STF"`
	PRA   *PRA  `hl7:"false,PRA"`
	ORG   []ORG `hl7:"false,ORG"`
	AFF   []AFF `hl7:"false,AFF"`
	LAN   []LAN `hl7:"false,LAN"`
	EDU   []EDU `hl7:"false,EDU"`
	Other []interface{}
}

func (s *RSP_K25_STAFF) MessageTypeSubStructName() string {
	return "RSP_K25_STAFF"
}

// RSP_Z82_COMMON_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z82_COMMON_ORDER struct {
	ORC           *ORC                   `hl7:"true,ORC"`
	ORDER_DETAIL  *RSP_Z82_ORDER_DETAIL  `hl7:"false,ORDER_DETAIL"`
	ENCODED_ORDER *RSP_Z82_ENCODED_ORDER `hl7:"false,ENCODED_ORDER"`
	RXD           *RXD                   `hl7:"true,RXD"`
	RXR           []RXR                  `hl7:"true,RXR"`
	RXC           []RXC                  `hl7:"false,RXC"`
	OBSERVATION   []RSP_Z82_OBSERVATION  `hl7:"true,OBSERVATION"`
	Other         []interface{}
}

func (s *RSP_Z82_COMMON_ORDER) MessageTypeSubStructName() string {
	return "RSP_Z82_COMMON_ORDER"
}

// RSP_Z82 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z82 struct {
	MSH            *MSH                     `hl7:"true,MSH"`
	MSA            *MSA                     `hl7:"true,MSA"`
	ERR            *ERR                     `hl7:"false,ERR"`
	QAK            *QAK                     `hl7:"true,QAK"`
	QPD            *QPD                     `hl7:"true,QPD"`
	RCP            *RCP                     `hl7:"true,RCP"`
	QUERY_RESPONSE []RSP_Z82_QUERY_RESPONSE `hl7:"true,QUERY_RESPONSE"`
	DSC            *DSC                     `hl7:"false,DSC"`
	Other          []interface{}
}

func (s *RSP_Z82) MessageTypeName() string {
	return "RSP_Z82"
}

// RSP_Z82_ENCODED_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z82_ENCODED_ORDER struct {
	RXE   *RXE  `hl7:"true,RXE"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RSP_Z82_ENCODED_ORDER) MessageTypeSubStructName() string {
	return "RSP_Z82_ENCODED_ORDER"
}

// RSP_Z82_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z82_OBSERVATION struct {
	OBX   *OBX  `hl7:"false,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RSP_Z82_OBSERVATION) MessageTypeSubStructName() string {
	return "RSP_Z82_OBSERVATION"
}

// RSP_Z82_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z82_ORDER_DETAIL struct {
	RXO       *RXO               `hl7:"true,RXO"`
	NTE       []NTE              `hl7:"false,NTE"`
	RXR       []RXR              `hl7:"true,RXR"`
	TREATMENT *RSP_Z82_TREATMENT `hl7:"false,TREATMENT"`
	Other     []interface{}
}

func (s *RSP_Z82_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RSP_Z82_ORDER_DETAIL"
}

// RSP_Z82_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z82_PATIENT struct {
	PID          *PID                   `hl7:"true,PID"`
	PD1          *PD1                   `hl7:"false,PD1"`
	NTE          []NTE                  `hl7:"false,NTE"`
	VISIT        *RSP_Z82_VISIT         `hl7:"false,VISIT"`
	COMMON_ORDER []RSP_Z82_COMMON_ORDER `hl7:"true,COMMON_ORDER"`
	Other        []interface{}
}

func (s *RSP_Z82_PATIENT) MessageTypeSubStructName() string {
	return "RSP_Z82_PATIENT"
}

// RSP_Z82_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z82_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RSP_Z82_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "RSP_Z82_PATIENT_VISIT"
}

// RSP_Z82_QUERY_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z82_QUERY_RESPONSE struct {
	PATIENT *RSP_Z82_PATIENT `hl7:"false,PATIENT"`
	Other   []interface{}
}

func (s *RSP_Z82_QUERY_RESPONSE) MessageTypeSubStructName() string {
	return "RSP_Z82_QUERY_RESPONSE"
}

// RSP_Z82_TREATMENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z82_TREATMENT struct {
	RXC   []RXC `hl7:"true,RXC"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RSP_Z82_TREATMENT) MessageTypeSubStructName() string {
	return "RSP_Z82_TREATMENT"
}

// RSP_Z82_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z82_VISIT struct {
	AL1           []AL1                  `hl7:"true,AL1"`
	PATIENT_VISIT *RSP_Z82_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	Other         []interface{}
}

func (s *RSP_Z82_VISIT) MessageTypeSubStructName() string {
	return "RSP_Z82_VISIT"
}

// RSP_Z86_ADMINISTRATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z86_ADMINISTRATION struct {
	RXA   *RXA  `hl7:"true,RXA"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RSP_Z86_ADMINISTRATION) MessageTypeSubStructName() string {
	return "RSP_Z86_ADMINISTRATION"
}

// RSP_Z86_COMMON_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z86_COMMON_ORDER struct {
	ORC            *ORC                    `hl7:"true,ORC"`
	ORDER_DETAIL   *RSP_Z86_ORDER_DETAIL   `hl7:"false,ORDER_DETAIL"`
	ENCODED_ORDER  *RSP_Z86_ENCODED_ORDER  `hl7:"false,ENCODED_ORDER"`
	DISPENSE       *RSP_Z86_DISPENSE       `hl7:"false,DISPENSE"`
	GIVE           *RSP_Z86_GIVE           `hl7:"false,GIVE"`
	ADMINISTRATION *RSP_Z86_ADMINISTRATION `hl7:"false,ADMINISTRATION"`
	OBSERVATION    []RSP_Z86_OBSERVATION   `hl7:"true,OBSERVATION"`
	Other          []interface{}
}

func (s *RSP_Z86_COMMON_ORDER) MessageTypeSubStructName() string {
	return "RSP_Z86_COMMON_ORDER"
}

// RSP_Z86 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z86 struct {
	MSH            *MSH                     `hl7:"true,MSH"`
	MSA            *MSA                     `hl7:"true,MSA"`
	ERR            *ERR                     `hl7:"false,ERR"`
	QAK            *QAK                     `hl7:"true,QAK"`
	QPD            *QPD                     `hl7:"true,QPD"`
	QUERY_RESPONSE []RSP_Z86_QUERY_RESPONSE `hl7:"true,QUERY_RESPONSE"`
	DSC            *DSC                     `hl7:"false,DSC"`
	Other          []interface{}
}

func (s *RSP_Z86) MessageTypeName() string {
	return "RSP_Z86"
}

// RSP_Z86_DISPENSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z86_DISPENSE struct {
	RXD   *RXD  `hl7:"true,RXD"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RSP_Z86_DISPENSE) MessageTypeSubStructName() string {
	return "RSP_Z86_DISPENSE"
}

// RSP_Z86_ENCODED_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z86_ENCODED_ORDER struct {
	RXE   *RXE  `hl7:"true,RXE"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RSP_Z86_ENCODED_ORDER) MessageTypeSubStructName() string {
	return "RSP_Z86_ENCODED_ORDER"
}

// RSP_Z86_GIVE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z86_GIVE struct {
	RXG   *RXG  `hl7:"true,RXG"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RSP_Z86_GIVE) MessageTypeSubStructName() string {
	return "RSP_Z86_GIVE"
}

// RSP_Z86_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z86_OBSERVATION struct {
	OBX   *OBX  `hl7:"false,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RSP_Z86_OBSERVATION) MessageTypeSubStructName() string {
	return "RSP_Z86_OBSERVATION"
}

// RSP_Z86_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z86_ORDER_DETAIL struct {
	RXO   *RXO  `hl7:"true,RXO"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RSP_Z86_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RSP_Z86_ORDER_DETAIL"
}

// RSP_Z86_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z86_PATIENT struct {
	PID          *PID                   `hl7:"true,PID"`
	PD1          *PD1                   `hl7:"false,PD1"`
	NTE          []NTE                  `hl7:"false,NTE"`
	AL1          []AL1                  `hl7:"false,AL1"`
	COMMON_ORDER []RSP_Z86_COMMON_ORDER `hl7:"true,COMMON_ORDER"`
	Other        []interface{}
}

func (s *RSP_Z86_PATIENT) MessageTypeSubStructName() string {
	return "RSP_Z86_PATIENT"
}

// RSP_Z86_QUERY_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z86_QUERY_RESPONSE struct {
	PATIENT *RSP_Z86_PATIENT `hl7:"false,PATIENT"`
	Other   []interface{}
}

func (s *RSP_Z86_QUERY_RESPONSE) MessageTypeSubStructName() string {
	return "RSP_Z86_QUERY_RESPONSE"
}

// RSP_Z88_ALLERGY represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z88_ALLERGY struct {
	AL1   []AL1          `hl7:"true,AL1"`
	VISIT *RSP_Z88_VISIT `hl7:"false,VISIT"`
	Other []interface{}
}

func (s *RSP_Z88_ALLERGY) MessageTypeSubStructName() string {
	return "RSP_Z88_ALLERGY"
}

// RSP_Z88_COMMON_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z88_COMMON_ORDER struct {
	ORC           *ORC                   `hl7:"true,ORC"`
	ORDER_DETAIL  *RSP_Z88_ORDER_DETAIL  `hl7:"false,ORDER_DETAIL"`
	ORDER_ENCODED *RSP_Z88_ORDER_ENCODED `hl7:"false,ORDER_ENCODED"`
	RXD           *RXD                   `hl7:"true,RXD"`
	RXR           []RXR                  `hl7:"true,RXR"`
	RXC           []RXC                  `hl7:"false,RXC"`
	OBSERVATION   []RSP_Z88_OBSERVATION  `hl7:"true,OBSERVATION"`
	Other         []interface{}
}

func (s *RSP_Z88_COMMON_ORDER) MessageTypeSubStructName() string {
	return "RSP_Z88_COMMON_ORDER"
}

// RSP_Z88_COMPONENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z88_COMPONENT struct {
	RXC   []RXC `hl7:"true,RXC"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RSP_Z88_COMPONENT) MessageTypeSubStructName() string {
	return "RSP_Z88_COMPONENT"
}

// RSP_Z88 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z88 struct {
	MSH            *MSH                     `hl7:"true,MSH"`
	MSA            *MSA                     `hl7:"true,MSA"`
	ERR            *ERR                     `hl7:"false,ERR"`
	QAK            *QAK                     `hl7:"true,QAK"`
	QPD            *QPD                     `hl7:"true,QPD"`
	RCP            *RCP                     `hl7:"true,RCP"`
	QUERY_RESPONSE []RSP_Z88_QUERY_RESPONSE `hl7:"true,QUERY_RESPONSE"`
	DSC            *DSC                     `hl7:"true,DSC"`
	Other          []interface{}
}

func (s *RSP_Z88) MessageTypeName() string {
	return "RSP_Z88"
}

// RSP_Z88_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z88_OBSERVATION struct {
	OBX   *OBX  `hl7:"false,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RSP_Z88_OBSERVATION) MessageTypeSubStructName() string {
	return "RSP_Z88_OBSERVATION"
}

// RSP_Z88_ORDER_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z88_ORDER_DETAIL struct {
	RXO       *RXO               `hl7:"true,RXO"`
	NTE       []NTE              `hl7:"false,NTE"`
	RXR       []RXR              `hl7:"true,RXR"`
	COMPONENT *RSP_Z88_COMPONENT `hl7:"false,COMPONENT"`
	Other     []interface{}
}

func (s *RSP_Z88_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "RSP_Z88_ORDER_DETAIL"
}

// RSP_Z88_ORDER_ENCODED represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z88_ORDER_ENCODED struct {
	RXE   *RXE  `hl7:"true,RXE"`
	RXR   []RXR `hl7:"true,RXR"`
	RXC   []RXC `hl7:"false,RXC"`
	Other []interface{}
}

func (s *RSP_Z88_ORDER_ENCODED) MessageTypeSubStructName() string {
	return "RSP_Z88_ORDER_ENCODED"
}

// RSP_Z88_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z88_PATIENT struct {
	PID          *PID                   `hl7:"true,PID"`
	PD1          *PD1                   `hl7:"false,PD1"`
	NTE          []NTE                  `hl7:"false,NTE"`
	ALLERGY      *RSP_Z88_ALLERGY       `hl7:"false,ALLERGY"`
	COMMON_ORDER []RSP_Z88_COMMON_ORDER `hl7:"true,COMMON_ORDER"`
	Other        []interface{}
}

func (s *RSP_Z88_PATIENT) MessageTypeSubStructName() string {
	return "RSP_Z88_PATIENT"
}

// RSP_Z88_QUERY_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z88_QUERY_RESPONSE struct {
	PATIENT *RSP_Z88_PATIENT `hl7:"false,PATIENT"`
	Other   []interface{}
}

func (s *RSP_Z88_QUERY_RESPONSE) MessageTypeSubStructName() string {
	return "RSP_Z88_QUERY_RESPONSE"
}

// RSP_Z88_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z88_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RSP_Z88_VISIT) MessageTypeSubStructName() string {
	return "RSP_Z88_VISIT"
}

// RSP_Z90_COMMON_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z90_COMMON_ORDER struct {
	ORC         *ORC                  `hl7:"true,ORC"`
	OBR         *OBR                  `hl7:"true,OBR"`
	NTE         []NTE                 `hl7:"false,NTE"`
	CTD         *CTD                  `hl7:"false,CTD"`
	OBSERVATION []RSP_Z90_OBSERVATION `hl7:"true,OBSERVATION"`
	Other       []interface{}
}

func (s *RSP_Z90_COMMON_ORDER) MessageTypeSubStructName() string {
	return "RSP_Z90_COMMON_ORDER"
}

// RSP_Z90 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z90 struct {
	MSH            *MSH                     `hl7:"true,MSH"`
	MSA            *MSA                     `hl7:"true,MSA"`
	ERR            *ERR                     `hl7:"false,ERR"`
	QAK            *QAK                     `hl7:"true,QAK"`
	QPD            *QPD                     `hl7:"true,QPD"`
	RCP            *RCP                     `hl7:"true,RCP"`
	QUERY_RESPONSE []RSP_Z90_QUERY_RESPONSE `hl7:"true,QUERY_RESPONSE"`
	DSC            *DSC                     `hl7:"true,DSC"`
	Other          []interface{}
}

func (s *RSP_Z90) MessageTypeName() string {
	return "RSP_Z90"
}

// RSP_Z90_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z90_OBSERVATION struct {
	OBX   *OBX  `hl7:"false,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *RSP_Z90_OBSERVATION) MessageTypeSubStructName() string {
	return "RSP_Z90_OBSERVATION"
}

// RSP_Z90_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z90_PATIENT struct {
	PID   *PID           `hl7:"true,PID"`
	PD1   *PD1           `hl7:"false,PD1"`
	NK1   []NK1          `hl7:"false,NK1"`
	NTE   []NTE          `hl7:"false,NTE"`
	VISIT *RSP_Z90_VISIT `hl7:"false,VISIT"`
	Other []interface{}
}

func (s *RSP_Z90_PATIENT) MessageTypeSubStructName() string {
	return "RSP_Z90_PATIENT"
}

// RSP_Z90_QUERY_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z90_QUERY_RESPONSE struct {
	PATIENT      *RSP_Z90_PATIENT       `hl7:"false,PATIENT"`
	COMMON_ORDER []RSP_Z90_COMMON_ORDER `hl7:"true,COMMON_ORDER"`
	Other        []interface{}
}

func (s *RSP_Z90_QUERY_RESPONSE) MessageTypeSubStructName() string {
	return "RSP_Z90_QUERY_RESPONSE"
}

// RSP_Z90_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RSP_Z90_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *RSP_Z90_VISIT) MessageTypeSubStructName() string {
	return "RSP_Z90_VISIT"
}

// RTB_K13 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RTB_K13 struct {
	MSH            *MSH                    `hl7:"true,MSH"`
	MSA            *MSA                    `hl7:"true,MSA"`
	ERR            *ERR                    `hl7:"false,ERR"`
	QAK            *QAK                    `hl7:"true,QAK"`
	QPD            *QPD                    `hl7:"true,QPD"`
	ROW_DEFINITION *RTB_K13_ROW_DEFINITION `hl7:"false,ROW_DEFINITION"`
	DSC            *DSC                    `hl7:"false,DSC"`
	Other          []interface{}
}

func (s *RTB_K13) MessageTypeName() string {
	return "RTB_K13"
}

// RTB_K13_ROW_DEFINITION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RTB_K13_ROW_DEFINITION struct {
	RDF   *RDF  `hl7:"true,RDF"`
	RDT   []RDT `hl7:"false,RDT"`
	Other []interface{}
}

func (s *RTB_K13_ROW_DEFINITION) MessageTypeSubStructName() string {
	return "RTB_K13_ROW_DEFINITION"
}

// RTB_Knn represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RTB_Knn struct {
	MSH   *MSH `hl7:"true,MSH"`
	MSA   *MSA `hl7:"true,MSA"`
	ERR   *ERR `hl7:"false,ERR"`
	QAK   *QAK `hl7:"true,QAK"`
	QPD   *QPD `hl7:"true,QPD"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *RTB_Knn) MessageTypeName() string {
	return "RTB_Knn"
}

// RTB_Q13 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RTB_Q13 struct {
	MSH            *MSH                    `hl7:"true,MSH"`
	MSA            *MSA                    `hl7:"true,MSA"`
	ERR            *ERR                    `hl7:"false,ERR"`
	QAK            *QAK                    `hl7:"true,QAK"`
	QPD            *QPD                    `hl7:"true,QPD"`
	ROW_DEFINITION *RTB_Q13_ROW_DEFINITION `hl7:"false,ROW_DEFINITION"`
	DSC            *DSC                    `hl7:"false,DSC"`
	Other          []interface{}
}

func (s *RTB_Q13) MessageTypeName() string {
	return "RTB_Q13"
}

// RTB_Q13_ROW_DEFINITION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RTB_Q13_ROW_DEFINITION struct {
	RDF   *RDF  `hl7:"true,RDF"`
	RDT   []RDT `hl7:"false,RDT"`
	Other []interface{}
}

func (s *RTB_Q13_ROW_DEFINITION) MessageTypeSubStructName() string {
	return "RTB_Q13_ROW_DEFINITION"
}

// RTB_Z74 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RTB_Z74 struct {
	MSH            *MSH                    `hl7:"true,MSH"`
	MSA            *MSA                    `hl7:"true,MSA"`
	ERR            *ERR                    `hl7:"false,ERR"`
	QAK            *QAK                    `hl7:"true,QAK"`
	QPD            *QPD                    `hl7:"true,QPD"`
	ROW_DEFINITION *RTB_Z74_ROW_DEFINITION `hl7:"false,ROW_DEFINITION"`
	DSC            *DSC                    `hl7:"false,DSC"`
	Other          []interface{}
}

func (s *RTB_Z74) MessageTypeName() string {
	return "RTB_Z74"
}

// RTB_Z74_ROW_DEFINITION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type RTB_Z74_ROW_DEFINITION struct {
	RDF   *RDF  `hl7:"true,RDF"`
	RDT   []RDT `hl7:"false,RDT"`
	Other []interface{}
}

func (s *RTB_Z74_ROW_DEFINITION) MessageTypeSubStructName() string {
	return "RTB_Z74_ROW_DEFINITION"
}

// SIU_S12 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SIU_S12 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	SCH       *SCH                `hl7:"true,SCH"`
	NTE       []NTE               `hl7:"false,NTE"`
	PATIENT   []SIU_S12_PATIENT   `hl7:"false,PATIENT"`
	RESOURCES []SIU_S12_RESOURCES `hl7:"true,RESOURCES"`
	Other     []interface{}
}

func (s *SIU_S12) MessageTypeName() string {
	return "SIU_S12"
}

// SIU_S12_GENERAL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SIU_S12_GENERAL_RESOURCE struct {
	AIG   *AIG  `hl7:"true,AIG"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SIU_S12_GENERAL_RESOURCE) MessageTypeSubStructName() string {
	return "SIU_S12_GENERAL_RESOURCE"
}

// SIU_S12_LOCATION_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SIU_S12_LOCATION_RESOURCE struct {
	AIL   *AIL  `hl7:"true,AIL"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SIU_S12_LOCATION_RESOURCE) MessageTypeSubStructName() string {
	return "SIU_S12_LOCATION_RESOURCE"
}

// SIU_S12_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SIU_S12_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	PD1   *PD1  `hl7:"false,PD1"`
	PV1   *PV1  `hl7:"false,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	DG1   []DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *SIU_S12_PATIENT) MessageTypeSubStructName() string {
	return "SIU_S12_PATIENT"
}

// SIU_S12_PERSONNEL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SIU_S12_PERSONNEL_RESOURCE struct {
	AIP   *AIP  `hl7:"true,AIP"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SIU_S12_PERSONNEL_RESOURCE) MessageTypeSubStructName() string {
	return "SIU_S12_PERSONNEL_RESOURCE"
}

// SIU_S12_RESOURCES represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SIU_S12_RESOURCES struct {
	RGS                *RGS                         `hl7:"true,RGS"`
	SERVICE            []SIU_S12_SERVICE            `hl7:"false,SERVICE"`
	GENERAL_RESOURCE   []SIU_S12_GENERAL_RESOURCE   `hl7:"false,GENERAL_RESOURCE"`
	LOCATION_RESOURCE  []SIU_S12_LOCATION_RESOURCE  `hl7:"false,LOCATION_RESOURCE"`
	PERSONNEL_RESOURCE []SIU_S12_PERSONNEL_RESOURCE `hl7:"false,PERSONNEL_RESOURCE"`
	Other              []interface{}
}

func (s *SIU_S12_RESOURCES) MessageTypeSubStructName() string {
	return "SIU_S12_RESOURCES"
}

// SIU_S12_SERVICE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SIU_S12_SERVICE struct {
	AIS   *AIS  `hl7:"true,AIS"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SIU_S12_SERVICE) MessageTypeSubStructName() string {
	return "SIU_S12_SERVICE"
}

// SPQ_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type SPQ_Q01 struct {
	MSH   *MSH `hl7:"true,MSH"`
	SPR   *SPR `hl7:"true,SPR"`
	RDF   *RDF `hl7:"false,RDF"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *SPQ_Q01) MessageTypeName() string {
	return "SPQ_Q01"
}

// SPQ_Q08 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SPQ_Q08 struct {
	MSH   *MSH `hl7:"true,MSH"`
	SPR   *SPR `hl7:"true,SPR"`
	RDF   *RDF `hl7:"false,RDF"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *SPQ_Q08) MessageTypeName() string {
	return "SPQ_Q08"
}

// SQM_S25 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SQM_S25 struct {
	MSH     *MSH             `hl7:"true,MSH"`
	QRD     *QRD             `hl7:"true,QRD"`
	QRF     *QRF             `hl7:"false,QRF"`
	REQUEST *SQM_S25_REQUEST `hl7:"false,REQUEST"`
	DSC     *DSC             `hl7:"false,DSC"`
	Other   []interface{}
}

func (s *SQM_S25) MessageTypeName() string {
	return "SQM_S25"
}

// SQM_S25_GENERAL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SQM_S25_GENERAL_RESOURCE struct {
	AIG   *AIG `hl7:"true,AIG"`
	APR   *APR `hl7:"false,APR"`
	Other []interface{}
}

func (s *SQM_S25_GENERAL_RESOURCE) MessageTypeSubStructName() string {
	return "SQM_S25_GENERAL_RESOURCE"
}

// SQM_S25_LOCATION_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SQM_S25_LOCATION_RESOURCE struct {
	AIL   *AIL `hl7:"true,AIL"`
	APR   *APR `hl7:"false,APR"`
	Other []interface{}
}

func (s *SQM_S25_LOCATION_RESOURCE) MessageTypeSubStructName() string {
	return "SQM_S25_LOCATION_RESOURCE"
}

// SQM_S25_PERSONNEL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SQM_S25_PERSONNEL_RESOURCE struct {
	AIP   *AIP `hl7:"true,AIP"`
	APR   *APR `hl7:"false,APR"`
	Other []interface{}
}

func (s *SQM_S25_PERSONNEL_RESOURCE) MessageTypeSubStructName() string {
	return "SQM_S25_PERSONNEL_RESOURCE"
}

// SQM_S25_REQUEST represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SQM_S25_REQUEST struct {
	ARQ       *ARQ                `hl7:"true,ARQ"`
	APR       *APR                `hl7:"false,APR"`
	PID       *PID                `hl7:"false,PID"`
	RESOURCES []SQM_S25_RESOURCES `hl7:"true,RESOURCES"`
	Other     []interface{}
}

func (s *SQM_S25_REQUEST) MessageTypeSubStructName() string {
	return "SQM_S25_REQUEST"
}

// SQM_S25_RESOURCES represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SQM_S25_RESOURCES struct {
	RGS                *RGS                         `hl7:"true,RGS"`
	SERVICE            []SQM_S25_SERVICE            `hl7:"false,SERVICE"`
	GENERAL_RESOURCE   []SQM_S25_GENERAL_RESOURCE   `hl7:"false,GENERAL_RESOURCE"`
	PERSONNEL_RESOURCE []SQM_S25_PERSONNEL_RESOURCE `hl7:"false,PERSONNEL_RESOURCE"`
	LOCATION_RESOURCE  []SQM_S25_LOCATION_RESOURCE  `hl7:"false,LOCATION_RESOURCE"`
	Other              []interface{}
}

func (s *SQM_S25_RESOURCES) MessageTypeSubStructName() string {
	return "SQM_S25_RESOURCES"
}

// SQM_S25_SERVICE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SQM_S25_SERVICE struct {
	AIS   *AIS `hl7:"true,AIS"`
	APR   *APR `hl7:"false,APR"`
	Other []interface{}
}

func (s *SQM_S25_SERVICE) MessageTypeSubStructName() string {
	return "SQM_S25_SERVICE"
}

// SQR_S25 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SQR_S25 struct {
	MSH      *MSH               `hl7:"true,MSH"`
	MSA      *MSA               `hl7:"true,MSA"`
	ERR      *ERR               `hl7:"false,ERR"`
	QAK      *QAK               `hl7:"true,QAK"`
	SCHEDULE []SQR_S25_SCHEDULE `hl7:"false,SCHEDULE"`
	DSC      *DSC               `hl7:"false,DSC"`
	Other    []interface{}
}

func (s *SQR_S25) MessageTypeName() string {
	return "SQR_S25"
}

// SQR_S25_GENERAL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SQR_S25_GENERAL_RESOURCE struct {
	AIG   *AIG  `hl7:"true,AIG"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SQR_S25_GENERAL_RESOURCE) MessageTypeSubStructName() string {
	return "SQR_S25_GENERAL_RESOURCE"
}

// SQR_S25_LOCATION_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SQR_S25_LOCATION_RESOURCE struct {
	AIL   *AIL  `hl7:"true,AIL"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SQR_S25_LOCATION_RESOURCE) MessageTypeSubStructName() string {
	return "SQR_S25_LOCATION_RESOURCE"
}

// SQR_S25_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SQR_S25_PATIENT struct {
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"false,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	DG1   *DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *SQR_S25_PATIENT) MessageTypeSubStructName() string {
	return "SQR_S25_PATIENT"
}

// SQR_S25_PERSONNEL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SQR_S25_PERSONNEL_RESOURCE struct {
	AIP   *AIP  `hl7:"true,AIP"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SQR_S25_PERSONNEL_RESOURCE) MessageTypeSubStructName() string {
	return "SQR_S25_PERSONNEL_RESOURCE"
}

// SQR_S25_RESOURCES represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SQR_S25_RESOURCES struct {
	RGS                *RGS                         `hl7:"true,RGS"`
	SERVICE            []SQR_S25_SERVICE            `hl7:"false,SERVICE"`
	GENERAL_RESOURCE   []SQR_S25_GENERAL_RESOURCE   `hl7:"false,GENERAL_RESOURCE"`
	PERSONNEL_RESOURCE []SQR_S25_PERSONNEL_RESOURCE `hl7:"false,PERSONNEL_RESOURCE"`
	LOCATION_RESOURCE  []SQR_S25_LOCATION_RESOURCE  `hl7:"false,LOCATION_RESOURCE"`
	Other              []interface{}
}

func (s *SQR_S25_RESOURCES) MessageTypeSubStructName() string {
	return "SQR_S25_RESOURCES"
}

// SQR_S25_SCHEDULE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SQR_S25_SCHEDULE struct {
	SCH       *SCH                `hl7:"true,SCH"`
	NTE       []NTE               `hl7:"false,NTE"`
	PATIENT   *SQR_S25_PATIENT    `hl7:"false,PATIENT"`
	RESOURCES []SQR_S25_RESOURCES `hl7:"true,RESOURCES"`
	Other     []interface{}
}

func (s *SQR_S25_SCHEDULE) MessageTypeSubStructName() string {
	return "SQR_S25_SCHEDULE"
}

// SQR_S25_SERVICE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SQR_S25_SERVICE struct {
	AIS   *AIS  `hl7:"true,AIS"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SQR_S25_SERVICE) MessageTypeSubStructName() string {
	return "SQR_S25_SERVICE"
}

// SRM_S01 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SRM_S01 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	ARQ       *ARQ                `hl7:"true,ARQ"`
	APR       *APR                `hl7:"false,APR"`
	NTE       []NTE               `hl7:"false,NTE"`
	PATIENT   []SRM_S01_PATIENT   `hl7:"false,PATIENT"`
	RESOURCES []SRM_S01_RESOURCES `hl7:"true,RESOURCES"`
	Other     []interface{}
}

func (s *SRM_S01) MessageTypeName() string {
	return "SRM_S01"
}

// SRM_S01_GENERAL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SRM_S01_GENERAL_RESOURCE struct {
	AIG   *AIG  `hl7:"true,AIG"`
	APR   *APR  `hl7:"false,APR"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SRM_S01_GENERAL_RESOURCE) MessageTypeSubStructName() string {
	return "SRM_S01_GENERAL_RESOURCE"
}

// SRM_S01_LOCATION_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SRM_S01_LOCATION_RESOURCE struct {
	AIL   *AIL  `hl7:"true,AIL"`
	APR   *APR  `hl7:"false,APR"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SRM_S01_LOCATION_RESOURCE) MessageTypeSubStructName() string {
	return "SRM_S01_LOCATION_RESOURCE"
}

// SRM_S01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SRM_S01_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"false,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	OBX   []OBX `hl7:"false,OBX"`
	DG1   []DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *SRM_S01_PATIENT) MessageTypeSubStructName() string {
	return "SRM_S01_PATIENT"
}

// SRM_S01_PERSONNEL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SRM_S01_PERSONNEL_RESOURCE struct {
	AIP   *AIP  `hl7:"true,AIP"`
	APR   *APR  `hl7:"false,APR"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SRM_S01_PERSONNEL_RESOURCE) MessageTypeSubStructName() string {
	return "SRM_S01_PERSONNEL_RESOURCE"
}

// SRM_S01_RESOURCES represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SRM_S01_RESOURCES struct {
	RGS                *RGS                         `hl7:"true,RGS"`
	SERVICE            []SRM_S01_SERVICE            `hl7:"false,SERVICE"`
	GENERAL_RESOURCE   []SRM_S01_GENERAL_RESOURCE   `hl7:"false,GENERAL_RESOURCE"`
	LOCATION_RESOURCE  []SRM_S01_LOCATION_RESOURCE  `hl7:"false,LOCATION_RESOURCE"`
	PERSONNEL_RESOURCE []SRM_S01_PERSONNEL_RESOURCE `hl7:"false,PERSONNEL_RESOURCE"`
	Other              []interface{}
}

func (s *SRM_S01_RESOURCES) MessageTypeSubStructName() string {
	return "SRM_S01_RESOURCES"
}

// SRM_S01_SERVICE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SRM_S01_SERVICE struct {
	AIS   *AIS  `hl7:"true,AIS"`
	APR   *APR  `hl7:"false,APR"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SRM_S01_SERVICE) MessageTypeSubStructName() string {
	return "SRM_S01_SERVICE"
}

// SRR_S01 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SRR_S01 struct {
	MSH      *MSH              `hl7:"true,MSH"`
	MSA      *MSA              `hl7:"true,MSA"`
	ERR      *ERR              `hl7:"false,ERR"`
	SCHEDULE *SRR_S01_SCHEDULE `hl7:"false,SCHEDULE"`
	Other    []interface{}
}

func (s *SRR_S01) MessageTypeName() string {
	return "SRR_S01"
}

// SRR_S01_GENERAL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SRR_S01_GENERAL_RESOURCE struct {
	AIG   *AIG  `hl7:"true,AIG"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SRR_S01_GENERAL_RESOURCE) MessageTypeSubStructName() string {
	return "SRR_S01_GENERAL_RESOURCE"
}

// SRR_S01_LOCATION_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SRR_S01_LOCATION_RESOURCE struct {
	AIL   *AIL  `hl7:"true,AIL"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SRR_S01_LOCATION_RESOURCE) MessageTypeSubStructName() string {
	return "SRR_S01_LOCATION_RESOURCE"
}

// SRR_S01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SRR_S01_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"false,PV1"`
	PV2   *PV2  `hl7:"false,PV2"`
	DG1   []DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *SRR_S01_PATIENT) MessageTypeSubStructName() string {
	return "SRR_S01_PATIENT"
}

// SRR_S01_PERSONNEL_RESOURCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SRR_S01_PERSONNEL_RESOURCE struct {
	AIP   *AIP  `hl7:"true,AIP"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SRR_S01_PERSONNEL_RESOURCE) MessageTypeSubStructName() string {
	return "SRR_S01_PERSONNEL_RESOURCE"
}

// SRR_S01_RESOURCES represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SRR_S01_RESOURCES struct {
	RGS                *RGS                         `hl7:"true,RGS"`
	SERVICE            []SRR_S01_SERVICE            `hl7:"false,SERVICE"`
	GENERAL_RESOURCE   []SRR_S01_GENERAL_RESOURCE   `hl7:"false,GENERAL_RESOURCE"`
	LOCATION_RESOURCE  []SRR_S01_LOCATION_RESOURCE  `hl7:"false,LOCATION_RESOURCE"`
	PERSONNEL_RESOURCE []SRR_S01_PERSONNEL_RESOURCE `hl7:"false,PERSONNEL_RESOURCE"`
	Other              []interface{}
}

func (s *SRR_S01_RESOURCES) MessageTypeSubStructName() string {
	return "SRR_S01_RESOURCES"
}

// SRR_S01_SCHEDULE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SRR_S01_SCHEDULE struct {
	SCH       *SCH                `hl7:"true,SCH"`
	NTE       []NTE               `hl7:"false,NTE"`
	PATIENT   []SRR_S01_PATIENT   `hl7:"false,PATIENT"`
	RESOURCES []SRR_S01_RESOURCES `hl7:"true,RESOURCES"`
	Other     []interface{}
}

func (s *SRR_S01_SCHEDULE) MessageTypeSubStructName() string {
	return "SRR_S01_SCHEDULE"
}

// SRR_S01_SERVICE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SRR_S01_SERVICE struct {
	AIS   *AIS  `hl7:"true,AIS"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *SRR_S01_SERVICE) MessageTypeSubStructName() string {
	return "SRR_S01_SERVICE"
}

// SSR_U04 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SSR_U04 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EQU   *EQU  `hl7:"true,EQU"`
	SAC   []SAC `hl7:"true,SAC"`
	ROL   *ROL  `hl7:"false,ROL"`
	Other []interface{}
}

func (s *SSR_U04) MessageTypeName() string {
	return "SSR_U04"
}

// SSU_U03 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SSU_U03 struct {
	MSH                *MSH                         `hl7:"true,MSH"`
	EQU                *EQU                         `hl7:"true,EQU"`
	SPECIMEN_CONTAINER []SSU_U03_SPECIMEN_CONTAINER `hl7:"true,SPECIMEN_CONTAINER"`
	ROL                *ROL                         `hl7:"false,ROL"`
	Other              []interface{}
}

func (s *SSU_U03) MessageTypeName() string {
	return "SSU_U03"
}

// SSU_U03_SPECIMEN_CONTAINER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SSU_U03_SPECIMEN_CONTAINER struct {
	SAC   *SAC `hl7:"true,SAC"`
	OBX   *OBX `hl7:"false,OBX"`
	Other []interface{}
}

func (s *SSU_U03_SPECIMEN_CONTAINER) MessageTypeSubStructName() string {
	return "SSU_U03_SPECIMEN_CONTAINER"
}

// SUR_P09 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SUR_P09 struct {
	MSH      *MSH               `hl7:"true,MSH"`
	FACILITY []SUR_P09_FACILITY `hl7:"true,FACILITY"`
	Other    []interface{}
}

func (s *SUR_P09) MessageTypeName() string {
	return "SUR_P09"
}

// SUR_P09_FACILITY represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SUR_P09_FACILITY struct {
	FAC             *FAC                      `hl7:"true,FAC"`
	PRODUCT         []SUR_P09_PRODUCT         `hl7:"true,PRODUCT"`
	PSH             *PSH                      `hl7:"true,PSH"`
	FACILITY_DETAIL []SUR_P09_FACILITY_DETAIL `hl7:"true,FACILITY_DETAIL"`
	Other           []interface{}
}

func (s *SUR_P09_FACILITY) MessageTypeSubStructName() string {
	return "SUR_P09_FACILITY"
}

// SUR_P09_FACILITY_DETAIL represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SUR_P09_FACILITY_DETAIL struct {
	FAC   *FAC `hl7:"true,FAC"`
	PDC   *PDC `hl7:"true,PDC"`
	NTE   *NTE `hl7:"true,NTE"`
	Other []interface{}
}

func (s *SUR_P09_FACILITY_DETAIL) MessageTypeSubStructName() string {
	return "SUR_P09_FACILITY_DETAIL"
}

// SUR_P09_PRODUCT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type SUR_P09_PRODUCT struct {
	PSH   *PSH `hl7:"true,PSH"`
	PDC   *PDC `hl7:"true,PDC"`
	Other []interface{}
}

func (s *SUR_P09_PRODUCT) MessageTypeSubStructName() string {
	return "SUR_P09_PRODUCT"
}

// TBR_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type TBR_Q01 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MSA   *MSA  `hl7:"true,MSA"`
	ERR   *ERR  `hl7:"false,ERR"`
	QAK   *QAK  `hl7:"true,QAK"`
	RDF   *RDF  `hl7:"true,RDF"`
	RDT   []RDT `hl7:"true,RDT"`
	DSC   *DSC  `hl7:"false,DSC"`
	Other []interface{}
}

func (s *TBR_Q01) MessageTypeName() string {
	return "TBR_Q01"
}

// TBR_R08 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type TBR_R08 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MSA   *MSA  `hl7:"true,MSA"`
	ERR   *ERR  `hl7:"false,ERR"`
	QAK   *QAK  `hl7:"true,QAK"`
	RDF   *RDF  `hl7:"true,RDF"`
	RDT   []RDT `hl7:"true,RDT"`
	DSC   *DSC  `hl7:"false,DSC"`
	Other []interface{}
}

func (s *TBR_R08) MessageTypeName() string {
	return "TBR_R08"
}

// TCU_U10 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type TCU_U10 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EQU   *EQU  `hl7:"true,EQU"`
	TCC   []TCC `hl7:"true,TCC"`
	ROL   *ROL  `hl7:"false,ROL"`
	Other []interface{}
}

func (s *TCU_U10) MessageTypeName() string {
	return "TCU_U10"
}

// UDM_Q05 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type UDM_Q05 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	URD   *URD  `hl7:"true,URD"`
	URS   *URS  `hl7:"false,URS"`
	DSP   []DSP `hl7:"true,DSP"`
	DSC   *DSC  `hl7:"false,DSC"`
	Other []interface{}
}

func (s *UDM_Q05) MessageTypeName() string {
	return "UDM_Q05"
}

// VQQ_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.3
type VQQ_Q01 struct {
	MSH   *MSH `hl7:"true,MSH"`
	VTQ   *VTQ `hl7:"true,VTQ"`
	RDF   *RDF `hl7:"false,RDF"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *VQQ_Q01) MessageTypeName() string {
	return "VQQ_Q01"
}

// VQQ_Q07 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type VQQ_Q07 struct {
	MSH   *MSH `hl7:"true,MSH"`
	VTQ   *VTQ `hl7:"true,VTQ"`
	RDF   *RDF `hl7:"false,RDF"`
	DSC   *DSC `hl7:"false,DSC"`
	Other []interface{}
}

func (s *VQQ_Q07) MessageTypeName() string {
	return "VQQ_Q07"
}

// VXQ_V01 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type VXQ_V01 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QRD   *QRD `hl7:"true,QRD"`
	QRF   *QRF `hl7:"false,QRF"`
	Other []interface{}
}

func (s *VXQ_V01) MessageTypeName() string {
	return "VXQ_V01"
}

// VXR_V03 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type VXR_V03 struct {
	MSH           *MSH                   `hl7:"true,MSH"`
	MSA           *MSA                   `hl7:"true,MSA"`
	QRD           *QRD                   `hl7:"true,QRD"`
	QRF           *QRF                   `hl7:"false,QRF"`
	PID           *PID                   `hl7:"true,PID"`
	PD1           *PD1                   `hl7:"false,PD1"`
	NK1           []NK1                  `hl7:"false,NK1"`
	PATIENT_VISIT *VXR_V03_PATIENT_VISIT `hl7:"false,PATIENT_VISIT"`
	GT1           []GT1                  `hl7:"false,GT1"`
	INSURANCE     []VXR_V03_INSURANCE    `hl7:"false,INSURANCE"`
	ORDER         []VXR_V03_ORDER        `hl7:"false,ORDER"`
	Other         []interface{}
}

func (s *VXR_V03) MessageTypeName() string {
	return "VXR_V03"
}

// VXR_V03_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type VXR_V03_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *VXR_V03_INSURANCE) MessageTypeSubStructName() string {
	return "VXR_V03_INSURANCE"
}

// VXR_V03_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type VXR_V03_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *VXR_V03_OBSERVATION) MessageTypeSubStructName() string {
	return "VXR_V03_OBSERVATION"
}

// VXR_V03_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type VXR_V03_ORDER struct {
	ORC         *ORC                  `hl7:"false,ORC"`
	RXA         *RXA                  `hl7:"true,RXA"`
	RXR         *RXR                  `hl7:"false,RXR"`
	OBSERVATION []VXR_V03_OBSERVATION `hl7:"false,OBSERVATION"`
	Other       []interface{}
}

func (s *VXR_V03_ORDER) MessageTypeSubStructName() string {
	return "VXR_V03_ORDER"
}

// VXR_V03_PATIENT_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type VXR_V03_PATIENT_VISIT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *VXR_V03_PATIENT_VISIT) MessageTypeSubStructName() string {
	return "VXR_V03_PATIENT_VISIT"
}

// VXU_V04 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type VXU_V04 struct {
	MSH       *MSH                `hl7:"true,MSH"`
	PID       *PID                `hl7:"true,PID"`
	PD1       *PD1                `hl7:"false,PD1"`
	NK1       []NK1               `hl7:"false,NK1"`
	PATIENT   *VXU_V04_PATIENT    `hl7:"false,PATIENT"`
	GT1       []GT1               `hl7:"false,GT1"`
	INSURANCE []VXU_V04_INSURANCE `hl7:"false,INSURANCE"`
	ORDER     []VXU_V04_ORDER     `hl7:"false,ORDER"`
	Other     []interface{}
}

func (s *VXU_V04) MessageTypeName() string {
	return "VXU_V04"
}

// VXU_V04_INSURANCE represents the corresponding HL7 message type.
// Definition from HL7 2.4
type VXU_V04_INSURANCE struct {
	IN1   *IN1 `hl7:"true,IN1"`
	IN2   *IN2 `hl7:"false,IN2"`
	IN3   *IN3 `hl7:"false,IN3"`
	Other []interface{}
}

func (s *VXU_V04_INSURANCE) MessageTypeSubStructName() string {
	return "VXU_V04_INSURANCE"
}

// VXU_V04_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.4
type VXU_V04_OBSERVATION struct {
	OBX   *OBX  `hl7:"true,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *VXU_V04_OBSERVATION) MessageTypeSubStructName() string {
	return "VXU_V04_OBSERVATION"
}

// VXU_V04_ORDER represents the corresponding HL7 message type.
// Definition from HL7 2.4
type VXU_V04_ORDER struct {
	ORC         *ORC                  `hl7:"false,ORC"`
	RXA         *RXA                  `hl7:"true,RXA"`
	RXR         *RXR                  `hl7:"false,RXR"`
	OBSERVATION []VXU_V04_OBSERVATION `hl7:"false,OBSERVATION"`
	Other       []interface{}
}

func (s *VXU_V04_ORDER) MessageTypeSubStructName() string {
	return "VXU_V04_ORDER"
}

// VXU_V04_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type VXU_V04_PATIENT struct {
	PV1   *PV1 `hl7:"true,PV1"`
	PV2   *PV2 `hl7:"false,PV2"`
	Other []interface{}
}

func (s *VXU_V04_PATIENT) MessageTypeSubStructName() string {
	return "VXU_V04_PATIENT"
}

// VXX_V02 represents the corresponding HL7 message type.
// Definition from HL7 2.4
type VXX_V02 struct {
	MSH     *MSH              `hl7:"true,MSH"`
	MSA     *MSA              `hl7:"true,MSA"`
	QRD     *QRD              `hl7:"true,QRD"`
	QRF     *QRF              `hl7:"false,QRF"`
	PATIENT []VXX_V02_PATIENT `hl7:"true,PATIENT"`
	Other   []interface{}
}

func (s *VXX_V02) MessageTypeName() string {
	return "VXX_V02"
}

// VXX_V02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.4
type VXX_V02_PATIENT struct {
	PID   *PID  `hl7:"true,PID"`
	NK1   []NK1 `hl7:"false,NK1"`
	Other []interface{}
}

func (s *VXX_V02_PATIENT) MessageTypeSubStructName() string {
	return "VXX_V02_PATIENT"
}

// GenericHL7Segment represents the corresponding HL7 segment type.
type GenericHL7Segment struct {
	segment []byte
}

func (s *GenericHL7Segment) SegmentName() string {
	return "GenericHL7Segment"
}

// ABS returns the first ABS segment within the message, or nil if there isn't one.
func (m *Message) ABS() (*ABS, error) {
	ps, err := m.parse("ABS")
	pst, ok := ps.(*ABS)
	if ok {
		return pst, err
	}
	return nil, err
}

// ACC returns the first ACC segment within the message, or nil if there isn't one.
func (m *Message) ACC() (*ACC, error) {
	ps, err := m.parse("ACC")
	pst, ok := ps.(*ACC)
	if ok {
		return pst, err
	}
	return nil, err
}

// ADD returns the first ADD segment within the message, or nil if there isn't one.
func (m *Message) ADD() (*ADD, error) {
	ps, err := m.parse("ADD")
	pst, ok := ps.(*ADD)
	if ok {
		return pst, err
	}
	return nil, err
}

// AFF returns the first AFF segment within the message, or nil if there isn't one.
func (m *Message) AFF() (*AFF, error) {
	ps, err := m.parse("AFF")
	pst, ok := ps.(*AFF)
	if ok {
		return pst, err
	}
	return nil, err
}

// AIG returns the first AIG segment within the message, or nil if there isn't one.
func (m *Message) AIG() (*AIG, error) {
	ps, err := m.parse("AIG")
	pst, ok := ps.(*AIG)
	if ok {
		return pst, err
	}
	return nil, err
}

// AIL returns the first AIL segment within the message, or nil if there isn't one.
func (m *Message) AIL() (*AIL, error) {
	ps, err := m.parse("AIL")
	pst, ok := ps.(*AIL)
	if ok {
		return pst, err
	}
	return nil, err
}

// AIP returns the first AIP segment within the message, or nil if there isn't one.
func (m *Message) AIP() (*AIP, error) {
	ps, err := m.parse("AIP")
	pst, ok := ps.(*AIP)
	if ok {
		return pst, err
	}
	return nil, err
}

// AIS returns the first AIS segment within the message, or nil if there isn't one.
func (m *Message) AIS() (*AIS, error) {
	ps, err := m.parse("AIS")
	pst, ok := ps.(*AIS)
	if ok {
		return pst, err
	}
	return nil, err
}

// AL1 returns the first AL1 segment within the message, or nil if there isn't one.
func (m *Message) AL1() (*AL1, error) {
	ps, err := m.parse("AL1")
	pst, ok := ps.(*AL1)
	if ok {
		return pst, err
	}
	return nil, err
}

// APR returns the first APR segment within the message, or nil if there isn't one.
func (m *Message) APR() (*APR, error) {
	ps, err := m.parse("APR")
	pst, ok := ps.(*APR)
	if ok {
		return pst, err
	}
	return nil, err
}

// ARQ returns the first ARQ segment within the message, or nil if there isn't one.
func (m *Message) ARQ() (*ARQ, error) {
	ps, err := m.parse("ARQ")
	pst, ok := ps.(*ARQ)
	if ok {
		return pst, err
	}
	return nil, err
}

// AUT returns the first AUT segment within the message, or nil if there isn't one.
func (m *Message) AUT() (*AUT, error) {
	ps, err := m.parse("AUT")
	pst, ok := ps.(*AUT)
	if ok {
		return pst, err
	}
	return nil, err
}

// BHS returns the first BHS segment within the message, or nil if there isn't one.
func (m *Message) BHS() (*BHS, error) {
	ps, err := m.parse("BHS")
	pst, ok := ps.(*BHS)
	if ok {
		return pst, err
	}
	return nil, err
}

// BLC returns the first BLC segment within the message, or nil if there isn't one.
func (m *Message) BLC() (*BLC, error) {
	ps, err := m.parse("BLC")
	pst, ok := ps.(*BLC)
	if ok {
		return pst, err
	}
	return nil, err
}

// BLG returns the first BLG segment within the message, or nil if there isn't one.
func (m *Message) BLG() (*BLG, error) {
	ps, err := m.parse("BLG")
	pst, ok := ps.(*BLG)
	if ok {
		return pst, err
	}
	return nil, err
}

// BTS returns the first BTS segment within the message, or nil if there isn't one.
func (m *Message) BTS() (*BTS, error) {
	ps, err := m.parse("BTS")
	pst, ok := ps.(*BTS)
	if ok {
		return pst, err
	}
	return nil, err
}

// CDM returns the first CDM segment within the message, or nil if there isn't one.
func (m *Message) CDM() (*CDM, error) {
	ps, err := m.parse("CDM")
	pst, ok := ps.(*CDM)
	if ok {
		return pst, err
	}
	return nil, err
}

// CM0 returns the first CM0 segment within the message, or nil if there isn't one.
func (m *Message) CM0() (*CM0, error) {
	ps, err := m.parse("CM0")
	pst, ok := ps.(*CM0)
	if ok {
		return pst, err
	}
	return nil, err
}

// CM1 returns the first CM1 segment within the message, or nil if there isn't one.
func (m *Message) CM1() (*CM1, error) {
	ps, err := m.parse("CM1")
	pst, ok := ps.(*CM1)
	if ok {
		return pst, err
	}
	return nil, err
}

// CM2 returns the first CM2 segment within the message, or nil if there isn't one.
func (m *Message) CM2() (*CM2, error) {
	ps, err := m.parse("CM2")
	pst, ok := ps.(*CM2)
	if ok {
		return pst, err
	}
	return nil, err
}

// CNS returns the first CNS segment within the message, or nil if there isn't one.
func (m *Message) CNS() (*CNS, error) {
	ps, err := m.parse("CNS")
	pst, ok := ps.(*CNS)
	if ok {
		return pst, err
	}
	return nil, err
}

// CSP returns the first CSP segment within the message, or nil if there isn't one.
func (m *Message) CSP() (*CSP, error) {
	ps, err := m.parse("CSP")
	pst, ok := ps.(*CSP)
	if ok {
		return pst, err
	}
	return nil, err
}

// CSR returns the first CSR segment within the message, or nil if there isn't one.
func (m *Message) CSR() (*CSR, error) {
	ps, err := m.parse("CSR")
	pst, ok := ps.(*CSR)
	if ok {
		return pst, err
	}
	return nil, err
}

// CSS returns the first CSS segment within the message, or nil if there isn't one.
func (m *Message) CSS() (*CSS, error) {
	ps, err := m.parse("CSS")
	pst, ok := ps.(*CSS)
	if ok {
		return pst, err
	}
	return nil, err
}

// CTD returns the first CTD segment within the message, or nil if there isn't one.
func (m *Message) CTD() (*CTD, error) {
	ps, err := m.parse("CTD")
	pst, ok := ps.(*CTD)
	if ok {
		return pst, err
	}
	return nil, err
}

// CTI returns the first CTI segment within the message, or nil if there isn't one.
func (m *Message) CTI() (*CTI, error) {
	ps, err := m.parse("CTI")
	pst, ok := ps.(*CTI)
	if ok {
		return pst, err
	}
	return nil, err
}

// DB1 returns the first DB1 segment within the message, or nil if there isn't one.
func (m *Message) DB1() (*DB1, error) {
	ps, err := m.parse("DB1")
	pst, ok := ps.(*DB1)
	if ok {
		return pst, err
	}
	return nil, err
}

// DG1 returns the first DG1 segment within the message, or nil if there isn't one.
func (m *Message) DG1() (*DG1, error) {
	ps, err := m.parse("DG1")
	pst, ok := ps.(*DG1)
	if ok {
		return pst, err
	}
	return nil, err
}

// DRG returns the first DRG segment within the message, or nil if there isn't one.
func (m *Message) DRG() (*DRG, error) {
	ps, err := m.parse("DRG")
	pst, ok := ps.(*DRG)
	if ok {
		return pst, err
	}
	return nil, err
}

// DSC returns the first DSC segment within the message, or nil if there isn't one.
func (m *Message) DSC() (*DSC, error) {
	ps, err := m.parse("DSC")
	pst, ok := ps.(*DSC)
	if ok {
		return pst, err
	}
	return nil, err
}

// DSP returns the first DSP segment within the message, or nil if there isn't one.
func (m *Message) DSP() (*DSP, error) {
	ps, err := m.parse("DSP")
	pst, ok := ps.(*DSP)
	if ok {
		return pst, err
	}
	return nil, err
}

// ECD returns the first ECD segment within the message, or nil if there isn't one.
func (m *Message) ECD() (*ECD, error) {
	ps, err := m.parse("ECD")
	pst, ok := ps.(*ECD)
	if ok {
		return pst, err
	}
	return nil, err
}

// ECR returns the first ECR segment within the message, or nil if there isn't one.
func (m *Message) ECR() (*ECR, error) {
	ps, err := m.parse("ECR")
	pst, ok := ps.(*ECR)
	if ok {
		return pst, err
	}
	return nil, err
}

// EDU returns the first EDU segment within the message, or nil if there isn't one.
func (m *Message) EDU() (*EDU, error) {
	ps, err := m.parse("EDU")
	pst, ok := ps.(*EDU)
	if ok {
		return pst, err
	}
	return nil, err
}

// EQL returns the first EQL segment within the message, or nil if there isn't one.
func (m *Message) EQL() (*EQL, error) {
	ps, err := m.parse("EQL")
	pst, ok := ps.(*EQL)
	if ok {
		return pst, err
	}
	return nil, err
}

// EQP returns the first EQP segment within the message, or nil if there isn't one.
func (m *Message) EQP() (*EQP, error) {
	ps, err := m.parse("EQP")
	pst, ok := ps.(*EQP)
	if ok {
		return pst, err
	}
	return nil, err
}

// EQU returns the first EQU segment within the message, or nil if there isn't one.
func (m *Message) EQU() (*EQU, error) {
	ps, err := m.parse("EQU")
	pst, ok := ps.(*EQU)
	if ok {
		return pst, err
	}
	return nil, err
}

// ERQ returns the first ERQ segment within the message, or nil if there isn't one.
func (m *Message) ERQ() (*ERQ, error) {
	ps, err := m.parse("ERQ")
	pst, ok := ps.(*ERQ)
	if ok {
		return pst, err
	}
	return nil, err
}

// ERR returns the first ERR segment within the message, or nil if there isn't one.
func (m *Message) ERR() (*ERR, error) {
	ps, err := m.parse("ERR")
	pst, ok := ps.(*ERR)
	if ok {
		return pst, err
	}
	return nil, err
}

// EVN returns the first EVN segment within the message, or nil if there isn't one.
func (m *Message) EVN() (*EVN, error) {
	ps, err := m.parse("EVN")
	pst, ok := ps.(*EVN)
	if ok {
		return pst, err
	}
	return nil, err
}

// FAC returns the first FAC segment within the message, or nil if there isn't one.
func (m *Message) FAC() (*FAC, error) {
	ps, err := m.parse("FAC")
	pst, ok := ps.(*FAC)
	if ok {
		return pst, err
	}
	return nil, err
}

// FHS returns the first FHS segment within the message, or nil if there isn't one.
func (m *Message) FHS() (*FHS, error) {
	ps, err := m.parse("FHS")
	pst, ok := ps.(*FHS)
	if ok {
		return pst, err
	}
	return nil, err
}

// FT1 returns the first FT1 segment within the message, or nil if there isn't one.
func (m *Message) FT1() (*FT1, error) {
	ps, err := m.parse("FT1")
	pst, ok := ps.(*FT1)
	if ok {
		return pst, err
	}
	return nil, err
}

// FTS returns the first FTS segment within the message, or nil if there isn't one.
func (m *Message) FTS() (*FTS, error) {
	ps, err := m.parse("FTS")
	pst, ok := ps.(*FTS)
	if ok {
		return pst, err
	}
	return nil, err
}

// GOL returns the first GOL segment within the message, or nil if there isn't one.
func (m *Message) GOL() (*GOL, error) {
	ps, err := m.parse("GOL")
	pst, ok := ps.(*GOL)
	if ok {
		return pst, err
	}
	return nil, err
}

// GP1 returns the first GP1 segment within the message, or nil if there isn't one.
func (m *Message) GP1() (*GP1, error) {
	ps, err := m.parse("GP1")
	pst, ok := ps.(*GP1)
	if ok {
		return pst, err
	}
	return nil, err
}

// GP2 returns the first GP2 segment within the message, or nil if there isn't one.
func (m *Message) GP2() (*GP2, error) {
	ps, err := m.parse("GP2")
	pst, ok := ps.(*GP2)
	if ok {
		return pst, err
	}
	return nil, err
}

// GT1 returns the first GT1 segment within the message, or nil if there isn't one.
func (m *Message) GT1() (*GT1, error) {
	ps, err := m.parse("GT1")
	pst, ok := ps.(*GT1)
	if ok {
		return pst, err
	}
	return nil, err
}

// IAM returns the first IAM segment within the message, or nil if there isn't one.
func (m *Message) IAM() (*IAM, error) {
	ps, err := m.parse("IAM")
	pst, ok := ps.(*IAM)
	if ok {
		return pst, err
	}
	return nil, err
}

// IN1 returns the first IN1 segment within the message, or nil if there isn't one.
func (m *Message) IN1() (*IN1, error) {
	ps, err := m.parse("IN1")
	pst, ok := ps.(*IN1)
	if ok {
		return pst, err
	}
	return nil, err
}

// IN2 returns the first IN2 segment within the message, or nil if there isn't one.
func (m *Message) IN2() (*IN2, error) {
	ps, err := m.parse("IN2")
	pst, ok := ps.(*IN2)
	if ok {
		return pst, err
	}
	return nil, err
}

// IN3 returns the first IN3 segment within the message, or nil if there isn't one.
func (m *Message) IN3() (*IN3, error) {
	ps, err := m.parse("IN3")
	pst, ok := ps.(*IN3)
	if ok {
		return pst, err
	}
	return nil, err
}

// INV returns the first INV segment within the message, or nil if there isn't one.
func (m *Message) INV() (*INV, error) {
	ps, err := m.parse("INV")
	pst, ok := ps.(*INV)
	if ok {
		return pst, err
	}
	return nil, err
}

// ISD returns the first ISD segment within the message, or nil if there isn't one.
func (m *Message) ISD() (*ISD, error) {
	ps, err := m.parse("ISD")
	pst, ok := ps.(*ISD)
	if ok {
		return pst, err
	}
	return nil, err
}

// LAN returns the first LAN segment within the message, or nil if there isn't one.
func (m *Message) LAN() (*LAN, error) {
	ps, err := m.parse("LAN")
	pst, ok := ps.(*LAN)
	if ok {
		return pst, err
	}
	return nil, err
}

// LCC returns the first LCC segment within the message, or nil if there isn't one.
func (m *Message) LCC() (*LCC, error) {
	ps, err := m.parse("LCC")
	pst, ok := ps.(*LCC)
	if ok {
		return pst, err
	}
	return nil, err
}

// LCH returns the first LCH segment within the message, or nil if there isn't one.
func (m *Message) LCH() (*LCH, error) {
	ps, err := m.parse("LCH")
	pst, ok := ps.(*LCH)
	if ok {
		return pst, err
	}
	return nil, err
}

// LDP returns the first LDP segment within the message, or nil if there isn't one.
func (m *Message) LDP() (*LDP, error) {
	ps, err := m.parse("LDP")
	pst, ok := ps.(*LDP)
	if ok {
		return pst, err
	}
	return nil, err
}

// LOC returns the first LOC segment within the message, or nil if there isn't one.
func (m *Message) LOC() (*LOC, error) {
	ps, err := m.parse("LOC")
	pst, ok := ps.(*LOC)
	if ok {
		return pst, err
	}
	return nil, err
}

// LRL returns the first LRL segment within the message, or nil if there isn't one.
func (m *Message) LRL() (*LRL, error) {
	ps, err := m.parse("LRL")
	pst, ok := ps.(*LRL)
	if ok {
		return pst, err
	}
	return nil, err
}

// MFA returns the first MFA segment within the message, or nil if there isn't one.
func (m *Message) MFA() (*MFA, error) {
	ps, err := m.parse("MFA")
	pst, ok := ps.(*MFA)
	if ok {
		return pst, err
	}
	return nil, err
}

// MFE returns the first MFE segment within the message, or nil if there isn't one.
func (m *Message) MFE() (*MFE, error) {
	ps, err := m.parse("MFE")
	pst, ok := ps.(*MFE)
	if ok {
		return pst, err
	}
	return nil, err
}

// MFI returns the first MFI segment within the message, or nil if there isn't one.
func (m *Message) MFI() (*MFI, error) {
	ps, err := m.parse("MFI")
	pst, ok := ps.(*MFI)
	if ok {
		return pst, err
	}
	return nil, err
}

// MRG returns the first MRG segment within the message, or nil if there isn't one.
func (m *Message) MRG() (*MRG, error) {
	ps, err := m.parse("MRG")
	pst, ok := ps.(*MRG)
	if ok {
		return pst, err
	}
	return nil, err
}

// MSA returns the first MSA segment within the message, or nil if there isn't one.
func (m *Message) MSA() (*MSA, error) {
	ps, err := m.parse("MSA")
	pst, ok := ps.(*MSA)
	if ok {
		return pst, err
	}
	return nil, err
}

// MSH returns the first MSH segment within the message, or nil if there isn't one.
func (m *Message) MSH() (*MSH, error) {
	ps, err := m.parse("MSH")
	pst, ok := ps.(*MSH)
	if ok {
		return pst, err
	}
	return nil, err
}

// NCK returns the first NCK segment within the message, or nil if there isn't one.
func (m *Message) NCK() (*NCK, error) {
	ps, err := m.parse("NCK")
	pst, ok := ps.(*NCK)
	if ok {
		return pst, err
	}
	return nil, err
}

// NDS returns the first NDS segment within the message, or nil if there isn't one.
func (m *Message) NDS() (*NDS, error) {
	ps, err := m.parse("NDS")
	pst, ok := ps.(*NDS)
	if ok {
		return pst, err
	}
	return nil, err
}

// NK1 returns the first NK1 segment within the message, or nil if there isn't one.
func (m *Message) NK1() (*NK1, error) {
	ps, err := m.parse("NK1")
	pst, ok := ps.(*NK1)
	if ok {
		return pst, err
	}
	return nil, err
}

// NPU returns the first NPU segment within the message, or nil if there isn't one.
func (m *Message) NPU() (*NPU, error) {
	ps, err := m.parse("NPU")
	pst, ok := ps.(*NPU)
	if ok {
		return pst, err
	}
	return nil, err
}

// NSC returns the first NSC segment within the message, or nil if there isn't one.
func (m *Message) NSC() (*NSC, error) {
	ps, err := m.parse("NSC")
	pst, ok := ps.(*NSC)
	if ok {
		return pst, err
	}
	return nil, err
}

// NST returns the first NST segment within the message, or nil if there isn't one.
func (m *Message) NST() (*NST, error) {
	ps, err := m.parse("NST")
	pst, ok := ps.(*NST)
	if ok {
		return pst, err
	}
	return nil, err
}

// NTE returns the first NTE segment within the message, or nil if there isn't one.
func (m *Message) NTE() (*NTE, error) {
	ps, err := m.parse("NTE")
	pst, ok := ps.(*NTE)
	if ok {
		return pst, err
	}
	return nil, err
}

// OBR returns the first OBR segment within the message, or nil if there isn't one.
func (m *Message) OBR() (*OBR, error) {
	ps, err := m.parse("OBR")
	pst, ok := ps.(*OBR)
	if ok {
		return pst, err
	}
	return nil, err
}

// OBX returns the first OBX segment within the message, or nil if there isn't one.
func (m *Message) OBX() (*OBX, error) {
	ps, err := m.parse("OBX")
	pst, ok := ps.(*OBX)
	if ok {
		return pst, err
	}
	return nil, err
}

// ODS returns the first ODS segment within the message, or nil if there isn't one.
func (m *Message) ODS() (*ODS, error) {
	ps, err := m.parse("ODS")
	pst, ok := ps.(*ODS)
	if ok {
		return pst, err
	}
	return nil, err
}

// ODT returns the first ODT segment within the message, or nil if there isn't one.
func (m *Message) ODT() (*ODT, error) {
	ps, err := m.parse("ODT")
	pst, ok := ps.(*ODT)
	if ok {
		return pst, err
	}
	return nil, err
}

// OM1 returns the first OM1 segment within the message, or nil if there isn't one.
func (m *Message) OM1() (*OM1, error) {
	ps, err := m.parse("OM1")
	pst, ok := ps.(*OM1)
	if ok {
		return pst, err
	}
	return nil, err
}

// OM2 returns the first OM2 segment within the message, or nil if there isn't one.
func (m *Message) OM2() (*OM2, error) {
	ps, err := m.parse("OM2")
	pst, ok := ps.(*OM2)
	if ok {
		return pst, err
	}
	return nil, err
}

// OM3 returns the first OM3 segment within the message, or nil if there isn't one.
func (m *Message) OM3() (*OM3, error) {
	ps, err := m.parse("OM3")
	pst, ok := ps.(*OM3)
	if ok {
		return pst, err
	}
	return nil, err
}

// OM4 returns the first OM4 segment within the message, or nil if there isn't one.
func (m *Message) OM4() (*OM4, error) {
	ps, err := m.parse("OM4")
	pst, ok := ps.(*OM4)
	if ok {
		return pst, err
	}
	return nil, err
}

// OM5 returns the first OM5 segment within the message, or nil if there isn't one.
func (m *Message) OM5() (*OM5, error) {
	ps, err := m.parse("OM5")
	pst, ok := ps.(*OM5)
	if ok {
		return pst, err
	}
	return nil, err
}

// OM6 returns the first OM6 segment within the message, or nil if there isn't one.
func (m *Message) OM6() (*OM6, error) {
	ps, err := m.parse("OM6")
	pst, ok := ps.(*OM6)
	if ok {
		return pst, err
	}
	return nil, err
}

// OM7 returns the first OM7 segment within the message, or nil if there isn't one.
func (m *Message) OM7() (*OM7, error) {
	ps, err := m.parse("OM7")
	pst, ok := ps.(*OM7)
	if ok {
		return pst, err
	}
	return nil, err
}

// ORC returns the first ORC segment within the message, or nil if there isn't one.
func (m *Message) ORC() (*ORC, error) {
	ps, err := m.parse("ORC")
	pst, ok := ps.(*ORC)
	if ok {
		return pst, err
	}
	return nil, err
}

// ORG returns the first ORG segment within the message, or nil if there isn't one.
func (m *Message) ORG() (*ORG, error) {
	ps, err := m.parse("ORG")
	pst, ok := ps.(*ORG)
	if ok {
		return pst, err
	}
	return nil, err
}

// ORO returns the first ORO segment within the message, or nil if there isn't one.
func (m *Message) ORO() (*ORO, error) {
	ps, err := m.parse("ORO")
	pst, ok := ps.(*ORO)
	if ok {
		return pst, err
	}
	return nil, err
}

// PCR returns the first PCR segment within the message, or nil if there isn't one.
func (m *Message) PCR() (*PCR, error) {
	ps, err := m.parse("PCR")
	pst, ok := ps.(*PCR)
	if ok {
		return pst, err
	}
	return nil, err
}

// PD1 returns the first PD1 segment within the message, or nil if there isn't one.
func (m *Message) PD1() (*PD1, error) {
	ps, err := m.parse("PD1")
	pst, ok := ps.(*PD1)
	if ok {
		return pst, err
	}
	return nil, err
}

// PDA returns the first PDA segment within the message, or nil if there isn't one.
func (m *Message) PDA() (*PDA, error) {
	ps, err := m.parse("PDA")
	pst, ok := ps.(*PDA)
	if ok {
		return pst, err
	}
	return nil, err
}

// PDC returns the first PDC segment within the message, or nil if there isn't one.
func (m *Message) PDC() (*PDC, error) {
	ps, err := m.parse("PDC")
	pst, ok := ps.(*PDC)
	if ok {
		return pst, err
	}
	return nil, err
}

// PEO returns the first PEO segment within the message, or nil if there isn't one.
func (m *Message) PEO() (*PEO, error) {
	ps, err := m.parse("PEO")
	pst, ok := ps.(*PEO)
	if ok {
		return pst, err
	}
	return nil, err
}

// PES returns the first PES segment within the message, or nil if there isn't one.
func (m *Message) PES() (*PES, error) {
	ps, err := m.parse("PES")
	pst, ok := ps.(*PES)
	if ok {
		return pst, err
	}
	return nil, err
}

// PID returns the first PID segment within the message, or nil if there isn't one.
func (m *Message) PID() (*PID, error) {
	ps, err := m.parse("PID")
	pst, ok := ps.(*PID)
	if ok {
		return pst, err
	}
	return nil, err
}

// PR1 returns the first PR1 segment within the message, or nil if there isn't one.
func (m *Message) PR1() (*PR1, error) {
	ps, err := m.parse("PR1")
	pst, ok := ps.(*PR1)
	if ok {
		return pst, err
	}
	return nil, err
}

// PRA returns the first PRA segment within the message, or nil if there isn't one.
func (m *Message) PRA() (*PRA, error) {
	ps, err := m.parse("PRA")
	pst, ok := ps.(*PRA)
	if ok {
		return pst, err
	}
	return nil, err
}

// PRB returns the first PRB segment within the message, or nil if there isn't one.
func (m *Message) PRB() (*PRB, error) {
	ps, err := m.parse("PRB")
	pst, ok := ps.(*PRB)
	if ok {
		return pst, err
	}
	return nil, err
}

// PRC returns the first PRC segment within the message, or nil if there isn't one.
func (m *Message) PRC() (*PRC, error) {
	ps, err := m.parse("PRC")
	pst, ok := ps.(*PRC)
	if ok {
		return pst, err
	}
	return nil, err
}

// PRD returns the first PRD segment within the message, or nil if there isn't one.
func (m *Message) PRD() (*PRD, error) {
	ps, err := m.parse("PRD")
	pst, ok := ps.(*PRD)
	if ok {
		return pst, err
	}
	return nil, err
}

// PSH returns the first PSH segment within the message, or nil if there isn't one.
func (m *Message) PSH() (*PSH, error) {
	ps, err := m.parse("PSH")
	pst, ok := ps.(*PSH)
	if ok {
		return pst, err
	}
	return nil, err
}

// PTH returns the first PTH segment within the message, or nil if there isn't one.
func (m *Message) PTH() (*PTH, error) {
	ps, err := m.parse("PTH")
	pst, ok := ps.(*PTH)
	if ok {
		return pst, err
	}
	return nil, err
}

// PV1 returns the first PV1 segment within the message, or nil if there isn't one.
func (m *Message) PV1() (*PV1, error) {
	ps, err := m.parse("PV1")
	pst, ok := ps.(*PV1)
	if ok {
		return pst, err
	}
	return nil, err
}

// PV2 returns the first PV2 segment within the message, or nil if there isn't one.
func (m *Message) PV2() (*PV2, error) {
	ps, err := m.parse("PV2")
	pst, ok := ps.(*PV2)
	if ok {
		return pst, err
	}
	return nil, err
}

// QAK returns the first QAK segment within the message, or nil if there isn't one.
func (m *Message) QAK() (*QAK, error) {
	ps, err := m.parse("QAK")
	pst, ok := ps.(*QAK)
	if ok {
		return pst, err
	}
	return nil, err
}

// QID returns the first QID segment within the message, or nil if there isn't one.
func (m *Message) QID() (*QID, error) {
	ps, err := m.parse("QID")
	pst, ok := ps.(*QID)
	if ok {
		return pst, err
	}
	return nil, err
}

// QPD returns the first QPD segment within the message, or nil if there isn't one.
func (m *Message) QPD() (*QPD, error) {
	ps, err := m.parse("QPD")
	pst, ok := ps.(*QPD)
	if ok {
		return pst, err
	}
	return nil, err
}

// QRD returns the first QRD segment within the message, or nil if there isn't one.
func (m *Message) QRD() (*QRD, error) {
	ps, err := m.parse("QRD")
	pst, ok := ps.(*QRD)
	if ok {
		return pst, err
	}
	return nil, err
}

// QRF returns the first QRF segment within the message, or nil if there isn't one.
func (m *Message) QRF() (*QRF, error) {
	ps, err := m.parse("QRF")
	pst, ok := ps.(*QRF)
	if ok {
		return pst, err
	}
	return nil, err
}

// QRI returns the first QRI segment within the message, or nil if there isn't one.
func (m *Message) QRI() (*QRI, error) {
	ps, err := m.parse("QRI")
	pst, ok := ps.(*QRI)
	if ok {
		return pst, err
	}
	return nil, err
}

// RCP returns the first RCP segment within the message, or nil if there isn't one.
func (m *Message) RCP() (*RCP, error) {
	ps, err := m.parse("RCP")
	pst, ok := ps.(*RCP)
	if ok {
		return pst, err
	}
	return nil, err
}

// RDF returns the first RDF segment within the message, or nil if there isn't one.
func (m *Message) RDF() (*RDF, error) {
	ps, err := m.parse("RDF")
	pst, ok := ps.(*RDF)
	if ok {
		return pst, err
	}
	return nil, err
}

// RDT returns the first RDT segment within the message, or nil if there isn't one.
func (m *Message) RDT() (*RDT, error) {
	ps, err := m.parse("RDT")
	pst, ok := ps.(*RDT)
	if ok {
		return pst, err
	}
	return nil, err
}

// RF1 returns the first RF1 segment within the message, or nil if there isn't one.
func (m *Message) RF1() (*RF1, error) {
	ps, err := m.parse("RF1")
	pst, ok := ps.(*RF1)
	if ok {
		return pst, err
	}
	return nil, err
}

// RGS returns the first RGS segment within the message, or nil if there isn't one.
func (m *Message) RGS() (*RGS, error) {
	ps, err := m.parse("RGS")
	pst, ok := ps.(*RGS)
	if ok {
		return pst, err
	}
	return nil, err
}

// RMI returns the first RMI segment within the message, or nil if there isn't one.
func (m *Message) RMI() (*RMI, error) {
	ps, err := m.parse("RMI")
	pst, ok := ps.(*RMI)
	if ok {
		return pst, err
	}
	return nil, err
}

// ROL returns the first ROL segment within the message, or nil if there isn't one.
func (m *Message) ROL() (*ROL, error) {
	ps, err := m.parse("ROL")
	pst, ok := ps.(*ROL)
	if ok {
		return pst, err
	}
	return nil, err
}

// RQ1 returns the first RQ1 segment within the message, or nil if there isn't one.
func (m *Message) RQ1() (*RQ1, error) {
	ps, err := m.parse("RQ1")
	pst, ok := ps.(*RQ1)
	if ok {
		return pst, err
	}
	return nil, err
}

// RQD returns the first RQD segment within the message, or nil if there isn't one.
func (m *Message) RQD() (*RQD, error) {
	ps, err := m.parse("RQD")
	pst, ok := ps.(*RQD)
	if ok {
		return pst, err
	}
	return nil, err
}

// RX1 returns the first RX1 segment within the message, or nil if there isn't one.
func (m *Message) RX1() (*RX1, error) {
	ps, err := m.parse("RX1")
	pst, ok := ps.(*RX1)
	if ok {
		return pst, err
	}
	return nil, err
}

// RXA returns the first RXA segment within the message, or nil if there isn't one.
func (m *Message) RXA() (*RXA, error) {
	ps, err := m.parse("RXA")
	pst, ok := ps.(*RXA)
	if ok {
		return pst, err
	}
	return nil, err
}

// RXC returns the first RXC segment within the message, or nil if there isn't one.
func (m *Message) RXC() (*RXC, error) {
	ps, err := m.parse("RXC")
	pst, ok := ps.(*RXC)
	if ok {
		return pst, err
	}
	return nil, err
}

// RXD returns the first RXD segment within the message, or nil if there isn't one.
func (m *Message) RXD() (*RXD, error) {
	ps, err := m.parse("RXD")
	pst, ok := ps.(*RXD)
	if ok {
		return pst, err
	}
	return nil, err
}

// RXE returns the first RXE segment within the message, or nil if there isn't one.
func (m *Message) RXE() (*RXE, error) {
	ps, err := m.parse("RXE")
	pst, ok := ps.(*RXE)
	if ok {
		return pst, err
	}
	return nil, err
}

// RXG returns the first RXG segment within the message, or nil if there isn't one.
func (m *Message) RXG() (*RXG, error) {
	ps, err := m.parse("RXG")
	pst, ok := ps.(*RXG)
	if ok {
		return pst, err
	}
	return nil, err
}

// RXO returns the first RXO segment within the message, or nil if there isn't one.
func (m *Message) RXO() (*RXO, error) {
	ps, err := m.parse("RXO")
	pst, ok := ps.(*RXO)
	if ok {
		return pst, err
	}
	return nil, err
}

// RXR returns the first RXR segment within the message, or nil if there isn't one.
func (m *Message) RXR() (*RXR, error) {
	ps, err := m.parse("RXR")
	pst, ok := ps.(*RXR)
	if ok {
		return pst, err
	}
	return nil, err
}

// SAC returns the first SAC segment within the message, or nil if there isn't one.
func (m *Message) SAC() (*SAC, error) {
	ps, err := m.parse("SAC")
	pst, ok := ps.(*SAC)
	if ok {
		return pst, err
	}
	return nil, err
}

// SCH returns the first SCH segment within the message, or nil if there isn't one.
func (m *Message) SCH() (*SCH, error) {
	ps, err := m.parse("SCH")
	pst, ok := ps.(*SCH)
	if ok {
		return pst, err
	}
	return nil, err
}

// SID returns the first SID segment within the message, or nil if there isn't one.
func (m *Message) SID() (*SID, error) {
	ps, err := m.parse("SID")
	pst, ok := ps.(*SID)
	if ok {
		return pst, err
	}
	return nil, err
}

// SPR returns the first SPR segment within the message, or nil if there isn't one.
func (m *Message) SPR() (*SPR, error) {
	ps, err := m.parse("SPR")
	pst, ok := ps.(*SPR)
	if ok {
		return pst, err
	}
	return nil, err
}

// STF returns the first STF segment within the message, or nil if there isn't one.
func (m *Message) STF() (*STF, error) {
	ps, err := m.parse("STF")
	pst, ok := ps.(*STF)
	if ok {
		return pst, err
	}
	return nil, err
}

// TCC returns the first TCC segment within the message, or nil if there isn't one.
func (m *Message) TCC() (*TCC, error) {
	ps, err := m.parse("TCC")
	pst, ok := ps.(*TCC)
	if ok {
		return pst, err
	}
	return nil, err
}

// TCD returns the first TCD segment within the message, or nil if there isn't one.
func (m *Message) TCD() (*TCD, error) {
	ps, err := m.parse("TCD")
	pst, ok := ps.(*TCD)
	if ok {
		return pst, err
	}
	return nil, err
}

// TXA returns the first TXA segment within the message, or nil if there isn't one.
func (m *Message) TXA() (*TXA, error) {
	ps, err := m.parse("TXA")
	pst, ok := ps.(*TXA)
	if ok {
		return pst, err
	}
	return nil, err
}

// UB1 returns the first UB1 segment within the message, or nil if there isn't one.
func (m *Message) UB1() (*UB1, error) {
	ps, err := m.parse("UB1")
	pst, ok := ps.(*UB1)
	if ok {
		return pst, err
	}
	return nil, err
}

// UB2 returns the first UB2 segment within the message, or nil if there isn't one.
func (m *Message) UB2() (*UB2, error) {
	ps, err := m.parse("UB2")
	pst, ok := ps.(*UB2)
	if ok {
		return pst, err
	}
	return nil, err
}

// URD returns the first URD segment within the message, or nil if there isn't one.
func (m *Message) URD() (*URD, error) {
	ps, err := m.parse("URD")
	pst, ok := ps.(*URD)
	if ok {
		return pst, err
	}
	return nil, err
}

// URS returns the first URS segment within the message, or nil if there isn't one.
func (m *Message) URS() (*URS, error) {
	ps, err := m.parse("URS")
	pst, ok := ps.(*URS)
	if ok {
		return pst, err
	}
	return nil, err
}

// VAR returns the first VAR segment within the message, or nil if there isn't one.
func (m *Message) VAR() (*VAR, error) {
	ps, err := m.parse("VAR")
	pst, ok := ps.(*VAR)
	if ok {
		return pst, err
	}
	return nil, err
}

// VTQ returns the first VTQ segment within the message, or nil if there isn't one.
func (m *Message) VTQ() (*VTQ, error) {
	ps, err := m.parse("VTQ")
	pst, ok := ps.(*VTQ)
	if ok {
		return pst, err
	}
	return nil, err
}

// AllABS returns a slice containing all ABS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllABS() ([]*ABS, error) {
	pss, err := m.ParseAll("ABS")
	return pss.([]*ABS), err
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

// AllAFF returns a slice containing all AFF segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllAFF() ([]*AFF, error) {
	pss, err := m.ParseAll("AFF")
	return pss.([]*AFF), err
}

// AllAIG returns a slice containing all AIG segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllAIG() ([]*AIG, error) {
	pss, err := m.ParseAll("AIG")
	return pss.([]*AIG), err
}

// AllAIL returns a slice containing all AIL segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllAIL() ([]*AIL, error) {
	pss, err := m.ParseAll("AIL")
	return pss.([]*AIL), err
}

// AllAIP returns a slice containing all AIP segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllAIP() ([]*AIP, error) {
	pss, err := m.ParseAll("AIP")
	return pss.([]*AIP), err
}

// AllAIS returns a slice containing all AIS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllAIS() ([]*AIS, error) {
	pss, err := m.ParseAll("AIS")
	return pss.([]*AIS), err
}

// AllAL1 returns a slice containing all AL1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllAL1() ([]*AL1, error) {
	pss, err := m.ParseAll("AL1")
	return pss.([]*AL1), err
}

// AllAPR returns a slice containing all APR segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllAPR() ([]*APR, error) {
	pss, err := m.ParseAll("APR")
	return pss.([]*APR), err
}

// AllARQ returns a slice containing all ARQ segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllARQ() ([]*ARQ, error) {
	pss, err := m.ParseAll("ARQ")
	return pss.([]*ARQ), err
}

// AllAUT returns a slice containing all AUT segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllAUT() ([]*AUT, error) {
	pss, err := m.ParseAll("AUT")
	return pss.([]*AUT), err
}

// AllBHS returns a slice containing all BHS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllBHS() ([]*BHS, error) {
	pss, err := m.ParseAll("BHS")
	return pss.([]*BHS), err
}

// AllBLC returns a slice containing all BLC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllBLC() ([]*BLC, error) {
	pss, err := m.ParseAll("BLC")
	return pss.([]*BLC), err
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

// AllCDM returns a slice containing all CDM segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCDM() ([]*CDM, error) {
	pss, err := m.ParseAll("CDM")
	return pss.([]*CDM), err
}

// AllCM0 returns a slice containing all CM0 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCM0() ([]*CM0, error) {
	pss, err := m.ParseAll("CM0")
	return pss.([]*CM0), err
}

// AllCM1 returns a slice containing all CM1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCM1() ([]*CM1, error) {
	pss, err := m.ParseAll("CM1")
	return pss.([]*CM1), err
}

// AllCM2 returns a slice containing all CM2 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCM2() ([]*CM2, error) {
	pss, err := m.ParseAll("CM2")
	return pss.([]*CM2), err
}

// AllCNS returns a slice containing all CNS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCNS() ([]*CNS, error) {
	pss, err := m.ParseAll("CNS")
	return pss.([]*CNS), err
}

// AllCSP returns a slice containing all CSP segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCSP() ([]*CSP, error) {
	pss, err := m.ParseAll("CSP")
	return pss.([]*CSP), err
}

// AllCSR returns a slice containing all CSR segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCSR() ([]*CSR, error) {
	pss, err := m.ParseAll("CSR")
	return pss.([]*CSR), err
}

// AllCSS returns a slice containing all CSS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCSS() ([]*CSS, error) {
	pss, err := m.ParseAll("CSS")
	return pss.([]*CSS), err
}

// AllCTD returns a slice containing all CTD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCTD() ([]*CTD, error) {
	pss, err := m.ParseAll("CTD")
	return pss.([]*CTD), err
}

// AllCTI returns a slice containing all CTI segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllCTI() ([]*CTI, error) {
	pss, err := m.ParseAll("CTI")
	return pss.([]*CTI), err
}

// AllDB1 returns a slice containing all DB1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllDB1() ([]*DB1, error) {
	pss, err := m.ParseAll("DB1")
	return pss.([]*DB1), err
}

// AllDG1 returns a slice containing all DG1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllDG1() ([]*DG1, error) {
	pss, err := m.ParseAll("DG1")
	return pss.([]*DG1), err
}

// AllDRG returns a slice containing all DRG segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllDRG() ([]*DRG, error) {
	pss, err := m.ParseAll("DRG")
	return pss.([]*DRG), err
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

// AllECD returns a slice containing all ECD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllECD() ([]*ECD, error) {
	pss, err := m.ParseAll("ECD")
	return pss.([]*ECD), err
}

// AllECR returns a slice containing all ECR segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllECR() ([]*ECR, error) {
	pss, err := m.ParseAll("ECR")
	return pss.([]*ECR), err
}

// AllEDU returns a slice containing all EDU segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllEDU() ([]*EDU, error) {
	pss, err := m.ParseAll("EDU")
	return pss.([]*EDU), err
}

// AllEQL returns a slice containing all EQL segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllEQL() ([]*EQL, error) {
	pss, err := m.ParseAll("EQL")
	return pss.([]*EQL), err
}

// AllEQP returns a slice containing all EQP segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllEQP() ([]*EQP, error) {
	pss, err := m.ParseAll("EQP")
	return pss.([]*EQP), err
}

// AllEQU returns a slice containing all EQU segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllEQU() ([]*EQU, error) {
	pss, err := m.ParseAll("EQU")
	return pss.([]*EQU), err
}

// AllERQ returns a slice containing all ERQ segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllERQ() ([]*ERQ, error) {
	pss, err := m.ParseAll("ERQ")
	return pss.([]*ERQ), err
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

// AllFAC returns a slice containing all FAC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllFAC() ([]*FAC, error) {
	pss, err := m.ParseAll("FAC")
	return pss.([]*FAC), err
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

// AllGOL returns a slice containing all GOL segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllGOL() ([]*GOL, error) {
	pss, err := m.ParseAll("GOL")
	return pss.([]*GOL), err
}

// AllGP1 returns a slice containing all GP1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllGP1() ([]*GP1, error) {
	pss, err := m.ParseAll("GP1")
	return pss.([]*GP1), err
}

// AllGP2 returns a slice containing all GP2 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllGP2() ([]*GP2, error) {
	pss, err := m.ParseAll("GP2")
	return pss.([]*GP2), err
}

// AllGT1 returns a slice containing all GT1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllGT1() ([]*GT1, error) {
	pss, err := m.ParseAll("GT1")
	return pss.([]*GT1), err
}

// AllIAM returns a slice containing all IAM segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllIAM() ([]*IAM, error) {
	pss, err := m.ParseAll("IAM")
	return pss.([]*IAM), err
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

// AllINV returns a slice containing all INV segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllINV() ([]*INV, error) {
	pss, err := m.ParseAll("INV")
	return pss.([]*INV), err
}

// AllISD returns a slice containing all ISD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllISD() ([]*ISD, error) {
	pss, err := m.ParseAll("ISD")
	return pss.([]*ISD), err
}

// AllLAN returns a slice containing all LAN segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllLAN() ([]*LAN, error) {
	pss, err := m.ParseAll("LAN")
	return pss.([]*LAN), err
}

// AllLCC returns a slice containing all LCC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllLCC() ([]*LCC, error) {
	pss, err := m.ParseAll("LCC")
	return pss.([]*LCC), err
}

// AllLCH returns a slice containing all LCH segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllLCH() ([]*LCH, error) {
	pss, err := m.ParseAll("LCH")
	return pss.([]*LCH), err
}

// AllLDP returns a slice containing all LDP segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllLDP() ([]*LDP, error) {
	pss, err := m.ParseAll("LDP")
	return pss.([]*LDP), err
}

// AllLOC returns a slice containing all LOC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllLOC() ([]*LOC, error) {
	pss, err := m.ParseAll("LOC")
	return pss.([]*LOC), err
}

// AllLRL returns a slice containing all LRL segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllLRL() ([]*LRL, error) {
	pss, err := m.ParseAll("LRL")
	return pss.([]*LRL), err
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

// AllNDS returns a slice containing all NDS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllNDS() ([]*NDS, error) {
	pss, err := m.ParseAll("NDS")
	return pss.([]*NDS), err
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

// AllOM7 returns a slice containing all OM7 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllOM7() ([]*OM7, error) {
	pss, err := m.ParseAll("OM7")
	return pss.([]*OM7), err
}

// AllORC returns a slice containing all ORC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllORC() ([]*ORC, error) {
	pss, err := m.ParseAll("ORC")
	return pss.([]*ORC), err
}

// AllORG returns a slice containing all ORG segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllORG() ([]*ORG, error) {
	pss, err := m.ParseAll("ORG")
	return pss.([]*ORG), err
}

// AllORO returns a slice containing all ORO segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllORO() ([]*ORO, error) {
	pss, err := m.ParseAll("ORO")
	return pss.([]*ORO), err
}

// AllPCR returns a slice containing all PCR segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPCR() ([]*PCR, error) {
	pss, err := m.ParseAll("PCR")
	return pss.([]*PCR), err
}

// AllPD1 returns a slice containing all PD1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPD1() ([]*PD1, error) {
	pss, err := m.ParseAll("PD1")
	return pss.([]*PD1), err
}

// AllPDA returns a slice containing all PDA segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPDA() ([]*PDA, error) {
	pss, err := m.ParseAll("PDA")
	return pss.([]*PDA), err
}

// AllPDC returns a slice containing all PDC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPDC() ([]*PDC, error) {
	pss, err := m.ParseAll("PDC")
	return pss.([]*PDC), err
}

// AllPEO returns a slice containing all PEO segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPEO() ([]*PEO, error) {
	pss, err := m.ParseAll("PEO")
	return pss.([]*PEO), err
}

// AllPES returns a slice containing all PES segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPES() ([]*PES, error) {
	pss, err := m.ParseAll("PES")
	return pss.([]*PES), err
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

// AllPRB returns a slice containing all PRB segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPRB() ([]*PRB, error) {
	pss, err := m.ParseAll("PRB")
	return pss.([]*PRB), err
}

// AllPRC returns a slice containing all PRC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPRC() ([]*PRC, error) {
	pss, err := m.ParseAll("PRC")
	return pss.([]*PRC), err
}

// AllPRD returns a slice containing all PRD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPRD() ([]*PRD, error) {
	pss, err := m.ParseAll("PRD")
	return pss.([]*PRD), err
}

// AllPSH returns a slice containing all PSH segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPSH() ([]*PSH, error) {
	pss, err := m.ParseAll("PSH")
	return pss.([]*PSH), err
}

// AllPTH returns a slice containing all PTH segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPTH() ([]*PTH, error) {
	pss, err := m.ParseAll("PTH")
	return pss.([]*PTH), err
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

// AllQAK returns a slice containing all QAK segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllQAK() ([]*QAK, error) {
	pss, err := m.ParseAll("QAK")
	return pss.([]*QAK), err
}

// AllQID returns a slice containing all QID segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllQID() ([]*QID, error) {
	pss, err := m.ParseAll("QID")
	return pss.([]*QID), err
}

// AllQPD returns a slice containing all QPD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllQPD() ([]*QPD, error) {
	pss, err := m.ParseAll("QPD")
	return pss.([]*QPD), err
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

// AllQRI returns a slice containing all QRI segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllQRI() ([]*QRI, error) {
	pss, err := m.ParseAll("QRI")
	return pss.([]*QRI), err
}

// AllRCP returns a slice containing all RCP segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRCP() ([]*RCP, error) {
	pss, err := m.ParseAll("RCP")
	return pss.([]*RCP), err
}

// AllRDF returns a slice containing all RDF segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRDF() ([]*RDF, error) {
	pss, err := m.ParseAll("RDF")
	return pss.([]*RDF), err
}

// AllRDT returns a slice containing all RDT segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRDT() ([]*RDT, error) {
	pss, err := m.ParseAll("RDT")
	return pss.([]*RDT), err
}

// AllRF1 returns a slice containing all RF1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRF1() ([]*RF1, error) {
	pss, err := m.ParseAll("RF1")
	return pss.([]*RF1), err
}

// AllRGS returns a slice containing all RGS segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRGS() ([]*RGS, error) {
	pss, err := m.ParseAll("RGS")
	return pss.([]*RGS), err
}

// AllRMI returns a slice containing all RMI segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRMI() ([]*RMI, error) {
	pss, err := m.ParseAll("RMI")
	return pss.([]*RMI), err
}

// AllROL returns a slice containing all ROL segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllROL() ([]*ROL, error) {
	pss, err := m.ParseAll("ROL")
	return pss.([]*ROL), err
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

// AllSAC returns a slice containing all SAC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllSAC() ([]*SAC, error) {
	pss, err := m.ParseAll("SAC")
	return pss.([]*SAC), err
}

// AllSCH returns a slice containing all SCH segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllSCH() ([]*SCH, error) {
	pss, err := m.ParseAll("SCH")
	return pss.([]*SCH), err
}

// AllSID returns a slice containing all SID segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllSID() ([]*SID, error) {
	pss, err := m.ParseAll("SID")
	return pss.([]*SID), err
}

// AllSPR returns a slice containing all SPR segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllSPR() ([]*SPR, error) {
	pss, err := m.ParseAll("SPR")
	return pss.([]*SPR), err
}

// AllSTF returns a slice containing all STF segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllSTF() ([]*STF, error) {
	pss, err := m.ParseAll("STF")
	return pss.([]*STF), err
}

// AllTCC returns a slice containing all TCC segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllTCC() ([]*TCC, error) {
	pss, err := m.ParseAll("TCC")
	return pss.([]*TCC), err
}

// AllTCD returns a slice containing all TCD segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllTCD() ([]*TCD, error) {
	pss, err := m.ParseAll("TCD")
	return pss.([]*TCD), err
}

// AllTXA returns a slice containing all TXA segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllTXA() ([]*TXA, error) {
	pss, err := m.ParseAll("TXA")
	return pss.([]*TXA), err
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

// AllVAR returns a slice containing all VAR segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllVAR() ([]*VAR, error) {
	pss, err := m.ParseAll("VAR")
	return pss.([]*VAR), err
}

// AllVTQ returns a slice containing all VTQ segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllVTQ() ([]*VTQ, error) {
	pss, err := m.ParseAll("VTQ")
	return pss.([]*VTQ), err
}

// Types maps the name of an HL7 segment or message type to the type of the struct that
// represents that segment or message type.
var Types = map[string]reflect.Type{}

func init() {
	Types["ABS"] = reflect.TypeOf(ABS{})
	Types["ACC"] = reflect.TypeOf(ACC{})
	Types["ADD"] = reflect.TypeOf(ADD{})
	Types["AFF"] = reflect.TypeOf(AFF{})
	Types["AIG"] = reflect.TypeOf(AIG{})
	Types["AIL"] = reflect.TypeOf(AIL{})
	Types["AIP"] = reflect.TypeOf(AIP{})
	Types["AIS"] = reflect.TypeOf(AIS{})
	Types["AL1"] = reflect.TypeOf(AL1{})
	Types["APR"] = reflect.TypeOf(APR{})
	Types["ARQ"] = reflect.TypeOf(ARQ{})
	Types["AUT"] = reflect.TypeOf(AUT{})
	Types["BHS"] = reflect.TypeOf(BHS{})
	Types["BLC"] = reflect.TypeOf(BLC{})
	Types["BLG"] = reflect.TypeOf(BLG{})
	Types["BTS"] = reflect.TypeOf(BTS{})
	Types["CDM"] = reflect.TypeOf(CDM{})
	Types["CM0"] = reflect.TypeOf(CM0{})
	Types["CM1"] = reflect.TypeOf(CM1{})
	Types["CM2"] = reflect.TypeOf(CM2{})
	Types["CNS"] = reflect.TypeOf(CNS{})
	Types["CSP"] = reflect.TypeOf(CSP{})
	Types["CSR"] = reflect.TypeOf(CSR{})
	Types["CSS"] = reflect.TypeOf(CSS{})
	Types["CTD"] = reflect.TypeOf(CTD{})
	Types["CTI"] = reflect.TypeOf(CTI{})
	Types["DB1"] = reflect.TypeOf(DB1{})
	Types["DG1"] = reflect.TypeOf(DG1{})
	Types["DRG"] = reflect.TypeOf(DRG{})
	Types["DSC"] = reflect.TypeOf(DSC{})
	Types["DSP"] = reflect.TypeOf(DSP{})
	Types["ECD"] = reflect.TypeOf(ECD{})
	Types["ECR"] = reflect.TypeOf(ECR{})
	Types["EDU"] = reflect.TypeOf(EDU{})
	Types["EQL"] = reflect.TypeOf(EQL{})
	Types["EQP"] = reflect.TypeOf(EQP{})
	Types["EQU"] = reflect.TypeOf(EQU{})
	Types["ERQ"] = reflect.TypeOf(ERQ{})
	Types["ERR"] = reflect.TypeOf(ERR{})
	Types["EVN"] = reflect.TypeOf(EVN{})
	Types["FAC"] = reflect.TypeOf(FAC{})
	Types["FHS"] = reflect.TypeOf(FHS{})
	Types["FT1"] = reflect.TypeOf(FT1{})
	Types["FTS"] = reflect.TypeOf(FTS{})
	Types["GOL"] = reflect.TypeOf(GOL{})
	Types["GP1"] = reflect.TypeOf(GP1{})
	Types["GP2"] = reflect.TypeOf(GP2{})
	Types["GT1"] = reflect.TypeOf(GT1{})
	Types["IAM"] = reflect.TypeOf(IAM{})
	Types["IN1"] = reflect.TypeOf(IN1{})
	Types["IN2"] = reflect.TypeOf(IN2{})
	Types["IN3"] = reflect.TypeOf(IN3{})
	Types["INV"] = reflect.TypeOf(INV{})
	Types["ISD"] = reflect.TypeOf(ISD{})
	Types["LAN"] = reflect.TypeOf(LAN{})
	Types["LCC"] = reflect.TypeOf(LCC{})
	Types["LCH"] = reflect.TypeOf(LCH{})
	Types["LDP"] = reflect.TypeOf(LDP{})
	Types["LOC"] = reflect.TypeOf(LOC{})
	Types["LRL"] = reflect.TypeOf(LRL{})
	Types["MFA"] = reflect.TypeOf(MFA{})
	Types["MFE"] = reflect.TypeOf(MFE{})
	Types["MFI"] = reflect.TypeOf(MFI{})
	Types["MRG"] = reflect.TypeOf(MRG{})
	Types["MSA"] = reflect.TypeOf(MSA{})
	Types["MSH"] = reflect.TypeOf(MSH{})
	Types["NCK"] = reflect.TypeOf(NCK{})
	Types["NDS"] = reflect.TypeOf(NDS{})
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
	Types["OM7"] = reflect.TypeOf(OM7{})
	Types["ORC"] = reflect.TypeOf(ORC{})
	Types["ORG"] = reflect.TypeOf(ORG{})
	Types["ORO"] = reflect.TypeOf(ORO{})
	Types["PCR"] = reflect.TypeOf(PCR{})
	Types["PD1"] = reflect.TypeOf(PD1{})
	Types["PDA"] = reflect.TypeOf(PDA{})
	Types["PDC"] = reflect.TypeOf(PDC{})
	Types["PEO"] = reflect.TypeOf(PEO{})
	Types["PES"] = reflect.TypeOf(PES{})
	Types["PID"] = reflect.TypeOf(PID{})
	Types["PR1"] = reflect.TypeOf(PR1{})
	Types["PRA"] = reflect.TypeOf(PRA{})
	Types["PRB"] = reflect.TypeOf(PRB{})
	Types["PRC"] = reflect.TypeOf(PRC{})
	Types["PRD"] = reflect.TypeOf(PRD{})
	Types["PSH"] = reflect.TypeOf(PSH{})
	Types["PTH"] = reflect.TypeOf(PTH{})
	Types["PV1"] = reflect.TypeOf(PV1{})
	Types["PV2"] = reflect.TypeOf(PV2{})
	Types["QAK"] = reflect.TypeOf(QAK{})
	Types["QID"] = reflect.TypeOf(QID{})
	Types["QPD"] = reflect.TypeOf(QPD{})
	Types["QRD"] = reflect.TypeOf(QRD{})
	Types["QRF"] = reflect.TypeOf(QRF{})
	Types["QRI"] = reflect.TypeOf(QRI{})
	Types["RCP"] = reflect.TypeOf(RCP{})
	Types["RDF"] = reflect.TypeOf(RDF{})
	Types["RDT"] = reflect.TypeOf(RDT{})
	Types["RF1"] = reflect.TypeOf(RF1{})
	Types["RGS"] = reflect.TypeOf(RGS{})
	Types["RMI"] = reflect.TypeOf(RMI{})
	Types["ROL"] = reflect.TypeOf(ROL{})
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
	Types["SAC"] = reflect.TypeOf(SAC{})
	Types["SCH"] = reflect.TypeOf(SCH{})
	Types["SID"] = reflect.TypeOf(SID{})
	Types["SPR"] = reflect.TypeOf(SPR{})
	Types["STF"] = reflect.TypeOf(STF{})
	Types["TCC"] = reflect.TypeOf(TCC{})
	Types["TCD"] = reflect.TypeOf(TCD{})
	Types["TXA"] = reflect.TypeOf(TXA{})
	Types["UB1"] = reflect.TypeOf(UB1{})
	Types["UB2"] = reflect.TypeOf(UB2{})
	Types["URD"] = reflect.TypeOf(URD{})
	Types["URS"] = reflect.TypeOf(URS{})
	Types["VAR"] = reflect.TypeOf(VAR{})
	Types["VTQ"] = reflect.TypeOf(VTQ{})
	Types["ACK"] = reflect.TypeOf(ACK{})
	Types["ACK_N02"] = reflect.TypeOf(ACK_N02{})
	Types["ADR_A19"] = reflect.TypeOf(ADR_A19{})
	Types["ADR_A19_INSURANCE"] = reflect.TypeOf(ADR_A19_INSURANCE{})
	Types["ADR_A19_PROCEDURE"] = reflect.TypeOf(ADR_A19_PROCEDURE{})
	Types["ADR_A19_QUERY_RESPONSE"] = reflect.TypeOf(ADR_A19_QUERY_RESPONSE{})
	Types["ADT_A01"] = reflect.TypeOf(ADT_A01{})
	Types["ADT_A01_INSURANCE"] = reflect.TypeOf(ADT_A01_INSURANCE{})
	Types["ADT_A01_PROCEDURE"] = reflect.TypeOf(ADT_A01_PROCEDURE{})
	Types["ADT_A02"] = reflect.TypeOf(ADT_A02{})
	Types["ADT_A03"] = reflect.TypeOf(ADT_A03{})
	Types["ADT_A03_PROCEDURE"] = reflect.TypeOf(ADT_A03_PROCEDURE{})
	Types["ADT_A04"] = reflect.TypeOf(ADT_A04{})
	Types["ADT_A04_INSURANCE"] = reflect.TypeOf(ADT_A04_INSURANCE{})
	Types["ADT_A05"] = reflect.TypeOf(ADT_A05{})
	Types["ADT_A05_INSURANCE"] = reflect.TypeOf(ADT_A05_INSURANCE{})
	Types["ADT_A05_PROCEDURE"] = reflect.TypeOf(ADT_A05_PROCEDURE{})
	Types["ADT_A06"] = reflect.TypeOf(ADT_A06{})
	Types["ADT_A06_INSURANCE"] = reflect.TypeOf(ADT_A06_INSURANCE{})
	Types["ADT_A06_PROCEDURE"] = reflect.TypeOf(ADT_A06_PROCEDURE{})
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
	Types["ADT_A38"] = reflect.TypeOf(ADT_A38{})
	Types["ADT_A39"] = reflect.TypeOf(ADT_A39{})
	Types["ADT_A39_PATIENT"] = reflect.TypeOf(ADT_A39_PATIENT{})
	Types["ADT_A43"] = reflect.TypeOf(ADT_A43{})
	Types["ADT_A43_PATIENT"] = reflect.TypeOf(ADT_A43_PATIENT{})
	Types["ADT_A44"] = reflect.TypeOf(ADT_A44{})
	Types["ADT_A44_PATIENT"] = reflect.TypeOf(ADT_A44_PATIENT{})
	Types["ADT_A45"] = reflect.TypeOf(ADT_A45{})
	Types["ADT_A45_MERGE_INFO"] = reflect.TypeOf(ADT_A45_MERGE_INFO{})
	Types["ADT_A50"] = reflect.TypeOf(ADT_A50{})
	Types["ADT_A52"] = reflect.TypeOf(ADT_A52{})
	Types["ADT_A54"] = reflect.TypeOf(ADT_A54{})
	Types["ADT_A60"] = reflect.TypeOf(ADT_A60{})
	Types["ADT_A61"] = reflect.TypeOf(ADT_A61{})
	Types["ARD_A19"] = reflect.TypeOf(ARD_A19{})
	Types["ARD_A19_INSURANCE"] = reflect.TypeOf(ARD_A19_INSURANCE{})
	Types["ARD_A19_PROCEDURE"] = reflect.TypeOf(ARD_A19_PROCEDURE{})
	Types["ARD_A19_QUERY_RESPONSE"] = reflect.TypeOf(ARD_A19_QUERY_RESPONSE{})
	Types["BAR_P01"] = reflect.TypeOf(BAR_P01{})
	Types["BAR_P01_INSURANCE"] = reflect.TypeOf(BAR_P01_INSURANCE{})
	Types["BAR_P01_PROCEDURE"] = reflect.TypeOf(BAR_P01_PROCEDURE{})
	Types["BAR_P01_VISIT"] = reflect.TypeOf(BAR_P01_VISIT{})
	Types["BAR_P02"] = reflect.TypeOf(BAR_P02{})
	Types["BAR_P02_PATIENT"] = reflect.TypeOf(BAR_P02_PATIENT{})
	Types["BAR_P05"] = reflect.TypeOf(BAR_P05{})
	Types["BAR_P05_INSURANCE"] = reflect.TypeOf(BAR_P05_INSURANCE{})
	Types["BAR_P05_PROCEDURE"] = reflect.TypeOf(BAR_P05_PROCEDURE{})
	Types["BAR_P05_VISIT"] = reflect.TypeOf(BAR_P05_VISIT{})
	Types["BAR_P06"] = reflect.TypeOf(BAR_P06{})
	Types["BAR_P06_PATIENT"] = reflect.TypeOf(BAR_P06_PATIENT{})
	Types["BAR_P10"] = reflect.TypeOf(BAR_P10{})
	Types["BAR_P10_PROCEDURE"] = reflect.TypeOf(BAR_P10_PROCEDURE{})
	Types["CRM_C01"] = reflect.TypeOf(CRM_C01{})
	Types["CRM_C01_PATIENT"] = reflect.TypeOf(CRM_C01_PATIENT{})
	Types["CSU_C09"] = reflect.TypeOf(CSU_C09{})
	Types["CSU_C09_PATIENT"] = reflect.TypeOf(CSU_C09_PATIENT{})
	Types["CSU_C09_RX_ADMIN"] = reflect.TypeOf(CSU_C09_RX_ADMIN{})
	Types["CSU_C09_STUDY_OBSERVATION"] = reflect.TypeOf(CSU_C09_STUDY_OBSERVATION{})
	Types["CSU_C09_STUDY_PHARM"] = reflect.TypeOf(CSU_C09_STUDY_PHARM{})
	Types["CSU_C09_STUDY_PHASE"] = reflect.TypeOf(CSU_C09_STUDY_PHASE{})
	Types["CSU_C09_STUDY_SCHEDULE"] = reflect.TypeOf(CSU_C09_STUDY_SCHEDULE{})
	Types["CSU_C09_VISIT"] = reflect.TypeOf(CSU_C09_VISIT{})
	Types["DFT_P03_COMMON_ORDER"] = reflect.TypeOf(DFT_P03_COMMON_ORDER{})
	Types["DFT_P03"] = reflect.TypeOf(DFT_P03{})
	Types["DFT_P03_FINANCIAL"] = reflect.TypeOf(DFT_P03_FINANCIAL{})
	Types["DFT_P03_FINANCIAL_COMMON_ORDER"] = reflect.TypeOf(DFT_P03_FINANCIAL_COMMON_ORDER{})
	Types["DFT_P03_FINANCIAL_OBSERVATION"] = reflect.TypeOf(DFT_P03_FINANCIAL_OBSERVATION{})
	Types["DFT_P03_FINANCIAL_ORDER"] = reflect.TypeOf(DFT_P03_FINANCIAL_ORDER{})
	Types["DFT_P03_FINANCIAL_PROCEDURE"] = reflect.TypeOf(DFT_P03_FINANCIAL_PROCEDURE{})
	Types["DFT_P03_INSURANCE"] = reflect.TypeOf(DFT_P03_INSURANCE{})
	Types["DFT_P03_OBSERVATION"] = reflect.TypeOf(DFT_P03_OBSERVATION{})
	Types["DFT_P03_ORDER"] = reflect.TypeOf(DFT_P03_ORDER{})
	Types["DFT_P11_COMMON_ORDER"] = reflect.TypeOf(DFT_P11_COMMON_ORDER{})
	Types["DFT_P11"] = reflect.TypeOf(DFT_P11{})
	Types["DFT_P11_FINANCIAL"] = reflect.TypeOf(DFT_P11_FINANCIAL{})
	Types["DFT_P11_FINANCIAL_COMMON_ORDER"] = reflect.TypeOf(DFT_P11_FINANCIAL_COMMON_ORDER{})
	Types["DFT_P11_FINANCIAL_INSURANCE"] = reflect.TypeOf(DFT_P11_FINANCIAL_INSURANCE{})
	Types["DFT_P11_FINANCIAL_OBSERVATION"] = reflect.TypeOf(DFT_P11_FINANCIAL_OBSERVATION{})
	Types["DFT_P11_FINANCIAL_ORDER"] = reflect.TypeOf(DFT_P11_FINANCIAL_ORDER{})
	Types["DFT_P11_FINANCIAL_PROCEDURE"] = reflect.TypeOf(DFT_P11_FINANCIAL_PROCEDURE{})
	Types["DFT_P11_INSURANCE"] = reflect.TypeOf(DFT_P11_INSURANCE{})
	Types["DFT_P11_OBSERVATION"] = reflect.TypeOf(DFT_P11_OBSERVATION{})
	Types["DFT_P11_ORDER"] = reflect.TypeOf(DFT_P11_ORDER{})
	Types["DOC_T12"] = reflect.TypeOf(DOC_T12{})
	Types["DOC_T12_RESULT"] = reflect.TypeOf(DOC_T12_RESULT{})
	Types["DSR_P04"] = reflect.TypeOf(DSR_P04{})
	Types["DSR_Q01"] = reflect.TypeOf(DSR_Q01{})
	Types["DSR_Q03"] = reflect.TypeOf(DSR_Q03{})
	Types["DSR_R03"] = reflect.TypeOf(DSR_R03{})
	Types["EAC_U07"] = reflect.TypeOf(EAC_U07{})
	Types["EAN_U09"] = reflect.TypeOf(EAN_U09{})
	Types["EAN_U09_NOTIFICATION"] = reflect.TypeOf(EAN_U09_NOTIFICATION{})
	Types["EAR_U08_COMMAND_RESPONSE"] = reflect.TypeOf(EAR_U08_COMMAND_RESPONSE{})
	Types["EAR_U08"] = reflect.TypeOf(EAR_U08{})
	Types["EDR_Q01"] = reflect.TypeOf(EDR_Q01{})
	Types["EDR_R07"] = reflect.TypeOf(EDR_R07{})
	Types["EQQ_Q01"] = reflect.TypeOf(EQQ_Q01{})
	Types["EQQ_Q04"] = reflect.TypeOf(EQQ_Q04{})
	Types["ERP_Q01"] = reflect.TypeOf(ERP_Q01{})
	Types["ERP_R09"] = reflect.TypeOf(ERP_R09{})
	Types["ESR_U02"] = reflect.TypeOf(ESR_U02{})
	Types["ESU_U01"] = reflect.TypeOf(ESU_U01{})
	Types["INR_U06"] = reflect.TypeOf(INR_U06{})
	Types["INU_U05"] = reflect.TypeOf(INU_U05{})
	Types["LSU_U12"] = reflect.TypeOf(LSU_U12{})
	Types["MCF_Q02"] = reflect.TypeOf(MCF_Q02{})
	Types["MDM_T01"] = reflect.TypeOf(MDM_T01{})
	Types["MDM_T02"] = reflect.TypeOf(MDM_T02{})
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
	Types["MFN_M04"] = reflect.TypeOf(MFN_M04{})
	Types["MFN_M04_MF_CDM"] = reflect.TypeOf(MFN_M04_MF_CDM{})
	Types["MFN_M05"] = reflect.TypeOf(MFN_M05{})
	Types["MFN_M05_MF_LOCATION"] = reflect.TypeOf(MFN_M05_MF_LOCATION{})
	Types["MFN_M05_MF_LOC_DEPT"] = reflect.TypeOf(MFN_M05_MF_LOC_DEPT{})
	Types["MFN_M06"] = reflect.TypeOf(MFN_M06{})
	Types["MFN_M06_MF_CDM"] = reflect.TypeOf(MFN_M06_MF_CDM{})
	Types["MFN_M06_MF_CLIN_STUDY"] = reflect.TypeOf(MFN_M06_MF_CLIN_STUDY{})
	Types["MFN_M06_MF_PHASE_SCHED_DETAIL"] = reflect.TypeOf(MFN_M06_MF_PHASE_SCHED_DETAIL{})
	Types["MFN_M07"] = reflect.TypeOf(MFN_M07{})
	Types["MFN_M07_MF_CLIN_STUDY"] = reflect.TypeOf(MFN_M07_MF_CLIN_STUDY{})
	Types["MFN_M07_MF_CLIN_STUDY_SCHED"] = reflect.TypeOf(MFN_M07_MF_CLIN_STUDY_SCHED{})
	Types["MFN_M07_MF_PHASE_SCHED_DETAIL"] = reflect.TypeOf(MFN_M07_MF_PHASE_SCHED_DETAIL{})
	Types["MFN_M08"] = reflect.TypeOf(MFN_M08{})
	Types["MFN_M08_MF_NUMERIC_OBSERVATION"] = reflect.TypeOf(MFN_M08_MF_NUMERIC_OBSERVATION{})
	Types["MFN_M08_MF_TEST_NUMERIC"] = reflect.TypeOf(MFN_M08_MF_TEST_NUMERIC{})
	Types["MFN_M09"] = reflect.TypeOf(MFN_M09{})
	Types["MFN_M09_MF_TEST_CATEGORICAL"] = reflect.TypeOf(MFN_M09_MF_TEST_CATEGORICAL{})
	Types["MFN_M09_MF_TEST_CAT_DETAIL"] = reflect.TypeOf(MFN_M09_MF_TEST_CAT_DETAIL{})
	Types["MFN_M10"] = reflect.TypeOf(MFN_M10{})
	Types["MFN_M10_MF_TEST_BATTERIES"] = reflect.TypeOf(MFN_M10_MF_TEST_BATTERIES{})
	Types["MFN_M10_MF_TEST_BATT_DETAIL"] = reflect.TypeOf(MFN_M10_MF_TEST_BATT_DETAIL{})
	Types["MFN_M11"] = reflect.TypeOf(MFN_M11{})
	Types["MFN_M11_MF_TEST_CALCULATED"] = reflect.TypeOf(MFN_M11_MF_TEST_CALCULATED{})
	Types["MFN_M11_MF_TEST_CALC_DETAIL"] = reflect.TypeOf(MFN_M11_MF_TEST_CALC_DETAIL{})
	Types["MFN_M12"] = reflect.TypeOf(MFN_M12{})
	Types["MFN_M12_MF_OBS_ATTRIBUTES"] = reflect.TypeOf(MFN_M12_MF_OBS_ATTRIBUTES{})
	Types["MFQ_M01"] = reflect.TypeOf(MFQ_M01{})
	Types["MFQ_M02"] = reflect.TypeOf(MFQ_M02{})
	Types["MFQ_M03"] = reflect.TypeOf(MFQ_M03{})
	Types["MFR_M01"] = reflect.TypeOf(MFR_M01{})
	Types["MFR_M01_MF"] = reflect.TypeOf(MFR_M01_MF{})
	Types["MFR_M01_MF_QUERY"] = reflect.TypeOf(MFR_M01_MF_QUERY{})
	Types["MFR_M02"] = reflect.TypeOf(MFR_M02{})
	Types["MFR_M02_MF_STAFF"] = reflect.TypeOf(MFR_M02_MF_STAFF{})
	Types["MFR_M03"] = reflect.TypeOf(MFR_M03{})
	Types["MFR_M03_MF_TEST"] = reflect.TypeOf(MFR_M03_MF_TEST{})
	Types["NMD_N01_APP_STATS"] = reflect.TypeOf(NMD_N01_APP_STATS{})
	Types["NMD_N01_APP_STATUS"] = reflect.TypeOf(NMD_N01_APP_STATUS{})
	Types["NMD_N01_CLOCK"] = reflect.TypeOf(NMD_N01_CLOCK{})
	Types["NMD_N01_CLOCK_AND_STATS_WITH_NOTES"] = reflect.TypeOf(NMD_N01_CLOCK_AND_STATS_WITH_NOTES{})
	Types["NMD_N01"] = reflect.TypeOf(NMD_N01{})
	Types["NMD_N02_APP_STATS"] = reflect.TypeOf(NMD_N02_APP_STATS{})
	Types["NMD_N02_APP_STATUS"] = reflect.TypeOf(NMD_N02_APP_STATUS{})
	Types["NMD_N02_CLOCK"] = reflect.TypeOf(NMD_N02_CLOCK{})
	Types["NMD_N02_CLOCK_AND_STATS_WITH_NOTES"] = reflect.TypeOf(NMD_N02_CLOCK_AND_STATS_WITH_NOTES{})
	Types["NMD_N02"] = reflect.TypeOf(NMD_N02{})
	Types["NMQ_N01_CLOCK_AND_STATISTICS"] = reflect.TypeOf(NMQ_N01_CLOCK_AND_STATISTICS{})
	Types["NMQ_N01"] = reflect.TypeOf(NMQ_N01{})
	Types["NMQ_N01_QRY_WITH_DETAIL"] = reflect.TypeOf(NMQ_N01_QRY_WITH_DETAIL{})
	Types["NMQ_N02_CLOCK_AND_STATISTICS"] = reflect.TypeOf(NMQ_N02_CLOCK_AND_STATISTICS{})
	Types["NMQ_N02"] = reflect.TypeOf(NMQ_N02{})
	Types["NMQ_N02_QRY_WITH_DETAIL"] = reflect.TypeOf(NMQ_N02_QRY_WITH_DETAIL{})
	Types["NMR_N01_CLOCK_AND_STATS_WITH_NOTES_ALT"] = reflect.TypeOf(NMR_N01_CLOCK_AND_STATS_WITH_NOTES_ALT{})
	Types["NMR_N01"] = reflect.TypeOf(NMR_N01{})
	Types["NMR_N02_CLOCK_AND_STATS_WITH_NOTES_ALT"] = reflect.TypeOf(NMR_N02_CLOCK_AND_STATS_WITH_NOTES_ALT{})
	Types["NMR_N02"] = reflect.TypeOf(NMR_N02{})
	Types["OMD_O01"] = reflect.TypeOf(OMD_O01{})
	Types["OMD_O01_DIET"] = reflect.TypeOf(OMD_O01_DIET{})
	Types["OMD_O01_INSURANCE"] = reflect.TypeOf(OMD_O01_INSURANCE{})
	Types["OMD_O01_OBSERVATION"] = reflect.TypeOf(OMD_O01_OBSERVATION{})
	Types["OMD_O01_ORDER_DIET"] = reflect.TypeOf(OMD_O01_ORDER_DIET{})
	Types["OMD_O01_ORDER_TRAY"] = reflect.TypeOf(OMD_O01_ORDER_TRAY{})
	Types["OMD_O01_PATIENT"] = reflect.TypeOf(OMD_O01_PATIENT{})
	Types["OMD_O01_PATIENT_VISIT"] = reflect.TypeOf(OMD_O01_PATIENT_VISIT{})
	Types["OMD_O03"] = reflect.TypeOf(OMD_O03{})
	Types["OMD_O03_DIET"] = reflect.TypeOf(OMD_O03_DIET{})
	Types["OMD_O03_INSURANCE"] = reflect.TypeOf(OMD_O03_INSURANCE{})
	Types["OMD_O03_OBSERVATION"] = reflect.TypeOf(OMD_O03_OBSERVATION{})
	Types["OMD_O03_ORDER_DIET"] = reflect.TypeOf(OMD_O03_ORDER_DIET{})
	Types["OMD_O03_ORDER_TRAY"] = reflect.TypeOf(OMD_O03_ORDER_TRAY{})
	Types["OMD_O03_PATIENT"] = reflect.TypeOf(OMD_O03_PATIENT{})
	Types["OMD_O03_PATIENT_VISIT"] = reflect.TypeOf(OMD_O03_PATIENT_VISIT{})
	Types["OMG_O19"] = reflect.TypeOf(OMG_O19{})
	Types["OMG_O19_INSURANCE"] = reflect.TypeOf(OMG_O19_INSURANCE{})
	Types["OMG_O19_OBSERVATION"] = reflect.TypeOf(OMG_O19_OBSERVATION{})
	Types["OMG_O19_OBSERVATION_PRIOR"] = reflect.TypeOf(OMG_O19_OBSERVATION_PRIOR{})
	Types["OMG_O19_ORDER"] = reflect.TypeOf(OMG_O19_ORDER{})
	Types["OMG_O19_ORDER_PRIOR"] = reflect.TypeOf(OMG_O19_ORDER_PRIOR{})
	Types["OMG_O19_PATIENT"] = reflect.TypeOf(OMG_O19_PATIENT{})
	Types["OMG_O19_PATIENT_PRIOR"] = reflect.TypeOf(OMG_O19_PATIENT_PRIOR{})
	Types["OMG_O19_PATIENT_VISIT"] = reflect.TypeOf(OMG_O19_PATIENT_VISIT{})
	Types["OMG_O19_PATIENT_VISIT_PRIOR"] = reflect.TypeOf(OMG_O19_PATIENT_VISIT_PRIOR{})
	Types["OMG_O19_PRIOR_RESULT"] = reflect.TypeOf(OMG_O19_PRIOR_RESULT{})
	Types["OML_O21_CONTAINER_1"] = reflect.TypeOf(OML_O21_CONTAINER_1{})
	Types["OML_O21_CONTAINER_2"] = reflect.TypeOf(OML_O21_CONTAINER_2{})
	Types["OML_O21"] = reflect.TypeOf(OML_O21{})
	Types["OML_O21_INSURANCE"] = reflect.TypeOf(OML_O21_INSURANCE{})
	Types["OML_O21_OBSERVATION"] = reflect.TypeOf(OML_O21_OBSERVATION{})
	Types["OML_O21_OBSERVATION_PRIOR"] = reflect.TypeOf(OML_O21_OBSERVATION_PRIOR{})
	Types["OML_O21_OBSERVATION_REQUEST"] = reflect.TypeOf(OML_O21_OBSERVATION_REQUEST{})
	Types["OML_O21_ORDER"] = reflect.TypeOf(OML_O21_ORDER{})
	Types["OML_O21_ORDER_GENERAL"] = reflect.TypeOf(OML_O21_ORDER_GENERAL{})
	Types["OML_O21_ORDER_PRIOR"] = reflect.TypeOf(OML_O21_ORDER_PRIOR{})
	Types["OML_O21_PATIENT"] = reflect.TypeOf(OML_O21_PATIENT{})
	Types["OML_O21_PATIENT_PRIOR"] = reflect.TypeOf(OML_O21_PATIENT_PRIOR{})
	Types["OML_O21_PATIENT_VISIT"] = reflect.TypeOf(OML_O21_PATIENT_VISIT{})
	Types["OML_O21_PATIENT_VISIT_PRIOR"] = reflect.TypeOf(OML_O21_PATIENT_VISIT_PRIOR{})
	Types["OML_O21_PRIOR_RESULT"] = reflect.TypeOf(OML_O21_PRIOR_RESULT{})
	Types["OMN_O01"] = reflect.TypeOf(OMN_O01{})
	Types["OMN_O01_INSURANCE"] = reflect.TypeOf(OMN_O01_INSURANCE{})
	Types["OMN_O01_OBSERVATION"] = reflect.TypeOf(OMN_O01_OBSERVATION{})
	Types["OMN_O01_ORDER"] = reflect.TypeOf(OMN_O01_ORDER{})
	Types["OMN_O01_ORDER_DETAIL"] = reflect.TypeOf(OMN_O01_ORDER_DETAIL{})
	Types["OMN_O01_PATIENT"] = reflect.TypeOf(OMN_O01_PATIENT{})
	Types["OMN_O01_PATIENT_VISIT"] = reflect.TypeOf(OMN_O01_PATIENT_VISIT{})
	Types["OMN_O07"] = reflect.TypeOf(OMN_O07{})
	Types["OMN_O07_INSURANCE"] = reflect.TypeOf(OMN_O07_INSURANCE{})
	Types["OMN_O07_OBSERVATION"] = reflect.TypeOf(OMN_O07_OBSERVATION{})
	Types["OMN_O07_ORDER"] = reflect.TypeOf(OMN_O07_ORDER{})
	Types["OMN_O07_PATIENT"] = reflect.TypeOf(OMN_O07_PATIENT{})
	Types["OMN_O07_PATIENT_VISIT"] = reflect.TypeOf(OMN_O07_PATIENT_VISIT{})
	Types["OMP_O09_COMPONENT"] = reflect.TypeOf(OMP_O09_COMPONENT{})
	Types["OMP_O09"] = reflect.TypeOf(OMP_O09{})
	Types["OMP_O09_INSURANCE"] = reflect.TypeOf(OMP_O09_INSURANCE{})
	Types["OMP_O09_OBSERVATION"] = reflect.TypeOf(OMP_O09_OBSERVATION{})
	Types["OMP_O09_ORDER"] = reflect.TypeOf(OMP_O09_ORDER{})
	Types["OMP_O09_PATIENT"] = reflect.TypeOf(OMP_O09_PATIENT{})
	Types["OMP_O09_PATIENT_VISIT"] = reflect.TypeOf(OMP_O09_PATIENT_VISIT{})
	Types["OMS_O01"] = reflect.TypeOf(OMS_O01{})
	Types["OMS_O01_INSURANCE"] = reflect.TypeOf(OMS_O01_INSURANCE{})
	Types["OMS_O01_OBSERVATION"] = reflect.TypeOf(OMS_O01_OBSERVATION{})
	Types["OMS_O01_ORDER"] = reflect.TypeOf(OMS_O01_ORDER{})
	Types["OMS_O01_ORDER_DETAIL"] = reflect.TypeOf(OMS_O01_ORDER_DETAIL{})
	Types["OMS_O01_PATIENT"] = reflect.TypeOf(OMS_O01_PATIENT{})
	Types["OMS_O01_PATIENT_VISIT"] = reflect.TypeOf(OMS_O01_PATIENT_VISIT{})
	Types["OMS_O05"] = reflect.TypeOf(OMS_O05{})
	Types["OMS_O05_INSURANCE"] = reflect.TypeOf(OMS_O05_INSURANCE{})
	Types["OMS_O05_OBSERVATION"] = reflect.TypeOf(OMS_O05_OBSERVATION{})
	Types["OMS_O05_ORDER"] = reflect.TypeOf(OMS_O05_ORDER{})
	Types["OMS_O05_PATIENT"] = reflect.TypeOf(OMS_O05_PATIENT{})
	Types["OMS_O05_PATIENT_VISIT"] = reflect.TypeOf(OMS_O05_PATIENT_VISIT{})
	Types["ORD_O02"] = reflect.TypeOf(ORD_O02{})
	Types["ORD_O02_ORDER_DIET"] = reflect.TypeOf(ORD_O02_ORDER_DIET{})
	Types["ORD_O02_ORDER_TRAY"] = reflect.TypeOf(ORD_O02_ORDER_TRAY{})
	Types["ORD_O02_PATIENT"] = reflect.TypeOf(ORD_O02_PATIENT{})
	Types["ORD_O02_RESPONSE"] = reflect.TypeOf(ORD_O02_RESPONSE{})
	Types["ORD_O04"] = reflect.TypeOf(ORD_O04{})
	Types["ORD_O04_ORDER_DIET"] = reflect.TypeOf(ORD_O04_ORDER_DIET{})
	Types["ORD_O04_ORDER_TRAY"] = reflect.TypeOf(ORD_O04_ORDER_TRAY{})
	Types["ORD_O04_PATIENT"] = reflect.TypeOf(ORD_O04_PATIENT{})
	Types["ORD_O04_RESPONSE"] = reflect.TypeOf(ORD_O04_RESPONSE{})
	Types["ORF_R04"] = reflect.TypeOf(ORF_R04{})
	Types["ORF_R04_OBSERVATION"] = reflect.TypeOf(ORF_R04_OBSERVATION{})
	Types["ORF_R04_ORDER"] = reflect.TypeOf(ORF_R04_ORDER{})
	Types["ORF_R04_PATIENT"] = reflect.TypeOf(ORF_R04_PATIENT{})
	Types["ORF_R04_QUERY_RESPONSE"] = reflect.TypeOf(ORF_R04_QUERY_RESPONSE{})
	Types["ORF_R04_RESPONSE"] = reflect.TypeOf(ORF_R04_RESPONSE{})
	Types["ORG_O20"] = reflect.TypeOf(ORG_O20{})
	Types["ORG_O20_ORDER"] = reflect.TypeOf(ORG_O20_ORDER{})
	Types["ORG_O20_PATIENT"] = reflect.TypeOf(ORG_O20_PATIENT{})
	Types["ORG_O20_RESPONSE"] = reflect.TypeOf(ORG_O20_RESPONSE{})
	Types["ORL_O22_CONTAINER"] = reflect.TypeOf(ORL_O22_CONTAINER{})
	Types["ORL_O22"] = reflect.TypeOf(ORL_O22{})
	Types["ORL_O22_GENERAL_ORDER"] = reflect.TypeOf(ORL_O22_GENERAL_ORDER{})
	Types["ORL_O22_OBSERVATION_REQUEST"] = reflect.TypeOf(ORL_O22_OBSERVATION_REQUEST{})
	Types["ORL_O22_ORDER"] = reflect.TypeOf(ORL_O22_ORDER{})
	Types["ORL_O22_PATIENT"] = reflect.TypeOf(ORL_O22_PATIENT{})
	Types["ORL_O22_RESPONSE"] = reflect.TypeOf(ORL_O22_RESPONSE{})
	Types["ORM_O01_CHOICE"] = reflect.TypeOf(ORM_O01_CHOICE{})
	Types["ORM_O01"] = reflect.TypeOf(ORM_O01{})
	Types["ORM_O01_INSURANCE"] = reflect.TypeOf(ORM_O01_INSURANCE{})
	Types["ORM_O01_OBSERVATION"] = reflect.TypeOf(ORM_O01_OBSERVATION{})
	Types["ORM_O01_ORDER"] = reflect.TypeOf(ORM_O01_ORDER{})
	Types["ORM_O01_ORDER_DETAIL"] = reflect.TypeOf(ORM_O01_ORDER_DETAIL{})
	Types["ORM_O01_PATIENT"] = reflect.TypeOf(ORM_O01_PATIENT{})
	Types["ORM_O01_PATIENT_VISIT"] = reflect.TypeOf(ORM_O01_PATIENT_VISIT{})
	Types["ORN_O02"] = reflect.TypeOf(ORN_O02{})
	Types["ORN_O02_ORDER"] = reflect.TypeOf(ORN_O02_ORDER{})
	Types["ORN_O02_PATIENT"] = reflect.TypeOf(ORN_O02_PATIENT{})
	Types["ORN_O02_RESPONSE"] = reflect.TypeOf(ORN_O02_RESPONSE{})
	Types["ORN_O08"] = reflect.TypeOf(ORN_O08{})
	Types["ORN_O08_ORDER"] = reflect.TypeOf(ORN_O08_ORDER{})
	Types["ORN_O08_PATIENT"] = reflect.TypeOf(ORN_O08_PATIENT{})
	Types["ORN_O08_RESPONSE"] = reflect.TypeOf(ORN_O08_RESPONSE{})
	Types["ORP_O10"] = reflect.TypeOf(ORP_O10{})
	Types["ORP_O10_ORDER"] = reflect.TypeOf(ORP_O10_ORDER{})
	Types["ORP_O10_ORDER_DETAIL"] = reflect.TypeOf(ORP_O10_ORDER_DETAIL{})
	Types["ORP_O10_PATIENT"] = reflect.TypeOf(ORP_O10_PATIENT{})
	Types["ORP_O10_RESPONSE"] = reflect.TypeOf(ORP_O10_RESPONSE{})
	Types["ORR_O02_CHOICE"] = reflect.TypeOf(ORR_O02_CHOICE{})
	Types["ORR_O02"] = reflect.TypeOf(ORR_O02{})
	Types["ORR_O02_ORDER"] = reflect.TypeOf(ORR_O02_ORDER{})
	Types["ORR_O02_ORDER_DETAIL"] = reflect.TypeOf(ORR_O02_ORDER_DETAIL{})
	Types["ORR_O02_PATIENT"] = reflect.TypeOf(ORR_O02_PATIENT{})
	Types["ORR_O02_RESPONSE"] = reflect.TypeOf(ORR_O02_RESPONSE{})
	Types["ORS_O02"] = reflect.TypeOf(ORS_O02{})
	Types["ORS_O02_ORDER"] = reflect.TypeOf(ORS_O02_ORDER{})
	Types["ORS_O02_PATIENT"] = reflect.TypeOf(ORS_O02_PATIENT{})
	Types["ORS_O02_RESPONSE"] = reflect.TypeOf(ORS_O02_RESPONSE{})
	Types["ORS_O06"] = reflect.TypeOf(ORS_O06{})
	Types["ORS_O06_ORDER"] = reflect.TypeOf(ORS_O06_ORDER{})
	Types["ORS_O06_PATIENT"] = reflect.TypeOf(ORS_O06_PATIENT{})
	Types["ORS_O06_RSPONSE"] = reflect.TypeOf(ORS_O06_RSPONSE{})
	Types["ORU_R01"] = reflect.TypeOf(ORU_R01{})
	Types["ORU_R01_OBSERVATION"] = reflect.TypeOf(ORU_R01_OBSERVATION{})
	Types["ORU_R01_ORDER_OBSERVATION"] = reflect.TypeOf(ORU_R01_ORDER_OBSERVATION{})
	Types["ORU_R01_PATIENT"] = reflect.TypeOf(ORU_R01_PATIENT{})
	Types["ORU_R01_PATIENT_RESULT"] = reflect.TypeOf(ORU_R01_PATIENT_RESULT{})
	Types["ORU_R01_RESPONSE"] = reflect.TypeOf(ORU_R01_RESPONSE{})
	Types["ORU_R01_VISIT"] = reflect.TypeOf(ORU_R01_VISIT{})
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
	Types["OSQ_Q06"] = reflect.TypeOf(OSQ_Q06{})
	Types["OSR_Q06_CHOICE"] = reflect.TypeOf(OSR_Q06_CHOICE{})
	Types["OSR_Q06"] = reflect.TypeOf(OSR_Q06{})
	Types["OSR_Q06_OBSERVATION"] = reflect.TypeOf(OSR_Q06_OBSERVATION{})
	Types["OSR_Q06_ORDER"] = reflect.TypeOf(OSR_Q06_ORDER{})
	Types["OSR_Q06_PATIENT"] = reflect.TypeOf(OSR_Q06_PATIENT{})
	Types["OSR_Q06_RESPONSE"] = reflect.TypeOf(OSR_Q06_RESPONSE{})
	Types["OUL_R21_CONTAINER"] = reflect.TypeOf(OUL_R21_CONTAINER{})
	Types["OUL_R21"] = reflect.TypeOf(OUL_R21{})
	Types["OUL_R21_OBSERVATION"] = reflect.TypeOf(OUL_R21_OBSERVATION{})
	Types["OUL_R21_ORDER_OBSERVATION"] = reflect.TypeOf(OUL_R21_ORDER_OBSERVATION{})
	Types["OUL_R21_PATIENT"] = reflect.TypeOf(OUL_R21_PATIENT{})
	Types["OUL_R21_VISIT"] = reflect.TypeOf(OUL_R21_VISIT{})
	Types["PEX_P07_ASSOCIATED_PERSON"] = reflect.TypeOf(PEX_P07_ASSOCIATED_PERSON{})
	Types["PEX_P07_ASSOCIATED_RX_ADMIN"] = reflect.TypeOf(PEX_P07_ASSOCIATED_RX_ADMIN{})
	Types["PEX_P07_ASSOCIATED_RX_ORDER"] = reflect.TypeOf(PEX_P07_ASSOCIATED_RX_ORDER{})
	Types["PEX_P07"] = reflect.TypeOf(PEX_P07{})
	Types["PEX_P07_EXPERIENCE"] = reflect.TypeOf(PEX_P07_EXPERIENCE{})
	Types["PEX_P07_PEX_CAUSE"] = reflect.TypeOf(PEX_P07_PEX_CAUSE{})
	Types["PEX_P07_PEX_OBSERVATION"] = reflect.TypeOf(PEX_P07_PEX_OBSERVATION{})
	Types["PEX_P07_RX_ADMINISTRATION"] = reflect.TypeOf(PEX_P07_RX_ADMINISTRATION{})
	Types["PEX_P07_RX_ORDER"] = reflect.TypeOf(PEX_P07_RX_ORDER{})
	Types["PEX_P07_STUDY"] = reflect.TypeOf(PEX_P07_STUDY{})
	Types["PEX_P07_VISIT"] = reflect.TypeOf(PEX_P07_VISIT{})
	Types["PGL_PC6_CHOICE"] = reflect.TypeOf(PGL_PC6_CHOICE{})
	Types["PGL_PC6"] = reflect.TypeOf(PGL_PC6{})
	Types["PGL_PC6_GOAL"] = reflect.TypeOf(PGL_PC6_GOAL{})
	Types["PGL_PC6_GOAL_ROLE"] = reflect.TypeOf(PGL_PC6_GOAL_ROLE{})
	Types["PGL_PC6_OBRRXO_SUPPGRP"] = reflect.TypeOf(PGL_PC6_OBRRXO_SUPPGRP{})
	Types["PGL_PC6_OBSERVATION"] = reflect.TypeOf(PGL_PC6_OBSERVATION{})
	Types["PGL_PC6_ORDER"] = reflect.TypeOf(PGL_PC6_ORDER{})
	Types["PGL_PC6_ORDER_DETAIL"] = reflect.TypeOf(PGL_PC6_ORDER_DETAIL{})
	Types["PGL_PC6_ORDER_OBSERVATION"] = reflect.TypeOf(PGL_PC6_ORDER_OBSERVATION{})
	Types["PGL_PC6_PATHWAY"] = reflect.TypeOf(PGL_PC6_PATHWAY{})
	Types["PGL_PC6_PATIENT_VISIT"] = reflect.TypeOf(PGL_PC6_PATIENT_VISIT{})
	Types["PGL_PC6_PROBLEM"] = reflect.TypeOf(PGL_PC6_PROBLEM{})
	Types["PGL_PC6_PROBLEM_OBSERVATION"] = reflect.TypeOf(PGL_PC6_PROBLEM_OBSERVATION{})
	Types["PGL_PC6_PROBLEM_ROLE"] = reflect.TypeOf(PGL_PC6_PROBLEM_ROLE{})
	Types["PIN_I07"] = reflect.TypeOf(PIN_I07{})
	Types["PIN_I07_GUARANTOR_INSURANCE"] = reflect.TypeOf(PIN_I07_GUARANTOR_INSURANCE{})
	Types["PIN_I07_INSURANCE"] = reflect.TypeOf(PIN_I07_INSURANCE{})
	Types["PIN_I07_PROVIDER"] = reflect.TypeOf(PIN_I07_PROVIDER{})
	Types["PMU_B01"] = reflect.TypeOf(PMU_B01{})
	Types["PMU_B03"] = reflect.TypeOf(PMU_B03{})
	Types["PMU_B04"] = reflect.TypeOf(PMU_B04{})
	Types["PPG_PCG_CHOICE"] = reflect.TypeOf(PPG_PCG_CHOICE{})
	Types["PPG_PCG"] = reflect.TypeOf(PPG_PCG{})
	Types["PPG_PCG_GOAL"] = reflect.TypeOf(PPG_PCG_GOAL{})
	Types["PPG_PCG_GOAL_OBSERVATION"] = reflect.TypeOf(PPG_PCG_GOAL_OBSERVATION{})
	Types["PPG_PCG_GOAL_ROLE"] = reflect.TypeOf(PPG_PCG_GOAL_ROLE{})
	Types["PPG_PCG_OBRRXO_SUPPGRP"] = reflect.TypeOf(PPG_PCG_OBRRXO_SUPPGRP{})
	Types["PPG_PCG_ORDER"] = reflect.TypeOf(PPG_PCG_ORDER{})
	Types["PPG_PCG_ORDER_DETAIL"] = reflect.TypeOf(PPG_PCG_ORDER_DETAIL{})
	Types["PPG_PCG_ORDER_OBSERVATION"] = reflect.TypeOf(PPG_PCG_ORDER_OBSERVATION{})
	Types["PPG_PCG_PATHWAY"] = reflect.TypeOf(PPG_PCG_PATHWAY{})
	Types["PPG_PCG_PATHWAY_ROLE"] = reflect.TypeOf(PPG_PCG_PATHWAY_ROLE{})
	Types["PPG_PCG_PATIENT_VISIT"] = reflect.TypeOf(PPG_PCG_PATIENT_VISIT{})
	Types["PPG_PCG_PROBLEM"] = reflect.TypeOf(PPG_PCG_PROBLEM{})
	Types["PPG_PCG_PROBLEM_OBSERVATION"] = reflect.TypeOf(PPG_PCG_PROBLEM_OBSERVATION{})
	Types["PPG_PCG_PROBLEM_ROLE"] = reflect.TypeOf(PPG_PCG_PROBLEM_ROLE{})
	Types["PPP_PCB_CHOICE"] = reflect.TypeOf(PPP_PCB_CHOICE{})
	Types["PPP_PCB"] = reflect.TypeOf(PPP_PCB{})
	Types["PPP_PCB_GOAL"] = reflect.TypeOf(PPP_PCB_GOAL{})
	Types["PPP_PCB_GOAL_OBSERVATION"] = reflect.TypeOf(PPP_PCB_GOAL_OBSERVATION{})
	Types["PPP_PCB_GOAL_ROLE"] = reflect.TypeOf(PPP_PCB_GOAL_ROLE{})
	Types["PPP_PCB_ORDER"] = reflect.TypeOf(PPP_PCB_ORDER{})
	Types["PPP_PCB_ORDER_DETAIL"] = reflect.TypeOf(PPP_PCB_ORDER_DETAIL{})
	Types["PPP_PCB_ORDER_OBSERVATION"] = reflect.TypeOf(PPP_PCB_ORDER_OBSERVATION{})
	Types["PPP_PCB_PATHWAY"] = reflect.TypeOf(PPP_PCB_PATHWAY{})
	Types["PPP_PCB_PATHWAY_ROLE"] = reflect.TypeOf(PPP_PCB_PATHWAY_ROLE{})
	Types["PPP_PCB_PATIENT_VISIT"] = reflect.TypeOf(PPP_PCB_PATIENT_VISIT{})
	Types["PPP_PCB_PROBLEM"] = reflect.TypeOf(PPP_PCB_PROBLEM{})
	Types["PPP_PCB_PROBLEM_OBSERVATION"] = reflect.TypeOf(PPP_PCB_PROBLEM_OBSERVATION{})
	Types["PPP_PCB_PROBLEM_ROLE"] = reflect.TypeOf(PPP_PCB_PROBLEM_ROLE{})
	Types["PPR_PC1_CHOICE"] = reflect.TypeOf(PPR_PC1_CHOICE{})
	Types["PPR_PC1"] = reflect.TypeOf(PPR_PC1{})
	Types["PPR_PC1_GOAL"] = reflect.TypeOf(PPR_PC1_GOAL{})
	Types["PPR_PC1_GOAL_OBSERVATION"] = reflect.TypeOf(PPR_PC1_GOAL_OBSERVATION{})
	Types["PPR_PC1_GOAL_ROLE"] = reflect.TypeOf(PPR_PC1_GOAL_ROLE{})
	Types["PPR_PC1_ORDER"] = reflect.TypeOf(PPR_PC1_ORDER{})
	Types["PPR_PC1_ORDER_DETAIL"] = reflect.TypeOf(PPR_PC1_ORDER_DETAIL{})
	Types["PPR_PC1_ORDER_OBSERVATION"] = reflect.TypeOf(PPR_PC1_ORDER_OBSERVATION{})
	Types["PPR_PC1_PATHWAY"] = reflect.TypeOf(PPR_PC1_PATHWAY{})
	Types["PPR_PC1_PATHWAY_OBSERVATION"] = reflect.TypeOf(PPR_PC1_PATHWAY_OBSERVATION{})
	Types["PPR_PC1_PATIENT_VISIT"] = reflect.TypeOf(PPR_PC1_PATIENT_VISIT{})
	Types["PPR_PC1_PROBLEM"] = reflect.TypeOf(PPR_PC1_PROBLEM{})
	Types["PPR_PC1_PROBLEM_OBSERVATION"] = reflect.TypeOf(PPR_PC1_PROBLEM_OBSERVATION{})
	Types["PPR_PC1_PROBLEM_ROLE"] = reflect.TypeOf(PPR_PC1_PROBLEM_ROLE{})
	Types["PPT_PCL_CHOICE"] = reflect.TypeOf(PPT_PCL_CHOICE{})
	Types["PPT_PCL"] = reflect.TypeOf(PPT_PCL{})
	Types["PPT_PCL_GOAL"] = reflect.TypeOf(PPT_PCL_GOAL{})
	Types["PPT_PCL_GOAL_OBSERVATION"] = reflect.TypeOf(PPT_PCL_GOAL_OBSERVATION{})
	Types["PPT_PCL_GOAL_ROLE"] = reflect.TypeOf(PPT_PCL_GOAL_ROLE{})
	Types["PPT_PCL_ORDER"] = reflect.TypeOf(PPT_PCL_ORDER{})
	Types["PPT_PCL_ORDER_DETAIL"] = reflect.TypeOf(PPT_PCL_ORDER_DETAIL{})
	Types["PPT_PCL_ORDER_OBSERVATION"] = reflect.TypeOf(PPT_PCL_ORDER_OBSERVATION{})
	Types["PPT_PCL_PATHWAY"] = reflect.TypeOf(PPT_PCL_PATHWAY{})
	Types["PPT_PCL_PATHWAY_ROLE"] = reflect.TypeOf(PPT_PCL_PATHWAY_ROLE{})
	Types["PPT_PCL_PATIENT"] = reflect.TypeOf(PPT_PCL_PATIENT{})
	Types["PPT_PCL_PATIENT_VISIT"] = reflect.TypeOf(PPT_PCL_PATIENT_VISIT{})
	Types["PPT_PCL_PROBLEM"] = reflect.TypeOf(PPT_PCL_PROBLEM{})
	Types["PPT_PCL_PROBLEM_OBSERVATION"] = reflect.TypeOf(PPT_PCL_PROBLEM_OBSERVATION{})
	Types["PPT_PCL_PROBLEM_ROLE"] = reflect.TypeOf(PPT_PCL_PROBLEM_ROLE{})
	Types["PPV_PCA_CHOICE"] = reflect.TypeOf(PPV_PCA_CHOICE{})
	Types["PPV_PCA"] = reflect.TypeOf(PPV_PCA{})
	Types["PPV_PCA_GOAL"] = reflect.TypeOf(PPV_PCA_GOAL{})
	Types["PPV_PCA_GOAL_OBSERVATION"] = reflect.TypeOf(PPV_PCA_GOAL_OBSERVATION{})
	Types["PPV_PCA_GOAL_PATHWAY"] = reflect.TypeOf(PPV_PCA_GOAL_PATHWAY{})
	Types["PPV_PCA_GOAL_ROLE"] = reflect.TypeOf(PPV_PCA_GOAL_ROLE{})
	Types["PPV_PCA_ORDER"] = reflect.TypeOf(PPV_PCA_ORDER{})
	Types["PPV_PCA_ORDER_DETAIL"] = reflect.TypeOf(PPV_PCA_ORDER_DETAIL{})
	Types["PPV_PCA_ORDER_OBSERVATION"] = reflect.TypeOf(PPV_PCA_ORDER_OBSERVATION{})
	Types["PPV_PCA_PATIENT"] = reflect.TypeOf(PPV_PCA_PATIENT{})
	Types["PPV_PCA_PATIENT_VISIT"] = reflect.TypeOf(PPV_PCA_PATIENT_VISIT{})
	Types["PPV_PCA_PROBLEM"] = reflect.TypeOf(PPV_PCA_PROBLEM{})
	Types["PPV_PCA_PROBLEM_OBSERVATION"] = reflect.TypeOf(PPV_PCA_PROBLEM_OBSERVATION{})
	Types["PPV_PCA_PROBLEM_ROLE"] = reflect.TypeOf(PPV_PCA_PROBLEM_ROLE{})
	Types["PRR_PC5_CHOICE"] = reflect.TypeOf(PRR_PC5_CHOICE{})
	Types["PRR_PC5"] = reflect.TypeOf(PRR_PC5{})
	Types["PRR_PC5_GOAL"] = reflect.TypeOf(PRR_PC5_GOAL{})
	Types["PRR_PC5_GOAL_OBSERVATION"] = reflect.TypeOf(PRR_PC5_GOAL_OBSERVATION{})
	Types["PRR_PC5_GOAL_ROLE"] = reflect.TypeOf(PRR_PC5_GOAL_ROLE{})
	Types["PRR_PC5_ORDER"] = reflect.TypeOf(PRR_PC5_ORDER{})
	Types["PRR_PC5_ORDER_DETAIL"] = reflect.TypeOf(PRR_PC5_ORDER_DETAIL{})
	Types["PRR_PC5_ORDER_OBSERVATION"] = reflect.TypeOf(PRR_PC5_ORDER_OBSERVATION{})
	Types["PRR_PC5_PATIENT"] = reflect.TypeOf(PRR_PC5_PATIENT{})
	Types["PRR_PC5_PATIENT_VISIT"] = reflect.TypeOf(PRR_PC5_PATIENT_VISIT{})
	Types["PRR_PC5_PROBLEM"] = reflect.TypeOf(PRR_PC5_PROBLEM{})
	Types["PRR_PC5_PROBLEM_OBSERVATION"] = reflect.TypeOf(PRR_PC5_PROBLEM_OBSERVATION{})
	Types["PRR_PC5_PROBLEM_PATHWAY"] = reflect.TypeOf(PRR_PC5_PROBLEM_PATHWAY{})
	Types["PRR_PC5_PROBLEM_ROLE"] = reflect.TypeOf(PRR_PC5_PROBLEM_ROLE{})
	Types["PTR_PCF_CHOICE"] = reflect.TypeOf(PTR_PCF_CHOICE{})
	Types["PTR_PCF"] = reflect.TypeOf(PTR_PCF{})
	Types["PTR_PCF_GOAL"] = reflect.TypeOf(PTR_PCF_GOAL{})
	Types["PTR_PCF_GOAL_OBSERVATION"] = reflect.TypeOf(PTR_PCF_GOAL_OBSERVATION{})
	Types["PTR_PCF_GOAL_ROLE"] = reflect.TypeOf(PTR_PCF_GOAL_ROLE{})
	Types["PTR_PCF_ORDER"] = reflect.TypeOf(PTR_PCF_ORDER{})
	Types["PTR_PCF_ORDER_DETAIL"] = reflect.TypeOf(PTR_PCF_ORDER_DETAIL{})
	Types["PTR_PCF_ORDER_OBSERVATION"] = reflect.TypeOf(PTR_PCF_ORDER_OBSERVATION{})
	Types["PTR_PCF_PATHWAY"] = reflect.TypeOf(PTR_PCF_PATHWAY{})
	Types["PTR_PCF_PATHWAY_ROLE"] = reflect.TypeOf(PTR_PCF_PATHWAY_ROLE{})
	Types["PTR_PCF_PATIENT"] = reflect.TypeOf(PTR_PCF_PATIENT{})
	Types["PTR_PCF_PATIENT_VISIT"] = reflect.TypeOf(PTR_PCF_PATIENT_VISIT{})
	Types["PTR_PCF_PROBLEM"] = reflect.TypeOf(PTR_PCF_PROBLEM{})
	Types["PTR_PCF_PROBLEM_OBSERVATION"] = reflect.TypeOf(PTR_PCF_PROBLEM_OBSERVATION{})
	Types["PTR_PCF_PROBLEM_ROLE"] = reflect.TypeOf(PTR_PCF_PROBLEM_ROLE{})
	Types["QBP_K13"] = reflect.TypeOf(QBP_K13{})
	Types["QBP_K13_ROW_DEFINITION"] = reflect.TypeOf(QBP_K13_ROW_DEFINITION{})
	Types["QBP_Q11"] = reflect.TypeOf(QBP_Q11{})
	Types["QBP_Q13"] = reflect.TypeOf(QBP_Q13{})
	Types["QBP_Q13_QBP"] = reflect.TypeOf(QBP_Q13_QBP{})
	Types["QBP_Q15"] = reflect.TypeOf(QBP_Q15{})
	Types["QBP_Q21"] = reflect.TypeOf(QBP_Q21{})
	Types["QBP_Qnn"] = reflect.TypeOf(QBP_Qnn{})
	Types["QBP_Z73"] = reflect.TypeOf(QBP_Z73{})
	Types["QCK_Q02"] = reflect.TypeOf(QCK_Q02{})
	Types["QCN_J01"] = reflect.TypeOf(QCN_J01{})
	Types["QRY_A19"] = reflect.TypeOf(QRY_A19{})
	Types["QRY_P04"] = reflect.TypeOf(QRY_P04{})
	Types["QRY_PC4"] = reflect.TypeOf(QRY_PC4{})
	Types["QRY_Q01"] = reflect.TypeOf(QRY_Q01{})
	Types["QRY_Q02"] = reflect.TypeOf(QRY_Q02{})
	Types["QRY_R02"] = reflect.TypeOf(QRY_R02{})
	Types["QRY_T12"] = reflect.TypeOf(QRY_T12{})
	Types["QSB_Q16"] = reflect.TypeOf(QSB_Q16{})
	Types["QVR_Q17"] = reflect.TypeOf(QVR_Q17{})
	Types["RAR_RAR"] = reflect.TypeOf(RAR_RAR{})
	Types["RAR_RAR_DEFINITION"] = reflect.TypeOf(RAR_RAR_DEFINITION{})
	Types["RAR_RAR_ENCODING"] = reflect.TypeOf(RAR_RAR_ENCODING{})
	Types["RAR_RAR_ORDER"] = reflect.TypeOf(RAR_RAR_ORDER{})
	Types["RAR_RAR_PATIENT"] = reflect.TypeOf(RAR_RAR_PATIENT{})
	Types["RAS_O01_COMPONENTS"] = reflect.TypeOf(RAS_O01_COMPONENTS{})
	Types["RAS_O01"] = reflect.TypeOf(RAS_O01{})
	Types["RAS_O01_ENCODING"] = reflect.TypeOf(RAS_O01_ENCODING{})
	Types["RAS_O01_OBSERVATION"] = reflect.TypeOf(RAS_O01_OBSERVATION{})
	Types["RAS_O01_ORDER"] = reflect.TypeOf(RAS_O01_ORDER{})
	Types["RAS_O01_ORDER_DETAIL"] = reflect.TypeOf(RAS_O01_ORDER_DETAIL{})
	Types["RAS_O01_ORDER_DETAIL_SUPPLEMENT"] = reflect.TypeOf(RAS_O01_ORDER_DETAIL_SUPPLEMENT{})
	Types["RAS_O01_PATIENT"] = reflect.TypeOf(RAS_O01_PATIENT{})
	Types["RAS_O01_PATIENT_VISIT"] = reflect.TypeOf(RAS_O01_PATIENT_VISIT{})
	Types["RAS_O17_COMPONENTS"] = reflect.TypeOf(RAS_O17_COMPONENTS{})
	Types["RAS_O17"] = reflect.TypeOf(RAS_O17{})
	Types["RAS_O17_ENCODING"] = reflect.TypeOf(RAS_O17_ENCODING{})
	Types["RAS_O17_OBSERVATION"] = reflect.TypeOf(RAS_O17_OBSERVATION{})
	Types["RAS_O17_ORDER"] = reflect.TypeOf(RAS_O17_ORDER{})
	Types["RAS_O17_ORDER_DETAIL"] = reflect.TypeOf(RAS_O17_ORDER_DETAIL{})
	Types["RAS_O17_ORDER_DETAIL_SUPPLEMENT"] = reflect.TypeOf(RAS_O17_ORDER_DETAIL_SUPPLEMENT{})
	Types["RAS_O17_PATIENT"] = reflect.TypeOf(RAS_O17_PATIENT{})
	Types["RAS_O17_PATIENT_VISIT"] = reflect.TypeOf(RAS_O17_PATIENT_VISIT{})
	Types["RCI_I05"] = reflect.TypeOf(RCI_I05{})
	Types["RCI_I05_OBSERVATION"] = reflect.TypeOf(RCI_I05_OBSERVATION{})
	Types["RCI_I05_PROVIDER"] = reflect.TypeOf(RCI_I05_PROVIDER{})
	Types["RCI_I05_RESULTS"] = reflect.TypeOf(RCI_I05_RESULTS{})
	Types["RCL_I06"] = reflect.TypeOf(RCL_I06{})
	Types["RCL_I06_PROVIDER"] = reflect.TypeOf(RCL_I06_PROVIDER{})
	Types["RDE_O01_COMPONENT"] = reflect.TypeOf(RDE_O01_COMPONENT{})
	Types["RDE_O01"] = reflect.TypeOf(RDE_O01{})
	Types["RDE_O01_INSURANCE"] = reflect.TypeOf(RDE_O01_INSURANCE{})
	Types["RDE_O01_OBSERVATION"] = reflect.TypeOf(RDE_O01_OBSERVATION{})
	Types["RDE_O01_ORDER"] = reflect.TypeOf(RDE_O01_ORDER{})
	Types["RDE_O01_ORDER_DETAIL"] = reflect.TypeOf(RDE_O01_ORDER_DETAIL{})
	Types["RDE_O01_PATIENT"] = reflect.TypeOf(RDE_O01_PATIENT{})
	Types["RDE_O01_PATIENT_VISIT"] = reflect.TypeOf(RDE_O01_PATIENT_VISIT{})
	Types["RDE_O11_COMPONENT"] = reflect.TypeOf(RDE_O11_COMPONENT{})
	Types["RDE_O11"] = reflect.TypeOf(RDE_O11{})
	Types["RDE_O11_INSURANCE"] = reflect.TypeOf(RDE_O11_INSURANCE{})
	Types["RDE_O11_OBSERVATION"] = reflect.TypeOf(RDE_O11_OBSERVATION{})
	Types["RDE_O11_ORDER"] = reflect.TypeOf(RDE_O11_ORDER{})
	Types["RDE_O11_ORDER_DETAIL"] = reflect.TypeOf(RDE_O11_ORDER_DETAIL{})
	Types["RDE_O11_PATIENT"] = reflect.TypeOf(RDE_O11_PATIENT{})
	Types["RDE_O11_PATIENT_VISIT"] = reflect.TypeOf(RDE_O11_PATIENT_VISIT{})
	Types["RDO_O01_COMPONENT"] = reflect.TypeOf(RDO_O01_COMPONENT{})
	Types["RDO_O01"] = reflect.TypeOf(RDO_O01{})
	Types["RDO_O01_INSURANCE"] = reflect.TypeOf(RDO_O01_INSURANCE{})
	Types["RDO_O01_OBSERVATION"] = reflect.TypeOf(RDO_O01_OBSERVATION{})
	Types["RDO_O01_ORDER"] = reflect.TypeOf(RDO_O01_ORDER{})
	Types["RDO_O01_ORDER_DETAIL"] = reflect.TypeOf(RDO_O01_ORDER_DETAIL{})
	Types["RDO_O01_PATIENT"] = reflect.TypeOf(RDO_O01_PATIENT{})
	Types["RDO_O01_PATIENT_VISIT"] = reflect.TypeOf(RDO_O01_PATIENT_VISIT{})
	Types["RDR_RDR"] = reflect.TypeOf(RDR_RDR{})
	Types["RDR_RDR_DEFINITION"] = reflect.TypeOf(RDR_RDR_DEFINITION{})
	Types["RDR_RDR_DISPENSE"] = reflect.TypeOf(RDR_RDR_DISPENSE{})
	Types["RDR_RDR_ENCODING"] = reflect.TypeOf(RDR_RDR_ENCODING{})
	Types["RDR_RDR_ORDER"] = reflect.TypeOf(RDR_RDR_ORDER{})
	Types["RDR_RDR_PATIENT"] = reflect.TypeOf(RDR_RDR_PATIENT{})
	Types["RDS_O01_COMPONENT"] = reflect.TypeOf(RDS_O01_COMPONENT{})
	Types["RDS_O01"] = reflect.TypeOf(RDS_O01{})
	Types["RDS_O01_ENCODING"] = reflect.TypeOf(RDS_O01_ENCODING{})
	Types["RDS_O01_OBSERVATION"] = reflect.TypeOf(RDS_O01_OBSERVATION{})
	Types["RDS_O01_ORDER"] = reflect.TypeOf(RDS_O01_ORDER{})
	Types["RDS_O01_ORDER_DETAIL"] = reflect.TypeOf(RDS_O01_ORDER_DETAIL{})
	Types["RDS_O01_ORDER_DETAIL_SUPPLEMENT"] = reflect.TypeOf(RDS_O01_ORDER_DETAIL_SUPPLEMENT{})
	Types["RDS_O01_PATIENT"] = reflect.TypeOf(RDS_O01_PATIENT{})
	Types["RDS_O01_PATIENT_VISIT"] = reflect.TypeOf(RDS_O01_PATIENT_VISIT{})
	Types["RDS_O13_COMPONENT"] = reflect.TypeOf(RDS_O13_COMPONENT{})
	Types["RDS_O13"] = reflect.TypeOf(RDS_O13{})
	Types["RDS_O13_ENCODING"] = reflect.TypeOf(RDS_O13_ENCODING{})
	Types["RDS_O13_OBSERVATION"] = reflect.TypeOf(RDS_O13_OBSERVATION{})
	Types["RDS_O13_ORDER"] = reflect.TypeOf(RDS_O13_ORDER{})
	Types["RDS_O13_ORDER_DETAIL"] = reflect.TypeOf(RDS_O13_ORDER_DETAIL{})
	Types["RDS_O13_ORDER_DETAIL_SUPPLEMENT"] = reflect.TypeOf(RDS_O13_ORDER_DETAIL_SUPPLEMENT{})
	Types["RDS_O13_PATIENT"] = reflect.TypeOf(RDS_O13_PATIENT{})
	Types["RDS_O13_PATIENT_VISIT"] = reflect.TypeOf(RDS_O13_PATIENT_VISIT{})
	Types["RDY_K15"] = reflect.TypeOf(RDY_K15{})
	Types["REF_I12_AUTCTD_SUPPGRP2"] = reflect.TypeOf(REF_I12_AUTCTD_SUPPGRP2{})
	Types["REF_I12_AUTHORIZATION"] = reflect.TypeOf(REF_I12_AUTHORIZATION{})
	Types["REF_I12_AUTHORIZATION_CONTACT"] = reflect.TypeOf(REF_I12_AUTHORIZATION_CONTACT{})
	Types["REF_I12"] = reflect.TypeOf(REF_I12{})
	Types["REF_I12_INSURANCE"] = reflect.TypeOf(REF_I12_INSURANCE{})
	Types["REF_I12_OBSERVATION"] = reflect.TypeOf(REF_I12_OBSERVATION{})
	Types["REF_I12_PATIENT_VISIT"] = reflect.TypeOf(REF_I12_PATIENT_VISIT{})
	Types["REF_I12_PROCEDURE"] = reflect.TypeOf(REF_I12_PROCEDURE{})
	Types["REF_I12_PROVIDER"] = reflect.TypeOf(REF_I12_PROVIDER{})
	Types["REF_I12_PROVIDER_CONTACT"] = reflect.TypeOf(REF_I12_PROVIDER_CONTACT{})
	Types["REF_I12_RESULTS"] = reflect.TypeOf(REF_I12_RESULTS{})
	Types["REF_I12_RESULTS_NOTES"] = reflect.TypeOf(REF_I12_RESULTS_NOTES{})
	Types["REF_I12_VISIT"] = reflect.TypeOf(REF_I12_VISIT{})
	Types["RER_RER"] = reflect.TypeOf(RER_RER{})
	Types["RER_RER_DEFINITION"] = reflect.TypeOf(RER_RER_DEFINITION{})
	Types["RER_RER_ORDER"] = reflect.TypeOf(RER_RER_ORDER{})
	Types["RER_RER_PATIENT"] = reflect.TypeOf(RER_RER_PATIENT{})
	Types["RGR_RGR"] = reflect.TypeOf(RGR_RGR{})
	Types["RGR_RGR_DEFINITION"] = reflect.TypeOf(RGR_RGR_DEFINITION{})
	Types["RGR_RGR_ENCODING"] = reflect.TypeOf(RGR_RGR_ENCODING{})
	Types["RGR_RGR_ORDER"] = reflect.TypeOf(RGR_RGR_ORDER{})
	Types["RGR_RGR_PATIENT"] = reflect.TypeOf(RGR_RGR_PATIENT{})
	Types["RGV_O01_COMPONENTS"] = reflect.TypeOf(RGV_O01_COMPONENTS{})
	Types["RGV_O01"] = reflect.TypeOf(RGV_O01{})
	Types["RGV_O01_ENCODING"] = reflect.TypeOf(RGV_O01_ENCODING{})
	Types["RGV_O01_GIVE"] = reflect.TypeOf(RGV_O01_GIVE{})
	Types["RGV_O01_OBSERVATION"] = reflect.TypeOf(RGV_O01_OBSERVATION{})
	Types["RGV_O01_ORDER"] = reflect.TypeOf(RGV_O01_ORDER{})
	Types["RGV_O01_ORDER_DETAIL"] = reflect.TypeOf(RGV_O01_ORDER_DETAIL{})
	Types["RGV_O01_ORDER_DETAIL_SUPPLEMENT"] = reflect.TypeOf(RGV_O01_ORDER_DETAIL_SUPPLEMENT{})
	Types["RGV_O01_PATIENT"] = reflect.TypeOf(RGV_O01_PATIENT{})
	Types["RGV_O01_PATIENT_VISIT"] = reflect.TypeOf(RGV_O01_PATIENT_VISIT{})
	Types["RGV_O15_COMPONENTS"] = reflect.TypeOf(RGV_O15_COMPONENTS{})
	Types["RGV_O15"] = reflect.TypeOf(RGV_O15{})
	Types["RGV_O15_ENCODING"] = reflect.TypeOf(RGV_O15_ENCODING{})
	Types["RGV_O15_GIVE"] = reflect.TypeOf(RGV_O15_GIVE{})
	Types["RGV_O15_OBSERVATION"] = reflect.TypeOf(RGV_O15_OBSERVATION{})
	Types["RGV_O15_ORDER"] = reflect.TypeOf(RGV_O15_ORDER{})
	Types["RGV_O15_ORDER_DETAIL"] = reflect.TypeOf(RGV_O15_ORDER_DETAIL{})
	Types["RGV_O15_ORDER_DETAIL_SUPPLEMENT"] = reflect.TypeOf(RGV_O15_ORDER_DETAIL_SUPPLEMENT{})
	Types["RGV_O15_PATIENT"] = reflect.TypeOf(RGV_O15_PATIENT{})
	Types["RGV_O15_PATIENT_VISIT"] = reflect.TypeOf(RGV_O15_PATIENT_VISIT{})
	Types["ROR_ROR"] = reflect.TypeOf(ROR_ROR{})
	Types["ROR_ROR_DEFINITION"] = reflect.TypeOf(ROR_ROR_DEFINITION{})
	Types["ROR_ROR_ORDER"] = reflect.TypeOf(ROR_ROR_ORDER{})
	Types["ROR_ROR_PATIENT"] = reflect.TypeOf(ROR_ROR_PATIENT{})
	Types["RPA_I08_AUTCTD_SUPPGRP2"] = reflect.TypeOf(RPA_I08_AUTCTD_SUPPGRP2{})
	Types["RPA_I08_AUTHORIZATION"] = reflect.TypeOf(RPA_I08_AUTHORIZATION{})
	Types["RPA_I08"] = reflect.TypeOf(RPA_I08{})
	Types["RPA_I08_INSURANCE"] = reflect.TypeOf(RPA_I08_INSURANCE{})
	Types["RPA_I08_OBSERVATION"] = reflect.TypeOf(RPA_I08_OBSERVATION{})
	Types["RPA_I08_PROCEDURE"] = reflect.TypeOf(RPA_I08_PROCEDURE{})
	Types["RPA_I08_PROVIDER"] = reflect.TypeOf(RPA_I08_PROVIDER{})
	Types["RPA_I08_RESULTS"] = reflect.TypeOf(RPA_I08_RESULTS{})
	Types["RPA_I08_VISIT"] = reflect.TypeOf(RPA_I08_VISIT{})
	Types["RPI_I01"] = reflect.TypeOf(RPI_I01{})
	Types["RPI_I01_GUARANTOR_INSURANCE"] = reflect.TypeOf(RPI_I01_GUARANTOR_INSURANCE{})
	Types["RPI_I01_INSURANCE"] = reflect.TypeOf(RPI_I01_INSURANCE{})
	Types["RPI_I01_PROVIDER"] = reflect.TypeOf(RPI_I01_PROVIDER{})
	Types["RPI_I04"] = reflect.TypeOf(RPI_I04{})
	Types["RPI_I04_GUARANTOR_INSURANCE"] = reflect.TypeOf(RPI_I04_GUARANTOR_INSURANCE{})
	Types["RPI_I04_INSURANCE"] = reflect.TypeOf(RPI_I04_INSURANCE{})
	Types["RPI_I04_PROVIDER"] = reflect.TypeOf(RPI_I04_PROVIDER{})
	Types["RPL_I02"] = reflect.TypeOf(RPL_I02{})
	Types["RPL_I02_PROVIDER"] = reflect.TypeOf(RPL_I02_PROVIDER{})
	Types["RPR_I03"] = reflect.TypeOf(RPR_I03{})
	Types["RPR_I03_PROVIDER"] = reflect.TypeOf(RPR_I03_PROVIDER{})
	Types["RQA_I08_AUTCTD_SUPPGRP2"] = reflect.TypeOf(RQA_I08_AUTCTD_SUPPGRP2{})
	Types["RQA_I08_AUTHORIZATION"] = reflect.TypeOf(RQA_I08_AUTHORIZATION{})
	Types["RQA_I08"] = reflect.TypeOf(RQA_I08{})
	Types["RQA_I08_GUARANTOR_INSURANCE"] = reflect.TypeOf(RQA_I08_GUARANTOR_INSURANCE{})
	Types["RQA_I08_INSURANCE"] = reflect.TypeOf(RQA_I08_INSURANCE{})
	Types["RQA_I08_OBSERVATION"] = reflect.TypeOf(RQA_I08_OBSERVATION{})
	Types["RQA_I08_PROCEDURE"] = reflect.TypeOf(RQA_I08_PROCEDURE{})
	Types["RQA_I08_PROVIDER"] = reflect.TypeOf(RQA_I08_PROVIDER{})
	Types["RQA_I08_RESULTS"] = reflect.TypeOf(RQA_I08_RESULTS{})
	Types["RQA_I08_VISIT"] = reflect.TypeOf(RQA_I08_VISIT{})
	Types["RQC_I05"] = reflect.TypeOf(RQC_I05{})
	Types["RQC_I05_PROVIDER"] = reflect.TypeOf(RQC_I05_PROVIDER{})
	Types["RQC_I06"] = reflect.TypeOf(RQC_I06{})
	Types["RQC_I06_PROVIDER"] = reflect.TypeOf(RQC_I06_PROVIDER{})
	Types["RQI_I01"] = reflect.TypeOf(RQI_I01{})
	Types["RQI_I01_GUARANTOR_INSURANCE"] = reflect.TypeOf(RQI_I01_GUARANTOR_INSURANCE{})
	Types["RQI_I01_INSURANCE"] = reflect.TypeOf(RQI_I01_INSURANCE{})
	Types["RQI_I01_PROVIDER"] = reflect.TypeOf(RQI_I01_PROVIDER{})
	Types["RQP_I04"] = reflect.TypeOf(RQP_I04{})
	Types["RQP_I04_PROVIDER"] = reflect.TypeOf(RQP_I04_PROVIDER{})
	Types["RQQ_Q01"] = reflect.TypeOf(RQQ_Q01{})
	Types["RQQ_Q09"] = reflect.TypeOf(RQQ_Q09{})
	Types["RRA_O02_ADMINISTRATION"] = reflect.TypeOf(RRA_O02_ADMINISTRATION{})
	Types["RRA_O02"] = reflect.TypeOf(RRA_O02{})
	Types["RRA_O02_ORDER"] = reflect.TypeOf(RRA_O02_ORDER{})
	Types["RRA_O02_PATIENT"] = reflect.TypeOf(RRA_O02_PATIENT{})
	Types["RRA_O02_RESPONSE"] = reflect.TypeOf(RRA_O02_RESPONSE{})
	Types["RRA_O18_ADMINISTRATION"] = reflect.TypeOf(RRA_O18_ADMINISTRATION{})
	Types["RRA_O18"] = reflect.TypeOf(RRA_O18{})
	Types["RRA_O18_ORDER"] = reflect.TypeOf(RRA_O18_ORDER{})
	Types["RRA_O18_PATIENT"] = reflect.TypeOf(RRA_O18_PATIENT{})
	Types["RRA_O18_RESPONSE"] = reflect.TypeOf(RRA_O18_RESPONSE{})
	Types["RRD_O02"] = reflect.TypeOf(RRD_O02{})
	Types["RRD_O02_DISPENSE"] = reflect.TypeOf(RRD_O02_DISPENSE{})
	Types["RRD_O02_ORDER"] = reflect.TypeOf(RRD_O02_ORDER{})
	Types["RRD_O02_PATIENT"] = reflect.TypeOf(RRD_O02_PATIENT{})
	Types["RRD_O02_RESPONSE"] = reflect.TypeOf(RRD_O02_RESPONSE{})
	Types["RRD_O14"] = reflect.TypeOf(RRD_O14{})
	Types["RRD_O14_DISPENSE"] = reflect.TypeOf(RRD_O14_DISPENSE{})
	Types["RRD_O14_ORDER"] = reflect.TypeOf(RRD_O14_ORDER{})
	Types["RRD_O14_PATIENT"] = reflect.TypeOf(RRD_O14_PATIENT{})
	Types["RRD_O14_RESPONSE"] = reflect.TypeOf(RRD_O14_RESPONSE{})
	Types["RRE_O02"] = reflect.TypeOf(RRE_O02{})
	Types["RRE_O02_ENCODING"] = reflect.TypeOf(RRE_O02_ENCODING{})
	Types["RRE_O02_ORDER"] = reflect.TypeOf(RRE_O02_ORDER{})
	Types["RRE_O02_PATIENT"] = reflect.TypeOf(RRE_O02_PATIENT{})
	Types["RRE_O02_RESPONSE"] = reflect.TypeOf(RRE_O02_RESPONSE{})
	Types["RRE_O12"] = reflect.TypeOf(RRE_O12{})
	Types["RRE_O12_ENCODING"] = reflect.TypeOf(RRE_O12_ENCODING{})
	Types["RRE_O12_ORDER"] = reflect.TypeOf(RRE_O12_ORDER{})
	Types["RRE_O12_PATIENT"] = reflect.TypeOf(RRE_O12_PATIENT{})
	Types["RRE_O12_RESPONSE"] = reflect.TypeOf(RRE_O12_RESPONSE{})
	Types["RRG_O02"] = reflect.TypeOf(RRG_O02{})
	Types["RRG_O02_GIVE"] = reflect.TypeOf(RRG_O02_GIVE{})
	Types["RRG_O02_ORDER"] = reflect.TypeOf(RRG_O02_ORDER{})
	Types["RRG_O02_PATIENT"] = reflect.TypeOf(RRG_O02_PATIENT{})
	Types["RRG_O02_RESPONSE"] = reflect.TypeOf(RRG_O02_RESPONSE{})
	Types["RRG_O16"] = reflect.TypeOf(RRG_O16{})
	Types["RRG_O16_GIVE"] = reflect.TypeOf(RRG_O16_GIVE{})
	Types["RRG_O16_ORDER"] = reflect.TypeOf(RRG_O16_ORDER{})
	Types["RRG_O16_PATIENT"] = reflect.TypeOf(RRG_O16_PATIENT{})
	Types["RRG_O16_RESPONSE"] = reflect.TypeOf(RRG_O16_RESPONSE{})
	Types["RRI_I12_AUTCTD_SUPPGRP2"] = reflect.TypeOf(RRI_I12_AUTCTD_SUPPGRP2{})
	Types["RRI_I12_AUTHORIZATION"] = reflect.TypeOf(RRI_I12_AUTHORIZATION{})
	Types["RRI_I12_AUTHORIZATION_CONTACT"] = reflect.TypeOf(RRI_I12_AUTHORIZATION_CONTACT{})
	Types["RRI_I12"] = reflect.TypeOf(RRI_I12{})
	Types["RRI_I12_OBSERVATION"] = reflect.TypeOf(RRI_I12_OBSERVATION{})
	Types["RRI_I12_PATIENT_VISIT"] = reflect.TypeOf(RRI_I12_PATIENT_VISIT{})
	Types["RRI_I12_PROCEDURE"] = reflect.TypeOf(RRI_I12_PROCEDURE{})
	Types["RRI_I12_PROVIDER"] = reflect.TypeOf(RRI_I12_PROVIDER{})
	Types["RRI_I12_PROVIDER_CONTACT"] = reflect.TypeOf(RRI_I12_PROVIDER_CONTACT{})
	Types["RRI_I12_RESULTS"] = reflect.TypeOf(RRI_I12_RESULTS{})
	Types["RRI_I12_RESULTS_NOTES"] = reflect.TypeOf(RRI_I12_RESULTS_NOTES{})
	Types["RRI_I12_VISIT"] = reflect.TypeOf(RRI_I12_VISIT{})
	Types["RRO_O02"] = reflect.TypeOf(RRO_O02{})
	Types["RRO_O02_ORDER"] = reflect.TypeOf(RRO_O02_ORDER{})
	Types["RRO_O02_ORDER_DETAIL"] = reflect.TypeOf(RRO_O02_ORDER_DETAIL{})
	Types["RRO_O02_PATIENT"] = reflect.TypeOf(RRO_O02_PATIENT{})
	Types["RRO_O02_RESPONSE"] = reflect.TypeOf(RRO_O02_RESPONSE{})
	Types["RSP_K11"] = reflect.TypeOf(RSP_K11{})
	Types["RSP_K13"] = reflect.TypeOf(RSP_K13{})
	Types["RSP_K13_ROW_DEFINITION"] = reflect.TypeOf(RSP_K13_ROW_DEFINITION{})
	Types["RSP_K15"] = reflect.TypeOf(RSP_K15{})
	Types["RSP_K21"] = reflect.TypeOf(RSP_K21{})
	Types["RSP_K21_QUERY_RESPONSE"] = reflect.TypeOf(RSP_K21_QUERY_RESPONSE{})
	Types["RSP_K22"] = reflect.TypeOf(RSP_K22{})
	Types["RSP_K22_QUERY_RESPONSE"] = reflect.TypeOf(RSP_K22_QUERY_RESPONSE{})
	Types["RSP_K23"] = reflect.TypeOf(RSP_K23{})
	Types["RSP_K24"] = reflect.TypeOf(RSP_K24{})
	Types["RSP_K25"] = reflect.TypeOf(RSP_K25{})
	Types["RSP_K25_STAFF"] = reflect.TypeOf(RSP_K25_STAFF{})
	Types["RSP_Z82_COMMON_ORDER"] = reflect.TypeOf(RSP_Z82_COMMON_ORDER{})
	Types["RSP_Z82"] = reflect.TypeOf(RSP_Z82{})
	Types["RSP_Z82_ENCODED_ORDER"] = reflect.TypeOf(RSP_Z82_ENCODED_ORDER{})
	Types["RSP_Z82_OBSERVATION"] = reflect.TypeOf(RSP_Z82_OBSERVATION{})
	Types["RSP_Z82_ORDER_DETAIL"] = reflect.TypeOf(RSP_Z82_ORDER_DETAIL{})
	Types["RSP_Z82_PATIENT"] = reflect.TypeOf(RSP_Z82_PATIENT{})
	Types["RSP_Z82_PATIENT_VISIT"] = reflect.TypeOf(RSP_Z82_PATIENT_VISIT{})
	Types["RSP_Z82_QUERY_RESPONSE"] = reflect.TypeOf(RSP_Z82_QUERY_RESPONSE{})
	Types["RSP_Z82_TREATMENT"] = reflect.TypeOf(RSP_Z82_TREATMENT{})
	Types["RSP_Z82_VISIT"] = reflect.TypeOf(RSP_Z82_VISIT{})
	Types["RSP_Z86_ADMINISTRATION"] = reflect.TypeOf(RSP_Z86_ADMINISTRATION{})
	Types["RSP_Z86_COMMON_ORDER"] = reflect.TypeOf(RSP_Z86_COMMON_ORDER{})
	Types["RSP_Z86"] = reflect.TypeOf(RSP_Z86{})
	Types["RSP_Z86_DISPENSE"] = reflect.TypeOf(RSP_Z86_DISPENSE{})
	Types["RSP_Z86_ENCODED_ORDER"] = reflect.TypeOf(RSP_Z86_ENCODED_ORDER{})
	Types["RSP_Z86_GIVE"] = reflect.TypeOf(RSP_Z86_GIVE{})
	Types["RSP_Z86_OBSERVATION"] = reflect.TypeOf(RSP_Z86_OBSERVATION{})
	Types["RSP_Z86_ORDER_DETAIL"] = reflect.TypeOf(RSP_Z86_ORDER_DETAIL{})
	Types["RSP_Z86_PATIENT"] = reflect.TypeOf(RSP_Z86_PATIENT{})
	Types["RSP_Z86_QUERY_RESPONSE"] = reflect.TypeOf(RSP_Z86_QUERY_RESPONSE{})
	Types["RSP_Z88_ALLERGY"] = reflect.TypeOf(RSP_Z88_ALLERGY{})
	Types["RSP_Z88_COMMON_ORDER"] = reflect.TypeOf(RSP_Z88_COMMON_ORDER{})
	Types["RSP_Z88_COMPONENT"] = reflect.TypeOf(RSP_Z88_COMPONENT{})
	Types["RSP_Z88"] = reflect.TypeOf(RSP_Z88{})
	Types["RSP_Z88_OBSERVATION"] = reflect.TypeOf(RSP_Z88_OBSERVATION{})
	Types["RSP_Z88_ORDER_DETAIL"] = reflect.TypeOf(RSP_Z88_ORDER_DETAIL{})
	Types["RSP_Z88_ORDER_ENCODED"] = reflect.TypeOf(RSP_Z88_ORDER_ENCODED{})
	Types["RSP_Z88_PATIENT"] = reflect.TypeOf(RSP_Z88_PATIENT{})
	Types["RSP_Z88_QUERY_RESPONSE"] = reflect.TypeOf(RSP_Z88_QUERY_RESPONSE{})
	Types["RSP_Z88_VISIT"] = reflect.TypeOf(RSP_Z88_VISIT{})
	Types["RSP_Z90_COMMON_ORDER"] = reflect.TypeOf(RSP_Z90_COMMON_ORDER{})
	Types["RSP_Z90"] = reflect.TypeOf(RSP_Z90{})
	Types["RSP_Z90_OBSERVATION"] = reflect.TypeOf(RSP_Z90_OBSERVATION{})
	Types["RSP_Z90_PATIENT"] = reflect.TypeOf(RSP_Z90_PATIENT{})
	Types["RSP_Z90_QUERY_RESPONSE"] = reflect.TypeOf(RSP_Z90_QUERY_RESPONSE{})
	Types["RSP_Z90_VISIT"] = reflect.TypeOf(RSP_Z90_VISIT{})
	Types["RTB_K13"] = reflect.TypeOf(RTB_K13{})
	Types["RTB_K13_ROW_DEFINITION"] = reflect.TypeOf(RTB_K13_ROW_DEFINITION{})
	Types["RTB_Knn"] = reflect.TypeOf(RTB_Knn{})
	Types["RTB_Q13"] = reflect.TypeOf(RTB_Q13{})
	Types["RTB_Q13_ROW_DEFINITION"] = reflect.TypeOf(RTB_Q13_ROW_DEFINITION{})
	Types["RTB_Z74"] = reflect.TypeOf(RTB_Z74{})
	Types["RTB_Z74_ROW_DEFINITION"] = reflect.TypeOf(RTB_Z74_ROW_DEFINITION{})
	Types["SIU_S12"] = reflect.TypeOf(SIU_S12{})
	Types["SIU_S12_GENERAL_RESOURCE"] = reflect.TypeOf(SIU_S12_GENERAL_RESOURCE{})
	Types["SIU_S12_LOCATION_RESOURCE"] = reflect.TypeOf(SIU_S12_LOCATION_RESOURCE{})
	Types["SIU_S12_PATIENT"] = reflect.TypeOf(SIU_S12_PATIENT{})
	Types["SIU_S12_PERSONNEL_RESOURCE"] = reflect.TypeOf(SIU_S12_PERSONNEL_RESOURCE{})
	Types["SIU_S12_RESOURCES"] = reflect.TypeOf(SIU_S12_RESOURCES{})
	Types["SIU_S12_SERVICE"] = reflect.TypeOf(SIU_S12_SERVICE{})
	Types["SPQ_Q01"] = reflect.TypeOf(SPQ_Q01{})
	Types["SPQ_Q08"] = reflect.TypeOf(SPQ_Q08{})
	Types["SQM_S25"] = reflect.TypeOf(SQM_S25{})
	Types["SQM_S25_GENERAL_RESOURCE"] = reflect.TypeOf(SQM_S25_GENERAL_RESOURCE{})
	Types["SQM_S25_LOCATION_RESOURCE"] = reflect.TypeOf(SQM_S25_LOCATION_RESOURCE{})
	Types["SQM_S25_PERSONNEL_RESOURCE"] = reflect.TypeOf(SQM_S25_PERSONNEL_RESOURCE{})
	Types["SQM_S25_REQUEST"] = reflect.TypeOf(SQM_S25_REQUEST{})
	Types["SQM_S25_RESOURCES"] = reflect.TypeOf(SQM_S25_RESOURCES{})
	Types["SQM_S25_SERVICE"] = reflect.TypeOf(SQM_S25_SERVICE{})
	Types["SQR_S25"] = reflect.TypeOf(SQR_S25{})
	Types["SQR_S25_GENERAL_RESOURCE"] = reflect.TypeOf(SQR_S25_GENERAL_RESOURCE{})
	Types["SQR_S25_LOCATION_RESOURCE"] = reflect.TypeOf(SQR_S25_LOCATION_RESOURCE{})
	Types["SQR_S25_PATIENT"] = reflect.TypeOf(SQR_S25_PATIENT{})
	Types["SQR_S25_PERSONNEL_RESOURCE"] = reflect.TypeOf(SQR_S25_PERSONNEL_RESOURCE{})
	Types["SQR_S25_RESOURCES"] = reflect.TypeOf(SQR_S25_RESOURCES{})
	Types["SQR_S25_SCHEDULE"] = reflect.TypeOf(SQR_S25_SCHEDULE{})
	Types["SQR_S25_SERVICE"] = reflect.TypeOf(SQR_S25_SERVICE{})
	Types["SRM_S01"] = reflect.TypeOf(SRM_S01{})
	Types["SRM_S01_GENERAL_RESOURCE"] = reflect.TypeOf(SRM_S01_GENERAL_RESOURCE{})
	Types["SRM_S01_LOCATION_RESOURCE"] = reflect.TypeOf(SRM_S01_LOCATION_RESOURCE{})
	Types["SRM_S01_PATIENT"] = reflect.TypeOf(SRM_S01_PATIENT{})
	Types["SRM_S01_PERSONNEL_RESOURCE"] = reflect.TypeOf(SRM_S01_PERSONNEL_RESOURCE{})
	Types["SRM_S01_RESOURCES"] = reflect.TypeOf(SRM_S01_RESOURCES{})
	Types["SRM_S01_SERVICE"] = reflect.TypeOf(SRM_S01_SERVICE{})
	Types["SRR_S01"] = reflect.TypeOf(SRR_S01{})
	Types["SRR_S01_GENERAL_RESOURCE"] = reflect.TypeOf(SRR_S01_GENERAL_RESOURCE{})
	Types["SRR_S01_LOCATION_RESOURCE"] = reflect.TypeOf(SRR_S01_LOCATION_RESOURCE{})
	Types["SRR_S01_PATIENT"] = reflect.TypeOf(SRR_S01_PATIENT{})
	Types["SRR_S01_PERSONNEL_RESOURCE"] = reflect.TypeOf(SRR_S01_PERSONNEL_RESOURCE{})
	Types["SRR_S01_RESOURCES"] = reflect.TypeOf(SRR_S01_RESOURCES{})
	Types["SRR_S01_SCHEDULE"] = reflect.TypeOf(SRR_S01_SCHEDULE{})
	Types["SRR_S01_SERVICE"] = reflect.TypeOf(SRR_S01_SERVICE{})
	Types["SSR_U04"] = reflect.TypeOf(SSR_U04{})
	Types["SSU_U03"] = reflect.TypeOf(SSU_U03{})
	Types["SSU_U03_SPECIMEN_CONTAINER"] = reflect.TypeOf(SSU_U03_SPECIMEN_CONTAINER{})
	Types["SUR_P09"] = reflect.TypeOf(SUR_P09{})
	Types["SUR_P09_FACILITY"] = reflect.TypeOf(SUR_P09_FACILITY{})
	Types["SUR_P09_FACILITY_DETAIL"] = reflect.TypeOf(SUR_P09_FACILITY_DETAIL{})
	Types["SUR_P09_PRODUCT"] = reflect.TypeOf(SUR_P09_PRODUCT{})
	Types["TBR_Q01"] = reflect.TypeOf(TBR_Q01{})
	Types["TBR_R08"] = reflect.TypeOf(TBR_R08{})
	Types["TCU_U10"] = reflect.TypeOf(TCU_U10{})
	Types["UDM_Q05"] = reflect.TypeOf(UDM_Q05{})
	Types["VQQ_Q01"] = reflect.TypeOf(VQQ_Q01{})
	Types["VQQ_Q07"] = reflect.TypeOf(VQQ_Q07{})
	Types["VXQ_V01"] = reflect.TypeOf(VXQ_V01{})
	Types["VXR_V03"] = reflect.TypeOf(VXR_V03{})
	Types["VXR_V03_INSURANCE"] = reflect.TypeOf(VXR_V03_INSURANCE{})
	Types["VXR_V03_OBSERVATION"] = reflect.TypeOf(VXR_V03_OBSERVATION{})
	Types["VXR_V03_ORDER"] = reflect.TypeOf(VXR_V03_ORDER{})
	Types["VXR_V03_PATIENT_VISIT"] = reflect.TypeOf(VXR_V03_PATIENT_VISIT{})
	Types["VXU_V04"] = reflect.TypeOf(VXU_V04{})
	Types["VXU_V04_INSURANCE"] = reflect.TypeOf(VXU_V04_INSURANCE{})
	Types["VXU_V04_OBSERVATION"] = reflect.TypeOf(VXU_V04_OBSERVATION{})
	Types["VXU_V04_ORDER"] = reflect.TypeOf(VXU_V04_ORDER{})
	Types["VXU_V04_PATIENT"] = reflect.TypeOf(VXU_V04_PATIENT{})
	Types["VXX_V02"] = reflect.TypeOf(VXX_V02{})
	Types["VXX_V02_PATIENT"] = reflect.TypeOf(VXX_V02_PATIENT{})
	Types["GenericHL7Segment"] = reflect.TypeOf(GenericHL7Segment{})
}
