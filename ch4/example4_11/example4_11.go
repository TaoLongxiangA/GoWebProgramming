package main

import (
	"encoding/json"
	"net/http"
)

type post struct {
	User    string
	Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	post := &post{
		User:    "taolx",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/jsonExample", jsonExample)
	server.ListenAndServe()
}
