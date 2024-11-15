package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HelloResponse represents the JSON structure for the "Hello" response
type HelloResponse struct {
	Message string `json:"message"`
}

// helloHandler handles the "Hello" endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Only allow GET method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	// Create and send the response
	response := HelloResponse{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

// main function initializes and starts the server
func main() {
	http.HandleFunc("/hello", helloHandler)

	port := 8080
	fmt.Printf("Server running on http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
