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
	"fmt"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/simhospital/pkg/config"
	"github.com/google/simhospital/pkg/constants"
	"github.com/google/simhospital/pkg/gender"
	"github.com/google/simhospital/pkg/generator/names"
	"github.com/google/simhospital/pkg/ir"
	"github.com/google/simhospital/pkg/pathway"
	"github.com/google/simhospital/pkg/test"
	"github.com/google/simhospital/pkg/test/testclock"
	"github.com/google/simhospital/pkg/test/testid"
	"github.com/google/simhospital/pkg/test/testperson"
	"github.com/google/simhospital/pkg/test/testwrite"
)

var (
	defaultNow = time.Date(2019, 1, 20, 0, 0, 0, 0, time.UTC)

	maleName         = "John"
	maleSurname      = "Smith"
	malePrefix       = "Mr"
	maleSuffix       = "Jr"
	maleDegree       = "Dr"
	maleGender       = "M"
	defaultEthnicity = &ir.Ethnicity{
		ID:   "White",
		Text: "White",
	}
	defaultAddress = &ir.Address{
		FirstLine:  "111 Big House",
		SecondLine: "Long Way",
		City:       "London",
		PostalCode: "AB1 2CD",
		Country:    "GB",
		Type:       "HOME",
	}
	defaultMalePerson = &ir.Person{
		Prefix:     malePrefix,
		FirstName:  maleName,
		MiddleName: maleName,
		Surname:    maleSurname,
		Suffix:     maleSuffix,
		Degree:     maleDegree,
		Gender:     maleGender,
		Ethnicity:  defaultEthnicity,
		Address:    defaultAddress,
		MRN:        "1",
	}
)

func TestNewPerson(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name string
		p    *pathway.Person
		want *ir.Person
	}{
		{
			name: "All values generated",
			p:    &pathway.Person{},
			want: defaultMalePerson,
		}, {
			name: "Set MRN, FirstName, Surname and Address",
			p: &pathway.Person{
				MRN:       "123456",
				FirstName: "Bob",
				Surname:   "Dylan",
				Address: &pathway.Address{
					FirstLine:  "999 Small House",
					SecondLine: "Short Street",
					City:       "Croydon",
					Postcode:   "XY9 8ZZ",
					Country:    "UK",
					Type:       "HOME",
				},
			},
			want: &ir.Person{
				Prefix:     malePrefix,
				FirstName:  "Bob",
				MiddleName: maleName,
				Surname:    "Dylan",
				Suffix:     maleSuffix,
				Degree:     maleDegree,
				Gender:     maleGender,
				Ethnicity:  defaultEthnicity,
				Address: &ir.Address{
					FirstLine:  "999 Small House",
					SecondLine: "Short Street",
					City:       "Croydon",
					PostalCode: "XY9 8ZZ",
					Country:    "UK",
					Type:       "HOME",
				},
				MRN: "123456",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := simpleMaleGenerator(ctx, t, defaultNow)
			got := g.NewPerson(tc.p)

			ignoreFields := cmpopts.IgnoreFields(ir.Person{}, "Birth", "PhoneNumber", "NHS")
			if diff := cmp.Diff(tc.want, got, ignoreFields); diff != "" {
				t.Errorf("g.NewPerson(%v) diff: (-want, +got):\n%s", tc.p, diff)
			}

			if got, want := got.Birth.Year(), defaultNow.Year(); got > want {
				t.Errorf("g.NewPerson(%v) got person.Birth.Year() = %d, want at most %v", tc.p, got, want)
			}
			if err := validatePhoneNumber(got.PhoneNumber); err != nil {
				t.Errorf("g.NewPerson(%v) got person.PhoneNumber = %s is invalid: %q", tc.p, got.PhoneNumber, err)
			}
			if err := validateNHS(got.NHS); err != nil {
				t.Errorf("g.NewPerson(%v) got person.NHS = %s is invalid: %q", tc.p, got.NHS, err)
			}
		})
	}
}

