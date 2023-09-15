package main

import (
	"os"
	"os/exec"
	"runtime"
)

func commandClear() error {
	clearCommand := ""
	if runtime.GOOS == "windows" {
		clearCommand = "cls"
	} else {
		clearCommand = "clear"
	}

	cmd := exec.Command(clearCommand)
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}
