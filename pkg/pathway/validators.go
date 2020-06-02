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
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/google/simhospital/pkg/clock"
	"github.com/google/simhospital/pkg/constants"
	"github.com/google/simhospital/pkg/doctor"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/location"
	"github.com/google/simhospital/pkg/orderprofile"
)

func (d *Delay) valid() error {
	if d == nil {
		return nil
	}
	if d.From.Seconds() > d.To.Seconds() {
		return errors.New("parameter To needs to be greater or equal than From")
	}
	if d.From.Seconds() < 0 {
		return errors.New("both endpoints need to be positive")
	}
	return nil
}

func (u *UsePatient) valid() error {
	if u == nil {
		return nil
	}
	if u.Patient == "" || u.Patient == Current {
		return fmt.Errorf("use_patient requires a set 'patient' different from %s", Current)
	}
	return nil
}

func validLocation(loc string, lm *location.Manager) error {
	if loc == "" {
		return errors.New("location not provided")
	}
	if _, ok := lm.RoomManagers[loc]; !ok {
		locs := locationNames(lm)
		return fmt.Errorf("unknown location %q, supported locations are [%v]", loc, strings.Join(locs, ","))
	}
	return nil
}

func locationNames(lm *location.Manager) []string {
	var locs []string
	for l := range lm.RoomManagers {
		locs = append(locs, l)
	}
	return locs
}

func (p *PreAdmission) valid(lm *location.Manager) error {
	if p == nil {
		return nil
	}
	if err := validLocation(p.Loc, lm); err != nil {
		return errors.Wrap(err, "error validating location in pre-admission")
	}
	if p.ExpectedAdmissionTimeFromNow == nil || *p.ExpectedAdmissionTimeFromNow < time.Duration(0) {
		return errors.New("a pre-admission requires a positive expected_admission_time_from_now")
	}
	return nil
}

func (p *PendingAdmission) valid(lm *location.Manager) error {
	if p == nil {
		return nil
	}
	if err := validLocation(p.Loc, lm); err != nil {
		return errors.Wrap(err, "error validating location in pending admission")
	}
	if p.ExpectedAdmissionTimeFromNow == nil || *p.ExpectedAdmissionTimeFromNow < time.Duration(0) {
		return errors.New("a pending admission requires a positive expected_admission_time_from_now")
	}
	return nil
}

func (a *Admission) valid(lm *location.Manager) error {
	if a == nil {
		return nil
	}
	if err := validLocation(a.Loc, lm); err != nil {
		return errors.Wrap(err, "error validating location in admission")
	}
	return nil
}

func (t *Transfer) valid(lm *location.Manager) error {
	if t == nil {
		return nil
	}

	if err := validLocation(t.Loc, lm); err != nil {
		return errors.Wrap(err, "error validating location in transfer")
	}
	return nil
}

func (t *TransferInError) valid(lm *location.Manager) error {
	if t == nil {
		return nil
	}

	if err := validLocation(t.Loc, lm); err != nil {
		return errors.Wrap(err, "error validating location in transfer_in_error")
	}
	return nil
}

func (p *PendingDischarge) valid() error {
	if p == nil {
		return nil
	}
	if p.ExpectedDischargeTimeFromNow == nil || *p.ExpectedDischargeTimeFromNow < time.Duration(0) {
		return errors.New("a pending discharge requires a positive expected_discharge_time_from_now")
	}
	return nil
}

func (p *PendingTransfer) valid(lm *location.Manager) error {
	if p == nil {
		return nil
	}
	if err := validLocation(p.Loc, lm); err != nil {
		return errors.Wrap(err, "error validating location in pending transfer")
	}
	if p.ExpectedTransferTimeFromNow == nil || *p.ExpectedTransferTimeFromNow < time.Duration(0) {
		return errors.New("a pending transfer requires a positive expected_transfer_time_from_now")
	}
	return nil
}

func (t *TrackDeparture) valid(lm *location.Manager) error {
	if t == nil {
		return nil
	}
	if t.DestinationLoc == "" {
		return errors.New("departure tracking requires a destination_loc")
	}
	if t.Mode != TrackMode && t.Mode != TransitMode && t.Mode != TemporaryMode {
		return fmt.Errorf("departure tracking mode is %q; must be one of [%v, %v, %v]", t.Mode, TrackMode, TransitMode, TemporaryMode)
	}
	if t.DestinationBed != "" && t.Mode != TrackMode && t.Mode != TransitMode {
		return fmt.Errorf("departure tracking destination_bed is %q for mode %q; destination_bed is only supported for modes [%v, %v]", t.DestinationBed, t.Mode, TrackMode, TransitMode)
	}
	if t.Mode != TemporaryMode {
		if err := validLocation(t.DestinationLoc, lm); err != nil {
			return errors.Wrapf(err, "error validating location in departure tracking in mode %s", t.Mode)
		}
	}
	return nil
}

