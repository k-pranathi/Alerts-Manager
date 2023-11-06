package main

import (
	"fmt"

	"github.com/alerts-manager/api/handlers"
	"github.com/alerts-manager/api/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	memoryStorage := storage.NewMemoryStorage()
	handlers := handlers.NewAlertHandlers(memoryStorage)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/alerts", handlers.GetAlerts)
	r.POST("/alerts", handlers.SaveAlerts)

	fmt.Println("Running Alert System API")
	err := r.Run()
	if err != nil {
		fmt.Println("Error starting server")
	}

}
