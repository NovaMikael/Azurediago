package server

import (
	"log"
	"net/http"
)

// StartServer initializes and starts the HTTP server
func StartServer() {
	http.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {})

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
