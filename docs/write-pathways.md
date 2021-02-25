# Write pathways

*   [Overview](#overview)
*   [Sections of a pathway](#sections-of-a-pathway)
    +   [Persons](#persons)
    +   [Consultant](#consultant)
    +   [Percentage of patients](#percentage-of-patients)
    +   [Historical Data](#historical-data)
    +   [Pathway](#pathway)
*   [Steps](#steps)
    +   [Use Patient](#use-patient)
    +   [Delay](#delay)
    +   [Admission](#admission)
    +   [Transfer](#transfer)
    +   [Transfer in Error](#transfer-in-error)
    +   [Document](#document)
    +   [Discharge](#discharge)
    +   [Discharge in Error](#discharge-in-error)
    +   [Registration](#registration)
    +   [Cancel Admit/ Visit](#cancel-admit-visit)
    +   [Cancel Transfer](#cancel-transfer)
    +   [Cancel Discharge](#cancel-discharge)
    +   [Add Person](#add-person)
    +   [Update Person](#update-person)
    +   [Pre Admission](#pre-admission)
    +   [Pending Admission](#pending-admission)
    +   [Pending Transfer](#pending-transfer)
    +   [Pending Discharge](#pending-discharge)
    +   [Delete Visit](#delete-visit)
    +   [Cancel Pending Admission](#cancel-pending-admission)
    +   [Cancel Pending Transfer](#cancel-pending-transfer)
    +   [Cancel Pending Discharge](#cancel-pending-discharge)
    +   [Order](#order)
        -   [Order Acknowledgement](#order-acknowledgement)
    +   [Results](#results)
        -   [Midnight Case](#midnight-case)
    +   [Merge](#merge)
    +   [Bed swap](#bed-swap)
    +   [Track Departure / Track Arrival](#track-departure-track-arrival)
    +   [AutoGenerate](#autogenerate)
    +   [Clinical Note](#clinical-note)
    +   [Hardcoded message](#hardcoded-message)
    +   [Generic](#generic)
    +   [GenerateResources](#generate-resources)
*   [Order profiles](#order-profiles)
    +   [Explicitly specify results for each test type in the order profile
        (recommended)](#explicitly-specify-results-for-each-test-type-in-the-order-profile-recommended)
    +   [Only specify results for a subset of test types](#only-specify-results-for-a-subset-of-test-types)
    +   [Explicitly specify result as random](#explicitly-specify-result-as-random)
    +   [Do not specify results - implicitly random](#do-not-specify-results-implicitly-random)
    +   [Results for a random order profile](#results-for-a-random-order-profile)
    +   [Results for non-existing order profile](#results-for-non-existing-order-profile)
    +   [Results with an unknown name for an existing order profile](#results-with-an-unknown-name-for-an-existing-order-profile)
*   [Default values for Order Profiles](#default-values-for-order-profiles)
    +   [Reference range](#reference-range)
*   [Pathway with multiple Orders and Results](#pathway-with-multiple-orders-and-results)
*   [Pathway with Result but no Order](#pathway-with-result-but-no-order)
*   [Pathways with multiple Results with the same order_id](#pathways-with-multiple-results-with-the-same-order-id)
    +   [Amendments and corrections](#amendments-and-corrections)
*   [Step parameters](#step-parameters)
*   [Allergies](#allergies)
*   [Locations](#locations)
*   [Appendix](#appendix)
    +   [Messages types and pathway events](#messages-types-and-pathway-events)

This page explains how to write pathways to be run by Simulated Hospital.

## Overview

A pathway is a sequence of events that happen to patients in a hospital.
Simulated Hospital generates patients that follow pathways. Pathway definitions
can be written in YAML or JSON format.

At startup, Simulated Hospital loads pathways from the directory specified by
the [`pathways_dir`](./arguments.md#pathways) command-line argument. All files
inside that directory are read and loaded. Each file can contain one or more
pathways. All pathways defined in a single file must use the same format (either
YAML or JSON, but not a combination of both). All pathways need to be valid
pathways; if not, Simulated Hospital won't start.

The examples in this guide use YAML for brevity.

## Sections of a pathway

A pathway consists of the following parts:

```yaml
pathway_name:             # The unique name of this pathway.
  persons:                # A description of the person or persons the pathway relates to.
    ...
  consultant:             # The consultant to use whenever it is needed in the pathway.
    ...
  percentage_of_patients: # The percentage of patients the pathway relates to.
    ...
  historical_data:        # Events that happened in the past.
    ...
  pathway:                # Actual events in the pathway.
    ...
```

All sections are optional, but either `historical_data` or `pathway`, or both,
must be present.

Events are also referred to as *steps*.

### Persons

The `persons` section configures the person or persons the pathway refers to.
This data is used to construct PID segments.

The `persons` section is optional. If the persons section is not set, the data
about the patient is generated randomly. This applies to individual fields too:
if a pathway only specifies a subset of fields, the rest of them will be
generated randomly. The [how to configure data](./arguments.md#configure-data)
page explains how data about persons is generated and how you can change it.

See an example of a `persons` section with a single patient `main_patient` that
sets the MRN, first name, surname, age and postcode:

```yaml
sample_pathway:
  persons:
    main_patient:
      mrn: "291847192817319"
      first_name: "Elizabeth"
      surname: "Smith"
      age:
        from: 20
        to: 25
      address:
        postcode: "SW1 4DN"
  ...
```

A persons section can refer to multiple patients. This is useful for pathways
where two medical records interact, such as pathways that do bed swaps or
patient merges. See an example of a `persons` section with two patients:

```yaml
sample_pathway:
  persons:
    main_patient:
      first_name: "Main Patient"
    unconscious_patient:
      first_name: "Unconscious"
  ...
```

The steps in the [Pathway](#pathway) section can refer to each of these persons
by their identifiers (`main_patient` or `unconscious_patient`). See
[Use Patient](#use-patient) for an example.

The following list contains the fields that can be configured in a pathway.
Remember that if a field is not specified, it is generated randomly.

*   `nhs`:
    [NHS number (National Health System number)](https://www.nhs.uk/using-the-nhs/about-the-nhs/what-is-an-nhs-number/).
    Simulated Hospital generates valid NHS numbers only, but there is no
    validation when the NHS number is specified directly in a pathway.
*   `mrn`: Medical Record Number (MRN). By default this is a random 32-bit
    integer. See [how to extend Simulated Hospital](./extend-sh.md) to learn how
    to configure this.
*   `age`: the age of the patient in years. You can specify a range using the
    `from` and `to` attributes, and the specific day in a year the person was
    born in the `day_of_year` attribute.
*   `date_of_birth`: the specific date of birth for the patient. Only one of
    `age` or `date_of_birth` can be set.
*   `gender`: the gender of the person.
*   `first_name`: the first name of the person. Names and surnames allow
    patients to be easily found when searching for examples of specific
    situations, and therefore many of the pathways provided by default set the
    first name of the patients to something that indicates the pathway the
    patient is in, e.g. *AKI Scenario 1*.
*   `surname`: the surname of the person.
*   `address` has the following subfields:
    *   `first_line`
    *   `second_line`
    *   `city`
    *   `postcode`
    *   `country`: the country code, e.g. GB
    *   `type` is the type of address, e.g. HOME or WORK
    *   `all_random` is a "true" or "false" value that indicates whether to
        populate all fields with randomly generated values.

The following fields have special generation rules:

*   Name. When the name not provided, it is chosen based on the gender and the
    date of birth.
*   Address. Every field in the address can be overridden separately, except
    `second_line`. If `first_line` is set, `second_line` will be re-generated.

The following demographics are always generated randomly:

*   Middle name is generated based on gender. It can be empty.
*   Ethnicity: by default it is based on the ethnicity frequency in London from
    [wikipedia](https://en.wikipedia.org/wiki/Ethnic_groups_in_London).
*   Patient Class: based on the patient class and type distribution and modified
    according to the steps, i.e if the patient is admitted their class is set to
    “INPATIENT”.
*   Prefix: based on gender.
*   Suffix and Degree can be empty.
*   Telephone number.

### Consultant

The `consultant` section configures the consultant that will be used whenever a
consultant is needed in the pathway. This data is used to populate information
about clinicians in EVN, OBR, OBX, PV1, DG1 and PR1 segments.

Diagnoses and procedures always use an arbitrary consultant to indicate that
past diagnoses and procedures are generally done by other clinicians.

The `consultant` section is optional. If the consultant section is not set, the
consultant is randomly chosen from the set of doctors defined in the
[`doctors_file` argument]./arguments.md#data-configuration]. To use an existing
predefined doctor, set the `id` field only, for example:

```yaml
consultant:
  id: C12345678
```

The identifier _"C12345678"_ must match an pre-defined doctor.

To use a doctor that is not part of the pre-loaded pool of doctors, set all
fields, for example:

```yaml
consultant:
  id: C12345678
  prefix: Dr.
  first_name: John
  surname: Smith
...
```

This new doctor will be added to the pool of doctors, and pathways can now refer
to it by the `id` only in the current run of Simulated Hospital. Simulated
Hospital sets the specialty from the `hospital_service` field specified in the
[HL7 message configuration file](./arguments.md#data-configuration).

### Percentage of patients

The `percentage_of_patients` is a floating number that represents the percentage
of patients the given pathway relates to. This number will determine the
likelihood of the pathway being run in comparison to other pathways. The greater
the value, the more often this pathway will be run.

The `percentage_of_patients` values for all pathways should sum up to 100. If
the pathway does not have `percentage_of_patients` specified, its
`percentage_of_patients` will be assigned from the remaining budget, ie: if the
subset of the pathways have the `percentage_of_patients` specified and it sums
up to 70%, the remaining 30% would be split equally across all remaining
pathways.

There following flags can be set to control the percentage calculation:

*   `max_significant_digits` - significant digits accepted for
    `percentage_of_patients` (default: 3)
*   `strict_percentage_validation` - whether to fail if the sum of
    `percentage_of_patients` differs from 100 more than `percentage_tolerance`
    (default: false)
*   `percentage_tolerance` - How much the accumulated `percentage_of_patients`
    can deviate from 100 (default: 10)
*   `default_percentage` - Default `percentage_of_patients` for pathways that
    don't specify it (default: 1). Note: this is only used if it is impossible
    to calculate the percentage by splitting the remaining budget across
    pathways that do not have the percentage specified.

### Historical Data

The `historical_data` section contains the list of events that happened in the
past and might be relevant to the patient's current pathway, for instance, it
may contain historical test results in the patient's medical history. These
events are performed before the actual pathway starts.

Any supported events, except `Delay` (see [Steps](#steps)), can be included in
this section. All events in this section must have a negative `time_from_now`,
that represents when in the past the event took place, see the following
example:

```yaml
pathway_name:
  historical_data:
    - result:
        order_profile: UREA AND ELECTROLYTES
        results:
          - test_name: Creatinine
            value: 126.00
            unit: UMOLL
            abnormal_flag: HIGH
      parameters:
        time_from_now: -8783h
```

### Pathway

The `pathway` section contains the actual events in the pathway. Any supported
events can be included in this section.

The following is an example of a pathway that performs an admission, a test
result, and a discharge of a patient, with a few delays in between them.

```yaml
aki_scenario_1:
  pathway:
    - admission:
        loc: Renal
    - delay:
        from: 10m
        to: 30m
    - result:
        order_profile: UREA AND ELECTROLYTES
        results:
          - test_name: Creatinine
            value: 254.00
            unit: UMOLL
            abnormal_flag: HIGH
    - delay:
        from: 30m
        to: 60m
    - discharge: {}
```

## Steps

Simulated Hospital supports multiple step types that refer to different events
at a hospital, for instance, admissions, discharges, bed swaps or test results.
Each step type translates to a different HL7 message. For instance, an
`admission` step generates an ADT A01 message.
[The Messages types and pathway events](#messages-types-and-pathway-events)
table connects message types and the steps that generate them.

### Use Patient

The `use_patient` step defines which patient should be used from now on in the
pathway. The patient can be referred to by the person identifier specified in
the Persons section in the pathway, or by their MRN.

```yaml
sample_pathway:
  persons:
    main_patient:
      first_name: "Main Patient"
    unconscious_patient:
      first_name: "Unconscious"
  pathway:
    - use_patient:
        patient: unconscious_patient
```

### Delay

The `delay` step adds a delay between 2 steps. Delay steps represent delays
between events in hospitals and enable building more realistic pathways. Delay
steps don't generate messages.

The duration specified in a Delay step runs in real time. For instance, the
following steps generate a random delay between an admission and a discharge of
30 to 60 minutes:

```yaml
sample_pathway:
  pathway:
    - admission:
        loc: Renal
    - delay:
        from: 30m
        to: 60m
    - discharge: {}
```

When the pathway runs, Simulated Hospital Patients will wait for 30 to 60
minutes before issuing the discharge after the admission.

### Admission

An `admission` step represents a hospital admission and produces an A01 message.
This step type requires the name of the point of care (`loc`) the patient is
being admitted to, for example:

```yaml
admission_only:
  pathway:
    - admission:
        loc: Renal
```

Optionally you can specify an admit reason for the patient. This will add the
admission reason to the PV2.3 Admit Reason field, for example:

```yaml
admission_with_reason:
  pathway:
    - admission:
        loc: Renal
        admit_reason: Kidney problems
```

The point of care must be one of the locations available in the hospital, see
the [Locations section](#locations).

An admission does the following:

*   Generate a location in the given point of care
*   Generate a visit identifier
*   Set the admission date to the current time
*   Set the patient class to *INPATIENT*

### Transfer

A `transfer` step represents a transfer from one location to another and
produces an A02 message. This step requires the name of the point of care
(`loc`) the patient will be transferred to. If the patient is currently
occupying a bed somewhere else, such bed will be marked as free, and a new bed
will be assigned in the new point of care. Example:

```yaml
pathway_with_transfer:
  pathway:
    - admission:
        loc: Non-renal
    - transfer:
        loc: Renal
```

When the pathway runs, the patient is admitted to the `Non-renal` ward and then
transferred to the `Renal` ward.

### Transfer in Error

A `transfer_in_error` step behaves like a [Transfer](#transfer) step but does
not free the current bed or allocate a new bed. This step represents a transfer
that was requested in error and will be cancelled with a
[Cancel Transfer](#cancel-transfer) step. Since the transfer is going to be
cancelled, Simulated Hospital Patients does not need to do any space management.

### Document

A `document` step creates documents such as an autopsy reports, diagnostic
images, discharge summaries and cardiodiagnostics. It produces a MDM^T02 message
with an TXA and an OBX segment.

This step allows customizing the document properties and/or the content.

#### Document properties

This step has the following optional parameters to set the document properties:

*   `document_type` populates the TXA.2-Document Type field. If not set in the
    pathway, the document type is chosen at random from the configurable list in
    the hl7 message configuration file.
*   `completion_status` populates the TXA.17-Document Completion Status field.
    If not set in the pathway, it defaults to *DO* (Documented).
*   `observation_identifier_id` populates the ID of the OBX.3-Observation
    Identifier field. If not set in the pathway, it defaults to *Established
    Patient 15*.
*   `observation_identifier_text` populates the Text of the OBX.3-Observation
    Identifier field. If not set in the pathway, it defaults to *Established
    Patient 15*.
*   `observation_identifier_coding_system` populates the Coding System of the
    OBX.3-Observation Identifier field. If not set in the pathway, it defaults
    to *Simulation*.

Example of a new document:

```yaml
pathway_with_document:
  pathway:
  - document:
      document_type: DS
      completion_status: IP
      observation_identifier_id: obs-id
      observation_identifier_text: obs-text
      observation_identifier_coding_system: coding-system
```

#### Content

The `document` step generates random textual content divided into lines, and
each line of the content becomes an *OBX* segment.

It is possible to set fixed content for the header, ending, or both, in order to
make the document follow the expected structure of a clinical system.

The following optional fields control the document's content:

*   `header_content_lines` sets the first lines in the document if fixed content
    needs to be set.
*   `ending_content_lines` sets the last lines in the document if fixed content
    needs to be set.
*   `num_random_content_lines` specifies the range of content lines using the
    `from` and `to` attributes. If not set in the pathway, Simulated Hospital
    generates a random number of lines between 10 and 50. If you do not want
    random content, set this to 0 or set as empty.

Example of a document with a fixed header and a fixed ending, and random content
in between:

```yaml
  - document:
      header_content_lines:
        - header-text-line-1
        - header-text-line-2
      ending_content_lines:
        - ending-text-line-1
        - ending-text-line-2
```

Example of a document with fixed content only:

```yaml
  - document:
      header_content_lines:
        - content-text-line-1
        - content-text-line-2
        - content-text-line-3
      num_random_content_lines: {}
```

#### Document updates

The `document` step includes an option to update the content lines for an
existing document given a pathway Document ID and an update type to be one of
the following:

*   append
*   overwrite

The following required fields control the document update:

*   `id` is used to link documents within a pathway. This should be set when
    generating a new document if there is intention to update this document.
    This does not set field TXA.12-Document_ID, which is always randomly
    generated.
*   `update_type` sets the type of update to perform.

Example of a document update of type `append`:

```yaml
pathway_with_document:
  pathway:
  - document:
      id: doc-id1
  - document:
      id: doc-id1
      update_type: "append"
      ending_content_lines:
        - ending-text-line-1
```

This pathway generates two messages for the same document. In the first message,
the document will have between 10 and 50 OBX segments with randomly generated
content. The second message will have an additional OBX with the content
"ending-text-line-1" after the previous randomly generated content.

Example of a document update of type `overwrite`:

```yaml
pathway_with_document:
  pathway:
  - document:
      id: doc-id1
  - document:
      id: doc-id1
      update_type: "overwrite"
      header_content_lines:
        - header-text-line-1
      num_content_lines:
        from: 10
        to: 10
```

This pathway generated two messages for the same document. In the first message,
the document will have between 10 and 50 OBX segments with randomly generated
content. The second message will replace these segments with an OBX with the
content "header-text-line-1" followed by 10 OBX segments with randomly generated
content.

### Discharge

A `discharge` step represents a discharge and produces an A03 message. This step
does not have parameters. A discharge step does the following:

*   Set the patient's discharge date
*   Set the patient class to *OUTPATIENT*
*   Mark the patient's bed as available

### Discharge in Error

A `discharge_in_error` step behaves as a [Discharge](#discharge) step but does
not mark the patient bed as available. This step represents a discharge that was
issued in error and will be cancelled with a
[Cancel Discharge](#cancel-discharge).

This step does not have parameters.

### Registration

A `registration` registers a patient and produces an A04 message. An optional
`patient_class` field sets a specific patient class, see example:

```yaml
pathway_with_registration:
  pathway:
    - registration:
        patient_class: EMERGENCY
```

If the `patient_class` is set, it will be used for both the Patient Class
(PV1.2) and Patient Type (PV1.18). If no `patient_class` field is set, the class
is generated randomly based on the distribution from the `patient_class.csv`
file.

### Cancel Admit/ Visit

A `cancel_visit` step cancels the latest admission or visit and produces a A11
message.

This step takes no parameters.

A `cancel_visit` step removes the cancelled admission or visit from the
patient's history.

### Cancel Transfer

A `cancel_transfer` cancels a transfer and produces an A12 message.

Use [Transfer in error](#transfer-in-error) before `cancel_transfer` so that the
patient location is not marked as available between the erronous transfer and
the cancellation.

This step takes no parameters.

### Cancel Discharge

A `cancel_discharge` step cancels the latest discharge and sends an A13 message.

Use [Discharge in error](#discharge-in-error) before `cancel_discharge` so that
the patient location is not marked as available between the erronous discharge
and the cancellation.

This step takes no parameters.

### Add Person

An `add_person` creates a person and sends an A28 message.

It uses the person specified in the pathway. It is only allowed as the first
step of the pathway.

### Update Person

An `update_person` updates an existing person. It sends an A08 message (Update
patient information) if the person is an inpatient, and an A31 (Update person
information) if the person is not an inpatient.

`update_person` might contain a valid person as the parameter. If it contains a
person, the current patient’s details will be updated with those of the person.

A person inside an `update_person` can have some fields set to an automatic
random value (for instance, the address and/or the surname), which updates the
patient’s details with new random values. This is particularly useful when we
want those values to be updated but we don’t want to set fixed values for all
the patients in such pathway.

To set a random value for a surname, use the keyword RANDOM. For addresses,
every individual field inside an address can also be set to RANDOM.
Alternatively, you can use `all_random: true` to generate a new random full
address.

If an `update_person` changes or sets the gender, the middlename and prefix will
be generated again only if the new gender is different than the previous one.

An `update_person` can also have a set of diagnoses and procedures specified.

If the code or description (or both) is set to RANDOM, a random diagnosis or
procedure is generated from the list of all valid diagnoses / procedures. The
`time_from_now` can be omitted in this case - it will be randomly set as maximum
1 year from now.

If only code is set for the given diagnosis / procedure, and it matches one of
the diagnosis / procedure loaded, the description is also populated. Same works
the other way round (i.e., if only description is specified in the pathway).

Example:

```yaml
- update_person:
   person:
     first_name: "Updated Person"
     surname: RANDOM
     address:
       all_random: true
   diagnoses:
     - type: Working
       code: A01.0
       description: Typhoid fever
       datetime:
         time_from_now: -48h
   procedures:
     - code: RANDOM
       datetime:
         time_from_now: -24h
```

### Pre Admission

A `pre_admission` generates a pre admission event and an A05 message. A
`pre_admission` requires the following parameters:

*   `loc`: The location where the patient will be admitted
*   `expected_admission_time_from_now`: The duration of time until the admission
    takes place

The HL7 segments of A05 messages are syntactically identical to those of A14
(Pending Admission) but A05 implies that an account should be opened for the
purposes of tests prior to the admission.

### Pending Admission

A `pending_admission` generates a pending admission event and an A14 message. A
`pending_admission` requires the following parameters:

*   `loc`: The location where the patient will be admitted
*   `expected_admission_time_from_now`: The duration of time until the admission
    takes place

### Pending Transfer

A `pending_transfer` generates a pending transfer event and an A15 message. A
`pending_transfer` requires the following parameters:

*   `loc`: The location where the patient will be transferred
*   `expected_transfer_time_from_now`: The duration of time until the transfer
    takes place

### Pending Discharge

A `pending_discharge` generates a pending discharge event and an A16 message. A
`pending_discharge` requires the following parameters:

*   `expected_discharge_time_from_now`: The duration of time until the discharge
    takes place

### Delete Visit

A `delete_visit` deletes a patient visit and sends a A23 message. The
`delete_visit` step is used to delete a previously discharged or canceled visit
record. To cancel an ongoing visit, use `cancel_visit` instead. In the current
implementation, we default to deleting the most recent discharged or canceled
visit, so no parameters are needed.

### Cancel Pending Admission

A `cancel_pending_admission` cancels a pending admission and generates an A27
message. In the A27 message, PV2-1 field contains the location where the patient
is no longer being admitted to. A `cancel_pending_admission` step must always be
preceded by a `pending_admission` step. No parameters are required.

### Cancel Pending Transfer

A `cancel_pending_transfer` cancels a pending transfer and generates an A26
message. In the A26 message, the PV2-1 field contains the location where the
patient is no longer being transferred to. A `cancel_pending_transfer` step must
always be preceded by a `pending_transfer` step. No parameters are required.

### Cancel Pending Discharge

A `cancel_pending_discharge` cancels a pending discharge and generates an A25
message. A `cancel_pending_discharge` step must always be preceded by a
`pending_discharge` step. No parameters are required.

### Order

An `order` event places an order and generates an ORM message.

An `order` event has the following optional fields:

*   `order_id`: a unique identifier for the order within the pathway. Order
    identifiers are used to link `order` and `result` events within the same
    pathway, so that Simulated Hospital knows that those events relate do the
    same order.
*   `order_profile`: the type of order which populates the _"OBR.4 - Universal
    Service Identifier"_ field. This can be either an order profile from the
    pre-loaded [order profiles](#order-profiles) or a different value.
*   `order_status`: the status of the order to set in the _"OBR.5 - Order
    Status"_. If not specified, Simulated Hospital uses the
    `order_status.in_process` value from the
    [HL7 config](./arguments.md#hl7-config).

At least one of `order_id` or `order_profile` needs to be present.

Simulated Hospital generates a Placer number (_"ORC.4 - Placer Group Number"_)
for each new order.

#### Order Acknowledgement

After an `order` step generates an ORM message, by default it also generates an
ORR message 1-10 seconds afterwards. The ORR (O02) message acknowledges that the
ORM message was received.

The MSA segment of the ORR message contains the message control ID from the ORM
message, connecting the message to the order it acknowledges.

The ORC segment of the ORR message is the same as in the ORM message, with the
exception that the Order Control field is set to *OK*.

If you don't want an order acknowledgement to be sent after a particular order,
on the `order` step you can set the boolean parameter
`no_acknowledgement_message` to `true` and the ORR message won’t be generated.

### Results

A `results` step generates a set of results and an ORU message. Given an order
profile name and the list of results, the full list of results for each test
type for the [Order Profile](#order-profiles) is generated. For each test type:

*   Value and unit - if specified in the pathway, they are used, otherwise
    generate random value from the normal range for test type. Use `EMPTY` to
    force an empty value.
*   Value type - if result value is numerical, set it to "NM", otherwise to
    "TX". If the value is set to `EMPTY`, the value type defaults to the value
    specified in the Order Profile, or is set to an empty string if there is no
    matching Order Profile.
*   Reference range (`reference_range`) - if specified in the pathway, it is
    used, otherwise it is set to default reference range for the test type

This data is used to construct OBX segment.

Also, in this step the Filler number is generated and populated in the ORC and
OBR segment, on top of Placer number, which is already generated in the Order
step.

The times associated with a result:

*   Collected
*   Received in Lab
*   Reported

are all included within the ORUs. The max time between each of them is specified
by values set in the configuration. However the time differences are bounded by
the delay between when the ORM and the ORU were fired.

In order to model some of the behaviour we find in messages received from
partners, we have made it possible to specify blank times in the pathway for
`collected` and/or `received_in_lab` as follows:

```yaml
collected: EMPTY
received_in_lab: EMPTY
```

A `results` event generates a HL7 message of type `ORU^R01` by default. Use the
`trigger_event` field to generate an ORU message with a different trigger event:
`R01` (default if empty), `R03` or `R32`, all case insensitive. Any other value
will be considered invalid.

Evern item in the `results` field creates a new OBX segment. You can specify
notes for each of the results, which will become NTE segments associated to the
OBX segment. Example:

```yaml
    - result:
        order_id: creat1
        order_profile: UREA AND ELECTROLYTES
        results:
          - test_name: Creatinine
            value: 500.00
            unit: UMOLL
            reference_range: 49 - 92
            abnormal_flag: DEFAULT
            notes:
              - "First note for the Creatinine"
              - "Second note for the Creatinine"
```

#### Midnight Case

In order to recreate the midnight case (a collected time of 00:00), set the
following:

```yaml
collected: MIDNIGHT
```

### Merge

A `merge` merges two or more patients.

This step requires you to specify in advance one (or more) valid MRNs to merge
the current patient with, or you can define and use multiple Persons in the
pathway and use their identifiers.

The Merge message will be of type `ADT^A34` if there is only one MRN or
identifier in the `children` field, and of type`ADT^A40` if there are more or if
the (optional) field `force_a40`is set to `true`. Example:

```yaml
merge:
  parent: CURRENT
  children: [1234]
  force_a40: true
```

### Bed swap

A `bed_swap` performs a bed swap between two patients. It produces A17 message.

This step requires you to specify in advance one valid MRN or patient identifier
to bed swap the current patient with. You can define and use multiple Persons in
the pathway and use their identifiers.

Example of a bed swap:

```yaml
bed_swap:
  patient_1: CURRENT
  patient_2: 1234
```

Note that only one MRN can be set here, as opposed to Merge steps that allow a
list of MRNs.

### Track Departure / Track Arrival

Track Departure and Track Arrival events represent changes in a patient's
physical location (inpatient or outpatient) that do not change the official
census bed location. See http://www.hl7.eu/refactored/msgADT_A09.html and
http://www.hl7.eu/refactored/msgADT_A10.html.

A `track_departure` event generates an A09 message. The event is triggered when
there is about to be a change in the patient’s location, but an official A02
transfer hasn’t been issued. The patient could be leaving the floor or the
building, but must stay within the same healthcare institution. There are three
modes in which this event can occur. For two of these modes, the event must be
followed by a `track_arrival` event that completes the move, otherwise the
patient location will remain in an inconsistent state.

1.  *transit* mode: the patient is departing to a non-temporary location and
    there is a gap between departure and arrival times. Use the optional field
    `destination_bed` to specify the specific bed the patient is departing to.
    Add a `track_arrival` event with `transit` mode to complete the move. The
    `loc` field must be the same for both events.

1.  *temporary* mode: the patient is moving to a temporary location (e.g.,
    Hallway, O/R, X-RAY, Limbo). Add a `track_arrival` event with `temporary`
    mode to complete the move.

    NOTE: multiple `track_departure` events in this mode can be issued before a
    `track_arrival` event, e.g. the patient could be sent to Hallway and then
    X-ray before enough time has passed to issue a `track_arrival` event for the
    Hallway location.

1.  *track* mode: patient is departing to a new non-temporary location. A
    `track_arrival` event is not needed after this. Use the optional field
    `destination_bed` to set the specific bed the patient is departing to.

A `track_arrival` event generates an A10 message. There are multiple modes of
track arrival events that are similar to the `track_departure` events:

1.  *transit* mode: the patient has arrived to a non-temporary location. This
    event must follow a `track_departure` event with `transit` mode.

1.  *temporary* mode: the patient has arrived from a temporary location. The
    location the patient has arrived to can be temporary or permanent. If the
    location of arrival is temporary, set the `is_temporary` field. This field
    can only be set in `temporary` mode. If the location is permanent, use the
    optional field `bed` to set the specific bed the patient has arrived at.

1.  *track* mode: the patient has arrived to a new non-temporary location. Use
    the optional field `bed` to set the specific bed the patient has arrived at.

See the following examples:

Example of transit mode:

```yaml
- track_departure:
   mode: transit
   destination_loc: Ward 1
- track_arrival:
   mode: transit
   loc: Ward 1
```

Example of temporary mode:

```yaml
- track_departure:
   mode: temporary
   destination_loc: Hallway
- track_departure:
   mode: temporary
   destination_loc: X-RAY
- track_arrival:
   mode: temporary
   loc: X-RAY
   is_temporary: true
- track_departure:
   mode: temporary
   destination_loc: O/R
- track_arrival:
   mode: temporary
   loc: Ward 1
```

Example of track mode:

```yaml
- track_departure:
   mode: track
   destination_loc: Ward 2
- track_arrival:
   mode: track
   loc: Ward 1
```

### AutoGenerate

An `autogenerate` event inserts one or more Result steps into the pathway within
a specified time-frame and at a specific interval. Neither the position of, nor
Delays specified for this step have an effect on when the Result events are
generated.

This step is useful to easily generate multiple results for a patient, for
instance, vital signs.

Example:

```yaml
- autogenerate:
    result:
      order_profile: Vital Signs
    from: -50h
    to: 5h
    every: 2h
```

*   `from` and `to` define the timeframe of the generated events. Both are
    inclusive and are absolute times starting from when the pathway start. If
    `from` is negative, results will be inserted into the `historical_data`
    sectio: e.g., in the above example, the first result message will be
    generated when time is -50h and thus belongs to the `historical_data`
    section. If `to` and `from` are equal, a single result will be inserted.

*   `every` is the interval at which results will be inserted. If `every` is
    larger than the interval between `from` and `to`, only one result will be
    inserted.

    The following autogenerate step generates two results, one at time 0h and
    one at time 10h:

    ```yaml
    - autogenerate:
        from: 0h
        to: 10h
        every: 10h
    ```

    The following autogenerate step generates a single result at time 0h:

    ```yaml
    - autogenerate:
        from: 0h
        to: 10h
        every: 11h
    ```

*   `result` is the Result step to be inserted into the pathway. Any parameters
    (e.g., `order_profile`) will be copied to the generated steps and will
    produce the same behaviour they would have in a regular `result` step.

### Clinical Note

A `clinical_note` event sends a HL7 message containing a document with
information about a patient, which clinicians refer to as a *clinical note*.
Clinical notes include discharge notes, images, etc., and can have multiple
formats.

The resulting HL7 message is an `ORU^R01` message and thus is similar to the
message generated by a Results event. The message contains a single result, with
the content of the document in the `OBX-5-ObservationValue` field. The content
will be encoded in base64 if the document is pdf, jpg or png.

A `clinical_note` requires the `content_type` field to be set and it must be one
of `txt`, `rtf`, `pdf`, `jpg` and `png`. If `content_type` is `txt`, the field
`document_content` can contain the fixed value for the note.

The following event will generate an `ORU^R01` message with an embedded `pdf`
document. The content will be one of the predefined pdf documents from the
directory linked by the argument
[`--sample_notes_directory`](./arguments.md#configure-data)).

```yaml
- clinical_note:
    content_type: pdf
```

The following event will generate an `ORU^R01` message with a document of type
*Discharge notes* and fixed text:

```yaml
- clinical_note:
    content_type: txt
    document_content: "My super secret notes"
    document_type: "Discharge notes"
```

If an `document_id` is specified, an addendum can be sent on top of the existing note.
The following sends a clinical note with a pdf.

```yaml
- clinical_note:
    document_id: my-id
    content_type: pdf
    document_type: "Discharge notes"
    document_title: "Patient Overview"
```

Using the same `document_id = my-id` to send another clinical note will append a new
document to the existing document. Specifying a new `document_title` and
`document_type` in the new pathway will override the previous values. The
following adds a new OBX to the previous message with the new document content
and updates the `document_title` and `document_type`.

```yaml
- clinical_note:
    document_id: my-id
    content_type: txt
    document_type: "Updated Discharge notes"
    document_title: "Updated Patient Overview"
```

### Hardcoded message

A `hardcoded_message` event sends a pre-loaded message from the folder
configured by the
[`hardcoded_messages_dir` argument](./arguments.md#data-configuration). This
step requires a `regex` field with a comma separated list of regular expressions
that match one or more message names from the pre-loaded set.

When Simulated Hospital runs a `hardcoded_message` step, it matches the names of
the pre-loaded messages with the regular expression. If the regular expression
matches no messages, the step results in an error. If the regular expression
matches one or more messages, Simulated Hospital picks and sends a random
message from the matching set.

For instance, the following step sends a random pre-loaded hardcoded message
that starts with the string "Invalid":

```yaml
- hardcoded_message:
    regex: Invalid.*,
```

Use the special regular expression `.*` to match any pre-loaded hardcoded
message.

Pre-loaded hardcoded messages must have a specific format, see
[Data configuration](./arguments.md#data-configuration). When Simulated Hospital
sends a message, it sets the following values:

*   _"MSH.7 - Date/Time Of Message"_: The current timestamp.
*   _"MSH-10 Message Control ID"_: A unique identifier for the message.
*   PID segment (if the message contains the _"PID_SEGMENT_PLACEHOLDER"_
    keyword): The PID segment of the current patient.

### Generic

A `generic` event allows developers to specify situations that aren't covered by
other events, such as experimental code or events not supported yet in Simulated
Hospital.

Generic events do not have a default behavior. Developers need to pair all
generic events with matching
[Override Event Processors](./extend-sh.md#custom-event-and-message-processors).
Generic events without matching processors will fail to run.

Use the `parameters.custom` field to pass custom parameters to the events. Use
the `name` field to distinguish between different types of generic events.

The following example uses two generic events. The interpretation and the
behavior in each event is up to the developer to define, and is sent with event
processors (see [example](./extend-sh.md#example-generic-events)). For instance,
the first generic event could add a medication to the patient's record, and the
second one a diagnosis.

```yaml
- generic:
    name: add_medication
  parameters:
    custom:
      medication_name: paracetamol
- generic:
    name: add_diagnosis
  parameters:
    custom:
      diagnosis_name: appendicitis
```

### Generate Resources

A `generate_resources` event will trigger the generation of the entire patient
record (at the time of the event) for each pathway as FHIR resources.

Currently, the following resources are supported:

-   [`Bundle`](https://www.hl7.org/fhir/bundle.html)
    -   `type`
    -   `entry`
-   [`Patient`](https://www.hl7.org/fhir/patient.html)
    -   `identifier`
    -   `name`
    -   `gender`
    -   `telecom`
    -   `deceased`
    -   `address`
-   [`AllergyIntolerance`](https://www.hl7.org/fhir/allergyintolerance.html)
    -   `type`
    -   `category`
    -   `reaction`
    -   `code`
    -   `patient`
    -   `clinicalStatus`
-   [`Encounter`](https://www.hl7.org/fhir/encounter.html)
    -   `status`
    -   `statusHistory`
    -   `period`
    -   `diagnoses`
    -   `class`
-   [`Observation`](https://www.hl7.org/fhir/observation.html)
    -   `code`
    -   `encounter`
    -   `status`
    -   `value`
    -   `note`
    -   `subject`
-   [`Location`](https://www.hl7.org/fhir/location.html)
    -   `name`
-   [`Procedure`](https://www.hl7.org/fhir/procedure.html)
    -   `code`
    -   `category`
    -   `subject`
    -   `encounter`
    -   `performed`
    -   `performer`
    -   `statusCode`
-   [`Condition`](https://www.hl7.org/fhir/condition.html)
    -   `code`
    -   `subject`
    -   `encounter`
    -   `recordedDate`
    -   `recorder`

## Order profiles

Order profiles define the type of results that are generated. All order profiles
are loaded from a configuration file.

The yaml configuration file has the following format:

```yaml
UREA AND ELECTROLYTES:
 test_types:
   Creatinine:
     id: lpdc-2012
     ref_range: 49 - 92
     unit: UMOLL
     value: '382'
     value_type: NM
   Potassium:
     id: lpdc-2804
     ref_range: 3.5 - 5.1
     unit: MMOLL
     value: '3.6'
     value_type: NM
 universal_service_id: lpdc-3969

Vital Signs:
  test_types:
    AVPU:
      id: tt-0005-01
      value: Alert
      value_type: TX
    BowelMovement:
      id: tt-0005-02
      value: 'Yes'
      value_type: TX

[etc.]
```

In reality, UREA AND ELECTROLYTES order profile consists of more than 5 test
types, but for simplicity let's it contains Creatinine and Potassium only.

When specifying Order or Result step in the pathway, the Order Profile name
needs to be specified. It is used to look up the Order Profile from the
configuration file.

For the Order step, the matching order profile is only used to populate the
correct universal service name and id. For the Result step, it is also used to
populate data like reference ranges, or to generate the normal / abnormal value
for the result. There are multiple ways of how Results may be specified.

### Explicitly specify results for each test type in the order profile (recommended)

It is recommended to be explicit when defining the Result in the pathway, ie: to
include results for each test type in the pathway. Example:

```yaml
result:
 order_profile: UREA AND ELECTROLYTES
 results:
   - test_name: Creatinine
     value: 150
     unit: UML
     reference_range: 145 - 550
   - test_name: Potassium
     value: 4.6
     unit: MMOLL
     reference_range: 3.5 - 5.1
```

This will cause the message with 2 results (2 OBX segments) to be generated,
with exactly the same values as specified in the pathway. Note:

*   `reference_range` is optional. If not specified, the reference range for
    test type from the order profile will be used.
*   `value` may either be numerical or textual value; if the value is numerical,
    eg: 70, 5.5, <0.25, >=3.5 etc., then the Unit must also be set; if the value
    is textual, the Unit cannot be set.

### Only specify results for a subset of test types

If the scenario is focused on one particular test types, e.g. Creatinine for
scenarios that measure Acute Kidney Injury, some test types may be omitted from
the pathway. Note that Results for test types not specified in the pathway will
not be included.

Example:

```yaml
result:
 order_profile: UREA AND ELECTROLYTES
 results:
   - test_name: Creatinine
     value: 700
     unit: UML
     reference_range: 145 - 550
```

This will cause the message to be generated with one Creatinine result only, and
no Potassium result. This may lead to less realistic pathways, as normally all
test types for the order profile are included.

### Explicitly specify result as random

Results may be generated randomly based on the reference range. There are 3
options to generate random values: normal, abnormally high or abnormally low:

```yaml
result:
 order_profile: UREA AND ELECTROLYTES
 results:
   - test_name: Creatinine
     value: 150
     unit: UML
     reference_range: 145 - 550
   - test_name: Potassium
     value: NORMAL

result:
 order_profile: UREA AND ELECTROLYTES
 results:
   - test_name: Creatinine
     value: 150
     unit: UML
     reference_range: 145 - 550
   - test_name: Potassium
     value: ABNORMAL_HIGH

result:
 order_profile: UREA AND ELECTROLYTES
 results:
   - test_name: Creatinine
     value: 150
     unit: UML
     reference_range: 145 - 550
   - test_name: Potassium
     value: ABNORMAL_LOW
```

This will cause the message to be generated with 2 results:

*   Creatinine result with the value specified in the pathway
*   Potassium result with the value randomly generated from the normal, high or
    low ranges (respectively) for this test type

If the value is specified as NORMAL, ABNORMAL_HIGH or ABNORMAL_LOW, then:

*   `unit` must be omitted or also set to NORMAL, ABNORMAL_HIGH or ABNORMAL_LOW.
    However, if the reference range is specified, then the unit must also be
    specified.

*   `reference range` can be specified and a random value will be generated
    based on it. This reference range overrides the reference range from the
    order profile for the current message

### Do not specify results - implicitly random

If no results are specified, the results for each test type will be generated
with a random value from the normal range of the order profile.

Example:

```yaml
result:
  order_profile: UREA AND ELECTROLYTES
```

This will cause the message to be generated with the full set of results from
the Urea and Electrolytes profile: Creatinine and Potassium in this case. All of
them will have a random value from their normal ranges.

### Results for a random order profile

The order profile may be specified as RANDOM, in which case a set of random
(normal) results for the random order profile will be generated.

Example:

```yaml
result:
  order_profile: RANDOM
```

### Results for non-existing order profile

If there is no matching order profile found for the name specified in the
pathway, the results will still be generated as far as the value and unit are
set explicitly. For random values, reference ranges and units are required.

### Results with an unknown name for an existing order profile

If there is a matching order profile for the name specified in the pathway, but
the test names do not match the test names in the order profile, the pathway
will be considered invalid.

## Default values for Order Profiles

When generating random values for Test Types for Order Profiles, there is a set
of rules that are used:

*   If the value type is not numerical, the value is always set to exactly the
    same textual value as in the order profile config file
*   If the value is numerical and the reference ranges can be parsed, the value
    is randomly generated from the normal range
*   If the value is numerical but the reference range cannot be parsed, the
    value is always set to the value from the order profile file

Numerical values are of the form: 70, 5.5, >0.25, <=5.5.

### Reference range

Reference ranges is an interval, within which the result is being considered
normal. It can be defined in one of the following format:

```
from-to
from - to
[ from - to ]
[from-to]
from-to^from^to
```

where both: `from` and `to` are either positive or negative floating point
numbers;

or:

```
<to^^<to
<to^<to
<=to^^<=to
<=to^<=to
[ < to ]
[ <= to ]
[<to]
[<=to]
```

where `to` is either positive or negative floating point number; `from` is not
specified, meaning that the start of the range is open;

or:

```
>from^^>from
>from^>from
>=from^^>=from
>=from^>=from
[ > from ]
[ >= from ]
[>from]
[>=from]
```

where `from` is either positive or negative floating point number; `to` is not
specified, meaning that the end of the range is open.

When generating a normal random value from the reference range, the value is
always strictly within the range. Similar when generating the abnormal high or
abnormal low values - they are always strictly above or below the reference
range respectively.

When generating a normal value from half-open reference range, we always make
sure the generated value has a correct sign and is of a reasonable order of
magnitute. Eg:

If the reference range is right open, ie: `>from`, the normal value will be
generated using the following rules:

*   if `from` is positive, the normal value is generated from (`from`, `10 x
    from`).
*   if `from` is negative, the normal value is generated from (`from`, `0`) to
    only allow negative numbers.
*   if `from` is `0`, the normal value is generated from (`from`, `10`); this is
    an arbitrary choice, as we don't really know what order of magnitude the
    values should be in.

If the reference range is left open, ie: `<to`, the normal value will be
generated using the following rules:

*   if `to` is positive, the normal value is generated from (`0`, `to`) to only
    allow positive values.
*   if `to` is negative, the normal value is generated from (`10 x to`, `to`).
*   if `to` is 0, the normal value is generated from (`-10`, `to`); this is an
    arbitrary choice, as we don't really know what order of magnitude the values
    should be in.

Similar rules apply when generating abnormal high / abnormal low values from any
reference range. We always make sure the generated number is of the reasonable
order of magnitite and has the same sign as the high / low edge or the range,
respecivelly.

It is not possile to generate abnormal high value for right open reference
range, or abnormal low value for left open reference range. Attempting to do so
will result in an error.

## Pathway with multiple Orders and Results

If the pathway contains multiple orders and results, each order and result must
have the `order_id` set, to pair them together. The Result will use the same
data (placer number, order date, etc.) as the corresponding Order.

## Pathway with Result but no Order

If the Result is specified but no order with the same `order_id` was specified
in the pathway, then the simulator will assume that it belongs to a different
order and thus all Order specific data (placer number, order date, etc.) will be
generated when creating the result.

## Pathways with multiple Results with the same order_id

The first result with the `order_id` is by default treated as final result. All
consequent results are then treated as corrections. The result status can
alternatively be overridden in the pathway.

The following steps generate two messages that relate to the same order, each
message with one result. The result in the second message will be marked as a
correction of the result in the first message:

```yaml
   - result:
        order_id: 123
        order_profile: CRP
        results:
        - test_name: Serum C-Reactive Protein
          value: 40
          unit: MGL
          abnormal_flag: HIGH
    - result:
        order_id: 123
        order_profile: CRP
        results:
        - test_name: Serum C-Reactive Protein
          value: 44
          unit: MGL
          abnormal_flag: HIGH
```

### Amendments and corrections

Every new message with results for the same order has incremental numbers in the
OBX.SetID fields. In the example above, the *OBX.1 - Set ID* field for the two
messages will be 1 and 2, respectively.

If you need two steps related to the same order to keep the same SetID (for
instance, some downstream systems might require the same SetID in order to
process corrections or amendments) set the `expect_correction` field of a
`result`. This field indicates that we expect a correction or amendment for the
same order.

The next message for the same order will be treated as the amendment, and the
SetIDs in both the initial message and the amendment will start from the same
number. Note that downstream systems might require the results in the amendment
to be specified in the same order as the original message. Whis can be achieved
by specifying the results explicitly in the same order in both events. If
results are generated randomly, the order is not guaranteed.

If the downstream processing systems treat preliminary values for results in a
different way, use `expect_correction` together with the appropriate value for
`order_status` and `result_status`. Set `order_status` and `result_status` in
the last message to set a value different from the default value for a
correction. Example:

```yaml
   - result:
        order_id: 123
        order_profile: CRP
        expect_correction: true
        order_status: P
        results_status: P
        results:
        - test_name: Serum C-Reactive Protein
          value: 40
          unit: MGL
          abnormal_flag: HIGH
    - result:
        order_id: 123
        order_profile: CRP
        order_status: F
        results_status: F
        results:
        - test_name: Serum C-Reactive Protein
          value: 44
          unit: MGL
          abnormal_flag: HIGH
```

The parameters of the resulting messages for this example are the following:

*   First message:

    *   OBR.25 (Result Status): P
    *   OBX.1 (Set ID) = 1
    *   OBX.11 (Observation Result Status) = P

*   Second message:

    *   OBR.25 (Result Status): F
    *   OBX.1 (Set ID) = 1
    *   OBX.11 (Observation Result Status) = F

Subsequent messages for the same order will start with Set ID 2.

## Step parameters

Each step can contain a `parameters` field with the following parameters:

*   `delay_message`: the delay between when the event happened, and when the HL7
    message should be sent. This can be used to simulate delays in the hospital
    systems or messages out of order.
*   `time_from_now`: the time offset between now and when the event happened.
    This is mostly used for historical data with a negative value.
*   `death_indicator`: indication of death status change: practice is to use “N”
    or “” to declare undead, and “Y” or “DECEASED” to declare dead. If death
    declared, time_since_death must be specified.
*   `time_since_death`: (positive) time offset from the moment of death until
    now.
*   `sending_application`: the sending application (MSH.3) to set in the message
    related to this event.
*   `sending_facility`: the sending facility (MSH.4) to set in the message
    related to this event.
*   `receiving_application`: the receiving application (MSH.5) to set in the
    message related to this event.
*   `receiving_facility`: the receiving facility (MSH.6) to set in the message
    related to this event.
*   `custom`: a map of strings to strings, which can be used to pass arbitrary
    values for custom processing, see
    [Custom event and message processors](#custom-event-and-message-processors).

Example:

```yaml
sample_pathway:
  pathway:
    - admission:
        loc: "Ward"
      parameters:
        delay_message:
          from: 1s
          to: 60m
        sending_application: MyEHR
        custom:
          my_arbitrary_field: my_arbitrary_value
```

## Allergies

A list of allergies can be specified in the following steps: `update_person`,
`admission`, `discharge`, `registration`, `pre_admission`, `discharge_in_error`,
and `add_person`. The Allergies can be also explicitly set to an empty list to
indicate that the patient does not have allergies. If the `allergies` section is
not set, allergies are randomly generated based on the probability that a
patient will have allergies. This probability, together with the maximum
allegies to generate per patient, are specified in the `data.yml` configuration
file.

Allergies are included as AL1 segments in the corresponding ADT message.

Example:

```yaml
allergies:
- type: MEDICATION
  code: J30.1
  description: Allergic rhinitis due to pollen
  severity: MODERATE
  reaction: Skin rash
- type: FOOD
  code: K52.2
  description: Allergic and dietetic gastroenteritis and colitis
  severity: SEVERE
  reaction: Rash
```

## Locations

Simulated Hospital has some default locations defined in `locations.yml`. All
event types for which a location needs to be specified (e.g., an `admission`
event) need to refer to an existing location from that configuration file.

A pathway that refers to an unknown location fails validation. See
[configure data](./arguments.md#data-configuration) for more information.

## Appendix

### Messages types and pathway events

The following table describes what pathway events generate each type of HL7v2
message:

| Message Type | Segments                                    | Pathway Event                 |
| ------------ | ------------------------------------------- | ----------------------------  |
| ADT^A01      | MSH, EVN, PID, PD1, PV1, NK1, AL1           | admission                     |
| ADT^A02      | MSH, EVN, PID, PD1, PV1                     | transfer_in_error             |
| ADT^A03      | MSH, EVN, PID, PD1, PV1, AL1                | discharge, discharge_in_error |
| ADT^A04      | MSH, EVN, PID, PD1, PV1, NK1, AL1           | registration                  |
| ADT^A05      | MSH, EVN, PID, PD1, PV1, PV2, NK1, AL1, DG1 | pre_admission                 |
| ADT^A08      | MSH, EVN, PID, PD1, PV1, AL1, DG1, PR1      | update_person                 |
| ADT^A09      | MSH, EVN, PID, PD1, PV1                     | track_departure               |
| ADT^A10      | MSH, EVN, PID, PD1, PV1                     | track_arrival                 |
| ADT^A11      | MSH, EVN, PID, PD1, PV1                     | cancel_visit                  |
| ADT^A12      | MSH, EVN, PID, PD1, PV1                     | cancel_transfer               |
| ADT^A13      | MSH, EVN, PID, PD1, PV1                     | cancel_discharge              |
| ADT^A14      | MSH, EVN, PID, PD1, PV1, PV2                | pending_admission             |
| ADT^A15      | MSH, EVN, PID, PD1, PV1                     | pending_transfer              |
| ADT^A16      | MSH, EVN, PID, PD1, PV1, PV2                | pending_discharge             |
| ADT^A17      | MSH, EVN, PID, PD1, PV1, PID, PD1, PV1      | bed_swap                      |
| ADT^A23      | MSH, EVN, PID, PV1                          | delete_visit                  |
| ADT^A25      | MSH, EVN, PID, PD1, PV1, PV2                | cancel_pending_discharge      |
| ADT^A26      | MSH, EVN, PID, PD1, PV1, PV2                | cancel_pending_transfer       |
| ADT^A27      | MSH, EVN, PID, PD1, PV1, PV2                | cancel_pending_admission      |
| ADT^A28      | MSH, EVN, PID, PD1, PV1, AL1                | add_person                    |
| ADT^A31      | MSH, EVN, PID, PD1, PV1, AL1, DG1, PR1      | update_person                 |
| ADT^A34      | MSH, EVN, PID, PD1, MRG                     | merge                         |
| ADT^A40      | MSH, EVN, PID, PD1, MRG, PV1                | merge                         |
| MDM^T02      | MSH, EVN, PID, PV1, TXA, OBX                | document                      |
| ORM^O01      | MSH, PID, PV1, ORC, OBR, NTE, OBX, NTE      | order                         |
| ORR^O02      | MSH, MSA, PID, ORC                          | order                         |
| ORU^R01      | MSH, PID, PV1, ORC, OBR, OBX, NTE           | results, clinical_note        |
| ORU^R03      | MSH, PID, PV1, ORC, OBR, OBX, NTE           | results                       |
| ORU^R32      | MSH, PID, PV1, ORC, OBR, OBX, NTE           | results                       |
