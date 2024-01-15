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
	//"crypto/rand"
	//"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const indexDailyLog = "ID~Type"

// "DailyOperationsLogSmartContract" implements the fabric-contract-api-go programming model.
// Also, it is important to underline, that the MSP for this type of "asset(s)"
// is "ops-operations-com", "org-operations-com" and "mte-operations-com".
// That said, the access control is allowed only through "ops-operations-com", "org-operations-com" & "mte-operations-com" MSPs'.
type DailyOperationsLogSmartContract struct {
	contractapi.Contract
}

// SetValidation set the condition of the object to mark as new
func (ba *DailyOperationsLog) SetValidatated() {
	ba.DailyOpsLogValidation = "validated"
}

// SetNotValidated set the condition of the object to mark as used
func (ba *DailyOperationsLog) SetNotValidated() {
	ba.DailyOpsLogValidation = "not-validated"
}

// HistoryQueryResult structure used for returning result of history query
type HistoryQueryResultDailyOpsLog struct {
	Record    *DailyOperationsLog `json:"recordops"`
	TxId      string              `json:"txId"`
	Timestamp time.Time           `json:"timestamp"`
	IsDelete  bool                `json:"isDelete"`
}

// PaginatedQueryResult structure used for returning paginated query results and metadata
type PaginatedQueryResultDailyOpsLog struct {
	Records             []*DailyOperationsLog `json:"recordsops"`
	FetchedRecordsCount int32                 `json:"fetchedRecordsCount"`
	Bookmark            string                `json:"bookmark"`
}

//================================================================================================================================================
// CREATE NEW DAILY OPERATIONS LOG
//================================================================================================================================================

// CreateNewDailyOperationsLog creates a new work order on the ledger
func (t *DailyOperationsLogSmartContract) CreateNewDailyOperationsLog(ctx contractapi.TransactionContextInterface, daily_ops_log_id string, owner string, weather_drive WeatherDriveStation, weather_return WeatherReturnStation, personnel_on_duty_drive_station PersonnelOnDutyDriveStation, personnel_on_duty_return_station PersonnelOnDutyReturnStation, operations_start string, operations_end string, operation_hours int, number_of_carriers int, number_of_passengers int, total_operating_hours int, additional_comments string) error {

	exists, err := t.DoesDailyOperationsLogExists(ctx, daily_ops_log_id)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	}
	if exists {
		return fmt.Errorf("asset already exists you fucking stupid donkey: %s", daily_ops_log_id)
	}

	dailyopslog := &DailyOperationsLog{
		DocType:               "dailyopslog",
		DailyOpsLogID:         daily_ops_log_id,
		Owner:                 owner,
		WeatherDrive:          weather_drive,
		WeatherReturn:         weather_return,
		PersonnelOnDutyDrive:  personnel_on_duty_drive_station,
		PersonnelOnDutyReturn: personnel_on_duty_return_station,
		OperationsStart:       operations_start,
		OperationsEnd:         operations_end,
		OperationHours:        operation_hours,
		NumberOfCarriers:      number_of_carriers,
		NumberOfPassengers:    number_of_passengers,
		TotalOperatingHours:   total_operating_hours,
		AdditionalComments:    additional_comments,
	}

	// Setting the validation record into "no-validated"
	dailyopslog.SetNotValidated()

	assetBytes, err := json.Marshal(dailyopslog)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(daily_ops_log_id, assetBytes)
	if err != nil {
		return err
	}

	//  Create an index to enable daily operations log ID-type based range queries, e.g. return all assets with specific ID.
	//  An 'index' is a normal key-value entry in the ledger.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on indexName~PrevWorkOrderID~WorkOrderType.
	//  This will enable very efficient state range queries based on composite keys matching indexName~PrevWorkOrderID~*
	dailyOperationsLogIndexKey, err := ctx.GetStub().CreateCompositeKey(indexDailyLog, []string{dailyopslog.DailyOpsLogID, dailyopslog.Owner})
	if err != nil {
		return err
	}
	//  Save index entry to world state. Only the key name is needed, no need to store a duplicate copy of the asset.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	return ctx.GetStub().PutState(dailyOperationsLogIndexKey, value)
}

//================================================================================================================================================
// READ DAILY OPERATIONS LOG
//================================================================================================================================================

