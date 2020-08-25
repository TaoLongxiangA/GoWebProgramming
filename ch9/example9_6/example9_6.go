package main

import (
	"fmt"
	"sync"
	"time"
)

func printNum1(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func printLetters1(wg *sync.WaitGroup) {
	for i := 'A'; i < 10+'A'; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNum1(&wg)
	go printLetters1(&wg)
	wg.Wait()
}
