package main

import "fmt"

type Test struct {
}

func (t Test) Method(val int) {
	fmt.Printf("Method expression test: %d\n", val)
}

func main() {
	tm := Test.Method
	t1 := Test{}
	tm(t1, 10)
}
