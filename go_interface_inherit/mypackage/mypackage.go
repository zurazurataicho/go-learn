package mypackage

import "fmt"

// IMyPrint interface
type IMyPrint interface {
	Print(prefix string)
}

// MyPrint implements IMyPrint
type MyPrint struct {
	Message string
}

// NewMyPrint makes its own instance
func NewMyPrint(message string) *MyPrint {
	return &MyPrint{Message: message}
}

// Print belongs to MyPrint
func (p *MyPrint) Print(prefix string) {
	fmt.Println(prefix, p.Message)
}

// IMyPrintExtra interface
type IMyPrintExtra interface {
	IMyPrint
	PrintExtra(prefix string)
}

// MyPrintExtra implements IMyPrintExtra
type MyPrintExtra struct {
	Message string
	Deco    string
}

// NewMyPrintExtra makes its own instance
func NewMyPrintExtra(message string, deco string) *MyPrintExtra {
	return &MyPrintExtra{Message: message, Deco: deco}
}

// Print belongs to MyPrintExtra
func (p *MyPrintExtra) Print(prefix string) {
	fmt.Printf("%s%s\n", prefix, p.Message)
}

// PrintExtra belongs to MyPrintExtra
func (p *MyPrintExtra) PrintExtra(prefix string) {
	fmt.Printf("%s%s%s\n", prefix, p.Deco, p.Message)
}
