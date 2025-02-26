package main

import (
	"context"
	"finance/config"
	"finance/service"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"os"
	"strconv"
)

func setupHandlers(mux *http.ServeMux, connPool *pgxpool.Pool) {
	mux.HandleFunc("/v1/bonds", service.HandleBonds(connPool))
}

// main function initializes and starts the server
func main() {

	// Create database connection
	connPool, err := pgxpool.NewWithConfig(context.Background(), config.DbConfig())
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
	} else {
		fmt.Println("Connected to the database!!")
	}
	defer connPool.Close()

	mux := http.NewServeMux()
	setupHandlers(mux, connPool)
	port_str := os.Getenv("PORT")
	if port_str == "" {
		port_str = "8080"
	}
	// string to int
	port, err := strconv.Atoi(port_str)
	if err != nil {
		// ... handle error
		panic(err)
	}
	fmt.Printf("Server running on http://localhost:%d\n", port)
	if err := http.ListenAndServe(":"+port_str, config.CorsMiddleware(mux)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
