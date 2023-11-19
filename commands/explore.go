package commands

import (
	"fmt"
	"strings"
	"github.com/like2foxes/pokedexcli/commands/internal/fetch"
)

type ExploreCommand struct {
	name        string
	description string
	execute     func(arg any) error
}

func NewExploreCommand(config *Config) *ExploreCommand {
	return &ExploreCommand{
		name:        "explore",
		description: "Explore the map",
		execute: func(arg any) error {
			location := strings.ToLower(strings.Trim(arg.(string), " "))
			if arg == nil || location == "" {
				return fmt.Errorf("Invalid location")
			}
			fmt.Println("Exploring map...")
			var exploreResult fetch.ExploreResult
			var raw []byte
			var err error
			if cached, ok := config.Cache.Get(location); ok {
				exploreResult, err = fetch.ParseResult(exploreResult, cached)
			} else {
				url := config.basicUrl + location
				exploreResult, raw, err = fetch.GetData(url, exploreResult)
				config.Cache.Add(location, raw)
			}
			if err != nil {
				return err
			}
			fmt.Printf("Found Pokemon: \n")
			for _, pokemonEncouter := range exploreResult.PokemonEncounters {
				fmt.Printf(" - %v\n", pokemonEncouter.Pokemon.Name)
			}
			return nil
		},
	}
}

func (c *ExploreCommand) Name() string {
	return c.name
}

func (c *ExploreCommand) Description() string {
	return c.description
}

func (c *ExploreCommand) Execute(arg any) error {
	return c.execute(arg)
}
