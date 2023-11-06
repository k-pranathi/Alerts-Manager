package handlers

import (
	"github.com/Jaya/walmart-api/api/models"
	"github.com/gin-gonic/gin"
)

func (h *AlertHandlers) SaveAlerts(c *gin.Context) {
	requestBody := models.Alert{}
	if err := c.Bind(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if requestBody.AlertID == "" {
		c.JSON(400, gin.H{"error": "Alert ID is required"})
		return
	}

	if err := h.storage.SaveAlert(requestBody); err != nil {
		c.JSON(500, gin.H{"alert_id": requestBody.AlertID, "error": err.Error()})
	}

	c.JSON(201, gin.H{
		"alert_id": requestBody.AlertID,
		"error":    "",
	})
}
