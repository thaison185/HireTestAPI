package routes

import (
	"hiretest-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterOrganizationRoutes(v1 fiber.Router, h *handlers.Registry) {
	r := v1.Group("/org")
	r.Get("/members", h.Org.ListMembers)
	r.Patch("/members/:memberId/roles", h.Org.UpdateMemberRoles)
}
