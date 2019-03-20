package main

import (
	"encoding/json"
	"fmt"
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
	/*router.HandleFunc("/artists/{id}", getArtist).Methods("GET")
	router.HandleFunc("/artists", createArtist).Methods("POST")
	router.HandleFunc("/artists/{id}", deleteArtist).Methods("DELETE")*/

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getArtists(w http.ResponseWriter, r *http.Request) {

	m := make(map[string]Artist)
	for _, v := range client.Keys("Artist*").Val() {
		fmt.Print(v)
		a := client.HGetAll("Artist:" + v).Val()

		fmt.Print(a)
	}

	fmt.Print(m)

	_ = json.NewEncoder(w).Encode(m)
}

/*func getArtist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var a Artist
	if err := json.Unmarshal([]byte(v), &a); err != nil {
		log.Println(err)
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
}*/

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
		client.HSet(i, "FirstName", v.FirstName)
	}

	// check if it worked
	for i := 0; i < len(artists); i++ {
		a := client.HGetAll("Artist:" + artists[i].ID).Val()
		fmt.Print("%s\n", a)
	}
}
