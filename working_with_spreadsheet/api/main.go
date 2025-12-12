package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
	"time"
	"encoding/json"
	"io"
	"bytes"
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

func ReError(msg string, c *gin.Context) {
	c.JSON(400, gin.H{"err": msg})
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

func userHandler(c *gin.Context) {
	userName := c.Param("name")
	if userName == "" {
		ReError("Name cannot be empty", c)
		return
	}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Hello, %s", userName),
	})
}

func searchHandler(c *gin.Context) {
	term := c.Query("term")
	if term == "" {
		ReError("term is required", c)
		return
	}
	limit := c.DefaultQuery("limit", "5")
	limit_number, err := strconv.Atoi(limit)
	if err != nil {
		limit_number = 5
	}
	c.JSON(200, gin.H{"term": term, "limit": limit_number})
}

func filterHandler(c *gin.Context) {
	min := c.Query("min")
	max := c.Query("max")
	if min == "" && max == "" {
		ReError("at least one parameter is required", c)
		return
	} else if min == "" && max != "" {
		maxN, err := strconv.Atoi(max)
		if err != nil {
			ReError("max should be an integer", c)
			return
		}
		c.JSON(200, gin.H{"max": maxN})
		return
	} else if min != "" && max == "" {
		minN, err := strconv.Atoi(min)
		if err != nil {
			ReError("min should be an integer", c)
			return
		}
		c.JSON(200, gin.H{"min": minN})
		return
	} else {
		minN, err := strconv.Atoi(min)
		maxN, err2 := strconv.Atoi(max)
		if err != nil || err2 != nil {
			ReError("min and max values must be integers", c)
			return
		} else if minN > maxN {
			ReError("min cannot be greater than max", c)
			return
		} else {
			c.JSON(200, gin.H{
				"min": minN,
				"max": maxN,
			})
		}
	}
}

func Logger(c *gin.Context) {
	method := c.Request.Method 
	path := c.Request.URL.Path
	startTime := time.Now()
	c.Next()
	latency := time.Since(startTime)
	fmt.Println(method, " ", path, " - ", latency.String())
}

func AuthMiddleware(c *gin.Context) {
	AuthToken := c.GetHeader("X-Auth-Token")
	if AuthToken == "" {
		c.JSON(401, gin.H{
			"error": "Unauthorized: missing auth token",
		})
		c.Abort()
		return
	}
	c.Next()
}

var requestsMap = make(map[string][]time.Time) 

func RateMiddleware(c *gin.Context) {
	ip := c.ClientIP()
	_, ok := requestsMap[ip]
	now := time.Now()
	threshold := 60 * time.Second
	if !ok {
		requestsMap[ip] = []time.Time{now}
	} else {
		filtered := []time.Time{} 
		for _, value := range requestsMap[ip] {
			if now.Sub(value) <= threshold {
				filtered = append(filtered, value)
			}
		}
		requestsMap[ip] = filtered
		if len(requestsMap[ip]) >= 5 {
			c.JSON(429, gin.H{
				"error": "Rate limit exceeded",
			})
			c.Abort()
			return
		} else {
			requestsMap[ip] = append(requestsMap[ip], time.Now())
		}
	}
	c.Next()
}

func ValidatorMiddleware(c *gin.Context) {
	if c.Request.Method != "POST" && c.Request.Method != "PUT" {
		c.Next()
		return
	}
	data, err := c.GetRawData()
	if err != nil {
		ReError("Invalid JSON format during validation", c)
		c.Abort()
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(data))
	var dat map[string]interface{}
	if err := json.Unmarshal(data, &dat); err != nil {
		ReError("Invalid JSON format during validation", c)
		c.Abort()
		return
	}

	allowedSet := map[string]bool{"name": true, "age": true}
	for key := range dat {
		if !allowedSet[key] {
			ReError(fmt.Sprintf("Unexpected field: %v", key), c)
			c.Abort()
			return
		}
	}

	value, ok := dat["name"]
	if !ok {
		ReError("Field name doesn't exist", c)
		c.Abort()
		return
	}
	_, isString := value.(string)
	if !isString {
		ReError(fmt.Sprintf("Wrong type error: %v should be a string", value), c)
		c.Abort()
		return
	}

	age, ok := dat["age"]
	if !ok {
		ReError("Field age doesn't exist", c)
		c.Abort()
		return
	}
	_, isNumber := age.(float64)
	if !isNumber {
		ReError(fmt.Sprintf("Wrong type error: %v should be a number", age), c)
		c.Abort()
		return
	}
	c.Set("body", dat)
	c.Next()
}


func main() {
	r := gin.Default()
	
	r.Use(Logger)

	r.POST("/echo", echoHandler)
	r.GET("/hello", helloHandler)


	v1 := r.Group("/v1")
	v1.Use(RateMiddleware)
	// v1.Use(AuthMiddleware)
	v1.Use(ValidatorMiddleware)
	{
		v1.POST("/person", personHandler)
		v1.GET("/user/:name", userHandler)
		v1.GET("/search", searchHandler)
		v1.GET("/filter", filterHandler)
	}

	r.Run()
}