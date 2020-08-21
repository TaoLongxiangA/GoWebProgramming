package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Post struct {
	XMLName xml.Name  `xml:"post"`
	Id      string    `xml:"id"`
	Content string    `xml:"content"`
	Author  Author    `xml:"author"`
	Comment []Comment `xml:"comments>comment"`
}

type Author struct {
	Id      string `xml:"id"`
	Content string `xml:"content"`
}

type Comment struct {
	Id      string `xml:"id"`
	Content string `xml:"content"`
	Author  string `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("ch7/example7_6/post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokes:", err)
			return
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se)
				fmt.Println(comment)
			}
		}
	}
}
