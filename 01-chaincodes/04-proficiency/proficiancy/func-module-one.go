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

// In here we will create the object to manage, in this case we will call it ModuleOne

package main

import (
	time "time"
)

// ModuleOneTrainingCreatedDate set the date when the training module one was created.
func (ba *ModuleOne) SetModuleOneCreatedDate() {
	ba.ModuleOneTrainingCreatedDate = time.Now().Format("2006-01-02 15:04:05")
}

// SetStatusTrainingTypeInitial set the status of the training type to mark as initial.
func (ba *ModuleOne) SetStatusTrainingTypeInitial() {
	ba.TrainingTypeModuleOne = 0
}

// SetStatusTrainingTypeRefresher set the status of the training type to mark as refresher.
func (ba *ModuleOne) SetStatusTrainingTypeRefresher() {
	ba.TrainingTypeModuleOne = 1
}

//==============================================================================================================

// SetStatusChapterOneModuleOneOpen set the status of the training chapter to mark as open.
func (ba *ModuleOne) SetStatusChapterOneModuleOneOpen() {
	ba.SetStatusChapterOneModuleOne = 0
}

// SetStatusChapterOneModuleOneCompleted set the status of the training chapter to mark as completed.
func (ba *ModuleOne) SetStatusChapterOneModuleOneCompleted() {
	ba.SetStatusChapterOneModuleOne = 1
}

// SetStatusChapterTwoModuleOneOpen set the status of the training chapter to mark as open.
func (ba *ModuleOne) SetStatusChapterTwoModuleOneOpen() {
	ba.SetStatusChapterTwoModuleOne = 0
}

// SetStatusChapterTwoModuleOneCompleted set the status of the training chapter to mark as completed.
func (ba *ModuleOne) SetStatusChapterTwoModuleOneCompleted() {
	ba.SetStatusChapterTwoModuleOne = 1
}

// SetStatusChapterThreeModuleOneOpen set the status of the training chapter to mark as open.
func (ba *ModuleOne) SetStatusChapterThreeModuleOneOpen() {
	ba.SetStatusChapterThreeModuleOne = 0
}

// SetStatusChapterThreeModuleOneCompleted set the status of the training chapter to mark as completed.
func (ba *ModuleOne) SetStatusChapterThreeModuleOneCompleted() {
	ba.SetStatusChapterThreeModuleOne = 1
}

//==============================================================================================================

// SetStatusModuleOneOpen set the status of the training module to mark as open, as it is not completed.
func (ba *ModuleOne) SetStatusModuleOneOpen() {
	ba.TrainingStatusModuleOne = 1
}

// SetStatusModuleOneCompleted set the status of the training module to mark as completed.
func (ba *ModuleOne) SetStatusModuleOneCompleted() {
	ba.TrainingStatusModuleOne = 3
}

// SetStatusModuleOneCanceled set the status of the training module to mark as canceled, employee leave during the probation period.
func (ba *ModuleOne) SetStatusModuleOneCanceled() {
	ba.TrainingStatusModuleOne = 4
}

//==============================================================================================================

// SetTheoreticalAssessmentModuleOne set the value of the theoretical assessment of the training module one.
func (ba *ModuleOne) SetTheoreticalAssessmentModuleOne() {
	ba.TheoreticalAssessmentModuleOne = 0
}

// SetPracticalAssessmentModuleOne set the value of the theoretical assessment of the training module one.
func (ba *ModuleOne) SetPracticalAssessmentModuleOne() {
	ba.PracticalAssessmentModuleOne = 0
}

// SetTotalAssessmentModuleOne set the value of the theoretical assessment of the training module one.
func (ba *ModuleOne) SetTotalAssessmentModuleOne() {
	ba.AssessmentModuleOne = 0
}

//==============================================================================================================

// SetAssessmentsAttemptsToZero set the value of the both assessment of the training module one to zero.
func (ba *ModuleOne) SetAssessmentsAttemptsToZero() {
	ba.AssessmentsAttempts = 0
}

// SetAssessmentsAttemptsToOne set the value of the both assessment of the training module one to zero.
func (ba *ModuleOne) SetAssessmentsAttemptsToOne() {
	ba.AssessmentsAttempts = 1
}

// SetAssessmentsAttemptsToTwo set the value of the both assessment of the training module one to zero.
func (ba *ModuleOne) SetAssessmentsAttemptsToTwo() {
	ba.AssessmentsAttempts = 2
}

// SetAssessmentsAttemptsToThree set the value of the both assessment of the training module one to zero.
func (ba *ModuleOne) SetAssessmentsAttemptsToThree() {
	ba.AssessmentsAttempts = 3
}
