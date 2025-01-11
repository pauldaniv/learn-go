package service

import (
	"context"
	"encoding/json"
	"finance/model"
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

	// Query the database for bonds
	rows, err := dbPool.Query(context.Background(), "SELECT id, name, count, buy_price, sell_price, currency_type, start_date, end_date, created_at FROM bonds")
	if err != nil {
		http.Error(w, "Failed to fetch bonds from database", http.StatusInternalServerError)
		log.Printf("Database query error: %v\n", err)
		return
	}
	defer rows.Close()

	// Parse the results into a list of bonds
	var bonds []model.Bond
	for rows.Next() {
		var bond model.Bond
		err := rows.Scan(
			&bond.ID,
			&bond.Name,
			&bond.Count,
			&bond.BuyPrice,
			&bond.SellPrice,
			&bond.CurrencyType,
			&bond.StartDate,
			&bond.EndDate,
			&bond.CreatedAt,
		)
		if err != nil {
			http.Error(w, "Failed to parse database result", http.StatusInternalServerError)
			log.Printf("Row scan error: %v\n", err)
			return
		}
		bonds = append(bonds, bond)
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

	// Only allow POST method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

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

	// Insert the new bond into the database
	query := `INSERT INTO bonds (id, name, count, buy_price, sell_price, currency_type, start_date, end_date, created_at) 
				  VALUES (gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7, now())`
	_, err := dbPool.Exec(
		context.Background(),
		query,
		bond.Name,
		bond.Count,
		bond.BuyPrice,
		bond.SellPrice,
		bond.CurrencyType,
		bond.StartDate,
		bond.EndDate,
	)
	if err != nil {
		http.Error(w, "Failed to save bond", http.StatusInternalServerError)
		log.Printf("Database insert error: %v\n", err)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Bond added successfully"})
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
