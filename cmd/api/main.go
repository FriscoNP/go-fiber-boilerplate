package main

import (
	"log"

	"github.com/FriscoNP/go-fiber-boilerplate/internal/config"
	"github.com/FriscoNP/go-fiber-boilerplate/internal/server"
)

func main() {
	cfg := config.Load()

	app := server.New()

	if err := server.Start(app, cfg.AppPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
