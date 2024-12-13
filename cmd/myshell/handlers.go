package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CommandFunc func(args []string)

var commandHandlers = map[string]CommandFunc{}

func handleCommand() {
	fmt.Fprint(os.Stdout, "$ ")

	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
	command = strings.TrimSpace(command)

	parts := strings.Fields(command)
	if len(parts) == 0 {
		return
	}

	cmd, args := parts[0], parts[1:]

	if handler, exists := commandHandlers[cmd]; exists {
		handler(args)
	} else {
		fmt.Println(cmd + ": command not found")
	}
}

func registerCommand(command string, handler CommandFunc) {
	commandHandlers[command] = handler
}

// Handlers
func handleExit(args []string) {
	os.Exit(0)
}

func handleEcho(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func handleType(args []string) {
	if len(args) != 1 {
		fmt.Println("Handle type needs only 1 argument")
		return
	}
	cmd := args[0]

	if _, exists := commandHandlers[cmd]; exists {
		fmt.Println(cmd + " is a shell builtin")
	} else {
		fmt.Println(cmd + ": not found")
	}

}