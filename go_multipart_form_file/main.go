package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./index.html")
	if err == nil {
		w.Write(data)
	}
}

func multipartFormFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fileHandler := r.MultipartForm.File["uploaded"][0]
	file, err := fileHandler.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", home)
	http.HandleFunc("/mff", multipartFormFile)
	server.ListenAndServe()
}
