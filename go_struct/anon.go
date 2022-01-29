package main

import (
	"fmt"
)

type AnonCar struct {
	Maker string
	Wheel
	Engine
}

func anon() {
	var c AnonCar
	c.Maker = "Toyota"
	c.Type = "M15A-FKS"
	c.Rpm = 2000
	c.Degree = 20
	fmt.Println("> anonymous structure: c.Maker = ...;")
	fmt.Printf("c = %v\n\n", c)
}
