package main

import "fmt"

type Hoge string

func (h *Hoge) Write(p []byte) (n int, err error) {
	*h = Hoge(p)
	return len(p), nil
}

type Fuga struct {
	message string
}

func (f *Fuga) Write(p []byte) (n int, err error) {
	(*f).message = string(p)
	return len(p), nil
}

func (f Fuga) String() string {
	return f.message
}

func example2() {
	var h Hoge
	fmt.Fprintf(&h, "message = %s", "hoge fuga")
	fmt.Println(h)

	var f Fuga
	fmt.Fprintf(&f, "message = %s", "fuga piyo")
	fmt.Println(f)
}
