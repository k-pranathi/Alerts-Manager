package handlers

import "github.com/Jaya/walmart-api/api/storage"

type AlertHandlers struct {
	storage storage.Storage
}

func NewAlertHandlers(storage storage.Storage) *AlertHandlers {
	return &AlertHandlers{
		storage: storage,
	}
}
