package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/like2foxes/pokedexcli/commands"
)

func main() {
	commands := commands.CreateCommands()
	for {
		print("pokecli> ")
		var args []string
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		if err != nil {
			continue
		}
		words := strings.Split(input, " ")
		
		if len(words) == 0 {
			fmt.Println("Invalid command")
			continue
		}
		if len(words) == 1 {
			input = words[0]
		}
		if len(words) > 1 {
			input = words[0]
			for i := 1; i < len(words); i++ {
				args = append(args, words[i])
			}
		}

		command, ok := commands[input]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		err = command.Execute(strings.Join(args, " "))
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
