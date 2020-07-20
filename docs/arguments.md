# Command-line arguments



-   [Message destination](#message-destination)
-   [Resource destination](#resource-destination)
-   [Data configuration](#data-configuration)
-   [Pathways](#pathways)
-   [Tool setup](#tool-setup)
    *   [HL7 config](#hl7-config)
    *   [Dashboard](#dashboard)
    *   [Runtime](#runtime)

Command-line arguments (shortened here to _arguments_) change the default
behavior of Simulated Hospital. This means you can do the following:

*   Run different instances of the app, each with their own data and
    configuration.
*   Change behavior without having to edit source code.

Append arguments to your launch command to change the following:

*   [Message destination](#message-destination)—where messages go once they're
    generated.
*   [Resource destination](#resource-destination)-where resources go once
    they're generated.
*   [Data configuration](#data-configuration) lets you use your own sample data.
*   [Pathways](#pathways) adjust which messages (and how often) Simulated
    Hospital sends.
*   [Tool-setup](#tool-setup) arguments to configure the app and messages.

You can list all the command-line arguments by appending the `help` argument and
running Simulated Hospital. For example:

```shell
$ docker run --rm -it -p 8000:8000 bazel:simhospital_container_image health/simulator \
--help
```

## Message destination

Message destination arguments control where messages go once they're generated,
such as the destination host over a network connection, into a file, or printed
to the console. Add these arguments to your launch command:

`-output` (string)
:   Where the generated HL7 messages go once they're generated. You can use the
    following values:

*   `stdout`: Print the messages to the console.
*   `mllp`: Send the messages over an
    [mllp connection](https://www.hl7.org/implement/standards/product_brief.cfm?product_id=55).
*   `file`: Store the messages in a file.

If not set, Simulated Hospital uses _"stdout"_.

`-output_file` (string)
:   File path to write messages to if `-output=file`. If not set, Simulated
    Hospital uses _"messages.out"_.

You can use `docker` to access the output file and copy it to a local folder,
for example:

```shell
  docker cp simulated_hospital:/health/messages.out .
```

`-mllp_destination` (string)
:   Host:Port to which MLLP messages will be sent; only relevant if
    `-output=mllp`. Since this argument depends on your specific setup,
    Simulated Hospital does not set a default value.

`-mllp_keep_alive` (boolean)
:   Whether to send keep-alive messages on the MLLP connection; only relevant if
    `-output=mllp`. If this is not set, Simulated Hospital does not send
    keep-alive messages.

`-mllp_keep_alive_interval` (duration)
:   Interval between keep-alive messages; only relevant if `-output=mllp` and
    `-mllp_keep_alive=true` (default 1m0s)

Here's an example that sets values for these arguments:

```shell
$ docker run --rm -it -p 8000:8000 bazel:simhospital_container_image health/simulator \
-output mllp \
-mllp_destination 127.0.0.1:6661
```

## Resource destination

Similarly to message destination arguments, resource destination arguments
control where resources go once they're generated.

`-resource_output` (string)
:   Where the generated resources go once they're generated. You can use the
    following values:

*   `stdout`: Print the resources to the console.
*   `file`: Store each resource in a separate file.
*   `cloud`: Send resources to a Cloud FHIR store.

If not set, Simulated Hospital uses _"stdout"_.

`-resource_output_dir` (string)
:   Path to the output directory for resource files; only relevant if
    `-resource_output=file`. Each file stores resources for a single patient at
    a specific point in time. Files are never updated. If not set, Simulated
    Hospital uses _"resources"_.

`-resource_format` (string)
:   The format in which to generate resources. You can use the following values:

*   `json`: Generate resources as JSON.
*   `proto`: Generate resources as text protocol buffers.

If not set, Simulated Hospital uses _"json"_.

The following arguments allow Simulated Hospital to directly populate a Cloud
FHIR store.

`-cloud_project_id` (string)
:   Project ID of the Cloud FHIR store; only relevant if -resource_output=cloud.
    Simulated Hospital does not have a default value.

`-cloud_location` (string)
:   Location of the Cloud FHIR store; only relevant if -resource_output=cloud.
    Simulated Hospital does not have a default value.

`-cloud_dataset` (string)
:   Dataset of the Cloud FHIR store; only relevant if -resource_output=cloud.
    Simulated Hospital does not have a default value.

`-cloud_datastore` (string)
:   Datastore of the Cloud FHIR store; only relevant if -resource_output=cloud.
    Simulated Hospital does not have a default value.

Before using this functionality, you must have:

*   Edit access for a
    [Cloud FHIR store](https://cloud.google.com/healthcare/docs/how-tos/fhir),
    You can create a new one by following the instructions at
    [_"Creating and managing FHIR stores"_](https://cloud.google.com/healthcare/docs/how-tos/fhir#creating_a_fhir_store).
*   Set
    [Application Default Credentials](https://cloud.google.com/sdk/gcloud/reference/auth/application-default)
    for the store. You can do this using `gcloud` with the following:

    ```shell
    $ gcloud auth application-default login
    $ gcloud auth application-default set-quota-project <PROJECT_ID>
    ```

*   A `generate_resources` step in your pathway(s).

Now Simulated Hospital will be able to write to the specified Cloud FHIR store.
You will need to [mount a volume](https://docs.docker.com/storage/volumes/),
mapping the host directory where your `gcloud` credentials are stored so that
Docker can access them. For example:

```shell
$ docker run -v ~/.config/gcloud:/root/.config/gcloud --rm -it -p 8000:8000 bazel:simhospital_container_image health/simulator \
--resource_output=cloud \
--cloud_project_id=<PROJECT_ID> \
--cloud_location=<LOCATION> \
--cloud_dataset=<DATASET_ID> \
--cloud_datastore=<DATASTORE_ID>
```

Note that if invalid arguments are passed, Simulated Hospital will display an
error when attempting to write to the Cloud FHIR store.

## Data configuration

Data configuration arguments allow you to use your own custom clinical,
administrative, and demographic sample data.

If you use a prebuilt Docker image and you want to load your own configuration
files instead of the default ones, you need to inject the files into the
container using Docker volumes.

For instance, run the following command to modify the allergies that are loaded
by default, by copying a local file with allergies to the default path for the
allergies configuration file:

```shell
docker run --rm -it -p 8000:8000 -v ABSOLUTE_PATH_TO_LOCAL_ALLERGIES_FILE:/configs/hl7_messages/allergies.csv $IMAGE health/simulator
```

Alternatively, you can copy the file somewhere else, and use the command line
arguments to point to it:

```shell
docker run --rm -it -p 8000:8000 -v ABSOLUTE_PATH_TO_LOCAL_ALLERGIES_FILE:/configs/allergies.csv $IMAGE health/simulator -allergies_file=configs/allergies.csv
```

Append these arguments to your launch command:

`-allergies_file` (string)
:   Path to a CSV file containing allergies and how often they occur. Simulated
    Hospital generates random allergies for the patients from these values. If
    not set, Simulated Hospital uses _"configs/hl7\_messages/allergies.csv"_.

The columns in the file denote the code, description and relative rate, in this
order. The following file declares two allergies; "Allergy to nuts" with a rate
of 2 will be used twice as much as "Allergy to penicillin":

```txt
    419199007,"Allergy to penicillin",1
    416098002,"Allergy to nuts",2
```

`-boys_names` (string)
:   Path to a CSV file containing historical boys names. Simulated Hospital
    generates patient names from these values based on the patient age. If not
    set, Simulated Hospital uses
    _"configs/hl7\_messages/third\_party/historicname\_tcm77-254032-boys.csv"_.

The format of the file is as follows:

```csv
# This is a comment. Lines starting with "#" are ignored.
RANK,1904,1914,1924
1,WILLIAM,JOHN,JOHN
2,JOHN,WILLIAM,WILLIAM
```

The first line must be a line that starts with RANK followed by comma separated
years in chronological order. The rest of the lines contain the rank number
followed by the names with that rank per given year. In the example above the
first most popular name in 1904 was William, and the 2nd most popular name was
John.

`-clinical_note_types_file` (string)
:   Path to a text file containing the types of Clinical Notes, with one type
    per row. Simulated Hospital assigns values from this file when the type of
    clinical note is not specified in the pathway. If not set, Simulated
    Hospital uses _"configs/hl7\_messages/third\_party/note\_types.txt"_.

`-data_config_file` (string)
:   Path to a YAML file containing the configuration for data to populate HL7
    fields that are not relevant to the use of the HL7 standard. If not set,
    Simulated Hospital uses _"configs/hl7\_messages/data.yml"_.

See the package `config` for the format.

TODO: Write documentation instead of referring to the package.

`-diagnoses_file` (string)
:   Path to a CSV file containing the diagnoses and how often they occur.
    Simulated Hospital generates random diagnoses for patients from these
    values. If not set, Simulated Hospital uses
    _"configs/hl7\_messages/diagnoses.csv"_.

See `allergies_file` for the format.

`-doctors_file` (string)
:   Path to a YAML file containing the doctors. Simulated Hospital assigns
    doctors to patients from these values. If not set, Simulated Hospital uses
    _"configs/hl7\_messages/doctors.yml"_.

Example of a doctor in the file:

```
- id: "C002"
  surname: "Smith"
  firstname: "Elizabeth"
  prefix: "Dr"
  specialty: "SUR"
```

All fields are required.

Pathways can refer to either doctors pre-defined in this file or to
[new doctors](./write-pathways.md#consultant).

`-ethnicity_file` (string)
:   Path to a CSV file containing the ethnicities and how often they occur.
    Simulated Hospital generates ethnicities for patients from these values. If
    not set, Simulated Hospital uses _"configs/hl7\_messages/ethnicity.csv"_.

See `allergies_file` for the format. Add a row with `nil,nil,X` to specify the
proportion of patients with no ethnicity.

`-girls_names` (string)
:   Path to a CSV file containing historical girls names. If not set, Simulated
    Hospital uses
    _"configs/hl7\_messages/third\_party/historicname\_tcm77-254032-girls.csv"_.

See `-boys_names` for the format.

`-hardcoded_messages_dir` (string)
:   Path to a directory containing YAML files with hardcoded messages.
    Developers can create pathways that send hardcoded messages to trigger
    specific conditions that are not supported in pathway events yet. Simulated
    Hospital pre-loads the hardcoded messages from the specified directory. If
    not set, Simulated Hospital uses _"configs/hardcoded\_messages"_.

The hardcoded messages contain a name and a list of segments, see example:

```yaml
InvalidMessageWithInvalidNHSNumber:
  segments:
  - "MSH|^~\\&|SIMHOSP|SFAC|RAPP|RFAC|%s||ADT^A01|%s|T|2.3|||AL||44|ASCII"
  - "PID||20180709162444249^^^^MRN|20180709162444249^^^^MRN~2278322867^^^NHSNMBR||FN20180709162444^Invalid NHS^^^MR^^CURRENT||19610306|M|"
```

Set "%s" in the _"MSH.7 - Date/Time Of Message"_ and _"MSH-10 Message Control
ID"_. When Simulated Hospital sends this message, it replaces the placeholders
with the current time and a unique identifier for the message, respectively.

Optionally, use the keyword _"PID\_SEGMENT\_PLACEHOLDER"_ instead of the PID
segment, for example:

```yaml
InvalidOru_MissingPlacerAndFiller:
  segments:
  - "MSH|^~\\&|SIMHOSP|SFAC|RAPP|RFAC|%s||ORU^R01|%s|T|2.3|||AL||44|ASCII"
  - "PID_SEGMENT_PLACEHOLDER"
  - "PV1|1|INPATIENT|ED^ED^Chair 06^Simulated Hospital^^BED^^2^||||123^dr. J Smith|||161||||19||||INPATIENT|2018070916325657^^^^visitid|||||||||||||||||||||||||20180403000000||||||||"
  - "ORC|RE||||||||20180709163256||||||20180709163256||||"
  - "OBR|1|||us-0003^UREA AND ELECTROLYTES||20180709163256|20180709163256|||||||||456^DR. J.G. BLACK|||||306360202|20180709163256||1|f||1^^^20180709163256^^r|^||"
  - "OBX|1|NM|tt-0003-01^Creatinine^||126|UMOLL|49-92|HH|||F|||||"
```

When Simulated Hospital sends this message in a patient pathway, it replaces the
_"PID\_SEGMENT\_PLACEHOLDER"_ keyword with the PID segment of the patient.

See also the [`hardcoded_message`](./write-pathways.md#hardcoded-message)
pathway step.

`-local_path` (string)
:   Absolute path to the directory where Simulated Hospital is located. This
    path is added as a prefix to all arguments that relate to paths, if they are
    not set explicitly in the command line.

`-locations_file` (string)
:   Path to a YAML file containing the definition of patient locations that are
    supported in the patient pathways. If a pathway refers to a location that is
    not present in the file, the pathway will be considered invalid. If not set,
    Simulated Hospital uses _"configs/hl7\_messages/locations.yml"_.

The location `ED` corresponds to the _Accident & Emergency_ location and is
required. This is the default location where patients are admitted if there's no
previous admission information. If the file does not contain an `ED` location,
Simulated Hospital fails to start.

Example:

```yaml
ED:
  poc: ED
  facility: Simulated Hospital
  building:
  floor: 1
  room: ED Room
  type: ED
```

The `type` is used for the *PL.6 - Person Location Type* field in the *PV1*
segment. If not set, Simulated Hospital uses *BED*.

If you use a custom Location Manager with no `ED` location defined, the default
location provided by Simulated Hospital will have a type of ED, and the rest of
the fields will be left blank.

`-nouns_file` (string)
:   Path to a text file containing English nouns that Simulated Hospital uses to
    generate arbitrary content such as notes or addresses. If not set, Simulated
    Hospital uses _"configs/hl7\_messages/third\_party/nouns.txt"_.

`-order_profile_file` (string)
:   Path to a YAML file containing the definition of order profiles that
    Simulated Hospital uses to generate orders and results. If not set,
    Simulated Hospital uses _"configs/hl7\_messages/order\_profiles.yml"_.

This file has the following format:

```
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

See [Order Profiles](./write-pathways.md#order-profiles) for how to use order
profiles when writing pathways.

`-patient_class_file` (string)
:   Path to a CSV file containing the list of patient classes and types and how
    often they occur. Simulated Hospital generates patient classes and types
    from these values. If not set, Simulated Hospital uses
    _"configs/hl7\_messages/patient\_class.csv"_.

See `allergies_file` for the format.

`-procedures_file` (string)
:   Path to a CSV file containing the list of procedures and how often they
    occur. Simulated Hospital generates random procedures for patients from
    these values. If not set, Simulated Hospital uses
    _"configs/hl7\_messages/procedures.csv"_.

See `allergies_file` for the format.

`-sample_notes_directory` (string)
:   Path to a directory containing sample notes in different formats that
    Simulated Hospital uses to generate clinical notes. If not set, Simulated
    Hospital uses _"configs/hl7\_messages/third\_party/notes"_.

`-surnames_file` (string)
:   Path to a text file containing surnames. Simulated Hospital uses these
    values to generate patient surnames. If not set, Simulated Hospital uses
    _"configs/hl7\_messages/third\_party/surnames.txt"_.

## Pathways

Pathways arguments adjust which messages (and how often) Simulated Hospital
sends.

`-pathway_names` (string)
:   Comma-separated list of pathway names for pathways to run. This argument
    controls the pathways that Simulated Hospital starts automatically. All
    pathways from `-pathways_dir` can be started from the
    [dashboard](./dashboard.md).

*   If `-pathway_manager_type=deterministic`, this argument must contain at
    least one element, and pathways are run in the same order as specified in
    this argument. The list wraps after all of the pathways are run.
*   If `-pathway_manager_type=distribution`, this argument can include regular
    expressions, or be empty - if empty, all pathways are included.

`-pathways_dir` (string)
:   Path to a directory containing YAML or JSON files with definitions of
    pathways. Simulated Hospital pre-loads all these pathways at startup, and
    generates patients that follow those pathways.
    [Write pathways](./write-pathways) describes how to write valid pathway
    definitions. If not set, Simulated Hospital uses _"configs/pathways"_.

`-pathway_manager_type` (string)
:   The way pathways are picked to be run. You can use the following values:

*   `distribution`: Pathways are run randomly, based on their distribution
    specified in the
    [`percentage_of_patients` section](./write-pathways.md#percentage-of-patients).
*   `deterministic`: Pathways are run in a deterministic order specified by the
    `-pathway_names` argument.

`-pathways_per_hour` (float)
:   Number of pathways that start per hour. If not set, Simulated Hospital uses
    `1`.

`-max_pathways` (int)
:   Number of pathways to run before stopping. Pathways run from the dashboard
    do not count towards this limit. If negative or not set, Simulated Hospital
    will keep running pathways indefinitely.

`-exclude_pathway_names` (string)
:   Comma-separated list of pathway names, or regular expressions that match
    pathway names, for the pathways to exclude from running. Pathways that match
    both `-pathway_names` and `-exclude_pathway_names` are excluded. This
    argument controls the pathways that Simulated Hospital starts automatically.
    This argument is used only if `-pathway_manager_type=distribution`. All
    pathways from `-pathways_dir` can be started from the
    [dashboard](./dashboard.md).

When using the *distribution* pathway manager type, you can also configure what
pathways run or not with the
[`percentage_of_patients` section](./write-pathways.md#percentage-of-patients)
in the pathway's YAML definition.

## Tool setup

Tool setup arguments configure the app and messages.

### HL7 config

To change how Simulated Hospital adds common metadata to your HL7 messages, add
these arguments to your launch command:

`-hl7_timezone` (string)
:   The name of a system timezone. The HL7 messages that Simulated Hospital
    generates will contain dates and times in this timezone. You need to check
    the timezone file exists in the system before using. If not set, Simulated
    hospital uses _"UTC"_.

`-header_config_file` (string)
:   Path to a YAML file containing values for the HL7 message header. For
    example, it includes the names of the sending and receiving applications. If
    not set, Simulated Hospital uses _"configs/hl7\_messages/header.yml"_.

`-hl7_config_file` (string)
:   Path to a YAML file containing codings for HL7 messages. If not set,
    Simulated hospital uses _"configs/hl7\_messages/hl7.yml"_.

Here's an example that sets values for these arguments:

```shell
$ docker run --rm -it -p 8000:8000 bazel:simhospital_container_image health/simulator \
-hl7_timezone Europe/Paris -header_config_file configs/hospital-1/header.yml \
-hl7_config_file configs/hospital-1/hl7.yml
```

### Dashboard

By default, the Simulated Hospital dashboard starts on
http://localhost:8000/simulated-hospital/. To change the TCP port or the URL
path, add one (or both) of the following command arguments:

`-dashboard_address` (string)
:   Address for the dashboard to control Simulated Hospital. If not set,
    Simulated Hospital uses _":8000"_.

`-dashboard_uri` (string)
:   Base URI at which the dashboard and endpoints are available. If not set,
    Simulated Hospital uses _"simulated-hospital"_.

For example, to serve Simulated Hospital from the URL
`http://localhost:8080/hospital-1/`, launch it with the following command:

```shell
$ docker run --rm -it -p 8000:8000 bazel:simhospital_container_image health/simulator \
-dashboard_address :8080 -dashboard_uri hospital-1
```

To use a different directory for CSS and JavaScript resources, use the
`static_dir` argument:

`-static_dir` (string)
:   Directory for static assets. If not set, Simulated Hospital uses _"static"_.

For example:

```shell
$ docker run --rm -it -p 8000:8000 bazel:simhospital_container_image health/simulator \
--static_dir site-resources/hospital-1
```

### Runtime

To change the runtime behavior of Simulated Hospital, add these arguments to
your launch command:

`-log_level` (string)
:   Set the logging granularity to _PANIC_, _FATAL_, _ERROR_, _WARN_, _INFO_, or
    _DEBUG_. If you don't set a level, Simulated hospital uses _"INFO"_.

`-metrics_listen_address` (string)
:   The TCP port that Simulated Hospital uses to serve metrics. To learn how to
    monitor with Prometheus, see [Monitor Simulated Hospital](./monitor). If you
    don't set a port, Simulated hospital uses _":9095"_.

`-sleep_for` (duration)
:   How long Simulated Hospital sleeps for before checking if any new messages
    need to be generated. You can use one of these time units: _ns_, _us_, _ms_,
    _s_, _m_, or _h_. For example, _"1.5s"_ or _"1500ms"_. If you don't set a
    duration, Simulated hospital uses _"1s"_.

`-delete_patients_from_memory` (boolean)
:   Whether Simulated Hospital deletes patients after their pathways finish.
    Deleting saves memory but means you can't reuse the patient in another
    pathway. If you don't set this, Simulated Hospital keeps patients in memory.

If you need to handle many patients at the same time and you want your patients
to be available for future pathways, consider implementing an
[Item Syncer](./extend-sh.md#item-syncers).

Here's an example that shows values for these arguments:

```shell
$ docker run --rm -it -p 8000:8000 bazel:simhospital_container_image health/simulator \
-log_level ERROR -metrics_listen_address :9096 \
-sleep_for 2s -delete_patients_from_memory true
```
