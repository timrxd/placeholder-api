package api

import (
	"github.com/gorilla/mux"
)

// Item is a placeholder struct for data
type Item struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// db is a placeholder for a database
var db map[int]*Item

// Create server initializes the router and database
func CreateServer() *mux.Router {

	// Initialize server
	router := mux.NewRouter()
	db = map[int]*Item{}

	// Assign handlers to routes
	router.HandleFunc("/items", GetItems).Methods("GET")
	router.HandleFunc("/item/{id}", GetItem).Methods("GET")
	router.HandleFunc("/item", CreateItem).Methods("POST")
	router.HandleFunc("/item/{id}", DeleteItem).Methods("DELETE")
	router.HandleFunc("/item/{id}", UpdateItem).Methods("PUT")

	return router
}

// LoadData directly initializes the db if necessary
func LoadData(data map[int]*Item) {
	db = data
}
