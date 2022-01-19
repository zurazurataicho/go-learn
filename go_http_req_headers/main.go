package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

func acceptEncodingHeader(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Accept-Encoding"]
	fmt.Fprintln(w, h)
}

func acceptHeader(w http.ResponseWriter, r *http.Request) {
	h := r.Header.Get("Accept")
	fmt.Fprintln(w, h)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/accept-encoding", acceptEncodingHeader)
	http.HandleFunc("/accept", acceptHeader)
	server.ListenAndServe()
}
