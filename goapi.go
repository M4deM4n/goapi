package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome to my Go API.")
}

func authHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Authentication endpoint")
}

func main() {
	http.HandleFunc("/", homeHandler)
	log.Println("Attempting to listen on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
