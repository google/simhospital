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

// Package hl7tofhir contains utility functions to convert HL7v2 to FHIR.
package hl7tofhir

import (
	"strings"

	cpb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/codes_go_proto"
)

// AbstractTypeCode converts the given string to a cpb.AbstractTypeCode_Value.
// If the string doesn't match exactly one of the supported values, AbstractTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AbstractTypeCode(s string) cpb.AbstractTypeCode_Value {
	if c.AbstractTypeCodeMap != nil {
		return c.AbstractTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultAbstractTypeCodeMap[strings.ToUpper(s)]
}

// AccountStatusCode converts the given string to a cpb.AccountStatusCode_Value.
// If the string doesn't match exactly one of the supported values, AccountStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AccountStatusCode(s string) cpb.AccountStatusCode_Value {
	if c.AccountStatusCodeMap != nil {
		return c.AccountStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultAccountStatusCodeMap[strings.ToUpper(s)]
}

// ActionCardinalityBehaviorCode converts the given string to a cpb.ActionCardinalityBehaviorCode_Value.
// If the string doesn't match exactly one of the supported values, ActionCardinalityBehaviorCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ActionCardinalityBehaviorCode(s string) cpb.ActionCardinalityBehaviorCode_Value {
	if c.ActionCardinalityBehaviorCodeMap != nil {
		return c.ActionCardinalityBehaviorCodeMap[strings.ToUpper(s)]
	}
	return DefaultActionCardinalityBehaviorCodeMap[strings.ToUpper(s)]
}

// ActionConditionKindCode converts the given string to a cpb.ActionConditionKindCode_Value.
// If the string doesn't match exactly one of the supported values, ActionConditionKindCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ActionConditionKindCode(s string) cpb.ActionConditionKindCode_Value {
	if c.ActionConditionKindCodeMap != nil {
		return c.ActionConditionKindCodeMap[strings.ToUpper(s)]
	}
	return DefaultActionConditionKindCodeMap[strings.ToUpper(s)]
}

// ActionGroupingBehaviorCode converts the given string to a cpb.ActionGroupingBehaviorCode_Value.
// If the string doesn't match exactly one of the supported values, ActionGroupingBehaviorCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ActionGroupingBehaviorCode(s string) cpb.ActionGroupingBehaviorCode_Value {
	if c.ActionGroupingBehaviorCodeMap != nil {
		return c.ActionGroupingBehaviorCodeMap[strings.ToUpper(s)]
	}
	return DefaultActionGroupingBehaviorCodeMap[strings.ToUpper(s)]
}

// ActionParticipantTypeCode converts the given string to a cpb.ActionParticipantTypeCode_Value.
// If the string doesn't match exactly one of the supported values, ActionParticipantTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ActionParticipantTypeCode(s string) cpb.ActionParticipantTypeCode_Value {
	if c.ActionParticipantTypeCodeMap != nil {
		return c.ActionParticipantTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultActionParticipantTypeCodeMap[strings.ToUpper(s)]
}

// ActionPrecheckBehaviorCode converts the given string to a cpb.ActionPrecheckBehaviorCode_Value.
// If the string doesn't match exactly one of the supported values, ActionPrecheckBehaviorCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ActionPrecheckBehaviorCode(s string) cpb.ActionPrecheckBehaviorCode_Value {
	if c.ActionPrecheckBehaviorCodeMap != nil {
		return c.ActionPrecheckBehaviorCodeMap[strings.ToUpper(s)]
	}
	return DefaultActionPrecheckBehaviorCodeMap[strings.ToUpper(s)]
}

// ActionRelationshipTypeCode converts the given string to a cpb.ActionRelationshipTypeCode_Value.
// If the string doesn't match exactly one of the supported values, ActionRelationshipTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ActionRelationshipTypeCode(s string) cpb.ActionRelationshipTypeCode_Value {
	if c.ActionRelationshipTypeCodeMap != nil {
		return c.ActionRelationshipTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultActionRelationshipTypeCodeMap[strings.ToUpper(s)]
}

// ActionRequiredBehaviorCode converts the given string to a cpb.ActionRequiredBehaviorCode_Value.
// If the string doesn't match exactly one of the supported values, ActionRequiredBehaviorCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ActionRequiredBehaviorCode(s string) cpb.ActionRequiredBehaviorCode_Value {
	if c.ActionRequiredBehaviorCodeMap != nil {
		return c.ActionRequiredBehaviorCodeMap[strings.ToUpper(s)]
	}
	return DefaultActionRequiredBehaviorCodeMap[strings.ToUpper(s)]
}

// ActionSelectionBehaviorCode converts the given string to a cpb.ActionSelectionBehaviorCode_Value.
// If the string doesn't match exactly one of the supported values, ActionSelectionBehaviorCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ActionSelectionBehaviorCode(s string) cpb.ActionSelectionBehaviorCode_Value {
	if c.ActionSelectionBehaviorCodeMap != nil {
		return c.ActionSelectionBehaviorCodeMap[strings.ToUpper(s)]
	}
	return DefaultActionSelectionBehaviorCodeMap[strings.ToUpper(s)]
}

// AddressTypeCode converts the given string to a cpb.AddressTypeCode_Value.
// If the string doesn't match exactly one of the supported values, AddressTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AddressTypeCode(s string) cpb.AddressTypeCode_Value {
	if c.AddressTypeCodeMap != nil {
		return c.AddressTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultAddressTypeCodeMap[strings.ToUpper(s)]
}

// AddressUseCode converts the given string to a cpb.AddressUseCode_Value.
// If the string doesn't match exactly one of the supported values, AddressUseCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AddressUseCode(s string) cpb.AddressUseCode_Value {
	if c.AddressUseCodeMap != nil {
		return c.AddressUseCodeMap[strings.ToUpper(s)]
	}
	return DefaultAddressUseCodeMap[strings.ToUpper(s)]
}

// AdministrativeGenderCode converts the given string to a cpb.AdministrativeGenderCode_Value.
// If the string doesn't match exactly one of the supported values, AdministrativeGenderCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AdministrativeGenderCode(s string) cpb.AdministrativeGenderCode_Value {
	if c.AdministrativeGenderCodeMap != nil {
		return c.AdministrativeGenderCodeMap[strings.ToUpper(s)]
	}
	return DefaultAdministrativeGenderCodeMap[strings.ToUpper(s)]
}

// AdverseEventActualityCode converts the given string to a cpb.AdverseEventActualityCode_Value.
// If the string doesn't match exactly one of the supported values, AdverseEventActualityCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AdverseEventActualityCode(s string) cpb.AdverseEventActualityCode_Value {
	if c.AdverseEventActualityCodeMap != nil {
		return c.AdverseEventActualityCodeMap[strings.ToUpper(s)]
	}
	return DefaultAdverseEventActualityCodeMap[strings.ToUpper(s)]
}

// AdverseEventOutcomeCode converts the given string to a cpb.AdverseEventOutcomeCode_Value.
// If the string doesn't match exactly one of the supported values, AdverseEventOutcomeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AdverseEventOutcomeCode(s string) cpb.AdverseEventOutcomeCode_Value {
	if c.AdverseEventOutcomeCodeMap != nil {
		return c.AdverseEventOutcomeCodeMap[strings.ToUpper(s)]
	}
	return DefaultAdverseEventOutcomeCodeMap[strings.ToUpper(s)]
}

// AdverseEventSeverityCode converts the given string to a cpb.AdverseEventSeverityCode_Value.
// If the string doesn't match exactly one of the supported values, AdverseEventSeverityCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AdverseEventSeverityCode(s string) cpb.AdverseEventSeverityCode_Value {
	if c.AdverseEventSeverityCodeMap != nil {
		return c.AdverseEventSeverityCodeMap[strings.ToUpper(s)]
	}
	return DefaultAdverseEventSeverityCodeMap[strings.ToUpper(s)]
}

// AggregationModeCode converts the given string to a cpb.AggregationModeCode_Value.
// If the string doesn't match exactly one of the supported values, AggregationModeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AggregationModeCode(s string) cpb.AggregationModeCode_Value {
	if c.AggregationModeCodeMap != nil {
		return c.AggregationModeCodeMap[strings.ToUpper(s)]
	}
	return DefaultAggregationModeCodeMap[strings.ToUpper(s)]
}

// AllergyIntoleranceCategoryCode converts the given string to a cpb.AllergyIntoleranceCategoryCode_Value.
// If the string doesn't match exactly one of the supported values, AllergyIntoleranceCategoryCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AllergyIntoleranceCategoryCode(s string) cpb.AllergyIntoleranceCategoryCode_Value {
	if c.AllergyIntoleranceCategoryCodeMap != nil {
		return c.AllergyIntoleranceCategoryCodeMap[strings.ToUpper(s)]
	}
	return DefaultAllergyIntoleranceCategoryCodeMap[strings.ToUpper(s)]
}

// AllergyIntoleranceClinicalStatusCode converts the given string to a cpb.AllergyIntoleranceClinicalStatusCode_Value.
// If the string doesn't match exactly one of the supported values, AllergyIntoleranceClinicalStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AllergyIntoleranceClinicalStatusCode(s string) cpb.AllergyIntoleranceClinicalStatusCode_Value {
	if c.AllergyIntoleranceClinicalStatusCodeMap != nil {
		return c.AllergyIntoleranceClinicalStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultAllergyIntoleranceClinicalStatusCodeMap[strings.ToUpper(s)]
}

// AllergyIntoleranceCriticalityCode converts the given string to a cpb.AllergyIntoleranceCriticalityCode_Value.
// If the string doesn't match exactly one of the supported values, AllergyIntoleranceCriticalityCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AllergyIntoleranceCriticalityCode(s string) cpb.AllergyIntoleranceCriticalityCode_Value {
	if c.AllergyIntoleranceCriticalityCodeMap != nil {
		return c.AllergyIntoleranceCriticalityCodeMap[strings.ToUpper(s)]
	}
	return DefaultAllergyIntoleranceCriticalityCodeMap[strings.ToUpper(s)]
}

// AllergyIntoleranceSeverityCode converts the given string to a cpb.AllergyIntoleranceSeverityCode_Value.
// If the string doesn't match exactly one of the supported values, AllergyIntoleranceSeverityCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AllergyIntoleranceSeverityCode(s string) cpb.AllergyIntoleranceSeverityCode_Value {
	if c.AllergyIntoleranceSeverityCodeMap != nil {
		return c.AllergyIntoleranceSeverityCodeMap[strings.ToUpper(s)]
	}
	return DefaultAllergyIntoleranceSeverityCodeMap[strings.ToUpper(s)]
}

// AllergyIntoleranceSubstanceExposureRiskCode converts the given string to a cpb.AllergyIntoleranceSubstanceExposureRiskCode_Value.
// If the string doesn't match exactly one of the supported values, AllergyIntoleranceSubstanceExposureRiskCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AllergyIntoleranceSubstanceExposureRiskCode(s string) cpb.AllergyIntoleranceSubstanceExposureRiskCode_Value {
	if c.AllergyIntoleranceSubstanceExposureRiskCodeMap != nil {
		return c.AllergyIntoleranceSubstanceExposureRiskCodeMap[strings.ToUpper(s)]
	}
	return DefaultAllergyIntoleranceSubstanceExposureRiskCodeMap[strings.ToUpper(s)]
}

// AllergyIntoleranceTypeCode converts the given string to a cpb.AllergyIntoleranceTypeCode_Value.
// If the string doesn't match exactly one of the supported values, AllergyIntoleranceTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AllergyIntoleranceTypeCode(s string) cpb.AllergyIntoleranceTypeCode_Value {
	if c.AllergyIntoleranceTypeCodeMap != nil {
		return c.AllergyIntoleranceTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultAllergyIntoleranceTypeCodeMap[strings.ToUpper(s)]
}

// AllergyIntoleranceVerificationStatusCode converts the given string to a cpb.AllergyIntoleranceVerificationStatusCode_Value.
// If the string doesn't match exactly one of the supported values, AllergyIntoleranceVerificationStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AllergyIntoleranceVerificationStatusCode(s string) cpb.AllergyIntoleranceVerificationStatusCode_Value {
	if c.AllergyIntoleranceVerificationStatusCodeMap != nil {
		return c.AllergyIntoleranceVerificationStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultAllergyIntoleranceVerificationStatusCodeMap[strings.ToUpper(s)]
}

// AppointmentStatusCode converts the given string to a cpb.AppointmentStatusCode_Value.
// If the string doesn't match exactly one of the supported values, AppointmentStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AppointmentStatusCode(s string) cpb.AppointmentStatusCode_Value {
	if c.AppointmentStatusCodeMap != nil {
		return c.AppointmentStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultAppointmentStatusCodeMap[strings.ToUpper(s)]
}

// AssertionDirectionTypeCode converts the given string to a cpb.AssertionDirectionTypeCode_Value.
// If the string doesn't match exactly one of the supported values, AssertionDirectionTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AssertionDirectionTypeCode(s string) cpb.AssertionDirectionTypeCode_Value {
	if c.AssertionDirectionTypeCodeMap != nil {
		return c.AssertionDirectionTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultAssertionDirectionTypeCodeMap[strings.ToUpper(s)]
}

// AssertionOperatorTypeCode converts the given string to a cpb.AssertionOperatorTypeCode_Value.
// If the string doesn't match exactly one of the supported values, AssertionOperatorTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AssertionOperatorTypeCode(s string) cpb.AssertionOperatorTypeCode_Value {
	if c.AssertionOperatorTypeCodeMap != nil {
		return c.AssertionOperatorTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultAssertionOperatorTypeCodeMap[strings.ToUpper(s)]
}

// AssertionResponseTypesCode converts the given string to a cpb.AssertionResponseTypesCode_Value.
// If the string doesn't match exactly one of the supported values, AssertionResponseTypesCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AssertionResponseTypesCode(s string) cpb.AssertionResponseTypesCode_Value {
	if c.AssertionResponseTypesCodeMap != nil {
		return c.AssertionResponseTypesCodeMap[strings.ToUpper(s)]
	}
	return DefaultAssertionResponseTypesCodeMap[strings.ToUpper(s)]
}

// AuditEventActionCode converts the given string to a cpb.AuditEventActionCode_Value.
// If the string doesn't match exactly one of the supported values, AuditEventActionCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AuditEventActionCode(s string) cpb.AuditEventActionCode_Value {
	if c.AuditEventActionCodeMap != nil {
		return c.AuditEventActionCodeMap[strings.ToUpper(s)]
	}
	return DefaultAuditEventActionCodeMap[strings.ToUpper(s)]
}

// AuditEventAgentNetworkTypeCode converts the given string to a cpb.AuditEventAgentNetworkTypeCode_Value.
// If the string doesn't match exactly one of the supported values, AuditEventAgentNetworkTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AuditEventAgentNetworkTypeCode(s string) cpb.AuditEventAgentNetworkTypeCode_Value {
	if c.AuditEventAgentNetworkTypeCodeMap != nil {
		return c.AuditEventAgentNetworkTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultAuditEventAgentNetworkTypeCodeMap[strings.ToUpper(s)]
}

// AuditEventOutcomeCode converts the given string to a cpb.AuditEventOutcomeCode_Value.
// If the string doesn't match exactly one of the supported values, AuditEventOutcomeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) AuditEventOutcomeCode(s string) cpb.AuditEventOutcomeCode_Value {
	if c.AuditEventOutcomeCodeMap != nil {
		return c.AuditEventOutcomeCodeMap[strings.ToUpper(s)]
	}
	return DefaultAuditEventOutcomeCodeMap[strings.ToUpper(s)]
}

// BenefitCostApplicabilityCode converts the given string to a cpb.BenefitCostApplicabilityCode_Value.
// If the string doesn't match exactly one of the supported values, BenefitCostApplicabilityCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) BenefitCostApplicabilityCode(s string) cpb.BenefitCostApplicabilityCode_Value {
	if c.BenefitCostApplicabilityCodeMap != nil {
		return c.BenefitCostApplicabilityCodeMap[strings.ToUpper(s)]
	}
	return DefaultBenefitCostApplicabilityCodeMap[strings.ToUpper(s)]
}

// BindingStrengthCode converts the given string to a cpb.BindingStrengthCode_Value.
// If the string doesn't match exactly one of the supported values, BindingStrengthCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) BindingStrengthCode(s string) cpb.BindingStrengthCode_Value {
	if c.BindingStrengthCodeMap != nil {
		return c.BindingStrengthCodeMap[strings.ToUpper(s)]
	}
	return DefaultBindingStrengthCodeMap[strings.ToUpper(s)]
}

// BiologicallyDerivedProductCategoryCode converts the given string to a cpb.BiologicallyDerivedProductCategoryCode_Value.
// If the string doesn't match exactly one of the supported values, BiologicallyDerivedProductCategoryCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) BiologicallyDerivedProductCategoryCode(s string) cpb.BiologicallyDerivedProductCategoryCode_Value {
	if c.BiologicallyDerivedProductCategoryCodeMap != nil {
		return c.BiologicallyDerivedProductCategoryCodeMap[strings.ToUpper(s)]
	}
	return DefaultBiologicallyDerivedProductCategoryCodeMap[strings.ToUpper(s)]
}