func (t *TrackArrival) valid(lm *location.Manager) error {
	if t == nil {
		return nil
	}
	if t.Loc == "" {
		return errors.New("arrival tracking requires a loc")
	}
	if t.Mode != TrackMode && t.Mode != TransitMode && t.Mode != TemporaryMode {
		return fmt.Errorf("arrival tracking mode is %q; must be one of [%v, %v, %v]", t.Mode, TrackMode, TransitMode, TemporaryMode)
	}
	if t.IsTemporary && t.Mode != TemporaryMode {
		return fmt.Errorf("field is_temporary in departure tracking  event can only be set for mode %q", TemporaryMode)
	}
	if t.Bed != "" && t.Mode != TrackMode && t.Mode != TemporaryMode {
		return fmt.Errorf("departure tracking bed is %q for mode %q; bed is only supported for modes [%v, %v]", t.Bed, t.Mode, TrackMode, TemporaryMode)
	}
	if !t.IsTemporary {
		if err := validLocation(t.Loc, lm); err != nil {
			return errors.Wrapf(err, "error validating location in arrival tracking in mode %s", t.Mode)
		}
	}
	return nil
}

func (a *AutoGenerate) valid() error {
	if a == nil {
		return nil
	}
	if a.Result == nil {
		a.Result = &Results{OrderProfile: constants.RandomString}
	}
	if a.From == nil || a.To == nil {
		return errors.New("autogenerate requires From and To to be set")
	}
	if a.From.Seconds() > a.To.Seconds() {
		return errors.New("autogenerate requires a positive time interval")
	}
	if a.From.Seconds() == a.To.Seconds() && a.Every != nil {
		return errors.New("autogenerate cannot have Every set if From is equal to To")
	}
	if a.From.Seconds() != a.To.Seconds() && a.Every == nil {
		return errors.New("autogenerate with different From and To requires Every to be set")
	}
	if a.Every != nil && a.Every.Seconds() <= 0 {
		return errors.New("autogenerate requires a positive Every")
	}
	return a.Result.valid()
}

func (n *ClinicalNote) valid() error {
	if n == nil {
		return nil
	}
	if n.ContentType == "" {
		return errors.New("ClinicalNote requires the ContentType to be set")
	}
	if n.DocumentContent != "" && n.ContentType != "txt" {
		return errors.Errorf("ClinicalNote requires ContentType to be a txt file when DocumentContent is explicitly given; given ContentType: %s", n.ContentType)
	}
	return nil
}

// validAllergies returns whether the step contains valid allergies. If the step doesn't contain
// allergies, it returns true.
func (s Step) validAllergies() error {
	var a []Allergy // We're looking for an array of allergies.
	relevantType := reflect.ValueOf(a).Type()

	var ec error
	v := reflect.ValueOf(&s).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.IsNil() {
			continue
		}
		if f.Kind() != reflect.Ptr {
			continue
		}
		f = f.Elem() // All fields of a Step are pointers.
		for j := 0; j < f.NumField(); j++ {
			innerF := f.Field(j)
			if innerF.Type() != relevantType || innerF.IsNil() {
				continue
			}
			for sliceI := 0; sliceI < innerF.Len(); sliceI++ {
				a := innerF.Index(sliceI).Interface().(Allergy)
				if a.Code == "" && a.Description == "" {
					ec = combineErrors(ec, fmt.Errorf("invalid allergy %v: neither Code nor Description specified", a))
				}
				if a.IdentificationDateTime != nil && !a.IdentificationDateTime.valid() {
					ec = combineErrors(ec, fmt.Errorf("invalid IdentificationDate provided: %v", a.IdentificationDateTime))
				}
			}
		}
	}
	return ec
}

func (h *HardcodedMessage) valid() error {
	if h == nil || h.Regex != "" {
		return nil
	}
	return errors.New("a hardcoded message requires a regex")
}

func validateResultValue(r *Result) error {
	if r.Value == "" {
		return fmt.Errorf("parameter Value is missing in result: %v. If value is to be randomised, it should be set to one of: %v", r, randomValues)
	}

	if r.IsValueRandom() {
		var ec error
		if r.ReferenceRange != "" && (r.Unit == "" || randomValues[r.Unit]) {
			// if reference range is set, the random value is generated from the overridden ref ranges;
			// in this case, unit also needs to be specified
			ec = combineErrors(ec, fmt.Errorf("parameter Value set to %s and custom reference range specified, but Unit is %s. "+
				"When specifying randomised numerical Value with custom reference ranges, Unit must be specified as "+
				"non-random value", r.Value, r.Unit))
		} else if r.ReferenceRange == "" && r.Unit != r.Value && r.Unit != "" {
			// no ref ranges is specified, then the unit cannot be set, or can be set to random value as well
			ec = combineErrors(ec, fmt.Errorf("parameter Value set to %s and no custom reference ranges, but Unit is %s. When specifying "+
				"randomised numerical Value, Unit must be omitted or set to the same string as value", r.Value, r.Unit))
		}

		if r.AbnormalFlag != "" {
			ec = combineErrors(ec, fmt.Errorf("parameter Value set to %s, but AbnormalFlag is specified. AbnormalFlag can only be "+
				"overridden when Value is not random", r.Value))
		}
		return ec
	}

	if r.GetValueType() == constants.NumericalValueType {
		if r.Unit == "" {
			return fmt.Errorf("parameter Unit missing for numerical value in result: %v", r)
		}
		return nil
	}

	if r.Value != constants.EmptyString {
		// empty value may either mean: numerical or textual, so the unit may or may not be specified.
		// if the value is not empty and is not numerical - it means it is textual, and the unit should not be specified
		if r.Unit != "" && r.Unit != constants.EmptyString {
			return errors.New("textual Value has the Unit specified. Unit can only be specified for numerical Value")
		}
		return nil
	}
	return nil
}

