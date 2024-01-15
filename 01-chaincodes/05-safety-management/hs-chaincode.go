/*
 SPDX-License-Identifier: Apache-2.0
*/

/*
====CHAINCODE EXECUTION SAMPLES (CLI) ==================

==== Invoke assets ====
peer chaincode invoke -C myc1 -n asset_transfer -c '{"Args":["CreateAsset","asset1","blue","5","tom","35"]}'
peer chaincode invoke -C myc1 -n asset_transfer -c '{"Args":["CreateAsset","asset2","red","4","tom","50"]}'
peer chaincode invoke -C myc1 -n asset_transfer -c '{"Args":["CreateAsset","asset3","blue","6","tom","70"]}'
peer chaincode invoke -C myc1 -n asset_transfer -c '{"Args":["TransferAsset","asset2","jerry"]}'
peer chaincode invoke -C myc1 -n asset_transfer -c '{"Args":["TransferAssetByColor","blue","jerry"]}'
peer chaincode invoke -C myc1 -n asset_transfer -c '{"Args":["DeleteAsset","asset1"]}'

==== Query assets ====
peer chaincode query -C myc1 -n asset_transfer -c '{"Args":["ReadAsset","asset1"]}'
peer chaincode query -C myc1 -n asset_transfer -c '{"Args":["GetAssetsByRange","asset1","asset3"]}'
peer chaincode query -C myc1 -n asset_transfer -c '{"Args":["GetRiskAssessmentHistory","asset1"]}'

Rich Query (Only supported if CouchDB is used as state database):
peer chaincode query -C myc1 -n asset_transfer -c '{"Args":["QueryAssetsByOwner","tom"]}'
peer chaincode query -C myc1 -n asset_transfer -c '{"Args":["QueryAssets","{\"selector\":{\"owner\":\"tom\"}}"]}'

Rich Query with Pagination (Only supported if CouchDB is used as state database):
peer chaincode query -C myc1 -n asset_transfer -c '{"Args":["QueryAssetsWithPagination","{\"selector\":{\"owner\":\"tom\"}}","3",""]}'

INDEXES TO SUPPORT COUCHDB RICH QUERIES

Indexes in CouchDB are required in order to make JSON queries efficient and are required for
any JSON query with a sort. Indexes may be packaged alongside
chaincode in a META-INF/statedb/couchdb/indexes directory. Each index must be defined in its own
text file with extension *.json with the index definition formatted in JSON following the
CouchDB index JSON syntax as documented at:
http://docs.couchdb.org/en/2.3.1/api/database/find.html#db-index

This asset transfer ledger example chaincode demonstrates a packaged
index which you can find in META-INF/statedb/couchdb/indexes/indexOwner.json.

If you have access to the your peer's CouchDB state database in a development environment,
you may want to iteratively test various indexes in support of your chaincode queries.  You
can use the CouchDB Fauxton interface or a command line curl utility to create and update
indexes. Then once you finalize an index, include the index definition alongside your
chaincode in the META-INF/statedb/couchdb/indexes directory, for packaging and deployment
to managed environments.

In the examples below you can find index definitions that support asset transfer ledger
chaincode queries, along with the syntax that you can use in development environments
to create the indexes in the CouchDB Fauxton interface or a curl command line utility.


Index for docType, owner.

Example curl command line to define index in the CouchDB channel_chaincode database
curl -i -X POST -H "Content-Type: application/json" -d "{\"index\":{\"fields\":[\"docType\",\"owner\"]},\"name\":\"indexOwner\",\"ddoc\":\"indexOwnerDoc\",\"type\":\"json\"}" http://hostname:port/myc1_assets/_index


Index for docType, owner, size (descending order).

Example curl command line to define index in the CouchDB channel_chaincode database:
curl -i -X POST -H "Content-Type: application/json" -d "{\"index\":{\"fields\":[{\"size\":\"desc\"},{\"docType\":\"desc\"},{\"owner\":\"desc\"}]},\"ddoc\":\"indexSizeSortDoc\", \"name\":\"indexSizeSortDesc\",\"type\":\"json\"}" http://hostname:port/myc1_assets/_index

Rich Query with index design doc and index name specified (Only supported if CouchDB is used as state database):
peer chaincode query -C myc1 -n asset_transfer -c '{"Args":["QueryAssets","{\"selector\":{\"docType\":\"asset\",\"owner\":\"tom\"}, \"use_index\":[\"_design/indexOwnerDoc\", \"indexOwner\"]}"]}'

Rich Query with index design doc specified only (Only supported if CouchDB is used as state database):
peer chaincode query -C myc1 -n asset_transfer -c '{"Args":["QueryAssets","{\"selector\":{\"docType\":{\"$eq\":\"asset\"},\"owner\":{\"$eq\":\"tom\"},\"size\":{\"$gt\":0}},\"fields\":[\"docType\",\"owner\",\"size\"],\"sort\":[{\"size\":\"desc\"}],\"use_index\":\"_design/indexSizeSortDoc\"}"]}'
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const index = "activity~ID"
const incoincident = "incident~ID"

// Health$SafetyChaincode implements the fabric-contract-api-go programming model
type HealthandSafetyChaincode struct {
	contractapi.Contract
}

// HistoryQueryResult structure used for returning result of history query
type HistoryQueryResult struct {
	Record    *RiskAssessment `json:"record"`
	TxId      string          `json:"txId"`
	Timestamp time.Time       `json:"timestamp"`
	IsDelete  bool            `json:"isDelete"`
}

// PaginatedQueryResult structure used for returning paginated query results and metadata
type PaginatedQueryResult struct {
	Records             []*RiskAssessment `json:"records"`
	FetchedRecordsCount int32             `json:"fetchedRecordsCount"`
	Bookmark            string            `json:"bookmark"`
}

// HistoryQueryResult structure used for returning result of history query
type HistoryQueryResultAccidentIncidentReport struct {
	Record    *AccidentIncidentReport `json:"record"`
	TxId      string                  `json:"txId"`
	Timestamp time.Time               `json:"timestamp_accident_incident_report"`
	IsDelete  bool                    `json:"isDelete"`
}

// PaginatedQueryResult structure used for returning paginated query results and metadata
type PaginatedQueryResultAccidentIncidentReport struct {
	Records             []*AccidentIncidentReport `json:"records"`
	FetchedRecordsCount int32                     `json:"fetchedRecordsCount"`
	Bookmark            string                    `json:"bookmark"`
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////
//                                     ////////////////////////////////////////////////////////////
//          RISK ASSESSMENTS           ////////////////////////////////////////////////////////////
//                                     ////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

/*
// HistoryQueryResult structure used for returning result of history query
type HistoryQueryResult struct {
	Record    *RiskAssessment `json:"record"`
	TxId      string          `json:"txId"`
	Timestamp time.Time       `json:"timestamp"`
	IsDelete  bool            `json:"isDelete"`
}

// PaginatedQueryResult structure used for returning paginated query results and metadata
type PaginatedQueryResult struct {
	Records             []*RiskAssessment `json:"records"`
	FetchedRecordsCount int32             `json:"fetchedRecordsCount"`
	Bookmark            string            `json:"bookmark"`
}

// Setting the risk assessment ID
func (ba *RiskAssessment) SetRiskAssessmentID() error {
	// Generates a random ID of 8 bytes
	idBytes := make([]byte, 5)
	_, err := rand.Read(idBytes)
	if err != nil {
		return err
	}

	// Set the risk assessment ID
	risk_assessment_id := "RISK-ASSESSMENT-ID-" + hex.EncodeToString(idBytes)
	ba.RiskAssessmentID = risk_assessment_id

	return nil
}

// Set the assessment date
func (ba *RiskAssessment) SetRiskAssessmentDate() {
	ba.RiskAssessmentDate = time.Now().Format("2006-01-02")
}

// Set the the next risk assessment review date
func (ba *RiskAssessment) SetRiskAssessmentDateNextReview() {
	ba.RiskAssessmentDateNextReview = time.Now().AddDate(0, 6, 0).Format("2006-01-02")
}

// SetRiskAssessmentActivityToNotSet set the condition of the work order to mark as good
func (ba *RiskAssessment) SetRiskAssessmentActivityToNotSet() {
	ba.RiskAssessmentActivity = "not-set"
}

// SetRiskAssessmentCreatedBy function pickup the user ID from the certificate and set it as the creator of the risk assessment
func (ba *RiskAssessment) SetRiskAssessmentCreatedBy(ctx contractapi.TransactionContextInterface) error {
	// Get the creator of the risk assessment
	creator, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return err
	}

	// Set the creator of the risk assessment, not as the certificate, but as the user ID
	ba.RiskAssessmentCreatedBy = creator

	return nil
}

// SetRiskAssessmentAssessedByToNotSet set the condition of the work order to mark as not-good
func (ba *RiskAssessment) SetRiskAssessmentAssessedByToNotSet() {
	ba.RiskAssessmentAssessedBy = "not-set"
}

// SetRiskAssessmentApprovedByToNotSet set the condition of the work order to mark as good
func (ba *RiskAssessment) SetRiskAssessmentApprovedByToNotSet() {
	ba.RiskAssessmentApprovedBy = "not-set"
}
*/

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// CreateRiskAssessment initializes a new risk assessment as an digital asset
// in the ledger, also we need to create a new risk assessment ID, which is a string, and then the
// risk assessment itself, which is a struct. The struct RiskAssessment is defined in utils-hs.go.
// The risk assessment ID is a string that is a concatenation of the string "risk-assessment-" and a
// hexadecimal string that is generated by the system. The risk assessment ID is the key to the risk
// assessment struct.
// The risk assessment struct is defined in utils-hs.go.
// The risk assessment struct is then marshalled into a byte array, which is then stored in the ledger.
// The risk assessment struct is then unmarshalled from the byte array and returned to the caller.

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// CreateRiskAssessment initializes a new risk assessment as a digital asset
// in the ledger, and generates a new risk assessment ID using the SetRiskAssessmentID function.
func (t *HealthandSafetyChaincode) CreateRiskAssessment(ctx contractapi.TransactionContextInterface,
	docType string,
	risk_assessment_id string,
	risk_assessment_date string,
	risk_assessment_date_next_review string,
	risk_assessment_activity string,
	risk_assessment_created_by string,
	risk_assessment_assessed_by string,
	risk_assessment_approved_by string,
	risk_assessment_hazard_list_one string,
	risk_assessment_hazard_list_two string,
	risk_assessment_hazard_list_three string,
	risk_assessment_hazard_list_four string,
	risk_assessment_hazard_list_five string) error {

	// Check if the risk assessment already exists
	exists, err := t.RiskAssessmentExists(ctx, risk_assessment_id)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("the risk assessment %s already exists", risk_assessment_id)
	}

	// Generate a new RiskAssessment struct
	riskassessment := &RiskAssessment{
		DocType:                       "Risk-Assessment",
		RiskAssessmentID:              risk_assessment_id,
		RiskAssessmentDate:            risk_assessment_date,
		RiskAssessmentDateNextReview:  risk_assessment_date_next_review,
		RiskAssessmentActivity:        risk_assessment_activity,
		RiskAssessmentCreatedBy:       risk_assessment_created_by,
		RiskAssessmentAssessedBy:      risk_assessment_assessed_by,
		RiskAssessmentApprovedBy:      risk_assessment_approved_by,
		RiskAssessmentHazardListOne:   risk_assessment_hazard_list_one,
		RiskAssessmentHazardListTwo:   risk_assessment_hazard_list_two,
		RiskAssessmentHazardListThree: risk_assessment_hazard_list_three,
		RiskAssessmentHazardListFour:  risk_assessment_hazard_list_four,
		RiskAssessmentHazardListFive:  risk_assessment_hazard_list_five,
	}

	// Marshal the RiskAssessment struct to JSON and store it in the ledger
	assetBytes, err := json.Marshal(riskassessment)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(riskassessment.RiskAssessmentID, assetBytes)
	if err != nil {
		return err
	}

	// Create an index to enable range queries based on risk assessment activity and assessed by fields
	activityIndexKey, err := ctx.GetStub().CreateCompositeKey(index, []string{riskassessment.RiskAssessmentID, riskassessment.RiskAssessmentActivity})
	if err != nil {
		return err
	}

	// Save the index entry to the ledger
	value := []byte{0x00}
	return ctx.GetStub().PutState(activityIndexKey, value)
}

