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
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	// "github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const indexEmployeePosition = "employee~position"

// SmartContract provides functions for managing Employees records, in this case we will represent
// and Employee as an Asset (as it trully is).
// Also, it is important to underline, that the MSP for this type of "asset"
// is "org-operations-com". That said, the access control is allowed only through "org-operations-com" MSP.

//================================================================================================================================================
// SMART CONTRACT STRUCT
//================================================================================================================================================

type EmployeeSmartContract struct {
	contractapi.Contract             // embedded contractapi.Contract
	StatusTrainingModuleOne   string // define the StatusTrainingModuleOne field as a string
	StatusTrainingModuleTwo   string // define the StatusTrainingModuleTwo field as a string
	StatusTrainingModuleThree string // define the StatusTrainingModuleThree field as a string
}

// HistoryQueryResult structure used for returning result of history query
type HistoryQueryResultEmployee struct {
	Record    *Employee `json:"record"`
	TxId      string    `json:"txId"`
	Timestamp time.Time `json:"timestamp"`
	IsDelete  bool      `json:"isDelete"`
}

//================================================================================================================================================
// InitLedger
//================================================================================================================================================

// InitLedger adds records of the employees to the ledger
func (s *EmployeeSmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	assets := []Employee{
		{
			DocType:         "employee",
			CompanySite:     "Macao",
			ID:              "Pass-11111111",
			Name:            "Dalibor",
			LastName:        "Vodenicarski",
			Gender:          "M",
			PositionEntry:   "Maintenance Supervisor",
			CurrentPosition: "General Manager",
			LevelPT:         "Employed",
			Birthday:        "January 1, 1979",
			ContractSigned:  "November 1, 2015",
			StartingDate:    "",
			YearsInService:  "7 years",
			EmployeeReview:  "Very Good",
			CustomTraining:  "",
			Salary:          35000,
			Address:         "Shun Tak, Macao",
			Cell:            "(853)62123456",
			WorkVisa:        "2020-2022",
		},
		{
			DocType:         "employee",
			CompanySite:     "Macao",
			ID:              "MID-12345",
			Name:            "Alice",
			LastName:        "Irving",
			Gender:          "F",
			PositionEntry:   "System Operator",
			CurrentPosition: "System Operator",
			LevelPT:         "Employed",
			Birthday:        "January 15, 1998",
			ContractSigned:  "November 1, 2022",
			StartingDate:    "",
			YearsInService:  "0.5 years",
			EmployeeReview:  "Good",
			CustomTraining:  "",
			Salary:          10000,
			Address:         "Sec Pai Van, One Oasis Tower 2, 28/B",
			Cell:            "(853)63111222",
			WorkVisa:        "n/a",
		},
		{
			DocType:         "employee",
			CompanySite:     "Macao",
			ID:              "MID-678910",
			Name:            "Leyton",
			LastName:        "Wyatt",
			Gender:          "M",
			PositionEntry:   "Maintenance Technician",
			CurrentPosition: "Maintenance Technician",
			LevelPT:         "Employed",
			Birthday:        "September 10, 1989",
			ContractSigned:  "November 15, 2020",
			StartingDate:    "",
			YearsInService:  "2 years",
			EmployeeReview:  "Very Good",
			CustomTraining:  "",
			Salary:          20000,
			Address:         "Rua do Campo, Build 2, Andar 5/22",
			Cell:            "(853)63222333",
			WorkVisa:        "n/a",
		},
		{
			DocType:         "employee",
			CompanySite:     "Macao",
			ID:              "MID-112233",
			Name:            "Haitao",
			LastName:        "Zhāng",
			Gender:          "M",
			PositionEntry:   "Maintenance Technician",
			CurrentPosition: "Maintenance Technician",
			LevelPT:         "Employed",
			Birthday:        "October 1, 1992",
			ContractSigned:  "November 1, 2018",
			StartingDate:    "",
			YearsInService:  "4 years",
			EmployeeReview:  "Very Good",
			CustomTraining:  "",
			Salary:          25000,
			Address:         "Praca Ponte E Horta, Tower 5, And. 2/5",
			Cell:            "(853)63444555",
			WorkVisa:        "n/a",
		},
		{
			DocType:         "employee",
			CompanySite:     "Macao",
			ID:              "MID-445566",
			Name:            "Ronaldo",
			LastName:        "Gomes Carvalho",
			Gender:          "M",
			PositionEntry:   "Maintenance Supervisor",
			CurrentPosition: "Maintenance Manager",
			LevelPT:         "Employed",
			Birthday:        "December 1, 1985",
			ContractSigned:  "August 2, 2017",
			StartingDate:    "",
			YearsInService:  "5 years",
			EmployeeReview:  "Very Good",
			CustomTraining:  "",
			Salary:          30000,
			Address:         "Sec Pai Van, One Oasis Tower 5, 22/F",
			Cell:            "(853)63666777",
			WorkVisa:        "n/a",
		},
		{
			DocType:         "employee",
			CompanySite:     "Macao",
			ID:              "MID-778899",
			Name:            "Ernesto",
			LastName:        "Eugenio",
			Gender:          "M",
			PositionEntry:   "System Operator",
			CurrentPosition: "System Operator",
			LevelPT:         "Employed",
			Birthday:        "June 16, 1999",
			ContractSigned:  "May 1, 2019",
			StartingDate:    "",
			YearsInService:  "3 years",
			EmployeeReview:  "Poor Performance",
			CustomTraining:  "",
			Salary:          15000,
			Address:         "NAPE, Edf. Londres 22/2",
			Cell:            "(853)63888777",
			WorkVisa:        "n/a",
		},
	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

//================================================================================================================================================
// ADD NEW EMPLOYEE
//================================================================================================================================================

// AddNewEmployee issues a new record of a new employee to the world state with given details.
func (s *EmployeeSmartContract) AddNewEmployee(ctx contractapi.TransactionContextInterface, companySite string, id string, name string, lastName string, gender string, positionEntry string, currentPosition string, levelPT string, birthday string, contractSigned string, startingDate string, yearsInService string, employeeReview string, customTraining string, salary int, address string, cell string, workVisa string) error {

	exists, err := s.EmployeeExists(ctx, id)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	// Declaring the existing variable 'em' of Employee struct type
	employee := Employee{
		DocType:         "employee",
		CompanySite:     companySite,
		ID:              id,
		Name:            name,
		LastName:        lastName,
		Gender:          gender,
		PositionEntry:   positionEntry,
		CurrentPosition: currentPosition,
		LevelPT:         levelPT,
		Birthday:        birthday,
		ContractSigned:  contractSigned,
		StartingDate:    startingDate,
		YearsInService:  yearsInService,
		EmployeeReview:  employeeReview,
		CustomTraining:  customTraining,
		Salary:          salary,
		Address:         address,
		Cell:            cell,
		WorkVisa:        workVisa,
	}

	// Set the StatusTrainingModuleOne field to "Open" for the new employee
	employee.SetTrainingModuleOneOpen()

	// Set the StatusTrainingModuleTwo field to "Open" for the new employee
	employee.SetTrainingModuleTwoOpen()

	// Set the StatusTrainingModuleThree field to "Open" for the new employee
	employee.SetTrainingModuleThreeOpen()

	assetBytes, err := json.Marshal(employee)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(id, assetBytes)
	if err != nil {
		return fmt.Errorf("failed to put to world state: %v", err)
	}

	// Create the composite key that will allow us to query for all employees by ID
	// The key is a combination of the prefix 'employee' and the employee ID
	employeeCompositeKey, err := ctx.GetStub().CreateCompositeKey(indexEmployeePosition, []string{employee.ID, employee.CurrentPosition})
	if err != nil {
		return fmt.Errorf("failed to create the composite key for the new employee: %v", err)
	}

	// Save the composite key index
	value := []byte{0x00}

	// Save the employee to the world state
	return ctx.GetStub().PutState(employeeCompositeKey, value)
}

//================================================================================================================================================
// READ EMPLOYEE
//================================================================================================================================================

// ReadEmployee returns the asset stored in the world state with given id.
func (s *EmployeeSmartContract) ReadEmployee(ctx contractapi.TransactionContextInterface, id string) (*Employee, error) {

	assetJSON, err := ctx.GetStub().GetState(id)

	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}

	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var employee Employee
	err = json.Unmarshal(assetJSON, &employee)
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

//================================================================================================================================================
// UPDATE EMPLOYEE STATUS FOR MODULE ONE TRAINING INTO "IN PROGRESS"
//================================================================================================================================================

// UpdateEmployeeStatusForModuleOneTrainingInProgress updates an existing asset in the world state with provided parameters.
func (s *EmployeeSmartContract) UpdateEmployeeStatusForModuleOneTrainingInProgress(ctx contractapi.TransactionContextInterface, id string, newStatusModuelOne string) (string, error) {

	// retrieve the Employee struct from the world state
	employee, err := s.ReadEmployee(ctx, id)
	if err != nil {
		return "", err
	}

	oldStatus := employee.StatusTrainingModuleOne

	// update the status of the training module one to in progress
	employee.SetTrainingModuleOneInProgress()

	employee.StatusTrainingModuleOne = newStatusModuelOne

	assetJSON, err := json.Marshal(employee)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(id, assetJSON)
	if err != nil {
		return "", err
	}

	return oldStatus, nil
}

//================================================================================================================================================
// UPDATE EMPLOYEE STATUS FOR MODULE TWO TRAINING INTO "IN PROGRESS"
//================================================================================================================================================

// UpdateEmployeeStatusForModuleTwoTrainingInProgress updates an existing asset in the world state with provided parameters.

func (s *EmployeeSmartContract) UpdateEmployeeStatusForModuleTwoTrainingInProgress(ctx contractapi.TransactionContextInterface, id string, newStatusModuelTwo string) (string, error) {

	// retrieve the Employee struct from the world state
	employee, err := s.ReadEmployee(ctx, id)
	if err != nil {
		return "", err
	}

	oldStatus := employee.StatusTrainingModuleTwo
	// update the status of the training module two to in progress
	employee.SetTrainingModuleTwoInProgress()

	employee.StatusTrainingModuleTwo = newStatusModuelTwo

	assetJSON, err := json.Marshal(employee)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(id, assetJSON)
	if err != nil {
		return "", err
	}

	return oldStatus, nil
}

//================================================================================================================================================
// UPDATE EMPLOYEE STATUS FOR MODULE THREE TRAINING INTO "IN PROGRESS"
//================================================================================================================================================

// UpdateEmployeeStatusForModuleThreeTrainingInProgress updates an existing asset in the world state with provided parameters.

func (s *EmployeeSmartContract) UpdateEmployeeStatusForModuleThreeTrainingInProgress(ctx contractapi.TransactionContextInterface, id string, newStatusModuelThree string) (string, error) {

	// retrieve the Employee struct from the world state
	employee, err := s.ReadEmployee(ctx, id)
	if err != nil {
		return "", err
	}

	oldStatus := employee.StatusTrainingModuleThree
	// update the status of the training module three to in progress

	employee.SetTrainingModuleThreeInProgress()

	employee.StatusTrainingModuleThree = newStatusModuelThree

	assetJSON, err := json.Marshal(employee)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(id, assetJSON)
	if err != nil {
		return "", err
	}

	return oldStatus, nil
}

//================================================================================================================================================
// UPDATE EMPLOYEE STATUS FOR MODULE ONE TRAINING INTO "COMPLETED"
//================================================================================================================================================

// UpdateEmployeeStatusForModuleOneTrainingCompleted updates an existing asset in the world state with provided parameters.

func (s *EmployeeSmartContract) UpdateEmployeeStatusForModuleOneTrainingCompleted(ctx contractapi.TransactionContextInterface, id string, newStatusModuelOne string) (string, error) {

	// retrieve the Employee struct from the world state
	employee, err := s.ReadEmployee(ctx, id)
	if err != nil {
		return "", err
	}

	oldStatus := employee.StatusTrainingModuleOne
	// update the status of the training module one to completed

	employee.SetTrainingModuleOneCompleted()

	employee.StatusTrainingModuleOne = newStatusModuelOne

	assetJSON, err := json.Marshal(employee)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(id, assetJSON)
	if err != nil {
		return "", err
	}

	return oldStatus, nil
}

//================================================================================================================================================
// UPDATE EMPLOYEE STATUS FOR MODULE TWO TRAINING INTO "COMPLETED"
//================================================================================================================================================

// UpdateEmployeeStatusForModuleTwoTrainingCompleted updates an existing asset in the world state with provided parameters.

func (s *EmployeeSmartContract) UpdateEmployeeStatusForModuleTwoTrainingCompleted(ctx contractapi.TransactionContextInterface, id string, newStatusModuelTwo string) (string, error) {

	// retrieve the Employee struct from the world state
	employee, err := s.ReadEmployee(ctx, id)
	if err != nil {
		return "", err
	}

	oldStatus := employee.StatusTrainingModuleTwo
	// update the status of the training module two to completed

	employee.SetTrainingModuleTwoCompleted()

	employee.StatusTrainingModuleTwo = newStatusModuelTwo

	assetJSON, err := json.Marshal(employee)

	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(id, assetJSON)

	if err != nil {
		return "", err
	}

	return oldStatus, nil
}

//================================================================================================================================================
// UPDATE EMPLOYEE STATUS FOR MODULE THREE TRAINING INTO "COMPLETED"
//================================================================================================================================================

// UpdateEmployeeStatusForModuleThreeTrainingCompleted updates an existing asset in the world state with provided parameters.

func (s *EmployeeSmartContract) UpdateEmployeeStatusForModuleThreeTrainingCompleted(ctx contractapi.TransactionContextInterface, id string, newStatusModuelThree string) (string, error) {

	// retrieve the Employee struct from the world state
	employee, err := s.ReadEmployee(ctx, id)
	if err != nil {
		return "", err
	}

	oldStatus := employee.StatusTrainingModuleThree
	// update the status of the training module three to completed

	employee.SetTrainingModuleThreeCompleted()

	employee.StatusTrainingModuleThree = newStatusModuelThree

	assetJSON, err := json.Marshal(employee)

	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(id, assetJSON)

	if err != nil {
		return "", err
	}

	return oldStatus, nil
}

//================================================================================================================================================
// UPDATE EMPLOYEE
//================================================================================================================================================

// UpdateEmployee updates an existing asset in the world state with provided parameters.
func (s *EmployeeSmartContract) UpdateEmployee(ctx contractapi.TransactionContextInterface, companySite string, id string, name string, lastName string, gender string, positionEntry string, currentPosition string, levelPT string, birthday string, contractSigned string, startingDate string, yearsInService string, employeeReview string, salary int, address string, cell string, workVisa string) error {

	exists, err := s.EmployeeExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	// Overwriting original asset with new asset
	asset := Employee{
		DocType:         "employee",
		CompanySite:     companySite,
		ID:              id,
		Name:            name,
		LastName:        lastName,
		Gender:          gender,
		PositionEntry:   positionEntry,
		CurrentPosition: currentPosition,
		LevelPT:         levelPT,
		Birthday:        birthday,
		ContractSigned:  contractSigned,
		StartingDate:    startingDate,
		YearsInService:  yearsInService,
		EmployeeReview:  employeeReview,
		Salary:          salary,
		Address:         address,
		Cell:            cell,
		WorkVisa:        workVisa,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

//================================================================================================================================================
// DELETE EMPLOYEE RECORDS
//================================================================================================================================================

// DeleteAsset deletes an given asset from the world state.
// The function is working, adding in "Swagger" just an employee ID.
func (s *EmployeeSmartContract) DeleteEmployeeRecords(ctx contractapi.TransactionContextInterface, id string) error {

	exists, err := s.EmployeeExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

//================================================================================================================================================
// EMPLOYEE EXISTS
//================================================================================================================================================

// AssetExists returns true when asset with given ID exists in world state
func (s *EmployeeSmartContract) EmployeeExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {

	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

//================================================================================================================================================
// TRANSFER OF THE EMPLOYEE
//================================================================================================================================================

// TransferEmployee updates the employee Company site field of the asset with the given id in the world state and returns the old owner.
// Transfer the asset; this could be when Employee is temporarily dislocated to another Company site.
func (s *EmployeeSmartContract) TransferEmployee(ctx contractapi.TransactionContextInterface, id string, newCompanySite string) (string, error) {

	asset, err := s.ReadEmployee(ctx, id)
	if err != nil {
		return "", err
	}

	oldCompanySite := asset.CompanySite
	asset.CompanySite = newCompanySite

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(id, assetJSON)
	if err != nil {
		return "", err
	}

	return oldCompanySite, nil
}

//================================================================================================================================================
// GET ALL EMPLOYEES
//================================================================================================================================================

// GetAllEmployees returns all assets found in world state
func (s *EmployeeSmartContract) GetAllEmployees(ctx contractapi.TransactionContextInterface) ([]*Employee, error) {

	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIteratorEmployee, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIteratorEmployee.Close()

	var assets []*Employee
	for resultsIteratorEmployee.HasNext() {
		queryResponse, err := resultsIteratorEmployee.Next()
		if err != nil {
			return nil, err
		}

		var asset Employee
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}

		assets = append(assets, &asset)
	}

	return assets, nil
}

//================================================================================================================================================
// QUERY EMPLOYEE BY QUERY STRING
//================================================================================================================================================

// QueryEmployeeByQueryString uses a query string to perform a query for work orders.
// Query string matching state database syntax is passed in and executed as is.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the "QueryAssetsForOwner" example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Ad hoc rich query
func (s *EmployeeSmartContract) QueryEmployeeByQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]*Employee, error) {

	return getQueryResultForQueryStringEmployee(ctx, queryString)
}

//================================================================================================================================================
// GET QUERY RESULT FOR QUERY STRING
//================================================================================================================================================

// getQueryResultForQueryStringEmployee executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryStringEmployee(ctx contractapi.TransactionContextInterface, queryString string) ([]*Employee, error) {
	resultsIteratorEmployee, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorEmployee.Close()

	return constructQueryResponseFromIteratorEmployee(resultsIteratorEmployee)
}

//================================================================================================================================================
// CONSTRUCT QUERY RESPONSE FROM ITERATOR
//================================================================================================================================================

// constructQueryResponseFromIteratorEmployee constructs a slice of assets from the resultsIteratorPrev
func constructQueryResponseFromIteratorEmployee(resultsIteratorEmployee shim.StateQueryIteratorInterface) ([]*Employee, error) {
	var assets []*Employee
	for resultsIteratorEmployee.HasNext() {
		queryResult, err := resultsIteratorEmployee.Next()
		if err != nil {
			return nil, err
		}
		var employee Employee
		err = json.Unmarshal(queryResult.Value, &employee)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &employee)
	}

	return assets, nil
}

//================================================================================================================================================
// GET EMPLOYEE HISTORY
//================================================================================================================================================

// GetEmployeeHistory returns the chain of custody for an asset since issuance.
func (s *EmployeeSmartContract) GetEmployeeHistory(ctx contractapi.TransactionContextInterface, id string) ([]HistoryQueryResultEmployee, error) {

	log.Printf("GetAssetHistory: ID %v", id)

	resultsIteratorEmployee, err := ctx.GetStub().GetHistoryForKey(id)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorEmployee.Close()

	var records []HistoryQueryResultEmployee
	for resultsIteratorEmployee.HasNext() {
		response, err := resultsIteratorEmployee.Next()
		if err != nil {
			return nil, err
		}

		var employee Employee
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &employee)
			if err != nil {
				return nil, err
			}
		} else {
			employee = Employee{
				ID: id,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			return nil, err
		}

		record := HistoryQueryResultEmployee{
			TxId:      response.TxId,
			Timestamp: timestamp,
			Record:    &employee,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}

//================================================================================================================================================
// PROMOTION OF THE EMPLOYEE
//================================================================================================================================================

// PromoteEmployee updates the employee Current Position field of the "digital asset" with the given id in the world state and returns the previous
// Current Position.
// Promote the employee (team member - TM); When employee is completing the Competency Module 1/2/3 his/her ledger information's will be updated accordingly.
func (s *EmployeeSmartContract) PromoteEmployee(ctx contractapi.TransactionContextInterface, id string, newCurrentPosition string) (string, error) {

	asset, err := s.ReadEmployee(ctx, id)
	if err != nil {
		return "", err
	}

	previousCurrentPosition := asset.CurrentPosition
	asset.CurrentPosition = newCurrentPosition

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(id, assetJSON)
	if err != nil {
		return "", err
	}

	return previousCurrentPosition, nil
}

//================================================================================================================================================
// UPDATING THE RECORDS FOR A CUSTOM TRAINING
//================================================================================================================================================

// UpdatedCustomTrainingEmployee updates the employee Custom Training field of the "digital asset" with the given id in the world state
// and returns the previous
// Updated the employee Custom Training records (team member - TM); When employee is completing the Competency Module Custom Training his/her ledger information's will be updated accordingly.

func (s *EmployeeSmartContract) UpdatedCustomTrainingEmployee(ctx contractapi.TransactionContextInterface, id string, newCustomTraining string) (string, error) {

	asset, err := s.ReadEmployee(ctx, id)
	if err != nil {
		return "", err
	}

	previousCustomTraining := asset.CustomTraining
	asset.CustomTraining = newCustomTraining

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(id, assetJSON)
	if err != nil {
		return "", err
	}

	return previousCustomTraining, nil
}
