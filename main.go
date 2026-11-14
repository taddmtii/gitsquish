package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	service := NewService()

	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Invalid command.")
		os.Exit(1)
	}

	command := args[1]
	switch command {
	case "ship":
		shipFlags := flag.NewFlagSet("ship", flag.ExitOnError)
		message := shipFlags.String("m", "", "Commit Message")
		shipFlags.Parse(args[2:])
		if *message == "" {
			fmt.Println("Error: commit message cannot be empty")
			os.Exit(1)
		}
		err := service.Ship(*message)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		os.Exit(0)
	case "stats":
		service.Stats()
		os.Exit(0)
	}
}
