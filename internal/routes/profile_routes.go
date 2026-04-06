package routes

import (
	"hiretest-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterProfileRoutes(v1 fiber.Router, h *handlers.Registry) {
	v1.Get("/profile", h.Profile.Get)
}
