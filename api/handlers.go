package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

// GetItems returns all items in database
func GetItems(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /items")

	// Output items as list
	list := []*Item{}
	for _, v := range db {
		list = append(list, v)
	}

	// Sort by ID
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})

	json.NewEncoder(w).Encode(list)
}

// GetItem returns a single item based on ID
func GetItem(w http.ResponseWriter, r *http.Request) {
	// Get ID from params
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Invalid ID")
		return
	}
	log.Printf("GET /item/%v\n", id)

	// Get item from database
	item := db[id]
	if item == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("ID not found")
		return
	}

	// Return item
	json.NewEncoder(w).Encode(item)
}

// CreateItem creates a new item in the database
func CreateItem(w http.ResponseWriter, r *http.Request) {
	log.Println("POST /item")

	// Read in item from request
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid item")
		return
	}

	// Check if id is free
	if db[newItem.ID] != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID already in use")
		return
	}

	// Add item to database
	db[newItem.ID] = &newItem
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

// DeleteItem removes an item from the database
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	// Get ID from params
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Invalid ID")
		return
	}
	log.Printf("DELETE /item/%v\n", id)

	// Check if id exists
	if db[id] == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID does not exist")
		return
	}

	// Remove item from database
	delete(db, id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Item deleted")
}

// UpdateItem changes an item in the database
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	// Get ID from params
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Invalid ID")
		return
	}
	log.Printf("PUT /item/%v\n", id)

	// Check if id exists
	if db[id] == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID does not exist")
		return
	}

	// Read in item from request
	var newItem Item
	err = json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid item")
		return
	}

	// Ignore any id specified in body
	newItem.ID = id

	// Update item in database
	db[id] = &newItem
	json.NewEncoder(w).Encode(newItem)
}
