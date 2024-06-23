package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebhookPayload struct {
	Event     string `json:"event"`
	TimeStamp string `json:"time"`
	Payload   struct {
		Event struct {
			UUID string `json:"uuid"`
		} `json:"event"`
		Invitee struct {
			Email string `json:"email"`
		} `json:"invitee"`
	} `json:"payload"`
}

func main() {
	r := gin.Default()

	r.POST("/webhook", func(c *gin.Context) {
		var payload WebhookPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Webhook received: %+v\n", payload)

		// Handle the webhook payload here

		c.JSON(http.StatusOK, gin.H{"status": "Webhook received"})
	})

	r.Run(":8080") // Run on port 8080
}
