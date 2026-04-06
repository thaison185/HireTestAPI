package routes

import (
	"hiretest-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterCandidateRoutes(v1 fiber.Router, h *handlers.Registry) {
	r := v1.Group("/candidates")
	r.Post("/", h.Candidate.Create)
	r.Get("/", h.Candidate.List)
	r.Get("/:candidateId", h.Candidate.Detail)
	r.Patch("/:candidateId", h.Candidate.Update)
}
