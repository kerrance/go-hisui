package models

type Type struct {
	Number uint8  `njson:"slot"`
	Name   string `njson:"type.name"`
}
