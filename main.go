package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"database/sql"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Bin represents a text bin with text content.
type BinRequest struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

type Bin struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	Title     string `json:"title"`
	Timestamp string `json:"timestamp"`
}

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
	var newBin BinRequest

	psqlInfo := fmt.Sprintf("user=ubuntu password=pwd dbname=postgres sslmode=disable")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Parse the request body.
	err = json.NewDecoder(r.Body).Decode(&newBin)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate a new UUID for the bin.
	id := uuid.New().String()

	_, err = db.Exec("INSERT INTO pastebin(id, timestamp, title, content) VALUES($1, $2, $3, $4)", id, time.Now().UTC().Unix(), newBin.Title, newBin.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the UUID.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

// getBinHandler retrieves a bin by UUID.
func getBinHandler(w http.ResponseWriter, r *http.Request) {
	// Get the UUID from the URL path.
	vars := mux.Vars(r)
	id := vars["uuid"]

	psqlInfo := fmt.Sprintf("user=ubuntu password=pwd dbname=postgres sslmode=disable")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var bin Bin
	err = db.QueryRow("SELECT id, timestamp, title, content FROM pastebin WHERE id = $1", id).Scan(&bin.Id, &bin.Timestamp, &bin.Title, &bin.Text)

	if err != nil {
		http.Error(w, "Resource Not Found", http.StatusNotFound)
		return
	}

	fmt.Println(bin)
	// Respond with the bin's text.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bin)
}
