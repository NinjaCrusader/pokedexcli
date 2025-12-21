package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/NinjaCrusader/pokedexcli/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
	mapcallback func() internal.Maps
}

// Cleans the User input before we use it
// Lowercases the input
// Trims any extra whitespace
// Breaks the string into a splice by word
func cleanInput(text string) []string {
	newText := strings.ToLower(text)
	newText = strings.TrimSpace(newText)
	result := strings.Fields(newText)
	return result
}

// Will exit the application when called
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// Will print out the commands available to the User
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

// Will reach out to the PokeAPI and return 20 location-area points and display them to the User
func commandMap() (internal.Maps, error) {
	mapData, err := internal.GetMapHelper()
	if err != nil {
		fmt.Println("Getting map information has failed")
		os.Exit(0)
	}
	return nil
}
