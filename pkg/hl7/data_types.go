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

package hl7

import (
	"bytes"
	"errors"
	"reflect"
	"strconv"
	"time"

	"golang.org/x/text/encoding/unicode"
)

// ST represents a HL7 ST value (String).
type ST string

var _ Primitive = (*ST)(nil)

// NewST returns a new ST.
func NewST(st ST) *ST {
	return &st
}

// Marshal marshals the ST value.
func (st *ST) Marshal(_ *Context) ([]byte, error) {
	return []byte(string(*st)), nil
}

// Unmarshal unmarshals the ST value.
func (st *ST) Unmarshal(field []byte, c *Context) error {
	parsed, err := parseText(field, c)
	if err != nil {
		*st = ""
		return err
	}
	*st = ST(parsed)
	return nil
}

// SanitizeString sanitizes the provided string value.
// SanitizeString returns an empty string if the value is a HL7 NULL e.g. '""', else it is a no-op.
func SanitizeString(s string) string {
	if isHL7Null([]byte(s)) {
		return ""
	}
	return s
}

// SanitizedString sanitizes the ST string value.
func (st *ST) SanitizedString() string {
	return SanitizeString(st.String())
}

func (st *ST) String() string {
	if st == nil {
		return ""
	}

	return string(*st)
}

// Empty returns whether ST is nil or empty.
func (st *ST) Empty() bool {
	return st == nil || *st == ""
}

// ID represents a HL7 ID value (Coded value for HL7 tables).
type ID string

var _ Primitive = (*ID)(nil)

// NewID returns a new ID.
func NewID(id ID) *ID {
	return &id
}

// Marshal marshals the ID value.
func (id *ID) Marshal(_ *Context) ([]byte, error) {
	return []byte(string(*id)), nil
}

// Unmarshal unmarshals the ID value.
func (id *ID) Unmarshal(field []byte, _ *Context) error {
	*id = ID(field)
	return nil
}

// SanitizedString sanitizes the ID string value.
func (id *ID) SanitizedString() string {
	return SanitizeString(id.String())
}

func (id *ID) String() string {
	if id == nil {
		return ""
	}

	return string(*id)
}

// Empty returns whether ID is nil or empty.
func (id *ID) Empty() bool {
	return id == nil || *id == ""
}

// SI represents a HL7 SI value (Sequence ID).
type SI struct {
	Value uint64
	Valid bool // "false" if the SI is an HL7 null.
}

var _ Primitive = (*SI)(nil)

// NewSI returns a new SI.
func NewSI(value uint64) *SI {
	return &SI{value, true}
}

// Marshal marshals the SI value.
func (si *SI) Marshal(_ *Context) ([]byte, error) {
	return []byte(strconv.FormatUint(si.Value, 10)), nil
}

// Unmarshal unmarshals the SI value.
func (si *SI) Unmarshal(field []byte, _ *Context) error {
	si.Valid = !isHL7Null(field)
	if !si.Valid {
		return nil
	}
	i, err := strconv.ParseInt(string(field), 10, 64)
	if err != nil {
		err = ErrBadValue
	} else if i < 0 {
		err = errors.New("SI must be non-negative")
	} else {
		si.Value = uint64(i)
	}
	return err
}

// NM represents a HL7 NM value (Numeric).
type NM struct {
	Value float64
	Valid bool // "false" if the NM is an HL7 null.
}

var _ Primitive = (*NM)(nil)

// NewNM returns a new NM.
func NewNM(value float64) *NM {
	return &NM{Value: value, Valid: true}
}

// Marshal marshals the NM value.
func (nm *NM) Marshal(_ *Context) ([]byte, error) {
	return []byte(strconv.FormatFloat(nm.Value, 'f', -1, 64)), nil
}

// Unmarshal unmarshals the NM value.
func (nm *NM) Unmarshal(field []byte, _ *Context) error {
	nm.Valid = !isHL7Null(field)
	if !nm.Valid {
		return nil
	}
	i, err := strconv.ParseFloat(string(field), 64)
	if err != nil {
		return ErrBadValue
	}
	nm.Value = i
	return nil
}

// IS represents a HL7 IS value (Coded value for user-defined tables).
type IS string

var _ Primitive = (*IS)(nil)

// NewIS returns a new IS.
func NewIS(is IS) *IS {
	return &is
}

// Marshal marshals the IS value.
func (is *IS) Marshal(_ *Context) ([]byte, error) {
	return []byte(string(*is)), nil
}

// Unmarshal unmarshals the IS value.
func (is *IS) Unmarshal(field []byte, _ *Context) error {
	*is = IS(field)
	return nil
}

// SanitizedString sanitizes the IS string value.
func (is *IS) SanitizedString() string {
	return SanitizeString(is.String())
}

func (is *IS) String() string {
	if is == nil {
		return ""
	}

	return string(*is)
}

// Empty returns whether IS is nil or empty.
func (is *IS) Empty() bool {
	return is == nil || *is == ""
}

// DT represents a HL7 DT value (Date).
type DT string

var _ Primitive = (*DT)(nil)

// NewDT returns a new DT.
func NewDT(dt DT) *DT {
	return &dt
}

// Marshal marshals the DT value.
func (dt *DT) Marshal(_ *Context) ([]byte, error) {
	return []byte(string(*dt)), nil
}

// Unmarshal unmarshals the DT value.
func (dt *DT) Unmarshal(field []byte, _ *Context) error {
	*dt = DT(field)
	return nil
}

// TM represents a HL7 TM value (Time).
type TM string

var _ Primitive = (*TM)(nil)

// NewTM returns a new TM.
func NewTM(tm TM) *TM {
	return &tm
}

// Marshal marshals the TM value.
func (tm *TM) Marshal(_ *Context) ([]byte, error) {
	return []byte(string(*tm)), nil
}

// Unmarshal unmarshals the TM value.
func (tm *TM) Unmarshal(field []byte, _ *Context) error {
	*tm = TM(field)
	return nil
}

// TSPrecision represents the precision of a time stamp (TS) value.
type TSPrecision int

// The various TSPrecision values.
const (
	YearPrecision TSPrecision = iota
	MonthPrecision
	DayPrecision
	HourPrecision
	MinutePrecision
	SecondPrecision
	TenthSecondPrecision
	HundredthSecondPrecision
	ThousandthSecondPrecision
	TenThousandthSecondPrecision
)

func (p TSPrecision) formatString() string {
	return [...]string{
		"2006",
		"200601",
		"20060102",
		"2006010215",
		"200601021504",
		"20060102150405",
		"20060102150405.0",
		"20060102150405.00",
		"20060102150405.000",
		"20060102150405.0000",
	}[p]
}

// TS represents a HL7 TS value (Timestamp).
type TS struct {
	IsHL7Null     bool
	Time          time.Time
	Precision     TSPrecision
	ElideTimezone bool
}

var _ Primitive = (*TS)(nil)

// Marshal marshals a TS value.
func (ts *TS) Marshal(c *Context) ([]byte, error) {
	time := ts.Time.In(c.TimezoneLoc)
	format := ts.Precision.formatString()
	if c.IncludeTimezone && !ts.ElideTimezone {
		format += "-0700"
	}
	return []byte(time.Format(format)), nil
}

// Unmarshal unmarshals a TS value, as described in 2.8.42 of the HL7 2.3
// specification. Values have the following format:
//     YYYY[MM[DD[HHMM[SS[.S[S[S[S]]]]]]]][+/-ZZZZ]^<degree of precision>
// Paraphrasing the comments on precision: In the current and future
// versions of HL7, the precision is indicated by limiting the number
// of digits used, unless the optional second component is present. In
// prior versions of HL7, an optional second component indicates the
// degree of precision of the time stamp. This optional second
// component is retained only for purposes of backward compatibility.
func (ts *TS) Unmarshal(field []byte, c *Context) error {
	// Because TS is actually a composite type, we need to use the appropriate delimiter, based
	// on the nesting level in the context.

	ts.IsHL7Null = isHL7Null(field)
	if ts.IsHL7Null {
		return nil
	}
	components := c.Delimiters.splitComponents(Token{Value: field}, c.Nesting)
	var tsWithoutTz []byte
	tzIndex := bytes.LastIndexAny(components[0].Value, "+-")
	if tzIndex > 0 {
		tsWithoutTz = components[0].Value[0:tzIndex]
	} else {
		tsWithoutTz = components[0].Value
	}
	var format string
	switch len(tsWithoutTz) {
	case 4:
		format = "2006"
		ts.Precision = YearPrecision
	case 6:
		format = "200601"
		ts.Precision = MonthPrecision
	case 8:
		format = "20060102"
		ts.Precision = DayPrecision
	case 10:
		format = "2006010215"
		ts.Precision = HourPrecision
	case 12:
		format = "200601021504"
		ts.Precision = MinutePrecision
	case 14:
		format = "20060102150405"
		ts.Precision = SecondPrecision
	case 16:
		format = "20060102150405.0"
		ts.Precision = TenthSecondPrecision
	case 17:
		format = "20060102150405.00"
		ts.Precision = HundredthSecondPrecision
	case 18:
		format = "20060102150405.000"
		ts.Precision = ThousandthSecondPrecision
	case 19:
		format = "20060102150405.0000"
		ts.Precision = TenThousandthSecondPrecision
	default:
		return errors.New("bad TS value: invalid length")
	}
	if len(components) > 1 {
		switch string(components[1].Value) {
		case "Y":
			ts.Precision = YearPrecision
		case "L":
			ts.Precision = MonthPrecision
		case "D":
			ts.Precision = DayPrecision
		case "H":
			ts.Precision = HourPrecision
		case "M":
			ts.Precision = MinutePrecision
		case "S":
			ts.Precision = SecondPrecision
		default:
			return errors.New("bad TS value: unknown precision")
		}
	}
	var err error
	if tzIndex < 0 && ts.Precision > DayPrecision {
		ts.Time, err = time.ParseInLocation(format, string(tsWithoutTz), c.TimezoneLoc)
		return err
	}
	if tzIndex < 0 && ts.Precision <= DayPrecision {
		ts.Time, err = time.Parse(format, string(tsWithoutTz))
		return err
	}
	// The -0700 is only here to show time.parse() how timezones are
	// represented, the 7h value itself is unused.
	// See: https://golang.org/pkg/time/#Parse
	ts.Time, err = time.Parse(format+"-0700", string(components[0].Value))
	return err
}

