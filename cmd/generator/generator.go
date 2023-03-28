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

// Package main generates Go code for parsing HL7 messages, using the XML schemas.
// We generate a single Go module from all HL7 specification versions, using
// the most recent definition of each type. This relies on HL7 maintaining
// backwards compatibility.
// For an example of the generated code, see schema.go.
// TODO:
// - Represent HL7 null values
// - Consider using go generate (blog.golang.org/generate)
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Element represents the elements with a XML schema complex type.
type Element struct {
	Ref        string `xml:"ref,attr"`
	MinOccurs  string `xml:"minOccurs,attr"`
	MaxOccurs  string `xml:"maxOccurs,attr"`
	Deprecated bool
}

// IsHL7Segment returns whether the element is an HL7 segment.
func (e Element) IsHL7Segment() bool {
	return len(e.Ref) == 3
}

// IsHL7DataType returns whether the element is an HL7 data type.
func (e Element) IsHL7DataType() bool {
	return len(e.Ref) == 2
}

// ComplexType represents a XML schema complex type, used within the HL7 XML
// schemas. These represent segments and their fields in segments.xsd and
// fields.xsd respectively, and composite datatypes within datatypes.xsd.
type ComplexType struct {
	Name          string    `xml:"name,attr"`
	Elements      []Element `xml:"sequence>element"`
	Choices       []Element `xml:"choice>element"`
	Item          string    `xml:"annotation>appinfo>Item"`
	UpperCaseType string    `xml:"annotation>appinfo>Type"`
	// Some schemas use lower-case "type".
	LowerCaseType string `xml:"annotation>appinfo>type"`
	LongName      string `xml:"annotation>appinfo>LongName"`
	Version       string
}

// Type returns the type.
func (c ComplexType) Type() string {
	t := c.UpperCaseType
	if t == "" {
		t = c.LowerCaseType
	}
	if t == "varies" {
		return "Any"
	}
	return t
}

// IsHL7Segment returns true if this ComplexType represents a HL7 segment.
func (c ComplexType) IsHL7Segment() bool {
	// Note that it's a completely arbitrary property of the XML schema mapping
	// choice that segments don't have types and fields do.
	return c.Type() == ""
}

// IsHL7CompositeType returns true if this ComplexType represents a HL7
// composite type.
func (c ComplexType) IsHL7CompositeType() bool {
	return len(c.Elements) > 0
}

// SegmentName returns the name of the segment, eg MSH or PID, that c
// represents, or an undefined value if c doesn't represent a segment.
func (c ComplexType) SegmentName() string {
	return c.Name[0:3]
}

// MessageTypeName returns the name of the message type, eg ADT_A01, that c
// represents, or an undefined value if c doesn't represent a message.
func (c ComplexType) MessageTypeName() string {
	return nameify(c.Name)
}

// Schema represents a HL7 XML schema (ie segments.xsd, fields.xsd or
// datatypes.xsd)
type Schema struct {
	ComplexTypes []*ComplexType `xml:"complexType"`
}

// GoType returns the Go type name for a HL7 data type
func GoType(t string) string {
	// The HL7 specification uses a type of '*' or 'var' to denote that the type of
	// a field is contingent on the value of another field, and hence can't be
	// statically known
	if t == "*" || t == "var" || t == "varies" {
		return "Any"
	}
	return t
}

var (
	nonWordCharacters                 = regexp.MustCompile(`\W+`)           // Not a letter, digit or underscore.
	nonWordAndNotUnderscoreCharacters = regexp.MustCompile(`[^a-zA-Z0-9]+`) // Neither a letter nor a digit.
	// Matches a trailing trailing HL7 type in, e.g., "ID number (NM)".
	trailingType = regexp.MustCompile(` \([A-Z]{2,3}\)$`)
	// Matches how field names are represented in the XML schema.
	fieldTypeName = regexp.MustCompile(`^[A-Z0-9]{3}\.\d+\.CONTENT$`)
	// Matches how segment names are represented in the XML schema.
	segmentName = regexp.MustCompile(`^[A-Z0-9]{3}\.CONTENT$`)
	// Matches how message type names are represented in the XML schema.
	messageTypeName = regexp.MustCompile(`^[A-Z0-9]{3}\_[A-Za-z0-9]{3}\.CONTENT$`)
	// Matches how message subtype names are represented in the XML schema.
	// The extra "\.?" catches cases like PPV_PCA..ORDER_DETAIL_SEGMENT.CONTENT.
	// The hyphen "-" catches cases like SDR_S31.ANTI-MICROBIAL_DEVICE_DATA.CONTENT.
	messageSubtypeName = regexp.MustCompile(`^[A-Z0-9]{3}\_[A-Za-z0-9]{3}\.?\.[A-Za-z0-9_-]+\.CONTENT$`)
)

func remove(r *regexp.Regexp, s string) string {
	return r.ReplaceAllString(s, "")
}

// toFieldName returns a valid go field name derived from the long description
// for a HL7 field.
func toFieldName(longName string) string {
	return strings.Replace(strings.Title(nonWordCharacters.ReplaceAllString(remove(trailingType, removeHL7Prefix(longName)), " ")), " ", "", -1)
}

func toFieldNameWithoutUnderscore(longName string) string {
	return strings.Replace(strings.Title(nonWordAndNotUnderscoreCharacters.ReplaceAllString(remove(trailingType, removeHL7Prefix(longName)), " ")), " ", "", -1)
}

