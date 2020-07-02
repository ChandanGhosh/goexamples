package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ResposeData struct {
	Name     string    `json:"name"`
	Pokemons []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNum int     `json:"entry_number"`
	Species  Species `json:"pokemon_species"`
}

type Species struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
	// rawHttpGet()
	structHttpGet()
}

func structHttpGet() {
	res, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		log.Panicln(err.Error())
		os.Exit(1)
	}
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
	}
	var resposeData ResposeData
	json.Unmarshal(resData, &resposeData)
	fmt.Println(resposeData.Name)
	fmt.Println(len(resposeData.Pokemons))
}

func rawHttpGet() {
	res, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		log.Panicln(err.Error())
		os.Exit(1)
	}
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(string(resData))
}
