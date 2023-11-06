package storage

import "github.com/Jaya/walmart-api/api/models"

type Storage interface {
	SaveAlert(alert models.Alert) error
	GetAlerts(alertQuery models.AlertQuery) ([]models.Alert, error)
}
