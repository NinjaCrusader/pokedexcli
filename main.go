package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//This is where we are defining our commands the user can use
	supportedCommands := map[string]cliCommand{
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
	}

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

		//Checking if the first value is a command within the supported commands
		cmd, ok := supportedCommands[firstValue]
		if ok {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
