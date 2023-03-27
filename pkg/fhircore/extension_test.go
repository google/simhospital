// Copyright 2023 Google LLC
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

package fhircore

import (
	"testing"

	pb "github.com/google/fhir/go/proto/google/fhir/proto/r4/core/datatypes_go_proto"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

const targetURL = "target_url"

var (
	ext1            = StringExtension(targetURL, "1")
	ext2            = StringExtension(targetURL, "2")
	ext3            = StringExtension(targetURL, "3")
	extNotMatching  = StringExtension("different_url", "4")
	extNotMatching2 = StringExtension("different_url", "5")
)

func TestFindFirstExtension(t *testing.T) {
	tcs := []struct {
		name string
		in   []*pb.Extension
		want *pb.Extension
	}{{
		name: "single extension returns the only extension",
		in:   []*pb.Extension{ext1},
		want: ext1,
	}, {
		name: "multiple extensions with the same URL returns the first one",
		in:   []*pb.Extension{ext1, ext2, ext3},
		want: ext1,
	}, {
		name: "multiple extensions with different URLs returns the first matching",
		in:   []*pb.Extension{extNotMatching, ext2, ext3},
		want: ext2,
	}, {
		name: "empty extensions returns nil",
		in:   []*pb.Extension{},
		want: nil,
	}, {
		name: "nil extensions returns nil",
		in:   nil,
		want: nil,
	}, {
		name: "non matching URL returns nil",
		in:   []*pb.Extension{extNotMatching},
		want: nil,
	}, {
		name: "URL without value returns is skipped",
		in:   []*pb.Extension{&pb.Extension{Url: &pb.Uri{}}, ext1},
		want: ext1,
	}, {
		name: "Empty extension is skipped",
		in:   []*pb.Extension{&pb.Extension{}, ext1},
		want: ext1,
	}}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := FindFirstExtension(tc.in, targetURL)
			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("FindFirstExtension(%v, %s) mismatch (-want, +got)=\n%s", tc.in, targetURL, diff)
			}
		})
	}
}

func TestFindLastExtension(t *testing.T) {
	tcs := []struct {
		name string
		in   []*pb.Extension
		want *pb.Extension
	}{{
		name: "single extension returns the only extension",
		in:   []*pb.Extension{ext1},
		want: ext1,
	}, {
		name: "multiple extensions with the same URL returns the last one",
		in:   []*pb.Extension{ext1, ext2, ext3},
		want: ext3,
	}, {
		name: "multiple extensions with different URLs returns the last matching",
		in:   []*pb.Extension{extNotMatching, ext2, ext3},
		want: ext3,
	}, {
		name: "empty extensions returns nil",
		in:   []*pb.Extension{},
		want: nil,
	}, {
		name: "nil extensions returns nil",
		in:   nil,
		want: nil,
	}, {
		name: "non matching URL returns nil",
		in:   []*pb.Extension{extNotMatching},
		want: nil,
	}, {
		name: "URL without value returns is skipped",
		in:   []*pb.Extension{&pb.Extension{Url: &pb.Uri{}}, ext1},
		want: ext1,
	}, {
		name: "Empty extension is skipped",
		in:   []*pb.Extension{&pb.Extension{}, ext1},
		want: ext1,
	}}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := FindLastExtension(tc.in, targetURL)
			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("FindLastExtension(%v, %s) mismatch (-want, +got)=\n%s", tc.in, targetURL, diff)
			}
			// (White-box testing) Make sure that invoking FindLastExtension doesn't leave the extensions in reversed order.
			got = FindLastExtension(tc.in, targetURL)
			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("Second invocation of FindLastExtension(%v, %s) mismatch (-want, +got)=\n%s", tc.in, targetURL, diff)
			}
		})
	}
}

