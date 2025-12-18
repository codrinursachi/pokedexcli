package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show available commands",
			callback:    commandHelp,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		user_command := cleanInput(text)
		if cmd, exists := commands[user_command[0]]; exists {
			if err := cmd.callback(); err != nil {
				fmt.Printf("%v\n", err)
			}
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}
