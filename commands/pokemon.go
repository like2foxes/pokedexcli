package commands

import (
	"fmt"
	"slices"
	"strings"

	"github.com/like2foxes/pokedexcli/commands/internal/fetch"
)

type Pokemon struct {
	Name string
	Height int
	Weight int
	Types []string
	Stats map[string]int
}

func NewPokemon(pokemonResult fetch.CatchResult) Pokemon {
	var types []string
	for _, t := range pokemonResult.Types {
		if slices.Contains(types, t.Type.Name) {
			continue
		}
		types = append(types, t.Type.Name)
	}

	var stats map[string]int = make(map[string]int)
	for _, s := range pokemonResult.Stats {
		stats[s.Stat.Name] = s.BaseStat
	}
	return Pokemon{
		Name: pokemonResult.Name,
		Height: pokemonResult.Height,
		Weight: pokemonResult.Weight,
		Types: types,
		Stats: stats,
	}
}

func (p Pokemon) ToString() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Name: %s\n", p.Name))
	sb.WriteString(fmt.Sprintf("Height: %d\n", p.Height))
	sb.WriteString(fmt.Sprintf("Weight: %d\n", p.Weight))
	sb.WriteString(fmt.Sprint("Stats:\n"))
	for k, v := range p.Stats {
		sb.WriteString(fmt.Sprintf(" - %s: %d\n", k, v))
	}
	sb.WriteString(fmt.Sprint("Types:\n"))
	for _, t := range p.Types {
		sb.WriteString(fmt.Sprintf(" - %s\n", t))
	}
	return sb.String()
}
