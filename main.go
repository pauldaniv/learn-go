package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pauldaniv/learn-go/service"
	"log"
	"net/http"
	"os"
)

// main function initializes and starts the server
func main() {

	// Create database connection
	connPool, err := pgxpool.NewWithConfig(context.Background(), DbConfig())
	if err != nil {
		log.Fatal("Error while creating connection to the database!!")
	}

	connection, err := connPool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error while acquiring connection from the database pool!!")
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		log.Fatal("Could not ping database")
	}

	fmt.Println("Connected to the database!!")

	defer connPool.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/bonds", listBonds(connPool))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on http://localhost:%d\n", port)
	if err := http.ListenAndServe(":"+port, corsMiddleware(mux)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
