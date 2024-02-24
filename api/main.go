package main

import (
	"net/http"

	"github.com/yourusername/yourpackage/handler"
)

func main() {
	// Create a new mux router
	r := http.NewServeMux()

	// Set up the handler for the root path
	r.HandleFunc("/", handler.Handler)

	// Start the server
	http.ListenAndServe(":8080", r)
}
