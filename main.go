package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/m7shapan/njson"
)

type Pokemon struct {
	Name      string    `njson:"name"`
	ID        int       `njson:"id"`
	Height    int       `njson:"height"`
	Weight    int       `njson:"weight"`
	Type1     string    `njson:"types.0.type.name"`
	Type2     string    `njson:"types.1.type.name"`
	Abilities []Ability `njson:"abilities"`
}

type Ability struct {
	Slot     uint8  `njson:"slot"`
	Name     string `njson:"ability.name"`
	IsHidden bool   `njson:"is_hidden"`
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

	fmt.Println("Height:", convertDecimetersToMeters(pokemon.Height))
	fmt.Println("Weight:", convertHectogramsToKilograms(pokemon.Weight))

	if pokemon.Type2 == "" {
		fmt.Println("Type:", pokemon.Type1)
	} else {
		fmt.Println("Type 1:", pokemon.Type1)
		fmt.Println("Type 2:", pokemon.Type2)
	}

	fmt.Println("----- Abilities -----")
	for i, pokemonAbilities := range pokemon.Abilities {
		printPokemonAbilities(pokemonAbilities)
		i++
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

func printPokemonAbilities(ability Ability) {
	if ability.IsHidden == true {
		fmt.Print("Hidden Ability:")
	} else {
		fmt.Print("Ability ", ability.Slot, ":")
	}

	fmt.Printf(" %+v\n", strings.ToTitle(ability.Name))
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

	pokeJson.Name = strings.ToTitle(pokeJson.Name)
	pokeJson.Type1 = strings.ToTitle(pokeJson.Type1)
	pokeJson.Type2 = strings.ToTitle(pokeJson.Type2)

	return pokeJson
}
