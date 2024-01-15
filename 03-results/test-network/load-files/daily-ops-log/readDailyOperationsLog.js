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
class QueryDailyOperationalLogsWorkload extends WorkloadModuleBase {
  constructor() {
    super();
    this.txIndex = 0;
    this.creator = ["Alice", "Bob", "Charley", "Dave", "Eleanore"];
    this.weather = ["Sunny", "Rainy", "Cloudy", "Snowy"];
    this.humidity = ["95%", "96%", "97%", "98%", "99%"];
    this.temperature = ["+31 C", "+32 C", "+33 C", "+34 C", "+35 C"];
    this.comments = [
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus sit.",
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus sit.",
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus sit."
    ];
    this.maintenanceTechnician = [
      "Maintenance Technician Tier 1",
      "Maintenance Technician Tier 2",
      "Maintenance Technician Tier 3"
    ];
    this.systemOperator = [
      "System Operator Tier 1",
      "System Operator Tier 2",
      "System Operator Tier 3"
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
      //  const daily_ops_log_id = `${this.workerIndex}_${i}`;

      const daily_ops_log_id = `${this.workerIndex}_${i}`;

      let owner = this.creator[Math.floor(Math.random() * this.creator.length)];

      const weatherDriveStation = {
        weather_drive_station_condition: this.weather[
          Math.floor(Math.random() * this.weather.length)
        ],
        weather_drive_station_temperature: this.temperature[
          Math.floor(Math.random() * this.temperature.length)
        ],
        weather_drive_station_humidity: this.humidity[
          Math.floor(Math.random() * this.humidity.length)
        ]
      };

      const weatherReturnStation = {
        weather_return_station_condition: this.weather[
          Math.floor(Math.random() * this.weather.length)
        ],
        weather_return_station_temperature: this.temperature[
          Math.floor(Math.random() * this.temperature.length)
        ],
        weather_return_station_humidity: this.humidity[
          Math.floor(Math.random() * this.humidity.length)
        ]
      };

      const personnelOnDutyDriveStation = {
        personnel_on_duty_drive_station_shift_manager: "Shift Manager",
        personnel_on_duty_drive_station_maintenance_technician_one: this
          .maintenanceTechnician[
          Math.floor(Math.random() * this.maintenanceTechnician.length)
        ],
        personnel_on_duty_drive_station_maintenance_technician_two: this
          .maintenanceTechnician[
          Math.floor(Math.random() * this.maintenanceTechnician.length)
        ],
        personnel_on_duty_drive_station_system_operator_one: this
          .systemOperator[
          Math.floor(Math.random() * this.systemOperator.length)
        ],
        personnel_on_duty_drive_station_system_operator_two: this
          .systemOperator[
          Math.floor(Math.random() * this.systemOperator.length)
        ]
      };

      const personnelOnDutyReturnStation = {
        personnel_on_duty_return_station_maintenance_technician: this
          .maintenanceTechnician[
          Math.floor(Math.random() * this.maintenanceTechnician.length)
        ],
        personnel_on_duty_return_station_system_operator: this.systemOperator[
          Math.floor(Math.random() * this.systemOperator.length)
        ]
      };

      let operations_start =
        "operations-start-at-" +
        `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;

      let operations_end =
        "operations-end-at" +
        `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;

      let operation_hours = Math.floor(Math.random() * 10);

      let number_of_carriers = Math.floor(Math.random() * 10);

      let number_of_passengers = Math.floor(Math.random() * 100);

      let total_operating_hours = Math.floor(Math.random() * 10000);

      let additional_comments =
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus sit. -" +
        `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;

      console.log(
        `Worker ${this
          .workerIndex}: Creating daily operations log ${daily_ops_log_id}`
      );
      const request = {
        contractId: this.roundArguments.contractId,
        contractFunction: "CreateNewDailyOperationsLog",
        invokerIdentity: "User1",
        contractArguments: [
          //"dailyopslog",
          daily_ops_log_id,
          owner,
          JSON.stringify(weatherDriveStation),
          JSON.stringify(weatherReturnStation),
          JSON.stringify(personnelOnDutyDriveStation),
          JSON.stringify(personnelOnDutyReturnStation),
          operations_start,
          operations_end,
          operation_hours,
          number_of_carriers,
          number_of_passengers,
          total_operating_hours,
          additional_comments
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
      contractFunction: "ReadDailyOperationsLog",
      invokerIdentity: "User1",
      // Pass the correct value of daily_ops_log_id that was created in the initializeWorkloadModule function
      contractArguments: [`${this.workerIndex}_${randomId}`],
      readOnly: true
    };

    await this.sutAdapter.sendRequests(myArgs);
  }

}

function createWorkloadModule() {
  return new QueryDailyOperationalLogsWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;
