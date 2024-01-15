package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	//chaincode, err := contractapi.NewChaincode(&PreventiveMaintenanceChaincode{}, &CorrectiveMaintenanceChaincode{})
	chaincode, err := contractapi.NewChaincode(
		&EmployeeSmartContract{})
	if err != nil {
		log.Panicf("Error creating maintenance work order chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting maintenance work order chaincode: %v", err)
	}
}
