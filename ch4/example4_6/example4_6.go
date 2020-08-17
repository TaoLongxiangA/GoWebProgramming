package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	//r.ParseMultipartForm(1024)
	//fileHeader := r.MultipartForm.File["upload"][0]
	//file, err := fileHeader.Open()
	file, _, err := r.FormFile("upload")
	if err == nil {
		date, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(date))
		}
	}
	//fmt.Fprintln(w, fileHeader)
	fmt.Fprintln(w, r.Form)
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