func TestNewPersonWithNHS(t *testing.T) {
	ctx := context.Background()
	g := simpleMaleGenerator(ctx, t, defaultNow)
	want := "0714630667"
	pathway := &pathway.Person{
		NHS: want,
	}
	got := g.NewPerson(pathway)
	if got.NHS != want {
		t.Errorf("g.NewPerson(%v) got person.NHS=%q, want=%q", pathway, got.NHS, want)
	}
}

func TestNewPersonWithDOB(t *testing.T) {
	ctx := context.Background()
	birthdate := time.Date(1984, 2, 12, 0, 0, 0, 0, time.UTC)
	g, _, dataConfig := testGenerator(ctx, t, defaultNow)
	cases := []struct {
		name          string
		p             *pathway.Person
		wantExactDate *time.Time
		wantDOBFrom   int
		wantDOBTo     int
		// First name is chosen based on the DOB.
		// See: simulated_hospital/pkg/test/data/historicname_boys_test.csv
		wantName []string
	}{
		{
			name:        "Default age",
			p:           &pathway.Person{Gender: pathway.Male},
			wantDOBFrom: 0,
			wantDOBTo:   defaultNow.Year(),
			wantName:    dataConfig.FirstNames.Boys.All,
		}, {
			name: "Age range",
			p: &pathway.Person{
				Age: &pathway.Age{From: 20, To: 30},
			},
			wantDOBFrom: defaultNow.Year() - 30,
			wantDOBTo:   defaultNow.Year() - 20,
			// 2019 - 30 = 1989 -> bucket 1994
			// 2019 - 20 = 1999 -> bucket 1994
			wantName: dataConfig.FirstNames.Boys.ByYear[1994],
		}, {
			name: "Exact Birthdate",
			p: &pathway.Person{
				DateOfBirth: &birthdate,
			},
			wantExactDate: &birthdate,
			wantName:      dataConfig.FirstNames.Boys.ByYear[1984],
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			got := g.NewPerson(tc.p)

			if tc.wantExactDate != nil {
				if got.Birth.Time != *tc.wantExactDate {
					t.Errorf("g.NewPerson(%v) got person.Birth.Time=%v, want=%v", tc.p, got.Birth.Time, *tc.wantExactDate)
				}
				return
			}
			if got := got.Birth.Year(); got < tc.wantDOBFrom || got > tc.wantDOBTo {
				t.Errorf("g.NewPerson(%v) got person.Birth.Year()=%d, want between (%d, %d)", tc.p, got, tc.wantDOBFrom, tc.wantDOBTo)
			}
		})
	}
}

