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

class CreateAccidentIncidentReport extends WorkloadModuleBase {
    /**
     * Initializes the workload module instance.
     */
    constructor() {
        super();
        this.txIndex = 0;
        this.subject = ['Carrier collision causes a prolong downtime during public operating hours.', 'Employee slips on wet floor, resulting in a broken arm.',
        'Team member falls from scaffolding, sustaining serious injuries.', 'Chemical spill in Drive Room leads to evacuation of building.'];
        this.classification = ['Near Miss', 'First Aid', 'Medical Treatment', 'Lost Time Injury', 'Fatality'];
        this.description = ['Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod, diam id aliquam ultricies, nunc ipsum aliquet nunc, eget. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod, diam id aliquam ultricies, nunc ipsum aliquet nunc, eget.'];
        this.creator = ['Alice', 'Bob', 'Charley', 'Dave', 'Eleanore'];
        this.validator = ['Fox', 'Gina', 'Honey', 'Ivy', 'John'];
        this.peopleInvolved = [1, 2, 3, 4, 5];
        this.injured = ['Karl', 'Lily', 'Mia', 'Nina', 'Oscar'];
        this.witness = ['Peter', 'Quinn', 'Rita', 'Sam', 'Tina'];
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

    let docType = 'accident-incident-report-' + this.workerIndex + this.txIndex.toString();
    const accident_incident_report_id = 'accident-incident-report-id-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
    let subject_accident_incident_report_id = 'accident-incident-report-subject-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}_${this.subject}`;
    let accident_incident_report_date = 'accident-incident-report-date-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
    let accident_incident_report_time_start = 'accident-incident-report-time-start-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
    let accident_incident_report_time_end = 'accident-incident-report-time-end-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
    let accident_incident_report_location = 'accident-incident-report-location-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
    let accident_incident_report_hs_aspects = 'accident-incident-report-hs-aspects-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
    let accident_incident_report_classification = 'accident-incident-report-classification-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}_${this.classification}`;
    let accident_incident_report_description = 'accident-incident-report-description-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}_${this.description}`;
    let accident_incident_report_immediate_action = 'accident-incident-report-immediate-action-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
    let accident_incident_report_follow_up_action = 'accident-incident-report-follow-up-action-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
    let accident_incident_report_created_by = 'accident-incident-report-created-by-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}_${this.creator}`;
    let accident_incident_report_validated_by = 'accident-incident-report-validated-by-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}_${this.validator}`;
    let accident_incident_report_status = 'accident-incident-report-status-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
    
    // the following field shall be pass as an integer
    let accident_incident_report_number_of_people_involved = this.peopleInvolved[Math.floor(Math.random() * this.peopleInvolved.length)];

    let accident_incident_report_person_injured = 'accident-incident-report-person-injured-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}_${this.injured}`;
    let accident_incident_report_witnesses = 'accident-incident-report-witness-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}_${this.witness}`;
    let accident_incident_report_event_overview = 'accident-incident-report-event-overview-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;
    let accident_incident_report_line_of_communication = 'accident-incident-report-line-of-communication-' + `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;

    const request = {
        contractId: this.roundArguments.contractId,
        contractFunction: 'CreateAccidentIncidentReport',
        invokerIdentity: 'User1',
        contractArguments: [
            docType,
            accident_incident_report_id,
            subject_accident_incident_report_id,
            accident_incident_report_date,
            accident_incident_report_time_start,
            accident_incident_report_time_end,
            accident_incident_report_location,
            accident_incident_report_hs_aspects,
            accident_incident_report_classification,
            accident_incident_report_description,
            accident_incident_report_immediate_action,  
            accident_incident_report_follow_up_action,
            accident_incident_report_created_by,
            accident_incident_report_validated_by,
            accident_incident_report_status,
            accident_incident_report_number_of_people_involved,
            accident_incident_report_person_injured,
            accident_incident_report_witnesses,
            accident_incident_report_event_overview,
            accident_incident_report_line_of_communication],
        readOnly: false
    };

    console.info(this.txIndex);
    await this.sutAdapter.sendRequests(request);
}
}

function createWorkloadModule() {
    return new CreateAccidentIncidentReport();
}

/*
return new CreateRiskAssessmentWorkload();
}
*/

module.exports.createWorkloadModule = createWorkloadModule;





