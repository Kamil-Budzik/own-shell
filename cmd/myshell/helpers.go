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
	inDoubleQuote := false
	escapeNext := false

	for _, char := range input {
		switch {
		case escapeNext:
			current.WriteRune(char)
			escapeNext = false

		case char == '\\' && inDoubleQuote:
			escapeNext = true

		case char == '\'' && !inDoubleQuote:
			inSingleQuote = !inSingleQuote

		case char == '"' && !inSingleQuote:
			inDoubleQuote = !inDoubleQuote

		case char == ' ' && !inSingleQuote && !inDoubleQuote:
			if current.Len() > 0 {
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

	if inSingleQuote || inDoubleQuote {
		fmt.Fprintln(os.Stderr, "Error: unmatched quote")
		return nil
	}

	return args
}
