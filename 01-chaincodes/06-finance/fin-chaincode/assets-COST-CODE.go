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

//=======================================================================
// Cost Codes which can be used
//=======================================================================

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

/*
// CostCode 6000 => Maintenance Parts Tools and Equipment
type MaintenancePartsToolsAndEquipment struct {
	Consumables       int `json:"consumables_6001"`
	MaintenanceParts  int `json:"maintenance_parts_6002"`
	ToolsAndEquipment int `json:"tools_and_equipment_6003"`
}
*/

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
