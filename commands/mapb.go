package commands
import (
	"errors"
	"fmt"
	"github.com/like2foxes/pokedexcli/commands/internal/fetch"
)

type MapbCommand struct {
	name        string
	description string
	execute     func(arg any) error
}

func NewMapbCommand(config *Config) *MapbCommand {
	return &MapbCommand{
		name:        "mapb",
		description: "Display the previous map",
		execute: func(arg any) error {
			if config.prev == "" {
				return errors.New("No previous map")
			}
			fmt.Println("Loading map...")
			var mapResult fetch.MapResult
			var raw []byte
			var err error
			if cached, ok := config.Cache.Get(config.prev); ok {
				mapResult, err = fetch.ParseResult(mapResult, cached)
			} else {
				mapResult, raw, err = fetch.GetData(config.prev, mapResult)
				config.Cache.Add(config.prev, raw)
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

func (c *MapbCommand) Name() string {
	return c.name
}

func (c *MapbCommand) Description() string {
	return c.description
}

func (c *MapbCommand) Execute(arg any) error {
	return c.execute(arg)
}
