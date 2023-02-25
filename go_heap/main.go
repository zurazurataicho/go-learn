package main

import (
	"fmt"
	"zura.org/go_heap/heap"
)

func main() {
	heap_test()
}

func heap_test() {
	h := heap.NewHeap(10)
	data := []int{10, 5, 20, 3}
	for _, v := range data {
		fmt.Println(v)
		if err := h.Add(v); err != nil {
			fmt.Println(err)
			return
		}
	}
	h.Show()
}
