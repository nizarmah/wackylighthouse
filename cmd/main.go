package main

import (
	"log"
	"net/http"
)

// HTTP server for static files so we can run Lighthouse tests locally.
func main() {
	// Create the file server.
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Start the server.
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
