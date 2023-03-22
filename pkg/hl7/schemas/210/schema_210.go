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

// This file contains the schemas for HL7 messages, segments and values for HL7v2 version 2.1.0.
// It has been auto-generated from the HL7v2 specification.

package hl7

import "reflect"

// CE represents the corresponding HL7 datatype.
// Definition from HL7 2.1
type CE struct {
	Identifier                  *ID `hl7:"false,Identifier"`
	Text                        *ST `hl7:"false,Text"`
	NameOfCodingSystem          *ST `hl7:"false,Name Of Coding System"`
	AlternateIdentifier         *ST `hl7:"false,Alternate Identifier"`
	AlternateText               *ST `hl7:"false,Alternate Text"`
	NameOfAlternateCodingSystem *ST `hl7:"false,Name Of Alternate Coding System"`
}

// ACC represents the corresponding HL7 segment.
// Definition from HL7 2.1
type ACC struct {
	ACCIDENTDATETIME *TS `hl7:"false,ACCIDENT DATE/TIME"` // ACC-1
	ACCIDENTCODE     *ID `hl7:"false,ACCIDENT CODE"`      // ACC-2
	ACCIDENTLOCATION *ST `hl7:"false,ACCIDENT LOCATION"`  // ACC-3
}

func (s *ACC) SegmentName() string {
	return "ACC"
}

// ADD represents the corresponding HL7 segment.
// Definition from HL7 2.1
type ADD struct {
	ADDENDUMCONTINUATIONPOINTER *ST `hl7:"false,ADDENDUM CONTINUATION POINTER"` // ADD-1
}

func (s *ADD) SegmentName() string {
	return "ADD"
}

// BHS represents the corresponding HL7 segment.
// Definition from HL7 2.1
type BHS struct {
	BATCHFIELDSEPARATOR       *ST `hl7:"true,BATCH FIELD SEPARATOR"`        // BHS-1
	BATCHENCODINGCHARACTERS   *ST `hl7:"true,BATCH ENCODING CHARACTERS"`    // BHS-2
	BATCHSENDINGAPPLICATION   *ST `hl7:"false,BATCH SENDING APPLICATION"`   // BHS-3
	BATCHSENDINGFACILITY      *ST `hl7:"false,BATCH SENDING FACILITY"`      // BHS-4
	BATCHRECEIVINGAPPLICATION *ST `hl7:"false,BATCH RECEIVING APPLICATION"` // BHS-5
	BATCHRECEIVINGFACILITY    *ST `hl7:"false,BATCH RECEIVING FACILITY"`    // BHS-6
	BATCHCREATIONDATETIME     *TS `hl7:"false,BATCH CREATION DATE/TIME"`    // BHS-7
	BATCHSECURITY             *ST `hl7:"false,BATCH SECURITY"`              // BHS-8
	BATCHNAMEIDTYPE           *ST `hl7:"false,BATCH NAME/ID/TYPE"`          // BHS-9
	BATCHCOMMENT              *ST `hl7:"false,BATCH COMMENT"`               // BHS-10
	BATCHCONTROLID            *ST `hl7:"false,BATCH CONTROL ID"`            // BHS-11
	REFERENCEBATCHCONTROLID   *ST `hl7:"false,REFERENCE BATCH CONTROL ID"`  // BHS-12
}

func (s *BHS) SegmentName() string {
	return "BHS"
}

// BLG represents the corresponding HL7 segment.
// Definition from HL7 2.1
type BLG struct {
	WHENTOCHARGE *CM `hl7:"false,WHEN TO CHARGE"` // BLG-1
	CHARGETYPE   *ID `hl7:"false,CHARGE TYPE"`    // BLG-2
	ACCOUNTID    *CM `hl7:"false,ACCOUNT ID"`     // BLG-3
}

func (s *BLG) SegmentName() string {
	return "BLG"
}

// BTS represents the corresponding HL7 segment.
// Definition from HL7 2.1
type BTS struct {
	BATCHMESSAGECOUNT *ST `hl7:"false,BATCH MESSAGE COUNT"` // BTS-1
	BATCHCOMMENT      *ST `hl7:"false,BATCH COMMENT"`       // BTS-2
	BATCHTOTALS       *CM `hl7:"false,BATCH TOTALS"`        // BTS-3
}

func (s *BTS) SegmentName() string {
	return "BTS"
}

// DG1 represents the corresponding HL7 segment.
// Definition from HL7 2.1
type DG1 struct {
	SETIDDIAGNOSIS          *SI `hl7:"true,SET ID - DIAGNOSIS"`         // DG1-1
	DIAGNOSISCODINGMETHOD   *ID `hl7:"true,DIAGNOSIS CODING METHOD"`    // DG1-2
	DIAGNOSISCODE           *ID `hl7:"false,DIAGNOSIS CODE"`            // DG1-3
	DIAGNOSISDESCRIPTION    *ST `hl7:"false,DIAGNOSIS DESCRIPTION"`     // DG1-4
	DIAGNOSISDATETIME       *TS `hl7:"false,DIAGNOSIS DATE/TIME"`       // DG1-5
	DIAGNOSISDRGTYPE        *ID `hl7:"true,DIAGNOSIS/DRG TYPE"`         // DG1-6
	MAJORDIAGNOSTICCATEGORY *ST `hl7:"false,MAJOR DIAGNOSTIC CATEGORY"` // DG1-7
	DIAGNOSTICRELATEDGROUP  *ID `hl7:"false,DIAGNOSTIC RELATED GROUP"`  // DG1-8
	DRGAPPROVALINDICATOR    *ID `hl7:"false,DRG APPROVAL INDICATOR"`    // DG1-9
	DRGGROUPERREVIEWCODE    *ID `hl7:"false,DRG GROUPER REVIEW CODE"`   // DG1-10
	OUTLIERTYPE             *ID `hl7:"false,OUTLIER TYPE"`              // DG1-11
	OUTLIERDAYS             *NM `hl7:"false,OUTLIER DAYS"`              // DG1-12
	OUTLIERCOST             *NM `hl7:"false,OUTLIER COST"`              // DG1-13
	GROUPERVERSIONANDTYPE   *ST `hl7:"false,GROUPER VERSION AND TYPE"`  // DG1-14
}

func (s *DG1) SegmentName() string {
	return "DG1"
}

// DSC represents the corresponding HL7 segment.
// Definition from HL7 2.1
type DSC struct {
	CONTINUATIONPOINTER *ST `hl7:"false,CONTINUATION POINTER"` // DSC-1
}

func (s *DSC) SegmentName() string {
	return "DSC"
}

// DSP represents the corresponding HL7 segment.
// Definition from HL7 2.1
type DSP struct {
	SETIDDISPLAYDATA  *SI `hl7:"false,SET ID - DISPLAY DATA"` // DSP-1
	DISPLAYLEVEL      *SI `hl7:"false,DISPLAY LEVEL"`         // DSP-2
	DATALINE          *TX `hl7:"true,DATA LINE"`              // DSP-3
	LOGICALBREAKPOINT *ST `hl7:"false,LOGICAL BREAK POINT"`   // DSP-4
	RESULTID          *TX `hl7:"false,RESULT ID"`             // DSP-5
}

func (s *DSP) SegmentName() string {
	return "DSP"
}

// ERR represents the corresponding HL7 segment.
// Definition from HL7 2.1
type ERR struct {
	ERRORCODEANDLOCATION []ID `hl7:"true,ERROR CODE AND LOCATION"` // ERR-1
}

func (s *ERR) SegmentName() string {
	return "ERR"
}

// EVN represents the corresponding HL7 segment.
// Definition from HL7 2.1
type EVN struct {
	EVENTTYPECODE        *ID `hl7:"true,EVENT TYPE CODE"`          // EVN-1
	DATETIMEOFEVENT      *TS `hl7:"true,DATE/TIME OF EVENT"`       // EVN-2
	DATETIMEPLANNEDEVENT *TS `hl7:"false,DATE/TIME PLANNED EVENT"` // EVN-3
	EVENTREASONCODE      *ID `hl7:"false,EVENT REASON CODE"`       // EVN-4
}

func (s *EVN) SegmentName() string {
	return "EVN"
}

// FHS represents the corresponding HL7 segment.
// Definition from HL7 2.1
type FHS struct {
	FILEFIELDSEPARATOR       *ST `hl7:"true,FILE FIELD SEPARATOR"`        // FHS-1
	FILEENCODINGCHARACTERS   *ST `hl7:"true,FILE ENCODING CHARACTERS"`    // FHS-2
	FILESENDINGAPPLICATION   *ST `hl7:"false,FILE SENDING APPLICATION"`   // FHS-3
	FILESENDINGFACILITY      *ST `hl7:"false,FILE SENDING FACILITY"`      // FHS-4
	FILERECEIVINGAPPLICATION *ST `hl7:"false,FILE RECEIVING APPLICATION"` // FHS-5
	FILERECEIVINGFACILITY    *ST `hl7:"false,FILE RECEIVING FACILITY"`    // FHS-6
	DATETIMEOFFILECREATION   *TS `hl7:"false,DATE/TIME OF FILE CREATION"` // FHS-7
	FILESECURITY             *ST `hl7:"false,FILE SECURITY"`              // FHS-8
	FILENAMEID               *ST `hl7:"false,FILE NAME/ID"`               // FHS-9
	FILEHEADERCOMMENT        *ST `hl7:"false,FILE HEADER COMMENT"`        // FHS-10
	FILECONTROLID            *ST `hl7:"false,FILE CONTROL ID"`            // FHS-11
	REFERENCEFILECONTROLID   *ST `hl7:"false,REFERENCE FILE CONTROL ID"`  // FHS-12
}

func (s *FHS) SegmentName() string {
	return "FHS"
}

// FT1 represents the corresponding HL7 segment.
// Definition from HL7 2.1
type FT1 struct {
	SETIDFINANCIALTRANSACTION *SI `hl7:"false,SET ID - FINANCIAL TRANSACTION"` // FT1-1
	TRANSACTIONID             *ST `hl7:"false,TRANSACTION ID"`                 // FT1-2
	TRANSACTIONBATCHID        *ST `hl7:"false,TRANSACTION BATCH ID"`           // FT1-3
	TRANSACTIONDATE           *DT `hl7:"true,TRANSACTION DATE"`                // FT1-4
	TRANSACTIONPOSTINGDATE    *DT `hl7:"false,TRANSACTION POSTING DATE"`       // FT1-5
	TRANSACTIONTYPE           *ID `hl7:"true,TRANSACTION TYPE"`                // FT1-6
	TRANSACTIONCODE           *ID `hl7:"true,TRANSACTION CODE"`                // FT1-7
	TRANSACTIONDESCRIPTION    *ST `hl7:"false,TRANSACTION DESCRIPTION"`        // FT1-8
	TRANSACTIONDESCRIPTIONALT *ST `hl7:"false,TRANSACTION DESCRIPTION - ALT"`  // FT1-9
	TRANSACTIONAMOUNTEXTENDED *NM `hl7:"false,TRANSACTION AMOUNT - EXTENDED"`  // FT1-10
	TRANSACTIONQUANTITY       *NM `hl7:"false,TRANSACTION QUANTITY"`           // FT1-11
	TRANSACTIONAMOUNTUNIT     *NM `hl7:"false,TRANSACTION AMOUNT - UNIT"`      // FT1-12
	DEPARTMENTCODE            *ST `hl7:"false,DEPARTMENT CODE"`                // FT1-13
	INSURANCEPLANID           *ID `hl7:"false,INSURANCE PLAN ID"`              // FT1-14
	INSURANCEAMOUNT           *NM `hl7:"false,INSURANCE AMOUNT"`               // FT1-15
	PATIENTLOCATION           *ST `hl7:"false,PATIENT LOCATION"`               // FT1-16
	FEESCHEDULE               *ID `hl7:"false,FEE SCHEDULE"`                   // FT1-17
	PATIENTTYPE               *ID `hl7:"false,PATIENT TYPE"`                   // FT1-18
	DIAGNOSISCODE             *ID `hl7:"false,DIAGNOSIS CODE"`                 // FT1-19
	PERFORMEDBYCODE           *CN `hl7:"false,PERFORMED BY CODE"`              // FT1-20
	ORDEREDBYCODE             *CN `hl7:"false,ORDERED BY CODE"`                // FT1-21
	UNITCOST                  *NM `hl7:"false,UNIT COST"`                      // FT1-22
}

