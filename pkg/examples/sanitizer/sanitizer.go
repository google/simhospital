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

// Package sanitizer illustrates how the Rewrite capabilities of the HL7 library can be
// utilised to deal with tweaks of HL7v2.
package sanitizer

import (
	"bytes"
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts/cmpopts"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/google/simhospital/pkg/hl7"
	"github.com/google/simhospital/pkg/hl7ids"
	"github.com/google/simhospital/pkg/logging"
	"github.com/google/simhospital/pkg/monitoring"
)

var (
	log = logging.ForCallerPackage()

	postcodeRegexp        = regexp.MustCompile(`^[A-Za-z]{1,2}[0-9][0-9A-Za-z]?\s?[0-9][A-Za-z]{2}$`)
	halfPostcodeRegexp    = regexp.MustCompile(`^[A-Za-z]{1,2}[0-9][0-9A-Za-z]?$`)
	extractPostcodeRegexp = regexp.MustCompile(`([A-Za-z]{1,2}[0-9][0-9A-Za-z]?\s?[0-9][A-Za-z]{2})`)

	// Struct representing a default empty ZCM segment. All fields from the ZCM segment that need to
	// be populated must be present here.
	// ZCM segments are internal segments. If they cause an error, such error should be considered an
	// error of the sanitization stage (HL7_SANITIZATION_FAILURE).
	// This differs from errors with other segments that signal invalid input messages, e.g. errors
	// with PID segments.
	defaultZCM = &hl7.ZCM{
		SetIDZCM:                       hl7.NewSI(1),
		OrderDiscipline:                hl7.NewST(""),
		PatientPrimaryFacilityIDNumber: hl7.NewST(""),
		CustomMessageType:              hl7.NewST(""),
	}

	messageSanitizerCounters struct {
		HL7 struct {
			UnrecognizedDeathIndicatorValue prometheus.Counter `help:"Number of messages that use an unrecognized death indicator value"`
		}
		MessageSanitizer struct {
			RewriteWarningTotal *prometheus.CounterVec `help:"Number of situations we consider unusual behaviour but not necessarily erroneous or alerting worthy" labels:"rewrite_name,reason"`
		}
	}
)

// Error occurs when there's an error in sanitizing a message.
type Error struct {
	E    error
	Name string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error %s: %v", e.Name, e.E)
}

// MessageSanitizerI is an interface for a message sanitizer.
type MessageSanitizerI interface {
	SanitizeMessage(context.Context, Message) (*hl7.Message, error)
}

// MessageSanitizer is a message sanitizer.
// The fields represent configuration that might be useful in a real setting.
type MessageSanitizer struct {
	// The list of Universal Service Identifiers that define a message as a Vital Signs message.
	VitalSignsUSIs []string
	// OBXRemovalStatus is a list of observation statuses that, if found, the OBX segment will be removed.
	OBXRemovalStatus []string
	// UnescapeOBXValues determines whether to unescape obx value field.
	UnescapeOBXValues bool
	// Whether to rewrite patient identifiers
	RewritePatientIdentifiers  bool
	AddMRN                     bool
	RewriteIdentifierLocations []string
	// Whether to squash OBXs
	// Generally this will depend on parameters such as the sendingApplication, sendingFacility, messageType...
	ShouldSquashOBXs func(message *hl7.Message) bool
	// Map from old sending facilities to the new one
	OldSendingFacilities map[string]string
	// Whether to strip zeros from MRN fields
	RemoveLeadingZerosFromMRN bool
	// Map of death indicators to the correct one
	DeathIndicatorNormalizer map[string]string
}

// CommonNormalizedDeathIndicator contains the most common mappings for death indicators.
var CommonNormalizedDeathIndicator = map[string]string{
	"y":        "Y",
	"yes":      "Y",
	"Deceased": "Y",
	"deceased": "Y",
	"n":        "N",
	"no":       "N",
	"Y":        "Y",
	"YES":      "Y",
	"DECEASED": "Y",
	"N":        "N",
	"NO":       "N",
}

func init() {
	monitoring.CreateAndRegisterMetricsFromStruct(&messageSanitizerCounters)
}

// Message is a message used as input for the Sanitizer.
type Message struct{ Message string }