// ReadDailyOperationsLog retrieves an asset from the ledger.
func (t *DailyOperationsLogSmartContract) ReadDailyOperationsLog(ctx contractapi.TransactionContextInterface, daily_ops_log_id string) (*DailyOperationsLog, error) {

	assetBytes, err := ctx.GetStub().GetState(daily_ops_log_id)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset %s: %v", daily_ops_log_id, err)
	}
	if assetBytes == nil {
		return nil, fmt.Errorf("asset %s does not exist moron", daily_ops_log_id)
	}

	var dailyopslog DailyOperationsLog
	err = json.Unmarshal(assetBytes, &dailyopslog)
	if err != nil {
		return nil, err
	}

	return &dailyopslog, nil
}

//================================================================================================================================================
// DELETE DAILY OPERATIONS LOG
//================================================================================================================================================

// DeleteDailyOperationsLog removes an asset key-value pair from the ledger

func (t *DailyOperationsLogSmartContract) DeleteDailyOperationsLog(ctx contractapi.TransactionContextInterface, daily_ops_log_id string) error {

	dailyopslog, err := t.ReadDailyOperationsLog(ctx, daily_ops_log_id)
	if err != nil {
		return err
	}

	err = ctx.GetStub().DelState(daily_ops_log_id)

	if err != nil {
		return fmt.Errorf("failed to delete asset %s: %v", daily_ops_log_id, err)
	}

	dailyOperationsLogIndexKey, err := ctx.GetStub().CreateCompositeKey(indexDailyLog, []string{dailyopslog.DailyOpsLogID, dailyopslog.Owner})
	if err != nil {
		return err
	}

	// Delete index entry
	return ctx.GetStub().DelState(dailyOperationsLogIndexKey)
}

//================================================================================================================================================
// CONSTRUCT QUERY RESPONSE FROM ITERATOR
//================================================================================================================================================

// constructQueryResponseFromIteratorDailyOpsLog constructs a slice of assets from the resultsIteratorDailyOpsLog
func constructQueryResponseFromIteratorDailyOpsLog(resultsIteratorDailyOpsLog shim.StateQueryIteratorInterface) ([]*DailyOperationsLog, error) {
	var assetsops []*DailyOperationsLog
	for resultsIteratorDailyOpsLog.HasNext() {
		queryResult, err := resultsIteratorDailyOpsLog.Next()
		if err != nil {
			return nil, err
		}
		var dailyopslog DailyOperationsLog
		err = json.Unmarshal(queryResult.Value, &dailyopslog)
		if err != nil {
			return nil, err
		}
		assetsops = append(assetsops, &dailyopslog)
	}

	return assetsops, nil
}

//================================================================================================================================================
// GET DAILY OPERATIONS LOGS BY RANGE
//================================================================================================================================================

// GetDailyOperationsLogByRange performs a range query based on the start and end keys provided.
// Read-only function results are not typically submitted to ordering. If the read-only
// results are submitted to ordering, or if the query is used in an update transaction
// and submitted to ordering, then the committing peers will re-execute to guarantee that
// result sets are stable between endorsement time and commit time. The transaction is
// invalidated by the committing peers if the result set has changed between endorsement
// time and commit time.
// Therefore, range queries are a safe option for performing update transactions based on query results.
func (t *DailyOperationsLogSmartContract) GetDailyOperationsLogByRange(ctx contractapi.TransactionContextInterface, startKey, endKey string) ([]*DailyOperationsLog, error) {

	resultsIteratorDailyOpsLog, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorDailyOpsLog.Close()

	return constructQueryResponseFromIteratorDailyOpsLog(resultsIteratorDailyOpsLog)
}

/*
//================================================================================================================================================
// QUERY DAILY OPERATIONS LOG BY OWNER - NO NEED AS WE QUERY THE LEDGER BY TEH QUERY STRING!!!
//================================================================================================================================================

// QueryDailyOperationsLogByType queries for assets based on the types name.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (type).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *DailyOperationsLogSmartContract) QueryDailyOperationsLogByOwner(ctx contractapi.TransactionContextInterface, owner string) ([]*DailyOperationsLog, error) {
	queryStringOwner := fmt.Sprintf(`{"selector":{"docType":"dailyopslog","owner":"%s"}}`, owner)
	return getQueryResultForQueryStringDailyOpsLog(ctx, queryStringOwner)
}
*/