func (s *FT1) SegmentName() string {
	return "FT1"
}

// FTS represents the corresponding HL7 segment.
// Definition from HL7 2.1
type FTS struct {
	FILEBATCHCOUNT     *ST `hl7:"false,FILE BATCH COUNT"`     // FTS-1
	FILETRAILERCOMMENT *ST `hl7:"false,FILE TRAILER COMMENT"` // FTS-2
}

func (s *FTS) SegmentName() string {
	return "FTS"
}

// GT1 represents the corresponding HL7 segment.
// Definition from HL7 2.1
type GT1 struct {
	SETIDGUARANTOR            *SI `hl7:"true,SET ID - GUARANTOR"`           // GT1-1
	GUARANTORNUMBER           *ID `hl7:"false,GUARANTOR NUMBER"`            // GT1-2
	GUARANTORNAME             *PN `hl7:"true,GUARANTOR NAME"`               // GT1-3
	GUARANTORSPOUSENAME       *PN `hl7:"false,GUARANTOR SPOUSE NAME"`       // GT1-4
	GUARANTORADDRESS          *AD `hl7:"false,GUARANTOR ADDRESS"`           // GT1-5
	GUARANTORPHNUMHOME        *TN `hl7:"false,GUARANTOR PH. NUM.- HOME"`    // GT1-6
	GUARANTORPHNUMBUSINESS    *TN `hl7:"false,GUARANTOR PH. NUM-BUSINESS"`  // GT1-7
	GUARANTORDATEOFBIRTH      *DT `hl7:"false,GUARANTOR DATE OF BIRTH"`     // GT1-8
	GUARANTORSEX              *ID `hl7:"false,GUARANTOR SEX"`               // GT1-9
	GUARANTORTYPE             *ID `hl7:"false,GUARANTOR TYPE"`              // GT1-10
	GUARANTORRELATIONSHIP     *ID `hl7:"false,GUARANTOR RELATIONSHIP"`      // GT1-11
	GUARANTORSSN              *ST `hl7:"false,GUARANTOR SSN"`               // GT1-12
	GUARANTORDATEBEGIN        *DT `hl7:"false,GUARANTOR DATE - BEGIN"`      // GT1-13
	GUARANTORDATEEND          *DT `hl7:"false,GUARANTOR DATE - END"`        // GT1-14
	GUARANTORPRIORITY         *NM `hl7:"false,GUARANTOR PRIORITY"`          // GT1-15
	GUARANTOREMPLOYERNAME     *ST `hl7:"false,GUARANTOR EMPLOYER NAME"`     // GT1-16
	GUARANTOREMPLOYERADDRESS  *AD `hl7:"false,GUARANTOR EMPLOYER ADDRESS"`  // GT1-17
	GUARANTOREMPLOYPHONE      *TN `hl7:"false,GUARANTOR EMPLOY PHONE #"`    // GT1-18
	GUARANTOREMPLOYEEIDNUM    *ST `hl7:"false,GUARANTOR EMPLOYEE ID NUM"`   // GT1-19
	GUARANTOREMPLOYMENTSTATUS *ID `hl7:"false,GUARANTOR EMPLOYMENT STATUS"` // GT1-20
}

func (s *GT1) SegmentName() string {
	return "GT1"
}

// IN1 represents the corresponding HL7 segment.
// Definition from HL7 2.1
type IN1 struct {
	SETIDINSURANCE                *SI `hl7:"true,SET ID - INSURANCE"`                 // IN1-1
	INSURANCEPLANID               *ID `hl7:"true,INSURANCE PLAN ID"`                  // IN1-2
	INSURANCECOMPANYID            *ST `hl7:"true,INSURANCE COMPANY ID"`               // IN1-3
	INSURANCECOMPANYNAME          *ST `hl7:"false,INSURANCE COMPANY NAME"`            // IN1-4
	INSURANCECOMPANYADDRESS       *AD `hl7:"false,INSURANCE COMPANY ADDRESS"`         // IN1-5
	INSURANCECOCONTACTPERS        *PN `hl7:"false,INSURANCE CO. CONTACT PERS"`        // IN1-6
	INSURANCECOPHONENUMBER        *TN `hl7:"false,INSURANCE CO PHONE NUMBER"`         // IN1-7
	GROUPNUMBER                   *ST `hl7:"false,GROUP NUMBER"`                      // IN1-8
	GROUPNAME                     *ST `hl7:"false,GROUP NAME"`                        // IN1-9
	INSUREDSGROUPEMPID            *ST `hl7:"false,INSURED'S GROUP EMP. ID"`           // IN1-10
	INSUREDSGROUPEMPNAME          *ST `hl7:"false,INSURED'S GROUP EMP. NAME"`         // IN1-11
	PLANEFFECTIVEDATE             *DT `hl7:"false,PLAN EFFECTIVE DATE"`               // IN1-12
	PLANEXPIRATIONDATE            *DT `hl7:"false,PLAN EXPIRATION DATE"`              // IN1-13
	AUTHORIZATIONINFORMATION      *ST `hl7:"false,AUTHORIZATION INFORMATION"`         // IN1-14
	PLANTYPE                      *ID `hl7:"false,PLAN TYPE"`                         // IN1-15
	NAMEOFINSURED                 *PN `hl7:"false,NAME OF INSURED"`                   // IN1-16
	INSUREDSRELATIONSHIPTOPATIENT *ID `hl7:"false,INSURED'S RELATIONSHIP TO PATIENT"` // IN1-17
	INSUREDSDATEOFBIRTH           *DT `hl7:"false,INSURED'S DATE OF BIRTH"`           // IN1-18
	INSUREDSADDRESS               *AD `hl7:"false,INSURED'S ADDRESS"`                 // IN1-19
	ASSIGNMENTOFBENEFITS          *ID `hl7:"false,ASSIGNMENT OF BENEFITS"`            // IN1-20
	COORDINATIONOFBENEFITS        *ID `hl7:"false,COORDINATION OF BENEFITS"`          // IN1-21
	COORDOFBENPRIORITY            *ST `hl7:"false,COORD OF BEN. PRIORITY"`            // IN1-22
	NOTICEOFADMISSIONCODE         *ID `hl7:"false,NOTICE OF ADMISSION CODE"`          // IN1-23
	NOTICEOFADMISSIONDATE         *DT `hl7:"false,NOTICE OF ADMISSION DATE"`          // IN1-24
	RPTOFELIGIBILITYCODE          *ID `hl7:"false,RPT OF ELIGIBILITY CODE"`           // IN1-25
	RPTOFELIGIBILITYDATE          *DT `hl7:"false,RPT OF ELIGIBILITY DATE"`           // IN1-26
	RELEASEINFORMATIONCODE        *ID `hl7:"false,RELEASE INFORMATION CODE"`          // IN1-27
	PREADMITCERT                  *ST `hl7:"false,PRE-ADMIT CERT."`                   // IN1-28
	VERIFICATIONDATE              *DT `hl7:"false,VERIFICATION DATE"`                 // IN1-29
	VERIFICATIONBY                *CM `hl7:"false,VERIFICATION BY"`                   // IN1-30
	TYPEOFAGREEMENTCODE           *ID `hl7:"false,TYPE OF AGREEMENT CODE"`            // IN1-31
	BILLINGSTATUS                 *ID `hl7:"false,BILLING STATUS"`                    // IN1-32
	LIFETIMERESERVEDAYS           *NM `hl7:"false,LIFETIME RESERVE DAYS"`             // IN1-33
	DELAYBEFORELRDAY              *NM `hl7:"false,DELAY BEFORE L. R. DAY"`            // IN1-34
	COMPANYPLANCODE               *ST `hl7:"false,COMPANY PLAN CODE"`                 // IN1-35
	POLICYNUMBER                  *ST `hl7:"false,POLICY NUMBER"`                     // IN1-36
	POLICYDEDUCTIBLE              *NM `hl7:"false,POLICY DEDUCTIBLE"`                 // IN1-37
	POLICYLIMITAMOUNT             *NM `hl7:"false,POLICY LIMIT - AMOUNT"`             // IN1-38
	POLICYLIMITDAYS               *NM `hl7:"false,POLICY LIMIT - DAYS"`               // IN1-39
	ROOMRATESEMIPRIVATE           *NM `hl7:"false,ROOM RATE - SEMI-PRIVATE"`          // IN1-40
	ROOMRATEPRIVATE               *NM `hl7:"false,ROOM RATE - PRIVATE"`               // IN1-41
	INSUREDSEMPLOYMENTSTATUS      *ID `hl7:"false,INSURED'S EMPLOYMENT STATUS"`       // IN1-42
	INSUREDSSEX                   *ID `hl7:"false,INSURED'S SEX"`                     // IN1-43
	INSUREDSEMPLOYERADDRESS       *AD `hl7:"false,INSURED'S EMPLOYER ADDRESS"`        // IN1-44
}

func (s *IN1) SegmentName() string {
	return "IN1"
}

// MRG represents the corresponding HL7 segment.
// Definition from HL7 2.1
type MRG struct {
	PRIORPATIENTIDINTERNAL    *CK `hl7:"true,PRIOR PATIENT ID - INTERNAL"`   // MRG-1
	PRIORALTERNATEPATIENTID   *CK `hl7:"false,PRIOR ALTERNATE PATIENT ID"`   // MRG-2
	PRIORPATIENTACCOUNTNUMBER *CK `hl7:"false,PRIOR PATIENT ACCOUNT NUMBER"` // MRG-3
}

func (s *MRG) SegmentName() string {
	return "MRG"
}

// MSA represents the corresponding HL7 segment.
// Definition from HL7 2.1
type MSA struct {
	ACKNOWLEDGMENTCODE        *ID `hl7:"true,ACKNOWLEDGMENT CODE"`          // MSA-1
	MESSAGECONTROLID          *ST `hl7:"true,MESSAGE CONTROL ID"`           // MSA-2
	TEXTMESSAGE               *ST `hl7:"false,TEXT MESSAGE"`                // MSA-3
	EXPECTEDSEQUENCENUMBER    *NM `hl7:"false,EXPECTED SEQUENCE NUMBER"`    // MSA-4
	DELAYEDACKNOWLEDGMENTTYPE *ID `hl7:"false,DELAYED ACKNOWLEDGMENT TYPE"` // MSA-5
}

func (s *MSA) SegmentName() string {
	return "MSA"
}