// SanitizeMessage removes known partner-specific nuances in the HL7 format. For instance,
// even if there's a reserved field to send a patient's zip code, some partners might use a different
// field. The sanitization process modifies the underlying HL7 message and unifies it for all
// partners so that downstream services don't have to deal with such nuances.
// The sanitization step also takes care of known problems within a partner's messages,
// for instance, it can fix known parsing errors without failing the parsing stage.
// The MSH segment in the message needs to be present and parsable, otherwise an error is returned.
// After the sanitization, if there were no errors, the returned message will have the following
// properties:
// * All messages will have valid and parsable MSH, PID and ZCM segments.
// * Vital Signs messages will also have parsable OBR and OBX segments, if present.
// If an error is returned, it will be a SanitizerError.
// An error of type MESSAGE_SANITIZATION_FAILURE means that something went wrong while sanitizing;
// any other error indicates a problem with the message.
func (s MessageSanitizer) SanitizeMessage(ctx context.Context, m *Message) (*hl7.Message, error) {
	zcm, err := hl7.MarshalSegment(defaultZCM, hl7.DefaultContextWithoutLocation)
	if err != nil {
		return nil, &Error{
			E:    errors.Wrap(err, "could not marshal ZCM segment"),
			Name: "MESSAGE_SANITIZATION_FAILURE",
		}
	}
	m.Message = fmt.Sprintf("%s\r%s", m.Message, string(zcm))

	if err = doBasicValidation(m.Message); err != nil {
		return nil, err
	}

	hl7Msg, _ := hl7.ParseMessage([]byte(m.Message))
	var rewrites []hl7.Rewrite

	if s.RewritePatientIdentifiers {
		rewrites = append(rewrites, rewritePatientIdentifiers(s.AddMRN, s.RewriteIdentifierLocations...))
	}

	rewrites = append(rewrites, rewriteSendingFacility(s.OldSendingFacilities))

	rewriteZipcode, err := s.rewriteZipCodeFromAddress(ctx, hl7Msg)
	if err != nil {
		return nil, err
	}
	if rewriteZipcode != nil {
		rewrites = append(rewrites, rewriteZipcode)
	}

	rewrites = append(rewrites, rewriteDuplicatedPostcode)

	// Remove zeros from the front of MRN numbers. This requires the MRN keyword to be present, so it
	// needs to be done after rewritePatientIdentifiers.
	if s.RemoveLeadingZerosFromMRN {
		rewrites = append(rewrites, rewriteRemoveLeadingZerosFromMRN)
	}

	if s.ShouldSquashOBXs != nil && s.ShouldSquashOBXs(hl7Msg) {
		if squashOBXes := s.squashOBXes(ctx, hl7Msg); squashOBXes != nil {
			log.Debug("Including squashOBXes rewrite")
			rewrites = append(rewrites, squashOBXes)
		}
	}

	rewrites = append(rewrites, rewritePIDDateTimeOfBirth)
	rewrites = append(rewrites, s.rewriteUnitsCEElements())

	rewritePatientDeathInd := s.rewritePIDPatientDeathIndicator(ctx)
	if rewritePatientDeathInd != nil {
		rewrites = append(rewrites, rewritePatientDeathInd)
	}

	rewritePrimaryFac := s.rewritePatientPrimaryFacility(ctx, hl7Msg)
	if rewritePrimaryFac != nil {
		rewrites = append(rewrites, rewritePrimaryFac)
	}

	if s.UnescapeOBXValues {
		rewrites = append(rewrites, rewriteOBXUnescapeValue())
	}

	rHL7Msg, err := doRewrites(rewrites, m)
	if err != nil {
		return nil, err
	}
	// After the rewrites are applied, make sure that the PID segment is present and parsable.
	if _, err = getParsedPID(rHL7Msg); err != nil {
		return nil, err
	}
	// After the rewrites are applied, make sure that the ZCM segment is present and parsable.
	if err := getParsedZCMErr(rHL7Msg); err != nil {
		return nil, err
	}
	return rHL7Msg, nil
}

func doRewrites(rewrites []hl7.Rewrite, m *Message) (*hl7.Message, error) {
	mo := hl7.NewParseMessageOptions()
	mo.Rewrites = &rewrites
	hl7Msg, err := hl7.ParseMessageWithOptions([]byte(m.Message), mo)
	if err != nil {
		return nil, &Error{
			E:    errors.Wrap(err, "Could not apply rewrite"),
			Name: "MESSAGE_SANITIZATION_FAILURE",
		}
	}

	// Force an initial rewrite function in the PD1 segment. This is a horrible hack, see
	// rewritePatientPrimaryFacility.
	hl7Msg.PD1()

	return hl7Msg, nil
}

