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

// The workload script in JavaScript for the benchmark testing in Hyperledger Caliper
// is shown below:

"use strict";

const { WorkloadModuleBase } = require("@hyperledger/caliper-core");

class CreateDailyOperationLogWorkload extends WorkloadModuleBase {
  /**
     * Initializes the workload module instance.
     */
  constructor() {
    super();
    this.txIndex = 0;
    this.creator = ['Alice', 'Bob', 'Charley', 'Dave', 'Eleanore'];
    this.weather = ['Sunny', 'Rainy', 'Cloudy', 'Snowy'];
    this.humidity = ['95%', '96%', '97%', '98%', '99%'];
    this.temperature = ['+31 C', '+32 C', '+33 C', '+34 C', '+35 C'];
    this.comments = ['Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus sit.', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus sit.'];
    this.maintenanceTechnician = ['Maintenance Technician Tier 1', 'Maintenance Technician Tier 2', 'Maintenance Technician Tier 3'];
    this.systemOperator = ['System Operator Tier 1', 'System Operator Tier 2', 'System Operator Tier 3'];
  }

  /**
     * Initialize the workload module with the given parameters.
     * @param {number} workerIndex The 0-based index of the worker instantiating the workload module.
     * @param {number} totalWorkers The total number of workers participating in the round.
     * @param {number} roundIndex The 0-based index of the currently executing round.
     * @param {Object} roundArguments The user-provided arguments for the round from the benchmark configuration file.
     * @param {ConnectorBase} sutAdapter The adapter of the underlying SUT.
     * @param {Object} sutContext The custom context object provided by the SUT adapter.
     * @async
     */
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
  }

  async submitTransaction() {
	this.txIndex++;
  
	let dailyopslog =
	  "dailyopslog-" + this.worker + this.txIndex.toString();
  
	const daily_ops_log_id =
	  "daily-ops-log-ID-" +
	  `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
  
	let owner = this.creator[Math.floor(Math.random() * this.creator.length)];

  const weatherDriveStation = {
    weather_drive_station_condition: this.weather[Math.floor(Math.random() * this.weather.length)],
    weather_drive_station_temperature: this.temperature[Math.floor(Math.random() * this.temperature.length)],
    weather_drive_station_humidity: this.humidity[Math.floor(Math.random() * this.humidity.length)]
  };

  const weatherReturnStation = {
    weather_return_station_condition: this.weather[Math.floor(Math.random() * this.weather.length)],
    weather_return_station_temperature: this.temperature[Math.floor(Math.random() * this.temperature.length)],
    weather_return_station_humidity: this.humidity[Math.floor(Math.random() * this.humidity.length)]
  };

  const personnelOnDutyDriveStation = {
    personnel_on_duty_drive_station_shift_manager: "Shift Manager",
    personnel_on_duty_drive_station_maintenance_technician_one: this.maintenanceTechnician[Math.floor(Math.random() * this.maintenanceTechnician.length)],
    personnel_on_duty_drive_station_maintenance_technician_two: this.maintenanceTechnician[Math.floor(Math.random() * this.maintenanceTechnician.length)],
    personnel_on_duty_drive_station_system_operator_one: this.systemOperator[Math.floor(Math.random() * this.systemOperator.length)],
    personnel_on_duty_drive_station_system_operator_two: this.systemOperator[Math.floor(Math.random() * this.systemOperator.length)]
  };

  const personnelOnDutyReturnStation = {
    personnel_on_duty_return_station_maintenance_technician: this.maintenanceTechnician[Math.floor(Math.random() * this.maintenanceTechnician.length)],
    personnel_on_duty_return_station_system_operator: this.systemOperator[Math.floor(Math.random() * this.systemOperator.length)]
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
  
	const request = {
	  contractId: this.roundArguments.contractId,
	  contractFunction: "CreateNewDailyOperationsLog",
	  invokerIdentity: "User1",
	  contractArguments: [
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
  
	console.info(this.txIndex);
	await this.sutAdapter.sendRequests(request);
  }

    async cleanupWorkloadModule() {
    // Do nothing
  }
}

function createWorkloadModule() {
  return new CreateDailyOperationLogWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;

// Path: benchmarks/samples/fabric/daily-ops-log/createDailyOperationsLog.js
// Compare this snippet from benchmarks/samples/fabric/daily-ops-log/readDailyOperationsLog.js:
