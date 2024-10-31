package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func createBinHandler(w http.ResponseWriter, r *http.Request) {
	var newBin BinRequest

	psqlInfo := fmt.Sprintf("user=ubuntu password=pwd dbname=postgres sslmode=disable")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	err = json.NewDecoder(r.Body).Decode(&newBin)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()

	_, err = db.Exec("INSERT INTO pastebin(id, timestamp, title, content) VALUES($1, $2, $3, $4)", id, time.Now().UTC().Unix(), newBin.Title, newBin.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func starBinHandler(w http.ResponseWriter, r *http.Request) {
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

	// Add +1 to the star counter
	_, err = db.Exec("UPDATE pastebin SET star_counter = star_counter + 1 WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the UUID.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// getBinHandler retrieves a bin by UUID.
func getBinHandler(w http.ResponseWriter, r *http.Request) {
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
	err = db.QueryRow("SELECT id, timestamp, title, content, seen_counter, star_counter FROM pastebin WHERE id = $1", id).Scan(&bin.Id, &bin.Timestamp, &bin.Title, &bin.Text, &bin.SeenCounter, &bin.StarCounter)

	if err != nil {
		http.Error(w, "Resource Not Found", http.StatusNotFound)
		return
	}

	// Add +1 to the seen counter
	_, err = db.Exec("UPDATE pastebin SET seen_counter = seen_counter + 1 WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	bin.SeenCounter += 1

	fmt.Println(bin)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bin)
}

func getLeaderboardHandler(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("user=ubuntu password=pwd dbname=postgres sslmode=disable")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT title, content, star_counter FROM pastebin ORDER BY star_counter desc")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var bins []Bin
	for rows.Next() {
		var bin Bin
		if err := rows.Scan(&bin.Title, &bin.Text, &bin.StarCounter); err != nil {
			fmt.Println(err)
		}
		bins = append(bins, bin)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(bins)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bins)
}
