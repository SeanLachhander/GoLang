package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// CD Struct (Model)

type CD struct {
	ID     string  `json:"id"`
	UPC    string  `json:"upc"`
	Title  string  `json:"title"`
	Artist *Artist `json:"artist"`
}

// Init CD variable as a slice CD struct
var CDs []CD

// Artist Struct

type Artist struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Get all CDs

func getCDs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CDs)
}

// Get a CD

func getCD(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through the music library, and find the one with matching ID
	for _, item := range CDs {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&CD{})
}

// Create a new CD

func createCD(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cd CD
	_ = json.NewDecoder(r.Body).Decode(&cd)
	cd.ID = strconv.Itoa(rand.Intn(10000000))
	CDs = append(CDs, cd)
	json.NewEncoder(w).Encode(cd)
}

// Update a CD

func updateCD(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range CDs {
		if item.ID == params["id"] {
			CDs = append(CDs[:index], CDs[index+1:]...)
			var cd CD
			_ = json.NewDecoder(r.Body).Decode(&cd)
			cd.ID = params["id"]
			CDs = append(CDs, cd)
			json.NewEncoder(w).Encode(cd)
			return
		}
	}
	json.NewEncoder(w).Encode(CDs)
}

// Delete a CD.

func deleteCD(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range CDs {
		if item.ID == params["id"] {
			CDs = append(CDs[:index], CDs[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(CDs)
}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Mock Data
	CDs = append(CDs, CD{ID: "1", UPC: "12345", Title: "The King of Dancehall", Artist: &Artist{Firstname: "Adidja", Lastname: "Palmer"}})
	CDs = append(CDs, CD{ID: "2", UPC: "67890", Title: "Greatest Hits", Artist: &Artist{Firstname: "Tupac", Lastname: "Shakur"}})

	// Route Handlers -- Endpoints
	router.HandleFunc("/api/CDs", getCDs).Methods("GET")
	router.HandleFunc("/api/CDs/{id}", getCD).Methods("GET")
	router.HandleFunc("/api/CDs", createCD).Methods("POST")
	router.HandleFunc("/api/CDs/{id}", updateCD).Methods("PUT")
	router.HandleFunc("/api/CDs/{id}", deleteCD).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":1234", router))

}
