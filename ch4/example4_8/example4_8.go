package main

import (
	"net/http"
)

func write(w http.ResponseWriter, r *http.Request) {
	str := "l am a SDE"
	w.Write([]byte(str))
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/write", write)
	server.ListenAndServe()
}
