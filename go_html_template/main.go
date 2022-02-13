package main

import (
	"html/template"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	w := os.Stdout
	t, _ := template.ParseFiles(
		"template1.html",
		"template2.html",
		"template3.html",
		"template4.html",
	)

	// replace to first template
	t.Execute(w, "hoge")

	// replace to second template
	t.ExecuteTemplate(w, "template2.html", "fuga")

	// execute condition action on third template
	t.ExecuteTemplate(w, "template3.html", rand.Intn(10)%2 == 0)

	// execute iterator action on forth template
	list1 := []string{"Foo", "Bar", "Baz"}
	list2 := []string{}
	t.ExecuteTemplate(w, "template4.html", list1)
	t.ExecuteTemplate(w, "template4.html", list2)
}
