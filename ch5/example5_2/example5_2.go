package main

import (
	"log"
	"net/http"
	"text/template"
)

func process(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("ch5/example5_2/temp.html") //parse template file
	if err != nil {
		log.Println(err)
	}
	temp.Execute(w, "hello,world!")
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
