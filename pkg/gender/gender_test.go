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

package gender

import (
	"context"
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/logging"
	"github.com/google/simhospital/pkg/pathway"
	"github.com/google/simhospital/pkg/test"
)

func TestMain(m *testing.M) {
	logging.SetLogLevel(logrus.DebugLevel)

	retCode := m.Run()

	os.Exit(retCode)
}

func TestConvertorPathwayToHL7(t *testing.T) {
	ctx := context.Background()
	hl7Config, err := config.LoadHL7Config(ctx, test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}

	wantMapping := map[pathway.Gender]string{
		pathway.Male:   hl7Config.Gender.Male,
		pathway.Female: hl7Config.Gender.Female,
		"NOT A GENDER": hl7Config.Gender.Unknown,
	}

	c := NewConvertor(hl7Config)
	for k, v := range wantMapping {
		t.Run(fmt.Sprintf("%v-%v", k, v), func(t *testing.T) {
			if got, want := c.PathwayToHL7(k), v; got != want {
				t.Errorf("PathwayToHL7(%v)=%v, want %v", k, got, want)
			}
		})
	}
}

func TestConvertorInternalToHL7(t *testing.T) {
	ctx := context.Background()
	hl7Config, err := config.LoadHL7Config(ctx, test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}

	wantMapping := map[Internal]string{
		Male:    hl7Config.Gender.Male,
		Female:  hl7Config.Gender.Female,
		Unknown: hl7Config.Gender.Unknown,
	}
	c := NewConvertor(hl7Config)

	for k, v := range wantMapping {
		t.Run(fmt.Sprintf("%v-%v", k, v), func(t *testing.T) {
			if got, want := c.InternalToHL7(k), v; got != want {
				t.Errorf("c.InternalToHL7(%v)=%v, want %v", k, got, want)
			}
		})
	}
}

func TestConvertorHL7ToInternal(t *testing.T) {
	ctx := context.Background()
	hl7Config, err := config.LoadHL7Config(ctx, test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}

	wantMapping := map[string]Internal{
		hl7Config.Gender.Male:    Male,
		hl7Config.Gender.Female:  Female,
		hl7Config.Gender.Unknown: Unknown,
	}
	c := NewConvertor(hl7Config)

	for k, v := range wantMapping {
		t.Run(fmt.Sprintf("%v-%v", k, v), func(t *testing.T) {
			if got, want := c.HL7ToInternal(k), v; got != want {
				t.Errorf("c.HL7ToInternal(%v)=%v, want %v", k, got, want)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	want := []Internal{Male, Female}
	gotFreqPertem := map[Internal]int{}
	runs := 1000

	for i := 0; i < runs; i++ {
		got := Random()
		if !contains(got, want) {
			t.Errorf("Random()=%v; want one of %v", got, want)
		}
		gotFreqPertem[got]++
	}

	delta := float64(runs) / 5.0
	wantFreq := float64(runs) / float64(len(want))
	for _, item := range want {
		if gotFreq := gotFreqPertem[item]; math.Abs(float64(gotFreq)-wantFreq) > delta {
			t.Errorf("gotFreq[%v] = %d, want within %v of %f", item, gotFreq, delta, wantFreq)
		}
	}
}

func contains(i Internal, items []Internal) bool {
	for _, item := range items {
		if i == item {
			return true
		}
	}
	return false
}
