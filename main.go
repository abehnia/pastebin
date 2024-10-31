package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Bin represents a text bin with text content.
type Bin struct {
	Text string `json:"text"`
}

// In-memory storage for bins.
var bins = make(map[string]Bin)
var binsMutex sync.Mutex

func main() {
	// Initialize the router.
	r := mux.NewRouter()

	// Define the routes.
	r.HandleFunc("/bins", createBinHandler).Methods("POST")
	r.HandleFunc("/bins/{uuid}", getBinHandler).Methods("GET")

	// Start the HTTP server.
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// createBinHandler handles the creation of a new bin.
func createBinHandler(w http.ResponseWriter, r *http.Request) {
	var newBin Bin

	// Parse the request body.
	err := json.NewDecoder(r.Body).Decode(&newBin)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate a new UUID for the bin.
	id := uuid.New().String()

	// Store the bin in the map with a lock.
	binsMutex.Lock()
	bins[id] = newBin
	binsMutex.Unlock()

	// Respond with the UUID.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

// getBinHandler retrieves a bin by UUID.
func getBinHandler(w http.ResponseWriter, r *http.Request) {
	// Get the UUID from the URL path.
	vars := mux.Vars(r)
	id := vars["uuid"]

	// Retrieve the bin from the map with a lock.
	binsMutex.Lock()
	bin, exists := bins[id]
	binsMutex.Unlock()

	if !exists {
		http.Error(w, "Bin not found", http.StatusNotFound)
		return
	}

	// Respond with the bin's text.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bin)
}
