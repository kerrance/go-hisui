package app

import (
	"fmt"
	"github.com/kerrance/go-hisui/app/models"
	"github.com/m7shapan/njson"
	"io"
	"log"
	"net/http"
	"strings"
)

func CreatePokemon(enteredPokemonNameOrPokedexNumber string) models.Pokemon {
	pokemonUrl := "https://pokeapi.co/api/v2/pokemon/"
	req, _ := http.Get(pokemonUrl + strings.ToLower(enteredPokemonNameOrPokedexNumber))

	if req.StatusCode != 200 {
		log.Fatalln("Pokémon not found. Please correct your search term and try again.")
	}
	defer req.Body.Close()

	json, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal("An unexpected error occurred:", err)
	}

	pokemonJson := models.Pokemon{}

	if err := njson.Unmarshal(json, &pokemonJson); err != nil {
		fmt.Println(err)
	}

	return pokemonJson
}
