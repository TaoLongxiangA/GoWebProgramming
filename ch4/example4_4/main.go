package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	//log.Println(len(r.Form))
	//log.Println(len(r.PostForm))
	//fmt.Fprintln(w, r.Form)
	//fmt.Fprintln(w, r.PostForm)
	//r.ParseForm()
	//log.Println(len(r.Form))
	//log.Println(len(r.PostForm))
	//fmt.Fprintln(w, r.Form)
	//fmt.Fprintln(w, r.PostForm)

	fmt.Fprintln(w, r.PostFormValue("hello"))
	fmt.Fprintln(w, r.PostFormValue("post"))

	fmt.Fprintln(w, r.FormValue("hello"))
	fmt.Fprintln(w, r.FormValue("post"))

	//fmt.Fprintln(w, "next line is MultipartForm")
	//r.ParseMultipartForm(1024)
	//fmt.Fprintln(w, r.MultipartForm)
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8081"}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
