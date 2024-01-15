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
	//"github.com/hyperledger/fabric-chaincode-go/pkg/cid"

	//"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const indexPrev = "ID~Type"

// const index = "ID~Type"

// "MaintenanceSmartContract" implements the fabric-contract-api-go programming model
type MaintenanceSmartContract struct {
	contractapi.Contract
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////                                                                                ///////////////////////////////////////
///////////////////////////           PREVENTIVE MAINTENANCE WORK ORDERS                                   ///////////////////////////////////////
///////////////////////////                                                                                ///////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// HistoryQueryResult structure used for returning result of history query
type HistoryQueryResult struct {
	Record    *PreventiveMaintenance `json:"record"`
	TxId      string                 `json:"txId"`
	Timestamp time.Time              `json:"timestamp"`
	IsDelete  bool                   `json:"isDelete"`
}

// PaginatedQueryResult structure used for returning paginated query results and metadata
type PaginatedQueryResult struct {
	Records             []*PreventiveMaintenance `json:"records"`
	FetchedRecordsCount int32                    `json:"fetchedRecordsCount"`
	Bookmark            string                   `json:"bookmark"`
}

// SetPrevWorkOrderTimestamp sets the timestamp of the work order to the current time
func (ba *PreventiveMaintenance) SetPrevWorkOrderTimestamp() {
	ba.PreventiveWorkOrderTimestamp = time.Now().Format("2006-01-02 15:04:05")
}

/*
// SetPrevWorkOrderCreator sets and get the MSP ID of the creator of the work order
func (ba *PreventiveMaintenance) SetPrevWorkOrderCreator(stub cid.ChaincodeStubInterface) {
	ba.PreventiveWorkOrderCreator, _ = cid.GetMSPID(stub)
}

// SetPrevWorkOrderConditionGood set the condition of the work order to mark as good
func (ba *PreventiveMaintenance) SetPrevWorkOrderConditionGood() {
	ba.PrevCondition = "good"
}

// SetPrevWorkOrderConditionNotSet set the condition of the work order to mark as not-good
func (ba *PreventiveMaintenance) SetPrevWorkOrderConditionNotSet() {
	ba.PrevCondition = "not-set"
}

// SetPrevWorkOrderValidatated set the condition of the work order to mark as good
func (ba *PreventiveMaintenance) SetPrevWorkOrderValidatated() {
	ba.PrevValidation = "validated"
}

// SetPrevWorkOrderNotValidatated set the condition of the work order to mark as not-good
func (ba *PreventiveMaintenance) SetPrevWorkOrderNotValidatated() {
	ba.PrevValidation = "not-set"
}
*/

//================================================================================================================================================
// CREATE PREVENTIVE WORK ORDER
//================================================================================================================================================

// CreatePreventiveWorkOrder creates a new work order on the ledger
func (t *MaintenanceSmartContract) CreatePreventiveWorkOrder(ctx contractapi.TransactionContextInterface, prev_work_order_id string, prev_work_order_type string, prev_work_order_description string, prev_general_instructions string, prev_planned_labour PlannedLabour, prev_maintenance_process_steps MaintenanceProcessSteps, prev_spare_parts_used MaintenanceParts) error {

	/*
		mspId, err := cid.GetMSPID(ctx.GetStub())

		if err != nil {
			return fmt.Errorf("failed while getting identity. %s", err.Error())
		}

		if mspId != "Org1MSP" {
			return fmt.Errorf("you are not authorized to create a new daily operations log")
		}
	*/

	exists, err := t.DoesPreventiveWorkOrderExists(ctx, prev_work_order_id)

	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	}
	if exists {
		return fmt.Errorf("asset already exists: %s", prev_work_order_id)
	}

	prevmtewo := &PreventiveMaintenance{
		DocType:                      "prevmtewo",
		PrevWorkOrderID:              prev_work_order_id,
		PrevWorkOrderType:            prev_work_order_type,
		PrevWorkOrderDescription:     prev_work_order_description,
		PrevGeneralHandSInstructions: prev_general_instructions,
		PrevPlannedLabour:            prev_planned_labour,
		PrevMaintenanceProcessSteps:  prev_maintenance_process_steps,
		PrevSparePartsUsed:           prev_spare_parts_used,
	}

	// Setting the MSP ID of the creator of the new work order.
	prevmtewo.SetPrevWorkOrderCreator(ctx.GetStub())

	// Setting the timestamp of the new work order to the current time.
	prevmtewo.SetPrevWorkOrderTimestamp()

	// Setting the next work order date adding one month to the current date.
	prevmtewo.SetNextPreventiveWorkOrder()

	// Setting the condition record to "not-set".
	prevmtewo.SetPrevWorkOrderConditionNotSet()

	// Setting the validation record for the new work order to "not-set" or not-validated.
	prevmtewo.SetPrevWorkOrderNotValidatated()

	assetBytes, err := json.Marshal(prevmtewo)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(prev_work_order_id, assetBytes)
	if err != nil {
		return err
	}

	//  Create an index to enable work order ID-type based range queries, e.g. return all assets with specific ID.
	//  An 'index' is a normal key-value entry in the ledger.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on indexName~PrevWorkOrderID~WorkOrderType.
	//  This will enable very efficient state range queries based on composite keys matching indexName~PrevWorkOrderID~*
	preventiveWorkOrderNameIndexKey, err := ctx.GetStub().CreateCompositeKey(indexPrev, []string{prevmtewo.PrevWorkOrderID, prevmtewo.PrevWorkOrderType})
	if err != nil {
		return err
	}
	//  Save index entry to world state. Only the key name is needed, no need to store a duplicate copy of the asset.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	return ctx.GetStub().PutState(preventiveWorkOrderNameIndexKey, value)
}

//================================================================================================================================================
// READ WORK ORDER
//================================================================================================================================================

// ReadWorkOrder retrieves an asset from the ledger
func (t *MaintenanceSmartContract) ReadPreventiveWorkOrder(ctx contractapi.TransactionContextInterface, prev_work_order_id string) (*PreventiveMaintenance, error) {
	assetBytes, err := ctx.GetStub().GetState(prev_work_order_id)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset %s: %v", prev_work_order_id, err)
	}
	if assetBytes == nil {
		return nil, fmt.Errorf("asset %s does not exist", prev_work_order_id)
	}

	var prevmtewo PreventiveMaintenance
	err = json.Unmarshal(assetBytes, &prevmtewo)
	if err != nil {
		return nil, err
	}

	return &prevmtewo, nil
}

//================================================================================================================================================
// DELETE WORK ORDER
//================================================================================================================================================

// DeleteWorkOrder removes an asset key-value pair from the ledger
func (t *MaintenanceSmartContract) DeletePreventiveWorkOrder(ctx contractapi.TransactionContextInterface, prev_work_order_id string) error {

	/*
		mspId, err := cid.GetMSPID(ctx.GetStub())
		if err != nil {
			return fmt.Errorf("failed while getting identity. %s", err.Error())
		}

		if mspId != "mte-operations-com" {
			return fmt.Errorf("you are not authorized to create a new daily operations log")
		}
	*/

	prevmtewo, err := t.ReadPreventiveWorkOrder(ctx, prev_work_order_id)
	if err != nil {
		return err
	}

	err = ctx.GetStub().DelState(prev_work_order_id)
	if err != nil {
		return fmt.Errorf("failed to delete asset %s: %v", prev_work_order_id, err)
	}

	preventiveWorkOrderNameIndexKey, err := ctx.GetStub().CreateCompositeKey(indexPrev, []string{prevmtewo.PrevWorkOrderID, prevmtewo.PrevWorkOrderType})
	if err != nil {
		return err
	}

	// Delete index entry
	return ctx.GetStub().DelState(preventiveWorkOrderNameIndexKey)
}

