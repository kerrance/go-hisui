package models

type Pokemon struct {
	Name      string    `njson:"name"`
	ID        int       `njson:"id"`
	Height    int       `njson:"height"`
	Weight    int       `njson:"weight"`
	Types     []Type    `njson:"types"`
	Abilities []Ability `njson:"abilities"`
}
