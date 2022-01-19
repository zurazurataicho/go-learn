package main

import "fmt"
import "zura.org/greetings"

func main() {
	message := greetings.Hello("zura")
	fmt.Println(message)
}
