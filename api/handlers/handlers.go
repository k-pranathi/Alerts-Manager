package handlers

import "github.com/alerts-manager/api/storage"

// AlertHandlers ...
type AlertHandlers struct {
	storage storage.Storage
}

// NewAlertHandlers ...
func NewAlertHandlers(storage storage.Storage) *AlertHandlers {
	return &AlertHandlers{
		storage: storage,
	}
}
