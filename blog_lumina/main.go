package main

import (
	"net/http"
	"lumina/handlers"
	"lumina/db"
	"html/template"
	"lumina/app"
	"log"
)


func AppInit() {
	var err error
	app.DB = db.Connect()

	app.Tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	AppInit()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/post/{id}", handlers.Post)
	http.ListenAndServe(":80", nil)
}
