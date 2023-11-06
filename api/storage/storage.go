package storage

import "github.com/alerts-manager/api/models"

// Storage ...
type Storage interface {
	SaveAlert(alert models.Alert) error
	GetAlerts(alertQuery models.AlertQuery) ([]models.Alert, error)
}
