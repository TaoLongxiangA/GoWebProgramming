package main

import (
	"fmt"
	"time"
)

func callerA(c chan string) {
	c <- "Hello,world!"
}

func callerB(c chan string) {
	c <- "Hala,Mundo!"
}

func main() {
	chan1, chan2 := make(chan string), make(chan string)
	go callerA(chan1)
	go callerB(chan2)
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Millisecond)
		select {
		case msg := <-chan1:
			fmt.Printf("%s from A\n", msg)
		case msg := <-chan2:
			fmt.Printf("%s from B\n", msg)
		default:
			fmt.Println("DEFAULT")
		}
	}
}
