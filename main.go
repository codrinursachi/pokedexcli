package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		commands := cleanInput(text)
		fmt.Printf("Your command was: %v\n",commands[0])
	}
}
