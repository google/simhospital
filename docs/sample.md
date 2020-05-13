# Sample output

-   [Sample messages](#sample-messages)

## Sample messages

The following two messages have been generated with the default Simulated
Hospital configuration:

```
MSH|^~\&|SIMHOSP|SFAC|RAPP|RFAC|20200501140643||ORU^R01|1|T|2.3|||AL||44|ASCII
PID|1|2590157853^^^SIMULATOR MRN^MRN|2590157853^^^SIMULATOR
MRN^MRN~2478684691^^^NHSNBR^NHSNMBR||Esterkin^AKI Scenario
6^^^Miss^^CURRENT||19890118000000|F|||170 Juice Place^^London^^RW21
6KC^GBR^HOME||020 5368 1665^HOME|||||||||R^Other - Chinese^^^||||||||
PV1|1|O|ED^^^Simulated
Hospital^^ED^^|28b|||C006^Woolfson^Kathleen^^^Dr^^^DRNBR^PRSNL^^^ORGDR|||MED||||||||||||||||||||||||||||||||||20200501140643||
ORC|RE|1892929505|4262718364||CM||||20200501140643
OBR|1|1892929505|4262718364|us-0003^UREA AND
ELECTROLYTES^WinPath^^||20200501140643|20200501140643|||||||20200501140643||||||||20200501140643|||F||1
OBX|1|NM|tt-0003-01^Creatinine^WinPath^^||98.00|UMOLL|49 -
92|H|||F|||20200501140643|| NTE|0||Task cow administration| NTE|1||Grapefruit
garlic resale camera|

MSH|^~\&|SIMHOSP|SFAC|RAPP|RFAC|20200508130643||ADT^A01|5|T|2.3|||AL||44|ASCII
EVN|A01|20200508130643|||C006^Woolfson^Kathleen^^^Dr^^^DRNBR^PRSNL^^^ORGDR|
PID|1|2590157853^^^SIMULATOR MRN^MRN|2590157853^^^SIMULATOR
MRN^MRN~2478684691^^^NHSNBR^NHSNMBR||Esterkin^AKI Scenario
6^^^Miss^^CURRENT||19890118000000|F|||170 Juice Place^^London^^RW21
6KC^GBR^HOME||020 5368 1665^HOME|||||||||R^Other - Chinese^^^||||||||
PD1|||FAMILY PRACTICE^^12345| PV1|1|I|RenalWard^MainRoom^Bed 1^Simulated
Hospital^^BED^Main
Building^5|28b|||C006^Woolfson^Kathleen^^^Dr^^^DRNBR^PRSNL^^^ORGDR|||MED|||||||||6145914547062969032^^^^visitid||||||||||||||||||||||ARRIVED|||20200508130643||
```

Click [here](./artifacts/messages.out) to download a file with 1013 sample
messages. They have been generated with the following command:

```shell
LOCAL_DIR=<absolute path to your local dir with Simulated Hospital, e.g. /tmp/simulated_hospital>

bazel run //cmd/simulator:simulator -- \
  --local_path=${LOCAL_DIR} \
  --pathways_per_hour=3600 \
  --output=file \
  --output_file=${LOCAL_DIR}/messages.out
```

Have a look at [the command-line arguments](./arguments.md) to learn about the
configuration files and how to make Simulated Hospital generate messages with
different values.
