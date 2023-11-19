package commands

import (
	"fmt"
	"github.com/like2foxes/pokedexcli/commands/internal/fetch"
)

type MapCommand struct {
	name        string
	description string
	execute     func(arg any) error
}

func NewMapCommand(config *Config) *MapCommand {
	return &MapCommand{
		name:        "map",
		description: "Display the map",
		execute: func(arg any) error {
			fmt.Println("Loading map...")
			var mapResult fetch.MapResult
			var raw []byte
			var err error
			if cached, ok := config.Cache.Get(config.next); ok {
				mapResult, err = fetch.ParseResult(mapResult, cached)
			} else {
				mapResult, raw, err = fetch.GetData(config.next, mapResult)
				config.Cache.Add(config.next, raw)
			}

			if err != nil {
				return err
			}
			config.next = mapResult.Next
			config.prev = mapResult.Previous
			for _, location := range mapResult.Results {
				fmt.Printf("%v\n", location.Name)
			}
			return nil
		},
	}
}

func (c *MapCommand) Name() string {
	return c.name
}

func (c *MapCommand) Description() string {
	return c.description
}

func (c *MapCommand) Execute(arg any) error {
	return c.execute(arg)
}
