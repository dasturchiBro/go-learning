package main

import (
	"time"
	"net/http"
	"html/template"
	"log"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Post struct {
	Id int
	Title string
	Body string
	Date time.Time
}



func connect() *pgxpool.Pool {
	db, err := pgxpool.New(context.Background(), "postgres://postgres:12345@localhost:5432/blog")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func home_page(w http.ResponseWriter, r *http.Request) {
	db := connect()
	rows, err := db.Query(context.Background(), "SELECT id, title, body, date FROM posts")
	var posts []Post  
	for rows.Next() {
		var p Post 
		if err := rows.Scan(&p.Id, &p.Title, &p.Body, &p.Date); err != nil {
			log.Fatal(err)
		}
		posts = append(posts, p)
	}

	tmpl := template.Must(template.ParseFiles(
		"./templates/index.html", 
		"./templates/header.html", 
		"./templates/footer.html",
	))
	err = tmpl.ExecuteTemplate(w, "index.html", posts)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":80", nil)
}