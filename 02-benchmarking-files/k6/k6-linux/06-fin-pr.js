// Javascript file to create a new purchase order for maintenance spare parts
// The script is executed using the k6 load testing tool
// The script is executed using the following command:
// k6 run 06-fin-pr.js --vus 1 --duration 1s --iterations 1 --summary-trend-stats "avg,min,med,max,p(95),p(99)"
// k6 run 06-fin-pr.js --vus 10 --duration 60s --iterations 10 --summary-trend-stats "avg,min,med,max,p(95),p(99)"

// Import the http module to make the http requests
import http from "k6/http";

// Function to generate a random number as the id, padded to 5 digits
export default function() {
  const url =
    "https://k0f9wp2ou5:[your-password]-k0uccgu26l-connect.kr0-aws-ws.kaleido.io/transactions";

  // Function to generate a random number as the id, padded to 5 digits
  function generateRandomId() {
    const randomNumber = Math.floor(Math.random() * 100000);
    return `purchase-request-id-k6-${randomNumber.toString().padStart(5, "0")}`;
  }

  // Create the data object to be sent in the request
  const data = {
    headers: {
      type: "SendTransaction",
      signer: "Supervisor",
      channel: "default-channel",
      chaincode: "finance"
    },
    func: "NewPurchaseRequestMtncParts",
    args: [
      "purchase-request",
      "maintenance-parts",
      generateRandomId(), // Generate the random id
      '{"cost_code_item_one":"11111", "item_description":"Maintenance Part","item_one_cost":100,"item_one_quantity":2,"item_one_total_cost":200}',
      '{"cost_code_item_two":"22222", "item_description":"Maintenance Part","item_two_cost":100,"item_two_quantity":2,"item_two_total_cost":200}',
      '{"cost_code_item_three":"33333", "item_description":"Maintenance Part","item_three_cost":100,"item_three_quantity":2,"item_three_total_cost":200}',
      '{"cost_code_item_four":"44444", "item_description":"Maintenance Part","item_four_cost":100,"item_four_quantity":2,"item_four_total_cost":200}',
      '{"cost_code_item_five":"55555", "item_description":"Maintenance Part","item_five_cost":100,"item_five_quantity":2,"item_five_total_cost":200}',
      "Supplier ID: 123456",
      "The Purchase Order is issued for the purchase of the maintenance spare parts.",
      "No remarks"
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
