package main

import (
	"fmt"
)

type EmbedCar struct {
	Maker  string
	Wheel  Wheel
	Engine Engine
}

func embed() {
	var c EmbedCar
	c.Maker = "Toyota"
	c.Engine.Type = "M15A-FKS"
	c.Engine.Rpm = 2000
	c.Wheel.Degree = 20
	fmt.Println("> embedded structure: c.Maker = ...;")
	fmt.Printf("c = %v\n\n", c)
}
