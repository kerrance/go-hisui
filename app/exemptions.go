package app

import "strings"

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
