[Main H&S page.](README.md)

# Kaleido Testing

Create binaries:

```
GOOS=linux GOARCH=amd64 go build -o [chaincode-name].bin
```
Create the Risk Assessment:

```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "Maintenance Technician",
    "channel": "default-channel",
    "chaincode": "health_and_safety_smart_contract"
  },
  "func": "CreateRiskAssessment",
  "args": [],
  "init": false
}
```
Response body:

```json
{
  "headers": {
    "id": "a320a337-aa1a-4dbf-71c6-a7744b9c6959",
    "type": "TransactionSuccess",
    "timeReceived": "2023-08-11T05:29:29.623818554Z",
    "timeElapsed": 0.431078281,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 586,
  "signerMSP": "u0mk35n6yh",
  "signer": "Maintenance Technician",
  "transactionID": "1cebd899bf97e98060e0a33e11fe2e27348c8623a3037b93f9e99286b5fdfa5f",
  "status": "VALID"
}
```

The created transactions carry the following data:

```json
{
  "docType": "risk-assessment",
  "risk_assessment_id": "risk-assessment-ID-5d765099c0",
  "risk_assessment_date": "2023-08-11",
  "risk_assessment_date_next_review": "2024-02-11",
  "risk_assessment_activity": "not-set",
  "risk_assessment_assessed_by": "not-set",
  "risk_assessment_approved_by": "not-set"
}
```
Create the Risk Assessment with the list of four hazards:

```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "General Manager",
    "channel": "default-channel",
    "chaincode": "health_and_safety_smart_contract"
  },
  "func": "CreateRiskAssessment",
  "args": [
    "[\"one\"]",
    "[\"two\"]",
    "[\"three\"]",
    "[\"four\"]",
    "[\"five\"]",
    "[\"six\"]",
    "[\"seven\"]",
    "[\"eight\"]",
    "[\"nine\"]",
    "[\"ten\"]",
    "[\"eleven\"]",
    "[\"twelve\"]",
    "[\"thirteen\"]",
    "[\"fourteen\"]",
    "[\"fifteen\"]"
  ],
  "init": false
}
```

Response bode:

```json
{
  "headers": {
    "id": "4124463a-a4c9-4d90-6b10-72438300d79e",
    "type": "TransactionSuccess",
    "timeReceived": "2023-08-15T00:58:57.412932028Z",
    "timeElapsed": 0.424045898,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 600,
  "signerMSP": "u0mk35n6yh",
  "signer": "General Manager",
  "transactionID": "ae86da6d2118d94564fc32c75d1af5341cbe82e543fc91e7f372d3801645d8cf",
  "status": "VALID"
}
```

The created record on the ledger is shown below:

```json
{
  "docType": "risk-assessment",
  "risk_assessment_id": "risk-assessment-ID-f024be98a1",
  "risk_assessment_date": "2023-08-15",
  "risk_assessment_date_next_review": "2024-02-15",
  "risk_assessment_activity": "not-set",
  "risk_assessment_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
  "risk_assessment_assessed_by": "not-set",
  "risk_assessment_approved_by": "not-set",
  "risk_assessment_hazard_list_one": ["one"],
  "risk_assessment_hazard_list_two": ["two"],
  "risk_assessment_hazard_list_three": ["three"],
  "risk_assessment_hazard_list_four": ["four"],
  "risk_assessment_hazard_list_five": ["five"],
  "risk_assessment_hazard_list_six": ["six"],
  "risk_assessment_hazard_list_seven": ["seven"],
  "risk_assessment_hazard_list_eight": ["eight"],
  "risk_assessment_hazard_list_nine": ["nine"],
  "risk_assessment_hazard_list_ten": ["ten"],
  "risk_assessment_hazard_list_eleven": ["eleven"],
  "risk_assessment_hazard_list_twelve": ["twelve"],
  "risk_assessment_hazard_list_thirteen": ["thirteen"],
  "risk_assessment_hazard_list_fourteen": ["fourteen"],
  "risk_assessment_hazard_list_fifteen": ["fifteen"]
}
```

The next record will be storing the data as per the list on the hazard list.
The function that creates a transaction on the leger will be as follows:

```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "General Manager",
    "channel": "default-channel",
    "chaincode": "health_and_safety_smart_contract"
  },
  "func": "CreateRiskAssessment",
  "args": [
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]"
  ],
  "init": false
}
```

Response body:

```json
{
  "docType": "risk-assessment",
  "risk_assessment_id": "risk-assessment-ID-f77677e0c9",
  "risk_assessment_date": "2023-08-15",
  "risk_assessment_date_next_review": "2024-02-15",
  "risk_assessment_activity": "not-set",
  "risk_assessment_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
  "risk_assessment_assessed_by": "not-set",
  "risk_assessment_approved_by": "not-set",
  "risk_assessment_hazard_list_one": [
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD"
  ],
  "risk_assessment_hazard_list_two": [
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD"
  ],
  "risk_assessment_hazard_list_three": [
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD"
  ],
  "risk_assessment_hazard_list_four": [
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD"
  ],
  "risk_assessment_hazard_list_five": [
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD"
  ],
  "risk_assessment_hazard_list_six": [
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD"
  ],
  "risk_assessment_hazard_list_seven": [
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD"
  ],
  "risk_assessment_hazard_list_eight": [
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD"
  ],
  "risk_assessment_hazard_list_nine": [
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD"
  ],
  "risk_assessment_hazard_list_ten": [
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD"
  ],
  "risk_assessment_hazard_list_eleven": [
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD"
  ],
  "risk_assessment_hazard_list_twelve": [
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD"
  ],
  "risk_assessment_hazard_list_thirteen": [
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD"
  ],
  "risk_assessment_hazard_list_fourteen": [
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD"
  ],
  "risk_assessment_hazard_list_fifteen": [
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD",
    "TBD"
  ]
}
```

With pre-defined hazard list one (the first row)

```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "General Manager",
    "channel": "default-channel",
    "chaincode": "health_and_safety_smart_contract"
  },
  "func": "CreateRiskAssessment",
  "args": [
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]",
    "[\"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\", \"TBD\"]"
  ],
  "init": false
}
```

Response body:

```json
{
  "headers": {
    "id": "91585813-1bf2-44dd-4e79-08b675f81284",
    "type": "TransactionSuccess",
    "timeReceived": "2023-08-15T01:28:02.127129485Z",
    "timeElapsed": 0.475153813,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 606,
  "signerMSP": "u0mk35n6yh",
  "signer": "General Manager",
  "transactionID": "33617c2e7c726b9e7d51809cca849b4c65ab4cb476cb35ee0d5a35bf66d62a5e",
  "status": "VALID"
}
```

How it is stored on the ledger:

```json

```

**Typical **record for creating the new** risk assessment `POST /transaction`:**

```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "General Manager",
    "channel": "default-channel",
    "chaincode": "health_and_safety_smart_contract"
  },
  "func": "CreateRiskAssessment",
  "args": [],
  "init": false
}
```

The stored data:

```json
{
  "docType": "risk-assessment",
  "risk_assessment_id": "RISK-ASSESSMENT-ID-622077bd3b",
  "risk_assessment_date": "2023-08-15",
  "risk_assessment_date_next_review": "2024-02-15",
  "risk_assessment_activity": "not-set",
  "risk_assessment_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
  "risk_assessment_assessed_by": "not-set",
  "risk_assessment_approved_by": "not-set",
  "risk_assessment_hazard_list_one": [
    "TBD1.1Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus sapien urna, congue a volutpat vel, efficitur pellentesque dui. Aliquam in.",
    "EMP, CON",
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla metus nibh, pharetra vel ipsum a, ornare dictum sem. Nullam eget hendrerit tortor, a tempor turpis. In sit amet nisl posuere, rhoncus lorem eu, mollis leo. Nulla ut vestibulum ligula. Fusce pretium elit ut orci rutrum, pellentesque ultrices lectus ultricies. Aliquam suscipit dictum orci eget condimentum. Aliquam iaculis, arcu nec mollis aliquet, velit leo eleifend dolor, id aliquam enim tellus at libero. Donec lobortis leo vitae tellus elementum, eleifend rhoncus quam luctus. Donec eu quam vehicula, euismod felis fermentum, egestas dui. In nisi orci, dictum id elit iaculis, sodales sagittis metus. Sed ut tellus fermentum, luctus dolor ac, maximus lectus. Duis viverra auctor lorem, sit amet finibus orci lobortis ut. Integer in nisl velit. Proin lacinia nulla sit amet tortor ullamcorper luctus.",
    "2",
    "None.",
    "2",
    "n/a",
    "n/a"
  ],
  "risk_assessment_hazard_list_two": [
    "TBD2.1",
    "TBD2.2",
    "TBD2.3",
    "TBD2.4",
    "TBD2.5",
    "TBD2.6",
    "TBD2.7",
    "TBD2.8"
  ],
  "risk_assessment_hazard_list_three": [
    "TBD3.1",
    "TBD3.2",
    "TBD3.3",
    "TBD3.4",
    "TBD3.5",
    "TBD3.6",
    "TBD3.7",
    "TBD3.8"
  ],
  "risk_assessment_hazard_list_four": [
    "TBD4.1",
    "TBD4.2",
    "TBD4.3",
    "TBD4.4",
    "TBD4.5",
    "TBD4.6",
    "TBD4.7",
    "TBD4.8"
  ],
  "risk_assessment_hazard_list_five": [
    "TBD5.1",
    "TBD5.2",
    "TBD5.3",
    "TBD5.4",
    "TBD5.5",
    "TBD5.6",
    "TBD5.7",
    "TBD5.8"
  ],
  "risk_assessment_hazard_list_six": [
    "TBD6.1",
    "TBD6.2",
    "TBD6.3",
    "TBD6.4",
    "TBD6.5",
    "TBD6.6",
    "TBD6.7",
    "TBD6.8"
  ],
  "risk_assessment_hazard_list_seven": [
    "TBD7.1",
    "TBD7.2",
    "TBD7.3",
    "TBD7.4",
    "TBD7.5",
    "TBD7.6",
    "TBD7.7",
    "TBD7.8"
  ],
  "risk_assessment_hazard_list_eight": [
    "TBD8.1",
    "TBD8.2",
    "TBD8.3",
    "TBD8.4",
    "TBD8.5",
    "TBD8.6",
    "TBD8.7",
    "TBD8.8"
  ],
  "risk_assessment_hazard_list_nine": [
    "TBD9.1",
    "TBD9.2",
    "TBD9.3",
    "TBD9.4",
    "TBD9.5",
    "TBD9.6",
    "TBD9.7",
    "TBD9.8"
  ],
  "risk_assessment_hazard_list_ten": [
    "TBD10.1",
    "TBD10.2",
    "TBD10.3",
    "TBD10.4",
    "TBD10.5",
    "TBD10.6",
    "TBD10.7",
    "TBD10.8"
  ],
  "risk_assessment_hazard_list_eleven": [
    "TBD11.1",
    "TBD11.2",
    "TBD11.3",
    "TBD11.4",
    "TBD11.5",
    "TBD11.6",
    "TBD11.7",
    "TBD11.8"
  ],
  "risk_assessment_hazard_list_twelve": [
    "TBD12.1",
    "TBD12.2",
    "TBD12.3",
    "TBD12.4",
    "TBD12.5",
    "TBD12.6",
    "TBD12.7",
    "TBD12.8"
  ],
  "risk_assessment_hazard_list_thirteen": [
    "TBD13.1",
    "TBD13.2",
    "TBD13.3",
    "TBD13.4",
    "TBD13.5",
    "TBD13.6",
    "TBD13.7",
    "TBD13.8"
  ],
  "risk_assessment_hazard_list_fourteen": [
    "TBD14.1",
    "TBD14.2",
    "TBD14.3",
    "TBD14.4",
    "TBD14.5",
    "TBD14.6",
    "TBD14.7",
    "TBD14.8"
  ],
  "risk_assessment_hazard_list_fifteen": [
    "TBD15.1",
    "TBD15.2",
    "TBD15.3",
    "TBD15.4",
    "TBD15.5",
    "TBD15.6",
    "TBD15.7",
    "TBD15.8"
  ]
}
```

