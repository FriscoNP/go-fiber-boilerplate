---
type: Playbook
title: Containerization
description: Status and intended content of the Dockerfile and Compose file.
tags: [operations, docker, deployment]
status: draft
generated: { by: human:frisco, at: 2026-10-11T00:00:00Z }
sources:
  - id: docker-docs
    resource: https://docs.docker.com/build/building/multi-stage/
    title: Docker multi-stage builds
    author: team:docker
---

# Status: placeholder

Both `Dockerfile` and `docker-compose.yml` exist in the repository but
are **zero bytes**. Nothing builds or runs a container today. This
concept is `status: draft` for that reason.[^docker-docs]

| File | Size | Effect |
|------|------|--------|
| `Dockerfile` | 0 bytes | `docker build .` fails on an empty Dockerfile. |
| `docker-compose.yml` | 0 bytes | `docker compose up` fails to parse an empty file. |

# Intended shape

The following is the target design, not current behavior:

```dockerfile
FROM golang:1.25 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/api ./cmd/api

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /out/api ./api
# config.local.yaml must be copied in or mounted, since the loader
# reads ./internal/config/config.local.yaml relative to WORKDIR.
ENTRYPOINT ["./api"]
```

```yaml
services:
  api:
    build: .
    ports:
      - "8080:8080"
    volumes:
      - ./internal/config/config.local.yaml:/app/internal/config/config.local.yaml:ro
```

# Blocking constraint

`config.Load()` hard-codes the relative path `./internal/config` and
calls `log.Fatalf` when the file is absent. Any container image must
therefore provide `internal/config/config.local.yaml` at that path
relative to the container's working directory, or the container exits
immediately. Resolving this - by honoring an environment variable or a
`--config` flag - is a prerequisite for a clean image, and is out of
scope for this bundle.

# Work plan

1. Decide how configuration reaches the container (mount vs. baked in
   vs. env-var override).
2. Write the multi-stage `Dockerfile`.
3. Write `docker-compose.yml` with port mapping and a health check on
   `GET /api/health`.
4. Add a non-root runtime user and a `HEALTHCHECK` instruction.
5. Change this concept to `status: stable`.

# Related concepts

* [Configuration](/operations/configuration.md) for the file the image
  must supply.
* [Run locally](/operations/run-local.md) for the non-container path.

[^docker-docs]: Docker multi-stage builds
