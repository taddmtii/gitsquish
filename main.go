package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	service := NewService()

	// Read arguments from the command line
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
		service.ship(*message)
		fmt.Println("Code shipped!")
		os.Exit(0)
	}
	

	
}
