# Contributing

Thanks for helping improve go-fiber-boilerplate!

## Setup

1. Fork and clone the repository.
2. `cp internal/config/example.config.local.yaml internal/config/config.local.yaml`
3. `make run` and check `GET /api/health`.

## Workflow

- Branch from `main`: `feature/<name>`, `fix/<name>` or `docs/<name>`.
- Use [Conventional Commits](https://www.conventionalcommits.org/)
  (`feat:`, `fix:`, `docs:`, `chore:` ...).
- Before opening a PR run `make fmt lint test`.
- Keep PRs focused and fill in the PR template.

## Reporting issues

Use the issue templates. For security problems see [SECURITY.md](SECURITY.md).
