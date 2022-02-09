package types

import "gorm.io/gorm"

type PokemonRaw struct {
	Name       string `csv:"name"`
	Generation string `csv:"generation"`
}

type Pokemon struct {
	Name       string `csv:"name"`
	Generation int    `csv:"generation"`
}

type PokeGorm struct {
	gorm.Model
	Name       string `csv:"name"`
	Generation int    `csv:"generation"`
}