// BiologicallyDerivedProductStatusCode converts the given string to a cpb.BiologicallyDerivedProductStatusCode_Value.
// If the string doesn't match exactly one of the supported values, BiologicallyDerivedProductStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) BiologicallyDerivedProductStatusCode(s string) cpb.BiologicallyDerivedProductStatusCode_Value {
	if c.BiologicallyDerivedProductStatusCodeMap != nil {
		return c.BiologicallyDerivedProductStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultBiologicallyDerivedProductStatusCodeMap[strings.ToUpper(s)]
}

// BiologicallyDerivedProductStorageScaleCode converts the given string to a cpb.BiologicallyDerivedProductStorageScaleCode_Value.
// If the string doesn't match exactly one of the supported values, BiologicallyDerivedProductStorageScaleCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) BiologicallyDerivedProductStorageScaleCode(s string) cpb.BiologicallyDerivedProductStorageScaleCode_Value {
	if c.BiologicallyDerivedProductStorageScaleCodeMap != nil {
		return c.BiologicallyDerivedProductStorageScaleCodeMap[strings.ToUpper(s)]
	}
	return DefaultBiologicallyDerivedProductStorageScaleCodeMap[strings.ToUpper(s)]
}

// BundleTypeCode converts the given string to a cpb.BundleTypeCode_Value.
// If the string doesn't match exactly one of the supported values, BundleTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) BundleTypeCode(s string) cpb.BundleTypeCode_Value {
	if c.BundleTypeCodeMap != nil {
		return c.BundleTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultBundleTypeCodeMap[strings.ToUpper(s)]
}

// CanonicalStatusCodesForFHIRResourcesCode converts the given string to a cpb.CanonicalStatusCodesForFHIRResourcesCode_Value.
// If the string doesn't match exactly one of the supported values, CanonicalStatusCodesForFHIRResourcesCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) CanonicalStatusCodesForFHIRResourcesCode(s string) cpb.CanonicalStatusCodesForFHIRResourcesCode_Value {
	if c.CanonicalStatusCodesForFHIRResourcesCodeMap != nil {
		return c.CanonicalStatusCodesForFHIRResourcesCodeMap[strings.ToUpper(s)]
	}
	return DefaultCanonicalStatusCodesForFHIRResourcesCodeMap[strings.ToUpper(s)]
}

// CapabilityStatementKindCode converts the given string to a cpb.CapabilityStatementKindCode_Value.
// If the string doesn't match exactly one of the supported values, CapabilityStatementKindCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) CapabilityStatementKindCode(s string) cpb.CapabilityStatementKindCode_Value {
	if c.CapabilityStatementKindCodeMap != nil {
		return c.CapabilityStatementKindCodeMap[strings.ToUpper(s)]
	}
	return DefaultCapabilityStatementKindCodeMap[strings.ToUpper(s)]
}

// CarePlanActivityStatusCode converts the given string to a cpb.CarePlanActivityStatusCode_Value.
// If the string doesn't match exactly one of the supported values, CarePlanActivityStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) CarePlanActivityStatusCode(s string) cpb.CarePlanActivityStatusCode_Value {
	if c.CarePlanActivityStatusCodeMap != nil {
		return c.CarePlanActivityStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultCarePlanActivityStatusCodeMap[strings.ToUpper(s)]
}

// CareTeamStatusCode converts the given string to a cpb.CareTeamStatusCode_Value.
// If the string doesn't match exactly one of the supported values, CareTeamStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) CareTeamStatusCode(s string) cpb.CareTeamStatusCode_Value {
	if c.CareTeamStatusCodeMap != nil {
		return c.CareTeamStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultCareTeamStatusCodeMap[strings.ToUpper(s)]
}

// CatalogEntryRelationTypeCode converts the given string to a cpb.CatalogEntryRelationTypeCode_Value.
// If the string doesn't match exactly one of the supported values, CatalogEntryRelationTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) CatalogEntryRelationTypeCode(s string) cpb.CatalogEntryRelationTypeCode_Value {
	if c.CatalogEntryRelationTypeCodeMap != nil {
		return c.CatalogEntryRelationTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultCatalogEntryRelationTypeCodeMap[strings.ToUpper(s)]
}

// ChargeItemStatusCode converts the given string to a cpb.ChargeItemStatusCode_Value.
// If the string doesn't match exactly one of the supported values, ChargeItemStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ChargeItemStatusCode(s string) cpb.ChargeItemStatusCode_Value {
	if c.ChargeItemStatusCodeMap != nil {
		return c.ChargeItemStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultChargeItemStatusCodeMap[strings.ToUpper(s)]
}

// ChoiceListOrientationCode converts the given string to a cpb.ChoiceListOrientationCode_Value.
// If the string doesn't match exactly one of the supported values, ChoiceListOrientationCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ChoiceListOrientationCode(s string) cpb.ChoiceListOrientationCode_Value {
	if c.ChoiceListOrientationCodeMap != nil {
		return c.ChoiceListOrientationCodeMap[strings.ToUpper(s)]
	}
	return DefaultChoiceListOrientationCodeMap[strings.ToUpper(s)]
}

// ClaimProcessingCode converts the given string to a cpb.ClaimProcessingCode_Value.
// If the string doesn't match exactly one of the supported values, ClaimProcessingCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ClaimProcessingCode(s string) cpb.ClaimProcessingCode_Value {
	if c.ClaimProcessingCodeMap != nil {
		return c.ClaimProcessingCodeMap[strings.ToUpper(s)]
	}
	return DefaultClaimProcessingCodeMap[strings.ToUpper(s)]
}

// CodeSearchSupportCode converts the given string to a cpb.CodeSearchSupportCode_Value.
// If the string doesn't match exactly one of the supported values, CodeSearchSupportCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) CodeSearchSupportCode(s string) cpb.CodeSearchSupportCode_Value {
	if c.CodeSearchSupportCodeMap != nil {
		return c.CodeSearchSupportCodeMap[strings.ToUpper(s)]
	}
	return DefaultCodeSearchSupportCodeMap[strings.ToUpper(s)]
}

// CodeSystemContentModeCode converts the given string to a cpb.CodeSystemContentModeCode_Value.
// If the string doesn't match exactly one of the supported values, CodeSystemContentModeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) CodeSystemContentModeCode(s string) cpb.CodeSystemContentModeCode_Value {
	if c.CodeSystemContentModeCodeMap != nil {
		return c.CodeSystemContentModeCodeMap[strings.ToUpper(s)]
	}
	return DefaultCodeSystemContentModeCodeMap[strings.ToUpper(s)]
}

// CodeSystemHierarchyMeaningCode converts the given string to a cpb.CodeSystemHierarchyMeaningCode_Value.
// If the string doesn't match exactly one of the supported values, CodeSystemHierarchyMeaningCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) CodeSystemHierarchyMeaningCode(s string) cpb.CodeSystemHierarchyMeaningCode_Value {
	if c.CodeSystemHierarchyMeaningCodeMap != nil {
		return c.CodeSystemHierarchyMeaningCodeMap[strings.ToUpper(s)]
	}
	return DefaultCodeSystemHierarchyMeaningCodeMap[strings.ToUpper(s)]
}

// CompartmentTypeCode converts the given string to a cpb.CompartmentTypeCode_Value.
// If the string doesn't match exactly one of the supported values, CompartmentTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) CompartmentTypeCode(s string) cpb.CompartmentTypeCode_Value {
	if c.CompartmentTypeCodeMap != nil {
		return c.CompartmentTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultCompartmentTypeCodeMap[strings.ToUpper(s)]
}

// CompositionAttestationModeCode converts the given string to a cpb.CompositionAttestationModeCode_Value.
// If the string doesn't match exactly one of the supported values, CompositionAttestationModeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) CompositionAttestationModeCode(s string) cpb.CompositionAttestationModeCode_Value {
	if c.CompositionAttestationModeCodeMap != nil {
		return c.CompositionAttestationModeCodeMap[strings.ToUpper(s)]
	}
	return DefaultCompositionAttestationModeCodeMap[strings.ToUpper(s)]
}

// CompositionStatusCode converts the given string to a cpb.CompositionStatusCode_Value.
// If the string doesn't match exactly one of the supported values, CompositionStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) CompositionStatusCode(s string) cpb.CompositionStatusCode_Value {
	if c.CompositionStatusCodeMap != nil {
		return c.CompositionStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultCompositionStatusCodeMap[strings.ToUpper(s)]
}

// ConceptMapEquivalenceCode converts the given string to a cpb.ConceptMapEquivalenceCode_Value.
// If the string doesn't match exactly one of the supported values, ConceptMapEquivalenceCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ConceptMapEquivalenceCode(s string) cpb.ConceptMapEquivalenceCode_Value {
	if c.ConceptMapEquivalenceCodeMap != nil {
		return c.ConceptMapEquivalenceCodeMap[strings.ToUpper(s)]
	}
	return DefaultConceptMapEquivalenceCodeMap[strings.ToUpper(s)]
}

// ConceptMapGroupUnmappedModeCode converts the given string to a cpb.ConceptMapGroupUnmappedModeCode_Value.
// If the string doesn't match exactly one of the supported values, ConceptMapGroupUnmappedModeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ConceptMapGroupUnmappedModeCode(s string) cpb.ConceptMapGroupUnmappedModeCode_Value {
	if c.ConceptMapGroupUnmappedModeCodeMap != nil {
		return c.ConceptMapGroupUnmappedModeCodeMap[strings.ToUpper(s)]
	}
	return DefaultConceptMapGroupUnmappedModeCodeMap[strings.ToUpper(s)]
}

// ConditionClinicalStatusCode converts the given string to a cpb.ConditionClinicalStatusCode_Value.
// If the string doesn't match exactly one of the supported values, ConditionClinicalStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ConditionClinicalStatusCode(s string) cpb.ConditionClinicalStatusCode_Value {
	if c.ConditionClinicalStatusCodeMap != nil {
		return c.ConditionClinicalStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultConditionClinicalStatusCodeMap[strings.ToUpper(s)]
}

// ConditionVerificationStatusCode converts the given string to a cpb.ConditionVerificationStatusCode_Value.
// If the string doesn't match exactly one of the supported values, ConditionVerificationStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ConditionVerificationStatusCode(s string) cpb.ConditionVerificationStatusCode_Value {
	if c.ConditionVerificationStatusCodeMap != nil {
		return c.ConditionVerificationStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultConditionVerificationStatusCodeMap[strings.ToUpper(s)]
}

// ConditionalDeleteStatusCode converts the given string to a cpb.ConditionalDeleteStatusCode_Value.
// If the string doesn't match exactly one of the supported values, ConditionalDeleteStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ConditionalDeleteStatusCode(s string) cpb.ConditionalDeleteStatusCode_Value {
	if c.ConditionalDeleteStatusCodeMap != nil {
		return c.ConditionalDeleteStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultConditionalDeleteStatusCodeMap[strings.ToUpper(s)]
}

// ConditionalReadStatusCode converts the given string to a cpb.ConditionalReadStatusCode_Value.
// If the string doesn't match exactly one of the supported values, ConditionalReadStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ConditionalReadStatusCode(s string) cpb.ConditionalReadStatusCode_Value {
	if c.ConditionalReadStatusCodeMap != nil {
		return c.ConditionalReadStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultConditionalReadStatusCodeMap[strings.ToUpper(s)]
}

// ConformanceExpectationCode converts the given string to a cpb.ConformanceExpectationCode_Value.
// If the string doesn't match exactly one of the supported values, ConformanceExpectationCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ConformanceExpectationCode(s string) cpb.ConformanceExpectationCode_Value {
	if c.ConformanceExpectationCodeMap != nil {
		return c.ConformanceExpectationCodeMap[strings.ToUpper(s)]
	}
	return DefaultConformanceExpectationCodeMap[strings.ToUpper(s)]
}

// ConsentDataMeaningCode converts the given string to a cpb.ConsentDataMeaningCode_Value.
// If the string doesn't match exactly one of the supported values, ConsentDataMeaningCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ConsentDataMeaningCode(s string) cpb.ConsentDataMeaningCode_Value {
	if c.ConsentDataMeaningCodeMap != nil {
		return c.ConsentDataMeaningCodeMap[strings.ToUpper(s)]
	}
	return DefaultConsentDataMeaningCodeMap[strings.ToUpper(s)]
}

// ConsentProvisionTypeCode converts the given string to a cpb.ConsentProvisionTypeCode_Value.
// If the string doesn't match exactly one of the supported values, ConsentProvisionTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ConsentProvisionTypeCode(s string) cpb.ConsentProvisionTypeCode_Value {
	if c.ConsentProvisionTypeCodeMap != nil {
		return c.ConsentProvisionTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultConsentProvisionTypeCodeMap[strings.ToUpper(s)]
}

// ConsentStateCode converts the given string to a cpb.ConsentStateCode_Value.
// If the string doesn't match exactly one of the supported values, ConsentStateCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ConsentStateCode(s string) cpb.ConsentStateCode_Value {
	if c.ConsentStateCodeMap != nil {
		return c.ConsentStateCodeMap[strings.ToUpper(s)]
	}
	return DefaultConsentStateCodeMap[strings.ToUpper(s)]
}

// ConstraintSeverityCode converts the given string to a cpb.ConstraintSeverityCode_Value.
// If the string doesn't match exactly one of the supported values, ConstraintSeverityCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ConstraintSeverityCode(s string) cpb.ConstraintSeverityCode_Value {
	if c.ConstraintSeverityCodeMap != nil {
		return c.ConstraintSeverityCodeMap[strings.ToUpper(s)]
	}
	return DefaultConstraintSeverityCodeMap[strings.ToUpper(s)]
}

// ContactPointSystemCode converts the given string to a cpb.ContactPointSystemCode_Value.
// If the string doesn't match exactly one of the supported values, ContactPointSystemCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ContactPointSystemCode(s string) cpb.ContactPointSystemCode_Value {
	if c.ContactPointSystemCodeMap != nil {
		return c.ContactPointSystemCodeMap[strings.ToUpper(s)]
	}
	return DefaultContactPointSystemCodeMap[strings.ToUpper(s)]
}

// ContactPointUseCode converts the given string to a cpb.ContactPointUseCode_Value.
// If the string doesn't match exactly one of the supported values, ContactPointUseCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ContactPointUseCode(s string) cpb.ContactPointUseCode_Value {
	if c.ContactPointUseCodeMap != nil {
		return c.ContactPointUseCodeMap[strings.ToUpper(s)]
	}
	return DefaultContactPointUseCodeMap[strings.ToUpper(s)]
}

// ContractResourcePublicationStatusCode converts the given string to a cpb.ContractResourcePublicationStatusCode_Value.
// If the string doesn't match exactly one of the supported values, ContractResourcePublicationStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ContractResourcePublicationStatusCode(s string) cpb.ContractResourcePublicationStatusCode_Value {
	if c.ContractResourcePublicationStatusCodeMap != nil {
		return c.ContractResourcePublicationStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultContractResourcePublicationStatusCodeMap[strings.ToUpper(s)]
}

// ContractResourceStatusCode converts the given string to a cpb.ContractResourceStatusCode_Value.
// If the string doesn't match exactly one of the supported values, ContractResourceStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ContractResourceStatusCode(s string) cpb.ContractResourceStatusCode_Value {
	if c.ContractResourceStatusCodeMap != nil {
		return c.ContractResourceStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultContractResourceStatusCodeMap[strings.ToUpper(s)]
}

// ContributorTypeCode converts the given string to a cpb.ContributorTypeCode_Value.
// If the string doesn't match exactly one of the supported values, ContributorTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ContributorTypeCode(s string) cpb.ContributorTypeCode_Value {
	if c.ContributorTypeCodeMap != nil {
		return c.ContributorTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultContributorTypeCodeMap[strings.ToUpper(s)]
}

// DataAbsentReasonCode converts the given string to a cpb.DataAbsentReasonCode_Value.
// If the string doesn't match exactly one of the supported values, DataAbsentReasonCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) DataAbsentReasonCode(s string) cpb.DataAbsentReasonCode_Value {
	if c.DataAbsentReasonCodeMap != nil {
		return c.DataAbsentReasonCodeMap[strings.ToUpper(s)]
	}
	return DefaultDataAbsentReasonCodeMap[strings.ToUpper(s)]
}

// DataTypeCode converts the given string to a cpb.DataTypeCode_Value.
// If the string doesn't match exactly one of the supported values, DataTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) DataTypeCode(s string) cpb.DataTypeCode_Value {
	if c.DataTypeCodeMap != nil {
		return c.DataTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultDataTypeCodeMap[strings.ToUpper(s)]
}

// DaysOfWeekCode converts the given string to a cpb.DaysOfWeekCode_Value.
// If the string doesn't match exactly one of the supported values, DaysOfWeekCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) DaysOfWeekCode(s string) cpb.DaysOfWeekCode_Value {
	if c.DaysOfWeekCodeMap != nil {
		return c.DaysOfWeekCodeMap[strings.ToUpper(s)]
	}
	return DefaultDaysOfWeekCodeMap[strings.ToUpper(s)]
}

