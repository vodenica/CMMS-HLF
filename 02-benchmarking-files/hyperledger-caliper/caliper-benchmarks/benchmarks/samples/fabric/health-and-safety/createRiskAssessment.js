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

// The workload script in JavaScript for the benchmark testing in Hyperledger Caliper is shown below:

'use strict';

const { WorkloadModuleBase } = require('@hyperledger/caliper-core');

class CreateRiskAssessmentWorkload extends WorkloadModuleBase {
    /**
     * Initializes the workload module instance.
     */
    constructor() {
        super();
        this.txIndex = 0;
        this.activity = ['Working on towers', 'Grip maintenance', 'Cabins maintenance', 'Working at the tyre conveyor are', 'Electrical works'];
        this.creator = ['Alice', 'Bob', 'Charley', 'Dave', 'Eleanore'];
        this.assessor = ['Fox', 'Gina', 'Honey', 'Ivy', 'John'];
        this.approver = ['Karl', 'Lily', 'Mia', 'Nina', 'Oscar'];
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
async initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext) {
    await super.initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext);
}

async submitTransaction() {
    this.txIndex++;

    let docType = 'risk-assessment-' + this.workerIndex + this.txIndex.toString();
    const risk_assessment_id = 'risk-assessment-id-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
    let risk_assessment_date = 'risk-assessment-date-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
    let risk_assessment_date_next_review = 'risk-assessment-date-next-review-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
    let risk_assessment_activity = this.activity[Math.floor(Math.random() * this.activity.length)];
    let risk_assessment_created_by = this.creator[Math.floor(Math.random() * this.creator.length)];
    let risk_assessment_assessed_by = this.assessor[Math.floor(Math.random() * this.assessor.length)];
    let risk_assessment_approved_by = this.approver[Math.floor(Math.random() * this.approver.length)];
    let risk_assessment_hazard_list_one = 'risk-assessment-hazard-list-one-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
    let risk_assessment_hazard_list_two = 'risk-assessment-hazard-list-two-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
    let risk_assessment_hazard_list_three = 'risk-assessment-hazard-list-three-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
    let risk_assessment_hazard_list_four = 'risk-assessment-hazard-list-four-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
    let risk_assessment_hazard_list_five = 'risk-assessment-hazard-list-five-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;

/*
    let event_type = this.type[Math.floor(Math.random() * this.type.length)];
    let event_description = this.description[Math.floor(Math.random() * this.description.length)];
    let event_start = 'event-started_' + this.workerIndex + this.txIndex.toString();
    let event_ends   = 'event-ended_' + this.workerIndex + this.txIndex.toString();
    */

    const request = {
        contractId: this.roundArguments.contractId,
        contractFunction: 'CreateRiskAssessment',
        invokerIdentity: 'User1',
        contractArguments: [docType,
            risk_assessment_id,
            risk_assessment_date,
            risk_assessment_date_next_review,
            risk_assessment_activity,
            risk_assessment_created_by,
            risk_assessment_assessed_by,
            risk_assessment_approved_by,
            risk_assessment_hazard_list_one,
            risk_assessment_hazard_list_two,
            risk_assessment_hazard_list_three,
            risk_assessment_hazard_list_four,
            risk_assessment_hazard_list_five],
        readOnly: false
    };

    console.info(this.txIndex);
    await this.sutAdapter.sendRequests(request);
}
}

function createWorkloadModule() {
return new CreateRiskAssessmentWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;





