package main

import (
	"log"

	con "github.com/otyamura/pokewordle_solver/internal/connection"
)

func main() {
	_, r := con.CreateConnection()
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
