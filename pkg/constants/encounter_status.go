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

package constants

// EncounterStatus is the status of an Encounter.
// Full set of values:
// https://www.hl7.org/fhir/codesystem-encounter-status.html

const (
	// EncounterStatusPlanned denotes that the Encounter has not yet started.
	EncounterStatusPlanned = "planned"
	// EncounterStatusArrived denotes that the patient is present for the Encounter, however is not currently meeting with the practicioner.
	EncounterStatusArrived = "arrived"
	// EncounterStatusInProgress denotes that the Encounter has begun and the patient is present.
	EncounterStatusInProgress = "in-progress"
	// EncounterStatusFinished denotes that the Encounter has ended.
	EncounterStatusFinished = "finished"
	// EncounterStatusCancelled denotes that the Encounter has ended before it has begun.
	EncounterStatusCancelled = "cancelled"
	// EncounterStatusUnknown denotes that the encounter status is unknown.
	EncounterStatusUnknown = "unknown"
)