// rewriteZipCodeFromAddress returns a function that rewrites the PatientAddress.ZipOrPostalCode
// field if the postcode has been sent in a different field of the PatientAddress.
// This method allows for postcodes to be included as part of larger field, for instance it
// successfully extracts the postcode "N1C 4AB" from "LONDON N1C 4AB".
// The postcode is only rewritten if the PatientAddress.ZipOrPostalCode is empty, in order to avoid
// losing information, even if such field might be malformed.
// Only postcodes with a valid UK format are rewritten into ZipOrPostalCode.
// The original field the postcode came from is not modified.
// This function returns nil if no replacements are needed.
func (s MessageSanitizer) rewriteZipCodeFromAddress(ctx context.Context, hl7Msg *hl7.Message) (hl7.Rewrite, error) {
	pid, err := getParsedPID(hl7Msg)
	if err != nil {
		return nil, err
	}
	pids, err := hl7Msg.AllPID()
	if err != nil {
		return nil, &Error{
			E:    errors.Wrap(err, "Could not get retrieve all PID segments"),
			Name: "PID_RETRIEVAL_FAILURE",
		}
	}
	if len(pids) > 1 {
		log.Warning("Multiple PID segments in message. Things might not work.")
	}

	// A map of address types to replace the corresponding zip code. The keys represents the
	// indexes within the list of patient addresses that need replacement.
	addressTypes := map[int]string{}
	for i := range pid.PatientAddress {
		address := pid.PatientAddress[i]
		if address.ZipOrPostalCode.String() != "" {
			if extractPostcode(address.ZipOrPostalCode.String()) == "" {
				log.WithContext(ctx).Warning("Not rewriting malformed postcode")
			}
			continue
		}

		// Try to extract the postcode from other fields.
		// Give preference to Address Type; some messages communicate
		// the postcode in this field instead.
		addressTypes[i] = extractPostcode(address.AddressType.String())
		if addressTypes[i] != "" {
			continue
		}
		if address.StreetAddress != nil {
			addressTypes[i] = extractPostcode(address.StreetAddress.StreetOrMailingAddress.String())
			if addressTypes[i] != "" {
				continue
			}
		}
		addressTypes[i] = extractPostcode(address.OtherDesignation.String())
		if addressTypes[i] != "" {
			continue
		}
		addressTypes[i] = extractPostcode(address.City.String())
		if addressTypes[i] != "" {
			continue
		}
		addressTypes[i] = extractPostcode(address.StateOrProvince.String())
		if addressTypes[i] != "" {
			continue
		}
		addressTypes[i] = extractPostcode(address.Country.String())
	}
	if len(addressTypes) == 0 {
		return nil, nil
	}

	// A map of offsets already rewritten, to avoid incrementing our indexes twice for the same
	// value (the Rewrite function processes some values twice).
	addressOffsetSeen := map[int]bool{}
	addressIndex := -1

	return func(t hl7.Token) *hl7.RewriteResult {
		if t.Location == "PID-2-Patient ID" {
			// We need to reset the indexes every time the PID is parsed, otherwise the
			// message would be rewritten only for the first call to message.PID().
			addressIndex = -1
			addressOffsetSeen = map[int]bool{}
		}
		if t.Location == "PID-11-Patient Address" && !addressOffsetSeen[t.Offset] {
			addressOffsetSeen[t.Offset] = true
			addressIndex++
		}
		if t.Location == "PID-11-Patient Address/XAD-5-Zip Or Postal Code" {
			if addressTypes[addressIndex] != "" {
				return hl7.RewriteResultReplaceValue([]byte(addressTypes[addressIndex]))
			}
		}
		return hl7.RewriteResultNoop()
	}, nil
}

// extractPostcode extracts a valid postcode from the given string.
// It returns the empty string if no valid postcode can be found in the field.
// If the field contains more than one item that looks like a postcode, it returns the first one.
func extractPostcode(s string) string {
	r := extractPostcodeRegexp.FindStringSubmatch(s)
	if len(r) > 0 {
		return r[0]
	}
	return ""
}

