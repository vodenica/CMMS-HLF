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
	"math/rand"
	"time"
)

//=================================================================================================================================
// OPERATIONAL SERVICES INVOICE
//=================================================================================================================================
/*
// SetInvoiceDocType set the document type of the invoice
func (in *InvoiceOps) SetInvoiceDocType() {
	in.DocType = "invoice-ops"
}
*/

// SetInvoiceID set the invoice ID form the current date and time converted into running 5 digits HEX numbers (e.g. INV-OPS-20181001-00001)

func (in *InvoiceOps) SetInvoiceOpsID() {
	in.InvoiceOpsID = "INV-OPS-" + time.Now().Format("20060102") + "-" + RandStringRunes(5)
}

func RandStringRunes(i int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, i)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// SetInvoiceDescription set the description of the invoice
func (in *InvoiceOps) SetInvoiceDescription() {
	in.InvoiceDescription = "Invoice for Operational Services" + " " + time.Now().Format("January 2006")
}

// SetCompanyRegistrationNumber set the company registration number
func (in *InvoiceOps) SetCompanyRegistrationNumber() {
	in.CompanyRegistrationNumber = "Reg.No. 123456789"
}

// SetInvoiceDate set the date of the invoice
func (in *InvoiceOps) SetInvoiceOpsDateCreated() {
	in.InvoiceOpsDateCreated = time.Now().Format("2006-01-02")
}

// SetPaymentTerms set the payment terms
func (in *InvoiceOps) SetPaymentTerms() {
	in.PaymentTerms = "30 days"
}

// SetBillingPeriodOpsServices set the billing period of the invoice which shall cover the previous month
func (in *InvoiceOps) SetBillingPeriodOpsServices() {
	in.BillingPeriodOpsServices = "Billing Period For Operational Services Rendered in " + time.Now().AddDate(0, -1, 0).Format("January 2006")
}

// SetInvoiceCostCode set the cost code as integer
func (in *InvoiceOps) SetInvoiceCostCode() {
	in.InvoiceCostCode = 1000
}

// SetInvoiceOpsApprovalStatusPending set the condition of the Invoice to PENDING
func (in *InvoiceOps) SetInvoiceOpsApprovalStatusPending() {
	in.InvoiceApprovalStatus = "Pending"
}

// SetInvoiceApprovalStatusApproved set the condition of the Invoice to APPROVED
func (in *InvoiceOps) SetInvoiceOpsApprovalStatusApproved() {
	in.InvoiceApprovalStatus = "Approved"
}

// SetInvoiceApprovalStatusDeclined set the condition of the Invoice to DECLINED
func (in *InvoiceOps) SetInvoiceOpsApprovalStatusDeclined() {
	in.InvoiceApprovalStatus = "Declined"
}

//=================================================================================================================================
// SetInvoiceStatusOpen set the status of the Invoice to ISSUED
func (in *InvoiceOps) SetInvoiceOpsStatusIssued() {
	in.InvoiceStatus = "Issued"
}

// SetInvoiceStatusRevision set the status of the Invoice to REVISION
func (in *InvoiceOps) SetInvoiceOpsStatusRevision() {
	in.InvoiceStatus = "Revision"
}

// SetInvoiceStatusPaid set the status of the Invoice to PAID
func (in *InvoiceOps) SetInvoiceOpsStatusPaid() {
	in.InvoiceStatus = "Paid"
}

//=================================================================================================================================
// ADDITIONAL WORK INVOICE
//=================================================================================================================================
// SetInvoiceID set the invoice ID form the current date and time converted into running 5 digits HEX numbers (e.g. INV-OPS-20181001-00001)

func (iw *InvoiceWork) SetInvoiceWorkID() {
	iw.InvoiceWorkID = "INV-WORK-" + time.Now().Format("20060102") + "-" + RandStringRunes(5)
}

// SetCompanyRegistrationNumber set the company registration number
func (iw *InvoiceWork) SetCompanyRegistrationNumber() {
	iw.CompanyRegistrationNumber = "Reg.No. 123456789"
}

