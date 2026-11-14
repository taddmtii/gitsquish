# gitsquish

A small command-line tool written in Go that simplifies common Git workflows.

## Usage

Run from inside a Git repository:

```sh
gitsquish ship -m "Your commit message"
```

The `ship` command runs these steps in order:

1. Stage changes with `git add .`.
2. Commit with the provided message.
3. Push with `git push`.

Use it when you want to stage, commit, and push in a single command. Git must be installed, and your branch must be configured for pushing to a remote.

## Install locally

With Go installed, run this from the project directory:

```sh
go install .
```

Make sure your Go binary directory (`GOBIN`, or `$(go env GOPATH)/bin` by default) is on your `PATH`.
