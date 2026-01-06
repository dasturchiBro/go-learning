package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	// "github.com/jackc/pgx/v5/pgxpool"
	// "context"
	// "log"
)

type Book struct {
	Id int `json:"id"`
	Name string `json:"name"`
	NumberOfPages int `json:"numberofpages"`
	Author string `json:"author"`
}

// const DB *pgxpool.Pool


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

func getBookByID(c *gin.Context, books *([]Book)) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"message": "bad request",
		})
	}

	for _, book := range *books {
		if book.Id == id {
			c.IndentedJSON(200, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "book not found",
	})
}


// func AppInit() {
// 	var err error
// 	DB, err = pgxpool.New(context.Background(), "postgres://postgres:12345@localhost:5432/lumina")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

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
	router.GET("/:id", func(c *gin.Context) {
		getBookByID(c, &books)
	})
	router.POST("/", func(c *gin.Context) {
		postBooks(c, &books)
	})
	router.Run()
}