// TN represents a HL7 TN value (Telephone number).
type TN string

var _ Primitive = (*TN)(nil)

// NewTN returns a new TN.
func NewTN(tn TN) *TN {
	return &tn
}

// Marshal marshals the TN value.
func (tn *TN) Marshal(_ *Context) ([]byte, error) {
	return []byte(string(*tn)), nil
}

// Unmarshal unmarshals the TN value.
func (tn *TN) Unmarshal(field []byte, _ *Context) error {
	*tn = TN(field)
	return nil
}

// FT represents a HL7 FT value (Formatted text).
type FT string

var _ Primitive = (*FT)(nil)

// NewFT returns a new FT.
func NewFT(ft FT) *FT {
	return &ft
}

// Marshal marshals the FT value.
func (ft *FT) Marshal(c *Context) ([]byte, error) {
	return marshalText([]byte(*ft), c), nil
}

// Unmarshal unmarshals the FT value.
func (ft *FT) Unmarshal(field []byte, c *Context) error {
	parsed, err := parseText(field, c)
	if err != nil {
		*ft = ""
		return err
	}
	*ft = FT(parsed)
	return nil
}

// TX represents a HL7 TX value (Text data).
type TX string

var _ Primitive = (*TX)(nil)

// NewTX returns a new TX.
func NewTX(tx TX) *TX {
	return &tx
}

// Marshal marshals the TX value.
func (tx *TX) Marshal(_ *Context) ([]byte, error) {
	return []byte(string(*tx)), nil
}

// Unmarshal unmarshals the TX value.
func (tx *TX) Unmarshal(field []byte, _ *Context) error {
	*tx = TX(field)
	return nil
}

// SanitizedString sanitizes the HD string value.
func (hd *HD) SanitizedString() string {
	return SanitizeString(hd.String())
}

// HD (Hierarchic designator) helper functions.
func (hd *HD) String() string {
	if hd == nil {
		return ""
	}

	c := &Context{
		Decoder:     unicode.UTF8.NewDecoder(),
		Delimiters:  DefaultDelimiters,
		Nesting:     0,
		TimezoneLoc: Location,
	}
	b, err := marshalCompositeValue(reflect.ValueOf(*hd), c)
	if err != nil {
		return "unable to marshal HD"
	}

	return string(b)
}

// CM represents a HL7 CM value (Composite).
// CM was deprecated in HL7 2.4, and most instances have been replaced, so
// we don't provide a working implementation here.
type CM []byte

var _ Primitive = (*CM)(nil)

// NewCM returns a new CM.
func NewCM(cm CM) *CM {
	return &cm
}

// Marshal marshals the CM value.
func (cm *CM) Marshal(_ *Context) ([]byte, error) {
	return []byte(*cm), nil
}

// Unmarshal unmarshals the CM value.
func (cm *CM) Unmarshal(field []byte, _ *Context) error {
	*cm = CM(field)
	return nil
}

// GTS represents a HL7 GTS value (General Timing Specification).
type GTS string

var _ Primitive = (*GTS)(nil)

// Marshal a GTS value.
func (gts *GTS) Marshal(_ *Context) ([]byte, error) {
	return []byte(string(*gts)), nil
}

// Unmarshal a GTS value.
func (gts *GTS) Unmarshal(field []byte, _ *Context) error {
	*gts = GTS(field)
	return nil
}

// NUL represents a HL7 NUL value (Null).
type NUL string

var _ Primitive = (*NUL)(nil)

// Marshal a NUL value.
func (nul *NUL) Marshal(_ *Context) ([]byte, error) {
	return []byte(string(*nul)), nil
}

// Unmarshal a NUL value.
func (nul *NUL) Unmarshal(field []byte, _ *Context) error {
	*nul = NUL(field)
	return nil
}

// Any represents any potential HL7 field type.
type Any []byte

var _ Primitive = (*Any)(nil)

// NewAny returns a new Any.
func NewAny(any Any) *Any {
	return &any
}

// Marshal marshals Any value.
func (any *Any) Marshal(_ *Context) ([]byte, error) {
	return *any, nil
}

// Unmarshal unmarshals Any value.
func (any *Any) Unmarshal(field []byte, _ *Context) error {
	*any = Any(field)
	return nil
}
