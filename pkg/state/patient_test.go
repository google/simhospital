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

package state

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/ir"
)

func TestPatient_GetOrder(t *testing.T) {
	p := Patient{
		Orders: make(map[string]*ir.Order),
	}
	orderID1 := "orderID1"

	if p.GetOrder(orderID1) != nil {
		t.Error("p.GetOrder(orderID1) is something, want <nil>")
	}

	// Add order with non-empty id.
	order1 := &ir.Order{
		OrderProfile: &ir.CodedElement{ID: "1", Text: "profile1"},
	}
	p.AddOrder(orderID1, order1)
	if diff := cmp.Diff(order1, p.GetOrder(orderID1)); diff != "" {
		t.Errorf("Patient.GetOrder(%q) mismatch (-want +got):\n%s", orderID1, diff)
	}
	if len(p.Orders) != 1 {
		t.Errorf("len(p.Orders) = %d, want %d", len(p.Orders), 1)
	}

	// Update the existing order.
	order1.Filler = "1234"
	p.AddOrder(orderID1, order1)
	if diff := cmp.Diff(order1, p.GetOrder(orderID1)); diff != "" {
		t.Errorf("Patient.GetOrder(%q) mismatch (-want, +got):\n%s", orderID1, diff)
	}
	if len(p.Orders) != 1 {
		t.Errorf("len(p.Orders) = %d, want %d", len(p.Orders), 1)
	}

	// Add an order with an empty ID
	// The ID is generated, and every order with an empty ID is treated as an unique order.
	orderNoID := &ir.Order{
		OrderProfile: &ir.CodedElement{ID: "2", Text: "profile2"},
	}
	p.AddOrder("", orderNoID)
	wantOrderID := fmt.Sprintf(generatedOrderIDPattern, 1)
	got := p.GetOrder(wantOrderID)
	if diff := cmp.Diff(orderNoID, got); diff != "" {
		t.Errorf("Patient.GetOrder(%q) mismatch (-want +got):\n%s", wantOrderID, diff)
	}
	if len(p.Orders) != 2 {
		t.Errorf("len(p.Orders) = %d, want %d", len(p.Orders), 2)
	}

	// Adding an identical order without order ID does not override the existing one.
	p.AddOrder("", orderNoID)
	wantOrderID = fmt.Sprintf(generatedOrderIDPattern, 2)
	got = p.GetOrder(wantOrderID)
	if diff := cmp.Diff(orderNoID, got); diff != "" {
		t.Errorf("Patient.GetOrder(%q) mismatch (-want +got):\n%s", wantOrderID, diff)
	}

	if len(p.Orders) != 3 {
		t.Errorf("len(p.Orders) = %d, want %d", len(p.Orders), 3)
	}
}

func TestPatient_PushPastVisit_PopPastVisit(t *testing.T) {
	p := Patient{}

	if _, err := p.PopPastVisit(); err == nil {
		t.Error("PopPastVisit() got nil error, want error because there are no past visits")
	}
	visitID1, visitID2 := uint64(12345), uint64(67890)
	p.PushPastVisit(visitID1)
	p.PushPastVisit(visitID2)

	var gotVisits []uint64
	for i := 0; i < 2; i++ {
		v, err := p.PopPastVisit()
		if err != nil {
			t.Fatalf("PopPastVisit() failed with %v", err)
		}
		gotVisits = append(gotVisits, v)
	}
	wantVisits := []uint64{visitID2, visitID1}
	if diff := cmp.Diff(wantVisits, gotVisits); diff != "" {
		t.Errorf("PopPastVisits() got diff (-want, +got):\n%s", diff)
	}
	if _, err := p.PopPastVisit(); err == nil {
		t.Error("PopPastVisit() got nil error, want error because there are no past visits")
	}
}
