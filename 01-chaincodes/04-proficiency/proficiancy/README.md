[Main page](../../../README.md)

# Proficiency Services Chaincode

Proficiency services [chaincode](../04-proficiency/proficiency) is a part of the CMMS solution, which is responsible for keeping and updating records on the ledger of the training, custom training and assessment of the employee. I've used a complex smart contract repository from [fabric-contract-api-go](https://github.com/vodenica/fabric-contract-api-go.git) to create a proficiency chaincode.
The smart contracts or chaincodes are also considered assessment parts, where smart contracts are used to update the state of the ledger with the results of the assessments, both theoretical and practical. Also, the total value is added to determine if the employee passes the evaluations from the dedicated model.
Chaincode is tested on the [Kaleido](https://www.kaleido.io/) platform, and details of how these tests are performed can be found [here](Kaleido/README.md).

## Testing platforms
* Hyperledger Fabric `test-network` with `Caliper` benchmarking testing framework ([`Caliper Testing`](#caliper-testing) & [`Test Network`](#test-network)).
* Hyperledger Fabric on the `Kaleido` Web3 BaaS platform, [`Kaleido Testing`](#kaleido-testing).
* REST API testing with the [`k6`](../../../02-benchmarking-files/k6/k6-linux/README.md) load testing framework.

## Proficiency module of the proposed CMMS

This part of the project consists of three main modules, such as:

- Module 1 - Training for System Operators
- Module 2 - Training for Maintenance Technicians
- Module 3 - Training for the Management of the Enterprise

## Spin up the `test-network`

Follow the instructions from the `caliper` testing repository.

Cleanup the terminal:

```bash
PS1='\e[1;32m\u@\H:\e[0m\e[1;34m\W\e[0m\e[1;33m $:\e[0m '
```

Before we start installing the test-network, we have to ensure that Docker is up and running, and then we navigate to the test-network folder.

```bash
cd ~/go/src/github.com/hyperledger/fabric-samples/test-network/
```

Spin up the `test-network`:

```bash
./network.sh up createChannel -ca -s couchdb
```

Install the `Proficiency Operations` Smart Contract:

```bash
./network.sh deployCC -ccn proficiency -ccv 1.0 -ccp /home/vodenica/Desktop/test-network-smart-contracts/05-CMS-competency/DEV-CMS-test -ccl go
```

The new version of the smart contract is deployed to the `test-network` and the `test-network` is ready for testing. It is important to update the version in the `main.go` file before deploying the smart contract.

```bash
./network.sh deployCC -ccn proficiency -ccv 2.0 -ccp /home/vodenica/Desktop/test-network-smart-contracts/03-MTE-maintenance -ccl go
```

To test the installed `Proficiency Operations Smart Contract`, navigate to `fabric-samples/test-network` and define all necessary environment variables for `Org1MSP`. Note that TLS is enabled in `test-network`.
Run the following command:

```bash
cd $HOME/go/src/github.com/hyperledger/fabric-samples/test-network
```

```bash
export FABRIC_CFG_PATH=$PWD/../config/
```

```bash
export PATH=${PWD}/../bin:$PATH
```

#### Environmental set-up for `Org1MSP`:

```bash
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
```

During the development environment setup phase, we downloaded Hyperledger Fabric
binaries including peer. They are located in the `fabric-samples/bin` folder and utilize
configurations stored in `fabric-samples/config. Therefore, we can update the `PATH variable
and set `FABRIC_CFG_PATH` to simplify `peer binary usage.

```bash
export PATH=${PWD}/../bin:$PATH
```

```bash
export FABRIC_CFG_PATH=$PWD/../config/
```

### Environmental variables for `Org2MSP`.

```bash
cd $HOME/go/src/github.com/hyperledger/fabric-samples/test-network
```

```bash
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051
```

```bash
export PATH=${PWD}/../bin:$PATH
```

```bash
export FABRIC_CFG_PATH=$PWD/../config/
```

Now, in the Org1MSP terminal window we can run a peer channel list command to confirm that we are able to use the peer binary without further adjustments.

```bash
peer channel list
```

Response:

```bash
Channels peers has joined:
mychannel
```

The commit transaction is submitted to peers of both `Org1MSP` and `Org2MSP`. The Smart Contract definition is committed to the channel if all targeted peers return successful responses. To confirm this, use the peer lifecycle chaincode `querycommitted` command.

```bash
peer lifecycle chaincode querycommitted --channelID mychannel --name proficiency --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

Response:

```bash
Committed chaincode definition for chaincode 'proficiency' on channel 'mychannel':
Version: 1.0, Sequence: 1, Endorsement Plugin: escc, Validation Plugin: vscc, Approvals: [Org1MSP: true, Org2MSP: true]
```

# Test Network

## Smart contract function `CreateNewModuleOne`

The function `CreateNewModuelOne` is used to create a proficiency module one for training purposes. The function takes 6 arguments. The first argument is the ID of the module. The remaining 5 arguments are the values of the module.

```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n proficiency --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"CreateNewModuleOne","Args":["proficiency-module-one-ID-0001", "{\"trainer_module_one_id\": \"Trainer-ID\", \"trainer_module_one_name\": \"Trainer Name\", \"trainer_module_one_surname\": \"Trainer Surname\"}", "{\"trainee_module_one_id\": \"Trainee-ID\", \"trainee_module_one_name\": \"Trainee Name\", \"trainee_module_one_surname\": \"Trainee Surname\"}", "{\"module_one_chapter_one_session_one\":\"Module One - Chapter 01 - Seesion 03\",\"module_one_chapter_one_session_two\":\"Module One - Chapter 01 - Seesion 03\", \"module_one_chapter_one_session_three\":\"Module One - Chapter 01 - Seesion 03\"}", "{\"module_one_chapter_two_session_one\":\"Module One - Chapter 02 - Seesion 03\",\"module_one_chapter_two_session_two\":\"Module One - Chapter 02 - Seesion 03\", \"module_one_chapter_two_session_three\":\"Module One - Chapter 02 - Seesion 03\"}", "{\"module_one_chapter_three_session_one\":\"Module One - Chapter 03 - Seesion 03\",\"module_one_chapter_three_session_two\":\"Module One - Chapter 03 - Seesion 03\", \"module_one_chapter_three_session_three\":\"Module One - Chapter 03 - Seesion 03\"}"]}'
```

The `JSON` format data of the created proficiency module one on the ledger is as follows:

```json
{
  "_id": "proficiency-module-one-ID-0001",
  "_rev": "1-1bf1cd59dbd811e6f817e79fbc01dd0a",
  "assessment_module_one": 0,
  "assessments_attempts": 0,
  "module_one_chapter_one": {
    "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
    "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03",
    "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03"
  },
  "module_one_chapter_three": {
    "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
    "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03",
    "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03"
  },
  "module_one_chapter_two": {
    "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
    "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03",
    "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03"
  },
  "module_one_id": "proficiency-module-one-ID-0001",
  "module_one_training_created_date": "2023-09-18 01:52:16",
  "practical_assessment_module_one": 0,
  "set_status_chapter_one_module_one": 0,
  "set_status_chapter_three_module_one": 0,
  "set_status_chapter_two_module_one": 0,
  "theoretical_assessment_module_one": 0,
  "trainee_module_one": {
    "trainee_module_one_id": "Trainee-ID",
    "trainee_module_one_name": "Trainee Name",
    "trainee_module_one_surname": "Trainee Surname"
  },
  "trainer_module_one": {
    "trainer_module_one_id": "Trainer-ID",
    "trainer_module_one_name": "Trainer Name",
    "trainer_module_one_surname": "Trainer Surname"
  },
  "training_status_module_one": 1,
  "training_type_module_one": 0,
  "~version": "CgMBBgA="
}
```

## Smart contract function `GetHistoryRecordsForModuleOne`

```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n proficiency --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"GetHistoryRecordsForModuleOne","Args":["proficiency-module-one-ID-0001"]}'
```

Response body:

```json
[
  {
    "record": {
      "module_one_id": "proficiency-module-one-ID-0001",
      "trainer_module_one": {
        "trainer_module_one_id": "Trainer-ID",
        "trainer_module_one_name": "Trainer Name",
        "trainer_module_one_surname": "Trainer Surname"
      },
      "trainee_module_one": {
        "trainee_module_one_id": "Trainee-ID",
        "trainee_module_one_name": "Trainee Name",
        "trainee_module_one_surname": "Trainee Surname"
      },
      "module_one_training_created_date": "2023-09-15",
      "module_one_chapter_one": {
        "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
        "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03",
        "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03"
      },
      "set_status_chapter_one_module_one": 0,
      "module_one_chapter_two": {
        "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
        "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03",
        "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03"
      },
      "set_status_chapter_two_module_one": 0,
      "module_one_chapter_three": {
        "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
        "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03",
        "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03"
      },
      "set_status_chapter_three_module_one": 0,
      "training_type_module_one": 0,
      "theoretical_assessment_module_one": 0,
      "practical_assessment_module_one": 0,
      "assessment_module_one": 0,
      "training_status_module_one": 1,
      "assessments_attempts": 0
    },
    "txId": "3e39bd13e452db17d6356c458da9a13a823b410d348cdbdbb146dc5d29bde0e0",
    "timestamp": "2023-09-15T23:28:20.93105461Z",
    "isDelete": false
  }
]
```

## Smart contract function `UpdateStatusChapterOneModuleOneToCompleted`

Request body:

```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n proficiency --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"UpdateStatusChapterOneModuleOneToCompleted","Args":["proficiency-module-one-ID-0001"]}'
```

The `JSON` data format response:

```json
[
  {
    "record": {
      "module_one_id": "proficiency-module-one-ID-0001",
      "trainer_module_one": {
        "trainer_module_one_id": "Trainer-ID",
        "trainer_module_one_name": "Trainer Name",
        "trainer_module_one_surname": "Trainer Surname"
      },
      "trainee_module_one": {
        "trainee_module_one_id": "Trainee-ID",
        "trainee_module_one_name": "Trainee Name",
        "trainee_module_one_surname": "Trainee Surname"
      },
      "module_one_training_created_date": "2023-09-15",
      "module_one_chapter_one": {
        "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
        "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03",
        "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03"
      },
      "set_status_chapter_one_module_one": 1,
      "module_one_chapter_two": {
        "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
        "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03",
        "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03"
      },
      "set_status_chapter_two_module_one": 0,
      "module_one_chapter_three": {
        "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
        "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03",
        "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03"
      },
      "set_status_chapter_three_module_one": 0,
      "training_type_module_one": 0,
      "theoretical_assessment_module_one": 0,
      "practical_assessment_module_one": 0,
      "assessment_module_one": 0,
      "training_status_module_one": 1,
      "assessments_attempts": 0
    },
    "txId": "f89e2f0b83a01b9a7ab7f9a9df04a98a1f8c18fa9e9c89f7411e70e4d7f9e391",
    "timestamp": "2023-09-15T23:43:55.591826251Z",
    "isDelete": false
  },
  {
    "record": {
      "module_one_id": "proficiency-module-one-ID-0001",
      "trainer_module_one": {
        "trainer_module_one_id": "Trainer-ID",
        "trainer_module_one_name": "Trainer Name",
        "trainer_module_one_surname": "Trainer Surname"
      },
      "trainee_module_one": {
        "trainee_module_one_id": "Trainee-ID",
        "trainee_module_one_name": "Trainee Name",
        "trainee_module_one_surname": "Trainee Surname"
      },
      "module_one_training_created_date": "2023-09-15",
      "module_one_chapter_one": {
        "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
        "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03",
        "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03"
      },
      "set_status_chapter_one_module_one": 0,
      "module_one_chapter_two": {
        "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
        "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03",
        "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03"
      },
      "set_status_chapter_two_module_one": 0,
      "module_one_chapter_three": {
        "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
        "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03",
        "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03"
      },
      "set_status_chapter_three_module_one": 0,
      "training_type_module_one": 0,
      "theoretical_assessment_module_one": 0,
      "practical_assessment_module_one": 0,
      "assessment_module_one": 0,
      "training_status_module_one": 1,
      "assessments_attempts": 0
    },
    "txId": "3e39bd13e452db17d6356c458da9a13a823b410d348cdbdbb146dc5d29bde0e0",
    "timestamp": "2023-09-15T23:28:20.93105461Z",
    "isDelete": false
  }
]
```

## Smart contract function `UpdateValueTheoreticalAssessment`

Assuming that the candidate has passed the theoretical assessment, we can update the value of the theoretical assessment to 75%. The command can be passed as follows below.

Request body:

```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n proficiency --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"UpdateValueTheoreticalAssessment","Args":["proficiency-module-one-ID-0001", "75"]}'
```

## Smart contract function `UpdateValuePracticalAndTotalAssessment`

Assuming that the candidate has passed the practical assessment, we can update the value of the practical assessment to 100%. The command can be passed as follows below.

Request body:

```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n proficiency --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"UpdateValuePracticalAndTotalAssessment","Args":["proficiency-module-one-ID-0001", "100"]}'
```

Response `JSON` format data on the ledger:

```json
{
  "_id": "proficiency-module-one-ID-0001",
  "_rev": "5-ffe4ec27099cb0460d1083029641a313",
  "assessment_module_one": 175,
  "assessments_attempts": 1,
  "module_one_chapter_one": {
    "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
    "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03",
    "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03"
  },
  "module_one_chapter_three": {
    "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
    "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03",
    "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03"
  },
  "module_one_chapter_two": {
    "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
    "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03",
    "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03"
  },
  "module_one_id": "proficiency-module-one-ID-0001",
  "module_one_training_created_date": "2023-09-18 01:52:16",
  "practical_assessment_module_one": 100,
  "set_status_chapter_one_module_one": 1,
  "set_status_chapter_three_module_one": 0,
  "set_status_chapter_two_module_one": 0,
  "theoretical_assessment_module_one": 75,
  "trainee_module_one": {
    "trainee_module_one_id": "Trainee-ID",
    "trainee_module_one_name": "Trainee Name",
    "trainee_module_one_surname": "Trainee Surname"
  },
  "trainer_module_one": {
    "trainer_module_one_id": "Trainer-ID",
    "trainer_module_one_name": "Trainer Name",
    "trainer_module_one_surname": "Trainer Surname"
  },
  "training_status_module_one": 3,
  "training_type_module_one": 0,
  "~version": "CgMBCwA="
}
```

## Smart contract function `GetHistoryRecordsForModuleOne`

```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n proficiency --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"GetHistoryRecordsForModuleOne","Args":["proficiency-module-one-ID-0001"]}'
```

Response body:

```json
[
  {
    "record": {
      "module_one_id": "proficiency-module-one-ID-0001",
      "trainer_module_one": {
        "trainer_module_one_id": "Trainer-ID",
        "trainer_module_one_name": "Trainer Name",
        "trainer_module_one_surname": "Trainer Surname"
      },
      "trainee_module_one": {
        "trainee_module_one_id": "Trainee-ID",
        "trainee_module_one_name": "Trainee Name",
        "trainee_module_one_surname": "Trainee Surname"
      },
      "module_one_training_created_date": "2023-09-18 01:52:16",
      "module_one_chapter_one": {
        "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
        "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03",
        "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03"
      },
      "set_status_chapter_one_module_one": 1,
      "module_one_chapter_two": {
        "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
        "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03",
        "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03"
      },
      "set_status_chapter_two_module_one": 0,
      "module_one_chapter_three": {
        "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
        "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03",
        "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03"
      },
      "set_status_chapter_three_module_one": 0,
      "training_type_module_one": 0,
      "theoretical_assessment_module_one": 75,
      "practical_assessment_module_one": 200,
      "assessment_module_one": 275,
      "training_status_module_one": 3,
      "assessments_attempts": 2
    },
    "txId": "acbc51de237ab874c281c6139bd6f25913c04f689ea3c72275f3ca4c5520c34a",
    "timestamp": "2023-09-18T02:06:10.848797157Z",
    "isDelete": false
  },
  {
    "record": {
      "module_one_id": "proficiency-module-one-ID-0001",
      "trainer_module_one": {
        "trainer_module_one_id": "Trainer-ID",
        "trainer_module_one_name": "Trainer Name",
        "trainer_module_one_surname": "Trainer Surname"
      },
      "trainee_module_one": {
        "trainee_module_one_id": "Trainee-ID",
        "trainee_module_one_name": "Trainee Name",
        "trainee_module_one_surname": "Trainee Surname"
      },
      "module_one_training_created_date": "2023-09-18 01:52:16",
      "module_one_chapter_one": {
        "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
        "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03",
        "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03"
      },
      "set_status_chapter_one_module_one": 1,
      "module_one_chapter_two": {
        "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
        "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03",
        "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03"
      },
      "set_status_chapter_two_module_one": 0,
      "module_one_chapter_three": {
        "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
        "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03",
        "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03"
      },
      "set_status_chapter_three_module_one": 0,
      "training_type_module_one": 0,
      "theoretical_assessment_module_one": 75,
      "practical_assessment_module_one": 100,
      "assessment_module_one": 175,
      "training_status_module_one": 3,
      "assessments_attempts": 1
    },
    "txId": "6a7e38ea8c3ed9636b20a8faa356da5234d918f0838aa2cdee18fecab7ababd8",
    "timestamp": "2023-09-18T02:03:57.678710055Z",
    "isDelete": false
  },
  {
    "record": {
      "module_one_id": "proficiency-module-one-ID-0001",
      "trainer_module_one": {
        "trainer_module_one_id": "Trainer-ID",
        "trainer_module_one_name": "Trainer Name",
        "trainer_module_one_surname": "Trainer Surname"
      },
      "trainee_module_one": {
        "trainee_module_one_id": "Trainee-ID",
        "trainee_module_one_name": "Trainee Name",
        "trainee_module_one_surname": "Trainee Surname"
      },
      "module_one_training_created_date": "2023-09-18 01:52:16",
      "module_one_chapter_one": {
        "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
        "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03",
        "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03"
      },
      "set_status_chapter_one_module_one": 1,
      "module_one_chapter_two": {
        "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
        "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03",
        "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03"
      },
      "set_status_chapter_two_module_one": 0,
      "module_one_chapter_three": {
        "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
        "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03",
        "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03"
      },
      "set_status_chapter_three_module_one": 0,
      "training_type_module_one": 0,
      "theoretical_assessment_module_one": 75,
      "practical_assessment_module_one": 0,
      "assessment_module_one": 0,
      "training_status_module_one": 1,
      "assessments_attempts": 0
    },
    "txId": "f4fd77054c3d21665a3905046d9ca6411163d7ce10f98526bdcaa2a5ddf96e23",
    "timestamp": "2023-09-18T01:57:01.135520726Z",
    "isDelete": false
  },
  {
    "record": {
      "module_one_id": "proficiency-module-one-ID-0001",
      "trainer_module_one": {
        "trainer_module_one_id": "Trainer-ID",
        "trainer_module_one_name": "Trainer Name",
        "trainer_module_one_surname": "Trainer Surname"
      },
      "trainee_module_one": {
        "trainee_module_one_id": "Trainee-ID",
        "trainee_module_one_name": "Trainee Name",
        "trainee_module_one_surname": "Trainee Surname"
      },
      "module_one_training_created_date": "2023-09-18 01:52:16",
      "module_one_chapter_one": {
        "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
        "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03",
        "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03"
      },
      "set_status_chapter_one_module_one": 1,
      "module_one_chapter_two": {
        "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
        "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03",
        "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03"
      },
      "set_status_chapter_two_module_one": 0,
      "module_one_chapter_three": {
        "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
        "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03",
        "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03"
      },
      "set_status_chapter_three_module_one": 0,
      "training_type_module_one": 0,
      "theoretical_assessment_module_one": 0,
      "practical_assessment_module_one": 0,
      "assessment_module_one": 0,
      "training_status_module_one": 1,
      "assessments_attempts": 0
    },
    "txId": "a93235556e1e0aeb7f898abe23c37c9cccb684aaa4fdf4d9a6560457e83e7b39",
    "timestamp": "2023-09-18T01:54:11.477099575Z",
    "isDelete": false
  },
  {
    "record": {
      "module_one_id": "proficiency-module-one-ID-0001",
      "trainer_module_one": {
        "trainer_module_one_id": "Trainer-ID",
        "trainer_module_one_name": "Trainer Name",
        "trainer_module_one_surname": "Trainer Surname"
      },
      "trainee_module_one": {
        "trainee_module_one_id": "Trainee-ID",
        "trainee_module_one_name": "Trainee Name",
        "trainee_module_one_surname": "Trainee Surname"
      },
      "module_one_training_created_date": "2023-09-18 01:52:16",
      "module_one_chapter_one": {
        "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
        "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03",
        "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03"
      },
      "set_status_chapter_one_module_one": 0,
      "module_one_chapter_two": {
        "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
        "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03",
        "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03"
      },
      "set_status_chapter_two_module_one": 0,
      "module_one_chapter_three": {
        "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
        "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03",
        "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03"
      },
      "set_status_chapter_three_module_one": 0,
      "training_type_module_one": 0,
      "theoretical_assessment_module_one": 0,
      "practical_assessment_module_one": 0,
      "assessment_module_one": 0,
      "training_status_module_one": 1,
      "assessments_attempts": 0
    },
    "txId": "060ba341e14e51b519451ddcd6bd1b61e677d6662f170d4fb27447640ff11b4a",
    "timestamp": "2023-09-18T01:52:16.090698727Z",
    "isDelete": false
  }
]
```

[Back to top](#testing-platforms)

# Kaleido testing

## Smart contract function `CreateNewModuleOne`

Request body:

```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "Operations Supervisor",
    "channel": "default-channel",
    "chaincode": "competency"
  },
  "func": "CreateNewModuleOne",
  "args": [
    "proficiency-module-one-ID-0001X",
    "{\"trainer_module_one_id\": \"Trainer-ID\", \"trainer_module_one_name\": \"Trainer Name\", \"trainer_module_one_surname\": \"Trainer Surname\"}",
    "{\"trainee_module_one_id\": \"Trainee-ID\", \"trainee_module_one_name\": \"Trainee Name\", \"trainee_module_one_surname\": \"Trainee Surname\"}",
    "{\"module_one_chapter_one_session_one\":\"Module One - Chapter 01 - Seesion 03\",\"module_one_chapter_one_session_two\":\"Module One - Chapter 01 - Seesion 03\", \"module_one_chapter_one_session_three\":\"Module One - Chapter 01 - Seesion 03\"}",
    "{\"module_one_chapter_two_session_one\":\"Module One - Chapter 02 - Seesion 03\",\"module_one_chapter_two_session_two\":\"Module One - Chapter 02 - Seesion 03\", \"module_one_chapter_two_session_three\":\"Module One - Chapter 02 - Seesion 03\"}",
    "{\"module_one_chapter_three_session_one\":\"Module One - Chapter 03 - Seesion 03\",\"module_one_chapter_three_session_two\":\"Module One - Chapter 03 - Seesion 03\", \"module_one_chapter_three_session_three\":\"Module One - Chapter 03 - Seesion 03\"}"
  ],
  "init": false
}
```

Response body:

```json
{
  "headers": {
    "id": "d4967daa-ee4a-470d-6978-7f74ad3a67b0",
    "type": "TransactionSuccess",
    "timeReceived": "2023-09-21T01:26:17.813386084Z",
    "timeElapsed": 0.316428618,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 1123,
  "signerMSP": "u0mk35n6yh",
  "signer": "Operations Supervisor",
  "transactionID": "014e4c5b464b3c44a07f00f5b1b456ee78692ada4681e29b83b738019d02616e",
  "status": "VALID"
}
```

Tx id:

```
014e4c5b464b3c44a07f00f5b1b456ee78692ada4681e29b83b738019d02616e
```

Response body in `JSON` format of the created asset:

```json
{
  "module_one_id": "proficiency-module-one-ID-0001X",
  "trainer_module_one": {
    "trainer_module_one_id": "Trainer-ID",
    "trainer_module_one_name": "Trainer Name",
    "trainer_module_one_surname": "Trainer Surname"
  },
  "trainee_module_one": {
    "trainee_module_one_id": "Trainee-ID",
    "trainee_module_one_name": "Trainee Name",
    "trainee_module_one_surname": "Trainee Surname"
  },
  "module_one_training_created_date": "2023-09-21 01:26:17",
  "module_one_chapter_one": {
    "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
    "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03",
    "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03"
  },
  "set_status_chapter_one_module_one": 0,
  "module_one_chapter_two": {
    "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
    "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03",
    "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03"
  },
  "set_status_chapter_two_module_one": 0,
  "module_one_chapter_three": {
    "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
    "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03",
    "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03"
  },
  "set_status_chapter_three_module_one": 0,
  "training_type_module_one": 0,
  "theoretical_assessment_module_one": 0,
  "practical_assessment_module_one": 0,
  "assessment_module_one": 0,
  "training_status_module_one": 1,
  "assessments_attempts": 0
}
```

## Smart contract function `UpdateValueTheoreticalAssessment`

Assuming that the candidate has passed the theoretical assessment, we can update the value of the theoretical assessment to 75%. The command can be passed as follows below.

Request body `POST /transactions`:

```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "Operations Supervisor",
    "channel": "default-channel",
    "chaincode": "competency"
  },
  "func": "UpdateValueTheoreticalAssessment",
  "args": ["proficiency-module-one-ID-0001X", "75"],
  "init": false
}
```

Response body:

```json
{
  "headers": {
    "id": "65ac01b2-deba-40fe-7b8b-1047a87d6c83",
    "type": "TransactionSuccess",
    "timeReceived": "2023-09-21T02:15:28.716852893Z",
    "timeElapsed": 0.312208839,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 1124,
  "signerMSP": "u0mk35n6yh",
  "signer": "Operations Supervisor",
  "transactionID": "24670a5bac7be204fde289d063de5350415674e39fe271e3a975d44d1bd40f22",
  "status": "VALID"
}
```

Reposne body in `JSON` format recorded on the ledger:

```json
{
  "module_one_id": "proficiency-module-one-ID-0001X",
  "trainer_module_one": {
    "trainer_module_one_id": "Trainer-ID",
    "trainer_module_one_name": "Trainer Name",
    "trainer_module_one_surname": "Trainer Surname"
  },
  "trainee_module_one": {
    "trainee_module_one_id": "Trainee-ID",
    "trainee_module_one_name": "Trainee Name",
    "trainee_module_one_surname": "Trainee Surname"
  },
  "module_one_training_created_date": "2023-09-21 01:26:17",
  "module_one_chapter_one": {
    "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
    "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03",
    "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03"
  },
  "set_status_chapter_one_module_one": 0,
  "module_one_chapter_two": {
    "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
    "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03",
    "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03"
  },
  "set_status_chapter_two_module_one": 0,
  "module_one_chapter_three": {
    "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
    "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03",
    "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03"
  },
  "set_status_chapter_three_module_one": 0,
  "training_type_module_one": 0,
  "theoretical_assessment_module_one": 75,
  "practical_assessment_module_one": 0,
  "assessment_module_one": 0,
  "training_status_module_one": 1,
  "assessments_attempts": 0
}
```

## Smart contract function `UpdateStatusChapterOneModuleOneToCompleted`

Assuming that the candidate has passed the theoretical assessment, we can update the status of the chapter one to completed. The command can be passed as follows below.

Request body `POST /transactions`:

```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "Operations Supervisor",
    "channel": "default-channel",
    "chaincode": "competency"
  },
  "func": "UpdateStatusChapterOneModuleOneToCompleted",
  "args": ["proficiency-module-one-ID-0001X"],
  "init": false
}
```

Response body:

```json
{
  "headers": {
    "id": "2bce359c-4ee4-48d1-6b2e-6eec12b00c9d",
    "type": "TransactionSuccess",
    "timeReceived": "2023-09-21T02:22:22.745055457Z",
    "timeElapsed": 0.319503092,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 1125,
  "signerMSP": "u0mk35n6yh",
  "signer": "Operations Supervisor",
  "transactionID": "75590fc6a161c27ebd72f478b75f0aacdfc54ae12206edea17f722ce14a3a24b",
  "status": "VALID"
}
```

Response body in `JSON` format as is recorded on the ledger:

```json
{
  "module_one_id": "proficiency-module-one-ID-0001X",
  "trainer_module_one": {
    "trainer_module_one_id": "Trainer-ID",
    "trainer_module_one_name": "Trainer Name",
    "trainer_module_one_surname": "Trainer Surname"
  },
  "trainee_module_one": {
    "trainee_module_one_id": "Trainee-ID",
    "trainee_module_one_name": "Trainee Name",
    "trainee_module_one_surname": "Trainee Surname"
  },
  "module_one_training_created_date": "2023-09-21 01:26:17",
  "module_one_chapter_one": {
    "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
    "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03",
    "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03"
  },
  "set_status_chapter_one_module_one": 1,
  "module_one_chapter_two": {
    "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
    "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03",
    "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03"
  },
  "set_status_chapter_two_module_one": 0,
  "module_one_chapter_three": {
    "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
    "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03",
    "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03"
  },
  "set_status_chapter_three_module_one": 0,
  "training_type_module_one": 0,
  "theoretical_assessment_module_one": 75,
  "practical_assessment_module_one": 0,
  "assessment_module_one": 0,
  "training_status_module_one": 1,
  "assessments_attempts": 0
}
```

## Smart contract function `UpdateStatusChapterTwoModuleOneToCompleted`

Assuming that the candidate has passed the theoretical assessment, we can update the status of the chapter two to completed. The command can be passed as follows below.

Request body `POST /transactions`:

```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "Operations Supervisor",
    "channel": "default-channel",
    "chaincode": "competency"
  },
  "func": "UpdateStatusChapterTwoModuleOneToCompleted",
  "args": ["proficiency-module-one-ID-0001X"],
  "init": false
}
```

Response body:

```json
{
  "headers": {
    "id": "dd1a795f-fb55-4b23-5018-0a23dab6faaf",
    "type": "TransactionSuccess",
    "timeReceived": "2023-09-21T02:24:49.068730753Z",
    "timeElapsed": 0.289837282,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 1126,
  "signerMSP": "u0mk35n6yh",
  "signer": "Operations Supervisor",
  "transactionID": "91db7b5d873df4f2c69dc29d171f146a15a472489682921b4c7a75e62a9c1c46",
  "status": "VALID"
}
```

Response body in `JSON` format as recorded on the ledger:

```json
{
  "module_one_id": "proficiency-module-one-ID-0001X",
  "trainer_module_one": {
    "trainer_module_one_id": "Trainer-ID",
    "trainer_module_one_name": "Trainer Name",
    "trainer_module_one_surname": "Trainer Surname"
  },
  "trainee_module_one": {
    "trainee_module_one_id": "Trainee-ID",
    "trainee_module_one_name": "Trainee Name",
    "trainee_module_one_surname": "Trainee Surname"
  },
  "module_one_training_created_date": "2023-09-21 01:26:17",
  "module_one_chapter_one": {
    "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
    "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03",
    "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03"
  },
  "set_status_chapter_one_module_one": 1,
  "module_one_chapter_two": {
    "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
    "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03",
    "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03"
  },
  "set_status_chapter_two_module_one": 1,
  "module_one_chapter_three": {
    "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
    "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03",
    "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03"
  },
  "set_status_chapter_three_module_one": 0,
  "training_type_module_one": 0,
  "theoretical_assessment_module_one": 75,
  "practical_assessment_module_one": 0,
  "assessment_module_one": 0,
  "training_status_module_one": 1,
  "assessments_attempts": 0
}
```

## Smart contract function `UpdateStatusChapterThreeModuleOneToCompleted`

Assuming that the candidate has passed the theoretical assessment, we can update the status of the chapter three to completed. The command can be passed as follows below.

Request body `POST /transactions`:

```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "Operations Supervisor",
    "channel": "default-channel",
    "chaincode": "competency"
  },
  "func": "UpdateStatusChapterThreeModuleOneToCompleted",
  "args": ["proficiency-module-one-ID-0001X"],
  "init": false
}
```

Response body:

```json
{
  "headers": {
    "id": "755066c2-867f-43dd-6bbb-89d985af4737",
    "type": "TransactionSuccess",
    "timeReceived": "2023-09-21T02:27:47.939742325Z",
    "timeElapsed": 0.3231204,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 1127,
  "signerMSP": "u0mk35n6yh",
  "signer": "Operations Supervisor",
  "transactionID": "044b6d087b6cd2ac5b1e6ade5b58a781593172f430b715de38ad460f88a570e1",
  "status": "VALID"
}
```

Response body in `JSON` format as recorded on the ledger:

```json
{
  "module_one_id": "proficiency-module-one-ID-0001X",
  "trainer_module_one": {
    "trainer_module_one_id": "Trainer-ID",
    "trainer_module_one_name": "Trainer Name",
    "trainer_module_one_surname": "Trainer Surname"
  },
  "trainee_module_one": {
    "trainee_module_one_id": "Trainee-ID",
    "trainee_module_one_name": "Trainee Name",
    "trainee_module_one_surname": "Trainee Surname"
  },
  "module_one_training_created_date": "2023-09-21 01:26:17",
  "module_one_chapter_one": {
    "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
    "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03",
    "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03"
  },
  "set_status_chapter_one_module_one": 1,
  "module_one_chapter_two": {
    "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
    "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03",
    "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03"
  },
  "set_status_chapter_two_module_one": 1,
  "module_one_chapter_three": {
    "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
    "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03",
    "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03"
  },
  "set_status_chapter_three_module_one": 1,
  "training_type_module_one": 0,
  "theoretical_assessment_module_one": 75,
  "practical_assessment_module_one": 0,
  "assessment_module_one": 0,
  "training_status_module_one": 1,
  "assessments_attempts": 0
}
```

## Samrt contract function `UpdateValuePracticalAndTotalAssessment`

Request body `POST /transactions`

```json
{
  "headers": {
    "type": "SendTransaction",
    "signer": "Operations Supervisor",
    "channel": "default-channel",
    "chaincode": "competency"
  },
  "func": "UpdateValuePracticalAndTotalAssessment",
  "args": ["proficiency-module-one-ID-0001X", "100"],
  "init": false
}
```

Response body:

```json
{
  "headers": {
    "id": "82099053-c029-4086-5b5b-a3a3ffdbc1e9",
    "type": "TransactionSuccess",
    "timeReceived": "2023-09-21T02:41:03.841520739Z",
    "timeElapsed": 0.325246392,
    "requestOffset": "",
    "requestId": ""
  },
  "blockNumber": 1128,
  "signerMSP": "u0mk35n6yh",
  "signer": "Operations Supervisor",
  "transactionID": "b5c4375a5699493e1ebef40ac5242a9458fc2f752c0d41acd640ba630834e8d0",
  "status": "VALID"
}
```

Response body in `JSON` formate as recorded on the ledger:

```json
{
  "module_one_id": "proficiency-module-one-ID-0001X",
  "trainer_module_one": {
    "trainer_module_one_id": "Trainer-ID",
    "trainer_module_one_name": "Trainer Name",
    "trainer_module_one_surname": "Trainer Surname"
  },
  "trainee_module_one": {
    "trainee_module_one_id": "Trainee-ID",
    "trainee_module_one_name": "Trainee Name",
    "trainee_module_one_surname": "Trainee Surname"
  },
  "module_one_training_created_date": "2023-09-21 01:26:17",
  "module_one_chapter_one": {
    "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
    "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03",
    "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03"
  },
  "set_status_chapter_one_module_one": 1,
  "module_one_chapter_two": {
    "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
    "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03",
    "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03"
  },
  "set_status_chapter_two_module_one": 1,
  "module_one_chapter_three": {
    "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
    "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03",
    "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03"
  },
  "set_status_chapter_three_module_one": 1,
  "training_type_module_one": 0,
  "theoretical_assessment_module_one": 75,
  "practical_assessment_module_one": 100,
  "assessment_module_one": 175,
  "training_status_module_one": 3,
  "assessments_attempts": 1
}
```

## Smart contract function `GetHistoryRecordsForModuleOne`

Request body `POST /query`:

```json
{
  "headers": {
    "type": "Query",
    "signer": "Operations Supervisor",
    "channel": "default-channel",
    "chaincode": "competency"
  },
  "func": "GetHistoryRecordsForModuleOne",
  "args": ["proficiency-module-one-ID-0001X"],
  "init": false
}
```

Response body in `JSON` format afte querying the ledger:

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
        "assessment_module_one": 175,
        "assessments_attempts": 1,
        "module_one_chapter_one": {
          "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
          "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03",
          "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03"
        },
        "module_one_chapter_three": {
          "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
          "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03",
          "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03"
        },
        "module_one_chapter_two": {
          "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
          "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03",
          "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03"
        },
        "module_one_id": "proficiency-module-one-ID-0001X",
        "module_one_training_created_date": "2023-09-21 01:26:17",
        "practical_assessment_module_one": 100,
        "set_status_chapter_one_module_one": 1,
        "set_status_chapter_three_module_one": 1,
        "set_status_chapter_two_module_one": 1,
        "theoretical_assessment_module_one": 75,
        "trainee_module_one": {
          "trainee_module_one_id": "Trainee-ID",
          "trainee_module_one_name": "Trainee Name",
          "trainee_module_one_surname": "Trainee Surname"
        },
        "trainer_module_one": {
          "trainer_module_one_id": "Trainer-ID",
          "trainer_module_one_name": "Trainer Name",
          "trainer_module_one_surname": "Trainer Surname"
        },
        "training_status_module_one": 3,
        "training_type_module_one": 0
      },
      "timestamp": "2023-09-21T02:41:03.841804208Z",
      "txId": "b5c4375a5699493e1ebef40ac5242a9458fc2f752c0d41acd640ba630834e8d0"
    },
    {
      "isDelete": false,
      "record": {
        "assessment_module_one": 0,
        "assessments_attempts": 0,
        "module_one_chapter_one": {
          "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
          "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03",
          "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03"
        },
        "module_one_chapter_three": {
          "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
          "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03",
          "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03"
        },
        "module_one_chapter_two": {
          "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
          "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03",
          "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03"
        },
        "module_one_id": "proficiency-module-one-ID-0001X",
        "module_one_training_created_date": "2023-09-21 01:26:17",
        "practical_assessment_module_one": 0,
        "set_status_chapter_one_module_one": 1,
        "set_status_chapter_three_module_one": 1,
        "set_status_chapter_two_module_one": 1,
        "theoretical_assessment_module_one": 75,
        "trainee_module_one": {
          "trainee_module_one_id": "Trainee-ID",
          "trainee_module_one_name": "Trainee Name",
          "trainee_module_one_surname": "Trainee Surname"
        },
        "trainer_module_one": {
          "trainer_module_one_id": "Trainer-ID",
          "trainer_module_one_name": "Trainer Name",
          "trainer_module_one_surname": "Trainer Surname"
        },
        "training_status_module_one": 1,
        "training_type_module_one": 0
      },
      "timestamp": "2023-09-21T02:27:47.940017917Z",
      "txId": "044b6d087b6cd2ac5b1e6ade5b58a781593172f430b715de38ad460f88a570e1"
    },
    {
      "isDelete": false,
      "record": {
        "assessment_module_one": 0,
        "assessments_attempts": 0,
        "module_one_chapter_one": {
          "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
          "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03",
          "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03"
        },
        "module_one_chapter_three": {
          "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
          "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03",
          "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03"
        },
        "module_one_chapter_two": {
          "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
          "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03",
          "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03"
        },
        "module_one_id": "proficiency-module-one-ID-0001X",
        "module_one_training_created_date": "2023-09-21 01:26:17",
        "practical_assessment_module_one": 0,
        "set_status_chapter_one_module_one": 1,
        "set_status_chapter_three_module_one": 0,
        "set_status_chapter_two_module_one": 1,
        "theoretical_assessment_module_one": 75,
        "trainee_module_one": {
          "trainee_module_one_id": "Trainee-ID",
          "trainee_module_one_name": "Trainee Name",
          "trainee_module_one_surname": "Trainee Surname"
        },
        "trainer_module_one": {
          "trainer_module_one_id": "Trainer-ID",
          "trainer_module_one_name": "Trainer Name",
          "trainer_module_one_surname": "Trainer Surname"
        },
        "training_status_module_one": 1,
        "training_type_module_one": 0
      },
      "timestamp": "2023-09-21T02:24:49.06897534Z",
      "txId": "91db7b5d873df4f2c69dc29d171f146a15a472489682921b4c7a75e62a9c1c46"
    },
    {
      "isDelete": false,
      "record": {
        "assessment_module_one": 0,
        "assessments_attempts": 0,
        "module_one_chapter_one": {
          "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
          "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03",
          "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03"
        },
        "module_one_chapter_three": {
          "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
          "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03",
          "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03"
        },
        "module_one_chapter_two": {
          "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
          "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03",
          "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03"
        },
        "module_one_id": "proficiency-module-one-ID-0001X",
        "module_one_training_created_date": "2023-09-21 01:26:17",
        "practical_assessment_module_one": 0,
        "set_status_chapter_one_module_one": 1,
        "set_status_chapter_three_module_one": 0,
        "set_status_chapter_two_module_one": 0,
        "theoretical_assessment_module_one": 75,
        "trainee_module_one": {
          "trainee_module_one_id": "Trainee-ID",
          "trainee_module_one_name": "Trainee Name",
          "trainee_module_one_surname": "Trainee Surname"
        },
        "trainer_module_one": {
          "trainer_module_one_id": "Trainer-ID",
          "trainer_module_one_name": "Trainer Name",
          "trainer_module_one_surname": "Trainer Surname"
        },
        "training_status_module_one": 1,
        "training_type_module_one": 0
      },
      "timestamp": "2023-09-21T02:22:22.74530218Z",
      "txId": "75590fc6a161c27ebd72f478b75f0aacdfc54ae12206edea17f722ce14a3a24b"
    },
    {
      "isDelete": false,
      "record": {
        "assessment_module_one": 0,
        "assessments_attempts": 0,
        "module_one_chapter_one": {
          "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
          "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03",
          "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03"
        },
        "module_one_chapter_three": {
          "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
          "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03",
          "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03"
        },
        "module_one_chapter_two": {
          "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
          "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03",
          "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03"
        },
        "module_one_id": "proficiency-module-one-ID-0001X",
        "module_one_training_created_date": "2023-09-21 01:26:17",
        "practical_assessment_module_one": 0,
        "set_status_chapter_one_module_one": 0,
        "set_status_chapter_three_module_one": 0,
        "set_status_chapter_two_module_one": 0,
        "theoretical_assessment_module_one": 75,
        "trainee_module_one": {
          "trainee_module_one_id": "Trainee-ID",
          "trainee_module_one_name": "Trainee Name",
          "trainee_module_one_surname": "Trainee Surname"
        },
        "trainer_module_one": {
          "trainer_module_one_id": "Trainer-ID",
          "trainer_module_one_name": "Trainer Name",
          "trainer_module_one_surname": "Trainer Surname"
        },
        "training_status_module_one": 1,
        "training_type_module_one": 0
      },
      "timestamp": "2023-09-21T02:15:28.717109801Z",
      "txId": "24670a5bac7be204fde289d063de5350415674e39fe271e3a975d44d1bd40f22"
    },
    {
      "isDelete": false,
      "record": {
        "assessment_module_one": 0,
        "assessments_attempts": 0,
        "module_one_chapter_one": {
          "module_one_chapter_one_session_one": "Module One - Chapter 01 - Seesion 03",
          "module_one_chapter_one_session_three": "Module One - Chapter 01 - Seesion 03",
          "module_one_chapter_one_session_two": "Module One - Chapter 01 - Seesion 03"
        },
        "module_one_chapter_three": {
          "module_one_chapter_three_session_one": "Module One - Chapter 03 - Seesion 03",
          "module_one_chapter_three_session_three": "Module One - Chapter 03 - Seesion 03",
          "module_one_chapter_three_session_two": "Module One - Chapter 03 - Seesion 03"
        },
        "module_one_chapter_two": {
          "module_one_chapter_two_session_one": "Module One - Chapter 02 - Seesion 03",
          "module_one_chapter_two_session_three": "Module One - Chapter 02 - Seesion 03",
          "module_one_chapter_two_session_two": "Module One - Chapter 02 - Seesion 03"
        },
        "module_one_id": "proficiency-module-one-ID-0001X",
        "module_one_training_created_date": "2023-09-21 01:26:17",
        "practical_assessment_module_one": 0,
        "set_status_chapter_one_module_one": 0,
        "set_status_chapter_three_module_one": 0,
        "set_status_chapter_two_module_one": 0,
        "theoretical_assessment_module_one": 0,
        "trainee_module_one": {
          "trainee_module_one_id": "Trainee-ID",
          "trainee_module_one_name": "Trainee Name",
          "trainee_module_one_surname": "Trainee Surname"
        },
        "trainer_module_one": {
          "trainer_module_one_id": "Trainer-ID",
          "trainer_module_one_name": "Trainer Name",
          "trainer_module_one_surname": "Trainer Surname"
        },
        "training_status_module_one": 1,
        "training_type_module_one": 0
      },
      "timestamp": "2023-09-21T01:26:17.815093684Z",
      "txId": "014e4c5b464b3c44a07f00f5b1b456ee78692ada4681e29b83b738019d02616e"
    }
  ]
}
```

[Back on top](#testing-platforms)

# Caliper Testing

The certificate must be adjusted after running the test network. The code below is shown where we can take the MSP credentials, and adjust them into the [`test-network.yaml`](../../../02-benchmarking-files/hyperledger-caliper/caliper-benchmarks/networks/fabric/test-network.yaml) file.

The location of the `test-network.yaml` file can be reached by the command:

```bash
cd /home/vodenica/go/src/github.com/hyperledger/caliper-benchmarks/networks/fabric
```

The credentials to be picked up from this directory for `User1` using `ls` command:

```bash
cd /home/vodenica/go/src/github.com/hyperledger/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore
```

Ensure you are in the `caliper-benchmarks` directory. If not, navigate to the `caliper-benchmarks` directory:

```bash
cd /home/vodenica/go/src/github.com/hyperledger/caliper-benchmarks
```

Execution:

```bash
npx caliper launch manager --caliper-workspace ./ --caliper-networkconfig networks/fabric/test-network.yaml --caliper-benchconfig benchmarks/samples/fabric/proficiency/config.yaml --caliper-flow-only-test --caliper-fabric-gateway-enabled
```

[Back on top](#testing-platforms)
