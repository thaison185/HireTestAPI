package routes

import (
	"hiretest-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterProfileRoutes(v1 fiber.Router, h *handlers.Registry) {
	r := v1.Group("/profile")
	r.Get("/", h.Profile.Get)
}
