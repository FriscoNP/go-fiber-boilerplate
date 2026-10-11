---
okf_version: "0.2"
---

# Go Fiber Boilerplate Knowledge

A portable OKF knowledge bundle describing this repository for humans and AI
agents. Clone the repository and read any file directly; no tooling is needed.

# Start Here

* [Overview](/overview.md) - What this boilerplate is and how it is laid out.
* [Architecture](/architecture.md) - Request flow from `main` through the Fiber pipeline.
* [OKF conformance notes](/standards/okf.md) - The Open Knowledge Format rules this bundle follows.

# API

* [API](/api/) - HTTP endpoints exposed by the service.
* [Health check](/api/health.md) - `GET /api/health` liveness endpoint.

# Components

* [Components](/components/) - The internal packages that make up the service.
* [Config](/components/config.md) - viper-backed YAML configuration loader.
* [Server](/components/server.md) - Fiber app construction and startup.
* [Middleware](/components/middleware.md) - Recovery and request logging.
* [Routes](/components/routes.md) - Route registration and the `/api` group.
* [Handlers](/components/handlers.md) - Request handler conventions.

# Operations

* [Operations](/operations/) - How to run, configure, and package the service.
* [Run locally](/operations/run-local.md) - Build and run with the Makefile.
* [Configuration](/operations/configuration.md) - Config file location, keys, and failure modes.
* [Containerization](/operations/containerization.md) - Dockerfile and Compose status.

# Standards

* [Standards](/standards/) - Conventions this bundle and repository follow.
* [Open Knowledge Format](/standards/okf.md) - The OKF v0.2 spec summary and local rules.
