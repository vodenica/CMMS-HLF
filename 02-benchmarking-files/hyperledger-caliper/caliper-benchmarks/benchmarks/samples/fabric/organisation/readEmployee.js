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
class ReadEmployeeWorkload extends WorkloadModuleBase {
  constructor() {
    super();
    this.txIndex = 0;

    // The names and surnames were automatically generated using https://www.mockaroo.com/
    this.name = [
      "John",
      "Mary",
      "James",
      "Patricia",
      "Robert",
      "Jennifer",
      "Michael",
      "Linda",
      "William",
      "Elizabeth"
    ];

    this.lastName = [
      "Smith",
      "Jones",
      "Williams",
      "Taylor",
      "Davies",
      "Brown",
      "Wilson",
      "Evans",
      "Thomas",
      "Johnson"
    ];

    this.gender = ["female", "male"];

    this.positionEntry = ["System Operator", "Maintenance Technician"];

    this.position = [
      "Maintenance Technician Tier 1",
      "Maintenance Technician Tier 2",
      "Maintenance Technician Tier 3",
      "Maintenance Supervisor",
      "Maintenance Manager",
      "Operations Manager",
      "Operations Supervisor",
      "Operations Technician"
    ];

    this.salary = [15000, 20000, 25000, 30000, 35000, 40000, 45000, 50000];

    this.contractSignedDate = [
      "2019-01-01",
      "2019-02-01",
      "2019-03-01",
      "2019-04-01",
      "2019-05-01",
      "2019-06-01"
    ];

    this.startingDate = [
      "2019-01-01",
      "2019-02-01",
      "2019-03-01",
      "2019-04-01",
      "2019-05-01",
      "2019-06-01",
      "2019-07-01"
    ];

    this.birthday = [
      "1990-01-01",
      "1991-02-01",
      "1992-03-01",
      "1993-04-01",
      "1994-05-01",
      "1995-06-01",
      "1996-07-01"
    ];

    this.address = [
      "Address 1",
      "Address 2",
      "Address 3",
      "Address 4",
      "Address 5",
      "Address 6",
      "Address 7",
      "Address 8",
      "Address 9",
      "Address 10"
    ];

    this.cellNumber = [
      "+27 11 111 1111",
      "+27 22 222 2222",
      "+27 33 333 3333",
      "+27 44 444 4444",
      "+27 55 555 5555",
      "+27 66 666 6666",
      "+27 77 777 7777",
      "+27 88 888 8888",
      "+27 99 999 9999",
      "+27 10 101 1010"
    ];

    this.employeeReview = [
      "Employee Review 1",
      "Employee Review 2",
      "Employee Review 3",
      "Employee Review 4",
      "Employee Review 5",
      "Employee Review 6",
      "Employee Review 7",
      "Employee Review 8",
      "Employee Review 9",
      "Employee Review 10"
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

      const CompanySite = "Macao S.A.R., China";

      const ID = `${this.workerIndex}_${i}`;

      const Name = this.name[Math.floor(Math.random() * this.name.length)];

      const LastName = this.lastName[
        Math.floor(Math.random() * this.lastName.length)
      ];

      const Gender = this.gender[Math.round(Math.random())];

      const PositionEntry = this.positionEntry[
        Math.floor(Math.random() * this.positionEntry.length)
      ];

      const CurrentPosition = this.position[
        Math.floor(Math.random() * this.position.length)
      ];

      const LevelPT = "Full-time";

      const Birthday = this.birthday[
        Math.floor(Math.random() * this.birthday.length)
      ];

      const ContractSigned = this.contractSignedDate[
        Math.floor(Math.random() * this.contractSignedDate.length)
      ];

      const StartingDate = this.startingDate[
        Math.floor(Math.random() * this.startingDate.length)
      ];

      const YearsInService = Math.floor(Math.random() * 2);

      const EmployeeReview = this.employeeReview[
        Math.floor(Math.random() * this.employeeReview.length)
      ];

      const CompletedCustomTraining = "Completed Custom Training";

      const Salary = this.salary[
        Math.floor(Math.random() * this.salary.length)
      ];

      const Address = this.address[
        Math.floor(Math.random() * this.address.length)
      ];

      const Cell = this.cellNumber[
        Math.floor(Math.random() * this.cellNumber.length)
      ];

      const WorkVisa = "Work Visa Yes/No";

      console.log(
        `Worker ${this.workerIndex}: Adding new employees ${ID}`
      );
      const request = {
        contractId: this.roundArguments.contractId,
        contractFunction: "AddNewEmployee",
        invokerIdentity: "User1",
        contractArguments: [
          CompanySite,
          ID,
          Name,
          LastName,
          Gender,
          PositionEntry,
          CurrentPosition,
          LevelPT,
          Birthday,
          ContractSigned,
          StartingDate,
          YearsInService,
          EmployeeReview,
          CompletedCustomTraining,
          Salary,
          Address,
          Cell,
          WorkVisa
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
      contractFunction: "ReadEmployee",
      invokerIdentity: "User1",
      // Pass the correct value of corr_work_order_id that was created in the initializeWorkloadModule function
      contractArguments: [`${this.workerIndex}_${randomId}`],
      readOnly: true
    };

    await this.sutAdapter.sendRequests(myArgs);
  }
}

function createWorkloadModule() {
  return new ReadEmployeeWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;
