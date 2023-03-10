package app

type Pokemon struct {
	Name      string    `njson:"name"`
	ID        int       `njson:"id"`
	Height    int       `njson:"height"`
	Weight    int       `njson:"weight"`
	Types     []Type    `njson:"types"`
	Abilities []Ability `njson:"abilities"`
}

type Type struct {
	Slot uint8  `njson:"slot"`
	Name string `njson:"type.name"`
}

type Ability struct {
	Slot     uint8  `njson:"slot"`
	Name     string `njson:"ability.name"`
	IsHidden bool   `njson:"is_hidden"`
}
