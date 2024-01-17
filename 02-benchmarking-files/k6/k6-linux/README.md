[Main page.](../../../README.md)

# k6 load testing tool

In this repository are `JavaScript` files for REST API load testing with [k6](https://k6.io/). 

Representational Stateless Transfer Application Programming Interface (REST API) is a web service that allows interaction with cloud resources using HTTP requests. We use `JavaScript` to write the test scripts and load our exposed REST API endpoints with `k6`. 

To run the [load files](../k6-linux/), you need to install `k6` on your machine. You can find the installation instructions [here](https://k6.io/docs/getting-started/installation/).

## Running the load tests

Run these commands in the terminal to run the load tests:

```bash
k6 run [file_name].js or 
```
or
```bash
k6 run --vus 10 --duration 60s [file_name].js
```
As seen in one sample in the picture below, K6 generates and presents a thorough summary of the aggregated results at the end of a test.

![k6 results](../../../05-plots/images/01-ops.png)

The `k6` load testing results are comprehensively compiled from end-of-test summary reports for each chaincode function. These reports detail aggregate statistics for the primary aspects of the test. As illustrated in the subsequent tables: 
* A single table represents the test results for a single region [Table 1](#table-1---single-region-summary) and 
* Three separate tables depict the test results for three distinct regions [Table 2](#table-2---multi-region-summary-for-SE-region), [Table 3](#table-3---multi-region-summary-for-FR-region), and [Table 4](#table-4---multi-region-summary-for-OH-region).

### Table 1 - Single region summary
The table shows the `end-of-test` summary for the load test in a single region environment (_**SE region**_)
<table>
  <tr>
    <th align="center" rowspan="2">Chaincode</th>
    <th align="center" rowspan="2">Chaincode Function</th>
    <th align="center" colspan="2">Data received</th>
    <th align="center" colspan="2">Data sent</th>
    <th align="center" colspan="1">Avg. Latency</th>
    <th align="center" colspan="1">Total requests</th>
    <th align="center" colspan="1">Requests rate</th>
    <th align="center" colspan="2">Failed</th>
    <th align="center" colspan="1">p(95)</th>
  </tr>
  <tr>
    <td align="center">[kB]</td>
    <td align="center">[kB/s]</td>
    <td align="center">[kB]</td>
    <td align="center">[kB/s]</td>
    <td align="center">[s]</td>
    <td align="center">[#]</td>
    <td align="center">[TPS rate]</td>
    <td align="center">[%]</td>
    <td align="center">[#]</td>
    <td align="center">[ms]</td>
  </tr>
  <tr>
    <td align="left">Organisational Chaincode*</td>
    <td align="left">Adding a new employee to the organisation.</td>
    <td align="center">666</td>
    <td align="center">11</td>
    <td align="center">405</td>
    <td align="center">6.7</td>
    <td align="center">0.80</td>
    <td align="center">751</td>
    <td align="center">12.3626/s</td>
    <td align="center">0.13</td>
    <td align="center">1</td>
    <td align="center">870.02</td>
  </tr>
  <tr>
    <td align="left">Operational Chaincode*</td>
    <td align="left">Creating a new daily operations log.</td>
    <td align="center">707</td>
    <td align="center">12</td>
    <td align="center">1000</td>
    <td align="center">17</td>
    <td align="center">0.75</td>
    <td align="center">804</td>
    <td align="center">13.2384/s</td>
    <td align="center">4.60</td>
    <td align="center">37</td>
    <td align="center">847.23</td>
  </tr>
  <tr>
    <td align="left" rowspan="2">Maintenance Chaincode*</td>
    <td align="left">Creating preventive work orders.</td>
    <td align="center">755</td>
    <td align="center">12</td>
    <td align="center">1800</td>
    <td align="center">29</td>
    <td align="center">0.76</td>
    <td align="center">789</td>
    <td align="center">12.9931/s</td>
    <td align="center">0.38</td>
    <td align="center">3</td>
    <td align="center">824.78</td>
  </tr>
  <tr>
    <td align="left">Creating corrective work orders.</td>
    <td align="center">649</td>
    <td align="center">11</td>
    <td align="center">1100</td>
    <td align="center">18</td>
    <td align="center">0.83</td>
    <td align="center">721</td>
    <td align="center">11.8730/s</td>
    <td align="center">1.38</td>
    <td align="center">10</td>
    <td align="center">971.13</td>
  </tr>
  <tr>
    <td align="left">Proficiency Chaincode*</td>
    <td align="left">Creating training Module One.</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
  </tr>
  <tr>
    <td align="left" rowspan="2">Health & Safety Chaincode**</td>
    <td align="left">Creating incident and accident reports.</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
  </tr>
  <tr>
    <td align="left">Creating risk assessments.</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
  </tr>
  <tr>
    <td align="left" rowspan="3">Financial Chaincode**</td>
    <td align="left">Creating purchase order.</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
  </tr>
  <tr>
    <td align="left">Creating operational invoices.</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
  </tr>
  <tr>
    <td align="left">Creating invoices for additional work.</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
    <td align="center">x</td>
  </tr>
  <!-- Add rows here -->
</table>

___

### Table 2 - Multi region summary for SE region

### Table 3 - Multi region summary for FR region

### Table 4 - Multi region  summary for OH region

What is REST API? [Read more](https://www.redhat.com/en/topics/api/what-is-a-rest-api).

[Back to the main page.](../../../README.md)
