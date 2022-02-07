package preprocessing

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/otyamura/pokewordle_solver/types"
	"github.com/tidwall/gjson"
)

func Get_poke(id int) types.Pokemon {
	base_url := "https://pokeapi.co/api/v2/pokemon-species/"
	resp, err := http.Get(base_url + strconv.Itoa(id))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json := string(body)
	p := types.Pokemon{
		Name:       gjson.Get(json, "names.0.name").String(),
		Generation: gjson.Get(json, "generation.name").String(),
	}
	return p
}
