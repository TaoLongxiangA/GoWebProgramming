package main

import (
	"log"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("ch5/example5_3/temp.html") //parse template file
	if err != nil {
		log.Println(err)
	}
	rand.Seed(time.Now().Unix())
	temp.Execute(w, rand.Intn(10) > 5)
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