//================================================================================================================================================
// CONSTRUCT QUERY RESPONSE FROM ITERATOR
//================================================================================================================================================

// constructQueryResponseFromIterator constructs a slice of assets from the resultsIteratorPrev
func constructQueryResponseFromIterator(resultsIteratorPrev shim.StateQueryIteratorInterface) ([]*PreventiveMaintenance, error) {
	var assets []*PreventiveMaintenance
	for resultsIteratorPrev.HasNext() {
		queryResult, err := resultsIteratorPrev.Next()
		if err != nil {
			return nil, err
		}
		var prevmtewo PreventiveMaintenance
		err = json.Unmarshal(queryResult.Value, &prevmtewo)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &prevmtewo)
	}

	return assets, nil
}

//================================================================================================================================================
// GET WORK ORDERS BY RANGE
//================================================================================================================================================

// GetWorkOrdersByRange performs a range query based on the start and end keys provided.
// Read-only function results are not typically submitted to ordering. If the read-only
// results are submitted to ordering, or if the query is used in an update transaction
// and submitted to ordering, then the committing peers will re-execute to guarantee that
// result sets are stable between endorsement time and commit time. The transaction is
// invalidated by the committing peers if the result set has changed between endorsement
// time and commit time.
// Therefore, range queries are a safe option for performing update transactions based on query results.
func (t *MaintenanceSmartContract) GetPreventiveWorkOrdersByRange(ctx contractapi.TransactionContextInterface, startKey, endKey string) ([]*PreventiveMaintenance, error) {
	resultsIteratorPrev, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorPrev.Close()

	return constructQueryResponseFromIterator(resultsIteratorPrev)
}

//================================================================================================================================================
// QUERY WORK ORDER BY TYPE
//================================================================================================================================================

// QueryAssetsByType queries for assets based on the types name.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (type).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *MaintenanceSmartContract) QueryPreventiveWorkOrderByType(ctx contractapi.TransactionContextInterface, prev_work_order_type string) ([]*PreventiveMaintenance, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"prevmtewo","prev_work_order_type":"%s"}}`, prev_work_order_type)
	return getQueryResultForQueryString(ctx, queryString)
}

//================================================================================================================================================
// QUERY WORK ORDER BY CONDITION
//================================================================================================================================================

// QueryAssetsByCondition queries for assets based on their condition.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (condition).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *MaintenanceSmartContract) QueryPreventiveWorkOrderByCondition(ctx contractapi.TransactionContextInterface, prev_condition string) ([]*PreventiveMaintenance, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"prevmtewo","prev_condition":"%s"}}`, prev_condition)
	return getQueryResultForQueryString(ctx, queryString)
}

//================================================================================================================================================
// QUERY WORK ORDER BY VALIDATION
//================================================================================================================================================

// QueryPreventiveWorkOrderByValidation queries for work orders based on their validation.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (validation).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *MaintenanceSmartContract) QueryPreventiveWorkOrderByValidation(ctx contractapi.TransactionContextInterface, prev_validation string) ([]*PreventiveMaintenance, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"prevmtewo","prev_validation":"%s"}}`, prev_validation)
	return getQueryResultForQueryString(ctx, queryString)
}

//================================================================================================================================================
// QUERY WORK ORDER BY ID
//================================================================================================================================================

// QueryWorkOrderByID queries for work orders based on the its ID number.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (type).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *MaintenanceSmartContract) QueryPreventiveWorkOrderByID(ctx contractapi.TransactionContextInterface, prev_work_order_id string) ([]*PreventiveMaintenance, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"prevmtewo","prev_work_order_id":"%s"}}`, prev_work_order_id)
	return getQueryResultForQueryString(ctx, queryString)
}

//================================================================================================================================================
// QUERY WORK ORDER
//================================================================================================================================================

// QueryPreventiveWorkOrders uses a query string to perform a query for work orders.
// Query string matching state database syntax is passed in and executed as is.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the "QueryAssetsForOwner" example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Ad hoc rich query
func (t *MaintenanceSmartContract) QueryPreventiveWorkOrders(ctx contractapi.TransactionContextInterface, queryString string) ([]*PreventiveMaintenance, error) {
	return getQueryResultForQueryString(ctx, queryString)
}

//================================================================================================================================================
// GET QUERY RESULT FOR QUERY STRING
//================================================================================================================================================

// getQueryResultForQueryString executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]*PreventiveMaintenance, error) {
	resultsIteratorPrev, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorPrev.Close()

	return constructQueryResponseFromIterator(resultsIteratorPrev)
}

//================================================================================================================================================
// GET PREVENTIVE WORK ORDERS BY RANGE WITH PAGINATION
//================================================================================================================================================

