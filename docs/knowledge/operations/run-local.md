---
type: Playbook
title: Run locally
description: Steps to go from a clean checkout to a running service on port 8080.
tags: [operations, run, setup]
status: stable
generated: { by: human:frisco, at: 2026-10-11T00:00:00Z }
verified: { by: human:frisco, at: 2026-10-11T00:00:00Z }
---

# Prerequisites

- Go **1.25.3** or newer (the `go.mod` directive is `go 1.25.3`; a newer
  toolchain such as Go 1.27 compiles it without changes).
- Network access on first build to download module dependencies.
- The repository must be your **current working directory** - the config loader
  searches `./internal/config` relative to it.

# Steps

1. **Create the config file.** It is gitignored, so a fresh clone does not
   have one:

   ```bash
   cp internal/config/example.config.local.yaml internal/config/config.local.yaml
   ```

2. **Run from the repository root:**

   ```bash
   make run
   ```

   This is equivalent to `go run cmd/api/main.go`.

3. **Verify:**

   ```bash
   curl -s http://localhost:8080/api/health
   # {"message":"Fiber API is running smoothly!","status":"ok"}
   ```

4. **Stop** with `Ctrl+C`. There is no graceful-shutdown handler.

# Useful variants

```bash
go build ./...           # type-check and compile without running
go vet ./...             # static analysis
go run cmd/api/main.go   # what `make run` executes
```

# Troubleshooting

| Symptom | Cause | Fix |
|---------|-------|-----|
| `Error reading config file: ...` then exit | Missing `config.local.yaml`, or the working directory is not the repo root | Copy the example file and run from the repo root |
| `address already in use` | Port 8080 is taken | Free the port or change `app.port` in the config file |
| Module download errors | Offline or blocked proxy | Configure `GOPROXY` and retry |

# Related concepts

* [Configuration](/operations/configuration.md) for the config file details.
* [Health check](/api/health.md) for the verification endpoint.
* [Architecture](/architecture.md) for what happens during startup.
