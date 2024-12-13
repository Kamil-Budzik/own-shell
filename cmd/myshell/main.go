package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

type CommandFunc func(args []string)

var commandHandlers = map[string]CommandFunc{}

func main() {
	registerCommand("exit", handleExit)
	registerCommand("echo", handleEcho)

	for {
		handleCommand()
	}
}

func registerCommand(command string, handler CommandFunc) {
	commandHandlers[command] = handler
}
