package routes

import (
	"hiretest-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterInvitationRoutes(v1 fiber.Router, h *handlers.Registry) {
	r := v1.Group("/invitations")
	r.Post("/", h.Invitation.Create)
	r.Get("/", h.Invitation.List)
	r.Post("/:invitationId/resend", h.Invitation.Resend)
}
