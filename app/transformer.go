package app

import (
	"fmt"
	"github.com/kerrance/go-hisui/app/models"
)

func Show(pokemon models.Pokemon) {
	printPokemonName(pokemon.Name)
	fmt.Printf("National Pokédex number: %0*d\n", 4, pokemon.ID)
	fmt.Println("Height:", ConvertDecimetersToMeters(pokemon.Height))
	fmt.Println("Weight:", ConvertHectogramsToKilograms(pokemon.Weight))

	fmt.Println("------- Types -------")
	singleType := false
	if len(pokemon.Types) == 1 {
		singleType = true
	}
	for _, pokemonTypes := range pokemon.Types {
		printPokemonTypes(pokemonTypes, singleType)
	}

	fmt.Println("----- Abilities -----")
	if (len(pokemon.Abilities) == 2) && (pokemon.Abilities[0].Name == pokemon.Abilities[1].Name) {
		fmt.Print("Ability:")
		_, err := printPokemonAbility(pokemon.Abilities[0])
		if err != nil {
			return
		}
	} else {
		for _, pokemonAbilities := range pokemon.Abilities {
			printPokemonAbilities(pokemonAbilities)
		}
	}
}

func printPokemonName(pokemonName string) {
	pokemonName = DisplayPokemonNameWithPunctuationMarkIfNeeded(pokemonName)

	pokemonName = ConvertStringToTitleCase(pokemonName)
	if ShouldPokemonNameBeHyphenated(pokemonName) == true {
		pokemonName = ReplaceHyphensWithSpaces(pokemonName)
	} else {
		pokemonName = ReplaceSpacesWithHyphens(pokemonName)
	}

	fmt.Println("Name:", pokemonName)
}

func printPokemonTypes(pokemonType models.Type, singleType bool) {
	if singleType == true {
		fmt.Print("Type:")
	} else {
		fmt.Print("Type ", pokemonType.Number, ":")
	}

	fmt.Printf(" %+v\n", ConvertStringToTitleCase(pokemonType.Name))
}

func printPokemonAbilities(pokemonAbility models.Ability) {
	if pokemonAbility.IsHidden == true {
		fmt.Print("Hidden Ability:")
	} else {
		fmt.Print("Ability ", pokemonAbility.Number, ":")
	}

	_, err := printPokemonAbility(pokemonAbility)
	if err != nil {
		return
	}
}

func printPokemonAbility(pokemonAbility models.Ability) (int, error) {
	return fmt.Printf(" %+v\n", ReplaceSpacesWithHyphens(ConvertStringToTitleCase(pokemonAbility.Name)))
}
