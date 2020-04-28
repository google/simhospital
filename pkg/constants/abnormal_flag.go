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

// AbnormalFlag is the abnormal flag to be set on an observation.
// Full set of values:
// http://hl7-definition.caristix.com:9010/HL7%20v2.2/table/Default.aspx?version=HL7+v2.2&table=0078
type AbnormalFlag string

const (
	// AbnormalFlagLow is an abnormal flag indicating the value is below low normal.
	AbnormalFlagLow AbnormalFlag = "LOW"
	// AbnormalFlagHigh is an abnormal flag indicating the value is above high normal.
	AbnormalFlagHigh AbnormalFlag = "HIGH"
	// AbnormalFlagEmpty represents a normal abnormal flag. Equivalent to AbnormalFlagNormal.
	AbnormalFlagEmpty AbnormalFlag = ""
	// AbnormalFlagNormal represents a normal abnormal flag. Equivalent to AbnormalFlagEmpty.
	AbnormalFlagNormal AbnormalFlag = "NORMAL"
	// AbnormalFlagDefault indicates that the abnormal flag should be derived
	// from the reference ranges and the value.
	AbnormalFlagDefault AbnormalFlag = "DEFAULT"
)

// AbnormalFlagValues is a map of valid abnormal flag values.
var AbnormalFlagValues = map[AbnormalFlag]bool{
	AbnormalFlagLow:     true,
	AbnormalFlagHigh:    true,
	AbnormalFlagEmpty:   true,
	AbnormalFlagNormal:  true,
	AbnormalFlagDefault: true,
}

// IsNormalFlag returns whether s is a normal abnormal flag.
func IsNormalFlag(s AbnormalFlag) bool {
	return s == AbnormalFlagEmpty || s == AbnormalFlagNormal
}

// FromRandomType returns an abnormal flag value for the given randomType.
func FromRandomType(randomType string) AbnormalFlag {
	switch randomType {
	case AbnormalLow:
		return AbnormalFlagLow
	case AbnormalHigh:
		return AbnormalFlagHigh
	default:
		return AbnormalFlagEmpty
	}
}
