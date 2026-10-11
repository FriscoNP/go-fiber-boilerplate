package handler_test

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/FriscoNP/go-fiber-boilerplate/internal/handler"
	"github.com/gofiber/fiber/v3"
)

func TestHealthCheck(t *testing.T) {
	app := fiber.New()
	app.Get("/api/health", handler.HealthCheck)

	resp, err := app.Test(httptest.NewRequest("GET", "/api/health", nil))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	if !strings.Contains(string(body), `"status":"ok"`) {
		t.Fatalf("unexpected body: %s", body)
	}
}
