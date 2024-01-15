// Javascript file to create a new Proficiency Moduel 1 training record
// The script is executed using the k6 load testing tool
// The script is executed using the following command:
// k6 run 04-prf.js --vus 1 --duration 1s --iterations 1 --summary-trend-stats "avg,min,med,max,p(95),p(99)"
// k6 run 04-prf.js --vus 10 --duration 60s --iterations 10 --summary-trend-stats "avg,min,med,max,p(95),p(99)"
// k6 run --vus 10 --duration 60s 04-prf.js

// Import the http module to make the http requests
import http from "k6/http";

// Function to generate a random number as the id, padded to 5 digits
export default function() {
  const url =
    "https://k0f9wp2ou5:[your-password]-k0uccgu26l-connect.kr0-aws-ws.kaleido.io/transactions";

  // Function to generate a random number as the id, padded to 5 digits
  function generateRandomId() {
    const randomNumber = Math.floor(Math.random() * 100000);
    return `proficiency-module-one-${randomNumber.toString().padStart(5, "0")}`;
  }

  // Create the data object to be sent in the request
  const data = {
    headers: {
      type: "SendTransaction",
      signer: "Supervisor",
      channel: "default-channel",
      chaincode: "proficiency"
    },
    func: "CreateNewModuleOne",
    args: [
      generateRandomId(), // Generate the random id
      '{"trainer_module_one_id": "Trainer-ID", "trainer_module_one_name": "Trainer Name", "trainer_module_one_surname": "Trainer Surname"}',
      '{"trainee_module_one_id": "Trainee-ID", "trainee_module_one_name": "Trainee Name", "trainee_module_one_surname": "Trainee Surname"}',
      '{"module_one_chapter_one_session_one":"Module One - Chapter 01 - Session 03","module_one_chapter_one_session_two":"Module One - Chapter 01 - Session 03", "module_one_chapter_one_session_three":"Module One - Chapter 01 - Session 03"}',
      '{"module_one_chapter_two_session_one":"Module One - Chapter 02 - Session 03","module_one_chapter_two_session_two":"Module One - Chapter 02 - Session 03", "module_one_chapter_two_session_three":"Module One - Chapter 02 - Session 03"}',
      '{"module_one_chapter_three_session_one":"Module One - Chapter 03 - Session 03","module_one_chapter_three_session_two":"Module One - Chapter 03 - Session 03", "module_one_chapter_three_session_three":"Module One - Chapter 03 - Session 03"}'
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
