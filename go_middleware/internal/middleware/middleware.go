package middleware

import (
	"context"
	"log"
	"net/http"
)

type ContextKey string

func withUser(ctx context.Context, user ContextKey) context.Context {
	return context.WithValue(ctx, "user", user)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer valid-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		var user ContextKey = "AuthenticatedUser"
		ctx := withUser(r.Context(), user)
		log.Printf("Setting user in context: %s", user) // debug
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func ChainMiddleware(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	next := h
	for i := len(middlewares) - 1; i >= 0; i-- {
		next = middlewares[i](next)
	}
	return next
}
