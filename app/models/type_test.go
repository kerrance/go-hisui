package models_test

import (
	"github.com/kerrance/go-hisui/app/models"
	"testing"
)

func TestType(t *testing.T) {
	testType := models.Type{
		Number: 1,
		Name:   "water",
	}

	if testType.Number != 1 {
		t.Errorf("Expected Number to be 1, but got %d", testType.Number)
	}

	if testType.Name != "water" {
		t.Errorf("Expected Name to be \"water\", but got \"%s\"", testType.Name)
	}
}
