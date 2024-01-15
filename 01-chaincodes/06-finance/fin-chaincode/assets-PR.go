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

// RequestItemOne is the structure of the first item in the purchase request
type RequestItemOne struct {
	CostCodeItemOne  string `json:"cost_code_item_one"`
	ItemDescription  string `json:"item_description"`
	ItemOneCost      int    `json:"item_one_cost"`
	ItemOneQuantity  int    `json:"item_one_quantity"`
	ItemOneTotalCost int    `json:"item_one_total_cost"`
}

// RequestItemTwo is the structure of the second item in the purchase request
type RequestItemTwo struct {
	CostCodeItemTwo  string `json:"cost_code_item_two"`
	ItemDescription  string `json:"item_description"`
	ItemTwoCost      int    `json:"item_two_cost"`
	ItemTwoQuantity  int    `json:"item_two_quantity"`
	ItemTwoTotalCost int    `json:"item_two_total_cost"`
}

// RequestItemThree is the structure of the third item in the purchase request
type RequestItemThree struct {
	CostCodeItemThree  string `json:"cost_code_item_three"`
	ItemDescription    string `json:"item_description"`
	ItemThreeCost      int    `json:"item_three_cost"`
	ItemThreeQuantity  int    `json:"item_three_quantity"`
	ItemThreeTotalCost int    `json:"item_three_total_cost"`
}

// RequestItemFour is the structure of the fourth item in the purchase request
type RequestItemFour struct {
	CostCodeItemFour  string `json:"cost_code_item_four"`
	ItemDescription   string `json:"item_description"`
	ItemFourCost      int    `json:"item_four_cost"`
	ItemFourQuantity  int    `json:"item_four_quantity"`
	ItemFourTotalCost int    `json:"item_four_total_cost"`
}

// RequestItemFive is the structure of the fifth item in the purchase request
type RequestItemFive struct {
	CostCodeItemFive  string `json:"cost_code_item_five"`
	ItemDescription   string `json:"item_description"`
	ItemFiveCost      int    `json:"item_five_cost"`
	ItemFiveQuantity  int    `json:"item_five_quantity"`
	ItemFiveTotalCost int    `json:"item_five_total_cost"`
}

// Structure of the purchasing request
type PurchaseRequest struct {
	DocType                  string           `json:"docType"` // here we can put "purchase-request"
	PurchaseRequestType      string           `json:"purchase_request_type"`
	PurchaseOrderID          string           `json:"purchase_order_id"`
	PurchaseOrderCreated     string           `json:"purchase_order_created"`
	PurchaseOrderRequester   string           `json:"purchase_order_requestor"`
	PurchaseOrderApprover    string           `json:"purchase_order_approver"`
	PurchaseRequestItemOne   RequestItemOne   `json:"purchase_request_item_one"`
	PurchaseRequestItemTwo   RequestItemTwo   `json:"purchase_request_item_two"`
	PurchaseRequestItemThree RequestItemThree `json:"purchase_request_item_three"`
	PurchaseRequestItemFour  RequestItemFour  `json:"purchase_request_item_four"`
	PurchaseRequestItemFive  RequestItemFive  `json:"purchase_request_item_five"`
	SupplierOrderID          string           `json:"supplier_order_id"`
	SupplierOrderDescription string           `json:"supplier_order_description"`
	PurchaseOrderTotalCost   int              `json:"purchase_order_total_cost"`
	StatusOfTheUrgency       int              `json:"status_of_the_urgency"`
	StatusOfTheOrder         string           `json:"status_of_the_order"`
	VerifyDelivery           int              `json:"verify_delivery"`
	VerifyQuantity           int              `json:"verify_quantity"`
	VerifyCondition          int              `json:"verify_condition"`
	SignOffReceivedGoods     string           `json:"sign_off_received_goods"`
	PurchaseRequestRemarks   string           `json:"purchase_request_remarks"`
}
