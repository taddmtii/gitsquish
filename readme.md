# gitsquish

A small command-line tool written in Go that simplifies common Git workflows and summarizes repository stats.

## Usage

Run from inside a Git repository:

### Ship changes

```sh
gitsquish ship -m "Your commit message"
```

The `ship` command runs these steps in order:

1. Stage changes with `git add .`.
2. Commit with the provided message.
3. Push with `git push`.

Use it when you want to stage, commit, and push in a single command. Git must be installed, and your branch must be configured for pushing to a remote.

### Repository stats

```sh
gitsquish stats
```

Displays:

- Total commit count.
- Commit count for the last 30 days.
- A contributor table with commit counts, names, and email addresses, sorted by commit count.

Stats cover history reachable from the current commit (`HEAD`), including merged branches. The repository must have at least one commit.

## Install locally

With Go installed, run this from the project directory:

```sh
go install .
```

Make sure your Go binary directory (`GOBIN`, or `$(go env GOPATH)/bin` by default) is on your `PATH`.
