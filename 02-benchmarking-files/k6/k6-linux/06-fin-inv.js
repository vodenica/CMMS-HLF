// Javascript file to create a new invoice for operational services
// The script is executed using the k6 load testing tool
// The script is executed using the following command:
// k6 run 06-fin-inv.js --vus 1 --duration 1s --iterations 1 --summary-trend-stats "avg,min,med,max,p(95),p(99)"
// k6 run 06-fin-inv.js --vus 10 --duration 60s --iterations 10 --summary-trend-stats "avg,min,med,max,p(95),p(99)"
// k6 run --vus 5 --duration 60s 06-fin-inv.js

// Import the http module to make the http requests
import http from "k6/http";

// Function to generate a random number as the id, padded to 5 digits
export default function() {
  const url =
    "https://k0f9wp2ou5:[your-password]-k0uccgu26l-connect.kr0-aws-ws.kaleido.io/transactions";

  // Function to generate a random number as the id, padded to 5 digits
  function generateRandomId() {
    const randomNumber = Math.floor(Math.random() * 100000);
    return `invoice-ops-id-k6-${randomNumber.toString().padStart(5, "0")}`;
  }

  // Create the data object to be sent in the request
  const data = {
    headers: {
      type: "SendTransaction",
      signer: "Supervisor",
      channel: "default-channel",
      chaincode: "finance"
    },
    func: "NewInvoiceOperationalServices",
    args: [
      "invoice-ops",
      "operations",
      generateRandomId(), // Generate the random id
      "200000",
      "0",
      "10000",
      "8000",
      "5000"
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
