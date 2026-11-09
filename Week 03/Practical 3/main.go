package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// SongData struct representing data of a song in the application
type Song struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Artist    string `json:"artist"`
	Genre     string `json:"genre"`
	Language  string `json:"language"`
	Available bool   `json:"available"`
}

// Songs map with Song ID as the key
var songs = map[string]Song{
	"1": {ID: "1", Title: "Until I Drown", Artist: "Tears for Fears", Genre: "Rock", Language: "English", Available: true},
	"2": {ID: "2", Title: "Kangaroo Court", Artist: "Capital Cities", Genre: "Pop", Language: "English", Available: true},
	"3": {ID: "3", Title: "Frou-frou Foxes in Midsummer Fires", Artist: "Cocteau Twins", Genre: "Pop", Language: "English", Available: true},
	"4": {ID: "4", Title: "Instant Crush", Artist: "Daft Punk", Genre: "Rock", Language: "English", Available: false},
	"5": {ID: "5", Title: "A Gira", Artist: "Trio Ternura", Genre: "Samba", Language: "Portuguese", Available: false},
	"6": {ID: "6", Title: "Tarha Tadagh", Artist: "Imarhan", Genre: "Blues", Language: "Tamasheq", Available: true},
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome to the World Greatest Music Library API!",
	})
}

func getSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result := []Song{}
	for _, song := range songs {
		result = append(result, song)
	}
	json.NewEncoder(w).Encode(result)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/welcome", welcome).Methods("GET")
	router.HandleFunc("/api/v1/songs", getSongs).Methods("GET")

	fmt.Println("Server running on http://localhost:5000")
	http.ListenAndServe(":5000", router)
}
