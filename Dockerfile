FROM golang:1.26 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/api ./cmd/api

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /out/api /app/api
# Config is loaded from ./internal/config relative to the working dir.
COPY internal/config/example.config.local.yaml /app/internal/config/config.local.yaml
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/api"]
