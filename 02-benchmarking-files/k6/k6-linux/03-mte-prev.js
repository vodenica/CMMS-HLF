// Javascript file to create a new Preventive Work Order
// The script is executed using the k6 load testing tool
// The script is executed using the following command:
// k6 run 03-mte-prev.js or k6 run --vus 10 --duration 60s 03-mte-prev.js

// Import the http module to make the http requests
import http from "k6/http";

// Function to generate a random number as the id, padded to 5 digits
export default function() {
  const url =
    "https://k0f9wp2ou5:[your-password]-k0uccgu26l-connect.kr0-aws-ws.kaleido.io/transactions";

  // Function to generate a random number as the id, padded to 5 digits
  function generateRandomId() {
    const randomNumber = Math.floor(Math.random() * 100000);
    return `preventive-work-order-id-${randomNumber
      .toString()
      .padStart(5, "0")}`;
  }

  // Create the data object to be sent in the request
  const data = {
    headers: {
      type: "SendTransaction",
      signer: "Maintenance Technician",
      channel: "default-channel",
      chaincode: "maintenance"
    },
    func: "CreatePreventiveWorkOrder",
    args: [
      generateRandomId(), // Generate the random id
      "Preventive Maintenance Work Order",
      "Description of the preventive maintenance work order",
      "General Health and Safety instructions: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut at molestie tortor. Phasellus iaculis id enim ac bibendum. Nullam gravida quam eget lacus dictum, vel ullamcorper eros condimentum. Aenean laoreet elit ut magna gravida dapibus. Ut scelerisque sit amet justo eget auctor. Donec fringilla, eros pharetra molestie tristique, urna felis maximus turpis, sed interdum velit turpis ut leo. Nam commodo, est sed congue malesuada, neque lorem elementum odio, ut tempus dui neque non ex. Aliquam semper faucibus ante, eget feugiat erat ullamcorper a. Aenean sagittis lacus mi, a egestas nibh luctus et. Sed congue nisl et nunc tincidunt, eu tincidunt nunc vehicula. Donec dui lorem, vehicula ac lectus at, venenatis mattis odio.",
      '{"maintenance_supervisor":"Maintenance Supervisor","maintenance_technician_one":"Maintenance Technician 1","maintenance_technician_two":"Maintenance Technician 2", "maintenance_technician_three": "Maintenance Technician 3","system_operator":" System Operator"}',
      '{"maintenance_process_step_one":"Maintenance Process Step 1", "maintenance_process_step_two":"Maintenance Process Step 2", "maintenance_process_step_three":"Maintenance Process Step 3", "maintenance_process_step_four":"Maintenance Process Step 4", "maintenance_process_step_five":"Maintenance Process Step 5"}',
      '{"maintenance_part_one":"Part-ID-15467", "maintenance_part_two":"Part-ID-15468", "maintenance_part_three":"Part-ID-15469", "maintenance_part_four":"Part-ID-15470", "maintenance_part_five":"Part-ID-15471", "maintenance_part_six":"Part-ID-15472", "maintenance_part_seven":"Part-ID-15473", "maintenance_part_eight":"Part-ID-15474", "maintenance_part_nine":"Part-ID-15475", "maintenance_part_ten":"Part-ID-15476"}'
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
