package main

import (
	"log"

	"github.com/otyamura/pokewordle_solver/internal/connection"
	"github.com/otyamura/pokewordle_solver/internal/load"
	"github.com/otyamura/pokewordle_solver/types"
)

func main() {
	db := connection.CreateDBConnection()

	err := db.AutoMigrate(&types.Pokemon{})
	if err != nil {
		log.Fatal(err)
	}
	pokes := load.LoadPokes()
	db.Create(&pokes)
}
