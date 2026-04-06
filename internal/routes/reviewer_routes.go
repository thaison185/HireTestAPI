package routes

import (
	"hiretest-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterReviewerRoutes(v1 fiber.Router, h *handlers.Registry) {
	r := v1.Group("/reviewer")
	r.Get("/submissions/queue", h.Reviewer.Queue)
	r.Get("/submissions/:submissionId", h.Reviewer.Detail)
	r.Post("/submissions/:submissionId/score", h.Reviewer.Score)
}
