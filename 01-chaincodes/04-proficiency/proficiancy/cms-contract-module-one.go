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
	"errors"
	"fmt"

	"log"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	// "github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	// "github.com/hyperledger/fabric-contract-api-go/internal/functionaltests/contracts/utils"
)

// CompetencySmartContractModuleOne contract for handling BasicAssets
type CompetencySmartContractModuleOne struct {
	contractapi.Contract
}

// CustomTransactionContextInterface adds extra methods to basic context
// interface to give access to callData
type CustomTransactionContextInterface interface {
	contractapi.TransactionContextInterface

	SetCallData([]byte)
	GetCallData() []byte
}

// CustomTransactionContext adds extra field to contractapi.TransactionContext
// so that data can be between calls
type CustomTransactionContext struct {
	contractapi.TransactionContext
	callData []byte
}

// SetCallData sets the call data property
func (ctx *CustomTransactionContext) SetCallData(bytes []byte) {
	ctx.callData = bytes
}

// GetCallData gets the call data property
func (ctx *CustomTransactionContext) GetCallData() []byte {
	return ctx.callData
}

// GetWorldState takes a key and sets what is found in the world state for that
// key in the transaction context
func GetWorldState(ctx CustomTransactionContextInterface) error {
	_, params := ctx.GetStub().GetFunctionAndParameters()

	if len(params) < 1 {
		return errors.New("missing key for world state")
	}

	existing, err := ctx.GetStub().GetState(params[0])

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	ctx.SetCallData(existing)

	return nil
}

// UnknownTransactionHandler logs details of a bad transaction request
// and returns a shim error
func UnknownTransactionHandler(ctx CustomTransactionContextInterface) error {
	fcn, args := ctx.GetStub().GetFunctionAndParameters()
	return fmt.Errorf("invalid function %s passed with args [%s]", fcn, strings.Join(args, ", "))
}

type HistoryQueryResultModuleOne struct {
	Record    *ModuleOne `json:"record"`
	TxId      string     `json:"txId"`
	Timestamp time.Time  `json:"timestamp"`
	IsDelete  bool       `json:"isDelete"`
}

//==================================================================================================================
// CREATE NEW TRAINING MODULE ONE - MODULE 1
//==================================================================================================================

// CreateNewModuleOne adds a new basic asset to the world state using id as key
func (s *CompetencySmartContractModuleOne) CreateNewModuleOne(ctx CustomTransactionContextInterface, module_one_id string, trainer_module_one TrainerModuleOne, trainee_module_one TraineeModuleOne, module_one_chapter_one ModuleOneChapterOne, module_one_chapter_two ModuleOneChapterTwo, module_one_chapter_three ModuleOneChapterThree) error {

	// Add access control!!!
	/*
		mspId, err := cid.GetMSPID(ctx.GetStub())
		if err != nil {
			return fmt.Errorf("failed while getting identity. %s", err.Error())
		}

		if mspId != "org-operations-com" {
			return fmt.Errorf("you are not authorized to create a new training Module 1 contexts")
		}
	*/

	existing := ctx.GetCallData()

	if existing != nil {
		return fmt.Errorf("cannot create new Module 1 training contexts in world state as key %s already exists", module_one_id)
	}

	ba := new(ModuleOne)
	ba.ModuleOneID = module_one_id
	ba.ModuleOneTrainer = trainer_module_one
	ba.ModuleOneTrainee = trainee_module_one
	ba.ChapterOneModuleOne = module_one_chapter_one
	ba.ChapterTwoModuleOne = module_one_chapter_two
	ba.ChapterThreeModuleOne = module_one_chapter_three

	// These functions will set up the default values.
	ba.SetModuleOneCreatedDate()
	ba.SetStatusTrainingTypeInitial()
	ba.SetStatusChapterOneModuleOneOpen()
	ba.SetStatusChapterTwoModuleOneOpen()
	ba.SetStatusChapterThreeModuleOneOpen()
	ba.SetStatusModuleOneOpen()

	// Set the assessments, both theoretical and practical, and total. At this stage, all values shall be "0".
	ba.SetTheoreticalAssessmentModuleOne()
	ba.SetPracticalAssessmentModuleOne()
	ba.SetTotalAssessmentModuleOne()

	// Set the AssessmentsAttempts to "0"
	ba.SetAssessmentsAttemptsToZero()

	// Set the transaction ID

	baBytes, _ := json.Marshal(ba)

	err := ctx.GetStub().PutState(module_one_id, []byte(baBytes))

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	return nil
}

