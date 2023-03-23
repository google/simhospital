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

package hl7tofhir

import (
	cpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/codes_go_proto"
)

// DefaultAbstractTypeCodeMap maps from string to cpb.AbstractTypeCode_Value.
var DefaultAbstractTypeCodeMap = map[string]cpb.AbstractTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.AbstractTypeCode_INVALID_UNINITIALIZED,
	"TYPE":                  cpb.AbstractTypeCode_TYPE,
	"ANY":                   cpb.AbstractTypeCode_ANY,
}

// DefaultAccountStatusCodeMap maps from string to cpb.AccountStatusCode_Value.
var DefaultAccountStatusCodeMap = map[string]cpb.AccountStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.AccountStatusCode_INVALID_UNINITIALIZED,
	"ACTIVE":                cpb.AccountStatusCode_ACTIVE,
	"INACTIVE":              cpb.AccountStatusCode_INACTIVE,
	"ENTERED_IN_ERROR":      cpb.AccountStatusCode_ENTERED_IN_ERROR,
	"ON_HOLD":               cpb.AccountStatusCode_ON_HOLD,
	"UNKNOWN":               cpb.AccountStatusCode_UNKNOWN,
}

// DefaultActionCardinalityBehaviorCodeMap maps from string to cpb.ActionCardinalityBehaviorCode_Value.
var DefaultActionCardinalityBehaviorCodeMap = map[string]cpb.ActionCardinalityBehaviorCode_Value{
	"INVALID_UNINITIALIZED": cpb.ActionCardinalityBehaviorCode_INVALID_UNINITIALIZED,
	"SINGLE":                cpb.ActionCardinalityBehaviorCode_SINGLE,
	"MULTIPLE":              cpb.ActionCardinalityBehaviorCode_MULTIPLE,
}

// DefaultActionConditionKindCodeMap maps from string to cpb.ActionConditionKindCode_Value.
var DefaultActionConditionKindCodeMap = map[string]cpb.ActionConditionKindCode_Value{
	"INVALID_UNINITIALIZED": cpb.ActionConditionKindCode_INVALID_UNINITIALIZED,
	"APPLICABILITY":         cpb.ActionConditionKindCode_APPLICABILITY,
	"START":                 cpb.ActionConditionKindCode_START,
	"STOP":                  cpb.ActionConditionKindCode_STOP,
}

// DefaultActionGroupingBehaviorCodeMap maps from string to cpb.ActionGroupingBehaviorCode_Value.
var DefaultActionGroupingBehaviorCodeMap = map[string]cpb.ActionGroupingBehaviorCode_Value{
	"INVALID_UNINITIALIZED": cpb.ActionGroupingBehaviorCode_INVALID_UNINITIALIZED,
	"VISUAL_GROUP":          cpb.ActionGroupingBehaviorCode_VISUAL_GROUP,
	"LOGICAL_GROUP":         cpb.ActionGroupingBehaviorCode_LOGICAL_GROUP,
	"SENTENCE_GROUP":        cpb.ActionGroupingBehaviorCode_SENTENCE_GROUP,
}

// DefaultActionParticipantTypeCodeMap maps from string to cpb.ActionParticipantTypeCode_Value.
var DefaultActionParticipantTypeCodeMap = map[string]cpb.ActionParticipantTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ActionParticipantTypeCode_INVALID_UNINITIALIZED,
	"PATIENT":               cpb.ActionParticipantTypeCode_PATIENT,
	"PRACTITIONER":          cpb.ActionParticipantTypeCode_PRACTITIONER,
	"RELATED_PERSON":        cpb.ActionParticipantTypeCode_RELATED_PERSON,
	"DEVICE":                cpb.ActionParticipantTypeCode_DEVICE,
}

// DefaultActionPrecheckBehaviorCodeMap maps from string to cpb.ActionPrecheckBehaviorCode_Value.
var DefaultActionPrecheckBehaviorCodeMap = map[string]cpb.ActionPrecheckBehaviorCode_Value{
	"INVALID_UNINITIALIZED": cpb.ActionPrecheckBehaviorCode_INVALID_UNINITIALIZED,
	"YES":                   cpb.ActionPrecheckBehaviorCode_YES,
	"NO":                    cpb.ActionPrecheckBehaviorCode_NO,
}

// DefaultActionRelationshipTypeCodeMap maps from string to cpb.ActionRelationshipTypeCode_Value.
var DefaultActionRelationshipTypeCodeMap = map[string]cpb.ActionRelationshipTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ActionRelationshipTypeCode_INVALID_UNINITIALIZED,
	"BEFORE_START":          cpb.ActionRelationshipTypeCode_BEFORE_START,
	"BEFORE":                cpb.ActionRelationshipTypeCode_BEFORE,
	"BEFORE_END":            cpb.ActionRelationshipTypeCode_BEFORE_END,
	"CONCURRENT_WITH_START": cpb.ActionRelationshipTypeCode_CONCURRENT_WITH_START,
	"CONCURRENT":            cpb.ActionRelationshipTypeCode_CONCURRENT,
	"CONCURRENT_WITH_END":   cpb.ActionRelationshipTypeCode_CONCURRENT_WITH_END,
	"AFTER_START":           cpb.ActionRelationshipTypeCode_AFTER_START,
	"AFTER":                 cpb.ActionRelationshipTypeCode_AFTER,
	"AFTER_END":             cpb.ActionRelationshipTypeCode_AFTER_END,
}

// DefaultActionRequiredBehaviorCodeMap maps from string to cpb.ActionRequiredBehaviorCode_Value.
var DefaultActionRequiredBehaviorCodeMap = map[string]cpb.ActionRequiredBehaviorCode_Value{
	"INVALID_UNINITIALIZED":  cpb.ActionRequiredBehaviorCode_INVALID_UNINITIALIZED,
	"MUST":                   cpb.ActionRequiredBehaviorCode_MUST,
	"COULD":                  cpb.ActionRequiredBehaviorCode_COULD,
	"MUST_UNLESS_DOCUMENTED": cpb.ActionRequiredBehaviorCode_MUST_UNLESS_DOCUMENTED,
}

// DefaultActionSelectionBehaviorCodeMap maps from string to cpb.ActionSelectionBehaviorCode_Value.
var DefaultActionSelectionBehaviorCodeMap = map[string]cpb.ActionSelectionBehaviorCode_Value{
	"INVALID_UNINITIALIZED": cpb.ActionSelectionBehaviorCode_INVALID_UNINITIALIZED,
	"ANY":                   cpb.ActionSelectionBehaviorCode_ANY,
	"ALL":                   cpb.ActionSelectionBehaviorCode_ALL,
	"ALL_OR_NONE":           cpb.ActionSelectionBehaviorCode_ALL_OR_NONE,
	"EXACTLY_ONE":           cpb.ActionSelectionBehaviorCode_EXACTLY_ONE,
	"AT_MOST_ONE":           cpb.ActionSelectionBehaviorCode_AT_MOST_ONE,
	"ONE_OR_MORE":           cpb.ActionSelectionBehaviorCode_ONE_OR_MORE,
}

// DefaultAddressTypeCodeMap maps from string to cpb.AddressTypeCode_Value.
var DefaultAddressTypeCodeMap = map[string]cpb.AddressTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.AddressTypeCode_INVALID_UNINITIALIZED,
	"POSTAL":                cpb.AddressTypeCode_POSTAL,
	"PHYSICAL":              cpb.AddressTypeCode_PHYSICAL,
	"BOTH":                  cpb.AddressTypeCode_BOTH,
}

// DefaultAddressUseCodeMap maps from string to cpb.AddressUseCode_Value.
var DefaultAddressUseCodeMap = map[string]cpb.AddressUseCode_Value{
	"INVALID_UNINITIALIZED": cpb.AddressUseCode_INVALID_UNINITIALIZED,
	"HOME":                  cpb.AddressUseCode_HOME,
	"WORK":                  cpb.AddressUseCode_WORK,
	"TEMP":                  cpb.AddressUseCode_TEMP,
	"OLD":                   cpb.AddressUseCode_OLD,
	"BILLING":               cpb.AddressUseCode_BILLING,
}

// DefaultAdministrativeGenderCodeMap maps from string to cpb.AdministrativeGenderCode_Value.
var DefaultAdministrativeGenderCodeMap = map[string]cpb.AdministrativeGenderCode_Value{
	"INVALID_UNINITIALIZED": cpb.AdministrativeGenderCode_INVALID_UNINITIALIZED,
	"MALE":                  cpb.AdministrativeGenderCode_MALE,
	"FEMALE":                cpb.AdministrativeGenderCode_FEMALE,
	"OTHER":                 cpb.AdministrativeGenderCode_OTHER,
	"UNKNOWN":               cpb.AdministrativeGenderCode_UNKNOWN,
}

// DefaultAdverseEventActualityCodeMap maps from string to cpb.AdverseEventActualityCode_Value.
var DefaultAdverseEventActualityCodeMap = map[string]cpb.AdverseEventActualityCode_Value{
	"INVALID_UNINITIALIZED": cpb.AdverseEventActualityCode_INVALID_UNINITIALIZED,
	"ACTUAL":                cpb.AdverseEventActualityCode_ACTUAL,
	"POTENTIAL":             cpb.AdverseEventActualityCode_POTENTIAL,
}

// DefaultAdverseEventOutcomeCodeMap maps from string to cpb.AdverseEventOutcomeCode_Value.
var DefaultAdverseEventOutcomeCodeMap = map[string]cpb.AdverseEventOutcomeCode_Value{
	"INVALID_UNINITIALIZED":  cpb.AdverseEventOutcomeCode_INVALID_UNINITIALIZED,
	"RESOLVED":               cpb.AdverseEventOutcomeCode_RESOLVED,
	"RECOVERING":             cpb.AdverseEventOutcomeCode_RECOVERING,
	"ONGOING":                cpb.AdverseEventOutcomeCode_ONGOING,
	"RESOLVED_WITH_SEQUELAE": cpb.AdverseEventOutcomeCode_RESOLVED_WITH_SEQUELAE,
	"FATAL":                  cpb.AdverseEventOutcomeCode_FATAL,
	"UNKNOWN":                cpb.AdverseEventOutcomeCode_UNKNOWN,
}

// DefaultAdverseEventSeverityCodeMap maps from string to cpb.AdverseEventSeverityCode_Value.
var DefaultAdverseEventSeverityCodeMap = map[string]cpb.AdverseEventSeverityCode_Value{
	"INVALID_UNINITIALIZED": cpb.AdverseEventSeverityCode_INVALID_UNINITIALIZED,
	"MILD":                  cpb.AdverseEventSeverityCode_MILD,
	"MODERATE":              cpb.AdverseEventSeverityCode_MODERATE,
	"SEVERE":                cpb.AdverseEventSeverityCode_SEVERE,
}

// DefaultAggregationModeCodeMap maps from string to cpb.AggregationModeCode_Value.
var DefaultAggregationModeCodeMap = map[string]cpb.AggregationModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.AggregationModeCode_INVALID_UNINITIALIZED,
	"CONTAINED":             cpb.AggregationModeCode_CONTAINED,
	"REFERENCED":            cpb.AggregationModeCode_REFERENCED,
	"BUNDLED":               cpb.AggregationModeCode_BUNDLED,
}

// DefaultAllergyIntoleranceCategoryCodeMap maps from string to cpb.AllergyIntoleranceCategoryCode_Value.
var DefaultAllergyIntoleranceCategoryCodeMap = map[string]cpb.AllergyIntoleranceCategoryCode_Value{
	"INVALID_UNINITIALIZED": cpb.AllergyIntoleranceCategoryCode_INVALID_UNINITIALIZED,
	"FOOD":                  cpb.AllergyIntoleranceCategoryCode_FOOD,
	"MEDICATION":            cpb.AllergyIntoleranceCategoryCode_MEDICATION,
	"ENVIRONMENT":           cpb.AllergyIntoleranceCategoryCode_ENVIRONMENT,
	"BIOLOGIC":              cpb.AllergyIntoleranceCategoryCode_BIOLOGIC,
}

// DefaultAllergyIntoleranceClinicalStatusCodeMap maps from string to cpb.AllergyIntoleranceClinicalStatusCode_Value.
var DefaultAllergyIntoleranceClinicalStatusCodeMap = map[string]cpb.AllergyIntoleranceClinicalStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.AllergyIntoleranceClinicalStatusCode_INVALID_UNINITIALIZED,
	"ACTIVE":                cpb.AllergyIntoleranceClinicalStatusCode_ACTIVE,
	"INACTIVE":              cpb.AllergyIntoleranceClinicalStatusCode_INACTIVE,
	"RESOLVED":              cpb.AllergyIntoleranceClinicalStatusCode_RESOLVED,
}

// DefaultAllergyIntoleranceCriticalityCodeMap maps from string to cpb.AllergyIntoleranceCriticalityCode_Value.
var DefaultAllergyIntoleranceCriticalityCodeMap = map[string]cpb.AllergyIntoleranceCriticalityCode_Value{
	"INVALID_UNINITIALIZED": cpb.AllergyIntoleranceCriticalityCode_INVALID_UNINITIALIZED,
	"LOW":                   cpb.AllergyIntoleranceCriticalityCode_LOW,
	"HIGH":                  cpb.AllergyIntoleranceCriticalityCode_HIGH,
	"UNABLE_TO_ASSESS":      cpb.AllergyIntoleranceCriticalityCode_UNABLE_TO_ASSESS,
}

// DefaultAllergyIntoleranceSeverityCodeMap maps from string to cpb.AllergyIntoleranceSeverityCode_Value.
var DefaultAllergyIntoleranceSeverityCodeMap = map[string]cpb.AllergyIntoleranceSeverityCode_Value{
	"INVALID_UNINITIALIZED": cpb.AllergyIntoleranceSeverityCode_INVALID_UNINITIALIZED,
	"MILD":                  cpb.AllergyIntoleranceSeverityCode_MILD,
	"MODERATE":              cpb.AllergyIntoleranceSeverityCode_MODERATE,
	"SEVERE":                cpb.AllergyIntoleranceSeverityCode_SEVERE,
}

// DefaultAllergyIntoleranceSubstanceExposureRiskCodeMap maps from string to cpb.AllergyIntoleranceSubstanceExposureRiskCode_Value.
var DefaultAllergyIntoleranceSubstanceExposureRiskCodeMap = map[string]cpb.AllergyIntoleranceSubstanceExposureRiskCode_Value{
	"INVALID_UNINITIALIZED":  cpb.AllergyIntoleranceSubstanceExposureRiskCode_INVALID_UNINITIALIZED,
	"KNOWN_REACTION_RISK":    cpb.AllergyIntoleranceSubstanceExposureRiskCode_KNOWN_REACTION_RISK,
	"NO_KNOWN_REACTION_RISK": cpb.AllergyIntoleranceSubstanceExposureRiskCode_NO_KNOWN_REACTION_RISK,
}

// DefaultAllergyIntoleranceTypeCodeMap maps from string to cpb.AllergyIntoleranceTypeCode_Value.
var DefaultAllergyIntoleranceTypeCodeMap = map[string]cpb.AllergyIntoleranceTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.AllergyIntoleranceTypeCode_INVALID_UNINITIALIZED,
	"ALLERGY":               cpb.AllergyIntoleranceTypeCode_ALLERGY,
	"INTOLERANCE":           cpb.AllergyIntoleranceTypeCode_INTOLERANCE,
}

// DefaultAllergyIntoleranceVerificationStatusCodeMap maps from string to cpb.AllergyIntoleranceVerificationStatusCode_Value.
var DefaultAllergyIntoleranceVerificationStatusCodeMap = map[string]cpb.AllergyIntoleranceVerificationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.AllergyIntoleranceVerificationStatusCode_INVALID_UNINITIALIZED,
	"UNCONFIRMED":           cpb.AllergyIntoleranceVerificationStatusCode_UNCONFIRMED,
	"CONFIRMED":             cpb.AllergyIntoleranceVerificationStatusCode_CONFIRMED,
	"REFUTED":               cpb.AllergyIntoleranceVerificationStatusCode_REFUTED,
	"ENTERED_IN_ERROR":      cpb.AllergyIntoleranceVerificationStatusCode_ENTERED_IN_ERROR,
}

// DefaultAppointmentStatusCodeMap maps from string to cpb.AppointmentStatusCode_Value.
var DefaultAppointmentStatusCodeMap = map[string]cpb.AppointmentStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.AppointmentStatusCode_INVALID_UNINITIALIZED,
	"PROPOSED":              cpb.AppointmentStatusCode_PROPOSED,
	"PENDING":               cpb.AppointmentStatusCode_PENDING,
	"BOOKED":                cpb.AppointmentStatusCode_BOOKED,
	"ARRIVED":               cpb.AppointmentStatusCode_ARRIVED,
	"FULFILLED":             cpb.AppointmentStatusCode_FULFILLED,
	"CANCELLED":             cpb.AppointmentStatusCode_CANCELLED,
	"NOSHOW":                cpb.AppointmentStatusCode_NOSHOW,
	"ENTERED_IN_ERROR":      cpb.AppointmentStatusCode_ENTERED_IN_ERROR,
	"CHECKED_IN":            cpb.AppointmentStatusCode_CHECKED_IN,
	"WAITLIST":              cpb.AppointmentStatusCode_WAITLIST,
}

// DefaultAssertionDirectionTypeCodeMap maps from string to cpb.AssertionDirectionTypeCode_Value.
var DefaultAssertionDirectionTypeCodeMap = map[string]cpb.AssertionDirectionTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.AssertionDirectionTypeCode_INVALID_UNINITIALIZED,
	"RESPONSE":              cpb.AssertionDirectionTypeCode_RESPONSE,
	"REQUEST":               cpb.AssertionDirectionTypeCode_REQUEST,
}

// DefaultAssertionOperatorTypeCodeMap maps from string to cpb.AssertionOperatorTypeCode_Value.
var DefaultAssertionOperatorTypeCodeMap = map[string]cpb.AssertionOperatorTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.AssertionOperatorTypeCode_INVALID_UNINITIALIZED,
	"EQUALS":                cpb.AssertionOperatorTypeCode_EQUALS,
	"NOT_EQUALS":            cpb.AssertionOperatorTypeCode_NOT_EQUALS,
	"IN":                    cpb.AssertionOperatorTypeCode_IN,
	"NOT_IN":                cpb.AssertionOperatorTypeCode_NOT_IN,
	"GREATER_THAN":          cpb.AssertionOperatorTypeCode_GREATER_THAN,
	"LESS_THAN":             cpb.AssertionOperatorTypeCode_LESS_THAN,
	"EMPTY":                 cpb.AssertionOperatorTypeCode_EMPTY,
	"NOT_EMPTY":             cpb.AssertionOperatorTypeCode_NOT_EMPTY,
	"CONTAINS":              cpb.AssertionOperatorTypeCode_CONTAINS,
	"NOT_CONTAINS":          cpb.AssertionOperatorTypeCode_NOT_CONTAINS,
	"EVAL":                  cpb.AssertionOperatorTypeCode_EVAL,
}

// DefaultAssertionResponseTypesCodeMap maps from string to cpb.AssertionResponseTypesCode_Value.
var DefaultAssertionResponseTypesCodeMap = map[string]cpb.AssertionResponseTypesCode_Value{
	"INVALID_UNINITIALIZED": cpb.AssertionResponseTypesCode_INVALID_UNINITIALIZED,
	"OKAY":                  cpb.AssertionResponseTypesCode_OKAY,
	"CREATED":               cpb.AssertionResponseTypesCode_CREATED,
	"NO_CONTENT":            cpb.AssertionResponseTypesCode_NO_CONTENT,
	"NOT_MODIFIED":          cpb.AssertionResponseTypesCode_NOT_MODIFIED,
	"BAD":                   cpb.AssertionResponseTypesCode_BAD,
	"FORBIDDEN":             cpb.AssertionResponseTypesCode_FORBIDDEN,
	"NOT_FOUND":             cpb.AssertionResponseTypesCode_NOT_FOUND,
	"METHOD_NOT_ALLOWED":    cpb.AssertionResponseTypesCode_METHOD_NOT_ALLOWED,
	"CONFLICT":              cpb.AssertionResponseTypesCode_CONFLICT,
	"GONE":                  cpb.AssertionResponseTypesCode_GONE,
	"PRECONDITION_FAILED":   cpb.AssertionResponseTypesCode_PRECONDITION_FAILED,
	"UNPROCESSABLE":         cpb.AssertionResponseTypesCode_UNPROCESSABLE,
}

// DefaultAuditEventActionCodeMap maps from string to cpb.AuditEventActionCode_Value.
var DefaultAuditEventActionCodeMap = map[string]cpb.AuditEventActionCode_Value{
	"INVALID_UNINITIALIZED": cpb.AuditEventActionCode_INVALID_UNINITIALIZED,
	"C":                     cpb.AuditEventActionCode_C,
	"R":                     cpb.AuditEventActionCode_R,
	"U":                     cpb.AuditEventActionCode_U,
	"D":                     cpb.AuditEventActionCode_D,
	"E":                     cpb.AuditEventActionCode_E,
}

// DefaultAuditEventAgentNetworkTypeCodeMap maps from string to cpb.AuditEventAgentNetworkTypeCode_Value.
var DefaultAuditEventAgentNetworkTypeCodeMap = map[string]cpb.AuditEventAgentNetworkTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.AuditEventAgentNetworkTypeCode_INVALID_UNINITIALIZED,
	"MACHINE_NAME":          cpb.AuditEventAgentNetworkTypeCode_MACHINE_NAME,
	"IP_ADDRESS":            cpb.AuditEventAgentNetworkTypeCode_IP_ADDRESS,
	"TELEPHONE_NUMBER":      cpb.AuditEventAgentNetworkTypeCode_TELEPHONE_NUMBER,
	"EMAIL_ADDRESS":         cpb.AuditEventAgentNetworkTypeCode_EMAIL_ADDRESS,
	"URI":                   cpb.AuditEventAgentNetworkTypeCode_URI,
}

// DefaultAuditEventOutcomeCodeMap maps from string to cpb.AuditEventOutcomeCode_Value.
var DefaultAuditEventOutcomeCodeMap = map[string]cpb.AuditEventOutcomeCode_Value{
	"INVALID_UNINITIALIZED": cpb.AuditEventOutcomeCode_INVALID_UNINITIALIZED,
	"SUCCESS":               cpb.AuditEventOutcomeCode_SUCCESS,
	"MINOR_FAILURE":         cpb.AuditEventOutcomeCode_MINOR_FAILURE,
	"SERIOUS_FAILURE":       cpb.AuditEventOutcomeCode_SERIOUS_FAILURE,
	"MAJOR_FAILURE":         cpb.AuditEventOutcomeCode_MAJOR_FAILURE,
}

// DefaultBenefitCostApplicabilityCodeMap maps from string to cpb.BenefitCostApplicabilityCode_Value.
var DefaultBenefitCostApplicabilityCodeMap = map[string]cpb.BenefitCostApplicabilityCode_Value{
	"INVALID_UNINITIALIZED": cpb.BenefitCostApplicabilityCode_INVALID_UNINITIALIZED,
	"IN_NETWORK":            cpb.BenefitCostApplicabilityCode_IN_NETWORK,
	"OUT_OF_NETWORK":        cpb.BenefitCostApplicabilityCode_OUT_OF_NETWORK,
	"OTHER":                 cpb.BenefitCostApplicabilityCode_OTHER,
}

// DefaultBindingStrengthCodeMap maps from string to cpb.BindingStrengthCode_Value.
var DefaultBindingStrengthCodeMap = map[string]cpb.BindingStrengthCode_Value{
	"INVALID_UNINITIALIZED": cpb.BindingStrengthCode_INVALID_UNINITIALIZED,
	"REQUIRED":              cpb.BindingStrengthCode_REQUIRED,
	"EXTENSIBLE":            cpb.BindingStrengthCode_EXTENSIBLE,
	"PREFERRED":             cpb.BindingStrengthCode_PREFERRED,
	"EXAMPLE":               cpb.BindingStrengthCode_EXAMPLE,
}

// DefaultBiologicallyDerivedProductCategoryCodeMap maps from string to cpb.BiologicallyDerivedProductCategoryCode_Value.
var DefaultBiologicallyDerivedProductCategoryCodeMap = map[string]cpb.BiologicallyDerivedProductCategoryCode_Value{
	"INVALID_UNINITIALIZED": cpb.BiologicallyDerivedProductCategoryCode_INVALID_UNINITIALIZED,
	"ORGAN":                 cpb.BiologicallyDerivedProductCategoryCode_ORGAN,
	"TISSUE":                cpb.BiologicallyDerivedProductCategoryCode_TISSUE,
	"FLUID":                 cpb.BiologicallyDerivedProductCategoryCode_FLUID,
	"CELLS":                 cpb.BiologicallyDerivedProductCategoryCode_CELLS,
	"BIOLOGICAL_AGENT":      cpb.BiologicallyDerivedProductCategoryCode_BIOLOGICAL_AGENT,
}

// DefaultBiologicallyDerivedProductStatusCodeMap maps from string to cpb.BiologicallyDerivedProductStatusCode_Value.
var DefaultBiologicallyDerivedProductStatusCodeMap = map[string]cpb.BiologicallyDerivedProductStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.BiologicallyDerivedProductStatusCode_INVALID_UNINITIALIZED,
	"AVAILABLE":             cpb.BiologicallyDerivedProductStatusCode_AVAILABLE,
	"UNAVAILABLE":           cpb.BiologicallyDerivedProductStatusCode_UNAVAILABLE,
}

// DefaultBiologicallyDerivedProductStorageScaleCodeMap maps from string to cpb.BiologicallyDerivedProductStorageScaleCode_Value.
var DefaultBiologicallyDerivedProductStorageScaleCodeMap = map[string]cpb.BiologicallyDerivedProductStorageScaleCode_Value{
	"INVALID_UNINITIALIZED": cpb.BiologicallyDerivedProductStorageScaleCode_INVALID_UNINITIALIZED,
	"FARENHEIT":             cpb.BiologicallyDerivedProductStorageScaleCode_FARENHEIT,
	"CELSIUS":               cpb.BiologicallyDerivedProductStorageScaleCode_CELSIUS,
	"KELVIN":                cpb.BiologicallyDerivedProductStorageScaleCode_KELVIN,
}

// DefaultBundleTypeCodeMap maps from string to cpb.BundleTypeCode_Value.
var DefaultBundleTypeCodeMap = map[string]cpb.BundleTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.BundleTypeCode_INVALID_UNINITIALIZED,
	"DOCUMENT":              cpb.BundleTypeCode_DOCUMENT,
	"MESSAGE":               cpb.BundleTypeCode_MESSAGE,
	"TRANSACTION":           cpb.BundleTypeCode_TRANSACTION,
	"TRANSACTION_RESPONSE":  cpb.BundleTypeCode_TRANSACTION_RESPONSE,
	"BATCH":                 cpb.BundleTypeCode_BATCH,
	"BATCH_RESPONSE":        cpb.BundleTypeCode_BATCH_RESPONSE,
	"HISTORY":               cpb.BundleTypeCode_HISTORY,
	"SEARCHSET":             cpb.BundleTypeCode_SEARCHSET,
	"COLLECTION":            cpb.BundleTypeCode_COLLECTION,
}

// DefaultCanonicalStatusCodesForFHIRResourcesCodeMap maps from string to cpb.CanonicalStatusCodesForFHIRResourcesCode_Value.
var DefaultCanonicalStatusCodesForFHIRResourcesCodeMap = map[string]cpb.CanonicalStatusCodesForFHIRResourcesCode_Value{
	"INVALID_UNINITIALIZED": cpb.CanonicalStatusCodesForFHIRResourcesCode_INVALID_UNINITIALIZED,
	"ERROR":                 cpb.CanonicalStatusCodesForFHIRResourcesCode_ERROR,
	"PROPOSED":              cpb.CanonicalStatusCodesForFHIRResourcesCode_PROPOSED,
	"PLANNED":               cpb.CanonicalStatusCodesForFHIRResourcesCode_PLANNED,
	"DRAFT":                 cpb.CanonicalStatusCodesForFHIRResourcesCode_DRAFT,
	"REQUESTED":             cpb.CanonicalStatusCodesForFHIRResourcesCode_REQUESTED,
	"RECEIVED":              cpb.CanonicalStatusCodesForFHIRResourcesCode_RECEIVED,
	"DECLINED":              cpb.CanonicalStatusCodesForFHIRResourcesCode_DECLINED,
	"ACCEPTED":              cpb.CanonicalStatusCodesForFHIRResourcesCode_ACCEPTED,
	"ARRIVED":               cpb.CanonicalStatusCodesForFHIRResourcesCode_ARRIVED,
	"ACTIVE":                cpb.CanonicalStatusCodesForFHIRResourcesCode_ACTIVE,
	"SUSPENDED":             cpb.CanonicalStatusCodesForFHIRResourcesCode_SUSPENDED,
	"FAILED":                cpb.CanonicalStatusCodesForFHIRResourcesCode_FAILED,
	"REPLACED":              cpb.CanonicalStatusCodesForFHIRResourcesCode_REPLACED,
	"COMPLETE":              cpb.CanonicalStatusCodesForFHIRResourcesCode_COMPLETE,
	"INACTIVE":              cpb.CanonicalStatusCodesForFHIRResourcesCode_INACTIVE,
	"ABANDONED":             cpb.CanonicalStatusCodesForFHIRResourcesCode_ABANDONED,
	"UNKNOWN":               cpb.CanonicalStatusCodesForFHIRResourcesCode_UNKNOWN,
	"UNCONFIRMED":           cpb.CanonicalStatusCodesForFHIRResourcesCode_UNCONFIRMED,
	"CONFIRMED":             cpb.CanonicalStatusCodesForFHIRResourcesCode_CONFIRMED,
	"RESOLVED":              cpb.CanonicalStatusCodesForFHIRResourcesCode_RESOLVED,
	"REFUTED":               cpb.CanonicalStatusCodesForFHIRResourcesCode_REFUTED,
	"DIFFERENTIAL":          cpb.CanonicalStatusCodesForFHIRResourcesCode_DIFFERENTIAL,
	"PARTIAL":               cpb.CanonicalStatusCodesForFHIRResourcesCode_PARTIAL,
	"BUSY_UNAVAILABLE":      cpb.CanonicalStatusCodesForFHIRResourcesCode_BUSY_UNAVAILABLE,
	"FREE":                  cpb.CanonicalStatusCodesForFHIRResourcesCode_FREE,
	"ON_TARGET":             cpb.CanonicalStatusCodesForFHIRResourcesCode_ON_TARGET,
	"AHEAD_OF_TARGET":       cpb.CanonicalStatusCodesForFHIRResourcesCode_AHEAD_OF_TARGET,
	"BEHIND_TARGET":         cpb.CanonicalStatusCodesForFHIRResourcesCode_BEHIND_TARGET,
	"NOT_READY":             cpb.CanonicalStatusCodesForFHIRResourcesCode_NOT_READY,
	"TRANSDUC_DISCON":       cpb.CanonicalStatusCodesForFHIRResourcesCode_TRANSDUC_DISCON,
	"HW_DISCON":             cpb.CanonicalStatusCodesForFHIRResourcesCode_HW_DISCON,
}

// DefaultCapabilityStatementKindCodeMap maps from string to cpb.CapabilityStatementKindCode_Value.
var DefaultCapabilityStatementKindCodeMap = map[string]cpb.CapabilityStatementKindCode_Value{
	"INVALID_UNINITIALIZED": cpb.CapabilityStatementKindCode_INVALID_UNINITIALIZED,
	"INSTANCE":              cpb.CapabilityStatementKindCode_INSTANCE,
	"CAPABILITY":            cpb.CapabilityStatementKindCode_CAPABILITY,
	"REQUIREMENTS":          cpb.CapabilityStatementKindCode_REQUIREMENTS,
}

// DefaultCarePlanActivityStatusCodeMap maps from string to cpb.CarePlanActivityStatusCode_Value.
var DefaultCarePlanActivityStatusCodeMap = map[string]cpb.CarePlanActivityStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.CarePlanActivityStatusCode_INVALID_UNINITIALIZED,
	"NOT_STARTED":           cpb.CarePlanActivityStatusCode_NOT_STARTED,
	"SCHEDULED":             cpb.CarePlanActivityStatusCode_SCHEDULED,
	"IN_PROGRESS":           cpb.CarePlanActivityStatusCode_IN_PROGRESS,
	"ON_HOLD":               cpb.CarePlanActivityStatusCode_ON_HOLD,
	"COMPLETED":             cpb.CarePlanActivityStatusCode_COMPLETED,
	"CANCELLED":             cpb.CarePlanActivityStatusCode_CANCELLED,
	"STOPPED":               cpb.CarePlanActivityStatusCode_STOPPED,
	"UNKNOWN":               cpb.CarePlanActivityStatusCode_UNKNOWN,
	"ENTERED_IN_ERROR":      cpb.CarePlanActivityStatusCode_ENTERED_IN_ERROR,
}

// DefaultCareTeamStatusCodeMap maps from string to cpb.CareTeamStatusCode_Value.
var DefaultCareTeamStatusCodeMap = map[string]cpb.CareTeamStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.CareTeamStatusCode_INVALID_UNINITIALIZED,
	"PROPOSED":              cpb.CareTeamStatusCode_PROPOSED,
	"ACTIVE":                cpb.CareTeamStatusCode_ACTIVE,
	"SUSPENDED":             cpb.CareTeamStatusCode_SUSPENDED,
	"INACTIVE":              cpb.CareTeamStatusCode_INACTIVE,
	"ENTERED_IN_ERROR":      cpb.CareTeamStatusCode_ENTERED_IN_ERROR,
}

// DefaultCatalogEntryRelationTypeCodeMap maps from string to cpb.CatalogEntryRelationTypeCode_Value.
var DefaultCatalogEntryRelationTypeCodeMap = map[string]cpb.CatalogEntryRelationTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.CatalogEntryRelationTypeCode_INVALID_UNINITIALIZED,
	"TRIGGERS":              cpb.CatalogEntryRelationTypeCode_TRIGGERS,
	"IS_REPLACED_BY":        cpb.CatalogEntryRelationTypeCode_IS_REPLACED_BY,
}

// DefaultChargeItemStatusCodeMap maps from string to cpb.ChargeItemStatusCode_Value.
var DefaultChargeItemStatusCodeMap = map[string]cpb.ChargeItemStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ChargeItemStatusCode_INVALID_UNINITIALIZED,
	"PLANNED":               cpb.ChargeItemStatusCode_PLANNED,
	"BILLABLE":              cpb.ChargeItemStatusCode_BILLABLE,
	"NOT_BILLABLE":          cpb.ChargeItemStatusCode_NOT_BILLABLE,
	"ABORTED":               cpb.ChargeItemStatusCode_ABORTED,
	"BILLED":                cpb.ChargeItemStatusCode_BILLED,
	"ENTERED_IN_ERROR":      cpb.ChargeItemStatusCode_ENTERED_IN_ERROR,
	"UNKNOWN":               cpb.ChargeItemStatusCode_UNKNOWN,
}

// DefaultChoiceListOrientationCodeMap maps from string to cpb.ChoiceListOrientationCode_Value.
var DefaultChoiceListOrientationCodeMap = map[string]cpb.ChoiceListOrientationCode_Value{
	"INVALID_UNINITIALIZED": cpb.ChoiceListOrientationCode_INVALID_UNINITIALIZED,
	"HORIZONTAL":            cpb.ChoiceListOrientationCode_HORIZONTAL,
	"VERTICAL":              cpb.ChoiceListOrientationCode_VERTICAL,
}

