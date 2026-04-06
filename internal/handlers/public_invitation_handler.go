package handlers

import "github.com/gofiber/fiber/v2"

type PublicInvitationHandler struct{}

func (h *PublicInvitationHandler) DetailByToken(c *fiber.Ctx) error { return ok(c, "public invitation detail", fiber.Map{"token": c.Params("inviteToken")}) }
