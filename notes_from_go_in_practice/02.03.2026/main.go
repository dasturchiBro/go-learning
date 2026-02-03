package main

import (
	"fmt"
	"net/http"
	"time"
)

func greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Hello World!!!
		<script type="text/javascript">
            // Optionally force a reload, ignoring the browser cache
            window.location.reload(true);
        </script>`)
}


func count(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}

}

func main() {
	http.HandleFunc("/", greetings)
	http.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
    	fmt.Fprintf(w, `
	        <script type="text/javascript">
	            // Optionally force a reload, ignoring the browser cache
	            window.location.reload(true);
	        </script>
	    `)
	})
	go count(100)
	http.ListenAndServe("localhost:81", nil)
}