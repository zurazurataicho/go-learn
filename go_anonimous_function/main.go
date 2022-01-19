package main

import "fmt"

func genAnon() func() int {
	var i int
	fmt.Printf("init i = %d\n", i)
	return func() int {
		i += 1
		return i
	}
}

func main() {
	fn := genAnon()
	fmt.Printf("fn() = %d\n", fn())
	fmt.Printf("fn() = %d\n", fn())
	fmt.Printf("fn() = %d\n", fn())
	fmt.Printf("fn() = %d\n", fn())
	fmt.Printf("fn() = %d\n", fn())
}