// DetectedIssueSeverityCode converts the given string to a cpb.DetectedIssueSeverityCode_Value.
// If the string doesn't match exactly one of the supported values, DetectedIssueSeverityCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) DetectedIssueSeverityCode(s string) cpb.DetectedIssueSeverityCode_Value {
	if c.DetectedIssueSeverityCodeMap != nil {
		return c.DetectedIssueSeverityCodeMap[strings.ToUpper(s)]
	}
	return DefaultDetectedIssueSeverityCodeMap[strings.ToUpper(s)]
}

// DeviceMetricCalibrationStateCode converts the given string to a cpb.DeviceMetricCalibrationStateCode_Value.
// If the string doesn't match exactly one of the supported values, DeviceMetricCalibrationStateCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) DeviceMetricCalibrationStateCode(s string) cpb.DeviceMetricCalibrationStateCode_Value {
	if c.DeviceMetricCalibrationStateCodeMap != nil {
		return c.DeviceMetricCalibrationStateCodeMap[strings.ToUpper(s)]
	}
	return DefaultDeviceMetricCalibrationStateCodeMap[strings.ToUpper(s)]
}

// DeviceMetricCalibrationTypeCode converts the given string to a cpb.DeviceMetricCalibrationTypeCode_Value.
// If the string doesn't match exactly one of the supported values, DeviceMetricCalibrationTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) DeviceMetricCalibrationTypeCode(s string) cpb.DeviceMetricCalibrationTypeCode_Value {
	if c.DeviceMetricCalibrationTypeCodeMap != nil {
		return c.DeviceMetricCalibrationTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultDeviceMetricCalibrationTypeCodeMap[strings.ToUpper(s)]
}

// DeviceMetricCategoryCode converts the given string to a cpb.DeviceMetricCategoryCode_Value.
// If the string doesn't match exactly one of the supported values, DeviceMetricCategoryCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) DeviceMetricCategoryCode(s string) cpb.DeviceMetricCategoryCode_Value {
	if c.DeviceMetricCategoryCodeMap != nil {
		return c.DeviceMetricCategoryCodeMap[strings.ToUpper(s)]
	}
	return DefaultDeviceMetricCategoryCodeMap[strings.ToUpper(s)]
}

// DeviceMetricColorCode converts the given string to a cpb.DeviceMetricColorCode_Value.
// If the string doesn't match exactly one of the supported values, DeviceMetricColorCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) DeviceMetricColorCode(s string) cpb.DeviceMetricColorCode_Value {
	if c.DeviceMetricColorCodeMap != nil {
		return c.DeviceMetricColorCodeMap[strings.ToUpper(s)]
	}
	return DefaultDeviceMetricColorCodeMap[strings.ToUpper(s)]
}

// DeviceMetricOperationalStatusCode converts the given string to a cpb.DeviceMetricOperationalStatusCode_Value.
// If the string doesn't match exactly one of the supported values, DeviceMetricOperationalStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) DeviceMetricOperationalStatusCode(s string) cpb.DeviceMetricOperationalStatusCode_Value {
	if c.DeviceMetricOperationalStatusCodeMap != nil {
		return c.DeviceMetricOperationalStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultDeviceMetricOperationalStatusCodeMap[strings.ToUpper(s)]
}

// DeviceNameTypeCode converts the given string to a cpb.DeviceNameTypeCode_Value.
// If the string doesn't match exactly one of the supported values, DeviceNameTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) DeviceNameTypeCode(s string) cpb.DeviceNameTypeCode_Value {
	if c.DeviceNameTypeCodeMap != nil {
		return c.DeviceNameTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultDeviceNameTypeCodeMap[strings.ToUpper(s)]
}

// DeviceUseStatementStatusCode converts the given string to a cpb.DeviceUseStatementStatusCode_Value.
// If the string doesn't match exactly one of the supported values, DeviceUseStatementStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) DeviceUseStatementStatusCode(s string) cpb.DeviceUseStatementStatusCode_Value {
	if c.DeviceUseStatementStatusCodeMap != nil {
		return c.DeviceUseStatementStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultDeviceUseStatementStatusCodeMap[strings.ToUpper(s)]
}

// DiagnosticReportStatusCode converts the given string to a cpb.DiagnosticReportStatusCode_Value.
// If the string doesn't match exactly one of the supported values, DiagnosticReportStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) DiagnosticReportStatusCode(s string) cpb.DiagnosticReportStatusCode_Value {
	if c.DiagnosticReportStatusCodeMap != nil {
		return c.DiagnosticReportStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultDiagnosticReportStatusCodeMap[strings.ToUpper(s)]
}

// DiscriminatorTypeCode converts the given string to a cpb.DiscriminatorTypeCode_Value.
// If the string doesn't match exactly one of the supported values, DiscriminatorTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) DiscriminatorTypeCode(s string) cpb.DiscriminatorTypeCode_Value {
	if c.DiscriminatorTypeCodeMap != nil {
		return c.DiscriminatorTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultDiscriminatorTypeCodeMap[strings.ToUpper(s)]
}

// DocumentModeCode converts the given string to a cpb.DocumentModeCode_Value.
// If the string doesn't match exactly one of the supported values, DocumentModeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) DocumentModeCode(s string) cpb.DocumentModeCode_Value {
	if c.DocumentModeCodeMap != nil {
		return c.DocumentModeCodeMap[strings.ToUpper(s)]
	}
	return DefaultDocumentModeCodeMap[strings.ToUpper(s)]
}

// DocumentReferenceStatusCode converts the given string to a cpb.DocumentReferenceStatusCode_Value.
// If the string doesn't match exactly one of the supported values, DocumentReferenceStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) DocumentReferenceStatusCode(s string) cpb.DocumentReferenceStatusCode_Value {
	if c.DocumentReferenceStatusCodeMap != nil {
		return c.DocumentReferenceStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultDocumentReferenceStatusCodeMap[strings.ToUpper(s)]
}

// DocumentRelationshipTypeCode converts the given string to a cpb.DocumentRelationshipTypeCode_Value.
// If the string doesn't match exactly one of the supported values, DocumentRelationshipTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) DocumentRelationshipTypeCode(s string) cpb.DocumentRelationshipTypeCode_Value {
	if c.DocumentRelationshipTypeCodeMap != nil {
		return c.DocumentRelationshipTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultDocumentRelationshipTypeCodeMap[strings.ToUpper(s)]
}

// EligibilityRequestPurposeCode converts the given string to a cpb.EligibilityRequestPurposeCode_Value.
// If the string doesn't match exactly one of the supported values, EligibilityRequestPurposeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) EligibilityRequestPurposeCode(s string) cpb.EligibilityRequestPurposeCode_Value {
	if c.EligibilityRequestPurposeCodeMap != nil {
		return c.EligibilityRequestPurposeCodeMap[strings.ToUpper(s)]
	}
	return DefaultEligibilityRequestPurposeCodeMap[strings.ToUpper(s)]
}

// EligibilityResponsePurposeCode converts the given string to a cpb.EligibilityResponsePurposeCode_Value.
// If the string doesn't match exactly one of the supported values, EligibilityResponsePurposeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) EligibilityResponsePurposeCode(s string) cpb.EligibilityResponsePurposeCode_Value {
	if c.EligibilityResponsePurposeCodeMap != nil {
		return c.EligibilityResponsePurposeCodeMap[strings.ToUpper(s)]
	}
	return DefaultEligibilityResponsePurposeCodeMap[strings.ToUpper(s)]
}

// EnableWhenBehaviorCode converts the given string to a cpb.EnableWhenBehaviorCode_Value.
// If the string doesn't match exactly one of the supported values, EnableWhenBehaviorCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) EnableWhenBehaviorCode(s string) cpb.EnableWhenBehaviorCode_Value {
	if c.EnableWhenBehaviorCodeMap != nil {
		return c.EnableWhenBehaviorCodeMap[strings.ToUpper(s)]
	}
	return DefaultEnableWhenBehaviorCodeMap[strings.ToUpper(s)]
}

// EncounterLocationStatusCode converts the given string to a cpb.EncounterLocationStatusCode_Value.
// If the string doesn't match exactly one of the supported values, EncounterLocationStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) EncounterLocationStatusCode(s string) cpb.EncounterLocationStatusCode_Value {
	if c.EncounterLocationStatusCodeMap != nil {
		return c.EncounterLocationStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultEncounterLocationStatusCodeMap[strings.ToUpper(s)]
}

// EncounterStatusCode converts the given string to a cpb.EncounterStatusCode_Value.
// If the string doesn't match exactly one of the supported values, EncounterStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) EncounterStatusCode(s string) cpb.EncounterStatusCode_Value {
	if c.EncounterStatusCodeMap != nil {
		return c.EncounterStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultEncounterStatusCodeMap[strings.ToUpper(s)]
}

// EndpointStatusCode converts the given string to a cpb.EndpointStatusCode_Value.
// If the string doesn't match exactly one of the supported values, EndpointStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) EndpointStatusCode(s string) cpb.EndpointStatusCode_Value {
	if c.EndpointStatusCodeMap != nil {
		return c.EndpointStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultEndpointStatusCodeMap[strings.ToUpper(s)]
}

// EpisodeOfCareStatusCode converts the given string to a cpb.EpisodeOfCareStatusCode_Value.
// If the string doesn't match exactly one of the supported values, EpisodeOfCareStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) EpisodeOfCareStatusCode(s string) cpb.EpisodeOfCareStatusCode_Value {
	if c.EpisodeOfCareStatusCodeMap != nil {
		return c.EpisodeOfCareStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultEpisodeOfCareStatusCodeMap[strings.ToUpper(s)]
}

// EventCapabilityModeCode converts the given string to a cpb.EventCapabilityModeCode_Value.
// If the string doesn't match exactly one of the supported values, EventCapabilityModeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) EventCapabilityModeCode(s string) cpb.EventCapabilityModeCode_Value {
	if c.EventCapabilityModeCodeMap != nil {
		return c.EventCapabilityModeCodeMap[strings.ToUpper(s)]
	}
	return DefaultEventCapabilityModeCodeMap[strings.ToUpper(s)]
}

// EventStatusCode converts the given string to a cpb.EventStatusCode_Value.
// If the string doesn't match exactly one of the supported values, EventStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) EventStatusCode(s string) cpb.EventStatusCode_Value {
	if c.EventStatusCodeMap != nil {
		return c.EventStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultEventStatusCodeMap[strings.ToUpper(s)]
}

// EventTimingCode converts the given string to a cpb.EventTimingCode_Value.
// If the string doesn't match exactly one of the supported values, EventTimingCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) EventTimingCode(s string) cpb.EventTimingCode_Value {
	if c.EventTimingCodeMap != nil {
		return c.EventTimingCodeMap[strings.ToUpper(s)]
	}
	return DefaultEventTimingCodeMap[strings.ToUpper(s)]
}

// EvidenceVariableTypeCode converts the given string to a cpb.EvidenceVariableTypeCode_Value.
// If the string doesn't match exactly one of the supported values, EvidenceVariableTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) EvidenceVariableTypeCode(s string) cpb.EvidenceVariableTypeCode_Value {
	if c.EvidenceVariableTypeCodeMap != nil {
		return c.EvidenceVariableTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultEvidenceVariableTypeCodeMap[strings.ToUpper(s)]
}

// ExampleScenarioActorTypeCode converts the given string to a cpb.ExampleScenarioActorTypeCode_Value.
// If the string doesn't match exactly one of the supported values, ExampleScenarioActorTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ExampleScenarioActorTypeCode(s string) cpb.ExampleScenarioActorTypeCode_Value {
	if c.ExampleScenarioActorTypeCodeMap != nil {
		return c.ExampleScenarioActorTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultExampleScenarioActorTypeCodeMap[strings.ToUpper(s)]
}

// ExpansionParameterSourceCode converts the given string to a cpb.ExpansionParameterSourceCode_Value.
// If the string doesn't match exactly one of the supported values, ExpansionParameterSourceCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ExpansionParameterSourceCode(s string) cpb.ExpansionParameterSourceCode_Value {
	if c.ExpansionParameterSourceCodeMap != nil {
		return c.ExpansionParameterSourceCodeMap[strings.ToUpper(s)]
	}
	return DefaultExpansionParameterSourceCodeMap[strings.ToUpper(s)]
}

// ExpansionProcessingRuleCode converts the given string to a cpb.ExpansionProcessingRuleCode_Value.
// If the string doesn't match exactly one of the supported values, ExpansionProcessingRuleCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ExpansionProcessingRuleCode(s string) cpb.ExpansionProcessingRuleCode_Value {
	if c.ExpansionProcessingRuleCodeMap != nil {
		return c.ExpansionProcessingRuleCodeMap[strings.ToUpper(s)]
	}
	return DefaultExpansionProcessingRuleCodeMap[strings.ToUpper(s)]
}

// ExplanationOfBenefitStatusCode converts the given string to a cpb.ExplanationOfBenefitStatusCode_Value.
// If the string doesn't match exactly one of the supported values, ExplanationOfBenefitStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ExplanationOfBenefitStatusCode(s string) cpb.ExplanationOfBenefitStatusCode_Value {
	if c.ExplanationOfBenefitStatusCodeMap != nil {
		return c.ExplanationOfBenefitStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultExplanationOfBenefitStatusCodeMap[strings.ToUpper(s)]
}

// ExposureStateCode converts the given string to a cpb.ExposureStateCode_Value.
// If the string doesn't match exactly one of the supported values, ExposureStateCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ExposureStateCode(s string) cpb.ExposureStateCode_Value {
	if c.ExposureStateCodeMap != nil {
		return c.ExposureStateCodeMap[strings.ToUpper(s)]
	}
	return DefaultExposureStateCodeMap[strings.ToUpper(s)]
}

// ExtensionContextTypeCode converts the given string to a cpb.ExtensionContextTypeCode_Value.
// If the string doesn't match exactly one of the supported values, ExtensionContextTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ExtensionContextTypeCode(s string) cpb.ExtensionContextTypeCode_Value {
	if c.ExtensionContextTypeCodeMap != nil {
		return c.ExtensionContextTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultExtensionContextTypeCodeMap[strings.ToUpper(s)]
}

// FHIRDeviceStatusCode converts the given string to a cpb.FHIRDeviceStatusCode_Value.
// If the string doesn't match exactly one of the supported values, FHIRDeviceStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) FHIRDeviceStatusCode(s string) cpb.FHIRDeviceStatusCode_Value {
	if c.FHIRDeviceStatusCodeMap != nil {
		return c.FHIRDeviceStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultFHIRDeviceStatusCodeMap[strings.ToUpper(s)]
}

// FHIRRestfulInteractionsCode converts the given string to a cpb.FHIRRestfulInteractionsCode_Value.
// If the string doesn't match exactly one of the supported values, FHIRRestfulInteractionsCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) FHIRRestfulInteractionsCode(s string) cpb.FHIRRestfulInteractionsCode_Value {
	if c.FHIRRestfulInteractionsCodeMap != nil {
		return c.FHIRRestfulInteractionsCodeMap[strings.ToUpper(s)]
	}
	return DefaultFHIRRestfulInteractionsCodeMap[strings.ToUpper(s)]
}

// FHIRSubstanceStatusCode converts the given string to a cpb.FHIRSubstanceStatusCode_Value.
// If the string doesn't match exactly one of the supported values, FHIRSubstanceStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) FHIRSubstanceStatusCode(s string) cpb.FHIRSubstanceStatusCode_Value {
	if c.FHIRSubstanceStatusCodeMap != nil {
		return c.FHIRSubstanceStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultFHIRSubstanceStatusCodeMap[strings.ToUpper(s)]
}

// FHIRVersionCode converts the given string to a cpb.FHIRVersionCode_Value.
// If the string doesn't match exactly one of the supported values, FHIRVersionCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) FHIRVersionCode(s string) cpb.FHIRVersionCode_Value {
	if c.FHIRVersionCodeMap != nil {
		return c.FHIRVersionCodeMap[strings.ToUpper(s)]
	}
	return DefaultFHIRVersionCodeMap[strings.ToUpper(s)]
}

// FamilyHistoryStatusCode converts the given string to a cpb.FamilyHistoryStatusCode_Value.
// If the string doesn't match exactly one of the supported values, FamilyHistoryStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) FamilyHistoryStatusCode(s string) cpb.FamilyHistoryStatusCode_Value {
	if c.FamilyHistoryStatusCodeMap != nil {
		return c.FamilyHistoryStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultFamilyHistoryStatusCodeMap[strings.ToUpper(s)]
}

// FilterOperatorCode converts the given string to a cpb.FilterOperatorCode_Value.
// If the string doesn't match exactly one of the supported values, FilterOperatorCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) FilterOperatorCode(s string) cpb.FilterOperatorCode_Value {
	if c.FilterOperatorCodeMap != nil {
		return c.FilterOperatorCodeMap[strings.ToUpper(s)]
	}
	return DefaultFilterOperatorCodeMap[strings.ToUpper(s)]
}

