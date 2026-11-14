package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/tabwriter"
)

type Service struct{}

func NewService() *Service { return &Service{} }

// Usage: gitsquish ship -m "commit message"
func (s *Service) Ship(message string) error {
	args := []string{"add", "."}
	_, err := s.RunCommand(args)
	if err != nil {
		return fmt.Errorf("failed to stage changes: %w", err)
	}
	fmt.Println("Changes staged.")

	args = []string{"commit", "-m", message}
	_, err = s.RunCommand(args)
	if err != nil {
		return fmt.Errorf("failed to create commit: %w", err)
	}
	fmt.Println("Commit created.")

	args = []string{"push"}
	_, err = s.RunCommand(args)
	if err != nil {
		return fmt.Errorf("commit created locally, but push failed: %w", err)
	}
	fmt.Println("Changes pushed.")
	return nil
}

// Usage: gitsquish stats
func (s *Service) Stats() {
	// git rev-list --count HEAD
	args := []string{"rev-list", "--count", "HEAD"}
	output, err := s.RunCommand(args)
	if err != nil {
		fmt.Println("Something went wrong with getting total commits:", err)
		return
	}
	fmt.Println("Total Commits: ", strings.TrimSpace(output))

	// git rev-list --count --since="30 days ago" HEAD
	args = []string{"rev-list", "--count", "--since=\"30 days ago\"", "HEAD"}
	output, err = s.RunCommand(args)
	if err != nil {
		fmt.Println("Something went wrong with getting commits in last 30 days:", err)
		return
	}
	fmt.Println("Commits in last 30 days: ", strings.TrimSpace(output)+"\n")

	// git shortlog -sne HEAD
	args = []string{"shortlog", "-sne", "HEAD"}
	output, err = s.RunCommand(args)
	if err != nil {
		fmt.Println("Something went wrong with getting shortlog:", err)
		return
	}
	// Create a table to format output correctly
	table := tabwriter.NewWriter(os.Stdout, 0, 4, 3, ' ', 0)

	fmt.Fprintln(table, "COMMITS\tCONTRIBUTOR")

	for _, line := range strings.Split(strings.TrimSpace(output), "\n") {
		count, author, found := strings.Cut(strings.TrimSpace(line), "\t")
		if !found {
			continue
		}

		fmt.Fprintf(table, "%s\t%s\n", count, author)
	}
	table.Flush()

}

func (s *Service) RunCommand(args []string) (string, error) {
	cmd := exec.Command("git", args...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