// DefaultClaimProcessingCodeMap maps from string to cpb.ClaimProcessingCode_Value.
var DefaultClaimProcessingCodeMap = map[string]cpb.ClaimProcessingCode_Value{
	"INVALID_UNINITIALIZED": cpb.ClaimProcessingCode_INVALID_UNINITIALIZED,
	"QUEUED":                cpb.ClaimProcessingCode_QUEUED,
	"COMPLETE":              cpb.ClaimProcessingCode_COMPLETE,
	"ERROR":                 cpb.ClaimProcessingCode_ERROR,
	"PARTIAL":               cpb.ClaimProcessingCode_PARTIAL,
}

// DefaultCodeSearchSupportCodeMap maps from string to cpb.CodeSearchSupportCode_Value.
var DefaultCodeSearchSupportCodeMap = map[string]cpb.CodeSearchSupportCode_Value{
	"INVALID_UNINITIALIZED": cpb.CodeSearchSupportCode_INVALID_UNINITIALIZED,
	"EXPLICIT":              cpb.CodeSearchSupportCode_EXPLICIT,
	"ALL":                   cpb.CodeSearchSupportCode_ALL,
}

// DefaultCodeSystemContentModeCodeMap maps from string to cpb.CodeSystemContentModeCode_Value.
var DefaultCodeSystemContentModeCodeMap = map[string]cpb.CodeSystemContentModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.CodeSystemContentModeCode_INVALID_UNINITIALIZED,
	"NOT_PRESENT":           cpb.CodeSystemContentModeCode_NOT_PRESENT,
	"EXAMPLE":               cpb.CodeSystemContentModeCode_EXAMPLE,
	"FRAGMENT":              cpb.CodeSystemContentModeCode_FRAGMENT,
	"COMPLETE":              cpb.CodeSystemContentModeCode_COMPLETE,
	"SUPPLEMENT":            cpb.CodeSystemContentModeCode_SUPPLEMENT,
}

// DefaultCodeSystemHierarchyMeaningCodeMap maps from string to cpb.CodeSystemHierarchyMeaningCode_Value.
var DefaultCodeSystemHierarchyMeaningCodeMap = map[string]cpb.CodeSystemHierarchyMeaningCode_Value{
	"INVALID_UNINITIALIZED": cpb.CodeSystemHierarchyMeaningCode_INVALID_UNINITIALIZED,
	"GROUPED_BY":            cpb.CodeSystemHierarchyMeaningCode_GROUPED_BY,
	"IS_A":                  cpb.CodeSystemHierarchyMeaningCode_IS_A,
	"PART_OF":               cpb.CodeSystemHierarchyMeaningCode_PART_OF,
	"CLASSIFIED_WITH":       cpb.CodeSystemHierarchyMeaningCode_CLASSIFIED_WITH,
}

// DefaultCompartmentTypeCodeMap maps from string to cpb.CompartmentTypeCode_Value.
var DefaultCompartmentTypeCodeMap = map[string]cpb.CompartmentTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.CompartmentTypeCode_INVALID_UNINITIALIZED,
	"PATIENT":               cpb.CompartmentTypeCode_PATIENT,
	"ENCOUNTER":             cpb.CompartmentTypeCode_ENCOUNTER,
	"RELATED_PERSON":        cpb.CompartmentTypeCode_RELATED_PERSON,
	"PRACTITIONER":          cpb.CompartmentTypeCode_PRACTITIONER,
	"DEVICE":                cpb.CompartmentTypeCode_DEVICE,
}

// DefaultCompositionAttestationModeCodeMap maps from string to cpb.CompositionAttestationModeCode_Value.
var DefaultCompositionAttestationModeCodeMap = map[string]cpb.CompositionAttestationModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.CompositionAttestationModeCode_INVALID_UNINITIALIZED,
	"PERSONAL":              cpb.CompositionAttestationModeCode_PERSONAL,
	"PROFESSIONAL":          cpb.CompositionAttestationModeCode_PROFESSIONAL,
	"LEGAL":                 cpb.CompositionAttestationModeCode_LEGAL,
	"OFFICIAL":              cpb.CompositionAttestationModeCode_OFFICIAL,
}

// DefaultCompositionStatusCodeMap maps from string to cpb.CompositionStatusCode_Value.
var DefaultCompositionStatusCodeMap = map[string]cpb.CompositionStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.CompositionStatusCode_INVALID_UNINITIALIZED,
	"PRELIMINARY":           cpb.CompositionStatusCode_PRELIMINARY,
	"FINAL":                 cpb.CompositionStatusCode_FINAL,
	"AMENDED":               cpb.CompositionStatusCode_AMENDED,
	"ENTERED_IN_ERROR":      cpb.CompositionStatusCode_ENTERED_IN_ERROR,
}

// DefaultConceptMapEquivalenceCodeMap maps from string to cpb.ConceptMapEquivalenceCode_Value.
var DefaultConceptMapEquivalenceCodeMap = map[string]cpb.ConceptMapEquivalenceCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConceptMapEquivalenceCode_INVALID_UNINITIALIZED,
	"RELATEDTO":             cpb.ConceptMapEquivalenceCode_RELATEDTO,
	"EQUIVALENT":            cpb.ConceptMapEquivalenceCode_EQUIVALENT,
	"EQUAL":                 cpb.ConceptMapEquivalenceCode_EQUAL,
	"WIDER":                 cpb.ConceptMapEquivalenceCode_WIDER,
	"SUBSUMES":              cpb.ConceptMapEquivalenceCode_SUBSUMES,
	"NARROWER":              cpb.ConceptMapEquivalenceCode_NARROWER,
	"SPECIALIZES":           cpb.ConceptMapEquivalenceCode_SPECIALIZES,
	"INEXACT":               cpb.ConceptMapEquivalenceCode_INEXACT,
	"UNMATCHED":             cpb.ConceptMapEquivalenceCode_UNMATCHED,
	"DISJOINT":              cpb.ConceptMapEquivalenceCode_DISJOINT,
}

// DefaultConceptMapGroupUnmappedModeCodeMap maps from string to cpb.ConceptMapGroupUnmappedModeCode_Value.
var DefaultConceptMapGroupUnmappedModeCodeMap = map[string]cpb.ConceptMapGroupUnmappedModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConceptMapGroupUnmappedModeCode_INVALID_UNINITIALIZED,
	"PROVIDED":              cpb.ConceptMapGroupUnmappedModeCode_PROVIDED,
	"FIXED":                 cpb.ConceptMapGroupUnmappedModeCode_FIXED,
	"OTHER_MAP":             cpb.ConceptMapGroupUnmappedModeCode_OTHER_MAP,
}

// DefaultConditionClinicalStatusCodeMap maps from string to cpb.ConditionClinicalStatusCode_Value.
var DefaultConditionClinicalStatusCodeMap = map[string]cpb.ConditionClinicalStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConditionClinicalStatusCode_INVALID_UNINITIALIZED,
	"ACTIVE":                cpb.ConditionClinicalStatusCode_ACTIVE,
	"RECURRENCE":            cpb.ConditionClinicalStatusCode_RECURRENCE,
	"RELAPSE":               cpb.ConditionClinicalStatusCode_RELAPSE,
	"INACTIVE":              cpb.ConditionClinicalStatusCode_INACTIVE,
	"REMISSION":             cpb.ConditionClinicalStatusCode_REMISSION,
	"RESOLVED":              cpb.ConditionClinicalStatusCode_RESOLVED,
}

// DefaultConditionVerificationStatusCodeMap maps from string to cpb.ConditionVerificationStatusCode_Value.
var DefaultConditionVerificationStatusCodeMap = map[string]cpb.ConditionVerificationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConditionVerificationStatusCode_INVALID_UNINITIALIZED,
	"UNCONFIRMED":           cpb.ConditionVerificationStatusCode_UNCONFIRMED,
	"PROVISIONAL":           cpb.ConditionVerificationStatusCode_PROVISIONAL,
	"DIFFERENTIAL":          cpb.ConditionVerificationStatusCode_DIFFERENTIAL,
	"CONFIRMED":             cpb.ConditionVerificationStatusCode_CONFIRMED,
	"REFUTED":               cpb.ConditionVerificationStatusCode_REFUTED,
	"ENTERED_IN_ERROR":      cpb.ConditionVerificationStatusCode_ENTERED_IN_ERROR,
}

// DefaultConditionalDeleteStatusCodeMap maps from string to cpb.ConditionalDeleteStatusCode_Value.
var DefaultConditionalDeleteStatusCodeMap = map[string]cpb.ConditionalDeleteStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConditionalDeleteStatusCode_INVALID_UNINITIALIZED,
	"NOT_SUPPORTED":         cpb.ConditionalDeleteStatusCode_NOT_SUPPORTED,
	"SINGLE":                cpb.ConditionalDeleteStatusCode_SINGLE,
	"MULTIPLE":              cpb.ConditionalDeleteStatusCode_MULTIPLE,
}

// DefaultConditionalReadStatusCodeMap maps from string to cpb.ConditionalReadStatusCode_Value.
var DefaultConditionalReadStatusCodeMap = map[string]cpb.ConditionalReadStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConditionalReadStatusCode_INVALID_UNINITIALIZED,
	"NOT_SUPPORTED":         cpb.ConditionalReadStatusCode_NOT_SUPPORTED,
	"MODIFIED_SINCE":        cpb.ConditionalReadStatusCode_MODIFIED_SINCE,
	"NOT_MATCH":             cpb.ConditionalReadStatusCode_NOT_MATCH,
	"FULL_SUPPORT":          cpb.ConditionalReadStatusCode_FULL_SUPPORT,
}

// DefaultConformanceExpectationCodeMap maps from string to cpb.ConformanceExpectationCode_Value.
var DefaultConformanceExpectationCodeMap = map[string]cpb.ConformanceExpectationCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConformanceExpectationCode_INVALID_UNINITIALIZED,
	"SHALL":                 cpb.ConformanceExpectationCode_SHALL,
	"SHOULD":                cpb.ConformanceExpectationCode_SHOULD,
	"MAY":                   cpb.ConformanceExpectationCode_MAY,
	"SHOULD_NOT":            cpb.ConformanceExpectationCode_SHOULD_NOT,
}

// DefaultConsentDataMeaningCodeMap maps from string to cpb.ConsentDataMeaningCode_Value.
var DefaultConsentDataMeaningCodeMap = map[string]cpb.ConsentDataMeaningCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConsentDataMeaningCode_INVALID_UNINITIALIZED,
	"INSTANCE":              cpb.ConsentDataMeaningCode_INSTANCE,
	"RELATED":               cpb.ConsentDataMeaningCode_RELATED,
	"DEPENDENTS":            cpb.ConsentDataMeaningCode_DEPENDENTS,
	"AUTHOREDBY":            cpb.ConsentDataMeaningCode_AUTHOREDBY,
}

// DefaultConsentProvisionTypeCodeMap maps from string to cpb.ConsentProvisionTypeCode_Value.
var DefaultConsentProvisionTypeCodeMap = map[string]cpb.ConsentProvisionTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConsentProvisionTypeCode_INVALID_UNINITIALIZED,
	"DENY":                  cpb.ConsentProvisionTypeCode_DENY,
	"PERMIT":                cpb.ConsentProvisionTypeCode_PERMIT,
}

// DefaultConsentStateCodeMap maps from string to cpb.ConsentStateCode_Value.
var DefaultConsentStateCodeMap = map[string]cpb.ConsentStateCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConsentStateCode_INVALID_UNINITIALIZED,
	"DRAFT":                 cpb.ConsentStateCode_DRAFT,
	"PROPOSED":              cpb.ConsentStateCode_PROPOSED,
	"ACTIVE":                cpb.ConsentStateCode_ACTIVE,
	"REJECTED":              cpb.ConsentStateCode_REJECTED,
	"INACTIVE":              cpb.ConsentStateCode_INACTIVE,
	"ENTERED_IN_ERROR":      cpb.ConsentStateCode_ENTERED_IN_ERROR,
}

// DefaultConstraintSeverityCodeMap maps from string to cpb.ConstraintSeverityCode_Value.
var DefaultConstraintSeverityCodeMap = map[string]cpb.ConstraintSeverityCode_Value{
	"INVALID_UNINITIALIZED": cpb.ConstraintSeverityCode_INVALID_UNINITIALIZED,
	"ERROR":                 cpb.ConstraintSeverityCode_ERROR,
	"WARNING":               cpb.ConstraintSeverityCode_WARNING,
}

// DefaultContactPointSystemCodeMap maps from string to cpb.ContactPointSystemCode_Value.
var DefaultContactPointSystemCodeMap = map[string]cpb.ContactPointSystemCode_Value{
	"INVALID_UNINITIALIZED": cpb.ContactPointSystemCode_INVALID_UNINITIALIZED,
	"PHONE":                 cpb.ContactPointSystemCode_PHONE,
	"FAX":                   cpb.ContactPointSystemCode_FAX,
	"EMAIL":                 cpb.ContactPointSystemCode_EMAIL,
	"PAGER":                 cpb.ContactPointSystemCode_PAGER,
	"URL":                   cpb.ContactPointSystemCode_URL,
	"SMS":                   cpb.ContactPointSystemCode_SMS,
	"OTHER":                 cpb.ContactPointSystemCode_OTHER,
}

// DefaultContactPointUseCodeMap maps from string to cpb.ContactPointUseCode_Value.
var DefaultContactPointUseCodeMap = map[string]cpb.ContactPointUseCode_Value{
	"INVALID_UNINITIALIZED": cpb.ContactPointUseCode_INVALID_UNINITIALIZED,
	"HOME":                  cpb.ContactPointUseCode_HOME,
	"WORK":                  cpb.ContactPointUseCode_WORK,
	"TEMP":                  cpb.ContactPointUseCode_TEMP,
	"OLD":                   cpb.ContactPointUseCode_OLD,
	"MOBILE":                cpb.ContactPointUseCode_MOBILE,
}

// DefaultContractResourcePublicationStatusCodeMap maps from string to cpb.ContractResourcePublicationStatusCode_Value.
var DefaultContractResourcePublicationStatusCodeMap = map[string]cpb.ContractResourcePublicationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ContractResourcePublicationStatusCode_INVALID_UNINITIALIZED,
	"AMENDED":               cpb.ContractResourcePublicationStatusCode_AMENDED,
	"APPENDED":              cpb.ContractResourcePublicationStatusCode_APPENDED,
	"CANCELLED":             cpb.ContractResourcePublicationStatusCode_CANCELLED,
	"DISPUTED":              cpb.ContractResourcePublicationStatusCode_DISPUTED,
	"ENTERED_IN_ERROR":      cpb.ContractResourcePublicationStatusCode_ENTERED_IN_ERROR,
	"EXECUTABLE":            cpb.ContractResourcePublicationStatusCode_EXECUTABLE,
	"EXECUTED":              cpb.ContractResourcePublicationStatusCode_EXECUTED,
	"NEGOTIABLE":            cpb.ContractResourcePublicationStatusCode_NEGOTIABLE,
	"OFFERED":               cpb.ContractResourcePublicationStatusCode_OFFERED,
	"POLICY":                cpb.ContractResourcePublicationStatusCode_POLICY,
	"REJECTED":              cpb.ContractResourcePublicationStatusCode_REJECTED,
	"RENEWED":               cpb.ContractResourcePublicationStatusCode_RENEWED,
	"REVOKED":               cpb.ContractResourcePublicationStatusCode_REVOKED,
	"RESOLVED":              cpb.ContractResourcePublicationStatusCode_RESOLVED,
	"TERMINATED":            cpb.ContractResourcePublicationStatusCode_TERMINATED,
}

// DefaultContractResourceStatusCodeMap maps from string to cpb.ContractResourceStatusCode_Value.
var DefaultContractResourceStatusCodeMap = map[string]cpb.ContractResourceStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ContractResourceStatusCode_INVALID_UNINITIALIZED,
	"AMENDED":               cpb.ContractResourceStatusCode_AMENDED,
	"APPENDED":              cpb.ContractResourceStatusCode_APPENDED,
	"CANCELLED":             cpb.ContractResourceStatusCode_CANCELLED,
	"DISPUTED":              cpb.ContractResourceStatusCode_DISPUTED,
	"ENTERED_IN_ERROR":      cpb.ContractResourceStatusCode_ENTERED_IN_ERROR,
	"EXECUTABLE":            cpb.ContractResourceStatusCode_EXECUTABLE,
	"EXECUTED":              cpb.ContractResourceStatusCode_EXECUTED,
	"NEGOTIABLE":            cpb.ContractResourceStatusCode_NEGOTIABLE,
	"OFFERED":               cpb.ContractResourceStatusCode_OFFERED,
	"POLICY":                cpb.ContractResourceStatusCode_POLICY,
	"REJECTED":              cpb.ContractResourceStatusCode_REJECTED,
	"RENEWED":               cpb.ContractResourceStatusCode_RENEWED,
	"REVOKED":               cpb.ContractResourceStatusCode_REVOKED,
	"RESOLVED":              cpb.ContractResourceStatusCode_RESOLVED,
	"TERMINATED":            cpb.ContractResourceStatusCode_TERMINATED,
}

// DefaultContributorTypeCodeMap maps from string to cpb.ContributorTypeCode_Value.
var DefaultContributorTypeCodeMap = map[string]cpb.ContributorTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ContributorTypeCode_INVALID_UNINITIALIZED,
	"AUTHOR":                cpb.ContributorTypeCode_AUTHOR,
	"EDITOR":                cpb.ContributorTypeCode_EDITOR,
	"REVIEWER":              cpb.ContributorTypeCode_REVIEWER,
	"ENDORSER":              cpb.ContributorTypeCode_ENDORSER,
}

// DefaultDataAbsentReasonCodeMap maps from string to cpb.DataAbsentReasonCode_Value.
var DefaultDataAbsentReasonCodeMap = map[string]cpb.DataAbsentReasonCode_Value{
	"INVALID_UNINITIALIZED": cpb.DataAbsentReasonCode_INVALID_UNINITIALIZED,
	"UNKNOWN":               cpb.DataAbsentReasonCode_UNKNOWN,
	"ASKED_UNKNOWN":         cpb.DataAbsentReasonCode_ASKED_UNKNOWN,
	"TEMP_UNKNOWN":          cpb.DataAbsentReasonCode_TEMP_UNKNOWN,
	"NOT_ASKED":             cpb.DataAbsentReasonCode_NOT_ASKED,
	"ASKED_DECLINED":        cpb.DataAbsentReasonCode_ASKED_DECLINED,
	"MASKED":                cpb.DataAbsentReasonCode_MASKED,
	"NOT_APPLICABLE":        cpb.DataAbsentReasonCode_NOT_APPLICABLE,
	"UNSUPPORTED":           cpb.DataAbsentReasonCode_UNSUPPORTED,
	"AS_TEXT":               cpb.DataAbsentReasonCode_AS_TEXT,
	"ERROR":                 cpb.DataAbsentReasonCode_ERROR,
	"NOT_A_NUMBER":          cpb.DataAbsentReasonCode_NOT_A_NUMBER,
	"NEGATIVE_INFINITY":     cpb.DataAbsentReasonCode_NEGATIVE_INFINITY,
	"POSITIVE_INFINITY":     cpb.DataAbsentReasonCode_POSITIVE_INFINITY,
	"NOT_PERFORMED":         cpb.DataAbsentReasonCode_NOT_PERFORMED,
	"NOT_PERMITTED":         cpb.DataAbsentReasonCode_NOT_PERMITTED,
}

// DefaultDataTypeCodeMap maps from string to cpb.DataTypeCode_Value.
var DefaultDataTypeCodeMap = map[string]cpb.DataTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.DataTypeCode_INVALID_UNINITIALIZED,
	"ADDRESS":               cpb.DataTypeCode_ADDRESS,
	"AGE":                   cpb.DataTypeCode_AGE,
	"ANNOTATION":            cpb.DataTypeCode_ANNOTATION,
	"ATTACHMENT":            cpb.DataTypeCode_ATTACHMENT,
	"BACKBONE_ELEMENT":      cpb.DataTypeCode_BACKBONE_ELEMENT,
	"CODEABLE_CONCEPT":      cpb.DataTypeCode_CODEABLE_CONCEPT,
	"CODING":                cpb.DataTypeCode_CODING,
	"CONTACT_DETAIL":        cpb.DataTypeCode_CONTACT_DETAIL,
	"CONTACT_POINT":         cpb.DataTypeCode_CONTACT_POINT,
	"CONTRIBUTOR":           cpb.DataTypeCode_CONTRIBUTOR,
	"COUNT":                 cpb.DataTypeCode_COUNT,
	"DATA_REQUIREMENT":      cpb.DataTypeCode_DATA_REQUIREMENT,
	"DISTANCE":              cpb.DataTypeCode_DISTANCE,
	"DOSAGE":                cpb.DataTypeCode_DOSAGE,
	"DURATION":              cpb.DataTypeCode_DURATION,
	"ELEMENT":               cpb.DataTypeCode_ELEMENT,
	"ELEMENT_DEFINITION":    cpb.DataTypeCode_ELEMENT_DEFINITION,
	"EXPRESSION":            cpb.DataTypeCode_EXPRESSION,
	"EXTENSION":             cpb.DataTypeCode_EXTENSION,
	"HUMAN_NAME":            cpb.DataTypeCode_HUMAN_NAME,
	"IDENTIFIER":            cpb.DataTypeCode_IDENTIFIER,
	"MARKETING_STATUS":      cpb.DataTypeCode_MARKETING_STATUS,
	"META":                  cpb.DataTypeCode_META,
	"MONEY":                 cpb.DataTypeCode_MONEY,
	"MONEY_QUANTITY":        cpb.DataTypeCode_MONEY_QUANTITY,
	"NARRATIVE":             cpb.DataTypeCode_NARRATIVE,
	"PARAMETER_DEFINITION":  cpb.DataTypeCode_PARAMETER_DEFINITION,
	"PERIOD":                cpb.DataTypeCode_PERIOD,
	"POPULATION":            cpb.DataTypeCode_POPULATION,
	"PROD_CHARACTERISTIC":   cpb.DataTypeCode_PROD_CHARACTERISTIC,
	"PRODUCT_SHELF_LIFE":    cpb.DataTypeCode_PRODUCT_SHELF_LIFE,
	"QUANTITY":              cpb.DataTypeCode_QUANTITY,
	"RANGE":                 cpb.DataTypeCode_RANGE,
	"RATIO":                 cpb.DataTypeCode_RATIO,
	"REFERENCE":             cpb.DataTypeCode_REFERENCE,
	"RELATED_ARTIFACT":      cpb.DataTypeCode_RELATED_ARTIFACT,
	"SAMPLED_DATA":          cpb.DataTypeCode_SAMPLED_DATA,
	"SIGNATURE":             cpb.DataTypeCode_SIGNATURE,
	"SIMPLE_QUANTITY":       cpb.DataTypeCode_SIMPLE_QUANTITY,
	"SUBSTANCE_AMOUNT":      cpb.DataTypeCode_SUBSTANCE_AMOUNT,
	"TIMING":                cpb.DataTypeCode_TIMING,
	"TRIGGER_DEFINITION":    cpb.DataTypeCode_TRIGGER_DEFINITION,
	"USAGE_CONTEXT":         cpb.DataTypeCode_USAGE_CONTEXT,
	"BASE64_BINARY":         cpb.DataTypeCode_BASE64_BINARY,
	"BOOLEAN":               cpb.DataTypeCode_BOOLEAN,
	"CANONICAL":             cpb.DataTypeCode_CANONICAL,
	"CODE":                  cpb.DataTypeCode_CODE,
	"DATE":                  cpb.DataTypeCode_DATE,
	"DATE_TIME":             cpb.DataTypeCode_DATE_TIME,
	"DECIMAL":               cpb.DataTypeCode_DECIMAL,
	"ID":                    cpb.DataTypeCode_ID,
	"INSTANT":               cpb.DataTypeCode_INSTANT,
	"INTEGER":               cpb.DataTypeCode_INTEGER,
	"MARKDOWN":              cpb.DataTypeCode_MARKDOWN,
	"OID":                   cpb.DataTypeCode_OID,
	"POSITIVE_INT":          cpb.DataTypeCode_POSITIVE_INT,
	"STRING":                cpb.DataTypeCode_STRING,
	"TIME":                  cpb.DataTypeCode_TIME,
	"UNSIGNED_INT":          cpb.DataTypeCode_UNSIGNED_INT,
	"URI":                   cpb.DataTypeCode_URI,
	"URL":                   cpb.DataTypeCode_URL,
	"UUID":                  cpb.DataTypeCode_UUID,
	"XHTML":                 cpb.DataTypeCode_XHTML,
}

// DefaultDaysOfWeekCodeMap maps from string to cpb.DaysOfWeekCode_Value.
var DefaultDaysOfWeekCodeMap = map[string]cpb.DaysOfWeekCode_Value{
	"INVALID_UNINITIALIZED": cpb.DaysOfWeekCode_INVALID_UNINITIALIZED,
	"MON":                   cpb.DaysOfWeekCode_MON,
	"TUE":                   cpb.DaysOfWeekCode_TUE,
	"WED":                   cpb.DaysOfWeekCode_WED,
	"THU":                   cpb.DaysOfWeekCode_THU,
	"FRI":                   cpb.DaysOfWeekCode_FRI,
	"SAT":                   cpb.DaysOfWeekCode_SAT,
	"SUN":                   cpb.DaysOfWeekCode_SUN,
}

// DefaultDetectedIssueSeverityCodeMap maps from string to cpb.DetectedIssueSeverityCode_Value.
var DefaultDetectedIssueSeverityCodeMap = map[string]cpb.DetectedIssueSeverityCode_Value{
	"INVALID_UNINITIALIZED": cpb.DetectedIssueSeverityCode_INVALID_UNINITIALIZED,
	"HIGH":                  cpb.DetectedIssueSeverityCode_HIGH,
	"MODERATE":              cpb.DetectedIssueSeverityCode_MODERATE,
	"LOW":                   cpb.DetectedIssueSeverityCode_LOW,
}

// DefaultDeviceMetricCalibrationStateCodeMap maps from string to cpb.DeviceMetricCalibrationStateCode_Value.
var DefaultDeviceMetricCalibrationStateCodeMap = map[string]cpb.DeviceMetricCalibrationStateCode_Value{
	"INVALID_UNINITIALIZED": cpb.DeviceMetricCalibrationStateCode_INVALID_UNINITIALIZED,
	"NOT_CALIBRATED":        cpb.DeviceMetricCalibrationStateCode_NOT_CALIBRATED,
	"CALIBRATION_REQUIRED":  cpb.DeviceMetricCalibrationStateCode_CALIBRATION_REQUIRED,
	"CALIBRATED":            cpb.DeviceMetricCalibrationStateCode_CALIBRATED,
	"UNSPECIFIED":           cpb.DeviceMetricCalibrationStateCode_UNSPECIFIED,
}

// DefaultDeviceMetricCalibrationTypeCodeMap maps from string to cpb.DeviceMetricCalibrationTypeCode_Value.
var DefaultDeviceMetricCalibrationTypeCodeMap = map[string]cpb.DeviceMetricCalibrationTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.DeviceMetricCalibrationTypeCode_INVALID_UNINITIALIZED,
	"UNSPECIFIED":           cpb.DeviceMetricCalibrationTypeCode_UNSPECIFIED,
	"OFFSET":                cpb.DeviceMetricCalibrationTypeCode_OFFSET,
	"GAIN":                  cpb.DeviceMetricCalibrationTypeCode_GAIN,
	"TWO_POINT":             cpb.DeviceMetricCalibrationTypeCode_TWO_POINT,
}

// DefaultDeviceMetricCategoryCodeMap maps from string to cpb.DeviceMetricCategoryCode_Value.
var DefaultDeviceMetricCategoryCodeMap = map[string]cpb.DeviceMetricCategoryCode_Value{
	"INVALID_UNINITIALIZED": cpb.DeviceMetricCategoryCode_INVALID_UNINITIALIZED,
	"MEASUREMENT":           cpb.DeviceMetricCategoryCode_MEASUREMENT,
	"SETTING":               cpb.DeviceMetricCategoryCode_SETTING,
	"CALCULATION":           cpb.DeviceMetricCategoryCode_CALCULATION,
	"UNSPECIFIED":           cpb.DeviceMetricCategoryCode_UNSPECIFIED,
}

// DefaultDeviceMetricColorCodeMap maps from string to cpb.DeviceMetricColorCode_Value.
var DefaultDeviceMetricColorCodeMap = map[string]cpb.DeviceMetricColorCode_Value{
	"INVALID_UNINITIALIZED": cpb.DeviceMetricColorCode_INVALID_UNINITIALIZED,
	"BLACK":                 cpb.DeviceMetricColorCode_BLACK,
	"RED":                   cpb.DeviceMetricColorCode_RED,
	"GREEN":                 cpb.DeviceMetricColorCode_GREEN,
	"YELLOW":                cpb.DeviceMetricColorCode_YELLOW,
	"BLUE":                  cpb.DeviceMetricColorCode_BLUE,
	"MAGENTA":               cpb.DeviceMetricColorCode_MAGENTA,
	"CYAN":                  cpb.DeviceMetricColorCode_CYAN,
	"WHITE":                 cpb.DeviceMetricColorCode_WHITE,
}

// DefaultDeviceMetricOperationalStatusCodeMap maps from string to cpb.DeviceMetricOperationalStatusCode_Value.
var DefaultDeviceMetricOperationalStatusCodeMap = map[string]cpb.DeviceMetricOperationalStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.DeviceMetricOperationalStatusCode_INVALID_UNINITIALIZED,
	"ON":                    cpb.DeviceMetricOperationalStatusCode_ON,
	"OFF":                   cpb.DeviceMetricOperationalStatusCode_OFF,
	"STANDBY":               cpb.DeviceMetricOperationalStatusCode_STANDBY,
	"ENTERED_IN_ERROR":      cpb.DeviceMetricOperationalStatusCode_ENTERED_IN_ERROR,
}

// DefaultDeviceNameTypeCodeMap maps from string to cpb.DeviceNameTypeCode_Value.
var DefaultDeviceNameTypeCodeMap = map[string]cpb.DeviceNameTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.DeviceNameTypeCode_INVALID_UNINITIALIZED,
	"UDI_LABEL_NAME":        cpb.DeviceNameTypeCode_UDI_LABEL_NAME,
	"USER_FRIENDLY_NAME":    cpb.DeviceNameTypeCode_USER_FRIENDLY_NAME,
	"PATIENT_REPORTED_NAME": cpb.DeviceNameTypeCode_PATIENT_REPORTED_NAME,
	"MANUFACTURER_NAME":     cpb.DeviceNameTypeCode_MANUFACTURER_NAME,
	"MODEL_NAME":            cpb.DeviceNameTypeCode_MODEL_NAME,
	"OTHER":                 cpb.DeviceNameTypeCode_OTHER,
}

// DefaultDeviceUseStatementStatusCodeMap maps from string to cpb.DeviceUseStatementStatusCode_Value.
var DefaultDeviceUseStatementStatusCodeMap = map[string]cpb.DeviceUseStatementStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.DeviceUseStatementStatusCode_INVALID_UNINITIALIZED,
	"ACTIVE":                cpb.DeviceUseStatementStatusCode_ACTIVE,
	"COMPLETED":             cpb.DeviceUseStatementStatusCode_COMPLETED,
	"ENTERED_IN_ERROR":      cpb.DeviceUseStatementStatusCode_ENTERED_IN_ERROR,
	"INTENDED":              cpb.DeviceUseStatementStatusCode_INTENDED,
	"STOPPED":               cpb.DeviceUseStatementStatusCode_STOPPED,
	"ON_HOLD":               cpb.DeviceUseStatementStatusCode_ON_HOLD,
}

// DefaultDiagnosticReportStatusCodeMap maps from string to cpb.DiagnosticReportStatusCode_Value.
var DefaultDiagnosticReportStatusCodeMap = map[string]cpb.DiagnosticReportStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.DiagnosticReportStatusCode_INVALID_UNINITIALIZED,
	"REGISTERED":            cpb.DiagnosticReportStatusCode_REGISTERED,
	"PARTIAL":               cpb.DiagnosticReportStatusCode_PARTIAL,
	"PRELIMINARY":           cpb.DiagnosticReportStatusCode_PRELIMINARY,
	"FINAL":                 cpb.DiagnosticReportStatusCode_FINAL,
	"AMENDED":               cpb.DiagnosticReportStatusCode_AMENDED,
	"CORRECTED":             cpb.DiagnosticReportStatusCode_CORRECTED,
	"APPENDED":              cpb.DiagnosticReportStatusCode_APPENDED,
	"CANCELLED":             cpb.DiagnosticReportStatusCode_CANCELLED,
	"ENTERED_IN_ERROR":      cpb.DiagnosticReportStatusCode_ENTERED_IN_ERROR,
	"UNKNOWN":               cpb.DiagnosticReportStatusCode_UNKNOWN,
}

// DefaultDiscriminatorTypeCodeMap maps from string to cpb.DiscriminatorTypeCode_Value.
var DefaultDiscriminatorTypeCodeMap = map[string]cpb.DiscriminatorTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.DiscriminatorTypeCode_INVALID_UNINITIALIZED,
	"VALUE":                 cpb.DiscriminatorTypeCode_VALUE,
	"EXISTS":                cpb.DiscriminatorTypeCode_EXISTS,
	"PATTERN":               cpb.DiscriminatorTypeCode_PATTERN,
	"TYPE":                  cpb.DiscriminatorTypeCode_TYPE,
	"PROFILE":               cpb.DiscriminatorTypeCode_PROFILE,
}

// DefaultDocumentModeCodeMap maps from string to cpb.DocumentModeCode_Value.
var DefaultDocumentModeCodeMap = map[string]cpb.DocumentModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.DocumentModeCode_INVALID_UNINITIALIZED,
	"PRODUCER":              cpb.DocumentModeCode_PRODUCER,
	"CONSUMER":              cpb.DocumentModeCode_CONSUMER,
}

// DefaultDocumentReferenceStatusCodeMap maps from string to cpb.DocumentReferenceStatusCode_Value.
var DefaultDocumentReferenceStatusCodeMap = map[string]cpb.DocumentReferenceStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.DocumentReferenceStatusCode_INVALID_UNINITIALIZED,
	"CURRENT":               cpb.DocumentReferenceStatusCode_CURRENT,
	"SUPERSEDED":            cpb.DocumentReferenceStatusCode_SUPERSEDED,
	"ENTERED_IN_ERROR":      cpb.DocumentReferenceStatusCode_ENTERED_IN_ERROR,
}

// DefaultDocumentRelationshipTypeCodeMap maps from string to cpb.DocumentRelationshipTypeCode_Value.
var DefaultDocumentRelationshipTypeCodeMap = map[string]cpb.DocumentRelationshipTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.DocumentRelationshipTypeCode_INVALID_UNINITIALIZED,
	"REPLACES":              cpb.DocumentRelationshipTypeCode_REPLACES,
	"TRANSFORMS":            cpb.DocumentRelationshipTypeCode_TRANSFORMS,
	"SIGNS":                 cpb.DocumentRelationshipTypeCode_SIGNS,
	"APPENDS":               cpb.DocumentRelationshipTypeCode_APPENDS,
}

