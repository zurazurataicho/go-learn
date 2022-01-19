package main

import (
	"fmt"
	"net/http"
)

func parse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.FormValue("first_name"))
	fmt.Fprintln(w, r.PostFormValue("first_name"))
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.PostForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/parse", parse)
	server.ListenAndServe()
}
