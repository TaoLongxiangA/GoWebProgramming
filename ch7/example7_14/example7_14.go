package main

import "database/sql"

var Db *sql.DB

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

func init() {
	var err error
	Db, err = sql.Open("mysql", "root@123456")
	if err != nil {
		panic(err)
	}
}

func retrieve(id int) (post Post, err error) {
	err = Db.QueryRow("select * from Post where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}