// DefaultEligibilityRequestPurposeCodeMap maps from string to cpb.EligibilityRequestPurposeCode_Value.
var DefaultEligibilityRequestPurposeCodeMap = map[string]cpb.EligibilityRequestPurposeCode_Value{
	"INVALID_UNINITIALIZED": cpb.EligibilityRequestPurposeCode_INVALID_UNINITIALIZED,
	"AUTH_REQUIREMENTS":     cpb.EligibilityRequestPurposeCode_AUTH_REQUIREMENTS,
	"BENEFITS":              cpb.EligibilityRequestPurposeCode_BENEFITS,
	"DISCOVERY":             cpb.EligibilityRequestPurposeCode_DISCOVERY,
	"VALIDATION":            cpb.EligibilityRequestPurposeCode_VALIDATION,
}

// DefaultEligibilityResponsePurposeCodeMap maps from string to cpb.EligibilityResponsePurposeCode_Value.
var DefaultEligibilityResponsePurposeCodeMap = map[string]cpb.EligibilityResponsePurposeCode_Value{
	"INVALID_UNINITIALIZED": cpb.EligibilityResponsePurposeCode_INVALID_UNINITIALIZED,
	"AUTH_REQUIREMENTS":     cpb.EligibilityResponsePurposeCode_AUTH_REQUIREMENTS,
	"BENEFITS":              cpb.EligibilityResponsePurposeCode_BENEFITS,
	"DISCOVERY":             cpb.EligibilityResponsePurposeCode_DISCOVERY,
	"VALIDATION":            cpb.EligibilityResponsePurposeCode_VALIDATION,
}

// DefaultEnableWhenBehaviorCodeMap maps from string to cpb.EnableWhenBehaviorCode_Value.
var DefaultEnableWhenBehaviorCodeMap = map[string]cpb.EnableWhenBehaviorCode_Value{
	"INVALID_UNINITIALIZED": cpb.EnableWhenBehaviorCode_INVALID_UNINITIALIZED,
	"ALL":                   cpb.EnableWhenBehaviorCode_ALL,
	"ANY":                   cpb.EnableWhenBehaviorCode_ANY,
}

// DefaultEncounterLocationStatusCodeMap maps from string to cpb.EncounterLocationStatusCode_Value.
var DefaultEncounterLocationStatusCodeMap = map[string]cpb.EncounterLocationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.EncounterLocationStatusCode_INVALID_UNINITIALIZED,
	"PLANNED":               cpb.EncounterLocationStatusCode_PLANNED,
	"ACTIVE":                cpb.EncounterLocationStatusCode_ACTIVE,
	"RESERVED":              cpb.EncounterLocationStatusCode_RESERVED,
	"COMPLETED":             cpb.EncounterLocationStatusCode_COMPLETED,
}

// DefaultEncounterStatusCodeMap maps from string to cpb.EncounterStatusCode_Value.
var DefaultEncounterStatusCodeMap = map[string]cpb.EncounterStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.EncounterStatusCode_INVALID_UNINITIALIZED,
	"PLANNED":               cpb.EncounterStatusCode_PLANNED,
	"ARRIVED":               cpb.EncounterStatusCode_ARRIVED,
	"TRIAGED":               cpb.EncounterStatusCode_TRIAGED,
	"IN_PROGRESS":           cpb.EncounterStatusCode_IN_PROGRESS,
	"ONLEAVE":               cpb.EncounterStatusCode_ONLEAVE,
	"FINISHED":              cpb.EncounterStatusCode_FINISHED,
	"CANCELLED":             cpb.EncounterStatusCode_CANCELLED,
	"ENTERED_IN_ERROR":      cpb.EncounterStatusCode_ENTERED_IN_ERROR,
	"UNKNOWN":               cpb.EncounterStatusCode_UNKNOWN,
}

// DefaultEndpointStatusCodeMap maps from string to cpb.EndpointStatusCode_Value.
var DefaultEndpointStatusCodeMap = map[string]cpb.EndpointStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.EndpointStatusCode_INVALID_UNINITIALIZED,
	"ACTIVE":                cpb.EndpointStatusCode_ACTIVE,
	"SUSPENDED":             cpb.EndpointStatusCode_SUSPENDED,
	"ERROR":                 cpb.EndpointStatusCode_ERROR,
	"OFF":                   cpb.EndpointStatusCode_OFF,
	"ENTERED_IN_ERROR":      cpb.EndpointStatusCode_ENTERED_IN_ERROR,
	"TEST":                  cpb.EndpointStatusCode_TEST,
}

// DefaultEpisodeOfCareStatusCodeMap maps from string to cpb.EpisodeOfCareStatusCode_Value.
var DefaultEpisodeOfCareStatusCodeMap = map[string]cpb.EpisodeOfCareStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.EpisodeOfCareStatusCode_INVALID_UNINITIALIZED,
	"PLANNED":               cpb.EpisodeOfCareStatusCode_PLANNED,
	"WAITLIST":              cpb.EpisodeOfCareStatusCode_WAITLIST,
	"ACTIVE":                cpb.EpisodeOfCareStatusCode_ACTIVE,
	"ONHOLD":                cpb.EpisodeOfCareStatusCode_ONHOLD,
	"FINISHED":              cpb.EpisodeOfCareStatusCode_FINISHED,
	"CANCELLED":             cpb.EpisodeOfCareStatusCode_CANCELLED,
	"ENTERED_IN_ERROR":      cpb.EpisodeOfCareStatusCode_ENTERED_IN_ERROR,
}

// DefaultEventCapabilityModeCodeMap maps from string to cpb.EventCapabilityModeCode_Value.
var DefaultEventCapabilityModeCodeMap = map[string]cpb.EventCapabilityModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.EventCapabilityModeCode_INVALID_UNINITIALIZED,
	"SENDER":                cpb.EventCapabilityModeCode_SENDER,
	"RECEIVER":              cpb.EventCapabilityModeCode_RECEIVER,
}

// DefaultEventStatusCodeMap maps from string to cpb.EventStatusCode_Value.
var DefaultEventStatusCodeMap = map[string]cpb.EventStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.EventStatusCode_INVALID_UNINITIALIZED,
	"PREPARATION":           cpb.EventStatusCode_PREPARATION,
	"IN_PROGRESS":           cpb.EventStatusCode_IN_PROGRESS,
	"NOT_DONE":              cpb.EventStatusCode_NOT_DONE,
	"ON_HOLD":               cpb.EventStatusCode_ON_HOLD,
	"STOPPED":               cpb.EventStatusCode_STOPPED,
	"COMPLETED":             cpb.EventStatusCode_COMPLETED,
	"ENTERED_IN_ERROR":      cpb.EventStatusCode_ENTERED_IN_ERROR,
	"UNKNOWN":               cpb.EventStatusCode_UNKNOWN,
}

// DefaultEventTimingCodeMap maps from string to cpb.EventTimingCode_Value.
var DefaultEventTimingCodeMap = map[string]cpb.EventTimingCode_Value{
	"INVALID_UNINITIALIZED": cpb.EventTimingCode_INVALID_UNINITIALIZED,
	"MORN":                  cpb.EventTimingCode_MORN,
	"MORN_EARLY":            cpb.EventTimingCode_MORN_EARLY,
	"MORN_LATE":             cpb.EventTimingCode_MORN_LATE,
	"NOON":                  cpb.EventTimingCode_NOON,
	"AFT":                   cpb.EventTimingCode_AFT,
	"AFT_EARLY":             cpb.EventTimingCode_AFT_EARLY,
	"AFT_LATE":              cpb.EventTimingCode_AFT_LATE,
	"EVE":                   cpb.EventTimingCode_EVE,
	"EVE_EARLY":             cpb.EventTimingCode_EVE_EARLY,
	"EVE_LATE":              cpb.EventTimingCode_EVE_LATE,
	"NIGHT":                 cpb.EventTimingCode_NIGHT,
	"PHS":                   cpb.EventTimingCode_PHS,
}

// DefaultEvidenceVariableTypeCodeMap maps from string to cpb.EvidenceVariableTypeCode_Value.
var DefaultEvidenceVariableTypeCodeMap = map[string]cpb.EvidenceVariableTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.EvidenceVariableTypeCode_INVALID_UNINITIALIZED,
	"DICHOTOMOUS":           cpb.EvidenceVariableTypeCode_DICHOTOMOUS,
	"CONTINUOUS":            cpb.EvidenceVariableTypeCode_CONTINUOUS,
	"DESCRIPTIVE":           cpb.EvidenceVariableTypeCode_DESCRIPTIVE,
}

// DefaultExampleScenarioActorTypeCodeMap maps from string to cpb.ExampleScenarioActorTypeCode_Value.
var DefaultExampleScenarioActorTypeCodeMap = map[string]cpb.ExampleScenarioActorTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ExampleScenarioActorTypeCode_INVALID_UNINITIALIZED,
	"PERSON":                cpb.ExampleScenarioActorTypeCode_PERSON,
	"ENTITY":                cpb.ExampleScenarioActorTypeCode_ENTITY,
}

// DefaultExpansionParameterSourceCodeMap maps from string to cpb.ExpansionParameterSourceCode_Value.
var DefaultExpansionParameterSourceCodeMap = map[string]cpb.ExpansionParameterSourceCode_Value{
	"INVALID_UNINITIALIZED": cpb.ExpansionParameterSourceCode_INVALID_UNINITIALIZED,
	"INPUT":                 cpb.ExpansionParameterSourceCode_INPUT,
	"SERVER":                cpb.ExpansionParameterSourceCode_SERVER,
	"CODESYSTEM":            cpb.ExpansionParameterSourceCode_CODESYSTEM,
}

// DefaultExpansionProcessingRuleCodeMap maps from string to cpb.ExpansionProcessingRuleCode_Value.
var DefaultExpansionProcessingRuleCodeMap = map[string]cpb.ExpansionProcessingRuleCode_Value{
	"INVALID_UNINITIALIZED": cpb.ExpansionProcessingRuleCode_INVALID_UNINITIALIZED,
	"ALL_CODES":             cpb.ExpansionProcessingRuleCode_ALL_CODES,
	"UNGROUPED":             cpb.ExpansionProcessingRuleCode_UNGROUPED,
	"GROUPS_ONLY":           cpb.ExpansionProcessingRuleCode_GROUPS_ONLY,
}

// DefaultExplanationOfBenefitStatusCodeMap maps from string to cpb.ExplanationOfBenefitStatusCode_Value.
var DefaultExplanationOfBenefitStatusCodeMap = map[string]cpb.ExplanationOfBenefitStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ExplanationOfBenefitStatusCode_INVALID_UNINITIALIZED,
	"ACTIVE":                cpb.ExplanationOfBenefitStatusCode_ACTIVE,
	"CANCELLED":             cpb.ExplanationOfBenefitStatusCode_CANCELLED,
	"DRAFT":                 cpb.ExplanationOfBenefitStatusCode_DRAFT,
	"ENTERED_IN_ERROR":      cpb.ExplanationOfBenefitStatusCode_ENTERED_IN_ERROR,
}

// DefaultExposureStateCodeMap maps from string to cpb.ExposureStateCode_Value.
var DefaultExposureStateCodeMap = map[string]cpb.ExposureStateCode_Value{
	"INVALID_UNINITIALIZED": cpb.ExposureStateCode_INVALID_UNINITIALIZED,
	"EXPOSURE":              cpb.ExposureStateCode_EXPOSURE,
	"EXPOSURE_ALTERNATIVE":  cpb.ExposureStateCode_EXPOSURE_ALTERNATIVE,
}

// DefaultExtensionContextTypeCodeMap maps from string to cpb.ExtensionContextTypeCode_Value.
var DefaultExtensionContextTypeCodeMap = map[string]cpb.ExtensionContextTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ExtensionContextTypeCode_INVALID_UNINITIALIZED,
	"FHIRPATH":              cpb.ExtensionContextTypeCode_FHIRPATH,
	"ELEMENT":               cpb.ExtensionContextTypeCode_ELEMENT,
	"EXTENSION":             cpb.ExtensionContextTypeCode_EXTENSION,
}

// DefaultFHIRDeviceStatusCodeMap maps from string to cpb.FHIRDeviceStatusCode_Value.
var DefaultFHIRDeviceStatusCodeMap = map[string]cpb.FHIRDeviceStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.FHIRDeviceStatusCode_INVALID_UNINITIALIZED,
	"ACTIVE":                cpb.FHIRDeviceStatusCode_ACTIVE,
	"INACTIVE":              cpb.FHIRDeviceStatusCode_INACTIVE,
	"ENTERED_IN_ERROR":      cpb.FHIRDeviceStatusCode_ENTERED_IN_ERROR,
	"UNKNOWN":               cpb.FHIRDeviceStatusCode_UNKNOWN,
}

// DefaultFHIRRestfulInteractionsCodeMap maps from string to cpb.FHIRRestfulInteractionsCode_Value.
var DefaultFHIRRestfulInteractionsCodeMap = map[string]cpb.FHIRRestfulInteractionsCode_Value{
	"INVALID_UNINITIALIZED": cpb.FHIRRestfulInteractionsCode_INVALID_UNINITIALIZED,
	"READ":                  cpb.FHIRRestfulInteractionsCode_READ,
	"VREAD":                 cpb.FHIRRestfulInteractionsCode_VREAD,
	"UPDATE":                cpb.FHIRRestfulInteractionsCode_UPDATE,
	"PATCH":                 cpb.FHIRRestfulInteractionsCode_PATCH,
	"DELETE":                cpb.FHIRRestfulInteractionsCode_DELETE,
	"HISTORY":               cpb.FHIRRestfulInteractionsCode_HISTORY,
	"HISTORY_INSTANCE":      cpb.FHIRRestfulInteractionsCode_HISTORY_INSTANCE,
	"HISTORY_TYPE":          cpb.FHIRRestfulInteractionsCode_HISTORY_TYPE,
	"HISTORY_SYSTEM":        cpb.FHIRRestfulInteractionsCode_HISTORY_SYSTEM,
	"CREATE":                cpb.FHIRRestfulInteractionsCode_CREATE,
	"SEARCH":                cpb.FHIRRestfulInteractionsCode_SEARCH,
	"SEARCH_TYPE":           cpb.FHIRRestfulInteractionsCode_SEARCH_TYPE,
	"SEARCH_SYSTEM":         cpb.FHIRRestfulInteractionsCode_SEARCH_SYSTEM,
	"CAPABILITIES":          cpb.FHIRRestfulInteractionsCode_CAPABILITIES,
	"TRANSACTION":           cpb.FHIRRestfulInteractionsCode_TRANSACTION,
	"BATCH":                 cpb.FHIRRestfulInteractionsCode_BATCH,
	"OPERATION":             cpb.FHIRRestfulInteractionsCode_OPERATION,
}

// DefaultFHIRSubstanceStatusCodeMap maps from string to cpb.FHIRSubstanceStatusCode_Value.
var DefaultFHIRSubstanceStatusCodeMap = map[string]cpb.FHIRSubstanceStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.FHIRSubstanceStatusCode_INVALID_UNINITIALIZED,
	"ACTIVE":                cpb.FHIRSubstanceStatusCode_ACTIVE,
	"INACTIVE":              cpb.FHIRSubstanceStatusCode_INACTIVE,
	"ENTERED_IN_ERROR":      cpb.FHIRSubstanceStatusCode_ENTERED_IN_ERROR,
}

// DefaultFHIRVersionCodeMap maps from string to cpb.FHIRVersionCode_Value.
var DefaultFHIRVersionCodeMap = map[string]cpb.FHIRVersionCode_Value{
	"INVALID_UNINITIALIZED": cpb.FHIRVersionCode_INVALID_UNINITIALIZED,
	"V_0_01":                cpb.FHIRVersionCode_V_0_01,
	"V_0_05":                cpb.FHIRVersionCode_V_0_05,
	"V_0_06":                cpb.FHIRVersionCode_V_0_06,
	"V_0_11":                cpb.FHIRVersionCode_V_0_11,
	"V_0_0_80":              cpb.FHIRVersionCode_V_0_0_80,
	"V_0_0_81":              cpb.FHIRVersionCode_V_0_0_81,
	"V_0_0_82":              cpb.FHIRVersionCode_V_0_0_82,
	"V_0_4_0":               cpb.FHIRVersionCode_V_0_4_0,
	"V_0_5_0":               cpb.FHIRVersionCode_V_0_5_0,
	"V_1_0_0":               cpb.FHIRVersionCode_V_1_0_0,
	"V_1_0_1":               cpb.FHIRVersionCode_V_1_0_1,
	"V_1_0_2":               cpb.FHIRVersionCode_V_1_0_2,
	"V_1_1_0":               cpb.FHIRVersionCode_V_1_1_0,
	"V_1_4_0":               cpb.FHIRVersionCode_V_1_4_0,
	"V_1_6_0":               cpb.FHIRVersionCode_V_1_6_0,
	"V_1_8_0":               cpb.FHIRVersionCode_V_1_8_0,
	"V_3_0_0":               cpb.FHIRVersionCode_V_3_0_0,
	"V_3_0_1":               cpb.FHIRVersionCode_V_3_0_1,
	"V_3_3_0":               cpb.FHIRVersionCode_V_3_3_0,
	"V_3_5_0":               cpb.FHIRVersionCode_V_3_5_0,
	"V_4_0_0":               cpb.FHIRVersionCode_V_4_0_0,
	"V_4_0_1":               cpb.FHIRVersionCode_V_4_0_1,
}

// DefaultFamilyHistoryStatusCodeMap maps from string to cpb.FamilyHistoryStatusCode_Value.
var DefaultFamilyHistoryStatusCodeMap = map[string]cpb.FamilyHistoryStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.FamilyHistoryStatusCode_INVALID_UNINITIALIZED,
	"PARTIAL":               cpb.FamilyHistoryStatusCode_PARTIAL,
	"COMPLETED":             cpb.FamilyHistoryStatusCode_COMPLETED,
	"ENTERED_IN_ERROR":      cpb.FamilyHistoryStatusCode_ENTERED_IN_ERROR,
	"HEALTH_UNKNOWN":        cpb.FamilyHistoryStatusCode_HEALTH_UNKNOWN,
}

// DefaultFilterOperatorCodeMap maps from string to cpb.FilterOperatorCode_Value.
var DefaultFilterOperatorCodeMap = map[string]cpb.FilterOperatorCode_Value{
	"INVALID_UNINITIALIZED": cpb.FilterOperatorCode_INVALID_UNINITIALIZED,
	"EQUALS":                cpb.FilterOperatorCode_EQUALS,
	"IS_A":                  cpb.FilterOperatorCode_IS_A,
	"DESCENDENT_OF":         cpb.FilterOperatorCode_DESCENDENT_OF,
	"IS_NOT_A":              cpb.FilterOperatorCode_IS_NOT_A,
	"REGEX":                 cpb.FilterOperatorCode_REGEX,
	"IN":                    cpb.FilterOperatorCode_IN,
	"NOT_IN":                cpb.FilterOperatorCode_NOT_IN,
	"GENERALIZES":           cpb.FilterOperatorCode_GENERALIZES,
	"EXISTS":                cpb.FilterOperatorCode_EXISTS,
}

// DefaultFinancialResourceStatusCodeMap maps from string to cpb.FinancialResourceStatusCode_Value.
var DefaultFinancialResourceStatusCodeMap = map[string]cpb.FinancialResourceStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.FinancialResourceStatusCode_INVALID_UNINITIALIZED,
	"ACTIVE":                cpb.FinancialResourceStatusCode_ACTIVE,
	"CANCELLED":             cpb.FinancialResourceStatusCode_CANCELLED,
	"DRAFT":                 cpb.FinancialResourceStatusCode_DRAFT,
	"ENTERED_IN_ERROR":      cpb.FinancialResourceStatusCode_ENTERED_IN_ERROR,
}

// DefaultFlagStatusCodeMap maps from string to cpb.FlagStatusCode_Value.
var DefaultFlagStatusCodeMap = map[string]cpb.FlagStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.FlagStatusCode_INVALID_UNINITIALIZED,
	"ACTIVE":                cpb.FlagStatusCode_ACTIVE,
	"INACTIVE":              cpb.FlagStatusCode_INACTIVE,
	"ENTERED_IN_ERROR":      cpb.FlagStatusCode_ENTERED_IN_ERROR,
}

// DefaultGoalAcceptanceStatusCodeMap maps from string to cpb.GoalAcceptanceStatusCode_Value.
var DefaultGoalAcceptanceStatusCodeMap = map[string]cpb.GoalAcceptanceStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.GoalAcceptanceStatusCode_INVALID_UNINITIALIZED,
	"AGREE":                 cpb.GoalAcceptanceStatusCode_AGREE,
	"DISAGREE":              cpb.GoalAcceptanceStatusCode_DISAGREE,
	"PENDING":               cpb.GoalAcceptanceStatusCode_PENDING,
}

// DefaultGoalLifecycleStatusCodeMap maps from string to cpb.GoalLifecycleStatusCode_Value.
var DefaultGoalLifecycleStatusCodeMap = map[string]cpb.GoalLifecycleStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.GoalLifecycleStatusCode_INVALID_UNINITIALIZED,
	"PROPOSED":              cpb.GoalLifecycleStatusCode_PROPOSED,
	"PLANNED":               cpb.GoalLifecycleStatusCode_PLANNED,
	"ACCEPTED":              cpb.GoalLifecycleStatusCode_ACCEPTED,
	"ACTIVE":                cpb.GoalLifecycleStatusCode_ACTIVE,
	"ON_HOLD":               cpb.GoalLifecycleStatusCode_ON_HOLD,
	"COMPLETED":             cpb.GoalLifecycleStatusCode_COMPLETED,
	"CANCELLED":             cpb.GoalLifecycleStatusCode_CANCELLED,
	"ENTERED_IN_ERROR":      cpb.GoalLifecycleStatusCode_ENTERED_IN_ERROR,
	"REJECTED":              cpb.GoalLifecycleStatusCode_REJECTED,
}

// DefaultGraphCompartmentRuleCodeMap maps from string to cpb.GraphCompartmentRuleCode_Value.
var DefaultGraphCompartmentRuleCodeMap = map[string]cpb.GraphCompartmentRuleCode_Value{
	"INVALID_UNINITIALIZED": cpb.GraphCompartmentRuleCode_INVALID_UNINITIALIZED,
	"IDENTICAL":             cpb.GraphCompartmentRuleCode_IDENTICAL,
	"MATCHING":              cpb.GraphCompartmentRuleCode_MATCHING,
	"DIFFERENT":             cpb.GraphCompartmentRuleCode_DIFFERENT,
	"CUSTOM":                cpb.GraphCompartmentRuleCode_CUSTOM,
}

// DefaultGraphCompartmentUseCodeMap maps from string to cpb.GraphCompartmentUseCode_Value.
var DefaultGraphCompartmentUseCodeMap = map[string]cpb.GraphCompartmentUseCode_Value{
	"INVALID_UNINITIALIZED": cpb.GraphCompartmentUseCode_INVALID_UNINITIALIZED,
	"CONDITION":             cpb.GraphCompartmentUseCode_CONDITION,
	"REQUIREMENT":           cpb.GraphCompartmentUseCode_REQUIREMENT,
}

// DefaultGroupMeasureCodeMap maps from string to cpb.GroupMeasureCode_Value.
var DefaultGroupMeasureCodeMap = map[string]cpb.GroupMeasureCode_Value{
	"INVALID_UNINITIALIZED": cpb.GroupMeasureCode_INVALID_UNINITIALIZED,
	"MEAN":                  cpb.GroupMeasureCode_MEAN,
	"MEDIAN":                cpb.GroupMeasureCode_MEDIAN,
	"MEAN_OF_MEAN":          cpb.GroupMeasureCode_MEAN_OF_MEAN,
	"MEAN_OF_MEDIAN":        cpb.GroupMeasureCode_MEAN_OF_MEDIAN,
	"MEDIAN_OF_MEAN":        cpb.GroupMeasureCode_MEDIAN_OF_MEAN,
	"MEDIAN_OF_MEDIAN":      cpb.GroupMeasureCode_MEDIAN_OF_MEDIAN,
}

// DefaultGroupTypeCodeMap maps from string to cpb.GroupTypeCode_Value.
var DefaultGroupTypeCodeMap = map[string]cpb.GroupTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.GroupTypeCode_INVALID_UNINITIALIZED,
	"PERSON":                cpb.GroupTypeCode_PERSON,
	"ANIMAL":                cpb.GroupTypeCode_ANIMAL,
	"PRACTITIONER":          cpb.GroupTypeCode_PRACTITIONER,
	"DEVICE":                cpb.GroupTypeCode_DEVICE,
	"MEDICATION":            cpb.GroupTypeCode_MEDICATION,
	"SUBSTANCE":             cpb.GroupTypeCode_SUBSTANCE,
}

// DefaultGuidanceResponseStatusCodeMap maps from string to cpb.GuidanceResponseStatusCode_Value.
var DefaultGuidanceResponseStatusCodeMap = map[string]cpb.GuidanceResponseStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.GuidanceResponseStatusCode_INVALID_UNINITIALIZED,
	"SUCCESS":               cpb.GuidanceResponseStatusCode_SUCCESS,
	"DATA_REQUESTED":        cpb.GuidanceResponseStatusCode_DATA_REQUESTED,
	"DATA_REQUIRED":         cpb.GuidanceResponseStatusCode_DATA_REQUIRED,
	"IN_PROGRESS":           cpb.GuidanceResponseStatusCode_IN_PROGRESS,
	"FAILURE":               cpb.GuidanceResponseStatusCode_FAILURE,
	"ENTERED_IN_ERROR":      cpb.GuidanceResponseStatusCode_ENTERED_IN_ERROR,
}

// DefaultGuidePageGenerationCodeMap maps from string to cpb.GuidePageGenerationCode_Value.
var DefaultGuidePageGenerationCodeMap = map[string]cpb.GuidePageGenerationCode_Value{
	"INVALID_UNINITIALIZED": cpb.GuidePageGenerationCode_INVALID_UNINITIALIZED,
	"HTML":                  cpb.GuidePageGenerationCode_HTML,
	"MARKDOWN":              cpb.GuidePageGenerationCode_MARKDOWN,
	"XML":                   cpb.GuidePageGenerationCode_XML,
	"GENERATED":             cpb.GuidePageGenerationCode_GENERATED,
}

// DefaultGuideParameterCodeMap maps from string to cpb.GuideParameterCode_Value.
var DefaultGuideParameterCodeMap = map[string]cpb.GuideParameterCode_Value{
	"INVALID_UNINITIALIZED": cpb.GuideParameterCode_INVALID_UNINITIALIZED,
	"APPLY":                 cpb.GuideParameterCode_APPLY,
	"PATH_RESOURCE":         cpb.GuideParameterCode_PATH_RESOURCE,
	"PATH_PAGES":            cpb.GuideParameterCode_PATH_PAGES,
	"PATH_TX_CACHE":         cpb.GuideParameterCode_PATH_TX_CACHE,
	"EXPANSION_PARAMETER":   cpb.GuideParameterCode_EXPANSION_PARAMETER,
	"RULE_BROKEN_LINKS":     cpb.GuideParameterCode_RULE_BROKEN_LINKS,
	"GENERATE_XML":          cpb.GuideParameterCode_GENERATE_XML,
	"GENERATE_JSON":         cpb.GuideParameterCode_GENERATE_JSON,
	"GENERATE_TURTLE":       cpb.GuideParameterCode_GENERATE_TURTLE,
	"HTML_TEMPLATE":         cpb.GuideParameterCode_HTML_TEMPLATE,
}

// DefaultHL7WorkgroupCodeMap maps from string to cpb.HL7WorkgroupCode_Value.
var DefaultHL7WorkgroupCodeMap = map[string]cpb.HL7WorkgroupCode_Value{
	"INVALID_UNINITIALIZED": cpb.HL7WorkgroupCode_INVALID_UNINITIALIZED,
	"CBCC":                  cpb.HL7WorkgroupCode_CBCC,
	"CDS":                   cpb.HL7WorkgroupCode_CDS,
	"CQI":                   cpb.HL7WorkgroupCode_CQI,
	"CG":                    cpb.HL7WorkgroupCode_CG,
	"DEV":                   cpb.HL7WorkgroupCode_DEV,
	"EHR":                   cpb.HL7WorkgroupCode_EHR,
	"FHIR":                  cpb.HL7WorkgroupCode_FHIR,
	"FM":                    cpb.HL7WorkgroupCode_FM,
	"HSI":                   cpb.HL7WorkgroupCode_HSI,
	"II":                    cpb.HL7WorkgroupCode_II,
	"INM":                   cpb.HL7WorkgroupCode_INM,
	"ITS":                   cpb.HL7WorkgroupCode_ITS,
	"MNM":                   cpb.HL7WorkgroupCode_MNM,
	"OO":                    cpb.HL7WorkgroupCode_OO,
	"PA":                    cpb.HL7WorkgroupCode_PA,
	"PC":                    cpb.HL7WorkgroupCode_PC,
	"PHER":                  cpb.HL7WorkgroupCode_PHER,
	"PHX":                   cpb.HL7WorkgroupCode_PHX,
	"BRR":                   cpb.HL7WorkgroupCode_BRR,
	"SD":                    cpb.HL7WorkgroupCode_SD,
	"SEC":                   cpb.HL7WorkgroupCode_SEC,
	"US":                    cpb.HL7WorkgroupCode_US,
	"VOCAB":                 cpb.HL7WorkgroupCode_VOCAB,
	"AID":                   cpb.HL7WorkgroupCode_AID,
}

// DefaultHTTPVerbCodeMap maps from string to cpb.HTTPVerbCode_Value.
var DefaultHTTPVerbCodeMap = map[string]cpb.HTTPVerbCode_Value{
	"INVALID_UNINITIALIZED": cpb.HTTPVerbCode_INVALID_UNINITIALIZED,
	"GET":                   cpb.HTTPVerbCode_GET,
	"HEAD":                  cpb.HTTPVerbCode_HEAD,
	"POST":                  cpb.HTTPVerbCode_POST,
	"PUT":                   cpb.HTTPVerbCode_PUT,
	"DELETE":                cpb.HTTPVerbCode_DELETE,
	"PATCH":                 cpb.HTTPVerbCode_PATCH,
}

// DefaultHumanNameAssemblyOrderCodeMap maps from string to cpb.HumanNameAssemblyOrderCode_Value.
var DefaultHumanNameAssemblyOrderCodeMap = map[string]cpb.HumanNameAssemblyOrderCode_Value{
	"INVALID_UNINITIALIZED": cpb.HumanNameAssemblyOrderCode_INVALID_UNINITIALIZED,
	"NL1":                   cpb.HumanNameAssemblyOrderCode_NL1,
	"NL2":                   cpb.HumanNameAssemblyOrderCode_NL2,
	"NL3":                   cpb.HumanNameAssemblyOrderCode_NL3,
	"NL4":                   cpb.HumanNameAssemblyOrderCode_NL4,
}

// DefaultIdentifierUseCodeMap maps from string to cpb.IdentifierUseCode_Value.
var DefaultIdentifierUseCodeMap = map[string]cpb.IdentifierUseCode_Value{
	"INVALID_UNINITIALIZED": cpb.IdentifierUseCode_INVALID_UNINITIALIZED,
	"USUAL":                 cpb.IdentifierUseCode_USUAL,
	"OFFICIAL":              cpb.IdentifierUseCode_OFFICIAL,
	"TEMP":                  cpb.IdentifierUseCode_TEMP,
	"SECONDARY":             cpb.IdentifierUseCode_SECONDARY,
	"OLD":                   cpb.IdentifierUseCode_OLD,
}

// DefaultIdentityAssuranceLevelCodeMap maps from string to cpb.IdentityAssuranceLevelCode_Value.
var DefaultIdentityAssuranceLevelCodeMap = map[string]cpb.IdentityAssuranceLevelCode_Value{
	"INVALID_UNINITIALIZED": cpb.IdentityAssuranceLevelCode_INVALID_UNINITIALIZED,
	"LEVEL1":                cpb.IdentityAssuranceLevelCode_LEVEL1,
	"LEVEL2":                cpb.IdentityAssuranceLevelCode_LEVEL2,
	"LEVEL3":                cpb.IdentityAssuranceLevelCode_LEVEL3,
	"LEVEL4":                cpb.IdentityAssuranceLevelCode_LEVEL4,
}

// DefaultImagingStudyStatusCodeMap maps from string to cpb.ImagingStudyStatusCode_Value.
var DefaultImagingStudyStatusCodeMap = map[string]cpb.ImagingStudyStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ImagingStudyStatusCode_INVALID_UNINITIALIZED,
	"REGISTERED":            cpb.ImagingStudyStatusCode_REGISTERED,
	"AVAILABLE":             cpb.ImagingStudyStatusCode_AVAILABLE,
	"CANCELLED":             cpb.ImagingStudyStatusCode_CANCELLED,
	"ENTERED_IN_ERROR":      cpb.ImagingStudyStatusCode_ENTERED_IN_ERROR,
	"UNKNOWN":               cpb.ImagingStudyStatusCode_UNKNOWN,
}

// DefaultImplantStatusCodeMap maps from string to cpb.ImplantStatusCode_Value.
var DefaultImplantStatusCodeMap = map[string]cpb.ImplantStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ImplantStatusCode_INVALID_UNINITIALIZED,
	"FUNCTIONAL":            cpb.ImplantStatusCode_FUNCTIONAL,
	"NON_FUNCTIONAL":        cpb.ImplantStatusCode_NON_FUNCTIONAL,
	"DISABLED":              cpb.ImplantStatusCode_DISABLED,
	"UNKNOWN":               cpb.ImplantStatusCode_UNKNOWN,
}

// DefaultInvoicePriceComponentTypeCodeMap maps from string to cpb.InvoicePriceComponentTypeCode_Value.
var DefaultInvoicePriceComponentTypeCodeMap = map[string]cpb.InvoicePriceComponentTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.InvoicePriceComponentTypeCode_INVALID_UNINITIALIZED,
	"BASE":                  cpb.InvoicePriceComponentTypeCode_BASE,
	"SURCHARGE":             cpb.InvoicePriceComponentTypeCode_SURCHARGE,
	"DEDUCTION":             cpb.InvoicePriceComponentTypeCode_DEDUCTION,
	"DISCOUNT":              cpb.InvoicePriceComponentTypeCode_DISCOUNT,
	"TAX":                   cpb.InvoicePriceComponentTypeCode_TAX,
	"INFORMATIONAL":         cpb.InvoicePriceComponentTypeCode_INFORMATIONAL,
}

// DefaultInvoiceStatusCodeMap maps from string to cpb.InvoiceStatusCode_Value.
var DefaultInvoiceStatusCodeMap = map[string]cpb.InvoiceStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.InvoiceStatusCode_INVALID_UNINITIALIZED,
	"DRAFT":                 cpb.InvoiceStatusCode_DRAFT,
	"ISSUED":                cpb.InvoiceStatusCode_ISSUED,
	"BALANCED":              cpb.InvoiceStatusCode_BALANCED,
	"CANCELLED":             cpb.InvoiceStatusCode_CANCELLED,
	"ENTERED_IN_ERROR":      cpb.InvoiceStatusCode_ENTERED_IN_ERROR,
}