// Following the function CreateRiskAssessment, we need to create a function that will read
// the all risk assessment from the ledger. This function will be used to query the ledger
// for all risk assessment. The function will return a slice of pointers to RiskAssessment structs.
// The function will return an error if it encounters an error while reading the ledger.
// ReadRiskAssessment will read the risk assessment from the ledger.
// The function will return a pointer to a RiskAssessment struct.
// The function will return an error if it encounters an error while reading the ledger.
// The function will return an error if the risk assessment does not exist.
// ReadRiskAssessment will read the risk assessment from the ledger.
// The function will return a pointer to a RiskAssessment struct.

func (t *HealthandSafetyChaincode) ReadAllRiskAssessments(ctx contractapi.TransactionContextInterface) ([]*RiskAssessment, error) {
	// Create an empty collection of type RiskAssessment
	var riskassessments []*RiskAssessment

	// Create a key value pair iterator
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}

	// Close the iterator when the function returns
	defer resultsIterator.Close()

	// Iterate through the iterator
	for resultsIterator.HasNext() {
		// Get the next key value pair
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		// Create an empty RiskAssessment struct
		var riskassessment RiskAssessment

		// Unmarshal the JSON byte string to the RiskAssessment struct
		err = json.Unmarshal(queryResponse.Value, &riskassessment)
		if err != nil {
			return nil, err
		}

		// Append the RiskAssessment struct to the collection
		riskassessments = append(riskassessments, &riskassessment)
	}

	// Return the collection of RiskAssessment structs to the calling function
	return riskassessments, nil
}

