/*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

// The workload script for ReadDailyOperationsLog in JavaScript for the benchmark testing in Hyperledger Caliper
// is shown below:

"use strict";

const { WorkloadModuleBase } = require("@hyperledger/caliper-core");

/**
 * Workload module for the benchmark round.
 */
class ReadTrainingModuleOneWorkload extends WorkloadModuleBase {
  constructor() {
    super();
    this.txIndex = 0;

    // The names and surnames were automatically generated using https://www.mockaroo.com/
    this.trainerModuleOneName = [
      "John", "Mary", "James", "Patricia", "Robert", "Jennifer", "Michael", "Linda", "William", "Elizabeth"
    ];
    this.trainerModuleOneSurname = [
      "Smith", "Jones", "Williams", "Taylor", "Davies", "Brown", "Wilson", "Evans", "Thomas", "Johnson"
    ];
    // The names and surnames were automatically generated using https://www.mockaroo.com/
    this.traineeModuleOneName = [
      "David", "Susan", "Richard", "Karen", "Joseph", "Nancy", "Charles", "Margaret", "Christopher", "Lisa"
    ];
    this.traineeModuleOneSurname = [
      "Miller", "Wilson", "Moore", "Taylor", "Anderson", "Thomas", "Jackson", "White", "Harris", "Martin"
    ];

    this.moduleOneChapterOne = [
      "Session A: Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
      "Session B: Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
      "Session C: Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
    ];

    this.moduleOneChapterTwo = [
      "Session D: Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
      "Session E: Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
      "Session F: Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
    ];

    this.moduleOneChapterThree = [
      "Session G: Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
      "Session H: Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
      "Session J: Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
    ];
  }

  async initializeWorkloadModule(
    workerIndex,
    totalWorkers,
    roundIndex,
    roundArguments,
    sutAdapter,
    sutContext
  ) {
    await super.initializeWorkloadModule(
      workerIndex,
      totalWorkers,
      roundIndex,
      roundArguments,
      sutAdapter,
      sutContext
    );

    for (let i = 0; i < this.roundArguments.assets; i++) {

      const module_one_id = `${this.workerIndex}_${i}`;


      const moduleOneTrainer = {
        trainer_module_one_id: "trainer-module-one-ID-" + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`,
        trainer_module_one_name: this.trainerModuleOneName[
          Math.floor(Math.random() * this.trainerModuleOneName.length)
        ],
        trainer_module_one_surname: this.trainerModuleOneSurname[
          Math.floor(Math.random() * this.trainerModuleOneSurname.length)
        ]
      };
  
      const moduleOneTrainee = {
        trainee_module_one_id: "trainee-module-one-ID-" + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`,
        trainee_module_one_name: this.traineeModuleOneName[
          Math.floor(Math.random() * this.traineeModuleOneName.length)
        ],
        trainee_module_one_surname: this.traineeModuleOneSurname[
          Math.floor(Math.random() * this.traineeModuleOneSurname.length)
        ]
      };
  
      const moduleOneChapterOne = {
        module_one_chapter_one_session_one: this.moduleOneChapterOne[
          Math.floor(Math.random() * this.moduleOneChapterOne.length)
        ],
        module_one_chapter_one_session_two: this.moduleOneChapterOne[
          Math.floor(Math.random() * this.moduleOneChapterOne.length)
        ],
        module_one_chapter_one_session_three: this.moduleOneChapterOne[
          Math.floor(Math.random() * this.moduleOneChapterOne.length)
        ]
      };
  
      const moduleOneChapterTwo = {
        module_one_chapter_two_session_one: this.moduleOneChapterTwo[
          Math.floor(Math.random() * this.moduleOneChapterTwo.length)
        ],
        module_one_chapter_two_session_two: this.moduleOneChapterTwo[
          Math.floor(Math.random() * this.moduleOneChapterTwo.length)
        ],
        module_one_chapter_two_session_three: this.moduleOneChapterTwo[
          Math.floor(Math.random() * this.moduleOneChapterTwo.length)
        ]
      };
  
      const moduleOneChapterThree = {
        module_one_chapter_three_session_one: this.moduleOneChapterThree[
          Math.floor(Math.random() * this.moduleOneChapterThree.length)
        ],
        module_one_chapter_three_session_two: this.moduleOneChapterThree[
          Math.floor(Math.random() * this.moduleOneChapterThree.length)
        ],
        module_one_chapter_three_session_three: this.moduleOneChapterThree[
          Math.floor(Math.random() * this.moduleOneChapterThree.length)
        ]
      };
  
      let module_one_training_created_date = Date.now();
  
      let set_status_chapter_one_module_one = Math.floor(Math.random() * 1);
  
      let set_status_chapter_two_module_one = Math.floor(Math.random() * 1);
  
      let set_status_chapter_three_module_one = Math.floor(Math.random() * 1);
  
      let training_type_module_one = Math.floor(Math.random() * 1);
  
      let theoretical_assessment_module_one = Math.floor(Math.random() * 100);
  
      let practical_assessment_module_one = Math.floor(Math.random() * 100);
  
      let assessment_module_one = Math.floor(Math.random() * 100);
  
      let training_status_module_one = Math.floor(Math.random() * 1);
  
      let assessments_attempts = Math.floor(Math.random() * 1);

      console.log(
        `Worker ${this
          .workerIndex}: Creating Training Module One ${module_one_id}`
      );
      const request = {
        contractId: this.roundArguments.contractId,
        contractFunction: "CreateNewModuleOne",
        invokerIdentity: "User1",
        contractArguments: [
          module_one_id,
          JSON.stringify(moduleOneTrainer),
          JSON.stringify(moduleOneTrainee),
          JSON.stringify(moduleOneChapterOne),
          JSON.stringify(moduleOneChapterTwo),
          JSON.stringify(moduleOneChapterThree),
          module_one_training_created_date,
          set_status_chapter_one_module_one,
          set_status_chapter_two_module_one,
          set_status_chapter_three_module_one,
          training_type_module_one,
          theoretical_assessment_module_one,
          practical_assessment_module_one,
          assessment_module_one,
          training_status_module_one,
          assessments_attempts
        ],
        readOnly: false
      };

      await this.sutAdapter.sendRequests(request);
    }

    this.limitIndex = this.roundArguments.assets;
  }

  async submitTransaction() {
    const randomId = Math.floor(Math.random() * this.roundArguments.assets);
    const myArgs = {
      contractId: this.roundArguments.contractId,
      contractFunction: "GetModuleOneRecords",
      invokerIdentity: "User1",
      // Pass the correct value of corr_work_order_id that was created in the initializeWorkloadModule function
      contractArguments: [`${this.workerIndex}_${randomId}`],
      readOnly: true
    };

    await this.sutAdapter.sendRequests(myArgs);
  }
}

function createWorkloadModule() {
  return new ReadTrainingModuleOneWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;
