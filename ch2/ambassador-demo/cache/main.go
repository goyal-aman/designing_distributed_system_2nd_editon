package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

var store = make(map[string]string)
var mu sync.RWMutex

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in get")
	key := r.URL.Path[len("/get/"):]
	fmt.Println("key:", key, " store:", store, " store[key]:", store[key])
	mu.RLock()
	val, ok := store[key]
	mu.RUnlock()

	if !ok {
		http.Error(w, "not found", 404)
		return
	}

	w.Write([]byte(val))
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in set")

	path := r.URL.Path[len("/set/"):]

	parts := strings.SplitN(path, "/", 2)
	if len(parts) != 2 {
		http.Error(w, "invalid format", 400)
		return
	}

	key := parts[0]
	value := parts[1]

	fmt.Println("key:", key, "value:", value)

	mu.Lock()
	store[key] = value
	mu.Unlock()

	fmt.Println("store:", store)

	w.Write([]byte("ok"))
}
func main() {
	http.HandleFunc("/get/", getHandler)
	http.HandleFunc("/set/", setHandler)

	fmt.Println("Cache running on :8082")
	http.ListenAndServe(":8082", nil)
}