// nameify returns a valid go field name derived from the message type or
// segment name by removing the trailing "CONTENT" and replacing . with _
// A repeated "_" is replaced by a single "__". This supports issues like "PPV_PCA..ORDER_DETAIL_SEGMENT.CONTENT".
func nameify(rawName string) string {
	return strings.Replace(strings.Replace(strings.Replace(strings.TrimSuffix(removeHL7Prefix(rawName), ".CONTENT"), ".", "_", -1), "-", "_", -1), "__", "_", -1)
}

// removeHL7Prefix removes the HL7 prefix from field names that appear with such prefix in the specification.
func removeHL7Prefix(s string) string {
	keep := map[string]bool{
		"HL7 Data Type":  true,
		"HL7 Error Code": true,
	}
	if keep[s] {
		return s
	}

	s = strings.TrimPrefix(s, "HL7 ")
	s = strings.TrimPrefix(s, "HL7")

	return s
}

// docify cleans up the long name for a HL7 field by removing the trailing type,
// if present, and putting it into Title case.
func docify(longName string) string {
	return strings.Title(remove(trailingType, removeHL7Prefix(longName)))
}

func deduplicate(names []string) []string {
	counts := map[string]int{}
	for _, n := range names {
		counts[n]++
	}
	suffixes := map[string]int{}
	deduplicated := make([]string, len(names))
	for i, n := range names {
		if counts[n] == 1 {
			deduplicated[i] = n
		} else {
			suffix, ok := suffixes[n]
			if !ok {
				suffix = 1
			}
			deduplicated[i] = fmt.Sprintf("%s%d", n, suffix)
			suffixes[n] = suffix + 1
		}
	}
	return deduplicated
}

func tag(name string, required bool) string {
	// We remove any quotes from the name, or it will produce a malformed tag.
	return fmt.Sprintf("hl7:\"%t,%s\"", required, strings.Replace(name, "\"", "", -1))
}

// outputCompositeType writes a Go struct to p that represents the HL7 complex
// type described by c. The generated code looks something like this:
//
//	type AD struct {
//	  StreetAddress *ST // AD.1.CONTENT
//	  OtherDesignation *ST // AD.2.CONTENT
//	  City *ST // AD.3.CONTENT
//	  ...
//	}
//
// Deprecated types have type *NUL.
func outputCompositeType(p *Printer, c *ComplexType, fields map[string]*ComplexType) {
	p.P("")
	p.P("// %s represents the corresponding HL7 datatype.", c.Name)
	p.P("// Definition from HL7 %s", c.Version)
	p.P("type %s struct {", c.Name)
	{
		p.In()
		for _, e := range c.Elements {
			f, ok := fields[e.Ref+".CONTENT"]
			if ok {
				if e.MaxOccurs != "1" {
					log.Fatalf("Datatype %s got e.MaxOccurs=%s, want 1: datatypes shouldn't have repeated components", e.Ref, e.MaxOccurs)
				}
				if e.Deprecated {
					p.P("Deprecated%s *NUL `%s`", toFieldNameWithoutUnderscore(f.LongName), tag(docify(f.LongName), e.MinOccurs != "0"))
				} else {
					p.P("%s *%s `%s`", toFieldNameWithoutUnderscore(f.LongName), f.Type(), tag(docify(f.LongName), e.MinOccurs != "0"))
				}
			} else {
				log.Fatalf("Missing composite data type field: %s", e.Ref)
			}
		}
		p.Out()
	}
	p.P("}")
}

// outputMessageType writes a Go struct to p that represents the HL7 message
// described by c. The generated code looks something like this:
//
// type ADT_A01 struct {
//
//		MSH       *MSH
//		PID       *PID
//		ROL1      []ROL
//		ROL2      []ROL
//		PROCEDURE []ADT_A01_PROCEDURE
//	 ...
//		Other     []Segment
//
// }
func outputMessageType(p *Printer, c *ComplexType, spec *Specification) {
	name := nameify(c.Name)
	if !messageTypeName.MatchString(c.Name) && !messageSubtypeName.MatchString(c.Name) && c.Name != "ACK.CONTENT" {
		log.Fatalf("Not a message name: %s", c.Name)
	}

	elements := make([]Element, 0, len(c.Elements)+len(c.Choices))
	elements = append(elements, c.Elements...)
	elements = append(elements, c.Choices...)
	firstChoiceElement := len(c.Elements)

	typeNames := make([]string, len(elements))
	for i, e := range elements {
		s, ok := spec.Segments[e.Ref+".CONTENT"]
		if !ok {
			s, ok = spec.MessageTypes[e.Ref+".CONTENT"]
		}
		if ok {
			typeNames[i] = toFieldName(nameify(s.Name))
		}
	}
	segNames := make([]string, len(typeNames))
	segNameExists := make([]bool, len(typeNames))
	for i, n := range typeNames {
		// Remove the message type from the field name for sub-structs.
		segNames[i] = strings.TrimPrefix(n, strings.Split(c.Name, ".")[0]+"_")
		if n != "" {
			segNameExists[i] = true
		}
	}
	segNames = deduplicate(segNames)

	p.P("// %s represents the corresponding HL7 message type.", name)
	p.P("// Definition from HL7 %s", c.Version)
	p.P("type %s struct {", name)
	{
		p.In()
		for i, e := range elements {
			if i == firstChoiceElement {
				p.P("// Only one of the following fields will be set.")
			}
			if segNameExists[i] {
				t := GoType(typeNames[i])
				if e.MaxOccurs == "1" {
					t = "*" + t
				} else {
					t = "[]" + t
				}
				p.P("%s %s `%s`", segNames[i], t, tag(segNames[i], e.MinOccurs != "0"))
			} else {
				p.P("// Missing: %s", e.Ref)
			}
		}
		// Add field for all 'other' segs, e.g. z segments
		p.P("%s %s", "Other", "[]interface{}")
		p.Out()
	}
	p.P("}")
	p.P("")
	// Use different function names here so we can differentiate msg types and substructs.
	if messageSubtypeName.MatchString(c.Name) {
		p.P("func (s *%s) MessageTypeSubStructName() string {", name)
	} else {
		p.P("func (s *%s) MessageTypeName() string {", name)
	}
	p.In()
	p.P("return %q", name)
	p.Out()
	p.P("}")
}