// RiskAssessmentExist returns true when asset with given ID exists in the ledger.
func (t *HealthandSafetyChaincode) RiskAssessmentExists(ctx contractapi.TransactionContextInterface, risk_assessment_id string) (bool, error) {
	assetBytes, err := ctx.GetStub().GetState(risk_assessment_id)
	if err != nil {
		return false, fmt.Errorf("failed to read asset %s from world state. %v", risk_assessment_id, err)
	}

	return assetBytes != nil, nil
}

// ReadRiskAssessment retrieves a risk assessment from the ledger - AI generated function and it is fucking working!!!
func (t *HealthandSafetyChaincode) ReadRiskAssessment(ctx contractapi.TransactionContextInterface,
	risk_assessment_id string) (*RiskAssessment, error) {
	riskassessmentBytes, err := ctx.GetStub().GetState(risk_assessment_id)
	if err != nil {
		return nil, fmt.Errorf("failed to read risk assessment %s from world state. %v", risk_assessment_id, err)
	}
	if riskassessmentBytes == nil {
		return nil, fmt.Errorf("risk assessment %s does not exist", risk_assessment_id)
	}

	var riskassessment RiskAssessment
	err = json.Unmarshal(riskassessmentBytes, &riskassessment)
	if err != nil {
		return nil, err
	}

	return &riskassessment, nil
}

