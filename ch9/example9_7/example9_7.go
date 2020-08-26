package main

import (
	"fmt"
	"time"
)

func printNum1(w chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	w <- true
}

func printLetters1(w chan bool) {
	for i := 'A'; i < 10+'A'; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	w <- true
}

func main() {
	chan1, chan2 := make(chan bool), make(chan bool)
	go printNum1(chan1)
	go printLetters1(chan2)
	<-chan1
	<-chan2
}
