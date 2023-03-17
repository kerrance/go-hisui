package main

import (
	"fmt"
	"github.com/kerrance/go-hisui/app"
	"log"
)

func main() {
	var enteredPokemonNameOrPokedexNumber string

	fmt.Println("Search for a Pokémon by name: ")

	_, err := fmt.Scanf("%s", &enteredPokemonNameOrPokedexNumber)
	if err != nil {
		log.Fatal("An unexpected error occurred:", err)
	}

	app.Show(app.FindPokemon(enteredPokemonNameOrPokedexNumber))
}
