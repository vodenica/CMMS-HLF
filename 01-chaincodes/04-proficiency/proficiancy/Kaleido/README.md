# Kaleido platform

Starting with the installation of the chaincode on the platform Kaleido
Below is a sample build command to be used when we are creating a new version of the chaincode/smart contract:
```
GOOS=linux GOARCH=amd64 go build -o cmsCompetency.bin
```
## Enrolling the user on "Kaleido" platform
Before we start working with the ledger, we must enrol a user, and in Swagger, we run the following nameplate. 
```
POST /identities (Enroll user)
```
### Response for "userTEST3"
The response from the action is shown below:
```json
{
  "name": "userTEST3",
  "secret": "MJbHUTWDwYch"
}
```
We use "secret" to enrol the user
## Transaction for _CreateNewModuleOne_ 
Swagger - to be added with the following records: 
```
POST /transaction
```
Add these JSON file to create new Module 1 
```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "userTEST3",
    "channel": "default-channel",
    "chaincode": "competency"
  },
  "func": "CreateNewModuleOne",
  "args": ["CMS-Module-1-0001", "{\"trainer_module_one_id\": \"12345678\", \"trainer_module_one_name\": \"Dalibor\", \"trainer_module_one_surname\": \"Vodenicarski\"}", "{\"trainee_module_one_id\": \"MID-12345\", \"trainee_module_one_name\": \"Alice\", \"trainee_module_one_surname\": \"Irving\"}", "{\"module_one_chapter_one_session_one\": \"Session 1 - Introduction\", \"module_one_chapter_one_session_two\": \"Session 2 - Basic of Operational Services\", \"module_one_chapter_one_session_three\": \"Session 3 - Introduction to the Health and Safety\", \"module_one_chapter_one_session_four\": \"Session 4 - System & Site Orientation\"}", "{\"module_one_chapter_two_session_one\": \"Session 1 - General Concept of a Ropeway\", \"module_one_chapter_two_session_two\": \"Session 2 - Specific Components in the System\", \"module_one_chapter_two_session_three\": \"Session 3 - Normal & Degraded Operation Procedures\"}", "{\"module_one_chapter_three_session_one\": \"Session 1 - Troubleshooting Procedures\", \"module_one_chapter_three_session_two\": \"Session 2 - Stop Circuit - Basic Troubleshooting\", \"module_one_chapter_three_session_three\": \"Session 3 - Emergency Stop Circuit - Basic Troublesh\"}"],
  "init": false
}
```
### Transaction body:
```json
{
  "headers": {
    "id": "d11978f9-0fde-4dea-7378-fb4467784ed2",
    "type": "TransactionSuccess",
    "timeReceived": "2023-02-18T11:50:03.714021002Z",
    "timeElapsed": 0.296805054,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 124,
  "signerMSP": "u0mk35n6yh",
  "signer": "userTEST3",
  "transactionID": "8aed70436c2d68198d34dffe9805eeb2038bd88317eb50a9684c97354157fe64",
  "status": "VALID"
}
```

#### The end of the file ...