// DefaultIssueSeverityCodeMap maps from string to cpb.IssueSeverityCode_Value.
var DefaultIssueSeverityCodeMap = map[string]cpb.IssueSeverityCode_Value{
	"INVALID_UNINITIALIZED": cpb.IssueSeverityCode_INVALID_UNINITIALIZED,
	"FATAL":                 cpb.IssueSeverityCode_FATAL,
	"ERROR":                 cpb.IssueSeverityCode_ERROR,
	"WARNING":               cpb.IssueSeverityCode_WARNING,
	"INFORMATION":           cpb.IssueSeverityCode_INFORMATION,
}

// DefaultIssueTypeCodeMap maps from string to cpb.IssueTypeCode_Value.
var DefaultIssueTypeCodeMap = map[string]cpb.IssueTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.IssueTypeCode_INVALID_UNINITIALIZED,
	"INVALID":               cpb.IssueTypeCode_INVALID,
	"STRUCTURE":             cpb.IssueTypeCode_STRUCTURE,
	"REQUIRED":              cpb.IssueTypeCode_REQUIRED,
	"VALUE":                 cpb.IssueTypeCode_VALUE,
	"INVARIANT":             cpb.IssueTypeCode_INVARIANT,
	"SECURITY":              cpb.IssueTypeCode_SECURITY,
	"LOGIN":                 cpb.IssueTypeCode_LOGIN,
	"UNKNOWN":               cpb.IssueTypeCode_UNKNOWN,
	"EXPIRED":               cpb.IssueTypeCode_EXPIRED,
	"FORBIDDEN":             cpb.IssueTypeCode_FORBIDDEN,
	"SUPPRESSED":            cpb.IssueTypeCode_SUPPRESSED,
	"PROCESSING":            cpb.IssueTypeCode_PROCESSING,
	"NOT_SUPPORTED":         cpb.IssueTypeCode_NOT_SUPPORTED,
	"DUPLICATE":             cpb.IssueTypeCode_DUPLICATE,
	"MULTIPLE_MATCHES":      cpb.IssueTypeCode_MULTIPLE_MATCHES,
	"NOT_FOUND":             cpb.IssueTypeCode_NOT_FOUND,
	"DELETED":               cpb.IssueTypeCode_DELETED,
	"TOO_LONG":              cpb.IssueTypeCode_TOO_LONG,
	"CODE_INVALID":          cpb.IssueTypeCode_CODE_INVALID,
	"EXTENSION":             cpb.IssueTypeCode_EXTENSION,
	"TOO_COSTLY":            cpb.IssueTypeCode_TOO_COSTLY,
	"BUSINESS_RULE":         cpb.IssueTypeCode_BUSINESS_RULE,
	"CONFLICT":              cpb.IssueTypeCode_CONFLICT,
	"TRANSIENT":             cpb.IssueTypeCode_TRANSIENT,
	"LOCK_ERROR":            cpb.IssueTypeCode_LOCK_ERROR,
	"NO_STORE":              cpb.IssueTypeCode_NO_STORE,
	"EXCEPTION":             cpb.IssueTypeCode_EXCEPTION,
	"TIMEOUT":               cpb.IssueTypeCode_TIMEOUT,
	"INCOMPLETE":            cpb.IssueTypeCode_INCOMPLETE,
	"THROTTLED":             cpb.IssueTypeCode_THROTTLED,
	"INFORMATIONAL":         cpb.IssueTypeCode_INFORMATIONAL,
}

// DefaultLinkTypeCodeMap maps from string to cpb.LinkTypeCode_Value.
var DefaultLinkTypeCodeMap = map[string]cpb.LinkTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.LinkTypeCode_INVALID_UNINITIALIZED,
	"REPLACED_BY":           cpb.LinkTypeCode_REPLACED_BY,
	"REPLACES":              cpb.LinkTypeCode_REPLACES,
	"REFER":                 cpb.LinkTypeCode_REFER,
	"SEEALSO":               cpb.LinkTypeCode_SEEALSO,
}

// DefaultLinkageTypeCodeMap maps from string to cpb.LinkageTypeCode_Value.
var DefaultLinkageTypeCodeMap = map[string]cpb.LinkageTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.LinkageTypeCode_INVALID_UNINITIALIZED,
	"SOURCE":                cpb.LinkageTypeCode_SOURCE,
	"ALTERNATE":             cpb.LinkageTypeCode_ALTERNATE,
	"HISTORICAL":            cpb.LinkageTypeCode_HISTORICAL,
}

// DefaultListModeCodeMap maps from string to cpb.ListModeCode_Value.
var DefaultListModeCodeMap = map[string]cpb.ListModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ListModeCode_INVALID_UNINITIALIZED,
	"WORKING":               cpb.ListModeCode_WORKING,
	"SNAPSHOT":              cpb.ListModeCode_SNAPSHOT,
	"CHANGES":               cpb.ListModeCode_CHANGES,
}

// DefaultListStatusCodeMap maps from string to cpb.ListStatusCode_Value.
var DefaultListStatusCodeMap = map[string]cpb.ListStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ListStatusCode_INVALID_UNINITIALIZED,
	"CURRENT":               cpb.ListStatusCode_CURRENT,
	"RETIRED":               cpb.ListStatusCode_RETIRED,
	"ENTERED_IN_ERROR":      cpb.ListStatusCode_ENTERED_IN_ERROR,
}

// DefaultLocationModeCodeMap maps from string to cpb.LocationModeCode_Value.
var DefaultLocationModeCodeMap = map[string]cpb.LocationModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.LocationModeCode_INVALID_UNINITIALIZED,
	"INSTANCE":              cpb.LocationModeCode_INSTANCE,
	"KIND":                  cpb.LocationModeCode_KIND,
}

// DefaultLocationStatusCodeMap maps from string to cpb.LocationStatusCode_Value.
var DefaultLocationStatusCodeMap = map[string]cpb.LocationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.LocationStatusCode_INVALID_UNINITIALIZED,
	"ACTIVE":                cpb.LocationStatusCode_ACTIVE,
	"SUSPENDED":             cpb.LocationStatusCode_SUSPENDED,
	"INACTIVE":              cpb.LocationStatusCode_INACTIVE,
}

// DefaultMatchGradeCodeMap maps from string to cpb.MatchGradeCode_Value.
var DefaultMatchGradeCodeMap = map[string]cpb.MatchGradeCode_Value{
	"INVALID_UNINITIALIZED": cpb.MatchGradeCode_INVALID_UNINITIALIZED,
	"CERTAIN":               cpb.MatchGradeCode_CERTAIN,
	"PROBABLE":              cpb.MatchGradeCode_PROBABLE,
	"POSSIBLE":              cpb.MatchGradeCode_POSSIBLE,
	"CERTAINLY_NOT":         cpb.MatchGradeCode_CERTAINLY_NOT,
}

// DefaultMeasureImprovementNotationCodeMap maps from string to cpb.MeasureImprovementNotationCode_Value.
var DefaultMeasureImprovementNotationCodeMap = map[string]cpb.MeasureImprovementNotationCode_Value{
	"INVALID_UNINITIALIZED": cpb.MeasureImprovementNotationCode_INVALID_UNINITIALIZED,
	"INCREASE":              cpb.MeasureImprovementNotationCode_INCREASE,
	"DECREASE":              cpb.MeasureImprovementNotationCode_DECREASE,
}

// DefaultMeasureReportStatusCodeMap maps from string to cpb.MeasureReportStatusCode_Value.
var DefaultMeasureReportStatusCodeMap = map[string]cpb.MeasureReportStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.MeasureReportStatusCode_INVALID_UNINITIALIZED,
	"COMPLETE":              cpb.MeasureReportStatusCode_COMPLETE,
	"PENDING":               cpb.MeasureReportStatusCode_PENDING,
	"ERROR":                 cpb.MeasureReportStatusCode_ERROR,
}

// DefaultMeasureReportTypeCodeMap maps from string to cpb.MeasureReportTypeCode_Value.
var DefaultMeasureReportTypeCodeMap = map[string]cpb.MeasureReportTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.MeasureReportTypeCode_INVALID_UNINITIALIZED,
	"INDIVIDUAL":            cpb.MeasureReportTypeCode_INDIVIDUAL,
	"SUBJECT_LIST":          cpb.MeasureReportTypeCode_SUBJECT_LIST,
	"SUMMARY":               cpb.MeasureReportTypeCode_SUMMARY,
	"DATA_COLLECTION":       cpb.MeasureReportTypeCode_DATA_COLLECTION,
}

// DefaultMedicationAdministrationStatusCodeMap maps from string to cpb.MedicationAdministrationStatusCode_Value.
var DefaultMedicationAdministrationStatusCodeMap = map[string]cpb.MedicationAdministrationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.MedicationAdministrationStatusCode_INVALID_UNINITIALIZED,
	"IN_PROGRESS":           cpb.MedicationAdministrationStatusCode_IN_PROGRESS,
	"NOT_DONE":              cpb.MedicationAdministrationStatusCode_NOT_DONE,
	"ON_HOLD":               cpb.MedicationAdministrationStatusCode_ON_HOLD,
	"COMPLETED":             cpb.MedicationAdministrationStatusCode_COMPLETED,
	"ENTERED_IN_ERROR":      cpb.MedicationAdministrationStatusCode_ENTERED_IN_ERROR,
	"STOPPED":               cpb.MedicationAdministrationStatusCode_STOPPED,
	"UNKNOWN":               cpb.MedicationAdministrationStatusCode_UNKNOWN,
}

// DefaultMedicationDispenseStatusCodeMap maps from string to cpb.MedicationDispenseStatusCode_Value.
var DefaultMedicationDispenseStatusCodeMap = map[string]cpb.MedicationDispenseStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.MedicationDispenseStatusCode_INVALID_UNINITIALIZED,
	"PREPARATION":           cpb.MedicationDispenseStatusCode_PREPARATION,
	"IN_PROGRESS":           cpb.MedicationDispenseStatusCode_IN_PROGRESS,
	"CANCELLED":             cpb.MedicationDispenseStatusCode_CANCELLED,
	"ON_HOLD":               cpb.MedicationDispenseStatusCode_ON_HOLD,
	"COMPLETED":             cpb.MedicationDispenseStatusCode_COMPLETED,
	"ENTERED_IN_ERROR":      cpb.MedicationDispenseStatusCode_ENTERED_IN_ERROR,
	"STOPPED":               cpb.MedicationDispenseStatusCode_STOPPED,
	"DECLINED":              cpb.MedicationDispenseStatusCode_DECLINED,
	"UNKNOWN":               cpb.MedicationDispenseStatusCode_UNKNOWN,
}

// DefaultMedicationKnowledgeStatusCodeMap maps from string to cpb.MedicationKnowledgeStatusCode_Value.
var DefaultMedicationKnowledgeStatusCodeMap = map[string]cpb.MedicationKnowledgeStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.MedicationKnowledgeStatusCode_INVALID_UNINITIALIZED,
	"ACTIVE":                cpb.MedicationKnowledgeStatusCode_ACTIVE,
	"INACTIVE":              cpb.MedicationKnowledgeStatusCode_INACTIVE,
	"ENTERED_IN_ERROR":      cpb.MedicationKnowledgeStatusCode_ENTERED_IN_ERROR,
}

// DefaultMedicationRequestIntentCodeMap maps from string to cpb.MedicationRequestIntentCode_Value.
var DefaultMedicationRequestIntentCodeMap = map[string]cpb.MedicationRequestIntentCode_Value{
	"INVALID_UNINITIALIZED": cpb.MedicationRequestIntentCode_INVALID_UNINITIALIZED,
	"PROPOSAL":              cpb.MedicationRequestIntentCode_PROPOSAL,
	"PLAN":                  cpb.MedicationRequestIntentCode_PLAN,
	"ORDER":                 cpb.MedicationRequestIntentCode_ORDER,
	"ORIGINAL_ORDER":        cpb.MedicationRequestIntentCode_ORIGINAL_ORDER,
	"REFLEX_ORDER":          cpb.MedicationRequestIntentCode_REFLEX_ORDER,
	"FILLER_ORDER":          cpb.MedicationRequestIntentCode_FILLER_ORDER,
	"INSTANCE_ORDER":        cpb.MedicationRequestIntentCode_INSTANCE_ORDER,
	"OPTION":                cpb.MedicationRequestIntentCode_OPTION,
}

// DefaultMedicationStatementStatusCodesMap maps from string to cpb.MedicationStatementStatusCodes_Value.
var DefaultMedicationStatementStatusCodesMap = map[string]cpb.MedicationStatementStatusCodes_Value{
	"INVALID_UNINITIALIZED": cpb.MedicationStatementStatusCodes_INVALID_UNINITIALIZED,
	"ACTIVE":                cpb.MedicationStatementStatusCodes_ACTIVE,
	"COMPLETED":             cpb.MedicationStatementStatusCodes_COMPLETED,
	"ENTERED_IN_ERROR":      cpb.MedicationStatementStatusCodes_ENTERED_IN_ERROR,
	"INTENDED":              cpb.MedicationStatementStatusCodes_INTENDED,
	"STOPPED":               cpb.MedicationStatementStatusCodes_STOPPED,
	"ON_HOLD":               cpb.MedicationStatementStatusCodes_ON_HOLD,
	"UNKNOWN":               cpb.MedicationStatementStatusCodes_UNKNOWN,
	"NOT_TAKEN":             cpb.MedicationStatementStatusCodes_NOT_TAKEN,
}

// DefaultMedicationStatusCodeMap maps from string to cpb.MedicationStatusCode_Value.
var DefaultMedicationStatusCodeMap = map[string]cpb.MedicationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.MedicationStatusCode_INVALID_UNINITIALIZED,
	"ACTIVE":                cpb.MedicationStatusCode_ACTIVE,
	"INACTIVE":              cpb.MedicationStatusCode_INACTIVE,
	"ENTERED_IN_ERROR":      cpb.MedicationStatusCode_ENTERED_IN_ERROR,
}

// DefaultMedicationrequestStatusCodeMap maps from string to cpb.MedicationrequestStatusCode_Value.
var DefaultMedicationrequestStatusCodeMap = map[string]cpb.MedicationrequestStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.MedicationrequestStatusCode_INVALID_UNINITIALIZED,
	"ACTIVE":                cpb.MedicationrequestStatusCode_ACTIVE,
	"ON_HOLD":               cpb.MedicationrequestStatusCode_ON_HOLD,
	"CANCELLED":             cpb.MedicationrequestStatusCode_CANCELLED,
	"COMPLETED":             cpb.MedicationrequestStatusCode_COMPLETED,
	"ENTERED_IN_ERROR":      cpb.MedicationrequestStatusCode_ENTERED_IN_ERROR,
	"STOPPED":               cpb.MedicationrequestStatusCode_STOPPED,
	"DRAFT":                 cpb.MedicationrequestStatusCode_DRAFT,
	"UNKNOWN":               cpb.MedicationrequestStatusCode_UNKNOWN,
}

// DefaultMessageSignificanceCategoryCodeMap maps from string to cpb.MessageSignificanceCategoryCode_Value.
var DefaultMessageSignificanceCategoryCodeMap = map[string]cpb.MessageSignificanceCategoryCode_Value{
	"INVALID_UNINITIALIZED": cpb.MessageSignificanceCategoryCode_INVALID_UNINITIALIZED,
	"CONSEQUENCE":           cpb.MessageSignificanceCategoryCode_CONSEQUENCE,
	"CURRENCY":              cpb.MessageSignificanceCategoryCode_CURRENCY,
	"NOTIFICATION":          cpb.MessageSignificanceCategoryCode_NOTIFICATION,
}

// DefaultMessageheaderResponseRequestCodeMap maps from string to cpb.MessageheaderResponseRequestCode_Value.
var DefaultMessageheaderResponseRequestCodeMap = map[string]cpb.MessageheaderResponseRequestCode_Value{
	"INVALID_UNINITIALIZED": cpb.MessageheaderResponseRequestCode_INVALID_UNINITIALIZED,
	"ALWAYS":                cpb.MessageheaderResponseRequestCode_ALWAYS,
	"ON_ERROR":              cpb.MessageheaderResponseRequestCode_ON_ERROR,
	"NEVER":                 cpb.MessageheaderResponseRequestCode_NEVER,
	"ON_SUCCESS":            cpb.MessageheaderResponseRequestCode_ON_SUCCESS,
}

// DefaultNameUseCodeMap maps from string to cpb.NameUseCode_Value.
var DefaultNameUseCodeMap = map[string]cpb.NameUseCode_Value{
	"INVALID_UNINITIALIZED": cpb.NameUseCode_INVALID_UNINITIALIZED,
	"USUAL":                 cpb.NameUseCode_USUAL,
	"OFFICIAL":              cpb.NameUseCode_OFFICIAL,
	"TEMP":                  cpb.NameUseCode_TEMP,
	"NICKNAME":              cpb.NameUseCode_NICKNAME,
	"ANONYMOUS":             cpb.NameUseCode_ANONYMOUS,
	"OLD":                   cpb.NameUseCode_OLD,
	"MAIDEN":                cpb.NameUseCode_MAIDEN,
}

// DefaultNamingSystemIdentifierTypeCodeMap maps from string to cpb.NamingSystemIdentifierTypeCode_Value.
var DefaultNamingSystemIdentifierTypeCodeMap = map[string]cpb.NamingSystemIdentifierTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.NamingSystemIdentifierTypeCode_INVALID_UNINITIALIZED,
	"OID":                   cpb.NamingSystemIdentifierTypeCode_OID,
	"UUID":                  cpb.NamingSystemIdentifierTypeCode_UUID,
	"URI":                   cpb.NamingSystemIdentifierTypeCode_URI,
	"OTHER":                 cpb.NamingSystemIdentifierTypeCode_OTHER,
}

// DefaultNamingSystemTypeCodeMap maps from string to cpb.NamingSystemTypeCode_Value.
var DefaultNamingSystemTypeCodeMap = map[string]cpb.NamingSystemTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.NamingSystemTypeCode_INVALID_UNINITIALIZED,
	"CODESYSTEM":            cpb.NamingSystemTypeCode_CODESYSTEM,
	"IDENTIFIER":            cpb.NamingSystemTypeCode_IDENTIFIER,
	"ROOT":                  cpb.NamingSystemTypeCode_ROOT,
}

// DefaultNarrativeStatusCodeMap maps from string to cpb.NarrativeStatusCode_Value.
var DefaultNarrativeStatusCodeMap = map[string]cpb.NarrativeStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.NarrativeStatusCode_INVALID_UNINITIALIZED,
	"GENERATED":             cpb.NarrativeStatusCode_GENERATED,
	"EXTENSIONS":            cpb.NarrativeStatusCode_EXTENSIONS,
	"ADDITIONAL":            cpb.NarrativeStatusCode_ADDITIONAL,
	"EMPTY":                 cpb.NarrativeStatusCode_EMPTY,
}

// DefaultNoteTypeCodeMap maps from string to cpb.NoteTypeCode_Value.
var DefaultNoteTypeCodeMap = map[string]cpb.NoteTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.NoteTypeCode_INVALID_UNINITIALIZED,
	"DISPLAY":               cpb.NoteTypeCode_DISPLAY,
	"PRINT":                 cpb.NoteTypeCode_PRINT,
	"PRINTOPER":             cpb.NoteTypeCode_PRINTOPER,
}

// DefaultObservationDataTypeCodeMap maps from string to cpb.ObservationDataTypeCode_Value.
var DefaultObservationDataTypeCodeMap = map[string]cpb.ObservationDataTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ObservationDataTypeCode_INVALID_UNINITIALIZED,
	"QUANTITY":              cpb.ObservationDataTypeCode_QUANTITY,
	"CODEABLE_CONCEPT":      cpb.ObservationDataTypeCode_CODEABLE_CONCEPT,
	"STRING":                cpb.ObservationDataTypeCode_STRING,
	"BOOLEAN":               cpb.ObservationDataTypeCode_BOOLEAN,
	"INTEGER":               cpb.ObservationDataTypeCode_INTEGER,
	"RANGE":                 cpb.ObservationDataTypeCode_RANGE,
	"RATIO":                 cpb.ObservationDataTypeCode_RATIO,
	"SAMPLED_DATA":          cpb.ObservationDataTypeCode_SAMPLED_DATA,
	"TIME":                  cpb.ObservationDataTypeCode_TIME,
	"DATE_TIME":             cpb.ObservationDataTypeCode_DATE_TIME,
	"PERIOD":                cpb.ObservationDataTypeCode_PERIOD,
}

// DefaultObservationRangeCategoryCodeMap maps from string to cpb.ObservationRangeCategoryCode_Value.
var DefaultObservationRangeCategoryCodeMap = map[string]cpb.ObservationRangeCategoryCode_Value{
	"INVALID_UNINITIALIZED": cpb.ObservationRangeCategoryCode_INVALID_UNINITIALIZED,
	"REFERENCE":             cpb.ObservationRangeCategoryCode_REFERENCE,
	"CRITICAL":              cpb.ObservationRangeCategoryCode_CRITICAL,
	"ABSOLUTE":              cpb.ObservationRangeCategoryCode_ABSOLUTE,
}

// DefaultObservationStatusCodeMap maps from string to cpb.ObservationStatusCode_Value.
var DefaultObservationStatusCodeMap = map[string]cpb.ObservationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ObservationStatusCode_INVALID_UNINITIALIZED,
	"REGISTERED":            cpb.ObservationStatusCode_REGISTERED,
	"PRELIMINARY":           cpb.ObservationStatusCode_PRELIMINARY,
	"FINAL":                 cpb.ObservationStatusCode_FINAL,
	"AMENDED":               cpb.ObservationStatusCode_AMENDED,
	"CORRECTED":             cpb.ObservationStatusCode_CORRECTED,
	"CANCELLED":             cpb.ObservationStatusCode_CANCELLED,
	"ENTERED_IN_ERROR":      cpb.ObservationStatusCode_ENTERED_IN_ERROR,
	"UNKNOWN":               cpb.ObservationStatusCode_UNKNOWN,
}

// DefaultOperationKindCodeMap maps from string to cpb.OperationKindCode_Value.
var DefaultOperationKindCodeMap = map[string]cpb.OperationKindCode_Value{
	"INVALID_UNINITIALIZED": cpb.OperationKindCode_INVALID_UNINITIALIZED,
	"OPERATION":             cpb.OperationKindCode_OPERATION,
	"QUERY":                 cpb.OperationKindCode_QUERY,
}

// DefaultOperationParameterUseCodeMap maps from string to cpb.OperationParameterUseCode_Value.
var DefaultOperationParameterUseCodeMap = map[string]cpb.OperationParameterUseCode_Value{
	"INVALID_UNINITIALIZED": cpb.OperationParameterUseCode_INVALID_UNINITIALIZED,
	"IN":                    cpb.OperationParameterUseCode_IN,
	"OUT":                   cpb.OperationParameterUseCode_OUT,
}

// DefaultOrientationTypeCodeMap maps from string to cpb.OrientationTypeCode_Value.
var DefaultOrientationTypeCodeMap = map[string]cpb.OrientationTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.OrientationTypeCode_INVALID_UNINITIALIZED,
	"SENSE":                 cpb.OrientationTypeCode_SENSE,
	"ANTISENSE":             cpb.OrientationTypeCode_ANTISENSE,
}

// DefaultParticipantRequiredCodeMap maps from string to cpb.ParticipantRequiredCode_Value.
var DefaultParticipantRequiredCodeMap = map[string]cpb.ParticipantRequiredCode_Value{
	"INVALID_UNINITIALIZED": cpb.ParticipantRequiredCode_INVALID_UNINITIALIZED,
	"REQUIRED":              cpb.ParticipantRequiredCode_REQUIRED,
	"OPTIONAL":              cpb.ParticipantRequiredCode_OPTIONAL,
	"INFORMATION_ONLY":      cpb.ParticipantRequiredCode_INFORMATION_ONLY,
}

// DefaultParticipationStatusCodeMap maps from string to cpb.ParticipationStatusCode_Value.
var DefaultParticipationStatusCodeMap = map[string]cpb.ParticipationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ParticipationStatusCode_INVALID_UNINITIALIZED,
	"ACCEPTED":              cpb.ParticipationStatusCode_ACCEPTED,
	"DECLINED":              cpb.ParticipationStatusCode_DECLINED,
	"TENTATIVE":             cpb.ParticipationStatusCode_TENTATIVE,
	"NEEDS_ACTION":          cpb.ParticipationStatusCode_NEEDS_ACTION,
}

// DefaultPropertyRepresentationCodeMap maps from string to cpb.PropertyRepresentationCode_Value.
var DefaultPropertyRepresentationCodeMap = map[string]cpb.PropertyRepresentationCode_Value{
	"INVALID_UNINITIALIZED": cpb.PropertyRepresentationCode_INVALID_UNINITIALIZED,
	"XML_ATTR":              cpb.PropertyRepresentationCode_XML_ATTR,
	"XML_TEXT":              cpb.PropertyRepresentationCode_XML_TEXT,
	"TYPE_ATTR":             cpb.PropertyRepresentationCode_TYPE_ATTR,
	"CDA_TEXT":              cpb.PropertyRepresentationCode_CDA_TEXT,
	"XHTML":                 cpb.PropertyRepresentationCode_XHTML,
}

// DefaultPropertyTypeCodeMap maps from string to cpb.PropertyTypeCode_Value.
var DefaultPropertyTypeCodeMap = map[string]cpb.PropertyTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.PropertyTypeCode_INVALID_UNINITIALIZED,
	"CODE":                  cpb.PropertyTypeCode_CODE,
	"CODING":                cpb.PropertyTypeCode_CODING,
	"STRING":                cpb.PropertyTypeCode_STRING,
	"INTEGER":               cpb.PropertyTypeCode_INTEGER,
	"BOOLEAN":               cpb.PropertyTypeCode_BOOLEAN,
	"DATE_TIME":             cpb.PropertyTypeCode_DATE_TIME,
	"DECIMAL":               cpb.PropertyTypeCode_DECIMAL,
}

// DefaultProvenanceEntityRoleCodeMap maps from string to cpb.ProvenanceEntityRoleCode_Value.
var DefaultProvenanceEntityRoleCodeMap = map[string]cpb.ProvenanceEntityRoleCode_Value{
	"INVALID_UNINITIALIZED": cpb.ProvenanceEntityRoleCode_INVALID_UNINITIALIZED,
	"DERIVATION":            cpb.ProvenanceEntityRoleCode_DERIVATION,
	"REVISION":              cpb.ProvenanceEntityRoleCode_REVISION,
	"QUOTATION":             cpb.ProvenanceEntityRoleCode_QUOTATION,
	"SOURCE":                cpb.ProvenanceEntityRoleCode_SOURCE,
	"REMOVAL":               cpb.ProvenanceEntityRoleCode_REMOVAL,
}

// DefaultPublicationStatusCodeMap maps from string to cpb.PublicationStatusCode_Value.
var DefaultPublicationStatusCodeMap = map[string]cpb.PublicationStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.PublicationStatusCode_INVALID_UNINITIALIZED,
	"DRAFT":                 cpb.PublicationStatusCode_DRAFT,
	"ACTIVE":                cpb.PublicationStatusCode_ACTIVE,
	"RETIRED":               cpb.PublicationStatusCode_RETIRED,
	"UNKNOWN":               cpb.PublicationStatusCode_UNKNOWN,
}

// DefaultQualityTypeCodeMap maps from string to cpb.QualityTypeCode_Value.
var DefaultQualityTypeCodeMap = map[string]cpb.QualityTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.QualityTypeCode_INVALID_UNINITIALIZED,
	"INDEL":                 cpb.QualityTypeCode_INDEL,
	"SNP":                   cpb.QualityTypeCode_SNP,
	"UNKNOWN":               cpb.QualityTypeCode_UNKNOWN,
}

// DefaultQuantityComparatorCodeMap maps from string to cpb.QuantityComparatorCode_Value.
var DefaultQuantityComparatorCodeMap = map[string]cpb.QuantityComparatorCode_Value{
	"INVALID_UNINITIALIZED":    cpb.QuantityComparatorCode_INVALID_UNINITIALIZED,
	"LESS_THAN":                cpb.QuantityComparatorCode_LESS_THAN,
	"LESS_THAN_OR_EQUAL_TO":    cpb.QuantityComparatorCode_LESS_THAN_OR_EQUAL_TO,
	"GREATER_THAN_OR_EQUAL_TO": cpb.QuantityComparatorCode_GREATER_THAN_OR_EQUAL_TO,
	"GREATER_THAN":             cpb.QuantityComparatorCode_GREATER_THAN,
}

// DefaultQuestionnaireItemOperatorCodeMap maps from string to cpb.QuestionnaireItemOperatorCode_Value.
var DefaultQuestionnaireItemOperatorCodeMap = map[string]cpb.QuestionnaireItemOperatorCode_Value{
	"INVALID_UNINITIALIZED":    cpb.QuestionnaireItemOperatorCode_INVALID_UNINITIALIZED,
	"EXISTS":                   cpb.QuestionnaireItemOperatorCode_EXISTS,
	"EQUALS":                   cpb.QuestionnaireItemOperatorCode_EQUALS,
	"NOT_EQUAL_TO":             cpb.QuestionnaireItemOperatorCode_NOT_EQUAL_TO,
	"GREATER_THAN":             cpb.QuestionnaireItemOperatorCode_GREATER_THAN,
	"LESS_THAN":                cpb.QuestionnaireItemOperatorCode_LESS_THAN,
	"GREATER_THAN_OR_EQUAL_TO": cpb.QuestionnaireItemOperatorCode_GREATER_THAN_OR_EQUAL_TO,
	"LESS_THAN_OR_EQUAL_TO":    cpb.QuestionnaireItemOperatorCode_LESS_THAN_OR_EQUAL_TO,
}

// DefaultQuestionnaireItemTypeCodeMap maps from string to cpb.QuestionnaireItemTypeCode_Value.
var DefaultQuestionnaireItemTypeCodeMap = map[string]cpb.QuestionnaireItemTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.QuestionnaireItemTypeCode_INVALID_UNINITIALIZED,
	"GROUP":                 cpb.QuestionnaireItemTypeCode_GROUP,
	"DISPLAY":               cpb.QuestionnaireItemTypeCode_DISPLAY,
	"QUESTION":              cpb.QuestionnaireItemTypeCode_QUESTION,
	"BOOLEAN":               cpb.QuestionnaireItemTypeCode_BOOLEAN,
	"DECIMAL":               cpb.QuestionnaireItemTypeCode_DECIMAL,
	"INTEGER":               cpb.QuestionnaireItemTypeCode_INTEGER,
	"DATE":                  cpb.QuestionnaireItemTypeCode_DATE,
	"DATE_TIME":             cpb.QuestionnaireItemTypeCode_DATE_TIME,
	"TIME":                  cpb.QuestionnaireItemTypeCode_TIME,
	"STRING":                cpb.QuestionnaireItemTypeCode_STRING,
	"TEXT":                  cpb.QuestionnaireItemTypeCode_TEXT,
	"URL":                   cpb.QuestionnaireItemTypeCode_URL,
	"CHOICE":                cpb.QuestionnaireItemTypeCode_CHOICE,
	"OPEN_CHOICE":           cpb.QuestionnaireItemTypeCode_OPEN_CHOICE,
	"ATTACHMENT":            cpb.QuestionnaireItemTypeCode_ATTACHMENT,
	"REFERENCE":             cpb.QuestionnaireItemTypeCode_REFERENCE,
	"QUANTITY":              cpb.QuestionnaireItemTypeCode_QUANTITY,
}

// DefaultQuestionnaireItemUsageModeCodeMap maps from string to cpb.QuestionnaireItemUsageModeCode_Value.
var DefaultQuestionnaireItemUsageModeCodeMap = map[string]cpb.QuestionnaireItemUsageModeCode_Value{
	"INVALID_UNINITIALIZED":     cpb.QuestionnaireItemUsageModeCode_INVALID_UNINITIALIZED,
	"CAPTURE_DISPLAY":           cpb.QuestionnaireItemUsageModeCode_CAPTURE_DISPLAY,
	"CAPTURE":                   cpb.QuestionnaireItemUsageModeCode_CAPTURE,
	"DISPLAY":                   cpb.QuestionnaireItemUsageModeCode_DISPLAY,
	"DISPLAY_NON_EMPTY":         cpb.QuestionnaireItemUsageModeCode_DISPLAY_NON_EMPTY,
	"CAPTURE_DISPLAY_NON_EMPTY": cpb.QuestionnaireItemUsageModeCode_CAPTURE_DISPLAY_NON_EMPTY,
}

// DefaultQuestionnaireResponseStatusCodeMap maps from string to cpb.QuestionnaireResponseStatusCode_Value.
var DefaultQuestionnaireResponseStatusCodeMap = map[string]cpb.QuestionnaireResponseStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.QuestionnaireResponseStatusCode_INVALID_UNINITIALIZED,
	"IN_PROGRESS":           cpb.QuestionnaireResponseStatusCode_IN_PROGRESS,
	"COMPLETED":             cpb.QuestionnaireResponseStatusCode_COMPLETED,
	"AMENDED":               cpb.QuestionnaireResponseStatusCode_AMENDED,
	"ENTERED_IN_ERROR":      cpb.QuestionnaireResponseStatusCode_ENTERED_IN_ERROR,
	"STOPPED":               cpb.QuestionnaireResponseStatusCode_STOPPED,
}

// DefaultReferenceHandlingPolicyCodeMap maps from string to cpb.ReferenceHandlingPolicyCode_Value.
var DefaultReferenceHandlingPolicyCodeMap = map[string]cpb.ReferenceHandlingPolicyCode_Value{
	"INVALID_UNINITIALIZED": cpb.ReferenceHandlingPolicyCode_INVALID_UNINITIALIZED,
	"LITERAL":               cpb.ReferenceHandlingPolicyCode_LITERAL,
	"LOGICAL":               cpb.ReferenceHandlingPolicyCode_LOGICAL,
	"RESOLVES":              cpb.ReferenceHandlingPolicyCode_RESOLVES,
	"ENFORCED":              cpb.ReferenceHandlingPolicyCode_ENFORCED,
	"LOCAL":                 cpb.ReferenceHandlingPolicyCode_LOCAL,
}

// DefaultReferenceVersionRulesCodeMap maps from string to cpb.ReferenceVersionRulesCode_Value.
var DefaultReferenceVersionRulesCodeMap = map[string]cpb.ReferenceVersionRulesCode_Value{
	"INVALID_UNINITIALIZED": cpb.ReferenceVersionRulesCode_INVALID_UNINITIALIZED,
	"EITHER":                cpb.ReferenceVersionRulesCode_EITHER,
	"INDEPENDENT":           cpb.ReferenceVersionRulesCode_INDEPENDENT,
	"SPECIFIC":              cpb.ReferenceVersionRulesCode_SPECIFIC,
}

// DefaultRelatedArtifactTypeCodeMap maps from string to cpb.RelatedArtifactTypeCode_Value.
var DefaultRelatedArtifactTypeCodeMap = map[string]cpb.RelatedArtifactTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.RelatedArtifactTypeCode_INVALID_UNINITIALIZED,
	"DOCUMENTATION":         cpb.RelatedArtifactTypeCode_DOCUMENTATION,
	"JUSTIFICATION":         cpb.RelatedArtifactTypeCode_JUSTIFICATION,
	"CITATION":              cpb.RelatedArtifactTypeCode_CITATION,
	"PREDECESSOR":           cpb.RelatedArtifactTypeCode_PREDECESSOR,
	"SUCCESSOR":             cpb.RelatedArtifactTypeCode_SUCCESSOR,
	"DERIVED_FROM":          cpb.RelatedArtifactTypeCode_DERIVED_FROM,
	"DEPENDS_ON":            cpb.RelatedArtifactTypeCode_DEPENDS_ON,
	"COMPOSED_OF":           cpb.RelatedArtifactTypeCode_COMPOSED_OF,
}

