package main

import (
	"log"
	"os"

	con "github.com/otyamura/pokewordle_solver/internal/connection"
)

func main() {
	_, r := con.CreateConnection()
	p := os.Getenv("PORT")
	if p != "" {
		p = ":" + p
	} else {
		p = ":8080"
	}
	err := r.Run(p)
	if err != nil {
		log.Fatal(err)
	}
}