// outputSegment writes a Go struct to p that represents the HL7 segment
// described by c. The generated code looks something like this:
//
//	type PID struct {
//	  SetIDPID *SI  // PID.1.CONTENT
//	  PatientID *CX  // PID.2.CONTENT
//	  ...
//	}
//
// Deprecated fields have type *NUL.
func outputSegment(p *Printer, c *ComplexType, fields map[string]*ComplexType) {
	name := c.Name[0:3] // HL7 segment names are always the first three characters
	if !segmentName.MatchString(c.Name) {
		log.Fatalf("Not a segment name: %s", c.Name)
	}

	amendElements(c)
	fieldNames := make([]string, len(c.Elements))
	for i, e := range c.Elements {
		f, ok := fields[e.Ref+".CONTENT"]
		if ok {
			fieldNames[i] = toFieldName(f.LongName)
		}
	}
	fieldNames = deduplicate(fieldNames)

	p.P("// %s represents the corresponding HL7 segment.", name)
	p.P("// Definition from HL7 %s", c.Version)
	p.P("type %s struct {", name)
	{
		p.In()
		for i, e := range c.Elements {
			f, ok := fields[e.Ref+".CONTENT"]
			if ok {
				// TODO: Find a way to represent HL7 NULLs
				t := GoType(f.Type())
				if e.MaxOccurs == "1" {
					t = "*" + t
				} else {
					t = "[]" + t
				}
				if e.Deprecated {
					p.P("Deprecated%s *NUL `%s` // %s-%d", strings.Replace(fieldNames[i], "_", "", -1), tag(docify(f.LongName), e.MinOccurs != "0"), name, i+1)
				} else {
					p.P("%s %s `%s` // %s-%d", strings.Replace(fieldNames[i], "_", "", -1), t, tag(docify(f.LongName), e.MinOccurs != "0"), name, i+1)
				}
			} else {
				p.P("// Missing: %s", e.Ref)
			}
		}
		p.Out()
	}
	p.P("}")
	p.P("")
	p.P("func (s *%s) SegmentName() string {", name)
	p.In()
	p.P("return %q", name)
	p.Out()
	p.P("}")
	p.P("")
}

// outputGenericHL7Segment writes a Go struct to p that represents a generic HL7 segment
// e.g., Z segments.
func outputGenericHL7Segment(p *Printer) {
	name := "GenericHL7Segment"
	elements := map[string]string{"segment": "[]byte"}
	p.P("// %s represents the corresponding HL7 segment type.", name)
	p.P("type %s struct {", name)
	{
		p.In()
		for e, t := range elements {
			p.P("%s %s", e, t)
		}
		p.Out()
	}
	p.P("}")
	p.P("")
	p.P("func (s *%s) SegmentName() string {", name)
	p.In()
	p.P("return %q", name)
	p.Out()
	p.P("}")
}

func outputParseSegmentFunctions(p *Printer, spec *Specification) {
	for _, k := range sortedMapKeys(spec.Segments) {
		c := spec.Segments[k]
		name := c.SegmentName()
		p.P("// %s returns the first %s segment within the message, or nil if there isn't one.", name, name)
		p.P("func (m *Message) %s() (*%s, error) {", name, name)
		{
			p.In()
			p.P("ps, err := m.Parse(%q)", name)
			p.P("pst, ok := ps.(*%s)", name)
			p.P("if ok {")
			{
				p.In()
				p.P("return pst, err")
				p.Out()
			}
			p.P("}")
			p.P("return nil, err")
			p.Out()
		}
		p.P("}")
	}
}

func outputParseAllSegmentsFunctions(p *Printer, spec *Specification) {
	for _, k := range sortedMapKeys(spec.Segments) {
		c := spec.Segments[k]
		name := c.SegmentName()
		p.P("// All%s returns a slice containing all %s segments within the message,", name, name)
		p.P("// or an empty slice if there aren't any.")
		p.P("func (m *Message) All%s() ([]*%s, error) {", name, name)
		{
			p.In()
			p.P("pss, err := m.ParseAll(%q)", name)
			p.P("return pss.([]*%s), err", name)
			p.Out()
		}
		p.P("}")
	}
}

// outputMessageFunctions writes Go methods on Message to p that allow segments
// described in s to be parsed, unless blocklisted. These include a method to
// parse all segments within a message:
//
//	m.ParseAll()
//
// Methods to parse the first occurrence of a given segment:
//
//	pid, err := m.PID()
//
// Methods to parse every occurrence of a given segment, returning a slice:
//
//	pids, err := m.AllPID()
func outputMessageFunctions(p *Printer, spec *Specification) {
	outputParseSegmentFunctions(p, spec)
	p.P("")
	outputParseAllSegmentsFunctions(p, spec)
}

func outputSegmentTypeMap(p *Printer, spec *Specification) {
	for _, k := range sortedMapKeys(spec.Segments) {
		c := spec.Segments[k]
		p.P("Types[%q] = reflect.TypeOf(%s{})", c.SegmentName(), c.SegmentName())
	}
}

