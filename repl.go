package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/NinjaCrusader/pokedexcli/internal"
	"github.com/NinjaCrusader/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	Next     *string
	Previous *string
	Cache    *pokecache.Cache
}

// Will return available commands that can be used
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commanMapB,
		},
		"explore": {
			name:        "explore",
			description: "Displays the names of the Pokemon that can be found in an area",
			callback:    commandExplore,
		},
	}
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
func commandExit(cfg *config, area string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// Will print out the commands available to the User
func commandHelp(cfg *config, area string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Displays the names of 20 location areas in the Pokemon world")
	fmt.Println("explore: It takes the name of a location area and prints the Pokemon in that area (example: explore <area_name>)")
	return nil
}

func commandExplore(cfg *config, area string) error {
	if len(area) == 0 {
		err := fmt.Errorf("Please provide an area")
		return err
	}

	fmt.Printf("Exploring %v...\n", area)

	areaData, err := internal.GetAreaInformationHelper(area, cfg.Cache)
	if err != nil {
		fmt.Println("Getting area information has failed")
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, r := range areaData.PokemonEncounters {
		fmt.Printf(" - %v\n", r.Pokemon.Name)
	}

	return nil
}

// Will reach out to the PokeAPI and return 20 location-area points and display them to the User
func commandMap(cfg *config, area string) error {
	var url string

	if cfg.Next == nil {
		url = ""
	} else {
		url = *cfg.Next
	}
	mapData, err := internal.GetMapHelper(url, cfg.Cache)
	if err != nil {
		fmt.Println("Getting map information has failed")
		return err
	}

	for _, r := range mapData.Results {
		fmt.Println(r.Name)
	}

	cfg.Next = mapData.Next
	cfg.Previous = mapData.Previous

	return nil
}

func commanMapB(cfg *config, area string) error {
	var url string

	if cfg.Previous == nil {
		url = ""
	} else {
		url = *cfg.Previous
	}

	mapData, err := internal.GetMapHelper(url, cfg.Cache)
	if err != nil {
		fmt.Println("getting map information has failed")
		return err
	}

	for _, r := range mapData.Results {
		fmt.Println(r.Name)
	}

	cfg.Next = mapData.Next
	cfg.Previous = mapData.Previous

	return nil
}

func startRepl(cfg *config) {

	//Waiting for user input
	scanner := bufio.NewScanner(os.Stdin)

	//Starting the REPL loop
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanUserInput := cleanInput(userInput)
		if len(cleanUserInput) == 0 {
			fmt.Println("Please enter a valid command")
			continue
		}

		//Checking the first value within the user input
		firstValue := cleanUserInput[0]
		var argumentValue string
		if len(cleanUserInput) > 1 {
			argumentValue = cleanUserInput[1]
		}

		//Checking if the first value is a command within the supported commands
		commands := getCommands()
		cmd, ok := commands[firstValue]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(cfg, argumentValue); err != nil {
			fmt.Println(err)
		}
	}
}