// DefaultRepositoryTypeCodeMap maps from string to cpb.RepositoryTypeCode_Value.
var DefaultRepositoryTypeCodeMap = map[string]cpb.RepositoryTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.RepositoryTypeCode_INVALID_UNINITIALIZED,
	"DIRECTLINK":            cpb.RepositoryTypeCode_DIRECTLINK,
	"OPENAPI":               cpb.RepositoryTypeCode_OPENAPI,
	"LOGIN":                 cpb.RepositoryTypeCode_LOGIN,
	"OAUTH":                 cpb.RepositoryTypeCode_OAUTH,
	"OTHER":                 cpb.RepositoryTypeCode_OTHER,
}

// DefaultRequestIntentCodeMap maps from string to cpb.RequestIntentCode_Value.
var DefaultRequestIntentCodeMap = map[string]cpb.RequestIntentCode_Value{
	"INVALID_UNINITIALIZED": cpb.RequestIntentCode_INVALID_UNINITIALIZED,
	"PROPOSAL":              cpb.RequestIntentCode_PROPOSAL,
	"PLAN":                  cpb.RequestIntentCode_PLAN,
	"DIRECTIVE":             cpb.RequestIntentCode_DIRECTIVE,
	"ORDER":                 cpb.RequestIntentCode_ORDER,
	"ORIGINAL_ORDER":        cpb.RequestIntentCode_ORIGINAL_ORDER,
	"REFLEX_ORDER":          cpb.RequestIntentCode_REFLEX_ORDER,
	"FILLER_ORDER":          cpb.RequestIntentCode_FILLER_ORDER,
	"INSTANCE_ORDER":        cpb.RequestIntentCode_INSTANCE_ORDER,
	"OPTION":                cpb.RequestIntentCode_OPTION,
}

// DefaultRequestPriorityCodeMap maps from string to cpb.RequestPriorityCode_Value.
var DefaultRequestPriorityCodeMap = map[string]cpb.RequestPriorityCode_Value{
	"INVALID_UNINITIALIZED": cpb.RequestPriorityCode_INVALID_UNINITIALIZED,
	"ROUTINE":               cpb.RequestPriorityCode_ROUTINE,
	"URGENT":                cpb.RequestPriorityCode_URGENT,
	"ASAP":                  cpb.RequestPriorityCode_ASAP,
	"STAT":                  cpb.RequestPriorityCode_STAT,
}

// DefaultRequestResourceTypeCodeMap maps from string to cpb.RequestResourceTypeCode_Value.
var DefaultRequestResourceTypeCodeMap = map[string]cpb.RequestResourceTypeCode_Value{
	"INVALID_UNINITIALIZED":       cpb.RequestResourceTypeCode_INVALID_UNINITIALIZED,
	"APPOINTMENT":                 cpb.RequestResourceTypeCode_APPOINTMENT,
	"APPOINTMENT_RESPONSE":        cpb.RequestResourceTypeCode_APPOINTMENT_RESPONSE,
	"CARE_PLAN":                   cpb.RequestResourceTypeCode_CARE_PLAN,
	"CLAIM":                       cpb.RequestResourceTypeCode_CLAIM,
	"COMMUNICATION_REQUEST":       cpb.RequestResourceTypeCode_COMMUNICATION_REQUEST,
	"CONTRACT":                    cpb.RequestResourceTypeCode_CONTRACT,
	"DEVICE_REQUEST":              cpb.RequestResourceTypeCode_DEVICE_REQUEST,
	"ENROLLMENT_REQUEST":          cpb.RequestResourceTypeCode_ENROLLMENT_REQUEST,
	"IMMUNIZATION_RECOMMENDATION": cpb.RequestResourceTypeCode_IMMUNIZATION_RECOMMENDATION,
	"MEDICATION_REQUEST":          cpb.RequestResourceTypeCode_MEDICATION_REQUEST,
	"NUTRITION_ORDER":             cpb.RequestResourceTypeCode_NUTRITION_ORDER,
	"SERVICE_REQUEST":             cpb.RequestResourceTypeCode_SERVICE_REQUEST,
	"SUPPLY_REQUEST":              cpb.RequestResourceTypeCode_SUPPLY_REQUEST,
	"TASK":                        cpb.RequestResourceTypeCode_TASK,
	"VISION_PRESCRIPTION":         cpb.RequestResourceTypeCode_VISION_PRESCRIPTION,
}

// DefaultRequestStatusCodeMap maps from string to cpb.RequestStatusCode_Value.
var DefaultRequestStatusCodeMap = map[string]cpb.RequestStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.RequestStatusCode_INVALID_UNINITIALIZED,
	"DRAFT":                 cpb.RequestStatusCode_DRAFT,
	"ACTIVE":                cpb.RequestStatusCode_ACTIVE,
	"ON_HOLD":               cpb.RequestStatusCode_ON_HOLD,
	"REVOKED":               cpb.RequestStatusCode_REVOKED,
	"COMPLETED":             cpb.RequestStatusCode_COMPLETED,
	"ENTERED_IN_ERROR":      cpb.RequestStatusCode_ENTERED_IN_ERROR,
	"UNKNOWN":               cpb.RequestStatusCode_UNKNOWN,
}

// DefaultResearchElementTypeCodeMap maps from string to cpb.ResearchElementTypeCode_Value.
var DefaultResearchElementTypeCodeMap = map[string]cpb.ResearchElementTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ResearchElementTypeCode_INVALID_UNINITIALIZED,
	"POPULATION":            cpb.ResearchElementTypeCode_POPULATION,
	"EXPOSURE":              cpb.ResearchElementTypeCode_EXPOSURE,
	"OUTCOME":               cpb.ResearchElementTypeCode_OUTCOME,
}

// DefaultResearchStudyStatusCodeMap maps from string to cpb.ResearchStudyStatusCode_Value.
var DefaultResearchStudyStatusCodeMap = map[string]cpb.ResearchStudyStatusCode_Value{
	"INVALID_UNINITIALIZED":              cpb.ResearchStudyStatusCode_INVALID_UNINITIALIZED,
	"ACTIVE":                             cpb.ResearchStudyStatusCode_ACTIVE,
	"ADMINISTRATIVELY_COMPLETED":         cpb.ResearchStudyStatusCode_ADMINISTRATIVELY_COMPLETED,
	"APPROVED":                           cpb.ResearchStudyStatusCode_APPROVED,
	"CLOSED_TO_ACCRUAL":                  cpb.ResearchStudyStatusCode_CLOSED_TO_ACCRUAL,
	"CLOSED_TO_ACCRUAL_AND_INTERVENTION": cpb.ResearchStudyStatusCode_CLOSED_TO_ACCRUAL_AND_INTERVENTION,
	"COMPLETED":                          cpb.ResearchStudyStatusCode_COMPLETED,
	"DISAPPROVED":                        cpb.ResearchStudyStatusCode_DISAPPROVED,
	"IN_REVIEW":                          cpb.ResearchStudyStatusCode_IN_REVIEW,
	"TEMPORARILY_CLOSED_TO_ACCRUAL":      cpb.ResearchStudyStatusCode_TEMPORARILY_CLOSED_TO_ACCRUAL,
	"TEMPORARILY_CLOSED_TO_ACCRUAL_AND_INTERVENTION": cpb.ResearchStudyStatusCode_TEMPORARILY_CLOSED_TO_ACCRUAL_AND_INTERVENTION,
	"WITHDRAWN": cpb.ResearchStudyStatusCode_WITHDRAWN,
}

// DefaultResearchSubjectStatusCodeMap maps from string to cpb.ResearchSubjectStatusCode_Value.
var DefaultResearchSubjectStatusCodeMap = map[string]cpb.ResearchSubjectStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.ResearchSubjectStatusCode_INVALID_UNINITIALIZED,
	"CANDIDATE":             cpb.ResearchSubjectStatusCode_CANDIDATE,
	"ELIGIBLE":              cpb.ResearchSubjectStatusCode_ELIGIBLE,
	"FOLLOW_UP":             cpb.ResearchSubjectStatusCode_FOLLOW_UP,
	"INELIGIBLE":            cpb.ResearchSubjectStatusCode_INELIGIBLE,
	"NOT_REGISTERED":        cpb.ResearchSubjectStatusCode_NOT_REGISTERED,
	"OFF_STUDY":             cpb.ResearchSubjectStatusCode_OFF_STUDY,
	"ON_STUDY":              cpb.ResearchSubjectStatusCode_ON_STUDY,
	"ON_STUDY_INTERVENTION": cpb.ResearchSubjectStatusCode_ON_STUDY_INTERVENTION,
	"ON_STUDY_OBSERVATION":  cpb.ResearchSubjectStatusCode_ON_STUDY_OBSERVATION,
	"PENDING_ON_STUDY":      cpb.ResearchSubjectStatusCode_PENDING_ON_STUDY,
	"POTENTIAL_CANDIDATE":   cpb.ResearchSubjectStatusCode_POTENTIAL_CANDIDATE,
	"SCREENING":             cpb.ResearchSubjectStatusCode_SCREENING,
	"WITHDRAWN":             cpb.ResearchSubjectStatusCode_WITHDRAWN,
}

// DefaultResourceSecurityCategoryCodeMap maps from string to cpb.ResourceSecurityCategoryCode_Value.
var DefaultResourceSecurityCategoryCodeMap = map[string]cpb.ResourceSecurityCategoryCode_Value{
	"INVALID_UNINITIALIZED": cpb.ResourceSecurityCategoryCode_INVALID_UNINITIALIZED,
	"ANONYMOUS":             cpb.ResourceSecurityCategoryCode_ANONYMOUS,
	"BUSINESS":              cpb.ResourceSecurityCategoryCode_BUSINESS,
	"INDIVIDUAL":            cpb.ResourceSecurityCategoryCode_INDIVIDUAL,
	"PATIENT":               cpb.ResourceSecurityCategoryCode_PATIENT,
	"NOT_CLASSIFIED":        cpb.ResourceSecurityCategoryCode_NOT_CLASSIFIED,
}

// DefaultResourceTypeCodeMap maps from string to cpb.ResourceTypeCode_Value.
var DefaultResourceTypeCodeMap = map[string]cpb.ResourceTypeCode_Value{
	"INVALID_UNINITIALIZED":                cpb.ResourceTypeCode_INVALID_UNINITIALIZED,
	"ACCOUNT":                              cpb.ResourceTypeCode_ACCOUNT,
	"ACTIVITY_DEFINITION":                  cpb.ResourceTypeCode_ACTIVITY_DEFINITION,
	"ADVERSE_EVENT":                        cpb.ResourceTypeCode_ADVERSE_EVENT,
	"ALLERGY_INTOLERANCE":                  cpb.ResourceTypeCode_ALLERGY_INTOLERANCE,
	"APPOINTMENT":                          cpb.ResourceTypeCode_APPOINTMENT,
	"APPOINTMENT_RESPONSE":                 cpb.ResourceTypeCode_APPOINTMENT_RESPONSE,
	"AUDIT_EVENT":                          cpb.ResourceTypeCode_AUDIT_EVENT,
	"BASIC":                                cpb.ResourceTypeCode_BASIC,
	"BINARY":                               cpb.ResourceTypeCode_BINARY,
	"BIOLOGICALLY_DERIVED_PRODUCT":         cpb.ResourceTypeCode_BIOLOGICALLY_DERIVED_PRODUCT,
	"BODY_STRUCTURE":                       cpb.ResourceTypeCode_BODY_STRUCTURE,
	"BUNDLE":                               cpb.ResourceTypeCode_BUNDLE,
	"CAPABILITY_STATEMENT":                 cpb.ResourceTypeCode_CAPABILITY_STATEMENT,
	"CARE_PLAN":                            cpb.ResourceTypeCode_CARE_PLAN,
	"CARE_TEAM":                            cpb.ResourceTypeCode_CARE_TEAM,
	"CATALOG_ENTRY":                        cpb.ResourceTypeCode_CATALOG_ENTRY,
	"CHARGE_ITEM":                          cpb.ResourceTypeCode_CHARGE_ITEM,
	"CHARGE_ITEM_DEFINITION":               cpb.ResourceTypeCode_CHARGE_ITEM_DEFINITION,
	"CLAIM":                                cpb.ResourceTypeCode_CLAIM,
	"CLAIM_RESPONSE":                       cpb.ResourceTypeCode_CLAIM_RESPONSE,
	"CLINICAL_IMPRESSION":                  cpb.ResourceTypeCode_CLINICAL_IMPRESSION,
	"CODE_SYSTEM":                          cpb.ResourceTypeCode_CODE_SYSTEM,
	"COMMUNICATION":                        cpb.ResourceTypeCode_COMMUNICATION,
	"COMMUNICATION_REQUEST":                cpb.ResourceTypeCode_COMMUNICATION_REQUEST,
	"COMPARTMENT_DEFINITION":               cpb.ResourceTypeCode_COMPARTMENT_DEFINITION,
	"COMPOSITION":                          cpb.ResourceTypeCode_COMPOSITION,
	"CONCEPT_MAP":                          cpb.ResourceTypeCode_CONCEPT_MAP,
	"CONDITION":                            cpb.ResourceTypeCode_CONDITION,
	"CONSENT":                              cpb.ResourceTypeCode_CONSENT,
	"CONTRACT":                             cpb.ResourceTypeCode_CONTRACT,
	"COVERAGE":                             cpb.ResourceTypeCode_COVERAGE,
	"COVERAGE_ELIGIBILITY_REQUEST":         cpb.ResourceTypeCode_COVERAGE_ELIGIBILITY_REQUEST,
	"COVERAGE_ELIGIBILITY_RESPONSE":        cpb.ResourceTypeCode_COVERAGE_ELIGIBILITY_RESPONSE,
	"DETECTED_ISSUE":                       cpb.ResourceTypeCode_DETECTED_ISSUE,
	"DEVICE":                               cpb.ResourceTypeCode_DEVICE,
	"DEVICE_DEFINITION":                    cpb.ResourceTypeCode_DEVICE_DEFINITION,
	"DEVICE_METRIC":                        cpb.ResourceTypeCode_DEVICE_METRIC,
	"DEVICE_REQUEST":                       cpb.ResourceTypeCode_DEVICE_REQUEST,
	"DEVICE_USE_STATEMENT":                 cpb.ResourceTypeCode_DEVICE_USE_STATEMENT,
	"DIAGNOSTIC_REPORT":                    cpb.ResourceTypeCode_DIAGNOSTIC_REPORT,
	"DOCUMENT_MANIFEST":                    cpb.ResourceTypeCode_DOCUMENT_MANIFEST,
	"DOCUMENT_REFERENCE":                   cpb.ResourceTypeCode_DOCUMENT_REFERENCE,
	"DOMAIN_RESOURCE":                      cpb.ResourceTypeCode_DOMAIN_RESOURCE,
	"EFFECT_EVIDENCE_SYNTHESIS":            cpb.ResourceTypeCode_EFFECT_EVIDENCE_SYNTHESIS,
	"ENCOUNTER":                            cpb.ResourceTypeCode_ENCOUNTER,
	"ENDPOINT":                             cpb.ResourceTypeCode_ENDPOINT,
	"ENROLLMENT_REQUEST":                   cpb.ResourceTypeCode_ENROLLMENT_REQUEST,
	"ENROLLMENT_RESPONSE":                  cpb.ResourceTypeCode_ENROLLMENT_RESPONSE,
	"EPISODE_OF_CARE":                      cpb.ResourceTypeCode_EPISODE_OF_CARE,
	"EVENT_DEFINITION":                     cpb.ResourceTypeCode_EVENT_DEFINITION,
	"EVIDENCE":                             cpb.ResourceTypeCode_EVIDENCE,
	"EVIDENCE_VARIABLE":                    cpb.ResourceTypeCode_EVIDENCE_VARIABLE,
	"EXAMPLE_SCENARIO":                     cpb.ResourceTypeCode_EXAMPLE_SCENARIO,
	"EXPLANATION_OF_BENEFIT":               cpb.ResourceTypeCode_EXPLANATION_OF_BENEFIT,
	"FAMILY_MEMBER_HISTORY":                cpb.ResourceTypeCode_FAMILY_MEMBER_HISTORY,
	"FLAG":                                 cpb.ResourceTypeCode_FLAG,
	"GOAL":                                 cpb.ResourceTypeCode_GOAL,
	"GRAPH_DEFINITION":                     cpb.ResourceTypeCode_GRAPH_DEFINITION,
	"GROUP":                                cpb.ResourceTypeCode_GROUP,
	"GUIDANCE_RESPONSE":                    cpb.ResourceTypeCode_GUIDANCE_RESPONSE,
	"HEALTHCARE_SERVICE":                   cpb.ResourceTypeCode_HEALTHCARE_SERVICE,
	"IMAGING_STUDY":                        cpb.ResourceTypeCode_IMAGING_STUDY,
	"IMMUNIZATION":                         cpb.ResourceTypeCode_IMMUNIZATION,
	"IMMUNIZATION_EVALUATION":              cpb.ResourceTypeCode_IMMUNIZATION_EVALUATION,
	"IMMUNIZATION_RECOMMENDATION":          cpb.ResourceTypeCode_IMMUNIZATION_RECOMMENDATION,
	"IMPLEMENTATION_GUIDE":                 cpb.ResourceTypeCode_IMPLEMENTATION_GUIDE,
	"INSURANCE_PLAN":                       cpb.ResourceTypeCode_INSURANCE_PLAN,
	"INVOICE":                              cpb.ResourceTypeCode_INVOICE,
	"LIBRARY":                              cpb.ResourceTypeCode_LIBRARY,
	"LINKAGE":                              cpb.ResourceTypeCode_LINKAGE,
	"LIST":                                 cpb.ResourceTypeCode_LIST,
	"LOCATION":                             cpb.ResourceTypeCode_LOCATION,
	"MEASURE":                              cpb.ResourceTypeCode_MEASURE,
	"MEASURE_REPORT":                       cpb.ResourceTypeCode_MEASURE_REPORT,
	"MEDIA":                                cpb.ResourceTypeCode_MEDIA,
	"MEDICATION":                           cpb.ResourceTypeCode_MEDICATION,
	"MEDICATION_ADMINISTRATION":            cpb.ResourceTypeCode_MEDICATION_ADMINISTRATION,
	"MEDICATION_DISPENSE":                  cpb.ResourceTypeCode_MEDICATION_DISPENSE,
	"MEDICATION_KNOWLEDGE":                 cpb.ResourceTypeCode_MEDICATION_KNOWLEDGE,
	"MEDICATION_REQUEST":                   cpb.ResourceTypeCode_MEDICATION_REQUEST,
	"MEDICATION_STATEMENT":                 cpb.ResourceTypeCode_MEDICATION_STATEMENT,
	"MEDICINAL_PRODUCT":                    cpb.ResourceTypeCode_MEDICINAL_PRODUCT,
	"MEDICINAL_PRODUCT_AUTHORIZATION":      cpb.ResourceTypeCode_MEDICINAL_PRODUCT_AUTHORIZATION,
	"MEDICINAL_PRODUCT_CONTRAINDICATION":   cpb.ResourceTypeCode_MEDICINAL_PRODUCT_CONTRAINDICATION,
	"MEDICINAL_PRODUCT_INDICATION":         cpb.ResourceTypeCode_MEDICINAL_PRODUCT_INDICATION,
	"MEDICINAL_PRODUCT_INGREDIENT":         cpb.ResourceTypeCode_MEDICINAL_PRODUCT_INGREDIENT,
	"MEDICINAL_PRODUCT_INTERACTION":        cpb.ResourceTypeCode_MEDICINAL_PRODUCT_INTERACTION,
	"MEDICINAL_PRODUCT_MANUFACTURED":       cpb.ResourceTypeCode_MEDICINAL_PRODUCT_MANUFACTURED,
	"MEDICINAL_PRODUCT_PACKAGED":           cpb.ResourceTypeCode_MEDICINAL_PRODUCT_PACKAGED,
	"MEDICINAL_PRODUCT_PHARMACEUTICAL":     cpb.ResourceTypeCode_MEDICINAL_PRODUCT_PHARMACEUTICAL,
	"MEDICINAL_PRODUCT_UNDESIRABLE_EFFECT": cpb.ResourceTypeCode_MEDICINAL_PRODUCT_UNDESIRABLE_EFFECT,
	"MESSAGE_DEFINITION":                   cpb.ResourceTypeCode_MESSAGE_DEFINITION,
	"MESSAGE_HEADER":                       cpb.ResourceTypeCode_MESSAGE_HEADER,
	"MOLECULAR_SEQUENCE":                   cpb.ResourceTypeCode_MOLECULAR_SEQUENCE,
	"NAMING_SYSTEM":                        cpb.ResourceTypeCode_NAMING_SYSTEM,
	"NUTRITION_ORDER":                      cpb.ResourceTypeCode_NUTRITION_ORDER,
	"OBSERVATION":                          cpb.ResourceTypeCode_OBSERVATION,
	"OBSERVATION_DEFINITION":               cpb.ResourceTypeCode_OBSERVATION_DEFINITION,
	"OPERATION_DEFINITION":                 cpb.ResourceTypeCode_OPERATION_DEFINITION,
	"OPERATION_OUTCOME":                    cpb.ResourceTypeCode_OPERATION_OUTCOME,
	"ORGANIZATION":                         cpb.ResourceTypeCode_ORGANIZATION,
	"ORGANIZATION_AFFILIATION":             cpb.ResourceTypeCode_ORGANIZATION_AFFILIATION,
	"PARAMETERS":                           cpb.ResourceTypeCode_PARAMETERS,
	"PATIENT":                              cpb.ResourceTypeCode_PATIENT,
	"PAYMENT_NOTICE":                       cpb.ResourceTypeCode_PAYMENT_NOTICE,
	"PAYMENT_RECONCILIATION":               cpb.ResourceTypeCode_PAYMENT_RECONCILIATION,
	"PERSON":                               cpb.ResourceTypeCode_PERSON,
	"PLAN_DEFINITION":                      cpb.ResourceTypeCode_PLAN_DEFINITION,
	"PRACTITIONER":                         cpb.ResourceTypeCode_PRACTITIONER,
	"PRACTITIONER_ROLE":                    cpb.ResourceTypeCode_PRACTITIONER_ROLE,
	"PROCEDURE":                            cpb.ResourceTypeCode_PROCEDURE,
	"PROVENANCE":                           cpb.ResourceTypeCode_PROVENANCE,
	"QUESTIONNAIRE":                        cpb.ResourceTypeCode_QUESTIONNAIRE,
	"QUESTIONNAIRE_RESPONSE":               cpb.ResourceTypeCode_QUESTIONNAIRE_RESPONSE,
	"RELATED_PERSON":                       cpb.ResourceTypeCode_RELATED_PERSON,
	"REQUEST_GROUP":                        cpb.ResourceTypeCode_REQUEST_GROUP,
	"RESEARCH_DEFINITION":                  cpb.ResourceTypeCode_RESEARCH_DEFINITION,
	"RESEARCH_ELEMENT_DEFINITION":          cpb.ResourceTypeCode_RESEARCH_ELEMENT_DEFINITION,
	"RESEARCH_STUDY":                       cpb.ResourceTypeCode_RESEARCH_STUDY,
	"RESEARCH_SUBJECT":                     cpb.ResourceTypeCode_RESEARCH_SUBJECT,
	"RESOURCE":                             cpb.ResourceTypeCode_RESOURCE,
	"RISK_ASSESSMENT":                      cpb.ResourceTypeCode_RISK_ASSESSMENT,
	"RISK_EVIDENCE_SYNTHESIS":              cpb.ResourceTypeCode_RISK_EVIDENCE_SYNTHESIS,
	"SCHEDULE":                             cpb.ResourceTypeCode_SCHEDULE,
	"SEARCH_PARAMETER":                     cpb.ResourceTypeCode_SEARCH_PARAMETER,
	"SERVICE_REQUEST":                      cpb.ResourceTypeCode_SERVICE_REQUEST,
	"SLOT":                                 cpb.ResourceTypeCode_SLOT,
	"SPECIMEN":                             cpb.ResourceTypeCode_SPECIMEN,
	"SPECIMEN_DEFINITION":                  cpb.ResourceTypeCode_SPECIMEN_DEFINITION,
	"STRUCTURE_DEFINITION":                 cpb.ResourceTypeCode_STRUCTURE_DEFINITION,
	"STRUCTURE_MAP":                        cpb.ResourceTypeCode_STRUCTURE_MAP,
	"SUBSCRIPTION":                         cpb.ResourceTypeCode_SUBSCRIPTION,
	"SUBSTANCE":                            cpb.ResourceTypeCode_SUBSTANCE,
	"SUBSTANCE_NUCLEIC_ACID":               cpb.ResourceTypeCode_SUBSTANCE_NUCLEIC_ACID,
	"SUBSTANCE_POLYMER":                    cpb.ResourceTypeCode_SUBSTANCE_POLYMER,
	"SUBSTANCE_PROTEIN":                    cpb.ResourceTypeCode_SUBSTANCE_PROTEIN,
	"SUBSTANCE_REFERENCE_INFORMATION":      cpb.ResourceTypeCode_SUBSTANCE_REFERENCE_INFORMATION,
	"SUBSTANCE_SOURCE_MATERIAL":            cpb.ResourceTypeCode_SUBSTANCE_SOURCE_MATERIAL,
	"SUBSTANCE_SPECIFICATION":              cpb.ResourceTypeCode_SUBSTANCE_SPECIFICATION,
	"SUPPLY_DELIVERY":                      cpb.ResourceTypeCode_SUPPLY_DELIVERY,
	"SUPPLY_REQUEST":                       cpb.ResourceTypeCode_SUPPLY_REQUEST,
	"TASK":                                 cpb.ResourceTypeCode_TASK,
	"TERMINOLOGY_CAPABILITIES":             cpb.ResourceTypeCode_TERMINOLOGY_CAPABILITIES,
	"TEST_REPORT":                          cpb.ResourceTypeCode_TEST_REPORT,
	"TEST_SCRIPT":                          cpb.ResourceTypeCode_TEST_SCRIPT,
	"VALUE_SET":                            cpb.ResourceTypeCode_VALUE_SET,
	"VERIFICATION_RESULT":                  cpb.ResourceTypeCode_VERIFICATION_RESULT,
	"VISION_PRESCRIPTION":                  cpb.ResourceTypeCode_VISION_PRESCRIPTION,
}

// DefaultResourceVersionPolicyCodeMap maps from string to cpb.ResourceVersionPolicyCode_Value.
var DefaultResourceVersionPolicyCodeMap = map[string]cpb.ResourceVersionPolicyCode_Value{
	"INVALID_UNINITIALIZED": cpb.ResourceVersionPolicyCode_INVALID_UNINITIALIZED,
	"NO_VERSION":            cpb.ResourceVersionPolicyCode_NO_VERSION,
	"VERSIONED":             cpb.ResourceVersionPolicyCode_VERSIONED,
	"VERSIONED_UPDATE":      cpb.ResourceVersionPolicyCode_VERSIONED_UPDATE,
}

// DefaultResponseTypeCodeMap maps from string to cpb.ResponseTypeCode_Value.
var DefaultResponseTypeCodeMap = map[string]cpb.ResponseTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.ResponseTypeCode_INVALID_UNINITIALIZED,
	"OK":                    cpb.ResponseTypeCode_OK,
	"TRANSIENT_ERROR":       cpb.ResponseTypeCode_TRANSIENT_ERROR,
	"FATAL_ERROR":           cpb.ResponseTypeCode_FATAL_ERROR,
}

// DefaultRestfulCapabilityModeCodeMap maps from string to cpb.RestfulCapabilityModeCode_Value.
var DefaultRestfulCapabilityModeCodeMap = map[string]cpb.RestfulCapabilityModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.RestfulCapabilityModeCode_INVALID_UNINITIALIZED,
	"CLIENT":                cpb.RestfulCapabilityModeCode_CLIENT,
	"SERVER":                cpb.RestfulCapabilityModeCode_SERVER,
}

