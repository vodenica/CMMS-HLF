/*
 * Copyright 2018 IBM All Rights Reserved.
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
	"time"

	"github.com/golang/protobuf/ptypes"
	//"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	//"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
)

const indexPr = "type~status"
const indexInv = "type~status"

// FinSmartContract contract for handling the business logic of a basic financial services of an enterprise
type FinSmartContract struct {
	contractapi.Contract
}

// PurchaseRequestHistoryQueryResult structure used for returning result of history query
type PurchaseRequestHistoryQueryResult struct {
	Record    *PurchaseRequest `json:"record"`
	TxId      string           `json:"txId"`
	Timestamp time.Time        `json:"timestamp"`
	IsDelete  bool             `json:"isDelete"`
}

// InvoiceOpsHistoryQueryResult structure used for returning result of history query
type InvoiceOpsHistoryQueryResult struct {
	Record    *InvoiceOps `json:"record"`
	TxId      string      `json:"txId"`
	Timestamp time.Time   `json:"timestamp"`
	IsDelete  bool        `json:"isDelete"`
}

// InvoiceWorkHistoryQueryResult structure used for returning result of history query
type InvoiceWorkHistoryQueryResult struct {
	Record    *InvoiceWork `json:"record"`
	TxId      string       `json:"txId"`
	Timestamp time.Time    `json:"timestamp"`
	IsDelete  bool         `json:"isDelete"`
}

//==========================================================================================================================
// NEW PURCHASE REQUEST - WORKING
//==========================================================================================================================
// NewPurchaseRequestMtncParts adds a new basic object to the world state using id as key
func (c *FinSmartContract) NewPurchaseRequestMtncParts(ctx CustomTransactionContextInterface,
	docType string,
	purchase_request_type string,
	//purchase_order_id string,
	purchase_request_item_one RequestItemOne,
	purchase_request_item_two RequestItemTwo,
	purchase_request_item_three RequestItemThree,
	purchase_request_item_four RequestItemFour,
	purchase_request_item_five RequestItemFive,
	supplier_order_id string,
	supplier_order_description string,
	purchase_request_remarks string) error {

	var pr PurchaseRequest

	existing := ctx.GetCallData()

	if existing != nil {
		return fmt.Errorf("cannot create new object in world state as key %s already exists", pr.PurchaseOrderID)
	}

	// Working with identities.
	requestor, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(requestor) == "User1" || string(requestor) == "Maintenance Supervisor" {
		return fmt.Errorf("you are not authorized to create a new purchase order")
	}

	/*
		pr := PurchaseRequest{}
	*/

	pr.DocType = docType
	pr.PurchaseRequestType = purchase_request_type

	//pr.PurchaseOrderID = purchase_order_id

	pr.PurchaseRequestItemOne = purchase_request_item_one
	pr.PurchaseRequestItemTwo = purchase_request_item_two
	pr.PurchaseRequestItemThree = purchase_request_item_three
	pr.PurchaseRequestItemFour = purchase_request_item_four
	pr.PurchaseRequestItemFive = purchase_request_item_five

	pr.SupplierOrderID = supplier_order_id
	pr.SupplierOrderDescription = supplier_order_description

	// pr.SignOffReceivedGoods = sign_off_received_goods
	pr.PurchaseRequestRemarks = purchase_request_remarks

	// Setting the Purchase Order ID
	pr.SetPurchaseOrderID()

	pr.SetPurchaseOrderCreationDate()

	pr.SetPurchaseOrderCreatorIdentity(ctx.GetStub())

	pr.SetPurchaseOrderApproverIdentityToNotSet()

	// Setting total cost of the purchase request:
	newTotalCostOfPurchaseOrder := purchase_request_item_one.ItemOneTotalCost + purchase_request_item_two.ItemTwoTotalCost + purchase_request_item_three.ItemThreeTotalCost + purchase_request_item_four.ItemFourTotalCost + purchase_request_item_five.ItemFiveTotalCost

	pr.PurchaseOrderTotalCost = newTotalCostOfPurchaseOrder

	pr.SetPurchaseRequestNonUrgent()
	pr.SetPurchaseRequestOpen()

	pr.SetVerifyDeliveryToNotDelivered()
	pr.SetVerifyQuantityNotDelivered()
	pr.SetVerifyConditionOpen()

	// The following function shall be updated with the corresponding function to set the sign off received goods
	pr.SetSignOffReceivedGoodsToNotSet()

	// Check if the purchase order exists on the world state
	exists, err := c.DoesPurchaseOrderExists(ctx, pr.PurchaseOrderID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the purchase order %s already exists", pr.PurchaseOrderID)
	}

	baBytes, _ := json.Marshal(pr)

	err2 := ctx.GetStub().PutState(pr.PurchaseOrderID, []byte(baBytes))

	if err2 != nil {
		return errors.New("unable to interact with world state")
	}

	//  Create an index to enable color-based range queries, e.g. return all blue assets.
	//  An 'index' is a normal key-value entry in the ledger.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on indexName~color~name.
	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
	statusPurchaseRequestIndexKey, err := ctx.GetStub().CreateCompositeKey(indexPr, []string{pr.DocType, pr.PurchaseRequestType})
	if err != nil {
		return err
	}
	//  Save index entry to world state. Only the key name is needed, no need to store a duplicate copy of the asset.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	return ctx.GetStub().PutState(statusPurchaseRequestIndexKey, value)

	// return nil
}

