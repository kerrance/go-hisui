package app_test

import (
	"github.com/kerrance/go-hisui/app"
	"testing"
)

func TestShouldPokemonNameBeHyphenated(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Provide a hypenated Pokémon name", "chi-yu", true},
		{"Provide a hypenated Pokémon name", "chien-pao", true},
		{"Provide a hypenated Pokémon name", "hakamo-o", true},
		{"Provide a hypenated Pokémon name", "HO-OH", true},
		{"Provide a hypenated Pokémon name", "jangmo-o", true},
		{"Provide a hypenated Pokémon name", "kommo-o", true},
		{"Provide a hypenated Pokémon name", "PORYGON-Z", true},
		{"Provide a hypenated Pokémon name", "ting-lu", true},
		{"Provide a hypenated Pokémon name", "wo-CHIEN", true},
		{"Provide a non-hypenated Pokémon name", "pikachu", false},
		{"Provide a non-hypenated Pokémon name", "CHARIZARD", false},
		{"Provide a non-hypenated Pokémon name", "mr. mime", false},
		{"Provide a non-hypenated Pokémon name", "type: null", false},
		{"Provide a non-hypenated Pokémon name", "iron LEAVES", false},
		{"Provide a non-hypenated Pokémon name", "mew two", false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(testing *testing.T) {
			result := app.ShouldPokemonNameBeHyphenated(testCase.input)
			if result != testCase.expected {
				testing.Errorf(
					"Expected ShouldPokemonNameBeHyphenated(%q) to be %v, but got %v",
					testCase.input,
					testCase.expected,
					result,
				)
			}
		})
	}
}
