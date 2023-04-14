package app_test

import (
	"github.com/kerrance/go-hisui/app"
	"testing"
)

func TestSanitisePokemonNameForSearch(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Provide a Pokémon name with an apostrophe", "farfetch'd", "farfetchd"},
		{"Provide a Pokémon name with a diacritic", "flabebé", "flabebe"},
		{"Provide a Pokémon name with a diacritic", "flabébe", "flabebe"},
		{"Provide a Pokémon name with a diacritic", "flabébé", "flabebe"},
		{"Provide a Pokémon name with no space which needs one", "mimejr", "mime-jr"},
		{"Provide a Pokémon name with a full-stop", "mimejr.", "mime-jr"},
		{"Provide a Pokémon name with a space and a full stop", "mime jr.", "mime-jr"},
		{"Provide a Pokémon name with no space which needs one", "mrmime", "mr-mime"},
		{"Provide a Pokémon name with a full-stop", "mr.mime", "mr-mime"},
		{"Provide a Pokémon name with a space and a full stop", "mr. mime", "mr-mime"},
		{"Provide a Pokémon name with no space which needs one", "mrrime", "mr-rime"},
		{"Provide a Pokémon name with a full-stop", "mr.rime", "mr-rime"},
		{"Provide a Pokémon name with a space and a full stop", "mr. rime", "mr-rime"},
		{"Provide a Pokémon name with an apostrophe", "sirfetch'd", "sirfetchd"},
		{"Provide a Pokémon name with no space which needs one", "typenull", "type-null"},
		{"Provide a Pokémon name with a colon", "type:null", "type-null"},
		{"Provide a Pokémon name with a space and a colon", "type: null", "type-null"},
		{"Provide a Pokémon name which should not be sanitised", "pikachu", "pikachu"},
		{"Provide a Pokémon name which should not be sanitised", "CHARIZARD", "CHARIZARD"},
		{"Provide a Pokémon name which should not be sanitised", "iron LEAVES", "iron LEAVES"},
		{"Provide a Pokémon name which should not be sanitised", "mew two", "mew two"},
		{"Provide a Pokémon name which should not be sanitised", "tapu-koko", "tapu-koko"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(testing *testing.T) {
			result := app.SanitisePokemonNameForSearch(testCase.input)
			if result != testCase.expected {
				testing.Errorf(
					"Expected SanitisePokemonNameForSearch(%q) to be %v, but got %v",
					testCase.input,
					testCase.expected,
					result,
				)
			}
		})
	}
}

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
		{"Provide a non-hypenated Pokémon name", "tapu-koko", false},
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

func TestDisplayPokemonNameWithPunctuationMarkIfNeeded(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Provide a Pokémon name which requires an apostrophe", "farfetchd", "farfetch'd"},
		{"Provide a Pokémon name which requires a diacritic", "flabebe", "flabébé"},
		{"Provide a Pokémon name which requires a space and a full stop", "mime-jr", "mime jr."},
		{"Provide a Pokémon name which requires a space and a full stop", "mr-mime", "mr. mime"},
		{"Provide a Pokémon name which requires a space and a full stop", "mr-rime", "mr. rime"},
		{"Provide a Pokémon name which requires an apostrophe", "sirfetchd", "sirfetch'd"},
		{"Provide a Pokémon name which requires a space and a colon", "type-null", "type: null"},
		{"Provide a Pokémon name which should not be sanitised", "pikachu", "pikachu"},
		{"Provide a Pokémon name which should not be sanitised", "CHARIZARD", "CHARIZARD"},
		{"Provide a Pokémon name which should not be sanitised", "iron LEAVES", "iron LEAVES"},
		{"Provide a Pokémon name which should not be sanitised", "mew two", "mew two"},
		{"Provide a Pokémon name which should not be sanitised", "tapu-koko", "tapu-koko"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(testing *testing.T) {
			result := app.DisplayPokemonNameWithPunctuationMarkIfNeeded(testCase.input)
			if result != testCase.expected {
				testing.Errorf(
					"Expected DisplayPokemonNameWithPunctuationMarkIfNeeded(%q) to be %v, but got %v",
					testCase.input,
					testCase.expected,
					result,
				)
			}
		})
	}
}