// SetInvoiceDate set the date of the invoice
func (iw *InvoiceWork) SetInvoiceWorkDateCreated() {
	iw.InvoiceWorkDateCreated = time.Now().Format("2006-01-02")
}

// SetPaymentTerms set the payment terms
func (iw *InvoiceWork) SetPaymentTerms() {
	iw.PaymentTerms = "30 days"
}

// SetBillingPeriodForAdditionalWork set the billing period of the invoice
func (iw *InvoiceWork) SetBillingPeriodForAdditionalWork() {
	iw.BillingPeriodAdditionalWork = "Billing Period For Additional Work Rendered from 2020-01-01 to 2020-01-31"
}

// SetInvoiceOpsApprovalStatusPending set the condition of the Invoice to PENDING
func (iw *InvoiceWork) SetInvoiceWorkApprovalStatusPending() {
	iw.InvoiceApprovalStatus = "Pending"
}

// SetInvoiceApprovalStatusApproved set the condition of the Invoice to APPROVED
func (iw *InvoiceWork) SetInvoiceWorkApprovalStatusApproved() {
	iw.InvoiceApprovalStatus = "Approved"
}

// SetInvoiceApprovalStatusDeclined set the condition of the Invoice to DECLINED
func (iw *InvoiceWork) SetInvoiceWorkApprovalStatusDeclined() {
	iw.InvoiceApprovalStatus = "Declined"
}

//=================================================================================================================================
// SetInvoiceStatusOpen set the status of the Invoice to ISSUED
func (iw *InvoiceWork) SetInvoiceWorkStatusIssued() {
	iw.InvoiceStatus = "Issued"
}

// SetInvoiceStatusRevision set the status of the Invoice to REVISION
func (iw *InvoiceWork) SetInvoiceWorkStatusRevision() {
	iw.InvoiceStatus = "Revision"
}

// SetInvoiceStatusPaid set the status of the Invoice to PAID
func (iw *InvoiceWork) SetInvoiceWorkStatusPaid() {
	iw.InvoiceStatus = "Paid"
}

/*
//=================================================================================================================================
// Cost Codes which will be used
//=================================================================================================================================

// Cost Code => 1000 Payroll
type Payroll struct {
	SalariesAndBonuses int `json:"salaries_and_bonuses_1001"`
	Overtime           int `json:"overtime_1002"`
}

// Cost Code => 2000 Office Equipment
type OfficeExpenses struct {
	OfficeEquipment int `json:"office_equipment_2001"`
	ITEquipment     int `json:"it_equipment_20002"`
}

// Cost Code => 3000 Training & Certification
type TrainingAndCertification struct {
	SafetyTraining      int `json:"safety_training_3001"`
	MaintenanceTraining int `json:"maintenance_training_3001"`
}

// Cost Code 5000 => Subcontracted Maintenance Services
type SubcontractedMaintenanceServices struct {
	MRTAndNDTInspections int `json:"mrt_and_ndt_inspections_5001"`
	CertificationOfTE    int `json:"certification_of_tools_end_equipment_5002"`
}

// CostCode 6000 => Maintenance Parts Tools and Equipment
type MaintenancePartsToolsAndEquipment struct {
	Consumables       int `json:"consumables_6001"`
	MaintenanceParts  int `json:"maintenance_parts_6002"`
	ToolsAndEquipment int `json:"tools_and_equipment_6003"`
}

// Cost Code 7000 => Expenses
type Expenses struct {
	ExpensesAdditionalWork int `json:"expenses_additional_work_7002"`
	ExpensesCARP           int `json:"expenses_carp_7003"`
}

// Cost Code 8000 => Income
type Income struct {
	IncomeServiceAgreement int `json:"income_service_agreement_8001"`
	IncomeAdditionalWork   int `json:"income_additional_work_8002"`
	IncomeCARP             int `json:"income_carp_8003"`
}
*/
