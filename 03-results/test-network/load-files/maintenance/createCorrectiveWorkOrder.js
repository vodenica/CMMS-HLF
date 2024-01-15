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

class CreateCorrectiveWorkOrderWorkload extends WorkloadModuleBase {
  /**
     * Initializes the workload module instance.
     */
  constructor() {
    super();
    this.txIndex = 0;
    this.plannedLabourSupervisors = [
      "Maintenance Supervisor",
      "Operations Supervisor"
    ];
    this.plannedLabourTechnicians = [
      "Maintenance Technician Tier 1",
      "Maintenance Technician Tier 2",
      "Maintenance Technician Tier 3"
    ];
    this.plannedLabourOperators = [
      "System Operator Tier 1",
      "System Operator Tier 2",
      "System Operator Tier 3"
    ];
    this.processSteps = [
      "Process Step N: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer pellentesque convallis est a commodo. Nulla sit amet vehicula justo. Sed at arcu placerat, dignissim ante.",
      "Process Step N: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer pellentesque convallis est a commodo. Nulla sit amet vehicula justo. Sed at arcu placerat, dignissim ante.",
      "Process Step N: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer pellentesque convallis est a commodo. Nulla sit amet vehicula justo. Sed at arcu placerat, dignissim ante.",
      "Process Step N: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer pellentesque convallis est a commodo. Nulla sit amet vehicula justo. Sed at arcu placerat, dignissim ante.",
      "Process Step N: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer pellentesque convallis est a commodo. Nulla sit amet vehicula justo. Sed at arcu placerat, dignissim ante."
    ];
    this.maintenanceParts = [
      "Maintenance Part 1",
      "Maintenance Part 2",
      "Maintenance Part 3",
      "Maintenance Part 4",
      "Maintenance Part 5",
      "Maintenance Part 6",
      "Maintenance Part 7",
      "Maintenance Part 8",
      "Maintenance Part 9",
      "Maintenance Part 10"
    ];
    this.description = [
      "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer pellentesque convallis est a commodo. Nulla sit amet vehicula justo. Sed at arcu placerat, dignissim ante."
    ];
    this.healthandsafety = [
      "Health & Safety Instruction: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut at molestie tortor. Phasellus iaculis id enim ac bibendum. Nullam gravida quam eget lacus dictum, vel ullamcorper eros condimentum. Aenean laoreet elit ut magna gravida dapibus. Ut scelerisque sit amet justo eget auctor. Donec fringilla, eros pharetra molestie tristique, urna felis maximus turpis, sed interdum velit turpis ut leo. Nam commodo, est sed congue malesuada, neque lorem elementum odio, ut tempus dui neque non ex. Aliquam semper faucibus ante, eget feugiat erat ullamcorper a. Aenean sagittis lacus mi, a egestas nibh luctus et. Sed congue nisl et nunc tincidunt, eu tincidunt nunc vehicula. Donec dui lorem, vehicula ac lectus at, venenatis mattis odio."
    ];
    this.creator = ["Alice", "Bob", "Charlie", "Dave", "Eve", "Frank"];
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

    let corrmtewo = "docType-corrmtewo" + this.worker + this.txIndex.toString();

    const corr_work_order_id =
      "corrective-work-order-ID-" +
      `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;

    let corr_work_order_type = "Corrective Work Order";

    let corr_work_order_description = this.description[
      Math.floor(Math.random() * this.description.length)
    ];

    let corr_general_instructions = this.healthandsafety[
      Math.floor(Math.random() * this.healthandsafety.length)
    ];

    const plannedLabour = {
      maintenance_supervisor: this.plannedLabourSupervisors[
        Math.floor(Math.random() * this.plannedLabourSupervisors.length)
      ],
      maintenance_technician_one: this.plannedLabourTechnicians[
        Math.floor(Math.random() * this.plannedLabourTechnicians.length)
      ],
      maintenance_technician_two: this.plannedLabourTechnicians[
        Math.floor(Math.random() * this.plannedLabourTechnicians.length)
      ],
      maintenance_technician_three: this.plannedLabourTechnicians[
        Math.floor(Math.random() * this.plannedLabourTechnicians.length)
      ],
      system_operator: this.plannedLabourOperators[
        Math.floor(Math.random() * this.plannedLabourOperators.length)
      ]
    };

    const correctiveMaintenanceProcessSteps = {
      maintenance_process_step_one: this.processSteps[
        Math.floor(Math.random() * this.processSteps.length)
      ],
      maintenance_process_step_two: this.processSteps[
        Math.floor(Math.random() * this.processSteps.length)
      ],
      maintenance_process_step_three: this.processSteps[
        Math.floor(Math.random() * this.processSteps.length)
      ],
      maintenance_process_step_four: this.processSteps[
        Math.floor(Math.random() * this.processSteps.length)
      ],
      maintenance_process_step_five: this.processSteps[
        Math.floor(Math.random() * this.processSteps.length)
      ]
    };

    const maintenanceParts = {
      maintenance_part_one: this.maintenanceParts[
        Math.floor(Math.random() * this.maintenanceParts.length)
      ],
      maintenance_part_two: this.maintenanceParts[
        Math.floor(Math.random() * this.maintenanceParts.length)
      ],
      maintenance_part_three: this.maintenanceParts[
        Math.floor(Math.random() * this.maintenanceParts.length)
      ],
      maintenance_part_four: this.maintenanceParts[
        Math.floor(Math.random() * this.maintenanceParts.length)
      ],
      maintenance_part_five: this.maintenanceParts[
        Math.floor(Math.random() * this.maintenanceParts.length)
      ],
      maintenance_part_six: this.maintenanceParts[
        Math.floor(Math.random() * this.maintenanceParts.length)
      ],
      maintenance_part_seven: this.maintenanceParts[
        Math.floor(Math.random() * this.maintenanceParts.length)
      ],
      maintenance_part_eight: this.maintenanceParts[
        Math.floor(Math.random() * this.maintenanceParts.length)
      ],
      maintenance_part_nine: this.maintenanceParts[
        Math.floor(Math.random() * this.maintenanceParts.length)
      ],
      maintenance_part_ten: this.maintenanceParts[
        Math.floor(Math.random() * this.maintenanceParts.length)
      ]
    };

    let corr_work_order_creator = this.creator[Math.floor(Math.random() * this.creator.length)];

    let corr_work_order_tiemstamp = Date.now();

    let corr_condition = "not-good";

    let corr_validation = "not-validated";

    const request = {
      contractId: this.roundArguments.contractId,
      contractFunction: "CreateCorrectiveWorkOrder",
      invokerIdentity: "User1",
      contractArguments: [
        corr_work_order_id,
        corr_work_order_type,
        corr_work_order_description,
        corr_general_instructions,
        JSON.stringify(plannedLabour),
        JSON.stringify(correctiveMaintenanceProcessSteps),
        JSON.stringify(maintenanceParts),
        corr_work_order_creator,
        corr_work_order_tiemstamp,
        corr_condition,
        corr_validation
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
  return new CreateCorrectiveWorkOrderWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;

