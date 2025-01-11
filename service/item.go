package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

// itemsHandler handles the "/items" endpoint
func itemsHandler(dbPool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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

		// Query the database for items
		rows, err := dbPool.Query(context.Background(), "SELECT id, name, price FROM items")
		if err != nil {
			http.Error(w, "Failed to fetch items from database", http.StatusInternalServerError)
			log.Printf("Database query error: %v\n", err)
			return
		}
		defer rows.Close()

		// Parse the results into a list of items
		var items []Item
		for rows.Next() {
			var item Item
			err := rows.Scan(&item.ID, &item.Name, &item.Price)
			if err != nil {
				http.Error(w, "Failed to parse database result", http.StatusInternalServerError)
				log.Printf("Row scan error: %v\n", err)
				return
			}
			items = append(items, item)
		}

		// Send the list of items as JSON response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(items); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			log.Printf("JSON encoding error: %v\n", err)
		}
	}
}
