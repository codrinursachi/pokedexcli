package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     string
	previous string
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
		"map": {
			name:        "map",
			description: "Show map of pokedex locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show map of pokedex locations (back)",
			callback:    commandMapB,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{
		next:     "https://pokeapi.co/api/v2/location/",
		previous: "",
	}
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		user_command := cleanInput(text)
		if cmd, exists := commands[user_command[0]]; exists {
			if err := cmd.callback(cfg); err != nil {
				fmt.Printf("%v\n", err)
			}
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}
