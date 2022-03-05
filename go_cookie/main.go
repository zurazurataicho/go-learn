package main

import (
	_ "encoding/base64"
	"fmt"
	"net/http"
	"time"
)

type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string
	MaxAge     int
	Secure     bool
	HttpOnly   bool
	Raw        string
	Unparsed   []string
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	// c1 := http.Cookie{
	// 	Name:     "first_cookie",
	// 	Value:    "Go Rust Haskell",
	// 	HttpOnly: true,
	// }
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Foo Bar Baz",
		HttpOnly: true,
	}
	// http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
	fmt.Println("setCookie")
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get first_cookie")
	}

	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
	fmt.Println("getCookie")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
