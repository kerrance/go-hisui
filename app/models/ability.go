package models

type Ability struct {
	Number   uint8  `njson:"slot"`
	Name     string `njson:"ability.name"`
	IsHidden bool   `njson:"is_hidden"`
}
