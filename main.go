package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Item represents an individual item in the list
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

// addCORSHeaders adds CORS headers to the response
func addCORSHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// itemsHandler handles the "/items" endpoint
func itemsHandler(w http.ResponseWriter, r *http.Request) {
	addCORSHeaders(w, r)

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only allow GET method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	// Create a list of items
	items := []Item{
		{ID: 1, Name: "Laptop", Price: "$999"},
		{ID: 2, Name: "Smartphone", Price: "$799"},
		{ID: 3, Name: "Tablet", Price: "$499"},
	}

	// Send the list of items as JSON response
	json.NewEncoder(w).Encode(items)
}

// helloHandler handles the "/hello" endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	addCORSHeaders(w, r)

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only allow GET method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	// Create and send the response
	response := map[string]string{"message": "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

// main function initializes and starts the server
func main() {
	http.HandleFunc("/v1/items", itemsHandler)

	port := 8080
	fmt.Printf("Server running on http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
