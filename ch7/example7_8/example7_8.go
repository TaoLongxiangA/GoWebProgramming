package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"Post"`
	Id      string   `xml:"id,attr"` //xml属性
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		//XMLName: "post1",
		Id:      "1",
		Content: "Hello,World!",
		Author: Author{
			Id:   "1",
			Name: "Jack",
		},
	}

	xmlFile, err := os.Create("post.xml")
	if err != nil {
		fmt.Println("Error creating XML file:", err)
	}

	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(post)
	if err != nil {
		fmt.Println("Error encoding XML to file:", err)
	}
}
