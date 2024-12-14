package main

import (
	"fmt"
	"os"
	"os/exec"
)

func executeFile(cmd string, args []string) {
	command := exec.Command(cmd, args...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	err := command.Run()
	if err != nil {
		fmt.Printf("%s: command not found\n", cmd)
	}
}
