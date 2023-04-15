package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ClarifyPokemonNameForSearch(pokemonName string) string {
	switch strings.ToLower(pokemonName) {
	case "nidoran":
		for {
			fmt.Println("Did you mean Nidoran♀ or Nidoran♂?")

			nidoranGenderInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				log.Fatal("An unexpected error occurred:", err)
			}
			enteredNidoranGender := strings.TrimSpace(strings.ToLower(nidoranGenderInput))

			if enteredNidoranGender == "♀" || enteredNidoranGender == "female" || enteredNidoranGender == "f" {
				pokemonName = "nidoran-f"
				break
			} else if enteredNidoranGender == "♂" || enteredNidoranGender == "male" || enteredNidoranGender == "m" {
				pokemonName = "nidoran-m"
				break
			} else {
				fmt.Println("Invalid gender input. Please enter '♀', '♂', 'female' or 'male'")
			}
		}
	case "morpeko":
		for {
			fmt.Println("Did you mean Morpeko Hangry Mode or Morpeko Full Belly Mode?")

			morpekoModeInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				log.Fatal("An unexpected error occurred:", err)
			}
			enteredMorpekoMode := strings.TrimSpace(strings.ToLower(morpekoModeInput))

			if enteredMorpekoMode == "hangry mode" || enteredMorpekoMode == "hangry" {
				pokemonName = "morpeko-hangry"
				break
			} else if enteredMorpekoMode == "full belly mode" || enteredMorpekoMode == "full belly" {
				pokemonName = "morpeko-full-belly"
				break
			} else {
				fmt.Println("Invalid mode input. Please enter 'Hangry Mode' or 'Full Belly Mode'")
			}
		}
	}

	return pokemonName
}

func SanitisePokemonNameForSearch(pokemonName string) string {
	switch strings.ToLower(pokemonName) {
	case "farfetch'd":
		return "farfetchd"
	case "flabebé", "flabébe", "flabébé":
		return "flabebe"
	case "mimejr", "mimejr.", "mime jr.":
		return "mime-jr"
	case "mrmime", "mr.mime", "mr. mime":
		return "mr-mime"
	case "mrrime", "mr.rime", "mr. rime":
		return "mr-rime"
	case "sirfetch'd":
		return "sirfetchd"
	case "typenull", "type:null", "type: null":
		return "type-null"
	}
	return pokemonName
}

func ShouldPokemonNameBeHyphenated(pokemonName string) bool {
	switch strings.ToLower(pokemonName) {
	case
		"chi-yu",
		"chien-pao",
		"hakamo-o",
		"ho-oh",
		"jangmo-o",
		"kommo-o",
		"porygon-z",
		"ting-lu",
		"wo-chien":
		return true
	}

	return false
}

func DisplayPokemonNameWithPunctuationMarkIfNeeded(pokemonName string) string {
	switch strings.ToLower(pokemonName) {
	case "farfetchd":
		return "farfetch'd"
	case "flabebe":
		return "flabébé"
	case "mime-jr":
		return "mime jr."
	case "mr-mime":
		return "mr. mime"
	case "mr-rime":
		return "mr. rime"
	case "sirfetchd":
		return "sirfetch'd"
	case "type-null":
		return "type: null"
	}

	return pokemonName
}
