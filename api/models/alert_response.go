package models

// AlertResponseBody ...
type AlertResponseBody struct {
	ServiceID   string          `json:"service_id"`
	ServiceName string          `json:"service_name"`
	Alerts      []AlertResponse `json:"alerts"`
}

// AlertResponse ...
type AlertResponse struct {
	AlertID   string `json:"alert_id"`
	Model     string `json:"model"`
	AlertType string `json:"alert_type"`
	AlertTs   string `json:"alert_ts"`
	Severity  string `json:"severity"`
	TeamSlack string `json:"team_slack"`
}
