package main

import (
	"fmt"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post *Post) {
	PostById[post.Id] = post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], post)
}

func mem() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello", Author: "Hoge"}
	post2 := Post{Id: 2, Content: "World", Author: "Fuga"}
	post3 := Post{Id: 3, Content: "This is", Author: "Piyo"}
	post4 := Post{Id: 4, Content: "a pen", Author: "Pugya"}

	store(&post1)
	store(&post2)
	store(&post3)
	store(&post4)
	post3.Id = 100

	fmt.Println("=== mem ===")
	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Hoge"] {
		fmt.Println(post)
	}
	fmt.Println("------")
	for _, post := range PostsByAuthor["Piyo"] {
		fmt.Println(post)
	}
}