// MSH represents the corresponding HL7 segment.
// Definition from HL7 2.1
type MSH struct {
	// Missing: MSH.1
	ENCODINGCHARACTERS   *Delimiters `hl7:"true,ENCODING CHARACTERS"`    // MSH-2
	SENDINGAPPLICATION   *ST         `hl7:"false,SENDING APPLICATION"`   // MSH-3
	SENDINGFACILITY      *ST         `hl7:"false,SENDING FACILITY"`      // MSH-4
	RECEIVINGAPPLICATION *ST         `hl7:"false,RECEIVING APPLICATION"` // MSH-5
	RECEIVINGFACILITY    *ST         `hl7:"false,RECEIVING FACILITY"`    // MSH-6
	DATETIMEOFMESSAGE    *TS         `hl7:"false,DATE/TIME OF MESSAGE"`  // MSH-7
	Security             *ST         `hl7:"false,Security"`              // MSH-8
	MESSAGETYPE          *ID         `hl7:"true,MESSAGE TYPE"`           // MSH-9
	MESSAGECONTROLID     *ST         `hl7:"true,MESSAGE CONTROL ID"`     // MSH-10
	PROCESSINGID         *ID         `hl7:"true,PROCESSING ID"`          // MSH-11
	VERSIONID            *NM         `hl7:"true,VERSION ID"`             // MSH-12
	SEQUENCENUMBER       *NM         `hl7:"false,SEQUENCE NUMBER"`       // MSH-13
	CONTINUATIONPOINTER  *ST         `hl7:"false,CONTINUATION POINTER"`  // MSH-14
}

func (s *MSH) SegmentName() string {
	return "MSH"
}

// NCK represents the corresponding HL7 segment.
// Definition from HL7 2.1
type NCK struct {
	SYSTEMDATETIME *TS `hl7:"true,SYSTEM DATE/TIME"` // NCK-1
}

func (s *NCK) SegmentName() string {
	return "NCK"
}

// NK1 represents the corresponding HL7 segment.
// Definition from HL7 2.1
type NK1 struct {
	SETIDNEXTOFKIN        *SI  `hl7:"true,SET ID - NEXT OF KIN"`        // NK1-1
	NEXTOFKINNAME         *PN  `hl7:"false,NEXT OF KIN NAME"`           // NK1-2
	NEXTOFKINRELATIONSHIP *ST  `hl7:"false,NEXT OF KIN RELATIONSHIP"`   // NK1-3
	NEXTOFKINADDRESS      *AD  `hl7:"false,NEXT OF KIN - ADDRESS"`      // NK1-4
	NEXTOFKINPHONENUMBER  []TN `hl7:"false,NEXT OF KIN - PHONE NUMBER"` // NK1-5
}

func (s *NK1) SegmentName() string {
	return "NK1"
}

// NPU represents the corresponding HL7 segment.
// Definition from HL7 2.1
type NPU struct {
	BEDLOCATION *ID `hl7:"true,BED LOCATION"` // NPU-1
	BEDSTATUS   *ID `hl7:"false,BED STATUS"`  // NPU-2
}

func (s *NPU) SegmentName() string {
	return "NPU"
}

// NSC represents the corresponding HL7 segment.
// Definition from HL7 2.1
type NSC struct {
	NETWORKCHANGETYPE  *ID `hl7:"true,NETWORK CHANGE TYPE"`  // NSC-1
	CURRENTCPU         *ST `hl7:"false,CURRENT CPU"`         // NSC-2
	CURRENTFILESERVER  *ST `hl7:"false,CURRENT FILESERVER"`  // NSC-3
	CURRENTAPPLICATION *ST `hl7:"false,CURRENT APPLICATION"` // NSC-4
	CURRENTFACILITY    *ST `hl7:"false,CURRENT FACILITY"`    // NSC-5
	NEWCPU             *ST `hl7:"false,NEW CPU"`             // NSC-6
	NEWFILESERVER      *ST `hl7:"false,NEW FILESERVER"`      // NSC-7
	NEWAPPLICATION     *ST `hl7:"false,NEW APPLICATION"`     // NSC-8
	NEWFACILITY        *ST `hl7:"false,NEW FACILITY"`        // NSC-9
}

func (s *NSC) SegmentName() string {
	return "NSC"
}

// NST represents the corresponding HL7 segment.
// Definition from HL7 2.1
type NST struct {
	STATISTICSAVAILABLE    *ID `hl7:"true,STATISTICS AVAILABLE"`      // NST-1
	SOURCEIDENTIFIER       *ST `hl7:"false,SOURCE IDENTIFIER"`        // NST-2
	SOURCETYPE             *ID `hl7:"false,SOURCE TYPE"`              // NST-3
	STATISTICSSTART        *TS `hl7:"false,STATISTICS START"`         // NST-4
	STATISTICSEND          *TS `hl7:"false,STATISTICS END"`           // NST-5
	RECEIVECHARACTERCOUNT  *NM `hl7:"false,RECEIVE CHARACTER COUNT"`  // NST-6
	SENDCHARACTERCOUNT     *NM `hl7:"false,SEND CHARACTER COUNT"`     // NST-7
	MESSAGESRECEIVED       *NM `hl7:"false,MESSAGES RECEIVED"`        // NST-8
	MESSAGESSENT           *NM `hl7:"false,MESSAGES SENT"`            // NST-9
	CHECKSUMERRORSRECEIVED *NM `hl7:"false,CHECKSUM ERRORS RECEIVED"` // NST-10
	LENGTHERRORSRECEIVED   *NM `hl7:"false,LENGTH ERRORS RECEIVED"`   // NST-11
	OTHERERRORSRECEIVED    *NM `hl7:"false,OTHER ERRORS RECEIVED"`    // NST-12
	CONNECTTIMEOUTS        *NM `hl7:"false,CONNECT TIMEOUTS"`         // NST-13
	RECEIVETIMEOUTS        *NM `hl7:"false,RECEIVE TIMEOUTS"`         // NST-14
	NETWORKERRORS          *NM `hl7:"false,NETWORK ERRORS"`           // NST-15
}

func (s *NST) SegmentName() string {
	return "NST"
}

