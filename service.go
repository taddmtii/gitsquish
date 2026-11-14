package main

import (
	"fmt"
	"os/exec"
)

type Service struct {}

func NewService() *Service { return &Service{} }

// Usage: gitsquish ship -m "commit message"
func (s *Service) ship(message string) {

	stageCmd := exec.Command("git", "add", ".")
	err := stageCmd.Run()
	if err != nil {
		fmt.Println("Something went wrong with staging:", err)
		return
	}

	commitCmd := exec.Command("git", "commit", "-m", message) 
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
