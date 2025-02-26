package persistence

import (
	"context"
	"finance/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

// itemsHandler handles the "/items" endpoint
func FindAll(dbPool *pgxpool.Pool) ([]model.Bond, error) {

	// Query the database for bonds
	rows, err := dbPool.Query(context.Background(), "SELECT id, name, count, buy_price, sell_price, currency_type, start_date, end_date, created_at FROM bonds")
	if err != nil {
		log.Printf("Database query error: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	// Parse the results into a list of bonds
	bonds := []model.Bond{}
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

		bonds = append(bonds, bond)
		if err != nil {
			log.Printf("Row scan error: %v\n", err)
			return nil, err
		}

	}

	// Check for any errors encountered during iteration
	if rows.Err() != nil {
		log.Printf("Row iteration error: %v\n", rows.Err())
		return nil, rows.Err()
	}
	return bonds, nil

}

// AddBond handles POST requests to add a new bond
func StoreBond(bond model.Bond, dbPool *pgxpool.Pool) (model.Bond, error) {

	// Insert the new bond into the database and return the inserted record
	query := `INSERT INTO bonds (id, name, isin, count, buy_price, sell_price, currency_type, start_date, end_date, created_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, now())
			  RETURNING id, name, isin, count, buy_price, sell_price, currency_type, start_date, end_date, created_at`

	var insertedBond model.Bond
	err := dbPool.QueryRow(
		context.Background(),
		query,
		bond.ID,
		bond.Name,
		bond.IsIn,
		bond.Count,
		bond.BuyPrice,
		bond.SellPrice,
		bond.CurrencyType,
		bond.StartDate,
		bond.EndDate,
	).Scan(
		&insertedBond.ID,
		&insertedBond.Name,
		&insertedBond.IsIn,
		&insertedBond.Count,
		&insertedBond.BuyPrice,
		&insertedBond.SellPrice,
		&insertedBond.CurrencyType,
		&insertedBond.StartDate,
		&insertedBond.EndDate,
		&insertedBond.CreatedAt,
	)
	if err != nil {
		log.Printf("Database insert error: %v\n", err)
		return model.Bond{}, err
	}
	return insertedBond, nil
}
