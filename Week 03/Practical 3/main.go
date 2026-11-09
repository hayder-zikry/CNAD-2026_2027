package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

// SongData struct representing data of a song in the application
type SongData struct {
	id        string
	title     string
	artist    string
	genre     string
	language  string
	available bool
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome to the World Greatest Music Library API!"
	})
}
