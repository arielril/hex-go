# Go CLI - Template

[![Go Report Card](https://goreportcard.com/badge/github.com/arielril/golang-cli)](https://goreportcard.com/report/github.com/arielril/golang-cli)

This is a template repository to create CLI using golang. It creates a powerful cli using Cobra and Viper as main packages.

## Setup

Create your own `.config.yaml` file to configure the CLI.

To set up the git hooks run the following command.

```bash
cp ./scripts/pre-commit-hook.sh .git/hooks/pre-commit
chmod +x .git/hooks/pre-commit
```

## Running Locally

```bash
go run ./cmd/golang-cli/main.go help
```

## Dependency Management

This package uses [go modules](https://github.com/golang/go/wiki/Modules). Whenever you run `go build` or `go test` all dependencies will be downloaded automatically. See [go.mod](./go.mod) and [go.sum](go.sum) to see all modules used by this package.

## Configuration

### Runtime

Many settings are set throught the `.config.yaml` file and read using the [Viper](https://github.com/spf13/viper) package to set the configuration of the project and to retrieve the configuration values.

An example of the configuration file can be found at `configs/.config.example.yaml`.

## Project Structure

The project structure used follows the project layout defined in [Golang Standards - Project Layout](https://github.com/golang-standards/project-layout), which used a lot of definitions found in the web and compiled an unified layout.
