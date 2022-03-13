package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

// type Post struct {
// 	Id      int
// 	Content string
// 	Author  string
// }

func gob_store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func gob_load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func gobfile() {
	storePost := Post{Id: 2, Content: "Hello, world second!", Author: "Piyo Piyo"}
	gob_store(&storePost, "gob_post")

	var loadPost Post
	gob_load(&loadPost, "gob_post")

	fmt.Println("=== gob ===")
	fmt.Println(loadPost)
}
