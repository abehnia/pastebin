package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/bins", createBinHandler).Methods("POST")
	r.HandleFunc("/bins/{uuid}", getBinHandler).Methods("GET")
	r.HandleFunc("/bins/{uuid}/star", starBinHandler).Methods("POST")
	r.HandleFunc("/leaderboard", getLeaderboardHandler).Methods("GET")
	return r
}
