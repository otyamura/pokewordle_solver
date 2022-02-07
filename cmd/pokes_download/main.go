package main

import (
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/gocarina/gocsv"
	pre "github.com/otyamura/pokewordle_solver/internal/preprocessing"
)

// const MAX_POKES = 898

const MAX_POKES = 10

func main() {
	pokes := []pre.Pokemon{}
	for i := 1; i <= MAX_POKES; i++ {
		p := pre.Get_poke(i)
		if utf8.RuneCountInString(p.Name) == 5 {
			fmt.Println(p.Name)
			pokes = append(pokes, p)
		}
	}

	file, _ := os.OpenFile("./csv/poke.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()

	// csvファイルを書き出し
	gocsv.MarshalFile(pokes, file)
}