func outputMessageTypeMap(p *Printer, spec *Specification) {
	for _, k := range sortedMapKeys(spec.MessageTypes) {
		c := spec.MessageTypes[k]
		p.P("Types[%q] = reflect.TypeOf(%s{})", c.MessageTypeName(), c.MessageTypeName())
	}
}

func parseXML(name string, v interface{}) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	return xml.NewDecoder(f).Decode(v)
}

func sortableVersion(v string) int {
	n, _ := strconv.Atoi(strings.Replace(v, ".", "", -1))
	if n < 200 {
		n *= 10
	}
	return n
}

func toHl7VersionName(v int) string {
	str := strconv.Itoa(v)
	var result string
	for i, c := range str {
		if i > 0 {
			result += "."
		}
		result += string(c)
	}
	return result
}

type ByVersion []string

func (v ByVersion) Len() int           { return len(v) }
func (v ByVersion) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v ByVersion) Less(i, j int) bool { return sortableVersion(v[i]) < sortableVersion(v[j]) }

type Specification struct {
	Fields         map[string]*ComplexType
	Segments       map[string]*ComplexType
	CompositeTypes map[string]*ComplexType
	MessageTypes   map[string]*ComplexType
}

func NewSpecification() *Specification {
	return &Specification{
		Fields:         map[string]*ComplexType{},
		Segments:       map[string]*ComplexType{},
		CompositeTypes: map[string]*ComplexType{},
		MessageTypes:   map[string]*ComplexType{},
	}
}

// ParseSpecification parses the HL7 specification from the specified directory.
func ParseSpecification(directory string, version string, blockListed map[string]bool, s *Specification) error {
	messageSchema := &Schema{}
	err := parseXML(filepath.Join(directory, "fields.xsd"), messageSchema)
	if err != nil {
		return err
	}
	err = parseXML(filepath.Join(directory, "segments.xsd"), messageSchema)
	if err != nil {
		return err
	}

	dataTypeSchema := &Schema{}
	err = parseXML(filepath.Join(directory, "datatypes.xsd"), dataTypeSchema)
	if err != nil {
		return err
	}

	messageTypeSchema := &Schema{}
	messageTypeFiles, err := filepath.Glob(filepath.Join(directory, "???_???.xsd"))
	if err != nil {
		return err
	}
	// Prepend ACK.xsd to keep the alphabetical order of message type structs in generated.go
	messageTypeFiles = append([]string{filepath.Join(directory, "ACK.xsd")}, messageTypeFiles...)
	for _, f := range messageTypeFiles {
		err = parseXML(f, messageTypeSchema)
		if err != nil {
			return err
		}
	}

	for _, c := range messageSchema.ComplexTypes {
		if blockListed[c.Name] {
			continue
		}
		c.Version = version
		if c.IsHL7Segment() {
			if segmentName.MatchString(c.Name) {
				s.Segments[c.Name] = c
			}
			continue
		}
		if !fieldTypeName.MatchString(c.Name) {
			log.Fatalf("Not a field name: %s", c.Name)
		}
		s.Fields[c.Name] = c
	}

	for _, c := range dataTypeSchema.ComplexTypes {
		if blockListed[c.Name] {
			continue
		}
		c.Version = version
		s.Fields[c.Name] = c
	}

	for _, c := range dataTypeSchema.ComplexTypes {
		if blockListed[c.Name] {
			continue
		}
		c.Version = version
		if c.IsHL7CompositeType() {
			s.CompositeTypes[c.Name] = c
		}
	}

	for _, c := range messageTypeSchema.ComplexTypes {
		if blockListed[c.Name] {
			continue
		}
		c.Version = version
		if c.IsHL7Segment() { // TODO: should be IsHL7MessageType
			s.MessageTypes[c.Name] = c
		}
	}

	return nil
}

// parseExtras parses the locally-synthesized specs for extra messages, segments, or fields.
func parseExtras(messageFiles []string, blockListed map[string]bool, s *Specification) error {
	messageTypeSchema := &Schema{}
	for _, f := range messageFiles {
		log.Printf("Processing extra file %s", f)
		err := parseXML(f, messageTypeSchema)
		if err != nil {
			return err
		}
	}
	for _, c := range messageTypeSchema.ComplexTypes {
		if blockListed[c.Name] {
			continue
		}
		c.Version = "SYNTHETIC"
		switch {
		case c.IsHL7Segment() && segmentName.MatchString(c.Name):
			s.Segments[c.Name] = c
		case c.IsHL7Segment():
			s.MessageTypes[c.Name] = c
		case fieldTypeName.MatchString(c.Name):
			s.Fields[c.Name] = c
		}
	}
	return nil
}

func allElements(c *ComplexType) []Element {
	elements := make([]Element, 0, len(c.Elements)+len(c.Choices))
	elements = append(elements, c.Elements...)
	elements = append(elements, c.Choices...)
	return elements
}

// escape escapes s for use as a Go identifier
func escape(s string) string {
	if s == "var" {
		return s + "_"
	}
	return s
}

// Group represents a group of related HL7 segments, and sub groups, for which
// we'll generate a struct. Examples would include a top level message type,
// like ADT_A01, or a group of related segments within a message, for example
// those referring to a given PID segment within an ORU_R01 message.
type Group struct {
	Name     string
	Elements []GroupElement
}

// NameForType returns the Go type we'll use for this group - ie the name of
// the generated group.
func (g Group) NameForType() string {
	// TODO: Remove the v2 suffix once the older API has been removed,
	// it's there now to prevent name collisions
	return escape(g.Name + "v2")
}

