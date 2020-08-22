package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("ch7/example7_10/post.json")
	if err != nil {
		fmt.Println("Error opening file :", err)
		return
	}
	defer jsonFile.Close()

	jsonDate, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading json date:", err)
		return
	}

	var post Post
	json.Unmarshal(jsonDate, &post)
	fmt.Println(post)
}
