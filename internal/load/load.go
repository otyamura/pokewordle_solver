package load

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/otyamura/pokewordle_solver/types"
)

func LoadPokeRaws() []types.PokemonRaw {
	f, err := os.OpenFile("./csv/pokes.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	pokes := []types.PokemonRaw{}
	if err := gocsv.UnmarshalFile(f, &pokes); err != nil {
		log.Fatal(err)
	}
	return pokes
}

func LoadPokes() []types.Pokemon {
	f, err := os.OpenFile("./csv/clean.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	pokes := []types.Pokemon{}
	if err := gocsv.UnmarshalFile(f, &pokes); err != nil {
		log.Fatal(err)
	}
	return pokes
}