//==========================================================================================================================
// UPDATE PURCHASE ORDER TO APPROVED - WORKING
//==========================================================================================================================
// UpdateValue changes the value of a basic object to add the value passed
// Adding number 1, the status will change into completed
func (c *FinSmartContract) UpdatePurchaseOrderToApproved(ctx CustomTransactionContextInterface,
	purchase_order_id string,
	updateOrderApprover string,
	valueApprove string) error {
	existing := ctx.GetCallData()

	if existing == nil {
		return fmt.Errorf("cannot update object in world state as key %s does not exist", purchase_order_id)
	}

	// Working with identities!!!

	approver, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(approver) == "User1" {
		return fmt.Errorf("you are not authorized to approve a new purchase order")
	}

	pr := PurchaseRequest{}

	err := json.Unmarshal(existing, &pr)

	if err != nil {
		return fmt.Errorf("data retrieved from world state for key %s was not of type BasicObject", purchase_order_id)
	}

	pr.PurchaseOrderApprover = updateOrderApprover
	pr.StatusOfTheOrder = valueApprove

	baBytes, _ := json.Marshal(pr)

	err = ctx.GetStub().PutState(purchase_order_id, []byte(baBytes))

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	return nil
}

//==========================================================================================================================
// UPDATE PURCHASE ORDER TO DECLINED
//==========================================================================================================================
// UpdatePurchaseOrderToDecline  changes the value of a basic object to add the value passed
// Adding number 1, the status will change into completed
func (c *FinSmartContract) UpdatePurchaseOrderToDecline(ctx CustomTransactionContextInterface, purchase_order_id string, updateOrderApprover string, valueDecline string) error {
	existing := ctx.GetCallData()

	if existing == nil {
		return fmt.Errorf("cannot update object in world state as key %s does not exist", purchase_order_id)
	}

	// Working with identities!!!

	approver, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(approver) == "User1" {
		return fmt.Errorf("you are not authorized to approve a new purchase order")
	}

	pr := PurchaseRequest{}

	err := json.Unmarshal(existing, &pr)

	if err != nil {
		return fmt.Errorf("data retrieved from world state for key %s was not of type BasicObject", purchase_order_id)
	}

	pr.PurchaseOrderApprover = updateOrderApprover
	pr.StatusOfTheOrder = valueDecline

	baBytes, _ := json.Marshal(pr)

	err = ctx.GetStub().PutState(purchase_order_id, []byte(baBytes))

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	return nil
}

//==================================================================================================================
// GET THE PURCHASE REQUEST BY ID - WORKING
//==================================================================================================================
// GetPurchaseRequestByID returns the object with id given from the world state
func (c *FinSmartContract) GetPurchaseRequestByID(ctx CustomTransactionContextInterface, purchase_order_id string) (*PurchaseRequest, error) {
	existing := ctx.GetCallData()

	if existing == nil {
		return nil, fmt.Errorf("cannot read world state pair with key %s. Does not exist", purchase_order_id)
	}

	// Working with identities!!!
	reviewer, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return nil, fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(reviewer) == "User1" {
		return nil, fmt.Errorf("you are not authorized to approve a new purchase order")
	}

	pr := new(PurchaseRequest)

	err2 := json.Unmarshal(existing, pr)

	if err2 != nil {
		return nil, fmt.Errorf("data retrieved from world state for key %s was not of type BasicObject", purchase_order_id)
	}

	return pr, nil
}

//==========================================================================================================================
// DOES WORK ORDER EXISTS - WORKING
//==========================================================================================================================
// DoesPurchaseOrderExists returns true when asset with given ID exists in the ledger.
func (c *FinSmartContract) DoesPurchaseOrderExists(ctx CustomTransactionContextInterface, purchase_order_id string) (bool, error) {

	// Working with identities!!!

	reviewer, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return false, fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(reviewer) == "User1" {
		return false, fmt.Errorf("you are not authorized to approve a new purchase order")
	}

	assetBytes, err := ctx.GetStub().GetState(purchase_order_id)
	if err != nil {
		return false, fmt.Errorf("failed to read asset %s from world state. %v", purchase_order_id, err)
	}

	return assetBytes != nil, nil
}

//==========================================================================================================================
// FUNCTION CONSTRUCT QUERY RESPONSE FROM ITERATOR [QUERY]
//==========================================================================================================================

// constructQueryResponseFromIterator constructs a slice of assets from the resultsIterator
func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) ([]*PurchaseRequest, error) {

	var assetspurchaserequest []*PurchaseRequest
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var assetpurchaserequest PurchaseRequest
		err = json.Unmarshal(queryResult.Value, &assetpurchaserequest)
		if err != nil {
			return nil, err
		}
		assetspurchaserequest = append(assetspurchaserequest, &assetpurchaserequest)
	}

	return assetspurchaserequest, nil
}

//=========================================================================================================================
// UPDATED THE SIGN-OFF  OF THE DELIVERED GOODS
//=========================================================================================================================
func (c *FinSmartContract) UpdateSignOffReceivedGoods(ctx CustomTransactionContextInterface, purchase_order_id string, updateSignOffReceivedGoods string) error {
	existing := ctx.GetCallData()

	if existing == nil {
		return fmt.Errorf("cannot update object in world state as key %s does not exist", purchase_order_id)
	}

	// Working with identities!!!

	approver, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(approver) == "User1" {
		return fmt.Errorf("you are not authorized to approve a new purchase order")
	}

	pr := PurchaseRequest{}

	err := json.Unmarshal(existing, &pr)

	if err != nil {
		return fmt.Errorf("data retrieved from world state for key %s was not of type BasicObject", purchase_order_id)
	}

	// Optional, we can set up the identity which runs the function
	pr.SignOffReceivedGoods = updateSignOffReceivedGoods

	baBytes, _ := json.Marshal(pr)

	err = ctx.GetStub().PutState(purchase_order_id, []byte(baBytes))

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	return nil
}

//=========================================================================================================================
// FUNCTION CONSTRUCT QUERY RESPONSE FROM ITERATOR (BY STATUS) [QUERY]
//=========================================================================================================================

