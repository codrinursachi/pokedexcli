package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/codrinursachi/pokedexcli/internal/pokecache"
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
	for word := range strings.FieldsSeq(strings.ToLower(text)) {
		trimmed := strings.TrimSpace(word)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func commandExit(config *config, cache *pokecache.Cache) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(config *config, cache *pokecache.Cache) error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n\n")
	fmt.Print("help: Displays a help message\n")
	fmt.Print("exit: Exit the Pokedex\n")
	return nil
}

func commandMap(config *config, cache *pokecache.Cache) error {
	resultsBytes, ok := cache.Get(config.next)
	if !ok {
		req, err := http.Get(config.next)
		if err != nil {
			return err
		}
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)
		locationsResp := pokedexLocations{}
		err = decoder.Decode(&locationsResp)
		if err != nil {
			return err
		}
		resultsBytes, err = json.Marshal(locationsResp)
		if err != nil {
			return err
		}
		cache.Add(config.next, resultsBytes)
	}
	var pokedexLocations pokedexLocations
	err := json.Unmarshal(resultsBytes, &pokedexLocations)
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

func commandMapB(config *config, cache *pokecache.Cache) error {
	if config.previous == "" {
		fmt.Print("you're on the first page\n")
		return nil
	}
	resultsBytes, ok := cache.Get(config.previous)
	if !ok {
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
		resultsBytes, err = json.Marshal(pokedexLocations)
		if err != nil {
			return err
		}
		cache.Add(config.previous, resultsBytes)
	}
	var locations pokedexLocations
	err := json.Unmarshal(resultsBytes, &locations)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Printf("%s-area\n", location.Name)
	}
	config.next = locations.Next
	config.previous = locations.Previous
	return nil
}