// DefaultSPDXLicenseCodeMap maps from string to cpb.SPDXLicenseCode_Value.
var DefaultSPDXLicenseCodeMap = map[string]cpb.SPDXLicenseCode_Value{
	"INVALID_UNINITIALIZED":                cpb.SPDXLicenseCode_INVALID_UNINITIALIZED,
	"NOT_OPEN_SOURCE":                      cpb.SPDXLicenseCode_NOT_OPEN_SOURCE,
	"BSD_ZERO_CLAUSE_LICENSE":              cpb.SPDXLicenseCode_BSD_ZERO_CLAUSE_LICENSE,
	"AAL":                                  cpb.SPDXLicenseCode_AAL,
	"ABSTYLES":                             cpb.SPDXLicenseCode_ABSTYLES,
	"ADOBE_2006":                           cpb.SPDXLicenseCode_ADOBE_2006,
	"ADOBE_GLYPH":                          cpb.SPDXLicenseCode_ADOBE_GLYPH,
	"ADSL":                                 cpb.SPDXLicenseCode_ADSL,
	"AFL_1_1":                              cpb.SPDXLicenseCode_AFL_1_1,
	"AFL_1_2":                              cpb.SPDXLicenseCode_AFL_1_2,
	"AFL_2_0":                              cpb.SPDXLicenseCode_AFL_2_0,
	"AFL_2_1":                              cpb.SPDXLicenseCode_AFL_2_1,
	"AFL_3_0":                              cpb.SPDXLicenseCode_AFL_3_0,
	"AFMPARSE":                             cpb.SPDXLicenseCode_AFMPARSE,
	"AGPL_1_0_ONLY":                        cpb.SPDXLicenseCode_AGPL_1_0_ONLY,
	"AGPL_1_0_OR_LATER":                    cpb.SPDXLicenseCode_AGPL_1_0_OR_LATER,
	"AGPL_3_0_ONLY":                        cpb.SPDXLicenseCode_AGPL_3_0_ONLY,
	"AGPL_3_0_OR_LATER":                    cpb.SPDXLicenseCode_AGPL_3_0_OR_LATER,
	"ALADDIN":                              cpb.SPDXLicenseCode_ALADDIN,
	"AMDPLPA":                              cpb.SPDXLicenseCode_AMDPLPA,
	"AML":                                  cpb.SPDXLicenseCode_AML,
	"AMPAS":                                cpb.SPDXLicenseCode_AMPAS,
	"ANTLR_PD":                             cpb.SPDXLicenseCode_ANTLR_PD,
	"APACHE_1_0":                           cpb.SPDXLicenseCode_APACHE_1_0,
	"APACHE_1_1":                           cpb.SPDXLicenseCode_APACHE_1_1,
	"APACHE_2_0":                           cpb.SPDXLicenseCode_APACHE_2_0,
	"APAFML":                               cpb.SPDXLicenseCode_APAFML,
	"APL_1_0":                              cpb.SPDXLicenseCode_APL_1_0,
	"APSL_1_0":                             cpb.SPDXLicenseCode_APSL_1_0,
	"APSL_1_1":                             cpb.SPDXLicenseCode_APSL_1_1,
	"APSL_1_2":                             cpb.SPDXLicenseCode_APSL_1_2,
	"APSL_2_0":                             cpb.SPDXLicenseCode_APSL_2_0,
	"ARTISTIC_1_0_CL8":                     cpb.SPDXLicenseCode_ARTISTIC_1_0_CL8,
	"ARTISTIC_1_0_PERL":                    cpb.SPDXLicenseCode_ARTISTIC_1_0_PERL,
	"ARTISTIC_1_0":                         cpb.SPDXLicenseCode_ARTISTIC_1_0,
	"ARTISTIC_2_0":                         cpb.SPDXLicenseCode_ARTISTIC_2_0,
	"BAHYPH":                               cpb.SPDXLicenseCode_BAHYPH,
	"BARR":                                 cpb.SPDXLicenseCode_BARR,
	"BEERWARE":                             cpb.SPDXLicenseCode_BEERWARE,
	"BIT_TORRENT_1_0":                      cpb.SPDXLicenseCode_BIT_TORRENT_1_0,
	"BIT_TORRENT_1_1":                      cpb.SPDXLicenseCode_BIT_TORRENT_1_1,
	"BORCEUX":                              cpb.SPDXLicenseCode_BORCEUX,
	"BSD_1_CLAUSE":                         cpb.SPDXLicenseCode_BSD_1_CLAUSE,
	"BSD_2_CLAUSE_FREE_BSD":                cpb.SPDXLicenseCode_BSD_2_CLAUSE_FREE_BSD,
	"BSD_2_CLAUSE_NET_BSD":                 cpb.SPDXLicenseCode_BSD_2_CLAUSE_NET_BSD,
	"BSD_2_CLAUSE_PATENT":                  cpb.SPDXLicenseCode_BSD_2_CLAUSE_PATENT,
	"BSD_2_CLAUSE":                         cpb.SPDXLicenseCode_BSD_2_CLAUSE,
	"BSD_3_CLAUSE_ATTRIBUTION":             cpb.SPDXLicenseCode_BSD_3_CLAUSE_ATTRIBUTION,
	"BSD_3_CLAUSE_CLEAR":                   cpb.SPDXLicenseCode_BSD_3_CLAUSE_CLEAR,
	"BSD_3_CLAUSE_LBNL":                    cpb.SPDXLicenseCode_BSD_3_CLAUSE_LBNL,
	"BSD_3_CLAUSE_NO_NUCLEAR_LICENSE_2014": cpb.SPDXLicenseCode_BSD_3_CLAUSE_NO_NUCLEAR_LICENSE_2014,
	"BSD_3_CLAUSE_NO_NUCLEAR_LICENSE":      cpb.SPDXLicenseCode_BSD_3_CLAUSE_NO_NUCLEAR_LICENSE,
	"BSD_3_CLAUSE_NO_NUCLEAR_WARRANTY":     cpb.SPDXLicenseCode_BSD_3_CLAUSE_NO_NUCLEAR_WARRANTY,
	"BSD_3_CLAUSE":                         cpb.SPDXLicenseCode_BSD_3_CLAUSE,
	"BSD_4_CLAUSE_UC":                      cpb.SPDXLicenseCode_BSD_4_CLAUSE_UC,
	"BSD_4_CLAUSE":                         cpb.SPDXLicenseCode_BSD_4_CLAUSE,
	"BSD_PROTECTION":                       cpb.SPDXLicenseCode_BSD_PROTECTION,
	"BSD_SOURCE_CODE":                      cpb.SPDXLicenseCode_BSD_SOURCE_CODE,
	"BSL_1_0":                              cpb.SPDXLicenseCode_BSL_1_0,
	"BZIP2_1_0_5":                          cpb.SPDXLicenseCode_BZIP2_1_0_5,
	"BZIP2_1_0_6":                          cpb.SPDXLicenseCode_BZIP2_1_0_6,
	"CALDERA":                              cpb.SPDXLicenseCode_CALDERA,
	"CATOSL_1_1":                           cpb.SPDXLicenseCode_CATOSL_1_1,
	"CC_BY_1_0":                            cpb.SPDXLicenseCode_CC_BY_1_0,
	"CC_BY_2_0":                            cpb.SPDXLicenseCode_CC_BY_2_0,
	"CC_BY_2_5":                            cpb.SPDXLicenseCode_CC_BY_2_5,
	"CC_BY_3_0":                            cpb.SPDXLicenseCode_CC_BY_3_0,
	"CC_BY_4_0":                            cpb.SPDXLicenseCode_CC_BY_4_0,
	"CC_BY_NC_1_0":                         cpb.SPDXLicenseCode_CC_BY_NC_1_0,
	"CC_BY_NC_2_0":                         cpb.SPDXLicenseCode_CC_BY_NC_2_0,
	"CC_BY_NC_2_5":                         cpb.SPDXLicenseCode_CC_BY_NC_2_5,
	"CC_BY_NC_3_0":                         cpb.SPDXLicenseCode_CC_BY_NC_3_0,
	"CC_BY_NC_4_0":                         cpb.SPDXLicenseCode_CC_BY_NC_4_0,
	"CC_BY_NC_ND_1_0":                      cpb.SPDXLicenseCode_CC_BY_NC_ND_1_0,
	"CC_BY_NC_ND_2_0":                      cpb.SPDXLicenseCode_CC_BY_NC_ND_2_0,
	"CC_BY_NC_ND_2_5":                      cpb.SPDXLicenseCode_CC_BY_NC_ND_2_5,
	"CC_BY_NC_ND_3_0":                      cpb.SPDXLicenseCode_CC_BY_NC_ND_3_0,
	"CC_BY_NC_ND_4_0":                      cpb.SPDXLicenseCode_CC_BY_NC_ND_4_0,
	"CC_BY_NC_SA_1_0":                      cpb.SPDXLicenseCode_CC_BY_NC_SA_1_0,
	"CC_BY_NC_SA_2_0":                      cpb.SPDXLicenseCode_CC_BY_NC_SA_2_0,
	"CC_BY_NC_SA_2_5":                      cpb.SPDXLicenseCode_CC_BY_NC_SA_2_5,
	"CC_BY_NC_SA_3_0":                      cpb.SPDXLicenseCode_CC_BY_NC_SA_3_0,
	"CC_BY_NC_SA_4_0":                      cpb.SPDXLicenseCode_CC_BY_NC_SA_4_0,
	"CC_BY_ND_1_0":                         cpb.SPDXLicenseCode_CC_BY_ND_1_0,
	"CC_BY_ND_2_0":                         cpb.SPDXLicenseCode_CC_BY_ND_2_0,
	"CC_BY_ND_2_5":                         cpb.SPDXLicenseCode_CC_BY_ND_2_5,
	"CC_BY_ND_3_0":                         cpb.SPDXLicenseCode_CC_BY_ND_3_0,
	"CC_BY_ND_4_0":                         cpb.SPDXLicenseCode_CC_BY_ND_4_0,
	"CC_BY_SA_1_0":                         cpb.SPDXLicenseCode_CC_BY_SA_1_0,
	"CC_BY_SA_2_0":                         cpb.SPDXLicenseCode_CC_BY_SA_2_0,
	"CC_BY_SA_2_5":                         cpb.SPDXLicenseCode_CC_BY_SA_2_5,
	"CC_BY_SA_3_0":                         cpb.SPDXLicenseCode_CC_BY_SA_3_0,
	"CC_BY_SA_4_0":                         cpb.SPDXLicenseCode_CC_BY_SA_4_0,
	"CC0_1_0":                              cpb.SPDXLicenseCode_CC0_1_0,
	"CDDL_1_0":                             cpb.SPDXLicenseCode_CDDL_1_0,
	"CDDL_1_1":                             cpb.SPDXLicenseCode_CDDL_1_1,
	"CDLA_PERMISSIVE_1_0":                  cpb.SPDXLicenseCode_CDLA_PERMISSIVE_1_0,
	"CDLA_SHARING_1_0":                     cpb.SPDXLicenseCode_CDLA_SHARING_1_0,
	"CECILL_1_0":                           cpb.SPDXLicenseCode_CECILL_1_0,
	"CECILL_1_1":                           cpb.SPDXLicenseCode_CECILL_1_1,
	"CECILL_2_0":                           cpb.SPDXLicenseCode_CECILL_2_0,
	"CECILL_2_1":                           cpb.SPDXLicenseCode_CECILL_2_1,
	"CECILL_B":                             cpb.SPDXLicenseCode_CECILL_B,
	"CECILL_C":                             cpb.SPDXLicenseCode_CECILL_C,
	"CL_ARTISTIC":                          cpb.SPDXLicenseCode_CL_ARTISTIC,
	"CNRI_JYTHON":                          cpb.SPDXLicenseCode_CNRI_JYTHON,
	"CNRI_PYTHON_GPL_COMPATIBLE":           cpb.SPDXLicenseCode_CNRI_PYTHON_GPL_COMPATIBLE,
	"CNRI_PYTHON":                          cpb.SPDXLicenseCode_CNRI_PYTHON,
	"CONDOR_1_1":                           cpb.SPDXLicenseCode_CONDOR_1_1,
	"CPAL_1_0":                             cpb.SPDXLicenseCode_CPAL_1_0,
	"CPL_1_0":                              cpb.SPDXLicenseCode_CPL_1_0,
	"CPOL_1_02":                            cpb.SPDXLicenseCode_CPOL_1_02,
	"CROSSWORD":                            cpb.SPDXLicenseCode_CROSSWORD,
	"CRYSTAL_STACKER":                      cpb.SPDXLicenseCode_CRYSTAL_STACKER,
	"CUA_OPL_1_0":                          cpb.SPDXLicenseCode_CUA_OPL_1_0,
	"CUBE":                                 cpb.SPDXLicenseCode_CUBE,
	"CURL":                                 cpb.SPDXLicenseCode_CURL,
	"D_FSL_1_0":                            cpb.SPDXLicenseCode_D_FSL_1_0,
	"DIFFMARK":                             cpb.SPDXLicenseCode_DIFFMARK,
	"DOC":                                  cpb.SPDXLicenseCode_DOC,
	"DOTSEQN":                              cpb.SPDXLicenseCode_DOTSEQN,
	"DSDP":                                 cpb.SPDXLicenseCode_DSDP,
	"DVIPDFM":                              cpb.SPDXLicenseCode_DVIPDFM,
	"ECL_1_0":                              cpb.SPDXLicenseCode_ECL_1_0,
	"ECL_2_0":                              cpb.SPDXLicenseCode_ECL_2_0,
	"EFL_1_0":                              cpb.SPDXLicenseCode_EFL_1_0,
	"EFL_2_0":                              cpb.SPDXLicenseCode_EFL_2_0,
	"E_GENIX":                              cpb.SPDXLicenseCode_E_GENIX,
	"ENTESSA":                              cpb.SPDXLicenseCode_ENTESSA,
	"EPL_1_0":                              cpb.SPDXLicenseCode_EPL_1_0,
	"EPL_2_0":                              cpb.SPDXLicenseCode_EPL_2_0,
	"ERL_PL_1_1":                           cpb.SPDXLicenseCode_ERL_PL_1_1,
	"EU_DATAGRID":                          cpb.SPDXLicenseCode_EU_DATAGRID,
	"EUPL_1_0":                             cpb.SPDXLicenseCode_EUPL_1_0,
	"EUPL_1_1":                             cpb.SPDXLicenseCode_EUPL_1_1,
	"EUPL_1_2":                             cpb.SPDXLicenseCode_EUPL_1_2,
	"EUROSYM":                              cpb.SPDXLicenseCode_EUROSYM,
	"FAIR":                                 cpb.SPDXLicenseCode_FAIR,
	"FRAMEWORX_1_0":                        cpb.SPDXLicenseCode_FRAMEWORX_1_0,
	"FREE_IMAGE":                           cpb.SPDXLicenseCode_FREE_IMAGE,
	"FSFAP":                                cpb.SPDXLicenseCode_FSFAP,
	"FSFUL":                                cpb.SPDXLicenseCode_FSFUL,
	"FSFULLR":                              cpb.SPDXLicenseCode_FSFULLR,
	"FTL":                                  cpb.SPDXLicenseCode_FTL,
	"GFDL_1_1_ONLY":                        cpb.SPDXLicenseCode_GFDL_1_1_ONLY,
	"GFDL_1_1_OR_LATER":                    cpb.SPDXLicenseCode_GFDL_1_1_OR_LATER,
	"GFDL_1_2_ONLY":                        cpb.SPDXLicenseCode_GFDL_1_2_ONLY,
	"GFDL_1_2_OR_LATER":                    cpb.SPDXLicenseCode_GFDL_1_2_OR_LATER,
	"GFDL_1_3_ONLY":                        cpb.SPDXLicenseCode_GFDL_1_3_ONLY,
	"GFDL_1_3_OR_LATER":                    cpb.SPDXLicenseCode_GFDL_1_3_OR_LATER,
	"GIFTWARE":                             cpb.SPDXLicenseCode_GIFTWARE,
	"GL2PS":                                cpb.SPDXLicenseCode_GL2PS,
	"GLIDE":                                cpb.SPDXLicenseCode_GLIDE,
	"GLULXE":                               cpb.SPDXLicenseCode_GLULXE,
	"GNUPLOT":                              cpb.SPDXLicenseCode_GNUPLOT,
	"GPL_1_0_ONLY":                         cpb.SPDXLicenseCode_GPL_1_0_ONLY,
	"GPL_1_0_OR_LATER":                     cpb.SPDXLicenseCode_GPL_1_0_OR_LATER,
	"GPL_2_0_ONLY":                         cpb.SPDXLicenseCode_GPL_2_0_ONLY,
	"GPL_2_0_OR_LATER":                     cpb.SPDXLicenseCode_GPL_2_0_OR_LATER,
	"GPL_3_0_ONLY":                         cpb.SPDXLicenseCode_GPL_3_0_ONLY,
	"GPL_3_0_OR_LATER":                     cpb.SPDXLicenseCode_GPL_3_0_OR_LATER,
	"G_SOAP_1_3B":                          cpb.SPDXLicenseCode_G_SOAP_1_3B,
	"HASKELL_REPORT":                       cpb.SPDXLicenseCode_HASKELL_REPORT,
	"HPND":                                 cpb.SPDXLicenseCode_HPND,
	"IBM_PIBS":                             cpb.SPDXLicenseCode_IBM_PIBS,
	"ICU":                                  cpb.SPDXLicenseCode_ICU,
	"IJG":                                  cpb.SPDXLicenseCode_IJG,
	"IMAGE_MAGICK":                         cpb.SPDXLicenseCode_IMAGE_MAGICK,
	"I_MATIX":                              cpb.SPDXLicenseCode_I_MATIX,
	"IMLIB2":                               cpb.SPDXLicenseCode_IMLIB2,
	"INFO_ZIP":                             cpb.SPDXLicenseCode_INFO_ZIP,
	"INTEL_ACPI":                           cpb.SPDXLicenseCode_INTEL_ACPI,
	"INTEL":                                cpb.SPDXLicenseCode_INTEL,
	"INTERBASE_1_0":                        cpb.SPDXLicenseCode_INTERBASE_1_0,
	"IPA":                                  cpb.SPDXLicenseCode_IPA,
	"IPL_1_0":                              cpb.SPDXLicenseCode_IPL_1_0,
	"ISC":                                  cpb.SPDXLicenseCode_ISC,
	"JAS_PER_2_0":                          cpb.SPDXLicenseCode_JAS_PER_2_0,
	"JSON":                                 cpb.SPDXLicenseCode_JSON,
	"LAL_1_2":                              cpb.SPDXLicenseCode_LAL_1_2,
	"LAL_1_3":                              cpb.SPDXLicenseCode_LAL_1_3,
	"LATEX2E":                              cpb.SPDXLicenseCode_LATEX2E,
	"LEPTONICA":                            cpb.SPDXLicenseCode_LEPTONICA,
	"LGPL_2_0_ONLY":                        cpb.SPDXLicenseCode_LGPL_2_0_ONLY,
	"LGPL_2_0_OR_LATER":                    cpb.SPDXLicenseCode_LGPL_2_0_OR_LATER,
	"LGPL_2_1_ONLY":                        cpb.SPDXLicenseCode_LGPL_2_1_ONLY,
	"LGPL_2_1_OR_LATER":                    cpb.SPDXLicenseCode_LGPL_2_1_OR_LATER,
	"LGPL_3_0_ONLY":                        cpb.SPDXLicenseCode_LGPL_3_0_ONLY,
	"LGPL_3_0_OR_LATER":                    cpb.SPDXLicenseCode_LGPL_3_0_OR_LATER,
	"LGPLLR":                               cpb.SPDXLicenseCode_LGPLLR,
	"LIBPNG":                               cpb.SPDXLicenseCode_LIBPNG,
	"LIBTIFF":                              cpb.SPDXLicenseCode_LIBTIFF,
	"LI_LI_Q_P_1_1":                        cpb.SPDXLicenseCode_LI_LI_Q_P_1_1,
	"LI_LI_Q_R_1_1":                        cpb.SPDXLicenseCode_LI_LI_Q_R_1_1,
	"LI_LI_Q_RPLUS_1_1":                    cpb.SPDXLicenseCode_LI_LI_Q_RPLUS_1_1,
	"LINUX_OPEN_IB":                        cpb.SPDXLicenseCode_LINUX_OPEN_IB,
	"LPL_1_0":                              cpb.SPDXLicenseCode_LPL_1_0,
	"LPL_1_02":                             cpb.SPDXLicenseCode_LPL_1_02,
	"LPPL_1_0":                             cpb.SPDXLicenseCode_LPPL_1_0,
	"LPPL_1_1":                             cpb.SPDXLicenseCode_LPPL_1_1,
	"LPPL_1_2":                             cpb.SPDXLicenseCode_LPPL_1_2,
	"LPPL_1_3A":                            cpb.SPDXLicenseCode_LPPL_1_3A,
	"LPPL_1_3C":                            cpb.SPDXLicenseCode_LPPL_1_3C,
	"MAKE_INDEX":                           cpb.SPDXLicenseCode_MAKE_INDEX,
	"MIR_OS":                               cpb.SPDXLicenseCode_MIR_OS,
	"MIT_0":                                cpb.SPDXLicenseCode_MIT_0,
	"MIT_ADVERTISING":                      cpb.SPDXLicenseCode_MIT_ADVERTISING,
	"MIT_CMU":                              cpb.SPDXLicenseCode_MIT_CMU,
	"MIT_ENNA":                             cpb.SPDXLicenseCode_MIT_ENNA,
	"MIT_FEH":                              cpb.SPDXLicenseCode_MIT_FEH,
	"MIT":                                  cpb.SPDXLicenseCode_MIT,
	"MITNFA":                               cpb.SPDXLicenseCode_MITNFA,
	"MOTOSOTO":                             cpb.SPDXLicenseCode_MOTOSOTO,
	"MPICH2":                               cpb.SPDXLicenseCode_MPICH2,
	"MPL_1_0":                              cpb.SPDXLicenseCode_MPL_1_0,
	"MPL_1_1":                              cpb.SPDXLicenseCode_MPL_1_1,
	"MPL_2_0_NO_COPYLEFT_EXCEPTION":        cpb.SPDXLicenseCode_MPL_2_0_NO_COPYLEFT_EXCEPTION,
	"MPL_2_0":                              cpb.SPDXLicenseCode_MPL_2_0,
	"MS_PL":                                cpb.SPDXLicenseCode_MS_PL,
	"MS_RL":                                cpb.SPDXLicenseCode_MS_RL,
	"MTLL":                                 cpb.SPDXLicenseCode_MTLL,
	"MULTICS":                              cpb.SPDXLicenseCode_MULTICS,
	"MUP":                                  cpb.SPDXLicenseCode_MUP,
	"NASA_1_3":                             cpb.SPDXLicenseCode_NASA_1_3,
	"NAUMEN":                               cpb.SPDXLicenseCode_NAUMEN,
	"NBPL_1_0":                             cpb.SPDXLicenseCode_NBPL_1_0,
	"NCSA":                                 cpb.SPDXLicenseCode_NCSA,
	"NET_SNMP":                             cpb.SPDXLicenseCode_NET_SNMP,
	"NET_CDF":                              cpb.SPDXLicenseCode_NET_CDF,
	"NEWSLETR":                             cpb.SPDXLicenseCode_NEWSLETR,
	"NGPL":                                 cpb.SPDXLicenseCode_NGPL,
	"NLOD_1_0":                             cpb.SPDXLicenseCode_NLOD_1_0,
	"NLPL":                                 cpb.SPDXLicenseCode_NLPL,
	"NOKIA":                                cpb.SPDXLicenseCode_NOKIA,
	"NOSL":                                 cpb.SPDXLicenseCode_NOSL,
	"NOWEB":                                cpb.SPDXLicenseCode_NOWEB,
	"NPL_1_0":                              cpb.SPDXLicenseCode_NPL_1_0,
	"NPL_1_1":                              cpb.SPDXLicenseCode_NPL_1_1,
	"NPOSL_3_0":                            cpb.SPDXLicenseCode_NPOSL_3_0,
	"NRL":                                  cpb.SPDXLicenseCode_NRL,
	"NTP":                                  cpb.SPDXLicenseCode_NTP,
	"OCCT_PL":                              cpb.SPDXLicenseCode_OCCT_PL,
	"OCLC_2_0":                             cpb.SPDXLicenseCode_OCLC_2_0,
	"O_DB_L_1_0":                           cpb.SPDXLicenseCode_O_DB_L_1_0,
	"OFL_1_0":                              cpb.SPDXLicenseCode_OFL_1_0,
	"OFL_1_1":                              cpb.SPDXLicenseCode_OFL_1_1,
	"OGTSL":                                cpb.SPDXLicenseCode_OGTSL,
	"OLDAP_1_1":                            cpb.SPDXLicenseCode_OLDAP_1_1,
	"OLDAP_1_2":                            cpb.SPDXLicenseCode_OLDAP_1_2,
	"OLDAP_1_3":                            cpb.SPDXLicenseCode_OLDAP_1_3,
	"OLDAP_1_4":                            cpb.SPDXLicenseCode_OLDAP_1_4,
	"OLDAP_2_0_1":                          cpb.SPDXLicenseCode_OLDAP_2_0_1,
	"OLDAP_2_0":                            cpb.SPDXLicenseCode_OLDAP_2_0,
	"OLDAP_2_1":                            cpb.SPDXLicenseCode_OLDAP_2_1,
	"OLDAP_2_2_1":                          cpb.SPDXLicenseCode_OLDAP_2_2_1,
	"OLDAP_2_2_2":                          cpb.SPDXLicenseCode_OLDAP_2_2_2,
	"OLDAP_2_2":                            cpb.SPDXLicenseCode_OLDAP_2_2,
	"OLDAP_2_3":                            cpb.SPDXLicenseCode_OLDAP_2_3,
	"OLDAP_2_4":                            cpb.SPDXLicenseCode_OLDAP_2_4,
	"OLDAP_2_5":                            cpb.SPDXLicenseCode_OLDAP_2_5,
	"OLDAP_2_6":                            cpb.SPDXLicenseCode_OLDAP_2_6,
	"OLDAP_2_7":                            cpb.SPDXLicenseCode_OLDAP_2_7,
	"OLDAP_2_8":                            cpb.SPDXLicenseCode_OLDAP_2_8,
	"OML":                                  cpb.SPDXLicenseCode_OML,
	"OPEN_SSL":                             cpb.SPDXLicenseCode_OPEN_SSL,
	"OPL_1_0":                              cpb.SPDXLicenseCode_OPL_1_0,
	"OSET_PL_2_1":                          cpb.SPDXLicenseCode_OSET_PL_2_1,
	"OSL_1_0":                              cpb.SPDXLicenseCode_OSL_1_0,
	"OSL_1_1":                              cpb.SPDXLicenseCode_OSL_1_1,
	"OSL_2_0":                              cpb.SPDXLicenseCode_OSL_2_0,
	"OSL_2_1":                              cpb.SPDXLicenseCode_OSL_2_1,
	"OSL_3_0":                              cpb.SPDXLicenseCode_OSL_3_0,
	"PDDL_1_0":                             cpb.SPDXLicenseCode_PDDL_1_0,
	"PHP_3_0":                              cpb.SPDXLicenseCode_PHP_3_0,
	"PHP_3_01":                             cpb.SPDXLicenseCode_PHP_3_01,
	"PLEXUS":                               cpb.SPDXLicenseCode_PLEXUS,
	"POSTGRE_SQL":                          cpb.SPDXLicenseCode_POSTGRE_SQL,
	"PSFRAG":                               cpb.SPDXLicenseCode_PSFRAG,
	"PSUTILS":                              cpb.SPDXLicenseCode_PSUTILS,
	"PYTHON_2_0":                           cpb.SPDXLicenseCode_PYTHON_2_0,
	"QHULL":                                cpb.SPDXLicenseCode_QHULL,
	"QPL_1_0":                              cpb.SPDXLicenseCode_QPL_1_0,
	"RDISC":                                cpb.SPDXLicenseCode_RDISC,
	"R_HE_COS_1_1":                         cpb.SPDXLicenseCode_R_HE_COS_1_1,
	"RPL_1_1":                              cpb.SPDXLicenseCode_RPL_1_1,
	"RPL_1_5":                              cpb.SPDXLicenseCode_RPL_1_5,
	"RPSL_1_0":                             cpb.SPDXLicenseCode_RPSL_1_0,
	"RSA_MD":                               cpb.SPDXLicenseCode_RSA_MD,
	"RSCPL":                                cpb.SPDXLicenseCode_RSCPL,
	"RUBY":                                 cpb.SPDXLicenseCode_RUBY,
	"SAX_PD":                               cpb.SPDXLicenseCode_SAX_PD,
	"SAXPATH":                              cpb.SPDXLicenseCode_SAXPATH,
	"SCEA":                                 cpb.SPDXLicenseCode_SCEA,
	"SENDMAIL":                             cpb.SPDXLicenseCode_SENDMAIL,
	"SGI_B_1_0":                            cpb.SPDXLicenseCode_SGI_B_1_0,
	"SGI_B_1_1":                            cpb.SPDXLicenseCode_SGI_B_1_1,
	"SGI_B_2_0":                            cpb.SPDXLicenseCode_SGI_B_2_0,
	"SIM_PL_2_0":                           cpb.SPDXLicenseCode_SIM_PL_2_0,
	"SISSL_1_2":                            cpb.SPDXLicenseCode_SISSL_1_2,
	"SISSL":                                cpb.SPDXLicenseCode_SISSL,
	"SLEEPYCAT":                            cpb.SPDXLicenseCode_SLEEPYCAT,
	"SMLNJ":                                cpb.SPDXLicenseCode_SMLNJ,
	"SMPPL":                                cpb.SPDXLicenseCode_SMPPL,
	"SNIA":                                 cpb.SPDXLicenseCode_SNIA,
	"SPENCER_86":                           cpb.SPDXLicenseCode_SPENCER_86,
	"SPENCER_94":                           cpb.SPDXLicenseCode_SPENCER_94,
	"SPENCER_99":                           cpb.SPDXLicenseCode_SPENCER_99,
	"SPL_1_0":                              cpb.SPDXLicenseCode_SPL_1_0,
	"SUGAR_CRM_1_1_3":                      cpb.SPDXLicenseCode_SUGAR_CRM_1_1_3,
	"SWL":                                  cpb.SPDXLicenseCode_SWL,
	"TCL":                                  cpb.SPDXLicenseCode_TCL,
	"TCP_WRAPPERS":                         cpb.SPDXLicenseCode_TCP_WRAPPERS,
	"T_MATE":                               cpb.SPDXLicenseCode_T_MATE,
	"TORQUE_1_1":                           cpb.SPDXLicenseCode_TORQUE_1_1,
	"TOSL":                                 cpb.SPDXLicenseCode_TOSL,
	"UNICODE_DFS_2015":                     cpb.SPDXLicenseCode_UNICODE_DFS_2015,
	"UNICODE_DFS_2016":                     cpb.SPDXLicenseCode_UNICODE_DFS_2016,
	"UNICODE_TOU":                          cpb.SPDXLicenseCode_UNICODE_TOU,
	"UNLICENSE":                            cpb.SPDXLicenseCode_UNLICENSE,
	"UPL_1_0":                              cpb.SPDXLicenseCode_UPL_1_0,
	"VIM":                                  cpb.SPDXLicenseCode_VIM,
	"VOSTROM":                              cpb.SPDXLicenseCode_VOSTROM,
	"VSL_1_0":                              cpb.SPDXLicenseCode_VSL_1_0,
	"W3C_19980720":                         cpb.SPDXLicenseCode_W3C_19980720,
	"W3C_20150513":                         cpb.SPDXLicenseCode_W3C_20150513,
	"W3C":                                  cpb.SPDXLicenseCode_W3C,
	"WATCOM_1_0":                           cpb.SPDXLicenseCode_WATCOM_1_0,
	"WSUIPA":                               cpb.SPDXLicenseCode_WSUIPA,
	"WTFPL":                                cpb.SPDXLicenseCode_WTFPL,
	"X11":                                  cpb.SPDXLicenseCode_X11,
	"XEROX":                                cpb.SPDXLicenseCode_XEROX,
	"X_FREE86_1_1":                         cpb.SPDXLicenseCode_X_FREE86_1_1,
	"XINETD":                               cpb.SPDXLicenseCode_XINETD,
	"XNET":                                 cpb.SPDXLicenseCode_XNET,
	"XPP":                                  cpb.SPDXLicenseCode_XPP,
	"X_SKAT":                               cpb.SPDXLicenseCode_X_SKAT,
	"YPL_1_0":                              cpb.SPDXLicenseCode_YPL_1_0,
	"YPL_1_1":                              cpb.SPDXLicenseCode_YPL_1_1,
	"ZED":                                  cpb.SPDXLicenseCode_ZED,
	"ZEND_2_0":                             cpb.SPDXLicenseCode_ZEND_2_0,
	"ZIMBRA_1_3":                           cpb.SPDXLicenseCode_ZIMBRA_1_3,
	"ZIMBRA_1_4":                           cpb.SPDXLicenseCode_ZIMBRA_1_4,
	"ZLIB_ACKNOWLEDGEMENT":                 cpb.SPDXLicenseCode_ZLIB_ACKNOWLEDGEMENT,
	"ZLIB":                                 cpb.SPDXLicenseCode_ZLIB,
	"ZPL_1_1":                              cpb.SPDXLicenseCode_ZPL_1_1,
	"ZPL_2_0":                              cpb.SPDXLicenseCode_ZPL_2_0,
	"ZPL_2_1":                              cpb.SPDXLicenseCode_ZPL_2_1,
}

// DefaultSearchComparatorCodeMap maps from string to cpb.SearchComparatorCode_Value.
var DefaultSearchComparatorCodeMap = map[string]cpb.SearchComparatorCode_Value{
	"INVALID_UNINITIALIZED": cpb.SearchComparatorCode_INVALID_UNINITIALIZED,
	"EQ":                    cpb.SearchComparatorCode_EQ,
	"NE":                    cpb.SearchComparatorCode_NE,
	"GT":                    cpb.SearchComparatorCode_GT,
	"LT":                    cpb.SearchComparatorCode_LT,
	"GE":                    cpb.SearchComparatorCode_GE,
	"LE":                    cpb.SearchComparatorCode_LE,
	"SA":                    cpb.SearchComparatorCode_SA,
	"EB":                    cpb.SearchComparatorCode_EB,
	"AP":                    cpb.SearchComparatorCode_AP,
}

// DefaultSearchEntryModeCodeMap maps from string to cpb.SearchEntryModeCode_Value.
var DefaultSearchEntryModeCodeMap = map[string]cpb.SearchEntryModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.SearchEntryModeCode_INVALID_UNINITIALIZED,
	"MATCH":                 cpb.SearchEntryModeCode_MATCH,
	"INCLUDE":               cpb.SearchEntryModeCode_INCLUDE,
	"OUTCOME":               cpb.SearchEntryModeCode_OUTCOME,
}

// DefaultSearchModifierCodeMap maps from string to cpb.SearchModifierCode_Value.
var DefaultSearchModifierCodeMap = map[string]cpb.SearchModifierCode_Value{
	"INVALID_UNINITIALIZED": cpb.SearchModifierCode_INVALID_UNINITIALIZED,
	"MISSING":               cpb.SearchModifierCode_MISSING,
	"EXACT":                 cpb.SearchModifierCode_EXACT,
	"CONTAINS":              cpb.SearchModifierCode_CONTAINS,
	"NOT":                   cpb.SearchModifierCode_NOT,
	"TEXT":                  cpb.SearchModifierCode_TEXT,
	"IN":                    cpb.SearchModifierCode_IN,
	"NOT_IN":                cpb.SearchModifierCode_NOT_IN,
	"BELOW":                 cpb.SearchModifierCode_BELOW,
	"ABOVE":                 cpb.SearchModifierCode_ABOVE,
	"TYPE":                  cpb.SearchModifierCode_TYPE,
	"IDENTIFIER":            cpb.SearchModifierCode_IDENTIFIER,
	"OF_TYPE":               cpb.SearchModifierCode_OF_TYPE,
}

// DefaultSearchParamTypeCodeMap maps from string to cpb.SearchParamTypeCode_Value.
var DefaultSearchParamTypeCodeMap = map[string]cpb.SearchParamTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.SearchParamTypeCode_INVALID_UNINITIALIZED,
	"NUMBER":                cpb.SearchParamTypeCode_NUMBER,
	"DATE":                  cpb.SearchParamTypeCode_DATE,
	"STRING":                cpb.SearchParamTypeCode_STRING,
	"TOKEN":                 cpb.SearchParamTypeCode_TOKEN,
	"REFERENCE":             cpb.SearchParamTypeCode_REFERENCE,
	"COMPOSITE":             cpb.SearchParamTypeCode_COMPOSITE,
	"QUANTITY":              cpb.SearchParamTypeCode_QUANTITY,
	"URI":                   cpb.SearchParamTypeCode_URI,
	"SPECIAL":               cpb.SearchParamTypeCode_SPECIAL,
}

// DefaultSequenceTypeCodeMap maps from string to cpb.SequenceTypeCode_Value.
var DefaultSequenceTypeCodeMap = map[string]cpb.SequenceTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.SequenceTypeCode_INVALID_UNINITIALIZED,
	"AA":                    cpb.SequenceTypeCode_AA,
	"DNA":                   cpb.SequenceTypeCode_DNA,
	"RNA":                   cpb.SequenceTypeCode_RNA,
}

// DefaultSlicingRulesCodeMap maps from string to cpb.SlicingRulesCode_Value.
var DefaultSlicingRulesCodeMap = map[string]cpb.SlicingRulesCode_Value{
	"INVALID_UNINITIALIZED": cpb.SlicingRulesCode_INVALID_UNINITIALIZED,
	"CLOSED":                cpb.SlicingRulesCode_CLOSED,
	"OPEN":                  cpb.SlicingRulesCode_OPEN,
	"OPEN_AT_END":           cpb.SlicingRulesCode_OPEN_AT_END,
}

// DefaultSlotStatusCodeMap maps from string to cpb.SlotStatusCode_Value.
var DefaultSlotStatusCodeMap = map[string]cpb.SlotStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.SlotStatusCode_INVALID_UNINITIALIZED,
	"BUSY":                  cpb.SlotStatusCode_BUSY,
	"FREE":                  cpb.SlotStatusCode_FREE,
	"BUSY_UNAVAILABLE":      cpb.SlotStatusCode_BUSY_UNAVAILABLE,
	"BUSY_TENTATIVE":        cpb.SlotStatusCode_BUSY_TENTATIVE,
	"ENTERED_IN_ERROR":      cpb.SlotStatusCode_ENTERED_IN_ERROR,
}

// DefaultSmartCapabilitiesCodeMap maps from string to cpb.SmartCapabilitiesCode_Value.
var DefaultSmartCapabilitiesCodeMap = map[string]cpb.SmartCapabilitiesCode_Value{
	"INVALID_UNINITIALIZED":         cpb.SmartCapabilitiesCode_INVALID_UNINITIALIZED,
	"LAUNCH_EHR":                    cpb.SmartCapabilitiesCode_LAUNCH_EHR,
	"LAUNCH_STANDALONE":             cpb.SmartCapabilitiesCode_LAUNCH_STANDALONE,
	"CLIENT_PUBLIC":                 cpb.SmartCapabilitiesCode_CLIENT_PUBLIC,
	"CLIENT_CONFIDENTIAL_SYMMETRIC": cpb.SmartCapabilitiesCode_CLIENT_CONFIDENTIAL_SYMMETRIC,
	"SSO_OPENID_CONNECT":            cpb.SmartCapabilitiesCode_SSO_OPENID_CONNECT,
	"CONTEXT_PASSTHROUGH_BANNER":    cpb.SmartCapabilitiesCode_CONTEXT_PASSTHROUGH_BANNER,
	"CONTEXT_PASSTHROUGH_STYLE":     cpb.SmartCapabilitiesCode_CONTEXT_PASSTHROUGH_STYLE,
	"CONTEXT_EHR_PATIENT":           cpb.SmartCapabilitiesCode_CONTEXT_EHR_PATIENT,
	"CONTEXT_EHR_ENCOUNTER":         cpb.SmartCapabilitiesCode_CONTEXT_EHR_ENCOUNTER,
	"CONTEXT_STANDALONE_PATIENT":    cpb.SmartCapabilitiesCode_CONTEXT_STANDALONE_PATIENT,
	"CONTEXT_STANDALONE_ENCOUNTER":  cpb.SmartCapabilitiesCode_CONTEXT_STANDALONE_ENCOUNTER,
	"PERMISSION_OFFLINE":            cpb.SmartCapabilitiesCode_PERMISSION_OFFLINE,
	"PERMISSION_PATIENT":            cpb.SmartCapabilitiesCode_PERMISSION_PATIENT,
	"PERMISSION_USER":               cpb.SmartCapabilitiesCode_PERMISSION_USER,
}

// DefaultSortDirectionCodeMap maps from string to cpb.SortDirectionCode_Value.
var DefaultSortDirectionCodeMap = map[string]cpb.SortDirectionCode_Value{
	"INVALID_UNINITIALIZED": cpb.SortDirectionCode_INVALID_UNINITIALIZED,
	"ASCENDING":             cpb.SortDirectionCode_ASCENDING,
	"DESCENDING":            cpb.SortDirectionCode_DESCENDING,
}

// DefaultSpecimenContainedPreferenceCodeMap maps from string to cpb.SpecimenContainedPreferenceCode_Value.
var DefaultSpecimenContainedPreferenceCodeMap = map[string]cpb.SpecimenContainedPreferenceCode_Value{
	"INVALID_UNINITIALIZED": cpb.SpecimenContainedPreferenceCode_INVALID_UNINITIALIZED,
	"PREFERRED":             cpb.SpecimenContainedPreferenceCode_PREFERRED,
	"ALTERNATE":             cpb.SpecimenContainedPreferenceCode_ALTERNATE,
}

// DefaultSpecimenStatusCodeMap maps from string to cpb.SpecimenStatusCode_Value.
var DefaultSpecimenStatusCodeMap = map[string]cpb.SpecimenStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.SpecimenStatusCode_INVALID_UNINITIALIZED,
	"AVAILABLE":             cpb.SpecimenStatusCode_AVAILABLE,
	"UNAVAILABLE":           cpb.SpecimenStatusCode_UNAVAILABLE,
	"UNSATISFACTORY":        cpb.SpecimenStatusCode_UNSATISFACTORY,
	"ENTERED_IN_ERROR":      cpb.SpecimenStatusCode_ENTERED_IN_ERROR,
}

