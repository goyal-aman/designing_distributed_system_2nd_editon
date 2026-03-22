package main

import (
	"fmt"
	"io"
	"net/http"
)

// curl --request GET --url http://ambassador.default.svc.cluster.local:8081
const ambassadorURL = "http://ambassador.default.svc.cluster.local:8081"

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	var resp *http.Response
	var err error

	if r.Method == http.MethodGet {
		resp, err = http.Get(ambassadorURL + "/get/" + path)
	} else {
		resp, err = http.Post(ambassadorURL+"/set/"+path, "text/plain", nil)
	}

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	w.Write(body)
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("App running on :8080")
	http.ListenAndServe(":8080", nil)
}
