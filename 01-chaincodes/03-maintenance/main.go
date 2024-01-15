package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	chaincode, err := contractapi.NewChaincode(&MaintenanceSmartContract{})

	if err != nil {
		log.Panicf("error creating maintenance work order chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("error starting maintenance work order chaincode: %v", err)
	}

	/*
	chaincode2, err2 := contractapi.NewChaincode(&CorrectiveMaintenanceSmartContract{})

	if err2 != nil {
		log.Panicf("error creating maintenance work order chaincode: %v", err2)
	}

	if err := chaincode2.Start(); err != nil {
		log.Panicf("error starting maintenance work order chaincode: %v", err)
	}

	chaincode3, err3 := contractapi.NewChaincode(&FailureMaintenanceSmartContract{})

	if err3 != nil {
		log.Panicf("error creating maintenance work order chaincode: %v", err3)
	}

	if err := chaincode3.Start(); err != nil {
		log.Panicf("error starting maintenance work order chaincode: %v", err)
	}
	*/

}