// rewriteDuplicatedPostcode rewrites postcode fields that are of the form "POSTCODEPOSTCODE" or
// "POSTCODE POSTCODE", i.e., two concatenated equal postcodes with or without a space in between.
// For instance, "N1C 4ABN1C 4AB" is rewritten to "N1C 4AB".
// It also rewrites postcodes of the form "N1C N1C" (the postcode is only the first half of a
// full postcode).
// If the postcode field does not have any of those formats, or the postcode does not have the
// format of a valid UK postcode, this is a no-op.
func rewriteDuplicatedPostcode(t hl7.Token) *hl7.RewriteResult {
	if t.Location == "PID-11-Patient Address/XAD-5-Zip Or Postal Code" {
		value := string(t.Value)
		if len(value) >= 2 {
			mid := len(value) / 2
			firstHalf := value[0:mid]
			secondHalf := value[mid:]

			if (firstHalf == secondHalf || secondHalf == fmt.Sprintf(" %s", firstHalf)) &&
				(postcodeRegexp.MatchString(firstHalf) || halfPostcodeRegexp.MatchString(firstHalf)) {
				return hl7.RewriteResultReplaceValue([]byte(firstHalf))
			}
		}
	}
	return hl7.RewriteResultNoop()
}

// squashOBXes tries to join the values of all the OBXes into a single OBX.
// The use case is for Radiology reports where we receive them as one OBX per line of the report.
// The squash will only be performed if all OBXes are identical except for ObservationValue and SetID and
// all ObservationValues are single-value.
func (s MessageSanitizer) squashOBXes(ctx context.Context, m *hl7.Message) hl7.Rewrite {
	// Ignore parsing errors in the OBX segments, since those should be handled in a different
	// place, and some could even be solved by the sanitizer.
	obxes, _ := m.AllOBX()
	logLocal := log.WithContext(ctx).WithField("num_obxes", len(obxes))
	if len(obxes) <= 1 {
		logLocal.Debug("Not squashing OBXes. There's one or less OBXes in this message.")
		return nil
	}
	logLocal.Debug("Trying to squash OBXes")

	// Check if all the OBXes share metadata
	var firstOBX = obxes[0]
	var values []hl7.Any
	for _, obx := range obxes {
		switch len(obx.ObservationValue) {
		case 0:
			values = append(values, hl7.Any(""))
		case 1:
			values = append(values, obx.ObservationValue[0])
		default:
			logLocal.WithFields(
				logrus.Fields{
					"set_id":     obx.SetIDOBX.Value,
					"num_values": len(obx.ObservationValue),
				}).
				Warning("Not squashing OBXes. OBX has multiple values.")
			return nil
		}
		// TODO: bail if there's any notes.
		if !compareOBXesForRadiologySquash(firstOBX, obx) {
			logLocal.WithField("set_id", obx.SetIDOBX.Value).
				Warning("Not squashing OBXes. OBX is not equal to the first OBX ignoring SetID and Value")
			return nil
		}
	}
	// All OBXes are identical except for the SetID and Value. Proceed with re-write.
	logLocal.WithField("num_values_to_join", len(values)).Debug("Proceeding with OBX squashing.")
	byteValues := make([][]byte, len(values))
	for i, v := range values {
		byteValues[i] = []byte(v)
	}
	newValue := bytes.Join([][]byte(byteValues), []byte("\n"))

	firstOBX.ObservationValue[0] = newValue
	newOBX, err := hl7.MarshalSegment(firstOBX, m.Context)
	if err != nil {
		logLocal.WithError(err).
			Error("Not squashing OBXes. Failed to marshal squashed OBX.")
		return nil
	}
	firstSetID, err := firstOBX.SetIDOBX.Marshal(m.Context)
	if err != nil {
		logLocal.WithError(err).
			Error("Not squashing OBXes. Failed to get SetId for squashed OBX.")
		return nil
	}
	firstSetID = bytes.Join([][]byte{[]byte("OBX|"), firstSetID, {m.Context.Delimiters.Field}}, []byte{})
	return func(t hl7.Token) *hl7.RewriteResult {
		if t.Location == "OBX" {
			// Set the first OBX by SetId to the new value.
			if bytes.Index(t.Value, firstSetID) == 0 {
				return hl7.RewriteResultReplaceValue(newOBX)
			}
			// Delete all other OBXes.
			return hl7.RewriteResultDeleteToken()
		}
		return hl7.RewriteResultNoop()
	}
}