// DefaultStandardsStatusCodeMap maps from string to cpb.StandardsStatusCode_Value.
var DefaultStandardsStatusCodeMap = map[string]cpb.StandardsStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.StandardsStatusCode_INVALID_UNINITIALIZED,
	"DRAFT":                 cpb.StandardsStatusCode_DRAFT,
	"NORMATIVE":             cpb.StandardsStatusCode_NORMATIVE,
	"TRIAL_USE":             cpb.StandardsStatusCode_TRIAL_USE,
	"INFORMATIVE":           cpb.StandardsStatusCode_INFORMATIVE,
	"DEPRECATED":            cpb.StandardsStatusCode_DEPRECATED,
	"EXTERNAL":              cpb.StandardsStatusCode_EXTERNAL,
}

// DefaultStatusCodeMap maps from string to cpb.StatusCode_Value.
var DefaultStatusCodeMap = map[string]cpb.StatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.StatusCode_INVALID_UNINITIALIZED,
	"ATTESTED":              cpb.StatusCode_ATTESTED,
	"VALIDATED":             cpb.StatusCode_VALIDATED,
	"IN_PROCESS":            cpb.StatusCode_IN_PROCESS,
	"REQ_REVALID":           cpb.StatusCode_REQ_REVALID,
	"VAL_FAIL":              cpb.StatusCode_VAL_FAIL,
	"REVAL_FAIL":            cpb.StatusCode_REVAL_FAIL,
}

// DefaultStrandTypeCodeMap maps from string to cpb.StrandTypeCode_Value.
var DefaultStrandTypeCodeMap = map[string]cpb.StrandTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.StrandTypeCode_INVALID_UNINITIALIZED,
	"WATSON":                cpb.StrandTypeCode_WATSON,
	"CRICK":                 cpb.StrandTypeCode_CRICK,
}

// DefaultStructureDefinitionKindCodeMap maps from string to cpb.StructureDefinitionKindCode_Value.
var DefaultStructureDefinitionKindCodeMap = map[string]cpb.StructureDefinitionKindCode_Value{
	"INVALID_UNINITIALIZED": cpb.StructureDefinitionKindCode_INVALID_UNINITIALIZED,
	"PRIMITIVE_TYPE":        cpb.StructureDefinitionKindCode_PRIMITIVE_TYPE,
	"COMPLEX_TYPE":          cpb.StructureDefinitionKindCode_COMPLEX_TYPE,
	"RESOURCE":              cpb.StructureDefinitionKindCode_RESOURCE,
	"LOGICAL":               cpb.StructureDefinitionKindCode_LOGICAL,
}

// DefaultStructureMapContextTypeCodeMap maps from string to cpb.StructureMapContextTypeCode_Value.
var DefaultStructureMapContextTypeCodeMap = map[string]cpb.StructureMapContextTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.StructureMapContextTypeCode_INVALID_UNINITIALIZED,
	"TYPE":                  cpb.StructureMapContextTypeCode_TYPE,
	"VARIABLE":              cpb.StructureMapContextTypeCode_VARIABLE,
}

// DefaultStructureMapGroupTypeModeCodeMap maps from string to cpb.StructureMapGroupTypeModeCode_Value.
var DefaultStructureMapGroupTypeModeCodeMap = map[string]cpb.StructureMapGroupTypeModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.StructureMapGroupTypeModeCode_INVALID_UNINITIALIZED,
	"NONE":                  cpb.StructureMapGroupTypeModeCode_NONE,
	"TYPES":                 cpb.StructureMapGroupTypeModeCode_TYPES,
	"TYPE_AND_TYPES":        cpb.StructureMapGroupTypeModeCode_TYPE_AND_TYPES,
}

// DefaultStructureMapInputModeCodeMap maps from string to cpb.StructureMapInputModeCode_Value.
var DefaultStructureMapInputModeCodeMap = map[string]cpb.StructureMapInputModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.StructureMapInputModeCode_INVALID_UNINITIALIZED,
	"SOURCE":                cpb.StructureMapInputModeCode_SOURCE,
	"TARGET":                cpb.StructureMapInputModeCode_TARGET,
}

// DefaultStructureMapModelModeCodeMap maps from string to cpb.StructureMapModelModeCode_Value.
var DefaultStructureMapModelModeCodeMap = map[string]cpb.StructureMapModelModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.StructureMapModelModeCode_INVALID_UNINITIALIZED,
	"SOURCE":                cpb.StructureMapModelModeCode_SOURCE,
	"QUERIED":               cpb.StructureMapModelModeCode_QUERIED,
	"TARGET":                cpb.StructureMapModelModeCode_TARGET,
	"PRODUCED":              cpb.StructureMapModelModeCode_PRODUCED,
}

// DefaultStructureMapSourceListModeCodeMap maps from string to cpb.StructureMapSourceListModeCode_Value.
var DefaultStructureMapSourceListModeCodeMap = map[string]cpb.StructureMapSourceListModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.StructureMapSourceListModeCode_INVALID_UNINITIALIZED,
	"FIRST":                 cpb.StructureMapSourceListModeCode_FIRST,
	"NOT_FIRST":             cpb.StructureMapSourceListModeCode_NOT_FIRST,
	"LAST":                  cpb.StructureMapSourceListModeCode_LAST,
	"NOT_LAST":              cpb.StructureMapSourceListModeCode_NOT_LAST,
	"ONLY_ONE":              cpb.StructureMapSourceListModeCode_ONLY_ONE,
}

// DefaultStructureMapTargetListModeCodeMap maps from string to cpb.StructureMapTargetListModeCode_Value.
var DefaultStructureMapTargetListModeCodeMap = map[string]cpb.StructureMapTargetListModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.StructureMapTargetListModeCode_INVALID_UNINITIALIZED,
	"FIRST":                 cpb.StructureMapTargetListModeCode_FIRST,
	"SHARE":                 cpb.StructureMapTargetListModeCode_SHARE,
	"LAST":                  cpb.StructureMapTargetListModeCode_LAST,
	"COLLATE":               cpb.StructureMapTargetListModeCode_COLLATE,
}

// DefaultStructureMapTransformCodeMap maps from string to cpb.StructureMapTransformCode_Value.
var DefaultStructureMapTransformCodeMap = map[string]cpb.StructureMapTransformCode_Value{
	"INVALID_UNINITIALIZED": cpb.StructureMapTransformCode_INVALID_UNINITIALIZED,
	"CREATE":                cpb.StructureMapTransformCode_CREATE,
	"COPY":                  cpb.StructureMapTransformCode_COPY,
	"TRUNCATE":              cpb.StructureMapTransformCode_TRUNCATE,
	"ESCAPE":                cpb.StructureMapTransformCode_ESCAPE,
	"CAST":                  cpb.StructureMapTransformCode_CAST,
	"APPEND":                cpb.StructureMapTransformCode_APPEND,
	"TRANSLATE":             cpb.StructureMapTransformCode_TRANSLATE,
	"REFERENCE":             cpb.StructureMapTransformCode_REFERENCE,
	"DATE_OP":               cpb.StructureMapTransformCode_DATE_OP,
	"UUID":                  cpb.StructureMapTransformCode_UUID,
	"POINTER":               cpb.StructureMapTransformCode_POINTER,
	"EVALUATE":              cpb.StructureMapTransformCode_EVALUATE,
	"CC":                    cpb.StructureMapTransformCode_CC,
	"C":                     cpb.StructureMapTransformCode_C,
	"QTY":                   cpb.StructureMapTransformCode_QTY,
	"ID":                    cpb.StructureMapTransformCode_ID,
	"CP":                    cpb.StructureMapTransformCode_CP,
}

// DefaultSubscriptionChannelTypeCodeMap maps from string to cpb.SubscriptionChannelTypeCode_Value.
var DefaultSubscriptionChannelTypeCodeMap = map[string]cpb.SubscriptionChannelTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.SubscriptionChannelTypeCode_INVALID_UNINITIALIZED,
	"REST_HOOK":             cpb.SubscriptionChannelTypeCode_REST_HOOK,
	"WEBSOCKET":             cpb.SubscriptionChannelTypeCode_WEBSOCKET,
	"EMAIL":                 cpb.SubscriptionChannelTypeCode_EMAIL,
	"SMS":                   cpb.SubscriptionChannelTypeCode_SMS,
	"MESSAGE":               cpb.SubscriptionChannelTypeCode_MESSAGE,
}

// DefaultSubscriptionStatusCodeMap maps from string to cpb.SubscriptionStatusCode_Value.
var DefaultSubscriptionStatusCodeMap = map[string]cpb.SubscriptionStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.SubscriptionStatusCode_INVALID_UNINITIALIZED,
	"REQUESTED":             cpb.SubscriptionStatusCode_REQUESTED,
	"ACTIVE":                cpb.SubscriptionStatusCode_ACTIVE,
	"ERROR":                 cpb.SubscriptionStatusCode_ERROR,
	"OFF":                   cpb.SubscriptionStatusCode_OFF,
}

// DefaultSupplyDeliveryStatusCodeMap maps from string to cpb.SupplyDeliveryStatusCode_Value.
var DefaultSupplyDeliveryStatusCodeMap = map[string]cpb.SupplyDeliveryStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.SupplyDeliveryStatusCode_INVALID_UNINITIALIZED,
	"IN_PROGRESS":           cpb.SupplyDeliveryStatusCode_IN_PROGRESS,
	"COMPLETED":             cpb.SupplyDeliveryStatusCode_COMPLETED,
	"ABANDONED":             cpb.SupplyDeliveryStatusCode_ABANDONED,
	"ENTERED_IN_ERROR":      cpb.SupplyDeliveryStatusCode_ENTERED_IN_ERROR,
}

// DefaultSupplyItemTypeCodeMap maps from string to cpb.SupplyItemTypeCode_Value.
var DefaultSupplyItemTypeCodeMap = map[string]cpb.SupplyItemTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.SupplyItemTypeCode_INVALID_UNINITIALIZED,
	"MEDICATION":            cpb.SupplyItemTypeCode_MEDICATION,
	"DEVICE":                cpb.SupplyItemTypeCode_DEVICE,
}

// DefaultSupplyRequestStatusCodeMap maps from string to cpb.SupplyRequestStatusCode_Value.
var DefaultSupplyRequestStatusCodeMap = map[string]cpb.SupplyRequestStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.SupplyRequestStatusCode_INVALID_UNINITIALIZED,
	"DRAFT":                 cpb.SupplyRequestStatusCode_DRAFT,
	"ACTIVE":                cpb.SupplyRequestStatusCode_ACTIVE,
	"SUSPENDED":             cpb.SupplyRequestStatusCode_SUSPENDED,
	"CANCELLED":             cpb.SupplyRequestStatusCode_CANCELLED,
	"COMPLETED":             cpb.SupplyRequestStatusCode_COMPLETED,
	"ENTERED_IN_ERROR":      cpb.SupplyRequestStatusCode_ENTERED_IN_ERROR,
	"UNKNOWN":               cpb.SupplyRequestStatusCode_UNKNOWN,
}

// DefaultTaskIntentCodeMap maps from string to cpb.TaskIntentCode_Value.
var DefaultTaskIntentCodeMap = map[string]cpb.TaskIntentCode_Value{
	"INVALID_UNINITIALIZED": cpb.TaskIntentCode_INVALID_UNINITIALIZED,
	"UNKNOWN":               cpb.TaskIntentCode_UNKNOWN,
}

// DefaultTaskStatusCodeMap maps from string to cpb.TaskStatusCode_Value.
var DefaultTaskStatusCodeMap = map[string]cpb.TaskStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.TaskStatusCode_INVALID_UNINITIALIZED,
	"DRAFT":                 cpb.TaskStatusCode_DRAFT,
	"REQUESTED":             cpb.TaskStatusCode_REQUESTED,
	"RECEIVED":              cpb.TaskStatusCode_RECEIVED,
	"ACCEPTED":              cpb.TaskStatusCode_ACCEPTED,
	"REJECTED":              cpb.TaskStatusCode_REJECTED,
	"READY":                 cpb.TaskStatusCode_READY,
	"CANCELLED":             cpb.TaskStatusCode_CANCELLED,
	"IN_PROGRESS":           cpb.TaskStatusCode_IN_PROGRESS,
	"ON_HOLD":               cpb.TaskStatusCode_ON_HOLD,
	"FAILED":                cpb.TaskStatusCode_FAILED,
	"COMPLETED":             cpb.TaskStatusCode_COMPLETED,
	"ENTERED_IN_ERROR":      cpb.TaskStatusCode_ENTERED_IN_ERROR,
}

// DefaultTemplateStatusCodeLifeCycleCodeMap maps from string to cpb.TemplateStatusCodeLifeCycleCode_Value.
var DefaultTemplateStatusCodeLifeCycleCodeMap = map[string]cpb.TemplateStatusCodeLifeCycleCode_Value{
	"INVALID_UNINITIALIZED": cpb.TemplateStatusCodeLifeCycleCode_INVALID_UNINITIALIZED,
	"DRAFT":                 cpb.TemplateStatusCodeLifeCycleCode_DRAFT,
	"PENDING":               cpb.TemplateStatusCodeLifeCycleCode_PENDING,
	"ACTIVE":                cpb.TemplateStatusCodeLifeCycleCode_ACTIVE,
	"REVIEW":                cpb.TemplateStatusCodeLifeCycleCode_REVIEW,
	"CANCELLED":             cpb.TemplateStatusCodeLifeCycleCode_CANCELLED,
	"REJECTED":              cpb.TemplateStatusCodeLifeCycleCode_REJECTED,
	"RETIRED":               cpb.TemplateStatusCodeLifeCycleCode_RETIRED,
	"TERMINATED":            cpb.TemplateStatusCodeLifeCycleCode_TERMINATED,
}

// DefaultTestReportActionResultCodeMap maps from string to cpb.TestReportActionResultCode_Value.
var DefaultTestReportActionResultCodeMap = map[string]cpb.TestReportActionResultCode_Value{
	"INVALID_UNINITIALIZED": cpb.TestReportActionResultCode_INVALID_UNINITIALIZED,
	"PASS":                  cpb.TestReportActionResultCode_PASS,
	"SKIP":                  cpb.TestReportActionResultCode_SKIP,
	"FAIL":                  cpb.TestReportActionResultCode_FAIL,
	"WARNING":               cpb.TestReportActionResultCode_WARNING,
	"ERROR":                 cpb.TestReportActionResultCode_ERROR,
}

// DefaultTestReportParticipantTypeCodeMap maps from string to cpb.TestReportParticipantTypeCode_Value.
var DefaultTestReportParticipantTypeCodeMap = map[string]cpb.TestReportParticipantTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.TestReportParticipantTypeCode_INVALID_UNINITIALIZED,
	"TEST_ENGINE":           cpb.TestReportParticipantTypeCode_TEST_ENGINE,
	"CLIENT":                cpb.TestReportParticipantTypeCode_CLIENT,
	"SERVER":                cpb.TestReportParticipantTypeCode_SERVER,
}

// DefaultTestReportResultCodeMap maps from string to cpb.TestReportResultCode_Value.
var DefaultTestReportResultCodeMap = map[string]cpb.TestReportResultCode_Value{
	"INVALID_UNINITIALIZED": cpb.TestReportResultCode_INVALID_UNINITIALIZED,
	"PASS":                  cpb.TestReportResultCode_PASS,
	"FAIL":                  cpb.TestReportResultCode_FAIL,
	"PENDING":               cpb.TestReportResultCode_PENDING,
}

// DefaultTestReportStatusCodeMap maps from string to cpb.TestReportStatusCode_Value.
var DefaultTestReportStatusCodeMap = map[string]cpb.TestReportStatusCode_Value{
	"INVALID_UNINITIALIZED": cpb.TestReportStatusCode_INVALID_UNINITIALIZED,
	"COMPLETED":             cpb.TestReportStatusCode_COMPLETED,
	"IN_PROGRESS":           cpb.TestReportStatusCode_IN_PROGRESS,
	"WAITING":               cpb.TestReportStatusCode_WAITING,
	"STOPPED":               cpb.TestReportStatusCode_STOPPED,
	"ENTERED_IN_ERROR":      cpb.TestReportStatusCode_ENTERED_IN_ERROR,
}

// DefaultTestScriptRequestMethodCodeMap maps from string to cpb.TestScriptRequestMethodCode_Value.
var DefaultTestScriptRequestMethodCodeMap = map[string]cpb.TestScriptRequestMethodCode_Value{
	"INVALID_UNINITIALIZED": cpb.TestScriptRequestMethodCode_INVALID_UNINITIALIZED,
	"DELETE":                cpb.TestScriptRequestMethodCode_DELETE,
	"GET":                   cpb.TestScriptRequestMethodCode_GET,
	"OPTIONS":               cpb.TestScriptRequestMethodCode_OPTIONS,
	"PATCH":                 cpb.TestScriptRequestMethodCode_PATCH,
	"POST":                  cpb.TestScriptRequestMethodCode_POST,
	"PUT":                   cpb.TestScriptRequestMethodCode_PUT,
	"HEAD":                  cpb.TestScriptRequestMethodCode_HEAD,
}

// DefaultTriggerTypeCodeMap maps from string to cpb.TriggerTypeCode_Value.
var DefaultTriggerTypeCodeMap = map[string]cpb.TriggerTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.TriggerTypeCode_INVALID_UNINITIALIZED,
	"NAMED_EVENT":           cpb.TriggerTypeCode_NAMED_EVENT,
	"PERIODIC":              cpb.TriggerTypeCode_PERIODIC,
	"DATA_CHANGED":          cpb.TriggerTypeCode_DATA_CHANGED,
	"DATA_ADDED":            cpb.TriggerTypeCode_DATA_ADDED,
	"DATA_MODIFIED":         cpb.TriggerTypeCode_DATA_MODIFIED,
	"DATA_REMOVED":          cpb.TriggerTypeCode_DATA_REMOVED,
	"DATA_ACCESSED":         cpb.TriggerTypeCode_DATA_ACCESSED,
	"DATA_ACCESS_ENDED":     cpb.TriggerTypeCode_DATA_ACCESS_ENDED,
}

// DefaultTypeDerivationRuleCodeMap maps from string to cpb.TypeDerivationRuleCode_Value.
var DefaultTypeDerivationRuleCodeMap = map[string]cpb.TypeDerivationRuleCode_Value{
	"INVALID_UNINITIALIZED": cpb.TypeDerivationRuleCode_INVALID_UNINITIALIZED,
	"SPECIALIZATION":        cpb.TypeDerivationRuleCode_SPECIALIZATION,
	"CONSTRAINT":            cpb.TypeDerivationRuleCode_CONSTRAINT,
}

// DefaultUDIEntryTypeCodeMap maps from string to cpb.UDIEntryTypeCode_Value.
var DefaultUDIEntryTypeCodeMap = map[string]cpb.UDIEntryTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.UDIEntryTypeCode_INVALID_UNINITIALIZED,
	"BARCODE":               cpb.UDIEntryTypeCode_BARCODE,
	"RFID":                  cpb.UDIEntryTypeCode_RFID,
	"MANUAL":                cpb.UDIEntryTypeCode_MANUAL,
	"CARD":                  cpb.UDIEntryTypeCode_CARD,
	"SELF_REPORTED":         cpb.UDIEntryTypeCode_SELF_REPORTED,
	"UNKNOWN":               cpb.UDIEntryTypeCode_UNKNOWN,
}

// DefaultUseCodeMap maps from string to cpb.UseCode_Value.
var DefaultUseCodeMap = map[string]cpb.UseCode_Value{
	"INVALID_UNINITIALIZED": cpb.UseCode_INVALID_UNINITIALIZED,
	"CLAIM":                 cpb.UseCode_CLAIM,
	"PREAUTHORIZATION":      cpb.UseCode_PREAUTHORIZATION,
	"PREDETERMINATION":      cpb.UseCode_PREDETERMINATION,
}

// DefaultV20444CodeMap maps from string to cpb.V20444Code_Value.
var DefaultV20444CodeMap = map[string]cpb.V20444Code_Value{
	"INVALID_UNINITIALIZED": cpb.V20444Code_INVALID_UNINITIALIZED,
	"F":                     cpb.V20444Code_F,
	"G":                     cpb.V20444Code_G,
}

// DefaultV3AddressUseCodeMap maps from string to cpb.V3AddressUseCode_Value.
var DefaultV3AddressUseCodeMap = map[string]cpb.V3AddressUseCode_Value{
	"INVALID_UNINITIALIZED":         cpb.V3AddressUseCode_INVALID_UNINITIALIZED,
	"GENERAL_ADDRESS_USE":           cpb.V3AddressUseCode_GENERAL_ADDRESS_USE,
	"BAD":                           cpb.V3AddressUseCode_BAD,
	"CONF":                          cpb.V3AddressUseCode_CONF,
	"H":                             cpb.V3AddressUseCode_H,
	"HP":                            cpb.V3AddressUseCode_HP,
	"HV":                            cpb.V3AddressUseCode_HV,
	"OLD":                           cpb.V3AddressUseCode_OLD,
	"TMP":                           cpb.V3AddressUseCode_TMP,
	"WP":                            cpb.V3AddressUseCode_WP,
	"DIR":                           cpb.V3AddressUseCode_DIR,
	"PUB":                           cpb.V3AddressUseCode_PUB,
	"POSTAL_ADDRESS_USE":            cpb.V3AddressUseCode_POSTAL_ADDRESS_USE,
	"PHYS":                          cpb.V3AddressUseCode_PHYS,
	"PST":                           cpb.V3AddressUseCode_PST,
	"TELECOMMUNICATION_ADDRESS_USE": cpb.V3AddressUseCode_TELECOMMUNICATION_ADDRESS_USE,
	"AS":                            cpb.V3AddressUseCode_AS,
	"EC":                            cpb.V3AddressUseCode_EC,
	"MC":                            cpb.V3AddressUseCode_MC,
	"PG":                            cpb.V3AddressUseCode_PG,
}

// DefaultV3ConfidentialityCodeMap maps from string to cpb.V3ConfidentialityCode_Value.
var DefaultV3ConfidentialityCodeMap = map[string]cpb.V3ConfidentialityCode_Value{
	"INVALID_UNINITIALIZED":          cpb.V3ConfidentialityCode_INVALID_UNINITIALIZED,
	"CONFIDENTIALITY":                cpb.V3ConfidentialityCode_CONFIDENTIALITY,
	"L":                              cpb.V3ConfidentialityCode_L,
	"M":                              cpb.V3ConfidentialityCode_M,
	"N":                              cpb.V3ConfidentialityCode_N,
	"R":                              cpb.V3ConfidentialityCode_R,
	"U":                              cpb.V3ConfidentialityCode_U,
	"V":                              cpb.V3ConfidentialityCode_V,
	"CONFIDENTIALITY_BY_ACCESS_KIND": cpb.V3ConfidentialityCode_CONFIDENTIALITY_BY_ACCESS_KIND,
	"B":                              cpb.V3ConfidentialityCode_B,
	"D":                              cpb.V3ConfidentialityCode_D,
	"I":                              cpb.V3ConfidentialityCode_I,
	"CONFIDENTIALITY_BY_INFO_TYPE":   cpb.V3ConfidentialityCode_CONFIDENTIALITY_BY_INFO_TYPE,
	"ETH":                            cpb.V3ConfidentialityCode_ETH,
	"HIV":                            cpb.V3ConfidentialityCode_HIV,
	"PSY":                            cpb.V3ConfidentialityCode_PSY,
	"SDV":                            cpb.V3ConfidentialityCode_SDV,
	"CONFIDENTIALITY_MODIFIERS":      cpb.V3ConfidentialityCode_CONFIDENTIALITY_MODIFIERS,
	"C":                              cpb.V3ConfidentialityCode_C,
	"S":                              cpb.V3ConfidentialityCode_S,
	"T":                              cpb.V3ConfidentialityCode_T,
}

// DefaultV3EntityNamePartQualifierCodeMap maps from string to cpb.V3EntityNamePartQualifierCode_Value.
var DefaultV3EntityNamePartQualifierCodeMap = map[string]cpb.V3EntityNamePartQualifierCode_Value{
	"INVALID_UNINITIALIZED":            cpb.V3EntityNamePartQualifierCode_INVALID_UNINITIALIZED,
	"ORGANIZATION_NAME_PART_QUALIFIER": cpb.V3EntityNamePartQualifierCode_ORGANIZATION_NAME_PART_QUALIFIER,
	"AC":                               cpb.V3EntityNamePartQualifierCode_AC,
	"AD":                               cpb.V3EntityNamePartQualifierCode_AD,
	"BR":                               cpb.V3EntityNamePartQualifierCode_BR,
	"CL":                               cpb.V3EntityNamePartQualifierCode_CL,
	"IN":                               cpb.V3EntityNamePartQualifierCode_IN,
	"LS":                               cpb.V3EntityNamePartQualifierCode_LS,
	"NB":                               cpb.V3EntityNamePartQualifierCode_NB,
	"PR":                               cpb.V3EntityNamePartQualifierCode_PR,
	"SP":                               cpb.V3EntityNamePartQualifierCode_SP,
	"TITLE":                            cpb.V3EntityNamePartQualifierCode_TITLE,
	"VV":                               cpb.V3EntityNamePartQualifierCode_VV,
	"PHARMACEUTICAL_ENTITY_NAME_PART_QUALIFIERS": cpb.V3EntityNamePartQualifierCode_PHARMACEUTICAL_ENTITY_NAME_PART_QUALIFIERS,
	"CON":                               cpb.V3EntityNamePartQualifierCode_CON,
	"DEV":                               cpb.V3EntityNamePartQualifierCode_DEV,
	"FLAV":                              cpb.V3EntityNamePartQualifierCode_FLAV,
	"FORMUL":                            cpb.V3EntityNamePartQualifierCode_FORMUL,
	"FRM":                               cpb.V3EntityNamePartQualifierCode_FRM,
	"INV":                               cpb.V3EntityNamePartQualifierCode_INV,
	"POPUL":                             cpb.V3EntityNamePartQualifierCode_POPUL,
	"SCI":                               cpb.V3EntityNamePartQualifierCode_SCI,
	"STR":                               cpb.V3EntityNamePartQualifierCode_STR,
	"TIME":                              cpb.V3EntityNamePartQualifierCode_TIME,
	"TMK":                               cpb.V3EntityNamePartQualifierCode_TMK,
	"USE":                               cpb.V3EntityNamePartQualifierCode_USE,
	"PERSON_NAME_PART_QUALIFIER":        cpb.V3EntityNamePartQualifierCode_PERSON_NAME_PART_QUALIFIER,
	"PERSON_NAME_PART_AFFIX_TYPES":      cpb.V3EntityNamePartQualifierCode_PERSON_NAME_PART_AFFIX_TYPES,
	"PERSON_NAME_PART_CHANGE_QUALIFIER": cpb.V3EntityNamePartQualifierCode_PERSON_NAME_PART_CHANGE_QUALIFIER,
	"PERSON_NAME_PART_MISC_QUALIFIER":   cpb.V3EntityNamePartQualifierCode_PERSON_NAME_PART_MISC_QUALIFIER,
}

// DefaultV3EntityNamePartQualifierR2CodeMap maps from string to cpb.V3EntityNamePartQualifierR2Code_Value.
var DefaultV3EntityNamePartQualifierR2CodeMap = map[string]cpb.V3EntityNamePartQualifierR2Code_Value{
	"INVALID_UNINITIALIZED": cpb.V3EntityNamePartQualifierR2Code_INVALID_UNINITIALIZED,
	"AD":                    cpb.V3EntityNamePartQualifierR2Code_AD,
	"SP":                    cpb.V3EntityNamePartQualifierR2Code_SP,
	"BR":                    cpb.V3EntityNamePartQualifierR2Code_BR,
	"CL":                    cpb.V3EntityNamePartQualifierR2Code_CL,
	"IN":                    cpb.V3EntityNamePartQualifierR2Code_IN,
	"LS":                    cpb.V3EntityNamePartQualifierR2Code_LS,
	"MID":                   cpb.V3EntityNamePartQualifierR2Code_MID,
	"PFX":                   cpb.V3EntityNamePartQualifierR2Code_PFX,
	"PHARMACEUTICAL_ENTITY_NAME_PART_QUALIFIERS": cpb.V3EntityNamePartQualifierR2Code_PHARMACEUTICAL_ENTITY_NAME_PART_QUALIFIERS,
	"CON":          cpb.V3EntityNamePartQualifierR2Code_CON,
	"DEV":          cpb.V3EntityNamePartQualifierR2Code_DEV,
	"FLAV":         cpb.V3EntityNamePartQualifierR2Code_FLAV,
	"FORMUL":       cpb.V3EntityNamePartQualifierR2Code_FORMUL,
	"FRM":          cpb.V3EntityNamePartQualifierR2Code_FRM,
	"INV":          cpb.V3EntityNamePartQualifierR2Code_INV,
	"POPUL":        cpb.V3EntityNamePartQualifierR2Code_POPUL,
	"SCI":          cpb.V3EntityNamePartQualifierR2Code_SCI,
	"STR":          cpb.V3EntityNamePartQualifierR2Code_STR,
	"TIME":         cpb.V3EntityNamePartQualifierR2Code_TIME,
	"TMK":          cpb.V3EntityNamePartQualifierR2Code_TMK,
	"USE":          cpb.V3EntityNamePartQualifierR2Code_USE,
	"SFX":          cpb.V3EntityNamePartQualifierR2Code_SFX,
	"TITLE_STYLES": cpb.V3EntityNamePartQualifierR2Code_TITLE_STYLES,
	"AC":           cpb.V3EntityNamePartQualifierR2Code_AC,
	"HON":          cpb.V3EntityNamePartQualifierR2Code_HON,
	"NB":           cpb.V3EntityNamePartQualifierR2Code_NB,
	"PR":           cpb.V3EntityNamePartQualifierR2Code_PR,
}

// DefaultV3EntityNameUseCodeMap maps from string to cpb.V3EntityNameUseCode_Value.
var DefaultV3EntityNameUseCodeMap = map[string]cpb.V3EntityNameUseCode_Value{
	"INVALID_UNINITIALIZED":   cpb.V3EntityNameUseCode_INVALID_UNINITIALIZED,
	"NAME_REPRESENTATION_USE": cpb.V3EntityNameUseCode_NAME_REPRESENTATION_USE,
	"ABC":                     cpb.V3EntityNameUseCode_ABC,
	"IDE":                     cpb.V3EntityNameUseCode_IDE,
	"SYL":                     cpb.V3EntityNameUseCode_SYL,
	"ASGN":                    cpb.V3EntityNameUseCode_ASGN,
	"C":                       cpb.V3EntityNameUseCode_C,
	"I":                       cpb.V3EntityNameUseCode_I,
	"L":                       cpb.V3EntityNameUseCode_L,
	"OR":                      cpb.V3EntityNameUseCode_OR,
	"P":                       cpb.V3EntityNameUseCode_P,
	"A":                       cpb.V3EntityNameUseCode_A,
	"R":                       cpb.V3EntityNameUseCode_R,
	"SRCH":                    cpb.V3EntityNameUseCode_SRCH,
	"PHON":                    cpb.V3EntityNameUseCode_PHON,
	"SNDX":                    cpb.V3EntityNameUseCode_SNDX,
}

// DefaultV3EntityNameUseR2CodeMap maps from string to cpb.V3EntityNameUseR2Code_Value.
var DefaultV3EntityNameUseR2CodeMap = map[string]cpb.V3EntityNameUseR2Code_Value{
	"INVALID_UNINITIALIZED":   cpb.V3EntityNameUseR2Code_INVALID_UNINITIALIZED,
	"ASSUMED":                 cpb.V3EntityNameUseR2Code_ASSUMED,
	"A":                       cpb.V3EntityNameUseR2Code_A,
	"ANON":                    cpb.V3EntityNameUseR2Code_ANON,
	"I":                       cpb.V3EntityNameUseR2Code_I,
	"P":                       cpb.V3EntityNameUseR2Code_P,
	"R":                       cpb.V3EntityNameUseR2Code_R,
	"C":                       cpb.V3EntityNameUseR2Code_C,
	"M":                       cpb.V3EntityNameUseR2Code_M,
	"NAME_REPRESENTATION_USE": cpb.V3EntityNameUseR2Code_NAME_REPRESENTATION_USE,
	"ABC":                     cpb.V3EntityNameUseR2Code_ABC,
	"IDE":                     cpb.V3EntityNameUseR2Code_IDE,
	"SYL":                     cpb.V3EntityNameUseR2Code_SYL,
	"OLD":                     cpb.V3EntityNameUseR2Code_OLD,
	"DN":                      cpb.V3EntityNameUseR2Code_DN,
	"OR":                      cpb.V3EntityNameUseR2Code_OR,
	"PHON":                    cpb.V3EntityNameUseR2Code_PHON,
	"SRCH":                    cpb.V3EntityNameUseR2Code_SRCH,
	"T":                       cpb.V3EntityNameUseR2Code_T,
}

// DefaultV3NullFlavorCodeMap maps from string to cpb.V3NullFlavorCode_Value.
var DefaultV3NullFlavorCodeMap = map[string]cpb.V3NullFlavorCode_Value{
	"INVALID_UNINITIALIZED": cpb.V3NullFlavorCode_INVALID_UNINITIALIZED,
	"NI":                    cpb.V3NullFlavorCode_NI,
	"INV":                   cpb.V3NullFlavorCode_INV,
	"DER":                   cpb.V3NullFlavorCode_DER,
	"OTH":                   cpb.V3NullFlavorCode_OTH,
	"NINF":                  cpb.V3NullFlavorCode_NINF,
	"PINF":                  cpb.V3NullFlavorCode_PINF,
	"UNC":                   cpb.V3NullFlavorCode_UNC,
	"MSK":                   cpb.V3NullFlavorCode_MSK,
	"NA":                    cpb.V3NullFlavorCode_NA,
	"UNK":                   cpb.V3NullFlavorCode_UNK,
	"ASKU":                  cpb.V3NullFlavorCode_ASKU,
	"NAV":                   cpb.V3NullFlavorCode_NAV,
	"NASK":                  cpb.V3NullFlavorCode_NASK,
	"NAVU":                  cpb.V3NullFlavorCode_NAVU,
	"QS":                    cpb.V3NullFlavorCode_QS,
	"TRC":                   cpb.V3NullFlavorCode_TRC,
	"NP":                    cpb.V3NullFlavorCode_NP,
}

