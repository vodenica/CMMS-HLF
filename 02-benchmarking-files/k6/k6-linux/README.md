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
As seen in one sample in the picture below, K6 generates and presents a thorough summary of the aggregated results at the end of a test. As seen below in Figure 1.

![k6 results](../../../05-plots/images/01-ops.png)

**Figure 1** - End-of-test summary

## End of test summary
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
    <th align="center" colspan="1">p(95)***</th>
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
    <td align="center">725</td>
    <td align="center">12</td>
    <td align="center">1100</td>
    <td align="center">18</td>
    <td align="center">0.73</td>
    <td align="center">822</td>
    <td align="center">13.5449/s</td>
    <td align="center">0.12</td>
    <td align="center">1</td>
    <td align="center">817.56</td>
  </tr>
  <tr>
    <td align="left" rowspan="2">Health & Safety Chaincode**</td>
    <td align="left">Creating incident and accident reports.</td>
    <td align="center">616</td>
    <td align="center">10</td>
    <td align="center">690</td>
    <td align="center">11</td>
    <td align="center">0.42</td>
    <td align="center">714</td>
    <td align="center">11.8308/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">484.55</td>
  </tr>
  <tr>
    <td align="left">Creating risk assessments.</td>
    <td align="center">539</td>
    <td align="center">8.9</td>
    <td align="center">158</td>
    <td align="center">2.6</td>
    <td align="center">0.48</td>
    <td align="center">621</td>
    <td align="center">10.2629/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">596.36</td>
  </tr>
  <tr>
    <td align="left" rowspan="3">Financial Chaincode**</td>
    <td align="left">Creating purchase order.</td>
    <td align="center">577</td>
    <td align="center">9.5</td>
    <td align="center">748</td>
    <td align="center">12</td>
    <td align="center">0.49</td>
    <td align="center">608</td>
    <td align="center">10.0534/s</td>
    <td align="center">1.15</td>
    <td align="center">7</td>
    <td align="center">589.29</td>
  </tr>
  <tr>
    <td align="left">Creating operational invoices.</td>
    <td align="center">524</td>
    <td align="center">8.7</td>
    <td align="center">205</td>
    <td align="center">3.4</td>
    <td align="center">0.49</td>
    <td align="center">608</td>
    <td align="center">10.0631/s</td>
    <td align="center">1.15</td>
    <td align="center">7</td>
    <td align="center">576.20</td>
  </tr>
  <tr>
    <td align="left">Creating invoices for additional work.</td>
    <td align="center">533</td>
    <td align="center">8.8</td>
    <td align="center">548</td>
    <td align="center">9</td>
    <td align="center">0.48</td>
    <td align="center">620</td>
    <td align="center">10.2148/s</td>
    <td align="center">1.93</td>
    <td align="center">12</td>
    <td align="center">593.84</td>
  </tr>
  <!-- Add rows here -->
</table>

Note:

* `*` represents ten (10) users.
*  `**` represents five (5) users.
*  `***` the 95th percentile represents a value below which 95% of the data points fall.