// GetPreventiveWorkOrdersByRangeWithPagination performs a range query based on the start and end key,
// page size and a bookmark.
// The number of fetched records will be equal to or lesser than the page size.
// Paginated range queries are only valid for read only transactions.
// Example: Pagination with Range Query
func (t *MaintenanceSmartContract) GetPreventiveWorkOrdersByRangeWithPagination(ctx contractapi.TransactionContextInterface, startKey string, endKey string, pageSize int, bookmark string) (*PaginatedQueryResult, error) {

	resultsIteratorPrev, responseMetadata, err := ctx.GetStub().GetStateByRangeWithPagination(startKey, endKey, int32(pageSize), bookmark)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorPrev.Close()

	assets, err := constructQueryResponseFromIterator(resultsIteratorPrev)
	if err != nil {
		return nil, err
	}

	return &PaginatedQueryResult{
		Records:             assets,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil
}

//================================================================================================================================================
// QUERY PREVENTIVE WORK ORDERS WITH PAGINATION
//================================================================================================================================================

// QueryPreventiveWithPagination uses a query string, page size and a bookmark to perform a query
// for assets. Query string matching state database syntax is passed in and executed as is.
// The number of fetched records would be equal to or lesser than the specified page size.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the QueryAssetsForOwner example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// Paginated queries are only valid for read only transactions.
// Example: Pagination with Ad hoc Rich Query
func (t *MaintenanceSmartContract) QueryPreventiveWorkOrdersWithPagination(ctx contractapi.TransactionContextInterface, queryString string, pageSize int, bookmark string) (*PaginatedQueryResult, error) {

	return getQueryResultForQueryStringWithPagination(ctx, queryString, int32(pageSize), bookmark)
}

//================================================================================================================================================
// GET QUERY RESULT FOR QUERY STRING WITH PAGINATION
//================================================================================================================================================

// getQueryResultForQueryStringWithPagination executes the passed in query string with
// pagination info. The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryStringWithPagination(ctx contractapi.TransactionContextInterface, queryString string, pageSize int32, bookmark string) (*PaginatedQueryResult, error) {

	resultsIteratorPrev, responseMetadata, err := ctx.GetStub().GetQueryResultWithPagination(queryString, pageSize, bookmark)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorPrev.Close()

	assets, err := constructQueryResponseFromIterator(resultsIteratorPrev)
	if err != nil {
		return nil, err
	}

	return &PaginatedQueryResult{
		Records:             assets,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil
}

//================================================================================================================================================
// GET WORK ORDER HISTORY
//================================================================================================================================================

// GetPreventiveWorkOrderHistory returns the chain of custody for an asset since issuance.
func (t *MaintenanceSmartContract) GetPreventiveWorkOrderHistory(ctx contractapi.TransactionContextInterface, prev_work_order_id string) ([]HistoryQueryResult, error) {
	log.Printf("GetAssetHistory: ID %v", prev_work_order_id)

	resultsIteratorPrev, err := ctx.GetStub().GetHistoryForKey(prev_work_order_id)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorPrev.Close()

	var records []HistoryQueryResult
	for resultsIteratorPrev.HasNext() {
		response, err := resultsIteratorPrev.Next()
		if err != nil {
			return nil, err
		}

		var prevmtewo PreventiveMaintenance
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &prevmtewo)
			if err != nil {
				return nil, err
			}
		} else {
			prevmtewo = PreventiveMaintenance{
				PrevWorkOrderID: prev_work_order_id,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			return nil, err
		}

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: timestamp,
			Record:    &prevmtewo,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}

//================================================================================================================================================
// DOES WORK ORDER EXISTS
//================================================================================================================================================

// DoesWorkOrderExists returns true when asset with given ID exists in the ledger.
func (t *MaintenanceSmartContract) DoesPreventiveWorkOrderExists(ctx contractapi.TransactionContextInterface, prev_work_order_id string) (bool, error) {
	assetBytes, err := ctx.GetStub().GetState(prev_work_order_id)
	if err != nil {
		return false, fmt.Errorf("failed to read asset %s from world state. %v", prev_work_order_id, err)
	}

	return assetBytes != nil, nil
}

//================================================================================================================================================
// UPDATE WORK ORDER
//================================================================================================================================================

// UpdateWorkOrder updates an existing work order in the world state with provided parameters.
func (t *MaintenanceSmartContract) UpdateWorkOrder(ctx contractapi.TransactionContextInterface, prev_work_order_id string, prev_work_order_type string, prev_work_order_description string, prev_general_instructions string, prev_planned_labour PlannedLabour, prev_maintenance_process_steps MaintenanceProcessSteps, prev_spare_parts_used MaintenanceParts) error {
	exists, err := t.DoesPreventiveWorkOrderExists(ctx, prev_work_order_id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", prev_work_order_id)
	}

	// overwriting original asset with new asset
	prevmtewo := PreventiveMaintenance{
		DocType:                      "prevmtewo",
		PrevWorkOrderID:              prev_work_order_id,
		PrevWorkOrderType:            prev_work_order_type,
		PrevWorkOrderDescription:     prev_work_order_description,
		PrevGeneralHandSInstructions: prev_general_instructions,
		PrevPlannedLabour:            prev_planned_labour,
		PrevMaintenanceProcessSteps:  prev_maintenance_process_steps,
		PrevSparePartsUsed:           prev_spare_parts_used,
	}

	// Setting the MSP ID of the creator of the new work order.
	prevmtewo.SetPrevWorkOrderCreator(ctx.GetStub())

	// Setting the timestamp of the new work order to the current time.
	prevmtewo.SetPrevWorkOrderTimestamp()

	// Setting the condition record to "good".
	prevmtewo.SetPrevWorkOrderConditionGood()

	// Setting the validation record for the new work order to "validated".
	prevmtewo.SetPrevWorkOrderValidatated()

	assetJSON, err := json.Marshal(prevmtewo)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(prev_work_order_id, assetJSON)
}

/*
//================================================================================================================================================
// GET ALL WORK ORDERS
//================================================================================================================================================

// GetAllWorkOrders returns all assets found in world state
func (t *MaintenanceSmartContract) GetAllPreventiveWorkOrders(ctx contractapi.TransactionContextInterface) ([]*PreventiveMaintenance, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.

	resultsIteratorPrev, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIteratorPrev.Close()

	var assets []*PreventiveMaintenance
	for resultsIteratorPrev.HasNext() {
		queryResponse, err := resultsIteratorPrev.Next()
		if err != nil {
			return nil, err
		}

		var asset PreventiveMaintenance
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
*/

//================================================================================================================================================
// INIT LEDGER
//================================================================================================================================================
/*
// InitLedger creates the initial set of daily, monthly, and annually work orders in the ledger.
func (t *MaintenanceSmartContract) InitLedgerWithPreventiveWorkOrders(ctx contractapi.TransactionContextInterface) error {
	assets := []PreventiveMaintenance{
		{DocType: "prevmtewo", PrevWorkOrderID: "D-MTE-1010", PrevWorkOrderType: "Preventive Maintenance Work Order", PrevWorkOrderDescription: "Daily Startup Inspection - Drive Station", PrevGeneralHandSInstructions: "General Health and Safety instructions", PrevPlannedLabour: "Maintenance Technician", PrevMaintenanceProcessSteps: "[\"Maintenance Process Step 1\",\"Maintenance Process Step 2\", \"Maintenance Process Step 3\", \"Maintenance Process Step 4\", \"Maintenance Process Step 5\"]", PrevSparePartsUsed: "n/a", PrevCondition: "not-set", PrevValidation: "not-set"}, {DocType: "prevmtewo", PrevWorkOrderID: "D-MTE-1011", PrevWorkOrderType: "Preventive Maintenance Work Order", PrevWorkOrderDescription: "Daily Startup Inspection - Return Station", PrevGeneralHandSInstructions: "General Health and Safety instructions", PrevPlannedLabour: "Maintenance Technician", PrevMaintenanceProcessSteps: "[\"Maintenance Process Step 1\",\"Maintenance Process Step 2\", \"Maintenance Process Step 3\", \"Maintenance Process Step 4\", \"Maintenance Process Step 5\"]", PrevSparePartsUsed: "n/a", PrevCondition: "not-set", PrevValidation: "not-set"},
		{DocType: "prevmtewo", PrevWorkOrderID: "D-MTE-1012", PrevWorkOrderType: "Preventive Maintenance Work Order", PrevWorkOrderDescription: "Daily Startup Inspection - Test Run", PrevGeneralHandSInstructions: "General Health and Safety instructions", PrevPlannedLabour: "Maintenance Technician", PrevMaintenanceProcessSteps: "[\"Maintenance Process Step 1\",\"Maintenance Process Step 2\", \"Maintenance Process Step 3\", \"Maintenance Process Step 4\", \"Maintenance Process Step 5\"]", PrevSparePartsUsed: "n/a", PrevCondition: "not-set", PrevValidation: "not-set"},
		{DocType: "prevmtewo", PrevWorkOrderID: "M-MTE-1020", PrevWorkOrderType: "Preventive Maintenance Work Order", PrevWorkOrderDescription: "Monthly Carriers Inspection", PrevGeneralHandSInstructions: "General Health and Safety instructions", PrevPlannedLabour: "Maintenance Technician", PrevMaintenanceProcessSteps: "[\"Maintenance Process Step 1\",\"Maintenance Process Step 2\", \"Maintenance Process Step 3\", \"Maintenance Process Step 4\", \"Maintenance Process Step 5\"]", PrevSparePartsUsed: "n/a", PrevCondition: "not-set", PrevValidation: "not-set"},
		{DocType: "prevmtewo", PrevWorkOrderID: "M-MTE-1021", PrevWorkOrderType: "Preventive Maintenance Work Order", PrevWorkOrderDescription: "Monthly Line Equipment Inspection", PrevGeneralHandSInstructions: "General Health and Safety instructions", PrevPlannedLabour: "Maintenance Technician", PrevMaintenanceProcessSteps: "[\"Maintenance Process Step 1\",\"Maintenance Process Step 2\", \"Maintenance Process Step 3\", \"Maintenance Process Step 4\"]", PrevSparePartsUsed: "n/a", PrevCondition: "not-set", PrevValidation: "not-set"},
		{DocType: "prevmtewo", PrevWorkOrderID: "Y-MTE-1060", PrevWorkOrderType: "Preventive Maintenance Work Order", PrevWorkOrderDescription: "Annual Station Inspection", PrevGeneralHandSInstructions: "General Health and Safety instructions", PrevPlannedLabour: "Maintenance Technician", PrevMaintenanceProcessSteps: "[\"Maintenance Process Step 1\",\"Maintenance Process Step 2\", \"Maintenance Process Step 3\", \"Maintenance Process Step 4\", \"Maintenance Process Step 5\",\"Maintenance Process Step 6\"]", PrevSparePartsUsed: "n/a", PrevCondition: "not-set", PrevValidation: "not-set"},
	}

	for _, asset := range assets {
		err := t.CreatePreventiveWorkOrder(ctx, asset.PrevWorkOrderID, asset.PrevWorkOrderType, asset.PrevWorkOrderDescription, asset.PrevGeneralHandSInstructions, asset.PrevPlannedLabour, asset.PrevMaintenanceProcessSteps, asset.PrevSparePartsUsed)
		if err != nil {
			return err
		}
	}

	/*
		for _, asset := range assets {
			err := t.CreatePreventiveWorkOrder(ctx, asset.PrevWorkOrderID, asset.PrevWorkOrderType, asset.PrevWorkOrderDescription, asset.PrevGeneralHandSInstructions, asset.PrevPlannedLabour, asset.PrevMaintenanceProcessSteps, asset.PrevSparePartsUsed, asset.PrevCondition, asset.PrevValidation)
			if err != nil {
				return err
			}
		}
*/
/*
	return nil
}
*/

//================================================================================================================================================
// MAIN FUNCTION
//================================================================================================================================================
/*
func main() {
	chaincode, err := contractapi.NewChaincode(&PreventiveMaintenanceChaincode{})
	if err != nil {
		log.Panicf("Error creating maintenance work order chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting maintenance work order chaincode: %v", err)
	}
}
*/

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////                                                                                ///////////////////////////////////////
///////////////////////////           CORRECTIVE MAINTENANCE WORK ORDERS                                   ///////////////////////////////////////
///////////////////////////                                                                                ///////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const indexCorr = "ID~Type"

// CorrectiveHistoryQueryResult structure used for returning result of history query
type CorrectiveHistoryQueryResult struct {
	Record    *CorrectiveMaintenance `json:"record"`
	TxId      string                 `json:"txId"`
	Timestamp time.Time              `json:"timestamp"`
	IsDelete  bool                   `json:"isDelete"`
}

// CorrectivePaginatedQueryRequest structure used for returning paginated query results and metadata
type CorrectivePaginatedQueryResult struct {
	Records             []*CorrectiveMaintenance `json:"records"`
	FetchedRecordsCount int32                    `json:"fetchedRecordsCount"`
	Bookmark            string                   `json:"bookmark"`
}

//================================================================================================================================================
// CREATE CORRECTIVE WORK ORDER
//================================================================================================================================================

// SetPrevWorkOrderTimestamp sets the timestamp of the work order to the current time
func (ba *CorrectiveMaintenance) SetCorrWorkOrderTimestamp() {
	ba.CorrWorkOrderTimestamp = time.Now().Format("2006-01-02 15:04:05")
}

// CreateCorrectiveWorkOrder initializes a new asset in the ledger
func (t *MaintenanceSmartContract) CreateCorrectiveWorkOrder(ctx contractapi.TransactionContextInterface, corr_work_order_id string, corr_work_order_type string, corr_work_order_description string, corr_general_instructions string, corr_planned_labour PlannedLabour, corr_maintenance_process_steps MaintenanceProcessSteps, corr_spare_parts_used MaintenanceParts) error {

	exists, err := t.DoesCorrectiveWorkOrderExists(ctx, corr_work_order_id)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	}
	if exists {
		return fmt.Errorf("asset already exists: %s", corr_work_order_id)
	}

	corrmtewo := &CorrectiveMaintenance{
		DocType:                      "corrmtewo",
		CorrWorkOrderID:              corr_work_order_id,
		CorrWorkOrderType:            corr_work_order_type,
		CorrWorkOrderDescription:     corr_work_order_description,
		CorrGeneralHandSInstructions: corr_general_instructions,
		CorrPlannedLabour:            corr_planned_labour,
		CorrMaintenanceProcessSteps:  corr_maintenance_process_steps,
		CorrSparePartsUsed:           corr_spare_parts_used,
	}

	// Setting work order creator
	corrmtewo.SetCorrWorkOrderCreator(ctx.GetStub())

	// Setting the timestamp of the new work order to the current time.
	corrmtewo.SetCorrWorkOrderTimestamp()

	// Setting the date of completion of the new work order adding 15 day, for example.
	corrmtewo.SetCorrWorkOrderCompletionDate()

	// Setting the condition record to not set.
	corrmtewo.SetCorrWorkOrderConditionNotGood()

	// Setting the validation record for the new work order to not validated.
	corrmtewo.SetCorrWorkOrderNotValidated()

	assetBytes, err := json.Marshal(corrmtewo)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(corr_work_order_id, assetBytes)
	if err != nil {
		return err
	}

	//  Create an index to enable work order ID-type based range queries, e.g. return all assets with specific ID.
	//  An 'index' is a normal key-value entry in the ledger.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on indexNCorrWorkOrderID~WorkOrderType.
	//  This will enable very efficient state range queries based on composite keys matching indexNCorrWorkOrderID~*
	workOrderNameIndexKey, err := ctx.GetStub().CreateCompositeKey(indexCorr, []string{corrmtewo.CorrWorkOrderID, corrmtewo.CorrWorkOrderType})
	if err != nil {
		return err
	}
	//  Save index entry to world state. Only the key name is needed, no need to store a duplicate copy of the asset.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	return ctx.GetStub().PutState(workOrderNameIndexKey, value)
}

