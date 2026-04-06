package routes

import (
	"hiretest-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuditRoutes(v1 fiber.Router, h *handlers.Registry) {
	r := v1.Group("/audit-logs")
	r.Get("/", h.Audit.List)
}
