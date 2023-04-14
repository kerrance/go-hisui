package app

import "strings"

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
