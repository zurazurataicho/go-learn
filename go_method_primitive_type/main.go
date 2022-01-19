package main

import (
	"fmt"
	"strconv"
)

type Integer int

func (i Integer) ToString() string {
	return strconv.Itoa(int(i))
}

func main() {
	var i Integer
	i = 10
	fmt.Printf("i = %v\n", i)
	fmt.Printf("i.ToString() = %v\n", i.ToString())
}
