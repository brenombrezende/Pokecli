package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
}

func StartRepl(cfg *config) {
	commandClear(cfg)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		commandName := cleanInput(scanner.Text())
		if len(commandName) == 0 {
			continue
		}
		command, exists := getCommands()[commandName[0]]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Print("ERROR - ")
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Command not found")
			continue
		}

	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"clear": {
			name:        "clear",
			description: "Clear the terminal",
			callback:    commandClear,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations, and so on.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Same as map - but displays the names of the previous 20 areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}
