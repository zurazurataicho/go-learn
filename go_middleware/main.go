package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"zura.org/middleware/middleware"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(string)
	fmt.Fprintln(w, "Hello, %s!", user)
}

func goodByeHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(string)
	fmt.Fprintln(w, "Good bye, %s!", user)
}

func run(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/hello", middleware.ChainMiddleware(http.HandlerFunc(helloHandler), middleware.AuthMiddleware, middleware.LogMiddleware))

	s := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		log.Println("Starting server on :8080")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("server failed: %w", err)
		}
		return nil
	})

	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("Server shutdown failed: %v\n", err)
	}
	log.Println("Server gracefully stopped")

	return eg.Wait()
}
