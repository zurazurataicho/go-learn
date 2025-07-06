package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"encoding/json"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var (
	oauthConfig = &oauth2.Config{
		ClientID:     "1088542013789-mem896r2s0g3odj7acvv7vrfd69dotcb.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	oauthStateString = "pseudo-random"
	jwtSecret = []byte("your-jwt-secret")
)

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)
	http.HandleFunc("/protected", handleProtectedRoute)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL(oauthStateString)
	json.NewEncoder(w).Encode(map[string]string{"url": url})
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		http.Error(w, "State mismatch", http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v1/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	fmt.Fprint(w, "Login successful!<br>")
	resp.Write(w)
}
