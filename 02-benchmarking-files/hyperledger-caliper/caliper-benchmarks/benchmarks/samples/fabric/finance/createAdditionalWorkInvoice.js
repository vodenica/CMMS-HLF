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

class CreatePurchaseOrderWorkload extends WorkloadModuleBase {
  /**
     * Initializes the workload module instance.
     */
  constructor() {
    super();

    this.txIndex = 0;

    this.costCode = [11111, 22222, 33333, 44444, 55555];

    // Note: here we are working with integers, but we could also work with floats
    this.itemCost = [100, 200, 300, 400, 500, 600, 700, 800, 900, 1000];

    this.itemQuantity = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

    this.itemTotalCost = [
      10000,
      20000,
      30000,
      40000,
      50000,
      60000,
      70000,
      80000
    ];

    this.invoiceVariations = [10000, 12000, 15000, 17000, 18000, 20000];
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

    const docType = "invoice-work";

    const invoice_type_work =
      "invoice-work" +
      `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;

    const invoice_work_id =
      "Invoice Work ID-" +
      `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;

    const invoice_description =
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.";

    const invoice_work_cost_code = this.costCode[
      Math.floor(Math.random() * this.costCode.length)
    ];

    const variation_additional_expenses_add_work = this.invoiceVariations[
      Math.floor(Math.random() * this.invoiceVariations.length)
    ];

    const invoiceItemOne = {
      invoice_item_one_qty: this.itemQuantity[
        Math.floor(Math.random() * this.itemQuantity.length)
      ],
      invoice_item_one_item_cost: this.itemCost[
        Math.floor(Math.random() * this.itemCost.length)
      ],
      invoice_item_one_cost: this.itemTotalCost[
        Math.floor(Math.random() * this.itemTotalCost.length)
      ]
    };

    const invoiceItemTwo = {
      invoice_item_two_qty: this.itemQuantity[
        Math.floor(Math.random() * this.itemQuantity.length)
      ],
      invoice_item_two_item_cost: this.itemCost[
        Math.floor(Math.random() * this.itemCost.length)
      ],
      invoice_item_two_cost: this.itemTotalCost[
        Math.floor(Math.random() * this.itemTotalCost.length)
      ]
    };

    const invoiceItemThree = {
      invoice_item_three_qty: this.itemQuantity[
        Math.floor(Math.random() * this.itemQuantity.length)
      ],
      invoice_item_three_item_cost: this.itemCost[
        Math.floor(Math.random() * this.itemCost.length)
      ],
      invoice_item_three_cost: this.itemTotalCost[
        Math.floor(Math.random() * this.itemTotalCost.length)
      ]
    };

    const invoiceItemFour = {
      invoice_item_four_qty: this.itemQuantity[
        Math.floor(Math.random() * this.itemQuantity.length)
      ],
      invoice_item_four_item_cost: this.itemCost[
        Math.floor(Math.random() * this.itemCost.length)
      ],
      invoice_item_four_cost: this.itemTotalCost[
        Math.floor(Math.random() * this.itemTotalCost.length)
      ]
    };

    const invoiceItemFive = {
      invoice_item_five_qty: this.itemQuantity[
        Math.floor(Math.random() * this.itemQuantity.length)
      ],
      invoice_item_five_item_cost: this.itemCost[
        Math.floor(Math.random() * this.itemCost.length)
      ],
      invoice_item_five_cost: this.itemTotalCost[
        Math.floor(Math.random() * this.itemTotalCost.length)
      ]
    };

    const request = {
      contractId: this.roundArguments.contractId,
      contractFunction: "NewInvoiceAdditionalWork",
      invokerIdentity: "User1",
      contractArguments: [
        docType,
        invoice_type_work,
        invoice_work_id,
        invoice_description,
        invoice_work_cost_code,
        variation_additional_expenses_add_work,
        JSON.stringify(invoiceItemOne),
        JSON.stringify(invoiceItemTwo),
        JSON.stringify(invoiceItemThree),
        JSON.stringify(invoiceItemFour),
        JSON.stringify(invoiceItemFive)
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
  return new CreatePurchaseOrderWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;
