package main

type RiskAssessment struct {
	DocType                       string `json:"docType"`                          // This is the discriminator, "risk-assessment"
	RiskAssessmentID              string `json:"risk_assessment_id"`               // This is the key, e.g. "risk-assessment-HEX"
	RiskAssessmentDate            string `json:"risk_assessment_date"`             // This is the date of the risk assessment was created by the HS responsible person
	RiskAssessmentDateNextReview  string `json:"risk_assessment_date_next_review"` // This is the date of the next review
	RiskAssessmentActivity        string `json:"risk_assessment_activity"`         // This is the activity being assessed
	RiskAssessmentCreatedBy       string `json:"risk_assessment_created_by"`       // This is the person who created the assessment
	RiskAssessmentAssessedBy      string `json:"risk_assessment_assessed_by"`      // This is the person who did the assessment
	RiskAssessmentApprovedBy      string `json:"risk_assessment_approved_by"`      // This is the person who approved the assessment
	RiskAssessmentHazardListOne   string `json:"risk_assessment_hazard_list_one"`
	RiskAssessmentHazardListTwo   string `json:"risk_assessment_hazard_list_two"`
	RiskAssessmentHazardListThree string `json:"risk_assessment_hazard_list_three"`
	RiskAssessmentHazardListFour  string `json:"risk_assessment_hazard_list_four"`
	RiskAssessmentHazardListFive  string `json:"risk_assessment_hazard_list_five"`
}

type AccidentIncidentReport struct {
	DocType                                      string `json:"docType"`                                            // This is the discriminator, "accident-incident-report"
	AccidentIncidentReportID                     string `json:"accident_incident_report_id"`                        // This is the key, e.g. "accident-incident-report-HEX"
	SubjectAccidentIncidentReport                string `json:"subject_accident_incident_report"`                   // This is the subject of the accident or incident report
	AccidentIncidentReportDate                   string `json:"accident_incident_report_date"`                      // This is the date of the accident or incident report
	AccidentIncidentReportTimeStart              string `json:"accident_incident_report_time_start"`                // This is the time of the accident or incident report
	AccidentIncidentReportTimeEnd                string `json:"accident_incident_report_time_end"`                  // This is the time of the accident or incident report
	AccidentIncidentReportLocation               string `json:"accident_incident_report_location"`                  // This is the location of the accident or incident report
	AccidentIncidentReportHSAspects              string `json:"accident_incident_report_hs_aspects"`                // This is the HS aspects of the accident or incident report
	AccidentIncidentReportClassification         string `json:"accident_incident_report_classification"`            // This is the classification of the accident or incident report
	AccidentIncidentReportDescription            string `json:"accident_incident_report_description"`               // This is the description of the accident or incident report
	AccidentIncidentReportImmediateAction        string `json:"accident_incident_report_immediate_action"`          // This is the immediate action of the accident or incident report
	AccidentIncidentReportFollowUpActions        string `json:"accident_incident_report_follow_up_actions"`         // This is the follow up action of the accident or incident report
	AccidentIncidentReportCreatedBy              string `json:"accident_incident_report_created_by"`                // This is the person who created the accident or incident report
	AccidentIncidentReportValidatedBy            string `json:"accident_incident_report_validated_by"`              // This is the person who validated the accident or incident report
	AccidentIncidentReportStatus                 string `json:"accident_incident_report_status"`                    // This is the status of the accident or incident report
	AccidentIncidentReportNumberOfPeopleInvolved int    `json:"accident_incident_report_number_of_people_involved"` // This is the number of people involved in the accident or incident report
	AccidentIncidentReportPersonInjured          string `json:"accident_incident_report_person_injured"`            // This is the list of people injured in the accident or incident report
	AccidentIncidentReportWitnesses              string `json:"accident_incident_report_witnesses"`                 // This is the list of witnesses in the accident or incident report
	AccidentIncidentReportEventOverview          string `json:"accident_incident_report_event_overview"`            // This is the list of event overviews in the accident or incident report
	AccidentIncidentReportLineOfCommunication    string `json:"accident_incident_report_line_of_communication"`     // This is the list of line of communication in the accident or incident report
}