// DefaultV3ParticipationModeCodeMap maps from string to cpb.V3ParticipationModeCode_Value.
var DefaultV3ParticipationModeCodeMap = map[string]cpb.V3ParticipationModeCode_Value{
	"INVALID_UNINITIALIZED": cpb.V3ParticipationModeCode_INVALID_UNINITIALIZED,
	"ELECTRONIC":            cpb.V3ParticipationModeCode_ELECTRONIC,
	"PHYSICAL":              cpb.V3ParticipationModeCode_PHYSICAL,
	"REMOTE":                cpb.V3ParticipationModeCode_REMOTE,
	"VERBAL":                cpb.V3ParticipationModeCode_VERBAL,
	"DICTATE":               cpb.V3ParticipationModeCode_DICTATE,
	"FACE":                  cpb.V3ParticipationModeCode_FACE,
	"PHONE":                 cpb.V3ParticipationModeCode_PHONE,
	"VIDEOCONF":             cpb.V3ParticipationModeCode_VIDEOCONF,
	"WRITTEN":               cpb.V3ParticipationModeCode_WRITTEN,
	"FAXWRIT":               cpb.V3ParticipationModeCode_FAXWRIT,
	"HANDWRIT":              cpb.V3ParticipationModeCode_HANDWRIT,
	"MAILWRIT":              cpb.V3ParticipationModeCode_MAILWRIT,
	"ONLINEWRIT":            cpb.V3ParticipationModeCode_ONLINEWRIT,
	"EMAILWRIT":             cpb.V3ParticipationModeCode_EMAILWRIT,
	"TYPEWRIT":              cpb.V3ParticipationModeCode_TYPEWRIT,
}

// DefaultV3ProbabilityDistributionTypeCodeMap maps from string to cpb.V3ProbabilityDistributionTypeCode_Value.
var DefaultV3ProbabilityDistributionTypeCodeMap = map[string]cpb.V3ProbabilityDistributionTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.V3ProbabilityDistributionTypeCode_INVALID_UNINITIALIZED,
	"B":                     cpb.V3ProbabilityDistributionTypeCode_B,
	"E":                     cpb.V3ProbabilityDistributionTypeCode_E,
	"F":                     cpb.V3ProbabilityDistributionTypeCode_F,
	"G":                     cpb.V3ProbabilityDistributionTypeCode_G,
	"LN":                    cpb.V3ProbabilityDistributionTypeCode_LN,
	"N":                     cpb.V3ProbabilityDistributionTypeCode_N,
	"T":                     cpb.V3ProbabilityDistributionTypeCode_T,
	"U":                     cpb.V3ProbabilityDistributionTypeCode_U,
	"X2":                    cpb.V3ProbabilityDistributionTypeCode_X2,
}

// DefaultV3RoleCodeMap maps from string to cpb.V3RoleCode_Value.
var DefaultV3RoleCodeMap = map[string]cpb.V3RoleCode_Value{
	"INVALID_UNINITIALIZED":      cpb.V3RoleCode_INVALID_UNINITIALIZED,
	"AFFILIATION_ROLE_TYPE":      cpb.V3RoleCode_AFFILIATION_ROLE_TYPE,
	"AGENT_ROLE_TYPE":            cpb.V3RoleCode_AGENT_ROLE_TYPE,
	"AMENDER":                    cpb.V3RoleCode_AMENDER,
	"CLASSIFIER":                 cpb.V3RoleCode_CLASSIFIER,
	"CONSENTER":                  cpb.V3RoleCode_CONSENTER,
	"CONSWIT":                    cpb.V3RoleCode_CONSWIT,
	"COPART":                     cpb.V3RoleCode_COPART,
	"DECLASSIFIER":               cpb.V3RoleCode_DECLASSIFIER,
	"DELEGATEE":                  cpb.V3RoleCode_DELEGATEE,
	"DELEGATOR":                  cpb.V3RoleCode_DELEGATOR,
	"DOWNGRDER":                  cpb.V3RoleCode_DOWNGRDER,
	"DRIVCLASSIFIER":             cpb.V3RoleCode_DRIVCLASSIFIER,
	"GRANTEE":                    cpb.V3RoleCode_GRANTEE,
	"GRANTOR":                    cpb.V3RoleCode_GRANTOR,
	"INTPRTER":                   cpb.V3RoleCode_INTPRTER,
	"REVIEWER":                   cpb.V3RoleCode_REVIEWER,
	"VALIDATOR":                  cpb.V3RoleCode_VALIDATOR,
	"COVERAGE_SPONSOR_ROLE_TYPE": cpb.V3RoleCode_COVERAGE_SPONSOR_ROLE_TYPE,
	"FULLINS":                    cpb.V3RoleCode_FULLINS,
	"SELFINS":                    cpb.V3RoleCode_SELFINS,
	"PAYOR_ROLE_TYPE":            cpb.V3RoleCode_PAYOR_ROLE_TYPE,
	"ENROLBKR":                   cpb.V3RoleCode_ENROLBKR,
	"TPA":                        cpb.V3RoleCode_TPA,
	"UMO":                        cpb.V3RoleCode_UMO,
	"RESPRSN":                    cpb.V3RoleCode_RESPRSN,
	"EXCEST":                     cpb.V3RoleCode_EXCEST,
	"GUADLTM":                    cpb.V3RoleCode_GUADLTM,
	"GUARD":                      cpb.V3RoleCode_GUARD,
	"POWATT":                     cpb.V3RoleCode_POWATT,
	"DPOWATT":                    cpb.V3RoleCode_DPOWATT,
	"HPOWATT":                    cpb.V3RoleCode_HPOWATT,
	"SPOWATT":                    cpb.V3RoleCode_SPOWATT,
	"ASSIGNED_ROLE_TYPE":         cpb.V3RoleCode_ASSIGNED_ROLE_TYPE,
	"ASSIGNED_NON_PERSON_LIVING_SUBJECT_ROLE_TYPE": cpb.V3RoleCode_ASSIGNED_NON_PERSON_LIVING_SUBJECT_ROLE_TYPE,
	"ASSIST":                               cpb.V3RoleCode_ASSIST,
	"BIOTH":                                cpb.V3RoleCode_BIOTH,
	"ANTIBIOT":                             cpb.V3RoleCode_ANTIBIOT,
	"DEBR":                                 cpb.V3RoleCode_DEBR,
	"CCO":                                  cpb.V3RoleCode_CCO,
	"SEE":                                  cpb.V3RoleCode_SEE,
	"SNIFF":                                cpb.V3RoleCode_SNIFF,
	"CERTIFIED_ENTITY_TYPE":                cpb.V3RoleCode_CERTIFIED_ENTITY_TYPE,
	"CITIZEN_ROLE_TYPE":                    cpb.V3RoleCode_CITIZEN_ROLE_TYPE,
	"CAS":                                  cpb.V3RoleCode_CAS,
	"CASM":                                 cpb.V3RoleCode_CASM,
	"CN":                                   cpb.V3RoleCode_CN,
	"CNRP":                                 cpb.V3RoleCode_CNRP,
	"CNRPM":                                cpb.V3RoleCode_CNRPM,
	"CPCA":                                 cpb.V3RoleCode_CPCA,
	"CRP":                                  cpb.V3RoleCode_CRP,
	"CRPM":                                 cpb.V3RoleCode_CRPM,
	"CONTACT_ROLE_TYPE":                    cpb.V3RoleCode_CONTACT_ROLE_TYPE,
	"ADMINISTRATIVE_CONTACT_ROLE_TYPE":     cpb.V3RoleCode_ADMINISTRATIVE_CONTACT_ROLE_TYPE,
	"BILL":                                 cpb.V3RoleCode_BILL,
	"ORG":                                  cpb.V3RoleCode_ORG,
	"PAYOR":                                cpb.V3RoleCode_PAYOR,
	"ECON":                                 cpb.V3RoleCode_ECON,
	"NOK":                                  cpb.V3RoleCode_NOK,
	"IDENTIFIED_ENTITY_TYPE":               cpb.V3RoleCode_IDENTIFIED_ENTITY_TYPE,
	"LOCATION_IDENTIFIED_ENTITY_ROLE_CODE": cpb.V3RoleCode_LOCATION_IDENTIFIED_ENTITY_ROLE_CODE,
	"ACHFID":                               cpb.V3RoleCode_ACHFID,
	"JURID":                                cpb.V3RoleCode_JURID,
	"LOCHFID":                              cpb.V3RoleCode_LOCHFID,
	"LIVING_SUBJECT_PRODUCTION_CLASS":      cpb.V3RoleCode_LIVING_SUBJECT_PRODUCTION_CLASS,
	"BF":                                   cpb.V3RoleCode_BF,
	"BL":                                   cpb.V3RoleCode_BL,
	"BR":                                   cpb.V3RoleCode_BR,
	"CO":                                   cpb.V3RoleCode_CO,
	"DA":                                   cpb.V3RoleCode_DA,
	"DR":                                   cpb.V3RoleCode_DR,
	"DU":                                   cpb.V3RoleCode_DU,
	"FI":                                   cpb.V3RoleCode_FI,
	"LY":                                   cpb.V3RoleCode_LY,
	"MT":                                   cpb.V3RoleCode_MT,
	"MU":                                   cpb.V3RoleCode_MU,
	"PL":                                   cpb.V3RoleCode_PL,
	"RC":                                   cpb.V3RoleCode_RC,
	"SH":                                   cpb.V3RoleCode_SH,
	"VL":                                   cpb.V3RoleCode_VL,
	"WL":                                   cpb.V3RoleCode_WL,
	"WO":                                   cpb.V3RoleCode_WO,
	"MEDICATION_GENERALIZATION_ROLE_TYPE":  cpb.V3RoleCode_MEDICATION_GENERALIZATION_ROLE_TYPE,
	"DC":                                   cpb.V3RoleCode_DC,
	"GD":                                   cpb.V3RoleCode_GD,
	"GDF":                                  cpb.V3RoleCode_GDF,
	"GDS":                                  cpb.V3RoleCode_GDS,
	"GDSF":                                 cpb.V3RoleCode_GDSF,
	"MGDSF":                                cpb.V3RoleCode_MGDSF,
	"MEMBER_ROLE_TYPE":                     cpb.V3RoleCode_MEMBER_ROLE_TYPE,
	"TRB":                                  cpb.V3RoleCode_TRB,
	"PERSONAL_RELATIONSHIP_ROLE_TYPE":      cpb.V3RoleCode_PERSONAL_RELATIONSHIP_ROLE_TYPE,
	"FAMMEMB":                              cpb.V3RoleCode_FAMMEMB,
	"CHILD":                                cpb.V3RoleCode_CHILD,
	"CHLDADOPT":                            cpb.V3RoleCode_CHLDADOPT,
	"DAUADOPT":                             cpb.V3RoleCode_DAUADOPT,
	"SONADOPT":                             cpb.V3RoleCode_SONADOPT,
	"CHLDFOST":                             cpb.V3RoleCode_CHLDFOST,
	"DAUFOST":                              cpb.V3RoleCode_DAUFOST,
	"SONFOST":                              cpb.V3RoleCode_SONFOST,
	"DAUC":                                 cpb.V3RoleCode_DAUC,
	"DAU":                                  cpb.V3RoleCode_DAU,
	"STPDAU":                               cpb.V3RoleCode_STPDAU,
	"NCHILD":                               cpb.V3RoleCode_NCHILD,
	"SON":                                  cpb.V3RoleCode_SON,
	"SONC":                                 cpb.V3RoleCode_SONC,
	"STPSON":                               cpb.V3RoleCode_STPSON,
	"STPCHLD":                              cpb.V3RoleCode_STPCHLD,
	"EXT":                                  cpb.V3RoleCode_EXT,
	"AUNT":                                 cpb.V3RoleCode_AUNT,
	"MAUNT":                                cpb.V3RoleCode_MAUNT,
	"PAUNT":                                cpb.V3RoleCode_PAUNT,
	"COUSN":                                cpb.V3RoleCode_COUSN,
	"MCOUSN":                               cpb.V3RoleCode_MCOUSN,
	"PCOUSN":                               cpb.V3RoleCode_PCOUSN,
	"GGRPRN":                               cpb.V3RoleCode_GGRPRN,
	"GGRFTH":                               cpb.V3RoleCode_GGRFTH,
	"MGGRFTH":                              cpb.V3RoleCode_MGGRFTH,
	"PGGRFTH":                              cpb.V3RoleCode_PGGRFTH,
	"GGRMTH":                               cpb.V3RoleCode_GGRMTH,
	"MGGRMTH":                              cpb.V3RoleCode_MGGRMTH,
	"PGGRMTH":                              cpb.V3RoleCode_PGGRMTH,
	"MGGRPRN":                              cpb.V3RoleCode_MGGRPRN,
	"PGGRPRN":                              cpb.V3RoleCode_PGGRPRN,
	"GRNDCHILD":                            cpb.V3RoleCode_GRNDCHILD,
	"GRNDDAU":                              cpb.V3RoleCode_GRNDDAU,
	"GRNDSON":                              cpb.V3RoleCode_GRNDSON,
	"GRPRN":                                cpb.V3RoleCode_GRPRN,
	"GRFTH":                                cpb.V3RoleCode_GRFTH,
	"MGRFTH":                               cpb.V3RoleCode_MGRFTH,
	"PGRFTH":                               cpb.V3RoleCode_PGRFTH,
	"GRMTH":                                cpb.V3RoleCode_GRMTH,
	"MGRMTH":                               cpb.V3RoleCode_MGRMTH,
	"PGRMTH":                               cpb.V3RoleCode_PGRMTH,
	"MGRPRN":                               cpb.V3RoleCode_MGRPRN,
	"PGRPRN":                               cpb.V3RoleCode_PGRPRN,
	"INLAW":                                cpb.V3RoleCode_INLAW,
	"CHLDINLAW":                            cpb.V3RoleCode_CHLDINLAW,
	"DAUINLAW":                             cpb.V3RoleCode_DAUINLAW,
	"SONINLAW":                             cpb.V3RoleCode_SONINLAW,
	"PRNINLAW":                             cpb.V3RoleCode_PRNINLAW,
	"FTHINLAW":                             cpb.V3RoleCode_FTHINLAW,
	"MTHINLAW":                             cpb.V3RoleCode_MTHINLAW,
	"SIBINLAW":                             cpb.V3RoleCode_SIBINLAW,
	"BROINLAW":                             cpb.V3RoleCode_BROINLAW,
	"SISINLAW":                             cpb.V3RoleCode_SISINLAW,
	"NIENEPH":                              cpb.V3RoleCode_NIENEPH,
	"NEPHEW":                               cpb.V3RoleCode_NEPHEW,
	"NIECE":                                cpb.V3RoleCode_NIECE,
	"UNCLE":                                cpb.V3RoleCode_UNCLE,
	"MUNCLE":                               cpb.V3RoleCode_MUNCLE,
	"PUNCLE":                               cpb.V3RoleCode_PUNCLE,
	"PRN":                                  cpb.V3RoleCode_PRN,
	"ADOPTP":                               cpb.V3RoleCode_ADOPTP,
	"ADOPTF":                               cpb.V3RoleCode_ADOPTF,
	"ADOPTM":                               cpb.V3RoleCode_ADOPTM,
	"FTH":                                  cpb.V3RoleCode_FTH,
	"FTHFOST":                              cpb.V3RoleCode_FTHFOST,
	"NFTH":                                 cpb.V3RoleCode_NFTH,
	"NFTHF":                                cpb.V3RoleCode_NFTHF,
	"STPFTH":                               cpb.V3RoleCode_STPFTH,
	"MTH":                                  cpb.V3RoleCode_MTH,
	"GESTM":                                cpb.V3RoleCode_GESTM,
	"MTHFOST":                              cpb.V3RoleCode_MTHFOST,
	"NMTH":                                 cpb.V3RoleCode_NMTH,
	"NMTHF":                                cpb.V3RoleCode_NMTHF,
	"STPMTH":                               cpb.V3RoleCode_STPMTH,
	"NPRN":                                 cpb.V3RoleCode_NPRN,
	"PRNFOST":                              cpb.V3RoleCode_PRNFOST,
	"STPPRN":                               cpb.V3RoleCode_STPPRN,
	"SIB":                                  cpb.V3RoleCode_SIB,
	"BRO":                                  cpb.V3RoleCode_BRO,
	"HBRO":                                 cpb.V3RoleCode_HBRO,
	"NBRO":                                 cpb.V3RoleCode_NBRO,
	"TWINBRO":                              cpb.V3RoleCode_TWINBRO,
	"FTWINBRO":                             cpb.V3RoleCode_FTWINBRO,
	"ITWINBRO":                             cpb.V3RoleCode_ITWINBRO,
	"STPBRO":                               cpb.V3RoleCode_STPBRO,
	"HSIB":                                 cpb.V3RoleCode_HSIB,
	"HSIS":                                 cpb.V3RoleCode_HSIS,
	"NSIB":                                 cpb.V3RoleCode_NSIB,
	"NSIS":                                 cpb.V3RoleCode_NSIS,
	"TWINSIS":                              cpb.V3RoleCode_TWINSIS,
	"FTWINSIS":                             cpb.V3RoleCode_FTWINSIS,
	"ITWINSIS":                             cpb.V3RoleCode_ITWINSIS,
	"TWIN":                                 cpb.V3RoleCode_TWIN,
	"FTWIN":                                cpb.V3RoleCode_FTWIN,
	"ITWIN":                                cpb.V3RoleCode_ITWIN,
	"SIS":                                  cpb.V3RoleCode_SIS,
	"STPSIS":                               cpb.V3RoleCode_STPSIS,
	"STPSIB":                               cpb.V3RoleCode_STPSIB,
	"SIGOTHR":                              cpb.V3RoleCode_SIGOTHR,
	"DOMPART":                              cpb.V3RoleCode_DOMPART,
	"FMRSPS":                               cpb.V3RoleCode_FMRSPS,
	"SPS":                                  cpb.V3RoleCode_SPS,
	"HUSB":                                 cpb.V3RoleCode_HUSB,
	"WIFE":                                 cpb.V3RoleCode_WIFE,
	"FRND":                                 cpb.V3RoleCode_FRND,
	"NBOR":                                 cpb.V3RoleCode_NBOR,
	"ONESELF":                              cpb.V3RoleCode_ONESELF,
	"ROOM":                                 cpb.V3RoleCode_ROOM,
	"POLICY_OR_PROGRAM_COVERAGE_ROLE_TYPE": cpb.V3RoleCode_POLICY_OR_PROGRAM_COVERAGE_ROLE_TYPE,
	"COVERAGE_ROLE_TYPE":                   cpb.V3RoleCode_COVERAGE_ROLE_TYPE,
	"FAMDEP":                               cpb.V3RoleCode_FAMDEP,
	"HANDIC":                               cpb.V3RoleCode_HANDIC,
	"INJ":                                  cpb.V3RoleCode_INJ,
	"SELF":                                 cpb.V3RoleCode_SELF,
	"SPON":                                 cpb.V3RoleCode_SPON,
	"STUD":                                 cpb.V3RoleCode_STUD,
	"FSTUD":                                cpb.V3RoleCode_FSTUD,
	"PSTUD":                                cpb.V3RoleCode_PSTUD,
	"ADOPT":                                cpb.V3RoleCode_ADOPT,
	"GCHILD":                               cpb.V3RoleCode_GCHILD,
	"GPARNT":                               cpb.V3RoleCode_GPARNT,
	"NAT":                                  cpb.V3RoleCode_NAT,
	"NIENE":                                cpb.V3RoleCode_NIENE,
	"PARNT":                                cpb.V3RoleCode_PARNT,
	"SPSE":                                 cpb.V3RoleCode_SPSE,
	"STEP":                                 cpb.V3RoleCode_STEP,
	"COVERED_PARTY_ROLE_TYPE":              cpb.V3RoleCode_COVERED_PARTY_ROLE_TYPE,
	"CLAIMANT_COVERED_PARTY_ROLE_TYPE":     cpb.V3RoleCode_CLAIMANT_COVERED_PARTY_ROLE_TYPE,
	"CRIMEVIC":                             cpb.V3RoleCode_CRIMEVIC,
	"INJWKR":                               cpb.V3RoleCode_INJWKR,
	"DEPENDENT_COVERED_PARTY_ROLE_TYPE":    cpb.V3RoleCode_DEPENDENT_COVERED_PARTY_ROLE_TYPE,
	"COCBEN":                               cpb.V3RoleCode_COCBEN,
	"DIFFABL":                              cpb.V3RoleCode_DIFFABL,
	"WARD":                                 cpb.V3RoleCode_WARD,
	"INDIVIDUAL_INSURED_PARTY_ROLE_TYPE":   cpb.V3RoleCode_INDIVIDUAL_INSURED_PARTY_ROLE_TYPE,
	"RETIREE":                              cpb.V3RoleCode_RETIREE,
	"PROGRAM_ELIGIBLE_PARTY_ROLE_TYPE":     cpb.V3RoleCode_PROGRAM_ELIGIBLE_PARTY_ROLE_TYPE,
	"INDIG":                                cpb.V3RoleCode_INDIG,
	"MIL":                                  cpb.V3RoleCode_MIL,
	"ACTMIL":                               cpb.V3RoleCode_ACTMIL,
	"RETMIL":                               cpb.V3RoleCode_RETMIL,
	"VET":                                  cpb.V3RoleCode_VET,
	"SUBSCRIBER_COVERED_PARTY_ROLE_TYPE":   cpb.V3RoleCode_SUBSCRIBER_COVERED_PARTY_ROLE_TYPE,
	"RESEARCH_SUBJECT_ROLE_BASIS":          cpb.V3RoleCode_RESEARCH_SUBJECT_ROLE_BASIS,
	"ERL":                                  cpb.V3RoleCode_ERL,
	"SCN":                                  cpb.V3RoleCode_SCN,
	"SERVICE_DELIVERY_LOCATION_ROLE_TYPE":  cpb.V3RoleCode_SERVICE_DELIVERY_LOCATION_ROLE_TYPE,
	"DEDICATED_SERVICE_DELIVERY_LOCATION_ROLE_TYPE": cpb.V3RoleCode_DEDICATED_SERVICE_DELIVERY_LOCATION_ROLE_TYPE,
	"DEDICATED_CLINICAL_LOCATION_ROLE_TYPE":         cpb.V3RoleCode_DEDICATED_CLINICAL_LOCATION_ROLE_TYPE,
	"DX":                                            cpb.V3RoleCode_DX,
	"CVDX":                                          cpb.V3RoleCode_CVDX,
	"CATH":                                          cpb.V3RoleCode_CATH,
	"ECHO":                                          cpb.V3RoleCode_ECHO,
	"GIDX":                                          cpb.V3RoleCode_GIDX,
	"ENDOS":                                         cpb.V3RoleCode_ENDOS,
	"RADDX":                                         cpb.V3RoleCode_RADDX,
	"RADO":                                          cpb.V3RoleCode_RADO,
	"RNEU":                                          cpb.V3RoleCode_RNEU,
	"HOSP":                                          cpb.V3RoleCode_HOSP,
	"CHR":                                           cpb.V3RoleCode_CHR,
	"GACH":                                          cpb.V3RoleCode_GACH,
	"MHSP":                                          cpb.V3RoleCode_MHSP,
	"PSYCHF":                                        cpb.V3RoleCode_PSYCHF,
	"RH":                                            cpb.V3RoleCode_RH,
	"RHAT":                                          cpb.V3RoleCode_RHAT,
	"RHII":                                          cpb.V3RoleCode_RHII,
	"RHMAD":                                         cpb.V3RoleCode_RHMAD,
	"RHPI":                                          cpb.V3RoleCode_RHPI,
	"RHPIH":                                         cpb.V3RoleCode_RHPIH,
	"RHPIMS":                                        cpb.V3RoleCode_RHPIMS,
	"RHPIVS":                                        cpb.V3RoleCode_RHPIVS,
	"RHYAD":                                         cpb.V3RoleCode_RHYAD,
	"HU":                                            cpb.V3RoleCode_HU,
	"BMTU":                                          cpb.V3RoleCode_BMTU,
	"CCU":                                           cpb.V3RoleCode_CCU,
	"CHEST":                                         cpb.V3RoleCode_CHEST,
	"EPIL":                                          cpb.V3RoleCode_EPIL,
	"ER":                                            cpb.V3RoleCode_ER,
	"ETU":                                           cpb.V3RoleCode_ETU,
	"HD":                                            cpb.V3RoleCode_HD,
	"HLAB":                                          cpb.V3RoleCode_HLAB,
	"INLAB":                                         cpb.V3RoleCode_INLAB,
	"OUTLAB":                                        cpb.V3RoleCode_OUTLAB,
	"HRAD":                                          cpb.V3RoleCode_HRAD,
	"HUSCS":                                         cpb.V3RoleCode_HUSCS,
	"ICU":                                           cpb.V3RoleCode_ICU,
	"PEDICU":                                        cpb.V3RoleCode_PEDICU,
	"PEDNICU":                                       cpb.V3RoleCode_PEDNICU,
	"INPHARM":                                       cpb.V3RoleCode_INPHARM,
	"MBL":                                           cpb.V3RoleCode_MBL,
	"NCCS":                                          cpb.V3RoleCode_NCCS,
	"NS":                                            cpb.V3RoleCode_NS,
	"OUTPHARM":                                      cpb.V3RoleCode_OUTPHARM,
	"PEDU":                                          cpb.V3RoleCode_PEDU,
	"PHU":                                           cpb.V3RoleCode_PHU,
	"RHU":                                           cpb.V3RoleCode_RHU,
	"SLEEP":                                         cpb.V3RoleCode_SLEEP,
	"NCCF":                                          cpb.V3RoleCode_NCCF,
	"SNF":                                           cpb.V3RoleCode_SNF,
	"OF":                                            cpb.V3RoleCode_OF,
	"ALL":                                           cpb.V3RoleCode_ALL,
	"AMPUT":                                         cpb.V3RoleCode_AMPUT,
	"BMTC":                                          cpb.V3RoleCode_BMTC,
	"BREAST":                                        cpb.V3RoleCode_BREAST,
	"CANC":                                          cpb.V3RoleCode_CANC,
	"CAPC":                                          cpb.V3RoleCode_CAPC,
	"CARD":                                          cpb.V3RoleCode_CARD,
	"PEDCARD":                                       cpb.V3RoleCode_PEDCARD,
	"COAG":                                          cpb.V3RoleCode_COAG,
	"CRS":                                           cpb.V3RoleCode_CRS,
	"DERM":                                          cpb.V3RoleCode_DERM,
	"ENDO":                                          cpb.V3RoleCode_ENDO,
	"PEDE":                                          cpb.V3RoleCode_PEDE,
	"ENT":                                           cpb.V3RoleCode_ENT,
	"FMC":                                           cpb.V3RoleCode_FMC,
	"GI":                                            cpb.V3RoleCode_GI,
	"PEDGI":                                         cpb.V3RoleCode_PEDGI,
	"GIM":                                           cpb.V3RoleCode_GIM,
	"GYN":                                           cpb.V3RoleCode_GYN,
	"HEM":                                           cpb.V3RoleCode_HEM,
	"PEDHEM":                                        cpb.V3RoleCode_PEDHEM,
	"HTN":                                           cpb.V3RoleCode_HTN,
	"IEC":                                           cpb.V3RoleCode_IEC,
	"INFD":                                          cpb.V3RoleCode_INFD,
	"PEDID":                                         cpb.V3RoleCode_PEDID,
	"INV":                                           cpb.V3RoleCode_INV,
	"LYMPH":                                         cpb.V3RoleCode_LYMPH,
	"MGEN":                                          cpb.V3RoleCode_MGEN,
	"NEPH":                                          cpb.V3RoleCode_NEPH,
	"PEDNEPH":                                       cpb.V3RoleCode_PEDNEPH,
	"NEUR":                                          cpb.V3RoleCode_NEUR,
	"OB":                                            cpb.V3RoleCode_OB,
	"OMS":                                           cpb.V3RoleCode_OMS,
	"ONCL":                                          cpb.V3RoleCode_ONCL,
	"PEDHO":                                         cpb.V3RoleCode_PEDHO,
	"OPH":                                           cpb.V3RoleCode_OPH,
	"OPTC":                                          cpb.V3RoleCode_OPTC,
	"ORTHO":                                         cpb.V3RoleCode_ORTHO,
	"HAND":                                          cpb.V3RoleCode_HAND,
	"PAINCL":                                        cpb.V3RoleCode_PAINCL,
	"PC":                                            cpb.V3RoleCode_PC,
	"PEDC":                                          cpb.V3RoleCode_PEDC,
	"PEDRHEUM":                                      cpb.V3RoleCode_PEDRHEUM,
	"POD":                                           cpb.V3RoleCode_POD,
	"PREV":                                          cpb.V3RoleCode_PREV,
	"PROCTO":                                        cpb.V3RoleCode_PROCTO,
	"PROFF":                                         cpb.V3RoleCode_PROFF,
	"PROS":                                          cpb.V3RoleCode_PROS,
	"PSI":                                           cpb.V3RoleCode_PSI,
	"PSY":                                           cpb.V3RoleCode_PSY,
	"RHEUM":                                         cpb.V3RoleCode_RHEUM,
	"SPMED":                                         cpb.V3RoleCode_SPMED,
	"SU":                                            cpb.V3RoleCode_SU,
	"PLS":                                           cpb.V3RoleCode_PLS,
	"URO":                                           cpb.V3RoleCode_URO,
	"TR":                                            cpb.V3RoleCode_TR,
	"TRAVEL":                                        cpb.V3RoleCode_TRAVEL,
	"WND":                                           cpb.V3RoleCode_WND,
	"RTF":                                           cpb.V3RoleCode_RTF,
	"PRC":                                           cpb.V3RoleCode_PRC,
	"SURF":                                          cpb.V3RoleCode_SURF,
	"DEDICATED_NON_CLINICAL_LOCATION_ROLE_TYPE": cpb.V3RoleCode_DEDICATED_NON_CLINICAL_LOCATION_ROLE_TYPE,
	"DADDR": cpb.V3RoleCode_DADDR,
	"MOBL":  cpb.V3RoleCode_MOBL,
	"AMB":   cpb.V3RoleCode_AMB,
	"PHARM": cpb.V3RoleCode_PHARM,
	"INCIDENTAL_SERVICE_DELIVERY_LOCATION_ROLE_TYPE": cpb.V3RoleCode_INCIDENTAL_SERVICE_DELIVERY_LOCATION_ROLE_TYPE,
	"ACC":                  cpb.V3RoleCode_ACC,
	"COMM":                 cpb.V3RoleCode_COMM,
	"CSC":                  cpb.V3RoleCode_CSC,
	"PTRES":                cpb.V3RoleCode_PTRES,
	"SCHOOL":               cpb.V3RoleCode_SCHOOL,
	"UPC":                  cpb.V3RoleCode_UPC,
	"WORK":                 cpb.V3RoleCode_WORK,
	"SPECIMEN_ROLE_TYPE":   cpb.V3RoleCode_SPECIMEN_ROLE_TYPE,
	"C":                    cpb.V3RoleCode_C,
	"G":                    cpb.V3RoleCode_G,
	"L":                    cpb.V3RoleCode_L,
	"P":                    cpb.V3RoleCode_P,
	"Q":                    cpb.V3RoleCode_Q,
	"B":                    cpb.V3RoleCode_B,
	"E":                    cpb.V3RoleCode_E,
	"F":                    cpb.V3RoleCode_F,
	"O":                    cpb.V3RoleCode_O,
	"V":                    cpb.V3RoleCode_V,
	"R":                    cpb.V3RoleCode_R,
	"CLAIM":                cpb.V3RoleCode_CLAIM,
	"COMMUNITY_LABORATORY": cpb.V3RoleCode_COMMUNITY_LABORATORY,
	"GT":                   cpb.V3RoleCode_GT,
	"HOME_HEALTH":          cpb.V3RoleCode_HOME_HEALTH,
	"LABORATORY":           cpb.V3RoleCode_LABORATORY,
	"PATHOLOGIST":          cpb.V3RoleCode_PATHOLOGIST,
	"PH":                   cpb.V3RoleCode_PH,
	"PHLEBOTOMIST":         cpb.V3RoleCode_PHLEBOTOMIST,
	"PROG":                 cpb.V3RoleCode_PROG,
	"PT":                   cpb.V3RoleCode_PT,
	"SUBJECT":              cpb.V3RoleCode_SUBJECT,
	"THIRD_PARTY":          cpb.V3RoleCode_THIRD_PARTY,
	"DEP":                  cpb.V3RoleCode_DEP,
	"DEPEN":                cpb.V3RoleCode_DEPEN,
	"FM":                   cpb.V3RoleCode_FM,
	"INDIV":                cpb.V3RoleCode_INDIV,
	"NAMED":                cpb.V3RoleCode_NAMED,
	"PSYCHCF":              cpb.V3RoleCode_PSYCHCF,
	"SUBSCR":               cpb.V3RoleCode_SUBSCR,
}

// DefaultV3TimingEventCodeMap maps from string to cpb.V3TimingEventCode_Value.
var DefaultV3TimingEventCodeMap = map[string]cpb.V3TimingEventCode_Value{
	"INVALID_UNINITIALIZED": cpb.V3TimingEventCode_INVALID_UNINITIALIZED,
	"AC":                    cpb.V3TimingEventCode_AC,
	"ACD":                   cpb.V3TimingEventCode_ACD,
	"ACM":                   cpb.V3TimingEventCode_ACM,
	"ACV":                   cpb.V3TimingEventCode_ACV,
	"C":                     cpb.V3TimingEventCode_C,
	"CD":                    cpb.V3TimingEventCode_CD,
	"CM":                    cpb.V3TimingEventCode_CM,
	"CV":                    cpb.V3TimingEventCode_CV,
	"HS":                    cpb.V3TimingEventCode_HS,
	"IC":                    cpb.V3TimingEventCode_IC,
	"ICD":                   cpb.V3TimingEventCode_ICD,
	"ICM":                   cpb.V3TimingEventCode_ICM,
	"ICV":                   cpb.V3TimingEventCode_ICV,
	"PC":                    cpb.V3TimingEventCode_PC,
	"PCD":                   cpb.V3TimingEventCode_PCD,
	"PCM":                   cpb.V3TimingEventCode_PCM,
	"PCV":                   cpb.V3TimingEventCode_PCV,
	"WAKE":                  cpb.V3TimingEventCode_WAKE,
}

// DefaultVisionBaseCodeMap maps from string to cpb.VisionBaseCode_Value.
var DefaultVisionBaseCodeMap = map[string]cpb.VisionBaseCode_Value{
	"INVALID_UNINITIALIZED": cpb.VisionBaseCode_INVALID_UNINITIALIZED,
	"UP":                    cpb.VisionBaseCode_UP,
	"DOWN":                  cpb.VisionBaseCode_DOWN,
	"IN":                    cpb.VisionBaseCode_IN,
	"OUT":                   cpb.VisionBaseCode_OUT,
}

// DefaultVisionEyesCodeMap maps from string to cpb.VisionEyesCode_Value.
var DefaultVisionEyesCodeMap = map[string]cpb.VisionEyesCode_Value{
	"INVALID_UNINITIALIZED": cpb.VisionEyesCode_INVALID_UNINITIALIZED,
	"RIGHT":                 cpb.VisionEyesCode_RIGHT,
	"LEFT":                  cpb.VisionEyesCode_LEFT,
}

// DefaultXPathUsageTypeCodeMap maps from string to cpb.XPathUsageTypeCode_Value.
var DefaultXPathUsageTypeCodeMap = map[string]cpb.XPathUsageTypeCode_Value{
	"INVALID_UNINITIALIZED": cpb.XPathUsageTypeCode_INVALID_UNINITIALIZED,
	"NORMAL":                cpb.XPathUsageTypeCode_NORMAL,
	"PHONETIC":              cpb.XPathUsageTypeCode_PHONETIC,
	"NEARBY":                cpb.XPathUsageTypeCode_NEARBY,
	"DISTANCE":              cpb.XPathUsageTypeCode_DISTANCE,
	"OTHER":                 cpb.XPathUsageTypeCode_OTHER,
}
