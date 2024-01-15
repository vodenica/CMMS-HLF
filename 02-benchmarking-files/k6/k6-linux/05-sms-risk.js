// Javascript file to create a new Risk Assessment record
// The script is executed using the k6 load testing tool
// The script is executed using the following command:
// k6 run 05-sms-risk.js --vus 1 --duration 1s --iterations 1 --summary-trend-stats "avg,min,med,max,p(95),p(99)"

// Import the http module to make the http requests
import http from "k6/http";

// Function to generate a random number as the id, padded to 5 digits
export default function() {
  const url =
    "https://k0f9wp2ou5:[your-password]-k0uccgu26l-connect.kr0-aws-ws.kaleido.io/transactions";

  /*
    // Function to generate a random number as the id, padded to 5 digits
  function generateRandomId() {
    const randomNumber = Math.floor(Math.random() * 100000);
    return `risk-assessment-id-${randomNumber
      .toString()
      .padStart(5, "0")}`;
  }
  */

  // Create the data object to be sent in the request
  const data = {
    headers: {
      type: "SendTransaction",
      signer: "General Manager",
      channel: "default-channel",
      chaincode: "healthandsafety"
    },
    func: "CreateRiskAssessment",
    args: [],
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
