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

package person

import (
	"context"
	"math"
	"testing"

	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/test"
	"github.com/google/simhospital/pkg/test/testwrite"
)

func TestEthnicityGenerator(t *testing.T) {
	ctx := context.Background()
	hl7Config, err := config.LoadHL7Config(ctx, test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	ethnicity := `
White,White,1
Asian,Asian,2
Black,Black,2
Other,Other,2
nil,nil,3
`
	fEthnicity := testwrite.BytesToFile(t, []byte(ethnicity))
	dataF := test.DataFiles[test.Test]
	dataF.Ethnicities = fEthnicity
	dataConfig, err := config.LoadData(ctx, dataF, hl7Config)
	if err != nil {
		t.Fatalf("LoadData(%+v, %+v) failed with %v", dataF, hl7Config, err)
	}

	want := map[string]int{
		"White": 1,
		"Asian": 2,
		"Black": 2,
		"Other": 2,
		"nil":   3,
	}
	all := 0
	for _, v := range want {
		all += v
	}

	gotPerKey := map[string]int{}

	eg := NewEthnicityGenerator(dataConfig)

	runs := 1000
	for i := 0; i < runs; i++ {
		got := eg.Random()
		if got == nil {
			gotPerKey["nil"]++
		} else if _, ok := want[got.ID]; ok {
			gotPerKey[got.ID]++
		} else {
			t.Errorf("eg.Random()=%v, want nil or one of: %v", got, want)
		}
	}

	delta := float64(runs) / 5.0
	for k, v := range want {
		wantFreq := float64(runs) * (float64(v) / float64(all))
		if gotFreq := gotPerKey[k]; math.Abs(float64(gotFreq)-wantFreq) > delta {
			t.Errorf("eg.Random() got %q freq=%d, want within %v of %f", k, gotFreq, delta, wantFreq)
		}
	}
}
