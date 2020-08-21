package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	ID      int
	Content string
	Author  string
}

func store(date interface{}, filename string) {
	//creat buffer
	buffer := new(bytes.Buffer)
	//encoder
	encoder := gob.NewEncoder(buffer)
	//encode
	err := encoder.Encode(date)
	if err != nil {
		panic(err)
	}
	//write file
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func load(date interface{}, filename string) {
	//read content
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	decoder := gob.NewDecoder(buffer)
	//write
	err = decoder.Decode(date)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{
		ID:      1,
		Content: "HELLO,WORLD!",
		Author:  "John",
	}
	store(post, "file")
	var post1 Post
	load(&post1, "file")
	fmt.Println(post1)
}
