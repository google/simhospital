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
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/simhospital/pkg/ir"
)

func TestPatient_GetOrder(t *testing.T) {
	p := Patient{
		Orders:      make(map[string]*ir.Order),
		PatientInfo: &ir.PatientInfo{},
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
	ec1 := p.PatientInfo.LatestEncounter()
	if diff := cmp.Diff([]*ir.Order{order1}, ec1.Orders); diff != "" {
		t.Errorf("ec.Orders mismatch (-want +got):\n%s", diff)
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
	if len(p.PatientInfo.Encounters) != 1 {
		t.Errorf("len(p.PatientInfo.Encounter) = %d, want %d", len(p.PatientInfo.Encounters), 1)
	}
	if diff := cmp.Diff([]*ir.Order{order1}, ec1.Orders); diff != "" {
		t.Errorf("ec.Orders mismatch (-want +got):\n%s", diff)
	}

	// Add an order with an empty ID
	// The ID is generated, and every order with an empty ID is treated as an unique order.
	orderNoID := &ir.Order{
		OrderProfile: &ir.CodedElement{ID: "2", Text: "profile2"},
	}
	p.AddOrder("", orderNoID)
	wantOrderID := "generated-1"
	got := p.GetOrder(wantOrderID)
	if diff := cmp.Diff(orderNoID, got); diff != "" {
		t.Errorf("Patient.GetOrder(%q) mismatch (-want +got):\n%s", wantOrderID, diff)
	}
	if len(p.Orders) != 2 {
		t.Errorf("len(p.Orders) = %d, want %d", len(p.Orders), 2)
	}
	if len(p.PatientInfo.Encounters) != 2 {
		t.Errorf("len(p.PatientInfo.Encounters) = %d, want %d", len(p.PatientInfo.Encounters), 2)
	}
	ec2 := p.PatientInfo.LatestEncounter()
	if diff := cmp.Diff([]*ir.Order{got}, ec2.Orders); diff != "" {
		t.Errorf("ec.Orders mismatch (-want +got):\n%s", diff)
	}

	// Adding an identical order without order ID does not override the existing one.
	p.AddOrder("", orderNoID)
	wantOrderID = "generated-2"
	got = p.GetOrder(wantOrderID)
	if diff := cmp.Diff(orderNoID, got); diff != "" {
		t.Errorf("Patient.GetOrder(%q) mismatch (-want +got):\n%s", wantOrderID, diff)
	}
	if len(p.PatientInfo.Encounters) != 3 {
		t.Errorf("len(p.PatientInfo.Encounters) = %d, want %d", len(p.PatientInfo.Encounters), 3)
	}
	ec3 := p.PatientInfo.LatestEncounter()
	if diff := cmp.Diff([]*ir.Order{got}, ec3.Orders); diff != "" {
		t.Errorf("ec.Orders mismatch (-want +got):\n%s", diff)
	}

	if len(p.Orders) != 3 {
		t.Errorf("len(p.Orders) = %d, want %d", len(p.Orders), 3)
	}
}

func TestPatient_GetDocument(t *testing.T) {
	p := Patient{
		Documents: make(map[string]*ir.Document),
	}

	docid1 := "docid1"

	// Add document with non-empty id.
	doc1 := &ir.Document{ContentLine: []string{"sample-text1"}}
	p.AddDocument(docid1, doc1)
	if diff := cmp.Diff(doc1, p.GetDocument(docid1)); diff != "" {
		t.Errorf("Patient.GetDocument(%v) mismatch (-want +got):\n%s", docid1, diff)
	}

	// Update the existing document.
	doc1.ContentLine = append(doc1.ContentLine, "sample-text2")
	if diff := cmp.Diff(doc1, p.GetDocument(docid1)); diff != "" {
		t.Errorf("Patient.Document(%v) mismatch (-want +got):\n%s", docid1, diff)
	}

	if len(p.Documents) != 1 {
		t.Errorf("len(p.Documents) = %d, want %d", len(p.Documents), 1)
	}

	// Add a Document with an empty ID.
	// The ID is generated, and every Document with an empty ID is treated as an unique document.
	docNoID := &ir.Document{}
	p.AddDocument("", docNoID)
	wantDocID := "generated-1"
	got := p.GetDocument(wantDocID)
	if diff := cmp.Diff(docNoID, got); diff != "" {
		t.Errorf("Patient.GetDocument(%v) mismatch (-want +got):\n%s", wantDocID, diff)
	}
	if len(p.Documents) != 2 {
		t.Errorf("len(p.Documents) = %d, want %d", len(p.Documents), 2)
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
