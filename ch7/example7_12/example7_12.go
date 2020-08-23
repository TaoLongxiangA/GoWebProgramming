package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
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

	jsonFile, err := json.MarshalIndent(post, "", "\t")
	if err != nil {
		fmt.Println("Error marshal struct to json:", err)
		return
	}

	fmt.Println("root json file")
	fmt.Println("type of jsonFile:", reflect.TypeOf(jsonFile))
	fmt.Println(string(jsonFile))

	err = ioutil.WriteFile("output.json", jsonFile, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

}
