package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./home.html")
	if err == nil {
		w.Write(data)
	}
}

func writeHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No services")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", home)
	http.HandleFunc("/writeheader", writeHeader)
	server.ListenAndServe()
}