// FinancialResourceStatusCode converts the given string to a cpb.FinancialResourceStatusCode_Value.
// If the string doesn't match exactly one of the supported values, FinancialResourceStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) FinancialResourceStatusCode(s string) cpb.FinancialResourceStatusCode_Value {
	if c.FinancialResourceStatusCodeMap != nil {
		return c.FinancialResourceStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultFinancialResourceStatusCodeMap[strings.ToUpper(s)]
}

// FlagStatusCode converts the given string to a cpb.FlagStatusCode_Value.
// If the string doesn't match exactly one of the supported values, FlagStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) FlagStatusCode(s string) cpb.FlagStatusCode_Value {
	if c.FlagStatusCodeMap != nil {
		return c.FlagStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultFlagStatusCodeMap[strings.ToUpper(s)]
}

// GoalAcceptanceStatusCode converts the given string to a cpb.GoalAcceptanceStatusCode_Value.
// If the string doesn't match exactly one of the supported values, GoalAcceptanceStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) GoalAcceptanceStatusCode(s string) cpb.GoalAcceptanceStatusCode_Value {
	if c.GoalAcceptanceStatusCodeMap != nil {
		return c.GoalAcceptanceStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultGoalAcceptanceStatusCodeMap[strings.ToUpper(s)]
}

// GoalLifecycleStatusCode converts the given string to a cpb.GoalLifecycleStatusCode_Value.
// If the string doesn't match exactly one of the supported values, GoalLifecycleStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) GoalLifecycleStatusCode(s string) cpb.GoalLifecycleStatusCode_Value {
	if c.GoalLifecycleStatusCodeMap != nil {
		return c.GoalLifecycleStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultGoalLifecycleStatusCodeMap[strings.ToUpper(s)]
}

// GraphCompartmentRuleCode converts the given string to a cpb.GraphCompartmentRuleCode_Value.
// If the string doesn't match exactly one of the supported values, GraphCompartmentRuleCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) GraphCompartmentRuleCode(s string) cpb.GraphCompartmentRuleCode_Value {
	if c.GraphCompartmentRuleCodeMap != nil {
		return c.GraphCompartmentRuleCodeMap[strings.ToUpper(s)]
	}
	return DefaultGraphCompartmentRuleCodeMap[strings.ToUpper(s)]
}

// GraphCompartmentUseCode converts the given string to a cpb.GraphCompartmentUseCode_Value.
// If the string doesn't match exactly one of the supported values, GraphCompartmentUseCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) GraphCompartmentUseCode(s string) cpb.GraphCompartmentUseCode_Value {
	if c.GraphCompartmentUseCodeMap != nil {
		return c.GraphCompartmentUseCodeMap[strings.ToUpper(s)]
	}
	return DefaultGraphCompartmentUseCodeMap[strings.ToUpper(s)]
}

// GroupMeasureCode converts the given string to a cpb.GroupMeasureCode_Value.
// If the string doesn't match exactly one of the supported values, GroupMeasureCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) GroupMeasureCode(s string) cpb.GroupMeasureCode_Value {
	if c.GroupMeasureCodeMap != nil {
		return c.GroupMeasureCodeMap[strings.ToUpper(s)]
	}
	return DefaultGroupMeasureCodeMap[strings.ToUpper(s)]
}

// GroupTypeCode converts the given string to a cpb.GroupTypeCode_Value.
// If the string doesn't match exactly one of the supported values, GroupTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) GroupTypeCode(s string) cpb.GroupTypeCode_Value {
	if c.GroupTypeCodeMap != nil {
		return c.GroupTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultGroupTypeCodeMap[strings.ToUpper(s)]
}

// GuidanceResponseStatusCode converts the given string to a cpb.GuidanceResponseStatusCode_Value.
// If the string doesn't match exactly one of the supported values, GuidanceResponseStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) GuidanceResponseStatusCode(s string) cpb.GuidanceResponseStatusCode_Value {
	if c.GuidanceResponseStatusCodeMap != nil {
		return c.GuidanceResponseStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultGuidanceResponseStatusCodeMap[strings.ToUpper(s)]
}

// GuidePageGenerationCode converts the given string to a cpb.GuidePageGenerationCode_Value.
// If the string doesn't match exactly one of the supported values, GuidePageGenerationCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) GuidePageGenerationCode(s string) cpb.GuidePageGenerationCode_Value {
	if c.GuidePageGenerationCodeMap != nil {
		return c.GuidePageGenerationCodeMap[strings.ToUpper(s)]
	}
	return DefaultGuidePageGenerationCodeMap[strings.ToUpper(s)]
}

// GuideParameterCode converts the given string to a cpb.GuideParameterCode_Value.
// If the string doesn't match exactly one of the supported values, GuideParameterCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) GuideParameterCode(s string) cpb.GuideParameterCode_Value {
	if c.GuideParameterCodeMap != nil {
		return c.GuideParameterCodeMap[strings.ToUpper(s)]
	}
	return DefaultGuideParameterCodeMap[strings.ToUpper(s)]
}

// HL7WorkgroupCode converts the given string to a cpb.HL7WorkgroupCode_Value.
// If the string doesn't match exactly one of the supported values, HL7WorkgroupCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) HL7WorkgroupCode(s string) cpb.HL7WorkgroupCode_Value {
	if c.HL7WorkgroupCodeMap != nil {
		return c.HL7WorkgroupCodeMap[strings.ToUpper(s)]
	}
	return DefaultHL7WorkgroupCodeMap[strings.ToUpper(s)]
}

// HTTPVerbCode converts the given string to a cpb.HTTPVerbCode_Value.
// If the string doesn't match exactly one of the supported values, HTTPVerbCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) HTTPVerbCode(s string) cpb.HTTPVerbCode_Value {
	if c.HTTPVerbCodeMap != nil {
		return c.HTTPVerbCodeMap[strings.ToUpper(s)]
	}
	return DefaultHTTPVerbCodeMap[strings.ToUpper(s)]
}

// HumanNameAssemblyOrderCode converts the given string to a cpb.HumanNameAssemblyOrderCode_Value.
// If the string doesn't match exactly one of the supported values, HumanNameAssemblyOrderCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) HumanNameAssemblyOrderCode(s string) cpb.HumanNameAssemblyOrderCode_Value {
	if c.HumanNameAssemblyOrderCodeMap != nil {
		return c.HumanNameAssemblyOrderCodeMap[strings.ToUpper(s)]
	}
	return DefaultHumanNameAssemblyOrderCodeMap[strings.ToUpper(s)]
}

// IdentifierUseCode converts the given string to a cpb.IdentifierUseCode_Value.
// If the string doesn't match exactly one of the supported values, IdentifierUseCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) IdentifierUseCode(s string) cpb.IdentifierUseCode_Value {
	if c.IdentifierUseCodeMap != nil {
		return c.IdentifierUseCodeMap[strings.ToUpper(s)]
	}
	return DefaultIdentifierUseCodeMap[strings.ToUpper(s)]
}

// IdentityAssuranceLevelCode converts the given string to a cpb.IdentityAssuranceLevelCode_Value.
// If the string doesn't match exactly one of the supported values, IdentityAssuranceLevelCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) IdentityAssuranceLevelCode(s string) cpb.IdentityAssuranceLevelCode_Value {
	if c.IdentityAssuranceLevelCodeMap != nil {
		return c.IdentityAssuranceLevelCodeMap[strings.ToUpper(s)]
	}
	return DefaultIdentityAssuranceLevelCodeMap[strings.ToUpper(s)]
}

// ImagingStudyStatusCode converts the given string to a cpb.ImagingStudyStatusCode_Value.
// If the string doesn't match exactly one of the supported values, ImagingStudyStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ImagingStudyStatusCode(s string) cpb.ImagingStudyStatusCode_Value {
	if c.ImagingStudyStatusCodeMap != nil {
		return c.ImagingStudyStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultImagingStudyStatusCodeMap[strings.ToUpper(s)]
}

// ImplantStatusCode converts the given string to a cpb.ImplantStatusCode_Value.
// If the string doesn't match exactly one of the supported values, ImplantStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ImplantStatusCode(s string) cpb.ImplantStatusCode_Value {
	if c.ImplantStatusCodeMap != nil {
		return c.ImplantStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultImplantStatusCodeMap[strings.ToUpper(s)]
}

// InvoicePriceComponentTypeCode converts the given string to a cpb.InvoicePriceComponentTypeCode_Value.
// If the string doesn't match exactly one of the supported values, InvoicePriceComponentTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) InvoicePriceComponentTypeCode(s string) cpb.InvoicePriceComponentTypeCode_Value {
	if c.InvoicePriceComponentTypeCodeMap != nil {
		return c.InvoicePriceComponentTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultInvoicePriceComponentTypeCodeMap[strings.ToUpper(s)]
}

// InvoiceStatusCode converts the given string to a cpb.InvoiceStatusCode_Value.
// If the string doesn't match exactly one of the supported values, InvoiceStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) InvoiceStatusCode(s string) cpb.InvoiceStatusCode_Value {
	if c.InvoiceStatusCodeMap != nil {
		return c.InvoiceStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultInvoiceStatusCodeMap[strings.ToUpper(s)]
}

// IssueSeverityCode converts the given string to a cpb.IssueSeverityCode_Value.
// If the string doesn't match exactly one of the supported values, IssueSeverityCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) IssueSeverityCode(s string) cpb.IssueSeverityCode_Value {
	if c.IssueSeverityCodeMap != nil {
		return c.IssueSeverityCodeMap[strings.ToUpper(s)]
	}
	return DefaultIssueSeverityCodeMap[strings.ToUpper(s)]
}

// IssueTypeCode converts the given string to a cpb.IssueTypeCode_Value.
// If the string doesn't match exactly one of the supported values, IssueTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) IssueTypeCode(s string) cpb.IssueTypeCode_Value {
	if c.IssueTypeCodeMap != nil {
		return c.IssueTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultIssueTypeCodeMap[strings.ToUpper(s)]
}

// LinkTypeCode converts the given string to a cpb.LinkTypeCode_Value.
// If the string doesn't match exactly one of the supported values, LinkTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) LinkTypeCode(s string) cpb.LinkTypeCode_Value {
	if c.LinkTypeCodeMap != nil {
		return c.LinkTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultLinkTypeCodeMap[strings.ToUpper(s)]
}

// LinkageTypeCode converts the given string to a cpb.LinkageTypeCode_Value.
// If the string doesn't match exactly one of the supported values, LinkageTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) LinkageTypeCode(s string) cpb.LinkageTypeCode_Value {
	if c.LinkageTypeCodeMap != nil {
		return c.LinkageTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultLinkageTypeCodeMap[strings.ToUpper(s)]
}

// ListModeCode converts the given string to a cpb.ListModeCode_Value.
// If the string doesn't match exactly one of the supported values, ListModeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ListModeCode(s string) cpb.ListModeCode_Value {
	if c.ListModeCodeMap != nil {
		return c.ListModeCodeMap[strings.ToUpper(s)]
	}
	return DefaultListModeCodeMap[strings.ToUpper(s)]
}

// ListStatusCode converts the given string to a cpb.ListStatusCode_Value.
// If the string doesn't match exactly one of the supported values, ListStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ListStatusCode(s string) cpb.ListStatusCode_Value {
	if c.ListStatusCodeMap != nil {
		return c.ListStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultListStatusCodeMap[strings.ToUpper(s)]
}

// LocationModeCode converts the given string to a cpb.LocationModeCode_Value.
// If the string doesn't match exactly one of the supported values, LocationModeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) LocationModeCode(s string) cpb.LocationModeCode_Value {
	if c.LocationModeCodeMap != nil {
		return c.LocationModeCodeMap[strings.ToUpper(s)]
	}
	return DefaultLocationModeCodeMap[strings.ToUpper(s)]
}

// LocationStatusCode converts the given string to a cpb.LocationStatusCode_Value.
// If the string doesn't match exactly one of the supported values, LocationStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) LocationStatusCode(s string) cpb.LocationStatusCode_Value {
	if c.LocationStatusCodeMap != nil {
		return c.LocationStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultLocationStatusCodeMap[strings.ToUpper(s)]
}

// MatchGradeCode converts the given string to a cpb.MatchGradeCode_Value.
// If the string doesn't match exactly one of the supported values, MatchGradeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) MatchGradeCode(s string) cpb.MatchGradeCode_Value {
	if c.MatchGradeCodeMap != nil {
		return c.MatchGradeCodeMap[strings.ToUpper(s)]
	}
	return DefaultMatchGradeCodeMap[strings.ToUpper(s)]
}

// MeasureImprovementNotationCode converts the given string to a cpb.MeasureImprovementNotationCode_Value.
// If the string doesn't match exactly one of the supported values, MeasureImprovementNotationCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) MeasureImprovementNotationCode(s string) cpb.MeasureImprovementNotationCode_Value {
	if c.MeasureImprovementNotationCodeMap != nil {
		return c.MeasureImprovementNotationCodeMap[strings.ToUpper(s)]
	}
	return DefaultMeasureImprovementNotationCodeMap[strings.ToUpper(s)]
}

// MeasureReportStatusCode converts the given string to a cpb.MeasureReportStatusCode_Value.
// If the string doesn't match exactly one of the supported values, MeasureReportStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) MeasureReportStatusCode(s string) cpb.MeasureReportStatusCode_Value {
	if c.MeasureReportStatusCodeMap != nil {
		return c.MeasureReportStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultMeasureReportStatusCodeMap[strings.ToUpper(s)]
}

// MeasureReportTypeCode converts the given string to a cpb.MeasureReportTypeCode_Value.
// If the string doesn't match exactly one of the supported values, MeasureReportTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) MeasureReportTypeCode(s string) cpb.MeasureReportTypeCode_Value {
	if c.MeasureReportTypeCodeMap != nil {
		return c.MeasureReportTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultMeasureReportTypeCodeMap[strings.ToUpper(s)]
}

// MedicationAdministrationStatusCode converts the given string to a cpb.MedicationAdministrationStatusCode_Value.
// If the string doesn't match exactly one of the supported values, MedicationAdministrationStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) MedicationAdministrationStatusCode(s string) cpb.MedicationAdministrationStatusCode_Value {
	if c.MedicationAdministrationStatusCodeMap != nil {
		return c.MedicationAdministrationStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultMedicationAdministrationStatusCodeMap[strings.ToUpper(s)]
}

// MedicationDispenseStatusCode converts the given string to a cpb.MedicationDispenseStatusCode_Value.
// If the string doesn't match exactly one of the supported values, MedicationDispenseStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) MedicationDispenseStatusCode(s string) cpb.MedicationDispenseStatusCode_Value {
	if c.MedicationDispenseStatusCodeMap != nil {
		return c.MedicationDispenseStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultMedicationDispenseStatusCodeMap[strings.ToUpper(s)]
}

// MedicationKnowledgeStatusCode converts the given string to a cpb.MedicationKnowledgeStatusCode_Value.
// If the string doesn't match exactly one of the supported values, MedicationKnowledgeStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) MedicationKnowledgeStatusCode(s string) cpb.MedicationKnowledgeStatusCode_Value {
	if c.MedicationKnowledgeStatusCodeMap != nil {
		return c.MedicationKnowledgeStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultMedicationKnowledgeStatusCodeMap[strings.ToUpper(s)]
}

// MedicationRequestIntentCode converts the given string to a cpb.MedicationRequestIntentCode_Value.
// If the string doesn't match exactly one of the supported values, MedicationRequestIntentCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) MedicationRequestIntentCode(s string) cpb.MedicationRequestIntentCode_Value {
	if c.MedicationRequestIntentCodeMap != nil {
		return c.MedicationRequestIntentCodeMap[strings.ToUpper(s)]
	}
	return DefaultMedicationRequestIntentCodeMap[strings.ToUpper(s)]
}

// MedicationStatementStatusCodes converts the given string to a cpb.MedicationStatementStatusCodes_Value.
// If the string doesn't match exactly one of the supported values, MedicationStatementStatusCodes returns
// INVALID_UNINITIALIZED.
func (c *Convertor) MedicationStatementStatusCodes(s string) cpb.MedicationStatementStatusCodes_Value {
	if c.MedicationStatementStatusCodesMap != nil {
		return c.MedicationStatementStatusCodesMap[strings.ToUpper(s)]
	}
	return DefaultMedicationStatementStatusCodesMap[strings.ToUpper(s)]
}

// MedicationStatusCode converts the given string to a cpb.MedicationStatusCode_Value.
// If the string doesn't match exactly one of the supported values, MedicationStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) MedicationStatusCode(s string) cpb.MedicationStatusCode_Value {
	if c.MedicationStatusCodeMap != nil {
		return c.MedicationStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultMedicationStatusCodeMap[strings.ToUpper(s)]
}

// MedicationrequestStatusCode converts the given string to a cpb.MedicationrequestStatusCode_Value.
// If the string doesn't match exactly one of the supported values, MedicationrequestStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) MedicationrequestStatusCode(s string) cpb.MedicationrequestStatusCode_Value {
	if c.MedicationrequestStatusCodeMap != nil {
		return c.MedicationrequestStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultMedicationrequestStatusCodeMap[strings.ToUpper(s)]
}