// compareOBXesForRadiologySquash compares two OBXes ignoring SetID and ObservationValue.
func compareOBXesForRadiologySquash(o1, o2 *hl7.OBX) bool {
	return cmp.Equal(o1, o2, cmpopts.IgnoreFields(hl7.OBX{}, "SetIDOBX", "ObservationValue"))
}

// rewritePatientPrimaryFacility removes the IDNumber of the XON field in PD1.3 (Patient Primary
// Facility-ID Number, type NM) and puts it in ZCM-3-Patient Primary Facility ID Number (type ST).
// The ID Number is defined as an NM field in HL7 but in reality it does not always contain a
// number, which makes our parsing fail.
func (s MessageSanitizer) rewritePatientPrimaryFacility(ctx context.Context, hl7Msg *hl7.Message) hl7.Rewrite {
	pd1, _ := hl7Msg.PD1()
	if pd1 == nil || len(pd1.PatientPrimaryFacility) == 0 {
		return nil
	}

	if len(pd1.PatientPrimaryFacility) > 1 {
		log.WithContext(ctx).
			Warning("Multiple PD1.PatientPrimaryFacility in message. Things might not work.")
	}
	return s.rewriteXONIDNumber(ctx, "PD1-3-Patient Primary Facility", "ZCM-3-Patient Primary Facility ID Number")
}

// rewriteXONIDNumber removes the IDNumber of an XON field (type NM) and puts it in the destLocation
// field of the message.
func (s MessageSanitizer) rewriteXONIDNumber(ctx context.Context, xonLocation string, destLocation string) hl7.Rewrite {
	// id will cache the value from ID Number so that it can be written into destLocation on a
	// second pass.
	// Note destLocation should be located after xonLocation in the message, so that by the time the
	// rewrite function gets to it, "id" has already been populated.
	id := ""
	return func(t hl7.Token) *hl7.RewriteResult {
		if t.Location == xonLocation+"/XON-3-ID Number" {
			// Populated only once. In repeated fields, this will get the value of the first item.
			if id == "" {
				id = string(t.Value)
			}
			return hl7.RewriteResultDeleteToken()
		}
		if t.Location == destLocation {
			return hl7.RewriteResultReplaceValue([]byte(id))
		}
		return hl7.RewriteResultNoop()
	}
}

// rewriteSendingFacility replaces old sending facilities that old messages have, with new sending
// facilities.
func rewriteSendingFacility(m map[string]string) hl7.Rewrite {
	// TODO: This, and I guess other, rewrites would be easier to do
	// after the message has been parsed. However - each segment gets parsed
	// multiple times within the code at the moment, so we have to use a rewrite
	// instead. Once the message type structs ensure that segments are only parsed
	// once, we can simplify these.
	return func(t hl7.Token) *hl7.RewriteResult {
		if t.Location == "MSH-3-Sending Facility" {
			if v, ok := m[string(t.Value)]; ok {
				return hl7.RewriteResultReplaceValue([]byte(v))
			}
		}
		return hl7.RewriteResultNoop()
	}
}

const dateLength = len("20060102")

// rewritePIDDateTimeOfBirth ensures that the PID date of birth field is no
// more precise than a day. If the value was more precise, it would be
// converted from the message timezone to UTC, potentially changing the date
// and rendering it incorrect should the timezone adjustment cross midnight.
// TODO: Work out a more robust way of handling this that allows
// a precise time of birth to be stored, potentially by storing the date in
// the original timezone along with all UTC-converted timestamps.
func rewritePIDDateTimeOfBirth(t hl7.Token) *hl7.RewriteResult {
	if t.Location == "PID-7-Date/Time Of Birth" && len(t.Value) > dateLength {
		return hl7.RewriteResultReplaceValue(t.Value[0:dateLength])
	}
	return hl7.RewriteResultNoop()
}