[Back on top](#end-of-test-summary)
___

### Table 2 - Multi region summary for SE region
The table shows the `end-of-test` summary for the load test performed in a multi-region environment, where we observed a single region environment or SE-region (_**SE region**_, _SE - Seoul, South Korea_)
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
    <th align="center" colspan="1">p(95)***</th>
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
    <td align="left">Organisational Chaincode**</td>
    <td align="left">Adding a new employee to the organisation.</td>
    <td align="center">321</td>
    <td align="center">3.5</td>
    <td align="center">198</td>
    <td align="center">2.1</td>
    <td align="center">1.26</td>
    <td align="center">359</td>
    <td align="center">13.8782/s</td>
    <td align="center">1.39</td>
    <td align="center">5</td>
    <td align="center">2.47</td>
  </tr>
  <tr>
    <td align="left">Operational Chaincode**</td>
    <td align="left">Creating a new daily operations log.</td>
    <td align="center">412</td>
    <td align="center">4.5</td>
    <td align="center">600</td>
    <td align="center">6.6</td>
    <td align="center">0.96</td>
    <td align="center">469</td>
    <td align="center">5.1570/s</td>
    <td align="center">0.32</td>
    <td align="center">2</td>
    <td align="center">1.71</td>
  </tr>
  <tr>
    <td align="left" rowspan="2">Maintenance Chaincode**</td>
    <td align="left">Creating preventive work orders.</td>
    <td align="center">535</td>
    <td align="center">5.9</td>
    <td align="center">1400</td>
    <td align="center">15</td>
    <td align="center">0.73</td>
    <td align="center">617</td>
    <td align="center">6.8111/s</td>
    <td align="center">0.32</td>
    <td align="center">2</td>
    <td align="center">1.00</td>
  </tr>
  <tr>
    <td align="left">Creating corrective work orders.</td>
    <td align="center">437</td>
    <td align="center">4.8</td>
    <td align="center">743</td>
    <td align="center">8.2</td>
    <td align="center">0.90</td>
    <td align="center">500</td>
    <td align="center">5.5019/s</td>
    <td align="center">0.8</td>
    <td align="center">4</td>
    <td align="center">1.14</td>
  </tr>
  <tr>
    <td align="left">Proficiency Chaincode**</td>
    <td align="left">Creating training Module One.</td>
    <td align="center">484</td>
    <td align="center">5.3</td>
    <td align="center">702</td>
    <td align="center">7.7</td>
    <td align="center">0.86</td>
    <td align="center">526</td>
    <td align="center">5.7851/s</td>
    <td align="center">0.57</td>
    <td align="center">3</td>
    <td align="center">1.16</td>
  </tr>
  <tr>
    <td align="left" rowspan="2">Health & Safety Chaincode**</td>
    <td align="left">Creating incident and accident reports.</td>
    <td align="center">375</td>
    <td align="center">4.1</td>
    <td align="center">412</td>
    <td align="center">4.6</td>
    <td align="center">1.06</td>
    <td align="center">424</td>
    <td align="center">4.6857/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">1.36</td>
  </tr>
  <tr>
    <td align="left">Creating risk assessments.</td>
    <td align="center">322</td>
    <td align="center">3.5</td>
    <td align="center">94</td>
    <td align="center">4.0</td>
    <td align="center">1.26</td>
    <td align="center">360</td>
    <td align="center">3.9019/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">2.76</td>
  </tr>
  <tr>
    <td align="left" rowspan="3">Financial Chaincode**</td>
    <td align="left">Creating purchase order.</td>
    <td align="center">381</td>
    <td align="center">4.2</td>
    <td align="center">444</td>
    <td align="center">4.9</td>
    <td align="center">1.22</td>
    <td align="center">365</td>
    <td align="center">4.0329/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">1.58</td>
  </tr>
  <tr>
    <td align="left">Creating operational invoices.</td>
    <td align="center">283</td>
    <td align="center">3.1</td>
    <td align="center">100</td>
    <td align="center">1.1</td>
    <td align="center">1.44</td>
    <td align="center">312</td>
    <td align="center">3.4426/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">2.91</td>
  </tr>
  <tr>
    <td align="left">Creating invoices for additional work.</td>
    <td align="center">481</td>
    <td align="center">5.3</td>
    <td align="center">470</td>
    <td align="center">5.2</td>
    <td align="center">0.82</td>
    <td align="center">552</td>
    <td align="center">6.0189</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">1.22</td>
  </tr>
</table>

Note:

* `*` represents ten (10) users.
*  `**` represents five (5) users.
*  `***` the 95th percentile represents a value below which 95% of the data points fall.

[Back on top](#end-of-test-summary)

___

### Table 3 - Multi region summary for FR region
The table shows the `end-of-test` summary for the load test performed in a multi-region environment, where we observed a single region environment or FR-region (_**FR region**_, _FR - Frankfurt, Germany_)
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
    <th align="center" colspan="1">p(95)***</th>
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
    <td align="left">Organisational Chaincode**</td>
    <td align="left">Adding a new employee to the organisation.</td>
    <td align="center">285</td>
    <td align="center">3.1</td>
    <td align="center">175</td>
    <td align="center">1.9</td>
    <td align="center">1.42</td>
    <td align="center">316</td>
    <td align="center">3.4542/s</td>
    <td align="center">0.94</td>
    <td align="center">3</td>
    <td align="center">2.25</td>
  </tr>
  <tr>
    <td align="left">Operational Chaincode**</td>
    <td align="left">Creating a new daily operations log.</td>
    <td align="center">374</td>
    <td align="center">4.1</td>
    <td align="center">452</td>
    <td align="center">4.9</td>
    <td align="center">1.27</td>
    <td align="center">346</td>
    <td align="center">3.7777/s</td>
    <td align="center">0.57</td>
    <td align="center">2</td>
    <td align="center">2.06</td>
  </tr>
  <tr>
    <td align="left" rowspan="2">Maintenance Chaincode*</td>
    <td align="left">Creating preventive work orders.</td>
    <td align="center">251</td>
    <td align="center">2.7</td>
    <td align="center">614</td>
    <td align="center">6.7</td>
    <td align="center">1.63</td>
    <td align="center">275</td>
    <td align="center">2.9878/s</td>
    <td align="center">1.81</td>
    <td align="center">5</td>
    <td align="center">2.25</td>
  </tr>
  <tr>
    <td align="left">Creating corrective work orders.</td>
    <td align="center">295</td>
    <td align="center">3.3</td>
    <td align="center">488</td>
    <td align="center">5.4</td>
    <td align="center">1.36</td>
    <td align="center">328</td>
    <td align="center">3.6145/s</td>
    <td align="center">0.91</td>
    <td align="center">3</td>
    <td align="center">1.65</td>
  </tr>
  <tr>
    <td align="left">Proficiency Chaincode**</td>
    <td align="left">Creating training Module One.</td>
    <td align="center">193</td>
    <td align="center">2.1</td>
    <td align="center">231</td>
    <td align="center">2.5</td>
    <td align="center">2.63</td>
    <td align="center">169</td>
    <td align="center">1.8591/s</td>
    <td align="center">1.18</td>
    <td align="center">2</td>
    <td align="center">6.17</td>
  </tr>
  <tr>
    <td align="left" rowspan="2">Health & Safety Chaincode**</td>
    <td align="left">Creating incident and accident reports.</td>
    <td align="center">245</td>
    <td align="center">2.7</td>
    <td align="center">261</td>
    <td align="center">2.9</td>
    <td align="center">1.62</td>
    <td align="center">267</td>
    <td align="center">2.9296/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">1.96</td>
  </tr>
  <tr>
    <td align="left">Creating risk assessments.</td>
    <td align="center">367</td>
    <td align="center">4.0</td>
    <td align="center">89</td>
    <td align="center">1.0</td>
    <td align="center">1.52</td>
    <td align="center">285</td>
    <td align="center">3.0708/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">2.30</td>
  </tr>
  <tr>
    <td align="left" rowspan="3">Financial Chaincode**</td>
    <td align="left">Creating purchase order.</td>
    <td align="center">321</td>
    <td align="center">3.5</td>
    <td align="center">334</td>
    <td align="center">3.6</td>
    <td align="center">1.63</td>
    <td align="center">270</td>
    <td align="center">2.9518/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">2.11</td>
  </tr>
  <tr>
    <td align="left">Creating operational invoices.</td>
    <td align="center">233</td>
    <td align="center">2.5</td>
    <td align="center">68</td>
    <td align="center">0.7</td>
    <td align="center">2.44</td>
    <td align="center">181</td>
    <td align="center">1.9526/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">4.63</td>
  </tr>
  <tr>
    <td align="left">Creating invoices for additional work.</td>
    <td align="center">180</td>
    <td align="center">2.0</td>
    <td align="center">162</td>
    <td align="center">1.8</td>
    <td align="center">2.38</td>
    <td align="center">188</td>
    <td align="center">2.0671/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">5.84</td>
  </tr>
</table>

Note:

* `*` represents ten (10) users.
*  `**` represents five (5) users.
*  `***` the 95th percentile represents a value below which 95% of the data points fall.

[Back on top](#end-of-test-summary)

___

### Table 4 - Multi region summary for OH region
The table shows the `end-of-test` summary for the load test performed in a multi-region environment, where we observed a single region environment or OH-region (_**OH region**_, _OH - Ohio, United States_)
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
    <th align="center" colspan="1">p(95)***</th>
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
    <td align="left">Organisational Chaincode**</td>
    <td align="left">Adding a new employee to the organisation.</td>
    <td align="center">394</td>
    <td align="center">4.3</td>
    <td align="center">203</td>
    <td align="center">2.2</td>
    <td align="center">1.26</td>
    <td align="center">347</td>
    <td align="center">3.8170/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">3.01</td>
  </tr>
  <tr>
    <td align="left">Operational Chaincode**</td>
    <td align="left">Creating a new daily operations log.</td>
    <td align="center">335</td>
    <td align="center">3.7</td>
    <td align="center">480</td>
    <td align="center">5.3</td>
    <td align="center">1.19</td>
    <td align="center">3.75</td>
    <td align="center">4.1294/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">2.63</td>
  </tr>
  <tr>
    <td align="left" rowspan="2">Maintenance Chaincode**</td>
    <td align="left">Creating preventive work orders.</td>
    <td align="center">302</td>
    <td align="center">3.3</td>
    <td align="center">747</td>
    <td align="center">8.2</td>
    <td align="center">1.33</td>
    <td align="center">335</td>
    <td align="center">3.7006/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">2.13</td>
  </tr>
  <tr>
    <td align="left">Creating corrective work orders.</td>
    <td align="center">407</td>
    <td align="center">4.5</td>
    <td align="center">553</td>
    <td align="center">6.1</td>
    <td align="center">1.21</td>
    <td align="center">364</td>
    <td align="center">4.0146/s</td>
    <td align="center">0.82</td>
    <td align="center">3</td>
    <td align="center">1.48</td>
  </tr>
  <tr>
    <td align="left">Proficiency Chaincode**</td>
    <td align="left">Creating training Module One.</td>
    <td align="center">397</td>
    <td align="center">4.4</td>
    <td align="center">600</td>
    <td align="center">6.6</td>
    <td align="center">0.99</td>
    <td align="center">451</td>
    <td align="center">4.9691/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">1.25</td>
  </tr>
  <tr>
    <td align="left" rowspan="2">Health & Safety Chaincode**</td>
    <td align="left">Creating incident and accident reports.</td>
    <td align="center">482</td>
    <td align="center">5.3</td>
    <td align="center">537</td>
    <td align="center">5.9</td>
    <td align="center">0.81</td>
    <td align="center">553</td>
    <td align="center">6.1057/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">1.00</td>
  </tr>
  <tr>
    <td align="left">Creating risk assessments.</td>
    <td align="center">320</td>
    <td align="center">3.5</td>
    <td align="center">93</td>
    <td align="center">1.0</td>
    <td align="center">1.25</td>
    <td align="center">357</td>
    <td align="center">3.9287/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">1.72</td>
  </tr>
  <tr>
    <td align="left" rowspan="3">Financial Chaincode**</td>
    <td align="left">Creating purchase order.</td>
    <td align="center">577</td>
    <td align="center">9.5</td>
    <td align="center">748</td>
    <td align="center">12</td>
    <td align="center">0.49</td>
    <td align="center">608</td>
    <td align="center">10.0534/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">1.96</td>
  </tr>
  <tr>
    <td align="left">Creating operational invoices.</td>
    <td align="center">524</td>
    <td align="center">8.7</td>
    <td align="center">205</td>
    <td align="center">3.4</td>
    <td align="center">0.49</td>
    <td align="center">608</td>
    <td align="center">10.0631/s</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">2.10</td>
  </tr>
  <tr>
    <td align="left">Creating invoices for additional work.</td>
    <td align="center">533</td>
    <td align="center">8.8</td>
    <td align="center">548</td>
    <td align="center">9</td>
    <td align="center">0.48</td>
    <td align="center">620</td>
    <td align="center">10.2148</td>
    <td align="center">0.00</td>
    <td align="center">0</td>
    <td align="center">2.39</td>
  </tr>
  <!-- Add rows here -->
</table>

Note:

* `*` represents ten (10) users.
*  `**` represents five (5) users.
*  `***` the 95th percentile represents a value below which 95% of the data points fall.

[Back on top](#end-of-test-summary)

___

What is REST API? [Read more](https://www.redhat.com/en/topics/api/what-is-a-rest-api).

[Back to the main page.](../../../README.md)
