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
class ReadPurchaseOrderOneWorkload extends WorkloadModuleBase {
  constructor() {
    super();

    this.txIndex = 0;

    this.costCode = ["11111", "22222", "33333", "44444", "55555"];

    this.itemDescription = [
      "Item Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit."
    ];

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

      const docType = "purchase-request";

      const purchase_request_type =
        "purchase-request-maintenance-parts" +
        `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;

      const purchase_order_id = `purchase-order-${this.workerIndex}_${i}`;

      const requestItemOne = {
        cost_code_item_one: this.costCode[
          Math.floor(Math.random() * this.costCode.length)
        ],
        item_description: this.itemDescription[
          Math.floor(Math.random() * this.itemDescription.length)
        ],
        item_one_cost: this.itemCost[
          Math.floor(Math.random() * this.itemCost.length)
        ],
        item_one_quantity: this.itemQuantity[
          Math.floor(Math.random() * this.itemQuantity.length)
        ],
        item_one_total_cost: this.itemTotalCost[
          Math.floor(Math.random() * this.itemTotalCost.length)
        ]
      };

      const requestItemTwo = {
        cost_code_item_two: this.costCode[
          Math.floor(Math.random() * this.costCode.length)
        ],
        item_description: this.itemDescription[
          Math.floor(Math.random() * this.itemDescription.length)
        ],
        item_two_cost: this.itemCost[
          Math.floor(Math.random() * this.itemCost.length)
        ],
        item_two_quantity: this.itemQuantity[
          Math.floor(Math.random() * this.itemQuantity.length)
        ],
        item_two_total_cost:
          this.itemCost[Math.floor(Math.random() * this.itemCost.length)] *
          this.itemQuantity[
            Math.floor(Math.random() * this.itemQuantity.length)
          ]
      };

      const requestItemThree = {
        cost_code_item_three: this.costCode[
          Math.floor(Math.random() * this.costCode.length)
        ],
        item_description: this.itemDescription[
          Math.floor(Math.random() * this.itemDescription.length)
        ],
        item_three_cost: this.itemCost[
          Math.floor(Math.random() * this.itemCost.length)
        ],
        item_three_quantity: this.itemQuantity[
          Math.floor(Math.random() * this.itemQuantity.length)
        ],
        item_three_total_cost:
          this.itemCost[Math.floor(Math.random() * this.itemCost.length)] *
          this.itemQuantity[
            Math.floor(Math.random() * this.itemQuantity.length)
          ]
      };

      const requestItemFour = {
        cost_code_item_four: this.costCode[
          Math.floor(Math.random() * this.costCode.length)
        ],
        item_description: this.itemDescription[
          Math.floor(Math.random() * this.itemDescription.length)
        ],
        item_four_cost: this.itemCost[
          Math.floor(Math.random() * this.itemCost.length)
        ],
        item_four_quantity: this.itemQuantity[
          Math.floor(Math.random() * this.itemQuantity.length)
        ],
        item_four_total_cost:
          this.itemCost[Math.floor(Math.random() * this.itemCost.length)] *
          this.itemQuantity[
            Math.floor(Math.random() * this.itemQuantity.length)
          ]
      };

      const requestItemFive = {
        cost_code_item_five: this.costCode[
          Math.floor(Math.random() * this.costCode.length)
        ],
        item_description: this.itemDescription[
          Math.floor(Math.random() * this.itemDescription.length)
        ],
        item_five_cost: this.itemCost[
          Math.floor(Math.random() * this.itemCost.length)
        ],
        item_five_quantity: this.itemQuantity[
          Math.floor(Math.random() * this.itemQuantity.length)
        ],
        item_five_total_cost:
          this.itemCost[Math.floor(Math.random() * this.itemCost.length)] *
          this.itemQuantity[
            Math.floor(Math.random() * this.itemQuantity.length)
          ]
      };

      let supplier_order_id =
        "supplier-order-ID-" +
        `${this.roundIndex}_${this.workerIndex}_${this.txIndex}_${Date.now()}`;

      let supplier_order_description =
        "Supplier Order Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit.";

      let purchase_request_remarks =
        "Purchase Request Remarks: Lorem ipsum dolor sit amet, consectetur adipiscing elit.";

      console.log(
        `Worker ${this
          .workerIndex}: Creating Purchase Order ${purchase_order_id}`
      );
      const request = {
        contractId: this.roundArguments.contractId,
        contractFunction: "NewPurchaseRequestMtncParts",
        invokerIdentity: "User1",
        contractArguments: [
          docType,
          purchase_request_type,
          purchase_order_id,

          JSON.stringify(requestItemOne),
          JSON.stringify(requestItemTwo),
          JSON.stringify(requestItemThree),
          JSON.stringify(requestItemFour),
          JSON.stringify(requestItemFive),

          supplier_order_id,
          supplier_order_description,
          purchase_request_remarks
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
      contractFunction: "GetPurchaseRequestByID",
      invokerIdentity: "User1",
      // Pass the correct value of purchase_order_id that was created in the initializeWorkloadModule function
      contractArguments: [`purchase-order-${this.workerIndex}_${randomId}`],
      
      readOnly: true
    };

    await this.sutAdapter.sendRequests(myArgs);
  }
}

function createWorkloadModule() {
  return new ReadPurchaseOrderOneWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;
