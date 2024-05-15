package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type Pokemons []Pokemon

type Pokemon struct {
	ID   int `json:"id"`
	Name struct {
		English  string `json:"english"`
		Japanese string `json:"japanese"`
		Chinese  string `json:"chinese"`
		French   string `json:"french"`
	} `json:"name"`
	Type []string `json:"type"`
	Base struct {
		Hp        int `json:"HP"`
		Attack    int `json:"Attack"`
		Defense   int `json:"Defense"`
		SpAttack  int `json:"Sp. Attack"`
		SpDefense int `json:"Sp. Defense"`
		Speed     int `json:"Speed"`
	} `json:"base"`
}

var (
	ErrFailed = errors.New("Failed")
)

func main() {
	// POKEMON
	file, err := os.Open("pokemon.json")
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
	}
	defer file.Close()
	var pokemons Pokemons
	err = json.NewDecoder(file).Decode(&pokemons)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
	}
	fmt.Println(len(pokemons))

	httpClient := http.DefaultClient
	res, err := httpClient.Get("https://raw.githubusercontent.com/fanzeyi/pokemon.json/master/pokedex.json")
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
	}
	defer res.Body.Close()
	var pokemons2 Pokemons
	err = json.NewDecoder(res.Body).Decode(&pokemons2)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
	}
	fmt.Println(len(pokemons2))
}
