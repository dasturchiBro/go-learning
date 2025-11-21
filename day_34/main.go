package main 

import (
	"net/http"
	"html/template"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"context"
	"fmt"
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

func index(w http.ResponseWriter, r *http.Request) {
	db := connect()
	rows, err := db.Query(context.Background(), "SELECT id, title, body, date FROM posts")
	if err != nil {
		log.Fatal(err)
	}
	var posts []Post
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.Id, &p.Title, &p.Body, &p.Date); err != nil {
			log.Fatal(err)
		}
		posts = append(posts, p)
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html"))
	if err := tmpl.ExecuteTemplate(w, "index", posts); err != nil {
		log.Fatal(err)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	p := Post{1, "Test", "test", time.Now()}
	tmpl := template.Must(template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html"))
	if err := tmpl.ExecuteTemplate(w, "create", p); err != nil {
		log.Fatal(err)
	}
}

func save_post(w http.ResponseWriter, r *http.Request) {
	db := connect() 
	title := r.FormValue("title")
	body := r.FormValue("body")
	id := r.FormValue("id")
	date := time.Now()
	fmt.Println(title, body, id)
	fmt.Println("Hello")
	if _, err := db.Query(context.Background(), fmt.Sprintf("INSERT INTO \"posts\" (\"id\", \"title\", \"body\", \"date\") VALUES('%v', '%v', '%v', '%v'::DATE)", id, title, body, date)); err != nil {
		log.Fatal(err)
	}
}

func handleFunc() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create/", create)
	http.HandleFunc("/save_post/", save_post)
	http.ListenAndServe(":80", nil)
}

func main() {
	handleFunc()
}