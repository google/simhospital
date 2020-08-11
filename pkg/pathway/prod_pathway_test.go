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

package pathway

import (
	"context"
	"testing"
	"time"

	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/doctor"
	"github.com/google/simhospital/pkg/location"
	"github.com/google/simhospital/pkg/orderprofile"
	"github.com/google/simhospital/pkg/test"
	"github.com/google/simhospital/pkg/test/testclock"
)

func TestValidateProdPathways(t *testing.T) {
	ctx := context.Background()
	hl7Config, err := config.LoadHL7Config(ctx, test.MessageConfigProd)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigProd, err)
	}
	d, err := doctor.LoadDoctors(ctx, test.DoctorsConfigProd)
	if err != nil {
		t.Fatalf("LoadDoctors(%s) failed with %v", test.DoctorsConfigProd, err)
	}
	op, err := orderprofile.Load(ctx, test.OrderProfilesConfigProd, hl7Config)
	if err != nil {
		t.Fatalf("orderprofile.Load(%s, %+v) failed with %v", test.OrderProfilesConfigProd, hl7Config, err)
	}
	lm, err := location.NewManager(ctx, test.LocationsConfigProd)
	if err != nil {
		t.Fatalf("location.NewManager(%s) failed with %v", test.LocationsConfigProd, err)
	}
	p := &Parser{Clock: testclock.New(time.Now()), OrderProfiles: op, Doctors: d, LocationManager: lm}
	pathways, err := p.ParsePathways(ctx, test.PathwaysDirProd)
	if err != nil {
		t.Fatalf("ParsePathways(%s) failed with %v", test.PathwaysDirProd, err)
	}

	if len(pathways) == 0 {
		t.Fatalf("ParsePathways(%s) got empty pathways, want non empty", test.PathwaysDirProd)
	}
}
