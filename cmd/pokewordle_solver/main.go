package main

import con "github.com/otyamura/pokewordle_solver/internal/connection"

func main() {
	_, r := con.CreateConnection()
	r.Run()
}
