package main

import "fmt"

// Interface
type IMyPrint interface {
	Print()
}

// MyPrint implements IMyPrint
type MyPrint struct {
	Message string
}

func (p *MyPrint) Print() {
	fmt.Println("MyPrint: ", p.Message)
}

// MyPrintf implements IMyPrint
type MyPrintf struct {
	Message string
}

func (p *MyPrintf) Print() {
	fmt.Printf("MyPrintf: %s\n", p.Message)
}

func main() {
	var imp IMyPrint

	imp = &MyPrint{Message: "Hello, MyPrint!"}
	imp.Print()

	imp = &MyPrintf{Message: "Hello, MyPrintf!!!"}
	imp.Print()
}