func validateResultAgainstOrderProfile(r *Result, orderProfile *orderprofile.OrderProfile) error {
	var ec error
	if orderProfile == nil {
		// no matching order profile
		if r.IsValueRandom() && r.ReferenceRange == "" {
			ec = combineErrors(ec, fmt.Errorf("parameter Value set to %s, but no order profile found and no custom reference ranges set", r.Value))
		}
		if r.ReferenceRange != "" {
			ec = combineErrors(ec, r.validValueAndRefRange(r.ReferenceRange))
		} else if r.AbnormalFlag == constants.AbnormalFlagDefault {
			ec = combineErrors(ec, errors.New("cannot derive abnormal flag if the reference range is not specified"))
		}
	} else {
		// matching order profile
		if tt, ok := orderProfile.TestTypes[r.TestName]; ok {
			if r.ReferenceRange != "" {
				ec = combineErrors(ec, r.validValueAndRefRange(r.ReferenceRange))
			} else if tt.RefRange != "" {
				ec = combineErrors(ec, r.validValueAndRefRange(tt.RefRange))
			} else if r.AbnormalFlag == constants.AbnormalFlagDefault {
				ec = combineErrors(ec, errors.New("cannot derive abnormal flag if the reference range is not specified"))
			}
		} else {
			ec = combineErrors(ec, fmt.Errorf("test type %q doesn't exist in the order profile %s", r.TestName, orderProfile.UniversalService.Text))
		}
	}
	return ec
}

func (r *Result) valid(orderProfile *orderprofile.OrderProfile) error {
	if r == nil {
		return nil
	}

	var ec error
	if r.TestName == "" {
		ec = combineErrors(ec, fmt.Errorf("parameter TestName is missing in result: %v", r))
	}
	if r.AbnormalFlag == constants.AbnormalFlagDefault && (r.Value == constants.EmptyString || r.GetValueType() == constants.TextualValueType) {
		// If the value is textual or set to empty string, it doesn't matter what the ref range is set to; all values:
		// HIGH / LOW / NORMAL for abnormal flag are acceptable, as this is the only way to indicate whether the value
		// is normal or abnormal.
		// Only DEFAULT value is not accepted in this case, as there is no way to derive it from the textual value.
		ec = combineErrors(ec, fmt.Errorf("cannot derive abnormal flag from textual or empty value: %q", r.AbnormalFlag))
	}
	if _, ok := constants.AbnormalFlagValues[r.AbnormalFlag]; !ok {
		ec = combineErrors(ec, fmt.Errorf("invalid abnormal flag %v. Abnormal flag should be set to one of: %v", r.AbnormalFlag, constants.AbnormalFlagValues))
	}

	if err := validateResultValue(r); err != nil {
		ec = combineErrors(ec, err)
	}
	if err := validateResultAgainstOrderProfile(r, orderProfile); err != nil {
		ec = combineErrors(ec, err)
	}
	return ec
}

func (r *Result) validValueAndRefRange(refRange string) error {
	g, err := orderprofile.ValueGeneratorFromRange(refRange)
	if err != nil {
		// If the reference range cannot be parsed (eg: is set to "-") and the value is numerical, the abnormal
		// flag should not be set to HIGH / LOW / DEFAULT, as it is not possible to validate it.
		if r.GetValueType() == constants.NumericalValueType && !constants.IsNormalFlag(r.AbnormalFlag) {
			return errors.Wrapf(err, "cannot validate abnormal flag %s against the reference range for numerical value", r.AbnormalFlag)
		}
		if r.IsValueRandom() {
			return errors.Wrapf(err, "cannot generate random value from invalid reference range %q", refRange)
		}
	} else if r.IsValueRandom() {
		_, err := g.Random(r.Value)
		if err != nil {
			return errors.Wrap(err, "cannot generate random value")
		}
	} else if r.GetValueType() == constants.NumericalValueType && r.AbnormalFlag != constants.AbnormalFlagDefault {
		_, v, _ := orderprofile.ValueFromString(r.Value)
		if g.IsNormal(v) && !constants.IsNormalFlag(r.AbnormalFlag) {
			return fmt.Errorf("abnormal flag %s doesn't match expected value of: %s", r.AbnormalFlag, constants.AbnormalFlagNormal)
		}
		if g.IsHigh(v) && r.AbnormalFlag != constants.AbnormalFlagHigh {
			return fmt.Errorf("abnormal flag %s doesn't match expected value of: %s", r.AbnormalFlag, constants.AbnormalFlagHigh)
		}
		if g.IsLow(v) && r.AbnormalFlag != constants.AbnormalFlagLow {
			return fmt.Errorf("abnormal flag %s doesn't match expected value of: %s", r.AbnormalFlag, constants.AbnormalFlagLow)
		}
	}
	return nil
}

