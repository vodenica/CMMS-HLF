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
	"time"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
)

//==============================================================================================================================
// PREVENTIVE MAINTENANCE
//==============================================================================================================================

// SetPrevWorkOrderCreator sets and get the MSP ID of the creator of the work order
func (ba *PreventiveMaintenance) SetPrevWorkOrderCreator(stub cid.ChaincodeStubInterface) {
	ba.PreventiveWorkOrderCreator, _ = cid.GetMSPID(stub)
}

// Setting the next work order date adding one month to the current date.
func (ba *PreventiveMaintenance) SetNextPreventiveWorkOrder() {
	ba.NextPreventiveWorkOrder = time.Now().AddDate(0, 1, 0).Format("2006-01-02 15:04:05")
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

//==============================================================================================================================
// CORRECTIVE MAINTENANCE
//==============================================================================================================================

// SetCorrWorkOrderCreator sets and get the MSP ID of the creator of the work order
func (ba *CorrectiveMaintenance) SetCorrWorkOrderCreator(stub cid.ChaincodeStubInterface) {
	ba.CorrectiveWorkOrderCreator, _ = cid.GetMSPID(stub)
}

// SetCorrWorkOrderCompletionDate sets the completion date of the work order adding 15 days to the current date.
func (ba *CorrectiveMaintenance) SetCorrWorkOrderCompletionDate() {
	ba.CorrWorkOrderCompletionDate = time.Now().AddDate(0, 0, 15).Format("2006-01-02 15:04:05")
}

// SetCorrWorkOrderConditionGood set the condition of the work order to mark as good
func (ba *CorrectiveMaintenance) SetCorrWorkOrderConditionGood() {
	ba.CorrCondition = "good"
}

// SetCorrWorkOrderConditionNotGood set the condition of the work order to mark as not-good
func (ba *CorrectiveMaintenance) SetCorrWorkOrderConditionNotGood() {
	ba.CorrCondition = "not-good"
}

// SetCorrWorkOrderValidated set the condition of the work order to mark as good
func (ba *CorrectiveMaintenance) SetCorrWorkOrderValidated() {
	ba.CorrValidation = "validated"
}

// SetCorrWorkOrderConditionNotGood set the condition of the work order to mark as not-good
func (ba *CorrectiveMaintenance) SetCorrWorkOrderNotValidated() {
	ba.CorrValidation = "not-validated"
}

//==============================================================================================================================
// FAILURE MAINTENANCE
//==============================================================================================================================

// SetFailWorkOrderCreator sets and get the MSP ID of the creator of the work order
func (ba *FailureMaintenance) SetFailWorkOrderCreator(stub cid.ChaincodeStubInterface) {
	ba.FailureWorkOrderCreator, _ = cid.GetMSPID(stub)
}

/*
// Set FailWorkOrderTimestamp sets the timestamp of the work order
func (ba *FailureMaintenance) SetFailWorkOrderTimestamp() {
	ba.FailWorkOrderTimestamp = time.Now().Format("2006-01-02 15:04:05")
}
*/

// SetFailWorkOrderCompletionDate sets the completion date of the work order adding 15 days to the current date.
func (ba *FailureMaintenance) SetFailWorkOrderCompletionDate() {
	ba.FailWorkOrderCompletionDate = time.Now().AddDate(0, 0, 15).Format("2006-01-02 15:04:05")
}

// SetFailWorkOrderConditionGood set the condition of the work order to mark as good
func (ba *FailureMaintenance) SetFailWorkOrderConditionGood() {
	ba.FailCondition = "good"
}

// SetFailWorkOrderConditionNotGood set the condition of the work order to mark as not-good
func (ba *FailureMaintenance) SetFailWorkOrderConditionNotGood() {
	ba.FailCondition = "not-good"
}

// SetFailWorkOrderValidated set the condition of the work order to mark as good
func (ba *FailureMaintenance) SetFailWorkOrderValidated() {
	ba.FailValidation = "validated"
}

// SetFailWorkOrderConditionNotGood set the condition of the work order to mark as not-good
func (ba *FailureMaintenance) SetFailWorkOrderNotValidated() {
	ba.FailValidation = "not-validated"
}

//==============================================================================================================================