// constructQueryResponseFromIterator constructs a slice of assets from the resultsIterator
func constructQueryResponseFromIteratorByTheStatus(resultsIterator shim.StateQueryIteratorInterface) ([]*PurchaseRequest, error) {
	var assetspurchaserequestbythestatus []*PurchaseRequest
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var assetbythestatus PurchaseRequest
		err = json.Unmarshal(queryResult.Value, &assetbythestatus)
		if err != nil {
			return nil, err
		}
		assetspurchaserequestbythestatus = append(assetspurchaserequestbythestatus, &assetbythestatus)
	}

	return assetspurchaserequestbythestatus, nil
}

//=========================================================================================================================
// GET PURCHASE REQUEST HISTORY - WORKING
//=========================================================================================================================

// GetPurchaseRequestHistory returns the chain of custody for an asset since issuance.
func (c *FinSmartContract) GetPurchaseRequestHistory(ctx contractapi.TransactionContextInterface, purchase_order_id string) ([]PurchaseRequestHistoryQueryResult, error) {

	// Working with identities!!!
	reviewer, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return nil, fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(reviewer) == "User1" {
		return nil, fmt.Errorf("you are not authorized to approve a new purchase order")
	}

	log.Printf("GetPurchaseOrderHistory: ID %v", purchase_order_id)

	resultsIteratorPO, err := ctx.GetStub().GetHistoryForKey(purchase_order_id)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorPO.Close()

	var records []PurchaseRequestHistoryQueryResult
	for resultsIteratorPO.HasNext() {
		response, err := resultsIteratorPO.Next()
		if err != nil {
			return nil, err
		}

		var purchaserequest PurchaseRequest
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &purchaserequest)
			if err != nil {
				return nil, err
			}
		} else {
			purchaserequest = PurchaseRequest{
				PurchaseOrderID: purchase_order_id,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)

		if err != nil {
			return nil, err
		}

		record := PurchaseRequestHistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: timestamp,
			Record:    &purchaserequest,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}

//=========================================================================================================================
// QUERY THE PURCHASE REQUEST [QUERY]
//=========================================================================================================================

// QueryPurchaseRequest uses a query string to perform a query for purchase orders.
// Query string matching state database syntax is passed in and executed as is.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the QueryPurchaseRequestForStatus example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
func (c *FinSmartContract) QueryPurchaseRequest(ctx contractapi.TransactionContextInterface, queryStringPurchaseRequest string) ([]*PurchaseRequest, error) {

	// Working with identities!!!
	reviewer, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return nil, fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(reviewer) == "User1" {
		return nil, fmt.Errorf("you are not authorized to approve a new purchase order")
	}

	return getQueryResultForQueryString(ctx, queryStringPurchaseRequest)
}

//=========================================================================================================================
// FUNCTION GET QUERY RESULTS FOR QUERY STRING [QUERY]
//=========================================================================================================================

// getQueryResultForQueryString executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]*PurchaseRequest, error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructQueryResponseFromIterator(resultsIterator)
}

//=========================================================================================================================
// GET PURCHASE REQUESTS BY RANGE [QUERY] - WORKING
// NOTE:Some of the POs' were change the filed "status_of_the_order" and unmarshal does not working.
//=========================================================================================================================

// GetPurchaseRequestsByRange performs a range query based on the start and end keys provided.
// Read-only function results are not typically submitted to ordering. If the read-only
// results are submitted to ordering, or if the query is used in an update transaction
// and submitted to ordering, then the committing peers will re-execute to guarantee that
// result sets are stable between endorsement time and commit time. The transaction is
// invalidated by the committing peers if the result set has changed between endorsement
// time and commit time.
// Therefore, range queries are a safe option for performing update transactions based on query results.
func (c *FinSmartContract) GetPurchaseRequestsByRange(ctx contractapi.TransactionContextInterface, startKey, endKey string) ([]*PurchaseRequest, error) {

	// Working with identities!!!
	reviewer, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return nil, fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(reviewer) == "User1" {
		return nil, fmt.Errorf("you are not authorized to approve a new purchase order")
	}

	resultsIteratorPurchaseRequest, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorPurchaseRequest.Close()

	return constructQueryResponseFromIterator(resultsIteratorPurchaseRequest)
}

//=========================================================================================================================
// QUERY PURCHASE REQUESTS BY THE STATUS [QUERY] - WORKING
//=========================================================================================================================

// QueryAssetsByTheStatusOfThePurchaseRequest queries for assets based on the owners name.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (owner).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (c *FinSmartContract) QueryAssetsByTheStatusOfThePurchaseRequest(ctx contractapi.TransactionContextInterface, status_of_the_order string) ([]*PurchaseRequest, error) {
	queryStringByTheStatus := fmt.Sprintf(`{"selector":{"docType":"purchase-order","status_of_the_order":"%s"}}`, status_of_the_order)
	return getQueryResultForQueryStringQueryByTheStatus(ctx, queryStringByTheStatus)
}

//=========================================================================================================================
// FUNCTION GET QUERY RESULTS BY STATUS [QUERY]
//=========================================================================================================================

