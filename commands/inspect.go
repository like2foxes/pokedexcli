package commands

import (
	"fmt"
	"strings"
)

type Inspect struct {
	name string
	description string
	execute func(arg any) error
}

func NewInspectCommand(config *Config) *Inspect {
	return &Inspect{
		name: "inspect",
		description: "Inspect a pokemon",
		execute: func(arg any) error {
			pokemonName := strings.ToLower(strings.Trim(arg.(string), ""))
			if arg == nil || pokemonName == "" {
				return fmt.Errorf("no pokemon name provided")
			}
			if pokemon, ok := config.Pokemons[pokemonName]; ok {
				fmt.Println(pokemon.ToString())
				return nil
			}
			return fmt.Errorf("this pokemon was not found")
		},
	}
}

func (i *Inspect) Name() string {
	return i.name
}

func (i *Inspect) Description() string {
	return i.description
}

func (i *Inspect) Execute(arg any) error {
	return i.execute(arg)
}