/*
//================================================================================================================================================
// QUERY DAILY OPERATIONS LOGS BY CONDITION
//================================================================================================================================================

// QueryDailyOperationsLogs queries for assets based on their condition.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (condition).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *DailyOperationsLogSmartContract) QueryDailyOperationsLogByCondition(ctx contractapi.TransactionContextInterface, dailyopslog_condition string) ([]*DailyOperationsLog, error) {
	queryStringCond := fmt.Sprintf(`{"selector":{"docType":"dailyopslog","dailyopslog_condition":"%s"}}`, dailyopslog_condition)
	return getQueryResultForQueryStringDailyOpsLog(ctx, queryStringCond)
}
*/

/*
//================================================================================================================================================
// QUERY DAILY OPERATIONS LOGS BY VALIDATION - NO NEED AS WE QUERY THE LEDGER BY TEH QUERY STRING!!!
//================================================================================================================================================

// QueryDailyOperationsLogByValidation queries for work orders based on their validation.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (validation).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *DailyOperationsLogSmartContract) QueryDailyOperationsLogByValidation(ctx contractapi.TransactionContextInterface, dailyopslog_validation string) ([]*DailyOperationsLog, error) {
	queryStringVal := fmt.Sprintf(`{"selector":{"docType":"dailyopslog","dailyopslog_validation":"%s"}}`, dailyopslog_validation)
	return getQueryResultForQueryStringDailyOpsLog(ctx, queryStringVal)
}
*/

/*
//================================================================================================================================================
// QUERY DAILY OPERATIONS LOGS BY ID - NO NEED AS WE QUERY THE LEDGER BY TEH QUERY STRING!!!
//================================================================================================================================================

// QueryDailyOperationsLogByID queries for work orders based on the its ID number.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (type).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *DailyOperationsLogSmartContract) QueryDailyOperationsLogByID(ctx contractapi.TransactionContextInterface, daily_ops_log_id string) ([]*DailyOperationsLog, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"dailyopslog", daily_ops_log_id":"%s"}}`, daily_ops_log_id)
	return getQueryResultForQueryStringDailyOpsLog(ctx, queryString)
}
*/

//================================================================================================================================================
// QUERY DAILY OPERATIONS LOG!!!
//================================================================================================================================================

// QueryDailyOperationsLog uses a query string to perform a query for work orders.
// Query string matching state database syntax is passed in and executed as is.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the "QueryAssetsForOwner" example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Ad hoc rich query
func (t *DailyOperationsLogSmartContract) QueryDailyOperationsLog(ctx contractapi.TransactionContextInterface, queryString string) ([]*DailyOperationsLog, error) {

	return getQueryResultForQueryStringDailyOpsLog(ctx, queryString)
}

//================================================================================================================================================
// GET QUERY RESULT FOR QUERY STRING
//================================================================================================================================================

// getQueryResultForQueryStringDailyOpsLog executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryStringDailyOpsLog(ctx contractapi.TransactionContextInterface, queryString string) ([]*DailyOperationsLog, error) {
	resultsIteratorDailyOpsLog, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorDailyOpsLog.Close()

	return constructQueryResponseFromIteratorDailyOpsLog(resultsIteratorDailyOpsLog)
}

//================================================================================================================================================
// GET DAILY OPERATIONS LOGS BY RANGE WITH PAGINATION
//================================================================================================================================================

