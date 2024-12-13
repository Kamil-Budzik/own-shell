package main

import (
	"fmt"
	"os"
	"strings"
)

func handleExit(args []string) {
	os.Exit(0)
}

func handleEcho(args []string) {
	fmt.Println(strings.Join(args, " "))
}
