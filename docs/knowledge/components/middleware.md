---
type: Component
title: Middleware
description: Registers panic recovery and a single-line request logger globally on the Fiber app.
resource: internal/middleware/middleware.go
tags: [middleware, logging, recovery]
status: stable
generated: { by: human:frisco, at: 2026-10-11T00:00:00Z }
verified: { by: human:frisco, at: 2026-10-11T00:00:00Z }
---

# Registered middleware

```go
app.Use(recoverer.New())
app.Use(logger.New(logger.Config{
    Format: "[${time}] ${status} - ${method} ${path}\n",
}))
```

| Order | Middleware | Package |
|-------|-----------|---------|
| 1 | recover | `fiber/v3/middleware/recover` |
| 2 | logger | `fiber/v3/middleware/logger` |

`app.Use` with no path matches every method and path, so both wrap all routes.

# Recover middleware

- Uses `recover.New()` with no config, so `EnableStackTrace` is `false`. A
  panic is converted into an `error` that Fiber's error handler renders as a
  `500`, and the stack trace handler is not invoked - the panic message is
  **not** printed to stderr by default.
- Registered first, making it the outermost middleware.

# Logger middleware

- Writes a single line per request to `os.Stdout` (the default `Stream`).
- Format resolves to something like `[14:32:07] 200 - GET /api/health`, followed
  by a newline. `${time}` uses the default `TimeFormat` of `15:04:05` and the
  default `TimeZone` of `Local`.
- Colors are enabled when stdout is a TTY (`DisableColors` defaults to `false`),
  so captured output may contain ANSI escape codes.
- On the first request it snapshots `c.App().ErrorHandler` and afterwards
  manually invokes that handler whenever a downstream handler returns a
  non-nil error, then renders the log line. The configured `Format` does not
  include `${error}`, so errors are not visible in the log output even though
  they are still handled.

# Consequence of the ordering

Because recover is registered **before** logger, recover is the outer frame.
A panic in a handler unwinds through the logger's `c.Next()` call, so the
logger's post-processing never runs for that request. Panicking requests are
therefore converted to a `500` by recover but are **not** written to the access
log. Registering logger first would flip that trade-off: the line would be
logged, and recover would still catch the panic.

# Behavioral notes

- No request IDs, no body/latency capture, no log level filtering.
- No CORS, compression, authentication, rate limiting, or timeout middleware.

# Related concepts

* [Architecture](/architecture.md) for where middleware sits in the request flow.
* [Health check](/api/health.md) for the resulting log output on a live probe.