// GetDailyOperationsLogByRangeWithPagination performs a range query based on the start and end key,
// page size and a bookmark.
// The number of fetched records will be equal to or lesser than the page size.
// Paginated range queries are only valid for read only transactions.
// Example: Pagination with Range Query
func (t *DailyOperationsLogSmartContract) GetDailyOperationsLogByRangeWithPagination(ctx contractapi.TransactionContextInterface, startKey string, endKey string, pageSize int, bookmark string) (*PaginatedQueryResultDailyOpsLog, error) {

	resultsIteratorDailyOpsLog, responseMetadata, err := ctx.GetStub().GetStateByRangeWithPagination(startKey, endKey, int32(pageSize), bookmark)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorDailyOpsLog.Close()

	assetsops, err := constructQueryResponseFromIteratorDailyOpsLog(resultsIteratorDailyOpsLog)
	if err != nil {
		return nil, err
	}

	return &PaginatedQueryResultDailyOpsLog{
		Records:             assetsops,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil
}

//================================================================================================================================================
// QUERY DAILY OPERATIONS LOGS WITH PAGINATION
//================================================================================================================================================

// QueryDailyOperationsLogWithPagination uses a query string, page size and a bookmark to perform a query
// for assets. Query string matching state database syntax is passed in and executed as is.
// The number of fetched records would be equal to or lesser than the specified page size.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the QueryAssetsForOwner example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// Paginated queries are only valid for read only transactions.
// Example: Pagination with Ad hoc Rich Query
func (t *DailyOperationsLogSmartContract) QueryDailyOperationsLogWithPagination(ctx contractapi.TransactionContextInterface, queryString string, pageSize int, bookmark string) (*PaginatedQueryResultDailyOpsLog, error) {

	return getQueryResultForQueryStringDailyOpsLogWithPagination(ctx, queryString, int32(pageSize), bookmark)
}

//================================================================================================================================================
// GET QUERY RESULT FOR QUERY STRING WITH PAGINATION
//================================================================================================================================================

// getQueryResultForQueryStringDailyOpsLogWithPagination executes the passed in query string with
// pagination info. The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryStringDailyOpsLogWithPagination(ctx contractapi.TransactionContextInterface, queryString string, pageSize int32, bookmark string) (*PaginatedQueryResultDailyOpsLog, error) {

	resultsIteratorDailyOpsLog, responseMetadata, err := ctx.GetStub().GetQueryResultWithPagination(queryString, pageSize, bookmark)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorDailyOpsLog.Close()

	assetsops, err := constructQueryResponseFromIteratorDailyOpsLog(resultsIteratorDailyOpsLog)
	if err != nil {
		return nil, err
	}

	return &PaginatedQueryResultDailyOpsLog{
		Records:             assetsops,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil
}

//================================================================================================================================================
// GET DAILY OPERATIONS LOG HISTORY
//================================================================================================================================================

// GetDailyOperationsLogHistory returns the chain of custody for an asset since issuance.
func (t *DailyOperationsLogSmartContract) GetDailyOperationsLogHistory(ctx contractapi.TransactionContextInterface, daily_ops_log_id string) ([]HistoryQueryResultDailyOpsLog, error) {

	log.Printf("GetAssetHistory: ID %v", daily_ops_log_id)

	resultsIteratorDailyOpsLog, err := ctx.GetStub().GetHistoryForKey(daily_ops_log_id)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorDailyOpsLog.Close()

	var recordsops []HistoryQueryResultDailyOpsLog
	for resultsIteratorDailyOpsLog.HasNext() {
		response, err := resultsIteratorDailyOpsLog.Next()
		if err != nil {
			return nil, err
		}

		var dailyopslog DailyOperationsLog
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &dailyopslog)
			if err != nil {
				return nil, err
			}
		} else {
			dailyopslog = DailyOperationsLog{
				DailyOpsLogID: daily_ops_log_id,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			return nil, err
		}

		recordops := HistoryQueryResultDailyOpsLog{
			TxId:      response.TxId,
			Timestamp: timestamp,
			Record:    &dailyopslog,
			IsDelete:  response.IsDelete,
		}
		recordsops = append(recordsops, recordops)
	}

	return recordsops, nil
}

//================================================================================================================================================
// DOES DAILY OPERATIONS LOG EXISTS/*
//================================================================================================================================================

// DoesDailyOperationsLogExists returns true when asset with given ID exists in the ledger.
func (t *DailyOperationsLogSmartContract) DoesDailyOperationsLogExists(ctx contractapi.TransactionContextInterface, daily_ops_log_id string) (bool, error) {

	assetBytes, err := ctx.GetStub().GetState(daily_ops_log_id)
	if err != nil {
		return false, fmt.Errorf("failed to read asset %s from world state. %v", daily_ops_log_id, err)
	}

	return assetBytes != nil, nil
}

//================================================================================================================================================
// UPDATE DAILY OPERATIONS LOG
//================================================================================================================================================

