package commands

import (
	"fmt"
)

type HelpCommand struct {
	name string
	description string
	execute func(arg any) error
}

func NewHelpCommand(config *Config) *HelpCommand {
	return &HelpCommand{
		name: "help",
		description: "Display help information",
		execute: func(arg any) error {
			fmt.Println("Welcome to the Pokemon CLI!")
			fmt.Println("This CLI will allow you to explore the Pokemon world.")
			fmt.Println("To get started, type 'map' to see the map.")
			fmt.Println("To exit, type 'exit'.")
			fmt.Println("Commands:")
			for _, command := range config.Commands {
				fmt.Printf("%s - %s\n", command.Name(), command.Description())
			}
			return nil
		},
	}
}

func (c *HelpCommand) Name() string {
	return c.name
}

func (c *HelpCommand) Description() string {
	return c.description
}

func (c *HelpCommand) Execute(arg any) error {
	return c.execute(arg)
}
