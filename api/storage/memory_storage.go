package storage

import (
	"github.com/alerts-manager/api/models"
	"github.com/alerts-manager/api/util"
)

// MemoryStorage ...
type MemoryStorage struct {
	alertStorage []models.Alert
}

// NewMemoryStorage ...
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		alertStorage: []models.Alert{},
	}
}

// SaveAlert ...
func (m *MemoryStorage) SaveAlert(alert models.Alert) error {
	m.alertStorage = append(m.alertStorage, alert)
	return nil
}

// GetAlerts ...
func (m *MemoryStorage) GetAlerts(alertQuery models.AlertQuery) ([]models.Alert, error) {
	filteredAlerts := m.alertStorage

	if alertQuery.ServiceID != "" {
		alerts, err := m.filterAlertsByServiceId(alertQuery.ServiceID, filteredAlerts)
		if err != nil {
			return nil, err
		}
		filteredAlerts = alerts
	}

	if alertQuery.StartTs != "" {
		alerts, err := m.filterAlertsByStartTime(alertQuery.StartTs, filteredAlerts)
		if err != nil {
			return nil, err
		}
		filteredAlerts = alerts
	}

	if alertQuery.EndTs != "" {
		alerts, err := m.filterAlertsByEndTime(alertQuery.EndTs, filteredAlerts)
		if err != nil {
			return nil, err
		}
		filteredAlerts = alerts
	}

	return filteredAlerts, nil
}

func (m *MemoryStorage) filterAlertsByStartTime(startTime string, alerts []models.Alert) ([]models.Alert, error) {
	filteredAlerts := []models.Alert{}

	startTimeEpoch, err := util.GetTimeFromEpoch(startTime)
	if err != nil {
		return nil, err
	}

	for _, alert := range alerts {
		alertTimeStamp, err := util.GetTimeFromEpoch(alert.AlertTs)
		if err != nil {
			return nil, err
		}

		if alertTimeStamp.After(*startTimeEpoch) {
			filteredAlerts = append(filteredAlerts, alert)
		}
	}

	return filteredAlerts, nil
}

func (m *MemoryStorage) filterAlertsByEndTime(endTime string, alerts []models.Alert) ([]models.Alert, error) {
	filteredAlerts := []models.Alert{}

	endTimeEpoch, err := util.GetTimeFromEpoch(endTime)
	if err != nil {
		return nil, err
	}

	for _, alert := range alerts {
		alertTimeStamp, err := util.GetTimeFromEpoch(alert.AlertTs)
		if err != nil {
			return nil, err
		}

		if alertTimeStamp.Before(*endTimeEpoch) {
			filteredAlerts = append(filteredAlerts, alert)
		}
	}

	return filteredAlerts, nil
}

func (m *MemoryStorage) filterAlertsByServiceId(serviceId string, alerts []models.Alert) ([]models.Alert, error) {
	filteredAlerts := []models.Alert{}

	for _, alert := range alerts {
		if alert.ServiceID == serviceId {
			filteredAlerts = append(filteredAlerts, alert)
		}
	}

	return filteredAlerts, nil
}
