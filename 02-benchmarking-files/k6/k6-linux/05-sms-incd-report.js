// Javascript file to create a new Incident / Accident Report
// The script is executed using the k6 load testing tool
// The script is executed using the following command:
// k6 run 05-sms-incd-report.js --vus 1 --duration 1s --iterations 1 --summary-trend-stats "avg,min,med,max,p(95),p(99)"

// Import the http module to make the http requests
import http from "k6/http";

// Function to generate a random number as the id, padded to 5 digits
export default function() {
  const url =
    "https://k0f9wp2ou5:[your-password]-k0uccgu26l-connect.kr0-aws-ws.kaleido.io/transactions";

  // Function to generate a random number as the id, padded to 5 digits
  function generateRandomId() {
    const randomNumber = Math.floor(Math.random() * 100000);
    return `incident-report-k6-id-${randomNumber.toString().padStart(5, "0")}`;
  }

  // Create the data object to be sent in the request
  const data = {
    headers: {
      type: "SendTransaction",
      signer: "General Manager",
      channel: "default-channel",
      chaincode: "healthandsafety"
    },
    func: "CreateAccidentIncidentReport",
    args: [
      "Subject(K6): Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
      "Location: Lorem ipsum dolor sit amet.",
      "HS: Aspects: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
      "Description: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nibh velit, egestas nec laoreet a, efficitur vitae purus. Suspendisse finibus est nec est ultrices placerat. Cras condimentum pharetra dui et semper.",
      "Classification: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas placerat.",
      '["Lorem ipsum dolor sit amet consectetur adipiscing elit Maecenas placerat", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Maecenas placerat."]'
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
