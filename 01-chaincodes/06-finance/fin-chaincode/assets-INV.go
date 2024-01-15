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

// Structure of the OPS Invoice object/asset
type InvoiceOps struct {
	DocType                     string `json:"docType"`          // here we can put "invoice"
	InvoiceTypeOps              string `json:"invoice_type_ops"` // here we can put "operations"
	InvoiceOpsID                string `json:"invoice__ops_id"`  // e.g. INV-OPS-0001
	CompanyRegistrationNumber   string `json:"company_registration_number"`
	InvoiceOpsDateCreated       string `json:"invoice_ops_date_created"`
	PaymentTerms                string `json:"payment_terms"`
	BillingPeriodOpsServices    string `json:"billing_period_ops_services"` // Dates => from - to
	InvoiceDescription          string `json:"invoice_description"`
	InvoiceCostCode             int    `json:"invoice_cost_code"`
	InvoiceOperationalCost      int    `json:"invoice_operational_cost"`
	VariationAdditionalExpenses int    `json:"variation_additional_expenses"` // to be added cost, e.g. additional operating hours
	DowntimeDeduction           int    `json:"downtime_deduction"`
	CPIAdjustment               int    `json:"cpi_adjustment"`
	OtherAgreedReductions       int    `json:"other_agreed_reductions"`
	TotalServiceInvoiceCost     uint   `json:"total_service_invoice_cost"` // How to calculate all these inputs
	InvoiceApprovalStatus       string `json:"invoice_approval_status"`
	InvoiceStatus               string `json:"invoice_status"`
}

type InvoiceItemOne struct {
	InvoiceItemOneQty      int `json:"invoice_item_one_qty"`
	InvoiceItemOneItemCost int `json:"invoice_item_one_item_cost"`
	InvoiceItemOneCost     int `json:"invoice_item_one_cost"`
}

type InvoiceItemTwo struct {
	InvoiceItemTwoQty      int `json:"invoice_item_two_qty"`
	InvoiceItemTwoItemCost int `json:"invoice_item_two_item_cost"`
	InvoiceItemTwoCost     int `json:"invoice_item_two_cost"`
}

type InvoiceItemThree struct {
	InvoiceItemThreeQty      int `json:"invoice_item_three_qty"`
	InvoiceItemThreeItemCost int `json:"invoice_item_three_item_cost"`
	InvoiceItemThreeCost     int `json:"invoice_item_three_cost"`
}

type InvoiceItemFour struct {
	InvoiceItemFourQty      int `json:"invoice_item_four_qty"`
	InvoiceItemFourItemCost int `json:"invoice_item_four_item_cost"`
	InvoiceItemFourCost     int `json:"invoice_item_four_cost"`
}

type InvoiceItemFive struct {
	InvoiceItemFiveQty      int `json:"invoice_item_five_qty"`
	InvoiceItemFiveItemCost int `json:"invoice_item_five_item_cost"`
	InvoiceItemFiveCost     int `json:"invoice_item_five_cost"`
}

// Structure of the OPS Invoice object/asset
type InvoiceWork struct {
	DocType                            string           `json:"docType"`           // here we can put "invoice"
	InvoiceTypeWork                    string           `json:"invoice_type_work"` // here we can put "additional-work"
	InvoiceWorkID                      string           `json:"invoice_work_id"`   // e.g. INV-WORK-20200101-0001
	CompanyRegistrationNumber          string           `json:"company_registration_number"`
	InvoiceWorkDateCreated             string           `json:"invoice_work_date_created"`
	PaymentTerms                       string           `json:"payment_terms"`
	BillingPeriodAdditionalWork        string           `json:"billing_period_additional_work"` // Dates => from - to
	InvoiceDescription                 string           `json:"invoice_description"`
	InvoiceWorkCostCode                int              `json:"invoice_work_cost_code"`
	VariationAdditionalExpensesAddWork int              `json:"variation_additional_expenses_add_work"`
	InvoiceWorkItemOne                 InvoiceItemOne   `json:"invoice_item_one"`            // here we enter all items which needs to be listed in the invoice
	InvoiceWorkItemTwo                 InvoiceItemTwo   `json:"invoice_item_two"`            // here we enter all items which needs to be listed in the invoice
	InvoiceWorkItemThree               InvoiceItemThree `json:"invoice_item_three"`          // here we enter all items which needs to be listed in the invoice
	InvoiceWorkItemFour                InvoiceItemFour  `json:"invoice_item_four"`           // here we enter all items which needs to be listed in the invoice
	InvoiceWorkItemFive                InvoiceItemFive  `json:"invoice_item_five"`           // here we enter all items which needs to be listed in the invoice
	TotalAddWorkInvoiceCost            uint             `json:"total_add_work_invoice_cost"` // How to calculate all these inputs
	InvoiceApprovalStatus              string           `json:"invoice_approval_status"`
	InvoiceStatus                      string           `json:"invoice_status"`
}
