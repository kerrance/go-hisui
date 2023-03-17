package app

import (
	"fmt"
	"github.com/kerrance/go-hisui/app/models"
)

func Show(pokemon models.Pokemon) {
	fmt.Println("Name:", ConvertStringToTitleCase(pokemon.Name))
	fmt.Println("National Pokédex number:", pokemon.ID)

	fmt.Println("Height:", ConvertDecimetersToMeters(pokemon.Height))
	fmt.Println("Weight:", ConvertHectogramsToKilograms(pokemon.Weight))

	fmt.Println("------- Types -------")
	moreThanOneType := false
	if len(pokemon.Types) > 1 {
		moreThanOneType = true
	}
	for _, pokemonTypes := range pokemon.Types {
		printPokemonTypes(pokemonTypes, moreThanOneType)
	}

	fmt.Println("----- Abilities -----")
	for _, pokemonAbilities := range pokemon.Abilities {
		printPokemonAbilities(pokemonAbilities)
	}
}

func printPokemonTypes(pokemonType models.Type, single bool) {
	if single == true {
		fmt.Print("Type ", pokemonType.Number, ":")
	} else {
		fmt.Print("Type:")
	}

	fmt.Printf(" %+v\n", ConvertStringToTitleCase(pokemonType.Name))
}

func printPokemonAbilities(pokemonAbility models.Ability) {
	if pokemonAbility.IsHidden == true {
		fmt.Print("Hidden Ability:")
	} else {
		fmt.Print("Ability ", pokemonAbility.Number, ":")
	}

	fmt.Printf(" %+v\n", pokemonAbility.Name)
}
