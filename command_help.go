package main

import "fmt"

func commandHelp(cfg *config) error {
	allCommands := getCommands()
	fmt.Println("----------")
	fmt.Println("Here's a list of all available commands:")
	for _, c := range allCommands {
		fmt.Println(c.name, " - ", c.description)
	}
	fmt.Println("----------")
	return nil
}
