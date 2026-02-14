package main

import (
	"log"
	"strconv"

	"github.com/FriscoNP/go-fiber-boilerplate/internal/config"
	"github.com/FriscoNP/go-fiber-boilerplate/internal/server"
)

func main() {
	cfg := config.Load()

	app := server.New()

	appPort := strconv.Itoa(cfg.App.Port)

	if err := server.Start(app, appPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
