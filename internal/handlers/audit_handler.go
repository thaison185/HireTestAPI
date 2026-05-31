package handlers

import (
	"hiretest-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AuditHandler struct {
	Service *services.AuditService
}

func NewAuditHandler(s *services.AuditService) *AuditHandler {
	return &AuditHandler{Service: s}
}

func (h *AuditHandler) List(c *fiber.Ctx) error { return ok(c, "audit logs", []string{}) }
