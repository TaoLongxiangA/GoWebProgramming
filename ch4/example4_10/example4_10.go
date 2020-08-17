package main

import (
	"net/http"
)

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location","http://www.baidu.com")
	w.WriteHeader(302)
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/headerExample", headerExample)
	server.ListenAndServe()
}