func TestNewPersonWithGender(t *testing.T) {
	ctx := context.Background()
	g, hl7Config, dataConfig := testGeneratorWithCustomHL7Gender(ctx, t, defaultNow)
	cases := []struct {
		name           string
		p              *pathway.Person
		wantGender     []string
		wantFirstName  []string
		wantMiddleName []string
	}{
		{
			name: "Male",
			p: &pathway.Person{
				Gender: pathway.Male,
			},
			wantGender:     []string{hl7Config.Gender.Male},
			wantFirstName:  dataConfig.FirstNames.Boys.All,
			wantMiddleName: append(dataConfig.FirstNames.Boys.All, ""),
		}, {
			name: "Female",
			p: &pathway.Person{
				Gender: pathway.Female,
			},
			wantGender:     []string{hl7Config.Gender.Female},
			wantFirstName:  dataConfig.FirstNames.Girls.All,
			wantMiddleName: append(dataConfig.FirstNames.Girls.All, ""),
		}, {
			name: "Female with name explicit from the pathway",
			p: &pathway.Person{
				Gender:    pathway.Female,
				FirstName: "Alice",
			},
			wantGender:     []string{hl7Config.Gender.Female},
			wantFirstName:  []string{"Alice"},
			wantMiddleName: append(dataConfig.FirstNames.Girls.All, ""),
		}, {
			name: "Random",
			p: &pathway.Person{
				Gender: constants.RandomString,
			},
			wantGender:     []string{hl7Config.Gender.Male, hl7Config.Gender.Female},
			wantFirstName:  append(dataConfig.FirstNames.Boys.All, dataConfig.FirstNames.Girls.All...),
			wantMiddleName: append(append(dataConfig.FirstNames.Boys.All, ""), dataConfig.FirstNames.Girls.All...),
		}, {
			name:           "Not specified",
			p:              &pathway.Person{},
			wantGender:     []string{hl7Config.Gender.Male, hl7Config.Gender.Female},
			wantFirstName:  append(dataConfig.FirstNames.Boys.All, dataConfig.FirstNames.Girls.All...),
			wantMiddleName: append(append(dataConfig.FirstNames.Boys.All, ""), dataConfig.FirstNames.Girls.All...),
		},
	}

	for _, tc := range cases {
		for i := 0; i < 100; i++ {
			t.Run(fmt.Sprintf("%s#%d", tc.name, i), func(t *testing.T) {
				got := g.NewPerson(tc.p)

				if !contains(got.Gender, tc.wantGender) {
					t.Errorf("g.NewPerson(%v) got person.Gender=%v, want in %v", tc.p, got.Gender, tc.wantGender)
				}
				if !contains(got.FirstName, tc.wantFirstName) {
					t.Errorf("g.NewPerson(%v) got person.FirstName=%v, want in %v", tc.p, got.FirstName, tc.wantFirstName)
				}
				if !contains(got.MiddleName, tc.wantMiddleName) {
					t.Errorf("g.NewPerson(%v) got person.MiddleName=%v, want in %v", tc.p, got.MiddleName, tc.wantMiddleName)
				}
			})
		}
	}
}

func TestUpdatePersonFromPathway(t *testing.T) {
	ctx := context.Background()
	birthdate := time.Date(1984, 2, 12, 0, 0, 0, 0, time.UTC)
	cases := []struct {
		name    string
		person  func() *ir.Person
		pathway *pathway.Person
		want    func() *ir.Person
	}{
		{
			name:    "No changes",
			person:  testperson.New,
			pathway: &pathway.Person{},
			want: func() *ir.Person {
				return testperson.New()
			},
		}, {
			name:   "Update first line of address",
			person: testperson.New,
			pathway: &pathway.Person{
				Address: &pathway.Address{FirstLine: "12 New House"},
			},
			want: func() *ir.Person {
				newPerson := testperson.New()
				newPerson.Address.FirstLine = "12 New House"
				newPerson.Address.SecondLine = ""
				return newPerson
			},
		}, {
			name:   "Update first and second line of address",
			person: testperson.New,
			pathway: &pathway.Person{
				Address: &pathway.Address{FirstLine: "12 New House", SecondLine: "New Street"},
			},
			want: func() *ir.Person {
				newPerson := testperson.New()
				newPerson.Address.FirstLine = "12 New House"
				newPerson.Address.SecondLine = "New Street"
				return newPerson
			},
		}, {
			name:   "Update random address",
			person: testperson.New,
			pathway: &pathway.Person{
				Address: &pathway.Address{AllRandom: true},
			},
			want: func() *ir.Person {
				newPerson := testperson.New()
				newPerson.Address = defaultAddress
				return newPerson
			},
		}, {
			name:    "Empty address - no change",
			person:  testperson.New,
			pathway: &pathway.Person{Address: &pathway.Address{}},
			want: func() *ir.Person {
				return testperson.New()
			},
		}, {
			name:    "Update gender - change gender, prefix and middle name",
			person:  testperson.New,
			pathway: &pathway.Person{Gender: pathway.Male},
			want: func() *ir.Person {
				newPerson := testperson.New()
				newPerson.Gender = maleGender
				newPerson.Prefix = malePrefix
				newPerson.MiddleName = maleName
				return newPerson
			},
		}, {
			name:    "Update gender and name - change name, gender, prefix and middle name",
			person:  testperson.New,
			pathway: &pathway.Person{FirstName: "Adam", Gender: pathway.Male},
			want: func() *ir.Person {
				newPerson := testperson.New()
				newPerson.FirstName = "Adam"
				newPerson.Gender = maleGender
				newPerson.Prefix = malePrefix
				newPerson.MiddleName = maleName
				return newPerson
			},
		}, {
			name:    "Update MRN and NHS",
			person:  testperson.New,
			pathway: &pathway.Person{MRN: "987654", NHS: "0714630333"},
			want: func() *ir.Person {
				newPerson := testperson.New()
				newPerson.MRN = "987654"
				newPerson.NHS = "0714630333"
				return newPerson
			},
		}, {
			name:    "Update date of birth",
			person:  testperson.New,
			pathway: &pathway.Person{DateOfBirth: &birthdate},
			want: func() *ir.Person {
				newPerson := testperson.New()
				newPerson.Birth = ir.NewValidTime(birthdate)
				return newPerson
			},
		}, {
			name: "Update empty fields",
			person: func() *ir.Person {
				return &ir.Person{
					NHS:   "0714630667",
					Birth: ir.NewValidTime(defaultNow),
				}
			},
			pathway: &pathway.Person{},
			want: func() *ir.Person {
				return &ir.Person{
					Prefix:     malePrefix,
					FirstName:  maleName,
					MiddleName: maleName,
					Surname:    maleSurname,
					Gender:     maleGender,
					Address:    defaultAddress,
					MRN:        "1",
					NHS:        "0714630667",
					Birth:      ir.NewValidTime(defaultNow),
				}
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := simpleMaleGenerator(ctx, t, defaultNow)
			got := tc.person()
			g.UpdatePersonFromPathway(got, tc.pathway)

			if diff := cmp.Diff(tc.want(), got); diff != "" {
				t.Errorf("g.UpdatePersonFromPathway(%+v, %+v) diff: (-want, +got):\n%s", tc.person(), tc.pathway, diff)
			}
		})
	}
}

