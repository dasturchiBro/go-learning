package handlers

import (
	"lumina/models"
	"net/http"
	"lumina/db"
	"time"
	"strconv"
)

func SaveUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if _, err := r.Cookie("session_id"); err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "unable to parse form", 500)
		return
	}
	user := models.User{
		Name: r.FormValue("name"),
		Email: r.FormValue("email"),
		Hash_pass: r.FormValue("password"),
		Role: "user",
	}
	UserId, err := db.InsertUser(user)
	if err == "email" {
		http.Error(w, "User with this email already exists!", 500)
	} else if err == "error" {
		http.Error(w, "something went wrong, please try again", 500)
	}
	expiry := time.Now().Add(24 * time.Hour)
	session := models.Session{
		SessionId: user.Email,
		UserId: UserId,
		ExpiresAt: expiry,
	}
	err2 := db.InsertSession(session)
	if err2 != nil {
		http.Error(w, "Something went wrong. Please try again!", 500)
	}
	http.SetCookie(w, &http.Cookie{
		Name: "session_id",
		Value: strconv.Itoa(session.UserId) + session.SessionId,
		Path: "/",
		HttpOnly: false,
		Secure: false,
		Expires: expiry,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}