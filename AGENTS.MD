# Agent Instructions

## Project

This is a Go project developed collaboratively by Roberto and Codex.

## Go rules

- Keep changes small and focused.
- Use idiomatic Go.
- Run `gofmt` on changed Go files.
- Run `go test ./...` before saying a task is complete.
- Run `go vet ./...` when behavior or API code changes.
- Prefer simple structs and clear interfaces over over-engineering.
- Do not add dependencies unless needed.
- Explain why any new dependency is necessary.

## Git workflow

- Work on feature branches.
- Do not push directly to `main`.
- Open pull requests for changes.
- Summarize files changed, tests run, and remaining risks.

## Safety

- Do not commit secrets, tokens, passwords, local machine paths, or private credentials.
