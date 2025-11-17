package main

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	var answer string
	fmt.Println("What to do?")
	fmt.Scan(&answer)
	fmt.Println(answer)
	fmt.Fprintf(w, answer)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/",home_page)
	http.ListenAndServe(":"+port, nil)
}