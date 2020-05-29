# Extend Simulated Hospital

-   [Overview](#overview)
-   [Custom event and message processors](#custom-event-and-message-processors)
    *   [Example: An event that generates multiple messages](#example-an-event-that-generates-multiple-messages)
    *   [Example: Generic events](#example-generic-events)
-   [Validation functions](#validation-functions)
-   [Item syncers](#item-syncers)
-   [Data generators](#data-generators)
    *   [Identifiers](#identifiers)
    *   [Addresses of patients](#addresses-of-patients)
-   [Arbitrary patient data](#arbitrary-patient-data)

This page explains how you can write source code to extend the functionality of
Simulated Hospital. You will need to import the hospital library and create your
own instance of Simulated Hospital, as done in `simulator.go`.

## Overview

Custom logic can be injected in Simulated Hospital by means of the following
mechanisms:

*   [Custom event and message processors](#custom-event-and-message-processors)
    modify the default event and message processing logic.
*   [Validation functions](#validation-functions) allow to set custom validation
    for pathways.
*   [Item syncers](#item-syncers) sync operations to add and delete events,
    messages and patients.
*   [Data generators](#data-generators) allow to generate custom data for
    patient identifiers and addresses.
*   [Arbitrary patient data](#arbitrary-patient-data) stores arbitrary
    information that is not included in Simulated Hospital.

All of these are set through the `AdditionalConfig` struct that is sent when a
Hospital is created. See below for more details on each.

## Custom event and message processors

When an event is processed, the default logic in `event_types.go` builds the HL7
message related to the event (for event types that generate HL7 messages) and
queues the message. The default message processing logic in `message.go` sends
the message through the configured sender.

The rest of the section describes how to override this default logic, and also
inject logic that runs before and/or after an event or a message is processed.

Write your custom logic in implementations of `EventProcessor` and/or
`MessageProcessor`, and send them in the `Processors` field inside the
`AdditionalConfig` struct. Use the `-Pre`, `-Post` or `-Override` fields
depending on whether you want the logic to run before, after, or instead the
default logic.

Both `EventProcessor` and `MessageProcessor` have a `Matches` method that checks
whether a specific processor matches the current event or message depending on
the type of processor.

The logic that is executed for a certain message or event is as follows:

1.  Run all matching `-Pre` processors in order.
1.  Run all matching `-Override` processors in order. If none of them match, run
    the default processing logic for the event or message.
1.  Run all matching `-Post` processors in order.

If a processor triggers an error, the processing of that item immediately stops.

Some examples of situations where implementing custom logic could be useful:

*   An event processor that runs instead of the custom logic (`EventOverride`)
    and creates multiple related HL7 messages for a single event.
*   An event processor that runs after an event has been processed and
    optionally sends a related HL7 message - for instance, the logic that
    creates an Acknowledgement message after an Order event could be implemented
    with an `EventPost` processor.
*   A message processor that runs after a message has been processed
    (`MessagePost`) and checks that the message has been processed correctly by
    the downstreams systems.
*   A message processor that runs before a message is processed (`MessagePre`)
    and removes sensitive strings.

Use the generic `custom` map inside the `parameters` field of an event type to
send data that can be useful for these processors.

### Example: An event that generates multiple messages

This example walks you through the steps to add support for an event that
generates multiple messages related to a specific order profile - for instance,
some Electronic Health Records send multiple messages related to the same set of
Vital Signs. These messages can be complicated, so for simplicity let's assume
that we want to send two ORU^R01 messages with the same value in the Placer
Order Number (OBR.2 and ORC.2). We also want the Placer Order Number to be
configurable in the YAML file.

We can create an `EventProcessor` that captures when a specific order profile is
being used and runs custom logic instead of the default processing logic. This
can be implemented as follows:

1.  Implement the `Matches` method of the `EventProcessor` interface to trigger
    the custom logic when the event is of type `Result` and it refers to a
    specific order profile.

1.  Implement the `Process` method of the `EventProcessor` interface to build
    and return the HL7 messages. Use the generic `parameters.custom` field of
    the event to send further configuration - in this case, the Placer Order
    Number that we want to set.

1.  Send the custom processor in the `AdditionalConfig.EventOverride` field.

YAML file:

```yaml
pathway_with_custom_processing:
  pathway:
    - results:
        # Our event processor will trigger for this order profile.
        order_profile: target_order_profile
      custom:
        # We will set this Placer Order Number in the generated message.
        placer_order_number: my_placer_order_number
```

Processor:

```go
// myProcessor is a custom event processor that creates two messages with a specific Placer Order Number.
type myProcessor struct{}

// Matches returns whether the given event can be processed by our custom processor.
func (p *myProcessor) Matches(e *state.Event) bool {
    return e.Step.Result != nil && e.Step.Result.OrderProfile == "target_order_profile"
}

// Process generates messages from the given event and patient information.
func (p *myProcessor) Process(e *state.Event, patientInfo *ir.PatientInfo, cfg *processor.Config) ([]*message.HL7Message, error) {
  g := cfg.Generator
  o, err := g.SetResults(nil, e.Step.Result, e.EventTime)
  if err != nil {
      return errors.Wrap(err, "cannot set results in Results event")
  }
  o.Placer = e.Step.Parameters.Custom["placer_order_number"]
  header1 := g.NewHeader(&e.Step)
  msg1, err := message.BuildResultORUR01(header1, patientInfo, o, e.EventTime, e.MessageTime)
  if err != nil {
      return errors.Wrap(err, "cannot build the first ORU^R01 message")
  }
  // We want the two message headers to be different.
  header2 := g.NewHeader(&e.Step)
  msg2, err := message.BuildResultORUR01(header2, patientInfo, o, e.EventTime, e.MessageTime)
  if err != nil {
      return errors.Wrap(err, "cannot build the second ORU^R01 message")
  }
  return []*message.Hl7Message{msg1, msg2}, nil
}
```

Creation of the hospital:

```go
ac := hospital.AdditionalConfig{
  Processors: hospital.Processors{EventOverride: []hospital.EventProcessor{&myProcessor{}}},
}

h, err := hospital.NewHospital(hospital.Config{
   AdditionalConfig:        ac,
   ...[other settings here]...
})
```

### Example: Generic events

[Generic events](./write-pathways.md#generic) allow to inject custom behavior
for situations that do not fit the supported event types in Simulated Hospital.
Every generic step requires a matching Override Event Processor.

Generic events have a `name` field that can be used to distinguish between
different types of generic events in your pathways.

For instance, imagine that you want to have two events to add medications to a
patient's record, and another event to print the medications on the standard
output. This is currently not possible in Simulated Hospital, but you can use
generic events to implement this functionality.

The YAML file could look like this:

```yaml
my_pathway:
  pathway:
    - generic:
        name: add_medication
      parameters:
        custom:
          medication_name: paracetamol
    - generic:
        name: add_medication
      parameters:
        custom:
          medication_name: ibuprofen
    - generic:
        name: print_medications
```

This pathway is valid, but the generic events will fail to process because we
have not specified the logic that needs to run for them. We need to implement
this logic using event processors. In addition to that, there are no fields in
the patient record that can store medications, so we define our own structure
for medication names and use the `AdditionalData` field (see
[Arbitrary patient data](#arbitrary-patient-data)) to store it. Full example:

```go
// addMedicationProc is a custom event processor that adds medications to a patient's medical record.
type addMedicationProc struct{}

// Matches returns whether the given event can be processed by our custom processor.
func (p *addMedicationProc) Matches(e *state.Event) bool {
    return e.Step.Generic != nil && e.Step.Generic.Name == "add_medication"
}

// AdditionalData is the type of *ir.PatientInfo.AdditionalData.
type AdditionalData struct {
    Medications []string
}

// Process adds a medication to the patient's medical record.
func (p *addMedicationProc) Process(e *state.Event, patientInfo *ir.PatientInfo, cfg *processor.Config) ([]*message.HL7Message, error) {
    // Medications aren't part of the regular fields of patientInfo so we use the AdditionalData field.
    newMedication := e.Step.Parameters.Custom["medication_name"]
    var ad AdditionalData
    if patientInfo.AdditionalData != nil {
        ad = patientInfo.AdditionalData.(AdditionalData)
    }

    ad.Medications = append(ad.Medications, newMedication)
    patientInfo.AdditionalData = ad
    return nil, nil
}

// printMedicationsProc is a custom event processor that prints the medications from a patient's medical record.
type printMedicationsProc struct{}

// Matches returns whether the given event can be processed by our custom processor.
func (p *printMedicationsProc) Matches(e *state.Event) bool {
    return e.Step.Generic != nil && e.Step.Generic.Name == "print_medications"
}

// Process prints the medications from the patient's medical record.
func (p *printMedicationsProc) Process(e *state.Event, patientInfo *ir.PatientInfo, cfg *processor.Config) ([]*message.HL7Message, error) {
  ad := patientInfo.AdditionalData.(AdditionalData)
  fmt.Println("Medications:")
  fmt.Println(ad.Medications)
  return nil, nil
}
```

Finally, we send the two event processors when we create the hospital. Event
processors of generic events must be sent in the `EventOverride` field since we
want to override the default empty logic:

```go
ac := hospital.AdditionalConfig{
  Processors: hospital.Processors{EventOverride: []hospital.EventProcessor{&addMedicationProc{}, &printMedicationsProc{}}},
}

h, err := hospital.NewHospital(hospital.Config{
   AdditionalConfig:        ac,
   ...[other settings here]...
})
```

After the above pathway runs, Simulated Hospital will print "\[paracetamol,
ibuprofen\]".

## Validation functions

Simulated Hospital validates the pathways that are loaded at startup and the
pathways that are started through the Control Panel. Basic validation includes,
for instance, checking that all historical steps have a negative
`time_from_now`. By validating and failing early we decrease the probability of
unexpected surprises.

Additional validation functions can be set by means of the
`AdditionalConfig.ValidFn` field. For instance, in the example mentioned in the
previous section, it would be possible to create a validation function that
makes sure that every message with the *target_order_profile* Order Profile also
sets the `custom.placer_order_number` field.

## Item syncers

Information about events, messages and patients is kept in in-memory storage
while Simulated Hospital is running. Events and messages are stored in queues,
and patients in a map. When Simulated Hospital stops, all internal state is
lost.

Item Syncers can be set for each of these items. If an item syncer is set for
one of those data structures, all operations to such structure (e.g., put an
item in a queue, extract an item from a queue, add an item to the patients map)
are forwarded to the item syncers. This allows implementing functionality to
sync the internal data structures with, for instance, a database, and recover
the data in subsequent runs of Simulated Hospital.

## Data generators

Simulated Hospital supports sending custom generators for identifiers and
addresses.

### Identifiers

The identifiers that allow customization are Medical Record Numbers (MRN),
Placer Order Number and Filler Order Number. By default, their value is a random
unsigned 32-bit.

For each of these fields separately you can set your own generators that
implement the following interface:

```go
// IDGenerator is an interface to generate identifiers.
type IDGenerator interface {
    NewID() string
}
```

### Addresses of patients

By default, the addresses of the patients are British addresses that contain
British postcodes. You can create your own generator that implements the
following interface:

```go
// AddressGenerator is an interface to generate addresses.
type AddressGenerator interface {
    Random() *ir.Address
}
```

## Arbitrary patient data

Patient data is stored in `ir.PatientInfo`. If you need to store other data
types that aren't included there, use the `AdditionalData` field. This field is
an interface so you can store any type you want, but you are responsible for
casting it back and forth.

The following code snippet illustrates how to use this field. This is taken from
the [Generic events](#example-generic-events) example above.

We define a struct to store inside AdditionalData.

```go
// AdditionalData is the type of *ir.PatientInfo.AdditionalData.
type AdditionalData struct {
    Medications []string
}
```

We can use the following snippet in a custom events processor to store
medication names:

```go
// Medications aren't part of the regular fields of patientInfo so we use the AdditionalData field.
newMedication := e.Step.Parameters.Custom["medication_name"]
var ad AdditionalData
if patientInfo.AdditionalData != nil {
    ad = patientInfo.AdditionalData.(AdditionalData)
}
ad.Medications = append(ad.Medications, newMedication)
patientInfo.AdditionalData = ad
```

Then, in a different event processor, we can read it and print it like this:

```go
ad := patientInfo.AdditionalData.(AdditionalData)
fmt.Println("Medications:")
fmt.Println(ad.Medications)
```
