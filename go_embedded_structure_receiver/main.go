package main

import "fmt"

type Turn int

const (
	Left Turn = iota
	Right
)

type Engine struct {
	Type string
	Rpm  int
}

type Wheel struct {
	Degree int
}

type Car struct {
	Maker string
	Wheel
	Engine
}

func (c *Car) Turn(d Turn) {
	if d == Right {
		c.Degree += 90
		return
	}
	if d == Left {
		c.Degree -= 90
		return
	}
}

func main() {
	c := Car{"Honda", Wheel{0}, Engine{"B20B", 2500}}

	fmt.Println()
}
