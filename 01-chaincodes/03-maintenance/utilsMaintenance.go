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

type PlannedLabour struct {
	MaintenanceSupervisor      string `json:"maintenance_supervisor"`
	MaintenanceTechnicianOne   string `json:"maintenance_technician_one"`
	MaintenanceTechnicianTwo   string `json:"maintenance_technician_two"`
	MaintenanceTechnicianThree string `json:"maintenance_technician_three"`
	SystemOperator             string `json:"system_operator"`
}

type MaintenanceProcessSteps struct {
	MaintenanceProcessStepOne   string `json:"maintenance_process_step_one"`
	MaintenanceProcessStepTwo   string `json:"maintenance_process_step_two"`
	MaintenanceProcessStepThree string `json:"maintenance_process_step_three"`
	MaintenanceProcessStepFour  string `json:"maintenance_process_step_four"`
	MaintenanceProcessStepFive  string `json:"maintenance_process_step_five"`
}

type MaintenanceParts struct {
	MaintenancePartOne   string `json:"maintenance_part_one"`
	MaintenancePartTwo   string `json:"maintenance_part_two"`
	MaintenancePartThree string `json:"maintenance_part_three"`
	MaintenancePartFour  string `json:"maintenance_part_four"`
	MaintenancePartFive  string `json:"maintenance_part_five"`
	MaintenancePartSix   string `json:"maintenance_part_six"`
	MaintenancePartSeven string `json:"maintenance_part_seven"`
	MaintenancePartEight string `json:"maintenance_part_eight"`
	MaintenancePartNine  string `json:"maintenance_part_nine"`
	MaintenancePartTen   string `json:"maintenance_part_ten"`
}

type PreventiveMaintenance struct {
	DocType                      string                  `json:"docType"`
	PrevWorkOrderID              string                  `json:"prev_work_order_id"`
	PreventiveWorkOrderTimestamp string                  `json:"prev_work_order_timestamp"`
	NextPreventiveWorkOrder      string                  `json:"next_prev_work_order"`
	PreventiveWorkOrderCreator   string                  `json:"prev_work_order_creator"`
	PrevWorkOrderType            string                  `json:"prev_work_order_type"`
	PrevWorkOrderDescription     string                  `json:"prev_work_order_description"`
	PrevGeneralHandSInstructions string                  `json:"prev_general_instructions"`
	PrevPlannedLabour            PlannedLabour           `json:"prev_planned_labour"`
	PrevMaintenanceProcessSteps  MaintenanceProcessSteps `json:"prev_maintenance_process_steps"`
	PrevSparePartsUsed           MaintenanceParts        `json:"prev_spare_parts_used"`
	PrevCondition                string                  `json:"prev_condition"`
	PrevValidation               string                  `json:"prev_validation"`
}

type CorrectiveMaintenance struct {
	DocType                      string                  `json:"docType"`
	CorrWorkOrderID              string                  `json:"corr_work_order_id"`
	CorrWorkOrderTimestamp       string                  `json:"corr_work_order_timestamp"`
	CorrWorkOrderCompletionDate  string                  `json:"corr_work_order_completion_date"`
	CorrectiveWorkOrderCreator   string                  `json:"corr_work_order_creator"`
	CorrWorkOrderType            string                  `json:"corr_work_order_type"`
	CorrWorkOrderDescription     string                  `json:"corr_work_order_description"`
	CorrGeneralHandSInstructions string                  `json:"corr_general_instructions"`
	CorrPlannedLabour            PlannedLabour           `json:"corr_planned_labour"`
	CorrMaintenanceProcessSteps  MaintenanceProcessSteps `json:"corr_maintenance_process_steps"`
	CorrSparePartsUsed           MaintenanceParts        `json:"corr_spare_parts_used"`
	CorrCondition                string                  `json:"corr_condition"`
	CorrValidation               string                  `json:"corr_validation"`
}

type FailureMaintenance struct {
	DocType                      string                  `json:"docType"`
	FailWorkOrderID              string                  `json:"fail_work_order_id"`
	FailWorkOrderTimestamp       string                  `json:"fail_work_order_timestamp"`
	FailWorkOrderCompletionDate  string                  `json:"fail_work_order_completion_date"`
	FailureWorkOrderCreator      string                  `json:"fail_work_order_creator"`
	FailWorkOrderType            string                  `json:"fail_work_order_type"`
	FailFaultID                  string                  `json:"fault_id"`
	FailWorkOrderDescription     string                  `json:"fail_work_order_description"`
	FailGeneralHandSInstructions string                  `json:"fail_general_H&S_instructions"`
	FailPlannedLabour            PlannedLabour           `json:"fail_planned_labour"`
	FailMaintenanceProcessSteps  MaintenanceProcessSteps `json:"fail_maintenance_process_steps"`
	FailSparePartsUsed           MaintenanceParts        `json:"fail_spare_parts_used"`
	FailCondition                string                  `json:"fail_condition"`
	FailValidation               string                  `json:"fail_validation"`
}