func (r *Results) valid() error {
	if r == nil {
		return nil
	}
	var ec error
	if r.OrderID == "" && r.OrderProfile == "" {
		ec = combineErrors(ec, errors.New("neither OrderId nor OrderProfile specified. Result must either relate to an Order / "+
			"other Result, or have Order Profile specified."))
	}
	if r.OrderProfile == constants.RandomString && len(r.Results) > 0 {
		ec = combineErrors(ec, fmt.Errorf("parameter OrderProfile is set to %s, but Results are specified. Results can only be specified "+
			"for the non-random OrderProfile", constants.RandomString))
	}
	if r.OrderStatus != "" && r.ResultStatus == "" {
		ec = combineErrors(ec, fmt.Errorf("parameter OrderStatus is set to %s, but ResultStatus is not set. When overriding status, "+
			"both: OrderStatus and ResultStatus need to be set.", r.OrderStatus))
	}
	if r.ResultStatus != "" && r.OrderStatus == "" {
		ec = combineErrors(ec, fmt.Errorf("parameter ResultStatus is set to %s, but OrderStatus is not set. When overriding status, "+
			"both: OrderStatus and ResultStatus need to be set.", r.ResultStatus))
	}
	if !validDate(r.CollectedDateTime) {
		ec = combineErrors(ec, fmt.Errorf("invalid CollectedDateTime: %s", r.CollectedDateTime))
	}
	if !validDate(r.ReceivedInLabDateTime) {
		ec = combineErrors(ec, fmt.Errorf("invalid ReceivedInLabDateTime: %s", r.ReceivedInLabDateTime))
	}
	te := strings.ToUpper(r.TriggerEvent)
	if te != "" && te != constants.R01 && te != constants.R03 && te != constants.R32 {
		ec = combineErrors(ec, fmt.Errorf("invalid trigger_event: %s; want R01, R03, R32 or empty", r.TriggerEvent))
	}
	return ec
}

func validDate(d string) bool {
	return d == "" || d == constants.MidnightDate || d == constants.EmptyString
}

func (o *Order) valid() error {
	if o == nil {
		return nil
	}
	if o.OrderID == "" && o.OrderProfile == "" {
		return errors.New("neither OrderId nor OrderProfile specified. Order must either relate to another Order / " +
			"Result, or have OrderProfile specified")
	}
	return nil
}

func (i *Interval) valid() error {
	if i == nil {
		return nil
	}
	if i.From > i.To {
		return errors.New("parameter To needs to be greater than or equal to From")
	}
	if i.From < 0 || i.To < 0 {
		return errors.New("parameters To and From need to be greater than or equal to 0")
	}
	return nil
}

func (a *Age) valid() error {
	if a == nil {
		return nil
	}
	if a.DayOfYear < 0 || a.DayOfYear > 365 {
		return fmt.Errorf("day of year must be in range [0, 365] if set (0 indicates random), got: %d", a.DayOfYear)
	}
	if a.From > a.To {
		return errors.New("parameter To needs to be greater than or equal to From")
	}
	return nil
}

// validatePersons makes sure that the Persons section (if present) is valid.
// All persons in Persons need to be valid persons (e.g., valid gender, etc.) as per person.valid().
// If there is more than one person, all of the following need to be met:
// * the first step needs to be a UsePatient event so that we know which person to use first.
// * all persons need to be used in at least one UsePatient step.
func (p *Pathway) validatePersons() error {
	if p.Persons == nil {
		return nil
	}
	persons := *p.Persons
	for k, v := range persons {
		if err := v.valid(); err != nil {
			return errors.Wrapf(err, "invalid person %s in pathway", k)
		}
	}

	if len(persons) == 1 {
		return nil
	}

	// Is the first step a UsePatient?
	steps := append(p.History, p.Pathway...)
	if steps[0].UsePatient == nil {
		return errors.New("use_patient needs to be the first step of the pathway")
	}

	// Are all persons used in a UsePatient step?
	unusedPersons := map[PatientID]bool{}
	for k := range persons {
		unusedPersons[k] = true
	}
	for _, s := range steps {
		if s.UsePatient != nil {
			// If it's not in the map (it's an MRN already) deleting is a no-op.
			delete(unusedPersons, s.UsePatient.Patient)
		}
	}
	if len(unusedPersons) != 0 {
		return fmt.Errorf("there are unused persons in Persons: %v", unusedPersons)
	}
	return nil
}

