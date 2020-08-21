package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"` //id属性
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"` //原始XML数据
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"` //xml元素的字符串数据
}

func main() {
	xmlFile, err := os.Open("ch7/example7_2/post.xml")
	if err != nil {
		fmt.Println("Error oping XML file:", err)
		return
	}

	date, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML file:", err)
		return
	}

	var post Post
	xml.Unmarshal(date, &post)
	fmt.Println(post)
}
