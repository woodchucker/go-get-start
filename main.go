package main

import (
	"net/http"

	handler "github.com/woodchucker/go-get-start/api"
)

func main() {
	http.HandleFunc("/", handler.Handler) // Match the Vercel endpoint
	http.ListenAndServe(":3000", nil)
}
