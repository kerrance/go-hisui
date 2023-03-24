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
	pokemonApiUrl := "https://pokeapi.co/api/v2/pokemon/"
	req, _ := http.Get(pokemonApiUrl + ConvertStringToKebabCase(enteredPokemonNameOrPokedexNumber))
	if req.StatusCode != 200 {
		log.Fatalln("Pokémon not found. Please correct your search term and try again.")
	}

	json, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal("An unexpected error occurred:", err)
	}

	foundPokemonJson := models.Pokemon{}
	if err := njson.Unmarshal(json, &foundPokemonJson); err != nil {
		fmt.Println(err)
	}

	return foundPokemonJson
}
