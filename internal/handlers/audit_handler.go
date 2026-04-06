package handlers

import "github.com/gofiber/fiber/v2"

type AuditHandler struct{}

func (h *AuditHandler) List(c *fiber.Ctx) error { return ok(c, "audit logs", []string{}) }
