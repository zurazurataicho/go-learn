package main

import "fmt"

type User struct {
	name string
	age  int
}

func (u User) Show() {
	fmt.Println(u)
}

func (u User) AgeValue() {
	u.age++
}

func (u *User) AgePointer() {
	u.age++
}

func main() {
	uv := User{"Taro", 20}
	up := &User{"Hanako", 30}

	fmt.Println("uv")
	uv.Show() // {Taro 20}
	uv.AgeValue()
	uv.Show() // {Taro 20}
	uv.AgePointer()
	uv.Show() // {Taro 21}

	fmt.Println("up")
	up.Show() // {Hanako 30}
	up.AgeValue()
	up.Show() // {Hanako 30}
	up.AgePointer()
	up.Show() // {Hanako 31}
}