// rewritePIDPatientDeathIndicator ensures that the PID death indicator field
// is always set to one of the allowed values in the map.
func (s MessageSanitizer) rewritePIDPatientDeathIndicator(ctx context.Context) hl7.Rewrite {
	counterIncremented := false
	return func(t hl7.Token) *hl7.RewriteResult {
		if t.Location != "PID-30-Patient Death Indicator" || len(t.Value) == 0 || s.DeathIndicatorNormalizer == nil || len(s.DeathIndicatorNormalizer) == 0 {
			return hl7.RewriteResultNoop()
		}
		if hl7.IsHL7Null(t.Value) {
			return hl7.RewriteResultReplaceValue(hl7.Null)
		}

		normalized, ok := s.DeathIndicatorNormalizer[string(t.Value)]
		if !ok {
			// Unrecognized value. Increment the UnrecognizedDeathIndicatorValue
			// counter, and ignore the value by clearing it (not to be
			// confused with setting it to HL7 Null).
			if !counterIncremented {
				messageSanitizerCounters.HL7.UnrecognizedDeathIndicatorValue.Inc()
				counterIncremented = true
			}
			return hl7.RewriteResultDeleteToken()
		}
		return hl7.RewriteResultReplaceValue([]byte(normalized))
	}
}

// rewriteUnitsCEElements will fill the CE-2-Text inside units with an OBX segment if it is empty,
// as this is where we persist units from. We populate this with CE-1-Identifier.
// This only works in cases where we will visit the CE-2-Text location, e.g. |mmHg^^|.
func (s MessageSanitizer) rewriteUnitsCEElements() hl7.Rewrite {
	// id will cache the value from CE-1-Identifier so that it can be written into CE-2-Text on a second pass.
	var id []byte
	return func(t hl7.Token) *hl7.RewriteResult {
		if t.Location == "OBX-6-Units" && !strings.Contains(string(t.Value), "^") {
			// Add the placeholder for CE-2-Text, otherwise we won't visit it.
			return hl7.RewriteResultReplaceValue([]byte(string(t.Value) + "^"))
		}
		if t.Location == "OBX-6-Units/CE-1-Identifier" {
			id = t.Value
		} else if t.Location == "OBX-6-Units/CE-2-Text" && len(t.Value) == 0 {
			return hl7.RewriteResultReplaceValue(id)
		}
		return hl7.RewriteResultNoop()
	}
}

func contains(target string, in ...string) bool {
	for _, i := range in {
		if i == target {
			return true
		}
	}
	return false
}

// rewritePatientIdentifiers returns a function that rewrites the identifiers for the given
// locations (typically PID-2 and/or PID-3 fields), adding the words
// MRN or NHS to the IdentifierTypeCode field.
// If the value is a valid NHS number, it adds the keyword for NHS numbers.
// If the value is present and it's not a valid NHS number, and rewriteMRN is true, it adds the
// keyword for MRNs.
// This is useful for Point of Care (POC) messages that only contain MRN or NHS numbers but do not
// contain the type of identifier. Without this method, the downstream code might fail to parse
// these as valid MRN or NHS numbers.
// Example of POC messages include NamespaceID == "cobas IT 1000" or "Rapidcomm".
func rewritePatientIdentifiers(rewriteMRN bool, fromLocations ...string) hl7.Rewrite {
	return func(t hl7.Token) *hl7.RewriteResult {

		if !contains(t.Location, fromLocations...) {
			return hl7.RewriteResultNoop()
		}

		val := string(t.Value)
		joinDelimiter := string(hl7.DefaultContextWithoutLocation.Delimiters.Component)

		// If the Location we're at has multiple items, such as multiple identifiers in
		// "PID-3-Patient Identifier List", we get here at different levels: first for the entire
		// field (that will contain multiple items), and then for each individual item.
		// We need to get to the individual items separately in order to rewrite them, so if we're still
		// at the list level, don't rewrite the value yet.
		if hasMultipleItems(t.Value) {
			return hl7.RewriteResultNoop()
		}

		parts := strings.Split(val, string(hl7.DefaultContextWithoutLocation.Delimiters.Component))

		if len(parts) != 1 {
			return hl7.RewriteResultNoop()
		}
		value := parts[0]

		var keyword string
		switch {
		case hl7ids.NHSNumberIsValid(value):
			keyword = "NHSNMBR"
		case value != "" && rewriteMRN:
			keyword = "MRN"
		default:
			return hl7.RewriteResultNoop()
		}

		returnVal := strings.Join([]string{value, joinDelimiter, joinDelimiter, joinDelimiter, joinDelimiter, keyword}, "")
		return hl7.RewriteResultReplaceValue([]byte(returnVal))
	}
}

