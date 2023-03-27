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

package fhircore

import (
	pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
)

func refID(id string) *pb.ReferenceId {
	return &pb.ReferenceId{Value: id}
}

func PatientRef(id string) *pb.Reference {
	return &pb.Reference{Reference: &pb.Reference_PatientId{refID(id)}}
}

func OrganizationRef(id string) *pb.Reference {
	return &pb.Reference{Reference: &pb.Reference_OrganizationId{refID(id)}}
}

func PersonRef(id string) *pb.Reference {
	return &pb.Reference{Reference: &pb.Reference_PersonId{refID(id)}}
}

func PractitionerRef(id string) *pb.Reference {
	return &pb.Reference{Reference: &pb.Reference_PractitionerId{refID(id)}}
}

func LocationRef(id string) *pb.Reference {
	return &pb.Reference{Reference: &pb.Reference_LocationId{refID(id)}}
}

func ObservationRef(id string) *pb.Reference {
	return &pb.Reference{Reference: &pb.Reference_ObservationId{refID(id)}}
}

func DiagnosticReportRef(id string) *pb.Reference {
	return &pb.Reference{Reference: &pb.Reference_DiagnosticReportId{refID(id)}}
}

func ServiceRequestRef(id string) *pb.Reference {
	return &pb.Reference{Reference: &pb.Reference_ServiceRequestId{refID(id)}}
}

func EncounterRef(id string) *pb.Reference {
	return &pb.Reference{Reference: &pb.Reference_EncounterId{refID(id)}}
}

func AllergyIntoleranceRef(id string) *pb.Reference {
	return &pb.Reference{Reference: &pb.Reference_AllergyIntoleranceId{refID(id)}}
}

func DocumentReferenceRef(id string) *pb.Reference {
	return &pb.Reference{Reference: &pb.Reference_DocumentReferenceId{refID(id)}}
}

func ProcedureRef(id string) *pb.Reference {
	return &pb.Reference{Reference: &pb.Reference_ProcedureId{refID(id)}}
}

func ConditionRef(id string) *pb.Reference {
	return &pb.Reference{Reference: &pb.Reference_ConditionId{refID(id)}}
}
