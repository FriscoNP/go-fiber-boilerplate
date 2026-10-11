---
type: Architecture
title: Architecture
description: How a request travels from main through the Fiber pipeline to a handler, and how each package participates.
tags: [architecture, request-flow, fiber]
status: stable
generated: { by: human:frisco, at: 2026-10-11T00:00:00Z }
verified: { by: human:frisco, at: 2026-10-11T00:00:00Z }
---

# Request lifecycle

```
main()                          cmd/api/main.go
  ├─ config.Load()              internal/config
  ├─ server.New()               internal/server
  │    ├─ middleware.Register() internal/middleware
  │    └─ routes.Register()     internal/routes
  │         └─ handler.*        internal/handler
  └─ server.Start(app, port)    internal/server
       └─ app.Listen(":<port>")
```

1. `main` calls `config.Load()`, which reads `internal/config/config.local.yaml`
   via viper and unmarshals it into `config.Config`. A missing or malformed file
   calls `log.Fatalf` - there is no fallback.
2. `main` builds the app with `server.New()`.
3. `server.New()` registers middleware first, then routes, so middleware wraps
   every route handler in registration order.
4. `main` derives the port with `strconv.Itoa(cfg.App.Port)` and calls
   `server.Start(app, appPort)`, which is a thin wrapper over `app.Listen`.
5. Fiber dispatches the request through the recovery middleware, then the
   logger, then the matching route handler.

# Package responsibilities

| Package | Owns | Depends on |
|---------|------|------------|
| `cmd/api` | Composition and process lifecycle | `config`, `server` |
| `internal/config` | Config schema and loading | `viper` |
| `internal/server` | App construction, listen address | `middleware`, `routes`, `fiber` |
| `internal/middleware` | Recoverer and request logger | `fiber` |
| `internal/routes` | URL space | `handler`, `fiber` |
| `internal/handler` | Response shaping | `fiber` |

# Design notes

- **Registration order matters.** Middleware is registered before routes, so
  `recover` and `logger` apply to all handlers. Adding a route in
  `server.New()` instead of `routes.Register()` would bypass that intent.
- **No application-level error type.** Handlers return `error` straight to
  Fiber, which renders Fiber's default error page.
- **Composition root is `main`.** The wiring order lives in
  `cmd/api/main.go`; the packages themselves do not import each other sideways
  except through `server`.
- **Config is loaded eagerly and fatally.** Fail-fast by design, which keeps
  the app from starting in a half-configured state.

# Related concepts

- [Components](/components/) for per-package detail.
- [Run locally](/operations/run-local.md) for the startup procedure.
