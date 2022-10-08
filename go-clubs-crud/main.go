package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Club struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Coach *Coach `json:"coach"`
}

type Coach struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}

var clubs []Club

func getClubs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clubs)

}

func getClub(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range clubs {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createClub(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var club Club
	_ = json.NewDecoder(r.Body).Decode(&club)
	club.ID = strconv.Itoa(rand.Intn(100000000))
	clubs = append(clubs, club)
	json.NewEncoder(w).Encode(club)
}

func updateClub(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range clubs {
		if item.ID == params["id"] {
			clubs = append(clubs[:index], clubs[index+1:]...)
			var club Club
			_ = json.NewDecoder(r.Body).Decode(&club)
			club.ID = params["id"]
			clubs = append(clubs, club)
			json.NewEncoder(w).Encode(club)
		}
	}
}

func deleteClub(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range clubs {
		if item.ID == params["id"] {
			clubs = append(clubs[:index], clubs[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(clubs)
}

func main() {
	r := mux.NewRouter()

	clubs = append(clubs, Club{ID: "1", Name: "Paradise", Coach: &Coach{Name: "Ollie Paradise", Level: "UKCC Level 1"}})
	clubs = append(clubs, Club{ID: "2", Name: "Heart", Coach: &Coach{Name: "Wayne Smith", Level: "HPC"}})
	r.HandleFunc("/clubs", getClubs).Methods("GET")
	r.HandleFunc("/clubs/{id}", getClub).Methods("GET")
	r.HandleFunc("/clubs", createClub).Methods("POST")
	r.HandleFunc("/clubs/{id}", updateClub).Methods("PUT")
	r.HandleFunc("/clubs/{id}", deleteClub).Methods("DELETE")

	fmt.Printf("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