func TestUpdatePersonFromPathwayUpdateDateOfBirth(t *testing.T) {
	ctx := context.Background()
	birthdate := time.Date(1984, 2, 12, 0, 0, 0, 0, time.UTC)
	cases := []struct {
		name        string
		person      func() *ir.Person
		pathway     *pathway.Person
		wantDOBFrom int
		wantDOBTo   int
	}{
		{
			name:        "No changes",
			person:      testperson.New,
			pathway:     &pathway.Person{},
			wantDOBFrom: testperson.New().Birth.Year(),
			wantDOBTo:   testperson.New().Birth.Year(),
		}, {
			name:        "Empty Age - update Age From:0 To:0",
			person:      testperson.New,
			pathway:     &pathway.Person{Age: &pathway.Age{}},
			wantDOBFrom: defaultNow.Year(),
			wantDOBTo:   defaultNow.Year(),
		}, {
			name:   "Update DOB from Age",
			person: testperson.New,
			pathway: &pathway.Person{
				Age: &pathway.Age{From: 20, To: 25},
			},
			wantDOBFrom: defaultNow.Year() - 25,
			wantDOBTo:   defaultNow.Year() - 20,
		}, {
			name:   "Update DOB from DateOfBirth",
			person: testperson.New,
			pathway: &pathway.Person{
				DateOfBirth: &birthdate,
			},
			wantDOBFrom: birthdate.Year(),
			wantDOBTo:   birthdate.Year(),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := simpleMaleGenerator(ctx, t, defaultNow)
			got := tc.person()
			g.UpdatePersonFromPathway(got, tc.pathway)

			if got := got.Birth.Year(); got < tc.wantDOBFrom || got > tc.wantDOBTo {
				t.Errorf("g.UpdatePersonFromPathway(%+v, %+v) got person.Birth.Year()=%d, want between (%d, %d)", tc.person(), tc.pathway, got, tc.wantDOBFrom, tc.wantDOBTo)
			}
		})
	}
}

