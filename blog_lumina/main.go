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

	http.HandleFunc("/", handlers.Index)
	http.ListenAndServe(":80", nil)
}