// MessageSignificanceCategoryCode converts the given string to a cpb.MessageSignificanceCategoryCode_Value.
// If the string doesn't match exactly one of the supported values, MessageSignificanceCategoryCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) MessageSignificanceCategoryCode(s string) cpb.MessageSignificanceCategoryCode_Value {
	if c.MessageSignificanceCategoryCodeMap != nil {
		return c.MessageSignificanceCategoryCodeMap[strings.ToUpper(s)]
	}
	return DefaultMessageSignificanceCategoryCodeMap[strings.ToUpper(s)]
}

// MessageheaderResponseRequestCode converts the given string to a cpb.MessageheaderResponseRequestCode_Value.
// If the string doesn't match exactly one of the supported values, MessageheaderResponseRequestCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) MessageheaderResponseRequestCode(s string) cpb.MessageheaderResponseRequestCode_Value {
	if c.MessageheaderResponseRequestCodeMap != nil {
		return c.MessageheaderResponseRequestCodeMap[strings.ToUpper(s)]
	}
	return DefaultMessageheaderResponseRequestCodeMap[strings.ToUpper(s)]
}

// NameUseCode converts the given string to a cpb.NameUseCode_Value.
// If the string doesn't match exactly one of the supported values, NameUseCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) NameUseCode(s string) cpb.NameUseCode_Value {
	if c.NameUseCodeMap != nil {
		return c.NameUseCodeMap[strings.ToUpper(s)]
	}
	return DefaultNameUseCodeMap[strings.ToUpper(s)]
}

// NamingSystemIdentifierTypeCode converts the given string to a cpb.NamingSystemIdentifierTypeCode_Value.
// If the string doesn't match exactly one of the supported values, NamingSystemIdentifierTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) NamingSystemIdentifierTypeCode(s string) cpb.NamingSystemIdentifierTypeCode_Value {
	if c.NamingSystemIdentifierTypeCodeMap != nil {
		return c.NamingSystemIdentifierTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultNamingSystemIdentifierTypeCodeMap[strings.ToUpper(s)]
}

// NamingSystemTypeCode converts the given string to a cpb.NamingSystemTypeCode_Value.
// If the string doesn't match exactly one of the supported values, NamingSystemTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) NamingSystemTypeCode(s string) cpb.NamingSystemTypeCode_Value {
	if c.NamingSystemTypeCodeMap != nil {
		return c.NamingSystemTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultNamingSystemTypeCodeMap[strings.ToUpper(s)]
}

// NarrativeStatusCode converts the given string to a cpb.NarrativeStatusCode_Value.
// If the string doesn't match exactly one of the supported values, NarrativeStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) NarrativeStatusCode(s string) cpb.NarrativeStatusCode_Value {
	if c.NarrativeStatusCodeMap != nil {
		return c.NarrativeStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultNarrativeStatusCodeMap[strings.ToUpper(s)]
}

// NoteTypeCode converts the given string to a cpb.NoteTypeCode_Value.
// If the string doesn't match exactly one of the supported values, NoteTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) NoteTypeCode(s string) cpb.NoteTypeCode_Value {
	if c.NoteTypeCodeMap != nil {
		return c.NoteTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultNoteTypeCodeMap[strings.ToUpper(s)]
}

// ObservationDataTypeCode converts the given string to a cpb.ObservationDataTypeCode_Value.
// If the string doesn't match exactly one of the supported values, ObservationDataTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ObservationDataTypeCode(s string) cpb.ObservationDataTypeCode_Value {
	if c.ObservationDataTypeCodeMap != nil {
		return c.ObservationDataTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultObservationDataTypeCodeMap[strings.ToUpper(s)]
}

// ObservationRangeCategoryCode converts the given string to a cpb.ObservationRangeCategoryCode_Value.
// If the string doesn't match exactly one of the supported values, ObservationRangeCategoryCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ObservationRangeCategoryCode(s string) cpb.ObservationRangeCategoryCode_Value {
	if c.ObservationRangeCategoryCodeMap != nil {
		return c.ObservationRangeCategoryCodeMap[strings.ToUpper(s)]
	}
	return DefaultObservationRangeCategoryCodeMap[strings.ToUpper(s)]
}

// ObservationStatusCode converts the given string to a cpb.ObservationStatusCode_Value.
// If the string doesn't match exactly one of the supported values, ObservationStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ObservationStatusCode(s string) cpb.ObservationStatusCode_Value {
	if c.ObservationStatusCodeMap != nil {
		return c.ObservationStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultObservationStatusCodeMap[strings.ToUpper(s)]
}

// OperationKindCode converts the given string to a cpb.OperationKindCode_Value.
// If the string doesn't match exactly one of the supported values, OperationKindCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) OperationKindCode(s string) cpb.OperationKindCode_Value {
	if c.OperationKindCodeMap != nil {
		return c.OperationKindCodeMap[strings.ToUpper(s)]
	}
	return DefaultOperationKindCodeMap[strings.ToUpper(s)]
}

// OperationParameterUseCode converts the given string to a cpb.OperationParameterUseCode_Value.
// If the string doesn't match exactly one of the supported values, OperationParameterUseCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) OperationParameterUseCode(s string) cpb.OperationParameterUseCode_Value {
	if c.OperationParameterUseCodeMap != nil {
		return c.OperationParameterUseCodeMap[strings.ToUpper(s)]
	}
	return DefaultOperationParameterUseCodeMap[strings.ToUpper(s)]
}

// OrientationTypeCode converts the given string to a cpb.OrientationTypeCode_Value.
// If the string doesn't match exactly one of the supported values, OrientationTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) OrientationTypeCode(s string) cpb.OrientationTypeCode_Value {
	if c.OrientationTypeCodeMap != nil {
		return c.OrientationTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultOrientationTypeCodeMap[strings.ToUpper(s)]
}

// ParticipantRequiredCode converts the given string to a cpb.ParticipantRequiredCode_Value.
// If the string doesn't match exactly one of the supported values, ParticipantRequiredCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ParticipantRequiredCode(s string) cpb.ParticipantRequiredCode_Value {
	if c.ParticipantRequiredCodeMap != nil {
		return c.ParticipantRequiredCodeMap[strings.ToUpper(s)]
	}
	return DefaultParticipantRequiredCodeMap[strings.ToUpper(s)]
}

// ParticipationStatusCode converts the given string to a cpb.ParticipationStatusCode_Value.
// If the string doesn't match exactly one of the supported values, ParticipationStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ParticipationStatusCode(s string) cpb.ParticipationStatusCode_Value {
	if c.ParticipationStatusCodeMap != nil {
		return c.ParticipationStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultParticipationStatusCodeMap[strings.ToUpper(s)]
}

// PropertyRepresentationCode converts the given string to a cpb.PropertyRepresentationCode_Value.
// If the string doesn't match exactly one of the supported values, PropertyRepresentationCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) PropertyRepresentationCode(s string) cpb.PropertyRepresentationCode_Value {
	if c.PropertyRepresentationCodeMap != nil {
		return c.PropertyRepresentationCodeMap[strings.ToUpper(s)]
	}
	return DefaultPropertyRepresentationCodeMap[strings.ToUpper(s)]
}

// PropertyTypeCode converts the given string to a cpb.PropertyTypeCode_Value.
// If the string doesn't match exactly one of the supported values, PropertyTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) PropertyTypeCode(s string) cpb.PropertyTypeCode_Value {
	if c.PropertyTypeCodeMap != nil {
		return c.PropertyTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultPropertyTypeCodeMap[strings.ToUpper(s)]
}

// ProvenanceEntityRoleCode converts the given string to a cpb.ProvenanceEntityRoleCode_Value.
// If the string doesn't match exactly one of the supported values, ProvenanceEntityRoleCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ProvenanceEntityRoleCode(s string) cpb.ProvenanceEntityRoleCode_Value {
	if c.ProvenanceEntityRoleCodeMap != nil {
		return c.ProvenanceEntityRoleCodeMap[strings.ToUpper(s)]
	}
	return DefaultProvenanceEntityRoleCodeMap[strings.ToUpper(s)]
}

// PublicationStatusCode converts the given string to a cpb.PublicationStatusCode_Value.
// If the string doesn't match exactly one of the supported values, PublicationStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) PublicationStatusCode(s string) cpb.PublicationStatusCode_Value {
	if c.PublicationStatusCodeMap != nil {
		return c.PublicationStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultPublicationStatusCodeMap[strings.ToUpper(s)]
}

// QualityTypeCode converts the given string to a cpb.QualityTypeCode_Value.
// If the string doesn't match exactly one of the supported values, QualityTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) QualityTypeCode(s string) cpb.QualityTypeCode_Value {
	if c.QualityTypeCodeMap != nil {
		return c.QualityTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultQualityTypeCodeMap[strings.ToUpper(s)]
}

// QuantityComparatorCode converts the given string to a cpb.QuantityComparatorCode_Value.
// If the string doesn't match exactly one of the supported values, QuantityComparatorCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) QuantityComparatorCode(s string) cpb.QuantityComparatorCode_Value {
	if c.QuantityComparatorCodeMap != nil {
		return c.QuantityComparatorCodeMap[strings.ToUpper(s)]
	}
	return DefaultQuantityComparatorCodeMap[strings.ToUpper(s)]
}

// QuestionnaireItemOperatorCode converts the given string to a cpb.QuestionnaireItemOperatorCode_Value.
// If the string doesn't match exactly one of the supported values, QuestionnaireItemOperatorCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) QuestionnaireItemOperatorCode(s string) cpb.QuestionnaireItemOperatorCode_Value {
	if c.QuestionnaireItemOperatorCodeMap != nil {
		return c.QuestionnaireItemOperatorCodeMap[strings.ToUpper(s)]
	}
	return DefaultQuestionnaireItemOperatorCodeMap[strings.ToUpper(s)]
}

// QuestionnaireItemTypeCode converts the given string to a cpb.QuestionnaireItemTypeCode_Value.
// If the string doesn't match exactly one of the supported values, QuestionnaireItemTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) QuestionnaireItemTypeCode(s string) cpb.QuestionnaireItemTypeCode_Value {
	if c.QuestionnaireItemTypeCodeMap != nil {
		return c.QuestionnaireItemTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultQuestionnaireItemTypeCodeMap[strings.ToUpper(s)]
}

// QuestionnaireItemUsageModeCode converts the given string to a cpb.QuestionnaireItemUsageModeCode_Value.
// If the string doesn't match exactly one of the supported values, QuestionnaireItemUsageModeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) QuestionnaireItemUsageModeCode(s string) cpb.QuestionnaireItemUsageModeCode_Value {
	if c.QuestionnaireItemUsageModeCodeMap != nil {
		return c.QuestionnaireItemUsageModeCodeMap[strings.ToUpper(s)]
	}
	return DefaultQuestionnaireItemUsageModeCodeMap[strings.ToUpper(s)]
}

// QuestionnaireResponseStatusCode converts the given string to a cpb.QuestionnaireResponseStatusCode_Value.
// If the string doesn't match exactly one of the supported values, QuestionnaireResponseStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) QuestionnaireResponseStatusCode(s string) cpb.QuestionnaireResponseStatusCode_Value {
	if c.QuestionnaireResponseStatusCodeMap != nil {
		return c.QuestionnaireResponseStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultQuestionnaireResponseStatusCodeMap[strings.ToUpper(s)]
}

// ReferenceHandlingPolicyCode converts the given string to a cpb.ReferenceHandlingPolicyCode_Value.
// If the string doesn't match exactly one of the supported values, ReferenceHandlingPolicyCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ReferenceHandlingPolicyCode(s string) cpb.ReferenceHandlingPolicyCode_Value {
	if c.ReferenceHandlingPolicyCodeMap != nil {
		return c.ReferenceHandlingPolicyCodeMap[strings.ToUpper(s)]
	}
	return DefaultReferenceHandlingPolicyCodeMap[strings.ToUpper(s)]
}

// ReferenceVersionRulesCode converts the given string to a cpb.ReferenceVersionRulesCode_Value.
// If the string doesn't match exactly one of the supported values, ReferenceVersionRulesCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ReferenceVersionRulesCode(s string) cpb.ReferenceVersionRulesCode_Value {
	if c.ReferenceVersionRulesCodeMap != nil {
		return c.ReferenceVersionRulesCodeMap[strings.ToUpper(s)]
	}
	return DefaultReferenceVersionRulesCodeMap[strings.ToUpper(s)]
}

// RelatedArtifactTypeCode converts the given string to a cpb.RelatedArtifactTypeCode_Value.
// If the string doesn't match exactly one of the supported values, RelatedArtifactTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) RelatedArtifactTypeCode(s string) cpb.RelatedArtifactTypeCode_Value {
	if c.RelatedArtifactTypeCodeMap != nil {
		return c.RelatedArtifactTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultRelatedArtifactTypeCodeMap[strings.ToUpper(s)]
}

// RepositoryTypeCode converts the given string to a cpb.RepositoryTypeCode_Value.
// If the string doesn't match exactly one of the supported values, RepositoryTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) RepositoryTypeCode(s string) cpb.RepositoryTypeCode_Value {
	if c.RepositoryTypeCodeMap != nil {
		return c.RepositoryTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultRepositoryTypeCodeMap[strings.ToUpper(s)]
}

// RequestIntentCode converts the given string to a cpb.RequestIntentCode_Value.
// If the string doesn't match exactly one of the supported values, RequestIntentCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) RequestIntentCode(s string) cpb.RequestIntentCode_Value {
	if c.RequestIntentCodeMap != nil {
		return c.RequestIntentCodeMap[strings.ToUpper(s)]
	}
	return DefaultRequestIntentCodeMap[strings.ToUpper(s)]
}

// RequestPriorityCode converts the given string to a cpb.RequestPriorityCode_Value.
// If the string doesn't match exactly one of the supported values, RequestPriorityCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) RequestPriorityCode(s string) cpb.RequestPriorityCode_Value {
	if c.RequestPriorityCodeMap != nil {
		return c.RequestPriorityCodeMap[strings.ToUpper(s)]
	}
	return DefaultRequestPriorityCodeMap[strings.ToUpper(s)]
}

// RequestResourceTypeCode converts the given string to a cpb.RequestResourceTypeCode_Value.
// If the string doesn't match exactly one of the supported values, RequestResourceTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) RequestResourceTypeCode(s string) cpb.RequestResourceTypeCode_Value {
	if c.RequestResourceTypeCodeMap != nil {
		return c.RequestResourceTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultRequestResourceTypeCodeMap[strings.ToUpper(s)]
}

// RequestStatusCode converts the given string to a cpb.RequestStatusCode_Value.
// If the string doesn't match exactly one of the supported values, RequestStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) RequestStatusCode(s string) cpb.RequestStatusCode_Value {
	if c.RequestStatusCodeMap != nil {
		return c.RequestStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultRequestStatusCodeMap[strings.ToUpper(s)]
}

// ResearchElementTypeCode converts the given string to a cpb.ResearchElementTypeCode_Value.
// If the string doesn't match exactly one of the supported values, ResearchElementTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ResearchElementTypeCode(s string) cpb.ResearchElementTypeCode_Value {
	if c.ResearchElementTypeCodeMap != nil {
		return c.ResearchElementTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultResearchElementTypeCodeMap[strings.ToUpper(s)]
}

// ResearchStudyStatusCode converts the given string to a cpb.ResearchStudyStatusCode_Value.
// If the string doesn't match exactly one of the supported values, ResearchStudyStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ResearchStudyStatusCode(s string) cpb.ResearchStudyStatusCode_Value {
	if c.ResearchStudyStatusCodeMap != nil {
		return c.ResearchStudyStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultResearchStudyStatusCodeMap[strings.ToUpper(s)]
}

// ResearchSubjectStatusCode converts the given string to a cpb.ResearchSubjectStatusCode_Value.
// If the string doesn't match exactly one of the supported values, ResearchSubjectStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ResearchSubjectStatusCode(s string) cpb.ResearchSubjectStatusCode_Value {
	if c.ResearchSubjectStatusCodeMap != nil {
		return c.ResearchSubjectStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultResearchSubjectStatusCodeMap[strings.ToUpper(s)]
}

// ResourceSecurityCategoryCode converts the given string to a cpb.ResourceSecurityCategoryCode_Value.
// If the string doesn't match exactly one of the supported values, ResourceSecurityCategoryCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ResourceSecurityCategoryCode(s string) cpb.ResourceSecurityCategoryCode_Value {
	if c.ResourceSecurityCategoryCodeMap != nil {
		return c.ResourceSecurityCategoryCodeMap[strings.ToUpper(s)]
	}
	return DefaultResourceSecurityCategoryCodeMap[strings.ToUpper(s)]
}

