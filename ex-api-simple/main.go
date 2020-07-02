package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for _, item := range people {
		if item.ID == id {
			_ = json.NewEncoder(w).Encode(item)
			return
		}
	}
	_ = json.NewEncoder(w).Encode(&Person{})
}
func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(people)
}
func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = id
	people = append(people, person)
	_ = json.NewEncoder(w).Encode(person)
}
func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for index, item := range people {
		if item.ID == id {
			people = append(people[:index], people[index+1:]...)
			break
		}
		_ = json.NewEncoder(w).Encode(people)
	}
}

func main() {
	people = append(people, Person{ID: "1", Firstname: "Chandan", Lastname: "Ghosh", Address: &Address{City: "Baidyapur", State: "WB"}})
	people = append(people, Person{ID: "2", Firstname: "Arnab", Lastname: "Ghosh", Address: &Address{City: "Sapoorji", State: "WB"}})
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
