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
	// "log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {

	/*

		chaincode, err := contractapi.NewChaincode(&CompetencySmartContractModuleOne{})

		if err != nil {
			log.Panicf("Error creating competency Module One chaincode: %v", err)
		}

		if err := chaincode.Start(); err != nil {
			log.Panicf("Error starting competency Module One chaincode: %v", err)
		}
	*/

	complexContract := new(CompetencySmartContractModuleOne)
	complexContract.TransactionContextHandler = new(CustomTransactionContext)
	complexContract.BeforeTransaction = GetWorldState

	cc, err := contractapi.NewChaincode(complexContract)

	if err != nil {
		panic(err.Error())
	}

	if err := cc.Start(); err != nil {
		panic(err.Error())
	}

}
