package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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
	xmlFile, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshaling file:", err)
		return
	}
	err = ioutil.WriteFile("post.xml", []byte(xml.Header+string(xmlFile)), 0644)
	if err != nil {
		fmt.Println("Error write file :", err)
		return
	}
}
