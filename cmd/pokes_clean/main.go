package main

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/otyamura/pokewordle_solver/internal/preprocessing"
)

func main() {
	pokes := preprocessing.Cleaning()
	file, _ := os.OpenFile("./csv/clean.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()

	// csvファイルを書き出し
	err := gocsv.MarshalFile(pokes, file)
	if err != nil {
		log.Fatal(err)
	}
}
