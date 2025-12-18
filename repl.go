package main

import (
	"fmt"
	"os"
	"strings"
)

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

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n\n")
	fmt.Print("help: Displays a help message\n")
	fmt.Print("exit: Exit the Pokedex\n")
	return nil
}