// UpdateDailyOperationsLog updates an existing work order in the world state with provided parameters.
func (t *DailyOperationsLogSmartContract) UpdateDailyOperationsLog(ctx contractapi.TransactionContextInterface, daily_ops_log_id string, owner string, weather_drive WeatherDriveStation, weather_return WeatherReturnStation, personnel_on_duty_drive_station PersonnelOnDutyDriveStation, personnel_on_duty_return_station PersonnelOnDutyReturnStation, operations_start string, operations_end string, operation_hours int, number_of_carriers int, number_of_passengers int, total_operating_hours int, additional_comments string) error {

	exists, err := t.DoesDailyOperationsLogExists(ctx, daily_ops_log_id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", daily_ops_log_id)
	}

	// overwriting original asset with new asset
	dailyopslog := &DailyOperationsLog{
		DocType:               "dailyopslog",
		DailyOpsLogID:         daily_ops_log_id,
		Owner:                 owner,
		WeatherDrive:          weather_drive,
		WeatherReturn:         weather_return,
		PersonnelOnDutyDrive:  personnel_on_duty_drive_station,
		PersonnelOnDutyReturn: personnel_on_duty_return_station,
		OperationsStart:       operations_start,
		OperationsEnd:         operations_end,
		OperationHours:        operation_hours,
		NumberOfCarriers:      number_of_carriers,
		NumberOfPassengers:    number_of_passengers,
		TotalOperatingHours:   total_operating_hours,
		AdditionalComments:    additional_comments,
	}

	dailyopslog.SetValidatated()

	// dailyopslog.SetDowntimeValidated()

	assetJSON, err := json.Marshal(dailyopslog)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(daily_ops_log_id, assetJSON)
}

//================================================================================================================================================
// OPEN TO UPDATE DAILY OPERATIONS LOG
//================================================================================================================================================

// UpdateDailyOperationsLog updates an existing work order in the world state with provided parameters.
func (t *DailyOperationsLogSmartContract) OpenToUpdatedDailyOperationsLog(ctx contractapi.TransactionContextInterface, daily_ops_log_id string, owner string, weather_drive WeatherDriveStation, weather_return WeatherReturnStation, personnel_on_duty_drive_station PersonnelOnDutyDriveStation, personnel_on_duty_return_station PersonnelOnDutyReturnStation, operations_start string, operations_end string, operation_hours int, additional_comments string) error {

	exists, err := t.DoesDailyOperationsLogExists(ctx, daily_ops_log_id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", daily_ops_log_id)
	}

	// overwriting original asset with new asset
	dailyopslog := &DailyOperationsLog{
		DocType:               "dailyopslog",
		DailyOpsLogID:         daily_ops_log_id,
		Owner:                 owner,
		WeatherDrive:          weather_drive,
		WeatherReturn:         weather_return,
		PersonnelOnDutyDrive:  personnel_on_duty_drive_station,
		PersonnelOnDutyReturn: personnel_on_duty_return_station,
		OperationsStart:       operations_start,
		OperationsEnd:         operations_end,
		OperationHours:        operation_hours,
		AdditionalComments:    additional_comments,
	}

	dailyopslog.SetNotValidated()

	assetJSON, err := json.Marshal(dailyopslog)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(daily_ops_log_id, assetJSON)
}

//================================================================================================================================================
// GET ALL DAILY OPERATIONS LOGS
//================================================================================================================================================

// GetAllDailyOperationsLogs returns all assets found in world state
func (t *DailyOperationsLogSmartContract) GetAllDailyOperationsLogs(ctx contractapi.TransactionContextInterface) ([]*DailyOperationsLog, error) {

	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.

	resultsIteratorDailyOpsLog, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIteratorDailyOpsLog.Close()

	var assetsops []*DailyOperationsLog
	for resultsIteratorDailyOpsLog.HasNext() {
		queryResponse, err := resultsIteratorDailyOpsLog.Next()
		if err != nil {
			return nil, err
		}

		var asset DailyOperationsLog
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assetsops = append(assetsops, &asset)
	}

	return assetsops, nil
}

//================================================================================================================================================
// CALL FUNCTION FROM OTHER CHAINCODE (TEST)
//================================================================================================================================================

// GetANewCorrectiveMaintenanceWorkOrder call for a function form the other smart contract,
// in this case calling the smart contract for corrective maintenance.
// "employee" chaincode is from my working folder, chaincode for employees in minifabric (tested).
//
// IMPORTANT:
// "functionName" => function that needs to be called, when tested I called "ReadAsset"
// "documentID"   => this is the parameter from the function we use, in this case we use ID number e.g. "MID-112233".

func (t *DailyOperationsLogSmartContract) CallingCrossChaincodeFunctionEmployee(ctx contractapi.TransactionContextInterface, functionName string, documentDataEmployeeID string) (string, error) {

	if len(documentDataEmployeeID) == 0 {
		return "", fmt.Errorf("please provide correct document data for")
	}

	params := []string{"EmployeeSmartContract:ReadEmployee", documentDataEmployeeID}

	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := ctx.GetStub().InvokeChaincode("orgemployee", queryArgs, "default-channel")

	return string(response.Payload), nil
}