//================================================================================================================================================
// READ CORRECTIVE WORK ORDER
//================================================================================================================================================

// ReadCorrectiveWorkOrder retrieves an asset from the ledger
func (t *MaintenanceSmartContract) ReadCorrectiveWorkOrder(ctx contractapi.TransactionContextInterface, corr_work_order_id string) (*CorrectiveMaintenance, error) {

	assetBytes, err := ctx.GetStub().GetState(corr_work_order_id)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset %s: %v", corr_work_order_id, err)
	}
	if assetBytes == nil {
		return nil, fmt.Errorf("asset %s does not exist", corr_work_order_id)
	}

	var corrmtewo CorrectiveMaintenance
	err = json.Unmarshal(assetBytes, &corrmtewo)
	if err != nil {
		return nil, err
	}

	return &corrmtewo, nil
}

//================================================================================================================================================
// DELETE CORRECTIVE WORK ORDER
//================================================================================================================================================

// DeleteCorrectiveWorkOrder removes an asset key-value pair from the ledger
func (t *MaintenanceSmartContract) DeleteCorrectiveWorkOrder(ctx contractapi.TransactionContextInterface, corr_work_order_id string) error {

	/*
		mspId, err := cid.GetMSPID(ctx.GetStub())
		if err != nil {
			return fmt.Errorf("failed while getting identity. %s", err.Error())
		}

		if mspId != "mte-operations-com" {
			return fmt.Errorf("you are not authorized to create a new daily operations log")
		}
	*/

	corrmtewo, err := t.ReadCorrectiveWorkOrder(ctx, corr_work_order_id)
	if err != nil {
		return err
	}

	err = ctx.GetStub().DelState(corr_work_order_id)
	if err != nil {
		return fmt.Errorf("failed to delete asset %s: %v", corr_work_order_id, err)
	}

	workorderNameIndexKey, err := ctx.GetStub().CreateCompositeKey(indexCorr, []string{corrmtewo.CorrWorkOrderID, corrmtewo.CorrWorkOrderType})
	if err != nil {
		return err
	}

	// Delete index entry
	return ctx.GetStub().DelState(workorderNameIndexKey)
}

//================================================================================================================================================
// CONSTRUCT QUERY RESPONSE FROM ITERATOR
//================================================================================================================================================

