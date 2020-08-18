package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	log.Println("func formatDate")
	layout := "2020-09-01"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	log.Println("func process")
	funMap := template.FuncMap{"fdate": formatDate}
	t := template.New("ch5/example5_14/dateFormat.html").Funcs(funMap)
	t, _ = t.ParseFiles("ch5/example5_14/dateFormat.html")
	log.Println(time.Now())
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