func validateWithRelativePositions(steps []Step, now time.Time, lm *location.Manager) error {
	var ec error
	for i, s := range steps {
		if err := s.valid(now, lm); err != nil {
			ec = combineErrors(ec, fmt.Errorf("invalid step: %v", err))
		}
		if s.StepType() == StepAddPerson && i != 0 {
			ec = combineErrors(ec, errors.New("add_person should be the first step"))
		}
	}
	return ec
}

type orderIDAndProfileValidator struct {
	orderProfiles         *orderprofile.OrderProfiles
	orderIDSeen           map[string]bool
	orderIDToOrderProfile map[string]string
}

func (v *orderIDAndProfileValidator) addOrderIDAndProfile(s Step) error {
	var ec error
	if s.Order != nil && s.Order.OrderID != "" {
		ec = combineErrors(ec, v.validateOrderIDAndOrderProfile(s.Order.OrderID, s.Order.OrderProfile))
	}
	if s.Result != nil && s.Result.OrderID != "" {
		ec = combineErrors(ec, v.validateOrderIDAndOrderProfile(s.Result.OrderID, s.Result.OrderProfile))
	}
	if s.Result != nil {
		ec = combineErrors(ec, v.validateResultAgainstOrderProfile(s.Result))
	}
	return ec
}

func (v *orderIDAndProfileValidator) validateResultAgainstOrderProfile(result *Results) error {
	var profileName string
	if result.OrderProfile != "" {
		profileName = result.OrderProfile
	} else {
		profileName = v.orderIDToOrderProfile[result.OrderID]
	}

	op, _ := v.orderProfiles.Get(profileName)
	var ec error
	for _, r := range result.Results {
		ec = combineErrors(ec, r.valid(op))
	}
	return ec
}

func (v *orderIDAndProfileValidator) validateOrderIDAndOrderProfile(orderID string, orderProfile string) error {
	var ec error
	if _, ok := v.orderIDSeen[orderID]; !ok {
		if orderProfile == "" {
			ec = combineErrors(ec, fmt.Errorf("order id %q declared first time, but no order profile specified", orderID))
		}
		v.orderIDToOrderProfile[orderID] = orderProfile
	} else if orderProfile != "" && orderProfile != v.orderIDToOrderProfile[orderID] {
		// Order Profile may be specified in all steps related to the same Order Id, as this improves readability
		// of long and complicated pathways.
		ec = combineErrors(ec, fmt.Errorf("order id %q is re-used, but order profile %v does not match previous order profile %v", orderID, orderProfile, v.orderIDToOrderProfile[orderID]))
	}

	v.orderIDSeen[orderID] = true
	return ec
}

func (d *Document) valid() error {
	if d == nil {
		return nil
	}
	var ec error
	if d.NumRandomContentLines != nil {
		if err := d.NumRandomContentLines.valid(); err != nil {
			ec = combineErrors(ec, errors.New("invalid Document.NumContentLines"))
		}
	}
	if d.UpdateType == "" {
		return ec
	}
	if d.UpdateType != Append && d.UpdateType != Overwrite {
		ec = combineErrors(ec, fmt.Errorf("Document.UpdateType must be set to `append` or `overwrite`, but was set to: %s", d.UpdateType))
	}
	if d.ID == "" {
		ec = combineErrors(ec, errors.New("Document.ID is required to update a document"))
	}
	if d.UpdateType == Append && len(d.HeaderContentLines) == 0 && len(d.EndingContentLines) == 0 && d.NumRandomContentLines != nil && d.NumRandomContentLines.From == 0 && d.NumRandomContentLines.To == 0 {
		ec = combineErrors(ec, errors.New("cannot append 0 lines to document"))
	}
	return ec
}

