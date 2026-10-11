BIN := bin/api

.PHONY: run build test lint fmt vet tidy docker-build docker-up

run:
	go run ./cmd/api

build:
	CGO_ENABLED=0 go build -o $(BIN) ./cmd/api

test:
	go test -race -cover ./...

lint:
	golangci-lint run

fmt:
	gofmt -w .

vet:
	go vet ./...

tidy:
	go mod tidy

docker-build:
	docker build -t go-fiber-boilerplate .

docker-up:
	docker compose up --build
