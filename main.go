package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Artist struct {
	ID        string
	FirstName string
	LastName  string
	Pseudonym string
}

var artists []Artist

func main() {
	router := mux.NewRouter()

	popupateArtists()

	router.HandleFunc("/artists", getArtists).Methods("GET")
	router.HandleFunc("/artists/{id}", getArtist).Methods("GET")
	router.HandleFunc("/artists", createArtist).Methods("POST")
	router.HandleFunc("/artists/{id}", deleteArtist).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getArtists(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(artists)
}

func getArtist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range artists {
		if item.ID == params["id"] {
			_ = json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createArtist(w http.ResponseWriter, r *http.Request) {
	var artist Artist

	_ = json.NewDecoder(r.Body).Decode(&artist)
	artists = append(artists, artist)
	_ = json.NewEncoder(w).Encode(artists)
}

func deleteArtist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range artists {
		if item.ID == params["id"] {
			artists = append(artists[:index], artists[index+1:]...)
			_ = json.NewEncoder(w).Encode(artists)
			return
		}
	}
}

func popupateArtists() {
	artists = append(artists,
		Artist{"1", "Jamie N", "Commons", ""},
		Artist{"2", "Joshua", "James", ""},
		Artist{"3", "Jake", "Smith", "The White Buffalo"},
		Artist{"4", "Dan", "Auerbach", ""},
		Artist{"5", "Troy", "Baker", ""},
	)
}