//  Constructs a slice of assets from the resultsIteratorCorr
func constructCorrectiveQueryResponseFromIterator(resultsIteratorCorr shim.StateQueryIteratorInterface) ([]*CorrectiveMaintenance, error) {
	var assetsCorr []*CorrectiveMaintenance
	for resultsIteratorCorr.HasNext() {
		queryResult, err := resultsIteratorCorr.Next()
		if err != nil {
			return nil, err
		}
		var corrmtewo CorrectiveMaintenance
		err = json.Unmarshal(queryResult.Value, &corrmtewo)
		if err != nil {
			return nil, err
		}
		assetsCorr = append(assetsCorr, &corrmtewo)
	}

	return assetsCorr, nil
}

//================================================================================================================================================
// GET CORRECTIVE WORK ORDERS BY RANGE
//================================================================================================================================================

// GetCorrectiveWorkOrdersByRange performs a range query based on the start and end keys provided.
// Read-only function results are not typically submitted to ordering. If the read-only
// results are submitted to ordering, or if the query is used in an update transaction
// and submitted to ordering, then the committing peers will re-execute to guarantee that
// result sets are stable between endorsement time and commit time. The transaction is
// invalidated by the committing peers if the result set has changed between endorsement
// time and commit time.
// Therefore, range queries are a safe option for performing update transactions based on query results.
func (t *MaintenanceSmartContract) GetCorrectiveWorkOrdersByRange(ctx contractapi.TransactionContextInterface, startKey, endKey string) ([]*CorrectiveMaintenance, error) {
	resultsIteratorCorr, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorCorr.Close()

	return constructCorrectiveQueryResponseFromIterator(resultsIteratorCorr)
}

//================================================================================================================================================
// QUERY CORRECTIVE WORK ORDER BY TYPE
//================================================================================================================================================

// QueryCorrectiveWorkOrderByType queries for assets based on the types name.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (type).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *MaintenanceSmartContract) QueryCorrectiveWorkOrderByType(ctx contractapi.TransactionContextInterface, corr_work_order_type string) ([]*CorrectiveMaintenance, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"corrmtewo","corr_work_order_type":"%s"}}`, corr_work_order_type)
	return getCorrectiveQueryResultForQueryString(ctx, queryString)
}

//================================================================================================================================================
// QUERY CORRECTIVE WORK ORDER BY CONDITION
//================================================================================================================================================

// QueryCorrectiveWorkOrderByCondition queries for assets based on their condition.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (condition).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *MaintenanceSmartContract) QueryCorrectiveWorkOrderByCondition(ctx contractapi.TransactionContextInterface, corr_condition string) ([]*CorrectiveMaintenance, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"corrmtewo","corr_condition":"%s"}}`, corr_condition)
	return getCorrectiveQueryResultForQueryString(ctx, queryString)
}

//================================================================================================================================================
// QUERY CORRECTIVE WORK ORDER BY VALIDATION
//================================================================================================================================================

// QueryCorrectiveWorkOrderByValidation queries for work orders based on their validation.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (validation).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *MaintenanceSmartContract) QueryCorrectiveWorkOrderByValidation(ctx contractapi.TransactionContextInterface, corr_validation string) ([]*CorrectiveMaintenance, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"corrmtewo","corr_validation":"%s"}}`, corr_validation)
	return getCorrectiveQueryResultForQueryString(ctx, queryString)
}

//================================================================================================================================================
// QUERY CORRECTIVE WORK ORDER BY ID
//================================================================================================================================================

// QueryCorrectiveWorkOrderByID queries for work orders based on the its ID number.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (type).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *MaintenanceSmartContract) QueryCorrectiveWorkOrderByID(ctx contractapi.TransactionContextInterface, corr_work_order_id string) ([]*CorrectiveMaintenance, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"corrmtewo","corr_work_order_id":"%s"}}`, corr_work_order_id)
	return getCorrectiveQueryResultForQueryString(ctx, queryString)
}

//================================================================================================================================================
// QUERY CORRECTIVE WORK ORDER
//================================================================================================================================================

// QueryCorrectiveWorkOrders uses a query string to perform a query for work orders.
// Query string matching state database syntax is passed in and executed as is.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the "QueryAssetsForOwner" example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Ad hoc rich query
func (t *MaintenanceSmartContract) QueryCorrectiveWorkOrders(ctx contractapi.TransactionContextInterface, queryString string) ([]*CorrectiveMaintenance, error) {
	return getQueryResultForQueryStringCorrective(ctx, queryString)
}

//================================================================================================================================================
// GET QUERY RESULT FOR QUERY STRING CORRECTIVE
//================================================================================================================================================

// getQueryResultForQueryString executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryStringCorrective(ctx contractapi.TransactionContextInterface, queryString string) ([]*CorrectiveMaintenance, error) {
	resultsIteratorCorr, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorCorr.Close()

	return constructCorrectiveQueryResponseFromIterator(resultsIteratorCorr)
}

//================================================================================================================================================
// QUERY CORRECTIVE WORK ORDER
//================================================================================================================================================

// QueryCorrectiveWorkOrders uses a query string to perform a query for work orders.
// Query string matching resultsIteratorCorr
// getCorrectiveCorrectiveQueryResultForQueryString executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getCorrectiveQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]*CorrectiveMaintenance, error) {
	resultsIteratorCorr, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorCorr.Close()

	return constructCorrectiveQueryResponseFromIterator(resultsIteratorCorr)
}

//================================================================================================================================================
// GET CORRECTIVE WORK ORDERS BY RANGE WITH PAGINATION
//================================================================================================================================================