// getQueryResultForQueryStringQueryByTheStatus executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryStringQueryByTheStatus(ctx contractapi.TransactionContextInterface, queryStringByTheStatus string) ([]*PurchaseRequest, error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryStringByTheStatus)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructQueryResponseFromIteratorByTheStatus(resultsIterator)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////                           ///////////////////////////////////////////////////////////////////
/////////////////////////////       I N V O C E S       ///////////////////////////////////////////////////////////////////
/////////////////////////////                           ///////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//=========================================================================================================================
// ISSUING THE NEW OPERATIONAL SERVICES INVOICE
//=========================================================================================================================
// NewInvoiceOperationalServices adds a new basic object to the world state using id as key
func (c *FinSmartContract) NewInvoiceOperationalServices(ctx CustomTransactionContextInterface,
	docType string,
	invoice_type_ops string,
	//invoice_id string,
	invoice_operational_cost int,
	variation_additional_expenses int,
	downtime_deduction int,
	cpi_adjustment int,
	other_agreed_reductions int) error {

	// Declaring the existing variable 'in' for the InvoiceOps struct
	var in InvoiceOps

	existing := ctx.GetCallData()

	// If function to check if the created Invoices with specific ID exists in the world state.
	// If exists, then return error message.
	if existing != nil {
		return fmt.Errorf("cannot create new object in world state as key %s already exists", in.InvoiceOpsID)
	}

	// Working with identities.
	requestor, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(requestor) == "User1" || string(requestor) == "General Manager" {
		return fmt.Errorf("you are not authorized to create a new purchase order")
	}

	// Setting the docType of the invoice "invoice"
	in.DocType = docType

	// Setting the Invoice type, such as "operations", "additional work", "maintenance", etc.
	in.InvoiceTypeOps = invoice_type_ops

	// Setting the running ID number of the invoice, e.g INV-OPS-20230920-0001
	//in.InvoiceID = invoice_id
	in.SetInvoiceOpsID()

	// Setting Company Registration Number
	in.SetCompanyRegistrationNumber()

	// Setting the date when the invoice was created
	in.SetInvoiceOpsDateCreated()

	// Setting payment terms
	in.SetPaymentTerms()

	// Setting the billing period of the invoice for the previous month
	in.SetBillingPeriodOpsServices()
	//in.BillingPeriodOpsServices = billing_period_ops_services // Dates => from - to

	// Setting Invoice description
	in.SetInvoiceDescription()

	// Setting Invoice Cost Code
	in.SetInvoiceCostCode()

	in.InvoiceOperationalCost = invoice_operational_cost
	in.VariationAdditionalExpenses = variation_additional_expenses
	in.DowntimeDeduction = downtime_deduction
	in.CPIAdjustment = cpi_adjustment
	in.OtherAgreedReductions = other_agreed_reductions

	// This part calculates the total amount of the Operational Services Invoice.
	newTotalServiceInvoiceCost := invoice_operational_cost + variation_additional_expenses - downtime_deduction + cpi_adjustment - other_agreed_reductions

	if newTotalServiceInvoiceCost < 0 {
		newTotalServiceInvoiceCost = 0
	}

	// Setting the total amount of the Operational Services Invoice.
	in.TotalServiceInvoiceCost = uint(newTotalServiceInvoiceCost)

	in.SetInvoiceOpsApprovalStatusPending()
	in.SetInvoiceOpsStatusIssued()

	// Check for existing data in the world state
	existing, err := ctx.GetStub().GetState(in.InvoiceOpsID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if existing != nil {
		return fmt.Errorf("the asset %s already exists", in.InvoiceOpsID)
	}

	baBytes, _ := json.Marshal(in)

	err = ctx.GetStub().PutState(in.InvoiceOpsID, []byte(baBytes))

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	// Create an index to enable pending-base range queries, e.g. return all pending invoices
	// An 'index' is a normal key-value entry in the ledger.
	// The key is a composite key, with the elements that you want to range query on listed first.
	// In our case, the composite key is based on indexName~InvoiceID~status.
	// This will enable very efficient state range queries based on composite keys matching indexName~color~*
	statusInvoiceOpsIndexKey, err := ctx.GetStub().CreateCompositeKey(indexInv, []string{in.DocType, in.InvoiceTypeOps})
	if err != nil {
		return err
	}
	// Save index entry to world state. Only the key name is needed, no need to store a duplicate copy of the asset.
	// Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	return ctx.GetStub().PutState(statusInvoiceOpsIndexKey, value)

}

//=========================================================================================================================
// ISSUING THE NEW INVOICE FOR ADDITIONAL WORK
//=========================================================================================================================
// NewInvoiceAdditionalWork adds a new basic object to the world state using id as key
func (c *FinSmartContract) NewInvoiceAdditionalWork(ctx CustomTransactionContextInterface,
	docType string,
	invoice_type_work string,
	invoice_description string,
	invoice_work_cost_code int,
	variation_additional_expenses_add_work int,
	invoice_item_one InvoiceItemOne,
	invoice_item_two InvoiceItemTwo,
	invoice_item_three InvoiceItemThree,
	invoice_item_four InvoiceItemFour,
	invoice_item_five InvoiceItemFive) error {

	var iw InvoiceWork

	existing := ctx.GetCallData()

	if existing != nil {
		return fmt.Errorf("cannot create new object in world state as key %s already exists", iw.InvoiceWorkID)
	}

	// Working with identities.
	requestor, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(requestor) == "User1" || string(requestor) == "General Manager" {
		return fmt.Errorf("you are not authorized to create a new purchase order")
	}

	iw.DocType = docType                   // here we can put "invoice-work"
	iw.InvoiceTypeWork = invoice_type_work // here we can put "additional-work"

	iw.SetInvoiceWorkID() // e.g. INV-WORK-20230920-0001

	// Setting Company Registration Number
	iw.SetCompanyRegistrationNumber()
	//iw.CompanyRegistrationNumber = company_registration_number

	// Setting the date when the invoice was created
	iw.SetInvoiceWorkDateCreated()
	//iw.InvoiceDate = invoice_date

	// Setting payment terms of the invoice
	iw.SetPaymentTerms()
	//iw.PaymentTerms = payment_terms

	// Setting the billing period of the invoice
	iw.SetBillingPeriodForAdditionalWork()
	//iw.BillingPeriodAdditionalWork = billing_period_additional_work

	iw.InvoiceDescription = invoice_description

	iw.InvoiceWorkCostCode = invoice_work_cost_code

	iw.VariationAdditionalExpensesAddWork = variation_additional_expenses_add_work

	iw.InvoiceWorkItemOne = invoice_item_one // here we enter all items which needs to be listed in the invoice

	iw.InvoiceWorkItemTwo = invoice_item_two // here we enter all items which needs to be listed in the invoice

	iw.InvoiceWorkItemThree = invoice_item_three // here we enter all items which needs to be listed in the invoice

	iw.InvoiceWorkItemFour = invoice_item_four // here we enter all items which needs to be listed in the invoice

	iw.InvoiceWorkItemFive = invoice_item_five // here we enter all items which needs to be listed in the invoice

	// this part calculates the total amount of the Operational Services Invoice.
	newTotalAddWorkInvoiceCost := variation_additional_expenses_add_work + invoice_item_one.InvoiceItemOneCost + invoice_item_two.InvoiceItemTwoCost + invoice_item_three.InvoiceItemThreeCost + invoice_item_four.InvoiceItemFourCost + invoice_item_five.InvoiceItemFiveCost

	if newTotalAddWorkInvoiceCost < 0 {
		newTotalAddWorkInvoiceCost = 0
	}

	iw.TotalAddWorkInvoiceCost = uint(newTotalAddWorkInvoiceCost)

	iw.SetInvoiceWorkApprovalStatusPending()
	iw.SetInvoiceWorkStatusIssued()

	// Check existing data in the world state
	existing, err := ctx.GetStub().GetState(iw.InvoiceWorkID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if existing != nil {
		return fmt.Errorf("the asset %s already exists", iw.InvoiceWorkID)
	}

	baBytes, _ := json.Marshal(iw)

	err = ctx.GetStub().PutState(iw.InvoiceWorkID, []byte(baBytes))

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	//  Create an index to enable color-based range queries, e.g. return all blue assets.
	//  An 'index' is a normal key-value entry in the ledger.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on indexName~color~name.
	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
	statusInvoiceWorkIndexKey, err := ctx.GetStub().CreateCompositeKey(indexInv, []string{iw.DocType, iw.InvoiceTypeWork})
	if err != nil {
		return err
	}
	//  Save index entry to world state. Only the key name is needed, no need to store a duplicate copy of the asset.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	return ctx.GetStub().PutState(statusInvoiceWorkIndexKey, value)

}

//=========================================================================================================================
// UPDATE OPERATIONAL SERVICES INVOICE STATUS TO APPROVED
//=========================================================================================================================
// UpdateInvoiceOpsStatusToApproved changes the value of a basic object to add the value passed
func (c *FinSmartContract) UpdateInvoiceOpsStatusToApproved(ctx CustomTransactionContextInterface, invoice_ops_id string) error {
	existing := ctx.GetCallData()

	if existing == nil {
		return fmt.Errorf("cannot update object in world state as key %s does not exist", invoice_ops_id)
	}

	// Working with identities!!!
	approver, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(approver) == "General Manager" {
		return fmt.Errorf("you are not authorized to approve the new Invoice")
	}

	in := InvoiceOps{}

	err := json.Unmarshal(existing, &in)

	if err != nil {
		return fmt.Errorf("data retrieved from world state for key %s was not of type BasicObject", invoice_ops_id)
	}

	in.SetInvoiceOpsApprovalStatusApproved()

	baBytes, _ := json.Marshal(in)

	err = ctx.GetStub().PutState(invoice_ops_id, []byte(baBytes))

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	return nil
}

//=========================================================================================================================
// UPDATE ADDITIONAL WORK INVOICE STATUS TO APPROVED
//=========================================================================================================================
// UpdateInvoiceWorkStatusToApproved changes the value of a basic object to add the value passed
func (c *FinSmartContract) UpdateInvoiceWorkStatusToApproved(ctx CustomTransactionContextInterface, invoice_ops_id string) error {
	existing := ctx.GetCallData()

	if existing == nil {
		return fmt.Errorf("cannot update object in world state as key %s does not exist", invoice_ops_id)
	}

	// Working with identities!!!
	approver, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(approver) == "General Manager" {
		return fmt.Errorf("you are not authorized to approve the new Invoice")
	}

	iw := InvoiceWork{}

	err := json.Unmarshal(existing, &iw)

	if err != nil {
		return fmt.Errorf("data retrieved from world state for key %s was not of type BasicObject", invoice_ops_id)
	}

	iw.SetInvoiceWorkApprovalStatusApproved()

	baBytes, _ := json.Marshal(iw)

	err = ctx.GetStub().PutState(invoice_ops_id, []byte(baBytes))

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	return nil
}

//=========================================================================================================================
// UPDATE THE STATUS OF THE ADDITIONAL WORK INVOICE INTO DECLINED
//=========================================================================================================================
// UpdateInvoiceWorkStatusToDecline  changes the value of a basic object to add the value passed.
func (c *FinSmartContract) UpdateInvoiceWorkStatusToDecline(ctx CustomTransactionContextInterface, invoice_id string) error {
	existing := ctx.GetCallData()

	if existing == nil {
		return fmt.Errorf("cannot update object in world state as key %s does not exist", invoice_id)
	}

	// Working with identities!!!
	approver, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "General Manager" - as only General Manager cna decline/cancel the Invoice.
	if string(approver) == "General Manager" {
		return fmt.Errorf("you are not authorized to decline the Invoice")
	}

	iw := InvoiceWork{}

	err := json.Unmarshal(existing, &iw)

	if err != nil {
		return fmt.Errorf("data retrieved from world state for key %s was not of type BasicObject", invoice_id)
	}

	iw.SetInvoiceWorkApprovalStatusDeclined()

	baBytes, _ := json.Marshal(iw)

	err = ctx.GetStub().PutState(invoice_id, []byte(baBytes))

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	return nil
}

//=========================================================================================================================
// GET THE OPERATIONAL SERVICES INVOICE REQUEST BY ID
//=========================================================================================================================
// GetInvoiceOpsByID returns the object with id given from the world state
func (c *FinSmartContract) GetInvoiceOpsByID(ctx CustomTransactionContextInterface, invoice_id string) (*InvoiceOps, error) {
	existing := ctx.GetCallData()

	if existing == nil {
		return nil, fmt.Errorf("cannot read world state pair with key %s. Does not exist", invoice_id)
	}

	// Working with identities!!!
	reviewer, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return nil, fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(reviewer) == "General Manager" {
		return nil, fmt.Errorf("you are not authorized to approve a new purchase order")
	}

	in := new(InvoiceOps)

	err2 := json.Unmarshal(existing, in)

	if err2 != nil {
		return nil, fmt.Errorf("data retrieved from world state for key %s was not of type BasicObject", invoice_id)
	}

	return in, nil
}

//=========================================================================================================================
// GET THE ADDITIONAL WORK INVOICE REQUEST BY ID
//=========================================================================================================================
// GetInvoiceWorkByID returns the object with id given from the world state
func (c *FinSmartContract) GetInvoiceWorkByID(ctx CustomTransactionContextInterface, invoice_id string) (*InvoiceWork, error) {
	existing := ctx.GetCallData()

	if existing == nil {
		return nil, fmt.Errorf("cannot read world state pair with key %s. Does not exist", invoice_id)
	}

	// Working with identities!!!
	reviewer, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return nil, fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(reviewer) == "General Manager" {
		return nil, fmt.Errorf("you are not authorized to approve a new purchase order")
	}

	iw := new(InvoiceWork)

	err2 := json.Unmarshal(existing, iw)

	if err2 != nil {
		return nil, fmt.Errorf("data retrieved from world state for key %s was not of type BasicObject", invoice_id)
	}

	return iw, nil
}

//=========================================================================================================================
// DOES THE INVOICE EXISTS - WORKING!
//=========================================================================================================================
// DoesInvoiceExists returns true when asset with given ID exists in the ledger.
func (c *FinSmartContract) DoesInvoiceExists(ctx CustomTransactionContextInterface, invoice_id string) (bool, error) {

	// Working with identities!!!

	reviewer, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return false, fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(reviewer) == "User1" {
		return false, fmt.Errorf("you are not authorized to approve a new purchase order")
	}

	assetBytes, err := ctx.GetStub().GetState(invoice_id)
	if err != nil {
		return false, fmt.Errorf("failed to read asset %s from world state. %v", invoice_id, err)
	}

	return assetBytes != nil, nil
}

//=========================================================================================================================
//*************************************************************************************************************************
//**************************************      QUERYING THE OPERATIONAL SERVICES INVOICE      ******************************
//*************************************************************************************************************************
//=========================================================================================================================
//=========================================================================================================================
// GET OPERATIONAL SERVICES INVOICE HISTORY
//=========================================================================================================================

// GetInvoiceOpsRequestHistory returns the chain of custody for an asset since issuance.
func (c *FinSmartContract) GetInvoiceOpsRequestHistory(ctx contractapi.TransactionContextInterface, invoice_ops_id string) ([]InvoiceOpsHistoryQueryResult, error) {

	// Working with identities!!!
	reviewer, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return nil, fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(reviewer) == "User1" {
		return nil, fmt.Errorf("you are not authorized to approve a new purchase order")
	}

	log.Printf("GetOperationalServicesHistory: ID %v", invoice_ops_id)

	resultsIteratorInvOps, err := ctx.GetStub().GetHistoryForKey(invoice_ops_id)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorInvOps.Close()

	var records []InvoiceOpsHistoryQueryResult
	for resultsIteratorInvOps.HasNext() {
		response, err := resultsIteratorInvOps.Next()
		if err != nil {
			return nil, err
		}

		var invoiceops InvoiceOps
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &invoiceops)
			if err != nil {
				return nil, err
			}
		} else {
			invoiceops = InvoiceOps{
				InvoiceOpsID: invoice_ops_id,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)

		if err != nil {
			return nil, err
		}

		record := InvoiceOpsHistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: timestamp,
			Record:    &invoiceops,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}

// GetInvoiceWorkRequestHistory returns the chain of custody for an asset since issuance.
func (c *FinSmartContract) GetInvoiceWorkRequestHistory(ctx contractapi.TransactionContextInterface, invoice_work_id string) ([]InvoiceWorkHistoryQueryResult, error) {

	// Working with identities!!!
	reviewer, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return nil, fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	if string(reviewer) == "User1" {
		return nil, fmt.Errorf("you are not authorized to approve a new purchase order")
	}

	log.Printf("GetInvoiceWorkRequestHistory: ID %v", invoice_work_id)

	resultsIteratorInvWork, err := ctx.GetStub().GetHistoryForKey(invoice_work_id)
	if err != nil {
		return nil, err
	}

	defer resultsIteratorInvWork.Close()

	var records []InvoiceWorkHistoryQueryResult
	for resultsIteratorInvWork.HasNext() {
		response, err := resultsIteratorInvWork.Next()
		if err != nil {
			return nil, err
		}

		var invoicework InvoiceWork
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &invoicework)
			if err != nil {
				return nil, err
			}
		} else {
			invoicework = InvoiceWork{
				InvoiceWorkID: invoice_work_id,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)

		if err != nil {
			return nil, err
		}

		record := InvoiceWorkHistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: timestamp,
			Record:    &invoicework,
			IsDelete:  response.IsDelete,
		}

		records = append(records, record)
	}

	return records, nil
}

