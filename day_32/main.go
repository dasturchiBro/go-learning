package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)


type Person struct {
	Name string
	Age int
	Projects []string
}

func (p Person) getInfo() string {
	return fmt.Sprintf("Name: %s \nAge: %d\n", p.Name, p.Age)
}

func connect() *pgxpool.Pool {
	db, err := pgxpool.New(context.Background(), "postgres://postgres:12345@localhost:5432/fastapi_new")
	if err != nil {
		log.Fatal(err)
	}
	return db
}


func home_page(w http.ResponseWriter, r *http.Request) {
	john := Person{"John Doe", 19, []string{"IoT Smart Home", "Green House Project", "Learning Program"}}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, john)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":80", nil)
}

// go get github.com/jackc/pgx/v5