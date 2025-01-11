package service

import (
	"encoding/json"
	"finance/model"
	"finance/persistence"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
)

// itemsHandler handles the "/items" endpoint
func ListBonds(w http.ResponseWriter, r *http.Request, dbPool *pgxpool.Pool) {

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
	bonds, err := persistence.FindAll(dbPool)

	if err != nil {
		http.Error(w, "Failed to parse database result", http.StatusInternalServerError)
		log.Printf("Row scan error: %v\n", err)
		return
	}
	// Send the list of bonds as JSON response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bonds); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("JSON encoding error: %v\n", err)
	}
}

// AddBond handles POST requests to add a new bond
func AddBond(w http.ResponseWriter, r *http.Request, dbPool *pgxpool.Pool) {

	// Parse the incoming JSON request body
	var bond model.Bond
	if err := json.NewDecoder(r.Body).Decode(&bond); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Failed to decode request body: %v\n", err)
		return
	}

	// Validate required fields
	if bond.Name == "" || bond.Count <= 0 || bond.BuyPrice <= 0 || bond.SellPrice <= 0 || bond.CurrencyType == "" || bond.EndDate.IsZero() {
		http.Error(w, "Missing or invalid fields", http.StatusBadRequest)
		log.Println("Validation failed for input fields")
		return
	}

	record, err := persistence.StoreBond(bond, dbPool)

	if err != nil {
		http.Error(w, "Failed to save bond", http.StatusInternalServerError)
		log.Printf("Database insert error: %v\n", err)
		return
	}

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(record)
}

// HandleBonds handles both GET and POST requests for bonds
func HandleBonds(dbPool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ListBonds(w, r, dbPool)
		case http.MethodPost:
			AddBond(w, r, dbPool)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}
