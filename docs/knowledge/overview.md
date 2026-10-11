---
type: Reference
title: Overview
description: What this boilerplate is and how the repository is organized.
tags: [overview, architecture, go, fiber]
status: stable
generated: { by: human:frisco, at: 2026-10-11T00:00:00Z }
verified: { by: human:frisco, at: 2026-10-11T00:00:00Z }
sources:
  - id: fiber-docs
    resource: https://gofiber.io/
    title: Fiber project home
    author: team:gofiber
    last_modified: 2026-01-01T00:00:00Z
  - id: viper-docs
    resource: https://github.com/spf13/viper
    title: spf13/viper repository
    author: human:spf13
    last_modified: 2026-06-01T00:00:00Z
---

# Summary

`go-fiber-boilerplate` is a deliberately small Go service skeleton built
on Fiber v3[^fiber-docs] with viper[^viper-docs] configuration
loading. It ships a single liveness endpoint and the package layout
needed to grow a real API without restructuring.

The module path is `github.com/FriscoNP/go-fiber-boilerplate` and the
Go toolchain directive is `go 1.25.3`.

# Repository layout

| Path | Purpose |
|------|---------|
| `cmd/api/` | Application entry point. |
| `internal/config/` | Config struct plus the viper loader. |
| `internal/server/` | Fiber app construction and `Listen`. |
| `internal/middleware/` | Middleware registration. |
| `internal/routes/` | Route registration. |
| `internal/handler/` | HTTP handlers. |
| `Makefile` | `make run` shortcut. |
| `Dockerfile`, `docker-compose.yml` | Empty container placeholders. |
| `go.mod`, `go.sum` | Module definition and checksums. |

# Dependencies

Direct dependencies are `github.com/gofiber/fiber/v3` (v3.0.0) and
`github.com/spf13/viper` (v1.21.0). Everything else in `go.mod` is
indirect.

# What it does not ship

* No database, cache, or queue integration.
* No authentication, authorization, or session handling.
* No structured logging beyond Fiber's request logger.
* No test files yet.
* No container runtime configuration; both Docker files are empty.

# Related concepts

* [Architecture](/architecture.md) for the request lifecycle.
* [Configuration](/operations/configuration.md) for required config
  files.
* [Health check](/api/health.md) for the only implemented endpoint.

[^fiber-docs]: Fiber project home
[^viper-docs]: spf13/viper repository
