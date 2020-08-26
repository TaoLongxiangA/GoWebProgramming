package main

import (
	"fmt"
	"log"
)

func callerA(c chan string) {
	c <- "Hello,world!"
	close(c)
}

func callerB(c chan string) {
	c <- "Halo,Mound!"
	close(c)
}

func main() {
	chan1, chan2 := make(chan string), make(chan string)
	go callerA(chan1)
	go callerB(chan2)
	var msg string
	ok1, ok2 := true, true
	for ok1 || ok2 {
		log.Println(ok1)
		log.Println(ok2)
		select {
		case msg, ok1 = <-chan1:
			log.Println("chan1 start:", ok1)
			if ok1 {
				fmt.Printf("%s from A\n", msg)
			}
			log.Println("chan1 end:", ok1)
		case msg, ok2 = <-chan2:
			log.Println("chan2 start:", ok1)
			if ok2 {
				fmt.Printf("%s from B\n", msg)
			}
			log.Println("chan2 end:", ok1)
		}
	}
}
