package main

import (
	"fmt"
	"os/exec"
)

type Service struct{}

func NewService() *Service { return &Service{} }

// Usage: gitsquish ship -m "commit message"
func (s *Service) Ship(message string) {
	args := []string{"add", "."}
	err := s.RunCommand(args)
	if err != nil {
		fmt.Println("Something went wrong with git add command.")
		return
	}

	args = []string{"commit", "-m", message}
	err = s.RunCommand(args)
	if err != nil {
		fmt.Println("Something went wrong with git commit command.")
		return
	}


	args = []string{"push"}
	err = s.RunCommand(args)
	if err != nil {
		fmt.Println("Something went wrong with git push command.")
		return
	}

}

// Usage: gitsquish stats
func (s *Service) Stats() {
	args := []string{"rev-list", "--count", "HEAD"}
	err := s.RunCommand(args); 
	if err != nil {
		fmt.Println("Something went wrong with getting stats:", err)
	}
}

func (s *Service) RunCommand(args []string) error {
	cmd := exec.Command("git", args...)
	return cmd.Run()
}
