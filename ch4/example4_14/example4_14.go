package main

import (
	"log"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "1",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "2",
		HttpOnly: true,
	}
	log.Println(c1)
	log.Println(c2)
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/setCookie", setCookie)
	server.ListenAndServe()
}
