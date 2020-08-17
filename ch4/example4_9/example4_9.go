package main

import (
	"fmt"
	"net/http"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := "l am a SDE"
	w.Write([]byte(str))
}

func writeHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w,"No such service ,try next door")
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/headerExample", writeExample)
	http.HandleFunc("/writeHeader", writeHeader)
	server.ListenAndServe()
}
