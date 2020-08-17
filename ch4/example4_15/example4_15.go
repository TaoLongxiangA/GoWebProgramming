package main

import (
	"fmt"
	"net/http"
)

func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie := r.Header.Get("Cookie")
	fmt.Fprintln(w, cookie)
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/getCookie", getCookie)
	server.ListenAndServe()
}
