package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func (p *Point) Swap() {
	p.X, p.Y = p.Y, p.X
}

func point() {
	p := Point{1, 2}
	fmt.Println("> p := Point{1, 2}")
	fmt.Printf("p = %v\n\n", p)

	var pp *Point
	pp = &p
	pp.X++
	fmt.Println("> var pp *Point; pp = &p; pp.X++")
	fmt.Printf("p = %v\n", p)
	fmt.Printf("pp = %v\n\n", *pp)

	p.Y--
	fmt.Println("> p.Y--")
	fmt.Printf("p = %v\n", p)
	fmt.Printf("pp = %v\n\n", *pp)

	p1 := Point{10, 20}
	p2 := Point{20, 10}
	fmt.Println("> p1 := Point{10, 20}; p2 := Point{20, 10}; Compair to p1 and p2")
	fmt.Println(p1 == p2)

	fmt.Println("> p1.Swap; Compair to p1 and p2")
	p1.Swap()
	fmt.Printf("%v\n\n", p1 == p2)
}
