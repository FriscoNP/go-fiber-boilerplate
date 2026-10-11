---
type: Component
title: Handlers
description: The handler package where request handlers live; one today.
resource: internal/handler/health.go
tags: [handlers, http, patterns]
status: stable
generated: { by: human:frisco, at: 2026-10-11T00:00:00Z }
verified: { by: human:frisco, at: 2026-10-11T00:00:00Z }
---

# Contents

`internal/handler` currently holds a single file, `health.go`, with one
exported function:

```go
func HealthCheck(ctx fiber.Ctx) error {
    return ctx.JSON(fiber.Map{
        "status":  "ok",
        "message": "Fiber API is running smoothly!",
    })
}
```

# Conventions

Handlers follow the standard Fiber signature and are registered by
[routes](/components/routes.md):

* Signature `func(ctx fiber.Ctx) error`.
* Return `nil` on success; return an `error` to hand control to Fiber's
  error handler.
* Build responses with `ctx.JSON(...)` using a `fiber.Map` or a typed
  struct.
* One function per endpoint, named for the action (`HealthCheck`,
  `ListWidgets`).

# Behavioral notes

* `ctx.JSON` sets `Content-Type: application/json` and serializes with
  Fiber's default JSON encoder (Go's `encoding/json`).
* No request binding is used today; handlers that need input should call
  `ctx.BodyParser(&dst)`.
* The handler does not validate anything or consult configuration - it is
  a pure liveness response.

# Example: adding a handler

```go
package handler

type widget struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

func ListWidgets(ctx fiber.Ctx) error {
    return ctx.JSON([]widget{{ID: "w-1", Name: "Example"}})
}
```

Then wire it in `internal/routes/routes.go`:

```go
api.Get("/widgets", handler.ListWidgets)
```

# Related concepts

* [Health check](/api/health.md) - the endpoint this handler serves.
* [Routes](/components/routes.md) - registration and the URL space.
