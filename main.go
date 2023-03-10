package main

import (
	"fmt"
	"github.com/kerrance/go-hisui/app"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/m7shapan/njson"
)

func main() {
	var enteredPokemonName string

	fmt.Println("Search for a Pokémon by name: ")
	fmt.Scanf("%s", &enteredPokemonName)
	show(createPokemon(enteredPokemonName))
}

func show(pokemon app.Pokemon) {
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("National Pokédex number:", pokemon.ID)

	fmt.Println("Height:", convertDecimetersToMeters(pokemon.Height))
	fmt.Println("Weight:", convertHectogramsToKilograms(pokemon.Weight))

	fmt.Println("------- Types -------")
	for _, pokemonTypes := range pokemon.Types {
		fmt.Print("Type ", pokemonTypes.Slot, ":")

		fmt.Printf(" %+v\n", strings.ToTitle(pokemonTypes.Name))
	}

	fmt.Println("----- Abilities -----")
	for _, pokemonAbilities := range pokemon.Abilities {
		printPokemonAbilities(pokemonAbilities)
	}
}

func convertDecimetersToMeters(height int) string {
	if height >= 10 {
		var heightAsFloat = float64(height)
		var adjustedHeight = heightAsFloat / float64(10)
		return fmt.Sprintf("%g", adjustedHeight) + "m"
	} else {
		return "0." + strconv.Itoa(height) + "m"
	}
}

func convertHectogramsToKilograms(weight int) string {
	if weight >= 10 {
		var weightAsFloat = float64(weight)
		var adjustedWeight = weightAsFloat / float64(10)
		return fmt.Sprintf("%g", adjustedWeight) + "kg"
	} else {
		var weightAsFloat = float64(weight)
		var adjustedWeight = weightAsFloat * 2.2
		var weightAsPounds = adjustedWeight / float64(10)
		return fmt.Sprintf("%.1f", weightAsPounds) + "lbs"
	}
}

func printPokemonAbilities(ability app.Ability) {
	if ability.IsHidden == true {
		fmt.Print("Hidden Ability:")
	} else {
		fmt.Print("Ability ", ability.Slot, ":")
	}

	fmt.Printf(" %+v\n", strings.ToTitle(ability.Name))
}

func createPokemon(name string) app.Pokemon {
	pokemonUrl := "https://pokeapi.co/api/v2/pokemon/"
	req, _ := http.Get(pokemonUrl + strings.ToLower(name))

	if req.StatusCode != 200 {
		log.Fatalln("Pokémon not found. Please correct your search term and try again.")
	}
	defer req.Body.Close()

	pokeJson := app.Pokemon{}

	json, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal("An unexpected error occurred:", err)
	}

	if err := njson.Unmarshal(json, &pokeJson); err != nil {
		fmt.Println(err)
	}

	pokeJson.Name = strings.ToTitle(pokeJson.Name)

	return pokeJson
}