// DeleteRiskAssessment removes an asset key-value pair from the ledger
func (t *HealthandSafetyChaincode) DeleteRiskAssessment(ctx contractapi.TransactionContextInterface,
	risk_assessment_id string) error {
	riskassessment, err := t.ReadRiskAssessment(ctx, risk_assessment_id)
	if err != nil {
		return err
	}

	err = ctx.GetStub().DelState(risk_assessment_id)
	if err != nil {
		return fmt.Errorf("failed to delete asset %s: %v", risk_assessment_id, err)
	}

	activityIndexKey, err := ctx.GetStub().CreateCompositeKey(index, []string{riskassessment.RiskAssessmentActivity, riskassessment.RiskAssessmentID})
	if err != nil {
		return err
	}

	// Delete index entry
	return ctx.GetStub().DelState(activityIndexKey)
}

// constructQueryResponseFromIterator constructs a slice of assets from the resultsIterator
func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) ([]*RiskAssessment, error) {
	var assets []*RiskAssessment
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset RiskAssessment
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

// GetRiskAssessmentByRange queries for assets based on a passed in activity, performs a range query based
// on the start and end keys provided.
// Read-only function results are not typically submitted to ordering. If the read-only
// results are submitted to ordering, or if the query is used in an update transaction
// and submitted to ordering, then the committing peers will re-execute to guarantee that
// result sets are stable between endorsement time and commit time. The transaction is
// invalidated by the committing peers if the result set has changed between endorsement
// time and commit time.
// Therefore, range queries are a safe option for performing update transactions based on query results.
func (t *HealthandSafetyChaincode) GetRiskAssessmentByRange(ctx contractapi.TransactionContextInterface, startKey, endKey string) ([]*RiskAssessment, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructQueryResponseFromIterator(resultsIterator)
}

// WORKING!!!!
// Update the risk assessment with corresponding "risk_assessment_id" and with the new value of "risk_assessment_activity"
// UpdateRiskAssessmentActivity will update the risk assessment activity of a given risk assessment id.
// Uses GetStateByPartialCompositeKey (range query) against activity~ID 'index'.
// Committing peers will re-execute range queries to guarantee that result sets are stable
// between endorsement time and commit time. The transaction is invalidated by the
// committing peers if the result set has changed between endorsement time and commit time.
// Therefore, range queries are a safe option for performing update transactions based on query results.
// Example: GetStateByPartialCompositeKey/RangeQuery
func (t *HealthandSafetyChaincode) UpdateRiskAssessmentActivity(ctx contractapi.TransactionContextInterface, risk_assessment_id, updateActivity string) error {
	riskassessment, err := t.ReadRiskAssessment(ctx, risk_assessment_id)
	if err != nil {
		return err
	}

	riskassessment.RiskAssessmentActivity = updateActivity
	assetBytes, err := json.Marshal(riskassessment)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(risk_assessment_id, assetBytes)
}

// Update the corresponding risk assessment with "risk"assessment_id" with the new value of "risk_assessment_assessed_by"
// UpdateRiskAssessmentAssessedBy will updated the risk assessment assessed by of a given risk assessment id.
// Uses GetStateByPartialCompositeKey (range query) against activity~ID 'index'.
// Committing peers will re-execute range queries to guarantee that result sets are stable
// between endorsement time and commit time. The transaction is invalidated by the
// committing peers if the result set has changed between endorsement time and commit time.
// Therefore, range queries are a safe option for performing update transactions based on query results.
// Example: GetStateByPartialCompositeKey/RangeQuery
func (t *HealthandSafetyChaincode) UpdateRiskAssessmentAssessedBy(ctx contractapi.TransactionContextInterface, risk_assessment_id, updateAssessedBy string) error {
	riskassessment, err := t.ReadRiskAssessment(ctx, risk_assessment_id)
	if err != nil {
		return err
	}

	riskassessment.RiskAssessmentAssessedBy = updateAssessedBy
	assetBytes, err := json.Marshal(riskassessment)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(risk_assessment_id, assetBytes)
}

// Update the corresponding risk assessment with "risk_assessment_id" with the new value of "risk_assessment_approved_by"
// UpdateRiskAssessmentApprovedBy will updated the risk assessment approved by of a given risk assessment id.
// Uses GetStateByPartialCompositeKey (range query) against activity~ID 'index'.
// Committing peers will re-execute range queries to guarantee that result sets are stable
// between endorsement time and commit time. The transaction is invalidated by the
// committing peers if the result set has changed between endorsement time and commit time.
// Therefore, range queries are a safe option for performing update transactions based on query results.
// Example: GetStateByPartialCompositeKey/RangeQuery.
// Add function that only identity with the role of "General Manager" can approve a risk assessment.
func (t *HealthandSafetyChaincode) UpdateRiskAssessmentApprovedBy(ctx contractapi.TransactionContextInterface, risk_assessment_id, updateApprovedBy string) error {

	// Get the client identity
	clientIdentity, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return fmt.Errorf("failed to get client identity: %v", err)
	}

	// Check if the client identity is the General Manager
	// "General Manager ID" => "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg=="
	if clientIdentity != "eDUwOTo6Q049R2VuZXJhbCBNYW5hZ2VyLE9VPWNsaWVudDo6Q049ZmFicmljLWNhLXNlcnZlcg==" {
		return fmt.Errorf("access denied. Only the General Manager can execute this function")
	}

	riskassessment, err := t.ReadRiskAssessment(ctx, risk_assessment_id)
	if err != nil {
		return err
	}

	riskassessment.RiskAssessmentApprovedBy = updateApprovedBy
	assetBytes, err := json.Marshal(riskassessment)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(risk_assessment_id, assetBytes)
}

