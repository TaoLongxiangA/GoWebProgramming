package main

import (
	"html/template"
	"log"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("ch5/example5_10/t1.html", "ch5/example5_10/t2.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "hello")
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