//=========================================================================================================================
//=========================================================================================================================
//=========================================================================================================================
// 											QUERYING FUNCTIONS FOR OPERATIONAL SERVICE INVOICES
//=========================================================================================================================
// FUNCTION CONSTRUCT QUERY RESPONSE FROM ITERATOR (BY STATUS) [QUERY]
//=========================================================================================================================

// constructQueryResponseFromIterator constructs a slice of assets from the resultsIterator
func constructQueryResponseFromIteratorByTheStatusInvoiceOps(resultsIterator shim.StateQueryIteratorInterface) ([]*InvoiceOps, error) {
	var assetsinvoiceopsbythestatus []*InvoiceOps
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var assetinvoiceopsbythestatus InvoiceOps
		err = json.Unmarshal(queryResult.Value, &assetinvoiceopsbythestatus)
		if err != nil {
			return nil, err
		}
		assetsinvoiceopsbythestatus = append(assetsinvoiceopsbythestatus, &assetinvoiceopsbythestatus)
	}

	return assetsinvoiceopsbythestatus, nil
}

//=========================================================================================================================
// QUERY THE INVOICE BY THE STATUS [QUERY] - WORKING
//=========================================================================================================================

// QueryAssetsByTheStatusOfTheInvoiceOps queries for assets based on the owners name.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (owner).
// Only available on state databases that support rich query (e.g. CouchDB)
// Example: Parameterized rich query
func (c *FinSmartContract) QueryAssetsByTheStatusOfTheInvoiceOps(ctx contractapi.TransactionContextInterface, invoice_status string) ([]*InvoiceOps, error) {
	queryStringByTheStatusInvoiceOps := fmt.Sprintf(`{"selector":{"docType":"invoice","invoice_status":"%s"}}`, invoice_status)
	return getQueryResultForQueryStringQueryByTheStatusInvoiceOps(ctx, queryStringByTheStatusInvoiceOps)
}

