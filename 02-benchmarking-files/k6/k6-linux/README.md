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
  <tr>
    <th rowspan="2">Chaincode</th>
    <th rowspan="2">Chaincode Definition</th>
    <th colspan="2">Data received</th>
    <th colspan="2">Data sent</th>
    <th rowspan="2">Avg. Latency</th>
    <th rowspan="2">Total requests</th>
    <th rowspan="2">Requests rate</th>
    <th colspan="2">Failed</th>
    <th rowspan="2">p(95)</th>
  </tr>
  <tr>
    <td>[kB]</td>
    <td>[kB/s]</td>
    <td>[kB]</td>
    <td>[kB/s]</td>
    <td>[%]</td>
    <td>[#]</td>
  </tr>
  <!-- Add rows here -->
</table>



___

Here is the HTML code for the table:

<table>
  <tr>
    <th rowspan="2">Chaincode</th>
    <th rowspan="2">Chaincode Definition</th>
    <th colspan="2">Data received</th>
    <th colspan="2">Data sent</th>
    <th rowspan="2">Avg. Latency</th>
    <th rowspan="2">Total requests</th>
    <th rowspan="2">Requests rate</th>
    <th colspan="2">Failed</th>
    <th rowspan="2">p(95)</th>
  </tr>
  <tr>
    <td>[kB]</td>
    <td>[kB/s]</td>
    <td>[kB]</td>
    <td>[kB/s]</td>
    <td>[%]</td>
    <td>[#]</td>
  </tr>
  <tr>
    <td rowspan="2">Organisational Chaincode*</td>
    <td>Adding a new employee to the organisation.</td>
    <td>666</td>
    <td>405</td>
    <td>0.80</td>
    <td>751</td>
    <td>12.3626/s</td>
    <td>0.13</td>
    <td>1</td>
    <td>870.02</td>
  </tr>
  <tr>
    <td colspan="9"></td>
  </tr>
  <tr>
    <td rowspan="2">Operational Chaincode*</td>
    <td>Creating a new daily operations log.</td>
    <td>707</td>
    <td>1000</td>
    <td>0.75</td>
    <td>804</td>
    <td>13.2384/s</td>
    <td>4.60</td>
    <td>37</td>
    <td>847.23</td>
  </tr>
  <tr>
    <td colspan="9"></td>
  </tr>
  <tr>
    <td rowspan="3">Maintenance Chaincode*</td>
    <td>Creating preventive work orders.</td>
    <td>755</td>
    <td>1800</td>
    <td>0.76</td>
    <td>789</td>
    <td>12.9931/s</td>
    <td>0.38</td>
    <td>3</td>
    <td>824.78</td>
  </tr>
  <tr>
    <td>Creating corrective orders.</td>
    <td>649</td>
    <td>1100</td>
    <td>0.83</td>
    <td>721</td>
    <td>11.8730/s</td>
    <td>1.38</td>
    <td>10</td>
    <td>971.13</td>
  </tr>
  <tr>
    <td colspan="9"></td>
  </tr>
  <tr>
    <td rowspan="2">Proficiency Chaincode*</td>
    <td>Creating training Module One.</td>
    <td>725</td>
    <td>1100</td>
    <td>0.73</td>
    <td>822</td>
    <td>13.5449/s</td>
    <td>0.12</td>
    <td>1</td>
    <td>817.56</td>
  </tr>
  <tr>
    <td colspan="9"></td>
  </tr>
  <tr>
    <td rowspan="3">Health & Safety Chaincode**</td>
    <td>Creating incident and accident reports.</td>
    <td>616</td>
    <td>690</td>
    <td>0.42</td>
    <td>714</td>
    <td>11.8308/s</td>
    <td>0.00</td>
    <td>0</td>
    <td>484.55</td>
  </tr>
  <tr>
    <td>Creating risk assessments.</td>
    <td>539</td>
    <td>158</td>
    <td>0.48</td>
    <td>621</td>
    <td>10.2629/s</td>
    <td>0.00</td>
    <td>0</td>
    <td>596.36</td>
  </tr>
  <tr>
    <td colspan="9"></td>
  </tr>
  <tr>
    <td rowspan="4">Financial Chaincode**</td>
    <td>Creating purchase order.</td>
    <td>577</td>
    <td>748</td>
    <td>0.49</td>
    <td>608</td>
    <td>10.0534/s</td>
    <td>1.15</td>
    <td>7</td>
    <td>589.29</td>
  </tr>
  <tr>
    <td>Creating operational invoices.</td>
    <td>524</td>
    <td>205</td>
    <td>0.49</td>
    <td>608</td>
    <td>10.0631/s</td>
    <td>1.15</td>
    <td>7</td>
    <td>576.20</td>
  </tr>
  <tr>
    <td>Creating invoices for additional work.</td>
    <td>533</td>
    <td>548</td>
    <td>0.48</td>
    <td>620</td>
    <td>10.2148/s</td>
    <td>1.93</td>
    <td>12</td>
    <td>593.84</td>
  </tr>
  <tr>
    <td colspan="9"></td>
  </tr>
</table>



What is REST API? [Read more](https://www.redhat.com/en/topics/api/what-is-a-rest-api).

[Back to the main page.](../../../README.md)
