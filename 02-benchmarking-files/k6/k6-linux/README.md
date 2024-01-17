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
* A single table represents the test results for a single region and 
* Three separate tables depict the test results for three distinct regions.

**Table** - Single region `end-of-the-test` summary (SE - region)

| Chaincode | Chaincode Definition | Data received [kB] | Data sent [kB/s] | Avg. Latency [s] | Total requests [#] | Requests rate [TPS rate] | Failed [%] | Failed [#] | p(95) [ms] |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Organisational Chaincode* | Adding a new employee to the organisation. | 666 | 11 | 0.80 | 751 | 12.3626/s | 0.13 | 1 | 870.02 |
| Operational Chaincode* | Creating a new daily operations log. | 707 | 12 | 0.75 | 804 | 13.2384/s | 4.60 | 37 | 847.23 |
| Maintenance Chaincode* | Creating preventive work orders. | 755 | 12 | 0.76 | 789 | 12.9931/s | 0.38 | 3 | 824.78 |
|  | Creating corrective orders. | 649 | 11 | 0.83 | 721 | 11.8730/s | 1.38 | 10 | 971.13 |
| Proficiency Chaincode* | Creating training Module One. | 725 | 12 | 0.73 | 822 | 13.5449/s | 0.12 | 1 | 817.56 |
| Health & Safety Chaincode** | Creating incident and accident reports. | 616 | 10 | 0.42 | 714 | 11.8308/s | 0.00 | 0 | 484.55 |
|  | Creating risk assessments. | 539 | 8.9 | 0.48 | 621 | 10.2629/s | 0.00 | 0 | 596.36 |
| Financial Chaincode** | Creating purchase order. | 577 | 9.5 | 0.49 | 608 | 10.0534/s | 1.15 | 7 | 589.29 |
|  | Creating operational invoices. | 524 | 8.7 | 0.49 | 608 | 10.0631/s | 1.15 | 7 | 576.20 |
|  | Creating invoices for additional work. | 533 | 8.8 | 0.48 | 620 | 10.2148/s | 1.93 | 12 | 593.84 |

<table>
  <style>
    td {
      text-align: center;
    }
  </style>
  <tr>
    <th rowspan="2">Chaincode</th>
    <th rowspan="2">Chaincode Definition</th>
    <th colspan="2">Data received</th>
    <th colspan="2">Data sent</th>
    <th colspan="1">Avg. Latency</th>
    <th colspan="1">Total requests</th>
    <th colspan="1">Requests rate</th>
    <th colspan="2">Failed</th>
    <th colspan="1">p(95)</th>
  </tr>
  <tr>
    <td>[kB]</td>
    <td>[kB/s]</td>
    <td>[kB]</td>
    <td>[kB/s]</td>
    <td>[s]</td>
    <td>[#]</td>
    <td>[TPS rate]</td>
    <td>[%]</td>
    <td>[#]</td>
    <td>[ms]</td>
  </tr>
  <!-- Add rows here -->
</table>



___

Here is the HTML code for the table:





What is REST API? [Read more](https://www.redhat.com/en/topics/api/what-is-a-rest-api).

[Back to the main page.](../../../README.md)
