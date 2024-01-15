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

// TrainerModuleOne contains the full name of an trainer who will perform training in module one.
type TrainerModuleOne struct {
	TrainerModuleOneID      string `json:"trainer_module_one_id"`
	TrainerModuleOneName    string `json:"trainer_module_one_name"`
	TrainerModuleOneSurname string `json:"trainer_module_one_surname"`
}

// TraineeModuleOne contains the full name of an trainer who will perform training in module one.
type TraineeModuleOne struct {
	TraineeModuleOneID      string `json:"trainee_module_one_id"`
	TraineeModuleOneName    string `json:"trainee_module_one_name"`
	TraineeModuleOneSurname string `json:"trainee_module_one_surname"`
}

// ModuleOneChapterOne is an asset which will carry thi information's about the topics to be trained.
type ModuleOneChapterOne struct {
	ModuleOneChapterOneSessionOne   string `json:"module_one_chapter_one_session_one"`
	ModuleOneChapterOneSessionTwo   string `json:"module_one_chapter_one_session_two"`
	ModuleOneChapterOneSessionThree string `json:"module_one_chapter_one_session_three"`
	// SetStatusChapterOneModuleOne    int    `json:"set_status_chapter_one_module_one"`
}

// ModuleOneChapterTwo is an asset which will carry thi information's about the topics to be trained.
type ModuleOneChapterTwo struct {
	ModuleOneChapterTwoSessionOne   string `json:"module_one_chapter_two_session_one"`
	ModuleOneChapterTwoSessionTwo   string `json:"module_one_chapter_two_session_two"`
	ModuleOneChapterTwoSessionThree string `json:"module_one_chapter_two_session_three"`
	// SetStatusChapterTwoModuleOne    int    `json:"set_status_chapter_two_module_one"`
}

// ModuleOneChapterThree is an asset which will carry thi information's about the topics to be trained.
type ModuleOneChapterThree struct {
	ModuleOneChapterThreeSessionOne   string `json:"module_one_chapter_three_session_one"`
	ModuleOneChapterThreeSessionTwo   string `json:"module_one_chapter_three_session_two"`
	ModuleOneChapterThreeSessionThree string `json:"module_one_chapter_three_session_three"`
	// SetStatusChapterThreeModuleOne    int    `json:"set_status_chapter_three_module_one"`
}

// Training "Module-1" as ModuleOne is basic asset, which needs to be updated following the progress of the training itself.
type ModuleOne struct {
	ModuleOneID                    string                `json:"module_one_id"`
	ModuleOneTrainer               TrainerModuleOne      `json:"trainer_module_one"`
	ModuleOneTrainee               TraineeModuleOne      `json:"trainee_module_one"`
	ModuleOneTrainingCreatedDate   string                `json:"module_one_training_created_date"`
	ChapterOneModuleOne            ModuleOneChapterOne   `json:"module_one_chapter_one"`
	SetStatusChapterOneModuleOne   int                   `json:"set_status_chapter_one_module_one"`
	ChapterTwoModuleOne            ModuleOneChapterTwo   `json:"module_one_chapter_two"`
	SetStatusChapterTwoModuleOne   int                   `json:"set_status_chapter_two_module_one"`
	ChapterThreeModuleOne          ModuleOneChapterThree `json:"module_one_chapter_three"`
	SetStatusChapterThreeModuleOne int                   `json:"set_status_chapter_three_module_one"`
	TrainingTypeModuleOne          int                   `json:"training_type_module_one"`
	TheoreticalAssessmentModuleOne uint                  `json:"theoretical_assessment_module_one"`
	PracticalAssessmentModuleOne   uint                  `json:"practical_assessment_module_one"`
	AssessmentModuleOne            uint                  `json:"assessment_module_one"`
	TrainingStatusModuleOne        int                   `json:"training_status_module_one"`
	AssessmentsAttempts            int                   `json:"assessments_attempts"`
}
