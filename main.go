package main

import (
	"fmt"
	"os"
)

func main() {
	command := os.Args[1]
	switch command {
	case "input":
		input(os.Args[2:])
	case "leaderboard":
		leaderboard(os.Args[2:])
	default:
		_, _ = fmt.Fprintf(os.Stderr, "Invalid command: %s\n", command)
		os.Exit(1)
	}
}

func input(args []string) {
	fmt.Println("input")
}

func leaderboard(args []string) {
	fmt.Println("leaderboard")
}
