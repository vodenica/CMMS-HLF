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
	//"math/rand"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	//"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//=================================================================================================================================
// SetInvoiceID set the invoice ID form the current date and time converted into running 5 digits HEX numbers (e.g. INV-OPS-20181001-00001)

func (pr *PurchaseRequest) SetPurchaseOrderID() {
	pr.PurchaseOrderID = "PURCH-" + time.Now().Format("20060102") + "-" + RandStringRunes(5)
}

// SetPurchaseOrderCreationDate set the creation date of the purchase order
func (pr *PurchaseRequest) SetPurchaseOrderCreationDate() {
	pr.PurchaseOrderCreated = time.Now().Format("2006-01-02")
}

// Setting the Purchase Order Requestor identity with function
func (pr *PurchaseRequest) SetPurchaseOrderCreatorIdentity(stub cid.ChaincodeStubInterface) {
	pr.PurchaseOrderRequester, _ = cid.GetMSPID(stub)
}

// Setting the Purchase Order Approver identity with function to "not-set"
func (pr *PurchaseRequest) SetPurchaseOrderApproverIdentityToNotSet() {
	pr.PurchaseOrderApprover = "not-set"
}

// SetPurchaseRequestUrgent set the condition of the purchase order to URGENT
func (pr *PurchaseRequest) SetPurchaseRequestNonUrgent() {
	pr.StatusOfTheUrgency = 0
}

// SetPurchaseRequestUrgent set the condition of the purchase order to URGENT
func (pr *PurchaseRequest) SetPurchaseRequestUrgent() {
	pr.StatusOfTheUrgency = 1
}

// SetPurchaseRequestApproved set the condition of the purchase order to APPROVED
func (pr *PurchaseRequest) SetPurchaseRequestOpen() {
	pr.StatusOfTheOrder = "Open"
}

// SetPurchaseRequestApproved set the condition of the purchase order to APPROVED
func (pr *PurchaseRequest) SetPurchaseRequestApproved() {
	pr.StatusOfTheOrder = "Approved"
}

// SetPurchaseRequestDeclined set the condition of the purchase order to DECLINED
func (pr *PurchaseRequest) SetPurchaseRequestDeclined() {
	pr.StatusOfTheOrder = "Declined"
}

// SetPurchaseRequestDeclined set the condition of the purchase order to COMPLETED
func (pr *PurchaseRequest) SetPurchaseRequestCompleted() {
	pr.StatusOfTheOrder = "Completed"
}

//=================================================================================================================================
// SetSignOffReceivedGoods to not set
func (pr *PurchaseRequest) SetSignOffReceivedGoodsToNotSet() {
	pr.SignOffReceivedGoods = "not-set"
}

//=================================================================================================================================
// SetVerifyQuantityToNotDelivered set the condition of the purchase order to NOT-COMPLETE-DELIVERED
func (pr *PurchaseRequest) SetVerifyDeliveryToNotDelivered() {
	pr.VerifyDelivery = 0
}

// SetVerifyDeliveryToPartiallyDelivered set the condition of the purchase order to NOT-COMPLETE-DELIVERED
func (pr *PurchaseRequest) SetVerifyDeliveryToPartiallyDelivered() {
	pr.VerifyDelivery = 1
}

// SetVerifyDeliveryToDelivered set the condition of the purchase order to NOT-COMPLETE-DELIVERED
func (pr *PurchaseRequest) SetVerifyDeliveryToDelivered() {
	pr.VerifyDelivery = 2
}

//=================================================================================================================================
// SetVerifyQuantityNotDelivered set the condition of the purchase order to NOT-COMPLETE-DELIVERED
func (pr *PurchaseRequest) SetVerifyQuantityNotDelivered() {
	pr.VerifyQuantity = 0
}

// SetVerifyQuantityAllDelivered set the condition of the purchase order to NOT-COMPLETE-DELIVERED
func (pr *PurchaseRequest) SetVerifyQuantityAllDelivered() {
	pr.VerifyQuantity = 1
}

//=================================================================================================================================
// SetVerifyQuantityNotDelivered set the condition of the purchase order to NOT-COMPLETE-DELIVERED
func (pr *PurchaseRequest) SetVerifyConditionOpen() {
	pr.VerifyCondition = 0
}

// SetVerifyQuantityAllDelivered set the condition of the purchase order to NOT-COMPLETE-DELIVERED
func (pr *PurchaseRequest) SetVerifyConditionNotGood() {
	pr.VerifyCondition = 1
}

// SetVerifyQuantityAllDelivered set the condition of the purchase order to NOT-COMPLETE-DELIVERED
func (pr *PurchaseRequest) SetVerifyConditionGood() {
	pr.VerifyCondition = 2
}
