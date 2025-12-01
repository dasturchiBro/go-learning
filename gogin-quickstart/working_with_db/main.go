package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
	"log"
	"strconv"
	"time"
)

type Post struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Category string `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

var DB *pgxpool.Pool

func AppInit() {
	var err error
	DB, err = pgxpool.New(context.Background(), "postgres://postgres:12345@localhost:5432/lumina")
	if err != nil {
		log.Fatal(err)
	}
}

func getPosts(c *gin.Context) {
	rows, err := DB.Query(context.Background(), "SELECT id, title, body, category, created_at FROM posts")
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"message": "something went wrong, please try again",
		})
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		var post Post 
		rows.Scan(&post.Id, &post.Title, &post.Body, &post.Category, &post.CreatedAt)
		posts = append(posts, post)
	}
	c.IndentedJSON(http.StatusOK, posts)
}

func getPostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(500, gin.H{
			"message":"make sure the id is a number",
		})
	}
	var post Post
	err = DB.QueryRow(context.Background(), "SELECT id, title, body, category, created_at FROM posts WHERE id = $1", id).Scan(&post.Id, &post.Title, &post.Body, &post.Category, &post.CreatedAt)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "post not found",
		})
		return
	}
	c.IndentedJSON(200, post)
}

func insertPost(c *gin.Context) {

}

func main() {
	AppInit()
	router := gin.Default()
	router.GET("/posts", getPosts)
	router.GET("/posts/:id", getPostById)
	router.POST("/posts", insertPost)
	router.Run("localhost:80")
}