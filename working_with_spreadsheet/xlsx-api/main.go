package main

import (
	"github.com/gin-gonic/gin"
	"xlsx/handlers"
)

func main() {
	r := gin.Default()
	r.Static("/downloads", "./downloads")
	v1 := r.Group("/v1")
	{
		v1.POST("/xlsx", handlers.XLSXHandler)
	}

	r.Run()
}