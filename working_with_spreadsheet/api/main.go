package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi :)",
	})
}

func echoHandler(c *gin.Context) {
	var json map[string]interface{}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"err": "Invalid JSON"})
		return
	}
	c.JSON(200, json)
}

func ReError(error string, c *gin.Context) {
	c.JSON(400, gin.H{"err": error})
	return
}

func personHandler(c *gin.Context) {
	var person Person
	if err := c.BindJSON(&person); err != nil {
		ReError("Invalid JSON format", c)
		return
	} else if person.Name == "" {
		ReError("Name cannot be empty", c)
		return
	} else if person.Age <= 0 {
		ReError("Age must be greater than zero", c)
		return
	}
	c.JSON(200, person)
}

func main() {
	r := gin.Default()

	r.GET("/hello", helloHandler)
	r.POST("/echo", echoHandler)
	r.POST("/person", personHandler)

	r.Run()
}