// UniqueKeyForElement returns a string that can be used as a globally unique
// key for a given field. It's generated by concatenating the name of the
// group with the name of the field, to make, for example ADT_A01.PID.
func (g Group) UniqueKeyForElement(i int) string {
	return g.NameForType() + "." + g.Elements[i].NameForField()
}

// GroupElement represents either a segment, or a sub group, within a Group.
type GroupElement struct {
	DataType         string // One of DataType or Segment or Group will be set
	Segment          string
	Group            *Group
	Required         bool
	Repeated         bool
	DeduplicatedName string
}

// NameForType returns the Go type we'll use for values of this field's type.
func (g GroupElement) NameForType() string {
	if g.Group != nil {
		return g.Group.NameForType()
	} else if len(g.Segment) > 0 {
		return escape(g.Segment)
	} else {
		return escape(g.DataType)
	}
}

// NameForAccessorMethod returns the name of the method we'll generate to access
// this field. Individual segments are of the form MSH(), repeated segments are
// of the form AllNTE(), sub groups are of the form GroupByPID().
func (g GroupElement) NameForAccessorMethod() string {
	name := ""
	if g.Group != nil {
		name = "GroupBy" + g.DeduplicatedName
	} else {
		name = g.DeduplicatedName
		if g.Repeated {
			name = "All" + name
		}
	}
	return escape(name)
}

// NameForField returns the name of the (unexposed) field we'll generate to
// hold the value for this element.
func (g GroupElement) NameForField() string {
	return escape(strings.ToLower(g.DeduplicatedName))
}

// buildGroup takes the definition of a message type from the specification,
// and returns a Group.
// The function first builds a set of nested groups directly from the HL7
// specification. These end up being pretty deeply nested, eg:
// ORU_R01 > PATIENT_RESULT > PATIENT > ORDER_OBSERVATION
// It then folds subgroups into their parents where possible, leaving
// something like this:
// ORU_R01 > (PATIENT_RESULT + PATIENT) > ORDER_OBSERVATION
// Before using the name of a segment within the group as it's name, as that's
// generally how people think about messages, eg:
// ORU_R01 > PID > OBR
func buildGroup(c *ComplexType, spec *Specification) *Group {
	group := buildRawGroup(c, spec)
	inlineSubGroups(group)
	generateNames(group)
	return group
}

// buildRawGroup takes the definition of a message type from the specification,
// and returns a group that directly reflects the underlying message type
// without further transformation.
func buildRawGroup(c *ComplexType, spec *Specification) *Group {
	group := &Group{
		Name: strings.Replace(c.Name, ".CONTENT", "", 1),
	}
	elements := allElements(c)
	group.Elements = make([]GroupElement, 0, len(elements))
	for _, e := range elements {
		if e.IsHL7Segment() {
			group.Elements = append(group.Elements, GroupElement{Segment: e.Ref})
		} else if e.IsHL7DataType() {
			group.Elements = append(group.Elements, GroupElement{DataType: e.Ref})
		} else {
			next, ok := spec.MessageTypes[e.Ref+".CONTENT"]
			if ok {
				group.Elements = append(group.Elements, GroupElement{Group: buildGroup(next, spec)})
			} else {
				if e.Ref != "anyZSegment" && e.Ref != "anyHL7Segment" && e.Ref != "IsHL7DataType" {
					panic("Bad ref: " + e.Ref)
				}
				continue
			}
		}
		tail := len(group.Elements) - 1
		if e.MinOccurs == "1" {
			group.Elements[tail].Required = true
		}
		if e.MaxOccurs != "1" {
			group.Elements[tail].Repeated = true
		}
	}
	return group
}

// inlineSubGroups merges the elements of a sub group into the parent group, if
// possible. Merging is possible if the group isn't repeated, and none of the
// segments within it collide with those in the parent.
// For example, the HL7 message type definition defines the structure
// of the ORU^R01 message to be:
// ORU_R01 > PATIENT_RESULT > PATIENT > ORDER_OBSERVATION
// However, PATIENT is not repeated, and doesn't contain any segments that are
// already used within PATIENT_RESULT, so we can restructure as:
// ORU_R01 > (PATIENT_RESULT + PATIENT) > ORDER_OBSERVATION
func inlineSubGroups(g *Group) {
	segments := map[string]bool{}
	for _, element := range g.Elements {
		if element.Group != nil {
			inlineSubGroups(element.Group)
		} else {
			segments[element.Segment] = true
		}
	}

	newElements := make([]GroupElement, 0, len(g.Elements))
	for _, element := range g.Elements {
		if element.Group == nil {
			newElements = append(newElements, element)
		} else {
			canInline := true
			for _, subElement := range element.Group.Elements {
				if subElement.Group == nil && segments[subElement.Segment] {
					canInline = false
					break
				}
			}
			if canInline && !element.Repeated {
				for _, subElement := range element.Group.Elements {
					if subElement.Group == nil {
						segments[subElement.Segment] = true
					}
					newElements = append(newElements, subElement)
				}
			} else {
				newElements = append(newElements, element)
			}
		}
	}
	g.Elements = newElements
}

// generateNames creates names for all groups, and the elements within them,
// based on key segments (defined the name of the first required segment within
// that group, if it has any, otherwise the first segment). Names of sub groups
// are prefixed with the name of parent groups to avoid collisions.
func generateNames(g *Group) {
	names := make([]string, 0, len(g.Elements))
	for _, e := range g.Elements {
		if e.Group != nil {
			names = append(names, keySegment(e.Group))
		} else if len(e.Segment) > 0 {
			names = append(names, e.Segment)
		} else {
			names = append(names, e.DataType)
		}
	}
	names = deduplicate(names)
	for i, e := range g.Elements {
		g.Elements[i].DeduplicatedName = names[i]
		if e.Group != nil {
			// TODO: Some of these type names can get pretty long, eg:
			// ORU_R01_PID_OBR_OBX. If this becomes an issue, we could also
			// use the shortest unique prefix, making that example ORU_R01_OBX.
			// On the other hand, it makes the origin of the group more explicit.
			e.Group.Name = g.Name + "_" + names[i]
			generateNames(e.Group)
		}
	}
}