func (s Step) valid(now time.Time, lm *location.Manager) error {
	if s.StepType() == stepInvalid {
		return errors.New("cannot detect step type, exactly one field must be set")
	}
	if err := s.Order.valid(); err != nil {
		return errors.Wrap(err, "invalid Order step")
	}
	if err := s.Result.valid(); err != nil {
		return errors.Wrap(err, "invalid Result step")
	}
	if err := s.Admission.valid(lm); err != nil {
		return errors.Wrap(err, "invalid Admission step")
	}
	if err := s.Transfer.valid(lm); err != nil {
		return errors.Wrap(err, "invalid Transfer step")
	}
	if err := s.TransferInError.valid(lm); err != nil {
		return errors.Wrap(err, "invalid TransferInError step")
	}
	if s.Delay != nil && s.Parameters != nil && s.Parameters.DelayMessage != nil {
		return errors.New("a Delay step cannot have delay_message")
	}
	if err := s.Document.valid(); err != nil {
		return errors.Wrap(err, "invalid Document step")
	}

	if s.Parameters != nil {
		if err := s.Parameters.DelayMessage.valid(); err != nil {
			return errors.Wrapf(err, "invalid delay in parameters.delay_message")
		}
		if s.Parameters.Status != nil && s.Parameters.Status.TimeOfDeath != nil && s.Parameters.Status.TimeSinceDeath != nil {
			return errors.New("only one of TimeOfDeath and TimeSinceDeath may be set in the same step")
		}
	}
	if s.Merge != nil {
		if len(s.Merge.Children) == 0 {
			return errors.New("merge steps require at least 1 child patient")
		}
		if len(s.Merge.Children) != 1 && s.Merge.ForceA40 {
			return fmt.Errorf("force_a40 requires only one child patient, got %d", len(s.Merge.Children))
		}
		if s.Merge.Parent == "" {
			return errors.New("merge steps require one parent patient")
		}
	}
	if s.BedSwap != nil && (s.BedSwap.Patient1 == "" || s.BedSwap.Patient2 == "") {
		return errors.New("a bed swap requires patient_1 and patient_2 to be set")
	}
	if err := s.PendingAdmission.valid(lm); err != nil {
		return errors.Wrap(err, "invalid PendingAdmission step")
	}
	if err := s.PendingDischarge.valid(); err != nil {
		return errors.Wrap(err, "invalid PendingDischarge step")
	}
	if err := s.PendingTransfer.valid(lm); err != nil {
		return errors.Wrap(err, "invalid PendingTransfer step")
	}
	if err := s.TrackDeparture.valid(lm); err != nil {
		return errors.Wrap(err, "invalid TrackDeparture step")
	}
	if err := s.TrackArrival.valid(lm); err != nil {
		return errors.Wrap(err, "invalid TrackArrival step")
	}
	if err := s.Delay.valid(); err != nil {
		return errors.Wrap(err, "invalid Delay step")
	}
	if err := s.UpdatePerson.valid(now); err != nil {
		return errors.Wrap(err, "Invalid UpdatePerson step")
	}
	if err := s.PreAdmission.valid(lm); err != nil {
		return errors.Wrap(err, "invalid PreAdmission step")
	}
	if err := s.AutoGenerate.valid(); err != nil {
		return errors.Wrap(err, "invalid AutoGenerate step")
	}
	if err := s.validAllergies(); err != nil {
		return errors.Wrapf(err, "invalid %s", s.StepType())
	}
	if err := s.UsePatient.valid(); err != nil {
		return errors.Wrap(err, "invalid UsePatient step")
	}
	if s.UsePatient != nil && s.Parameters != nil && s.Parameters.DelayMessage != nil {
		return errors.New("invalid UsePatient step: delay_message is not allowed")
	}
	if err := s.ClinicalNote.valid(); err != nil {
		return errors.Wrap(err, "invalid SendNote step")
	}
	if err := s.HardcodedMessage.valid(); err != nil {
		return errors.Wrap(err, "invalid HardcodedMessage step")
	}
	return nil
}

func (up *UpdatePerson) valid(now time.Time) error {
	if up == nil {
		return nil
	}
	ec := up.Person.valid()
	for _, d := range up.Diagnoses {
		if err := d.valid(now); err != nil {
			ec = combineErrors(ec, errors.Wrap(err, "invalid diagnosis"))
		}
	}
	for _, p := range up.Procedures {
		if err := p.valid(now); err != nil {
			ec = combineErrors(ec, errors.Wrap(err, "invalid procedure"))
		}
	}
	return ec
}

func (d *DateTime) valid() bool {
	timeWasProvided := d.Time != nil
	timeFromNowWasProvided := d.TimeFromNow != nil
	noDateTimeRecordedWasProvided := d.NoDateTimeRecorded

	numProvided := 0
	for _, b := range []bool{timeWasProvided, timeFromNowWasProvided, noDateTimeRecordedWasProvided} {
		if b {
			numProvided++
		}
	}
	return numProvided == 1
}

func (d *DiagnosisOrProcedure) valid(now time.Time) error {
	if d.Code == constants.RandomString || d.Description == constants.RandomString {
		if !(d.Code == constants.RandomString || d.Code == "") || !(d.Description == constants.RandomString || d.Description == "") {
			return fmt.Errorf("code=%s but description=%s; either both need to be RANDOM, or one of them needs to be omitted", d.Code, d.Description)
		}
		if d.Type != "" {
			return fmt.Errorf("type cannot be set for random item, but was set to: %s", d.Type)
		}
		if d.DateTime != nil && (!d.DateTime.valid() || !d.DateTime.IsBefore(now)) {
			return fmt.Errorf("invalid or future datetime provided: %v", *d.DateTime)
		}
	} else {
		if d.DateTime != nil && (!d.DateTime.valid() || !d.DateTime.IsBefore(now)) {
			return fmt.Errorf("invalid or future datetime provided: %v", *d.DateTime)
		}
		if d.Code == "" && d.Description == "" {
			return fmt.Errorf("neither Code nor Description specified")
		}
	}
	return nil
}

