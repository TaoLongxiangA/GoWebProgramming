package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello,world!")
	c1 := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c1)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	//c1 := r.Header["Cookie"]
	//fmt.Fprintln(w, c1)
	c2, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash",
			Expires: time.Unix(1, 0),
			MaxAge:  -1,
		}
		http.SetCookie(w, &rc)
		c, _ := base64.URLEncoding.DecodeString(c2.Value)
		fmt.Fprintln(w, string(c))
	}
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/setMessage", setMessage)
	http.HandleFunc("/showMessage", showMessage)
	server.ListenAndServe()
}
