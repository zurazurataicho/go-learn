package handler

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(string)
	fmt.Fprintln(w, "Hello, %s!", user)
}

func GoodByeHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(string)
	fmt.Fprintln(w, "Good bye, %s!", user)
}
