package main

import (
	"github.com/otyamura/pokewordle_solver/internal/connection"
	"github.com/otyamura/pokewordle_solver/internal/load"
	"github.com/otyamura/pokewordle_solver/types"
)

func main() {
	db := connection.CreateDBConnection()

	db.AutoMigrate(&types.Pokemon{})
	pokes := load.LoadPokes()
	db.Create(&pokes)
}
