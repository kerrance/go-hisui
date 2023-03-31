package models_test

import (
	"github.com/kerrance/go-hisui/app/models"
	"testing"
)

func TestPokemon(t *testing.T) {
	testPokemon := models.Pokemon{
		Name:   "marill",
		ID:     183,
		Height: 4,
		Weight: 85,
		Types: []models.Type{
			{Number: 1, Name: "water"},
			{Number: 2, Name: "fairy"},
		},
		Abilities: []models.Ability{
			{Number: 1, Name: "thick-fat", IsHidden: false},
			{Number: 2, Name: "huge-power", IsHidden: false},
			{Number: 3, Name: "sap-sipper", IsHidden: true},
		},
	}

	if testPokemon.Name != "marill" {
		t.Errorf("Expected Name to be \"marill\", but got \"%s\"", testPokemon.Name)
	}

	if testPokemon.ID != 183 {
		t.Errorf("Expected ID to be 183, but got %d", testPokemon.ID)
	}

	if testPokemon.Height != 4 {
		t.Errorf("Expected Height to be 4, but got %d", testPokemon.Height)
	}

	if testPokemon.Weight != 85 {
		t.Errorf("Expected Weight to be 85, but got %d", testPokemon.Weight)
	}

	if testPokemon.Types[0].Name != "water" {
		t.Errorf("Expected first type to be \"water\", but got \"%s\"", testPokemon.Types[0].Name)
	}

	if testPokemon.Types[1].Name != "fairy" {
		t.Errorf("Expected second type to be \"fairy\", but got \"%s\"", testPokemon.Types[1].Name)
	}

	if len(testPokemon.Types) != 2 {
		t.Errorf("Expected 2 types, but got %d", len(testPokemon.Types))
	}

	if testPokemon.Abilities[0].Name != "thick-fat" {
		t.Errorf("Expected second ability to be \"thick-fat\", but got \"%s\"", testPokemon.Abilities[1].Name)
	}

	if testPokemon.Abilities[0].IsHidden != false {
		t.Errorf("Expected second ability IsHidden to be false, but got %v", testPokemon.Abilities[1].IsHidden)
	}

	if testPokemon.Abilities[1].Name != "huge-power" {
		t.Errorf("Expected second ability to be \"huge-power\", but got \"%s\"", testPokemon.Abilities[1].Name)
	}

	if testPokemon.Abilities[1].IsHidden != false {
		t.Errorf("Expected second ability IsHidden to be false, but got %v", testPokemon.Abilities[1].IsHidden)
	}

	if testPokemon.Abilities[2].Name != "sap-sipper" {
		t.Errorf("Expected first ability to be \"sap-sipper\", but got \"%s\"", testPokemon.Abilities[0].Name)
	}

	if testPokemon.Abilities[2].IsHidden != true {
		t.Errorf("Expected first ability IsHidden to be true, but got %v", testPokemon.Abilities[0].IsHidden)
	}

	if len(testPokemon.Abilities) != 3 {
		t.Errorf("Expected 3 abilities, but got %d", len(testPokemon.Abilities))
	}
}
