package commands

import (
	"fmt"
)

type PokedexCommand struct {
	name        string
	description string
	execute     func(arg any) error
}

func NewPokedexCommand(config *Config) *PokedexCommand {
	return &PokedexCommand{
		name:        "pokedex",
		description: "Display the pokedex",
		execute: func(arg any) error {
			fmt.Println("Loading pokedex...")
			if len(config.Pokemons) == 0 {
				return fmt.Errorf("no pokemon in the pokedex")
			}
			for _, pokemon := range config.Pokemons {
				fmt.Printf(" - %s\n", pokemon.Name)
			}
			return nil
		},
	}
}

func (p *PokedexCommand) Name() string {
	return p.name
}

func (p *PokedexCommand) Description() string {
	return p.description
}

func (p *PokedexCommand) Execute(arg any) error {
	return p.execute(arg)
}
