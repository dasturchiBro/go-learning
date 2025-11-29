package main

import (
	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "What's up?",
	})
}

func main() {
	router := gin.Default() 
	router.GET("/blog", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	router.GET("/", home)
	router.Run()
}


// 1. Full REST APIs
// User registration, login, JWT auth, CRUD operations, admin dashboards. The boring stuff every real engineer needs to master.

// 2. Microservices
// Split your system into small Go services: auth service, payment service, notification service. Gin makes them quick to spin up.

// 3. Real-time dashboards
// Use Gin to collect data, serve JSON, feed your frontend charts. You can track your runs, your apps’ logs, network traffic, whatever.

// 4. File upload systems
// Images, PDFs, profile pictures. Handle multipart forms. Validate, store, serve. Useful everywhere.

// 5. Webhooks
// Receive events from third-party services. Telegram bots. Payment systems. Anything that sends HTTP requests.

// 6. Reverse proxy / API gateway (simple version)
// Route requests, add middleware, log metrics, throttle requests. You’ll feel your brain growing.

// 7. Authentication systems
// Sessions, cookies, JWT, OAuth. The stuff weaker programmers avoid because it scares them.

// 8. A backend for your YouTube projects, your blog, your portfolio, whatever startup idea hits you at 2 AM
// Stop thinking and start building.