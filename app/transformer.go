package app

import (
	"fmt"
	"github.com/kerrance/go-hisui/app/models"
	"strings"
)

func Show(pokemon models.Pokemon) {
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("National Pokédex number:", pokemon.ID)

	fmt.Println("Height:", ConvertDecimetersToMeters(pokemon.Height))
	fmt.Println("Weight:", ConvertHectogramsToKilograms(pokemon.Weight))

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

func printPokemonAbilities(ability models.Ability) {
	if ability.IsHidden == true {
		fmt.Print("Hidden Ability:")
	} else {
		fmt.Print("Ability ", ability.Slot, ":")
	}

	fmt.Printf(" %+v\n", strings.ToTitle(ability.Name))
}