// TO BE TESTED!!!!
// QueryRiskAssessmentsByActivity queries for assets based on the owners name.
// This is an example of risk-assessment query where the query logic is baked into the chaincode,
// and accepting a single query parameter (owner).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *HealthandSafetyChaincode) QueryRiskAssessmentsByActivity(ctx contractapi.TransactionContextInterface, risk_assessment_activity string) ([]*RiskAssessment, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"Risk-Assessment","risk_assessment_activity":"%s"}}`, risk_assessment_activity)
	return getQueryResultForQueryString(ctx, queryString)
}

// QueryRiskAssessmentsByAssessor queries for assets based on the owners name.
// This is an example of risk-assessment query where the query logic is baked into the chaincode,
// and accepting a single query parameter (owner).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *HealthandSafetyChaincode) QueryRiskAssessmentsByAssessor(ctx contractapi.TransactionContextInterface, risk_assessment_assessed_by string) ([]*RiskAssessment, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"Risk-Assessment","risk_assessment_assessed_by":"%s"}}`, risk_assessment_assessed_by)
	return getQueryResultForQueryString(ctx, queryString)
}

// QueryRiskAssessmentsByAssessor queries for assets based on the owners name.
// This is an example of risk-assessment query where the query logic is baked into the chaincode,
// and accepting a single query parameter (owner).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (t *HealthandSafetyChaincode) QueryRiskAssessmentsByApprover(ctx contractapi.TransactionContextInterface, risk_assessment_approved_by string) ([]*RiskAssessment, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"Risk-Assessment","risk_assessment_approved_by":"%s"}}`, risk_assessment_approved_by)
	return getQueryResultForQueryString(ctx, queryString)
}

// QueryRiskAssessments uses a query string to perform a query for assets.
// Query string matching state database syntax is passed in and executed as is.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the QueryRiskAssessmentsForOwner example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Ad hoc rich query
func (t *HealthandSafetyChaincode) QueryRiskAssessments(ctx contractapi.TransactionContextInterface, queryString string) ([]*RiskAssessment, error) {
	return getQueryResultForQueryString(ctx, queryString)
}

// getQueryResultForQueryString executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]*RiskAssessment, error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructQueryResponseFromIterator(resultsIterator)
}

// GetAssetsByRangeWithPagination performs a range query based on the start and end key,
// page size and a bookmark.
// The number of fetched records will be equal to or lesser than the page size.
// Paginated range queries are only valid for read only transactions.
// Example: Pagination with Range Query
func (t *HealthandSafetyChaincode) GetAssetsByRangeWithPagination(ctx contractapi.TransactionContextInterface, startKey string, endKey string, pageSize int, bookmark string) ([]*RiskAssessment, error) {

	resultsIterator, _, err := ctx.GetStub().GetStateByRangeWithPagination(startKey, endKey, int32(pageSize), bookmark)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructQueryResponseFromIterator(resultsIterator)
}

// QueryAssetsWithPagination uses a query string, page size and a bookmark to perform a query
// for assets. Query string matching state database syntax is passed in and executed as is.
// The number of fetched records would be equal to or lesser than the specified page size.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the QueryAssetsForOwner example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// Paginated queries are only valid for read only transactions.
// Example: Pagination with Ad hoc Rich Query
func (t *HealthandSafetyChaincode) QueryAssetsWithPagination(ctx contractapi.TransactionContextInterface, queryString string, pageSize int, bookmark string) (*PaginatedQueryResult, error) {

	return getQueryResultForQueryStringWithPagination(ctx, queryString, int32(pageSize), bookmark)
}

// getQueryResultForQueryStringWithPagination executes the passed in query string with
// pagination info. The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryStringWithPagination(ctx contractapi.TransactionContextInterface, queryString string, pageSize int32, bookmark string) (*PaginatedQueryResult, error) {

	resultsIterator, responseMetadata, err := ctx.GetStub().GetQueryResultWithPagination(queryString, pageSize, bookmark)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	assets, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	return &PaginatedQueryResult{
		Records:             assets,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil
}

