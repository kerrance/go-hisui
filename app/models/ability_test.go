package models_test

import (
	"github.com/kerrance/go-hisui/app/models"
	"testing"
)

func TestAbility(t *testing.T) {
	testAbility := models.Ability{
		Number:   1,
		Name:     "swift-swim",
		IsHidden: true,
	}

	if testAbility.Number != 1 {
		t.Errorf("Expected Number to be 1, but got %d", testAbility.Number)
	}

	if testAbility.Name != "swift-swim" {
		t.Errorf("Expected Name to be \"swift-swim\", but got \"%s\"", testAbility.Name)
	}

	if testAbility.IsHidden != true {
		t.Errorf("Expected IsHidden to be true, but got %v", testAbility.IsHidden)
	}
}
