package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func responseHtml(w http.ResponseWriter, r *http.Request) {
	str := `
<html>
<head>
<title>Go Web Programming</title>
</head>
<body>
<h1>Hello, World</h1>
</body>
</html>`
	w.Write([]byte(str))
	fmt.Println("responseHtml")
}

func responseNoService(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No service.")
	fmt.Println("responseNoService")
}

func responseRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com/")
	w.WriteHeader(302)
	fmt.Println("responseRedirect")
}

func responseJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Atsushi Ezura",
		Threads: []string{"First", "Second", "Third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
	fmt.Println("responseJson")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/html", responseHtml)
	http.HandleFunc("/noservice", responseNoService)
	http.HandleFunc("/redirect", responseRedirect)
	http.HandleFunc("/json", responseJson)

	server.ListenAndServe()
}