func TestUpdatePersonFromPathwaySetNHSIfEmpty(t *testing.T) {
	ctx := context.Background()
	g := simpleMaleGenerator(ctx, t, defaultNow)
	pathway := &pathway.Person{}
	person := testperson.New()
	person.NHS = ""

	original := new(ir.Person)
	*original = *person
	g.UpdatePersonFromPathway(person, pathway)

	if err := validateNHS(person.NHS); err != nil {
		t.Errorf("g.UpdatePersonFromPathway(%+v, %+v) got person.NHS = %s is invalid: %q", original, pathway, person.NHS, err)
	}
}

func TestUpdatePersonFromPathwayNilPathway(t *testing.T) {
	ctx := context.Background()
	g := simpleMaleGenerator(ctx, t, defaultNow)
	var pathway *pathway.Person
	// Suffix, Degree,  Ethnicity and PhoneNumber cannot be set in the pathway,
	// so they're never updated in UpdatePersonFromPathway.
	// Other fields will be set to random values.
	want := *defaultMalePerson
	want.Suffix = ""
	want.Degree = ""
	want.Ethnicity = nil
	want.PhoneNumber = ""
	got := &ir.Person{}
	original := new(ir.Person)
	*original = *got

	g.UpdatePersonFromPathway(got, pathway)

	ignoreFields := cmpopts.IgnoreFields(ir.Person{}, "Birth", "NHS")
	if diff := cmp.Diff(&want, got, ignoreFields); diff != "" {
		t.Errorf("g.UpdatePersonFromPathway(%v, nil) diff: (-want, +got):\n%s", original, diff)
	}

	if got, want := got.Birth.Year(), defaultNow.Year(); got > want {
		t.Errorf("g.UpdatePersonFromPathway(%v, nil) got person.Birth.Year() = %d, want at most %v", original, got, want)
	}
	if err := validateNHS(got.NHS); err != nil {
		t.Errorf("g.UpdatePersonFromPathway(%v, nil) got person.NHS = %s is invalid: %q", original, got.NHS, err)
	}
}

func validatePhoneNumber(num string) error {
	if num == "" {
		return fmt.Errorf("got empty phone number %q, want it to be nonempty", num)
	}
	regex := "^[0-9]{3} [0-9]{4} [0-9]{4}$"
	m, err := regexp.MatchString(regex, num)
	if err != nil {
		return fmt.Errorf("regexp.MatchString(%q, %q) got err=%q, want no error", regex, num, err)
	}
	if !m {
		return fmt.Errorf("got phone number %q, want it to match %q", num, regex)
	}
	return nil
}

// validateNHS is a helper to check NHS number for validity; see
// http://www.datadictionary.nhs.uk/version2/data_dictionary/data_field_notes/n/nhs_number_de.asp?shownav=0
func validateNHS(nhs string) error {
	if len(nhs) != 10 {
		return fmt.Errorf("got NHS number %q, want it to be 10 digits in length", nhs)
	}
	nhsNumber, err := strconv.Atoi(nhs)
	if err != nil {
		return fmt.Errorf("strconv.Atoi(%q) returned unexpected err=%q", nhs, err)
	}
	sum := 0
	leadingDigits := nhsNumber / 10
	for i := 2; i <= 10; i++ {
		sum += (leadingDigits % 10) * i
		leadingDigits /= 10
	}
	remainder := sum % 11
	expectedCheckDigit := (11 - remainder) % 11
	if nhsNumber%10 != expectedCheckDigit {
		return fmt.Errorf("nhsNumber %% 10 = %d, want check digit %d", nhsNumber%10, expectedCheckDigit)
	}
	return nil
}

