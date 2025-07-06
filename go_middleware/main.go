package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	middleware "zura.org/middleware/internal/middleware"
	handler "zura.org/middleware/internal/handler"
)

func run(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/hello", middleware.ChainMiddleware(http.HandlerFunc(handler.HelloHandler), middleware.AuthMiddleware, middleware.LogMiddleware))
	mux.Handle("/goodbye", middleware.ChainMiddleware(http.HandlerFunc(handler.GoodByeHandler), middleware.AuthMiddleware, middleware.LogMiddleware))

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

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("Error: %v\n", err)
	}
}