func TestFindAllExtensions(t *testing.T) {
	tcs := []struct {
		name string
		in   []*pb.Extension
		want []*pb.Extension
	}{{
		name: "single extension returns the only extension",
		in:   []*pb.Extension{ext1},
		want: []*pb.Extension{ext1},
	}, {
		name: "multiple extensions returns all matching",
		in:   []*pb.Extension{extNotMatching, ext2, ext3},
		want: []*pb.Extension{ext2, ext3},
	}, {
		name: "empty extensions returns empty",
		in:   []*pb.Extension{},
		want: []*pb.Extension{},
	}, {
		name: "nil extensions returns empty",
		in:   nil,
		want: []*pb.Extension{},
	}, {
		name: "non matching URL returns empty",
		in:   []*pb.Extension{extNotMatching},
		want: []*pb.Extension{},
	}, {
		name: "URL without value is skipped",
		in:   []*pb.Extension{&pb.Extension{Url: &pb.Uri{}}, ext1},
		want: []*pb.Extension{ext1},
	}, {
		name: "Empty extension is skipped",
		in:   []*pb.Extension{&pb.Extension{}, ext1},
		want: []*pb.Extension{ext1},
	}}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := FindAllExtensions(tc.in, targetURL)
			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("FindAllExtensions(%v, %s) mismatch (-want, +got)=\n%s", tc.in, targetURL, diff)
			}
		})
	}
}

func TestDeleteAllExtensions(t *testing.T) {
	tcs := []struct {
		name string
		in   []*pb.Extension
		want []*pb.Extension
	}{{
		name: "single extension deletes the only extension",
		in:   []*pb.Extension{ext1},
		want: []*pb.Extension{},
	}, {
		name: "deletes if first one",
		in:   []*pb.Extension{ext1, extNotMatching, extNotMatching2},
		want: []*pb.Extension{extNotMatching, extNotMatching2},
	}, {
		name: "deletes if last one",
		in:   []*pb.Extension{extNotMatching, extNotMatching2, ext1},
		want: []*pb.Extension{extNotMatching, extNotMatching2},
	}, {
		name: "multiple extensions deletes all matching",
		in:   []*pb.Extension{ext2, extNotMatching, ext3, extNotMatching2},
		want: []*pb.Extension{extNotMatching, extNotMatching2},
	}, {
		name: "empty extensions doesn't delete anything",
		in:   []*pb.Extension{},
		want: []*pb.Extension{},
	}, {
		name: "non matching URL returns empty",
		in:   []*pb.Extension{extNotMatching},
		want: []*pb.Extension{extNotMatching},
	}, {
		name: "URL without value is skipped",
		in:   []*pb.Extension{&pb.Extension{Url: &pb.Uri{}}, ext1},
		want: []*pb.Extension{&pb.Extension{Url: &pb.Uri{}}},
	}, {
		name: "Empty extension is skipped",
		in:   []*pb.Extension{&pb.Extension{}, ext1},
		want: []*pb.Extension{&pb.Extension{}},
	}}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := DeleteAllExtensions(tc.in, targetURL)
			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("DeleteAllExtensions(%v, %s) mismatch (-want, +got)=\n%s", tc.in, targetURL, diff)
			}
		})
	}
}

// TODO: Add tests for the other GetXExtension methods.
func TestGetStringExtension(t *testing.T) {
	tcs := []struct {
		name    string
		in      []*pb.Extension
		wantErr bool
		want    string
	}{{
		name: "single extension returns the only extension",
		in:   []*pb.Extension{ext1},
		want: "1",
	}, {
		name:    "multiple extensions with the same URL returns error",
		in:      []*pb.Extension{ext1, ext2, ext3},
		wantErr: true,
	}, {
		name: "multiple extensions returns the only matching",
		in:   []*pb.Extension{extNotMatching, ext2, extNotMatching},
		want: "2",
	}, {
		name:    "extension without Value returns error",
		in:      []*pb.Extension{{Url: Uri(targetURL)}},
		wantErr: true,
	}, {
		name:    "extension with empty Extension_ValueX returns err",
		in:      []*pb.Extension{{Url: Uri(targetURL), Value: &pb.Extension_ValueX{}}},
		wantErr: true,
	}, {
		name: "extension with empty Extension_ValueX_StringValue returns err",
		in: []*pb.Extension{{Url: Uri(targetURL), Value: &pb.Extension_ValueX{
			Choice: &pb.Extension_ValueX_StringValue{},
		}}},
		wantErr: true,
	}, {
		name: "extension with empty Value returns empty",
		in:   []*pb.Extension{StringExtension(targetURL, "")},
		want: "",
	}, {
		name:    "empty extensions returns error",
		in:      []*pb.Extension{},
		wantErr: true,
	}, {
		name:    "nil extensions returns error",
		in:      nil,
		wantErr: true,
	}, {
		name:    "non matching URL returns error",
		in:      []*pb.Extension{extNotMatching},
		wantErr: true,
	}, {
		name:    "extension with different type returns error",
		in:      []*pb.Extension{BooleanExtension(targetURL, true)},
		wantErr: true,
	}, {
		name: "URL without value returns is skipped",
		in:   []*pb.Extension{&pb.Extension{Url: &pb.Uri{}}, ext1},
		want: "1",
	}, {
		name: "Empty extension is skipped",
		in:   []*pb.Extension{&pb.Extension{}, ext1},
		want: "1",
	}}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetStringExtension(tc.in, targetURL)
			if gotErr := (err != nil); gotErr != tc.wantErr {
				t.Errorf("GetStringExtension(%v, %s) got diff error: gotErr: %v, wantErr? %t", tc.in, targetURL, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("GetStringExtension(%v, %s) got %v, want %s", tc.in, targetURL, got, tc.want)
			}
		})
	}
}

