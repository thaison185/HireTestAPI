package routes

import (
	"hiretest-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterQuestionRoutes(r fiber.Router, h *handlers.Registry) {
	r.Post("/", h.Question.Create)
	r.Get("/", h.Question.List)
	r.Get("/:id", h.Question.Detail)
	r.Patch("/:id", h.Question.Update)
	r.Delete("/:id", h.Question.Delete)
	// r.Post("/:id/publish", h.Question.Publish)
}
