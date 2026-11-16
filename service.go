package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
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
func (s *Service) Stats() error {
	var wg sync.WaitGroup

	type CommitStat int

	const (
		TotalCommits CommitStat = iota
		LastMonthCommits
		Contributors
	)

	type Result struct {
		statType CommitStat
		output   string
		err      error
	}

	commandResults := make(chan Result, 3)

	wg.Add(3)

	go func() {
		defer wg.Done()

		var result Result
		args := []string{"rev-list", "--count", "HEAD"}
		output, err := s.RunCommand(args)
		if err != nil {
			result.err = fmt.Errorf("failed to get total commits: %w", err)
			commandResults <- result
			return
		}
		result.output = strings.TrimSpace(output)
		result.statType = TotalCommits
		commandResults <- result
	}()

	go func() {
		defer wg.Done()

		var result Result
		args := []string{"rev-list", "--count", "--since=30 days ago", "HEAD"}
		output, err := s.RunCommand(args)
		if err != nil {
			result.err = fmt.Errorf("failed to get commits in the last 30 days: %w", err)
			commandResults <- result
			return
		}

		result.output = strings.TrimSpace(output)
		result.statType = LastMonthCommits
		commandResults <- result
	}()

	go func() {
		defer wg.Done()

		var result Result
		args := []string{"shortlog", "-sne", "HEAD"}
		output, err := s.RunCommand(args)
		if err != nil {
			result.err = fmt.Errorf("failed to get contributors: %w", err)
			commandResults <- result
			return
		}

		var contributorTable strings.Builder
		table := tabwriter.NewWriter(&contributorTable, 0, 4, 3, ' ', 0)
		fmt.Fprintln(table, "COMMITS\tCONTRIBUTOR")

		for _, line := range strings.Split(strings.TrimSpace(output), "\n") {
			count, author, found := strings.Cut(strings.TrimSpace(line), "\t")
			if !found {
				continue
			}

			fmt.Fprintf(table, "%s\t%s\n", count, author)
		}

		err = table.Flush()
		if err != nil {
			result.err = fmt.Errorf("failed to write contributor table: %w", err)
			commandResults <- result
			return
		}

		result.output = contributorTable.String()
		result.statType = Contributors
		commandResults <- result
	}()

	wg.Wait()
	close(commandResults)

	outputs := make(map[CommitStat]string)

	for result := range commandResults {
		if result.err != nil {
			return result.err
		}

		outputs[result.statType] = result.output
	}

	output := fmt.Sprintf(
		"Total Commits: %s\nCommits in last 30 days: %s\n\n%s",
		outputs[TotalCommits],
		outputs[LastMonthCommits],
		outputs[Contributors],
	)

	_, err := fmt.Fprint(os.Stdout, output)
	if err != nil {
		return fmt.Errorf("failed to write stats: %w", err)
	}

	return nil
}

func (s *Service) RunCommand(args []string) (string, error) {
	cmd := exec.Command("git", args...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
