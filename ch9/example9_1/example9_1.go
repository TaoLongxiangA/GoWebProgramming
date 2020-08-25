package main

import (
	"fmt"
	"time"
)

func printNum1() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func printLetters1() {
	for i := 'A'; i <= 10+'A'; i++ {
		fmt.Printf("%c ", i)
	}
}

func print1() {
	printNum1()
	printLetters1()
}

func goPrint1() {
	go printNum1()
	go printLetters1()
}

func printNum2() {
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
}

func printLetters2() {
	for i := 'A'; i <= 10+'A'; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
}

func print2() {
	printNum2()
	printLetters2()
}

func goPrint2() {
	go printNum2()
	go printLetters2()
}
func main() {

}
