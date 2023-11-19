package commands

import (
	"time"
	"github.com/like2foxes/pokedexcli/commands/internal/cache"
)

const areaUrl = "https://pokeapi.co/api/v2/location-area/"
const pokemonUrl = "https://pokeapi.co/api/v2/pokemon/"
const firstUrl = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
type Command interface {
	Name() string
	Description() string
	Execute(arg any) error
}

type Config struct {
	Commands map[string]Command 
	Pokemons map[string]Pokemon
	Cache cache.Cache
	basicUrl string
	pokemonUrl string
	next string
	prev string
}

func CreateCommands() map[string]Command {
	cache := cache.NewCache(time.Second * 5)
	config := &Config{
		Pokemons: map[string]Pokemon{},
		basicUrl: areaUrl,
		pokemonUrl: pokemonUrl,
		next: firstUrl,
		prev: "",
		Cache: *cache,
	}
	var commands = map[string]Command{
		"map": NewMapCommand(config),
		"mapb": NewMapbCommand(config),
		"explore": NewExploreCommand(config),
		"catch": NewCatchCommand(config),
		"inspect": NewInspectCommand(config),
		"pokedex": NewPokedexCommand(config),
		"exit": NewExitCommand(config),
		"help": NewHelpCommand(config),
	}

	config.Commands = commands

	return config.Commands
}
