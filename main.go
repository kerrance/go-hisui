package main

import (
	"fmt"
	"github.com/404a10/MicroDex"
)

func main() {
	var enteredPokemonName string

	fmt.Println("Search for a Pokémon by name: ")
	fmt.Scanf("%s", &enteredPokemonName)
	show(microdex.CreatePokemon(enteredPokemonName))
}

func show(pokemon microdex.Pokemon) {
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("ID:", pokemon.ID)
	if pokemon.Type_2 == "" {
		fmt.Println("Type:", pokemon.Type_1)
	} else {
		fmt.Println("Type 1:", pokemon.Type_1)
		fmt.Println("Type 2:", pokemon.Type_2)
	}
}
