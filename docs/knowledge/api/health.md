---
type: API Endpoint
title: Health check
description: Liveness endpoint that confirms the Fiber API process is running.
resource: /api/health
tags: [api, observability, health]
status: stable
generated: { by: human:frisco, at: 2026-10-11T00:00:00Z }
verified: { by: human:frisco, at: 2026-10-11T00:00:00Z }
---

# Request

```http
GET /api/health
```

No path parameters, query parameters, headers, or request body.

# Response

`200 OK`, `Content-Type: application/json`.

```json
{
  "status": "ok",
  "message": "Fiber API is running smoothly!"
}
```

| Field | Type | Description |
|-------|------|-------------|
| `status` | string | Always `"ok"`. |
| `message` | string | Fixed human-readable confirmation. |

# Behavior

* Always returns `200` while the process is alive; it performs no
  dependency checks (no database or downstream probing).
* Reached through the `/api` group defined in
  [routes](/components/routes.md), so the effective path includes the
  `/api` prefix.
* Implemented by `HealthCheck` in [handlers](/components/handlers.md).

# Examples

```bash
curl -s http://localhost:8080/api/health
# {"message":"Fiber API is running smoothly!","status":"ok"}
```

# See also

* [Routes](/components/routes.md) - where the path is registered.
* [Run locally](/operations/run-local.md) - starting the server to reach
  it.
