package models

type Ability struct {
	Slot     uint8  `njson:"slot"`
	Name     string `njson:"ability.name"`
	IsHidden bool   `njson:"is_hidden"`
}
