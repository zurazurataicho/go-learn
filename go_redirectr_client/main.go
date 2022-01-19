package main

import (
	"io/ioutil"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./index.html")
	if err == nil {
		w.Write(data)
	}
}

func redirectClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://zura.org/")
	w.WriteHeader(301)
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/", home)
	http.HandleFunc("/redirect", redirectClient)
	server.ListenAndServe()
}
