package handlers

import (
	"net/http"
	"lumina/db"
	"lumina/app"
)


func Index(w http.ResponseWriter, r *http.Request) {
	posts := db.GetPosts()
	if err := app.Tmpl.ExecuteTemplate(w, "index", posts); err != nil {
		http.Error(w, err.Error(), 500)
	}
}