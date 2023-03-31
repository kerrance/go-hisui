package app

func ShouldPokemonNameBeHyphenated(pokemonName string) bool {
	switch pokemonName {
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
