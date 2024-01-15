// Javascript file to create a new invoice for additional works
// The script is executed using the k6 load testing tool
// The script is executed using the following command:
// k6 run --vus 5 --duration 60s 06-fin-inv-add-work.js

// Import the http module to make the http requests
import http from "k6/http";

// Function to generate a random number as the id, padded to 5 digits
export default function() {
  const url =
    "https://k0f9wp2ou5:[your-password]-k0uccgu26l-connect.kr0-aws-ws.kaleido.io/transactions";

  // Function to generate a random number as the id, padded to 5 digits
  function generateRandomId() {
    const randomNumber = Math.floor(Math.random() * 100000);
    return `invoice-additional-work-id-k6-${randomNumber
      .toString()
      .padStart(5, "0")}`;
  }

  // Create the data object to be sent in the request
  const data = {
    headers: {
      type: "SendTransaction",
      signer: "Supervisor",
      channel: "default-channel",
      chaincode: "finance"
    },
    func: "NewInvoiceAdditionalWork",
    args: [
      "invoice-work",
      "additional-work",
      generateRandomId(), // Generate the random id
      "The Invoice is issued for the completed additional works.",
      "1000",
      "1",
      '{"invoice_item_one_qty":2, "invoice_item_one_item_cost":100, "invoice_item_one_cost":200}',
      '{"invoice_item_two_qty":1, "invoice_item_two_item_cost":50, "invoice_item_two_cost":50}',
      '{"invoice_item_three_qty":4, "invoice_item_three_item_cost":25, "invoice_item_three_cost":100}',
      '{"invoice_item_four_qty":3, "invoice_item_four_item_cost":50, "invoice_item_four_cost":150}',
      '{"invoice_item_five_qty":30, "invoice_item_five_item_cost":4, "invoice_item_five_cost":120}'
    ],
    init: false
  };

  const params = {
    headers: {
      "Content-Type": "application/json"
    }
  };

  const response = http.post(url, JSON.stringify(data), params);

  console.log("Status code:", response.status);
  console.log("Response body:", response.body);
}