// keySegment returns the name of the first required segment within the group,
// or the name of the first segment if all are optional.
func keySegment(g *Group) string {
	name := ""
	for _, e := range g.Elements {
		if e.Group == nil {
			if name == "" {
				name = e.Segment
			}
			if e.Required {
				name = e.Segment
				break
			}
		}
	}
	return name
}

// collectGroups returns a list containing the group g, and all subgroups below it.
func collectGroups(g *Group) []*Group {
	collected := make([]*Group, 1, 1)
	collected[0] = g
	for _, e := range g.Elements {
		if e.Group != nil {
			collected = append(collected, collectGroups(e.Group)...)
		}
	}
	return collected
}

// outputAccessor outputs a Go method to p that will provide access to element e
// within group g
func outputAccessor(p *Printer, g *Group, e GroupElement) {
	var t string
	if e.Repeated {
		t = "[]*" + e.NameForType()
	} else {
		t = "*" + e.NameForType()
	}
	p.P("func (m *%s) %s() %s {", g.NameForType(), e.NameForAccessorMethod(), t)
	p.In()
	p.P("return m.%s", e.NameForField())
	p.Out()
	p.P("}")
	p.P("")
}

// outputYAMLMarshaler outputs a Go method to p that will customise YAML
// marshalling for the group. This is necessary as no fields are exported.
// TODO: This is mainly used for manual inspection while debugging.
// There may be a format that's easier to read.
func outputYAMLMarshaler(p *Printer, g *Group) {
	p.P("func (m %s) MarshalYAML() (interface{}, error) {", g.NameForType())
	p.In()
	p.P("return map[string]interface{}{")
	p.In()
	for _, e := range g.Elements {
		p.P("%q: m.%s,", e.NameForField(), e.NameForField())
	}
	p.Out()
	p.P("}, nil")
	p.Out()
	p.P("}")
}

// outputGroup outputs a Go structure definition to p for the group g, together
// with the accessor methods for the fields within it.
func outputGroup(p *Printer, g *Group) {
	// TODO: Print out a summary of the group structure in text at the top of the structure definition.
	gs := collectGroups(g)
	for _, g := range gs {
		p.P("type %s struct {", g.NameForType())
		p.In()
		for _, e := range g.Elements {
			comment := ""
			if e.Required {
				comment = " // Required"
			}
			if e.Repeated {
				p.P("%s []*%s%s", e.NameForField(), e.NameForType(), comment)
			} else {
				p.P("%s *%s%s", e.NameForField(), e.NameForType(), comment)
			}
		}
		p.Out()
		p.P("}")
		p.P("")
		for _, e := range g.Elements {
			outputAccessor(p, g, e)
		}
		p.P("")
		outputYAMLMarshaler(p, g)
	}
}

// GroupPath represents a path through one set of nested fields within a
// group. Each element of the path holds a group, and the index of a field
// within in.
type GroupPath []GroupPathElement

type GroupPathElement struct {
	Group *Group
	Index int
}

func (p GroupPath) Copy() GroupPath {
	c := make([]GroupPathElement, len(p))
	copy(c, p)
	return c
}

// Segment returns the segment at the end of this path, ie OBR in
// ORU_R01 > PID > OBR. Will panic if p doesn't actually end in a segment.
func (p GroupPath) Segment() string {
	return p[len(p)-1].Group.Elements[p[len(p)-1].Index].Segment
}

// traverseSegments walks a nested set of groups, starting from start,
// calling f(p) for each field that holds a segment, where p is the path to
// that field. Modifying p within p() will change traversal.
func traverseSegments(start GroupPath, f func(GroupPath)) {
	path := start.Copy()
	for {
		tail := len(path) - 1
		if path[tail].Index >= len(path[tail].Group.Elements) {
			// Traversal of this group has finished. Pop the last element off the end
			// of the path and move to the next field within the parent group.
			path = path[0:tail]
			if len(path) == 0 {
				break
			} else {
				path[tail-1].Index++
				continue
			}
		}
		element := path[tail].Group.Elements[path[tail].Index]
		if element.Group != nil {
			path = append(path, GroupPathElement{element.Group, 0})
		} else {
			f(path)
			path[tail].Index++
		}
	}
}

type StringSet map[string]bool

// buildFollowSets computes the follow set for all segments within g and it's
// nested sub groups. A follow set is the set of segments that may
// legitimately follow a given segment within a message, according to the
// specification. This is used during parsing to decide whether a given
// segment should be ignored or not.
// The follow set is a map that looks like the following:
// "ADT_A01.pid": StringSet{
//
//	 "PD1": true,
//	 "OBX": true,
//	 ...
//	}
//
// Indicating that segments PD1 and OBX (and others) can legitimately follow
// the PID segment within an ADT_A01 message.
func buildFollowSets(g *Group, fs map[string]StringSet) {
	start := GroupPath{{g, 0}}
	traverseSegments(start, func(path GroupPath) {
		tail := path[len(path)-1]
		f := StringSet{}
		fs[tail.Group.UniqueKeyForElement(tail.Index)] = f
		buildFollowSet(path, f)
	})
}