//==================================================================================================================
// UPDATE THE STATUS OF THE CHAPTER ONE OF TRAINING MODULE ONE - MODUEL 1
//==================================================================================================================

// UpdateStatusChapterOneModuleOneToCompleted changes the ownership of a basic asset and mark it as used

func (cc *CompetencySmartContractModuleOne) UpdateStatusChapterOneModuleOneToCompleted(ctx CustomTransactionContextInterface, module_one_id string) error {

	// Add access control!!!
	/*
		mspId, err := cid.GetMSPID(ctx.GetStub())
		if err != nil {
			return fmt.Errorf("failed while getting identity. %s", err.Error())
		}

		if mspId != "org-operations-com" {
			return fmt.Errorf("you are not authorized to create a new training Module 1 contexts")
		}
	*/

	existing := ctx.GetCallData()

	if existing == nil {
		return fmt.Errorf("cannot update asset in world state as key %s does not exist", module_one_id)
	}

	ba := new(ModuleOne)

	err := json.Unmarshal(existing, ba)

	if err != nil {
		return fmt.Errorf("data retrieved from world state for key %s was not of type ModuleOne", module_one_id)
	}

	// ba.ChapterOneModuleOne = newChapterOneModuleOne
	ba.SetStatusChapterOneModuleOneCompleted()

	baBytes, _ := json.Marshal(ba)

	err = ctx.GetStub().PutState(module_one_id, []byte(baBytes))

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	return nil
}

//==================================================================================================================
// UPDATE THE STATUS OF THE CHAPTER TWO OF TRAINING MODULE ONE - Module 1
//==================================================================================================================

// UpdateStatusChapterTwoModuleOneToCompleted changes the ownership of a basic asset and mark it as used

func (cc *CompetencySmartContractModuleOne) UpdateStatusChapterTwoModuleOneToCompleted(ctx CustomTransactionContextInterface, module_one_id string) error {

	// Add access control!!!
	/*
		mspId, err := cid.GetMSPID(ctx.GetStub())
		if err != nil {
			return fmt.Errorf("failed while getting identity. %s", err.Error())
		}

		if mspId != "org-operations-com" {
			return fmt.Errorf("you are not authorized to create a new training Module 1 contexts")
		}
	*/

	existing := ctx.GetCallData()

	if existing == nil {
		return fmt.Errorf("cannot update asset in world state as key %s does not exist", module_one_id)
	}

	ba := new(ModuleOne)

	err := json.Unmarshal(existing, ba)

	if err != nil {
		return fmt.Errorf("data retrieved from world state for key %s was not of type ModuleOne", module_one_id)
	}

	// ba.ChapterOneModuleOne = newChapterOneModuleOne
	ba.SetStatusChapterTwoModuleOneCompleted()

	baBytes, _ := json.Marshal(ba)

	err = ctx.GetStub().PutState(module_one_id, []byte(baBytes))

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	return nil
}

//==================================================================================================================
// UPDATE THE STATUS OF THE CHAPTER THREE OF TRAINING MODULE ONE
//==================================================================================================================

// UpdateStatusChapterThreeModuleOneToCompleted changes the ownership of a basic asset and mark it as used

func (cc *CompetencySmartContractModuleOne) UpdateStatusChapterThreeModuleOneToCompleted(ctx CustomTransactionContextInterface, module_one_id string) error {

	// Add access control!!!
	/*
		mspId, err := cid.GetMSPID(ctx.GetStub())
		if err != nil {
			return fmt.Errorf("failed while getting identity. %s", err.Error())
		}

		if mspId != "org-operations-com" {
			return fmt.Errorf("you are not authorized to create a new training Module 1 contexts")
		}
	*/

	existing := ctx.GetCallData()

	if existing == nil {
		return fmt.Errorf("cannot update asset in world state as key %s does not exist", module_one_id)
	}

	ba := new(ModuleOne)

	err := json.Unmarshal(existing, ba)

	if err != nil {
		return fmt.Errorf("data retrieved from world state for key %s was not of type ModuleOne", module_one_id)
	}

	// ba.ChapterOneModuleOne = newChapterOneModuleOne
	ba.SetStatusChapterThreeModuleOneCompleted()

	baBytes, _ := json.Marshal(ba)

	err = ctx.GetStub().PutState(module_one_id, []byte(baBytes))

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	return nil
}

//==================================================================================================================
// UPDATING THE RESULTS FOR THE THEORETICAL ASSESSMENT OF THE TRAINING MODULE ONE
//==================================================================================================================

