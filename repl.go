package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type pokedexLocation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type pokedexLocations struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []pokedexLocation `json:"results"`
}

func cleanInput(text string) []string {
	result := []string{}
	for word := range strings.SplitSeq(strings.ToLower(text), " ") {
		trimmed := strings.TrimSpace(word)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func commandExit(config *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n\n")
	fmt.Print("help: Displays a help message\n")
	fmt.Print("exit: Exit the Pokedex\n")
	return nil
}

func commandMap(config *config) error {
	req, err := http.Get(config.next)
	if err != nil {
		return err
	}
	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)
	pokedexLocations := pokedexLocations{}
	err = decoder.Decode(&pokedexLocations)
	if err != nil {
		return err
	}
	for _, location := range pokedexLocations.Results {
		fmt.Printf("%s-area\n", location.Name)
	}
	config.next = pokedexLocations.Next
	config.previous = pokedexLocations.Previous
	return nil
}

func commandMapB(config *config) error {
	if config.previous == "" {
		fmt.Print("you're on the first page\n")
		return nil
	}
	req, err := http.Get(config.previous)
	if err != nil {
		return err
	}
	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)
	pokedexLocations := pokedexLocations{}
	err = decoder.Decode(&pokedexLocations)
	if err != nil {
		return err
	}
	for _, location := range pokedexLocations.Results {
		fmt.Printf("%s-area\n", location.Name)
	}
	config.next = pokedexLocations.Next
	config.previous = pokedexLocations.Previous
	return nil
}
