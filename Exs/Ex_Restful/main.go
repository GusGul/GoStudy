// http://pokeapi.co/api/v2/pokedex/kanto/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Name    string    `json:"name"`
	Pokemon []pokemon `json:"pokemon_entries"`
}

type pokemon struct {
	Number int            `json:"entry_number"`
	Specie pokemonSpecies `json:"pokemon_species"`
}

type pokemonSpecies struct {
	Name string `json:"name"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)

	for _, pokemon := range responseObject.Pokemon {
		fmt.Println(pokemon.Number, pokemon.Specie.Name)
	}
}