// UpdateValueTheoreticalAssessment changes the value of a basic asset to add the value passed
func (cc *CompetencySmartContractModuleOne) UpdateValueTheoreticalAssessment(ctx CustomTransactionContextInterface, module_one_id string, valueAdd uint) error {

	existing := ctx.GetCallData()

	if existing == nil {
		return fmt.Errorf("cannot update asset in world state as key %s does not exist", module_one_id)
	}

	// to be added check if the all modules are completed or not. Solution would be with the applications, which will check if chapters are completed.

	ba := new(ModuleOne)

	err := json.Unmarshal(existing, ba)

	if err != nil {
		return fmt.Errorf("data retrieved from world state for key %s was not of type BasicAsset", module_one_id)
	}

	ba.TheoreticalAssessmentModuleOne += valueAdd

	// this part of the code must be put "above" to delete whenever we want to do a new assessment.
	if ba.TheoreticalAssessmentModuleOne < 60 {
		ba.SetTheoreticalAssessmentModuleOne()
		// fmt.Println("theoretical assessment field, you have less than 60 points:", ba.TheoreticalAssessmentModuleOne)
	}

	baBytes, _ := json.Marshal(ba)

	err = ctx.GetStub().PutState(module_one_id, []byte(baBytes))

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	return nil
}

//==================================================================================================================
// UPDATING THE RESULTS FOR THE PRACTICAL ASSESSMENT OF THE TRAINING MODULE ONE
//==================================================================================================================

// UpdateValuePracticalAndTotalAssessment changes the value of a basic asset to add the value passed
func (cc *CompetencySmartContractModuleOne) UpdateValuePracticalAndTotalAssessment(ctx CustomTransactionContextInterface, module_one_id string, valueAdd uint) error {
	existing := ctx.GetCallData()

	if existing == nil {
		return fmt.Errorf("cannot update asset in world state as key %s does not exist", module_one_id)
	}

	// to be added check if the all modules are completed or not. Solution would be with the applications, which will check if chapters are completed.

	ba := new(ModuleOne)

	err := json.Unmarshal(existing, ba)

	if err != nil {
		return fmt.Errorf("data retrieved from world state for key %s was not of type BasicAsset", module_one_id)
	}

	// Checking how many attempts you already consumed.

	if ba.AssessmentsAttempts == 3 {
		return errors.New("there is no more assessment attempts left, fucking stupid moron")
	}

	ba.PracticalAssessmentModuleOne += valueAdd

	if ba.PracticalAssessmentModuleOne < 60 {
		ba.SetPracticalAssessmentModuleOne()
		// fmt.Println("theoretical assessment field, you have less than 60 points:", ba.TheoreticalAssessmentModuleOne)
	}

	// Here must be stopped that the result form the practical assessment is added to the theoretical assessments!!!

	// adding total value for the both assessments.

	ba.AssessmentModuleOne = ba.PracticalAssessmentModuleOne + ba.TheoreticalAssessmentModuleOne

	// changing the status of the module one completed if total value of points if equal or greater then 120 points (60+60).

	if ba.AssessmentModuleOne < 120 {
		fmt.Println("total number of points is less than 120, please re-do your practical assessment in 15 days:", ba.AssessmentModuleOne)
		// return fmt.Errorf("practical assessment failed %s please redo it in 15 days", module_one_id)
	} else {
		ba.SetStatusModuleOneOpen()
		// fmt.Println("Successfully completed Module One:", ba.AssessmentModuleOne)
		// ba.SetStatusModuleOneCompleted()
	}

	if ba.AssessmentModuleOne > 120 {
		ba.SetStatusModuleOneCompleted()
		// fmt.Println("Successfully completed Module One:", ba.AssessmentModuleOne)
		// ba.SetStatusModuleOneCompleted()
	}

	ba.AssessmentsAttempts += 1

	/*
		if ba.AssessmentsAttempts == 1 {
			fmt.Println("you have two remaining attempts attempt")
		} else {
			ba.SetAssessmentsAttemptsToOne()
		}
	*/

	baBytes, _ := json.Marshal(ba)

	err = ctx.GetStub().PutState(module_one_id, []byte(baBytes))

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	return nil
}

//==================================================================================================================
// READ THE RECORDS OF THE TRAINING MODULE ONE
//==================================================================================================================

