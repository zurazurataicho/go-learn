package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	contentLenght := r.ContentLength
	if contentLenght == 0 {
		fmt.Fprintln(w, "empty")
		return
	}
	requestBody := make([]byte, contentLenght)
	r.Body.Read(requestBody)
	fmt.Fprintln(w, string(requestBody))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}
