package main

import (
	"fmt"
	"sync"
)

//通过wait group解决goroutine运行完，未打印的问题
func thrower(c chan int, wg *sync.WaitGroup) {
	//log.Printf("%v ", c)
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("thrower>>", i)
	}
	wg.Done()
}

func catcher(c chan int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("catcher<<", num)
	}
	wg.Done()
}

//通过time.sheep解决打印问题
//func thrower(c chan int) {
//	for i := 0; i < 5; i++ {
//		c <- i
//		 fmt.Println("thrower>>", i)
//	}
//}
//
//func catcher(c chan int) {
//	for i := 0; i < 5; i++ {
//		num := <-c
//		fmt.Println("catcher<<", num)
//	}
//}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan int)
	go thrower(c, &wg)
	go catcher(c, &wg)
	wg.Wait()

	/*=====================================*/
	//go thrower(c)
	//go catcher(c)
	//time.Sleep(100 * time.Millisecond)
}
