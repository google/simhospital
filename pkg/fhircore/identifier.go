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
	cpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/codes_go_proto"
	pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
	"google.golang.org/protobuf/proto"
)

// It also contains functions to deal with Patient and Person identifiers (Medical Record Numbers)
// and UK National Health System identifiers, as well as generic identifiers.

const (
	MRNIdentifierCode    = "MR"
	MRNIdentifierDisplay = "Medical Record Number"
	NHSIdentifierCode    = "NH"
	NHSIdentifierDisplay = "UK National Health System identifier"

	VisitNumberIdentifierCode    = "VN"
	VisitNumberIdentifierDisplay = "Visit number"

	SecondaryMRNIdentifierCode   = "MR_S"
	SecondaryMRIdentifierDisplay = "Medical record number (secondary)"

	IdentifierAssigningAuthorityExt = "IdentifierAssigningAuthority"
	IdentifierAssigningFacilityExt  = "IdentifierAssigningFacility"
	IdentifierTypeCodeExt           = "IdentifierTypeCode"

	// IdentifierTypeCodeSystem represents values in http://hl7.org/fhir/valueset-identifier-type.html
	IdentifierTypeCodeSystem = "http://hl7.org/fhir/ValueSet/identifier-type"
)

// Identifier creates an identifier proto with the given parameters.
func Identifier(value string, code string) *pb.Identifier {
	return &pb.Identifier{
		Use: &pb.Identifier_UseCode{
			Value: cpb.IdentifierUseCode_OFFICIAL,
		},
		Value: String(value),
		Type:  &pb.CodeableConcept{Coding: []*pb.Coding{Coding(code, IdentifierTypeCodeSystem, "")}},
	}
}

func IdentifierNHS(value string) *pb.Identifier {
	i := Identifier(value, NHSIdentifierCode)
	i.Type.Text = String(NHSIdentifierDisplay)
	i.Type.Coding[0].Display = String(NHSIdentifierDisplay)
	return i
}

func IdentifierMRN(value string) *pb.Identifier {
	i := Identifier(value, MRNIdentifierCode)
	i.Type.Coding[0].Display = String(MRNIdentifierDisplay)
	i.Type.Text = String(MRNIdentifierDisplay)
	return i
}

func IdentifierVisitNumber(value string) *pb.Identifier {
	i := Identifier(value, VisitNumberIdentifierCode)
	i.Type.Coding[0].Display = String(VisitNumberIdentifierDisplay)
	i.Type.Text = String(VisitNumberIdentifierDisplay)
	return i
}

// IdentifierSecondaryMRN creates a secondary MRN identifier.
//
// We use all of authority, facility and typeCode, even if they are empty, because MRNs need to
// match all of those in order to be considered equal.
func IdentifierSecondaryMRN(value string, typeCode string, authority string, facility string) *pb.Identifier {
	i := Identifier(value, SecondaryMRNIdentifierCode)
	i.Type.Coding[0].Display = String(SecondaryMRIdentifierDisplay)
	i.Type.Text = String(SecondaryMRIdentifierDisplay)
	addIdentifierExtension(i, IdentifierAssigningAuthorityExt, authority)
	addIdentifierExtension(i, IdentifierAssigningFacilityExt, facility)
	addIdentifierExtension(i, IdentifierTypeCodeExt, typeCode)
	return i
}

func addIdentifierExtension(i *pb.Identifier, extensionURL string, value string) {
	i.Extension = append(i.Extension, StringExtension(extensionURL, value))
}

// identifierMatches returns whether the given identifier matches the given code.
func identifierMatchesCode(id *pb.Identifier, code string) bool {
	for _, c := range id.GetType().GetCoding() {
		if code == c.GetCode().GetValue() {
			return true
		}
	}
	return false
}

// identifierMatches returns whether the given identifier matches the given system.
func identifierMatchesSystem(id *pb.Identifier, system string) bool {
	return system == id.GetSystem().GetValue()
}

// FindIdentifiers returns the list of identifiers that match the given code.
func FindIdentifiers(identifiers []*pb.Identifier, code string) []*pb.Identifier {
	ids := []*pb.Identifier{}
	for _, id := range identifiers {
		if identifierMatchesCode(id, code) {
			ids = append(ids, id)
		}
	}
	return ids
}

// AddOrUpdateIdentifier will either update the provided identifier, if there already is one with
// the same Code, or will insert a new one.
// AddOrUpdateIdentifier returns the updated list of identifiers.
// The entire identifier is updated. If the matched identifier has multiple Coding items, all of those
// will be lost if the provided identifier has a single Coding (which is usually the case).
func AddOrUpdateIdentifier(identifiers []*pb.Identifier, identifier *pb.Identifier) []*pb.Identifier {
	code := ""
	if identifier.Type != nil && len(identifier.Type.Coding) > 0 && identifier.Type.Coding[0].Code != nil {
		code = identifier.Type.Coding[0].Code.Value
	}
	if code == "" {
		return append(identifiers, identifier)
	}
	for i, id := range identifiers {
		if identifierMatchesCode(id, code) {
			// There should only be one identifier for each type, except for
			// secondary MRNs that should match everything.
			if code != SecondaryMRNIdentifierCode || proto.Equal(identifier, id) {
				identifiers[i] = identifier
				return identifiers
			}
		}
	}
	// No existing identifier.
	return append(identifiers, identifier)
}

func GetValue(id *pb.Identifier) string {
	return id.GetValue().GetValue()
}

func extractIDs(identifiers []*pb.Identifier, code string) []string {
	ids := []string{}
	for _, i := range FindIdentifiers(identifiers, code) {
		ids = append(ids, GetValue(i))
	}
	return ids
}

// GetMRN returns the MRNs in the provided list of Identifiers.
func GetMRN(ids []*pb.Identifier) []string {
	return extractIDs(ids, MRNIdentifierCode)
}

// GetNHS returns the NHS numbers in the provided list of Identifiers.
// Generally there will be just one.
func GetNHS(ids []*pb.Identifier) []string {
	return extractIDs(ids, NHSIdentifierCode)
}
