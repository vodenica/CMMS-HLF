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
class ReadOperationalInvoiceWorkload extends WorkloadModuleBase {
  constructor() {
    super();

    this.txIndex = 0;

    // Note: here we are working with integers, but we could also work with floats
    this.costCode = [11111, 22222, 33333, 44444, 55555];

    this.invoiceOpsCost = [1500000, 175000, 200000, 225000, 250000];

    this.invoiceVariations = [10000, 12000, 15000, 17000, 18000, 20000];

    this.invoiceDowntimeDeduction = [10000, 11000, 12000, 15000, 18000];

    this.itemDescription = [
      "Item Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit."
    ];

    this.adjustmentCPI = [25000, 30000, 35000];

    this.invoiceOtherReductions = [
      1000,
      2000,
      3000,
      4000,
      5000,
      6000,
      7000,
      8000,
      9000,
      10000
    ];

    this.invoiceTotalCost = [
      10000,
      20000,
      30000,
      40000,
      50000,
      60000,
      70000,
      80000
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
      const docType = "invoice-ops";

      const invoice_type_ops =
        "operations-invoice" +
        `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;

      const invoice_ops_id = `INV-OPS-${this.workerIndex}_${i}`;

      let invoice_operational_cost = this.invoiceOpsCost[
        Math.floor(Math.random() * this.invoiceOpsCost.length)
      ];

      let variation_additional_expenses = this.invoiceVariations[
        Math.floor(Math.random() * this.invoiceVariations.length)
      ];

      let downtime_deduction = this.invoiceDowntimeDeduction[
        Math.floor(Math.random() * this.invoiceDowntimeDeduction.length)
      ];

      let cpi_adjustment = this.adjustmentCPI[
        Math.floor(Math.random() * this.adjustmentCPI.length)
      ];

      let other_agreed_reductions = this.invoiceOtherReductions[
        Math.floor(Math.random() * this.invoiceOtherReductions.length)
      ];

      console.log(
        `Worker ${this
          .workerIndex}: Creating Invoice for Operational Services ${invoice_ops_id}`
      );
      const request = {
        contractId: this.roundArguments.contractId,
        contractFunction: "NewInvoiceOperationalServices",
        invokerIdentity: "User1",
        contractArguments: [
          docType,
          invoice_type_ops,
          invoice_ops_id,

          invoice_operational_cost,
          variation_additional_expenses,
          downtime_deduction,
          cpi_adjustment,
          other_agreed_reductions
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
      contractFunction: "GetInvoiceOpsByID",
      invokerIdentity: "User1",
      // Pass the correct value of purchase_order_id that was created in the initializeWorkloadModule function
      contractArguments: [`INV-OPS-${this.workerIndex}_${randomId}`],

      readOnly: true
    };

    await this.sutAdapter.sendRequests(myArgs);
  }
}

function createWorkloadModule() {
  return new ReadOperationalInvoiceWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;