// GetCorrectiveWorkOrdersByRangeWithPagination performs a range query based on the start and end key,
// page size and a bookmark.
// The number of fetched records will be equal to or lesser than the page size.
// Paginated range queries are only valid for read only transactions.
// Example: Pagination with Range Query
func (t *MaintenanceSmartContract) GetCorrectiveWorkOrdersByRangeWithPagination(ctx contractapi.TransactionContextInterface, startKey string, endKey string, pageSize int, bookmark string) (*CorrectivePaginatedQueryResult, error) {

	resultsIteratorCorr, responseMetadata, err := ctx.GetStub().GetStateByRangeWithPagination(startKey, endKey, int32(pageSize), bookmark)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorCorr.Close()

	assetsCorr, err := constructCorrectiveQueryResponseFromIterator(resultsIteratorCorr)
	if err != nil {
		return nil, err
	}

	return &CorrectivePaginatedQueryResult{
		Records:             assetsCorr,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil
}

//================================================================================================================================================
// QUERY CORRECTIVE WORK ORDERS WITH PAGINATION
//================================================================================================================================================

// QueryCorrectiveWorkOrdersWithPagination uses a query string, page size and a bookmark to perform a query
// for assets. Query string matching state database syntax is passed in and executed as is.
// The number of fetched records would be equal to or lesser than the specified page size.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the QueryAssetsForOwner example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// Paginated queries are only valid for read only transactions.
// Example: Pagination with Ad hoc Rich Query
func (t *MaintenanceSmartContract) QueryCorrectiveWorkOrdersWithPagination(ctx contractapi.TransactionContextInterface, queryString string, pageSize int, bookmark string) (*CorrectivePaginatedQueryResult, error) {

	return getCorrectiveQueryResultForQueryStringWithPagination(ctx, queryString, int32(pageSize), bookmark)
}

//================================================================================================================================================
// GET QUERY RESULT FOR QUERY STRING WITH PAGINATION
//================================================================================================================================================

// getCorrectiveQueryResultForQueryStringWithPagination executes the passed in query string with
// pagination info. The result set is built and returned as a byte array containing the JSON results.
func getCorrectiveQueryResultForQueryStringWithPagination(ctx contractapi.TransactionContextInterface, queryString string, pageSize int32, bookmark string) (*CorrectivePaginatedQueryResult, error) {

	resultsIteratorCorr, responseMetadata, err := ctx.GetStub().GetQueryResultWithPagination(queryString, pageSize, bookmark)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorCorr.Close()

	assetsCorr, err := constructCorrectiveQueryResponseFromIterator(resultsIteratorCorr)
	if err != nil {
		return nil, err
	}

	return &CorrectivePaginatedQueryResult{
		Records:             assetsCorr,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil
}

//================================================================================================================================================
// GET CORRECTIVE WORK ORDER HISTORY
//================================================================================================================================================

// GetCorrectiveWorkOrderHistory returns the chain of custody for an asset since issuance.
func (t *MaintenanceSmartContract) GetCorrectiveWorkOrderHistory(ctx contractapi.TransactionContextInterface, corr_work_order_id string) ([]CorrectiveHistoryQueryResult, error) {
	log.Printf("GetAssetHistory: ID %v", corr_work_order_id)

	resultsIteratorCorr, err := ctx.GetStub().GetHistoryForKey(corr_work_order_id)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorCorr.Close()

	var records []CorrectiveHistoryQueryResult
	for resultsIteratorCorr.HasNext() {
		response, err := resultsIteratorCorr.Next()
		if err != nil {
			return nil, err
		}

		var corrmtewo CorrectiveMaintenance
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &corrmtewo)
			if err != nil {
				return nil, err
			}
		} else {
			corrmtewo = CorrectiveMaintenance{
				CorrWorkOrderID: corr_work_order_id,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)

		if err != nil {
			return nil, err
		}

		record := CorrectiveHistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: timestamp,
			Record:    &corrmtewo,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}

//================================================================================================================================================
// DOES CORRECTIVE WORK ORDER EXISTS
//================================================================================================================================================

// DoesCorrectiveWorkOrderExists returns true when asset with given ID exists in the ledger.
func (t *MaintenanceSmartContract) DoesCorrectiveWorkOrderExists(ctx contractapi.TransactionContextInterface, corr_work_order_id string) (bool, error) {
	assetBytes, err := ctx.GetStub().GetState(corr_work_order_id)
	if err != nil {
		return false, fmt.Errorf("failed to read asset %s from world state. %v", corr_work_order_id, err)
	}

	return assetBytes != nil, nil
}

//================================================================================================================================================
// UPDATE CORRECTIVE WORK ORDER
//================================================================================================================================================

// UpdateCorrectiveWorkOrder updates an existing work order in the world state with provided parameters.
func (t *MaintenanceSmartContract) UpdateCorrectiveWorkOrder(ctx contractapi.TransactionContextInterface, corr_work_order_id string, corr_work_order_type string, corr_work_order_description string, corr_general_instructions string, corr_planned_labour PlannedLabour, corr_maintenance_process_steps MaintenanceProcessSteps, corr_spare_parts_used MaintenanceParts) error {

	/*
		mspId, err := cid.GetMSPID(ctx.GetStub())
		if err != nil {
			return fmt.Errorf("failed while getting identity. %s", err.Error())
		}

		if mspId != "mte-operations-com" {
			return fmt.Errorf("you are not authorized to create a new daily operations log")
		}
	*/

	exists, err := t.DoesCorrectiveWorkOrderExists(ctx, corr_work_order_id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", corr_work_order_id)
	}

	// Overwriting original asset with new asset
	corrmtewo := CorrectiveMaintenance{
		DocType:                      "corrmtewo",
		CorrWorkOrderID:              corr_work_order_id,
		CorrWorkOrderType:            corr_work_order_type,
		CorrWorkOrderDescription:     corr_work_order_description,
		CorrGeneralHandSInstructions: corr_general_instructions,
		CorrPlannedLabour:            corr_planned_labour,
		CorrMaintenanceProcessSteps:  corr_maintenance_process_steps,
		CorrSparePartsUsed:           corr_spare_parts_used,
	}

	// Setting work order creator
	corrmtewo.SetCorrWorkOrderCreator(ctx.GetStub())

	// Setting the timestamp of the new work order to the current time.
	corrmtewo.SetCorrWorkOrderTimestamp()

	// Setting the condition record to good.
	corrmtewo.SetCorrWorkOrderConditionGood()

	// Setting the validation record for the new work order to validated.
	corrmtewo.SetCorrWorkOrderValidated()

	assetJSON, err := json.Marshal(corrmtewo)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(corr_work_order_id, assetJSON)
}

/*
//================================================================================================================================================
// GET ALL CORRECTIVE WORK ORDERS
//================================================================================================================================================

// GetAllCorrectiveWorkOrders returns all assets found in world state
func (t *MaintenanceSmartContract) GetAllCorrectiveWorkOrders(ctx contractapi.TransactionContextInterface) ([]*CorrectiveMaintenance, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIteratorCorr, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIteratorCorr.Close()

	var assets []*CorrectiveMaintenance
	for resultsIteratorCorr.HasNext() {
		queryResponse, err := resultsIteratorCorr.Next()
		if err != nil {
			return nil, err
		}

		var asset CorrectiveMaintenance
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
*/

//================================================================================================================================================
// MAIN FUNCTION
//================================================================================================================================================

/*
func main() {
	chaincode, err := contractapi.NewChaincode(&CorrectiveMaintenanceChaincode{})
	if err != nil {
		log.Panicf("Error creating maintenance work order chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting maintenance work order chaincode: %v", err)
	}
}
*/

//================================================================================================================================================
// CALL FUNCTION FROM OTHER CHAINCODE (TEST)
//================================================================================================================================================

// CallingCrossChaincodeFunctionCorrectiveMaintenance call for a function form the other smart contract,
// in this case calling the smart contract for corrective maintenance.
// "employee" chaincode is from my working folder, chaincode for employees in minifabric (tested).
//
// IMPORTANT:
// "documentDataCorrectiveMaintenance"   => this is the parameter from the function we use, in this case we use ID number
// e.g. "C-MTE-1000" to query corrective maintenance work order

func (t *MaintenanceSmartContract) CallingCrossChaincodeFunctionReadAsset(ctx contractapi.TransactionContextInterface, asset_id string) (string, error) {

	if len(asset_id) == 0 {
		return "", fmt.Errorf("please provide correct document data for")
	}

	params := []string{"AssetsSmartContract:ReadAsset", asset_id}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := ctx.GetStub().InvokeChaincode("assetoperations", queryArgs, "default-channel")

	return string(response.Payload), nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////                                                                                ///////////////////////////////////////
///////////////////////////           FAILURE MAINTENANCE WORK ORDERS                                      ///////////////////////////////////////
///////////////////////////                                                                                ///////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const indexFail = "ID~Type"

// "FailureHistoryQueryResult" structure used for returning result of history query
type FailureHistoryQueryResult struct {
	Record    *FailureMaintenance `json:"record"`
	TxId      string              `json:"txId"`
	Timestamp time.Time           `json:"timestamp"`
	IsDelete  bool                `json:"isDelete"`
}

// "FailurePaginatedQueryRequest" structure used for returning paginated query results and metadata
type FailurePaginatedQueryResult struct {
	Records             []*FailureMaintenance `json:"records"`
	FetchedRecordsCount int32                 `json:"fetchedRecordsCount"`
	Bookmark            string                `json:"bookmark"`
}

// Set FailWorkOrderTimestamp sets the timestamp of the work order
func (ba *FailureMaintenance) SetFailWorkOrderTimestamp() {
	ba.FailWorkOrderTimestamp = time.Now().Format("2006-01-02 15:04:05")
}

/*
// SetFailConditionGood set the condition of the work order to mark as good
func (ba *FailureMaintenance) SetFailWorkOrderConditionGood() {
	ba.FailCondition = "good"
}

// SetFailConditionNotGood set the condition of the work order to mark as not-good
func (ba *FailureMaintenance) SetFailWorkOrderConditionNotGood() {
	ba.FailCondition = "not-good"
}

// SetFailWorkOrderValidated set the condition of the work order to mark as good
func (ba *FailureMaintenance) SetFailWorkOrderValidated() {
	ba.FailValidation = "validated"
}

// SetFailWorkOrderNotValidated set the condition of the work order to mark as good
func (ba *FailureMaintenance) SetFailWorkOrderNotValidated() {
	ba.FailValidation = "not-validated"
}
*/

//================================================================================================================================================
// CREATE FAILURE WORK ORDER
//================================================================================================================================================

// CreateFailureWorkOrder initializes a new asset in the ledger
func (t *MaintenanceSmartContract) CreateFailureWorkOrder(ctx contractapi.TransactionContextInterface, fail_work_order_id string, fail_work_order_type string, fault_id string, fail_work_order_description string, fail_general_instructions string, fail_planned_labour PlannedLabour, fail_maintenance_process_steps MaintenanceProcessSteps, fail_spare_parts_used MaintenanceParts) error {

	exists, err := t.DoesFailureWorkOrderExists(ctx, fail_work_order_id)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	}
	if exists {
		return fmt.Errorf("asset already exists: %s", fail_work_order_id)
	}

	failmtewo := &FailureMaintenance{
		DocType:                      "failmtewo",
		FailWorkOrderID:              fail_work_order_id,
		FailWorkOrderType:            fail_work_order_type,
		FailFaultID:                  fault_id,
		FailWorkOrderDescription:     fail_work_order_description,
		FailGeneralHandSInstructions: fail_general_instructions,
		FailPlannedLabour:            fail_planned_labour,
		FailMaintenanceProcessSteps:  fail_maintenance_process_steps,
		FailSparePartsUsed:           fail_spare_parts_used,
	}

	// Setting work order creator
	failmtewo.SetFailWorkOrderCreator(ctx.GetStub())

	// Setting the failure maintenance work order timestamp
	failmtewo.SetFailWorkOrderTimestamp()

	// Setting the date of completion of the new work order adding 15 day, for example.
	failmtewo.SetFailWorkOrderCompletionDate()

	// Setting the condition record to not set.
	failmtewo.SetFailWorkOrderConditionNotGood()

	// Setting the validation record for the new work order to not validated.
	failmtewo.SetFailWorkOrderNotValidated()

	assetBytes, err := json.Marshal(failmtewo)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(fail_work_order_id, assetBytes)
	if err != nil {
		return err
	}

	//  Create an index to enable work order ID-type based range queries, e.g. return all assets with specific ID.
	//  An 'index' is a normal key-value entry in the ledger.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on indexNCorrWorkOrderID~WorkOrderType.
	//  This will enable very efficient state range queries based on composite keys matching indexNCorrWorkOrderID~*
	workOrderNameFailIndexKey, err := ctx.GetStub().CreateCompositeKey(indexFail, []string{failmtewo.FailWorkOrderID, failmtewo.FailWorkOrderType})
	if err != nil {
		return err
	}
	//  Save index entry to world state. Only the key name is needed, no need to store a duplicate copy of the asset.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	return ctx.GetStub().PutState(workOrderNameFailIndexKey, value)
}

//================================================================================================================================================
// READ FAILURE WORK ORDER
//================================================================================================================================================

// ReadFailureWorkOrder retrieves an asset from the ledger
func (t *MaintenanceSmartContract) ReadFailureWorkOrder(ctx contractapi.TransactionContextInterface, fail_work_order_id string) (*FailureMaintenance, error) {
	assetBytes, err := ctx.GetStub().GetState(fail_work_order_id)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset %s: %v", fail_work_order_id, err)
	}
	if assetBytes == nil {
		return nil, fmt.Errorf("asset %s does not exist", fail_work_order_id)
	}

	var failmtewo FailureMaintenance
	err = json.Unmarshal(assetBytes, &failmtewo)
	if err != nil {
		return nil, err
	}

	return &failmtewo, nil
}

//================================================================================================================================================
// DELETE FAILURE WORK ORDER
//================================================================================================================================================

// DeleteFailureWorkOrder removes an asset key-value pair from the ledger
func (t *MaintenanceSmartContract) DeleteFailureWorkOrder(ctx contractapi.TransactionContextInterface, fail_work_order_id string) error {
	failmtewo, err := t.ReadFailureWorkOrder(ctx, fail_work_order_id)
	if err != nil {
		return err
	}

	err = ctx.GetStub().DelState(fail_work_order_id)
	if err != nil {
		return fmt.Errorf("failed to delete failure work order %s: %v", fail_work_order_id, err)
	}

	workorderNameIndexKey, err := ctx.GetStub().CreateCompositeKey(indexFail, []string{failmtewo.FailWorkOrderID, failmtewo.FailWorkOrderType})
	if err != nil {
		return err
	}

	// Delete index entry
	return ctx.GetStub().DelState(workorderNameIndexKey)
}

//================================================================================================================================================
// CONSTRUCT QUERY RESPONSE FROM ITERATOR
//================================================================================================================================================

//  constructFailureQueryResponseFromIterator a slice of assets from the resultsIterator
func constructFailureQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) ([]*FailureMaintenance, error) {
	var assetsFail []*FailureMaintenance
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var failmtewo FailureMaintenance
		err = json.Unmarshal(queryResult.Value, &failmtewo)
		if err != nil {
			return nil, err
		}
		assetsFail = append(assetsFail, &failmtewo)
	}

	return assetsFail, nil
}