func strExists(str *string) bool {
	return str != nil && *str != ""
}

// Consultant Validation
// [ID only]: Look in the Doctors map. If the ID isn't present, fail validation.
// If the ID is present, validation succeeds. When we run the pathway, load the doctor by the ID.
// [LastName, FirstName]: Validation succeeds. When the pathway runs, if the combination is present,
// load such doctor. If not, create a new doctor with that data, with a new ID. Store the new Doctor
// in the map so that it can be reused in the future.
// [ID, LastName, FirstName]: Look in the Doctors map. If the ID is present in the map but the rest
// of the fields don't match with the content of the map, fail validation. If the ID is not present
// in the map, but the combination of LastName and FirstName exists in the map, fail validation.
// Otherwise, validation succeeds.
// The reason field is not used if we return true, but has been added to improve readability.
func (c *Consultant) valid(doctors *doctor.Doctors) (*Consultant, error) {
	if c == nil {
		return nil, nil
	}
	if strExists(c.FirstName) && strExists(c.Surname) && strExists(c.ID) {
		// Given both ID and name
		doctor := doctors.GetByID(*c.ID)
		if doctor != nil {
			if doctor.FirstName != *c.FirstName || doctor.Surname != *c.Surname {
				return nil, fmt.Errorf("consultant ID matched but name did not. Expected: %s %s Actual: %s %s",
					doctor.FirstName, doctor.Surname, *c.FirstName, *c.Surname)
			}
			updatePrefix(c, doctor)
			return c, nil
		}
		doctor = doctors.GetByName(*c.FirstName, *c.Surname)
		if doctor != nil && doctor.ID != *c.ID {
			return nil, fmt.Errorf("consultant name matched but ID did not. Expected: %s Actual: %s", doctor.ID, *c.ID)
		}
		updatePrefix(c, doctor)
		return c, nil
	} else if !strExists(c.FirstName) && !strExists(c.Surname) && strExists(c.ID) {
		// ID only; name blank
		doctor := doctors.GetByID(*c.ID)
		if doctor == nil {
			return nil, fmt.Errorf("consultant ID %s not found. If the consultant name is not specified, the ID must be an existing ID", *c.ID)
		}
		c.FirstName = &doctor.FirstName
		c.Surname = &doctor.Surname
		updatePrefix(c, doctor)
		return c, nil
	} else if strExists(c.FirstName) && strExists(c.Surname) && !strExists(c.ID) {
		// name only; ID blank
		doctor := doctors.GetByName(*c.FirstName, *c.Surname)
		if doctor != nil {
			c.ID = &doctor.ID
			updatePrefix(c, doctor)
			return c, nil
		}
		newID := newConsultantID(doctors)
		c.ID = &newID
		return c, nil
	}

	return nil, errors.New("consultant validation failed. Either ID or First Name and Surname must be provided for a consultant to be valid")
}

func updatePrefix(c *Consultant, withPrefix *ir.Doctor) {
	if withPrefix != nil && c.Prefix == nil {
		c.Prefix = &withPrefix.Prefix
	}
}

func newConsultantID(doctors *doctor.Doctors) string {
	for {
		var sb strings.Builder
		sb.WriteString("C")
		for i := 0; i < 7; i++ {
			sb.WriteString(strconv.Itoa(rand.Intn(10)))
		}
		newID := sb.String()
		doctor := doctors.GetByID(newID)
		if doctor == nil {
			return newID
		}
	}
}

func (p *Person) valid() error {
	if p == nil {
		// Person is optional.
		return nil
	}
	if err := p.Address.valid(); err != nil {
		return errors.Wrapf(err, "invalid address")
	}
	if p.Age != nil && p.DateOfBirth != nil {
		return errors.New("only one of Age or DateOfBirth may be set for the same Person")
	}
	if err := p.Age.valid(); err != nil {
		return errors.Wrapf(err, "invalid age")
	}
	g := p.Gender
	if g == "" || g == constants.RandomString || g == Male || g == Female {
		// Gender can be "", a string of the form "RANDOM...", Male, or Female.
		return nil
	}
	return fmt.Errorf("unknown gender: %s", g)
}

