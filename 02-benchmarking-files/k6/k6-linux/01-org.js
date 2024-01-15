// Javascript file to add a new employee
// The script is executed using the k6 load testing tool
// The script is executed using the following command:
// k6 run 01-org.js or k6 run --vus 10 --duration 60s 01-org.js

// Import the http module to make the http requests
import http from "k6/http";

// Function to generate a random number as the id, padded to 5 digits
export default function() {
  const url =
    "https://k0f9wp2ou5:[your-password]-k0uccgu26l-connect.kr0-aws-ws.kaleido.io/transactions";

  // Function to generate a random number as the id, padded to 5 digits
  function generateRandomId() {
    const randomNumber = Math.floor(Math.random() * 100000);
    return `employee-id-${randomNumber.toString().padStart(5, "0")}`;
  }

  // Create the data object to be sent in the request
  const data = {
    headers: {
      type: "SendTransaction",
      signer: "Supervisor",
      channel: "default-channel",
      chaincode: "organization"
    },
    func: "AddNewEmployee",
    args: [
      "S.A.R. Macao, China",
      generateRandomId(), // Generate the random id
      "Name",
      "Last Name",
      "Gender",
      "Position entry",
      "Current position",
      "Level FT/PT",
      "Birthday",
      "Contract Signed Date",
      "Starting Date",
      "Years in service",
      "Employee review",
      "Completed custom training",
      "100000",
      "Address",
      "Cell-phone-number",
      "Working visa Yes/No"
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
