package main

import (
	"log"

	"github.com/spike510/url-shortener/internal/http"
)

func main() {
	r := http.NewRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	log.Println("Server running on :8080")
}
