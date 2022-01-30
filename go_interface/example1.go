package main

import "fmt"

type Pet interface {
	Cry()
}

type Cat struct {
}

type Dog struct {
}

func (c Cat) Cry() {
	fmt.Println("Meow")
}

func (d Dog) Cry() {
	fmt.Println("Vow")
}

func MewOrVow(p Pet) {
	p.Cry()
}

func example1() {
	c := Cat{}
	d := Dog{}
	MewOrVow(c)
	MewOrVow(d)
}
