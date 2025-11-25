package handlers

import (
	"net/http"
	"lumina/db"
	"strings"
	"strconv"
	"lumina/app"
)

func SplitId(path string) (int, bool) {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return 0, true
	}
	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, true
	}
	return id, false
}

func Post(w http.ResponseWriter, r *http.Request) {
	id, err := SplitId(r.URL.Path)
	if err != false {
		http.NotFound(w, r)
	}
	post, isTherePost := db.GetPostById(id)
	if !isTherePost {
		http.NotFound(w, r)
		return
	}
	if err := app.Tmpl.ExecuteTemplate(w, "post", post); err != nil {
		http.Error(w, err.Error(), 500)
	}
}