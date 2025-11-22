package main

import (
	// "fmt"
	"log"
	"net/http"
	"html/template"
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
	"time"
	"strconv"
)

var (
	db *pgxpool.Pool 
	tmpl *template.Template
)

func initApp() {
	db = connect()
	tmpl = template.Must(template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html", "templates/create.html"))
}


type Post struct {
	Id int
	Title string  
	Body string  
	Date time.Time
}

func connect() *pgxpool.Pool {
	ctx := context.Background()
	db, err := pgxpool.New(ctx, "postgres://postgres:12345@localhost:5432/blog") 
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rows, err := db.Query(ctx, "SELECT id, title, body, date FROM posts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.Id, &p.Title, &p.Body, &p.Date); err != nil {
			log.Fatal(err)
		}
		posts = append(posts, p)
	}
	if err := tmpl.ExecuteTemplate(w, "index", posts); err != nil {
		log.Fatal(err)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "create", ""); err != nil {
		log.Fatal(err)
	}
}

func save_post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse.", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "Unable to parse.", http.StatusBadRequest)
		return
	}
	post := Post{
		id,
		r.FormValue("title"),
		r.FormValue("body"),
		time.Now(),
	}
	if err := insertPost(post); err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		log.Print(err)
	}
	http.Redirect(w,r,"/", http.StatusSeeOther)
}

func insertPost(p Post) error {
	_, err := db.Exec(context.Background(), "INSERT INTO posts (id, title, body, date) VALUES ($1, $2, $3, $4)", p.Id, p.Title, p.Body, p.Date)
	return err
}

func handleFunc() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create/", create)
	http.HandleFunc("/save_post", save_post)
	http.ListenAndServe(":80", nil)
}

func main() {
	initApp()
	handleFunc()
}