// rewriteRemoveLeadingZerosFromMRN will rewrite the PID-2, PID-3 and PID-4 fields. It will remove 0's from
// the front of the identifier, if this identifier is an MRN.
func rewriteRemoveLeadingZerosFromMRN(t hl7.Token) *hl7.RewriteResult {
	if t.Location == "PID-2-Patient ID" || t.Location == "PID-3-Patient Identifier List" || t.Location == "PID-4-Alternate Patient ID - PID" {
		val := string(t.Value)

		if val == "" {
			return hl7.RewriteResultNoop()
		}

		// If the Location we're at has multiple items, such as multiple identifiers in
		// "PID-3-Patient Identifier List", we get here at different levels: first for the entire
		// field (that will contain multiple items), and then for each individual item.
		// We need to get to the individual items separately in order to rewrite them, so if we're still
		// at the list level, don't rewrite the value yet.
		if hasMultipleItems(t.Value) {
			return hl7.RewriteResultNoop()
		}

		joinDelimiter := string(hl7.DefaultContextWithoutLocation.Delimiters.Component)
		parts := strings.Split(val, string(hl7.DefaultContextWithoutLocation.Delimiters.Component))

		returnVal := strings.Join([]string{parts[0], joinDelimiter}, "")
		if hasMRNKeyword(parts...) {
			returnVal = strings.Join([]string{strings.TrimLeft(parts[0], "0"), joinDelimiter}, "")
		}

		for _, part := range parts[1:] {
			returnVal = strings.Join([]string{returnVal, part, joinDelimiter}, "")
		}
		return hl7.RewriteResultReplaceValue([]byte(returnVal))
	}
	return hl7.RewriteResultNoop()
}

func hasMRNKeyword(mrns ...string) bool {
	for _, s := range mrns {
		if strings.Contains(s, "MRN") {
			return true
		}
	}
	return false
}

// getParsedPID returns a valid PID segment from the message.
// It returns an Error if the PID segment is invalid or there is no PID segment.
func getParsedPID(hl7Msg *hl7.Message) (*hl7.PID, error) {
	pid, err := hl7Msg.PID()
	if err != nil {
		return nil, &Error{
			E:    errors.Wrap(err, "Could not retrieve PID"),
			Name: "PID_RETRIEVAL_FAILURE",
		}
	}
	if pid == nil {
		return nil, &Error{
			E:    errors.New("No PID segment in message"),
			Name: "NO_PID",
		}
	}
	return pid, nil
}

// getParsedZCMErr returns an Error of type HL7_SANITIZATION_FAILURE if the message doesn't
// have exactly one valid ZCM segment.
func getParsedZCMErr(hl7Msg *hl7.Message) error {
	// Is the ZCM segment parsable?
	_, err := hl7Msg.ZCM()
	if err != nil {
		return &Error{
			E:    errors.Wrap(err, "could not retrieve ZCM"),
			Name: "MESSAGE_SANITIZATION_FAILURE",
		}
	}
	// Is there exactly one ZCM segment?
	allZCM, _ := hl7Msg.AllZCM()
	if len(allZCM) != 1 {
		// If we're here, it could be because (a) there is a bug in our sanitization logic that either did not attach
		// the ZCM or attached too many, or (b) the input message had a ZCM segment already.
		return &Error{
			E:    errors.Errorf("expected 1 ZCM segment after sanitization, got %d", len(allZCM)),
			Name: "MESSAGE_SANITIZATION_FAILURE",
		}
	}
	return nil
}

