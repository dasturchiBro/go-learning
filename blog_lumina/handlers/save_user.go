package handlers

import (
	"lumina/models"
	"net/http"
)

func SaveUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "unable to parse form", 500)
		return
	}
	user := models.User{
		Name: r.FormValue("name")
		Email: r.FormValue("email")
		Hash_pass: r.FormValue("password")
	}
	if err := InsertUser(user); err == "email" {
		http.Error(w, "User with this email already exists!", 500)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}