Read the Risk Assessment, queried by ID `RISK-ASSESSMENT-ID-622077bd3b`.
Response body:

```json
{
  "headers": {
    "channel": "default-channel",
    "timeReceived": "",
    "timeElapsed": 0,
    "requestOffset": "",
    "requestId": ""
  },
  "result": {
    "docType": "risk-assessment",
    "risk_assessment_activity": "not-set",
    "risk_assessment_approved_by": "not-set",
    "risk_assessment_assessed_by": "not-set",
    "risk_assessment_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
    "risk_assessment_date": "2023-08-15",
    "risk_assessment_date_next_review": "2024-02-15",
    "risk_assessment_hazard_list_eight": [
      "TBD8.1",
      "TBD8.2",
      "TBD8.3",
      "TBD8.4",
      "TBD8.5",
      "TBD8.6",
      "TBD8.7",
      "TBD8.8"
    ],
    "risk_assessment_hazard_list_eleven": [
      "TBD11.1",
      "TBD11.2",
      "TBD11.3",
      "TBD11.4",
      "TBD11.5",
      "TBD11.6",
      "TBD11.7",
      "TBD11.8"
    ],
    "risk_assessment_hazard_list_fifteen": [
      "TBD15.1",
      "TBD15.2",
      "TBD15.3",
      "TBD15.4",
      "TBD15.5",
      "TBD15.6",
      "TBD15.7",
      "TBD15.8"
    ],
    "risk_assessment_hazard_list_five": [
      "TBD5.1",
      "TBD5.2",
      "TBD5.3",
      "TBD5.4",
      "TBD5.5",
      "TBD5.6",
      "TBD5.7",
      "TBD5.8"
    ],
    "risk_assessment_hazard_list_four": [
      "TBD4.1",
      "TBD4.2",
      "TBD4.3",
      "TBD4.4",
      "TBD4.5",
      "TBD4.6",
      "TBD4.7",
      "TBD4.8"
    ],
    "risk_assessment_hazard_list_fourteen": [
      "TBD14.1",
      "TBD14.2",
      "TBD14.3",
      "TBD14.4",
      "TBD14.5",
      "TBD14.6",
      "TBD14.7",
      "TBD14.8"
    ],
    "risk_assessment_hazard_list_nine": [
      "TBD9.1",
      "TBD9.2",
      "TBD9.3",
      "TBD9.4",
      "TBD9.5",
      "TBD9.6",
      "TBD9.7",
      "TBD9.8"
    ],
    "risk_assessment_hazard_list_one": [
      "TBD1.1Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus sapien urna, congue a volutpat vel, efficitur pellentesque dui. Aliquam in.",
      "EMP, CON",
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla metus nibh, pharetra vel ipsum a, ornare dictum sem. Nullam eget hendrerit tortor, a tempor turpis. In sit amet nisl posuere, rhoncus lorem eu, mollis leo. Nulla ut vestibulum ligula. Fusce pretium elit ut orci rutrum, pellentesque ultrices lectus ultricies. Aliquam suscipit dictum orci eget condimentum. Aliquam iaculis, arcu nec mollis aliquet, velit leo eleifend dolor, id aliquam enim tellus at libero. Donec lobortis leo vitae tellus elementum, eleifend rhoncus quam luctus. Donec eu quam vehicula, euismod felis fermentum, egestas dui. In nisi orci, dictum id elit iaculis, sodales sagittis metus. Sed ut tellus fermentum, luctus dolor ac, maximus lectus. Duis viverra auctor lorem, sit amet finibus orci lobortis ut. Integer in nisl velit. Proin lacinia nulla sit amet tortor ullamcorper luctus.",
      "2",
      "None.",
      "2",
      "n/a",
      "n/a"
    ],
    "risk_assessment_hazard_list_seven": [
      "TBD7.1",
      "TBD7.2",
      "TBD7.3",
      "TBD7.4",
      "TBD7.5",
      "TBD7.6",
      "TBD7.7",
      "TBD7.8"
    ],
    "risk_assessment_hazard_list_six": [
      "TBD6.1",
      "TBD6.2",
      "TBD6.3",
      "TBD6.4",
      "TBD6.5",
      "TBD6.6",
      "TBD6.7",
      "TBD6.8"
    ],
    "risk_assessment_hazard_list_ten": [
      "TBD10.1",
      "TBD10.2",
      "TBD10.3",
      "TBD10.4",
      "TBD10.5",
      "TBD10.6",
      "TBD10.7",
      "TBD10.8"
    ],
    "risk_assessment_hazard_list_thirteen": [
      "TBD13.1",
      "TBD13.2",
      "TBD13.3",
      "TBD13.4",
      "TBD13.5",
      "TBD13.6",
      "TBD13.7",
      "TBD13.8"
    ],
    "risk_assessment_hazard_list_three": [
      "TBD3.1",
      "TBD3.2",
      "TBD3.3",
      "TBD3.4",
      "TBD3.5",
      "TBD3.6",
      "TBD3.7",
      "TBD3.8"
    ],
    "risk_assessment_hazard_list_twelve": [
      "TBD12.1",
      "TBD12.2",
      "TBD12.3",
      "TBD12.4",
      "TBD12.5",
      "TBD12.6",
      "TBD12.7",
      "TBD12.8"
    ],
    "risk_assessment_hazard_list_two": [
      "TBD2.1",
      "TBD2.2",
      "TBD2.3",
      "TBD2.4",
      "TBD2.5",
      "TBD2.6",
      "TBD2.7",
      "TBD2.8"
    ],
    "risk_assessment_id": "RISK-ASSESSMENT-ID-622077bd3b"
  }
}
```

Delete the Risk Assessment:

```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "General Manager",
    "channel": "default-channel",
    "chaincode": "health_and_safety_smart_contract"
  },
  "func": "DeleteRiskAssessment",
  "args": ["RISK-ASSESSMENT-ID-622077bd3b"],
  "init": false
}
```

Response body:

```json
{
  "headers": {
    "id": "c79df194-be43-49f2-64c5-4382e8479804",
    "type": "TransactionSuccess",
    "timeReceived": "2023-08-15T23:16:09.535459169Z",
    "timeElapsed": 0.317427061,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 617,
  "signerMSP": "u0mk35n6yh",
  "signer": "General Manager",
  "transactionID": "df7bfd1233598b3e03f5aaa09622604e6f66418ee856b8ac241792e8eb066266",
  "status": "VALID"
}
```

Updating the `RiskAssessmentActivity` argument on the newly created risk assessment:

```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "General Manager",
    "channel": "default-channel",
    "chaincode": "health_and_safety_smart_contract"
  },
  "func": "UpdateRiskAssessmentActivity",
  "args": ["risk-assessment-ID-f77677e0c9", "Working on towers"],
  "init": false
}
```

**General Manager ID**

```json
eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==
```

**Maintenance Technician ID**

```json
eDUwOTo6Q049TWFpbnRlbmFuY2UgVGVjaG5pY2lhbixPVT1jbGllbnQ6OkNOPWZhYnJpYy1jYS1zZXJ2ZXI=
```

Query the ledger `POST /query`:

```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "General Manager",
    "channel": "default-channel",
    "chaincode": "health_and_safety_smart_contract"
  },
  "func": "QueryRiskAssessmentsByActivity",
  "args": ["Working on towers."],
  "init": false
}
```

Response body:

```json
{
  "headers": {
    "channel": "default-channel",
    "timeReceived": "",
    "timeElapsed": 0,
    "requestOffset": "",
    "requestId": ""
  },
  "result": [
    {
      "docType": "risk-assessment",
      "risk_assessment_activity": "Working on towers.",
      "risk_assessment_approved_by": "Bob",
      "risk_assessment_assessed_by": "Alice",
      "risk_assessment_created_by": "eDUwOTo6Q049TWFpbnRlbmFuY2UgVGVjaG5pY2lhbixPVT1jbGllbnQ6OkNOPWZhYnJpYy1jYS1zZXJ2ZXI=",
      "risk_assessment_date": "2023-08-16",
      "risk_assessment_date_next_review": "2024-02-16",
      "risk_assessment_hazard_list_eight": [
        "TBD8.1",
        "TBD8.2",
        "TBD8.3",
        "TBD8.4",
        "TBD8.5",
        "TBD8.6",
        "TBD8.7",
        "TBD8.8"
      ],
      "risk_assessment_hazard_list_eleven": [
        "TBD11.1",
        "TBD11.2",
        "TBD11.3",
        "TBD11.4",
        "TBD11.5",
        "TBD11.6",
        "TBD11.7",
        "TBD11.8"
      ],
      "risk_assessment_hazard_list_fifteen": [
        "TBD15.1",
        "TBD15.2",
        "TBD15.3",
        "TBD15.4",
        "TBD15.5",
        "TBD15.6",
        "TBD15.7",
        "TBD15.8"
      ],
      "risk_assessment_hazard_list_five": [
        "TBD5.1",
        "TBD5.2",
        "TBD5.3",
        "TBD5.4",
        "TBD5.5",
        "TBD5.6",
        "TBD5.7",
        "TBD5.8"
      ],
      "risk_assessment_hazard_list_four": [
        "TBD4.1",
        "TBD4.2",
        "TBD4.3",
        "TBD4.4",
        "TBD4.5",
        "TBD4.6",
        "TBD4.7",
        "TBD4.8"
      ],
      "risk_assessment_hazard_list_fourteen": [
        "TBD14.1",
        "TBD14.2",
        "TBD14.3",
        "TBD14.4",
        "TBD14.5",
        "TBD14.6",
        "TBD14.7",
        "TBD14.8"
      ],
      "risk_assessment_hazard_list_nine": [
        "TBD9.1",
        "TBD9.2",
        "TBD9.3",
        "TBD9.4",
        "TBD9.5",
        "TBD9.6",
        "TBD9.7",
        "TBD9.8"
      ],
      "risk_assessment_hazard_list_one": [
        "TBD1.1",
        "TBD1.2",
        "TBD1.3",
        "2",
        "None.",
        "2",
        "n/a",
        "n/a"
      ],
      "risk_assessment_hazard_list_seven": [
        "TBD7.1",
        "TBD7.2",
        "TBD7.3",
        "TBD7.4",
        "TBD7.5",
        "TBD7.6",
        "TBD7.7",
        "TBD7.8"
      ],
      "risk_assessment_hazard_list_six": [
        "TBD6.1",
        "TBD6.2",
        "TBD6.3",
        "TBD6.4",
        "TBD6.5",
        "TBD6.6",
        "TBD6.7",
        "TBD6.8"
      ],
      "risk_assessment_hazard_list_ten": [
        "TBD10.1",
        "TBD10.2",
        "TBD10.3",
        "TBD10.4",
        "TBD10.5",
        "TBD10.6",
        "TBD10.7",
        "TBD10.8"
      ],
      "risk_assessment_hazard_list_thirteen": [
        "TBD13.1",
        "TBD13.2",
        "TBD13.3",
        "TBD13.4",
        "TBD13.5",
        "TBD13.6",
        "TBD13.7",
        "TBD13.8"
      ],
      "risk_assessment_hazard_list_three": [
        "TBD3.1",
        "TBD3.2",
        "TBD3.3",
        "TBD3.4",
        "TBD3.5",
        "TBD3.6",
        "TBD3.7",
        "TBD3.8"
      ],
      "risk_assessment_hazard_list_twelve": [
        "TBD12.1",
        "TBD12.2",
        "TBD12.3",
        "TBD12.4",
        "TBD12.5",
        "TBD12.6",
        "TBD12.7",
        "TBD12.8"
      ],
      "risk_assessment_hazard_list_two": [
        "TBD2.1",
        "TBD2.2",
        "TBD2.3",
        "TBD2.4",
        "TBD2.5",
        "TBD2.6",
        "TBD2.7",
        "TBD2.8"
      ],
      "risk_assessment_id": "RISK-ASSESSMENT-ID-2009fdf235"
    }
  ]
}
```

Query the ledger `POST /query`

```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "General Manager",
    "channel": "default-channel",
    "chaincode": "health_and_safety_smart_contract"
  },
  "func": "QueryRiskAssessments",
  "args": [
    "{\"selector\":{\"risk_assessment_activity\":\"Working on towers.\"}}"
  ],
  "init": false
}
```

Response body:

```json
{
  "headers": {
    "channel": "default-channel",
    "timeReceived": "",
    "timeElapsed": 0,
    "requestOffset": "",
    "requestId": ""
  },
  "result": [
    {
      "docType": "risk-assessment",
      "risk_assessment_activity": "Working on towers.",
      "risk_assessment_approved_by": "Bob",
      "risk_assessment_assessed_by": "Alice",
      "risk_assessment_created_by": "eDUwOTo6Q049TWFpbnRlbmFuY2UgVGVjaG5pY2lhbixPVT1jbGllbnQ6OkNOPWZhYnJpYy1jYS1zZXJ2ZXI=",
      "risk_assessment_date": "2023-08-16",
      "risk_assessment_date_next_review": "2024-02-16",
      "risk_assessment_hazard_list_eight": ["TBD8.1", "TBD8.2", "TBD8.3", "TBD8.4", "TBD8.5", "TBD8.6", "TBD8.7", "TBD8.8"],
      "risk_assessment_hazard_list_eleven": ["TBD11.1","TBD11.2","TBD11.3","TBD11.4","TBD11.5","TBD11.6","TBD11.7","TBD11.8"],
      "risk_assessment_hazard_list_fifteen": ["TBD15.1","TBD15.2","TBD15.3","TBD15.4","TBD15.5","TBD15.6","TBD15.7","TBD15.8"],
      "risk_assessment_hazard_list_five": [
        "TBD5.1",
        "TBD5.2",
        "TBD5.3",
        "TBD5.4",
        "TBD5.5",
        "TBD5.6",
        "TBD5.7",
        "TBD5.8"
      ],
      "risk_assessment_hazard_list_four": [
        "TBD4.1",
        "TBD4.2",
        "TBD4.3",
        "TBD4.4",
        "TBD4.5",
        "TBD4.6",
        "TBD4.7",
        "TBD4.8"
      ],
      "risk_assessment_hazard_list_fourteen": [
        "TBD14.1",
        "TBD14.2",
        "TBD14.3",
        "TBD14.4",
        "TBD14.5",
        "TBD14.6",
        "TBD14.7",
        "TBD14.8"
      ],
      "risk_assessment_hazard_list_nine": [
        "TBD9.1",
        "TBD9.2",
        "TBD9.3",
        "TBD9.4",
        "TBD9.5",
        "TBD9.6",
        "TBD9.7",
        "TBD9.8"
      ],
      "risk_assessment_hazard_list_one": [
        "TBD1.1",
        "TBD1.2",
        "TBD1.3",
        "2",
        "None.",
        "2",
        "n/a",
        "n/a"
      ],
      "risk_assessment_hazard_list_seven": [
        "TBD7.1",
        "TBD7.2",
        "TBD7.3",
        "TBD7.4",
        "TBD7.5",
        "TBD7.6",
        "TBD7.7",
        "TBD7.8"
      ],
      "risk_assessment_hazard_list_six": [
        "TBD6.1",
        "TBD6.2",
        "TBD6.3",
        "TBD6.4",
        "TBD6.5",
        "TBD6.6",
        "TBD6.7",
        "TBD6.8"
      ],
      "risk_assessment_hazard_list_ten": [
        "TBD10.1",
        "TBD10.2",
        "TBD10.3",
        "TBD10.4",
        "TBD10.5",
        "TBD10.6",
        "TBD10.7",
        "TBD10.8"
      ],
      "risk_assessment_hazard_list_thirteen": [
        "TBD13.1",
        "TBD13.2",
        "TBD13.3",
        "TBD13.4",
        "TBD13.5",
        "TBD13.6",
        "TBD13.7",
        "TBD13.8"
      ],
      "risk_assessment_hazard_list_three": [
        "TBD3.1",
        "TBD3.2",
        "TBD3.3",
        "TBD3.4",
        "TBD3.5",
        "TBD3.6",
        "TBD3.7",
        "TBD3.8"
      ],
      "risk_assessment_hazard_list_twelve": [
        "TBD12.1",
        "TBD12.2",
        "TBD12.3",
        "TBD12.4",
        "TBD12.5",
        "TBD12.6",
        "TBD12.7",
        "TBD12.8"
      ],
      "risk_assessment_hazard_list_two": [
        "TBD2.1",
        "TBD2.2",
        "TBD2.3",
        "TBD2.4",
        "TBD2.5",
        "TBD2.6",
        "TBD2.7",
        "TBD2.8"
      ],
      "risk_assessment_id": "RISK-ASSESSMENT-ID-2009fdf235"
    }
  ]
}
```

**Get Risk Assessment History `POST /query`**:

```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "General Manager",
    "channel": "default-channel",
    "chaincode": "health_and_safety_smart_contract"
  },
  "func": "GetRiskAssessmentHistory",
  "args": ["RISK-ASSESSMENT-ID-2009fdf235"],
  "init": false
}
```
Response body:

```json
{
  "headers": {
    "channel": "default-channel",
    "timeReceived": "",
    "timeElapsed": 0,
    "requestOffset": "",
    "requestId": ""
  },
  "result": [
    {
      "isDelete": false,
      "record": {
        "docType": "risk-assessment",
        "risk_assessment_activity": "Working on towers.",
        "risk_assessment_approved_by": "Bob",
        "risk_assessment_assessed_by": "Alice",
        "risk_assessment_created_by": "eDUwOTo6Q049TWFpbnRlbmFuY2UgVGVjaG5pY2lhbixPVT1jbGllbnQ6OkNOPWZhYnJpYy1jYS1zZXJ2ZXI=",
        "risk_assessment_date": "2023-08-16",
        "risk_assessment_date_next_review": "2024-02-16",
        "risk_assessment_hazard_list_eight": [
          "TBD8.1",
          "TBD8.2",
          "TBD8.3",
          "TBD8.4",
          "TBD8.5",
          "TBD8.6",
          "TBD8.7",
          "TBD8.8"
        ],
        "risk_assessment_hazard_list_eleven": [
          "TBD11.1",
          "TBD11.2",
          "TBD11.3",
          "TBD11.4",
          "TBD11.5",
          "TBD11.6",
          "TBD11.7",
          "TBD11.8"
        ],
        "risk_assessment_hazard_list_fifteen": [
          "TBD15.1",
          "TBD15.2",
          "TBD15.3",
          "TBD15.4",
          "TBD15.5",
          "TBD15.6",
          "TBD15.7",
          "TBD15.8"
        ],
        "risk_assessment_hazard_list_five": [
          "TBD5.1",
          "TBD5.2",
          "TBD5.3",
          "TBD5.4",
          "TBD5.5",
          "TBD5.6",
          "TBD5.7",
          "TBD5.8"
        ],
        "risk_assessment_hazard_list_four": [
          "TBD4.1",
          "TBD4.2",
          "TBD4.3",
          "TBD4.4",
          "TBD4.5",
          "TBD4.6",
          "TBD4.7",
          "TBD4.8"
        ],
        "risk_assessment_hazard_list_fourteen": [
          "TBD14.1",
          "TBD14.2",
          "TBD14.3",
          "TBD14.4",
          "TBD14.5",
          "TBD14.6",
          "TBD14.7",
          "TBD14.8"
        ],
        "risk_assessment_hazard_list_nine": [
          "TBD9.1",
          "TBD9.2",
          "TBD9.3",
          "TBD9.4",
          "TBD9.5",
          "TBD9.6",
          "TBD9.7",
          "TBD9.8"
        ],
        "risk_assessment_hazard_list_one": [
          "TBD1.1",
          "TBD1.2",
          "TBD1.3",
          "2",
          "None.",
          "2",
          "n/a",
          "n/a"
        ],
        "risk_assessment_hazard_list_seven": [
          "TBD7.1",
          "TBD7.2",
          "TBD7.3",
          "TBD7.4",
          "TBD7.5",
          "TBD7.6",
          "TBD7.7",
          "TBD7.8"
        ],
        "risk_assessment_hazard_list_six": [
          "TBD6.1",
          "TBD6.2",
          "TBD6.3",
          "TBD6.4",
          "TBD6.5",
          "TBD6.6",
          "TBD6.7",
          "TBD6.8"
        ],
        "risk_assessment_hazard_list_ten": [
          "TBD10.1",
          "TBD10.2",
          "TBD10.3",
          "TBD10.4",
          "TBD10.5",
          "TBD10.6",
          "TBD10.7",
          "TBD10.8"
        ],
        "risk_assessment_hazard_list_thirteen": [
          "TBD13.1",
          "TBD13.2",
          "TBD13.3",
          "TBD13.4",
          "TBD13.5",
          "TBD13.6",
          "TBD13.7",
          "TBD13.8"
        ],
        "risk_assessment_hazard_list_three": [
          "TBD3.1",
          "TBD3.2",
          "TBD3.3",
          "TBD3.4",
          "TBD3.5",
          "TBD3.6",
          "TBD3.7",
          "TBD3.8"
        ],
        "risk_assessment_hazard_list_twelve": [
          "TBD12.1",
          "TBD12.2",
          "TBD12.3",
          "TBD12.4",
          "TBD12.5",
          "TBD12.6",
          "TBD12.7",
          "TBD12.8"
        ],
        "risk_assessment_hazard_list_two": [
          "TBD2.1",
          "TBD2.2",
          "TBD2.3",
          "TBD2.4",
          "TBD2.5",
          "TBD2.6",
          "TBD2.7",
          "TBD2.8"
        ],
        "risk_assessment_id": "RISK-ASSESSMENT-ID-2009fdf235"
      },
      "timestamp": "2023-08-16T15:38:17.240692986Z",
      "txId": "a2c475fcc24ea3a7c751a5c28942833a3df80eeaf09dda3075b59abcfa82fb25"
    },
    {
      "isDelete": false,
      "record": {
        "docType": "risk-assessment",
        "risk_assessment_activity": "Working on towers.",
        "risk_assessment_approved_by": "not-set",
        "risk_assessment_assessed_by": "Alice",
        "risk_assessment_created_by": "eDUwOTo6Q049TWFpbnRlbmFuY2UgVGVjaG5pY2lhbixPVT1jbGllbnQ6OkNOPWZhYnJpYy1jYS1zZXJ2ZXI=",
        "risk_assessment_date": "2023-08-16",
        "risk_assessment_date_next_review": "2024-02-16",
        "risk_assessment_hazard_list_eight": [
          "TBD8.1",
          "TBD8.2",
          "TBD8.3",
          "TBD8.4",
          "TBD8.5",
          "TBD8.6",
          "TBD8.7",
          "TBD8.8"
        ],
        "risk_assessment_hazard_list_eleven": [
          "TBD11.1",
          "TBD11.2",
          "TBD11.3",
          "TBD11.4",
          "TBD11.5",
          "TBD11.6",
          "TBD11.7",
          "TBD11.8"
        ],
        "risk_assessment_hazard_list_fifteen": [
          "TBD15.1",
          "TBD15.2",
          "TBD15.3",
          "TBD15.4",
          "TBD15.5",
          "TBD15.6",
          "TBD15.7",
          "TBD15.8"
        ],
        "risk_assessment_hazard_list_five": [
          "TBD5.1",
          "TBD5.2",
          "TBD5.3",
          "TBD5.4",
          "TBD5.5",
          "TBD5.6",
          "TBD5.7",
          "TBD5.8"
        ],
        "risk_assessment_hazard_list_four": [
          "TBD4.1",
          "TBD4.2",
          "TBD4.3",
          "TBD4.4",
          "TBD4.5",
          "TBD4.6",
          "TBD4.7",
          "TBD4.8"
        ],
        "risk_assessment_hazard_list_fourteen": [
          "TBD14.1",
          "TBD14.2",
          "TBD14.3",
          "TBD14.4",
          "TBD14.5",
          "TBD14.6",
          "TBD14.7",
          "TBD14.8"
        ],
        "risk_assessment_hazard_list_nine": [
          "TBD9.1",
          "TBD9.2",
          "TBD9.3",
          "TBD9.4",
          "TBD9.5",
          "TBD9.6",
          "TBD9.7",
          "TBD9.8"
        ],
        "risk_assessment_hazard_list_one": [
          "TBD1.1",
          "TBD1.2",
          "TBD1.3",
          "2",
          "None.",
          "2",
          "n/a",
          "n/a"
        ],
        "risk_assessment_hazard_list_seven": [
          "TBD7.1",
          "TBD7.2",
          "TBD7.3",
          "TBD7.4",
          "TBD7.5",
          "TBD7.6",
          "TBD7.7",
          "TBD7.8"
        ],
        "risk_assessment_hazard_list_six": [
          "TBD6.1",
          "TBD6.2",
          "TBD6.3",
          "TBD6.4",
          "TBD6.5",
          "TBD6.6",
          "TBD6.7",
          "TBD6.8"
        ],
        "risk_assessment_hazard_list_ten": [
          "TBD10.1",
          "TBD10.2",
          "TBD10.3",
          "TBD10.4",
          "TBD10.5",
          "TBD10.6",
          "TBD10.7",
          "TBD10.8"
        ],
        "risk_assessment_hazard_list_thirteen": [
          "TBD13.1",
          "TBD13.2",
          "TBD13.3",
          "TBD13.4",
          "TBD13.5",
          "TBD13.6",
          "TBD13.7",
          "TBD13.8"
        ],
        "risk_assessment_hazard_list_three": [
          "TBD3.1",
          "TBD3.2",
          "TBD3.3",
          "TBD3.4",
          "TBD3.5",
          "TBD3.6",
          "TBD3.7",
          "TBD3.8"
        ],
        "risk_assessment_hazard_list_twelve": [
          "TBD12.1",
          "TBD12.2",
          "TBD12.3",
          "TBD12.4",
          "TBD12.5",
          "TBD12.6",
          "TBD12.7",
          "TBD12.8"
        ],
        "risk_assessment_hazard_list_two": [
          "TBD2.1",
          "TBD2.2",
          "TBD2.3",
          "TBD2.4",
          "TBD2.5",
          "TBD2.6",
          "TBD2.7",
          "TBD2.8"
        ],
        "risk_assessment_id": "RISK-ASSESSMENT-ID-2009fdf235"
      },
      "timestamp": "2023-08-16T15:17:21.127306345Z",
      "txId": "16f93e9378ed4bfd109ae41fa59898b86fbf733b7441e04202621f1d0f9b12e1"
    },
    {
      "isDelete": false,
      "record": {
        "docType": "risk-assessment",
        "risk_assessment_activity": "Working on towers.",
        "risk_assessment_approved_by": "not-set",
        "risk_assessment_assessed_by": "not-set",
        "risk_assessment_created_by": "eDUwOTo6Q049TWFpbnRlbmFuY2UgVGVjaG5pY2lhbixPVT1jbGllbnQ6OkNOPWZhYnJpYy1jYS1zZXJ2ZXI=",
        "risk_assessment_date": "2023-08-16",
        "risk_assessment_date_next_review": "2024-02-16",
        "risk_assessment_hazard_list_eight": [
          "TBD8.1",
          "TBD8.2",
          "TBD8.3",
          "TBD8.4",
          "TBD8.5",
          "TBD8.6",
          "TBD8.7",
          "TBD8.8"
        ],
        "risk_assessment_hazard_list_eleven": [
          "TBD11.1",
          "TBD11.2",
          "TBD11.3",
          "TBD11.4",
          "TBD11.5",
          "TBD11.6",
          "TBD11.7",
          "TBD11.8"
        ],
        "risk_assessment_hazard_list_fifteen": [
          "TBD15.1",
          "TBD15.2",
          "TBD15.3",
          "TBD15.4",
          "TBD15.5",
          "TBD15.6",
          "TBD15.7",
          "TBD15.8"
        ],
        "risk_assessment_hazard_list_five": [
          "TBD5.1",
          "TBD5.2",
          "TBD5.3",
          "TBD5.4",
          "TBD5.5",
          "TBD5.6",
          "TBD5.7",
          "TBD5.8"
        ],
        "risk_assessment_hazard_list_four": [
          "TBD4.1",
          "TBD4.2",
          "TBD4.3",
          "TBD4.4",
          "TBD4.5",
          "TBD4.6",
          "TBD4.7",
          "TBD4.8"
        ],
        "risk_assessment_hazard_list_fourteen": [
          "TBD14.1",
          "TBD14.2",
          "TBD14.3",
          "TBD14.4",
          "TBD14.5",
          "TBD14.6",
          "TBD14.7",
          "TBD14.8"
        ],
        "risk_assessment_hazard_list_nine": [
          "TBD9.1",
          "TBD9.2",
          "TBD9.3",
          "TBD9.4",
          "TBD9.5",
          "TBD9.6",
          "TBD9.7",
          "TBD9.8"
        ],
        "risk_assessment_hazard_list_one": [
          "TBD1.1",
          "TBD1.2",
          "TBD1.3",
          "2",
          "None.",
          "2",
          "n/a",
          "n/a"
        ],
        "risk_assessment_hazard_list_seven": [
          "TBD7.1",
          "TBD7.2",
          "TBD7.3",
          "TBD7.4",
          "TBD7.5",
          "TBD7.6",
          "TBD7.7",
          "TBD7.8"
        ],
        "risk_assessment_hazard_list_six": [
          "TBD6.1",
          "TBD6.2",
          "TBD6.3",
          "TBD6.4",
          "TBD6.5",
          "TBD6.6",
          "TBD6.7",
          "TBD6.8"
        ],
        "risk_assessment_hazard_list_ten": [
          "TBD10.1",
          "TBD10.2",
          "TBD10.3",
          "TBD10.4",
          "TBD10.5",
          "TBD10.6",
          "TBD10.7",
          "TBD10.8"
        ],
        "risk_assessment_hazard_list_thirteen": [
          "TBD13.1",
          "TBD13.2",
          "TBD13.3",
          "TBD13.4",
          "TBD13.5",
          "TBD13.6",
          "TBD13.7",
          "TBD13.8"
        ],
        "risk_assessment_hazard_list_three": [
          "TBD3.1",
          "TBD3.2",
          "TBD3.3",
          "TBD3.4",
          "TBD3.5",
          "TBD3.6",
          "TBD3.7",
          "TBD3.8"
        ],
        "risk_assessment_hazard_list_twelve": [
          "TBD12.1",
          "TBD12.2",
          "TBD12.3",
          "TBD12.4",
          "TBD12.5",
          "TBD12.6",
          "TBD12.7",
          "TBD12.8"
        ],
        "risk_assessment_hazard_list_two": [
          "TBD2.1",
          "TBD2.2",
          "TBD2.3",
          "TBD2.4",
          "TBD2.5",
          "TBD2.6",
          "TBD2.7",
          "TBD2.8"
        ],
        "risk_assessment_id": "RISK-ASSESSMENT-ID-2009fdf235"
      },
      "timestamp": "2023-08-16T15:16:50.961833461Z",
      "txId": "f16a7967834880ffd382e443600164d12ad3a222d4ef96b704282373019ba13b"
    },
    {
      "isDelete": false,
      "record": {
        "docType": "risk-assessment",
        "risk_assessment_activity": "not-set",
        "risk_assessment_approved_by": "not-set",
        "risk_assessment_assessed_by": "not-set",
        "risk_assessment_created_by": "eDUwOTo6Q049TWFpbnRlbmFuY2UgVGVjaG5pY2lhbixPVT1jbGllbnQ6OkNOPWZhYnJpYy1jYS1zZXJ2ZXI=",
        "risk_assessment_date": "2023-08-16",
        "risk_assessment_date_next_review": "2024-02-16",
        "risk_assessment_hazard_list_eight": [
          "TBD8.1",
          "TBD8.2",
          "TBD8.3",
          "TBD8.4",
          "TBD8.5",
          "TBD8.6",
          "TBD8.7",
          "TBD8.8"
        ],
        "risk_assessment_hazard_list_eleven": [
          "TBD11.1",
          "TBD11.2",
          "TBD11.3",
          "TBD11.4",
          "TBD11.5",
          "TBD11.6",
          "TBD11.7",
          "TBD11.8"
        ],
        "risk_assessment_hazard_list_fifteen": [
          "TBD15.1",
          "TBD15.2",
          "TBD15.3",
          "TBD15.4",
          "TBD15.5",
          "TBD15.6",
          "TBD15.7",
          "TBD15.8"
        ],
        "risk_assessment_hazard_list_five": [
          "TBD5.1",
          "TBD5.2",
          "TBD5.3",
          "TBD5.4",
          "TBD5.5",
          "TBD5.6",
          "TBD5.7",
          "TBD5.8"
        ],
        "risk_assessment_hazard_list_four": [
          "TBD4.1",
          "TBD4.2",
          "TBD4.3",
          "TBD4.4",
          "TBD4.5",
          "TBD4.6",
          "TBD4.7",
          "TBD4.8"
        ],
        "risk_assessment_hazard_list_fourteen": [
          "TBD14.1",
          "TBD14.2",
          "TBD14.3",
          "TBD14.4",
          "TBD14.5",
          "TBD14.6",
          "TBD14.7",
          "TBD14.8"
        ],
        "risk_assessment_hazard_list_nine": [
          "TBD9.1",
          "TBD9.2",
          "TBD9.3",
          "TBD9.4",
          "TBD9.5",
          "TBD9.6",
          "TBD9.7",
          "TBD9.8"
        ],
        "risk_assessment_hazard_list_one": [
          "TBD1.1",
          "TBD1.2",
          "TBD1.3",
          "2",
          "None.",
          "2",
          "n/a",
          "n/a"
        ],
        "risk_assessment_hazard_list_seven": [
          "TBD7.1",
          "TBD7.2",
          "TBD7.3",
          "TBD7.4",
          "TBD7.5",
          "TBD7.6",
          "TBD7.7",
          "TBD7.8"
        ],
        "risk_assessment_hazard_list_six": [
          "TBD6.1",
          "TBD6.2",
          "TBD6.3",
          "TBD6.4",
          "TBD6.5",
          "TBD6.6",
          "TBD6.7",
          "TBD6.8"
        ],
        "risk_assessment_hazard_list_ten": [
          "TBD10.1",
          "TBD10.2",
          "TBD10.3",
          "TBD10.4",
          "TBD10.5",
          "TBD10.6",
          "TBD10.7",
          "TBD10.8"
        ],
        "risk_assessment_hazard_list_thirteen": [
          "TBD13.1",
          "TBD13.2",
          "TBD13.3",
          "TBD13.4",
          "TBD13.5",
          "TBD13.6",
          "TBD13.7",
          "TBD13.8"
        ],
        "risk_assessment_hazard_list_three": [
          "TBD3.1",
          "TBD3.2",
          "TBD3.3",
          "TBD3.4",
          "TBD3.5",
          "TBD3.6",
          "TBD3.7",
          "TBD3.8"
        ],
        "risk_assessment_hazard_list_twelve": [
          "TBD12.1",
          "TBD12.2",
          "TBD12.3",
          "TBD12.4",
          "TBD12.5",
          "TBD12.6",
          "TBD12.7",
          "TBD12.8"
        ],
        "risk_assessment_hazard_list_two": [
          "TBD2.1",
          "TBD2.2",
          "TBD2.3",
          "TBD2.4",
          "TBD2.5",
          "TBD2.6",
          "TBD2.7",
          "TBD2.8"
        ],
        "risk_assessment_id": "RISK-ASSESSMENT-ID-2009fdf235"
      },
      "timestamp": "2023-08-16T15:12:32.42533398Z",
      "txId": "7b05362c5b7ee103199bbd1b8b47575b58907c32b26d25cc1709d5905107cede"
    }
  ]
}
```
Crete Accident/Incident Report:
```json
{
	"headers": {
		"type": "SendTransaction",
		"signer": "General Manager",
		"channel": "default-channel",
		"chaincode": "health_and_safety_smart_contract"
	},
	"func": "CreateAccidentIncidentReport",
	"args": ["test1", "test2", "test3", "[]"],
	"init": false
}
```
Response body:
```json
{
	"docType": "accident-incident-report",
	"accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-1654ab54a5",
	"subject_accident_incident_report": "test1",
	"accident_incident_report_date": "2023-08-16",
	"accident_incident_report_time_start": "2024-02-16 18:05:27",
	"accident_incident_report_time_end": "not-set",
	"accident_incident_report_location": "not-set",
	"accident_incident_report_hs_aspects": "test2",
	"accident_incident_report_classification": "test3",
	"accident_incident_report_description": "not-set",
	"accident_incident_report_immediate_action": [],
	"accident_incident_report_follow_up_action": ["not-set"],
	"accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
	"accident_incident_report_approved_by": "not-set",
	"accident_incident_report_status": "open",
	"accident_incident_report_date_closed": "not-set",
	"accident_incident_report_number_of_people_involved": 0,
	"accident_incident_report_person_injured": ["not-set"],
	"accident_incident_report_witnesses": ["not-set"],
	"accident_incident_report_event_overviews": ["not-set"],
	"accident_incident_report_line_of_communication": ["not-set"]
}
```
Another request for creating an accident/incident report:
```json
{
	"headers": {
		"type": "SendTransaction",
		"signer": "General Manager",
		"channel": "default-channel",
		"chaincode": "health_and_safety_smart_contract"
	},
	"func": "CreateAccidentIncidentReport",
	"args": [
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.", 
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.", 
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.", 
    "[\"Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat\", \"Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.\"]"
    ],
	"init": false
}
```
Response body:
```json
{
	"docType": "accident-incident-report",
	"accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-99e9928497",
	"subject_accident_incident_report": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
	"accident_incident_report_date": "2023-08-16",
	"accident_incident_report_time_start": "2024-02-16 18:11:19",
	"accident_incident_report_time_end": "not-set",
	"accident_incident_report_location": "not-set",
	"accident_incident_report_hs_aspects": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
	"accident_incident_report_classification": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
	"accident_incident_report_description": "not-set",
	"accident_incident_report_immediate_action": [
    "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat", 
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
    ],
	"accident_incident_report_follow_up_action": ["not-set"],
	"accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
	"accident_incident_report_approved_by": "not-set",
	"accident_incident_report_status": "open",
	"accident_incident_report_date_closed": "not-set",
	"accident_incident_report_number_of_people_involved": 0,
	"accident_incident_report_person_injured": ["not-set"],
	"accident_incident_report_witnesses": ["not-set"],
	"accident_incident_report_event_overviews": ["not-set"],
	"accident_incident_report_line_of_communication": ["not-set"]
}
```
Another request for creating an accident/incident report:
```json
{
	"headers": {
		"type": "SendTransaction",
		"signer": "General Manager",
		"channel": "default-channel",
		"chaincode": "health_and_safety_smart_contract"
	},
	"func": "CreateAccidentIncidentReport",
	"args": [
    "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "Location: Lorem ipsum dolor sit amet.", 
    "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.", 
    "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
    "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.", 
    "[\"Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat\", \"Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.\"]"
    ],
	"init": false
}
```
Response body:
```json
{
  "headers": {
    "id": "a85d285a-70aa-418a-70c5-4281cc4c2cad",
    "type": "TransactionSuccess",
    "timeReceived": "2023-08-17T21:41:43.034672534Z",
    "timeElapsed": 0.288646169,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 671,
  "signerMSP": "u0mk35n6yh",
  "signer": "General Manager",
  "transactionID": "567603d03a5db22dbfdb2e47fd8b5a43ffc8a55ba62d2b146f512eacc577ae97",
  "status": "VALID"
}
```
Content of the transaction using `POST /query` endpoint and passing the Accident/Incident Report ID:
```json
{
  "headers": {
    "channel": "default-channel",
    "timeReceived": "",
    "timeElapsed": 0,
    "requestOffset": "",
    "requestId": ""
  },
  "result": {
    "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
    "accident_incident_report_date": "2023-08-17",
    "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
    "accident_incident_report_event_overview": [
      "not-set"
    ],
    "accident_incident_report_follow_up_actions": [
      "not-set"
    ],
    "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
    "accident_incident_report_immediate_action": [
      "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
    ],
    "accident_incident_report_line_of_communication": [
      "not-set"
    ],
    "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
    "accident_incident_report_number_of_people_involved": 0,
    "accident_incident_report_person_injured": [
      "not-set"
    ],
    "accident_incident_report_status": "open",
    "accident_incident_report_time_end": "not-set",
    "accident_incident_report_time_start": "2024-02-17 21:41:43",
    "accident_incident_report_validated": "not-set",
    "accident_incident_report_witnesses": [
      "not-set"
    ],
    "docType": "accident-incident-report",
    "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
  }
}
```
**Updating the Accident/Incident Report**
Updating the accident/incident report with `time-end`
Request body `POST /transactions`:
```json
{
	"headers": {
		"type": "SendTransaction",
		"signer": "General Manager",
		"channel": "default-channel",
		"chaincode": "health_and_safety_smart_contract"
	},
	"func": "UpdateAccidentIncidentReportTimeEnd",
	"args": [
    "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
    "2024-02-17 17:00:00"
    ],
	"init": false
}
```
Content of the transaction using `POST /query` endpoint and passing the Accident/Incident Report ID with function `ReadAccidentIncidentReport`:
```json
{
  "headers": {
    "channel": "default-channel",
    "timeReceived": "",
    "timeElapsed": 0,
    "requestOffset": "",
    "requestId": ""
  },
  "result": {
    "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
    "accident_incident_report_date": "2023-08-17",
    "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
    "accident_incident_report_event_overview": [
      "not-set"
    ],
    "accident_incident_report_follow_up_actions": [
      "not-set"
    ],
    "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
    "accident_incident_report_immediate_action": [
      "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
    ],
    "accident_incident_report_line_of_communication": [
      "not-set"
    ],
    "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
    "accident_incident_report_number_of_people_involved": 0,
    "accident_incident_report_person_injured": [
      "not-set"
    ],
    "accident_incident_report_status": "open",
    "accident_incident_report_time_end": "2024-02-17 17:00:00",
    "accident_incident_report_time_start": "2024-02-17 21:41:43",
    "accident_incident_report_validated": "not-set",
    "accident_incident_report_witnesses": [
      "not-set"
    ],
    "docType": "accident-incident-report",
    "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
  }
}
```
Updating the accident/incident report with `follow-up-actions`
Request body `POST /transactions`:
```json
{
	"headers": {
		"type": "SendTransaction",
		"signer": "General Manager",
		"channel": "default-channel",
		"chaincode": "health_and_safety_smart_contract"
	},
	"func": "UpdateAccidentIncidentReportFollowUpActions",
	"args": [
    "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
    "[\"Action:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat\", \"Action:2 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.\", \"Action:3 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.\", \"Action:4 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.\", \"Action:5 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.\"]"
    ],
	"init": false
}
```
Response body:
```json
{
  "headers": {
    "id": "cd0f959b-2256-4a9c-7e42-6b18981f9484",
    "type": "TransactionSuccess",
    "timeReceived": "2023-08-17T21:54:53.658215197Z",
    "timeElapsed": 0.294977787,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 673,
  "signerMSP": "u0mk35n6yh",
  "signer": "General Manager",
  "transactionID": "d02c2082011703952e2c69c00a8b20ddee8e10bf99834594b54ffd3a89cb3d68",
  "status": "VALID"
}
```

