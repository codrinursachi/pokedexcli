package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/codrinursachi/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache, string, pokedex) error
}

type config struct {
	next     string
	previous string
}

type pokedex struct {
	pokemons map[string]Pokemon
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
		"explore": {
			name:        "explore",
			description: "Explore a specific location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a specific Pokemon",
			callback:    commandCatch,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{
		next:     "https://pokeapi.co/api/v2/location-area/",
		previous: "",
	}
	cache := pokecache.NewCache(5 * time.Second)
	pokedex := pokedex{
		pokemons: make(map[string]Pokemon),
	}
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		user_command := cleanInput(text)
		if cmd, exists := commands[user_command[0]]; exists {
			if len(user_command) < 2 {
				user_command = append(user_command, "")
			}
			if err := cmd.callback(cfg, cache, user_command[1], pokedex); err != nil {
				fmt.Printf("%v\n", err)
			}
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}