func TestAddOrUpdateExtension(t *testing.T) {
	tcs := []struct {
		name       string
		inExisting []*pb.Extension
		inNew      *pb.Extension
		wantErr    bool
		want       []*pb.Extension
	}{{
		name:       "add new",
		inExisting: []*pb.Extension{},
		inNew:      ext1,
		want:       []*pb.Extension{ext1},
	}, {
		name:       "update existing",
		inExisting: []*pb.Extension{ext1},
		inNew:      ext2,
		want:       []*pb.Extension{ext2},
	}, {
		name:       "add even if there are other extensions",
		inExisting: []*pb.Extension{extNotMatching, ext1},
		inNew:      ext2,
		want:       []*pb.Extension{extNotMatching, ext2},
	}, {
		name:       "multiple matching extensions returns error",
		inExisting: []*pb.Extension{ext1, ext2},
		inNew:      ext3,
		wantErr:    true,
	}, {
		name:       "update even if existing has different type",
		inExisting: []*pb.Extension{BooleanExtension(targetURL, true)},
		inNew:      ext3,
		want:       []*pb.Extension{ext3},
	}, {
		name:       "can add with empty url",
		inExisting: []*pb.Extension{},
		inNew:      StringExtension("", "value"),
		want:       []*pb.Extension{StringExtension("", "value")},
	}, {
		name:       "can update with empty url",
		inExisting: []*pb.Extension{StringExtension("", "value1")},
		inNew:      StringExtension("", "value2"),
		want:       []*pb.Extension{StringExtension("", "value2")},
	}}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := AddOrUpdateExtension(tc.inExisting, tc.inNew)
			if gotErr := (err != nil); gotErr != tc.wantErr {
				t.Errorf("AddOrUpdateExtension(%v, %v) got diff error: gotErr: %v, wantErr? %t", tc.inExisting, tc.inNew, err, tc.wantErr)
			}
			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("AddOrUpdateExtension(%v, %v) mismatch (-want, +got)=\n%s", tc.inExisting, tc.inNew, diff)
			}
		})
	}
}

func TestAddNewExtensionsOnly(t *testing.T) {
	tcs := []struct {
		name       string
		inExisting []*pb.Extension
		inNew      []*pb.Extension
		want       []*pb.Extension
	}{{
		name:       "add new",
		inExisting: []*pb.Extension{},
		inNew:      []*pb.Extension{ext1},
		want:       []*pb.Extension{ext1},
	}, {
		name:       "don't update existing",
		inExisting: []*pb.Extension{ext1},
		inNew:      []*pb.Extension{ext2},
		want:       []*pb.Extension{ext1},
	}, {
		name:       "add even if there are other extensions",
		inExisting: []*pb.Extension{extNotMatching},
		inNew:      []*pb.Extension{ext1},
		want:       []*pb.Extension{extNotMatching, ext1},
	}, {
		name:       "only add once",
		inExisting: []*pb.Extension{},
		inNew:      []*pb.Extension{ext1, ext2},
		want:       []*pb.Extension{ext1},
	}, {
		name:       "add multiple",
		inExisting: []*pb.Extension{},
		inNew:      []*pb.Extension{ext1, StringExtension("other-url", "6")},
		want:       []*pb.Extension{ext1, StringExtension("other-url", "6")},
	}}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := AddNewExtensionsOnly(tc.inExisting, tc.inNew)
			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("AddNewExtensionsOnly(%v, %v) mismatch (-want, +got)=\n%s", tc.inExisting, tc.inNew, diff)
			}
		})
	}
}
