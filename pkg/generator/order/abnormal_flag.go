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

package order

import (
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/constants"
)

// AbnormalFlagConvertor converts abnormal flag values.
type AbnormalFlagConvertor struct {
	mapping map[constants.AbnormalFlag]string
}

// NewAbnormalFlagConvertor returns a new abnormalFlagConvertor based on
// the HL7 abnormal flag values provided in the HL7Config.
func NewAbnormalFlagConvertor(c *config.HL7Config) AbnormalFlagConvertor {
	return AbnormalFlagConvertor{
		mapping: map[constants.AbnormalFlag]string{
			constants.AbnormalFlagLow:  c.AbnormalFlags.BelowLowNormal,
			constants.AbnormalFlagHigh: c.AbnormalFlags.AboveHighNormal,
		},
	}
}

// ToHL7 returns an HL7 value of the abnormal flag value provided.
func (c *AbnormalFlagConvertor) ToHL7(f constants.AbnormalFlag) string {
	return c.mapping[f]
}
