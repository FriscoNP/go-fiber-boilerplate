# go-fiber-boilerplate

[![CI](https://github.com/FriscoNP/go-fiber-boilerplate/actions/workflows/ci.yml/badge.svg)](https://github.com/FriscoNP/go-fiber-boilerplate/actions/workflows/ci.yml)
[![CodeQL](https://github.com/FriscoNP/go-fiber-boilerplate/actions/workflows/codeql.yml/badge.svg)](https://github.com/FriscoNP/go-fiber-boilerplate/actions/workflows/codeql.yml)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/FriscoNP/go-fiber-boilerplate)](go.mod)

A small, opinionated Go service skeleton built on [Fiber v3](https://gofiber.io/)
with [Viper](https://github.com/spf13/viper) YAML configuration. It ships one
liveness endpoint and a package layout that can grow into a real API without
restructuring.

## Features

- Fiber v3 HTTP server
- YAML configuration via Viper
- Clear `cmd/` + `internal/` layout (config, server, middleware, routes, handler)
- Multi-stage Docker build and Compose file
- CI, CodeQL, Dependabot and tagged releases via GitHub Actions

## Quick start

Prerequisites: Go 1.25+ (and optionally Docker).

```bash
git clone git@github.com:FriscoNP/go-fiber-boilerplate.git
cd go-fiber-boilerplate

cp internal/config/example.config.local.yaml internal/config/config.local.yaml
make run

curl http://localhost:8080/api/health
```

## Configuration

Config is read from `internal/config/config.local.yaml` (git-ignored).

```yaml
app:
  name: "Golang Service with Fiber"
  port: 8080
  env: "local"
```

## Endpoints

| Method | Path          | Description    |
|--------|---------------|----------------|
| GET    | `/api/health` | Liveness check |

## Project layout

```
cmd/api/              application entry point
internal/config/      config struct and Viper loader
internal/server/      Fiber app construction and listen
internal/middleware/  middleware registration
internal/routes/      route registration
internal/handler/     HTTP handlers
docs/knowledge/       project knowledge bundle (OKF)
```

## Development

```bash
make run          # run the API
make test         # tests with race detector and coverage
make lint         # golangci-lint
make fmt          # gofmt
make build        # build binary into ./bin
```

## Docker

```bash
make docker-build
make docker-up
```

## Documentation

Architecture, components and operations notes live in
[`docs/knowledge/`](docs/knowledge/index.md).

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). Please report vulnerabilities as
described in [SECURITY.md](SECURITY.md).

## License

[MIT](LICENSE)
