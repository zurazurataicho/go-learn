package main

import (
	"fmt"
)

func main() {
	i := [...]int{1, 2, 3}
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", i[0:3])
	fmt.Printf("%T\n", i[1:2])
	fmt.Printf("%T\n", i[1:3])
	fmt.Printf("%T\n", i[1:])
	fmt.Printf("%T\n", i[:2])
}
