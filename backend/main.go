package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {

	name, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hostname: %q\n", html.EscapeString(name))
		for name, headers := range r.Header {
			for _, h := range headers {
				fmt.Fprintf(w, "%v: %v\n", name, h)
			}
		}

	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
