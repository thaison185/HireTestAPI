package routes

import (
	"hiretest-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterTestRoutes(v1 fiber.Router, h *handlers.Registry) {
	r := v1.Group("/tests")
	r.Post("/", h.Test.Create)
	r.Get("/", h.Test.List)
	r.Get("/:testId", h.Test.Detail)
	r.Patch("/:testId", h.Test.Update)
	r.Delete("/:testId", h.Test.Delete)
	r.Post("/:testId/questions", h.Test.AddQuestion)
	r.Post("/:testId/publish", h.Test.Publish)
}
