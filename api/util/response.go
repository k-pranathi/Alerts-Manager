package util

import "github.com/Jaya/walmart-api/api/models"

func GetAlertResponse(alerts []models.Alert) *models.AlertResponseBody {
	alertResponseBody := &models.AlertResponseBody{}

	for _, alert := range alerts {
		alertResponseBody.ServiceID = alert.ServiceID
		alertResponseBody.ServiceName = alert.ServiceName
		alertResponseBody.Alerts = append(alertResponseBody.Alerts, models.AlertResponse{
			AlertID:   alert.AlertID,
			Model:     alert.Model,
			AlertType: alert.AlertType,
			AlertTs:   alert.AlertTs,
			Severity:  alert.Severity,
			TeamSlack: alert.TeamSlack,
		})
	}

	return alertResponseBody
}
