package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"log"
	"net/http"

	"github.com/chandanghosh/goexercises/ex-embedfs/data"
	"github.com/gorilla/mux"
)

//go:embed data/people/people.json
var peopleStr string

//go:embed data/people/people.json
var peopleBytes []byte

type People struct {
	People []Person `json:"people"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(bytes.NewReader(peopleBytes))
	var people People
	_ = dec.Decode(&people)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(people)
}

func main() {
	// fmt.Println(peopleStr)
	// fmt.Println(string(peopleBytes))

	data.GetPeopleFromEmbeddedFile()

	// router
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	log.Println("Server started at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
