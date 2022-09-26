package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/m7shapan/njson"
)

type Pokemon struct {
	Name   string `njson:"name"`
	ID     int    `njson:"id"`
	Height int    `njson:"height"`
	Weight int    `njson:"weight"`
	Type_1 string `njson:"types.0.type.name"`
	Type_2 string `njson:"types.1.type.name"`
}

func main() {
	var enteredPokemonName string

	fmt.Println("Search for a Pokémon by name: ")
	fmt.Scanf("%s", &enteredPokemonName)
	show(createPokemon(enteredPokemonName))
}

func show(pokemon Pokemon) {
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("National Pokédex number:", pokemon.ID)

	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)

	if pokemon.Type_2 == "" {
		fmt.Println("Type:", pokemon.Type_1)
	} else {
		fmt.Println("Type 1:", pokemon.Type_1)
		fmt.Println("Type 2:", pokemon.Type_2)
	}
}

func createPokemon(name string) Pokemon {
	pokemonUrl := "https://pokeapi.co/api/v2/pokemon/"
	req, _ := http.Get(pokemonUrl + strings.ToLower(name))

	if req.StatusCode != 200 {
		log.Fatalln("Pokémon not found. Please correct your search term and try again.")
	}
	defer req.Body.Close()

	pokeJson := Pokemon{}

	json, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal("An unexpected error occurred:", err)
	}

	if err := njson.Unmarshal(json, &pokeJson); err != nil {
		fmt.Println(err)
	}

	pokeJson.Name = strings.Title(pokeJson.Name)
	pokeJson.Type_1 = strings.Title(pokeJson.Type_1)
	pokeJson.Type_2 = strings.Title(pokeJson.Type_2)

	var pokeJsonHeight int
	pokeJsonHeight = pokeJson.Height
	pokeJson.Height = convertHeight(pokeJson.Height)

	return pokeJson
}

func convertHeight(height int) string {
	fmt.Println(height)

	return "test height"
}