// NTE represents the corresponding HL7 segment.
// Definition from HL7 2.1
type NTE struct {
	SETIDNOTESANDCOMMENTS *SI  `hl7:"false,SET ID - NOTES AND COMMENTS"` // NTE-1
	SOURCEOFCOMMENT       *ID  `hl7:"false,SOURCE OF COMMENT"`           // NTE-2
	COMMENT               []TX `hl7:"true,COMMENT"`                      // NTE-3
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

// ORC represents the corresponding HL7 segment.
// Definition from HL7 2.1
type ORC struct {
	ORDERCONTROL          *ST  `hl7:"true,ORDER CONTROL"`             // ORC-1
	PLACERORDER           *CM  `hl7:"false,PLACER ORDER #"`           // ORC-2
	FILLERORDER           *CM  `hl7:"false,FILLER ORDER #"`           // ORC-3
	PLACERGROUP           *CM  `hl7:"false,PLACER GROUP #"`           // ORC-4
	ORDERSTATUS           *ST  `hl7:"false,ORDER STATUS"`             // ORC-5
	RESPONSEFLAG          *ST  `hl7:"false,RESPONSE FLAG"`            // ORC-6
	TIMINGQUANTITY        *CM  `hl7:"false,TIMING/QUANTITY"`          // ORC-7
	PARENT                *CM  `hl7:"false,PARENT"`                   // ORC-8
	DATETIMEOFTRANSACTION *TS  `hl7:"false,DATE/TIME OF TRANSACTION"` // ORC-9
	ENTEREDBY             *CN  `hl7:"false,ENTERED BY"`               // ORC-10
	VERIFIEDBY            *CN  `hl7:"false,VERIFIED BY"`              // ORC-11
	ORDERINGPROVIDER      *CN  `hl7:"false,ORDERING PROVIDER"`        // ORC-12
	ENTERERSLOCATION      *CM  `hl7:"false,ENTERER'S LOCATION"`       // ORC-13
	CALLBACKPHONENUMBER   []TN `hl7:"false,CALL BACK PHONE NUMBER"`   // ORC-14
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
// Definition from HL7 2.1
type PID struct {
	SETIDPATIENTID              *SI  `hl7:"false,SET ID - PATIENT ID"`               // PID-1
	PATIENTIDEXTERNALEXTERNALID *CK  `hl7:"false,PATIENT ID EXTERNAL (EXTERNAL ID)"` // PID-2
	PATIENTIDINTERNALINTERNALID *CK  `hl7:"true,PATIENT ID INTERNAL (INTERNAL ID)"`  // PID-3
	ALTERNATEPATIENTID          *ST  `hl7:"false,ALTERNATE PATIENT ID"`              // PID-4
	PATIENTNAME                 *PN  `hl7:"true,PATIENT NAME"`                       // PID-5
	MOTHERSMAIDENNAME           *ST  `hl7:"false,MOTHER'S MAIDEN NAME"`              // PID-6
	DATEOFBIRTH                 *DT  `hl7:"false,DATE OF BIRTH"`                     // PID-7
	SEX                         *ID  `hl7:"false,SEX"`                               // PID-8
	PATIENTALIAS                []PN `hl7:"false,PATIENT ALIAS"`                     // PID-9
	Race                        *CWE `hl7:"false,Race"`                              // PID-10
	PATIENTADDRESS              *AD  `hl7:"false,PATIENT ADDRESS"`                   // PID-11
	COUNTYCODE                  *ID  `hl7:"false,COUNTY CODE"`                       // PID-12
	PHONENUMBERHOME             []TN `hl7:"false,PHONE NUMBER - HOME"`               // PID-13
	PHONENUMBERBUSINESS         []TN `hl7:"false,PHONE NUMBER - BUSINESS"`           // PID-14
	LANGUAGEPATIENT             *ST  `hl7:"false,LANGUAGE - PATIENT"`                // PID-15
	MARITALSTATUS               *ID  `hl7:"false,MARITAL STATUS"`                    // PID-16
	RELIGION                    *ID  `hl7:"false,RELIGION"`                          // PID-17
	PATIENTACCOUNTNUMBER        *CK  `hl7:"false,PATIENT ACCOUNT NUMBER"`            // PID-18
	SSNNUMBERPATIENT            *ST  `hl7:"false,SSN NUMBER - PATIENT"`              // PID-19
	DRIVERSLICNUMPATIENT        *CM  `hl7:"false,DRIVER'S LIC NUM - PATIENT"`        // PID-20
}

func (s *PID) SegmentName() string {
	return "PID"
}

// PR1 represents the corresponding HL7 segment.
// Definition from HL7 2.1
type PR1 struct {
	SETIDPROCEDURE        []SI `hl7:"true,SET ID - PROCEDURE"`       // PR1-1
	PROCEDURECODINGMETHOD *ID  `hl7:"true,PROCEDURE CODING METHOD."` // PR1-2
	PROCEDURECODE         *ID  `hl7:"true,PROCEDURE CODE"`           // PR1-3
	PROCEDUREDESCRIPTION  *ST  `hl7:"false,PROCEDURE DESCRIPTION"`   // PR1-4
	PROCEDUREDATETIME     *TS  `hl7:"true,PROCEDURE DATE/TIME"`      // PR1-5
	PROCEDURETYPE         *ID  `hl7:"true,PROCEDURE TYPE"`           // PR1-6
	PROCEDUREMINUTES      *NM  `hl7:"false,PROCEDURE MINUTES"`       // PR1-7
	ANESTHESIOLOGIST      *CN  `hl7:"false,ANESTHESIOLOGIST"`        // PR1-8
	ANESTHESIACODE        *ID  `hl7:"false,ANESTHESIA CODE"`         // PR1-9
	ANESTHESIAMINUTES     *NM  `hl7:"false,ANESTHESIA MINUTES"`      // PR1-10
	SURGEON               *CN  `hl7:"false,SURGEON"`                 // PR1-11
	RESIDENTCODE          *CN  `hl7:"false,RESIDENT CODE"`           // PR1-12
	CONSENTCODE           *ID  `hl7:"false,CONSENT CODE"`            // PR1-13
}

func (s *PR1) SegmentName() string {
	return "PR1"
}

// PV1 represents the corresponding HL7 segment.
// Definition from HL7 2.1
type PV1 struct {
	SETIDPATIENTVISIT       *SI  `hl7:"false,SET ID - PATIENT VISIT"`    // PV1-1
	PATIENTCLASS            *ID  `hl7:"true,PATIENT CLASS"`              // PV1-2
	ASSIGNEDPATIENTLOCATION *ID  `hl7:"true,ASSIGNED PATIENT LOCATION"`  // PV1-3
	ADMISSIONTYPE           *ID  `hl7:"false,ADMISSION TYPE"`            // PV1-4
	PREADMITNUMBER          *ST  `hl7:"false,PRE-ADMIT NUMBER"`          // PV1-5
	PRIORPATIENTLOCATION    *ID  `hl7:"false,PRIOR PATIENT LOCATION"`    // PV1-6
	ATTENDINGDOCTOR         *CN  `hl7:"false,ATTENDING DOCTOR"`          // PV1-7
	REFERRINGDOCTOR         *CN  `hl7:"false,REFERRING DOCTOR"`          // PV1-8
	CONSULTINGDOCTOR        []CN `hl7:"false,CONSULTING DOCTOR"`         // PV1-9
	HOSPITALSERVICE         *ID  `hl7:"false,HOSPITAL SERVICE"`          // PV1-10
	TEMPORARYLOCATION       *ID  `hl7:"false,TEMPORARY LOCATION"`        // PV1-11
	PREADMITTESTINDICATOR   *ID  `hl7:"false,PRE-ADMIT TEST INDICATOR"`  // PV1-12
	READMISSIONINDICATOR    *ID  `hl7:"false,RE-ADMISSION INDICATOR"`    // PV1-13
	ADMITSOURCE             *ID  `hl7:"false,ADMIT SOURCE"`              // PV1-14
	AMBULATORYSTATUS        *ID  `hl7:"false,AMBULATORY STATUS"`         // PV1-15
	VIPINDICATOR            *ID  `hl7:"false,VIP INDICATOR"`             // PV1-16
	ADMITTINGDOCTOR         *CN  `hl7:"false,ADMITTING DOCTOR"`          // PV1-17
	PATIENTTYPE             *ID  `hl7:"false,PATIENT TYPE"`              // PV1-18
	VISITNUMBER             *NM  `hl7:"false,VISIT NUMBER"`              // PV1-19
	FINANCIALCLASS          []ID `hl7:"false,FINANCIAL CLASS"`           // PV1-20
	CHARGEPRICEINDICATOR    *ID  `hl7:"false,CHARGE PRICE INDICATOR"`    // PV1-21
	COURTESYCODE            *ID  `hl7:"false,COURTESY CODE"`             // PV1-22
	CREDITRATING            *ID  `hl7:"false,CREDIT RATING"`             // PV1-23
	CONTRACTCODE            []ID `hl7:"false,CONTRACT CODE"`             // PV1-24
	CONTRACTEFFECTIVEDATE   []DT `hl7:"false,CONTRACT EFFECTIVE DATE"`   // PV1-25
	CONTRACTAMOUNT          []NM `hl7:"false,CONTRACT AMOUNT"`           // PV1-26
	CONTRACTPERIOD          []NM `hl7:"false,CONTRACT PERIOD"`           // PV1-27
	INTERESTCODE            *ID  `hl7:"false,INTEREST CODE"`             // PV1-28
	TRANSFERTOBADDEBTCODE   *ID  `hl7:"false,TRANSFER TO BAD DEBT CODE"` // PV1-29
	TRANSFERTOBADDEBTDATE   *DT  `hl7:"false,TRANSFER TO BAD DEBT DATE"` // PV1-30
	BADDEBTAGENCYCODE       *ST  `hl7:"false,BAD DEBT AGENCY CODE"`      // PV1-31
	BADDEBTTRANSFERAMOUNT   *NM  `hl7:"false,BAD DEBT TRANSFER AMOUNT"`  // PV1-32
	BADDEBTRECOVERYAMOUNT   *NM  `hl7:"false,BAD DEBT RECOVERY AMOUNT"`  // PV1-33
	DELETEACCOUNTINDICATOR  *ID  `hl7:"false,DELETE ACCOUNT INDICATOR"`  // PV1-34
	DELETEACCOUNTDATE       *DT  `hl7:"false,DELETE ACCOUNT DATE"`       // PV1-35
	DISCHARGEDISPOSITION    *ID  `hl7:"false,DISCHARGE DISPOSITION"`     // PV1-36
	DISCHARGEDTOLOCATION    *ID  `hl7:"false,DISCHARGED TO LOCATION"`    // PV1-37
	DIETTYPE                *ID  `hl7:"false,DIET TYPE"`                 // PV1-38
	SERVICINGFACILITY       *ID  `hl7:"false,SERVICING FACILITY"`        // PV1-39
	BEDSTATUS               *ID  `hl7:"false,BED STATUS"`                // PV1-40
	ACCOUNTSTATUS           *ID  `hl7:"false,ACCOUNT STATUS"`            // PV1-41
	PENDINGLOCATION         *ID  `hl7:"false,PENDING LOCATION"`          // PV1-42
	PRIORTEMPORARYLOCATION  *ID  `hl7:"false,PRIOR TEMPORARY LOCATION"`  // PV1-43
	ADMITDATETIME           *TS  `hl7:"false,ADMIT DATE/TIME"`           // PV1-44
	DISCHARGEDATETIME       *TS  `hl7:"false,DISCHARGE DATE/TIME"`       // PV1-45
	CURRENTPATIENTBALANCE   *NM  `hl7:"false,CURRENT PATIENT BALANCE"`   // PV1-46
	TOTALCHARGES            *NM  `hl7:"false,TOTAL CHARGES"`             // PV1-47
	TOTALADJUSTMENTS        *NM  `hl7:"false,TOTAL ADJUSTMENTS"`         // PV1-48
	TOTALPAYMENTS           *NM  `hl7:"false,TOTAL PAYMENTS"`            // PV1-49
}

func (s *PV1) SegmentName() string {
	return "PV1"
}

// QRD represents the corresponding HL7 segment.
// Definition from HL7 2.1
type QRD struct {
	QUERYDATETIME            *TS  `hl7:"true,QUERY DATE/TIME"`              // QRD-1
	QUERYFORMATCODE          *ID  `hl7:"true,QUERY FORMAT CODE"`            // QRD-2
	QUERYPRIORITY            *ID  `hl7:"true,QUERY PRIORITY"`               // QRD-3
	QUERYID                  *ST  `hl7:"true,QUERY ID"`                     // QRD-4
	DEFERREDRESPONSETYPE     *ID  `hl7:"false,DEFERRED RESPONSE TYPE"`      // QRD-5
	DEFERREDRESPONSEDATETIME *TS  `hl7:"false,DEFERRED RESPONSE DATE/TIME"` // QRD-6
	QUANTITYLIMITEDREQUEST   *CQ  `hl7:"true,QUANTITY LIMITED REQUEST"`     // QRD-7
	WHOSUBJECTFILTER         []ST `hl7:"true,WHO SUBJECT FILTER"`           // QRD-8
	WHATSUBJECTFILTER        []ID `hl7:"true,WHAT SUBJECT FILTER"`          // QRD-9
	WHATDEPARTMENTDATACODE   []ST `hl7:"true,WHAT DEPARTMENT DATA CODE"`    // QRD-10
	WHATDATACODEVALUEQUAL    []ST `hl7:"false,WHAT DATA CODE VALUE QUAL."`  // QRD-11
	QUERYRESULTSLEVEL        *ID  `hl7:"false,QUERY RESULTS LEVEL"`         // QRD-12
}

func (s *QRD) SegmentName() string {
	return "QRD"
}

// QRF represents the corresponding HL7 segment.
// Definition from HL7 2.1
type QRF struct {
	WHERESUBJECTFILTER    []ST `hl7:"true,WHERE SUBJECT FILTER"`       // QRF-1
	WHENDATASTARTDATETIME *TS  `hl7:"false,WHEN DATA START DATE/TIME"` // QRF-2
	WHENDATAENDDATETIME   *TS  `hl7:"false,WHEN DATA END DATE/TIME"`   // QRF-3
	WHATUSERQUALIFIER     []ST `hl7:"false,WHAT USER QUALIFIER"`       // QRF-4
	OTHERQRYSUBJECTFILTER []ST `hl7:"false,OTHER QRY SUBJECT FILTER"`  // QRF-5
}

func (s *QRF) SegmentName() string {
	return "QRF"
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

// UB1 represents the corresponding HL7 segment.
// Definition from HL7 2.1
type UB1 struct {
	SETIDUB82                 *SI  `hl7:"false,SET ID - UB82"`                  // UB1-1
	BLOODDEDUCTIBLE           *ST  `hl7:"false,BLOOD DEDUCTIBLE"`               // UB1-2
	BLOODFURNPINTSOF40        *ST  `hl7:"false,BLOOD FURN.-PINTS OF (40)"`      // UB1-3
	BLOODREPLACEDPINTS41      *ST  `hl7:"false,BLOOD REPLACED-PINTS (41)"`      // UB1-4
	BLOODNOTRPLCDPINTS42      *ST  `hl7:"false,BLOOD NOT RPLCD-PINTS(42)"`      // UB1-5
	COINSURANCEDAYS25         *ST  `hl7:"false,CO-INSURANCE DAYS (25)"`         // UB1-6
	CONDITIONCODE             []ID `hl7:"false,CONDITION CODE"`                 // UB1-7
	COVEREDDAYS23             *ST  `hl7:"false,COVERED DAYS - (23)"`            // UB1-8
	NONCOVEREDDAYS24          *ST  `hl7:"false,NON COVERED DAYS - (24)"`        // UB1-9
	VALUEAMOUNTCODE           []CM `hl7:"false,VALUE AMOUNT & CODE"`            // UB1-10
	NUMBEROFGRACEDAYS90       *ST  `hl7:"false,NUMBER OF GRACE DAYS (90)"`      // UB1-11
	SPECPROGINDICATOR44       *ID  `hl7:"false,SPEC. PROG. INDICATOR(44)"`      // UB1-12
	PSROURAPPROVALIND87       *ID  `hl7:"false,PSRO/UR APPROVAL IND. (87)"`     // UB1-13
	PSROURAPRVDSTAYFM88       *DT  `hl7:"false,PSRO/UR APRVD STAY-FM(88)"`      // UB1-14
	PSROURAPRVDSTAYTO89       *DT  `hl7:"false,PSRO/UR APRVD STAY-TO(89)"`      // UB1-15
	OCCURRENCE2832            []ID `hl7:"false,OCCURRENCE (28-32)"`             // UB1-16
	OCCURRENCESPAN33          *ID  `hl7:"false,OCCURRENCE SPAN (33)"`           // UB1-17
	OCCURRENCESPANSTARTDATE33 *DT  `hl7:"false,OCCURRENCE SPAN START DATE(33)"` // UB1-18
	OCCURSPANENDDATE33        *DT  `hl7:"false,OCCUR. SPAN END DATE (33)"`      // UB1-19
	UB82LOCATOR2              *ST  `hl7:"false,UB-82 LOCATOR 2"`                // UB1-20
	UB82LOCATOR9              *ST  `hl7:"false,UB-82 LOCATOR 9"`                // UB1-21
	UB82LOCATOR27             *ST  `hl7:"false,UB-82 LOCATOR 27"`               // UB1-22
	UB82LOCATOR45             *ST  `hl7:"false,UB-82 LOCATOR 45"`               // UB1-23
}

func (s *UB1) SegmentName() string {
	return "UB1"
}

// URD represents the corresponding HL7 segment.
// Definition from HL7 2.1
type URD struct {
	RUDATETIME              *TS  `hl7:"false,R/U DATE/TIME"`               // URD-1
	REPORTPRIORITY          *ID  `hl7:"false,REPORT PRIORITY"`             // URD-2
	RUWHOSUBJECTDEFINITION  []ST `hl7:"true,R/U WHO SUBJECT DEFINITION"`   // URD-3
	RUWHATSUBJECTDEFINITION []ID `hl7:"false,R/U WHAT SUBJECT DEFINITION"` // URD-4
	RUWHATDEPARTMENTCODE    []ST `hl7:"false,R/U WHAT DEPARTMENT CODE"`    // URD-5
	RUDISPLAYPRINTLOCATIONS []ST `hl7:"false,R/U DISPLAY/PRINT LOCATIONS"` // URD-6
	RURESULTSLEVEL          *ID  `hl7:"false,R/U RESULTS LEVEL"`           // URD-7
}

func (s *URD) SegmentName() string {
	return "URD"
}

// URS represents the corresponding HL7 segment.
// Definition from HL7 2.1
type URS struct {
	RUWHERESUBJECTDEFINITION    []ST `hl7:"true,R/U WHERE SUBJECT DEFINITION"`      // URS-1
	RUWHENDATASTARTDATETIME     *TS  `hl7:"false,R/U WHEN DATA START DATE/TIME"`    // URS-2
	RUWHENDATAENDDATETIME       *TS  `hl7:"false,R/U WHEN DATA END DATE/TIME"`      // URS-3
	RUWHATUSERQUALIFIER         []ST `hl7:"false,R/U WHAT USER QUALIFIER"`          // URS-4
	RUOTHERRESULTSSUBJECTDEFINI []ST `hl7:"false,R/U OTHER RESULTS SUBJECT DEFINI"` // URS-5
}

func (s *URS) SegmentName() string {
	return "URS"
}

// ACK represents the corresponding HL7 message type.
// Definition from HL7 2.1
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
// Definition from HL7 2.1
type ADR_A19 struct {
	MSH            *MSH                     `hl7:"true,MSH"`
	MSA            *MSA                     `hl7:"true,MSA"`
	QRD            *QRD                     `hl7:"true,QRD"`
	QUERY_RESPONSE []ADR_A19_QUERY_RESPONSE `hl7:"true,QUERY_RESPONSE"`
	DSC            *DSC                     `hl7:"false,DSC"`
	Other          []interface{}
}

func (s *ADR_A19) MessageTypeName() string {
	return "ADR_A19"
}

// ADR_A19_QUERY_RESPONSE represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADR_A19_QUERY_RESPONSE struct {
	EVN   *EVN `hl7:"false,EVN"`
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"true,PV1"`
	Other []interface{}
}

func (s *ADR_A19_QUERY_RESPONSE) MessageTypeSubStructName() string {
	return "ADR_A19_QUERY_RESPONSE"
}

// ADT_A01 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A01 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	NK1   *NK1 `hl7:"true,NK1"`
	PV1   *PV1 `hl7:"true,PV1"`
	DG1   *DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A01) MessageTypeName() string {
	return "ADT_A01"
}

// ADT_A02 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A02 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"true,PV1"`
	Other []interface{}
}

func (s *ADT_A02) MessageTypeName() string {
	return "ADT_A02"
}

// ADT_A03 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A03 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"true,PV1"`
	Other []interface{}
}

func (s *ADT_A03) MessageTypeName() string {
	return "ADT_A03"
}

// ADT_A04 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A04 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	NK1   *NK1 `hl7:"true,NK1"`
	PV1   *PV1 `hl7:"true,PV1"`
	DG1   *DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A04) MessageTypeName() string {
	return "ADT_A04"
}

