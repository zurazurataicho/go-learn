package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type userName struct {
	First string `json:"first_name"`
	Last  string `json:"last_name"`
}

func isApplicationJson(r *http.Request) bool {
	if r.Method != "POST" {
		return false
	}

	if r.Header.Get("Content-Type") != "application/json" {
		return false
	}

	return true
}

func parseJSON(w http.ResponseWriter, r *http.Request) {
	if !isApplicationJson(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var name userName
	json.NewDecoder(r.Body).Decode(&name)

	fmt.Fprintln(w, name)
	fmt.Fprintf(w, "First: %s, Last: %s", name.First, name.Last)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/json", parseJSON)
	server.ListenAndServe()
}