// doBasicValidation returns an error if the message is not valid for the sanitization stage.
// A message is considered valid for the sanitization stage if:
// * It is parsable, and
// * It contains a MSH segment, and
// * The sending facility, sending application and message type are present.
// * The DateTimeOfMessage is present.
// Note that some of these cases might have already been checked by upstream services before the
// messages get here.
func doBasicValidation(message string) error {
	hl7Msg, err := hl7.ParseMessage([]byte(message))
	if err != nil {
		return &Error{
			E:    errors.Wrap(err, "could not parse message"),
			Name: "HL7_PARSE_FAILURE",
		}
	}
	// The MSH segment is a special segment and it cannot be nil or have parsing errors if ParseMessage succeeded.
	msh, _ := hl7Msg.MSH()
	if msh.SendingFacility == nil || msh.SendingFacility.NamespaceID.String() == "" {
		return &Error{
			E:    errors.New("could not get sending facility in MSH segment"),
			Name: "MISSING_HL7_FIELD",
		}
	}
	if msh.MessageType == nil || msh.MessageType.MessageCode.String() == "" {
		return &Error{
			E:    errors.New("could not get message type in MSH segment"),
			Name: "MISSING_HL7_FIELD",
		}
	}
	if msh.SendingApplication == nil || msh.SendingApplication.NamespaceID.String() == "" {
		return &Error{
			E:    errors.New("could not get sending application in MSH segment"),
			Name: "MISSING_HL7_FIELD",
		}
	}
	if msh.DateTimeOfMessage == nil {
		return &Error{
			E:    errors.New("could not get message time stamp in MSH Segment"),
			Name: "MISSING_HL7_FIELD",
		}
	}
	return nil
}

// hasMultipleItems returns whether the given value is a list of items.
func hasMultipleItems(value []byte) bool {
	return bytes.Contains(value, []byte{hl7.DefaultContextWithoutLocation.Delimiters.Repetition})
}

// rewriteOBXUnescapeValue will unescape and rewrite the OBX.5 field. Standard HL7
// unescaping which is used to unescape ST and FT fields will be used. If an error
// occurs during unescaping, this function will be a no-op.
func rewriteOBXUnescapeValue() hl7.Rewrite {
	/*
		Rewrites are applied when a segment is requested off a parsed message.
		For example, consider the following message `m`:
					MSH|...
					PID|...
					OBR|...
					OBX|1||||one~two~three||||...
					OBX|2||||four/.br/five||||...
					OBX|3||||six/E/seven~eight|||...

		Calling m.OBX() will parse the first OBX segment out of this message. This is when the rewrites are
		applied to this segment. Rewrites are applied top down,
			* Firstly, all the rewrite methods are called on the segment itself (first OBX).
			* Then all the rewrites are called on each field.
					* If a field is repeated (OBX.5 is repeated), rewrites are called on the whole field as well as
						individual sub fields. For example, for an OBX.5 value x~y~z, rewrites are applied in this order
						[x~y~z, x, y, z].

		Calling m.AllOBX() does all this to all the obx segments in the message. We can exploit knowing the fact that
		"rewrites are applied top down to a segment" to make sure that rewrites are not applied twice on a single field.
		This is achieved by noting where the end of the OBX.5 field is and then keeping track of when we have parsed the
		last element so we can reset the end of field index.
	*/
	endOfValueFieldIndex := -1
	return func(t hl7.Token) *hl7.RewriteResult {
		if t.Location == "OBX-5-Observation Value" {

			endOfCurrFieldIndex := t.Offset + len(t.Value)
			if endOfValueFieldIndex == -1 {
				// We are just starting to unescape an entire OBX.5 field.
				// So mark endOfValueFieldIndex to be the end of the OBX.5 index.
				// And quit without doing the unescaping.
				endOfValueFieldIndex = endOfCurrFieldIndex
				return hl7.RewriteResultNoop()
			} else if endOfCurrFieldIndex < endOfValueFieldIndex {
				// We are sanitizing one of the repeated fields which is not the last item.
				// So do nothing and continue with unescaping.
				// This if-else branch doesn't do anything, but keeping it here for the sake of clarity.
			} else if endOfCurrFieldIndex == endOfValueFieldIndex {
				// This item in the repeated array is either the only item in the OBX.5 field
				// or this is the last item. So now set endOfValueFieldIndex to be -1 to mark that the
				//entire OBX.5 field has been unescaped. Then continue with unescaping this item.
				endOfValueFieldIndex = -1
			}

			// TODO: Use delimiters from the message instead of default ones
			unescaped, err := hl7.UnescapeText(t.Value, hl7.DefaultDelimiters, false /* isST=false*/)
			if err != nil {
				// TODO: Increment a metric if we fail unescaping
				return hl7.RewriteResultNoop()
			}
			return hl7.RewriteResultReplaceValue(unescaped)
		}
		return hl7.RewriteResultNoop()
	}
}
