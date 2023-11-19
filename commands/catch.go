package commands

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/like2foxes/pokedexcli/commands/internal/fetch"
)

type CatchCommand struct {
	name        string
	description string
	execute     func(arg any) error
}

func NewCatchCommand(config *Config) *CatchCommand {
	return &CatchCommand{
		name:        "catch",
		description: "Catch a pokemon",
		execute: func(arg any) error {
			pokemonName := strings.ToLower(strings.Trim(arg.(string), " "))
			if arg == nil || pokemonName == "" {
				return fmt.Errorf("Invalid pokemon name")
			}
			var catchResult fetch.CatchResult
			var raw []byte
			var err error
			if cached, ok := config.Cache.Get(pokemonName); ok {
				catchResult, err = fetch.ParseResult(catchResult, cached)
			} else {
				url := config.pokemonUrl + pokemonName
				catchResult, raw, err = fetch.GetData(url, catchResult)
				config.Cache.Add(pokemonName, raw)
			}
			if err != nil {
				return err
			}
			fmt.Printf(" - %v\n", catchResult.Name)
			fmt.Printf("Throwing a Pokeball on %v...\n", catchResult.Name)
			if(fight(catchResult.BaseExperience)) {
				fmt.Printf(" - %v was caught!\n", catchResult.Name)
				if _, ok := config.Pokemons[pokemonName]; !ok {
					config.Pokemons[pokemonName] = NewPokemon(catchResult)
				}
			} else {
				fmt.Printf(" - %v ran away!\n", catchResult.Name)
			}

			return nil
		},
	}
}

func fight(exp int) bool {
	num := rand.Intn(400)
	return exp < num
}

func (c *CatchCommand) Name() string {
	return c.name
}

func (c *CatchCommand) Description() string {
	return c.description
}

func (c *CatchCommand) Execute(arg any) error {
	return c.execute(arg)
}
