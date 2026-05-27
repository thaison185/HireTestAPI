package routes

import (
	"hiretest-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(r fiber.Router, h *handlers.Registry) {
	r.Post("/login", h.Auth.Login)
	r.Post("/refresh", h.Auth.Refresh)
	r.Post("/logout", h.Auth.Logout)
}