// GetRiskAssessmentHistory returns the chain of custody for an asset since issuance. WORKING!!!
func (t *HealthandSafetyChaincode) GetRiskAssessmentHistory(ctx contractapi.TransactionContextInterface, risk_assessment_id string) ([]HistoryQueryResult, error) {
	log.Printf("GetRiskAssessmentHistory: ID %v", risk_assessment_id)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(risk_assessment_id)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset RiskAssessment
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &asset)
			if err != nil {
				return nil, err
			}
		} else {
			asset = RiskAssessment{
				RiskAssessmentID: risk_assessment_id,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			return nil, err
		}

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: timestamp,
			Record:    &asset,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////                                           //////////////////////////////////////////////////////
///////////      INCIDENT & ACCIDENT REPORTING        //////////////////////////////////////////////////////
///////////                                           //////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Create an Incident Report following the structure defined in the schema
// Example: CreateAccidentIncidentReport - create an incident report
// CreateAccidentIncidentReport initializes a new report as an digital asset
// in the ledger, and generates a new accident/incident report ID using the SetAccidentIncidentReportID function.
func (t *HealthandSafetyChaincode) CreateAccidentIncidentReport(ctx contractapi.TransactionContextInterface,
	docType string,
	accident_incident_report_id string,
	subject_accident_incident_report string,
	accident_incident_report_date string,
	accident_incident_report_time_start string,
	accident_incident_report_time_end string,
	accident_incident_report_location string,
	accident_incident_report_hs_aspects string,
	accident_incident_report_classification string,
	accident_incident_report_description string,
	accident_incident_report_immediate_action string,
	accident_incident_report_follow_up_actions string,
	accident_incident_report_created_by string,
	accident_incident_report_validated_by string,
	accident_incident_report_status string,
	accident_incident_report_number_of_people_involved int,
	accident_incident_report_person_injured string,
	accident_incident_report_witnesses string,
	accident_incident_report_event_overview string,
	accident_incident_report_line_of_communication string) error {

	// Generate a new AccidentIncidentReport struct
	accidentincidentreport := &AccidentIncidentReport{
		DocType:                                      "accident-incident-report",
		AccidentIncidentReportID:                     accident_incident_report_id,
		SubjectAccidentIncidentReport:                subject_accident_incident_report,
		AccidentIncidentReportDate:                   accident_incident_report_date,
		AccidentIncidentReportTimeStart:              accident_incident_report_time_start,
		AccidentIncidentReportTimeEnd:                accident_incident_report_time_end,
		AccidentIncidentReportLocation:               accident_incident_report_location,
		AccidentIncidentReportHSAspects:              accident_incident_report_hs_aspects,
		AccidentIncidentReportClassification:         accident_incident_report_classification,
		AccidentIncidentReportDescription:            accident_incident_report_description,
		AccidentIncidentReportImmediateAction:        accident_incident_report_immediate_action,
		AccidentIncidentReportFollowUpActions:        accident_incident_report_follow_up_actions,
		AccidentIncidentReportCreatedBy:              accident_incident_report_created_by,
		AccidentIncidentReportValidatedBy:            accident_incident_report_validated_by,
		AccidentIncidentReportStatus:                 accident_incident_report_status,
		AccidentIncidentReportNumberOfPeopleInvolved: accident_incident_report_number_of_people_involved,
		AccidentIncidentReportPersonInjured:          accident_incident_report_person_injured,
		AccidentIncidentReportWitnesses:              accident_incident_report_witnesses,
		AccidentIncidentReportEventOverview:          accident_incident_report_event_overview,
		AccidentIncidentReportLineOfCommunication:    accident_incident_report_line_of_communication,
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// Marshal the RiskAssessment struct to JSON and store it in the ledger
	assetBytes, err := json.Marshal(accidentincidentreport)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(accidentincidentreport.AccidentIncidentReportID, assetBytes)
	if err != nil {
		return err
	}

	// Create an index to enable range queries based on risk assessment activity and assessed by fields
	incidentAccidentIndexKey, err := ctx.GetStub().CreateCompositeKey(incoincident, []string{accidentincidentreport.AccidentIncidentReportID, accidentincidentreport.SubjectAccidentIncidentReport})
	if err != nil {
		return err
	}

	// Save the index entry to the ledger
	value := []byte{0x00}
	return ctx.GetStub().PutState(incidentAccidentIndexKey, value)
}

// Read an Accident/Incident Report from the ledger function
// Example: ReadAccidentIncidentReport - read an accident/incident report
// ReadAccidentIncidentReport retrieves the accident/incident report from the ledger
// based on its ID
func (t *HealthandSafetyChaincode) ReadAccidentIncidentReport(ctx contractapi.TransactionContextInterface, accident_incident_report_id string) (*AccidentIncidentReport, error) {
	// Read the Accident/Incident Report from the ledger
	assetJSON, err := ctx.GetStub().GetState(accident_incident_report_id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the accident/incident report %s does not exist", accident_incident_report_id)
	}

	var accidentincidentreport AccidentIncidentReport
	err = json.Unmarshal(assetJSON, &accidentincidentreport)
	if err != nil {
		return nil, err
	}

	return &accidentincidentreport, nil
}

// Update an Accident/Incident Report function with corresponding ID.
// Example: UpdateAccidentIncidentReportTimeEnd - update an accident/incident report
// UpdateAccidentIncidentReportTimeEnd updates the accident/incident report time end
// based on the accident/incident report ID
func (t *HealthandSafetyChaincode) UpdateAccidentIncidentReportTimeEnd(ctx contractapi.TransactionContextInterface, accident_incident_report_id, updateTimeEnd string) error {
	accidentincidentreport, err := t.ReadAccidentIncidentReport(ctx, accident_incident_report_id)
	if err != nil {
		return err
	}

	// Update the accident/incident report time end
	accidentincidentreport.AccidentIncidentReportTimeEnd = updateTimeEnd

	// Marshal the Accident/Incident Report struct to JSON
	assetJSON, err := json.Marshal(accidentincidentreport)
	if err != nil {
		return err
	}

	// Write the Accident/Incident Report to the ledger
	return ctx.GetStub().PutState(accident_incident_report_id, assetJSON)
}

// Update an Accident/Incident Report function with corresponding ID.
// Example: UpdateAccidentIncidentReportFollowUpActions - update an accident/incident report
// UpdateAccidentIncidentReportFollowUpActions updates the accident/incident report follow up actions
// based on the accident/incident report ID
func (t *HealthandSafetyChaincode) UpdateAccidentIncidentReportFollowUpActions(ctx contractapi.TransactionContextInterface, accident_incident_report_id string, updateFollowUpActions string) error {
	accidentincidentreport, err := t.ReadAccidentIncidentReport(ctx, accident_incident_report_id)
	if err != nil {
		return err
	}

	// Update the accident/incident report follow up actions
	accidentincidentreport.AccidentIncidentReportFollowUpActions = updateFollowUpActions

	// Marshal the Accident/Incident Report struct to JSON
	assetJSON, err := json.Marshal(accidentincidentreport)
	if err != nil {
		return err
	}

	// Write the Accident/Incident Report to the ledger
	return ctx.GetStub().PutState(accident_incident_report_id, assetJSON)

}

// Update an Accident/Incident Report function with corresponding ID.
// Example: UpdateAccidentIncidentReportValidated - update an accident/incident report
// UpdateAccidentIncidentReportValidated updates the accident/incident report validated
// based on the accident/incident report ID and at the same time, it will updated the accident/incident report status to closed.
// the argument updateValidated is a string that will be passed shall be the ID name of the person who validated the accident/incident report
// and executed the transaction/function.
func (t *HealthandSafetyChaincode) UpdateAccidentIncidentReportValidated(ctx contractapi.TransactionContextInterface, accident_incident_report_id string) error {
	accidentincidentreport, err := t.ReadAccidentIncidentReport(ctx, accident_incident_report_id)
	if err != nil {
		return err
	}

	// Get the identity name of the person who validated the accident/incident report
	// and executed the transaction/function
	identityName, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return err
	}

	// convert the identity name to string, that is showing the name of the identity
	identityNameString := string(identityName)

	// Pass the identityNameString to the updateValidated argument
	//updateValidated = identityNameString

	// Update the accident/incident report validated
	accidentincidentreport.AccidentIncidentReportValidatedBy = identityNameString

	// Update the accident/incident report status to closed
	accidentincidentreport.AccidentIncidentReportStatus = "closed"

	// Marshal the Accident/Incident Report struct to JSON
	assetJSON, err := json.Marshal(accidentincidentreport)
	if err != nil {
		return err
	}

	// Write the Accident/Incident Report to the ledger
	return ctx.GetStub().PutState(accident_incident_report_id, assetJSON)

}

// Update an Accident/Incident Report function with corresponding ID.
// Example: UpdateAccidentIncidentReportNumberOfPeopleInvolved - update an accident/incident report
// UpdateAccidentIncidentReportNumberOfPeopleInvolved updates the accident/incident report number of people involved
// based on the accident/incident report ID.
// the argument updateNumberOfPeopleInvolved is an integer that will be passed shall be the number of people involved in the accident/incident.
func (t *HealthandSafetyChaincode) UpdateAccidentIncidentReportNumberOfPeopleInvolved(ctx contractapi.TransactionContextInterface, accident_incident_report_id string, updateNumberOfPeopleInvolved int) error {
	accidentincidentreport, err := t.ReadAccidentIncidentReport(ctx, accident_incident_report_id)
	if err != nil {
		return err
	}

	// Update the accident/incident report number of people involved
	accidentincidentreport.AccidentIncidentReportNumberOfPeopleInvolved = updateNumberOfPeopleInvolved

	// Marshal the Accident/Incident Report struct to JSON
	assetJSON, err := json.Marshal(accidentincidentreport)
	if err != nil {
		return err
	}

	// Write the Accident/Incident Report to the ledger
	return ctx.GetStub().PutState(accident_incident_report_id, assetJSON)

}

// Update an Accident/Incident Report function with corresponding ID.
// Example: UpdateAccidentIncidentReportPersonInjured - update an accident/incident report
// UpdateAccidentIncidentReportPersonInjured updates the accident/incident report person injured
// based on the accident/incident report ID.
// the argument updatePersonInjured is a "[]string" that will be passed shall be the name of the person injured in the accident/incident.
func (t *HealthandSafetyChaincode) UpdateAccidentIncidentReportPersonInjured(ctx contractapi.TransactionContextInterface, accident_incident_report_id string, updatePersonInjured string) error {
	accidentincidentreport, err := t.ReadAccidentIncidentReport(ctx, accident_incident_report_id)
	if err != nil {
		return err
	}

	// Update the accident/incident report person injured
	accidentincidentreport.AccidentIncidentReportPersonInjured = updatePersonInjured

	// Marshal the Accident/Incident Report struct to JSON
	assetJSON, err := json.Marshal(accidentincidentreport)
	if err != nil {
		return err
	}

	// Write the Accident/Incident Report to the ledger
	return ctx.GetStub().PutState(accident_incident_report_id, assetJSON)

}

// Update an Accident/Incident Report function with corresponding ID.
// Example: UpdateAccidentIncidentReportWitnesses - update an accident/incident report
// UpdateAccidentIncidentReportWitnesses updates the accident/incident report witnesses
// based on the accident/incident report ID.
// the argument updateWitnesses is a "[]string" that will be passed shall be the name of the witnesses in the accident/incident.
func (t *HealthandSafetyChaincode) UpdateAccidentIncidentReportWitnesses(ctx contractapi.TransactionContextInterface, accident_incident_report_id string, updateWitnesses string) error {
	accidentincidentreport, err := t.ReadAccidentIncidentReport(ctx, accident_incident_report_id)
	if err != nil {
		return err
	}

	// Update the accident/incident report witnesses
	accidentincidentreport.AccidentIncidentReportWitnesses = updateWitnesses

	// Marshal the Accident/Incident Report struct to JSON
	assetJSON, err := json.Marshal(accidentincidentreport)
	if err != nil {
		return err
	}

	// Write the Accident/Incident Report to the ledger
	return ctx.GetStub().PutState(accident_incident_report_id, assetJSON)

}

// Update an Accident/Incident Report function with corresponding ID.
// Example: UpdateAccidentIncidentReportEventOverview - update an accident/incident report
// UpdateAccidentIncidentReportEventOverview updates the accident/incident report event overview
// based on the accident/incident report ID.
// the argument updateEventOverview is a "[]string" that will be passed shall be the event overview of the accident/incident.
func (t *HealthandSafetyChaincode) UpdateAccidentIncidentReportEventOverview(ctx contractapi.TransactionContextInterface, accident_incident_report_id string, updateEventOverview string) error {
	accidentincidentreport, err := t.ReadAccidentIncidentReport(ctx, accident_incident_report_id)
	if err != nil {
		return err
	}

	// Update the accident/incident report event overview
	accidentincidentreport.AccidentIncidentReportEventOverview = updateEventOverview

	// Marshal the Accident/Incident Report struct to JSON
	assetJSON, err := json.Marshal(accidentincidentreport)
	if err != nil {
		return err
	}

	// Write the Accident/Incident Report to the ledger
	return ctx.GetStub().PutState(accident_incident_report_id, assetJSON)

}

// Update an Accident/Incident Report function with corresponding ID.
// Example: UpdateAccidentIncidentReportLineOfCommunication - update an accident/incident report.
// UpdateAccidentIncidentReportLineOfCommunication updates the accident/incident report line of communication
// based on the accident/incident report ID.
// the argument updateLineOfCommunication is a "[]string" that will be passed shall be the line of communication of the accident/incident.
func (t *HealthandSafetyChaincode) UpdateAccidentIncidentReportLineOfCommunication(ctx contractapi.TransactionContextInterface, accident_incident_report_id string, updateLineOfCommunication string) error {
	accidentincidentreport, err := t.ReadAccidentIncidentReport(ctx, accident_incident_report_id)
	if err != nil {
		return err
	}

	// Update the accident/incident report line of communication
	accidentincidentreport.AccidentIncidentReportLineOfCommunication = updateLineOfCommunication

	// Marshal the Accident/Incident Report struct to JSON
	assetJSON, err := json.Marshal(accidentincidentreport)
	if err != nil {
		return err
	}

	// Write the Accident/Incident Report to the ledger
	return ctx.GetStub().PutState(accident_incident_report_id, assetJSON)

}

// QueryIncidentReports uses a query string to perform a query for assets.
// Query string matching state database syntax is passed in and executed as is.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the QueryIncidentReports example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Ad hoc rich query
func (t *HealthandSafetyChaincode) QueryIncidentReports(ctx contractapi.TransactionContextInterface, queryString string) ([]*AccidentIncidentReport, error) {
	return getQueryResultForQueryStringAccidentIncidentReport(ctx, queryString)
}

// getQueryResultForQueryString executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryStringAccidentIncidentReport(ctx contractapi.TransactionContextInterface, queryString string) ([]*AccidentIncidentReport, error) {

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	accidentincidentreports := []*AccidentIncidentReport{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		accidentincidentreport := new(AccidentIncidentReport)
		err = json.Unmarshal(queryResponse.Value, accidentincidentreport)
		if err != nil {
			return nil, err
		}
		accidentincidentreports = append(accidentincidentreports, accidentincidentreport)
	}

	return accidentincidentreports, nil
}

