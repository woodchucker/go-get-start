package main

import (
	"log"
	"net/http"

	handler "github.com/woodchucker/go-get-start/api"
)

func main() {
	http.HandleFunc("/", handler.Handler) // Match the Vercel endpoint
	// Start local server
	println("Server running on http://localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
