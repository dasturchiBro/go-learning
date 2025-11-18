package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
)

type Post struct {
	Id uint  
	Title string
	Body string
	Rate float64
}

func (p Post) getPostRate() string {
	return fmt.Sprintf("The rate of the post is %.2f", p.Rate)
}

func (p *Post) setRate(newRate float64) {
	p.Rate = newRate
}

func home_page(w http.ResponseWriter, r *http.Request) {
	// posts := map[int]Post{
	// 	1: Post{1, "Go is super easy!", "Beginners in programming find it hard to learn programming languages, but there is ...", 4.5},
	// 	2: Post{1, "Go is super easy!", "Beginners in programming find it hard to learn programming languages, but there is ...", 4.5},
	// 	3: Post{1, "Go is super easy!", "Beginners in programming find it hard to learn programming languages, but there is ...", 4.5},
	// 	4: Post{1, "Go is super easy!", "Beginners in programming find it hard to learn programming languages, but there is ...", 4.5},
	// }

	posts := []Post{Post{1, "Go is super easy!", "Beginners in programming find it hard to learn programming languages, but there is ...", 4.5},
					Post{2, "Go is super easy!", "Beginners in programming find it hard to learn programming languages, but there is ...", 4.5},}
	tmpl, err := template.ParseFiles("templates/test.html")
	if err != nil {
		log.Fatal("Something went wrong: ", err)
	}
	tmpl.Execute(w, posts)
}

func blog_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World!!!</h1>")
}

func main() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/blog/", blog_page)
	http.ListenAndServe(":80", nil)


}