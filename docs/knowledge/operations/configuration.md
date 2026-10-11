---
type: Reference
title: Configuration
description: Location, schema, and failure modes of the YAML configuration file loaded by viper.
tags: [operations, config, yaml, viper]
status: stable
generated: { by: human:frisco, at: 2026-10-11T00:00:00Z }
verified: { by: human:frisco, at: 2026-10-11T00:00:00Z }
---

# File location

| Item | Value |
|------|-------|
| Search directory | `./internal/config` (relative to the process working directory) |
| Filename stem | `config.local` |
| Full path | `internal/config/config.local.yaml` |
| Format | YAML, forced via `viper.SetConfigType("yaml")` |
| Template | `internal/config/example.config.local.yaml` |
| VCS status | `.gitignore` contains `config.local.yaml`, so it is never committed |

viper resolves the search path relative to where the binary is executed, not
relative to the source file. Run from the repository root.

# Schema

```yaml
app:
  name: "Golang Service with Fiber"
  port: 8080
  env: "local"
```

| Key | Type | Consumed by | Notes |
|-----|------|-------------|-------|
| `app.name` | string | nothing yet | Label only; retained for future use. |
| `app.port` | int | `cmd/api/main.go` -> `strconv.Itoa` -> `server.Start` | Determines the listen port. |
| `app.env` | string | nothing yet | Environment label; retained for future use. |

# Setup

```bash
cp internal/config/example.config.local.yaml internal/config/config.local.yaml
```

# Failure modes

- **File missing** -> `log.Fatalf("Error reading config file: ...")`; the
  process exits with a non-zero status before the server starts.
- **File malformed** -> same fatal path.
- **Wrong working directory** -> the relative search path misses the file,
  producing the same "config file not found" exit.
- **Unknown keys** -> ignored; viper does not enforce strict key matching.
- **Missing keys** -> `app.port` defaults to `0` under weak decoding, which
  binds an ephemeral port rather than failing.

# Not supported today

- Environment variable overrides (`APP_PORT` and similar are not read; there is
  no `AutomaticEnv` call).
- Hot reload (`WatchConfig` is not called; restart to pick up changes).
- Separate per-environment files such as `config.prod.yaml`.
- Secrets management of any kind.

# Related concepts

* [Config component](/components/config.md) for the loader implementation.
* [Run locally](/operations/run-local.md) for the setup steps.
