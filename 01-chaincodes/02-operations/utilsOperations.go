/*
 * Copyright 2022 IBM All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the 'License');
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an 'AS IS' BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

/*
To be added:
	1. Adding and updating the total number of operating hours so that the number of operating hours is taken
	from the previous day and added to the number of operating hours when the Operational log is created.

	2. Adding and updating the total number of passengers so that the number of passengers is taken
	from the previous day and added to the number of passengers when the Operational log is created.
*/

// Downtime stuct was removed as new smart contract was made, so-called "DowntimeEventContract"

type WeatherDriveStation struct {
	WeatherDriveStationCondition   string `json:"weather_drive_station_condition"`
	WeatherDriveStationTemperature string `json:"weather_drive_station_temperature"`
	WeatherDriveStationHumidity    string `json:"weather_drive_station_humidity"`
}

type WeatherReturnStation struct {
	WeatherReturnStationCondition   string `json:"weather_return_station_condition"`
	WeatherReturnStationTemperature string `json:"weather_return_station_temperature"`
	WeatherReturnStationHumidity    string `json:"weather_return_station_humidity"`
}

type PersonnelOnDutyDriveStation struct {
	PersonnelOnDutyDriveStationShiftManager             string `json:"personnel_on_duty_drive_station_shift_manager"`
	PersonnelOnDutyDriveStationMaintenanceTechnicianOne string `json:"personnel_on_duty_drive_station_maintenance_technician_one"`
	PersonnelOnDutyDriveStationMaintenanceTechnicianTwo string `json:"personnel_on_duty_drive_station_maintenance_technician_two"`
	PersonnelOnDutyDriveStationSystemOperatorOne        string `json:"personnel_on_duty_drive_station_system_operator_one"`
	PersonnelOnDutyDriveStationSystemOperatorTwo        string `json:"personnel_on_duty_drive_station_system_operator_two"`
}

type PersonnelOnDutyReturnStation struct {
	PersonnelOnDutyReturnStationMaintenanceTechnician string `json:"personnel_on_duty_return_station_maintenance_technician"`
	PersonnelOnDutyReturnStationSystemOperator        string `json:"personnel_on_duty_return_station_system_operator"`
}

type DailyOperationsLog struct {
	DocType               string                       `json:"dailyopslog"`
	DailyOpsLogID         string                       `json:"daily_ops_log_id"`
	Owner                 string                       `json:"owner"`
	WeatherDrive          WeatherDriveStation          `json:"weather_drive_station"`
	WeatherReturn         WeatherReturnStation         `json:"weather_return"`
	PersonnelOnDutyDrive  PersonnelOnDutyDriveStation  `json:"personnel_on_duty_drive_station"`
	PersonnelOnDutyReturn PersonnelOnDutyReturnStation `json:"personnel_on_duty_return_station"`
	OperationsStart       string                       `json:"operations_start"`
	OperationsEnd         string                       `json:"operations_end"`
	OperationHours        int                          `json:"operation_hours"`
	DailyOpsLogValidation string                       `json:"dailyopslog_validation"`
	NumberOfCarriers      int                          `json:"number_of_carriers"`
	NumberOfPassengers    int                          `json:"number_of_passengers"`
	TotalOperatingHours   int                          `json:"total_operating_hours"`
	AdditionalComments    string                       `json:"additional_comments"`
}