// buildFollowSet computes the follow set for the segment described by path.
// This will include the segment itself if it's repeated, all segments that
// follow it in a depth-first traversal of the nested groups, and precede it
// if any parent groups can repeat.
// Technically, traversal can stop at the first required segment, as if that's
// missing the message is invalid. We currently choose to include all segments
// to better handle broken messages.
func buildFollowSet(path GroupPath, f StringSet) {
	path = path.Copy()
	tail := len(path) - 1
	// Add the current segment if it repeats
	element := path[tail].Group.Elements[path[tail].Index]
	if !element.Repeated {
		path[tail].Index++
	}
	// Add all following segments from a depth-first traversal
	traverseSegments(path, func(p GroupPath) {
		f[p.Segment()] = true
	})
	// Add preceding segments if parent groups can repeat
	for i := len(path) - 2; i >= 0; i-- {
		if path[i].Index < len(path[i].Group.Elements) && path[i].Group.Elements[path[i].Index].Repeated {
			path[i+1].Index = 0
		} else {
			path[i+1].Index = len(path[i+1].Group.Elements)
		}
	}
	path[0].Index = len(path[0].Group.Elements) // The first group can't repeat
	traverseSegments(path, func(p GroupPath) {
		f[p.Segment()] = true
	})
}

func outputFollowSets(p *Printer, fs map[string]StringSet) {
	p.P("var FollowSets = map[string]StringSet{")
	p.In()
	fsKeys := make([]string, 0)
	for k := range fs {
		fsKeys = append(fsKeys, k)
	}
	sort.Strings(fsKeys)
	for _, key := range fsKeys {
		f := fs[key]
		p.P("%q: StringSet{", key)
		p.In()
		fKeys := make([]string, 0)
		for k := range f {
			fKeys = append(fKeys, k)
		}
		sort.Strings(fKeys)
		for _, segment := range fKeys {
			p.P("%q: true,", segment)
		}
		p.Out()
		p.P("},")
	}
	p.Out()
	p.P("}")
}

// A Printer writes formatted lines to an output stream, keeping track of indentation level.
type Printer struct {
	w             io.Writer
	codeIndent    uint
	commentIndent uint
}

// NewPrinter returns a new Printer that writes to w.
func NewPrinter(w io.Writer) *Printer {
	return &Printer{
		w: w,
	}
}

// In increases the level of indentation for subsequent lines.
func (p *Printer) In() {
	p.codeIndent++
}

// Out decreases the level of indentation for subsequent lines.
func (p *Printer) Out() {
	p.codeIndent--
}

// P prints a formatted line to the output stream, indenting it by the current indentation level
// using tab characters.
func (p *Printer) P(format string, a ...interface{}) error {
	var err error

	if strings.HasPrefix(format, "// ") {
		format = "//" + strings.Repeat(" ", int(p.commentIndent)+1) + format[3:]
	}
	if _, err = fmt.Fprintf(p.w, strings.Repeat("\t", int(p.codeIndent))); err != nil {
		return err
	}

	if _, err = fmt.Fprintf(p.w, format, a...); err != nil {
		return err
	}

	_, err = fmt.Fprintf(p.w, "\n")
	return err
}

type sliceFlags []string

func (i *sliceFlags) String() string {
	return strings.Join(*i, " ")
}

func (i *sliceFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	schemaDirectory := flag.String("schema", "", "Directory containing the HL7 schema.")
	extraDirectory := flag.String("extra_dir", "", "Directory containing synthetic message XSD files.")
	maxVersion := flag.Int("max_hl7_version", 251, "The maximum HL7v2 version to generate schemas from, as a three digit integer: 240 for 2.4, 251 for 2.5.1, etc.")
	var inputBlockListed sliceFlags
	flag.Var(&inputBlockListed, "block_list", "Segments/messages/types to skip when parsing from xsd.  This flag can be specified multiple times.  i.e. --block_list=PPX --block_list=ORU")
	flag.Parse()

	blockListed := loadBlocklisted(inputBlockListed)

	s := buildSpecification(blockListed, *maxVersion, *schemaDirectory, *extraDirectory)
	p := NewPrinter(os.Stdout)

	log.Println("Generating code...")
	outputHeader(s, p, *maxVersion)
	outputSpecification(s, p, blockListed)
}

func loadBlocklisted(inputBlockListed sliceFlags) map[string]bool {
	blockListed := map[string]bool{
		// Don't generate code for composite types for which we provide our own
		// bespoke implementation.
		"FT":           true,
		"TS":           true,
		"TX":           true,
		"TX_CHALLENGE": true,

		// Don't generate code for the composite types representing an arbitrary
		// HL7 segment, as we'll provide our own bespoke implementation
		"anyHL7Segment.TYPE": true,
		"anyZSegment.TYPE":   true,
	}
	for _, name := range inputBlockListed {
		blockListed[name] = true
	}
	return blockListed
}

func loadVersions(schemaDirectory string) []string {
	directories, err := filepath.Glob(filepath.Join(schemaDirectory, "*"))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found directories: %s", directories)

	versions := []string{}
	for _, d := range directories {
		s, err := os.Stat(d)
		if err == nil && s.IsDir() {
			versions = append(versions, filepath.Base(d))
		}
	}

	return versions
}

