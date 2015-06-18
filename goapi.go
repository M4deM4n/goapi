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
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
