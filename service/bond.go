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
func ListBonds(dbPool *pgxpool.Pool) http.HandlerFunc {
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
}
