package main

import (
	"fmt"
	"log"
	"net/http"
)


func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	fmt.Println(r)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w,"This is GET method.")
	} else if r.Method == http.MethodPost {
		fmt.Fprintf(w,"This is POST method.")
	}
	fmt.Println("Something has happened!")
}

func main() {
	mux := http.NewServeMux()

	http.HandleFunc("/", homeHandler)
	mux.HandleFunc("/api", apiHandler)

	fmt.Println("Server starting on port 80...")
	log.Fatal(http.ListenAndServe(":80", nil))

}