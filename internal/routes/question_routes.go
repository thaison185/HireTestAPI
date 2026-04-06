package routes

import (
	"hiretest-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterQuestionRoutes(v1 fiber.Router, h *handlers.Registry) {
	r := v1.Group("/questions")
	r.Post("/", h.Question.Create)
	r.Get("/", h.Question.List)
	r.Get("/:questionId", h.Question.Detail)
	r.Patch("/:questionId", h.Question.Update)
	r.Delete("/:questionId", h.Question.Delete)
	r.Post("/:questionId/publish", h.Question.Publish)
}
