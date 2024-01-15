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

type Employee struct {
	DocType                   string `json:"employee"`
	CompanySite               string `json:"CompanySite"`
	ID                        string `json:"ID"`
	Name                      string `json:"Name"`
	LastName                  string `json:"LastName"`
	Gender                    string `json:"Gender"`
	PositionEntry             string `json:"PositionEntry"`
	CurrentPosition           string `json:"CurrentPosition"` // For the current position, e.g. Maintenance Technician Tier 1 or System Operator Tier 3 (Technical Operator) etc.
	LevelPT                   string `json:"LevelPT"`         // It will be used only to check if the Company still employs Employee(s).
	Birthday                  string `json:"Birthday"`
	ContractSigned            string `json:"ContractSigned"`
	StartingDate              string `json:"StartingDate"`
	YearsInService            string `json:"YearsInService"`
	EmployeeReview            string `json:"EmployeeReview"`
	CustomTraining            string `json:"CompletedCustomTrining"`
	StatusTrainingModuleOne   string `json:"StatusTrainingModuleOne"`
	StatusTrainingModuleTwo   string `json:"StatusTrainingModuleTwo"`
	StatusTrainingModuleThree string `json:"StatusTrainingModuleThree"`
	Salary                    int    `json:"Salary"`
	Address                   string `json:"Address"`
	Cell                      string `json:"Cell"`
	WorkVisa                  string `json:"WorkVisa"`
}