//================================================================================================================================================
// GET FAILURE WORK ORDERS BY RANGE
//================================================================================================================================================

// GetFailureWorkOrdersByRange performs a range query based on the start and end keys provided.
// Read-only function results are not typically submitted to ordering. If the read-only
// results are submitted to ordering, or if the query is used in an update transaction
// and submitted to ordering, then the committing peers will re-execute to guarantee that
// result sets are stable between endorsement time and commit time. The transaction is
// invalidated by the committing peers if the result set has changed between endorsement
// time and commit time.
// Therefore, range queries are a safe option for performing update transactions based on query results.
func (t *MaintenanceSmartContract) GetFailureWorkOrdersByRange(ctx contractapi.TransactionContextInterface, startKey, endKey string) ([]*FailureMaintenance, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructFailureQueryResponseFromIterator(resultsIterator)
}

//================================================================================================================================================
// QUERY FAILURE WORK ORDER BY TYPE
//================================================================================================================================================

// QueryFAilureWorkOrderByType queries for assets based on the types name.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (type).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *MaintenanceSmartContract) QueryFailureWorkOrderByType(ctx contractapi.TransactionContextInterface, fail_work_order_type string) ([]*FailureMaintenance, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"failmtewo","fail_work_order_type":"%s"}}`, fail_work_order_type)
	return getFailureQueryResultForQueryString(ctx, queryString)
}

//================================================================================================================================================
// QUERY FAILURE WORK ORDER BY CONDITION
//================================================================================================================================================

// QueryFailureWorkOrderByCondition queries for assets based on their condition.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (condition).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *MaintenanceSmartContract) QueryFailureWorkOrderByCondition(ctx contractapi.TransactionContextInterface, fail_condition string) ([]*FailureMaintenance, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"failmtewo","fail_condition":"%s"}}`, fail_condition)
	return getFailureQueryResultForQueryString(ctx, queryString)
}

//================================================================================================================================================
// QUERY FAILURE WORK ORDER BY VALIDATION
//================================================================================================================================================

// QueryFailureWorkOrderByValidation queries for work orders based on their validation.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (validation).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *MaintenanceSmartContract) QueryFailureWorkOrderByValidation(ctx contractapi.TransactionContextInterface, fail_validation string) ([]*FailureMaintenance, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"failmtewo","fail_validation":"%s"}}`, fail_validation)
	return getFailureQueryResultForQueryString(ctx, queryString)
}

//================================================================================================================================================
// QUERY FAILURE WORK ORDER BY ID
//================================================================================================================================================

