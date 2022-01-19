package main

import (
	"fmt"
	"strconv"
)

func map1() {
	m := make(map[string]int)
	for i := 0; i < 100; i++ {
		m["map"+strconv.Itoa(i)] = i
	}
	fmt.Printf("len = %d\n", len(m))
	fmt.Printf("type = %T\n", m)
	fmt.Printf("map = %v\n", m)
}

func map2() {
	m1 := make(map[string]int)
	var m2 map[string]int
	fmt.Println(m2 == nil)
	fmt.Println(len(m2) == 0)
	m1["hoge"] = 1
	m1["fuga"] = 2
	m2 = m1
	fmt.Printf("m1 = %v\n", m1)
	fmt.Printf("m2 = %v\n", m2)

	m2["piyo"] = 3
	fmt.Printf("m1 = %v\n", m1)
	fmt.Printf("m2 = %v\n", m2)
}

func map3() {
	// nil map
	var m1 map[string]int
	fmt.Println(m1)
	if val, ok := m1["hoge"]; !ok {
		fmt.Printf("m1: No key, val = %d\n", val)
	}

	// empty map
	m2 := make(map[string]int)
	fmt.Println(m2)
	if val, ok := m2["hoge"]; !ok {
		fmt.Printf("m2: No key, val = %d\n", val)
	}

	m3 := make(map[string]int)
	m3["piyo"] = 3
	fmt.Println(m3)
	if val, ok := m3["hoge"]; !ok {
		fmt.Printf("m3: No key, val = %d\n", val)
	}
	val, ok := m3["piyo"]
	fmt.Printf("val=%d, ok=%v\n", val, ok)
}

func main() {
	map1()
	map2()
	map3()
}
