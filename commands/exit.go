package commands

import (
	"fmt"
	"os"
)

type ExitCommand struct {
	name string
	description string
	execute func(arg any) error
}

func NewExitCommand(config *Config) *ExitCommand {
	return &ExitCommand{
		name: "exit",
		description: "Exit the program",
		execute: func(arg any) error {
			fmt.Println("Exiting...")
			os.Exit(0)
			return nil
		},
	}
}

func (c *ExitCommand) Name() string {
	return c.name
}

func (c *ExitCommand) Description() string {
	return c.description
}

func (c *ExitCommand) Execute(arg any) error {
	return c.execute(arg)
}