// ADT_A05 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A05 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	NK1   *NK1 `hl7:"true,NK1"`
	PV1   *PV1 `hl7:"true,PV1"`
	DG1   *DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A05) MessageTypeName() string {
	return "ADT_A05"
}

// ADT_A06 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A06 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"true,PV1"`
	Other []interface{}
}

func (s *ADT_A06) MessageTypeName() string {
	return "ADT_A06"
}

// ADT_A07 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A07 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"true,PV1"`
	Other []interface{}
}

func (s *ADT_A07) MessageTypeName() string {
	return "ADT_A07"
}

// ADT_A08 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A08 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	NK1   *NK1 `hl7:"true,NK1"`
	PV1   *PV1 `hl7:"true,PV1"`
	DG1   *DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A08) MessageTypeName() string {
	return "ADT_A08"
}

// ADT_A09 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A09 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"true,PV1"`
	DG1   *DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A09) MessageTypeName() string {
	return "ADT_A09"
}

// ADT_A10 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A10 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"true,PV1"`
	DG1   *DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A10) MessageTypeName() string {
	return "ADT_A10"
}

// ADT_A11 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A11 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"true,PV1"`
	DG1   *DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A11) MessageTypeName() string {
	return "ADT_A11"
}

// ADT_A12 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A12 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"true,PV1"`
	DG1   *DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A12) MessageTypeName() string {
	return "ADT_A12"
}

// ADT_A13 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A13 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"true,PV1"`
	DG1   *DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A13) MessageTypeName() string {
	return "ADT_A13"
}

// ADT_A14 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A14 struct {
	MSH *MSH `hl7:"true,MSH"`
	EVN *EVN `hl7:"true,EVN"`
	PID *PID `hl7:"true,PID"`
	// Missing: PD1
	NK1   *NK1 `hl7:"true,NK1"`
	PV1   *PV1 `hl7:"true,PV1"`
	DG1   *DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A14) MessageTypeName() string {
	return "ADT_A14"
}

// ADT_A15 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A15 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"true,PV1"`
	DG1   *DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A15) MessageTypeName() string {
	return "ADT_A15"
}

// ADT_A16 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A16 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"true,PV1"`
	DG1   *DG1 `hl7:"false,DG1"`
	Other []interface{}
}

func (s *ADT_A16) MessageTypeName() string {
	return "ADT_A16"
}

// ADT_A17 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A17 struct {
	MSH     *MSH              `hl7:"true,MSH"`
	EVN     *EVN              `hl7:"true,EVN"`
	PATIENT []ADT_A17_PATIENT `hl7:"true,PATIENT"`
	Other   []interface{}
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
// Definition from HL7 2.1
type ADT_A18 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	MRG   *MRG `hl7:"true,MRG"`
	PV1   *PV1 `hl7:"false,PV1"`
	Other []interface{}
}

func (s *ADT_A18) MessageTypeName() string {
	return "ADT_A18"
}

// ADT_A20 represents the corresponding HL7 message type.
// Definition from HL7 2.1
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
// Definition from HL7 2.1
type ADT_A21 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"true,PV1"`
	Other []interface{}
}

func (s *ADT_A21) MessageTypeName() string {
	return "ADT_A21"
}

// ADT_A22 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A22 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"true,PV1"`
	Other []interface{}
}

func (s *ADT_A22) MessageTypeName() string {
	return "ADT_A22"
}

// ADT_A23 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A23 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"true,PV1"`
	Other []interface{}
}

func (s *ADT_A23) MessageTypeName() string {
	return "ADT_A23"
}

// ADT_A24 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ADT_A24 struct {
	MSH   *MSH `hl7:"true,MSH"`
	EVN   *EVN `hl7:"true,EVN"`
	PID1  *PID `hl7:"true,PID1"`
	PV1   *PV1 `hl7:"true,PV1"`
	PID2  *PID `hl7:"true,PID2"`
	Other []interface{}
}

func (s *ADT_A24) MessageTypeName() string {
	return "ADT_A24"
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
// Definition from HL7 2.1
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

// BAR_P01_VISIT represents the corresponding HL7 message type.
// Definition from HL7 2.1
type BAR_P01_VISIT struct {
	PV1   *PV1  `hl7:"false,PV1"`
	DG1   []DG1 `hl7:"false,DG1"`
	PR1   []PR1 `hl7:"false,PR1"`
	GT1   []GT1 `hl7:"false,GT1"`
	NK1   []NK1 `hl7:"false,NK1"`
	IN1   []IN1 `hl7:"false,IN1"`
	ACC   *ACC  `hl7:"false,ACC"`
	UB1   *UB1  `hl7:"false,UB1"`
	Other []interface{}
}

func (s *BAR_P01_VISIT) MessageTypeSubStructName() string {
	return "BAR_P01_VISIT"
}

// BAR_P02 represents the corresponding HL7 message type.
// Definition from HL7 2.1
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
// Definition from HL7 2.1
type BAR_P02_PATIENT struct {
	PID   *PID `hl7:"true,PID"`
	PV1   *PV1 `hl7:"false,PV1"`
	Other []interface{}
}

func (s *BAR_P02_PATIENT) MessageTypeSubStructName() string {
	return "BAR_P02_PATIENT"
}

// DFT_P03 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type DFT_P03 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	EVN   *EVN  `hl7:"true,EVN"`
	PID   *PID  `hl7:"true,PID"`
	PV1   *PV1  `hl7:"false,PV1"`
	FT1   []FT1 `hl7:"false,FT1"`
	Other []interface{}
}

func (s *DFT_P03) MessageTypeName() string {
	return "DFT_P03"
}

// DSR_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type DSR_Q01 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	MSA   *MSA  `hl7:"true,MSA"`
	QRD   *QRD  `hl7:"true,QRD"`
	QRF   *QRF  `hl7:"false,QRF"`
	DSP   []DSP `hl7:"true,DSP"`
	DSC   *DSC  `hl7:"true,DSC"`
	Other []interface{}
}

func (s *DSR_Q01) MessageTypeName() string {
	return "DSR_Q01"
}

// DSR_Q03 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type DSR_Q03 struct {
	MSH   *MSH  `hl7:"true,MSH"`
	QRD   *QRD  `hl7:"true,QRD"`
	QRF   *QRF  `hl7:"false,QRF"`
	DSP   []DSP `hl7:"true,DSP"`
	DSC   *DSC  `hl7:"true,DSC"`
	Other []interface{}
}

