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
	"context"
	"fmt"
	"testing"

	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/constants"
	"github.com/google/simhospital/pkg/test"
)

func TestToHL7(t *testing.T) {
	ctx := context.Background()
	c, err := config.LoadHL7Config(ctx, test.MessageConfigProd)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigProd, err)
	}
	convertor := NewAbnormalFlagConvertor(c)

	cases := []struct {
		input constants.AbnormalFlag
		want  string
	}{
		{
			input: constants.AbnormalFlagLow,
			want:  c.AbnormalFlags.BelowLowNormal,
		},
		{
			input: constants.AbnormalFlagHigh,
			want:  c.AbnormalFlags.AboveHighNormal,
		},
		{
			input: constants.AbnormalFlagEmpty,
			want:  "",
		},
		{
			input: constants.AbnormalFlagNormal,
			want:  "",
		},
		{
			input: constants.AbnormalFlagDefault,
			want:  "",
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			if got := convertor.ToHL7(tc.input); got != tc.want {
				t.Errorf("ToHL7(%v) = %s, want %s", tc.input, got, tc.want)
			}
		})
	}
}
