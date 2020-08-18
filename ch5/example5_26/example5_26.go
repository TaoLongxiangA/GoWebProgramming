package main

import (
	"html/template"
	"math/rand"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(10)
	if num > 5 {
		t, _ := template.ParseFiles("ch5/example5_26/layout.html", "ch5/example5_26/red.html")
		t.ExecuteTemplate(w, "layout", "") //parse template
	} else {
		t, _ := template.ParseFiles("ch5/example5_26/layout.html")
		t.ExecuteTemplate(w, "layout", "")
	}
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
