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

import (
	"fmt"
	"testing"
)

func TestFromRandomType(t *testing.T) {
	cases := []struct {
		randomType string
		want       AbnormalFlag
	}{
		{randomType: AbnormalLow, want: AbnormalFlagLow},
		{randomType: AbnormalHigh, want: AbnormalFlagHigh},
		{randomType: EmptyString, want: AbnormalFlagEmpty},
		{randomType: "any", want: AbnormalFlagEmpty},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s-%s", tc.randomType, tc.want), func(t *testing.T) {
			got := FromRandomType(tc.randomType)
			if got != tc.want {
				t.Errorf("FromRandomType(%s) = %v, want %v", tc.randomType, got, tc.want)
			}
		})
	}
}
