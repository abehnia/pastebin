package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"database/sql"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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

	psqlInfo := fmt.Sprintf("user=ubuntu password=pwd dbname=postgres sslmode=disable")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Parse the request body.
	err = json.NewDecoder(r.Body).Decode(&newBin)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate a new UUID for the bin.
	id := uuid.New().String()

	_, err = db.Exec("INSERT INTO pastebin(id, timestamp, title, content) VALUES($1, $2, $3, $4)", id, time.Now().UTC().Unix(), "my-title", newBin.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
