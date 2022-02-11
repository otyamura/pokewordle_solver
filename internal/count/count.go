package count

import (
	"sort"

	"github.com/otyamura/pokewordle_solver/internal/load"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func GetCharRanking() PairList {
	m := countChar()
	i := 0
	pl := make(PairList, len(m))
	for k, v := range m {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func countChar() map[string]int {
	pokes := load.LoadPokeRaws()
	m := make(map[string]int)
	for _, p := range pokes {
		for _, c := range p.Name {
			m[string(c)]++
		}
	}

	return m
}
