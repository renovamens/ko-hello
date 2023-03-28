package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	log.Println("start http server...")

	http.HandleFunc("/api/hello", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("hello, received a request")
	fmt.Fprintf(w, "Hello, this is a http service for first ko build.")
}
