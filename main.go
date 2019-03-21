package main

import (
	"encoding/json"
	"github.com/cwd-nial/go-api/storages"
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

var client = storages.GetRedisClient()

func main() {
	router := mux.NewRouter()

	populateArtists()

	router.HandleFunc("/artists", getArtists).Methods("GET")
	router.HandleFunc("/artists/{id}", getArtist).Methods("GET")
	router.HandleFunc("/artists", createArtist).Methods("POST")
	router.HandleFunc("/artists/{id}", deleteArtist).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getArtists(w http.ResponseWriter, r *http.Request) {

	var artists []Artist

	for _, v := range client.Keys("Artist:*").Val() {
		itemData := client.HGetAll(v).Val()
		artists = append(artists, Artist{
			itemData["ID"],
			itemData["FirstName"],
			itemData["LastName"],
			itemData["Pseudonym"],
		})
	}
	_ = json.NewEncoder(w).Encode(artists)
}

func getArtist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var artist Artist
	itemData := client.HGetAll("Artist:" + params["id"]).Val()
	artist = Artist{
		itemData["ID"],
		itemData["FirstName"],
		itemData["LastName"],
		itemData["Pseudonym"],
	}
	_ = json.NewEncoder(w).Encode(artist)
}

func createArtist(w http.ResponseWriter, r *http.Request) {
	var artist Artist

	_ = json.NewDecoder(r.Body).Decode(&artist)
	i := "Artist:" + artist.ID
	client.HSet(i, "ID", artist.ID)
	client.HSet(i, "FirstName", artist.FirstName)
	client.HSet(i, "LastName", artist.LastName)
	client.HSet(i, "Pseudonym", artist.Pseudonym)
	_ = json.NewEncoder(w).Encode(artist)
}

func deleteArtist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	client.Del("Artist:" + params["id"])
}

func populateArtists() {

	// hard-coded artists
	var artists []Artist
	artists = append(artists,
		Artist{"1", "Jamie N", "Commons", ""},
		Artist{"2", "Joshua", "James", ""},
		Artist{"3", "Jake", "Smith", "The White Buffalo"},
		Artist{"4", "Dan", "Auerbach", ""},
		Artist{"5", "Troy", "Baker", ""},
	)

	// add to redis DB
	for _, v := range artists {
		i := "Artist:" + v.ID
		client.HSet(i, "ID", v.ID)
		client.HSet(i, "FirstName", v.FirstName)
		client.HSet(i, "LastName", v.LastName)
		client.HSet(i, "Pseudonym", v.Pseudonym)
	}
}
