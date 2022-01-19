package main

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Show() {
	fmt.Println(u)
}

func (u *User) Birthday() {
	u.Age++
}

func main() {
	u := User{1, "Taro", 20}
	up := &User{2, "Hanako", 18}

	// Same receiver specify both receiver parameter and method argument
	u.Show()
	up.Birthday()

	// Receiver parameter is a variable, receiver argument is a pointer
	u.Birthday()

	// Receiver parameter is a point, receiver argument is a variable
	up.Show()

	// Literal can call a method that has a variable receiver argument
	User{3, "Kenta", 23}.Show()

	// Literal cannot call a method that has a pointer receiver argument
	// User{3, "Kenta", 23}.Brithday() // Compile Error
}