// ResourceTypeCode converts the given string to a cpb.ResourceTypeCode_Value.
// If the string doesn't match exactly one of the supported values, ResourceTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ResourceTypeCode(s string) cpb.ResourceTypeCode_Value {
	if c.ResourceTypeCodeMap != nil {
		return c.ResourceTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultResourceTypeCodeMap[strings.ToUpper(s)]
}

// ResourceVersionPolicyCode converts the given string to a cpb.ResourceVersionPolicyCode_Value.
// If the string doesn't match exactly one of the supported values, ResourceVersionPolicyCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ResourceVersionPolicyCode(s string) cpb.ResourceVersionPolicyCode_Value {
	if c.ResourceVersionPolicyCodeMap != nil {
		return c.ResourceVersionPolicyCodeMap[strings.ToUpper(s)]
	}
	return DefaultResourceVersionPolicyCodeMap[strings.ToUpper(s)]
}

// ResponseTypeCode converts the given string to a cpb.ResponseTypeCode_Value.
// If the string doesn't match exactly one of the supported values, ResponseTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) ResponseTypeCode(s string) cpb.ResponseTypeCode_Value {
	if c.ResponseTypeCodeMap != nil {
		return c.ResponseTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultResponseTypeCodeMap[strings.ToUpper(s)]
}

// RestfulCapabilityModeCode converts the given string to a cpb.RestfulCapabilityModeCode_Value.
// If the string doesn't match exactly one of the supported values, RestfulCapabilityModeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) RestfulCapabilityModeCode(s string) cpb.RestfulCapabilityModeCode_Value {
	if c.RestfulCapabilityModeCodeMap != nil {
		return c.RestfulCapabilityModeCodeMap[strings.ToUpper(s)]
	}
	return DefaultRestfulCapabilityModeCodeMap[strings.ToUpper(s)]
}

// SPDXLicenseCode converts the given string to a cpb.SPDXLicenseCode_Value.
// If the string doesn't match exactly one of the supported values, SPDXLicenseCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) SPDXLicenseCode(s string) cpb.SPDXLicenseCode_Value {
	if c.SPDXLicenseCodeMap != nil {
		return c.SPDXLicenseCodeMap[strings.ToUpper(s)]
	}
	return DefaultSPDXLicenseCodeMap[strings.ToUpper(s)]
}

// SearchComparatorCode converts the given string to a cpb.SearchComparatorCode_Value.
// If the string doesn't match exactly one of the supported values, SearchComparatorCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) SearchComparatorCode(s string) cpb.SearchComparatorCode_Value {
	if c.SearchComparatorCodeMap != nil {
		return c.SearchComparatorCodeMap[strings.ToUpper(s)]
	}
	return DefaultSearchComparatorCodeMap[strings.ToUpper(s)]
}

// SearchEntryModeCode converts the given string to a cpb.SearchEntryModeCode_Value.
// If the string doesn't match exactly one of the supported values, SearchEntryModeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) SearchEntryModeCode(s string) cpb.SearchEntryModeCode_Value {
	if c.SearchEntryModeCodeMap != nil {
		return c.SearchEntryModeCodeMap[strings.ToUpper(s)]
	}
	return DefaultSearchEntryModeCodeMap[strings.ToUpper(s)]
}

// SearchModifierCode converts the given string to a cpb.SearchModifierCode_Value.
// If the string doesn't match exactly one of the supported values, SearchModifierCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) SearchModifierCode(s string) cpb.SearchModifierCode_Value {
	if c.SearchModifierCodeMap != nil {
		return c.SearchModifierCodeMap[strings.ToUpper(s)]
	}
	return DefaultSearchModifierCodeMap[strings.ToUpper(s)]
}

// SearchParamTypeCode converts the given string to a cpb.SearchParamTypeCode_Value.
// If the string doesn't match exactly one of the supported values, SearchParamTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) SearchParamTypeCode(s string) cpb.SearchParamTypeCode_Value {
	if c.SearchParamTypeCodeMap != nil {
		return c.SearchParamTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultSearchParamTypeCodeMap[strings.ToUpper(s)]
}

// SequenceTypeCode converts the given string to a cpb.SequenceTypeCode_Value.
// If the string doesn't match exactly one of the supported values, SequenceTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) SequenceTypeCode(s string) cpb.SequenceTypeCode_Value {
	if c.SequenceTypeCodeMap != nil {
		return c.SequenceTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultSequenceTypeCodeMap[strings.ToUpper(s)]
}

// SlicingRulesCode converts the given string to a cpb.SlicingRulesCode_Value.
// If the string doesn't match exactly one of the supported values, SlicingRulesCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) SlicingRulesCode(s string) cpb.SlicingRulesCode_Value {
	if c.SlicingRulesCodeMap != nil {
		return c.SlicingRulesCodeMap[strings.ToUpper(s)]
	}
	return DefaultSlicingRulesCodeMap[strings.ToUpper(s)]
}

// SlotStatusCode converts the given string to a cpb.SlotStatusCode_Value.
// If the string doesn't match exactly one of the supported values, SlotStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) SlotStatusCode(s string) cpb.SlotStatusCode_Value {
	if c.SlotStatusCodeMap != nil {
		return c.SlotStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultSlotStatusCodeMap[strings.ToUpper(s)]
}

// SmartCapabilitiesCode converts the given string to a cpb.SmartCapabilitiesCode_Value.
// If the string doesn't match exactly one of the supported values, SmartCapabilitiesCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) SmartCapabilitiesCode(s string) cpb.SmartCapabilitiesCode_Value {
	if c.SmartCapabilitiesCodeMap != nil {
		return c.SmartCapabilitiesCodeMap[strings.ToUpper(s)]
	}
	return DefaultSmartCapabilitiesCodeMap[strings.ToUpper(s)]
}

// SortDirectionCode converts the given string to a cpb.SortDirectionCode_Value.
// If the string doesn't match exactly one of the supported values, SortDirectionCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) SortDirectionCode(s string) cpb.SortDirectionCode_Value {
	if c.SortDirectionCodeMap != nil {
		return c.SortDirectionCodeMap[strings.ToUpper(s)]
	}
	return DefaultSortDirectionCodeMap[strings.ToUpper(s)]
}

// SpecimenContainedPreferenceCode converts the given string to a cpb.SpecimenContainedPreferenceCode_Value.
// If the string doesn't match exactly one of the supported values, SpecimenContainedPreferenceCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) SpecimenContainedPreferenceCode(s string) cpb.SpecimenContainedPreferenceCode_Value {
	if c.SpecimenContainedPreferenceCodeMap != nil {
		return c.SpecimenContainedPreferenceCodeMap[strings.ToUpper(s)]
	}
	return DefaultSpecimenContainedPreferenceCodeMap[strings.ToUpper(s)]
}

// SpecimenStatusCode converts the given string to a cpb.SpecimenStatusCode_Value.
// If the string doesn't match exactly one of the supported values, SpecimenStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) SpecimenStatusCode(s string) cpb.SpecimenStatusCode_Value {
	if c.SpecimenStatusCodeMap != nil {
		return c.SpecimenStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultSpecimenStatusCodeMap[strings.ToUpper(s)]
}

// StandardsStatusCode converts the given string to a cpb.StandardsStatusCode_Value.
// If the string doesn't match exactly one of the supported values, StandardsStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) StandardsStatusCode(s string) cpb.StandardsStatusCode_Value {
	if c.StandardsStatusCodeMap != nil {
		return c.StandardsStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultStandardsStatusCodeMap[strings.ToUpper(s)]
}

// StatusCode converts the given string to a cpb.StatusCode_Value.
// If the string doesn't match exactly one of the supported values, StatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) StatusCode(s string) cpb.StatusCode_Value {
	if c.StatusCodeMap != nil {
		return c.StatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultStatusCodeMap[strings.ToUpper(s)]
}

// StrandTypeCode converts the given string to a cpb.StrandTypeCode_Value.
// If the string doesn't match exactly one of the supported values, StrandTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) StrandTypeCode(s string) cpb.StrandTypeCode_Value {
	if c.StrandTypeCodeMap != nil {
		return c.StrandTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultStrandTypeCodeMap[strings.ToUpper(s)]
}

// StructureDefinitionKindCode converts the given string to a cpb.StructureDefinitionKindCode_Value.
// If the string doesn't match exactly one of the supported values, StructureDefinitionKindCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) StructureDefinitionKindCode(s string) cpb.StructureDefinitionKindCode_Value {
	if c.StructureDefinitionKindCodeMap != nil {
		return c.StructureDefinitionKindCodeMap[strings.ToUpper(s)]
	}
	return DefaultStructureDefinitionKindCodeMap[strings.ToUpper(s)]
}

// StructureMapContextTypeCode converts the given string to a cpb.StructureMapContextTypeCode_Value.
// If the string doesn't match exactly one of the supported values, StructureMapContextTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) StructureMapContextTypeCode(s string) cpb.StructureMapContextTypeCode_Value {
	if c.StructureMapContextTypeCodeMap != nil {
		return c.StructureMapContextTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultStructureMapContextTypeCodeMap[strings.ToUpper(s)]
}

// StructureMapGroupTypeModeCode converts the given string to a cpb.StructureMapGroupTypeModeCode_Value.
// If the string doesn't match exactly one of the supported values, StructureMapGroupTypeModeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) StructureMapGroupTypeModeCode(s string) cpb.StructureMapGroupTypeModeCode_Value {
	if c.StructureMapGroupTypeModeCodeMap != nil {
		return c.StructureMapGroupTypeModeCodeMap[strings.ToUpper(s)]
	}
	return DefaultStructureMapGroupTypeModeCodeMap[strings.ToUpper(s)]
}

// StructureMapInputModeCode converts the given string to a cpb.StructureMapInputModeCode_Value.
// If the string doesn't match exactly one of the supported values, StructureMapInputModeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) StructureMapInputModeCode(s string) cpb.StructureMapInputModeCode_Value {
	if c.StructureMapInputModeCodeMap != nil {
		return c.StructureMapInputModeCodeMap[strings.ToUpper(s)]
	}
	return DefaultStructureMapInputModeCodeMap[strings.ToUpper(s)]
}

// StructureMapModelModeCode converts the given string to a cpb.StructureMapModelModeCode_Value.
// If the string doesn't match exactly one of the supported values, StructureMapModelModeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) StructureMapModelModeCode(s string) cpb.StructureMapModelModeCode_Value {
	if c.StructureMapModelModeCodeMap != nil {
		return c.StructureMapModelModeCodeMap[strings.ToUpper(s)]
	}
	return DefaultStructureMapModelModeCodeMap[strings.ToUpper(s)]
}

// StructureMapSourceListModeCode converts the given string to a cpb.StructureMapSourceListModeCode_Value.
// If the string doesn't match exactly one of the supported values, StructureMapSourceListModeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) StructureMapSourceListModeCode(s string) cpb.StructureMapSourceListModeCode_Value {
	if c.StructureMapSourceListModeCodeMap != nil {
		return c.StructureMapSourceListModeCodeMap[strings.ToUpper(s)]
	}
	return DefaultStructureMapSourceListModeCodeMap[strings.ToUpper(s)]
}

// StructureMapTargetListModeCode converts the given string to a cpb.StructureMapTargetListModeCode_Value.
// If the string doesn't match exactly one of the supported values, StructureMapTargetListModeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) StructureMapTargetListModeCode(s string) cpb.StructureMapTargetListModeCode_Value {
	if c.StructureMapTargetListModeCodeMap != nil {
		return c.StructureMapTargetListModeCodeMap[strings.ToUpper(s)]
	}
	return DefaultStructureMapTargetListModeCodeMap[strings.ToUpper(s)]
}

// StructureMapTransformCode converts the given string to a cpb.StructureMapTransformCode_Value.
// If the string doesn't match exactly one of the supported values, StructureMapTransformCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) StructureMapTransformCode(s string) cpb.StructureMapTransformCode_Value {
	if c.StructureMapTransformCodeMap != nil {
		return c.StructureMapTransformCodeMap[strings.ToUpper(s)]
	}
	return DefaultStructureMapTransformCodeMap[strings.ToUpper(s)]
}

// SubscriptionChannelTypeCode converts the given string to a cpb.SubscriptionChannelTypeCode_Value.
// If the string doesn't match exactly one of the supported values, SubscriptionChannelTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) SubscriptionChannelTypeCode(s string) cpb.SubscriptionChannelTypeCode_Value {
	if c.SubscriptionChannelTypeCodeMap != nil {
		return c.SubscriptionChannelTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultSubscriptionChannelTypeCodeMap[strings.ToUpper(s)]
}

// SubscriptionStatusCode converts the given string to a cpb.SubscriptionStatusCode_Value.
// If the string doesn't match exactly one of the supported values, SubscriptionStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) SubscriptionStatusCode(s string) cpb.SubscriptionStatusCode_Value {
	if c.SubscriptionStatusCodeMap != nil {
		return c.SubscriptionStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultSubscriptionStatusCodeMap[strings.ToUpper(s)]
}

// SupplyDeliveryStatusCode converts the given string to a cpb.SupplyDeliveryStatusCode_Value.
// If the string doesn't match exactly one of the supported values, SupplyDeliveryStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) SupplyDeliveryStatusCode(s string) cpb.SupplyDeliveryStatusCode_Value {
	if c.SupplyDeliveryStatusCodeMap != nil {
		return c.SupplyDeliveryStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultSupplyDeliveryStatusCodeMap[strings.ToUpper(s)]
}

// SupplyItemTypeCode converts the given string to a cpb.SupplyItemTypeCode_Value.
// If the string doesn't match exactly one of the supported values, SupplyItemTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) SupplyItemTypeCode(s string) cpb.SupplyItemTypeCode_Value {
	if c.SupplyItemTypeCodeMap != nil {
		return c.SupplyItemTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultSupplyItemTypeCodeMap[strings.ToUpper(s)]
}

// SupplyRequestStatusCode converts the given string to a cpb.SupplyRequestStatusCode_Value.
// If the string doesn't match exactly one of the supported values, SupplyRequestStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) SupplyRequestStatusCode(s string) cpb.SupplyRequestStatusCode_Value {
	if c.SupplyRequestStatusCodeMap != nil {
		return c.SupplyRequestStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultSupplyRequestStatusCodeMap[strings.ToUpper(s)]
}

// TaskIntentCode converts the given string to a cpb.TaskIntentCode_Value.
// If the string doesn't match exactly one of the supported values, TaskIntentCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) TaskIntentCode(s string) cpb.TaskIntentCode_Value {
	if c.TaskIntentCodeMap != nil {
		return c.TaskIntentCodeMap[strings.ToUpper(s)]
	}
	return DefaultTaskIntentCodeMap[strings.ToUpper(s)]
}

// TaskStatusCode converts the given string to a cpb.TaskStatusCode_Value.
// If the string doesn't match exactly one of the supported values, TaskStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) TaskStatusCode(s string) cpb.TaskStatusCode_Value {
	if c.TaskStatusCodeMap != nil {
		return c.TaskStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultTaskStatusCodeMap[strings.ToUpper(s)]
}

// TemplateStatusCodeLifeCycleCode converts the given string to a cpb.TemplateStatusCodeLifeCycleCode_Value.
// If the string doesn't match exactly one of the supported values, TemplateStatusCodeLifeCycleCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) TemplateStatusCodeLifeCycleCode(s string) cpb.TemplateStatusCodeLifeCycleCode_Value {
	if c.TemplateStatusCodeLifeCycleCodeMap != nil {
		return c.TemplateStatusCodeLifeCycleCodeMap[strings.ToUpper(s)]
	}
	return DefaultTemplateStatusCodeLifeCycleCodeMap[strings.ToUpper(s)]
}

// TestReportActionResultCode converts the given string to a cpb.TestReportActionResultCode_Value.
// If the string doesn't match exactly one of the supported values, TestReportActionResultCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) TestReportActionResultCode(s string) cpb.TestReportActionResultCode_Value {
	if c.TestReportActionResultCodeMap != nil {
		return c.TestReportActionResultCodeMap[strings.ToUpper(s)]
	}
	return DefaultTestReportActionResultCodeMap[strings.ToUpper(s)]
}

// TestReportParticipantTypeCode converts the given string to a cpb.TestReportParticipantTypeCode_Value.
// If the string doesn't match exactly one of the supported values, TestReportParticipantTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) TestReportParticipantTypeCode(s string) cpb.TestReportParticipantTypeCode_Value {
	if c.TestReportParticipantTypeCodeMap != nil {
		return c.TestReportParticipantTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultTestReportParticipantTypeCodeMap[strings.ToUpper(s)]
}

// TestReportResultCode converts the given string to a cpb.TestReportResultCode_Value.
// If the string doesn't match exactly one of the supported values, TestReportResultCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) TestReportResultCode(s string) cpb.TestReportResultCode_Value {
	if c.TestReportResultCodeMap != nil {
		return c.TestReportResultCodeMap[strings.ToUpper(s)]
	}
	return DefaultTestReportResultCodeMap[strings.ToUpper(s)]
}