func (a *Address) valid() error {
	if a == nil {
		return nil
	}
	if a.AllRandom {
		// No other field can be set.
		v := reflect.ValueOf(a).Elem()
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			if f != v.FieldByName("AllRandom") && f.Len() > 0 {
				return fmt.Errorf("%s was set. If all_random is enabled, no other field can be set", v.Type().Field(i).Name)
			}
		}
		return nil
	}

	if a.SecondLine.IsSet() && !a.FirstLine.IsSet() {
		return errors.New("address.SecondLine was set but address.FirstLine was not; both lines must be set together")
	}
	return nil
}

func (p *Percentage) valid() error {
	if p == nil {
		return nil
	}
	if *p < 0 {
		return fmt.Errorf("expected >=0, got %v%%", *p)
	}
	if *p != 0 {
		significantDigits := p.significantDigits()
		if significantDigits > maxSignificantDigits {
			return fmt.Errorf("expected %v decimal points at most, got %v (%v)", maxSignificantDigits, significantDigits, *p)
		}
	}
	return nil
}

func validateHistory(history []Step, clock clock.Clock, lm *location.Manager, validator orderIDAndProfileValidator) error {
	var ec error
	if err := validateWithRelativePositions(history, clock.Now(), lm); err != nil {
		ec = combineErrors(ec, err)
	}

	for _, s := range history {
		if s.Delay != nil {
			ec = combineErrors(ec, errors.New("delays in historical steps are not supported"))
		}
		if s.AutoGenerate != nil {
			ec = combineErrors(ec, errors.New("step AutoGenerate in historical steps is not supported"))
		}
		if s.UsePatient == nil {
			if s.Parameters == nil || s.Parameters.TimeFromNow == nil || s.Parameters.TimeFromNow.Seconds() >= 0 {
				ec = combineErrors(ec, errors.New("parameters.time_from_now must be set and negative for a historical step"))
			}
		}
		ec = combineErrors(ec, validator.addOrderIDAndProfile(s))
	}
	return ec
}

func validatePathway(pathway []Step, clock clock.Clock, lm *location.Manager, validator orderIDAndProfileValidator) error {
	var ec error
	if err := validateWithRelativePositions(pathway, clock.Now(), lm); err != nil {
		ec = combineErrors(ec, err)
	}

	for _, s := range pathway {
		if s.Parameters != nil && s.Parameters.TimeFromNow != nil {
			ec = combineErrors(ec, errors.New("parameters.time_from_now in Pathway steps is not supported"))
		}
		ec = combineErrors(ec, validator.addOrderIDAndProfile(s))
	}
	return ec
}

// Valid returns whether the pathway is valid.
// It applies custom validation that depends on whether the steps are historical or not.
// Returns an error if the pathway is invalid.
func (p *Pathway) Valid(clock clock.Clock, orderProfiles *orderprofile.OrderProfiles, doctors *doctor.Doctors, lm *location.Manager, validFn func(*Pathway) error) error {
	var ec error

	if len(p.History) == 0 && len(p.Pathway) == 0 {
		ec = combineErrors(ec, errors.New("pathway, historical_data and merge cannot be empty at the same time"))
	}

	consultant, err := p.Consultant.valid(doctors)
	if err != nil {
		ec = combineErrors(ec, errors.Wrap(err, "invalid consultant"))
	}
	p.Consultant = consultant

	if err = p.Percentage.valid(); err != nil {
		ec = combineErrors(ec, errors.Wrap(err, "invalid percentage of patients"))
	}

	if validFn != nil {
		if err = validFn(p); err != nil {
			ec = combineErrors(ec, errors.Wrap(err, "invalid based on valid function"))
		}
	}

	if err := p.validatePersons(); err != nil {
		ec = combineErrors(ec, err)
	}

	validator := orderIDAndProfileValidator{
		orderProfiles:         orderProfiles,
		orderIDSeen:           make(map[string]bool),
		orderIDToOrderProfile: make(map[string]string),
	}
	if err := validateHistory(p.History, clock, lm, validator); err != nil {
		ec = combineErrors(ec, err)
	}
	if err := validatePathway(p.Pathway, clock, lm, validator); err != nil {
		ec = combineErrors(ec, err)
	}

	if ec != nil {
		log.WithField("pathway_name", p.Name()).Error(ec)
	}
	return ec
}

// errorCollection is a specialization type of the error interface. It allows multiple
// error objects to be combined into a single meta-error.
type errorCollection []error

// Error serializes all the errors within an errorCollection.
func (err errorCollection) Error() string {
	var errStrings []string
	for _, suberr := range err {
		errStrings = append(errStrings, suberr.Error())
	}
	return strings.Join(errStrings, "\n")
}

// combineErrors merges a sequence of errors (an error may be an errorCollection) into a single error.
func combineErrors(otherErrors ...error) error {
	var err errorCollection
	for _, otherError := range otherErrors {
		if otherError == nil {
			continue
		}
		if otherEC, ok := otherError.(errorCollection); ok {
			err = append(err, otherEC...)
		} else {
			err = append(err, otherError)
		}
	}
	if len(err) != 0 {
		return err
	}
	return nil
}
