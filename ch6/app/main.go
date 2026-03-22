package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/ready", ready)
	fmt.Println("listening on 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("err starting server at 8080", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("hello")
	fmt.Fprintf(w, "hello world")
}

func ready(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ready")
}
