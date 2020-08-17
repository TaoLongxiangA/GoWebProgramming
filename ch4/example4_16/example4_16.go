package main

import (
	"fmt"
	"net/http"
)

func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "cannot get the first_cookie")
	}
	cookies := r.Cookies()
	fmt.Fprintln(w, cookies)
	fmt.Fprintln(w, cookie)
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/getCookie", getCookie)
	server.ListenAndServe()
}
