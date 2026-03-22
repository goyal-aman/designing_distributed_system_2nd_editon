package main

import (
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"net/http"
	"strings"
)

var cachePods = []string{
	"http://cache-0.cache.default.svc.cluster.local:8082",
	"http://cache-1.cache.default.svc.cluster.local:8082",
	"http://cache-2.cache.default.svc.cluster.local:8082",
}

func pickNode(key string) string {
	h := fnv.New32a()
	h.Write([]byte(key))
	return cachePods[int(h.Sum32())%len(cachePods)]
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/get/"):]
	node := pickNode(key)
	log.Println("in get, path:", key)

	resp, err := http.Get(node + "/get/" + key)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	w.Write(body)
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/set/"):]
	key := strings.Split(path, "/")[0]
	log.Println("in set, path:", key)
	node := pickNode(key)

	resp, err := http.Post(node+"/set/"+path, "text/plain", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	w.Write(body)
}

func main() {
	http.HandleFunc("/get/", getHandler)
	http.HandleFunc("/set/", setHandler)

	fmt.Println("Ambassador running on :8081")
	http.ListenAndServe(":8081", nil)
}
