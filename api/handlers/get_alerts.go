package handlers

import (
	"github.com/alerts-manager/api/models"
	"github.com/alerts-manager/api/util"
	"github.com/gin-gonic/gin"
)

// GetAlerts ...
func (h *AlertHandlers) GetAlerts(c *gin.Context) {
	queryParam := c.Request.URL.Query()
	alertQuery := models.AlertQuery{}

	if serviceID := queryParam.Get("service_id"); serviceID != "" {
		alertQuery.ServiceID = serviceID
	}

	if endTs := queryParam.Get("end_ts"); endTs != "" {
		alertQuery.EndTs = endTs
	}

	if startTs := queryParam.Get("start_ts"); startTs != "" {
		alertQuery.StartTs = startTs
	}

	alerts, _ := h.storage.GetAlerts(alertQuery)

	alertResponse := util.GetAlertResponse(alerts)

	c.JSON(200, gin.H{"alerts": alertResponse})
}
