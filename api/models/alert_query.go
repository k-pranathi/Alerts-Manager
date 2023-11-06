package models

// AlertQuery is a struct that represents a query for alerts
type AlertQuery struct {
	ServiceID string `json:"service_id"`
	StartTs   string `json:"start_ts"`
	EndTs     string `json:"end_ts"`
}