// GetModuleRecords returns the basic asset with id given from the world state
func (cc *CompetencySmartContractModuleOne) GetModuleOneRecords(ctx CustomTransactionContextInterface, module_one_id string) (*ModuleOne, error) {

	existing := ctx.GetCallData()

	if existing == nil {
		return nil, fmt.Errorf("cannot read world state pair with key %s. Does not exist", module_one_id)
	}

	ba := new(ModuleOne)

	err := json.Unmarshal(existing, ba)

	// err := json.Unmarshal(existing, ba) => this line was deleted!!!

	if err != nil {
		return nil, fmt.Errorf("data retrieved from world state for key %s was not of type BasicAsset", module_one_id)
	}

	return ba, nil
}

//==================================================================================================================
// GET HISTORY OF RECORDS OF THE TRAINING MODULE ONE
//==================================================================================================================

// GetHistoryRecordsForModuleOne returns the chain of custody for an asset since issuance.
func (cc *CompetencySmartContractModuleOne) GetHistoryRecordsForModuleOne(ctx contractapi.TransactionContextInterface, module_one_id string) ([]HistoryQueryResultModuleOne, error) {
	log.Printf("GetHistoryRecordsForModuleOne: ID %v", module_one_id)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(module_one_id)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []HistoryQueryResultModuleOne
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset ModuleOne
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &asset)
			if err != nil {
				return nil, err
			}
		} else {
			asset = ModuleOne{
				ModuleOneID: module_one_id,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			return nil, err
		}

		record := HistoryQueryResultModuleOne{
			TxId:      response.TxId,
			Timestamp: timestamp,
			Record:    &asset,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}

//==================================================================================================================
// SET THEORETICAL AND PRACTICAL ASSESSMENTS TO "0" AT TRAINING MODULE ONE
//==================================================================================================================

// SetTheoreticalAndPracticalAssessmentsToZero changes the ownership of a basic asset and mark it as used

func (cc *CompetencySmartContractModuleOne) SetTheoreticalAndPracticalAssessmentsToZero(ctx CustomTransactionContextInterface, module_one_id string) error {

	// Add access control!!!
	/*
		mspId, err := cid.GetMSPID(ctx.GetStub())
		if err != nil {
			return fmt.Errorf("failed while getting identity. %s", err.Error())
		}

		if mspId != "org-operations-com" {
			return fmt.Errorf("you are not authorized to create a new training Module 1 contexts")
		}
	*/

	existing := ctx.GetCallData()

	if existing == nil {
		return fmt.Errorf("cannot update asset in world state as key %s does not exist", module_one_id)
	}

	ba := new(ModuleOne)

	err := json.Unmarshal(existing, ba)

	if err != nil {
		return fmt.Errorf("data retrieved from world state for key %s was not of type ModuleOne", module_one_id)
	}

	// ba.ChapterOneModuleOne = newChapterOneModuleOne
	// ba.SetStatusChapterThreeModuleOneCompleted()
	ba.SetTheoreticalAssessmentModuleOne()
	ba.SetPracticalAssessmentModuleOne()

	ba.AssessmentsAttempts += 1

	// here must be created a function which will stop updating as the attempts shall be limited to 3

	baBytes, _ := json.Marshal(ba)

	err = ctx.GetStub().PutState(module_one_id, []byte(baBytes))

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	return nil
}

//==================================================================================================================
// QUERY THE TRAINING MODULE ONE
//==================================================================================================================

// constructQueryResponseFromIteratorModuleOne constructs a slice of assets from the resultsIterator
func constructQueryResponseFromIteratorModuleOne(resultsIteratorQuery shim.StateQueryIteratorInterface) ([]*ModuleOne, error) {
	var assets []*ModuleOne
	for resultsIteratorQuery.HasNext() {
		queryResult, err := resultsIteratorQuery.Next()
		if err != nil {
			return nil, err
		}
		var asset ModuleOne
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

// QueryModuleOne uses a query string to perform a query for the Module 1 records.
// Query string matching state database syntax is passed in and executed as is.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the QueryModuleOne example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Ad hoc rich query
func (cc *CompetencySmartContractModuleOne) QueryModuleOne(ctx contractapi.TransactionContextInterface, queryString string) ([]*ModuleOne, error) {
	return getQueryResultForQueryStringModuleOne(ctx, queryString)
}

// getQueryResultForQueryStringModuleOne executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryStringModuleOne(ctx contractapi.TransactionContextInterface, queryString string) ([]*ModuleOne, error) {
	resultsIteratorQuery, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorQuery.Close()

	return constructQueryResponseFromIteratorModuleOne(resultsIteratorQuery)
}
