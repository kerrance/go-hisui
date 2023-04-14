package app

import (
	"fmt"
	"github.com/kerrance/go-hisui/app/models"
	"github.com/m7shapan/njson"
	"io"
	"log"
	"net/http"
)

func FindPokemon(enteredPokemonNameOrPokedexNumber string) models.Pokemon {
	pokeApiPokemonUrl, _ := http.Get(
		"https://pokeapi.co/api/v2/pokemon/" + ConvertStringToKebabCase(enteredPokemonNameOrPokedexNumber),
	)
	if pokeApiPokemonUrl.StatusCode != 200 {
		log.Fatalln("Pokémon not found. Please correct your search term and try again.")
	}

	pokemonJson, err := io.ReadAll(pokeApiPokemonUrl.Body)
	if err != nil {
		log.Fatal("An unexpected error occurred:", err)
	}

	foundPokemonJson := models.Pokemon{}
	if err := njson.Unmarshal(pokemonJson, &foundPokemonJson); err != nil {
		fmt.Println(err)
	}

	return foundPokemonJson
}