func contains(str string, strs []string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

type fakeAddressGenerator struct {
	want *ir.Address
}

func (f *fakeAddressGenerator) Random() *ir.Address {
	return f.want
}

func simpleMaleGenerator(ctx context.Context, t *testing.T, date time.Time) *Generator {
	t.Helper()

	hl7Config, err := config.LoadHL7Config(ctx, test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	hl7Config.Gender = config.Gender{
		Male:   maleGender,
		Female: maleGender,
	}

	name := `
RANK,1904
1,John
`
	fNames := testwrite.BytesToFile(t, []byte(name))
	surnames := `Smith`
	fSurnames := testwrite.BytesToFile(t, []byte(surnames))
	ethnicity := `White,White,1`
	fEthnicity := testwrite.BytesToFile(t, []byte(ethnicity))
	data := `
patient_name:
  degrees:
    - "Dr"
  degree_percentage: 100
  suffixes:
    - "Jr"
  suffix_percentage: 100
  female_prefixes:
    - "Mr"
  male_prefixes:
    - "Mr"
  middlename_percentage: 100
`
	fData := testwrite.BytesToFile(t, []byte(data))

	df := test.DataFiles[test.Test]
	df.Boys = fNames
	df.Girls = fNames
	df.Surnames = fSurnames
	df.Ethnicities = fEthnicity
	df.DataConfig = fData

	dataCFG, err := config.LoadData(ctx, df, hl7Config)
	if err != nil {
		t.Fatalf("LoadData(%+v, %+v) failed with %v", df, hl7Config, err)
	}

	return &Generator{
		Clock:              testclock.New(date),
		NameGenerator:      &names.Generator{Data: dataCFG},
		GenderConvertor:    gender.NewConvertor(hl7Config),
		EthnicityGenerator: NewEthnicityGenerator(dataCFG),
		AddressGenerator: &fakeAddressGenerator{
			want: defaultAddress,
		},
		MRNGenerator: &testid.Generator{},
	}
}

func testGenerator(ctx context.Context, t *testing.T, date time.Time) (*Generator, *config.HL7Config, *config.Data) {
	t.Helper()
	hl7Config, err := config.LoadHL7Config(ctx, test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	dataF := test.DataFiles[test.Test]
	dataCFG, err := config.LoadData(ctx, dataF, hl7Config)
	if err != nil {
		t.Fatalf("LoadData(%+v, %+v) failed with %v", dataF, hl7Config, err)
	}

	return &Generator{
		Clock:              testclock.New(date),
		NameGenerator:      &names.Generator{Data: dataCFG},
		GenderConvertor:    gender.NewConvertor(hl7Config),
		EthnicityGenerator: NewEthnicityGenerator(dataCFG),
		AddressGenerator:   &fakeAddressGenerator{},
		MRNGenerator:       &testid.Generator{},
	}, hl7Config, dataCFG
}

func testGeneratorWithCustomHL7Gender(ctx context.Context, t *testing.T, date time.Time) (*Generator, *config.HL7Config, *config.Data) {
	t.Helper()
	hl7Config, err := config.LoadHL7Config(ctx, test.MessageConfigTest)
	if err != nil {
		t.Fatalf("LoadHL7Config(%s) failed with %v", test.MessageConfigTest, err)
	}
	hl7Config.Gender.Male = "HL7_M"
	hl7Config.Gender.Female = "HL7_F"
	hl7Config.Gender.Unknown = "HL7_U"
	dataF := test.DataFiles[test.Test]
	dataCFG, err := config.LoadData(ctx, dataF, hl7Config)
	if err != nil {
		t.Fatalf("LoadData(%+v, %+v) failed with %v", dataF, hl7Config, err)
	}

	return &Generator{
		Clock:              testclock.New(date),
		NameGenerator:      &names.Generator{Data: dataCFG},
		GenderConvertor:    gender.NewConvertor(hl7Config),
		EthnicityGenerator: NewEthnicityGenerator(dataCFG),
		AddressGenerator:   &fakeAddressGenerator{},
		MRNGenerator:       &testid.Generator{},
	}, hl7Config, dataCFG
}
