package main

import (
	"bufio"
	"fmt"
	"os"
)

func handleCommand() {
	fmt.Fprint(os.Stdout, "$ ")

	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
	parsedCommand := command[:len(command)-1]

	switch parsedCommand {
	case "exit 0":
		os.Exit(0)
	default:
		fmt.Println(parsedCommand + ": command not found")

	}

}

func main() {

	for {
		handleCommand()
	}
}
