// Javascript file to create a new Daily Operations Log
// This script is used to create a new Daily Operations Log
// The script is executed using the k6 load testing tool
// The script is executed using the following command:
// k6 run 02-ops.js or k6 run --vus 10 --duration 60s 02-ops.js

// Import the http module to make the http requests
import http from "k6/http";

// Function to generate a random number as the id, padded to 5 digits
export default function() {
  const url =
    "https://k0f9wp2ou5:[your-password]-k0uccgu26l-connect.kr0-aws-ws.kaleido.io/transactions";

  // Function to generate a random number as the id, padded to 5 digits
  function generateRandomId() {
    const randomNumber = Math.floor(Math.random() * 100000);
    return `daily-ops-log-test-id-${randomNumber.toString().padStart(5, "0")}`;
  }

  // Create the data object to be sent in the request
  const data = {
    headers: {
      type: "SendTransaction",
      signer: "System Operator",
      channel: "default-channel",
      chaincode: "operations"
    },
    func: "CreateNewDailyOperationsLog",
    args: [
      generateRandomId(), // Generate the random id
      "Creator of the Daily Operational Log",
      '{"weather_drive_station_condition":"Sunny", "weather_drive_station_temperature":"+25 C", "weather_drive_station_humidity":"95 %"}',
      '{"weather_return_station_condition":"Cloudy", "weather_return_station_temperature":"+24 C", "weather_return_station_humidity":"95 %"}',
      '{"personnel_on_duty_drive_station_shift_manager":"Shift Manager", "personnel_on_duty_drive_station_maintenance_technician_one":"Maintenance Technician", "personnel_on_duty_drive_station_maintenance_technician_two":"Maintenance Technician", "personnel_on_duty_drive_station_system_operator_one":"System Operator", "personnel_on_duty_drive_station_system_operator_two":"System Operator"}',
      '{"personnel_on_duty_return_station_maintenance_technician":"Maintenance Technician", "personnel_on_duty_return_station_system_operator":"System Operator"}',
      "Operations start",
      "Operations end",
      "10000",
      "35",
      "2500",
      "No comments so far."
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
