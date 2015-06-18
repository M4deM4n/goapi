package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome to my Go API.")
}

func main() {
	http.HandleFunc("/", homeHandler)
	log.Println("Attempting to listen on port 8080", port)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
