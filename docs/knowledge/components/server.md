---
type: Component
title: Server construction
description: Builds the Fiber app and listens on the configured port.
resource: internal/server/server.go
tags: [server, fiber, lifecycle]
status: stable
generated: { by: human:frisco, at: 2026-10-11T00:00:00Z }
verified: { by: human:frisco, at: 2026-10-11T00:00:00Z }
---

# API surface

```go
func New() *fiber.App
func Start(app *fiber.App, port string) error
```

`New()` returns a fully wired `*fiber.App`. `Start()` is a thin wrapper
over `app.Listen(fmt.Sprintf(":%s", port))`, which binds all interfaces
on the given port and returns the error from Fiber if the server stops.

# Construction order

```go
app := fiber.New()
middleware.Register(app)   // 1. recover, then logger
routes.Register(app)       // 2. /api group
```

Middleware must be registered before routes so that it wraps every
handler. `fiber.New()` is called with no options, so Fiber defaults
apply: no custom logger, no custom error handler, default proxy/header
behavior.

# Startup wiring

`cmd/api/main.go` calls:

```go
cfg := config.Load()
app := server.New()
appPort := strconv.Itoa(cfg.App.Port)
server.Start(app, appPort)
```

# Behavioral notes

* **No graceful shutdown.** There is no `SIGINT`/`SIGTERM` handler and no
  `app.Shutdown()` call; the process is killed by the signal.
* **Port is passed as a string.** It is stringified in `main` and
  formatted with `:%s`, so `0` is a legal value and asks the OS for an
  ephemeral port.
* **No health/readiness gating at startup.** Nothing is loaded or checked
  between construction and `Listen`.
* **Single listening socket.** No TLS configuration, so traffic is plain
  HTTP unless terminated upstream.

# Related concepts

* [Architecture](/architecture.md) for the full startup sequence.
* [Middleware](/components/middleware.md) and
  [routes](/components/routes.md) for what `New()` wires up.