func (s *DSR_Q03) MessageTypeName() string {
	return "DSR_Q03"
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

// ORM_O01_CHOICE represents the corresponding HL7 message type.
// Definition from HL7 2.1
type ORM_O01_CHOICE struct {
	// Only one of the following fields will be set.
	OBR   *OBR `hl7:"true,OBR"`
	ORO   *ORO `hl7:"true,ORO"`
	RX1   *RX1 `hl7:"true,RX1"`
	Other []interface{}
}

func (s *ORM_O01_CHOICE) MessageTypeSubStructName() string {
	return "ORM_O01_CHOICE"
}

// ORM_O01 represents the corresponding HL7 message type.
// Definition from HL7 2.1
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
// Definition from HL7 2.1
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
// Definition from HL7 2.1
type ORM_O01_ORDER_DETAIL struct {
	CHOICE *ORM_O01_CHOICE `hl7:"true,CHOICE"`
	NTE1   []NTE           `hl7:"false,NTE1"`
	OBX    []OBX           `hl7:"false,OBX"`
	NTE2   []NTE           `hl7:"false,NTE2"`
	Other  []interface{}
}

func (s *ORM_O01_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "ORM_O01_ORDER_DETAIL"
}

// ORM_O01_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.1
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
// Definition from HL7 2.1
type ORR_O02_CHOICE struct {
	// Only one of the following fields will be set.
	OBR   *OBR `hl7:"true,OBR"`
	ORO   *ORO `hl7:"true,ORO"`
	RX1   *RX1 `hl7:"true,RX1"`
	Other []interface{}
}

func (s *ORR_O02_CHOICE) MessageTypeSubStructName() string {
	return "ORR_O02_CHOICE"
}

// ORR_O02 represents the corresponding HL7 message type.
// Definition from HL7 2.1
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
// Definition from HL7 2.1
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
// Definition from HL7 2.1
type ORR_O02_ORDER_DETAIL struct {
	CHOICE *ORR_O02_CHOICE `hl7:"true,CHOICE"`
	Other  []interface{}
}

func (s *ORR_O02_ORDER_DETAIL) MessageTypeSubStructName() string {
	return "ORR_O02_ORDER_DETAIL"
}

// ORR_O02_PATIENT represents the corresponding HL7 message type.
// Definition from HL7 2.1
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
// Definition from HL7 2.1
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
// Definition from HL7 2.1
type ORU_R01_OBSERVATION struct {
	OBX   *OBX  `hl7:"false,OBX"`
	NTE   []NTE `hl7:"false,NTE"`
	Other []interface{}
}

func (s *ORU_R01_OBSERVATION) MessageTypeSubStructName() string {
	return "ORU_R01_OBSERVATION"
}

// ORU_R01_ORDER_OBSERVATION represents the corresponding HL7 message type.
// Definition from HL7 2.1
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
// Definition from HL7 2.1
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
// Definition from HL7 2.1
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
	PV1 *PV1 `hl7:"true,PV1"`
	// Missing: PV2
	Other []interface{}
}

func (s *ORU_R32_VISIT) MessageTypeSubStructName() string {
	return "ORU_R32_VISIT"
}

// QRY_A19 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type QRY_A19 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QRD   *QRD `hl7:"true,QRD"`
	Other []interface{}
}

func (s *QRY_A19) MessageTypeName() string {
	return "QRY_A19"
}

// QRY_Q01 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type QRY_Q01 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QRD   *QRD `hl7:"true,QRD"`
	QRF   *QRF `hl7:"false,QRF"`
	DSC   *DSC `hl7:"true,DSC"`
	Other []interface{}
}

func (s *QRY_Q01) MessageTypeName() string {
	return "QRY_Q01"
}

// QRY_Q02 represents the corresponding HL7 message type.
// Definition from HL7 2.1
type QRY_Q02 struct {
	MSH   *MSH `hl7:"true,MSH"`
	QRD   *QRD `hl7:"true,QRD"`
	QRF   *QRF `hl7:"false,QRF"`
	DSC   *DSC `hl7:"true,DSC"`
	Other []interface{}
}

func (s *QRY_Q02) MessageTypeName() string {
	return "QRY_Q02"
}

// UDM_Q05 represents the corresponding HL7 message type.
// Definition from HL7 2.1
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

// PV1 returns the first PV1 segment within the message, or nil if there isn't one.
func (m *Message) PV1() (*PV1, error) {
	ps, err := m.Parse("PV1")
	pst, ok := ps.(*PV1)
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

// RX1 returns the first RX1 segment within the message, or nil if there isn't one.
func (m *Message) RX1() (*RX1, error) {
	ps, err := m.Parse("RX1")
	pst, ok := ps.(*RX1)
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

// AllPV1 returns a slice containing all PV1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllPV1() ([]*PV1, error) {
	pss, err := m.ParseAll("PV1")
	return pss.([]*PV1), err
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

// AllRX1 returns a slice containing all RX1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllRX1() ([]*RX1, error) {
	pss, err := m.ParseAll("RX1")
	return pss.([]*RX1), err
}

// AllUB1 returns a slice containing all UB1 segments within the message,
// or an empty slice if there aren't any.
func (m *Message) AllUB1() ([]*UB1, error) {
	pss, err := m.ParseAll("UB1")
	return pss.([]*UB1), err
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
	msh *MSH             // Required
	msa *MSA             // Required
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
		"qrd": m.qrd,
		"pid": m.pid,
		"dsc": m.dsc,
	}, nil
}

type ADR_A19_PIDv2 struct {
	evn *EVN
	pid *PID // Required
	pv1 *PV1 // Required
}

func (m *ADR_A19_PIDv2) EVN() *EVN {
	return m.evn
}

func (m *ADR_A19_PIDv2) PID() *PID {
	return m.pid
}

func (m *ADR_A19_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m ADR_A19_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
	}, nil
}

type ADT_A01v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 *NK1 // Required
	pv1 *PV1 // Required
	dg1 *DG1
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

func (m *ADT_A01v2) NK1() *NK1 {
	return m.nk1
}

func (m *ADT_A01v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A01v2) DG1() *DG1 {
	return m.dg1
}

func (m ADT_A01v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"dg1": m.dg1,
	}, nil
}

type ADT_A02v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
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

func (m ADT_A02v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
	}, nil
}

type ADT_A03v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
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

func (m ADT_A03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
	}, nil
}

type ADT_A04v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 *NK1 // Required
	pv1 *PV1 // Required
	dg1 *DG1
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

func (m *ADT_A04v2) NK1() *NK1 {
	return m.nk1
}

func (m *ADT_A04v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A04v2) DG1() *DG1 {
	return m.dg1
}

func (m ADT_A04v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"dg1": m.dg1,
	}, nil
}

type ADT_A05v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 *NK1 // Required
	pv1 *PV1 // Required
	dg1 *DG1
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

func (m *ADT_A05v2) NK1() *NK1 {
	return m.nk1
}

func (m *ADT_A05v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A05v2) DG1() *DG1 {
	return m.dg1
}

func (m ADT_A05v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"dg1": m.dg1,
	}, nil
}

type ADT_A06v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
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

func (m *ADT_A06v2) PV1() *PV1 {
	return m.pv1
}

func (m ADT_A06v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
	}, nil
}

type ADT_A07v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
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

func (m *ADT_A07v2) PV1() *PV1 {
	return m.pv1
}

func (m ADT_A07v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
	}, nil
}

type ADT_A08v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	nk1 *NK1 // Required
	pv1 *PV1 // Required
	dg1 *DG1
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

func (m *ADT_A08v2) NK1() *NK1 {
	return m.nk1
}

func (m *ADT_A08v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A08v2) DG1() *DG1 {
	return m.dg1
}

func (m ADT_A08v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"dg1": m.dg1,
	}, nil
}

type ADT_A09v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	dg1 *DG1
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

func (m *ADT_A09v2) DG1() *DG1 {
	return m.dg1
}

func (m ADT_A09v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"dg1": m.dg1,
	}, nil
}

type ADT_A10v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	dg1 *DG1
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

func (m *ADT_A10v2) DG1() *DG1 {
	return m.dg1
}

func (m ADT_A10v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"dg1": m.dg1,
	}, nil
}

type ADT_A11v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	dg1 *DG1
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

func (m *ADT_A11v2) DG1() *DG1 {
	return m.dg1
}

func (m ADT_A11v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"dg1": m.dg1,
	}, nil
}

type ADT_A12v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	dg1 *DG1
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

func (m *ADT_A12v2) DG1() *DG1 {
	return m.dg1
}

func (m ADT_A12v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"dg1": m.dg1,
	}, nil
}

type ADT_A13v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	dg1 *DG1
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

func (m *ADT_A13v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A13v2) DG1() *DG1 {
	return m.dg1
}

func (m ADT_A13v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"dg1": m.dg1,
	}, nil
}

type ADT_A14v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pd1 *PD1 // Required
	nk1 *NK1 // Required
	pv1 *PV1 // Required
	dg1 *DG1
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

func (m *ADT_A14v2) PD1() *PD1 {
	return m.pd1
}

func (m *ADT_A14v2) NK1() *NK1 {
	return m.nk1
}

func (m *ADT_A14v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A14v2) DG1() *DG1 {
	return m.dg1
}

func (m ADT_A14v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pd1": m.pd1,
		"nk1": m.nk1,
		"pv1": m.pv1,
		"dg1": m.dg1,
	}, nil
}

type ADT_A15v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	dg1 *DG1
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

func (m *ADT_A15v2) DG1() *DG1 {
	return m.dg1
}

func (m ADT_A15v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"dg1": m.dg1,
	}, nil
}

type ADT_A16v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
	dg1 *DG1
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

func (m *ADT_A16v2) DG1() *DG1 {
	return m.dg1
}

func (m ADT_A16v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"dg1": m.dg1,
	}, nil
}

type ADT_A17v2 struct {
	msh *MSH             // Required
	evn *EVN             // Required
	pid []*ADT_A17_PIDv2 // Required
}

func (m *ADT_A17v2) MSH() *MSH {
	return m.msh
}

func (m *ADT_A17v2) EVN() *EVN {
	return m.evn
}

func (m *ADT_A17v2) GroupByPID() []*ADT_A17_PIDv2 {
	return m.pid
}

func (m ADT_A17v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
	}, nil
}

type ADT_A17_PIDv2 struct {
	pid *PID // Required
	pv1 *PV1 // Required
}

func (m *ADT_A17_PIDv2) PID() *PID {
	return m.pid
}

func (m *ADT_A17_PIDv2) PV1() *PV1 {
	return m.pv1
}

func (m ADT_A17_PIDv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pid": m.pid,
		"pv1": m.pv1,
	}, nil
}

type ADT_A18v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	mrg *MRG // Required
	pv1 *PV1
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

func (m ADT_A21v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
	}, nil
}

type ADT_A22v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
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

func (m ADT_A22v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
	}, nil
}

type ADT_A23v2 struct {
	msh *MSH // Required
	evn *EVN // Required
	pid *PID // Required
	pv1 *PV1 // Required
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

func (m ADT_A23v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
	}, nil
}

type ADT_A24v2 struct {
	msh  *MSH // Required
	evn  *EVN // Required
	pid1 *PID // Required
	pv1  *PV1 // Required
	pid2 *PID // Required
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

func (m *ADT_A24v2) PV1() *PV1 {
	return m.pv1
}

func (m *ADT_A24v2) PID2() *PID {
	return m.pid2
}

func (m ADT_A24v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh":  m.msh,
		"evn":  m.evn,
		"pid1": m.pid1,
		"pv1":  m.pv1,
		"pid2": m.pid2,
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
	dg1 []*DG1
	pr1 []*PR1
	gt1 []*GT1
	nk1 []*NK1
	in1 []*IN1
	acc *ACC
	ub1 *UB1
}

func (m *BAR_P01_PV1v2) PV1() *PV1 {
	return m.pv1
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

func (m *BAR_P01_PV1v2) AllIN1() []*IN1 {
	return m.in1
}

func (m *BAR_P01_PV1v2) ACC() *ACC {
	return m.acc
}

func (m *BAR_P01_PV1v2) UB1() *UB1 {
	return m.ub1
}

func (m BAR_P01_PV1v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"pv1": m.pv1,
		"dg1": m.dg1,
		"pr1": m.pr1,
		"gt1": m.gt1,
		"nk1": m.nk1,
		"in1": m.in1,
		"acc": m.acc,
		"ub1": m.ub1,
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
	ft1 []*FT1
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

func (m *DFT_P03v2) AllFT1() []*FT1 {
	return m.ft1
}

func (m DFT_P03v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"evn": m.evn,
		"pid": m.pid,
		"pv1": m.pv1,
		"ft1": m.ft1,
	}, nil
}

type DSR_Q01v2 struct {
	msh *MSH // Required
	msa *MSA // Required
	qrd *QRD // Required
	qrf *QRF
	dsp []*DSP // Required
	dsc *DSC   // Required
}

func (m *DSR_Q01v2) MSH() *MSH {
	return m.msh
}

func (m *DSR_Q01v2) MSA() *MSA {
	return m.msa
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
		"qrd": m.qrd,
		"qrf": m.qrf,
		"dsp": m.dsp,
		"dsc": m.dsc,
	}, nil
}

type DSR_Q03v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	dsp []*DSP // Required
	dsc *DSC   // Required
}

