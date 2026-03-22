package main

import (
	"fmt"
	"net/http"

	"log"

	"github.com/brendandburns/topz/pkg/topz"
)

func main() {
	http.HandleFunc("/topz", topz.HandleTopz)
	http.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		fmt.Print("hello")
		res.Write([]byte("hello"))
	})
	fmt.Print("listening")

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
