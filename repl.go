package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearCmdLine() {
	clearCommand := ""
	if runtime.GOOS == "windows" {
		clearCommand = "cls"
	} else {
		clearCommand = "clear"
	}

	cmd := exec.Command(clearCommand)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		switch scanner.Text() {
		case "help":
			commandHelp()
		case "exit":
			return
		default:
			fmt.Println("Unknown command, please use a valid prompt")
		}
	}
}

func commandHelp() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
}
