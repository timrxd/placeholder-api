package main

import (
	"log"
	"net/http"

	"github.com/timrxd/placeholder-api/api"
	"github.com/timrxd/placeholder-api/data"
)

func main() {
	log.Println("Starting server...")

	// Initialize router
	server := api.CreateServer()

	// Load placeholder data
	data, err := data.ImportData("data/placeholder.json")
	if err != nil {
		log.Fatal(err)
	}
	api.LoadData(data)

	// Serve API on port 8080
	log.Fatal(http.ListenAndServe(":8080", server))
}
