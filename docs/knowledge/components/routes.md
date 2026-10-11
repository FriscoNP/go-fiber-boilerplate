---
type: Component
title: Routes
description: Owns the URL space by creating the /api group and registering every application endpoint.
resource: internal/routes/routes.go
tags: [routes, http, api]
status: stable
generated: { by: human:frisco, at: 2026-10-11T00:00:00Z }
verified: { by: human:frisco, at: 2026-10-11T00:00:00Z }
---

# Implementation

```go
func Register(app *fiber.App) {
    api := app.Group("/api")
    api.Get("/health", handler.HealthCheck)
}
```

`Register` is the single entry point called by
[server.New()](/components/server.md). All application routes belong inside
the group created here.

# URL space

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| `GET` | `/api/health` | `handler.HealthCheck` | [Health check](/api/health.md) |

The group prefix `/api` is applied by Fiber, so handlers always see the full
path `/api/health`.

# Extending

Add a new endpoint by appending to the group:

```go
api.Get("/widgets", handler.ListWidgets)
api.Post("/widgets", handler.CreateWidget)
```

For a versioned or nested prefix, chain another group:

```go
v1 := api.Group("/v1")
v1.Get("/widgets", handler.ListWidgets)
```

# Behavioral notes

- No `HEAD`-only, `OPTIONS`, or `Static` routes are registered, and no 404
  handler is installed, so unmatched paths fall through to Fiber's default
  `404` error page.
- Middleware registered in [middleware](/components/middleware.md) applies to
  these routes because it was registered first at the app level.
- Only `GET` is implemented today; other methods against `/api/health` return
  Fiber's default `405`-style error page rather than a JSON body.

# Related concepts

* [Handlers](/components/handlers.md) for what the registered functions do.
* [Architecture](/architecture.md) for registration order.
