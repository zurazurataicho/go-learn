package main

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point) ScaleBy(scale int) {
	p.X *= scale
	p.Y *= scale
}

func main() {
	pv := Point{1, 2}
	scaleBy := pv.ScaleBy
	scaleBy(10)
	fmt.Println(pv)

	pe := (*Point).ScaleBy
	p := Point{3, 4}
	pe(&p, 2)
	fmt.Println(p)
}
