package routes

import (
	"hiretest-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterProfileRoutes(r fiber.Router, h *handlers.Registry) {
	r.Get("/", h.Profile.Get)
}
