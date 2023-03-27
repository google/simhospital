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

package hl7tofhirmap

import (
	cpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/codes_go_proto"
)

// DefaultAbstractTypeCodeMap maps from string to cpb.AbstractTypeCode_Value.
var DefaultAbstractTypeCodeMap = map[string]cpb.AbstractTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.AbstractTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"ANY":                   cpb.AbstractTypeCode_ANY,                   // Enum 2
	"TYPE":                  cpb.AbstractTypeCode_TYPE,                  // Enum 1
}

// DefaultAccountStatusCodeMap maps from string to cpb.AccountStatusCode_Value.
var DefaultAccountStatusCodeMap = map[string]cpb.AccountStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.AccountStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.AccountStatusCode_ACTIVE,                // Enum 1
	"ENTERED_IN_ERROR":      cpb.AccountStatusCode_ENTERED_IN_ERROR,      // Enum 3
	"INACTIVE":              cpb.AccountStatusCode_INACTIVE,              // Enum 2
	"ON_HOLD":               cpb.AccountStatusCode_ON_HOLD,               // Enum 4
	"UNKNOWN":               cpb.AccountStatusCode_UNKNOWN,               // Enum 5
}

// DefaultActionCardinalityBehaviorCodeMap maps from string to cpb.ActionCardinalityBehaviorCode_Value.
var DefaultActionCardinalityBehaviorCodeMap = map[string]cpb.ActionCardinalityBehaviorCode_Value{
	"INVALID_UNINITIALIZED": cpb.ActionCardinalityBehaviorCode_INVALID_UNINITIALIZED, // Enum 0
	"MULTIPLE":              cpb.ActionCardinalityBehaviorCode_MULTIPLE,              // Enum 2
	"SINGLE":                cpb.ActionCardinalityBehaviorCode_SINGLE,                // Enum 1
}

// DefaultActionConditionKindCodeMap maps from string to cpb.ActionConditionKindCode_Value.
var DefaultActionConditionKindCodeMap = map[string]cpb.ActionConditionKindCode_Value{
	"INVALID_UNINITIALIZED": cpb.ActionConditionKindCode_INVALID_UNINITIALIZED, // Enum 0
	"APPLICABILITY":         cpb.ActionConditionKindCode_APPLICABILITY,         // Enum 1
	"START":                 cpb.ActionConditionKindCode_START,                 // Enum 2
	"STOP":                  cpb.ActionConditionKindCode_STOP,                  // Enum 3
}

// DefaultActionGroupingBehaviorCodeMap maps from string to cpb.ActionGroupingBehaviorCode_Value.
var DefaultActionGroupingBehaviorCodeMap = map[string]cpb.ActionGroupingBehaviorCode_Value{
	"INVALID_UNINITIALIZED": cpb.ActionGroupingBehaviorCode_INVALID_UNINITIALIZED, // Enum 0
	"LOGICAL_GROUP":         cpb.ActionGroupingBehaviorCode_LOGICAL_GROUP,         // Enum 2
	"SENTENCE_GROUP":        cpb.ActionGroupingBehaviorCode_SENTENCE_GROUP,        // Enum 3
	"VISUAL_GROUP":          cpb.ActionGroupingBehaviorCode_VISUAL_GROUP,          // Enum 1
}

// DefaultActionParticipantTypeCodeMap maps from string to cpb.ActionParticipantTypeCode_Value.
var DefaultActionParticipantTypeCodeMap = map[string]cpb.ActionParticipantTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ActionParticipantTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"DEVICE":                cpb.ActionParticipantTypeCode_DEVICE,                // Enum 4
	"PATIENT":               cpb.ActionParticipantTypeCode_PATIENT,               // Enum 1
	"PRACTITIONER":          cpb.ActionParticipantTypeCode_PRACTITIONER,          // Enum 2
	"RELATED_PERSON":        cpb.ActionParticipantTypeCode_RELATED_PERSON,        // Enum 3
}

// DefaultActionPrecheckBehaviorCodeMap maps from string to cpb.ActionPrecheckBehaviorCode_Value.
var DefaultActionPrecheckBehaviorCodeMap = map[string]cpb.ActionPrecheckBehaviorCode_Value{
	"INVALID_UNINITIALIZED": cpb.ActionPrecheckBehaviorCode_INVALID_UNINITIALIZED, // Enum 0
	"NO":                    cpb.ActionPrecheckBehaviorCode_NO,                    // Enum 2
	"YES":                   cpb.ActionPrecheckBehaviorCode_YES,                   // Enum 1
}

// DefaultActionRelationshipTypeCodeMap maps from string to cpb.ActionRelationshipTypeCode_Value.
var DefaultActionRelationshipTypeCodeMap = map[string]cpb.ActionRelationshipTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ActionRelationshipTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"AFTER":                 cpb.ActionRelationshipTypeCode_AFTER,                 // Enum 8
	"AFTER_END":             cpb.ActionRelationshipTypeCode_AFTER_END,             // Enum 9
	"AFTER_START":           cpb.ActionRelationshipTypeCode_AFTER_START,           // Enum 7
	"BEFORE":                cpb.ActionRelationshipTypeCode_BEFORE,                // Enum 2
	"BEFORE_END":            cpb.ActionRelationshipTypeCode_BEFORE_END,            // Enum 3
	"BEFORE_START":          cpb.ActionRelationshipTypeCode_BEFORE_START,          // Enum 1
	"CONCURRENT":            cpb.ActionRelationshipTypeCode_CONCURRENT,            // Enum 5
	"CONCURRENT_WITH_END":   cpb.ActionRelationshipTypeCode_CONCURRENT_WITH_END,   // Enum 6
	"CONCURRENT_WITH_START": cpb.ActionRelationshipTypeCode_CONCURRENT_WITH_START, // Enum 4
}

// DefaultActionRequiredBehaviorCodeMap maps from string to cpb.ActionRequiredBehaviorCode_Value.
var DefaultActionRequiredBehaviorCodeMap = map[string]cpb.ActionRequiredBehaviorCode_Value{
	"INVALID_UNINITIALIZED":  cpb.ActionRequiredBehaviorCode_INVALID_UNINITIALIZED,  // Enum 0
	"COULD":                  cpb.ActionRequiredBehaviorCode_COULD,                  // Enum 2
	"MUST":                   cpb.ActionRequiredBehaviorCode_MUST,                   // Enum 1
	"MUST_UNLESS_DOCUMENTED": cpb.ActionRequiredBehaviorCode_MUST_UNLESS_DOCUMENTED, // Enum 3
}

// DefaultActionSelectionBehaviorCodeMap maps from string to cpb.ActionSelectionBehaviorCode_Value.
var DefaultActionSelectionBehaviorCodeMap = map[string]cpb.ActionSelectionBehaviorCode_Value{
	"INVALID_UNINITIALIZED": cpb.ActionSelectionBehaviorCode_INVALID_UNINITIALIZED, // Enum 0
	"ALL":                   cpb.ActionSelectionBehaviorCode_ALL,                   // Enum 2
	"ALL_OR_NONE":           cpb.ActionSelectionBehaviorCode_ALL_OR_NONE,           // Enum 3
	"ANY":                   cpb.ActionSelectionBehaviorCode_ANY,                   // Enum 1
	"AT_MOST_ONE":           cpb.ActionSelectionBehaviorCode_AT_MOST_ONE,           // Enum 5
	"EXACTLY_ONE":           cpb.ActionSelectionBehaviorCode_EXACTLY_ONE,           // Enum 4
	"ONE_OR_MORE":           cpb.ActionSelectionBehaviorCode_ONE_OR_MORE,           // Enum 6
}

// DefaultAddressTypeCodeMap maps from string to cpb.AddressTypeCode_Value.
var DefaultAddressTypeCodeMap = map[string]cpb.AddressTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.AddressTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"BOTH":                  cpb.AddressTypeCode_BOTH,                  // Enum 3
	"PHYSICAL":              cpb.AddressTypeCode_PHYSICAL,              // Enum 2
	"POSTAL":                cpb.AddressTypeCode_POSTAL,                // Enum 1
}

// DefaultAddressUseCodeMap maps from string to cpb.AddressUseCode_Value.
var DefaultAddressUseCodeMap = map[string]cpb.AddressUseCode_Value{
	"INVALID_UNINITIALIZED": cpb.AddressUseCode_INVALID_UNINITIALIZED, // Enum 0
	"BILLING":               cpb.AddressUseCode_BILLING,               // Enum 5
	"HOME":                  cpb.AddressUseCode_HOME,                  // Enum 1
	"OLD":                   cpb.AddressUseCode_OLD,                   // Enum 4
	"TEMP":                  cpb.AddressUseCode_TEMP,                  // Enum 3
	"WORK":                  cpb.AddressUseCode_WORK,                  // Enum 2
}

// DefaultAdministrativeGenderCodeMap maps from string to cpb.AdministrativeGenderCode_Value.
var DefaultAdministrativeGenderCodeMap = map[string]cpb.AdministrativeGenderCode_Value{
	"INVALID_UNINITIALIZED": cpb.AdministrativeGenderCode_INVALID_UNINITIALIZED, // Enum 0
	"FEMALE":                cpb.AdministrativeGenderCode_FEMALE,                // Enum 2
	"MALE":                  cpb.AdministrativeGenderCode_MALE,                  // Enum 1
	"OTHER":                 cpb.AdministrativeGenderCode_OTHER,                 // Enum 3
	"UNKNOWN":               cpb.AdministrativeGenderCode_UNKNOWN,               // Enum 4
}

// DefaultAdverseEventActualityCodeMap maps from string to cpb.AdverseEventActualityCode_Value.
var DefaultAdverseEventActualityCodeMap = map[string]cpb.AdverseEventActualityCode_Value{
	"INVALID_UNINITIALIZED": cpb.AdverseEventActualityCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTUAL":                cpb.AdverseEventActualityCode_ACTUAL,                // Enum 1
	"POTENTIAL":             cpb.AdverseEventActualityCode_POTENTIAL,             // Enum 2
}

// DefaultAdverseEventOutcomeCodeMap maps from string to cpb.AdverseEventOutcomeCode_Value.
var DefaultAdverseEventOutcomeCodeMap = map[string]cpb.AdverseEventOutcomeCode_Value{
	"INVALID_UNINITIALIZED":  cpb.AdverseEventOutcomeCode_INVALID_UNINITIALIZED,  // Enum 0
	"FATAL":                  cpb.AdverseEventOutcomeCode_FATAL,                  // Enum 5
	"ONGOING":                cpb.AdverseEventOutcomeCode_ONGOING,                // Enum 3
	"RECOVERING":             cpb.AdverseEventOutcomeCode_RECOVERING,             // Enum 2
	"RESOLVED":               cpb.AdverseEventOutcomeCode_RESOLVED,               // Enum 1
	"RESOLVED_WITH_SEQUELAE": cpb.AdverseEventOutcomeCode_RESOLVED_WITH_SEQUELAE, // Enum 4
	"UNKNOWN":                cpb.AdverseEventOutcomeCode_UNKNOWN,                // Enum 6
}

// DefaultAdverseEventSeverityCodeMap maps from string to cpb.AdverseEventSeverityCode_Value.
var DefaultAdverseEventSeverityCodeMap = map[string]cpb.AdverseEventSeverityCode_Value{
	"INVALID_UNINITIALIZED": cpb.AdverseEventSeverityCode_INVALID_UNINITIALIZED, // Enum 0
	"MILD":                  cpb.AdverseEventSeverityCode_MILD,                  // Enum 1
	"MODERATE":              cpb.AdverseEventSeverityCode_MODERATE,              // Enum 2
	"SEVERE":                cpb.AdverseEventSeverityCode_SEVERE,                // Enum 3
}

// DefaultAggregationModeCodeMap maps from string to cpb.AggregationModeCode_Value.
var DefaultAggregationModeCodeMap = map[string]cpb.AggregationModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.AggregationModeCode_INVALID_UNINITIALIZED, // Enum 0
	"BUNDLED":               cpb.AggregationModeCode_BUNDLED,               // Enum 3
	"CONTAINED":             cpb.AggregationModeCode_CONTAINED,             // Enum 1
	"REFERENCED":            cpb.AggregationModeCode_REFERENCED,            // Enum 2
}

// DefaultAllergyIntoleranceCategoryCodeMap maps from string to cpb.AllergyIntoleranceCategoryCode_Value.
var DefaultAllergyIntoleranceCategoryCodeMap = map[string]cpb.AllergyIntoleranceCategoryCode_Value{
	"INVALID_UNINITIALIZED": cpb.AllergyIntoleranceCategoryCode_INVALID_UNINITIALIZED, // Enum 0
	"BIOLOGIC":              cpb.AllergyIntoleranceCategoryCode_BIOLOGIC,              // Enum 4
	"ENVIRONMENT":           cpb.AllergyIntoleranceCategoryCode_ENVIRONMENT,           // Enum 3
	"FOOD":                  cpb.AllergyIntoleranceCategoryCode_FOOD,                  // Enum 1
	"MEDICATION":            cpb.AllergyIntoleranceCategoryCode_MEDICATION,            // Enum 2
}

// DefaultAllergyIntoleranceClinicalStatusCodeMap maps from string to cpb.AllergyIntoleranceClinicalStatusCode_Value.
var DefaultAllergyIntoleranceClinicalStatusCodeMap = map[string]cpb.AllergyIntoleranceClinicalStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.AllergyIntoleranceClinicalStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.AllergyIntoleranceClinicalStatusCode_ACTIVE,                // Enum 1
	"INACTIVE":              cpb.AllergyIntoleranceClinicalStatusCode_INACTIVE,              // Enum 2
	"RESOLVED":              cpb.AllergyIntoleranceClinicalStatusCode_RESOLVED,              // Enum 3
}

// DefaultAllergyIntoleranceCriticalityCodeMap maps from string to cpb.AllergyIntoleranceCriticalityCode_Value.
var DefaultAllergyIntoleranceCriticalityCodeMap = map[string]cpb.AllergyIntoleranceCriticalityCode_Value{
	"INVALID_UNINITIALIZED": cpb.AllergyIntoleranceCriticalityCode_INVALID_UNINITIALIZED, // Enum 0
	"HIGH":                  cpb.AllergyIntoleranceCriticalityCode_HIGH,                  // Enum 2
	"LOW":                   cpb.AllergyIntoleranceCriticalityCode_LOW,                   // Enum 1
	"UNABLE_TO_ASSESS":      cpb.AllergyIntoleranceCriticalityCode_UNABLE_TO_ASSESS,      // Enum 3
}

// DefaultAllergyIntoleranceSeverityCodeMap maps from string to cpb.AllergyIntoleranceSeverityCode_Value.
var DefaultAllergyIntoleranceSeverityCodeMap = map[string]cpb.AllergyIntoleranceSeverityCode_Value{
	"INVALID_UNINITIALIZED": cpb.AllergyIntoleranceSeverityCode_INVALID_UNINITIALIZED, // Enum 0
	"MILD":                  cpb.AllergyIntoleranceSeverityCode_MILD,                  // Enum 1
	"MODERATE":              cpb.AllergyIntoleranceSeverityCode_MODERATE,              // Enum 2
	"SEVERE":                cpb.AllergyIntoleranceSeverityCode_SEVERE,                // Enum 3
}

// DefaultAllergyIntoleranceSubstanceExposureRiskCodeMap maps from string to cpb.AllergyIntoleranceSubstanceExposureRiskCode_Value.
var DefaultAllergyIntoleranceSubstanceExposureRiskCodeMap = map[string]cpb.AllergyIntoleranceSubstanceExposureRiskCode_Value{
	"INVALID_UNINITIALIZED":  cpb.AllergyIntoleranceSubstanceExposureRiskCode_INVALID_UNINITIALIZED,  // Enum 0
	"KNOWN_REACTION_RISK":    cpb.AllergyIntoleranceSubstanceExposureRiskCode_KNOWN_REACTION_RISK,    // Enum 1
	"NO_KNOWN_REACTION_RISK": cpb.AllergyIntoleranceSubstanceExposureRiskCode_NO_KNOWN_REACTION_RISK, // Enum 2
}

// DefaultAllergyIntoleranceTypeCodeMap maps from string to cpb.AllergyIntoleranceTypeCode_Value.
var DefaultAllergyIntoleranceTypeCodeMap = map[string]cpb.AllergyIntoleranceTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.AllergyIntoleranceTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"ALLERGY":               cpb.AllergyIntoleranceTypeCode_ALLERGY,               // Enum 1
	"INTOLERANCE":           cpb.AllergyIntoleranceTypeCode_INTOLERANCE,           // Enum 2
}

// DefaultAllergyIntoleranceVerificationStatusCodeMap maps from string to cpb.AllergyIntoleranceVerificationStatusCode_Value.
var DefaultAllergyIntoleranceVerificationStatusCodeMap = map[string]cpb.AllergyIntoleranceVerificationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.AllergyIntoleranceVerificationStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"CONFIRMED":             cpb.AllergyIntoleranceVerificationStatusCode_CONFIRMED,             // Enum 2
	"ENTERED_IN_ERROR":      cpb.AllergyIntoleranceVerificationStatusCode_ENTERED_IN_ERROR,      // Enum 4
	"REFUTED":               cpb.AllergyIntoleranceVerificationStatusCode_REFUTED,               // Enum 3
	"UNCONFIRMED":           cpb.AllergyIntoleranceVerificationStatusCode_UNCONFIRMED,           // Enum 1
}

// DefaultAppointmentStatusCodeMap maps from string to cpb.AppointmentStatusCode_Value.
var DefaultAppointmentStatusCodeMap = map[string]cpb.AppointmentStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.AppointmentStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ARRIVED":               cpb.AppointmentStatusCode_ARRIVED,               // Enum 4
	"BOOKED":                cpb.AppointmentStatusCode_BOOKED,                // Enum 3
	"CANCELLED":             cpb.AppointmentStatusCode_CANCELLED,             // Enum 6
	"CHECKED_IN":            cpb.AppointmentStatusCode_CHECKED_IN,            // Enum 9
	"ENTERED_IN_ERROR":      cpb.AppointmentStatusCode_ENTERED_IN_ERROR,      // Enum 8
	"FULFILLED":             cpb.AppointmentStatusCode_FULFILLED,             // Enum 5
	"NOSHOW":                cpb.AppointmentStatusCode_NOSHOW,                // Enum 7
	"PENDING":               cpb.AppointmentStatusCode_PENDING,               // Enum 2
	"PROPOSED":              cpb.AppointmentStatusCode_PROPOSED,              // Enum 1
	"WAITLIST":              cpb.AppointmentStatusCode_WAITLIST,              // Enum 10
}

// DefaultAssertionDirectionTypeCodeMap maps from string to cpb.AssertionDirectionTypeCode_Value.
var DefaultAssertionDirectionTypeCodeMap = map[string]cpb.AssertionDirectionTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.AssertionDirectionTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"REQUEST":               cpb.AssertionDirectionTypeCode_REQUEST,               // Enum 2
	"RESPONSE":              cpb.AssertionDirectionTypeCode_RESPONSE,              // Enum 1
}

// DefaultAssertionOperatorTypeCodeMap maps from string to cpb.AssertionOperatorTypeCode_Value.
var DefaultAssertionOperatorTypeCodeMap = map[string]cpb.AssertionOperatorTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.AssertionOperatorTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"<":                     cpb.AssertionOperatorTypeCode_LESS_THAN,             // Enum 6
	"=":                     cpb.AssertionOperatorTypeCode_EQUALS,                // Enum 1
	">":                     cpb.AssertionOperatorTypeCode_GREATER_THAN,          // Enum 5
	"CONTAINS":              cpb.AssertionOperatorTypeCode_CONTAINS,              // Enum 9
	"EMPTY":                 cpb.AssertionOperatorTypeCode_EMPTY,                 // Enum 7
	"EQUALS":                cpb.AssertionOperatorTypeCode_EQUALS,                // Enum 1
	"EVAL":                  cpb.AssertionOperatorTypeCode_EVAL,                  // Enum 11
	"GREATER_THAN":          cpb.AssertionOperatorTypeCode_GREATER_THAN,          // Enum 5
	"IN":                    cpb.AssertionOperatorTypeCode_IN,                    // Enum 3
	"LESS_THAN":             cpb.AssertionOperatorTypeCode_LESS_THAN,             // Enum 6
	"NOT_CONTAINS":          cpb.AssertionOperatorTypeCode_NOT_CONTAINS,          // Enum 10
	"NOT_EMPTY":             cpb.AssertionOperatorTypeCode_NOT_EMPTY,             // Enum 8
	"NOT_EQUALS":            cpb.AssertionOperatorTypeCode_NOT_EQUALS,            // Enum 2
	"NOT_IN":                cpb.AssertionOperatorTypeCode_NOT_IN,                // Enum 4
}

// DefaultAssertionResponseTypesCodeMap maps from string to cpb.AssertionResponseTypesCode_Value.
var DefaultAssertionResponseTypesCodeMap = map[string]cpb.AssertionResponseTypesCode_Value{
	"INVALID_UNINITIALIZED": cpb.AssertionResponseTypesCode_INVALID_UNINITIALIZED, // Enum 0
	"BAD":                   cpb.AssertionResponseTypesCode_BAD,                   // Enum 5
	"CONFLICT":              cpb.AssertionResponseTypesCode_CONFLICT,              // Enum 9
	"CREATED":               cpb.AssertionResponseTypesCode_CREATED,               // Enum 2
	"FORBIDDEN":             cpb.AssertionResponseTypesCode_FORBIDDEN,             // Enum 6
	"GONE":                  cpb.AssertionResponseTypesCode_GONE,                  // Enum 10
	"METHOD_NOT_ALLOWED":    cpb.AssertionResponseTypesCode_METHOD_NOT_ALLOWED,    // Enum 8
	"NOT_FOUND":             cpb.AssertionResponseTypesCode_NOT_FOUND,             // Enum 7
	"NOT_MODIFIED":          cpb.AssertionResponseTypesCode_NOT_MODIFIED,          // Enum 4
	"NO_CONTENT":            cpb.AssertionResponseTypesCode_NO_CONTENT,            // Enum 3
	"OKAY":                  cpb.AssertionResponseTypesCode_OKAY,                  // Enum 1
	"PRECONDITION_FAILED":   cpb.AssertionResponseTypesCode_PRECONDITION_FAILED,   // Enum 11
	"UNPROCESSABLE":         cpb.AssertionResponseTypesCode_UNPROCESSABLE,         // Enum 12
}

// DefaultAuditEventActionCodeMap maps from string to cpb.AuditEventActionCode_Value.
var DefaultAuditEventActionCodeMap = map[string]cpb.AuditEventActionCode_Value{
	"INVALID_UNINITIALIZED": cpb.AuditEventActionCode_INVALID_UNINITIALIZED, // Enum 0
	"C":                     cpb.AuditEventActionCode_C,                     // Enum 1
	"D":                     cpb.AuditEventActionCode_D,                     // Enum 4
	"E":                     cpb.AuditEventActionCode_E,                     // Enum 5
	"R":                     cpb.AuditEventActionCode_R,                     // Enum 2
	"U":                     cpb.AuditEventActionCode_U,                     // Enum 3
}

// DefaultAuditEventAgentNetworkTypeCodeMap maps from string to cpb.AuditEventAgentNetworkTypeCode_Value.
var DefaultAuditEventAgentNetworkTypeCodeMap = map[string]cpb.AuditEventAgentNetworkTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.AuditEventAgentNetworkTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"EMAIL_ADDRESS":         cpb.AuditEventAgentNetworkTypeCode_EMAIL_ADDRESS,         // Enum 4
	"IP_ADDRESS":            cpb.AuditEventAgentNetworkTypeCode_IP_ADDRESS,            // Enum 2
	"MACHINE_NAME":          cpb.AuditEventAgentNetworkTypeCode_MACHINE_NAME,          // Enum 1
	"TELEPHONE_NUMBER":      cpb.AuditEventAgentNetworkTypeCode_TELEPHONE_NUMBER,      // Enum 3
	"URI":                   cpb.AuditEventAgentNetworkTypeCode_URI,                   // Enum 5
}

// DefaultAuditEventOutcomeCodeMap maps from string to cpb.AuditEventOutcomeCode_Value.
var DefaultAuditEventOutcomeCodeMap = map[string]cpb.AuditEventOutcomeCode_Value{
	"INVALID_UNINITIALIZED": cpb.AuditEventOutcomeCode_INVALID_UNINITIALIZED, // Enum 0
	"MAJOR_FAILURE":         cpb.AuditEventOutcomeCode_MAJOR_FAILURE,         // Enum 4
	"MINOR_FAILURE":         cpb.AuditEventOutcomeCode_MINOR_FAILURE,         // Enum 2
	"SERIOUS_FAILURE":       cpb.AuditEventOutcomeCode_SERIOUS_FAILURE,       // Enum 3
	"SUCCESS":               cpb.AuditEventOutcomeCode_SUCCESS,               // Enum 1
}

// DefaultBenefitCostApplicabilityCodeMap maps from string to cpb.BenefitCostApplicabilityCode_Value.
var DefaultBenefitCostApplicabilityCodeMap = map[string]cpb.BenefitCostApplicabilityCode_Value{
	"INVALID_UNINITIALIZED": cpb.BenefitCostApplicabilityCode_INVALID_UNINITIALIZED, // Enum 0
	"IN_NETWORK":            cpb.BenefitCostApplicabilityCode_IN_NETWORK,            // Enum 1
	"OTHER":                 cpb.BenefitCostApplicabilityCode_OTHER,                 // Enum 3
	"OUT_OF_NETWORK":        cpb.BenefitCostApplicabilityCode_OUT_OF_NETWORK,        // Enum 2
}

// DefaultBindingStrengthCodeMap maps from string to cpb.BindingStrengthCode_Value.
var DefaultBindingStrengthCodeMap = map[string]cpb.BindingStrengthCode_Value{
	"INVALID_UNINITIALIZED": cpb.BindingStrengthCode_INVALID_UNINITIALIZED, // Enum 0
	"EXAMPLE":               cpb.BindingStrengthCode_EXAMPLE,               // Enum 4
	"EXTENSIBLE":            cpb.BindingStrengthCode_EXTENSIBLE,            // Enum 2
	"PREFERRED":             cpb.BindingStrengthCode_PREFERRED,             // Enum 3
	"REQUIRED":              cpb.BindingStrengthCode_REQUIRED,              // Enum 1
}

// DefaultBiologicallyDerivedProductCategoryCodeMap maps from string to cpb.BiologicallyDerivedProductCategoryCode_Value.
var DefaultBiologicallyDerivedProductCategoryCodeMap = map[string]cpb.BiologicallyDerivedProductCategoryCode_Value{
	"INVALID_UNINITIALIZED": cpb.BiologicallyDerivedProductCategoryCode_INVALID_UNINITIALIZED, // Enum 0
	"BIOLOGICAL_AGENT":      cpb.BiologicallyDerivedProductCategoryCode_BIOLOGICAL_AGENT,      // Enum 5
	"CELLS":                 cpb.BiologicallyDerivedProductCategoryCode_CELLS,                 // Enum 4
	"FLUID":                 cpb.BiologicallyDerivedProductCategoryCode_FLUID,                 // Enum 3
	"ORGAN":                 cpb.BiologicallyDerivedProductCategoryCode_ORGAN,                 // Enum 1
	"TISSUE":                cpb.BiologicallyDerivedProductCategoryCode_TISSUE,                // Enum 2
}

// DefaultBiologicallyDerivedProductStatusCodeMap maps from string to cpb.BiologicallyDerivedProductStatusCode_Value.
var DefaultBiologicallyDerivedProductStatusCodeMap = map[string]cpb.BiologicallyDerivedProductStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.BiologicallyDerivedProductStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"AVAILABLE":             cpb.BiologicallyDerivedProductStatusCode_AVAILABLE,             // Enum 1
	"UNAVAILABLE":           cpb.BiologicallyDerivedProductStatusCode_UNAVAILABLE,           // Enum 2
}

// DefaultBiologicallyDerivedProductStorageScaleCodeMap maps from string to cpb.BiologicallyDerivedProductStorageScaleCode_Value.
var DefaultBiologicallyDerivedProductStorageScaleCodeMap = map[string]cpb.BiologicallyDerivedProductStorageScaleCode_Value{
	"INVALID_UNINITIALIZED": cpb.BiologicallyDerivedProductStorageScaleCode_INVALID_UNINITIALIZED, // Enum 0
	"CELSIUS":               cpb.BiologicallyDerivedProductStorageScaleCode_CELSIUS,               // Enum 2
	"FARENHEIT":             cpb.BiologicallyDerivedProductStorageScaleCode_FARENHEIT,             // Enum 1
	"KELVIN":                cpb.BiologicallyDerivedProductStorageScaleCode_KELVIN,                // Enum 3
}

// DefaultBundleTypeCodeMap maps from string to cpb.BundleTypeCode_Value.
var DefaultBundleTypeCodeMap = map[string]cpb.BundleTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.BundleTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"BATCH":                 cpb.BundleTypeCode_BATCH,                 // Enum 5
	"BATCH_RESPONSE":        cpb.BundleTypeCode_BATCH_RESPONSE,        // Enum 6
	"COLLECTION":            cpb.BundleTypeCode_COLLECTION,            // Enum 9
	"DOCUMENT":              cpb.BundleTypeCode_DOCUMENT,              // Enum 1
	"HISTORY":               cpb.BundleTypeCode_HISTORY,               // Enum 7
	"MESSAGE":               cpb.BundleTypeCode_MESSAGE,               // Enum 2
	"SEARCHSET":             cpb.BundleTypeCode_SEARCHSET,             // Enum 8
	"TRANSACTION":           cpb.BundleTypeCode_TRANSACTION,           // Enum 3
	"TRANSACTION_RESPONSE":  cpb.BundleTypeCode_TRANSACTION_RESPONSE,  // Enum 4
}

// DefaultCanonicalStatusCodesForFHIRResourcesCodeMap maps from string to cpb.CanonicalStatusCodesForFHIRResourcesCode_Value.
var DefaultCanonicalStatusCodesForFHIRResourcesCodeMap = map[string]cpb.CanonicalStatusCodesForFHIRResourcesCode_Value{
	"INVALID_UNINITIALIZED": cpb.CanonicalStatusCodesForFHIRResourcesCode_INVALID_UNINITIALIZED, // Enum 0
	"ABANDONED":             cpb.CanonicalStatusCodesForFHIRResourcesCode_ABANDONED,             // Enum 16
	"ACCEPTED":              cpb.CanonicalStatusCodesForFHIRResourcesCode_ACCEPTED,              // Enum 8
	"ACTIVE":                cpb.CanonicalStatusCodesForFHIRResourcesCode_ACTIVE,                // Enum 10
	"AHEAD_OF_TARGET":       cpb.CanonicalStatusCodesForFHIRResourcesCode_AHEAD_OF_TARGET,       // Enum 27
	"ARRIVED":               cpb.CanonicalStatusCodesForFHIRResourcesCode_ARRIVED,               // Enum 9
	"BEHIND_TARGET":         cpb.CanonicalStatusCodesForFHIRResourcesCode_BEHIND_TARGET,         // Enum 28
	"BUSY_UNAVAILABLE":      cpb.CanonicalStatusCodesForFHIRResourcesCode_BUSY_UNAVAILABLE,      // Enum 24
	"COMPLETE":              cpb.CanonicalStatusCodesForFHIRResourcesCode_COMPLETE,              // Enum 14
	"CONFIRMED":             cpb.CanonicalStatusCodesForFHIRResourcesCode_CONFIRMED,             // Enum 19
	"DECLINED":              cpb.CanonicalStatusCodesForFHIRResourcesCode_DECLINED,              // Enum 7
	"DIFFERENTIAL":          cpb.CanonicalStatusCodesForFHIRResourcesCode_DIFFERENTIAL,          // Enum 22
	"DRAFT":                 cpb.CanonicalStatusCodesForFHIRResourcesCode_DRAFT,                 // Enum 4
	"ERROR":                 cpb.CanonicalStatusCodesForFHIRResourcesCode_ERROR,                 // Enum 1
	"FAILED":                cpb.CanonicalStatusCodesForFHIRResourcesCode_FAILED,                // Enum 12
	"FREE":                  cpb.CanonicalStatusCodesForFHIRResourcesCode_FREE,                  // Enum 25
	"HW_DISCON":             cpb.CanonicalStatusCodesForFHIRResourcesCode_HW_DISCON,             // Enum 31
	"INACTIVE":              cpb.CanonicalStatusCodesForFHIRResourcesCode_INACTIVE,              // Enum 15
	"NOT_READY":             cpb.CanonicalStatusCodesForFHIRResourcesCode_NOT_READY,             // Enum 29
	"ON_TARGET":             cpb.CanonicalStatusCodesForFHIRResourcesCode_ON_TARGET,             // Enum 26
	"PARTIAL":               cpb.CanonicalStatusCodesForFHIRResourcesCode_PARTIAL,               // Enum 23
	"PLANNED":               cpb.CanonicalStatusCodesForFHIRResourcesCode_PLANNED,               // Enum 3
	"PROPOSED":              cpb.CanonicalStatusCodesForFHIRResourcesCode_PROPOSED,              // Enum 2
	"RECEIVED":              cpb.CanonicalStatusCodesForFHIRResourcesCode_RECEIVED,              // Enum 6
	"REFUTED":               cpb.CanonicalStatusCodesForFHIRResourcesCode_REFUTED,               // Enum 21
	"REPLACED":              cpb.CanonicalStatusCodesForFHIRResourcesCode_REPLACED,              // Enum 13
	"REQUESTED":             cpb.CanonicalStatusCodesForFHIRResourcesCode_REQUESTED,             // Enum 5
	"RESOLVED":              cpb.CanonicalStatusCodesForFHIRResourcesCode_RESOLVED,              // Enum 20
	"SUSPENDED":             cpb.CanonicalStatusCodesForFHIRResourcesCode_SUSPENDED,             // Enum 11
	"TRANSDUC_DISCON":       cpb.CanonicalStatusCodesForFHIRResourcesCode_TRANSDUC_DISCON,       // Enum 30
	"UNCONFIRMED":           cpb.CanonicalStatusCodesForFHIRResourcesCode_UNCONFIRMED,           // Enum 18
	"UNKNOWN":               cpb.CanonicalStatusCodesForFHIRResourcesCode_UNKNOWN,               // Enum 17
}

// DefaultCapabilityStatementKindCodeMap maps from string to cpb.CapabilityStatementKindCode_Value.
var DefaultCapabilityStatementKindCodeMap = map[string]cpb.CapabilityStatementKindCode_Value{
	"INVALID_UNINITIALIZED": cpb.CapabilityStatementKindCode_INVALID_UNINITIALIZED, // Enum 0
	"CAPABILITY":            cpb.CapabilityStatementKindCode_CAPABILITY,            // Enum 2
	"INSTANCE":              cpb.CapabilityStatementKindCode_INSTANCE,              // Enum 1
	"REQUIREMENTS":          cpb.CapabilityStatementKindCode_REQUIREMENTS,          // Enum 3
}

// DefaultCarePlanActivityStatusCodeMap maps from string to cpb.CarePlanActivityStatusCode_Value.
var DefaultCarePlanActivityStatusCodeMap = map[string]cpb.CarePlanActivityStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.CarePlanActivityStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"CANCELLED":             cpb.CarePlanActivityStatusCode_CANCELLED,             // Enum 6
	"COMPLETED":             cpb.CarePlanActivityStatusCode_COMPLETED,             // Enum 5
	"ENTERED_IN_ERROR":      cpb.CarePlanActivityStatusCode_ENTERED_IN_ERROR,      // Enum 9
	"IN_PROGRESS":           cpb.CarePlanActivityStatusCode_IN_PROGRESS,           // Enum 3
	"NOT_STARTED":           cpb.CarePlanActivityStatusCode_NOT_STARTED,           // Enum 1
	"ON_HOLD":               cpb.CarePlanActivityStatusCode_ON_HOLD,               // Enum 4
	"SCHEDULED":             cpb.CarePlanActivityStatusCode_SCHEDULED,             // Enum 2
	"STOPPED":               cpb.CarePlanActivityStatusCode_STOPPED,               // Enum 7
	"UNKNOWN":               cpb.CarePlanActivityStatusCode_UNKNOWN,               // Enum 8
}

// DefaultCareTeamStatusCodeMap maps from string to cpb.CareTeamStatusCode_Value.
var DefaultCareTeamStatusCodeMap = map[string]cpb.CareTeamStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.CareTeamStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.CareTeamStatusCode_ACTIVE,                // Enum 2
	"ENTERED_IN_ERROR":      cpb.CareTeamStatusCode_ENTERED_IN_ERROR,      // Enum 5
	"INACTIVE":              cpb.CareTeamStatusCode_INACTIVE,              // Enum 4
	"PROPOSED":              cpb.CareTeamStatusCode_PROPOSED,              // Enum 1
	"SUSPENDED":             cpb.CareTeamStatusCode_SUSPENDED,             // Enum 3
}

// DefaultCatalogEntryRelationTypeCodeMap maps from string to cpb.CatalogEntryRelationTypeCode_Value.
var DefaultCatalogEntryRelationTypeCodeMap = map[string]cpb.CatalogEntryRelationTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.CatalogEntryRelationTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"IS_REPLACED_BY":        cpb.CatalogEntryRelationTypeCode_IS_REPLACED_BY,        // Enum 2
	"TRIGGERS":              cpb.CatalogEntryRelationTypeCode_TRIGGERS,              // Enum 1
}

// DefaultChargeItemStatusCodeMap maps from string to cpb.ChargeItemStatusCode_Value.
var DefaultChargeItemStatusCodeMap = map[string]cpb.ChargeItemStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ChargeItemStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ABORTED":               cpb.ChargeItemStatusCode_ABORTED,               // Enum 4
	"BILLABLE":              cpb.ChargeItemStatusCode_BILLABLE,              // Enum 2
	"BILLED":                cpb.ChargeItemStatusCode_BILLED,                // Enum 5
	"ENTERED_IN_ERROR":      cpb.ChargeItemStatusCode_ENTERED_IN_ERROR,      // Enum 6
	"NOT_BILLABLE":          cpb.ChargeItemStatusCode_NOT_BILLABLE,          // Enum 3
	"PLANNED":               cpb.ChargeItemStatusCode_PLANNED,               // Enum 1
	"UNKNOWN":               cpb.ChargeItemStatusCode_UNKNOWN,               // Enum 7
}

// DefaultChoiceListOrientationCodeMap maps from string to cpb.ChoiceListOrientationCode_Value.
var DefaultChoiceListOrientationCodeMap = map[string]cpb.ChoiceListOrientationCode_Value{
	"INVALID_UNINITIALIZED": cpb.ChoiceListOrientationCode_INVALID_UNINITIALIZED, // Enum 0
	"HORIZONTAL":            cpb.ChoiceListOrientationCode_HORIZONTAL,            // Enum 1
	"VERTICAL":              cpb.ChoiceListOrientationCode_VERTICAL,              // Enum 2
}

// DefaultClaimProcessingCodeMap maps from string to cpb.ClaimProcessingCode_Value.
var DefaultClaimProcessingCodeMap = map[string]cpb.ClaimProcessingCode_Value{
	"INVALID_UNINITIALIZED": cpb.ClaimProcessingCode_INVALID_UNINITIALIZED, // Enum 0
	"COMPLETE":              cpb.ClaimProcessingCode_COMPLETE,              // Enum 2
	"ERROR":                 cpb.ClaimProcessingCode_ERROR,                 // Enum 3
	"PARTIAL":               cpb.ClaimProcessingCode_PARTIAL,               // Enum 4
	"QUEUED":                cpb.ClaimProcessingCode_QUEUED,                // Enum 1
}

// DefaultCodeSearchSupportCodeMap maps from string to cpb.CodeSearchSupportCode_Value.
var DefaultCodeSearchSupportCodeMap = map[string]cpb.CodeSearchSupportCode_Value{
	"INVALID_UNINITIALIZED": cpb.CodeSearchSupportCode_INVALID_UNINITIALIZED, // Enum 0
	"ALL":                   cpb.CodeSearchSupportCode_ALL,                   // Enum 2
	"EXPLICIT":              cpb.CodeSearchSupportCode_EXPLICIT,              // Enum 1
}

// DefaultCodeSystemContentModeCodeMap maps from string to cpb.CodeSystemContentModeCode_Value.
var DefaultCodeSystemContentModeCodeMap = map[string]cpb.CodeSystemContentModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.CodeSystemContentModeCode_INVALID_UNINITIALIZED, // Enum 0
	"COMPLETE":              cpb.CodeSystemContentModeCode_COMPLETE,              // Enum 4
	"EXAMPLE":               cpb.CodeSystemContentModeCode_EXAMPLE,               // Enum 2
	"FRAGMENT":              cpb.CodeSystemContentModeCode_FRAGMENT,              // Enum 3
	"NOT_PRESENT":           cpb.CodeSystemContentModeCode_NOT_PRESENT,           // Enum 1
	"SUPPLEMENT":            cpb.CodeSystemContentModeCode_SUPPLEMENT,            // Enum 5
}

// DefaultCodeSystemHierarchyMeaningCodeMap maps from string to cpb.CodeSystemHierarchyMeaningCode_Value.
var DefaultCodeSystemHierarchyMeaningCodeMap = map[string]cpb.CodeSystemHierarchyMeaningCode_Value{
	"INVALID_UNINITIALIZED": cpb.CodeSystemHierarchyMeaningCode_INVALID_UNINITIALIZED, // Enum 0
	"CLASSIFIED_WITH":       cpb.CodeSystemHierarchyMeaningCode_CLASSIFIED_WITH,       // Enum 4
	"GROUPED_BY":            cpb.CodeSystemHierarchyMeaningCode_GROUPED_BY,            // Enum 1
	"IS_A":                  cpb.CodeSystemHierarchyMeaningCode_IS_A,                  // Enum 2
	"PART_OF":               cpb.CodeSystemHierarchyMeaningCode_PART_OF,               // Enum 3
}

// DefaultCompartmentTypeCodeMap maps from string to cpb.CompartmentTypeCode_Value.
var DefaultCompartmentTypeCodeMap = map[string]cpb.CompartmentTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.CompartmentTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"DEVICE":                cpb.CompartmentTypeCode_DEVICE,                // Enum 5
	"ENCOUNTER":             cpb.CompartmentTypeCode_ENCOUNTER,             // Enum 2
	"PATIENT":               cpb.CompartmentTypeCode_PATIENT,               // Enum 1
	"PRACTITIONER":          cpb.CompartmentTypeCode_PRACTITIONER,          // Enum 4
	"RELATED_PERSON":        cpb.CompartmentTypeCode_RELATED_PERSON,        // Enum 3
}

// DefaultCompositionAttestationModeCodeMap maps from string to cpb.CompositionAttestationModeCode_Value.
var DefaultCompositionAttestationModeCodeMap = map[string]cpb.CompositionAttestationModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.CompositionAttestationModeCode_INVALID_UNINITIALIZED, // Enum 0
	"LEGAL":                 cpb.CompositionAttestationModeCode_LEGAL,                 // Enum 3
	"OFFICIAL":              cpb.CompositionAttestationModeCode_OFFICIAL,              // Enum 4
	"PERSONAL":              cpb.CompositionAttestationModeCode_PERSONAL,              // Enum 1
	"PROFESSIONAL":          cpb.CompositionAttestationModeCode_PROFESSIONAL,          // Enum 2
}

// DefaultCompositionStatusCodeMap maps from string to cpb.CompositionStatusCode_Value.
var DefaultCompositionStatusCodeMap = map[string]cpb.CompositionStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.CompositionStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"AMENDED":               cpb.CompositionStatusCode_AMENDED,               // Enum 3
	"ENTERED_IN_ERROR":      cpb.CompositionStatusCode_ENTERED_IN_ERROR,      // Enum 4
	"FINAL":                 cpb.CompositionStatusCode_FINAL,                 // Enum 2
	"PRELIMINARY":           cpb.CompositionStatusCode_PRELIMINARY,           // Enum 1
}

// DefaultConceptMapEquivalenceCodeMap maps from string to cpb.ConceptMapEquivalenceCode_Value.
var DefaultConceptMapEquivalenceCodeMap = map[string]cpb.ConceptMapEquivalenceCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConceptMapEquivalenceCode_INVALID_UNINITIALIZED, // Enum 0
	"DISJOINT":              cpb.ConceptMapEquivalenceCode_DISJOINT,              // Enum 10
	"EQUAL":                 cpb.ConceptMapEquivalenceCode_EQUAL,                 // Enum 3
	"EQUIVALENT":            cpb.ConceptMapEquivalenceCode_EQUIVALENT,            // Enum 2
	"INEXACT":               cpb.ConceptMapEquivalenceCode_INEXACT,               // Enum 8
	"NARROWER":              cpb.ConceptMapEquivalenceCode_NARROWER,              // Enum 6
	"RELATEDTO":             cpb.ConceptMapEquivalenceCode_RELATEDTO,             // Enum 1
	"SPECIALIZES":           cpb.ConceptMapEquivalenceCode_SPECIALIZES,           // Enum 7
	"SUBSUMES":              cpb.ConceptMapEquivalenceCode_SUBSUMES,              // Enum 5
	"UNMATCHED":             cpb.ConceptMapEquivalenceCode_UNMATCHED,             // Enum 9
	"WIDER":                 cpb.ConceptMapEquivalenceCode_WIDER,                 // Enum 4
}

// DefaultConceptMapGroupUnmappedModeCodeMap maps from string to cpb.ConceptMapGroupUnmappedModeCode_Value.
var DefaultConceptMapGroupUnmappedModeCodeMap = map[string]cpb.ConceptMapGroupUnmappedModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConceptMapGroupUnmappedModeCode_INVALID_UNINITIALIZED, // Enum 0
	"FIXED":                 cpb.ConceptMapGroupUnmappedModeCode_FIXED,                 // Enum 2
	"OTHER_MAP":             cpb.ConceptMapGroupUnmappedModeCode_OTHER_MAP,             // Enum 3
	"PROVIDED":              cpb.ConceptMapGroupUnmappedModeCode_PROVIDED,              // Enum 1
}

// DefaultConditionClinicalStatusCodeMap maps from string to cpb.ConditionClinicalStatusCode_Value.
var DefaultConditionClinicalStatusCodeMap = map[string]cpb.ConditionClinicalStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConditionClinicalStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.ConditionClinicalStatusCode_ACTIVE,                // Enum 1
	"INACTIVE":              cpb.ConditionClinicalStatusCode_INACTIVE,              // Enum 4
	"RECURRENCE":            cpb.ConditionClinicalStatusCode_RECURRENCE,            // Enum 2
	"RELAPSE":               cpb.ConditionClinicalStatusCode_RELAPSE,               // Enum 3
	"REMISSION":             cpb.ConditionClinicalStatusCode_REMISSION,             // Enum 5
	"RESOLVED":              cpb.ConditionClinicalStatusCode_RESOLVED,              // Enum 6
}

// DefaultConditionVerificationStatusCodeMap maps from string to cpb.ConditionVerificationStatusCode_Value.
var DefaultConditionVerificationStatusCodeMap = map[string]cpb.ConditionVerificationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConditionVerificationStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"CONFIRMED":             cpb.ConditionVerificationStatusCode_CONFIRMED,             // Enum 4
	"DIFFERENTIAL":          cpb.ConditionVerificationStatusCode_DIFFERENTIAL,          // Enum 3
	"ENTERED_IN_ERROR":      cpb.ConditionVerificationStatusCode_ENTERED_IN_ERROR,      // Enum 6
	"PROVISIONAL":           cpb.ConditionVerificationStatusCode_PROVISIONAL,           // Enum 2
	"REFUTED":               cpb.ConditionVerificationStatusCode_REFUTED,               // Enum 5
	"UNCONFIRMED":           cpb.ConditionVerificationStatusCode_UNCONFIRMED,           // Enum 1
}

// DefaultConditionalDeleteStatusCodeMap maps from string to cpb.ConditionalDeleteStatusCode_Value.
var DefaultConditionalDeleteStatusCodeMap = map[string]cpb.ConditionalDeleteStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConditionalDeleteStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"MULTIPLE":              cpb.ConditionalDeleteStatusCode_MULTIPLE,              // Enum 3
	"NOT_SUPPORTED":         cpb.ConditionalDeleteStatusCode_NOT_SUPPORTED,         // Enum 1
	"SINGLE":                cpb.ConditionalDeleteStatusCode_SINGLE,                // Enum 2
}

// DefaultConditionalReadStatusCodeMap maps from string to cpb.ConditionalReadStatusCode_Value.
var DefaultConditionalReadStatusCodeMap = map[string]cpb.ConditionalReadStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConditionalReadStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"FULL_SUPPORT":          cpb.ConditionalReadStatusCode_FULL_SUPPORT,          // Enum 4
	"MODIFIED_SINCE":        cpb.ConditionalReadStatusCode_MODIFIED_SINCE,        // Enum 2
	"NOT_MATCH":             cpb.ConditionalReadStatusCode_NOT_MATCH,             // Enum 3
	"NOT_SUPPORTED":         cpb.ConditionalReadStatusCode_NOT_SUPPORTED,         // Enum 1
}

// DefaultConformanceExpectationCodeMap maps from string to cpb.ConformanceExpectationCode_Value.
var DefaultConformanceExpectationCodeMap = map[string]cpb.ConformanceExpectationCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConformanceExpectationCode_INVALID_UNINITIALIZED, // Enum 0
	"MAY":                   cpb.ConformanceExpectationCode_MAY,                   // Enum 3
	"SHALL":                 cpb.ConformanceExpectationCode_SHALL,                 // Enum 1
	"SHOULD":                cpb.ConformanceExpectationCode_SHOULD,                // Enum 2
	"SHOULD_NOT":            cpb.ConformanceExpectationCode_SHOULD_NOT,            // Enum 4
}

// DefaultConsentDataMeaningCodeMap maps from string to cpb.ConsentDataMeaningCode_Value.
var DefaultConsentDataMeaningCodeMap = map[string]cpb.ConsentDataMeaningCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConsentDataMeaningCode_INVALID_UNINITIALIZED, // Enum 0
	"AUTHOREDBY":            cpb.ConsentDataMeaningCode_AUTHOREDBY,            // Enum 4
	"DEPENDENTS":            cpb.ConsentDataMeaningCode_DEPENDENTS,            // Enum 3
	"INSTANCE":              cpb.ConsentDataMeaningCode_INSTANCE,              // Enum 1
	"RELATED":               cpb.ConsentDataMeaningCode_RELATED,               // Enum 2
}

// DefaultConsentProvisionTypeCodeMap maps from string to cpb.ConsentProvisionTypeCode_Value.
var DefaultConsentProvisionTypeCodeMap = map[string]cpb.ConsentProvisionTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConsentProvisionTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"DENY":                  cpb.ConsentProvisionTypeCode_DENY,                  // Enum 1
	"PERMIT":                cpb.ConsentProvisionTypeCode_PERMIT,                // Enum 2
}

// DefaultConsentStateCodeMap maps from string to cpb.ConsentStateCode_Value.
var DefaultConsentStateCodeMap = map[string]cpb.ConsentStateCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConsentStateCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.ConsentStateCode_ACTIVE,                // Enum 3
	"DRAFT":                 cpb.ConsentStateCode_DRAFT,                 // Enum 1
	"ENTERED_IN_ERROR":      cpb.ConsentStateCode_ENTERED_IN_ERROR,      // Enum 6
	"INACTIVE":              cpb.ConsentStateCode_INACTIVE,              // Enum 5
	"PROPOSED":              cpb.ConsentStateCode_PROPOSED,              // Enum 2
	"REJECTED":              cpb.ConsentStateCode_REJECTED,              // Enum 4
}

// DefaultConstraintSeverityCodeMap maps from string to cpb.ConstraintSeverityCode_Value.
var DefaultConstraintSeverityCodeMap = map[string]cpb.ConstraintSeverityCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConstraintSeverityCode_INVALID_UNINITIALIZED, // Enum 0
	"ERROR":                 cpb.ConstraintSeverityCode_ERROR,                 // Enum 1
	"WARNING":               cpb.ConstraintSeverityCode_WARNING,               // Enum 2
}

// DefaultContactPointSystemCodeMap maps from string to cpb.ContactPointSystemCode_Value.
var DefaultContactPointSystemCodeMap = map[string]cpb.ContactPointSystemCode_Value{
	"INVALID_UNINITIALIZED": cpb.ContactPointSystemCode_INVALID_UNINITIALIZED, // Enum 0
	"EMAIL":                 cpb.ContactPointSystemCode_EMAIL,                 // Enum 3
	"FAX":                   cpb.ContactPointSystemCode_FAX,                   // Enum 2
	"OTHER":                 cpb.ContactPointSystemCode_OTHER,                 // Enum 7
	"PAGER":                 cpb.ContactPointSystemCode_PAGER,                 // Enum 4
	"PHONE":                 cpb.ContactPointSystemCode_PHONE,                 // Enum 1
	"SMS":                   cpb.ContactPointSystemCode_SMS,                   // Enum 6
	"URL":                   cpb.ContactPointSystemCode_URL,                   // Enum 5
}

// DefaultContactPointUseCodeMap maps from string to cpb.ContactPointUseCode_Value.
var DefaultContactPointUseCodeMap = map[string]cpb.ContactPointUseCode_Value{
	"INVALID_UNINITIALIZED": cpb.ContactPointUseCode_INVALID_UNINITIALIZED, // Enum 0
	"HOME":                  cpb.ContactPointUseCode_HOME,                  // Enum 1
	"MOBILE":                cpb.ContactPointUseCode_MOBILE,                // Enum 5
	"OLD":                   cpb.ContactPointUseCode_OLD,                   // Enum 4
	"TEMP":                  cpb.ContactPointUseCode_TEMP,                  // Enum 3
	"WORK":                  cpb.ContactPointUseCode_WORK,                  // Enum 2
}

// DefaultContractResourcePublicationStatusCodeMap maps from string to cpb.ContractResourcePublicationStatusCode_Value.
var DefaultContractResourcePublicationStatusCodeMap = map[string]cpb.ContractResourcePublicationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ContractResourcePublicationStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"AMENDED":               cpb.ContractResourcePublicationStatusCode_AMENDED,               // Enum 1
	"APPENDED":              cpb.ContractResourcePublicationStatusCode_APPENDED,              // Enum 2
	"CANCELLED":             cpb.ContractResourcePublicationStatusCode_CANCELLED,             // Enum 3
	"DISPUTED":              cpb.ContractResourcePublicationStatusCode_DISPUTED,              // Enum 4
	"ENTERED_IN_ERROR":      cpb.ContractResourcePublicationStatusCode_ENTERED_IN_ERROR,      // Enum 5
	"EXECUTABLE":            cpb.ContractResourcePublicationStatusCode_EXECUTABLE,            // Enum 6
	"EXECUTED":              cpb.ContractResourcePublicationStatusCode_EXECUTED,              // Enum 7
	"NEGOTIABLE":            cpb.ContractResourcePublicationStatusCode_NEGOTIABLE,            // Enum 8
	"OFFERED":               cpb.ContractResourcePublicationStatusCode_OFFERED,               // Enum 9
	"POLICY":                cpb.ContractResourcePublicationStatusCode_POLICY,                // Enum 10
	"REJECTED":              cpb.ContractResourcePublicationStatusCode_REJECTED,              // Enum 11
	"RENEWED":               cpb.ContractResourcePublicationStatusCode_RENEWED,               // Enum 12
	"RESOLVED":              cpb.ContractResourcePublicationStatusCode_RESOLVED,              // Enum 14
	"REVOKED":               cpb.ContractResourcePublicationStatusCode_REVOKED,               // Enum 13
	"TERMINATED":            cpb.ContractResourcePublicationStatusCode_TERMINATED,            // Enum 15
}

// DefaultContractResourceStatusCodeMap maps from string to cpb.ContractResourceStatusCode_Value.
var DefaultContractResourceStatusCodeMap = map[string]cpb.ContractResourceStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ContractResourceStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"AMENDED":               cpb.ContractResourceStatusCode_AMENDED,               // Enum 1
	"APPENDED":              cpb.ContractResourceStatusCode_APPENDED,              // Enum 2
	"CANCELLED":             cpb.ContractResourceStatusCode_CANCELLED,             // Enum 3
	"DISPUTED":              cpb.ContractResourceStatusCode_DISPUTED,              // Enum 4
	"ENTERED_IN_ERROR":      cpb.ContractResourceStatusCode_ENTERED_IN_ERROR,      // Enum 5
	"EXECUTABLE":            cpb.ContractResourceStatusCode_EXECUTABLE,            // Enum 6
	"EXECUTED":              cpb.ContractResourceStatusCode_EXECUTED,              // Enum 7
	"NEGOTIABLE":            cpb.ContractResourceStatusCode_NEGOTIABLE,            // Enum 8
	"OFFERED":               cpb.ContractResourceStatusCode_OFFERED,               // Enum 9
	"POLICY":                cpb.ContractResourceStatusCode_POLICY,                // Enum 10
	"REJECTED":              cpb.ContractResourceStatusCode_REJECTED,              // Enum 11
	"RENEWED":               cpb.ContractResourceStatusCode_RENEWED,               // Enum 12
	"RESOLVED":              cpb.ContractResourceStatusCode_RESOLVED,              // Enum 14
	"REVOKED":               cpb.ContractResourceStatusCode_REVOKED,               // Enum 13
	"TERMINATED":            cpb.ContractResourceStatusCode_TERMINATED,            // Enum 15
}

// DefaultContributorTypeCodeMap maps from string to cpb.ContributorTypeCode_Value.
var DefaultContributorTypeCodeMap = map[string]cpb.ContributorTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ContributorTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"AUTHOR":                cpb.ContributorTypeCode_AUTHOR,                // Enum 1
	"EDITOR":                cpb.ContributorTypeCode_EDITOR,                // Enum 2
	"ENDORSER":              cpb.ContributorTypeCode_ENDORSER,              // Enum 4
	"REVIEWER":              cpb.ContributorTypeCode_REVIEWER,              // Enum 3
}

// DefaultDataAbsentReasonCodeMap maps from string to cpb.DataAbsentReasonCode_Value.
var DefaultDataAbsentReasonCodeMap = map[string]cpb.DataAbsentReasonCode_Value{
	"INVALID_UNINITIALIZED": cpb.DataAbsentReasonCode_INVALID_UNINITIALIZED, // Enum 0
	"ASKED_DECLINED":        cpb.DataAbsentReasonCode_ASKED_DECLINED,        // Enum 5
	"ASKED_UNKNOWN":         cpb.DataAbsentReasonCode_ASKED_UNKNOWN,         // Enum 2
	"AS_TEXT":               cpb.DataAbsentReasonCode_AS_TEXT,               // Enum 9
	"ERROR":                 cpb.DataAbsentReasonCode_ERROR,                 // Enum 10
	"MASKED":                cpb.DataAbsentReasonCode_MASKED,                // Enum 6
	"NEGATIVE_INFINITY":     cpb.DataAbsentReasonCode_NEGATIVE_INFINITY,     // Enum 12
	"NOT_APPLICABLE":        cpb.DataAbsentReasonCode_NOT_APPLICABLE,        // Enum 7
	"NOT_ASKED":             cpb.DataAbsentReasonCode_NOT_ASKED,             // Enum 4
	"NOT_A_NUMBER":          cpb.DataAbsentReasonCode_NOT_A_NUMBER,          // Enum 11
	"NOT_PERFORMED":         cpb.DataAbsentReasonCode_NOT_PERFORMED,         // Enum 14
	"NOT_PERMITTED":         cpb.DataAbsentReasonCode_NOT_PERMITTED,         // Enum 15
	"POSITIVE_INFINITY":     cpb.DataAbsentReasonCode_POSITIVE_INFINITY,     // Enum 13
	"TEMP_UNKNOWN":          cpb.DataAbsentReasonCode_TEMP_UNKNOWN,          // Enum 3
	"UNKNOWN":               cpb.DataAbsentReasonCode_UNKNOWN,               // Enum 1
	"UNSUPPORTED":           cpb.DataAbsentReasonCode_UNSUPPORTED,           // Enum 8
}

// DefaultDataTypeCodeMap maps from string to cpb.DataTypeCode_Value.
var DefaultDataTypeCodeMap = map[string]cpb.DataTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.DataTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"ADDRESS":               cpb.DataTypeCode_ADDRESS,               // Enum 1
	"AGE":                   cpb.DataTypeCode_AGE,                   // Enum 2
	"ANNOTATION":            cpb.DataTypeCode_ANNOTATION,            // Enum 3
	"ATTACHMENT":            cpb.DataTypeCode_ATTACHMENT,            // Enum 4
	"BACKBONE_ELEMENT":      cpb.DataTypeCode_BACKBONE_ELEMENT,      // Enum 5
	"BASE64_BINARY":         cpb.DataTypeCode_BASE64_BINARY,         // Enum 44
	"BOOLEAN":               cpb.DataTypeCode_BOOLEAN,               // Enum 45
	"CANONICAL":             cpb.DataTypeCode_CANONICAL,             // Enum 46
	"CODE":                  cpb.DataTypeCode_CODE,                  // Enum 47
	"CODEABLE_CONCEPT":      cpb.DataTypeCode_CODEABLE_CONCEPT,      // Enum 6
	"CODING":                cpb.DataTypeCode_CODING,                // Enum 7
	"CONTACT_DETAIL":        cpb.DataTypeCode_CONTACT_DETAIL,        // Enum 8
	"CONTACT_POINT":         cpb.DataTypeCode_CONTACT_POINT,         // Enum 9
	"CONTRIBUTOR":           cpb.DataTypeCode_CONTRIBUTOR,           // Enum 10
	"COUNT":                 cpb.DataTypeCode_COUNT,                 // Enum 11
	"DATA_REQUIREMENT":      cpb.DataTypeCode_DATA_REQUIREMENT,      // Enum 12
	"DATE":                  cpb.DataTypeCode_DATE,                  // Enum 48
	"DATE_TIME":             cpb.DataTypeCode_DATE_TIME,             // Enum 49
	"DECIMAL":               cpb.DataTypeCode_DECIMAL,               // Enum 50
	"DISTANCE":              cpb.DataTypeCode_DISTANCE,              // Enum 13
	"DOSAGE":                cpb.DataTypeCode_DOSAGE,                // Enum 14
	"DURATION":              cpb.DataTypeCode_DURATION,              // Enum 15
	"ELEMENT":               cpb.DataTypeCode_ELEMENT,               // Enum 16
	"ELEMENT_DEFINITION":    cpb.DataTypeCode_ELEMENT_DEFINITION,    // Enum 17
	"EXPRESSION":            cpb.DataTypeCode_EXPRESSION,            // Enum 18
	"EXTENSION":             cpb.DataTypeCode_EXTENSION,             // Enum 19
	"HUMAN_NAME":            cpb.DataTypeCode_HUMAN_NAME,            // Enum 20
	"ID":                    cpb.DataTypeCode_ID,                    // Enum 51
	"IDENTIFIER":            cpb.DataTypeCode_IDENTIFIER,            // Enum 21
	"INSTANT":               cpb.DataTypeCode_INSTANT,               // Enum 52
	"INTEGER":               cpb.DataTypeCode_INTEGER,               // Enum 53
	"MARKDOWN":              cpb.DataTypeCode_MARKDOWN,              // Enum 54
	"MARKETING_STATUS":      cpb.DataTypeCode_MARKETING_STATUS,      // Enum 22
	"META":                  cpb.DataTypeCode_META,                  // Enum 23
	"MONEY":                 cpb.DataTypeCode_MONEY,                 // Enum 24
	"MONEY_QUANTITY":        cpb.DataTypeCode_MONEY_QUANTITY,        // Enum 25
	"NARRATIVE":             cpb.DataTypeCode_NARRATIVE,             // Enum 26
	"OID":                   cpb.DataTypeCode_OID,                   // Enum 55
	"PARAMETER_DEFINITION":  cpb.DataTypeCode_PARAMETER_DEFINITION,  // Enum 27
	"PERIOD":                cpb.DataTypeCode_PERIOD,                // Enum 28
	"POPULATION":            cpb.DataTypeCode_POPULATION,            // Enum 29
	"POSITIVE_INT":          cpb.DataTypeCode_POSITIVE_INT,          // Enum 56
	"PRODUCT_SHELF_LIFE":    cpb.DataTypeCode_PRODUCT_SHELF_LIFE,    // Enum 31
	"PROD_CHARACTERISTIC":   cpb.DataTypeCode_PROD_CHARACTERISTIC,   // Enum 30
	"QUANTITY":              cpb.DataTypeCode_QUANTITY,              // Enum 32
	"RANGE":                 cpb.DataTypeCode_RANGE,                 // Enum 33
	"RATIO":                 cpb.DataTypeCode_RATIO,                 // Enum 34
	"REFERENCE":             cpb.DataTypeCode_REFERENCE,             // Enum 35
	"RELATED_ARTIFACT":      cpb.DataTypeCode_RELATED_ARTIFACT,      // Enum 36
	"SAMPLED_DATA":          cpb.DataTypeCode_SAMPLED_DATA,          // Enum 37
	"SIGNATURE":             cpb.DataTypeCode_SIGNATURE,             // Enum 38
	"SIMPLE_QUANTITY":       cpb.DataTypeCode_SIMPLE_QUANTITY,       // Enum 39
	"STRING":                cpb.DataTypeCode_STRING,                // Enum 57
	"SUBSTANCE_AMOUNT":      cpb.DataTypeCode_SUBSTANCE_AMOUNT,      // Enum 40
	"TIME":                  cpb.DataTypeCode_TIME,                  // Enum 58
	"TIMING":                cpb.DataTypeCode_TIMING,                // Enum 41
	"TRIGGER_DEFINITION":    cpb.DataTypeCode_TRIGGER_DEFINITION,    // Enum 42
	"UNSIGNED_INT":          cpb.DataTypeCode_UNSIGNED_INT,          // Enum 59
	"URI":                   cpb.DataTypeCode_URI,                   // Enum 60
	"URL":                   cpb.DataTypeCode_URL,                   // Enum 61
	"USAGE_CONTEXT":         cpb.DataTypeCode_USAGE_CONTEXT,         // Enum 43
	"UUID":                  cpb.DataTypeCode_UUID,                  // Enum 62
	"XHTML":                 cpb.DataTypeCode_XHTML,                 // Enum 63
}

// DefaultDaysOfWeekCodeMap maps from string to cpb.DaysOfWeekCode_Value.
var DefaultDaysOfWeekCodeMap = map[string]cpb.DaysOfWeekCode_Value{
	"INVALID_UNINITIALIZED": cpb.DaysOfWeekCode_INVALID_UNINITIALIZED, // Enum 0
	"FRI":                   cpb.DaysOfWeekCode_FRI,                   // Enum 5
	"MON":                   cpb.DaysOfWeekCode_MON,                   // Enum 1
	"SAT":                   cpb.DaysOfWeekCode_SAT,                   // Enum 6
	"SUN":                   cpb.DaysOfWeekCode_SUN,                   // Enum 7
	"THU":                   cpb.DaysOfWeekCode_THU,                   // Enum 4
	"TUE":                   cpb.DaysOfWeekCode_TUE,                   // Enum 2
	"WED":                   cpb.DaysOfWeekCode_WED,                   // Enum 3
}

// DefaultDetectedIssueSeverityCodeMap maps from string to cpb.DetectedIssueSeverityCode_Value.
var DefaultDetectedIssueSeverityCodeMap = map[string]cpb.DetectedIssueSeverityCode_Value{
	"INVALID_UNINITIALIZED": cpb.DetectedIssueSeverityCode_INVALID_UNINITIALIZED, // Enum 0
	"HIGH":                  cpb.DetectedIssueSeverityCode_HIGH,                  // Enum 1
	"LOW":                   cpb.DetectedIssueSeverityCode_LOW,                   // Enum 3
	"MODERATE":              cpb.DetectedIssueSeverityCode_MODERATE,              // Enum 2
}

// DefaultDeviceMetricCalibrationStateCodeMap maps from string to cpb.DeviceMetricCalibrationStateCode_Value.
var DefaultDeviceMetricCalibrationStateCodeMap = map[string]cpb.DeviceMetricCalibrationStateCode_Value{
	"INVALID_UNINITIALIZED": cpb.DeviceMetricCalibrationStateCode_INVALID_UNINITIALIZED, // Enum 0
	"CALIBRATED":            cpb.DeviceMetricCalibrationStateCode_CALIBRATED,            // Enum 3
	"CALIBRATION_REQUIRED":  cpb.DeviceMetricCalibrationStateCode_CALIBRATION_REQUIRED,  // Enum 2
	"NOT_CALIBRATED":        cpb.DeviceMetricCalibrationStateCode_NOT_CALIBRATED,        // Enum 1
	"UNSPECIFIED":           cpb.DeviceMetricCalibrationStateCode_UNSPECIFIED,           // Enum 4
}

// DefaultDeviceMetricCalibrationTypeCodeMap maps from string to cpb.DeviceMetricCalibrationTypeCode_Value.
var DefaultDeviceMetricCalibrationTypeCodeMap = map[string]cpb.DeviceMetricCalibrationTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.DeviceMetricCalibrationTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"GAIN":                  cpb.DeviceMetricCalibrationTypeCode_GAIN,                  // Enum 3
	"OFFSET":                cpb.DeviceMetricCalibrationTypeCode_OFFSET,                // Enum 2
	"TWO_POINT":             cpb.DeviceMetricCalibrationTypeCode_TWO_POINT,             // Enum 4
	"UNSPECIFIED":           cpb.DeviceMetricCalibrationTypeCode_UNSPECIFIED,           // Enum 1
}

// DefaultDeviceMetricCategoryCodeMap maps from string to cpb.DeviceMetricCategoryCode_Value.
var DefaultDeviceMetricCategoryCodeMap = map[string]cpb.DeviceMetricCategoryCode_Value{
	"INVALID_UNINITIALIZED": cpb.DeviceMetricCategoryCode_INVALID_UNINITIALIZED, // Enum 0
	"CALCULATION":           cpb.DeviceMetricCategoryCode_CALCULATION,           // Enum 3
	"MEASUREMENT":           cpb.DeviceMetricCategoryCode_MEASUREMENT,           // Enum 1
	"SETTING":               cpb.DeviceMetricCategoryCode_SETTING,               // Enum 2
	"UNSPECIFIED":           cpb.DeviceMetricCategoryCode_UNSPECIFIED,           // Enum 4
}

// DefaultDeviceMetricColorCodeMap maps from string to cpb.DeviceMetricColorCode_Value.
var DefaultDeviceMetricColorCodeMap = map[string]cpb.DeviceMetricColorCode_Value{
	"INVALID_UNINITIALIZED": cpb.DeviceMetricColorCode_INVALID_UNINITIALIZED, // Enum 0
	"BLACK":                 cpb.DeviceMetricColorCode_BLACK,                 // Enum 1
	"BLUE":                  cpb.DeviceMetricColorCode_BLUE,                  // Enum 5
	"CYAN":                  cpb.DeviceMetricColorCode_CYAN,                  // Enum 7
	"GREEN":                 cpb.DeviceMetricColorCode_GREEN,                 // Enum 3
	"MAGENTA":               cpb.DeviceMetricColorCode_MAGENTA,               // Enum 6
	"RED":                   cpb.DeviceMetricColorCode_RED,                   // Enum 2
	"WHITE":                 cpb.DeviceMetricColorCode_WHITE,                 // Enum 8
	"YELLOW":                cpb.DeviceMetricColorCode_YELLOW,                // Enum 4
}

// DefaultDeviceMetricOperationalStatusCodeMap maps from string to cpb.DeviceMetricOperationalStatusCode_Value.
var DefaultDeviceMetricOperationalStatusCodeMap = map[string]cpb.DeviceMetricOperationalStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.DeviceMetricOperationalStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ENTERED_IN_ERROR":      cpb.DeviceMetricOperationalStatusCode_ENTERED_IN_ERROR,      // Enum 4
	"OFF":                   cpb.DeviceMetricOperationalStatusCode_OFF,                   // Enum 2
	"ON":                    cpb.DeviceMetricOperationalStatusCode_ON,                    // Enum 1
	"STANDBY":               cpb.DeviceMetricOperationalStatusCode_STANDBY,               // Enum 3
}

// DefaultDeviceNameTypeCodeMap maps from string to cpb.DeviceNameTypeCode_Value.
var DefaultDeviceNameTypeCodeMap = map[string]cpb.DeviceNameTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.DeviceNameTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"MANUFACTURER_NAME":     cpb.DeviceNameTypeCode_MANUFACTURER_NAME,     // Enum 4
	"MODEL_NAME":            cpb.DeviceNameTypeCode_MODEL_NAME,            // Enum 5
	"OTHER":                 cpb.DeviceNameTypeCode_OTHER,                 // Enum 6
	"PATIENT_REPORTED_NAME": cpb.DeviceNameTypeCode_PATIENT_REPORTED_NAME, // Enum 3
	"UDI_LABEL_NAME":        cpb.DeviceNameTypeCode_UDI_LABEL_NAME,        // Enum 1
	"USER_FRIENDLY_NAME":    cpb.DeviceNameTypeCode_USER_FRIENDLY_NAME,    // Enum 2
}

// DefaultDeviceUseStatementStatusCodeMap maps from string to cpb.DeviceUseStatementStatusCode_Value.
var DefaultDeviceUseStatementStatusCodeMap = map[string]cpb.DeviceUseStatementStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.DeviceUseStatementStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.DeviceUseStatementStatusCode_ACTIVE,                // Enum 1
	"COMPLETED":             cpb.DeviceUseStatementStatusCode_COMPLETED,             // Enum 2
	"ENTERED_IN_ERROR":      cpb.DeviceUseStatementStatusCode_ENTERED_IN_ERROR,      // Enum 3
	"INTENDED":              cpb.DeviceUseStatementStatusCode_INTENDED,              // Enum 4
	"ON_HOLD":               cpb.DeviceUseStatementStatusCode_ON_HOLD,               // Enum 6
	"STOPPED":               cpb.DeviceUseStatementStatusCode_STOPPED,               // Enum 5
}

// DefaultDiagnosticReportStatusCodeMap maps from string to cpb.DiagnosticReportStatusCode_Value.
var DefaultDiagnosticReportStatusCodeMap = map[string]cpb.DiagnosticReportStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.DiagnosticReportStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"AMENDED":               cpb.DiagnosticReportStatusCode_AMENDED,               // Enum 5
	"APPENDED":              cpb.DiagnosticReportStatusCode_APPENDED,              // Enum 7
	"CANCELLED":             cpb.DiagnosticReportStatusCode_CANCELLED,             // Enum 8
	"CORRECTED":             cpb.DiagnosticReportStatusCode_CORRECTED,             // Enum 6
	"ENTERED_IN_ERROR":      cpb.DiagnosticReportStatusCode_ENTERED_IN_ERROR,      // Enum 9
	"FINAL":                 cpb.DiagnosticReportStatusCode_FINAL,                 // Enum 4
	"PARTIAL":               cpb.DiagnosticReportStatusCode_PARTIAL,               // Enum 2
	"PRELIMINARY":           cpb.DiagnosticReportStatusCode_PRELIMINARY,           // Enum 3
	"REGISTERED":            cpb.DiagnosticReportStatusCode_REGISTERED,            // Enum 1
	"UNKNOWN":               cpb.DiagnosticReportStatusCode_UNKNOWN,               // Enum 10
}

// DefaultDiscriminatorTypeCodeMap maps from string to cpb.DiscriminatorTypeCode_Value.
var DefaultDiscriminatorTypeCodeMap = map[string]cpb.DiscriminatorTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.DiscriminatorTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"EXISTS":                cpb.DiscriminatorTypeCode_EXISTS,                // Enum 2
	"PATTERN":               cpb.DiscriminatorTypeCode_PATTERN,               // Enum 3
	"PROFILE":               cpb.DiscriminatorTypeCode_PROFILE,               // Enum 5
	"TYPE":                  cpb.DiscriminatorTypeCode_TYPE,                  // Enum 4
	"VALUE":                 cpb.DiscriminatorTypeCode_VALUE,                 // Enum 1
}

// DefaultDocumentModeCodeMap maps from string to cpb.DocumentModeCode_Value.
var DefaultDocumentModeCodeMap = map[string]cpb.DocumentModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.DocumentModeCode_INVALID_UNINITIALIZED, // Enum 0
	"CONSUMER":              cpb.DocumentModeCode_CONSUMER,              // Enum 2
	"PRODUCER":              cpb.DocumentModeCode_PRODUCER,              // Enum 1
}

// DefaultDocumentReferenceStatusCodeMap maps from string to cpb.DocumentReferenceStatusCode_Value.
var DefaultDocumentReferenceStatusCodeMap = map[string]cpb.DocumentReferenceStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.DocumentReferenceStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"CURRENT":               cpb.DocumentReferenceStatusCode_CURRENT,               // Enum 1
	"ENTERED_IN_ERROR":      cpb.DocumentReferenceStatusCode_ENTERED_IN_ERROR,      // Enum 3
	"SUPERSEDED":            cpb.DocumentReferenceStatusCode_SUPERSEDED,            // Enum 2
}

// DefaultDocumentRelationshipTypeCodeMap maps from string to cpb.DocumentRelationshipTypeCode_Value.
var DefaultDocumentRelationshipTypeCodeMap = map[string]cpb.DocumentRelationshipTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.DocumentRelationshipTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"APPENDS":               cpb.DocumentRelationshipTypeCode_APPENDS,               // Enum 4
	"REPLACES":              cpb.DocumentRelationshipTypeCode_REPLACES,              // Enum 1
	"SIGNS":                 cpb.DocumentRelationshipTypeCode_SIGNS,                 // Enum 3
	"TRANSFORMS":            cpb.DocumentRelationshipTypeCode_TRANSFORMS,            // Enum 2
}

// DefaultEligibilityRequestPurposeCodeMap maps from string to cpb.EligibilityRequestPurposeCode_Value.
var DefaultEligibilityRequestPurposeCodeMap = map[string]cpb.EligibilityRequestPurposeCode_Value{
	"INVALID_UNINITIALIZED": cpb.EligibilityRequestPurposeCode_INVALID_UNINITIALIZED, // Enum 0
	"AUTH_REQUIREMENTS":     cpb.EligibilityRequestPurposeCode_AUTH_REQUIREMENTS,     // Enum 1
	"BENEFITS":              cpb.EligibilityRequestPurposeCode_BENEFITS,              // Enum 2
	"DISCOVERY":             cpb.EligibilityRequestPurposeCode_DISCOVERY,             // Enum 3
	"VALIDATION":            cpb.EligibilityRequestPurposeCode_VALIDATION,            // Enum 4
}

// DefaultEligibilityResponsePurposeCodeMap maps from string to cpb.EligibilityResponsePurposeCode_Value.
var DefaultEligibilityResponsePurposeCodeMap = map[string]cpb.EligibilityResponsePurposeCode_Value{
	"INVALID_UNINITIALIZED": cpb.EligibilityResponsePurposeCode_INVALID_UNINITIALIZED, // Enum 0
	"AUTH_REQUIREMENTS":     cpb.EligibilityResponsePurposeCode_AUTH_REQUIREMENTS,     // Enum 1
	"BENEFITS":              cpb.EligibilityResponsePurposeCode_BENEFITS,              // Enum 2
	"DISCOVERY":             cpb.EligibilityResponsePurposeCode_DISCOVERY,             // Enum 3
	"VALIDATION":            cpb.EligibilityResponsePurposeCode_VALIDATION,            // Enum 4
}

// DefaultEnableWhenBehaviorCodeMap maps from string to cpb.EnableWhenBehaviorCode_Value.
var DefaultEnableWhenBehaviorCodeMap = map[string]cpb.EnableWhenBehaviorCode_Value{
	"INVALID_UNINITIALIZED": cpb.EnableWhenBehaviorCode_INVALID_UNINITIALIZED, // Enum 0
	"ALL":                   cpb.EnableWhenBehaviorCode_ALL,                   // Enum 1
	"ANY":                   cpb.EnableWhenBehaviorCode_ANY,                   // Enum 2
}

// DefaultEncounterLocationStatusCodeMap maps from string to cpb.EncounterLocationStatusCode_Value.
var DefaultEncounterLocationStatusCodeMap = map[string]cpb.EncounterLocationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.EncounterLocationStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.EncounterLocationStatusCode_ACTIVE,                // Enum 2
	"COMPLETED":             cpb.EncounterLocationStatusCode_COMPLETED,             // Enum 4
	"PLANNED":               cpb.EncounterLocationStatusCode_PLANNED,               // Enum 1
	"RESERVED":              cpb.EncounterLocationStatusCode_RESERVED,              // Enum 3
}

// DefaultEncounterStatusCodeMap maps from string to cpb.EncounterStatusCode_Value.
var DefaultEncounterStatusCodeMap = map[string]cpb.EncounterStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.EncounterStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ARRIVED":               cpb.EncounterStatusCode_ARRIVED,               // Enum 2
	"CANCELLED":             cpb.EncounterStatusCode_CANCELLED,             // Enum 7
	"ENTERED_IN_ERROR":      cpb.EncounterStatusCode_ENTERED_IN_ERROR,      // Enum 8
	"FINISHED":              cpb.EncounterStatusCode_FINISHED,              // Enum 6
	"IN_PROGRESS":           cpb.EncounterStatusCode_IN_PROGRESS,           // Enum 4
	"ONLEAVE":               cpb.EncounterStatusCode_ONLEAVE,               // Enum 5
	"PLANNED":               cpb.EncounterStatusCode_PLANNED,               // Enum 1
	"TRIAGED":               cpb.EncounterStatusCode_TRIAGED,               // Enum 3
	"UNKNOWN":               cpb.EncounterStatusCode_UNKNOWN,               // Enum 9
}

// DefaultEndpointStatusCodeMap maps from string to cpb.EndpointStatusCode_Value.
var DefaultEndpointStatusCodeMap = map[string]cpb.EndpointStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.EndpointStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.EndpointStatusCode_ACTIVE,                // Enum 1
	"ENTERED_IN_ERROR":      cpb.EndpointStatusCode_ENTERED_IN_ERROR,      // Enum 5
	"ERROR":                 cpb.EndpointStatusCode_ERROR,                 // Enum 3
	"OFF":                   cpb.EndpointStatusCode_OFF,                   // Enum 4
	"SUSPENDED":             cpb.EndpointStatusCode_SUSPENDED,             // Enum 2
	"TEST":                  cpb.EndpointStatusCode_TEST,                  // Enum 6
}

// DefaultEpisodeOfCareStatusCodeMap maps from string to cpb.EpisodeOfCareStatusCode_Value.
var DefaultEpisodeOfCareStatusCodeMap = map[string]cpb.EpisodeOfCareStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.EpisodeOfCareStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.EpisodeOfCareStatusCode_ACTIVE,                // Enum 3
	"CANCELLED":             cpb.EpisodeOfCareStatusCode_CANCELLED,             // Enum 6
	"ENTERED_IN_ERROR":      cpb.EpisodeOfCareStatusCode_ENTERED_IN_ERROR,      // Enum 7
	"FINISHED":              cpb.EpisodeOfCareStatusCode_FINISHED,              // Enum 5
	"ONHOLD":                cpb.EpisodeOfCareStatusCode_ONHOLD,                // Enum 4
	"PLANNED":               cpb.EpisodeOfCareStatusCode_PLANNED,               // Enum 1
	"WAITLIST":              cpb.EpisodeOfCareStatusCode_WAITLIST,              // Enum 2
}

// DefaultEventCapabilityModeCodeMap maps from string to cpb.EventCapabilityModeCode_Value.
var DefaultEventCapabilityModeCodeMap = map[string]cpb.EventCapabilityModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.EventCapabilityModeCode_INVALID_UNINITIALIZED, // Enum 0
	"RECEIVER":              cpb.EventCapabilityModeCode_RECEIVER,              // Enum 2
	"SENDER":                cpb.EventCapabilityModeCode_SENDER,                // Enum 1
}

// DefaultEventStatusCodeMap maps from string to cpb.EventStatusCode_Value.
var DefaultEventStatusCodeMap = map[string]cpb.EventStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.EventStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"COMPLETED":             cpb.EventStatusCode_COMPLETED,             // Enum 6
	"ENTERED_IN_ERROR":      cpb.EventStatusCode_ENTERED_IN_ERROR,      // Enum 7
	"IN_PROGRESS":           cpb.EventStatusCode_IN_PROGRESS,           // Enum 2
	"NOT_DONE":              cpb.EventStatusCode_NOT_DONE,              // Enum 3
	"ON_HOLD":               cpb.EventStatusCode_ON_HOLD,               // Enum 4
	"PREPARATION":           cpb.EventStatusCode_PREPARATION,           // Enum 1
	"STOPPED":               cpb.EventStatusCode_STOPPED,               // Enum 5
	"UNKNOWN":               cpb.EventStatusCode_UNKNOWN,               // Enum 8
}

// DefaultEventTimingCodeMap maps from string to cpb.EventTimingCode_Value.
var DefaultEventTimingCodeMap = map[string]cpb.EventTimingCode_Value{
	"INVALID_UNINITIALIZED": cpb.EventTimingCode_INVALID_UNINITIALIZED, // Enum 0
	"AFT":                   cpb.EventTimingCode_AFT,                   // Enum 5
	"AFT_EARLY":             cpb.EventTimingCode_AFT_EARLY,             // Enum 6
	"AFT_LATE":              cpb.EventTimingCode_AFT_LATE,              // Enum 7
	"EVE":                   cpb.EventTimingCode_EVE,                   // Enum 8
	"EVE_EARLY":             cpb.EventTimingCode_EVE_EARLY,             // Enum 9
	"EVE_LATE":              cpb.EventTimingCode_EVE_LATE,              // Enum 10
	"MORN":                  cpb.EventTimingCode_MORN,                  // Enum 1
	"MORN_EARLY":            cpb.EventTimingCode_MORN_EARLY,            // Enum 2
	"MORN_LATE":             cpb.EventTimingCode_MORN_LATE,             // Enum 3
	"NIGHT":                 cpb.EventTimingCode_NIGHT,                 // Enum 11
	"NOON":                  cpb.EventTimingCode_NOON,                  // Enum 4
	"PHS":                   cpb.EventTimingCode_PHS,                   // Enum 12
}

// DefaultEvidenceVariableTypeCodeMap maps from string to cpb.EvidenceVariableTypeCode_Value.
var DefaultEvidenceVariableTypeCodeMap = map[string]cpb.EvidenceVariableTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.EvidenceVariableTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"CONTINUOUS":            cpb.EvidenceVariableTypeCode_CONTINUOUS,            // Enum 2
	"DESCRIPTIVE":           cpb.EvidenceVariableTypeCode_DESCRIPTIVE,           // Enum 3
	"DICHOTOMOUS":           cpb.EvidenceVariableTypeCode_DICHOTOMOUS,           // Enum 1
}

// DefaultExampleScenarioActorTypeCodeMap maps from string to cpb.ExampleScenarioActorTypeCode_Value.
var DefaultExampleScenarioActorTypeCodeMap = map[string]cpb.ExampleScenarioActorTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ExampleScenarioActorTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"ENTITY":                cpb.ExampleScenarioActorTypeCode_ENTITY,                // Enum 2
	"PERSON":                cpb.ExampleScenarioActorTypeCode_PERSON,                // Enum 1
}

// DefaultExpansionParameterSourceCodeMap maps from string to cpb.ExpansionParameterSourceCode_Value.
var DefaultExpansionParameterSourceCodeMap = map[string]cpb.ExpansionParameterSourceCode_Value{
	"INVALID_UNINITIALIZED": cpb.ExpansionParameterSourceCode_INVALID_UNINITIALIZED, // Enum 0
	"CODESYSTEM":            cpb.ExpansionParameterSourceCode_CODESYSTEM,            // Enum 3
	"INPUT":                 cpb.ExpansionParameterSourceCode_INPUT,                 // Enum 1
	"SERVER":                cpb.ExpansionParameterSourceCode_SERVER,                // Enum 2
}

// DefaultExpansionProcessingRuleCodeMap maps from string to cpb.ExpansionProcessingRuleCode_Value.
var DefaultExpansionProcessingRuleCodeMap = map[string]cpb.ExpansionProcessingRuleCode_Value{
	"INVALID_UNINITIALIZED": cpb.ExpansionProcessingRuleCode_INVALID_UNINITIALIZED, // Enum 0
	"ALL_CODES":             cpb.ExpansionProcessingRuleCode_ALL_CODES,             // Enum 1
	"GROUPS_ONLY":           cpb.ExpansionProcessingRuleCode_GROUPS_ONLY,           // Enum 3
	"UNGROUPED":             cpb.ExpansionProcessingRuleCode_UNGROUPED,             // Enum 2
}

// DefaultExplanationOfBenefitStatusCodeMap maps from string to cpb.ExplanationOfBenefitStatusCode_Value.
var DefaultExplanationOfBenefitStatusCodeMap = map[string]cpb.ExplanationOfBenefitStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ExplanationOfBenefitStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.ExplanationOfBenefitStatusCode_ACTIVE,                // Enum 1
	"CANCELLED":             cpb.ExplanationOfBenefitStatusCode_CANCELLED,             // Enum 2
	"DRAFT":                 cpb.ExplanationOfBenefitStatusCode_DRAFT,                 // Enum 3
	"ENTERED_IN_ERROR":      cpb.ExplanationOfBenefitStatusCode_ENTERED_IN_ERROR,      // Enum 4
}

// DefaultExposureStateCodeMap maps from string to cpb.ExposureStateCode_Value.
var DefaultExposureStateCodeMap = map[string]cpb.ExposureStateCode_Value{
	"INVALID_UNINITIALIZED": cpb.ExposureStateCode_INVALID_UNINITIALIZED, // Enum 0
	"EXPOSURE":              cpb.ExposureStateCode_EXPOSURE,              // Enum 1
	"EXPOSURE_ALTERNATIVE":  cpb.ExposureStateCode_EXPOSURE_ALTERNATIVE,  // Enum 2
}

// DefaultExtensionContextTypeCodeMap maps from string to cpb.ExtensionContextTypeCode_Value.
var DefaultExtensionContextTypeCodeMap = map[string]cpb.ExtensionContextTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ExtensionContextTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"ELEMENT":               cpb.ExtensionContextTypeCode_ELEMENT,               // Enum 2
	"EXTENSION":             cpb.ExtensionContextTypeCode_EXTENSION,             // Enum 3
	"FHIRPATH":              cpb.ExtensionContextTypeCode_FHIRPATH,              // Enum 1
}

// DefaultFHIRDeviceStatusCodeMap maps from string to cpb.FHIRDeviceStatusCode_Value.
var DefaultFHIRDeviceStatusCodeMap = map[string]cpb.FHIRDeviceStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.FHIRDeviceStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.FHIRDeviceStatusCode_ACTIVE,                // Enum 1
	"ENTERED_IN_ERROR":      cpb.FHIRDeviceStatusCode_ENTERED_IN_ERROR,      // Enum 3
	"INACTIVE":              cpb.FHIRDeviceStatusCode_INACTIVE,              // Enum 2
	"UNKNOWN":               cpb.FHIRDeviceStatusCode_UNKNOWN,               // Enum 4
}

// DefaultFHIRRestfulInteractionsCodeMap maps from string to cpb.FHIRRestfulInteractionsCode_Value.
var DefaultFHIRRestfulInteractionsCodeMap = map[string]cpb.FHIRRestfulInteractionsCode_Value{
	"INVALID_UNINITIALIZED": cpb.FHIRRestfulInteractionsCode_INVALID_UNINITIALIZED, // Enum 0
	"BATCH":                 cpb.FHIRRestfulInteractionsCode_BATCH,                 // Enum 16
	"CAPABILITIES":          cpb.FHIRRestfulInteractionsCode_CAPABILITIES,          // Enum 14
	"CREATE":                cpb.FHIRRestfulInteractionsCode_CREATE,                // Enum 10
	"DELETE":                cpb.FHIRRestfulInteractionsCode_DELETE,                // Enum 5
	"HISTORY":               cpb.FHIRRestfulInteractionsCode_HISTORY,               // Enum 6
	"HISTORY_INSTANCE":      cpb.FHIRRestfulInteractionsCode_HISTORY_INSTANCE,      // Enum 7
	"HISTORY_SYSTEM":        cpb.FHIRRestfulInteractionsCode_HISTORY_SYSTEM,        // Enum 9
	"HISTORY_TYPE":          cpb.FHIRRestfulInteractionsCode_HISTORY_TYPE,          // Enum 8
	"OPERATION":             cpb.FHIRRestfulInteractionsCode_OPERATION,             // Enum 17
	"PATCH":                 cpb.FHIRRestfulInteractionsCode_PATCH,                 // Enum 4
	"READ":                  cpb.FHIRRestfulInteractionsCode_READ,                  // Enum 1
	"SEARCH":                cpb.FHIRRestfulInteractionsCode_SEARCH,                // Enum 11
	"SEARCH_SYSTEM":         cpb.FHIRRestfulInteractionsCode_SEARCH_SYSTEM,         // Enum 13
	"SEARCH_TYPE":           cpb.FHIRRestfulInteractionsCode_SEARCH_TYPE,           // Enum 12
	"TRANSACTION":           cpb.FHIRRestfulInteractionsCode_TRANSACTION,           // Enum 15
	"UPDATE":                cpb.FHIRRestfulInteractionsCode_UPDATE,                // Enum 3
	"VREAD":                 cpb.FHIRRestfulInteractionsCode_VREAD,                 // Enum 2
}

// DefaultFHIRSubstanceStatusCodeMap maps from string to cpb.FHIRSubstanceStatusCode_Value.
var DefaultFHIRSubstanceStatusCodeMap = map[string]cpb.FHIRSubstanceStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.FHIRSubstanceStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.FHIRSubstanceStatusCode_ACTIVE,                // Enum 1
	"ENTERED_IN_ERROR":      cpb.FHIRSubstanceStatusCode_ENTERED_IN_ERROR,      // Enum 3
	"INACTIVE":              cpb.FHIRSubstanceStatusCode_INACTIVE,              // Enum 2
}

// DefaultFHIRVersionCodeMap maps from string to cpb.FHIRVersionCode_Value.
var DefaultFHIRVersionCodeMap = map[string]cpb.FHIRVersionCode_Value{
	"INVALID_UNINITIALIZED": cpb.FHIRVersionCode_INVALID_UNINITIALIZED, // Enum 0
	"V_0_01":                cpb.FHIRVersionCode_V_0_01,                // Enum 1
	"V_0_05":                cpb.FHIRVersionCode_V_0_05,                // Enum 2
	"V_0_06":                cpb.FHIRVersionCode_V_0_06,                // Enum 3
	"V_0_0_80":              cpb.FHIRVersionCode_V_0_0_80,              // Enum 5
	"V_0_0_81":              cpb.FHIRVersionCode_V_0_0_81,              // Enum 6
	"V_0_0_82":              cpb.FHIRVersionCode_V_0_0_82,              // Enum 7
	"V_0_11":                cpb.FHIRVersionCode_V_0_11,                // Enum 4
	"V_0_4_0":               cpb.FHIRVersionCode_V_0_4_0,               // Enum 8
	"V_0_5_0":               cpb.FHIRVersionCode_V_0_5_0,               // Enum 9
	"V_1_0_0":               cpb.FHIRVersionCode_V_1_0_0,               // Enum 10
	"V_1_0_1":               cpb.FHIRVersionCode_V_1_0_1,               // Enum 11
	"V_1_0_2":               cpb.FHIRVersionCode_V_1_0_2,               // Enum 12
	"V_1_1_0":               cpb.FHIRVersionCode_V_1_1_0,               // Enum 13
	"V_1_4_0":               cpb.FHIRVersionCode_V_1_4_0,               // Enum 14
	"V_1_6_0":               cpb.FHIRVersionCode_V_1_6_0,               // Enum 15
	"V_1_8_0":               cpb.FHIRVersionCode_V_1_8_0,               // Enum 16
	"V_3_0_0":               cpb.FHIRVersionCode_V_3_0_0,               // Enum 17
	"V_3_0_1":               cpb.FHIRVersionCode_V_3_0_1,               // Enum 18
	"V_3_3_0":               cpb.FHIRVersionCode_V_3_3_0,               // Enum 19
	"V_3_5_0":               cpb.FHIRVersionCode_V_3_5_0,               // Enum 20
	"V_4_0_0":               cpb.FHIRVersionCode_V_4_0_0,               // Enum 21
	"V_4_0_1":               cpb.FHIRVersionCode_V_4_0_1,               // Enum 22
}

// DefaultFamilyHistoryStatusCodeMap maps from string to cpb.FamilyHistoryStatusCode_Value.
var DefaultFamilyHistoryStatusCodeMap = map[string]cpb.FamilyHistoryStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.FamilyHistoryStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"COMPLETED":             cpb.FamilyHistoryStatusCode_COMPLETED,             // Enum 2
	"ENTERED_IN_ERROR":      cpb.FamilyHistoryStatusCode_ENTERED_IN_ERROR,      // Enum 3
	"HEALTH_UNKNOWN":        cpb.FamilyHistoryStatusCode_HEALTH_UNKNOWN,        // Enum 4
	"PARTIAL":               cpb.FamilyHistoryStatusCode_PARTIAL,               // Enum 1
}

// DefaultFilterOperatorCodeMap maps from string to cpb.FilterOperatorCode_Value.
var DefaultFilterOperatorCodeMap = map[string]cpb.FilterOperatorCode_Value{
	"INVALID_UNINITIALIZED": cpb.FilterOperatorCode_INVALID_UNINITIALIZED, // Enum 0
	"=":                     cpb.FilterOperatorCode_EQUALS,                // Enum 1
	"DESCENDENT_OF":         cpb.FilterOperatorCode_DESCENDENT_OF,         // Enum 3
	"EQUALS":                cpb.FilterOperatorCode_EQUALS,                // Enum 1
	"EXISTS":                cpb.FilterOperatorCode_EXISTS,                // Enum 9
	"GENERALIZES":           cpb.FilterOperatorCode_GENERALIZES,           // Enum 8
	"IN":                    cpb.FilterOperatorCode_IN,                    // Enum 6
	"IS_A":                  cpb.FilterOperatorCode_IS_A,                  // Enum 2
	"IS_NOT_A":              cpb.FilterOperatorCode_IS_NOT_A,              // Enum 4
	"NOT_IN":                cpb.FilterOperatorCode_NOT_IN,                // Enum 7
	"REGEX":                 cpb.FilterOperatorCode_REGEX,                 // Enum 5
}

// DefaultFinancialResourceStatusCodeMap maps from string to cpb.FinancialResourceStatusCode_Value.
var DefaultFinancialResourceStatusCodeMap = map[string]cpb.FinancialResourceStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.FinancialResourceStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.FinancialResourceStatusCode_ACTIVE,                // Enum 1
	"CANCELLED":             cpb.FinancialResourceStatusCode_CANCELLED,             // Enum 2
	"DRAFT":                 cpb.FinancialResourceStatusCode_DRAFT,                 // Enum 3
	"ENTERED_IN_ERROR":      cpb.FinancialResourceStatusCode_ENTERED_IN_ERROR,      // Enum 4
}

// DefaultFlagStatusCodeMap maps from string to cpb.FlagStatusCode_Value.
var DefaultFlagStatusCodeMap = map[string]cpb.FlagStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.FlagStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.FlagStatusCode_ACTIVE,                // Enum 1
	"ENTERED_IN_ERROR":      cpb.FlagStatusCode_ENTERED_IN_ERROR,      // Enum 3
	"INACTIVE":              cpb.FlagStatusCode_INACTIVE,              // Enum 2
}

// DefaultGoalAcceptanceStatusCodeMap maps from string to cpb.GoalAcceptanceStatusCode_Value.
var DefaultGoalAcceptanceStatusCodeMap = map[string]cpb.GoalAcceptanceStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.GoalAcceptanceStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"AGREE":                 cpb.GoalAcceptanceStatusCode_AGREE,                 // Enum 1
	"DISAGREE":              cpb.GoalAcceptanceStatusCode_DISAGREE,              // Enum 2
	"PENDING":               cpb.GoalAcceptanceStatusCode_PENDING,               // Enum 3
}

// DefaultGoalLifecycleStatusCodeMap maps from string to cpb.GoalLifecycleStatusCode_Value.
var DefaultGoalLifecycleStatusCodeMap = map[string]cpb.GoalLifecycleStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.GoalLifecycleStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACCEPTED":              cpb.GoalLifecycleStatusCode_ACCEPTED,              // Enum 3
	"ACTIVE":                cpb.GoalLifecycleStatusCode_ACTIVE,                // Enum 4
	"CANCELLED":             cpb.GoalLifecycleStatusCode_CANCELLED,             // Enum 7
	"COMPLETED":             cpb.GoalLifecycleStatusCode_COMPLETED,             // Enum 6
	"ENTERED_IN_ERROR":      cpb.GoalLifecycleStatusCode_ENTERED_IN_ERROR,      // Enum 8
	"ON_HOLD":               cpb.GoalLifecycleStatusCode_ON_HOLD,               // Enum 5
	"PLANNED":               cpb.GoalLifecycleStatusCode_PLANNED,               // Enum 2
	"PROPOSED":              cpb.GoalLifecycleStatusCode_PROPOSED,              // Enum 1
	"REJECTED":              cpb.GoalLifecycleStatusCode_REJECTED,              // Enum 9
}

// DefaultGraphCompartmentRuleCodeMap maps from string to cpb.GraphCompartmentRuleCode_Value.
var DefaultGraphCompartmentRuleCodeMap = map[string]cpb.GraphCompartmentRuleCode_Value{
	"INVALID_UNINITIALIZED": cpb.GraphCompartmentRuleCode_INVALID_UNINITIALIZED, // Enum 0
	"CUSTOM":                cpb.GraphCompartmentRuleCode_CUSTOM,                // Enum 4
	"DIFFERENT":             cpb.GraphCompartmentRuleCode_DIFFERENT,             // Enum 3
	"IDENTICAL":             cpb.GraphCompartmentRuleCode_IDENTICAL,             // Enum 1
	"MATCHING":              cpb.GraphCompartmentRuleCode_MATCHING,              // Enum 2
}

// DefaultGraphCompartmentUseCodeMap maps from string to cpb.GraphCompartmentUseCode_Value.
var DefaultGraphCompartmentUseCodeMap = map[string]cpb.GraphCompartmentUseCode_Value{
	"INVALID_UNINITIALIZED": cpb.GraphCompartmentUseCode_INVALID_UNINITIALIZED, // Enum 0
	"CONDITION":             cpb.GraphCompartmentUseCode_CONDITION,             // Enum 1
	"REQUIREMENT":           cpb.GraphCompartmentUseCode_REQUIREMENT,           // Enum 2
}

// DefaultGroupMeasureCodeMap maps from string to cpb.GroupMeasureCode_Value.
var DefaultGroupMeasureCodeMap = map[string]cpb.GroupMeasureCode_Value{
	"INVALID_UNINITIALIZED": cpb.GroupMeasureCode_INVALID_UNINITIALIZED, // Enum 0
	"MEAN":                  cpb.GroupMeasureCode_MEAN,                  // Enum 1
	"MEAN_OF_MEAN":          cpb.GroupMeasureCode_MEAN_OF_MEAN,          // Enum 3
	"MEAN_OF_MEDIAN":        cpb.GroupMeasureCode_MEAN_OF_MEDIAN,        // Enum 4
	"MEDIAN":                cpb.GroupMeasureCode_MEDIAN,                // Enum 2
	"MEDIAN_OF_MEAN":        cpb.GroupMeasureCode_MEDIAN_OF_MEAN,        // Enum 5
	"MEDIAN_OF_MEDIAN":      cpb.GroupMeasureCode_MEDIAN_OF_MEDIAN,      // Enum 6
}

// DefaultGroupTypeCodeMap maps from string to cpb.GroupTypeCode_Value.
var DefaultGroupTypeCodeMap = map[string]cpb.GroupTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.GroupTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"ANIMAL":                cpb.GroupTypeCode_ANIMAL,                // Enum 2
	"DEVICE":                cpb.GroupTypeCode_DEVICE,                // Enum 4
	"MEDICATION":            cpb.GroupTypeCode_MEDICATION,            // Enum 5
	"PERSON":                cpb.GroupTypeCode_PERSON,                // Enum 1
	"PRACTITIONER":          cpb.GroupTypeCode_PRACTITIONER,          // Enum 3
	"SUBSTANCE":             cpb.GroupTypeCode_SUBSTANCE,             // Enum 6
}

// DefaultGuidanceResponseStatusCodeMap maps from string to cpb.GuidanceResponseStatusCode_Value.
var DefaultGuidanceResponseStatusCodeMap = map[string]cpb.GuidanceResponseStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.GuidanceResponseStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"DATA_REQUESTED":        cpb.GuidanceResponseStatusCode_DATA_REQUESTED,        // Enum 2
	"DATA_REQUIRED":         cpb.GuidanceResponseStatusCode_DATA_REQUIRED,         // Enum 3
	"ENTERED_IN_ERROR":      cpb.GuidanceResponseStatusCode_ENTERED_IN_ERROR,      // Enum 6
	"FAILURE":               cpb.GuidanceResponseStatusCode_FAILURE,               // Enum 5
	"IN_PROGRESS":           cpb.GuidanceResponseStatusCode_IN_PROGRESS,           // Enum 4
	"SUCCESS":               cpb.GuidanceResponseStatusCode_SUCCESS,               // Enum 1
}

// DefaultGuidePageGenerationCodeMap maps from string to cpb.GuidePageGenerationCode_Value.
var DefaultGuidePageGenerationCodeMap = map[string]cpb.GuidePageGenerationCode_Value{
	"INVALID_UNINITIALIZED": cpb.GuidePageGenerationCode_INVALID_UNINITIALIZED, // Enum 0
	"GENERATED":             cpb.GuidePageGenerationCode_GENERATED,             // Enum 4
	"HTML":                  cpb.GuidePageGenerationCode_HTML,                  // Enum 1
	"MARKDOWN":              cpb.GuidePageGenerationCode_MARKDOWN,              // Enum 2
	"XML":                   cpb.GuidePageGenerationCode_XML,                   // Enum 3
}

// DefaultGuideParameterCodeMap maps from string to cpb.GuideParameterCode_Value.
var DefaultGuideParameterCodeMap = map[string]cpb.GuideParameterCode_Value{
	"INVALID_UNINITIALIZED": cpb.GuideParameterCode_INVALID_UNINITIALIZED, // Enum 0
	"APPLY":                 cpb.GuideParameterCode_APPLY,                 // Enum 1
	"EXPANSION_PARAMETER":   cpb.GuideParameterCode_EXPANSION_PARAMETER,   // Enum 5
	"GENERATE_JSON":         cpb.GuideParameterCode_GENERATE_JSON,         // Enum 8
	"GENERATE_TURTLE":       cpb.GuideParameterCode_GENERATE_TURTLE,       // Enum 9
	"GENERATE_XML":          cpb.GuideParameterCode_GENERATE_XML,          // Enum 7
	"HTML_TEMPLATE":         cpb.GuideParameterCode_HTML_TEMPLATE,         // Enum 10
	"PATH_PAGES":            cpb.GuideParameterCode_PATH_PAGES,            // Enum 3
	"PATH_RESOURCE":         cpb.GuideParameterCode_PATH_RESOURCE,         // Enum 2
	"PATH_TX_CACHE":         cpb.GuideParameterCode_PATH_TX_CACHE,         // Enum 4
	"RULE_BROKEN_LINKS":     cpb.GuideParameterCode_RULE_BROKEN_LINKS,     // Enum 6
}

// DefaultHL7WorkgroupCodeMap maps from string to cpb.HL7WorkgroupCode_Value.
var DefaultHL7WorkgroupCodeMap = map[string]cpb.HL7WorkgroupCode_Value{
	"INVALID_UNINITIALIZED": cpb.HL7WorkgroupCode_INVALID_UNINITIALIZED, // Enum 0
	"AID":                   cpb.HL7WorkgroupCode_AID,                   // Enum 24
	"BRR":                   cpb.HL7WorkgroupCode_BRR,                   // Enum 19
	"CBCC":                  cpb.HL7WorkgroupCode_CBCC,                  // Enum 1
	"CDS":                   cpb.HL7WorkgroupCode_CDS,                   // Enum 2
	"CG":                    cpb.HL7WorkgroupCode_CG,                    // Enum 4
	"CQI":                   cpb.HL7WorkgroupCode_CQI,                   // Enum 3
	"DEV":                   cpb.HL7WorkgroupCode_DEV,                   // Enum 5
	"EHR":                   cpb.HL7WorkgroupCode_EHR,                   // Enum 6
	"FHIR":                  cpb.HL7WorkgroupCode_FHIR,                  // Enum 7
	"FM":                    cpb.HL7WorkgroupCode_FM,                    // Enum 8
	"HSI":                   cpb.HL7WorkgroupCode_HSI,                   // Enum 9
	"II":                    cpb.HL7WorkgroupCode_II,                    // Enum 10
	"INM":                   cpb.HL7WorkgroupCode_INM,                   // Enum 11
	"ITS":                   cpb.HL7WorkgroupCode_ITS,                   // Enum 12
	"MNM":                   cpb.HL7WorkgroupCode_MNM,                   // Enum 13
	"OO":                    cpb.HL7WorkgroupCode_OO,                    // Enum 14
	"PA":                    cpb.HL7WorkgroupCode_PA,                    // Enum 15
	"PC":                    cpb.HL7WorkgroupCode_PC,                    // Enum 16
	"PHER":                  cpb.HL7WorkgroupCode_PHER,                  // Enum 17
	"PHX":                   cpb.HL7WorkgroupCode_PHX,                   // Enum 18
	"SD":                    cpb.HL7WorkgroupCode_SD,                    // Enum 20
	"SEC":                   cpb.HL7WorkgroupCode_SEC,                   // Enum 21
	"US":                    cpb.HL7WorkgroupCode_US,                    // Enum 22
	"VOCAB":                 cpb.HL7WorkgroupCode_VOCAB,                 // Enum 23
}

// DefaultHTTPVerbCodeMap maps from string to cpb.HTTPVerbCode_Value.
var DefaultHTTPVerbCodeMap = map[string]cpb.HTTPVerbCode_Value{
	"INVALID_UNINITIALIZED": cpb.HTTPVerbCode_INVALID_UNINITIALIZED, // Enum 0
	"DELETE":                cpb.HTTPVerbCode_DELETE,                // Enum 5
	"GET":                   cpb.HTTPVerbCode_GET,                   // Enum 1
	"HEAD":                  cpb.HTTPVerbCode_HEAD,                  // Enum 2
	"PATCH":                 cpb.HTTPVerbCode_PATCH,                 // Enum 6
	"POST":                  cpb.HTTPVerbCode_POST,                  // Enum 3
	"PUT":                   cpb.HTTPVerbCode_PUT,                   // Enum 4
}

// DefaultHumanNameAssemblyOrderCodeMap maps from string to cpb.HumanNameAssemblyOrderCode_Value.
var DefaultHumanNameAssemblyOrderCodeMap = map[string]cpb.HumanNameAssemblyOrderCode_Value{
	"INVALID_UNINITIALIZED": cpb.HumanNameAssemblyOrderCode_INVALID_UNINITIALIZED, // Enum 0
	"NL1":                   cpb.HumanNameAssemblyOrderCode_NL1,                   // Enum 1
	"NL2":                   cpb.HumanNameAssemblyOrderCode_NL2,                   // Enum 2
	"NL3":                   cpb.HumanNameAssemblyOrderCode_NL3,                   // Enum 3
	"NL4":                   cpb.HumanNameAssemblyOrderCode_NL4,                   // Enum 4
}

// DefaultIdentifierUseCodeMap maps from string to cpb.IdentifierUseCode_Value.
var DefaultIdentifierUseCodeMap = map[string]cpb.IdentifierUseCode_Value{
	"INVALID_UNINITIALIZED": cpb.IdentifierUseCode_INVALID_UNINITIALIZED, // Enum 0
	"OFFICIAL":              cpb.IdentifierUseCode_OFFICIAL,              // Enum 2
	"OLD":                   cpb.IdentifierUseCode_OLD,                   // Enum 5
	"SECONDARY":             cpb.IdentifierUseCode_SECONDARY,             // Enum 4
	"TEMP":                  cpb.IdentifierUseCode_TEMP,                  // Enum 3
	"USUAL":                 cpb.IdentifierUseCode_USUAL,                 // Enum 1
}

// DefaultIdentityAssuranceLevelCodeMap maps from string to cpb.IdentityAssuranceLevelCode_Value.
var DefaultIdentityAssuranceLevelCodeMap = map[string]cpb.IdentityAssuranceLevelCode_Value{
	"INVALID_UNINITIALIZED": cpb.IdentityAssuranceLevelCode_INVALID_UNINITIALIZED, // Enum 0
	"LEVEL1":                cpb.IdentityAssuranceLevelCode_LEVEL1,                // Enum 1
	"LEVEL2":                cpb.IdentityAssuranceLevelCode_LEVEL2,                // Enum 2
	"LEVEL3":                cpb.IdentityAssuranceLevelCode_LEVEL3,                // Enum 3
	"LEVEL4":                cpb.IdentityAssuranceLevelCode_LEVEL4,                // Enum 4
}

// DefaultImagingStudyStatusCodeMap maps from string to cpb.ImagingStudyStatusCode_Value.
var DefaultImagingStudyStatusCodeMap = map[string]cpb.ImagingStudyStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ImagingStudyStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"AVAILABLE":             cpb.ImagingStudyStatusCode_AVAILABLE,             // Enum 2
	"CANCELLED":             cpb.ImagingStudyStatusCode_CANCELLED,             // Enum 3
	"ENTERED_IN_ERROR":      cpb.ImagingStudyStatusCode_ENTERED_IN_ERROR,      // Enum 4
	"REGISTERED":            cpb.ImagingStudyStatusCode_REGISTERED,            // Enum 1
	"UNKNOWN":               cpb.ImagingStudyStatusCode_UNKNOWN,               // Enum 5
}

// DefaultImplantStatusCodeMap maps from string to cpb.ImplantStatusCode_Value.
var DefaultImplantStatusCodeMap = map[string]cpb.ImplantStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ImplantStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"DISABLED":              cpb.ImplantStatusCode_DISABLED,              // Enum 3
	"FUNCTIONAL":            cpb.ImplantStatusCode_FUNCTIONAL,            // Enum 1
	"NON_FUNCTIONAL":        cpb.ImplantStatusCode_NON_FUNCTIONAL,        // Enum 2
	"UNKNOWN":               cpb.ImplantStatusCode_UNKNOWN,               // Enum 4
}

// DefaultInvoicePriceComponentTypeCodeMap maps from string to cpb.InvoicePriceComponentTypeCode_Value.
var DefaultInvoicePriceComponentTypeCodeMap = map[string]cpb.InvoicePriceComponentTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.InvoicePriceComponentTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"BASE":                  cpb.InvoicePriceComponentTypeCode_BASE,                  // Enum 1
	"DEDUCTION":             cpb.InvoicePriceComponentTypeCode_DEDUCTION,             // Enum 3
	"DISCOUNT":              cpb.InvoicePriceComponentTypeCode_DISCOUNT,              // Enum 4
	"INFORMATIONAL":         cpb.InvoicePriceComponentTypeCode_INFORMATIONAL,         // Enum 6
	"SURCHARGE":             cpb.InvoicePriceComponentTypeCode_SURCHARGE,             // Enum 2
	"TAX":                   cpb.InvoicePriceComponentTypeCode_TAX,                   // Enum 5
}

// DefaultInvoiceStatusCodeMap maps from string to cpb.InvoiceStatusCode_Value.
var DefaultInvoiceStatusCodeMap = map[string]cpb.InvoiceStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.InvoiceStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"BALANCED":              cpb.InvoiceStatusCode_BALANCED,              // Enum 3
	"CANCELLED":             cpb.InvoiceStatusCode_CANCELLED,             // Enum 4
	"DRAFT":                 cpb.InvoiceStatusCode_DRAFT,                 // Enum 1
	"ENTERED_IN_ERROR":      cpb.InvoiceStatusCode_ENTERED_IN_ERROR,      // Enum 5
	"ISSUED":                cpb.InvoiceStatusCode_ISSUED,                // Enum 2
}

// DefaultIssueSeverityCodeMap maps from string to cpb.IssueSeverityCode_Value.
var DefaultIssueSeverityCodeMap = map[string]cpb.IssueSeverityCode_Value{
	"INVALID_UNINITIALIZED": cpb.IssueSeverityCode_INVALID_UNINITIALIZED, // Enum 0
	"ERROR":                 cpb.IssueSeverityCode_ERROR,                 // Enum 2
	"FATAL":                 cpb.IssueSeverityCode_FATAL,                 // Enum 1
	"INFORMATION":           cpb.IssueSeverityCode_INFORMATION,           // Enum 4
	"WARNING":               cpb.IssueSeverityCode_WARNING,               // Enum 3
}

// DefaultIssueTypeCodeMap maps from string to cpb.IssueTypeCode_Value.
var DefaultIssueTypeCodeMap = map[string]cpb.IssueTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.IssueTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"BUSINESS_RULE":         cpb.IssueTypeCode_BUSINESS_RULE,         // Enum 22
	"CODE_INVALID":          cpb.IssueTypeCode_CODE_INVALID,          // Enum 19
	"CONFLICT":              cpb.IssueTypeCode_CONFLICT,              // Enum 23
	"DELETED":               cpb.IssueTypeCode_DELETED,               // Enum 17
	"DUPLICATE":             cpb.IssueTypeCode_DUPLICATE,             // Enum 14
	"EXCEPTION":             cpb.IssueTypeCode_EXCEPTION,             // Enum 27
	"EXPIRED":               cpb.IssueTypeCode_EXPIRED,               // Enum 9
	"EXTENSION":             cpb.IssueTypeCode_EXTENSION,             // Enum 20
	"FORBIDDEN":             cpb.IssueTypeCode_FORBIDDEN,             // Enum 10
	"INCOMPLETE":            cpb.IssueTypeCode_INCOMPLETE,            // Enum 29
	"INFORMATIONAL":         cpb.IssueTypeCode_INFORMATIONAL,         // Enum 31
	"INVALID":               cpb.IssueTypeCode_INVALID,               // Enum 1
	"INVARIANT":             cpb.IssueTypeCode_INVARIANT,             // Enum 5
	"LOCK_ERROR":            cpb.IssueTypeCode_LOCK_ERROR,            // Enum 25
	"LOGIN":                 cpb.IssueTypeCode_LOGIN,                 // Enum 7
	"MULTIPLE_MATCHES":      cpb.IssueTypeCode_MULTIPLE_MATCHES,      // Enum 15
	"NOT_FOUND":             cpb.IssueTypeCode_NOT_FOUND,             // Enum 16
	"NOT_SUPPORTED":         cpb.IssueTypeCode_NOT_SUPPORTED,         // Enum 13
	"NO_STORE":              cpb.IssueTypeCode_NO_STORE,              // Enum 26
	"PROCESSING":            cpb.IssueTypeCode_PROCESSING,            // Enum 12
	"REQUIRED":              cpb.IssueTypeCode_REQUIRED,              // Enum 3
	"SECURITY":              cpb.IssueTypeCode_SECURITY,              // Enum 6
	"STRUCTURE":             cpb.IssueTypeCode_STRUCTURE,             // Enum 2
	"SUPPRESSED":            cpb.IssueTypeCode_SUPPRESSED,            // Enum 11
	"THROTTLED":             cpb.IssueTypeCode_THROTTLED,             // Enum 30
	"TIMEOUT":               cpb.IssueTypeCode_TIMEOUT,               // Enum 28
	"TOO_COSTLY":            cpb.IssueTypeCode_TOO_COSTLY,            // Enum 21
	"TOO_LONG":              cpb.IssueTypeCode_TOO_LONG,              // Enum 18
	"TRANSIENT":             cpb.IssueTypeCode_TRANSIENT,             // Enum 24
	"UNKNOWN":               cpb.IssueTypeCode_UNKNOWN,               // Enum 8
	"VALUE":                 cpb.IssueTypeCode_VALUE,                 // Enum 4
}

// DefaultLinkTypeCodeMap maps from string to cpb.LinkTypeCode_Value.
var DefaultLinkTypeCodeMap = map[string]cpb.LinkTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.LinkTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"REFER":                 cpb.LinkTypeCode_REFER,                 // Enum 3
	"REPLACED_BY":           cpb.LinkTypeCode_REPLACED_BY,           // Enum 1
	"REPLACES":              cpb.LinkTypeCode_REPLACES,              // Enum 2
	"SEEALSO":               cpb.LinkTypeCode_SEEALSO,               // Enum 4
}

// DefaultLinkageTypeCodeMap maps from string to cpb.LinkageTypeCode_Value.
var DefaultLinkageTypeCodeMap = map[string]cpb.LinkageTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.LinkageTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"ALTERNATE":             cpb.LinkageTypeCode_ALTERNATE,             // Enum 2
	"HISTORICAL":            cpb.LinkageTypeCode_HISTORICAL,            // Enum 3
	"SOURCE":                cpb.LinkageTypeCode_SOURCE,                // Enum 1
}

// DefaultListModeCodeMap maps from string to cpb.ListModeCode_Value.
var DefaultListModeCodeMap = map[string]cpb.ListModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ListModeCode_INVALID_UNINITIALIZED, // Enum 0
	"CHANGES":               cpb.ListModeCode_CHANGES,               // Enum 3
	"SNAPSHOT":              cpb.ListModeCode_SNAPSHOT,              // Enum 2
	"WORKING":               cpb.ListModeCode_WORKING,               // Enum 1
}

// DefaultListStatusCodeMap maps from string to cpb.ListStatusCode_Value.
var DefaultListStatusCodeMap = map[string]cpb.ListStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ListStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"CURRENT":               cpb.ListStatusCode_CURRENT,               // Enum 1
	"ENTERED_IN_ERROR":      cpb.ListStatusCode_ENTERED_IN_ERROR,      // Enum 3
	"RETIRED":               cpb.ListStatusCode_RETIRED,               // Enum 2
}

// DefaultLocationModeCodeMap maps from string to cpb.LocationModeCode_Value.
var DefaultLocationModeCodeMap = map[string]cpb.LocationModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.LocationModeCode_INVALID_UNINITIALIZED, // Enum 0
	"INSTANCE":              cpb.LocationModeCode_INSTANCE,              // Enum 1
	"KIND":                  cpb.LocationModeCode_KIND,                  // Enum 2
}

// DefaultLocationStatusCodeMap maps from string to cpb.LocationStatusCode_Value.
var DefaultLocationStatusCodeMap = map[string]cpb.LocationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.LocationStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.LocationStatusCode_ACTIVE,                // Enum 1
	"INACTIVE":              cpb.LocationStatusCode_INACTIVE,              // Enum 3
	"SUSPENDED":             cpb.LocationStatusCode_SUSPENDED,             // Enum 2
}

// DefaultMatchGradeCodeMap maps from string to cpb.MatchGradeCode_Value.
var DefaultMatchGradeCodeMap = map[string]cpb.MatchGradeCode_Value{
	"INVALID_UNINITIALIZED": cpb.MatchGradeCode_INVALID_UNINITIALIZED, // Enum 0
	"CERTAIN":               cpb.MatchGradeCode_CERTAIN,               // Enum 1
	"CERTAINLY_NOT":         cpb.MatchGradeCode_CERTAINLY_NOT,         // Enum 4
	"POSSIBLE":              cpb.MatchGradeCode_POSSIBLE,              // Enum 3
	"PROBABLE":              cpb.MatchGradeCode_PROBABLE,              // Enum 2
}

// DefaultMeasureImprovementNotationCodeMap maps from string to cpb.MeasureImprovementNotationCode_Value.
var DefaultMeasureImprovementNotationCodeMap = map[string]cpb.MeasureImprovementNotationCode_Value{
	"INVALID_UNINITIALIZED": cpb.MeasureImprovementNotationCode_INVALID_UNINITIALIZED, // Enum 0
	"DECREASE":              cpb.MeasureImprovementNotationCode_DECREASE,              // Enum 2
	"INCREASE":              cpb.MeasureImprovementNotationCode_INCREASE,              // Enum 1
}

// DefaultMeasureReportStatusCodeMap maps from string to cpb.MeasureReportStatusCode_Value.
var DefaultMeasureReportStatusCodeMap = map[string]cpb.MeasureReportStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.MeasureReportStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"COMPLETE":              cpb.MeasureReportStatusCode_COMPLETE,              // Enum 1
	"ERROR":                 cpb.MeasureReportStatusCode_ERROR,                 // Enum 3
	"PENDING":               cpb.MeasureReportStatusCode_PENDING,               // Enum 2
}

// DefaultMeasureReportTypeCodeMap maps from string to cpb.MeasureReportTypeCode_Value.
var DefaultMeasureReportTypeCodeMap = map[string]cpb.MeasureReportTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.MeasureReportTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"DATA_COLLECTION":       cpb.MeasureReportTypeCode_DATA_COLLECTION,       // Enum 4
	"INDIVIDUAL":            cpb.MeasureReportTypeCode_INDIVIDUAL,            // Enum 1
	"SUBJECT_LIST":          cpb.MeasureReportTypeCode_SUBJECT_LIST,          // Enum 2
	"SUMMARY":               cpb.MeasureReportTypeCode_SUMMARY,               // Enum 3
}

// DefaultMedicationAdministrationStatusCodeMap maps from string to cpb.MedicationAdministrationStatusCode_Value.
var DefaultMedicationAdministrationStatusCodeMap = map[string]cpb.MedicationAdministrationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.MedicationAdministrationStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"COMPLETED":             cpb.MedicationAdministrationStatusCode_COMPLETED,             // Enum 4
	"ENTERED_IN_ERROR":      cpb.MedicationAdministrationStatusCode_ENTERED_IN_ERROR,      // Enum 5
	"IN_PROGRESS":           cpb.MedicationAdministrationStatusCode_IN_PROGRESS,           // Enum 1
	"NOT_DONE":              cpb.MedicationAdministrationStatusCode_NOT_DONE,              // Enum 2
	"ON_HOLD":               cpb.MedicationAdministrationStatusCode_ON_HOLD,               // Enum 3
	"STOPPED":               cpb.MedicationAdministrationStatusCode_STOPPED,               // Enum 6
	"UNKNOWN":               cpb.MedicationAdministrationStatusCode_UNKNOWN,               // Enum 7
}

// DefaultMedicationDispenseStatusCodeMap maps from string to cpb.MedicationDispenseStatusCode_Value.
var DefaultMedicationDispenseStatusCodeMap = map[string]cpb.MedicationDispenseStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.MedicationDispenseStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"CANCELLED":             cpb.MedicationDispenseStatusCode_CANCELLED,             // Enum 3
	"COMPLETED":             cpb.MedicationDispenseStatusCode_COMPLETED,             // Enum 5
	"DECLINED":              cpb.MedicationDispenseStatusCode_DECLINED,              // Enum 8
	"ENTERED_IN_ERROR":      cpb.MedicationDispenseStatusCode_ENTERED_IN_ERROR,      // Enum 6
	"IN_PROGRESS":           cpb.MedicationDispenseStatusCode_IN_PROGRESS,           // Enum 2
	"ON_HOLD":               cpb.MedicationDispenseStatusCode_ON_HOLD,               // Enum 4
	"PREPARATION":           cpb.MedicationDispenseStatusCode_PREPARATION,           // Enum 1
	"STOPPED":               cpb.MedicationDispenseStatusCode_STOPPED,               // Enum 7
	"UNKNOWN":               cpb.MedicationDispenseStatusCode_UNKNOWN,               // Enum 9
}

// DefaultMedicationKnowledgeStatusCodeMap maps from string to cpb.MedicationKnowledgeStatusCode_Value.
var DefaultMedicationKnowledgeStatusCodeMap = map[string]cpb.MedicationKnowledgeStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.MedicationKnowledgeStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.MedicationKnowledgeStatusCode_ACTIVE,                // Enum 1
	"ENTERED_IN_ERROR":      cpb.MedicationKnowledgeStatusCode_ENTERED_IN_ERROR,      // Enum 3
	"INACTIVE":              cpb.MedicationKnowledgeStatusCode_INACTIVE,              // Enum 2
}

// DefaultMedicationRequestIntentCodeMap maps from string to cpb.MedicationRequestIntentCode_Value.
var DefaultMedicationRequestIntentCodeMap = map[string]cpb.MedicationRequestIntentCode_Value{
	"INVALID_UNINITIALIZED": cpb.MedicationRequestIntentCode_INVALID_UNINITIALIZED, // Enum 0
	"FILLER_ORDER":          cpb.MedicationRequestIntentCode_FILLER_ORDER,          // Enum 6
	"INSTANCE_ORDER":        cpb.MedicationRequestIntentCode_INSTANCE_ORDER,        // Enum 7
	"OPTION":                cpb.MedicationRequestIntentCode_OPTION,                // Enum 8
	"ORDER":                 cpb.MedicationRequestIntentCode_ORDER,                 // Enum 3
	"ORIGINAL_ORDER":        cpb.MedicationRequestIntentCode_ORIGINAL_ORDER,        // Enum 4
	"PLAN":                  cpb.MedicationRequestIntentCode_PLAN,                  // Enum 2
	"PROPOSAL":              cpb.MedicationRequestIntentCode_PROPOSAL,              // Enum 1
	"REFLEX_ORDER":          cpb.MedicationRequestIntentCode_REFLEX_ORDER,          // Enum 5
}

// DefaultMedicationStatementStatusCodesMap maps from string to cpb.MedicationStatementStatusCodes_Value.
var DefaultMedicationStatementStatusCodesMap = map[string]cpb.MedicationStatementStatusCodes_Value{
	"INVALID_UNINITIALIZED": cpb.MedicationStatementStatusCodes_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.MedicationStatementStatusCodes_ACTIVE,                // Enum 1
	"COMPLETED":             cpb.MedicationStatementStatusCodes_COMPLETED,             // Enum 2
	"ENTERED_IN_ERROR":      cpb.MedicationStatementStatusCodes_ENTERED_IN_ERROR,      // Enum 3
	"INTENDED":              cpb.MedicationStatementStatusCodes_INTENDED,              // Enum 4
	"NOT_TAKEN":             cpb.MedicationStatementStatusCodes_NOT_TAKEN,             // Enum 8
	"ON_HOLD":               cpb.MedicationStatementStatusCodes_ON_HOLD,               // Enum 6
	"STOPPED":               cpb.MedicationStatementStatusCodes_STOPPED,               // Enum 5
	"UNKNOWN":               cpb.MedicationStatementStatusCodes_UNKNOWN,               // Enum 7
}

// DefaultMedicationStatusCodeMap maps from string to cpb.MedicationStatusCode_Value.
var DefaultMedicationStatusCodeMap = map[string]cpb.MedicationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.MedicationStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.MedicationStatusCode_ACTIVE,                // Enum 1
	"ENTERED_IN_ERROR":      cpb.MedicationStatusCode_ENTERED_IN_ERROR,      // Enum 3
	"INACTIVE":              cpb.MedicationStatusCode_INACTIVE,              // Enum 2
}

// DefaultMedicationrequestStatusCodeMap maps from string to cpb.MedicationrequestStatusCode_Value.
var DefaultMedicationrequestStatusCodeMap = map[string]cpb.MedicationrequestStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.MedicationrequestStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.MedicationrequestStatusCode_ACTIVE,                // Enum 1
	"CANCELLED":             cpb.MedicationrequestStatusCode_CANCELLED,             // Enum 3
	"COMPLETED":             cpb.MedicationrequestStatusCode_COMPLETED,             // Enum 4
	"DRAFT":                 cpb.MedicationrequestStatusCode_DRAFT,                 // Enum 7
	"ENTERED_IN_ERROR":      cpb.MedicationrequestStatusCode_ENTERED_IN_ERROR,      // Enum 5
	"ON_HOLD":               cpb.MedicationrequestStatusCode_ON_HOLD,               // Enum 2
	"STOPPED":               cpb.MedicationrequestStatusCode_STOPPED,               // Enum 6
	"UNKNOWN":               cpb.MedicationrequestStatusCode_UNKNOWN,               // Enum 8
}

// DefaultMessageSignificanceCategoryCodeMap maps from string to cpb.MessageSignificanceCategoryCode_Value.
var DefaultMessageSignificanceCategoryCodeMap = map[string]cpb.MessageSignificanceCategoryCode_Value{
	"INVALID_UNINITIALIZED": cpb.MessageSignificanceCategoryCode_INVALID_UNINITIALIZED, // Enum 0
	"CONSEQUENCE":           cpb.MessageSignificanceCategoryCode_CONSEQUENCE,           // Enum 1
	"CURRENCY":              cpb.MessageSignificanceCategoryCode_CURRENCY,              // Enum 2
	"NOTIFICATION":          cpb.MessageSignificanceCategoryCode_NOTIFICATION,          // Enum 3
}

// DefaultMessageheaderResponseRequestCodeMap maps from string to cpb.MessageheaderResponseRequestCode_Value.
var DefaultMessageheaderResponseRequestCodeMap = map[string]cpb.MessageheaderResponseRequestCode_Value{
	"INVALID_UNINITIALIZED": cpb.MessageheaderResponseRequestCode_INVALID_UNINITIALIZED, // Enum 0
	"ALWAYS":                cpb.MessageheaderResponseRequestCode_ALWAYS,                // Enum 1
	"NEVER":                 cpb.MessageheaderResponseRequestCode_NEVER,                 // Enum 3
	"ON_ERROR":              cpb.MessageheaderResponseRequestCode_ON_ERROR,              // Enum 2
	"ON_SUCCESS":            cpb.MessageheaderResponseRequestCode_ON_SUCCESS,            // Enum 4
}

// DefaultNameUseCodeMap maps from string to cpb.NameUseCode_Value.
var DefaultNameUseCodeMap = map[string]cpb.NameUseCode_Value{
	"INVALID_UNINITIALIZED": cpb.NameUseCode_INVALID_UNINITIALIZED, // Enum 0
	"ANONYMOUS":             cpb.NameUseCode_ANONYMOUS,             // Enum 5
	"MAIDEN":                cpb.NameUseCode_MAIDEN,                // Enum 7
	"NICKNAME":              cpb.NameUseCode_NICKNAME,              // Enum 4
	"OFFICIAL":              cpb.NameUseCode_OFFICIAL,              // Enum 2
	"OLD":                   cpb.NameUseCode_OLD,                   // Enum 6
	"TEMP":                  cpb.NameUseCode_TEMP,                  // Enum 3
	"USUAL":                 cpb.NameUseCode_USUAL,                 // Enum 1
}

// DefaultNamingSystemIdentifierTypeCodeMap maps from string to cpb.NamingSystemIdentifierTypeCode_Value.
var DefaultNamingSystemIdentifierTypeCodeMap = map[string]cpb.NamingSystemIdentifierTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.NamingSystemIdentifierTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"OID":                   cpb.NamingSystemIdentifierTypeCode_OID,                   // Enum 1
	"OTHER":                 cpb.NamingSystemIdentifierTypeCode_OTHER,                 // Enum 4
	"URI":                   cpb.NamingSystemIdentifierTypeCode_URI,                   // Enum 3
	"UUID":                  cpb.NamingSystemIdentifierTypeCode_UUID,                  // Enum 2
}

// DefaultNamingSystemTypeCodeMap maps from string to cpb.NamingSystemTypeCode_Value.
var DefaultNamingSystemTypeCodeMap = map[string]cpb.NamingSystemTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.NamingSystemTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"CODESYSTEM":            cpb.NamingSystemTypeCode_CODESYSTEM,            // Enum 1
	"IDENTIFIER":            cpb.NamingSystemTypeCode_IDENTIFIER,            // Enum 2
	"ROOT":                  cpb.NamingSystemTypeCode_ROOT,                  // Enum 3
}

// DefaultNarrativeStatusCodeMap maps from string to cpb.NarrativeStatusCode_Value.
var DefaultNarrativeStatusCodeMap = map[string]cpb.NarrativeStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.NarrativeStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ADDITIONAL":            cpb.NarrativeStatusCode_ADDITIONAL,            // Enum 3
	"EMPTY":                 cpb.NarrativeStatusCode_EMPTY,                 // Enum 4
	"EXTENSIONS":            cpb.NarrativeStatusCode_EXTENSIONS,            // Enum 2
	"GENERATED":             cpb.NarrativeStatusCode_GENERATED,             // Enum 1
}

// DefaultNoteTypeCodeMap maps from string to cpb.NoteTypeCode_Value.
var DefaultNoteTypeCodeMap = map[string]cpb.NoteTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.NoteTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"DISPLAY":               cpb.NoteTypeCode_DISPLAY,               // Enum 1
	"PRINT":                 cpb.NoteTypeCode_PRINT,                 // Enum 2
	"PRINTOPER":             cpb.NoteTypeCode_PRINTOPER,             // Enum 3
}

// DefaultObservationDataTypeCodeMap maps from string to cpb.ObservationDataTypeCode_Value.
var DefaultObservationDataTypeCodeMap = map[string]cpb.ObservationDataTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ObservationDataTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"BOOLEAN":               cpb.ObservationDataTypeCode_BOOLEAN,               // Enum 4
	"CODEABLE_CONCEPT":      cpb.ObservationDataTypeCode_CODEABLE_CONCEPT,      // Enum 2
	"DATE_TIME":             cpb.ObservationDataTypeCode_DATE_TIME,             // Enum 10
	"INTEGER":               cpb.ObservationDataTypeCode_INTEGER,               // Enum 5
	"PERIOD":                cpb.ObservationDataTypeCode_PERIOD,                // Enum 11
	"QUANTITY":              cpb.ObservationDataTypeCode_QUANTITY,              // Enum 1
	"RANGE":                 cpb.ObservationDataTypeCode_RANGE,                 // Enum 6
	"RATIO":                 cpb.ObservationDataTypeCode_RATIO,                 // Enum 7
	"SAMPLED_DATA":          cpb.ObservationDataTypeCode_SAMPLED_DATA,          // Enum 8
	"STRING":                cpb.ObservationDataTypeCode_STRING,                // Enum 3
	"TIME":                  cpb.ObservationDataTypeCode_TIME,                  // Enum 9
}

// DefaultObservationRangeCategoryCodeMap maps from string to cpb.ObservationRangeCategoryCode_Value.
var DefaultObservationRangeCategoryCodeMap = map[string]cpb.ObservationRangeCategoryCode_Value{
	"INVALID_UNINITIALIZED": cpb.ObservationRangeCategoryCode_INVALID_UNINITIALIZED, // Enum 0
	"ABSOLUTE":              cpb.ObservationRangeCategoryCode_ABSOLUTE,              // Enum 3
	"CRITICAL":              cpb.ObservationRangeCategoryCode_CRITICAL,              // Enum 2
	"REFERENCE":             cpb.ObservationRangeCategoryCode_REFERENCE,             // Enum 1
}

// DefaultObservationStatusCodeMap maps from string to cpb.ObservationStatusCode_Value.
var DefaultObservationStatusCodeMap = map[string]cpb.ObservationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ObservationStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"AMENDED":               cpb.ObservationStatusCode_AMENDED,               // Enum 4
	"CANCELLED":             cpb.ObservationStatusCode_CANCELLED,             // Enum 6
	"CORRECTED":             cpb.ObservationStatusCode_CORRECTED,             // Enum 5
	"ENTERED_IN_ERROR":      cpb.ObservationStatusCode_ENTERED_IN_ERROR,      // Enum 7
	"FINAL":                 cpb.ObservationStatusCode_FINAL,                 // Enum 3
	"PRELIMINARY":           cpb.ObservationStatusCode_PRELIMINARY,           // Enum 2
	"REGISTERED":            cpb.ObservationStatusCode_REGISTERED,            // Enum 1
	"UNKNOWN":               cpb.ObservationStatusCode_UNKNOWN,               // Enum 8
}

// DefaultOperationKindCodeMap maps from string to cpb.OperationKindCode_Value.
var DefaultOperationKindCodeMap = map[string]cpb.OperationKindCode_Value{
	"INVALID_UNINITIALIZED": cpb.OperationKindCode_INVALID_UNINITIALIZED, // Enum 0
	"OPERATION":             cpb.OperationKindCode_OPERATION,             // Enum 1
	"QUERY":                 cpb.OperationKindCode_QUERY,                 // Enum 2
}

// DefaultOperationParameterUseCodeMap maps from string to cpb.OperationParameterUseCode_Value.
var DefaultOperationParameterUseCodeMap = map[string]cpb.OperationParameterUseCode_Value{
	"INVALID_UNINITIALIZED": cpb.OperationParameterUseCode_INVALID_UNINITIALIZED, // Enum 0
	"IN":                    cpb.OperationParameterUseCode_IN,                    // Enum 1
	"OUT":                   cpb.OperationParameterUseCode_OUT,                   // Enum 2
}

// DefaultOrientationTypeCodeMap maps from string to cpb.OrientationTypeCode_Value.
var DefaultOrientationTypeCodeMap = map[string]cpb.OrientationTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.OrientationTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"ANTISENSE":             cpb.OrientationTypeCode_ANTISENSE,             // Enum 2
	"SENSE":                 cpb.OrientationTypeCode_SENSE,                 // Enum 1
}

// DefaultParticipantRequiredCodeMap maps from string to cpb.ParticipantRequiredCode_Value.
var DefaultParticipantRequiredCodeMap = map[string]cpb.ParticipantRequiredCode_Value{
	"INVALID_UNINITIALIZED": cpb.ParticipantRequiredCode_INVALID_UNINITIALIZED, // Enum 0
	"INFORMATION_ONLY":      cpb.ParticipantRequiredCode_INFORMATION_ONLY,      // Enum 3
	"OPTIONAL":              cpb.ParticipantRequiredCode_OPTIONAL,              // Enum 2
	"REQUIRED":              cpb.ParticipantRequiredCode_REQUIRED,              // Enum 1
}

// DefaultParticipationStatusCodeMap maps from string to cpb.ParticipationStatusCode_Value.
var DefaultParticipationStatusCodeMap = map[string]cpb.ParticipationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ParticipationStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACCEPTED":              cpb.ParticipationStatusCode_ACCEPTED,              // Enum 1
	"DECLINED":              cpb.ParticipationStatusCode_DECLINED,              // Enum 2
	"NEEDS_ACTION":          cpb.ParticipationStatusCode_NEEDS_ACTION,          // Enum 4
	"TENTATIVE":             cpb.ParticipationStatusCode_TENTATIVE,             // Enum 3
}

// DefaultPropertyRepresentationCodeMap maps from string to cpb.PropertyRepresentationCode_Value.
var DefaultPropertyRepresentationCodeMap = map[string]cpb.PropertyRepresentationCode_Value{
	"INVALID_UNINITIALIZED": cpb.PropertyRepresentationCode_INVALID_UNINITIALIZED, // Enum 0
	"CDA_TEXT":              cpb.PropertyRepresentationCode_CDA_TEXT,              // Enum 4
	"TYPE_ATTR":             cpb.PropertyRepresentationCode_TYPE_ATTR,             // Enum 3
	"XHTML":                 cpb.PropertyRepresentationCode_XHTML,                 // Enum 5
	"XML_ATTR":              cpb.PropertyRepresentationCode_XML_ATTR,              // Enum 1
	"XML_TEXT":              cpb.PropertyRepresentationCode_XML_TEXT,              // Enum 2
}

// DefaultPropertyTypeCodeMap maps from string to cpb.PropertyTypeCode_Value.
var DefaultPropertyTypeCodeMap = map[string]cpb.PropertyTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.PropertyTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"BOOLEAN":               cpb.PropertyTypeCode_BOOLEAN,               // Enum 5
	"CODE":                  cpb.PropertyTypeCode_CODE,                  // Enum 1
	"CODING":                cpb.PropertyTypeCode_CODING,                // Enum 2
	"DATE_TIME":             cpb.PropertyTypeCode_DATE_TIME,             // Enum 6
	"DECIMAL":               cpb.PropertyTypeCode_DECIMAL,               // Enum 7
	"INTEGER":               cpb.PropertyTypeCode_INTEGER,               // Enum 4
	"STRING":                cpb.PropertyTypeCode_STRING,                // Enum 3
}

// DefaultProvenanceEntityRoleCodeMap maps from string to cpb.ProvenanceEntityRoleCode_Value.
var DefaultProvenanceEntityRoleCodeMap = map[string]cpb.ProvenanceEntityRoleCode_Value{
	"INVALID_UNINITIALIZED": cpb.ProvenanceEntityRoleCode_INVALID_UNINITIALIZED, // Enum 0
	"DERIVATION":            cpb.ProvenanceEntityRoleCode_DERIVATION,            // Enum 1
	"QUOTATION":             cpb.ProvenanceEntityRoleCode_QUOTATION,             // Enum 3
	"REMOVAL":               cpb.ProvenanceEntityRoleCode_REMOVAL,               // Enum 5
	"REVISION":              cpb.ProvenanceEntityRoleCode_REVISION,              // Enum 2
	"SOURCE":                cpb.ProvenanceEntityRoleCode_SOURCE,                // Enum 4
}

// DefaultPublicationStatusCodeMap maps from string to cpb.PublicationStatusCode_Value.
var DefaultPublicationStatusCodeMap = map[string]cpb.PublicationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.PublicationStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.PublicationStatusCode_ACTIVE,                // Enum 2
	"DRAFT":                 cpb.PublicationStatusCode_DRAFT,                 // Enum 1
	"RETIRED":               cpb.PublicationStatusCode_RETIRED,               // Enum 3
	"UNKNOWN":               cpb.PublicationStatusCode_UNKNOWN,               // Enum 4
}

// DefaultQualityTypeCodeMap maps from string to cpb.QualityTypeCode_Value.
var DefaultQualityTypeCodeMap = map[string]cpb.QualityTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.QualityTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"INDEL":                 cpb.QualityTypeCode_INDEL,                 // Enum 1
	"SNP":                   cpb.QualityTypeCode_SNP,                   // Enum 2
	"UNKNOWN":               cpb.QualityTypeCode_UNKNOWN,               // Enum 3
}

// DefaultQuantityComparatorCodeMap maps from string to cpb.QuantityComparatorCode_Value.
var DefaultQuantityComparatorCodeMap = map[string]cpb.QuantityComparatorCode_Value{
	"INVALID_UNINITIALIZED":    cpb.QuantityComparatorCode_INVALID_UNINITIALIZED,    // Enum 0
	"<":                        cpb.QuantityComparatorCode_LESS_THAN,                // Enum 1
	"<=":                       cpb.QuantityComparatorCode_LESS_THAN_OR_EQUAL_TO,    // Enum 2
	">":                        cpb.QuantityComparatorCode_GREATER_THAN,             // Enum 4
	">=":                       cpb.QuantityComparatorCode_GREATER_THAN_OR_EQUAL_TO, // Enum 3
	"GREATER_THAN":             cpb.QuantityComparatorCode_GREATER_THAN,             // Enum 4
	"GREATER_THAN_OR_EQUAL_TO": cpb.QuantityComparatorCode_GREATER_THAN_OR_EQUAL_TO, // Enum 3
	"LESS_THAN":                cpb.QuantityComparatorCode_LESS_THAN,                // Enum 1
	"LESS_THAN_OR_EQUAL_TO":    cpb.QuantityComparatorCode_LESS_THAN_OR_EQUAL_TO,    // Enum 2
}

// DefaultQuestionnaireItemOperatorCodeMap maps from string to cpb.QuestionnaireItemOperatorCode_Value.
var DefaultQuestionnaireItemOperatorCodeMap = map[string]cpb.QuestionnaireItemOperatorCode_Value{
	"INVALID_UNINITIALIZED":    cpb.QuestionnaireItemOperatorCode_INVALID_UNINITIALIZED,    // Enum 0
	"!=":                       cpb.QuestionnaireItemOperatorCode_NOT_EQUAL_TO,             // Enum 3
	"<":                        cpb.QuestionnaireItemOperatorCode_LESS_THAN,                // Enum 5
	"<=":                       cpb.QuestionnaireItemOperatorCode_LESS_THAN_OR_EQUAL_TO,    // Enum 7
	"=":                        cpb.QuestionnaireItemOperatorCode_EQUALS,                   // Enum 2
	">":                        cpb.QuestionnaireItemOperatorCode_GREATER_THAN,             // Enum 4
	">=":                       cpb.QuestionnaireItemOperatorCode_GREATER_THAN_OR_EQUAL_TO, // Enum 6
	"EQUALS":                   cpb.QuestionnaireItemOperatorCode_EQUALS,                   // Enum 2
	"EXISTS":                   cpb.QuestionnaireItemOperatorCode_EXISTS,                   // Enum 1
	"GREATER_THAN":             cpb.QuestionnaireItemOperatorCode_GREATER_THAN,             // Enum 4
	"GREATER_THAN_OR_EQUAL_TO": cpb.QuestionnaireItemOperatorCode_GREATER_THAN_OR_EQUAL_TO, // Enum 6
	"LESS_THAN":                cpb.QuestionnaireItemOperatorCode_LESS_THAN,                // Enum 5
	"LESS_THAN_OR_EQUAL_TO":    cpb.QuestionnaireItemOperatorCode_LESS_THAN_OR_EQUAL_TO,    // Enum 7
	"NOT_EQUAL_TO":             cpb.QuestionnaireItemOperatorCode_NOT_EQUAL_TO,             // Enum 3
}

// DefaultQuestionnaireItemTypeCodeMap maps from string to cpb.QuestionnaireItemTypeCode_Value.
var DefaultQuestionnaireItemTypeCodeMap = map[string]cpb.QuestionnaireItemTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.QuestionnaireItemTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"ATTACHMENT":            cpb.QuestionnaireItemTypeCode_ATTACHMENT,            // Enum 15
	"BOOLEAN":               cpb.QuestionnaireItemTypeCode_BOOLEAN,               // Enum 4
	"CHOICE":                cpb.QuestionnaireItemTypeCode_CHOICE,                // Enum 13
	"DATE":                  cpb.QuestionnaireItemTypeCode_DATE,                  // Enum 7
	"DATE_TIME":             cpb.QuestionnaireItemTypeCode_DATE_TIME,             // Enum 8
	"DECIMAL":               cpb.QuestionnaireItemTypeCode_DECIMAL,               // Enum 5
	"DISPLAY":               cpb.QuestionnaireItemTypeCode_DISPLAY,               // Enum 2
	"GROUP":                 cpb.QuestionnaireItemTypeCode_GROUP,                 // Enum 1
	"INTEGER":               cpb.QuestionnaireItemTypeCode_INTEGER,               // Enum 6
	"OPEN_CHOICE":           cpb.QuestionnaireItemTypeCode_OPEN_CHOICE,           // Enum 14
	"QUANTITY":              cpb.QuestionnaireItemTypeCode_QUANTITY,              // Enum 17
	"QUESTION":              cpb.QuestionnaireItemTypeCode_QUESTION,              // Enum 3
	"REFERENCE":             cpb.QuestionnaireItemTypeCode_REFERENCE,             // Enum 16
	"STRING":                cpb.QuestionnaireItemTypeCode_STRING,                // Enum 10
	"TEXT":                  cpb.QuestionnaireItemTypeCode_TEXT,                  // Enum 11
	"TIME":                  cpb.QuestionnaireItemTypeCode_TIME,                  // Enum 9
	"URL":                   cpb.QuestionnaireItemTypeCode_URL,                   // Enum 12
}

// DefaultQuestionnaireItemUsageModeCodeMap maps from string to cpb.QuestionnaireItemUsageModeCode_Value.
var DefaultQuestionnaireItemUsageModeCodeMap = map[string]cpb.QuestionnaireItemUsageModeCode_Value{
	"INVALID_UNINITIALIZED":     cpb.QuestionnaireItemUsageModeCode_INVALID_UNINITIALIZED,     // Enum 0
	"CAPTURE":                   cpb.QuestionnaireItemUsageModeCode_CAPTURE,                   // Enum 2
	"CAPTURE_DISPLAY":           cpb.QuestionnaireItemUsageModeCode_CAPTURE_DISPLAY,           // Enum 1
	"CAPTURE_DISPLAY_NON_EMPTY": cpb.QuestionnaireItemUsageModeCode_CAPTURE_DISPLAY_NON_EMPTY, // Enum 5
	"DISPLAY":                   cpb.QuestionnaireItemUsageModeCode_DISPLAY,                   // Enum 3
	"DISPLAY_NON_EMPTY":         cpb.QuestionnaireItemUsageModeCode_DISPLAY_NON_EMPTY,         // Enum 4
}

// DefaultQuestionnaireResponseStatusCodeMap maps from string to cpb.QuestionnaireResponseStatusCode_Value.
var DefaultQuestionnaireResponseStatusCodeMap = map[string]cpb.QuestionnaireResponseStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.QuestionnaireResponseStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"AMENDED":               cpb.QuestionnaireResponseStatusCode_AMENDED,               // Enum 3
	"COMPLETED":             cpb.QuestionnaireResponseStatusCode_COMPLETED,             // Enum 2
	"ENTERED_IN_ERROR":      cpb.QuestionnaireResponseStatusCode_ENTERED_IN_ERROR,      // Enum 4
	"IN_PROGRESS":           cpb.QuestionnaireResponseStatusCode_IN_PROGRESS,           // Enum 1
	"STOPPED":               cpb.QuestionnaireResponseStatusCode_STOPPED,               // Enum 5
}

// DefaultReferenceHandlingPolicyCodeMap maps from string to cpb.ReferenceHandlingPolicyCode_Value.
var DefaultReferenceHandlingPolicyCodeMap = map[string]cpb.ReferenceHandlingPolicyCode_Value{
	"INVALID_UNINITIALIZED": cpb.ReferenceHandlingPolicyCode_INVALID_UNINITIALIZED, // Enum 0
	"ENFORCED":              cpb.ReferenceHandlingPolicyCode_ENFORCED,              // Enum 4
	"LITERAL":               cpb.ReferenceHandlingPolicyCode_LITERAL,               // Enum 1
	"LOCAL":                 cpb.ReferenceHandlingPolicyCode_LOCAL,                 // Enum 5
	"LOGICAL":               cpb.ReferenceHandlingPolicyCode_LOGICAL,               // Enum 2
	"RESOLVES":              cpb.ReferenceHandlingPolicyCode_RESOLVES,              // Enum 3
}

// DefaultReferenceVersionRulesCodeMap maps from string to cpb.ReferenceVersionRulesCode_Value.
var DefaultReferenceVersionRulesCodeMap = map[string]cpb.ReferenceVersionRulesCode_Value{
	"INVALID_UNINITIALIZED": cpb.ReferenceVersionRulesCode_INVALID_UNINITIALIZED, // Enum 0
	"EITHER":                cpb.ReferenceVersionRulesCode_EITHER,                // Enum 1
	"INDEPENDENT":           cpb.ReferenceVersionRulesCode_INDEPENDENT,           // Enum 2
	"SPECIFIC":              cpb.ReferenceVersionRulesCode_SPECIFIC,              // Enum 3
}

// DefaultRelatedArtifactTypeCodeMap maps from string to cpb.RelatedArtifactTypeCode_Value.
var DefaultRelatedArtifactTypeCodeMap = map[string]cpb.RelatedArtifactTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.RelatedArtifactTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"CITATION":              cpb.RelatedArtifactTypeCode_CITATION,              // Enum 3
	"COMPOSED_OF":           cpb.RelatedArtifactTypeCode_COMPOSED_OF,           // Enum 8
	"DEPENDS_ON":            cpb.RelatedArtifactTypeCode_DEPENDS_ON,            // Enum 7
	"DERIVED_FROM":          cpb.RelatedArtifactTypeCode_DERIVED_FROM,          // Enum 6
	"DOCUMENTATION":         cpb.RelatedArtifactTypeCode_DOCUMENTATION,         // Enum 1
	"JUSTIFICATION":         cpb.RelatedArtifactTypeCode_JUSTIFICATION,         // Enum 2
	"PREDECESSOR":           cpb.RelatedArtifactTypeCode_PREDECESSOR,           // Enum 4
	"SUCCESSOR":             cpb.RelatedArtifactTypeCode_SUCCESSOR,             // Enum 5
}

// DefaultRepositoryTypeCodeMap maps from string to cpb.RepositoryTypeCode_Value.
var DefaultRepositoryTypeCodeMap = map[string]cpb.RepositoryTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.RepositoryTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"DIRECTLINK":            cpb.RepositoryTypeCode_DIRECTLINK,            // Enum 1
	"LOGIN":                 cpb.RepositoryTypeCode_LOGIN,                 // Enum 3
	"OAUTH":                 cpb.RepositoryTypeCode_OAUTH,                 // Enum 4
	"OPENAPI":               cpb.RepositoryTypeCode_OPENAPI,               // Enum 2
	"OTHER":                 cpb.RepositoryTypeCode_OTHER,                 // Enum 5
}

// DefaultRequestIntentCodeMap maps from string to cpb.RequestIntentCode_Value.
var DefaultRequestIntentCodeMap = map[string]cpb.RequestIntentCode_Value{
	"INVALID_UNINITIALIZED": cpb.RequestIntentCode_INVALID_UNINITIALIZED, // Enum 0
	"DIRECTIVE":             cpb.RequestIntentCode_DIRECTIVE,             // Enum 3
	"FILLER_ORDER":          cpb.RequestIntentCode_FILLER_ORDER,          // Enum 7
	"INSTANCE_ORDER":        cpb.RequestIntentCode_INSTANCE_ORDER,        // Enum 8
	"OPTION":                cpb.RequestIntentCode_OPTION,                // Enum 9
	"ORDER":                 cpb.RequestIntentCode_ORDER,                 // Enum 4
	"ORIGINAL_ORDER":        cpb.RequestIntentCode_ORIGINAL_ORDER,        // Enum 5
	"PLAN":                  cpb.RequestIntentCode_PLAN,                  // Enum 2
	"PROPOSAL":              cpb.RequestIntentCode_PROPOSAL,              // Enum 1
	"REFLEX_ORDER":          cpb.RequestIntentCode_REFLEX_ORDER,          // Enum 6
}

// DefaultRequestPriorityCodeMap maps from string to cpb.RequestPriorityCode_Value.
var DefaultRequestPriorityCodeMap = map[string]cpb.RequestPriorityCode_Value{
	"INVALID_UNINITIALIZED": cpb.RequestPriorityCode_INVALID_UNINITIALIZED, // Enum 0
	"ASAP":                  cpb.RequestPriorityCode_ASAP,                  // Enum 3
	"ROUTINE":               cpb.RequestPriorityCode_ROUTINE,               // Enum 1
	"STAT":                  cpb.RequestPriorityCode_STAT,                  // Enum 4
	"URGENT":                cpb.RequestPriorityCode_URGENT,                // Enum 2
}

// DefaultRequestResourceTypeCodeMap maps from string to cpb.RequestResourceTypeCode_Value.
var DefaultRequestResourceTypeCodeMap = map[string]cpb.RequestResourceTypeCode_Value{
	"INVALID_UNINITIALIZED":       cpb.RequestResourceTypeCode_INVALID_UNINITIALIZED,       // Enum 0
	"APPOINTMENT":                 cpb.RequestResourceTypeCode_APPOINTMENT,                 // Enum 1
	"APPOINTMENT_RESPONSE":        cpb.RequestResourceTypeCode_APPOINTMENT_RESPONSE,        // Enum 2
	"CARE_PLAN":                   cpb.RequestResourceTypeCode_CARE_PLAN,                   // Enum 3
	"CLAIM":                       cpb.RequestResourceTypeCode_CLAIM,                       // Enum 4
	"COMMUNICATION_REQUEST":       cpb.RequestResourceTypeCode_COMMUNICATION_REQUEST,       // Enum 5
	"CONTRACT":                    cpb.RequestResourceTypeCode_CONTRACT,                    // Enum 6
	"DEVICE_REQUEST":              cpb.RequestResourceTypeCode_DEVICE_REQUEST,              // Enum 7
	"ENROLLMENT_REQUEST":          cpb.RequestResourceTypeCode_ENROLLMENT_REQUEST,          // Enum 8
	"IMMUNIZATION_RECOMMENDATION": cpb.RequestResourceTypeCode_IMMUNIZATION_RECOMMENDATION, // Enum 9
	"MEDICATION_REQUEST":          cpb.RequestResourceTypeCode_MEDICATION_REQUEST,          // Enum 10
	"NUTRITION_ORDER":             cpb.RequestResourceTypeCode_NUTRITION_ORDER,             // Enum 11
	"SERVICE_REQUEST":             cpb.RequestResourceTypeCode_SERVICE_REQUEST,             // Enum 12
	"SUPPLY_REQUEST":              cpb.RequestResourceTypeCode_SUPPLY_REQUEST,              // Enum 13
	"TASK":                        cpb.RequestResourceTypeCode_TASK,                        // Enum 14
	"VISION_PRESCRIPTION":         cpb.RequestResourceTypeCode_VISION_PRESCRIPTION,         // Enum 15
}

// DefaultRequestStatusCodeMap maps from string to cpb.RequestStatusCode_Value.
var DefaultRequestStatusCodeMap = map[string]cpb.RequestStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.RequestStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.RequestStatusCode_ACTIVE,                // Enum 2
	"COMPLETED":             cpb.RequestStatusCode_COMPLETED,             // Enum 5
	"DRAFT":                 cpb.RequestStatusCode_DRAFT,                 // Enum 1
	"ENTERED_IN_ERROR":      cpb.RequestStatusCode_ENTERED_IN_ERROR,      // Enum 6
	"ON_HOLD":               cpb.RequestStatusCode_ON_HOLD,               // Enum 3
	"REVOKED":               cpb.RequestStatusCode_REVOKED,               // Enum 4
	"UNKNOWN":               cpb.RequestStatusCode_UNKNOWN,               // Enum 7
}

// DefaultResearchElementTypeCodeMap maps from string to cpb.ResearchElementTypeCode_Value.
var DefaultResearchElementTypeCodeMap = map[string]cpb.ResearchElementTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ResearchElementTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"EXPOSURE":              cpb.ResearchElementTypeCode_EXPOSURE,              // Enum 2
	"OUTCOME":               cpb.ResearchElementTypeCode_OUTCOME,               // Enum 3
	"POPULATION":            cpb.ResearchElementTypeCode_POPULATION,            // Enum 1
}

// DefaultResearchStudyStatusCodeMap maps from string to cpb.ResearchStudyStatusCode_Value.
var DefaultResearchStudyStatusCodeMap = map[string]cpb.ResearchStudyStatusCode_Value{
	"INVALID_UNINITIALIZED":              cpb.ResearchStudyStatusCode_INVALID_UNINITIALIZED,              // Enum 0
	"ACTIVE":                             cpb.ResearchStudyStatusCode_ACTIVE,                             // Enum 1
	"ADMINISTRATIVELY_COMPLETED":         cpb.ResearchStudyStatusCode_ADMINISTRATIVELY_COMPLETED,         // Enum 2
	"APPROVED":                           cpb.ResearchStudyStatusCode_APPROVED,                           // Enum 3
	"CLOSED_TO_ACCRUAL":                  cpb.ResearchStudyStatusCode_CLOSED_TO_ACCRUAL,                  // Enum 4
	"CLOSED_TO_ACCRUAL_AND_INTERVENTION": cpb.ResearchStudyStatusCode_CLOSED_TO_ACCRUAL_AND_INTERVENTION, // Enum 5
	"COMPLETED":                          cpb.ResearchStudyStatusCode_COMPLETED,                          // Enum 6
	"DISAPPROVED":                        cpb.ResearchStudyStatusCode_DISAPPROVED,                        // Enum 7
	"IN_REVIEW":                          cpb.ResearchStudyStatusCode_IN_REVIEW,                          // Enum 8
	"TEMPORARILY_CLOSED_TO_ACCRUAL":      cpb.ResearchStudyStatusCode_TEMPORARILY_CLOSED_TO_ACCRUAL,      // Enum 9
	"TEMPORARILY_CLOSED_TO_ACCRUAL_AND_INTERVENTION": cpb.ResearchStudyStatusCode_TEMPORARILY_CLOSED_TO_ACCRUAL_AND_INTERVENTION, // Enum 10
	"WITHDRAWN": cpb.ResearchStudyStatusCode_WITHDRAWN, // Enum 11
}

// DefaultResearchSubjectStatusCodeMap maps from string to cpb.ResearchSubjectStatusCode_Value.
var DefaultResearchSubjectStatusCodeMap = map[string]cpb.ResearchSubjectStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ResearchSubjectStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"CANDIDATE":             cpb.ResearchSubjectStatusCode_CANDIDATE,             // Enum 1
	"ELIGIBLE":              cpb.ResearchSubjectStatusCode_ELIGIBLE,              // Enum 2
	"FOLLOW_UP":             cpb.ResearchSubjectStatusCode_FOLLOW_UP,             // Enum 3
	"INELIGIBLE":            cpb.ResearchSubjectStatusCode_INELIGIBLE,            // Enum 4
	"NOT_REGISTERED":        cpb.ResearchSubjectStatusCode_NOT_REGISTERED,        // Enum 5
	"OFF_STUDY":             cpb.ResearchSubjectStatusCode_OFF_STUDY,             // Enum 6
	"ON_STUDY":              cpb.ResearchSubjectStatusCode_ON_STUDY,              // Enum 7
	"ON_STUDY_INTERVENTION": cpb.ResearchSubjectStatusCode_ON_STUDY_INTERVENTION, // Enum 8
	"ON_STUDY_OBSERVATION":  cpb.ResearchSubjectStatusCode_ON_STUDY_OBSERVATION,  // Enum 9
	"PENDING_ON_STUDY":      cpb.ResearchSubjectStatusCode_PENDING_ON_STUDY,      // Enum 10
	"POTENTIAL_CANDIDATE":   cpb.ResearchSubjectStatusCode_POTENTIAL_CANDIDATE,   // Enum 11
	"SCREENING":             cpb.ResearchSubjectStatusCode_SCREENING,             // Enum 12
	"WITHDRAWN":             cpb.ResearchSubjectStatusCode_WITHDRAWN,             // Enum 13
}

// DefaultResourceSecurityCategoryCodeMap maps from string to cpb.ResourceSecurityCategoryCode_Value.
var DefaultResourceSecurityCategoryCodeMap = map[string]cpb.ResourceSecurityCategoryCode_Value{
	"INVALID_UNINITIALIZED": cpb.ResourceSecurityCategoryCode_INVALID_UNINITIALIZED, // Enum 0
	"ANONYMOUS":             cpb.ResourceSecurityCategoryCode_ANONYMOUS,             // Enum 1
	"BUSINESS":              cpb.ResourceSecurityCategoryCode_BUSINESS,              // Enum 2
	"INDIVIDUAL":            cpb.ResourceSecurityCategoryCode_INDIVIDUAL,            // Enum 3
	"NOT_CLASSIFIED":        cpb.ResourceSecurityCategoryCode_NOT_CLASSIFIED,        // Enum 5
	"PATIENT":               cpb.ResourceSecurityCategoryCode_PATIENT,               // Enum 4
}

// DefaultResourceTypeCodeMap maps from string to cpb.ResourceTypeCode_Value.
var DefaultResourceTypeCodeMap = map[string]cpb.ResourceTypeCode_Value{
	"INVALID_UNINITIALIZED":                cpb.ResourceTypeCode_INVALID_UNINITIALIZED,                // Enum 0
	"ACCOUNT":                              cpb.ResourceTypeCode_ACCOUNT,                              // Enum 1
	"ACTIVITY_DEFINITION":                  cpb.ResourceTypeCode_ACTIVITY_DEFINITION,                  // Enum 2
	"ADVERSE_EVENT":                        cpb.ResourceTypeCode_ADVERSE_EVENT,                        // Enum 3
	"ALLERGY_INTOLERANCE":                  cpb.ResourceTypeCode_ALLERGY_INTOLERANCE,                  // Enum 4
	"APPOINTMENT":                          cpb.ResourceTypeCode_APPOINTMENT,                          // Enum 5
	"APPOINTMENT_RESPONSE":                 cpb.ResourceTypeCode_APPOINTMENT_RESPONSE,                 // Enum 6
	"AUDIT_EVENT":                          cpb.ResourceTypeCode_AUDIT_EVENT,                          // Enum 7
	"BASIC":                                cpb.ResourceTypeCode_BASIC,                                // Enum 8
	"BINARY":                               cpb.ResourceTypeCode_BINARY,                               // Enum 9
	"BIOLOGICALLY_DERIVED_PRODUCT":         cpb.ResourceTypeCode_BIOLOGICALLY_DERIVED_PRODUCT,         // Enum 10
	"BODY_STRUCTURE":                       cpb.ResourceTypeCode_BODY_STRUCTURE,                       // Enum 11
	"BUNDLE":                               cpb.ResourceTypeCode_BUNDLE,                               // Enum 12
	"CAPABILITY_STATEMENT":                 cpb.ResourceTypeCode_CAPABILITY_STATEMENT,                 // Enum 13
	"CARE_PLAN":                            cpb.ResourceTypeCode_CARE_PLAN,                            // Enum 14
	"CARE_TEAM":                            cpb.ResourceTypeCode_CARE_TEAM,                            // Enum 15
	"CATALOG_ENTRY":                        cpb.ResourceTypeCode_CATALOG_ENTRY,                        // Enum 16
	"CHARGE_ITEM":                          cpb.ResourceTypeCode_CHARGE_ITEM,                          // Enum 17
	"CHARGE_ITEM_DEFINITION":               cpb.ResourceTypeCode_CHARGE_ITEM_DEFINITION,               // Enum 18
	"CLAIM":                                cpb.ResourceTypeCode_CLAIM,                                // Enum 19
	"CLAIM_RESPONSE":                       cpb.ResourceTypeCode_CLAIM_RESPONSE,                       // Enum 20
	"CLINICAL_IMPRESSION":                  cpb.ResourceTypeCode_CLINICAL_IMPRESSION,                  // Enum 21
	"CODE_SYSTEM":                          cpb.ResourceTypeCode_CODE_SYSTEM,                          // Enum 22
	"COMMUNICATION":                        cpb.ResourceTypeCode_COMMUNICATION,                        // Enum 23
	"COMMUNICATION_REQUEST":                cpb.ResourceTypeCode_COMMUNICATION_REQUEST,                // Enum 24
	"COMPARTMENT_DEFINITION":               cpb.ResourceTypeCode_COMPARTMENT_DEFINITION,               // Enum 25
	"COMPOSITION":                          cpb.ResourceTypeCode_COMPOSITION,                          // Enum 26
	"CONCEPT_MAP":                          cpb.ResourceTypeCode_CONCEPT_MAP,                          // Enum 27
	"CONDITION":                            cpb.ResourceTypeCode_CONDITION,                            // Enum 28
	"CONSENT":                              cpb.ResourceTypeCode_CONSENT,                              // Enum 29
	"CONTRACT":                             cpb.ResourceTypeCode_CONTRACT,                             // Enum 30
	"COVERAGE":                             cpb.ResourceTypeCode_COVERAGE,                             // Enum 31
	"COVERAGE_ELIGIBILITY_REQUEST":         cpb.ResourceTypeCode_COVERAGE_ELIGIBILITY_REQUEST,         // Enum 32
	"COVERAGE_ELIGIBILITY_RESPONSE":        cpb.ResourceTypeCode_COVERAGE_ELIGIBILITY_RESPONSE,        // Enum 33
	"DETECTED_ISSUE":                       cpb.ResourceTypeCode_DETECTED_ISSUE,                       // Enum 34
	"DEVICE":                               cpb.ResourceTypeCode_DEVICE,                               // Enum 35
	"DEVICE_DEFINITION":                    cpb.ResourceTypeCode_DEVICE_DEFINITION,                    // Enum 36
	"DEVICE_METRIC":                        cpb.ResourceTypeCode_DEVICE_METRIC,                        // Enum 37
	"DEVICE_REQUEST":                       cpb.ResourceTypeCode_DEVICE_REQUEST,                       // Enum 38
	"DEVICE_USE_STATEMENT":                 cpb.ResourceTypeCode_DEVICE_USE_STATEMENT,                 // Enum 39
	"DIAGNOSTIC_REPORT":                    cpb.ResourceTypeCode_DIAGNOSTIC_REPORT,                    // Enum 40
	"DOCUMENT_MANIFEST":                    cpb.ResourceTypeCode_DOCUMENT_MANIFEST,                    // Enum 41
	"DOCUMENT_REFERENCE":                   cpb.ResourceTypeCode_DOCUMENT_REFERENCE,                   // Enum 42
	"DOMAIN_RESOURCE":                      cpb.ResourceTypeCode_DOMAIN_RESOURCE,                      // Enum 43
	"EFFECT_EVIDENCE_SYNTHESIS":            cpb.ResourceTypeCode_EFFECT_EVIDENCE_SYNTHESIS,            // Enum 44
	"ENCOUNTER":                            cpb.ResourceTypeCode_ENCOUNTER,                            // Enum 45
	"ENDPOINT":                             cpb.ResourceTypeCode_ENDPOINT,                             // Enum 46
	"ENROLLMENT_REQUEST":                   cpb.ResourceTypeCode_ENROLLMENT_REQUEST,                   // Enum 47
	"ENROLLMENT_RESPONSE":                  cpb.ResourceTypeCode_ENROLLMENT_RESPONSE,                  // Enum 48
	"EPISODE_OF_CARE":                      cpb.ResourceTypeCode_EPISODE_OF_CARE,                      // Enum 49
	"EVENT_DEFINITION":                     cpb.ResourceTypeCode_EVENT_DEFINITION,                     // Enum 50
	"EVIDENCE":                             cpb.ResourceTypeCode_EVIDENCE,                             // Enum 51
	"EVIDENCE_VARIABLE":                    cpb.ResourceTypeCode_EVIDENCE_VARIABLE,                    // Enum 52
	"EXAMPLE_SCENARIO":                     cpb.ResourceTypeCode_EXAMPLE_SCENARIO,                     // Enum 53
	"EXPLANATION_OF_BENEFIT":               cpb.ResourceTypeCode_EXPLANATION_OF_BENEFIT,               // Enum 54
	"FAMILY_MEMBER_HISTORY":                cpb.ResourceTypeCode_FAMILY_MEMBER_HISTORY,                // Enum 55
	"FLAG":                                 cpb.ResourceTypeCode_FLAG,                                 // Enum 56
	"GOAL":                                 cpb.ResourceTypeCode_GOAL,                                 // Enum 57
	"GRAPH_DEFINITION":                     cpb.ResourceTypeCode_GRAPH_DEFINITION,                     // Enum 58
	"GROUP":                                cpb.ResourceTypeCode_GROUP,                                // Enum 59
	"GUIDANCE_RESPONSE":                    cpb.ResourceTypeCode_GUIDANCE_RESPONSE,                    // Enum 60
	"HEALTHCARE_SERVICE":                   cpb.ResourceTypeCode_HEALTHCARE_SERVICE,                   // Enum 61
	"IMAGING_STUDY":                        cpb.ResourceTypeCode_IMAGING_STUDY,                        // Enum 62
	"IMMUNIZATION":                         cpb.ResourceTypeCode_IMMUNIZATION,                         // Enum 63
	"IMMUNIZATION_EVALUATION":              cpb.ResourceTypeCode_IMMUNIZATION_EVALUATION,              // Enum 64
	"IMMUNIZATION_RECOMMENDATION":          cpb.ResourceTypeCode_IMMUNIZATION_RECOMMENDATION,          // Enum 65
	"IMPLEMENTATION_GUIDE":                 cpb.ResourceTypeCode_IMPLEMENTATION_GUIDE,                 // Enum 66
	"INSURANCE_PLAN":                       cpb.ResourceTypeCode_INSURANCE_PLAN,                       // Enum 67
	"INVOICE":                              cpb.ResourceTypeCode_INVOICE,                              // Enum 68
	"LIBRARY":                              cpb.ResourceTypeCode_LIBRARY,                              // Enum 69
	"LINKAGE":                              cpb.ResourceTypeCode_LINKAGE,                              // Enum 70
	"LIST":                                 cpb.ResourceTypeCode_LIST,                                 // Enum 71
	"LOCATION":                             cpb.ResourceTypeCode_LOCATION,                             // Enum 72
	"MEASURE":                              cpb.ResourceTypeCode_MEASURE,                              // Enum 73
	"MEASURE_REPORT":                       cpb.ResourceTypeCode_MEASURE_REPORT,                       // Enum 74
	"MEDIA":                                cpb.ResourceTypeCode_MEDIA,                                // Enum 75
	"MEDICATION":                           cpb.ResourceTypeCode_MEDICATION,                           // Enum 76
	"MEDICATION_ADMINISTRATION":            cpb.ResourceTypeCode_MEDICATION_ADMINISTRATION,            // Enum 77
	"MEDICATION_DISPENSE":                  cpb.ResourceTypeCode_MEDICATION_DISPENSE,                  // Enum 78
	"MEDICATION_KNOWLEDGE":                 cpb.ResourceTypeCode_MEDICATION_KNOWLEDGE,                 // Enum 79
	"MEDICATION_REQUEST":                   cpb.ResourceTypeCode_MEDICATION_REQUEST,                   // Enum 80
	"MEDICATION_STATEMENT":                 cpb.ResourceTypeCode_MEDICATION_STATEMENT,                 // Enum 81
	"MEDICINAL_PRODUCT":                    cpb.ResourceTypeCode_MEDICINAL_PRODUCT,                    // Enum 82
	"MEDICINAL_PRODUCT_AUTHORIZATION":      cpb.ResourceTypeCode_MEDICINAL_PRODUCT_AUTHORIZATION,      // Enum 83
	"MEDICINAL_PRODUCT_CONTRAINDICATION":   cpb.ResourceTypeCode_MEDICINAL_PRODUCT_CONTRAINDICATION,   // Enum 84
	"MEDICINAL_PRODUCT_INDICATION":         cpb.ResourceTypeCode_MEDICINAL_PRODUCT_INDICATION,         // Enum 85
	"MEDICINAL_PRODUCT_INGREDIENT":         cpb.ResourceTypeCode_MEDICINAL_PRODUCT_INGREDIENT,         // Enum 86
	"MEDICINAL_PRODUCT_INTERACTION":        cpb.ResourceTypeCode_MEDICINAL_PRODUCT_INTERACTION,        // Enum 87
	"MEDICINAL_PRODUCT_MANUFACTURED":       cpb.ResourceTypeCode_MEDICINAL_PRODUCT_MANUFACTURED,       // Enum 88
	"MEDICINAL_PRODUCT_PACKAGED":           cpb.ResourceTypeCode_MEDICINAL_PRODUCT_PACKAGED,           // Enum 89
	"MEDICINAL_PRODUCT_PHARMACEUTICAL":     cpb.ResourceTypeCode_MEDICINAL_PRODUCT_PHARMACEUTICAL,     // Enum 90
	"MEDICINAL_PRODUCT_UNDESIRABLE_EFFECT": cpb.ResourceTypeCode_MEDICINAL_PRODUCT_UNDESIRABLE_EFFECT, // Enum 91
	"MESSAGE_DEFINITION":                   cpb.ResourceTypeCode_MESSAGE_DEFINITION,                   // Enum 92
	"MESSAGE_HEADER":                       cpb.ResourceTypeCode_MESSAGE_HEADER,                       // Enum 93
	"MOLECULAR_SEQUENCE":                   cpb.ResourceTypeCode_MOLECULAR_SEQUENCE,                   // Enum 94
	"NAMING_SYSTEM":                        cpb.ResourceTypeCode_NAMING_SYSTEM,                        // Enum 95
	"NUTRITION_ORDER":                      cpb.ResourceTypeCode_NUTRITION_ORDER,                      // Enum 96
	"OBSERVATION":                          cpb.ResourceTypeCode_OBSERVATION,                          // Enum 97
	"OBSERVATION_DEFINITION":               cpb.ResourceTypeCode_OBSERVATION_DEFINITION,               // Enum 98
	"OPERATION_DEFINITION":                 cpb.ResourceTypeCode_OPERATION_DEFINITION,                 // Enum 99
	"OPERATION_OUTCOME":                    cpb.ResourceTypeCode_OPERATION_OUTCOME,                    // Enum 100
	"ORGANIZATION":                         cpb.ResourceTypeCode_ORGANIZATION,                         // Enum 101
	"ORGANIZATION_AFFILIATION":             cpb.ResourceTypeCode_ORGANIZATION_AFFILIATION,             // Enum 102
	"PARAMETERS":                           cpb.ResourceTypeCode_PARAMETERS,                           // Enum 103
	"PATIENT":                              cpb.ResourceTypeCode_PATIENT,                              // Enum 104
	"PAYMENT_NOTICE":                       cpb.ResourceTypeCode_PAYMENT_NOTICE,                       // Enum 105
	"PAYMENT_RECONCILIATION":               cpb.ResourceTypeCode_PAYMENT_RECONCILIATION,               // Enum 106
	"PERSON":                               cpb.ResourceTypeCode_PERSON,                               // Enum 107
	"PLAN_DEFINITION":                      cpb.ResourceTypeCode_PLAN_DEFINITION,                      // Enum 108
	"PRACTITIONER":                         cpb.ResourceTypeCode_PRACTITIONER,                         // Enum 109
	"PRACTITIONER_ROLE":                    cpb.ResourceTypeCode_PRACTITIONER_ROLE,                    // Enum 110
	"PROCEDURE":                            cpb.ResourceTypeCode_PROCEDURE,                            // Enum 111
	"PROVENANCE":                           cpb.ResourceTypeCode_PROVENANCE,                           // Enum 112
	"QUESTIONNAIRE":                        cpb.ResourceTypeCode_QUESTIONNAIRE,                        // Enum 113
	"QUESTIONNAIRE_RESPONSE":               cpb.ResourceTypeCode_QUESTIONNAIRE_RESPONSE,               // Enum 114
	"RELATED_PERSON":                       cpb.ResourceTypeCode_RELATED_PERSON,                       // Enum 115
	"REQUEST_GROUP":                        cpb.ResourceTypeCode_REQUEST_GROUP,                        // Enum 116
	"RESEARCH_DEFINITION":                  cpb.ResourceTypeCode_RESEARCH_DEFINITION,                  // Enum 117
	"RESEARCH_ELEMENT_DEFINITION":          cpb.ResourceTypeCode_RESEARCH_ELEMENT_DEFINITION,          // Enum 118
	"RESEARCH_STUDY":                       cpb.ResourceTypeCode_RESEARCH_STUDY,                       // Enum 119
	"RESEARCH_SUBJECT":                     cpb.ResourceTypeCode_RESEARCH_SUBJECT,                     // Enum 120
	"RESOURCE":                             cpb.ResourceTypeCode_RESOURCE,                             // Enum 121
	"RISK_ASSESSMENT":                      cpb.ResourceTypeCode_RISK_ASSESSMENT,                      // Enum 122
	"RISK_EVIDENCE_SYNTHESIS":              cpb.ResourceTypeCode_RISK_EVIDENCE_SYNTHESIS,              // Enum 123
	"SCHEDULE":                             cpb.ResourceTypeCode_SCHEDULE,                             // Enum 124
	"SEARCH_PARAMETER":                     cpb.ResourceTypeCode_SEARCH_PARAMETER,                     // Enum 125
	"SERVICE_REQUEST":                      cpb.ResourceTypeCode_SERVICE_REQUEST,                      // Enum 126
	"SLOT":                                 cpb.ResourceTypeCode_SLOT,                                 // Enum 127
	"SPECIMEN":                             cpb.ResourceTypeCode_SPECIMEN,                             // Enum 128
	"SPECIMEN_DEFINITION":                  cpb.ResourceTypeCode_SPECIMEN_DEFINITION,                  // Enum 129
	"STRUCTURE_DEFINITION":                 cpb.ResourceTypeCode_STRUCTURE_DEFINITION,                 // Enum 130
	"STRUCTURE_MAP":                        cpb.ResourceTypeCode_STRUCTURE_MAP,                        // Enum 131
	"SUBSCRIPTION":                         cpb.ResourceTypeCode_SUBSCRIPTION,                         // Enum 132
	"SUBSTANCE":                            cpb.ResourceTypeCode_SUBSTANCE,                            // Enum 133
	"SUBSTANCE_NUCLEIC_ACID":               cpb.ResourceTypeCode_SUBSTANCE_NUCLEIC_ACID,               // Enum 134
	"SUBSTANCE_POLYMER":                    cpb.ResourceTypeCode_SUBSTANCE_POLYMER,                    // Enum 135
	"SUBSTANCE_PROTEIN":                    cpb.ResourceTypeCode_SUBSTANCE_PROTEIN,                    // Enum 136
	"SUBSTANCE_REFERENCE_INFORMATION":      cpb.ResourceTypeCode_SUBSTANCE_REFERENCE_INFORMATION,      // Enum 137
	"SUBSTANCE_SOURCE_MATERIAL":            cpb.ResourceTypeCode_SUBSTANCE_SOURCE_MATERIAL,            // Enum 138
	"SUBSTANCE_SPECIFICATION":              cpb.ResourceTypeCode_SUBSTANCE_SPECIFICATION,              // Enum 139
	"SUPPLY_DELIVERY":                      cpb.ResourceTypeCode_SUPPLY_DELIVERY,                      // Enum 140
	"SUPPLY_REQUEST":                       cpb.ResourceTypeCode_SUPPLY_REQUEST,                       // Enum 141
	"TASK":                                 cpb.ResourceTypeCode_TASK,                                 // Enum 142
	"TERMINOLOGY_CAPABILITIES":             cpb.ResourceTypeCode_TERMINOLOGY_CAPABILITIES,             // Enum 143
	"TEST_REPORT":                          cpb.ResourceTypeCode_TEST_REPORT,                          // Enum 144
	"TEST_SCRIPT":                          cpb.ResourceTypeCode_TEST_SCRIPT,                          // Enum 145
	"VALUE_SET":                            cpb.ResourceTypeCode_VALUE_SET,                            // Enum 146
	"VERIFICATION_RESULT":                  cpb.ResourceTypeCode_VERIFICATION_RESULT,                  // Enum 147
	"VISION_PRESCRIPTION":                  cpb.ResourceTypeCode_VISION_PRESCRIPTION,                  // Enum 148
}

// DefaultResourceVersionPolicyCodeMap maps from string to cpb.ResourceVersionPolicyCode_Value.
var DefaultResourceVersionPolicyCodeMap = map[string]cpb.ResourceVersionPolicyCode_Value{
	"INVALID_UNINITIALIZED": cpb.ResourceVersionPolicyCode_INVALID_UNINITIALIZED, // Enum 0
	"NO_VERSION":            cpb.ResourceVersionPolicyCode_NO_VERSION,            // Enum 1
	"VERSIONED":             cpb.ResourceVersionPolicyCode_VERSIONED,             // Enum 2
	"VERSIONED_UPDATE":      cpb.ResourceVersionPolicyCode_VERSIONED_UPDATE,      // Enum 3
}

// DefaultResponseTypeCodeMap maps from string to cpb.ResponseTypeCode_Value.
var DefaultResponseTypeCodeMap = map[string]cpb.ResponseTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ResponseTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"FATAL_ERROR":           cpb.ResponseTypeCode_FATAL_ERROR,           // Enum 3
	"OK":                    cpb.ResponseTypeCode_OK,                    // Enum 1
	"TRANSIENT_ERROR":       cpb.ResponseTypeCode_TRANSIENT_ERROR,       // Enum 2
}

// DefaultRestfulCapabilityModeCodeMap maps from string to cpb.RestfulCapabilityModeCode_Value.
var DefaultRestfulCapabilityModeCodeMap = map[string]cpb.RestfulCapabilityModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.RestfulCapabilityModeCode_INVALID_UNINITIALIZED, // Enum 0
	"CLIENT":                cpb.RestfulCapabilityModeCode_CLIENT,                // Enum 1
	"SERVER":                cpb.RestfulCapabilityModeCode_SERVER,                // Enum 2
}

// DefaultSPDXLicenseCodeMap maps from string to cpb.SPDXLicenseCode_Value.
var DefaultSPDXLicenseCodeMap = map[string]cpb.SPDXLicenseCode_Value{
	"INVALID_UNINITIALIZED":                cpb.SPDXLicenseCode_INVALID_UNINITIALIZED,                // Enum 0
	"AAL":                                  cpb.SPDXLicenseCode_AAL,                                  // Enum 3
	"ABSTYLES":                             cpb.SPDXLicenseCode_ABSTYLES,                             // Enum 4
	"ADOBE_2006":                           cpb.SPDXLicenseCode_ADOBE_2006,                           // Enum 5
	"ADOBE_GLYPH":                          cpb.SPDXLicenseCode_ADOBE_GLYPH,                          // Enum 6
	"ADSL":                                 cpb.SPDXLicenseCode_ADSL,                                 // Enum 7
	"AFL_1_1":                              cpb.SPDXLicenseCode_AFL_1_1,                              // Enum 8
	"AFL_1_2":                              cpb.SPDXLicenseCode_AFL_1_2,                              // Enum 9
	"AFL_2_0":                              cpb.SPDXLicenseCode_AFL_2_0,                              // Enum 10
	"AFL_2_1":                              cpb.SPDXLicenseCode_AFL_2_1,                              // Enum 11
	"AFL_3_0":                              cpb.SPDXLicenseCode_AFL_3_0,                              // Enum 12
	"AFMPARSE":                             cpb.SPDXLicenseCode_AFMPARSE,                             // Enum 13
	"AGPL_1_0_ONLY":                        cpb.SPDXLicenseCode_AGPL_1_0_ONLY,                        // Enum 14
	"AGPL_1_0_OR_LATER":                    cpb.SPDXLicenseCode_AGPL_1_0_OR_LATER,                    // Enum 15
	"AGPL_3_0_ONLY":                        cpb.SPDXLicenseCode_AGPL_3_0_ONLY,                        // Enum 16
	"AGPL_3_0_OR_LATER":                    cpb.SPDXLicenseCode_AGPL_3_0_OR_LATER,                    // Enum 17
	"ALADDIN":                              cpb.SPDXLicenseCode_ALADDIN,                              // Enum 18
	"AMDPLPA":                              cpb.SPDXLicenseCode_AMDPLPA,                              // Enum 19
	"AML":                                  cpb.SPDXLicenseCode_AML,                                  // Enum 20
	"AMPAS":                                cpb.SPDXLicenseCode_AMPAS,                                // Enum 21
	"ANTLR_PD":                             cpb.SPDXLicenseCode_ANTLR_PD,                             // Enum 22
	"APACHE_1_0":                           cpb.SPDXLicenseCode_APACHE_1_0,                           // Enum 23
	"APACHE_1_1":                           cpb.SPDXLicenseCode_APACHE_1_1,                           // Enum 24
	"APACHE_2_0":                           cpb.SPDXLicenseCode_APACHE_2_0,                           // Enum 25
	"APAFML":                               cpb.SPDXLicenseCode_APAFML,                               // Enum 26
	"APL_1_0":                              cpb.SPDXLicenseCode_APL_1_0,                              // Enum 27
	"APSL_1_0":                             cpb.SPDXLicenseCode_APSL_1_0,                             // Enum 28
	"APSL_1_1":                             cpb.SPDXLicenseCode_APSL_1_1,                             // Enum 29
	"APSL_1_2":                             cpb.SPDXLicenseCode_APSL_1_2,                             // Enum 30
	"APSL_2_0":                             cpb.SPDXLicenseCode_APSL_2_0,                             // Enum 31
	"ARTISTIC_1_0":                         cpb.SPDXLicenseCode_ARTISTIC_1_0,                         // Enum 34
	"ARTISTIC_1_0_CL8":                     cpb.SPDXLicenseCode_ARTISTIC_1_0_CL8,                     // Enum 32
	"ARTISTIC_1_0_PERL":                    cpb.SPDXLicenseCode_ARTISTIC_1_0_PERL,                    // Enum 33
	"ARTISTIC_2_0":                         cpb.SPDXLicenseCode_ARTISTIC_2_0,                         // Enum 35
	"BAHYPH":                               cpb.SPDXLicenseCode_BAHYPH,                               // Enum 36
	"BARR":                                 cpb.SPDXLicenseCode_BARR,                                 // Enum 37
	"BEERWARE":                             cpb.SPDXLicenseCode_BEERWARE,                             // Enum 38
	"BIT_TORRENT_1_0":                      cpb.SPDXLicenseCode_BIT_TORRENT_1_0,                      // Enum 39
	"BIT_TORRENT_1_1":                      cpb.SPDXLicenseCode_BIT_TORRENT_1_1,                      // Enum 40
	"BORCEUX":                              cpb.SPDXLicenseCode_BORCEUX,                              // Enum 41
	"BSD_1_CLAUSE":                         cpb.SPDXLicenseCode_BSD_1_CLAUSE,                         // Enum 42
	"BSD_2_CLAUSE":                         cpb.SPDXLicenseCode_BSD_2_CLAUSE,                         // Enum 46
	"BSD_2_CLAUSE_FREE_BSD":                cpb.SPDXLicenseCode_BSD_2_CLAUSE_FREE_BSD,                // Enum 43
	"BSD_2_CLAUSE_NET_BSD":                 cpb.SPDXLicenseCode_BSD_2_CLAUSE_NET_BSD,                 // Enum 44
	"BSD_2_CLAUSE_PATENT":                  cpb.SPDXLicenseCode_BSD_2_CLAUSE_PATENT,                  // Enum 45
	"BSD_3_CLAUSE":                         cpb.SPDXLicenseCode_BSD_3_CLAUSE,                         // Enum 53
	"BSD_3_CLAUSE_ATTRIBUTION":             cpb.SPDXLicenseCode_BSD_3_CLAUSE_ATTRIBUTION,             // Enum 47
	"BSD_3_CLAUSE_CLEAR":                   cpb.SPDXLicenseCode_BSD_3_CLAUSE_CLEAR,                   // Enum 48
	"BSD_3_CLAUSE_LBNL":                    cpb.SPDXLicenseCode_BSD_3_CLAUSE_LBNL,                    // Enum 49
	"BSD_3_CLAUSE_NO_NUCLEAR_LICENSE":      cpb.SPDXLicenseCode_BSD_3_CLAUSE_NO_NUCLEAR_LICENSE,      // Enum 51
	"BSD_3_CLAUSE_NO_NUCLEAR_LICENSE_2014": cpb.SPDXLicenseCode_BSD_3_CLAUSE_NO_NUCLEAR_LICENSE_2014, // Enum 50
	"BSD_3_CLAUSE_NO_NUCLEAR_WARRANTY":     cpb.SPDXLicenseCode_BSD_3_CLAUSE_NO_NUCLEAR_WARRANTY,     // Enum 52
	"BSD_4_CLAUSE":                         cpb.SPDXLicenseCode_BSD_4_CLAUSE,                         // Enum 55
	"BSD_4_CLAUSE_UC":                      cpb.SPDXLicenseCode_BSD_4_CLAUSE_UC,                      // Enum 54
	"BSD_PROTECTION":                       cpb.SPDXLicenseCode_BSD_PROTECTION,                       // Enum 56
	"BSD_SOURCE_CODE":                      cpb.SPDXLicenseCode_BSD_SOURCE_CODE,                      // Enum 57
	"BSD_ZERO_CLAUSE_LICENSE":              cpb.SPDXLicenseCode_BSD_ZERO_CLAUSE_LICENSE,              // Enum 2
	"BSL_1_0":                              cpb.SPDXLicenseCode_BSL_1_0,                              // Enum 58
	"BZIP2_1_0_5":                          cpb.SPDXLicenseCode_BZIP2_1_0_5,                          // Enum 59
	"BZIP2_1_0_6":                          cpb.SPDXLicenseCode_BZIP2_1_0_6,                          // Enum 60
	"CALDERA":                              cpb.SPDXLicenseCode_CALDERA,                              // Enum 61
	"CATOSL_1_1":                           cpb.SPDXLicenseCode_CATOSL_1_1,                           // Enum 62
	"CC0_1_0":                              cpb.SPDXLicenseCode_CC0_1_0,                              // Enum 93
	"CC_BY_1_0":                            cpb.SPDXLicenseCode_CC_BY_1_0,                            // Enum 63
	"CC_BY_2_0":                            cpb.SPDXLicenseCode_CC_BY_2_0,                            // Enum 64
	"CC_BY_2_5":                            cpb.SPDXLicenseCode_CC_BY_2_5,                            // Enum 65
	"CC_BY_3_0":                            cpb.SPDXLicenseCode_CC_BY_3_0,                            // Enum 66
	"CC_BY_4_0":                            cpb.SPDXLicenseCode_CC_BY_4_0,                            // Enum 67
	"CC_BY_NC_1_0":                         cpb.SPDXLicenseCode_CC_BY_NC_1_0,                         // Enum 68
	"CC_BY_NC_2_0":                         cpb.SPDXLicenseCode_CC_BY_NC_2_0,                         // Enum 69
	"CC_BY_NC_2_5":                         cpb.SPDXLicenseCode_CC_BY_NC_2_5,                         // Enum 70
	"CC_BY_NC_3_0":                         cpb.SPDXLicenseCode_CC_BY_NC_3_0,                         // Enum 71
	"CC_BY_NC_4_0":                         cpb.SPDXLicenseCode_CC_BY_NC_4_0,                         // Enum 72
	"CC_BY_NC_ND_1_0":                      cpb.SPDXLicenseCode_CC_BY_NC_ND_1_0,                      // Enum 73
	"CC_BY_NC_ND_2_0":                      cpb.SPDXLicenseCode_CC_BY_NC_ND_2_0,                      // Enum 74
	"CC_BY_NC_ND_2_5":                      cpb.SPDXLicenseCode_CC_BY_NC_ND_2_5,                      // Enum 75
	"CC_BY_NC_ND_3_0":                      cpb.SPDXLicenseCode_CC_BY_NC_ND_3_0,                      // Enum 76
	"CC_BY_NC_ND_4_0":                      cpb.SPDXLicenseCode_CC_BY_NC_ND_4_0,                      // Enum 77
	"CC_BY_NC_SA_1_0":                      cpb.SPDXLicenseCode_CC_BY_NC_SA_1_0,                      // Enum 78
	"CC_BY_NC_SA_2_0":                      cpb.SPDXLicenseCode_CC_BY_NC_SA_2_0,                      // Enum 79
	"CC_BY_NC_SA_2_5":                      cpb.SPDXLicenseCode_CC_BY_NC_SA_2_5,                      // Enum 80
	"CC_BY_NC_SA_3_0":                      cpb.SPDXLicenseCode_CC_BY_NC_SA_3_0,                      // Enum 81
	"CC_BY_NC_SA_4_0":                      cpb.SPDXLicenseCode_CC_BY_NC_SA_4_0,                      // Enum 82
	"CC_BY_ND_1_0":                         cpb.SPDXLicenseCode_CC_BY_ND_1_0,                         // Enum 83
	"CC_BY_ND_2_0":                         cpb.SPDXLicenseCode_CC_BY_ND_2_0,                         // Enum 84
	"CC_BY_ND_2_5":                         cpb.SPDXLicenseCode_CC_BY_ND_2_5,                         // Enum 85
	"CC_BY_ND_3_0":                         cpb.SPDXLicenseCode_CC_BY_ND_3_0,                         // Enum 86
	"CC_BY_ND_4_0":                         cpb.SPDXLicenseCode_CC_BY_ND_4_0,                         // Enum 87
	"CC_BY_SA_1_0":                         cpb.SPDXLicenseCode_CC_BY_SA_1_0,                         // Enum 88
	"CC_BY_SA_2_0":                         cpb.SPDXLicenseCode_CC_BY_SA_2_0,                         // Enum 89
	"CC_BY_SA_2_5":                         cpb.SPDXLicenseCode_CC_BY_SA_2_5,                         // Enum 90
	"CC_BY_SA_3_0":                         cpb.SPDXLicenseCode_CC_BY_SA_3_0,                         // Enum 91
	"CC_BY_SA_4_0":                         cpb.SPDXLicenseCode_CC_BY_SA_4_0,                         // Enum 92
	"CDDL_1_0":                             cpb.SPDXLicenseCode_CDDL_1_0,                             // Enum 94
	"CDDL_1_1":                             cpb.SPDXLicenseCode_CDDL_1_1,                             // Enum 95
	"CDLA_PERMISSIVE_1_0":                  cpb.SPDXLicenseCode_CDLA_PERMISSIVE_1_0,                  // Enum 96
	"CDLA_SHARING_1_0":                     cpb.SPDXLicenseCode_CDLA_SHARING_1_0,                     // Enum 97
	"CECILL_1_0":                           cpb.SPDXLicenseCode_CECILL_1_0,                           // Enum 98
	"CECILL_1_1":                           cpb.SPDXLicenseCode_CECILL_1_1,                           // Enum 99
	"CECILL_2_0":                           cpb.SPDXLicenseCode_CECILL_2_0,                           // Enum 100
	"CECILL_2_1":                           cpb.SPDXLicenseCode_CECILL_2_1,                           // Enum 101
	"CECILL_B":                             cpb.SPDXLicenseCode_CECILL_B,                             // Enum 102
	"CECILL_C":                             cpb.SPDXLicenseCode_CECILL_C,                             // Enum 103
	"CL_ARTISTIC":                          cpb.SPDXLicenseCode_CL_ARTISTIC,                          // Enum 104
	"CNRI_JYTHON":                          cpb.SPDXLicenseCode_CNRI_JYTHON,                          // Enum 105
	"CNRI_PYTHON":                          cpb.SPDXLicenseCode_CNRI_PYTHON,                          // Enum 107
	"CNRI_PYTHON_GPL_COMPATIBLE":           cpb.SPDXLicenseCode_CNRI_PYTHON_GPL_COMPATIBLE,           // Enum 106
	"CONDOR_1_1":                           cpb.SPDXLicenseCode_CONDOR_1_1,                           // Enum 108
	"CPAL_1_0":                             cpb.SPDXLicenseCode_CPAL_1_0,                             // Enum 109
	"CPL_1_0":                              cpb.SPDXLicenseCode_CPL_1_0,                              // Enum 110
	"CPOL_1_02":                            cpb.SPDXLicenseCode_CPOL_1_02,                            // Enum 111
	"CROSSWORD":                            cpb.SPDXLicenseCode_CROSSWORD,                            // Enum 112
	"CRYSTAL_STACKER":                      cpb.SPDXLicenseCode_CRYSTAL_STACKER,                      // Enum 113
	"CUA_OPL_1_0":                          cpb.SPDXLicenseCode_CUA_OPL_1_0,                          // Enum 114
	"CUBE":                                 cpb.SPDXLicenseCode_CUBE,                                 // Enum 115
	"CURL":                                 cpb.SPDXLicenseCode_CURL,                                 // Enum 116
	"DIFFMARK":                             cpb.SPDXLicenseCode_DIFFMARK,                             // Enum 118
	"DOC":                                  cpb.SPDXLicenseCode_DOC,                                  // Enum 119
	"DOTSEQN":                              cpb.SPDXLicenseCode_DOTSEQN,                              // Enum 120
	"DSDP":                                 cpb.SPDXLicenseCode_DSDP,                                 // Enum 121
	"DVIPDFM":                              cpb.SPDXLicenseCode_DVIPDFM,                              // Enum 122
	"D_FSL_1_0":                            cpb.SPDXLicenseCode_D_FSL_1_0,                            // Enum 117
	"ECL_1_0":                              cpb.SPDXLicenseCode_ECL_1_0,                              // Enum 123
	"ECL_2_0":                              cpb.SPDXLicenseCode_ECL_2_0,                              // Enum 124
	"EFL_1_0":                              cpb.SPDXLicenseCode_EFL_1_0,                              // Enum 125
	"EFL_2_0":                              cpb.SPDXLicenseCode_EFL_2_0,                              // Enum 126
	"ENTESSA":                              cpb.SPDXLicenseCode_ENTESSA,                              // Enum 128
	"EPL_1_0":                              cpb.SPDXLicenseCode_EPL_1_0,                              // Enum 129
	"EPL_2_0":                              cpb.SPDXLicenseCode_EPL_2_0,                              // Enum 130
	"ERL_PL_1_1":                           cpb.SPDXLicenseCode_ERL_PL_1_1,                           // Enum 131
	"EUPL_1_0":                             cpb.SPDXLicenseCode_EUPL_1_0,                             // Enum 133
	"EUPL_1_1":                             cpb.SPDXLicenseCode_EUPL_1_1,                             // Enum 134
	"EUPL_1_2":                             cpb.SPDXLicenseCode_EUPL_1_2,                             // Enum 135
	"EUROSYM":                              cpb.SPDXLicenseCode_EUROSYM,                              // Enum 136
	"EU_DATAGRID":                          cpb.SPDXLicenseCode_EU_DATAGRID,                          // Enum 132
	"E_GENIX":                              cpb.SPDXLicenseCode_E_GENIX,                              // Enum 127
	"FAIR":                                 cpb.SPDXLicenseCode_FAIR,                                 // Enum 137
	"FRAMEWORX_1_0":                        cpb.SPDXLicenseCode_FRAMEWORX_1_0,                        // Enum 138
	"FREE_IMAGE":                           cpb.SPDXLicenseCode_FREE_IMAGE,                           // Enum 139
	"FSFAP":                                cpb.SPDXLicenseCode_FSFAP,                                // Enum 140
	"FSFUL":                                cpb.SPDXLicenseCode_FSFUL,                                // Enum 141
	"FSFULLR":                              cpb.SPDXLicenseCode_FSFULLR,                              // Enum 142
	"FTL":                                  cpb.SPDXLicenseCode_FTL,                                  // Enum 143
	"GFDL_1_1_ONLY":                        cpb.SPDXLicenseCode_GFDL_1_1_ONLY,                        // Enum 144
	"GFDL_1_1_OR_LATER":                    cpb.SPDXLicenseCode_GFDL_1_1_OR_LATER,                    // Enum 145
	"GFDL_1_2_ONLY":                        cpb.SPDXLicenseCode_GFDL_1_2_ONLY,                        // Enum 146
	"GFDL_1_2_OR_LATER":                    cpb.SPDXLicenseCode_GFDL_1_2_OR_LATER,                    // Enum 147
	"GFDL_1_3_ONLY":                        cpb.SPDXLicenseCode_GFDL_1_3_ONLY,                        // Enum 148
	"GFDL_1_3_OR_LATER":                    cpb.SPDXLicenseCode_GFDL_1_3_OR_LATER,                    // Enum 149
	"GIFTWARE":                             cpb.SPDXLicenseCode_GIFTWARE,                             // Enum 150
	"GL2PS":                                cpb.SPDXLicenseCode_GL2PS,                                // Enum 151
	"GLIDE":                                cpb.SPDXLicenseCode_GLIDE,                                // Enum 152
	"GLULXE":                               cpb.SPDXLicenseCode_GLULXE,                               // Enum 153
	"GNUPLOT":                              cpb.SPDXLicenseCode_GNUPLOT,                              // Enum 154
	"GPL_1_0_ONLY":                         cpb.SPDXLicenseCode_GPL_1_0_ONLY,                         // Enum 155
	"GPL_1_0_OR_LATER":                     cpb.SPDXLicenseCode_GPL_1_0_OR_LATER,                     // Enum 156
	"GPL_2_0_ONLY":                         cpb.SPDXLicenseCode_GPL_2_0_ONLY,                         // Enum 157
	"GPL_2_0_OR_LATER":                     cpb.SPDXLicenseCode_GPL_2_0_OR_LATER,                     // Enum 158
	"GPL_3_0_ONLY":                         cpb.SPDXLicenseCode_GPL_3_0_ONLY,                         // Enum 159
	"GPL_3_0_OR_LATER":                     cpb.SPDXLicenseCode_GPL_3_0_OR_LATER,                     // Enum 160
	"G_SOAP_1_3B":                          cpb.SPDXLicenseCode_G_SOAP_1_3B,                          // Enum 161
	"HASKELL_REPORT":                       cpb.SPDXLicenseCode_HASKELL_REPORT,                       // Enum 162
	"HPND":                                 cpb.SPDXLicenseCode_HPND,                                 // Enum 163
	"IBM_PIBS":                             cpb.SPDXLicenseCode_IBM_PIBS,                             // Enum 164
	"ICU":                                  cpb.SPDXLicenseCode_ICU,                                  // Enum 165
	"IJG":                                  cpb.SPDXLicenseCode_IJG,                                  // Enum 166
	"IMAGE_MAGICK":                         cpb.SPDXLicenseCode_IMAGE_MAGICK,                         // Enum 167
	"IMLIB2":                               cpb.SPDXLicenseCode_IMLIB2,                               // Enum 169
	"INFO_ZIP":                             cpb.SPDXLicenseCode_INFO_ZIP,                             // Enum 170
	"INTEL":                                cpb.SPDXLicenseCode_INTEL,                                // Enum 172
	"INTEL_ACPI":                           cpb.SPDXLicenseCode_INTEL_ACPI,                           // Enum 171
	"INTERBASE_1_0":                        cpb.SPDXLicenseCode_INTERBASE_1_0,                        // Enum 173
	"IPA":                                  cpb.SPDXLicenseCode_IPA,                                  // Enum 174
	"IPL_1_0":                              cpb.SPDXLicenseCode_IPL_1_0,                              // Enum 175
	"ISC":                                  cpb.SPDXLicenseCode_ISC,                                  // Enum 176
	"I_MATIX":                              cpb.SPDXLicenseCode_I_MATIX,                              // Enum 168
	"JAS_PER_2_0":                          cpb.SPDXLicenseCode_JAS_PER_2_0,                          // Enum 177
	"JSON":                                 cpb.SPDXLicenseCode_JSON,                                 // Enum 178
	"LAL_1_2":                              cpb.SPDXLicenseCode_LAL_1_2,                              // Enum 179
	"LAL_1_3":                              cpb.SPDXLicenseCode_LAL_1_3,                              // Enum 180
	"LATEX2E":                              cpb.SPDXLicenseCode_LATEX2E,                              // Enum 181
	"LEPTONICA":                            cpb.SPDXLicenseCode_LEPTONICA,                            // Enum 182
	"LGPLLR":                               cpb.SPDXLicenseCode_LGPLLR,                               // Enum 189
	"LGPL_2_0_ONLY":                        cpb.SPDXLicenseCode_LGPL_2_0_ONLY,                        // Enum 183
	"LGPL_2_0_OR_LATER":                    cpb.SPDXLicenseCode_LGPL_2_0_OR_LATER,                    // Enum 184
	"LGPL_2_1_ONLY":                        cpb.SPDXLicenseCode_LGPL_2_1_ONLY,                        // Enum 185
	"LGPL_2_1_OR_LATER":                    cpb.SPDXLicenseCode_LGPL_2_1_OR_LATER,                    // Enum 186
	"LGPL_3_0_ONLY":                        cpb.SPDXLicenseCode_LGPL_3_0_ONLY,                        // Enum 187
	"LGPL_3_0_OR_LATER":                    cpb.SPDXLicenseCode_LGPL_3_0_OR_LATER,                    // Enum 188
	"LIBPNG":                               cpb.SPDXLicenseCode_LIBPNG,                               // Enum 190
	"LIBTIFF":                              cpb.SPDXLicenseCode_LIBTIFF,                              // Enum 191
	"LINUX_OPEN_IB":                        cpb.SPDXLicenseCode_LINUX_OPEN_IB,                        // Enum 195
	"LI_LI_Q_P_1_1":                        cpb.SPDXLicenseCode_LI_LI_Q_P_1_1,                        // Enum 192
	"LI_LI_Q_RPLUS_1_1":                    cpb.SPDXLicenseCode_LI_LI_Q_RPLUS_1_1,                    // Enum 194
	"LI_LI_Q_R_1_1":                        cpb.SPDXLicenseCode_LI_LI_Q_R_1_1,                        // Enum 193
	"LPL_1_0":                              cpb.SPDXLicenseCode_LPL_1_0,                              // Enum 196
	"LPL_1_02":                             cpb.SPDXLicenseCode_LPL_1_02,                             // Enum 197
	"LPPL_1_0":                             cpb.SPDXLicenseCode_LPPL_1_0,                             // Enum 198
	"LPPL_1_1":                             cpb.SPDXLicenseCode_LPPL_1_1,                             // Enum 199
	"LPPL_1_2":                             cpb.SPDXLicenseCode_LPPL_1_2,                             // Enum 200
	"LPPL_1_3A":                            cpb.SPDXLicenseCode_LPPL_1_3A,                            // Enum 201
	"LPPL_1_3C":                            cpb.SPDXLicenseCode_LPPL_1_3C,                            // Enum 202
	"MAKE_INDEX":                           cpb.SPDXLicenseCode_MAKE_INDEX,                           // Enum 203
	"MIR_OS":                               cpb.SPDXLicenseCode_MIR_OS,                               // Enum 204
	"MIT":                                  cpb.SPDXLicenseCode_MIT,                                  // Enum 210
	"MITNFA":                               cpb.SPDXLicenseCode_MITNFA,                               // Enum 211
	"MIT_0":                                cpb.SPDXLicenseCode_MIT_0,                                // Enum 205
	"MIT_ADVERTISING":                      cpb.SPDXLicenseCode_MIT_ADVERTISING,                      // Enum 206
	"MIT_CMU":                              cpb.SPDXLicenseCode_MIT_CMU,                              // Enum 207
	"MIT_ENNA":                             cpb.SPDXLicenseCode_MIT_ENNA,                             // Enum 208
	"MIT_FEH":                              cpb.SPDXLicenseCode_MIT_FEH,                              // Enum 209
	"MOTOSOTO":                             cpb.SPDXLicenseCode_MOTOSOTO,                             // Enum 212
	"MPICH2":                               cpb.SPDXLicenseCode_MPICH2,                               // Enum 213
	"MPL_1_0":                              cpb.SPDXLicenseCode_MPL_1_0,                              // Enum 214
	"MPL_1_1":                              cpb.SPDXLicenseCode_MPL_1_1,                              // Enum 215
	"MPL_2_0":                              cpb.SPDXLicenseCode_MPL_2_0,                              // Enum 217
	"MPL_2_0_NO_COPYLEFT_EXCEPTION":        cpb.SPDXLicenseCode_MPL_2_0_NO_COPYLEFT_EXCEPTION,        // Enum 216
	"MS_PL":                                cpb.SPDXLicenseCode_MS_PL,                                // Enum 218
	"MS_RL":                                cpb.SPDXLicenseCode_MS_RL,                                // Enum 219
	"MTLL":                                 cpb.SPDXLicenseCode_MTLL,                                 // Enum 220
	"MULTICS":                              cpb.SPDXLicenseCode_MULTICS,                              // Enum 221
	"MUP":                                  cpb.SPDXLicenseCode_MUP,                                  // Enum 222
	"NASA_1_3":                             cpb.SPDXLicenseCode_NASA_1_3,                             // Enum 223
	"NAUMEN":                               cpb.SPDXLicenseCode_NAUMEN,                               // Enum 224
	"NBPL_1_0":                             cpb.SPDXLicenseCode_NBPL_1_0,                             // Enum 225
	"NCSA":                                 cpb.SPDXLicenseCode_NCSA,                                 // Enum 226
	"NET_CDF":                              cpb.SPDXLicenseCode_NET_CDF,                              // Enum 228
	"NET_SNMP":                             cpb.SPDXLicenseCode_NET_SNMP,                             // Enum 227
	"NEWSLETR":                             cpb.SPDXLicenseCode_NEWSLETR,                             // Enum 229
	"NGPL":                                 cpb.SPDXLicenseCode_NGPL,                                 // Enum 230
	"NLOD_1_0":                             cpb.SPDXLicenseCode_NLOD_1_0,                             // Enum 231
	"NLPL":                                 cpb.SPDXLicenseCode_NLPL,                                 // Enum 232
	"NOKIA":                                cpb.SPDXLicenseCode_NOKIA,                                // Enum 233
	"NOSL":                                 cpb.SPDXLicenseCode_NOSL,                                 // Enum 234
	"NOT_OPEN_SOURCE":                      cpb.SPDXLicenseCode_NOT_OPEN_SOURCE,                      // Enum 1
	"NOWEB":                                cpb.SPDXLicenseCode_NOWEB,                                // Enum 235
	"NPL_1_0":                              cpb.SPDXLicenseCode_NPL_1_0,                              // Enum 236
	"NPL_1_1":                              cpb.SPDXLicenseCode_NPL_1_1,                              // Enum 237
	"NPOSL_3_0":                            cpb.SPDXLicenseCode_NPOSL_3_0,                            // Enum 238
	"NRL":                                  cpb.SPDXLicenseCode_NRL,                                  // Enum 239
	"NTP":                                  cpb.SPDXLicenseCode_NTP,                                  // Enum 240
	"OCCT_PL":                              cpb.SPDXLicenseCode_OCCT_PL,                              // Enum 241
	"OCLC_2_0":                             cpb.SPDXLicenseCode_OCLC_2_0,                             // Enum 242
	"OFL_1_0":                              cpb.SPDXLicenseCode_OFL_1_0,                              // Enum 244
	"OFL_1_1":                              cpb.SPDXLicenseCode_OFL_1_1,                              // Enum 245
	"OGTSL":                                cpb.SPDXLicenseCode_OGTSL,                                // Enum 246
	"OLDAP_1_1":                            cpb.SPDXLicenseCode_OLDAP_1_1,                            // Enum 247
	"OLDAP_1_2":                            cpb.SPDXLicenseCode_OLDAP_1_2,                            // Enum 248
	"OLDAP_1_3":                            cpb.SPDXLicenseCode_OLDAP_1_3,                            // Enum 249
	"OLDAP_1_4":                            cpb.SPDXLicenseCode_OLDAP_1_4,                            // Enum 250
	"OLDAP_2_0":                            cpb.SPDXLicenseCode_OLDAP_2_0,                            // Enum 252
	"OLDAP_2_0_1":                          cpb.SPDXLicenseCode_OLDAP_2_0_1,                          // Enum 251
	"OLDAP_2_1":                            cpb.SPDXLicenseCode_OLDAP_2_1,                            // Enum 253
	"OLDAP_2_2":                            cpb.SPDXLicenseCode_OLDAP_2_2,                            // Enum 256
	"OLDAP_2_2_1":                          cpb.SPDXLicenseCode_OLDAP_2_2_1,                          // Enum 254
	"OLDAP_2_2_2":                          cpb.SPDXLicenseCode_OLDAP_2_2_2,                          // Enum 255
	"OLDAP_2_3":                            cpb.SPDXLicenseCode_OLDAP_2_3,                            // Enum 257
	"OLDAP_2_4":                            cpb.SPDXLicenseCode_OLDAP_2_4,                            // Enum 258
	"OLDAP_2_5":                            cpb.SPDXLicenseCode_OLDAP_2_5,                            // Enum 259
	"OLDAP_2_6":                            cpb.SPDXLicenseCode_OLDAP_2_6,                            // Enum 260
	"OLDAP_2_7":                            cpb.SPDXLicenseCode_OLDAP_2_7,                            // Enum 261
	"OLDAP_2_8":                            cpb.SPDXLicenseCode_OLDAP_2_8,                            // Enum 262
	"OML":                                  cpb.SPDXLicenseCode_OML,                                  // Enum 263
	"OPEN_SSL":                             cpb.SPDXLicenseCode_OPEN_SSL,                             // Enum 264
	"OPL_1_0":                              cpb.SPDXLicenseCode_OPL_1_0,                              // Enum 265
	"OSET_PL_2_1":                          cpb.SPDXLicenseCode_OSET_PL_2_1,                          // Enum 266
	"OSL_1_0":                              cpb.SPDXLicenseCode_OSL_1_0,                              // Enum 267
	"OSL_1_1":                              cpb.SPDXLicenseCode_OSL_1_1,                              // Enum 268
	"OSL_2_0":                              cpb.SPDXLicenseCode_OSL_2_0,                              // Enum 269
	"OSL_2_1":                              cpb.SPDXLicenseCode_OSL_2_1,                              // Enum 270
	"OSL_3_0":                              cpb.SPDXLicenseCode_OSL_3_0,                              // Enum 271
	"O_DB_L_1_0":                           cpb.SPDXLicenseCode_O_DB_L_1_0,                           // Enum 243
	"PDDL_1_0":                             cpb.SPDXLicenseCode_PDDL_1_0,                             // Enum 272
	"PHP_3_0":                              cpb.SPDXLicenseCode_PHP_3_0,                              // Enum 273
	"PHP_3_01":                             cpb.SPDXLicenseCode_PHP_3_01,                             // Enum 274
	"PLEXUS":                               cpb.SPDXLicenseCode_PLEXUS,                               // Enum 275
	"POSTGRE_SQL":                          cpb.SPDXLicenseCode_POSTGRE_SQL,                          // Enum 276
	"PSFRAG":                               cpb.SPDXLicenseCode_PSFRAG,                               // Enum 277
	"PSUTILS":                              cpb.SPDXLicenseCode_PSUTILS,                              // Enum 278
	"PYTHON_2_0":                           cpb.SPDXLicenseCode_PYTHON_2_0,                           // Enum 279
	"QHULL":                                cpb.SPDXLicenseCode_QHULL,                                // Enum 280
	"QPL_1_0":                              cpb.SPDXLicenseCode_QPL_1_0,                              // Enum 281
	"RDISC":                                cpb.SPDXLicenseCode_RDISC,                                // Enum 282
	"RPL_1_1":                              cpb.SPDXLicenseCode_RPL_1_1,                              // Enum 284
	"RPL_1_5":                              cpb.SPDXLicenseCode_RPL_1_5,                              // Enum 285
	"RPSL_1_0":                             cpb.SPDXLicenseCode_RPSL_1_0,                             // Enum 286
	"RSA_MD":                               cpb.SPDXLicenseCode_RSA_MD,                               // Enum 287
	"RSCPL":                                cpb.SPDXLicenseCode_RSCPL,                                // Enum 288
	"RUBY":                                 cpb.SPDXLicenseCode_RUBY,                                 // Enum 289
	"R_HE_COS_1_1":                         cpb.SPDXLicenseCode_R_HE_COS_1_1,                         // Enum 283
	"SAXPATH":                              cpb.SPDXLicenseCode_SAXPATH,                              // Enum 291
	"SAX_PD":                               cpb.SPDXLicenseCode_SAX_PD,                               // Enum 290
	"SCEA":                                 cpb.SPDXLicenseCode_SCEA,                                 // Enum 292
	"SENDMAIL":                             cpb.SPDXLicenseCode_SENDMAIL,                             // Enum 293
	"SGI_B_1_0":                            cpb.SPDXLicenseCode_SGI_B_1_0,                            // Enum 294
	"SGI_B_1_1":                            cpb.SPDXLicenseCode_SGI_B_1_1,                            // Enum 295
	"SGI_B_2_0":                            cpb.SPDXLicenseCode_SGI_B_2_0,                            // Enum 296
	"SIM_PL_2_0":                           cpb.SPDXLicenseCode_SIM_PL_2_0,                           // Enum 297
	"SISSL":                                cpb.SPDXLicenseCode_SISSL,                                // Enum 299
	"SISSL_1_2":                            cpb.SPDXLicenseCode_SISSL_1_2,                            // Enum 298
	"SLEEPYCAT":                            cpb.SPDXLicenseCode_SLEEPYCAT,                            // Enum 300
	"SMLNJ":                                cpb.SPDXLicenseCode_SMLNJ,                                // Enum 301
	"SMPPL":                                cpb.SPDXLicenseCode_SMPPL,                                // Enum 302
	"SNIA":                                 cpb.SPDXLicenseCode_SNIA,                                 // Enum 303
	"SPENCER_86":                           cpb.SPDXLicenseCode_SPENCER_86,                           // Enum 304
	"SPENCER_94":                           cpb.SPDXLicenseCode_SPENCER_94,                           // Enum 305
	"SPENCER_99":                           cpb.SPDXLicenseCode_SPENCER_99,                           // Enum 306
	"SPL_1_0":                              cpb.SPDXLicenseCode_SPL_1_0,                              // Enum 307
	"SUGAR_CRM_1_1_3":                      cpb.SPDXLicenseCode_SUGAR_CRM_1_1_3,                      // Enum 308
	"SWL":                                  cpb.SPDXLicenseCode_SWL,                                  // Enum 309
	"TCL":                                  cpb.SPDXLicenseCode_TCL,                                  // Enum 310
	"TCP_WRAPPERS":                         cpb.SPDXLicenseCode_TCP_WRAPPERS,                         // Enum 311
	"TORQUE_1_1":                           cpb.SPDXLicenseCode_TORQUE_1_1,                           // Enum 313
	"TOSL":                                 cpb.SPDXLicenseCode_TOSL,                                 // Enum 314
	"T_MATE":                               cpb.SPDXLicenseCode_T_MATE,                               // Enum 312
	"UNICODE_DFS_2015":                     cpb.SPDXLicenseCode_UNICODE_DFS_2015,                     // Enum 315
	"UNICODE_DFS_2016":                     cpb.SPDXLicenseCode_UNICODE_DFS_2016,                     // Enum 316
	"UNICODE_TOU":                          cpb.SPDXLicenseCode_UNICODE_TOU,                          // Enum 317
	"UNLICENSE":                            cpb.SPDXLicenseCode_UNLICENSE,                            // Enum 318
	"UPL_1_0":                              cpb.SPDXLicenseCode_UPL_1_0,                              // Enum 319
	"VIM":                                  cpb.SPDXLicenseCode_VIM,                                  // Enum 320
	"VOSTROM":                              cpb.SPDXLicenseCode_VOSTROM,                              // Enum 321
	"VSL_1_0":                              cpb.SPDXLicenseCode_VSL_1_0,                              // Enum 322
	"W3C":                                  cpb.SPDXLicenseCode_W3C,                                  // Enum 325
	"W3C_19980720":                         cpb.SPDXLicenseCode_W3C_19980720,                         // Enum 323
	"W3C_20150513":                         cpb.SPDXLicenseCode_W3C_20150513,                         // Enum 324
	"WATCOM_1_0":                           cpb.SPDXLicenseCode_WATCOM_1_0,                           // Enum 326
	"WSUIPA":                               cpb.SPDXLicenseCode_WSUIPA,                               // Enum 327
	"WTFPL":                                cpb.SPDXLicenseCode_WTFPL,                                // Enum 328
	"X11":                                  cpb.SPDXLicenseCode_X11,                                  // Enum 329
	"XEROX":                                cpb.SPDXLicenseCode_XEROX,                                // Enum 330
	"XINETD":                               cpb.SPDXLicenseCode_XINETD,                               // Enum 332
	"XNET":                                 cpb.SPDXLicenseCode_XNET,                                 // Enum 333
	"XPP":                                  cpb.SPDXLicenseCode_XPP,                                  // Enum 334
	"X_FREE86_1_1":                         cpb.SPDXLicenseCode_X_FREE86_1_1,                         // Enum 331
	"X_SKAT":                               cpb.SPDXLicenseCode_X_SKAT,                               // Enum 335
	"YPL_1_0":                              cpb.SPDXLicenseCode_YPL_1_0,                              // Enum 336
	"YPL_1_1":                              cpb.SPDXLicenseCode_YPL_1_1,                              // Enum 337
	"ZED":                                  cpb.SPDXLicenseCode_ZED,                                  // Enum 338
	"ZEND_2_0":                             cpb.SPDXLicenseCode_ZEND_2_0,                             // Enum 339
	"ZIMBRA_1_3":                           cpb.SPDXLicenseCode_ZIMBRA_1_3,                           // Enum 340
	"ZIMBRA_1_4":                           cpb.SPDXLicenseCode_ZIMBRA_1_4,                           // Enum 341
	"ZLIB":                                 cpb.SPDXLicenseCode_ZLIB,                                 // Enum 343
	"ZLIB_ACKNOWLEDGEMENT":                 cpb.SPDXLicenseCode_ZLIB_ACKNOWLEDGEMENT,                 // Enum 342
	"ZPL_1_1":                              cpb.SPDXLicenseCode_ZPL_1_1,                              // Enum 344
	"ZPL_2_0":                              cpb.SPDXLicenseCode_ZPL_2_0,                              // Enum 345
	"ZPL_2_1":                              cpb.SPDXLicenseCode_ZPL_2_1,                              // Enum 346
}

// DefaultSearchComparatorCodeMap maps from string to cpb.SearchComparatorCode_Value.
var DefaultSearchComparatorCodeMap = map[string]cpb.SearchComparatorCode_Value{
	"INVALID_UNINITIALIZED": cpb.SearchComparatorCode_INVALID_UNINITIALIZED, // Enum 0
	"AP":                    cpb.SearchComparatorCode_AP,                    // Enum 9
	"EB":                    cpb.SearchComparatorCode_EB,                    // Enum 8
	"EQ":                    cpb.SearchComparatorCode_EQ,                    // Enum 1
	"GE":                    cpb.SearchComparatorCode_GE,                    // Enum 5
	"GT":                    cpb.SearchComparatorCode_GT,                    // Enum 3
	"LE":                    cpb.SearchComparatorCode_LE,                    // Enum 6
	"LT":                    cpb.SearchComparatorCode_LT,                    // Enum 4
	"NE":                    cpb.SearchComparatorCode_NE,                    // Enum 2
	"SA":                    cpb.SearchComparatorCode_SA,                    // Enum 7
}

// DefaultSearchEntryModeCodeMap maps from string to cpb.SearchEntryModeCode_Value.
var DefaultSearchEntryModeCodeMap = map[string]cpb.SearchEntryModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.SearchEntryModeCode_INVALID_UNINITIALIZED, // Enum 0
	"INCLUDE":               cpb.SearchEntryModeCode_INCLUDE,               // Enum 2
	"MATCH":                 cpb.SearchEntryModeCode_MATCH,                 // Enum 1
	"OUTCOME":               cpb.SearchEntryModeCode_OUTCOME,               // Enum 3
}

// DefaultSearchModifierCodeMap maps from string to cpb.SearchModifierCode_Value.
var DefaultSearchModifierCodeMap = map[string]cpb.SearchModifierCode_Value{
	"INVALID_UNINITIALIZED": cpb.SearchModifierCode_INVALID_UNINITIALIZED, // Enum 0
	"ABOVE":                 cpb.SearchModifierCode_ABOVE,                 // Enum 9
	"BELOW":                 cpb.SearchModifierCode_BELOW,                 // Enum 8
	"CONTAINS":              cpb.SearchModifierCode_CONTAINS,              // Enum 3
	"EXACT":                 cpb.SearchModifierCode_EXACT,                 // Enum 2
	"IDENTIFIER":            cpb.SearchModifierCode_IDENTIFIER,            // Enum 11
	"IN":                    cpb.SearchModifierCode_IN,                    // Enum 6
	"MISSING":               cpb.SearchModifierCode_MISSING,               // Enum 1
	"NOT":                   cpb.SearchModifierCode_NOT,                   // Enum 4
	"NOT_IN":                cpb.SearchModifierCode_NOT_IN,                // Enum 7
	"OF_TYPE":               cpb.SearchModifierCode_OF_TYPE,               // Enum 12
	"TEXT":                  cpb.SearchModifierCode_TEXT,                  // Enum 5
	"TYPE":                  cpb.SearchModifierCode_TYPE,                  // Enum 10
}

// DefaultSearchParamTypeCodeMap maps from string to cpb.SearchParamTypeCode_Value.
var DefaultSearchParamTypeCodeMap = map[string]cpb.SearchParamTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.SearchParamTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"COMPOSITE":             cpb.SearchParamTypeCode_COMPOSITE,             // Enum 6
	"DATE":                  cpb.SearchParamTypeCode_DATE,                  // Enum 2
	"NUMBER":                cpb.SearchParamTypeCode_NUMBER,                // Enum 1
	"QUANTITY":              cpb.SearchParamTypeCode_QUANTITY,              // Enum 7
	"REFERENCE":             cpb.SearchParamTypeCode_REFERENCE,             // Enum 5
	"SPECIAL":               cpb.SearchParamTypeCode_SPECIAL,               // Enum 9
	"STRING":                cpb.SearchParamTypeCode_STRING,                // Enum 3
	"TOKEN":                 cpb.SearchParamTypeCode_TOKEN,                 // Enum 4
	"URI":                   cpb.SearchParamTypeCode_URI,                   // Enum 8
}

// DefaultSequenceTypeCodeMap maps from string to cpb.SequenceTypeCode_Value.
var DefaultSequenceTypeCodeMap = map[string]cpb.SequenceTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.SequenceTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"AA":                    cpb.SequenceTypeCode_AA,                    // Enum 1
	"DNA":                   cpb.SequenceTypeCode_DNA,                   // Enum 2
	"RNA":                   cpb.SequenceTypeCode_RNA,                   // Enum 3
}

// DefaultSlicingRulesCodeMap maps from string to cpb.SlicingRulesCode_Value.
var DefaultSlicingRulesCodeMap = map[string]cpb.SlicingRulesCode_Value{
	"INVALID_UNINITIALIZED": cpb.SlicingRulesCode_INVALID_UNINITIALIZED, // Enum 0
	"CLOSED":                cpb.SlicingRulesCode_CLOSED,                // Enum 1
	"OPEN":                  cpb.SlicingRulesCode_OPEN,                  // Enum 2
	"OPEN_AT_END":           cpb.SlicingRulesCode_OPEN_AT_END,           // Enum 3
}

// DefaultSlotStatusCodeMap maps from string to cpb.SlotStatusCode_Value.
var DefaultSlotStatusCodeMap = map[string]cpb.SlotStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.SlotStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"BUSY":                  cpb.SlotStatusCode_BUSY,                  // Enum 1
	"BUSY_TENTATIVE":        cpb.SlotStatusCode_BUSY_TENTATIVE,        // Enum 4
	"BUSY_UNAVAILABLE":      cpb.SlotStatusCode_BUSY_UNAVAILABLE,      // Enum 3
	"ENTERED_IN_ERROR":      cpb.SlotStatusCode_ENTERED_IN_ERROR,      // Enum 5
	"FREE":                  cpb.SlotStatusCode_FREE,                  // Enum 2
}

// DefaultSmartCapabilitiesCodeMap maps from string to cpb.SmartCapabilitiesCode_Value.
var DefaultSmartCapabilitiesCodeMap = map[string]cpb.SmartCapabilitiesCode_Value{
	"INVALID_UNINITIALIZED":         cpb.SmartCapabilitiesCode_INVALID_UNINITIALIZED,         // Enum 0
	"CLIENT_CONFIDENTIAL_SYMMETRIC": cpb.SmartCapabilitiesCode_CLIENT_CONFIDENTIAL_SYMMETRIC, // Enum 4
	"CLIENT_PUBLIC":                 cpb.SmartCapabilitiesCode_CLIENT_PUBLIC,                 // Enum 3
	"CONTEXT_EHR_ENCOUNTER":         cpb.SmartCapabilitiesCode_CONTEXT_EHR_ENCOUNTER,         // Enum 9
	"CONTEXT_EHR_PATIENT":           cpb.SmartCapabilitiesCode_CONTEXT_EHR_PATIENT,           // Enum 8
	"CONTEXT_PASSTHROUGH_BANNER":    cpb.SmartCapabilitiesCode_CONTEXT_PASSTHROUGH_BANNER,    // Enum 6
	"CONTEXT_PASSTHROUGH_STYLE":     cpb.SmartCapabilitiesCode_CONTEXT_PASSTHROUGH_STYLE,     // Enum 7
	"CONTEXT_STANDALONE_ENCOUNTER":  cpb.SmartCapabilitiesCode_CONTEXT_STANDALONE_ENCOUNTER,  // Enum 11
	"CONTEXT_STANDALONE_PATIENT":    cpb.SmartCapabilitiesCode_CONTEXT_STANDALONE_PATIENT,    // Enum 10
	"LAUNCH_EHR":                    cpb.SmartCapabilitiesCode_LAUNCH_EHR,                    // Enum 1
	"LAUNCH_STANDALONE":             cpb.SmartCapabilitiesCode_LAUNCH_STANDALONE,             // Enum 2
	"PERMISSION_OFFLINE":            cpb.SmartCapabilitiesCode_PERMISSION_OFFLINE,            // Enum 12
	"PERMISSION_PATIENT":            cpb.SmartCapabilitiesCode_PERMISSION_PATIENT,            // Enum 13
	"PERMISSION_USER":               cpb.SmartCapabilitiesCode_PERMISSION_USER,               // Enum 14
	"SSO_OPENID_CONNECT":            cpb.SmartCapabilitiesCode_SSO_OPENID_CONNECT,            // Enum 5
}

// DefaultSortDirectionCodeMap maps from string to cpb.SortDirectionCode_Value.
var DefaultSortDirectionCodeMap = map[string]cpb.SortDirectionCode_Value{
	"INVALID_UNINITIALIZED": cpb.SortDirectionCode_INVALID_UNINITIALIZED, // Enum 0
	"ASCENDING":             cpb.SortDirectionCode_ASCENDING,             // Enum 1
	"DESCENDING":            cpb.SortDirectionCode_DESCENDING,            // Enum 2
}

// DefaultSpecimenContainedPreferenceCodeMap maps from string to cpb.SpecimenContainedPreferenceCode_Value.
var DefaultSpecimenContainedPreferenceCodeMap = map[string]cpb.SpecimenContainedPreferenceCode_Value{
	"INVALID_UNINITIALIZED": cpb.SpecimenContainedPreferenceCode_INVALID_UNINITIALIZED, // Enum 0
	"ALTERNATE":             cpb.SpecimenContainedPreferenceCode_ALTERNATE,             // Enum 2
	"PREFERRED":             cpb.SpecimenContainedPreferenceCode_PREFERRED,             // Enum 1
}

// DefaultSpecimenStatusCodeMap maps from string to cpb.SpecimenStatusCode_Value.
var DefaultSpecimenStatusCodeMap = map[string]cpb.SpecimenStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.SpecimenStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"AVAILABLE":             cpb.SpecimenStatusCode_AVAILABLE,             // Enum 1
	"ENTERED_IN_ERROR":      cpb.SpecimenStatusCode_ENTERED_IN_ERROR,      // Enum 4
	"UNAVAILABLE":           cpb.SpecimenStatusCode_UNAVAILABLE,           // Enum 2
	"UNSATISFACTORY":        cpb.SpecimenStatusCode_UNSATISFACTORY,        // Enum 3
}

// DefaultStandardsStatusCodeMap maps from string to cpb.StandardsStatusCode_Value.
var DefaultStandardsStatusCodeMap = map[string]cpb.StandardsStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.StandardsStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"DEPRECATED":            cpb.StandardsStatusCode_DEPRECATED,            // Enum 5
	"DRAFT":                 cpb.StandardsStatusCode_DRAFT,                 // Enum 1
	"EXTERNAL":              cpb.StandardsStatusCode_EXTERNAL,              // Enum 6
	"INFORMATIVE":           cpb.StandardsStatusCode_INFORMATIVE,           // Enum 4
	"NORMATIVE":             cpb.StandardsStatusCode_NORMATIVE,             // Enum 2
	"TRIAL_USE":             cpb.StandardsStatusCode_TRIAL_USE,             // Enum 3
}

// DefaultStatusCodeMap maps from string to cpb.StatusCode_Value.
var DefaultStatusCodeMap = map[string]cpb.StatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.StatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ATTESTED":              cpb.StatusCode_ATTESTED,              // Enum 1
	"IN_PROCESS":            cpb.StatusCode_IN_PROCESS,            // Enum 3
	"REQ_REVALID":           cpb.StatusCode_REQ_REVALID,           // Enum 4
	"REVAL_FAIL":            cpb.StatusCode_REVAL_FAIL,            // Enum 6
	"VALIDATED":             cpb.StatusCode_VALIDATED,             // Enum 2
	"VAL_FAIL":              cpb.StatusCode_VAL_FAIL,              // Enum 5
}

// DefaultStrandTypeCodeMap maps from string to cpb.StrandTypeCode_Value.
var DefaultStrandTypeCodeMap = map[string]cpb.StrandTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.StrandTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"CRICK":                 cpb.StrandTypeCode_CRICK,                 // Enum 2
	"WATSON":                cpb.StrandTypeCode_WATSON,                // Enum 1
}

// DefaultStructureDefinitionKindCodeMap maps from string to cpb.StructureDefinitionKindCode_Value.
var DefaultStructureDefinitionKindCodeMap = map[string]cpb.StructureDefinitionKindCode_Value{
	"INVALID_UNINITIALIZED": cpb.StructureDefinitionKindCode_INVALID_UNINITIALIZED, // Enum 0
	"COMPLEX_TYPE":          cpb.StructureDefinitionKindCode_COMPLEX_TYPE,          // Enum 2
	"LOGICAL":               cpb.StructureDefinitionKindCode_LOGICAL,               // Enum 4
	"PRIMITIVE_TYPE":        cpb.StructureDefinitionKindCode_PRIMITIVE_TYPE,        // Enum 1
	"RESOURCE":              cpb.StructureDefinitionKindCode_RESOURCE,              // Enum 3
}

// DefaultStructureMapContextTypeCodeMap maps from string to cpb.StructureMapContextTypeCode_Value.
var DefaultStructureMapContextTypeCodeMap = map[string]cpb.StructureMapContextTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.StructureMapContextTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"TYPE":                  cpb.StructureMapContextTypeCode_TYPE,                  // Enum 1
	"VARIABLE":              cpb.StructureMapContextTypeCode_VARIABLE,              // Enum 2
}

// DefaultStructureMapGroupTypeModeCodeMap maps from string to cpb.StructureMapGroupTypeModeCode_Value.
var DefaultStructureMapGroupTypeModeCodeMap = map[string]cpb.StructureMapGroupTypeModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.StructureMapGroupTypeModeCode_INVALID_UNINITIALIZED, // Enum 0
	"NONE":                  cpb.StructureMapGroupTypeModeCode_NONE,                  // Enum 1
	"TYPES":                 cpb.StructureMapGroupTypeModeCode_TYPES,                 // Enum 2
	"TYPE_AND_TYPES":        cpb.StructureMapGroupTypeModeCode_TYPE_AND_TYPES,        // Enum 3
}

// DefaultStructureMapInputModeCodeMap maps from string to cpb.StructureMapInputModeCode_Value.
var DefaultStructureMapInputModeCodeMap = map[string]cpb.StructureMapInputModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.StructureMapInputModeCode_INVALID_UNINITIALIZED, // Enum 0
	"SOURCE":                cpb.StructureMapInputModeCode_SOURCE,                // Enum 1
	"TARGET":                cpb.StructureMapInputModeCode_TARGET,                // Enum 2
}

// DefaultStructureMapModelModeCodeMap maps from string to cpb.StructureMapModelModeCode_Value.
var DefaultStructureMapModelModeCodeMap = map[string]cpb.StructureMapModelModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.StructureMapModelModeCode_INVALID_UNINITIALIZED, // Enum 0
	"PRODUCED":              cpb.StructureMapModelModeCode_PRODUCED,              // Enum 4
	"QUERIED":               cpb.StructureMapModelModeCode_QUERIED,               // Enum 2
	"SOURCE":                cpb.StructureMapModelModeCode_SOURCE,                // Enum 1
	"TARGET":                cpb.StructureMapModelModeCode_TARGET,                // Enum 3
}

// DefaultStructureMapSourceListModeCodeMap maps from string to cpb.StructureMapSourceListModeCode_Value.
var DefaultStructureMapSourceListModeCodeMap = map[string]cpb.StructureMapSourceListModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.StructureMapSourceListModeCode_INVALID_UNINITIALIZED, // Enum 0
	"FIRST":                 cpb.StructureMapSourceListModeCode_FIRST,                 // Enum 1
	"LAST":                  cpb.StructureMapSourceListModeCode_LAST,                  // Enum 3
	"NOT_FIRST":             cpb.StructureMapSourceListModeCode_NOT_FIRST,             // Enum 2
	"NOT_LAST":              cpb.StructureMapSourceListModeCode_NOT_LAST,              // Enum 4
	"ONLY_ONE":              cpb.StructureMapSourceListModeCode_ONLY_ONE,              // Enum 5
}

// DefaultStructureMapTargetListModeCodeMap maps from string to cpb.StructureMapTargetListModeCode_Value.
var DefaultStructureMapTargetListModeCodeMap = map[string]cpb.StructureMapTargetListModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.StructureMapTargetListModeCode_INVALID_UNINITIALIZED, // Enum 0
	"COLLATE":               cpb.StructureMapTargetListModeCode_COLLATE,               // Enum 4
	"FIRST":                 cpb.StructureMapTargetListModeCode_FIRST,                 // Enum 1
	"LAST":                  cpb.StructureMapTargetListModeCode_LAST,                  // Enum 3
	"SHARE":                 cpb.StructureMapTargetListModeCode_SHARE,                 // Enum 2
}

// DefaultStructureMapTransformCodeMap maps from string to cpb.StructureMapTransformCode_Value.
var DefaultStructureMapTransformCodeMap = map[string]cpb.StructureMapTransformCode_Value{
	"INVALID_UNINITIALIZED": cpb.StructureMapTransformCode_INVALID_UNINITIALIZED, // Enum 0
	"APPEND":                cpb.StructureMapTransformCode_APPEND,                // Enum 6
	"C":                     cpb.StructureMapTransformCode_C,                     // Enum 14
	"CAST":                  cpb.StructureMapTransformCode_CAST,                  // Enum 5
	"CC":                    cpb.StructureMapTransformCode_CC,                    // Enum 13
	"COPY":                  cpb.StructureMapTransformCode_COPY,                  // Enum 2
	"CP":                    cpb.StructureMapTransformCode_CP,                    // Enum 17
	"CREATE":                cpb.StructureMapTransformCode_CREATE,                // Enum 1
	"DATE_OP":               cpb.StructureMapTransformCode_DATE_OP,               // Enum 9
	"ESCAPE":                cpb.StructureMapTransformCode_ESCAPE,                // Enum 4
	"EVALUATE":              cpb.StructureMapTransformCode_EVALUATE,              // Enum 12
	"ID":                    cpb.StructureMapTransformCode_ID,                    // Enum 16
	"POINTER":               cpb.StructureMapTransformCode_POINTER,               // Enum 11
	"QTY":                   cpb.StructureMapTransformCode_QTY,                   // Enum 15
	"REFERENCE":             cpb.StructureMapTransformCode_REFERENCE,             // Enum 8
	"TRANSLATE":             cpb.StructureMapTransformCode_TRANSLATE,             // Enum 7
	"TRUNCATE":              cpb.StructureMapTransformCode_TRUNCATE,              // Enum 3
	"UUID":                  cpb.StructureMapTransformCode_UUID,                  // Enum 10
}

// DefaultSubscriptionChannelTypeCodeMap maps from string to cpb.SubscriptionChannelTypeCode_Value.
var DefaultSubscriptionChannelTypeCodeMap = map[string]cpb.SubscriptionChannelTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.SubscriptionChannelTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"EMAIL":                 cpb.SubscriptionChannelTypeCode_EMAIL,                 // Enum 3
	"MESSAGE":               cpb.SubscriptionChannelTypeCode_MESSAGE,               // Enum 5
	"REST_HOOK":             cpb.SubscriptionChannelTypeCode_REST_HOOK,             // Enum 1
	"SMS":                   cpb.SubscriptionChannelTypeCode_SMS,                   // Enum 4
	"WEBSOCKET":             cpb.SubscriptionChannelTypeCode_WEBSOCKET,             // Enum 2
}

// DefaultSubscriptionStatusCodeMap maps from string to cpb.SubscriptionStatusCode_Value.
var DefaultSubscriptionStatusCodeMap = map[string]cpb.SubscriptionStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.SubscriptionStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.SubscriptionStatusCode_ACTIVE,                // Enum 2
	"ERROR":                 cpb.SubscriptionStatusCode_ERROR,                 // Enum 3
	"OFF":                   cpb.SubscriptionStatusCode_OFF,                   // Enum 4
	"REQUESTED":             cpb.SubscriptionStatusCode_REQUESTED,             // Enum 1
}

// DefaultSupplyDeliveryStatusCodeMap maps from string to cpb.SupplyDeliveryStatusCode_Value.
var DefaultSupplyDeliveryStatusCodeMap = map[string]cpb.SupplyDeliveryStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.SupplyDeliveryStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ABANDONED":             cpb.SupplyDeliveryStatusCode_ABANDONED,             // Enum 3
	"COMPLETED":             cpb.SupplyDeliveryStatusCode_COMPLETED,             // Enum 2
	"ENTERED_IN_ERROR":      cpb.SupplyDeliveryStatusCode_ENTERED_IN_ERROR,      // Enum 4
	"IN_PROGRESS":           cpb.SupplyDeliveryStatusCode_IN_PROGRESS,           // Enum 1
}

// DefaultSupplyItemTypeCodeMap maps from string to cpb.SupplyItemTypeCode_Value.
var DefaultSupplyItemTypeCodeMap = map[string]cpb.SupplyItemTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.SupplyItemTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"DEVICE":                cpb.SupplyItemTypeCode_DEVICE,                // Enum 2
	"MEDICATION":            cpb.SupplyItemTypeCode_MEDICATION,            // Enum 1
}

// DefaultSupplyRequestStatusCodeMap maps from string to cpb.SupplyRequestStatusCode_Value.
var DefaultSupplyRequestStatusCodeMap = map[string]cpb.SupplyRequestStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.SupplyRequestStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.SupplyRequestStatusCode_ACTIVE,                // Enum 2
	"CANCELLED":             cpb.SupplyRequestStatusCode_CANCELLED,             // Enum 4
	"COMPLETED":             cpb.SupplyRequestStatusCode_COMPLETED,             // Enum 5
	"DRAFT":                 cpb.SupplyRequestStatusCode_DRAFT,                 // Enum 1
	"ENTERED_IN_ERROR":      cpb.SupplyRequestStatusCode_ENTERED_IN_ERROR,      // Enum 6
	"SUSPENDED":             cpb.SupplyRequestStatusCode_SUSPENDED,             // Enum 3
	"UNKNOWN":               cpb.SupplyRequestStatusCode_UNKNOWN,               // Enum 7
}

// DefaultTaskIntentCodeMap maps from string to cpb.TaskIntentCode_Value.
var DefaultTaskIntentCodeMap = map[string]cpb.TaskIntentCode_Value{
	"INVALID_UNINITIALIZED": cpb.TaskIntentCode_INVALID_UNINITIALIZED, // Enum 0
	"UNKNOWN":               cpb.TaskIntentCode_UNKNOWN,               // Enum 1
}

// DefaultTaskStatusCodeMap maps from string to cpb.TaskStatusCode_Value.
var DefaultTaskStatusCodeMap = map[string]cpb.TaskStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.TaskStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"ACCEPTED":              cpb.TaskStatusCode_ACCEPTED,              // Enum 4
	"CANCELLED":             cpb.TaskStatusCode_CANCELLED,             // Enum 7
	"COMPLETED":             cpb.TaskStatusCode_COMPLETED,             // Enum 11
	"DRAFT":                 cpb.TaskStatusCode_DRAFT,                 // Enum 1
	"ENTERED_IN_ERROR":      cpb.TaskStatusCode_ENTERED_IN_ERROR,      // Enum 12
	"FAILED":                cpb.TaskStatusCode_FAILED,                // Enum 10
	"IN_PROGRESS":           cpb.TaskStatusCode_IN_PROGRESS,           // Enum 8
	"ON_HOLD":               cpb.TaskStatusCode_ON_HOLD,               // Enum 9
	"READY":                 cpb.TaskStatusCode_READY,                 // Enum 6
	"RECEIVED":              cpb.TaskStatusCode_RECEIVED,              // Enum 3
	"REJECTED":              cpb.TaskStatusCode_REJECTED,              // Enum 5
	"REQUESTED":             cpb.TaskStatusCode_REQUESTED,             // Enum 2
}

// DefaultTemplateStatusCodeLifeCycleCodeMap maps from string to cpb.TemplateStatusCodeLifeCycleCode_Value.
var DefaultTemplateStatusCodeLifeCycleCodeMap = map[string]cpb.TemplateStatusCodeLifeCycleCode_Value{
	"INVALID_UNINITIALIZED": cpb.TemplateStatusCodeLifeCycleCode_INVALID_UNINITIALIZED, // Enum 0
	"ACTIVE":                cpb.TemplateStatusCodeLifeCycleCode_ACTIVE,                // Enum 3
	"CANCELLED":             cpb.TemplateStatusCodeLifeCycleCode_CANCELLED,             // Enum 5
	"DRAFT":                 cpb.TemplateStatusCodeLifeCycleCode_DRAFT,                 // Enum 1
	"PENDING":               cpb.TemplateStatusCodeLifeCycleCode_PENDING,               // Enum 2
	"REJECTED":              cpb.TemplateStatusCodeLifeCycleCode_REJECTED,              // Enum 6
	"RETIRED":               cpb.TemplateStatusCodeLifeCycleCode_RETIRED,               // Enum 7
	"REVIEW":                cpb.TemplateStatusCodeLifeCycleCode_REVIEW,                // Enum 4
	"TERMINATED":            cpb.TemplateStatusCodeLifeCycleCode_TERMINATED,            // Enum 8
}

// DefaultTestReportActionResultCodeMap maps from string to cpb.TestReportActionResultCode_Value.
var DefaultTestReportActionResultCodeMap = map[string]cpb.TestReportActionResultCode_Value{
	"INVALID_UNINITIALIZED": cpb.TestReportActionResultCode_INVALID_UNINITIALIZED, // Enum 0
	"ERROR":                 cpb.TestReportActionResultCode_ERROR,                 // Enum 5
	"FAIL":                  cpb.TestReportActionResultCode_FAIL,                  // Enum 3
	"PASS":                  cpb.TestReportActionResultCode_PASS,                  // Enum 1
	"SKIP":                  cpb.TestReportActionResultCode_SKIP,                  // Enum 2
	"WARNING":               cpb.TestReportActionResultCode_WARNING,               // Enum 4
}

// DefaultTestReportParticipantTypeCodeMap maps from string to cpb.TestReportParticipantTypeCode_Value.
var DefaultTestReportParticipantTypeCodeMap = map[string]cpb.TestReportParticipantTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.TestReportParticipantTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"CLIENT":                cpb.TestReportParticipantTypeCode_CLIENT,                // Enum 2
	"SERVER":                cpb.TestReportParticipantTypeCode_SERVER,                // Enum 3
	"TEST_ENGINE":           cpb.TestReportParticipantTypeCode_TEST_ENGINE,           // Enum 1
}

// DefaultTestReportResultCodeMap maps from string to cpb.TestReportResultCode_Value.
var DefaultTestReportResultCodeMap = map[string]cpb.TestReportResultCode_Value{
	"INVALID_UNINITIALIZED": cpb.TestReportResultCode_INVALID_UNINITIALIZED, // Enum 0
	"FAIL":                  cpb.TestReportResultCode_FAIL,                  // Enum 2
	"PASS":                  cpb.TestReportResultCode_PASS,                  // Enum 1
	"PENDING":               cpb.TestReportResultCode_PENDING,               // Enum 3
}

// DefaultTestReportStatusCodeMap maps from string to cpb.TestReportStatusCode_Value.
var DefaultTestReportStatusCodeMap = map[string]cpb.TestReportStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.TestReportStatusCode_INVALID_UNINITIALIZED, // Enum 0
	"COMPLETED":             cpb.TestReportStatusCode_COMPLETED,             // Enum 1
	"ENTERED_IN_ERROR":      cpb.TestReportStatusCode_ENTERED_IN_ERROR,      // Enum 5
	"IN_PROGRESS":           cpb.TestReportStatusCode_IN_PROGRESS,           // Enum 2
	"STOPPED":               cpb.TestReportStatusCode_STOPPED,               // Enum 4
	"WAITING":               cpb.TestReportStatusCode_WAITING,               // Enum 3
}

// DefaultTestScriptRequestMethodCodeMap maps from string to cpb.TestScriptRequestMethodCode_Value.
var DefaultTestScriptRequestMethodCodeMap = map[string]cpb.TestScriptRequestMethodCode_Value{
	"INVALID_UNINITIALIZED": cpb.TestScriptRequestMethodCode_INVALID_UNINITIALIZED, // Enum 0
	"DELETE":                cpb.TestScriptRequestMethodCode_DELETE,                // Enum 1
	"GET":                   cpb.TestScriptRequestMethodCode_GET,                   // Enum 2
	"HEAD":                  cpb.TestScriptRequestMethodCode_HEAD,                  // Enum 7
	"OPTIONS":               cpb.TestScriptRequestMethodCode_OPTIONS,               // Enum 3
	"PATCH":                 cpb.TestScriptRequestMethodCode_PATCH,                 // Enum 4
	"POST":                  cpb.TestScriptRequestMethodCode_POST,                  // Enum 5
	"PUT":                   cpb.TestScriptRequestMethodCode_PUT,                   // Enum 6
}

// DefaultTriggerTypeCodeMap maps from string to cpb.TriggerTypeCode_Value.
var DefaultTriggerTypeCodeMap = map[string]cpb.TriggerTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.TriggerTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"DATA_ACCESSED":         cpb.TriggerTypeCode_DATA_ACCESSED,         // Enum 7
	"DATA_ACCESS_ENDED":     cpb.TriggerTypeCode_DATA_ACCESS_ENDED,     // Enum 8
	"DATA_ADDED":            cpb.TriggerTypeCode_DATA_ADDED,            // Enum 4
	"DATA_CHANGED":          cpb.TriggerTypeCode_DATA_CHANGED,          // Enum 3
	"DATA_MODIFIED":         cpb.TriggerTypeCode_DATA_MODIFIED,         // Enum 5
	"DATA_REMOVED":          cpb.TriggerTypeCode_DATA_REMOVED,          // Enum 6
	"NAMED_EVENT":           cpb.TriggerTypeCode_NAMED_EVENT,           // Enum 1
	"PERIODIC":              cpb.TriggerTypeCode_PERIODIC,              // Enum 2
}

// DefaultTypeDerivationRuleCodeMap maps from string to cpb.TypeDerivationRuleCode_Value.
var DefaultTypeDerivationRuleCodeMap = map[string]cpb.TypeDerivationRuleCode_Value{
	"INVALID_UNINITIALIZED": cpb.TypeDerivationRuleCode_INVALID_UNINITIALIZED, // Enum 0
	"CONSTRAINT":            cpb.TypeDerivationRuleCode_CONSTRAINT,            // Enum 2
	"SPECIALIZATION":        cpb.TypeDerivationRuleCode_SPECIALIZATION,        // Enum 1
}

// DefaultUDIEntryTypeCodeMap maps from string to cpb.UDIEntryTypeCode_Value.
var DefaultUDIEntryTypeCodeMap = map[string]cpb.UDIEntryTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.UDIEntryTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"BARCODE":               cpb.UDIEntryTypeCode_BARCODE,               // Enum 1
	"CARD":                  cpb.UDIEntryTypeCode_CARD,                  // Enum 4
	"MANUAL":                cpb.UDIEntryTypeCode_MANUAL,                // Enum 3
	"RFID":                  cpb.UDIEntryTypeCode_RFID,                  // Enum 2
	"SELF_REPORTED":         cpb.UDIEntryTypeCode_SELF_REPORTED,         // Enum 5
	"UNKNOWN":               cpb.UDIEntryTypeCode_UNKNOWN,               // Enum 6
}

// DefaultUseCodeMap maps from string to cpb.UseCode_Value.
var DefaultUseCodeMap = map[string]cpb.UseCode_Value{
	"INVALID_UNINITIALIZED": cpb.UseCode_INVALID_UNINITIALIZED, // Enum 0
	"CLAIM":                 cpb.UseCode_CLAIM,                 // Enum 1
	"PREAUTHORIZATION":      cpb.UseCode_PREAUTHORIZATION,      // Enum 2
	"PREDETERMINATION":      cpb.UseCode_PREDETERMINATION,      // Enum 3
}

// DefaultV20444CodeMap maps from string to cpb.V20444Code_Value.
var DefaultV20444CodeMap = map[string]cpb.V20444Code_Value{
	"INVALID_UNINITIALIZED": cpb.V20444Code_INVALID_UNINITIALIZED, // Enum 0
	"F":                     cpb.V20444Code_F,                     // Enum 1
	"G":                     cpb.V20444Code_G,                     // Enum 2
}

// DefaultV3AddressUseCodeMap maps from string to cpb.V3AddressUseCode_Value.
var DefaultV3AddressUseCodeMap = map[string]cpb.V3AddressUseCode_Value{
	"INVALID_UNINITIALIZED":         cpb.V3AddressUseCode_INVALID_UNINITIALIZED,         // Enum 0
	"AS":                            cpb.V3AddressUseCode_AS,                            // Enum 16
	"BAD":                           cpb.V3AddressUseCode_BAD,                           // Enum 2
	"CONF":                          cpb.V3AddressUseCode_CONF,                          // Enum 3
	"DIR":                           cpb.V3AddressUseCode_DIR,                           // Enum 10
	"EC":                            cpb.V3AddressUseCode_EC,                            // Enum 17
	"GENERAL_ADDRESS_USE":           cpb.V3AddressUseCode_GENERAL_ADDRESS_USE,           // Enum 1
	"H":                             cpb.V3AddressUseCode_H,                             // Enum 4
	"HP":                            cpb.V3AddressUseCode_HP,                            // Enum 5
	"HV":                            cpb.V3AddressUseCode_HV,                            // Enum 6
	"MC":                            cpb.V3AddressUseCode_MC,                            // Enum 18
	"OLD":                           cpb.V3AddressUseCode_OLD,                           // Enum 7
	"PG":                            cpb.V3AddressUseCode_PG,                            // Enum 19
	"PHYS":                          cpb.V3AddressUseCode_PHYS,                          // Enum 13
	"POSTAL_ADDRESS_USE":            cpb.V3AddressUseCode_POSTAL_ADDRESS_USE,            // Enum 12
	"PST":                           cpb.V3AddressUseCode_PST,                           // Enum 14
	"PUB":                           cpb.V3AddressUseCode_PUB,                           // Enum 11
	"TELECOMMUNICATION_ADDRESS_USE": cpb.V3AddressUseCode_TELECOMMUNICATION_ADDRESS_USE, // Enum 15
	"TMP":                           cpb.V3AddressUseCode_TMP,                           // Enum 8
	"WP":                            cpb.V3AddressUseCode_WP,                            // Enum 9
}

// DefaultV3ConfidentialityCodeMap maps from string to cpb.V3ConfidentialityCode_Value.
var DefaultV3ConfidentialityCodeMap = map[string]cpb.V3ConfidentialityCode_Value{
	"INVALID_UNINITIALIZED":          cpb.V3ConfidentialityCode_INVALID_UNINITIALIZED,          // Enum 0
	"B":                              cpb.V3ConfidentialityCode_B,                              // Enum 9
	"C":                              cpb.V3ConfidentialityCode_C,                              // Enum 18
	"CONFIDENTIALITY":                cpb.V3ConfidentialityCode_CONFIDENTIALITY,                // Enum 1
	"CONFIDENTIALITY_BY_ACCESS_KIND": cpb.V3ConfidentialityCode_CONFIDENTIALITY_BY_ACCESS_KIND, // Enum 8
	"CONFIDENTIALITY_BY_INFO_TYPE":   cpb.V3ConfidentialityCode_CONFIDENTIALITY_BY_INFO_TYPE,   // Enum 12
	"CONFIDENTIALITY_MODIFIERS":      cpb.V3ConfidentialityCode_CONFIDENTIALITY_MODIFIERS,      // Enum 17
	"D":                              cpb.V3ConfidentialityCode_D,                              // Enum 10
	"ETH":                            cpb.V3ConfidentialityCode_ETH,                            // Enum 13
	"HIV":                            cpb.V3ConfidentialityCode_HIV,                            // Enum 14
	"I":                              cpb.V3ConfidentialityCode_I,                              // Enum 11
	"L":                              cpb.V3ConfidentialityCode_L,                              // Enum 2
	"M":                              cpb.V3ConfidentialityCode_M,                              // Enum 3
	"N":                              cpb.V3ConfidentialityCode_N,                              // Enum 4
	"PSY":                            cpb.V3ConfidentialityCode_PSY,                            // Enum 15
	"R":                              cpb.V3ConfidentialityCode_R,                              // Enum 5
	"S":                              cpb.V3ConfidentialityCode_S,                              // Enum 19
	"SDV":                            cpb.V3ConfidentialityCode_SDV,                            // Enum 16
	"T":                              cpb.V3ConfidentialityCode_T,                              // Enum 20
	"U":                              cpb.V3ConfidentialityCode_U,                              // Enum 6
	"V":                              cpb.V3ConfidentialityCode_V,                              // Enum 7
}

// DefaultV3EntityNamePartQualifierCodeMap maps from string to cpb.V3EntityNamePartQualifierCode_Value.
var DefaultV3EntityNamePartQualifierCodeMap = map[string]cpb.V3EntityNamePartQualifierCode_Value{
	"INVALID_UNINITIALIZED":             cpb.V3EntityNamePartQualifierCode_INVALID_UNINITIALIZED,             // Enum 0
	"AC":                                cpb.V3EntityNamePartQualifierCode_AC,                                // Enum 2
	"AD":                                cpb.V3EntityNamePartQualifierCode_AD,                                // Enum 3
	"BR":                                cpb.V3EntityNamePartQualifierCode_BR,                                // Enum 4
	"CL":                                cpb.V3EntityNamePartQualifierCode_CL,                                // Enum 5
	"CON":                               cpb.V3EntityNamePartQualifierCode_CON,                               // Enum 14
	"DEV":                               cpb.V3EntityNamePartQualifierCode_DEV,                               // Enum 15
	"FLAV":                              cpb.V3EntityNamePartQualifierCode_FLAV,                              // Enum 16
	"FORMUL":                            cpb.V3EntityNamePartQualifierCode_FORMUL,                            // Enum 17
	"FRM":                               cpb.V3EntityNamePartQualifierCode_FRM,                               // Enum 18
	"IN":                                cpb.V3EntityNamePartQualifierCode_IN,                                // Enum 6
	"INV":                               cpb.V3EntityNamePartQualifierCode_INV,                               // Enum 19
	"LS":                                cpb.V3EntityNamePartQualifierCode_LS,                                // Enum 7
	"NB":                                cpb.V3EntityNamePartQualifierCode_NB,                                // Enum 8
	"ORGANIZATION_NAME_PART_QUALIFIER":  cpb.V3EntityNamePartQualifierCode_ORGANIZATION_NAME_PART_QUALIFIER,  // Enum 1
	"PERSON_NAME_PART_AFFIX_TYPES":      cpb.V3EntityNamePartQualifierCode_PERSON_NAME_PART_AFFIX_TYPES,      // Enum 27
	"PERSON_NAME_PART_CHANGE_QUALIFIER": cpb.V3EntityNamePartQualifierCode_PERSON_NAME_PART_CHANGE_QUALIFIER, // Enum 28
	"PERSON_NAME_PART_MISC_QUALIFIER":   cpb.V3EntityNamePartQualifierCode_PERSON_NAME_PART_MISC_QUALIFIER,   // Enum 29
	"PERSON_NAME_PART_QUALIFIER":        cpb.V3EntityNamePartQualifierCode_PERSON_NAME_PART_QUALIFIER,        // Enum 26
	"PHARMACEUTICAL_ENTITY_NAME_PART_QUALIFIERS": cpb.V3EntityNamePartQualifierCode_PHARMACEUTICAL_ENTITY_NAME_PART_QUALIFIERS, // Enum 13
	"POPUL": cpb.V3EntityNamePartQualifierCode_POPUL, // Enum 20
	"PR":    cpb.V3EntityNamePartQualifierCode_PR,    // Enum 9
	"SCI":   cpb.V3EntityNamePartQualifierCode_SCI,   // Enum 21
	"SP":    cpb.V3EntityNamePartQualifierCode_SP,    // Enum 10
	"STR":   cpb.V3EntityNamePartQualifierCode_STR,   // Enum 22
	"TIME":  cpb.V3EntityNamePartQualifierCode_TIME,  // Enum 23
	"TITLE": cpb.V3EntityNamePartQualifierCode_TITLE, // Enum 11
	"TMK":   cpb.V3EntityNamePartQualifierCode_TMK,   // Enum 24
	"USE":   cpb.V3EntityNamePartQualifierCode_USE,   // Enum 25
	"VV":    cpb.V3EntityNamePartQualifierCode_VV,    // Enum 12
}

// DefaultV3EntityNamePartQualifierR2CodeMap maps from string to cpb.V3EntityNamePartQualifierR2Code_Value.
var DefaultV3EntityNamePartQualifierR2CodeMap = map[string]cpb.V3EntityNamePartQualifierR2Code_Value{
	"INVALID_UNINITIALIZED": cpb.V3EntityNamePartQualifierR2Code_INVALID_UNINITIALIZED, // Enum 0
	"AC":                    cpb.V3EntityNamePartQualifierR2Code_AC,                    // Enum 24
	"AD":                    cpb.V3EntityNamePartQualifierR2Code_AD,                    // Enum 1
	"BR":                    cpb.V3EntityNamePartQualifierR2Code_BR,                    // Enum 3
	"CL":                    cpb.V3EntityNamePartQualifierR2Code_CL,                    // Enum 4
	"CON":                   cpb.V3EntityNamePartQualifierR2Code_CON,                   // Enum 10
	"DEV":                   cpb.V3EntityNamePartQualifierR2Code_DEV,                   // Enum 11
	"FLAV":                  cpb.V3EntityNamePartQualifierR2Code_FLAV,                  // Enum 12
	"FORMUL":                cpb.V3EntityNamePartQualifierR2Code_FORMUL,                // Enum 13
	"FRM":                   cpb.V3EntityNamePartQualifierR2Code_FRM,                   // Enum 14
	"HON":                   cpb.V3EntityNamePartQualifierR2Code_HON,                   // Enum 25
	"IN":                    cpb.V3EntityNamePartQualifierR2Code_IN,                    // Enum 5
	"INV":                   cpb.V3EntityNamePartQualifierR2Code_INV,                   // Enum 15
	"LS":                    cpb.V3EntityNamePartQualifierR2Code_LS,                    // Enum 6
	"MID":                   cpb.V3EntityNamePartQualifierR2Code_MID,                   // Enum 7
	"NB":                    cpb.V3EntityNamePartQualifierR2Code_NB,                    // Enum 26
	"PFX":                   cpb.V3EntityNamePartQualifierR2Code_PFX,                   // Enum 8
	"PHARMACEUTICAL_ENTITY_NAME_PART_QUALIFIERS": cpb.V3EntityNamePartQualifierR2Code_PHARMACEUTICAL_ENTITY_NAME_PART_QUALIFIERS, // Enum 9
	"POPUL":        cpb.V3EntityNamePartQualifierR2Code_POPUL,        // Enum 16
	"PR":           cpb.V3EntityNamePartQualifierR2Code_PR,           // Enum 27
	"SCI":          cpb.V3EntityNamePartQualifierR2Code_SCI,          // Enum 17
	"SFX":          cpb.V3EntityNamePartQualifierR2Code_SFX,          // Enum 22
	"SP":           cpb.V3EntityNamePartQualifierR2Code_SP,           // Enum 2
	"STR":          cpb.V3EntityNamePartQualifierR2Code_STR,          // Enum 18
	"TIME":         cpb.V3EntityNamePartQualifierR2Code_TIME,         // Enum 19
	"TITLE_STYLES": cpb.V3EntityNamePartQualifierR2Code_TITLE_STYLES, // Enum 23
	"TMK":          cpb.V3EntityNamePartQualifierR2Code_TMK,          // Enum 20
	"USE":          cpb.V3EntityNamePartQualifierR2Code_USE,          // Enum 21
}

// DefaultV3EntityNameUseCodeMap maps from string to cpb.V3EntityNameUseCode_Value.
var DefaultV3EntityNameUseCodeMap = map[string]cpb.V3EntityNameUseCode_Value{
	"INVALID_UNINITIALIZED":   cpb.V3EntityNameUseCode_INVALID_UNINITIALIZED,   // Enum 0
	"A":                       cpb.V3EntityNameUseCode_A,                       // Enum 11
	"ABC":                     cpb.V3EntityNameUseCode_ABC,                     // Enum 2
	"ASGN":                    cpb.V3EntityNameUseCode_ASGN,                    // Enum 5
	"C":                       cpb.V3EntityNameUseCode_C,                       // Enum 6
	"I":                       cpb.V3EntityNameUseCode_I,                       // Enum 7
	"IDE":                     cpb.V3EntityNameUseCode_IDE,                     // Enum 3
	"L":                       cpb.V3EntityNameUseCode_L,                       // Enum 8
	"NAME_REPRESENTATION_USE": cpb.V3EntityNameUseCode_NAME_REPRESENTATION_USE, // Enum 1
	"OR":                      cpb.V3EntityNameUseCode_OR,                      // Enum 9
	"P":                       cpb.V3EntityNameUseCode_P,                       // Enum 10
	"PHON":                    cpb.V3EntityNameUseCode_PHON,                    // Enum 14
	"R":                       cpb.V3EntityNameUseCode_R,                       // Enum 12
	"SNDX":                    cpb.V3EntityNameUseCode_SNDX,                    // Enum 15
	"SRCH":                    cpb.V3EntityNameUseCode_SRCH,                    // Enum 13
	"SYL":                     cpb.V3EntityNameUseCode_SYL,                     // Enum 4
}

// DefaultV3EntityNameUseR2CodeMap maps from string to cpb.V3EntityNameUseR2Code_Value.
var DefaultV3EntityNameUseR2CodeMap = map[string]cpb.V3EntityNameUseR2Code_Value{
	"INVALID_UNINITIALIZED":   cpb.V3EntityNameUseR2Code_INVALID_UNINITIALIZED,   // Enum 0
	"A":                       cpb.V3EntityNameUseR2Code_A,                       // Enum 2
	"ABC":                     cpb.V3EntityNameUseR2Code_ABC,                     // Enum 10
	"ANON":                    cpb.V3EntityNameUseR2Code_ANON,                    // Enum 3
	"ASSUMED":                 cpb.V3EntityNameUseR2Code_ASSUMED,                 // Enum 1
	"C":                       cpb.V3EntityNameUseR2Code_C,                       // Enum 7
	"DN":                      cpb.V3EntityNameUseR2Code_DN,                      // Enum 14
	"I":                       cpb.V3EntityNameUseR2Code_I,                       // Enum 4
	"IDE":                     cpb.V3EntityNameUseR2Code_IDE,                     // Enum 11
	"M":                       cpb.V3EntityNameUseR2Code_M,                       // Enum 8
	"NAME_REPRESENTATION_USE": cpb.V3EntityNameUseR2Code_NAME_REPRESENTATION_USE, // Enum 9
	"OLD":                     cpb.V3EntityNameUseR2Code_OLD,                     // Enum 13
	"OR":                      cpb.V3EntityNameUseR2Code_OR,                      // Enum 15
	"P":                       cpb.V3EntityNameUseR2Code_P,                       // Enum 5
	"PHON":                    cpb.V3EntityNameUseR2Code_PHON,                    // Enum 16
	"R":                       cpb.V3EntityNameUseR2Code_R,                       // Enum 6
	"SRCH":                    cpb.V3EntityNameUseR2Code_SRCH,                    // Enum 17
	"SYL":                     cpb.V3EntityNameUseR2Code_SYL,                     // Enum 12
	"T":                       cpb.V3EntityNameUseR2Code_T,                       // Enum 18
}

// DefaultV3NullFlavorCodeMap maps from string to cpb.V3NullFlavorCode_Value.
var DefaultV3NullFlavorCodeMap = map[string]cpb.V3NullFlavorCode_Value{
	"INVALID_UNINITIALIZED": cpb.V3NullFlavorCode_INVALID_UNINITIALIZED, // Enum 0
	"ASKU":                  cpb.V3NullFlavorCode_ASKU,                  // Enum 11
	"DER":                   cpb.V3NullFlavorCode_DER,                   // Enum 3
	"INV":                   cpb.V3NullFlavorCode_INV,                   // Enum 2
	"MSK":                   cpb.V3NullFlavorCode_MSK,                   // Enum 8
	"NA":                    cpb.V3NullFlavorCode_NA,                    // Enum 9
	"NASK":                  cpb.V3NullFlavorCode_NASK,                  // Enum 13
	"NAV":                   cpb.V3NullFlavorCode_NAV,                   // Enum 12
	"NAVU":                  cpb.V3NullFlavorCode_NAVU,                  // Enum 14
	"NI":                    cpb.V3NullFlavorCode_NI,                    // Enum 1
	"NINF":                  cpb.V3NullFlavorCode_NINF,                  // Enum 5
	"NP":                    cpb.V3NullFlavorCode_NP,                    // Enum 17
	"OTH":                   cpb.V3NullFlavorCode_OTH,                   // Enum 4
	"PINF":                  cpb.V3NullFlavorCode_PINF,                  // Enum 6
	"QS":                    cpb.V3NullFlavorCode_QS,                    // Enum 15
	"TRC":                   cpb.V3NullFlavorCode_TRC,                   // Enum 16
	"UNC":                   cpb.V3NullFlavorCode_UNC,                   // Enum 7
	"UNK":                   cpb.V3NullFlavorCode_UNK,                   // Enum 10
}

// DefaultV3ParticipationModeCodeMap maps from string to cpb.V3ParticipationModeCode_Value.
var DefaultV3ParticipationModeCodeMap = map[string]cpb.V3ParticipationModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.V3ParticipationModeCode_INVALID_UNINITIALIZED, // Enum 0
	"DICTATE":               cpb.V3ParticipationModeCode_DICTATE,               // Enum 5
	"ELECTRONIC":            cpb.V3ParticipationModeCode_ELECTRONIC,            // Enum 1
	"EMAILWRIT":             cpb.V3ParticipationModeCode_EMAILWRIT,             // Enum 14
	"FACE":                  cpb.V3ParticipationModeCode_FACE,                  // Enum 6
	"FAXWRIT":               cpb.V3ParticipationModeCode_FAXWRIT,               // Enum 10
	"HANDWRIT":              cpb.V3ParticipationModeCode_HANDWRIT,              // Enum 11
	"MAILWRIT":              cpb.V3ParticipationModeCode_MAILWRIT,              // Enum 12
	"ONLINEWRIT":            cpb.V3ParticipationModeCode_ONLINEWRIT,            // Enum 13
	"PHONE":                 cpb.V3ParticipationModeCode_PHONE,                 // Enum 7
	"PHYSICAL":              cpb.V3ParticipationModeCode_PHYSICAL,              // Enum 2
	"REMOTE":                cpb.V3ParticipationModeCode_REMOTE,                // Enum 3
	"TYPEWRIT":              cpb.V3ParticipationModeCode_TYPEWRIT,              // Enum 15
	"VERBAL":                cpb.V3ParticipationModeCode_VERBAL,                // Enum 4
	"VIDEOCONF":             cpb.V3ParticipationModeCode_VIDEOCONF,             // Enum 8
	"WRITTEN":               cpb.V3ParticipationModeCode_WRITTEN,               // Enum 9
}

// DefaultV3ProbabilityDistributionTypeCodeMap maps from string to cpb.V3ProbabilityDistributionTypeCode_Value.
var DefaultV3ProbabilityDistributionTypeCodeMap = map[string]cpb.V3ProbabilityDistributionTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.V3ProbabilityDistributionTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"B":                     cpb.V3ProbabilityDistributionTypeCode_B,                     // Enum 1
	"E":                     cpb.V3ProbabilityDistributionTypeCode_E,                     // Enum 2
	"F":                     cpb.V3ProbabilityDistributionTypeCode_F,                     // Enum 3
	"G":                     cpb.V3ProbabilityDistributionTypeCode_G,                     // Enum 4
	"LN":                    cpb.V3ProbabilityDistributionTypeCode_LN,                    // Enum 5
	"N":                     cpb.V3ProbabilityDistributionTypeCode_N,                     // Enum 6
	"T":                     cpb.V3ProbabilityDistributionTypeCode_T,                     // Enum 7
	"U":                     cpb.V3ProbabilityDistributionTypeCode_U,                     // Enum 8
	"X2":                    cpb.V3ProbabilityDistributionTypeCode_X2,                    // Enum 9
}

// DefaultV3RoleCodeMap maps from string to cpb.V3RoleCode_Value.
var DefaultV3RoleCodeMap = map[string]cpb.V3RoleCode_Value{
	"INVALID_UNINITIALIZED":            cpb.V3RoleCode_INVALID_UNINITIALIZED,            // Enum 0
	"ACC":                              cpb.V3RoleCode_ACC,                              // Enum 360
	"ACHFID":                           cpb.V3RoleCode_ACHFID,                           // Enum 61
	"ACTMIL":                           cpb.V3RoleCode_ACTMIL,                           // Enum 232
	"ADMINISTRATIVE_CONTACT_ROLE_TYPE": cpb.V3RoleCode_ADMINISTRATIVE_CONTACT_ROLE_TYPE, // Enum 53
	"ADOPT":                            cpb.V3RoleCode_ADOPT,                            // Enum 211
	"ADOPTF":                           cpb.V3RoleCode_ADOPTF,                           // Enum 154
	"ADOPTM":                           cpb.V3RoleCode_ADOPTM,                           // Enum 155
	"ADOPTP":                           cpb.V3RoleCode_ADOPTP,                           // Enum 153
	"AFFILIATION_ROLE_TYPE":            cpb.V3RoleCode_AFFILIATION_ROLE_TYPE,            // Enum 1
	"AGENT_ROLE_TYPE":                  cpb.V3RoleCode_AGENT_ROLE_TYPE,                  // Enum 2
	"ALL":                              cpb.V3RoleCode_ALL,                              // Enum 293
	"AMB":                              cpb.V3RoleCode_AMB,                              // Enum 357
	"AMENDER":                          cpb.V3RoleCode_AMENDER,                          // Enum 3
	"AMPUT":                            cpb.V3RoleCode_AMPUT,                            // Enum 294
	"ANTIBIOT":                         cpb.V3RoleCode_ANTIBIOT,                         // Enum 37
	"ASSIGNED_NON_PERSON_LIVING_SUBJECT_ROLE_TYPE": cpb.V3RoleCode_ASSIGNED_NON_PERSON_LIVING_SUBJECT_ROLE_TYPE, // Enum 34
	"ASSIGNED_ROLE_TYPE":                           cpb.V3RoleCode_ASSIGNED_ROLE_TYPE,                           // Enum 33
	"ASSIST":                                       cpb.V3RoleCode_ASSIST,                                       // Enum 35
	"AUNT":                                         cpb.V3RoleCode_AUNT,                                         // Enum 109
	"B":                                            cpb.V3RoleCode_B,                                            // Enum 373
	"BF":                                           cpb.V3RoleCode_BF,                                           // Enum 65
	"BILL":                                         cpb.V3RoleCode_BILL,                                         // Enum 54
	"BIOTH":                                        cpb.V3RoleCode_BIOTH,                                        // Enum 36
	"BL":                                           cpb.V3RoleCode_BL,                                           // Enum 66
	"BMTC":                                         cpb.V3RoleCode_BMTC,                                         // Enum 295
	"BMTU":                                         cpb.V3RoleCode_BMTU,                                         // Enum 266
	"BR":                                           cpb.V3RoleCode_BR,                                           // Enum 67
	"BREAST":                                       cpb.V3RoleCode_BREAST,                                       // Enum 296
	"BRO":                                          cpb.V3RoleCode_BRO,                                          // Enum 171
	"BROINLAW":                                     cpb.V3RoleCode_BROINLAW,                                     // Enum 144
	"C":                                            cpb.V3RoleCode_C,                                            // Enum 368
	"CANC":                                         cpb.V3RoleCode_CANC,                                         // Enum 297
	"CAPC":                                         cpb.V3RoleCode_CAPC,                                         // Enum 298
	"CARD":                                         cpb.V3RoleCode_CARD,                                         // Enum 299
	"CAS":                                          cpb.V3RoleCode_CAS,                                          // Enum 44
	"CASM":                                         cpb.V3RoleCode_CASM,                                         // Enum 45
	"CATH":                                         cpb.V3RoleCode_CATH,                                         // Enum 244
	"CCO":                                          cpb.V3RoleCode_CCO,                                          // Enum 39
	"CCU":                                          cpb.V3RoleCode_CCU,                                          // Enum 267
	"CERTIFIED_ENTITY_TYPE":                        cpb.V3RoleCode_CERTIFIED_ENTITY_TYPE,                        // Enum 42
	"CHEST":                                        cpb.V3RoleCode_CHEST,                                        // Enum 268
	"CHILD":                                        cpb.V3RoleCode_CHILD,                                        // Enum 93
	"CHLDADOPT":                                    cpb.V3RoleCode_CHLDADOPT,                                    // Enum 94
	"CHLDFOST":                                     cpb.V3RoleCode_CHLDFOST,                                     // Enum 97
	"CHLDINLAW":                                    cpb.V3RoleCode_CHLDINLAW,                                    // Enum 137
	"CHR":                                          cpb.V3RoleCode_CHR,                                          // Enum 252
	"CITIZEN_ROLE_TYPE":                            cpb.V3RoleCode_CITIZEN_ROLE_TYPE,                            // Enum 43
	"CLAIM":                                        cpb.V3RoleCode_CLAIM,                                        // Enum 379
	"CLAIMANT_COVERED_PARTY_ROLE_TYPE":             cpb.V3RoleCode_CLAIMANT_COVERED_PARTY_ROLE_TYPE,             // Enum 220
	"CLASSIFIER":                                   cpb.V3RoleCode_CLASSIFIER,                                   // Enum 4
	"CN":                                           cpb.V3RoleCode_CN,                                           // Enum 46
	"CNRP":                                         cpb.V3RoleCode_CNRP,                                         // Enum 47
	"CNRPM":                                        cpb.V3RoleCode_CNRPM,                                        // Enum 48
	"CO":                                           cpb.V3RoleCode_CO,                                           // Enum 68
	"COAG":                                         cpb.V3RoleCode_COAG,                                         // Enum 301
	"COCBEN":                                       cpb.V3RoleCode_COCBEN,                                       // Enum 224
	"COMM":                                         cpb.V3RoleCode_COMM,                                         // Enum 361
	"COMMUNITY_LABORATORY":                         cpb.V3RoleCode_COMMUNITY_LABORATORY,                         // Enum 380
	"CONSENTER":                                    cpb.V3RoleCode_CONSENTER,                                    // Enum 5
	"CONSWIT":                                      cpb.V3RoleCode_CONSWIT,                                      // Enum 6
	"CONTACT_ROLE_TYPE":                            cpb.V3RoleCode_CONTACT_ROLE_TYPE,                            // Enum 52
	"COPART":                                       cpb.V3RoleCode_COPART,                                       // Enum 7
	"COUSN":                                        cpb.V3RoleCode_COUSN,                                        // Enum 112
	"COVERAGE_ROLE_TYPE":                           cpb.V3RoleCode_COVERAGE_ROLE_TYPE,                           // Enum 202
	"COVERAGE_SPONSOR_ROLE_TYPE":                   cpb.V3RoleCode_COVERAGE_SPONSOR_ROLE_TYPE,                   // Enum 18
	"COVERED_PARTY_ROLE_TYPE":                      cpb.V3RoleCode_COVERED_PARTY_ROLE_TYPE,                      // Enum 219
	"CPCA":                                         cpb.V3RoleCode_CPCA,                                         // Enum 49
	"CRIMEVIC":                                     cpb.V3RoleCode_CRIMEVIC,                                     // Enum 221
	"CRP":                                          cpb.V3RoleCode_CRP,                                          // Enum 50
	"CRPM":                                         cpb.V3RoleCode_CRPM,                                         // Enum 51
	"CRS":                                          cpb.V3RoleCode_CRS,                                          // Enum 302
	"CSC":                                          cpb.V3RoleCode_CSC,                                          // Enum 362
	"CVDX":                                         cpb.V3RoleCode_CVDX,                                         // Enum 243
	"DA":                                           cpb.V3RoleCode_DA,                                           // Enum 69
	"DADDR":                                        cpb.V3RoleCode_DADDR,                                        // Enum 355
	"DAU":                                          cpb.V3RoleCode_DAU,                                          // Enum 101
	"DAUADOPT":                                     cpb.V3RoleCode_DAUADOPT,                                     // Enum 95
	"DAUC":                                         cpb.V3RoleCode_DAUC,                                         // Enum 100
	"DAUFOST":                                      cpb.V3RoleCode_DAUFOST,                                      // Enum 98
	"DAUINLAW":                                     cpb.V3RoleCode_DAUINLAW,                                     // Enum 138
	"DC":                                           cpb.V3RoleCode_DC,                                           // Enum 83
	"DEBR":                                         cpb.V3RoleCode_DEBR,                                         // Enum 38
	"DECLASSIFIER":                                 cpb.V3RoleCode_DECLASSIFIER,                                 // Enum 8
	"DEDICATED_CLINICAL_LOCATION_ROLE_TYPE":        cpb.V3RoleCode_DEDICATED_CLINICAL_LOCATION_ROLE_TYPE,        // Enum 241
	"DEDICATED_NON_CLINICAL_LOCATION_ROLE_TYPE":     cpb.V3RoleCode_DEDICATED_NON_CLINICAL_LOCATION_ROLE_TYPE,     // Enum 354
	"DEDICATED_SERVICE_DELIVERY_LOCATION_ROLE_TYPE": cpb.V3RoleCode_DEDICATED_SERVICE_DELIVERY_LOCATION_ROLE_TYPE, // Enum 240
	"DELEGATEE":                         cpb.V3RoleCode_DELEGATEE,                         // Enum 9
	"DELEGATOR":                         cpb.V3RoleCode_DELEGATOR,                         // Enum 10
	"DEP":                               cpb.V3RoleCode_DEP,                               // Enum 391
	"DEPEN":                             cpb.V3RoleCode_DEPEN,                             // Enum 392
	"DEPENDENT_COVERED_PARTY_ROLE_TYPE": cpb.V3RoleCode_DEPENDENT_COVERED_PARTY_ROLE_TYPE, // Enum 223
	"DERM":                              cpb.V3RoleCode_DERM,                              // Enum 303
	"DIFFABL":                           cpb.V3RoleCode_DIFFABL,                           // Enum 225
	"DOMPART":                           cpb.V3RoleCode_DOMPART,                           // Enum 192
	"DOWNGRDER":                         cpb.V3RoleCode_DOWNGRDER,                         // Enum 11
	"DPOWATT":                           cpb.V3RoleCode_DPOWATT,                           // Enum 30
	"DR":                                cpb.V3RoleCode_DR,                                // Enum 70
	"DRIVCLASSIFIER":                    cpb.V3RoleCode_DRIVCLASSIFIER,                    // Enum 12
	"DU":                                cpb.V3RoleCode_DU,                                // Enum 71
	"DX":                                cpb.V3RoleCode_DX,                                // Enum 242
	"E":                                 cpb.V3RoleCode_E,                                 // Enum 374
	"ECHO":                              cpb.V3RoleCode_ECHO,                              // Enum 245
	"ECON":                              cpb.V3RoleCode_ECON,                              // Enum 57
	"ENDO":                              cpb.V3RoleCode_ENDO,                              // Enum 304
	"ENDOS":                             cpb.V3RoleCode_ENDOS,                             // Enum 247
	"ENROLBKR":                          cpb.V3RoleCode_ENROLBKR,                          // Enum 22
	"ENT":                               cpb.V3RoleCode_ENT,                               // Enum 306
	"EPIL":                              cpb.V3RoleCode_EPIL,                              // Enum 269
	"ER":                                cpb.V3RoleCode_ER,                                // Enum 270
	"ERL":                               cpb.V3RoleCode_ERL,                               // Enum 237
	"ETU":                               cpb.V3RoleCode_ETU,                               // Enum 271
	"EXCEST":                            cpb.V3RoleCode_EXCEST,                            // Enum 26
	"EXT":                               cpb.V3RoleCode_EXT,                               // Enum 108
	"F":                                 cpb.V3RoleCode_F,                                 // Enum 375
	"FAMDEP":                            cpb.V3RoleCode_FAMDEP,                            // Enum 203
	"FAMMEMB":                           cpb.V3RoleCode_FAMMEMB,                           // Enum 92
	"FI":                                cpb.V3RoleCode_FI,                                // Enum 72
	"FM":                                cpb.V3RoleCode_FM,                                // Enum 393
	"FMC":                               cpb.V3RoleCode_FMC,                               // Enum 307
	"FMRSPS":                            cpb.V3RoleCode_FMRSPS,                            // Enum 193
	"FRND":                              cpb.V3RoleCode_FRND,                              // Enum 197
	"FSTUD":                             cpb.V3RoleCode_FSTUD,                             // Enum 209
	"FTH":                               cpb.V3RoleCode_FTH,                               // Enum 156
	"FTHFOST":                           cpb.V3RoleCode_FTHFOST,                           // Enum 157
	"FTHINLAW":                          cpb.V3RoleCode_FTHINLAW,                          // Enum 141
	"FTWIN":                             cpb.V3RoleCode_FTWIN,                             // Enum 186
	"FTWINBRO":                          cpb.V3RoleCode_FTWINBRO,                          // Enum 175
	"FTWINSIS":                          cpb.V3RoleCode_FTWINSIS,                          // Enum 183
	"FULLINS":                           cpb.V3RoleCode_FULLINS,                           // Enum 19
	"G":                                 cpb.V3RoleCode_G,                                 // Enum 369
	"GACH":                              cpb.V3RoleCode_GACH,                              // Enum 253
	"GCHILD":                            cpb.V3RoleCode_GCHILD,                            // Enum 212
	"GD":                                cpb.V3RoleCode_GD,                                // Enum 84
	"GDF":                               cpb.V3RoleCode_GDF,                               // Enum 85
	"GDS":                               cpb.V3RoleCode_GDS,                               // Enum 86
	"GDSF":                              cpb.V3RoleCode_GDSF,                              // Enum 87
	"GESTM":                             cpb.V3RoleCode_GESTM,                             // Enum 162
	"GGRFTH":                            cpb.V3RoleCode_GGRFTH,                            // Enum 116
	"GGRMTH":                            cpb.V3RoleCode_GGRMTH,                            // Enum 119
	"GGRPRN":                            cpb.V3RoleCode_GGRPRN,                            // Enum 115
	"GI":                                cpb.V3RoleCode_GI,                                // Enum 308
	"GIDX":                              cpb.V3RoleCode_GIDX,                              // Enum 246
	"GIM":                               cpb.V3RoleCode_GIM,                               // Enum 310
	"GPARNT":                            cpb.V3RoleCode_GPARNT,                            // Enum 213
	"GRANTEE":                           cpb.V3RoleCode_GRANTEE,                           // Enum 13
	"GRANTOR":                           cpb.V3RoleCode_GRANTOR,                           // Enum 14
	"GRFTH":                             cpb.V3RoleCode_GRFTH,                             // Enum 128
	"GRMTH":                             cpb.V3RoleCode_GRMTH,                             // Enum 131
	"GRNDCHILD":                         cpb.V3RoleCode_GRNDCHILD,                         // Enum 124
	"GRNDDAU":                           cpb.V3RoleCode_GRNDDAU,                           // Enum 125
	"GRNDSON":                           cpb.V3RoleCode_GRNDSON,                           // Enum 126
	"GRPRN":                             cpb.V3RoleCode_GRPRN,                             // Enum 127
	"GT":                                cpb.V3RoleCode_GT,                                // Enum 381
	"GUADLTM":                           cpb.V3RoleCode_GUADLTM,                           // Enum 27
	"GUARD":                             cpb.V3RoleCode_GUARD,                             // Enum 28
	"GYN":                               cpb.V3RoleCode_GYN,                               // Enum 311
	"HAND":                              cpb.V3RoleCode_HAND,                              // Enum 331
	"HANDIC":                            cpb.V3RoleCode_HANDIC,                            // Enum 204
	"HBRO":                              cpb.V3RoleCode_HBRO,                              // Enum 172
	"HD":                                cpb.V3RoleCode_HD,                                // Enum 272
	"HEM":                               cpb.V3RoleCode_HEM,                               // Enum 312
	"HLAB":                              cpb.V3RoleCode_HLAB,                              // Enum 273
	"HOME_HEALTH":                       cpb.V3RoleCode_HOME_HEALTH,                       // Enum 382
	"HOSP":                              cpb.V3RoleCode_HOSP,                              // Enum 251
	"HPOWATT":                           cpb.V3RoleCode_HPOWATT,                           // Enum 31
	"HRAD":                              cpb.V3RoleCode_HRAD,                              // Enum 276
	"HSIB":                              cpb.V3RoleCode_HSIB,                              // Enum 178
	"HSIS":                              cpb.V3RoleCode_HSIS,                              // Enum 179
	"HTN":                               cpb.V3RoleCode_HTN,                               // Enum 314
	"HU":                                cpb.V3RoleCode_HU,                                // Enum 265
	"HUSB":                              cpb.V3RoleCode_HUSB,                              // Enum 195
	"HUSCS":                             cpb.V3RoleCode_HUSCS,                             // Enum 277
	"ICU":                               cpb.V3RoleCode_ICU,                               // Enum 278
	"IDENTIFIED_ENTITY_TYPE":            cpb.V3RoleCode_IDENTIFIED_ENTITY_TYPE,            // Enum 59
	"IEC":                               cpb.V3RoleCode_IEC,                               // Enum 315
	"INCIDENTAL_SERVICE_DELIVERY_LOCATION_ROLE_TYPE": cpb.V3RoleCode_INCIDENTAL_SERVICE_DELIVERY_LOCATION_ROLE_TYPE, // Enum 359
	"INDIG":                                cpb.V3RoleCode_INDIG,                                // Enum 230
	"INDIV":                                cpb.V3RoleCode_INDIV,                                // Enum 394
	"INDIVIDUAL_INSURED_PARTY_ROLE_TYPE":   cpb.V3RoleCode_INDIVIDUAL_INSURED_PARTY_ROLE_TYPE,   // Enum 227
	"INFD":                                 cpb.V3RoleCode_INFD,                                 // Enum 316
	"INJ":                                  cpb.V3RoleCode_INJ,                                  // Enum 205
	"INJWKR":                               cpb.V3RoleCode_INJWKR,                               // Enum 222
	"INLAB":                                cpb.V3RoleCode_INLAB,                                // Enum 274
	"INLAW":                                cpb.V3RoleCode_INLAW,                                // Enum 136
	"INPHARM":                              cpb.V3RoleCode_INPHARM,                              // Enum 281
	"INTPRTER":                             cpb.V3RoleCode_INTPRTER,                             // Enum 15
	"INV":                                  cpb.V3RoleCode_INV,                                  // Enum 318
	"ITWIN":                                cpb.V3RoleCode_ITWIN,                                // Enum 187
	"ITWINBRO":                             cpb.V3RoleCode_ITWINBRO,                             // Enum 176
	"ITWINSIS":                             cpb.V3RoleCode_ITWINSIS,                             // Enum 184
	"JURID":                                cpb.V3RoleCode_JURID,                                // Enum 62
	"L":                                    cpb.V3RoleCode_L,                                    // Enum 370
	"LABORATORY":                           cpb.V3RoleCode_LABORATORY,                           // Enum 383
	"LIVING_SUBJECT_PRODUCTION_CLASS":      cpb.V3RoleCode_LIVING_SUBJECT_PRODUCTION_CLASS,      // Enum 64
	"LOCATION_IDENTIFIED_ENTITY_ROLE_CODE": cpb.V3RoleCode_LOCATION_IDENTIFIED_ENTITY_ROLE_CODE, // Enum 60
	"LOCHFID":                              cpb.V3RoleCode_LOCHFID,                              // Enum 63
	"LY":                                   cpb.V3RoleCode_LY,                                   // Enum 73
	"LYMPH":                                cpb.V3RoleCode_LYMPH,                                // Enum 319
	"MAUNT":                                cpb.V3RoleCode_MAUNT,                                // Enum 110
	"MBL":                                  cpb.V3RoleCode_MBL,                                  // Enum 282
	"MCOUSN":                               cpb.V3RoleCode_MCOUSN,                               // Enum 113
	"MEDICATION_GENERALIZATION_ROLE_TYPE":  cpb.V3RoleCode_MEDICATION_GENERALIZATION_ROLE_TYPE,  // Enum 82
	"MEMBER_ROLE_TYPE":                     cpb.V3RoleCode_MEMBER_ROLE_TYPE,                     // Enum 89
	"MGDSF":                                cpb.V3RoleCode_MGDSF,                                // Enum 88
	"MGEN":                                 cpb.V3RoleCode_MGEN,                                 // Enum 320
	"MGGRFTH":                              cpb.V3RoleCode_MGGRFTH,                              // Enum 117
	"MGGRMTH":                              cpb.V3RoleCode_MGGRMTH,                              // Enum 120
	"MGGRPRN":                              cpb.V3RoleCode_MGGRPRN,                              // Enum 122
	"MGRFTH":                               cpb.V3RoleCode_MGRFTH,                               // Enum 129
	"MGRMTH":                               cpb.V3RoleCode_MGRMTH,                               // Enum 132
	"MGRPRN":                               cpb.V3RoleCode_MGRPRN,                               // Enum 134
	"MHSP":                                 cpb.V3RoleCode_MHSP,                                 // Enum 254
	"MIL":                                  cpb.V3RoleCode_MIL,                                  // Enum 231
	"MOBL":                                 cpb.V3RoleCode_MOBL,                                 // Enum 356
	"MT":                                   cpb.V3RoleCode_MT,                                   // Enum 74
	"MTH":                                  cpb.V3RoleCode_MTH,                                  // Enum 161
	"MTHFOST":                              cpb.V3RoleCode_MTHFOST,                              // Enum 163
	"MTHINLAW":                             cpb.V3RoleCode_MTHINLAW,                             // Enum 142
	"MU":                                   cpb.V3RoleCode_MU,                                   // Enum 75
	"MUNCLE":                               cpb.V3RoleCode_MUNCLE,                               // Enum 150
	"NAMED":                                cpb.V3RoleCode_NAMED,                                // Enum 395
	"NAT":                                  cpb.V3RoleCode_NAT,                                  // Enum 214
	"NBOR":                                 cpb.V3RoleCode_NBOR,                                 // Enum 198
	"NBRO":                                 cpb.V3RoleCode_NBRO,                                 // Enum 173
	"NCCF":                                 cpb.V3RoleCode_NCCF,                                 // Enum 290
	"NCCS":                                 cpb.V3RoleCode_NCCS,                                 // Enum 283
	"NCHILD":                               cpb.V3RoleCode_NCHILD,                               // Enum 103
	"NEPH":                                 cpb.V3RoleCode_NEPH,                                 // Enum 321
	"NEPHEW":                               cpb.V3RoleCode_NEPHEW,                               // Enum 147
	"NEUR":                                 cpb.V3RoleCode_NEUR,                                 // Enum 323
	"NFTH":                                 cpb.V3RoleCode_NFTH,                                 // Enum 158
	"NFTHF":                                cpb.V3RoleCode_NFTHF,                                // Enum 159
	"NIECE":                                cpb.V3RoleCode_NIECE,                                // Enum 148
	"NIENE":                                cpb.V3RoleCode_NIENE,                                // Enum 215
	"NIENEPH":                              cpb.V3RoleCode_NIENEPH,                              // Enum 146
	"NMTH":                                 cpb.V3RoleCode_NMTH,                                 // Enum 164
	"NMTHF":                                cpb.V3RoleCode_NMTHF,                                // Enum 165
	"NOK":                                  cpb.V3RoleCode_NOK,                                  // Enum 58
	"NPRN":                                 cpb.V3RoleCode_NPRN,                                 // Enum 167
	"NS":                                   cpb.V3RoleCode_NS,                                   // Enum 284
	"NSIB":                                 cpb.V3RoleCode_NSIB,                                 // Enum 180
	"NSIS":                                 cpb.V3RoleCode_NSIS,                                 // Enum 181
	"O":                                    cpb.V3RoleCode_O,                                    // Enum 376
	"OB":                                   cpb.V3RoleCode_OB,                                   // Enum 324
	"OF":                                   cpb.V3RoleCode_OF,                                   // Enum 292
	"OMS":                                  cpb.V3RoleCode_OMS,                                  // Enum 325
	"ONCL":                                 cpb.V3RoleCode_ONCL,                                 // Enum 326
	"ONESELF":                              cpb.V3RoleCode_ONESELF,                              // Enum 199
	"OPH":                                  cpb.V3RoleCode_OPH,                                  // Enum 328
	"OPTC":                                 cpb.V3RoleCode_OPTC,                                 // Enum 329
	"ORG":                                  cpb.V3RoleCode_ORG,                                  // Enum 55
	"ORTHO":                                cpb.V3RoleCode_ORTHO,                                // Enum 330
	"OUTLAB":                               cpb.V3RoleCode_OUTLAB,                               // Enum 275
	"OUTPHARM":                             cpb.V3RoleCode_OUTPHARM,                             // Enum 285
	"P":                                    cpb.V3RoleCode_P,                                    // Enum 371
	"PAINCL":                               cpb.V3RoleCode_PAINCL,                               // Enum 332
	"PARNT":                                cpb.V3RoleCode_PARNT,                                // Enum 216
	"PATHOLOGIST":                          cpb.V3RoleCode_PATHOLOGIST,                          // Enum 384
	"PAUNT":                                cpb.V3RoleCode_PAUNT,                                // Enum 111
	"PAYOR":                                cpb.V3RoleCode_PAYOR,                                // Enum 56
	"PAYOR_ROLE_TYPE":                      cpb.V3RoleCode_PAYOR_ROLE_TYPE,                      // Enum 21
	"PC":                                   cpb.V3RoleCode_PC,                                   // Enum 333
	"PCOUSN":                               cpb.V3RoleCode_PCOUSN,                               // Enum 114
	"PEDC":                                 cpb.V3RoleCode_PEDC,                                 // Enum 334
	"PEDCARD":                              cpb.V3RoleCode_PEDCARD,                              // Enum 300
	"PEDE":                                 cpb.V3RoleCode_PEDE,                                 // Enum 305
	"PEDGI":                                cpb.V3RoleCode_PEDGI,                                // Enum 309
	"PEDHEM":                               cpb.V3RoleCode_PEDHEM,                               // Enum 313
	"PEDHO":                                cpb.V3RoleCode_PEDHO,                                // Enum 327
	"PEDICU":                               cpb.V3RoleCode_PEDICU,                               // Enum 279
	"PEDID":                                cpb.V3RoleCode_PEDID,                                // Enum 317
	"PEDNEPH":                              cpb.V3RoleCode_PEDNEPH,                              // Enum 322
	"PEDNICU":                              cpb.V3RoleCode_PEDNICU,                              // Enum 280
	"PEDRHEUM":                             cpb.V3RoleCode_PEDRHEUM,                             // Enum 335
	"PEDU":                                 cpb.V3RoleCode_PEDU,                                 // Enum 286
	"PERSONAL_RELATIONSHIP_ROLE_TYPE":      cpb.V3RoleCode_PERSONAL_RELATIONSHIP_ROLE_TYPE,      // Enum 91
	"PGGRFTH":                              cpb.V3RoleCode_PGGRFTH,                              // Enum 118
	"PGGRMTH":                              cpb.V3RoleCode_PGGRMTH,                              // Enum 121
	"PGGRPRN":                              cpb.V3RoleCode_PGGRPRN,                              // Enum 123
	"PGRFTH":                               cpb.V3RoleCode_PGRFTH,                               // Enum 130
	"PGRMTH":                               cpb.V3RoleCode_PGRMTH,                               // Enum 133
	"PGRPRN":                               cpb.V3RoleCode_PGRPRN,                               // Enum 135
	"PH":                                   cpb.V3RoleCode_PH,                                   // Enum 385
	"PHARM":                                cpb.V3RoleCode_PHARM,                                // Enum 358
	"PHLEBOTOMIST":                         cpb.V3RoleCode_PHLEBOTOMIST,                         // Enum 386
	"PHU":                                  cpb.V3RoleCode_PHU,                                  // Enum 287
	"PL":                                   cpb.V3RoleCode_PL,                                   // Enum 76
	"PLS":                                  cpb.V3RoleCode_PLS,                                  // Enum 346
	"POD":                                  cpb.V3RoleCode_POD,                                  // Enum 336
	"POLICY_OR_PROGRAM_COVERAGE_ROLE_TYPE": cpb.V3RoleCode_POLICY_OR_PROGRAM_COVERAGE_ROLE_TYPE, // Enum 201
	"POWATT":                               cpb.V3RoleCode_POWATT,                               // Enum 29
	"PRC":                                  cpb.V3RoleCode_PRC,                                  // Enum 352
	"PREV":                                 cpb.V3RoleCode_PREV,                                 // Enum 337
	"PRN":                                  cpb.V3RoleCode_PRN,                                  // Enum 152
	"PRNFOST":                              cpb.V3RoleCode_PRNFOST,                              // Enum 168
	"PRNINLAW":                             cpb.V3RoleCode_PRNINLAW,                             // Enum 140
	"PROCTO":                               cpb.V3RoleCode_PROCTO,                               // Enum 338
	"PROFF":                                cpb.V3RoleCode_PROFF,                                // Enum 339
	"PROG":                                 cpb.V3RoleCode_PROG,                                 // Enum 387
	"PROGRAM_ELIGIBLE_PARTY_ROLE_TYPE":     cpb.V3RoleCode_PROGRAM_ELIGIBLE_PARTY_ROLE_TYPE,     // Enum 229
	"PROS":                                 cpb.V3RoleCode_PROS,                                 // Enum 340
	"PSI":                                  cpb.V3RoleCode_PSI,                                  // Enum 341
	"PSTUD":                                cpb.V3RoleCode_PSTUD,                                // Enum 210
	"PSY":                                  cpb.V3RoleCode_PSY,                                  // Enum 342
	"PSYCHCF":                              cpb.V3RoleCode_PSYCHCF,                              // Enum 396
	"PSYCHF":                               cpb.V3RoleCode_PSYCHF,                               // Enum 255
	"PT":                                   cpb.V3RoleCode_PT,                                   // Enum 388
	"PTRES":                                cpb.V3RoleCode_PTRES,                                // Enum 363
	"PUNCLE":                               cpb.V3RoleCode_PUNCLE,                               // Enum 151
	"Q":                                    cpb.V3RoleCode_Q,                                    // Enum 372
	"R":                                    cpb.V3RoleCode_R,                                    // Enum 378
	"RADDX":                                cpb.V3RoleCode_RADDX,                                // Enum 248
	"RADO":                                 cpb.V3RoleCode_RADO,                                 // Enum 249
	"RC":                                   cpb.V3RoleCode_RC,                                   // Enum 77
	"RESEARCH_SUBJECT_ROLE_BASIS":          cpb.V3RoleCode_RESEARCH_SUBJECT_ROLE_BASIS,          // Enum 236
	"RESPRSN":                              cpb.V3RoleCode_RESPRSN,                              // Enum 25
	"RETIREE":                              cpb.V3RoleCode_RETIREE,                              // Enum 228
	"RETMIL":                               cpb.V3RoleCode_RETMIL,                               // Enum 233
	"REVIEWER":                             cpb.V3RoleCode_REVIEWER,                             // Enum 16
	"RH":                                   cpb.V3RoleCode_RH,                                   // Enum 256
	"RHAT":                                 cpb.V3RoleCode_RHAT,                                 // Enum 257
	"RHEUM":                                cpb.V3RoleCode_RHEUM,                                // Enum 343
	"RHII":                                 cpb.V3RoleCode_RHII,                                 // Enum 258
	"RHMAD":                                cpb.V3RoleCode_RHMAD,                                // Enum 259
	"RHPI":                                 cpb.V3RoleCode_RHPI,                                 // Enum 260
	"RHPIH":                                cpb.V3RoleCode_RHPIH,                                // Enum 261
	"RHPIMS":                               cpb.V3RoleCode_RHPIMS,                               // Enum 262
	"RHPIVS":                               cpb.V3RoleCode_RHPIVS,                               // Enum 263
	"RHU":                                  cpb.V3RoleCode_RHU,                                  // Enum 288
	"RHYAD":                                cpb.V3RoleCode_RHYAD,                                // Enum 264
	"RNEU":                                 cpb.V3RoleCode_RNEU,                                 // Enum 250
	"ROOM":                                 cpb.V3RoleCode_ROOM,                                 // Enum 200
	"RTF":                                  cpb.V3RoleCode_RTF,                                  // Enum 351
	"SCHOOL":                               cpb.V3RoleCode_SCHOOL,                               // Enum 364
	"SCN":                                  cpb.V3RoleCode_SCN,                                  // Enum 238
	"SEE":                                  cpb.V3RoleCode_SEE,                                  // Enum 40
	"SELF":                                 cpb.V3RoleCode_SELF,                                 // Enum 206
	"SELFINS":                              cpb.V3RoleCode_SELFINS,                              // Enum 20
	"SERVICE_DELIVERY_LOCATION_ROLE_TYPE":  cpb.V3RoleCode_SERVICE_DELIVERY_LOCATION_ROLE_TYPE,  // Enum 239
	"SH":                                   cpb.V3RoleCode_SH,                                   // Enum 78
	"SIB":                                  cpb.V3RoleCode_SIB,                                  // Enum 170
	"SIBINLAW":                             cpb.V3RoleCode_SIBINLAW,                             // Enum 143
	"SIGOTHR":                              cpb.V3RoleCode_SIGOTHR,                              // Enum 191
	"SIS":                                  cpb.V3RoleCode_SIS,                                  // Enum 188
	"SISINLAW":                             cpb.V3RoleCode_SISINLAW,                             // Enum 145
	"SLEEP":                                cpb.V3RoleCode_SLEEP,                                // Enum 289
	"SNF":                                  cpb.V3RoleCode_SNF,                                  // Enum 291
	"SNIFF":                                cpb.V3RoleCode_SNIFF,                                // Enum 41
	"SON":                                  cpb.V3RoleCode_SON,                                  // Enum 104
	"SONADOPT":                             cpb.V3RoleCode_SONADOPT,                             // Enum 96
	"SONC":                                 cpb.V3RoleCode_SONC,                                 // Enum 105
	"SONFOST":                              cpb.V3RoleCode_SONFOST,                              // Enum 99
	"SONINLAW":                             cpb.V3RoleCode_SONINLAW,                             // Enum 139
	"SPECIMEN_ROLE_TYPE":                   cpb.V3RoleCode_SPECIMEN_ROLE_TYPE,                   // Enum 367
	"SPMED":                                cpb.V3RoleCode_SPMED,                                // Enum 344
	"SPON":                                 cpb.V3RoleCode_SPON,                                 // Enum 207
	"SPOWATT":                              cpb.V3RoleCode_SPOWATT,                              // Enum 32
	"SPS":                                  cpb.V3RoleCode_SPS,                                  // Enum 194
	"SPSE":                                 cpb.V3RoleCode_SPSE,                                 // Enum 217
	"STEP":                                 cpb.V3RoleCode_STEP,                                 // Enum 218
	"STPBRO":                               cpb.V3RoleCode_STPBRO,                               // Enum 177
	"STPCHLD":                              cpb.V3RoleCode_STPCHLD,                              // Enum 107
	"STPDAU":                               cpb.V3RoleCode_STPDAU,                               // Enum 102
	"STPFTH":                               cpb.V3RoleCode_STPFTH,                               // Enum 160
	"STPMTH":                               cpb.V3RoleCode_STPMTH,                               // Enum 166
	"STPPRN":                               cpb.V3RoleCode_STPPRN,                               // Enum 169
	"STPSIB":                               cpb.V3RoleCode_STPSIB,                               // Enum 190
	"STPSIS":                               cpb.V3RoleCode_STPSIS,                               // Enum 189
	"STPSON":                               cpb.V3RoleCode_STPSON,                               // Enum 106
	"STUD":                                 cpb.V3RoleCode_STUD,                                 // Enum 208
	"SU":                                   cpb.V3RoleCode_SU,                                   // Enum 345
	"SUBJECT":                              cpb.V3RoleCode_SUBJECT,                              // Enum 389
	"SUBSCR":                               cpb.V3RoleCode_SUBSCR,                               // Enum 397
	"SUBSCRIBER_COVERED_PARTY_ROLE_TYPE":   cpb.V3RoleCode_SUBSCRIBER_COVERED_PARTY_ROLE_TYPE,   // Enum 235
	"SURF":                                 cpb.V3RoleCode_SURF,                                 // Enum 353
	"THIRD_PARTY":                          cpb.V3RoleCode_THIRD_PARTY,                          // Enum 390
	"TPA":                                  cpb.V3RoleCode_TPA,                                  // Enum 23
	"TR":                                   cpb.V3RoleCode_TR,                                   // Enum 348
	"TRAVEL":                               cpb.V3RoleCode_TRAVEL,                               // Enum 349
	"TRB":                                  cpb.V3RoleCode_TRB,                                  // Enum 90
	"TWIN":                                 cpb.V3RoleCode_TWIN,                                 // Enum 185
	"TWINBRO":                              cpb.V3RoleCode_TWINBRO,                              // Enum 174
	"TWINSIS":                              cpb.V3RoleCode_TWINSIS,                              // Enum 182
	"UMO":                                  cpb.V3RoleCode_UMO,                                  // Enum 24
	"UNCLE":                                cpb.V3RoleCode_UNCLE,                                // Enum 149
	"UPC":                                  cpb.V3RoleCode_UPC,                                  // Enum 365
	"URO":                                  cpb.V3RoleCode_URO,                                  // Enum 347
	"V":                                    cpb.V3RoleCode_V,                                    // Enum 377
	"VALIDATOR":                            cpb.V3RoleCode_VALIDATOR,                            // Enum 17
	"VET":                                  cpb.V3RoleCode_VET,                                  // Enum 234
	"VL":                                   cpb.V3RoleCode_VL,                                   // Enum 79
	"WARD":                                 cpb.V3RoleCode_WARD,                                 // Enum 226
	"WIFE":                                 cpb.V3RoleCode_WIFE,                                 // Enum 196
	"WL":                                   cpb.V3RoleCode_WL,                                   // Enum 80
	"WND":                                  cpb.V3RoleCode_WND,                                  // Enum 350
	"WO":                                   cpb.V3RoleCode_WO,                                   // Enum 81
	"WORK":                                 cpb.V3RoleCode_WORK,                                 // Enum 366
}

// DefaultV3TimingEventCodeMap maps from string to cpb.V3TimingEventCode_Value.
var DefaultV3TimingEventCodeMap = map[string]cpb.V3TimingEventCode_Value{
	"INVALID_UNINITIALIZED": cpb.V3TimingEventCode_INVALID_UNINITIALIZED, // Enum 0
	"AC":                    cpb.V3TimingEventCode_AC,                    // Enum 1
	"ACD":                   cpb.V3TimingEventCode_ACD,                   // Enum 2
	"ACM":                   cpb.V3TimingEventCode_ACM,                   // Enum 3
	"ACV":                   cpb.V3TimingEventCode_ACV,                   // Enum 4
	"C":                     cpb.V3TimingEventCode_C,                     // Enum 5
	"CD":                    cpb.V3TimingEventCode_CD,                    // Enum 6
	"CM":                    cpb.V3TimingEventCode_CM,                    // Enum 7
	"CV":                    cpb.V3TimingEventCode_CV,                    // Enum 8
	"HS":                    cpb.V3TimingEventCode_HS,                    // Enum 9
	"IC":                    cpb.V3TimingEventCode_IC,                    // Enum 10
	"ICD":                   cpb.V3TimingEventCode_ICD,                   // Enum 11
	"ICM":                   cpb.V3TimingEventCode_ICM,                   // Enum 12
	"ICV":                   cpb.V3TimingEventCode_ICV,                   // Enum 13
	"PC":                    cpb.V3TimingEventCode_PC,                    // Enum 14
	"PCD":                   cpb.V3TimingEventCode_PCD,                   // Enum 15
	"PCM":                   cpb.V3TimingEventCode_PCM,                   // Enum 16
	"PCV":                   cpb.V3TimingEventCode_PCV,                   // Enum 17
	"WAKE":                  cpb.V3TimingEventCode_WAKE,                  // Enum 18
}

// DefaultVisionBaseCodeMap maps from string to cpb.VisionBaseCode_Value.
var DefaultVisionBaseCodeMap = map[string]cpb.VisionBaseCode_Value{
	"INVALID_UNINITIALIZED": cpb.VisionBaseCode_INVALID_UNINITIALIZED, // Enum 0
	"DOWN":                  cpb.VisionBaseCode_DOWN,                  // Enum 2
	"IN":                    cpb.VisionBaseCode_IN,                    // Enum 3
	"OUT":                   cpb.VisionBaseCode_OUT,                   // Enum 4
	"UP":                    cpb.VisionBaseCode_UP,                    // Enum 1
}

// DefaultVisionEyesCodeMap maps from string to cpb.VisionEyesCode_Value.
var DefaultVisionEyesCodeMap = map[string]cpb.VisionEyesCode_Value{
	"INVALID_UNINITIALIZED": cpb.VisionEyesCode_INVALID_UNINITIALIZED, // Enum 0
	"LEFT":                  cpb.VisionEyesCode_LEFT,                  // Enum 2
	"RIGHT":                 cpb.VisionEyesCode_RIGHT,                 // Enum 1
}

// DefaultXPathUsageTypeCodeMap maps from string to cpb.XPathUsageTypeCode_Value.
var DefaultXPathUsageTypeCodeMap = map[string]cpb.XPathUsageTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.XPathUsageTypeCode_INVALID_UNINITIALIZED, // Enum 0
	"DISTANCE":              cpb.XPathUsageTypeCode_DISTANCE,              // Enum 4
	"NEARBY":                cpb.XPathUsageTypeCode_NEARBY,                // Enum 3
	"NORMAL":                cpb.XPathUsageTypeCode_NORMAL,                // Enum 1
	"OTHER":                 cpb.XPathUsageTypeCode_OTHER,                 // Enum 5
	"PHONETIC":              cpb.XPathUsageTypeCode_PHONETIC,              // Enum 2
}
