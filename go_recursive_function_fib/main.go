package main

import "fmt"

func main() {
	var fib func(int) int
	fib = func(n int) int {
		if n == 0 || n == 1 {
			return n
		}
		return fib(n-2) + fib(n-1)
	}

	for i := 0; i < 20; i++ {
		fmt.Println(fib(i))
	}
}
