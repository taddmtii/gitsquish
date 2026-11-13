package main

import (
	"fmt"
	"flag"
)

func main() {
	// Get and parse commit message
	commitFlag := flag.String("m", "", "Commit Message")
	flag.Parse()
	fmt.Println("Commit message: ", *commitFlag)

	// Define command line arguments
	commands := []string {
		"git", "add", ".",
		"git", "commit", "-m", *commitFlag,
		"git", "push"
	}
}
