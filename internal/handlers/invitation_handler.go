package handlers

import "github.com/gofiber/fiber/v2"

type InvitationHandler struct{}

func (h *InvitationHandler) Create(c *fiber.Ctx) error { return ok(c, "invitation created", nil) }
func (h *InvitationHandler) List(c *fiber.Ctx) error { return ok(c, "invitation list", []string{}) }
func (h *InvitationHandler) Resend(c *fiber.Ctx) error { return ok(c, "invitation resent", nil) }