//=========================================================================================================================
// FUNCTION GET QUERY RESULTS BY STATUS [QUERY]
//=========================================================================================================================

// getQueryResultForQueryStringQueryByTheStatusInvoiceOps executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryStringQueryByTheStatusInvoiceOps(ctx contractapi.TransactionContextInterface, queryStringByTheStatusInvoiceOps string) ([]*InvoiceOps, error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryStringByTheStatusInvoiceOps)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructQueryResponseFromIteratorByTheStatusInvoiceOps(resultsIterator)
}

//=========================================================================================================================
//=========================================================================================================================
//=========================================================================================================================
// FUNCTION CONSTRUCT QUERY RESPONSE FROM ITERATOR [QUERY]
//=========================================================================================================================

// constructQueryResponseFromIterator constructs a slice of assets from the resultsIterator
func constructQueryResponseFromIteratorInvoiceOps(resultsIterator shim.StateQueryIteratorInterface) ([]*InvoiceOps, error) {

	var assetsinvoiceops []*InvoiceOps
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var assetinvoiceops InvoiceOps
		err = json.Unmarshal(queryResult.Value, &assetinvoiceops)
		if err != nil {
			return nil, err
		}
		assetsinvoiceops = append(assetsinvoiceops, &assetinvoiceops)
	}

	return assetsinvoiceops, nil
}