func (m *DSR_Q03v2) MSH() *MSH {
	return m.msh
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
	oro  *ORO // Required
	rx1  *RX1 // Required
	nte1 []*NTE
	obx  []*OBX
	nte2 []*NTE
	blg  *BLG
}

func (m *ORM_O01_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *ORM_O01_ORCv2) OBR() *OBR {
	return m.obr
}

func (m *ORM_O01_ORCv2) ORO() *ORO {
	return m.oro
}

func (m *ORM_O01_ORCv2) RX1() *RX1 {
	return m.rx1
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
		"oro":  m.oro,
		"rx1":  m.rx1,
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
	oro *ORO // Required
	rx1 *RX1 // Required
	nte []*NTE
}

func (m *ORR_O02_PID_ORCv2) ORC() *ORC {
	return m.orc
}

func (m *ORR_O02_PID_ORCv2) OBR() *OBR {
	return m.obr
}

func (m *ORR_O02_PID_ORCv2) ORO() *ORO {
	return m.oro
}

func (m *ORR_O02_PID_ORCv2) RX1() *RX1 {
	return m.rx1
}

func (m *ORR_O02_PID_ORCv2) AllNTE() []*NTE {
	return m.nte
}

func (m ORR_O02_PID_ORCv2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"orc": m.orc,
		"obr": m.obr,
		"oro": m.oro,
		"rx1": m.rx1,
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
}

func (m *QRY_A19v2) MSH() *MSH {
	return m.msh
}

func (m *QRY_A19v2) QRD() *QRD {
	return m.qrd
}

func (m QRY_A19v2) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"msh": m.msh,
		"qrd": m.qrd,
	}, nil
}

