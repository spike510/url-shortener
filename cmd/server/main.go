package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spike510/url-shortener/internal/generator"
	internalhttp "github.com/spike510/url-shortener/internal/http"
	"github.com/spike510/url-shortener/internal/storage"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = fmt.Sprintf("http://localhost:%s", port)
	}

	h := internalhttp.NewHandler(baseURL, generator.NewCodeGenerator(), storage.NewInMemoryStorage())

	r := internalhttp.NewRouter()
	r.POST("/api/shorten", h.Shorten)
	r.GET("/:code", h.Redirect)

	log.Printf("Server running on :%s", port)

	go func() {
		if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}
