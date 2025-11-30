package main

import (
	"github.com/gin-gonic/gin"
)

type Book struct {
	Id int `json:"id"`
	Name string `json:"name"`
	NumberOfPages int `json:"numberofpages"`
	Author string `json:"author"`
}



func getAllBooks(c *gin.Context, books []Book) {
	
	c.IndentedJSON(200, books)
}

func postBooks(c *gin.Context, books *([]Book)) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	*books = append(*books, newBook)
	c.IndentedJSON(200, books)
}

func main() {
	books := []Book{
		{
			1,
			"Sherlock Holmes",
			514,
			"Unknown",
		},
		{
			2,
			"Atomic Habits",
			214,
			"James Clear",
		},
		{
			3,
			"Anna Karenina",
			1008,
			"Leo Tosltoy",
		},
	}
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		getAllBooks(c, books)
	})
	router.POST("/", func(c *gin.Context) {
		postBooks(c, &books)
	})
	router.Run()
}