---
type: Component
title: Configuration loader
description: Reads internal/config/config.local.yaml with viper into Config.
resource: internal/config/config.go
tags: [config, viper, yaml]
status: stable
generated: { by: human:frisco, at: 2026-10-11T00:00:00Z }
verified: { by: human:frisco, at: 2026-10-11T00:00:00Z }
sources:
  - id: viper-unmarshal
    resource: https://github.com/spf13/viper
    title: spf13/viper repository
    author: human:spf13
    last_modified: 2026-06-01T00:00:00Z
---

# Location

`internal/config/config.go`, plus `internal/config/example.config.local.yaml`
as the checked-in template.

# Schema

```go
type Config struct {
    App App `yaml:"app"`
}

type App struct {
    Name string `yaml:"name"`
    Port int    `yaml:"port"`
    Env  string `yaml:"env"`
}
```

| Key | Type | Example | Meaning |
|-----|------|---------|---------|
| `app.name` | string | `Golang Service with Fiber` | Human-readable service name. Currently unused at runtime. |
| `app.port` | int | `8080` | TCP port passed to `app.Listen`. |
| `app.env` | string | `local` | Environment label. Currently unused at runtime. |

# Loading procedure

`Load()` performs four steps and returns `Config` by value:

1. `viper.SetConfigName("config.local")` - the filename without
   extension.
2. `viper.SetConfigType("yaml")` - forces the YAML decoder.
3. `viper.AddConfigPath("./internal/config")` - a **relative** search
   path.
4. `viper.ReadInConfig()` followed by `viper.Unmarshal(&config)`.

Both error branches call `log.Fatalf`, so the process exits rather than
starting with partial configuration.

# Behavioral notes

* **Path resolution is relative to the process working directory.** The
  documented invocation is `make run` from the repository root, because
  `./internal/config` only resolves when the working directory is the
  repo root. Running the binary from another directory fails at startup.
* **`config.local.yaml` is gitignored.** A fresh clone has no config
  file, so `make run` fails immediately until
  `example.config.local.yaml` is copied to `config.local.yaml`.
* **The `yaml:` struct tags are decorative.** viper's `Unmarshal`
  decodes with `go-viper/mapstructure/v2`, whose default tag name is
  `mapstructure`, and viper does not override it.[^viper-unmarshal] Field
  matching therefore falls back to case-insensitive field-name matching,
  which happens to align with the `app.*` keys. Rename a Go field and the
  key stops resolving - add a `mapstructure:` tag if explicit mapping is
  wanted.
* **Decoding is weakly typed.** viper configures `WeaklyTypedInput: true`,
  so a quoted `"8080"` would still coerce to the `int` port.
* **Not environment-variable driven.** No `viper.AutomaticEnv` or
  `SetEnvPrefix` call exists, so `APP_PORT`-style overrides are not read.
* **No reload.** `fsnotify` is present as an indirect dependency, but no
  `WatchConfig` call is made; changes require a restart.

# Related concepts

* [Configuration operations](/operations/configuration.md) for the file
  format and setup steps.
* [Architecture](/architecture.md) for where loading happens in startup.

[^viper-unmarshal]: spf13/viper repository
