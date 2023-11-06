package models

type Alert struct {
	AlertID     string `json:"alert_id"`
	ServiceID   string `json:"service_id"`
	ServiceName string `json:"service_name"`
	Model       string `json:"model"`
	AlertType   string `json:"alert_type"`
	AlertTs     string `json:"alert_ts"`
	Severity    string `json:"severity"`
	TeamSlack   string `json:"team_slack"`
}
