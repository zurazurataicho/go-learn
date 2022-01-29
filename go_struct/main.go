package main

type Engine struct {
	Type string
	Rpm  int
}

type Wheel struct {
	Degree int
}

func main() {
	point()
	embed()
	anon()
}
