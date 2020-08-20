package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Post struct {
	ID     int
	Conten string
	Author string
}

func main() {
	csvFile, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	posts := []Post{
		Post{
			ID:     1,
			Conten: "HELLO,WORLD!",
			Author: "Sau Sheong",
		},
		Post{
			ID:     2,
			Conten: "Bonjour Monde!",
			Author: "Pierre",
		},
		Post{
			ID:     3,
			Conten: "Hala Mundo!",
			Author: "Pedor",
		},
		Post{
			ID:     4,
			Conten: "Greetings Earthlings",
			Author: "Sau Sheong",
		},
	}
	writer := csv.NewWriter(csvFile)

	for _, post := range posts {
		line := []string{strconv.Itoa(post.ID), post.Conten, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush() //!!!notice
}
