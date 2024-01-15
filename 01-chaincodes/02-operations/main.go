/*
 * Copyright 2022 IBM All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the 'License');
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an 'AS IS' BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//  Create a line of code that assigns a value to a version-related variable and update it as needed.
// 	Example: const version = "1.0"

const (
	chaincodeName = "daily_ops_log"
	version       = "1.0" // Update this version number as needed
)

func main() {
	chaincode, err := contractapi.NewChaincode(
		&DailyOperationsLogSmartContract{},
	)
	if err != nil {
		log.Panicf("Error creating maintenance work order chaincode: %v", err)
	}

	chaincode.Info.Title = "Daily Operations Log Smart Contract"
	chaincode.Info.Version = version
	chaincode.Info.Description = "Smart contract for creating and updating daily operations logs"

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting maintenance work order chaincode: %v", err)
	}
}
