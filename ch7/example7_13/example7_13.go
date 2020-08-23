package main

import (
	"encoding/json"
	"fmt"
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
	post := Post{
		Id:      1,
		Content: "Hello,World!",
		Author: Author{
			Id:   2,
			Name: "Jack",
		},
		Comments: []Comment{
			{
				Id:      3,
				Content: "Have a good day!",
				Author:  "Adam",
			},
			{
				Id:      4,
				Content: "How are you!",
				Author:  "John",
			},
		},
	}

	jsonFile, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating json file!:", err)
		return
	}

	encoder := json.NewEncoder(jsonFile)

	//format
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error writing json file!:", err)
	}
}
