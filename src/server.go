package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("running")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "css/style.css")
	})
	log.Fatal(http.ListenAndServe(":4000", nil))
}
