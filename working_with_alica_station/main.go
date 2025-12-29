package main

import (
	"github.com/gin-gonic/gin"
)

type Req struct {
	Request struct {
		Type string `json:"type"`
		Command string `json:"command"`
	} `json:"request"`
	Session struct {
		New       bool   `json:"new"`
		SessionID string `json:"session_id"`
		MessageID int    `json:"message_id"`
		UserID    string `json:"user_id"`
	} `json:"session"`
}

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		var req Req
		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		text := "You studied 5 hours today."
		if req.Request.Type == "AccountLinkingRequest" {
			text = "Account linking not supported."
		}

		if req.Request.Command == "hi" {
			text = "Hello"
		}

		c.JSON(200, gin.H{
			"version": "1.0",
			"session": req.Session,
			"response": gin.H{
				"text":        text,
				"end_session": false,
			},
		})
	})

	r.Run(":9000")
}