// Get the history for an Accident/Incident Report function with the corresponding ID.
// Example: GetAccidentIncidentReportHistory - get the history for an accident/incident report.
// GetAccidentIncidentReportHistory gets the accident/incident report history
// based on the accident/incident report ID.
func (t *HealthandSafetyChaincode) GetAccidentIncidentReportHistory(ctx contractapi.TransactionContextInterface, accident_incident_report_id string) ([]*HistoryQueryResultAccidentIncidentReport, error) {
	log.Printf("Get Accident/Incident Report History: ID: " + accident_incident_report_id)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(accident_incident_report_id)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []*HistoryQueryResultAccidentIncidentReport
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var accidentincidentreport AccidentIncidentReport
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &accidentincidentreport)
			if err != nil {
				return nil, err
			}
		} else {
			accidentincidentreport = AccidentIncidentReport{
				AccidentIncidentReportID: accident_incident_report_id,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			return nil, err
		}

		record := &HistoryQueryResultAccidentIncidentReport{
			TxId:      response.TxId,
			Timestamp: timestamp,
			Record:    &accidentincidentreport,
			IsDelete:  response.IsDelete,
		}

		records = append(records, record)
	}

	return records, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	chaincode, err := contractapi.NewChaincode(&HealthandSafetyChaincode{})
	if err != nil {
		log.Panicf("Error creating asset chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting asset chaincode: %v", err)
	}

	// Expose metrics endpoint for Prometheus to scrape
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":8080", nil)
}