// TestReportStatusCode converts the given string to a cpb.TestReportStatusCode_Value.
// If the string doesn't match exactly one of the supported values, TestReportStatusCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) TestReportStatusCode(s string) cpb.TestReportStatusCode_Value {
	if c.TestReportStatusCodeMap != nil {
		return c.TestReportStatusCodeMap[strings.ToUpper(s)]
	}
	return DefaultTestReportStatusCodeMap[strings.ToUpper(s)]
}

// TestScriptRequestMethodCode converts the given string to a cpb.TestScriptRequestMethodCode_Value.
// If the string doesn't match exactly one of the supported values, TestScriptRequestMethodCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) TestScriptRequestMethodCode(s string) cpb.TestScriptRequestMethodCode_Value {
	if c.TestScriptRequestMethodCodeMap != nil {
		return c.TestScriptRequestMethodCodeMap[strings.ToUpper(s)]
	}
	return DefaultTestScriptRequestMethodCodeMap[strings.ToUpper(s)]
}

// TriggerTypeCode converts the given string to a cpb.TriggerTypeCode_Value.
// If the string doesn't match exactly one of the supported values, TriggerTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) TriggerTypeCode(s string) cpb.TriggerTypeCode_Value {
	if c.TriggerTypeCodeMap != nil {
		return c.TriggerTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultTriggerTypeCodeMap[strings.ToUpper(s)]
}

// TypeDerivationRuleCode converts the given string to a cpb.TypeDerivationRuleCode_Value.
// If the string doesn't match exactly one of the supported values, TypeDerivationRuleCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) TypeDerivationRuleCode(s string) cpb.TypeDerivationRuleCode_Value {
	if c.TypeDerivationRuleCodeMap != nil {
		return c.TypeDerivationRuleCodeMap[strings.ToUpper(s)]
	}
	return DefaultTypeDerivationRuleCodeMap[strings.ToUpper(s)]
}

// UDIEntryTypeCode converts the given string to a cpb.UDIEntryTypeCode_Value.
// If the string doesn't match exactly one of the supported values, UDIEntryTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) UDIEntryTypeCode(s string) cpb.UDIEntryTypeCode_Value {
	if c.UDIEntryTypeCodeMap != nil {
		return c.UDIEntryTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultUDIEntryTypeCodeMap[strings.ToUpper(s)]
}

// UseCode converts the given string to a cpb.UseCode_Value.
// If the string doesn't match exactly one of the supported values, UseCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) UseCode(s string) cpb.UseCode_Value {
	if c.UseCodeMap != nil {
		return c.UseCodeMap[strings.ToUpper(s)]
	}
	return DefaultUseCodeMap[strings.ToUpper(s)]
}

// V20444Code converts the given string to a cpb.V20444Code_Value.
// If the string doesn't match exactly one of the supported values, V20444Code returns
// INVALID_UNINITIALIZED.
func (c *Convertor) V20444Code(s string) cpb.V20444Code_Value {
	if c.V20444CodeMap != nil {
		return c.V20444CodeMap[strings.ToUpper(s)]
	}
	return DefaultV20444CodeMap[strings.ToUpper(s)]
}

// V3AddressUseCode converts the given string to a cpb.V3AddressUseCode_Value.
// If the string doesn't match exactly one of the supported values, V3AddressUseCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) V3AddressUseCode(s string) cpb.V3AddressUseCode_Value {
	if c.V3AddressUseCodeMap != nil {
		return c.V3AddressUseCodeMap[strings.ToUpper(s)]
	}
	return DefaultV3AddressUseCodeMap[strings.ToUpper(s)]
}

// V3ConfidentialityCode converts the given string to a cpb.V3ConfidentialityCode_Value.
// If the string doesn't match exactly one of the supported values, V3ConfidentialityCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) V3ConfidentialityCode(s string) cpb.V3ConfidentialityCode_Value {
	if c.V3ConfidentialityCodeMap != nil {
		return c.V3ConfidentialityCodeMap[strings.ToUpper(s)]
	}
	return DefaultV3ConfidentialityCodeMap[strings.ToUpper(s)]
}

// V3EntityNamePartQualifierCode converts the given string to a cpb.V3EntityNamePartQualifierCode_Value.
// If the string doesn't match exactly one of the supported values, V3EntityNamePartQualifierCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) V3EntityNamePartQualifierCode(s string) cpb.V3EntityNamePartQualifierCode_Value {
	if c.V3EntityNamePartQualifierCodeMap != nil {
		return c.V3EntityNamePartQualifierCodeMap[strings.ToUpper(s)]
	}
	return DefaultV3EntityNamePartQualifierCodeMap[strings.ToUpper(s)]
}

// V3EntityNamePartQualifierR2Code converts the given string to a cpb.V3EntityNamePartQualifierR2Code_Value.
// If the string doesn't match exactly one of the supported values, V3EntityNamePartQualifierR2Code returns
// INVALID_UNINITIALIZED.
func (c *Convertor) V3EntityNamePartQualifierR2Code(s string) cpb.V3EntityNamePartQualifierR2Code_Value {
	if c.V3EntityNamePartQualifierR2CodeMap != nil {
		return c.V3EntityNamePartQualifierR2CodeMap[strings.ToUpper(s)]
	}
	return DefaultV3EntityNamePartQualifierR2CodeMap[strings.ToUpper(s)]
}

// V3EntityNameUseCode converts the given string to a cpb.V3EntityNameUseCode_Value.
// If the string doesn't match exactly one of the supported values, V3EntityNameUseCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) V3EntityNameUseCode(s string) cpb.V3EntityNameUseCode_Value {
	if c.V3EntityNameUseCodeMap != nil {
		return c.V3EntityNameUseCodeMap[strings.ToUpper(s)]
	}
	return DefaultV3EntityNameUseCodeMap[strings.ToUpper(s)]
}

// V3EntityNameUseR2Code converts the given string to a cpb.V3EntityNameUseR2Code_Value.
// If the string doesn't match exactly one of the supported values, V3EntityNameUseR2Code returns
// INVALID_UNINITIALIZED.
func (c *Convertor) V3EntityNameUseR2Code(s string) cpb.V3EntityNameUseR2Code_Value {
	if c.V3EntityNameUseR2CodeMap != nil {
		return c.V3EntityNameUseR2CodeMap[strings.ToUpper(s)]
	}
	return DefaultV3EntityNameUseR2CodeMap[strings.ToUpper(s)]
}

// V3NullFlavorCode converts the given string to a cpb.V3NullFlavorCode_Value.
// If the string doesn't match exactly one of the supported values, V3NullFlavorCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) V3NullFlavorCode(s string) cpb.V3NullFlavorCode_Value {
	if c.V3NullFlavorCodeMap != nil {
		return c.V3NullFlavorCodeMap[strings.ToUpper(s)]
	}
	return DefaultV3NullFlavorCodeMap[strings.ToUpper(s)]
}

// V3ParticipationModeCode converts the given string to a cpb.V3ParticipationModeCode_Value.
// If the string doesn't match exactly one of the supported values, V3ParticipationModeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) V3ParticipationModeCode(s string) cpb.V3ParticipationModeCode_Value {
	if c.V3ParticipationModeCodeMap != nil {
		return c.V3ParticipationModeCodeMap[strings.ToUpper(s)]
	}
	return DefaultV3ParticipationModeCodeMap[strings.ToUpper(s)]
}

// V3ProbabilityDistributionTypeCode converts the given string to a cpb.V3ProbabilityDistributionTypeCode_Value.
// If the string doesn't match exactly one of the supported values, V3ProbabilityDistributionTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) V3ProbabilityDistributionTypeCode(s string) cpb.V3ProbabilityDistributionTypeCode_Value {
	if c.V3ProbabilityDistributionTypeCodeMap != nil {
		return c.V3ProbabilityDistributionTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultV3ProbabilityDistributionTypeCodeMap[strings.ToUpper(s)]
}

// V3RoleCode converts the given string to a cpb.V3RoleCode_Value.
// If the string doesn't match exactly one of the supported values, V3RoleCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) V3RoleCode(s string) cpb.V3RoleCode_Value {
	if c.V3RoleCodeMap != nil {
		return c.V3RoleCodeMap[strings.ToUpper(s)]
	}
	return DefaultV3RoleCodeMap[strings.ToUpper(s)]
}

// V3TimingEventCode converts the given string to a cpb.V3TimingEventCode_Value.
// If the string doesn't match exactly one of the supported values, V3TimingEventCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) V3TimingEventCode(s string) cpb.V3TimingEventCode_Value {
	if c.V3TimingEventCodeMap != nil {
		return c.V3TimingEventCodeMap[strings.ToUpper(s)]
	}
	return DefaultV3TimingEventCodeMap[strings.ToUpper(s)]
}

// VisionBaseCode converts the given string to a cpb.VisionBaseCode_Value.
// If the string doesn't match exactly one of the supported values, VisionBaseCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) VisionBaseCode(s string) cpb.VisionBaseCode_Value {
	if c.VisionBaseCodeMap != nil {
		return c.VisionBaseCodeMap[strings.ToUpper(s)]
	}
	return DefaultVisionBaseCodeMap[strings.ToUpper(s)]
}

// VisionEyesCode converts the given string to a cpb.VisionEyesCode_Value.
// If the string doesn't match exactly one of the supported values, VisionEyesCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) VisionEyesCode(s string) cpb.VisionEyesCode_Value {
	if c.VisionEyesCodeMap != nil {
		return c.VisionEyesCodeMap[strings.ToUpper(s)]
	}
	return DefaultVisionEyesCodeMap[strings.ToUpper(s)]
}

// XPathUsageTypeCode converts the given string to a cpb.XPathUsageTypeCode_Value.
// If the string doesn't match exactly one of the supported values, XPathUsageTypeCode returns
// INVALID_UNINITIALIZED.
func (c *Convertor) XPathUsageTypeCode(s string) cpb.XPathUsageTypeCode_Value {
	if c.XPathUsageTypeCodeMap != nil {
		return c.XPathUsageTypeCodeMap[strings.ToUpper(s)]
	}
	return DefaultXPathUsageTypeCodeMap[strings.ToUpper(s)]
}

