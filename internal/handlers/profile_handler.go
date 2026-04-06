package handlers

import "github.com/gofiber/fiber/v2"

type ProfileHandler struct{}

func (h *ProfileHandler) Get(c *fiber.Ctx) error { return ok(c, "current user", fiber.Map{"id": "me"}) }
