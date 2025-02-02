package main

import (
	"log"
	"net/http"

	"moneytool/api"
	"moneytool/storage"
)

func main() {
	store, err := storage.NewPostgresStorage("host=postgres user=moneytool password=moneytool dbname=moneytool port=5432 sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	handler := api.NewHandler(store)

	http.HandleFunc("/api/v1/account", handler.CreateAccount)
	http.HandleFunc("/api/v1/accounts", handler.GetAccounts)
	http.HandleFunc("/api/v1/accounts/{id}", handler.GetAccount)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