func buildSpecification(blockListed map[string]bool, maxVersion int, schemaDirectory, extraDirectory string) *Specification {
	versions := loadVersions(schemaDirectory)
	sort.Sort(ByVersion(versions))
	log.Printf("Found versions: %s", versions)

	s := NewSpecification()
	for _, v := range versions {
		sv := sortableVersion(v)
		f := filepath.Join(schemaDirectory, v)
		if sv <= maxVersion {
			log.Printf("Processing version %s from %s", v, f)
			err := ParseSpecification(f, v, blockListed, s)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Printf("Skipping version %s and above (--max_hl7_version=%d)", v, maxVersion)
			break
		}
	}

	for _, k := range sortedMapKeys(s.Fields) {
		amendElements(s.Fields[k])
	}

	// Ignoring errors since the only possibility is for a malformed pattern
	syntheticMsgFiles, _ := filepath.Glob(filepath.Join(extraDirectory, "???_???.xsd"))
	if err := parseExtras(syntheticMsgFiles, blockListed, s); err != nil {
		log.Fatal(err)
	}
	log.Printf("Finished processing. Specification: #CompositeTypes: %d, #Segments: %d, #MessageTypes: %d, #Fields: %d", len(s.CompositeTypes), len(s.Segments), len(s.MessageTypes), len(s.Fields))

	// Correct issues in the specification that hinder code generation
	// The long name of XTN.1 is a format, that can't be turned into a valid name
	if maxVersion > 230 {
		s.Fields["XTN.1.CONTENT"].LongName = "Number"
	}
	// ED has two fields named data
	if maxVersion > 220 {
		s.Fields["ED.3.CONTENT"].LongName = "Data Subtype"
	}
	// The XML schema defines a field for the field separator character (ie |)
	// which is implicit in the text format
	delete(s.Fields, "MSH.1.CONTENT")
	// Special case the MSH-2-Encoding characters field, because while it's
	// defined as being of type ST, it actually contains the unescaped delimiters.
	s.Fields["MSH.2.CONTENT"].UpperCaseType = "Delimiters"
	// CNS changed from a datatype to a segment in HL7 2.4, so we rename here to
	// avoid a conflict.
	if maxVersion > 230 {
		s.CompositeTypes["CNS"].Name = "CNS231"
	}

	return s
}

// amendElements fixes schemas that omit fields to indicate that they're deprecated.
// The generator assumes all fields are present and consecutive. Without amendElements, the output
// segments/types have fewer elements and with the wrong number, and the resulting schema is not
// backwards compatible. The absent fields are marked as Optional and Deprecated.
func amendElements(c *ComplexType) {
	elements := c.Elements
	if len(elements) == 0 {
		return
	}

	expectNext := 1
	var finalElements []Element

	for _, e := range elements {
		re := regexp.MustCompile(`\.([0-9]+)`)
		re.MatchString(e.Ref)
		ms := re.FindAllStringSubmatch(e.Ref, -1)
		current, _ := strconv.Atoi(ms[0][1])

		if current != expectNext {
			for j := expectNext; j < current; j++ {
				newE := e
				newE.Ref = e.Ref[:4] + strconv.Itoa(j)
				newE.MinOccurs = "0"
				newE.MaxOccurs = "1"
				newE.Deprecated = true
				log.Printf("Backfilling deprecated field %v (from %v); Added: %+v", c.Name, e.Ref, newE)
				finalElements = append(finalElements, newE)
			}
		}
		expectNext = current + 1
		finalElements = append(finalElements, e)
	}

	c.Elements = finalElements
}

func outputHeader(s *Specification, p *Printer, maxVersion int) {
	p.P("// This file contains the schemas for HL7 messages, segments and values for HL7v2 version %s.", toHl7VersionName(maxVersion))
	p.P("// It has been auto-generated from the HL7v2 specification.")
	p.P("")
	p.P("package hl7")
	p.P("")
	p.P("import \"reflect\"")

	p.P("")
}

func outputSpecification(s *Specification, p *Printer, blockListed map[string]bool) {
	for _, k := range sortedMapKeys(s.CompositeTypes) {
		c := s.CompositeTypes[k]
		if !blockListed[c.Name] {
			outputCompositeType(p, c, s.Fields)
		}
	}

	p.P("")
	for _, k := range sortedMapKeys(s.Segments) {
		c := s.Segments[k]
		if !blockListed[c.Name] {
			outputSegment(p, c, s.Fields)
		}
	}

	p.P("")
	for _, k := range sortedMapKeys(s.MessageTypes) {
		c := s.MessageTypes[k]
		outputMessageType(p, c, s)
	}

	p.P("")
	outputGenericHL7Segment(p)

	p.P("")
	outputMessageFunctions(p, s)

	p.P("")
	p.P("// v2 API")

	names := make([]string, len(s.MessageTypes), len(s.MessageTypes))
	i := 0
	for name := range s.MessageTypes {
		names[i] = name
		i++
	}
	sort.Strings(names)
	fs := map[string]StringSet{}
	groups := make([]*Group, 0)
	for _, name := range names {
		if messageSubtypeName.MatchString(name) {
			continue
		}
		group := buildGroup(s.MessageTypes[name], s)
		groups = append(groups, group)
		outputGroup(p, group)
		buildFollowSets(group, fs)
	}

	p.P("")
	p.P("// Types maps the name of an HL7 segment or message type to the type of the struct that")
	p.P("// represents that segment or message type.")
	p.P("var Types = map[string]reflect.Type{}")
	p.P("func init() {")
	p.In()
	outputSegmentTypeMap(p, s)
	outputMessageTypeMap(p, s)
	p.P("Types[%q] = reflect.TypeOf(%s{})", "GenericHL7Segment", "GenericHL7Segment")
	for _, g := range groups {
		for _, sg := range collectGroups(g) {
			p.P("Types[%q] = reflect.TypeOf(%s{})", sg.NameForType(), sg.NameForType())
		}
	}
	p.Out()
	p.P("}")
	outputFollowSets(p, fs)
}

func sortedMapKeys(m map[string]*ComplexType) []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