Content of the transaction using `POST /query` endpoint and passing the Accident/Incident Report ID with function `ReadAccidentIncidentReport`:
```json
{
  "headers": {
    "channel": "default-channel",
    "timeReceived": "",
    "timeElapsed": 0,
    "requestOffset": "",
    "requestId": ""
  },
  "result": {
    "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
    "accident_incident_report_date": "2023-08-17",
    "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
    "accident_incident_report_event_overview": [
      "not-set"
    ],
    "accident_incident_report_follow_up_actions": [
      "Action:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Action:2 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:3 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:4 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:5 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
    ],
    "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
    "accident_incident_report_immediate_action": [
      "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
    ],
    "accident_incident_report_line_of_communication": [
      "not-set"
    ],
    "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
    "accident_incident_report_number_of_people_involved": 0,
    "accident_incident_report_person_injured": [
      "not-set"
    ],
    "accident_incident_report_status": "open",
    "accident_incident_report_time_end": "2024-02-17 17:00:00",
    "accident_incident_report_time_start": "2024-02-17 21:41:43",
    "accident_incident_report_validated": "not-set",
    "accident_incident_report_witnesses": [
      "not-set"
    ],
    "docType": "accident-incident-report",
    "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
  }
}
```
Updating the accident/incident report with `validated` and `status` fields set to `name-of-validator` and `closed` respectively.
Request body `POST /transactions`:
```json
{
	"headers": {
		"type": "SendTransaction",
		"signer": "General Manager",
		"channel": "default-channel",
		"chaincode": "health_and_safety_smart_contract"
	},
	"func": "UpdateAccidentIncidentReportValidated",
	"args": [
    "ACCIDENT-INCIDENT-REPORT-ID-20570865f5"
    ],
	"init": false
}
```
Response body:
```json
{
  "headers": {
    "id": "c124dbb7-032e-423f-5bed-700c3126054d",
    "type": "TransactionSuccess",
    "timeReceived": "2023-08-17T22:01:10.906762972Z",
    "timeElapsed": 0.324156187,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 674,
  "signerMSP": "u0mk35n6yh",
  "signer": "General Manager",
  "transactionID": "401de99cb15b56cc41ab22b2311cd4243a8290296ef6d32cde2b765454e4bcb2",
  "status": "VALID"
}
```
Content of the transaction using `POST /query` endpoint and passing the Accident/Incident Report ID with function `ReadAccidentIncidentReport`:
```json
{
  "headers": {
    "channel": "default-channel",
    "timeReceived": "",
    "timeElapsed": 0,
    "requestOffset": "",
    "requestId": ""
  },
  "result": {
    "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
    "accident_incident_report_date": "2023-08-17",
    "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
    "accident_incident_report_event_overview": [
      "not-set"
    ],
    "accident_incident_report_follow_up_actions": [
      "Action:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Action:2 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:3 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:4 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:5 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
    ],
    "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
    "accident_incident_report_immediate_action": [
      "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
    ],
    "accident_incident_report_line_of_communication": [
      "not-set"
    ],
    "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
    "accident_incident_report_number_of_people_involved": 0,
    "accident_incident_report_person_injured": [
      "not-set"
    ],
    "accident_incident_report_status": "closed",
    "accident_incident_report_time_end": "2024-02-17 17:00:00",
    "accident_incident_report_time_start": "2024-02-17 21:41:43",
    "accident_incident_report_validated": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
    "accident_incident_report_witnesses": [
      "not-set"
    ],
    "docType": "accident-incident-report",
    "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
  }
}
```
Updating the accident/incident report with the `number_of_people_involved` field set to `1`.
Request body `POST /transactions`:
```json
{
	"headers": {
		"type": "SendTransaction",
		"signer": "General Manager",
		"channel": "default-channel",
		"chaincode": "health_and_safety_smart_contract"
	},
	"func": "UpdateAccidentIncidentReportNumberOfPeopleInvolved",
	"args": [
    "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
    "1"
    ],
	"init": false
}
```
Response body:
```json
{
  "headers": {
    "id": "2614ae1e-9e61-4a42-562e-e5ca267eec2a",
    "type": "TransactionSuccess",
    "timeReceived": "2023-08-17T22:06:54.705470637Z",
    "timeElapsed": 0.34035394,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 675,
  "signerMSP": "u0mk35n6yh",
  "signer": "General Manager",
  "transactionID": "58162f3eac752790c0db5a44a75bba2652fb4d02983b43a92303a3844effd71e",
  "status": "VALID"
}
```
Content of the transaction using `POST /query` endpoint and passing the Accident/Incident Report ID with function `ReadAccidentIncidentReport`:
```json
{
  "headers": {
    "channel": "default-channel",
    "timeReceived": "",
    "timeElapsed": 0,
    "requestOffset": "",
    "requestId": ""
  },
  "result": {
    "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
    "accident_incident_report_date": "2023-08-17",
    "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
    "accident_incident_report_event_overview": [
      "not-set"
    ],
    "accident_incident_report_follow_up_actions": [
      "Action:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Action:2 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:3 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:4 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:5 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
    ],
    "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
    "accident_incident_report_immediate_action": [
      "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
    ],
    "accident_incident_report_line_of_communication": [
      "not-set"
    ],
    "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
    "accident_incident_report_number_of_people_involved": 1,
    "accident_incident_report_person_injured": [
      "not-set"
    ],
    "accident_incident_report_status": "closed",
    "accident_incident_report_time_end": "2024-02-17 17:00:00",
    "accident_incident_report_time_start": "2024-02-17 21:41:43",
    "accident_incident_report_validated": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
    "accident_incident_report_witnesses": [
      "not-set"
    ],
    "docType": "accident-incident-report",
    "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
  }
}
```
Updating the accident/incident report with the `person_injured` field set to `["John Doe"]`.
Request body `POST /transactions`:
```json
{
	"headers": {
		"type": "SendTransaction",
		"signer": "General Manager",
		"channel": "default-channel",
		"chaincode": "health_and_safety_smart_contract"
	},
	"func": "UpdateAccidentIncidentReportPersonInjured",
	"args": [
    "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
    "[\"John Doe\"]"
    ],
	"init": false
}
```
Response body:
```json
{
  "headers": {
    "id": "91d45119-58ae-4b68-6101-fed1b5b3af20",
    "type": "TransactionSuccess",
    "timeReceived": "2023-08-17T22:11:48.158806473Z",
    "timeElapsed": 0.290648925,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 676,
  "signerMSP": "u0mk35n6yh",
  "signer": "General Manager",
  "transactionID": "aaf5e3058f7f7a14505a724dc5debce1f18c2602c4262fb5f07c1328862300da",
  "status": "VALID"
}
```
Content of the transaction using `POST /query` endpoint and passing the Accident/Incident Report ID with function `ReadAccidentIncidentReport`:
```json
{
  "headers": {
    "channel": "default-channel",
    "timeReceived": "",
    "timeElapsed": 0,
    "requestOffset": "",
    "requestId": ""
  },
  "result": {
    "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
    "accident_incident_report_date": "2023-08-17",
    "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
    "accident_incident_report_event_overview": [
      "not-set"
    ],
    "accident_incident_report_follow_up_actions": [
      "Action:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Action:2 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:3 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:4 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:5 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
    ],
    "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
    "accident_incident_report_immediate_action": [
      "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
    ],
    "accident_incident_report_line_of_communication": [
      "not-set"
    ],
    "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
    "accident_incident_report_number_of_people_involved": 1,
    "accident_incident_report_person_injured": [
      "John Doe"
    ],
    "accident_incident_report_status": "closed",
    "accident_incident_report_time_end": "2024-02-17 17:00:00",
    "accident_incident_report_time_start": "2024-02-17 21:41:43",
    "accident_incident_report_validated": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
    "accident_incident_report_witnesses": [
      "not-set"
    ],
    "docType": "accident-incident-report",
    "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
  }
}
```
Updating the accident/incident report with `witnesess` updating the argumnts with `["Beth A. Foley", "Nena H. Woodcock"]`.
Request body `POST /tranasctions`:
```json
{
	"headers": {
		"type": "SendTransaction",
		"signer": "General Manager",
		"channel": "default-channel",
		"chaincode": "health_and_safety_smart_contract"
	},
	"func": "UpdateAccidentIncidentReportWitnesses",
	"args": [
    "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
    "[\"Beth A. Foley\", \"Nena H. Woodcock\"]"
    ],
	"init": false
}
```
Response body:
```json
{
  "headers": {
    "id": "95fc99b5-ca75-4b91-7a0f-fb5c335f7d4e",
    "type": "TransactionSuccess",
    "timeReceived": "2023-08-17T22:16:20.225694806Z",
    "timeElapsed": 0.29183525,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 677,
  "signerMSP": "u0mk35n6yh",
  "signer": "General Manager",
  "transactionID": "a1e0a1282c83b2d127fc5bd9b3377fd2e54fb3b969447e46c368981f8c0b4983",
  "status": "VALID"
}
```
Content of the transaction using `POST /query` endpoint and passing the Accident/Incident Report ID with function `ReadAccidentIncidentReport`:
```json
{
  "headers": {
    "channel": "default-channel",
    "timeReceived": "",
    "timeElapsed": 0,
    "requestOffset": "",
    "requestId": ""
  },
  "result": {
    "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
    "accident_incident_report_date": "2023-08-17",
    "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
    "accident_incident_report_event_overview": [
      "not-set"
    ],
    "accident_incident_report_follow_up_actions": [
      "Action:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Action:2 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:3 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:4 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:5 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
    ],
    "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
    "accident_incident_report_immediate_action": [
      "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
    ],
    "accident_incident_report_line_of_communication": [
      "not-set"
    ],
    "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
    "accident_incident_report_number_of_people_involved": 1,
    "accident_incident_report_person_injured": [
      "John Doe"
    ],
    "accident_incident_report_status": "closed",
    "accident_incident_report_time_end": "2024-02-17 17:00:00",
    "accident_incident_report_time_start": "2024-02-17 21:41:43",
    "accident_incident_report_validated": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
    "accident_incident_report_witnesses": [
      "Beth A. Foley",
      "Nena H. Woodcock"
    ],
    "docType": "accident-incident-report",
    "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
  }
}
```
Updating the accident/incident report with `event_overview` updating the arguments with `["Event:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat", "Event:2 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat"]`.
Request body `POST /transactions`:
```json
{
	"headers": {
		"type": "SendTransaction",
		"signer": "General Manager",
		"channel": "default-channel",
		"chaincode": "health_and_safety_smart_contract"
	},
	"func": "UpdateAccidentIncidentReportEventOverview",
	"args": [
    "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
    "[\"Event:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat\", \"Event:2 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat\"]"
    ],
	"init": false
}
```
Response body:
```json
{
  "headers": {
    "id": "7f750321-b826-4f7d-4c8d-93cd400fc0c6",
    "type": "TransactionSuccess",
    "timeReceived": "2023-08-17T22:19:13.596701197Z",
    "timeElapsed": 0.287967258,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 678,
  "signerMSP": "u0mk35n6yh",
  "signer": "General Manager",
  "transactionID": "cab95c84a0aff2da02514fe4e7ca46a272471e2bc1282c4430675c97bef12a64",
  "status": "VALID"
}
```
Content of the transaction using `POST /query` endpoint and passing the Accident/Incident Report ID with function `ReadAccidentIncidentReport`:
```json
{
  "headers": {
    "channel": "default-channel",
    "timeReceived": "",
    "timeElapsed": 0,
    "requestOffset": "",
    "requestId": ""
  },
  "result": {
    "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
    "accident_incident_report_date": "2023-08-17",
    "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
    "accident_incident_report_event_overview": [
      "Event:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Event:2 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat"
    ],
    "accident_incident_report_follow_up_actions": [
      "Action:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Action:2 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:3 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:4 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:5 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
    ],
    "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
    "accident_incident_report_immediate_action": [
      "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
    ],
    "accident_incident_report_line_of_communication": [
      "not-set"
    ],
    "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
    "accident_incident_report_number_of_people_involved": 1,
    "accident_incident_report_person_injured": [
      "John Doe"
    ],
    "accident_incident_report_status": "closed",
    "accident_incident_report_time_end": "2024-02-17 17:00:00",
    "accident_incident_report_time_start": "2024-02-17 21:41:43",
    "accident_incident_report_validated": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
    "accident_incident_report_witnesses": [
      "Beth A. Foley",
      "Nena H. Woodcock"
    ],
    "docType": "accident-incident-report",
    "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
  }
}
```
Updating the accident/incident report with `line-of-communication` with `["Line of communication:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat", "Line of communication:2 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat"]`.
Request body `POST /tranasctions`:
```json
{
	"headers": {
		"type": "SendTransaction",
		"signer": "General Manager",
		"channel": "default-channel",
		"chaincode": "health_and_safety_smart_contract"
	},
	"func": "UpdateAccidentIncidentReportLineOfCommunication",
	"args": [
    "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
    "[\"Line of communication:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat\", \"Line of communication:2 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat\"]"
    ],
	"init": false
}
```
Response body:
```json
{
  "headers": {
    "id": "fc6c69f9-f115-48a6-564b-35e65cbcda9d",
    "type": "TransactionSuccess",
    "timeReceived": "2023-08-17T22:23:41.669690666Z",
    "timeElapsed": 0.297019499,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 679,
  "signerMSP": "u0mk35n6yh",
  "signer": "General Manager",
  "transactionID": "c255de0b6cee3e07243f191dba8d6b9d349a34a723a110283d8f9081c5e2a85b",
  "status": "VALID"
}
```
Content of the transaction using `POST /query` endpoint and passing the Accident/Incident Report ID with function `ReadAccidentIncidentReport`:
```json
{
  "headers": {
    "channel": "default-channel",
    "timeReceived": "",
    "timeElapsed": 0,
    "requestOffset": "",
    "requestId": ""
  },
  "result": {
    "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
    "accident_incident_report_date": "2023-08-17",
    "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
    "accident_incident_report_event_overview": [
      "Event:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Event:2 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat"
    ],
    "accident_incident_report_follow_up_actions": [
      "Action:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Action:2 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:3 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:4 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
      "Action:5 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
    ],
    "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
    "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
    "accident_incident_report_immediate_action": [
      "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
    ],
    "accident_incident_report_line_of_communication": [
      "Line of communication:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
      "Line of communication:2 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat"
    ],
    "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
    "accident_incident_report_number_of_people_involved": 1,
    "accident_incident_report_person_injured": [
      "John Doe"
    ],
    "accident_incident_report_status": "closed",
    "accident_incident_report_time_end": "2024-02-17 17:00:00",
    "accident_incident_report_time_start": "2024-02-17 21:41:43",
    "accident_incident_report_validated": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
    "accident_incident_report_witnesses": [
      "Beth A. Foley",
      "Nena H. Woodcock"
    ],
    "docType": "accident-incident-report",
    "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
  }
}
```
Qureying the accident/incident report with `subject` 
Request body `POST /query`:
```json
{
	"headers": {
		"type": "SendTransaction",
		"signer": "General Manager",
		"channel": "default-channel",
		"chaincode": "health_and_safety_smart_contract"
	},
	"func": "QueryIncidentReports",
	"args": [
    "{\"selector\":{\"subject_accident_incident_report\":\"Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.\"}}"
    ],
	"init": false
}
```
Response body and Content of the transaction using `POST /query` endpoint:
```json
{
  "headers": {
    "channel": "default-channel",
    "timeReceived": "",
    "timeElapsed": 0,
    "requestOffset": "",
    "requestId": ""
  },
  "result": [
    {
      "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
      "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
      "accident_incident_report_date": "2023-08-17",
      "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
      "accident_incident_report_event_overview": [
        "Event:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
        "Event:2 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat"
      ],
      "accident_incident_report_follow_up_actions": [
        "Action:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
        "Action:2 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
        "Action:3 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
        "Action:4 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
        "Action:5 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
      ],
      "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
      "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
      "accident_incident_report_immediate_action": [
        "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
      ],
      "accident_incident_report_line_of_communication": [
        "Line of communication:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
        "Line of communication:2 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat"
      ],
      "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
      "accident_incident_report_number_of_people_involved": 1,
      "accident_incident_report_person_injured": [
        "John Doe"
      ],
      "accident_incident_report_status": "closed",
      "accident_incident_report_time_end": "2024-02-17 17:00:00",
      "accident_incident_report_time_start": "2024-02-17 21:41:43",
      "accident_incident_report_validated": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
      "accident_incident_report_witnesses": [
        "Beth A. Foley",
        "Nena H. Woodcock"
      ],
      "docType": "accident-incident-report",
      "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
    }
  ]
}
```
**Get Accident/Incident History `POST /query`**:
```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "General Manager",
    "channel": "default-channel",
    "chaincode": "health_and_safety_smart_contract"
  },
  "func": "GetAccidentIncidentReportHistory",
  "args": ["ACCIDENT-INCIDENT-REPORT-ID-20570865f5"],
  "init": false
}
```
Response body:
```json
{
  "headers": {
    "channel": "default-channel",
    "timeReceived": "",
    "timeElapsed": 0,
    "requestOffset": "",
    "requestId": ""
  },
  "result": [
    {
      "isDelete": false,
      "record": {
        "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
        "accident_incident_report_date": "2023-08-17",
        "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
        "accident_incident_report_event_overview": [
          "Event:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Event:2 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat"
        ],
        "accident_incident_report_follow_up_actions": [
          "Action:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Action:2 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:3 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:4 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:5 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
        ],
        "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
        "accident_incident_report_immediate_action": [
          "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
        ],
        "accident_incident_report_line_of_communication": [
          "Line of communication:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Line of communication:2 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat"
        ],
        "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
        "accident_incident_report_number_of_people_involved": 1,
        "accident_incident_report_person_injured": [
          "John Doe"
        ],
        "accident_incident_report_status": "closed",
        "accident_incident_report_time_end": "2024-02-17 17:00:00",
        "accident_incident_report_time_start": "2024-02-17 21:41:43",
        "accident_incident_report_validated": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
        "accident_incident_report_witnesses": [
          "Beth A. Foley",
          "Nena H. Woodcock"
        ],
        "docType": "accident-incident-report",
        "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
      },
      "timestamp_accident_incident_report": "2023-08-17T22:23:41.669982648Z",
      "txId": "c255de0b6cee3e07243f191dba8d6b9d349a34a723a110283d8f9081c5e2a85b"
    },
    {
      "isDelete": false,
      "record": {
        "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
        "accident_incident_report_date": "2023-08-17",
        "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
        "accident_incident_report_event_overview": [
          "Event:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Event:2 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat"
        ],
        "accident_incident_report_follow_up_actions": [
          "Action:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Action:2 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:3 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:4 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:5 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
        ],
        "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
        "accident_incident_report_immediate_action": [
          "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
        ],
        "accident_incident_report_line_of_communication": [
          "not-set"
        ],
        "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
        "accident_incident_report_number_of_people_involved": 1,
        "accident_incident_report_person_injured": [
          "John Doe"
        ],
        "accident_incident_report_status": "closed",
        "accident_incident_report_time_end": "2024-02-17 17:00:00",
        "accident_incident_report_time_start": "2024-02-17 21:41:43",
        "accident_incident_report_validated": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
        "accident_incident_report_witnesses": [
          "Beth A. Foley",
          "Nena H. Woodcock"
        ],
        "docType": "accident-incident-report",
        "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
      },
      "timestamp_accident_incident_report": "2023-08-17T22:19:13.596967836Z",
      "txId": "cab95c84a0aff2da02514fe4e7ca46a272471e2bc1282c4430675c97bef12a64"
    },
    {
      "isDelete": false,
      "record": {
        "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
        "accident_incident_report_date": "2023-08-17",
        "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
        "accident_incident_report_event_overview": [
          "not-set"
        ],
        "accident_incident_report_follow_up_actions": [
          "Action:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Action:2 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:3 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:4 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:5 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
        ],
        "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
        "accident_incident_report_immediate_action": [
          "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
        ],
        "accident_incident_report_line_of_communication": [
          "not-set"
        ],
        "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
        "accident_incident_report_number_of_people_involved": 1,
        "accident_incident_report_person_injured": [
          "John Doe"
        ],
        "accident_incident_report_status": "closed",
        "accident_incident_report_time_end": "2024-02-17 17:00:00",
        "accident_incident_report_time_start": "2024-02-17 21:41:43",
        "accident_incident_report_validated": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
        "accident_incident_report_witnesses": [
          "Beth A. Foley",
          "Nena H. Woodcock"
        ],
        "docType": "accident-incident-report",
        "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
      },
      "timestamp_accident_incident_report": "2023-08-17T22:16:20.225953435Z",
      "txId": "a1e0a1282c83b2d127fc5bd9b3377fd2e54fb3b969447e46c368981f8c0b4983"
    },
    {
      "isDelete": false,
      "record": {
        "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
        "accident_incident_report_date": "2023-08-17",
        "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
        "accident_incident_report_event_overview": [
          "not-set"
        ],
        "accident_incident_report_follow_up_actions": [
          "Action:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Action:2 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:3 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:4 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:5 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
        ],
        "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
        "accident_incident_report_immediate_action": [
          "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
        ],
        "accident_incident_report_line_of_communication": [
          "not-set"
        ],
        "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
        "accident_incident_report_number_of_people_involved": 1,
        "accident_incident_report_person_injured": [
          "John Doe"
        ],
        "accident_incident_report_status": "closed",
        "accident_incident_report_time_end": "2024-02-17 17:00:00",
        "accident_incident_report_time_start": "2024-02-17 21:41:43",
        "accident_incident_report_validated": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
        "accident_incident_report_witnesses": [
          "not-set"
        ],
        "docType": "accident-incident-report",
        "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
      },
      "timestamp_accident_incident_report": "2023-08-17T22:11:48.159061546Z",
      "txId": "aaf5e3058f7f7a14505a724dc5debce1f18c2602c4262fb5f07c1328862300da"
    },
    {
      "isDelete": false,
      "record": {
        "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
        "accident_incident_report_date": "2023-08-17",
        "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
        "accident_incident_report_event_overview": [
          "not-set"
        ],
        "accident_incident_report_follow_up_actions": [
          "Action:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Action:2 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:3 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:4 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:5 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
        ],
        "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
        "accident_incident_report_immediate_action": [
          "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
        ],
        "accident_incident_report_line_of_communication": [
          "not-set"
        ],
        "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
        "accident_incident_report_number_of_people_involved": 1,
        "accident_incident_report_person_injured": [
          "not-set"
        ],
        "accident_incident_report_status": "closed",
        "accident_incident_report_time_end": "2024-02-17 17:00:00",
        "accident_incident_report_time_start": "2024-02-17 21:41:43",
        "accident_incident_report_validated": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
        "accident_incident_report_witnesses": [
          "not-set"
        ],
        "docType": "accident-incident-report",
        "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
      },
      "timestamp_accident_incident_report": "2023-08-17T22:06:54.705798602Z",
      "txId": "58162f3eac752790c0db5a44a75bba2652fb4d02983b43a92303a3844effd71e"
    },
    {
      "isDelete": false,
      "record": {
        "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
        "accident_incident_report_date": "2023-08-17",
        "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
        "accident_incident_report_event_overview": [
          "not-set"
        ],
        "accident_incident_report_follow_up_actions": [
          "Action:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Action:2 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:3 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:4 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:5 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
        ],
        "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
        "accident_incident_report_immediate_action": [
          "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
        ],
        "accident_incident_report_line_of_communication": [
          "not-set"
        ],
        "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
        "accident_incident_report_number_of_people_involved": 0,
        "accident_incident_report_person_injured": [
          "not-set"
        ],
        "accident_incident_report_status": "closed",
        "accident_incident_report_time_end": "2024-02-17 17:00:00",
        "accident_incident_report_time_start": "2024-02-17 21:41:43",
        "accident_incident_report_validated": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
        "accident_incident_report_witnesses": [
          "not-set"
        ],
        "docType": "accident-incident-report",
        "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
      },
      "timestamp_accident_incident_report": "2023-08-17T22:01:10.907024746Z",
      "txId": "401de99cb15b56cc41ab22b2311cd4243a8290296ef6d32cde2b765454e4bcb2"
    },
    {
      "isDelete": false,
      "record": {
        "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
        "accident_incident_report_date": "2023-08-17",
        "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
        "accident_incident_report_event_overview": [
          "not-set"
        ],
        "accident_incident_report_follow_up_actions": [
          "Action:1 - Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Action:2 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:3 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:4 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat.",
          "Action:5 -Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
        ],
        "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
        "accident_incident_report_immediate_action": [
          "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
        ],
        "accident_incident_report_line_of_communication": [
          "not-set"
        ],
        "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
        "accident_incident_report_number_of_people_involved": 0,
        "accident_incident_report_person_injured": [
          "not-set"
        ],
        "accident_incident_report_status": "open",
        "accident_incident_report_time_end": "2024-02-17 17:00:00",
        "accident_incident_report_time_start": "2024-02-17 21:41:43",
        "accident_incident_report_validated": "not-set",
        "accident_incident_report_witnesses": [
          "not-set"
        ],
        "docType": "accident-incident-report",
        "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
      },
      "timestamp_accident_incident_report": "2023-08-17T21:54:53.65852685Z",
      "txId": "d02c2082011703952e2c69c00a8b20ddee8e10bf99834594b54ffd3a89cb3d68"
    },
    {
      "isDelete": false,
      "record": {
        "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
        "accident_incident_report_date": "2023-08-17",
        "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
        "accident_incident_report_event_overview": [
          "not-set"
        ],
        "accident_incident_report_follow_up_actions": [
          "not-set"
        ],
        "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
        "accident_incident_report_immediate_action": [
          "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
        ],
        "accident_incident_report_line_of_communication": [
          "not-set"
        ],
        "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
        "accident_incident_report_number_of_people_involved": 0,
        "accident_incident_report_person_injured": [
          "not-set"
        ],
        "accident_incident_report_status": "open",
        "accident_incident_report_time_end": "2024-02-17 17:00:00",
        "accident_incident_report_time_start": "2024-02-17 21:41:43",
        "accident_incident_report_validated": "not-set",
        "accident_incident_report_witnesses": [
          "not-set"
        ],
        "docType": "accident-incident-report",
        "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
      },
      "timestamp_accident_incident_report": "2023-08-17T21:49:04.882853279Z",
      "txId": "e86cdcfdc25ac1c61932450dce917127ed6aee2a619bcf4e79cf3529290552c2"
    },
    {
      "isDelete": false,
      "record": {
        "accident_incident_report_classification": "Calssification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_created_by": "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==",
        "accident_incident_report_date": "2023-08-17",
        "accident_incident_report_description": "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
        "accident_incident_report_event_overview": [
          "not-set"
        ],
        "accident_incident_report_follow_up_actions": [
          "not-set"
        ],
        "accident_incident_report_hs_aspects": "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
        "accident_incident_report_id": "ACCIDENT-INCIDENT-REPORT-ID-20570865f5",
        "accident_incident_report_immediate_action": [
          "Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat",
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."
        ],
        "accident_incident_report_line_of_communication": [
          "not-set"
        ],
        "accident_incident_report_location": "Location: Lorem ipsum dolor sit amet.",
        "accident_incident_report_number_of_people_involved": 0,
        "accident_incident_report_person_injured": [
          "not-set"
        ],
        "accident_incident_report_status": "open",
        "accident_incident_report_time_end": "not-set",
        "accident_incident_report_time_start": "2024-02-17 21:41:43",
        "accident_incident_report_validated": "not-set",
        "accident_incident_report_witnesses": [
          "not-set"
        ],
        "docType": "accident-incident-report",
        "subject_accident_incident_report": "Subject: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat."
      },
      "timestamp_accident_incident_report": "2023-08-17T21:41:43.034990733Z",
      "txId": "567603d03a5db22dbfdb2e47fd8b5a43ffc8a55ba62d2b146f512eacc577ae97"
    }
  ]
}
```






