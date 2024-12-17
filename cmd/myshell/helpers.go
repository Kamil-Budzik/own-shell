package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

func parseCommand(input string) []string {
	var args []string
	var current strings.Builder
	inSingleQuote := false

	for _, char := range input {
		switch char {
		case '\'':
			inSingleQuote = !inSingleQuote
		case ' ':
			if inSingleQuote {
				current.WriteRune(char)
			} else if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
		default:
			current.WriteRune(char)
		}
	}

	if current.Len() > 0 {
		args = append(args, current.String())
	}

	if inSingleQuote {
		fmt.Fprintln(os.Stderr, "Error: unmatched single quote")
		return nil
	}

	return args
}
