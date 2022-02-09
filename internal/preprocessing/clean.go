package preprocessing

import (
	"github.com/otyamura/pokewordle_solver/internal/load"
	"github.com/otyamura/pokewordle_solver/types"
)

func Cleaning() []types.Pokemon {
	raws := load.Load_pokes()
	var pokes []types.Pokemon
	var pre = "generation-"
	for _, raw := range raws {
		var gen int
		switch raw.Generation {
		case pre + "i":
			gen = 1
		case pre + "ii":
			gen = 2
		case pre + "iii":
			gen = 3
		case pre + "iv":
			gen = 4
		case pre + "v":
			gen = 5
		case pre + "vi":
			gen = 6
		case pre + "vii":
			gen = 7
		case pre + "viii":
			gen = 8
		default:
			gen = 0
		}
		poke := types.Pokemon{
			Name:       raw.Name,
			Generation: gen,
		}
		pokes = append(pokes, poke)
	}
	return pokes
}
