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

	ec := EmbedCar{"Honda", Wheel{0}, Engine{"B20B", 2500}}

	fmt.Println("> embedded structure: c.Maker = ...; ec := EmbedCar{...}")
	fmt.Printf("c = %v\n", c)
	fmt.Printf("ec = %v\n\n", ec)
}