//=========================================================================================================================
// FUNCTION CONSTRUCT QUERY RESPONSE FROM ITERATOR [QUERY]
//=========================================================================================================================

// constructQueryResponseFromIterator constructs a slice of assets from the resultsIterator
func constructQueryResponseFromIteratorInvoiceWork(resultsIterator shim.StateQueryIteratorInterface) ([]*InvoiceWork, error) {

	var assetsinvoicework []*InvoiceWork
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var assetinvoicework InvoiceWork
		err = json.Unmarshal(queryResult.Value, &assetinvoicework)
		if err != nil {
			return nil, err
		}
		assetsinvoicework = append(assetsinvoicework, &assetinvoicework)
	}

	return assetsinvoicework, nil
}

//=========================================================================================================================
// GET OPERATIONAL SERVICE INVOICES BY RANGE [QUERY] - WORKING!
//=========================================================================================================================

// GetInvoiceOpsByRange performs a range query based on the start and end keys provided.
// Read-only function results are not typically submitted to ordering. If the read-only
// results are submitted to ordering, or if the query is used in an update transaction
// and submitted to ordering, then the committing peers will re-execute to guarantee that
// result sets are stable between endorsement time and commit time. The transaction is
// invalidated by the committing peers if the result set has changed between endorsement
// time and commit time.
// Therefore, range queries are a safe option for performing update transactions based on query results.
func (c *FinSmartContract) GetInvoiceOpsByRange(ctx contractapi.TransactionContextInterface, startKey, endKey string) ([]*InvoiceOps, error) {

	// Working with identities!!!
	reviewer, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return nil, fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(reviewer) == "User1" {
		return nil, fmt.Errorf("you are not authorized to approve a new purchase order")
	}

	resultsIteratorInvoiceOps, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}
	defer resultsIteratorInvoiceOps.Close()

	return constructQueryResponseFromIteratorInvoiceOps(resultsIteratorInvoiceOps)
}

//=========================================================================================================================
//=========================================================================================================================
//=========================================================================================================================
// QUERY THE OPERATIONAL SERVICE INVOICES [QUERY] - WORKING!
//=========================================================================================================================

// QueryInvoiceOps uses a query string to perform a query for purchase orders.
// Query string matching state database syntax is passed in and executed as is.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the QueryInvoiceOpsForStatus example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
func (c *FinSmartContract) QueryInvoiceOps(ctx contractapi.TransactionContextInterface, queryStringInvoiceOps string) ([]*InvoiceOps, error) {

	// Working with identities!!!
	reviewer, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return nil, fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(reviewer) == "User1" {
		return nil, fmt.Errorf("you are not authorized to approve a new purchase order")
	}

	return getQueryResultForQueryStringInvoiceOps(ctx, queryStringInvoiceOps)
}

