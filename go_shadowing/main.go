package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)
	fmt := 100
	fmt.Println(fmt) // fmt.Println undefined (type int has no field or method Println)
}
