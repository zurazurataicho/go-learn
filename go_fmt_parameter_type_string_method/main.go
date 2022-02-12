package main

import "fmt"

type MyType struct {
	Message string
}

func (p *MyType) String() string {
	return p.Message
}

func main() {
	v := MyType{Message: "hoge"}
	fmt.Printf("%v\n", v)
}