//=========================================================================================================================
// QUERY THE ADDITIONAL WORK INVOICES [QUERY]
//=========================================================================================================================

// QueryInvoiceWork uses a query string to perform a query for purchase orders.
// Query string matching state database syntax is passed in and executed as is.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the QueryInvoiceWorkForStatus example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
func (c *FinSmartContract) QueryInvoiceWork(ctx contractapi.TransactionContextInterface, queryStringInvoiceWork string) ([]*InvoiceWork, error) {

	// Working with identities!!!
	reviewer, err1 := ctx.GetStub().GetCreator()
	if err1 != nil {
		return nil, fmt.Errorf("failed while getting identity. %s", err1.Error())
	}

	// Put identity from the Kaleido platform!!!
	// For now, I will put my favorite one "User1"
	if string(reviewer) == "User1" {
		return nil, fmt.Errorf("you are not authorized to approve a new purchase order")
	}

	return getQueryResultForQueryStringInvoiceWork(ctx, queryStringInvoiceWork)
}

//=========================================================================================================================
// FUNCTION GET QUERY RESULTS FOR QUERY STRING [QUERY]
//=========================================================================================================================

// getQueryResultForQueryString executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryStringInvoiceOps(ctx contractapi.TransactionContextInterface, queryString string) ([]*InvoiceOps, error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructQueryResponseFromIteratorInvoiceOps(resultsIterator)
}

//=========================================================================================================================
// FUNCTION GET QUERY RESULTS FOR QUERY STRING [QUERY]
//=========================================================================================================================

// getQueryResultForQueryString executes the passed in query string.
// The result set is built and returned as a byte array containing the JSON results.
func getQueryResultForQueryStringInvoiceWork(ctx contractapi.TransactionContextInterface, queryString string) ([]*InvoiceWork, error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructQueryResponseFromIteratorInvoiceWork(resultsIterator)
}

//=========================================================================================================================
//=========================================================================================================================
//=========================================================================================================================
// 											CROSS-CHAINCODE CALLS OPERATIONAL SERVICE INVOICES
//=========================================================================================================================
//
// These two functions of the chaincode are called for two reasons, one is that we check the Downtime Event by ID and we check its current status.
// Current status means: "0" - not validated, and "1" - validated;
//
//=========================================================================================================================
// CALL FUNCTION FROM OTHER CHAINCODE QUERYING THE DOWNTIME EVENT BY ID - WORKING!
//=========================================================================================================================

// CallingCrossChaincodeDowntimeEventByID call for a function form the other smart contract,
// in this case calling the smart contract for corrective maintenance.
// "downtimeevent" chaincode is from my working folder, chaincode for employees in minifabric (tested).
//
// IMPORTANT:
// "functionName" => function that needs to be called, when tested I called "GetDowntimeEventByID"
// "documentID"   => this is the parameter from the function we use, in this case we use ID number e.g. "DE-0005".

func (c *FinSmartContract) CallingCrossChaincodeDowntimeEventByID(ctx contractapi.TransactionContextInterface, functionName string, documentData string) (string, error) {

	if len(documentData) == 0 {
		return "", fmt.Errorf("please provide correct document data for")
	}

	params := []string{functionName, documentData}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := ctx.GetStub().InvokeChaincode("downtimeevent", queryArgs, "default-channel")

	return string(response.Payload), nil
}

//=========================================================================================================================
// CALL FUNCTION FROM OTHER CHAINCODE QUERYING THE DOWNTIME EVENT BY ITS STATUS - WORKING!
//=========================================================================================================================

// CallingCrossChaincodeDowntimeEventByTheConditionStatus call for a function form the other smart contract,
// in this case calling the smart contract for corrective maintenance.
// "downtimeevent" chaincode is installed on Kaleido.
//
// IMPORTANT:
// "functionName" => function that needs to be called, when tested I called "GetTheStatusOfDowntimeEventCondition"
// "documentID"   => this is the parameter from the function we use, in this case we use ID number e.g. "DE-0005".
// The result shall be "1", as this downtime event is validated

func (c *FinSmartContract) CallingCrossChaincodeDowntimeEventByTheConditionStatus(ctx contractapi.TransactionContextInterface, functionName string, documentData string) (string, error) {

	if len(documentData) == 0 {
		return "", fmt.Errorf("please provide correct document data for")
	}

	params := []string{functionName, documentData}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := ctx.GetStub().InvokeChaincode("downtimeevent", queryArgs, "default-channel")

	return string(response.Payload), nil
}

//=========================================================================================================================
// CALL FUNCTION FROM OTHER CHAINCODE QUERYING THE DOWNTIME EVENT BY ITS STATUS - WORKING!
//=========================================================================================================================

// CallingCrossChaincodeDowntimeEventByTheExclusion call for a function form the other smart contract,
// in this case calling the smart contract for corrective maintenance.
// "downtimeevent" chaincode is installed on Kaleido.
//
// IMPORTANT:
// "functionName" => function that needs to be called, when tested I called "QueryDowntimeEventsByTheExclusionStatus"
// "documentID"   => this is the parameter from the function we use, in this case we use ID number e.g. "Downtime Event Exclusion - YES".
// The result shall be "1", as this downtime event is validated

func (c *FinSmartContract) CallingCrossChaincodeDowntimeEventByTheExclusion(ctx contractapi.TransactionContextInterface, functionName string, documentData string) (string, error) {

	if len(documentData) == 0 {
		return "", fmt.Errorf("please provide correct document data for")
	}

	params := []string{functionName, documentData}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}

	response := ctx.GetStub().InvokeChaincode("downtimeevent", queryArgs, "default-channel")

	return string(response.Payload), nil
}
