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

package testhl7

import (
	"fmt"
)

const (
	// https://corepointhealth.com/resource-center/hl7-resources/hl7-msh-message-header
	// The message type field here is set to a placerholder "%s". Users should set this type based
	// on the message type they want to parse, e.g. ADT^A01 or ORU_R01.
	//                ┌ Encoding Characters
	//                |     ┌ Sending Application
	//                |     |   ┌ Sending Facility
	//                |     |   |   ┌ Receiving Application
	//                |     |   |   |        ┌ Receiving Facility
	//                |     |   |   |        |    ┌ Date/Time Of Message
	//                |     |   |   |        |    |              ┌ Security
	//                |     |   |   |        |    |              |┌ Message Type
	//                |     |   |   |        |    |              ||  ┌ Message ControlID
	//                |     |   |   |        |    |              ||  |                      ┌ Processing ID
	//                |     |   |   |        |    |              ||  |                      | ┌ Version ID
	//                |     |   |   |        |    |              ||  |                      | |   ┌ Sequence Number
	//                |     |   |   |        |    |              ||  |                      | |   |┌ Continuation Pointer
	//                |     |   |   |        |    |              ||  |                      | |   ||┌ Accept Acknowledgment Type
	//                |     |   |   |        |    |              ||  |                      | |   |||  ┌ Application Acknowledgment Type
	//                |     |   |   |        |    |              ||  |                      | |   |||  |┌ Country Code
	//                |     |   |   |        |    |              ||  |                      | |   |||  ||  ┌ Character Set
	TemplateMSH = "MSH|^~\\&|SAP|SFA|RECEAPPL|RFAC|20141128001635||%s|2014112800163507740000|T|2.3|||AL||44|ASCII"

	// https://corepointhealth.com/resource-center/hl7-resources/hl7-pv1-patient-visit-information-segment
	//               ┌ ID
	//               │ ┌ Patient Class
	//               │ │         ┌ Assigned Patient Location
	//               │ │         │                                       ┌ Admission Type
	//               │ │         │                                       │   ┌ Preadmit Number
	//               │ │         │                                       │   │┌ Prior Patient Location
	SegmentPV1 = "PV1|1|INPATIENT|SFA 10 EAST^Bay01^Bed10^SFA RF^^BED^SFA|28b||^^^^^^" +
		//    ┌ Attending Doctor
		"" + "|C3335432^Jensen^Mark^^^^^^NHSCSLTNBR^PRSNL^^^NONGP^~891257458214^Jensen^Mark^^^^^^DRNBR^PRSNL^^^ORGDR^" +
		//    ┌ Referring Doctor
		"" + "|C3335432^Jensen^Mark^^^^^^NHSCSLTNBR^PRSNL^^^NONGP^~891257458214^Jensen^Mark^^^^^^DRNBR^PRSNL^^^ORGDR^" +
		//    ┌ Consulting Doctor
		//    │┌ Hospital Service
		//    ││   ┌ Temporary Location
		//    ││   │┌ Preadmit  Indicator
		//    ││   ││┌ Re-admission Indicator
		//    ││   │││┌ Admit Source
		//    ││   ││││  ┌ Ambulatory Status
		//    ││   ││││  │┌ VIP Indicator
		//    ││   ││││  ││┌ Admitting Doctor
		//    ││   ││││  │││┌ Patient Type
		//    ││   ││││  ││││         ┌ Visit Number
		//    ││   ││││  ││││         │                  ┌ Financial Class
		//    ││   ││││  ││││         │                  │┌ Charge Price Indicator
		//    ││   ││││  ││││         │                  ││┌ Courtesy Code
		//    ││   ││││  ││││         │                  │││┌ Credit Rating
		//    ││   ││││  ││││         │                  ││││┌ Contract Code
		//    ││   ││││  ││││         │                  │││││┌ Contract Effective Date
		//    ││   ││││  ││││         │                  ││││││┌ Contract Amount
		//    ││   ││││  ││││         │                  │││││││┌ Contract Period
		//    ││   ││││  ││││         │                  ││││││││┌ Interest Code
		"" + "||361||||19||||INPATIENT|6153479^^^^VISITID|||||||||" +
		//    ┌ Transfer to Bad Debt Code
		//    │┌ Transfer to Bad Debt Date
		//    ││┌ Bad Debt Agency Code
		//    │││┌ Bad Debt Transfer Amount
		//    ││││┌ Bad Debt Recovery Amount
		//    │││││┌ Delete Account Indicator
		//    ││││││┌ Delete Account Date
		//    │││││││┌ Discharge Disposition
		//    ││││││││┌ Discharged to Location
		//    │││││││││┌ Diet Type
		//    ││││││││││┌ Servicing Facility
		//    │││││││││││      ┌ Bed Status
		//    │││││││││││      │┌ Account Status
		//    │││││││││││      ││      ┌ Pending Location
		//    │││││││││││      ││      │┌ Prior Temporary Location
		//    │││││││││││      ││      ││┌ Admit Date/Time
		"" + "|||||||||||SFA RF||ACTIVE|||20141127180000"

	// https://corepointhealth.com/resource-center/hl7-resources/hl7-obr-segment
	//               ┌ ID
	//               │ ┌ Placer Order Number
	//               │ │              ┌ Filler Order Number
	//               │ │              │┌ Universal Service ID
	//               │ │              ││                   ┌ Priority
	//               │ │              ││                   │  ┌ Requested Date/Time
	//               │ │              ││                   │  │              ┌ Observation Date/Timer
	//               │ │              ││                   │  │              │              ┌ Observation End Date/Time
	//               │ │              ││                   │  │              │              │              ┌ Collection Volume
	//               │ │              ││                   │  │              │              │              │┌ Collector Identifier
	//               │ │              ││                   │  │              │              │              ││┌ Specimen Action Code
	//               │ │              ││                   │  │              │              │              │││┌ Danger Code
	//               │ │              ││                   │  │              │              │              ││││┌ Relevant Clinical Info.
	//               │ │              ││                   │  │              │              │              │││││┌ Specimen Received Date/Time
	//               │ │              ││                   │  │              │              │              ││││││┌ Specimen Source
	//               │ │              ││                   │  │              │              │              │││││││      ┌ Ordering Provider
	//               │ │              ││                   │  │              │              │              │││││││      │               ┌ Order Callback Phone Number
	//               │ │              ││                   │  │              │              │              │││││││      │               │┌ Placer Field 1
	//               │ │              ││                   │  │              │              │              │││││││      │               ││┌ Placer Field 2
	//               │ │              ││                   │  │              │              │              │││││││      │               │││┌ Filler Field 1
	//               │ │              ││                   │  │              │              │              │││││││      │               ││││┌ Filler Field 2
	//               │ │              ││                   │  │              │              │              │││││││      │               │││││┌ Results Rpt/Status Chng - Date/Time
	//               │ │              ││                   │  │              │              │              │││││││      │               ││││││              ┌ Charge to Practice
	//               │ │              ││                   │  │              │              │              │││││││      │               ││││││              │┌ Diagnostic Serv Sect ID
	//               │ │              ││                   │  │              │              │              │││││││      │               ││││││              ││  ┌ Result Status
	SegmentOBR = "OBR|1|20060307110114||003038^Urinalysis^L|HI|20060307110113|20060307110114|20060307110114|||||||source|C3335432^Jensen||||||20060307110115||HM|F|"

	// http://hl7-definition.caristix.com:9010/Default.aspx?version=HL7+v2.3.1&segment=ORC
	//               ┌ ID
	//               │ ┌ Placer Order Number
	//               │ │              ┌ Filler Order Number
	//               │ │              │      ┌ Placer Group Number
	//               │ │              │      │┌ Order status
	//               │ │              │      ││ ┌ Response Flag
	//               │ │              │      ││ │┌ Quantity/Timing
	//               │ │              │      ││ ││┌ Parent Order
	//               │ │              │      ││ │││┌ Transaction Date/Time
	SegmentORC = "ORC|1|20060307110114|XXYYZZ||A||||20060307110114"

	// http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/segment/AL1
	//                 ┌ Allergy Type
	//                 │
	//                 │    ┌ Allergy Identifier, Description and Coding System
	//                 │    │								                                               ┌ Allergy Severity
	//                 │    │								                                               │      ┌ Allergy Reaction
	//                 │    │								                                               │      │    ┌ Allergy Identification Date
	//                 │    │								                                               │      │    │
	SegmentAL1 = "AL1|1|DRUG|##NOMEN##,AL1,ceStruct,allergy,1234,1234567^penicillin^ALLERGY|SEVERE|rash|20180428000000"

	// http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/segment/OBX
	//               ┌ Value Type
	//               │    ┌ Observation Identifier
	//               │    │                            ┌ Observation Sub-ID
	//               │    │                            │┌ Observation Value
	//               │    │                            ││   ┌ Units
	//               │    │                            ││   │           ┌ Reference Range
	//               │    │                            ││   │           │        ┌ Abnormal Flags
	//               │    │                            ││   │           │        │ ┌ Probability
	//               │    │                            ││   │           │        │ │┌ Nature of Abnormal
	//               │    │                            ││   │           │        │ ││┌ Observation Result Status
	//               │    │                            ││   │           │        │ │││ ┌ Date Last Obs Normal Values
	//               │    │                            ││   │           │        │ │││ │┌ User Defined Access Checks
	//               │    │                            ││   │           │        │ │││ ││┌ Date/Time of the Observation
	//               │    │                            ││   │           │        │ │││ │││┌ Producer's ID
	//               │    │                            ││   │           │        │ │││ ││││┌ Responsible Observer
	//               │    │                            ││   │           │        │ │││ │││││
	SegmentOBX  = "OBX|1|NM|lpdc-2012^Creatinine^WinPath||112|UMOLL^UMOLL|66 - 112|A|||F|||||"
	SegmentOBX2 = "OBX|2|ST|testcode^^system||text content||||||F|||||"
	SegmentOBX3 = "OBX|3|DT|testcode^name of testcode^system||20130707||||||F|||||"
	SegmentOBX4 = "OBX|4|TS|testcode^name of testcode^system||20121201120200||||||F|||||"
	SegmentOBX5 = "OBX|5|NM|lpdc-2012^Creatinine^WinPath||[12.0]|UMOLL^UMOLL|66 - 112|A|||F|||||"
	SegmentOBX6 = "OBX|6|ST|testcode^name of testcode^system||text content~||||||F|||||"
	SegmentOBX7 = "OBX|7|NM|lpdc-2012^Creatinine^WinPath||  [12.0]  |UMOLL^UMOLL|66 - 112|A|||F|||||"
	// http://hl7-definition.caristix.com:9010/HL7%20v2.3.1/segment/NTE
	//                 ┌ Source of Comment
	//                 │ ┌ Comment
	//                 │ │                         ┌ Comment Type
	SegmentNTE = "NTE|1|L|NOTE: Submission of serum|"
	// https://corepointhealth.com/resource-center/hl7-resources/hl7-pid-segment
	SegmentPID = "PID|1|8888^^^SF MRN^MRN^|8888^^^SF MRN^MRN~4010232137^^^NHSNBR^NHSNMBR|2345^^^SF MRN^MRN^~5678^^^SF MRN^MRN^|Mogollon^Elisa^^^MR^^CURRENT||20000524|FEMALE|ZZZTEST^PAUL^^^MR^^PREVIOUS||FIRST LINE^SECOND LINE^CITY^^ABC DEF^COUNTRY^HOME^^||0205551234^PRN~^^^something@gmail.com|0205552345^BUSINESS|||CATHOLIC|3393034^^^Encounter Num^FINNBR^||||C|||0||||20180212133000|"
	//                  ┌ Reason ID
	//                  │        ┌ Reason Text
	//                  │        │          ┌ Reason Code System
	SegmentPV2  = "PV2||1|reasonID^reasonText^reasonCodeSystem||||||||||||||||||||^^718004"
	SegmentEVN  = "EVN|A01|20151127180000|||216865551019^ZZZDOCTOR^Foo^Bar Baz^^^^^DRNBR^PRSNL^^^ORGDR^"
	SegmentNK11 = "NK1|1|ZZZTESTWIFE^JANE^^^MS^^CURRENT|SPOUSE|96 The Street^second^LONDON^^ZZ99 1AA^GBR^^^|02011115555||FAMILYMEM||||||||||||||||||||||||||||||"
	SegmentNK12 = "NK1|1|JONES^BOB^^^MR^^CURRENT|BROTHER|96 The Street^second^LONDON^^ZZ99 1AA^GBR^^^|02011115555||FRIEND||||||||||||||||||||||||||||||"
	//                 ┌ Order discipline
	SegmentZCM = "ZCM|1|discipline||||"
)

// SegmentMSH is an arbitrary MSH segment, for when the type does not matter.
var SegmentMSH = fmt.Sprintf(TemplateMSH, "ADT^A01")
