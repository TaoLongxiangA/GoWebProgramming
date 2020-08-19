package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	date := []byte("Hello,World!")
	err := ioutil.WriteFile("file1", date, 0644)
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("file1")
	fmt.Println(string(read1))

	file1, _ := os.Create("date2")
	defer file1.Close()

	count, _ := file1.Write(date)
	fmt.Printf("write %d bytes\n", count)

	file2, _ := os.Open("date2")
	defer file2.Close()

	read2 := make([]byte, len(date))
	count1, _ := file2.Read(read2)
	fmt.Printf("write %d bytes\n", count1)
	fmt.Println(string(read2))
}
