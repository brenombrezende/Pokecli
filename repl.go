package main

import (
	"bufio"
	"fmt"
	"os"
)

func StartRepl() {
	commandClear()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		commandName := scanner.Text()
		command, exists := getCommands()[commandName]
		if exists {
			command.callback()
			continue
		} else {
			fmt.Println("Command not found")
			continue
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func()
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
	}
}