//================================================================================================================================================
// CALL FUNCTION FROM OTHER CHAINCODE (TEST)
//================================================================================================================================================

// CallingCrossChaincodeFunctionCorrectiveMaintenance call for a function form the other smart contract,
// in this case calling the smart contract for corrective maintenance.
// "employee" chaincode is from my working folder, chaincode for employees in minifabric (tested).
//
// IMPORTANT:
// "documentDataCorrectiveMaintenance"   => this is the parameter from the function we use, in this case we use ID number e.g. "C-MTE-1000" to query corrective maintenance work order

func (t *DailyOperationsLogSmartContract) CallingCrossChaincodeFunctionReadCorrectiveMaintenanceWorkOrder(ctx contractapi.TransactionContextInterface, documentDataCorrectiveMaintenance string) (string, error) {

	if len(documentDataCorrectiveMaintenance) == 0 {
		return "", fmt.Errorf("please provide correct document data for")
	}

	params := []string{"MaintenanceSmartContract:ReadCorrectiveWorkOrder", documentDataCorrectiveMaintenance}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := ctx.GetStub().InvokeChaincode("mtemaintenance", queryArgs, "default-channel")

	return string(response.Payload), nil
}

//================================================================================================================================================
// CALL FUNCTION FROM OTHER CHAINCODE (TEST)
//================================================================================================================================================

// CallingCrossChaincodeFunctionCorrectiveMaintenance call for a function form the other smart contract,
// in this case calling the smart contract for corrective maintenance.
// "employee" chaincode is from my working folder, chaincode for employees in minifabric (tested).
//
// IMPORTANT:
// "documentDataCorrectiveMaintenance"   => this is the parameter from the function we use, in this case we use ID number e.g. "C-MTE-1000" to query corrective maintenance work order

func (t *DailyOperationsLogSmartContract) CallingCrossChaincodeFunctionCreateCorrectiveMaintenanceWorkOrder(ctx contractapi.TransactionContextInterface, documentParam0 string, documentParam1 string, documentParam2 string, documentParam3 string, documentParam4 string, documentParam5 string, documentParam6 string) (string, error) {

	if len(documentParam0) == 0 {
		return "", fmt.Errorf("please provide correct document data for")
	}

	params := []string{"MaintenanceSmartContract:CreateCorrectiveWorkOrder", documentParam0, documentParam1, documentParam2, documentParam3, documentParam4, documentParam5, documentParam6}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := ctx.GetStub().InvokeChaincode("mtemaintenance", queryArgs, "default-channel")

	return string(response.Payload), nil
}

//================================================================================================================================================
// CALL FUNCTION FROM OTHER CHAINCODE CREATING NEW DOWNTIME EVENT
//================================================================================================================================================

// CallingCrossChaincodeFunctionCreateNewDowntimeEvent call for a function form the other smart contract,
// in this case calling the smart contract for downtimeevent.
// The "downtimeEvent" chaincode is from my working folder, chaincode for employees in minifabric (tested).
//
// IMPORTANT:
// "documentDataCorrectiveMaintenance"   => this is the parameter from the function we use, in this case we use ID number e.g. "C-MTE-1000" to query corrective maintenance work order

func (t *DailyOperationsLogSmartContract) CallingCrossChaincodeFunctionCreateNewDowntimeEvent(ctx contractapi.TransactionContextInterface, documentParam0 string, documentParam1 string, documentParam2 string, documentParam3 string, documentParam4 string, documentParam5 string, documentParam6 string, documentParam7 string, documentParam8 string, documentParam9 string, documentParam10 string) (string, error) {

	if len(documentParam0) == 0 {
		return "", fmt.Errorf("please provide correct document data for")
	}

	paramsDowntimeEvent := []string{"DowntimeEventContract:NewDowntimeEvent", documentParam0, documentParam1, documentParam2, documentParam3, documentParam4, documentParam5, documentParam6, documentParam7, documentParam8, documentParam9, documentParam10}
	queryArgs := make([][]byte, len(paramsDowntimeEvent))
	for i, arg := range paramsDowntimeEvent {
		queryArgs[i] = []byte(arg)
	}

	response := ctx.GetStub().InvokeChaincode("downtimeevent", queryArgs, "default-channel")

	return string(response.Payload), nil
}