// QueryFailureWorkOrderByID queries for work orders based on the its ID number.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (type).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *MaintenanceSmartContract) QueryFailureWorkOrderByID(ctx contractapi.TransactionContextInterface, fail_work_order_id string) ([]*FailureMaintenance, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"failmtewo","fail_work_order_id":"%s"}}`, fail_work_order_id)
	return getFailureQueryResultForQueryString(ctx, queryString)
}

//================================================================================================================================================
// QUERY FAILURE WORK ORDER
//================================================================================================================================================

// QueryFailureWorkOrders uses a query string to perform a query for work orders.
// Query string matching state database syntax is passed in and executed as is.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the "QueryAssetsForOwner" example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Ad hoc rich query
func (t *MaintenanceSmartContract) QueryFailureWorkOrders(ctx contractapi.TransactionContextInterface, queryString string) ([]*FailureMaintenance, error) {
	return getQueryResultForQueryStringFailure(ctx, queryString)
}

//================================================================================================================================================
// QUERY FAILURE WORK ORDER BY FAULT ID
//================================================================================================================================================
func (t *MaintenanceSmartContract) QueryFailureWorkOrderByFaultID(ctx contractapi.TransactionContextInterface, fail_fault_id string) ([]*FailureMaintenance, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"failmtewo","fail_fault_id":"%s"}}`, fail_fault_id)
	return getFailureQueryResultForQueryString(ctx, queryString)
}

//================================================================================================================================================
// GET QUERY RESULT FOR QUERY STRING FAILURE WORK ORDERS
//================================================================================================================================================

// getQueryResultForQueryString executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryStringFailure(ctx contractapi.TransactionContextInterface, queryString string) ([]*FailureMaintenance, error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructFailureQueryResponseFromIterator(resultsIterator)
}

//================================================================================================================================================
// QUERY FAILURE WORK ORDER
//================================================================================================================================================

// QueryFailureWorkOrders uses a query string to perform a query for work orders.
// Query string matching resultsIterator
// getFailureFailureQueryResultForQueryString executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getFailureQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]*FailureMaintenance, error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructFailureQueryResponseFromIterator(resultsIterator)
}

//================================================================================================================================================
// GET FAILURE WORK ORDERS BY RANGE WITH PAGINATION
//================================================================================================================================================

// GetFailureWorkOrdersByRangeWithPagination performs a range query based on the start and end key,
// page size and a bookmark.
// The number of fetched records will be equal to or lesser than the page size.
// Paginated range queries are only valid for read only transactions.
// Example: Pagination with Range Query
func (t *MaintenanceSmartContract) GetFailureWorkOrdersByRangeWithPagination(ctx contractapi.TransactionContextInterface, startKey string, endKey string, pageSize int, bookmark string) (*FailurePaginatedQueryResult, error) {

	resultsIterator, responseMetadata, err := ctx.GetStub().GetStateByRangeWithPagination(startKey, endKey, int32(pageSize), bookmark)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	assetsFail, err := constructFailureQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	return &FailurePaginatedQueryResult{
		Records:             assetsFail,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil
}

//================================================================================================================================================
// QUERY FAILURE WORK ORDERS WITH PAGINATION
//================================================================================================================================================

// QueryFailureWorkOrdersWithPagination uses a query string, page size and a bookmark to perform a query
// for assets. Query string matching state database syntax is passed in and executed as is.
// The number of fetched records would be equal to or lesser than the specified page size.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the QueryAssetsForOwner example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// Paginated queries are only valid for read only transactions.
// Example: Pagination with Ad hoc Rich Query
func (t *MaintenanceSmartContract) QueryFailureWorkOrdersWithPagination(ctx contractapi.TransactionContextInterface, queryString string, pageSize int, bookmark string) (*FailurePaginatedQueryResult, error) {

	return getFailureQueryResultForQueryStringWithPagination(ctx, queryString, int32(pageSize), bookmark)
}

//================================================================================================================================================
// GET QUERY RESULT FOR QUERY STRING WITH PAGINATION
//================================================================================================================================================

// getFailureQueryResultForQueryStringWithPagination executes the passed in query string with
// pagination info. The result set is built and returned as a byte array containing the JSON results.
func getFailureQueryResultForQueryStringWithPagination(ctx contractapi.TransactionContextInterface, queryString string, pageSize int32, bookmark string) (*FailurePaginatedQueryResult, error) {

	resultsIterator, responseMetadata, err := ctx.GetStub().GetQueryResultWithPagination(queryString, pageSize, bookmark)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	assetsFail, err := constructFailureQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	return &FailurePaginatedQueryResult{
		Records:             assetsFail,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil
}

//================================================================================================================================================
// GET FAILURE WORK ORDER HISTORY
//================================================================================================================================================

// GetFailureWorkOrderHistory returns the chain of custody for an asset since issuance.
func (t *MaintenanceSmartContract) GetFailureWorkOrderHistory(ctx contractapi.TransactionContextInterface, fail_work_order_id string) ([]FailureHistoryQueryResult, error) {
	log.Printf("GetFailureWorkOrderHistory: ID %v", fail_work_order_id)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(fail_work_order_id)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []FailureHistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var failmtewo FailureMaintenance
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &failmtewo)
			if err != nil {
				return nil, err
			}
		} else {
			failmtewo = FailureMaintenance{
				FailWorkOrderID: fail_work_order_id,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)

		if err != nil {
			return nil, err
		}

		record := FailureHistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: timestamp,
			Record:    &failmtewo,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}

//================================================================================================================================================
// DOES FAILURE WORK ORDER EXISTS
//================================================================================================================================================

// DoesFailureWorkOrderExists returns true when asset with given ID exists in the ledger.
func (t *MaintenanceSmartContract) DoesFailureWorkOrderExists(ctx contractapi.TransactionContextInterface, fail_work_order_id string) (bool, error) {
	assetBytes, err := ctx.GetStub().GetState(fail_work_order_id)
	if err != nil {
		return false, fmt.Errorf("failed to read asset %s from world state. %v", fail_work_order_id, err)
	}

	return assetBytes != nil, nil
}

//================================================================================================================================================
// UPDATE FAILURE WORK ORDER
//================================================================================================================================================

// UpdateFailureWorkOrder updates an existing work order in the world state with provided parameters.
func (t *MaintenanceSmartContract) UpdateFailureWorkOrder(ctx contractapi.TransactionContextInterface, fail_work_order_id string, fail_work_order_type string, fault_id string, fail_work_order_description string, fail_general_instructions string, fail_planned_labour PlannedLabour, fail_maintenance_process_steps MaintenanceProcessSteps, fail_spare_parts_used MaintenanceParts) error {
	exists, err := t.DoesFailureWorkOrderExists(ctx, fail_work_order_id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", fail_work_order_id)
	}

	// overwriting original asset with new asset
	failmtewo := FailureMaintenance{
		DocType:                      "failmtewo",
		FailWorkOrderID:              fail_work_order_id,
		FailWorkOrderType:            fail_work_order_type,
		FailFaultID:                  fault_id,
		FailWorkOrderDescription:     fail_work_order_description,
		FailGeneralHandSInstructions: fail_general_instructions,
		FailPlannedLabour:            fail_planned_labour,
		FailMaintenanceProcessSteps:  fail_maintenance_process_steps,
		FailSparePartsUsed:           fail_spare_parts_used,
	}

	// Setting work order creator
	failmtewo.SetFailWorkOrderCreator(ctx.GetStub())

	// Setting the timestamp of the new work order to the current time.
	failmtewo.SetFailWorkOrderTimestamp()

	// Setting the condition record to "good".
	failmtewo.SetFailWorkOrderConditionGood()

	// Setting the validation record for the new work order to "validated".
	failmtewo.SetFailWorkOrderValidated()

	assetJSON, err := json.Marshal(failmtewo)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(fail_work_order_id, assetJSON)
}
