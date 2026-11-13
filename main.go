package main

import (
	"fmt"
	"flag"
	"os/exec"
	"os"
)

func main() {
	// Get and parse commit message
	commitFlag := flag.String("m", "", "Commit Message")
	flag.Parse()
	if *commitFlag == "" {
		fmt.Println("Error: commit message cannot be empty")
		os.Exit(1)
	}
	
	stageCmd := exec.Command("git", "add", ".")
	err := stageCmd.Run()
	if err != nil {
		fmt.Println("Something went wrong with staging:", err)
		return
	}

	commitCmd := exec.Command("git", "commit", "-m", *commitFlag)

	commitCmd.Stdout = os.Stdout
	commitCmd.Stderr = os.Stderr

	err = commitCmd.Run()
	if err != nil {
		fmt.Println("Something went wrong with committing:", err)
		return
	}

	pushCmd := exec.Command("git", "push")
	err = pushCmd.Run()
	if err != nil {
		fmt.Println("Something went wrong with committing:", err)
		return
	}

}