type QRY_Q01v2 struct {
	msh *MSH // Required
	qrd *QRD // Required
	qrf *QRF
	dsc *DSC // Required
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
	dsc *DSC // Required
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
	Types["ORC"] = reflect.TypeOf(ORC{})
	Types["ORO"] = reflect.TypeOf(ORO{})
	Types["PID"] = reflect.TypeOf(PID{})
	Types["PR1"] = reflect.TypeOf(PR1{})
	Types["PV1"] = reflect.TypeOf(PV1{})
	Types["QRD"] = reflect.TypeOf(QRD{})
	Types["QRF"] = reflect.TypeOf(QRF{})
	Types["RX1"] = reflect.TypeOf(RX1{})
	Types["UB1"] = reflect.TypeOf(UB1{})
	Types["URD"] = reflect.TypeOf(URD{})
	Types["URS"] = reflect.TypeOf(URS{})
	Types["ACK"] = reflect.TypeOf(ACK{})
	Types["ADR_A19"] = reflect.TypeOf(ADR_A19{})
	Types["ADR_A19_QUERY_RESPONSE"] = reflect.TypeOf(ADR_A19_QUERY_RESPONSE{})
	Types["ADT_A01"] = reflect.TypeOf(ADT_A01{})
	Types["ADT_A02"] = reflect.TypeOf(ADT_A02{})
	Types["ADT_A03"] = reflect.TypeOf(ADT_A03{})
	Types["ADT_A04"] = reflect.TypeOf(ADT_A04{})
	Types["ADT_A05"] = reflect.TypeOf(ADT_A05{})
	Types["ADT_A06"] = reflect.TypeOf(ADT_A06{})
	Types["ADT_A07"] = reflect.TypeOf(ADT_A07{})
	Types["ADT_A08"] = reflect.TypeOf(ADT_A08{})
	Types["ADT_A09"] = reflect.TypeOf(ADT_A09{})
	Types["ADT_A10"] = reflect.TypeOf(ADT_A10{})
	Types["ADT_A11"] = reflect.TypeOf(ADT_A11{})
	Types["ADT_A12"] = reflect.TypeOf(ADT_A12{})
	Types["ADT_A13"] = reflect.TypeOf(ADT_A13{})
	Types["ADT_A14"] = reflect.TypeOf(ADT_A14{})
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
	Types["ADT_A40"] = reflect.TypeOf(ADT_A40{})
	Types["ADT_A40_PATIENT"] = reflect.TypeOf(ADT_A40_PATIENT{})
	Types["ADT_A44"] = reflect.TypeOf(ADT_A44{})
	Types["ADT_A44_PATIENT"] = reflect.TypeOf(ADT_A44_PATIENT{})
	Types["BAR_P01"] = reflect.TypeOf(BAR_P01{})
	Types["BAR_P01_VISIT"] = reflect.TypeOf(BAR_P01_VISIT{})
	Types["BAR_P02"] = reflect.TypeOf(BAR_P02{})
	Types["BAR_P02_PATIENT"] = reflect.TypeOf(BAR_P02_PATIENT{})
	Types["DFT_P03"] = reflect.TypeOf(DFT_P03{})
	Types["DSR_Q01"] = reflect.TypeOf(DSR_Q01{})
	Types["DSR_Q03"] = reflect.TypeOf(DSR_Q03{})
	Types["MCF_Q02"] = reflect.TypeOf(MCF_Q02{})
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
	Types["QRY_Q01"] = reflect.TypeOf(QRY_Q01{})
	Types["QRY_Q02"] = reflect.TypeOf(QRY_Q02{})
	Types["UDM_Q05"] = reflect.TypeOf(UDM_Q05{})
	Types["GenericHL7Segment"] = reflect.TypeOf(GenericHL7Segment{})
	Types["ACKv2"] = reflect.TypeOf(ACKv2{})
	Types["ADR_A19v2"] = reflect.TypeOf(ADR_A19v2{})
	Types["ADR_A19_PIDv2"] = reflect.TypeOf(ADR_A19_PIDv2{})
	Types["ADT_A01v2"] = reflect.TypeOf(ADT_A01v2{})
	Types["ADT_A02v2"] = reflect.TypeOf(ADT_A02v2{})
	Types["ADT_A03v2"] = reflect.TypeOf(ADT_A03v2{})
	Types["ADT_A04v2"] = reflect.TypeOf(ADT_A04v2{})
	Types["ADT_A05v2"] = reflect.TypeOf(ADT_A05v2{})
	Types["ADT_A06v2"] = reflect.TypeOf(ADT_A06v2{})
	Types["ADT_A07v2"] = reflect.TypeOf(ADT_A07v2{})
	Types["ADT_A08v2"] = reflect.TypeOf(ADT_A08v2{})
	Types["ADT_A09v2"] = reflect.TypeOf(ADT_A09v2{})
	Types["ADT_A10v2"] = reflect.TypeOf(ADT_A10v2{})
	Types["ADT_A11v2"] = reflect.TypeOf(ADT_A11v2{})
	Types["ADT_A12v2"] = reflect.TypeOf(ADT_A12v2{})
	Types["ADT_A13v2"] = reflect.TypeOf(ADT_A13v2{})
	Types["ADT_A14v2"] = reflect.TypeOf(ADT_A14v2{})
	Types["ADT_A15v2"] = reflect.TypeOf(ADT_A15v2{})
	Types["ADT_A16v2"] = reflect.TypeOf(ADT_A16v2{})
	Types["ADT_A17v2"] = reflect.TypeOf(ADT_A17v2{})
	Types["ADT_A17_PIDv2"] = reflect.TypeOf(ADT_A17_PIDv2{})
	Types["ADT_A18v2"] = reflect.TypeOf(ADT_A18v2{})
	Types["ADT_A20v2"] = reflect.TypeOf(ADT_A20v2{})
	Types["ADT_A21v2"] = reflect.TypeOf(ADT_A21v2{})
	Types["ADT_A22v2"] = reflect.TypeOf(ADT_A22v2{})
	Types["ADT_A23v2"] = reflect.TypeOf(ADT_A23v2{})
	Types["ADT_A24v2"] = reflect.TypeOf(ADT_A24v2{})
	Types["ADT_A40v2"] = reflect.TypeOf(ADT_A40v2{})
	Types["ADT_A40_PIDv2"] = reflect.TypeOf(ADT_A40_PIDv2{})
	Types["ADT_A44v2"] = reflect.TypeOf(ADT_A44v2{})
	Types["ADT_A44_PIDv2"] = reflect.TypeOf(ADT_A44_PIDv2{})
	Types["BAR_P01v2"] = reflect.TypeOf(BAR_P01v2{})
	Types["BAR_P01_PV1v2"] = reflect.TypeOf(BAR_P01_PV1v2{})
	Types["BAR_P02v2"] = reflect.TypeOf(BAR_P02v2{})
	Types["BAR_P02_PIDv2"] = reflect.TypeOf(BAR_P02_PIDv2{})
	Types["DFT_P03v2"] = reflect.TypeOf(DFT_P03v2{})
	Types["DSR_Q01v2"] = reflect.TypeOf(DSR_Q01v2{})
	Types["DSR_Q03v2"] = reflect.TypeOf(DSR_Q03v2{})
	Types["MCF_Q02v2"] = reflect.TypeOf(MCF_Q02v2{})
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
	Types["QRY_Q01v2"] = reflect.TypeOf(QRY_Q01v2{})
	Types["QRY_Q02v2"] = reflect.TypeOf(QRY_Q02v2{})
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
	"ADR_A19_PIDv2.evn": StringSet{
		"DSC": true,
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADR_A19_PIDv2.pid": StringSet{
		"DSC": true,
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADR_A19_PIDv2.pv1": StringSet{
		"DSC": true,
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADR_A19v2.dsc": StringSet{},
	"ADR_A19v2.msa": StringSet{
		"DSC": true,
		"EVN": true,
		"PID": true,
		"PV1": true,
		"QRD": true,
	},
	"ADR_A19v2.msh": StringSet{
		"DSC": true,
		"EVN": true,
		"MSA": true,
		"PID": true,
		"PV1": true,
		"QRD": true,
	},
	"ADR_A19v2.qrd": StringSet{
		"DSC": true,
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A01v2.dg1": StringSet{},
	"ADT_A01v2.evn": StringSet{
		"DG1": true,
		"NK1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A01v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"NK1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A01v2.nk1": StringSet{
		"DG1": true,
		"PV1": true,
	},
	"ADT_A01v2.pid": StringSet{
		"DG1": true,
		"NK1": true,
		"PV1": true,
	},
	"ADT_A01v2.pv1": StringSet{
		"DG1": true,
	},
	"ADT_A02v2.evn": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A02v2.msh": StringSet{
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A02v2.pid": StringSet{
		"PV1": true,
	},
	"ADT_A02v2.pv1": StringSet{},
	"ADT_A03v2.evn": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A03v2.msh": StringSet{
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A03v2.pid": StringSet{
		"PV1": true,
	},
	"ADT_A03v2.pv1": StringSet{},
	"ADT_A04v2.dg1": StringSet{},
	"ADT_A04v2.evn": StringSet{
		"DG1": true,
		"NK1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A04v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"NK1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A04v2.nk1": StringSet{
		"DG1": true,
		"PV1": true,
	},
	"ADT_A04v2.pid": StringSet{
		"DG1": true,
		"NK1": true,
		"PV1": true,
	},
	"ADT_A04v2.pv1": StringSet{
		"DG1": true,
	},
	"ADT_A05v2.dg1": StringSet{},
	"ADT_A05v2.evn": StringSet{
		"DG1": true,
		"NK1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A05v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"NK1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A05v2.nk1": StringSet{
		"DG1": true,
		"PV1": true,
	},
	"ADT_A05v2.pid": StringSet{
		"DG1": true,
		"NK1": true,
		"PV1": true,
	},
	"ADT_A05v2.pv1": StringSet{
		"DG1": true,
	},
	"ADT_A06v2.evn": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A06v2.msh": StringSet{
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A06v2.pid": StringSet{
		"PV1": true,
	},
	"ADT_A06v2.pv1": StringSet{},
	"ADT_A07v2.evn": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A07v2.msh": StringSet{
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A07v2.pid": StringSet{
		"PV1": true,
	},
	"ADT_A07v2.pv1": StringSet{},
	"ADT_A08v2.dg1": StringSet{},
	"ADT_A08v2.evn": StringSet{
		"DG1": true,
		"NK1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A08v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"NK1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A08v2.nk1": StringSet{
		"DG1": true,
		"PV1": true,
	},
	"ADT_A08v2.pid": StringSet{
		"DG1": true,
		"NK1": true,
		"PV1": true,
	},
	"ADT_A08v2.pv1": StringSet{
		"DG1": true,
	},
	"ADT_A09v2.dg1": StringSet{},
	"ADT_A09v2.evn": StringSet{
		"DG1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A09v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A09v2.pid": StringSet{
		"DG1": true,
		"PV1": true,
	},
	"ADT_A09v2.pv1": StringSet{
		"DG1": true,
	},
	"ADT_A10v2.dg1": StringSet{},
	"ADT_A10v2.evn": StringSet{
		"DG1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A10v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A10v2.pid": StringSet{
		"DG1": true,
		"PV1": true,
	},
	"ADT_A10v2.pv1": StringSet{
		"DG1": true,
	},
	"ADT_A11v2.dg1": StringSet{},
	"ADT_A11v2.evn": StringSet{
		"DG1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A11v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A11v2.pid": StringSet{
		"DG1": true,
		"PV1": true,
	},
	"ADT_A11v2.pv1": StringSet{
		"DG1": true,
	},
	"ADT_A12v2.dg1": StringSet{},
	"ADT_A12v2.evn": StringSet{
		"DG1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A12v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A12v2.pid": StringSet{
		"DG1": true,
		"PV1": true,
	},
	"ADT_A12v2.pv1": StringSet{
		"DG1": true,
	},
	"ADT_A13v2.dg1": StringSet{},
	"ADT_A13v2.evn": StringSet{
		"DG1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A13v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A13v2.pid": StringSet{
		"DG1": true,
		"PV1": true,
	},
	"ADT_A13v2.pv1": StringSet{
		"DG1": true,
	},
	"ADT_A14v2.dg1": StringSet{},
	"ADT_A14v2.evn": StringSet{
		"DG1": true,
		"NK1": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A14v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"NK1": true,
		"PD1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A14v2.nk1": StringSet{
		"DG1": true,
		"PV1": true,
	},
	"ADT_A14v2.pd1": StringSet{
		"DG1": true,
		"NK1": true,
		"PV1": true,
	},
	"ADT_A14v2.pid": StringSet{
		"DG1": true,
		"NK1": true,
		"PD1": true,
		"PV1": true,
	},
	"ADT_A14v2.pv1": StringSet{
		"DG1": true,
	},
	"ADT_A15v2.dg1": StringSet{},
	"ADT_A15v2.evn": StringSet{
		"DG1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A15v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A15v2.pid": StringSet{
		"DG1": true,
		"PV1": true,
	},
	"ADT_A15v2.pv1": StringSet{
		"DG1": true,
	},
	"ADT_A16v2.dg1": StringSet{},
	"ADT_A16v2.evn": StringSet{
		"DG1": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A16v2.msh": StringSet{
		"DG1": true,
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A16v2.pid": StringSet{
		"DG1": true,
		"PV1": true,
	},
	"ADT_A16v2.pv1": StringSet{
		"DG1": true,
	},
	"ADT_A17_PIDv2.pid": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A17_PIDv2.pv1": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A17v2.evn": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A17v2.msh": StringSet{
		"EVN": true,
		"PID": true,
		"PV1": true,
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
		"PID": true,
		"PV1": true,
	},
	"ADT_A21v2.msh": StringSet{
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A21v2.pid": StringSet{
		"PV1": true,
	},
	"ADT_A21v2.pv1": StringSet{},
	"ADT_A22v2.evn": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A22v2.msh": StringSet{
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A22v2.pid": StringSet{
		"PV1": true,
	},
	"ADT_A22v2.pv1": StringSet{},
	"ADT_A23v2.evn": StringSet{
		"PID": true,
		"PV1": true,
	},
	"ADT_A23v2.msh": StringSet{
		"EVN": true,
		"PID": true,
		"PV1": true,
	},
	"ADT_A23v2.pid": StringSet{
		"PV1": true,
	},
	"ADT_A23v2.pv1": StringSet{},
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
	"ADT_A24v2.pid2": StringSet{},
	"ADT_A24v2.pv1": StringSet{
		"PID": true,
	},
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
	"BAR_P01_PV1v2.acc": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"NK1": true,
		"PR1": true,
		"PV1": true,
		"UB1": true,
	},
	"BAR_P01_PV1v2.dg1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"NK1": true,
		"PR1": true,
		"PV1": true,
		"UB1": true,
	},
	"BAR_P01_PV1v2.gt1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"NK1": true,
		"PR1": true,
		"PV1": true,
		"UB1": true,
	},
	"BAR_P01_PV1v2.in1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"NK1": true,
		"PR1": true,
		"PV1": true,
		"UB1": true,
	},
	"BAR_P01_PV1v2.nk1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"NK1": true,
		"PR1": true,
		"PV1": true,
		"UB1": true,
	},
	"BAR_P01_PV1v2.pr1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"NK1": true,
		"PR1": true,
		"PV1": true,
		"UB1": true,
	},
	"BAR_P01_PV1v2.pv1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"NK1": true,
		"PR1": true,
		"PV1": true,
		"UB1": true,
	},
	"BAR_P01_PV1v2.ub1": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"NK1": true,
		"PR1": true,
		"PV1": true,
		"UB1": true,
	},
	"BAR_P01v2.evn": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"NK1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"UB1": true,
	},
	"BAR_P01v2.msh": StringSet{
		"ACC": true,
		"DG1": true,
		"EVN": true,
		"GT1": true,
		"IN1": true,
		"NK1": true,
		"PID": true,
		"PR1": true,
		"PV1": true,
		"UB1": true,
	},
	"BAR_P01v2.pid": StringSet{
		"ACC": true,
		"DG1": true,
		"GT1": true,
		"IN1": true,
		"NK1": true,
		"PR1": true,
		"PV1": true,
		"UB1": true,
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
		"PID": true,
		"PV1": true,
	},
	"DFT_P03v2.ft1": StringSet{
		"FT1": true,
	},
	"DFT_P03v2.msh": StringSet{
		"EVN": true,
		"FT1": true,
		"PID": true,
		"PV1": true,
	},
	"DFT_P03v2.pid": StringSet{
		"FT1": true,
		"PV1": true,
	},
	"DFT_P03v2.pv1": StringSet{
		"FT1": true,
	},
	"DSR_Q01v2.dsc": StringSet{},
	"DSR_Q01v2.dsp": StringSet{
		"DSC": true,
		"DSP": true,
	},
	"DSR_Q01v2.msa": StringSet{
		"DSC": true,
		"DSP": true,
		"QRD": true,
		"QRF": true,
	},
	"DSR_Q01v2.msh": StringSet{
		"DSC": true,
		"DSP": true,
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
	"DSR_Q03v2.msh": StringSet{
		"DSC": true,
		"DSP": true,
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
	"MCF_Q02v2.msa": StringSet{},
	"MCF_Q02v2.msh": StringSet{
		"MSA": true,
	},
	"ORM_O01_ORCv2.blg": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"ORO": true,
		"RX1": true,
	},
	"ORM_O01_ORCv2.nte1": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"ORO": true,
		"RX1": true,
	},
	"ORM_O01_ORCv2.nte2": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"ORO": true,
		"RX1": true,
	},
	"ORM_O01_ORCv2.obr": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"ORO": true,
		"RX1": true,
	},
	"ORM_O01_ORCv2.obx": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"ORO": true,
		"RX1": true,
	},
	"ORM_O01_ORCv2.orc": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"ORO": true,
		"RX1": true,
	},
	"ORM_O01_ORCv2.oro": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"ORO": true,
		"RX1": true,
	},
	"ORM_O01_ORCv2.rx1": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"ORO": true,
		"RX1": true,
	},
	"ORM_O01_PIDv2.nte": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"ORO": true,
		"PV1": true,
		"RX1": true,
	},
	"ORM_O01_PIDv2.pid": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"ORO": true,
		"PV1": true,
		"RX1": true,
	},
	"ORM_O01_PIDv2.pv1": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"ORO": true,
		"RX1": true,
	},
	"ORM_O01v2.msh": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"ORO": true,
		"PID": true,
		"PV1": true,
		"RX1": true,
	},
	"ORM_O01v2.nte": StringSet{
		"BLG": true,
		"NTE": true,
		"OBR": true,
		"OBX": true,
		"ORC": true,
		"ORO": true,
		"PID": true,
		"PV1": true,
		"RX1": true,
	},
	"ORR_O02_PID_ORCv2.nte": StringSet{
		"NTE": true,
		"OBR": true,
		"ORC": true,
		"ORO": true,
		"RX1": true,
	},
	"ORR_O02_PID_ORCv2.obr": StringSet{
		"NTE": true,
		"OBR": true,
		"ORC": true,
		"ORO": true,
		"RX1": true,
	},
	"ORR_O02_PID_ORCv2.orc": StringSet{
		"NTE": true,
		"OBR": true,
		"ORC": true,
		"ORO": true,
		"RX1": true,
	},
	"ORR_O02_PID_ORCv2.oro": StringSet{
		"NTE": true,
		"OBR": true,
		"ORC": true,
		"ORO": true,
		"RX1": true,
	},
	"ORR_O02_PID_ORCv2.rx1": StringSet{
		"NTE": true,
		"OBR": true,
		"ORC": true,
		"ORO": true,
		"RX1": true,
	},
	"ORR_O02_PIDv2.nte": StringSet{
		"NTE": true,
		"OBR": true,
		"ORC": true,
		"ORO": true,
		"RX1": true,
	},
	"ORR_O02_PIDv2.pid": StringSet{
		"NTE": true,
		"OBR": true,
		"ORC": true,
		"ORO": true,
		"RX1": true,
	},
	"ORR_O02v2.msa": StringSet{
		"NTE": true,
		"OBR": true,
		"ORC": true,
		"ORO": true,
		"PID": true,
		"RX1": true,
	},
	"ORR_O02v2.msh": StringSet{
		"MSA": true,
		"NTE": true,
		"OBR": true,
		"ORC": true,
		"ORO": true,
		"PID": true,
		"RX1": true,
	},
	"ORR_O02v2.nte": StringSet{
		"NTE": true,
		"OBR": true,
		"ORC": true,
		"ORO": true,
		"PID": true,
		"RX1": true,
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
	},
	"QRY_A19v2.qrd": StringSet{},
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
