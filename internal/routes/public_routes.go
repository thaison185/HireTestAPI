package routes

import (
	"hiretest-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterPublicRoutes(v1 fiber.Router, h *handlers.Registry) {
	r := v1.Group("/public")
	r.Get("/invitations/:inviteToken", h.PublicInv.DetailByToken)
	r.Post("/submissions/start", h.PublicSub.Start)
	r.Patch("/submissions/:submissionId/autosave", h.PublicSub.Autosave)
	r.Post("/submissions/:submissionId/final-submit", h.PublicSub.FinalSubmit)
}
