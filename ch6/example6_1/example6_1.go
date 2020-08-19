package main

import (
	"fmt"
)

var (
	PostByid     map[int]*Post
	PostByAuthor map[string][]*Post
)

type Post struct {
	ID      int
	Content string
	Author  string
}

func store(post Post) {
	PostByid[post.ID] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], &post)
}

func main() {
	PostByid = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)

	post1 := Post{
		ID:      1,
		Content: "Hello World!",
		Author:  "Sau Sheong",
	}
	post2 := Post{
		ID:      2,
		Content: "Bonjour Monda!",
		Author:  "Pierre",
	}
	post3 := Post{
		ID:      3,
		Content: "Halo Mundo",
		Author:  "pedro",
	}
	post4 := Post{
		ID:      4,
		Content: "Greetings Earthlings",
		Author:  "Sau Sheong",
	}
	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostByid[1])
	fmt.Println(PostByid[2])
	for _, post := range PostByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}
	for _, post := range PostByAuthor["pedro"] {
		fmt.Println(post)
	}
}