// Convertor is a convertor of HL7v2 values to FHIR.
// The contained mappings can be overridden to support custom configurations.
// If there's no override, Convertor will use the default mapping.
type Convertor struct {
	AbstractTypeCodeMap                            map[string]cpb.AbstractTypeCode_Value
	AccountStatusCodeMap                           map[string]cpb.AccountStatusCode_Value
	ActionCardinalityBehaviorCodeMap               map[string]cpb.ActionCardinalityBehaviorCode_Value
	ActionConditionKindCodeMap                     map[string]cpb.ActionConditionKindCode_Value
	ActionGroupingBehaviorCodeMap                  map[string]cpb.ActionGroupingBehaviorCode_Value
	ActionParticipantTypeCodeMap                   map[string]cpb.ActionParticipantTypeCode_Value
	ActionPrecheckBehaviorCodeMap                  map[string]cpb.ActionPrecheckBehaviorCode_Value
	ActionRelationshipTypeCodeMap                  map[string]cpb.ActionRelationshipTypeCode_Value
	ActionRequiredBehaviorCodeMap                  map[string]cpb.ActionRequiredBehaviorCode_Value
	ActionSelectionBehaviorCodeMap                 map[string]cpb.ActionSelectionBehaviorCode_Value
	AddressTypeCodeMap                             map[string]cpb.AddressTypeCode_Value
	AddressUseCodeMap                              map[string]cpb.AddressUseCode_Value
	AdministrativeGenderCodeMap                    map[string]cpb.AdministrativeGenderCode_Value
	AdverseEventActualityCodeMap                   map[string]cpb.AdverseEventActualityCode_Value
	AdverseEventOutcomeCodeMap                     map[string]cpb.AdverseEventOutcomeCode_Value
	AdverseEventSeverityCodeMap                    map[string]cpb.AdverseEventSeverityCode_Value
	AggregationModeCodeMap                         map[string]cpb.AggregationModeCode_Value
	AllergyIntoleranceCategoryCodeMap              map[string]cpb.AllergyIntoleranceCategoryCode_Value
	AllergyIntoleranceClinicalStatusCodeMap        map[string]cpb.AllergyIntoleranceClinicalStatusCode_Value
	AllergyIntoleranceCriticalityCodeMap           map[string]cpb.AllergyIntoleranceCriticalityCode_Value
	AllergyIntoleranceSeverityCodeMap              map[string]cpb.AllergyIntoleranceSeverityCode_Value
	AllergyIntoleranceSubstanceExposureRiskCodeMap map[string]cpb.AllergyIntoleranceSubstanceExposureRiskCode_Value
	AllergyIntoleranceTypeCodeMap                  map[string]cpb.AllergyIntoleranceTypeCode_Value
	AllergyIntoleranceVerificationStatusCodeMap    map[string]cpb.AllergyIntoleranceVerificationStatusCode_Value
	AppointmentStatusCodeMap                       map[string]cpb.AppointmentStatusCode_Value
	AssertionDirectionTypeCodeMap                  map[string]cpb.AssertionDirectionTypeCode_Value
	AssertionOperatorTypeCodeMap                   map[string]cpb.AssertionOperatorTypeCode_Value
	AssertionResponseTypesCodeMap                  map[string]cpb.AssertionResponseTypesCode_Value
	AuditEventActionCodeMap                        map[string]cpb.AuditEventActionCode_Value
	AuditEventAgentNetworkTypeCodeMap              map[string]cpb.AuditEventAgentNetworkTypeCode_Value
	AuditEventOutcomeCodeMap                       map[string]cpb.AuditEventOutcomeCode_Value
	BenefitCostApplicabilityCodeMap                map[string]cpb.BenefitCostApplicabilityCode_Value
	BindingStrengthCodeMap                         map[string]cpb.BindingStrengthCode_Value
	BiologicallyDerivedProductCategoryCodeMap      map[string]cpb.BiologicallyDerivedProductCategoryCode_Value
	BiologicallyDerivedProductStatusCodeMap        map[string]cpb.BiologicallyDerivedProductStatusCode_Value
	BiologicallyDerivedProductStorageScaleCodeMap  map[string]cpb.BiologicallyDerivedProductStorageScaleCode_Value
	BundleTypeCodeMap                              map[string]cpb.BundleTypeCode_Value
	CanonicalStatusCodesForFHIRResourcesCodeMap    map[string]cpb.CanonicalStatusCodesForFHIRResourcesCode_Value
	CapabilityStatementKindCodeMap                 map[string]cpb.CapabilityStatementKindCode_Value
	CarePlanActivityStatusCodeMap                  map[string]cpb.CarePlanActivityStatusCode_Value
	CareTeamStatusCodeMap                          map[string]cpb.CareTeamStatusCode_Value
	CatalogEntryRelationTypeCodeMap                map[string]cpb.CatalogEntryRelationTypeCode_Value
	ChargeItemStatusCodeMap                        map[string]cpb.ChargeItemStatusCode_Value
	ChoiceListOrientationCodeMap                   map[string]cpb.ChoiceListOrientationCode_Value
	ClaimProcessingCodeMap                         map[string]cpb.ClaimProcessingCode_Value
	CodeSearchSupportCodeMap                       map[string]cpb.CodeSearchSupportCode_Value
	CodeSystemContentModeCodeMap                   map[string]cpb.CodeSystemContentModeCode_Value
	CodeSystemHierarchyMeaningCodeMap              map[string]cpb.CodeSystemHierarchyMeaningCode_Value
	CompartmentTypeCodeMap                         map[string]cpb.CompartmentTypeCode_Value
	CompositionAttestationModeCodeMap              map[string]cpb.CompositionAttestationModeCode_Value
	CompositionStatusCodeMap                       map[string]cpb.CompositionStatusCode_Value
	ConceptMapEquivalenceCodeMap                   map[string]cpb.ConceptMapEquivalenceCode_Value
	ConceptMapGroupUnmappedModeCodeMap             map[string]cpb.ConceptMapGroupUnmappedModeCode_Value
	ConditionClinicalStatusCodeMap                 map[string]cpb.ConditionClinicalStatusCode_Value
	ConditionVerificationStatusCodeMap             map[string]cpb.ConditionVerificationStatusCode_Value
	ConditionalDeleteStatusCodeMap                 map[string]cpb.ConditionalDeleteStatusCode_Value
	ConditionalReadStatusCodeMap                   map[string]cpb.ConditionalReadStatusCode_Value
	ConformanceExpectationCodeMap                  map[string]cpb.ConformanceExpectationCode_Value
	ConsentDataMeaningCodeMap                      map[string]cpb.ConsentDataMeaningCode_Value
	ConsentProvisionTypeCodeMap                    map[string]cpb.ConsentProvisionTypeCode_Value
	ConsentStateCodeMap                            map[string]cpb.ConsentStateCode_Value
	ConstraintSeverityCodeMap                      map[string]cpb.ConstraintSeverityCode_Value
	ContactPointSystemCodeMap                      map[string]cpb.ContactPointSystemCode_Value
	ContactPointUseCodeMap                         map[string]cpb.ContactPointUseCode_Value
	ContractResourcePublicationStatusCodeMap       map[string]cpb.ContractResourcePublicationStatusCode_Value
	ContractResourceStatusCodeMap                  map[string]cpb.ContractResourceStatusCode_Value
	ContributorTypeCodeMap                         map[string]cpb.ContributorTypeCode_Value
	DataAbsentReasonCodeMap                        map[string]cpb.DataAbsentReasonCode_Value
	DataTypeCodeMap                                map[string]cpb.DataTypeCode_Value
	DaysOfWeekCodeMap                              map[string]cpb.DaysOfWeekCode_Value
	DetectedIssueSeverityCodeMap                   map[string]cpb.DetectedIssueSeverityCode_Value
	DeviceMetricCalibrationStateCodeMap            map[string]cpb.DeviceMetricCalibrationStateCode_Value
	DeviceMetricCalibrationTypeCodeMap             map[string]cpb.DeviceMetricCalibrationTypeCode_Value
	DeviceMetricCategoryCodeMap                    map[string]cpb.DeviceMetricCategoryCode_Value
	DeviceMetricColorCodeMap                       map[string]cpb.DeviceMetricColorCode_Value
	DeviceMetricOperationalStatusCodeMap           map[string]cpb.DeviceMetricOperationalStatusCode_Value
	DeviceNameTypeCodeMap                          map[string]cpb.DeviceNameTypeCode_Value
	DeviceUseStatementStatusCodeMap                map[string]cpb.DeviceUseStatementStatusCode_Value
	DiagnosticReportStatusCodeMap                  map[string]cpb.DiagnosticReportStatusCode_Value
	DiscriminatorTypeCodeMap                       map[string]cpb.DiscriminatorTypeCode_Value
	DocumentModeCodeMap                            map[string]cpb.DocumentModeCode_Value
	DocumentReferenceStatusCodeMap                 map[string]cpb.DocumentReferenceStatusCode_Value
	DocumentRelationshipTypeCodeMap                map[string]cpb.DocumentRelationshipTypeCode_Value
	EligibilityRequestPurposeCodeMap               map[string]cpb.EligibilityRequestPurposeCode_Value
	EligibilityResponsePurposeCodeMap              map[string]cpb.EligibilityResponsePurposeCode_Value
	EnableWhenBehaviorCodeMap                      map[string]cpb.EnableWhenBehaviorCode_Value
	EncounterLocationStatusCodeMap                 map[string]cpb.EncounterLocationStatusCode_Value
	EncounterStatusCodeMap                         map[string]cpb.EncounterStatusCode_Value
	EndpointStatusCodeMap                          map[string]cpb.EndpointStatusCode_Value
	EpisodeOfCareStatusCodeMap                     map[string]cpb.EpisodeOfCareStatusCode_Value
	EventCapabilityModeCodeMap                     map[string]cpb.EventCapabilityModeCode_Value
	EventStatusCodeMap                             map[string]cpb.EventStatusCode_Value
	EventTimingCodeMap                             map[string]cpb.EventTimingCode_Value
	EvidenceVariableTypeCodeMap                    map[string]cpb.EvidenceVariableTypeCode_Value
	ExampleScenarioActorTypeCodeMap                map[string]cpb.ExampleScenarioActorTypeCode_Value
	ExpansionParameterSourceCodeMap                map[string]cpb.ExpansionParameterSourceCode_Value
	ExpansionProcessingRuleCodeMap                 map[string]cpb.ExpansionProcessingRuleCode_Value
	ExplanationOfBenefitStatusCodeMap              map[string]cpb.ExplanationOfBenefitStatusCode_Value
	ExposureStateCodeMap                           map[string]cpb.ExposureStateCode_Value
	ExtensionContextTypeCodeMap                    map[string]cpb.ExtensionContextTypeCode_Value
	FHIRDeviceStatusCodeMap                        map[string]cpb.FHIRDeviceStatusCode_Value
	FHIRRestfulInteractionsCodeMap                 map[string]cpb.FHIRRestfulInteractionsCode_Value
	FHIRSubstanceStatusCodeMap                     map[string]cpb.FHIRSubstanceStatusCode_Value
	FHIRVersionCodeMap                             map[string]cpb.FHIRVersionCode_Value
	FamilyHistoryStatusCodeMap                     map[string]cpb.FamilyHistoryStatusCode_Value
	FilterOperatorCodeMap                          map[string]cpb.FilterOperatorCode_Value
	FinancialResourceStatusCodeMap                 map[string]cpb.FinancialResourceStatusCode_Value
	FlagStatusCodeMap                              map[string]cpb.FlagStatusCode_Value
	GoalAcceptanceStatusCodeMap                    map[string]cpb.GoalAcceptanceStatusCode_Value
	GoalLifecycleStatusCodeMap                     map[string]cpb.GoalLifecycleStatusCode_Value
	GraphCompartmentRuleCodeMap                    map[string]cpb.GraphCompartmentRuleCode_Value
	GraphCompartmentUseCodeMap                     map[string]cpb.GraphCompartmentUseCode_Value
	GroupMeasureCodeMap                            map[string]cpb.GroupMeasureCode_Value
	GroupTypeCodeMap                               map[string]cpb.GroupTypeCode_Value
	GuidanceResponseStatusCodeMap                  map[string]cpb.GuidanceResponseStatusCode_Value
	GuidePageGenerationCodeMap                     map[string]cpb.GuidePageGenerationCode_Value
	GuideParameterCodeMap                          map[string]cpb.GuideParameterCode_Value
	HL7WorkgroupCodeMap                            map[string]cpb.HL7WorkgroupCode_Value
	HTTPVerbCodeMap                                map[string]cpb.HTTPVerbCode_Value
	HumanNameAssemblyOrderCodeMap                  map[string]cpb.HumanNameAssemblyOrderCode_Value
	IdentifierUseCodeMap                           map[string]cpb.IdentifierUseCode_Value
	IdentityAssuranceLevelCodeMap                  map[string]cpb.IdentityAssuranceLevelCode_Value
	ImagingStudyStatusCodeMap                      map[string]cpb.ImagingStudyStatusCode_Value
	ImplantStatusCodeMap                           map[string]cpb.ImplantStatusCode_Value
	InvoicePriceComponentTypeCodeMap               map[string]cpb.InvoicePriceComponentTypeCode_Value
	InvoiceStatusCodeMap                           map[string]cpb.InvoiceStatusCode_Value
	IssueSeverityCodeMap                           map[string]cpb.IssueSeverityCode_Value
	IssueTypeCodeMap                               map[string]cpb.IssueTypeCode_Value
	LinkTypeCodeMap                                map[string]cpb.LinkTypeCode_Value
	LinkageTypeCodeMap                             map[string]cpb.LinkageTypeCode_Value
	ListModeCodeMap                                map[string]cpb.ListModeCode_Value
	ListStatusCodeMap                              map[string]cpb.ListStatusCode_Value
	LocationModeCodeMap                            map[string]cpb.LocationModeCode_Value
	LocationStatusCodeMap                          map[string]cpb.LocationStatusCode_Value
	MatchGradeCodeMap                              map[string]cpb.MatchGradeCode_Value
	MeasureImprovementNotationCodeMap              map[string]cpb.MeasureImprovementNotationCode_Value
	MeasureReportStatusCodeMap                     map[string]cpb.MeasureReportStatusCode_Value
	MeasureReportTypeCodeMap                       map[string]cpb.MeasureReportTypeCode_Value
	MedicationAdministrationStatusCodeMap          map[string]cpb.MedicationAdministrationStatusCode_Value
	MedicationDispenseStatusCodeMap                map[string]cpb.MedicationDispenseStatusCode_Value
	MedicationKnowledgeStatusCodeMap               map[string]cpb.MedicationKnowledgeStatusCode_Value
	MedicationRequestIntentCodeMap                 map[string]cpb.MedicationRequestIntentCode_Value
	MedicationStatementStatusCodesMap              map[string]cpb.MedicationStatementStatusCodes_Value
	MedicationStatusCodeMap                        map[string]cpb.MedicationStatusCode_Value
	MedicationrequestStatusCodeMap                 map[string]cpb.MedicationrequestStatusCode_Value
	MessageSignificanceCategoryCodeMap             map[string]cpb.MessageSignificanceCategoryCode_Value
	MessageheaderResponseRequestCodeMap            map[string]cpb.MessageheaderResponseRequestCode_Value
	NameUseCodeMap                                 map[string]cpb.NameUseCode_Value
	NamingSystemIdentifierTypeCodeMap              map[string]cpb.NamingSystemIdentifierTypeCode_Value
	NamingSystemTypeCodeMap                        map[string]cpb.NamingSystemTypeCode_Value
	NarrativeStatusCodeMap                         map[string]cpb.NarrativeStatusCode_Value
	NoteTypeCodeMap                                map[string]cpb.NoteTypeCode_Value
	ObservationDataTypeCodeMap                     map[string]cpb.ObservationDataTypeCode_Value
	ObservationRangeCategoryCodeMap                map[string]cpb.ObservationRangeCategoryCode_Value
	ObservationStatusCodeMap                       map[string]cpb.ObservationStatusCode_Value
	OperationKindCodeMap                           map[string]cpb.OperationKindCode_Value
	OperationParameterUseCodeMap                   map[string]cpb.OperationParameterUseCode_Value
	OrientationTypeCodeMap                         map[string]cpb.OrientationTypeCode_Value
	ParticipantRequiredCodeMap                     map[string]cpb.ParticipantRequiredCode_Value
	ParticipationStatusCodeMap                     map[string]cpb.ParticipationStatusCode_Value
	PropertyRepresentationCodeMap                  map[string]cpb.PropertyRepresentationCode_Value
	PropertyTypeCodeMap                            map[string]cpb.PropertyTypeCode_Value
	ProvenanceEntityRoleCodeMap                    map[string]cpb.ProvenanceEntityRoleCode_Value
	PublicationStatusCodeMap                       map[string]cpb.PublicationStatusCode_Value
	QualityTypeCodeMap                             map[string]cpb.QualityTypeCode_Value
	QuantityComparatorCodeMap                      map[string]cpb.QuantityComparatorCode_Value
	QuestionnaireItemOperatorCodeMap               map[string]cpb.QuestionnaireItemOperatorCode_Value
	QuestionnaireItemTypeCodeMap                   map[string]cpb.QuestionnaireItemTypeCode_Value
	QuestionnaireItemUsageModeCodeMap              map[string]cpb.QuestionnaireItemUsageModeCode_Value
	QuestionnaireResponseStatusCodeMap             map[string]cpb.QuestionnaireResponseStatusCode_Value
	ReferenceHandlingPolicyCodeMap                 map[string]cpb.ReferenceHandlingPolicyCode_Value
	ReferenceVersionRulesCodeMap                   map[string]cpb.ReferenceVersionRulesCode_Value
	RelatedArtifactTypeCodeMap                     map[string]cpb.RelatedArtifactTypeCode_Value
	RepositoryTypeCodeMap                          map[string]cpb.RepositoryTypeCode_Value
	RequestIntentCodeMap                           map[string]cpb.RequestIntentCode_Value
	RequestPriorityCodeMap                         map[string]cpb.RequestPriorityCode_Value
	RequestResourceTypeCodeMap                     map[string]cpb.RequestResourceTypeCode_Value
	RequestStatusCodeMap                           map[string]cpb.RequestStatusCode_Value
	ResearchElementTypeCodeMap                     map[string]cpb.ResearchElementTypeCode_Value
	ResearchStudyStatusCodeMap                     map[string]cpb.ResearchStudyStatusCode_Value
	ResearchSubjectStatusCodeMap                   map[string]cpb.ResearchSubjectStatusCode_Value
	ResourceSecurityCategoryCodeMap                map[string]cpb.ResourceSecurityCategoryCode_Value
	ResourceTypeCodeMap                            map[string]cpb.ResourceTypeCode_Value
	ResourceVersionPolicyCodeMap                   map[string]cpb.ResourceVersionPolicyCode_Value
	ResponseTypeCodeMap                            map[string]cpb.ResponseTypeCode_Value
	RestfulCapabilityModeCodeMap                   map[string]cpb.RestfulCapabilityModeCode_Value
	SPDXLicenseCodeMap                             map[string]cpb.SPDXLicenseCode_Value
	SearchComparatorCodeMap                        map[string]cpb.SearchComparatorCode_Value
	SearchEntryModeCodeMap                         map[string]cpb.SearchEntryModeCode_Value
	SearchModifierCodeMap                          map[string]cpb.SearchModifierCode_Value
	SearchParamTypeCodeMap                         map[string]cpb.SearchParamTypeCode_Value
	SequenceTypeCodeMap                            map[string]cpb.SequenceTypeCode_Value
	SlicingRulesCodeMap                            map[string]cpb.SlicingRulesCode_Value
	SlotStatusCodeMap                              map[string]cpb.SlotStatusCode_Value
	SmartCapabilitiesCodeMap                       map[string]cpb.SmartCapabilitiesCode_Value
	SortDirectionCodeMap                           map[string]cpb.SortDirectionCode_Value
	SpecimenContainedPreferenceCodeMap             map[string]cpb.SpecimenContainedPreferenceCode_Value
	SpecimenStatusCodeMap                          map[string]cpb.SpecimenStatusCode_Value
	StandardsStatusCodeMap                         map[string]cpb.StandardsStatusCode_Value
	StatusCodeMap                                  map[string]cpb.StatusCode_Value
	StrandTypeCodeMap                              map[string]cpb.StrandTypeCode_Value
	StructureDefinitionKindCodeMap                 map[string]cpb.StructureDefinitionKindCode_Value
	StructureMapContextTypeCodeMap                 map[string]cpb.StructureMapContextTypeCode_Value
	StructureMapGroupTypeModeCodeMap               map[string]cpb.StructureMapGroupTypeModeCode_Value
	StructureMapInputModeCodeMap                   map[string]cpb.StructureMapInputModeCode_Value
	StructureMapModelModeCodeMap                   map[string]cpb.StructureMapModelModeCode_Value
	StructureMapSourceListModeCodeMap              map[string]cpb.StructureMapSourceListModeCode_Value
	StructureMapTargetListModeCodeMap              map[string]cpb.StructureMapTargetListModeCode_Value
	StructureMapTransformCodeMap                   map[string]cpb.StructureMapTransformCode_Value
	SubscriptionChannelTypeCodeMap                 map[string]cpb.SubscriptionChannelTypeCode_Value
	SubscriptionStatusCodeMap                      map[string]cpb.SubscriptionStatusCode_Value
	SupplyDeliveryStatusCodeMap                    map[string]cpb.SupplyDeliveryStatusCode_Value
	SupplyItemTypeCodeMap                          map[string]cpb.SupplyItemTypeCode_Value
	SupplyRequestStatusCodeMap                     map[string]cpb.SupplyRequestStatusCode_Value
	TaskIntentCodeMap                              map[string]cpb.TaskIntentCode_Value
	TaskStatusCodeMap                              map[string]cpb.TaskStatusCode_Value
	TemplateStatusCodeLifeCycleCodeMap             map[string]cpb.TemplateStatusCodeLifeCycleCode_Value
	TestReportActionResultCodeMap                  map[string]cpb.TestReportActionResultCode_Value
	TestReportParticipantTypeCodeMap               map[string]cpb.TestReportParticipantTypeCode_Value
	TestReportResultCodeMap                        map[string]cpb.TestReportResultCode_Value
	TestReportStatusCodeMap                        map[string]cpb.TestReportStatusCode_Value
	TestScriptRequestMethodCodeMap                 map[string]cpb.TestScriptRequestMethodCode_Value
	TriggerTypeCodeMap                             map[string]cpb.TriggerTypeCode_Value
	TypeDerivationRuleCodeMap                      map[string]cpb.TypeDerivationRuleCode_Value
	UDIEntryTypeCodeMap                            map[string]cpb.UDIEntryTypeCode_Value
	UseCodeMap                                     map[string]cpb.UseCode_Value
	V20444CodeMap                                  map[string]cpb.V20444Code_Value
	V3AddressUseCodeMap                            map[string]cpb.V3AddressUseCode_Value
	V3ConfidentialityCodeMap                       map[string]cpb.V3ConfidentialityCode_Value
	V3EntityNamePartQualifierCodeMap               map[string]cpb.V3EntityNamePartQualifierCode_Value
	V3EntityNamePartQualifierR2CodeMap             map[string]cpb.V3EntityNamePartQualifierR2Code_Value
	V3EntityNameUseCodeMap                         map[string]cpb.V3EntityNameUseCode_Value
	V3EntityNameUseR2CodeMap                       map[string]cpb.V3EntityNameUseR2Code_Value
	V3NullFlavorCodeMap                            map[string]cpb.V3NullFlavorCode_Value
	V3ParticipationModeCodeMap                     map[string]cpb.V3ParticipationModeCode_Value
	V3ProbabilityDistributionTypeCodeMap           map[string]cpb.V3ProbabilityDistributionTypeCode_Value
	V3RoleCodeMap                                  map[string]cpb.V3RoleCode_Value
	V3TimingEventCodeMap                           map[string]cpb.V3TimingEventCode_Value
	VisionBaseCodeMap                              map[string]cpb.VisionBaseCode_Value
	VisionEyesCodeMap                              map[string]cpb.VisionEyesCode_Value
	XPathUsageTypeCodeMap                          map[string]cpb.XPathUsageTypeCode_Value
}
