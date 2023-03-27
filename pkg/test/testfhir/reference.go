// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testfhir

import pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"

// GetReferenceID returns the ID embedded in `ref`.
func GetReferenceID(ref *pb.Reference) string {
	for _, v := range []string{
		ref.GetPatientId().GetValue(),
		ref.GetOrganizationId().GetValue(),
		ref.GetPersonId().GetValue(),
		ref.GetPractitionerId().GetValue(),
		ref.GetLocationId().GetValue(),
		ref.GetObservationId().GetValue(),
		ref.GetDiagnosticReportId().GetValue(),
		ref.GetServiceRequestId().GetValue(),
		ref.GetEncounterId().GetValue(),
		ref.GetAllergyIntoleranceId().GetValue(),
		ref.GetDocumentReferenceId().GetValue(),
		ref.GetProcedureId().GetValue(),
		ref.GetConditionId().GetValue(),
	} {
		if v != "" {
			return v
		}
	}
	return ""
}
